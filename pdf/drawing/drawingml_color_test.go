package drawing

import (
	"math"
	"testing"

	"github.com/connerohnesorge/goffice/drawingml"
)

func TestParseSRGBHex(t *testing.T) {
	tests := []struct {
		hex     string
		r, g, b float64
	}{
		{"FF0000", 1.0, 0.0, 0.0},
		{"00FF00", 0.0, 1.0, 0.0},
		{"0000FF", 0.0, 0.0, 1.0},
		{"FFFFFF", 1.0, 1.0, 1.0},
		{"000000", 0.0, 0.0, 0.0},
		{
			"#FF0000",
			1.0,
			0.0,
			0.0,
		}, // With hash prefix
		{"808080", 0.502, 0.502, 0.502}, // Gray
	}

	for _, tt := range tests {
		c := parseSRGBHex(tt.hex)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"parseSRGBHex(%q): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.hex,
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

func TestParseSRGBHex_Invalid(t *testing.T) {
	invalidInputs := []string{
		"",
		"F00",     // Too short
		"FF00",    // Wrong length
		"FF00000", // Too long
		"GGGGGG",  // Invalid hex
	}

	for _, input := range invalidInputs {
		c := parseSRGBHex(input)
		if c != Black {
			t.Errorf(
				"parseSRGBHex(%q): expected Black for invalid input, got %v",
				input,
				c,
			)
		}
	}
}

func TestResolveHSLColor(t *testing.T) {
	tests := []struct {
		hue, sat, lum int
		r, g, b       float64
	}{
		// Red: H=0, S=100%, L=50%
		{0, 100000, 50000, 1.0, 0.0, 0.0},
		// Green: H=120, S=100%, L=50%
		{7200000, 100000, 50000, 0.0, 1.0, 0.0},
		// Blue: H=240, S=100%, L=50%
		{14400000, 100000, 50000, 0.0, 0.0, 1.0},
		// White: L=100%
		{0, 0, 100000, 1.0, 1.0, 1.0},
		// Black: L=0%
		{0, 0, 0, 0.0, 0.0, 0.0},
		// Gray: S=0%, L=50%
		{0, 0, 50000, 0.5, 0.5, 0.5},
	}

	for _, tt := range tests {
		c := resolveHSLColor(
			tt.hue,
			tt.sat,
			tt.lum,
		)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"resolveHSLColor(%d, %d, %d): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.hue,
				tt.sat,
				tt.lum,
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

func TestResolveSCRGBColor(t *testing.T) {
	tests := []struct {
		r, g, b    int // In percentage units (0-100000)
		er, eg, eb float64
	}{
		{
			100000,
			0,
			0,
			1.0,
			0.0,
			0.0,
		}, // Red
		{
			0,
			100000,
			0,
			0.0,
			1.0,
			0.0,
		}, // Green
		{
			0,
			0,
			100000,
			0.0,
			0.0,
			1.0,
		}, // Blue
		{
			50000,
			50000,
			50000,
			0.5,
			0.5,
			0.5,
		}, // Gray
	}

	for _, tt := range tests {
		c := resolveSCRGBColor(tt.r, tt.g, tt.b)
		if math.Abs(c.R-tt.er) > 0.001 ||
			math.Abs(c.G-tt.eg) > 0.001 ||
			math.Abs(c.B-tt.eb) > 0.001 {
			t.Errorf(
				"resolveSCRGBColor(%d, %d, %d): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.r,
				tt.g,
				tt.b,
				tt.er,
				tt.eg,
				tt.eb,
				c.R,
				c.G,
				c.B,
			)
		}
	}
}

func TestResolvePresetColor(t *testing.T) {
	tests := []struct {
		name    string
		r, g, b uint8
	}{
		{"black", 0, 0, 0},
		{"white", 255, 255, 255},
		{"red", 255, 0, 0},
		{"blue", 0, 0, 255},
		{"coral", 255, 127, 80},
		{"dkBlue", 0, 0, 139},
	}

	for _, tt := range tests {
		c := resolvePresetColor(tt.name)
		r, g, b := c.ToRGB8()
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf(
				"resolvePresetColor(%q): expected (%d, %d, %d), got (%d, %d, %d)",
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

func TestResolvePresetColor_Unknown(
	t *testing.T,
) {
	c := resolvePresetColor("notAColor")
	if c != Black {
		t.Errorf(
			"resolvePresetColor for unknown color should return Black, got %v",
			c,
		)
	}
}

func TestResolveSystemColor(t *testing.T) {
	// Test with lastColor fallback.
	c := resolveSystemColor("window", "FF0000")
	if c.R != 1.0 || c.G != 0.0 || c.B != 0.0 {
		t.Errorf(
			"resolveSystemColor with lastColor: expected red, got %v",
			c,
		)
	}

	// Test without lastColor.
	c = resolveSystemColor("window", "")
	if c != White {
		t.Errorf(
			"resolveSystemColor('window'): expected White, got %v",
			c,
		)
	}

	c = resolveSystemColor("windowText", "")
	if c != Black {
		t.Errorf(
			"resolveSystemColor('windowText'): expected Black, got %v",
			c,
		)
	}
}

func TestResolveSchemeColor(t *testing.T) {
	palette := NewThemePalette()

	tests := []struct {
		name  string
		color Color
	}{
		{"dk1", palette.Dark1},
		{"lt1", palette.Light1},
		{"accent1", palette.Accent1},
		{"hlink", palette.Hyperlink},
	}

	for _, tt := range tests {
		c := resolveSchemeColor(tt.name, palette)
		if c != tt.color {
			t.Errorf(
				"resolveSchemeColor(%q): expected %v, got %v",
				tt.name,
				tt.color,
				c,
			)
		}
	}
}

func TestResolveSchemeColor_NilPalette(
	t *testing.T,
) {
	// Should fall back to default palette.
	c := resolveSchemeColor("accent1", nil)
	expected := DefaultThemePalette.Accent1
	if c != expected {
		t.Errorf(
			"resolveSchemeColor with nil palette: expected default accent1 %v, got %v",
			expected,
			c,
		)
	}
}

func TestDrawingMLColor_Resolve(t *testing.T) {
	// Test sRGB color.
	dmlColor := &DrawingMLColor{
		Type:  ColorTypeSRGB,
		Value: "FF0000",
	}
	c := dmlColor.Resolve(nil)
	if c.R != 1.0 || c.G != 0.0 || c.B != 0.0 {
		t.Errorf(
			"Resolve sRGB: expected red, got %v",
			c,
		)
	}

	// Test HSL color.
	dmlColor = &DrawingMLColor{
		Type:       ColorTypeHSL,
		Hue:        0,
		Saturation: 100000,
		Luminance:  50000,
	}
	c = dmlColor.Resolve(nil)
	if math.Abs(c.R-1.0) > 0.01 || c.G > 0.01 ||
		c.B > 0.01 {
		t.Errorf(
			"Resolve HSL (red): expected approximately red, got %v",
			c,
		)
	}

	// Test preset color.
	dmlColor = &DrawingMLColor{
		Type:  ColorTypePreset,
		Value: "coral",
	}
	c = dmlColor.Resolve(nil)
	r, g, b := c.ToRGB8()
	if r != 255 || g != 127 || b != 80 {
		t.Errorf(
			"Resolve preset coral: expected (255, 127, 80), got (%d, %d, %d)",
			r,
			g,
			b,
		)
	}
}

func TestDrawingMLColor_ResolveWithTransformations(
	t *testing.T,
) {
	// Red with 50% tint (should lighten toward white).
	dmlColor := &DrawingMLColor{
		Type:  ColorTypeSRGB,
		Value: "FF0000",
		Transformations: []ColorTransformation{
			{Type: "tint", Value: 50000},
		},
	}
	c := dmlColor.Resolve(nil)

	// After 50% tint, red should move toward white.
	// R stays 1.0, G and B should be ~0.5.
	if c.R != 1.0 || math.Abs(c.G-0.5) > 0.01 ||
		math.Abs(c.B-0.5) > 0.01 {
		t.Errorf(
			"Resolve sRGB with tint: expected (1.0, ~0.5, ~0.5), got %v",
			c,
		)
	}
}

func TestDrawingMLColor_ResolveNil(t *testing.T) {
	var dmlColor *DrawingMLColor
	c := dmlColor.Resolve(nil)
	if c != Black {
		t.Errorf(
			"Resolve nil color: expected Black, got %v",
			c,
		)
	}
}

func TestDrawingMLColor_ResolveUnknown(
	t *testing.T,
) {
	dmlColor := &DrawingMLColor{
		Type: ColorTypeUnknown,
	}
	c := dmlColor.Resolve(nil)
	if c != Black {
		t.Errorf(
			"Resolve unknown color type: expected Black, got %v",
			c,
		)
	}
}

func TestThemePalette_GetColor(t *testing.T) {
	palette := NewThemePalette()

	// Test direct color lookups.
	tests := []struct {
		value drawingml.SchemeColorValue
		color Color
	}{
		{
			drawingml.SchemeColorDark1,
			palette.Dark1,
		},
		{
			drawingml.SchemeColorLight1,
			palette.Light1,
		},
		{
			drawingml.SchemeColorAccent1,
			palette.Accent1,
		},
		{
			drawingml.SchemeColorAccent2,
			palette.Accent2,
		},
		{
			drawingml.SchemeColorHyperlink,
			palette.Hyperlink,
		},
	}

	for _, tt := range tests {
		c, ok := palette.GetColor(tt.value)
		if !ok {
			t.Errorf(
				"GetColor(%s): expected to find color",
				tt.value,
			)
		}
		if c != tt.color {
			t.Errorf(
				"GetColor(%s): expected %v, got %v",
				tt.value,
				tt.color,
				c,
			)
		}
	}
}

func TestThemePalette_GetColorAliases(
	t *testing.T,
) {
	palette := NewThemePalette()

	// Test color aliases.
	tests := []struct {
		alias    drawingml.SchemeColorValue
		resolved Color
	}{
		{
			drawingml.SchemeColorBackground1,
			palette.Light1,
		},
		{
			drawingml.SchemeColorText1,
			palette.Dark1,
		},
		{
			drawingml.SchemeColorBackground2,
			palette.Light2,
		},
		{
			drawingml.SchemeColorText2,
			palette.Dark2,
		},
	}

	for _, tt := range tests {
		c, ok := palette.GetColor(tt.alias)
		if !ok {
			t.Errorf(
				"GetColor(%s): expected to find aliased color",
				tt.alias,
			)
		}
		if c != tt.resolved {
			t.Errorf(
				"GetColor(%s): expected %v, got %v",
				tt.alias,
				tt.resolved,
				c,
			)
		}
	}
}

func TestThemePalette_GetColorOrDefault(
	t *testing.T,
) {
	palette := NewThemePalette()

	// Known color.
	c := palette.GetColorOrDefault(
		drawingml.SchemeColorAccent1,
	)
	if c != palette.Accent1 {
		t.Errorf(
			"GetColorOrDefault(accent1): expected Accent1, got %v",
			c,
		)
	}

	// Unknown color should return black.
	c = palette.GetColorOrDefault("unknownColor")
	if c != Black {
		t.Errorf(
			"GetColorOrDefault(unknown): expected Black, got %v",
			c,
		)
	}
}

func TestThemePalette_SetColor(t *testing.T) {
	palette := NewThemePalette()
	newColor := ParseHex("#123456")

	palette.SetColor(
		drawingml.SchemeColorAccent1,
		newColor,
	)

	c, _ := palette.GetColor(
		drawingml.SchemeColorAccent1,
	)
	if c != newColor {
		t.Errorf(
			"SetColor: expected %v, got %v",
			newColor,
			c,
		)
	}
}

func TestThemePalette_Clone(t *testing.T) {
	original := NewThemePalette()
	original.Accent1 = Red

	clone := original.Clone()

	// Verify clone has same values.
	if clone.Accent1 != Red {
		t.Errorf(
			"Clone Accent1: expected Red, got %v",
			clone.Accent1,
		)
	}

	// Modify original, verify clone is unaffected.
	original.Accent1 = Blue
	if clone.Accent1 != Red {
		t.Errorf(
			"Clone should be independent: expected Red, got %v",
			clone.Accent1,
		)
	}
}

func TestThemePalette_AccentColors(t *testing.T) {
	palette := NewThemePalette()
	accents := palette.AccentColors()

	if len(accents) != 6 {
		t.Errorf(
			"AccentColors: expected 6 colors, got %d",
			len(accents),
		)
	}

	if accents[0] != palette.Accent1 {
		t.Errorf(
			"AccentColors[0]: expected Accent1",
		)
	}
	if accents[5] != palette.Accent6 {
		t.Errorf(
			"AccentColors[5]: expected Accent6",
		)
	}
}

func TestThemePalette_AccentColor(t *testing.T) {
	palette := NewThemePalette()

	tests := []struct {
		index    int
		expected Color
	}{
		{1, palette.Accent1},
		{2, palette.Accent2},
		{3, palette.Accent3},
		{4, palette.Accent4},
		{5, palette.Accent5},
		{6, palette.Accent6},
		{0, palette.Accent1},  // Invalid index
		{7, palette.Accent1},  // Invalid index
		{-1, palette.Accent1}, // Invalid index
	}

	for _, tt := range tests {
		c := palette.AccentColor(tt.index)
		if c != tt.expected {
			t.Errorf(
				"AccentColor(%d): expected %v, got %v",
				tt.index,
				tt.expected,
				c,
			)
		}
	}
}

func TestOfficeThemes(t *testing.T) {
	// Verify Office themes are defined.
	themes := []string{
		"Office",
		"Grayscale",
		"Blue",
		"Green",
		"Red",
		"Purple",
	}

	for _, name := range themes {
		theme, ok := OfficeThemes[name]
		if !ok {
			t.Errorf(
				"OfficeThemes[%q]: expected to find theme",
				name,
			)

			continue
		}

		// Verify theme has valid colors.
		if theme.Accent1 == (Color{}) {
			t.Errorf(
				"OfficeThemes[%q]: Accent1 should not be zero value",
				name,
			)
		}
	}
}

func TestApplyTint(t *testing.T) {
	// Pure red with various tint amounts.
	tests := []struct {
		value   int
		r, g, b float64
	}{
		{
			0,
			1.0,
			0.0,
			0.0,
		}, // No tint = original red
		{50000, 1.0, 0.5, 0.5}, // 50% tint
		{
			100000,
			1.0,
			1.0,
			1.0,
		}, // 100% tint = white
	}

	for _, tt := range tests {
		c := ApplyTint(Red, tt.value)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"ApplyTint(Red, %d): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.value,
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

func TestApplyShade(t *testing.T) {
	// Pure red with various shade amounts.
	tests := []struct {
		value   int
		r, g, b float64
	}{
		{
			100000,
			1.0,
			0.0,
			0.0,
		}, // 100% shade = original red
		{50000, 0.5, 0.0, 0.0}, // 50% shade
		{
			0,
			0.0,
			0.0,
			0.0,
		}, // 0% shade = black
	}

	for _, tt := range tests {
		c := ApplyShade(Red, tt.value)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"ApplyShade(Red, %d): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.value,
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

func TestApplyAlpha(t *testing.T) {
	tests := []struct {
		value int
		alpha float64
	}{
		{100000, 1.0},
		{50000, 0.5},
		{0, 0.0},
	}

	for _, tt := range tests {
		c := ApplyAlpha(Red, tt.value)
		if math.Abs(c.A-tt.alpha) > 0.01 {
			t.Errorf(
				"ApplyAlpha(Red, %d): expected alpha %.2f, got %.2f",
				tt.value,
				tt.alpha,
				c.A,
			)
		}
	}
}

func TestApplyAlphaOffset(t *testing.T) {
	// Start with 50% alpha.
	base := Red.WithAlpha(0.5)

	tests := []struct {
		value int
		alpha float64
	}{
		{25000, 0.75},  // +25%
		{-25000, 0.25}, // -25%
		{50000, 1.0},   // +50% (clamped to 1.0)
		{-60000, 0.0},  // -60% (clamped to 0.0)
	}

	for _, tt := range tests {
		c := ApplyAlphaOffset(base, tt.value)
		if math.Abs(c.A-tt.alpha) > 0.01 {
			t.Errorf(
				"ApplyAlphaOffset(0.5, %d): expected alpha %.2f, got %.2f",
				tt.value,
				tt.alpha,
				c.A,
			)
		}
	}
}

func TestApplySaturationModulation(t *testing.T) {
	// Create a saturated color.
	hsl := NewHSL(0, 1.0, 0.5)
	base := hsl.ToRGB()

	c := ApplySaturationModulation(
		base,
		50000,
	) // 50%
	resultHSL := RGBToHSL(c)

	if math.Abs(resultHSL.S-0.5) > 0.02 {
		t.Errorf(
			"ApplySaturationModulation(50%%): expected S~0.5, got %.2f",
			resultHSL.S,
		)
	}
}

func TestApplyLuminanceModulation(t *testing.T) {
	// Gray color with L=0.5.
	gray := NewRGB(0.5, 0.5, 0.5)

	c := ApplyLuminanceModulation(
		gray,
		50000,
	) // 50%
	resultHSL := RGBToHSL(c)

	if math.Abs(resultHSL.L-0.25) > 0.02 {
		t.Errorf(
			"ApplyLuminanceModulation(50%%): expected L~0.25, got %.2f",
			resultHSL.L,
		)
	}
}

func TestApplyLuminanceOffset(t *testing.T) {
	// Gray color with L=0.5.
	gray := NewRGB(0.5, 0.5, 0.5)

	c := ApplyLuminanceOffset(gray, 25000) // +25%
	resultHSL := RGBToHSL(c)

	if math.Abs(resultHSL.L-0.75) > 0.02 {
		t.Errorf(
			"ApplyLuminanceOffset(+25%%): expected L~0.75, got %.2f",
			resultHSL.L,
		)
	}
}

func TestApplyComplement(t *testing.T) {
	// Complement of red should be cyan.
	c := ApplyComplement(Red)
	resultHSL := RGBToHSL(c)

	// Red hue is 0, complement should be 180 (cyan).
	if math.Abs(resultHSL.H-180) > 1 {
		t.Errorf(
			"ApplyComplement(Red): expected H~180, got %.2f",
			resultHSL.H,
		)
	}
}

func TestApplyInverse(t *testing.T) {
	c := ApplyInverse(Red)

	// Inverse of red (1,0,0) should be cyan (0,1,1).
	if c.R != 0 || c.G != 1 || c.B != 1 {
		t.Errorf(
			"ApplyInverse(Red): expected (0, 1, 1), got (%.2f, %.2f, %.2f)",
			c.R,
			c.G,
			c.B,
		)
	}
}

func TestApplyGrayscale(t *testing.T) {
	c := ApplyGrayscale(Red)

	// All RGB components should be equal.
	if math.Abs(c.R-c.G) > 0.001 ||
		math.Abs(c.G-c.B) > 0.001 {
		t.Errorf(
			"ApplyGrayscale: expected R=G=B, got (%.4f, %.4f, %.4f)",
			c.R,
			c.G,
			c.B,
		)
	}

	// Using standard luminosity weights.
	expected := 0.299 // Red contributes 29.9% to grayscale
	if math.Abs(c.R-expected) > 0.01 {
		t.Errorf(
			"ApplyGrayscale(Red): expected ~%.3f, got %.3f",
			expected,
			c.R,
		)
	}
}

func TestHSVColor(t *testing.T) {
	// Test red: H=0, S=1, V=1.
	hsv := NewHSV(0, 1, 1)
	rgb := hsv.ToRGB()

	if math.Abs(rgb.R-1.0) > 0.01 ||
		rgb.G > 0.01 ||
		rgb.B > 0.01 {
		t.Errorf(
			"HSV(0, 1, 1) to RGB: expected red, got (%.2f, %.2f, %.2f)",
			rgb.R,
			rgb.G,
			rgb.B,
		)
	}

	// Test green: H=120, S=1, V=1.
	hsv = NewHSV(120, 1, 1)
	rgb = hsv.ToRGB()

	if rgb.R > 0.01 ||
		math.Abs(rgb.G-1.0) > 0.01 ||
		rgb.B > 0.01 {
		t.Errorf(
			"HSV(120, 1, 1) to RGB: expected green, got (%.2f, %.2f, %.2f)",
			rgb.R,
			rgb.G,
			rgb.B,
		)
	}
}

func TestRGBToHSV(t *testing.T) {
	// Red.
	hsv := RGBToHSV(Red)
	if math.Abs(hsv.H) > 1 ||
		math.Abs(hsv.S-1.0) > 0.01 ||
		math.Abs(hsv.V-1.0) > 0.01 {
		t.Errorf(
			"RGBToHSV(Red): expected (0, 1, 1), got (%.2f, %.2f, %.2f)",
			hsv.H,
			hsv.S,
			hsv.V,
		)
	}

	// Gray.
	gray := NewRGB(0.5, 0.5, 0.5)
	hsv = RGBToHSV(gray)
	if hsv.S > 0.01 {
		t.Errorf(
			"RGBToHSV(Gray): expected S=0, got %.2f",
			hsv.S,
		)
	}
}

func TestInterpolate(t *testing.T) {
	// Interpolate between black and white.
	tests := []struct {
		t       float64
		r, g, b float64
	}{
		{0.0, 0.0, 0.0, 0.0},
		{0.5, 0.5, 0.5, 0.5},
		{1.0, 1.0, 1.0, 1.0},
	}

	for _, tt := range tests {
		c := Interpolate(Black, White, tt.t)
		if math.Abs(c.R-tt.r) > 0.01 ||
			math.Abs(c.G-tt.g) > 0.01 ||
			math.Abs(c.B-tt.b) > 0.01 {
			t.Errorf(
				"Interpolate(Black, White, %.1f): expected (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
				tt.t,
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

func TestInterpolateHSL(t *testing.T) {
	// Interpolate between red and blue (should go through magenta, not cyan).
	c := InterpolateHSL(Red, Blue, 0.5)
	hsl := RGBToHSL(c)

	// Midpoint between red (0) and blue (240) going the short way is 300 (magenta).
	// But the interpolation should actually be 120 (green) if going the other way,
	// or 300 (magenta) taking the short path clockwise.
	// The short path from 0 to 240 is through 300 (magenta).
	if math.Abs(hsl.H-300) > 5 {
		t.Errorf(
			"InterpolateHSL(Red, Blue, 0.5): expected H~300 (magenta), got %.2f",
			hsl.H,
		)
	}
}

func TestIsKnownTransformation(t *testing.T) {
	known := []string{
		"tint", "shade", "alpha", "satMod", "lumMod", "lumOff",
		"comp", "inv", "gray", "gamma", "invGamma",
	}

	for _, name := range known {
		if !isKnownTransformation(name) {
			t.Errorf(
				"isKnownTransformation(%q): expected true",
				name,
			)
		}
	}

	unknown := []string{"unknown", "foo", "bar"}
	for _, name := range unknown {
		if isKnownTransformation(name) {
			t.Errorf(
				"isKnownTransformation(%q): expected false",
				name,
			)
		}
	}
}

func TestApplyTransformations_Multiple(
	t *testing.T,
) {
	// Apply multiple transformations in sequence.
	transforms := []ColorTransformation{
		{
			Type:  "shade",
			Value: 50000,
		}, // Darken to 50%
		{
			Type:  "alpha",
			Value: 50000,
		}, // Set alpha to 50%
	}

	c := applyTransformations(Red, transforms)

	// Red should be darkened to ~0.5, and alpha should be 0.5.
	if math.Abs(c.R-0.5) > 0.01 {
		t.Errorf(
			"ApplyTransformations: expected R~0.5, got %.2f",
			c.R,
		)
	}
	if math.Abs(c.A-0.5) > 0.01 {
		t.Errorf(
			"ApplyTransformations: expected A~0.5, got %.2f",
			c.A,
		)
	}
}

func TestApplyHueOffset(t *testing.T) {
	// Offset red hue by 120 degrees (should become green).
	offset := 120 * 60000 // 120 degrees in 60000ths
	c := ApplyHueOffset(Red, offset)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.H-120) > 1 {
		t.Errorf(
			"ApplyHueOffset(Red, 120deg): expected H~120, got %.2f",
			hsl.H,
		)
	}
}

func TestApplyRGBModifiers(t *testing.T) {
	// Test red component modifiers.
	c := ApplyRed(Gray, 100000) // Set red to 100%
	if c.R != 1.0 {
		t.Errorf(
			"ApplyRed: expected R=1.0, got %.2f",
			c.R,
		)
	}

	c = ApplyRedOffset(
		NewRGB(0.5, 0.5, 0.5),
		25000,
	) // +25%
	if math.Abs(c.R-0.75) > 0.01 {
		t.Errorf(
			"ApplyRedOffset: expected R~0.75, got %.2f",
			c.R,
		)
	}

	c = ApplyRedModulation(Red, 50000) // 50%
	if math.Abs(c.R-0.5) > 0.01 {
		t.Errorf(
			"ApplyRedModulation: expected R~0.5, got %.2f",
			c.R,
		)
	}

	// Test green component.
	c = ApplyGreen(Gray, 100000)
	if c.G != 1.0 {
		t.Errorf(
			"ApplyGreen: expected G=1.0, got %.2f",
			c.G,
		)
	}

	// Test blue component.
	c = ApplyBlue(Gray, 100000)
	if c.B != 1.0 {
		t.Errorf(
			"ApplyBlue: expected B=1.0, got %.2f",
			c.B,
		)
	}
}

func TestGammaCorrection(t *testing.T) {
	gray := NewRGB(0.5, 0.5, 0.5)

	// Apply gamma and inverse gamma should return to original.
	c := ApplyGamma(gray)
	c = ApplyInverseGamma(c)

	if math.Abs(c.R-0.5) > 0.01 ||
		math.Abs(c.G-0.5) > 0.01 ||
		math.Abs(c.B-0.5) > 0.01 {
		t.Errorf(
			"Gamma round-trip: expected (0.5, 0.5, 0.5), got (%.2f, %.2f, %.2f)",
			c.R,
			c.G,
			c.B,
		)
	}
}

func TestThemePaletteFromHexColors(t *testing.T) {
	palette := ThemePaletteFromHexColors(
		"#000000", "#FFFFFF",
		"#333333", "#CCCCCC",
		"#FF0000", "#00FF00", "#0000FF",
		"#FFFF00", "#FF00FF", "#00FFFF",
		"#0066CC", "#9900CC",
	)

	if palette.Dark1 != Black {
		t.Errorf(
			"ThemePaletteFromHexColors: Dark1 should be Black",
		)
	}
	if palette.Light1 != White {
		t.Errorf(
			"ThemePaletteFromHexColors: Light1 should be White",
		)
	}
	if palette.Accent1 != Red {
		t.Errorf(
			"ThemePaletteFromHexColors: Accent1 should be Red",
		)
	}
}

func TestApplyAlphaModulation(t *testing.T) {
	base := Red.WithAlpha(0.5)

	tests := []struct {
		value int
		alpha float64
	}{
		{100000, 0.5}, // 100% = no change
		{50000, 0.25}, // 50% = half
		{
			200000,
			1.0,
		}, // 200% = double (clamped to 1.0)
	}

	for _, tt := range tests {
		c := ApplyAlphaModulation(base, tt.value)
		if math.Abs(c.A-tt.alpha) > 0.01 {
			t.Errorf(
				"ApplyAlphaModulation(0.5, %d): expected alpha %.2f, got %.2f",
				tt.value,
				tt.alpha,
				c.A,
			)
		}
	}
}

func TestApplyHue(t *testing.T) {
	// Set hue to 120 degrees (green).
	hueValue := 120 * 60000 // In 60000ths of a degree
	c := ApplyHue(Red, hueValue)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.H-120) > 1 {
		t.Errorf(
			"ApplyHue(Red, 120deg): expected H~120, got %.2f",
			hsl.H,
		)
	}
}

func TestApplyHueModulation(t *testing.T) {
	// Create a color with hue=60 (yellow).
	yellow := NewHSL(60, 1.0, 0.5).ToRGB()

	// Modulate hue by 200% (should double to 120 = green).
	c := ApplyHueModulation(yellow, 200000)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.H-120) > 1 {
		t.Errorf(
			"ApplyHueModulation(60, 200%%): expected H~120, got %.2f",
			hsl.H,
		)
	}
}

func TestApplySaturation(t *testing.T) {
	// Set saturation to 50%.
	c := ApplySaturation(Red, 50000)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.S-0.5) > 0.02 {
		t.Errorf(
			"ApplySaturation(Red, 50%%): expected S~0.5, got %.2f",
			hsl.S,
		)
	}
}

func TestApplyLuminance(t *testing.T) {
	// Set luminance to 75%.
	c := ApplyLuminance(Red, 75000)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.L-0.75) > 0.02 {
		t.Errorf(
			"ApplyLuminance(Red, 75%%): expected L~0.75, got %.2f",
			hsl.L,
		)
	}
}

func TestApplySaturationOffset(t *testing.T) {
	// Create a color with 50% saturation.
	base := NewHSL(0, 0.5, 0.5).ToRGB()

	// Add 25% saturation.
	c := ApplySaturationOffset(base, 25000)
	hsl := RGBToHSL(c)

	if math.Abs(hsl.S-0.75) > 0.02 {
		t.Errorf(
			"ApplySaturationOffset(0.5, +25%%): expected S~0.75, got %.2f",
			hsl.S,
		)
	}
}

func TestHSVNormalization(t *testing.T) {
	// Test hue normalization.
	hsv := NewHSV(-90, 1, 1)
	if hsv.H < 0 || hsv.H >= 360 {
		t.Errorf(
			"HSV should normalize negative hue, got %.2f",
			hsv.H,
		)
	}

	hsv = NewHSV(450, 1, 1)
	if hsv.H < 0 || hsv.H >= 360 {
		t.Errorf(
			"HSV should normalize hue > 360, got %.2f",
			hsv.H,
		)
	}
}

func TestThemePalette_Nil(t *testing.T) {
	var palette *ThemePalette

	// GetColor on nil palette should return false.
	_, ok := palette.GetColor(
		drawingml.SchemeColorAccent1,
	)
	if ok {
		t.Error(
			"GetColor on nil palette should return false",
		)
	}

	// SetColor on nil palette should not panic.
	palette.SetColor(
		drawingml.SchemeColorAccent1,
		Red,
	) // Should not panic

	// Clone of nil should return nil.
	clone := palette.Clone()
	if clone != nil {
		t.Error(
			"Clone of nil palette should return nil",
		)
	}
}

func TestSystemColorsMap(t *testing.T) {
	// Verify some system colors are defined.
	expectedColors := []string{
		"window", "windowText", "btnFace", "highlight", "highlightText",
	}

	for _, name := range expectedColors {
		if _, ok := SystemColors[name]; !ok {
			t.Errorf(
				"SystemColors[%q]: expected to find color",
				name,
			)
		}
	}
}

func TestPresetColorsMap(t *testing.T) {
	// Verify some preset colors are defined.
	expectedColors := []string{
		"black", "white", "red", "blue", "coral", "dkBlue", "ltGray",
	}

	for _, name := range expectedColors {
		if _, ok := PresetColors[name]; !ok {
			t.Errorf(
				"PresetColors[%q]: expected to find color",
				name,
			)
		}
	}
}
