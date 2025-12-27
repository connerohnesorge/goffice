package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// base10 is the base used for integer to string conversion.
const base10 = 10

// bitSize64 is the bit size for parsing 64-bit integers.
const bitSize64 = 64

// Element name constants.
const (
	elemOff = "off"
	elemExt = "ext"
)

// Transform2D represents the a:xfrm element for 2D transformations.
// This element specifies the location and size of a drawing element,
// along with optional rotation and flip properties.
type Transform2D struct {
	*openxml.CompositeElementBase
}

// NewTransform2D creates a new Transform2D element with the specified
// offset and extent values.
func NewTransform2D(
	offX, offY, extCx, extCy EMU,
) *Transform2D {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"xfrm",
		PrefixMain,
	)
	xfrm := &Transform2D{
		CompositeElementBase: elem,
	}

	// Add offset element (a:off)
	off := openxml.NewCompositeElement(
		NamespaceMain,
		elemOff,
		PrefixMain,
	)
	off.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(
				offX.Int64(),
				base10,
			),
		),
	)
	off.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(
				offY.Int64(),
				base10,
			),
		),
	)
	xfrm.AppendChild(off)

	// Add extent element (a:ext)
	ext := openxml.NewCompositeElement(
		NamespaceMain,
		elemExt,
		PrefixMain,
	)
	ext.SetAttribute(
		openxml.NewAttribute(
			"",
			"cx",
			"",
			strconv.FormatInt(
				extCx.Int64(),
				base10,
			),
		),
	)
	ext.SetAttribute(
		openxml.NewAttribute(
			"",
			"cy",
			"",
			strconv.FormatInt(
				extCy.Int64(),
				base10,
			),
		),
	)
	xfrm.AppendChild(ext)

	return xfrm
}

// NewTransform2DWithOffset creates a new Transform2D from coordinate types.
func NewTransform2DWithOffset(
	offset Offset,
	extent Extent,
) *Transform2D {
	return NewTransform2D(
		offset.X,
		offset.Y,
		extent.Cx,
		extent.Cy,
	)
}

// NewTransform2DWithPoint creates a new Transform2D from a Point2D and Extent.
func NewTransform2DWithPoint(
	point Point2D,
	extent Extent,
) *Transform2D {
	return NewTransform2D(
		point.X,
		point.Y,
		extent.Cx,
		extent.Cy,
	)
}

// Offset returns the offset (position) of the transform.
func (t *Transform2D) Offset() Offset {
	offElem := t.GetElement(
		elemOff,
		NamespaceMain,
	)
	if offElem == nil {
		return Offset{}
	}

	var x, y EMU
	if xAttr, found := offElem.GetAttribute("x", ""); found {
		val, _ := strconv.ParseInt(
			xAttr.Value(),
			base10,
			bitSize64,
		)
		x = EMU(val)
	}
	if yAttr, found := offElem.GetAttribute("y", ""); found {
		val, _ := strconv.ParseInt(
			yAttr.Value(),
			base10,
			bitSize64,
		)
		y = EMU(val)
	}

	return Offset{X: x, Y: y}
}

// SetOffset sets the offset (position) of the transform.
func (t *Transform2D) SetOffset(offset Offset) {
	offElem := t.GetElement(
		elemOff,
		NamespaceMain,
	)
	if offElem == nil {
		offElem = openxml.NewCompositeElement(
			NamespaceMain,
			elemOff,
			PrefixMain,
		)
		// Insert at beginning
		t.PrependChild(offElem)
	}

	offElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(
				offset.X.Int64(),
				base10,
			),
		),
	)
	offElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(
				offset.Y.Int64(),
				base10,
			),
		),
	)
}

// Extent returns the extent (size) of the transform.
func (t *Transform2D) Extent() Extent {
	extElem := t.GetElement(
		elemExt,
		NamespaceMain,
	)
	if extElem == nil {
		return Extent{}
	}

	var cx, cy EMU
	if cxAttr, found := extElem.GetAttribute("cx", ""); found {
		val, _ := strconv.ParseInt(
			cxAttr.Value(),
			base10,
			bitSize64,
		)
		cx = EMU(val)
	}
	if cyAttr, found := extElem.GetAttribute("cy", ""); found {
		val, _ := strconv.ParseInt(
			cyAttr.Value(),
			base10,
			bitSize64,
		)
		cy = EMU(val)
	}

	return Extent{Cx: cx, Cy: cy}
}

// SetExtent sets the extent (size) of the transform.
func (t *Transform2D) SetExtent(extent Extent) {
	extElem := t.GetElement(
		elemExt,
		NamespaceMain,
	)
	if extElem == nil {
		extElem = openxml.NewCompositeElement(
			NamespaceMain,
			elemExt,
			PrefixMain,
		)
		t.AppendChild(extElem)
	}

	extElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"cx",
			"",
			strconv.FormatInt(
				extent.Cx.Int64(),
				base10,
			),
		),
	)
	extElem.SetAttribute(
		openxml.NewAttribute(
			"",
			"cy",
			"",
			strconv.FormatInt(
				extent.Cy.Int64(),
				base10,
			),
		),
	)
}
