package elements

//revive:disable:file-length-limit many border properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// BorderStyle represents the border line style.
type BorderStyle string

const (
	// BorderStyleNone indicates no border.
	BorderStyleNone BorderStyle = "none"
	// BorderStyleThin indicates a thin border.
	BorderStyleThin BorderStyle = "thin"
	// BorderStyleMedium indicates a medium border.
	BorderStyleMedium BorderStyle = "medium"
	// BorderStyleDashed indicates a dashed border.
	BorderStyleDashed BorderStyle = "dashed"
	// BorderStyleDotted indicates a dotted border.
	BorderStyleDotted BorderStyle = "dotted"
	// BorderStyleThick indicates a thick border.
	BorderStyleThick BorderStyle = "thick"
	// BorderStyleDouble indicates a double border.
	BorderStyleDouble BorderStyle = "double"
	// BorderStyleHair indicates a hair-thin border.
	BorderStyleHair BorderStyle = "hair"
	// BorderStyleMediumDashed indicates a medium dashed border.
	BorderStyleMediumDashed BorderStyle = "mediumDashed"
	// BorderStyleDashDot indicates a dash-dot border.
	BorderStyleDashDot BorderStyle = "dashDot"
	// BorderStyleMediumDashDot indicates a medium dash-dot border.
	BorderStyleMediumDashDot BorderStyle = "mediumDashDot"
	// BorderStyleDashDotDot indicates a dash-dot-dot border.
	BorderStyleDashDotDot BorderStyle = "dashDotDot"
	// BorderStyleMediumDashDotDot indicates a medium dash-dot-dot border.
	BorderStyleMediumDashDotDot BorderStyle = "mediumDashDotDot"
	// BorderStyleSlantDashDot indicates a slant dash-dot border.
	BorderStyleSlantDashDot BorderStyle = "slantDashDot"
)

// Borders represents the borders container element (x:borders).
type Borders struct {
	*openxml.CompositeElementBase
}

// NewBorders creates a new Borders element.
func NewBorders() *Borders {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"borders",
		PrefixDefault,
	)

	return &Borders{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (b *Borders) Count() uint32 {
	attr, found := b.GetAttribute("count", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10, //nolint:revive // add-constant
		32, //nolint:revive // add-constant
	)

	return uint32(val)
}

// SetCount sets the count attribute.
func (b *Borders) SetCount(count uint32) {
	b.SetAttribute(
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

// Borders returns an iterator over all Border elements.
func (b *Borders) Borders() iter.Seq[*Border] {
	return func(yield func(*Border) bool) {
		for child := range b.Children() {
			if child.LocalName() != "border" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var border *Border
			switch v := child.(type) {
			case *Border:
				border = v
			case *openxml.CompositeElementBase:
				border = &Border{CompositeElementBase: v}
			}
			if border != nil && !yield(border) {
				return
			}
		}
	}
}

// GetBorder returns the border at the given index, or nil if out of range.
func (b *Borders) GetBorder(
	index uint32,
) *Border {
	var i uint32
	for border := range b.Borders() {
		if i == index {
			return border
		}
		i++
	}

	return nil
}

// AddBorder adds a new border and returns it.
func (b *Borders) AddBorder() *Border {
	border := NewBorder()
	b.AppendChild(border)
	b.SetCount(b.Count() + 1)

	return border
}

// ItemCount returns the actual number of Border children.
func (b *Borders) ItemCount() int {
	count := 0
	for range b.Borders() {
		count++
	}

	return count
}

// Clone creates a deep copy of this Borders element.
func (b *Borders) Clone() openxml.Element {
	cloned := b.CompositeElementBase.Clone()

	return &Borders{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Borders element.
func (b *Borders) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.CompositeElementBase.CloneNode(
		deep,
	)

	return &Borders{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Border represents a border element (x:border) in the stylesheet.
type Border struct {
	*openxml.CompositeElementBase
}

// NewBorder creates a new Border element.
func NewBorder() *Border {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"border",
		PrefixDefault,
	)

	return &Border{CompositeElementBase: elem}
}

// DiagonalUp returns whether the diagonal-up border is enabled.
func (b *Border) DiagonalUp() bool {
	attr, found := b.GetAttribute(
		"diagonalUp",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDiagonalUp sets whether the diagonal-up border is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *Border) SetDiagonalUp(value bool) {
	if value {
		b.SetAttribute(
			openxml.NewAttribute(
				"",
				"diagonalUp",
				"",
				attrValueTrue,
			),
		)
	} else {
		b.RemoveAttribute("diagonalUp", "")
	}
}

// DiagonalDown returns whether the diagonal-down border is enabled.
func (b *Border) DiagonalDown() bool {
	attr, found := b.GetAttribute(
		"diagonalDown",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDiagonalDown sets whether the diagonal-down border is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *Border) SetDiagonalDown(value bool) {
	if value {
		b.SetAttribute(
			openxml.NewAttribute(
				"",
				"diagonalDown",
				"",
				attrValueTrue,
			),
		)
	} else {
		b.RemoveAttribute("diagonalDown", "")
	}
}

// Outline returns whether the outline border is enabled.
func (b *Border) Outline() bool {
	attr, found := b.GetAttribute("outline", "")
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetOutline sets whether the outline border is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (b *Border) SetOutline(value bool) {
	if value {
		b.RemoveAttribute("outline", "")
	} else {
		b.SetAttribute(
			openxml.NewAttribute("", "outline", "", attrValueFalse),
		)
	}
}

// Left returns the left border element, or nil if not present.
func (b *Border) Left() *BorderPr {
	elem := b.GetElement("left", NamespaceSML)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*BorderPr); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BorderPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateLeft returns the left border, creating if needed.
func (b *Border) GetOrCreateLeft() *BorderPr {
	bp := b.Left()
	if bp != nil {
		return bp
	}
	bp = NewBorderPrLeft()
	b.insertBorderPr(bp, "left")

	return bp
}

// Right returns the right border element, or nil if not present.
func (b *Border) Right() *BorderPr {
	elem := b.GetElement("right", NamespaceSML)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*BorderPr); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BorderPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRight returns the right border, creating if needed.
func (b *Border) GetOrCreateRight() *BorderPr {
	bp := b.Right()
	if bp != nil {
		return bp
	}
	bp = NewBorderPrRight()
	b.insertBorderPr(bp, "right")

	return bp
}

// Top returns the top border element, or nil if not present.
func (b *Border) Top() *BorderPr {
	elem := b.GetElement("top", NamespaceSML)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*BorderPr); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BorderPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTop returns the top border, creating if needed.
func (b *Border) GetOrCreateTop() *BorderPr {
	bp := b.Top()
	if bp != nil {
		return bp
	}
	bp = NewBorderPrTop()
	b.insertBorderPr(bp, "top")

	return bp
}

// Bottom returns the bottom border element, or nil if not present.
func (b *Border) Bottom() *BorderPr {
	elem := b.GetElement("bottom", NamespaceSML)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*BorderPr); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BorderPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBottom returns the bottom border, creating if needed.
func (b *Border) GetOrCreateBottom() *BorderPr {
	bp := b.Bottom()
	if bp != nil {
		return bp
	}
	bp = NewBorderPrBottom()
	b.insertBorderPr(bp, "bottom")

	return bp
}

// Diagonal returns the diagonal border element, or nil if not present.
func (b *Border) Diagonal() *BorderPr {
	elem := b.GetElement("diagonal", NamespaceSML)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*BorderPr); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BorderPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDiagonal returns the diagonal border, creating if needed.
func (b *Border) GetOrCreateDiagonal() *BorderPr {
	bp := b.Diagonal()
	if bp != nil {
		return bp
	}
	bp = NewBorderPrDiagonal()
	b.insertBorderPr(bp, "diagonal")

	return bp
}

// insertBorderPr inserts a border property element in the correct order.
// Order is: left, right, top, bottom, diagonal.
func (b *Border) insertBorderPr(
	bp *BorderPr,
	name string,
) {
	// Border elements must be in order: left, right, top, bottom, diagonal
	//nolint:revive // add-constant: border element order positions
	order := map[string]int{
		"left":     0,
		"right":    1,
		"top":      2,
		"bottom":   3, //nolint:revive // add-constant
		"diagonal": 4, //nolint:revive // add-constant
	}

	targetOrder := order[name]

	// Find the first element that should come after this one
	for child := range b.Children() {
		if childOrder, ok := order[child.LocalName()]; ok &&
			childOrder > targetOrder {
			b.InsertBefore(bp, child)

			return
		}
	}

	// No element found that should come after, so append
	b.AppendChild(bp)
}

// Clone creates a deep copy of this Border element.
func (b *Border) Clone() openxml.Element {
	cloned := b.CompositeElementBase.Clone()

	return &Border{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Border element.
func (b *Border) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.CompositeElementBase.CloneNode(
		deep,
	)

	return &Border{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// BorderPr represents a border property element
// (left, right, top, bottom, diagonal).
type BorderPr struct {
	*openxml.CompositeElementBase
}

// NewBorderPr creates a new BorderPr element with the given local name.
func NewBorderPr(localName string) *BorderPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		localName,
		PrefixDefault,
	)

	return &BorderPr{CompositeElementBase: elem}
}

// NewBorderPrLeft creates a new left border element.
func NewBorderPrLeft() *BorderPr {
	return NewBorderPr(
		"left", //nolint:revive // add-constant: border element name
	)
}

// NewBorderPrRight creates a new right border element.
func NewBorderPrRight() *BorderPr {
	return NewBorderPr(
		"right", //nolint:revive // add-constant: border element name
	)
}

// NewBorderPrTop creates a new top border element.
func NewBorderPrTop() *BorderPr {
	return NewBorderPr(
		"top", //nolint:revive // add-constant: border element name
	)
}

// NewBorderPrBottom creates a new bottom border element.
func NewBorderPrBottom() *BorderPr {
	return NewBorderPr(
		"bottom", //nolint:revive // add-constant: border element name
	)
}

// NewBorderPrDiagonal creates a new diagonal border element.
func NewBorderPrDiagonal() *BorderPr {
	return NewBorderPr(
		"diagonal", //nolint:revive // add-constant: border element name
	)
}

// Style returns the border style.
func (bp *BorderPr) Style() BorderStyle {
	attr, found := bp.GetAttribute("style", "")
	if !found {
		return BorderStyleNone
	}

	return BorderStyle(attr.Value())
}

// SetStyle sets the border style.
func (bp *BorderPr) SetStyle(style BorderStyle) {
	if style == "" || style == BorderStyleNone {
		bp.RemoveAttribute("style", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"style",
			"",
			string(style),
		),
	)
}

// Color returns the color element, or nil if not present.
func (bp *BorderPr) Color() *Color {
	elem := bp.GetElement("color", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Color); ok {
		return c
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &Color{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateColor returns the color element, creating if needed.
func (bp *BorderPr) GetOrCreateColor() *Color {
	c := bp.Color()
	if c != nil {
		return c
	}
	c = NewColor()
	bp.AppendChild(c)

	return c
}

// Clone creates a deep copy of this BorderPr element.
func (bp *BorderPr) Clone() openxml.Element {
	cloned := bp.CompositeElementBase.Clone()

	return &BorderPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this BorderPr element.
func (bp *BorderPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := bp.CompositeElementBase.CloneNode(
		deep,
	)

	return &BorderPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
