package drawing

import (
	"math"
	"strings"
	"testing"
)

func TestNewColor(t *testing.T) {
	c := NewColor(0.5, 0.25, 0.75, 0.9)
	if c.R != 0.5 || c.G != 0.25 || c.B != 0.75 ||
		c.A != 0.9 {
		t.Errorf(
			"NewColor: expected (0.5, 0.25, 0.75, 0.9), got (%v, %v, %v, %v)",
			c.R,
			c.G,
			c.B,
			c.A,
		)
	}
}

func TestNewColor_Clamping(t *testing.T) {
	c := NewColor(-0.5, 1.5, 0.5, 2.0)
	if c.R != 0 || c.G != 1 || c.A != 1 {
		t.Errorf(
			"NewColor should clamp values to [0,1], got (%v, %v, %v, %v)",
			c.R,
			c.G,
			c.B,
			c.A,
		)
	}
}

func TestNewRGB(t *testing.T) {
	c := NewRGB(0.5, 0.25, 0.75)
	if c.R != 0.5 || c.G != 0.25 || c.B != 0.75 ||
		c.A != 1.0 {
		t.Errorf(
			"NewRGB: expected alpha=1.0, got %v",
			c.A,
		)
	}
}

func TestNewRGB8(t *testing.T) {
	c := NewRGB8(255, 128, 0)
	if c.R != 1.0 {
		t.Errorf(
			"NewRGB8: expected R=1.0, got %v",
			c.R,
		)
	}
	if math.Abs(c.G-0.502) > 0.01 {
		t.Errorf(
			"NewRGB8: expected G~0.502, got %v",
			c.G,
		)
	}
	if c.B != 0 {
		t.Errorf(
			"NewRGB8: expected B=0, got %v",
			c.B,
		)
	}
}

func TestNewRGBA8(t *testing.T) {
	c := NewRGBA8(255, 0, 0, 128)
	if c.R != 1.0 || c.B != 0 {
		t.Errorf(
			"NewRGBA8: unexpected RGB values",
		)
	}
	if math.Abs(c.A-0.502) > 0.01 {
		t.Errorf(
			"NewRGBA8: expected A~0.502, got %v",
			c.A,
		)
	}
}

func TestNewGray(t *testing.T) {
	c := NewGray(0.5)
	if c.R != 0.5 || c.G != 0.5 || c.B != 0.5 {
		t.Errorf(
			"NewGray: expected all channels to be 0.5, got (%v, %v, %v)",
			c.R,
			c.G,
			c.B,
		)
	}
}

func TestNewGray8(t *testing.T) {
	c := NewGray8(128)
	expected := 128.0 / 255.0
	if math.Abs(c.R-expected) > 0.01 {
		t.Errorf(
			"NewGray8: expected ~%v, got %v",
			expected,
			c.R,
		)
	}
}

func TestParseHex(t *testing.T) {
	tests := []struct {
		input   string
		r, g, b uint8
	}{
		{"#FF0000", 255, 0, 0},
		{"#00FF00", 0, 255, 0},
		{"#0000FF", 0, 0, 255},
		{"FF0000", 255, 0, 0}, // Without #
		{"#F00", 255, 0, 0},   // Short format
		{
			"F00",
			255,
			0,
			0,
		}, // Short format without #
		{
			"#abc",
			170,
			187,
			204,
		}, // Short lowercase
		{
			"#ABC",
			170,
			187,
			204,
		}, // Short uppercase
		{"#ffffff", 255, 255, 255},
		{"#000000", 0, 0, 0},
	}

	for _, tt := range tests {
		c := ParseHex(tt.input)
		r, g, b := c.ToRGB8()
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf(
				"ParseHex(%q): expected (%d, %d, %d), got (%d, %d, %d)",
				tt.input,
				tt.r,
				tt.g,
				tt.b,
				r,
				g,
				b,
			)
		}
	}
}

func TestParseHexError(t *testing.T) {
	invalidInputs := []string{
		"",
		"#",
		"#GG0000",
		"#12345",
		"#1234567",
	}

	for _, input := range invalidInputs {
		_, err := ParseHexError(input)
		if err == nil {
			t.Errorf(
				"ParseHexError(%q): expected error, got nil",
				input,
			)
		}
	}
}

func TestColor_ToHex(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{Red, "#FF0000"},
		{Green, "#008000"},
		{Blue, "#0000FF"},
		{Black, "#000000"},
		{White, "#FFFFFF"},
	}

	for _, tt := range tests {
		result := tt.color.ToHex()
		if result != tt.expected {
			t.Errorf(
				"ToHex: expected %q, got %q",
				tt.expected,
				result,
			)
		}
	}
}

func TestColor_ToRGB8(t *testing.T) {
	c := NewRGB(1.0, 0.5, 0.0)
	r, g, b := c.ToRGB8()
	if r != 255 || b != 0 {
		t.Errorf(
			"ToRGB8: unexpected values (%d, %d, %d)",
			r,
			g,
			b,
		)
	}
	// g should be around 128
	if g < 127 || g > 128 {
		t.Errorf(
			"ToRGB8: expected g~128, got %d",
			g,
		)
	}
}

func TestColor_ToRGBA8(t *testing.T) {
	c := NewColor(1.0, 0.5, 0.0, 0.5)
	r, g, b, a := c.ToRGBA8()
	if r != 255 || b != 0 {
		t.Errorf(
			"ToRGBA8: unexpected RGB values (%d, %d, %d)",
			r,
			g,
			b,
		)
	}
	if a < 127 || a > 128 {
		t.Errorf(
			"ToRGBA8: expected a~128, got %d",
			a,
		)
	}
}

func TestColor_IsOpaque(t *testing.T) {
	opaque := NewRGB(1, 0, 0)
	if !opaque.IsOpaque() {
		t.Error(
			"IsOpaque: should return true for alpha=1",
		)
	}

	transparent := NewColor(1, 0, 0, 0.5)
	if transparent.IsOpaque() {
		t.Error(
			"IsOpaque: should return false for alpha<1",
		)
	}
}

func TestColor_IsTransparent(t *testing.T) {
	transparent := NewColor(1, 0, 0, 0)
	if !transparent.IsTransparent() {
		t.Error(
			"IsTransparent: should return true for alpha=0",
		)
	}

	opaque := NewRGB(1, 0, 0)
	if opaque.IsTransparent() {
		t.Error(
			"IsTransparent: should return false for alpha>0",
		)
	}
}

func TestColor_WithAlpha(t *testing.T) {
	c := Red.WithAlpha(0.5)
	if c.R != 1 || c.A != 0.5 {
		t.Errorf(
			"WithAlpha: expected (1, 0, 0, 0.5), got (%v, %v, %v, %v)",
			c.R,
			c.G,
			c.B,
			c.A,
		)
	}
}

func TestColor_Blend(t *testing.T) {
	c := Black.Blend(White, 0.5)
	if math.Abs(c.R-0.5) > 0.001 ||
		math.Abs(c.G-0.5) > 0.001 ||
		math.Abs(c.B-0.5) > 0.001 {
		t.Errorf(
			"Blend: expected (0.5, 0.5, 0.5), got (%v, %v, %v)",
			c.R,
			c.G,
			c.B,
		)
	}
}

func TestColor_Lighten(t *testing.T) {
	c := Black.Lighten(0.5)
	if math.Abs(c.R-0.5) > 0.001 {
		t.Errorf(
			"Lighten: expected R=0.5, got %v",
			c.R,
		)
	}
}

func TestColor_Darken(t *testing.T) {
	c := White.Darken(0.5)
	if math.Abs(c.R-0.5) > 0.001 {
		t.Errorf(
			"Darken: expected R=0.5, got %v",
			c.R,
		)
	}
}

func TestCMYKColor(t *testing.T) {
	cmyk := NewCMYK(0, 1, 1, 0)
	rgb := cmyk.ToRGB()
	// Cyan=0, Magenta=1, Yellow=1, Black=0 should give Red
	if rgb.R != 1 || rgb.G != 0 || rgb.B != 0 {
		t.Errorf(
			"CMYK to RGB: expected red, got (%v, %v, %v)",
			rgb.R,
			rgb.G,
			rgb.B,
		)
	}
}

func TestRGBToCMYK(t *testing.T) {
	cmyk := RGBToCMYK(Red)
	// Red should be C=0, M=1, Y=1, K=0
	if cmyk.C != 0 || cmyk.M != 1 ||
		cmyk.Y != 1 ||
		cmyk.K != 0 {
		t.Errorf(
			"RGB to CMYK: expected (0, 1, 1, 0), got (%v, %v, %v, %v)",
			cmyk.C,
			cmyk.M,
			cmyk.Y,
			cmyk.K,
		)
	}
}

func TestRGBToCMYK_Black(t *testing.T) {
	cmyk := RGBToCMYK(Black)
	// Black should be K=1
	if cmyk.K != 1 {
		t.Errorf(
			"RGB to CMYK for black: expected K=1, got %v",
			cmyk.K,
		)
	}
}

func TestHSLColor(t *testing.T) {
	hsl := NewHSL(0, 1, 0.5)
	rgb := hsl.ToRGB()
	// H=0, S=1, L=0.5 should give Red
	if math.Abs(rgb.R-1) > 0.01 || rgb.G > 0.01 ||
		rgb.B > 0.01 {
		t.Errorf(
			"HSL to RGB: expected red, got (%v, %v, %v)",
			rgb.R,
			rgb.G,
			rgb.B,
		)
	}
}

func TestHSLColor_Gray(t *testing.T) {
	hsl := NewHSL(0, 0, 0.5)
	rgb := hsl.ToRGB()
	// S=0 should give gray
	if math.Abs(rgb.R-0.5) > 0.01 ||
		math.Abs(rgb.G-0.5) > 0.01 ||
		math.Abs(rgb.B-0.5) > 0.01 {
		t.Errorf(
			"HSL to RGB gray: expected (0.5, 0.5, 0.5), got (%v, %v, %v)",
			rgb.R,
			rgb.G,
			rgb.B,
		)
	}
}

func TestRGBToHSL(t *testing.T) {
	hsl := RGBToHSL(Red)
	// Red should be H=0, S=1, L=0.5
	if math.Abs(hsl.H) > 0.01 ||
		math.Abs(hsl.S-1) > 0.01 ||
		math.Abs(hsl.L-0.5) > 0.01 {
		t.Errorf(
			"RGB to HSL: expected (0, 1, 0.5), got (%v, %v, %v)",
			hsl.H,
			hsl.S,
			hsl.L,
		)
	}
}

func TestRGBToHSL_Gray(t *testing.T) {
	hsl := RGBToHSL(Gray)
	// Gray should have S=0
	if hsl.S > 0.01 {
		t.Errorf(
			"RGB to HSL gray: expected S=0, got %v",
			hsl.S,
		)
	}
}

func TestSetFillRGB(t *testing.T) {
	result := Red.SetFillRGB()
	if result != "1 0 0 rg" {
		t.Errorf(
			"SetFillRGB: expected '1 0 0 rg', got %q",
			result,
		)
	}
}

func TestSetStrokeRGB(t *testing.T) {
	result := Blue.SetStrokeRGB()
	if result != "0 0 1 RG" {
		t.Errorf(
			"SetStrokeRGB: expected '0 0 1 RG', got %q",
			result,
		)
	}
}

func TestSetFillGray(t *testing.T) {
	result := SetFillGray(0.5)
	if result != "0.5 g" {
		t.Errorf(
			"SetFillGray: expected '0.5 g', got %q",
			result,
		)
	}
}

func TestSetStrokeGray(t *testing.T) {
	result := SetStrokeGray(0.25)
	if result != "0.25 G" {
		t.Errorf(
			"SetStrokeGray: expected '0.25 G', got %q",
			result,
		)
	}
}

func TestSetFillCMYK(t *testing.T) {
	cmyk := NewCMYK(0, 1, 1, 0)
	result := cmyk.SetFillCMYK()
	if result != "0 1 1 0 k" {
		t.Errorf(
			"SetFillCMYK: expected '0 1 1 0 k', got %q",
			result,
		)
	}
}

func TestSetStrokeCMYK(t *testing.T) {
	cmyk := NewCMYK(1, 0, 0, 0)
	result := cmyk.SetStrokeCMYK()
	if result != "1 0 0 0 K" {
		t.Errorf(
			"SetStrokeCMYK: expected '1 0 0 0 K', got %q",
			result,
		)
	}
}

func TestSetColorSpace(t *testing.T) {
	result := SetColorSpace("DeviceRGB", false)
	if result != "/DeviceRGB cs" {
		t.Errorf(
			"SetColorSpace fill: expected '/DeviceRGB cs', got %q",
			result,
		)
	}

	result = SetColorSpace("DeviceCMYK", true)
	if result != "/DeviceCMYK CS" {
		t.Errorf(
			"SetColorSpace stroke: expected '/DeviceCMYK CS', got %q",
			result,
		)
	}
}

func TestSetColor(t *testing.T) {
	result := SetColor([]float64{1, 0, 0}, false)
	if result != "1 0 0 sc" {
		t.Errorf(
			"SetColor fill: expected '1 0 0 sc', got %q",
			result,
		)
	}

	result = SetColor([]float64{0, 0, 1}, true)
	if result != "0 0 1 SC" {
		t.Errorf(
			"SetColor stroke: expected '0 0 1 SC', got %q",
			result,
		)
	}
}

func TestNamedColors(t *testing.T) {
	// Test a few named colors
	tests := []struct {
		name    string
		r, g, b uint8
	}{
		{"black", 0, 0, 0},
		{"white", 255, 255, 255},
		{"red", 255, 0, 0},
		{"blue", 0, 0, 255},
		{"coral", 255, 127, 80},
	}

	for _, tt := range tests {
		c := NamedColor(tt.name)
		r, g, b := c.ToRGB8()
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf(
				"NamedColor(%q): expected (%d, %d, %d), got (%d, %d, %d)",
				tt.name,
				tt.r,
				tt.g,
				tt.b,
				r,
				g,
				b,
			)
		}
	}
}

func TestNamedColor_CaseInsensitive(
	t *testing.T,
) {
	c1 := NamedColor("red")
	c2 := NamedColor("RED")
	c3 := NamedColor("Red")

	if c1.R != c2.R || c2.R != c3.R {
		t.Error(
			"NamedColor should be case-insensitive",
		)
	}
}

func TestNamedColor_NotFound(t *testing.T) {
	c := NamedColor("notacolor")
	if c != Black {
		t.Error(
			"NamedColor should return black for unknown colors",
		)
	}
}

func TestNamedColorLookup(t *testing.T) {
	c, found := NamedColorLookup("red")
	if !found {
		t.Error(
			"NamedColorLookup should find 'red'",
		)
	}
	if c.R != 1 {
		t.Error(
			"NamedColorLookup: 'red' should have R=1",
		)
	}

	_, found = NamedColorLookup("notacolor")
	if found {
		t.Error(
			"NamedColorLookup should not find 'notacolor'",
		)
	}
}

func TestParseColor(t *testing.T) {
	tests := []struct {
		input   string
		r, g, b float64
	}{
		{"red", 1, 0, 0},
		{"#FF0000", 1, 0, 0},
		{"FF0000", 1, 0, 0},
		{"rgb(255, 0, 0)", 1, 0, 0},
		{"rgb(100%, 0%, 0%)", 1, 0, 0},
	}

	for _, tt := range tests {
		c := ParseColor(tt.input)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"ParseColor(%q): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.input,
				tt.r,
				tt.g,
				tt.b,
				c.R,
				c.G,
				c.B,
			)
		}
	}
}

func TestParseColorError(t *testing.T) {
	_, err := ParseColorError("")
	if err == nil {
		t.Error(
			"ParseColorError should return error for empty string",
		)
	}

	_, err = ParseColorError("notaformat")
	if err == nil {
		t.Error(
			"ParseColorError should return error for unrecognized format",
		)
	}
}

func TestParseColor_RGBA(t *testing.T) {
	c := ParseColor("rgba(255, 0, 0, 0.5)")
	if math.Abs(c.A-0.5) > 0.01 {
		t.Errorf(
			"ParseColor rgba: expected A=0.5, got %v",
			c.A,
		)
	}
}

func TestColor_String(t *testing.T) {
	opaque := Red.String()
	if opaque != "#FF0000" {
		t.Errorf(
			"String opaque: expected '#FF0000', got %q",
			opaque,
		)
	}

	transparent := NewColor(1, 0, 0, 0.5).String()
	if !strings.Contains(transparent, "rgba") {
		t.Errorf(
			"String transparent: expected rgba format, got %q",
			transparent,
		)
	}
}

func TestBasicColors(t *testing.T) {
	// Verify basic color constants are correct
	if Black.R != 0 || Black.G != 0 ||
		Black.B != 0 {
		t.Error("Black should be (0, 0, 0)")
	}
	if White.R != 1 || White.G != 1 ||
		White.B != 1 {
		t.Error("White should be (1, 1, 1)")
	}
	if Red.R != 1 || Red.G != 0 || Red.B != 0 {
		t.Error("Red should be (1, 0, 0)")
	}
	if Blue.R != 0 || Blue.G != 0 || Blue.B != 1 {
		t.Error("Blue should be (0, 0, 1)")
	}
	if Yellow.R != 1 || Yellow.G != 1 ||
		Yellow.B != 0 {
		t.Error("Yellow should be (1, 1, 0)")
	}
	if Cyan.R != 0 || Cyan.G != 1 || Cyan.B != 1 {
		t.Error("Cyan should be (0, 1, 1)")
	}
	if Magenta.R != 1 || Magenta.G != 0 ||
		Magenta.B != 1 {
		t.Error("Magenta should be (1, 0, 1)")
	}
	if !Transparent.IsTransparent() {
		t.Error("Transparent should have A=0")
	}
}

func TestHSLNormalization(t *testing.T) {
	// Test hue normalization
	hsl := NewHSL(-90, 1, 0.5)
	if hsl.H < 0 || hsl.H >= 360 {
		t.Errorf(
			"HSL should normalize negative hue, got %v",
			hsl.H,
		)
	}

	hsl = NewHSL(450, 1, 0.5)
	if hsl.H < 0 || hsl.H >= 360 {
		t.Errorf(
			"HSL should normalize hue > 360, got %v",
			hsl.H,
		)
	}
}

func TestClamp01(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-1, 0},
		{0, 0},
		{0.5, 0.5},
		{1, 1},
		{2, 1},
	}

	for _, tt := range tests {
		result := clamp01(tt.input)
		if result != tt.expected {
			t.Errorf(
				"clamp01(%v): expected %v, got %v",
				tt.input,
				tt.expected,
				result,
			)
		}
	}
}

func TestHueToRGB(t *testing.T) {
	// Test wraparound cases
	result := hueToRGB(
		0,
		1,
		-0.5,
	) // Should wrap to 0.5
	if result < 0 || result > 1 {
		t.Errorf(
			"hueToRGB should return value in [0,1], got %v",
			result,
		)
	}

	result = hueToRGB(
		0,
		1,
		1.5,
	) // Should wrap to 0.5
	if result < 0 || result > 1 {
		t.Errorf(
			"hueToRGB should return value in [0,1], got %v",
			result,
		)
	}
}
