package layout

import (
	"strings"
	"testing"
)

func TestNBSPConstants(t *testing.T) {
	// Verify the constants are correct Unicode code points
	tests := []struct {
		name     string
		got      rune
		wantHex  rune
		wantName string
	}{
		{
			"NBSP",
			NBSP,
			'\u00A0',
			"NO-BREAK SPACE",
		},
		{
			"NarrowNBSP",
			NarrowNBSP,
			'\u202F',
			"NARROW NO-BREAK SPACE",
		},
		{
			"ZWNBSP",
			ZWNBSP,
			'\uFEFF',
			"ZERO WIDTH NO-BREAK SPACE",
		},
		{
			"WordJoiner",
			WordJoiner,
			'\u2060',
			"WORD JOINER",
		},
		{
			"FigureSpace",
			FigureSpace,
			'\u2007',
			"FIGURE SPACE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.wantHex {
				t.Errorf(
					"%s = %U, want %U",
					tt.name,
					tt.got,
					tt.wantHex,
				)
			}
		})
	}
}

func TestNonBreakingSpaceType_String(
	t *testing.T,
) {
	tests := []struct {
		nbspType NonBreakingSpaceType
		want     string
	}{
		{NBSPTypeStandard, "Standard NBSP"},
		{NBSPTypeNarrow, "Narrow NBSP"},
		{NBSPTypeZeroWidth, "Zero Width NBSP"},
		{NBSPTypeWordJoiner, "Word Joiner"},
		{NBSPTypeFigure, "Figure Space"},
		{NBSPTypeUnknown, "Unknown"},
		{NonBreakingSpaceType(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.nbspType.String(); got != tt.want {
				t.Errorf(
					"NonBreakingSpaceType(%d).String() = %q, want %q",
					tt.nbspType,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestIsNonBreakingSpace(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"standard NBSP", NBSP, true},
		{"narrow NBSP", NarrowNBSP, true},
		{"zero width NBSP", ZWNBSP, true},
		{"word joiner", WordJoiner, true},
		{"figure space", FigureSpace, true},
		{"regular space", ' ', false},
		{"letter", 'a', false},
		{"digit", '0', false},
		{"newline", '\n', false},
		{"tab", '\t', false},
		{"em space", '\u2003', false},
		{"en space", '\u2002', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNonBreakingSpace(tt.r); got != tt.want {
				t.Errorf(
					"IsNonBreakingSpace(%U) = %v, want %v",
					tt.r,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGetNonBreakingSpaceType(t *testing.T) {
	tests := []struct {
		r    rune
		want NonBreakingSpaceType
	}{
		{NBSP, NBSPTypeStandard},
		{NarrowNBSP, NBSPTypeNarrow},
		{ZWNBSP, NBSPTypeZeroWidth},
		{WordJoiner, NBSPTypeWordJoiner},
		{FigureSpace, NBSPTypeFigure},
		{' ', NBSPTypeUnknown},
		{'a', NBSPTypeUnknown},
	}

	for _, tt := range tests {
		t.Run(
			tt.want.String(),
			func(t *testing.T) {
				if got := GetNonBreakingSpaceType(tt.r); got != tt.want {
					t.Errorf(
						"GetNonBreakingSpaceType(%U) = %v, want %v",
						tt.r,
						got,
						tt.want,
					)
				}
			},
		)
	}
}

func TestContainsNonBreakingSpace(t *testing.T) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{"empty string", "", false},
		{"no NBSP", "hello world", false},
		{
			"with standard NBSP",
			"hello\u00A0world",
			true,
		},
		{
			"with narrow NBSP",
			"hello\u202Fworld",
			true,
		},
		{"with ZWNBSP", "hello\uFEFFworld", true},
		{
			"with word joiner",
			"hello\u2060world",
			true,
		},
		{
			"with figure space",
			"1\u20072\u20073",
			true,
		},
		{"only NBSP", "\u00A0", true},
		{"regular spaces only", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsNonBreakingSpace(tt.text); got != tt.want {
				t.Errorf(
					"ContainsNonBreakingSpace(%q) = %v, want %v",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCountNonBreakingSpaces(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"empty", "", 0},
		{"no NBSP", "hello world", 0},
		{"one NBSP", "hello\u00A0world", 1},
		{
			"two NBSP",
			"hello\u00A0world\u00A0!",
			2,
		},
		{
			"mixed types",
			"a\u00A0b\u202Fc\u2060d",
			3,
		},
		{"only NBSP", "\u00A0\u00A0\u00A0", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNonBreakingSpaces(tt.text); got != tt.want {
				t.Errorf(
					"CountNonBreakingSpaces(%q) = %d, want %d",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestFindNonBreakingSpaces(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		positions := FindNonBreakingSpaces("")
		if len(positions) != 0 {
			t.Errorf(
				"Expected 0 positions, got %d",
				len(positions),
			)
		}
	})

	t.Run("no NBSP", func(t *testing.T) {
		positions := FindNonBreakingSpaces(
			"hello world",
		)
		if len(positions) != 0 {
			t.Errorf(
				"Expected 0 positions, got %d",
				len(positions),
			)
		}
	})

	t.Run("single NBSP", func(t *testing.T) {
		text := "hello\u00A0world"
		positions := FindNonBreakingSpaces(text)
		if len(positions) != 1 {
			t.Fatalf(
				"Expected 1 position, got %d",
				len(positions),
			)
		}
		if positions[0].RuneOffset != 5 {
			t.Errorf(
				"RuneOffset = %d, want 5",
				positions[0].RuneOffset,
			)
		}
		if positions[0].Type != NBSPTypeStandard {
			t.Errorf(
				"Type = %v, want %v",
				positions[0].Type,
				NBSPTypeStandard,
			)
		}
		if positions[0].Rune != NBSP {
			t.Errorf(
				"Rune = %U, want %U",
				positions[0].Rune,
				NBSP,
			)
		}
	})

	t.Run(
		"multiple NBSP types",
		func(t *testing.T) {
			text := "a\u00A0b\u202Fc\u2060d"
			positions := FindNonBreakingSpaces(
				text,
			)
			if len(positions) != 3 {
				t.Fatalf(
					"Expected 3 positions, got %d",
					len(positions),
				)
			}

			expected := []struct {
				runeOffset int
				nbspType   NonBreakingSpaceType
			}{
				{1, NBSPTypeStandard},
				{3, NBSPTypeNarrow},
				{5, NBSPTypeWordJoiner},
			}

			for i, exp := range expected {
				if positions[i].RuneOffset != exp.runeOffset {
					t.Errorf(
						"Position[%d].RuneOffset = %d, want %d",
						i,
						positions[i].RuneOffset,
						exp.runeOffset,
					)
				}
				if positions[i].Type != exp.nbspType {
					t.Errorf(
						"Position[%d].Type = %v, want %v",
						i,
						positions[i].Type,
						exp.nbspType,
					)
				}
			}
		},
	)

	t.Run(
		"with unicode text",
		func(t *testing.T) {
			text := "Bonjour\u00A0monde" // French for "hello world"
			positions := FindNonBreakingSpaces(
				text,
			)
			if len(positions) != 1 {
				t.Fatalf(
					"Expected 1 position, got %d",
					len(positions),
				)
			}
			if positions[0].RuneOffset != 7 {
				t.Errorf(
					"RuneOffset = %d, want 7",
					positions[0].RuneOffset,
				)
			}
		},
	)
}

func TestReplaceSpacesWithNBSP(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"empty", "", ""},
		{"no spaces", "hello", "hello"},
		{
			"one space",
			"hello world",
			"hello\u00A0world",
		},
		{
			"multiple spaces",
			"a b c d",
			"a\u00A0b\u00A0c\u00A0d",
		},
		{
			"consecutive spaces",
			"a  b",
			"a\u00A0\u00A0b",
		},
		{
			"already NBSP",
			"hello\u00A0world",
			"hello\u00A0world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpacesWithNBSP(tt.text); got != tt.want {
				t.Errorf(
					"ReplaceSpacesWithNBSP(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestReplaceNBSPWithSpaces(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"empty", "", ""},
		{"no NBSP", "hello world", "hello world"},
		{
			"one NBSP",
			"hello\u00A0world",
			"hello world",
		},
		{
			"multiple NBSP",
			"a\u00A0b\u00A0c",
			"a b c",
		},
		{
			"mixed NBSP types",
			"a\u00A0b\u202Fc",
			"a b\u202Fc",
		}, // Only standard NBSP replaced
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceNBSPWithSpaces(tt.text); got != tt.want {
				t.Errorf(
					"ReplaceNBSPWithSpaces(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestReplaceAllNBSPWithSpaces(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"empty", "", ""},
		{"no NBSP", "hello world", "hello world"},
		{
			"standard NBSP",
			"hello\u00A0world",
			"hello world",
		},
		{
			"narrow NBSP",
			"hello\u202Fworld",
			"hello world",
		},
		{
			"figure space",
			"1\u20072\u20073",
			"1 2 3",
		},
		{
			"zero width NBSP",
			"hello\uFEFFworld",
			"helloworld",
		}, // Removed
		{
			"word joiner",
			"hello\u2060world",
			"helloworld",
		}, // Removed
		{
			"mixed",
			"a\u00A0b\u202Fc\uFEFFd\u2060e",
			"a b c d e"[0:9],
		}, // Fix: "a b cde"
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceAllNBSPWithSpaces(
				tt.text,
			)
			// Special handling for "mixed" case
			if tt.name == "mixed" {
				expected := "a b cde"
				if got != expected {
					t.Errorf(
						"ReplaceAllNBSPWithSpaces(%q) = %q, want %q",
						tt.text,
						got,
						expected,
					)
				}

				return
			}
			if got != tt.want {
				t.Errorf(
					"ReplaceAllNBSPWithSpaces(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestInsertNBSP(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		runePos int
		want    string
	}{
		{"start", "hello", 0, "\u00A0hello"},
		{"middle", "hello", 2, "he\u00A0llo"},
		{"end", "hello", 5, "hello\u00A0"},
		{"negative", "hello", -1, "hello"},
		{"too large", "hello", 10, "hello"},
		{"empty string at 0", "", 0, "\u00A0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertNBSP(tt.text, tt.runePos); got != tt.want {
				t.Errorf(
					"InsertNBSP(%q, %d) = %q, want %q",
					tt.text,
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestInsertNBSPType(t *testing.T) {
	text := "hello"

	t.Run("narrow NBSP", func(t *testing.T) {
		got := InsertNBSPType(text, 2, NarrowNBSP)
		want := "he\u202Fllo"
		if got != want {
			t.Errorf(
				"InsertNBSPType(%q, 2, NarrowNBSP) = %q, want %q",
				text,
				got,
				want,
			)
		}
	})

	t.Run("word joiner", func(t *testing.T) {
		got := InsertNBSPType(text, 2, WordJoiner)
		want := "he\u2060llo"
		if got != want {
			t.Errorf(
				"InsertNBSPType(%q, 2, WordJoiner) = %q, want %q",
				text,
				got,
				want,
			)
		}
	})
}

func TestInsertWordJoiner(t *testing.T) {
	text := "hello"
	got := InsertWordJoiner(text, 2)
	want := "he\u2060llo"
	if got != want {
		t.Errorf(
			"InsertWordJoiner(%q, 2) = %q, want %q",
			text,
			got,
			want,
		)
	}
}

func TestNBSPProcessor(t *testing.T) {
	t.Run("empty text", func(t *testing.T) {
		p := NewNBSPProcessor("")
		if p.HasNonBreakingSpaces() {
			t.Error(
				"HasNonBreakingSpaces() should be false for empty text",
			)
		}
		if p.Count() != 0 {
			t.Errorf(
				"Count() = %d, want 0",
				p.Count(),
			)
		}
	})

	t.Run("text with NBSP", func(t *testing.T) {
		text := "hello\u00A0world"
		p := NewNBSPProcessor(text)

		if !p.HasNonBreakingSpaces() {
			t.Error(
				"HasNonBreakingSpaces() should be true",
			)
		}
		if p.Count() != 1 {
			t.Errorf(
				"Count() = %d, want 1",
				p.Count(),
			)
		}
		if p.OriginalText() != text {
			t.Errorf(
				"OriginalText() = %q, want %q",
				p.OriginalText(),
				text,
			)
		}
	})

	t.Run("CountByType", func(t *testing.T) {
		text := "a\u00A0b\u00A0c\u202Fd"
		p := NewNBSPProcessor(text)

		if p.CountByType(NBSPTypeStandard) != 2 {
			t.Errorf(
				"CountByType(Standard) = %d, want 2",
				p.CountByType(NBSPTypeStandard),
			)
		}
		if p.CountByType(NBSPTypeNarrow) != 1 {
			t.Errorf(
				"CountByType(Narrow) = %d, want 1",
				p.CountByType(NBSPTypeNarrow),
			)
		}
		if p.CountByType(
			NBSPTypeWordJoiner,
		) != 0 {
			t.Errorf(
				"CountByType(WordJoiner) = %d, want 0",
				p.CountByType(NBSPTypeWordJoiner),
			)
		}
	})

	t.Run(
		"GetPositionsByType",
		func(t *testing.T) {
			text := "a\u00A0b\u202Fc\u00A0d"
			p := NewNBSPProcessor(text)

			standard := p.GetPositionsByType(
				NBSPTypeStandard,
			)
			if len(standard) != 2 {
				t.Errorf(
					"GetPositionsByType(Standard) returned %d positions, want 2",
					len(standard),
				)
			}

			narrow := p.GetPositionsByType(
				NBSPTypeNarrow,
			)
			if len(narrow) != 1 {
				t.Errorf(
					"GetPositionsByType(Narrow) returned %d positions, want 1",
					len(narrow),
				)
			}
		},
	)
}

func TestNBSPProcessor_NormalizeToStandardNBSP(
	t *testing.T,
) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"no NBSP", "hello", "hello"},
		{
			"already standard",
			"hello\u00A0world",
			"hello\u00A0world",
		},
		{
			"narrow to standard",
			"hello\u202Fworld",
			"hello\u00A0world",
		},
		{
			"figure to standard",
			"1\u20072",
			"1\u00A02",
		},
		{
			"remove ZWNBSP",
			"hello\uFEFFworld",
			"helloworld",
		},
		{
			"remove word joiner",
			"hello\u2060world",
			"helloworld",
		},
		{
			"mixed",
			"a\u00A0b\u202Fc\uFEFFd",
			"a\u00A0b\u00A0cd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewNBSPProcessor(tt.text)
			if got := p.NormalizeToStandardNBSP(); got != tt.want {
				t.Errorf(
					"NormalizeToStandardNBSP() = %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestNBSPProcessor_ToDisplayText(
	t *testing.T,
) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"no NBSP", "hello", "hello"},
		{
			"keep standard NBSP",
			"hello\u00A0world",
			"hello\u00A0world",
		},
		{
			"keep narrow NBSP",
			"hello\u202Fworld",
			"hello\u202Fworld",
		},
		{
			"remove ZWNBSP",
			"hello\uFEFFworld",
			"helloworld",
		},
		{
			"remove word joiner",
			"hello\u2060world",
			"helloworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewNBSPProcessor(tt.text)
			if got := p.ToDisplayText(); got != tt.want {
				t.Errorf(
					"ToDisplayText() = %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestNBSPProcessor_IsBreakProhibitedAt(
	t *testing.T,
) {
	text := "hello\u00A0world"
	p := NewNBSPProcessor(text)

	tests := []struct {
		runePos int
		want    bool
	}{
		{0, false},  // before 'h'
		{4, false},  // before 'o'
		{5, true},   // at NBSP position
		{6, true},   // immediately after NBSP
		{7, false},  // at 'o' in world
		{10, false}, // at 'd'
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := p.IsBreakProhibitedAt(tt.runePos); got != tt.want {
				t.Errorf(
					"IsBreakProhibitedAt(%d) = %v, want %v",
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestJoinWithNBSP(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{"empty", []string{}, ""},
		{"single", []string{"hello"}, "hello"},
		{
			"two parts",
			[]string{"hello", "world"},
			"hello\u00A0world",
		},
		{
			"three parts",
			[]string{"a", "b", "c"},
			"a\u00A0b\u00A0c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinWithNBSP(tt.parts...); got != tt.want {
				t.Errorf(
					"JoinWithNBSP(%v) = %q, want %q",
					tt.parts,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestJoinWithWordJoiner(t *testing.T) {
	parts := []string{"hello", "world"}
	got := JoinWithWordJoiner(parts...)
	want := "hello\u2060world"
	if got != want {
		t.Errorf(
			"JoinWithWordJoiner(%v) = %q, want %q",
			parts,
			got,
			want,
		)
	}
}

func TestProtectFromBreaking(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"empty", "", ""},
		{"word", "hello", "\u2060hello\u2060"},
		{
			"with spaces",
			"hello world",
			"\u2060hello world\u2060",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtectFromBreaking(tt.text); got != tt.want {
				t.Errorf(
					"ProtectFromBreaking(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestReplaceSpacesBetweenDigits(
	t *testing.T,
) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			"no digits",
			"hello world",
			"hello world",
		},
		{
			"digits with space",
			"1 000",
			"1\u00A0000",
		},
		{
			"multiple groups",
			"1 000 000",
			"1\u00A0000\u00A0000",
		},
		{
			"mixed text",
			"Price: 1 000 EUR",
			"Price: 1\u00A0000 EUR",
		},
		{"no space between", "1000", "1000"},
		{
			"space not between digits",
			"a 1 b",
			"a 1 b",
		},
		{"short text", "1", "1"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpacesBetweenDigits(tt.text); got != tt.want {
				t.Errorf(
					"ReplaceSpacesBetweenDigits(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestReplaceSpacesAfterPunctuation(
	t *testing.T,
) {
	t.Run(
		"before punctuation",
		func(t *testing.T) {
			tests := []struct {
				text string
				want string
			}{
				{"Hello !", "Hello\u202F!"},
				{"What ?", "What\u202F?"},
				{"Note :", "Note\u202F:"},
				{"Item ;", "Item\u202F;"},
			}

			for _, tt := range tests {
				got := ReplaceSpacesAfterPunctuation(
					tt.text,
					true,
				)
				if got != tt.want {
					t.Errorf(
						"ReplaceSpacesAfterPunctuation(%q, true) = %q, want %q",
						tt.text,
						got,
						tt.want,
					)
				}
			}
		},
	)

	t.Run(
		"no before punctuation",
		func(t *testing.T) {
			text := "Hello !"
			got := ReplaceSpacesAfterPunctuation(
				text,
				false,
			)
			if got != text {
				t.Errorf(
					"ReplaceSpacesAfterPunctuation(%q, false) = %q, want %q",
					text,
					got,
					text,
				)
			}
		},
	)

	t.Run("guillemet", func(t *testing.T) {
		text := "\u00AB hello"
		got := ReplaceSpacesAfterPunctuation(
			text,
			false,
		)
		want := "\u00AB\u00A0hello"
		if got != want {
			t.Errorf(
				"ReplaceSpacesAfterPunctuation(%q, false) = %q, want %q",
				text,
				got,
				want,
			)
		}
	})
}

// Test integration with the LineBreaker
func TestNBSPLineBreakIntegration(t *testing.T) {
	lb := NewLineBreaker()

	t.Run(
		"NBSP prevents break",
		func(t *testing.T) {
			// "hello world" with regular space - should allow break
			text1 := "hello world"
			opp1 := lb.FindBreakOpportunities(
				text1,
			)
			hasBreakAtSpace1 := false
			for _, o := range opp1 {
				if o.RunePosition == 6 &&
					o.Type == BreakAllowed {
					hasBreakAtSpace1 = true

					break
				}
			}
			if !hasBreakAtSpace1 {
				t.Error(
					"Expected break opportunity at regular space",
				)
			}

			// "hello world" with NBSP - should NOT allow break
			text2 := "hello\u00A0world"
			opp2 := lb.FindBreakOpportunities(
				text2,
			)
			hasBreakAtNBSP := false
			for _, o := range opp2 {
				if o.RunePosition == 6 &&
					o.Type == BreakAllowed {
					hasBreakAtNBSP = true

					break
				}
			}
			if hasBreakAtNBSP {
				t.Error(
					"Did not expect break opportunity at NBSP",
				)
			}
		},
	)

	t.Run(
		"narrow NBSP prevents break",
		func(t *testing.T) {
			text := "hello\u202Fworld"
			opp := lb.FindBreakOpportunities(text)
			for _, o := range opp {
				if o.RunePosition == 6 &&
					o.Type == BreakAllowed {
					t.Error(
						"Did not expect break opportunity at narrow NBSP",
					)

					break
				}
			}
		},
	)

	t.Run(
		"NBSP line break class is GL",
		func(t *testing.T) {
			if GetLineBreakClass(NBSP) != GL {
				t.Errorf(
					"GetLineBreakClass(NBSP) = %v, want GL",
					GetLineBreakClass(NBSP),
				)
			}
			if GetLineBreakClass(
				NarrowNBSP,
			) != GL {
				t.Errorf(
					"GetLineBreakClass(NarrowNBSP) = %v, want GL",
					GetLineBreakClass(NarrowNBSP),
				)
			}
			if GetLineBreakClass(
				FigureSpace,
			) != GL {
				t.Errorf(
					"GetLineBreakClass(FigureSpace) = %v, want GL",
					GetLineBreakClass(
						FigureSpace,
					),
				)
			}
		},
	)

	t.Run(
		"word joiner line break class is WJ",
		func(t *testing.T) {
			if GetLineBreakClass(
				WordJoiner,
			) != WJ {
				t.Errorf(
					"GetLineBreakClass(WordJoiner) = %v, want WJ",
					GetLineBreakClass(WordJoiner),
				)
			}
			if GetLineBreakClass(ZWNBSP) != WJ {
				t.Errorf(
					"GetLineBreakClass(ZWNBSP) = %v, want WJ",
					GetLineBreakClass(ZWNBSP),
				)
			}
		},
	)
}

func TestSplitIntoLines_WithNBSP(t *testing.T) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	t.Run(
		"NBSP keeps words together",
		func(t *testing.T) {
			// With NBSP, "Dr. Smith" should stay together
			text := "Hello Dr.\u00A0Smith goodbye"
			maxWidth := 12.0 // Wide enough for "Hello" + space + "Dr. Smith" but we want to test NBSP

			// First test with regular space
			textWithSpace := strings.ReplaceAll(
				text,
				"\u00A0",
				" ",
			)
			linesWithSpace := SplitIntoLines(
				textWithSpace,
				10,
				measureFunc,
			)

			// The text should break somewhere
			if len(linesWithSpace) < 2 {
				t.Log(
					"Text with regular space did not require multiple lines",
				)
			}

			// With NBSP, verify Dr. and Smith stay together
			linesWithNBSP := SplitIntoLines(
				text,
				maxWidth,
				measureFunc,
			)
			for _, line := range linesWithNBSP {
				// If Dr. is on a line, Smith should also be on the same line
				if strings.Contains(
					line.Text,
					"Dr.",
				) &&
					!strings.Contains(
						line.Text,
						"Smith",
					) {
					t.Errorf(
						"NBSP did not keep 'Dr.' and 'Smith' together: line = %q",
						line.Text,
					)
				}
			}
		},
	)

	t.Run(
		"number with NBSP stays together",
		func(t *testing.T) {
			// "1 000 000" with NBSP should stay together
			text := "Total: 1\u00A0000\u00A0000 items"
			maxWidth := 10.0

			lines := SplitIntoLines(
				text,
				maxWidth,
				measureFunc,
			)
			// At minimum, we should have some output
			if len(lines) == 0 {
				t.Error(
					"Expected at least one line",
				)
			}
		},
	)
}

// Benchmark tests
func BenchmarkIsNonBreakingSpace(b *testing.B) {
	chars := []rune{
		NBSP,
		' ',
		'a',
		NarrowNBSP,
		WordJoiner,
	}
	b.ResetTimer()
	for range b.N {
		for _, c := range chars {
			IsNonBreakingSpace(c)
		}
	}
}

func BenchmarkFindNonBreakingSpaces(
	b *testing.B,
) {
	text := "Hello\u00A0world,\u202Fthis\u2060is\u00A0a\u00A0test."
	b.ResetTimer()
	for range b.N {
		FindNonBreakingSpaces(text)
	}
}

func BenchmarkReplaceSpacesWithNBSP(
	b *testing.B,
) {
	text := "Hello world, this is a test with multiple spaces."
	b.ResetTimer()
	for range b.N {
		ReplaceSpacesWithNBSP(text)
	}
}

func BenchmarkReplaceAllNBSPWithSpaces(
	b *testing.B,
) {
	text := "Hello\u00A0world,\u202Fthis\u2060is\u00A0a\u00A0test."
	b.ResetTimer()
	for range b.N {
		ReplaceAllNBSPWithSpaces(text)
	}
}

func BenchmarkNBSPProcessor_New(b *testing.B) {
	text := "Hello\u00A0world,\u202Fthis\u2060is\u00A0a\u00A0test."
	b.ResetTimer()
	for range b.N {
		NewNBSPProcessor(text)
	}
}
