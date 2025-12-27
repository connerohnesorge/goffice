package core

import (
	"math"
	"testing"
)

// floatEqualsWithTolerance compares two floats with a custom epsilon tolerance.
func floatEqualsWithTolerance(
	a, b, epsilon float64,
) bool {
	return math.Abs(a-b) < epsilon
}

func TestNewCoordTransformer(t *testing.T) {
	ct := NewCoordTransformer(
		792.0,
	) // Letter height in points

	if ct.pageHeight != 792.0 {
		t.Errorf(
			"pageHeight = %v, want 792.0",
			ct.pageHeight,
		)
	}
	if ct.scaleX != 1.0 || ct.scaleY != 1.0 {
		t.Errorf(
			"scale = (%v, %v), want (1.0, 1.0)",
			ct.scaleX,
			ct.scaleY,
		)
	}
	if ct.offsetX != 0 || ct.offsetY != 0 {
		t.Errorf(
			"offset = (%v, %v), want (0, 0)",
			ct.offsetX,
			ct.offsetY,
		)
	}
}

func TestCoordTransformer_SetScale(t *testing.T) {
	ct := NewCoordTransformer(792.0)
	ct.SetScale(2.0, 0.5)

	if ct.scaleX != 2.0 || ct.scaleY != 0.5 {
		t.Errorf(
			"scale = (%v, %v), want (2.0, 0.5)",
			ct.scaleX,
			ct.scaleY,
		)
	}
}

func TestCoordTransformer_SetOffset(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)
	ct.SetOffset(72.0, 36.0)

	if ct.offsetX != 72.0 || ct.offsetY != 36.0 {
		t.Errorf(
			"offset = (%v, %v), want (72.0, 36.0)",
			ct.offsetX,
			ct.offsetY,
		)
	}
}

func TestCoordTransformer_EMUToPoint(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)

	// 914400 EMU = 1 inch = 72 points
	x, y := ct.EMUToPoint(914400, 914400)
	if !floatEqualsWithTolerance(
		x,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			72.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPoint(914400, 914400) = (%v, %v), want (72.0, 72.0)",
			x,
			y,
		)
	}

	// Test with zero
	x, y = ct.EMUToPoint(0, 0)
	if x != 0 || y != 0 {
		t.Errorf(
			"EMUToPoint(0, 0) = (%v, %v), want (0, 0)",
			x,
			y,
		)
	}
}

func TestCoordTransformer_EMUToPDFCoord(
	t *testing.T,
) {
	ct := NewCoordTransformer(
		792.0,
	) // Letter height

	// Test origin (0, 0) in OOXML should be (0, 792) in PDF
	x, y := ct.EMUToPDFCoord(0, 0)
	if x != 0 ||
		!floatEqualsWithTolerance(
			y,
			792.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPDFCoord(0, 0) = (%v, %v), want (0, 792.0)",
			x,
			y,
		)
	}

	// Test 1 inch from top in OOXML should be 10 inches from bottom in PDF (for 11" page)
	x, y = ct.EMUToPDFCoord(
		0,
		914400,
	) // 1 inch = 72 points from top
	if x != 0 ||
		!floatEqualsWithTolerance(
			y,
			720.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPDFCoord(0, 914400) = (%v, %v), want (0, 720.0)",
			x,
			y,
		)
	}

	// Test point at 1", 1" from top-left in OOXML
	x, y = ct.EMUToPDFCoord(914400, 914400)
	if !floatEqualsWithTolerance(
		x,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			720.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPDFCoord(914400, 914400) = (%v, %v), want (72.0, 720.0)",
			x,
			y,
		)
	}
}

func TestCoordTransformer_PointToPDFCoord(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)

	// Origin in OOXML space
	x, y := ct.PointToPDFCoord(0, 0)
	if x != 0 ||
		!floatEqualsWithTolerance(
			y,
			792.0,
			0.001,
		) {
		t.Errorf(
			"PointToPDFCoord(0, 0) = (%v, %v), want (0, 792.0)",
			x,
			y,
		)
	}

	// Point at (72, 72) in OOXML = (72, 720) in PDF
	x, y = ct.PointToPDFCoord(72, 72)
	if !floatEqualsWithTolerance(
		x,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			720.0,
			0.001,
		) {
		t.Errorf(
			"PointToPDFCoord(72, 72) = (%v, %v), want (72.0, 720.0)",
			x,
			y,
		)
	}
}

func TestCoordTransformer_EMUToPDFSize(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)

	// 914400 EMU = 72 points
	w, h := ct.EMUToPDFSize(914400, 914400)
	if !floatEqualsWithTolerance(
		w,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			h,
			72.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPDFSize(914400, 914400) = (%v, %v), want (72.0, 72.0)",
			w,
			h,
		)
	}
}

func TestCoordTransformer_WithScale(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)
	ct.SetScale(2.0, 2.0)

	// With 2x scale, 1 inch (914400 EMU) should become 144 points
	x, y := ct.EMUToPoint(914400, 914400)
	if !floatEqualsWithTolerance(
		x,
		144.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			144.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPoint with 2x scale = (%v, %v), want (144.0, 144.0)",
			x,
			y,
		)
	}
}

func TestCoordTransformer_WithOffset(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)
	ct.SetOffset(72.0, 72.0)

	// With 72pt offset, origin should shift
	x, y := ct.EMUToPoint(0, 0)
	if !floatEqualsWithTolerance(
		x,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			72.0,
			0.001,
		) {
		t.Errorf(
			"EMUToPoint with offset = (%v, %v), want (72.0, 72.0)",
			x,
			y,
		)
	}
}

func TestOOXMLTransform_Basic(t *testing.T) {
	xfrm := NewOOXMLTransform(
		914400,
		914400,
		914400*2,
		914400*3,
	)

	offX, offY := xfrm.OffsetPoints()
	if !floatEqualsWithTolerance(
		offX,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			offY,
			72.0,
			0.001,
		) {
		t.Errorf(
			"OffsetPoints() = (%v, %v), want (72.0, 72.0)",
			offX,
			offY,
		)
	}

	extX, extY := xfrm.ExtentPoints()
	if !floatEqualsWithTolerance(
		extX,
		144.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			extY,
			216.0,
			0.001,
		) {
		t.Errorf(
			"ExtentPoints() = (%v, %v), want (144.0, 216.0)",
			extX,
			extY,
		)
	}

	centerX, centerY := xfrm.CenterPoints()
	if !floatEqualsWithTolerance(
		centerX,
		72.0+72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			centerY,
			72.0+108.0,
			0.001,
		) {
		t.Errorf(
			"CenterPoints() = (%v, %v), want (144.0, 180.0)",
			centerX,
			centerY,
		)
	}
}

func TestOOXMLTransform_Rotation(t *testing.T) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)

	// Test 90 degrees (5400000 units)
	xfrm = xfrm.WithRotationDegrees(90)
	if !floatEqualsWithTolerance(
		xfrm.RotationDegrees(),
		90.0,
		0.001,
	) {
		t.Errorf(
			"RotationDegrees() = %v, want 90.0",
			xfrm.RotationDegrees(),
		)
	}

	// Test in radians
	if !floatEqualsWithTolerance(
		xfrm.RotationRadians(),
		math.Pi/2,
		0.0001,
	) {
		t.Errorf(
			"RotationRadians() = %v, want %v",
			xfrm.RotationRadians(),
			math.Pi/2,
		)
	}

	// Test HasRotation
	if !xfrm.HasRotation() {
		t.Error(
			"HasRotation() should be true for 90 degree rotation",
		)
	}

	// Test zero rotation
	xfrm = xfrm.WithRotation(0)
	if xfrm.HasRotation() {
		t.Error(
			"HasRotation() should be false for zero rotation",
		)
	}
}

func TestOOXMLTransform_Flips(t *testing.T) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)

	// Test flip states
	if xfrm.FlipH || xfrm.FlipV {
		t.Error(
			"Initial transform should have no flips",
		)
	}

	xfrm = xfrm.WithFlipH(true)
	if !xfrm.FlipH || xfrm.FlipV {
		t.Error(
			"FlipH should be true, FlipV should be false",
		)
	}

	xfrm = xfrm.WithFlipV(true)
	if !xfrm.FlipH || !xfrm.FlipV {
		t.Error("Both flips should be true")
	}

	if !xfrm.HasFlip() {
		t.Error(
			"HasFlip() should be true when flips are set",
		)
	}
}

func TestPDFMatrix_Identity(t *testing.T) {
	m := IdentityMatrix()

	if !m.IsIdentity() {
		t.Error(
			"IdentityMatrix() should be identity",
		)
	}

	// Transform a point - should return same point
	x, y := m.TransformPoint(10, 20)
	if x != 10 || y != 20 {
		t.Errorf(
			"Identity transform of (10, 20) = (%v, %v), want (10, 20)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_Translation(t *testing.T) {
	m := TranslationMatrix(100, 50)

	x, y := m.TransformPoint(10, 20)
	if x != 110 || y != 70 {
		t.Errorf(
			"Translation(100, 50) of (10, 20) = (%v, %v), want (110, 70)",
			x,
			y,
		)
	}

	// Translation should not affect vectors
	vx, vy := m.TransformVector(10, 20)
	if vx != 10 || vy != 20 {
		t.Errorf(
			"Translation vector transform = (%v, %v), want (10, 20)",
			vx,
			vy,
		)
	}
}

func TestPDFMatrix_Scale(t *testing.T) {
	m := ScaleMatrix(2, 3)

	x, y := m.TransformPoint(10, 20)
	if x != 20 || y != 60 {
		t.Errorf(
			"Scale(2, 3) of (10, 20) = (%v, %v), want (20, 60)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_Rotation(t *testing.T) {
	// 90 degree rotation counter-clockwise
	m := RotationMatrix(math.Pi / 2)

	x, y := m.TransformPoint(1, 0)
	if !floatEqualsWithTolerance(x, 0, 0.0001) ||
		!floatEqualsWithTolerance(y, 1, 0.0001) {
		t.Errorf(
			"90 degree rotation of (1, 0) = (%v, %v), want (0, 1)",
			x,
			y,
		)
	}

	// 180 degree rotation
	m = RotationMatrix(math.Pi)
	x, y = m.TransformPoint(1, 0)
	if !floatEqualsWithTolerance(x, -1, 0.0001) ||
		!floatEqualsWithTolerance(y, 0, 0.0001) {
		t.Errorf(
			"180 degree rotation of (1, 0) = (%v, %v), want (-1, 0)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_RotateDegrees(t *testing.T) {
	m := IdentityMatrix().RotateDegrees(90)

	x, y := m.TransformPoint(1, 0)
	if !floatEqualsWithTolerance(x, 0, 0.0001) ||
		!floatEqualsWithTolerance(y, 1, 0.0001) {
		t.Errorf(
			"90 degree rotation of (1, 0) = (%v, %v), want (0, 1)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_FlipHorizontal(t *testing.T) {
	m := IdentityMatrix().FlipHorizontal()

	x, y := m.TransformPoint(10, 20)
	if x != -10 || y != 20 {
		t.Errorf(
			"FlipHorizontal of (10, 20) = (%v, %v), want (-10, 20)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_FlipVertical(t *testing.T) {
	m := IdentityMatrix().FlipVertical()

	x, y := m.TransformPoint(10, 20)
	if x != 10 || y != -20 {
		t.Errorf(
			"FlipVertical of (10, 20) = (%v, %v), want (10, -20)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_Multiply(t *testing.T) {
	// Translate then scale
	t1 := TranslationMatrix(10, 10)
	s := ScaleMatrix(2, 2)
	combined := s.Multiply(t1)

	// Point (5, 5) -> translate to (15, 15) -> scale to (30, 30)
	x, y := combined.TransformPoint(5, 5)
	if x != 30 || y != 30 {
		t.Errorf(
			"Translate then scale = (%v, %v), want (30, 30)",
			x,
			y,
		)
	}
}

func TestPDFMatrix_Inverse(t *testing.T) {
	m := TranslationMatrix(100, 50).Scale(2, 3)
	inv := m.Inverse()

	// m * inv should be identity
	result := m.Multiply(inv)
	if !result.IsIdentity() {
		t.Errorf(
			"Matrix * Inverse should be identity, got %+v",
			result,
		)
	}

	// inv * m should also be identity
	result = inv.Multiply(m)
	if !result.IsIdentity() {
		t.Errorf(
			"Inverse * Matrix should be identity, got %+v",
			result,
		)
	}
}

func TestPDFMatrix_Determinant(t *testing.T) {
	// Identity has determinant 1
	m := IdentityMatrix()
	if m.Determinant() != 1 {
		t.Errorf(
			"Identity determinant = %v, want 1",
			m.Determinant(),
		)
	}

	// Scale(2, 3) has determinant 6
	m = ScaleMatrix(2, 3)
	if m.Determinant() != 6 {
		t.Errorf(
			"Scale(2, 3) determinant = %v, want 6",
			m.Determinant(),
		)
	}

	// Rotation should have determinant 1
	m = RotationMatrix(math.Pi / 4)
	if !floatEqualsWithTolerance(
		m.Determinant(),
		1,
		0.0001,
	) {
		t.Errorf(
			"Rotation determinant = %v, want 1",
			m.Determinant(),
		)
	}
}

func TestPDFMatrix_Array(t *testing.T) {
	m := PDFMatrix{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
		F: 6,
	}
	arr := m.Array()

	expected := [6]float64{1, 2, 3, 4, 5, 6}
	if arr != expected {
		t.Errorf(
			"Array() = %v, want %v",
			arr,
			expected,
		)
	}
}

func TestOOXMLTransformToPDFMatrix_NoTransform(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	)
	m := OOXMLTransformToPDFMatrix(xfrm, 792.0)

	// No transform should return identity
	if !m.IsIdentity() {
		t.Errorf(
			"No transform should return identity, got %+v",
			m,
		)
	}
}

func TestOOXMLTransformToPDFMatrix_Rotation(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	) // 1 inch square at origin
	xfrm = xfrm.WithRotationDegrees(90)

	m := OOXMLTransformToPDFMatrix(xfrm, 792.0)

	// The matrix should not be identity
	if m.IsIdentity() {
		t.Error(
			"90 degree rotation should not produce identity matrix",
		)
	}

	// Check that determinant is preserved (should be 1 for pure rotation)
	det := m.Determinant()
	if !floatEqualsWithTolerance(det, 1, 0.0001) {
		t.Errorf(
			"Rotation matrix determinant = %v, want 1",
			det,
		)
	}
}

func TestOOXMLTransformToPDFMatrix_FlipH(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)
	xfrm = xfrm.WithFlipH(true)

	m := OOXMLTransformToPDFMatrix(xfrm, 792.0)

	// Determinant should be -1 for horizontal flip
	det := m.Determinant()
	if !floatEqualsWithTolerance(
		det,
		-1,
		0.0001,
	) {
		t.Errorf(
			"FlipH matrix determinant = %v, want -1",
			det,
		)
	}
}

func TestOOXMLTransformToPDFMatrix_FlipV(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)
	xfrm = xfrm.WithFlipV(true)

	m := OOXMLTransformToPDFMatrix(xfrm, 792.0)

	// Determinant should be -1 for vertical flip
	det := m.Determinant()
	if !floatEqualsWithTolerance(
		det,
		-1,
		0.0001,
	) {
		t.Errorf(
			"FlipV matrix determinant = %v, want -1",
			det,
		)
	}
}

func TestOOXMLTransformToPDFMatrix_BothFlips(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)
	xfrm = xfrm.WithFlipH(true).WithFlipV(true)

	m := OOXMLTransformToPDFMatrix(xfrm, 792.0)

	// Two flips should give determinant 1 (equivalent to 180 rotation)
	det := m.Determinant()
	if !floatEqualsWithTolerance(det, 1, 0.0001) {
		t.Errorf(
			"Both flips matrix determinant = %v, want 1",
			det,
		)
	}
}

func TestCoordTransformer_OOXMLToPDFBounds(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)

	// 1 inch square at 1 inch from top-left
	xfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	)

	x, y, w, h := ct.OOXMLToPDFBounds(xfrm)

	// X should be 72 points from left
	if !floatEqualsWithTolerance(x, 72.0, 0.001) {
		t.Errorf("x = %v, want 72.0", x)
	}

	// Width and height should be 72 points
	if !floatEqualsWithTolerance(
		w,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			h,
			72.0,
			0.001,
		) {
		t.Errorf(
			"size = (%v, %v), want (72.0, 72.0)",
			w,
			h,
		)
	}

	// Y in PDF: page height (792) - OOXML y (72) - height (72) = 648
	if !floatEqualsWithTolerance(
		y,
		648.0,
		0.001,
	) {
		t.Errorf("y = %v, want 648.0", y)
	}
}

func TestTransformedOOXMLBounds_NoRotation(
	t *testing.T,
) {
	xfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	)

	x, _, w, h := TransformedOOXMLBounds(
		xfrm,
		792.0,
	)

	if !floatEqualsWithTolerance(x, 72.0, 0.001) {
		t.Errorf("x = %v, want 72.0", x)
	}
	if !floatEqualsWithTolerance(
		w,
		72.0,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			h,
			72.0,
			0.001,
		) {
		t.Errorf(
			"size = (%v, %v), want (72.0, 72.0)",
			w,
			h,
		)
	}
}

func TestTransformedOOXMLBounds_45DegreeRotation(
	t *testing.T,
) {
	// 1 inch square at origin, rotated 45 degrees
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	)
	xfrm = xfrm.WithRotationDegrees(45)

	x, _, w, h := TransformedOOXMLBounds(
		xfrm,
		792.0,
	)

	// A 72x72 square rotated 45 degrees has a bounding box of ~102x102
	expectedSize := 72.0 * math.Sqrt(2)
	if !floatEqualsWithTolerance(
		w,
		expectedSize,
		0.5,
	) ||
		!floatEqualsWithTolerance(
			h,
			expectedSize,
			0.5,
		) {
		t.Errorf(
			"45 degree rotated bounds size = (%v, %v), want (%v, %v)",
			w,
			h,
			expectedSize,
			expectedSize,
		)
	}

	// The bounding box should be centered around the original shape's center
	// Bounds should be non-negative for a shape at origin
	if x > 0 {
		t.Errorf(
			"x = %v, should be negative or zero for centered rotation",
			x,
		)
	}
}

func TestNestedTransformer_Basic(t *testing.T) {
	nt := NewNestedTransformer(792.0)

	if nt.Depth() != 0 {
		t.Errorf(
			"Initial depth = %v, want 0",
			nt.Depth(),
		)
	}

	if !nt.CurrentMatrix().IsIdentity() {
		t.Error(
			"Initial matrix should be identity",
		)
	}
}

func TestNestedTransformer_PushPop(t *testing.T) {
	nt := NewNestedTransformer(792.0)

	// Push a transform with rotation
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	).WithRotationDegrees(90)
	nt.PushTransform(xfrm)

	if nt.Depth() != 1 {
		t.Errorf(
			"After push depth = %v, want 1",
			nt.Depth(),
		)
	}

	if nt.CurrentMatrix().IsIdentity() {
		t.Error(
			"After push matrix should not be identity",
		)
	}

	// Pop the transform
	nt.PopTransform()

	if nt.Depth() != 0 {
		t.Errorf(
			"After pop depth = %v, want 0",
			nt.Depth(),
		)
	}

	if !nt.CurrentMatrix().IsIdentity() {
		t.Error(
			"After pop matrix should be identity",
		)
	}
}

func TestNestedTransformer_MultipleLevels(
	t *testing.T,
) {
	nt := NewNestedTransformer(792.0)

	// Push multiple transforms
	xfrm1 := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	).WithRotationDegrees(45)
	xfrm2 := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	).WithRotationDegrees(45)

	nt.PushTransform(xfrm1)
	nt.PushTransform(xfrm2)

	if nt.Depth() != 2 {
		t.Errorf(
			"After two pushes depth = %v, want 2",
			nt.Depth(),
		)
	}

	// Pop one level
	nt.PopTransform()
	if nt.Depth() != 1 {
		t.Errorf(
			"After one pop depth = %v, want 1",
			nt.Depth(),
		)
	}

	// Reset should clear all
	nt.Reset()
	if nt.Depth() != 0 {
		t.Errorf(
			"After reset depth = %v, want 0",
			nt.Depth(),
		)
	}
}

func TestNestedTransformer_TransformPoint(
	t *testing.T,
) {
	nt := NewNestedTransformer(792.0)

	// With identity matrix, point should be unchanged
	x, y := nt.TransformPoint(100, 100)
	if x != 100 || y != 100 {
		t.Errorf(
			"Identity TransformPoint(100, 100) = (%v, %v), want (100, 100)",
			x,
			y,
		)
	}
}

func TestGroupTransform_Basic(t *testing.T) {
	// Create a group transform where child space is 1000x1000 EMU
	// and it maps to a 500x500 EMU area at offset (100, 100)
	groupXfrm := NewOOXMLTransform(
		100000,
		100000,
		500000,
		500000,
	) // offset and extent in EMU
	gt := NewGroupTransform(
		0,
		0,
		1000000,
		1000000,
		groupXfrm,
	) // child origin at 0,0, child extent 1000000 EMU

	matrix := gt.ChildToGroupMatrix()

	// A point at (500000, 500000) in child space (center of child)
	// should map to the center of the group's extent
	// After transformation: 500000 * 0.5 + 100000 = 350000 EMU = ~27.56 points
	// We need to check the scaling is correct
	if matrix.IsIdentity() {
		t.Error(
			"GroupTransform matrix should not be identity when scaling is applied",
		)
	}
}

func TestGroupTransform_ChildToGroupMatrix(
	t *testing.T,
) {
	// Group has 2:1 scale ratio (child is twice as big as group extent)
	groupXfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	) // 1" offset, 1" extent
	gt := NewGroupTransform(
		0,
		0,
		914400*2,
		914400*2,
		groupXfrm,
	) // child extent is 2"

	matrix := gt.ChildToGroupMatrix()

	// The matrix should scale by 0.5 (group extent / child extent)
	// A point at (914400, 914400) in child (1" from origin)
	// should map to 0.5" in the group's space, then offset by 1"
	// Result: 72 * 0.5 + 72 = 108 points

	testX := float64(
		914400,
	) * PointsPerEMU // 72 points
	testY := float64(914400) * PointsPerEMU

	resultX, resultY := matrix.TransformPoint(
		testX,
		testY,
	)

	expectedX := 72.0*0.5 + 72.0 // 108 points
	expectedY := 72.0*0.5 + 72.0

	if !floatEqualsWithTolerance(
		resultX,
		expectedX,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			resultY,
			expectedY,
			0.001,
		) {
		t.Errorf(
			"ChildToGroupMatrix transform = (%v, %v), want (%v, %v)",
			resultX,
			resultY,
			expectedX,
			expectedY,
		)
	}
}

func TestNestedTransformer_PushGroupTransform(
	t *testing.T,
) {
	nt := NewNestedTransformer(792.0)

	// Create a simple group transform
	groupXfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	)
	gt := NewGroupTransform(
		0,
		0,
		914400,
		914400,
		groupXfrm,
	)

	nt.PushGroupTransform(gt)

	if nt.Depth() != 1 {
		t.Errorf(
			"After PushGroupTransform depth = %v, want 1",
			nt.Depth(),
		)
	}

	// Pop should work
	nt.PopTransform()
	if nt.Depth() != 0 {
		t.Errorf(
			"After PopTransform depth = %v, want 0",
			nt.Depth(),
		)
	}
}

func TestCoordinateConversion_RoundTrip(
	t *testing.T,
) {
	ct := NewCoordTransformer(792.0)

	// Test that various EMU values convert correctly
	testCases := []struct {
		emuX, emuY int64
	}{
		{0, 0},
		{914400, 0},
		{0, 914400},
		{914400, 914400},
		{914400 * 4, 914400 * 5}, // 4" x 5"
	}

	for _, tc := range testCases {
		// Convert to PDF coordinates
		pdfX, pdfY := ct.EMUToPDFCoord(
			tc.emuX,
			tc.emuY,
		)

		// The conversion should be consistent
		expectedPdfX := float64(
			tc.emuX,
		) * PointsPerEMU
		expectedPdfY := ct.pageHeight - float64(
			tc.emuY,
		)*PointsPerEMU

		if !floatEqualsWithTolerance(
			pdfX,
			expectedPdfX,
			0.0001,
		) ||
			!floatEqualsWithTolerance(
				pdfY,
				expectedPdfY,
				0.0001,
			) {
			t.Errorf(
				"EMUToPDFCoord(%d, %d) = (%v, %v), want (%v, %v)",
				tc.emuX,
				tc.emuY,
				pdfX,
				pdfY,
				expectedPdfX,
				expectedPdfY,
			)
		}
	}
}

func TestOOXMLAngleUnits(t *testing.T) {
	// Verify OOXML angle unit conversion
	// OOXML uses 60000ths of a degree

	// 90 degrees = 5400000 units
	xfrm := NewOOXMLTransform(
		0,
		0,
		100,
		100,
	).WithRotation(5400000)
	if !floatEqualsWithTolerance(
		xfrm.RotationDegrees(),
		90.0,
		0.0001,
	) {
		t.Errorf(
			"5400000 units = %v degrees, want 90",
			xfrm.RotationDegrees(),
		)
	}

	// 180 degrees = 10800000 units
	xfrm = xfrm.WithRotation(10800000)
	if !floatEqualsWithTolerance(
		xfrm.RotationDegrees(),
		180.0,
		0.0001,
	) {
		t.Errorf(
			"10800000 units = %v degrees, want 180",
			xfrm.RotationDegrees(),
		)
	}

	// 45 degrees = 2700000 units
	xfrm = xfrm.WithRotationDegrees(45)
	if xfrm.Rotation != 2700000 {
		t.Errorf(
			"45 degrees = %v units, want 2700000",
			xfrm.Rotation,
		)
	}
}

func BenchmarkCoordTransformer_EMUToPDFCoord(
	b *testing.B,
) {
	ct := NewCoordTransformer(792.0)

	for range b.N {
		ct.EMUToPDFCoord(914400, 914400)
	}
}

func BenchmarkPDFMatrix_Multiply(b *testing.B) {
	m1 := TranslationMatrix(100, 100)
	m2 := RotationMatrix(math.Pi / 4)

	for range b.N {
		_ = m1.Multiply(m2)
	}
}

func BenchmarkOOXMLTransformToPDFMatrix(
	b *testing.B,
) {
	xfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	).
		WithRotationDegrees(45).
		WithFlipH(true)

	for range b.N {
		_ = OOXMLTransformToPDFMatrix(xfrm, 792.0)
	}
}

// TestPDFMatrixEdgeCases tests edge cases in PDF matrix operations.
func TestPDFMatrixEdgeCases(t *testing.T) {
	t.Run(
		"singular matrix inverse",
		func(t *testing.T) {
			// A singular matrix (determinant = 0) should return identity when inverted
			singular := PDFMatrix{
				A: 0,
				B: 0,
				C: 0,
				D: 0,
				E: 100,
				F: 100,
			}
			inv := singular.Inverse()

			if !inv.IsIdentity() {
				t.Errorf(
					"Inverse of singular matrix should be identity, got %+v",
					inv,
				)
			}
		},
	)

	t.Run(
		"near-singular matrix inverse",
		func(t *testing.T) {
			// A matrix with very small determinant
			nearSingular := PDFMatrix{
				A: 1e-11,
				B: 0,
				C: 0,
				D: 1e-11,
				E: 0,
				F: 0,
			}
			inv := nearSingular.Inverse()

			// Should return identity since det < 1e-10
			if !inv.IsIdentity() {
				t.Errorf(
					"Inverse of near-singular matrix should be identity, got %+v",
					inv,
				)
			}
		},
	)

	t.Run(
		"double inverse returns original",
		func(t *testing.T) {
			original := TranslationMatrix(
				100,
				50,
			).Scale(2, 3).
				Rotate(math.Pi / 4)
			inv := original.Inverse()
			doubleInv := inv.Inverse()

			// Double inverse should approximately equal original
			if !floatEqualsWithTolerance(
				doubleInv.A,
				original.A,
				1e-9,
			) ||
				!floatEqualsWithTolerance(
					doubleInv.B,
					original.B,
					1e-9,
				) ||
				!floatEqualsWithTolerance(
					doubleInv.C,
					original.C,
					1e-9,
				) ||
				!floatEqualsWithTolerance(
					doubleInv.D,
					original.D,
					1e-9,
				) ||
				!floatEqualsWithTolerance(
					doubleInv.E,
					original.E,
					1e-9,
				) ||
				!floatEqualsWithTolerance(
					doubleInv.F,
					original.F,
					1e-9,
				) {
				t.Errorf(
					"Double inverse should equal original:\noriginal: %+v\ndouble inv: %+v",
					original,
					doubleInv,
				)
			}
		},
	)

	t.Run(
		"matrix multiplication associativity",
		func(t *testing.T) {
			m1 := TranslationMatrix(10, 20)
			m2 := ScaleMatrix(2, 3)
			m3 := RotationMatrix(math.Pi / 6)

			// (m1 * m2) * m3 should equal m1 * (m2 * m3)
			left := m1.Multiply(m2).Multiply(m3)
			right := m1.Multiply(m2.Multiply(m3))

			if !floatEqualsWithTolerance(
				left.A,
				right.A,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					left.B,
					right.B,
					1e-10,
				) ||
				!floatEqualsWithTolerance(
					left.C,
					right.C,
					1e-10,
				) ||
				!floatEqualsWithTolerance(
					left.D,
					right.D,
					1e-10,
				) ||
				!floatEqualsWithTolerance(
					left.E,
					right.E,
					1e-10,
				) ||
				!floatEqualsWithTolerance(
					left.F,
					right.F,
					1e-10,
				) {
				t.Errorf(
					"Matrix multiplication not associative:\n(m1*m2)*m3: %+v\nm1*(m2*m3): %+v",
					left,
					right,
				)
			}
		},
	)

	t.Run(
		"rotation by multiples of 90 degrees",
		func(t *testing.T) {
			// 90 degree rotation: (1,0) -> (0,1)
			rot90 := RotationMatrix(math.Pi / 2)
			x, y := rot90.TransformPoint(1, 0)
			if !floatEqualsWithTolerance(
				x,
				0,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					y,
					1,
					1e-10,
				) {
				t.Errorf(
					"90 degree rotation: (1,0) -> (%v,%v), want (0,1)",
					x,
					y,
				)
			}

			// 180 degree rotation: (1,0) -> (-1,0)
			rot180 := RotationMatrix(math.Pi)
			x, y = rot180.TransformPoint(1, 0)
			if !floatEqualsWithTolerance(
				x,
				-1,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					y,
					0,
					1e-10,
				) {
				t.Errorf(
					"180 degree rotation: (1,0) -> (%v,%v), want (-1,0)",
					x,
					y,
				)
			}

			// 270 degree rotation: (1,0) -> (0,-1)
			rot270 := RotationMatrix(
				3 * math.Pi / 2,
			)
			x, y = rot270.TransformPoint(1, 0)
			if !floatEqualsWithTolerance(
				x,
				0,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					y,
					-1,
					1e-10,
				) {
				t.Errorf(
					"270 degree rotation: (1,0) -> (%v,%v), want (0,-1)",
					x,
					y,
				)
			}

			// 360 degree rotation: (1,0) -> (1,0)
			rot360 := RotationMatrix(2 * math.Pi)
			x, y = rot360.TransformPoint(1, 0)
			if !floatEqualsWithTolerance(
				x,
				1,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					y,
					0,
					1e-10,
				) {
				t.Errorf(
					"360 degree rotation: (1,0) -> (%v,%v), want (1,0)",
					x,
					y,
				)
			}
		},
	)

	t.Run(
		"negative rotation",
		func(t *testing.T) {
			// -90 degree rotation: (1,0) -> (0,-1)
			rotNeg90 := RotationMatrix(
				-math.Pi / 2,
			)
			x, y := rotNeg90.TransformPoint(1, 0)
			if !floatEqualsWithTolerance(
				x,
				0,
				1e-10,
			) ||
				!floatEqualsWithTolerance(
					y,
					-1,
					1e-10,
				) {
				t.Errorf(
					"-90 degree rotation: (1,0) -> (%v,%v), want (0,-1)",
					x,
					y,
				)
			}
		},
	)
}

// TestCoordTransformerEdgeCases tests edge cases in coordinate transformation.
func TestCoordTransformerEdgeCases(t *testing.T) {
	t.Run("zero page height", func(t *testing.T) {
		ct := NewCoordTransformer(0)

		// Y-flip with zero page height should work (but result in negative Y)
		x, y := ct.EMUToPDFCoord(914400, 914400)
		expectedX := 72.0
		expectedY := 0 - 72.0 // pageHeight - y = 0 - 72 = -72

		if !floatEqualsWithTolerance(
			x,
			expectedX,
			0.001,
		) {
			t.Errorf(
				"Zero page height X: got %v, want %v",
				x,
				expectedX,
			)
		}
		if !floatEqualsWithTolerance(
			y,
			expectedY,
			0.001,
		) {
			t.Errorf(
				"Zero page height Y: got %v, want %v",
				y,
				expectedY,
			)
		}
	})

	t.Run("very large page", func(t *testing.T) {
		// 1000 inches = 72000 points
		largePageHeight := 72000.0
		ct := NewCoordTransformer(largePageHeight)

		x, y := ct.EMUToPDFCoord(0, 0)
		if x != 0 || y != largePageHeight {
			t.Errorf(
				"Large page origin: got (%v,%v), want (0,%v)",
				x,
				y,
				largePageHeight,
			)
		}
	})

	t.Run(
		"negative EMU coordinates",
		func(t *testing.T) {
			ct := NewCoordTransformer(792.0)

			// Negative X (to the left of origin)
			x, y := ct.EMUToPDFCoord(-914400, 0)
			if !floatEqualsWithTolerance(
				x,
				-72.0,
				0.001,
			) {
				t.Errorf(
					"Negative X: got %v, want -72",
					x,
				)
			}
			if !floatEqualsWithTolerance(
				y,
				792.0,
				0.001,
			) {
				t.Errorf(
					"Negative X, Y should still be page height: got %v, want 792",
					y,
				)
			}

			// Negative Y (above origin in OOXML, below page height in PDF)
			x, y = ct.EMUToPDFCoord(0, -914400)
			if !floatEqualsWithTolerance(
				x,
				0,
				0.001,
			) {
				t.Errorf(
					"Negative Y, X should be 0: got %v",
					x,
				)
			}
			expectedY := 792.0 - (-72.0) // pageHeight - (-72) = 864
			if !floatEqualsWithTolerance(
				y,
				expectedY,
				0.001,
			) {
				t.Errorf(
					"Negative Y: got %v, want %v",
					y,
					expectedY,
				)
			}
		},
	)

	t.Run(
		"combined scale and offset",
		func(t *testing.T) {
			ct := NewCoordTransformer(792.0)
			ct.SetScale(2.0, 0.5)
			ct.SetOffset(50, 25)

			// Test EMUToPoint (no Y-flip)
			x, y := ct.EMUToPoint(914400, 914400)
			expectedX := 72.0*2.0 + 50 // 194
			expectedY := 72.0*0.5 + 25 // 61
			if !floatEqualsWithTolerance(
				x,
				expectedX,
				0.001,
			) ||
				!floatEqualsWithTolerance(
					y,
					expectedY,
					0.001,
				) {
				t.Errorf(
					"Scale+offset: got (%v,%v), want (%v,%v)",
					x,
					y,
					expectedX,
					expectedY,
				)
			}
		},
	)
}

// TestOOXMLTransformEdgeCases tests edge cases in OOXML transform handling.
func TestOOXMLTransformEdgeCases(t *testing.T) {
	t.Run("zero extent", func(t *testing.T) {
		// A shape with zero size
		xfrm := NewOOXMLTransform(
			914400,
			914400,
			0,
			0,
		)

		offX, offY := xfrm.OffsetPoints()
		if !floatEqualsWithTolerance(
			offX,
			72,
			0.001,
		) ||
			!floatEqualsWithTolerance(
				offY,
				72,
				0.001,
			) {
			t.Errorf(
				"Zero extent offset: got (%v,%v), want (72,72)",
				offX,
				offY,
			)
		}

		extCx, extCy := xfrm.ExtentPoints()
		if extCx != 0 || extCy != 0 {
			t.Errorf(
				"Zero extent size: got (%v,%v), want (0,0)",
				extCx,
				extCy,
			)
		}

		// Center should be same as offset for zero extent
		centerX, centerY := xfrm.CenterPoints()
		if !floatEqualsWithTolerance(
			centerX,
			72,
			0.001,
		) ||
			!floatEqualsWithTolerance(
				centerY,
				72,
				0.001,
			) {
			t.Errorf(
				"Zero extent center: got (%v,%v), want (72,72)",
				centerX,
				centerY,
			)
		}
	})

	t.Run("negative offset", func(t *testing.T) {
		xfrm := NewOOXMLTransform(
			-914400,
			-914400,
			914400,
			914400,
		)

		offX, offY := xfrm.OffsetPoints()
		if !floatEqualsWithTolerance(
			offX,
			-72,
			0.001,
		) ||
			!floatEqualsWithTolerance(
				offY,
				-72,
				0.001,
			) {
			t.Errorf(
				"Negative offset: got (%v,%v), want (-72,-72)",
				offX,
				offY,
			)
		}
	})

	t.Run(
		"very large rotation",
		func(t *testing.T) {
			// 720 degrees (2 full rotations)
			xfrm := NewOOXMLTransform(
				0,
				0,
				914400,
				914400,
			).WithRotationDegrees(720)

			if !floatEqualsWithTolerance(
				xfrm.RotationDegrees(),
				720,
				0.001,
			) {
				t.Errorf(
					"Large rotation: got %v degrees, want 720",
					xfrm.RotationDegrees(),
				)
			}

			// RadianS should also work
			expectedRadians := 720.0 * math.Pi / 180.0
			if !floatEqualsWithTolerance(
				xfrm.RotationRadians(),
				expectedRadians,
				0.0001,
			) {
				t.Errorf(
					"Large rotation radians: got %v, want %v",
					xfrm.RotationRadians(),
					expectedRadians,
				)
			}
		},
	)

	t.Run(
		"combined rotation and flip",
		func(t *testing.T) {
			xfrm := NewOOXMLTransform(
				0,
				0,
				914400,
				914400,
			).
				WithRotationDegrees(90).
				WithFlipH(true).
				WithFlipV(true)

			if !xfrm.HasTransform() {
				t.Error(
					"HasTransform should be true",
				)
			}
			if !xfrm.HasRotation() {
				t.Error(
					"HasRotation should be true",
				)
			}
			if !xfrm.HasFlip() {
				t.Error("HasFlip should be true")
			}

			// Generate the matrix and verify it's not identity
			m := OOXMLTransformToPDFMatrix(
				xfrm,
				792.0,
			)
			if m.IsIdentity() {
				t.Error(
					"Combined rotation+flip should not produce identity matrix",
				)
			}
		},
	)
}

// TestNestedTransformerEdgeCases tests edge cases in nested transformation.
func TestNestedTransformerEdgeCases(
	t *testing.T,
) {
	t.Run("deep nesting", func(t *testing.T) {
		nt := NewNestedTransformer(792.0)

		// Push 10 levels of transforms
		for range 10 {
			xfrm := NewOOXMLTransform(
				0,
				0,
				914400,
				914400,
			).WithRotationDegrees(10)
			nt.PushTransform(xfrm)
		}

		if nt.Depth() != 10 {
			t.Errorf(
				"After 10 pushes, depth = %d, want 10",
				nt.Depth(),
			)
		}

		// Pop all levels
		for range 10 {
			nt.PopTransform()
		}

		if nt.Depth() != 0 {
			t.Errorf(
				"After 10 pops, depth = %d, want 0",
				nt.Depth(),
			)
		}

		// Should not go below 0
		nt.PopTransform()
		if nt.Depth() != 0 {
			t.Errorf(
				"After extra pop, depth = %d, want 0",
				nt.Depth(),
			)
		}
	})

	t.Run("reset clears all", func(t *testing.T) {
		nt := NewNestedTransformer(792.0)

		// Push some transforms
		for i := range 5 {
			xfrm := NewOOXMLTransform(
				914400*int64(i),
				914400*int64(i),
				914400,
				914400,
			)
			nt.PushTransform(xfrm)
		}

		nt.Reset()

		if nt.Depth() != 0 {
			t.Errorf(
				"After reset, depth = %d, want 0",
				nt.Depth(),
			)
		}
		if !nt.CurrentMatrix().IsIdentity() {
			t.Error(
				"After reset, matrix should be identity",
			)
		}
	})

	t.Run(
		"empty stack CurrentMatrix",
		func(t *testing.T) {
			// Create a nested transformer with empty stack (edge case)
			nt := &NestedTransformer{
				CoordTransformer: NewCoordTransformer(
					792.0,
				),
				stack: make([]PDFMatrix, 0),
			}

			// Should return identity for empty stack
			m := nt.CurrentMatrix()
			if !m.IsIdentity() {
				t.Errorf(
					"Empty stack CurrentMatrix should be identity, got %+v",
					m,
				)
			}
		},
	)
}

// TestGroupTransformEdgeCases tests edge cases in group transformation.
func TestGroupTransformEdgeCases(t *testing.T) {
	t.Run(
		"zero child extent",
		func(t *testing.T) {
			// Child extent of zero should not cause division by zero
			groupXfrm := NewOOXMLTransform(
				914400,
				914400,
				914400,
				914400,
			)
			gt := NewGroupTransform(
				0,
				0,
				0,
				0,
				groupXfrm,
			)

			matrix := gt.ChildToGroupMatrix()

			// Scale should be 1 (fallback for zero extent)
			// The matrix should still be valid (not NaN or Inf)
			if math.IsNaN(matrix.A) ||
				math.IsInf(matrix.A, 0) {
				t.Errorf(
					"Zero child extent produced invalid matrix: %+v",
					matrix,
				)
			}
		},
	)

	t.Run(
		"equal child and group extent",
		func(t *testing.T) {
			// When child extent equals group extent, scale should be 1:1
			groupXfrm := NewOOXMLTransform(
				914400,
				914400,
				914400,
				914400,
			)
			gt := NewGroupTransform(
				0,
				0,
				914400,
				914400,
				groupXfrm,
			)

			matrix := gt.ChildToGroupMatrix()

			// Scale should be 1
			if !floatEqualsWithTolerance(
				matrix.A,
				1.0,
				0.001,
			) ||
				!floatEqualsWithTolerance(
					matrix.D,
					1.0,
					0.001,
				) {
				t.Errorf(
					"Equal extent scale should be 1, got A=%v D=%v",
					matrix.A,
					matrix.D,
				)
			}
		},
	)

	t.Run(
		"scaled group with rotation",
		func(t *testing.T) {
			// Group with 2:1 scale and 45 degree rotation
			groupXfrm := NewOOXMLTransform(
				0,
				0,
				914400,
				914400,
			).WithRotationDegrees(45)
			gt := NewGroupTransform(
				0,
				0,
				914400*2,
				914400*2,
				groupXfrm,
			)

			nt := NewNestedTransformer(792.0)
			nt.PushGroupTransform(gt)

			if nt.Depth() != 1 {
				t.Errorf(
					"After PushGroupTransform, depth = %d, want 1",
					nt.Depth(),
				)
			}

			// The matrix should not be identity
			if nt.CurrentMatrix().IsIdentity() {
				t.Error(
					"Scaled+rotated group should not produce identity matrix",
				)
			}
		},
	)
}

// TestCompleteOOXMLToPDFPipeline tests the complete conversion pipeline.
func TestCompleteOOXMLToPDFPipeline(
	t *testing.T,
) {
	t.Run(
		"shape at known position",
		func(t *testing.T) {
			// A 1" x 1" shape at (1", 1") from top-left on a Letter page (8.5" x 11")
			pageHeight := 792.0 // 11 inches

			xfrm := NewOOXMLTransform(
				914400, // 1 inch X offset
				914400, // 1 inch Y offset from top
				914400, // 1 inch width
				914400, // 1 inch height
			)

			ct := NewCoordTransformer(pageHeight)
			x, y, w, h := ct.OOXMLToPDFBounds(
				xfrm,
			)

			// Expected PDF coordinates:
			// X = 72 (1 inch from left)
			// Y = 792 - 72 - 72 = 648 (bottom of shape in PDF)
			// W = 72
			// H = 72
			if !floatEqualsWithTolerance(
				x,
				72,
				0.001,
			) {
				t.Errorf("PDF X = %v, want 72", x)
			}
			if !floatEqualsWithTolerance(
				y,
				648,
				0.001,
			) {
				t.Errorf(
					"PDF Y = %v, want 648",
					y,
				)
			}
			if !floatEqualsWithTolerance(
				w,
				72,
				0.001,
			) {
				t.Errorf("PDF W = %v, want 72", w)
			}
			if !floatEqualsWithTolerance(
				h,
				72,
				0.001,
			) {
				t.Errorf("PDF H = %v, want 72", h)
			}
		},
	)

	t.Run(
		"rotated shape bounds",
		func(t *testing.T) {
			// A 1" x 1" shape rotated 45 degrees
			pageHeight := 792.0

			xfrm := NewOOXMLTransform(
				100*12700,
				100*12700,
				914400,
				914400,
			).
				WithRotationDegrees(45)

			x, y, w, h := TransformedOOXMLBounds(
				xfrm,
				pageHeight,
			)

			// For a 45-degree rotated square, the bounding box diagonal = original side * sqrt(2)
			expectedDiagonal := 72.0 * math.Sqrt(
				2,
			)

			if !floatEqualsWithTolerance(
				w,
				expectedDiagonal,
				0.5,
			) {
				t.Errorf(
					"Rotated bounds width = %v, want ~%v",
					w,
					expectedDiagonal,
				)
			}
			if !floatEqualsWithTolerance(
				h,
				expectedDiagonal,
				0.5,
			) {
				t.Errorf(
					"Rotated bounds height = %v, want ~%v",
					h,
					expectedDiagonal,
				)
			}

			// Position should be adjusted for the expanded bounding box
			_ = x
			_ = y // Just verify they're calculated without error
		},
	)

	t.Run(
		"nested group shapes",
		func(t *testing.T) {
			// Simulate a shape inside a group
			pageHeight := 792.0
			nt := NewNestedTransformer(pageHeight)

			// Group at (1", 1") with 2" x 2" size, child space is 4" x 4"
			groupXfrm := NewOOXMLTransform(
				914400,
				914400,
				914400*2,
				914400*2,
			)
			gt := NewGroupTransform(
				0,
				0,
				914400*4,
				914400*4,
				groupXfrm,
			)

			nt.PushGroupTransform(gt)

			// A point at (2", 2") in child space should map to:
			// - Scale by 0.5 (group extent / child extent = 2/4 = 0.5)
			// - Add group offset (1", 1")
			// Result: (1" + 2"*0.5, 1" + 2"*0.5) = (2", 2") in points = (144, 144)
			childPtX := 144.0 // 2 inches
			childPtY := 144.0

			m := nt.CurrentMatrix()
			resultX, resultY := m.TransformPoint(
				childPtX,
				childPtY,
			)

			// Expected: 144 * 0.5 + 72 = 144
			expectedX := 144.0
			expectedY := 144.0

			if !floatEqualsWithTolerance(
				resultX,
				expectedX,
				0.5,
			) {
				t.Errorf(
					"Nested X = %v, want %v",
					resultX,
					expectedX,
				)
			}
			if !floatEqualsWithTolerance(
				resultY,
				expectedY,
				0.5,
			) {
				t.Errorf(
					"Nested Y = %v, want %v",
					resultY,
					expectedY,
				)
			}
		},
	)
}

// TestDrawingMLXfrmParsing tests parsing of real DrawingML xfrm values.
func TestDrawingMLXfrmParsing(t *testing.T) {
	// These values are from actual OOXML documents
	testCases := []struct {
		name           string
		offX, offY     int64
		extCx, extCy   int64
		rotation       int32
		flipH, flipV   bool
		expectedPtX    float64
		expectedPtY    float64 // offset in points
		expectedWidth  float64
		expectedHeight float64
	}{
		{
			name:           "PowerPoint shape at origin",
			offX:           0,
			offY:           0,
			extCx:          914400,
			extCy:          457200,
			rotation:       0,
			expectedPtX:    0,
			expectedPtY:    0,
			expectedWidth:  72,
			expectedHeight: 36,
		},
		{
			name:           "Word inline image",
			offX:           914400,
			offY:           1828800,
			extCx:          3657600,
			extCy:          2743200,
			rotation:       0,
			expectedPtX:    72,
			expectedPtY:    144,
			expectedWidth:  288,
			expectedHeight: 216,
		},
		{
			name:           "Rotated text box",
			offX:           457200,
			offY:           457200,
			extCx:          1828800,
			extCy:          914400,
			rotation:       5400000, // 90 degrees
			flipH:          false,
			flipV:          false,
			expectedPtX:    36,
			expectedPtY:    36,
			expectedWidth:  144,
			expectedHeight: 72,
		},
		{
			name:           "Flipped shape",
			offX:           0,
			offY:           0,
			extCx:          914400,
			extCy:          914400,
			flipH:          true,
			flipV:          false,
			expectedPtX:    0,
			expectedPtY:    0,
			expectedWidth:  72,
			expectedHeight: 72,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xfrm := NewOOXMLTransform(
				tc.offX,
				tc.offY,
				tc.extCx,
				tc.extCy,
			).
				WithRotation(tc.rotation).
				WithFlipH(tc.flipH).
				WithFlipV(tc.flipV)

			offX, offY := xfrm.OffsetPoints()
			extCx, extCy := xfrm.ExtentPoints()

			if !floatEqualsWithTolerance(
				offX,
				tc.expectedPtX,
				0.001,
			) {
				t.Errorf(
					"Offset X = %v, want %v",
					offX,
					tc.expectedPtX,
				)
			}
			if !floatEqualsWithTolerance(
				offY,
				tc.expectedPtY,
				0.001,
			) {
				t.Errorf(
					"Offset Y = %v, want %v",
					offY,
					tc.expectedPtY,
				)
			}
			if !floatEqualsWithTolerance(
				extCx,
				tc.expectedWidth,
				0.001,
			) {
				t.Errorf(
					"Width = %v, want %v",
					extCx,
					tc.expectedWidth,
				)
			}
			if !floatEqualsWithTolerance(
				extCy,
				tc.expectedHeight,
				0.001,
			) {
				t.Errorf(
					"Height = %v, want %v",
					extCy,
					tc.expectedHeight,
				)
			}
		})
	}
}

// TestOriginTransformation tests the origin transformation (top-left to bottom-left).
func TestOriginTransformation(t *testing.T) {
	t.Run(
		"OOXML origin to PDF origin",
		func(t *testing.T) {
			pageHeight := 792.0 // Letter height

			// OOXML origin (0,0) is top-left
			// PDF origin (0,0) is bottom-left
			ct := NewCoordTransformer(pageHeight)

			// Point at OOXML origin should be at (0, pageHeight) in PDF
			x, y := ct.PointToPDFCoord(0, 0)
			if x != 0 ||
				!floatEqualsWithTolerance(
					y,
					pageHeight,
					0.001,
				) {
				t.Errorf(
					"OOXML (0,0) -> PDF (%v,%v), want (0,%v)",
					x,
					y,
					pageHeight,
				)
			}

			// Point at OOXML (0, pageHeight) should be at PDF (0, 0)
			x, y = ct.PointToPDFCoord(
				0,
				pageHeight,
			)
			if x != 0 ||
				!floatEqualsWithTolerance(
					y,
					0,
					0.001,
				) {
				t.Errorf(
					"OOXML (0,%v) -> PDF (%v,%v), want (0,0)",
					pageHeight,
					x,
					y,
				)
			}

			// Point at center of page
			centerY := pageHeight / 2
			x, y = ct.PointToPDFCoord(0, centerY)
			if x != 0 ||
				!floatEqualsWithTolerance(
					y,
					centerY,
					0.001,
				) {
				t.Errorf(
					"OOXML (0,%v) -> PDF (%v,%v), want (0,%v)",
					centerY,
					x,
					y,
					centerY,
				)
			}
		},
	)

	t.Run(
		"various page sizes",
		func(t *testing.T) {
			pageSizes := []float64{
				792.0,  // Letter height
				841.89, // A4 height
				1008.0, // Legal height
				612.0,  // Letter width (landscape)
			}

			for _, pageHeight := range pageSizes {
				ct := NewCoordTransformer(
					pageHeight,
				)

				// Top of page in OOXML -> top of page in PDF
				_, y := ct.PointToPDFCoord(0, 0)
				if !floatEqualsWithTolerance(
					y,
					pageHeight,
					0.001,
				) {
					t.Errorf(
						"Page %v: top -> %v, want %v",
						pageHeight,
						y,
						pageHeight,
					)
				}

				// Bottom of page in OOXML -> bottom of page in PDF
				_, y = ct.PointToPDFCoord(
					0,
					pageHeight,
				)
				if !floatEqualsWithTolerance(
					y,
					0,
					0.001,
				) {
					t.Errorf(
						"Page %v: bottom -> %v, want 0",
						pageHeight,
						y,
					)
				}
			}
		},
	)
}
