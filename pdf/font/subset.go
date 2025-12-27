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

	// Determine the old glyph IDs we need to include
	// Start by collecting the glyphs for our runes
	oldGlyphIDs := make(map[uint16]bool)
	oldGlyphIDs[0] = true // Always include .notdef

	// Map runes to old glyph IDs using the original font's cmap
	runeToOldGlyphID := make(map[rune]uint16)
	for _, r := range runes {
		if glyphData, ok := b.font.GlyphData[r]; ok {
			// We need to find the original glyph ID
			// Since we have the metrics, we search hmtx
			oldGID := b.findGlyphIDForRune(
				r,
				tables,
			)
			if oldGID != 0 || r == 0 {
				oldGlyphIDs[oldGID] = true
				runeToOldGlyphID[r] = oldGID
			}
			_ = glyphData
		}
	}

	// For glyf/loca subsetting, we need to handle composite glyphs
	// which reference other glyphs
	if tables["glyf"] != nil &&
		tables["loca"] != nil {
		// Expand glyph set to include composite dependencies
		if err := b.expandCompositeGlyphs(
			tables,
			oldGlyphIDs,
		); err != nil {
			return nil, fmt.Errorf(
				"%w: failed to expand composite glyphs: %v",
				ErrSubsetFailed,
				err,
			)
		}
	}

	// Check if subsetting is worthwhile
	numGlyphs := getNumGlyphsFromMaxp(
		tables["maxp"],
	)
	if len(oldGlyphIDs) > numGlyphs/2 {
		// If using more than half the glyphs, embedding the whole font
		// is more efficient than subsetting
		return nil, fmt.Errorf(
			"%w: using %d of %d glyphs, full embedding is more efficient",
			ErrSubsetFailed,
			len(oldGlyphIDs),
			numGlyphs,
		)
	}

	// Build the subset font with proper table rewriting
	return b.buildSubsetFont(
		tables,
		runes,
		glyphMapping,
		oldGlyphIDs,
		runeToOldGlyphID,
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
// This implementation properly rewrites glyf, loca, hmtx, and other tables.
func (b *SubsetBuilder) buildSubsetFont(
	tables map[string][]byte,
	runes []rune,
	glyphMapping map[rune]uint16,
	oldGlyphIDs map[uint16]bool,
	runeToOldGlyphID map[rune]uint16,
) ([]byte, error) {
	// Check if glyf table exists (TrueType outlines)
	hasGlyf := tables["glyf"] != nil &&
		tables["loca"] != nil

	// Create old-to-new glyph ID mapping
	oldToNewGlyphID := make(map[uint16]uint16)
	newGlyphID := uint16(0)

	// Sort old glyph IDs for consistent output
	sortedOldGIDs := make(
		[]uint16,
		0,
		len(oldGlyphIDs),
	)
	for gid := range oldGlyphIDs {
		sortedOldGIDs = append(sortedOldGIDs, gid)
	}
	sort.Slice(
		sortedOldGIDs,
		func(i, j int) bool {
			return sortedOldGIDs[i] < sortedOldGIDs[j]
		},
	)

	// Build mapping
	for _, oldGID := range sortedOldGIDs {
		oldToNewGlyphID[oldGID] = newGlyphID
		newGlyphID++
	}

	newGlyphCount := newGlyphID

	// Build modified tables
	newTables := make(map[string][]byte)

	// Copy tables that don't need modification
	copyTables := []string{
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

	// Get loca format from head table
	locaFormat := b.getLocaFormat(tables["head"])

	// Rewrite glyf and loca tables with only used glyphs
	if hasGlyf {
		newGlyf, newLoca, err := b.rewriteGlyfLoca(
			tables["glyf"],
			tables["loca"],
			locaFormat,
			oldToNewGlyphID,
			sortedOldGIDs,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"%w: failed to rewrite glyf/loca: %v",
				ErrSubsetFailed,
				err,
			)
		}
		newTables["glyf"] = newGlyf
		newTables["loca"] = newLoca
	}

	// Rewrite hmtx with only used glyph metrics
	newHmtx, numHMetrics, err := b.rewriteHmtx(
		tables["hmtx"],
		tables["hhea"],
		oldToNewGlyphID,
		sortedOldGIDs,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: failed to rewrite hmtx: %v",
			ErrSubsetFailed,
			err,
		)
	}
	newTables["hmtx"] = newHmtx

	// Update hhea with new numberOfHMetrics
	newTables["hhea"] = b.updateHhea(
		tables["hhea"],
		numHMetrics,
	)

	// Update maxp with new glyph count
	newTables["maxp"] = b.createMaxpSubset(
		tables["maxp"],
		newGlyphCount,
	)

	// Copy head table (we'll update checksum later if needed)
	if head, ok := tables["head"]; ok {
		newTables["head"] = head
	}

	// Create a new cmap that maps our subset runes to new glyph IDs
	newCmap := b.createCmapSubsetWithMapping(
		runes,
		glyphMapping,
		runeToOldGlyphID,
		oldToNewGlyphID,
	)
	newTables["cmap"] = newCmap

	// Include CFF table if present (OpenType CFF font)
	if cff, ok := tables["CFF "]; ok {
		newTables["CFF "] = cff
	}

	// Build the final font file
	return b.assembleTrueTypeFont(newTables)
}

// findGlyphIDForRune finds the original glyph ID for a rune by parsing the cmap.
func (b *SubsetBuilder) findGlyphIDForRune(
	r rune,
	tables map[string][]byte,
) uint16 {
	cmapData, ok := tables["cmap"]
	if !ok {
		return 0
	}

	if len(cmapData) < 4 {
		return 0
	}

	reader := bytes.NewReader(cmapData)
	var version, numTables uint16
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&version,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&numTables,
	)

	// Find best cmap subtable
	var bestOffset uint32
	for range numTables {
		var platformID, encodingID uint16
		var offset uint32
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&platformID,
		)
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&encodingID,
		)
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&offset,
		)

		// Prefer Windows Unicode
		if platformID == 3 &&
			(encodingID == 1 || encodingID == 10) {
			bestOffset = offset
			break
		}
	}

	if bestOffset == 0 ||
		int(bestOffset) >= len(cmapData) {
		return 0
	}

	subtable := cmapData[bestOffset:]
	if len(subtable) < 2 {
		return 0
	}

	format := binary.BigEndian.Uint16(
		subtable[0:2],
	)
	switch format {
	case 4:
		return b.findGlyphInFormat4(subtable, r)
	case 12:
		return b.findGlyphInFormat12(subtable, r)
	default:
		return 0
	}
}

// findGlyphInFormat4 finds a glyph ID in a format 4 cmap subtable.
func (b *SubsetBuilder) findGlyphInFormat4(
	data []byte,
	r rune,
) uint16 {
	if len(data) < 14 || r > 0xFFFF {
		return 0
	}

	c := uint16(r)
	reader := bytes.NewReader(data)

	var format, length, language, segCountX2 uint16
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&format,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&length,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&language,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&segCountX2,
	)

	segCount := int(segCountX2 / 2)
	if segCount == 0 {
		return 0
	}

	// Skip searchRange, entrySelector, rangeShift
	reader.Seek(14, 0)

	// Read endCodes
	endCodes := make([]uint16, segCount)
	for i := range segCount {
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&endCodes[i],
		)
	}

	// Skip reservedPad
	var pad uint16
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&pad,
	)

	// Read startCodes
	startCodes := make([]uint16, segCount)
	for i := range segCount {
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&startCodes[i],
		)
	}

	// Read idDeltas
	idDeltas := make([]int16, segCount)
	for i := range segCount {
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&idDeltas[i],
		)
	}

	// Find the segment containing our character
	for seg := range segCount {
		if c >= startCodes[seg] &&
			c <= endCodes[seg] {
			return uint16(
				int16(c) + idDeltas[seg],
			)
		}
	}

	return 0
}

// findGlyphInFormat12 finds a glyph ID in a format 12 cmap subtable.
func (b *SubsetBuilder) findGlyphInFormat12(
	data []byte,
	r rune,
) uint16 {
	if len(data) < 16 {
		return 0
	}

	c := uint32(r)
	reader := bytes.NewReader(data)

	var format uint16
	var reserved uint16
	var length, language, numGroups uint32
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&format,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&reserved,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&length,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&language,
	)
	_ = binary.Read(
		reader,
		binary.BigEndian,
		&numGroups,
	)

	for range numGroups {
		var startCharCode, endCharCode, startGlyphID uint32
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&startCharCode,
		)
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&endCharCode,
		)
		_ = binary.Read(
			reader,
			binary.BigEndian,
			&startGlyphID,
		)

		if c >= startCharCode &&
			c <= endCharCode {
			return uint16(
				startGlyphID + (c - startCharCode),
			)
		}
	}

	return 0
}

// expandCompositeGlyphs expands the glyph set to include all composite dependencies.
func (b *SubsetBuilder) expandCompositeGlyphs(
	tables map[string][]byte,
	glyphIDs map[uint16]bool,
) error {
	glyfData := tables["glyf"]
	locaData := tables["loca"]
	headData := tables["head"]

	if glyfData == nil || locaData == nil ||
		headData == nil {
		return nil
	}

	locaFormat := b.getLocaFormat(headData)
	numGlyphs := getNumGlyphsFromMaxp(
		tables["maxp"],
	)

	// Parse loca table to get glyph offsets
	glyphOffsets, err := b.parseLocaTable(
		locaData,
		locaFormat,
		numGlyphs,
	)
	if err != nil {
		return err
	}

	// Iteratively expand composite glyphs
	// We need to repeat until no new glyphs are added
	changed := true
	for changed {
		changed = false
		currentGlyphs := make(
			[]uint16,
			0,
			len(glyphIDs),
		)
		for gid := range glyphIDs {
			currentGlyphs = append(
				currentGlyphs,
				gid,
			)
		}

		for _, gid := range currentGlyphs {
			if int(gid) >= len(glyphOffsets)-1 {
				continue
			}

			offset := glyphOffsets[gid]
			nextOffset := glyphOffsets[gid+1]

			if offset >= nextOffset ||
				int(offset) >= len(glyfData) {
				continue
			}

			glyphData := glyfData[offset:nextOffset]
			if len(glyphData) < 10 {
				continue // Empty or invalid glyph
			}

			// Check if this is a composite glyph
			numberOfContours := int16(
				binary.BigEndian.Uint16(
					glyphData[0:2],
				),
			)
			if numberOfContours >= 0 {
				continue // Simple glyph
			}

			// Parse composite glyph components
			componentGIDs := b.parseCompositeGlyph(
				glyphData,
			)
			for _, componentGID := range componentGIDs {
				if !glyphIDs[componentGID] {
					glyphIDs[componentGID] = true
					changed = true
				}
			}
		}
	}

	return nil
}

// getLocaFormat extracts the indexToLocFormat from the head table.
func (b *SubsetBuilder) getLocaFormat(
	headData []byte,
) int16 {
	if len(headData) < 54 {
		return 0
	}
	return int16(
		binary.BigEndian.Uint16(headData[50:52]),
	)
}

// parseLocaTable parses the loca table and returns glyph offsets.
func (b *SubsetBuilder) parseLocaTable(
	locaData []byte,
	locaFormat int16,
	numGlyphs int,
) ([]uint32, error) {
	offsets := make([]uint32, numGlyphs+1)

	if locaFormat == 0 {
		// Short format: offsets are uint16 * 2
		expectedSize := (numGlyphs + 1) * 2
		if len(locaData) < expectedSize {
			return nil, fmt.Errorf(
				"loca table too short for short format",
			)
		}

		for i := 0; i <= numGlyphs; i++ {
			offset := binary.BigEndian.Uint16(
				locaData[i*2 : i*2+2],
			)
			offsets[i] = uint32(offset) * 2
		}
	} else {
		// Long format: offsets are uint32
		expectedSize := (numGlyphs + 1) * 4
		if len(locaData) < expectedSize {
			return nil, fmt.Errorf("loca table too short for long format")
		}

		for i := 0; i <= numGlyphs; i++ {
			offsets[i] = binary.BigEndian.Uint32(locaData[i*4 : i*4+4])
		}
	}

	return offsets, nil
}

// parseCompositeGlyph extracts component glyph IDs from a composite glyph.
func (b *SubsetBuilder) parseCompositeGlyph(
	glyphData []byte,
) []uint16 {
	if len(glyphData) < 10 {
		return nil
	}

	var components []uint16
	offset := 10 // Skip header

	const (
		flagArg1And2AreWords   = 0x0001
		flagWeHaveAScale       = 0x0008
		flagMoreComponents     = 0x0020
		flagWeHaveAnXAndYScale = 0x0040
		flagWeHaveATwoByTwo    = 0x0080
		flagWeHaveInstructions = 0x0100
	)

	for offset+4 <= len(glyphData) {
		flags := binary.BigEndian.Uint16(
			glyphData[offset : offset+2],
		)
		glyphIndex := binary.BigEndian.Uint16(
			glyphData[offset+2 : offset+4],
		)
		components = append(
			components,
			glyphIndex,
		)
		offset += 4

		// Skip arguments
		if flags&flagArg1And2AreWords != 0 {
			offset += 4 // Two int16 arguments
		} else {
			offset += 2 // Two int8 arguments
		}

		// Skip transformation matrix
		if flags&flagWeHaveAScale != 0 {
			offset += 2 // One F2DOT14
		} else if flags&flagWeHaveAnXAndYScale != 0 {
			offset += 4 // Two F2DOT14
		} else if flags&flagWeHaveATwoByTwo != 0 {
			offset += 8 // Four F2DOT14
		}

		if flags&flagMoreComponents == 0 {
			break
		}
	}

	return components
}

// rewriteGlyfLoca creates new glyf and loca tables with only the subset glyphs.
func (b *SubsetBuilder) rewriteGlyfLoca(
	glyfData, locaData []byte,
	locaFormat int16,
	oldToNewGlyphID map[uint16]uint16,
	sortedOldGIDs []uint16,
) ([]byte, []byte, error) {
	numGlyphs := len(sortedOldGIDs)

	// Calculate number of glyphs from loca table size
	var numOriginalGlyphs int
	if locaFormat == 0 {
		numOriginalGlyphs = (len(locaData) / 2) - 1
	} else {
		numOriginalGlyphs = (len(locaData) / 4) - 1
	}

	if numOriginalGlyphs < 0 {
		numOriginalGlyphs = 0
	}

	// Parse original loca table
	originalOffsets, err := b.parseLocaTable(
		locaData,
		locaFormat,
		numOriginalGlyphs,
	)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"failed to parse loca table: %w",
			err,
		)
	}

	// Build new glyf data
	newGlyfBuf := &bytes.Buffer{}
	newOffsets := make([]uint32, numGlyphs+1)

	for i, oldGID := range sortedOldGIDs {
		newOffsets[i] = uint32(newGlyfBuf.Len())

		if int(oldGID) >= len(originalOffsets)-1 {
			continue
		}

		offset := originalOffsets[oldGID]
		nextOffset := originalOffsets[oldGID+1]

		if offset >= nextOffset ||
			int(offset) >= len(glyfData) {
			continue // Empty glyph
		}

		glyphBytes := glyfData[offset:nextOffset]

		// For composite glyphs, we need to update component references
		if len(glyphBytes) >= 10 {
			numberOfContours := int16(
				binary.BigEndian.Uint16(
					glyphBytes[0:2],
				),
			)
			if numberOfContours < 0 {
				// Composite glyph - rewrite component references
				glyphBytes = b.rewriteCompositeGlyph(
					glyphBytes,
					oldToNewGlyphID,
				)
			}
		}

		newGlyfBuf.Write(glyphBytes)
	}

	// Final offset
	newOffsets[numGlyphs] = uint32(
		newGlyfBuf.Len(),
	)

	// Build new loca table
	newLocaBuf := &bytes.Buffer{}
	if locaFormat == 0 {
		// Short format: check if all offsets fit in uint16
		maxOffset := newOffsets[numGlyphs]
		if maxOffset > 0x1FFFE {
			// Need to use long format instead
			locaFormat = 1
		}
	}

	if locaFormat == 0 {
		// Short format
		for _, offset := range newOffsets {
			_ = binary.Write(
				newLocaBuf,
				binary.BigEndian,
				uint16(offset/2),
			)
		}
	} else {
		// Long format
		for _, offset := range newOffsets {
			_ = binary.Write(
				newLocaBuf,
				binary.BigEndian,
				offset,
			)
		}
	}

	return newGlyfBuf.Bytes(), newLocaBuf.Bytes(), nil
}

// rewriteCompositeGlyph updates component glyph IDs in a composite glyph.
func (b *SubsetBuilder) rewriteCompositeGlyph(
	glyphData []byte,
	oldToNewGlyphID map[uint16]uint16,
) []byte {
	if len(glyphData) < 10 {
		return glyphData
	}

	result := make([]byte, len(glyphData))
	copy(result, glyphData)

	offset := 10 // Skip header

	const (
		flagArg1And2AreWords   = 0x0001
		flagWeHaveAScale       = 0x0008
		flagMoreComponents     = 0x0020
		flagWeHaveAnXAndYScale = 0x0040
		flagWeHaveATwoByTwo    = 0x0080
	)

	for offset+4 <= len(result) {
		flags := binary.BigEndian.Uint16(
			result[offset : offset+2],
		)
		oldGlyphIndex := binary.BigEndian.Uint16(
			result[offset+2 : offset+4],
		)

		// Remap glyph index
		if newGlyphIndex, ok := oldToNewGlyphID[oldGlyphIndex]; ok {
			binary.BigEndian.PutUint16(
				result[offset+2:offset+4],
				newGlyphIndex,
			)
		}

		offset += 4

		// Skip arguments
		if flags&flagArg1And2AreWords != 0 {
			offset += 4
		} else {
			offset += 2
		}

		// Skip transformation
		if flags&flagWeHaveAScale != 0 {
			offset += 2
		} else if flags&flagWeHaveAnXAndYScale != 0 {
			offset += 4
		} else if flags&flagWeHaveATwoByTwo != 0 {
			offset += 8
		}

		if flags&flagMoreComponents == 0 {
			break
		}
	}

	return result
}

// rewriteHmtx creates a new hmtx table with only the subset glyph metrics.
func (b *SubsetBuilder) rewriteHmtx(
	hmtxData, hheaData []byte,
	oldToNewGlyphID map[uint16]uint16,
	sortedOldGIDs []uint16,
) ([]byte, uint16, error) {
	if len(hheaData) < 36 {
		return nil, 0, fmt.Errorf(
			"hhea table too short",
		)
	}

	numHMetrics := binary.BigEndian.Uint16(
		hheaData[34:36],
	)

	buf := &bytes.Buffer{}

	for _, oldGID := range sortedOldGIDs {
		var advanceWidth uint16
		var lsb int16

		if oldGID < numHMetrics {
			// Read full metric
			offset := int(oldGID) * 4
			if offset+4 <= len(hmtxData) {
				advanceWidth = binary.BigEndian.Uint16(
					hmtxData[offset : offset+2],
				)
				lsb = int16(
					binary.BigEndian.Uint16(
						hmtxData[offset+2 : offset+4],
					),
				)
			}
		} else {
			// Use last advance width
			if numHMetrics > 0 {
				offset := int(numHMetrics-1) * 4
				if offset+2 <= len(hmtxData) {
					advanceWidth = binary.BigEndian.Uint16(
						hmtxData[offset : offset+2],
					)
				}
			}

			// Read LSB from remaining entries
			lsbOffset := int(numHMetrics)*4 + int(oldGID-numHMetrics)*2
			if lsbOffset+2 <= len(hmtxData) {
				lsb = int16(
					binary.BigEndian.Uint16(hmtxData[lsbOffset : lsbOffset+2]),
				)
			}
		}

		_ = binary.Write(
			buf,
			binary.BigEndian,
			advanceWidth,
		)
		_ = binary.Write(
			buf,
			binary.BigEndian,
			lsb,
		)
	}

	// New numberOfHMetrics is the number of glyphs with full metrics
	newNumHMetrics := uint16(len(sortedOldGIDs))

	return buf.Bytes(), newNumHMetrics, nil
}

// updateHhea updates the hhea table with new numberOfHMetrics.
func (b *SubsetBuilder) updateHhea(
	hheaData []byte,
	numHMetrics uint16,
) []byte {
	if len(hheaData) < 36 {
		return hheaData
	}

	result := make([]byte, len(hheaData))
	copy(result, hheaData)
	binary.BigEndian.PutUint16(
		result[34:36],
		numHMetrics,
	)

	return result
}

// createCmapSubsetWithMapping creates a cmap table mapping subset runes to new glyph IDs.
func (b *SubsetBuilder) createCmapSubsetWithMapping(
	runes []rune,
	glyphMapping map[rune]uint16,
	runeToOldGlyphID map[rune]uint16,
	oldToNewGlyphID map[uint16]uint16,
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
	) // platformID
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint16(1),
	) // encodingID
	_ = binary.Write(
		buf,
		binary.BigEndian,
		uint32(12),
	) // offset

	// Build Format 4 subtable with new glyph IDs
	bmpRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if r <= 0xFFFF {
			bmpRunes = append(bmpRunes, r)
		}
	}

	sort.Slice(bmpRunes, func(i, j int) bool {
		return bmpRunes[i] < bmpRunes[j]
	})

	type segment struct {
		startCode uint16
		endCode   uint16
		idDelta   int16
	}

	segments := make([]segment, 0)

	for _, r := range bmpRunes {
		oldGID, ok := runeToOldGlyphID[r]
		if !ok {
			continue
		}

		newGID, ok := oldToNewGlyphID[oldGID]
		if !ok {
			continue
		}

		delta := int16(newGID) - int16(r)
		segments = append(segments, segment{
			startCode: uint16(r),
			endCode:   uint16(r),
			idDelta:   delta,
		})
	}

	// Add terminating segment
	segments = append(segments, segment{
		startCode: 0xFFFF,
		endCode:   0xFFFF,
		idDelta:   1,
	})

	segCount := uint16(len(segments))
	segCountX2 := segCount * 2

	searchRange := uint16(1)
	entrySelector := uint16(0)
	for searchRange*2 <= segCount {
		searchRange *= 2
		entrySelector++
	}
	searchRange *= 2
	rangeShift := segCountX2 - searchRange

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

	// idRangeOffset array (all zeros)
	for range segments {
		_ = binary.Write(
			buf,
			binary.BigEndian,
			uint16(0),
		)
	}

	return buf.Bytes()
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
