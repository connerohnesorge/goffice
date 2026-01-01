package transform

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func TestIdentity(t *testing.T) {
	m := Identity()
	if m.A != 1 || m.B != 0 || m.C != 0 || m.D != 1 || m.E != 0 || m.F != 0 {
		t.Errorf("Identity() = %+v, want identity matrix", m)
	}

	// Identity should leave points unchanged
	x, y := m.TransformPoint(10, 20)
	if x != 10 || y != 20 {
		t.Errorf("Identity().TransformPoint(10, 20) = (%f, %f), want (10, 20)", x, y)
	}
}

func TestTranslate(t *testing.T) {
	m := Translate(100, 200)
	x, y := m.TransformPoint(10, 20)
	if math.Abs(x-110) > epsilon || math.Abs(y-220) > epsilon {
		t.Errorf("Translate(100, 200).TransformPoint(10, 20) = (%f, %f), want (110, 220)", x, y)
	}
}

func TestRotate(t *testing.T) {
	// Rotate 90 degrees counter-clockwise
	m := Rotate(math.Pi / 2)
	x, y := m.TransformPoint(1, 0)
	if math.Abs(x-0) > epsilon || math.Abs(y-1) > epsilon {
		t.Errorf("Rotate(90°).TransformPoint(1, 0) = (%f, %f), want (0, 1)", x, y)
	}

	// Rotate 180 degrees
	m = Rotate(math.Pi)
	x, y = m.TransformPoint(1, 0)
	if math.Abs(x-(-1)) > epsilon || math.Abs(y-0) > epsilon {
		t.Errorf("Rotate(180°).TransformPoint(1, 0) = (%f, %f), want (-1, 0)", x, y)
	}
}

func TestScale(t *testing.T) {
	m := Scale(2, 3)
	x, y := m.TransformPoint(10, 20)
	if math.Abs(x-20) > epsilon || math.Abs(y-60) > epsilon {
		t.Errorf("Scale(2, 3).TransformPoint(10, 20) = (%f, %f), want (20, 60)", x, y)
	}
}

func TestMultiply(t *testing.T) {
	// Translate then rotate should give different result than rotate then translate
	translate := Translate(100, 0)
	rotate := Rotate(math.Pi / 2) // 90 degrees

	// Translate THEN rotate (multiply in reverse order: rotate * translate)
	m1 := Multiply(rotate, translate)
	x1, y1 := m1.TransformPoint(0, 0)
	// Point (0,0) -> (100, 0) after translate -> (0, 100) after rotate
	if math.Abs(x1-0) > epsilon || math.Abs(y1-100) > epsilon {
		t.Errorf("Rotate * Translate applied to (0, 0) = (%f, %f), want (0, 100)", x1, y1)
	}

	// Rotate THEN translate (multiply in reverse order: translate * rotate)
	m2 := Multiply(translate, rotate)
	x2, y2 := m2.TransformPoint(0, 0)
	// Point (0,0) -> (0, 0) after rotate -> (100, 0) after translate
	if math.Abs(x2-100) > epsilon || math.Abs(y2-0) > epsilon {
		t.Errorf("Translate * Rotate applied to (0, 0) = (%f, %f), want (100, 0)", x2, y2)
	}
}

func TestInvert(t *testing.T) {
	// Test basic matrix inversion
	m := Matrix{A: 2, B: 0, C: 0, D: 3, E: 10, F: 20}
	inv := m.Invert()

	// m * inv should equal identity
	result := Multiply(m, inv)

	if math.Abs(result.A-1) > epsilon {
		t.Errorf("(M * M⁻¹).A = %f, want 1", result.A)
	}
	if math.Abs(result.B-0) > epsilon {
		t.Errorf("(M * M⁻¹).B = %f, want 0", result.B)
	}
	if math.Abs(result.C-0) > epsilon {
		t.Errorf("(M * M⁻¹).C = %f, want 0", result.C)
	}
	if math.Abs(result.D-1) > epsilon {
		t.Errorf("(M * M⁻¹).D = %f, want 1", result.D)
	}
	if math.Abs(result.E-0) > epsilon {
		t.Errorf("(M * M⁻¹).E = %f, want 0", result.E)
	}
	if math.Abs(result.F-0) > epsilon {
		t.Errorf("(M * M⁻¹).F = %f, want 0", result.F)
	}
}

func TestInvertSingularMatrix(t *testing.T) {
	// Zero determinant (scale 0 in X direction)
	m := Matrix{A: 0, B: 0, C: 0, D: 1, E: 0, F: 0}
	inv := m.Invert()

	// Should return identity as fallback
	identity := Identity()
	if inv != identity {
		t.Errorf("Invert(singular matrix) = %+v, want identity %+v", inv, identity)
	}
}

func TestTransformPoint(t *testing.T) {
	// Complex transformation: scale then rotate then translate
	scale := Scale(2, 2)
	rotate := Rotate(math.Pi / 4) // 45 degrees
	translate := Translate(100, 200)

	// Compose: translate * rotate * scale (applied right to left)
	m := Multiply(translate, Multiply(rotate, scale))

	// Transform point (1, 0)
	x, y := m.TransformPoint(1, 0)

	// Point (1, 0) -> (2, 0) after scale -> (√2, √2) after rotate -> (100+√2, 200+√2) after translate
	sqrt2 := math.Sqrt(2)
	expectedX := 100 + sqrt2
	expectedY := 200 + sqrt2

	if math.Abs(x-expectedX) > epsilon || math.Abs(y-expectedY) > epsilon {
		t.Errorf("Complex transform of (1, 0) = (%f, %f), want (%f, %f)", x, y, expectedX, expectedY)
	}
}

func TestMatrixComposition(t *testing.T) {
	// Test that composing 3 transforms works correctly
	m1 := Translate(10, 20)
	m2 := Rotate(math.Pi / 6) // 30 degrees
	m3 := Scale(2, 3)

	// Compose: m1 * m2 * m3
	result := Multiply(m1, Multiply(m2, m3))

	// Apply to test point
	x, y := result.TransformPoint(5, 5)

	// Manual calculation:
	// (5, 5) -> (10, 15) after scale
	// -> rotate 30°: x' = 10*cos(30°) - 15*sin(30°), y' = 10*sin(30°) + 15*cos(30°)
	cos30 := math.Cos(math.Pi / 6)
	sin30 := math.Sin(math.Pi / 6)
	x2 := 10*cos30 - 15*sin30
	y2 := 10*sin30 + 15*cos30
	// -> translate: (x2 + 10, y2 + 20)
	expectedX := x2 + 10
	expectedY := y2 + 20

	if math.Abs(x-expectedX) > epsilon || math.Abs(y-expectedY) > epsilon {
		t.Errorf("3-matrix composition transform = (%f, %f), want (%f, %f)", x, y, expectedX, expectedY)
	}
}
