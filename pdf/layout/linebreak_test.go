package layout

import (
	"strings"
	"testing"
)

func TestLineBreakClass_String(t *testing.T) {
	tests := []struct {
		class LineBreakClass
		want  string
	}{
		{BK, "BK"},
		{CR, "CR"},
		{LF, "LF"},
		{SP, "SP"},
		{AL, "AL"},
		{NU, "NU"},
		{ID, "ID"},
		{XX, "XX"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.class.String(); got != tt.want {
				t.Errorf(
					"LineBreakClass.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestBreakOpportunityType_String(
	t *testing.T,
) {
	tests := []struct {
		typ  BreakOpportunityType
		want string
	}{
		{BreakProhibited, "Prohibited"},
		{BreakAllowed, "Allowed"},
		{BreakMandatory, "Mandatory"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.typ.String(); got != tt.want {
				t.Errorf(
					"BreakOpportunityType.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGetLineBreakClass(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want LineBreakClass
	}{
		// Mandatory breaks
		{"line feed", '\n', LF},
		{"carriage return", '\r', CR},
		{"next line", '\u0085', NL},
		{"line separator", '\u2028', BK},
		{"paragraph separator", '\u2029', BK},

		// Spaces
		{"space", ' ', SP},
		{"no-break space", '\u00A0', GL},
		{"em space", '\u2003', SP},
		{"figure space", '\u2007', GL},

		// Zero-width
		{"zero width space", '\u200B', ZW},
		{"word joiner", '\u2060', WJ},
		{"soft hyphen", '\u00AD', BA},

		// Punctuation
		{"opening paren", '(', OP},
		{"closing paren", ')', CL},
		{"opening bracket", '[', OP},
		{"closing bracket", ']', CL},
		{"exclamation", '!', EX},
		{"question", '?', EX},

		// Quotes
		{"double quote", '"', QU},
		{"single quote", '\'', QU},
		{"left double quote", '\u201C', QU},
		{"right double quote", '\u201D', QU},

		// Hyphens
		{"hyphen-minus", '-', HY},
		{"hyphen", '\u2010', HY},
		{"em dash", '\u2014', BA},

		// Numeric
		{"digit", '0', NU},
		{"dollar", '$', PR},
		{"percent", '%', PO},
		{"period", '.', IS},
		{"comma", ',', IS},

		// Letters
		{"latin letter", 'A', AL},
		{"lowercase letter", 'a', AL},

		// CJK
		{"CJK ideograph", '\u4E00', ID},
		{"hiragana", '\u3042', ID},
		{"katakana", '\u30A2', ID},
		{"fullwidth comma", '\uFF0C', NS},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLineBreakClass(tt.r); got != tt.want {
				t.Errorf(
					"GetLineBreakClass(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestLineBreaker_FindBreakOpportunities_MandatoryBreaks(
	t *testing.T,
) {
	lb := NewLineBreaker()

	tests := []struct {
		name          string
		text          string
		wantMandatory int // number of mandatory breaks
	}{
		{"single line feed", "hello\nworld", 1},
		{"crlf", "hello\r\nworld", 1},
		{"multiple lf", "a\nb\nc", 2},
		{"trailing lf", "hello\n", 1},
		{"only lf", "\n", 1},
		{
			"paragraph separator",
			"hello\u2029world",
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := lb.FindBreakOpportunities(
				tt.text,
			)

			mandatoryCount := 0
			for _, opp := range opportunities {
				if opp.Type == BreakMandatory {
					mandatoryCount++
				}
			}

			if mandatoryCount != tt.wantMandatory {
				t.Errorf(
					"FindBreakOpportunities(%q) found %d mandatory breaks, want %d",
					tt.text,
					mandatoryCount,
					tt.wantMandatory,
				)
			}
		})
	}
}

func TestLineBreaker_FindBreakOpportunities_Spaces(
	t *testing.T,
) {
	lb := NewLineBreaker()

	tests := []struct {
		name      string
		text      string
		wantBreak bool // expect at least one break opportunity
	}{
		{"simple words", "hello world", true},
		{"multiple spaces", "hello  world", true},
		{"no spaces", "helloworld", false},
		{"only space", " ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := lb.FindBreakOpportunities(
				tt.text,
			)

			hasBreak := false
			for _, opp := range opportunities {
				if opp.Type == BreakAllowed ||
					opp.Type == BreakMandatory {
					hasBreak = true

					break
				}
			}

			if hasBreak != tt.wantBreak {
				t.Errorf(
					"FindBreakOpportunities(%q) hasBreak = %v, want %v",
					tt.text,
					hasBreak,
					tt.wantBreak,
				)
			}
		})
	}
}

func TestLineBreaker_FindBreakOpportunities_Punctuation(
	t *testing.T,
) {
	lb := NewLineBreaker()

	// Test that we don't break before closing punctuation
	opportunities := lb.FindBreakOpportunities(
		"word)",
	)

	for _, opp := range opportunities {
		// Position 4 is before ')'
		if opp.Position == 4 &&
			opp.Type != BreakProhibited {
			t.Errorf(
				"Should not break before closing parenthesis",
			)
		}
	}

	// Test that we don't break after opening punctuation
	opportunities = lb.FindBreakOpportunities(
		"(word",
	)

	for _, opp := range opportunities {
		// Position 1 is after '('
		if opp.Position == 1 &&
			opp.Type != BreakProhibited {
			t.Errorf(
				"Should not break after opening parenthesis",
			)
		}
	}
}

func TestLineBreaker_FindBreakOpportunities_CJK(
	t *testing.T,
) {
	lb := NewLineBreaker()

	// CJK text should have break opportunities between characters
	text := "\u4E00\u4E8C\u4E09" // One, Two, Three in Chinese
	opportunities := lb.FindBreakOpportunities(
		text,
	)

	hasBreak := false
	for _, opp := range opportunities {
		if opp.Type == BreakAllowed {
			hasBreak = true

			break
		}
	}

	if !hasBreak {
		t.Error(
			"Expected break opportunities between CJK characters",
		)
	}
}

func TestLineBreaker_FindBreakOpportunities_Numbers(
	t *testing.T,
) {
	lb := NewLineBreaker()

	tests := []struct {
		name           string
		text           string
		expectNoBreaks bool // true if we expect no breaks within the number
	}{
		{"integer", "12345", true},
		{"decimal", "123.45", true},
		{"currency", "$100", true},
		{"percentage", "50%", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := lb.FindBreakOpportunities(
				tt.text,
			)

			hasInternalBreak := false
			for _, opp := range opportunities {
				if opp.Position > 0 &&
					opp.Position < len(tt.text) &&
					opp.Type == BreakAllowed {
					hasInternalBreak = true

					break
				}
			}

			if tt.expectNoBreaks &&
				hasInternalBreak {
				t.Errorf(
					"Unexpected break in %q",
					tt.text,
				)
			}
		})
	}
}

func TestLineBreaker_CanBreakBefore(
	t *testing.T,
) {
	lb := NewLineBreaker()

	tests := []struct {
		name    string
		text    string
		runePos int
		want    bool
	}{
		{
			"before second word",
			"hello world",
			6,
			true,
		},
		{"start of string", "hello", 0, false},
		{
			"before closing paren",
			"word)",
			4,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lb.CanBreakBefore(
				tt.text,
				tt.runePos,
			)
			if got != tt.want {
				t.Errorf(
					"CanBreakBefore(%q, %d) = %v, want %v",
					tt.text,
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestLineBreaker_CanBreakAfter(t *testing.T) {
	lb := NewLineBreaker()

	tests := []struct {
		name    string
		text    string
		runePos int
		want    bool
	}{
		{"after space", "hello world", 5, true},
		{"after letter", "hello world", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lb.CanBreakAfter(
				tt.text,
				tt.runePos,
			)
			if got != tt.want {
				t.Errorf(
					"CanBreakAfter(%q, %d) = %v, want %v",
					tt.text,
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestFindFirstBreak(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			"simple words",
			"hello world",
			6,
		}, // After "hello "
		{
			"newline",
			"hello\nworld",
			6,
		}, // After "hello\n"
		{"no break", "helloworld", -1},
		{"empty", "", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindFirstBreak(tt.text)
			if got != tt.want {
				t.Errorf(
					"FindFirstBreak(%q) = %v, want %v",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestSplitIntoLines(t *testing.T) {
	// Simple measure function: 1 unit per character
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	tests := []struct {
		name      string
		text      string
		maxWidth  float64
		wantLines int
	}{
		{"single word fits", "hello", 10, 1},
		{
			"two words separate",
			"hello world",
			6,
			2,
		},
		{
			"multiple lines",
			"one two three four",
			10,
			2,
		},
		{"empty string", "", 10, 0},
		{"forced break", "hello\nworld", 100, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := SplitIntoLines(
				tt.text,
				tt.maxWidth,
				measureFunc,
			)
			if len(results) != tt.wantLines {
				t.Errorf(
					"SplitIntoLines(%q, %v) = %d lines, want %d",
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

func TestSplitIntoLines_WidthRespected(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	text := "The quick brown fox jumps over the lazy dog"
	maxWidth := 15.0

	results := SplitIntoLines(
		text,
		maxWidth,
		measureFunc,
	)

	for i, result := range results {
		if result.Width > maxWidth {
			t.Errorf(
				"Line %d exceeds max width: %q (width: %v, max: %v)",
				i,
				result.Text,
				result.Width,
				maxWidth,
			)
		}
	}
}

func TestSplitIntoLines_Reconstructs(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	text := "Hello world this is a test"
	maxWidth := 12.0

	results := SplitIntoLines(
		text,
		maxWidth,
		measureFunc,
	)

	// Reconstruct the text (accounting for spaces between lines)
	var reconstructed strings.Builder
	for _, result := range results {
		reconstructed.WriteString(
			strings.TrimSpace(result.Text),
		)
		if !result.ForcedBreak {
			reconstructed.WriteString(" ")
		}
	}

	// Trim trailing space
	got := strings.TrimSpace(
		reconstructed.String(),
	)
	want := strings.Join(
		strings.Fields(text),
		" ",
	)

	if got != want {
		t.Errorf(
			"Reconstructed text = %q, want %q",
			got,
			want,
		)
	}
}

func TestBreakText(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		wantBefore string
		wantAfter  string
	}{
		{
			"simple words",
			"hello world",
			"hello ",
			"world",
		},
		{
			"no break",
			"helloworld",
			"helloworld",
			"",
		},
		{
			"newline",
			"hello\nworld",
			"hello\n",
			"world",
		},
		{"empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before, after := BreakText(tt.text)
			if before != tt.wantBefore ||
				after != tt.wantAfter {
				t.Errorf(
					"BreakText(%q) = (%q, %q), want (%q, %q)",
					tt.text,
					before,
					after,
					tt.wantBefore,
					tt.wantAfter,
				)
			}
		})
	}
}

func TestNewLineBreaker(t *testing.T) {
	lb := NewLineBreaker()
	if lb == nil {
		t.Fatal("NewLineBreaker() returned nil")
	}
	if lb.strictMode {
		t.Error(
			"NewLineBreaker() should not be in strict mode",
		)
	}
}

func TestNewStrictLineBreaker(t *testing.T) {
	lb := NewStrictLineBreaker()
	if lb == nil {
		t.Fatal(
			"NewStrictLineBreaker() returned nil",
		)
	}
	if !lb.strictMode {
		t.Error(
			"NewStrictLineBreaker() should be in strict mode",
		)
	}
}

func TestIsCJKIdeograph(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"basic CJK", '\u4E00', true},
		{"end of CJK", '\u9FFF', true},
		{"extension A", '\u3400', true},
		{"latin letter", 'A', false},
		{"digit", '0', false},
		{
			"hiragana",
			'\u3042',
			false,
		}, // Hiragana is not ideographic
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCJKIdeograph(tt.r); got != tt.want {
				t.Errorf(
					"isCJKIdeograph(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsCJKNonStarter(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"small hiragana a", '\u3041', true},
		{"prolonged sound mark", '\u30FC', true},
		{"ideographic comma", '\u3001', true},
		{"regular hiragana", '\u3042', false},
		{"regular katakana", '\u30A2', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCJKNonStarter(tt.r); got != tt.want {
				t.Errorf(
					"isCJKNonStarter(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

// Benchmark tests
func BenchmarkFindBreakOpportunities_Short(
	b *testing.B,
) {
	lb := NewLineBreaker()
	text := "Hello, world!"

	b.ResetTimer()
	for range b.N {
		lb.FindBreakOpportunities(text)
	}
}

func BenchmarkFindBreakOpportunities_Long(
	b *testing.B,
) {
	lb := NewLineBreaker()
	text := "The quick brown fox jumps over the lazy dog. " +
		"Pack my box with five dozen liquor jugs. " +
		"How vexingly quick daft zebras jump!"

	b.ResetTimer()
	for range b.N {
		lb.FindBreakOpportunities(text)
	}
}

func BenchmarkFindBreakOpportunities_CJK(
	b *testing.B,
) {
	lb := NewLineBreaker()
	// Some Japanese text
	text := "\u3053\u3093\u306B\u3061\u306F\u4E16\u754C"

	b.ResetTimer()
	for range b.N {
		lb.FindBreakOpportunities(text)
	}
}

func BenchmarkSplitIntoLines(b *testing.B) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}
	text := "The quick brown fox jumps over the lazy dog. " +
		"Pack my box with five dozen liquor jugs."
	maxWidth := 40.0

	b.ResetTimer()
	for range b.N {
		SplitIntoLines(
			text,
			maxWidth,
			measureFunc,
		)
	}
}

func BenchmarkGetLineBreakClass(b *testing.B) {
	runes := []rune(
		"Hello, world! \u4E00\u4E8C\u4E09",
	)

	b.ResetTimer()
	for range b.N {
		for _, r := range runes {
			GetLineBreakClass(r)
		}
	}
}
