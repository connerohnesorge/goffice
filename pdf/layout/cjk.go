package layout

// CJK character breaking support for East Asian text layout.
// This file implements kinsoku shori (line breaking prohibition rules)
// for Chinese, Japanese, and Korean text.

// CJKScriptType represents the type of CJK script.
type CJKScriptType int

const (
	// CJKScriptNone indicates the character is not CJK.
	CJKScriptNone CJKScriptType = iota
	// CJKScriptHan represents Han/Chinese ideographs.
	CJKScriptHan
	// CJKScriptHiragana represents Japanese Hiragana.
	CJKScriptHiragana
	// CJKScriptKatakana represents Japanese Katakana.
	CJKScriptKatakana
	// CJKScriptHangul represents Korean Hangul.
	CJKScriptHangul
	// CJKScriptBopomofo represents Bopomofo (Zhuyin).
	CJKScriptBopomofo
	// CJKScriptSymbol represents CJK symbols and punctuation.
	CJKScriptSymbol
)

// String returns the string name of the CJK script type.
func (s CJKScriptType) String() string {
	switch s {
	case CJKScriptNone:
		return "None"
	case CJKScriptHan:
		return "Han"
	case CJKScriptHiragana:
		return "Hiragana"
	case CJKScriptKatakana:
		return "Katakana"
	case CJKScriptHangul:
		return "Hangul"
	case CJKScriptBopomofo:
		return "Bopomofo"
	case CJKScriptSymbol:
		return "Symbol"
	default:
		return "Unknown"
	}
}

// CJKBreakRule represents the line breaking rule for a CJK character.
type CJKBreakRule int

const (
	// CJKBreakNormal indicates normal breaking rules apply.
	CJKBreakNormal CJKBreakRule = iota
	// CJKBreakNoBreakBefore indicates no break is allowed before this character (kinsoku shori).
	CJKBreakNoBreakBefore
	// CJKBreakNoBreakAfter indicates no break is allowed after this character.
	CJKBreakNoBreakAfter
	// CJKBreakNoBreakEither indicates no break is allowed before or after this character.
	CJKBreakNoBreakEither
)

// String returns the string name of the CJK break rule.
func (r CJKBreakRule) String() string {
	switch r {
	case CJKBreakNormal:
		return "Normal"
	case CJKBreakNoBreakBefore:
		return "NoBreakBefore"
	case CJKBreakNoBreakAfter:
		return "NoBreakAfter"
	case CJKBreakNoBreakEither:
		return "NoBreakEither"
	default:
		return "Unknown"
	}
}

// IsCJK returns true if the rune is any CJK character.
// This includes Han ideographs, Hiragana, Katakana, Hangul, Bopomofo,
// and CJK symbols/punctuation.
func IsCJK(r rune) bool {
	return GetCJKScriptType(r) != CJKScriptNone
}

// GetCJKScriptType returns the CJK script type for a rune.
func GetCJKScriptType(r rune) CJKScriptType {
	// CJK Unified Ideographs
	if isHanCharacter(r) {
		return CJKScriptHan
	}

	// Hiragana
	if isHiragana(r) {
		return CJKScriptHiragana
	}

	// Katakana
	if isKatakana(r) {
		return CJKScriptKatakana
	}

	// Hangul
	if isHangul(r) {
		return CJKScriptHangul
	}

	// Bopomofo
	if isBopomofo(r) {
		return CJKScriptBopomofo
	}

	// CJK Symbols and Punctuation
	if isCJKSymbol(r) {
		return CJKScriptSymbol
	}

	return CJKScriptNone
}

// isHanCharacter returns true if the rune is a Han/Chinese ideograph.
func isHanCharacter(r rune) bool {
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
	// CJK Unified Ideographs Extension E
	if r >= 0x2B820 && r <= 0x2CEAF {
		return true
	}
	// CJK Unified Ideographs Extension F
	if r >= 0x2CEB0 && r <= 0x2EBEF {
		return true
	}
	// CJK Unified Ideographs Extension G
	if r >= 0x30000 && r <= 0x3134F {
		return true
	}
	// CJK Compatibility Ideographs
	if r >= 0xF900 && r <= 0xFAFF {
		return true
	}
	// CJK Compatibility Ideographs Supplement
	if r >= 0x2F800 && r <= 0x2FA1F {
		return true
	}

	return false
}

// isHiragana returns true if the rune is a Hiragana character.
func isHiragana(r rune) bool {
	// Hiragana block
	if r >= 0x3040 && r <= 0x309F {
		return true
	}

	// Small Kana Extension
	if r >= 0x1B130 && r <= 0x1B16F {
		return true
	}

	return false
}

// isKatakana returns true if the rune is a Katakana character.
func isKatakana(r rune) bool {
	// Katakana block
	if r >= 0x30A0 && r <= 0x30FF {
		return true
	}

	// Katakana Phonetic Extensions
	if r >= 0x31F0 && r <= 0x31FF {
		return true
	}

	// Halfwidth Katakana
	if r >= 0xFF65 && r <= 0xFF9F {
		return true
	}

	// Small Kana Extension
	if r >= 0x1B130 && r <= 0x1B16F {
		return true
	}

	return false
}

// isHangul returns true if the rune is a Hangul character.
func isHangul(r rune) bool {
	// Hangul Jamo
	if r >= 0x1100 && r <= 0x11FF {
		return true
	}

	// Hangul Compatibility Jamo
	if r >= 0x3130 && r <= 0x318F {
		return true
	}

	// Hangul Syllables
	if r >= 0xAC00 && r <= 0xD7AF {
		return true
	}

	// Hangul Jamo Extended-A
	if r >= 0xA960 && r <= 0xA97F {
		return true
	}

	// Hangul Jamo Extended-B
	if r >= 0xD7B0 && r <= 0xD7FF {
		return true
	}

	return false
}

// isBopomofo returns true if the rune is a Bopomofo character.
func isBopomofo(r rune) bool {
	// Bopomofo
	if r >= 0x3100 && r <= 0x312F {
		return true
	}
	// Bopomofo Extended
	if r >= 0x31A0 && r <= 0x31BF {
		return true
	}

	return false
}

// isCJKSymbol returns true if the rune is a CJK symbol or punctuation.
func isCJKSymbol(r rune) bool {
	// CJK Symbols and Punctuation
	if r >= 0x3000 && r <= 0x303F {
		return true
	}

	// CJK Compatibility Forms
	if r >= 0xFE30 && r <= 0xFE4F {
		return true
	}

	// Enclosed CJK Letters and Months
	if r >= 0x3200 && r <= 0x32FF {
		return true
	}

	// CJK Compatibility
	if r >= 0x3300 && r <= 0x33FF {
		return true
	}

	// Fullwidth Forms
	if r >= 0xFF00 && r <= 0xFF60 {
		return true
	}

	// Halfwidth Forms (CJK portion)
	if r >= 0xFFA0 && r <= 0xFFDC {
		return true
	}

	return false
}

// noBreakBeforeChars contains characters that should not start a line (kinsoku shori).
// These include closing brackets, periods, commas, small kana, etc.
var noBreakBeforeChars = map[rune]bool{
	// Closing brackets and parentheses
	'\u3001': true, // Ideographic comma
	'\u3002': true, // Ideographic full stop
	'\uFF0C': true, // Fullwidth comma
	'\uFF0E': true, // Fullwidth full stop
	'\uFF01': true, // Fullwidth exclamation mark
	'\uFF1F': true, // Fullwidth question mark
	'\uFF1A': true, // Fullwidth colon
	'\uFF1B': true, // Fullwidth semicolon
	'\u3009': true, // Right angle bracket
	'\u300B': true, // Right double angle bracket
	'\u300D': true, // Right corner bracket
	'\u300F': true, // Right white corner bracket
	'\u3011': true, // Right black lenticular bracket
	'\u3015': true, // Right tortoise shell bracket
	'\u3017': true, // Right white lenticular bracket
	'\u3019': true, // Right white tortoise shell bracket
	'\u301B': true, // Right white square bracket
	'\u301E': true, // Double prime quotation mark
	'\u301F': true, // Low double prime quotation mark
	'\uFF09': true, // Fullwidth right parenthesis
	'\uFF3D': true, // Fullwidth right square bracket
	'\uFF5D': true, // Fullwidth right curly bracket
	'\uFF60': true, // Fullwidth right white parenthesis
	'\u3005': true, // Ideographic iteration mark
	'\u303B': true, // Vertical ideographic iteration mark
	'\u309D': true, // Hiragana iteration mark
	'\u309E': true, // Hiragana voiced iteration mark
	'\u30FD': true, // Katakana iteration mark
	'\u30FE': true, // Katakana voiced iteration mark
	'\u30FC': true, // Katakana-Hiragana prolonged sound mark
	'\u30A0': true, // Katakana-Hiragana double hyphen

	// Small Hiragana (cannot start a line)
	'\u3041': true, // Small A
	'\u3043': true, // Small I
	'\u3045': true, // Small U
	'\u3047': true, // Small E
	'\u3049': true, // Small O
	'\u3063': true, // Small TSU
	'\u3083': true, // Small YA
	'\u3085': true, // Small YU
	'\u3087': true, // Small YO
	'\u308E': true, // Small WA
	'\u3095': true, // Small KA
	'\u3096': true, // Small KE

	// Small Katakana (cannot start a line)
	'\u30A1': true, // Small A
	'\u30A3': true, // Small I
	'\u30A5': true, // Small U
	'\u30A7': true, // Small E
	'\u30A9': true, // Small O
	'\u30C3': true, // Small TSU
	'\u30E3': true, // Small YA
	'\u30E5': true, // Small YU
	'\u30E7': true, // Small YO
	'\u30EE': true, // Small WA
	'\u30F5': true, // Small KA
	'\u30F6': true, // Small KE

	// Halfwidth Katakana small characters
	'\uFF67': true, // Halfwidth small A
	'\uFF68': true, // Halfwidth small I
	'\uFF69': true, // Halfwidth small U
	'\uFF6A': true, // Halfwidth small E
	'\uFF6B': true, // Halfwidth small O
	'\uFF6C': true, // Halfwidth small YA
	'\uFF6D': true, // Halfwidth small YU
	'\uFF6E': true, // Halfwidth small YO
	'\uFF6F': true, // Halfwidth small TSU
	'\uFF70': true, // Halfwidth prolonged sound mark

	// Additional closing punctuation
	'\u2019': true, // Right single quotation mark
	'\u201D': true, // Right double quotation mark
	'\u00BB': true, // Right-pointing double angle quotation mark
}

// noBreakAfterChars contains characters after which a line break should not occur.
// These include opening brackets, currency symbols, etc.
var noBreakAfterChars = map[rune]bool{
	// Opening brackets and parentheses
	'\u3008': true, // Left angle bracket
	'\u300A': true, // Left double angle bracket
	'\u300C': true, // Left corner bracket
	'\u300E': true, // Left white corner bracket
	'\u3010': true, // Left black lenticular bracket
	'\u3014': true, // Left tortoise shell bracket
	'\u3016': true, // Left white lenticular bracket
	'\u3018': true, // Left white tortoise shell bracket
	'\u301A': true, // Left white square bracket
	'\u301D': true, // Reversed double prime quotation mark
	'\uFF08': true, // Fullwidth left parenthesis
	'\uFF3B': true, // Fullwidth left square bracket
	'\uFF5B': true, // Fullwidth left curly bracket
	'\uFF5F': true, // Fullwidth left white parenthesis

	// Quotation marks
	'\u2018': true, // Left single quotation mark
	'\u201C': true, // Left double quotation mark
	'\u00AB': true, // Left-pointing double angle quotation mark

	// Currency symbols
	'\uFFE5': true, // Fullwidth yen sign
	'\uFF04': true, // Fullwidth dollar sign
	'\uFFE0': true, // Fullwidth cent sign
	'\uFFE1': true, // Fullwidth pound sign
}

// GetCJKBreakRule returns the line breaking rule for a CJK character.
func GetCJKBreakRule(r rune) CJKBreakRule {
	noBreakBefore := noBreakBeforeChars[r]
	noBreakAfter := noBreakAfterChars[r]

	if noBreakBefore && noBreakAfter {
		return CJKBreakNoBreakEither
	}
	if noBreakBefore {
		return CJKBreakNoBreakBefore
	}
	if noBreakAfter {
		return CJKBreakNoBreakAfter
	}

	return CJKBreakNormal
}

// IsNoBreakBefore returns true if a line break is prohibited before this character.
func IsNoBreakBefore(r rune) bool {
	rule := GetCJKBreakRule(r)

	return rule == CJKBreakNoBreakBefore ||
		rule == CJKBreakNoBreakEither
}

// IsNoBreakAfter returns true if a line break is prohibited after this character.
func IsNoBreakAfter(r rune) bool {
	rule := GetCJKBreakRule(r)

	return rule == CJKBreakNoBreakAfter ||
		rule == CJKBreakNoBreakEither
}

// IsSmallKana returns true if the rune is a small kana character.
func IsSmallKana(r rune) bool {
	switch r {
	// Small Hiragana
	case '\u3041',
		'\u3043',
		'\u3045',
		'\u3047',
		'\u3049',
		'\u3063',
		'\u3083',
		'\u3085',
		'\u3087',
		'\u308E',
		'\u3095',
		'\u3096':
		return true
	// Small Katakana
	case '\u30A1',
		'\u30A3',
		'\u30A5',
		'\u30A7',
		'\u30A9',
		'\u30C3',
		'\u30E3',
		'\u30E5',
		'\u30E7',
		'\u30EE',
		'\u30F5',
		'\u30F6':
		return true
	// Halfwidth small Katakana
	case '\uFF67',
		'\uFF68',
		'\uFF69',
		'\uFF6A',
		'\uFF6B',
		'\uFF6C',
		'\uFF6D',
		'\uFF6E',
		'\uFF6F':
		return true
	}

	return false
}

// IsProlongedSoundMark returns true if the rune is a prolonged sound mark.
func IsProlongedSoundMark(r rune) bool {
	switch r {
	case '\u30FC', // Katakana-Hiragana prolonged sound mark
		'\uFF70': // Halfwidth prolonged sound mark
		return true
	}

	return false
}

// IsIterationMark returns true if the rune is an iteration mark.
func IsIterationMark(r rune) bool {
	switch r {
	case '\u3005', // Ideographic iteration mark
		'\u303B',           // Vertical ideographic iteration mark
		'\u309D', '\u309E', // Hiragana iteration marks
		'\u30FD', '\u30FE': // Katakana iteration marks
		return true
	}

	return false
}

// CJKBreakChecker provides utilities for checking CJK line breaking rules.
type CJKBreakChecker struct {
	// Strictness level for break checking.
	// 0 = lenient (allow more breaks)
	// 1 = normal (standard kinsoku shori)
	// 2 = strict (additional restrictions)
	strictness int
}

// NewCJKBreakChecker creates a new CJK break checker with normal strictness.
func NewCJKBreakChecker() *CJKBreakChecker {
	return &CJKBreakChecker{strictness: 1}
}

// NewCJKBreakCheckerWithStrictness creates a new CJK break checker with specified strictness.
func NewCJKBreakCheckerWithStrictness(
	strictness int,
) *CJKBreakChecker {
	if strictness < 0 {
		strictness = 0
	}
	if strictness > 2 {
		strictness = 2
	}

	return &CJKBreakChecker{
		strictness: strictness,
	}
}

// CanBreakBetween returns true if a line break is allowed between the two characters.
func (c *CJKBreakChecker) CanBreakBetween(
	before, after rune,
) bool {
	// Check if break is prohibited after the first character
	if IsNoBreakAfter(before) {
		return false
	}

	// Check if break is prohibited before the second character
	if IsNoBreakBefore(after) {
		return false
	}

	// In strict mode, apply additional restrictions
	if c.strictness >= 2 {
		// Don't break between related characters
		if isRelatedCJKPair(before, after) {
			return false
		}
	}

	// Both characters must be CJK for automatic break allowance
	beforeCJK := IsCJK(before)
	afterCJK := IsCJK(after)

	// If both are CJK and neither has break restrictions, allow break
	if beforeCJK && afterCJK {
		return true
	}

	// Mixed content - use default rules
	return false
}

// isRelatedCJKPair checks if two characters form a related pair that shouldn't be broken.
func isRelatedCJKPair(before, after rune) bool {
	// Digit followed by CJK unit suffix
	if before >= '0' && before <= '9' {
		if isCJKUnitSuffix(after) {
			return true
		}
	}

	// CJK followed by its related mark
	if IsCJK(before) &&
		(IsIterationMark(after) || IsProlongedSoundMark(after)) {
		return true
	}

	return false
}

// isCJKUnitSuffix returns true if the rune is a common CJK unit suffix.
func isCJKUnitSuffix(r rune) bool {
	switch r {
	case '\u5E74', // year
		'\u6708', // month
		'\u65E5', // day
		'\u6642', // hour
		'\u5206', // minute
		'\u79D2', // second
		'\u5186', // yen (currency)
		'\u500B', // counter
		'\u4EBA', // person
		'\u672C', // book/long object counter
		'\u679A', // flat object counter
		'\u53F0', // machine/vehicle counter
		'\u56DE': // times/rounds counter
		return true
	}

	return false
}

// CanBreakAt checks if a line break is allowed at the specified rune position in the text.
// The break would occur before the character at the given position.
func (c *CJKBreakChecker) CanBreakAt(
	text string,
	runePos int,
) bool {
	runes := []rune(text)

	if runePos <= 0 || runePos >= len(runes) {
		return false
	}

	before := runes[runePos-1]
	after := runes[runePos]

	return c.CanBreakBetween(before, after)
}

// FindBreakOpportunities returns all positions in the text where line breaks are allowed
// according to CJK rules. Only positions within CJK character sequences are considered.
func (c *CJKBreakChecker) FindBreakOpportunities(
	text string,
) []int {
	runes := []rune(text)
	if len(runes) < 2 {
		return nil
	}

	var opportunities []int

	for i := 1; i < len(runes); i++ {
		before := runes[i-1]
		after := runes[i]

		// Only consider positions where at least one character is CJK
		if !IsCJK(before) && !IsCJK(after) {
			continue
		}

		if c.CanBreakBetween(before, after) {
			opportunities = append(
				opportunities,
				i,
			)
		}
	}

	return opportunities
}

// CJKTextInfo provides information about CJK content in text.
type CJKTextInfo struct {
	// HasCJK indicates whether the text contains any CJK characters.
	HasCJK bool
	// IsPurelyCJK indicates whether the text contains only CJK characters.
	IsPurelyCJK bool
	// HanCount is the number of Han/Chinese ideographs.
	HanCount int
	// HiraganaCount is the number of Hiragana characters.
	HiraganaCount int
	// KatakanaCount is the number of Katakana characters.
	KatakanaCount int
	// HangulCount is the number of Hangul characters.
	HangulCount int
	// SymbolCount is the number of CJK symbols and punctuation.
	SymbolCount int
	// TotalCJKCount is the total number of CJK characters.
	TotalCJKCount int
	// DominantScript is the most common CJK script type in the text.
	DominantScript CJKScriptType
}

// AnalyzeCJKContent analyzes the CJK content in a string.
func AnalyzeCJKContent(text string) CJKTextInfo {
	info := CJKTextInfo{
		DominantScript: CJKScriptNone,
	}

	if len(text) == 0 {
		return info
	}

	totalRunes := 0
	for _, r := range text {
		totalRunes++
		scriptType := GetCJKScriptType(r)

		switch scriptType {
		case CJKScriptHan:
			info.HanCount++
		case CJKScriptHiragana:
			info.HiraganaCount++
		case CJKScriptKatakana:
			info.KatakanaCount++
		case CJKScriptHangul:
			info.HangulCount++
		case CJKScriptSymbol:
			info.SymbolCount++
		case CJKScriptBopomofo:
			// Count as symbol or handle separately if needed
		case CJKScriptNone:
			// Not CJK
		}
	}

	info.TotalCJKCount = info.HanCount + info.HiraganaCount +
		info.KatakanaCount + info.HangulCount + info.SymbolCount
	info.HasCJK = info.TotalCJKCount > 0
	info.IsPurelyCJK = info.TotalCJKCount == totalRunes

	// Determine dominant script (excluding symbols)
	maxCount := 0
	if info.HanCount > maxCount {
		maxCount = info.HanCount
		info.DominantScript = CJKScriptHan
	}
	if info.HiraganaCount > maxCount {
		maxCount = info.HiraganaCount
		info.DominantScript = CJKScriptHiragana
	}
	if info.KatakanaCount > maxCount {
		maxCount = info.KatakanaCount
		info.DominantScript = CJKScriptKatakana
	}
	if info.HangulCount > maxCount {
		info.DominantScript = CJKScriptHangul
	}

	return info
}

// IsFullwidthPunctuation returns true if the rune is a fullwidth punctuation character.
func IsFullwidthPunctuation(r rune) bool {
	switch r {
	case '\uFF01', // Fullwidth exclamation mark
		'\uFF0C', // Fullwidth comma
		'\uFF0E', // Fullwidth full stop
		'\uFF1A', // Fullwidth colon
		'\uFF1B', // Fullwidth semicolon
		'\uFF1F', // Fullwidth question mark
		'\u3001', // Ideographic comma
		'\u3002': // Ideographic full stop
		return true
	}

	return false
}

// IsFullwidthBracket returns true if the rune is a fullwidth bracket character.
func IsFullwidthBracket(r rune) bool {
	switch r {
	case '\uFF08',
		'\uFF09', // Fullwidth parentheses
		'\uFF3B',
		'\uFF3D', // Fullwidth square brackets
		'\uFF5B',
		'\uFF5D', // Fullwidth curly brackets
		'\u3008',
		'\u3009', // Angle brackets
		'\u300A',
		'\u300B', // Double angle brackets
		'\u300C',
		'\u300D', // Corner brackets
		'\u300E',
		'\u300F', // White corner brackets
		'\u3010',
		'\u3011', // Black lenticular brackets
		'\u3014',
		'\u3015', // Tortoise shell brackets
		'\u3016',
		'\u3017', // White lenticular brackets
		'\u3018',
		'\u3019', // White tortoise shell brackets
		'\u301A',
		'\u301B': // White square brackets
		return true
	}

	return false
}

// CJKLineBreaker extends line breaking for CJK text.
type CJKLineBreaker struct {
	*LineBreaker
	cjkChecker *CJKBreakChecker
}

// NewCJKLineBreaker creates a new CJKLineBreaker.
func NewCJKLineBreaker() *CJKLineBreaker {
	return &CJKLineBreaker{
		LineBreaker: NewLineBreaker(),
		cjkChecker:  NewCJKBreakChecker(),
	}
}

// NewCJKLineBreakerWithStrictness creates a new CJKLineBreaker with specified strictness.
func NewCJKLineBreakerWithStrictness(
	strictness int,
) *CJKLineBreaker {
	return &CJKLineBreaker{
		LineBreaker: NewLineBreaker(),
		cjkChecker: NewCJKBreakCheckerWithStrictness(
			strictness,
		),
	}
}

// FindBreakOpportunities returns break opportunities with CJK-specific rules applied.
func (clb *CJKLineBreaker) FindBreakOpportunities(
	text string,
) []BreakOpportunity {
	if len(text) == 0 {
		return nil
	}

	// Get base opportunities from LineBreaker
	baseOpportunities := clb.LineBreaker.FindBreakOpportunities(
		text,
	)

	runes := []rune(text)
	if len(runes) < 2 {
		return baseOpportunities
	}

	// Build byte position map
	bytePositions := make([]int, len(runes)+1)
	pos := 0
	for i, r := range runes {
		bytePositions[i] = pos
		pos += len(string(r))
	}
	bytePositions[len(runes)] = pos

	// Create a map of existing opportunities by position
	oppMap := make(map[int]*BreakOpportunity)
	for i := range baseOpportunities {
		oppMap[baseOpportunities[i].Position] = &baseOpportunities[i]
	}

	// Apply CJK-specific rules
	var result []BreakOpportunity

	for i := 1; i <= len(runes); i++ {
		bytePos := bytePositions[i]
		existingOpp, hasExisting := oppMap[bytePos]

		// Check for mandatory breaks first
		if hasExisting &&
			existingOpp.Type == BreakMandatory {
			result = append(result, *existingOpp)

			continue
		}

		if i >= len(runes) {
			continue
		}

		before := runes[i-1]
		after := runes[i]

		beforeCJK := IsCJK(before)
		afterCJK := IsCJK(after)

		// If neither character is CJK, use base opportunity
		if !beforeCJK && !afterCJK {
			if hasExisting &&
				existingOpp.Type != BreakProhibited {
				result = append(
					result,
					*existingOpp,
				)
			}

			continue
		}

		// Apply CJK break rules
		var breakType BreakOpportunityType

		if clb.cjkChecker.CanBreakBetween(
			before,
			after,
		) {
			breakType = BreakAllowed
		} else {
			breakType = BreakProhibited
		}

		if breakType != BreakProhibited {
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
