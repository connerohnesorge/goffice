package table

import (
	"errors"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// NamespaceMain is the DrawingML namespace.
const NamespaceMain = drawingml.NamespaceMain

// PrefixMain is the DrawingML prefix.
const PrefixMain = drawingml.PrefixMain

// Common attribute values.
const (
	attrValueOne  = "1"
	attrValueTrue = "true"
)

// Table represents a DrawingML table (a:tbl).
// This is the core table structure used in presentations, spreadsheets, and diagrams.
type Table struct {
	*openxml.CompositeElementBase
}

// NewTable creates a new table with the specified number of rows and columns.
// All cells are initialized with empty text bodies.
func NewTable(rows, cols int) *Table {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tbl",
		PrefixMain,
	)
	tbl := &Table{CompositeElementBase: elem}

	// Add table properties
	tbl.AppendChild(NewTableProperties())

	// Add table grid with column definitions
	grid := NewTableGrid(cols)
	tbl.AppendChild(grid)

	// Add rows
	for range rows {
		row := NewTableRow(cols)
		tbl.AppendChild(row)
	}

	return tbl
}

// Properties returns the table properties.
func (t *Table) Properties() *TableProperties {
	elem := t.GetElement("tblPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if props, ok := elem.(*TableProperties); ok {
		return props
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableProperties{CompositeElementBase: comp}
	}

	return nil
}

// SetProperties sets the table properties.
func (t *Table) SetProperties(props *TableProperties) {
	if existing := t.GetElement("tblPr", NamespaceMain); existing != nil {
		t.RemoveChild(existing)
	}
	if props != nil {
		t.PrependChild(props)
	}
}

// Grid returns the table grid.
func (t *Table) Grid() *TableGrid {
	elem := t.GetElement("tblGrid", NamespaceMain)
	if elem == nil {
		return nil
	}
	if grid, ok := elem.(*TableGrid); ok {
		return grid
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableGrid{CompositeElementBase: comp}
	}

	return nil
}

// SetGrid sets the table grid.
func (t *Table) SetGrid(grid *TableGrid) {
	if existing := t.GetElement("tblGrid", NamespaceMain); existing != nil {
		t.RemoveChild(existing)
	}
	if grid == nil {
		return
	}
	// Insert after properties
	if props := t.Properties(); props != nil {
		t.InsertAfter(grid, props)
	} else {
		t.PrependChild(grid)
	}
}

// Rows returns all rows in the table.
func (t *Table) Rows() []*TableRow {
	var rows []*TableRow
	for child := range t.Children() {
		if child.LocalName() != "tr" || child.NamespaceURI() != NamespaceMain {
			continue
		}
		switch v := child.(type) {
		case *TableRow:
			rows = append(rows, v)
		case *openxml.CompositeElementBase:
			rows = append(rows, &TableRow{CompositeElementBase: v})
		}
	}

	return rows
}

// RowCount returns the number of rows in the table.
func (t *Table) RowCount() int {
	return len(t.Rows())
}

// ColumnCount returns the number of columns in the table.
func (t *Table) ColumnCount() int {
	grid := t.Grid()
	if grid == nil {
		return 0
	}

	return grid.ColumnCount()
}

// GetRow returns the row at the specified index (0-based).
func (t *Table) GetRow(index int) *TableRow {
	rows := t.Rows()
	if index < 0 || index >= len(rows) {
		return nil
	}

	return rows[index]
}

// AppendRow appends a new row with the correct number of columns.
func (t *Table) AppendRow() *TableRow {
	cols := t.ColumnCount()
	row := NewTableRow(cols)
	t.AppendChild(row)

	return row
}

// InsertRow inserts a new row at the specified index.
func (t *Table) InsertRow(index int) error {
	rows := t.Rows()
	if index < 0 || index > len(rows) {
		return errors.New("index out of range")
	}

	cols := t.ColumnCount()
	newRow := NewTableRow(cols)

	if index == len(rows) {
		// Append at end
		t.AppendChild(newRow)
	} else {
		// Insert before the row at index
		t.InsertBefore(newRow, rows[index])
	}

	return nil
}

// DeleteRow deletes the row at the specified index.
func (t *Table) DeleteRow(index int) error {
	rows := t.Rows()
	if index < 0 || index >= len(rows) {
		return errors.New("index out of range")
	}
	t.RemoveChild(rows[index])

	return nil
}

// GetCell returns the cell at the specified row and column (0-based).
func (t *Table) GetCell(row, col int) *TableCell {
	rowElem := t.GetRow(row)
	if rowElem == nil {
		return nil
	}

	return rowElem.GetCell(col)
}

// Clone creates a deep copy of this Table element.
func (t *Table) Clone() openxml.Element {
	return &Table{
		CompositeElementBase: t.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableProperties represents table-level styling (a:tblPr).
type TableProperties struct {
	*openxml.CompositeElementBase
}

// NewTableProperties creates a new table properties element.
func NewTableProperties() *TableProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tblPr",
		PrefixMain,
	)

	return &TableProperties{CompositeElementBase: elem}
}

// FirstRow returns whether the first row should be styled specially.
func (tp *TableProperties) FirstRow() bool {
	attr, found := tp.GetAttribute("firstRow", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetFirstRow sets whether the first row should be styled specially.
func (tp *TableProperties) SetFirstRow(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "firstRow", "", attrValueOne))
	} else {
		tp.RemoveAttribute("firstRow", "")
	}
}

// LastRow returns whether the last row should be styled specially.
func (tp *TableProperties) LastRow() bool {
	attr, found := tp.GetAttribute("lastRow", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetLastRow sets whether the last row should be styled specially.
func (tp *TableProperties) SetLastRow(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "lastRow", "", attrValueOne))
	} else {
		tp.RemoveAttribute("lastRow", "")
	}
}

// BandRow returns whether rows should be banded.
func (tp *TableProperties) BandRow() bool {
	attr, found := tp.GetAttribute("bandRow", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetBandRow sets whether rows should be banded.
func (tp *TableProperties) SetBandRow(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "bandRow", "", attrValueOne))
	} else {
		tp.RemoveAttribute("bandRow", "")
	}
}

// BandCol returns whether columns should be banded.
func (tp *TableProperties) BandCol() bool {
	attr, found := tp.GetAttribute("bandCol", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetBandCol sets whether columns should be banded.
func (tp *TableProperties) SetBandCol(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "bandCol", "", attrValueOne))
	} else {
		tp.RemoveAttribute("bandCol", "")
	}
}

// FirstCol returns whether the first column should be styled specially.
func (tp *TableProperties) FirstCol() bool {
	attr, found := tp.GetAttribute("firstCol", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetFirstCol sets whether the first column should be styled specially.
func (tp *TableProperties) SetFirstCol(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "firstCol", "", attrValueOne))
	} else {
		tp.RemoveAttribute("firstCol", "")
	}
}

// LastCol returns whether the last column should be styled specially.
func (tp *TableProperties) LastCol() bool {
	attr, found := tp.GetAttribute("lastCol", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne || attr.Value() == attrValueTrue
}

// SetLastCol sets whether the last column should be styled specially.
func (tp *TableProperties) SetLastCol(value bool) { //nolint:revive // flag-parameter is intentional for boolean setter
	if value {
		tp.SetAttribute(openxml.NewAttribute("", "lastCol", "", attrValueOne))
	} else {
		tp.RemoveAttribute("lastCol", "")
	}
}

// Clone creates a deep copy of this TableProperties element.
func (tp *TableProperties) Clone() openxml.Element {
	return &TableProperties{
		CompositeElementBase: tp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
