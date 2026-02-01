# Tasks: Fix Wordprocessing Enums and Re-enable Dependent Code

## Schema Analysis

- [ ] 1.1 Examine Open XML schema for VerticalAlignmentRunValues definition
- [ ] 1.2 Examine Open XML schema for PlaceholderValuesTbl definition
- [ ] 1.3 Check which schema file contains these enum definitions
- [ ] 1.4 Document the expected enum values
- [ ] 1.5 Verify schema files are current

## Code Generator Investigation

- [ ] 2.1 Review cmd/gen-go-wordprocessing/ enum generation logic
- [ ] 2.2 Identify why VerticalAlignmentRunValues is not generated
- [ ] 2.3 Identify why PlaceholderValuesTbl is not generated
- [ ] 2.4 Check for similar missing enums
- [ ] 2.5 Document root cause

## Code Generator Fix

- [ ] 3.1 Update enum generation logic to include VerticalAlignmentRunValues
- [ ] 3.2 Update enum generation logic to include PlaceholderValuesTbl
- [ ] 3.3 Add proper documentation comment generation
- [ ] 3.4 Ensure enum values match schema exactly
- [ ] 3.5 Test generator with sample schema

## Regeneration

- [ ] 4.1 Run code generator for wordprocessing elements
- [ ] 4.2 Verify VerticalAlignmentRunValues is generated
- [ ] 4.3 Verify PlaceholderValuesTbl is generated
- [ ] 4.4 Check for any other changes in generated files
- [ ] 4.5 Review all generated enum values for correctness

## Re-enable Commented Code

- [ ] 5.1 Re-enable code in enums_manual.go (VerticalAlignmentRunValues)
- [ ] 5.2 Re-enable code in run_properties.go line 341
- [ ] 5.3 Re-enable code in run_properties.go line 349
- [ ] 5.4 Re-enable code in run_properties.go line 360
- [ ] 5.5 Re-enable code in body.go (PlaceholderValuesTbl)
- [ ] 5.6 Re-enable code in header.go (PlaceholderValuesTbl)
- [ ] 5.7 Remove TODO comments from re-enabled code
- [ ] 5.8 Verify code compiles correctly

## Fix Any API Mismatches

- [ ] 6.1 Check if method signatures need updates
- [ ] 6.2 Check if property names need updates
- [ ] 6.3 Update any broken references
- [ ] 6.4 Ensure consistency across files

## Unit Tests

- [ ] 7.1 Create test: VerticalAlignmentRunValues enum values
- [ ] 7.2 Create test: PlaceholderValuesTbl enum values
- [ ] 7.3 Create test: VerticalAlignBaseline property access
- [ ] 7.4 Create test: PlaceholderValuesTbl in body
- [ ] 7.5 Create test: PlaceholderValuesTbl in header
- [ ] 7.6 Verify existing enum tests still pass

## Regression Tests

- [ ] 8.1 Run all wordprocessing element tests
- [ ] 8.2 Verify no test failures after regeneration
- [ ] 8.3 Check for any API changes that break callers
- [ ] 8.4 Run integration tests

## Documentation

- [ ] 9.1 Update wordprocessing/AGENTS.md with enum info
- [ ] 9.2 Document the regeneration process
- [ ] 9.3 Add code comments explaining enum usage
- [ ] 9.4 Update CHANGELOG with fixed enums

## Verification

- [ ] 10.1 All 4 TODO comments removed/resolved
- [ ] 10.2 Both enums are now generated automatically
- [ ] 10.3 Previously commented code is active
- [ ] 10.4 All tests pass
- [ ] 10.5 No regressions in existing functionality
