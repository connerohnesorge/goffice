package transform

import (
	"math"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
)

func TestFromTransform2D_Simple(t *testing.T) {
	// Create a simple transform with offset only
	xfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	off := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	off.SetAttribute(openxml.NewAttribute("", "x", "", "100"))
	off.SetAttribute(openxml.NewAttribute("", "y", "", "200"))
	xfrm.AppendChild(off)

	ext := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	xfrm.AppendChild(ext)

	m := FromTransform2D(xfrm)

	// Should be a simple translation
	x, y := m.TransformPoint(0, 0)
	if math.Abs(x-100) > epsilon || math.Abs(y-200) > epsilon {
		t.Errorf("Transform of (0, 0) = (%f, %f), want (100, 200)", x, y)
	}
}

func TestFromTransform2D_WithRotation(t *testing.T) {
	// Create transform with offset and rotation
	xfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	off := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	off.SetAttribute(openxml.NewAttribute("", "x", "", "0"))
	off.SetAttribute(openxml.NewAttribute("", "y", "", "0"))
	xfrm.AppendChild(off)

	ext := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	xfrm.AppendChild(ext)

	// 90 degrees = 5400000 (90 * 60000)
	xfrm.SetAttribute(openxml.NewAttribute("", "rot", "", "5400000"))

	m := FromTransform2D(xfrm)

	// Point (1, 0) should rotate to (0, 1)
	x, y := m.TransformPoint(1, 0)
	if math.Abs(x-0) > epsilon || math.Abs(y-1) > epsilon {
		t.Errorf("90° rotation of (1, 0) = (%f, %f), want (0, 1)", x, y)
	}
}

func TestFromTransform2D_WithFlip(t *testing.T) {
	// Create transform with horizontal flip
	xfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	off := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	off.SetAttribute(openxml.NewAttribute("", "x", "", "0"))
	off.SetAttribute(openxml.NewAttribute("", "y", "", "0"))
	xfrm.AppendChild(off)

	ext := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	xfrm.AppendChild(ext)

	xfrm.SetAttribute(openxml.NewAttribute("", "flipH", "", "1"))

	m := FromTransform2D(xfrm)

	// Point (10, 5) should flip horizontally to (-10, 5)
	x, y := m.TransformPoint(10, 5)
	if math.Abs(x-(-10)) > epsilon || math.Abs(y-5) > epsilon {
		t.Errorf("Horizontal flip of (10, 5) = (%f, %f), want (-10, 5)", x, y)
	}
}

func TestFromTransform2D_WithViewport(t *testing.T) {
	// Create group transform with child coordinate space
	xfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	off := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	off.SetAttribute(openxml.NewAttribute("", "x", "", "100"))
	off.SetAttribute(openxml.NewAttribute("", "y", "", "200"))
	xfrm.AppendChild(off)

	ext := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	xfrm.AppendChild(ext)

	// Child extent: 2000 x 1600 (will be scaled down to 1000 x 800)
	chExt := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "chExt", "a")
	chExt.SetAttribute(openxml.NewAttribute("", "cx", "", "2000"))
	chExt.SetAttribute(openxml.NewAttribute("", "cy", "", "1600"))
	xfrm.AppendChild(chExt)

	// Child offset: 0, 0
	chOff := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "chOff", "a")
	chOff.SetAttribute(openxml.NewAttribute("", "x", "", "0"))
	chOff.SetAttribute(openxml.NewAttribute("", "y", "", "0"))
	xfrm.AppendChild(chOff)

	m := FromTransform2D(xfrm)

	// Point (100, 100) in child space should be:
	// 1. Scaled by 0.5 (1000/2000, 800/1600) -> (50, 50)
	// 2. Translated by (100, 200) -> (150, 250)
	x, y := m.TransformPoint(100, 100)
	if math.Abs(x-150) > epsilon || math.Abs(y-250) > epsilon {
		t.Errorf("Viewport transform of (100, 100) = (%f, %f), want (150, 250)", x, y)
	}
}

func TestDecomposeTransform(t *testing.T) {
	// Create a composite transform: translate(10, 20) * rotate(30°) * scale(2, 3)
	translate := Translate(10, 20)
	rotate := Rotate(math.Pi / 6) // 30 degrees
	scale := Scale(2, 3)

	m := Multiply(translate, Multiply(rotate, scale))

	// Decompose
	comp := DecomposeTransform(m)

	// Check translation
	if math.Abs(comp.TranslateX-10) > epsilon || math.Abs(comp.TranslateY-20) > epsilon {
		t.Errorf("Decomposed translation = (%f, %f), want (10, 20)", comp.TranslateX, comp.TranslateY)
	}

	// Check rotation
	if math.Abs(comp.Rotation-math.Pi/6) > epsilon {
		t.Errorf("Decomposed rotation = %f, want %f (30°)", comp.Rotation, math.Pi/6)
	}

	// Check scale
	if math.Abs(comp.ScaleX-2) > epsilon || math.Abs(comp.ScaleY-3) > epsilon {
		t.Errorf("Decomposed scale = (%f, %f), want (2, 3)", comp.ScaleX, comp.ScaleY)
	}
}

func TestComputeAbsoluteTransform(t *testing.T) {
	// Create a simple hierarchy: root -> group -> shape
	root := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "spTree", "p")

	// Group with translation (100, 200)
	group := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSp", "p")
	grpSpPr := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSpPr", "p")
	groupXfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	groupOff := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	groupOff.SetAttribute(openxml.NewAttribute("", "x", "", "100"))
	groupOff.SetAttribute(openxml.NewAttribute("", "y", "", "200"))
	groupXfrm.AppendChild(groupOff)

	groupExt := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	groupExt.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	groupExt.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	groupXfrm.AppendChild(groupExt)

	grpSpPr.AppendChild(groupXfrm)
	group.AppendChild(grpSpPr)
	root.AppendChild(group)

	// Shape with translation (10, 20) relative to group
	shape := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "sp", "p")
	spPr := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "spPr", "a")
	shapeXfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	shapeOff := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	shapeOff.SetAttribute(openxml.NewAttribute("", "x", "", "10"))
	shapeOff.SetAttribute(openxml.NewAttribute("", "y", "", "20"))
	shapeXfrm.AppendChild(shapeOff)

	shapeExt := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	shapeExt.SetAttribute(openxml.NewAttribute("", "cx", "", "50"))
	shapeExt.SetAttribute(openxml.NewAttribute("", "cy", "", "50"))
	shapeXfrm.AppendChild(shapeExt)

	spPr.AppendChild(shapeXfrm)
	shape.AppendChild(spPr)
	group.AppendChild(shape)

	// Compute absolute transform for shape
	m := ComputeAbsoluteTransform(shape, root)

	// Shape at (10, 20) in group space should be at (110, 220) in absolute space
	x, y := m.TransformPoint(0, 0)
	if math.Abs(x-110) > epsilon || math.Abs(y-220) > epsilon {
		t.Errorf("Absolute transform of (0, 0) = (%f, %f), want (110, 220)", x, y)
	}
}

func TestLocalToAbsolute(t *testing.T) {
	// Create simple hierarchy for testing
	root := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "spTree", "p")
	group := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSp", "p")
	grpSpPr := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSpPr", "p")
	groupXfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	groupOff := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	groupOff.SetAttribute(openxml.NewAttribute("", "x", "", "100"))
	groupOff.SetAttribute(openxml.NewAttribute("", "y", "", "200"))
	groupXfrm.AppendChild(groupOff)

	groupExt := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	groupExt.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	groupExt.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	groupXfrm.AppendChild(groupExt)

	grpSpPr.AppendChild(groupXfrm)
	group.AppendChild(grpSpPr)
	root.AppendChild(group)

	// Convert (10, 20) in group's local space to absolute
	x, y := LocalToAbsolute(10, 20, group, root)
	if math.Abs(x-110) > epsilon || math.Abs(y-220) > epsilon {
		t.Errorf("LocalToAbsolute(10, 20) = (%f, %f), want (110, 220)", x, y)
	}
}

func TestAbsoluteToLocal(t *testing.T) {
	// Create simple hierarchy for testing
	root := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "spTree", "p")
	group := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSp", "p")
	grpSpPr := openxml.NewCompositeElement("http://schemas.openxmlformats.org/presentationml/2006/main", "grpSpPr", "p")
	groupXfrm := openxml.NewCompositeElement("http://schemas.openxmlformats.org/drawingml/2006/main", "xfrm", "a")

	groupOff := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "off", "a")
	groupOff.SetAttribute(openxml.NewAttribute("", "x", "", "100"))
	groupOff.SetAttribute(openxml.NewAttribute("", "y", "", "200"))
	groupXfrm.AppendChild(groupOff)

	groupExt := openxml.NewLeafElement("http://schemas.openxmlformats.org/drawingml/2006/main", "ext", "a")
	groupExt.SetAttribute(openxml.NewAttribute("", "cx", "", "1000"))
	groupExt.SetAttribute(openxml.NewAttribute("", "cy", "", "800"))
	groupXfrm.AppendChild(groupExt)

	grpSpPr.AppendChild(groupXfrm)
	group.AppendChild(grpSpPr)
	root.AppendChild(group)

	// Convert (110, 220) in absolute space to group's local space
	x, y := AbsoluteToLocal(110, 220, group, root)
	if math.Abs(x-10) > epsilon || math.Abs(y-20) > epsilon {
		t.Errorf("AbsoluteToLocal(110, 220) = (%f, %f), want (10, 20)", x, y)
	}
}
