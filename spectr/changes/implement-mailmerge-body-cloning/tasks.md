# Tasks: Implement Mailmerge Body Cloning for Multiple Records

## Implementation

- [ ] 1.1 Create body cloning utility function
- [ ] 1.2 Implement deep copy via XML serialization approach
- [ ] 1.3 Update Execute() to iterate through all data source records
- [ ] 1.4 Implement page break insertion between records
- [ ] 1.5 Ensure field replacement works correctly per record
- [ ] 1.6 Handle SKIPIF fields correctly in multi-record context
- [ ] 1.7 Add error handling for cloning failures

## Unit Tests

- [ ] 2.1 Create test: Clone simple paragraph body
- [ ] 2.2 Create test: Clone body with tables
- [ ] 2.3 Create test: Clone body with images
- [ ] 2.4 Create test: Clone body with complex formatting
- [ ] 2.5 Create test: Field replacement in cloned body
- [ ] 2.6 Create test: Page break insertion
- [ ] 2.7 Create test: Multiple records processing (3 records)
- [ ] 2.8 Create test: Many records processing (10+ records)

## Integration Tests

- [ ] 3.1 Create end-to-end test: Full mailmerge with 3 CSV records
- [ ] 3.2 Create end-to-end test: Full mailmerge with 5 JSON records
- [ ] 3.3 Create end-to-end test: Complex template with multiple records
- [ ] 3.4 Create end-to-end test: Verify formatting preserved across clones

## Performance Tests

- [ ] 4.1 Performance test: 10 records, simple template
- [ ] 4.2 Performance test: 50 records, simple template
- [ ] 4.3 Performance test: 100 records, simple template
- [ ] 4.4 Memory profiling: Ensure no leaks with multiple records

## Test Fixtures

- [ ] 5.1 Create test CSV: 3 records with standard fields
- [ ] 5.2 Create test CSV: 10 records with various data types
- [ ] 5.3 Create test JSON: Array of 5 objects
- [ ] 5.4 Create test template: Simple body with fields
- [ ] 5.5 Create test template: Body with tables
- [ ] 5.6 Create test template: Body with complex formatting

## Documentation

- [ ] 6.1 Update AGENTS.md with body cloning details
- [ ] 6.2 Update README.md with multi-record examples
- [ ] 6.3 Add code comments explaining cloning strategy
- [ ] 6.4 Document performance characteristics

## Verification

- [ ] 7.1 Run all mailmerge tests - ensure no regressions
- [ ] 7.2 Verify ExecuteToDocuments() still works correctly
- [ ] 7.3 Verify test coverage >90% for changed code
- [ ] 7.4 Manual test with real Word document and CSV
- [ ] 7.5 Performance benchmark meets criteria
