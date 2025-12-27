package font

import (
	"os"
	"path/filepath"
	"testing"
)

// TestNewSubset tests creating a new subset builder
func TestNewSubset(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
		},
	}

	builder := NewSubset(font)
	if builder == nil {
		t.Fatal("NewSubset returned nil")
	}

	if builder.font != font {
		t.Error(
			"Builder font reference is incorrect",
		)
	}

	if builder.usedRunes == nil {
		t.Error("Builder usedRunes map is nil")
	}
}

// TestSubsetBuilder_AddRune tests adding individual runes
func TestSubsetBuilder_AddRune(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
			'C': {AdvanceWidth: 700},
		},
	}

	builder := NewSubset(font)
	builder.AddRune('A')
	builder.AddRune('B')
	builder.AddRune(
		'A',
	) // Duplicate should not increase count

	if builder.UsedRuneCount() != 2 {
		t.Errorf(
			"UsedRuneCount = %d, want 2",
			builder.UsedRuneCount(),
		)
	}
}

// TestSubsetBuilder_AddString tests adding runes from a string
func TestSubsetBuilder_AddString(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'H': {AdvanceWidth: 700},
			'e': {AdvanceWidth: 500},
			'l': {AdvanceWidth: 300},
			'o': {AdvanceWidth: 550},
		},
	}

	builder := NewSubset(font)
	builder.AddString("Hello")

	// "Hello" has 4 unique characters: H, e, l, o
	if builder.UsedRuneCount() != 4 {
		t.Errorf(
			"UsedRuneCount = %d, want 4",
			builder.UsedRuneCount(),
		)
	}
}

// TestSubsetBuilder_AddRunes tests adding multiple runes at once
func TestSubsetBuilder_AddRunes(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
			'C': {AdvanceWidth: 700},
		},
	}

	builder := NewSubset(font)
	builder.AddRunes([]rune{'A', 'B', 'C'})

	if builder.UsedRuneCount() != 3 {
		t.Errorf(
			"UsedRuneCount = %d, want 3",
			builder.UsedRuneCount(),
		)
	}
}

// TestSubsetBuilder_UsedRunes tests getting sorted used runes
func TestSubsetBuilder_UsedRunes(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'C': {AdvanceWidth: 700},
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
		},
	}

	builder := NewSubset(font)
	builder.AddRunes([]rune{'C', 'A', 'B'})

	runes := builder.UsedRunes()
	if len(runes) != 3 {
		t.Fatalf(
			"UsedRunes length = %d, want 3",
			len(runes),
		)
	}

	// Should be sorted
	if runes[0] != 'A' || runes[1] != 'B' ||
		runes[2] != 'C' {
		t.Errorf(
			"UsedRunes not sorted: got %v",
			runes,
		)
	}
}

// TestSubsetBuilder_Build tests building a subset
func TestSubsetBuilder_Build(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'H': {AdvanceWidth: 700},
			'e': {AdvanceWidth: 500},
			'l': {AdvanceWidth: 300},
			'o': {AdvanceWidth: 550},
		},
		Data: buildMinimalFont(), // Use the minimal font from truetype_test.go
	}

	builder := NewSubset(font)
	builder.AddString("Hello")

	subset, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	if subset.Original != font {
		t.Error("Subset.Original is incorrect")
	}

	// 4 unique characters + 1 for .notdef = 5 glyphs total
	// But we only map characters the font actually has
	if subset.GlyphCount < 1 {
		t.Errorf(
			"GlyphCount = %d, want at least 1",
			subset.GlyphCount,
		)
	}

	if len(subset.Data) == 0 {
		t.Error("Subset.Data is empty")
	}
}

// TestSubsetBuilder_Build_NilFont tests building with nil font
func TestSubsetBuilder_Build_NilFont(
	t *testing.T,
) {
	builder := &SubsetBuilder{
		font:      nil,
		usedRunes: make(map[rune]bool),
	}
	builder.usedRunes['A'] = true

	_, err := builder.Build()
	if err == nil {
		t.Error("Expected error for nil font")
	}
}

// TestSubsetBuilder_Build_NoGlyphs tests building with no glyphs
func TestSubsetBuilder_Build_NoGlyphs(
	t *testing.T,
) {
	font := &Font{
		Family:    "Test",
		GlyphData: make(map[rune]GlyphMetrics),
		Data:      buildMinimalFont(),
	}

	builder := NewSubset(font)
	// Don't add any runes

	_, err := builder.Build()
	if err == nil {
		t.Error("Expected error for empty subset")
	}

	if err != ErrNoGlyphs {
		t.Errorf(
			"Expected ErrNoGlyphs, got %v",
			err,
		)
	}
}

// TestCreateSubset tests the convenience function
func TestCreateSubset(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'T': {AdvanceWidth: 600},
			'e': {AdvanceWidth: 500},
			's': {AdvanceWidth: 450},
			't': {AdvanceWidth: 400},
		},
		Data: buildMinimalFont(),
	}

	subset, err := CreateSubset(font, "Test")
	if err != nil {
		t.Fatalf("CreateSubset failed: %v", err)
	}

	if subset.Original != font {
		t.Error("Subset.Original is incorrect")
	}
}

// TestCreateSubset_NilFont tests CreateSubset with nil font
func TestCreateSubset_NilFont(t *testing.T) {
	_, err := CreateSubset(nil, "Test")
	if err == nil {
		t.Error("Expected error for nil font")
	}
}

// TestSubsetFont_GlyphID tests glyph ID lookup
func TestSubsetFont_GlyphID(t *testing.T) {
	subset := &SubsetFont{
		UsedGlyphs: map[rune]uint16{
			'A': 1,
			'B': 2,
			'C': 3,
		},
	}

	tests := []struct {
		r        rune
		expected uint16
	}{
		{'A', 1},
		{'B', 2},
		{'C', 3},
		{
			'Z',
			0,
		}, // Not in subset, returns .notdef
	}

	for _, tt := range tests {
		got := subset.GlyphID(tt.r)
		if got != tt.expected {
			t.Errorf(
				"GlyphID(%q) = %d, want %d",
				tt.r,
				got,
				tt.expected,
			)
		}
	}
}

// TestSubsetFont_HasRune tests rune presence check
func TestSubsetFont_HasRune(t *testing.T) {
	subset := &SubsetFont{
		UsedGlyphs: map[rune]uint16{
			'A': 1,
			'B': 2,
		},
	}

	if !subset.HasRune('A') {
		t.Error("HasRune('A') = false, want true")
	}

	if subset.HasRune('Z') {
		t.Error("HasRune('Z') = true, want false")
	}
}

// TestSubsetFont_Runes tests getting all runes in subset
func TestSubsetFont_Runes(t *testing.T) {
	subset := &SubsetFont{
		UsedGlyphs: map[rune]uint16{
			'C': 3,
			'A': 1,
			'B': 2,
		},
	}

	runes := subset.Runes()
	if len(runes) != 3 {
		t.Fatalf(
			"Runes length = %d, want 3",
			len(runes),
		)
	}

	// Should be sorted
	if runes[0] != 'A' || runes[1] != 'B' ||
		runes[2] != 'C' {
		t.Errorf(
			"Runes not sorted: got %v",
			runes,
		)
	}
}

// TestSubsetFont_EncodedText tests text encoding
func TestSubsetFont_EncodedText(t *testing.T) {
	subset := &SubsetFont{
		UsedGlyphs: map[rune]uint16{
			'H': 1,
			'i': 2,
		},
	}

	encoded := subset.EncodedText("Hi!")

	if len(encoded) != 3 {
		t.Fatalf(
			"EncodedText length = %d, want 3",
			len(encoded),
		)
	}

	// 'H' -> 1, 'i' -> 2, '!' -> 0 (not in subset)
	if encoded[0] != 1 || encoded[1] != 2 ||
		encoded[2] != 0 {
		t.Errorf(
			"EncodedText = %v, want [1, 2, 0]",
			encoded,
		)
	}
}

// TestSubsetFont_FontName tests font naming
func TestSubsetFont_FontName(t *testing.T) {
	font := &Font{Family: "Arial"}

	// Not a subset (full embedding)
	fullSubset := &SubsetFont{
		Original: font,
		IsSubset: false,
	}
	if name := fullSubset.FontName("ABCDEF"); name != "Arial" {
		t.Errorf(
			"FontName for full font = %q, want %q",
			name,
			"Arial",
		)
	}

	// Actual subset
	actualSubset := &SubsetFont{
		Original: font,
		IsSubset: true,
	}
	if name := actualSubset.FontName("XYZABC"); name != "XYZABC+Arial" {
		t.Errorf(
			"FontName for subset = %q, want %q",
			name,
			"XYZABC+Arial",
		)
	}

	// Short tag gets padded
	if name := actualSubset.FontName("AB"); name != "ABAAAA+Arial" {
		t.Errorf(
			"FontName with short tag = %q, want %q",
			name,
			"ABAAAA+Arial",
		)
	}
}

// TestCalculateTableChecksum tests checksum calculation
func TestCalculateTableChecksum(t *testing.T) {
	// Simple test: 4 bytes that add up to a known value
	data := []byte{0x00, 0x00, 0x00, 0x01}
	checksum := calculateTableChecksum(data)
	if checksum != 1 {
		t.Errorf(
			"calculateTableChecksum = %d, want 1",
			checksum,
		)
	}

	// Test padding
	data2 := []byte{0x00, 0x00, 0x00, 0x01, 0x00}
	checksum2 := calculateTableChecksum(data2)
	// Should be 1 + (0 padded to 4 bytes) = 1
	if checksum2 != 1 {
		t.Errorf(
			"calculateTableChecksum with padding = %d, want 1",
			checksum2,
		)
	}
}

// TestParseFontTables tests table parsing
func TestParseFontTables(t *testing.T) {
	fontData := buildMinimalFont()

	tables, err := parseFontTables(fontData)
	if err != nil {
		t.Fatalf(
			"parseFontTables failed: %v",
			err,
		)
	}

	// The minimal font from buildMinimalFont() contains these tables that are
	// fully within the data bounds. The name table is declared but extends
	// beyond the actual data, so it may not be parsed successfully.
	// We only require the essential tables for subsetting.
	requiredTables := []string{
		"head",
		"hhea",
		"maxp",
		"hmtx",
		"cmap",
	}
	for _, tag := range requiredTables {
		if _, ok := tables[tag]; !ok {
			t.Errorf(
				"Missing required table: %s",
				tag,
			)
		}
	}

	// Verify we got at least the required tables
	if len(tables) < len(requiredTables) {
		t.Errorf(
			"Expected at least %d tables, got %d",
			len(requiredTables),
			len(tables),
		)
	}
}

// TestParseFontTables_InvalidData tests parsing with invalid data
func TestParseFontTables_InvalidData(
	t *testing.T,
) {
	_, err := parseFontTables([]byte{0, 1, 2})
	if err == nil {
		t.Error("Expected error for invalid data")
	}
}

// TestGetNumGlyphsFromMaxp tests glyph count extraction
func TestGetNumGlyphsFromMaxp(t *testing.T) {
	// Valid maxp table header
	maxpData := []byte{
		0,
		0,
		0,
		0,
		0,
		10,
	} // version (4 bytes) + numGlyphs (2 bytes)
	count := getNumGlyphsFromMaxp(maxpData)
	if count != 10 {
		t.Errorf(
			"getNumGlyphsFromMaxp = %d, want 10",
			count,
		)
	}

	// Too short
	count2 := getNumGlyphsFromMaxp(
		[]byte{0, 0, 0},
	)
	if count2 != 0 {
		t.Errorf(
			"getNumGlyphsFromMaxp for short data = %d, want 0",
			count2,
		)
	}
}

// TestSubsetWithRealFont tests subsetting with a real font file if available
func TestSubsetWithRealFont(t *testing.T) {
	// Try to find a real font file for testing
	fontPaths := []string{
		"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
		"/usr/share/fonts/TTF/DejaVuSans.ttf",
		"/System/Library/Fonts/Helvetica.ttc",
	}

	var fontPath string
	for _, p := range fontPaths {
		if _, err := os.Stat(p); err == nil {
			fontPath = p

			break
		}
	}

	if fontPath == "" {
		t.Skip("No test font file available")
	}

	font, err := ParseFontFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse font %s: %v",
			fontPath,
			err,
		)
	}

	// Create a subset with some text
	subset, err := CreateSubset(
		font,
		"Hello, World!",
	)
	if err != nil {
		t.Fatalf("CreateSubset failed: %v", err)
	}

	// Verify subset properties
	if subset.Original != font {
		t.Error("Subset.Original is incorrect")
	}

	// Check that all characters are mapped
	for _, r := range "Hello, World!" {
		if _, ok := font.GlyphData[r]; ok {
			if !subset.HasRune(r) {
				t.Errorf(
					"Rune %q missing from subset",
					r,
				)
			}
		}
	}

	// Verify we have some data
	if len(subset.Data) == 0 {
		t.Error("Subset has no data")
	}

	t.Logf(
		"Created subset from %s: %d glyphs, %d bytes, isSubset=%v",
		filepath.Base(
			fontPath,
		),
		subset.GlyphCount,
		len(subset.Data),
		subset.IsSubset,
	)
}

// TestSubsetUnicode tests subsetting with Unicode characters
func TestSubsetUnicode(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			// Latin
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
			// Extended Latin
			'\u00E9': {
				AdvanceWidth: 500,
			}, // e with acute
			'\u00F1': {
				AdvanceWidth: 550,
			}, // n with tilde
			// Greek
			'\u03B1': {
				AdvanceWidth: 520,
			}, // alpha
			'\u03B2': {AdvanceWidth: 530}, // beta
		},
		Data: buildMinimalFont(),
	}

	builder := NewSubset(font)
	builder.AddString(
		"AB\u00E9\u00F1\u03B1\u03B2",
	)

	subset, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	// All 6 characters should be in the subset
	expectedRunes := []rune{
		'A',
		'B',
		'\u00E9',
		'\u00F1',
		'\u03B1',
		'\u03B2',
	}
	for _, r := range expectedRunes {
		if !subset.HasRune(r) {
			t.Errorf(
				"Rune %q (U+%04X) missing from subset",
				r,
				r,
			)
		}
	}
}

// TestSubsetEmptyString tests subsetting with empty string
func TestSubsetEmptyString(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
		},
		Data: buildMinimalFont(),
	}

	builder := NewSubset(font)
	builder.AddString("")

	_, err := builder.Build()
	if err != ErrNoGlyphs {
		t.Errorf(
			"Expected ErrNoGlyphs for empty string, got %v",
			err,
		)
	}
}

// TestSubsetMissingGlyphs tests subsetting when font is missing glyphs
func TestSubsetMissingGlyphs(t *testing.T) {
	font := &Font{
		Family: "Test",
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
		},
		Data: buildMinimalFont(),
	}

	// Request characters not in font
	subset, err := CreateSubset(font, "ABCXYZ")
	if err != nil {
		t.Fatalf("CreateSubset failed: %v", err)
	}

	// Only A and B should be in the subset
	if !subset.HasRune('A') ||
		!subset.HasRune('B') {
		t.Error(
			"Expected A and B to be in subset",
		)
	}

	// X, Y, Z should not be in subset (not in font)
	if subset.HasRune('X') ||
		subset.HasRune('Y') ||
		subset.HasRune('Z') {
		t.Error(
			"X, Y, Z should not be in subset (not in font)",
		)
	}

	// GlyphID should return 0 for missing characters
	if subset.GlyphID('X') != 0 {
		t.Error(
			"GlyphID for missing glyph should be 0",
		)
	}
}
