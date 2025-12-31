package core

// RGB represents a simple RGB color with components in the range [0, 1].
// This is a minimal color type used for basic rendering operations.
type RGB struct {
	R, G, B float64
}

// NewRGB creates a new RGB color with the specified values.
// Values are clamped to the range [0, 1].
func NewRGB(r, g, b float64) RGB {
	return RGB{
		R: clamp01(r),
		G: clamp01(g),
		B: clamp01(b),
	}
}

// clamp01 clamps a value to the range [0, 1].
func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}

	return v
}
