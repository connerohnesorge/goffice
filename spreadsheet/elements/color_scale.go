package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

const (
	elemNameColor = "color"
)

// ColorScale represents a color scale element (x:colorScale)
// for conditional formatting.
// A color scale applies a gradual color fill based on cell values.
// Contains 2-3 cfvo elements and 2-3 matching color elements.
type ColorScale struct {
	*openxml.CompositeElementBase
}

// NewColorScale creates a new ColorScale element.
func NewColorScale() *ColorScale {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"colorScale",
		PrefixDefault,
	)

	return &ColorScale{CompositeElementBase: elem}
}

// Cfvos returns an iterator over all Cfvo elements in this color scale.
func (cs *ColorScale) Cfvos() iter.Seq[*Cfvo] {
	return func(yield func(*Cfvo) bool) {
		for child := range cs.Children() {
			if child.LocalName() != elemNameCfvo ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var cfvo *Cfvo
			if c, ok := child.(*Cfvo); ok {
				cfvo = c
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				cfvo = &Cfvo{LeafElementBase: leaf}
			}
			if cfvo != nil && !yield(cfvo) {
				return
			}
		}
	}
}

// CfvoCount returns the number of Cfvo elements.
func (cs *ColorScale) CfvoCount() int {
	count := 0
	for range cs.Cfvos() {
		count++
	}

	return count
}

// AddCfvo adds a new Cfvo element to the color scale.
func (cs *ColorScale) AddCfvo() *Cfvo {
	cfvo := NewCfvo()
	// Insert before colors
	var insertBefore openxml.Element
	for child := range cs.Children() {
		if child.LocalName() == elemNameColor &&
			child.NamespaceURI() == NamespaceSML {
			insertBefore = child

			break
		}
	}
	if insertBefore != nil {
		cs.InsertBefore(cfvo, insertBefore)
	} else {
		cs.AppendChild(cfvo)
	}

	return cfvo
}

// Colors returns an iterator over all Color elements in this color scale.
func (cs *ColorScale) Colors() iter.Seq[*Color] {
	return func(yield func(*Color) bool) {
		for child := range cs.Children() {
			if child.LocalName() != elemNameColor ||
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

// ColorCount returns the number of Color elements.
func (cs *ColorScale) ColorCount() int {
	count := 0
	for range cs.Colors() {
		count++
	}

	return count
}

// AddColor adds a new Color element to the color scale.
func (cs *ColorScale) AddColor() *Color {
	color := NewColor()
	cs.AppendChild(color)

	return color
}

// SetTwoColorScale sets up a two-color scale with min and max values.
//
//nolint:revive // argument-limit
func (cs *ColorScale) SetTwoColorScale(
	minType CfvoType,
	minVal, minColor string,
	maxType CfvoType,
	maxVal, maxColor string,
) {
	// Clear existing
	cs.RemoveAllChildren()

	// Add cfvo elements
	minCfvo := cs.AddCfvo()
	minCfvo.SetType(minType)
	if minVal != "" {
		minCfvo.SetVal(minVal)
	}

	maxCfvo := cs.AddCfvo()
	maxCfvo.SetType(maxType)
	if maxVal != "" {
		maxCfvo.SetVal(maxVal)
	}

	// Add color elements
	minColorElem := cs.AddColor()
	minColorElem.SetRGB(minColor)

	maxColorElem := cs.AddColor()
	maxColorElem.SetRGB(maxColor)
}

// SetThreeColorScale sets up a three-color scale with min, mid, and max values.
//
//nolint:revive // argument-limit
func (cs *ColorScale) SetThreeColorScale(
	minType CfvoType,
	minVal, minColor string,
	midType CfvoType,
	midVal, midColor string,
	maxType CfvoType,
	maxVal, maxColor string,
) {
	// Clear existing
	cs.RemoveAllChildren()

	// Add cfvo elements
	minCfvo := cs.AddCfvo()
	minCfvo.SetType(minType)
	if minVal != "" {
		minCfvo.SetVal(minVal)
	}

	midCfvo := cs.AddCfvo()
	midCfvo.SetType(midType)
	if midVal != "" {
		midCfvo.SetVal(midVal)
	}

	maxCfvo := cs.AddCfvo()
	maxCfvo.SetType(maxType)
	if maxVal != "" {
		maxCfvo.SetVal(maxVal)
	}

	// Add color elements
	minColorElem := cs.AddColor()
	minColorElem.SetRGB(minColor)

	midColorElem := cs.AddColor()
	midColorElem.SetRGB(midColor)

	maxColorElem := cs.AddColor()
	maxColorElem.SetRGB(maxColor)
}

// Clone creates a deep copy of this ColorScale element.
func (cs *ColorScale) Clone() openxml.Element {
	cloned := cs.CompositeElementBase.Clone()

	return &ColorScale{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ColorScale element.
func (cs *ColorScale) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cs.CompositeElementBase.CloneNode(
		deep,
	)

	return &ColorScale{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
