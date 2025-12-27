package font

import (
	"testing"
)

func TestFontCategory_String(t *testing.T) {
	tests := []struct {
		category FontCategory
		expected string
	}{
		{CategorySerif, "Serif"},
		{CategorySansSerif, "Sans-Serif"},
		{CategoryMonospace, "Monospace"},
		{CategorySymbol, "Symbol"},
		{CategoryCursive, "Cursive"},
		{CategoryFantasy, "Fantasy"},
		{FontCategory(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.category.String(); got != tt.expected {
				t.Errorf(
					"FontCategory.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestNewFallbackChain(t *testing.T) {
	fc := NewFallbackChain()

	if fc == nil {
		t.Fatal("NewFallbackChain() returned nil")
	}

	// Check that default rules are present
	rules := fc.Rules()
	if len(rules) == 0 {
		t.Error(
			"Expected default rules to be present",
		)
	}

	// Check for some known default mappings
	foundArial := false
	foundTimesNewRoman := false
	foundCourierNew := false

	for _, rule := range rules {
		switch rule.Pattern {
		case "Arial":
			foundArial = true
		case "Times New Roman":
			foundTimesNewRoman = true
		case "Courier New":
			foundCourierNew = true
		}
	}

	if !foundArial {
		t.Error("Expected Arial fallback rule")
	}
	if !foundTimesNewRoman {
		t.Error(
			"Expected Times New Roman fallback rule",
		)
	}
	if !foundCourierNew {
		t.Error(
			"Expected Courier New fallback rule",
		)
	}
}

func TestFallbackChain_AddRule(t *testing.T) {
	fc := NewFallbackChain()
	initialCount := len(fc.Rules())

	fc.AddRule(
		"CustomFont",
		"FallbackFont",
		StyleRegular,
	)

	rules := fc.Rules()
	if len(rules) != initialCount+1 {
		t.Errorf(
			"Expected %d rules, got %d",
			initialCount+1,
			len(rules),
		)
	}

	// New rules should be prepended (higher priority)
	if rules[0].Pattern != "CustomFont" {
		t.Errorf(
			"Expected new rule to be first, got pattern: %s",
			rules[0].Pattern,
		)
	}
	if rules[0].Fallback != "FallbackFont" {
		t.Errorf(
			"Expected fallback to be FallbackFont, got: %s",
			rules[0].Fallback,
		)
	}
	if rules[0].Style != StyleRegular {
		t.Errorf(
			"Expected style to be StyleRegular, got: %v",
			rules[0].Style,
		)
	}
}

func TestFallbackChain_AddRuleToEnd(
	t *testing.T,
) {
	fc := NewFallbackChain()
	initialCount := len(fc.Rules())

	fc.AddRuleToEnd(
		"LastFont",
		"LastFallback",
		StyleBold,
	)

	rules := fc.Rules()
	if len(rules) != initialCount+1 {
		t.Errorf(
			"Expected %d rules, got %d",
			initialCount+1,
			len(rules),
		)
	}

	// Rule should be at the end
	last := rules[len(rules)-1]
	if last.Pattern != "LastFont" {
		t.Errorf(
			"Expected new rule to be last, got pattern: %s",
			last.Pattern,
		)
	}
}

func TestFallbackChain_SetDefault(t *testing.T) {
	fc := NewFallbackChain()

	// Create a mock font
	mockFont := &Font{
		Family: "DefaultSerif",
		Style:  StyleRegular,
	}

	fc.SetDefault(CategorySerif, mockFont)

	got := fc.GetDefault(CategorySerif)
	if got != mockFont {
		t.Error(
			"GetDefault did not return the set default font",
		)
	}

	// Check unset category returns nil
	got = fc.GetDefault(CategoryFantasy)
	if got != nil {
		t.Error("Expected nil for unset category")
	}
}

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		family   string
		expected bool
	}{
		// Exact match
		{"Arial", "Arial", true},
		{"Arial", "arial", true},
		{"arial", "Arial", true},
		{"Arial", "Helvetica", false},

		// Single wildcard at end
		{"Arial*", "Arial", true},
		{"Arial*", "Arial Bold", true},
		{"Arial*", "ArialMT", true},
		{"Arial*", "Helvetica", false},

		// Single wildcard at start
		{"*Mono", "DejaVu Mono", true},
		{"*Mono", "Mono", true},
		{"*Mono", "Monaco", false},

		// Wildcard in middle
		{
			"Liberation*Serif",
			"Liberation Serif",
			true,
		},
		{
			"Liberation*Serif",
			"LiberationSerif",
			true,
		},
		{
			"Liberation*Serif",
			"Liberation Sans",
			false,
		},

		// Multiple wildcards
		{"*Sans*", "DejaVu Sans", true},
		{"*Sans*", "Open Sans Condensed", true},
		{"*Sans*", "Liberation Sans Mono", true},
		{"*Sans*", "Arial", false},

		// Just wildcard
		{"*", "Anything", true},
		{"*", "", true},

		// Empty pattern
		{"", "", true},
		{"", "Arial", false},

		// Case insensitivity
		{"ARIAL", "arial", true},
		{"Arial*", "ARIAL BOLD", true},
	}

	for _, tt := range tests {
		t.Run(
			tt.pattern+"_"+tt.family,
			func(t *testing.T) {
				got := MatchPattern(
					tt.pattern,
					tt.family,
				)
				if got != tt.expected {
					t.Errorf(
						"MatchPattern(%q, %q) = %v, want %v",
						tt.pattern,
						tt.family,
						got,
						tt.expected,
					)
				}
			},
		)
	}
}

func TestClassifyFont(t *testing.T) {
	tests := []struct {
		family   string
		expected FontCategory
	}{
		// Serif fonts
		{"Times New Roman", CategorySerif},
		{"Georgia", CategorySerif},
		{"Cambria", CategorySerif},
		{"Liberation Serif", CategorySerif},
		{"DejaVu Serif", CategorySerif},

		// Sans-serif fonts
		{"Arial", CategorySansSerif},
		{"Calibri", CategorySansSerif},
		{"Helvetica", CategorySansSerif},
		{"Verdana", CategorySansSerif},
		{"Open Sans", CategorySansSerif},
		{"Liberation Sans", CategorySansSerif},
		{"Roboto", CategorySansSerif},

		// Monospace fonts
		{"Courier New", CategoryMonospace},
		{"Consolas", CategoryMonospace},
		{"Liberation Mono", CategoryMonospace},
		{"DejaVu Sans Mono", CategoryMonospace},
		{"Source Code Pro", CategoryMonospace},

		// Symbol fonts
		{"Symbol", CategorySymbol},
		{"Wingdings", CategorySymbol},
		{"Webdings", CategorySymbol},
		{"Emoji One", CategorySymbol},

		// Cursive fonts
		{"Comic Sans MS", CategoryCursive},
		{"Brush Script", CategoryCursive},
		{"Bradley Hand", CategoryCursive},

		// Fantasy fonts
		{"Blackletter", CategoryFantasy},
		{"Display Pro", CategoryFantasy},

		// Unknown defaults to serif
		{"Some Unknown Font", CategorySerif},
	}

	for _, tt := range tests {
		t.Run(tt.family, func(t *testing.T) {
			got := ClassifyFont(tt.family)
			if got != tt.expected {
				t.Errorf(
					"ClassifyFont(%q) = %v, want %v",
					tt.family,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestFallbackChain_Resolve(t *testing.T) {
	fc := NewFallbackChain()
	cache := NewFontCache(10)

	// Test with empty cache - should return error
	_, err := fc.Resolve(
		"Arial",
		StyleRegular,
		cache,
	)
	if err != ErrNoFallbackFound {
		t.Errorf(
			"Expected ErrNoFallbackFound, got: %v",
			err,
		)
	}

	// Add a fallback font to the cache
	fallbackFont := &Font{
		Family: "Liberation Sans",
		Style:  StyleRegular,
	}
	cache.Put(fallbackFont)

	// Now Arial should resolve to Liberation Sans
	resolved, err := fc.Resolve(
		"Arial",
		StyleRegular,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved != fallbackFont {
		t.Error(
			"Expected to resolve to fallback font",
		)
	}

	// Add exact font and it should be preferred
	exactFont := &Font{
		Family: "Arial",
		Style:  StyleRegular,
	}
	cache.Put(exactFont)

	resolved, err = fc.Resolve(
		"Arial",
		StyleRegular,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved != exactFont {
		t.Error(
			"Expected to resolve to exact font",
		)
	}
}

func TestFallbackChain_ResolveWithDefault(
	t *testing.T,
) {
	fc := NewFallbackChain()
	fc.ClearRules() // Start fresh without default rules
	cache := NewFontCache(10)

	// Set a default font for serif category
	defaultFont := &Font{
		Family: "Default Serif",
		Style:  StyleRegular,
	}
	fc.SetDefault(CategorySerif, defaultFont)

	// Unknown font should resolve to category default
	resolved, err := fc.Resolve(
		"Unknown Font",
		StyleRegular,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved != defaultFont {
		t.Error(
			"Expected to resolve to default font",
		)
	}
}

func TestFallbackChain_ResolveNilCache(
	t *testing.T,
) {
	fc := NewFallbackChain()

	_, err := fc.Resolve(
		"Arial",
		StyleRegular,
		nil,
	)
	if err == nil {
		t.Error("Expected error for nil cache")
	}
}

func TestFallbackChain_RemoveRule(t *testing.T) {
	fc := NewFallbackChain()

	// Find initial count
	initialCount := len(fc.Rules())

	// Remove an existing rule
	removed := fc.RemoveRule("Arial")
	if !removed {
		t.Error("Expected to remove Arial rule")
	}

	if len(fc.Rules()) != initialCount-1 {
		t.Errorf(
			"Expected %d rules after removal, got %d",
			initialCount-1,
			len(fc.Rules()),
		)
	}

	// Try to remove non-existent rule
	removed = fc.RemoveRule("NonExistentFont")
	if removed {
		t.Error(
			"Did not expect to remove non-existent rule",
		)
	}
}

func TestFallbackChain_FindRule(t *testing.T) {
	fc := NewFallbackChain()

	// Find existing rule
	rule, found := fc.FindRule("Arial")
	if !found {
		t.Error("Expected to find Arial rule")
	}
	if rule.Fallback != "Liberation Sans" {
		t.Errorf(
			"Expected Liberation Sans fallback, got: %s",
			rule.Fallback,
		)
	}

	// Find with wildcard match
	fc.AddRule(
		"Test*",
		"TestFallback",
		StyleRegular,
	)
	rule, found = fc.FindRule("TestFont")
	if !found {
		t.Error(
			"Expected to find rule for TestFont",
		)
	}
	if rule.Fallback != "TestFallback" {
		t.Errorf(
			"Expected TestFallback, got: %s",
			rule.Fallback,
		)
	}

	// Non-existent
	_, found = fc.FindRule(
		"CompletelyUnknownFont12345",
	)
	if found {
		t.Error(
			"Did not expect to find rule for unknown font",
		)
	}
}

func TestFallbackChain_ClearRules(t *testing.T) {
	fc := NewFallbackChain()

	if len(fc.Rules()) == 0 {
		t.Fatal("Expected initial rules")
	}

	fc.ClearRules()

	if len(fc.Rules()) != 0 {
		t.Error("Expected no rules after clear")
	}
}

func TestIsSerifFont(t *testing.T) {
	if !IsSerifFont("Times New Roman") {
		t.Error(
			"Expected Times New Roman to be serif",
		)
	}
	if !IsSerifFont(
		"times new roman",
	) { // Case insensitive
		t.Error(
			"Expected times new roman to be serif",
		)
	}
	if IsSerifFont("Arial") {
		t.Error("Expected Arial to not be serif")
	}
}

func TestIsSansSerifFont(t *testing.T) {
	if !IsSansSerifFont("Arial") {
		t.Error("Expected Arial to be sans-serif")
	}
	if !IsSansSerifFont("Calibri") {
		t.Error(
			"Expected Calibri to be sans-serif",
		)
	}
	if IsSansSerifFont("Times New Roman") {
		t.Error(
			"Expected Times New Roman to not be sans-serif",
		)
	}
}

func TestIsMonospaceFont(t *testing.T) {
	if !IsMonospaceFont("Courier New") {
		t.Error(
			"Expected Courier New to be monospace",
		)
	}
	if !IsMonospaceFont("Consolas") {
		t.Error(
			"Expected Consolas to be monospace",
		)
	}
	if IsMonospaceFont("Arial") {
		t.Error(
			"Expected Arial to not be monospace",
		)
	}
}

func TestIsSymbolFont(t *testing.T) {
	if !IsSymbolFont("Symbol") {
		t.Error(
			"Expected Symbol to be symbol font",
		)
	}
	if !IsSymbolFont("Wingdings") {
		t.Error(
			"Expected Wingdings to be symbol font",
		)
	}
	if IsSymbolFont("Arial") {
		t.Error(
			"Expected Arial to not be symbol font",
		)
	}
}

func TestFallbackChain_ResolveStyleInheritance(
	t *testing.T,
) {
	fc := NewFallbackChain()
	fc.ClearRules()
	cache := NewFontCache(10)

	// Add a rule that inherits style
	fc.AddRule("CustomFont", "Fallback", -1)

	// Add fallback fonts with different styles
	fallbackRegular := &Font{
		Family: "Fallback",
		Style:  StyleRegular,
	}
	fallbackBold := &Font{
		Family: "Fallback",
		Style:  StyleBold,
	}
	cache.Put(fallbackRegular)
	cache.Put(fallbackBold)

	// Request bold, should get bold fallback
	resolved, err := fc.Resolve(
		"CustomFont",
		StyleBold,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved.Style != StyleBold {
		t.Errorf(
			"Expected bold style, got: %v",
			resolved.Style,
		)
	}

	// Request regular, should get regular fallback
	resolved, err = fc.Resolve(
		"CustomFont",
		StyleRegular,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved.Style != StyleRegular {
		t.Errorf(
			"Expected regular style, got: %v",
			resolved.Style,
		)
	}
}

func TestFallbackChain_ResolveStyleOverride(
	t *testing.T,
) {
	fc := NewFallbackChain()
	fc.ClearRules()
	cache := NewFontCache(10)

	// Add a rule that forces regular style
	fc.AddRule(
		"ForcedRegular*",
		"Fallback",
		StyleRegular,
	)

	// Add only regular fallback
	fallbackRegular := &Font{
		Family: "Fallback",
		Style:  StyleRegular,
	}
	cache.Put(fallbackRegular)

	// Request bold, should still get regular due to override
	resolved, err := fc.Resolve(
		"ForcedRegularFont",
		StyleBold,
		cache,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resolved.Style != StyleRegular {
		t.Errorf(
			"Expected regular style override, got: %v",
			resolved.Style,
		)
	}
}

func TestFallbackChain_Concurrency(_ *testing.T) {
	fc := NewFallbackChain()
	cache := NewFontCache(100)

	// Add some fonts to cache
	for i := range 10 {
		cache.Put(
			&Font{
				Family: "Liberation Sans",
				Style:  FontStyle(i % 4),
			},
		)
	}

	// Run concurrent operations
	done := make(chan bool)
	for i := range 10 {
		go func(_ int) {
			for range 100 {
				_, _ = fc.Resolve(
					"Arial",
					StyleRegular,
					cache,
				)
				_, _ = fc.FindRule("Arial")
				_ = fc.Rules()
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines
	for range 10 {
		<-done
	}
}
