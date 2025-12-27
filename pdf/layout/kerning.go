// kerning.go provides kerning pair lookup for improved text spacing.
// Kerning adjusts the spacing between specific character pairs for better
// visual appearance (e.g., "AV", "To", "WA" typically need negative kerning).

package layout

import (
	"bytes"
	"encoding/binary"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// KerningTable provides kerning adjustments for character pairs.
// It extracts kerning data from TrueType/OpenType fonts and provides
// efficient lookup for text layout.
type KerningTable struct {
	// font is the underlying parsed font containing kerning data.
	font *font.Font

	// pairs stores kerning values indexed by (left<<16 | right) glyph IDs.
	pairs map[uint32]int

	// runePairs stores kerning values indexed by (left<<16 | right) runes.
	// This allows direct rune lookup without glyph ID conversion.
	runePairs map[uint64]int

	// hasKerning indicates whether any kerning data was found.
	hasKerning bool

	// glyphMetrics provides glyph ID lookup for rune-based kerning.
	glyphMetrics *GlyphMetrics
}

// KernTableHeader represents the header of the kern table.
type kernTableHeader struct {
	Version   uint16
	NumTables uint16
}

// KernSubtableHeader represents a kern subtable header.
type kernSubtableHeader struct {
	Version  uint16
	Length   uint16
	Coverage uint16
}

// KernFormat0Header represents the format 0 subtable header.
type kernFormat0Header struct {
	NPairs        uint16
	SearchRange   uint16
	EntrySelector uint16
	RangeShift    uint16
}

// KernFormat0Pair represents a single kerning pair in format 0.
type kernFormat0Pair struct {
	Left  uint16
	Right uint16
	Value int16
}

// Kern table tags and constants.
const (
	tagKern = "kern"
	tagGPOS = "GPOS"

	// Coverage field bit flags
	kernHorizontal  = 0x0001 // Horizontal kerning
	kernMinimum     = 0x0002 // Minimum values (not used in most fonts)
	kernCrossStream = 0x0004 // Cross-stream (vertical) adjustment
	kernOverride    = 0x0008 // Override previous kern values
)

// NewKerningTable creates a KerningTable from a parsed font.
// Returns nil if the font is nil.
func NewKerningTable(f *font.Font) *KerningTable {
	if f == nil {
		return nil
	}

	kt := &KerningTable{
		font:      f,
		pairs:     make(map[uint32]int),
		runePairs: make(map[uint64]int),
	}

	// Create glyph metrics for rune to glyph ID mapping
	kt.glyphMetrics = NewGlyphMetrics(f)

	// Try to parse kern table from font data
	kt.parseKernTable()

	// If no kern table, try GPOS table (basic support)
	if !kt.hasKerning {
		kt.parseGPOSKerning()
	}

	return kt
}

// parseKernTable parses the 'kern' table from font data.
func (kt *KerningTable) parseKernTable() {
	if kt.font == nil || len(kt.font.Data) < 12 {
		return
	}

	data := kt.font.Data

	// Read font header to find table directory
	r := bytes.NewReader(data)

	var sfntVersion uint32
	var numTables uint16
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
	_, _ = r.Seek(
		6,
		1,
	) // Skip searchRange, entrySelector, rangeShift

	// Find kern table
	var kernOffset, kernLength uint32
	for range numTables {
		var tag [4]byte
		var checksum, offset, length uint32
		_ = binary.Read(r, binary.BigEndian, &tag)
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

		if string(tag[:]) == tagKern {
			kernOffset = offset
			kernLength = length

			break
		}
	}

	if kernOffset == 0 || kernLength == 0 {
		return
	}

	// Validate bounds
	if int(kernOffset+kernLength) > len(data) {
		return
	}

	kernData := data[kernOffset : kernOffset+kernLength]
	kt.parseKernData(kernData)
}

// parseKernData parses the kern table data.
func (kt *KerningTable) parseKernData(
	data []byte,
) {
	if len(data) < 4 {
		return
	}

	r := bytes.NewReader(data)

	// Read header
	var header kernTableHeader
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return
	}

	// Process each subtable
	for range header.NumTables {
		if !kt.parseKernSubtable(r, data) {
			break
		}
	}
}

// parseKernSubtable parses a single kern subtable.
func (kt *KerningTable) parseKernSubtable(
	r *bytes.Reader,
	data []byte,
) bool {
	startPos, _ := r.Seek(0, 1)

	// Read subtable header
	var subHeader kernSubtableHeader
	if err := binary.Read(r, binary.BigEndian, &subHeader); err != nil {
		return false
	}

	// Check coverage flags
	// We only support horizontal kerning (not cross-stream)
	isHorizontal := (subHeader.Coverage & kernHorizontal) != 0
	isCrossStream := (subHeader.Coverage & kernCrossStream) != 0

	// Skip non-horizontal or cross-stream kerning
	if !isHorizontal || isCrossStream {
		// Seek to next subtable
		_, _ = r.Seek(
			startPos+int64(subHeader.Length),
			0,
		)

		return true
	}

	// Get format from high byte of coverage
	format := (subHeader.Coverage >> 8) & 0xFF

	switch format {
	case 0:
		kt.parseKernFormat0(r)
	default:
		// Unsupported format, skip to next subtable
	}

	// Seek to next subtable (in case we didn't read all data)
	_, _ = r.Seek(
		startPos+int64(subHeader.Length),
		0,
	)

	return true
}

// parseKernFormat0 parses format 0 kerning subtable.
// Format 0 is the most common kern table format with simple pair lists.
func (kt *KerningTable) parseKernFormat0(
	r *bytes.Reader,
) {
	var f0Header kernFormat0Header
	if err := binary.Read(r, binary.BigEndian, &f0Header); err != nil {
		return
	}

	// Read kerning pairs
	for range f0Header.NPairs {
		var pair kernFormat0Pair
		if err := binary.Read(r, binary.BigEndian, &pair); err != nil {
			break
		}

		// Store the kerning value indexed by glyph pair
		key := uint32(
			pair.Left,
		)<<16 | uint32(
			pair.Right,
		)
		kt.pairs[key] = int(pair.Value)
		kt.hasKerning = true
	}
}

// parseGPOSKerning attempts to extract basic kerning from GPOS table.
// GPOS is more complex than kern, so we only support simple pair positioning.
func (kt *KerningTable) parseGPOSKerning() {
	if kt.font == nil || len(kt.font.Data) < 12 {
		return
	}

	data := kt.font.Data

	// Read font header to find table directory
	r := bytes.NewReader(data)

	var sfntVersion uint32
	var numTables uint16
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
	_, _ = r.Seek(
		6,
		1,
	) // Skip searchRange, entrySelector, rangeShift

	// Find GPOS table
	var gposOffset, gposLength uint32
	for range numTables {
		var tag [4]byte
		var checksum, offset, length uint32
		_ = binary.Read(r, binary.BigEndian, &tag)
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

		if string(tag[:]) == tagGPOS {
			gposOffset = offset
			gposLength = length

			break
		}
	}

	if gposOffset == 0 || gposLength == 0 {
		return
	}

	// Validate bounds
	if int(gposOffset+gposLength) > len(data) {
		return
	}

	gposData := data[gposOffset : gposOffset+gposLength]
	kt.parseGPOSData(gposData)
}

// parseGPOSData parses GPOS table for kerning data.
// This is a simplified parser that looks for kern feature pair positioning.
func (kt *KerningTable) parseGPOSData(
	data []byte,
) {
	if len(data) < 10 {
		return
	}

	r := bytes.NewReader(data)

	// GPOS Header
	var majorVersion, minorVersion uint16
	var scriptListOffset, featureListOffset, lookupListOffset uint16

	_ = binary.Read(
		r,
		binary.BigEndian,
		&majorVersion,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&minorVersion,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&scriptListOffset,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&featureListOffset,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&lookupListOffset,
	)

	// We need at least version 1.0
	if majorVersion < 1 {
		return
	}

	// Validate offsets
	if int(featureListOffset) >= len(data) ||
		int(lookupListOffset) >= len(data) {
		return
	}

	// Find 'kern' feature in feature list
	kernLookupIndices := kt.findKernFeature(
		data,
		featureListOffset,
	)
	if len(kernLookupIndices) == 0 {
		return
	}

	// Parse lookup tables for kerning
	kt.parseGPOSLookups(
		data,
		lookupListOffset,
		kernLookupIndices,
	)
}

// findKernFeature finds the 'kern' feature in the feature list and returns lookup indices.
func (kt *KerningTable) findKernFeature(
	data []byte,
	featureListOffset uint16,
) []uint16 {
	if int(featureListOffset)+2 > len(data) {
		return nil
	}

	r := bytes.NewReader(data[featureListOffset:])

	var featureCount uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&featureCount,
	)

	for range featureCount {
		var tag [4]byte
		var offset uint16
		_ = binary.Read(r, binary.BigEndian, &tag)
		_ = binary.Read(
			r,
			binary.BigEndian,
			&offset,
		)

		if string(tag[:]) == "kern" {
			// Found kern feature, read its lookup indices
			featureOffset := int(
				featureListOffset,
			) + int(
				offset,
			)
			if featureOffset+4 > len(data) {
				return nil
			}

			fr := bytes.NewReader(
				data[featureOffset:],
			)
			var featureParams uint16
			var lookupIndexCount uint16
			_ = binary.Read(
				fr,
				binary.BigEndian,
				&featureParams,
			)
			_ = binary.Read(
				fr,
				binary.BigEndian,
				&lookupIndexCount,
			)

			indices := make(
				[]uint16,
				lookupIndexCount,
			)
			for j := range lookupIndexCount {
				_ = binary.Read(
					fr,
					binary.BigEndian,
					&indices[j],
				)
			}

			return indices
		}
	}

	return nil
}

// parseGPOSLookups parses the lookup tables for kerning data.
func (kt *KerningTable) parseGPOSLookups(
	data []byte,
	lookupListOffset uint16,
	indices []uint16,
) {
	if int(lookupListOffset)+2 > len(data) {
		return
	}

	r := bytes.NewReader(data[lookupListOffset:])

	var lookupCount uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&lookupCount,
	)

	// Read lookup offsets
	lookupOffsets := make([]uint16, lookupCount)
	for i := range lookupCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&lookupOffsets[i],
		)
	}

	// Process each kern lookup
	for _, idx := range indices {
		if idx >= lookupCount {
			continue
		}

		lookupOffset := int(
			lookupListOffset,
		) + int(
			lookupOffsets[idx],
		)
		if lookupOffset+6 > len(data) {
			continue
		}

		kt.parseGPOSLookup(data, lookupOffset)
	}
}

// parseGPOSLookup parses a single GPOS lookup table.
func (kt *KerningTable) parseGPOSLookup(
	data []byte,
	offset int,
) {
	if offset+6 > len(data) {
		return
	}

	r := bytes.NewReader(data[offset:])

	var lookupType, lookupFlag, subtableCount uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&lookupType,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&lookupFlag,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&subtableCount,
	)

	// Lookup type 2 is Pair Positioning (kerning)
	if lookupType != 2 {
		return
	}

	// Read subtable offsets
	subtableOffsets := make(
		[]uint16,
		subtableCount,
	)
	for i := range subtableCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&subtableOffsets[i],
		)
	}

	// Process each subtable
	for _, subOffset := range subtableOffsets {
		subtableOffset := offset + int(subOffset)
		if subtableOffset+2 > len(data) {
			continue
		}

		kt.parseGPOSPairPosSubtable(
			data,
			subtableOffset,
		)
	}
}

// parseGPOSPairPosSubtable parses a GPOS Pair Positioning subtable.
func (kt *KerningTable) parseGPOSPairPosSubtable(
	data []byte,
	offset int,
) {
	if offset+10 > len(data) {
		return
	}

	r := bytes.NewReader(data[offset:])

	var posFormat uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&posFormat,
	)

	// We support format 1 (individual pairs) for simplicity
	if posFormat != 1 {
		return
	}

	var coverageOffset, valueFormat1, valueFormat2, pairSetCount uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&coverageOffset,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&valueFormat1,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&valueFormat2,
	)
	_ = binary.Read(
		r,
		binary.BigEndian,
		&pairSetCount,
	)

	// Read coverage table to get first glyph IDs
	coverage := kt.parseCoverage(
		data,
		offset+int(coverageOffset),
	)
	if len(coverage) == 0 {
		return
	}

	// Read pair set offsets
	pairSetOffsets := make([]uint16, pairSetCount)
	for i := range pairSetCount {
		_ = binary.Read(
			r,
			binary.BigEndian,
			&pairSetOffsets[i],
		)
	}

	// Calculate value record sizes
	valueSize1 := kt.valueRecordSize(valueFormat1)
	valueSize2 := kt.valueRecordSize(valueFormat2)

	// Process each pair set
	for i, pairSetOffset := range pairSetOffsets {
		if i >= len(coverage) {
			break
		}

		leftGlyph := coverage[i]
		pairSetPos := offset + int(pairSetOffset)

		if pairSetPos+2 > len(data) {
			continue
		}

		kt.parsePairSet(
			data,
			pairSetPos,
			leftGlyph,
			valueFormat1,
			valueFormat2,
			valueSize1,
			valueSize2,
		)
	}
}

// parsePairSet parses a GPOS pair set.
func (kt *KerningTable) parsePairSet(
	data []byte,
	offset int,
	leftGlyph uint16,
	valueFormat1, valueFormat2 uint16,
	valueSize1, valueSize2 int,
) {
	if offset+2 > len(data) {
		return
	}

	r := bytes.NewReader(data[offset:])

	var pairCount uint16
	_ = binary.Read(
		r,
		binary.BigEndian,
		&pairCount,
	)

	// Each pair record: secondGlyph (2) + value1 (valueSize1) + value2 (valueSize2)
	pairRecordSize := 2 + valueSize1 + valueSize2

	for range pairCount {
		pos, _ := r.Seek(0, 1)
		if int(
			pos,
		)+pairRecordSize > len(
			data,
		)-offset {
			break
		}

		var secondGlyph uint16
		_ = binary.Read(
			r,
			binary.BigEndian,
			&secondGlyph,
		)

		// Read XAdvance from value1 if present (bit 0x0004)
		xAdvance := int16(0)
		if valueFormat1&0x0004 != 0 {
			// Skip XPlacement and YPlacement if present
			if valueFormat1&0x0001 != 0 {
				_, _ = r.Seek(2, 1) // XPlacement
			}
			if valueFormat1&0x0002 != 0 {
				_, _ = r.Seek(2, 1) // YPlacement
			}
			_ = binary.Read(
				r,
				binary.BigEndian,
				&xAdvance,
			)
		}

		// Skip remaining value1 fields and all of value2
		remaining := pairRecordSize - 2
		if valueFormat1&0x0004 != 0 {
			if valueFormat1&0x0001 != 0 {
				remaining -= 2
			}
			if valueFormat1&0x0002 != 0 {
				remaining -= 2
			}
			remaining -= 2 // XAdvance
		}
		_, _ = r.Seek(int64(remaining), 1)

		// Store kerning value
		if xAdvance != 0 {
			key := uint32(
				leftGlyph,
			)<<16 | uint32(
				secondGlyph,
			)
			kt.pairs[key] = int(xAdvance)
			kt.hasKerning = true
		}
	}
}

// parseCoverage parses a GPOS coverage table.
func (kt *KerningTable) parseCoverage(
	data []byte,
	offset int,
) []uint16 {
	if offset+4 > len(data) {
		return nil
	}

	r := bytes.NewReader(data[offset:])

	var format, count uint16
	_ = binary.Read(r, binary.BigEndian, &format)
	_ = binary.Read(r, binary.BigEndian, &count)

	if format == 1 {
		// Format 1: Array of glyph IDs
		glyphs := make([]uint16, count)
		for i := range count {
			_ = binary.Read(
				r,
				binary.BigEndian,
				&glyphs[i],
			)
		}

		return glyphs
	}

	// Format 2: Range records - not fully implemented for simplicity
	return nil
}

// valueRecordSize calculates the size of a GPOS value record.
func (kt *KerningTable) valueRecordSize(
	format uint16,
) int {
	size := 0
	if format&0x0001 != 0 {
		size += 2 // XPlacement
	}
	if format&0x0002 != 0 {
		size += 2 // YPlacement
	}
	if format&0x0004 != 0 {
		size += 2 // XAdvance
	}
	if format&0x0008 != 0 {
		size += 2 // YAdvance
	}
	if format&0x0010 != 0 {
		size += 2 // XPlaDevice
	}
	if format&0x0020 != 0 {
		size += 2 // YPlaDevice
	}
	if format&0x0040 != 0 {
		size += 2 // XAdvDevice
	}
	if format&0x0080 != 0 {
		size += 2 // YAdvDevice
	}

	return size
}

// GetKerning returns the kerning adjustment for a pair of runes in font units.
// Returns 0 if no kerning data exists for the pair or if either rune is not in the font.
func (kt *KerningTable) GetKerning(
	left, right rune,
) int {
	if kt == nil || !kt.hasKerning {
		return 0
	}

	// Check rune pair cache first
	runeKey := uint64(left)<<32 | uint64(right)
	if kern, ok := kt.runePairs[runeKey]; ok {
		return kern
	}

	// Look up glyph IDs and check glyph pair kerning
	if kt.glyphMetrics == nil {
		return 0
	}

	leftID := kt.glyphMetrics.GetGlyphID(left)
	rightID := kt.glyphMetrics.GetGlyphID(right)

	if leftID == 0 || rightID == 0 {
		return 0
	}

	kern := kt.GetKerningByGlyphID(
		leftID,
		rightID,
	)

	// Cache the result for future lookups
	kt.runePairs[runeKey] = kern

	return kern
}

// GetKerningByGlyphID returns the kerning adjustment for a pair of glyph IDs.
// Returns 0 if no kerning data exists for the pair.
func (kt *KerningTable) GetKerningByGlyphID(
	leftID, rightID uint16,
) int {
	if kt == nil || !kt.hasKerning {
		return 0
	}

	key := uint32(leftID)<<16 | uint32(rightID)
	if kern, ok := kt.pairs[key]; ok {
		return kern
	}

	return 0
}

// HasKerning returns true if the font contains any kerning data.
func (kt *KerningTable) HasKerning() bool {
	if kt == nil {
		return false
	}

	return kt.hasKerning
}

// MeasureStringWithKerning returns the total advance width for a string in font units,
// including kerning adjustments between adjacent characters.
func (kt *KerningTable) MeasureStringWithKerning(
	text string,
) int {
	if kt == nil || kt.glyphMetrics == nil ||
		len(text) == 0 {
		return 0
	}

	runes := []rune(text)
	if len(runes) == 0 {
		return 0
	}

	// Start with width of first character
	totalWidth := kt.glyphMetrics.GetAdvanceWidthForRune(
		runes[0],
	)

	// Add widths of remaining characters plus kerning
	for i := 1; i < len(runes); i++ {
		// Add kerning adjustment for the pair
		totalWidth += kt.GetKerning(
			runes[i-1],
			runes[i],
		)
		// Add advance width of current character
		totalWidth += kt.glyphMetrics.GetAdvanceWidthForRune(
			runes[i],
		)
	}

	return totalWidth
}

// MeasureStringWithKerningScaled returns the width of a string in points at the given font size,
// including kerning adjustments.
func (kt *KerningTable) MeasureStringWithKerningScaled(
	text string,
	fontSize float64,
) float64 {
	if kt == nil || kt.glyphMetrics == nil {
		return 0
	}

	unitsPerEm := kt.glyphMetrics.UnitsPerEm()
	if unitsPerEm == 0 {
		return 0
	}

	fontUnits := kt.MeasureStringWithKerning(text)

	return float64(
		fontUnits,
	) * fontSize / float64(
		unitsPerEm,
	)
}

// KernPairCount returns the number of kerning pairs in the table.
func (kt *KerningTable) KernPairCount() int {
	if kt == nil {
		return 0
	}

	return len(kt.pairs)
}

// Font returns the underlying font.
func (kt *KerningTable) Font() *font.Font {
	if kt == nil {
		return nil
	}

	return kt.font
}

// GlyphMetrics returns the underlying glyph metrics.
func (kt *KerningTable) GlyphMetrics() *GlyphMetrics {
	if kt == nil {
		return nil
	}

	return kt.glyphMetrics
}
