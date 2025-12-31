# Proposal Brief: Excel Row/Column Insert/Delete Operations

## Priority: TIER 1 - Critical
**Impact**: 9/10  
**Effort**: 2-3 weeks  
**Status**: Elements support, NO high-level API

## Gap Analysis

### Current State
**What exists:**
- ✅ **Cell manipulation** works (spreadsheet/cell.go)
- ✅ **Row/Column elements** exist (spreadsheet/elements/elements.go)
- ✅ **Sheet structure** supports rows/columns
- ✅ **Can read** documents with custom row heights, column widths

**What's missing:**
- ❌ No `Sheet.InsertRows(index, count)` method
- ❌ No `Sheet.DeleteRows(index, count)` method
- ❌ No `Sheet.InsertColumns(index, count)` method
- ❌ No `Sheet.DeleteColumns(index, count)` method
- ❌ No `Sheet.GroupRows(start, end)` for outlining
- ❌ No `Sheet.UngroupRows(start, end)`
- ❌ No automatic cell reference updates (formulas, named ranges, charts)

### OpenXML-SDK Comparison
**Microsoft's Excel Interop API provides:**
```csharp
worksheet.Rows[5].Insert(Shift:=xlShiftDown)
worksheet.Rows["5:7"].Delete(Shift:=xlShiftUp)
worksheet.Columns["C"].Insert(Shift:=xlShiftRight)
```

**Open-XML-SDK** doesn't provide high-level insert/delete (it's low-level) BUT applications using it need this functionality.

### Why This Is Hard
**Cell Reference Updates**: When inserting/deleting rows/columns, must update:
1. **Formulas**: `=SUM(A1:A10)` → `=SUM(A1:A12)` after inserting 2 rows at row 5
2. **Named Ranges**: `SalesData =A1:D10` → `A1:D12`
3. **Merged Cells**: `A1:C3` → adjust if within affected range
4. **Conditional Formatting**: Ranges must be updated
5. **Data Validation**: Applied ranges must shift
6. **Charts**: Data ranges must shift
7. **Tables**: Table ranges must extend/shrink

## Implementation Scope

### Files to Create/Modify

1. **spreadsheet/sheet_operations.go** (NEW)
   - `InsertRows(index uint, count uint) error`
   - `DeleteRows(index uint, count uint) error`
   - `InsertColumns(index uint, count uint) error`
   - `DeleteColumns(index uint, count uint) error`
   - `GroupRows(start uint, end uint) error`
   - `GroupColumns(start uint, end uint) error`

2. **spreadsheet/cell_reference_updater.go** (NEW)
   - `UpdateCellReferences(refString string, operation ShiftOperation) string`
   - Handle A1 notation, R1C1 notation, absolute/relative references
   - Handle range updates

3. **spreadsheet/sheet.go** (MODIFY)
   - Add insert/delete methods delegating to sheet_operations.go

### Core Design Decisions

**1. Two-Phase Operation**
```
Phase 1: Insert/Delete Physical Rows/Columns
  - Shift row/column indices in worksheet XML
  - Add/remove row/column elements
  
Phase 2: Update All References
  - Scan all formulas in worksheet
  - Scan all named ranges
  - Scan conditional formatting ranges
  - Scan data validation ranges
  - Scan merged cell ranges
  - Update chart data series ranges
```

**2. Cell Reference Parser**
Need robust parsing for:
- Simple references: `A1`, `$A$1`, `A$1`, `$A1`
- Ranges: `A1:D10`, `$A$1:$D$10`
- Cross-sheet: `Sheet2!A1`, `'Sheet Name'!A1:B5`
- Named ranges: `SalesData`, `_MyRange`
- R1C1 notation: `R1C1`, `R[-1]C[+2]`

**3. Shift Operation Types**
```go
type ShiftOperation struct {
    Type      ShiftType  // InsertRows, DeleteRows, InsertCols, DeleteCols
    Index     uint       // Starting index (0-based)
    Count     uint       // Number of rows/cols
    Worksheet string     // Which sheet (for cross-sheet refs)
}
```

**4. Formula Rewriting**
Use existing formula parsing or regex-based approach:
- Extract all cell references from formula
- Update each reference based on shift operation
- Reconstruct formula

**5. Grouped Rows/Columns (Outlining)**
Elements support `outlineLevel` attribute:
```go
row.SetOutlineLevel(types.NewUInt32Value(1))  // Level 1 grouping
```

## References

### Existing Code
- `spreadsheet/sheet.go` - Sheet API
- `spreadsheet/cell.go` - Cell reference handling
- `spreadsheet/elements/elements.go` - Row, Column, MergeCell elements
- `spreadsheet/formulas.go` - Formula handling (if exists)

### OpenXML-SDK References
Check @OpenXML-SDK/ for:
- How they handle cell reference parsing (if at all)
- Row/Column element structure
- Outline level attributes

### Similar Implementations
- **excelize** (Go Excel library) has insert/delete row/column - can reference for approach
- **openpyxl** (Python) has comprehensive insert/delete logic

## Success Criteria

- [ ] Can insert 5 rows at row 10 in a sheet with formulas
- [ ] All formulas below row 10 are updated correctly
- [ ] Absolute references ($A$1) are NOT shifted
- [ ] Relative references (A1) ARE shifted appropriately
- [ ] Mixed references ($A1, A$1) are handled correctly
- [ ] Named ranges are updated
- [ ] Merged cells are adjusted
- [ ] Conditional formatting ranges are updated
- [ ] Data validation ranges are updated
- [ ] Chart data series ranges are updated
- [ ] Can delete rows and formulas adjust correctly
- [ ] Can group rows and outline level is set
- [ ] Roundtrip test: insert rows → save → open → verify structure

## Spec Capabilities to Update

- **spreadsheet-document**: Add row/column operation requirements
- **spreadsheet-cells**: Add cell reference update requirements
- **spreadsheet-formulas**: Add formula rewriting requirements

## Dependencies

- Need cell reference parser (may need to create)
- Need formula tokenizer (may exist in spreadsheet/formulas.go)

## Out of Scope (Phase 2)

- Undo/Redo for operations
- Transaction support (rollback on error)
- Performance optimization for large sheets (batch updates)
- UI for row/column selection

## Example Usage (Target API)

```go
sheet := doc.Sheets()[0]

// Insert 3 rows starting at row 5 (shifts rows 5+ down)
sheet.InsertRows(4, 3)  // 0-based index

// Delete 2 columns starting at column C
sheet.DeleteColumns(2, 2)  // C and D

// Group rows 5-10 for outlining
sheet.GroupRows(4, 9)

// Formulas automatically updated:
// Before: =SUM(A1:A10)
// After:  =SUM(A1:A13)  (if inserted 3 rows above A10)

doc.Save()
```

## Algorithm Outline

```go
func (s *Sheet) InsertRows(index uint, count uint) error {
    // 1. Validate index is within bounds
    
    // 2. Shift existing rows down
    for i := s.MaxRowIndex(); i >= index; i-- {
        row := s.Row(i)
        row.SetR(types.NewUInt32Value(i + 1 + count))
    }
    
    // 3. Create new empty rows
    for i := uint(0); i < count; i++ {
        s.CreateRow(index + i)
    }
    
    // 4. Update all cell references in workbook
    updater := NewCellReferenceUpdater(ShiftOperation{
        Type: ShiftTypeInsertRows,
        Index: index,
        Count: count,
        Worksheet: s.Name(),
    })
    
    // Update formulas
    for row := range s.Rows() {
        for cell := range row.Cells() {
            if cell.Formula() != nil {
                newFormula := updater.UpdateFormula(cell.Formula().Text())
                cell.Formula().SetText(newFormula)
            }
        }
    }
    
    // Update named ranges
    for name := range s.Document().DefinedNames() {
        newRef := updater.UpdateCellReference(name.Reference())
        name.SetReference(newRef)
    }
    
    // Update merged cells
    for merge := range s.MergeCells() {
        newRange := updater.UpdateRange(merge.Ref())
        merge.SetRef(newRange)
    }
    
    // Update conditional formatting, data validation, charts...
    
    return nil
}
```
