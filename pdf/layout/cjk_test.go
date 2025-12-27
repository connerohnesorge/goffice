package layout

import (
	"testing"
)

func TestCJKScriptType_String(t *testing.T) {
	tests := []struct {
		scriptType CJKScriptType
		want       string
	}{
		{CJKScriptNone, "None"},
		{CJKScriptHan, "Han"},
		{CJKScriptHiragana, "Hiragana"},
		{CJKScriptKatakana, "Katakana"},
		{CJKScriptHangul, "Hangul"},
		{CJKScriptBopomofo, "Bopomofo"},
		{CJKScriptSymbol, "Symbol"},
		{CJKScriptType(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.scriptType.String(); got != tt.want {
				t.Errorf(
					"CJKScriptType.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCJKBreakRule_String(t *testing.T) {
	tests := []struct {
		rule CJKBreakRule
		want string
	}{
		{CJKBreakNormal, "Normal"},
		{CJKBreakNoBreakBefore, "NoBreakBefore"},
		{CJKBreakNoBreakAfter, "NoBreakAfter"},
		{CJKBreakNoBreakEither, "NoBreakEither"},
		{CJKBreakRule(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.rule.String(); got != tt.want {
				t.Errorf(
					"CJKBreakRule.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsCJK(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		// Han characters
		{"Han basic", '\u4E00', true},
		{"Han end of range", '\u9FFF', true},
		{"Han extension A", '\u3400', true},
		{"Han compatibility", '\uF900', true},

		// Hiragana
		{"Hiragana A", '\u3042', true},
		{"Hiragana small A", '\u3041', true},
		{"Hiragana end", '\u3096', true},

		// Katakana
		{"Katakana A", '\u30A2', true},
		{"Katakana small A", '\u30A1', true},
		{"Katakana prolonged", '\u30FC', true},
		{"Halfwidth Katakana", '\uFF66', true},

		// Hangul
		{"Hangul syllable", '\uAC00', true},
		{"Hangul jamo", '\u1100', true},
		{"Hangul compatibility", '\u3130', true},

		// Bopomofo
		{"Bopomofo", '\u3100', true},

		// CJK symbols
		{"CJK symbol", '\u3000', true},
		{"Ideographic comma", '\u3001', true},
		{"Fullwidth comma", '\uFF0C', true},

		// Non-CJK
		{"Latin letter", 'A', false},
		{"Digit", '0', false},
		{"Space", ' ', false},
		{"Greek", '\u03B1', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCJK(tt.r); got != tt.want {
				t.Errorf(
					"IsCJK(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGetCJKScriptType(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want CJKScriptType
	}{
		// Han
		{"Han ideograph", '\u4E00', CJKScriptHan},
		{"Han middle", '\u5B66', CJKScriptHan},

		// Hiragana
		{
			"Hiragana A",
			'\u3042',
			CJKScriptHiragana,
		},
		{
			"Hiragana N",
			'\u3093',
			CJKScriptHiragana,
		},

		// Katakana
		{
			"Katakana A",
			'\u30A2',
			CJKScriptKatakana,
		},
		{
			"Katakana prolonged",
			'\u30FC',
			CJKScriptKatakana,
		},

		// Hangul
		{"Hangul", '\uAC00', CJKScriptHangul},

		// Bopomofo
		{"Bopomofo", '\u3105', CJKScriptBopomofo},

		// Symbols
		{"CJK period", '\u3002', CJKScriptSymbol},
		{
			"Fullwidth exclaim",
			'\uFF01',
			CJKScriptSymbol,
		},

		// Non-CJK
		{"Latin", 'a', CJKScriptNone},
		{"Space", ' ', CJKScriptNone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCJKScriptType(tt.r); got != tt.want {
				t.Errorf(
					"GetCJKScriptType(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGetCJKBreakRule(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want CJKBreakRule
	}{
		// No break before (kinsoku shori)
		{
			"Ideographic comma",
			'\u3001',
			CJKBreakNoBreakBefore,
		},
		{
			"Ideographic period",
			'\u3002',
			CJKBreakNoBreakBefore,
		},
		{
			"Small hiragana a",
			'\u3041',
			CJKBreakNoBreakBefore,
		},
		{
			"Small katakana a",
			'\u30A1',
			CJKBreakNoBreakBefore,
		},
		{
			"Prolonged sound mark",
			'\u30FC',
			CJKBreakNoBreakBefore,
		},
		{
			"Iteration mark",
			'\u3005',
			CJKBreakNoBreakBefore,
		},
		{
			"Fullwidth comma",
			'\uFF0C',
			CJKBreakNoBreakBefore,
		},
		{
			"Fullwidth question",
			'\uFF1F',
			CJKBreakNoBreakBefore,
		},
		{
			"Right corner bracket",
			'\u300D',
			CJKBreakNoBreakBefore,
		},

		// No break after
		{
			"Left corner bracket",
			'\u300C',
			CJKBreakNoBreakAfter,
		},
		{
			"Left angle bracket",
			'\u3008',
			CJKBreakNoBreakAfter,
		},
		{
			"Fullwidth left paren",
			'\uFF08',
			CJKBreakNoBreakAfter,
		},
		{
			"Fullwidth yen",
			'\uFFE5',
			CJKBreakNoBreakAfter,
		},

		// Normal
		{"Regular han", '\u4E00', CJKBreakNormal},
		{
			"Regular hiragana",
			'\u3042',
			CJKBreakNormal,
		},
		{
			"Regular katakana",
			'\u30A2',
			CJKBreakNormal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCJKBreakRule(tt.r); got != tt.want {
				t.Errorf(
					"GetCJKBreakRule(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsNoBreakBefore(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Small hiragana", '\u3041', true},
		{"Prolonged sound", '\u30FC', true},
		{"Comma", '\u3001', true},
		{"Period", '\u3002', true},
		{"Regular hiragana", '\u3042', false},
		{"Han character", '\u4E00', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNoBreakBefore(tt.r); got != tt.want {
				t.Errorf(
					"IsNoBreakBefore(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsNoBreakAfter(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Left bracket", '\u300C', true},
		{"Fullwidth left paren", '\uFF08', true},
		{"Right bracket", '\u300D', false},
		{"Han character", '\u4E00', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNoBreakAfter(tt.r); got != tt.want {
				t.Errorf(
					"IsNoBreakAfter(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsSmallKana(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		// Small Hiragana
		{"Small hiragana a", '\u3041', true},
		{"Small hiragana i", '\u3043', true},
		{"Small hiragana tsu", '\u3063', true},

		// Small Katakana
		{"Small katakana a", '\u30A1', true},
		{"Small katakana tsu", '\u30C3', true},

		// Halfwidth small
		{"Halfwidth small a", '\uFF67', true},
		{"Halfwidth small tsu", '\uFF6F', true},

		// Regular kana
		{"Regular hiragana a", '\u3042', false},
		{"Regular katakana a", '\u30A2', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSmallKana(tt.r); got != tt.want {
				t.Errorf(
					"IsSmallKana(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsProlongedSoundMark(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Prolonged sound mark", '\u30FC', true},
		{"Halfwidth prolonged", '\uFF70', true},
		{"Regular katakana", '\u30A2', false},
		{"Han character", '\u4E00', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProlongedSoundMark(tt.r); got != tt.want {
				t.Errorf(
					"IsProlongedSoundMark(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsIterationMark(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Ideographic iteration", '\u3005', true},
		{"Vertical iteration", '\u303B', true},
		{"Hiragana iteration", '\u309D', true},
		{"Katakana iteration", '\u30FD', true},
		{"Regular hiragana", '\u3042', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIterationMark(tt.r); got != tt.want {
				t.Errorf(
					"IsIterationMark(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCJKBreakChecker_CanBreakBetween(
	t *testing.T,
) {
	checker := NewCJKBreakChecker()

	tests := []struct {
		name   string
		before rune
		after  rune
		want   bool
	}{
		// Normal CJK breaks
		{"Han to Han", '\u4E00', '\u4E8C', true},
		{
			"Hiragana to Hiragana",
			'\u3042',
			'\u3044',
			true,
		},
		{
			"Han to Hiragana",
			'\u4E00',
			'\u3042',
			true,
		},

		// No break before small kana
		{
			"Before small hiragana",
			'\u3042',
			'\u3041',
			false,
		},
		{
			"Before small katakana",
			'\u30A2',
			'\u30A1',
			false,
		},
		{
			"Before prolonged sound",
			'\u30A2',
			'\u30FC',
			false,
		},

		// No break after opening bracket
		{
			"After left bracket",
			'\u300C',
			'\u4E00',
			false,
		},
		{
			"After fullwidth paren",
			'\uFF08',
			'\u4E00',
			false,
		},

		// No break before closing bracket
		{
			"Before right bracket",
			'\u4E00',
			'\u300D',
			false,
		},
		{
			"Before comma",
			'\u4E00',
			'\u3001',
			false,
		},

		// Mixed with punctuation
		{
			"Before period",
			'\u4E00',
			'\u3002',
			false,
		},
		{
			"After comma normal",
			'\u3001',
			'\u4E00',
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker.CanBreakBetween(tt.before, tt.after); got != tt.want {
				t.Errorf(
					"CanBreakBetween(%q, %q) = %v, want %v",
					tt.before,
					tt.after,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCJKBreakChecker_CanBreakAt(
	t *testing.T,
) {
	checker := NewCJKBreakChecker()

	tests := []struct {
		name    string
		text    string
		runePos int
		want    bool
	}{
		// Japanese text
		{
			"Between Han",
			"\u4E00\u4E8C\u4E09",
			1,
			true,
		},
		{
			"Before small kana",
			"\u3042\u3041",
			1,
			false,
		},
		{
			"Before period",
			"\u3053\u3093\u306B\u3061\u306F\u3002",
			5,
			false,
		},

		// Boundary cases
		{"Position 0", "\u4E00\u4E8C", 0, false},
		{
			"Position at end",
			"\u4E00\u4E8C",
			2,
			false,
		},
		{"Empty string", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker.CanBreakAt(tt.text, tt.runePos); got != tt.want {
				t.Errorf(
					"CanBreakAt(%q, %d) = %v, want %v",
					tt.text,
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCJKBreakChecker_FindBreakOpportunities(
	t *testing.T,
) {
	checker := NewCJKBreakChecker()

	tests := []struct {
		name    string
		text    string
		wantLen int
		wantPos []int
	}{
		{
			name:    "Three Han characters",
			text:    "\u4E00\u4E8C\u4E09",
			wantLen: 2,
			wantPos: []int{1, 2},
		},
		{
			name:    "Han with small kana",
			text:    "\u3042\u3041", // A small-A
			wantLen: 0,              // No break before small kana
		},
		{
			name:    "Bracketed text",
			text:    "\u300C\u4E00\u300D", // "one" in brackets
			wantLen: 0,                    // No breaks inside brackets
		},
		{
			name:    "Short text",
			text:    "\u4E00",
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := checker.FindBreakOpportunities(
				tt.text,
			)

			if len(opportunities) != tt.wantLen {
				t.Errorf(
					"FindBreakOpportunities(%q) returned %d opportunities, want %d",
					tt.text,
					len(opportunities),
					tt.wantLen,
				)

				return
			}

			if tt.wantPos != nil {
				for i, pos := range tt.wantPos {
					if i < len(opportunities) &&
						opportunities[i] != pos {
						t.Errorf(
							"opportunity[%d] = %d, want %d",
							i,
							opportunities[i],
							pos,
						)
					}
				}
			}
		})
	}
}

func TestCJKBreakCheckerStrictness(t *testing.T) {
	lenient := NewCJKBreakCheckerWithStrictness(0)
	normal := NewCJKBreakCheckerWithStrictness(1)
	strict := NewCJKBreakCheckerWithStrictness(2)

	// Test that strict mode prevents more breaks
	tests := []struct {
		name    string
		before  rune
		after   rune
		lenient bool
		normal  bool
		strict  bool
	}{
		{
			name:    "Normal Han characters",
			before:  '\u4E00',
			after:   '\u4E8C',
			lenient: true,
			normal:  true,
			strict:  true,
		},
		{
			name:    "Before iteration mark",
			before:  '\u4E00',
			after:   '\u3005',
			lenient: false,
			normal:  false,
			strict:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenient.CanBreakBetween(tt.before, tt.after); got != tt.lenient {
				t.Errorf(
					"lenient.CanBreakBetween = %v, want %v",
					got,
					tt.lenient,
				)
			}
			if got := normal.CanBreakBetween(tt.before, tt.after); got != tt.normal {
				t.Errorf(
					"normal.CanBreakBetween = %v, want %v",
					got,
					tt.normal,
				)
			}
			if got := strict.CanBreakBetween(tt.before, tt.after); got != tt.strict {
				t.Errorf(
					"strict.CanBreakBetween = %v, want %v",
					got,
					tt.strict,
				)
			}
		})
	}
}

func TestAnalyzeCJKContent(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		wantHasCJK    bool
		wantPurelyCJK bool
		wantHan       int
		wantHiragana  int
		wantKatakana  int
		wantDominant  CJKScriptType
	}{
		{
			name:          "Chinese text",
			text:          "\u4E00\u4E8C\u4E09",
			wantHasCJK:    true,
			wantPurelyCJK: true,
			wantHan:       3,
			wantHiragana:  0,
			wantKatakana:  0,
			wantDominant:  CJKScriptHan,
		},
		{
			name:          "Japanese mixed",
			text:          "\u3053\u3093\u306B\u3061\u306F", // Hiragana "konnichiha"
			wantHasCJK:    true,
			wantPurelyCJK: true,
			wantHan:       0,
			wantHiragana:  5,
			wantKatakana:  0,
			wantDominant:  CJKScriptHiragana,
		},
		{
			name:          "Mixed with Latin",
			text:          "Hello\u4E16\u754C",
			wantHasCJK:    true,
			wantPurelyCJK: false,
			wantHan:       2,
			wantHiragana:  0,
			wantKatakana:  0,
			wantDominant:  CJKScriptHan,
		},
		{
			name:          "Latin only",
			text:          "Hello World",
			wantHasCJK:    false,
			wantPurelyCJK: false,
			wantHan:       0,
			wantHiragana:  0,
			wantKatakana:  0,
			wantDominant:  CJKScriptNone,
		},
		{
			name:          "Empty string",
			text:          "",
			wantHasCJK:    false,
			wantPurelyCJK: false,
			wantHan:       0,
			wantHiragana:  0,
			wantKatakana:  0,
			wantDominant:  CJKScriptNone,
		},
		{
			name:          "Katakana",
			text:          "\u30A2\u30A4\u30A6\u30A8\u30AA",
			wantHasCJK:    true,
			wantPurelyCJK: true,
			wantHan:       0,
			wantHiragana:  0,
			wantKatakana:  5,
			wantDominant:  CJKScriptKatakana,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := AnalyzeCJKContent(tt.text)

			if info.HasCJK != tt.wantHasCJK {
				t.Errorf(
					"HasCJK = %v, want %v",
					info.HasCJK,
					tt.wantHasCJK,
				)
			}
			if info.IsPurelyCJK != tt.wantPurelyCJK {
				t.Errorf(
					"IsPurelyCJK = %v, want %v",
					info.IsPurelyCJK,
					tt.wantPurelyCJK,
				)
			}
			if info.HanCount != tt.wantHan {
				t.Errorf(
					"HanCount = %d, want %d",
					info.HanCount,
					tt.wantHan,
				)
			}
			if info.HiraganaCount != tt.wantHiragana {
				t.Errorf(
					"HiraganaCount = %d, want %d",
					info.HiraganaCount,
					tt.wantHiragana,
				)
			}
			if info.KatakanaCount != tt.wantKatakana {
				t.Errorf(
					"KatakanaCount = %d, want %d",
					info.KatakanaCount,
					tt.wantKatakana,
				)
			}
			if info.DominantScript != tt.wantDominant {
				t.Errorf(
					"DominantScript = %v, want %v",
					info.DominantScript,
					tt.wantDominant,
				)
			}
		})
	}
}

func TestIsFullwidthPunctuation(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Fullwidth exclaim", '\uFF01', true},
		{"Fullwidth comma", '\uFF0C', true},
		{"Fullwidth period", '\uFF0E', true},
		{"Fullwidth colon", '\uFF1A', true},
		{"Ideographic comma", '\u3001', true},
		{"Ideographic period", '\u3002', true},
		{"ASCII period", '.', false},
		{"ASCII comma", ',', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFullwidthPunctuation(tt.r); got != tt.want {
				t.Errorf(
					"IsFullwidthPunctuation(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsFullwidthBracket(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Fullwidth left paren", '\uFF08', true},
		{"Fullwidth right paren", '\uFF09', true},
		{"Corner bracket left", '\u300C', true},
		{"Corner bracket right", '\u300D', true},
		{"Angle bracket left", '\u3008', true},
		{"ASCII paren", '(', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFullwidthBracket(tt.r); got != tt.want {
				t.Errorf(
					"IsFullwidthBracket(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCJKLineBreaker_FindBreakOpportunities(
	t *testing.T,
) {
	clb := NewCJKLineBreaker()

	tests := []struct {
		name           string
		text           string
		wantBreaks     int
		checkNoBreakAt []int // rune positions where no break should occur
	}{
		{
			name:       "Chinese sentence",
			text:       "\u4E00\u4E8C\u4E09\u56DB",
			wantBreaks: 3, // Between each character
		},
		{
			name:       "Japanese with period",
			text:       "\u3053\u3093\u306B\u3061\u306F\u3002",
			wantBreaks: 4, // Not before period
			checkNoBreakAt: []int{
				5,
			}, // Before period
		},
		{
			name:       "Quoted text",
			text:       "\u300C\u4E00\u300D",
			wantBreaks: 0, // No breaks inside quoted text
		},
		{
			name:       "With newline",
			text:       "\u4E00\n\u4E8C",
			wantBreaks: 1, // Mandatory break
		},
		{
			name:       "Empty",
			text:       "",
			wantBreaks: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opportunities := clb.FindBreakOpportunities(
				tt.text,
			)

			// Count non-prohibited breaks
			breakCount := 0
			for _, opp := range opportunities {
				if opp.Type != BreakProhibited {
					breakCount++
				}
			}

			if breakCount != tt.wantBreaks {
				t.Errorf(
					"FindBreakOpportunities(%q) found %d breaks, want %d",
					tt.text,
					breakCount,
					tt.wantBreaks,
				)
			}

			// Check positions where breaks should not occur
			for _, pos := range tt.checkNoBreakAt {
				for _, opp := range opportunities {
					if opp.RunePosition == pos &&
						opp.Type == BreakAllowed {
						t.Errorf(
							"Expected no break at position %d, but found one",
							pos,
						)
					}
				}
			}
		})
	}
}

func TestCJKLineBreaker_WithStrictness(
	t *testing.T,
) {
	strictBreaker := NewCJKLineBreakerWithStrictness(
		2,
	)

	// Test that strict mode is properly applied
	text := "\u4E00\u4E8C\u4E09"
	opportunities := strictBreaker.FindBreakOpportunities(
		text,
	)

	if len(opportunities) == 0 {
		t.Error(
			"Expected some break opportunities in strict mode",
		)
	}
}

// Japanese real-world text tests
func TestCJKBreaking_JapaneseText(t *testing.T) {
	checker := NewCJKBreakChecker()

	// Test kinsoku shori with real Japanese text patterns
	tests := []struct {
		name   string
		before rune
		after  rune
		want   bool
	}{
		// Small kana cannot start a line
		{
			"Before small tsu",
			'\u3063',
			'\u3066',
			true,
		}, // tsu followed by te
		{
			"te before small tsu",
			'\u3066',
			'\u3063',
			false,
		},

		// Prolonged sound mark
		{
			"Before prolonged mark",
			'\u30FC',
			'\u30A2',
			true,
		},
		{
			"A before prolonged",
			'\u30A2',
			'\u30FC',
			false,
		},

		// Iteration marks
		{
			"Before iteration",
			'\u3005',
			'\u4E00',
			true,
		},
		{
			"Han before iteration",
			'\u4E00',
			'\u3005',
			false,
		},

		// Punctuation
		{
			"Period cannot start",
			'\u3002',
			'\u3042',
			true,
		},
		{
			"Before period",
			'\u3042',
			'\u3002',
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checker.CanBreakBetween(
				tt.before,
				tt.after,
			)
			if got != tt.want {
				t.Errorf(
					"CanBreakBetween(%q, %q) = %v, want %v",
					tt.before,
					tt.after,
					got,
					tt.want,
				)
			}
		})
	}
}

// Korean text tests
func TestCJKBreaking_KoreanText(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Hangul syllable GA", '\uAC00', true},
		{"Hangul syllable NA", '\uB098', true},
		{"Hangul jamo", '\u1100', true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCJK(tt.r); got != tt.want {
				t.Errorf(
					"IsCJK(%q) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
			if got := GetCJKScriptType(tt.r); got != CJKScriptHangul {
				t.Errorf(
					"GetCJKScriptType(%q) = %v, want Hangul",
					tt.r,
					got,
				)
			}
		})
	}
}

// Benchmark tests
func BenchmarkIsCJK(b *testing.B) {
	runes := []rune(
		"\u4E00\u3042\u30A2\uAC00ABC123",
	)

	b.ResetTimer()
	for range b.N {
		for _, r := range runes {
			IsCJK(r)
		}
	}
}

func BenchmarkGetCJKScriptType(b *testing.B) {
	runes := []rune(
		"\u4E00\u3042\u30A2\uAC00ABC123",
	)

	b.ResetTimer()
	for range b.N {
		for _, r := range runes {
			GetCJKScriptType(r)
		}
	}
}

func BenchmarkCJKBreakChecker_CanBreakBetween(
	b *testing.B,
) {
	checker := NewCJKBreakChecker()
	pairs := [][2]rune{
		{'\u4E00', '\u4E8C'},
		{'\u3042', '\u3041'},
		{'\u300C', '\u4E00'},
		{'\u4E00', '\u300D'},
	}

	b.ResetTimer()
	for range b.N {
		for _, pair := range pairs {
			checker.CanBreakBetween(
				pair[0],
				pair[1],
			)
		}
	}
}

func BenchmarkCJKLineBreaker_FindBreakOpportunities(
	b *testing.B,
) {
	clb := NewCJKLineBreaker()
	text := "\u3053\u3093\u306B\u3061\u306F\u4E16\u754C\u3002\u4ECA\u65E5\u306F\u826F\u3044\u5929\u6C17\u3067\u3059\u3002"

	b.ResetTimer()
	for range b.N {
		clb.FindBreakOpportunities(text)
	}
}

func BenchmarkAnalyzeCJKContent(b *testing.B) {
	text := "\u3053\u3093\u306B\u3061\u306F\u4E16\u754C\u3002Hello World\u30AB\u30BF\u30AB\u30CA"

	b.ResetTimer()
	for range b.N {
		AnalyzeCJKContent(text)
	}
}
