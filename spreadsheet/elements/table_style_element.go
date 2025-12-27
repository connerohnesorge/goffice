package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// TableStyleElement represents a table style element (x:tableStyleElement).
type TableStyleElement struct {
	*openxml.LeafElementBase
}

// NewTableStyleElement creates a new TableStyleElement element.
func NewTableStyleElement() *TableStyleElement {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"tableStyleElement",
		PrefixDefault,
	)

	return &TableStyleElement{
		LeafElementBase: elem,
	}
}

// Type returns the element type.
func (t *TableStyleElement) Type() TableStyleElementType {
	attr, found := t.GetAttribute("type", "")
	if !found {
		return TableStyleElementWholeTable
	}

	return TableStyleElementType(attr.Value())
}

// SetType sets the element type.
func (t *TableStyleElement) SetType(
	elementType TableStyleElementType,
) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(elementType),
		),
	)
}

// Size returns the size (number of rows/columns for stripes).
func (t *TableStyleElement) Size() uint32 {
	attr, found := t.GetAttribute("size", "")
	if !found {
		return 1 // Default is 1
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		parseBase10,
		bitSize32,
	)

	return uint32(val)
}

// SetSize sets the size.
func (t *TableStyleElement) SetSize(size uint32) {
	if size == 1 {
		t.RemoveAttribute(
			"size",
			"",
		) // 1 is default

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"size",
			"",
			strconv.FormatUint(
				uint64(size),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// DxfId returns the differential format ID.
func (t *TableStyleElement) DxfId() uint32 {
	attr, found := t.GetAttribute("dxfId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant: base 10
		32, //nolint:revive // add-constant: bit size 32
	)

	return uint32(val)
}

// SetDxfId sets the differential format ID.
func (t *TableStyleElement) SetDxfId(id uint32) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"dxfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this TableStyleElement element.
func (t *TableStyleElement) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &TableStyleElement{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TableStyleElement element.
func (t *TableStyleElement) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.LeafElementBase.CloneNode(deep)

	return &TableStyleElement{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
