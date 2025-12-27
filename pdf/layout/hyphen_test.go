package layout

import (
	"strings"
	"testing"
)

func TestSoftHyphenConstants(t *testing.T) {
	// Verify the constants are correct
	if SoftHyphen != '\u00AD' {
		t.Errorf(
			"SoftHyphen = %U, want U+00AD",
			SoftHyphen,
		)
	}
	if VisibleHyphen != '-' {
		t.Errorf(
			"VisibleHyphen = %q, want '-'",
			VisibleHyphen,
		)
	}
}

func TestNewSoftHyphenProcessor(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		wantCount int
	}{
		{"empty string", "", 0},
		{"no soft hyphens", "hello world", 0},
		{
			"one soft hyphen",
			"interna\u00ADtional",
			1,
		},
		{
			"two soft hyphens",
			"in\u00ADterna\u00ADtional",
			2,
		},
		{"only soft hyphen", "\u00AD", 1},
		{
			"multiple consecutive",
			"\u00AD\u00AD\u00AD",
			3,
		},
		{"at start", "\u00ADhello", 1},
		{"at end", "hello\u00AD", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := NewSoftHyphenProcessor(
				tt.text,
			)
			if processor == nil {
				t.Fatal(
					"NewSoftHyphenProcessor returned nil",
				)
			}
			if processor.Count() != tt.wantCount {
				t.Errorf(
					"Count() = %d, want %d",
					processor.Count(),
					tt.wantCount,
				)
			}
		})
	}
}

func TestSoftHyphenProcessor_OriginalText(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	if processor.OriginalText() != text {
		t.Errorf(
			"OriginalText() = %q, want %q",
			processor.OriginalText(),
			text,
		)
	}
}

func TestSoftHyphenProcessor_CleanText(
	t *testing.T,
) {
	tests := []struct {
		name      string
		text      string
		wantClean string
	}{
		{
			"no soft hyphens",
			"hello world",
			"hello world",
		},
		{
			"one soft hyphen",
			"interna\u00ADtional",
			"international",
		},
		{
			"two soft hyphens",
			"in\u00ADterna\u00ADtional",
			"international",
		},
		{"at start", "\u00ADhello", "hello"},
		{"at end", "hello\u00AD", "hello"},
		{
			"multiple consecutive",
			"a\u00AD\u00AD\u00ADb",
			"ab",
		},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := NewSoftHyphenProcessor(
				tt.text,
			)
			if processor.CleanText() != tt.wantClean {
				t.Errorf(
					"CleanText() = %q, want %q",
					processor.CleanText(),
					tt.wantClean,
				)
			}
		})
	}
}

func TestSoftHyphenProcessor_HasSoftHyphens(
	t *testing.T,
) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			"has soft hyphen",
			"interna\u00ADtional",
			true,
		},
		{
			"no soft hyphen",
			"international",
			false,
		},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := NewSoftHyphenProcessor(
				tt.text,
			)
			if processor.HasSoftHyphens() != tt.want {
				t.Errorf(
					"HasSoftHyphens() = %v, want %v",
					processor.HasSoftHyphens(),
					tt.want,
				)
			}
		})
	}
}

func TestSoftHyphenProcessor_Positions(
	t *testing.T,
) {
	text := "in\u00ADterna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	positions := processor.Positions()
	if len(positions) != 2 {
		t.Fatalf(
			"Positions() returned %d positions, want 2",
			len(positions),
		)
	}

	// First soft hyphen is after "in"
	if positions[0].RuneOffset != 2 {
		t.Errorf(
			"First position RuneOffset = %d, want 2",
			positions[0].RuneOffset,
		)
	}

	// Second soft hyphen is after "terna" (but we need to account for first soft hyphen)
	if positions[1].RuneOffset != 8 {
		t.Errorf(
			"Second position RuneOffset = %d, want 8",
			positions[1].RuneOffset,
		)
	}
}

func TestSoftHyphenProcessor_FindBreakPoints(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	breakPoints := processor.FindBreakPoints()
	if len(breakPoints) != 1 {
		t.Fatalf(
			"FindBreakPoints() returned %d points, want 1",
			len(breakPoints),
		)
	}

	// Soft hyphen is at byte position 7 (after "interna"), and is 2 bytes
	// Break point should be after the soft hyphen
	expectedBreakPoint := 7 + len(
		string(SoftHyphen),
	)
	if breakPoints[0] != expectedBreakPoint {
		t.Errorf(
			"Break point = %d, want %d",
			breakPoints[0],
			expectedBreakPoint,
		)
	}
}

func TestSoftHyphenProcessor_BreakAt(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	result := processor.BreakAt(0)
	if result == nil {
		t.Fatal("BreakAt(0) returned nil")
	}

	// Before part should be "interna" + visible hyphen
	if result.BeforePartClean != "interna-" {
		t.Errorf(
			"BeforePartClean = %q, want %q",
			result.BeforePartClean,
			"interna-",
		)
	}

	// After part should be "tional"
	if result.AfterPartClean != "tional" {
		t.Errorf(
			"AfterPartClean = %q, want %q",
			result.AfterPartClean,
			"tional",
		)
	}

	if !result.HyphenAdded {
		t.Error("HyphenAdded should be true")
	}
}

func TestSoftHyphenProcessor_BreakAt_OutOfRange(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	if processor.BreakAt(-1) != nil {
		t.Error("BreakAt(-1) should return nil")
	}

	if processor.BreakAt(5) != nil {
		t.Error(
			"BreakAt(5) should return nil for text with 1 soft hyphen",
		)
	}
}

func TestSoftHyphenProcessor_BreakAtPosition(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	// Break at position after the soft hyphen
	softHyphenEnd := 7 + len(string(SoftHyphen))
	result := processor.BreakAtPosition(
		softHyphenEnd,
	)
	if result == nil {
		t.Fatal("BreakAtPosition returned nil")
	}

	if result.BeforePartClean != "interna-" {
		t.Errorf(
			"BeforePartClean = %q, want %q",
			result.BeforePartClean,
			"interna-",
		)
	}
}

func TestSoftHyphenProcessor_BreakAtPosition_NoMatch(
	t *testing.T,
) {
	text := "interna\u00ADtional"
	processor := NewSoftHyphenProcessor(text)

	// Break at position before the soft hyphen
	result := processor.BreakAtPosition(3)
	if result != nil {
		t.Error(
			"BreakAtPosition(3) should return nil when no soft hyphen at or before position",
		)
	}
}

func TestSoftHyphenProcessor_GetSegments(
	t *testing.T,
) {
	tests := []struct {
		name         string
		text         string
		wantSegments []string
	}{
		{
			"no soft hyphens",
			"hello",
			[]string{"hello"},
		},
		{
			"one soft hyphen",
			"inter\u00ADnal",
			[]string{"inter", "nal"},
		},
		{
			"two soft hyphens",
			"in\u00ADter\u00ADnal",
			[]string{"in", "ter", "nal"},
		},
		{
			"at start",
			"\u00ADhello",
			[]string{"hello"},
		},
		{
			"at end",
			"hello\u00AD",
			[]string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := NewSoftHyphenProcessor(
				tt.text,
			)
			segments := processor.GetSegments()

			if len(
				segments,
			) != len(
				tt.wantSegments,
			) {
				t.Fatalf(
					"GetSegments() returned %d segments, want %d: %v",
					len(segments),
					len(tt.wantSegments),
					segments,
				)
			}

			for i, seg := range segments {
				if seg != tt.wantSegments[i] {
					t.Errorf(
						"Segment[%d] = %q, want %q",
						i,
						seg,
						tt.wantSegments[i],
					)
				}
			}
		})
	}
}

func TestRemoveSoftHyphens(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			"no soft hyphens",
			"hello world",
			"hello world",
		},
		{
			"one soft hyphen",
			"interna\u00ADtional",
			"international",
		},
		{"multiple", "a\u00ADb\u00ADc", "abc"},
		{"empty", "", ""},
		{"only soft hyphen", "\u00AD", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSoftHyphens(tt.text); got != tt.want {
				t.Errorf(
					"RemoveSoftHyphens(%q) = %q, want %q",
					tt.text,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestContainsSoftHyphen(t *testing.T) {
	tests := []struct {
		text string
		want bool
	}{
		{"hello", false},
		{"interna\u00ADtional", true},
		{"", false},
		{"\u00AD", true},
	}

	for _, tt := range tests {
		if got := ContainsSoftHyphen(tt.text); got != tt.want {
			t.Errorf(
				"ContainsSoftHyphen(%q) = %v, want %v",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func TestCountSoftHyphens(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"hello", 0},
		{"interna\u00ADtional", 1},
		{"a\u00ADb\u00ADc\u00ADd", 3},
		{"", 0},
	}

	for _, tt := range tests {
		if got := CountSoftHyphens(tt.text); got != tt.want {
			t.Errorf(
				"CountSoftHyphens(%q) = %d, want %d",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func TestInsertSoftHyphen(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		runePos int
		want    string
	}{
		{
			"middle",
			"international",
			7,
			"interna\u00ADtional",
		},
		{"start", "hello", 0, "\u00ADhello"},
		{"end", "hello", 5, "hello\u00AD"},
		{"negative", "hello", -1, "hello"},
		{"too large", "hello", 10, "hello"},
		{"empty", "", 0, "\u00AD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertSoftHyphen(
				tt.text,
				tt.runePos,
			)
			if got != tt.want {
				t.Errorf(
					"InsertSoftHyphen(%q, %d) = %q, want %q",
					tt.text,
					tt.runePos,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestInsertSoftHyphens(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		positions []int
		want      string
	}{
		{
			"single position",
			"international",
			[]int{7},
			"interna\u00ADtional",
		},
		{
			"multiple positions",
			"international",
			[]int{2, 7},
			"in\u00ADterna\u00ADtional",
		},
		{
			"empty positions",
			"hello",
			[]int{},
			"hello",
		},
		{
			"out of order positions",
			"hello",
			[]int{3, 1},
			"hel\u00ADlo",
		}, // out of order positions are skipped
		{
			"invalid positions",
			"hello",
			[]int{-1, 10},
			"hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertSoftHyphens(
				tt.text,
				tt.positions,
			)
			if got != tt.want {
				t.Errorf(
					"InsertSoftHyphens(%q, %v) = %q, want %q",
					tt.text,
					tt.positions,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestFindSoftHyphenBreaks(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		wantBreaks int
	}{
		{"no soft hyphens", "hello world", 0},
		{
			"one soft hyphen",
			"interna\u00ADtional",
			1,
		},
		{
			"two soft hyphens",
			"in\u00ADterna\u00ADtional",
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			breaks := FindSoftHyphenBreaks(
				tt.text,
			)
			if len(breaks) != tt.wantBreaks {
				t.Errorf(
					"FindSoftHyphenBreaks(%q) returned %d breaks, want %d",
					tt.text,
					len(breaks),
					tt.wantBreaks,
				)
			}

			// All breaks should require visible hyphen
			for i, b := range breaks {
				if !b.RequiresVisibleHyphen {
					t.Errorf(
						"Break %d RequiresVisibleHyphen = false, want true",
						i,
					)
				}
				if b.BreakOpportunity.Type != BreakAllowed {
					t.Errorf(
						"Break %d Type = %v, want BreakAllowed",
						i,
						b.BreakOpportunity.Type,
					)
				}
			}
		})
	}
}

func TestProcessLineForSoftHyphens(t *testing.T) {
	tests := []struct {
		name               string
		lineText           string
		breakAtSoftHyphen  bool
		wantText           string
		wantEndsWithHyphen bool
	}{
		{
			"no soft hyphens",
			"hello",
			false,
			"hello",
			false,
		},
		{
			"ends with soft hyphen",
			"interna\u00AD",
			true,
			"interna-",
			true,
		},
		{
			"contains soft hyphen but doesn't end with it",
			"interna\u00ADtional",
			false,
			"international",
			false,
		},
		{
			"break at soft hyphen position",
			"interna\u00ADtional",
			true,
			"international",
			false,
		},
		{
			"empty",
			"",
			false,
			"",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProcessLineForSoftHyphens(
				tt.lineText,
				tt.breakAtSoftHyphen,
			)
			if result.Text != tt.wantText {
				t.Errorf(
					"Text = %q, want %q",
					result.Text,
					tt.wantText,
				)
			}
			if result.EndsWithHyphen != tt.wantEndsWithHyphen {
				t.Errorf(
					"EndsWithHyphen = %v, want %v",
					result.EndsWithHyphen,
					tt.wantEndsWithHyphen,
				)
			}
		})
	}
}

func TestSplitIntoLinesWithSoftHyphens(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	tests := []struct {
		name      string
		text      string
		maxWidth  float64
		wantLines int
	}{
		{"no soft hyphens", "hello world", 20, 1},
		{"break at space", "hello world", 6, 2},
		{
			"break at soft hyphen",
			"interna\u00ADtional",
			8,
			2,
		},
		{"empty string", "", 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := SplitIntoLinesWithSoftHyphens(
				tt.text,
				tt.maxWidth,
				measureFunc,
			)
			if len(results) != tt.wantLines {
				t.Errorf(
					"SplitIntoLinesWithSoftHyphens(%q, %v) = %d lines, want %d",
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

func TestSplitIntoLinesWithSoftHyphens_HyphenRendered(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	// Test that when breaking at a soft hyphen, the visible hyphen is added
	text := "interna\u00ADtional"
	maxWidth := 8.0 // Should break after "interna"

	results := SplitIntoLinesWithSoftHyphens(
		text,
		maxWidth,
		measureFunc,
	)

	if len(results) != 2 {
		t.Fatalf(
			"Expected 2 lines, got %d",
			len(results),
		)
	}

	// First line should end with hyphen
	if !strings.HasSuffix(results[0].Text, "-") {
		t.Errorf(
			"First line %q should end with hyphen",
			results[0].Text,
		)
	}

	// Verify the content
	if results[0].Text != "interna-" {
		t.Errorf(
			"First line = %q, want %q",
			results[0].Text,
			"interna-",
		)
	}
	if results[1].Text != "tional" {
		t.Errorf(
			"Second line = %q, want %q",
			results[1].Text,
			"tional",
		)
	}
}

func TestSplitIntoLinesWithSoftHyphens_NoBreakNeeded(
	t *testing.T,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}

	// When no break is needed, soft hyphens should be removed but no hyphen added
	text := "interna\u00ADtional"
	maxWidth := 20.0 // Large enough to fit the whole text

	results := SplitIntoLinesWithSoftHyphens(
		text,
		maxWidth,
		measureFunc,
	)

	if len(results) != 1 {
		t.Fatalf(
			"Expected 1 line, got %d",
			len(results),
		)
	}

	// Text should have soft hyphen removed but no visible hyphen added
	if results[0].Text != "international" {
		t.Errorf(
			"Line text = %q, want %q",
			results[0].Text,
			"international",
		)
	}
}

func TestSoftHyphenProcessor_MultipleSoftHyphens(
	t *testing.T,
) {
	// Test with a word that has multiple hyphenation points
	text := "un\u00ADbe\u00ADliev\u00ADable"
	processor := NewSoftHyphenProcessor(text)

	if processor.Count() != 3 {
		t.Errorf(
			"Count() = %d, want 3",
			processor.Count(),
		)
	}

	if processor.CleanText() != "unbelievable" {
		t.Errorf(
			"CleanText() = %q, want %q",
			processor.CleanText(),
			"unbelievable",
		)
	}

	// Test breaking at different positions
	for i := range 3 {
		result := processor.BreakAt(i)
		if result == nil {
			t.Errorf(
				"BreakAt(%d) returned nil",
				i,
			)

			continue
		}
		if !result.HyphenAdded {
			t.Errorf(
				"BreakAt(%d) HyphenAdded = false",
				i,
			)
		}
		if !strings.HasSuffix(
			result.BeforePartClean,
			"-",
		) {
			t.Errorf(
				"BreakAt(%d) BeforePartClean = %q, should end with hyphen",
				i,
				result.BeforePartClean,
			)
		}
	}
}

func TestSoftHyphenProcessor_UnicodeText(
	t *testing.T,
) {
	// Test with non-ASCII text
	text := "\u00FCber\u00ADsetzung" // ubersetzung with umlaut and soft hyphen
	processor := NewSoftHyphenProcessor(text)

	if processor.Count() != 1 {
		t.Errorf(
			"Count() = %d, want 1",
			processor.Count(),
		)
	}

	cleanText := processor.CleanText()
	if cleanText != "\u00FCbersetzung" {
		t.Errorf(
			"CleanText() = %q, want %q",
			cleanText,
			"\u00FCbersetzung",
		)
	}

	result := processor.BreakAt(0)
	if result == nil {
		t.Fatal("BreakAt(0) returned nil")
	}

	if result.BeforePartClean != "\u00FCber-" {
		t.Errorf(
			"BeforePartClean = %q, want %q",
			result.BeforePartClean,
			"\u00FCber-",
		)
	}
}

// Benchmark tests

func BenchmarkNewSoftHyphenProcessor_Short(
	b *testing.B,
) {
	text := "interna\u00ADtional"
	b.ResetTimer()
	for range b.N {
		NewSoftHyphenProcessor(text)
	}
}

func BenchmarkNewSoftHyphenProcessor_Long(
	b *testing.B,
) {
	text := "un\u00ADbe\u00ADliev\u00ADable in\u00ADterna\u00ADtional or\u00ADgani\u00ADza\u00ADtion"
	b.ResetTimer()
	for range b.N {
		NewSoftHyphenProcessor(text)
	}
}

func BenchmarkRemoveSoftHyphens(b *testing.B) {
	text := "un\u00ADbe\u00ADliev\u00ADable in\u00ADterna\u00ADtional"
	b.ResetTimer()
	for range b.N {
		RemoveSoftHyphens(text)
	}
}

func BenchmarkSplitIntoLinesWithSoftHyphens(
	b *testing.B,
) {
	measureFunc := func(text string) float64 {
		return float64(len([]rune(text)))
	}
	text := "The interna\u00ADtional organi\u00ADzation held an un\u00ADbe\u00ADliev\u00ADable event."
	maxWidth := 20.0

	b.ResetTimer()
	for range b.N {
		SplitIntoLinesWithSoftHyphens(
			text,
			maxWidth,
			measureFunc,
		)
	}
}

func BenchmarkFindSoftHyphenBreaks(b *testing.B) {
	text := "The interna\u00ADtional organi\u00ADzation held an event."

	b.ResetTimer()
	for range b.N {
		FindSoftHyphenBreaks(text)
	}
}
