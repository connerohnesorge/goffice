package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// attrValueOne is the string value "1" used for boolean attributes.
const attrValueOne = "1"

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
