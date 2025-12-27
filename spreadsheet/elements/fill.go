package elements

//revive:disable:file-length-limit SpreadsheetML fill types require many props

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// PatternType represents the fill pattern type.
type PatternType string

const (
	// PatternTypeNone indicates no fill pattern.
	PatternTypeNone PatternType = "none"
	// PatternTypeSolid indicates a solid fill.
	PatternTypeSolid PatternType = "solid"
	// PatternTypeMediumGray indicates a medium gray pattern.
	PatternTypeMediumGray PatternType = "mediumGray"
	// PatternTypeDarkGray indicates a dark gray pattern.
	PatternTypeDarkGray PatternType = "darkGray"
	// PatternTypeLightGray indicates a light gray pattern.
	PatternTypeLightGray PatternType = "lightGray"
	// PatternTypeDarkHorizontal indicates a dark horizontal pattern.
	PatternTypeDarkHorizontal PatternType = "darkHorizontal"
	// PatternTypeDarkVertical indicates a dark vertical pattern.
	PatternTypeDarkVertical PatternType = "darkVertical"
	// PatternTypeDarkDown indicates a dark down diagonal pattern.
	PatternTypeDarkDown PatternType = "darkDown"
	// PatternTypeDarkUp indicates a dark up diagonal pattern.
	PatternTypeDarkUp PatternType = "darkUp"
	// PatternTypeDarkGrid indicates a dark grid pattern.
	PatternTypeDarkGrid PatternType = "darkGrid"
	// PatternTypeDarkTrellis indicates a dark trellis pattern.
	PatternTypeDarkTrellis PatternType = "darkTrellis"
	// PatternTypeLightHorizontal indicates a light horizontal pattern.
	PatternTypeLightHorizontal PatternType = "lightHorizontal"
	// PatternTypeLightVertical indicates a light vertical pattern.
	PatternTypeLightVertical PatternType = "lightVertical"
	// PatternTypeLightDown indicates a light down diagonal pattern.
	PatternTypeLightDown PatternType = "lightDown"
	// PatternTypeLightUp indicates a light up diagonal pattern.
	PatternTypeLightUp PatternType = "lightUp"
	// PatternTypeLightGrid indicates a light grid pattern.
	PatternTypeLightGrid PatternType = "lightGrid"
	// PatternTypeLightTrellis indicates a light trellis pattern.
	PatternTypeLightTrellis PatternType = "lightTrellis"
	// PatternTypeGray125 indicates a 12.5% gray pattern.
	PatternTypeGray125 PatternType = "gray125"
	// PatternTypeGray0625 indicates a 6.25% gray pattern.
	PatternTypeGray0625 PatternType = "gray0625"
)

// Fills represents the fills container element (x:fills).
type Fills struct {
	*openxml.CompositeElementBase
}

// NewFills creates a new Fills element.
func NewFills() *Fills {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"fills",
		PrefixDefault,
	)

	return &Fills{CompositeElementBase: elem}
}

// Count returns the count attribute value.
func (f *Fills) Count() uint32 {
	attr, found := f.GetAttribute("count", "")
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
func (f *Fills) SetCount(count uint32) {
	f.SetAttribute(
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

// Fills returns an iterator over all Fill elements.
func (f *Fills) Fills() iter.Seq[*Fill] {
	return func(yield func(*Fill) bool) {
		for child := range f.Children() {
			if child.LocalName() != "fill" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var fill *Fill
			if fl, ok := child.(*Fill); ok {
				fill = fl
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				fill = &Fill{CompositeElementBase: comp}
			}
			if fill != nil && !yield(fill) {
				return
			}
		}
	}
}

// GetFill returns the fill at the given index, or nil if out of range.
func (f *Fills) GetFill(index uint32) *Fill {
	var i uint32
	for fill := range f.Fills() {
		if i == index {
			return fill
		}
		i++
	}

	return nil
}

// AddFill adds a new fill and returns it.
func (f *Fills) AddFill() *Fill {
	fill := NewFill()
	f.AppendChild(fill)
	f.SetCount(f.Count() + 1)

	return fill
}

// ItemCount returns the actual number of Fill children.
func (f *Fills) ItemCount() int {
	count := 0
	for range f.Fills() {
		count++
	}

	return count
}

// Clone creates a deep copy of this Fills element.
func (f *Fills) Clone() openxml.Element {
	cloned := f.CompositeElementBase.Clone()

	return &Fills{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Fills element.
func (f *Fills) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.CompositeElementBase.CloneNode(
		deep,
	)

	return &Fills{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Fill represents a fill element (x:fill) in the stylesheet.
type Fill struct {
	*openxml.CompositeElementBase
}

// NewFill creates a new Fill element.
func NewFill() *Fill {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"fill",
		PrefixDefault,
	)

	return &Fill{CompositeElementBase: elem}
}

// PatternFill returns the pattern fill element, or nil if not present.
func (f *Fill) PatternFill() *PatternFill {
	elem := f.GetElement(
		"patternFill",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pf, ok := elem.(*PatternFill); ok {
		return pf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PatternFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePatternFill returns the pattern fill, creating if needed.
// This will remove any existing gradient fill.
func (f *Fill) GetOrCreatePatternFill() *PatternFill {
	pf := f.PatternFill()
	if pf != nil {
		return pf
	}

	// Remove gradient fill if present
	if gf := f.GradientFill(); gf != nil {
		f.RemoveChild(gf)
	}

	pf = NewPatternFill()
	f.AppendChild(pf)

	return pf
}

// GradientFill returns the gradient fill element, or nil if not present.
func (f *Fill) GradientFill() *GradientFill {
	elem := f.GetElement(
		"gradientFill",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if gf, ok := elem.(*GradientFill); ok {
		return gf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GradientFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateGradientFill returns the gradient fill, creating if needed.
// This will remove any existing pattern fill.
func (f *Fill) GetOrCreateGradientFill() *GradientFill {
	gf := f.GradientFill()
	if gf != nil {
		return gf
	}

	// Remove pattern fill if present
	if pf := f.PatternFill(); pf != nil {
		f.RemoveChild(pf)
	}

	gf = NewGradientFill()
	f.AppendChild(gf)

	return gf
}

// Clone creates a deep copy of this Fill element.
func (f *Fill) Clone() openxml.Element {
	cloned := f.CompositeElementBase.Clone()

	return &Fill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Fill element.
func (f *Fill) CloneNode(
	deep bool,
) openxml.Element {
	cloned := f.CompositeElementBase.CloneNode(
		deep,
	)

	return &Fill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PatternFill represents the pattern fill element (x:patternFill).
type PatternFill struct {
	*openxml.CompositeElementBase
}

// NewPatternFill creates a new PatternFill element.
func NewPatternFill() *PatternFill {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"patternFill",
		PrefixDefault,
	)

	return &PatternFill{
		CompositeElementBase: elem,
	}
}

// PatternType returns the pattern type.
func (pf *PatternFill) PatternType() PatternType {
	attr, found := pf.GetAttribute(
		"patternType",
		"",
	)
	if !found {
		return PatternTypeNone
	}

	return PatternType(attr.Value())
}

// SetPatternType sets the pattern type.
func (pf *PatternFill) SetPatternType(
	patternType PatternType,
) {
	if patternType == "" ||
		patternType == PatternTypeNone {
		pf.RemoveAttribute("patternType", "")

		return
	}
	pf.SetAttribute(
		openxml.NewAttribute(
			"",
			"patternType",
			"",
			string(patternType),
		),
	)
}

// FgColor returns the foreground color element, or nil if not present.
func (pf *PatternFill) FgColor() *Color {
	elem := pf.GetElement("fgColor", NamespaceSML)
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

// GetOrCreateFgColor returns the foreground color, creating if needed.
func (pf *PatternFill) GetOrCreateFgColor() *Color {
	c := pf.FgColor()
	if c != nil {
		return c
	}
	c = NewFgColor()
	pf.AppendChild(c)

	return c
}

// BgColor returns the background color element, or nil if not present.
func (pf *PatternFill) BgColor() *Color {
	elem := pf.GetElement("bgColor", NamespaceSML)
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

// GetOrCreateBgColor returns the background color, creating if needed.
func (pf *PatternFill) GetOrCreateBgColor() *Color {
	c := pf.BgColor()
	if c != nil {
		return c
	}
	c = NewBgColor()
	pf.AppendChild(c)

	return c
}

// Clone creates a deep copy of this PatternFill element.
func (pf *PatternFill) Clone() openxml.Element {
	cloned := pf.CompositeElementBase.Clone()

	return &PatternFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PatternFill element.
func (pf *PatternFill) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pf.CompositeElementBase.CloneNode(
		deep,
	)

	return &PatternFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NewFgColor creates a new foreground color element (x:fgColor).
func NewFgColor() *Color {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"fgColor",
		PrefixDefault,
	)

	return &Color{LeafElementBase: elem}
}

// NewBgColor creates a new background color element (x:bgColor).
func NewBgColor() *Color {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"bgColor",
		PrefixDefault,
	)

	return &Color{LeafElementBase: elem}
}
