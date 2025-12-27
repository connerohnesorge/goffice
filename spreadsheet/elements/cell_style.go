package elements

//revive:disable:file-length-limit many style properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Built-in cell style IDs.
const (
	// BuiltinStyleNormal is the Normal style (index 0).
	BuiltinStyleNormal uint32 = 0
	// BuiltinStyleComma is the Comma style.
	BuiltinStyleComma uint32 = 3
	// BuiltinStyleCurrency is the Currency style.
	BuiltinStyleCurrency uint32 = 4
	// BuiltinStylePercent is the Percent style.
	BuiltinStylePercent uint32 = 5
	// BuiltinStyleComma0 is the Comma [0] style.
	BuiltinStyleComma0 uint32 = 6
	// BuiltinStyleCurrency0 is the Currency [0] style.
	BuiltinStyleCurrency0 uint32 = 7
	// BuiltinStyleHyperlink is the Hyperlink style.
	BuiltinStyleHyperlink uint32 = 8
	// BuiltinStyleHyperlinkFollowed is the Followed Hyperlink style.
	BuiltinStyleHyperlinkFollowed uint32 = 9
	// BuiltinStyleNote is the Note style.
	BuiltinStyleNote uint32 = 10
	// BuiltinStyleWarningText is the Warning Text style.
	BuiltinStyleWarningText uint32 = 11
	// BuiltinStyleTitle is the Title style.
	BuiltinStyleTitle uint32 = 15
	// BuiltinStyleHeading1 is the Heading 1 style.
	BuiltinStyleHeading1 uint32 = 16
	// BuiltinStyleHeading2 is the Heading 2 style.
	BuiltinStyleHeading2 uint32 = 17
	// BuiltinStyleHeading3 is the Heading 3 style.
	BuiltinStyleHeading3 uint32 = 18
	// BuiltinStyleHeading4 is the Heading 4 style.
	BuiltinStyleHeading4 uint32 = 19
	// BuiltinStyleInput is the Input style.
	BuiltinStyleInput uint32 = 20
	// BuiltinStyleOutput is the Output style.
	BuiltinStyleOutput uint32 = 21
	// BuiltinStyleCalculation is the Calculation style.
	BuiltinStyleCalculation uint32 = 22
	// BuiltinStyleCheckCell is the Check Cell style.
	BuiltinStyleCheckCell uint32 = 23
	// BuiltinStyleLinkedCell is the Linked Cell style.
	BuiltinStyleLinkedCell uint32 = 24
	// BuiltinStyleTotal is the Total style.
	BuiltinStyleTotal uint32 = 25
	// BuiltinStyleGood is the Good style.
	BuiltinStyleGood uint32 = 26
	// BuiltinStyleBad is the Bad style.
	BuiltinStyleBad uint32 = 27
	// BuiltinStyleNeutral is the Neutral style.
	BuiltinStyleNeutral uint32 = 28
	// BuiltinStyleAccent1 is the Accent1 style.
	BuiltinStyleAccent1 uint32 = 29
	// BuiltinStyleAccent2 is the Accent2 style.
	BuiltinStyleAccent2 uint32 = 33
	// BuiltinStyleAccent3 is the Accent3 style.
	BuiltinStyleAccent3 uint32 = 37
	// BuiltinStyleAccent4 is the Accent4 style.
	BuiltinStyleAccent4 uint32 = 41
	// BuiltinStyleAccent5 is the Accent5 style.
	BuiltinStyleAccent5 uint32 = 45
	// BuiltinStyleAccent6 is the Accent6 style.
	BuiltinStyleAccent6 uint32 = 49
	// BuiltinStyleExplanatoryText is the Explanatory Text style.
	BuiltinStyleExplanatoryText uint32 = 53
)

// BuiltinStyleNames maps built-in style IDs to their names.
var BuiltinStyleNames = map[uint32]string{
	0:  "Normal",
	3:  "Comma",
	4:  "Currency",
	5:  "Percent",
	6:  "Comma [0]",
	7:  "Currency [0]",
	8:  "Hyperlink",
	9:  "Followed Hyperlink",
	10: "Note",
	11: "Warning Text",
	15: "Title",
	16: "Heading 1",
	17: "Heading 2",
	18: "Heading 3",
	19: "Heading 4",
	20: "Input",
	21: "Output",
	22: "Calculation",
	23: "Check Cell",
	24: "Linked Cell",
	25: "Total",
	26: "Good",
	27: "Bad",
	28: "Neutral",
	29: "Accent1",
	33: "Accent2",
	37: "Accent3",
	41: "Accent4",
	45: "Accent5",
	49: "Accent6",
	53: "Explanatory Text",
}

// CellStyles represents the cell styles container element (x:cellStyles).
type CellStyles struct {
	*openxml.CompositeElementBase
}

// NewCellStyles creates a new CellStyles element.
func NewCellStyles() *CellStyles {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cellStyles",
		PrefixDefault,
	)

	return &CellStyles{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (c *CellStyles) Count() uint32 {
	attr, found := c.GetAttribute("count", "")
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

// SetCount sets the count attribute.
func (c *CellStyles) SetCount(count uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"count",
			"",
			strconv.FormatUint(
				uint64(count),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// CellStyles returns an iterator over all CellStyle elements.
func (c *CellStyles) CellStyles() iter.Seq[*CellStyle] {
	return func(yield func(*CellStyle) bool) {
		for child := range c.Children() {
			if child.LocalName() != "cellStyle" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var style *CellStyle
			if s, ok := child.(*CellStyle); ok {
				style = s
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				style = &CellStyle{LeafElementBase: leaf}
			}
			if style != nil && !yield(style) {
				return
			}
		}
	}
}

// GetCellStyle returns the cell style with the given name, or nil if not found.
func (c *CellStyles) GetCellStyle(
	name string,
) *CellStyle {
	for style := range c.CellStyles() {
		if style.Name() == name {
			return style
		}
	}

	return nil
}

// GetCellStyleByXfId returns the cell style with the given xfId, or nil if
// not found.
func (c *CellStyles) GetCellStyleByXfId(
	xfId uint32,
) *CellStyle {
	for style := range c.CellStyles() {
		if style.XfId() == xfId {
			return style
		}
	}

	return nil
}

// AddCellStyle adds a new cell style and returns it.
func (c *CellStyles) AddCellStyle() *CellStyle {
	style := NewCellStyle()
	c.AppendChild(style)
	c.SetCount(c.Count() + 1)

	return style
}

// ItemCount returns the actual number of CellStyle children.
func (c *CellStyles) ItemCount() int {
	count := 0
	for range c.CellStyles() {
		count++
	}

	return count
}

// Clone creates a deep copy of this CellStyles element.
func (c *CellStyles) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CellStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CellStyles element.
func (c *CellStyles) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CellStyles{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CellStyle represents a cell style element (x:cellStyle).
// This is a named style that references a format in cellStyleXfs.
type CellStyle struct {
	*openxml.LeafElementBase
}

// NewCellStyle creates a new CellStyle element.
func NewCellStyle() *CellStyle {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"cellStyle",
		PrefixDefault,
	)

	return &CellStyle{LeafElementBase: elem}
}

// Name returns the style name.
func (c *CellStyle) Name() string {
	attr, found := c.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the style name.
func (c *CellStyle) SetName(name string) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			name,
		),
	)
}

// XfId returns the index into cellStyleXfs.
func (c *CellStyle) XfId() uint32 {
	attr, found := c.GetAttribute("xfId", "")
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

// SetXfId sets the index into cellStyleXfs.
func (c *CellStyle) SetXfId(id uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"xfId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// BuiltinId returns the built-in style ID, or 0 if not a built-in style.
func (c *CellStyle) BuiltinId() uint32 {
	attr, found := c.GetAttribute("builtinId", "")
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

// SetBuiltinId sets the built-in style ID.
func (c *CellStyle) SetBuiltinId(id uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"builtinId",
			"",
			strconv.FormatUint(
				uint64(id),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// HasBuiltinId returns whether this style has a built-in ID.
func (c *CellStyle) HasBuiltinId() bool {
	_, found := c.GetAttribute("builtinId", "")

	return found
}

// ILevel returns the outline level for outline styles.
func (c *CellStyle) ILevel() uint32 {
	attr, found := c.GetAttribute("iLevel", "")
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

// SetILevel sets the outline level.
func (c *CellStyle) SetILevel(level uint32) {
	if level == 0 {
		c.RemoveAttribute("iLevel", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"iLevel",
			"",
			strconv.FormatUint(
				uint64(level),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Hidden returns whether the style is hidden.
func (c *CellStyle) Hidden() bool {
	attr, found := c.GetAttribute("hidden", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidden sets whether the style is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CellStyle) SetHidden(hidden bool) {
	if hidden {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"hidden",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("hidden", "")
	}
}

// CustomBuiltin returns whether this is a custom version of a built-in style.
func (c *CellStyle) CustomBuiltin() bool {
	attr, found := c.GetAttribute(
		"customBuiltin",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCustomBuiltin sets whether this is a custom version of a built-in style.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CellStyle) SetCustomBuiltin(
	custom bool,
) {
	if custom {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"customBuiltin",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("customBuiltin", "")
	}
}

// Clone creates a deep copy of this CellStyle element.
func (c *CellStyle) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &CellStyle{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this CellStyle element.
func (c *CellStyle) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &CellStyle{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
