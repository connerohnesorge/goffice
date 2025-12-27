//nolint:revive // comments-density
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FromMarker represents the from marker element (xdr:from).
// This specifies the starting cell position for an anchor.
type FromMarker struct {
	*openxml.CompositeElementBase
}

// NewFromMarker creates a new FromMarker element.
func NewFromMarker() *FromMarker {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"from",
		PrefixXDR,
	)

	return &FromMarker{CompositeElementBase: elem}
}

// Col returns the column index (0-based).
func (fm *FromMarker) Col() int {
	elem := fm.GetElement(
		elemNameCol,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.Atoi(leaf.InnerText())

		return val
	}

	return 0
}

// SetCol sets the column index (0-based).
func (fm *FromMarker) SetCol(col int) {
	elem := fm.GetElement(
		elemNameCol,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		elem = openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameCol,
			PrefixXDR,
			strconv.Itoa(col),
		)
		if first := fm.FirstChild(); first != nil {
			fm.InsertBefore(elem, first)
		} else {
			fm.AppendChild(elem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(strconv.Itoa(col))
	}
}

// ColOff returns the column offset in EMUs.
func (fm *FromMarker) ColOff() int64 {
	elem := fm.GetElement(
		elemNameColOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(
			leaf.InnerText(),
			parseBase10,
			bitSize64,
		)

		return val
	}

	return 0
}

// SetColOff sets the column offset in EMUs.
func (fm *FromMarker) SetColOff(offset int64) {
	elem := fm.GetElement(
		elemNameColOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		colElem := fm.GetElement(
			elemNameCol,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameColOff,
			PrefixXDR,
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
		if colElem != nil {
			fm.InsertAfter(newElem, colElem)
		} else {
			fm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
	}
}

// Row returns the row index (0-based).
func (fm *FromMarker) Row() int {
	elem := fm.GetElement(
		elemNameRow,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.Atoi(leaf.InnerText())

		return val
	}

	return 0
}

// SetRow sets the row index (0-based).
func (fm *FromMarker) SetRow(row int) {
	elem := fm.GetElement(
		elemNameRow,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		colOffElem := fm.GetElement(
			elemNameColOff,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameRow,
			PrefixXDR,
			strconv.Itoa(row),
		)
		if colOffElem != nil {
			fm.InsertAfter(newElem, colOffElem)
		} else {
			fm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(strconv.Itoa(row))
	}
}

// RowOff returns the row offset in EMUs.
func (fm *FromMarker) RowOff() int64 {
	elem := fm.GetElement(
		elemNameRowOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(
			leaf.InnerText(),
			parseBase10,
			bitSize64,
		)

		return val
	}

	return 0
}

// SetRowOff sets the row offset in EMUs.
func (fm *FromMarker) SetRowOff(offset int64) {
	elem := fm.GetElement(
		elemNameRowOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		rowElem := fm.GetElement(
			elemNameRow,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameRowOff,
			PrefixXDR,
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
		if rowElem != nil {
			fm.InsertAfter(newElem, rowElem)
		} else {
			fm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
	}
}

// SetPosition sets all position components at once.
func (fm *FromMarker) SetPosition(
	col int,
	colOff int64,
	row int,
	rowOff int64,
) {
	fm.SetCol(col)
	fm.SetColOff(colOff)
	fm.SetRow(row)
	fm.SetRowOff(rowOff)
}

// Clone creates a deep copy of this FromMarker element.
func (fm *FromMarker) Clone() openxml.Element {
	cloned := fm.CompositeElementBase.Clone()

	return &FromMarker{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FromMarker element.
func (fm *FromMarker) CloneNode(
	deep bool,
) openxml.Element {
	cloned := fm.CompositeElementBase.CloneNode(
		deep,
	)

	return &FromMarker{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ToMarker represents the to marker element (xdr:to).
// This specifies the ending cell position for a two-cell anchor.
type ToMarker struct {
	*openxml.CompositeElementBase
}

// NewToMarker creates a new ToMarker element.
func NewToMarker() *ToMarker {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"to",
		PrefixXDR,
	)

	return &ToMarker{CompositeElementBase: elem}
}

// Col returns the column index (0-based).
func (tm *ToMarker) Col() int {
	elem := tm.GetElement(
		elemNameCol,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.Atoi(leaf.InnerText())

		return val
	}

	return 0
}

// SetCol sets the column index (0-based).
func (tm *ToMarker) SetCol(col int) {
	elem := tm.GetElement(
		elemNameCol,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		elem = openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameCol,
			PrefixXDR,
			strconv.Itoa(col),
		)
		if first := tm.FirstChild(); first != nil {
			tm.InsertBefore(elem, first)
		} else {
			tm.AppendChild(elem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(strconv.Itoa(col))
	}
}

// ColOff returns the column offset in EMUs.
func (tm *ToMarker) ColOff() int64 {
	elem := tm.GetElement(
		elemNameColOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(
			leaf.InnerText(),
			parseBase10,
			bitSize64,
		)

		return val
	}

	return 0
}

// SetColOff sets the column offset in EMUs.
func (tm *ToMarker) SetColOff(offset int64) {
	elem := tm.GetElement(
		elemNameColOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		colElem := tm.GetElement(
			elemNameCol,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameColOff,
			PrefixXDR,
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
		if colElem != nil {
			tm.InsertAfter(newElem, colElem)
		} else {
			tm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
	}
}

// Row returns the row index (0-based).
func (tm *ToMarker) Row() int {
	elem := tm.GetElement(
		elemNameRow,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.Atoi(leaf.InnerText())

		return val
	}

	return 0
}

// SetRow sets the row index (0-based).
func (tm *ToMarker) SetRow(row int) {
	elem := tm.GetElement(
		elemNameRow,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		colOffElem := tm.GetElement(
			elemNameColOff,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameRow,
			PrefixXDR,
			strconv.Itoa(row),
		)
		if colOffElem != nil {
			tm.InsertAfter(newElem, colOffElem)
		} else {
			tm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(strconv.Itoa(row))
	}
}

// RowOff returns the row offset in EMUs.
func (tm *ToMarker) RowOff() int64 {
	elem := tm.GetElement(
		elemNameRowOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(
			leaf.InnerText(),
			parseBase10,
			bitSize64,
		)

		return val
	}

	return 0
}

// SetRowOff sets the row offset in EMUs.
func (tm *ToMarker) SetRowOff(offset int64) {
	elem := tm.GetElement(
		elemNameRowOff,
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		rowElem := tm.GetElement(
			elemNameRow,
			NamespaceSpreadsheetDrawing,
		)
		newElem := openxml.NewLeafElementWithText(
			NamespaceSpreadsheetDrawing,
			elemNameRowOff,
			PrefixXDR,
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
		if rowElem != nil {
			tm.InsertAfter(newElem, rowElem)
		} else {
			tm.AppendChild(newElem)
		}

		return
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(
			strconv.FormatInt(
				offset,
				parseBase10,
			),
		)
	}
}

// SetPosition sets all position components at once.
func (tm *ToMarker) SetPosition(
	col int,
	colOff int64,
	row int,
	rowOff int64,
) {
	tm.SetCol(col)
	tm.SetColOff(colOff)
	tm.SetRow(row)
	tm.SetRowOff(rowOff)
}

// Clone creates a deep copy of this ToMarker element.
func (tm *ToMarker) Clone() openxml.Element {
	cloned := tm.CompositeElementBase.Clone()

	return &ToMarker{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ToMarker element.
func (tm *ToMarker) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tm.CompositeElementBase.CloneNode(
		deep,
	)

	return &ToMarker{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
