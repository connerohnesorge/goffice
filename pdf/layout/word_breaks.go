package layout

import (
	"strings"
	"unicode"
)

// WordBreakOptions configures Word-specific line breaking behavior.
// These options extend the standard UAX #14 algorithm with rules that
// match Microsoft Word's text layout behavior.
type WordBreakOptions struct {
	// BreakAfterSlash enables breaking after forward slashes in paths and URLs.
	// Default: true
	BreakAfterSlash bool

	// BreakAfterBackslash enables breaking after backslashes in Windows paths.
	// Default: true
	BreakAfterBackslash bool

	// BreakAfterEquals enables breaking after equals signs in assignments.
	// Default: true
	BreakAfterEquals bool

	// BreakAfterColon enables breaking after colons (URLs, drive letters).
	// Default: true
	BreakAfterColon bool

	// KeepEmailsTogether prevents breaks within email addresses.
	// Default: true
	KeepEmailsTogether bool

	// KeepURLsTogether prevents breaks within URLs (except after / in long URLs).
	// Default: true
	KeepURLsTogether bool

	// KeepNumberAbbreviations keeps "No." with following numbers.
	// Default: true
	KeepNumberAbbreviations bool

	// KeepCurrencyWithNumbers keeps currency symbols with their numbers.
	// Default: true
	KeepCurrencyWithNumbers bool

	// KeepUnitsWithNumbers keeps measurement units with numbers (e.g., "10 cm").
	// Default: true
	KeepUnitsWithNumbers bool

	// KeepTitlesWithNames keeps title abbreviations with names (Dr., Mr., etc.).
	// Default: true
	KeepTitlesWithNames bool

	// HyphenateCompoundWords enables recognition of compound word hyphenation points.
	// Default: true
	HyphenateCompoundWords bool

	// HandleDashes enables proper handling of em-dash and en-dash.
	// Default: true
	HandleDashes bool

	// SupportSoftHyphen enables soft hyphen (U+00AD) as optional break point.
	// Default: true
	SupportSoftHyphen bool
}

// DefaultWordBreakOptions returns the default Word-compatible break options.
func DefaultWordBreakOptions() WordBreakOptions {
	return WordBreakOptions{
		BreakAfterSlash:         true,
		BreakAfterBackslash:     true,
		BreakAfterEquals:        true,
		BreakAfterColon:         true,
		KeepEmailsTogether:      true,
		KeepURLsTogether:        true,
		KeepNumberAbbreviations: true,
		KeepCurrencyWithNumbers: true,
		KeepUnitsWithNumbers:    true,
		KeepTitlesWithNames:     true,
		HyphenateCompoundWords:  true,
		HandleDashes:            true,
		SupportSoftHyphen:       true,
	}
}

// WordLineBreaker extends LineBreaker with Word-specific break rules.
type WordLineBreaker struct {
	*LineBreaker
	options WordBreakOptions
}

// NewWordLineBreaker creates a new WordLineBreaker with default options.
func NewWordLineBreaker() *WordLineBreaker {
	return &WordLineBreaker{
		LineBreaker: NewLineBreaker(),
		options:     DefaultWordBreakOptions(),
	}
}

// NewWordLineBreakerWithOptions creates a new WordLineBreaker with custom options.
func NewWordLineBreakerWithOptions(
	options WordBreakOptions,
) *WordLineBreaker {
	return &WordLineBreaker{
		LineBreaker: NewLineBreaker(),
		options:     options,
	}
}

// Options returns the current WordBreakOptions.
func (wlb *WordLineBreaker) Options() WordBreakOptions {
	return wlb.options
}

// SetOptions sets new WordBreakOptions.
func (wlb *WordLineBreaker) SetOptions(
	options WordBreakOptions,
) {
	wlb.options = options
}

// FindBreakOpportunities analyzes text and returns break opportunities
// with Word-specific rules applied.
func (wlb *WordLineBreaker) FindBreakOpportunities(
	text string,
) []BreakOpportunity {
	// First, get the standard UAX #14 break opportunities
	baseOpportunities := wlb.LineBreaker.FindBreakOpportunities(
		text,
	)

	if len(text) == 0 {
		return baseOpportunities
	}

	runes := []rune(text)

	// Calculate byte positions for each rune
	bytePositions := make([]int, len(runes)+1)
	pos := 0
	for i, r := range runes {
		bytePositions[i] = pos
		pos += len(string(r))
	}
	bytePositions[len(runes)] = pos

	// Build a map for quick lookup of existing opportunities
	oppMap := make(map[int]*BreakOpportunity)
	for i := range baseOpportunities {
		oppMap[baseOpportunities[i].Position] = &baseOpportunities[i]
	}

	// Detect special patterns in the text
	emailRanges := wlb.findEmailRanges(text)
	urlRanges := wlb.findURLRanges(text)
	titleRanges := wlb.findTitleRanges(
		text,
		runes,
		bytePositions,
	)
	unitRanges := wlb.findUnitRanges(
		text,
		runes,
		bytePositions,
	)
	numberAbbrevRanges := wlb.findNumberAbbrevRanges(
		text,
		runes,
		bytePositions,
	)

	// Apply Word-specific break modifications
	var result []BreakOpportunity

	for i := 1; i <= len(runes); i++ {
		bytePos := bytePositions[i]

		// Check existing opportunity
		existingOpp, hasExisting := oppMap[bytePos]

		var breakType BreakOpportunityType
		if hasExisting {
			breakType = existingOpp.Type
		} else {
			breakType = BreakProhibited
		}

		// Apply Word-specific rules

		// 1. Check if inside a protected range (emails, URLs, etc.)
		if breakType != BreakMandatory {
			if wlb.isInProtectedRange(
				bytePos,
				emailRanges,
				urlRanges,
				titleRanges,
				unitRanges,
				numberAbbrevRanges,
			) {
				// Check for special cases within URLs where breaks are allowed
				if wlb.options.KeepURLsTogether &&
					wlb.isBreakableURLPosition(
						text,
						bytePos,
						urlRanges,
					) {
					breakType = BreakAllowed
				} else {
					breakType = BreakProhibited
				}
			}
		}

		// 2. Add break opportunities after special characters
		if i > 0 && i < len(runes) {
			prevRune := runes[i-1]
			currRune := runes[i]

			// Break after slash (if not in URL or email)
			if wlb.options.BreakAfterSlash &&
				prevRune == '/' {
				if !wlb.isInRange(
					bytePos,
					emailRanges,
				) &&
					!wlb.isInRange(
						bytePos,
						urlRanges,
					) {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
			}

			// Break after backslash (Windows paths)
			if wlb.options.BreakAfterBackslash &&
				prevRune == '\\' {
				if !wlb.isInRange(
					bytePos,
					emailRanges,
				) {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
			}

			// Break after equals
			if wlb.options.BreakAfterEquals &&
				prevRune == '=' {
				if !wlb.isInRange(
					bytePos,
					urlRanges,
				) {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
			}

			// Break after colon (except in URLs and time notation)
			if wlb.options.BreakAfterColon &&
				prevRune == ':' {
				if !wlb.isInRange(
					bytePos,
					urlRanges,
				) &&
					!isTimeNotation(runes, i-1) {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
			}

			// Handle soft hyphen
			if wlb.options.SupportSoftHyphen &&
				prevRune == '\u00AD' {
				if breakType != BreakMandatory {
					breakType = BreakAllowed
				}
			}

			// Handle dashes (em-dash and en-dash)
			if wlb.options.HandleDashes {
				// Allow break after em-dash (U+2014) and en-dash (U+2013)
				if prevRune == '\u2014' ||
					prevRune == '\u2013' {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
				// Also allow break before em-dash
				if currRune == '\u2014' {
					if breakType != BreakMandatory {
						breakType = BreakAllowed
					}
				}
			}

			// Handle compound words with hyphens
			if wlb.options.HyphenateCompoundWords &&
				prevRune == '-' {
				// Allow break after hyphen in compound words
				if i > 1 && i < len(runes)-1 {
					prevPrevRune := runes[i-2]
					if unicode.IsLetter(
						prevPrevRune,
					) &&
						unicode.IsLetter(
							currRune,
						) {
						if breakType != BreakMandatory {
							breakType = BreakAllowed
						}
					}
				}
			}
		}

		// Only add the opportunity if it's not prohibited
		// (or if it's from the base opportunities)
		if breakType != BreakProhibited ||
			(hasExisting && existingOpp.Type != BreakProhibited) {
			result = append(
				result,
				BreakOpportunity{
					Position:     bytePos,
					RunePosition: i,
					Type:         breakType,
				},
			)
		}
	}

	return result
}

// byteRange represents a range in the text by byte positions.
type byteRange struct {
	start int
	end   int
}

// findEmailRanges finds email address ranges in the text.
func (wlb *WordLineBreaker) findEmailRanges(
	text string,
) []byteRange {
	if !wlb.options.KeepEmailsTogether {
		return nil
	}

	var ranges []byteRange

	// Simple email detection: look for @ and expand to word boundaries
	for i := range len(text) {
		if text[i] == '@' {
			// Find start (go back to space or start)
			start := i
			for start > 0 && !unicode.IsSpace(rune(text[start-1])) && text[start-1] != '<' && text[start-1] != '(' {
				start--
			}

			// Find end (go forward to space or end)
			end := i + 1
			for end < len(text) && !unicode.IsSpace(rune(text[end])) && text[end] != '>' && text[end] != ')' && text[end] != ',' {
				end++
			}

			// Validate it looks like an email (has @ and .)
			segment := text[start:end]
			atIdx := strings.Index(segment, "@")
			if atIdx != -1 &&
				strings.Contains(
					segment[atIdx:],
					".",
				) {
				ranges = append(
					ranges,
					byteRange{
						start: start,
						end:   end,
					},
				)
			}
		}
	}

	return ranges
}

// findURLRanges finds URL ranges in the text.
func (wlb *WordLineBreaker) findURLRanges(
	text string,
) []byteRange {
	if !wlb.options.KeepURLsTogether {
		return nil
	}

	var ranges []byteRange

	// Look for URL schemes
	schemes := []string{
		"http://",
		"https://",
		"ftp://",
		"file://",
		"mailto:",
	}

	for _, scheme := range schemes {
		idx := 0
		for {
			pos := strings.Index(
				text[idx:],
				scheme,
			)
			if pos == -1 {
				break
			}
			start := idx + pos

			// Find end (go to space or end)
			end := start + len(scheme)
			for end < len(text) && !unicode.IsSpace(rune(text[end])) && text[end] != '>' && text[end] != ')' && text[end] != '"' && text[end] != '\'' {
				end++
			}

			ranges = append(
				ranges,
				byteRange{start: start, end: end},
			)
			idx = end
		}
	}

	return ranges
}

// findTitleRanges finds title abbreviation ranges (Dr., Mr., Mrs., etc.).
func (wlb *WordLineBreaker) findTitleRanges(
	text string,
	runes []rune,
	bytePositions []int,
) []byteRange {
	if !wlb.options.KeepTitlesWithNames {
		return nil
	}

	var ranges []byteRange

	titles := []string{
		"Dr.",
		"Mr.",
		"Mrs.",
		"Ms.",
		"Prof.",
		"Rev.",
		"Sr.",
		"Jr.",
		"St.",
		"Hon.",
		"Pres.",
		"Gov.",
		"Sen.",
		"Rep.",
		"Gen.",
		"Col.",
		"Capt.",
		"Lt.",
		"Sgt.",
	}

	lowerText := strings.ToLower(text)

	for _, title := range titles {
		lowerTitle := strings.ToLower(title)
		idx := 0
		for {
			pos := strings.Index(
				lowerText[idx:],
				lowerTitle,
			)
			if pos == -1 {
				break
			}
			start := idx + pos
			end := start + len(title)

			// Check if followed by a space and then a capital letter (name)
			if end < len(text) &&
				text[end] == ' ' {
				// Skip the space
				nameStart := end + 1
				if nameStart < len(text) &&
					unicode.IsUpper(
						rune(text[nameStart]),
					) {
					// Find the end of the name
					nameEnd := nameStart
					for nameEnd < len(text) && !unicode.IsSpace(rune(text[nameEnd])) {
						nameEnd++
					}
					// Include title and the following name
					ranges = append(
						ranges,
						byteRange{
							start: start,
							end:   nameEnd,
						},
					)
				}
			}

			idx = end
		}
	}

	return ranges
}

// findUnitRanges finds number-unit ranges (e.g., "10 cm", "5 kg").
func (wlb *WordLineBreaker) findUnitRanges(
	text string,
	runes []rune,
	bytePositions []int,
) []byteRange {
	if !wlb.options.KeepUnitsWithNumbers {
		return nil
	}

	var ranges []byteRange

	// Common unit abbreviations
	units := []string{
		// Length
		"mm", "cm", "m", "km", "in", "ft", "yd", "mi",
		// Weight
		"mg", "g", "kg", "oz", "lb", "lbs",
		// Volume
		"ml", "l", "L", "gal", "qt", "pt",
		// Area
		"sq ft", "sq m", "sqft", "sqm",
		// Speed
		"mph", "kph", "km/h",
		// Temperature
		"C", "F", "K",
		// Data
		"KB", "MB", "GB", "TB", "PB",
		"kb", "mb", "gb", "tb", "pb",
		"Kb", "Mb", "Gb", "Tb", "Pb",
		// Time
		"ms", "sec", "min", "hr", "hrs",
		// Percentage
		"%",
	}

	// Find patterns like "123 unit" or "123unit"
	for i := 0; i < len(runes); i++ {
		// Look for end of number
		if unicode.IsDigit(runes[i]) {
			numStart := i
			for i < len(runes) && (unicode.IsDigit(runes[i]) || runes[i] == '.' || runes[i] == ',') {
				i++
			}
			numEnd := i

			if numEnd > numStart &&
				i < len(runes) {
				// Check for optional space
				unitStart := i
				if runes[i] == ' ' {
					unitStart = i + 1
				}

				if unitStart < len(runes) {
					// Check if what follows is a unit
					remaining := string(
						runes[unitStart:],
					)
					for _, unit := range units {
						if strings.HasPrefix(
							remaining,
							unit,
						) {
							unitEnd := unitStart + len(
								[]rune(unit),
							)
							// Make sure unit is at word boundary
							if unitEnd >= len(
								runes,
							) ||
								!unicode.IsLetter(
									runes[unitEnd],
								) {
								ranges = append(
									ranges,
									byteRange{
										start: bytePositions[numStart],
										end:   bytePositions[unitEnd],
									},
								)

								break
							}
						}
					}
				}
			}
		}
	}

	return ranges
}

// findNumberAbbrevRanges finds "No." abbreviation ranges.
func (wlb *WordLineBreaker) findNumberAbbrevRanges(
	text string,
	runes []rune,
	bytePositions []int,
) []byteRange {
	if !wlb.options.KeepNumberAbbreviations {
		return nil
	}

	var ranges []byteRange

	abbrevs := []string{
		"No.",
		"no.",
		"Nr.",
		"nr.",
		"#",
	}

	for _, abbrev := range abbrevs {
		idx := 0
		for {
			pos := strings.Index(
				text[idx:],
				abbrev,
			)
			if pos == -1 {
				break
			}
			start := idx + pos
			end := start + len(abbrev)

			// Skip any whitespace after abbreviation
			for end < len(text) && text[end] == ' ' {
				end++
			}

			// Find the number that follows
			if end < len(text) &&
				unicode.IsDigit(rune(text[end])) {
				numEnd := end
				for numEnd < len(text) && (unicode.IsDigit(rune(text[numEnd])) || text[numEnd] == '.' || text[numEnd] == ',' || text[numEnd] == '-') {
					numEnd++
				}
				ranges = append(
					ranges,
					byteRange{
						start: start,
						end:   numEnd,
					},
				)
			}

			idx = end
		}
	}

	return ranges
}

// isInProtectedRange checks if a position is inside any protected range.
func (wlb *WordLineBreaker) isInProtectedRange(
	pos int,
	ranges ...[]byteRange,
) bool {
	for _, rangeSet := range ranges {
		if wlb.isInRange(pos, rangeSet) {
			return true
		}
	}

	return false
}

// isInRange checks if a position is inside any of the given ranges.
func (wlb *WordLineBreaker) isInRange(
	pos int,
	ranges []byteRange,
) bool {
	for _, r := range ranges {
		if pos > r.start && pos < r.end {
			return true
		}
	}

	return false
}

// isBreakableURLPosition checks if a position within a URL can have a break.
// This allows breaking after / in long URLs.
func (wlb *WordLineBreaker) isBreakableURLPosition(
	text string,
	pos int,
	urlRanges []byteRange,
) bool {
	for _, r := range urlRanges {
		if pos > r.start && pos < r.end {
			// Allow break after / (but not ://)
			if pos > 0 && pos < len(text) &&
				text[pos-1] == '/' {
				// Don't break right after :// in scheme
				if pos >= 3 &&
					text[pos-3:pos] == "://" {
					return false
				}
				// Don't break after the first / in path
				schemeEnd := strings.Index(
					text[r.start:],
					"://",
				)
				if schemeEnd != -1 {
					pathStart := r.start + schemeEnd + 3
					// Find the first / in the path (after the host)
					hostEnd := strings.Index(
						text[pathStart:r.end],
						"/",
					)
					if hostEnd != -1 {
						firstSlashInPath := pathStart + hostEnd
						// Allow break after slashes in the path, but not the first one
						if pos > firstSlashInPath+1 {
							return true
						}
					}
				}
			}

			return false
		}
	}

	return false
}

// isTimeNotation checks if a colon at the given position is part of time notation.
func isTimeNotation(
	runes []rune,
	colonPos int,
) bool {
	// Check if there are digits before and after the colon (like 12:30)
	if colonPos <= 0 || colonPos >= len(runes)-1 {
		return false
	}

	hasPrevDigit := unicode.IsDigit(
		runes[colonPos-1],
	)
	hasNextDigit := unicode.IsDigit(
		runes[colonPos+1],
	)

	return hasPrevDigit && hasNextDigit
}

// SplitIntoLinesWord breaks text into lines using Word-specific rules.
func SplitIntoLinesWord(
	text string,
	maxWidth float64,
	measureFunc MeasureFunc,
) []LineBreakResult {
	return SplitIntoLinesWordWithOptions(
		text,
		maxWidth,
		measureFunc,
		DefaultWordBreakOptions(),
	)
}

// SplitIntoLinesWordWithOptions breaks text into lines using Word-specific rules
// with custom options.
func SplitIntoLinesWordWithOptions(
	text string,
	maxWidth float64,
	measureFunc MeasureFunc,
	options WordBreakOptions,
) []LineBreakResult {
	if len(text) == 0 {
		return nil
	}

	wlb := NewWordLineBreakerWithOptions(options)
	opportunities := wlb.FindBreakOpportunities(
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
		if !forcedBreak {
			remainingText := text[startByte:]
			remainingWidth := measureFunc(
				trimTrailingSpaces(remainingText),
			)
			if remainingWidth <= maxWidth {
				bestBreak = len(text)
				bestBreakOpp = nil
				forcedBreak = false
			}
		}

		// If no break found, take the rest of the text
		if bestBreak == -1 {
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

// IsCurrencySymbol returns true if the rune is a currency symbol.
func IsCurrencySymbol(r rune) bool {
	switch r {
	case '$',
		'\u00A2',
		'\u00A3',
		'\u00A5',
		'\u20AC',
		'\u20A3',
		'\u20A4',
		'\u20A7',
		'\u20A9',
		'\u20AA',
		'\u20AB',
		'\u20B9',
		'\u20BA',
		'\u20BC',
		'\u20BD',
		'\u20BF':
		return true
	}

	return unicode.Is(unicode.Sc, r)
}

// IsUnitAbbreviation returns true if the string is a common unit abbreviation.
func IsUnitAbbreviation(s string) bool {
	units := map[string]bool{
		// Length
		"mm": true, "cm": true, "m": true, "km": true,
		"in": true, "ft": true, "yd": true, "mi": true,
		// Weight
		"mg": true, "g": true, "kg": true, "oz": true, "lb": true, "lbs": true,
		// Volume
		"ml": true, "l": true, "L": true, "gal": true, "qt": true, "pt": true,
		// Speed
		"mph": true, "kph": true,
		// Temperature
		"C": true, "F": true, "K": true,
		// Data
		"KB": true, "MB": true, "GB": true, "TB": true, "PB": true,
		"kb": true, "mb": true, "gb": true, "tb": true, "pb": true,
		// Time
		"ms": true, "sec": true, "min": true, "hr": true, "hrs": true,
		// Other
		"%": true,
	}

	return units[s]
}

// IsTitleAbbreviation returns true if the string is a title abbreviation.
func IsTitleAbbreviation(s string) bool {
	titles := map[string]bool{
		"Dr.": true, "Mr.": true, "Mrs.": true, "Ms.": true,
		"Prof.": true, "Rev.": true, "Sr.": true, "Jr.": true,
		"St.": true, "Hon.": true, "Pres.": true, "Gov.": true,
		"Sen.": true, "Rep.": true, "Gen.": true, "Col.": true,
		"Capt.": true, "Lt.": true, "Sgt.": true,
	}

	return titles[s]
}
