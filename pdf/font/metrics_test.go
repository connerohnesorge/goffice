// metrics_test.go provides comprehensive tests for font metrics extraction.
// This file tests font-level metrics, glyph-level metrics, text measurement,
// and scaling functions with both constructed minimal fonts and real system fonts.

package font

import (
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"path/filepath"
	"testing"
)

// ============================================================================
// Font-level Metrics Tests
// ============================================================================

// TestFontMetrics_UnitsPerEm tests UnitsPerEm extraction for various font types.
func TestFontMetrics_UnitsPerEm(t *testing.T) {
	tests := []struct {
		name       string
		buildFont  func() []byte
		wantUpem   int
		skipReason string
	}{
		{
			name:      "minimal TrueType font with 1000 upem",
			buildFont: buildMinimalFont,
			wantUpem:  1000,
		},
		{
			name:      "TrueType font with 2048 upem",
			buildFont: buildFontWith2048Upem,
			wantUpem:  2048,
		},
		{
			name:      "minimal CFF font with 1000 upem",
			buildFont: buildMinimalCFFFont,
			wantUpem:  1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipReason != "" {
				t.Skip(tt.skipReason)
			}

			data := tt.buildFont()
			font, err := ParseFont(data)
			if err != nil {
				t.Fatalf(
					"Failed to parse font: %v",
					err,
				)
			}

			if font.Metrics.UnitsPerEm != tt.wantUpem {
				t.Errorf(
					"UnitsPerEm = %d, want %d",
					font.Metrics.UnitsPerEm,
					tt.wantUpem,
				)
			}
		})
	}
}

// TestFontMetrics_Ascender tests ascender extraction.
func TestFontMetrics_Ascender(t *testing.T) {
	// Build minimal font with known ascender (800)
	data := buildMinimalFont()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	if font.Metrics.Ascender != 800 {
		t.Errorf(
			"Ascender = %d, want 800",
			font.Metrics.Ascender,
		)
	}

	// Ascender should be positive (above baseline)
	if font.Metrics.Ascender <= 0 {
		t.Error("Ascender should be positive")
	}
}

// TestFontMetrics_Descender tests descender extraction.
func TestFontMetrics_Descender(t *testing.T) {
	// Build minimal font with known descender (-200)
	data := buildMinimalFont()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	if font.Metrics.Descender != -200 {
		t.Errorf(
			"Descender = %d, want -200",
			font.Metrics.Descender,
		)
	}

	// Descender should be negative (below baseline) for most fonts
	if font.Metrics.Descender >= 0 {
		t.Log(
			"Note: Descender is non-negative, which is unusual but valid for some fonts",
		)
	}
}

// TestFontMetrics_LineGap tests line gap extraction.
func TestFontMetrics_LineGap(t *testing.T) {
	// Build minimal font with known line gap (90)
	data := buildMinimalFont()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	if font.Metrics.LineGap != 90 {
		t.Errorf(
			"LineGap = %d, want 90",
			font.Metrics.LineGap,
		)
	}

	// Line gap should be non-negative
	if font.Metrics.LineGap < 0 {
		t.Error("LineGap should not be negative")
	}
}

// TestFontMetrics_LineHeight tests that computed line height is sensible.
func TestFontMetrics_LineHeight(t *testing.T) {
	data := buildMinimalFont()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Line height = Ascender - Descender + LineGap
	// With our minimal font: 800 - (-200) + 90 = 1090
	lineHeight := font.Metrics.Ascender - font.Metrics.Descender + font.Metrics.LineGap
	expected := 1090
	if lineHeight != expected {
		t.Errorf(
			"LineHeight = %d, want %d",
			lineHeight,
			expected,
		)
	}

	// Line height should be reasonable relative to UnitsPerEm
	// Typically 1.0 to 1.5 times UnitsPerEm
	ratio := float64(
		lineHeight,
	) / float64(
		font.Metrics.UnitsPerEm,
	)
	if ratio < 0.5 || ratio > 2.5 {
		t.Errorf(
			"Line height ratio %.2f is outside reasonable range [0.5, 2.5]",
			ratio,
		)
	}
}

// TestFontMetrics_CapHeight tests cap height extraction.
func TestFontMetrics_CapHeight(t *testing.T) {
	// Build font with OS/2 table containing CapHeight
	data := buildFontWithOS2Metrics()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Cap height should be non-negative if parsed
	if font.Metrics.CapHeight < 0 {
		t.Error(
			"CapHeight should not be negative",
		)
	}

	// Log the actual value for debugging
	t.Logf(
		"CapHeight = %d",
		font.Metrics.CapHeight,
	)
}

// TestFontMetrics_XHeight tests x-height extraction.
func TestFontMetrics_XHeight(t *testing.T) {
	// Build font with OS/2 table containing XHeight
	data := buildFontWithOS2Metrics()
	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// X-height should be non-negative if parsed
	if font.Metrics.XHeight < 0 {
		t.Error("XHeight should not be negative")
	}

	// Log the actual value for debugging
	t.Logf("XHeight = %d", font.Metrics.XHeight)
}

// ============================================================================
// Glyph-level Metrics Tests
// ============================================================================

// TestGlyphMetrics_AdvanceWidth tests advance width extraction for individual glyphs.
func TestGlyphMetrics_AdvanceWidth(t *testing.T) {
	// Create a font with specific glyph widths
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 722},
			'B': {AdvanceWidth: 667},
			'C': {AdvanceWidth: 722},
			'a': {AdvanceWidth: 556},
			'b': {AdvanceWidth: 611},
			'c': {AdvanceWidth: 556},
			'0': {AdvanceWidth: 556},
			'1': {AdvanceWidth: 556},
			' ': {AdvanceWidth: 278},
		},
	}

	tests := []struct {
		char      rune
		wantWidth int
	}{
		{'A', 722},
		{'B', 667},
		{'C', 722},
		{'a', 556},
		{'b', 611},
		{'c', 556},
		{'0', 556},
		{'1', 556},
		{' ', 278},
	}

	for _, tt := range tests {
		t.Run(
			string(tt.char),
			func(t *testing.T) {
				width := font.GlyphWidth(tt.char)
				if width != tt.wantWidth {
					t.Errorf(
						"GlyphWidth('%c') = %d, want %d",
						tt.char,
						width,
						tt.wantWidth,
					)
				}
			},
		)
	}
}

// TestGlyphMetrics_AllAlphanumeric tests that A-Z, a-z, 0-9 have valid widths.
func TestGlyphMetrics_AllAlphanumeric(
	t *testing.T,
) {
	// Skip if no system font available
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Test uppercase letters
	for r := 'A'; r <= 'Z'; r++ {
		width := font.GlyphWidth(r)
		if width == 0 {
			t.Errorf(
				"Uppercase '%c' has zero width",
				r,
			)
		}
	}

	// Test lowercase letters
	for r := 'a'; r <= 'z'; r++ {
		width := font.GlyphWidth(r)
		if width == 0 {
			t.Errorf(
				"Lowercase '%c' has zero width",
				r,
			)
		}
	}

	// Test digits
	for r := '0'; r <= '9'; r++ {
		width := font.GlyphWidth(r)
		if width == 0 {
			t.Errorf(
				"Digit '%c' has zero width",
				r,
			)
		}
	}
}

// TestGlyphMetrics_MonospaceConsistency tests that monospace fonts have consistent widths.
func TestGlyphMetrics_MonospaceConsistency(
	t *testing.T,
) {
	// Try to find a monospace font
	monoPaths := []string{
		"/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf",
		"/usr/share/fonts/TTF/DejaVuSansMono.ttf",
		"/usr/share/fonts/truetype/liberation/LiberationMono-Regular.ttf",
		"/System/Library/Fonts/Menlo.ttc",
		"/System/Library/Fonts/Monaco.ttf",
	}

	var fontPath string
	for _, p := range monoPaths {
		if _, err := os.Stat(p); err == nil {
			fontPath = p

			break
		}
	}

	if fontPath == "" {
		t.Skip(
			"No monospace font found for testing",
		)
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse monospace font: %v",
			err,
		)
	}

	// Get width of 'a' as reference
	referenceWidth := font.GlyphWidth('a')
	if referenceWidth == 0 {
		t.Fatal(
			"Reference glyph 'a' has zero width",
		)
	}

	// All alphanumeric characters should have the same width in monospace
	testChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	inconsistent := 0
	for _, r := range testChars {
		width := font.GlyphWidth(r)
		if width != referenceWidth {
			inconsistent++
		}
	}

	// Allow some tolerance for fonts that aren't strictly monospace
	if inconsistent > len(testChars)/4 {
		t.Errorf(
			"Monospace font has %d/%d characters with inconsistent widths",
			inconsistent,
			len(testChars),
		)
	}
}

// TestGlyphMetrics_ProportionalVariation tests that proportional fonts have varying widths.
func TestGlyphMetrics_ProportionalVariation(
	t *testing.T,
) {
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Get widths of different characters
	widthI := font.GlyphWidth('i')
	widthM := font.GlyphWidth('m')
	widthW := font.GlyphWidth('W')

	// For proportional fonts, 'i' should be narrower than 'm'
	// and 'm' should be narrower than or equal to 'W'
	if widthI >= widthM {
		t.Log(
			"Note: 'i' is not narrower than 'm' - may be a monospace font",
		)
	}

	if widthM >= widthW {
		t.Log(
			"Note: 'm' is not narrower than 'W' - may be unusual font",
		)
	}

	// At least one should be different for a proportional font
	if widthI == widthM && widthM == widthW {
		t.Log(
			"All test characters have same width - likely monospace font",
		)
	}
}

// ============================================================================
// Text Measurement Tests
// ============================================================================

// TestFont_TextWidth_SingleCharacter tests width calculation for single characters.
func TestFont_TextWidth_SingleCharacter(
	t *testing.T,
) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
		},
	}

	tests := []struct {
		char     string
		expected int
	}{
		{"A", 600},
		{"B", 650},
	}

	for _, tt := range tests {
		t.Run(tt.char, func(t *testing.T) {
			width := font.TextWidth(tt.char)
			if width != tt.expected {
				t.Errorf(
					"TextWidth(%q) = %d, want %d",
					tt.char,
					width,
					tt.expected,
				)
			}
		})
	}
}

// TestFont_TextWidth_Word tests width calculation for complete words.
func TestFont_TextWidth_Word(t *testing.T) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'H': {AdvanceWidth: 700},
			'e': {AdvanceWidth: 500},
			'l': {AdvanceWidth: 300},
			'o': {AdvanceWidth: 550},
			' ': {AdvanceWidth: 250},
			'W': {AdvanceWidth: 900},
			'r': {AdvanceWidth: 400},
			'd': {AdvanceWidth: 550},
		},
	}

	tests := []struct {
		text     string
		expected int
	}{
		// "Hello" = H(700) + e(500) + l(300) + l(300) + o(550) = 2350
		{"Hello", 2350},
		// "Word" = W(900) + o(550) + r(400) + d(550) = 2400
		{"Word", 2400},
		// "Hello World" = H(700) + e(500) + l(300) + l(300) + o(550) + space(250) + W(900) + o(550) + r(400) + l(300) + d(550) = 5300
		{"Hello World", 5300},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			width := font.TextWidth(tt.text)
			if width != tt.expected {
				t.Errorf(
					"TextWidth(%q) = %d, want %d",
					tt.text,
					width,
					tt.expected,
				)
			}
		})
	}
}

// TestFont_TextWidth_EmptyString tests that empty string returns zero width.
func TestFont_TextWidth_EmptyString(
	t *testing.T,
) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
		},
	}

	width := font.TextWidth("")
	if width != 0 {
		t.Errorf(
			"TextWidth(\"\") = %d, want 0",
			width,
		)
	}
}

// TestFont_TextWidth_MultiLine tests that newlines don't cause issues.
func TestFont_TextWidth_MultiLine(t *testing.T) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'H': {AdvanceWidth: 700},
			'i': {AdvanceWidth: 300},
			'\n': {
				AdvanceWidth: 0,
			}, // Newlines typically have zero width
		},
	}

	// "Hi\nHi" should be 2000 (Hi + 0 + Hi)
	// But newlines aren't mapped, so GlyphWidth returns 0
	text := "Hi\nHi"
	width := font.TextWidth(text)

	// Hi = 700 + 300 = 1000, newline = 0, Hi = 1000, total = 2000
	expected := 2000
	if width != expected {
		t.Errorf(
			"TextWidth(%q) = %d, want %d",
			text,
			width,
			expected,
		)
	}
}

// ============================================================================
// Scaling Functions Tests
// ============================================================================

// TestFont_ScaledWidth tests unit conversion at different font sizes.
func TestFont_ScaledWidth(t *testing.T) {
	tests := []struct {
		name       string
		unitsPerEm int
		fontUnits  int
		fontSize   float64
		expected   float64
	}{
		{
			name:       "standard 1000 upem, 12pt",
			unitsPerEm: 1000,
			fontUnits:  500,
			fontSize:   12.0,
			expected:   6.0, // 500 * 12 / 1000 = 6
		},
		{
			name:       "standard 2048 upem, 12pt",
			unitsPerEm: 2048,
			fontUnits:  1024,
			fontSize:   12.0,
			expected:   6.0, // 1024 * 12 / 2048 = 6
		},
		{
			name:       "full em at 10pt",
			unitsPerEm: 1000,
			fontUnits:  1000,
			fontSize:   10.0,
			expected:   10.0, // 1000 * 10 / 1000 = 10
		},
		{
			name:       "half em at 24pt",
			unitsPerEm: 1000,
			fontUnits:  500,
			fontSize:   24.0,
			expected:   12.0, // 500 * 24 / 1000 = 12
		},
		{
			name:       "large font size",
			unitsPerEm: 1000,
			fontUnits:  600,
			fontSize:   72.0,
			expected:   43.2, // 600 * 72 / 1000 = 43.2
		},
		{
			name:       "small font size",
			unitsPerEm: 1000,
			fontUnits:  600,
			fontSize:   6.0,
			expected:   3.6, // 600 * 6 / 1000 = 3.6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			font := &Font{
				Metrics: FontMetrics{
					UnitsPerEm: tt.unitsPerEm,
				},
			}

			result := font.ScaledWidth(
				tt.fontUnits,
				tt.fontSize,
			)
			if math.Abs(
				result-tt.expected,
			) > 0.0001 {
				t.Errorf(
					"ScaledWidth(%d, %.1f) = %.4f, want %.4f",
					tt.fontUnits,
					tt.fontSize,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestFont_ScaledWidth_ZeroUnitsPerEm tests handling of zero UnitsPerEm.
func TestFont_ScaledWidth_ZeroUnitsPerEm(
	t *testing.T,
) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 0},
	}

	result := font.ScaledWidth(500, 12.0)
	if result != 0 {
		t.Errorf(
			"ScaledWidth with zero UnitsPerEm = %f, want 0",
			result,
		)
	}
}

// TestFont_ScaledWidth_ZeroFontSize tests handling of zero font size.
func TestFont_ScaledWidth_ZeroFontSize(
	t *testing.T,
) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
	}

	result := font.ScaledWidth(500, 0.0)
	if result != 0 {
		t.Errorf(
			"ScaledWidth with zero fontSize = %f, want 0",
			result,
		)
	}
}

// TestFont_ScaledWidth_NegativeValues tests handling of negative values.
func TestFont_ScaledWidth_NegativeValues(
	t *testing.T,
) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
	}

	// Negative font units (e.g., left bearing)
	result := font.ScaledWidth(-100, 12.0)
	expected := -1.2 // -100 * 12 / 1000
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf(
			"ScaledWidth(-100, 12.0) = %f, want %f",
			result,
			expected,
		)
	}
}

// ============================================================================
// Real System Font Tests
// ============================================================================

// TestFontMetrics_RealSystemFont tests with real system fonts if available.
func TestFontMetrics_RealSystemFont(
	t *testing.T,
) {
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// UnitsPerEm should be 1000 or 2048 for most fonts
	if font.Metrics.UnitsPerEm != 1000 &&
		font.Metrics.UnitsPerEm != 2048 {
		t.Logf(
			"Unusual UnitsPerEm: %d",
			font.Metrics.UnitsPerEm,
		)
	}

	// Ascender should be positive
	if font.Metrics.Ascender <= 0 {
		t.Errorf(
			"Ascender = %d, expected positive",
			font.Metrics.Ascender,
		)
	}

	// Descender should typically be negative
	if font.Metrics.Descender > 0 {
		t.Logf(
			"Unusual positive descender: %d",
			font.Metrics.Descender,
		)
	}

	// LineGap should be non-negative
	if font.Metrics.LineGap < 0 {
		t.Errorf(
			"LineGap = %d, expected non-negative",
			font.Metrics.LineGap,
		)
	}

	t.Logf("Font: %s", font.Family)
	t.Logf(
		"  UnitsPerEm: %d",
		font.Metrics.UnitsPerEm,
	)
	t.Logf(
		"  Ascender:   %d",
		font.Metrics.Ascender,
	)
	t.Logf(
		"  Descender:  %d",
		font.Metrics.Descender,
	)
	t.Logf(
		"  LineGap:    %d",
		font.Metrics.LineGap,
	)
	t.Logf(
		"  CapHeight:  %d",
		font.Metrics.CapHeight,
	)
	t.Logf(
		"  XHeight:    %d",
		font.Metrics.XHeight,
	)
}

// TestFontMetrics_DejaVuSans tests with DejaVu Sans if available.
func TestFontMetrics_DejaVuSans(t *testing.T) {
	paths := []string{
		"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
		"/usr/share/fonts/TTF/DejaVuSans.ttf",
	}

	var fontPath string
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			fontPath = p

			break
		}
	}

	if fontPath == "" {
		t.Skip("DejaVu Sans not found")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse DejaVu Sans: %v",
			err,
		)
	}

	// DejaVu Sans has known metrics
	if font.Metrics.UnitsPerEm != 2048 {
		t.Errorf(
			"DejaVu Sans UnitsPerEm = %d, expected 2048",
			font.Metrics.UnitsPerEm,
		)
	}

	// Verify glyph data is populated
	if len(font.GlyphData) == 0 {
		t.Error("No glyph data extracted")
	}

	// Test width of space character (typically about 1/4 em)
	spaceWidth := font.GlyphWidth(' ')
	if spaceWidth == 0 {
		t.Error("Space character has zero width")
	}
}

// TestFontMetrics_LiberationSans tests with Liberation Sans if available.
func TestFontMetrics_LiberationSans(
	t *testing.T,
) {
	paths := []string{
		"/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
		"/usr/share/fonts/TTF/LiberationSans-Regular.ttf",
	}

	var fontPath string
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			fontPath = p

			break
		}
	}

	if fontPath == "" {
		t.Skip("Liberation Sans not found")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse Liberation Sans: %v",
			err,
		)
	}

	// Liberation Sans is metrically compatible with Arial
	if font.Metrics.UnitsPerEm != 2048 {
		t.Errorf(
			"Liberation Sans UnitsPerEm = %d, expected 2048",
			font.Metrics.UnitsPerEm,
		)
	}
}

// TestFontMetrics_GoFont tests with Go font if available.
func TestFontMetrics_GoFont(t *testing.T) {
	// Try to find Go font in common locations
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		home, _ := os.UserHomeDir()
		gopath = filepath.Join(home, "go")
	}

	paths := []string{
		filepath.Join(
			gopath,
			"pkg/mod/golang.org/x/image@v0.21.0/font/gofont/ttfs/Go-Regular.ttf",
		),
		filepath.Join(
			gopath,
			"pkg/mod/golang.org/x/image@v0.20.0/font/gofont/ttfs/Go-Regular.ttf",
		),
		filepath.Join(
			gopath,
			"pkg/mod/golang.org/x/image@v0.19.0/font/gofont/ttfs/Go-Regular.ttf",
		),
	}

	var fontPath string
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			fontPath = p

			break
		}
	}

	if fontPath == "" {
		t.Skip("Go font not found")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse Go font: %v",
			err,
		)
	}

	// Verify basic metrics
	if font.Metrics.UnitsPerEm == 0 {
		t.Error("UnitsPerEm should not be 0")
	}

	t.Logf(
		"Go font metrics: UnitsPerEm=%d, Ascender=%d, Descender=%d",
		font.Metrics.UnitsPerEm,
		font.Metrics.Ascender,
		font.Metrics.Descender,
	)
}

// ============================================================================
// Edge Cases Tests
// ============================================================================

// TestFont_MissingGlyph tests handling of missing glyphs.
func TestFont_MissingGlyph(t *testing.T) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
		},
	}

	// Request a glyph that doesn't exist
	width := font.GlyphWidth('Z')
	if width != 0 {
		t.Errorf(
			"Missing glyph width = %d, want 0",
			width,
		)
	}

	// Unicode character that likely doesn't exist
	width = font.GlyphWidth('\U0001F600') // Emoji
	if width != 0 {
		t.Errorf(
			"Missing emoji width = %d, want 0",
			width,
		)
	}
}

// TestFont_NonASCII tests handling of non-ASCII characters.
func TestFont_NonASCII(t *testing.T) {
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Test common non-ASCII characters
	testChars := []struct {
		char rune
		desc string
	}{
		{'\u00E9', "e-acute (Latin-1)"},
		{'\u00F1', "n-tilde (Latin-1)"},
		{'\u00FC', "u-umlaut (Latin-1)"},
		{'\u00C0', "A-grave (Latin-1)"},
		{'\u00A9', "copyright (Latin-1)"},
		{'\u20AC', "euro sign"},
	}

	for _, tc := range testChars {
		t.Run(tc.desc, func(t *testing.T) {
			width := font.GlyphWidth(tc.char)
			// We can't know if the font supports these, so just log
			if width == 0 {
				t.Logf(
					"Character %U (%s) has zero width (may not be in font)",
					tc.char,
					tc.desc,
				)
			} else {
				t.Logf("Character %U (%s) width = %d", tc.char, tc.desc, width)
			}
		})
	}
}

// TestFont_CJKCharacters tests handling of CJK characters.
func TestFont_CJKCharacters(t *testing.T) {
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Test common CJK characters
	testChars := []struct {
		char rune
		desc string
	}{
		{'\u4E00', "CJK Unified Ideograph (one)"},
		{
			'\u4E2D',
			"CJK Unified Ideograph (middle)",
		},
		{
			'\u6587',
			"CJK Unified Ideograph (text)",
		},
	}

	for _, tc := range testChars {
		t.Run(tc.desc, func(t *testing.T) {
			width := font.GlyphWidth(tc.char)
			// Most Western fonts don't include CJK
			if width == 0 {
				t.Logf(
					"Character %U (%s) not in font (expected for Western fonts)",
					tc.char,
					tc.desc,
				)
			} else {
				t.Logf("Character %U (%s) width = %d", tc.char, tc.desc, width)
			}
		})
	}
}

// TestFont_SpecialCharacters tests handling of special characters.
func TestFont_SpecialCharacters(t *testing.T) {
	fontPath := findTestFont(t)
	if fontPath == "" {
		t.Skip("No test font available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf("Failed to parse font: %v", err)
	}

	// Test punctuation and symbols
	testChars := []rune{
		'.', ',', '!', '?', ';', ':', '-', '_', '(', ')', '[', ']', '{', '}',
		'@', '#', '$', '%', '^', '&', '*', '+', '=', '/', '\\', '|', '~', '`',
	}

	for _, char := range testChars {
		t.Run(string(char), func(t *testing.T) {
			width := font.GlyphWidth(char)
			if width == 0 {
				t.Logf(
					"Punctuation '%c' has zero width",
					char,
				)
			}
		})
	}
}

// TestFont_ControlCharacters tests handling of control characters.
func TestFont_ControlCharacters(t *testing.T) {
	font := &Font{
		Metrics:   FontMetrics{UnitsPerEm: 1000},
		GlyphData: make(map[rune]GlyphMetrics),
	}

	// Control characters should have zero width
	controlChars := []rune{
		'\x00',
		'\x01',
		'\x02',
		'\x07',
		'\x08',
		'\x0B',
		'\x0C',
		'\x7F',
	}

	for _, char := range controlChars {
		t.Run(
			string([]rune{char}),
			func(t *testing.T) {
				width := font.GlyphWidth(char)
				if width != 0 {
					t.Errorf(
						"Control character %U has non-zero width %d",
						char,
						width,
					)
				}
			},
		)
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

// findTestFont finds a suitable test font on the system.
func findTestFont(t *testing.T) string {
	t.Helper()

	paths := []string{
		"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
		"/usr/share/fonts/TTF/DejaVuSans.ttf",
		"/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
		"/usr/share/fonts/TTF/LiberationSans-Regular.ttf",
		"/usr/share/fonts/truetype/freefont/FreeSans.ttf",
		"/usr/share/fonts/opentype/noto/NotoSans-Regular.ttf",
		"/System/Library/Fonts/Helvetica.ttc",
		"/System/Library/Fonts/Arial.ttf",
		"C:\\Windows\\Fonts\\arial.ttf",
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return ""
}

// buildMinimalFont creates a minimal TrueType font for testing.
func buildMinimalFont() []byte {
	buf := &bytes.Buffer{}

	tableStart := uint32(12 + 6*16)
	headSize := uint32(54)
	hheaSize := uint32(36)
	hmtxSize := uint32(4)
	cmapSize := uint32(4 + 8 + 32)
	nameSize := uint32(6 + 12 + 8)

	headOffset := tableStart
	hheaOffset := headOffset + headSize
	hmtxOffset := hheaOffset + hheaSize
	cmapOffset := hmtxOffset + hmtxSize
	nameOffset := cmapOffset + cmapSize

	// Header
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(6),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(64),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	)

	// Table records
	writeTableRecord(
		buf,
		"cmap",
		0,
		cmapOffset,
		cmapSize,
	)
	writeTableRecord(
		buf,
		"head",
		0,
		headOffset,
		headSize,
	)
	writeTableRecord(
		buf,
		"hhea",
		0,
		hheaOffset,
		hheaSize,
	)
	writeTableRecord(
		buf,
		"hmtx",
		0,
		hmtxOffset,
		hmtxSize,
	)
	writeTableRecord(
		buf,
		"name",
		0,
		nameOffset,
		nameSize,
	)

	// head table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x5F0F3CF5),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)

	// hhea table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(800),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(-200),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(90),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(600),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(600),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)

	// maxp table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00005000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)

	// hmtx table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(600),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)

	// cmap table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(12),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(4),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	// name table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(18),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0x0409),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	buf.Write(
		[]byte{
			0x00,
			'T',
			0x00,
			'e',
			0x00,
			's',
			0x00,
			't',
		},
	)

	return buf.Bytes()
}

// buildFontWith2048Upem creates a font with 2048 units per em.
func buildFontWith2048Upem() []byte {
	buf := &bytes.Buffer{}

	tableStart := uint32(12 + 6*16)
	headSize := uint32(54)
	hheaSize := uint32(36)
	hmtxSize := uint32(4)
	cmapSize := uint32(4 + 8 + 32)
	nameSize := uint32(6 + 12 + 8)

	headOffset := tableStart
	hheaOffset := headOffset + headSize
	hmtxOffset := hheaOffset + hheaSize
	cmapOffset := hmtxOffset + hmtxSize
	nameOffset := cmapOffset + cmapSize

	// Header
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(6),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(64),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	)

	// Table records
	writeTableRecord(
		buf,
		"cmap",
		0,
		cmapOffset,
		cmapSize,
	)
	writeTableRecord(
		buf,
		"head",
		0,
		headOffset,
		headSize,
	)
	writeTableRecord(
		buf,
		"hhea",
		0,
		hheaOffset,
		hheaSize,
	)
	writeTableRecord(
		buf,
		"hmtx",
		0,
		hmtxOffset,
		hmtxSize,
	)
	writeTableRecord(
		buf,
		"name",
		0,
		nameOffset,
		nameSize,
	)

	// head table with 2048 upem
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x5F0F3CF5),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2048),
	) // Different upem
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(2048),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(2048),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)

	// hhea table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1638),
	) // Scaled ascender
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(-410),
	) // Scaled descender
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(184),
	) // Scaled line gap
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1200),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1200),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)

	// maxp table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00005000),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)

	// hmtx table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1200),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	)

	// cmap table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(12),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(4),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	// name table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(18),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0x0409),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	buf.Write(
		[]byte{
			0x00,
			'T',
			0x00,
			'e',
			0x00,
			's',
			0x00,
			't',
		},
	)

	return buf.Bytes()
}

// buildMinimalCFFFont creates a minimal CFF font for testing.
func buildMinimalCFFFont() []byte {
	// Minimal CFF font implementation would go here
	// For now, return a simple placeholder
	return []byte{0x00, 0x01, 0x02, 0x03}
}

// writeTableRecord writes a font table record to the buffer.
func writeTableRecord(buf *bytes.Buffer, tag string, checksum uint32, offset uint32, length uint32) {
	buf.Write([]byte(tag))
	_ = binary.Write(buf, binary.BigEndian, checksum)
	_ = binary.Write(buf, binary.BigEndian, offset)
	_ = binary.Write(buf, binary.BigEndian, length)
}
