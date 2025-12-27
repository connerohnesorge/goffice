package layout

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// mockFontForParagraph creates a font for testing paragraph measurements.
func mockFontForParagraph() *font.Font {
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
			'i': {
				AdvanceWidth: 220,
				LeftBearing:  40,
			},
			's': {
				AdvanceWidth: 450,
				LeftBearing:  30,
			},
			't': {
				AdvanceWidth: 280,
				LeftBearing:  25,
			},
			'a': {
				AdvanceWidth: 520,
				LeftBearing:  20,
			},
			'n': {
				AdvanceWidth: 530,
				LeftBearing:  28,
			},
			'g': {
				AdvanceWidth: 520,
				LeftBearing:  22,
			},
			'p': {
				AdvanceWidth: 540,
				LeftBearing:  25,
			},
			'c': {
				AdvanceWidth: 480,
				LeftBearing:  28,
			},
			'b': {
				AdvanceWidth: 540,
				LeftBearing:  30,
			},
			'u': {
				AdvanceWidth: 530,
				LeftBearing:  25,
			},
			'm': {
				AdvanceWidth: 820,
				LeftBearing:  30,
			},
			'\uFFFD': {
				AdvanceWidth: 600,
				LeftBearing:  0,
			},
		},
	}
}

func TestParagraphAlignment(t *testing.T) {
	tests := []struct {
		align ParagraphAlignment
		want  string
	}{
		{AlignLeft, "left"},
		{AlignCenter, "center"},
		{AlignRight, "right"},
		{AlignJustify, "justify"},
		{AlignDistribute, "distribute"},
		{
			ParagraphAlignment(99),
			"left",
		}, // Unknown defaults to left
	}

	for _, tt := range tests {
		got := tt.align.String()
		if got != tt.want {
			t.Errorf(
				"ParagraphAlignment(%d).String() = %q, want %q",
				tt.align,
				got,
				tt.want,
			)
		}
	}
}

func TestLineSpacingRule(t *testing.T) {
	tests := []struct {
		rule LineSpacingRule
		want string
	}{
		{LineSpacingAuto, "auto"},
		{LineSpacingExact, "exact"},
		{LineSpacingAtLeast, "atLeast"},
		{LineSpacingMultiple, "multiple"},
		{
			LineSpacingRule(99),
			"auto",
		}, // Unknown defaults to auto
	}

	for _, tt := range tests {
		got := tt.rule.String()
		if got != tt.want {
			t.Errorf(
				"LineSpacingRule(%d).String() = %q, want %q",
				tt.rule,
				got,
				tt.want,
			)
		}
	}
}

func TestDefaultLineSpacing(t *testing.T) {
	ls := DefaultLineSpacing()

	if ls.Rule != LineSpacingAuto {
		t.Errorf(
			"Default Rule = %v, want LineSpacingAuto",
			ls.Rule,
		)
	}
	if ls.Before != 0 {
		t.Errorf(
			"Default Before = %v, want 0",
			ls.Before,
		)
	}
	if ls.After != 0 {
		t.Errorf(
			"Default After = %v, want 0",
			ls.After,
		)
	}
}

func TestParagraphIndentation_EffectiveIndents(
	t *testing.T,
) {
	tests := []struct {
		name           string
		indent         ParagraphIndentation
		wantFirst      float64
		wantSubsequent float64
	}{
		{
			name:           "No indentation",
			indent:         ParagraphIndentation{},
			wantFirst:      0,
			wantSubsequent: 0,
		},
		{
			name: "Left indent only",
			indent: ParagraphIndentation{
				Left: 36,
			},
			wantFirst:      36,
			wantSubsequent: 36,
		},
		{
			name: "First line indent",
			indent: ParagraphIndentation{
				Left:      36,
				FirstLine: 18,
			},
			wantFirst:      54, // 36 + 18
			wantSubsequent: 36,
		},
		{
			name: "Hanging indent",
			indent: ParagraphIndentation{
				Left:    36,
				Hanging: 18,
			},
			wantFirst:      36, // First line is at Left
			wantSubsequent: 54, // Subsequent lines at Left + Hanging
		},
		{
			name: "Right indent (doesn't affect left)",
			indent: ParagraphIndentation{
				Left:  36,
				Right: 24,
			},
			wantFirst:      36,
			wantSubsequent: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst := tt.indent.EffectiveFirstLineIndent()
			if gotFirst != tt.wantFirst {
				t.Errorf(
					"EffectiveFirstLineIndent() = %v, want %v",
					gotFirst,
					tt.wantFirst,
				)
			}

			gotSubsequent := tt.indent.EffectiveSubsequentIndent()
			if gotSubsequent != tt.wantSubsequent {
				t.Errorf(
					"EffectiveSubsequentIndent() = %v, want %v",
					gotSubsequent,
					tt.wantSubsequent,
				)
			}
		})
	}
}

func TestDefaultParagraphProperties(
	t *testing.T,
) {
	props := DefaultParagraphProperties()

	if props.Alignment != AlignLeft {
		t.Errorf(
			"Default Alignment = %v, want AlignLeft",
			props.Alignment,
		)
	}
	if !props.WidowControl {
		t.Error(
			"Default WidowControl should be true",
		)
	}
	if props.KeepLines {
		t.Error(
			"Default KeepLines should be false",
		)
	}
	if props.KeepWithNext {
		t.Error(
			"Default KeepWithNext should be false",
		)
	}
	if props.PageBreakBefore {
		t.Error(
			"Default PageBreakBefore should be false",
		)
	}
}

func TestNewParagraph(t *testing.T) {
	p := NewParagraph()

	if p == nil {
		t.Fatal("NewParagraph returned nil")
	}
	if len(p.Runs) != 0 {
		t.Errorf(
			"New paragraph should have no runs, got %d",
			len(p.Runs),
		)
	}
	if p.Properties.Alignment != AlignLeft {
		t.Errorf(
			"New paragraph alignment = %v, want AlignLeft",
			p.Properties.Alignment,
		)
	}
}

func TestNewParagraphWithRuns(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	runs := []*TextRun{
		NewTextRun("Hello ", gm, 12.0),
		NewTextRun("World", gm, 12.0),
	}

	p := NewParagraphWithRuns(runs)

	if p == nil {
		t.Fatal(
			"NewParagraphWithRuns returned nil",
		)
	}
	if len(p.Runs) != 2 {
		t.Errorf(
			"Paragraph should have 2 runs, got %d",
			len(p.Runs),
		)
	}
}

func TestParagraph_AddRun(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.AddRun(NewTextRun(" World", gm, 12.0))

	if len(p.Runs) != 2 {
		t.Errorf(
			"Expected 2 runs, got %d",
			len(p.Runs),
		)
	}

	// Adding nil should not crash or add
	p.AddRun(nil)
	if len(p.Runs) != 2 {
		t.Error(
			"Adding nil run should not increase run count",
		)
	}
}

func TestParagraph_Text(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.AddRun(NewTextRun(" ", gm, 12.0))
	p.AddRun(NewTextRun("World", gm, 12.0))

	text := p.Text()
	if text != "Hello World" {
		t.Errorf(
			"Text() = %q, want %q",
			text,
			"Hello World",
		)
	}
}

func TestParagraph_IsEmpty(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	// Nil paragraph
	var nilP *Paragraph
	if !nilP.IsEmpty() {
		t.Error("nil Paragraph should be empty")
	}

	// Empty paragraph
	p := NewParagraph()
	if !p.IsEmpty() {
		t.Error("New paragraph should be empty")
	}

	// Paragraph with empty run
	p.AddRun(NewTextRun("", gm, 12.0))
	if !p.IsEmpty() {
		t.Error(
			"Paragraph with only empty run should be empty",
		)
	}

	// Paragraph with content
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	if p.IsEmpty() {
		t.Error(
			"Paragraph with content should not be empty",
		)
	}
}

func TestParagraph_RunCount(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	var nilP *Paragraph
	if nilP.RunCount() != 0 {
		t.Error(
			"nil Paragraph should have 0 runs",
		)
	}

	p := NewParagraph()
	if p.RunCount() != 0 {
		t.Error(
			"New paragraph should have 0 runs",
		)
	}

	p.AddRun(NewTextRun("Hello", gm, 12.0))
	if p.RunCount() != 1 {
		t.Errorf(
			"RunCount() = %d, want 1",
			p.RunCount(),
		)
	}
}

func TestMeasureParagraph(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	// "Hello" = H(700) + e(500) + l(250) + l(250) + o(550) = 2250 font units
	// At 12pt with 1000 upem: 2250 * 12 / 1000 = 27.0
	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))

	width := MeasureParagraph(p)
	expected := 27.0

	if width != expected {
		t.Errorf(
			"MeasureParagraph = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureParagraphMultipleRuns(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	// "Hello" = 2250 font units = 27.0 pt
	// " " = 250 font units = 3.0 pt
	// "World" = W(900) + o(550) + r(350) + l(250) + d(550) = 2600 = 31.2 pt
	// Total = 27.0 + 3.0 + 31.2 = 61.2 pt

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.AddRun(NewTextRun(" ", gm, 12.0))
	p.AddRun(NewTextRun("World", gm, 12.0))

	width := MeasureParagraph(p)
	expected := 61.2

	if width != expected {
		t.Errorf(
			"MeasureParagraph(multiple runs) = %v, want %v",
			width,
			expected,
		)
	}
}

func TestMeasureParagraphNil(t *testing.T) {
	if MeasureParagraph(nil) != 0 {
		t.Error(
			"MeasureParagraph(nil) should return 0",
		)
	}

	p := NewParagraph()
	if MeasureParagraph(p) != 0 {
		t.Error(
			"MeasureParagraph(empty) should return 0",
		)
	}
}

func TestMeasureFirstLine(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	containerWidth := 500.0

	// No indentation
	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))

	available := MeasureFirstLine(
		p,
		containerWidth,
	)
	if available != containerWidth {
		t.Errorf(
			"MeasureFirstLine (no indent) = %v, want %v",
			available,
			containerWidth,
		)
	}

	// With left and right indentation
	p.Properties.Indentation.Left = 36
	p.Properties.Indentation.Right = 24
	available = MeasureFirstLine(
		p,
		containerWidth,
	)
	expected := containerWidth - 36 - 24 // 440
	if available != expected {
		t.Errorf(
			"MeasureFirstLine (with indent) = %v, want %v",
			available,
			expected,
		)
	}

	// With first-line indent
	p.Properties.Indentation.FirstLine = 18
	available = MeasureFirstLine(
		p,
		containerWidth,
	)
	expected = containerWidth - 36 - 18 - 24 // 422
	if available != expected {
		t.Errorf(
			"MeasureFirstLine (with first-line) = %v, want %v",
			available,
			expected,
		)
	}
}

func TestMeasureSubsequentLines(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	containerWidth := 500.0

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.Properties.Indentation.Left = 36
	p.Properties.Indentation.Right = 24
	p.Properties.Indentation.FirstLine = 18

	// Subsequent lines don't have first-line indent
	available := MeasureSubsequentLines(
		p,
		containerWidth,
	)
	expected := containerWidth - 36 - 24 // 440
	if available != expected {
		t.Errorf(
			"MeasureSubsequentLines = %v, want %v",
			available,
			expected,
		)
	}

	// With hanging indent
	p.Properties.Indentation.FirstLine = 0
	p.Properties.Indentation.Hanging = 18
	available = MeasureSubsequentLines(
		p,
		containerWidth,
	)
	expected = containerWidth - 36 - 18 - 24 // 422 (subsequent lines are more indented)
	if available != expected {
		t.Errorf(
			"MeasureSubsequentLines (hanging) = %v, want %v",
			available,
			expected,
		)
	}
}

func TestGetMinimumWidth(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	// "Hello World" - longest word is "World"
	p.AddRun(NewTextRun("Hello World", gm, 12.0))

	minWidth := GetMinimumWidth(p)

	// "World" = W(900) + o(550) + r(350) + l(250) + d(550) = 2600 font units = 31.2 pt
	expected := 31.2

	if minWidth != expected {
		t.Errorf(
			"GetMinimumWidth = %v, want %v",
			minWidth,
			expected,
		)
	}
}

func TestGetMinimumWidth_LongestWord(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	// The longest word should determine minimum width
	p.AddRun(NewTextRun("a test", gm, 12.0))

	minWidth := GetMinimumWidth(p)

	// "test" = t(280) + e(500) + s(450) + t(280) = 1510 font units = 18.12 pt
	expected := 18.12

	if !floatEquals(minWidth, expected, 0.01) {
		t.Errorf(
			"GetMinimumWidth = %v, want %v",
			minWidth,
			expected,
		)
	}
}

func TestGetMinimumWidth_WithIndentation(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.Properties.Indentation.Left = 36

	minWidth := GetMinimumWidth(p)

	// "Hello" = 27.0 pt + left indent 36
	expected := 27.0 + 36.0

	if minWidth != expected {
		t.Errorf(
			"GetMinimumWidth (with indent) = %v, want %v",
			minWidth,
			expected,
		)
	}
}

func TestGetMaximumWidth(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello World", gm, 12.0))

	maxWidth := GetMaximumWidth(p)

	// Full content width: "Hello World" = 61.2 pt
	expected := 61.2

	if maxWidth != expected {
		t.Errorf(
			"GetMaximumWidth = %v, want %v",
			maxWidth,
			expected,
		)
	}
}

func TestGetMaximumWidth_WithIndentation(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))
	p.Properties.Indentation.Left = 36
	p.Properties.Indentation.FirstLine = 18

	maxWidth := GetMaximumWidth(p)

	// "Hello" = 27.0 pt + left(36) + firstLine(18)
	expected := 27.0 + 36.0 + 18.0

	if maxWidth != expected {
		t.Errorf(
			"GetMaximumWidth (with indent) = %v, want %v",
			maxWidth,
			expected,
		)
	}
}

func TestSplitIntoWords(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{"", nil},
		{"Hello", []string{"Hello"}},
		{
			"Hello World",
			[]string{"Hello", "World"},
		},
		{
			"  Hello   World  ",
			[]string{"Hello", "World"},
		},
		{"a b c", []string{"a", "b", "c"}},
		{"   ", nil},
		{"\t\n", nil},
	}

	for _, tt := range tests {
		got := splitIntoWords(tt.text)
		if !stringSliceEqual(got, tt.want) {
			t.Errorf(
				"splitIntoWords(%q) = %v, want %v",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestParagraphMeasurer(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	pm := NewParagraphMeasurer(500.0)

	if pm.ContainerWidth() != 500.0 {
		t.Errorf(
			"ContainerWidth() = %v, want 500.0",
			pm.ContainerWidth(),
		)
	}

	pm.SetContainerWidth(400.0)
	if pm.ContainerWidth() != 400.0 {
		t.Errorf(
			"ContainerWidth() after Set = %v, want 400.0",
			pm.ContainerWidth(),
		)
	}

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))

	// Test various measurements
	natural := pm.MeasureNaturalWidth(p)
	if natural != 27.0 {
		t.Errorf(
			"MeasureNaturalWidth = %v, want 27.0",
			natural,
		)
	}

	firstAvail := pm.MeasureFirstLineAvailable(p)
	if firstAvail != 400.0 {
		t.Errorf(
			"MeasureFirstLineAvailable = %v, want 400.0",
			firstAvail,
		)
	}
}

func TestParagraphMeasurer_FitsWidth(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	pm := NewParagraphMeasurer(100.0)

	// "Hello" = 27.0 pt - should fit
	p := NewParagraph()
	p.AddRun(NewTextRun("Hello", gm, 12.0))

	if !pm.FitsWidth(p) {
		t.Error(
			"27pt paragraph should fit in 100pt container",
		)
	}

	if pm.RequiresWrapping(p) {
		t.Error(
			"27pt paragraph should not require wrapping in 100pt container",
		)
	}

	// Make container too small
	pm.SetContainerWidth(20.0)

	if pm.FitsWidth(p) {
		t.Error(
			"27pt paragraph should not fit in 20pt container",
		)
	}

	if !pm.RequiresWrapping(p) {
		t.Error(
			"27pt paragraph should require wrapping in 20pt container",
		)
	}
}

func TestMeasureParagraphDetailed(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello World", gm, 12.0))
	p.Properties.Indentation.Left = 36

	info := MeasureParagraphDetailed(p)

	if info.NaturalWidth != 61.2 {
		t.Errorf(
			"NaturalWidth = %v, want 61.2",
			info.NaturalWidth,
		)
	}

	if info.RunCount != 1 {
		t.Errorf(
			"RunCount = %d, want 1",
			info.RunCount,
		)
	}

	if info.CharacterCount != 11 {
		t.Errorf(
			"CharacterCount = %d, want 11",
			info.CharacterCount,
		)
	}

	if info.WordCount != 2 {
		t.Errorf(
			"WordCount = %d, want 2",
			info.WordCount,
		)
	}

	if info.FirstLineIndent != 36 {
		t.Errorf(
			"FirstLineIndent = %v, want 36",
			info.FirstLineIndent,
		)
	}
}

func TestCalculateTabPosition(t *testing.T) {
	p := NewParagraph()

	// No custom tabs - should use default (36pt intervals)
	pos, typ := CalculateTabPosition(p, 0)
	if pos != 36 {
		t.Errorf(
			"Default tab from 0 = %v, want 36",
			pos,
		)
	}
	if typ != TabStopLeft {
		t.Errorf(
			"Default tab type = %v, want TabStopLeft",
			typ,
		)
	}

	pos, _ = CalculateTabPosition(p, 35)
	if pos != 36 {
		t.Errorf(
			"Default tab from 35 = %v, want 36",
			pos,
		)
	}

	pos, _ = CalculateTabPosition(p, 36)
	if pos != 72 {
		t.Errorf(
			"Default tab from 36 = %v, want 72",
			pos,
		)
	}

	// With custom tabs
	p.Properties.TabStops = []TabStop{
		{Position: 72, Type: TabStopCenter},
		{Position: 144, Type: TabStopRight},
	}

	pos, typ = CalculateTabPosition(p, 0)
	if pos != 72 {
		t.Errorf(
			"Custom tab from 0 = %v, want 72",
			pos,
		)
	}
	if typ != TabStopCenter {
		t.Errorf(
			"Custom tab type = %v, want TabStopCenter",
			typ,
		)
	}

	pos, typ = CalculateTabPosition(p, 72)
	if pos != 144 {
		t.Errorf(
			"Custom tab from 72 = %v, want 144",
			pos,
		)
	}
	if typ != TabStopRight {
		t.Errorf(
			"Custom tab type = %v, want TabStopRight",
			typ,
		)
	}

	// Past all custom tabs - use default
	pos, _ = CalculateTabPosition(p, 200)
	if pos != 216 { // Next multiple of 36 after 200
		t.Errorf(
			"Default tab from 200 = %v, want 216",
			pos,
		)
	}
}

func TestEstimateLineCount(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	// Empty paragraph
	if EstimateLineCount(nil, 100) != 0 {
		t.Error(
			"nil paragraph should have 0 lines",
		)
	}

	// Fits on one line
	p := NewParagraph()
	p.AddRun(
		NewTextRun("Hello", gm, 12.0),
	) // 27pt

	count := EstimateLineCount(p, 100)
	if count != 1 {
		t.Errorf(
			"27pt in 100pt should be 1 line, got %d",
			count,
		)
	}

	// Needs multiple lines
	p = NewParagraph()
	p.AddRun(
		NewTextRun(
			"Hello World Hello World",
			gm,
			12.0,
		),
	) // Much wider

	count = EstimateLineCount(p, 50)
	if count < 2 {
		t.Errorf(
			"Wide paragraph in 50pt should need multiple lines, got %d",
			count,
		)
	}

	// Zero width
	count = EstimateLineCount(p, 0)
	if count != 0 {
		t.Errorf(
			"Zero width should give 0 lines, got %d",
			count,
		)
	}
}

func TestCalculateLineHeight(t *testing.T) {
	// Nil paragraph - default behavior
	height := CalculateLineHeight(nil, 12.0)
	expected := 12.0 * 1.2
	if !floatEquals(height, expected, 0.001) {
		t.Errorf(
			"nil paragraph line height = %v, want %v",
			height,
			expected,
		)
	}

	// Auto spacing
	p := NewParagraph()
	p.Properties.LineSpacing.Rule = LineSpacingAuto
	height = CalculateLineHeight(p, 12.0)
	expected = 12.0 * 1.15
	if !floatEquals(height, expected, 0.001) {
		t.Errorf(
			"Auto line height = %v, want %v",
			height,
			expected,
		)
	}

	// Exact spacing
	p.Properties.LineSpacing.Rule = LineSpacingExact
	p.Properties.LineSpacing.Value = 18.0
	height = CalculateLineHeight(p, 12.0)
	if !floatEquals(height, 18.0, 0.001) {
		t.Errorf(
			"Exact line height = %v, want 18.0",
			height,
		)
	}

	// At least - uses value when larger
	p.Properties.LineSpacing.Rule = LineSpacingAtLeast
	p.Properties.LineSpacing.Value = 20.0
	height = CalculateLineHeight(p, 12.0)
	if !floatEquals(height, 20.0, 0.001) {
		t.Errorf(
			"AtLeast line height (larger value) = %v, want 20.0",
			height,
		)
	}

	// At least - uses natural when larger
	p.Properties.LineSpacing.Value = 10.0
	height = CalculateLineHeight(p, 12.0)
	expected = 12.0 * 1.2 // Natural height
	if !floatEquals(height, expected, 0.001) {
		t.Errorf(
			"AtLeast line height (smaller value) = %v, want %v",
			height,
			expected,
		)
	}

	// Multiple
	p.Properties.LineSpacing.Rule = LineSpacingMultiple
	p.Properties.LineSpacing.Value = 2.0
	height = CalculateLineHeight(p, 12.0)
	if !floatEquals(height, 24.0, 0.001) {
		t.Errorf(
			"Multiple line height = %v, want 24.0",
			height,
		)
	}
}

func TestParagraphMeasurer_Nil(t *testing.T) {
	var pm *ParagraphMeasurer

	if pm.ContainerWidth() != 0 {
		t.Error(
			"nil ParagraphMeasurer.ContainerWidth() should return 0",
		)
	}

	if pm.MeasureFirstLineAvailable(nil) != 0 {
		t.Error(
			"nil ParagraphMeasurer methods should handle nil gracefully",
		)
	}

	// These should not panic
	pm.SetContainerWidth(100)
	pm.SetDefaultFontSize(12)
}

// Benchmark tests

func BenchmarkMeasureParagraph(b *testing.B) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(
		NewTextRun(
			"Hello World! This is a test paragraph with multiple words.",
			gm,
			12.0,
		),
	)

	b.ResetTimer()
	for range b.N {
		MeasureParagraph(p)
	}
}

func BenchmarkGetMinimumWidth(b *testing.B) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(
		NewTextRun(
			"Hello World! This is a test paragraph with multiple words.",
			gm,
			12.0,
		),
	)

	b.ResetTimer()
	for range b.N {
		GetMinimumWidth(p)
	}
}

func BenchmarkMeasureParagraphDetailed(
	b *testing.B,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)

	p := NewParagraph()
	p.AddRun(
		NewTextRun(
			"Hello World! This is a test paragraph with multiple words.",
			gm,
			12.0,
		),
	)
	p.Properties.Indentation.Left = 36
	p.Properties.Indentation.FirstLine = 18

	b.ResetTimer()
	for range b.N {
		MeasureParagraphDetailed(p)
	}
}

func BenchmarkSplitIntoWords(b *testing.B) {
	text := "Hello World! This is a test paragraph with multiple words and various spacing."

	b.ResetTimer()
	for range b.N {
		splitIntoWords(text)
	}
}
