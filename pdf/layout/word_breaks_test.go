package layout

import (
	"testing"
)

func TestDefaultWordBreakOptions(t *testing.T) {
	opts := DefaultWordBreakOptions()

	// Verify all defaults are true
	if !opts.BreakAfterSlash {
		t.Error(
			"BreakAfterSlash should be true by default",
		)
	}
	if !opts.BreakAfterBackslash {
		t.Error(
			"BreakAfterBackslash should be true by default",
		)
	}
	if !opts.BreakAfterEquals {
		t.Error(
			"BreakAfterEquals should be true by default",
		)
	}
	if !opts.BreakAfterColon {
		t.Error(
			"BreakAfterColon should be true by default",
		)
	}
	if !opts.KeepEmailsTogether {
		t.Error(
			"KeepEmailsTogether should be true by default",
		)
	}
	if !opts.KeepURLsTogether {
		t.Error(
			"KeepURLsTogether should be true by default",
		)
	}
	if !opts.KeepNumberAbbreviations {
		t.Error(
			"KeepNumberAbbreviations should be true by default",
		)
	}
	if !opts.KeepCurrencyWithNumbers {
		t.Error(
			"KeepCurrencyWithNumbers should be true by default",
		)
	}
	if !opts.KeepUnitsWithNumbers {
		t.Error(
			"KeepUnitsWithNumbers should be true by default",
		)
	}
	if !opts.KeepTitlesWithNames {
		t.Error(
			"KeepTitlesWithNames should be true by default",
		)
	}
	if !opts.HyphenateCompoundWords {
		t.Error(
			"HyphenateCompoundWords should be true by default",
		)
	}
	if !opts.HandleDashes {
		t.Error(
			"HandleDashes should be true by default",
		)
	}
	if !opts.SupportSoftHyphen {
		t.Error(
			"SupportSoftHyphen should be true by default",
		)
	}
}

func TestNewWordLineBreaker(t *testing.T) {
	wlb := NewWordLineBreaker()
	if wlb == nil {
		t.Fatal("NewWordLineBreaker returned nil")
	}
	if wlb.LineBreaker == nil {
		t.Error("LineBreaker should not be nil")
	}

	// Verify default options
	opts := wlb.Options()
	if !opts.BreakAfterSlash {
		t.Error(
			"Default options should have BreakAfterSlash = true",
		)
	}
}

func TestNewWordLineBreakerWithOptions(
	t *testing.T,
) {
	opts := WordBreakOptions{
		BreakAfterSlash:    false,
		KeepEmailsTogether: true,
	}

	wlb := NewWordLineBreakerWithOptions(opts)
	if wlb == nil {
		t.Fatal(
			"NewWordLineBreakerWithOptions returned nil",
		)
	}

	resultOpts := wlb.Options()
	if resultOpts.BreakAfterSlash {
		t.Error("BreakAfterSlash should be false")
	}
	if !resultOpts.KeepEmailsTogether {
		t.Error(
			"KeepEmailsTogether should be true",
		)
	}
}

func TestWordLineBreaker_SetOptions(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	newOpts := WordBreakOptions{
		BreakAfterSlash: false,
	}
	wlb.SetOptions(newOpts)

	if wlb.Options().BreakAfterSlash {
		t.Error(
			"BreakAfterSlash should be false after SetOptions",
		)
	}
}

func TestWordLineBreaker_BreakAfterSlash(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	// Test breaking in paths
	text := "path/to/file"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Should have break opportunities after slashes
	hasBreakAfterFirstSlash := false
	hasBreakAfterSecondSlash := false

	for _, opp := range opportunities {
		if opp.Type == BreakAllowed {
			if opp.Position == 5 { // After "path/"
				hasBreakAfterFirstSlash = true
			}
			if opp.Position == 8 { // After "path/to/"
				hasBreakAfterSecondSlash = true
			}
		}
	}

	if !hasBreakAfterFirstSlash {
		t.Error(
			"Expected break opportunity after first slash in 'path/to/file'",
		)
	}
	if !hasBreakAfterSecondSlash {
		t.Error(
			"Expected break opportunity after second slash in 'path/to/file'",
		)
	}
}

func TestWordLineBreaker_BreakAfterBackslash(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	// Test breaking in Windows paths
	text := "C:\\Users\\Documents"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Should have break opportunities after backslashes
	hasBreakAfterBackslash := false

	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position > 0 &&
			opp.Position < len(text) {
			// Check if previous char was backslash
			if text[opp.Position-1] == '\\' {
				hasBreakAfterBackslash = true

				break
			}
		}
	}

	if !hasBreakAfterBackslash {
		t.Error(
			"Expected break opportunity after backslash in Windows path",
		)
	}
}

func TestWordLineBreaker_BreakAfterEquals(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	text := "variable=value"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	hasBreakAfterEquals := false
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position == 9 { // After "variable="
			hasBreakAfterEquals = true

			break
		}
	}

	if !hasBreakAfterEquals {
		t.Error(
			"Expected break opportunity after equals sign",
		)
	}
}

func TestWordLineBreaker_BreakAfterColon(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	// Test colon NOT in time notation
	text := "label:value"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	hasBreakAfterColon := false
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position == 6 { // After "label:"
			hasBreakAfterColon = true

			break
		}
	}

	if !hasBreakAfterColon {
		t.Error(
			"Expected break opportunity after colon in 'label:value'",
		)
	}
}

func TestWordLineBreaker_NoBreakInTimeNotation(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	// Test colon in time notation - should NOT break
	text := "12:30"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position == 3 { // After "12:"
			t.Error(
				"Should not break in time notation '12:30'",
			)
		}
	}
}

func TestWordLineBreaker_KeepEmailsTogether(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	text := "contact user@example.com for info"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// There should be no breaks inside the email address
	emailStart := 8 // "user"
	emailEnd := 24  // after ".com"

	for _, opp := range opportunities {
		if opp.Position > emailStart &&
			opp.Position < emailEnd {
			if opp.Type == BreakAllowed {
				t.Errorf(
					"Should not break inside email address at position %d",
					opp.Position,
				)
			}
		}
	}
}

func TestWordLineBreaker_KeepURLsTogether(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	text := "visit https://example.com/page today"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// There should be limited breaks inside the URL
	urlStart := 6 // "https"
	urlEnd := 30  // after "/page"

	breakCount := 0
	for _, opp := range opportunities {
		if opp.Position > urlStart &&
			opp.Position < urlEnd {
			if opp.Type == BreakAllowed {
				breakCount++
			}
		}
	}

	// We may allow one or two breaks after / in path, but not many
	if breakCount > 2 {
		t.Errorf(
			"Too many breaks inside URL: %d",
			breakCount,
		)
	}
}

func TestWordLineBreaker_KeepUnitsWithNumbers(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	tests := []struct {
		name      string
		text      string
		unitStart int
		unitEnd   int
	}{
		{
			"centimeters",
			"The width is 10 cm wide",
			13,
			18,
		},
		{
			"kilograms",
			"Weight: 5 kg total",
			8,
			12,
		},
		{"megabytes", "Size: 100 MB free", 6, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := wlb.FindBreakOpportunities(
				tt.text,
			)

			// Check there's no break between number and unit
			for _, opp := range opportunities {
				if opp.Position > tt.unitStart &&
					opp.Position < tt.unitEnd {
					if opp.Type == BreakAllowed {
						t.Errorf(
							"Should not break between number and unit at position %d in %q",
							opp.Position,
							tt.text,
						)
					}
				}
			}
		})
	}
}

func TestWordLineBreaker_KeepNumberAbbreviations(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	text := "See No. 123 for details"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Find the "No." position
	noStart := 4 // "No."
	noEnd := 11  // After "123"

	// There should be no break between "No." and the number
	for _, opp := range opportunities {
		if opp.Position > noStart &&
			opp.Position < noEnd {
			if opp.Type == BreakAllowed {
				t.Errorf(
					"Should not break between 'No.' and number at position %d",
					opp.Position,
				)
			}
		}
	}
}

func TestWordLineBreaker_SoftHyphen(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	// Soft hyphen is U+00AD (2 bytes in UTF-8)
	text := "interna\u00ADtional"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Should have a break opportunity after the soft hyphen
	// "interna" = 7 bytes, soft hyphen = 2 bytes, so position should be 9
	hasSoftHyphenBreak := false
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed {
			// Check if this is after the soft hyphen (rune position 8)
			if opp.RunePosition == 8 {
				hasSoftHyphenBreak = true

				break
			}
		}
	}

	if !hasSoftHyphenBreak {
		t.Error(
			"Expected break opportunity at soft hyphen",
		)
	}
}

func TestWordLineBreaker_EmDash(t *testing.T) {
	wlb := NewWordLineBreaker()

	// Em dash is U+2014
	text := "word\u2014another"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Should have break opportunities around em-dash
	hasBreakAroundEmDash := false
	for _, opp := range opportunities {
		// Check for break after the 'd' in "word" (before em-dash) or after em-dash
		if opp.Type == BreakAllowed {
			if opp.Position == 4 ||
				opp.Position == 7 { // Before or after em-dash (em-dash is 3 bytes)
				hasBreakAroundEmDash = true

				break
			}
		}
	}

	if !hasBreakAroundEmDash {
		t.Error(
			"Expected break opportunity around em-dash",
		)
	}
}

func TestWordLineBreaker_CompoundWords(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	text := "self-aware"
	opportunities := wlb.FindBreakOpportunities(
		text,
	)

	// Should have break opportunity after hyphen in compound word
	hasBreakAfterHyphen := false
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position == 5 { // After "self-"
			hasBreakAfterHyphen = true

			break
		}
	}

	if !hasBreakAfterHyphen {
		t.Error(
			"Expected break opportunity after hyphen in compound word",
		)
	}
}

func TestSplitIntoLinesWord(t *testing.T) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	tests := []struct {
		name      string
		text      string
		maxWidth  float64
		wantLines int
	}{
		{"simple words", "hello world", 20, 1},
		{"split words", "hello world", 6, 2},
		{
			"path with slash",
			"very/long/path/here",
			10,
			2,
		},
		{"empty string", "", 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := SplitIntoLinesWord(
				tt.text,
				tt.maxWidth,
				measureFunc,
			)
			if len(results) != tt.wantLines {
				t.Errorf(
					"SplitIntoLinesWord(%q, %v) = %d lines, want %d",
					tt.text,
					tt.maxWidth,
					len(results),
					tt.wantLines,
				)
				for i, r := range results {
					t.Logf(
						"Line %d: %q (width: %v)",
						i,
						r.Text,
						r.Width,
					)
				}
			}
		})
	}
}

func TestSplitIntoLinesWordWithOptions(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	opts := WordBreakOptions{
		BreakAfterSlash: true,
	}

	text := "path/to/file"
	results := SplitIntoLinesWordWithOptions(
		text,
		6,
		measureFunc,
		opts,
	)

	if len(results) < 2 {
		t.Errorf(
			"Expected at least 2 lines for path with slashes, got %d",
			len(results),
		)
	}
}

func TestIsCurrencySymbol(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{'$', true},
		{'\u00A3', true}, // Pound
		{'\u00A5', true}, // Yen
		{'\u20AC', true}, // Euro
		{'A', false},
		{'1', false},
		{' ', false},
	}

	for _, tt := range tests {
		if got := IsCurrencySymbol(tt.r); got != tt.want {
			t.Errorf(
				"IsCurrencySymbol(%q) = %v, want %v",
				tt.r,
				got,
				tt.want,
			)
		}
	}
}

func TestIsUnitAbbreviation(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"cm", true},
		{"kg", true},
		{"MB", true},
		{"mph", true},
		{"%", true},
		{"foo", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := IsUnitAbbreviation(tt.s); got != tt.want {
			t.Errorf(
				"IsUnitAbbreviation(%q) = %v, want %v",
				tt.s,
				got,
				tt.want,
			)
		}
	}
}

func TestIsTitleAbbreviation(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"Dr.", true},
		{"Mr.", true},
		{"Mrs.", true},
		{"Prof.", true},
		{"Gen.", true},
		{"Doctor", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := IsTitleAbbreviation(tt.s); got != tt.want {
			t.Errorf(
				"IsTitleAbbreviation(%q) = %v, want %v",
				tt.s,
				got,
				tt.want,
			)
		}
	}
}

func TestIsTimeNotation(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		colonPos int
		want     bool
	}{
		{"time format", "12:30", 2, true},
		{"hour minute", "9:45", 1, true},
		{"not time start", ":30", 0, false},
		{"not time end", "12:", 2, false},
		{"label colon", "label:value", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runes := []rune(tt.text)
			if got := isTimeNotation(runes, tt.colonPos); got != tt.want {
				t.Errorf(
					"isTimeNotation(%q, %d) = %v, want %v",
					tt.text,
					tt.colonPos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestWordLineBreaker_EmptyText(t *testing.T) {
	wlb := NewWordLineBreaker()

	opportunities := wlb.FindBreakOpportunities(
		"",
	)
	if len(opportunities) != 0 {
		t.Errorf(
			"Expected no opportunities for empty text, got %d",
			len(opportunities),
		)
	}
}

func TestWordLineBreaker_SingleCharacter(
	t *testing.T,
) {
	wlb := NewWordLineBreaker()

	opportunities := wlb.FindBreakOpportunities(
		"a",
	)
	// Single character should have no internal break opportunities
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed &&
			opp.Position > 0 &&
			opp.Position < 1 {
			t.Error(
				"Unexpected break opportunity in single character",
			)
		}
	}
}

// Benchmark tests

func BenchmarkWordLineBreaker_FindBreakOpportunities_Simple(
	b *testing.B,
) {
	wlb := NewWordLineBreaker()
	text := "Hello, world! This is a test."

	b.ResetTimer()
	for range b.N {
		wlb.FindBreakOpportunities(text)
	}
}

func BenchmarkWordLineBreaker_FindBreakOpportunities_WithEmail(
	b *testing.B,
) {
	wlb := NewWordLineBreaker()
	text := "Contact us at support@example.com for assistance."

	b.ResetTimer()
	for range b.N {
		wlb.FindBreakOpportunities(text)
	}
}

func BenchmarkWordLineBreaker_FindBreakOpportunities_WithURL(
	b *testing.B,
) {
	wlb := NewWordLineBreaker()
	text := "Visit https://example.com/path/to/page for more information."

	b.ResetTimer()
	for range b.N {
		wlb.FindBreakOpportunities(text)
	}
}

func BenchmarkWordLineBreaker_FindBreakOpportunities_Complex(
	b *testing.B,
) {
	wlb := NewWordLineBreaker()
	text := "Dr. Smith sent 100 MB to user@example.com at 12:30. See No. 42 and visit https://example.com/path for C:\\Windows\\path."

	b.ResetTimer()
	for range b.N {
		wlb.FindBreakOpportunities(text)
	}
}

func BenchmarkSplitIntoLinesWord(b *testing.B) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}
	text := "The quick brown fox jumps over the lazy dog. Visit https://example.com for more."
	maxWidth := 40.0

	b.ResetTimer()
	for range b.N {
		SplitIntoLinesWord(
			text,
			maxWidth,
			measureFunc,
		)
	}
}
