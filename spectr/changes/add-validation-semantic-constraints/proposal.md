# Add Validation Semantic Constraints

## Overview
Implement comprehensive semantic validation constraints beyond schema validation, matching Open-XML-SDK's `DocumentFormat.OpenXml.Framework/Validation/Semantic/` namespace. This includes attribute relationship constraints, value range constraints, parent-child constraints, and cross-reference validation.

## Motivation
Current goffice validation only covers schema validation (element structure, data types, cardinality). The Open-XML-SDK provides 20+ semantic constraint types that catch logical errors like:
- Invalid attribute combinations (mutually exclusive attributes)
- Out-of-range values (percentages >100%, negative dimensions)
- Broken references (missing relationship targets, invalid bookmarks)
- Incorrect parent-child relationships
- Duplicate IDs or values that should be unique

These semantic errors pass schema validation but produce corrupt or invalid documents that Office applications reject.

## Goals

### Constraint Types to Implement
1. **AttributeValueSetConstraint** - Attribute value must be in allowed set
2. **AttributeValueRangeConstraint** - Attribute value must be in numeric range
3. **AttributeMutualExclusive** - Multiple attributes cannot coexist
4. **AttributeValueLessEqualToAnother** - Attribute A ≤ Attribute B
5. **AttributeValueConditionToAnother** - If A=X then B=Y
6. **AttributeAbsentConditionToNonValue** - If A absent, B cannot equal X
7. **AttributeAbsentConditionToValue** - If A absent, B must equal X
8. **AttributeRequiredConditionToValue** - If A=X, B is required
9. **AttributePairConstraint** - Attributes must appear together
10. **AttributeMinMaxConstraint** - Min attribute ≤ Max attribute
11. **AttributeValuePatternConstraint** - Attribute matches regex pattern
12. **UniqueAttributeValueConstraint** - Attribute values unique across elements
13. **ParentTypeConstraint** - Element's parent must be specific type
14. **RelationshipTypeConstraint** - Relationship must be specific type
15. **RelationshipExistConstraint** - Required relationship must exist
16. **ReferenceExistConstraint** - Referenced ID must exist in document

### Validation Infrastructure
- **SemanticConstraint** base interface/type
- **SemanticValidationLevel** enum (Full, Partial, None, Custom)
- Constraint registration per element type
- Configurable validation levels
- Clear error messages with XPath-style element paths
- Collect all errors (don't stop on first error)

### API Design
```go
// Validate with semantic constraints
validator := validation.NewValidator(validation.WithSemanticConstraints())
errors := validator.Validate(doc)

// Validate specific constraint levels
errors := validator.Validate(doc, validation.SemanticLevelFull)

// Custom constraint configuration
validator := validation.NewValidator(
    validation.WithConstraints(
        validation.AttributeMutualExclusive,
        validation.AttributeValueRange,
    ),
)
```

## Non-Goals
- Business logic validation (document-specific rules)
- Custom constraint languages or DSLs
- Real-time validation during document construction
- Auto-correction of validation errors
- Performance optimization for streaming validation (can be future enhancement)

## Dependencies
- **Requires**: `openxml/validation` framework (existing)
- **Requires**: Element metadata system for constraint definitions
- **Blocks**: Comprehensive validation test suite
- **Related**: Strict namespace support (different constraints for strict mode)

## Technical Approach

### 1. Constraint Definition
Each element type defines its semantic constraints:

```go
type Paragraph struct {
    // ... existing fields
}

func (p *Paragraph) SemanticConstraints() []validation.SemanticConstraint {
    return []validation.SemanticConstraint{
        validation.AttributeMutualExclusive("left", "right"),
        validation.AttributeValueRange("spacing", 0, 1584),
        validation.ParentTypeConstraint(Body{}, TableCell{}),
    }
}
```

### 2. Constraint Implementation
Each constraint type implements the `SemanticConstraint` interface:

```go
type SemanticConstraint interface {
    Validate(ctx *ValidationContext, element CommonWrapper) []ValidationError
    Description() string
    Level() SemanticValidationLevel
}

type AttributeValueRangeConstraint struct {
    AttributeName string
    MinValue, MaxValue int
}

func (c *AttributeValueRangeConstraint) Validate(ctx *ValidationContext, elem CommonWrapper) []ValidationError {
    val := elem.GetAttribute(c.AttributeName)
    if val == nil {
        return nil // Absent is valid
    }
    
    intVal := val.Int()
    if intVal < c.MinValue || intVal > c.MaxValue {
        return []ValidationError{{
            ErrorType: validation.SemanticError,
            Description: fmt.Sprintf("Attribute %s value %d out of range [%d, %d]",
                c.AttributeName, intVal, c.MinValue, c.MaxValue),
            Path: ctx.GetPath(elem),
        }}
    }
    return nil
}
```

### 3. Validation Context
Track document state during validation:

```go
type ValidationContext struct {
    Document     CommonWrapper
    IDMap        map[string]CommonWrapper  // For reference validation
    RelationshipMap map[string]*Relationship
    CurrentPath  []CommonWrapper
    Options      ValidationOptions
}

func (c *ValidationContext) GetPath(elem CommonWrapper) string {
    // Return XPath-like path: /document/body/p[1]/r[2]
}

func (c *ValidationContext) ResolveID(id string) (CommonWrapper, bool) {
    elem, exists := c.IDMap[id]
    return elem, exists
}
```

### 4. Code Generation
Generate constraint definitions from Open-XML-SDK metadata:

```bash
# Extract constraint definitions from SDK
go run ./cmd/gen-validation-constraints \
    -sdk ../Open-XML-SDK \
    -output openxml/validation/generated_constraints.go
```

## Implementation Plan

### Phase 1: Infrastructure (1 week)
- [ ] Define SemanticConstraint interface
- [ ] Implement ValidationContext with ID/relationship tracking
- [ ] Update Validator to run semantic constraints
- [ ] Add SemanticValidationLevel enum

### Phase 2: Core Constraints (2 weeks)
- [ ] AttributeValueSetConstraint
- [ ] AttributeValueRangeConstraint
- [ ] AttributeMutualExclusive
- [ ] ParentTypeConstraint
- [ ] ReferenceExistConstraint

### Phase 3: Advanced Constraints (2 weeks)
- [ ] AttributeValueLessEqualToAnother
- [ ] AttributeValueConditionToAnother
- [ ] AttributeAbsentConditionToNonValue
- [ ] AttributeAbsentConditionToValue
- [ ] AttributeRequiredConditionToValue
- [ ] AttributePairConstraint
- [ ] AttributeMinMaxConstraint

### Phase 4: Specialized Constraints (1 week)
- [ ] AttributeValuePatternConstraint (regex)
- [ ] UniqueAttributeValueConstraint
- [ ] RelationshipTypeConstraint
- [ ] RelationshipExistConstraint

### Phase 5: Integration & Testing (2 weeks)
- [ ] Add constraints to all element types
- [ ] Port semantic validation tests from SDK
- [ ] Performance profiling and optimization
- [ ] Documentation and examples

**Total Estimate: 8 weeks (1 engineer)**

## Test Strategy

### Test Sources
1. **Port SDK tests**: `DocumentFormat.OpenXml.Tests/Validation/`
2. **Malformed documents**: Intentionally broken Office files
3. **Real-world documents**: Documents from Office 2007-365

### Test Categories
- **Positive tests**: Valid documents pass validation
- **Negative tests**: Invalid documents fail with correct error
- **Edge cases**: Boundary values, absent attributes, empty elements
- **Performance tests**: Large documents validate in reasonable time

### Example Tests
```go
func TestAttributeMutualExclusive(t *testing.T) {
    para := wordprocessing.NewParagraph()
    para.SetAttribute("left", "100")
    para.SetAttribute("right", "200") // Invalid!
    
    validator := validation.NewValidator(validation.WithSemanticConstraints())
    errors := validator.ValidateElement(para)
    
    require.Len(t, errors, 1)
    assert.Contains(t, errors[0].Description, "mutually exclusive")
}

func TestReferenceExist(t *testing.T) {
    doc := wordprocessing.NewDocument()
    hyperlink := wordprocessing.NewHyperlink()
    hyperlink.Anchor = "NonExistentBookmark" // Invalid!
    
    errors := validation.Validate(doc)
    
    require.Len(t, errors, 1)
    assert.Contains(t, errors[0].Description, "bookmark")
}
```

## Risks & Mitigations

### Risk: Constraint definitions are verbose and error-prone
**Impact:** HIGH - Incorrect constraints produce false positives/negatives

**Mitigation:**
- Code generate constraints from SDK metadata where possible
- Comprehensive test suite comparing to SDK validation results
- Schema validation validates constraint definitions themselves

### Risk: Performance impact on large documents
**Impact:** MEDIUM - Semantic validation requires multiple document passes

**Mitigation:**
- Lazy validation (only when explicitly requested)
- Cache ID maps and relationship lookups
- Parallel validation of independent subtrees
- Provide "quick validation" mode (schema only)

### Risk: Constraint interactions are complex
**Impact:** MEDIUM - Multiple constraints may conflict or duplicate

**Mitigation:**
- Clear constraint priority and ordering
- Deduplicate errors from multiple constraints
- Document constraint interactions

### Risk: Office version differences in constraints
**Impact:** MEDIUM - Older Office versions may be more lenient

**Mitigation:**
- Version-aware constraints (Office2007, Office2010, etc.)
- Document which constraints apply to which versions
- Default to strictest validation, allow relaxation

## Success Criteria

1. ✅ All 16 semantic constraint types implemented
2. ✅ >90% of SDK semantic validation tests passing
3. ✅ Documents validated by goffice are accepted by Office applications
4. ✅ Clear, actionable error messages for all constraint violations
5. ✅ Validation performance <1s for typical documents (<1MB)
6. ✅ Comprehensive documentation with examples

## References

- **Open-XML-SDK Source**: `DocumentFormat.OpenXml.Framework/Validation/Semantic/`
- **ECMA-376 Part 1**: Section 17.18 (SimpleTypes), Annex L (Validation)
- **ISO/IEC 29500-1**: Semantic validation requirements
- **Office Compatibility Pack**: Validation behavior in Office 2007-2019

## Future Enhancements

Beyond initial implementation:
1. **Custom constraint plugins**: User-defined validation rules
2. **Streaming validation**: Validate during document construction
3. **Validation profiles**: Pre-configured constraint sets (strict, relaxed, government)
4. **Auto-repair**: Suggest or apply fixes for common errors
5. **IDE integration**: Real-time validation in document builders

---

**Estimated Effort:** 8 weeks (1 engineer)  
**Priority:** P0 - Critical for production use  
**Status:** Draft  
**Last Updated:** 2026-01-24
