package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// GradientStop represents a gradient stop element (x:stop).
type GradientStop struct {
	*openxml.CompositeElementBase
}

// NewGradientStop creates a new GradientStop element.
func NewGradientStop() *GradientStop {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"stop",
		PrefixDefault,
	)

	return &GradientStop{
		CompositeElementBase: elem,
	}
}

// Position returns the position of the stop (0.0-1.0).
func (gs *GradientStop) Position() float64 {
	attr, found := gs.GetAttribute("position", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetPosition sets the position of the stop (0.0-1.0).
func (gs *GradientStop) SetPosition(
	position float64,
) {
	gs.SetAttribute(
		openxml.NewAttribute(
			"",
			"position",
			"",
			strconv.FormatFloat(
				position,
				'f',
				-1,
				64, //nolint:revive // add-constant: bit size 64
			),
		),
	)
}

// Color returns the color element, or nil if not present.
func (gs *GradientStop) Color() *Color {
	elem := gs.GetElement("color", NamespaceSML)
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
func (gs *GradientStop) GetOrCreateColor() *Color {
	c := gs.Color()
	if c != nil {
		return c
	}
	c = NewColor()
	gs.AppendChild(c)

	return c
}

// Clone creates a deep copy of this GradientStop element.
func (gs *GradientStop) Clone() openxml.Element {
	cloned := gs.CompositeElementBase.Clone()

	return &GradientStop{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GradientStop element.
func (gs *GradientStop) CloneNode(
	deep bool,
) openxml.Element {
	cloned := gs.CompositeElementBase.CloneNode(
		deep,
	)

	return &GradientStop{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
