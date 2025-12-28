// cell_renderer.go handles rendering of individual cells with their content and formatting.

package spreadsheet

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// CellStyle contains the computed style for a cell.
type CellStyle struct {
	Font      *CellFont
	Fill      *CellFill
	Border    *CellBorder
	Alignment *CellAlignment
	NumFmt    string
}

// CellFont contains font properties for a cell.
type CellFont struct {
	Family string
	Size   float64
	Bold   bool
	Italic bool
	Color  core.RGB
}

// CellFill contains fill properties for a cell.
type CellFill struct {
	Type        string // "solid", "pattern", "gradient"
	FgColor     core.RGB
	BgColor     core.RGB
	PatternType string
}

// CellBorder contains border properties for a cell.
type CellBorder struct {
	Left   *BorderSide
	Right  *BorderSide
	Top    *BorderSide
	Bottom *BorderSide
}

// BorderSide contains properties for one side of a border.
type BorderSide struct {
	Style string
	Color core.RGB
	Width float64
}

// CellAlignment contains alignment properties for a cell.
type CellAlignment struct {
	Horizontal string // "left", "center", "right", "justify"
	Vertical   string // "top", "center", "bottom"
	WrapText   bool
	Rotation   int
	Indent     int
}

// renderCell renders a single cell to the PDF page.
func (r *SpreadsheetRenderer) renderCell(
	page *core.Page,
	cell *elements.Cell,
	layout *CellLayout,
	x, y, width, height float64,
	stylesheet *elements.Stylesheet,
) error {
	if width == 0 || height == 0 {
		// Cell is hidden or part of merged region (not top-left)
		return nil
	}

	// Get cell style
	style := r.getCellStyle(cell, stylesheet)

	// Render cell background
	if style.Fill != nil {
		r.renderCellFill(
			page,
			x,
			y,
			width,
			height,
			style.Fill,
		)
	}

	// Render cell borders
	if style.Border != nil {
		r.renderCellBorder(
			page,
			x,
			y,
			width,
			height,
			style.Border,
		)
	}

	// Render cell content
	content := r.getCellContent(cell, style)
	if content != "" {
		r.renderCellText(
			page,
			x,
			y,
			width,
			height,
			content,
			style,
		)
	}

	return nil
}

// getCellStyle computes the style for a cell.
func (r *SpreadsheetRenderer) getCellStyle(
	cell *elements.Cell,
	stylesheet *elements.Stylesheet,
) *CellStyle {
	style := &CellStyle{
		Font: &CellFont{
			Family: r.options.DefaultFontFamily,
			Size:   r.options.DefaultFontSize,
			Color:  core.RGB{R: 0, G: 0, B: 0},
		},
		Alignment: &CellAlignment{
			Horizontal: "left",
			Vertical:   "bottom",
		},
	}

	if stylesheet == nil {
		return style
	}

	// Get cell style index
	styleIdx := cell.StyleIndex()
	if styleIdx == 0 {
		return style
	}

	// Get cell XF (format)
	cellXfs := stylesheet.CellXfs()
	if cellXfs == nil {
		return style
	}

	var cellXf *elements.Xf
	idx := 0
	for xf := range cellXfs.Xfs() {
		if idx == int(styleIdx) {
			cellXf = xf
			break
		}
		idx++
	}

	if cellXf == nil {
		return style
	}

	// Apply font
	if cellXf.ApplyFont() {
		fontId := cellXf.FontId()
		if fontElem := r.getFont(stylesheet, int(fontId)); fontElem != nil {
			style.Font = r.parseFontElement(
				fontElem,
			)
		}
	}

	// Apply fill
	if cellXf.ApplyFill() {
		fillId := cellXf.FillId()
		if fillElem := r.getFill(stylesheet, int(fillId)); fillElem != nil {
			style.Fill = r.parseFillElement(
				fillElem,
			)
		}
	}

	// Apply border
	if cellXf.ApplyBorder() {
		borderId := cellXf.BorderId()
		if borderElem := r.getBorder(stylesheet, int(borderId)); borderElem != nil {
			style.Border = r.parseBorderElement(
				borderElem,
			)
		}
	}

	// Apply alignment
	if cellXf.ApplyAlignment() {
		if alignment := cellXf.Alignment(); alignment != nil {
			style.Alignment = r.parseAlignmentElement(
				alignment,
			)
		}
	}

	// Apply number format
	if cellXf.ApplyNumberFormat() {
		numFmtId := cellXf.NumFmtId()
		style.NumFmt = r.getNumberFormat(
			int(numFmtId),
		)
	}

	return style
}

// getFont retrieves a font from the stylesheet.
func (r *SpreadsheetRenderer) getFont(
	stylesheet *elements.Stylesheet,
	fontId int,
) *elements.Font {
	fonts := stylesheet.Fonts()
	if fonts == nil {
		return nil
	}

	idx := 0
	for font := range fonts.Fonts() {
		if idx == fontId {
			return font
		}
		idx++
	}

	return nil
}

// getFill retrieves a fill from the stylesheet.
func (r *SpreadsheetRenderer) getFill(
	stylesheet *elements.Stylesheet,
	fillId int,
) *elements.Fill {
	fills := stylesheet.Fills()
	if fills == nil {
		return nil
	}

	idx := 0
	for fill := range fills.Fills() {
		if idx == fillId {
			return fill
		}
		idx++
	}

	return nil
}

// getBorder retrieves a border from the stylesheet.
func (r *SpreadsheetRenderer) getBorder(
	stylesheet *elements.Stylesheet,
	borderId int,
) *elements.Border {
	borders := stylesheet.Borders()
	if borders == nil {
		return nil
	}

	idx := 0
	for border := range borders.Borders() {
		if idx == borderId {
			return border
		}
		idx++
	}

	return nil
}

// parseFontElement parses a font element into CellFont.
func (r *SpreadsheetRenderer) parseFontElement(
	fontElem *elements.Font,
) *CellFont {
	font := &CellFont{
		Family: r.options.DefaultFontFamily,
		Size:   r.options.DefaultFontSize,
		Color:  core.RGB{R: 0, G: 0, B: 0},
	}

	if sz := fontElem.Size(); sz != nil {
		font.Size = sz.Val()
	}

	if name := fontElem.Name(); name != nil {
		font.Family = name.Val()
	}

	if b := fontElem.Bold(); b != nil {
		font.Bold = b.Value()
	}

	if i := fontElem.Italic(); i != nil {
		font.Italic = i.Value()
	}

	if color := fontElem.Color(); color != nil {
		font.Color = r.parseColor(color)
	}

	return font
}

// parseFillElement parses a fill element into CellFill.
func (r *SpreadsheetRenderer) parseFillElement(
	fillElem *elements.Fill,
) *CellFill {
	fill := &CellFill{
		Type: "solid",
	}

	if patternFill := fillElem.PatternFill(); patternFill != nil {
		fill.PatternType = string(
			patternFill.PatternType(),
		)

		if fgColor := patternFill.FgColor(); fgColor != nil {
			fill.FgColor = r.parseColor(fgColor)
		}

		if bgColor := patternFill.BgColor(); bgColor != nil {
			fill.BgColor = r.parseColor(bgColor)
		}
	}

	return fill
}

// parseBorderElement parses a border element into CellBorder.
func (r *SpreadsheetRenderer) parseBorderElement(
	borderElem *elements.Border,
) *CellBorder {
	border := &CellBorder{}

	if left := borderElem.Left(); left != nil {
		border.Left = r.parseBorderSide(left)
	}

	if right := borderElem.Right(); right != nil {
		border.Right = r.parseBorderSide(right)
	}

	if top := borderElem.Top(); top != nil {
		border.Top = r.parseBorderSide(top)
	}

	if bottom := borderElem.Bottom(); bottom != nil {
		border.Bottom = r.parseBorderSide(bottom)
	}

	return border
}

// parseBorderSide parses a border side element.
func (r *SpreadsheetRenderer) parseBorderSide(
	elem *elements.BorderPr,
) *BorderSide {
	if elem == nil {
		return nil
	}

	style := string(elem.Style())
	if style == "" || style == "none" {
		return nil
	}

	side := &BorderSide{
		Style: style,
		Color: core.RGB{R: 0, G: 0, B: 0},
		Width: r.borderStyleToWidth(style),
	}

	if color := elem.Color(); color != nil {
		side.Color = r.parseColor(color)
	}

	return side
}

// borderStyleToWidth converts a border style to a line width in points.
func (r *SpreadsheetRenderer) borderStyleToWidth(
	style string,
) float64 {
	switch style {
	case "thin":
		return 0.5
	case "medium":
		return 1.0
	case "thick":
		return 1.5
	case "hair":
		return 0.25
	default:
		return 0.5
	}
}

// parseAlignmentElement parses an alignment element.
func (r *SpreadsheetRenderer) parseAlignmentElement(
	alignment *elements.Alignment,
) *CellAlignment {
	align := &CellAlignment{
		Horizontal: "left",
		Vertical:   "bottom",
	}

	if horizontal := alignment.Horizontal(); horizontal != "" {
		align.Horizontal = string(horizontal)
	}

	if vertical := alignment.Vertical(); vertical != "" {
		align.Vertical = string(vertical)
	}

	align.WrapText = alignment.WrapText()
	align.Rotation = int(alignment.TextRotation())
	align.Indent = int(alignment.Indent())

	return align
}

// parseColor parses a color element.
func (r *SpreadsheetRenderer) parseColor(
	color *elements.Color,
) core.RGB {
	// Try RGB value
	if rgb := color.RGB(); rgb != "" {
		return parseRGBString(rgb)
	}

	// Try indexed color
	if indexed := color.Indexed(); indexed > 0 {
		return getIndexedColor(int(indexed))
	}

	// Try theme color
	if theme := color.Theme(); theme >= 0 {
		return getThemeColor(int(theme))
	}

	// Default to black
	return core.RGB{R: 0, G: 0, B: 0}
}

// parseRGBString parses an ARGB hex string (e.g., "FFAABBCC").
func parseRGBString(rgb string) core.RGB {
	if len(rgb) < 6 {
		return core.RGB{R: 0, G: 0, B: 0}
	}

	// Skip alpha channel if present
	if len(rgb) == 8 {
		rgb = rgb[2:]
	}

	r, _ := strconv.ParseUint(rgb[0:2], 16, 8)
	g, _ := strconv.ParseUint(rgb[2:4], 16, 8)
	b, _ := strconv.ParseUint(rgb[4:6], 16, 8)

	return core.RGB{
		R: float64(r) / 255.0,
		G: float64(g) / 255.0,
		B: float64(b) / 255.0,
	}
}

// getIndexedColor returns the color for an indexed color value.
func getIndexedColor(index int) core.RGB {
	// Excel indexed colors (simplified)
	colors := []core.RGB{
		{R: 0, G: 0, B: 0}, // 0: Black
		{R: 1, G: 1, B: 1}, // 1: White
		{R: 1, G: 0, B: 0}, // 2: Red
		{R: 0, G: 1, B: 0}, // 3: Green
		{R: 0, G: 0, B: 1}, // 4: Blue
		{R: 1, G: 1, B: 0}, // 5: Yellow
		{
			R: 1,
			G: 0,
			B: 1,
		}, // 6: Magenta
		{R: 0, G: 1, B: 1}, // 7: Cyan
		{
			R: 128.0 / 255,
			G: 0,
			B: 0,
		}, // 8: Dark red
	}

	if index >= 0 && index < len(colors) {
		return colors[index]
	}

	return core.RGB{R: 0, G: 0, B: 0}
}

// getThemeColor returns the color for a theme color index.
func getThemeColor(theme int) core.RGB {
	// Theme colors (simplified)
	colors := []core.RGB{
		{R: 0, G: 0, B: 0}, // 0: dk1 (Text)
		{R: 1, G: 1, B: 1}, // 1: lt1 (Background)
		{
			R: 68.0 / 255,
			G: 84.0 / 255,
			B: 106.0 / 255,
		}, // 2: dk2
		{
			R: 238.0 / 255,
			G: 236.0 / 255,
			B: 225.0 / 255,
		}, // 3: lt2
		{
			R: 68.0 / 255,
			G: 114.0 / 255,
			B: 196.0 / 255,
		}, // 4: Accent1
		{
			R: 237.0 / 255,
			G: 125.0 / 255,
			B: 49.0 / 255,
		}, // 5: Accent2
	}

	if theme >= 0 && theme < len(colors) {
		return colors[theme]
	}

	return core.RGB{R: 0, G: 0, B: 0}
}

// getNumberFormat returns the number format string for a format ID.
func (r *SpreadsheetRenderer) getNumberFormat(
	numFmtId int,
) string {
	// Check custom formats first
	if format, found := r.numberFormats[numFmtId]; found {
		return format
	}

	// Built-in formats
	return getBuiltInNumberFormat(numFmtId)
}

// getBuiltInNumberFormat returns the built-in number format for an ID.
func getBuiltInNumberFormat(numFmtId int) string {
	formats := map[int]string{
		0:  "General",
		1:  "0",
		2:  "0.00",
		3:  "#,##0",
		4:  "#,##0.00",
		9:  "0%",
		10: "0.00%",
		11: "0.00E+00",
		12: "# ?/?",
		13: "# ??/??",
		14: "m/d/yyyy",
		15: "d-mmm-yy",
		16: "d-mmm",
		17: "mmm-yy",
		18: "h:mm AM/PM",
		19: "h:mm:ss AM/PM",
		20: "h:mm",
		21: "h:mm:ss",
		22: "m/d/yyyy h:mm",
		37: "#,##0 ;(#,##0)",
		38: "#,##0 ;[Red](#,##0)",
		39: "#,##0.00;(#,##0.00)",
		40: "#,##0.00;[Red](#,##0.00)",
		45: "mm:ss",
		46: "[h]:mm:ss",
		47: "mmss.0",
		48: "##0.0E+0",
		49: "@",
	}

	if format, found := formats[numFmtId]; found {
		return format
	}

	return "General"
}

// getCellContent gets the formatted content of a cell.
func (r *SpreadsheetRenderer) getCellContent(
	cell *elements.Cell,
	style *CellStyle,
) string {
	// Check if we should show formulas
	if r.options.RenderFormulas {
		if formula := cell.CellFormula(); formula != nil {
			return "=" + formula.Formula()
		}
	}

	// Get cell type
	cellType := cell.DataType()
	cellValue := cell.Value()
	if cellValue == "" {
		return ""
	}

	switch cellType {
	case elements.CellTypeSharedString:
		// Look up in shared strings table
		return r.getSharedString(
			cellValue,
		)

	case elements.CellTypeInlineString:
		if is := cell.InlineString(); is != nil {
			return is.PlainText()
		}
		return cellValue

	case elements.CellTypeBoolean:
		if cellValue == "1" ||
			cellValue == "true" {
			return "TRUE"
		}
		return "FALSE"

	case elements.CellTypeError:
		return cellValue

	case elements.CellTypeNumber,
		elements.CellTypeDate,
		"":
		// Format the number/date according to the number format
		return r.formatValue(
			cellValue,
			style.NumFmt,
		)

	default:
		return cellValue
	}
}

// getSharedString looks up a string in the shared strings table.
func (r *SpreadsheetRenderer) getSharedString(
	indexStr string,
) string {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return ""
	}

	workbookPart := r.doc.WorkbookPart()
	if workbookPart == nil {
		return ""
	}

	sharedStringsPart := workbookPart.SharedStringTablePart()
	if sharedStringsPart == nil {
		return ""
	}

	// Get the string from shared strings table
	str := sharedStringsPart.GetString(index)
	return str
}

// formatValue formats a numeric value according to the number format.
func (r *SpreadsheetRenderer) formatValue(
	value string,
	numFmt string,
) string {
	if value == "" {
		return ""
	}

	// Parse as float
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value
	}

	// Apply number format
	return applyNumberFormat(num, numFmt)
}

// applyNumberFormat applies a number format to a value.
func applyNumberFormat(
	value float64,
	format string,
) string {
	if format == "" || format == "General" {
		// General format
		if value == float64(int64(value)) {
			return fmt.Sprintf("%d", int64(value))
		}
		return fmt.Sprintf("%g", value)
	}

	// Check if it's a date format
	if isDateFormat(format) {
		return formatDate(value, format)
	}

	// Check if it's a percentage format
	if strings.Contains(format, "%") {
		return fmt.Sprintf("%.2f%%", value*100)
	}

	// Check if it's a currency format
	if strings.Contains(format, "$") {
		return fmt.Sprintf("$%.2f", value)
	}

	// Check decimal places
	if strings.Contains(format, ".00") {
		return fmt.Sprintf("%.2f", value)
	}

	if strings.Contains(format, ".0") {
		return fmt.Sprintf("%.1f", value)
	}

	// Default to integer or general
	if value == float64(int64(value)) {
		return fmt.Sprintf("%d", int64(value))
	}

	return fmt.Sprintf("%g", value)
}

// isDateFormat checks if a format string is a date format.
func isDateFormat(format string) bool {
	dateIndicators := []string{
		"y",
		"m",
		"d",
		"h",
		"s",
		"/",
		":",
	}
	for _, indicator := range dateIndicators {
		if strings.Contains(
			strings.ToLower(format),
			indicator,
		) {
			return true
		}
	}
	return false
}

// formatDate formats an Excel date serial number as a date string.
func formatDate(
	serial float64,
	format string,
) string {
	// Excel date serial: days since 1900-01-01 (with 1900 leap year bug)
	// Simplified conversion
	baseDate := time.Date(
		1899,
		12,
		30,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	days := int64(serial)
	fraction := serial - float64(days)

	date := baseDate.AddDate(0, 0, int(days))

	// Add time portion
	seconds := int64(fraction * 24 * 60 * 60)
	date = date.Add(
		time.Duration(seconds) * time.Second,
	)

	// Simple format mapping
	switch {
	case strings.Contains(format, "yyyy"):
		return date.Format("2006-01-02")
	case strings.Contains(format, "yy"):
		return date.Format("01/02/06")
	case strings.Contains(format, "h:mm:ss"):
		return date.Format("15:04:05")
	case strings.Contains(format, "h:mm"):
		return date.Format("15:04")
	default:
		return date.Format("01/02/2006")
	}
}

// renderCellFill renders the background fill for a cell.
func (r *SpreadsheetRenderer) renderCellFill(
	page *core.Page,
	x, y, width, height float64,
	fill *CellFill,
) {
	if fill == nil ||
		(fill.FgColor.R == 0 && fill.FgColor.G == 0 && fill.FgColor.B == 0) {
		return
	}

	// Draw filled rectangle
	drawFillRect(
		page,
		x,
		y,
		width,
		height,
		fill.FgColor,
		1.0,
	)
}

// renderCellBorder renders the borders for a cell.
func (r *SpreadsheetRenderer) renderCellBorder(
	page *core.Page,
	x, y, width, height float64,
	border *CellBorder,
) {
	// Draw left border
	if border.Left != nil {
		drawLine(
			page,
			x,
			y,
			x,
			y+height,
			border.Left.Color,
			border.Left.Width,
		)
	}

	// Draw right border
	if border.Right != nil {
		drawLine(
			page,
			x+width,
			y,
			x+width,
			y+height,
			border.Right.Color,
			border.Right.Width,
		)
	}

	// Draw top border
	if border.Top != nil {
		drawLine(
			page,
			x,
			y,
			x+width,
			y,
			border.Top.Color,
			border.Top.Width,
		)
	}

	// Draw bottom border
	if border.Bottom != nil {
		drawLine(
			page,
			x,
			y+height,
			x+width,
			y+height,
			border.Bottom.Color,
			border.Bottom.Width,
		)
	}
}

// renderCellText renders the text content of a cell.
func (r *SpreadsheetRenderer) renderCellText(
	page *core.Page,
	x, y, width, height float64,
	text string,
	style *CellStyle,
) {
	if text == "" {
		return
	}

	// Load font
	fontObj, err := r.loadFont(style.Font)
	if err != nil {
		return
	}

	// Calculate text position based on alignment
	textX, textY := r.calculateTextPosition(
		x, y, width, height,
		text, style, fontObj,
	)

	// Render text
	_ = drawText(
		r,
		page,
		fontObj,
		text,
		textX,
		textY,
		style.Font.Size,
		style.Font.Color,
	)
}

// calculateTextPosition calculates the position for text rendering based on alignment.
func (r *SpreadsheetRenderer) calculateTextPosition(
	x, y, width, height float64,
	text string,
	style *CellStyle,
	fontObj *font.Font,
) (float64, float64) {
	// Measure text (simplified - use approximate width)
	textWidth := float64(
		len(text),
	) * style.Font.Size * 0.5

	// Horizontal alignment
	textX := x + 2.0 // Default left padding
	switch style.Alignment.Horizontal {
	case "center":
		textX = x + (width-textWidth)/2
	case "right":
		textX = x + width - textWidth - 2.0
	}

	// Vertical alignment
	textY := y + height - 2.0 // Default bottom alignment
	switch style.Alignment.Vertical {
	case "center":
		textY = y + height/2 + style.Font.Size/3 // Approximate vertical center
	case "top":
		textY = y + style.Font.Size
	}

	return textX, textY
}

// loadFont loads a font with the specified properties.
func (r *SpreadsheetRenderer) loadFont(
	cellFont *CellFont,
) (*font.Font, error) {
	// Determine font style
	style := font.StyleRegular
	if cellFont.Bold && cellFont.Italic {
		style = font.StyleBoldItalic
	} else if cellFont.Bold {
		style = font.StyleBold
	} else if cellFont.Italic {
		style = font.StyleItalic
	}

	// Try to find system font
	fontPath, err := font.FindSystemFont(
		cellFont.Family,
		style,
	)
	if err != nil {
		// Fall back to default
		fontPath, err = font.FindSystemFont(
			r.options.DefaultFontFamily,
			font.StyleRegular,
		)
		if err != nil {
			return nil, err
		}
	}

	// Load from cache or file
	fontObj, err := r.fontCache.LoadFile(fontPath)
	if err != nil {
		return nil, err
	}

	return fontObj, nil
}
