package layout

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// mockFont creates a minimal font for testing glyph metrics.
func mockFont() *font.Font {
	return &font.Font{
		Family: "TestFont",
		Style:  font.StyleRegular,
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
			Ascender:   800,
			Descender:  -200,
			LineGap:    90,
			CapHeight:  700,
			XHeight:    500,
		},
		GlyphData: map[rune]font.GlyphMetrics{
			'H': {
				AdvanceWidth: 700,
				LeftBearing:  50,
			},
			'e': {
				AdvanceWidth: 500,
				LeftBearing:  30,
			},
			'l': {
				AdvanceWidth: 250,
				LeftBearing:  40,
			},
			'o': {
				AdvanceWidth: 550,
				LeftBearing:  25,
			},
			' ': {
				AdvanceWidth: 250,
				LeftBearing:  0,
			},
			'W': {
				AdvanceWidth: 900,
				LeftBearing:  20,
			},
			'r': {
				AdvanceWidth: 350,
				LeftBearing:  35,
			},
			'd': {
				AdvanceWidth: 550,
				LeftBearing:  28,
			},
			'!': {
				AdvanceWidth: 250,
				LeftBearing:  80,
			},
			'x': {
				AdvanceWidth: 480,
				LeftBearing:  15,
			},
			// Add some special characters
			'\u00A0': {
				AdvanceWidth: 250,
				LeftBearing:  0,
			}, // NBSP
			'\u2014': {
				AdvanceWidth: 800,
				LeftBearing:  0,
			}, // Em dash
			'\u4E2D': {
				AdvanceWidth: 1000,
				LeftBearing:  0,
			}, // CJK character
			'\uFFFD': {
				AdvanceWidth: 600,
				LeftBearing:  0,
			}, // Replacement character
		},
	}
}

func TestNewGlyphMetrics(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	if gm == nil {
		t.Fatal(
			"NewGlyphMetrics returned nil for valid font",
		)
	}

	if gm.font != f {
		t.Error(
			"GlyphMetrics does not reference the correct font",
		)
	}

	if gm.unitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			gm.unitsPerEm,
		)
	}

	// Space width should be cached
	if gm.spaceWidth != 250 {
		t.Errorf(
			"SpaceWidth = %d, want 250",
			gm.spaceWidth,
		)
	}
}

func TestNewGlyphMetricsNilFont(t *testing.T) {
	gm := NewGlyphMetrics(nil)

	if gm != nil {
		t.Error(
			"NewGlyphMetrics should return nil for nil font",
		)
	}
}

func TestGetGlyphID(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		r        rune
		wantZero bool // true if we expect glyph ID 0 (.notdef)
	}{
		{'H', false},
		{'e', false},
		{' ', false},
		{'z', true}, // Not in mock font
		{'\u4E2D', false},
		{'\u1234', true}, // Not in mock font
	}

	for _, tt := range tests {
		id := gm.GetGlyphID(tt.r)
		if tt.wantZero && id != 0 {
			t.Errorf(
				"GetGlyphID(%q) = %d, want 0",
				tt.r,
				id,
			)
		}
		if !tt.wantZero && id == 0 {
			t.Errorf(
				"GetGlyphID(%q) = 0, want non-zero",
				tt.r,
			)
		}
	}
}

func TestHasGlyph(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		r    rune
		want bool
	}{
		{'H', true},
		{'e', true},
		{' ', true},
		{'z', false},
		{'\u4E2D', true},
		{'\u1234', false},
	}

	for _, tt := range tests {
		if got := gm.HasGlyph(tt.r); got != tt.want {
			t.Errorf(
				"HasGlyph(%q) = %v, want %v",
				tt.r,
				got,
				tt.want,
			)
		}
	}
}

func TestGetAdvanceWidthForRune(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		r    rune
		want int
	}{
		{'H', 700},
		{'e', 500},
		{'l', 250},
		{'o', 550},
		{' ', 250},
		{'W', 900},
		{'\u4E2D', 1000},
	}

	for _, tt := range tests {
		if got := gm.GetAdvanceWidthForRune(tt.r); got != tt.want {
			t.Errorf(
				"GetAdvanceWidthForRune(%q) = %d, want %d",
				tt.r,
				got,
				tt.want,
			)
		}
	}
}

func TestGetAdvanceWidthForRuneMissing(
	t *testing.T,
) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// Missing glyph should return default width
	width := gm.GetAdvanceWidthForRune('z')

	// Default should be based on replacement character (600) or fallback calculation
	// In our mock font, FFFD has width 600
	if width != 600 {
		t.Errorf(
			"GetAdvanceWidthForRune('z') = %d, want default (600)",
			width,
		)
	}
}

func TestGetLeftBearing(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		r    rune
		want int
	}{
		{'H', 50},
		{'e', 30},
		{' ', 0},
		{'z', 0}, // Missing glyph should return 0
	}

	for _, tt := range tests {
		if got := gm.GetLeftBearing(tt.r); got != tt.want {
			t.Errorf(
				"GetLeftBearing(%q) = %d, want %d",
				tt.r,
				got,
				tt.want,
			)
		}
	}
}

func TestMeasureString(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// Correct calculations:
	// H=700, e=500, l=250, l=250, o=550 => Hello = 2250
	// W=900, o=550, r=350, l=250, d=550 => World = 2600
	// Hello World = 2250 + 250 + 2600 = 5100
	tests := []struct {
		text string
		want int
	}{
		{"", 0},
		{"H", 700},
		{"He", 1200}, // 700 + 500
		{
			"Hello",
			2250,
		}, // 700 + 500 + 250 + 250 + 550
		{
			"Hello World",
			5100,
		}, // Hello (2250) + space (250) + World (2600)
	}

	for _, tt := range tests {
		if got := gm.MeasureString(tt.text); got != tt.want {
			t.Errorf(
				"MeasureString(%q) = %d, want %d",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func TestMeasureStringScaled(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// "Hello" has width 2250 in font units with UnitsPerEm=1000
	// At 12pt: 2250 * 12 / 1000 = 27.0
	got := gm.MeasureStringScaled("Hello", 12.0)
	want := 27.0

	if got != want {
		t.Errorf(
			"MeasureStringScaled(\"Hello\", 12) = %v, want %v",
			got,
			want,
		)
	}

	// At 72pt (1 inch): 2250 * 72 / 1000 = 162.0
	got = gm.MeasureStringScaled("Hello", 72.0)
	want = 162.0

	if got != want {
		t.Errorf(
			"MeasureStringScaled(\"Hello\", 72) = %v, want %v",
			got,
			want,
		)
	}
}

func TestMeasureRuneScaled(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// 'H' has width 700 in font units with UnitsPerEm=1000
	// At 10pt: 700 * 10 / 1000 = 7.0
	got := gm.MeasureRuneScaled('H', 10.0)
	want := 7.0

	if got != want {
		t.Errorf(
			"MeasureRuneScaled('H', 10) = %v, want %v",
			got,
			want,
		)
	}
}

func TestFontMetricAccessors(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	if got := gm.UnitsPerEm(); got != 1000 {
		t.Errorf(
			"UnitsPerEm() = %d, want 1000",
			got,
		)
	}

	if got := gm.Ascender(); got != 800 {
		t.Errorf("Ascender() = %d, want 800", got)
	}

	if got := gm.Descender(); got != -200 {
		t.Errorf(
			"Descender() = %d, want -200",
			got,
		)
	}

	if got := gm.LineGap(); got != 90 {
		t.Errorf("LineGap() = %d, want 90", got)
	}

	// LineHeight = Ascender - Descender + LineGap = 800 - (-200) + 90 = 1090
	if got := gm.LineHeight(); got != 1090 {
		t.Errorf(
			"LineHeight() = %d, want 1090",
			got,
		)
	}
}

func TestLineHeightScaled(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// LineHeight = 1090, UnitsPerEm = 1000
	// At 12pt: 1090 * 12 / 1000 = 13.08
	got := gm.LineHeightScaled(12.0)
	want := 13.08

	if got != want {
		t.Errorf(
			"LineHeightScaled(12) = %v, want %v",
			got,
			want,
		)
	}
}

func TestSpaceWidth(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	if got := gm.SpaceWidth(); got != 250 {
		t.Errorf(
			"SpaceWidth() = %d, want 250",
			got,
		)
	}
}

func TestDefaultWidth(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// Default width should be the width of FFFD (600) in our mock font
	if got := gm.DefaultWidth(); got != 600 {
		t.Errorf(
			"DefaultWidth() = %d, want 600",
			got,
		)
	}
}

func TestDefaultWidthFallback(t *testing.T) {
	// Create a font without FFFD but with 'x'
	f := &font.Font{
		Family: "TestFont",
		Style:  font.StyleRegular,
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
		},
		GlyphData: map[rune]font.GlyphMetrics{
			'x': {
				AdvanceWidth: 480,
				LeftBearing:  15,
			},
		},
	}

	gm := NewGlyphMetrics(f)

	// Should fall back to 'x' width
	if got := gm.DefaultWidth(); got != 480 {
		t.Errorf(
			"DefaultWidth() = %d, want 480 (fallback to 'x')",
			got,
		)
	}
}

func TestTextMeasurer(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	tm := NewTextMeasurer(gm, 12.0)

	if tm.FontSize() != 12.0 {
		t.Errorf(
			"FontSize() = %v, want 12.0",
			tm.FontSize(),
		)
	}

	if tm.Metrics() != gm {
		t.Error(
			"Metrics() should return the underlying GlyphMetrics",
		)
	}

	// Test measurement
	// "Hello" = 2250 font units, at 12pt with 1000 upem = 27.0
	got := tm.Measure("Hello")
	want := 27.0

	if got != want {
		t.Errorf(
			"Measure(\"Hello\") = %v, want %v",
			got,
			want,
		)
	}
}

func TestTextMeasurerSetFontSize(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	tm := NewTextMeasurer(gm, 12.0)

	tm.SetFontSize(24.0)

	if tm.FontSize() != 24.0 {
		t.Errorf(
			"FontSize() after SetFontSize(24) = %v, want 24.0",
			tm.FontSize(),
		)
	}

	// "Hello" = 2250 font units, at 24pt with 1000 upem = 54.0
	got := tm.Measure("Hello")
	want := 54.0

	if got != want {
		t.Errorf(
			"Measure(\"Hello\") at 24pt = %v, want %v",
			got,
			want,
		)
	}
}

func TestMeasureFunc(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	tm := NewTextMeasurer(gm, 12.0)

	measureFunc := tm.MeasureFunc()

	// Verify it works as a MeasureFunc
	// "Hello" = 2250 font units, at 12pt with 1000 upem = 27.0
	got := measureFunc("Hello")
	want := 27.0

	if got != want {
		t.Errorf(
			"MeasureFunc()(\"Hello\") = %v, want %v",
			got,
			want,
		)
	}
}

func TestFindMissingGlyphs(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// Text with some missing glyphs: 'z', 'y' are not in mock font
	text := "Hello zy!"
	missing := gm.FindMissingGlyphs(text)

	if len(missing) != 2 {
		t.Errorf(
			"FindMissingGlyphs found %d missing, want 2",
			len(missing),
		)
	}

	// Check first missing glyph 'z'
	if len(missing) > 0 {
		if missing[0].Rune != 'z' {
			t.Errorf(
				"First missing glyph = %q, want 'z'",
				missing[0].Rune,
			)
		}
		if missing[0].RunePosition != 6 {
			t.Errorf(
				"First missing rune position = %d, want 6",
				missing[0].RunePosition,
			)
		}
	}

	// Check second missing glyph 'y'
	if len(missing) > 1 {
		if missing[1].Rune != 'y' {
			t.Errorf(
				"Second missing glyph = %q, want 'y'",
				missing[1].Rune,
			)
		}
	}
}

func TestFindMissingGlyphsEmpty(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// Text with no missing glyphs
	text := "Hello World!"
	missing := gm.FindMissingGlyphs(text)

	if len(missing) != 0 {
		t.Errorf(
			"FindMissingGlyphs found %d missing, want 0",
			len(missing),
		)
	}
}

func TestCoverage(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		text string
		want float64
	}{
		{
			"",
			1.0,
		}, // Empty string = 100% coverage
		{"Hello", 1.0}, // All glyphs present
		{"z", 0.0},     // No glyphs present
	}

	for _, tt := range tests {
		got := gm.Coverage(tt.text)
		if got != tt.want {
			t.Errorf(
				"Coverage(%q) = %v, want %v",
				tt.text,
				got,
				tt.want,
			)
		}
	}

	// Test "Hello z" separately since 6/7 = 0.857142857... (floating point)
	// "Hello z" = 7 chars: H, e, l, l, o, space, z - only z is missing = 6/7
	got := gm.Coverage("Hello z")
	want := 6.0 / 7.0
	if got != want {
		t.Errorf(
			"Coverage(\"Hello z\") = %v, want %v",
			got,
			want,
		)
	}
}

func TestCoveragePartial(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	// "Hz" - 'H' is present, 'z' is missing = 50%
	got := gm.Coverage("Hz")
	want := 0.5

	if got != want {
		t.Errorf(
			"Coverage(\"Hz\") = %v, want %v",
			got,
			want,
		)
	}
}

func TestFont(t *testing.T) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	if gm.Font() != f {
		t.Error(
			"Font() should return the underlying font",
		)
	}
}

func TestNilFontHandling(t *testing.T) {
	gm := &GlyphMetrics{
		font:         nil,
		defaultWidth: 500,
		unitsPerEm:   1000,
	}

	// All methods should handle nil font gracefully
	if gm.HasGlyph('H') {
		t.Error(
			"HasGlyph should return false for nil font",
		)
	}

	if gm.GetGlyphID('H') != 0 {
		t.Error(
			"GetGlyphID should return 0 for nil font",
		)
	}

	if gm.MeasureString("Hello") != 0 {
		t.Error(
			"MeasureString should return 0 for nil font",
		)
	}

	if gm.Ascender() != 0 {
		t.Error(
			"Ascender should return 0 for nil font",
		)
	}

	if gm.Descender() != 0 {
		t.Error(
			"Descender should return 0 for nil font",
		)
	}

	if gm.LineGap() != 0 {
		t.Error(
			"LineGap should return 0 for nil font",
		)
	}
}

func TestTextMeasurerNilMetrics(t *testing.T) {
	tm := NewTextMeasurer(nil, 12.0)

	if tm.Measure("Hello") != 0 {
		t.Error(
			"TextMeasurer.Measure should return 0 for nil metrics",
		)
	}
}

// Benchmark tests
func BenchmarkMeasureString(b *testing.B) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	text := "Hello World!"

	b.ResetTimer()
	for range b.N {
		gm.MeasureString(text)
	}
}

func BenchmarkMeasureStringScaled(b *testing.B) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	text := "Hello World!"

	b.ResetTimer()
	for range b.N {
		gm.MeasureStringScaled(text, 12.0)
	}
}

func BenchmarkGetAdvanceWidthForRune(
	b *testing.B,
) {
	f := mockFont()
	gm := NewGlyphMetrics(f)

	b.ResetTimer()
	for range b.N {
		gm.GetAdvanceWidthForRune('H')
	}
}

func BenchmarkFindMissingGlyphs(b *testing.B) {
	f := mockFont()
	gm := NewGlyphMetrics(f)
	text := "Hello World! This is a longer text with some missing chars like xyz."

	b.ResetTimer()
	for range b.N {
		gm.FindMissingGlyphs(text)
	}
}
