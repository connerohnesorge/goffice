package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

const (
	// numBase10 is the base for parsing decimal numbers.
	numBase10 = 10
	// numBits32 is the bit size for 32-bit unsigned integers.
	numBits32 = 32
)

// IconFilter represents the iconFilter element (x:iconFilter).
// It filters based on the icon shown in the cell from conditional formatting.
type IconFilter struct {
	*openxml.LeafElementBase
}

// NewIconFilter creates a new IconFilter element.
func NewIconFilter() *IconFilter {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"iconFilter",
		PrefixDefault,
	)

	return &IconFilter{LeafElementBase: elem}
}

// IconSet returns the icon set name. Attribute: iconSet.
// This corresponds to the IconSetType values used in conditional formatting.
func (i *IconFilter) IconSet() IconSetType {
	attr, found := i.GetAttribute("iconSet", "")
	if !found {
		return ""
	}

	return IconSetType(attr.Value())
}

// SetIconSet sets the icon set name. Attribute: iconSet.
func (i *IconFilter) SetIconSet(
	iconSet IconSetType,
) {
	if iconSet == "" {
		i.RemoveAttribute("iconSet", "")

		return
	}
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"iconSet",
			"",
			string(iconSet),
		),
	)
}

// IconId returns the zero-based icon index within the icon set.
// A value of 0 means to filter by cells with no icon.
// Attribute: iconId.
func (i *IconFilter) IconId() uint32 {
	attr, found := i.GetAttribute("iconId", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		numBase10,
		numBits32,
	)

	return uint32(val)
}

// SetIconId sets the zero-based icon index within the icon set.
// Attribute: iconId.
func (i *IconFilter) SetIconId(id uint32) {
	i.SetAttribute(
		openxml.NewAttribute(
			"",
			"iconId",
			"",
			strconv.FormatUint(
				uint64(id),
				numBase10,
			),
		),
	)
}

// Clone creates a deep copy of this IconFilter element.
func (i *IconFilter) Clone() openxml.Element {
	cloned := i.LeafElementBase.Clone()

	return &IconFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this IconFilter element.
func (i *IconFilter) CloneNode(
	deep bool,
) openxml.Element {
	cloned := i.LeafElementBase.CloneNode(deep)

	return &IconFilter{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
