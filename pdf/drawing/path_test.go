package drawing

import (
	"math"
	"strings"
	"testing"
)

func TestPathBuilder_MoveTo(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(100, 200)

	result := pb.String()
	if result != "100 200 m\n" {
		t.Errorf(
			"MoveTo: expected '100 200 m\\n', got %q",
			result,
		)
	}

	x, y := pb.CurrentPoint()
	if x != 100 || y != 200 {
		t.Errorf(
			"CurrentPoint: expected (100, 200), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestPathBuilder_LineTo(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0)
	pb.LineTo(100, 100)

	result := pb.String()
	expected := "0 0 m\n100 100 l\n"
	if result != expected {
		t.Errorf(
			"LineTo: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestPathBuilder_LineTo_AutoMoveTo(
	t *testing.T,
) {
	pb := NewPathBuilder()
	// LineTo without MoveTo should auto-move
	pb.LineTo(50, 50)

	result := pb.String()
	expected := "50 50 m\n"
	if result != expected {
		t.Errorf(
			"LineTo auto-move: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestPathBuilder_CurveTo(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0)
	pb.CurveTo(10, 20, 30, 40, 50, 50)

	result := pb.String()
	expected := "0 0 m\n10 20 30 40 50 50 c\n"
	if result != expected {
		t.Errorf(
			"CurveTo: expected %q, got %q",
			expected,
			result,
		)
	}

	x, y := pb.CurrentPoint()
	if x != 50 || y != 50 {
		t.Errorf(
			"CurrentPoint after CurveTo: expected (50, 50), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestPathBuilder_QuadraticCurveTo(
	t *testing.T,
) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0)
	pb.QuadraticCurveTo(50, 100, 100, 0)

	result := pb.String()
	// Quadratic is converted to cubic Bezier
	if !strings.Contains(result, " c\n") {
		t.Errorf(
			"QuadraticCurveTo should generate cubic curve, got %q",
			result,
		)
	}

	x, y := pb.CurrentPoint()
	if x != 100 || y != 0 {
		t.Errorf(
			"CurrentPoint after QuadraticCurveTo: expected (100, 0), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestPathBuilder_ClosePath(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0)
	pb.LineTo(100, 0)
	pb.LineTo(100, 100)
	pb.ClosePath()

	result := pb.String()
	expected := "0 0 m\n100 0 l\n100 100 l\nh\n"
	if result != expected {
		t.Errorf(
			"ClosePath: expected %q, got %q",
			expected,
			result,
		)
	}

	// Current point should be back at start
	x, y := pb.CurrentPoint()
	if x != 0 || y != 0 {
		t.Errorf(
			"CurrentPoint after ClosePath: expected (0, 0), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestPathBuilder_Rectangle(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(10, 20, 100, 50)

	result := pb.String()
	expected := "10 20 100 50 re\n"
	if result != expected {
		t.Errorf(
			"Rectangle: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestPathBuilder_RoundedRect(t *testing.T) {
	pb := NewPathBuilder()
	pb.RoundedRect(0, 0, 100, 50, 10)

	result := pb.String()
	// Should contain moveto, lineto, and curveto commands
	if !strings.Contains(result, " m\n") {
		t.Error(
			"RoundedRect should contain moveto",
		)
	}
	if !strings.Contains(result, " l\n") {
		t.Error(
			"RoundedRect should contain lineto",
		)
	}
	if !strings.Contains(result, " c\n") {
		t.Error(
			"RoundedRect should contain curveto",
		)
	}
	if !strings.Contains(result, "h\n") {
		t.Error(
			"RoundedRect should close the path",
		)
	}
}

func TestPathBuilder_Circle(t *testing.T) {
	pb := NewPathBuilder()
	pb.Circle(50, 50, 25)

	result := pb.String()
	// Circle is made of 4 cubic Bezier curves
	if strings.Count(result, " c\n") != 4 {
		t.Errorf(
			"Circle should have 4 curves, got: %q",
			result,
		)
	}
	if !strings.Contains(result, "h\n") {
		t.Error("Circle should close the path")
	}
}

func TestPathBuilder_Ellipse(t *testing.T) {
	pb := NewPathBuilder()
	pb.Ellipse(50, 50, 30, 20)

	result := pb.String()
	// Ellipse is made of 4 cubic Bezier curves
	if strings.Count(result, " c\n") != 4 {
		t.Errorf(
			"Ellipse should have 4 curves, got: %q",
			result,
		)
	}
}

func TestPathBuilder_Arc(t *testing.T) {
	pb := NewPathBuilder()
	pb.Arc(50, 50, 25, 0, 90)

	result := pb.String()
	// Arc should contain curve commands
	if !strings.Contains(result, " c\n") {
		t.Error(
			"Arc should contain curveto commands",
		)
	}
}

func TestPathBuilder_Polygon(t *testing.T) {
	pb := NewPathBuilder()
	pb.Polygon(0, 0, 100, 0, 50, 100)

	result := pb.String()
	expected := "0 0 m\n100 0 l\n50 100 l\nh\n"
	if result != expected {
		t.Errorf(
			"Polygon: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestPathBuilder_Polyline(t *testing.T) {
	pb := NewPathBuilder()
	pb.Polyline(0, 0, 100, 0, 50, 100)

	result := pb.String()
	expected := "0 0 m\n100 0 l\n50 100 l\n"
	if result != expected {
		t.Errorf(
			"Polyline: expected %q, got %q",
			expected,
			result,
		)
	}
	// Polyline should NOT close the path
	if strings.Contains(result, "h\n") {
		t.Error(
			"Polyline should not close the path",
		)
	}
}

func TestPathBuilder_RegularPolygon(
	t *testing.T,
) {
	pb := NewPathBuilder()
	pb.RegularPolygon(50, 50, 25, 6, 0) // Hexagon

	result := pb.String()
	// Should have moveto, 5 lineto, and closepath
	if !strings.Contains(result, " m\n") {
		t.Error(
			"RegularPolygon should contain moveto",
		)
	}
	if strings.Count(result, " l\n") != 5 {
		t.Errorf(
			"RegularPolygon(6 sides) should have 5 lineto commands, got %d",
			strings.Count(result, " l\n"),
		)
	}
	if !strings.Contains(result, "h\n") {
		t.Error(
			"RegularPolygon should close the path",
		)
	}
}

func TestPathBuilder_Star(t *testing.T) {
	pb := NewPathBuilder()
	pb.Star(
		50,
		50,
		25,
		10,
		5,
		-90,
	) // 5-point star

	result := pb.String()
	// Star with 5 points should have 10 line segments (alternating inner/outer)
	if !strings.Contains(result, " m\n") {
		t.Error("Star should contain moveto")
	}
	if strings.Count(result, " l\n") != 10 {
		t.Errorf(
			"Star(5 points) should have 10 lineto commands, got %d",
			strings.Count(result, " l\n"),
		)
	}
	if !strings.Contains(result, "h\n") {
		t.Error("Star should close the path")
	}
}

func TestPathBuilder_Stroke(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0)
	pb.LineTo(100, 100)

	result := pb.Stroke()
	if !strings.HasSuffix(result, "S\n") {
		t.Errorf(
			"Stroke: should end with 'S\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_Fill(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(0, 0, 100, 100)

	result := pb.Fill()
	if !strings.HasSuffix(result, "f\n") {
		t.Errorf(
			"Fill: should end with 'f\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_FillEvenOdd(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(0, 0, 100, 100)

	result := pb.FillEvenOdd()
	if !strings.HasSuffix(result, "f*\n") {
		t.Errorf(
			"FillEvenOdd: should end with 'f*\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_FillAndStroke(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(0, 0, 100, 100)

	result := pb.FillAndStroke()
	if !strings.HasSuffix(result, "B\n") {
		t.Errorf(
			"FillAndStroke: should end with 'B\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_ClipNonZero(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(0, 0, 100, 100)

	result := pb.ClipNonZero()
	if !strings.HasSuffix(result, "W n\n") {
		t.Errorf(
			"ClipNonZero: should end with 'W n\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_ClipEvenOdd(t *testing.T) {
	pb := NewPathBuilder()
	pb.Rectangle(0, 0, 100, 100)

	result := pb.ClipEvenOdd()
	if !strings.HasSuffix(result, "W* n\n") {
		t.Errorf(
			"ClipEvenOdd: should end with 'W* n\\n', got %q",
			result,
		)
	}
}

func TestPathBuilder_Reset(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(100, 100)
	pb.LineTo(200, 200)

	pb.Reset()

	if pb.String() != "" {
		t.Error(
			"Reset should clear the path operators",
		)
	}
	if pb.HasPath() {
		t.Error("Reset should clear hasPath flag")
	}
}

func TestPathBuilder_HasPath(t *testing.T) {
	pb := NewPathBuilder()
	if pb.HasPath() {
		t.Error(
			"New PathBuilder should not have a path",
		)
	}

	pb.MoveTo(0, 0)
	if !pb.HasPath() {
		t.Error(
			"PathBuilder should have a path after MoveTo",
		)
	}
}

func TestPathBuilder_FloatFormatting(
	t *testing.T,
) {
	pb := NewPathBuilder()
	pb.MoveTo(1.5, 2.333333)

	result := pb.String()
	// Should format nicely, not with excessive decimal places
	if strings.Contains(result, "2.333333333") {
		t.Errorf(
			"Float formatting should limit decimal places, got %q",
			result,
		)
	}
}

func TestPathBuilder_ComplexPath(t *testing.T) {
	pb := NewPathBuilder()
	pb.MoveTo(0, 0).
		LineTo(100, 0).
		LineTo(100, 100).
		LineTo(0, 100).
		ClosePath().
		MoveTo(25, 25).
		Circle(50, 50, 20)

	result := pb.String()
	// Should have multiple subpaths
	if strings.Count(result, " m\n") < 2 {
		t.Error(
			"Complex path should have multiple moveto commands",
		)
	}
}

// Transform tests

func TestTransform_Identity(t *testing.T) {
	tr := Identity()
	if !tr.IsIdentity() {
		t.Error(
			"Identity transform should return true for IsIdentity()",
		)
	}
}

func TestTransform_Translate(t *testing.T) {
	tr := Identity().Translate(10, 20)
	x, y := tr.TransformPoint(0, 0)
	if x != 10 || y != 20 {
		t.Errorf(
			"Translate: expected (10, 20), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_Scale(t *testing.T) {
	tr := Identity().Scale(2, 3)
	x, y := tr.TransformPoint(10, 10)
	if x != 20 || y != 30 {
		t.Errorf(
			"Scale: expected (20, 30), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_ScaleUniform(t *testing.T) {
	tr := Identity().ScaleUniform(2)
	x, y := tr.TransformPoint(10, 15)
	if x != 20 || y != 30 {
		t.Errorf(
			"ScaleUniform: expected (20, 30), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_Rotate(t *testing.T) {
	tr := Identity().Rotate(90)
	x, y := tr.TransformPoint(10, 0)
	// 90 degree rotation should move (10, 0) to approximately (0, 10)
	if math.Abs(x) > 0.0001 ||
		math.Abs(y-10) > 0.0001 {
		t.Errorf(
			"Rotate 90: expected approximately (0, 10), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_RotateAround(t *testing.T) {
	tr := Identity().RotateAround(180, 50, 50)
	x, y := tr.TransformPoint(100, 50)
	// 180 degree rotation around (50, 50) should move (100, 50) to (0, 50)
	if math.Abs(x) > 0.0001 ||
		math.Abs(y-50) > 0.0001 {
		t.Errorf(
			"RotateAround 180: expected approximately (0, 50), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_Skew(t *testing.T) {
	tr := Identity().Skew(45, 0)
	x, y := tr.TransformPoint(0, 10)
	// Skew X by 45 degrees means tan(45) = 1, so x += y
	if math.Abs(x-10) > 0.0001 || y != 10 {
		t.Errorf(
			"Skew: expected approximately (10, 10), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_FlipHorizontal(t *testing.T) {
	tr := Identity().FlipHorizontal()
	x, y := tr.TransformPoint(10, 20)
	if x != -10 || y != 20 {
		t.Errorf(
			"FlipHorizontal: expected (-10, 20), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_FlipVertical(t *testing.T) {
	tr := Identity().FlipVertical()
	x, y := tr.TransformPoint(10, 20)
	if x != 10 || y != -20 {
		t.Errorf(
			"FlipVertical: expected (10, -20), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_Concat(t *testing.T) {
	t1 := Identity().Translate(10, 0)
	t2 := Identity().Scale(2, 2)
	tr := t1.Concat(t2)

	// t1.Concat(t2) means: apply t2 first, then t1
	// So: scale(2,2) then translate(10,0)
	// Point (5, 5) -> scale -> (10, 10) -> translate -> (20, 10)
	x, y := tr.TransformPoint(5, 5)
	if x != 20 || y != 10 {
		t.Errorf(
			"Concat: expected (20, 10), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_Inverse(t *testing.T) {
	tr := Identity().Translate(10, 20).Scale(2, 2)
	inv := tr.Inverse()
	combined := tr.Concat(inv)

	if !combined.IsIdentity() {
		t.Errorf(
			"Transform concatenated with its inverse should be identity, got %v",
			combined,
		)
	}
}

func TestTransform_Determinant(t *testing.T) {
	tr := Identity()
	det := tr.Determinant()
	if det != 1 {
		t.Errorf(
			"Identity determinant should be 1, got %v",
			det,
		)
	}

	tr = Identity().Scale(2, 3)
	det = tr.Determinant()
	if det != 6 {
		t.Errorf(
			"Scale(2,3) determinant should be 6, got %v",
			det,
		)
	}
}

func TestTransform_TransformVector(t *testing.T) {
	tr := Identity().Translate(100, 100).
		Scale(2, 2)
	x, y := tr.TransformVector(10, 10)
	// Vector should ignore translation
	if x != 20 || y != 20 {
		t.Errorf(
			"TransformVector: expected (20, 20), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestTransform_ToCMOperator(t *testing.T) {
	tr := Identity().Translate(10, 20)
	result := tr.ToCMOperator()
	if !strings.Contains(result, "cm") {
		t.Errorf(
			"ToCMOperator should contain 'cm', got %q",
			result,
		)
	}
	if !strings.Contains(result, "10") ||
		!strings.Contains(result, "20") {
		t.Errorf(
			"ToCMOperator should contain translation values, got %q",
			result,
		)
	}
}

func TestTransform_Array(t *testing.T) {
	tr := Identity().Translate(10, 20)
	arr := tr.Array()
	if arr[4] != 10 || arr[5] != 20 {
		t.Errorf(
			"Array: expected e=10, f=20, got %v",
			arr,
		)
	}
}

func TestTransform_String(t *testing.T) {
	tr := Identity()
	s := tr.String()
	if !strings.Contains(s, "Transform") {
		t.Errorf(
			"String should contain 'Transform', got %q",
			s,
		)
	}
}

// GraphicsState tests

func TestGraphicsState_SaveRestore(t *testing.T) {
	gs := NewGraphicsState()
	gs.Save()
	gs.Translate(100, 100)
	gs.Restore()

	result := gs.String()
	if !strings.Contains(result, "q\n") {
		t.Error("Save should output 'q'")
	}
	if !strings.Contains(result, "Q\n") {
		t.Error("Restore should output 'Q'")
	}
}

func TestGraphicsState_Transform(t *testing.T) {
	gs := NewGraphicsState()
	initial := gs.Transform()
	if !initial.IsIdentity() {
		t.Error(
			"Initial transform should be identity",
		)
	}

	gs.Translate(50, 50)
	tr := gs.Transform()
	if tr.E != 50 || tr.F != 50 {
		t.Errorf(
			"Transform after Translate: expected E=50, F=50, got E=%v, F=%v",
			tr.E,
			tr.F,
		)
	}
}

func TestGraphicsState_ApplyTransform(
	t *testing.T,
) {
	gs := NewGraphicsState()
	tr := Identity().Scale(2, 2)
	gs.ApplyTransform(tr)

	result := gs.String()
	if !strings.Contains(result, "cm\n") {
		t.Error(
			"ApplyTransform should output 'cm' operator",
		)
	}
}

func TestGraphicsState_Reset(t *testing.T) {
	gs := NewGraphicsState()
	gs.Translate(100, 100)
	gs.Scale(2, 2)
	gs.Reset()

	if gs.String() != "" {
		t.Error("Reset should clear operators")
	}
	if !gs.Transform().IsIdentity() {
		t.Error(
			"Reset should reset transform to identity",
		)
	}
}

func TestGraphicsState_WriteOperator(
	t *testing.T,
) {
	gs := NewGraphicsState()
	gs.WriteOperator(
		"0.5 g",
	) // Set gray fill color

	result := gs.String()
	if result != "0.5 g\n" {
		t.Errorf(
			"WriteOperator: expected '0.5 g\\n', got %q",
			result,
		)
	}
}

func TestGraphicsState_NestedSaveRestore(
	t *testing.T,
) {
	gs := NewGraphicsState()
	gs.Save()
	gs.Translate(10, 10)
	gs.Save()
	gs.Scale(2, 2)
	gs.Restore()
	gs.Restore()

	result := gs.String()
	// Should have 2 saves and 2 restores
	if strings.Count(result, "q\n") != 2 {
		t.Errorf(
			"Should have 2 saves, got %d",
			strings.Count(result, "q\n"),
		)
	}
	if strings.Count(result, "Q\n") != 2 {
		t.Errorf(
			"Should have 2 restores, got %d",
			strings.Count(result, "Q\n"),
		)
	}
}

// Test formatFloat helper

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{-1, "-1"},
		{1.5, "1.5"},
		{
			1.0000001,
			"1",
		}, // Should round to 1 (7 decimal places)
		{1.123456, "1.123456"}, // Full precision
		{
			1.1234567,
			"1.123457",
		}, // Rounds to 6 decimal places
		{100, "100"},
		{-100.5, "-100.5"},
	}

	for _, tt := range tests {
		result := formatFloat(tt.input)
		if result != tt.expected {
			t.Errorf(
				"formatFloat(%v): expected %q, got %q",
				tt.input,
				tt.expected,
				result,
			)
		}
	}
}
