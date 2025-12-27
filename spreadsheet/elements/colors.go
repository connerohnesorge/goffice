package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Theme color indices.
const (
	// ThemeColorLight1 is typically white.
	ThemeColorLight1 uint32 = 0
	// ThemeColorDark1 is typically black.
	ThemeColorDark1 uint32 = 1
	// ThemeColorLight2 is the secondary light color.
	ThemeColorLight2 uint32 = 2
	// ThemeColorDark2 is the secondary dark color.
	ThemeColorDark2 uint32 = 3
	// ThemeColorAccent1 is the first accent color.
	ThemeColorAccent1 uint32 = 4
	// ThemeColorAccent2 is the second accent color.
	ThemeColorAccent2 uint32 = 5
	// ThemeColorAccent3 is the third accent color.
	ThemeColorAccent3 uint32 = 6
	// ThemeColorAccent4 is the fourth accent color.
	ThemeColorAccent4 uint32 = 7
	// ThemeColorAccent5 is the fifth accent color.
	ThemeColorAccent5 uint32 = 8
	// ThemeColorAccent6 is the sixth accent color.
	ThemeColorAccent6 uint32 = 9
)

// Standard indexed colors (legacy Excel color palette).
var StandardIndexedColors = []string{
	"00000000", // 0 - Black
	"00FFFFFF", // 1 - White
	"00FF0000", // 2 - Red
	"0000FF00", // 3 - Bright Green
	"000000FF", // 4 - Blue
	"00FFFF00", // 5 - Yellow
	"00FF00FF", // 6 - Pink
	"0000FFFF", // 7 - Turquoise
	"00000000", // 8 - Black
	"00FFFFFF", // 9 - White
	"00FF0000", // 10 - Red
	"0000FF00", // 11 - Bright Green
	"000000FF", // 12 - Blue
	"00FFFF00", // 13 - Yellow
	"00FF00FF", // 14 - Pink
	"0000FFFF", // 15 - Turquoise
	"00800000", // 16 - Dark Red
	"00008000", // 17 - Green
	"00000080", // 18 - Dark Blue
	"00808000", // 19 - Olive
	"00800080", // 20 - Purple
	"00008080", // 21 - Teal
	"00C0C0C0", // 22 - Silver
	"00808080", // 23 - Gray
	"009999FF", // 24 - Periwinkle
	"00993366", // 25 - Plum
	"00FFFFCC", // 26 - Ivory
	"00CCFFFF", // 27 - Light Turquoise
	"00660066", // 28 - Dark Purple
	"00FF8080", // 29 - Coral
	"000066CC", // 30 - Ocean Blue
	"00CCCCFF", // 31 - Ice Blue
	"00000080", // 32 - Dark Blue
	"00FF00FF", // 33 - Pink
	"00FFFF00", // 34 - Yellow
	"0000FFFF", // 35 - Turquoise
	"00800080", // 36 - Violet
	"00800000", // 37 - Dark Red
	"00008080", // 38 - Teal
	"000000FF", // 39 - Blue
	"0000CCFF", // 40 - Sky Blue
	"00CCFFFF", // 41 - Light Turquoise
	"00CCFFCC", // 42 - Light Green
	"00FFFF99", // 43 - Light Yellow
	"0099CCFF", // 44 - Pale Blue
	"00FF99CC", // 45 - Rose
	"00CC99FF", // 46 - Lavender
	"00FFCC99", // 47 - Tan
	"003366FF", // 48 - Light Blue
	"0033CCCC", // 49 - Aqua
	"0099CC00", // 50 - Lime
	"00FFCC00", // 51 - Gold
	"00FF9900", // 52 - Light Orange
	"00FF6600", // 53 - Orange
	"00666699", // 54 - Blue Gray
	"00969696", // 55 - Gray 40%
	"00003366", // 56 - Dark Teal
	"00339966", // 57 - Sea Green
	"00003300", // 58 - Dark Green
	"00333300", // 59 - Olive Green
	"00993300", // 60 - Brown
	"00993366", // 61 - Plum
	"00333399", // 62 - Indigo
	"00333333", // 63 - Gray 80%
}

// Colors represents the colors container element (x:colors).
// This element contains custom color palettes for the workbook.
type Colors struct {
	*openxml.CompositeElementBase
}

// NewColors creates a new Colors element.
func NewColors() *Colors {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"colors",
		PrefixDefault,
	)

	return &Colors{CompositeElementBase: elem}
}

// IndexedColors returns the indexed colors element, or nil if not present.
func (c *Colors) IndexedColors() *IndexedColors {
	elem := c.GetElement(
		"indexedColors",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ic, ok := elem.(*IndexedColors); ok {
		return ic
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &IndexedColors{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateIndexedColors returns the indexed colors, creating if needed.
func (c *Colors) GetOrCreateIndexedColors() *IndexedColors {
	ic := c.IndexedColors()
	if ic != nil {
		return ic
	}
	ic = NewIndexedColors()
	// IndexedColors comes before MruColors
	mru := c.MruColors()
	if mru != nil {
		c.InsertBefore(ic, mru)
	} else {
		c.AppendChild(ic)
	}

	return ic
}

// MruColors returns the MRU colors element, or nil if not present.
func (c *Colors) MruColors() *MruColors {
	elem := c.GetElement(
		"mruColors",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if mru, ok := elem.(*MruColors); ok {
		return mru
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &MruColors{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateMruColors returns the MRU colors, creating if needed.
func (c *Colors) GetOrCreateMruColors() *MruColors {
	mru := c.MruColors()
	if mru != nil {
		return mru
	}
	mru = NewMruColors()
	c.AppendChild(mru)

	return mru
}

// Clone creates a deep copy of this Colors element.
func (c *Colors) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Colors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Colors element.
func (c *Colors) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Colors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// IndexedColors represents the indexed colors element (x:indexedColors).
// This contains a custom color palette that overrides the default palette.
type IndexedColors struct {
	*openxml.CompositeElementBase
}

// NewIndexedColors creates a new IndexedColors element.
func NewIndexedColors() *IndexedColors {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"indexedColors",
		PrefixDefault,
	)

	return &IndexedColors{
		CompositeElementBase: elem,
	}
}

// RgbColors returns an iterator over all RgbColor elements.
func (ic *IndexedColors) RgbColors() iter.Seq[*RgbColor] {
	return func(yield func(*RgbColor) bool) {
		for child := range ic.Children() {
			if child.LocalName() != "rgbColor" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var rgb *RgbColor
			if r, ok := child.(*RgbColor); ok {
				rgb = r
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				rgb = &RgbColor{LeafElementBase: leaf}
			}
			if rgb != nil && !yield(rgb) {
				return
			}
		}
	}
}

// GetRgbColor returns the RGB color at the given index, or nil if out of range.
func (ic *IndexedColors) GetRgbColor(
	index int,
) *RgbColor {
	i := 0
	for rgb := range ic.RgbColors() {
		if i == index {
			return rgb
		}
		i++
	}

	return nil
}

// AddRgbColor adds a new RGB color with the given value.
func (ic *IndexedColors) AddRgbColor(
	rgb string,
) *RgbColor {
	color := NewRgbColor()
	color.SetRgb(rgb)
	ic.AppendChild(color)

	return color
}

// ColorCount returns the number of RGB colors.
func (ic *IndexedColors) ColorCount() int {
	count := 0
	for range ic.RgbColors() {
		count++
	}

	return count
}

// Clone creates a deep copy of this IndexedColors element.
func (ic *IndexedColors) Clone() openxml.Element {
	cloned := ic.CompositeElementBase.Clone()

	return &IndexedColors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this IndexedColors element.
func (ic *IndexedColors) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ic.CompositeElementBase.CloneNode(
		deep,
	)

	return &IndexedColors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// RgbColor represents an RGB color element (x:rgbColor).
type RgbColor struct {
	*openxml.LeafElementBase
}

// NewRgbColor creates a new RgbColor element.
func NewRgbColor() *RgbColor {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"rgbColor",
		PrefixDefault,
	)

	return &RgbColor{LeafElementBase: elem}
}

// Rgb returns the RGB color value (ARGB hex, e.g., "FFFF0000").
func (r *RgbColor) Rgb() string {
	attr, found := r.GetAttribute("rgb", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRgb sets the RGB color value (ARGB hex, e.g., "FFFF0000").
func (r *RgbColor) SetRgb(rgb string) {
	r.SetAttribute(
		openxml.NewAttribute(
			"",
			"rgb",
			"",
			rgb,
		),
	)
}

// Clone creates a deep copy of this RgbColor element.
func (r *RgbColor) Clone() openxml.Element {
	cloned := r.LeafElementBase.Clone()

	return &RgbColor{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RgbColor element.
func (r *RgbColor) CloneNode(
	deep bool,
) openxml.Element {
	cloned := r.LeafElementBase.CloneNode(deep)

	return &RgbColor{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// MruColors represents the most recently used colors element (x:mruColors).
// This contains colors that the user has recently selected.
type MruColors struct {
	*openxml.CompositeElementBase
}

// NewMruColors creates a new MruColors element.
func NewMruColors() *MruColors {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"mruColors",
		PrefixDefault,
	)

	return &MruColors{CompositeElementBase: elem}
}

// Colors returns an iterator over all Color elements.
func (m *MruColors) Colors() iter.Seq[*Color] {
	return func(yield func(*Color) bool) {
		for child := range m.Children() {
			if child.LocalName() != "color" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var color *Color
			if c, ok := child.(*Color); ok {
				color = c
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				color = &Color{LeafElementBase: leaf}
			}
			if color != nil && !yield(color) {
				return
			}
		}
	}
}

// AddColor adds a new color.
func (m *MruColors) AddColor() *Color {
	color := NewColor()
	m.AppendChild(color)

	return color
}

// ColorCount returns the number of colors.
func (m *MruColors) ColorCount() int {
	count := 0
	for range m.Colors() {
		count++
	}

	return count
}

// Clone creates a deep copy of this MruColors element.
func (m *MruColors) Clone() openxml.Element {
	cloned := m.CompositeElementBase.Clone()

	return &MruColors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this MruColors element.
func (m *MruColors) CloneNode(
	deep bool,
) openxml.Element {
	cloned := m.CompositeElementBase.CloneNode(
		deep,
	)

	return &MruColors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
