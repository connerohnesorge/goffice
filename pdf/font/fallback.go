// fallback.go provides configurable font fallback chains for font substitution when
// a requested font is not available. This is essential for PDF rendering when Office
// documents reference fonts that may not be installed on the rendering system.

package font

import (
	"errors"
	"strings"
	"sync"
)

// FontCategory represents the general category of a font for fallback purposes.
type FontCategory int

const (
	// CategorySerif represents serif fonts (e.g., Times New Roman, Georgia)
	CategorySerif FontCategory = iota
	// CategorySansSerif represents sans-serif fonts (e.g., Arial, Calibri)
	CategorySansSerif
	// CategoryMonospace represents monospace fonts (e.g., Courier New, Consolas)
	CategoryMonospace
	// CategorySymbol represents symbol fonts (e.g., Symbol, Wingdings)
	CategorySymbol
	// CategoryCursive represents cursive/script fonts
	CategoryCursive
	// CategoryFantasy represents decorative/fantasy fonts
	CategoryFantasy
)

// String returns the human-readable name of the font category.
func (c FontCategory) String() string {
	switch c {
	case CategorySerif:
		return "Serif"
	case CategorySansSerif:
		return "Sans-Serif"
	case CategoryMonospace:
		return "Monospace"
	case CategorySymbol:
		return "Symbol"
	case CategoryCursive:
		return "Cursive"
	case CategoryFantasy:
		return "Fantasy"
	default:
		return "Unknown"
	}
}

// FallbackRule defines a rule for font substitution.
// When a font family matches the Pattern, the Fallback font is used instead.
type FallbackRule struct {
	Pattern  string    // Font family pattern (supports * wildcard)
	Fallback string    // Fallback font family to use
	Style    FontStyle // Optional: specific style to use (-1 means inherit)
}

// ErrNoFallbackFound is returned when no fallback font could be found.
var ErrNoFallbackFound = errors.New(
	"no fallback font found",
)

// FallbackChain provides configurable font fallback functionality.
// It maintains rules for font substitution and default fonts for each category.
type FallbackChain struct {
	mu       sync.RWMutex
	rules    []FallbackRule
	defaults map[FontCategory]*Font
}

// NewFallbackChain creates a new FallbackChain with sensible defaults.
// It includes mappings for common Microsoft Office fonts to open alternatives.
func NewFallbackChain() *FallbackChain {
	fc := &FallbackChain{
		rules:    make([]FallbackRule, 0),
		defaults: make(map[FontCategory]*Font),
	}

	// Add default fallback rules for common MS Office fonts
	// These map proprietary fonts to their open-source alternatives or generic fallbacks

	// Serif fonts
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Times New Roman",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Georgia",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Cambria",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Garamond",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Book Antiqua",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Palatino Linotype",
			Fallback: "Liberation Serif",
			Style:    -1,
		},
	)

	// Sans-serif fonts
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Arial",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Calibri",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Helvetica",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Verdana",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Tahoma",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Segoe UI",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Trebuchet MS",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)

	// Monospace fonts
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Courier New",
			Fallback: "Liberation Mono",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Consolas",
			Fallback: "Liberation Mono",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Lucida Console",
			Fallback: "DejaVu Sans Mono",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Monaco",
			Fallback: "Liberation Mono",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Andale Mono",
			Fallback: "Liberation Mono",
			Style:    -1,
		},
	)

	// Symbol fonts - these are special and may need specific handling
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Symbol",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Wingdings",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Wingdings 2",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Wingdings 3",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Webdings",
			Fallback: "DejaVu Sans",
			Style:    -1,
		},
	)

	// Cursive/script fonts
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Comic Sans MS",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Bradley Hand*",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)
	fc.rules = append(
		fc.rules,
		FallbackRule{
			Pattern:  "Brush Script*",
			Fallback: "Liberation Sans",
			Style:    -1,
		},
	)

	return fc
}

// AddRule adds a custom fallback rule.
// The pattern supports * as a wildcard for any characters.
// Style can be set to -1 to inherit the original style.
func (fc *FallbackChain) AddRule(
	pattern, fallback string,
	style FontStyle,
) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	// Prepend new rules so they take priority over defaults
	fc.rules = append([]FallbackRule{{
		Pattern:  pattern,
		Fallback: fallback,
		Style:    style,
	}}, fc.rules...)
}

// AddRuleToEnd adds a custom fallback rule at the end of the chain.
// This is useful for adding low-priority fallbacks.
func (fc *FallbackChain) AddRuleToEnd(
	pattern, fallback string,
	style FontStyle,
) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	fc.rules = append(fc.rules, FallbackRule{
		Pattern:  pattern,
		Fallback: fallback,
		Style:    style,
	})
}

// SetDefault sets the default font for a specific category.
// This font is used when no specific rule matches.
func (fc *FallbackChain) SetDefault(
	category FontCategory,
	font *Font,
) {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.defaults[category] = font
}

// GetDefault returns the default font for a category.
func (fc *FallbackChain) GetDefault(
	category FontCategory,
) *Font {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	return fc.defaults[category]
}

// Resolve attempts to find a font, using fallback rules if the exact font is not available.
// It tries the following in order:
// 1. Get the exact font from the cache
// 2. Check rules for a matching pattern and try to get the fallback font
// 3. Classify the font and use the category default
// Returns ErrNoFallbackFound if no suitable font could be found.
func (fc *FallbackChain) Resolve(
	family string,
	style FontStyle,
	cache *FontCache,
) (*Font, error) {
	if cache == nil {
		return nil, errors.New(
			"font cache is nil",
		)
	}

	// Step 1: Try to get the exact font from cache
	if font, ok := cache.Get(family, style); ok {
		return font, nil
	}

	fc.mu.RLock()
	rules := fc.rules
	defaults := fc.defaults
	fc.mu.RUnlock()

	// Step 2: Check rules for matching patterns
	for _, rule := range rules {
		if MatchPattern(rule.Pattern, family) {
			// Determine which style to use
			fallbackStyle := style
			if rule.Style >= 0 {
				fallbackStyle = rule.Style
			}

			// Try to get the fallback font
			if font, ok := cache.Get(rule.Fallback, fallbackStyle); ok {
				return font, nil
			}

			// Try regular style if the specific style isn't available
			if fallbackStyle != StyleRegular {
				if font, ok := cache.Get(rule.Fallback, StyleRegular); ok {
					return font, nil
				}
			}
		}
	}

	// Step 3: Classify font and use category default
	category := ClassifyFont(family)
	if defaultFont, ok := defaults[category]; ok &&
		defaultFont != nil {
		return defaultFont, nil
	}

	// Step 4: Try any category default as last resort
	for _, font := range defaults {
		if font != nil {
			return font, nil
		}
	}

	return nil, ErrNoFallbackFound
}

// ClassifyFont determines the category of a font based on its family name.
// This uses heuristics based on common font naming conventions.
func ClassifyFont(family string) FontCategory {
	lower := strings.ToLower(family)

	// Check for symbol fonts first (most specific)
	symbolKeywords := []string{
		"symbol",
		"wingding",
		"webding",
		"dingbat",
		"icon",
		"emoji",
	}
	for _, kw := range symbolKeywords {
		if strings.Contains(lower, kw) {
			return CategorySymbol
		}
	}

	// Check for monospace fonts
	monoKeywords := []string{
		"mono",
		"courier",
		"consolas",
		"console",
		"fixed",
		"typewriter",
		"code",
		"terminal",
	}
	for _, kw := range monoKeywords {
		if strings.Contains(lower, kw) {
			return CategoryMonospace
		}
	}

	// Check for cursive/script fonts
	cursiveKeywords := []string{
		"script",
		"cursive",
		"hand",
		"brush",
		"comic",
	}
	for _, kw := range cursiveKeywords {
		if strings.Contains(lower, kw) {
			return CategoryCursive
		}
	}

	// Check for fantasy/decorative fonts
	fantasyKeywords := []string{
		"fantasy",
		"decorative",
		"display",
		"ornament",
		"blackletter",
		"gothic",
		"fraktur",
	}
	for _, kw := range fantasyKeywords {
		if strings.Contains(lower, kw) {
			return CategoryFantasy
		}
	}

	// Check for sans-serif fonts
	sansKeywords := []string{
		"sans",
		"arial",
		"calibri",
		"helvetica",
		"verdana",
		"tahoma",
		"segoe",
		"roboto",
		"open sans",
		"lato",
		"ubuntu",
	}
	for _, kw := range sansKeywords {
		if strings.Contains(lower, kw) {
			return CategorySansSerif
		}
	}

	// Check for serif fonts (explicit indicators)
	serifKeywords := []string{
		"serif",
		"times",
		"georgia",
		"cambria",
		"garamond",
		"palatino",
		"book",
		"roman",
	}
	for _, kw := range serifKeywords {
		if strings.Contains(lower, kw) {
			return CategorySerif
		}
	}

	// Default to serif (most common in documents)
	return CategorySerif
}

// MatchPattern checks if a font family name matches a pattern.
// The pattern supports * as a wildcard for any sequence of characters.
// Matching is case-insensitive.
func MatchPattern(pattern, family string) bool {
	pattern = strings.ToLower(pattern)
	family = strings.ToLower(family)

	// Handle exact match (no wildcards)
	if !strings.Contains(pattern, "*") {
		return pattern == family
	}

	// Split pattern by * and match each part
	parts := strings.Split(pattern, "*")

	// Filter out empty parts
	nonEmpty := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			nonEmpty = append(nonEmpty, p)
		}
	}

	// If pattern is just "*", match everything
	if len(nonEmpty) == 0 {
		return true
	}

	// Check if pattern starts with a fixed string (no leading *)
	pos := 0
	if !strings.HasPrefix(pattern, "*") {
		if !strings.HasPrefix(
			family,
			nonEmpty[0],
		) {
			return false
		}
		pos = len(nonEmpty[0])
		nonEmpty = nonEmpty[1:]
	}

	// Check if pattern ends with a fixed string (no trailing *)
	if !strings.HasSuffix(pattern, "*") &&
		len(nonEmpty) > 0 {
		lastPart := nonEmpty[len(nonEmpty)-1]
		if !strings.HasSuffix(family, lastPart) {
			return false
		}
		nonEmpty = nonEmpty[:len(nonEmpty)-1]
		family = family[:len(family)-len(lastPart)]
	}

	// Match remaining parts in order
	for _, part := range nonEmpty {
		idx := strings.Index(family[pos:], part)
		if idx == -1 {
			return false
		}
		pos += idx + len(part)
	}

	return true
}

// Rules returns a copy of the current fallback rules.
func (fc *FallbackChain) Rules() []FallbackRule {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	result := make([]FallbackRule, len(fc.rules))
	copy(result, fc.rules)

	return result
}

// ClearRules removes all fallback rules.
func (fc *FallbackChain) ClearRules() {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	fc.rules = make([]FallbackRule, 0)
}

// RemoveRule removes the first rule matching the given pattern.
// Returns true if a rule was removed.
func (fc *FallbackChain) RemoveRule(
	pattern string,
) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	for i, rule := range fc.rules {
		if strings.EqualFold(
			rule.Pattern,
			pattern,
		) {
			fc.rules = append(
				fc.rules[:i],
				fc.rules[i+1:]...)

			return true
		}
	}

	return false
}

// FindRule finds the first rule matching a font family.
// Returns the rule and true if found, or an empty rule and false if not.
func (fc *FallbackChain) FindRule(
	family string,
) (FallbackRule, bool) {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	for _, rule := range fc.rules {
		if MatchPattern(rule.Pattern, family) {
			return rule, true
		}
	}

	return FallbackRule{}, false
}

// commonSerifFonts lists fonts commonly classified as serif.
var commonSerifFonts = []string{
	"Times New Roman", "Georgia", "Cambria", "Garamond",
	"Book Antiqua", "Palatino Linotype", "Century", "Bookman",
	"Liberation Serif", "DejaVu Serif", "Noto Serif",
}

// commonSansSerifFonts lists fonts commonly classified as sans-serif.
var commonSansSerifFonts = []string{
	"Arial", "Calibri", "Helvetica", "Verdana", "Tahoma",
	"Segoe UI", "Trebuchet MS", "Century Gothic", "Candara",
	"Liberation Sans", "DejaVu Sans", "Noto Sans", "Roboto", "Open Sans",
}

// commonMonospaceFonts lists fonts commonly classified as monospace.
var commonMonospaceFonts = []string{
	"Courier New", "Consolas", "Lucida Console", "Monaco",
	"Andale Mono", "Liberation Mono", "DejaVu Sans Mono", "Noto Mono",
	"Source Code Pro", "Fira Mono",
}

// commonSymbolFonts lists fonts commonly classified as symbol.
var commonSymbolFonts = []string{
	"Symbol", "Wingdings", "Wingdings 2", "Wingdings 3",
	"Webdings", "Marlett", "MS Outlook",
}

// IsSerifFont returns true if the font family is a known serif font.
func IsSerifFont(family string) bool {
	lower := strings.ToLower(family)
	for _, f := range commonSerifFonts {
		if strings.ToLower(f) == lower {
			return true
		}
	}

	return false
}

// IsSansSerifFont returns true if the font family is a known sans-serif font.
func IsSansSerifFont(family string) bool {
	lower := strings.ToLower(family)
	for _, f := range commonSansSerifFonts {
		if strings.ToLower(f) == lower {
			return true
		}
	}

	return false
}

// IsMonospaceFont returns true if the font family is a known monospace font.
func IsMonospaceFont(family string) bool {
	lower := strings.ToLower(family)
	for _, f := range commonMonospaceFonts {
		if strings.ToLower(f) == lower {
			return true
		}
	}

	return false
}

// IsSymbolFont returns true if the font family is a known symbol font.
func IsSymbolFont(family string) bool {
	lower := strings.ToLower(family)
	for _, f := range commonSymbolFonts {
		if strings.ToLower(f) == lower {
			return true
		}
	}

	return false
}
