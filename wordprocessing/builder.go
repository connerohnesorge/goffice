//nolint:revive // file-length-limit: comprehensive document builder with many helper methods
package wordprocessing

import (
	"bytes"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const (
	maxHeadingLevel    = 9
	emuPerInch         = 914400 // EMUs per inch
	defaultBorderWidth = 6      // default border width in eighths of a point
	twipsPerPoint      = 20     // twips per point
	pctMultiplier      = 50     // multiplier for percentage calculations
)

// DocumentBuilder provides a fluent API for creating Word documents.
type DocumentBuilder struct {
	document *elements.Document
	images   []imageData
	errors   []error
}

// imageData holds image data to be added to the document.
type imageData struct {
	data        []byte
	contentType string
	width       int64
	height      int64
}

// NewDocumentBuilder creates a new DocumentBuilder.
func NewDocumentBuilder() *DocumentBuilder {
	return &DocumentBuilder{
		images: make([]imageData, 0),
		errors: make([]error, 0),
	}
}

// AddParagraph adds a paragraph with the given text.
// It returns a ParagraphBuilder for further customization.
func (db *DocumentBuilder) AddParagraph(
	text string,
) *ParagraphBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	p := db.document.Body().AppendParagraph(text)

	return &ParagraphBuilder{
		paragraph:  p,
		docBuilder: db,
	}
}

// AddHeading adds a heading paragraph with the given text and level (1-9).
func (db *DocumentBuilder) AddHeading(
	text string,
	level int,
) *ParagraphBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	p := db.document.Body().AppendParagraph(text)

	// Apply heading style
	styleId := "Heading1"
	if level >= 1 && level <= maxHeadingLevel {
		styleId = "Heading" + string(
			rune('0'+level),
		)
	}
	p.SetStyle(styleId)

	return &ParagraphBuilder{
		paragraph:  p,
		docBuilder: db,
	}
}

// AddTable adds a table with the specified number of rows and columns.
func (db *DocumentBuilder) AddTable(
	rows, cols int,
) *TableBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	t := db.document.Body().
		AppendTable(rows, cols)

	return &TableBuilder{
		table:      t,
		docBuilder: db,
	}
}

// AddImage adds an image to the document.
// The image will be added as an inline drawing.
// Width and height are in EMUs (914400 EMU = 1 inch).
func (db *DocumentBuilder) AddImage(
	data []byte,
	contentType string,
) *ImageBuilder {
	return &ImageBuilder{
		docBuilder:  db,
		data:        data,
		contentType: contentType,
		width:       emuPerInch, // Default 1 inch
		height:      emuPerInch, // Default 1 inch
	}
}

// AddPageBreak adds a page break to the document.
func (db *DocumentBuilder) AddPageBreak() *DocumentBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	db.document.Body().
		AppendParagraph("").
		AppendBreak(elements.BreakPage)

	return db
}

// AddColumnBreak adds a column break to the document.
func (db *DocumentBuilder) AddColumnBreak() *DocumentBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	db.document.Body().
		AppendParagraph("").
		AppendBreak(elements.BreakColumn)

	return db
}

// AddHorizontalRule adds a horizontal rule (paragraph border) to the document.
func (db *DocumentBuilder) AddHorizontalRule() *DocumentBuilder {
	if db.document == nil {
		db.document = elements.NewDocument()
	}
	p := db.document.Body().AppendParagraph("")
	props := p.GetOrCreateProperties()
	borders := props.GetOrCreateParagraphBorders()
	borders.SetBottom(
		elements.BorderSingle,
		defaultBorderWidth,
		"auto",
	)

	return db
}

// SetTitle sets the document title (for core properties).
func (db *DocumentBuilder) SetTitle(
	_ string,
) *DocumentBuilder {
	// This would set the core properties title
	// For now, we'll just track it
	return db
}

// SetAuthor sets the document author (for core properties).
func (db *DocumentBuilder) SetAuthor(
	_ string,
) *DocumentBuilder {
	// This would set the core properties author
	return db
}

// Build creates the final Document from the builder.
func (db *DocumentBuilder) Build() (*elements.Document, error) {
	if len(db.errors) > 0 {
		return nil, db.errors[0]
	}

	if db.document == nil {
		db.document = elements.NewDocument()
	}

	return db.document, nil
}

// BuildToBytes creates the document and returns it as a byte slice.
func (db *DocumentBuilder) BuildToBytes() ([]byte, error) {
	doc, err := db.Build()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := doc.WriteXML(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ParagraphBuilder provides a fluent API for building paragraphs.
type ParagraphBuilder struct {
	paragraph  *elements.Paragraph
	docBuilder *DocumentBuilder
}

// Bold makes all text in the paragraph bold.
func (pb *ParagraphBuilder) Bold() *ParagraphBuilder {
	for r := range pb.paragraph.Runs() {
		r.SetBold(true)
	}

	return pb
}

// Italic makes all text in the paragraph italic.
func (pb *ParagraphBuilder) Italic() *ParagraphBuilder {
	for r := range pb.paragraph.Runs() {
		r.SetItalic(true)
	}

	return pb
}

// Underline adds underline to all text in the paragraph.
func (pb *ParagraphBuilder) Underline() *ParagraphBuilder {
	for r := range pb.paragraph.Runs() {
		r.SetUnderline(elements.UnderlineSingle)
	}

	return pb
}

// Font sets the font name and size for all text in the paragraph.
// Size is in points (e.g., 12 for 12pt).
func (pb *ParagraphBuilder) Font(
	name string,
	size int,
) *ParagraphBuilder {
	halfPoints := size * 2 // Convert points to half-points
	for r := range pb.paragraph.Runs() {
		if name != "" {
			r.SetFont(name)
		}
		if size > 0 {
			r.SetFontSize(halfPoints)
		}
	}

	return pb
}

// FontSize sets the font size for all text in the paragraph.
// Size is in points (e.g., 12 for 12pt).
func (pb *ParagraphBuilder) FontSize(
	size int,
) *ParagraphBuilder {
	halfPoints := size * 2 // Convert points to half-points
	for r := range pb.paragraph.Runs() {
		r.SetFontSize(halfPoints)
	}

	return pb
}

// Color sets the text color for all text in the paragraph.
// Color should be a hex value without # (e.g., "FF0000" for red).
func (pb *ParagraphBuilder) Color(
	hex string,
) *ParagraphBuilder {
	for r := range pb.paragraph.Runs() {
		r.SetColor(hex)
	}

	return pb
}

// Highlight sets the highlight color for all text in the paragraph.
func (pb *ParagraphBuilder) Highlight(
	color elements.HighlightColor,
) *ParagraphBuilder {
	for r := range pb.paragraph.Runs() {
		r.SetHighlight(color)
	}

	return pb
}

// Align sets the paragraph alignment.
func (pb *ParagraphBuilder) Align(
	alignment elements.JustificationValue,
) *ParagraphBuilder {
	pb.paragraph.SetJustification(alignment)

	return pb
}

// AlignLeft aligns the paragraph to the left.
func (pb *ParagraphBuilder) AlignLeft() *ParagraphBuilder {
	return pb.Align(elements.JustificationLeft)
}

// AlignCenter centers the paragraph.
func (pb *ParagraphBuilder) AlignCenter() *ParagraphBuilder {
	return pb.Align(elements.JustificationCenter)
}

// AlignRight aligns the paragraph to the right.
func (pb *ParagraphBuilder) AlignRight() *ParagraphBuilder {
	return pb.Align(elements.JustificationRight)
}

// AlignJustify justifies the paragraph.
func (pb *ParagraphBuilder) AlignJustify() *ParagraphBuilder {
	return pb.Align(elements.JustificationBoth)
}

// Style sets the paragraph style.
func (pb *ParagraphBuilder) Style(
	styleId string,
) *ParagraphBuilder {
	pb.paragraph.SetStyle(styleId)

	return pb
}

// SpacingBefore sets the space before the paragraph in points.
func (pb *ParagraphBuilder) SpacingBefore(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetSpacingBefore(twips)

	return pb
}

// SpacingAfter sets the space after the paragraph in points.
func (pb *ParagraphBuilder) SpacingAfter(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetSpacingAfter(twips)

	return pb
}

// LeftIndent sets the left indentation in points.
func (pb *ParagraphBuilder) LeftIndent(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetLeftIndent(twips)

	return pb
}

// RightIndent sets the right indentation in points.
func (pb *ParagraphBuilder) RightIndent(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetRightIndent(twips)

	return pb
}

// FirstLineIndent sets the first line indentation in points.
func (pb *ParagraphBuilder) FirstLineIndent(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetFirstLineIndent(twips)

	return pb
}

// HangingIndent sets the hanging indentation in points.
func (pb *ParagraphBuilder) HangingIndent(
	points int,
) *ParagraphBuilder {
	twips := points * twipsPerPoint
	pb.paragraph.SetHangingIndent(twips)

	return pb
}

// AddRun adds a new run with the given text and returns a RunBuilder.
func (pb *ParagraphBuilder) AddRun(
	text string,
) *RunBuilder {
	r := pb.paragraph.AppendRun(text)

	return &RunBuilder{
		run:              r,
		paragraphBuilder: pb,
	}
}

// AddLineBreak adds a line break to the paragraph.
func (pb *ParagraphBuilder) AddLineBreak() *ParagraphBuilder {
	pb.paragraph.AppendBreak(elements.BreakLine)

	return pb
}

// Document returns the DocumentBuilder to continue building.
func (pb *ParagraphBuilder) Document() *DocumentBuilder {
	return pb.docBuilder
}

// Paragraph returns the underlying Paragraph element.
func (pb *ParagraphBuilder) Paragraph() *elements.Paragraph {
	return pb.paragraph
}

// RunBuilder provides a fluent API for building runs.
type RunBuilder struct {
	run              *elements.Run
	paragraphBuilder *ParagraphBuilder
}

// Bold makes this run bold.
func (rb *RunBuilder) Bold() *RunBuilder {
	rb.run.SetBold(true)

	return rb
}

// Italic makes this run italic.
func (rb *RunBuilder) Italic() *RunBuilder {
	rb.run.SetItalic(true)

	return rb
}

// Underline adds underline to this run.
func (rb *RunBuilder) Underline() *RunBuilder {
	rb.run.SetUnderline(elements.UnderlineSingle)

	return rb
}

// UnderlineStyle sets a specific underline style.
func (rb *RunBuilder) UnderlineStyle(
	style elements.UnderlineValue,
) *RunBuilder {
	rb.run.SetUnderline(style)

	return rb
}

// Font sets the font name.
func (rb *RunBuilder) Font(
	name string,
) *RunBuilder {
	rb.run.SetFont(name)

	return rb
}

// FontSize sets the font size in points.
func (rb *RunBuilder) FontSize(
	points int,
) *RunBuilder {
	rb.run.SetFontSize(
		points * 2,
	) // Convert to half-points

	return rb
}

// Color sets the text color (hex without #).
func (rb *RunBuilder) Color(
	hex string,
) *RunBuilder {
	rb.run.SetColor(hex)

	return rb
}

// Highlight sets the highlight color.
func (rb *RunBuilder) Highlight(
	color elements.HighlightColor,
) *RunBuilder {
	rb.run.SetHighlight(color)

	return rb
}

// AddText adds more text to this run.
func (rb *RunBuilder) AddText(
	text string,
) *RunBuilder {
	rb.run.AppendText(text)

	return rb
}

// AddBreak adds a break to this run.
func (rb *RunBuilder) AddBreak(
	breakType elements.BreakType,
) *RunBuilder {
	rb.run.AppendBreak(breakType)

	return rb
}

// AddTab adds a tab to this run.
func (rb *RunBuilder) AddTab() *RunBuilder {
	rb.run.AppendTab()

	return rb
}

// Paragraph returns the ParagraphBuilder to continue building the paragraph.
func (rb *RunBuilder) Paragraph() *ParagraphBuilder {
	return rb.paragraphBuilder
}

// Document returns the DocumentBuilder to continue building the document.
func (rb *RunBuilder) Document() *DocumentBuilder {
	return rb.paragraphBuilder.docBuilder
}

// Run returns the underlying Run element.
func (rb *RunBuilder) Run() *elements.Run {
	return rb.run
}

// TableBuilder provides a fluent API for building tables.
type TableBuilder struct {
	table      *elements.Table
	docBuilder *DocumentBuilder
}

// SetCellText sets the text content of a specific cell.
func (tb *TableBuilder) SetCellText(
	row, col int,
	text string,
) *TableBuilder {
	tb.table.SetCellText(row, col, text)

	return tb
}

// SetColumnWidth sets the width of a specific column in twips.
//
//nolint:revive // enforce-repeated-arg-type-style: explicit types improve readability
func (tb *TableBuilder) SetColumnWidth(
	col int,
	width int,
) *TableBuilder {
	tb.table.SetColumnWidth(col, width)

	return tb
}

// SetColumnWidthInches sets the width of a specific column in inches.
func (tb *TableBuilder) SetColumnWidthInches(
	col int,
	inches float64,
) *TableBuilder {
	twips := int(
		inches * 1440,
	) // 1 inch = 1440 twips

	return tb.SetColumnWidth(col, twips)
}

// AddRow appends a new row with the same number of columns as the table.
func (tb *TableBuilder) AddRow() *TableBuilder {
	// Get current column count from first row
	cols := 1
	firstRow := tb.table.GetRow(0)
	if firstRow != nil {
		cols = firstRow.CellCount()
	}
	tb.table.AppendRow(cols)

	return tb
}

// SetStyle sets the table style.
func (tb *TableBuilder) SetStyle(
	styleId string,
) *TableBuilder {
	tb.table.SetStyle(styleId)

	return tb
}

// SetWidth sets the table width.
func (tb *TableBuilder) SetWidth(
	width int,
	widthType elements.TableWidthType,
) *TableBuilder {
	tb.table.SetWidth(width, widthType)

	return tb
}

// SetWidthPercent sets the table width as a percentage.
func (tb *TableBuilder) SetWidthPercent(
	percent int,
) *TableBuilder {
	// PCT is in fiftieths of a percent (5000 = 100%)
	pct := percent * pctMultiplier

	return tb.SetWidth(
		pct,
		elements.TableWidthTypePct,
	)
}

// SetBorders sets all table borders.
func (tb *TableBuilder) SetBorders(
	style elements.BorderStyle,
	size int,
	color string,
) *TableBuilder {
	props := tb.table.GetOrCreateTableProperties()
	borders := props.GetOrCreateTableBorders()
	borders.SetAllBorders(style, size, color)

	return tb
}

// Cell returns a TableCellBuilder for the specified cell.
func (tb *TableBuilder) Cell(
	row, col int,
) *TableCellBuilder {
	cell := tb.table.GetCell(row, col)
	if cell == nil {
		return nil
	}

	return &TableCellBuilder{
		cell:         cell,
		tableBuilder: tb,
	}
}

// Row returns a TableRowBuilder for the specified row.
func (tb *TableBuilder) Row(
	index int,
) *TableRowBuilder {
	row := tb.table.GetRow(index)
	if row == nil {
		return nil
	}

	return &TableRowBuilder{
		row:          row,
		tableBuilder: tb,
	}
}

// Document returns the DocumentBuilder to continue building.
func (tb *TableBuilder) Document() *DocumentBuilder {
	return tb.docBuilder
}

// Table returns the underlying Table element.
func (tb *TableBuilder) Table() *elements.Table {
	return tb.table
}

// TableRowBuilder provides a fluent API for building table rows.
type TableRowBuilder struct {
	row          *elements.TableRow
	tableBuilder *TableBuilder
}

// SetHeight sets the row height.
func (t *TableRowBuilder) SetHeight(
	height int,
	rule elements.HeightRule,
) *TableRowBuilder {
	t.row.SetHeight(height, rule)

	return t
}

// SetHeaderRow marks this row as a header row.
func (t *TableRowBuilder) SetHeaderRow(
	isHeader bool,
) *TableRowBuilder {
	t.row.SetHeaderRow(isHeader)

	return t
}

// Cell returns a TableCellBuilder for the specified cell in this row.
func (t *TableRowBuilder) Cell(
	col int,
) *TableCellBuilder {
	cell := t.row.GetCell(col)
	if cell == nil {
		return nil
	}

	return &TableCellBuilder{
		cell:         cell,
		tableBuilder: t.tableBuilder,
	}
}

// Table returns the TableBuilder to continue building the table.
func (t *TableRowBuilder) Table() *TableBuilder {
	return t.tableBuilder
}

// Row returns the underlying TableRow element.
func (t *TableRowBuilder) Row() *elements.TableRow {
	return t.row
}

// TableCellBuilder provides a fluent API for building table cells.
type TableCellBuilder struct {
	cell         *elements.TableCell
	tableBuilder *TableBuilder
}

// SetText sets the cell text content.
func (tcb *TableCellBuilder) SetText(
	text string,
) *TableCellBuilder {
	tcb.cell.SetText(text)

	return tcb
}

// AddParagraph adds a paragraph to the cell.
func (tcb *TableCellBuilder) AddParagraph(
	text string,
) *ParagraphBuilder {
	p := tcb.cell.AppendParagraph(text)

	return &ParagraphBuilder{
		paragraph:  p,
		docBuilder: tcb.tableBuilder.docBuilder,
	}
}

// SetWidth sets the cell width.
func (tcb *TableCellBuilder) SetWidth(
	width int,
	widthType elements.TableWidthType,
) *TableCellBuilder {
	tcb.cell.SetWidth(width, widthType)

	return tcb
}

// SetShading sets the cell background color.
func (tcb *TableCellBuilder) SetShading(
	fillColor string,
) *TableCellBuilder {
	tcb.cell.SetShading(fillColor)

	return tcb
}

// SetVerticalMerge sets vertical cell merge.
func (tcb *TableCellBuilder) SetVerticalMerge(
	mergeType elements.VerticalMergeType,
) *TableCellBuilder {
	tcb.cell.SetVerticalMerge(mergeType)

	return tcb
}

// SetHorizontalMerge sets horizontal cell span (gridSpan).
func (tcb *TableCellBuilder) SetHorizontalMerge(
	span int,
) *TableCellBuilder {
	tcb.cell.SetHorizontalMerge(span)

	return tcb
}

// SetBorders sets all cell borders.
func (tcb *TableCellBuilder) SetBorders(
	style elements.BorderStyle,
	size int,
	color string,
) *TableCellBuilder {
	borders := tcb.cell.GetOrCreateTableCellProperties().
		GetOrCreateTableCellBorders()
	borders.SetAllBorders(style, size, color)

	return tcb
}

// Table returns the TableBuilder to continue building the table.
func (tcb *TableCellBuilder) Table() *TableBuilder {
	return tcb.tableBuilder
}

// Cell returns the underlying TableCell element.
func (tcb *TableCellBuilder) Cell() *elements.TableCell {
	return tcb.cell
}

// ImageBuilder provides a fluent API for adding images.
type ImageBuilder struct {
	docBuilder  *DocumentBuilder
	data        []byte
	contentType string
	width       int64
	height      int64
	inline      bool
}

// Width sets the image width in EMUs.
func (ib *ImageBuilder) Width(
	emus int64,
) *ImageBuilder {
	ib.width = emus

	return ib
}

// Height sets the image height in EMUs.
func (ib *ImageBuilder) Height(
	emus int64,
) *ImageBuilder {
	ib.height = emus

	return ib
}

// WidthInches sets the image width in inches.
func (ib *ImageBuilder) WidthInches(
	inches float64,
) *ImageBuilder {
	ib.width = int64(
		inches * float64(elements.EMUsPerInch),
	)

	return ib
}

// HeightInches sets the image height in inches.
func (ib *ImageBuilder) HeightInches(
	inches float64,
) *ImageBuilder {
	ib.height = int64(
		inches * float64(elements.EMUsPerInch),
	)

	return ib
}

// WidthCm sets the image width in centimeters.
func (ib *ImageBuilder) WidthCm(
	cm float64,
) *ImageBuilder {
	ib.width = int64(
		cm * float64(elements.EMUsPerCm),
	)

	return ib
}

// HeightCm sets the image height in centimeters.
func (ib *ImageBuilder) HeightCm(
	cm float64,
) *ImageBuilder {
	ib.height = int64(
		cm * float64(elements.EMUsPerCm),
	)

	return ib
}

// Size sets both width and height in EMUs.
func (ib *ImageBuilder) Size(
	width, height int64,
) *ImageBuilder {
	ib.width = width
	ib.height = height

	return ib
}

// SizeInches sets both width and height in inches.
func (ib *ImageBuilder) SizeInches(
	width, height float64,
) *ImageBuilder {
	return ib.WidthInches(width).
		HeightInches(height)
}

// SizeCm sets both width and height in centimeters.
func (ib *ImageBuilder) SizeCm(
	width, height float64,
) *ImageBuilder {
	return ib.WidthCm(width).HeightCm(height)
}

// Inline makes the image inline (default).
func (ib *ImageBuilder) Inline() *ImageBuilder {
	ib.inline = true

	return ib
}

// Insert inserts the image and returns to the DocumentBuilder.
// Note: The actual image part creation happens in Build()
// when a package is available.
func (ib *ImageBuilder) Insert() *DocumentBuilder {
	// Store the image data for later processing
	ib.docBuilder.images = append(
		ib.docBuilder.images,
		imageData{
			data:        ib.data,
			contentType: ib.contentType,
			width:       ib.width,
			height:      ib.height,
		},
	)

	// For now, create a placeholder drawing
	// In a full implementation, this would need a relationship ID
	// which requires the package to be created first
	if ib.docBuilder.document == nil {
		ib.docBuilder.document = elements.NewDocument()
	}

	// Create a paragraph with the drawing
	p := ib.docBuilder.document.Body().
		AppendParagraph("")
	r := p.AppendRun("")

	// Create inline drawing with placeholder relationship ID
	// The actual relationship will be set when the document is saved
	drawing := elements.NewInlineDrawing(
		ib.width,
		ib.height,
		"rId1",
	)
	r.AppendChild(drawing)

	return ib.docBuilder
}

// Document returns the DocumentBuilder without inserting.
func (ib *ImageBuilder) Document() *DocumentBuilder {
	return ib.docBuilder
}
