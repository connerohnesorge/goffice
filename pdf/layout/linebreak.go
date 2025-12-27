package layout

import (
	"unicode"
)

// LineBreakClass represents a Unicode UAX #14 line breaking class.
// Each character is assigned to one of these classes to determine
// where line breaks can occur.
type LineBreakClass int

const (
	// Mandatory break classes
	BK LineBreakClass = iota // Mandatory Break
	CR                       // Carriage Return
	LF                       // Line Feed
	NL                       // Next Line

	// Break opportunity classes
	SP // Space
	ZW // Zero Width Space

	// Non-break classes
	GL // Non-breaking (Glue)
	WJ // Word Joiner

	// Opening and closing
	OP // Opening Punctuation
	CL // Closing Punctuation
	CP // Close Parenthesis
	QU // Quotation

	// Exclamation/Interrogation
	EX // Exclamation/Interrogation

	// Alphabetic and related
	AL // Ordinary Alphabetic and Symbol Characters
	NU // Numeric
	PR // Prefix Numeric
	PO // Postfix Numeric
	HY // Hyphen
	BA // Break After
	BB // Break Before
	B2 // Break Opportunity Before and After
	NS // Nonstarter

	// Ideographic
	ID // Ideographic
	IN // Inseparable

	// Special
	CM // Combining Mark
	SY // Symbols Allowing Break After
	IS // Infix Numeric Separator
	SA // Complex Context Dependent (South East Asian)
	CB // Contingent Break Opportunity

	// Unknown/unassigned
	XX // Unknown
)

// lineBreakClassNames maps LineBreakClass to string names.
var lineBreakClassNames = [...]string{
	BK: "BK",
	CR: "CR",
	LF: "LF",
	NL: "NL",
	SP: "SP",
	ZW: "ZW",
	GL: "GL",
	WJ: "WJ",
	OP: "OP",
	CL: "CL",
	CP: "CP",
	QU: "QU",
	EX: "EX",
	AL: "AL",
	NU: "NU",
	PR: "PR",
	PO: "PO",
	HY: "HY",
	BA: "BA",
	BB: "BB",
	B2: "B2",
	NS: "NS",
	ID: "ID",
	IN: "IN",
	CM: "CM",
	SY: "SY",
	IS: "IS",
	SA: "SA",
	CB: "CB",
	XX: "XX",
}

// String returns the string name of the line break class.
func (c LineBreakClass) String() string {
	if int(c) < len(lineBreakClassNames) {
		return lineBreakClassNames[c]
	}

	return "XX"
}

// BreakOpportunityType represents the type of break opportunity at a position.
type BreakOpportunityType int

const (
	// BreakProhibited means no break is allowed at this position.
	BreakProhibited BreakOpportunityType = iota
	// BreakAllowed means a break is allowed at this position.
	BreakAllowed
	// BreakMandatory means a break is required at this position.
	BreakMandatory
)

// String returns the string representation of the break opportunity type.
func (t BreakOpportunityType) String() string {
	switch t {
	case BreakProhibited:
		return "Prohibited"
	case BreakAllowed:
		return "Allowed"
	case BreakMandatory:
		return "Mandatory"
	default:
		return "Unknown"
	}
}

// BreakOpportunity represents a potential line break position in text.
type BreakOpportunity struct {
	// Position is the byte offset in the string where the break can occur.
	// The break occurs before the character at this position.
	Position int
	// RunePosition is the rune index where the break can occur.
	RunePosition int
	// Type indicates whether this break is mandatory, allowed, or prohibited.
	Type BreakOpportunityType
}

// LineBreaker implements Unicode UAX #14 line breaking algorithm.
type LineBreaker struct {
	// strictMode enables strict conformance to UAX #14.
	// When false (default), some relaxations are allowed for
	// better compatibility with common text.
	strictMode bool
}

// NewLineBreaker creates a new LineBreaker with default settings.
func NewLineBreaker() *LineBreaker {
	return &LineBreaker{
		strictMode: false,
	}
}

// NewStrictLineBreaker creates a new LineBreaker with strict UAX #14 conformance.
func NewStrictLineBreaker() *LineBreaker {
	return &LineBreaker{
		strictMode: true,
	}
}

// GetLineBreakClass returns the UAX #14 line break class for a rune.
func GetLineBreakClass(r rune) LineBreakClass {
	// Mandatory break characters
	switch r {
	case '\n':
		return LF
	case '\r':
		return CR
	case '\u000B', '\u000C': // VT, FF
		return BK
	case '\u0085': // NEL (Next Line)
		return NL
	case '\u2028': // Line Separator
		return BK
	case '\u2029': // Paragraph Separator
		return BK
	}

	// Space characters
	switch r {
	case ' ':
		return SP
	case '\u00A0': // No-Break Space
		return GL
	case '\u1680',
		'\u2000',
		'\u2001',
		'\u2002',
		'\u2003',
		'\u2004',
		'\u2005',
		'\u2006',
		'\u2008',
		'\u2009',
		'\u200A',
		'\u205F',
		'\u3000':
		return SP // Various spaces
	case '\u2007': // Figure Space (non-breaking)
		return GL
	case '\u202F': // Narrow No-Break Space
		return GL
	}

	// Zero-width characters
	switch r {
	case '\u200B': // Zero Width Space
		return ZW
	case '\u2060',
		'\uFEFF': // Word Joiner, BOM (as WJ)
		return WJ
	case '\u00AD': // Soft Hyphen
		return BA
	}

	// Quotation marks
	switch r {
	case '"',
		'\'',
		'\u2018',
		'\u2019',
		'\u201A',
		'\u201B',
		'\u201C',
		'\u201D',
		'\u201E',
		'\u201F',
		'\u2039',
		'\u203A':
		return QU
	}

	// Opening punctuation
	switch r {
	case '(',
		'[',
		'{',
		'\u00AB',
		'\u2018',
		'\u201C',
		'\u2039':
		return OP
	}
	if unicode.Is(
		unicode.Ps,
		r,
	) { // Opening punctuation category
		return OP
	}

	// Closing punctuation
	switch r {
	case ')',
		']',
		'}',
		'\u00BB',
		'\u2019',
		'\u201D',
		'\u203A':
		return CL
	}
	if unicode.Is(
		unicode.Pe,
		r,
	) { // Closing punctuation category
		return CL
	}

	// Exclamation and question marks
	switch r {
	case '!',
		'?',
		'\u203C',
		'\u2047',
		'\u2048',
		'\u2049':
		return EX
	}

	// Hyphens
	switch r {
	case '-',
		'\u2010',
		'\u2011',
		'\u2012',
		'\u2013':
		return HY
	case '\u2014',
		'\u2015': // Em dash, horizontal bar
		return BA // Allow break after
	}

	// Numeric separators
	switch r {
	case '.':
		return IS
	case ',':
		return IS
	case ':':
		return IS
	case ';':
		return IS
	}

	// Prefix/postfix for numbers
	switch r {
	case '$',
		'\u00A3',
		'\u00A5',
		'\u20AC',
		'\u00A2': // Currency symbols
		return PR
	case '%',
		'\u2030',
		'\u2031': // Percent, per mille, per ten thousand
		return PO
	}

	// CJK characters (Ideographic)
	if isCJKIdeograph(r) {
		return ID
	}

	// Hangul syllables
	if unicode.Is(unicode.Hangul, r) {
		return ID
	}

	// Hiragana and Katakana
	if unicode.Is(unicode.Hiragana, r) ||
		unicode.Is(unicode.Katakana, r) {
		return ID
	}

	// Non-starters (CJK punctuation that cannot start a line)
	if isCJKNonStarter(r) {
		return NS
	}

	// Combining marks
	if unicode.Is(unicode.Mn, r) ||
		unicode.Is(unicode.Mc, r) ||
		unicode.Is(unicode.Me, r) {
		return CM
	}

	// Numeric
	if unicode.IsDigit(r) {
		return NU
	}

	// Default: treat as alphabetic
	if unicode.IsLetter(r) ||
		unicode.IsSymbol(r) {
		return AL
	}

	return XX
}

// isCJKIdeograph returns true if the rune is a CJK ideographic character.
func isCJKIdeograph(r rune) bool {
	// CJK Unified Ideographs
	if r >= 0x4E00 && r <= 0x9FFF {
		return true
	}
	// CJK Unified Ideographs Extension A
	if r >= 0x3400 && r <= 0x4DBF {
		return true
	}
	// CJK Unified Ideographs Extension B
	if r >= 0x20000 && r <= 0x2A6DF {
		return true
	}
	// CJK Unified Ideographs Extension C
	if r >= 0x2A700 && r <= 0x2B73F {
		return true
	}
	// CJK Unified Ideographs Extension D
	if r >= 0x2B740 && r <= 0x2B81F {
		return true
	}
	// CJK Compatibility Ideographs
	if r >= 0xF900 && r <= 0xFAFF {
		return true
	}

	return false
}

// isCJKNonStarter returns true if the rune is a CJK character that cannot start a line.
func isCJKNonStarter(r rune) bool {
	switch r {
	// Japanese small kana and prolonged sound mark
	case '\u3041',
		'\u3043',
		'\u3045',
		'\u3047',
		'\u3049', // Small hiragana
		'\u3063',
		'\u3083',
		'\u3085',
		'\u3087',
		'\u308E', // Small hiragana
		'\u30A1',
		'\u30A3',
		'\u30A5',
		'\u30A7',
		'\u30A9', // Small katakana
		'\u30C3',
		'\u30E3',
		'\u30E5',
		'\u30E7',
		'\u30EE', // Small katakana
		'\u30FC', // Prolonged sound mark
		'\u3001',
		'\u3002', // Ideographic comma and period
		'\u3005', // Ideographic iteration mark
		'\uFF01',
		'\uFF09',
		'\uFF0C',
		'\uFF0E',
		'\uFF1A',
		'\uFF1B',
		'\uFF1F', // Fullwidth punctuation
		'\uFF3D',
		'\uFF5D':
		return true
	}

	return false
}

// FindBreakOpportunities analyzes a string and returns all break opportunities.
// The returned slice contains positions where breaks can or must occur.
func (lb *LineBreaker) FindBreakOpportunities(
	text string,
) []BreakOpportunity {
	if len(text) == 0 {
		return nil
	}

	runes := []rune(text)
	if len(runes) == 0 {
		return nil
	}

	// Calculate byte positions for each rune
	bytePositions := make([]int, len(runes)+1)
	pos := 0
	for i, r := range runes {
		bytePositions[i] = pos
		pos += len(string(r))
	}
	bytePositions[len(runes)] = pos

	// Get line break classes for all runes
	classes := make([]LineBreakClass, len(runes))
	for i, r := range runes {
		classes[i] = GetLineBreakClass(r)
	}

	var opportunities []BreakOpportunity

	// Process each position between characters (and after the last)
	for i := 1; i <= len(runes); i++ {
		var breakType BreakOpportunityType
		var isMandatory bool

		prevClass := classes[i-1]

		// Rule LB4 & LB5: Mandatory breaks after BK, CR, LF, NL
		switch prevClass {
		case BK:
			breakType = BreakMandatory
			isMandatory = true
		case CR:
			// Check for CR+LF - if the current char is LF, this is part of CRLF
			// and we don't break here
			if i < len(runes) &&
				classes[i] == LF {
				continue // Skip - we'll break after the LF
			}
			breakType = BreakMandatory
			isMandatory = true
		case LF, NL:
			breakType = BreakMandatory
			isMandatory = true
		case SP,
			ZW,
			GL,
			WJ,
			OP,
			CL,
			CP,
			QU,
			EX,
			AL,
			NU,
			PR,
			PO,
			HY,
			BA,
			BB,
			B2,
			NS,
			ID,
			IN,
			CM,
			SY,
			IS,
			SA,
			CB,
			XX:
			// No mandatory break
		default:
			// No mandatory break
		}

		if isMandatory {
			opportunities = append(
				opportunities,
				BreakOpportunity{
					Position:     bytePositions[i],
					RunePosition: i,
					Type:         breakType,
				},
			)

			continue
		}

		// If we're at the end of the string, no more breaks to check
		if i >= len(runes) {
			continue
		}

		currClass := classes[i]

		// Determine break opportunity between prevClass and currClass
		breakType = lb.getBreakBetween(
			prevClass,
			currClass,
			runes,
			classes,
			i,
		)

		if breakType != BreakProhibited {
			opportunities = append(
				opportunities,
				BreakOpportunity{
					Position:     bytePositions[i],
					RunePosition: i,
					Type:         breakType,
				},
			)
		}
	}

	return opportunities
}

// getBreakBetween determines the break opportunity between two adjacent characters.
func (lb *LineBreaker) getBreakBetween(
	prevClass, currClass LineBreakClass,
	runes []rune,
	classes []LineBreakClass,
	pos int,
) BreakOpportunityType {
	// Rule LB6: Do not break before hard line breaks
	if currClass == BK || currClass == CR ||
		currClass == LF ||
		currClass == NL {
		return BreakProhibited
	}

	// Rule LB7: Do not break before spaces or zero width space
	if currClass == SP || currClass == ZW {
		return BreakProhibited
	}

	// Rule LB8: Break before any character following a zero-width space
	if prevClass == ZW {
		return BreakAllowed
	}

	// Rule LB18: Break after spaces
	// If the previous character is a space, we generally allow a break here.
	// The space acts as a break opportunity unless we're breaking before
	// something that specifically prohibits it (like NS or closing punctuation).
	if prevClass == SP {
		// Check if the current class prohibits breaks before it
		// This handles things like non-starters that can't start a line
		if currClass == NS {
			return BreakProhibited
		}
		// After space, we allow break to most character classes
		return BreakAllowed
	}

	// Apply the pair table
	return lb.lookupPairTable(
		prevClass,
		currClass,
	)
}

// lookupPairTable returns the break opportunity between two line break classes.
// This implements a simplified version of the UAX #14 pair table.
func (lb *LineBreaker) lookupPairTable(
	before, after LineBreakClass,
) BreakOpportunityType {
	// Rule LB9: Do not break before combining marks
	if after == CM {
		return BreakProhibited
	}

	// Rule LB10: Treat CM as AL if not following a base
	// (handled by treating CM as AL in the pair table)

	// Rule LB11: Do not break before or after Word Joiner
	if before == WJ || after == WJ {
		return BreakProhibited
	}

	// Rule LB12: Do not break after NBSP and related
	if before == GL {
		return BreakProhibited
	}

	// Rule LB12a: Do not break before NBSP and related
	if after == GL {
		return BreakProhibited
	}

	// Rule LB13: Do not break before closing punctuation
	if after == CL || after == CP ||
		after == EX ||
		after == IS ||
		after == SY {
		return BreakProhibited
	}

	// Rule LB14: Do not break after opening punctuation
	if before == OP {
		return BreakProhibited
	}

	// Rule LB15: Do not break within quotes
	if before == QU && after == OP {
		return BreakProhibited
	}

	// Rule LB16: Do not break between closing punctuation and non-starters
	if (before == CL || before == CP) &&
		after == NS {
		return BreakProhibited
	}

	// Rule LB17: Do not break within 'em dash em dash'
	if before == B2 && after == B2 {
		return BreakProhibited
	}

	// Rule LB18: Break after spaces
	if before == SP {
		return BreakAllowed
	}

	// Rule LB19: Do not break before or after quotation marks
	if before == QU || after == QU {
		return BreakProhibited
	}

	// Rule LB20: Break before and after CB (contingent break)
	if before == CB || after == CB {
		return BreakAllowed
	}

	// Rule LB21: Do not break before hyphen-minus, but allow after
	if after == BA || after == HY || after == NS {
		return BreakProhibited
	}

	// Rule LB21a: Don't break after Hebrew + Hyphen
	// (Simplified: we don't track Hebrew specifically)

	// Rule LB21b: Don't break between SY and HL
	// (Simplified: we don't track Hebrew specifically)

	// Rule LB22: Do not break between two ellipses
	if before == IN && after == IN {
		return BreakProhibited
	}

	// Rule LB23: Do not break between digits and letters
	if (before == AL || before == NU) &&
		after == NU {
		return BreakProhibited
	}
	if before == NU &&
		(after == AL || after == NU) {
		return BreakProhibited
	}

	// Rule LB23a: Do not break between numeric prefix and ideographs
	if before == PR && after == ID {
		return BreakProhibited
	}
	if before == ID && after == PO {
		return BreakProhibited
	}

	// Rule LB24: Do not break between prefix and letters
	if (before == PR || before == PO) &&
		(after == AL || after == NU) {
		return BreakProhibited
	}
	if (before == AL || before == NU) &&
		(after == PR || after == PO) {
		return BreakProhibited
	}

	// Rule LB25: Do not break between numeric values
	// (Complex rule for number+separator combinations)
	if before == NU && after == IS {
		return BreakProhibited
	}
	if before == IS && after == NU {
		return BreakProhibited
	}
	if before == NU && after == SY {
		return BreakProhibited
	}

	// Rule LB26: Korean syllable block rules
	// (Simplified: treat Hangul as ID)

	// Rule LB27: Korean/CJK with JL/JV/JT
	// (Simplified for common cases)

	// Rule LB28: Do not break between alphabetics
	if before == AL && after == AL {
		return BreakProhibited
	}

	// Rule LB29: Do not break between numeric punctuation and alphabetics
	if before == IS && after == AL {
		return BreakProhibited
	}

	// Rule LB30: Do not break between letters/numbers and opening/closing
	if (before == AL || before == NU) &&
		after == OP {
		return BreakProhibited
	}
	if before == CP &&
		(after == AL || after == NU) {
		return BreakProhibited
	}

	// Rule LB30a: East Asian Width specific rules
	// (Simplified)

	// Rule LB30b: Emoji handling
	// (Simplified)

	// Rule LB31: Break everywhere else
	// For ideographic characters, allow breaks
	if before == ID || after == ID {
		return BreakAllowed
	}

	// Default: allow break between other classes
	return BreakAllowed
}

// CanBreakBefore returns true if a line break is allowed before the
// character at the given rune position.
func (lb *LineBreaker) CanBreakBefore(
	text string,
	runePos int,
) bool {
	opportunities := lb.FindBreakOpportunities(
		text,
	)
	for _, opp := range opportunities {
		if opp.RunePosition == runePos {
			return opp.Type == BreakAllowed ||
				opp.Type == BreakMandatory
		}
	}

	return false
}

// CanBreakAfter returns true if a line break is allowed after the
// character at the given rune position.
func (lb *LineBreaker) CanBreakAfter(
	text string,
	runePos int,
) bool {
	return lb.CanBreakBefore(text, runePos+1)
}

// FindFirstBreak finds the first break opportunity in the text.
// Returns the byte position of the first break, or -1 if no break exists.
func FindFirstBreak(text string) int {
	lb := NewLineBreaker()
	opportunities := lb.FindBreakOpportunities(
		text,
	)
	if len(opportunities) > 0 {
		return opportunities[0].Position
	}

	return -1
}

// MeasureFunc is a function type that measures the width of a text string.
// It should return the width in the same units used for maxWidth.
type MeasureFunc func(text string) float64

// LineBreakResult represents the result of breaking text into lines.
type LineBreakResult struct {
	// Text is the text content of this line.
	Text string
	// StartByte is the byte offset in the original string where this line starts.
	StartByte int
	// EndByte is the byte offset in the original string where this line ends.
	EndByte int
	// Width is the measured width of this line (if a measure function was provided).
	Width float64
	// ForcedBreak is true if this line ended with a mandatory break.
	ForcedBreak bool
}

// SplitIntoLines breaks text into lines that fit within the specified width.
// The measureFunc is called to determine the width of text segments.
// Returns a slice of line results.
func SplitIntoLines(
	text string,
	maxWidth float64,
	measureFunc MeasureFunc,
) []LineBreakResult {
	if len(text) == 0 {
		return nil
	}

	lb := NewLineBreaker()
	opportunities := lb.FindBreakOpportunities(
		text,
	)

	var results []LineBreakResult
	startByte := 0

	// If no opportunities, return the whole text as one line
	if len(opportunities) == 0 {
		results = append(results, LineBreakResult{
			Text:        text,
			StartByte:   0,
			EndByte:     len(text),
			Width:       measureFunc(text),
			ForcedBreak: false,
		})

		return results
	}

	// Build lines by finding break points that fit
	for startByte < len(text) {
		bestBreak := -1
		var bestBreakOpp *BreakOpportunity
		var forcedBreak bool

		// Find the best break point that fits
		for i, opp := range opportunities {
			if opp.Position <= startByte {
				continue
			}

			// Check if we have a mandatory break
			if opp.Type == BreakMandatory {
				segment := text[startByte:opp.Position]
				width := measureFunc(segment)
				if width <= maxWidth ||
					bestBreak == -1 {
					bestBreak = opp.Position
					bestBreakOpp = &opportunities[i]
					forcedBreak = true

					break // Must break here
				}
			}

			// Check if this segment fits
			segment := text[startByte:opp.Position]
			width := measureFunc(segment)

			if width <= maxWidth {
				bestBreak = opp.Position
				bestBreakOpp = &opportunities[i]
				forcedBreak = false
			} else if bestBreak != -1 {
				// We've exceeded the width and have a valid break
				break
			}
		}

		// Also check if the remaining text fits entirely
		// (this handles the case where there are no more break opportunities
		// but the remaining text still fits)
		if !forcedBreak {
			remainingText := text[startByte:]
			remainingWidth := measureFunc(
				trimTrailingSpaces(remainingText),
			)
			if remainingWidth <= maxWidth {
				// The rest of the text fits, so take it all
				bestBreak = len(text)
				bestBreakOpp = nil
				forcedBreak = false
			}
		}

		// If no break found, we need to break at the end or force a break
		if bestBreak == -1 {
			// No more valid breaks, take the rest of the text
			results = append(
				results,
				LineBreakResult{
					Text:      text[startByte:],
					StartByte: startByte,
					EndByte:   len(text),
					Width: measureFunc(
						text[startByte:],
					),
					ForcedBreak: false,
				},
			)

			break
		}

		// Add this line
		lineText := text[startByte:bestBreak]

		// Trim trailing spaces for width measurement (but keep in text for proper rendering)
		trimmedText := trimTrailingSpaces(
			lineText,
		)

		results = append(results, LineBreakResult{
			Text:        lineText,
			StartByte:   startByte,
			EndByte:     bestBreak,
			Width:       measureFunc(trimmedText),
			ForcedBreak: forcedBreak,
		})

		startByte = bestBreak

		// Skip leading spaces on new line (except after mandatory break)
		if !forcedBreak || bestBreakOpp == nil ||
			bestBreakOpp.Type != BreakMandatory {
			for startByte < len(text) && text[startByte] == ' ' {
				startByte++
			}
		}
	}

	return results
}

// trimTrailingSpaces removes trailing spaces from a string.
func trimTrailingSpaces(s string) string {
	end := len(s)
	for end > 0 && s[end-1] == ' ' {
		end--
	}

	return s[:end]
}

// BreakText is a convenience function that breaks text at the first opportunity.
// It returns two strings: the text before the break and the text after.
// If no break is found, it returns the original text and an empty string.
func BreakText(
	text string,
) (before, after string) {
	pos := FindFirstBreak(text)
	if pos == -1 || pos >= len(text) {
		return text, ""
	}

	return text[:pos], text[pos:]
}
