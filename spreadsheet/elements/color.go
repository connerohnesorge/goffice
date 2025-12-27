package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Color represents a color element (x:color) used in conditional formatting
// and styles. Colors can be specified as auto, indexed, RGB, or theme colors.
type Color struct {
	*openxml.LeafElementBase
}

// NewColor creates a new Color element.
func NewColor() *Color {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"color",
		PrefixDefault,
	)

	return &Color{LeafElementBase: elem}
}

// Auto returns whether the automatic color should be used. Attribute: auto.
func (c *Color) Auto() bool {
	attr, found := c.GetAttribute("auto", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetAuto sets whether the automatic color should be used. Attribute: auto.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Color) SetAuto(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"auto",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("auto", "")
	}
}

// Indexed returns the indexed color value (0-64). Attribute: indexed.
func (c *Color) Indexed() uint32 {
	attr, found := c.GetAttribute("indexed", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:mnd,revive // base 10
		32, //nolint:mnd,revive // 32-bit size
	)

	return uint32(val)
}

// SetIndexed sets the indexed color value. Attribute: indexed.
func (c *Color) SetIndexed(index uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"indexed",
			"",
			strconv.FormatUint(
				uint64(index),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// RGB returns the RGB color value (ARGB hex, e.g., "FF0000FF"). Attribute: rgb.
func (c *Color) RGB() string {
	attr, found := c.GetAttribute("rgb", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRGB sets the RGB color value (ARGB hex, e.g., "FF0000FF"). Attribute: rgb.
func (c *Color) SetRGB(rgb string) {
	if rgb == "" {
		c.RemoveAttribute("rgb", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"rgb",
			"",
			rgb,
		),
	)
}

// Theme returns the theme color index. Attribute: theme.
func (c *Color) Theme() uint32 {
	attr, found := c.GetAttribute("theme", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:mnd,revive // base 10
		32, //nolint:mnd,revive // 32-bit size
	)

	return uint32(val)
}

// SetTheme sets the theme color index. Attribute: theme.
func (c *Color) SetTheme(theme uint32) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"theme",
			"",
			strconv.FormatUint(
				uint64(theme),
				10, //nolint:revive // add-constant
			),
		),
	)
}

// Tint returns the tint value (-1.0 to 1.0). Attribute: tint.
func (c *Color) Tint() float64 {
	attr, found := c.GetAttribute("tint", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetTint sets the tint value (-1.0 to 1.0). Attribute: tint.
func (c *Color) SetTint(tint float64) {
	if tint == 0 {
		c.RemoveAttribute("tint", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"tint",
			"",
			strconv.FormatFloat(
				tint,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Clone creates a deep copy of this Color element.
func (c *Color) Clone() openxml.Element {
	cloned := c.LeafElementBase.Clone()

	return &Color{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Color element.
func (c *Color) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.LeafElementBase.CloneNode(deep)

	return &Color{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}
