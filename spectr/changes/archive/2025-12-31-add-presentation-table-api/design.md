# PowerPoint Table API - Design

## Architecture Overview

The PowerPoint table implementation follows a three-layer architecture:

```
┌─────────────────────────────────────────┐
│    Presentation Layer (presentation/)   │
│  - Table (high-level API)               │
│  - Slide.AddTable()                     │
│  - Row/Cell manipulation                │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│   DrawingML Layer (drawingml/table/)    │
│  - Table, TableProperties, TableGrid    │
│  - TableRow, TableCell                  │
│  - XML serialization                    │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│  Presentation Elements (elements/)       │
│  - GraphicFrame (existing)              │
│  - Graphic, GraphicData                 │
└─────────────────────────────────────────┘
```

## DrawingML Table Element Design

### Core Types

```go
package table

// Table represents a DrawingML table (a:tbl)
type Table struct {
	openxml.CompositeElementBase

	TableProperties *TableProperties  `xml:"tblPr,omitempty"`
	TableGrid       *TableGrid        `xml:"tblGrid,omitempty"`
	Rows            []*TableRow       `xml:"tr,omitempty"`
}

func NewTable() *Table {
	t := &Table{}
	t.CompositeElementBase = openxml.NewCompositeElement(
		"http://schemas.openxmlformats.org/drawingml/2006/main",
		"tbl",
		"a",
	)
	t.TableProperties = NewTableProperties()
	t.TableGrid = NewTableGrid()
	return t
}

// TableProperties represents table-level properties (a:tblPr)
type TableProperties struct {
	openxml.CompositeElementBase

	FirstRow      *types.BoolValue   `xml:"firstRow,attr,omitempty"`
	FirstCol      *types.BoolValue   `xml:"firstCol,attr,omitempty"`
	LastRow       *types.BoolValue   `xml:"lastRow,attr,omitempty"`
	LastCol       *types.BoolValue   `xml:"lastCol,attr,omitempty"`
	BandRow       *types.BoolValue   `xml:"bandRow,attr,omitempty"`
	BandCol       *types.BoolValue   `xml:"bandCol,attr,omitempty"`

	TableStyleId  *TableStyleId      `xml:"tableStyleId,omitempty"`
}

// TableStyleId represents a GUID reference to a table style
type TableStyleId struct {
	openxml.CompositeElementBase
	Value string `xml:",chardata"`
}

// TableGrid represents column definitions (a:tblGrid)
type TableGrid struct {
	openxml.CompositeElementBase
	Columns []*GridColumn `xml:"gridCol,omitempty"`
}

// GridColumn represents a column width definition (a:gridCol)
type GridColumn struct {
	openxml.CompositeElementBase
	Width *types.Int64Value `xml:"w,attr,omitempty"`  // in EMUs
}

// TableRow represents a table row (a:tr)
type TableRow struct {
	openxml.CompositeElementBase

	Height *types.Int64Value `xml:"h,attr,omitempty"`  // in EMUs
	Cells  []*TableCell      `xml:"tc,omitempty"`
}

// TableCell represents a table cell (a:tc)
type TableCell struct {
	openxml.CompositeElementBase

	TextBody       *drawingml.TextBody  `xml:"txBody,omitempty"`
	CellProperties *CellProperties      `xml:"tcPr,omitempty"`

	// Merge attributes
	GridSpan  *types.IntValue `xml:"gridSpan,attr,omitempty"`
	RowSpan   *types.IntValue `xml:"rowSpan,attr,omitempty"`
	HMerge    *types.BoolValue `xml:"hMerge,attr,omitempty"`
	VMerge    *types.BoolValue `xml:"vMerge,attr,omitempty"`
}

// CellProperties represents cell-level properties (a:tcPr)
type CellProperties struct {
	openxml.CompositeElementBase

	MarginLeft   *types.Int32Value     `xml:"marL,attr,omitempty"`
	MarginRight  *types.Int32Value     `xml:"marR,attr,omitempty"`
	MarginTop    *types.Int32Value     `xml:"marT,attr,omitempty"`
	MarginBottom *types.Int32Value     `xml:"marB,attr,omitempty"`

	Anchor       *TextAnchoringType    `xml:"anchor,attr,omitempty"`
	AnchorCenter *types.BoolValue      `xml:"anchorCtr,attr,omitempty"`
	HorzOverflow *TextHorzOverflowType `xml:"horzOverflow,attr,omitempty"`

	SolidFill    *drawingml.SolidFill  `xml:"solidFill,omitempty"`
	// Additional fill types as needed

	BorderLeft   *CellBorder           `xml:"lnL,omitempty"`
	BorderRight  *CellBorder           `xml:"lnR,omitempty"`
	BorderTop    *CellBorder           `xml:"lnT,omitempty"`
	BorderBottom *CellBorder           `xml:"lnB,omitempty"`
}

// CellBorder represents a cell border line
type CellBorder struct {
	openxml.CompositeElementBase

	Width     *types.Int32Value    `xml:"w,attr,omitempty"`
	SolidFill *drawingml.SolidFill `xml:"solidFill,omitempty"`
}

// TextAnchoringType enum for vertical alignment
type TextAnchoringType string

const (
	TextAnchorTop    TextAnchoringType = "t"
	TextAnchorCenter TextAnchoringType = "ctr"
	TextAnchorBottom TextAnchoringType = "b"
)

// TextHorzOverflowType enum for horizontal overflow
type TextHorzOverflowType string

const (
	TextOverflowClip     TextHorzOverflowType = "clip"
	TextOverflowOverflow TextHorzOverflowType = "overflow"
)
```

## High-Level Presentation API Design

### Table Wrapper

```go
package presentation

// Table provides high-level API for PowerPoint tables
type Table struct {
	graphicFrame *elements.GraphicFrame
	table        *table.Table
	slide        *Slide
}

// CreateTable creates a new table on a slide
func (s *Slide) AddTable(rows, cols int) (*Table, error) {
	if rows < 1 || cols < 1 {
		return nil, fmt.Errorf("table must have at least 1 row and 1 column")
	}

	// Create GraphicFrame
	gf := s.ShapeTree().AddGraphicFrame()
	gf.SetPosition(100*units.PointToEMU, 100*units.PointToEMU) // Default position
	gf.SetSize(400*units.PointToEMU, 300*units.PointToEMU)     // Default size

	// Create table structure
	tbl := table.NewTable()

	// Create table grid (columns)
	defaultColWidth := int64(400 * units.PointToEMU / cols)
	for i := 0; i < cols; i++ {
		col := table.NewGridColumn()
		col.Width = types.NewInt64Value(defaultColWidth)
		tbl.TableGrid.Columns = append(tbl.TableGrid.Columns, col)
	}

	// Create rows
	for i := 0; i < rows; i++ {
		row := table.NewTableRow()
		row.Height = types.NewInt64Value(int64(50 * units.PointToEMU)) // Default row height

		// Create cells
		for j := 0; j < cols; j++ {
			cell := table.NewTableCell()
			cell.TextBody = drawingml.NewTextBody()
			cell.CellProperties = table.NewCellProperties()
			row.Cells = append(row.Cells, cell)
		}

		tbl.Rows = append(tbl.Rows, row)
	}

	// Link table to graphic frame
	elements.LinkGraphicFrameToTable(gf, tbl)

	return &Table{
		graphicFrame: gf,
		table:        tbl,
		slide:        s,
	}, nil
}

// LinkGraphicFrameToTable links a table to a graphic frame
// (Helper function in presentation/elements/helpers.go)
func LinkGraphicFrameToTable(gf *GraphicFrame, tbl *table.Table) {
	// Create Graphic element
	graphic := drawingml.NewGraphic()

	// Create GraphicData with table URI
	graphicData := drawingml.NewGraphicData()
	graphicData.Uri = types.NewStringValue("http://schemas.openxmlformats.org/drawingml/2006/table")

	// Append table to graphic data
	graphicData.AppendChild(tbl)
	graphic.GraphicData = graphicData

	// Set graphic on frame
	gf.Graphic = graphic
}
```

### Row Manipulation

```go
// RowCount returns the number of rows in the table
func (t *Table) RowCount() int {
	return len(t.table.Rows)
}

// ColumnCount returns the number of columns in the table
func (t *Table) ColumnCount() int {
	if len(t.table.TableGrid.Columns) > 0 {
		return len(t.table.TableGrid.Columns)
	}
	if len(t.table.Rows) > 0 {
		return len(t.table.Rows[0].Cells)
	}
	return 0
}

// GetRow returns the row at the specified index (0-based)
func (t *Table) GetRow(index int) (*table.TableRow, error) {
	if index < 0 || index >= len(t.table.Rows) {
		return nil, fmt.Errorf("row index %d out of bounds (0-%d)", index, len(t.table.Rows)-1)
	}
	return t.table.Rows[index], nil
}

// AppendRow adds a new row at the end of the table
func (t *Table) AppendRow() (*table.TableRow, error) {
	cols := t.ColumnCount()
	row := table.NewTableRow()
	row.Height = types.NewInt64Value(int64(50 * units.PointToEMU))

	// Create cells matching column count
	for i := 0; i < cols; i++ {
		cell := table.NewTableCell()
		cell.TextBody = drawingml.NewTextBody()
		cell.CellProperties = table.NewCellProperties()
		row.Cells = append(row.Cells, cell)
	}

	t.table.Rows = append(t.table.Rows, row)
	return row, nil
}

// InsertRow inserts a new row at the specified index
func (t *Table) InsertRow(index int) (*table.TableRow, error) {
	if index < 0 || index > len(t.table.Rows) {
		return nil, fmt.Errorf("row index %d out of bounds (0-%d)", index, len(t.table.Rows))
	}

	cols := t.ColumnCount()
	row := table.NewTableRow()
	row.Height = types.NewInt64Value(int64(50 * units.PointToEMU))

	// Create cells
	for i := 0; i < cols; i++ {
		cell := table.NewTableCell()
		cell.TextBody = drawingml.NewTextBody()
		cell.CellProperties = table.NewCellProperties()
		row.Cells = append(row.Cells, cell)
	}

	// Insert at index
	t.table.Rows = append(t.table.Rows[:index], append([]*table.TableRow{row}, t.table.Rows[index:]...)...)
	return row, nil
}

// DeleteRow removes the row at the specified index
func (t *Table) DeleteRow(index int) error {
	if index < 0 || index >= len(t.table.Rows) {
		return fmt.Errorf("row index %d out of bounds (0-%d)", index, len(t.table.Rows)-1)
	}
	t.table.Rows = append(t.table.Rows[:index], t.table.Rows[index+1:]...)
	return nil
}
```

### Cell Access and Content

```go
// GetCell returns the cell at the specified row and column (0-based)
func (t *Table) GetCell(row, col int) (*table.TableCell, error) {
	if row < 0 || row >= len(t.table.Rows) {
		return nil, fmt.Errorf("row index %d out of bounds", row)
	}
	if col < 0 || col >= len(t.table.Rows[row].Cells) {
		return nil, fmt.Errorf("column index %d out of bounds", col)
	}
	return t.table.Rows[row].Cells[col], nil
}

// SetCellText sets the text content of a cell
func (t *Table) SetCellText(row, col int, text string) error {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return err
	}

	// Ensure TextBody exists
	if cell.TextBody == nil {
		cell.TextBody = drawingml.NewTextBody()
	}

	// Clear existing paragraphs
	cell.TextBody.Paragraphs = nil

	// Add new paragraph with text
	para := drawingml.NewParagraph()
	run := drawingml.NewRun()
	run.Text = drawingml.NewText()
	run.Text.Value = text
	para.Runs = append(para.Runs, run)
	cell.TextBody.Paragraphs = append(cell.TextBody.Paragraphs, para)

	return nil
}

// GetCellText retrieves the text content of a cell
func (t *Table) GetCellText(row, col int) (string, error) {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return "", err
	}

	if cell.TextBody == nil || len(cell.TextBody.Paragraphs) == 0 {
		return "", nil
	}

	// Concatenate text from all runs in all paragraphs
	var text string
	for _, para := range cell.TextBody.Paragraphs {
		for _, run := range para.Runs {
			if run.Text != nil {
				text += run.Text.Value
			}
		}
		text += "\n"
	}

	return strings.TrimSpace(text), nil
}
```

### Column Width Management

```go
// SetColumnWidth sets the width of a column in points
func (t *Table) SetColumnWidth(col int, widthPt float64) error {
	if col < 0 || col >= len(t.table.TableGrid.Columns) {
		return fmt.Errorf("column index %d out of bounds", col)
	}

	widthEMU := int64(widthPt * units.PointToEMU)
	t.table.TableGrid.Columns[col].Width = types.NewInt64Value(widthEMU)
	return nil
}

// GetColumnWidth returns the width of a column in points
func (t *Table) GetColumnWidth(col int) (float64, error) {
	if col < 0 || col >= len(t.table.TableGrid.Columns) {
		return 0, fmt.Errorf("column index %d out of bounds", col)
	}

	if t.table.TableGrid.Columns[col].Width == nil {
		return 0, nil
	}

	widthEMU := t.table.TableGrid.Columns[col].Width.Value()
	return float64(widthEMU) / units.PointToEMU, nil
}

// SetEqualColumnWidths distributes table width equally across all columns
func (t *Table) SetEqualColumnWidths() {
	if len(t.table.TableGrid.Columns) == 0 {
		return
	}

	// Get table width from GraphicFrame
	_, _, width, _ := t.graphicFrame.GetBounds()

	colWidth := width / int64(len(t.table.TableGrid.Columns))
	for _, col := range t.table.TableGrid.Columns {
		col.Width = types.NewInt64Value(colWidth)
	}
}
```

### Table Styling

```go
// SetStyle applies a table style using a GUID reference
func (t *Table) SetStyle(styleGUID string) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}

	styleId := table.NewTableStyleId()
	styleId.Value = styleGUID
	t.table.TableProperties.TableStyleId = styleId
}

// SetBandedRows enables or disables banded row formatting
func (t *Table) SetBandedRows(enabled bool) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}
	t.table.TableProperties.BandRow = types.NewBoolValue(enabled)
}

// SetBandedColumns enables or disables banded column formatting
func (t *Table) SetBandedColumns(enabled bool) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}
	t.table.TableProperties.BandCol = types.NewBoolValue(enabled)
}

// SetFirstRowBold enables special formatting for the first row (header)
func (t *Table) SetFirstRowBold(enabled bool) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}
	t.table.TableProperties.FirstRow = types.NewBoolValue(enabled)
}

// SetFirstColumnBold enables special formatting for the first column
func (t *Table) SetFirstColumnBold(enabled bool) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}
	t.table.TableProperties.FirstCol = types.NewBoolValue(enabled)
}

// SetLastRowBold enables special formatting for the last row (total row)
func (t *Table) SetLastRowBold(enabled bool) {
	if t.table.TableProperties == nil {
		t.table.TableProperties = table.NewTableProperties()
	}
	t.table.TableProperties.LastRow = types.NewBoolValue(enabled)
}

// Built-in PowerPoint table style GUIDs (common styles)
const (
	TableStyleMedium2        = "{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}"
	TableStyleLight1         = "{2D5ABB26-0587-4C30-8999-92F81FD0307C}"
	TableStyleDark1          = "{D113A9D2-9D6B-4929-AA2D-F23B5EE8CBE7}"
	TableStyleMediumShading1 = "{793D81CF-94F2-401A-BA57-92F5A7B2D0C5}"
)
```

### Position and Size

```go
// SetPosition sets the table position on the slide in points
func (t *Table) SetPosition(xPt, yPt float64) {
	xEMU := int64(xPt * units.PointToEMU)
	yEMU := int64(yPt * units.PointToEMU)
	t.graphicFrame.SetPosition(xEMU, yEMU)
}

// GetPosition returns the table position in points
func (t *Table) GetPosition() (xPt, yPt float64) {
	x, y, _, _ := t.graphicFrame.GetBounds()
	return float64(x) / units.PointToEMU, float64(y) / units.PointToEMU
}

// SetSize sets the table size in points
func (t *Table) SetSize(widthPt, heightPt float64) {
	widthEMU := int64(widthPt * units.PointToEMU)
	heightEMU := int64(heightPt * units.PointToEMU)
	t.graphicFrame.SetSize(widthEMU, heightEMU)

	// Update table grid to match new width
	t.SetEqualColumnWidths()
}

// GetSize returns the table size in points
func (t *Table) GetSize() (widthPt, heightPt float64) {
	_, _, width, height := t.graphicFrame.GetBounds()
	return float64(width) / units.PointToEMU, float64(height) / units.PointToEMU
}
```

## Cell Formatting

```go
// SetCellBackgroundColor sets the background fill color of a cell
func (t *Table) SetCellBackgroundColor(row, col int, r, g, b uint8) error {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return err
	}

	if cell.CellProperties == nil {
		cell.CellProperties = table.NewCellProperties()
	}

	fill := drawingml.NewSolidFill()
	color := drawingml.NewRGBColor()
	color.R = r
	color.G = g
	color.B = b
	fill.Color = color

	cell.CellProperties.SolidFill = fill
	return nil
}

// SetCellAlignment sets the vertical alignment of cell content
func (t *Table) SetCellAlignment(row, col int, align table.TextAnchoringType) error {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return err
	}

	if cell.CellProperties == nil {
		cell.CellProperties = table.NewCellProperties()
	}

	cell.CellProperties.Anchor = &align
	return nil
}

// SetCellMargins sets the cell padding/margins in points
func (t *Table) SetCellMargins(row, col int, left, right, top, bottom float64) error {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return err
	}

	if cell.CellProperties == nil {
		cell.CellProperties = table.NewCellProperties()
	}

	cell.CellProperties.MarginLeft = types.NewInt32Value(int32(left * units.PointToEMU))
	cell.CellProperties.MarginRight = types.NewInt32Value(int32(right * units.PointToEMU))
	cell.CellProperties.MarginTop = types.NewInt32Value(int32(top * units.PointToEMU))
	cell.CellProperties.MarginBottom = types.NewInt32Value(int32(bottom * units.PointToEMU))

	return nil
}
```

## PDF Rendering Integration

```go
package presentation

// RenderTable renders a PowerPoint table to PDF
func (r *TableRenderer) RenderTable(ctx *RenderingContext, table *Table) error {
	// Get table bounds
	x, y, width, height := table.graphicFrame.GetBounds()
	xPt := float64(x) / units.PointToEMU
	yPt := float64(y) / units.PointToEMU
	widthPt := float64(width) / units.PointToEMU
	heightPt := float64(height) / units.PointToEMU

	// Calculate cell dimensions
	rows := table.RowCount()
	cols := table.ColumnCount()

	// Get column widths
	colWidths := make([]float64, cols)
	for i := 0; i < cols; i++ {
		w, _ := table.GetColumnWidth(i)
		colWidths[i] = w
	}

	// Render table grid (borders)
	r.renderTableGrid(ctx, xPt, yPt, widthPt, heightPt, rows, cols, colWidths)

	// Render cells
	currentY := yPt
	for row := 0; row < rows; row++ {
		rowHeight := r.getRowHeight(table, row)
		currentX := xPt

		for col := 0; col < cols; col++ {
			cellWidth := colWidths[col]

			// Render cell background
			r.renderCellBackground(ctx, table, row, col, currentX, currentY, cellWidth, rowHeight)

			// Render cell text
			r.renderCellText(ctx, table, row, col, currentX, currentY, cellWidth, rowHeight)

			currentX += cellWidth
		}

		currentY += rowHeight
	}

	return nil
}

func (r *TableRenderer) renderCellText(ctx *RenderingContext, table *Table, row, col int, x, y, width, height float64) {
	text, err := table.GetCellText(row, col)
	if err != nil || text == "" {
		return
	}

	// Get cell properties for alignment
	cell, _ := table.GetCell(row, col)
	var vAlign table.TextAnchoringType = table.TextAnchorTop
	if cell.CellProperties != nil && cell.CellProperties.Anchor != nil {
		vAlign = *cell.CellProperties.Anchor
	}

	// Calculate text position based on alignment
	textY := y + 5 // Default top with padding
	if vAlign == table.TextAnchorCenter {
		textY = y + height/2
	} else if vAlign == table.TextAnchorBottom {
		textY = y + height - 5
	}

	// Render text
	ctx.Page.SetFont("Arial", 11)
	ctx.Page.DrawText(x+5, textY, text)
}

func (r *TableRenderer) renderCellBackground(ctx *RenderingContext, table *Table, row, col int, x, y, width, height float64) {
	cell, err := table.GetCell(row, col)
	if err != nil {
		return
	}

	if cell.CellProperties == nil || cell.CellProperties.SolidFill == nil {
		return
	}

	// Get fill color
	color := cell.CellProperties.SolidFill.Color
	if rgbColor, ok := color.(*drawingml.RGBColor); ok {
		r := float64(rgbColor.R) / 255.0
		g := float64(rgbColor.G) / 255.0
		b := float64(rgbColor.B) / 255.0

		ctx.Page.SetFillColor(r, g, b)
		ctx.Page.DrawRectangle(x, y, width, height, true, false)
	}
}

func (r *TableRenderer) renderTableGrid(ctx *RenderingContext, x, y, width, height float64, rows, cols int, colWidths []float64) {
	// Set line color and width for borders
	ctx.Page.SetStrokeColor(0, 0, 0) // Black
	ctx.Page.SetLineWidth(0.5)

	// Draw horizontal lines
	currentY := y
	rowHeight := height / float64(rows)
	for i := 0; i <= rows; i++ {
		ctx.Page.DrawLine(x, currentY, x+width, currentY)
		currentY += rowHeight
	}

	// Draw vertical lines
	currentX := x
	for i := 0; i <= cols; i++ {
		ctx.Page.DrawLine(currentX, y, currentX, y+height)
		if i < cols {
			currentX += colWidths[i]
		}
	}
}
```

## Testing Strategy

### Unit Tests

```go
func TestTableCreation(t *testing.T) {
	slide := presentation.NewSlide()
	table, err := slide.AddTable(3, 4)
	if err != nil {
		t.Fatal(err)
	}

	if table.RowCount() != 3 {
		t.Errorf("Expected 3 rows, got %d", table.RowCount())
	}

	if table.ColumnCount() != 4 {
		t.Errorf("Expected 4 columns, got %d", table.ColumnCount())
	}
}

func TestCellTextManipulation(t *testing.T) {
	slide := presentation.NewSlide()
	table, _ := slide.AddTable(2, 2)

	// Set text
	table.SetCellText(0, 0, "Header")
	table.SetCellText(1, 1, "Data")

	// Get text
	text, err := table.GetCellText(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if text != "Header" {
		t.Errorf("Expected 'Header', got '%s'", text)
	}
}

func TestRowManipulation(t *testing.T) {
	slide := presentation.NewSlide()
	table, _ := slide.AddTable(2, 3)

	// Append row
	table.AppendRow()
	if table.RowCount() != 3 {
		t.Errorf("Expected 3 rows after append, got %d", table.RowCount())
	}

	// Insert row
	table.InsertRow(1)
	if table.RowCount() != 4 {
		t.Errorf("Expected 4 rows after insert, got %d", table.RowCount())
	}

	// Delete row
	table.DeleteRow(0)
	if table.RowCount() != 3 {
		t.Errorf("Expected 3 rows after delete, got %d", table.RowCount())
	}
}
```

### Integration Tests

```go
func TestTableRoundtrip(t *testing.T) {
	// Create presentation with table
	pres := presentation.NewPresentation()
	slide := pres.AddSlide()
	table, _ := slide.AddTable(3, 3)
	table.SetCellText(0, 0, "A1")
	table.SetCellText(0, 1, "B1")
	table.SetStyle(presentation.TableStyleMedium2)
	table.SetBandedRows(true)

	// Save to file
	tmpFile := "/tmp/test_table.pptx"
	pres.SaveToFile(tmpFile)
	defer os.Remove(tmpFile)

	// Load file
	pres2, err := presentation.OpenPresentation(tmpFile)
	if err != nil {
		t.Fatal(err)
	}

	// Verify table exists
	slide2 := pres2.GetSlide(0)
	frames := slide2.ShapeTree().GraphicFrames()
	if len(frames) != 1 {
		t.Fatalf("Expected 1 graphic frame, got %d", len(frames))
	}

	// Verify table content (would need table extraction from frame)
	// ...
}
```

## Example Usage

```go
package main

import "goffice/presentation"

func main() {
	// Create presentation
	pres := presentation.NewPresentation()
	slide := pres.AddSlide()

	// Create 4x3 table
	table, _ := slide.AddTable(4, 3)
	table.SetPosition(50, 50)   // 50pt from left/top
	table.SetSize(500, 200)     // 500pt wide, 200pt tall

	// Set header row
	table.SetCellText(0, 0, "Product")
	table.SetCellText(0, 1, "Price")
	table.SetCellText(0, 2, "Quantity")

	// Set data rows
	table.SetCellText(1, 0, "Widget A")
	table.SetCellText(1, 1, "$10.00")
	table.SetCellText(1, 2, "100")

	table.SetCellText(2, 0, "Widget B")
	table.SetCellText(2, 1, "$15.00")
	table.SetCellText(2, 2, "75")

	table.SetCellText(3, 0, "Total")
	table.SetCellText(3, 1, "$25.00")
	table.SetCellText(3, 2, "175")

	// Apply styling
	table.SetStyle(presentation.TableStyleMedium2)
	table.SetBandedRows(true)
	table.SetFirstRowBold(true)
	table.SetLastRowBold(true)

	// Customize column widths
	table.SetColumnWidth(0, 200) // 200pt for product column
	table.SetColumnWidth(1, 150) // 150pt for price
	table.SetColumnWidth(2, 150) // 150pt for quantity

	// Save presentation
	pres.SaveToFile("product_table.pptx")
}
```

## Open Questions

1. **Cell Merge/Split**: Should we implement cell merge in initial release or defer to future?
   - **Recommendation**: Defer to future enhancement, focus on basic functionality first

2. **Table Auto-Resize**: Should tables auto-resize based on content?
   - **Recommendation**: No auto-resize initially, users set explicit sizes

3. **Default Styles**: Which table styles should be included as constants?
   - **Recommendation**: Include 5-10 most common Office table styles

4. **Border Customization**: How detailed should border control be?
   - **Recommendation**: Basic border rendering initially, advanced customization in future

## Performance Considerations

- Table creation: O(rows × cols) for cell initialization
- Row append: O(cols) for creating new cells
- Cell access: O(1) with direct indexing
- XML serialization: O(rows × cols) for full table
- PDF rendering: O(rows × cols) for grid + text rendering

All operations are linear and suitable for typical presentation tables (< 100 rows).

## Migration from Existing Code

No breaking changes. Users with existing GraphicFrame code can continue using it. New table API is purely additive.
