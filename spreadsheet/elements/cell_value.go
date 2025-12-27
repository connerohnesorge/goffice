package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// CellValue represents the cell value element (x:v).
// The value is stored as text content:
// - Numbers are stored as decimal strings
// - Shared string indices are stored as integers
// - Booleans are stored as "0" or "1"
// - Errors are stored as error codes (e.g., "#DIV/0!")
type CellValue struct {
	*openxml.LeafElementBase
}

// NewCellValue creates a new CellValue element.
func NewCellValue() *CellValue {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"v",
		PrefixDefault,
	)

	return &CellValue{LeafElementBase: elem}
}

// NewCellValueWithText creates a new CellValue element with initial text.
func NewCellValueWithText(
	text string,
) *CellValue {
	elem := openxml.NewLeafElementWithText(
		NamespaceSML,
		"v",
		PrefixDefault,
		text,
	)

	return &CellValue{LeafElementBase: elem}
}

// Value returns the value as a string.
func (cv *CellValue) Value() string {
	return cv.InnerText()
}

// SetValue sets the value as a string.
func (cv *CellValue) SetValue(value string) {
	cv.SetInnerText(value)
}

// Clone creates a deep copy of this CellValue element.
func (cv *CellValue) Clone() openxml.Element {
	cloned := cv.LeafElementBase.Clone()

	return &CellValue{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CellValue element.
func (cv *CellValue) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cv.LeafElementBase.CloneNode(deep)

	return &CellValue{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
