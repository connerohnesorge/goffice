package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// SheetData represents the sheet data container element (x:sheetData).
// It contains all rows and cells in a worksheet.
type SheetData struct {
	*openxml.CompositeElementBase
}

// NewSheetData creates a new SheetData element.
func NewSheetData() *SheetData {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetData",
		PrefixDefault,
	)

	return &SheetData{CompositeElementBase: elem}
}

// Rows returns an iterator over all Row elements.
func (sd *SheetData) Rows() iter.Seq[*Row] {
	return func(yield func(*Row) bool) {
		for child := range sd.Children() {
			if child.LocalName() != "row" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var row *Row
			if r, ok := child.(*Row); ok {
				row = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				row = &Row{CompositeElementBase: comp}
			}
			if row != nil && !yield(row) {
				return
			}
		}
	}
}

// RowCount returns the number of rows.
func (sd *SheetData) RowCount() int {
	count := 0
	for range sd.Rows() {
		count++
	}

	return count
}

// GetRow returns the row at the given row index (1-based), or nil if not found.
func (sd *SheetData) GetRow(
	rowIndex uint32,
) *Row {
	for row := range sd.Rows() {
		if row.RowIndex() == rowIndex {
			return row
		}
	}

	return nil
}

// GetOrCreateRow returns the row at the given row index, creating it if it
// does not exist. Rows are kept in sorted order by index.
func (sd *SheetData) GetOrCreateRow(
	rowIndex uint32,
) *Row {
	// Try to find existing row
	for row := range sd.Rows() {
		if row.RowIndex() == rowIndex {
			return row
		}
	}

	// Create new row and insert in correct position
	return sd.AddRow(rowIndex)
}

// AddRow adds a new row at the given row index (1-based).
// The row is inserted in the correct position to maintain sorted order.
func (sd *SheetData) AddRow(
	rowIndex uint32,
) *Row {
	row := NewRow()
	row.SetRowIndex(rowIndex)

	// Find the correct position to insert the row
	var insertBefore openxml.Element
	for child := range sd.Children() {
		if child.LocalName() != "row" ||
			child.NamespaceURI() != NamespaceSML {
			continue
		}
		var existingRow *Row
		if r, ok := child.(*Row); ok {
			existingRow = r
		} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
			existingRow = &Row{CompositeElementBase: comp}
		}
		if existingRow != nil &&
			existingRow.RowIndex() > rowIndex {
			insertBefore = child

			break
		}
	}

	if insertBefore != nil {
		sd.InsertBefore(row, insertBefore)
	} else {
		sd.AppendChild(row)
	}

	return row
}

// RemoveRow removes the row at the given index.
func (sd *SheetData) RemoveRow(
	rowIndex uint32,
) bool {
	row := sd.GetRow(rowIndex)
	if row == nil {
		return false
	}

	return sd.RemoveChild(row)
}

// Clone creates a deep copy of this SheetData element.
func (sd *SheetData) Clone() openxml.Element {
	cloned := sd.CompositeElementBase.Clone()

	return &SheetData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetData element.
func (sd *SheetData) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sd.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
