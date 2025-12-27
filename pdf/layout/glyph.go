// glyph.go provides glyph advance width lookup for accurate text measurement and layout.
// This file bridges the font parsing layer with the text layout engine, providing
// efficient glyph metric lookups needed for line breaking and text positioning.

package layout

import (
	"github.com/connerohnesorge/goffice-pdf/font"
)

// GlyphMetrics provides glyph-level measurements from a parsed font.
// It wraps a font.Font to provide efficient metric lookups for text layout.
type GlyphMetrics struct {
	// font is the underlying parsed font containing glyph data.
	font *font.Font

	// defaultWidth is the fallback advance width for missing glyphs.
	// This is typically set to the average character width or a reasonable default.
	defaultWidth int

	// spaceWidth caches the width of the space character for efficiency.
	spaceWidth int

	// unitsPerEm stores the font's units per em for scaled calculations.
	unitsPerEm int
}

// NewGlyphMetrics creates a new GlyphMetrics instance from a parsed font.
// If the font is nil, returns nil.
func NewGlyphMetrics(f *font.Font) *GlyphMetrics {
	if f == nil {
		return nil
	}

	gm := &GlyphMetrics{
		font:       f,
		unitsPerEm: f.Metrics.UnitsPerEm,
	}

	// Calculate default width as fallback for missing glyphs.
	// Use .notdef glyph width if available, otherwise use half of unitsPerEm.
	gm.defaultWidth = gm.calculateDefaultWidth()

	// Cache space width for efficiency.
	gm.spaceWidth = gm.GetAdvanceWidthForRune(' ')
	if gm.spaceWidth == 0 {
		// Fallback: typical space is about 1/4 of em
		gm.spaceWidth = gm.unitsPerEm / 4
	}

	return gm
}

// calculateDefaultWidth determines a reasonable default width for missing glyphs.
// The priority is:
// 1. Width of the replacement character (U+FFFD)
// 2. Width of 'x' (typical character)
// 3. Half of unitsPerEm
func (gm *GlyphMetrics) calculateDefaultWidth() int {
	// Try replacement character first
	if w := gm.font.GlyphWidth('\uFFFD'); w > 0 {
		return w
	}

	// Try lowercase 'x' as a typical character
	if w := gm.font.GlyphWidth('x'); w > 0 {
		return w
	}

	// Try uppercase 'H' as another common reference
	if w := gm.font.GlyphWidth('H'); w > 0 {
		return w
	}

	// Fall back to half of unitsPerEm
	if gm.unitsPerEm > 0 {
		return gm.unitsPerEm / 2
	}

	// Ultimate fallback (shouldn't happen with valid fonts)
	return 500
}

// GetGlyphID returns the glyph ID for a Unicode codepoint.
// In OpenType/TrueType fonts, this would typically involve cmap table lookup.
// Since the font package pre-maps characters to metrics, we return a pseudo-ID.
// Returns 0 (the .notdef glyph ID) if the character is not in the font.
func (gm *GlyphMetrics) GetGlyphID(
	r rune,
) uint16 {
	if gm.font == nil {
		return 0
	}

	// Check if the character exists in the font's glyph data
	if _, ok := gm.font.GlyphData[r]; ok {
		// Return a non-zero ID to indicate the glyph exists.
		// The actual glyph ID is internal to the font parsing;
		// we use a hash of the rune as a pseudo-ID for external reference.
		// For practical purposes, non-zero means the glyph exists.
		return uint16((uint32(r) % 65534) + 1)
	}

	// Glyph not found, return .notdef (glyph ID 0)
	return 0
}

// HasGlyph returns true if the font contains a glyph for the given rune.
func (gm *GlyphMetrics) HasGlyph(r rune) bool {
	if gm.font == nil {
		return false
	}

	_, ok := gm.font.GlyphData[r]

	return ok
}

// GetAdvanceWidth returns the advance width for a glyph ID in font design units.
// Since we pre-map characters to metrics, this accepts a rune that was used
// to obtain the glyph ID. For missing glyphs, returns the default width.
func (gm *GlyphMetrics) GetAdvanceWidth(
	glyphID uint16,
) int {
	// Since our glyph IDs are pseudo-IDs based on runes, we can't directly
	// look up by ID. This method is provided for API consistency.
	// In practice, use GetAdvanceWidthForRune for direct lookups.
	if glyphID == 0 {
		return gm.defaultWidth
	}

	// Return default width since we can't reverse the ID mapping
	return gm.defaultWidth
}

// GetAdvanceWidthForRune returns the advance width for a Unicode codepoint
// in font design units. This is the combined cmap + hmtx lookup.
// For missing glyphs, returns a default width based on the font metrics.
func (gm *GlyphMetrics) GetAdvanceWidthForRune(
	r rune,
) int {
	if gm.font == nil {
		return gm.defaultWidth
	}

	// Use the font's GlyphWidth method which already handles the lookup
	width := gm.font.GlyphWidth(r)
	if width > 0 {
		return width
	}

	// Glyph not found, return default width
	return gm.defaultWidth
}

// GetLeftBearing returns the left side bearing for a rune in font design units.
// The left side bearing is the horizontal distance from the glyph origin
// to the leftmost point of the glyph outline.
func (gm *GlyphMetrics) GetLeftBearing(
	r rune,
) int {
	if gm.font == nil {
		return 0
	}

	if metrics, ok := gm.font.GlyphData[r]; ok {
		return metrics.LeftBearing
	}

	return 0
}

// MeasureString returns the total advance width for a string in font design units.
// This is the sum of all individual glyph advance widths.
func (gm *GlyphMetrics) MeasureString(
	text string,
) int {
	if gm.font == nil || len(text) == 0 {
		return 0
	}

	var totalWidth int
	for _, r := range text {
		totalWidth += gm.GetAdvanceWidthForRune(r)
	}

	return totalWidth
}

// MeasureStringScaled returns the width of a string in points at the given font size.
// This converts from font design units to actual display units.
func (gm *GlyphMetrics) MeasureStringScaled(
	text string,
	fontSize float64,
) float64 {
	if gm.unitsPerEm == 0 {
		return 0
	}

	fontUnits := gm.MeasureString(text)
	width := float64(
		fontUnits,
	) * fontSize / float64(
		gm.unitsPerEm,
	)

	return width
}

// MeasureRuneScaled returns the width of a single rune in points at the given font size.
func (gm *GlyphMetrics) MeasureRuneScaled(
	r rune,
	fontSize float64,
) float64 {
	if gm.unitsPerEm == 0 {
		return 0
	}

	fontUnits := gm.GetAdvanceWidthForRune(r)
	width := float64(
		fontUnits,
	) * fontSize / float64(
		gm.unitsPerEm,
	)

	return width
}

// UnitsPerEm returns the font's units per em value.
// This is typically 1000 for PostScript-based fonts and 2048 for TrueType.
func (gm *GlyphMetrics) UnitsPerEm() int {
	return gm.unitsPerEm
}

// SpaceWidth returns the cached width of the space character in font design units.
func (gm *GlyphMetrics) SpaceWidth() int {
	return gm.spaceWidth
}

// DefaultWidth returns the default width used for missing glyphs in font design units.
func (gm *GlyphMetrics) DefaultWidth() int {
	return gm.defaultWidth
}

// Ascender returns the font's ascender value in font design units.
// The ascender is the height above the baseline.
func (gm *GlyphMetrics) Ascender() int {
	if gm.font == nil {
		return 0
	}

	return gm.font.Metrics.Ascender
}

// Descender returns the font's descender value in font design units.
// The descender is the depth below the baseline (typically negative).
func (gm *GlyphMetrics) Descender() int {
	if gm.font == nil {
		return 0
	}

	return gm.font.Metrics.Descender
}

// LineGap returns the font's line gap value in font design units.
// This is the recommended additional spacing between lines.
func (gm *GlyphMetrics) LineGap() int {
	if gm.font == nil {
		return 0
	}

	return gm.font.Metrics.LineGap
}

// LineHeight returns the recommended line height in font design units.
// This is calculated as Ascender - Descender + LineGap.
func (gm *GlyphMetrics) LineHeight() int {
	if gm.font == nil {
		return 0
	}

	return gm.font.Metrics.Ascender - gm.font.Metrics.Descender + gm.font.Metrics.LineGap
}

// LineHeightScaled returns the recommended line height in points at the given font size.
func (gm *GlyphMetrics) LineHeightScaled(
	fontSize float64,
) float64 {
	if gm.unitsPerEm == 0 {
		return fontSize * 1.2 // Fallback: typical line height
	}

	return float64(
		gm.LineHeight(),
	) * fontSize / float64(
		gm.unitsPerEm,
	)
}

// Font returns the underlying font.Font for access to additional metadata.
func (gm *GlyphMetrics) Font() *font.Font {
	return gm.font
}

// TextMeasurer provides a MeasureFunc-compatible interface for use with
// the line breaking algorithm. It wraps GlyphMetrics with a specific font size.
type TextMeasurer struct {
	metrics  *GlyphMetrics
	fontSize float64
}

// NewTextMeasurer creates a TextMeasurer for a font at a specific size.
// The returned measurer can be used with SplitIntoLines.
func NewTextMeasurer(
	gm *GlyphMetrics,
	fontSize float64,
) *TextMeasurer {
	return &TextMeasurer{
		metrics:  gm,
		fontSize: fontSize,
	}
}

// Measure returns the width of text in points.
// This method satisfies the MeasureFunc type for use with line breaking.
func (tm *TextMeasurer) Measure(
	text string,
) float64 {
	if tm.metrics == nil {
		return 0
	}

	return tm.metrics.MeasureStringScaled(
		text,
		tm.fontSize,
	)
}

// MeasureFunc returns a MeasureFunc for use with SplitIntoLines.
func (tm *TextMeasurer) MeasureFunc() MeasureFunc {
	return tm.Measure
}

// FontSize returns the font size in points.
func (tm *TextMeasurer) FontSize() float64 {
	return tm.fontSize
}

// SetFontSize changes the font size for measurements.
func (tm *TextMeasurer) SetFontSize(
	fontSize float64,
) {
	tm.fontSize = fontSize
}

// Metrics returns the underlying GlyphMetrics.
func (tm *TextMeasurer) Metrics() *GlyphMetrics {
	return tm.metrics
}

// MissingGlyphInfo contains information about a missing glyph.
type MissingGlyphInfo struct {
	// Rune is the Unicode codepoint that was not found.
	Rune rune
	// Position is the byte position in the original string.
	Position int
	// RunePosition is the rune index in the original string.
	RunePosition int
}

// FindMissingGlyphs scans a string and returns information about any characters
// that don't have glyphs in the font. This is useful for font fallback logic.
func (gm *GlyphMetrics) FindMissingGlyphs(
	text string,
) []MissingGlyphInfo {
	if gm.font == nil {
		return nil
	}

	var missing []MissingGlyphInfo
	bytePos := 0
	runePos := 0

	for _, r := range text {
		if !gm.HasGlyph(r) {
			missing = append(
				missing,
				MissingGlyphInfo{
					Rune:         r,
					Position:     bytePos,
					RunePosition: runePos,
				},
			)
		}
		bytePos += len(string(r))
		runePos++
	}

	return missing
}

// Coverage returns the percentage of runes in the text that have glyphs.
// Returns 1.0 (100%) if all characters are covered, 0.0 if none are.
func (gm *GlyphMetrics) Coverage(
	text string,
) float64 {
	if len(text) == 0 {
		return 1.0
	}

	if gm.font == nil {
		return 0.0
	}

	total := 0
	covered := 0

	for _, r := range text {
		total++
		if gm.HasGlyph(r) {
			covered++
		}
	}

	if total == 0 {
		return 1.0
	}

	return float64(covered) / float64(total)
}
