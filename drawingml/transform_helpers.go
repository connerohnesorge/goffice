package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// attrValueOne is the string value "1" used for boolean attributes.
const attrValueOne = "1"

// localNameChOff is the local name for child offset element.
const localNameChOff = "chOff"

// Transform2D convenience methods for common operations.
// These methods provide simpler access to transform properties.

// Rotation returns the rotation in 60,000ths of a degree.
// A positive value indicates clockwise rotation.
func (t *Transform2D) Rotation() int {
	attr, found := t.GetAttribute("rot", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetRotation sets the rotation in 60,000ths of a degree.
// A positive value indicates clockwise rotation.
func (t *Transform2D) SetRotation(rot int) {
	if rot == 0 {
		// Remove attribute if zero
		t.RemoveAttribute("rot", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"rot",
			"",
			strconv.Itoa(rot),
		),
	)
}

// RotationDegrees returns the rotation in degrees.
func (t *Transform2D) RotationDegrees() float64 {
	return AngleUnitsToDegrees(t.Rotation())
}

// SetRotationDegrees sets the rotation in degrees.
func (t *Transform2D) SetRotationDegrees(
	degrees float64,
) {
	t.SetRotation(DegreesToAngleUnits(degrees))
}

// FlipH returns whether the shape is flipped horizontally.
func (t *Transform2D) FlipH() bool {
	attr, found := t.GetAttribute("flipH", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne ||
		attr.Value() == attrTrue
}

// SetFlipH sets whether the shape is flipped horizontally.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Transform2D) SetFlipH(flip bool) {
	if !flip {
		// Remove attribute if false
		t.RemoveAttribute("flipH", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"flipH",
			"",
			attrValueOne,
		),
	)
}

// FlipV returns whether the shape is flipped vertically.
func (t *Transform2D) FlipV() bool {
	attr, found := t.GetAttribute("flipV", "")
	if !found {
		return false
	}

	return attr.Value() == attrValueOne ||
		attr.Value() == attrTrue
}

// SetFlipV sets whether the shape is flipped vertically.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (t *Transform2D) SetFlipV(flip bool) {
	if !flip {
		// Remove attribute if false
		t.RemoveAttribute("flipV", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"flipV",
			"",
			attrValueOne,
		),
	)
}

// Clone creates a deep copy of this Transform2D element.
func (t *Transform2D) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &Transform2D{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Transform2D element.
func (t *Transform2D) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &Transform2D{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Width returns the width of the transform extent.
func (t *Transform2D) Width() EMU {
	return t.Extent().Cx
}

// SetWidth sets the width of the transform extent.
func (t *Transform2D) SetWidth(width EMU) {
	ext := t.Extent()
	ext.Cx = width
	t.SetExtent(ext)
}

// Height returns the height of the transform extent.
func (t *Transform2D) Height() EMU {
	return t.Extent().Cy
}

// SetHeight sets the height of the transform extent.
func (t *Transform2D) SetHeight(height EMU) {
	ext := t.Extent()
	ext.Cy = height
	t.SetExtent(ext)
}

// OffsetX returns the X offset of the transform.
func (t *Transform2D) OffsetX() EMU {
	return t.Offset().X
}

// SetOffsetX sets the X offset of the transform.
func (t *Transform2D) SetOffsetX(x EMU) {
	off := t.Offset()
	off.X = x
	t.SetOffset(off)
}

// OffsetY returns the Y offset of the transform.
func (t *Transform2D) OffsetY() EMU {
	return t.Offset().Y
}

// SetOffsetY sets the Y offset of the transform.
func (t *Transform2D) SetOffsetY(y EMU) {
	off := t.Offset()
	off.Y = y
	t.SetOffset(off)
}

// ScaleBy scales the transform extent by the given factor.
func (t *Transform2D) ScaleBy(factor float64) {
	ext := t.Extent()
	t.SetExtent(ext.Scale(factor))
}

// MoveTo moves the transform to the specified position.
func (t *Transform2D) MoveTo(x, y EMU) {
	t.SetOffset(Offset{X: x, Y: y})
}

// MoveBy moves the transform by the specified offset.
func (t *Transform2D) MoveBy(dx, dy EMU) {
	off := t.Offset()
	t.SetOffset(
		Offset{X: off.X + dx, Y: off.Y + dy},
	)
}

// Resize sets the transform extent to the specified dimensions.
func (t *Transform2D) Resize(width, height EMU) {
	t.SetExtent(Extent{Cx: width, Cy: height})
}

// ChildOffset returns the child coordinate space origin (a:chOff element).
// For group shapes, this defines where the child coordinate space starts.
// Returns (x, y, present) where present is false if the element doesn't exist.
func (t *Transform2D) ChildOffset() (x, y EMU, present bool) {
	chOffElem := t.GetElement(
		localNameChOff,
		NamespaceMain,
	)
	if chOffElem == nil {
		return 0, 0, false
	}

	var xVal, yVal EMU
	if xAttr, found := chOffElem.GetAttribute("x", ""); found {
		val, _ := strconv.ParseInt(
			xAttr.Value(),
			base10,
			bitSize64,
		)
		xVal = EMU(val)
	}
	if yAttr, found := chOffElem.GetAttribute("y", ""); found {
		val, _ := strconv.ParseInt(
			yAttr.Value(),
			base10,
			bitSize64,
		)
		yVal = EMU(val)
	}

	return xVal, yVal, true
}

// SetChildOffset sets the child coordinate space origin (a:chOff element).
// Creates the element if it doesn't exist.
func (t *Transform2D) SetChildOffset(x, y EMU) {
	chOffElem := t.GetElement(
		localNameChOff,
		NamespaceMain,
	)
	if chOffElem == nil {
		chOffElem = openxml.NewCompositeElement(
			NamespaceMain,
			localNameChOff,
			PrefixMain,
		)
		// Insert after extent element
		if extElem := t.GetElement("ext", NamespaceMain); extElem != nil {
			t.InsertAfter(chOffElem, extElem)
		} else {
			t.AppendChild(chOffElem)
		}
	}

	chOffElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(
				x.Int64(),
				base10,
			),
		),
	)
	chOffElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(
				y.Int64(),
				base10,
			),
		),
	)
}

// ChildExtent returns the child coordinate space size (a:chExt element).
// For group shapes, this defines the size of the child coordinate space.
// Returns (cx, cy, present) where present is false if the element doesn't exist.
func (t *Transform2D) ChildExtent() (cx, cy EMU, present bool) {
	chExtElem := t.GetElement(
		"chExt",
		NamespaceMain,
	)
	if chExtElem == nil {
		return 0, 0, false
	}

	var cxVal, cyVal EMU
	if cxAttr, found := chExtElem.GetAttribute("cx", ""); found {
		val, _ := strconv.ParseInt(
			cxAttr.Value(),
			base10,
			bitSize64,
		)
		cxVal = EMU(val)
	}
	if cyAttr, found := chExtElem.GetAttribute("cy", ""); found {
		val, _ := strconv.ParseInt(
			cyAttr.Value(),
			base10,
			bitSize64,
		)
		cyVal = EMU(val)
	}

	return cxVal, cyVal, true
}

// SetChildExtent sets the child coordinate space size (a:chExt element).
// Creates the element if it doesn't exist.
func (t *Transform2D) SetChildExtent(cx, cy EMU) {
	chExtElem := t.GetElement(
		"chExt",
		NamespaceMain,
	)
	if chExtElem == nil {
		chExtElem = openxml.NewCompositeElement(
			NamespaceMain,
			"chExt",
			PrefixMain,
		)
		// Insert after chOff element if it exists, otherwise after ext
		if chOffElem := t.GetElement(localNameChOff, NamespaceMain); chOffElem != nil {
			t.InsertAfter(chExtElem, chOffElem)
		} else if extElem := t.GetElement("ext", NamespaceMain); extElem != nil {
			t.InsertAfter(chExtElem, extElem)
		} else {
			t.AppendChild(chExtElem)
		}
	}

	chExtElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"cx",
			"",
			strconv.FormatInt(
				cx.Int64(),
				base10,
			),
		),
	)
	chExtElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"cy",
			"",
			strconv.FormatInt(
				cy.Int64(),
				base10,
			),
		),
	)
}
