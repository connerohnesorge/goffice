package font

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"
)

// TestDetectFontFormat tests font format detection
func TestDetectFontFormat(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected FontFormat
	}{
		{
			name:     "empty data",
			data:     nil,
			expected: FormatUnknown,
		},
		{
			name:     "too short",
			data:     []byte{0x00, 0x01},
			expected: FormatUnknown,
		},
		{
			name: "TrueType signature",
			data: []byte{
				0x00,
				0x01,
				0x00,
				0x00,
			},
			expected: FormatTrueType,
		},
		{
			name:     "true signature",
			data:     []byte{'t', 'r', 'u', 'e'},
			expected: FormatTrueType,
		},
		{
			name:     "OTTO signature",
			data:     []byte{'O', 'T', 'T', 'O'},
			expected: FormatOpenTypeCFF,
		},
		{
			name:     "TTC signature",
			data:     []byte{'t', 't', 'c', 'f'},
			expected: FormatTrueTypeCollection,
		},
		{
			name: "unknown signature",
			data: []byte{
				0xDE,
				0xAD,
				0xBE,
				0xEF,
			},
			expected: FormatUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectFontFormat(tt.data)
			if got != tt.expected {
				t.Errorf(
					"DetectFontFormat() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestFontFormat_String tests FontFormat string representation
func TestFontFormat_String(t *testing.T) {
	tests := []struct {
		format   FontFormat
		expected string
	}{
		{FormatUnknown, "Unknown"},
		{FormatTrueType, "TrueType"},
		{
			FormatOpenTypeTTF,
			"OpenType (TrueType outlines)",
		},
		{
			FormatOpenTypeCFF,
			"OpenType (CFF outlines)",
		},
		{
			FormatTrueTypeCollection,
			"TrueType Collection",
		},
		{FontFormat(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.format.String(); got != tt.expected {
				t.Errorf(
					"FontFormat.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestParseOpenType_InvalidData tests error handling for invalid data
func TestParseOpenType_InvalidData(t *testing.T) {
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
				0x00,
			},
			wantErr: true,
		},
		{
			name:    "invalid signature",
			data:    makeInvalidOTFont(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseOpenType(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ParseOpenType() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

// makeInvalidOTFont creates a font with an invalid magic number
func makeInvalidOTFont() []byte {
	buf := &bytes.Buffer{}
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0xDEADBEEF),
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
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	return buf.Bytes()
}

// TestParseOpenTypeFile_NonExistent tests error handling for non-existent files
func TestParseOpenTypeFile_NonExistent(
	t *testing.T,
) {
	_, err := ParseOpenTypeFile(
		"/nonexistent/path/to/font.otf",
	)
	if err == nil {
		t.Error(
			"ParseOpenTypeFile() expected error for non-existent file",
		)
	}
}

// TestParseFontFile_NonExistent tests error handling for non-existent files
func TestParseFontFile_NonExistent(t *testing.T) {
	_, err := ParseFontFile(
		"/nonexistent/path/to/font.otf",
	)
	if err == nil {
		t.Error(
			"ParseFontFile() expected error for non-existent file",
		)
	}
}

// TestParseFont_TrueType tests that ParseFont correctly delegates to TrueType parser
func TestParseFont_TrueType(t *testing.T) {
	// Use the minimal font builder from truetype_test.go
	data := buildMinimalFont()

	font, err := ParseFont(data)
	if err != nil {
		t.Fatalf("ParseFont() failed: %v", err)
	}

	if font.Metrics.UnitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			font.Metrics.UnitsPerEm,
		)
	}
}

// TestParseOpenType_TrueType tests that ParseOpenType handles TrueType fonts
func TestParseOpenType_TrueType(t *testing.T) {
	data := buildMinimalFont()

	font, err := ParseOpenType(data)
	if err != nil {
		t.Fatalf(
			"ParseOpenType() failed: %v",
			err,
		)
	}

	if font.Metrics.UnitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			font.Metrics.UnitsPerEm,
		)
	}
}

// TestParseOpenType_CFF tests parsing of a minimal CFF-based OpenType font
func TestParseOpenType_CFF(t *testing.T) {
	data := buildMinimalCFFFont()

	font, err := ParseOpenType(data)
	if err != nil {
		t.Fatalf(
			"ParseOpenType() failed for CFF font: %v",
			err,
		)
	}

	if font.Metrics.UnitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			font.Metrics.UnitsPerEm,
		)
	}
}

// buildMinimalCFFFont creates a minimal OpenType/CFF font for testing
func buildMinimalCFFFont() []byte {
	buf := &bytes.Buffer{}

	// Calculate offsets
	tableStart := uint32(
		12 + 7*16,
	) // header + 7 table records

	// Table sizes
	headSize := uint32(54)
	hheaSize := uint32(36)
	maxpSize := uint32(6)
	hmtxSize := uint32(4)
	cmapSize := uint32(4 + 8 + 32)
	nameSize := uint32(6 + 12 + 8)
	cffSize := uint32(32) // Minimal CFF table

	// Calculate offsets
	headOffset := tableStart
	hheaOffset := headOffset + headSize
	maxpOffset := hheaOffset + hheaSize
	hmtxOffset := maxpOffset + maxpSize
	cmapOffset := hmtxOffset + hmtxSize
	nameOffset := cmapOffset + cmapSize
	cffOffset := nameOffset + nameSize

	// Write header with OTTO signature
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x4F54544F),
	) // 'OTTO'
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(7),
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
		uint16(48),
	) // rangeShift

	// Write table records (must be sorted alphabetically by tag)
	writeTableRecord(
		buf,
		"CFF ",
		0,
		cffOffset,
		cffSize,
	)
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

	// Write maxp table (CFF version)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x00005000),
	) // version (0.5 for CFF)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // numGlyphs

	// Write hmtx table
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

	// Write cmap table
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
	)
	// Format 4 subtable
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

	// Write name table
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

	// Write minimal CFF table
	// CFF Header
	buf.WriteByte(1) // major version
	buf.WriteByte(0) // minor version
	buf.WriteByte(4) // header size
	buf.WriteByte(1) // offSize

	// Name INDEX
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // count
	buf.WriteByte(
		1,
	) // offSize
	buf.WriteByte(
		1,
	) // offset[0]
	buf.WriteByte(
		5,
	) // offset[1] (name length = 4)
	buf.Write(
		[]byte("Test"),
	) // name data

	// Top DICT INDEX (minimal)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // count
	buf.WriteByte(
		1,
	) // offSize
	buf.WriteByte(
		1,
	) // offset[0]
	buf.WriteByte(
		1,
	) // offset[1] (empty dict)

	// String INDEX (empty)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	// Global Subr INDEX (empty)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	return buf.Bytes()
}

// TestParseCFFInfo tests the CFF info parser
func TestParseCFFInfo(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		wantErr  bool
		wantName string
	}{
		{
			name:    "empty data",
			data:    nil,
			wantErr: true,
		},
		{
			name:    "too short",
			data:    []byte{1, 0},
			wantErr: true,
		},
		{
			name:     "valid CFF with name",
			data:     buildMinimalCFFData(),
			wantErr:  false,
			wantName: "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			font := &Font{}
			err := parseCFFInfo(tt.data, font)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"parseCFFInfo() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
			if !tt.wantErr && tt.wantName != "" &&
				font.Family != tt.wantName {
				t.Errorf(
					"parseCFFInfo() font.Family = %v, want %v",
					font.Family,
					tt.wantName,
				)
			}
		})
	}
}

// buildMinimalCFFData creates minimal CFF table data for testing
func buildMinimalCFFData() []byte {
	buf := &bytes.Buffer{}

	// CFF Header
	buf.WriteByte(1) // major version
	buf.WriteByte(0) // minor version
	buf.WriteByte(4) // header size
	buf.WriteByte(1) // offSize

	// Name INDEX
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // count
	buf.WriteByte(
		1,
	) // offSize
	buf.WriteByte(
		1,
	) // offset[0]
	buf.WriteByte(
		5,
	) // offset[1]
	buf.Write([]byte("Test"))

	return buf.Bytes()
}

// TestReadCFFOffset tests variable-length offset reading
func TestReadCFFOffset(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		offSize  int
		expected uint32
	}{
		{
			name:     "1 byte offset",
			data:     []byte{0x42},
			offSize:  1,
			expected: 0x42,
		},
		{
			name:     "2 byte offset",
			data:     []byte{0x01, 0x23},
			offSize:  2,
			expected: 0x0123,
		},
		{
			name:     "3 byte offset",
			data:     []byte{0x01, 0x23, 0x45},
			offSize:  3,
			expected: 0x012345,
		},
		{
			name: "4 byte offset",
			data: []byte{
				0x01,
				0x23,
				0x45,
				0x67,
			},
			offSize:  4,
			expected: 0x01234567,
		},
		{
			name:     "empty data",
			data:     nil,
			offSize:  1,
			expected: 0,
		},
		{
			name:     "data too short",
			data:     []byte{0x01},
			offSize:  2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readCFFOffset(
				tt.data,
				tt.offSize,
			)
			if got != tt.expected {
				t.Errorf(
					"readCFFOffset() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

// TestParseTrueTypeCollection tests TTC parsing
func TestParseTrueTypeCollection(t *testing.T) {
	// Create a minimal TTC with one font that has adjusted offsets
	data := buildMinimalTTC()

	fonts, err := ParseFontCollection(data)
	if err != nil {
		t.Fatalf(
			"ParseFontCollection() failed: %v",
			err,
		)
	}

	if len(fonts) != 1 {
		t.Errorf(
			"ParseFontCollection() returned %d fonts, want 1",
			len(fonts),
		)
	}

	if fonts[0].Metrics.UnitsPerEm != 1000 {
		t.Errorf(
			"UnitsPerEm = %d, want 1000",
			fonts[0].Metrics.UnitsPerEm,
		)
	}
}

// buildMinimalTTC creates a minimal TrueType Collection for testing
// In a TTC, the embedded fonts' table offsets are absolute (from TTC file start)
func buildMinimalTTC() []byte {
	buf := &bytes.Buffer{}

	// TTC Header (12 bytes)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(0x74746366),
	) // 'ttcf'
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(2),
	) // major version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // minor version
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(1),
	) // numFonts

	// Offset to first font header (TTC header 12 + offset table 4 = 16)
	ttcHeaderSize := uint32(16)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		ttcHeaderSize,
	)

	// Now build the font at this offset
	// Font header starts at offset 16
	fontHeaderStart := ttcHeaderSize

	// Calculate table offsets (they must be absolute from file start)
	// Font header: 12 bytes
	// 6 table records: 6 * 16 = 96 bytes
	tableStart := fontHeaderStart + uint32(
		12+6*16,
	)

	// Table sizes
	headSize := uint32(54)
	hheaSize := uint32(36)
	maxpSize := uint32(6)
	hmtxSize := uint32(4) // 1 metric
	cmapSize := uint32(4 + 8 + 32)
	nameSize := uint32(6 + 12 + 8)

	// Calculate absolute offsets
	headOffset := tableStart
	hheaOffset := headOffset + headSize
	maxpOffset := hheaOffset + hheaSize
	hmtxOffset := maxpOffset + maxpSize
	cmapOffset := hmtxOffset + hmtxSize
	nameOffset := cmapOffset + cmapSize

	// Write font header
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

	// Write table records with absolute offsets
	writeTTCTableRecord(
		buf,
		"cmap",
		0,
		cmapOffset,
		cmapSize,
	)
	writeTTCTableRecord(
		buf,
		"head",
		0,
		headOffset,
		headSize,
	)
	writeTTCTableRecord(
		buf,
		"hhea",
		0,
		hheaOffset,
		hheaSize,
	)
	writeTTCTableRecord(
		buf,
		"hmtx",
		0,
		hmtxOffset,
		hmtxSize,
	)
	writeTTCTableRecord(
		buf,
		"maxp",
		0,
		maxpOffset,
		maxpSize,
	)
	writeTTCTableRecord(
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

	// Write name table (minimal)
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

func writeTTCTableRecord(
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

// TestParseFontCollection_SingleFont tests that single fonts are handled correctly
func TestParseFontCollection_SingleFont(
	t *testing.T,
) {
	data := buildMinimalFont()

	fonts, err := ParseFontCollection(data)
	if err != nil {
		t.Fatalf(
			"ParseFontCollection() failed: %v",
			err,
		)
	}

	if len(fonts) != 1 {
		t.Errorf(
			"ParseFontCollection() returned %d fonts, want 1",
			len(fonts),
		)
	}
}

// TestParseFontCollectionFile_NonExistent tests error handling
func TestParseFontCollectionFile_NonExistent(
	t *testing.T,
) {
	_, err := ParseFontCollectionFile(
		"/nonexistent/path/to/font.ttc",
	)
	if err == nil {
		t.Error(
			"ParseFontCollectionFile() expected error for non-existent file",
		)
	}
}

// TestParseTrueTypeCollection_IndexOutOfRange tests error handling
func TestParseTrueTypeCollection_IndexOutOfRange(
	t *testing.T,
) {
	data := buildMinimalTTC()

	_, err := parseTrueTypeCollection(
		data,
		5,
	) // Invalid index
	if err == nil {
		t.Error(
			"parseTrueTypeCollection() expected error for invalid index",
		)
	}
}

// TestParseFont_RealOpenTypeFont tests with a real OpenType font if available
func TestParseFont_RealOpenTypeFont(
	t *testing.T,
) {
	// Try to find a real OpenType font for testing
	fontPaths := []string{
		// Common system font locations with OpenType fonts
		"/usr/share/fonts/opentype/",
		"/usr/share/fonts/OTF/",
		"/System/Library/Fonts/",
	}

	var otfPath string
	for _, dir := range fontPaths {
		if _, err := os.Stat(dir); os.IsNotExist(
			err,
		) {
			continue
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if filepath.Ext(
				entry.Name(),
			) == ".otf" {
				otfPath = filepath.Join(
					dir,
					entry.Name(),
				)

				break
			}
		}
		if otfPath != "" {
			break
		}
	}

	if otfPath == "" {
		t.Skip(
			"No OpenType font file available for testing",
		)
	}

	font, err := ParseFontFile(otfPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse OpenType font %s: %v",
			otfPath,
			err,
		)
	}

	if font.Metrics.UnitsPerEm == 0 {
		t.Error("UnitsPerEm should not be 0")
	}

	t.Logf(
		"Parsed OpenType font: %s, UnitsPerEm=%d, Glyphs=%d",
		font.Family,
		font.Metrics.UnitsPerEm,
		len(font.GlyphData),
	)
}

// TestParseFont_RealTTCFont tests with a real TTC font if available
func TestParseFont_RealTTCFont(t *testing.T) {
	// Try to find a real TTC font for testing
	fontPaths := []string{
		"/System/Library/Fonts/Helvetica.ttc",
		"/System/Library/Fonts/Times.ttc",
		"/usr/share/fonts/truetype/msttcorefonts/",
	}

	var ttcPath string
	for _, p := range fontPaths {
		if info, err := os.Stat(p); err == nil {
			if info.IsDir() {
				entries, err := os.ReadDir(p)
				if err != nil {
					continue
				}
				for _, entry := range entries {
					if filepath.Ext(
						entry.Name(),
					) == ".ttc" {
						ttcPath = filepath.Join(
							p,
							entry.Name(),
						)

						break
					}
				}
			} else if filepath.Ext(p) == ".ttc" {
				ttcPath = p
			}
		}
		if ttcPath != "" {
			break
		}
	}

	if ttcPath == "" {
		t.Skip(
			"No TTC font file available for testing",
		)
	}

	fonts, err := ParseFontCollectionFile(ttcPath)
	if err != nil {
		t.Fatalf(
			"Failed to parse TTC font %s: %v",
			ttcPath,
			err,
		)
	}

	if len(fonts) == 0 {
		t.Error(
			"Expected at least one font in collection",
		)
	}

	for i, font := range fonts {
		t.Logf(
			"Font %d: %s, UnitsPerEm=%d",
			i,
			font.Family,
			font.Metrics.UnitsPerEm,
		)
	}
}
