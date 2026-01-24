# Tasks: Add Validation Semantic Constraints

**Epic**: Semantic Validation Infrastructure  
**Priority**: P0  
**Estimated Effort**: 8 weeks  
**Owner**: TBD

---

## Phase 1: Infrastructure (Week 1)

### Task 1.1: Define Core Interfaces
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Create `SemanticConstraint` interface in `openxml/validation/semantic.go`
- [ ] Define `ValidationContext` struct with ID/relationship tracking
- [ ] Create `SemanticValidationLevel` enum (Full, Partial, None, Custom)
- [ ] Add `SemanticError` type extending `ValidationError`
- [ ] Write interface documentation with examples

**Acceptance Criteria:**
- Interfaces compile and pass `go vet`
- Documentation includes usage examples
- Code review approved

---

### Task 1.2: Implement ValidationContext
**Estimated**: 3 days  
**Assignee**: TBD  
**Depends on**: Task 1.1

- [ ] Implement `NewValidationContext(doc CommonWrapper, opts ValidationOptions)` 
- [ ] Implement `GetPath(elem CommonWrapper) string` - returns XPath-style path
- [ ] Implement `ResolveID(id string) (CommonWrapper, bool)` - ID lookup
- [ ] Implement `GetRelationship(id string) (*Relationship, bool)` - relationship lookup
- [ ] Build ID map during context initialization (walk document tree)
- [ ] Build relationship map from package
- [ ] Add context state tracking (current element path)
- [ ] Write unit tests for context operations

**Acceptance Criteria:**
- All context methods work correctly
- ID resolution handles missing IDs gracefully
- Path generation matches expected format
- Unit tests achieve >90% coverage

---

### Task 1.3: Extend Validator
**Estimated**: 2 days  
**Assignee**: TBD  
**Depends on**: Task 1.2

- [ ] Add `WithSemanticConstraints()` option to `NewValidator()`
- [ ] Add `WithSemanticLevel(level SemanticValidationLevel)` option
- [ ] Modify `Validate()` to run semantic constraints after schema validation
- [ ] Collect semantic errors alongside schema errors
- [ ] Add `ValidateElement()` for single element validation
- [ ] Update validation result structure to separate semantic/schema errors
- [ ] Write integration tests

**Acceptance Criteria:**
- Validator runs semantic constraints when enabled
- Errors are properly categorized
- Performance impact is measurable (<20% overhead)

---

## Phase 2: Core Constraints (Weeks 2-3)

### Task 2.1: AttributeValueSetConstraint
**Estimated**: 1 day  
**Assignee**: TBD  
**Depends on**: Task 1.3

- [ ] Implement `AttributeValueSetConstraint` struct
- [ ] Implement `Validate()` method
- [ ] Handle absent attribute (not an error)
- [ ] Handle case-sensitive vs case-insensitive comparison
- [ ] Write unit tests with valid/invalid cases
- [ ] Add to constraint catalog

**Test Cases:**
- Attribute value in allowed set → pass
- Attribute value not in allowed set → fail with error
- Attribute absent → pass
- Empty allowed set → fail on any value

---

### Task 2.2: AttributeValueRangeConstraint
**Estimated**: 1 day  
**Assignee**: TBD  
**Depends on**: Task 1.3

- [ ] Implement `AttributeValueRangeConstraint` struct (min, max)
- [ ] Support integer and float ranges
- [ ] Handle inclusive vs exclusive bounds
- [ ] Handle absent attribute
- [ ] Handle non-numeric values
- [ ] Write unit tests
- [ ] Add to constraint catalog

**Test Cases:**
- Value in range → pass
- Value below minimum → fail
- Value above maximum → fail
- Value at boundary → pass (if inclusive)
- Non-numeric value → fail with type error

---

### Task 2.3: AttributeMutualExclusive
**Estimated**: 2 days  
**Assignee**: TBD  
**Depends on**: Task 1.3

- [ ] Implement `AttributeMutualExclusive` constraint
- [ ] Support 2+ mutually exclusive attributes
- [ ] Check all combinations
- [ ] Provide clear error message listing conflicting attributes
- [ ] Write unit tests
- [ ] Add to constraint catalog

**Test Cases:**
- Only one attribute present → pass
- No attributes present → pass
- Two attributes present → fail
- Three attributes present → fail

---

### Task 2.4: ParentTypeConstraint
**Estimated**: 2 days  
**Assignee**: TBD  
**Depends on**: Task 1.3

- [ ] Implement `ParentTypeConstraint` struct
- [ ] Support multiple allowed parent types
- [ ] Handle document root (no parent)
- [ ] Check actual parent type against allowed types
- [ ] Write unit tests
- [ ] Add to constraint catalog

**Test Cases:**
- Element under allowed parent → pass
- Element under disallowed parent → fail
- Element at document root → handle gracefully

---

### Task 2.5: ReferenceExistConstraint
**Estimated**: 3 days  
**Assignee**: TBD  
**Depends on**: Task 1.2

- [ ] Implement `ReferenceExistConstraint` struct
- [ ] Support bookmark references
- [ ] Support style references
- [ ] Support numbering references
- [ ] Support custom XML references
- [ ] Use ValidationContext.ResolveID()
- [ ] Write comprehensive tests
- [ ] Add to constraint catalog

**Test Cases:**
- Reference to existing ID → pass
- Reference to missing ID → fail with descriptive error
- Multiple references → all checked
- Circular references → handle gracefully

---

## Phase 3: Advanced Constraints (Weeks 4-5)

### Task 3.1: AttributeValueLessEqualToAnother
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement comparison constraint (A ≤ B)
- [ ] Support integer, float, date comparisons
- [ ] Handle absent attributes
- [ ] Handle type mismatches
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.2: AttributeValueConditionToAnother
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement conditional constraint (if A=X then B=Y)
- [ ] Support multiple condition types (equal, not equal, greater, less)
- [ ] Handle absent attributes
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.3: AttributeAbsentConditionToNonValue
**Estimated**: 1 day  
**Assignee**: TBD

- [ ] Implement constraint: if A absent, B cannot equal X
- [ ] Handle all absence scenarios
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.4: AttributeAbsentConditionToValue
**Estimated**: 1 day  
**Assignee**: TBD

- [ ] Implement constraint: if A absent, B must equal X
- [ ] Handle all absence scenarios
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.5: AttributeRequiredConditionToValue
**Estimated**: 1 day  
**Assignee**: TBD

- [ ] Implement constraint: if A=X, B is required
- [ ] Handle absence checking
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.6: AttributePairConstraint
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement constraint: attributes must appear together
- [ ] Support N-way constraints (3+ attributes)
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 3.7: AttributeMinMaxConstraint
**Estimated**: 1 day  
**Assignee**: TBD

- [ ] Implement constraint: min attribute ≤ max attribute
- [ ] Handle absent attributes
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

## Phase 4: Specialized Constraints (Week 6)

### Task 4.1: AttributeValuePatternConstraint
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement regex pattern constraint
- [ ] Use Go `regexp` package
- [ ] Handle invalid regex patterns
- [ ] Provide helpful error messages with examples
- [ ] Write unit tests with various patterns
- [ ] Add to constraint catalog

**Test Cases:**
- Value matches pattern → pass
- Value doesn't match pattern → fail with example
- Invalid regex → error during constraint setup

---

### Task 4.2: UniqueAttributeValueConstraint
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement uniqueness constraint across elements
- [ ] Build value index during validation
- [ ] Support scoped uniqueness (within parent)
- [ ] Support global uniqueness (entire document)
- [ ] Write unit tests
- [ ] Add to constraint catalog

**Test Cases:**
- All values unique → pass
- Duplicate values → fail with locations
- Empty values → handle appropriately

---

### Task 4.3: RelationshipTypeConstraint
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Implement relationship type checking
- [ ] Use ValidationContext relationship map
- [ ] Check relationship type against allowed types
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

### Task 4.4: RelationshipExistConstraint
**Estimated**: 1 day  
**Assignee**: TBD

- [ ] Implement required relationship constraint
- [ ] Check relationship existence in package
- [ ] Write unit tests
- [ ] Add to constraint catalog

---

## Phase 5: Integration & Testing (Weeks 7-8)

### Task 5.1: Add Constraints to Elements
**Estimated**: 5 days  
**Assignee**: TBD  
**Depends on**: All constraint implementations

- [ ] Add `SemanticConstraints()` method to all element types
- [ ] Define constraints for wordprocessing elements
- [ ] Define constraints for spreadsheet elements
- [ ] Define constraints for presentation elements
- [ ] Define constraints for drawing elements
- [ ] Code generation tool for constraint definitions
- [ ] Review all constraint definitions

**Affected Elements**: ~300+ element types

---

### Task 5.2: Port SDK Validation Tests
**Estimated**: 4 days  
**Assignee**: TBD

- [ ] Identify relevant tests in `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests/Validation/`
- [ ] Port semantic validation tests to Go
- [ ] Adapt tests to goffice API
- [ ] Ensure all tests pass
- [ ] Add additional edge case tests

**Target**: ~200 ported tests

---

### Task 5.3: Performance Optimization
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Profile validation performance
- [ ] Optimize ID map building
- [ ] Optimize constraint checking order
- [ ] Cache constraint results where possible
- [ ] Benchmark before/after optimization
- [ ] Document performance characteristics

**Target**: <1s for typical documents

---

### Task 5.4: Documentation
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Write constraint catalog documentation
- [ ] Write validation guide for users
- [ ] Create example code for each constraint type
- [ ] Document validation options
- [ ] Document error messages
- [ ] Update package-level documentation

---

### Task 5.5: Integration Testing
**Estimated**: 2 days  
**Assignee**: TBD

- [ ] Test with real-world documents
- [ ] Test with intentionally malformed documents
- [ ] Test validation of Office 2007-365 documents
- [ ] Compare results with Open-XML-SDK validation
- [ ] Fix any discrepancies
- [ ] Document known differences

---

## Code Review Checkpoints

### Checkpoint 1 (End of Phase 1)
- [ ] Interface design review
- [ ] API consistency check
- [ ] Documentation review
- [ ] Performance baseline established

### Checkpoint 2 (End of Phase 2)
- [ ] Core constraints implementation review
- [ ] Test coverage review (target: >90%)
- [ ] Error message quality review

### Checkpoint 3 (End of Phase 3)
- [ ] Advanced constraints review
- [ ] Edge case coverage review
- [ ] Integration with existing validation

### Checkpoint 4 (End of Phase 5)
- [ ] Final code review
- [ ] Documentation review
- [ ] Performance review
- [ ] Release readiness

---

## Definition of Done

- [ ] All 16 constraint types implemented and tested
- [ ] Unit test coverage >90% for semantic validation code
- [ ] All ported SDK tests passing
- [ ] Performance benchmark established and met
- [ ] Documentation complete with examples
- [ ] Code review approved
- [ ] Integration testing complete
- [ ] No P0/P1 bugs outstanding

---

## Dependencies

**Upstream Dependencies:**
- openxml/validation framework (exists)
- Element metadata system (exists)
- Relationship management (exists)

**Downstream Dependents:**
- Comprehensive test suite
- Strict namespace support (uses semantic validation)

---

## Risks & Mitigation

**Risk 1**: Constraint definitions are error-prone  
**Mitigation**: Code generation from SDK metadata + comprehensive tests

**Risk 2**: Performance impact too high  
**Mitigation**: Profiling, optimization, make semantic validation opt-in

**Risk 3**: Complex constraint interactions  
**Mitigation**: Clear ordering, deduplicate errors, document interactions

---

**Last Updated**: 2026-01-24  
**Status**: Ready for implementation
