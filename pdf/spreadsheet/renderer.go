// renderer.go provides the entry point for SpreadsheetML to PDF conversion.
// It coordinates worksheet parsing, layout, and rendering to produce a PDF file.

package spreadsheet

import (
	"fmt"
	"strconv"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice/spreadsheet"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// SpreadsheetRenderer handles the conversion of a SpreadsheetML document to PDF.
type SpreadsheetRenderer struct {
	// doc is the source spreadsheet document.
	doc *spreadsheet.Document

	// pdf is the target PDF document.
	pdf *core.Document

	// options contains rendering configuration.
	options RenderOptions

	// fontCache caches loaded fonts.
	fontCache *font.FontCache

	// numberFormats tracks custom number formats.
	numberFormats map[int]string

	// currentSheet tracks the current sheet being rendered.
	currentSheet *spreadsheet.Sheet

	// currentPageNumber tracks the current page number (1-based).
	currentPageNumber int
}

// RenderOptions configures spreadsheet rendering behavior.
type RenderOptions struct {
	// DefaultPageSize is used if the worksheet doesn't specify one.
	DefaultPageSize core.PageSize

	// EmbedFonts controls whether fonts are embedded in the PDF.
	EmbedFonts bool

	// DefaultFontFamily is the font family to use for cells without specific font.
	DefaultFontFamily string

	// DefaultFontSize is the font size to use for cells without specific size (in points).
	DefaultFontSize float64

	// GridLines controls whether to render gridlines.
	GridLines bool

	// SheetsToRender specifies which sheets to render. If nil/empty, render all sheets.
	SheetsToRender []string

	// RenderFormulas controls whether to show formulas instead of values.
	RenderFormulas bool
}

// DefaultRenderOptions returns the default rendering configuration.
func DefaultRenderOptions() RenderOptions {
	return RenderOptions{
		DefaultPageSize:   core.PageSizeLetter,
		EmbedFonts:        true,
		DefaultFontFamily: "Calibri",
		DefaultFontSize:   11.0,
		GridLines:         false,
		SheetsToRender:    nil,
		RenderFormulas:    false,
	}
}

// NewSpreadsheetRenderer creates a new SpreadsheetRenderer for the specified document.
func NewSpreadsheetRenderer(
	doc *spreadsheet.Document,
) (*SpreadsheetRenderer, error) {
	pdf, err := core.NewDocument()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create PDF document: %w",
			err,
		)
	}

	cache := font.NewFontCache(
		50,
	) // Cache up to 50 fonts

	return &SpreadsheetRenderer{
		doc:               doc,
		pdf:               pdf,
		options:           DefaultRenderOptions(),
		fontCache:         cache,
		numberFormats:     make(map[int]string),
		currentPageNumber: 0,
	}, nil
}

// WithOptions sets custom rendering options.
func (r *SpreadsheetRenderer) WithOptions(
	opts RenderOptions,
) *SpreadsheetRenderer {
	r.options = opts
	return r
}

// Render converts the spreadsheet document to a PDF and writes it to the specified path.
func (r *SpreadsheetRenderer) Render(
	outputPath string,
) error {
	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return fmt.Errorf("no workbook part")
	}

	// Load number formats from styles
	if err := r.loadNumberFormats(); err != nil {
		return fmt.Errorf(
			"failed to load number formats: %w",
			err,
		)
	}

	// Determine which sheets to render
	sheetsToRender := r.getSheetsToRender()
	if len(sheetsToRender) == 0 {
		// No sheets to render, create an empty page
		_, err := r.pdf.AddPage(
			r.options.DefaultPageSize,
		)
		if err != nil {
			return err
		}
	}

	// Render each sheet
	for _, sheet := range sheetsToRender {
		r.currentSheet = sheet
		if err := r.renderSheet(sheet); err != nil {
			return fmt.Errorf(
				"failed to render sheet %s: %w",
				sheet.Name(),
				err,
			)
		}
	}

	// Save the PDF
	return r.pdf.WriteToFile(outputPath)
}

// getSheetsToRender returns the sheets that should be rendered.
func (r *SpreadsheetRenderer) getSheetsToRender() []*spreadsheet.Sheet {
	var result []*spreadsheet.Sheet

	if len(r.options.SheetsToRender) == 0 {
		// Render all sheets
		for sheet := range r.doc.Sheets() {
			result = append(result, sheet)
		}
		return result
	}

	// Render only specified sheets
	for _, sheetName := range r.options.SheetsToRender {
		if sheet, found := r.doc.SheetByName(sheetName); found {
			result = append(result, sheet)
		}
	}

	return result
}

// loadNumberFormats loads custom number formats from the styles.
func (r *SpreadsheetRenderer) loadNumberFormats() error {
	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return nil
	}

	stylesPart := workbookPart.StylesPart()
	if stylesPart == nil {
		return nil
	}

	// TODO: Fix stylesheet access from WorkbookStylesPart
	// The Stylesheet() method returns an interface that can't be type asserted
	// Need to find the proper way to access the concrete Stylesheet element

	return nil
}

// renderSheet renders a single worksheet to PDF pages.
func (r *SpreadsheetRenderer) renderSheet(
	sheet *spreadsheet.Sheet,
) error {
	worksheet := sheet.Worksheet()
	if worksheet == nil {
		return nil
	}

	// Get page setup
	pageSetup := worksheet.PageSetup()
	printOptions := worksheet.PrintOptions()
	pageMargins := worksheet.PageMargins()

	// Determine page size and orientation
	pageSize := r.getPageSize(pageSetup)

	// Get margins
	margins := r.getMargins(pageMargins)

	// Get print area (if specified)
	printArea := r.getPrintArea(worksheet)

	// Calculate cell layout
	layout := r.calculateCellLayout(
		worksheet,
		printArea,
	)

	// Render pages
	return r.renderPages(
		worksheet,
		layout,
		pageSize,
		margins,
		pageSetup,
		printOptions,
	)
}

// getPageSize determines the page size from page setup.
func (r *SpreadsheetRenderer) getPageSize(
	pageSetup *elements.PageSetup,
) core.PageSize {
	size := r.options.DefaultPageSize

	if pageSetup != nil {
		// Convert paper size to dimensions
		paperSize := pageSetup.PaperSize()
		size = paperSizeToPageSize(paperSize)

		// Check orientation
		orientation := pageSetup.Orientation()
		if orientation == elements.OrientationLandscape {
			// Swap width and height for landscape
			size.Width, size.Height = size.Height, size.Width
		}
	}

	return size
}

// getMargins determines the margins from page margins element.
func (r *SpreadsheetRenderer) getMargins(
	pageMargins *elements.PageMargins,
) core.Margins {
	margins := core.Margins{
		Top:    72.0, // 1 inch default
		Bottom: 72.0,
		Left:   72.0,
		Right:  72.0,
	}

	if pageMargins != nil {
		// Excel margins are in inches
		margins.Top = pageMargins.Top() * 72.0
		margins.Bottom = pageMargins.Bottom() * 72.0
		margins.Left = pageMargins.Left() * 72.0
		margins.Right = pageMargins.Right() * 72.0
	}

	return margins
}

// getPrintArea extracts the print area from the worksheet.
func (r *SpreadsheetRenderer) getPrintArea(
	worksheet *elements.Worksheet,
) *CellRange {
	// Look for defined names in the workbook
	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return nil
	}

	wb := workbookPart.Workbook()
	if wb == nil {
		return nil
	}

	workbook, ok := wb.(*elements.Workbook)
	if !ok {
		return nil
	}

	definedNames := workbook.DefinedNames()
	if definedNames == nil {
		return nil
	}

	// Find print area for current sheet
	sheetName := r.currentSheet.Name()
	for definedName := range definedNames.DefinedNames() {
		if definedName.Name() == "_xlnm.Print_Area" {
			// Check if it's for this sheet
			localSheetId := definedName.LocalSheetId()
			if localSheetId >= 0 {
				// Parse the print area reference
				ref := definedName.Formula()
				cellRange := parsePrintAreaRef(
					ref,
					sheetName,
				)
				if cellRange != nil {
					return cellRange
				}
			}
		}
	}

	return nil
}

// PDF returns the underlying PDF document.
func (r *SpreadsheetRenderer) PDF() *core.Document {
	return r.pdf
}

// paperSizeToPageSize converts an Excel paper size code to PDF page dimensions.
func paperSizeToPageSize(
	paperSize int,
) core.PageSize {
	switch paperSize {
	case 1: // Letter
		return core.PageSizeLetter
	case 5: // Legal
		return core.PageSizeLegal
	case 8: // A3
		return core.PageSizeA3
	case 9: // A4
		return core.PageSizeA4
	case 11: // A5
		return core.PageSize{
			Width:  419.53,
			Height: 595.28,
		} // 148 x 210 mm
	case 3: // Tabloid
		return core.PageSizeTabloid
	default:
		return core.PageSizeLetter
	}
}

// parsePrintAreaRef parses a print area reference string.
// Format: 'SheetName'!$A$1:$Z$100
func parsePrintAreaRef(
	ref, sheetName string,
) *CellRange {
	if ref == "" {
		return nil
	}

	// Simple parsing - extract the range part after !
	parts := splitReference(ref)
	if len(parts) < 2 {
		return nil
	}

	// Parse the range (e.g., $A$1:$Z$100)
	rangeStr := parts[1]
	return parseRange(rangeStr)
}

// splitReference splits a reference like 'Sheet1'!A1:B2 into parts.
func splitReference(ref string) []string {
	idx := -1
	for i, ch := range ref {
		if ch == '!' {
			idx = i
			break
		}
	}

	if idx == -1 {
		return []string{ref}
	}

	return []string{ref[:idx], ref[idx+1:]}
}

// parseRange parses a range string like A1:B2.
func parseRange(rangeStr string) *CellRange {
	// Remove $ signs
	rangeStr = stripDollarSigns(rangeStr)

	// Split by :
	parts := []string{}
	colonIdx := -1
	for i, ch := range rangeStr {
		if ch == ':' {
			colonIdx = i
			break
		}
	}

	if colonIdx == -1 {
		// Single cell
		parts = []string{rangeStr, rangeStr}
	} else {
		parts = []string{rangeStr[:colonIdx], rangeStr[colonIdx+1:]}
	}

	if len(parts) != 2 {
		return nil
	}

	start := parseCellRef(parts[0])
	end := parseCellRef(parts[1])

	if start == nil || end == nil {
		return nil
	}

	return &CellRange{
		StartRow: start.Row,
		StartCol: start.Col,
		EndRow:   end.Row,
		EndCol:   end.Col,
	}
}

// stripDollarSigns removes $ characters from a string.
func stripDollarSigns(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '$' {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// parseCellRef parses a cell reference like A1.
func parseCellRef(ref string) *CellRef {
	if ref == "" {
		return nil
	}

	// Split into column letters and row number
	colEnd := 0
	for i, ch := range ref {
		if ch >= 'A' && ch <= 'Z' {
			colEnd = i + 1
		} else {
			break
		}
	}

	if colEnd == 0 {
		return nil
	}

	colStr := ref[:colEnd]
	rowStr := ref[colEnd:]

	col := columnLetterToNumber(colStr)
	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return nil
	}

	return &CellRef{
		Row: row,
		Col: col,
	}
}

// columnLetterToNumber converts column letters to a number (A=1, Z=26, AA=27).
func columnLetterToNumber(letters string) int {
	result := 0
	for _, ch := range letters {
		result = result*26 + int(ch-'A'+1)
	}
	return result
}

// CellRef represents a cell reference.
type CellRef struct {
	Row int
	Col int
}

// CellRange represents a rectangular range of cells.
type CellRange struct {
	StartRow int
	StartCol int
	EndRow   int
	EndCol   int
}
