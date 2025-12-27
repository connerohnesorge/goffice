package layout

import (
	"strings"
)

// Non-breaking space character constants.
// These characters prevent line breaks from occurring at their position.
const (
	// NBSP is the standard no-break space (U+00A0).
	// This is the most common non-breaking space character, typically
	// used to keep words together that should not be separated.
	NBSP = '\u00A0'

	// NarrowNBSP is the narrow no-break space (U+202F).
	// This is a thinner non-breaking space, often used in French typography
	// before punctuation marks like : ; ! ? and for digit grouping.
	NarrowNBSP = '\u202F'

	// ZWNBSP is the zero width no-break space (U+FEFF).
	// Also known as the Byte Order Mark (BOM) when at the start of a file.
	// As a ZWNBSP, it has zero width and prevents line breaks.
	// Note: Modern Unicode recommends using Word Joiner (U+2060) instead.
	ZWNBSP = '\uFEFF'

	// WordJoiner is the word joiner character (U+2060).
	// This is the preferred zero-width non-breaking character in modern Unicode.
	// It prevents line breaks without adding visible space.
	WordJoiner = '\u2060'

	// FigureSpace is the figure space (U+2007).
	// A non-breaking space with the width of a digit, used for aligning
	// numbers in tables.
	FigureSpace = '\u2007'
)

// NonBreakingSpaceType represents the type of non-breaking space character.
type NonBreakingSpaceType int

const (
	// NBSPTypeStandard represents the standard no-break space (U+00A0).
	NBSPTypeStandard NonBreakingSpaceType = iota
	// NBSPTypeNarrow represents the narrow no-break space (U+202F).
	NBSPTypeNarrow
	// NBSPTypeZeroWidth represents the zero width no-break space (U+FEFF).
	NBSPTypeZeroWidth
	// NBSPTypeWordJoiner represents the word joiner (U+2060).
	NBSPTypeWordJoiner
	// NBSPTypeFigure represents the figure space (U+2007).
	NBSPTypeFigure
	// NBSPTypeUnknown represents an unknown or regular space.
	NBSPTypeUnknown
)

// String returns the string name of the non-breaking space type.
func (t NonBreakingSpaceType) String() string {
	switch t {
	case NBSPTypeStandard:
		return "Standard NBSP"
	case NBSPTypeNarrow:
		return "Narrow NBSP"
	case NBSPTypeZeroWidth:
		return "Zero Width NBSP"
	case NBSPTypeWordJoiner:
		return "Word Joiner"
	case NBSPTypeFigure:
		return "Figure Space"
	case NBSPTypeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

// NonBreakingSpacePosition represents the location of a non-breaking space in text.
type NonBreakingSpacePosition struct {
	// ByteOffset is the byte position of the NBSP in the string.
	ByteOffset int
	// RuneOffset is the rune index of the NBSP in the string.
	RuneOffset int
	// Type indicates the type of non-breaking space.
	Type NonBreakingSpaceType
	// Rune is the actual non-breaking space rune.
	Rune rune
}

// IsNonBreakingSpace returns true if the rune is any type of non-breaking space.
// This includes standard NBSP, narrow NBSP, zero-width NBSP, word joiner,
// and figure space.
func IsNonBreakingSpace(r rune) bool {
	switch r {
	case NBSP,
		NarrowNBSP,
		ZWNBSP,
		WordJoiner,
		FigureSpace:
		return true
	}

	return false
}

// GetNonBreakingSpaceType returns the type of non-breaking space for a rune.
// Returns NBSPTypeUnknown if the rune is not a non-breaking space.
func GetNonBreakingSpaceType(
	r rune,
) NonBreakingSpaceType {
	switch r {
	case NBSP:
		return NBSPTypeStandard
	case NarrowNBSP:
		return NBSPTypeNarrow
	case ZWNBSP:
		return NBSPTypeZeroWidth
	case WordJoiner:
		return NBSPTypeWordJoiner
	case FigureSpace:
		return NBSPTypeFigure
	default:
		return NBSPTypeUnknown
	}
}

// ContainsNonBreakingSpace returns true if the text contains any non-breaking space.
func ContainsNonBreakingSpace(text string) bool {
	for _, r := range text {
		if IsNonBreakingSpace(r) {
			return true
		}
	}

	return false
}

// CountNonBreakingSpaces returns the total count of non-breaking spaces in the text.
func CountNonBreakingSpaces(text string) int {
	count := 0
	for _, r := range text {
		if IsNonBreakingSpace(r) {
			count++
		}
	}

	return count
}

// FindNonBreakingSpaces returns all non-breaking space positions in the text.
func FindNonBreakingSpaces(
	text string,
) []NonBreakingSpacePosition {
	if len(text) == 0 {
		return nil
	}

	var positions []NonBreakingSpacePosition
	byteOffset := 0

	for runeOffset, r := range text {
		if IsNonBreakingSpace(r) {
			positions = append(
				positions,
				NonBreakingSpacePosition{
					ByteOffset: byteOffset,
					RuneOffset: runeOffset,
					Type: GetNonBreakingSpaceType(
						r,
					),
					Rune: r,
				},
			)
		}
		byteOffset += len(string(r))
	}

	// Recalculate rune offsets correctly
	runeIdx := 0
	byteOffset = 0
	posIdx := 0
	for _, r := range text {
		if posIdx < len(positions) &&
			byteOffset == positions[posIdx].ByteOffset {
			positions[posIdx].RuneOffset = runeIdx
			posIdx++
		}
		byteOffset += len(string(r))
		runeIdx++
	}

	return positions
}

// ReplaceSpacesWithNBSP replaces all regular spaces with non-breaking spaces.
// This is useful for ensuring text is kept together without line breaks.
func ReplaceSpacesWithNBSP(text string) string {
	return strings.ReplaceAll(
		text,
		" ",
		string(NBSP),
	)
}

// ReplaceNBSPWithSpaces replaces all standard non-breaking spaces with regular spaces.
// Other types of non-breaking spaces (narrow, zero-width, etc.) are not affected.
func ReplaceNBSPWithSpaces(text string) string {
	return strings.ReplaceAll(
		text,
		string(NBSP),
		" ",
	)
}

// ReplaceAllNBSPWithSpaces replaces all types of non-breaking spaces with regular spaces.
// This normalizes the text by converting all NBSP variants to standard spaces.
func ReplaceAllNBSPWithSpaces(
	text string,
) string {
	var builder strings.Builder
	builder.Grow(len(text))

	for _, r := range text {
		switch r {
		case NBSP, NarrowNBSP, FigureSpace:
			// Replace visible non-breaking spaces with regular space
			builder.WriteRune(' ')
		case ZWNBSP, WordJoiner:
			// Remove zero-width characters entirely
			continue
		default:
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

// InsertNBSP inserts a non-breaking space at the specified rune position.
// Returns the modified string. If the position is out of range,
// the original string is returned unchanged.
func InsertNBSP(text string, runePos int) string {
	return InsertNBSPType(text, runePos, NBSP)
}

// InsertNBSPType inserts a specific type of non-breaking space at the given position.
func InsertNBSPType(
	text string,
	runePos int,
	nbspRune rune,
) string {
	runes := []rune(text)
	if runePos < 0 || runePos > len(runes) {
		return text
	}

	result := make([]rune, 0, len(runes)+1)
	result = append(result, runes[:runePos]...)
	result = append(result, nbspRune)
	result = append(result, runes[runePos:]...)

	return string(result)
}

// InsertWordJoiner inserts a word joiner (zero-width non-breaking space) at the position.
// Word joiners are invisible and prevent line breaks at that position.
func InsertWordJoiner(
	text string,
	runePos int,
) string {
	return InsertNBSPType(
		text,
		runePos,
		WordJoiner,
	)
}

// NBSPProcessor handles detection and processing of non-breaking spaces in text.
type NBSPProcessor struct {
	originalText string
	positions    []NonBreakingSpacePosition
}

// NewNBSPProcessor creates a new NBSPProcessor for the given text.
func NewNBSPProcessor(
	text string,
) *NBSPProcessor {
	processor := &NBSPProcessor{
		originalText: text,
	}
	processor.positions = FindNonBreakingSpaces(
		text,
	)

	return processor
}

// OriginalText returns the original text.
func (p *NBSPProcessor) OriginalText() string {
	return p.originalText
}

// Positions returns all non-breaking space positions in the text.
func (p *NBSPProcessor) Positions() []NonBreakingSpacePosition {
	return p.positions
}

// HasNonBreakingSpaces returns true if the text contains any non-breaking spaces.
func (p *NBSPProcessor) HasNonBreakingSpaces() bool {
	return len(p.positions) > 0
}

// Count returns the number of non-breaking spaces in the text.
func (p *NBSPProcessor) Count() int {
	return len(p.positions)
}

// CountByType returns the count of non-breaking spaces of a specific type.
func (p *NBSPProcessor) CountByType(
	nbspType NonBreakingSpaceType,
) int {
	count := 0
	for _, pos := range p.positions {
		if pos.Type == nbspType {
			count++
		}
	}

	return count
}

// GetPositionsByType returns all positions of a specific NBSP type.
func (p *NBSPProcessor) GetPositionsByType(
	nbspType NonBreakingSpaceType,
) []NonBreakingSpacePosition {
	var result []NonBreakingSpacePosition
	for _, pos := range p.positions {
		if pos.Type == nbspType {
			result = append(result, pos)
		}
	}

	return result
}

// NormalizeToStandardNBSP converts all NBSP types to standard NBSP (U+00A0).
// Zero-width characters (ZWNBSP, Word Joiner) are converted to empty strings.
func (p *NBSPProcessor) NormalizeToStandardNBSP() string {
	if !p.HasNonBreakingSpaces() {
		return p.originalText
	}

	var builder strings.Builder
	builder.Grow(len(p.originalText))

	for _, r := range p.originalText {
		switch r {
		case NarrowNBSP, FigureSpace:
			builder.WriteRune(NBSP)
		case ZWNBSP, WordJoiner:
			// Remove zero-width characters
			continue
		default:
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

// ToDisplayText converts the text for display purposes.
// Non-breaking spaces are preserved, but zero-width characters are removed.
func (p *NBSPProcessor) ToDisplayText() string {
	if !p.HasNonBreakingSpaces() {
		return p.originalText
	}

	var builder strings.Builder
	builder.Grow(len(p.originalText))

	for _, r := range p.originalText {
		switch r {
		case ZWNBSP, WordJoiner:
			// Remove zero-width characters for display
			continue
		default:
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

// IsBreakProhibitedAt checks if a line break is prohibited at the given rune position
// due to a non-breaking space.
func (p *NBSPProcessor) IsBreakProhibitedAt(
	runePos int,
) bool {
	for _, pos := range p.positions {
		// A non-breaking space prohibits breaks both before and after it
		if pos.RuneOffset == runePos ||
			pos.RuneOffset == runePos-1 {
			return true
		}
	}

	return false
}

// JoinWithNBSP joins strings with non-breaking spaces instead of regular spaces.
func JoinWithNBSP(parts ...string) string {
	return strings.Join(parts, string(NBSP))
}

// JoinWithWordJoiner joins strings with word joiners (invisible non-breaking).
func JoinWithWordJoiner(parts ...string) string {
	return strings.Join(parts, string(WordJoiner))
}

// ProtectFromBreaking wraps the text with word joiners at the beginning and end
// to prevent breaks immediately before or after the text.
func ProtectFromBreaking(text string) string {
	if len(text) == 0 {
		return text
	}

	return string(
		WordJoiner,
	) + text + string(
		WordJoiner,
	)
}

// ReplaceSpacesBetweenDigits replaces regular spaces between digits with NBSP.
// This is useful for keeping numbers with spaces (like "1 000 000") together.
func ReplaceSpacesBetweenDigits(
	text string,
) string {
	runes := []rune(text)
	if len(runes) < 3 {
		return text
	}

	var builder strings.Builder
	builder.Grow(len(text))

	for i, r := range runes {
		if r == ' ' && i > 0 && i < len(runes)-1 {
			prevRune := runes[i-1]
			nextRune := runes[i+1]
			// Check if both adjacent characters are digits
			if isDigit(prevRune) &&
				isDigit(nextRune) {
				builder.WriteRune(NBSP)

				continue
			}
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

// isDigit returns true if the rune is a decimal digit.
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// ReplaceSpacesAfterPunctuation replaces spaces after certain punctuation with NBSP.
// This follows French typography rules where spaces before : ; ! ? are non-breaking.
// If beforePunctuation is true, spaces before punctuation are also replaced.
func ReplaceSpacesAfterPunctuation(
	text string,
	beforePunctuation bool,
) string {
	runes := []rune(text)
	if len(runes) < 2 {
		return text
	}

	var builder strings.Builder
	builder.Grow(len(text))

	punctuation := map[rune]bool{
		':': true,
		';': true,
		'!': true,
		'?': true,
	}

	for i, r := range runes {
		if r == ' ' {
			// Check for space before punctuation (French style)
			if beforePunctuation &&
				i < len(runes)-1 {
				nextRune := runes[i+1]
				if punctuation[nextRune] {
					builder.WriteRune(NarrowNBSP)

					continue
				}
			}
			// Check for space after opening punctuation that needs NBSP
			if i > 0 {
				prevRune := runes[i-1]
				if prevRune == '\u00AB' ||
					prevRune == '\u2039' { // Opening guillemets
					builder.WriteRune(NBSP)

					continue
				}
			}
		}
		builder.WriteRune(r)
	}

	return builder.String()
}
