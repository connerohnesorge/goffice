package presentation

import (
	"errors"
	"fmt"

	"github.com/connerohnesorge/goffice/drawingml"
	drawtable "github.com/connerohnesorge/goffice/drawingml/table"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// Table-related constants.
const (
	// Error messages.
	errTableNil = "table is nil"

	// Default table positioning in points.
	defaultTablePositionPt = 72  // 1 inch = 72 points
	defaultTableWidthPt    = 432 // 6 inches = 432 points
	defaultTableHeightPt   = 288 // 4 inches = 288 points
)

// Table represents a high-level table in a presentation slide.
// It wraps a GraphicFrame containing a DrawingML table.
type Table struct {
	graphicFrame *elements.GraphicFrame
	table        *drawtable.Table
}

// NewTableFromGraphicFrame wraps an existing GraphicFrame that contains a table.
// This is used for reading existing tables from documents.
func NewTableFromGraphicFrame(gf *elements.GraphicFrame) *Table {
	if gf == nil {
		return nil
	}

	// Try to find the table element in the graphic frame
	// Structure is: graphicFrame > graphic > graphicData > tbl
	graphicElem := gf.GetElement("graphic", elements.NamespaceDrawingML)
	if graphicElem == nil {
		return nil
	}

	// Cast to CompositeElement to access GetElement
	var graphicDataElem openxml.Element
	if compElem, ok := graphicElem.(openxml.CompositeElement); ok {
		graphicDataElem = compElem.GetElement("graphicData", elements.NamespaceDrawingML)
	}
	if graphicDataElem == nil {
		return nil
	}

	// Get the table element
	var tblElem openxml.Element
	if compElem, ok := graphicDataElem.(openxml.CompositeElement); ok {
		tblElem = compElem.GetElement("tbl", drawtable.NamespaceMain)
	}
	if tblElem == nil {
		return nil
	}

	// Wrap as drawtable.Table
	var tbl *drawtable.Table
	switch v := tblElem.(type) {
	case *drawtable.Table:
		tbl = v
	default:
		// If it's not already a *drawtable.Table, we can't use it
		return nil
	}

	return &Table{
		graphicFrame: gf,
		table:        tbl,
	}
}

// GraphicFrame returns the underlying GraphicFrame element.
func (t *Table) GraphicFrame() *elements.GraphicFrame {
	return t.graphicFrame
}

// DrawingMLTable returns the underlying DrawingML table element.
func (t *Table) DrawingMLTable() *drawtable.Table {
	return t.table
}

// RowCount returns the number of rows in the table.
func (t *Table) RowCount() int {
	if t.table == nil {
		return 0
	}

	return t.table.RowCount()
}

// ColumnCount returns the number of columns in the table.
func (t *Table) ColumnCount() int {
	if t.table == nil {
		return 0
	}

	return t.table.ColumnCount()
}

// AppendRow appends a new row to the table.
func (t *Table) AppendRow() *Table {
	if t.table != nil {
		t.table.AppendRow()
	}

	return t
}

// InsertRow inserts a new row at the specified index.
func (t *Table) InsertRow(index int) error {
	if t.table == nil {
		return errors.New(errTableNil)
	}

	return t.table.InsertRow(index)
}

// DeleteRow deletes the row at the specified index.
func (t *Table) DeleteRow(index int) error {
	if t.table == nil {
		return errors.New(errTableNil)
	}

	return t.table.DeleteRow(index)
}

// SetColumnWidth sets the width of the specified column in points.
func (t *Table) SetColumnWidth(col int, widthPt float64) error {
	if t.table == nil {
		return errors.New(errTableNil)
	}

	grid := t.table.Grid()
	if grid == nil {
		return errors.New("table grid is nil")
	}

	// Convert points to EMUs
	widthEMU := drawingml.PointsToEmu(widthPt)

	return grid.SetColumnWidth(col, int64(widthEMU))
}

// GetCell returns the cell at the specified row and column (0-based).
func (t *Table) GetCell(row, col int) (*TableCell, error) {
	if t.table == nil {
		return nil, errors.New(errTableNil)
	}

	cell := t.table.GetCell(row, col)
	if cell == nil {
		return nil, fmt.Errorf("cell at (%d, %d) not found", row, col)
	}

	return &TableCell{cell: cell}, nil
}

// SetCellText sets the text of the cell at the specified row and column.
func (t *Table) SetCellText(row, col int, text string) error {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return err
	}
	cell.SetText(text)

	return nil
}

// GetCellText returns the text of the cell at the specified row and column.
func (t *Table) GetCellText(row, col int) (string, error) {
	cell, err := t.GetCell(row, col)
	if err != nil {
		return "", err
	}

	return cell.GetText(), nil
}

// SetPosition sets the position of the table in points.
func (t *Table) SetPosition(xPt, yPt float64) {
	if t.graphicFrame == nil {
		return
	}

	// Convert points to EMUs
	xEMU := int(drawingml.PointsToEmu(xPt))
	yEMU := int(drawingml.PointsToEmu(yPt))

	t.graphicFrame.SetPosition(xEMU, yEMU)
}

// SetSize sets the size of the table in points.
func (t *Table) SetSize(widthPt, heightPt float64) {
	if t.graphicFrame == nil {
		return
	}

	// Convert points to EMUs
	widthEMU := int(drawingml.PointsToEmu(widthPt))
	heightEMU := int(drawingml.PointsToEmu(heightPt))

	t.graphicFrame.SetSize(widthEMU, heightEMU)
}

// SetFirstRow sets whether the first row should be styled specially.
func (t *Table) SetFirstRow(value bool) {
	if t.table == nil {
		return
	}
	if props := t.table.Properties(); props != nil {
		props.SetFirstRow(value)
	}
}

// SetBandedRows sets whether rows should be banded.
func (t *Table) SetBandedRows(value bool) {
	if t.table == nil {
		return
	}
	if props := t.table.Properties(); props != nil {
		props.SetBandRow(value)
	}
}

// SetBandedColumns sets whether columns should be banded.
func (t *Table) SetBandedColumns(value bool) {
	if t.table == nil {
		return
	}
	if props := t.table.Properties(); props != nil {
		props.SetBandCol(value)
	}
}

// TableCell represents a cell in a presentation table.
type TableCell struct {
	cell *drawtable.TableCell
}

// SetText sets the text content of the cell.
func (tc *TableCell) SetText(text string) {
	if tc.cell != nil {
		tc.cell.SetText(text)
	}
}

// GetText returns the text content of the cell.
func (tc *TableCell) GetText() string {
	if tc.cell == nil {
		return ""
	}

	return tc.cell.GetText()
}

// TextBody returns the underlying DrawingML TextBody for advanced formatting.
func (tc *TableCell) TextBody() *drawingml.TextBody {
	if tc.cell == nil {
		return nil
	}

	return tc.cell.TextBody()
}

// Properties returns the cell properties for advanced styling.
func (tc *TableCell) Properties() *drawtable.CellProperties {
	if tc.cell == nil {
		return nil
	}

	return tc.cell.Properties()
}

// NewTable creates a new table and adds it to the slide with the specified number of rows and columns.
// This is a convenience function that can be called from user code.
func NewTable(slide *elements.Slide, rows, cols int) *Table {
	if slide == nil {
		return nil
	}

	// Get or create the shape tree
	csd := slide.GetOrCreateCommonSlideData()
	if csd == nil {
		return nil
	}
	st := csd.GetOrCreateShapeTree()
	if st == nil {
		return nil
	}

	// Create GraphicFrame
	gf := st.AddGraphicFrame()

	// Create DrawingML table
	tbl := drawtable.NewTable(rows, cols)

	// Link table to graphic frame
	// TODO: Implement LinkGraphicFrameToTable
	_ = tbl
	_ = gf

	// Set default position and size
	// Default position: 1 inch from top-left
	// Default size: 6 inches wide, 4 inches tall
	xEMU := int(drawingml.PointsToEmu(defaultTablePositionPt))
	yEMU := int(drawingml.PointsToEmu(defaultTablePositionPt))
	wEMU := int(drawingml.PointsToEmu(defaultTableWidthPt))
	hEMU := int(drawingml.PointsToEmu(defaultTableHeightPt))

	gf.SetPosition(xEMU, yEMU)
	gf.SetSize(wEMU, hEMU)

	return &Table{
		graphicFrame: gf,
		table:        tbl,
	}
}
