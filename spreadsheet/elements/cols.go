package elements

//revive:disable:file-length-limit many column properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Cols represents the column definitions container element (x:cols).
type Cols struct {
	*openxml.CompositeElementBase
}

// NewCols creates a new Cols element.
func NewCols() *Cols {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"cols",
		PrefixDefault,
	)

	return &Cols{CompositeElementBase: elem}
}

// Cols returns an iterator over all Col elements.
func (c *Cols) Cols() iter.Seq[*Col] {
	return func(yield func(*Col) bool) {
		for child := range c.Children() {
			if child.LocalName() != "col" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var col *Col
			if cl, ok := child.(*Col); ok {
				col = cl
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				col = &Col{CompositeElementBase: comp}
			}
			if col != nil && !yield(col) {
				return
			}
		}
	}
}

// AddCol adds a new column definition and returns it.
func (c *Cols) AddCol(minCol, maxCol int) *Col {
	col := NewCol()
	col.SetMin(minCol)
	col.SetMax(maxCol)
	c.AppendChild(col)

	return col
}

// GetColByIndex returns the column definition that contains the given column
// index, or nil if not found.
func (c *Cols) GetColByIndex(index int) *Col {
	for col := range c.Cols() {
		if index >= col.Min() &&
			index <= col.Max() {
			return col
		}
	}

	return nil
}

// Clone creates a deep copy of this Cols element.
func (c *Cols) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Cols{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Cols element.
func (c *Cols) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Cols{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Col represents a column definition element (x:col).
type Col struct {
	*openxml.CompositeElementBase
}

// NewCol creates a new Col element.
func NewCol() *Col {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"col",
		PrefixDefault,
	)

	return &Col{CompositeElementBase: elem}
}

// Min returns the first column index of this definition (1-based).
func (c *Col) Min() int {
	attr, found := c.GetAttribute("min", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetMin sets the first column index of this definition.
func (c *Col) SetMin(minVal int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"min",
			"",
			strconv.Itoa(minVal),
		),
	)
}

// Max returns the last column index of this definition (1-based).
func (c *Col) Max() int {
	attr, found := c.GetAttribute("max", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetMax sets the last column index of this definition.
func (c *Col) SetMax(maxVal int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"max",
			"",
			strconv.Itoa(maxVal),
		),
	)
}

// Width returns the column width.
func (c *Col) Width() float64 {
	attr, found := c.GetAttribute("width", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetWidth sets the column width.
func (c *Col) SetWidth(width float64) {
	if width == 0 {
		c.RemoveAttribute("width", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"width",
			"",
			strconv.FormatFloat(
				width,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Style returns the style index for the column.
func (c *Col) Style() int {
	attr, found := c.GetAttribute("style", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetStyle sets the style index for the column.
func (c *Col) SetStyle(style int) {
	if style == 0 {
		c.RemoveAttribute("style", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"style",
			"",
			strconv.Itoa(style),
		),
	)
}

// Hidden returns whether the column is hidden.
func (c *Col) Hidden() bool {
	attr, found := c.GetAttribute("hidden", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetHidden sets whether the column is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Col) SetHidden(value bool) {
	if value {
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

// BestFit returns whether the column width is set to best fit.
func (c *Col) BestFit() bool {
	attr, found := c.GetAttribute("bestFit", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetBestFit sets whether the column width is set to best fit.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Col) SetBestFit(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"bestFit",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("bestFit", "")
	}
}

// CustomWidth returns whether a custom width was set.
func (c *Col) CustomWidth() bool {
	attr, found := c.GetAttribute(
		"customWidth",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCustomWidth sets whether a custom width was set.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Col) SetCustomWidth(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"customWidth",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("customWidth", "")
	}
}

// Collapsed returns whether the column outline is collapsed.
func (c *Col) Collapsed() bool {
	attr, found := c.GetAttribute("collapsed", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetCollapsed sets whether the column outline is collapsed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Col) SetCollapsed(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"collapsed",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("collapsed", "")
	}
}

// OutlineLevel returns the outline level of the column (0-7).
func (c *Col) OutlineLevel() int {
	attr, found := c.GetAttribute(
		"outlineLevel",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetOutlineLevel sets the outline level of the column.
func (c *Col) SetOutlineLevel(level int) {
	if level == 0 {
		c.RemoveAttribute("outlineLevel", "")

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"outlineLevel",
			"",
			strconv.Itoa(level),
		),
	)
}

// Phonetic returns whether phonetic information should be shown.
func (c *Col) Phonetic() bool {
	attr, found := c.GetAttribute("phonetic", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetPhonetic sets whether phonetic information should be shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *Col) SetPhonetic(value bool) {
	if value {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"phonetic",
				"",
				attrValueTrue,
			),
		)
	} else {
		c.RemoveAttribute("phonetic", "")
	}
}

// Clone creates a deep copy of this Col element.
func (c *Col) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Col{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Col element.
func (c *Col) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Col{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
