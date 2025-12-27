// bidi.go provides Unicode Bidirectional Algorithm (UAX #9) support.
// It handles reordering of mixed RTL and LTR text for visual rendering.

package layout

import (
	"strings"

	"golang.org/x/text/unicode/bidi"
)

// Direction represents the base direction of text.
type Direction int

const (
	// LTR represents Left-to-Right text.
	LTR Direction = iota
	// RTL represents Right-to-Left text.
	RTL
)

// ReorderedLine contains the visual order of text segments.
type ReorderedLine struct {
	// Text is the visuals-ordered text.
	Text string
	// Runs contains the directional runs in visual order.
	Runs []bidi.Run
}

// BidiProcessor handles bidirectional text analysis and reordering.
type BidiProcessor struct {
	paragraph bidi.Paragraph
}

// NewBidiProcessor creates a new BidiProcessor.
func NewBidiProcessor() *BidiProcessor {
	return &BidiProcessor{}
}

// Process analyzes the text and prepares it for reordering.
func (bp *BidiProcessor) Process(
	text string,
	baseDir Direction,
) error {
	if text == "" {
		return nil
	}

	var options []bidi.Option
	if baseDir == RTL {
		options = append(
			options,
			bidi.DefaultDirection(
				bidi.RightToLeft,
			),
		)
	} else {
		options = append(options, bidi.DefaultDirection(bidi.LeftToRight))
	}

	_, err := bp.paragraph.SetString(
		text,
		options...)
	return err
}

// Reorder returns the visual order of the text for a specific range (e.g., a line).
// The start and end parameters are byte offsets in the original string.
func (bp *BidiProcessor) Reorder(
	start, end int,
) (ReorderedLine, error) {
	if start >= end {
		return ReorderedLine{}, nil
	}

	// Line returns Ordering for a range
	order, err := bp.paragraph.Line(start, end)
	if err != nil {
		return ReorderedLine{}, err
	}

	var sb strings.Builder
	var runs []bidi.Run
	for i := 0; i < order.NumRuns(); i++ {
		r := order.Run(i)
		sb.WriteString(r.String())
		runs = append(runs, r)
	}

	return ReorderedLine{
		Text: sb.String(),
		Runs: runs,
	}, nil
}

// GetBaseDirection detects the base direction of the text.
func GetBaseDirection(text string) Direction {
	if text == "" {
		return LTR
	}

	p := bidi.Paragraph{}
	if _, err := p.SetString(text); err != nil {
		return LTR
	}

	// Based on the panic trace, let's try to get ordering first
	// as Paragraph.Direction() seems to rely on computed ordering.
	if _, err := p.Order(); err != nil {
		return LTR
	}

	if p.Direction() == bidi.RightToLeft {
		return RTL
	}
	return LTR
}
