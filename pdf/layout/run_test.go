package layout

import (
	"math"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// mockFontForRun creates a font for testing run measurements.
func mockFontForRun() *font.Font {
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
			'A': {
				AdvanceWidth: 700,
				LeftBearing:  20,
			},
			'V': {
				AdvanceWidth: 650,
				LeftBearing:  10,
			},
			'T': {
				AdvanceWidth: 600,
				LeftBearing:  5,
			},
			'x': {
				AdvanceWidth: 480,
				LeftBearing:  15,
			},
			// CJK character for mixed script testing
			'\u4E2D': {
				AdvanceWidth: 1000,
				LeftBearing:  0,
			},
			// Various space characters
			'\u2003': {
				AdvanceWidth: 1000,
				LeftBearing:  0,
			}, // Em Space
			'\u2002': {
				AdvanceWidth: 500,
				LeftBearing:  0,
			}, // En Space
			// Replacement character for missing glyph fallback
			'\uFFFD': {
				AdvanceWidth: 600,
				LeftBearing:  0,
			},
		},
	}
}

func TestNewTextRun(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	run := NewTextRun("Hello", gm, 12.0)

	if run == nil {
		t.Fatal("NewTextRun returned nil")
	}

	if run.Text != "Hello" {
		t.Errorf(
			"Text = %q, want %q",
			run.Text,
			"Hello",
		)
	}

	if run.GlyphMetrics != gm {
		t.Error("GlyphMetrics not set correctly")
	}

	if run.FontSize != 12.0 {
		t.Errorf(
			"FontSize = %v, want 12.0",
			run.FontSize,
		)
	}

	if run.HorizontalScaling != 1.0 {
		t.Errorf(
			"HorizontalScaling = %v, want 1.0",
			run.HorizontalScaling,
		)
	}
}

func TestNewTextRunWithKerning(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	kt := NewKerningTable(f)

	run := NewTextRunWithKerning(
		"Hello",
		gm,
		kt,
		12.0,
	)

	if run == nil {
		t.Fatal(
			"NewTextRunWithKerning returned nil",
		)
	}

	if run.KerningTable != kt {
		t.Error("KerningTable not set correctly")
	}
}

func TestMeasureRunNil(t *testing.T) {
	// Nil run
	if MeasureRun(nil) != 0 {
		t.Error("MeasureRun(nil) should return 0")
	}

	// Empty text
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := NewTextRun("", gm, 12.0)
	if MeasureRun(run) != 0 {
		t.Error(
			"MeasureRun with empty text should return 0",
		)
	}

	// Nil GlyphMetrics
	run = &TextRun{
		Text:              "Hello",
		GlyphMetrics:      nil,
		FontSize:          12.0,
		HorizontalScaling: 1.0,
	}
	if MeasureRun(run) != 0 {
		t.Error(
			"MeasureRun with nil GlyphMetrics should return 0",
		)
	}
}

func TestMeasureRunBasic(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello" = H(700) + e(500) + l(250) + l(250) + o(550) = 2250 font units
	// At 12pt with 1000 upem: 2250 * 12 / 1000 = 27.0
	run := NewTextRun("Hello", gm, 12.0)
	width := MeasureRun(run)
	expected := 27.0

	if width != expected {
		t.Errorf(
			"MeasureRun(\"Hello\") = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunSingleChar(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "H" = 700 font units
	// At 12pt with 1000 upem: 700 * 12 / 1000 = 8.4
	run := NewTextRun("H", gm, 12.0)
	width := MeasureRun(run)
	expected := 8.4

	if width != expected {
		t.Errorf(
			"MeasureRun(\"H\") = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithSpaces(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello World" = Hello(2250) + space(250) + World(2600) = 5100 font units
	// World = W(900) + o(550) + r(350) + l(250) + d(550) = 2600
	// At 12pt with 1000 upem: 5100 * 12 / 1000 = 61.2
	run := NewTextRun("Hello World", gm, 12.0)
	width := MeasureRun(run)
	expected := 61.2

	if width != expected {
		t.Errorf(
			"MeasureRun(\"Hello World\") = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithCharacterSpacing(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello" = 2250 font units = 27.0 points at 12pt
	// Character spacing of 1.0 pt between each character (4 gaps for 5 chars)
	// Total: 27.0 + 4*1.0 = 31.0
	run := &TextRun{
		Text:              "Hello",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  1.0,
		HorizontalScaling: 1.0,
	}
	width := MeasureRun(run)
	expected := 31.0

	if width != expected {
		t.Errorf(
			"MeasureRun with CharacterSpacing = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithNegativeCharacterSpacing(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello" = 27.0 points base
	// Character spacing of -0.5 pt between each character
	// Total: 27.0 + 4*(-0.5) = 27.0 - 2.0 = 25.0
	run := &TextRun{
		Text:              "Hello",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  -0.5,
		HorizontalScaling: 1.0,
	}
	width := MeasureRun(run)
	expected := 25.0

	if width != expected {
		t.Errorf(
			"MeasureRun with negative CharacterSpacing = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithWordSpacing(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello World" = 61.2 points base
	// Word spacing of 2.0 pt for the 1 space
	// Total: 61.2 + 2.0 = 63.2
	run := &TextRun{
		Text:              "Hello World",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		WordSpacing:       2.0,
		HorizontalScaling: 1.0,
	}
	width := MeasureRun(run)
	expected := 63.2

	if width != expected {
		t.Errorf(
			"MeasureRun with WordSpacing = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithMultipleSpaces(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "H H H" has 2 spaces
	// H(700) + space(250) + H(700) + space(250) + H(700) = 2600 font units
	// At 12pt: 2600 * 12 / 1000 = 31.2
	// Word spacing of 3.0 pt for 2 spaces: 31.2 + 2*3.0 = 37.2
	run := &TextRun{
		Text:              "H H H",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		WordSpacing:       3.0,
		HorizontalScaling: 1.0,
	}
	width := MeasureRun(run)
	expected := 37.2

	if width != expected {
		t.Errorf(
			"MeasureRun with multiple spaces = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunWithHorizontalScaling(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "Hello" = 27.0 points at normal width
	// At 50% scaling: 27.0 * 0.5 = 13.5
	run := &TextRun{
		Text:              "Hello",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		HorizontalScaling: 0.5,
	}
	width := MeasureRun(run)
	expected := 13.5

	if width != expected {
		t.Errorf(
			"MeasureRun at 50%% scaling = %v, want %v",
			width,
			expected,
		)
	}

	// At 200% scaling: 27.0 * 2.0 = 54.0
	run.HorizontalScaling = 2.0
	width = MeasureRun(run)
	expected = 54.0

	if width != expected {
		t.Errorf(
			"MeasureRun at 200%% scaling = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunZeroScalingUsesDefault(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// Zero scaling should be treated as 1.0
	run := &TextRun{
		Text:              "Hello",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		HorizontalScaling: 0, // Should default to 1.0
	}
	width := MeasureRun(run)
	expected := 27.0

	if width != expected {
		t.Errorf(
			"MeasureRun with zero scaling = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunCombinedOptions(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "H H" at 12pt
	// Base: H(700) + space(250) + H(700) = 1650 font units = 19.8 pt
	// Scaling 1.5: 19.8 * 1.5 = 29.7
	// Character spacing 0.5 (2 gaps): 29.7 + 2*0.5 = 30.7
	// Word spacing 1.0 (1 space): 30.7 + 1.0 = 31.7
	run := &TextRun{
		Text:              "H H",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  0.5,
		WordSpacing:       1.0,
		HorizontalScaling: 1.5,
	}
	width := MeasureRun(run)
	expected := 31.7

	if !floatEquals(width, expected, 0.001) {
		t.Errorf(
			"MeasureRun with combined options = %v, want %v",
			width,
			expected,
		)
	}
}

func TestTextRunWidth(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	run := NewTextRun("Hello", gm, 12.0)

	// Width() should call MeasureRun
	if run.Width() != MeasureRun(run) {
		t.Error(
			"TextRun.Width() should equal MeasureRun(run)",
		)
	}
}

func TestMeasureRunWithOptions(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	run := &TextRun{
		Text:              "H H",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  1.0,
		WordSpacing:       2.0,
		HorizontalScaling: 1.5,
	}

	// With all options disabled, should get base width only
	opts := RunMeasureOptions{
		IncludeKerning:     false,
		IncludeCharSpacing: false,
		IncludeWordSpacing: false,
		IncludeScaling:     false,
	}

	// Base: 1650 font units = 19.8 pt (no scaling, no spacing)
	width := MeasureRunWithOptions(run, opts)
	expected := 19.8

	if width != expected {
		t.Errorf(
			"MeasureRunWithOptions (all off) = %v, want %v",
			width,
			expected,
		)
	}

	// With only scaling
	opts.IncludeScaling = true
	// 19.8 * 1.5 = 29.7
	width = MeasureRunWithOptions(run, opts)
	expected = 29.7

	if !floatEquals(width, expected, 0.001) {
		t.Errorf(
			"MeasureRunWithOptions (scaling only) = %v, want %v",
			width,
			expected,
		)
	}

	// With scaling and char spacing
	opts.IncludeCharSpacing = true
	// 29.7 + 2*1.0 = 31.7
	width = MeasureRunWithOptions(run, opts)
	expected = 31.7

	if !floatEquals(width, expected, 0.001) {
		t.Errorf(
			"MeasureRunWithOptions (scaling + char) = %v, want %v",
			width,
			expected,
		)
	}
}

func TestDefaultRunMeasureOptions(t *testing.T) {
	opts := DefaultRunMeasureOptions()

	if !opts.IncludeKerning {
		t.Error("Default should include kerning")
	}
	if !opts.IncludeCharSpacing {
		t.Error(
			"Default should include char spacing",
		)
	}
	if !opts.IncludeWordSpacing {
		t.Error(
			"Default should include word spacing",
		)
	}
	if !opts.IncludeScaling {
		t.Error("Default should include scaling")
	}
}

func TestCountSpaces(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"", 0},
		{"Hello", 0},
		{" ", 1},
		{"Hello World", 1},
		{"a b c d", 3},
		{"  ", 2},
		{"\u2003", 1}, // Em Space
	}

	for _, tt := range tests {
		got := countSpaces(tt.text)
		if got != tt.want {
			t.Errorf(
				"countSpaces(%q) = %d, want %d",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func TestIsSpaceChar(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{' ', true},
		{'a', false},
		{'\u2003', true}, // Em Space
		{'\u2002', true}, // En Space
		{
			'\t',
			false,
		}, // Tab is not a "word" space
		{
			'\n',
			false,
		}, // Newline is not a "word" space
		{
			'\u00A0',
			false,
		}, // NBSP should not get word spacing
	}

	for _, tt := range tests {
		got := isSpaceChar(tt.r)
		if got != tt.want {
			t.Errorf(
				"isSpaceChar(%q) = %v, want %v",
				tt.r,
				got,
				tt.want,
			)
		}
	}
}

func TestSegmentRunForFallback(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// Text with some missing glyphs: 'z', 'y' are not in mock font
	run := &TextRun{
		Text:         "Hello zy World",
		GlyphMetrics: gm,
		FontSize:     12.0,
	}

	segments := SegmentRunForFallback(run)

	if len(segments) == 0 {
		t.Fatal(
			"SegmentRunForFallback returned no segments",
		)
	}

	// Verify we have the expected segments
	// "Hello " - covered
	// "zy" - needs fallback
	// " World" - covered
	if len(segments) != 3 {
		t.Errorf(
			"Expected 3 segments, got %d",
			len(segments),
		)
		for i, seg := range segments {
			t.Logf(
				"Segment %d: %q, NeedsFallback=%v",
				i,
				seg.Text,
				seg.NeedsFallback,
			)
		}

		return
	}

	if segments[0].NeedsFallback {
		t.Errorf(
			"First segment %q should not need fallback",
			segments[0].Text,
		)
	}

	if !segments[1].NeedsFallback {
		t.Errorf(
			"Second segment %q should need fallback",
			segments[1].Text,
		)
	}

	if segments[1].Text != "zy" {
		t.Errorf(
			"Second segment text = %q, want %q",
			segments[1].Text,
			"zy",
		)
	}

	if len(segments[1].MissingRunes) != 2 {
		t.Errorf(
			"Expected 2 missing runes, got %d",
			len(segments[1].MissingRunes),
		)
	}

	if segments[2].NeedsFallback {
		t.Errorf(
			"Third segment %q should not need fallback",
			segments[2].Text,
		)
	}
}

func TestSegmentRunForFallbackAllCovered(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	run := &TextRun{
		Text:         "Hello",
		GlyphMetrics: gm,
		FontSize:     12.0,
	}

	segments := SegmentRunForFallback(run)

	if len(segments) != 1 {
		t.Errorf(
			"Expected 1 segment, got %d",
			len(segments),
		)

		return
	}

	if segments[0].NeedsFallback {
		t.Error(
			"Segment should not need fallback",
		)
	}

	if segments[0].Text != "Hello" {
		t.Errorf(
			"Segment text = %q, want %q",
			segments[0].Text,
			"Hello",
		)
	}
}

func TestSegmentRunForFallbackNil(t *testing.T) {
	segments := SegmentRunForFallback(nil)
	if segments != nil {
		t.Error(
			"SegmentRunForFallback(nil) should return nil",
		)
	}

	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// Empty text
	run := &TextRun{
		Text:         "",
		GlyphMetrics: gm,
		FontSize:     12.0,
	}
	segments = SegmentRunForFallback(run)
	if segments != nil {
		t.Error(
			"SegmentRunForFallback with empty text should return nil",
		)
	}
}

func TestRunMeasurer(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	kt := NewKerningTable(f)

	rm := NewRunMeasurer(gm, kt, 12.0)

	if rm.FontSize() != 12.0 {
		t.Errorf(
			"FontSize() = %v, want 12.0",
			rm.FontSize(),
		)
	}

	// "Hello" = 27.0 at 12pt
	width := rm.Measure("Hello")
	if width != 27.0 {
		t.Errorf(
			"Measure(\"Hello\") = %v, want 27.0",
			width,
		)
	}

	// Test SetFontSize
	rm.SetFontSize(24.0)
	if rm.FontSize() != 24.0 {
		t.Errorf(
			"FontSize() after SetFontSize = %v, want 24.0",
			rm.FontSize(),
		)
	}

	// "Hello" = 54.0 at 24pt
	width = rm.Measure("Hello")
	if width != 54.0 {
		t.Errorf(
			"Measure(\"Hello\") at 24pt = %v, want 54.0",
			width,
		)
	}
}

func TestRunMeasurerWithSpacing(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	rm := NewRunMeasurer(gm, nil, 12.0)

	// "H H" at 12pt with spacing
	// Base: 19.8 + charSpacing 0.5*2 + wordSpacing 1.0*1 = 21.8
	width := rm.MeasureWithSpacing(
		"H H",
		0.5,
		1.0,
	)
	expected := 21.8

	if width != expected {
		t.Errorf(
			"MeasureWithSpacing = %v, want %v",
			width,
			expected,
		)
	}
}

func TestRunMeasurerMeasureFunc(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	rm := NewRunMeasurer(gm, nil, 12.0)
	mf := rm.MeasureFunc()

	// Should work as a MeasureFunc
	width := mf("Hello")
	if width != 27.0 {
		t.Errorf(
			"MeasureFunc(\"Hello\") = %v, want 27.0",
			width,
		)
	}
}

func TestRunMeasurerNil(t *testing.T) {
	var rm *RunMeasurer

	if rm.Measure("Hello") != 0 {
		t.Error(
			"nil RunMeasurer.Measure should return 0",
		)
	}

	if rm.MeasureWithSpacing(
		"Hello",
		1.0,
		1.0,
	) != 0 {
		t.Error(
			"nil RunMeasurer.MeasureWithSpacing should return 0",
		)
	}
}

func TestIsWhitespaceRun(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	tests := []struct {
		text string
		want bool
	}{
		{"", true},
		{" ", true},
		{"  ", true},
		{" \t\n", true},
		{"Hello", false},
		{" Hello ", false},
	}

	for _, tt := range tests {
		run := NewTextRun(tt.text, gm, 12.0)
		got := IsWhitespaceRun(run)
		if got != tt.want {
			t.Errorf(
				"IsWhitespaceRun(%q) = %v, want %v",
				tt.text,
				got,
				tt.want,
			)
		}
	}

	// Nil run should be considered whitespace
	if !IsWhitespaceRun(nil) {
		t.Error(
			"IsWhitespaceRun(nil) should return true",
		)
	}
}

func TestIsEmptyRun(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	if !IsEmptyRun(nil) {
		t.Error(
			"IsEmptyRun(nil) should return true",
		)
	}

	run := NewTextRun("", gm, 12.0)
	if !IsEmptyRun(run) {
		t.Error(
			"IsEmptyRun with empty text should return true",
		)
	}

	run = NewTextRun("Hello", gm, 12.0)
	if IsEmptyRun(run) {
		t.Error(
			"IsEmptyRun with text should return false",
		)
	}
}

func TestMeasureRunDetailed(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "H H" with various options
	run := &TextRun{
		Text:              "H H",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  0.5,
		WordSpacing:       1.0,
		HorizontalScaling: 1.5,
	}

	info := MeasureRunDetailed(run)

	// Base: 1650 font units = 19.8 pt
	if info.BaseWidth != 19.8 {
		t.Errorf(
			"BaseWidth = %v, want 19.8",
			info.BaseWidth,
		)
	}

	// No kerning in this font
	if info.KerningAdjustment != 0 {
		t.Errorf(
			"KerningAdjustment = %v, want 0",
			info.KerningAdjustment,
		)
	}

	// Char spacing: 0.5 * 2 = 1.0
	if info.CharSpacingTotal != 1.0 {
		t.Errorf(
			"CharSpacingTotal = %v, want 1.0",
			info.CharSpacingTotal,
		)
	}

	// Word spacing: 1.0 * 1 = 1.0
	if info.WordSpacingTotal != 1.0 {
		t.Errorf(
			"WordSpacingTotal = %v, want 1.0",
			info.WordSpacingTotal,
		)
	}

	// Scaled: 19.8 * 1.5 = 29.7
	if !floatEquals(
		info.ScaledWidth,
		29.7,
		0.001,
	) {
		t.Errorf(
			"ScaledWidth = %v, want 29.7",
			info.ScaledWidth,
		)
	}

	// Total: 29.7 + 1.0 + 1.0 = 31.7
	if !floatEquals(
		info.TotalWidth,
		31.7,
		0.001,
	) {
		t.Errorf(
			"TotalWidth = %v, want 31.7",
			info.TotalWidth,
		)
	}
}

func TestMeasureRunDetailedEmpty(t *testing.T) {
	info := MeasureRunDetailed(nil)

	if info.TotalWidth != 0 {
		t.Error(
			"MeasureRunDetailed(nil) should return zero info",
		)
	}

	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := NewTextRun("", gm, 12.0)

	info = MeasureRunDetailed(run)
	if info.TotalWidth != 0 {
		t.Error(
			"MeasureRunDetailed with empty text should return zero info",
		)
	}
}

func TestMeasureRunWithMixedScripts(
	t *testing.T,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// Mix of ASCII and CJK
	// "H" + CJK char + "e" = 700 + 1000 + 500 = 2200 font units
	// At 12pt: 2200 * 12 / 1000 = 26.4
	run := NewTextRun("H\u4E2De", gm, 12.0)
	width := MeasureRun(run)
	expected := 26.4

	if width != expected {
		t.Errorf(
			"MeasureRun with mixed scripts = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureRunOnlyWhitespace(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)

	// "   " = 3 spaces = 750 font units
	// At 12pt: 750 * 12 / 1000 = 9.0
	run := NewTextRun("   ", gm, 12.0)
	width := MeasureRun(run)
	expected := 9.0

	if width != expected {
		t.Errorf(
			"MeasureRun with only spaces = %v, want %v",
			width,
			expected,
		)
	}

	// With word spacing of 1.0 pt for 3 spaces: 9.0 + 3.0 = 12.0
	run.WordSpacing = 1.0
	run.HorizontalScaling = 1.0
	width = MeasureRun(run)
	expected = 12.0

	if width != expected {
		t.Errorf(
			"MeasureRun with only spaces + word spacing = %v, want %v",
			width,
			expected,
		)
	}
}

// floatEquals compares two float64 values with tolerance for floating point errors.
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestMeasureRunWithKerning(t *testing.T) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	kt := NewKerningTable(f)

	// Manually add kerning for A-V pair
	leftID := uint16((uint32('A') % 65534) + 1)
	rightID := uint16((uint32('V') % 65534) + 1)
	kt.pairs[uint32(leftID)<<16|uint32(rightID)] = -80
	kt.hasKerning = true

	// "AV" = A(700) + V(650) = 1350 font units
	// With kerning -80: 1350 - 80 = 1270 font units
	// At 12pt: 1270 * 12 / 1000 = 15.24
	run := NewTextRunWithKerning(
		"AV",
		gm,
		kt,
		12.0,
	)
	width := MeasureRun(run)
	expected := 15.24

	if !floatEquals(width, expected, 0.001) {
		t.Errorf(
			"MeasureRun with kerning = %v, want %v",
			width,
			expected,
		)
	}
}

// Benchmark tests

func BenchmarkMeasureRun(b *testing.B) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := NewTextRun("Hello World!", gm, 12.0)

	b.ResetTimer()
	for range b.N {
		MeasureRun(run)
	}
}

func BenchmarkMeasureRunWithSpacing(
	b *testing.B,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := &TextRun{
		Text:              "Hello World!",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  0.5,
		WordSpacing:       1.0,
		HorizontalScaling: 1.0,
	}

	b.ResetTimer()
	for range b.N {
		MeasureRun(run)
	}
}

func BenchmarkMeasureRunDetailed(b *testing.B) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := &TextRun{
		Text:              "Hello World!",
		GlyphMetrics:      gm,
		FontSize:          12.0,
		CharacterSpacing:  0.5,
		WordSpacing:       1.0,
		HorizontalScaling: 1.5,
	}

	b.ResetTimer()
	for range b.N {
		MeasureRunDetailed(run)
	}
}

func BenchmarkSegmentRunForFallback(
	b *testing.B,
) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	run := &TextRun{
		Text:         "Hello xyz World abc Test",
		GlyphMetrics: gm,
		FontSize:     12.0,
	}

	b.ResetTimer()
	for range b.N {
		SegmentRunForFallback(run)
	}
}

func BenchmarkRunMeasurerMeasure(b *testing.B) {
	f := mockFontForRun()
	gm := NewGlyphMetrics(f)
	rm := NewRunMeasurer(gm, nil, 12.0)
	text := "Hello World!"

	b.ResetTimer()
	for range b.N {
		rm.Measure(text)
	}
}
