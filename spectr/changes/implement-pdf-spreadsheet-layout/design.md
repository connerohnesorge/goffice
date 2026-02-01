# Proposal: Implement PDF Spreadsheet Layout with Actual Column Widths and Row Heights

## Summary

Update the PDF spreadsheet renderer to use actual column widths and row heights from the layout instead of hardcoded default values.

## Background

Currently, the PDF spreadsheet page layout code uses hardcoded default values for column widths (72 points) and row heights (15 points). This results in inaccurate rendering of spreadsheets where the actual layout dimensions from the Excel file should be used.

**Reference TODOs:**
- `pdf/spreadsheet/page_layout.go:1124` - "TODO: Use actual column widths from layout"
- `pdf/spreadsheet/page_layout.go:1132` - "TODO: Use actual row heights from layout"

Current code:
```go
// This is a simplified calculation
// TODO: Use actual column widths from layout
x += 72.0 // Default column width in points

// This is a simplified calculation
// TODO: Use actual row heights from layout
y += 15.0 // Default row height in points
```

## Motivation

Accurate PDF rendering of spreadsheets requires using the actual dimensions defined in the Excel file:
- Column widths may be explicitly set or auto-fitted
- Row heights may be manually adjusted or auto-fitted
- Using defaults results in misaligned content and poor visual fidelity

## Technical Design

### Current Implementation

The `calculateCellBounds()` function in `page_layout.go` currently:
1. Uses hardcoded 72.0 points for all column widths
2. Uses hardcoded 15.0 points for all row heights
3. Calculates X/Y positions by summing these defaults

### Proposed Implementation

Extend the layout calculation to:
1. Access the spreadsheet's column width information
2. Access the spreadsheet's row height information
3. Use actual dimensions when available
4. Fall back to defaults only when no explicit dimension is set

### Data Sources

Column widths can be obtained from:
- `worksheet.GetColumnWidth(colIndex)` or similar API
- Column definitions in the worksheet XML
- Default column width from worksheet properties

Row heights can be obtained from:
- `worksheet.GetRowHeight(rowIndex)` or similar API
- Row definitions in the worksheet XML
- Default row height from worksheet properties

### Code Structure

```go
func (r *SpreadsheetRenderer) calculateCellBounds(
    pageInfo PageInfo,
    fromRow, fromCol, toRow, toCol int,
) (x, y, width, height float64) {
    margins := r.options.Margins
    
    // Calculate X position (sum of actual column widths)
    x = margins.Left
    for col := pageInfo.ColStart; col < fromCol && col <= pageInfo.ColEnd; col++ {
        colWidth := r.getColumnWidth(col) // NEW: Get actual width
        x += colWidth
    }
    
    // Calculate Y position (sum of actual row heights)
    y = margins.Top
    for row := pageInfo.RowStart; row < fromRow && row <= pageInfo.RowEnd; row++ {
        rowHeight := r.getRowHeight(row) // NEW: Get actual height
        y += rowHeight
    }
    
    // Calculate width using actual column widths
    width = 0.0
    for col := fromCol; col < toCol; col++ {
        width += r.getColumnWidth(col) // NEW: Get actual width
    }
    
    // Calculate height using actual row heights
    height = 0.0
    for row := fromRow; row < toRow; row++ {
        height += r.getRowHeight(row) // NEW: Get actual height
    }
    
    return x, y, width, height
}

// NEW: Get actual column width
func (r *SpreadsheetRenderer) getColumnWidth(col int) float64 {
    if width, ok := r.columnWidths[col]; ok {
        return width
    }
    // Try to get from worksheet
    if r.worksheet != nil {
        if width := r.worksheet.GetColumnWidth(col); width > 0 {
            r.columnWidths[col] = width
            return width
        }
    }
    // Fall back to default
    return 72.0
}

// NEW: Get actual row height
func (r *SpreadsheetRenderer) getRowHeight(row int) float64 {
    if height, ok := r.rowHeights[row]; ok {
        return height
    }
    // Try to get from worksheet
    if r.worksheet != nil {
        if height := r.worksheet.GetRowHeight(row); height > 0 {
            r.rowHeights[row] = height
            return height
        }
    }
    // Fall back to default
    return 15.0
}
```

## Requirements

### SHALL Requirements

#### Requirement: Actual Column Width Usage
The spreadsheet PDF renderer SHALL use actual column widths from the Excel layout when available.

##### Scenario: Explicit Column Widths
Given a spreadsheet with columns of varying widths (e.g., 100pt, 150pt, 80pt)
When rendered to PDF
Then the PDF SHALL reflect the actual column widths, not default 72pt

##### Scenario: Auto-fitted Columns
Given a spreadsheet with auto-fitted columns
When rendered to PDF
Then the PDF SHALL use the calculated auto-fit widths

##### Scenario: Default Column Width
Given a spreadsheet with no explicit column widths
When rendered to PDF
Then the PDF SHALL use the worksheet's default column width or fall back to 72pt

#### Requirement: Actual Row Height Usage
The spreadsheet PDF renderer SHALL use actual row heights from the Excel layout when available.

##### Scenario: Explicit Row Heights
Given a spreadsheet with rows of varying heights (e.g., 20pt, 30pt, 15pt)
When rendered to PDF
Then the PDF SHALL reflect the actual row heights, not default 15pt

##### Scenario: Auto-fitted Rows
Given a spreadsheet with auto-fitted rows
When rendered to PDF
Then the PDF SHALL use the calculated auto-fit heights

##### Scenario: Wrapped Text Rows
Given a spreadsheet with wrapped text causing increased row height
When rendered to PDF
Then the PDF SHALL use the expanded row height

#### Requirement: Caching for Performance
The implementation SHALL cache column widths and row heights to avoid repeated lookups.

##### Scenario: Repeated Cell Rendering
Given a spreadsheet with many cells in the same column
When rendered to PDF
Then column widths SHALL be cached after first lookup

### SHOULD Requirements

#### Requirement: Width/Height Validation
The implementation SHOULD validate that retrieved widths/heights are within reasonable bounds.

#### Requirement: Units Consistency
The implementation SHOULD ensure all measurements use consistent units (points).

## Testing Strategy

### Unit Tests
- Test column width retrieval from worksheet
- Test row height retrieval from worksheet
- Test caching mechanism
- Test fallback to defaults
- Test with various measurement units

### Integration Tests
- End-to-end PDF rendering with explicit column widths
- End-to-end PDF rendering with explicit row heights
- End-to-end with auto-fitted dimensions

### Visual Tests
- Compare rendered PDF with expected layout
- Verify alignment of content in cells
- Verify page breaks at correct positions

## Implementation Plan

1. Add column width lookup method
2. Add row height lookup method
3. Update calculateCellBounds() to use actual dimensions
4. Add caching for performance
5. Add unit tests
6. Add integration tests
7. Visual regression testing

## Related Changes

- `pdf/spreadsheet/page_layout.go` - Main implementation
- `pdf/spreadsheet/renderer.go` - May need worksheet access updates
- `spreadsheet/worksheet.go` - Ensure width/height APIs are available

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Missing width/height APIs | High | Add methods to worksheet API if needed |
| Unit conversion issues | Medium | Ensure consistent point-based units |
| Performance with many lookups | Low | Implement caching mechanism |
| Zero/negative dimensions | Low | Add validation and fallback |

## Acceptance Criteria

- [ ] Actual column widths used when available
- [ ] Actual row heights used when available
- [ ] Proper fallback to defaults when not set
- [ ] Caching implemented for performance
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] Visual fidelity matches Excel output
