package font

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"
)

// TestParseTrueType_InvalidData tests error handling for invalid font data
func TestParseTrueType_InvalidData(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty data",
			data:    nil,
			wantErr: true,
		},
		{
			name: "too short",
			data: []byte{
				0x00,
				0x01,
				0x00,
				0x00,
			},
			wantErr: true,
		},
		{
			name:    "invalid magic number",
			data:    makeInvalidFont(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTrueType(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ParseTrueType() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

// makeInvalidFont creates a font with invalid magic number
func makeInvalidFont() []byte {
	buf := &bytes.Buffer{}
	// Write invalid sfnt version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0xDEADBEEF),
	)
	// Write numTables, searchRange, entrySelector, rangeShift
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
		uint16(0),
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	return buf.Bytes()
}

// TestParseTrueTypeFile_NonExistent tests error handling for non-existent files
func TestParseTrueTypeFile_NonExistent(
	t *testing.T,
) {
	_, err := ParseTrueTypeFile(
		"/nonexistent/path/to/font.ttf",
	)
	if err == nil {
		t.Error(
			"ParseTrueTypeFile() expected error for non-existent file",
		)
	}
}

// TestFontStyle_String tests the FontStyle String method
func TestFontStyle_String(t *testing.T) {
	tests := []struct {
		style    FontStyle
		expected string
	}{
		{StyleRegular, "Regular"},
		{StyleBold, "Bold"},
		{StyleItalic, "Italic"},
		{StyleBoldItalic, "Bold Italic"},
		{FontStyle(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.style.String(); got != tt.expected {
				t.Errorf(
					"FontStyle.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFont_TextWidth tests the text width calculation
func TestFont_TextWidth(t *testing.T) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
		GlyphData: map[rune]GlyphMetrics{
			'H': {AdvanceWidth: 700},
			'e': {AdvanceWidth: 500},
			'l': {AdvanceWidth: 300},
			'o': {AdvanceWidth: 550},
		},
	}

	// "Hello" = H(700) + e(500) + l(300) + l(300) + o(550) = 2350
	width := font.TextWidth("Hello")
	expected := 2350
	if width != expected {
		t.Errorf(
			"TextWidth(\"Hello\") = %d, want %d",
			width,
			expected,
		)
	}
}

// TestFont_GlyphWidth tests glyph width lookup
func TestFont_GlyphWidth(t *testing.T) {
	font := &Font{
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 600},
			'B': {AdvanceWidth: 650},
		},
	}

	if w := font.GlyphWidth('A'); w != 600 {
		t.Errorf(
			"GlyphWidth('A') = %d, want 600",
			w,
		)
	}

	if w := font.GlyphWidth('Z'); w != 0 {
		t.Errorf(
			"GlyphWidth('Z') = %d, want 0 (not found)",
			w,
		)
	}
}

// TestFont_ScaledWidth_Basic tests basic unit conversion scenarios
func TestFont_ScaledWidth_Basic(t *testing.T) {
	font := &Font{
		Metrics: FontMetrics{UnitsPerEm: 1000},
	}

	// 500 font units at 12pt with 1000 unitsPerEm = 6pt
	scaled := font.ScaledWidth(500, 12.0)
	expected := 6.0
	if scaled != expected {
		t.Errorf(
			"ScaledWidth(500, 12.0) = %f, want %f",
			scaled,
			expected,
		)
	}

	// Test with zero UnitsPerEm
	font.Metrics.UnitsPerEm = 0
	scaled = font.ScaledWidth(500, 12.0)
	if scaled != 0 {
		t.Errorf(
			"ScaledWidth with zero UnitsPerEm = %f, want 0",
			scaled,
		)
	}
}

// TestDecodeUTF16BE tests UTF-16BE decoding
func TestDecodeUTF16BE(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name: "simple ASCII",
			input: []byte{
				0x00,
				'H',
				0x00,
				'i',
			},
			expected: "Hi",
		},
		{
			name:     "empty",
			input:    nil,
			expected: "",
		},
		{
			name:     "odd length",
			input:    []byte{0x00},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeUTF16BE(tt.input); got != tt.expected {
				t.Errorf(
					"decodeUTF16BE() = %q, want %q",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestParseTrueType_MinimalFont tests parsing with a minimal constructed font
func TestParseTrueType_MinimalFont(t *testing.T) {
	// This test verifies the parser handles missing optional tables gracefully
	font := buildMinimalFont()

	parsed, err := ParseTrueType(font)
	if err != nil {
		t.Fatalf(
			"Failed to parse minimal font: %v",
			err,
		)
	}

	if parsed.Metrics.UnitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			parsed.Metrics.UnitsPerEm,
		)
	}
}

// buildMinimalFont constructs a minimal valid TrueType font for testing
func buildMinimalFont() []byte {
	buf := &bytes.Buffer{}

	// Calculate offsets
	// Header: 12 bytes
	// 6 table records: 6 * 16 = 96 bytes
	// Total directory: 108 bytes
	tableStart := uint32(12 + 6*16)

	// Table sizes
	headSize := uint32(54)
	hheaSize := uint32(36)
	maxpSize := uint32(6)
	hmtxSize := uint32(4) // 1 metric
	// cmap: header (4) + encoding record (8) + format 4 subtable (32 minimum for 1 segment)
	cmapSize := uint32(4 + 8 + 32)
	nameSize := uint32(
		6 + 12 + 8,
	) // header(6) + 1 record(12) + string data(8)

	// Calculate offsets for each table
	headOffset := tableStart
	hheaOffset := headOffset + headSize
	maxpOffset := hheaOffset + hheaSize
	hmtxOffset := maxpOffset + maxpSize
	cmapOffset := hmtxOffset + hmtxSize
	nameOffset := cmapOffset + cmapSize

	// Write header
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	) // sfnt version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(6),
	) // numTables
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(64),
	) // searchRange
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	) // entrySelector
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	) // rangeShift

	// Write table records
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
		"maxp",
		0,
		maxpOffset,
		maxpSize,
	)
	writeTableRecord(
		buf,
		"name",
		0,
		nameOffset,
		nameSize,
	)

	// Write head table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	) // version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	) // fontRevision
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0),
	) // checksumAdjustment
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x5F0F3CF5),
	) // magicNumber
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // flags
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1000),
	) // unitsPerEm
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	) // created
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int64(0),
	) // modified
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // xMin
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // yMin
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1000),
	) // xMax
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1000),
	) // yMax
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // macStyle
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	) // lowestRecPPEM
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(2),
	) // fontDirectionHint
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // indexToLocFormat
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // glyphDataFormat

	// Write hhea table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00010000),
	) // version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(800),
	) // ascender
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(-200),
	) // descender
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(90),
	) // lineGap
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(600),
	) // advanceWidthMax
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // minLeftSideBearing
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // minRightSideBearing
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(600),
	) // xMaxExtent
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	) // caretSlopeRise
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // caretSlopeRun
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // caretOffset
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // reserved1
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // reserved2
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // reserved3
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // reserved4
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // metricDataFormat
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // numberOfHMetrics

	// Write maxp table
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00005000),
	) // version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // numGlyphs

	// Write hmtx table (1 glyph)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(600),
	) // advanceWidth
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(0),
	) // lsb

	// Write cmap table with format 4 subtable
	// cmap header
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // numTables
	// Encoding record
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	) // platformID (Windows)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // encodingID (Unicode BMP)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(12),
	) // offset to subtable (4 + 8 = 12)

	// Format 4 subtable (1 segment mapping 'A' to glyph 0)
	// segCount = 1 (just the terminating 0xFFFF segment)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(4),
	) // format
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(32),
	) // length of subtable
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // language
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	) // segCountX2 (1 segment * 2)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	) // searchRange
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // entrySelector
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // rangeShift
	// endCode array
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	) // segment 0 end
	// reservedPad
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)
	// startCode array
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0xFFFF),
	) // segment 0 start
	// idDelta array
	_ = binary.Write(
		buf,
		binary.BigEndian,
		int16(1),
	) // segment 0 delta
	// idRangeOffset array
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // segment 0 range offset

	// Write name table (minimal)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // format
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // count
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(18),
	) // stringOffset (6 + 12 = 18)
	// Name record
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(3),
	) // platformID (Windows)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // encodingID (Unicode BMP)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0x0409),
	) // languageID (English US)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // nameID (font family)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(8),
	) // length
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // offset
	// String data "Test" in UTF-16BE
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

func writeTableRecord(
	buf *bytes.Buffer,
	tag string,
	checksum, offset, length uint32,
) {
	buf.Write([]byte(tag))
	_ = binary.Write(
		buf,
		binary.BigEndian,
		checksum,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		offset,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		length,
	)
}

// TestParseTrueType_RealFont tests parsing with a real font file if available
func TestParseTrueType_RealFont(t *testing.T) {
	// Try to find a real font file for testing
	fontPaths := []string{
		// Go font from x/image module
		filepath.Join(
			os.Getenv("GOPATH"),
			"pkg/mod/golang.org/x/image@v0.21.0/font/gofont/ttfs/Go-Regular.ttf",
		),
		// System fonts
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

	font, err := ParseTrueTypeFile(fontPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse font %s: %v",
			fontPath,
			err,
		)
	}

	// Verify basic metrics are populated
	if font.Metrics.UnitsPerEm == 0 {
		t.Error("UnitsPerEm should not be 0")
	}

	if font.Metrics.Ascender == 0 {
		t.Error("Ascender should not be 0")
	}

	if font.Metrics.Descender == 0 {
		t.Log(
			"Note: Descender is 0 (may be expected for some fonts)",
		)
	}

	// Verify some glyphs are mapped
	if len(font.GlyphData) == 0 {
		t.Error("No glyphs were mapped")
	}

	// Test that common ASCII characters have widths
	for _, r := range "AaBbCc123" {
		if font.GlyphWidth(r) == 0 {
			t.Logf(
				"Warning: glyph %q has zero width",
				r,
			)
		}
	}

	t.Logf(
		"Parsed font: %s (%s), UnitsPerEm=%d, Ascender=%d, Descender=%d, Glyphs=%d",
		font.Family,
		font.Style,
		font.Metrics.UnitsPerEm,
		font.Metrics.Ascender,
		font.Metrics.Descender,
		len(font.GlyphData),
	)
}

// TestGetTableData tests the table data extraction helper
func TestGetTableData(t *testing.T) {
	data := []byte("0123456789ABCDEF")
	tables := map[string]tableRecord{
		"test": {Offset: 4, Length: 6},
		"bad": {
			Offset: 100,
			Length: 10,
		}, // Invalid offset
	}

	// Valid table
	result, err := getTableData(
		data,
		tables,
		"test",
	)
	if err != nil {
		t.Errorf(
			"getTableData for valid table failed: %v",
			err,
		)
	}
	if string(result) != "456789" {
		t.Errorf(
			"getTableData returned %q, want %q",
			result,
			"456789",
		)
	}

	// Missing table
	_, err = getTableData(data, tables, "missing")
	if err == nil {
		t.Error(
			"getTableData should fail for missing table",
		)
	}

	// Invalid bounds
	_, err = getTableData(data, tables, "bad")
	if err == nil {
		t.Error(
			"getTableData should fail for invalid bounds",
		)
	}
}

// TestKerningTable tests the kerning table functionality
func TestKerningTable(t *testing.T) {
	t.Run("GetKerning", func(t *testing.T) {
		kern := &KerningTable{
			Pairs: map[uint32]int16{
				(uint32(65) << 16) | uint32(86): -100, // A-V kerning
				(uint32(84) << 16) | uint32(111): -80, // T-o kerning
			},
		}

		// Test existing kerning pairs
		if v := kern.GetKerning(65, 86); v != -100 {
			t.Errorf("GetKerning(A, V) = %d, want -100", v)
		}
		if v := kern.GetKerning(84, 111); v != -80 {
			t.Errorf("GetKerning(T, o) = %d, want -80", v)
		}

		// Test non-existing pair
		if v := kern.GetKerning(65, 65); v != 0 {
			t.Errorf("GetKerning(A, A) = %d, want 0", v)
		}
	})

	t.Run("AddKerning", func(t *testing.T) {
		kern := &KerningTable{}
		kern.AddKerning(65, 66, -50)

		if v := kern.GetKerning(65, 66); v != -50 {
			t.Errorf("GetKerning after AddKerning = %d, want -50", v)
		}
	})

	t.Run("NilKerningTable", func(t *testing.T) {
		var kern *KerningTable
		if v := kern.GetKerning(65, 66); v != 0 {
			t.Errorf("GetKerning on nil table = %d, want 0", v)
		}
	})
}

// TestParseKernTable tests parsing of kern tables
func TestParseKernTable(t *testing.T) {
	t.Run("Version0Format0", func(t *testing.T) {
		// Create a minimal kern table (version 0, format 0)
		buf := &bytes.Buffer{}

		// kern table header (version 0)
		_ = binary.Write(buf, binary.BigEndian, uint16(0)) // version
		_ = binary.Write(buf, binary.BigEndian, uint16(1)) // nTables

		// Subtable header
		_ = binary.Write(buf, binary.BigEndian, uint16(18))   // length
		_ = binary.Write(buf, binary.BigEndian, uint16(0x0001)) // coverage: horizontal=1
		_ = binary.Write(buf, binary.BigEndian, uint16(0))    // tupleIndex

		// Format 0 subtable header
		_ = binary.Write(buf, binary.BigEndian, uint16(2)) // nPairs
		_ = binary.Write(buf, binary.BigEndian, uint16(4)) // searchRange
		_ = binary.Write(buf, binary.BigEndian, uint16(1)) // entrySelector
		_ = binary.Write(buf, binary.BigEndian, uint16(0)) // rangeShift

		// Kerning pairs
		_ = binary.Write(buf, binary.BigEndian, uint16(65))  // left glyph (A)
		_ = binary.Write(buf, binary.BigEndian, uint16(86))  // right glyph (V)
		_ = binary.Write(buf, binary.BigEndian, int16(-100)) // kerning value

		_ = binary.Write(buf, binary.BigEndian, uint16(84))  // left glyph (T)
		_ = binary.Write(buf, binary.BigEndian, uint16(111)) // right glyph (o)
		_ = binary.Write(buf, binary.BigEndian, int16(-80))  // kerning value

		kernTable := parseKernTable(buf.Bytes(), map[string]tableRecord{
			"kern": {Offset: 0, Length: uint32(buf.Len())},
		})

		if kernTable == nil {
			t.Fatal("parseKernTable returned nil")
		}

		if len(kernTable.Pairs) != 2 {
			t.Errorf("Expected 2 kerning pairs, got %d", len(kernTable.Pairs))
		}

		// Verify kerning values
		key1 := (uint32(65) << 16) | uint32(86)
		if v, ok := kernTable.Pairs[key1]; !ok || v != -100 {
			t.Errorf("Kerning pair A-V = %d, want -100", v)
		}
	})

	t.Run("MissingKernTable", func(t *testing.T) {
		kernTable := parseKernTable([]byte{}, map[string]tableRecord{})
		if kernTable != nil {
			t.Error("parseKernTable should return nil for missing table")
		}
	})

	t.Run("EmptyKernTable", func(t *testing.T) {
		buf := &bytes.Buffer{}
		_ = binary.Write(buf, binary.BigEndian, uint16(0)) // version
		_ = binary.Write(buf, binary.BigEndian, uint16(0)) // nTables = 0

		kernTable := parseKernTable(buf.Bytes(), map[string]tableRecord{
			"kern": {Offset: 0, Length: uint32(buf.Len())},
		})

		if kernTable != nil {
			t.Error("parseKernTable should return nil for empty table")
		}
	})
}

// TestFontGetKerningForRunes tests font-level kerning lookup
func TestFontGetKerningForRunes(t *testing.T) {
	font := &Font{
		Kerning: &KerningTable{
			Pairs: map[uint32]int16{
				(uint32(65) << 16) | uint32(86): -100, // A-V
			},
		},
		glyphToRune: map[uint16]rune{
			65: 'A',
			86: 'V',
		},
	}

	// Test with glyph mapping
	if v := font.GetKerningForRunes('A', 'V'); v != -100 {
		t.Errorf("GetKerningForRunes(A, V) = %d, want -100", v)
	}

	// Test without glyph mapping (should return 0)
	if v := font.GetKerningForRunes('X', 'Y'); v != 0 {
		t.Errorf("GetKerningForRunes(X, Y) = %d, want 0", v)
	}
}

// TestFontTextWidthWithKerning tests text width calculation with kerning
func TestFontTextWidthWithKerning(t *testing.T) {
	font := &Font{
		GlyphData: map[rune]GlyphMetrics{
			'A': {AdvanceWidth: 500},
			'V': {AdvanceWidth: 500},
			'B': {AdvanceWidth: 500},
		},
		Kerning: &KerningTable{
			Pairs: map[uint32]int16{
				(uint32(65) << 16) | uint32(86): -100, // A-V kerning
			},
		},
		glyphToRune: map[uint16]rune{
			65: 'A',
			86: 'V',
			66: 'B',
		},
	}

	// Test without kerning (single character)
	width := font.TextWidthWithKerning("A")
	if width != 500 {
		t.Errorf("TextWidthWithKerning(\"A\") = %d, want 500", width)
	}

	// Test with kerning pair
	width = font.TextWidthWithKerning("AV")
	expected := 500 + 500 - 100 // 900
	if width != expected {
		t.Errorf("TextWidthWithKerning(\"AV\") = %d, want %d", width, expected)
	}

	// Test without kerning (no kerning pair)
	width = font.TextWidthWithKerning("AB")
	expected = 500 + 500 // 1000
	if width != expected {
		t.Errorf("TextWidthWithKerning(\"AB\") = %d, want %d", width, expected)
	}

	// Test with nil kerning (should fall back to regular TextWidth)
	font.Kerning = nil
	width = font.TextWidthWithKerning("AV")
	expected = 1000
	if width != expected {
		t.Errorf("TextWidthWithKerning(nil kerning) = %d, want %d", width, expected)
	}
}
