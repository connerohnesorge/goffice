# Tasks: Implement PDF Spreadsheet Layout with Actual Column Widths and Row Heights

## Implementation

- [ ] 1.1 Add column width lookup method to SpreadsheetRenderer
- [ ] 1.2 Add row height lookup method to SpreadsheetRenderer
- [ ] 1.3 Implement caching for column widths
- [ ] 1.4 Implement caching for row heights
- [ ] 1.5 Update calculateCellBounds() to use actual column widths
- [ ] 1.6 Update calculateCellBounds() to use actual row heights
- [ ] 1.7 Add fallback to default values when dimensions not set
- [ ] 1.8 Add validation for reasonable dimension bounds
- [ ] 1.9 Ensure consistent units (points) throughout

## Worksheet API (if needed)

- [ ] 2.1 Verify GetColumnWidth() method exists on worksheet
- [ ] 2.2 Verify GetRowHeight() method exists on worksheet
- [ ] 2.3 Add GetColumnWidth() if missing
- [ ] 2.4 Add GetRowHeight() if missing
- [ ] 2.5 Add GetDefaultColumnWidth() for fallback
- [ ] 2.6 Add GetDefaultRowHeight() for fallback

## Unit Tests

- [ ] 3.1 Create test: Column width retrieval from worksheet
- [ ] 3.2 Create test: Row height retrieval from worksheet
- [ ] 3.3 Create test: Column width caching
- [ ] 3.4 Create test: Row height caching
- [ ] 3.5 Create test: Fallback to default column width
- [ ] 3.6 Create test: Fallback to default row height
- [ ] 3.7 Create test: Zero/negative dimension handling
- [ ] 3.8 Create test: Unit conversion consistency

## Integration Tests

- [ ] 4.1 Create test: PDF rendering with explicit column widths
- [ ] 4.2 Create test: PDF rendering with explicit row heights
- [ ] 4.3 Create test: PDF rendering with auto-fitted columns
- [ ] 4.4 Create test: PDF rendering with auto-fitted rows
- [ ] 4.5 Create test: PDF rendering with mixed explicit and auto dimensions

## Visual/Regression Tests

- [ ] 5.1 Create visual test: Simple spreadsheet with varying column widths
- [ ] 5.2 Create visual test: Simple spreadsheet with varying row heights
- [ ] 5.3 Create visual test: Complex spreadsheet with merged cells
- [ ] 5.4 Compare output with reference images

## Test Fixtures

- [ ] 6.1 Create test spreadsheet: Explicit column widths (3 columns)
- [ ] 6.2 Create test spreadsheet: Explicit row heights (3 rows)
- [ ] 6.3 Create test spreadsheet: Auto-fitted columns
- [ ] 6.4 Create test spreadsheet: Auto-fitted rows with wrapped text
- [ ] 6.5 Create test spreadsheet: Mixed dimensions

## Documentation

- [ ] 7.1 Update pdf/AGENTS.md with layout calculation details
- [ ] 7.2 Add code comments explaining dimension lookup
- [ ] 7.3 Document caching strategy
- [ ] 7.4 Document fallback behavior

## Verification

- [ ] 8.1 Run all pdf spreadsheet tests - ensure no regressions
- [ ] 8.2 Verify test coverage >90% for changed code
- [ ] 8.3 Visual comparison with Excel PDF export
- [ ] 8.4 Performance benchmark shows improvement or no degradation
