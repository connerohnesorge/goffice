// truetype.go provides TrueType font parsing for extracting metrics needed for PDF text layout.

package font

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

// FontStyle represents the style of a font (regular, bold, italic, etc.)
type FontStyle int

const (
	// StyleRegular is the standard font weight and upright style
	StyleRegular FontStyle = iota
	// StyleBold is the bold font weight
	StyleBold
	// StyleItalic is the italic (oblique) style
	StyleItalic
	// StyleBoldItalic is bold weight with italic style
	StyleBoldItalic
)

// String returns the human-readable name of the font style.
func (s FontStyle) String() string {
	switch s {
	case StyleRegular:
		return "Regular"
	case StyleBold:
		return "Bold"
	case StyleItalic:
		return "Italic"
	case StyleBoldItalic:
		return "Bold Italic"
	default:
		return "Unknown"
	}
}

// FontMetrics contains font-level measurements in font design units.
type FontMetrics struct {
	UnitsPerEm int // Units per em square (typically 1000 or 2048)
	Ascender   int // Height above baseline
	Descender  int // Depth below baseline (typically negative)
	LineGap    int // Extra spacing between lines
	CapHeight  int // Height of capital letters
	XHeight    int // Height of lowercase 'x'
}

// GlyphMetrics contains per-glyph measurements in font design units.
type GlyphMetrics struct {
	AdvanceWidth int // Horizontal advance width
	LeftBearing  int // Left side bearing
}

// KerningPair represents a kerning adjustment between two glyphs.
type KerningPair struct {
	LeftGlyph  uint16 // Glyph index of the left character
	RightGlyph uint16 // Glyph index of the right character
	Value      int16  // Kerning value in font design units (negative = closer)
}

// KerningTable contains all kerning pairs for a font.
type KerningTable struct {
	Pairs map[uint32]int16 // Combined key: (left << 16) | right -> kerning value
}

// GetKerning returns the kerning adjustment for a glyph pair.
// The key combines left and right glyph indices for efficient lookup.
func (kt *KerningTable) GetKerning(leftGlyph, rightGlyph uint16) int16 {
	if kt == nil || kt.Pairs == nil {
		return 0
	}
	key := (uint32(leftGlyph) << 16) | uint32(rightGlyph)
	return kt.Pairs[key]
}

// AddKerning adds a kerning pair to the table.
func (kt *KerningTable) AddKerning(leftGlyph, rightGlyph uint16, value int16) {
	if kt.Pairs == nil {
		kt.Pairs = make(map[uint32]int16)
	}
	key := (uint32(leftGlyph) << 16) | uint32(rightGlyph)
	kt.Pairs[key] = value
}

// Font represents a parsed TrueType font with all data needed for PDF embedding.
type Font struct {
	Family    string                // Font family name
	Style     FontStyle             // Font style (regular, bold, italic, bolditalic)
	Metrics   FontMetrics           // Font-level metrics
	GlyphData map[rune]GlyphMetrics // Character to glyph metrics mapping
	Data      []byte                // Raw font data for embedding
	Kerning   *KerningTable         // Kerning pairs table (nil if no kerning)
	// glyphToRune maps glyph indices back to runes for kerning lookups
	glyphToRune map[uint16]rune
}

// Error definitions for font parsing
var (
	ErrInvalidFont = errors.New(
		"invalid TrueType font",
	)
	ErrMissingTable = errors.New(
		"required table not found",
	)
	ErrInvalidTableData = errors.New(
		"invalid table data",
	)
	ErrUnsupportedFormat = errors.New(
		"unsupported font format",
	)
)

// TrueType magic numbers and constants
const (
	ttfMagicTrueType uint32 = 0x00010000
	ttfMagicOTTO     uint32 = 0x4F54544F // 'OTTO' for OpenType with CFF
	ttfMagicTrue     uint32 = 0x74727565 // 'true'

	// Table tags
	tagHead = "head"
	tagHhea = "hhea"
	tagHmtx = "hmtx"
	tagCmap = "cmap"
	tagOS2  = "OS/2"
	tagName = "name"
	tagMaxp = "maxp"
	tagKern = "kern"
)

// ttfHeader represents the font file header
type ttfHeader struct {
	SfntVersion   uint32
	NumTables     uint16
	SearchRange   uint16
	EntrySelector uint16
	RangeShift    uint16
}

// tableRecord represents a single table directory entry
type tableRecord struct {
	Tag      [4]byte
	Checksum uint32
	Offset   uint32
	Length   uint32
}

// headTable represents the 'head' table data
type headTable struct {
	Version            uint32
	FontRevision       uint32
	ChecksumAdjustment uint32
	MagicNumber        uint32
	Flags              uint16
	UnitsPerEm         uint16
	Created            int64
	Modified           int64
	XMin               int16
	YMin               int16
	XMax               int16
	YMax               int16
	MacStyle           uint16
	LowestRecPPEM      uint16
	FontDirectionHint  int16
	IndexToLocFormat   int16
	GlyphDataFormat    int16
}

// hheaTable represents the 'hhea' table data
type hheaTable struct {
	Version             uint32
	Ascender            int16
	Descender           int16
	LineGap             int16
	AdvanceWidthMax     uint16
	MinLeftSideBearing  int16
	MinRightSideBearing int16
	XMaxExtent          int16
	CaretSlopeRise      int16
	CaretSlopeRun       int16
	CaretOffset         int16
	Reserved1           int16
	Reserved2           int16
	Reserved3           int16
	Reserved4           int16
	MetricDataFormat    int16
	NumberOfHMetrics    uint16
}

// os2Table represents the 'OS/2' table data (partial, we only need some fields)
//
//nolint:unused // Reserved for future OS/2 table parsing support
type os2Table struct {
	Version             uint16
	XAvgCharWidth       int16
	UsWeightClass       uint16
	UsWidthClass        uint16
	FsType              uint16
	YSubscriptXSize     int16
	YSubscriptYSize     int16
	YSubscriptXOffset   int16
	YSubscriptYOffset   int16
	YSuperscriptXSize   int16
	YSuperscriptYSize   int16
	YSuperscriptXOffset int16
	YSuperscriptYOffset int16
	YStrikeoutSize      int16
	YStrikeoutPosition  int16
	SFamilyClass        int16
	Panose              [10]byte
	UlUnicodeRange1     uint32
	UlUnicodeRange2     uint32
	UlUnicodeRange3     uint32
	UlUnicodeRange4     uint32
	AchVendID           [4]byte
	FsSelection         uint16
	UsFirstCharIndex    uint16
	UsLastCharIndex     uint16
	STypoAscender       int16
	STypoDescender      int16
	STypoLineGap        int16
	UsWinAscent         uint16
	UsWinDescent        uint16
	// Version 1+ fields
	UlCodePageRange1 uint32
	UlCodePageRange2 uint32
	// Version 2+ fields
	SxHeight   int16
	SCapHeight int16
}

// maxpTable represents the 'maxp' table (partial)
type maxpTable struct {
	Version   uint32
	NumGlyphs uint16
}

// longHorMetric represents a single horizontal metric entry
type longHorMetric struct {
	AdvanceWidth uint16
	Lsb          int16
}

// ParseTrueType parses TrueType font data and extracts metrics.
func ParseTrueType(data []byte) (*Font, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf(
			"%w: data too short",
			ErrInvalidFont,
		)
	}

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

	// Validate the magic number
	switch header.SfntVersion {
	case ttfMagicTrueType, ttfMagicTrue:
		// Valid TrueType
	case ttfMagicOTTO:
		// OpenType with CFF - we can still extract metrics
	default:
		return nil, fmt.Errorf(
			"%w: unknown sfnt version 0x%08X",
			ErrUnsupportedFormat,
			header.SfntVersion,
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

	font := &Font{
		GlyphData: make(map[rune]GlyphMetrics),
		Data:      data,
	}

	// Parse required tables

	// Parse 'head' table for unitsPerEm
	if err := parseHeadTable(data, tables, font); err != nil {
		return nil, err
	}

	// Parse 'maxp' table for numGlyphs (needed for hmtx)
	numGlyphs, err := parseMaxpTable(data, tables)
	if err != nil {
		return nil, err
	}

	// Parse 'hhea' table for ascender, descender, lineGap, numberOfHMetrics
	numHMetrics, err := parseHheaTable(
		data,
		tables,
		font,
	)
	if err != nil {
		return nil, err
	}

	// Parse 'OS/2' table for capHeight, xHeight, and style
	_ = parseOS2Table(data, tables, font)

	// Parse 'name' table for family name
	if err := parseNameTable(data, tables, font); err != nil {
		// name table is optional for our purposes, use a default
		font.Family = "Unknown"
	}

	// Parse 'hmtx' table for glyph advance widths
	glyphWidths, err := parseHmtxTable(
		data,
		tables,
		numHMetrics,
		numGlyphs,
	)
	if err != nil {
		return nil, err
	}

	// Parse 'cmap' table for character to glyph mapping
	if err := parseCmapTable(data, tables, glyphWidths, font); err != nil {
		return nil, err
	}

	// Parse 'kern' table for kerning pairs (optional)
	font.Kerning = parseKernTable(data, tables)

	return font, nil
}

// ParseTrueTypeFile loads and parses a TrueType font from a file.
func ParseTrueTypeFile(
	path string,
) (*Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read font file: %w",
			err,
		)
	}

	return ParseTrueType(data)
}

// getTableData retrieves the raw data for a table
func getTableData(
	data []byte,
	tables map[string]tableRecord,
	tag string,
) ([]byte, error) {
	record, ok := tables[tag]
	if !ok {
		return nil, fmt.Errorf(
			"%w: %s",
			ErrMissingTable,
			tag,
		)
	}

	start := int(record.Offset)
	end := start + int(record.Length)
	if start < 0 || end > len(data) ||
		start > end {
		return nil, fmt.Errorf(
			"%w: %s table has invalid bounds",
			ErrInvalidTableData,
			tag,
		)
	}

	return data[start:end], nil
}

// parseHeadTable parses the 'head' table
func parseHeadTable(
	data []byte,
	tables map[string]tableRecord,
	font *Font,
) error {
	tableData, err := getTableData(
		data,
		tables,
		tagHead,
	)
	if err != nil {
		return err
	}

	if len(tableData) < 54 {
		return fmt.Errorf(
			"%w: head table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)
	var head headTable
	if err := binary.Read(r, binary.BigEndian, &head); err != nil {
		return fmt.Errorf(
			"%w: failed to parse head table: %v",
			ErrInvalidTableData,
			err,
		)
	}

	// Validate magic number
	if head.MagicNumber != 0x5F0F3CF5 {
		return fmt.Errorf(
			"%w: invalid head table magic number",
			ErrInvalidTableData,
		)
	}

	font.Metrics.UnitsPerEm = int(head.UnitsPerEm)

	return nil
}

// parseHheaTable parses the 'hhea' table
func parseHheaTable(
	data []byte,
	tables map[string]tableRecord,
	font *Font,
) (uint16, error) {
	tableData, err := getTableData(
		data,
		tables,
		tagHhea,
	)
	if err != nil {
		return 0, err
	}

	if len(tableData) < 36 {
		return 0, fmt.Errorf(
			"%w: hhea table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)
	var hhea hheaTable
	if err := binary.Read(r, binary.BigEndian, &hhea); err != nil {
		return 0, fmt.Errorf(
			"%w: failed to parse hhea table: %v",
			ErrInvalidTableData,
			err,
		)
	}

	font.Metrics.Ascender = int(hhea.Ascender)
	font.Metrics.Descender = int(hhea.Descender)
	font.Metrics.LineGap = int(hhea.LineGap)

	return hhea.NumberOfHMetrics, nil
}

// parseOS2Table parses the 'OS/2' table
func parseOS2Table(
	data []byte,
	tables map[string]tableRecord,
	font *Font,
) error {
	tableData, err := getTableData(
		data,
		tables,
		tagOS2,
	)
	if err != nil {
		return err
	}

	if len(tableData) < 78 {
		return fmt.Errorf(
			"%w: OS/2 table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)

	// Read OS/2 table fields manually to handle version differences
	var version uint16
	if err := binary.Read(r, binary.BigEndian, &version); err != nil {
		return fmt.Errorf(
			"%w: failed to read OS/2 version",
			ErrInvalidTableData,
		)
	}

	// Read to fsSelection (at offset 62)
	_, _ = r.Seek(62, io.SeekStart)
	var fsSelection uint16
	if err := binary.Read(r, binary.BigEndian, &fsSelection); err != nil {
		return fmt.Errorf(
			"%w: failed to read fsSelection",
			ErrInvalidTableData,
		)
	}

	// Determine font style from fsSelection
	const (
		fsItalic = 1 << 0
		fsBold   = 1 << 5
	)

	isBold := fsSelection&fsBold != 0
	isItalic := fsSelection&fsItalic != 0

	switch {
	case isBold && isItalic:
		font.Style = StyleBoldItalic
	case isBold:
		font.Style = StyleBold
	case isItalic:
		font.Style = StyleItalic
	default:
		font.Style = StyleRegular
	}

	// Read sCapHeight and sxHeight (version 2+, at offsets 88 and 86)
	if version >= 2 && len(tableData) >= 96 {
		_, _ = r.Seek(86, io.SeekStart)
		var sxHeight, sCapHeight int16
		if err := binary.Read(r, binary.BigEndian, &sxHeight); err == nil {
			font.Metrics.XHeight = int(sxHeight)
		}
		if err := binary.Read(r, binary.BigEndian, &sCapHeight); err == nil {
			font.Metrics.CapHeight = int(
				sCapHeight,
			)
		}
	}

	return nil
}

// parseMaxpTable parses the 'maxp' table to get numGlyphs
func parseMaxpTable(
	data []byte,
	tables map[string]tableRecord,
) (uint16, error) {
	tableData, err := getTableData(
		data,
		tables,
		tagMaxp,
	)
	if err != nil {
		return 0, err
	}

	if len(tableData) < 6 {
		return 0, fmt.Errorf(
			"%w: maxp table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)
	var maxp maxpTable
	if err := binary.Read(r, binary.BigEndian, &maxp); err != nil {
		return 0, fmt.Errorf(
			"%w: failed to parse maxp table: %v",
			ErrInvalidTableData,
			err,
		)
	}

	return maxp.NumGlyphs, nil
}

// parseNameTable parses the 'name' table to extract font family name
func parseNameTable(
	data []byte,
	tables map[string]tableRecord,
	font *Font,
) error {
	tableData, err := getTableData(
		data,
		tables,
		tagName,
	)
	if err != nil {
		return err
	}

	if len(tableData) < 6 {
		return fmt.Errorf(
			"%w: name table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)

	var format, count, stringOffset uint16
	_ = binary.Read(r, binary.BigEndian, &format)
	_ = binary.Read(r, binary.BigEndian, &count)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&stringOffset,
	)

	// Name IDs we're interested in
	const (
		nameIDFontFamily        = 1
		nameIDFontSubfamily     = 2
		nameIDFullFontName      = 4
		nameIDTypographicFamily = 16
	)

	// Read name records to find family name
	for range count {
		var platformID, encodingID, languageID, nameID, length, offset uint16
		_ = binary.Read(
			r,
			binary.BigEndian,
			&platformID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&encodingID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&languageID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&nameID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&length,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&offset,
		)

		// Look for font family name (prefer platform 3 = Windows, encoding 1 = Unicode BMP)
		if nameID == nameIDFontFamily ||
			nameID == nameIDTypographicFamily {
			strStart := int(
				stringOffset,
			) + int(
				offset,
			)
			strEnd := strStart + int(length)
			if strStart >= 0 &&
				strEnd <= len(tableData) &&
				strStart < strEnd {
				strBytes := tableData[strStart:strEnd]

				// Platform 3 (Windows) uses UTF-16BE
				if platformID == 3 &&
					encodingID == 1 {
					font.Family = decodeUTF16BE(
						strBytes,
					)
					if nameID == nameIDTypographicFamily {
						return nil // Prefer typographic family
					}
				} else if platformID == 1 && font.Family == "" {
					// Platform 1 (Mac) uses mostly ASCII
					font.Family = string(strBytes)
				}
			}
		}
	}

	if font.Family == "" {
		font.Family = "Unknown"
	}

	return nil
}

// decodeUTF16BE decodes a UTF-16BE encoded string
func decodeUTF16BE(data []byte) string {
	if len(data)%2 != 0 {
		return ""
	}

	runes := make([]rune, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		runes[i/2] = rune(
			data[i],
		)<<8 | rune(
			data[i+1],
		)
	}

	return string(runes)
}

// parseHmtxTable parses the 'hmtx' table for glyph advance widths
func parseHmtxTable(
	data []byte,
	tables map[string]tableRecord,
	numHMetrics, numGlyphs uint16,
) ([]GlyphMetrics, error) {
	tableData, err := getTableData(
		data,
		tables,
		tagHmtx,
	)
	if err != nil {
		return nil, err
	}

	// Calculate expected size
	// numHMetrics long entries (4 bytes each) + remaining lsb entries (2 bytes each)
	expectedSize := int(
		numHMetrics,
	)*4 + (int(numGlyphs)-int(numHMetrics))*2
	if len(tableData) < int(numHMetrics)*4 {
		return nil, fmt.Errorf(
			"%w: hmtx table too short",
			ErrInvalidTableData,
		)
	}

	// Allow some flexibility in table size (some fonts have extra padding)
	_ = expectedSize

	r := bytes.NewReader(tableData)
	glyphMetrics := make(
		[]GlyphMetrics,
		numGlyphs,
	)

	// Read the full metric entries
	var lastAdvanceWidth uint16
	for i := range numHMetrics {
		var metric longHorMetric
		if err := binary.Read(r, binary.BigEndian, &metric); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to read hmtx entry: %v",
				ErrInvalidTableData,
				err,
			)
		}
		glyphMetrics[i] = GlyphMetrics{
			AdvanceWidth: int(
				metric.AdvanceWidth,
			),
			LeftBearing: int(metric.Lsb),
		}
		lastAdvanceWidth = metric.AdvanceWidth
	}

	// Remaining glyphs have the same advance width as the last one
	for i := numHMetrics; i < numGlyphs; i++ {
		var lsb int16
		if err := binary.Read(r, binary.BigEndian, &lsb); err != nil {
			// It's okay if we can't read additional lsb values
			break
		}
		glyphMetrics[i] = GlyphMetrics{
			AdvanceWidth: int(lastAdvanceWidth),
			LeftBearing:  int(lsb),
		}
	}

	return glyphMetrics, nil
}

// parseCmapTable parses the 'cmap' table to map characters to glyphs
func parseCmapTable(
	data []byte,
	tables map[string]tableRecord,
	glyphWidths []GlyphMetrics,
	font *Font,
) error {
	tableData, err := getTableData(
		data,
		tables,
		tagCmap,
	)
	if err != nil {
		return err
	}

	if len(tableData) < 4 {
		return fmt.Errorf(
			"%w: cmap table too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(tableData)

	var version, numTables uint16
	_ = binary.Read(r, binary.BigEndian, &version)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&numTables,
	)

	// Find the best encoding table
	// Prefer: Windows Unicode BMP (3,1), Windows Unicode Full (3,10), Unicode BMP (0,3)
	var bestOffset uint32
	var bestPriority int

	for range numTables {
		var platformID, encodingID uint16
		var offset uint32
		_ = binary.Read(
			r,
			binary.BigEndian,
			&platformID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&encodingID,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&offset,
		)

		priority := 0
		if platformID == 3 && encodingID == 10 {
			priority = 4 // Windows Unicode full repertoire
		} else if platformID == 3 && encodingID == 1 {
			priority = 3 // Windows Unicode BMP
		} else if platformID == 0 && encodingID == 4 {
			priority = 2 // Unicode full
		} else if platformID == 0 && encodingID == 3 {
			priority = 1 // Unicode BMP
		}

		if priority > bestPriority {
			bestPriority = priority
			bestOffset = offset
		}
	}

	if bestOffset == 0 {
		// No suitable cmap found, but we can still use the font
		return nil
	}

	// Parse the selected subtable
	if int(bestOffset)+2 > len(tableData) {
		return fmt.Errorf(
			"%w: cmap subtable offset out of bounds",
			ErrInvalidTableData,
		)
	}

	subtableData := tableData[bestOffset:]
	subtableReader := bytes.NewReader(
		subtableData,
	)

	var format uint16
	_ = binary.Read(
		subtableReader,
		binary.BigEndian,
		&format,
	)

	switch format {
	case 4:
		return parseCmapFormat4(
			subtableData,
			glyphWidths,
			font,
		)
	case 12:
		return parseCmapFormat12(
			subtableData,
			glyphWidths,
			font,
		)
	default:
		// Unsupported format, continue without character mapping
		return nil
	}
}

// parseCmapFormat4 parses format 4 cmap subtable (segment mapping to delta values)
func parseCmapFormat4(
	data []byte,
	glyphWidths []GlyphMetrics,
	font *Font,
) error {
	if len(data) < 14 {
		return fmt.Errorf(
			"%w: cmap format 4 too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(data)

	var format, length, language, segCountX2 uint16
	var searchRange, entrySelector, rangeShift uint16

	_ = binary.Read(r, binary.BigEndian, &format)
	_ = binary.Read(r, binary.BigEndian, &length)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&language,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&segCountX2,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&searchRange,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&entrySelector,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&rangeShift,
	)

	segCount := int(segCountX2 / 2)
	if segCount == 0 {
		return nil
	}

	// Read arrays
	endCodes := make([]uint16, segCount)
	startCodes := make([]uint16, segCount)
	idDeltas := make([]int16, segCount)
	idRangeOffsets := make([]uint16, segCount)

	for i := range segCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&endCodes[i],
		)
	}

	// Skip reservedPad
	var reservedPad uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&reservedPad,
	)

	for i := range segCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&startCodes[i],
		)
	}
	for i := range segCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&idDeltas[i],
		)
	}

	// Remember position for idRangeOffset calculations
	idRangeOffsetStart, _ := r.Seek(
		0,
		io.SeekCurrent,
	)

	for i := range segCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&idRangeOffsets[i],
		)
	}

	// Initialize glyph to rune mapping
	font.glyphToRune = make(map[uint16]rune)

	// Map characters to glyphs
	for seg := range segCount {
		start := startCodes[seg]
		end := endCodes[seg]
		delta := idDeltas[seg]
		rangeOffset := idRangeOffsets[seg]

		if start == 0xFFFF {
			continue // End marker
		}

		for c := start; c <= end; c++ {
			var glyphIndex uint16

			if rangeOffset == 0 {
				glyphIndex = uint16(
					int16(c) + delta,
				)
			} else {
				// Calculate position in glyph ID array
				offset := idRangeOffsetStart + int64(seg)*2 + int64(rangeOffset) + int64(c-start)*2
				if offset >= 0 && offset+2 <= int64(len(data)) {
					glyphIndex = uint16(data[offset])<<8 | uint16(data[offset+1])
					if glyphIndex != 0 {
						glyphIndex = uint16(int16(glyphIndex) + delta)
					}
				}
			}

			if glyphIndex > 0 &&
				int(
					glyphIndex,
				) < len(
					glyphWidths,
				) {
				font.GlyphData[rune(c)] = glyphWidths[glyphIndex]
				font.glyphToRune[glyphIndex] = rune(c)
			}
		}
	}

	return nil
}

// parseCmapFormat12 parses format 12 cmap subtable (segmented coverage)
func parseCmapFormat12(
	data []byte,
	glyphWidths []GlyphMetrics,
	font *Font,
) error {
	if len(data) < 16 {
		return fmt.Errorf(
			"%w: cmap format 12 too short",
			ErrInvalidTableData,
		)
	}

	r := bytes.NewReader(data)

	var format uint16
	var reserved uint16
	var length, language, numGroups uint32

	_ = binary.Read(r, binary.BigEndian, &format)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&reserved,
	)
	_ = binary.Read(r, binary.BigEndian, &length)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&language,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&numGroups,
	)

	// Limit to prevent memory exhaustion
	if numGroups > 100000 {
		numGroups = 100000
	}

	// Build glyph to rune mapping for kerning lookups
	font.glyphToRune = make(map[uint16]rune)

	for range numGroups {
		var startCharCode, endCharCode, startGlyphID uint32
		_ = binary.Read(
			r,
			binary.BigEndian,
			&startCharCode,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&endCharCode,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&startGlyphID,
		)

		// Limit range size to prevent memory exhaustion
		if endCharCode-startCharCode > 100000 {
			endCharCode = startCharCode + 100000
		}

		for c := startCharCode; c <= endCharCode; c++ {
			glyphIndex := startGlyphID + (c - startCharCode)
			if int(
				glyphIndex,
			) < len(
				glyphWidths,
			) {
				font.GlyphData[rune(c)] = glyphWidths[glyphIndex]
				font.glyphToRune[uint16(glyphIndex)] = rune(c)
			}
		}
	}

	return nil
}

// parseKernTable parses the 'kern' table for kerning pairs.
// Returns nil if the table is not present or empty.
func parseKernTable(
	data []byte,
	tables map[string]tableRecord,
) *KerningTable {
	tableData, err := getTableData(data, tables, tagKern)
	if err != nil {
		return nil // kern table is optional
	}

	if len(tableData) < 4 {
		return nil
	}

	r := bytes.NewReader(tableData)

	// Read version
	var version uint16
	if err := binary.Read(r, binary.BigEndian, &version); err != nil {
		return nil
	}

	kerningTable := &KerningTable{
		Pairs: make(map[uint32]int16),
	}

	if version == 0 {
		// Version 0 format (Apple-style)
		parseKernVersion0(tableData, kerningTable)
	} else if version == 1 {
		// Version 1 format (Microsoft-style)
		parseKernVersion1(tableData, kerningTable)
	}

	if len(kerningTable.Pairs) == 0 {
		return nil
	}

	return kerningTable
}

// parseKernVersion0 parses Apple-style kern table (version 0).
func parseKernVersion0(data []byte, kerningTable *KerningTable) {
	if len(data) < 4 {
		return
	}

	r := bytes.NewReader(data)

	var version, nTables uint16
	_ = binary.Read(r, binary.BigEndian, &version)
	_ = binary.Read(r, binary.BigEndian, &nTables)

	// Limit number of subtables
	if nTables > 100 {
		nTables = 100
	}

	for i := uint16(0); i < nTables; i++ {
		if r.Len() < 8 {
			break
		}

		var length, coverage uint16
		var tupleIndex uint16

		_ = binary.Read(r, binary.BigEndian, &length)
		_ = binary.Read(r, binary.BigEndian, &coverage)
		_ = binary.Read(r, binary.BigEndian, &tupleIndex)

		// Check if this is a horizontal kerning subtable
		// coverage bits: 0=horizontal, 1=minimum, 2=cross-stream, 3=override
		isHorizontal := coverage&0x0001 != 0
		isMinimum := coverage&0x0002 != 0
		isCrossStream := coverage&0x0004 != 0

		if !isHorizontal || isCrossStream || isMinimum {
			// Skip this subtable
			_, _ = r.Seek(int64(length)-8, io.SeekCurrent)
			continue
		}

		format := (coverage >> 8) & 0xFF

		switch format {
		case 0:
			parseKernFormat0(r, kerningTable)
		case 2:
			parseKernFormat2(r, kerningTable)
		default:
			// Skip unknown format
			_, _ = r.Seek(int64(length)-8, io.SeekCurrent)
		}
	}
}

// parseKernVersion1 parses Microsoft-style kern table (version 1).
func parseKernVersion1(data []byte, kerningTable *KerningTable) {
	if len(data) < 8 {
		return
	}

	r := bytes.NewReader(data)

	var version uint16
	var nTables uint32

	_ = binary.Read(r, binary.BigEndian, &version)
	_ = binary.Read(r, binary.BigEndian, &nTables)

	// Limit number of subtables
	if nTables > 100 {
		nTables = 100
	}

	for i := uint32(0); i < nTables; i++ {
		if r.Len() < 8 {
			break
		}

		var length uint32
		var coverage uint16
		var tupleIndex uint16

		_ = binary.Read(r, binary.BigEndian, &length)
		_ = binary.Read(r, binary.BigEndian, &coverage)
		_ = binary.Read(r, binary.BigEndian, &tupleIndex)

		// Check if this is a horizontal kerning subtable
		isHorizontal := coverage&0x0001 != 0
		isMinimum := coverage&0x0002 != 0
		isCrossStream := coverage&0x0004 != 0

		if !isHorizontal || isCrossStream || isMinimum {
			// Skip this subtable
			_, _ = r.Seek(int64(length)-8, io.SeekCurrent)
			continue
		}

		format := coverage & 0x00FF

		switch format {
		case 0:
			parseKernFormat0MS(r, kerningTable)
		default:
			// Skip unknown format
			_, _ = r.Seek(int64(length)-8, io.SeekCurrent)
		}
	}
}

// parseKernFormat0 parses format 0 kerning subtable (Apple-style).
func parseKernFormat0(r *bytes.Reader, kerningTable *KerningTable) {
	var nPairs, searchRange, entrySelector, rangeShift uint16

	_ = binary.Read(r, binary.BigEndian, &nPairs)
	_ = binary.Read(r, binary.BigEndian, &searchRange)
	_ = binary.Read(r, binary.BigEndian, &entrySelector)
	_ = binary.Read(r, binary.BigEndian, &rangeShift)

	// Limit number of pairs
	if nPairs > 10000 {
		nPairs = 10000
	}

	for i := uint16(0); i < nPairs; i++ {
		if r.Len() < 6 {
			break
		}

		var left, right uint16
		var value int16

		_ = binary.Read(r, binary.BigEndian, &left)
		_ = binary.Read(r, binary.BigEndian, &right)
		_ = binary.Read(r, binary.BigEndian, &value)

		if value != 0 {
			key := (uint32(left) << 16) | uint32(right)
			kerningTable.Pairs[key] = value
		}
	}
}

// parseKernFormat0MS parses format 0 kerning subtable (Microsoft-style).
func parseKernFormat0MS(r *bytes.Reader, kerningTable *KerningTable) {
	// Microsoft format 0 is similar but uses 32-bit length
	var nPairs, searchRange, entrySelector, rangeShift uint16

	_ = binary.Read(r, binary.BigEndian, &nPairs)
	_ = binary.Read(r, binary.BigEndian, &searchRange)
	_ = binary.Read(r, binary.BigEndian, &entrySelector)
	_ = binary.Read(r, binary.BigEndian, &rangeShift)

	// Limit number of pairs
	if nPairs > 10000 {
		nPairs = 10000
	}

	for i := uint16(0); i < nPairs; i++ {
		if r.Len() < 6 {
			break
		}

		var left, right uint16
		var value int16

		_ = binary.Read(r, binary.BigEndian, &left)
		_ = binary.Read(r, binary.BigEndian, &right)
		_ = binary.Read(r, binary.BigEndian, &value)

		if value != 0 {
			key := (uint32(left) << 16) | uint32(right)
			kerningTable.Pairs[key] = value
		}
	}
}

// parseKernFormat2 parses format 2 kerning subtable (class-based).
func parseKernFormat2(r *bytes.Reader, kerningTable *KerningTable) {
	// Format 2 uses class-based kerning
	var rowWidth uint16
	var leftOffsetTable, rightOffsetTable, arrayOffset uint16

	_ = binary.Read(r, binary.BigEndian, &rowWidth)
	_ = binary.Read(r, binary.BigEndian, &leftOffsetTable)
	_ = binary.Read(r, binary.BigEndian, &rightOffsetTable)
	_ = binary.Read(r, binary.BigEndian, &arrayOffset)

	// Format 2 is more complex and less common
	// For now, we skip detailed parsing
	// TODO: Implement full class-based kerning support
}

// GetKerningForRunes returns the kerning adjustment for a pair of characters.
// This looks up the kerning value based on glyph indices.
func (f *Font) GetKerningForRunes(left, right rune) int16 {
	if f.Kerning == nil {
		return 0
	}

	// Find glyph indices for the runes
	// We need to reverse-lookup from the cmap data
	leftGlyph := f.findGlyphIndex(left)
	rightGlyph := f.findGlyphIndex(right)

	if leftGlyph == 0 || rightGlyph == 0 {
		return 0
	}

	return f.Kerning.GetKerning(leftGlyph, rightGlyph)
}

// findGlyphIndex finds the glyph index for a rune.
// Returns 0 if not found (0 is typically the .notdef glyph).
func (f *Font) findGlyphIndex(r rune) uint16 {
	// This is a simplified lookup - in practice, we'd need to store
	// the reverse mapping from cmap parsing
	// For now, use a simple heuristic based on glyphData
	if f.glyphToRune != nil {
		for glyphIdx, char := range f.glyphToRune {
			if char == r {
				return glyphIdx
			}
		}
	}

	// Fallback: for ASCII characters, glyph index often equals char code
	if r < 256 {
		return uint16(r)
	}

	return 0
}

// TextWidthWithKerning calculates the total width of a string including kerning adjustments.
func (f *Font) TextWidthWithKerning(s string) int {
	if f.Kerning == nil || len(s) < 2 {
		return f.TextWidth(s)
	}

	var width int
	var prevRune rune
	var hasPrev bool

	for _, r := range s {
		width += f.GlyphWidth(r)

		if hasPrev {
			kern := f.GetKerningForRunes(prevRune, r)
			width += int(kern)
		}

		prevRune = r
		hasPrev = true
	}

	return width
}

// GlyphWidth returns the advance width for a character in font design units.
// Returns 0 if the glyph is not found.
func (f *Font) GlyphWidth(r rune) int {
	if m, ok := f.GlyphData[r]; ok {
		return m.AdvanceWidth
	}

	return 0
}

// TextWidth calculates the total width of a string in font design units.
func (f *Font) TextWidth(s string) int {
	var width int
	for _, r := range s {
		width += f.GlyphWidth(r)
	}

	return width
}

// ScaledWidth converts font units to points at a given font size.
func (f *Font) ScaledWidth(
	fontUnits int,
	fontSize float64,
) float64 {
	if f.Metrics.UnitsPerEm == 0 {
		return 0
	}

	return float64(
		fontUnits,
	) * fontSize / float64(
		f.Metrics.UnitsPerEm,
	)
}
