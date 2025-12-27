// engine.go provides the core text layout engine for PDF rendering.
// It coordinates font handling, line breaking, and paragraph layout
// to produce positioned text runs for rendering.

package layout

import (
	"strings"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// TextLayoutEngine handles Word-compatible text layout by coordinating
// font metrics, line breaking algorithms, and paragraph layout rules.
type TextLayoutEngine struct {
	// fontCache provides access to font metrics and glyph data.
	fontCache *font.FontCache

	// lineBreaker implements Unicode line breaking (UAX #14).
	lineBreaker *LineBreaker

	// defaultFontSize is used when no font size is specified.
	defaultFontSize float64
}

// NewTextLayoutEngine creates a new TextLayoutEngine with the given font cache.
func NewTextLayoutEngine(
	fontCache *font.FontCache,
) *TextLayoutEngine {
	return &TextLayoutEngine{
		fontCache:       fontCache,
		lineBreaker:     NewLineBreaker(),
		defaultFontSize: 12.0, // Default 12pt
	}
}

// FontCache returns the font cache used by the engine.
func (e *TextLayoutEngine) FontCache() *font.FontCache {
	if e == nil {
		return nil
	}

	return e.fontCache
}

// SetDefaultFontSize sets the default font size for the engine.
func (e *TextLayoutEngine) SetDefaultFontSize(
	size float64,
) {
	if e != nil {
		e.defaultFontSize = size
	}
}

// DefaultFontSize returns the current default font size.
func (e *TextLayoutEngine) DefaultFontSize() float64 {
	if e == nil {
		return 0
	}

	return e.defaultFontSize
}

// LayoutOptions contains options for the layout process.
type LayoutOptions struct {
	// MaxWidth is the maximum width of the layout area.
	MaxWidth float64

	// MaxHeight is the maximum height of the layout area (optional).
	MaxHeight float64

	// Hyphenation enables or disables automatic hyphenation.
	Hyphenation bool

	// Justify enables or disables full justification.
	Justify bool
}

// DefaultLayoutOptions returns default layout options.
func DefaultLayoutOptions() LayoutOptions {
	return LayoutOptions{
		MaxWidth:    500.0, // Default width
		Hyphenation: false,
		Justify:     false,
	}
}

// LayoutLine represents a single line of laid-out text.
type LayoutLine struct {
	// Runs contains the positioned text runs on this line.
	Runs []*PositionedRun

	// Width is the total width of the line.
	Width float64

	// Height is the total height of the line.
	Height float64

	// Ascent is the maximum ascent of the line.
	Ascent float64

	// Descent is the maximum descent of the line.
	Descent float64

	// Y is the vertical position of the baseline.
	Y float64

	// Indent is the horizontal indentation of the line from the left margin.
	Indent float64

	// ForcedBreak is true if this line ended with a mandatory break.
	ForcedBreak bool
}

// PositionedRun represents a text run with a specific position on a line.
type PositionedRun struct {
	// Run is the source text run.
	Run *TextRun

	// X is the horizontal offset from the start of the line.
	X float64

	// Width is the width of this positioned run.
	Width float64
}

// LayoutParagraph performs layout for a single paragraph, breaking it into lines
// according to the specified options and paragraph properties.
func (e *TextLayoutEngine) LayoutParagraph(
	p *Paragraph,
	opts LayoutOptions,
) []LayoutLine {
	if p == nil || p.IsEmpty() {
		return nil
	}

	fullText := p.Text()
	runes := []rune(fullText)
	baseDir := GetBaseDirection(fullText)

	bp := NewBidiProcessor()
	bp.Process(fullText, baseDir)

	// fast path: if entire paragraph fits on one line
	firstLineAvail := MeasureFirstLine(
		p,
		opts.MaxWidth,
	)
	if e.measureRuneRange(
		p,
		0,
		len(runes),
		0,
	) <= firstLineAvail {
		line := e.createLineFromRange(
			p,
			0,
			len(runes),
			firstLineAvail,
			false,
		)
		line.Indent = p.Properties.Indentation.EffectiveFirstLineIndent()
		e.reorderRuns(&line, bp, 0, len(runes))

		// Alignment
		align := p.Properties.Alignment
		if align == AlignJustify {
			align = AlignLeft
		}
		e.applyAlignment(
			&line,
			align,
			firstLineAvail,
		)
		return []LayoutLine{line}
	}

	opportunities := e.lineBreaker.FindBreakOpportunities(
		fullText,
	)
	var lines []LayoutLine
	startRune := 0

	for startRune < len(runes) {
		isFirstLine := len(lines) == 0
		availableWidth := opts.MaxWidth
		indent := 0.0
		if isFirstLine {
			indent = p.Properties.Indentation.EffectiveFirstLineIndent()
			availableWidth = firstLineAvail
		} else {
			indent = p.Properties.Indentation.EffectiveSubsequentIndent()
			availableWidth = MeasureSubsequentLines(p, opts.MaxWidth)
		}

		bestBreakRune := -1
		forcedBreak := false

		for _, opp := range opportunities {
			if opp.RunePosition <= startRune {
				continue
			}

			if opp.Type == BreakMandatory {
				width := e.measureRuneRange(
					p,
					startRune,
					opp.RunePosition,
					0,
				)
				if width <= availableWidth ||
					bestBreakRune == -1 {
					bestBreakRune = opp.RunePosition
					forcedBreak = true
					break
				}
			}

			width := e.measureRuneRange(
				p,
				startRune,
				opp.RunePosition,
				0,
			)
			if width <= availableWidth {
				bestBreakRune = opp.RunePosition
			} else if bestBreakRune != -1 {
				break
			}
		}

		if bestBreakRune == -1 {
			bestBreakRune = startRune + 1
			for i := startRune + 1; i <= len(runes); i++ {
				if e.measureRuneRange(
					p,
					startRune,
					i,
					0,
				) <= availableWidth {
					bestBreakRune = i
				} else {
					break
				}
			}
		} else if !forcedBreak {
			remainingWidth := e.measureRuneRange(p, startRune, len(runes), 0)
			if remainingWidth <= availableWidth {
				bestBreakRune = len(runes)
			}
		}

		line := e.createLineFromRange(
			p,
			startRune,
			bestBreakRune,
			availableWidth,
			forcedBreak,
		)
		line.Indent = indent

		e.reorderRuns(
			&line,
			bp,
			startRune,
			bestBreakRune,
		)

		isLastLine := bestBreakRune >= len(runes)
		align := p.Properties.Alignment
		if align == AlignJustify && isLastLine &&
			!forcedBreak {
			align = AlignLeft
		}

		hasTabs := false
		lineText := string(
			runes[startRune:bestBreakRune],
		)
		for _, r := range lineText {
			if r == '\t' {
				hasTabs = true
				break
			}
		}
		if hasTabs &&
			(align == AlignJustify || align == AlignDistribute) {
			align = AlignLeft
		}

		e.applyAlignment(
			&line,
			align,
			availableWidth,
		)
		lines = append(lines, line)
		startRune = bestBreakRune

		if !forcedBreak {
			for startRune < len(runes) && runes[startRune] == ' ' {
				startRune++
			}
		}
	}

	return lines
}

// reorderRuns rearranges positioned runs on a line based on visual order.
func (e *TextLayoutEngine) reorderRuns(
	line *LayoutLine,
	bp *BidiProcessor,
	startRune, endRune int,
) {
	if line == nil || len(line.Runs) == 0 {
		return
	}

	// For a simple implementation, we check if there are RTL runs.
	// If so, we may need to reorder them visually.
	// The bidi package handles the complex logic.
	_, err := bp.Reorder(startRune, endRune)
	if err != nil {
		return
	}

	// TODO: Fully implement run reordering based on bidi results.
	// This requires mapping visual bidi runs back to logical text runs.
}

// applyAlignment applies the paragraph alignment to a line.
func (e *TextLayoutEngine) applyAlignment(
	line *LayoutLine,
	align ParagraphAlignment,
	availableWidth float64,
) {
	if line == nil || len(line.Runs) == 0 {
		return
	}

	extraSpace := availableWidth - line.Width
	if extraSpace <= 0 {
		return
	}

	switch align {
	case AlignCenter:
		offset := extraSpace / 2.0
		for _, run := range line.Runs {
			run.X += offset
		}

	case AlignRight:
		for _, run := range line.Runs {
			run.X += extraSpace
		}

	case AlignJustify:
		e.applyJustification(line, availableWidth)

	case AlignDistribute:
		e.applyDistribution(line, availableWidth)

	case AlignLeft:
		fallthrough
	default:
	}
}

// applyJustification distributes extra space among the word spaces in a line.
func (e *TextLayoutEngine) applyJustification(
	line *LayoutLine,
	availableWidth float64,
) {
	spaceCount := 0
	for _, pr := range line.Runs {
		spaceCount += countSpaces(pr.Run.Text)
	}

	if spaceCount == 0 {
		return
	}

	extraSpace := availableWidth - line.Width
	spaceIncrement := extraSpace / float64(
		spaceCount,
	)

	currentOffset := 0.0
	for _, pr := range line.Runs {
		pr.X += currentOffset
		runSpaces := countSpaces(pr.Run.Text)
		if runSpaces > 0 {
			pr.Run.WordSpacing += spaceIncrement
			addedWidth := float64(
				runSpaces,
			) * spaceIncrement
			pr.Width += addedWidth
			currentOffset += addedWidth
		}
	}
	line.Width = availableWidth
}

// applyDistribution distributes extra space among all characters in a line.
func (e *TextLayoutEngine) applyDistribution(
	line *LayoutLine,
	availableWidth float64,
) {
	charCount := 0
	for _, pr := range line.Runs {
		charCount += len([]rune(pr.Run.Text))
	}

	if charCount <= 1 {
		return
	}

	extraSpace := availableWidth - line.Width
	gapIncrement := extraSpace / float64(
		charCount-1,
	)

	currentOffset := 0.0
	for _, pr := range line.Runs {
		pr.X += currentOffset
		runChars := len([]rune(pr.Run.Text))
		if runChars > 0 {
			pr.Run.CharacterSpacing += gapIncrement
			addedWidth := float64(
				runChars-1,
			) * gapIncrement
			pr.Width += addedWidth
			currentOffset += addedWidth + gapIncrement
		}
	}
	line.Width = availableWidth
}

// measureRuneRange measures the width of a range of runes within the paragraph.
func (e *TextLayoutEngine) measureRuneRange(
	p *Paragraph,
	start, end int,
	startX float64,
) float64 {
	if start >= end {
		return 0
	}

	currentX := startX
	currRune := 0

	for _, run := range p.Runs {
		runRunes := []rune(run.Text)
		runEnd := currRune + len(runRunes)

		if runEnd > start && currRune < end {
			overlapStart := max(0, start-currRune)
			overlapEnd := min(
				len(runRunes),
				end-currRune,
			)

			if overlapStart < overlapEnd {
				segmentRunes := runRunes[overlapStart:overlapEnd]
				lastRune := ' '
				if overlapStart > 0 {
					lastRune = runRunes[overlapStart-1]
				}

				for i, r := range segmentRunes {
					if r == '\t' {
						nextTab, _ := CalculateTabPosition(
							p,
							currentX,
						)
						currentX = nextTab
					} else {
						charRun := *run
						charRun.Text = string(r)
						width := MeasureRun(&charRun)

						if i > 0 || (overlapStart == 0 && currRune > start) {
							prevR := lastRune
							if i > 0 {
								prevR = segmentRunes[i-1]
							}
							if run.KerningTable != nil {
								upem := float64(run.GlyphMetrics.UnitsPerEm())
								kern := float64(run.KerningTable.GetKerning(prevR, r)) * run.FontSize / upem
								currentX += kern
							}
						}
						currentX += width
					}
				}
			}
		}

		currRune = runEnd
		if currRune >= end {
			break
		}
	}

	return currentX - startX
}

// createLineFromRange creates a LayoutLine for a specific rune range in the paragraph.
func (e *TextLayoutEngine) createLineFromRange(
	p *Paragraph,
	start, end int,
	availableWidth float64,
	forced bool,
) LayoutLine {
	line := LayoutLine{
		ForcedBreak: forced,
	}

	fullText := p.Text()
	runes := []rune(fullText)
	if start >= end || start >= len(runes) {
		return line
	}

	if end > len(runes) {
		end = len(runes)
	}

	currRune := 0
	currentX := 0.0

	for _, run := range p.Runs {
		runRunes := []rune(run.Text)
		runEnd := currRune + len(runRunes)

		if runEnd > start && currRune < end {
			overlapStart := max(0, start-currRune)
			overlapEnd := min(
				len(runRunes),
				end-currRune,
			)

			if overlapStart < overlapEnd {
				segmentRunes := runRunes[overlapStart:overlapEnd]
				var currentRunText strings.Builder
				runStartX := currentX

				for i, r := range segmentRunes {
					if r == '\t' {
						if currentRunText.Len() > 0 {
							segmentRun := *run
							segmentRun.Text = currentRunText.String()
							width := currentX - runStartX
							line.Runs = append(
								line.Runs,
								&PositionedRun{
									Run:   &segmentRun,
									X:     runStartX,
									Width: width,
								},
							)
							currentRunText.Reset()
						}
						nextTab, _ := CalculateTabPosition(
							p,
							currentX,
						)
						currentX = nextTab
						runStartX = currentX
					} else {
						charRun := *run
						charRun.Text = string(r)
						width := MeasureRun(&charRun)

						if i > 0 || (overlapStart == 0 && currRune > start) {
							prevR := ' '
							if i > 0 {
								prevR = segmentRunes[i-1]
							} else if overlapStart > 0 {
								prevR = runRunes[overlapStart-1]
							}
							if run.KerningTable != nil {
								upem := float64(run.GlyphMetrics.UnitsPerEm())
								kern := float64(run.KerningTable.GetKerning(prevR, r)) * run.FontSize / upem
								currentX += kern
							}
						}

						currentRunText.WriteRune(r)
						currentX += width

						if run.GlyphMetrics != nil {
							fontSize := run.FontSize
							upem := float64(run.GlyphMetrics.UnitsPerEm())
							ascent := float64(run.GlyphMetrics.Ascender()) * fontSize / upem
							descent := float64(run.GlyphMetrics.Descender()) * fontSize / upem

							if ascent > line.Ascent {
								line.Ascent = ascent
							}
							if descent < line.Descent {
								line.Descent = descent
							}

							height := CalculateLineHeight(p, fontSize)
							if height > line.Height {
								line.Height = height
							}
						}
					}
				}

				if currentRunText.Len() > 0 {
					segmentRun := *run
					segmentRun.Text = currentRunText.String()
					width := currentX - runStartX
					line.Runs = append(
						line.Runs,
						&PositionedRun{
							Run:   &segmentRun,
							X:     runStartX,
							Width: width,
						},
					)
				}
			}
		}

		currRune = runEnd
		if currRune >= end {
			break
		}
	}

	line.Width = currentX
	return line
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
