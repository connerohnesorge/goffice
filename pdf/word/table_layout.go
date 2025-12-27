// table_layout.go provides table layout algorithms for Word document rendering.
// It handles column width calculation, row height calculation, cell merging,
// and table pagination.

package word

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TableLayout represents the computed layout of a table.
type TableLayout struct {
	// Table is the source table element.
	Table *elements.Table

	// X and Y are the top-left coordinates of the table.
	X, Y float64

	// TotalWidth is the total width of the table.
	TotalWidth float64

	// TotalHeight is the total height of the table.
	TotalHeight float64

	// ColumnWidths contains the computed width of each column.
	ColumnWidths []float64

	// Rows contains the layout for each row.
	Rows []RowLayout
}

// RowLayout represents the layout of a single table row.
type RowLayout struct {
	// Row is the source row element.
	Row *elements.TableRow

	// Y is the top Y coordinate of this row.
	Y float64

	// Height is the computed height of this row.
	Height float64

	// Cells contains the layout for each cell.
	Cells []CellLayout
}

// CellLayout represents the layout of a single table cell.
type CellLayout struct {
	// Cell is the source cell element.
	Cell *elements.TableCell

	// X and Y are the top-left coordinates of the cell content area.
	X, Y float64

	// Width and Height are the dimensions of the cell.
	Width, Height float64

	// ColSpan is the number of columns this cell spans.
	ColSpan int

	// RowSpan is the number of rows this cell spans.
	RowSpan int

	// IsVMergeContinue indicates this cell continues a vertical merge.
	IsVMergeContinue bool

	// ContentLines contains the laid-out paragraph lines for this cell.
	ContentLines [][]layout.LayoutLine
}

// LayoutTable computes the complete layout for a table.
func (r *WordRenderer) LayoutTable(
	table *elements.Table,
	x, y, maxWidth float64,
) (*TableLayout, error) {
	tl := &TableLayout{
		Table: table,
		X:     x,
		Y:     y,
	}

	// 1. Calculate column widths
	if err := r.calculateColumnWidths(tl, maxWidth); err != nil {
		return nil, err
	}

	// 2. Layout each row
	currentY := y
	for row := range table.Rows() {
		rowLayout, err := r.layoutRow(
			tl,
			row,
			currentY,
		)
		if err != nil {
			return nil, err
		}
		tl.Rows = append(tl.Rows, rowLayout)
		currentY -= rowLayout.Height
	}

	// 3. Handle vertical cell merges (adjust heights)
	r.resolveVerticalMerges(tl)

	// 4. Calculate total dimensions
	tl.TotalWidth = 0
	for _, w := range tl.ColumnWidths {
		tl.TotalWidth += w
	}
	tl.TotalHeight = y - currentY

	return tl, nil
}

// calculateColumnWidths determines the width of each column.
func (r *WordRenderer) calculateColumnWidths(
	tl *TableLayout,
	maxWidth float64,
) error {
	// Get column count from table grid or first row
	grid := tl.Table.TableGrid()
	colCount := 0

	if grid != nil {
		// Count grid columns
		for range grid.GridColumns() {
			colCount++
		}
	}

	// If no grid, infer from first row
	if colCount == 0 {
		firstRow := tl.Table.GetRow(0)
		if firstRow != nil {
			for cell := range firstRow.Cells() {
				props := cell.TableCellProperties()
				span := 1
				if props != nil {
					span = props.GridSpan()
				}
				colCount += span
			}
		}
	}

	if colCount == 0 {
		return fmt.Errorf("table has no columns")
	}

	// Get table width properties
	tblProps := tl.Table.TableProperties()
	var tableWidth float64
	var widthType elements.TableWidthType

	if tblProps != nil {
		if tw := tblProps.TableWidth(); tw != nil {
			widthType = tw.Type()
			tableWidth = float64(tw.Width())
		}
	}

	// Determine actual table width
	actualWidth := maxWidth
	switch widthType {
	case elements.TableWidthTypeDxa:
		// Fixed width in twips
		actualWidth = tableWidth / 20.0
		if actualWidth > maxWidth {
			actualWidth = maxWidth
		}
	case elements.TableWidthTypePct:
		// Percentage (5000 = 100%)
		pct := tableWidth / 5000.0
		actualWidth = maxWidth * pct
	case elements.TableWidthTypeAuto,
		elements.TableWidthTypeNil:
		// Use max width
		actualWidth = maxWidth
	}

	// Get column widths from grid
	tl.ColumnWidths = make([]float64, colCount)
	explicitWidths := make([]bool, colCount)
	totalExplicitWidth := 0.0

	if grid != nil {
		i := 0
		for gc := range grid.GridColumns() {
			if i >= colCount {
				break
			}
			if w := gc.Width(); w > 0 {
				tl.ColumnWidths[i] = float64(
					w,
				) / 20.0
				explicitWidths[i] = true
				totalExplicitWidth += tl.ColumnWidths[i]
			}
			i++
		}
	}

	// Distribute remaining width to auto-width columns
	autoCount := 0
	for i := 0; i < colCount; i++ {
		if !explicitWidths[i] {
			autoCount++
		}
	}

	if autoCount > 0 {
		remainingWidth := actualWidth - totalExplicitWidth
		if remainingWidth > 0 {
			autoWidth := remainingWidth / float64(
				autoCount,
			)
			for i := 0; i < colCount; i++ {
				if !explicitWidths[i] {
					tl.ColumnWidths[i] = autoWidth
				}
			}
		}
	}

	// Scale if total exceeds maxWidth
	totalWidth := 0.0
	for _, w := range tl.ColumnWidths {
		totalWidth += w
	}
	if totalWidth > maxWidth {
		scale := maxWidth / totalWidth
		for i := range tl.ColumnWidths {
			tl.ColumnWidths[i] *= scale
		}
	}

	return nil
}

// layoutRow computes the layout for a single row.
func (r *WordRenderer) layoutRow(
	tl *TableLayout,
	row *elements.TableRow,
	y float64,
) (RowLayout, error) {
	rl := RowLayout{
		Row:    row,
		Y:      y,
		Height: 0,
	}

	// Get row height from properties
	minHeight := 0.0
	heightRule := elements.HeightRuleAuto
	if props := row.TableRowProperties(); props != nil {
		if trHeight := props.GetElement("trHeight", elements.NamespaceWML); trHeight != nil {
			if valAttr, found := trHeight.GetAttribute("val", elements.NamespaceWML); found {
				var h int
				fmt.Sscanf(
					valAttr.Value(),
					"%d",
					&h,
				)
				minHeight = float64(h) / 20.0
			}
			if ruleAttr, found := trHeight.GetAttribute("hRule", elements.NamespaceWML); found {
				heightRule = elements.HeightRule(
					ruleAttr.Value(),
				)
			}
		}
	}

	// Layout cells
	currentCol := 0
	currentX := tl.X

	for cell := range row.Cells() {
		if currentCol >= len(tl.ColumnWidths) {
			break
		}

		// Get cell properties
		props := cell.TableCellProperties()
		colSpan := 1
		isVMergeContinue := false

		if props != nil {
			colSpan = props.GridSpan()
			if vm := props.VerticalMerge(); vm != nil {
				if vm.Type() == elements.VerticalMergeContinue {
					isVMergeContinue = true
				}
			}
		}

		// Calculate cell width (sum of spanned columns)
		cellWidth := 0.0
		for i := 0; i < colSpan && currentCol+i < len(tl.ColumnWidths); i++ {
			cellWidth += tl.ColumnWidths[currentCol+i]
		}

		// Layout cell content
		cellLayout := CellLayout{
			Cell:             cell,
			X:                currentX,
			Y:                y,
			Width:            cellWidth,
			ColSpan:          colSpan,
			IsVMergeContinue: isVMergeContinue,
		}

		// Only layout content if not a vMerge continue cell
		if !isVMergeContinue {
			cellHeight, lines, err := r.layoutCellContent(
				cell,
				cellWidth,
			)
			if err != nil {
				return rl, err
			}
			cellLayout.Height = cellHeight
			cellLayout.ContentLines = lines

			// Update row height
			if cellHeight > rl.Height {
				rl.Height = cellHeight
			}
		}

		rl.Cells = append(rl.Cells, cellLayout)
		currentX += cellWidth
		currentCol += colSpan
	}

	// Apply height rule
	switch heightRule {
	case elements.HeightRuleExact:
		rl.Height = minHeight
	case elements.HeightRuleAtLeast:
		if rl.Height < minHeight {
			rl.Height = minHeight
		}
	case elements.HeightRuleAuto:
		// Use computed height, but respect minimum
		if minHeight > 0 &&
			rl.Height < minHeight {
			rl.Height = minHeight
		}
	}

	// Ensure minimum row height
	if rl.Height < 12.0 {
		rl.Height = 12.0
	}

	return rl, nil
}

// layoutCellContent layouts the paragraphs within a cell.
func (r *WordRenderer) layoutCellContent(
	cell *elements.TableCell,
	cellWidth float64,
) (float64, [][]layout.LayoutLine, error) {
	// Add padding (default 0.05 inch = 3.6 points on each side)
	padding := 3.6
	contentWidth := cellWidth - 2*padding

	if contentWidth < 10 {
		contentWidth = 10
	}

	var allLines [][]layout.LayoutLine
	totalHeight := 0.0

	for p := range cell.Paragraphs() {
		// Convert to layout paragraph
		lp := r.convertToLayoutParagraph(p)

		// Layout paragraph
		opts := layout.DefaultLayoutOptions()
		opts.MaxWidth = contentWidth
		lines := r.engine.LayoutParagraph(
			lp,
			opts,
		)

		// Calculate paragraph height
		paraHeight := 0.0
		for _, line := range lines {
			paraHeight += line.Height
		}

		// Add spacing
		props := p.Properties()
		if props != nil {
			if sp := props.SpacingBetweenLines(); sp != nil {
				paraHeight += float64(
					sp.Before(),
				) / 20.0
				paraHeight += float64(
					sp.After(),
				) / 20.0
			}
		}

		allLines = append(allLines, lines)
		totalHeight += paraHeight
	}

	// Add top and bottom padding
	totalHeight += 2 * padding

	return totalHeight, allLines, nil
}

// findCellIndexForColumn maps a grid column index to a cell index within a row.
// Returns -1 if the column index is not covered by any cell in the row.
//
// This function handles cells that span multiple columns (gridSpan).
// For example, if row has cells: [cell0(span=3), cell1(span=1), cell2(span=2)]
// - Grid columns 0, 1, 2 -> cell index 0
// - Grid column 3 -> cell index 1
// - Grid columns 4, 5 -> cell index 2
func findCellIndexForColumn(
	row *RowLayout,
	gridColIdx int,
) int {
	currentGridCol := 0

	for cellIdx, cell := range row.Cells {
		// Calculate the range of grid columns this cell covers
		cellStartCol := currentGridCol
		cellEndCol := currentGridCol + cell.ColSpan - 1

		// Check if gridColIdx falls within this cell's range
		if gridColIdx >= cellStartCol &&
			gridColIdx <= cellEndCol {
			return cellIdx
		}

		// Advance to next cell's starting column
		currentGridCol += cell.ColSpan
	}

	// Column index not found in this row
	return -1
}

// resolveVerticalMerges adjusts cell heights for vertically merged cells.
func (r *WordRenderer) resolveVerticalMerges(
	tl *TableLayout,
) {
	// Track vertical merge chains
	mergeStarts := make(
		map[int]int,
	) // col -> row index where merge starts

	for rowIdx, rowLayout := range tl.Rows {
		colIdx := 0
		for cellIdx := range rowLayout.Cells {
			cell := &tl.Rows[rowIdx].Cells[cellIdx]

			if cell.IsVMergeContinue {
				// This cell continues a merge, track it
				if startRow, ok := mergeStarts[colIdx]; ok {
					// Update the start cell's rowSpan
					tl.Rows[startRow].Cells[cellIdx].RowSpan++
				}
			} else {
				// Check if this starts a new merge
				props := cell.Cell.TableCellProperties()
				if props != nil {
					if vm := props.VerticalMerge(); vm != nil {
						if vm.Type() == elements.VerticalMergeRestart {
							mergeStarts[colIdx] = rowIdx
							cell.RowSpan = 1
						}
					}
				}
			}

			colIdx += cell.ColSpan
		}
	}

	// Calculate total heights for merged cells
	for colIdx, startRow := range mergeStarts {
		// Map grid column index to cell index
		// We need to find which cell in the row covers the grid column at colIdx
		cellIdx := findCellIndexForColumn(
			&tl.Rows[startRow],
			colIdx,
		)
		if cellIdx < 0 {
			// Column not found in row, skip
			continue
		}

		startCell := &tl.Rows[startRow].Cells[cellIdx]

		// Sum up heights of all spanned rows
		totalHeight := 0.0
		for i := 0; i < startCell.RowSpan && startRow+i < len(tl.Rows); i++ {
			totalHeight += tl.Rows[startRow+i].Height
		}
		startCell.Height = totalHeight
	}
}

// SplitTableAcrossPages splits a table layout if it doesn't fit on current page.
func (r *WordRenderer) SplitTableAcrossPages(
	tl *TableLayout,
	availableHeight float64,
) []*TableLayout {
	// Simple implementation: split by rows
	var splits []*TableLayout
	currentSplit := &TableLayout{
		Table:        tl.Table,
		X:            tl.X,
		Y:            tl.Y,
		ColumnWidths: tl.ColumnWidths,
	}

	currentHeight := 0.0

	for _, row := range tl.Rows {
		if currentHeight+row.Height > availableHeight &&
			len(currentSplit.Rows) > 0 {
			// Start a new split
			currentSplit.TotalWidth = tl.TotalWidth
			currentSplit.TotalHeight = currentHeight
			splits = append(splits, currentSplit)

			currentSplit = &TableLayout{
				Table:        tl.Table,
				X:            tl.X,
				Y:            tl.Y,
				ColumnWidths: tl.ColumnWidths,
			}
			currentHeight = 0
		}

		currentSplit.Rows = append(
			currentSplit.Rows,
			row,
		)
		currentHeight += row.Height
	}

	// Add final split
	if len(currentSplit.Rows) > 0 {
		currentSplit.TotalWidth = tl.TotalWidth
		currentSplit.TotalHeight = currentHeight
		splits = append(splits, currentSplit)
	}

	// If it all fits, return original
	if len(splits) <= 1 {
		return []*TableLayout{tl}
	}

	return splits
}
