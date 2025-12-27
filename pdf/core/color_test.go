package core

import "testing"

func TestNewRGB(t *testing.T) {
	tests := []struct {
		name     string
		r, g, b  float64
		expected RGB
	}{
		{
			name:     "normal values",
			r:        0.5,
			g:        0.7,
			b:        0.3,
			expected: RGB{R: 0.5, G: 0.7, B: 0.3},
		},
		{
			name:     "clamped high",
			r:        1.5,
			g:        2.0,
			b:        1.1,
			expected: RGB{R: 1.0, G: 1.0, B: 1.0},
		},
		{
			name:     "clamped low",
			r:        -0.5,
			g:        -1.0,
			b:        -0.1,
			expected: RGB{R: 0.0, G: 0.0, B: 0.0},
		},
		{
			name:     "black",
			r:        0.0,
			g:        0.0,
			b:        0.0,
			expected: RGB{R: 0.0, G: 0.0, B: 0.0},
		},
		{
			name:     "white",
			r:        1.0,
			g:        1.0,
			b:        1.0,
			expected: RGB{R: 1.0, G: 1.0, B: 1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewRGB(tt.r, tt.g, tt.b)
			if result != tt.expected {
				t.Errorf(
					"NewRGB(%v, %v, %v) = %v, want %v",
					tt.r,
					tt.g,
					tt.b,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestRGB(t *testing.T) {
	// Test direct struct creation
	rgb := RGB{R: 0.5, G: 0.6, B: 0.7}
	if rgb.R != 0.5 || rgb.G != 0.6 ||
		rgb.B != 0.7 {
		t.Errorf(
			"RGB struct creation failed: got %v",
			rgb,
		)
	}

	// Test with spreadsheet usage pattern
	black := RGB{R: 0, G: 0, B: 0}
	white := RGB{R: 1, G: 1, B: 1}

	if black.R != 0 || black.G != 0 ||
		black.B != 0 {
		t.Errorf(
			"black RGB incorrect: got %v",
			black,
		)
	}

	if white.R != 1 || white.G != 1 ||
		white.B != 1 {
		t.Errorf(
			"white RGB incorrect: got %v",
			white,
		)
	}
}

func TestClamp01(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"normal", 0.5, 0.5},
		{"zero", 0.0, 0.0},
		{"one", 1.0, 1.0},
		{"too high", 1.5, 1.0},
		{"too low", -0.5, 0.0},
		{"way too high", 100.0, 1.0},
		{"way too low", -100.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := clamp01(tt.input)
			if result != tt.expected {
				t.Errorf(
					"clamp01(%v) = %v, want %v",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}
