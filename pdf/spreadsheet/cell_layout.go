// cell_layout.go handles the calculation of cell positions and sizes.
// This module converts Excel's column widths (in character units) and row heights
// (in points) to PDF coordinates.

package spreadsheet

import (
	"math"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// CellLayout contains the calculated layout information for all cells in a worksheet.
type CellLayout struct {
	// ColumnWidths maps column index to width in points.
	ColumnWidths map[int]float64

	// RowHeights maps row index to height in points.
	RowHeights map[int]float64

	// DefaultColumnWidth is the default column width in points.
	DefaultColumnWidth float64

	// DefaultRowHeight is the default row height in points.
	DefaultRowHeight float64

	// TotalWidth is the total width of all columns in the layout.
	TotalWidth float64

	// TotalHeight is the total height of all rows in the layout.
	TotalHeight float64

	// VisibleRange is the range of cells to render.
	VisibleRange *CellRange

	// MergedCells contains all merged cell regions.
	MergedCells []*CellRange

	// HiddenRows tracks which rows are hidden.
	HiddenRows map[int]bool

	// HiddenColumns tracks which columns are hidden.
	HiddenColumns map[int]bool
}

// NewCellLayout creates a new cell layout.
func NewCellLayout() *CellLayout {
	return &CellLayout{
		ColumnWidths:       make(map[int]float64),
		RowHeights:         make(map[int]float64),
		DefaultColumnWidth: 64.0, // Default ~8.43 characters * 7.6 points/char
		DefaultRowHeight:   15.0, // Default 15 points
		MergedCells:        []*CellRange{},
		HiddenRows:         make(map[int]bool),
		HiddenColumns:      make(map[int]bool),
	}
}

// calculateCellLayout computes the layout for the worksheet.
func (r *SpreadsheetRenderer) calculateCellLayout(
	worksheet *elements.Worksheet,
	printArea *CellRange,
) *CellLayout {
	layout := NewCellLayout()

	// Get default column width from sheet format
	if sheetFormatPr := worksheet.SheetFormatPr(); sheetFormatPr != nil {
		// Excel default column width is in character units
		// 1 character unit ≈ 7.6 points (for default font)
		defaultWidth := sheetFormatPr.DefaultColWidth()
		if defaultWidth > 0 {
			layout.DefaultColumnWidth = defaultWidth * 7.6
		}

		// Default row height is in points
		defaultHeight := sheetFormatPr.DefaultRowHeight()
		if defaultHeight > 0 {
			layout.DefaultRowHeight = defaultHeight
		}
	}

	// Process column definitions
	if cols := worksheet.Cols(); cols != nil {
		for col := range cols.Cols() {
			min := col.Min()
			max := col.Max()
			width := col.Width()
			hidden := col.Hidden()

			// Convert Excel column width to points
			// Excel width is in character units: width = (characters * maxDigitWidth + padding) / maxDigitWidth
			// For default font (Calibri 11pt), 1 character ≈ 7.6 points
			widthPoints := width * 7.6

			for i := min; i <= max; i++ {
				layout.ColumnWidths[i] = widthPoints
				if hidden {
					layout.HiddenColumns[i] = true
				}
			}
		}
	}

	// Process row definitions and determine visible range
	minRow := math.MaxInt32
	maxRow := 0
	minCol := math.MaxInt32
	maxCol := 0

	if sheetData := worksheet.SheetData(); sheetData != nil {
		for row := range sheetData.Rows() {
			rowIdx := int(row.RowIndex())

			// Track visible range
			if rowIdx < minRow {
				minRow = rowIdx
			}
			if rowIdx > maxRow {
				maxRow = rowIdx
			}

			// Store row height
			if row.CustomHeight() {
				layout.RowHeights[rowIdx] = row.Height()
			}

			// Check if row is hidden
			if row.Hidden() {
				layout.HiddenRows[rowIdx] = true
			}

			// Process cells to determine column range
			for cell := range row.Cells() {
				ref := cell.Reference()
				cellRef := parseCellRef(ref)
				if cellRef != nil {
					if cellRef.Col < minCol {
						minCol = cellRef.Col
					}
					if cellRef.Col > maxCol {
						maxCol = cellRef.Col
					}
				}
			}
		}
	}

	// If print area is specified, use it to constrain visible range
	if printArea != nil {
		layout.VisibleRange = printArea
	} else if minRow != math.MaxInt32 {
		layout.VisibleRange = &CellRange{
			StartRow: minRow,
			StartCol: minCol,
			EndRow:   maxRow,
			EndCol:   maxCol,
		}
	}

	// Process merged cells
	if mergeCells := worksheet.MergeCells(); mergeCells != nil {
		for mergeCell := range mergeCells.GetMergeCells() {
			ref := mergeCell.Ref()
			cellRange := parseRange(ref)
			if cellRange != nil {
				layout.MergedCells = append(
					layout.MergedCells,
					cellRange,
				)
			}
		}
	}

	// Calculate total dimensions
	if layout.VisibleRange != nil {
		layout.TotalWidth = r.calculateTotalWidth(
			layout,
			layout.VisibleRange.StartCol,
			layout.VisibleRange.EndCol,
		)
		layout.TotalHeight = r.calculateTotalHeight(
			layout,
			layout.VisibleRange.StartRow,
			layout.VisibleRange.EndRow,
		)
	}

	return layout
}

// calculateTotalWidth computes the total width for a column range.
func (r *SpreadsheetRenderer) calculateTotalWidth(
	layout *CellLayout,
	startCol, endCol int,
) float64 {
	total := 0.0
	for col := startCol; col <= endCol; col++ {
		if layout.HiddenColumns[col] {
			continue
		}

		width, found := layout.ColumnWidths[col]
		if !found {
			width = layout.DefaultColumnWidth
		}
		total += width
	}

	return total
}

// calculateTotalHeight computes the total height for a row range.
func (r *SpreadsheetRenderer) calculateTotalHeight(
	layout *CellLayout,
	startRow, endRow int,
) float64 {
	total := 0.0
	for row := startRow; row <= endRow; row++ {
		if layout.HiddenRows[row] {
			continue
		}

		height, found := layout.RowHeights[row]
		if !found {
			height = layout.DefaultRowHeight
		}
		total += height
	}

	return total
}

// getCellPosition returns the position of a cell in points.
func (r *SpreadsheetRenderer) getCellPosition(
	layout *CellLayout,
	row, col int,
) (x, y float64) {
	if layout.VisibleRange == nil {
		return 0, 0
	}

	// Calculate X position (sum of column widths before this column)
	x = 0.0
	for c := layout.VisibleRange.StartCol; c < col; c++ {
		if layout.HiddenColumns[c] {
			continue
		}
		width, found := layout.ColumnWidths[c]
		if !found {
			width = layout.DefaultColumnWidth
		}
		x += width
	}

	// Calculate Y position (sum of row heights before this row)
	y = 0.0
	for r := layout.VisibleRange.StartRow; r < row; r++ {
		if layout.HiddenRows[r] {
			continue
		}
		height, found := layout.RowHeights[r]
		if !found {
			height = layout.DefaultRowHeight
		}
		y += height
	}

	return x, y
}

// getCellSize returns the size of a cell in points.
func (r *SpreadsheetRenderer) getCellSize(
	layout *CellLayout,
	row, col int,
) (width, height float64) {
	// Get column width
	width, found := layout.ColumnWidths[col]
	if !found {
		width = layout.DefaultColumnWidth
	}

	// Get row height
	height, found = layout.RowHeights[row]
	if !found {
		height = layout.DefaultRowHeight
	}

	// Check if this cell is part of a merged region
	for _, merged := range layout.MergedCells {
		if row >= merged.StartRow &&
			row <= merged.EndRow &&
			col >= merged.StartCol &&
			col <= merged.EndCol {
			// Cell is part of a merged region
			// Only calculate size for the top-left cell of the merged region
			if row == merged.StartRow &&
				col == merged.StartCol {
				// Calculate merged width
				width = 0.0
				for c := merged.StartCol; c <= merged.EndCol; c++ {
					if layout.HiddenColumns[c] {
						continue
					}
					w, found := layout.ColumnWidths[c]
					if !found {
						w = layout.DefaultColumnWidth
					}
					width += w
				}

				// Calculate merged height
				height = 0.0
				for r := merged.StartRow; r <= merged.EndRow; r++ {
					if layout.HiddenRows[r] {
						continue
					}
					h, found := layout.RowHeights[r]
					if !found {
						h = layout.DefaultRowHeight
					}
					height += h
				}
			} else {
				// Not the top-left cell, should not be rendered
				width = 0
				height = 0
			}

			break
		}
	}

	return width, height
}

// isCellMerged checks if a cell is part of a merged region.
func (r *SpreadsheetRenderer) isCellMerged(
	layout *CellLayout,
	row, col int,
) bool {
	for _, merged := range layout.MergedCells {
		if row >= merged.StartRow &&
			row <= merged.EndRow &&
			col >= merged.StartCol &&
			col <= merged.EndCol {
			return true
		}
	}

	return false
}

// isTopLeftOfMerge checks if a cell is the top-left cell of a merged region.
func (r *SpreadsheetRenderer) isTopLeftOfMerge(
	layout *CellLayout,
	row, col int,
) bool {
	for _, merged := range layout.MergedCells {
		if row == merged.StartRow &&
			col == merged.StartCol {
			return true
		}
	}

	return false
}
