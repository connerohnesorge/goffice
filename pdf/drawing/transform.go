package drawing

import (
	"fmt"
	"math"
	"strings"
)

// Transform represents a 2D affine transformation matrix.
// The matrix is represented as:
//
//	| a  b  0 |
//	| c  d  0 |
//	| e  f  1 |
//
// This corresponds to the PDF transformation matrix [a b c d e f].
// When applied to a point (x, y), the result is:
//
//	x' = a*x + c*y + e
//	y' = b*x + d*y + f
type Transform struct {
	A, B, C, D, E, F float64
}

// Identity returns the identity transformation matrix.
// This is the default transformation that has no effect.
func Identity() Transform {
	return Transform{
		A: 1, B: 0,
		C: 0, D: 1,
		E: 0, F: 0,
	}
}

// NewTransform creates a new identity transformation.
func NewTransform() Transform {
	return Identity()
}

// NewTransformFromMatrix creates a transform from matrix components [a b c d e f].
func NewTransformFromMatrix(
	a, b, c, d, e, f float64,
) Transform {
	return Transform{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

// Translate returns a new transform with a translation applied.
// The translation moves points by (tx, ty).
func (t Transform) Translate(
	tx, ty float64,
) Transform {
	// Translation matrix: [1 0 0 1 tx ty]
	// Multiply: this * translation
	return Transform{
		A: t.A,
		B: t.B,
		C: t.C,
		D: t.D,
		E: t.A*tx + t.C*ty + t.E,
		F: t.B*tx + t.D*ty + t.F,
	}
}

// Scale returns a new transform with scaling applied.
// The scaling multiplies x coordinates by sx and y coordinates by sy.
func (t Transform) Scale(
	sx, sy float64,
) Transform {
	// Scale matrix: [sx 0 0 sy 0 0]
	// Multiply: this * scale
	return Transform{
		A: t.A * sx,
		B: t.B * sx,
		C: t.C * sy,
		D: t.D * sy,
		E: t.E,
		F: t.F,
	}
}

// ScaleUniform returns a new transform with uniform scaling applied.
func (t Transform) ScaleUniform(
	s float64,
) Transform {
	return t.Scale(s, s)
}

// Rotate returns a new transform with rotation applied.
// The rotation is specified in degrees and is counter-clockwise.
func (t Transform) Rotate(
	degrees float64,
) Transform {
	return t.RotateRadians(
		degrees * math.Pi / 180,
	)
}

// RotateRadians returns a new transform with rotation applied.
// The rotation is specified in radians and is counter-clockwise.
func (t Transform) RotateRadians(
	radians float64,
) Transform {
	cos := math.Cos(radians)
	sin := math.Sin(radians)
	// Rotation matrix: [cos sin -sin cos 0 0]
	// Multiply: this * rotation
	return Transform{
		A: t.A*cos + t.C*sin,
		B: t.B*cos + t.D*sin,
		C: -t.A*sin + t.C*cos,
		D: -t.B*sin + t.D*cos,
		E: t.E,
		F: t.F,
	}
}

// RotateAround returns a new transform with rotation around a specific point.
// The rotation is specified in degrees and is counter-clockwise.
func (t Transform) RotateAround(
	degrees, cx, cy float64,
) Transform {
	return t.Translate(cx, cy).
		Rotate(degrees).
		Translate(-cx, -cy)
}

// Skew returns a new transform with skew (shear) applied.
// The skew angles are specified in degrees.
// skewX skews along the x-axis, skewY skews along the y-axis.
func (t Transform) Skew(
	skewXDegrees, skewYDegrees float64,
) Transform {
	tanX := math.Tan(skewXDegrees * math.Pi / 180)
	tanY := math.Tan(skewYDegrees * math.Pi / 180)
	// Skew matrix: [1 tanY tanX 1 0 0]
	// Multiply: this * skew
	return Transform{
		A: t.A + t.C*tanY,
		B: t.B + t.D*tanY,
		C: t.A*tanX + t.C,
		D: t.B*tanX + t.D,
		E: t.E,
		F: t.F,
	}
}

// FlipHorizontal returns a new transform with horizontal flip applied.
func (t Transform) FlipHorizontal() Transform {
	return t.Scale(-1, 1)
}

// FlipVertical returns a new transform with vertical flip applied.
func (t Transform) FlipVertical() Transform {
	return t.Scale(1, -1)
}

// Concat returns the concatenation of this transform with another.
// The result applies the other transform first, then this transform.
// This is equivalent to matrix multiplication: this * other.
func (t Transform) Concat(
	other Transform,
) Transform {
	return Transform{
		A: t.A*other.A + t.C*other.B,
		B: t.B*other.A + t.D*other.B,
		C: t.A*other.C + t.C*other.D,
		D: t.B*other.C + t.D*other.D,
		E: t.A*other.E + t.C*other.F + t.E,
		F: t.B*other.E + t.D*other.F + t.F,
	}
}

// Inverse returns the inverse of this transform.
// If the transform is not invertible (determinant is zero), the identity is returned.
func (t Transform) Inverse() Transform {
	det := t.Determinant()
	if math.Abs(det) < 1e-10 {
		return Identity()
	}
	invDet := 1 / det

	return Transform{
		A: t.D * invDet,
		B: -t.B * invDet,
		C: -t.C * invDet,
		D: t.A * invDet,
		E: (t.C*t.F - t.D*t.E) * invDet,
		F: (t.B*t.E - t.A*t.F) * invDet,
	}
}

// Determinant returns the determinant of the transform matrix.
func (t Transform) Determinant() float64 {
	return t.A*t.D - t.B*t.C
}

// IsIdentity returns true if this is the identity transform.
func (t Transform) IsIdentity() bool {
	const epsilon = 1e-10

	return math.Abs(t.A-1) < epsilon &&
		math.Abs(t.B) < epsilon &&
		math.Abs(t.C) < epsilon &&
		math.Abs(t.D-1) < epsilon &&
		math.Abs(t.E) < epsilon &&
		math.Abs(t.F) < epsilon
}

// TransformPoint applies the transform to a point and returns the result.
func (t Transform) TransformPoint(
	x, y float64,
) (tx, ty float64) {
	return t.A*x + t.C*y + t.E, t.B*x + t.D*y + t.F
}

// TransformVector applies the transform to a vector (ignoring translation).
func (t Transform) TransformVector(
	x, y float64,
) (vx, vy float64) {
	return t.A*x + t.C*y, t.B*x + t.D*y
}

// TransformDistance transforms a distance value.
// This is the average scale factor of the transform.
func (t Transform) TransformDistance(
	d float64,
) float64 {
	// Use the geometric mean of the scale factors
	scaleX := math.Sqrt(t.A*t.A + t.B*t.B)
	scaleY := math.Sqrt(t.C*t.C + t.D*t.D)

	return d * math.Sqrt(scaleX*scaleY)
}

// ToCMOperator returns the PDF cm (concatenate matrix) operator.
// This is used to set the current transformation matrix in a content stream.
func (t Transform) ToCMOperator() string {
	return fmt.Sprintf("%s %s %s %s %s %s cm",
		formatFloat(t.A), formatFloat(t.B),
		formatFloat(t.C), formatFloat(t.D),
		formatFloat(t.E), formatFloat(t.F))
}

// Array returns the transform as a PDF array [a b c d e f].
func (t Transform) Array() [6]float64 {
	return [6]float64{
		t.A,
		t.B,
		t.C,
		t.D,
		t.E,
		t.F,
	}
}

// String returns a string representation of the transform.
func (t Transform) String() string {
	return fmt.Sprintf(
		"Transform[a=%g b=%g c=%g d=%g e=%g f=%g]",
		t.A,
		t.B,
		t.C,
		t.D,
		t.E,
		t.F,
	)
}

// GraphicsState manages a stack of graphics states for PDF content streams.
// It tracks the current transformation matrix and other graphics state properties.
type GraphicsState struct {
	stack     []graphicsStateEntry
	current   graphicsStateEntry
	operators strings.Builder
}

type graphicsStateEntry struct {
	transform Transform
	// Future: add other graphics state properties like line width, color, etc.
}

// NewGraphicsState creates a new graphics state manager.
func NewGraphicsState() *GraphicsState {
	return &GraphicsState{
		current: graphicsStateEntry{
			transform: Identity(),
		},
	}
}

// Save saves the current graphics state onto the stack.
// This outputs the PDF 'q' operator.
func (gs *GraphicsState) Save() *GraphicsState {
	gs.stack = append(gs.stack, gs.current)
	gs.operators.WriteString("q\n")

	return gs
}

// Restore restores the graphics state from the stack.
// This outputs the PDF 'Q' operator.
func (gs *GraphicsState) Restore() *GraphicsState {
	if len(gs.stack) > 0 {
		gs.current = gs.stack[len(gs.stack)-1]
		gs.stack = gs.stack[:len(gs.stack)-1]
	}
	gs.operators.WriteString("Q\n")

	return gs
}

// Transform returns the current transformation matrix.
func (gs *GraphicsState) Transform() Transform {
	return gs.current.transform
}

// SetTransform sets the transformation matrix.
// This outputs the PDF 'cm' operator.
func (gs *GraphicsState) SetTransform(
	t Transform,
) *GraphicsState {
	// Calculate the delta transform from current to new
	delta := gs.current.transform.Inverse().
		Concat(t)
	if !delta.IsIdentity() {
		gs.operators.WriteString(
			delta.ToCMOperator(),
		)
		gs.operators.WriteByte('\n')
		gs.current.transform = t
	}

	return gs
}

// ApplyTransform concatenates a transform with the current transformation matrix.
// This outputs the PDF 'cm' operator.
func (gs *GraphicsState) ApplyTransform(
	t Transform,
) *GraphicsState {
	if !t.IsIdentity() {
		gs.operators.WriteString(t.ToCMOperator())
		gs.operators.WriteByte('\n')
		gs.current.transform = gs.current.transform.Concat(
			t,
		)
	}

	return gs
}

// Translate adds a translation to the current transformation.
func (gs *GraphicsState) Translate(
	tx, ty float64,
) *GraphicsState {
	return gs.ApplyTransform(
		Identity().Translate(tx, ty),
	)
}

// Scale adds scaling to the current transformation.
func (gs *GraphicsState) Scale(
	sx, sy float64,
) *GraphicsState {
	return gs.ApplyTransform(
		Identity().Scale(sx, sy),
	)
}

// Rotate adds rotation to the current transformation.
// The rotation is specified in degrees and is counter-clockwise.
func (gs *GraphicsState) Rotate(
	degrees float64,
) *GraphicsState {
	return gs.ApplyTransform(
		Identity().Rotate(degrees),
	)
}

// RotateAround adds rotation around a specific point.
func (gs *GraphicsState) RotateAround(
	degrees, cx, cy float64,
) *GraphicsState {
	return gs.ApplyTransform(
		Identity().RotateAround(degrees, cx, cy),
	)
}

// String returns the accumulated operators.
func (gs *GraphicsState) String() string {
	return gs.operators.String()
}

// Reset clears the accumulated operators and resets to identity transform.
func (gs *GraphicsState) Reset() *GraphicsState {
	gs.stack = gs.stack[:0]
	gs.current = graphicsStateEntry{
		transform: Identity(),
	}
	gs.operators.Reset()

	return gs
}

// WriteOperator writes a raw operator to the content stream.
func (gs *GraphicsState) WriteOperator(
	op string,
) *GraphicsState {
	gs.operators.WriteString(op)
	gs.operators.WriteByte('\n')

	return gs
}

// formatFloat formats a float64 for PDF output.
// It removes trailing zeros and uses a reasonable precision.
func formatFloat(f float64) string {
	// Round to 6 decimal places to avoid floating point noise
	f = math.Round(f*1000000) / 1000000

	// Handle special cases
	if f == 0 {
		return "0"
	}
	if f == 1 {
		return "1"
	}
	if f == -1 {
		return "-1"
	}

	// Format with up to 6 decimal places
	s := fmt.Sprintf("%.6f", f)

	// Trim trailing zeros
	s = strings.TrimRight(s, "0")

	// Trim trailing decimal point
	s = strings.TrimRight(s, ".")

	return s
}
