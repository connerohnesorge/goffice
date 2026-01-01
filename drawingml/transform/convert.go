package transform

import (
	"math"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Constants for parsing and conversion.
const (
	base10              = 10
	rotationDivisor     = 60000.0
	drawingMLNamespace  = "http://schemas.openxmlformats.org/drawingml/2006/main"
	presentationMLNS    = "http://schemas.openxmlformats.org/presentationml/2006/main"
)

// FromTransform2D converts a DrawingML Transform2D element to a transformation Matrix.
// For groups (when ChildOffset/ChildExtent are present), this includes viewport transformation.
// Returns Identity matrix if xfrm is nil.
//
// The transformation order is: viewport → flip → rotate → translate
// This matches the OOXML specification for group transforms.
//
//nolint:revive // function-length: complex transform logic requires many statements
func FromTransform2D(xfrm *openxml.CompositeElementBase) Matrix {
	if xfrm == nil {
		return Identity()
	}

	// Get offset (translation) - required child element a:off
	tx, ty := 0.0, 0.0
	if offElem := xfrm.GetElement("off", drawingMLNamespace); offElem != nil {
		if xAttr, found := offElem.GetAttribute("x", ""); found {
			if val, err := strconv.ParseInt(xAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				tx = float64(val)
			}
		}
		if yAttr, found := offElem.GetAttribute("y", ""); found {
			if val, err := strconv.ParseInt(yAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				ty = float64(val)
			}
		}
	}

	// Get rotation (stored as attribute on xfrm element, in 1/60000 degrees)
	angle := 0.0
	if rotAttr, found := xfrm.GetAttribute("rot", ""); found {
		if rotVal, err := strconv.ParseInt(rotAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
			// Convert from 1/60000 degrees to radians
			angle = float64(rotVal) / rotationDivisor * math.Pi / 180.0 //nolint:revive // 180 is degrees constant
		}
	}

	// Get flip attributes
	flipH := false
	flipV := false
	if flipHAttr, found := xfrm.GetAttribute("flipH", ""); found {
		flipH = flipHAttr.Value() == "1" || flipHAttr.Value() == "true"
	}
	if flipVAttr, found := xfrm.GetAttribute("flipV", ""); found {
		flipV = flipVAttr.Value() == "1" || flipVAttr.Value() == "true"
	}

	flipScaleX := 1.0
	flipScaleY := 1.0
	if flipH {
		flipScaleX = -1.0
	}
	if flipV {
		flipScaleY = -1.0
	}

	// Get extent (size) - required child element a:ext
	var extCx, extCy float64
	if extElem := xfrm.GetElement("ext", drawingMLNamespace); extElem != nil {
		if cxAttr, found := extElem.GetAttribute("cx", ""); found {
			if val, err := strconv.ParseInt(cxAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				extCx = float64(val)
			}
		}
		if cyAttr, found := extElem.GetAttribute("cy", ""); found {
			if val, err := strconv.ParseInt(cyAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				extCy = float64(val)
			}
		}
	}

	// Get viewport transformation (child coordinate space - for groups only)
	// ChildOffset (a:chOff) and ChildExtent (a:chExt) are optional child elements
	viewportMatrix := Identity()

	var chExtCx, chExtCy float64
	hasChExt := false
	if chExtElem := xfrm.GetElement("chExt", drawingMLNamespace); chExtElem != nil {
		hasChExt = true
		if cxAttr, found := chExtElem.GetAttribute("cx", ""); found {
			if val, err := strconv.ParseInt(cxAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				chExtCx = float64(val)
			}
		}
		if cyAttr, found := chExtElem.GetAttribute("cy", ""); found {
			if val, err := strconv.ParseInt(cyAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				chExtCy = float64(val)
			}
		}
	}

	var chOffX, chOffY float64
	hasChOff := false
	if chOffElem := xfrm.GetElement("chOff", drawingMLNamespace); chOffElem != nil {
		hasChOff = true
		if xAttr, found := chOffElem.GetAttribute("x", ""); found {
			if val, err := strconv.ParseInt(xAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				chOffX = float64(val)
			}
		}
		if yAttr, found := chOffElem.GetAttribute("y", ""); found {
			if val, err := strconv.ParseInt(yAttr.Value(), base10, 64); err == nil { //nolint:revive // 64-bit is standard
				chOffY = float64(val)
			}
		}
	}

	// If this is a group transform with viewport
	if hasChExt && extCx != 0 && extCy != 0 && chExtCx != 0 && chExtCy != 0 {
		// Scale factor: map ChildExtent to Extent
		scaleX := extCx / chExtCx
		scaleY := extCy / chExtCy

		// Translation: ChildOffset defines child space origin
		transX := 0.0
		transY := 0.0
		if hasChOff {
			transX = chOffX
			transY = chOffY
		}

		// Viewport transform: translate by negative child offset, then scale
		// This maps child coordinates to group's local space
		// P_group = (P_child - chOff) * scale
		viewportMatrix = Multiply(Scale(scaleX, scaleY), Translate(-transX, -transY))
	}

	// Build composite transform: translate * rotate * flip * viewport
	// Viewport is applied FIRST (rightmost) because it operates on child coordinates
	trans := Translate(tx, ty)
	rotate := Rotate(angle)
	flip := Scale(flipScaleX, flipScaleY)

	return Multiply(trans, Multiply(rotate, Multiply(flip, viewportMatrix)))
}

// ComputeAbsoluteTransform traverses the parent chain and composes transforms.
// This handles both the element's own transform and all parent group transforms.
// Returns the cumulative transformation matrix from the element to the root.
//
// elem: The element to compute the absolute transform for
// root: The root element (typically a ShapeTree or slide)
func ComputeAbsoluteTransform(elem, root openxml.Element) Matrix {
	var matrices []Matrix

	// First, get the element's own transform (if it has one)
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		if xfrm := getTransform(comp); xfrm != nil {
			matrices = append(matrices, FromTransform2D(xfrm))
		}
	}

	// Walk up parent chain collecting GROUP transforms only
	current := elem.Parent()
	for current != nil && current != root {
		comp, ok := current.(*openxml.CompositeElementBase)
		if ok && current.LocalName() == "grpSp" {
			if xfrm := getTransform(comp); xfrm != nil {
				// Prepend parent transform (parents come before children in multiplication)
				matrices = append([]Matrix{FromTransform2D(xfrm)}, matrices...)
			}
		}
		current = current.Parent()
	}

	// Multiply all transforms (parent → child order)
	// Result: T_parent * T_child means parent transform applies to child's transformed space
	result := Identity()
	for _, m := range matrices {
		result = Multiply(result, m)
	}

	return result
}

// getTransform extracts the transform element from a shape or group shape.
// Returns nil if no transform is found.
func getTransform(elem *openxml.CompositeElementBase) *openxml.CompositeElementBase {
	// For shapes: look for a:spPr/a:xfrm
	if xfrm := getTransformFromSpPr(elem); xfrm != nil {
		return xfrm
	}

	// For group shapes: look for p:grpSpPr/a:xfrm
	return getTransformFromGrpSpPr(elem)
}

func getTransformFromSpPr(elem *openxml.CompositeElementBase) *openxml.CompositeElementBase {
	spPr := elem.GetElement("spPr", drawingMLNamespace)
	if spPr == nil {
		return nil
	}
	compSpPr, ok := spPr.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}
	xfrm := compSpPr.GetElement("xfrm", drawingMLNamespace)
	if xfrm == nil {
		return nil
	}
	comp, ok := xfrm.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return comp
}

func getTransformFromGrpSpPr(elem *openxml.CompositeElementBase) *openxml.CompositeElementBase {
	grpSpPr := elem.GetElement("grpSpPr", presentationMLNS)
	if grpSpPr == nil {
		return nil
	}
	compGrpSpPr, ok := grpSpPr.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}
	xfrm := compGrpSpPr.GetElement("xfrm", drawingMLNamespace)
	if xfrm == nil {
		return nil
	}
	comp, ok := xfrm.(*openxml.CompositeElementBase)
	if !ok {
		return nil
	}

	return comp
}

// LocalToAbsolute converts a point from element's local coordinate space to absolute coordinates.
// This applies all parent group transforms.
//
// x, y: Coordinates in the element's local space
// elem: The element whose coordinate space we're converting from
// root: The root element (typically a ShapeTree or slide)
func LocalToAbsolute(x, y float64, elem, root openxml.Element) (absX, absY float64) {
	matrix := ComputeAbsoluteTransform(elem, root)

	return matrix.TransformPoint(x, y)
}

// AbsoluteToLocal converts a point from absolute coordinates to element's local coordinate space.
// This requires inverting all parent group transforms.
//
// x, y: Coordinates in absolute space
// elem: The element whose coordinate space we're converting to
// root: The root element (typically a ShapeTree or slide)
func AbsoluteToLocal(x, y float64, elem, root openxml.Element) (localX, localY float64) {
	matrix := ComputeAbsoluteTransform(elem, root)
	invMatrix := matrix.Invert()

	return invMatrix.TransformPoint(x, y)
}

// TransformComponents holds the decomposed parts of a transformation matrix.
type TransformComponents struct {
	TranslateX float64
	TranslateY float64
	Rotation   float64 // in radians
	ScaleX     float64
	ScaleY     float64
}

// DecomposeTransform extracts translation, rotation, and scale components from a matrix.
//
// Note: This assumes the matrix was constructed from simple transformations
// (translate, rotate, scale). Complex matrices with shear may not decompose accurately.
func DecomposeTransform(m Matrix) TransformComponents {
	// Translation is stored directly in E and F
	tx := m.E
	ty := m.F

	// Scale is the magnitude of the basis vectors
	// X basis vector: (A, B)
	// Y basis vector: (C, D)
	scaleX := math.Sqrt(m.A*m.A + m.B*m.B)
	scaleY := math.Sqrt(m.C*m.C + m.D*m.D)

	// Rotation is the angle of the X basis vector
	// atan2(B, A) gives the angle
	rotation := math.Atan2(m.B, m.A)

	return TransformComponents{
		TranslateX: tx,
		TranslateY: ty,
		Rotation:   rotation,
		ScaleX:     scaleX,
		ScaleY:     scaleY,
	}
}
