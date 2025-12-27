// opentype.go provides OpenType font parsing with support for both TrueType and CFF outlines.
// OpenType fonts can use either TrueType outlines (glyf table) or CFF outlines (CFF table).
// This file handles the detection and appropriate parsing of both variants.

package font

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// Font format detection constants
const (
	// signatureOTTO indicates OpenType with CFF outlines
	signatureOTTO uint32 = 0x4F54544F // 'OTTO'
	// signatureTTF indicates TrueType outlines (version 1.0)
	signatureTTF uint32 = 0x00010000
	// signatureTrue indicates TrueType outlines ('true')
	signatureTrue uint32 = 0x74727565 // 'true'
	// signatureTTC indicates TrueType Collection
	signatureTTC uint32 = 0x74746366 // 'ttcf'
)

// CFF table tag
const tagCFF = "CFF "

// Error definitions for OpenType parsing
var (
	ErrNotOpenType = errors.New(
		"not an OpenType font",
	)
	ErrCFFParseFailed = errors.New(
		"failed to parse CFF table",
	)
)

// FontFormat represents the detected format of a font file.
type FontFormat int

const (
	// FormatUnknown indicates the font format could not be determined
	FormatUnknown FontFormat = iota
	// FormatTrueType indicates a TrueType font (.ttf)
	FormatTrueType
	// FormatOpenTypeTTF indicates an OpenType font with TrueType outlines
	FormatOpenTypeTTF
	// FormatOpenTypeCFF indicates an OpenType font with CFF outlines
	FormatOpenTypeCFF
	// FormatTrueTypeCollection indicates a TrueType Collection (.ttc)
	FormatTrueTypeCollection
)

// String returns the human-readable name of the font format.
func (f FontFormat) String() string {
	switch f {
	case FormatTrueType:
		return "TrueType"
	case FormatOpenTypeTTF:
		return "OpenType (TrueType outlines)"
	case FormatOpenTypeCFF:
		return "OpenType (CFF outlines)"
	case FormatTrueTypeCollection:
		return "TrueType Collection"
	case FormatUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

// DetectFontFormat examines the font data and returns the detected format.
func DetectFontFormat(data []byte) FontFormat {
	if len(data) < 4 {
		return FormatUnknown
	}

	signature := binary.BigEndian.Uint32(data[:4])

	switch signature {
	case signatureOTTO:
		return FormatOpenTypeCFF
	case signatureTTF, signatureTrue:
		// Both 0x00010000 and 'true' indicate TrueType outlines
		// OpenType with TTF outlines uses the same signature as TTF
		return FormatTrueType
	case signatureTTC:
		return FormatTrueTypeCollection
	default:
		return FormatUnknown
	}
}

// ParseOpenType parses OpenType font data (either TrueType or CFF outlines).
// For TrueType-based OpenType fonts, this delegates to ParseTrueType.
// For CFF-based OpenType fonts, this extracts metrics from standard tables
// while recognizing the CFF outline format.
func ParseOpenType(data []byte) (*Font, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf(
			"%w: data too short",
			ErrInvalidFont,
		)
	}

	format := DetectFontFormat(data)

	switch format {
	case FormatTrueType, FormatOpenTypeTTF:
		// TrueType-based fonts use the same parsing logic
		return ParseTrueType(data)

	case FormatOpenTypeCFF:
		// CFF-based OpenType requires special handling
		return parseOpenTypeCFF(data)

	case FormatTrueTypeCollection:
		// TTC files contain multiple fonts; parse the first one by default
		return parseTrueTypeCollection(data, 0)

	case FormatUnknown:
		return nil, fmt.Errorf(
			"%w: unknown font format",
			ErrUnsupportedFormat,
		)

	default:
		return nil, fmt.Errorf(
			"%w: unknown font format",
			ErrUnsupportedFormat,
		)
	}
}

// ParseOpenTypeFile loads and parses an OpenType font from a file.
func ParseOpenTypeFile(
	path string,
) (*Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read font file: %w",
			err,
		)
	}

	return ParseOpenType(data)
}

// ParseFont auto-detects the font format and parses it appropriately.
// This is the recommended entry point for loading fonts of any type.
func ParseFont(data []byte) (*Font, error) {
	format := DetectFontFormat(data)

	switch format {
	case FormatTrueType, FormatOpenTypeTTF:
		return ParseTrueType(data)
	case FormatOpenTypeCFF:
		return parseOpenTypeCFF(data)
	case FormatTrueTypeCollection:
		return parseTrueTypeCollection(data, 0)
	case FormatUnknown:
		return nil, fmt.Errorf(
			"%w: unable to detect font format",
			ErrUnsupportedFormat,
		)
	default:
		return nil, fmt.Errorf(
			"%w: unable to detect font format",
			ErrUnsupportedFormat,
		)
	}
}

// ParseFontFile loads and parses a font file with auto-detection.
// This is the recommended entry point for loading font files.
func ParseFontFile(path string) (*Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read font file: %w",
			err,
		)
	}

	return ParseFont(data)
}

// parseOpenTypeCFF parses an OpenType font with CFF outlines.
// CFF fonts still use standard OpenType tables for metrics (head, hhea, hmtx, etc.)
// The CFF table contains the actual glyph outlines, but we don't need to fully parse it
// for metrics extraction.
func parseOpenTypeCFF(
	data []byte,
) (*Font, error) {
	r := bytes.NewReader(data)

	// Read the font header
	var header ttfHeader
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return nil, fmt.Errorf(
			"%w: failed to read header: %v",
			ErrInvalidFont,
			err,
		)
	}

	// Verify OTTO signature
	if header.SfntVersion != signatureOTTO {
		return nil, fmt.Errorf(
			"%w: expected OTTO signature",
			ErrNotOpenType,
		)
	}

	// Read table directory
	tables := make(map[string]tableRecord)
	for range header.NumTables {
		var record tableRecord
		if err := binary.Read(r, binary.BigEndian, &record); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to read table record: %v",
				ErrInvalidFont,
				err,
			)
		}
		tag := string(record.Tag[:])
		tables[tag] = record
	}

	// Verify CFF table exists
	if _, hasCFF := tables[tagCFF]; !hasCFF {
		return nil, fmt.Errorf(
			"%w: CFF table not found in OTTO font",
			ErrMissingTable,
		)
	}

	font := &Font{
		GlyphData: make(map[rune]GlyphMetrics),
		Data:      data,
	}

	// Parse standard tables for metrics (same as TrueType)
	if err := parseHeadTable(data, tables, font); err != nil {
		return nil, err
	}

	numGlyphs, err := parseMaxpTable(data, tables)
	if err != nil {
		return nil, err
	}

	numHMetrics, err := parseHheaTable(
		data,
		tables,
		font,
	)
	if err != nil {
		return nil, err
	}

	// OS/2 table is optional
	_ = parseOS2Table(data, tables, font)

	// name table is optional
	if err := parseNameTable(data, tables, font); err != nil {
		font.Family = "Unknown"
	}

	// Parse hmtx for glyph widths
	glyphWidths, err := parseHmtxTable(
		data,
		tables,
		numHMetrics,
		numGlyphs,
	)
	if err != nil {
		return nil, err
	}

	// Parse cmap for character-to-glyph mapping
	if err := parseCmapTable(data, tables, glyphWidths, font); err != nil {
		return nil, err
	}

	// Optionally parse CFF table for additional font info
	if cffData, err := getTableData(data, tables, tagCFF); err == nil {
		_ = parseCFFInfo(cffData, font)
	}

	return font, nil
}

// parseCFFInfo extracts font information from the CFF table.
// This is a minimal parser that extracts font name information without
// fully parsing the CFF format (which is complex and not needed for metrics).
func parseCFFInfo(data []byte, font *Font) error {
	if len(data) < 4 {
		return fmt.Errorf(
			"%w: CFF data too short",
			ErrCFFParseFailed,
		)
	}

	// CFF Header structure:
	// major (1 byte) - Major version number
	// minor (1 byte) - Minor version number
	// hdrSize (1 byte) - Header size
	// offSize (1 byte) - Absolute offset size

	major := data[0]
	minor := data[1]
	hdrSize := data[2]

	// Validate version (should be 1.0)
	if major != 1 || minor != 0 {
		// CFF2 or unknown version - we can still use metrics from other tables
		return nil
	}

	if int(hdrSize) > len(data) {
		return fmt.Errorf(
			"%w: invalid header size",
			ErrCFFParseFailed,
		)
	}

	// After the header is the Name INDEX
	// INDEX structure:
	// count (2 bytes) - Number of entries
	// offSize (1 byte) - Offset size
	// offset array - (count+1) offsets
	// data - The actual data

	offset := int(hdrSize)
	if offset+3 > len(data) {
		return nil // No Name INDEX, use existing font name
	}

	count := binary.BigEndian.Uint16(
		data[offset : offset+2],
	)
	if count == 0 {
		return nil // Empty Name INDEX
	}

	offSize := data[offset+2]
	if offSize == 0 || offSize > 4 {
		return nil // Invalid offset size
	}

	// Calculate offset array position and read first name
	offsetArrayStart := offset + 3
	offsetArrayEnd := offsetArrayStart + int(
		count+1,
	)*int(
		offSize,
	)

	if offsetArrayEnd > len(data) {
		return nil
	}

	// Read first two offsets to get name boundaries
	offset1 := readCFFOffset(
		data[offsetArrayStart:],
		int(offSize),
	)
	offset2 := readCFFOffset(
		data[offsetArrayStart+int(offSize):],
		int(offSize),
	)

	dataStart := offsetArrayEnd
	nameStart := dataStart + int(
		offset1,
	) - 1 // CFF offsets are 1-based
	nameEnd := dataStart + int(offset2) - 1

	if nameStart >= 0 && nameEnd <= len(data) &&
		nameStart < nameEnd {
		name := string(data[nameStart:nameEnd])
		// Only use CFF name if we don't already have one from name table
		if font.Family == "" ||
			font.Family == "Unknown" {
			font.Family = name
		}
	}

	return nil
}

// readCFFOffset reads a variable-length offset from CFF data.
func readCFFOffset(
	data []byte,
	offSize int,
) uint32 {
	if len(data) < offSize {
		return 0
	}

	var result uint32
	for i := range offSize {
		result = (result << 8) | uint32(data[i])
	}

	return result
}

// parseTrueTypeCollection parses a TrueType Collection file and extracts a font.
// The index parameter specifies which font to extract (0-based).
func parseTrueTypeCollection(
	data []byte,
	index int,
) (*Font, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf(
			"%w: TTC data too short",
			ErrInvalidFont,
		)
	}

	// TTC Header:
	// ttcTag (4 bytes) - 'ttcf'
	// majorVersion (2 bytes)
	// minorVersion (2 bytes)
	// numFonts (4 bytes)
	// offsetTable[numFonts] (4 bytes each)

	signature := binary.BigEndian.Uint32(data[:4])
	if signature != signatureTTC {
		return nil, fmt.Errorf(
			"%w: not a TTC file",
			ErrInvalidFont,
		)
	}

	r := bytes.NewReader(data)
	_, _ = r.Seek(
		8,
		0,
	) // Skip signature and version

	var numFonts uint32
	if err := binary.Read(r, binary.BigEndian, &numFonts); err != nil {
		return nil, fmt.Errorf(
			"%w: failed to read numFonts: %v",
			ErrInvalidFont,
			err,
		)
	}

	if index < 0 || index >= int(numFonts) {
		return nil, fmt.Errorf(
			"%w: font index %d out of range (0-%d)",
			ErrInvalidFont,
			index,
			numFonts-1,
		)
	}

	// Read offset table
	offsets := make([]uint32, numFonts)
	for i := range numFonts {
		if err := binary.Read(r, binary.BigEndian, &offsets[i]); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to read offset: %v",
				ErrInvalidFont,
				err,
			)
		}
	}

	// Get the offset for the requested font
	fontOffset := offsets[index]
	if int(fontOffset) >= len(data) {
		return nil, fmt.Errorf(
			"%w: font offset out of bounds",
			ErrInvalidFont,
		)
	}

	// Parse the font at the given offset
	// We need to adjust the parsing to account for the offset
	return parseFontAtOffset(data, fontOffset)
}

// parseFontAtOffset parses a font embedded within a larger data block.
// This is used for TTC files where multiple fonts share table data.
func parseFontAtOffset(
	data []byte,
	offset uint32,
) (*Font, error) {
	if int(offset)+12 > len(data) {
		return nil, fmt.Errorf(
			"%w: offset too large",
			ErrInvalidFont,
		)
	}

	// Read header at offset
	r := bytes.NewReader(data[offset:])

	var header ttfHeader
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return nil, fmt.Errorf(
			"%w: failed to read header: %v",
			ErrInvalidFont,
			err,
		)
	}

	// Determine format
	switch header.SfntVersion {
	case signatureTTF, signatureTrue:
		// Parse as TrueType, but the full data includes table offsets
		// that are relative to the beginning of the file, not the font
		return parseTrueTypeWithOffset(
			data,
			offset,
			header,
		)
	case signatureOTTO:
		return parseOpenTypeCFFWithOffset(
			data,
			offset,
			header,
		)
	default:
		return nil, fmt.Errorf(
			"%w: unknown sfnt version at offset",
			ErrUnsupportedFormat,
		)
	}
}

// parseTrueTypeWithOffset parses a TrueType font that starts at a given offset.
// Table offsets in the font are absolute (relative to data start).
func parseTrueTypeWithOffset(
	data []byte,
	offset uint32,
	header ttfHeader,
) (*Font, error) {
	r := bytes.NewReader(data[offset+12:])

	// Read table directory
	tables := make(map[string]tableRecord)
	for range header.NumTables {
		var record tableRecord
		if err := binary.Read(r, binary.BigEndian, &record); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to read table record: %v",
				ErrInvalidFont,
				err,
			)
		}
		tag := string(record.Tag[:])
		tables[tag] = record
	}

	font := &Font{
		GlyphData: make(map[rune]GlyphMetrics),
		Data:      data, // Keep full TTC data for proper table access
	}

	// Parse tables using the original data (offsets are absolute)
	if err := parseHeadTable(data, tables, font); err != nil {
		return nil, err
	}

	numGlyphs, err := parseMaxpTable(data, tables)
	if err != nil {
		return nil, err
	}

	numHMetrics, err := parseHheaTable(
		data,
		tables,
		font,
	)
	if err != nil {
		return nil, err
	}

	if err := parseOS2Table(data, tables, font); err != nil {
		// Optional
	}

	if err := parseNameTable(data, tables, font); err != nil {
		font.Family = "Unknown"
	}

	glyphWidths, err := parseHmtxTable(
		data,
		tables,
		numHMetrics,
		numGlyphs,
	)
	if err != nil {
		return nil, err
	}

	if err := parseCmapTable(data, tables, glyphWidths, font); err != nil {
		return nil, err
	}

	return font, nil
}

// parseOpenTypeCFFWithOffset parses an OpenType/CFF font that starts at a given offset.
func parseOpenTypeCFFWithOffset(
	data []byte,
	offset uint32,
	header ttfHeader,
) (*Font, error) {
	r := bytes.NewReader(data[offset+12:])

	// Read table directory
	tables := make(map[string]tableRecord)
	for range header.NumTables {
		var record tableRecord
		if err := binary.Read(r, binary.BigEndian, &record); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to read table record: %v",
				ErrInvalidFont,
				err,
			)
		}
		tag := string(record.Tag[:])
		tables[tag] = record
	}

	// Verify CFF table exists
	if _, hasCFF := tables[tagCFF]; !hasCFF {
		return nil, fmt.Errorf(
			"%w: CFF table not found",
			ErrMissingTable,
		)
	}

	font := &Font{
		GlyphData: make(map[rune]GlyphMetrics),
		Data:      data,
	}

	// Parse tables using the original data (offsets are absolute)
	if err := parseHeadTable(data, tables, font); err != nil {
		return nil, err
	}

	numGlyphs, err := parseMaxpTable(data, tables)
	if err != nil {
		return nil, err
	}

	numHMetrics, err := parseHheaTable(
		data,
		tables,
		font,
	)
	if err != nil {
		return nil, err
	}

	if err := parseOS2Table(data, tables, font); err != nil {
		// Optional
	}

	if err := parseNameTable(data, tables, font); err != nil {
		font.Family = "Unknown"
	}

	glyphWidths, err := parseHmtxTable(
		data,
		tables,
		numHMetrics,
		numGlyphs,
	)
	if err != nil {
		return nil, err
	}

	if err := parseCmapTable(data, tables, glyphWidths, font); err != nil {
		return nil, err
	}

	// Parse CFF info if available
	if cffData, err := getTableData(data, tables, tagCFF); err == nil {
		if err := parseCFFInfo(cffData, font); err != nil {
			// Non-fatal
		}
	}

	return font, nil
}

// ParseFontCollectionFile loads a TrueType Collection file and returns a list of fonts.
func ParseFontCollectionFile(
	path string,
) ([]*Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read font file: %w",
			err,
		)
	}

	return ParseFontCollection(data)
}

// ParseFontCollection parses all fonts in a TrueType Collection.
func ParseFontCollection(
	data []byte,
) ([]*Font, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf(
			"%w: data too short",
			ErrInvalidFont,
		)
	}

	signature := binary.BigEndian.Uint32(data[:4])
	if signature != signatureTTC {
		// Not a TTC, try parsing as single font
		font, err := ParseFont(data)
		if err != nil {
			return nil, err
		}

		return []*Font{font}, nil
	}

	r := bytes.NewReader(data)
	_, _ = r.Seek(8, 0)

	var numFonts uint32
	if err := binary.Read(r, binary.BigEndian, &numFonts); err != nil {
		return nil, fmt.Errorf(
			"%w: failed to read numFonts: %v",
			ErrInvalidFont,
			err,
		)
	}

	fonts := make([]*Font, 0, numFonts)
	for i := range numFonts {
		font, err := parseTrueTypeCollection(
			data,
			int(i),
		)
		if err != nil {
			// Log error but continue with other fonts
			continue
		}
		fonts = append(fonts, font)
	}

	if len(fonts) == 0 {
		return nil, fmt.Errorf(
			"%w: no fonts could be parsed from collection",
			ErrInvalidFont,
		)
	}

	return fonts, nil
}
