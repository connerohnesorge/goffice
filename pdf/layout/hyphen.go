package layout

import (
	"strings"
	"unicode/utf8"
)

// SoftHyphen is the Unicode soft hyphen character (U+00AD).
// It is an invisible character that indicates where a word may be hyphenated
// when line wrapping occurs. When a line break occurs at a soft hyphen,
// a visible hyphen should be rendered.
const SoftHyphen = '\u00AD'

// VisibleHyphen is the standard ASCII hyphen-minus character used
// when rendering a soft hyphen break.
const VisibleHyphen = '-'

// SoftHyphenPosition represents a soft hyphen location in text.
type SoftHyphenPosition struct {
	// ByteOffset is the byte position of the soft hyphen in the string.
	ByteOffset int
	// RuneOffset is the rune index of the soft hyphen in the string.
	RuneOffset int
	// BeforeText is the text segment before this soft hyphen (not including the soft hyphen).
	BeforeText string
	// AfterText is the text segment after this soft hyphen (not including the soft hyphen).
	AfterText string
}

// SoftHyphenProcessor handles detection and processing of soft hyphens in text.
// Soft hyphens (U+00AD) are invisible characters that indicate optional
// hyphenation points. When text is broken at a soft hyphen, a visible
// hyphen should be rendered at the end of the line.
type SoftHyphenProcessor struct {
	// positions caches the soft hyphen positions for the current text.
	positions []SoftHyphenPosition
	// originalText is the text being processed.
	originalText string
	// cleanText is the text with all soft hyphens removed.
	cleanText string
}

// NewSoftHyphenProcessor creates a new SoftHyphenProcessor for the given text.
func NewSoftHyphenProcessor(
	text string,
) *SoftHyphenProcessor {
	processor := &SoftHyphenProcessor{
		originalText: text,
	}
	processor.analyze()

	return processor
}

// analyze processes the text to find all soft hyphen positions.
func (shp *SoftHyphenProcessor) analyze() {
	if len(shp.originalText) == 0 {
		shp.cleanText = ""
		shp.positions = nil

		return
	}

	var positions []SoftHyphenPosition
	var cleanBuilder strings.Builder
	cleanBuilder.Grow(len(shp.originalText))

	runeOffset := 0
	lastSoftHyphenEnd := 0

	for i, r := range shp.originalText {
		if r == SoftHyphen {
			// Record the position
			pos := SoftHyphenPosition{
				ByteOffset: i,
				RuneOffset: runeOffset,
				BeforeText: shp.originalText[lastSoftHyphenEnd:i],
			}
			positions = append(positions, pos)

			// Skip writing the soft hyphen to clean text
			lastSoftHyphenEnd = i + utf8.RuneLen(
				r,
			)
		}
		runeOffset++
	}

	// Write any remaining text after the last soft hyphen
	if lastSoftHyphenEnd < len(shp.originalText) {
		cleanBuilder.WriteString(
			shp.originalText[lastSoftHyphenEnd:],
		)
	}

	// Build clean text
	shp.cleanText = shp.buildCleanText()

	// Calculate AfterText for each position
	for i := range positions {
		if i < len(positions)-1 {
			// AfterText goes until the next soft hyphen
			start := positions[i].ByteOffset + utf8.RuneLen(
				SoftHyphen,
			)
			end := positions[i+1].ByteOffset
			positions[i].AfterText = shp.originalText[start:end]
		} else {
			// Last soft hyphen - AfterText is everything after it
			start := positions[i].ByteOffset + utf8.RuneLen(SoftHyphen)
			positions[i].AfterText = shp.originalText[start:]
		}
	}

	shp.positions = positions
}

// buildCleanText constructs the text with all soft hyphens removed.
func (shp *SoftHyphenProcessor) buildCleanText() string {
	if !strings.ContainsRune(
		shp.originalText,
		SoftHyphen,
	) {
		return shp.originalText
	}

	var builder strings.Builder
	builder.Grow(len(shp.originalText))

	for _, r := range shp.originalText {
		if r != SoftHyphen {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

// OriginalText returns the original text including soft hyphens.
func (shp *SoftHyphenProcessor) OriginalText() string {
	return shp.originalText
}

// CleanText returns the text with all soft hyphens removed.
// This is useful for display when no line breaking occurs at soft hyphens.
func (shp *SoftHyphenProcessor) CleanText() string {
	return shp.cleanText
}

// Positions returns all soft hyphen positions in the text.
func (shp *SoftHyphenProcessor) Positions() []SoftHyphenPosition {
	return shp.positions
}

// HasSoftHyphens returns true if the text contains any soft hyphens.
func (shp *SoftHyphenProcessor) HasSoftHyphens() bool {
	return len(shp.positions) > 0
}

// Count returns the number of soft hyphens in the text.
func (shp *SoftHyphenProcessor) Count() int {
	return len(shp.positions)
}

// FindBreakPoints returns all byte positions where the text can be broken
// at a soft hyphen. Each position is where the break would occur
// (after the soft hyphen in the original text).
func (shp *SoftHyphenProcessor) FindBreakPoints() []int {
	breakPoints := make([]int, len(shp.positions))
	for i, pos := range shp.positions {
		breakPoints[i] = pos.ByteOffset + utf8.RuneLen(
			SoftHyphen,
		)
	}

	return breakPoints
}

// BreakResult represents the result of breaking text at a soft hyphen.
type BreakResult struct {
	// BeforePart is the text before the break point, with a visible hyphen appended.
	BeforePart string
	// AfterPart is the text after the break point (not including the soft hyphen).
	AfterPart string
	// BeforePartClean is the text before the break point without soft hyphens,
	// with a visible hyphen appended.
	BeforePartClean string
	// AfterPartClean is the text after the break point without soft hyphens.
	AfterPartClean string
	// BreakPosition is the byte position where the break occurred in the original text.
	BreakPosition int
	// HyphenAdded indicates that a visible hyphen was added to BeforePart.
	HyphenAdded bool
}

// BreakAt breaks the text at the specified soft hyphen index (0-based).
// Returns a BreakResult with the text before and after the break,
// with a visible hyphen added to the before part.
// Returns nil if the index is out of range.
func (shp *SoftHyphenProcessor) BreakAt(
	hyphenIndex int,
) *BreakResult {
	if hyphenIndex < 0 ||
		hyphenIndex >= len(shp.positions) {
		return nil
	}

	pos := shp.positions[hyphenIndex]
	breakPos := pos.ByteOffset + utf8.RuneLen(
		SoftHyphen,
	)

	// Get the before part (including the soft hyphen itself for position accuracy)
	beforeOriginal := shp.originalText[:pos.ByteOffset]
	afterOriginal := shp.originalText[breakPos:]

	// Clean versions (remove soft hyphens)
	beforeClean := RemoveSoftHyphens(
		beforeOriginal,
	)
	afterClean := RemoveSoftHyphens(afterOriginal)

	return &BreakResult{
		BeforePart: beforeOriginal + string(
			VisibleHyphen,
		),
		AfterPart: afterOriginal,
		BeforePartClean: beforeClean + string(
			VisibleHyphen,
		),
		AfterPartClean: afterClean,
		BreakPosition:  breakPos,
		HyphenAdded:    true,
	}
}

// BreakAtPosition breaks the text at the specified byte position.
// This method finds the soft hyphen at or before the given position
// and breaks there. Returns nil if no suitable soft hyphen is found.
func (shp *SoftHyphenProcessor) BreakAtPosition(
	bytePos int,
) *BreakResult {
	// Find the soft hyphen at or before this position
	for i := len(shp.positions) - 1; i >= 0; i-- {
		pos := shp.positions[i]
		softHyphenEnd := pos.ByteOffset + utf8.RuneLen(
			SoftHyphen,
		)
		if softHyphenEnd <= bytePos {
			return shp.BreakAt(i)
		}
	}

	return nil
}

// GetSegments returns all text segments separated by soft hyphens.
// Each segment can be used independently for rendering purposes.
func (shp *SoftHyphenProcessor) GetSegments() []string {
	if len(shp.positions) == 0 {
		return []string{shp.originalText}
	}

	segments := make(
		[]string,
		0,
		len(shp.positions)+1,
	)

	// Add text before first soft hyphen
	if shp.positions[0].ByteOffset > 0 {
		segments = append(
			segments,
			shp.originalText[:shp.positions[0].ByteOffset],
		)
	}

	// Add segments between soft hyphens
	for i, pos := range shp.positions {
		start := pos.ByteOffset + utf8.RuneLen(
			SoftHyphen,
		)
		var end int
		if i < len(shp.positions)-1 {
			end = shp.positions[i+1].ByteOffset
		} else {
			end = len(shp.originalText)
		}
		if start < end {
			segments = append(
				segments,
				shp.originalText[start:end],
			)
		}
	}

	return segments
}

// RemoveSoftHyphens removes all soft hyphen characters from the given text.
func RemoveSoftHyphens(text string) string {
	if !strings.ContainsRune(text, SoftHyphen) {
		return text
	}

	var builder strings.Builder
	builder.Grow(len(text))

	for _, r := range text {
		if r != SoftHyphen {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

// ContainsSoftHyphen returns true if the text contains any soft hyphen characters.
func ContainsSoftHyphen(text string) bool {
	return strings.ContainsRune(text, SoftHyphen)
}

// CountSoftHyphens returns the number of soft hyphens in the text.
func CountSoftHyphens(text string) int {
	count := 0
	for _, r := range text {
		if r == SoftHyphen {
			count++
		}
	}

	return count
}

// InsertSoftHyphen inserts a soft hyphen at the specified rune position.
// Returns the modified string. If the position is out of range,
// the original string is returned unchanged.
func InsertSoftHyphen(
	text string,
	runePos int,
) string {
	runes := []rune(text)
	if runePos < 0 || runePos > len(runes) {
		return text
	}

	result := make([]rune, 0, len(runes)+1)
	result = append(result, runes[:runePos]...)
	result = append(result, SoftHyphen)
	result = append(result, runes[runePos:]...)

	return string(result)
}

// InsertSoftHyphens inserts soft hyphens at multiple rune positions.
// Positions should be in ascending order. Returns the modified string.
func InsertSoftHyphens(
	text string,
	runePositions []int,
) string {
	if len(runePositions) == 0 {
		return text
	}

	runes := []rune(text)
	result := make(
		[]rune,
		0,
		len(runes)+len(runePositions),
	)

	lastPos := 0
	for _, pos := range runePositions {
		if pos < 0 || pos > len(runes) {
			continue
		}
		if pos < lastPos {
			// Positions not in order, skip
			continue
		}
		result = append(
			result,
			runes[lastPos:pos]...)
		result = append(result, SoftHyphen)
		lastPos = pos
	}

	// Append remaining text
	if lastPos < len(runes) {
		result = append(
			result,
			runes[lastPos:]...)
	}

	return string(result)
}

// SoftHyphenBreakInfo contains information about breaking at a soft hyphen
// during line layout.
type SoftHyphenBreakInfo struct {
	// BreakOpportunity is the break opportunity associated with this soft hyphen.
	BreakOpportunity BreakOpportunity
	// OriginalBytePos is the byte position of the soft hyphen in the original text.
	OriginalBytePos int
	// RequiresVisibleHyphen indicates that a visible hyphen should be rendered
	// if the text is broken at this position.
	RequiresVisibleHyphen bool
}

// FindSoftHyphenBreaks integrates with the line breaking system to find
// break opportunities at soft hyphen positions.
func FindSoftHyphenBreaks(
	text string,
) []SoftHyphenBreakInfo {
	processor := NewSoftHyphenProcessor(text)

	if !processor.HasSoftHyphens() {
		return nil
	}

	breaks := make(
		[]SoftHyphenBreakInfo,
		0,
		processor.Count(),
	)

	for i, pos := range processor.Positions() {
		breakBytePos := pos.ByteOffset + utf8.RuneLen(
			SoftHyphen,
		)

		breaks = append(
			breaks,
			SoftHyphenBreakInfo{
				BreakOpportunity: BreakOpportunity{
					Position:     breakBytePos,
					RunePosition: pos.RuneOffset + 1, // Break after the soft hyphen
					Type:         BreakAllowed,
				},
				OriginalBytePos:       pos.ByteOffset,
				RequiresVisibleHyphen: true,
			},
		)

		_ = i // Used in loop
	}

	return breaks
}

// ProcessLineBreakResult adjusts a line break result when the break
// occurs at a soft hyphen position. This ensures the visible hyphen
// is properly added to the line text.
type ProcessedLineResult struct {
	// Text is the line text with soft hyphens removed and visible hyphen added if needed.
	Text string
	// OriginalText is the original line text including soft hyphens.
	OriginalText string
	// EndsWithHyphen indicates that this line ends with a hyphen added due to soft hyphen break.
	EndsWithHyphen bool
	// Width is the rendered width of the text (to be calculated by caller).
	Width float64
}

// ProcessLineForSoftHyphens processes a line's text to handle soft hyphens correctly.
// If breakAtSoftHyphen is true and the line ends at a soft hyphen position,
// a visible hyphen is appended.
func ProcessLineForSoftHyphens(
	lineText string,
	breakAtSoftHyphen bool,
) ProcessedLineResult {
	result := ProcessedLineResult{
		OriginalText: lineText,
	}

	// Check if the line ends with a soft hyphen
	if len(lineText) > 0 {
		runes := []rune(lineText)
		lastRune := runes[len(runes)-1]

		if lastRune == SoftHyphen {
			// Remove the soft hyphen and add visible hyphen
			cleanText := RemoveSoftHyphens(
				lineText,
			)
			result.Text = cleanText + string(
				VisibleHyphen,
			)
			result.EndsWithHyphen = true

			return result
		}
	}

	// Check if we should add hyphen due to break position
	if breakAtSoftHyphen &&
		ContainsSoftHyphen(lineText) {
		// The line contains soft hyphens but doesn't end with one.
		// Just clean the text.
		result.Text = RemoveSoftHyphens(lineText)
		result.EndsWithHyphen = false

		return result
	}

	// No soft hyphen handling needed
	result.Text = RemoveSoftHyphens(lineText)
	result.EndsWithHyphen = false

	return result
}

// SplitIntoLinesWithSoftHyphens extends SplitIntoLines to properly handle
// soft hyphens, adding visible hyphens when breaks occur at soft hyphen positions.
func SplitIntoLinesWithSoftHyphens(
	text string,
	maxWidth float64,
	measureFunc MeasureFunc,
) []LineBreakResult {
	if len(text) == 0 {
		return nil
	}

	// Use WordLineBreaker which already supports soft hyphens
	wlb := NewWordLineBreaker()
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Find soft hyphen positions for post-processing
	processor := NewSoftHyphenProcessor(text)
	softHyphenBreakPositions := make(map[int]bool)

	for _, pos := range processor.Positions() {
		// The break occurs after the soft hyphen
		breakPos := pos.ByteOffset + utf8.RuneLen(
			SoftHyphen,
		)
		softHyphenBreakPositions[breakPos] = true
	}

	var results []LineBreakResult
	startByte := 0

	if len(opportunities) == 0 {
		lineText := text
		processed := ProcessLineForSoftHyphens(
			lineText,
			false,
		)
		results = append(results, LineBreakResult{
			Text:      processed.Text,
			StartByte: 0,
			EndByte:   len(text),
			Width: measureFunc(
				processed.Text,
			),
			ForcedBreak: false,
		})

		return results
	}

	for startByte < len(text) {
		bestBreak := -1
		var bestBreakOpp *BreakOpportunity
		var forcedBreak bool

		for i, opp := range opportunities {
			if opp.Position <= startByte {
				continue
			}

			if opp.Type == BreakMandatory {
				segment := text[startByte:opp.Position]
				processed := ProcessLineForSoftHyphens(
					segment,
					softHyphenBreakPositions[opp.Position],
				)
				width := measureFunc(
					processed.Text,
				)
				if width <= maxWidth ||
					bestBreak == -1 {
					bestBreak = opp.Position
					bestBreakOpp = &opportunities[i]
					forcedBreak = true

					break
				}
			}

			segment := text[startByte:opp.Position]
			processed := ProcessLineForSoftHyphens(
				segment,
				softHyphenBreakPositions[opp.Position],
			)
			width := measureFunc(processed.Text)

			if width <= maxWidth {
				bestBreak = opp.Position
				bestBreakOpp = &opportunities[i]
				forcedBreak = false
			} else if bestBreak != -1 {
				break
			}
		}

		if !forcedBreak {
			remainingText := text[startByte:]
			processed := ProcessLineForSoftHyphens(
				remainingText,
				false,
			)
			remainingWidth := measureFunc(
				trimTrailingSpaces(
					processed.Text,
				),
			)
			if remainingWidth <= maxWidth {
				bestBreak = len(text)
				bestBreakOpp = nil
				forcedBreak = false
			}
		}

		if bestBreak == -1 {
			lineText := text[startByte:]
			processed := ProcessLineForSoftHyphens(
				lineText,
				false,
			)
			results = append(
				results,
				LineBreakResult{
					Text:      processed.Text,
					StartByte: startByte,
					EndByte:   len(text),
					Width: measureFunc(
						processed.Text,
					),
					ForcedBreak: false,
				},
			)

			break
		}

		lineText := text[startByte:bestBreak]
		isSoftHyphenBreak := softHyphenBreakPositions[bestBreak]
		processed := ProcessLineForSoftHyphens(
			lineText,
			isSoftHyphenBreak,
		)

		// If this is a soft hyphen break, add the visible hyphen
		displayText := processed.Text
		if isSoftHyphenBreak &&
			!processed.EndsWithHyphen {
			displayText += string(
				VisibleHyphen,
			)
		}

		results = append(results, LineBreakResult{
			Text:      displayText,
			StartByte: startByte,
			EndByte:   bestBreak,
			Width: measureFunc(
				trimTrailingSpaces(displayText),
			),
			ForcedBreak: forcedBreak,
		})

		startByte = bestBreak

		if !forcedBreak || bestBreakOpp == nil ||
			bestBreakOpp.Type != BreakMandatory {
			for startByte < len(text) && text[startByte] == ' ' {
				startByte++
			}
		}
	}

	return results
}
