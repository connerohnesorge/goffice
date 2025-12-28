// page_layout.go handles the rendering of worksheet content across multiple PDF pages.
// This includes page breaks, scaling, print titles, and headers/footers.

package spreadsheet

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// PageInfo contains information about a page being rendered.
type PageInfo struct {
	Number     int
	RowStart   int
	RowEnd     int
	ColStart   int
	ColEnd     int
	OffsetX    float64
	OffsetY    float64
	Scale      float64
	RepeatRows *CellRange
	RepeatCols *CellRange
}

// renderPages renders the worksheet content across multiple pages.
func (r *SpreadsheetRenderer) renderPages(
	worksheet *elements.Worksheet,
	layout *CellLayout,
	pageSize core.PageSize,
	margins core.Margins,
	pageSetup *elements.PageSetup,
	printOptions *elements.PrintOptions,
) error {
	// Get stylesheet for cell formatting
	stylesheet := r.getStylesheet()

	// If there's no visible range, create an empty page
	if layout.VisibleRange == nil {
		_, err := r.pdf.AddPage(pageSize)
		if err != nil {
			return err
		}
		r.currentPageNumber++
		return nil
	}

	// Calculate page breaks and layout
	pages := r.calculatePageLayout(
		layout,
		pageSize,
		margins,
		pageSetup,
	)

	// Check for scaling
	scale := 1.0
	if pageSetup != nil {
		scalePercent := pageSetup.Scale()
		if scalePercent > 0 &&
			scalePercent != 100 {
			scale = float64(scalePercent) / 100.0
		}

		// Check fit-to-page settings
		fitToWidth := pageSetup.FitToWidth()
		fitToHeight := pageSetup.FitToHeight()
		if (fitToWidth > 0 && fitToWidth != 1) ||
			(fitToHeight > 0 && fitToHeight != 1) {
			// Calculate scaling to fit
			scale = r.calculateFitToScale(
				layout,
				pageSize,
				margins,
				fitToWidth,
				fitToHeight,
			)
		}
	}

	// Get print titles (repeat rows/columns)
	repeatRows, repeatCols := r.getPrintTitles(
		worksheet,
	)

	// Render each page
	for _, pageInfo := range pages {
		pageInfo.Scale = scale
		pageInfo.RepeatRows = repeatRows
		pageInfo.RepeatCols = repeatCols

		if err := r.renderPage(worksheet, layout, pageInfo, pageSize, margins, pageSetup, stylesheet); err != nil {
			return err
		}
	}

	return nil
}

// calculatePageLayout calculates how the worksheet should be split across pages.
func (r *SpreadsheetRenderer) calculatePageLayout(
	layout *CellLayout,
	pageSize core.PageSize,
	margins core.Margins,
	pageSetup *elements.PageSetup,
) []*PageInfo {
	if layout.VisibleRange == nil {
		return []*PageInfo{}
	}

	// Available space on page
	availWidth := pageSize.Width - margins.Left - margins.Right
	availHeight := pageSize.Height - margins.Top - margins.Bottom

	// Check for manual page breaks
	// For simplicity, we'll use automatic pagination based on available space
	pages := []*PageInfo{}

	// Calculate page breaks
	currentRow := layout.VisibleRange.StartRow
	currentCol := layout.VisibleRange.StartCol
	pageNum := 1

	for currentRow <= layout.VisibleRange.EndRow {
		// Determine how many rows fit on this page
		rowEnd := currentRow
		heightUsed := 0.0

		for rowEnd <= layout.VisibleRange.EndRow {
			if layout.HiddenRows[rowEnd] {
				rowEnd++
				continue
			}

			rowHeight, found := layout.RowHeights[rowEnd]
			if !found {
				rowHeight = layout.DefaultRowHeight
			}

			if heightUsed+rowHeight > availHeight &&
				heightUsed > 0 {
				break
			}

			heightUsed += rowHeight
			rowEnd++
		}

		// Determine how many columns fit on this page
		colEnd := currentCol
		widthUsed := 0.0

		for colEnd <= layout.VisibleRange.EndCol {
			if layout.HiddenColumns[colEnd] {
				colEnd++
				continue
			}

			colWidth, found := layout.ColumnWidths[colEnd]
			if !found {
				colWidth = layout.DefaultColumnWidth
			}

			if widthUsed+colWidth > availWidth &&
				widthUsed > 0 {
				break
			}

			widthUsed += colWidth
			colEnd++
		}

		// Create page info
		pages = append(pages, &PageInfo{
			Number:   pageNum,
			RowStart: currentRow,
			RowEnd:   rowEnd - 1,
			ColStart: currentCol,
			ColEnd:   colEnd - 1,
		})

		pageNum++

		// Move to next page
		if colEnd > layout.VisibleRange.EndCol {
			// Move to next row range
			currentRow = rowEnd
			currentCol = layout.VisibleRange.StartCol
		} else {
			// Move to next column range in same rows
			currentCol = colEnd
		}
	}

	return pages
}

// calculateFitToScale calculates the scaling factor to fit content to pages.
func (r *SpreadsheetRenderer) calculateFitToScale(
	layout *CellLayout,
	pageSize core.PageSize,
	margins core.Margins,
	fitToWidth, fitToHeight int,
) float64 {
	if layout.VisibleRange == nil {
		return 1.0
	}

	availWidth := pageSize.Width - margins.Left - margins.Right
	availHeight := pageSize.Height - margins.Top - margins.Bottom

	scaleX := 1.0
	scaleY := 1.0

	if fitToWidth > 0 {
		totalWidth := r.calculateTotalWidth(
			layout,
			layout.VisibleRange.StartCol,
			layout.VisibleRange.EndCol,
		)
		scaleX = (availWidth * float64(fitToWidth)) / totalWidth
	}

	if fitToHeight > 0 {
		totalHeight := r.calculateTotalHeight(
			layout,
			layout.VisibleRange.StartRow,
			layout.VisibleRange.EndRow,
		)
		scaleY = (availHeight * float64(fitToHeight)) / totalHeight
	}

	// Use the smaller scale to ensure content fits
	scale := scaleX
	if fitToHeight > 0 && scaleY < scale {
		scale = scaleY
	}

	// Limit scaling to reasonable range
	if scale < 0.1 {
		scale = 0.1
	}
	if scale > 1.0 {
		scale = 1.0
	}

	return scale
}

// getPrintTitles extracts the print titles (repeat rows/columns) from the worksheet.
func (r *SpreadsheetRenderer) getPrintTitles(
	worksheet *elements.Worksheet,
) (*CellRange, *CellRange) {
	// Look for defined names in the workbook
	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return nil, nil
	}

	wb := workbookPart.Workbook()
	if wb == nil {
		return nil, nil
	}

	workbook, ok := wb.(*elements.Workbook)
	if !ok {
		return nil, nil
	}

	definedNames := workbook.DefinedNames()
	if definedNames == nil {
		return nil, nil
	}

	var repeatRows, repeatCols *CellRange

	// Find print titles for current sheet
	sheetName := r.currentSheet.Name()
	for definedName := range definedNames.DefinedNames() {
		if definedName.Name() == "_xlnm.Print_Titles" {
			// Parse the print titles reference
			ref := definedName.Formula()
			repeatRows, repeatCols = parsePrintTitlesRef(
				ref,
				sheetName,
			)
			break
		}
	}

	return repeatRows, repeatCols
}

// parsePrintTitlesRef parses a print titles reference.
// Format: 'SheetName'!$1:$3,'SheetName'!$A:$B
func parsePrintTitlesRef(
	ref, sheetName string,
) (*CellRange, *CellRange) {
	// This is a simplified parser
	// In practice, print titles can be quite complex
	return nil, nil
}

// renderPage renders a single page of the worksheet.
func (r *SpreadsheetRenderer) renderPage(
	worksheet *elements.Worksheet,
	layout *CellLayout,
	pageInfo *PageInfo,
	pageSize core.PageSize,
	margins core.Margins,
	pageSetup *elements.PageSetup,
	stylesheet *elements.Stylesheet,
) error {
	// Create a new page
	page, err := r.pdf.AddPage(pageSize)
	if err != nil {
		return err
	}

	r.currentPageNumber++

	// Note: PDF coordinate system transformations would go here
	// For now, we'll apply margins and scaling directly to coordinates

	// Render print titles (repeat rows/columns on every page)
	repeatRowsHeight := 0.0
	repeatColsWidth := 0.0

	if pageInfo.RepeatRows != nil {
		// Render repeat rows at top of the page
		repeatRowsHeight = r.renderRepeatRows(
			page,
			worksheet,
			layout,
			pageInfo,
			margins,
			stylesheet,
		)
	}

	if pageInfo.RepeatCols != nil {
		// Render repeat columns at left of the page
		repeatColsWidth = r.renderRepeatColumns(
			page,
			worksheet,
			layout,
			pageInfo,
			margins,
			repeatRowsHeight,
			stylesheet,
		)
	}

	// Render cells
	if sheetData := worksheet.SheetData(); sheetData != nil {
		for row := range sheetData.Rows() {
			rowIdx := int(row.RowIndex())

			// Check if row is in this page's range
			if rowIdx < pageInfo.RowStart ||
				rowIdx > pageInfo.RowEnd {
				continue
			}

			// Skip rows that are in the repeat rows range (already rendered)
			if pageInfo.RepeatRows != nil &&
				rowIdx >= pageInfo.RepeatRows.StartRow &&
				rowIdx <= pageInfo.RepeatRows.EndRow {
				continue
			}

			// Check if row is hidden
			if layout.HiddenRows[rowIdx] {
				continue
			}

			// Render cells in this row
			for cell := range row.Cells() {
				ref := cell.Reference()
				cellRef := parseCellRef(ref)
				if cellRef == nil {
					continue
				}

				// Check if column is in this page's range
				if cellRef.Col < pageInfo.ColStart ||
					cellRef.Col > pageInfo.ColEnd {
					continue
				}

				// Skip columns that are in the repeat columns range (already rendered)
				if pageInfo.RepeatCols != nil &&
					cellRef.Col >= pageInfo.RepeatCols.StartCol &&
					cellRef.Col <= pageInfo.RepeatCols.EndCol {
					continue
				}

				// Check if column is hidden
				if layout.HiddenColumns[cellRef.Col] {
					continue
				}

				// Get cell position and size
				x, y := r.getCellPosition(
					layout,
					cellRef.Row,
					cellRef.Col,
				)
				width, height := r.getCellSize(
					layout,
					cellRef.Row,
					cellRef.Col,
				)

				// Adjust for page offset and repeat rows/columns
				x -= pageInfo.OffsetX
				y -= pageInfo.OffsetY
				x += margins.Left + repeatColsWidth
				y += margins.Top + repeatRowsHeight

				// Render the cell
				if err := r.renderCell(page, cell, layout, x, y, width, height, stylesheet); err != nil {
					return err
				}
			}
		}
	}

	// Render gridlines if enabled
	if r.options.GridLines {
		r.renderGridlines(page, layout, pageInfo)
	}

	// Render headers and footers
	if err := r.renderHeaderFooter(page, worksheet, pageInfo, pageSize, margins); err != nil {
		return err
	}

	return nil
}

// renderGridlines renders gridlines for the visible cells.
func (r *SpreadsheetRenderer) renderGridlines(
	page *core.Page,
	layout *CellLayout,
	pageInfo *PageInfo,
) {
	// Set gridline color (light gray)
	gridColor := core.RGB{R: 0.9, G: 0.9, B: 0.9}

	// Calculate total height for vertical gridlines
	totalHeight := 0.0
	for row := pageInfo.RowStart; row <= pageInfo.RowEnd; row++ {
		if layout.HiddenRows[row] {
			continue
		}
		rowHeight, found := layout.RowHeights[row]
		if !found {
			rowHeight = layout.DefaultRowHeight
		}
		totalHeight += rowHeight
	}

	// Draw vertical gridlines
	x := 0.0
	for col := pageInfo.ColStart; col <= pageInfo.ColEnd; col++ {
		if layout.HiddenColumns[col] {
			continue
		}

		colWidth, found := layout.ColumnWidths[col]
		if !found {
			colWidth = layout.DefaultColumnWidth
		}

		// Draw line
		drawLine(
			page,
			x,
			0,
			x,
			totalHeight,
			gridColor,
			0.5,
		)

		x += colWidth
	}

	// Draw final vertical line
	drawLine(
		page,
		x,
		0,
		x,
		totalHeight,
		gridColor,
		0.5,
	)

	// Draw horizontal gridlines
	y := 0.0
	for row := pageInfo.RowStart; row <= pageInfo.RowEnd; row++ {
		if layout.HiddenRows[row] {
			continue
		}

		rowHeight, found := layout.RowHeights[row]
		if !found {
			rowHeight = layout.DefaultRowHeight
		}

		// Draw line
		drawLine(page, 0, y, x, y, gridColor, 0.5)

		y += rowHeight
	}

	// Draw final horizontal line
	drawLine(page, 0, y, x, y, gridColor, 0.5)
}

// renderHeaderFooter renders the header and footer for a page.
func (r *SpreadsheetRenderer) renderHeaderFooter(
	page *core.Page,
	worksheet *elements.Worksheet,
	pageInfo *PageInfo,
	pageSize core.PageSize,
	margins core.Margins,
) error {
	headerFooter := worksheet.HeaderFooter()
	if headerFooter == nil {
		return nil
	}

	// Get header and footer content
	oddHeader := headerFooter.OddHeader()
	oddFooter := headerFooter.OddFooter()

	// Render header
	if oddHeader != "" {
		if err := r.renderHeaderFooterContent(page, oddHeader, pageInfo, pageSize, margins, true); err != nil {
			return err
		}
	}

	// Render footer
	if oddFooter != "" {
		if err := r.renderHeaderFooterContent(page, oddFooter, pageInfo, pageSize, margins, false); err != nil {
			return err
		}
	}

	return nil
}

// renderHeaderFooterContent renders header or footer content.
func (r *SpreadsheetRenderer) renderHeaderFooterContent(
	page *core.Page,
	content string,
	pageInfo *PageInfo,
	pageSize core.PageSize,
	margins core.Margins,
	isHeader bool,
) error {
	if content == "" {
		return nil
	}

	// Parse header/footer codes
	// Excel header/footer format: &L left, &C center, &R right
	// &P page number, &N total pages, &D date, &T time, etc.
	left, center, right := parseHeaderFooterContent(
		content,
		pageInfo,
	)

	// Determine Y position
	y := margins.Top / 2
	if !isHeader {
		y = pageSize.Height - margins.Bottom/2
	}

	// Load default font
	fontObj, err := r.loadFont(&CellFont{
		Family: r.options.DefaultFontFamily,
		Size:   10.0,
		Color:  core.RGB{R: 0, G: 0, B: 0},
	})
	if err != nil {
		return err
	}

	textColor := core.RGB{R: 0, G: 0, B: 0}

	// Render left section
	if left != "" {
		_ = drawText(
			r,
			page,
			fontObj,
			left,
			margins.Left,
			y,
			10.0,
			textColor,
		)
	}

	// Render center section
	if center != "" {
		// Approximate text width
		textWidth := float64(
			len(center),
		) * 10.0 * 0.5
		x := (pageSize.Width - textWidth) / 2
		_ = drawText(
			r,
			page,
			fontObj,
			center,
			x,
			y,
			10.0,
			textColor,
		)
	}

	// Render right section
	if right != "" {
		// Approximate text width
		textWidth := float64(
			len(right),
		) * 10.0 * 0.5
		x := pageSize.Width - margins.Right - textWidth
		_ = drawText(
			r,
			page,
			fontObj,
			right,
			x,
			y,
			10.0,
			textColor,
		)
	}

	return nil
}

// parseHeaderFooterContent parses header/footer content and returns left, center, right sections.
func parseHeaderFooterContent(
	content string,
	pageInfo *PageInfo,
) (string, string, string) {
	// Simplified parser for header/footer codes
	// Real implementation would handle all Excel codes

	left := ""
	center := ""
	right := ""

	currentSection := &center
	i := 0
	for i < len(content) {
		if content[i] == '&' &&
			i+1 < len(content) {
			code := content[i+1]
			switch code {
			case 'L':
				currentSection = &left
				i += 2
				continue
			case 'C':
				currentSection = &center
				i += 2
				continue
			case 'R':
				currentSection = &right
				i += 2
				continue
			case 'P':
				*currentSection += fmt.Sprintf(
					"%d",
					pageInfo.Number,
				)
				i += 2
				continue
			case 'D':
				*currentSection += "01/01/2024"
				i += 2
				continue
			case 'T':
				*currentSection += "12:00:00"
				i += 2
				continue
			}
		}

		*currentSection += string(content[i])
		i++
	}

	return left, center, right
}

// getStylesheet gets the stylesheet from the workbook.
func (r *SpreadsheetRenderer) getStylesheet() *elements.Stylesheet {
	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return nil
	}

	stylesPart := workbookPart.StylesPart()
	if stylesPart == nil {
		return nil
	}

	// Get the stylesheet from the styles part
	return stylesPart.Stylesheet()
}

// renderRepeatRows renders the repeat rows at the top of the page.
// Returns the total height of the repeated rows.
func (r *SpreadsheetRenderer) renderRepeatRows(
	page *core.Page,
	worksheet *elements.Worksheet,
	layout *CellLayout,
	pageInfo *PageInfo,
	margins core.Margins,
	stylesheet *elements.Stylesheet,
) float64 {
	if pageInfo.RepeatRows == nil {
		return 0.0
	}

	totalHeight := 0.0
	currentY := margins.Top

	// Get sheet data
	sheetData := worksheet.SheetData()
	if sheetData == nil {
		return 0.0
	}

	// Iterate through rows in the repeat range
	for row := range sheetData.Rows() {
		rowIdx := int(row.RowIndex())

		// Check if row is in the repeat range
		if rowIdx < pageInfo.RepeatRows.StartRow ||
			rowIdx > pageInfo.RepeatRows.EndRow {
			continue
		}

		// Check if row is hidden
		if layout.HiddenRows[rowIdx] {
			continue
		}

		// Get row height
		rowHeight, found := layout.RowHeights[rowIdx]
		if !found {
			rowHeight = layout.DefaultRowHeight
		}

		// Render cells in this row
		for cell := range row.Cells() {
			ref := cell.Reference()
			cellRef := parseCellRef(ref)
			if cellRef == nil {
				continue
			}

			// Check if column is in the page's range (or repeat cols range)
			if cellRef.Col < pageInfo.ColStart ||
				cellRef.Col > pageInfo.ColEnd {
				// Skip columns outside the current page range
				continue
			}

			// Check if column is hidden
			if layout.HiddenColumns[cellRef.Col] {
				continue
			}

			// Calculate cell position
			x := margins.Left
			for c := pageInfo.ColStart; c < cellRef.Col; c++ {
				if layout.HiddenColumns[c] {
					continue
				}
				colWidth, found := layout.ColumnWidths[c]
				if !found {
					colWidth = layout.DefaultColumnWidth
				}
				x += colWidth
			}

			width, height := r.getCellSize(
				layout,
				cellRef.Row,
				cellRef.Col,
			)

			// Render the cell
			if err := r.renderCell(page, cell, layout, x, currentY, width, height, stylesheet); err != nil {
				// Log error but continue rendering
				continue
			}
		}

		totalHeight += rowHeight
		currentY += rowHeight
	}

	return totalHeight
}

// renderRepeatColumns renders the repeat columns at the left of the page.
// Returns the total width of the repeated columns.
func (r *SpreadsheetRenderer) renderRepeatColumns(
	page *core.Page,
	worksheet *elements.Worksheet,
	layout *CellLayout,
	pageInfo *PageInfo,
	margins core.Margins,
	repeatRowsHeight float64,
	stylesheet *elements.Stylesheet,
) float64 {
	if pageInfo.RepeatCols == nil {
		return 0.0
	}

	totalWidth := 0.0

	// Get sheet data
	sheetData := worksheet.SheetData()
	if sheetData == nil {
		return 0.0
	}

	// Calculate total width of repeat columns
	for col := pageInfo.RepeatCols.StartCol; col <= pageInfo.RepeatCols.EndCol; col++ {
		if layout.HiddenColumns[col] {
			continue
		}
		colWidth, found := layout.ColumnWidths[col]
		if !found {
			colWidth = layout.DefaultColumnWidth
		}
		totalWidth += colWidth
	}

	// Iterate through rows in the page's range
	for row := range sheetData.Rows() {
		rowIdx := int(row.RowIndex())

		// Check if row is in this page's range
		if rowIdx < pageInfo.RowStart ||
			rowIdx > pageInfo.RowEnd {
			continue
		}

		// Check if row is hidden
		if layout.HiddenRows[rowIdx] {
			continue
		}

		// Calculate Y position for this row
		currentY := margins.Top + repeatRowsHeight
		for r := pageInfo.RowStart; r < rowIdx; r++ {
			if layout.HiddenRows[r] {
				continue
			}
			rh, found := layout.RowHeights[r]
			if !found {
				rh = layout.DefaultRowHeight
			}
			currentY += rh
		}

		// Render cells in the repeat columns for this row
		currentX := margins.Left
		for col := pageInfo.RepeatCols.StartCol; col <= pageInfo.RepeatCols.EndCol; col++ {
			if layout.HiddenColumns[col] {
				continue
			}

			// Find the cell at this position
			for cell := range row.Cells() {
				ref := cell.Reference()
				cellRef := parseCellRef(ref)
				if cellRef == nil {
					continue
				}

				if cellRef.Col != col {
					continue
				}

				width, height := r.getCellSize(
					layout,
					cellRef.Row,
					cellRef.Col,
				)

				// Render the cell
				if err := r.renderCell(page, cell, layout, currentX, currentY, width, height, stylesheet); err != nil {
					// Log error but continue rendering
					continue
				}
				break
			}

			// Move to next column position
			colWidth, found := layout.ColumnWidths[col]
			if !found {
				colWidth = layout.DefaultColumnWidth
			}
			currentX += colWidth
		}
	}

	return totalWidth
}
