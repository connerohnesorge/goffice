# Change: Add Row/Column Insert/Delete Operations

## Why

Excel row/column operations are **critical for spreadsheet manipulation** but completely absent from goffice. Users cannot programmatically insert or delete rows/columns, forcing manual editing or workarounds. This is a **Tier 1 priority gap** that blocks common automation scenarios:

- **Data import workflows**: Insert rows for new records
- **Report generation**: Add summary rows, delete empty rows
- **Template manipulation**: Expand/contract data sections dynamically
- **Batch processing**: Process files without manual intervention

**Current state**: Elements support exists (`Row`, `Column` in `spreadsheet/elements/`), cell manipulation works (`spreadsheet/cell.go`), but **no high-level API** for insert/delete operations. Users must manually shift row indices, which is error-prone and doesn't update cell references in formulas, named ranges, merged cells, conditional formatting, data validation, or charts.

**The hard part**: Cell reference updates. When inserting 3 rows at row 5:
- Formula `=SUM(A1:A10)` → `=SUM(A1:A13)` (range extends)
- Formula `=SUM($A$1:$A$10)` → `=SUM($A$1:$A$10)` (absolute refs unchanged)
- Named range `SalesData =A1:D10` → `A1:D13`
- Merged cells `A5:C5` → `A8:C8` (shifted down)
- Chart series ranges shift
- Conditional formatting ranges adjust
- Data validation ranges update

Without comprehensive reference updates, **documents become corrupt**.

## What Changes

**Core sheet operations API**:
- Add `Sheet.InsertRows(index, count)` - Insert rows at position, shift existing down
- Add `Sheet.DeleteRows(index, count)` - Delete rows, shift remaining up
- Add `Sheet.InsertColumns(index, count)` - Insert columns at position, shift existing right
- Add `Sheet.DeleteColumns(index, count)` - Delete columns, shift remaining left
- Add `Sheet.GroupRows(start, end)` - Set outline level for row grouping
- Add `Sheet.GroupColumns(start, end)` - Set outline level for column grouping
- Add `Sheet.UngroupRows(start, end)` - Remove row grouping
- Add `Sheet.UngroupColumns(start, end)` - Remove column grouping

**Cell reference update infrastructure**:
- Cell reference parser (A1 notation: `A1`, `$A$1`, `A$1`, `$A1`, `A1:D10`)
- R1C1 notation parser (`R1C1`, `R[-1]C[2]`) - **Phase 1: parse only, output as A1**
- Cross-sheet reference parsing (`Sheet2!A1`, `'Sheet Name'!A1:B5`)
- Formula tokenizer and rewriter
- Reference shift operations (insert/delete rows/cols)

**Comprehensive reference updates**:
- Scan and update all formulas in worksheet
- Update named ranges (workbook-level and sheet-scoped)
- Adjust merged cell ranges
- Update conditional formatting ranges
- Update data validation ranges
- Update chart data series ranges
- Update table ranges
- Handle absolute ($A$1) vs relative (A1) vs mixed ($A1, A$1) references

**Breaking changes**: None. This is entirely new functionality.

## Impact

**Affected specs**:
- `spreadsheet-document` (ADDED row/column operation requirements)
- `spreadsheet-cells` (ADDED cell reference update requirements)
- `spreadsheet-formulas` (ADDED formula rewriting requirements)

**New capabilities**:
- Programmatic row/column insertion and deletion
- Automatic cell reference updates across entire workbook
- Row/column grouping (outlining) for hierarchical data
- Batch document manipulation without corruption
- Template expansion/contraction

**Affected code**:
- `spreadsheet/sheet.go` - MODIFIED: Add operation methods
- `spreadsheet/sheet_operations.go` - NEW: Core insert/delete logic
- `spreadsheet/cell_reference.go` - NEW: Cell reference parser
- `spreadsheet/cell_reference_updater.go` - NEW: Reference update logic
- `spreadsheet/formula_rewriter.go` - NEW: Formula tokenization and rewriting
- `spreadsheet/range.go` - MODIFIED: Add range shift operations
- `spreadsheet/sheet_test.go` - MODIFIED: Add operation tests

**Dependencies**:
- Existing cell reference utilities in `spreadsheet/cell.go` (extend)
- Existing formula handling (if exists) or create new
- Element support for `Row.OutlineLevel`, `Column.OutlineLevel` (already exists)

## Key Design Decisions

### 1. Two-Phase Operation Architecture
**Decision**: Separate physical row/column manipulation from reference updates.

**Algorithm**:
```
Phase 1: Physical Shift
  - Adjust row/column indices in worksheet XML
  - Insert/delete row/column elements
  - Shift cell coordinates

Phase 2: Reference Update
  - Scan all formulas in all worksheets
  - Update named ranges (global and local)
  - Adjust merged cell ranges
  - Update conditional formatting
  - Update data validation
  - Update chart series ranges
  - Update table ranges
```

**Rationale**: Clear separation of concerns. Physical manipulation is straightforward. Reference updates are complex and cross-cutting. Separating them simplifies debugging and testing.

**Alternatives considered**:
- Single-pass update: Rejected - too complex, harder to test
- Lazy update on access: Rejected - leaves document in inconsistent state

### 2. Cell Reference Parser Design
**Decision**: Build dedicated parser supporting A1, R1C1, absolute/relative, cross-sheet.

**Parser structure**:
```go
type CellReference struct {
    Sheet        string  // Sheet name (empty if same sheet)
    Column       uint32  // 1-based column
    Row          uint32  // 1-based row
    ColumnAbs    bool    // true if $A
    RowAbs       bool    // true if $1
    IsR1C1       bool    // true if R1C1 notation
    R1C1RelCol   int     // R1C1 relative column offset
    R1C1RelRow   int     // R1C1 relative row offset
}

type RangeReference struct {
    Start CellReference
    End   CellReference
}
```

**Supported formats**:
- Simple: `A1`, `B5`, `XFD1048576`
- Absolute: `$A$1`, `$A1`, `A$1`
- Cross-sheet: `Sheet2!A1`, `'Sheet Name'!A1:B5`
- Ranges: `A1:D10`, `$A$1:$D$10`, `Sheet1!A1:Sheet1!D10`
- R1C1: `R1C1`, `R[-1]C[2]`, `RC[-1]`

**Rationale**: Comprehensive parsing is essential for correct updates. Excel supports many reference formats and we must handle all.

**Alternatives considered**:
- Regex-based parsing: Rejected - fragile, hard to maintain
- Reuse existing parsers: None exist with needed features
- Minimal parser (A1 only): Rejected - incomplete solution

### 3. Formula Rewriting Strategy
**Decision**: Tokenize formula, extract references, update each, reconstruct.

**Approach**:
```go
func UpdateFormula(formula string, operation ShiftOperation) string {
    // 1. Tokenize formula into operators, references, literals
    tokens := tokenizeFormula(formula)

    // 2. For each token that's a reference:
    for i, token := range tokens {
        if token.Type == TokenTypeReference {
            ref := parseCellReference(token.Value)
            updatedRef := shiftReference(ref, operation)
            tokens[i].Value = updatedRef.String()
        }
    }

    // 3. Reconstruct formula
    return reconstructFormula(tokens)
}
```

**Tokenization handles**:
- Function calls: `SUM()`, `VLOOKUP()`
- Operators: `+`, `-`, `*`, `/`, `&`, `=`, `<>`
- String literals: `"text"` (ignore)
- Cell references: `A1`, `$A$1`, `A1:D10`
- Named ranges: `SalesData` (check named range updates separately)

**Rationale**: Token-based approach is robust and extensible. Simple regex replacement is fragile (e.g., `"A1"` in string literal shouldn't be updated).

**Alternatives considered**:
- Full Excel formula parser: Rejected - too complex, not needed
- Regex replacement: Rejected - error-prone, misses edge cases
- AST-based approach: Rejected - overkill for reference updates

### 4. Shift Operation Design
**Decision**: Encapsulate shift details in operation struct.

```go
type ShiftType int

const (
    ShiftTypeInsertRows ShiftType = iota
    ShiftTypeDeleteRows
    ShiftTypeInsertColumns
    ShiftTypeDeleteColumns
)

type ShiftOperation struct {
    Type      ShiftType
    Index     uint32    // Starting row/column (1-based)
    Count     uint32    // Number of rows/columns
    Worksheet string    // Target worksheet name
}

func (op ShiftOperation) ShouldShiftReference(ref CellReference) bool {
    // Determine if reference should be updated based on:
    // - Is it in the affected worksheet?
    // - Is it in the affected range?
    // - Is it absolute (don't shift) or relative (shift)?
}

func (op ShiftOperation) ShiftReference(ref CellReference) CellReference {
    // Apply shift to reference
}
```

**Rationale**: Encapsulates shift logic, making it reusable across all reference types (formulas, named ranges, merged cells, etc.).

### 5. Reference Update Scope
**Decision**: Update all references in entire workbook, not just current sheet.

**Rationale**: Cross-sheet references exist (`Sheet2!A1`). When inserting rows in Sheet1, formulas in Sheet2 that reference Sheet1 must be updated. Workbook-scoped named ranges must update regardless of which sheet is modified.

**Performance consideration**: For large workbooks, scanning all sheets is expensive. Future optimization: Build reference index (which formulas reference which sheets) to avoid full scan.

### 6. Grouping (Outlining) API
**Decision**: Expose simple API wrapping `outlineLevel` attribute.

```go
func (s *Sheet) GroupRows(start, end uint32) error {
    // Set outlineLevel = 1 for rows start to end
    // If already grouped, increment level
}

func (s *Sheet) UngroupRows(start, end uint32) error {
    // Decrement outlineLevel for rows start to end
    // If level = 0, remove attribute
}
```

**Rationale**: Row/column grouping uses `outlineLevel` attribute (1-7 levels). Simple API wraps this. Excel handles expand/collapse UI.

**Out of scope**: Explicit collapse/expand state management (Excel infers from structure).

## Implementation Scope

**Phase 1: Cell Reference Infrastructure (1.5 weeks)**
- Implement cell reference parser (A1, R1C1, absolute/relative, cross-sheet)
- Implement range reference parser
- Add parser unit tests (200+ test cases for edge cases)
- Implement reference shift operations
- Add shift operation tests

**Phase 2: Formula Rewriting (1.5 weeks)**
- Implement formula tokenizer
- Implement reference extraction from tokens
- Implement reference update and formula reconstruction
- Add formula rewriting tests (100+ formulas)
- Handle edge cases (nested functions, string literals, named ranges)

**Phase 3: Insert/Delete Operations (1.5 weeks)**
- Implement `Sheet.InsertRows`, `Sheet.DeleteRows`
- Implement `Sheet.InsertColumns`, `Sheet.DeleteColumns`
- Implement physical row/column shifting
- Implement comprehensive reference updates (formulas, named ranges, merged cells, conditional formatting, data validation, charts, tables)
- Add operation tests (roundtrip: insert → save → open → verify)

**Phase 4: Grouping Operations (0.5 weeks)**
- Implement `Sheet.GroupRows`, `Sheet.UngroupRows`
- Implement `Sheet.GroupColumns`, `Sheet.UngroupColumns`
- Add grouping tests

**Phase 5: Integration & Testing (1 week)**
- Integration tests with complex workbooks
- Fidelity tests (compare with Excel's insert/delete behavior)
- Performance testing (large sheets: 10,000+ rows)
- Regression testing (ensure existing functionality unaffected)
- Documentation updates

**Total: 6 weeks**

## Out of Scope (Future Enhancements)

**Phase 2 features** (not in this proposal):
- Undo/redo support for operations
- Transaction support (rollback on error)
- Performance optimization for very large sheets (>100,000 rows)
- Batch operations API (`InsertRowsBatch` for multiple non-contiguous inserts)
- Smart reference update options (user-configurable update behavior)
- R1C1 formula output (Phase 1 parses R1C1 but outputs A1 notation only)
- Pivot table range updates (complex dependency tracking required, planned for dedicated pivot table proposal)
- Hyperlink target reference updates (requires hyperlink API, planned separately)

**Explicitly excluded**:
- Formula evaluation/calculation (use cached values)
- UI for row/column selection (server-side library)
- Excel compatibility validation (assumes standard Excel behavior)

**Rationale for Phase 2 deferrals**:
- **Pivot tables**: Require complex dependency analysis (pivot cache, field mappings, calculated items). Deferring to dedicated pivot table API proposal to avoid scope creep.
- **Hyperlinks**: Require hyperlink part navigation and relationship management. Better addressed in comprehensive hyperlink API.
- **R1C1 output**: Phase 1 focuses on correctness. R1C1 output is enhancement for specific use cases (VBA compatibility, localization).

## Success Criteria

**Functional**:
- [ ] Can insert 5 rows at row 10, all formulas below update correctly
- [ ] Can delete 3 columns at column C, all formulas adjust
- [ ] Absolute references ($A$1) remain unchanged
- [ ] Relative references (A1) shift appropriately
- [ ] Mixed references ($A1, A$1) behave correctly
- [ ] Named ranges update when affected ranges shift
- [ ] Merged cells adjust when within affected area
- [ ] Conditional formatting ranges update
- [ ] Data validation ranges update
- [ ] Chart data series ranges update
- [ ] Table ranges extend/shrink correctly
- [ ] Cross-sheet references update correctly
- [ ] Can group/ungroup rows and columns

**Quality**:
- [ ] Roundtrip test passes: insert rows → save → open → verify structure
- [ ] Fidelity test passes: goffice behavior matches Excel behavior
- [ ] All tests pass (unit, integration, fidelity)
- [ ] Code coverage >85% for new code
- [ ] No regressions in existing tests
- [ ] Documentation complete (API docs, examples, design notes)

**Performance**:
- [ ] Insert 100 rows in 1000-row sheet completes in <100ms
- [ ] Delete 50 columns in 100-column sheet completes in <50ms
- [ ] Reference update for 500 formulas completes in <200ms
- [ ] Large workbook (10 sheets, 10,000 rows each) operations complete in <1s

## Risk Mitigation

**Risk: Formula parsing bugs**
- **Impact**: High (corrupt formulas)
- **Probability**: Medium (formulas are complex)
- **Mitigation**: Extensive test suite with 100+ formula patterns, fidelity testing against Excel

**Risk: Edge cases in reference updates**
- **Impact**: High (corrupt documents)
- **Probability**: Medium (many reference types)
- **Mitigation**: Comprehensive test coverage for all reference types, cross-sheet refs, absolute/relative combinations

**Risk: Performance degradation**
- **Impact**: Medium (slow for large sheets)
- **Probability**: Low (operations are relatively simple)
- **Mitigation**: Performance benchmarks, profiling, optimize hot paths if needed

**Risk: Compatibility issues**
- **Impact**: Medium (documents don't open in Excel)
- **Probability**: Low (using standard OOXML)
- **Mitigation**: Roundtrip testing with Excel, validation against ECMA-376 schema

## Example Usage

```go
// Open workbook
doc, err := spreadsheet.Open("sales_report.xlsx", true)
if err != nil {
    log.Fatal(err)
}
defer doc.Close()

// Get first sheet
sheet := doc.Sheets()[0]

// Insert 3 blank rows at row 5 (for new data entry)
// Rows 5+ shift down to 8+
err = sheet.InsertRows(5, 3)
if err != nil {
    log.Fatal(err)
}

// Delete column C (remove obsolete data column)
// Columns D+ shift left to C+
err = sheet.DeleteColumns(3, 1)
if err != nil {
    log.Fatal(err)
}

// Group rows 10-20 for outlining (summary section)
err = sheet.GroupRows(10, 20)
if err != nil {
    log.Fatal(err)
}

// Formulas automatically updated:
// Before insert: =SUM(A1:A10)
// After insert:  =SUM(A1:A13)

// Before delete column: =B1+C1+D1
// After delete column:  =B1+C1     (C was deleted, D became C)

// Save changes
err = doc.Save()
if err != nil {
    log.Fatal(err)
}
```

## Dependencies

**Existing code to extend**:
- `spreadsheet/cell.go` - Cell reference utilities (extend with parser)
- `spreadsheet/range.go` - Range utilities (add shift operations)
- `spreadsheet/sheet.go` - Sheet API (add operation methods)
- Element support for `Row.OutlineLevel`, `Column.OutlineLevel` (already exists in schema-generated elements)

**New code to create**:
- `spreadsheet/sheet_operations.go` - Insert/delete implementation
- `spreadsheet/cell_reference.go` - Cell reference parser
- `spreadsheet/cell_reference_updater.go` - Reference shift logic
- `spreadsheet/formula_rewriter.go` - Formula tokenization and rewriting

**No external dependencies**: All functionality implemented using standard library and existing goffice packages.
