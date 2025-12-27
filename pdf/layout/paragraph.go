// paragraph.go provides paragraph-level text measurement and layout for PDF generation.
// A paragraph (corresponding to <w:p> in Word documents) contains multiple runs with
// potentially different formatting. This file provides measurement functions needed
// for line breaking and layout decisions.

package layout

import (
	"strings"
	"unicode"
)

// ParagraphAlignment represents text alignment within a paragraph.
type ParagraphAlignment int

const (
	// AlignLeft aligns text to the left margin.
	AlignLeft ParagraphAlignment = iota
	// AlignCenter centers text between margins.
	AlignCenter
	// AlignRight aligns text to the right margin.
	AlignRight
	// AlignJustify distributes text evenly between margins.
	AlignJustify
	// AlignDistribute distributes text and whitespace evenly.
	AlignDistribute
)

// String returns the string representation of the alignment.
func (a ParagraphAlignment) String() string {
	switch a {
	case AlignLeft:
		return "left"
	case AlignCenter:
		return "center"
	case AlignRight:
		return "right"
	case AlignJustify:
		return "justify"
	case AlignDistribute:
		return "distribute"
	default:
		return "left"
	}
}

// LineSpacingRule defines how line spacing is calculated.
type LineSpacingRule int

const (
	// LineSpacingAuto uses automatic line spacing based on font metrics.
	LineSpacingAuto LineSpacingRule = iota
	// LineSpacingExact uses an exact line height value.
	LineSpacingExact
	// LineSpacingAtLeast uses a minimum line height.
	LineSpacingAtLeast
	// LineSpacingMultiple uses a multiple of the default line height.
	LineSpacingMultiple
)

// String returns the string representation of the line spacing rule.
func (r LineSpacingRule) String() string {
	switch r {
	case LineSpacingAuto:
		return "auto"
	case LineSpacingExact:
		return "exact"
	case LineSpacingAtLeast:
		return "atLeast"
	case LineSpacingMultiple:
		return "multiple"
	default:
		return "auto"
	}
}

// LineSpacing defines the spacing between lines in a paragraph.
type LineSpacing struct {
	// Rule specifies how the Value is interpreted.
	Rule LineSpacingRule

	// Value is the line spacing value.
	// For LineSpacingExact and LineSpacingAtLeast: value in points.
	// For LineSpacingMultiple: multiplier (e.g., 1.5 for 150% spacing).
	// For LineSpacingAuto: this value is ignored.
	Value float64

	// Before is additional spacing before the paragraph in points.
	Before float64

	// After is additional spacing after the paragraph in points.
	After float64
}

// DefaultLineSpacing returns the default line spacing (auto with no extra spacing).
func DefaultLineSpacing() LineSpacing {
	return LineSpacing{
		Rule:   LineSpacingAuto,
		Value:  0,
		Before: 0,
		After:  0,
	}
}

// ParagraphIndentation defines the indentation settings for a paragraph.
type ParagraphIndentation struct {
	// Left is the left indentation in points from the left margin.
	Left float64

	// Right is the right indentation in points from the right margin.
	Right float64

	// FirstLine is the additional indentation for the first line only.
	// Positive values indent the first line further right.
	// Use Hanging instead for hanging indents.
	FirstLine float64

	// Hanging is a hanging indent in points.
	// When set, subsequent lines are indented by this amount relative to the first line.
	// This is mutually exclusive with FirstLine (only one should be non-zero).
	Hanging float64
}

// EffectiveFirstLineIndent returns the effective first line indentation.
// This accounts for both first-line indent and hanging indent.
func (i ParagraphIndentation) EffectiveFirstLineIndent() float64 {
	if i.Hanging > 0 {
		// With hanging indent, first line is at Left position,
		// subsequent lines are at Left + Hanging
		return i.Left
	}
	// With first-line indent, first line is at Left + FirstLine
	return i.Left + i.FirstLine
}

// EffectiveSubsequentIndent returns the effective indentation for lines after the first.
func (i ParagraphIndentation) EffectiveSubsequentIndent() float64 {
	if i.Hanging > 0 {
		// With hanging indent, subsequent lines are indented further
		return i.Left + i.Hanging
	}
	// Normal case: subsequent lines use the base left indent
	return i.Left
}

// TabStop represents a tab stop position and type.
type TabStop struct {
	// Position is the tab stop position in points from the left margin.
	Position float64

	// Type specifies how text aligns at this tab stop.
	Type TabStopType

	// Leader is the character used to fill space before the tab stop.
	Leader TabLeader
}

// TabStopType defines the alignment at a tab stop.
type TabStopType int

const (
	// TabStopLeft aligns text to the left of the tab position.
	TabStopLeft TabStopType = iota
	// TabStopCenter centers text at the tab position.
	TabStopCenter
	// TabStopRight aligns text to the right of the tab position.
	TabStopRight
	// TabStopDecimal aligns the decimal point at the tab position.
	TabStopDecimal
	// TabStopBar draws a vertical bar at the tab position.
	TabStopBar
)

// TabLeader defines the leader character for a tab stop.
type TabLeader int

const (
	// TabLeaderNone uses no leader.
	TabLeaderNone TabLeader = iota
	// TabLeaderDot uses dots as a leader.
	TabLeaderDot
	// TabLeaderHyphen uses hyphens as a leader.
	TabLeaderHyphen
	// TabLeaderUnderscore uses underscores as a leader.
	TabLeaderUnderscore
	// TabLeaderMiddleDot uses middle dots as a leader.
	TabLeaderMiddleDot
)

// ParagraphProperties contains formatting properties for a paragraph.
type ParagraphProperties struct {
	// Alignment specifies text alignment.
	Alignment ParagraphAlignment

	// Indentation specifies paragraph indentation.
	Indentation ParagraphIndentation

	// LineSpacing specifies line spacing settings.
	LineSpacing LineSpacing

	// TabStops lists custom tab stop positions.
	TabStops []TabStop

	// KeepLines prevents page breaks within the paragraph.
	KeepLines bool

	// KeepWithNext prevents a page break between this paragraph and the next.
	KeepWithNext bool

	// PageBreakBefore forces a page break before this paragraph.
	PageBreakBefore bool

	// WidowControl prevents widow and orphan lines.
	WidowControl bool
}

// DefaultParagraphProperties returns the default paragraph properties.
func DefaultParagraphProperties() ParagraphProperties {
	return ParagraphProperties{
		Alignment:       AlignLeft,
		Indentation:     ParagraphIndentation{},
		LineSpacing:     DefaultLineSpacing(),
		TabStops:        nil,
		KeepLines:       false,
		KeepWithNext:    false,
		PageBreakBefore: false,
		WidowControl:    true,
	}
}

// Paragraph represents a paragraph containing multiple text runs.
// It corresponds to a <w:p> element in Word documents.
type Paragraph struct {
	// Runs contains the text runs in this paragraph.
	Runs []*TextRun

	// Properties contains the paragraph formatting properties.
	Properties ParagraphProperties
}

// NewParagraph creates a new empty paragraph with default properties.
func NewParagraph() *Paragraph {
	return &Paragraph{
		Runs:       nil,
		Properties: DefaultParagraphProperties(),
	}
}

// NewParagraphWithRuns creates a new paragraph with the given runs.
func NewParagraphWithRuns(
	runs []*TextRun,
) *Paragraph {
	return &Paragraph{
		Runs:       runs,
		Properties: DefaultParagraphProperties(),
	}
}

// AddRun adds a text run to the paragraph.
func (p *Paragraph) AddRun(run *TextRun) {
	if p != nil && run != nil {
		p.Runs = append(p.Runs, run)
	}
}

// Text returns the concatenated text of all runs in the paragraph.
func (p *Paragraph) Text() string {
	if p == nil || len(p.Runs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, run := range p.Runs {
		if run != nil {
			sb.WriteString(run.Text)
		}
	}

	return sb.String()
}

// IsEmpty returns true if the paragraph contains no text.
func (p *Paragraph) IsEmpty() bool {
	if p == nil {
		return true
	}

	for _, run := range p.Runs {
		if run != nil && len(run.Text) > 0 {
			return false
		}
	}

	return true
}

// RunCount returns the number of runs in the paragraph.
func (p *Paragraph) RunCount() int {
	if p == nil {
		return 0
	}

	return len(p.Runs)
}

// MeasureParagraph calculates the total natural width of a paragraph in points.
// This is the sum of all run widths without any line breaking applied.
func MeasureParagraph(p *Paragraph) float64 {
	if p == nil || len(p.Runs) == 0 {
		return 0
	}

	var totalWidth float64
	for _, run := range p.Runs {
		totalWidth += MeasureRun(run)
	}

	return totalWidth
}

// MeasureFirstLine calculates the width available for the first line of a paragraph.
// This accounts for first-line indentation settings.
// The containerWidth parameter is the total available width for the paragraph.
func MeasureFirstLine(
	p *Paragraph,
	containerWidth float64,
) float64 {
	if p == nil {
		return containerWidth
	}

	indent := p.Properties.Indentation
	leftIndent := indent.EffectiveFirstLineIndent()
	rightIndent := indent.Right

	available := containerWidth - leftIndent - rightIndent
	if available < 0 {
		return 0
	}

	return available
}

// MeasureSubsequentLines calculates the width available for subsequent lines.
// This accounts for hanging indentation settings.
func MeasureSubsequentLines(
	p *Paragraph,
	containerWidth float64,
) float64 {
	if p == nil {
		return containerWidth
	}

	indent := p.Properties.Indentation
	leftIndent := indent.EffectiveSubsequentIndent()
	rightIndent := indent.Right

	available := containerWidth - leftIndent - rightIndent
	if available < 0 {
		return 0
	}

	return available
}

// GetMinimumWidth returns the minimum width required to render the paragraph.
// This is the width of the longest word (or unbreakable unit) across all runs.
func GetMinimumWidth(p *Paragraph) float64 {
	if p == nil || len(p.Runs) == 0 {
		return 0
	}

	var maxWordWidth float64

	for _, run := range p.Runs {
		if run == nil || len(run.Text) == 0 {
			continue
		}

		// Find all words in this run
		words := splitIntoWords(run.Text)

		for _, word := range words {
			// Measure this word using the run's metrics
			wordRun := &TextRun{
				Text:              word,
				GlyphMetrics:      run.GlyphMetrics,
				KerningTable:      run.KerningTable,
				FontSize:          run.FontSize,
				CharacterSpacing:  run.CharacterSpacing,
				WordSpacing:       0, // Don't add word spacing within a word
				HorizontalScaling: run.HorizontalScaling,
			}

			wordWidth := MeasureRun(wordRun)
			if wordWidth > maxWordWidth {
				maxWordWidth = wordWidth
			}
		}
	}

	// Add indentation to the minimum width requirement
	if p.Properties.Indentation.Left > 0 ||
		p.Properties.Indentation.Hanging > 0 ||
		p.Properties.Indentation.FirstLine > 0 {
		// The minimum needs to accommodate the largest indent plus the longest word
		maxIndent := p.Properties.Indentation.EffectiveFirstLineIndent()
		subsequentIndent := p.Properties.Indentation.EffectiveSubsequentIndent()
		if subsequentIndent > maxIndent {
			maxIndent = subsequentIndent
		}
		maxWordWidth += maxIndent
	}

	return maxWordWidth
}

// GetMaximumWidth returns the maximum width of the paragraph without wrapping.
// This is the natural width of all content laid out on a single line.
func GetMaximumWidth(p *Paragraph) float64 {
	if p == nil || len(p.Runs) == 0 {
		return 0
	}

	// Base content width
	contentWidth := MeasureParagraph(p)

	// Add first-line indentation (since no wrapping means single line)
	contentWidth += p.Properties.Indentation.EffectiveFirstLineIndent()

	return contentWidth
}

// splitIntoWords splits text into words, treating whitespace as word boundaries.
// Returns a slice of non-empty word strings.
func splitIntoWords(text string) []string {
	if len(text) == 0 {
		return nil
	}

	var words []string
	var currentWord strings.Builder
	inWord := false

	for _, r := range text {
		if unicode.IsSpace(r) {
			if inWord {
				// End of word
				word := currentWord.String()
				if len(word) > 0 {
					words = append(words, word)
				}
				currentWord.Reset()
				inWord = false
			}
			// Skip whitespace
		} else {
			// Part of a word
			currentWord.WriteRune(r)
			inWord = true
		}
	}

	// Don't forget the last word
	if inWord {
		word := currentWord.String()
		if len(word) > 0 {
			words = append(words, word)
		}
	}

	return words
}

// ParagraphMeasurer provides a reusable context for measuring paragraphs.
// It can be used when measuring multiple paragraphs with shared settings.
type ParagraphMeasurer struct {
	// containerWidth is the available width for paragraph content.
	containerWidth float64

	// defaultFontSize is the fallback font size when runs don't specify one.
	defaultFontSize float64
}

// NewParagraphMeasurer creates a new ParagraphMeasurer with the given container width.
func NewParagraphMeasurer(
	containerWidth float64,
) *ParagraphMeasurer {
	return &ParagraphMeasurer{
		containerWidth:  containerWidth,
		defaultFontSize: 12.0, // Default 12pt
	}
}

// SetContainerWidth sets the container width for measurements.
func (pm *ParagraphMeasurer) SetContainerWidth(
	width float64,
) {
	if pm != nil {
		pm.containerWidth = width
	}
}

// ContainerWidth returns the current container width.
func (pm *ParagraphMeasurer) ContainerWidth() float64 {
	if pm == nil {
		return 0
	}

	return pm.containerWidth
}

// SetDefaultFontSize sets the default font size for runs without a specified size.
func (pm *ParagraphMeasurer) SetDefaultFontSize(
	size float64,
) {
	if pm != nil {
		pm.defaultFontSize = size
	}
}

// MeasureNaturalWidth returns the natural width of the paragraph content.
func (pm *ParagraphMeasurer) MeasureNaturalWidth(
	p *Paragraph,
) float64 {
	return MeasureParagraph(p)
}

// MeasureFirstLineAvailable returns the available width for the first line.
func (pm *ParagraphMeasurer) MeasureFirstLineAvailable(
	p *Paragraph,
) float64 {
	if pm == nil {
		return 0
	}

	return MeasureFirstLine(p, pm.containerWidth)
}

// MeasureSubsequentLinesAvailable returns the available width for subsequent lines.
func (pm *ParagraphMeasurer) MeasureSubsequentLinesAvailable(
	p *Paragraph,
) float64 {
	if pm == nil {
		return 0
	}

	return MeasureSubsequentLines(
		p,
		pm.containerWidth,
	)
}

// MeasureMinimumWidth returns the minimum width required for the paragraph.
func (pm *ParagraphMeasurer) MeasureMinimumWidth(
	p *Paragraph,
) float64 {
	return GetMinimumWidth(p)
}

// MeasureMaximumWidth returns the maximum width without wrapping.
func (pm *ParagraphMeasurer) MeasureMaximumWidth(
	p *Paragraph,
) float64 {
	return GetMaximumWidth(p)
}

// FitsWidth returns true if the paragraph fits within the container width
// without requiring line wrapping.
func (pm *ParagraphMeasurer) FitsWidth(
	p *Paragraph,
) bool {
	if pm == nil || p == nil {
		return true
	}

	// Check first line fit
	firstLineWidth := MeasureParagraph(p)
	firstLineAvailable := pm.MeasureFirstLineAvailable(
		p,
	)

	return firstLineWidth <= firstLineAvailable
}

// RequiresWrapping returns true if the paragraph needs line breaking.
func (pm *ParagraphMeasurer) RequiresWrapping(
	p *Paragraph,
) bool {
	return !pm.FitsWidth(p)
}

// ParagraphWidthInfo contains detailed width information for a paragraph.
type ParagraphWidthInfo struct {
	// NaturalWidth is the total width of all runs without wrapping.
	NaturalWidth float64

	// MinimumWidth is the width of the longest unbreakable unit.
	MinimumWidth float64

	// MaximumWidth is the width with indentation, without wrapping.
	MaximumWidth float64

	// FirstLineIndent is the effective first line indentation.
	FirstLineIndent float64

	// SubsequentIndent is the effective indentation for subsequent lines.
	SubsequentIndent float64

	// RunCount is the number of runs in the paragraph.
	RunCount int

	// CharacterCount is the total number of characters.
	CharacterCount int

	// WordCount is the approximate number of words.
	WordCount int
}

// MeasureParagraphDetailed returns detailed width information for a paragraph.
func MeasureParagraphDetailed(
	p *Paragraph,
) ParagraphWidthInfo {
	info := ParagraphWidthInfo{}

	if p == nil || len(p.Runs) == 0 {
		return info
	}

	info.NaturalWidth = MeasureParagraph(p)
	info.MinimumWidth = GetMinimumWidth(p)
	info.MaximumWidth = GetMaximumWidth(p)
	info.FirstLineIndent = p.Properties.Indentation.EffectiveFirstLineIndent()
	info.SubsequentIndent = p.Properties.Indentation.EffectiveSubsequentIndent()
	info.RunCount = len(p.Runs)

	// Count characters and words
	for _, run := range p.Runs {
		if run != nil {
			info.CharacterCount += len(
				[]rune(run.Text),
			)
			info.WordCount += len(
				splitIntoWords(run.Text),
			)
		}
	}

	return info
}

// CalculateTabPosition finds the next tab stop position from the current position.
// Returns the tab position and the type of tab stop.
// If no custom tab stop is found, returns a default tab stop position.
func CalculateTabPosition(
	p *Paragraph,
	currentPos float64,
) (float64, TabStopType) {
	if p == nil {
		return nextDefaultTabStop(
			currentPos,
		), TabStopLeft
	}

	// Look for custom tab stops
	for _, tab := range p.Properties.TabStops {
		if tab.Position > currentPos {
			return tab.Position, tab.Type
		}
	}

	// No custom tab stop found, use default
	return nextDefaultTabStop(
		currentPos,
	), TabStopLeft
}

// DefaultTabInterval is the default interval between tab stops (0.5 inches = 36 points).
const DefaultTabInterval = 36.0

// nextDefaultTabStop returns the next default tab stop position.
func nextDefaultTabStop(
	currentPos float64,
) float64 {
	// Find the next multiple of DefaultTabInterval
	tabNumber := int(
		currentPos/DefaultTabInterval,
	) + 1

	return float64(tabNumber) * DefaultTabInterval
}

// EstimateLineCount estimates the number of lines needed to render the paragraph
// within the given width. This is an approximation useful for layout planning.
func EstimateLineCount(
	p *Paragraph,
	availableWidth float64,
) int {
	if p == nil || len(p.Runs) == 0 ||
		availableWidth <= 0 {
		return 0
	}

	naturalWidth := MeasureParagraph(p)
	if naturalWidth == 0 {
		return 0
	}

	// Account for first line having potentially different width
	firstLineWidth := MeasureFirstLine(
		p,
		availableWidth,
	)
	subsequentWidth := MeasureSubsequentLines(
		p,
		availableWidth,
	)

	if firstLineWidth <= 0 ||
		subsequentWidth <= 0 {
		// Cannot fit any content
		return 0
	}

	// Simple estimation: content fits on first line?
	if naturalWidth <= firstLineWidth {
		return 1
	}

	// Need more lines
	remainingAfterFirst := naturalWidth - firstLineWidth
	additionalLines := int(
		remainingAfterFirst/subsequentWidth,
	) + 1

	return 1 + additionalLines
}

// CalculateLineHeight calculates the line height for a line in points.
// The tallestRun parameter is the tallest run on the line (by font size).
func CalculateLineHeight(
	p *Paragraph,
	tallestRunHeight float64,
) float64 {
	if p == nil {
		return tallestRunHeight * 1.2 // Default 120% line height
	}

	spacing := p.Properties.LineSpacing

	switch spacing.Rule {
	case LineSpacingExact:
		return spacing.Value

	case LineSpacingAtLeast:
		natural := tallestRunHeight * 1.2
		if spacing.Value > natural {
			return spacing.Value
		}

		return natural

	case LineSpacingMultiple:
		return tallestRunHeight * spacing.Value

	case LineSpacingAuto:
		fallthrough
	default:
		// Auto spacing: typically 115-120% of font height
		return tallestRunHeight * 1.15
	}
}
