package transform

import (
	"math"
)

// singularityTolerance is the threshold for detecting singular matrices.
const singularityTolerance = 1e-10

// Matrix represents a 2D affine transformation matrix.
// The matrix is represented in the form:
//
//	[A  C  E]
//	[B  D  F]
//	[0  0  1]
//
// This allows for translation, rotation, scaling, and shearing transformations.
// Matrix multiplication order: rightmost matrix is applied FIRST to points.
type Matrix struct {
	A, B, C, D, E, F float64
}

// Identity returns the identity matrix (no transformation).
// The identity matrix leaves points unchanged when applied.
func Identity() Matrix {
	return Matrix{A: 1, B: 0, C: 0, D: 1, E: 0, F: 0}
}

// Translate creates a translation matrix.
// dx and dy specify the translation amounts in the X and Y directions.
func Translate(dx, dy float64) Matrix {
	return Matrix{A: 1, B: 0, C: 0, D: 1, E: dx, F: dy}
}

// Rotate creates a rotation matrix.
// angle is specified in radians. Positive angles rotate counter-clockwise.
func Rotate(angle float64) Matrix {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	return Matrix{A: cos, B: sin, C: -sin, D: cos, E: 0, F: 0}
}

// Scale creates a scaling matrix.
// sx and sy specify the scale factors in the X and Y directions.
func Scale(sx, sy float64) Matrix {
	return Matrix{A: sx, B: 0, C: 0, D: sy, E: 0, F: 0}
}

// Multiply returns the product of two matrices (m1 * m2).
// Note: order matters! Matrix multiplication is not commutative.
// The rightmost matrix is applied FIRST to points.
func Multiply(m1, m2 Matrix) Matrix {
	return Matrix{
		A: m1.A*m2.A + m1.C*m2.B,
		B: m1.B*m2.A + m1.D*m2.B,
		C: m1.A*m2.C + m1.C*m2.D,
		D: m1.B*m2.C + m1.D*m2.D,
		E: m1.A*m2.E + m1.C*m2.F + m1.E,
		F: m1.B*m2.E + m1.D*m2.F + m1.F,
	}
}

// Invert returns the inverse matrix.
// If the matrix is singular (determinant is zero or near-zero),
// returns the identity matrix as a fallback.
// Mathematically: m * m.Invert() = Identity()
func (m Matrix) Invert() Matrix {
	det := m.A*m.D - m.B*m.C
	// Check for near-zero determinant (singular matrix)
	if math.Abs(det) < singularityTolerance {
		return Identity() // Singular matrix - cannot invert
	}

	return Matrix{
		A: m.D / det,
		B: -m.B / det,
		C: -m.C / det,
		D: m.A / det,
		E: (m.C*m.F - m.D*m.E) / det,
		F: (m.B*m.E - m.A*m.F) / det,
	}
}

// TransformPoint applies the matrix transformation to a point.
// Returns the transformed coordinates.
func (m Matrix) TransformPoint(x, y float64) (newX, newY float64) {
	return m.A*x + m.C*y + m.E, m.B*x + m.D*y + m.F
}
