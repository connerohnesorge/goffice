// run.go provides run-level text measurement for PDF text layout.
// A "run" is a contiguous sequence of text with uniform formatting (same font,
// size, and style). In Word documents, runs are represented by <w:r> elements.
// This file bridges the glyph and kerning layers with higher-level text layout.

package layout

import (
	"unicode"
)

// TextRun represents a styled text segment with uniform formatting.
// All text within a run shares the same font, size, and styling properties.
type TextRun struct {
	// Text is the content of the run.
	Text string

	// GlyphMetrics provides glyph advance width lookup for the run's font.
	GlyphMetrics *GlyphMetrics

	// KerningTable provides kerning adjustments for character pairs.
	// May be nil if kerning is not available or disabled.
	KerningTable *KerningTable

	// FontSize is the font size in points.
	FontSize float64

	// CharacterSpacing is extra space added between each character, in points.
	// Positive values increase spacing, negative values decrease it.
	CharacterSpacing float64

	// WordSpacing is extra space added for space characters only, in points.
	// This is in addition to the normal space width and CharacterSpacing.
	WordSpacing float64

	// HorizontalScaling is a scaling factor applied to character widths.
	// 1.0 = normal, 0.5 = condensed to half width, 2.0 = expanded to double width.
	// Default should be 1.0.
	HorizontalScaling float64

	// Bold indicates if the text is bold.
	Bold bool

	// Italic indicates if the text is italic.
	Italic bool

	// Color is the text color (hex string, e.g. "FF0000").
	Color string

	// Underline indicates the underline style.
	Underline string

	// Strike indicates if the text has a strikethrough.
	Strike bool

	// Subscript indicates if the text is subscript.
	Subscript bool

	// Superscript indicates if the text is superscript.
	Superscript bool
}

// NewTextRun creates a new TextRun with default settings.
// HorizontalScaling defaults to 1.0 (normal width).
func NewTextRun(
	text string,
	gm *GlyphMetrics,
	fontSize float64,
) *TextRun {
	return &TextRun{
		Text:              text,
		GlyphMetrics:      gm,
		FontSize:          fontSize,
		HorizontalScaling: 1.0,
	}
}

// NewTextRunWithKerning creates a new TextRun with kerning support.
func NewTextRunWithKerning(
	text string,
	gm *GlyphMetrics,
	kt *KerningTable,
	fontSize float64,
) *TextRun {
	return &TextRun{
		Text:              text,
		GlyphMetrics:      gm,
		KerningTable:      kt,
		FontSize:          fontSize,
		HorizontalScaling: 1.0,
	}
}

// Width calculates the total width of the text run in points.
// This accounts for:
// - Glyph advance widths
// - Kerning adjustments (if KerningTable is set)
// - Character spacing
// - Word spacing (applied to space characters)
// - Horizontal scaling
func (r *TextRun) Width() float64 {
	return MeasureRun(r)
}

// MeasureRun calculates the width of a text run in points.
// This is the main entry point for run width calculation.
func MeasureRun(run *TextRun) float64 {
	if run == nil || len(run.Text) == 0 {
		return 0
	}

	if run.GlyphMetrics == nil {
		return 0
	}

	// Get the scaling factor (default to 1.0 if not set)
	scaling := run.HorizontalScaling
	if scaling == 0 {
		scaling = 1.0
	}

	runes := []rune(run.Text)
	if len(runes) == 0 {
		return 0
	}

	unitsPerEm := run.GlyphMetrics.UnitsPerEm()
	if unitsPerEm == 0 {
		return 0
	}

	// Calculate width in font units first
	var totalFontUnits int

	// First character width
	totalFontUnits = run.GlyphMetrics.GetAdvanceWidthForRune(
		runes[0],
	)

	// Add widths for remaining characters plus kerning
	for i := 1; i < len(runes); i++ {
		// Add kerning adjustment for the pair
		if run.KerningTable != nil {
			totalFontUnits += run.KerningTable.GetKerning(
				runes[i-1],
				runes[i],
			)
		}

		// Add advance width of current character
		totalFontUnits += run.GlyphMetrics.GetAdvanceWidthForRune(
			runes[i],
		)
	}

	// Convert font units to points
	baseWidth := float64(
		totalFontUnits,
	) * run.FontSize / float64(
		unitsPerEm,
	)

	// Apply horizontal scaling to the base glyph widths
	baseWidth *= scaling

	// Add character spacing for all characters except the last
	// (character spacing is the extra space *after* each character)
	if run.CharacterSpacing != 0 &&
		len(runes) > 0 {
		// Character spacing is added after each character except the last
		// (or we can consider it's between characters)
		baseWidth += run.CharacterSpacing * float64(
			len(runes)-1,
		)
	}

	// Add word spacing for space characters
	if run.WordSpacing != 0 {
		spaceCount := countSpaces(run.Text)
		baseWidth += run.WordSpacing * float64(
			spaceCount,
		)
	}

	return baseWidth
}

// countSpaces counts the number of space characters in a string.
// This includes regular spaces and other Unicode space characters.
func countSpaces(text string) int {
	count := 0
	for _, r := range text {
		if isSpaceChar(r) {
			count++
		}
	}

	return count
}

// isSpaceChar returns true if the rune is a space character that should
// receive word spacing adjustment.
func isSpaceChar(r rune) bool {
	// Primary space character
	if r == ' ' {
		return true
	}

	// Other common space characters that should receive word spacing
	switch r {
	case '\u2000', // En Quad
		'\u2001', // Em Quad
		'\u2002', // En Space
		'\u2003', // Em Space
		'\u2004', // Three-Per-Em Space
		'\u2005', // Four-Per-Em Space
		'\u2006', // Six-Per-Em Space
		'\u2008', // Punctuation Space
		'\u2009', // Thin Space
		'\u200A', // Hair Space
		'\u205F', // Medium Mathematical Space
		'\u3000': // Ideographic Space
		return true
	}

	return false
}

// MeasureRunWithOptions provides more control over run measurement.
type RunMeasureOptions struct {
	// IncludeKerning controls whether kerning is applied.
	IncludeKerning bool

	// IncludeCharSpacing controls whether character spacing is applied.
	IncludeCharSpacing bool

	// IncludeWordSpacing controls whether word spacing is applied.
	IncludeWordSpacing bool

	// IncludeScaling controls whether horizontal scaling is applied.
	IncludeScaling bool
}

// DefaultRunMeasureOptions returns options with all measurements enabled.
func DefaultRunMeasureOptions() RunMeasureOptions {
	return RunMeasureOptions{
		IncludeKerning:     true,
		IncludeCharSpacing: true,
		IncludeWordSpacing: true,
		IncludeScaling:     true,
	}
}

// MeasureRunWithOptions calculates run width with selective options.
func MeasureRunWithOptions(
	run *TextRun,
	opts RunMeasureOptions,
) float64 {
	if run == nil || len(run.Text) == 0 {
		return 0
	}

	if run.GlyphMetrics == nil {
		return 0
	}

	// Get the scaling factor
	scaling := 1.0
	if opts.IncludeScaling &&
		run.HorizontalScaling != 0 {
		scaling = run.HorizontalScaling
	}

	runes := []rune(run.Text)
	if len(runes) == 0 {
		return 0
	}

	unitsPerEm := run.GlyphMetrics.UnitsPerEm()
	if unitsPerEm == 0 {
		return 0
	}

	// Calculate width in font units first
	var totalFontUnits int

	// First character width
	totalFontUnits = run.GlyphMetrics.GetAdvanceWidthForRune(
		runes[0],
	)

	// Add widths for remaining characters plus kerning
	for i := 1; i < len(runes); i++ {
		// Add kerning adjustment for the pair
		if opts.IncludeKerning &&
			run.KerningTable != nil {
			totalFontUnits += run.KerningTable.GetKerning(
				runes[i-1],
				runes[i],
			)
		}

		// Add advance width of current character
		totalFontUnits += run.GlyphMetrics.GetAdvanceWidthForRune(
			runes[i],
		)
	}

	// Convert font units to points
	baseWidth := float64(
		totalFontUnits,
	) * run.FontSize / float64(
		unitsPerEm,
	)

	// Apply horizontal scaling to the base glyph widths
	baseWidth *= scaling

	// Add character spacing
	if opts.IncludeCharSpacing &&
		run.CharacterSpacing != 0 &&
		len(runes) > 0 {
		baseWidth += run.CharacterSpacing * float64(
			len(runes)-1,
		)
	}

	// Add word spacing for space characters
	if opts.IncludeWordSpacing &&
		run.WordSpacing != 0 {
		spaceCount := countSpaces(run.Text)
		baseWidth += run.WordSpacing * float64(
			spaceCount,
		)
	}

	return baseWidth
}

// RunSegment represents a portion of a run that may need different handling.
// This is useful for font fallback scenarios where different parts of a run
// may need different fonts.
type RunSegment struct {
	// Text is the text content of this segment.
	Text string

	// StartIndex is the rune index where this segment starts in the original run.
	StartIndex int

	// EndIndex is the rune index where this segment ends (exclusive).
	EndIndex int

	// NeedsFallback indicates if this segment contains characters
	// not covered by the primary font.
	NeedsFallback bool

	// MissingRunes contains the runes that are missing from the primary font.
	MissingRunes []rune
}

// SegmentRunForFallback splits a run into segments based on glyph coverage.
// Contiguous sequences of covered or uncovered characters form segments.
// This is useful for implementing font fallback.
func SegmentRunForFallback(
	run *TextRun,
) []RunSegment {
	if run == nil || len(run.Text) == 0 ||
		run.GlyphMetrics == nil {
		return nil
	}

	runes := []rune(run.Text)
	if len(runes) == 0 {
		return nil
	}

	var segments []RunSegment
	var currentSegment RunSegment
	var currentMissing bool
	var missingRunes []rune

	for i, r := range runes {
		hasCoverage := run.GlyphMetrics.HasGlyph(
			r,
		)
		needsFallback := !hasCoverage

		if i == 0 {
			// Initialize first segment
			currentMissing = needsFallback
			currentSegment = RunSegment{
				StartIndex:    0,
				NeedsFallback: needsFallback,
			}
			if needsFallback {
				missingRunes = append(
					missingRunes,
					r,
				)
			}

			continue
		}

		if needsFallback != currentMissing {
			// Segment boundary - save current segment
			currentSegment.EndIndex = i
			currentSegment.Text = string(
				runes[currentSegment.StartIndex:currentSegment.EndIndex],
			)
			if currentMissing {
				currentSegment.MissingRunes = missingRunes
			}
			segments = append(
				segments,
				currentSegment,
			)

			// Start new segment
			currentMissing = needsFallback
			missingRunes = nil
			currentSegment = RunSegment{
				StartIndex:    i,
				NeedsFallback: needsFallback,
			}
		}

		if needsFallback {
			missingRunes = append(missingRunes, r)
		}
	}

	// Don't forget the last segment
	currentSegment.EndIndex = len(runes)
	currentSegment.Text = string(
		runes[currentSegment.StartIndex:currentSegment.EndIndex],
	)
	if currentMissing {
		currentSegment.MissingRunes = missingRunes
	}
	segments = append(segments, currentSegment)

	return segments
}

// RunMeasurer provides a reusable run measurement context.
// It caches computed values for efficiency when measuring multiple runs
// with the same font.
type RunMeasurer struct {
	glyphMetrics *GlyphMetrics
	kerningTable *KerningTable
	fontSize     float64
}

// NewRunMeasurer creates a new RunMeasurer for a given font at a specific size.
func NewRunMeasurer(
	gm *GlyphMetrics,
	kt *KerningTable,
	fontSize float64,
) *RunMeasurer {
	return &RunMeasurer{
		glyphMetrics: gm,
		kerningTable: kt,
		fontSize:     fontSize,
	}
}

// Measure calculates the width of text in points using the measurer's settings.
func (rm *RunMeasurer) Measure(
	text string,
) float64 {
	if rm == nil || rm.glyphMetrics == nil {
		return 0
	}

	run := &TextRun{
		Text:              text,
		GlyphMetrics:      rm.glyphMetrics,
		KerningTable:      rm.kerningTable,
		FontSize:          rm.fontSize,
		HorizontalScaling: 1.0,
	}

	return MeasureRun(run)
}

// MeasureWithSpacing calculates the width with additional spacing options.
func (rm *RunMeasurer) MeasureWithSpacing(
	text string,
	charSpacing, wordSpacing float64,
) float64 {
	if rm == nil || rm.glyphMetrics == nil {
		return 0
	}

	run := &TextRun{
		Text:              text,
		GlyphMetrics:      rm.glyphMetrics,
		KerningTable:      rm.kerningTable,
		FontSize:          rm.fontSize,
		CharacterSpacing:  charSpacing,
		WordSpacing:       wordSpacing,
		HorizontalScaling: 1.0,
	}

	return MeasureRun(run)
}

// MeasureFunc returns a MeasureFunc compatible with SplitIntoLines.
func (rm *RunMeasurer) MeasureFunc() MeasureFunc {
	return rm.Measure
}

// SetFontSize updates the font size for subsequent measurements.
func (rm *RunMeasurer) SetFontSize(
	fontSize float64,
) {
	rm.fontSize = fontSize
}

// FontSize returns the current font size.
func (rm *RunMeasurer) FontSize() float64 {
	return rm.fontSize
}

// IsWhitespaceRun returns true if the run contains only whitespace characters.
func IsWhitespaceRun(run *TextRun) bool {
	if run == nil || len(run.Text) == 0 {
		return true
	}

	for _, r := range run.Text {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

// IsEmptyRun returns true if the run has no text content.
func IsEmptyRun(run *TextRun) bool {
	return run == nil || len(run.Text) == 0
}

// RunWidthInfo contains detailed width breakdown for a run.
type RunWidthInfo struct {
	// TotalWidth is the complete width in points.
	TotalWidth float64

	// BaseWidth is the width from glyph advances only (before scaling).
	BaseWidth float64

	// KerningAdjustment is the total kerning adjustment in points.
	KerningAdjustment float64

	// CharSpacingTotal is the total character spacing in points.
	CharSpacingTotal float64

	// WordSpacingTotal is the total word spacing in points.
	WordSpacingTotal float64

	// ScaledWidth is the width after horizontal scaling is applied.
	ScaledWidth float64
}

// MeasureRunDetailed returns detailed width information for a run.
func MeasureRunDetailed(
	run *TextRun,
) RunWidthInfo {
	info := RunWidthInfo{}

	if run == nil || len(run.Text) == 0 ||
		run.GlyphMetrics == nil {
		return info
	}

	runes := []rune(run.Text)
	if len(runes) == 0 {
		return info
	}

	unitsPerEm := run.GlyphMetrics.UnitsPerEm()
	if unitsPerEm == 0 {
		return info
	}

	scaling := run.HorizontalScaling
	if scaling == 0 {
		scaling = 1.0
	}

	// Calculate base width in font units (glyph advances only)
	var baseFontUnits int
	for _, r := range runes {
		baseFontUnits += run.GlyphMetrics.GetAdvanceWidthForRune(
			r,
		)
	}

	// Calculate kerning adjustment in font units
	var kerningFontUnits int
	if run.KerningTable != nil {
		for i := 1; i < len(runes); i++ {
			kerningFontUnits += run.KerningTable.GetKerning(
				runes[i-1],
				runes[i],
			)
		}
	}

	// Convert to points
	info.BaseWidth = float64(
		baseFontUnits,
	) * run.FontSize / float64(
		unitsPerEm,
	)
	info.KerningAdjustment = float64(
		kerningFontUnits,
	) * run.FontSize / float64(
		unitsPerEm,
	)

	// Character spacing (between characters, so n-1 for n characters)
	if run.CharacterSpacing != 0 &&
		len(runes) > 0 {
		info.CharSpacingTotal = run.CharacterSpacing * float64(
			len(runes)-1,
		)
	}

	// Word spacing
	if run.WordSpacing != 0 {
		spaceCount := countSpaces(run.Text)
		info.WordSpacingTotal = run.WordSpacing * float64(
			spaceCount,
		)
	}

	// Calculate scaled width (base + kerning, then scaled)
	info.ScaledWidth = (info.BaseWidth + info.KerningAdjustment) * scaling

	// Total width = scaled width + spacing adjustments
	info.TotalWidth = info.ScaledWidth + info.CharSpacingTotal + info.WordSpacingTotal

	return info
}
