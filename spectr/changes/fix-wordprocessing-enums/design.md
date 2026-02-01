# Proposal: Fix Wordprocessing Enums and Re-enable Dependent Code

## Summary

Fix the code generation for missing wordprocessing enum values and re-enable the dependent code that was commented out.

## Background

Several wordprocessing elements have TODOs indicating that code is disabled due to missing enum values:

**Reference TODOs:**
1. `wordprocessing/elements/enums_manual.go:135` - "Re-enable VerticalAlignmentRunValues when regenerated"
2. `wordprocessing/elements/run_properties.go:341,349,360` - "Re-enable VerticalAlignBaseline when regenerated" (3 locations)
3. `wordprocessing/elements/body.go:52` - "Re-enable PlaceholderValuesTbl when enum regenerated"
4. `wordprocessing/elements/header.go:51` - "Re-enable PlaceholderValuesTbl when enum regenerated"

These enums need to be properly generated from the Open XML schema, after which the commented code can be re-enabled.

## Motivation

Disabled features due to missing enums:
- Vertical alignment for runs (baseline alignment)
- Placeholder values for tables in body and header

These are important for document fidelity and feature completeness.

## Technical Design

### Analysis of Missing Enums

1. **VerticalAlignmentRunValues** - Used for run-level vertical alignment
   - Likely values: baseline, subscript, superscript
   
2. **PlaceholderValuesTbl** - Used for table placeholders
   - Likely values indicating table placeholder types

### Root Cause

The code generators in `cmd/gen-go-wordprocessing/` are not generating these enum values from the Open XML schema. This could be due to:
- Schema parsing issues
- Missing schema files
- Code generation logic gaps

### Proposed Solution

1. **Investigate Schema**
   - Examine Open XML schema for missing enum definitions
   - Verify schema files are up to date
   - Check if enums are in a different schema file

2. **Fix Code Generator**
   - Update `cmd/gen-go-wordprocessing/` to generate the missing enums
   - Ensure enum values match the Open XML specification
   - Add proper documentation comments

3. **Regenerate Enums**
   - Run the code generator
   - Verify new enum values are generated
   - Check for any other missing enums

4. **Re-enable Code**
   - Uncomment the disabled code in:
     - `enums_manual.go`
     - `run_properties.go`
     - `body.go`
     - `header.go`
   - Update references to use new enum values
   - Fix any API mismatches

### Code Changes

**Example re-enable in run_properties.go:**
```go
// Currently commented:
// TODO: Re-enable VerticalAlignBaseline when enum constant is regenerated
// if r.VerticalAlignBaseline != nil {
//     properties["VerticalAlignBaseline"] = r.VerticalAlignBaseline.Value
// }

// After fix:
if r.VerticalAlignBaseline != nil {
    properties["VerticalAlignBaseline"] = r.VerticalAlignBaseline.Value
}
```

## Requirements

### SHALL Requirements

#### Requirement: Generate Missing Enums
The code generator SHALL produce VerticalAlignmentRunValues enum with correct values from schema.

##### Scenario: VerticalAlignmentRunValues Enum
Given the Open XML schema defines vertical alignment values
When the code generator runs
Then it SHALL generate VerticalAlignmentRunValues enum with all schema-defined values

#### Requirement: PlaceholderValuesTbl Enum
The code generator SHALL produce PlaceholderValuesTbl enum with correct values.

##### Scenario: PlaceholderValuesTbl Enum Generated
Given the schema defines table placeholder values
When the code generator runs
Then it SHALL generate PlaceholderValuesTbl enum with all schema-defined values

#### Requirement: Re-enable VerticalAlignBaseline
The code in run_properties.go SHALL be re-enabled to use the regenerated enum.

##### Scenario: VerticalAlignBaseline Property
Given the VerticalAlignmentRunValues enum is available
When the code is re-enabled
Then run properties SHALL expose the VerticalAlignBaseline property

#### Requirement: Re-enable PlaceholderValuesTbl
The code in body.go and header.go SHALL be re-enabled.

##### Scenario: Body Placeholder Values
Given the PlaceholderValuesTbl enum is available
When the code is re-enabled
Then Body SHALL support table placeholder values

### SHOULD Requirements

#### Requirement: Schema Validation
The implementation SHOULD validate that generated enums match the schema.

#### Requirement: No Regressions
The fix SHOULD not break any existing enum generation.

## Implementation Plan

1. Examine Open XML schema files for missing enum definitions
2. Identify why code generator isn't producing these enums
3. Fix the code generator logic
4. Regenerate all wordprocessing elements
5. Re-enable commented code
6. Run tests to verify no regressions
7. Add tests for new enum values

## Related Changes

- `cmd/gen-go-wordprocessing/` - Code generator fixes
- `wordprocessing/elements/enums_manual.go` - Re-enable VerticalAlignmentRunValues
- `wordprocessing/elements/run_properties.go` - Re-enable VerticalAlignBaseline
- `wordprocessing/elements/body.go` - Re-enable PlaceholderValuesTbl
- `wordprocessing/elements/header.go` - Re-enable PlaceholderValuesTbl

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Schema doesn't define enums | High | Manual enum definition; report to Open-XML-SDK |
| Code generator changes break other enums | Medium | Comprehensive testing; incremental changes |
| Other dependencies on new enums | Low | Search codebase for related code |

## Acceptance Criteria

- [ ] VerticalAlignmentRunValues enum generated with correct values
- [ ] PlaceholderValuesTbl enum generated with correct values
- [ ] Commented code re-enabled in all 4 files
- [ ] All tests pass with no regressions
- [ ] New enum values tested
