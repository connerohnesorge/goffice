package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ColorFilter represents the colorFilter element (x:colorFilter).
// It filters based on cell background or font color.
type ColorFilter struct {
	*openxml.LeafElementBase
}

// NewColorFilter creates a new ColorFilter element.
func NewColorFilter() *ColorFilter {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"colorFilter",
		PrefixDefault,
	)

	return &ColorFilter{LeafElementBase: elem}
}

// DxfId returns the differential formatting ID that specifies the color.
// Attribute: dxfId.
func (cf *ColorFilter) DxfId() uint32 {
	attr, found := cf.GetAttribute("dxfId", "")
	if !found {
		return 0
	}
	//nolint:revive // add-constant: 10 and 32 are standard parse parameters
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetDxfId sets the differential formatting ID that specifies the color.
// Attribute: dxfId.
func (cf *ColorFilter) SetDxfId(id uint32) {
	cf.SetAttribute(
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

// CellColor returns whether to filter by cell background color (true)
// or font color (false). Attribute: cellColor.
// Default is true (cell background color).
func (cf *ColorFilter) CellColor() bool {
	attr, found := cf.GetAttribute(
		"cellColor",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCellColor sets whether to filter by cell background color (true)
// or font color (false). Attribute: cellColor.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cf *ColorFilter) SetCellColor(value bool) {
	if value {
		cf.RemoveAttribute(
			"cellColor",
			"",
		) // Default is true
	} else {
		cf.SetAttribute(
			openxml.NewAttribute(
				"",
				"cellColor",
				"",
				attrValueFalse,
			),
		)
	}
}

// Clone creates a deep copy of this ColorFilter element.
func (cf *ColorFilter) Clone() openxml.Element {
	cloned := cf.LeafElementBase.Clone()

	return &ColorFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this ColorFilter element.
func (cf *ColorFilter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cf.LeafElementBase.CloneNode(deep)

	return &ColorFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
