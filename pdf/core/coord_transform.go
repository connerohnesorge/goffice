// Package core provides PDF document creation and low-level PDF primitives.
package core

import (
	"math"
)

// CoordTransformer provides coordinate transformation from OOXML to PDF.
// OOXML uses top-left origin with EMU units, while PDF uses bottom-left origin
// with points. This transformer handles the conversion between these systems.
type CoordTransformer struct {
	// pageHeight is the page height in points, used for Y-axis flipping.
	pageHeight float64
	// scaleX and scaleY allow for additional scaling factors.
	scaleX, scaleY float64
	// offsetX and offsetY provide translation after coordinate conversion.
	offsetX, offsetY float64
}

// NewCoordTransformer creates a new coordinate transformer for OOXML to PDF conversion.
// pageHeight is the PDF page height in points, needed to flip the Y-axis.
func NewCoordTransformer(
	pageHeight float64,
) *CoordTransformer {
	return &CoordTransformer{
		pageHeight: pageHeight,
		scaleX:     1.0,
		scaleY:     1.0,
		offsetX:    0,
		offsetY:    0,
	}
}

// SetScale sets the scale factors for coordinate conversion.
func (ct *CoordTransformer) SetScale(
	scaleX, scaleY float64,
) *CoordTransformer {
	ct.scaleX = scaleX
	ct.scaleY = scaleY

	return ct
}

// SetOffset sets the translation offset in PDF points.
func (ct *CoordTransformer) SetOffset(
	offsetX, offsetY float64,
) *CoordTransformer {
	ct.offsetX = offsetX
	ct.offsetY = offsetY

	return ct
}

// SetPageHeight sets the page height for Y-axis flipping.
func (ct *CoordTransformer) SetPageHeight(
	height float64,
) *CoordTransformer {
	ct.pageHeight = height

	return ct
}

// PageHeight returns the current page height.
func (ct *CoordTransformer) PageHeight() float64 {
	return ct.pageHeight
}

// EMUToPoint converts an OOXML EMU coordinate to PDF points.
// This handles the EMU to points conversion without Y-axis flip.
func (ct *CoordTransformer) EMUToPoint(
	emuX, emuY int64,
) (x, y float64) {
	pdfX := float64(
		emuX,
	)*PointsPerEMU*ct.scaleX + ct.offsetX
	pdfY := float64(
		emuY,
	)*PointsPerEMU*ct.scaleY + ct.offsetY

	return pdfX, pdfY
}

// EMUToPDFCoord converts OOXML EMU coordinates to PDF coordinates.
// This performs the full conversion including:
// - EMU to points conversion
// - Y-axis flip (OOXML top-left to PDF bottom-left)
// - Scale and offset application
func (ct *CoordTransformer) EMUToPDFCoord(
	emuX, emuY int64,
) (x, y float64) {
	pdfX := float64(
		emuX,
	)*PointsPerEMU*ct.scaleX + ct.offsetX
	// Flip Y-axis: PDF Y increases upward from bottom, OOXML Y increases downward from top
	pdfY := ct.pageHeight - (float64(emuY)*PointsPerEMU*ct.scaleY + ct.offsetY)

	return pdfX, pdfY
}

// PointToPDFCoord converts a point in OOXML coordinate space to PDF coordinates.
// The input is in points with top-left origin, output is in points with bottom-left origin.
func (ct *CoordTransformer) PointToPDFCoord(
	x, y float64,
) (pdfX, pdfY float64) {
	outX := x*ct.scaleX + ct.offsetX
	outY := ct.pageHeight - (y*ct.scaleY + ct.offsetY)

	return outX, outY
}

// EMUToPDFSize converts OOXML EMU dimensions to PDF points.
// Unlike coordinates, dimensions don't need Y-axis flipping.
func (ct *CoordTransformer) EMUToPDFSize(
	emuWidth, emuHeight int64,
) (float64, float64) {
	pdfWidth := float64(
		emuWidth,
	) * PointsPerEMU * ct.scaleX
	pdfHeight := float64(
		emuHeight,
	) * PointsPerEMU * ct.scaleY

	return pdfWidth, pdfHeight
}

// OOXMLTransform represents a DrawingML xfrm (2D transform) element.
// This captures all transformation properties from OOXML.
type OOXMLTransform struct {
	// OffsetX and OffsetY are the position in EMU (from parent or page origin).
	OffsetX, OffsetY int64
	// ExtentCx and ExtentCy are the width and height in EMU.
	ExtentCx, ExtentCy int64
	// Rotation is the rotation angle in 60,000ths of a degree (clockwise).
	Rotation int32
	// FlipH indicates horizontal flip.
	FlipH bool
	// FlipV indicates vertical flip.
	FlipV bool
}

// NewOOXMLTransform creates a new OOXMLTransform with the specified position and size.
func NewOOXMLTransform(
	offX, offY, extCx, extCy int64,
) OOXMLTransform {
	return OOXMLTransform{
		OffsetX:  offX,
		OffsetY:  offY,
		ExtentCx: extCx,
		ExtentCy: extCy,
		Rotation: 0,
		FlipH:    false,
		FlipV:    false,
	}
}

// WithRotation sets the rotation angle in 60,000ths of a degree.
func (t OOXMLTransform) WithRotation(
	rot int32,
) OOXMLTransform {
	t.Rotation = rot

	return t
}

// WithRotationDegrees sets the rotation angle in degrees.
func (t OOXMLTransform) WithRotationDegrees(
	degrees float64,
) OOXMLTransform {
	t.Rotation = int32(
		math.Round(degrees * 60000.0),
	)

	return t
}

// WithFlipH sets the horizontal flip state.
func (t OOXMLTransform) WithFlipH(
	flip bool,
) OOXMLTransform {
	t.FlipH = flip

	return t
}

// WithFlipV sets the vertical flip state.
func (t OOXMLTransform) WithFlipV(
	flip bool,
) OOXMLTransform {
	t.FlipV = flip

	return t
}

// RotationDegrees returns the rotation in degrees.
func (t OOXMLTransform) RotationDegrees() float64 {
	return float64(t.Rotation) / 60000.0
}

// RotationRadians returns the rotation in radians.
func (t OOXMLTransform) RotationRadians() float64 {
	return t.RotationDegrees() * math.Pi / 180.0
}

// HasRotation returns true if there is any rotation applied.
func (t OOXMLTransform) HasRotation() bool {
	return t.Rotation != 0
}

// HasFlip returns true if either horizontal or vertical flip is applied.
func (t OOXMLTransform) HasFlip() bool {
	return t.FlipH || t.FlipV
}

// HasTransform returns true if any transformation (rotation or flip) is applied.
func (t OOXMLTransform) HasTransform() bool {
	return t.HasRotation() || t.HasFlip()
}

// OffsetPoints returns the offset in PDF points.
func (t OOXMLTransform) OffsetPoints() (float64, float64) {
	return float64(
			t.OffsetX,
		) * PointsPerEMU, float64(
			t.OffsetY,
		) * PointsPerEMU
}

// ExtentPoints returns the extent (size) in PDF points.
func (t OOXMLTransform) ExtentPoints() (float64, float64) {
	return float64(
			t.ExtentCx,
		) * PointsPerEMU, float64(
			t.ExtentCy,
		) * PointsPerEMU
}

// CenterPoints returns the center point in PDF points (relative to offset).
func (t OOXMLTransform) CenterPoints() (float64, float64) {
	offX, offY := t.OffsetPoints()
	extCx, extCy := t.ExtentPoints()

	return offX + extCx/2, offY + extCy/2
}

// PDFMatrix represents a 2D affine transformation matrix for PDF.
// The matrix is [a b c d e f] which transforms (x, y) to:
//
//	x' = a*x + c*y + e
//	y' = b*x + d*y + f
type PDFMatrix struct {
	A, B, C, D, E, F float64
}

// IdentityMatrix returns the identity transformation matrix.
func IdentityMatrix() PDFMatrix {
	return PDFMatrix{
		A: 1, B: 0,
		C: 0, D: 1,
		E: 0, F: 0,
	}
}

// TranslationMatrix creates a translation matrix.
func TranslationMatrix(tx, ty float64) PDFMatrix {
	return PDFMatrix{
		A: 1, B: 0,
		C: 0, D: 1,
		E: tx, F: ty,
	}
}

// ScaleMatrix creates a scaling matrix.
func ScaleMatrix(sx, sy float64) PDFMatrix {
	return PDFMatrix{
		A: sx, B: 0,
		C: 0, D: sy,
		E: 0, F: 0,
	}
}

// RotationMatrix creates a rotation matrix.
// The angle is in radians, counter-clockwise.
func RotationMatrix(radians float64) PDFMatrix {
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	return PDFMatrix{
		A: cos, B: sin,
		C: -sin, D: cos,
		E: 0, F: 0,
	}
}

// Multiply returns the matrix multiplication of this matrix with another.
// The result applies other first, then this matrix: result = this * other.
func (m PDFMatrix) Multiply(
	other PDFMatrix,
) PDFMatrix {
	return PDFMatrix{
		A: m.A*other.A + m.C*other.B,
		B: m.B*other.A + m.D*other.B,
		C: m.A*other.C + m.C*other.D,
		D: m.B*other.C + m.D*other.D,
		E: m.A*other.E + m.C*other.F + m.E,
		F: m.B*other.E + m.D*other.F + m.F,
	}
}

// Translate returns a new matrix with translation applied.
func (m PDFMatrix) Translate(
	tx, ty float64,
) PDFMatrix {
	return m.Multiply(TranslationMatrix(tx, ty))
}

// Scale returns a new matrix with scaling applied.
func (m PDFMatrix) Scale(
	sx, sy float64,
) PDFMatrix {
	return m.Multiply(ScaleMatrix(sx, sy))
}

// Rotate returns a new matrix with rotation applied.
// The angle is in radians, counter-clockwise.
func (m PDFMatrix) Rotate(
	radians float64,
) PDFMatrix {
	return m.Multiply(RotationMatrix(radians))
}

// RotateDegrees returns a new matrix with rotation applied.
// The angle is in degrees, counter-clockwise.
func (m PDFMatrix) RotateDegrees(
	degrees float64,
) PDFMatrix {
	return m.Rotate(degrees * math.Pi / 180.0)
}

// FlipHorizontal returns a new matrix with horizontal flip applied.
func (m PDFMatrix) FlipHorizontal() PDFMatrix {
	return m.Scale(-1, 1)
}

// FlipVertical returns a new matrix with vertical flip applied.
func (m PDFMatrix) FlipVertical() PDFMatrix {
	return m.Scale(1, -1)
}

// TransformPoint applies the matrix to a point.
func (m PDFMatrix) TransformPoint(
	x, y float64,
) (float64, float64) {
	return m.A*x + m.C*y + m.E, m.B*x + m.D*y + m.F
}

// TransformVector applies the matrix to a vector (ignores translation).
func (m PDFMatrix) TransformVector(
	x, y float64,
) (float64, float64) {
	return m.A*x + m.C*y, m.B*x + m.D*y
}

// Determinant returns the determinant of the matrix.
func (m PDFMatrix) Determinant() float64 {
	return m.A*m.D - m.B*m.C
}

// Inverse returns the inverse of the matrix.
// If the matrix is singular (determinant is zero), returns the identity matrix.
func (m PDFMatrix) Inverse() PDFMatrix {
	det := m.Determinant()
	if math.Abs(det) < 1e-10 {
		return IdentityMatrix()
	}
	invDet := 1.0 / det

	return PDFMatrix{
		A: m.D * invDet,
		B: -m.B * invDet,
		C: -m.C * invDet,
		D: m.A * invDet,
		E: (m.C*m.F - m.D*m.E) * invDet,
		F: (m.B*m.E - m.A*m.F) * invDet,
	}
}

// IsIdentity returns true if this is the identity matrix.
func (m PDFMatrix) IsIdentity() bool {
	const epsilon = 1e-10

	return math.Abs(m.A-1) < epsilon &&
		math.Abs(m.B) < epsilon &&
		math.Abs(m.C) < epsilon &&
		math.Abs(m.D-1) < epsilon &&
		math.Abs(m.E) < epsilon &&
		math.Abs(m.F) < epsilon
}

// Array returns the matrix as a 6-element array [a, b, c, d, e, f].
func (m PDFMatrix) Array() [6]float64 {
	return [6]float64{
		m.A,
		m.B,
		m.C,
		m.D,
		m.E,
		m.F,
	}
}

// OOXMLTransformToPDFMatrix converts an OOXML transform to a PDF transformation matrix.
// The resulting matrix is in PDF coordinate space and should be applied after
// coordinate conversion from OOXML to PDF space.
//
// The transformation order for DrawingML is:
// 1. Translate to center of shape
// 2. Apply flip (if any)
// 3. Apply rotation
// 4. Translate back from center
//
// Note: OOXML rotation is clockwise, but PDF rotation is counter-clockwise,
// so we negate the angle.
func OOXMLTransformToPDFMatrix(
	xfrm OOXMLTransform,
	pageHeight float64,
) PDFMatrix {
	// Get dimensions in points
	offX, offY := xfrm.OffsetPoints()
	extCx, extCy := xfrm.ExtentPoints()

	// Convert Y coordinates for PDF coordinate system (flip Y axis)
	pdfOffY := pageHeight - offY - extCy
	centerX := offX + extCx/2
	centerY := pdfOffY + extCy/2

	// Start with identity
	matrix := IdentityMatrix()

	// If no transformation needed, just return identity
	if !xfrm.HasTransform() {
		return matrix
	}

	// Apply transformations around the center point:
	// 1. Translate so center is at origin
	matrix = matrix.Translate(-centerX, -centerY)

	// 2. Apply flips (in shape's local coordinate space)
	if xfrm.FlipH {
		matrix = matrix.FlipHorizontal()
	}
	if xfrm.FlipV {
		matrix = matrix.FlipVertical()
	}

	// 3. Apply rotation (negate because OOXML is clockwise, PDF is counter-clockwise)
	if xfrm.HasRotation() {
		rotRadians := -xfrm.RotationRadians()
		matrix = matrix.Rotate(rotRadians)
	}

	// 4. Translate center back
	matrix = matrix.Translate(centerX, centerY)

	return matrix
}

// OOXMLToPDFBounds converts an OOXML bounding box to PDF coordinates.
// Returns (x, y, width, height) in PDF coordinate space.
func (ct *CoordTransformer) OOXMLToPDFBounds(
	xfrm OOXMLTransform,
) (x, y, width, height float64) {
	width, height = ct.EMUToPDFSize(
		xfrm.ExtentCx,
		xfrm.ExtentCy,
	)
	x, y = ct.EMUToPDFCoord(
		xfrm.OffsetX,
		xfrm.OffsetY,
	)
	// Adjust Y to be the bottom-left corner of the rectangle
	y -= height

	return x, y, width, height
}

// TransformedOOXMLBounds calculates the actual bounding box in PDF coordinates
// after applying rotation and flips. This is useful for determining the area
// that needs to be reserved for a rotated shape.
func TransformedOOXMLBounds(
	xfrm OOXMLTransform,
	pageHeight float64,
) (x, y, width, height float64) {
	// Get the base dimensions in points
	offX, offY := xfrm.OffsetPoints()
	extCx, extCy := xfrm.ExtentPoints()

	// Convert to PDF coordinates (flip Y)
	pdfOffY := pageHeight - offY - extCy

	// If no rotation, just return the basic bounds
	if !xfrm.HasRotation() {
		return offX, pdfOffY, extCx, extCy
	}

	// Calculate the corners of the unrotated rectangle
	corners := []struct{ x, y float64 }{
		{offX, pdfOffY},
		{offX + extCx, pdfOffY},
		{offX + extCx, pdfOffY + extCy},
		{offX, pdfOffY + extCy},
	}

	// Get the center point for rotation
	centerX := offX + extCx/2
	centerY := pdfOffY + extCy/2

	// Rotation angle (negate for PDF counter-clockwise)
	angle := -xfrm.RotationRadians()
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	// Rotate each corner and find the bounding box
	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	minY, maxY := math.MaxFloat64, -math.MaxFloat64

	for _, corner := range corners {
		// Translate to origin, rotate, translate back
		dx := corner.x - centerX
		dy := corner.y - centerY
		rotX := dx*cos - dy*sin + centerX
		rotY := dx*sin + dy*cos + centerY

		minX = math.Min(minX, rotX)
		maxX = math.Max(maxX, rotX)
		minY = math.Min(minY, rotY)
		maxY = math.Max(maxY, rotY)
	}

	return minX, minY, maxX - minX, maxY - minY
}

// NestedTransformer handles nested transformations (e.g., shapes inside groups).
// It maintains a stack of transformations to properly handle coordinate
// conversions through multiple levels of nesting.
type NestedTransformer struct {
	*CoordTransformer
	stack []PDFMatrix
}

// NewNestedTransformer creates a new nested transformer with the given page height.
func NewNestedTransformer(
	pageHeight float64,
) *NestedTransformer {
	return &NestedTransformer{
		CoordTransformer: NewCoordTransformer(
			pageHeight,
		),
		stack: []PDFMatrix{
			IdentityMatrix(),
		},
	}
}

// PushTransform pushes a new transformation onto the stack.
// Subsequent transformations will be relative to this transform.
func (nt *NestedTransformer) PushTransform(
	xfrm OOXMLTransform,
) {
	current := nt.CurrentMatrix()
	newMatrix := OOXMLTransformToPDFMatrix(
		xfrm,
		nt.pageHeight,
	)
	combined := current.Multiply(newMatrix)
	nt.stack = append(nt.stack, combined)
}

// PopTransform removes the most recent transformation from the stack.
func (nt *NestedTransformer) PopTransform() {
	if len(nt.stack) > 1 {
		nt.stack = nt.stack[:len(nt.stack)-1]
	}
}

// CurrentMatrix returns the current combined transformation matrix.
func (nt *NestedTransformer) CurrentMatrix() PDFMatrix {
	if len(nt.stack) == 0 {
		return IdentityMatrix()
	}

	return nt.stack[len(nt.stack)-1]
}

// TransformPoint applies the current transformation stack to a point.
func (nt *NestedTransformer) TransformPoint(
	x, y float64,
) (float64, float64) {
	return nt.CurrentMatrix().TransformPoint(x, y)
}

// Depth returns the current nesting depth (0 = no nesting).
func (nt *NestedTransformer) Depth() int {
	return len(nt.stack) - 1
}

// Reset clears the transformation stack.
func (nt *NestedTransformer) Reset() {
	nt.stack = []PDFMatrix{IdentityMatrix()}
}

// GroupTransform represents a group-level transformation in DrawingML.
// Groups have their own coordinate space that contains child shapes.
type GroupTransform struct {
	// ChildOffset is the offset of the child coordinate space origin (in EMU).
	ChildOffsetX, ChildOffsetY int64
	// ChildExtent is the extent of the child coordinate space (in EMU).
	ChildExtentCx, ChildExtentCy int64
	// Transform is the group's own transformation.
	Transform OOXMLTransform
}

// NewGroupTransform creates a new group transform.
func NewGroupTransform(
	childOffX, childOffY, childExtCx, childExtCy int64,
	groupXfrm OOXMLTransform,
) GroupTransform {
	return GroupTransform{
		ChildOffsetX:  childOffX,
		ChildOffsetY:  childOffY,
		ChildExtentCx: childExtCx,
		ChildExtentCy: childExtCy,
		Transform:     groupXfrm,
	}
}

// ChildToGroupMatrix creates a transformation matrix that converts coordinates
// from the child coordinate space to the group's space.
//
// The child coordinate space is defined by chOff and chExt in the group transform,
// and maps to the group's xfrm offset and extent.
//
// The transformation order when applied to a point P is:
// 1. Subtract child offset (move origin to child (0,0))
// 2. Scale from child extent to group extent
// 3. Add group offset
//
// P' = GroupOffset + Scale * (P - ChildOffset)
func (gt GroupTransform) ChildToGroupMatrix() PDFMatrix {
	// Calculate scale factors
	var scaleX, scaleY float64
	if gt.ChildExtentCx != 0 {
		scaleX = float64(
			gt.Transform.ExtentCx,
		) / float64(
			gt.ChildExtentCx,
		)
	} else {
		scaleX = 1.0
	}
	if gt.ChildExtentCy != 0 {
		scaleY = float64(
			gt.Transform.ExtentCy,
		) / float64(
			gt.ChildExtentCy,
		)
	} else {
		scaleY = 1.0
	}

	// Convert offsets to points
	childOffX := float64(
		gt.ChildOffsetX,
	) * PointsPerEMU
	childOffY := float64(
		gt.ChildOffsetY,
	) * PointsPerEMU
	groupOffX := float64(
		gt.Transform.OffsetX,
	) * PointsPerEMU
	groupOffY := float64(
		gt.Transform.OffsetY,
	) * PointsPerEMU

	// Build the matrix directly for the transformation:
	// P' = GroupOffset + Scale * (P - ChildOffset)
	// P' = Scale*P - Scale*ChildOffset + GroupOffset
	//
	// The matrix [a b c d e f] transforms (x, y) as:
	// x' = a*x + c*y + e
	// y' = b*x + d*y + f
	//
	// For our case (no rotation/shear):
	// x' = scaleX*x + (-scaleX*childOffX + groupOffX)
	// y' = scaleY*y + (-scaleY*childOffY + groupOffY)
	return PDFMatrix{
		A: scaleX,
		B: 0,
		C: 0,
		D: scaleY,
		E: -scaleX*childOffX + groupOffX,
		F: -scaleY*childOffY + groupOffY,
	}
}

// PushGroupTransform pushes a group transformation onto the nested transformer.
// This combines the child-to-group mapping with the group's own transformation.
func (nt *NestedTransformer) PushGroupTransform(
	gt GroupTransform,
) {
	// First, apply the child-to-group coordinate mapping
	childToGroup := gt.ChildToGroupMatrix()
	current := nt.CurrentMatrix()
	combined := current.Multiply(childToGroup)

	// Then apply the group's own transformation (rotation, flip)
	if gt.Transform.HasTransform() {
		groupMatrix := OOXMLTransformToPDFMatrix(
			gt.Transform,
			nt.pageHeight,
		)
		combined = combined.Multiply(groupMatrix)
	}

	nt.stack = append(nt.stack, combined)
}
