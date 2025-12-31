# Implementation Tasks

## Phase 1: Cell Reference Infrastructure (1.5 weeks)

### 1.1 Cell Reference Parser
- [ ] **Create `spreadsheet/cell_reference.go`**
  - Define `CellReference` struct (sheet, column, row, absolute flags, R1C1 fields)
  - Define `RangeReference` struct (start, end CellReference)
  - Implement `ParseCellRef(ref string) (*CellReference, error)` for A1 notation
  - Implement `ParseRangeRef(ref string) (*RangeReference, error)`
  - Implement `(cr *CellReference) String() string` to convert back to A1 notation
  - Implement `(rr *RangeReference) String() string`
  - Handle simple refs: `A1`, `XFD1048576`
  - Handle absolute refs: `$A$1`, `$A1`, `A$1`
  - Handle cross-sheet refs: `Sheet2!A1`, `'Sheet Name'!A1`
  - Handle quoted sheet names with special characters
  - Validation: Compiles, basic manual test

- [ ] **Add R1C1 notation support** to `cell_reference.go`
  - Implement R1C1 detection logic (starts with "R")
  - Parse absolute R1C1: `R5C10`
  - Parse relative R1C1: `R[-1]C[2]`, `RC[-1]`
  - Store in `CellReference` with `IsR1C1=true` and relative offsets
  - Validation: Compiles, basic manual test

- [ ] **Create column name utilities**
  - Implement `ColumnName(col uint32) string` (1→"A", 27→"AA")
  - Implement `ColumnIndex(name string) (uint32, error)` ("AA"→27)
  - Handle extended columns (up to XFD = 16384)
  - Validation: Compiles, basic manual test

- [ ] **Write cell reference parser unit tests** in `cell_reference_test.go`
  - Test A1 notation parsing (50+ cases): `A1`, `Z99`, `AA100`, `XFD1048576`
  - Test absolute parsing (30+ cases): `$A$1`, `$A1`, `A$1`
  - Test cross-sheet parsing (40+ cases): `Sheet2!A1`, `'Sheet Name'!A1`, `'Sheet (2024)'!A1`
  - Test range parsing (30+ cases): `A1:D10`, `$A$1:$D$10`, `Sheet1!A1:D10`
  - Test R1C1 parsing (30+ cases): `R1C1`, `R[-1]C[2]`, `RC[-1]`
  - Test invalid references (20+ cases): `@@@`, `A`, `1`, `Sheet!`, `A0`, `XFE1`
  - Test String() conversion (20+ cases): roundtrip parsing
  - Validation: All tests pass, >95% coverage

- [ ] **Write column utilities unit tests** in `cell_reference_test.go`
  - Test ColumnName (20+ cases): 1→"A", 26→"Z", 27→"AA", 16384→"XFD"
  - Test ColumnIndex (20+ cases): "A"→1, "Z"→26, "AA"→27, "XFD"→16384
  - Test edge cases: empty string, invalid chars, overflow
  - Validation: All tests pass, >95% coverage

## Phase 2: Formula Rewriting (1.5 weeks)

### 2.1 Formula Tokenizer
- [ ] **Create `spreadsheet/formula_rewriter.go`**
  - Define `TokenType` enum (Operator, Function, Reference, Literal, Paren, Comma)
  - Define `Token` struct (Type, Value)
  - Implement `tokenizeFormula(formula string) []Token`
  - Handle operators: `+`, `-`, `*`, `/`, `=`, `<>`, `<=`, `>=`, `&`
  - Handle functions: detect name followed by `(`
  - Handle cell references: use `ParseCellRef` and `ParseRangeRef`
  - Handle string literals: quoted strings (don't parse contents)
  - Handle parentheses and commas
  - Handle numeric literals
  - Validation: Compiles, basic manual test

- [ ] **Write tokenizer unit tests** in `formula_rewriter_test.go`
  - Test simple formulas (30+ cases): `=A1+B1`, `=SUM(A1:A10)`, `=A1*2`
  - Test complex formulas (30+ cases): `=IF(A1>0,SUM(B1:B10),0)`, nested functions
  - Test string literals (20+ cases): `="Value: "&A1`, `="Don't parse A1"`
  - Test operators (20+ cases): all operator types, precedence
  - Test cross-sheet refs (20+ cases): `=Sheet2!A1+Sheet3!B2`
  - Test edge cases: leading `=`, no `=`, empty formula, formulas with spaces
  - Validation: All tests pass, >90% coverage

### 2.2 Reference Shift Logic
- [ ] **Implement `CellReferenceUpdater` in `cell_reference_updater.go`**
  - Define `ShiftType` enum (InsertRows, DeleteRows, InsertColumns, DeleteColumns)
  - Define `ShiftOperation` struct (Type, Index, Count, Worksheet)
  - Create `CellReferenceUpdater` struct with operation
  - Implement `ShouldUpdate(ref *CellReference) bool`
    - Check worksheet match (empty = current sheet)
    - Check absolute vs relative addressing
    - Check if reference is in affected range
  - Implement `UpdateReference(ref *CellReference) *CellReference`
    - Apply shift to row or column
    - Handle deleted range (return invalid reference marker)
    - Handle absolute references (don't shift)
  - Implement `UpdateRange(rangeRef *RangeReference) *RangeReference`
    - Update both start and end
  - Validation: Compiles, basic manual test

- [ ] **Write reference updater unit tests** in `cell_reference_updater_test.go`
  - Test InsertRows (40+ cases): before insert, at insert, after insert, absolute vs relative
  - Test DeleteRows (40+ cases): before delete, in deleted range, after delete
  - Test InsertColumns (40+ cases): similar to InsertRows
  - Test DeleteColumns (40+ cases): similar to DeleteRows
  - Test cross-sheet scenarios (30+ cases): different sheet not affected
  - Test absolute references (30+ cases): $A$1 not shifted
  - Test mixed references (30+ cases): $A1 and A$1 behavior
  - Test edge cases: insert at row 1, delete at max row, whole column refs
  - Validation: All tests pass, >95% coverage

### 2.3 Formula Reconstruction
- [ ] **Implement formula rewriter** in `formula_rewriter.go`
  - Create `FormulaRewriter` struct
  - Implement `Rewrite(formula string, operation ShiftOperation) string`
    - Tokenize formula
    - For each token of type Reference:
      - Parse reference
      - Update using CellReferenceUpdater
      - Replace token value with updated reference
    - Reconstruct formula from tokens
  - Implement `reconstructFormula(tokens []Token) string`
    - Join tokens with appropriate spacing
    - Ensure leading `=` is present
  - Validation: Compiles, basic manual test

- [ ] **Write formula rewriter unit tests** in `formula_rewriter_test.go`
  - Test simple formula updates (30+ cases): `=A1+B1` after insert/delete
  - Test range formula updates (30+ cases): `=SUM(A1:A10)` range expansion
  - Test absolute reference preservation (20+ cases): `=$A$1+B1`
  - Test cross-sheet formulas (20+ cases): `=Sheet2!A1` when Sheet2 modified
  - Test complex formulas (30+ cases): nested functions, multiple references
  - Test string literal preservation (20+ cases): `="A1"` not updated
  - Test edge cases: references in deleted range → #REF!
  - Validation: All tests pass, >90% coverage

## Phase 3: Insert/Delete Operations (1.5 weeks)

### 3.1 Physical Row/Column Manipulation
- [ ] **Create `spreadsheet/sheet_operations.go`**
  - Define `SheetOperations` struct wrapping `*Sheet`
  - Validation: Compiles

- [ ] **Implement InsertRows in `sheet_operations.go`**
  - Implement `InsertRows(index, count uint32) error`
  - Validate parameters (index > 0, count > 0, within bounds)
  - Get SheetData element
  - Shift existing rows down:
    - Iterate rows in reverse order
    - Update row R attribute (row index)
    - Update cell R attributes (cell references)
  - Create new empty rows at insertion point
  - Insert new rows into SheetData at correct position
  - Update sheet dimensions if present
  - Validation: Compiles, basic manual test

- [ ] **Implement DeleteRows in `sheet_operations.go`**
  - Implement `DeleteRows(index, count uint32) error`
  - Validate parameters
  - Get SheetData element
  - Remove rows in deletion range
  - Shift remaining rows up:
    - Update row R attribute
    - Update cell R attributes
  - Update sheet dimensions
  - Validation: Compiles, basic manual test

- [ ] **Implement InsertColumns in `sheet_operations.go`**
  - Implement `InsertColumns(index, count uint32) error`
  - Similar to InsertRows but for columns
  - Shift column indices in all cells
  - Update column elements if present
  - Update sheet dimensions
  - Validation: Compiles, basic manual test

- [ ] **Implement DeleteColumns in `sheet_operations.go`**
  - Implement `DeleteColumns(index, count uint32) error`
  - Similar to DeleteRows but for columns
  - Delete cells in affected columns
  - Shift remaining columns left
  - Update column elements
  - Update sheet dimensions
  - Validation: Compiles, basic manual test

- [ ] **Write physical operation unit tests** in `sheet_operations_test.go`
  - Test InsertRows: verify row indices shifted, new rows created
  - Test DeleteRows: verify rows removed, indices shifted
  - Test InsertColumns: verify column indices shifted
  - Test DeleteColumns: verify columns removed
  - Test edge cases: insert at row 1, delete all rows, bounds validation
  - Validation: All tests pass, >85% coverage

### 3.2 Reference Update Integration
- [ ] **Implement formula updates in `sheet_operations.go`**
  - Add `updateFormulas(operation ShiftOperation)` helper
  - Iterate all worksheets in workbook
  - For each cell with formula:
    - Get formula text
    - Use FormulaRewriter to update
    - Set updated formula
  - Validation: Compiles, basic manual test

- [ ] **Implement named range updates in `sheet_operations.go`**
  - Add `updateNamedRanges(operation ShiftOperation)` helper
  - Get workbook DefinedNames element
  - For each DefinedName:
    - Get reference/formula
    - Determine if range reference or formula
    - Use CellReferenceUpdater or FormulaRewriter as appropriate
    - Set updated reference
  - Validation: Compiles, basic manual test

- [ ] **Implement merged cells updates in `sheet_operations.go`**
  - Add `updateMergeCells(operation ShiftOperation)` helper
  - Get worksheet MergeCells element
  - For each MergeCell:
    - Parse ref attribute (range reference)
    - Update using CellReferenceUpdater
    - Set updated ref
  - Validation: Compiles, basic manual test

- [ ] **Implement conditional formatting updates in `sheet_operations.go`**
  - Add `updateConditionalFormatting(operation ShiftOperation)` helper
  - Get worksheet ConditionalFormatting elements
  - For each ConditionalFormatting:
    - Update sqref attribute (range reference)
    - Update formula references in rules if present
  - Validation: Compiles, basic manual test

- [ ] **Implement data validation updates in `sheet_operations.go`**
  - Add `updateDataValidation(operation ShiftOperation)` helper
  - Get worksheet DataValidations element
  - For each DataValidation:
    - Update sqref attribute (range list)
    - Update formula1, formula2 if present
  - Validation: Compiles, basic manual test

- [ ] **Implement table range updates in `sheet_operations.go`**
  - Add `updateTableRanges(operation ShiftOperation)` helper
  - Get worksheet TableParts
  - For each TablePart:
    - Access Table element
    - Update ref attribute (range reference)
    - Update autoFilter ref if present
  - Validation: Compiles, basic manual test

- [ ] **Integrate reference updates into operations**
  - Modify `InsertRows` to call all update helpers after physical shift
  - Modify `DeleteRows` to call all update helpers
  - Modify `InsertColumns` to call all update helpers
  - Modify `DeleteColumns` to call all update helpers
  - Validation: Compiles, basic manual test

- [ ] **Write reference update integration tests** in `sheet_operations_test.go`
  - Test formula updates after InsertRows (20+ cases)
  - Test formula updates after DeleteRows (20+ cases)
  - Test named range updates (15+ cases)
  - Test merged cells updates (15+ cases)
  - Test conditional formatting updates (10+ cases)
  - Test data validation updates (10+ cases)
  - Test table range updates (10+ cases)
  - Test cross-sheet reference updates (15+ cases)
  - Test shared formula updates (20+ cases):
    - Master cell formula text updates
    - Master cell ref attribute updates
    - Shared index preservation (si attribute unchanged)
    - Child cell handling (no formula updates)
    - Shared formula with affected references
    - Shared formula when master cell is deleted
  - Test array formula updates (20+ cases):
    - CSE array formula ref attribute updates
    - CSE array anchor cell formula text updates
    - CSE array member cell ref updates
    - Dynamic array formula updates (no ref attribute)
    - Array formula when anchor deleted
    - Array formula with partially deleted range
  - Test data table formula updates (10+ cases):
    - R1 and R2 attribute updates
    - Data table range updates
  - Validation: All tests pass, >85% coverage

### 3.3 Sheet API Integration
- [ ] **Add operation methods to `spreadsheet/sheet.go`**
  - Implement `(s *Sheet) InsertRows(index, count uint32) error`
    - Create SheetOperations instance
    - Delegate to SheetOperations.InsertRows
    - Mark document dirty
  - Implement `(s *Sheet) DeleteRows(index, count uint32) error`
  - Implement `(s *Sheet) InsertColumns(index, count uint32) error`
  - Implement `(s *Sheet) DeleteColumns(index, count uint32) error`
  - Add godoc comments for all methods
  - Validation: Compiles, basic manual test

- [ ] **Write Sheet API tests** in `sheet_test.go`
  - Test InsertRows API: create sheet, insert rows, verify
  - Test DeleteRows API
  - Test InsertColumns API
  - Test DeleteColumns API
  - Test error cases: invalid parameters, nil sheet
  - Validation: All tests pass, >85% coverage

## Phase 4: Grouping Operations (0.5 weeks)

### 4.1 Row/Column Grouping
- [ ] **Implement row grouping in `sheet_operations.go`**
  - Implement `GroupRows(start, end uint32) error`
  - Validate parameters (start <= end, within bounds)
  - For each row in range:
    - Get or create Row element
    - Get current outlineLevel (default 0)
    - Set outlineLevel to current + 1 (max 7)
  - Validation: Compiles, basic manual test

- [ ] **Implement row ungrouping in `sheet_operations.go`**
  - Implement `UngroupRows(start, end uint32) error`
  - For each row in range:
    - Get Row element
    - Get current outlineLevel
    - Decrement outlineLevel (min 0)
    - If outlineLevel = 0, remove attribute
  - Validation: Compiles, basic manual test

- [ ] **Implement column grouping in `sheet_operations.go`**
  - Implement `GroupColumns(start, end uint32) error`
  - Similar to GroupRows but for Column elements
  - Validation: Compiles, basic manual test

- [ ] **Implement column ungrouping in `sheet_operations.go`**
  - Implement `UngroupColumns(start, end uint32) error`
  - Similar to UngroupRows
  - Validation: Compiles, basic manual test

- [ ] **Add grouping methods to `spreadsheet/sheet.go`**
  - Implement `(s *Sheet) GroupRows(start, end uint32) error`
  - Implement `(s *Sheet) UngroupRows(start, end uint32) error`
  - Implement `(s *Sheet) GroupColumns(start, end uint32) error`
  - Implement `(s *Sheet) UngroupColumns(start, end uint32) error`
  - Add godoc comments
  - Validation: Compiles, basic manual test

- [ ] **Write grouping tests** in `sheet_operations_test.go`
  - Test GroupRows: verify outlineLevel set correctly
  - Test UngroupRows: verify outlineLevel decremented
  - Test nested grouping: multiple levels
  - Test GroupColumns and UngroupColumns
  - Test edge cases: max level (7), ungroup at level 0
  - Validation: All tests pass, >85% coverage

## Phase 5: Integration & Testing (1 week)

### 5.1 Roundtrip Tests
- [ ] **Create roundtrip test suite** in `sheet_operations_integration_test.go`
  - Test: Create workbook → insert rows → save → reopen → verify
    - Create workbook with formulas, named ranges, merged cells
    - Insert 5 rows at row 10
    - Save to temp file
    - Reopen file
    - Verify formulas updated correctly
    - Verify named ranges updated
    - Verify merged cells updated
  - Test: Similar roundtrip for DeleteRows
  - Test: Similar roundtrip for InsertColumns
  - Test: Similar roundtrip for DeleteColumns
  - Test: Roundtrip with grouping (GroupRows, save, reopen, verify outlineLevel)
  - Validation: All tests pass

### 5.2 Complex Workbook Tests
- [ ] **Create complex workbook test** in `sheet_operations_integration_test.go`
  - Create workbook with:
    - Multiple sheets
    - Cross-sheet formulas
    - Named ranges (global and sheet-scoped)
    - Merged cells
    - Conditional formatting
    - Data validation
    - Tables
  - Perform insert/delete operations
  - Verify all references updated correctly
  - Save and reopen
  - Verify document validity
  - Validation: Test passes

### 5.3 Fidelity Tests
- [ ] **Create fidelity test suite** in `sheet_operations_fidelity_test.go`
  - Test: Compare goffice InsertRows behavior with Excel behavior
    - Create identical workbooks in goffice and Excel
    - Perform same insert operation
    - Compare formula results, named range values, merged cell positions
    - Assert identical behavior (or document differences)
  - Test: Similar fidelity tests for DeleteRows, InsertColumns, DeleteColumns
  - Test: Fidelity for edge cases (absolute refs, cross-sheet refs, deleted ranges)
  - Validation: >95% fidelity with Excel

### 5.4 Performance Tests
- [ ] **Create performance benchmarks** in `sheet_operations_bench_test.go`
  - Benchmark: Insert 100 rows in 1000-row sheet
    - Target: <100ms
  - Benchmark: Delete 50 columns in 100-column sheet
    - Target: <50ms
  - Benchmark: Update 500 formulas
    - Target: <200ms
  - Benchmark: Large workbook (10 sheets, 10k rows each) insert operation
    - Target: <1s
  - Run benchmarks, verify targets met
  - Profile if performance issues found
  - Validation: All benchmarks meet targets

### 5.5 Error Handling Tests
- [ ] **Create error handling tests** in `sheet_operations_test.go`
  - Test: Invalid parameters (index = 0, count = 0, out of bounds)
  - Test: Nil sheet, nil document
  - Test: Operations on read-only document
  - Test: Recovery from partial failures (reference update fails mid-operation)
  - Validation: All error cases handled gracefully

### 5.6 Regression Tests
- [ ] **Run full existing test suite**
  - Run all tests in `spreadsheet/` package
  - Verify no regressions in existing functionality
  - Fix any regressions found
  - Validation: All existing tests still pass

## Phase 6: Documentation (0.5 weeks)

### 6.1 API Documentation
- [ ] **Add godoc comments** to all public functions
  - Document `Sheet.InsertRows`, `Sheet.DeleteRows`, `Sheet.InsertColumns`, `Sheet.DeleteColumns`
  - Document `Sheet.GroupRows`, `Sheet.UngroupRows`, `Sheet.GroupColumns`, `Sheet.UngroupColumns`
  - Include parameter descriptions, return value descriptions, error conditions
  - Add usage examples in godoc comments
  - Validation: godoc spreadsheet looks professional

### 6.2 Design Documentation
- [ ] **Update `spreadsheet/` package documentation**
  - Add overview of row/column operations in package comment
  - Explain reference update behavior
  - Document absolute vs relative reference handling
  - Add examples
  - Validation: Documentation is clear and helpful

### 6.3 Usage Examples
- [ ] **Create example in `examples/spreadsheet-operations/`**
  - Create example program demonstrating insert/delete operations
  - Show formula updates in action
  - Show grouping/ungrouping
  - Add comments explaining each step
  - Ensure example compiles and runs
  - Validation: Example works correctly

- [ ] **Add example to `examples/README.md`**
  - Document the spreadsheet-operations example
  - Explain what it demonstrates
  - Validation: README is updated

### 6.4 Developer Guide
- [ ] **Create developer guide** in `spreadsheet/OPERATIONS.md`
  - Explain two-phase operation architecture
  - Document cell reference parser design
  - Document formula rewriter approach
  - Explain reference update logic
  - Include diagrams if helpful
  - Add troubleshooting section
  - Validation: Guide is comprehensive

## Validation Checkpoints

### After Phase 1 (Cell Reference Infrastructure)
- [ ] All cell reference parser tests pass (200+ tests)
- [ ] Coverage >95% for cell_reference.go
- [ ] Manual testing: can parse all supported reference formats
- [ ] No compilation errors

### After Phase 2 (Formula Rewriting)
- [ ] All tokenizer tests pass (100+ tests)
- [ ] All reference updater tests pass (150+ tests)
- [ ] All formula rewriter tests pass (100+ tests)
- [ ] Coverage >90% for formula_rewriter.go and cell_reference_updater.go
- [ ] Manual testing: can rewrite complex formulas correctly
- [ ] No compilation errors

### After Phase 3 (Insert/Delete Operations)
- [ ] All physical operation tests pass
- [ ] All reference update integration tests pass
- [ ] All Sheet API tests pass
- [ ] Coverage >85% for sheet_operations.go
- [ ] Manual testing: can insert/delete rows/columns, references update
- [ ] No compilation errors

### After Phase 4 (Grouping Operations)
- [ ] All grouping tests pass
- [ ] Coverage >85% for grouping code
- [ ] Manual testing: can group/ungroup rows and columns
- [ ] No compilation errors

### After Phase 5 (Integration & Testing)
- [ ] All roundtrip tests pass
- [ ] All complex workbook tests pass
- [ ] All fidelity tests pass (>95% fidelity with Excel)
- [ ] All performance benchmarks meet targets
- [ ] All error handling tests pass
- [ ] All existing tests still pass (no regressions)
- [ ] No compilation errors

### After Phase 6 (Documentation)
- [ ] All public APIs have godoc comments
- [ ] Package documentation is complete
- [ ] Examples compile and run correctly
- [ ] Developer guide is comprehensive
- [ ] Documentation reviewed for clarity

## Success Criteria

**Functional**:
- [ ] Can insert/delete rows and columns programmatically
- [ ] All cell references update correctly (formulas, named ranges, merged cells, etc.)
- [ ] Absolute references preserved, relative references shifted
- [ ] Cross-sheet references update correctly
- [ ] Can group/ungroup rows and columns
- [ ] Roundtrip tests pass (save → open → verify)
- [ ] Fidelity tests pass (>95% match with Excel behavior)

**Quality**:
- [ ] All tests pass (unit, integration, fidelity, performance)
- [ ] Coverage >85% for new code
- [ ] No regressions in existing functionality
- [ ] All error cases handled gracefully
- [ ] Code follows goffice style guidelines
- [ ] golangci-lint passes with no errors

**Performance**:
- [ ] Insert 100 rows in 1000-row sheet: <100ms
- [ ] Delete 50 columns in 100-column sheet: <50ms
- [ ] Update 500 formulas: <200ms
- [ ] Large workbook operations: <1s

**Documentation**:
- [ ] All public APIs documented
- [ ] Examples provided and working
- [ ] Developer guide complete
- [ ] Design decisions documented

## Dependencies

**Blocked by**: None (all dependencies already exist in goffice)

**Blocks**: None (this is a standalone feature)

## Parallelizable Work

**Phase 1** tasks can be split:
- Cell reference parser (one person)
- R1C1 notation (another person)
- Column utilities (another person)

**Phase 2** tasks can be split:
- Formula tokenizer (one person)
- Reference shift logic (another person)
- Formula reconstruction (another person)

**Phase 3** tasks must be sequential (dependencies between sub-tasks)

**Phase 4** tasks can be split:
- Row grouping (one person)
- Column grouping (another person)

**Phase 5** tests can be parallelized:
- Roundtrip tests (one person)
- Fidelity tests (another person)
- Performance tests (another person)

**Phase 6** documentation tasks can be parallelized

## Risk Mitigation

**Risk: Formula parsing bugs**
- Mitigation: Extensive test suite (100+ formula tests), fidelity testing
- Contingency: Fail-safe mode (preserve original formula if parsing fails)

**Risk: Performance issues**
- Mitigation: Performance benchmarks, profiling, optimization
- Contingency: Document performance limitations, recommend batch operations

**Risk: Compatibility issues**
- Mitigation: Roundtrip testing, Excel validation, schema validation
- Contingency: Document known limitations, provide workarounds

**Risk: Edge case bugs**
- Mitigation: Comprehensive test coverage (400+ test cases total)
- Contingency: Incremental rollout, user feedback, quick bug fixes
