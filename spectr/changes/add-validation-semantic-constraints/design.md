# Design: Semantic Validation Constraints

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                    Validator                            │
│  ┌──────────────┐       ┌──────────────┐              │
│  │   Schema     │       │  Semantic    │              │
│  │  Validator   │  -->  │  Validator   │              │
│  └──────────────┘       └──────────────┘              │
│         │                        │                      │
│         v                        v                      │
│  ┌────────────────────────────────────────┐           │
│  │      ValidationContext                 │           │
│  │  - Document tree                       │           │
│  │  - ID map                              │           │
│  │  - Relationship map                    │           │
│  │  - Element path tracking               │           │
│  └────────────────────────────────────────┘           │
└─────────────────────────────────────────────────────────┘
                         │
                         v
              ┌────────────────────┐
              │  Constraint Types  │
              ├────────────────────┤
              │ AttributeValueSet  │
              │ AttributeValueRange│
              │ AttributeMutualExcl│
              │ ParentTypeConstr   │
              │ ReferenceExist     │
              │ ...11 more         │
              └────────────────────┘
```

## Core Interfaces

### SemanticConstraint Interface

```go
// openxml/validation/semantic.go

package validation

// SemanticConstraint defines contract for semantic validation rules
type SemanticConstraint interface {
    // Validate checks constraint against element
    Validate(ctx *ValidationContext, elem CommonWrapper) []ValidationError
    
    // Description returns human-readable constraint description
    Description() string
    
    // Level returns validation strictness level
    Level() SemanticValidationLevel
}

// SemanticValidationLevel controls validation strictness
type SemanticValidationLevel int

const (
    SemanticLevelNone SemanticValidationLevel = iota
    SemanticLevelPartial  // Common constraints only
    SemanticLevelFull     // All constraints
    SemanticLevelCustom   // User-selected constraints
)
```

### ValidationContext

```go
// ValidationContext provides document state during validation
type ValidationContext struct {
    // Document being validated
    Document CommonWrapper
    
    // ID map for reference resolution
    idMap map[string]CommonWrapper
    
    // Relationship map
    relationships map[string]*packaging.Relationship
    
    // Current element path (for error reporting)
    currentPath []CommonWrapper
    
    // Validation options
    options ValidationOptions
    
    // Cached data
    cache map[string]interface{}
}

// NewValidationContext creates context with ID/relationship indexing
func NewValidationContext(doc CommonWrapper, opts ValidationOptions) *ValidationContext {
    ctx := &ValidationContext{
        Document:      doc,
        idMap:         make(map[string]CommonWrapper),
        relationships: make(map[string]*packaging.Relationship),
        options:       opts,
        cache:         make(map[string]interface{}),
    }
    
    // Build ID map by walking document tree
    ctx.buildIDMap(doc)
    
    // Build relationship map from package
    ctx.buildRelationshipMap(doc)
    
    return ctx
}

// GetPath returns XPath-style path to element
func (ctx *ValidationContext) GetPath(elem CommonWrapper) string {
    // Example: /document[1]/body[1]/p[3]/r[2]
    path := ""
    current := elem
    for current != nil {
        path = "/" + current.LocalName() + indexInParent(current) + path
        current = current.Parent()
    }
    return path
}

// ResolveID looks up element by ID attribute
func (ctx *ValidationContext) ResolveID(id string) (CommonWrapper, bool) {
    elem, exists := ctx.idMap[id]
    return elem, exists
}

// GetRelationship retrieves relationship by ID
func (ctx *ValidationContext) GetRelationship(id string) (*packaging.Relationship, bool) {
    rel, exists := ctx.relationships[id]
    return rel, exists
}

func (ctx *ValidationContext) buildIDMap(elem CommonWrapper) {
    // Extract ID from common ID attributes
    if id := elem.GetAttribute("id"); id != nil {
        ctx.idMap[id.String()] = elem
    }
    if id := elem.GetAttribute("Id"); id != nil {
        ctx.idMap[id.String()] = elem
    }
    
    // Recursively process children
    for _, child := range elem.GetChildren() {
        ctx.buildIDMap(child)
    }
}

func (ctx *ValidationContext) buildRelationshipMap(doc CommonWrapper) {
    // Get package from document
    pkg := getPackageFromDocument(doc)
    if pkg == nil {
        return
    }
    
    // Index all relationships
    for _, part := range pkg.GetParts() {
        for _, rel := range part.GetRelationships() {
            ctx.relationships[rel.ID] = rel
        }
    }
}
```

## Constraint Implementations

### AttributeValueSetConstraint

```go
// AttributeValueSetConstraint validates attribute is in allowed set
type AttributeValueSetConstraint struct {
    AttributeName string
    AllowedValues []string
    CaseSensitive bool
}

func (c *AttributeValueSetConstraint) Validate(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    attr := elem.GetAttribute(c.AttributeName)
    if attr == nil {
        // Absent attribute is valid
        return nil
    }
    
    value := attr.String()
    
    // Check if value in allowed set
    for _, allowed := range c.AllowedValues {
        if c.CaseSensitive {
            if value == allowed {
                return nil
            }
        } else {
            if strings.EqualFold(value, allowed) {
                return nil
            }
        }
    }
    
    return []ValidationError{{
        ErrorType:   SemanticError,
        Description: fmt.Sprintf(
            "Attribute '%s' has invalid value '%s'. Allowed values: %v",
            c.AttributeName, value, c.AllowedValues,
        ),
        Path: ctx.GetPath(elem),
        Severity: ErrorSeverity,
    }}
}

func (c *AttributeValueSetConstraint) Description() string {
    return fmt.Sprintf("Attribute '%s' must be one of: %v",
        c.AttributeName, c.AllowedValues)
}

func (c *AttributeValueSetConstraint) Level() SemanticValidationLevel {
    return SemanticLevelFull
}
```

### AttributeValueRangeConstraint

```go
// AttributeValueRangeConstraint validates numeric attribute in range
type AttributeValueRangeConstraint struct {
    AttributeName string
    MinValue      float64
    MaxValue      float64
    Inclusive     bool // true = [min,max], false = (min,max)
}

func (c *AttributeValueRangeConstraint) Validate(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    attr := elem.GetAttribute(c.AttributeName)
    if attr == nil {
        return nil
    }
    
    value, err := strconv.ParseFloat(attr.String(), 64)
    if err != nil {
        return []ValidationError{{
            ErrorType:   SemanticError,
            Description: fmt.Sprintf(
                "Attribute '%s' must be numeric, got '%s'",
                c.AttributeName, attr.String(),
            ),
            Path:     ctx.GetPath(elem),
            Severity: ErrorSeverity,
        }}
    }
    
    valid := false
    if c.Inclusive {
        valid = value >= c.MinValue && value <= c.MaxValue
    } else {
        valid = value > c.MinValue && value < c.MaxValue
    }
    
    if !valid {
        bounds := "(" + fmt.Sprint(c.MinValue) + ", " + fmt.Sprint(c.MaxValue) + ")"
        if c.Inclusive {
            bounds = "[" + fmt.Sprint(c.MinValue) + ", " + fmt.Sprint(c.MaxValue) + "]"
        }
        
        return []ValidationError{{
            ErrorType:   SemanticError,
            Description: fmt.Sprintf(
                "Attribute '%s' value %.2f out of range %s",
                c.AttributeName, value, bounds,
            ),
            Path:     ctx.GetPath(elem),
            Severity: ErrorSeverity,
        }}
    }
    
    return nil
}

func (c *AttributeValueRangeConstraint) Description() string {
    return fmt.Sprintf("Attribute '%s' must be in range [%.2f, %.2f]",
        c.AttributeName, c.MinValue, c.MaxValue)
}

func (c *AttributeValueRangeConstraint) Level() SemanticValidationLevel {
    return SemanticLevelFull
}
```

### AttributeMutualExclusiveConstraint

```go
// AttributeMutualExclusiveConstraint validates only one attribute present
type AttributeMutualExclusiveConstraint struct {
    Attributes []string
}

func (c *AttributeMutualExclusiveConstraint) Validate(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    var presentAttributes []string
    
    for _, attrName := range c.Attributes {
        if elem.GetAttribute(attrName) != nil {
            presentAttributes = append(presentAttributes, attrName)
        }
    }
    
    if len(presentAttributes) <= 1 {
        return nil
    }
    
    return []ValidationError{{
        ErrorType:   SemanticError,
        Description: fmt.Sprintf(
            "Attributes %v are mutually exclusive, but found: %v",
            c.Attributes, presentAttributes,
        ),
        Path:     ctx.GetPath(elem),
        Severity: ErrorSeverity,
    }}
}

func (c *AttributeMutualExclusiveConstraint) Description() string {
    return fmt.Sprintf("Only one of %v may be present", c.Attributes)
}

func (c *AttributeMutualExclusiveConstraint) Level() SemanticValidationLevel {
    return SemanticLevelFull
}
```

### ReferenceExistConstraint

```go
// ReferenceExistConstraint validates referenced ID exists
type ReferenceExistConstraint struct {
    AttributeName string
    ReferenceType string // "bookmark", "style", "numbering", etc.
}

func (c *ReferenceExistConstraint) Validate(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    attr := elem.GetAttribute(c.AttributeName)
    if attr == nil {
        return nil
    }
    
    refID := attr.String()
    _, exists := ctx.ResolveID(refID)
    
    if !exists {
        return []ValidationError{{
            ErrorType:   SemanticError,
            Description: fmt.Sprintf(
                "Referenced %s '%s' does not exist in document",
                c.ReferenceType, refID,
            ),
            Path:     ctx.GetPath(elem),
            Severity: ErrorSeverity,
        }}
    }
    
    return nil
}

func (c *ReferenceExistConstraint) Description() string {
    return fmt.Sprintf("Attribute '%s' must reference existing %s",
        c.AttributeName, c.ReferenceType)
}

func (c *ReferenceExistConstraint) Level() SemanticValidationLevel {
    return SemanticLevelFull
}
```

## Validator Integration

```go
// Validator extended to support semantic validation
type Validator struct {
    schemaValidator     *SchemaValidator
    semanticConstraints []SemanticConstraint
    semanticLevel       SemanticValidationLevel
}

// NewValidator creates validator with options
func NewValidator(opts ...ValidatorOption) *Validator {
    v := &Validator{
        schemaValidator: NewSchemaValidator(),
        semanticLevel:   SemanticLevelNone,
    }
    
    for _, opt := range opts {
        opt(v)
    }
    
    return v
}

// ValidatorOption configures validator
type ValidatorOption func(*Validator)

// WithSemanticConstraints enables semantic validation
func WithSemanticConstraints() ValidatorOption {
    return func(v *Validator) {
        v.semanticLevel = SemanticLevelFull
    }
}

// WithSemanticLevel sets validation level
func WithSemanticLevel(level SemanticValidationLevel) ValidatorOption {
    return func(v *Validator) {
        v.semanticLevel = level
    }
}

// Validate performs schema and semantic validation
func (v *Validator) Validate(doc CommonWrapper) []ValidationError {
    var errors []ValidationError
    
    // Schema validation first
    errors = append(errors, v.schemaValidator.Validate(doc)...)
    
    // Semantic validation if enabled
    if v.semanticLevel != SemanticLevelNone {
        ctx := NewValidationContext(doc, ValidationOptions{
            Level: v.semanticLevel,
        })
        errors = append(errors, v.validateSemantics(ctx, doc)...)
    }
    
    return errors
}

func (v *Validator) validateSemantics(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    var errors []ValidationError
    
    // Get constraints for this element type
    constraints := elem.SemanticConstraints()
    
    // Run applicable constraints
    for _, constraint := range constraints {
        if shouldRunConstraint(constraint, ctx.options.Level) {
            errors = append(errors, constraint.Validate(ctx, elem)...)
        }
    }
    
    // Recursively validate children
    for _, child := range elem.GetChildren() {
        errors = append(errors, v.validateSemantics(ctx, child)...)
    }
    
    return errors
}

func shouldRunConstraint(
    constraint SemanticConstraint,
    level SemanticValidationLevel,
) bool {
    switch level {
    case SemanticLevelNone:
        return false
    case SemanticLevelPartial:
        return constraint.Level() <= SemanticLevelPartial
    case SemanticLevelFull:
        return true
    case SemanticLevelCustom:
        // Check custom configuration
        return true
    default:
        return false
    }
}
```

## Element Integration

```go
// openxml/element.go

// CommonWrapper extended with semantic constraints
type CommonWrapper interface {
    // ... existing methods
    
    // SemanticConstraints returns constraints for this element type
    SemanticConstraints() []validation.SemanticConstraint
}

// Example: Paragraph element
// wordprocessing/elements/paragraph.go

func (p *Paragraph) SemanticConstraints() []validation.SemanticConstraint {
    return []validation.SemanticConstraint{
        // Indentation attributes are mutually exclusive with jc (justification)
        &validation.AttributeMutualExclusiveConstraint{
            Attributes: []string{"left", "right", "hanging"},
        },
        
        // Spacing must be in valid range (0-1584 points)
        &validation.AttributeValueRangeConstraint{
            AttributeName: "spacing",
            MinValue:      0,
            MaxValue:      1584,
            Inclusive:     true,
        },
        
        // Parent must be Body or TableCell
        &validation.ParentTypeConstraint{
            AllowedParents: []reflect.Type{
                reflect.TypeOf(&Body{}),
                reflect.TypeOf(&TableCell{}),
            },
        },
        
        // Style reference must exist
        &validation.ReferenceExistConstraint{
            AttributeName: "pStyle",
            ReferenceType: "style",
        },
    }
}
```

## Performance Considerations

### ID Map Caching
```go
// Cache ID map per document to avoid rebuilding
type validationCache struct {
    mu     sync.RWMutex
    idMaps map[*Document]map[string]CommonWrapper
}

var globalCache = &validationCache{
    idMaps: make(map[*Document]map[string]CommonWrapper),
}

func (ctx *ValidationContext) getOrBuildIDMap(doc *Document) map[string]CommonWrapper {
    globalCache.mu.RLock()
    idMap, exists := globalCache.idMaps[doc]
    globalCache.mu.RUnlock()
    
    if exists {
        return idMap
    }
    
    // Build and cache
    globalCache.mu.Lock()
    defer globalCache.mu.Unlock()
    
    idMap = make(map[string]CommonWrapper)
    ctx.buildIDMapInto(doc, idMap)
    globalCache.idMaps[doc] = idMap
    
    return idMap
}
```

### Parallel Validation
```go
// Validate subtrees in parallel
func (v *Validator) validateSemanticsParallel(
    ctx *ValidationContext,
    elem CommonWrapper,
) []ValidationError {
    children := elem.GetChildren()
    if len(children) < 4 {
        // Sequential for small trees
        return v.validateSemantics(ctx, elem)
    }
    
    // Parallel validation
    type result struct {
        errors []ValidationError
    }
    
    results := make(chan result, len(children))
    
    for _, child := range children {
        go func(c CommonWrapper) {
            errors := v.validateSemantics(ctx, c)
            results <- result{errors: errors}
        }(child)
    }
    
    var allErrors []ValidationError
    for range children {
        r := <-results
        allErrors = append(allErrors, r.errors...)
    }
    
    return allErrors
}
```

## Testing Strategy

### Unit Tests
```go
func TestAttributeValueSetConstraint(t *testing.T) {
    tests := []struct {
        name      string
        value     string
        allowed   []string
        wantError bool
    }{
        {"valid value", "left", []string{"left", "right", "center"}, false},
        {"invalid value", "invalid", []string{"left", "right"}, true},
        {"case insensitive", "LEFT", []string{"left", "right"}, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            elem := createMockElement(tt.value)
            constraint := &AttributeValueSetConstraint{
                AttributeName: "align",
                AllowedValues: tt.allowed,
                CaseSensitive: false,
            }
            
            ctx := NewValidationContext(elem, ValidationOptions{})
            errors := constraint.Validate(ctx, elem)
            
            if tt.wantError {
                assert.NotEmpty(t, errors)
            } else {
                assert.Empty(t, errors)
            }
        })
    }
}
```

---

**Last Updated**: 2026-01-24  
**Status**: Design approved
