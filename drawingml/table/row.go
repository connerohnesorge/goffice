package table

import (
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for row operations.
const (
	// defaultRowHeightFraction is the fraction of an inch for default row height.
	defaultRowHeightFraction = 4
	// base10Row is the numeric base for parsing/formatting row values.
	base10Row = 10
)

// TableRow represents a table row (a:tr).
type TableRow struct {
	*openxml.CompositeElementBase
}

// NewTableRow creates a new table row with the specified number of cells.
// Each cell is initialized with an empty text body.
// The row is given a default height of 0.25 inches (317500 EMUs).
func NewTableRow(cols int) *TableRow {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tr",
		PrefixMain,
	)
	row := &TableRow{CompositeElementBase: elem}

	// Set default row height (0.25 inches = 228600 EMUs)
	defaultHeight := drawingml.EMUsPerInch / defaultRowHeightFraction
	row.SetHeight(int64(defaultHeight))

	// Add cells
	for range cols {
		cell := NewTableCell()
		row.AppendChild(cell)
	}

	return row
}

// Height returns the row height in EMUs.
func (tr *TableRow) Height() int64 {
	attr, found := tr.GetAttribute("h", "")
	if !found {
		return 0
	}
	height, _ := strconv.ParseInt(attr.Value(), base10Row, 64) //nolint:revive // 64-bit integer is standard

	return height
}

// SetHeight sets the row height in EMUs.
func (tr *TableRow) SetHeight(heightEMU int64) {
	tr.SetAttribute(
		openxml.NewAttribute(
			"",
			"h",
			"",
			strconv.FormatInt(heightEMU, base10Row),
		),
	)
}

// Cells returns all cells in this row.
func (tr *TableRow) Cells() []*TableCell {
	var cells []*TableCell
	for child := range tr.Children() {
		if child.LocalName() != "tc" || child.NamespaceURI() != NamespaceMain {
			continue
		}
		switch v := child.(type) {
		case *TableCell:
			cells = append(cells, v)
		case *openxml.CompositeElementBase:
			cells = append(cells, &TableCell{CompositeElementBase: v})
		}
	}

	return cells
}

// CellCount returns the number of cells in this row.
func (tr *TableRow) CellCount() int {
	return len(tr.Cells())
}

// GetCell returns the cell at the specified index (0-based).
func (tr *TableRow) GetCell(index int) *TableCell {
	cells := tr.Cells()
	if index < 0 || index >= len(cells) {
		return nil
	}

	return cells[index]
}

// AppendCell appends a new cell to this row.
func (tr *TableRow) AppendCell() *TableCell {
	cell := NewTableCell()
	tr.AppendChild(cell)

	return cell
}

// Clone creates a deep copy of this TableRow element.
func (tr *TableRow) Clone() openxml.Element {
	return &TableRow{
		CompositeElementBase: tr.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
