package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// GradientType represents the gradient fill type.
type GradientType string

const (
	// GradientTypeLinear indicates a linear gradient.
	GradientTypeLinear GradientType = "linear"
	// GradientTypePath indicates a path gradient.
	GradientTypePath GradientType = "path"
)

// GradientFill represents the gradient fill element (x:gradientFill).
type GradientFill struct {
	*openxml.CompositeElementBase
}

// NewGradientFill creates a new GradientFill element.
func NewGradientFill() *GradientFill {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"gradientFill",
		PrefixDefault,
	)

	return &GradientFill{
		CompositeElementBase: elem,
	}
}

// Type returns the gradient type.
func (gf *GradientFill) Type() GradientType {
	attr, found := gf.GetAttribute("type", "")
	if !found {
		return GradientTypeLinear
	}

	return GradientType(attr.Value())
}

// SetType sets the gradient type.
func (gf *GradientFill) SetType(
	gradientType GradientType,
) {
	if gradientType == "" ||
		gradientType == GradientTypeLinear {
		gf.RemoveAttribute("type", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(gradientType),
		),
	)
}

// Degree returns the degree of rotation for linear gradients.
func (gf *GradientFill) Degree() float64 {
	attr, found := gf.GetAttribute("degree", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetDegree sets the degree of rotation for linear gradients.
func (gf *GradientFill) SetDegree(
	degree float64,
) {
	if degree == 0 {
		gf.RemoveAttribute("degree", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"degree",
			"",
			strconv.FormatFloat(
				degree,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Left returns the left position for path gradients (0.0-1.0).
func (gf *GradientFill) Left() float64 {
	attr, found := gf.GetAttribute("left", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetLeft sets the left position for path gradients (0.0-1.0).
func (gf *GradientFill) SetLeft(left float64) {
	if left == 0 {
		gf.RemoveAttribute("left", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"left",
			"",
			strconv.FormatFloat(
				left,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Right returns the right position for path gradients (0.0-1.0).
func (gf *GradientFill) Right() float64 {
	attr, found := gf.GetAttribute("right", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		bitSize64,
	)

	return val
}

// SetRight sets the right position for path gradients (0.0-1.0).
func (gf *GradientFill) SetRight(right float64) {
	if right == 0 {
		gf.RemoveAttribute("right", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"right",
			"",
			strconv.FormatFloat(
				right,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Top returns the top position for path gradients (0.0-1.0).
func (gf *GradientFill) Top() float64 {
	attr, found := gf.GetAttribute("top", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetTop sets the top position for path gradients (0.0-1.0).
func (gf *GradientFill) SetTop(top float64) {
	if top == 0 {
		gf.RemoveAttribute("top", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"top",
			"",
			strconv.FormatFloat(
				top,
				'f',
				-1,
				64, //nolint:revive // add-constant
			),
		),
	)
}

// Bottom returns the bottom position for path gradients (0.0-1.0).
func (gf *GradientFill) Bottom() float64 {
	attr, found := gf.GetAttribute("bottom", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseFloat(
		attr.Value(),
		64, //nolint:revive // add-constant
	)

	return val
}

// SetBottom sets the bottom position for path gradients (0.0-1.0).
func (gf *GradientFill) SetBottom(
	bottom float64,
) {
	if bottom == 0 {
		gf.RemoveAttribute("bottom", "")

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"bottom",
			"",
			strconv.FormatFloat(
				bottom,
				'f',
				-1,
				bitSize64,
			),
		),
	)
}

// Stops returns an iterator over all GradientStop elements.
func (gf *GradientFill) Stops() iter.Seq[*GradientStop] {
	return func(yield func(*GradientStop) bool) {
		for child := range gf.Children() {
			if child.LocalName() != "stop" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var stop *GradientStop
			switch v := child.(type) {
			case *GradientStop:
				stop = v
			case *openxml.CompositeElementBase:
				stop = &GradientStop{CompositeElementBase: v}
			}
			if stop != nil && !yield(stop) {
				return
			}
		}
	}
}

// AddStop adds a new gradient stop at the given position (0.0-1.0).
func (gf *GradientFill) AddStop(
	position float64,
) *GradientStop {
	stop := NewGradientStop()
	stop.SetPosition(position)
	gf.AppendChild(stop)

	return stop
}

// Clone creates a deep copy of this GradientFill element.
func (gf *GradientFill) Clone() openxml.Element {
	cloned := gf.CompositeElementBase.Clone()

	return &GradientFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GradientFill element.
func (gf *GradientFill) CloneNode(
	deep bool,
) openxml.Element {
	cloned := gf.CompositeElementBase.CloneNode(
		deep,
	)

	return &GradientFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
