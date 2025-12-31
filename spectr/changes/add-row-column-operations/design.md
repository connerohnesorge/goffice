# Design: Excel Row/Column Insert/Delete Operations

## Context

Excel row/column operations are fundamental spreadsheet manipulations, but implementing them correctly requires solving a complex problem: **comprehensive cell reference updates**. When rows or columns are inserted or deleted, every reference in the entire workbook must be examined and potentially updated.

**Key challenges**:
1. **Reference diversity**: A1 notation, R1C1 notation, absolute ($A$1), relative (A1), mixed ($A1, A$1), cross-sheet (Sheet2!A1)
2. **Reference locations**: Formulas, named ranges, merged cells, conditional formatting, data validation, charts, tables
3. **Correctness**: Missing a single reference update corrupts the document
4. **Performance**: Large workbooks have thousands of formulas; full scan must be reasonably fast

**Constraints**:
- Must maintain Excel compatibility (documents open correctly in Excel)
- Must preserve existing behavior (no regressions)
- Zero external dependencies (standard library only)
- Must support all Excel versions (2007+)

## Goals / Non-Goals

**Goals**:
- Provide high-level API for row/column insert/delete operations
- Automatically update all cell references in workbook
- Support row/column grouping (outlining)
- Handle all reference types (A1, R1C1, absolute, relative, mixed, cross-sheet)
- Maintain document validity after operations
- Reasonable performance for typical workbooks (<1000 formulas)

**Non-Goals**:
- Formula evaluation/calculation (use cached values)
- Undo/redo support (Phase 2 feature)
- Transaction/rollback support (Phase 2 feature)
- Performance optimization for massive workbooks (>10,000 rows) in Phase 1
- R1C1 formula output (Phase 1 parses R1C1 input, outputs A1 notation only)
- Pivot table range updates (Phase 2 feature - requires pivot cache dependency analysis)
- Hyperlink target updates (Phase 2 feature - requires hyperlink API)
- UI components (server-side library)

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                    Sheet Operations                     │
│  InsertRows, DeleteRows, InsertColumns, DeleteColumns  │
└────────────────────┬────────────────────────────────────┘
                     │
          ┌──────────┴──────────┐
          │                     │
    ┏━━━━━▼━━━━━┓         ┏━━━━━▼━━━━━━━━━━━━┓
    ┃  Phase 1  ┃         ┃     Phase 2      ┃
    ┃ Physical  ┃         ┃   Reference      ┃
    ┃   Shift   ┃         ┃     Update       ┃
    ┗━━━━━┬━━━━━┛         ┗━━━━━┬━━━━━━━━━━━━┛
          │                     │
          │                     │
┌─────────▼──────────┐  ┌───────▼────────────────────┐
│  Row/Column XML    │  │  Cell Reference Updater    │
│  - Shift indices   │  │  - Parse references        │
│  - Insert/delete   │  │  - Determine if affected   │
│  - Update cells    │  │  - Apply shift             │
└────────────────────┘  └───────┬────────────────────┘
                                │
                    ┌───────────┴────────────┐
                    │                        │
         ┌──────────▼────────┐    ┌─────────▼──────────┐
         │  Formula Rewriter │    │  Range Updater     │
         │  - Tokenize       │    │  - Merged cells    │
         │  - Extract refs   │    │  - Cond formatting │
         │  - Shift refs     │    │  - Data validation │
         │  - Reconstruct    │    │  - Charts, tables  │
         └───────────────────┘    └────────────────────┘
```

## Key Components

### 1. SheetOperations (Core Coordinator)

**Responsibility**: Orchestrate insert/delete operations.

**Interface**:
```go
type SheetOperations struct {
    sheet *Sheet
}

func (so *SheetOperations) InsertRows(index, count uint32) error {
    // 1. Validate parameters
    // 2. Phase 1: Physical shift (shift row indices, create new rows)
    // 3. Phase 2: Reference update (scan and update all references)
    // 4. Mark document dirty
    return nil
}

func (so *SheetOperations) DeleteRows(index, count uint32) error {
    // Similar to InsertRows but delete instead
}

func (so *SheetOperations) InsertColumns(index, count uint32) error {
    // Similar to InsertRows but for columns
}

func (so *SheetOperations) DeleteColumns(index, count uint32) error {
    // Similar to InsertRows but delete columns
}

func (so *SheetOperations) GroupRows(start, end uint32) error {
    // Set outlineLevel attribute on rows
}

func (so *SheetOperations) UngroupRows(start, end uint32) error {
    // Decrement outlineLevel attribute
}
```

**Design rationale**: Separate operations struct keeps Sheet API clean. Operations can maintain state during complex multi-phase updates.

### 2. CellReferenceParser

**Responsibility**: Parse all cell reference formats into structured representation.

**Data structures**:
```go
type CellReference struct {
    Sheet      string  // Sheet name ("" = current sheet, "Sheet2", "'Sheet Name'")
    Column     uint32  // 1-based column index
    Row        uint32  // 1-based row index
    ColumnAbs  bool    // true if $A (absolute column)
    RowAbs     bool    // true if $1 (absolute row)
    IsR1C1     bool    // true if R1C1 notation
    R1C1RelCol int     // R1C1 relative column offset (0 if absolute)
    R1C1RelRow int     // R1C1 relative row offset (0 if absolute)
}

type RangeReference struct {
    Start CellReference
    End   CellReference
}

func ParseCellRef(ref string) (*CellReference, error)
func ParseRangeRef(ref string) (*RangeReference, error)
func (cr *CellReference) String() string  // Convert back to A1 notation
```

**Parsing algorithm (A1 notation)**:
```
1. Extract sheet name if present:
   - Check for "!" separator
   - Handle quoted names: 'Sheet Name'!A1
   - Store sheet name, continue with rest

2. Detect absolute markers:
   - Check for $ before column
   - Check for $ before row

3. Parse column:
   - Read letters (A-Z, AA-ZZ, etc.)
   - Convert to 1-based index: A=1, Z=26, AA=27, etc.

4. Parse row:
   - Read digits
   - Convert to 1-based index

5. Validate:
   - Column <= 16384 (XFD = max column)
   - Row <= 1048576 (max row)
```

**R1C1 notation algorithm**:
```
1. Detect R1C1 format: starts with "R"
2. Parse row component:
   - R alone = relative row (offset 0)
   - R5 = absolute row 5
   - R[-2] = relative row (offset -2)
3. Parse column component (similar)
4. Store in CellReference with IsR1C1=true
5. Convert to A1 on output (Phase 1 limitation)
```

**Phase 1 R1C1 strategy**:
- **Parse**: Full R1C1 parsing support to handle existing formulas in R1C1 notation
- **Update**: Shift operations work correctly on R1C1 references (convert to internal representation, shift, output)
- **Output**: All formulas written as A1 notation (String() always outputs A1)
- **Rationale**: Ensures correctness without complexity of R1C1 output formatting. Users opening documents with R1C1 formulas will see them converted to A1 (acceptable for Phase 1).

**Design rationale**: Structured representation simplifies shift operations. Parser handles all formats consistently. String() method enables output in A1 notation. R1C1 output is Phase 2 enhancement.

**Edge cases handled**:
- Single cell as range: `A1` → `A1:A1`
- Whole column: `A:A`
- Whole row: `5:5`
- Cross-sheet range: `Sheet1!A1:Sheet1!D10`
- Quoted sheet names with special chars: `'Sheet Name (2024)'!A1`

### 3. CellReferenceUpdater

**Responsibility**: Determine if reference should shift and apply shift operation.

**Interface**:
```go
type ShiftOperation struct {
    Type      ShiftType  // InsertRows, DeleteRows, InsertColumns, DeleteColumns
    Index     uint32     // Starting row/column (1-based)
    Count     uint32     // Number of rows/columns
    Worksheet string     // Target worksheet name
}

type CellReferenceUpdater struct {
    operation ShiftOperation
}

func (cru *CellReferenceUpdater) ShouldUpdate(ref *CellReference) bool {
    // Determine if reference is affected by operation
}

func (cru *CellReferenceUpdater) UpdateReference(ref *CellReference) *CellReference {
    // Apply shift to reference, return new reference
}

func (cru *CellReferenceUpdater) UpdateRange(rangeRef *RangeReference) *RangeReference {
    // Update both start and end of range
}
```

**Shift decision logic**:
```go
func (cru *CellReferenceUpdater) ShouldUpdate(ref *CellReference) bool {
    // 1. Check if reference is in affected worksheet
    if ref.Sheet != "" && ref.Sheet != cru.operation.Worksheet {
        return false  // Different sheet, not affected
    }

    // 2. Check if reference uses absolute addressing
    switch cru.operation.Type {
    case ShiftTypeInsertRows, ShiftTypeDeleteRows:
        if ref.RowAbs {
            return false  // Absolute row reference, don't shift
        }
        // Check if row >= operation.Index
        return ref.Row >= cru.operation.Index

    case ShiftTypeInsertColumns, ShiftTypeDeleteColumns:
        if ref.ColumnAbs {
            return false  // Absolute column reference, don't shift
        }
        // Check if column >= operation.Index
        return ref.Column >= cru.operation.Index
    }

    return false
}
```

**Shift application**:
```go
func (cru *CellReferenceUpdater) UpdateReference(ref *CellReference) *CellReference {
    if !cru.ShouldUpdate(ref) {
        return ref  // No change
    }

    newRef := *ref  // Copy

    switch cru.operation.Type {
    case ShiftTypeInsertRows:
        newRef.Row += cru.operation.Count

    case ShiftTypeDeleteRows:
        if ref.Row < cru.operation.Index + cru.operation.Count {
            // Reference is in deleted range - invalidate
            newRef.Row = 0  // Signal invalid reference
        } else {
            newRef.Row -= cru.operation.Count
        }

    case ShiftTypeInsertColumns:
        newRef.Column += cru.operation.Count

    case ShiftTypeDeleteColumns:
        if ref.Column < cru.operation.Index + cru.operation.Count {
            // Reference is in deleted range - invalidate
            newRef.Column = 0
        } else {
            newRef.Column -= cru.operation.Count
        }
    }

    return &newRef
}
```

**Design rationale**: Centralized shift logic ensures consistent behavior across all reference types. Absolute vs relative handling is explicit and testable.

### 4. FormulaRewriter

**Responsibility**: Extract cell references from formulas, update them, reconstruct formula.

**Tokenization approach**:
```go
type TokenType int

const (
    TokenTypeOperator   TokenType = iota  // +, -, *, /, =, <, >, etc.
    TokenTypeFunction                     // SUM, VLOOKUP, IF, etc.
    TokenTypeReference                    // A1, $A$1, Sheet2!A1, A1:D10
    TokenTypeLiteral                      // 42, "text", TRUE
    TokenTypeParen                        // (, )
    TokenTypeComma                        // ,
)

type Token struct {
    Type  TokenType
    Value string
}

func tokenizeFormula(formula string) []Token {
    // State machine to tokenize formula
    // - Skip leading "=" if present
    // - Identify strings (quoted, don't parse contents)
    // - Identify functions (name followed by "(")
    // - Identify references (A1 pattern, R1C1 pattern, sheet!ref)
    // - Identify operators and delimiters
}
```

**Rewriting algorithm**:
```go
func (fr *FormulaRewriter) Rewrite(formula string, operation ShiftOperation) string {
    // 1. Tokenize
    tokens := tokenizeFormula(formula)

    // 2. Update reference tokens
    updater := NewCellReferenceUpdater(operation)
    for i, token := range tokens {
        if token.Type == TokenTypeReference {
            // Parse reference
            ref, err := parseCellOrRangeRef(token.Value)
            if err != nil {
                continue  // Skip invalid references
            }

            // Update reference
            updatedRef := updater.Update(ref)

            // Replace token value
            tokens[i].Value = updatedRef.String()
        }
    }

    // 3. Reconstruct formula
    return reconstructFormula(tokens)
}

func reconstructFormula(tokens []Token) string {
    // Join tokens, add spaces where appropriate
    // Ensure leading "=" is present
}
```

**Example tokenization**:
```
Formula: =SUM(A1:A10)+B1*2
Tokens:
  1. Operator "="
  2. Function "SUM"
  3. Paren "("
  4. Reference "A1:A10"
  5. Paren ")"
  6. Operator "+"
  7. Reference "B1"
  8. Operator "*"
  9. Literal "2"
```

**Design rationale**: Token-based approach is robust against edge cases (string literals, function names that look like references). Extensible for future enhancements (formula validation, simplification).

**Edge cases handled**:
- String literals with cell references: `="The value in A1 is " & A1` (don't update "A1" in string)
- Function names: `ADDRESS(1,1)` (ADDRESS is function, not reference)
- Quoted sheet names: `='Sheet Name'!A1`
- Named ranges: `SUM(SalesData)` (check named range updates separately)

### 5. RangeUpdater

**Responsibility**: Update ranges in non-formula contexts (merged cells, conditional formatting, etc.).

**Interface**:
```go
type RangeUpdater struct {
    operation ShiftOperation
}

func (ru *RangeUpdater) UpdateMergeCells(sheet *Sheet) error {
    // Iterate mergeCells elements, update ref attributes
}

func (ru *RangeUpdater) UpdateConditionalFormatting(sheet *Sheet) error {
    // Iterate conditionalFormatting elements, update range attributes
}

func (ru *RangeUpdater) UpdateDataValidation(sheet *Sheet) error {
    // Iterate dataValidation elements, update sqref attributes
}

func (ru *RangeUpdater) UpdateChartRanges(sheet *Sheet) error {
    // Find chart parts, update series ranges
}

func (ru *RangeUpdater) UpdateTableRanges(sheet *Sheet) error {
    // Update table ref attributes
}
```

**Merge cells example**:
```go
func (ru *RangeUpdater) UpdateMergeCells(sheet *Sheet) error {
    mergeCells := sheet.Worksheet().MergeCells()
    if mergeCells == nil {
        return nil
    }

    for _, mergeCell := range mergeCells.MergeCellElements() {
        // Get current range reference
        ref := mergeCell.Ref()
        if ref == "" {
            continue
        }

        // Parse range
        rangeRef, err := ParseRangeRef(ref)
        if err != nil {
            continue
        }

        // Update range
        updater := NewCellReferenceUpdater(ru.operation)
        updatedRange := updater.UpdateRange(rangeRef)

        // Set new reference
        mergeCell.SetRef(updatedRange.String())
    }

    return nil
}
```

**Design rationale**: Separate updater for each context simplifies logic. Each updater knows how to navigate the specific element structure and update the appropriate attributes.

### 6. Shared Formula Handler

**Responsibility**: Handle shared formula updates (both master and child cells).

**Interface**:
```go
type SharedFormulaHandler struct {
    operation ShiftOperation
}

func (sfh *SharedFormulaHandler) IsSharedFormulaMaster(formula *elements.CellFormula) bool {
    // Check if t="shared" and ref attribute exists
    return formula.FormulaType() == elements.FormulaTypeShared && formula.Ref() != ""
}

func (sfh *SharedFormulaHandler) IsSharedFormulaChild(formula *elements.CellFormula) bool {
    // Check if t="shared" and ref attribute does not exist
    return formula.FormulaType() == elements.FormulaTypeShared && formula.Ref() == ""
}

func (sfh *SharedFormulaHandler) UpdateSharedFormulaMaster(formula *elements.CellFormula) error {
    // 1. Get formula text and rewrite it
    formulaText := formula.Formula()
    if formulaText != "" {
        rewriter := NewFormulaRewriter()
        updatedFormula := rewriter.Rewrite(formulaText, sfh.operation)
        formula.SetFormula(updatedFormula)
    }

    // 2. Update ref attribute (range reference)
    ref := formula.Ref()
    rangeRef, err := ParseRangeRef(ref)
    if err != nil {
        return err
    }
    updater := NewCellReferenceUpdater(sfh.operation)
    updatedRange := updater.UpdateRange(rangeRef)
    formula.SetRef(updatedRange.String())

    // 3. Preserve si attribute (never changes during operations)
    // No action needed - si attribute is preserved automatically

    return nil
}

func (sfh *SharedFormulaHandler) UpdateSharedFormulaChild(formula *elements.CellFormula) error {
    // Child cells inherit from master - no updates needed
    // The si attribute points to the master, which was already updated
    return nil
}
```

**Design rationale**: Shared formulas are an Excel optimization where one formula is shared across many cells. Only the master cell contains the formula text and ref range. Child cells reference the master via si (shared index). During row/column operations, we update the master's formula and ref, but children remain unchanged.

### 7. Array Formula Handler

**Responsibility**: Handle array formula updates (CSE arrays and dynamic arrays).

**Interface**:
```go
type ArrayFormulaHandler struct {
    operation ShiftOperation
}

func (afh *ArrayFormulaHandler) IsArrayFormulaAnchor(formula *elements.CellFormula) bool {
    // CSE array: t="array" and has formula text
    // Dynamic array: no t attribute, just normal formula
    return formula.FormulaType() == elements.FormulaTypeArray && formula.Formula() != ""
}

func (afh *ArrayFormulaHandler) IsArrayFormulaMember(formula *elements.CellFormula) bool {
    // CSE array member: t="array" but no formula text
    return formula.FormulaType() == elements.FormulaTypeArray && formula.Formula() == ""
}

func (afh *ArrayFormulaHandler) UpdateCSEArrayAnchor(formula *elements.CellFormula) error {
    // 1. Update formula text
    formulaText := formula.Formula()
    if formulaText != "" {
        rewriter := NewFormulaRewriter()
        updatedFormula := rewriter.Rewrite(formulaText, afh.operation)
        formula.SetFormula(updatedFormula)
    }

    // 2. Update ref attribute (array range)
    ref := formula.Ref()
    if ref != "" {
        rangeRef, err := ParseRangeRef(ref)
        if err != nil {
            return err
        }
        updater := NewCellReferenceUpdater(afh.operation)
        updatedRange := updater.UpdateRange(rangeRef)
        formula.SetRef(updatedRange.String())
    }

    return nil
}

func (afh *ArrayFormulaHandler) UpdateCSEArrayMember(formula *elements.CellFormula) error {
    // Member cells have same ref as anchor - update ref only
    ref := formula.Ref()
    if ref != "" {
        rangeRef, err := ParseRangeRef(ref)
        if err != nil {
            return err
        }
        updater := NewCellReferenceUpdater(afh.operation)
        updatedRange := updater.UpdateRange(rangeRef)
        formula.SetRef(updatedRange.String())
    }

    return nil
}

func (afh *ArrayFormulaHandler) UpdateDynamicArray(formula *elements.CellFormula) error {
    // Dynamic arrays (Excel 365+) are just normal formulas that spill
    // No ref attribute, just update formula text
    formulaText := formula.Formula()
    rewriter := NewFormulaRewriter()
    updatedFormula := rewriter.Rewrite(formulaText, afh.operation)
    formula.SetFormula(updatedFormula)

    return nil
}
```

**Design rationale**: CSE (Ctrl+Shift+Enter) array formulas occupy a fixed range with all cells sharing the same ref attribute. Only the anchor (top-left cell) contains the formula text. Dynamic arrays (Excel 365+) are simpler - they're just normal formulas that automatically spill. Both need special handling during row/column operations.

### 8. Data Table Formula Handler

**Responsibility**: Handle data table formula updates.

**Interface**:
```go
type DataTableFormulaHandler struct {
    operation ShiftOperation
}

func (dtfh *DataTableFormulaHandler) UpdateDataTable(formula *elements.CellFormula) error {
    updater := NewCellReferenceUpdater(dtfh.operation)

    // 1. Update r1 attribute (first input cell)
    r1 := formula.R1()
    if r1 != "" {
        ref, err := ParseCellRef(r1)
        if err != nil {
            return err
        }
        updatedRef := updater.UpdateReference(ref)
        formula.SetR1(updatedRef.String())
    }

    // 2. Update r2 attribute (second input cell)
    r2 := formula.R2()
    if r2 != "" {
        ref, err := ParseCellRef(r2)
        if err != nil {
            return err
        }
        updatedRef := updater.UpdateReference(ref)
        formula.SetR2(updatedRef.String())
    }

    // 3. Update ref attribute (data table range)
    ref := formula.Ref()
    if ref != "" {
        rangeRef, err := ParseRangeRef(ref)
        if err != nil {
            return err
        }
        updatedRange := updater.UpdateRange(rangeRef)
        formula.SetRef(updatedRange.String())
    }

    return nil
}
```

**Design rationale**: Data tables (What-If analysis) have special attributes r1/r2 for input cells and ref for the table range. All three need to be updated during row/column operations.

## Two-Phase Operation Algorithm

### Phase 1: Physical Shift

**InsertRows algorithm**:
```
1. Validate parameters:
   - index > 0 and index <= maxRow
   - count > 0
   - index + count <= maxRow

2. Get SheetData element

3. Shift existing rows:
   FOR each row in SheetData (in reverse order):
       IF row.Index >= index:
           newIndex = row.Index + count
           row.SetR(newIndex)

           FOR each cell in row:
               // Update cell reference
               oldRef = cell.R()  // e.g., "A5"
               newRef = updateCellCoordinate(oldRef, newRowIndex)
               cell.SetR(newRef)

4. Create new empty rows:
   FOR i = 0 to count-1:
       newRow = CreateRow()
       newRow.SetR(index + i)
       Insert newRow into SheetData at appropriate position

5. Update sheet dimensions (if dimension element exists)
```

**DeleteRows algorithm**:
```
1. Validate parameters:
   - index > 0 and index <= maxRow
   - count > 0
   - index + count <= maxRow

2. Get SheetData element

3. Delete rows:
   FOR each row in SheetData:
       IF row.Index >= index AND row.Index < index + count:
           Remove row from SheetData
       ELSE IF row.Index >= index + count:
           newIndex = row.Index - count
           row.SetR(newIndex)

           FOR each cell in row:
               oldRef = cell.R()
               newRef = updateCellCoordinate(oldRef, newRowIndex)
               cell.SetR(newRef)

4. Update sheet dimensions
```

**InsertColumns algorithm**: Similar to InsertRows but for columns
**DeleteColumns algorithm**: Similar to DeleteRows but for columns

### Phase 2: Reference Update

**Algorithm**:
```
1. Create ShiftOperation struct with operation details

2. Update formulas in all worksheets:
   FOR each worksheet in workbook:
       FOR each row in worksheet:
           FOR each cell in row:
               IF cell has formula:
                   formulaElement = cell.Formula()
                   formulaType = formulaElement.FormulaType()

                   IF formulaType == "normal":
                       // Normal formula - just update text
                       formula = formulaElement.Formula()
                       updatedFormula = FormulaRewriter.Rewrite(formula, operation)
                       formulaElement.SetFormula(updatedFormula)

                   ELSE IF formulaType == "shared":
                       // Shared formula
                       sharedHandler = NewSharedFormulaHandler(operation)
                       IF sharedHandler.IsSharedFormulaMaster(formulaElement):
                           // Master cell - update formula text and ref
                           sharedHandler.UpdateSharedFormulaMaster(formulaElement)
                       ELSE:
                           // Child cell - no updates needed
                           sharedHandler.UpdateSharedFormulaChild(formulaElement)

                   ELSE IF formulaType == "array":
                       // Array formula (CSE)
                       arrayHandler = NewArrayFormulaHandler(operation)
                       IF arrayHandler.IsArrayFormulaAnchor(formulaElement):
                           // Anchor cell - update formula text and ref
                           arrayHandler.UpdateCSEArrayAnchor(formulaElement)
                       ELSE:
                           // Member cell - update ref only
                           arrayHandler.UpdateCSEArrayMember(formulaElement)

                   ELSE IF formulaType == "dataTable":
                       // Data table formula
                       dataTableHandler = NewDataTableFormulaHandler(operation)
                       dataTableHandler.UpdateDataTable(formulaElement)

3. Update named ranges:
   definedNames = workbook.DefinedNames()
   FOR each name in definedNames:
       ref = name.Reference()

       // Check if reference (could be formula or range)
       IF isRangeReference(ref):
           updatedRef = CellReferenceUpdater.UpdateRange(ref, operation)
           name.SetReference(updatedRef)
       ELSE IF isFormula(ref):
           updatedFormula = FormulaRewriter.Rewrite(ref, operation)
           name.SetReference(updatedFormula)

4. Update merged cells:
   RangeUpdater.UpdateMergeCells(sheet)

5. Update conditional formatting:
   RangeUpdater.UpdateConditionalFormatting(sheet)

6. Update data validation:
   RangeUpdater.UpdateDataValidation(sheet)

7. Update chart ranges:
   RangeUpdater.UpdateChartRanges(sheet)

8. Update table ranges:
   RangeUpdater.UpdateTableRanges(sheet)
```

## Reference Update Examples

### Example 1: Insert 3 rows at row 5

**Before**:
- Formula in B11: `=SUM(A1:A10)`
- Formula in C15: `=B$5+C10`
- Named range: `SalesData = Sheet1!$A$1:$D$10`
- Merged cells: `A5:C5`

**After inserting 3 rows at row 5**:
- Formula in B14: `=SUM(A1:A13)` (row shifted 5→8, range end 10→13)
- Formula in C18: `=B$5+C13` (row shifted 15→18, C10→C13, B$5 unchanged)
- Named range: `SalesData = Sheet1!$A$1:$D$13` (end row shifted)
- Merged cells: `A8:C8` (shifted down)

**Logic**:
- B11 shifts to B14 (11 >= 5, so 11+3=14)
- Formula `A1:A10`: A1 not affected (< 5), A10 affected (>= 5) → A13
- Formula `B$5`: Absolute row, not shifted
- Formula `C10`: 10 >= 5 → C13
- Named range: Start $A$1 absolute (not shifted), end $D$10 → $D$13
- Merged `A5:C5`: Both 5 >= 5 → A8:C8

### Example 2: Delete 2 columns at column C

**Before**:
- Formula in E5: `=A1+B2+C3+D4`
- Formula in F10: `=$C5+E5`

**After deleting columns C and D**:
- Formula in C5: `=A1+B2+#REF!+#REF!` (C3 and D4 deleted)
- Formula in D10: `=$C5+C5` (F→D, E→C, $C unchanged)

**Logic**:
- E5 shifts to C5 (E = col 5, delete 2 at col 3, so 5-2=3)
- Formula: A1 not affected (col 1 < 3), B2 not affected (col 2 < 3)
- C3 deleted (3 >= 3 and 3 < 3+2), becomes #REF!
- D4 deleted (4 >= 3 and 4 < 3+2), becomes #REF!
- F10 → D10 (col 6 → 6-2=4)
- $C5: Absolute column, not shifted
- E5 → C5 (col 5 → 5-2=3)

## Performance Considerations

### Optimization Strategies

**1. Lazy reference scanning**:
```go
// Instead of scanning all sheets every time:
func (so *SheetOperations) buildReferenceIndex() map[string][]CellLocation {
    // Build index: which sheets reference which other sheets
    // Only scan referenced sheets during update
}
```

**2. Batch row/column creation**:
```go
// Instead of creating rows one-by-one:
func (so *SheetOperations) createRowsBatch(start, count uint32) []*elements.Row {
    rows := make([]*elements.Row, count)
    for i := uint32(0); i < count; i++ {
        rows[i] = elements.NewRow()
        rows[i].SetR(start + i)
    }
    return rows
}
```

**3. Formula cache**:
```go
// Cache parsed formulas to avoid re-tokenizing
type FormulaCache struct {
    cache map[string][]Token
}

func (fc *FormulaCache) Tokenize(formula string) []Token {
    if tokens, ok := fc.cache[formula]; ok {
        return tokens
    }
    tokens := tokenizeFormula(formula)
    fc.cache[formula] = tokens
    return tokens
}
```

**4. Parallel processing** (future):
```go
// Process sheets in parallel (read-only analysis phase)
// Update sequentially (write phase)
```

### Performance Targets

| Operation | Workbook Size | Target Time |
|-----------|---------------|-------------|
| Insert 100 rows | 1000 rows, 100 formulas | <100ms |
| Delete 50 cols | 100 cols, 200 formulas | <50ms |
| Update references | 500 formulas | <200ms |
| Large workbook ops | 10 sheets, 10k rows each | <1s |

## Testing Strategy

### Unit Tests

**Cell reference parser tests** (200+ cases):
- A1 notation: `A1`, `Z99`, `AA100`, `XFD1048576`
- Absolute: `$A$1`, `$A1`, `A$1`
- Cross-sheet: `Sheet2!A1`, `'Sheet Name'!A1`
- R1C1: `R1C1`, `R[-1]C[2]`, `RC[-1]`
- Ranges: `A1:D10`, `$A$1:$D$10`
- Invalid cases: `@@@`, `A`, `1`, `Sheet!`

**Formula tokenizer tests** (100+ formulas):
- Simple: `=A1+B1`, `=SUM(A1:A10)`
- Complex: `=IF(A1>0,SUM(B1:B10),0)`
- String literals: `="Value: "&A1`
- Cross-sheet: `=Sheet2!A1+Sheet3!B2`
- Named ranges: `=SUM(SalesData)`

**Reference updater tests** (150+ cases):
- Insert rows: absolute vs relative, before/after insert point
- Delete rows: references in deleted range, before/after delete point
- Insert columns: similar variations
- Delete columns: similar variations
- Edge cases: references at boundary, whole column/row refs

### Integration Tests

**Roundtrip tests**:
```go
func TestInsertRowsRoundtrip(t *testing.T) {
    // 1. Create workbook with formulas, named ranges, merged cells
    doc := createTestWorkbook()

    // 2. Insert rows
    doc.Sheets()[0].InsertRows(5, 3)

    // 3. Save to temp file
    doc.SaveAs("test_insert.xlsx")

    // 4. Reopen
    doc2 := spreadsheet.Open("test_insert.xlsx")

    // 5. Verify structure
    verifyFormulasUpdated(t, doc2)
    verifyNamedRangesUpdated(t, doc2)
    verifyMergeCellsUpdated(t, doc2)
}
```

**Fidelity tests**:
Compare goffice behavior with Excel behavior:
```go
func TestInsertRowsFidelity(t *testing.T) {
    // 1. Create identical workbooks in goffice and Excel
    // 2. Perform same insert operation in both
    // 3. Compare formulas, named ranges, merged cells
    // 4. Assert identical results
}
```

### Regression Tests

Ensure existing functionality unaffected:
```go
func TestExistingFunctionalityPreserved(t *testing.T) {
    // Run all existing spreadsheet tests
    // Ensure no regressions
}
```

## Migration Path

**Phase 1: Core infrastructure (weeks 1-3)**
1. Implement cell reference parser
2. Implement formula tokenizer
3. Implement reference updater
4. Unit tests for all parsers and updaters

**Phase 2: Operations (weeks 4-5)**
1. Implement InsertRows, DeleteRows
2. Implement InsertColumns, DeleteColumns
3. Integrate reference updates
4. Integration tests

**Phase 3: Grouping & finalization (week 6)**
1. Implement GroupRows, UngroupRows, GroupColumns, UngroupColumns
2. Fidelity testing
3. Performance testing
4. Documentation

## Risks & Mitigation

### Risk: Formula parsing edge cases
**Impact**: High (corrupt formulas)
**Probability**: Medium
**Mitigation**:
- Extensive test suite with real-world formulas
- Fidelity testing against Excel
- Fail-safe: If tokenization fails, preserve original formula (no update)

### Risk: Performance degradation
**Impact**: Medium (slow for large workbooks)
**Probability**: Low
**Mitigation**:
- Performance benchmarks
- Profiling and optimization
- Lazy scanning (only scan referenced sheets)

### Risk: Compatibility issues
**Impact**: Medium (documents don't open in Excel)
**Probability**: Low
**Mitigation**:
- Roundtrip testing with Excel
- Schema validation
- Test with multiple Excel versions

### Risk: Edge cases in reference updates
**Impact**: High (corrupt references)
**Probability**: Medium
**Mitigation**:
- Comprehensive test coverage
- Manual testing with complex workbooks
- Validation against Excel behavior

## Open Questions

1. **R1C1 output support**: Should we support outputting formulas in R1C1 notation, or always use A1?
   - **Decision**: Phase 1 parses R1C1 input (for existing formulas), but outputs A1 notation only. R1C1 output is Phase 2 enhancement. This approach handles documents with R1C1 formulas but normalizes to A1, which is acceptable for Phase 1.

2. **Invalid reference handling**: When deleting rows/columns, references in deleted range become invalid. How to represent?
   - **Decision**: Use `#REF!` error (Excel standard behavior)

3. **Performance optimization**: At what workbook size should we implement lazy scanning?
   - **Decision**: Implement simple approach first, optimize if benchmarks show issues (>1s for typical workbooks)

4. **Named range scope**: How to handle named ranges with same name at workbook and sheet scope?
   - **Decision**: Update both, let Excel resolve conflicts per its rules

5. **Chart series updates**: Charts reference ranges - how to access chart parts?
   - **Decision**: Navigate from WorksheetPart → DrawingsPart → ChartPart, update series ranges

## Success Metrics

**Functional**:
- 100% of supported reference types update correctly
- 100% of roundtrip tests pass
- 95%+ fidelity with Excel behavior

**Quality**:
- 85%+ code coverage for new code
- Zero regressions in existing tests
- All integration tests pass

**Performance**:
- Insert/delete operations complete in <100ms for typical workbooks
- Reference updates complete in <200ms for 500 formulas
- Large workbook operations complete in <1s
