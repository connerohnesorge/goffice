package table

import (
	"errors"
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// base10 is the numeric base for parsing/formatting EMU values.
const base10 = 10

// TableGrid represents the table grid (a:tblGrid) which defines column widths.
type TableGrid struct {
	*openxml.CompositeElementBase
}

// NewTableGrid creates a new table grid with the specified number of columns.
// Each column is given a default width of 1 inch (914400 EMUs).
func NewTableGrid(cols int) *TableGrid {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tblGrid",
		PrefixMain,
	)
	grid := &TableGrid{CompositeElementBase: elem}

	// Add columns with default width
	defaultWidth := drawingml.EMUsPerInch // 1 inch = 914400 EMUs
	for range cols {
		col := NewGridColumn(int64(defaultWidth))
		grid.AppendChild(col)
	}

	return grid
}

// Columns returns all grid columns.
func (tg *TableGrid) Columns() []*GridColumn {
	var columns []*GridColumn
	for child := range tg.Children() {
		if child.LocalName() != "gridCol" || child.NamespaceURI() != NamespaceMain {
			continue
		}
		switch v := child.(type) {
		case *GridColumn:
			columns = append(columns, v)
		case *openxml.CompositeElementBase:
			columns = append(columns, &GridColumn{CompositeElementBase: v})
		}
	}

	return columns
}

// ColumnCount returns the number of columns.
func (tg *TableGrid) ColumnCount() int {
	return len(tg.Columns())
}

// GetColumn returns the column at the specified index (0-based).
func (tg *TableGrid) GetColumn(index int) *GridColumn {
	cols := tg.Columns()
	if index < 0 || index >= len(cols) {
		return nil
	}

	return cols[index]
}

// SetColumnWidth sets the width of the specified column in EMUs.
func (tg *TableGrid) SetColumnWidth(index int, widthEMU int64) error {
	col := tg.GetColumn(index)
	if col == nil {
		return errors.New("column index out of range")
	}
	col.SetWidth(widthEMU)

	return nil
}

// Clone creates a deep copy of this TableGrid element.
func (tg *TableGrid) Clone() openxml.Element {
	return &TableGrid{
		CompositeElementBase: tg.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GridColumn represents a single column in the table grid (a:gridCol).
type GridColumn struct {
	*openxml.CompositeElementBase
}

// NewGridColumn creates a new grid column with the specified width in EMUs.
func NewGridColumn(widthEMU int64) *GridColumn {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"gridCol",
		PrefixMain,
	)
	col := &GridColumn{CompositeElementBase: elem}
	col.SetWidth(widthEMU)

	return col
}

// Width returns the column width in EMUs.
func (gc *GridColumn) Width() int64 {
	attr, found := gc.GetAttribute("w", "")
	if !found {
		return 0
	}
	width, _ := strconv.ParseInt(attr.Value(), base10, 64) //nolint:revive // 64-bit integer is standard

	return width
}

// SetWidth sets the column width in EMUs.
func (gc *GridColumn) SetWidth(widthEMU int64) {
	gc.SetAttribute(
		openxml.NewAttribute(
			"",
			"w",
			"",
			strconv.FormatInt(widthEMU, base10),
		),
	)
}

// Clone creates a deep copy of this GridColumn element.
func (gc *GridColumn) Clone() openxml.Element {
	return &GridColumn{
		CompositeElementBase: gc.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
