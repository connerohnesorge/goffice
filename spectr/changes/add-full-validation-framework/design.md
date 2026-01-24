# Design: Comprehensive Validation Framework

## Overview
Complete ECMA-376 validation with semantic rules, custom validators, repair mode, and detailed error reporting.

## Core Types & Interfaces

### Validation Framework
```go
type ValidationContext struct {
    Strict         bool
    EnableRepair   bool
    MaxErrors      int
    ValidateLevels []SeverityLevel
}

type SeverityLevel int
const (
    Error SeverityLevel = iota
    Warning
    Info
)

type ValidationError interface {
    Severity() SeverityLevel
    Code() string
    Message() string
    Location() Location
    Suggestion() string
}

type Location struct {
    PartURI string
    Path    string
    Line    int
    Column  int
}
```

### Validation Rules
```go
type ValidationRule interface {
    Name() string
    Describe() string
    Validate(Element, ValidationContext) []ValidationError
    CanRepair() bool
    Repair(Element) error
}

type SchemaValidationRule interface {
    ValidationRule
    ElementName() string
    RequiredAttributes() []string
    AllowedChildren() []string
    Cardinality() CardinalityRule
}

type CustomValidator struct {
    Rules map[string]ValidationRule
}
```

### Validators
```go
type DocumentValidator interface {
    Validate(interface{}) ValidationResult
    ValidateStrict() ValidationResult
    RepairAndValidate() (ValidationResult, error)
}

type ValidationResult struct {
    IsValid  bool
    Errors   []ValidationError
    Warnings []ValidationError
    Infos    []ValidationError
}
```

## Architectural Decisions

### 1. Rule-Based Validation
**Decision**: Composable rules with composition

**Rationale**:
- Independent testing
- Enable/disable rules
- Custom rules easily added
- Repair logic encapsulated

### 2. Severity Levels
**Decision**: Three-level (error, warning, info)

**Rationale**:
- Distinguish breaking vs non-breaking
- Optimization suggestions
- Spec violation detection

### 3. Repair Mode
**Decision**: Built-in repair for common issues

**Rationale**:
- Auto-fix obvious problems
- Preserve structure
- User control
- Manual fixing reduced

### 4. Namespace Validation
**Decision**: Separate transitional and strict

**Rationale**:
- Two conformance levels
- Different rules per level
- Some elements transitional-only

## Validators Implementation

- Schema validation (elements, attributes, cardinality)
- Semantic validation (relationships, cross-references, IDs)
- Performance validation (element/relationship limits)
- Namespace validation (transitional vs strict)

## Implementation Strategy

### Phase 1: Core Framework (1.5 weeks)
- [ ] Validation interfaces
- [ ] Error collection
- [ ] Rule infrastructure
- [ ] Registry and runner

### Phase 2: Schema Rules (1.5 weeks)
- [ ] Element validation
- [ ] Attribute validation
- [ ] Cardinality checking
- [ ] Type validation

### Phase 3: Semantic Rules (1.5 weeks)
- [ ] Relationship validation
- [ ] Cross-reference validation
- [ ] ID uniqueness
- [ ] Required relationships

### Phase 4: Repair & Advanced (1.5 weeks)
- [ ] Repair mode
- [ ] Namespace validation
- [ ] Constraint validation
- [ ] Performance validation

### Phase 5: Integration (1 week)
- [ ] Document API hooks
- [ ] Incremental validation
- [ ] Examples

## Error Handling

```go
type ValidationErrorImpl struct {
    severity   SeverityLevel
    code       string
    message    string
    location   Location
    suggestion string
}
```

## Performance Considerations

1. Lazy validation (on demand)
2. Incremental validation (modified parts)
3. Schema caching
4. Parallel part validation
5. Early exit after max errors
