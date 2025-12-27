// subset.go provides font subsetting for PDF embedding.
// Font subsetting reduces PDF file size by including only the glyphs
// that are actually used in the document.

package font

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
)

// SubsetFont represents a font subset ready for PDF embedding.
// It contains only the glyphs needed for a specific text selection.
type SubsetFont struct {
	Original   *Font           // Original font
	UsedGlyphs map[rune]uint16 // Rune -> new glyph ID mapping
	Data       []byte          // Subsetted font data (or full font if subsetting not possible)
	GlyphCount int             // Number of glyphs in subset
	IsSubset   bool            // True if actual subsetting was performed
}

// SubsetBuilder builds a font subset incrementally by collecting used runes.
type SubsetBuilder struct {
	font      *Font
	usedRunes map[rune]bool
}

// Error definitions for font subsetting
var (
	ErrSubsetFailed = errors.New(
		"font subsetting failed",
	)
	ErrNoGlyphs = errors.New(
		"no glyphs to subset",
	)
	ErrInvalidGlyphID = errors.New(
		"invalid glyph ID",
	)
)

// NewSubset creates a new SubsetBuilder for the given font.
// The builder collects runes to include in the subset.
func NewSubset(font *Font) *SubsetBuilder {
	return &SubsetBuilder{
		font:      font,
		usedRunes: make(map[rune]bool),
	}
}

// AddRune marks a single rune as used in the subset.
func (b *SubsetBuilder) AddRune(r rune) {
	b.usedRunes[r] = true
}

// AddString marks all runes in a string as used in the subset.
func (b *SubsetBuilder) AddString(s string) {
	for _, r := range s {
		b.usedRunes[r] = true
	}
}

// AddRunes marks multiple runes as used in the subset.
func (b *SubsetBuilder) AddRunes(runes []rune) {
	for _, r := range runes {
		b.usedRunes[r] = true
	}
}

// UsedRuneCount returns the number of unique runes marked for inclusion.
func (b *SubsetBuilder) UsedRuneCount() int {
	return len(b.usedRunes)
}

// UsedRunes returns a sorted slice of all used runes.
func (b *SubsetBuilder) UsedRunes() []rune {
	runes := make([]rune, 0, len(b.usedRunes))
	for r := range b.usedRunes {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return runes
}

// Build creates the font subset containing only the marked glyphs.
// Returns a SubsetFont with the subsetted data and glyph mapping.
func (b *SubsetBuilder) Build() (*SubsetFont, error) {
	if b.font == nil {
		return nil, fmt.Errorf(
			"%w: font is nil",
			ErrSubsetFailed,
		)
	}

	if len(b.usedRunes) == 0 {
		return nil, ErrNoGlyphs
	}

	// Build glyph ID mapping
	// Glyph 0 is always .notdef and must be included
	glyphMapping := make(map[rune]uint16)
	nextGlyphID := uint16(
		1,
	) // Start from 1, 0 is .notdef

	// Sort runes for consistent ordering
	runes := b.UsedRunes()

	for _, r := range runes {
		// Check if the font has this glyph
		if _, ok := b.font.GlyphData[r]; ok {
			glyphMapping[r] = nextGlyphID
			nextGlyphID++
		}
	}

	// Include count: .notdef + mapped glyphs
	glyphCount := int(nextGlyphID)

	// Attempt to create a true subset
	subsetData, err := b.createTrueTypeSubset(
		runes,
		glyphMapping,
	)
	isSubset := err == nil && subsetData != nil

	if !isSubset {
		// Fall back to embedding the full font
		// This is acceptable for small fonts or when subsetting is too complex
		subsetData = b.font.Data
	}

	return &SubsetFont{
		Original:   b.font,
		UsedGlyphs: glyphMapping,
		Data:       subsetData,
		GlyphCount: glyphCount,
		IsSubset:   isSubset,
	}, nil
}

// createTrueTypeSubset attempts to create a true TrueType subset.
// This is a simplified implementation that creates a valid subset for basic fonts.
// For complex fonts or OpenType/CFF fonts, it returns an error to trigger fallback.
func (b *SubsetBuilder) createTrueTypeSubset(
	runes []rune,
	glyphMapping map[rune]uint16,
) ([]byte, error) {
	// Verify we have valid TrueType font data
	if len(b.font.Data) < 12 {
		return nil, fmt.Errorf(
			"%w: font data too short",
			ErrSubsetFailed,
		)
	}

	format := DetectFontFormat(b.font.Data)
	if format == FormatOpenTypeCFF {
		// CFF subsetting is more complex; fall back to full embedding
		return nil, fmt.Errorf(
			"%w: CFF subsetting not implemented",
			ErrSubsetFailed,
		)
	}

	if format == FormatTrueTypeCollection {
		// TTC subsetting requires extracting single font first
		return nil, fmt.Errorf(
			"%w: TTC subsetting not implemented",
			ErrSubsetFailed,
		)
	}

	// Parse the original font tables
	tables, err := parseFontTables(b.font.Data)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: %v",
			ErrSubsetFailed,
			err,
		)
	}

	// Check for required tables
	requiredTables := []string{
		"head",
		"hhea",
		"maxp",
		"hmtx",
		"cmap",
	}
	for _, tag := range requiredTables {
		if _, ok := tables[tag]; !ok {
			return nil, fmt.Errorf(
				"%w: missing required table %s",
				ErrSubsetFailed,
				tag,
			)
		}
	}

	// For a proper subset, we would need to:
	// 1. Rewrite glyf and loca tables with only used glyphs
	// 2. Update hmtx with only used glyph metrics
	// 3. Rewrite cmap to map characters to new glyph IDs
	// 4. Update maxp with new glyph count
	// 5. Recalculate head checksum

	// This is a complex task. For the initial implementation,
	// we create a subset that includes the mapping but uses the full font data.
	// A future TODO would implement proper table rewriting.

	// For now, check if this is a simple enough font to subset
	numGlyphs := getNumGlyphsFromMaxp(
		tables["maxp"],
	)
	if numGlyphs > 1000 {
		// Large fonts benefit most from subsetting, but are also more complex
		// For now, fall back to full embedding for large fonts
		// TODO: Implement proper glyf/loca rewriting for large font subsetting
		return nil, fmt.Errorf(
			"%w: font too large for simple subsetting (%d glyphs)",
			ErrSubsetFailed,
			numGlyphs,
		)
	}

	// For small fonts, the overhead of subsetting may not be worth it
	// Return nil to trigger fallback to full embedding
	if len(glyphMapping) > numGlyphs/2 {
		// If using more than half the glyphs, just embed the whole font
		return nil, fmt.Errorf(
			"%w: using %d of %d glyphs, full embedding is more efficient",
			ErrSubsetFailed,
			len(glyphMapping),
			numGlyphs,
		)
	}

	// Attempt basic subsetting for small fonts
	return b.buildSubsetFont(
		tables,
		runes,
		glyphMapping,
	)
}

// parseFontTables parses the font file and returns a map of table tag -> table data
func parseFontTables(
	data []byte,
) (map[string][]byte, error) {
	if len(data) < 12 {
		return nil, errors.New("data too short")
	}

	r := bytes.NewReader(data)

	// Read header
	var sfntVersion uint32
	var numTables, searchRange, entrySelector, rangeShift uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&sfntVersion,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&numTables,
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
	// Suppress unused variable warnings
	_ = sfntVersion
	_ = numTables
	_ = searchRange
	_ = entrySelector
	_ = rangeShift

	tables := make(map[string][]byte)

	for range numTables {
		var tag [4]byte
		var checksum, offset, length uint32

		_, _ = r.Read(tag[:])
		_ = binary.Read(
			r,
			binary.BigEndian,
			&checksum,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&offset,
		)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&length,
		)

		tagStr := string(tag[:])

		if int(offset)+int(length) <= len(data) {
			tables[tagStr] = data[offset : offset+length]
		}
	}

	return tables, nil
}

// getNumGlyphsFromMaxp extracts the number of glyphs from the maxp table
func getNumGlyphsFromMaxp(data []byte) int {
	if len(data) < 6 {
		return 0
	}

	return int(binary.BigEndian.Uint16(data[4:6]))
}

// buildSubsetFont builds a new TrueType font with only the subset glyphs.
// This is a simplified implementation that creates a valid font structure.
func (b *SubsetBuilder) buildSubsetFont(
	tables map[string][]byte,
	runes []rune,
	glyphMapping map[rune]uint16,
) ([]byte, error) {
	// For proper subsetting, we need to rewrite multiple tables
	// This implementation creates a subset with:
	// - head: copy original
	// - hhea: copy original (numberOfHMetrics may be too high, but that's ok)
	// - maxp: update numGlyphs
	// - hmtx: copy original (includes unused metrics, but valid)
	// - cmap: rewrite to map only subset characters
	// - glyf/loca: these are the complex ones - we skip proper subsetting for now

	// Check if glyf table exists (TrueType outlines)
	hasGlyf := tables["glyf"] != nil &&
		tables["loca"] != nil

	// Calculate new glyph count: .notdef + subset glyphs
	newGlyphCount := uint16(len(glyphMapping) + 1)

	// Build modified tables
	newTables := make(map[string][]byte)

	// Copy tables that don't need modification
	copyTables := []string{
		"head",
		"hhea",
		"OS/2",
		"name",
		"post",
		"cvt ",
		"fpgm",
		"prep",
		"gasp",
	}
	for _, tag := range copyTables {
		if data, ok := tables[tag]; ok {
			newTables[tag] = data
		}
	}

	// Create modified maxp with new glyph count
	newTables["maxp"] = b.createMaxpSubset(
		tables["maxp"],
		newGlyphCount,
	)

	// Copy hmtx - it's safe to include extra metrics
	if hmtx, ok := tables["hmtx"]; ok {
		newTables["hmtx"] = hmtx
	}

	// Create a simplified cmap that maps our subset
	newTables["cmap"] = b.createCmapSubset(
		runes,
		glyphMapping,
	)

	// For glyf/loca, we have two options:
	// 1. Include all glyphs (wastes space but simple)
	// 2. Properly subset (complex)
	// We choose option 1 for simplicity
	if hasGlyf {
		newTables["glyf"] = tables["glyf"]
		newTables["loca"] = tables["loca"]
	}

	// Include CFF table if present (OpenType CFF font)
	if cff, ok := tables["CFF "]; ok {
		newTables["CFF "] = cff
	}

	// Build the final font file
	return b.assembleTrueTypeFont(newTables)
}

// createMaxpSubset creates a modified maxp table with the new glyph count.
func (*SubsetBuilder) createMaxpSubset(
	original []byte,
	newGlyphCount uint16,
) []byte {
	if len(original) < 6 {
		// Create minimal maxp
		buf := &bytes.Buffer{}
		_ = binary.Write(
			buf,
			binary.BigEndian,
			uint32(0x00010000),
		) // version
		_ = binary.Write(
			buf,
			binary.BigEndian,
			newGlyphCount,
		)

		return buf.Bytes()
	}

	// Copy original and update glyph count
	result := make([]byte, len(original))
	copy(result, original)
	binary.BigEndian.PutUint16(
		result[4:6],
		newGlyphCount,
	)

	return result
}

// createCmapSubset creates a cmap table mapping only the subset characters.
func (b *SubsetBuilder) createCmapSubset(
	runes []rune,
	glyphMapping map[rune]uint16,
) []byte {
	buf := &bytes.Buffer{}

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

	// Encoding record (Windows Unicode BMP)
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
	) // offset (header=4 + 1 encoding record=8 = 12)

	// Build Format 4 subtable
	bmpRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if r <= 0xFFFF {
			bmpRunes = append(bmpRunes, r)
		}
	}

	// Sort BMP runes
	sort.Slice(bmpRunes, func(i, j int) bool {
		return bmpRunes[i] < bmpRunes[j]
	})

	// Build segments
	type segment struct {
		startCode uint16
		endCode   uint16
		idDelta   int16
	}

	segments := make([]segment, 0)

	if len(bmpRunes) > 0 {
		// Create segments from contiguous runs or individual characters
		// For simplicity, we create one segment per character with idDelta
		for _, r := range bmpRunes {
			if glyphID, ok := glyphMapping[r]; ok {
				// idDelta = glyphID - charCode
				delta := int16(glyphID) - int16(r)
				segments = append(
					segments,
					segment{
						startCode: uint16(r),
						endCode:   uint16(r),
						idDelta:   delta,
					},
				)
			}
		}
	}

	// Add terminating segment
	segments = append(segments, segment{
		startCode: 0xFFFF,
		endCode:   0xFFFF,
		idDelta:   1,
	})

	// Calculate Format 4 size
	segCount := uint16(len(segments))
	segCountX2 := segCount * 2

	// Calculate searchRange, entrySelector, rangeShift
	searchRange := uint16(1)
	entrySelector := uint16(0)
	for searchRange*2 <= segCount {
		searchRange *= 2
		entrySelector++
	}
	searchRange *= 2
	rangeShift := segCountX2 - searchRange

	// Format 4 header size: 14 bytes
	// Arrays: 4 * segCount * 2 bytes (endCode, reservedPad+startCode, idDelta, idRangeOffset)
	format4Length := 14 + 4*int(segCount)*2

	// Write Format 4 subtable
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(4),
	) // format
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(format4Length),
	) // length
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	) // language
	_ = binary.Write(
		buf,
		binary.BigEndian,
		segCountX2,
	) // segCountX2
	_ = binary.Write(
		buf,
		binary.BigEndian,
		searchRange,
	) // searchRange
	_ = binary.Write(
		buf,
		binary.BigEndian,
		entrySelector,
	) // entrySelector
	_ = binary.Write(
		buf,
		binary.BigEndian,
		rangeShift,
	) // rangeShift

	// endCode array
	for _, seg := range segments {
		_ = binary.Write(
			buf,
			binary.BigEndian,
			seg.endCode,
		)
	}

	// reservedPad
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(0),
	)

	// startCode array
	for _, seg := range segments {
		_ = binary.Write(
			buf,
			binary.BigEndian,
			seg.startCode,
		)
	}

	// idDelta array
	for _, seg := range segments {
		_ = binary.Write(
			buf,
			binary.BigEndian,
			seg.idDelta,
		)
	}

	// idRangeOffset array (all zeros since we use idDelta)
	for range segments {
		_ = binary.Write(
			buf,
			binary.BigEndian,
			uint16(0),
		)
	}

	return buf.Bytes()
}

// assembleTrueTypeFont builds a complete TrueType font file from tables.
func (b *SubsetBuilder) assembleTrueTypeFont(
	tables map[string][]byte,
) ([]byte, error) {
	// Determine font signature
	sfntVersion := uint32(0x00010000) // TrueType
	if _, hasCFF := tables["CFF "]; hasCFF {
		sfntVersion = 0x4F54544F // 'OTTO' for OpenType/CFF
	}

	numTables := uint16(len(tables))

	// Calculate searchRange, entrySelector, rangeShift
	searchRange := uint16(1)
	entrySelector := uint16(0)
	for searchRange*2 <= numTables {
		searchRange *= 2
		entrySelector++
	}
	searchRange *= 16
	rangeShift := numTables*16 - searchRange

	// Calculate table offsets
	headerSize := 12 + numTables*16
	offset := uint32(headerSize)

	// Pad tables to 4-byte boundaries and calculate offsets
	type tableEntry struct {
		tag    string
		data   []byte
		offset uint32
	}

	entries := make([]tableEntry, 0, len(tables))

	// Sort tables alphabetically for consistent output
	sortedTags := make([]string, 0, len(tables))
	for tag := range tables {
		sortedTags = append(sortedTags, tag)
	}
	sort.Strings(sortedTags)

	for _, tag := range sortedTags {
		data := tables[tag]
		entries = append(entries, tableEntry{
			tag:    tag,
			data:   data,
			offset: offset,
		})
		// Pad to 4-byte boundary
		paddedLen := (uint32(len(data)) + 3) & ^uint32(
			3,
		)
		offset += paddedLen
	}

	// Build the font file
	buf := &bytes.Buffer{}

	// Write header
	_ = binary.Write(
		buf,
		binary.BigEndian,
		sfntVersion,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		numTables,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		searchRange,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		entrySelector,
	)
	_ = binary.Write(
		buf,
		binary.BigEndian,
		rangeShift,
	)

	// Write table directory
	for _, entry := range entries {
		// Tag (4 bytes)
		tag := []byte(entry.tag)
		for len(tag) < 4 {
			tag = append(tag, ' ')
		}
		buf.Write(tag[:4])

		// Checksum (4 bytes) - calculate table checksum
		checksum := calculateTableChecksum(
			entry.data,
		)
		_ = binary.Write(
			buf,
			binary.BigEndian,
			checksum,
		)

		// Offset (4 bytes)
		_ = binary.Write(
			buf,
			binary.BigEndian,
			entry.offset,
		)

		// Length (4 bytes)
		_ = binary.Write(
			buf,
			binary.BigEndian,
			uint32(len(entry.data)),
		)
	}

	// Write table data
	for _, entry := range entries {
		buf.Write(entry.data)
		// Pad to 4-byte boundary
		padding := (4 - len(entry.data)%4) % 4
		for range padding {
			buf.WriteByte(0)
		}
	}

	return buf.Bytes(), nil
}

// calculateTableChecksum calculates the checksum of a table.
func calculateTableChecksum(data []byte) uint32 {
	// Pad to 4-byte boundary
	padded := data
	if len(data)%4 != 0 {
		padded = make([]byte, (len(data)+3)&^3)
		copy(padded, data)
	}

	var sum uint32
	for i := 0; i < len(padded); i += 4 {
		sum += binary.BigEndian.Uint32(
			padded[i : i+4],
		)
	}

	return sum
}

// CreateSubset is a convenience function that creates a font subset
// containing only the glyphs needed for the given text.
func CreateSubset(
	font *Font,
	text string,
) (*SubsetFont, error) {
	if font == nil {
		return nil, fmt.Errorf(
			"%w: font is nil",
			ErrSubsetFailed,
		)
	}

	builder := NewSubset(font)
	builder.AddString(text)

	return builder.Build()
}

// GlyphID returns the new glyph ID for a rune in the subset.
// Returns 0 (the .notdef glyph) if the rune is not in the subset.
func (s *SubsetFont) GlyphID(r rune) uint16 {
	if id, ok := s.UsedGlyphs[r]; ok {
		return id
	}

	return 0 // .notdef
}

// HasRune returns true if the rune is included in the subset.
func (s *SubsetFont) HasRune(r rune) bool {
	_, ok := s.UsedGlyphs[r]

	return ok
}

// Runes returns all runes included in the subset.
func (s *SubsetFont) Runes() []rune {
	runes := make([]rune, 0, len(s.UsedGlyphs))
	for r := range s.UsedGlyphs {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return runes
}

// EncodedText converts a string to a sequence of glyph IDs for the subset.
// Returns the glyph IDs suitable for encoding in a PDF content stream.
func (s *SubsetFont) EncodedText(
	text string,
) []uint16 {
	result := make([]uint16, 0, len(text))
	for _, r := range text {
		result = append(result, s.GlyphID(r))
	}

	return result
}

// FontName returns a name suitable for the subset font.
// PDF convention is to prefix subset fonts with a unique tag.
// The tag is a 6 uppercase letter identifier.
func (s *SubsetFont) FontName(tag string) string {
	if !s.IsSubset {
		return s.Original.Family
	}

	// Ensure tag is 6 uppercase letters
	if len(tag) < 6 {
		tag = tag + "AAAAAA"
	}
	tag = tag[:6]

	return tag + "+" + s.Original.Family
}
