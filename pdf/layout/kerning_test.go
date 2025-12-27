package layout

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// mockFontWithKerning creates a font for testing kerning functionality.
// Note: This creates a mock font without actual kern table data.
// The real kerning parsing is tested with actual font data.
func mockFontWithKerning() *font.Font {
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
			'o': {
				AdvanceWidth: 550,
				LeftBearing:  25,
			},
			'W': {
				AdvanceWidth: 900,
				LeftBearing:  15,
			},
			'a': {
				AdvanceWidth: 480,
				LeftBearing:  30,
			},
			'e': {
				AdvanceWidth: 500,
				LeftBearing:  25,
			},
			'l': {
				AdvanceWidth: 250,
				LeftBearing:  40,
			},
			'r': {
				AdvanceWidth: 350,
				LeftBearing:  35,
			},
			'n': {
				AdvanceWidth: 500,
				LeftBearing:  30,
			},
			'i': {
				AdvanceWidth: 230,
				LeftBearing:  50,
			},
			'g': {
				AdvanceWidth: 520,
				LeftBearing:  28,
			},
			' ': {
				AdvanceWidth: 250,
				LeftBearing:  0,
			},
			'H': {
				AdvanceWidth: 700,
				LeftBearing:  50,
			},
		},
		Data: nil, // No actual font data for mock
	}
}

func TestNewKerningTable(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	if kt == nil {
		t.Fatal(
			"NewKerningTable returned nil for valid font",
		)
	}

	if kt.font != f {
		t.Error(
			"KerningTable does not reference the correct font",
		)
	}

	if kt.glyphMetrics == nil {
		t.Error(
			"KerningTable should have GlyphMetrics initialized",
		)
	}
}

func TestNewKerningTableNilFont(t *testing.T) {
	kt := NewKerningTable(nil)

	if kt != nil {
		t.Error(
			"NewKerningTable should return nil for nil font",
		)
	}
}

func TestKerningTableNilMethods(t *testing.T) {
	var kt *KerningTable

	// All methods should handle nil receiver gracefully
	if kt.HasKerning() {
		t.Error(
			"nil KerningTable should return false for HasKerning",
		)
	}

	if kt.GetKerning('A', 'V') != 0 {
		t.Error(
			"nil KerningTable should return 0 for GetKerning",
		)
	}

	if kt.GetKerningByGlyphID(1, 2) != 0 {
		t.Error(
			"nil KerningTable should return 0 for GetKerningByGlyphID",
		)
	}

	if kt.MeasureStringWithKerning("Hello") != 0 {
		t.Error(
			"nil KerningTable should return 0 for MeasureStringWithKerning",
		)
	}

	if kt.MeasureStringWithKerningScaled(
		"Hello",
		12.0,
	) != 0 {
		t.Error(
			"nil KerningTable should return 0 for MeasureStringWithKerningScaled",
		)
	}

	if kt.KernPairCount() != 0 {
		t.Error(
			"nil KerningTable should return 0 for KernPairCount",
		)
	}

	if kt.Font() != nil {
		t.Error(
			"nil KerningTable should return nil for Font",
		)
	}

	if kt.GlyphMetrics() != nil {
		t.Error(
			"nil KerningTable should return nil for GlyphMetrics",
		)
	}
}

func TestHasKerningNoData(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Mock font has no kern table data, so HasKerning should be false
	if kt.HasKerning() {
		t.Error(
			"HasKerning should return false for font without kern data",
		)
	}
}

func TestGetKerningNoData(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// With no kerning data, all pairs should return 0
	tests := []struct {
		left  rune
		right rune
	}{
		{'A', 'V'},
		{'T', 'o'},
		{'W', 'a'},
		{'V', 'A'},
		{'L', 'T'},
	}

	for _, tt := range tests {
		kern := kt.GetKerning(tt.left, tt.right)
		if kern != 0 {
			t.Errorf(
				"GetKerning(%q, %q) = %d, want 0 (no kern data)",
				tt.left,
				tt.right,
				kern,
			)
		}
	}
}

func TestGetKerningByGlyphIDNoData(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// With no kerning data, all pairs should return 0
	kern := kt.GetKerningByGlyphID(1, 2)
	if kern != 0 {
		t.Errorf(
			"GetKerningByGlyphID(1, 2) = %d, want 0",
			kern,
		)
	}
}

func TestMeasureStringWithKerningNoData(
	t *testing.T,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Without kerning data, this should equal regular string measurement
	text := "AVTo"
	width := kt.MeasureStringWithKerning(text)

	// Calculate expected width without kerning
	// A=700, V=650, T=600, o=550 = 2500
	expectedWidth := 700 + 650 + 600 + 550

	if width != expectedWidth {
		t.Errorf(
			"MeasureStringWithKerning(%q) = %d, want %d",
			text,
			width,
			expectedWidth,
		)
	}
}

func TestMeasureStringWithKerningEmpty(
	t *testing.T,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	width := kt.MeasureStringWithKerning("")
	if width != 0 {
		t.Errorf(
			"MeasureStringWithKerning(\"\") = %d, want 0",
			width,
		)
	}
}

func TestMeasureStringWithKerningSingleChar(
	t *testing.T,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Single character should return its width
	width := kt.MeasureStringWithKerning("A")
	if width != 700 {
		t.Errorf(
			"MeasureStringWithKerning(\"A\") = %d, want 700",
			width,
		)
	}
}

func TestMeasureStringWithKerningScaled(
	t *testing.T,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// "AV" = 700 + 650 = 1350 font units
	// At 12pt with 1000 upem: 1350 * 12 / 1000 = 16.2
	width := kt.MeasureStringWithKerningScaled(
		"AV",
		12.0,
	)
	expected := 16.2

	if width != expected {
		t.Errorf(
			"MeasureStringWithKerningScaled(\"AV\", 12) = %v, want %v",
			width,
			expected,
		)
	}
}

func TestKernPairCount(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Mock font has no kern data
	if kt.KernPairCount() != 0 {
		t.Errorf(
			"KernPairCount() = %d, want 0",
			kt.KernPairCount(),
		)
	}
}

func TestFontAccessor(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	if kt.Font() != f {
		t.Error(
			"Font() should return the underlying font",
		)
	}
}

func TestGlyphMetricsAccessor(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	gm := kt.GlyphMetrics()
	if gm == nil {
		t.Fatal(
			"GlyphMetrics() should not return nil",
		)
	}

	// Verify it works correctly
	if gm.GetAdvanceWidthForRune('A') != 700 {
		t.Error(
			"GlyphMetrics should return correct advance widths",
		)
	}
}

// TestKerningTableWithManualPairs tests kerning by manually adding pairs.
// This simulates what would happen with actual kern table data.
func TestKerningTableWithManualPairs(
	t *testing.T,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Manually add some kerning pairs (simulating parsed data)
	// In real usage, these come from parsing the kern table
	kt.pairs[uint32(1)<<16|uint32(2)] = -50 // Glyph 1 + 2 = -50
	kt.pairs[uint32(3)<<16|uint32(4)] = -30 // Glyph 3 + 4 = -30
	kt.hasKerning = true

	if !kt.HasKerning() {
		t.Error(
			"HasKerning should return true after adding pairs",
		)
	}

	if kt.KernPairCount() != 2 {
		t.Errorf(
			"KernPairCount() = %d, want 2",
			kt.KernPairCount(),
		)
	}

	// Test glyph ID lookup
	if kt.GetKerningByGlyphID(1, 2) != -50 {
		t.Errorf(
			"GetKerningByGlyphID(1, 2) = %d, want -50",
			kt.GetKerningByGlyphID(1, 2),
		)
	}

	if kt.GetKerningByGlyphID(3, 4) != -30 {
		t.Errorf(
			"GetKerningByGlyphID(3, 4) = %d, want -30",
			kt.GetKerningByGlyphID(3, 4),
		)
	}

	// Non-existent pair should return 0
	if kt.GetKerningByGlyphID(1, 3) != 0 {
		t.Errorf(
			"GetKerningByGlyphID(1, 3) = %d, want 0",
			kt.GetKerningByGlyphID(1, 3),
		)
	}
}

// TestKerningCacheRunePairs tests that rune pair lookups are cached.
func TestKerningCacheRunePairs(t *testing.T) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	// Enable kerning by adding manual pairs
	// We need to add a pair that corresponds to the glyph IDs for 'A' and 'V'
	// Since the glyph ID calculation is (rune % 65534) + 1
	leftID := uint16((uint32('A') % 65534) + 1)
	rightID := uint16((uint32('V') % 65534) + 1)
	kt.pairs[uint32(leftID)<<16|uint32(rightID)] = -80
	kt.hasKerning = true

	// Initial lookup
	kern := kt.GetKerning('A', 'V')
	if kern != -80 {
		t.Errorf(
			"GetKerning('A', 'V') = %d, want -80",
			kern,
		)
	}

	// The rune pair should now be cached
	runeKey := uint64('A')<<32 | uint64('V')
	if _, ok := kt.runePairs[runeKey]; !ok {
		t.Error(
			"Rune pair should be cached after first lookup",
		)
	}

	// Verify cached value is correct
	if kt.runePairs[runeKey] != -80 {
		t.Errorf(
			"Cached kerning value = %d, want -80",
			kt.runePairs[runeKey],
		)
	}
}

// TestValueRecordSize tests the GPOS value record size calculation.
func TestValueRecordSize(t *testing.T) {
	kt := &KerningTable{}

	tests := []struct {
		format uint16
		want   int
	}{
		{0x0000, 0}, // No values
		{0x0001, 2}, // XPlacement only
		{0x0002, 2}, // YPlacement only
		{0x0004, 2}, // XAdvance only
		{0x0008, 2}, // YAdvance only
		{0x0003, 4}, // XPlacement + YPlacement
		{
			0x0007,
			6,
		}, // XPlacement + YPlacement + XAdvance
		{0x000F, 8}, // All position values
		{
			0x00FF,
			16,
		}, // All values including device tables
	}

	for _, tt := range tests {
		got := kt.valueRecordSize(tt.format)
		if got != tt.want {
			t.Errorf(
				"valueRecordSize(0x%04X) = %d, want %d",
				tt.format,
				got,
				tt.want,
			)
		}
	}
}

// Benchmark tests

func BenchmarkGetKerning(b *testing.B) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)

	b.ResetTimer()
	for range b.N {
		kt.GetKerning('A', 'V')
	}
}

func BenchmarkGetKerningByGlyphID(b *testing.B) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)
	kt.pairs[uint32(1)<<16|uint32(2)] = -50
	kt.hasKerning = true

	b.ResetTimer()
	for range b.N {
		kt.GetKerningByGlyphID(1, 2)
	}
}

func BenchmarkMeasureStringWithKerning(
	b *testing.B,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)
	text := "AVATAR Warning"

	b.ResetTimer()
	for range b.N {
		kt.MeasureStringWithKerning(text)
	}
}

func BenchmarkMeasureStringWithKerningScaled(
	b *testing.B,
) {
	f := mockFontWithKerning()
	kt := NewKerningTable(f)
	text := "AVATAR Warning"

	b.ResetTimer()
	for range b.N {
		kt.MeasureStringWithKerningScaled(
			text,
			12.0,
		)
	}
}
