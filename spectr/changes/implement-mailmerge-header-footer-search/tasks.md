# Tasks: Implement Mailmerge Header and Footer Field Search

## Implementation

- [ ] 1.1 Update `findMergeFields()` in field_finder.go to iterate document headers
- [ ] 1.2 Update `findMergeFields()` in field_finder.go to iterate document footers
- [ ] 1.3 Ensure proper handling of documents without headers/footers
- [ ] 1.4 Test with simple fields in headers
- [ ] 1.5 Test with complex fields (fldChar) in headers
- [ ] 1.6 Test with simple fields in footers
- [ ] 1.7 Test with complex fields (fldChar) in footers

## Unit Tests

- [ ] 2.1 Create test: Field detection in default header
- [ ] 2.2 Create test: Field detection in first page header
- [ ] 2.3 Create test: Field detection in default footer
- [ ] 2.4 Create test: Field detection in first page footer
- [ ] 2.5 Create test: Same field name in body and header
- [ ] 2.6 Create test: Same field name in body and footer
- [ ] 2.7 Create test: Document without headers
- [ ] 2.8 Create test: Document without footers
- [ ] 2.9 Create test: GREETINGLINE in header
- [ ] 2.10 Create test: Multiple field types across all locations

## Integration Tests

- [ ] 3.1 Create end-to-end test: Full mailmerge with header fields
- [ ] 3.2 Create end-to-end test: Full mailmerge with footer fields
- [ ] 3.3 Create end-to-end test: Multiple records with header/footer fields

## Test Fixtures

- [ ] 4.1 Create test document: Simple MERGEFIELD in header
- [ ] 4.2 Create test document: Complex field markers in header
- [ ] 4.3 Create test document: Simple MERGEFIELD in footer
- [ ] 4.4 Create test document: Complex field markers in footer
- [ ] 4.5 Create test document: Fields in body, header, and footer

## Documentation

- [ ] 5.1 Update AGENTS.md if needed
- [ ] 5.2 Update README.md with header/footer field examples
- [ ] 5.3 Add code comments explaining header/footer search

## Verification

- [ ] 6.1 Run all mailmerge tests - ensure no regressions
- [ ] 6.2 Verify test coverage >90% for changed code
- [ ] 6.3 Manual test with real Word document containing header/footer fields
