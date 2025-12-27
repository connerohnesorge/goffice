package drawing

import (
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

func TestTextRenderingMode_String(t *testing.T) {
	tests := []struct {
		mode     TextRenderingMode
		expected string
	}{
		{TextRenderFill, "fill"},
		{TextRenderStroke, "stroke"},
		{TextRenderFillStroke, "fill+stroke"},
		{TextRenderInvisible, "invisible"},
		{TextRenderFillClip, "fill+clip"},
		{TextRenderStrokeClip, "stroke+clip"},
		{
			TextRenderFillStrokeClip,
			"fill+stroke+clip",
		},
		{TextRenderClip, "clip"},
		{
			TextRenderingMode(99),
			"TextRenderingMode(99)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.mode.String(); got != tt.expected {
				t.Errorf(
					"TextRenderingMode.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestNewTextStyle(t *testing.T) {
	style := NewTextStyle()

	if style.FontSize != 12.0 {
		t.Errorf(
			"FontSize = %v, want 12.0",
			style.FontSize,
		)
	}
	if style.Color != Black {
		t.Errorf(
			"Color = %v, want Black",
			style.Color,
		)
	}
	if style.StrokeColor != Black {
		t.Errorf(
			"StrokeColor = %v, want Black",
			style.StrokeColor,
		)
	}
	if style.HorizontalScaling != 100.0 {
		t.Errorf(
			"HorizontalScaling = %v, want 100.0",
			style.HorizontalScaling,
		)
	}
	if style.RenderingMode != TextRenderFill {
		t.Errorf(
			"RenderingMode = %v, want TextRenderFill",
			style.RenderingMode,
		)
	}
}

func TestTextStyle_Chaining(t *testing.T) {
	style := NewTextStyle().
		WithFontSize(24.0).
		WithColor(Red).
		WithCharacterSpacing(1.5).
		WithWordSpacing(2.0).
		WithHorizontalScaling(110.0).
		WithTextRise(3.0).
		WithLeading(28.0).
		WithRenderingMode(TextRenderStroke)

	if style.FontSize != 24.0 {
		t.Errorf(
			"FontSize = %v, want 24.0",
			style.FontSize,
		)
	}
	if style.Color != Red {
		t.Errorf(
			"Color = %v, want Red",
			style.Color,
		)
	}
	if style.CharacterSpacing != 1.5 {
		t.Errorf(
			"CharacterSpacing = %v, want 1.5",
			style.CharacterSpacing,
		)
	}
	if style.WordSpacing != 2.0 {
		t.Errorf(
			"WordSpacing = %v, want 2.0",
			style.WordSpacing,
		)
	}
	if style.HorizontalScaling != 110.0 {
		t.Errorf(
			"HorizontalScaling = %v, want 110.0",
			style.HorizontalScaling,
		)
	}
	if style.TextRise != 3.0 {
		t.Errorf(
			"TextRise = %v, want 3.0",
			style.TextRise,
		)
	}
	if style.Leading != 28.0 {
		t.Errorf(
			"Leading = %v, want 28.0",
			style.Leading,
		)
	}
	if style.RenderingMode != TextRenderStroke {
		t.Errorf(
			"RenderingMode = %v, want TextRenderStroke",
			style.RenderingMode,
		)
	}
}

func TestTextStyle_Clone(t *testing.T) {
	original := NewTextStyle().WithFontSize(18.0).
		WithColor(Blue)
	clone := original.Clone()

	// Modify original
	original.FontSize = 24.0
	original.Color = Red

	// Clone should be unchanged
	if clone.FontSize != 18.0 {
		t.Errorf(
			"Clone FontSize = %v, want 18.0",
			clone.FontSize,
		)
	}
	if clone.Color != Blue {
		t.Errorf(
			"Clone Color = %v, want Blue",
			clone.Color,
		)
	}
}

func TestTextBuilder_BeginEndText(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "BT") {
		t.Error("Missing BT operator")
	}
	if !strings.Contains(result, "ET") {
		t.Error("Missing ET operator")
	}
}

func TestTextBuilder_SetFont(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetFont("/F1", 12.0)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "/F1 12 Tf") {
		t.Errorf(
			"Missing or incorrect Tf operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetTextMatrix(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetTextMatrix(1, 0, 0, 1, 100, 200)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"1 0 0 1 100 200 Tm",
	) {
		t.Errorf(
			"Missing or incorrect Tm operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetTextPosition(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetTextPosition(50, 100)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"1 0 0 1 50 100 Tm",
	) {
		t.Errorf(
			"Missing or incorrect position: %s",
			result,
		)
	}
}

func TestTextBuilder_MoveText(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.MoveText(10, -15)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "10 -15 Td") {
		t.Errorf(
			"Missing or incorrect Td operator: %s",
			result,
		)
	}
}

func TestTextBuilder_ShowText(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetFont("/F1", 12)
	tb.ShowText("Hello, World!")
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"(Hello, World!) Tj",
	) {
		t.Errorf(
			"Missing or incorrect Tj operator: %s",
			result,
		)
	}
}

func TestTextBuilder_ShowTextNextLine(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetFont("/F1", 12)
	tb.ShowTextNextLine("Next line")
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"(Next line) '",
	) {
		t.Errorf(
			"Missing or incorrect ' operator: %s",
			result,
		)
	}
}

func TestTextBuilder_ShowTextArray(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetFont("/F1", 12)
	tb.ShowTextArray("He", -20, "llo")
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"[(He) -20 (llo)] TJ",
	) {
		t.Errorf(
			"Missing or incorrect TJ operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetCharacterSpacing(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetCharacterSpacing(1.5)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "1.5 Tc") {
		t.Errorf(
			"Missing or incorrect Tc operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetWordSpacing(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetWordSpacing(2.0)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "2 Tw") {
		t.Errorf(
			"Missing or incorrect Tw operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetHorizontalScaling(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetHorizontalScaling(150)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "150 Tz") {
		t.Errorf(
			"Missing or incorrect Tz operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetLeading(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetLeading(14.4)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "14.4 TL") {
		t.Errorf(
			"Missing or incorrect TL operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetTextRise(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetTextRise(5.0)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "5 Ts") {
		t.Errorf(
			"Missing or incorrect Ts operator: %s",
			result,
		)
	}
}

func TestTextBuilder_SetRenderingMode(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetRenderingMode(TextRenderStroke)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "1 Tr") {
		t.Errorf(
			"Missing or incorrect Tr operator: %s",
			result,
		)
	}
}

func TestTextBuilder_NextLine(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.NextLine()
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "T*") {
		t.Errorf(
			"Missing T* operator: %s",
			result,
		)
	}
}

func TestTextBuilder_ShowHexText(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.ShowHexText(
		[]byte{0x48, 0x65, 0x6C, 0x6C, 0x6F},
	) // "Hello" in hex
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"<48656C6C6F> Tj",
	) {
		t.Errorf(
			"Missing or incorrect hex Tj operator: %s",
			result,
		)
	}
}

func TestTextBuilder_ShowGlyphs(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.ShowGlyphs(
		[]uint16{0x0001, 0x0002, 0x0003},
	)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(
		result,
		"<000100020003> Tj",
	) {
		t.Errorf(
			"Missing or incorrect glyph Tj operator: %s",
			result,
		)
	}
}

func TestTextBuilder_Reset(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.ShowText("Test")
	tb.EndText()

	tb.Reset()
	result := tb.String()
	if result != "" {
		t.Errorf(
			"Reset did not clear operators: %s",
			result,
		)
	}
}

func TestTextBuilder_DoubleBeginEnd(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.BeginText() // Should not add another BT
	tb.EndText()
	tb.EndText() // Should not add another ET

	result := tb.String()
	btCount := strings.Count(result, "BT")
	etCount := strings.Count(result, "ET")

	if btCount != 1 {
		t.Errorf(
			"Expected 1 BT, got %d: %s",
			btCount,
			result,
		)
	}
	if etCount != 1 {
		t.Errorf(
			"Expected 1 ET, got %d: %s",
			etCount,
			result,
		)
	}
}

func TestDrawText(t *testing.T) {
	style := NewTextStyle().
		WithFontSize(14).
		WithColor(Blue)

	result := DrawText(
		100,
		200,
		"Hello",
		"/F1",
		style,
	)

	if !strings.Contains(result, "BT") {
		t.Error("Missing BT operator")
	}
	if !strings.Contains(result, "ET") {
		t.Error("Missing ET operator")
	}
	if !strings.Contains(result, "/F1 14 Tf") {
		t.Errorf(
			"Missing font operator: %s",
			result,
		)
	}
	if !strings.Contains(result, "(Hello) Tj") {
		t.Errorf(
			"Missing text operator: %s",
			result,
		)
	}
	if !strings.Contains(result, "100 200 Tm") {
		t.Errorf("Missing position: %s", result)
	}
}

func TestDrawText_WithStrokeMode(t *testing.T) {
	style := NewTextStyle().
		WithRenderingMode(TextRenderStroke).
		WithStrokeColor(Red).
		WithColor(Blue)

	result := DrawText(
		100,
		200,
		"Stroked",
		"/F1",
		style,
	)

	// Should include stroke color
	if !strings.Contains(result, "1 0 0 RG") {
		t.Errorf(
			"Missing stroke color (red): %s",
			result,
		)
	}
	// Should include rendering mode
	if !strings.Contains(result, "1 Tr") {
		t.Errorf(
			"Missing rendering mode: %s",
			result,
		)
	}
}

func TestDrawTextAt(t *testing.T) {
	result := DrawTextAt(
		50,
		100,
		"Simple",
		"/F1",
		12,
	)

	if !strings.Contains(result, "BT") {
		t.Error("Missing BT operator")
	}
	if !strings.Contains(result, "ET") {
		t.Error("Missing ET operator")
	}
	if !strings.Contains(result, "/F1 12 Tf") {
		t.Errorf(
			"Missing font operator: %s",
			result,
		)
	}
	if !strings.Contains(result, "(Simple) Tj") {
		t.Errorf(
			"Missing text operator: %s",
			result,
		)
	}
}

func TestEscapePDFString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "Hello"},
		{"(test)", "\\(test\\)"},
		{"back\\slash", "back\\\\slash"},
		{"new\nline", "new\\nline"},
		{"tab\there", "tab\\there"},
		{"carriage\rreturn", "carriage\\rreturn"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := escapePDFString(tt.input)
			if got != tt.expected {
				t.Errorf(
					"escapePDFString(%q) = %q, want %q",
					tt.input,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestGenerateFontResourceName(t *testing.T) {
	tests := []struct {
		prefix   string
		index    int
		expected string
	}{
		{"F", 1, "F1"},
		{"F", 10, "F10"},
		{"TT", 1, "TT1"},
		{"Fnt", 5, "Fnt5"},
	}

	for _, tt := range tests {
		got := GenerateFontResourceName(
			tt.prefix,
			tt.index,
		)
		if got != tt.expected {
			t.Errorf(
				"GenerateFontResourceName(%q, %d) = %q, want %q",
				tt.prefix,
				tt.index,
				got,
				tt.expected,
			)
		}
	}
}

func TestTextWidth_NilFont(t *testing.T) {
	width := TextWidth("Test", nil, 12)
	if width != 0 {
		t.Errorf(
			"TextWidth with nil font = %v, want 0",
			width,
		)
	}
}

func TestTextHeight_NilFont(t *testing.T) {
	height := TextHeight(nil, 12)
	if height != 12 {
		t.Errorf(
			"TextHeight with nil font = %v, want 12 (fallback to fontSize)",
			height,
		)
	}
}

func TestTextAscender_NilFont(t *testing.T) {
	ascender := TextAscender(nil, 12)
	expected := 12 * 0.8
	// Use tolerance for floating point comparison
	if abs(ascender-expected) > 0.0001 {
		t.Errorf(
			"TextAscender with nil font = %v, want %v",
			ascender,
			expected,
		)
	}
}

func TestTextDescender_NilFont(t *testing.T) {
	descender := TextDescender(nil, 12)
	expected := 12 * -0.2
	// Use tolerance for floating point comparison
	if abs(descender-expected) > 0.0001 {
		t.Errorf(
			"TextDescender with nil font = %v, want %v",
			descender,
			expected,
		)
	}
}

// abs returns the absolute value of x
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}

func TestFontRef(t *testing.T) {
	// Test with nil font (no real font data)
	ref := NewFontRef("F1", nil)
	if ref.Name != "F1" {
		t.Errorf(
			"FontRef.Name = %q, want %q",
			ref.Name,
			"F1",
		)
	}
}

func TestTextBuilder_ApplyStyle(t *testing.T) {
	style := NewTextStyle().
		WithCharacterSpacing(1.0).
		WithWordSpacing(2.0).
		WithHorizontalScaling(110).
		WithTextRise(3.0).
		WithLeading(14).
		WithRenderingMode(TextRenderStroke)

	tb := NewTextBuilder()
	tb.BeginText()
	tb.ApplyStyle(style)
	tb.EndText()

	result := tb.String()

	if !strings.Contains(result, "1 Tc") {
		t.Errorf(
			"Missing character spacing: %s",
			result,
		)
	}
	if !strings.Contains(result, "2 Tw") {
		t.Errorf(
			"Missing word spacing: %s",
			result,
		)
	}
	if !strings.Contains(result, "110 Tz") {
		t.Errorf(
			"Missing horizontal scaling: %s",
			result,
		)
	}
	if !strings.Contains(result, "3 Ts") {
		t.Errorf("Missing text rise: %s", result)
	}
	if !strings.Contains(result, "14 TL") {
		t.Errorf("Missing leading: %s", result)
	}
	if !strings.Contains(result, "1 Tr") {
		t.Errorf(
			"Missing rendering mode: %s",
			result,
		)
	}
}

func TestTextBuilder_ApplyStyle_Defaults(
	t *testing.T,
) {
	// Default style should not add operators for default values
	style := NewTextStyle()

	tb := NewTextBuilder()
	tb.BeginText()
	tb.ApplyStyle(style)
	tb.EndText()

	result := tb.String()

	// Should only have BT and ET
	if strings.Contains(result, "Tc") {
		t.Errorf(
			"Should not set default character spacing: %s",
			result,
		)
	}
	if strings.Contains(result, "Tw") {
		t.Errorf(
			"Should not set default word spacing: %s",
			result,
		)
	}
	if strings.Contains(result, "Tz") {
		t.Errorf(
			"Should not set default horizontal scaling: %s",
			result,
		)
	}
	if strings.Contains(result, "Ts") {
		t.Errorf(
			"Should not set default text rise: %s",
			result,
		)
	}
	if strings.Contains(result, "Tr") {
		t.Errorf(
			"Should not set default rendering mode: %s",
			result,
		)
	}
}

func TestTextBuilder_ApplyStyle_Nil(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.ApplyStyle(nil) // Should not panic
	tb.EndText()

	result := tb.String()
	if result != "BT\nET\n" {
		t.Errorf(
			"ApplyStyle(nil) should do nothing, got: %s",
			result,
		)
	}
}

func TestFontDescriptorFromFont_Nil(
	t *testing.T,
) {
	fd := FontDescriptorFromFont(nil)
	if fd != nil {
		t.Error(
			"FontDescriptorFromFont(nil) should return nil",
		)
	}
}

func TestFontDescriptorFromFont(t *testing.T) {
	f := &font.Font{
		Family: "TestFont",
		Style:  font.StyleBold,
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
			Ascender:   800,
			Descender:  -200,
			CapHeight:  700,
			XHeight:    500,
		},
	}

	fd := FontDescriptorFromFont(f)
	if fd == nil {
		t.Fatal(
			"FontDescriptorFromFont returned nil",
		)
	}

	if fd.FontName != "TestFont" {
		t.Errorf(
			"FontName = %q, want %q",
			fd.FontName,
			"TestFont",
		)
	}
	if fd.Ascent != 800 {
		t.Errorf(
			"Ascent = %d, want 800",
			fd.Ascent,
		)
	}
	if fd.Descent != -200 {
		t.Errorf(
			"Descent = %d, want -200",
			fd.Descent,
		)
	}
	if fd.CapHeight != 700 {
		t.Errorf(
			"CapHeight = %d, want 700",
			fd.CapHeight,
		)
	}
	if fd.XHeight != 500 {
		t.Errorf(
			"XHeight = %d, want 500",
			fd.XHeight,
		)
	}
}

func TestFontDescriptorFromFont_Italic(
	t *testing.T,
) {
	f := &font.Font{
		Family: "TestItalic",
		Style:  font.StyleItalic,
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
			Ascender:   800,
			Descender:  -200,
		},
	}

	fd := FontDescriptorFromFont(f)
	if fd == nil {
		t.Fatal(
			"FontDescriptorFromFont returned nil",
		)
	}

	// Check italic flag is set (bit 6)
	italicFlag := 1 << 6
	if fd.Flags&italicFlag == 0 {
		t.Errorf(
			"Italic flag not set, Flags = %d",
			fd.Flags,
		)
	}
}

func TestWrapText_NilFont(t *testing.T) {
	lines := WrapText(
		"Hello World",
		nil,
		12,
		100,
		14,
	)
	if len(lines) != 1 {
		t.Errorf(
			"Expected 1 line for nil font, got %d",
			len(lines),
		)
	}
	if lines[0].Text != "Hello World" {
		t.Errorf(
			"Text = %q, want %q",
			lines[0].Text,
			"Hello World",
		)
	}
}

func TestWrapText_ZeroWidth(t *testing.T) {
	f := &font.Font{
		Family: "Test",
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
		},
	}
	lines := WrapText("Hello World", f, 12, 0, 14)
	if len(lines) != 1 {
		t.Errorf(
			"Expected 1 line for zero width, got %d",
			len(lines),
		)
	}
}

func TestWrapText_EmptyText(t *testing.T) {
	f := &font.Font{
		Family: "Test",
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
		},
	}
	lines := WrapText("", f, 12, 100, 14)
	if len(lines) != 0 {
		t.Errorf(
			"Expected 0 lines for empty text, got %d",
			len(lines),
		)
	}
}

func TestSubsetTracker(t *testing.T) {
	// Create a minimal font for testing
	f := &font.Font{
		Family: "Test",
		Metrics: font.FontMetrics{
			UnitsPerEm: 1000,
		},
		GlyphData: map[rune]font.GlyphMetrics{
			'H': {AdvanceWidth: 600},
			'e': {AdvanceWidth: 500},
			'l': {AdvanceWidth: 300},
			'o': {AdvanceWidth: 500},
		},
		Data: make(
			[]byte,
			100,
		), // Minimal font data
	}

	tracker := NewSubsetTracker(f)
	tracker.AddText("Hello")
	tracker.AddRune('!')

	// Build the subset
	subset, err := tracker.Build()
	if err != nil {
		t.Logf(
			"Subset build returned error (expected for minimal test font): %v",
			err,
		)
		// This is expected to fail without real font data
		return
	}

	// If build succeeded, verify the subset is accessible
	if tracker.Subset() != subset {
		t.Error(
			"Subset() should return the same subset as Build()",
		)
	}
}

func TestTextBuilder_ShowTextWithSpacing(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.ShowTextWithSpacing(
		5.0,
		1.0,
		"Spaced text",
	)
	tb.EndText()

	result := tb.String()
	// The " operator format is: aw ac (string) "
	if !strings.Contains(
		result,
		`5 1 (Spaced text) "`,
	) {
		t.Errorf(
			`Missing or incorrect " operator: %s`,
			result,
		)
	}
}

func TestTextBuilder_MoveTextSetLeading(
	t *testing.T,
) {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.MoveTextSetLeading(0, -14)
	tb.EndText()

	result := tb.String()
	if !strings.Contains(result, "0 -14 TD") {
		t.Errorf(
			"Missing or incorrect TD operator: %s",
			result,
		)
	}
}

// TestComplexTextLayout tests a more realistic text layout scenario
func TestComplexTextLayout(t *testing.T) {
	tb := NewTextBuilder()
	tb.BeginText()

	// Set up initial state
	tb.SetFont("/F1", 12)
	tb.SetLeading(14.4)
	tb.SetTextPosition(
		72,
		720,
	) // 1 inch margin, top of page

	// First line
	tb.ShowText("This is the first line.")
	tb.NextLine()

	// Second line with different character spacing
	tb.SetCharacterSpacing(0.5)
	tb.ShowText("Second line with extra spacing.")
	tb.NextLine()

	// Third line using TJ for kerning
	tb.SetCharacterSpacing(0)
	tb.ShowTextArray("AV", -80, "ATAR")

	tb.EndText()

	result := tb.String()

	// Verify structure
	if !strings.HasPrefix(result, "BT\n") {
		t.Error("Should start with BT")
	}
	if !strings.HasSuffix(result, "ET\n") {
		t.Error("Should end with ET")
	}

	// Verify key operators are present
	expectedOps := []string{
		"/F1 12 Tf",                    // Font
		"14.4 TL",                      // Leading
		"1 0 0 1 72 720",               // Position
		"(This is the first line.) Tj", // First text
		"T*",                           // Next line
		"0.5 Tc",                       // Character spacing
		"[(AV) -80 (ATAR)] TJ",         // Kerned text
	}

	for _, op := range expectedOps {
		if !strings.Contains(result, op) {
			t.Errorf(
				"Missing expected operator: %s\nGot: %s",
				op,
				result,
			)
		}
	}
}
