// text.go provides text rendering capabilities for PDF output.
// This includes text operators, font handling, and text styling.

package drawing

import (
	"fmt"
	"strings"

	"github.com/connerohnesorge/goffice-pdf/font"
)

// TextRenderingMode specifies how text should be rendered.
// This corresponds to the PDF Tr (text rendering mode) operator.
type TextRenderingMode int

const (
	// TextRenderFill fills the text (default).
	TextRenderFill TextRenderingMode = 0
	// TextRenderStroke strokes the text outlines.
	TextRenderStroke TextRenderingMode = 1
	// TextRenderFillStroke fills and then strokes the text.
	TextRenderFillStroke TextRenderingMode = 2
	// TextRenderInvisible renders text invisibly (for selection/search).
	TextRenderInvisible TextRenderingMode = 3
	// TextRenderFillClip fills text and adds to clipping path.
	TextRenderFillClip TextRenderingMode = 4
	// TextRenderStrokeClip strokes text and adds to clipping path.
	TextRenderStrokeClip TextRenderingMode = 5
	// TextRenderFillStrokeClip fills, strokes, and clips.
	TextRenderFillStrokeClip TextRenderingMode = 6
	// TextRenderClip adds text to clipping path only.
	TextRenderClip TextRenderingMode = 7
)

// String returns a human-readable name for the text rendering mode.
func (m TextRenderingMode) String() string {
	switch m {
	case TextRenderFill:
		return "fill"
	case TextRenderStroke:
		return "stroke"
	case TextRenderFillStroke:
		return "fill+stroke"
	case TextRenderInvisible:
		return "invisible"
	case TextRenderFillClip:
		return "fill+clip"
	case TextRenderStrokeClip:
		return "stroke+clip"
	case TextRenderFillStrokeClip:
		return "fill+stroke+clip"
	case TextRenderClip:
		return "clip"
	default:
		return fmt.Sprintf(
			"TextRenderingMode(%d)",
			m,
		)
	}
}

// TextStyle defines styling options for text rendering.
type TextStyle struct {
	// Font is the font to use for rendering. May be nil for default font.
	Font *font.Font
	// FontSize is the font size in points.
	FontSize float64
	// Color is the fill color for the text.
	Color Color
	// StrokeColor is the stroke color (for stroke rendering modes).
	StrokeColor Color
	// StrokeWidth is the line width for stroked text.
	StrokeWidth float64
	// CharacterSpacing is additional spacing between characters (in points).
	// Positive values increase spacing, negative values decrease it.
	CharacterSpacing float64
	// WordSpacing is additional spacing between words (in points).
	// Only affects space characters (ASCII 32).
	WordSpacing float64
	// HorizontalScaling is a percentage value for horizontal text scaling.
	// 100 is normal, 50 is half-width, 200 is double-width.
	HorizontalScaling float64
	// TextRise is vertical offset from the baseline (in points).
	// Positive values move text up, negative values move text down.
	TextRise float64
	// Leading is the distance between baselines (in points).
	// Used when moving to the next line with T* operator.
	Leading float64
	// RenderingMode specifies how text is painted.
	RenderingMode TextRenderingMode
}

// NewTextStyle creates a new text style with default values.
func NewTextStyle() *TextStyle {
	return &TextStyle{
		FontSize:          12.0,
		Color:             Black,
		StrokeColor:       Black,
		StrokeWidth:       1.0,
		HorizontalScaling: 100.0,
		Leading:           14.4, // 1.2 * 12pt default
		RenderingMode:     TextRenderFill,
	}
}

// WithFont sets the font and returns the style for chaining.
func (s *TextStyle) WithFont(
	f *font.Font,
) *TextStyle {
	s.Font = f

	return s
}

// WithFontSize sets the font size and returns the style for chaining.
func (s *TextStyle) WithFontSize(
	size float64,
) *TextStyle {
	s.FontSize = size

	return s
}

// WithColor sets the fill color and returns the style for chaining.
func (s *TextStyle) WithColor(
	c Color,
) *TextStyle {
	s.Color = c

	return s
}

// WithStrokeColor sets the stroke color and returns the style for chaining.
func (s *TextStyle) WithStrokeColor(
	c Color,
) *TextStyle {
	s.StrokeColor = c

	return s
}

// WithCharacterSpacing sets character spacing and returns the style for chaining.
func (s *TextStyle) WithCharacterSpacing(
	spacing float64,
) *TextStyle {
	s.CharacterSpacing = spacing

	return s
}

// WithWordSpacing sets word spacing and returns the style for chaining.
func (s *TextStyle) WithWordSpacing(
	spacing float64,
) *TextStyle {
	s.WordSpacing = spacing

	return s
}

// WithHorizontalScaling sets horizontal scaling and returns the style for chaining.
func (s *TextStyle) WithHorizontalScaling(
	scale float64,
) *TextStyle {
	s.HorizontalScaling = scale

	return s
}

// WithTextRise sets text rise and returns the style for chaining.
func (s *TextStyle) WithTextRise(
	rise float64,
) *TextStyle {
	s.TextRise = rise

	return s
}

// WithLeading sets the text leading and returns the style for chaining.
func (s *TextStyle) WithLeading(
	leading float64,
) *TextStyle {
	s.Leading = leading

	return s
}

// WithRenderingMode sets the rendering mode and returns the style for chaining.
func (s *TextStyle) WithRenderingMode(
	mode TextRenderingMode,
) *TextStyle {
	s.RenderingMode = mode

	return s
}

// Clone returns a deep copy of the text style.
func (s *TextStyle) Clone() *TextStyle {
	clone := *s

	return &clone
}

// FontRef represents a reference to a font resource in a PDF.
// This is used when the actual font name in the page resources is needed.
type FontRef struct {
	// Name is the resource name (e.g., "F1", "F2").
	Name string
	// Font is the underlying font data.
	Font *font.Font
	// Subset is the font subset if subsetting is used.
	Subset *font.SubsetFont
}

// NewFontRef creates a new font reference.
func NewFontRef(
	name string,
	f *font.Font,
) *FontRef {
	return &FontRef{
		Name: name,
		Font: f,
	}
}

// TextBuilder provides a fluent API for constructing PDF text content streams.
// It generates proper BT/ET blocks with text operators.
type TextBuilder struct {
	operators strings.Builder
	fontRef   *FontRef
	style     *TextStyle
	inText    bool // Track if we're inside a BT/ET block
	currentX  float64
	currentY  float64
}

// NewTextBuilder creates a new text builder.
func NewTextBuilder() *TextBuilder {
	return &TextBuilder{
		style: NewTextStyle(),
	}
}

// BeginText starts a text object (BT operator).
// All text operations must occur between BeginText and EndText.
func (tb *TextBuilder) BeginText() *TextBuilder {
	if !tb.inText {
		tb.operators.WriteString("BT\n")
		tb.inText = true
	}

	return tb
}

// EndText ends a text object (ET operator).
func (tb *TextBuilder) EndText() *TextBuilder {
	if tb.inText {
		tb.operators.WriteString("ET\n")
		tb.inText = false
	}

	return tb
}

// SetFont sets the font and size (Tf operator).
// The fontName should be the resource name (e.g., "/F1").
func (tb *TextBuilder) SetFont(
	fontName string,
	size float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s %s Tf\n",
			fontName,
			formatFloat(size),
		),
	)

	return tb
}

// SetFontRef sets the font using a FontRef.
func (tb *TextBuilder) SetFontRef(
	ref *FontRef,
	size float64,
) *TextBuilder {
	tb.fontRef = ref

	return tb.SetFont("/"+ref.Name, size)
}

// SetTextMatrix sets the text matrix (Tm operator).
// This positions text and can apply transformations.
// The matrix is [a b c d e f] where (e, f) is the position.
func (tb *TextBuilder) SetTextMatrix(
	a, b, c, d, e, f float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf("%s %s %s %s %s %s Tm\n",
			formatFloat(a), formatFloat(b),
			formatFloat(c), formatFloat(d),
			formatFloat(e), formatFloat(f)),
	)
	tb.currentX = e
	tb.currentY = f

	return tb
}

// SetTextPosition sets the text position using identity matrix (simple positioning).
// This is equivalent to SetTextMatrix(1, 0, 0, 1, x, y).
func (tb *TextBuilder) SetTextPosition(
	x, y float64,
) *TextBuilder {
	return tb.SetTextMatrix(1, 0, 0, 1, x, y)
}

// MoveText moves the text position (Td operator).
// The offset is relative to the start of the current line.
func (tb *TextBuilder) MoveText(
	tx, ty float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s %s Td\n",
			formatFloat(tx),
			formatFloat(ty),
		),
	)
	tb.currentX += tx
	tb.currentY += ty

	return tb
}

// MoveTextSetLeading moves the text position and sets leading (TD operator).
// This is equivalent to: -ty TL tx ty Td
func (tb *TextBuilder) MoveTextSetLeading(
	tx, ty float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s %s TD\n",
			formatFloat(tx),
			formatFloat(ty),
		),
	)
	tb.currentX += tx
	tb.currentY += ty

	return tb
}

// NextLine moves to the start of the next line (T* operator).
// This uses the current leading value.
func (tb *TextBuilder) NextLine() *TextBuilder {
	tb.operators.WriteString("T*\n")

	return tb
}

// SetCharacterSpacing sets character spacing (Tc operator).
// Value is in unscaled text space units.
func (tb *TextBuilder) SetCharacterSpacing(
	spacing float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s Tc\n",
			formatFloat(spacing),
		),
	)

	return tb
}

// SetWordSpacing sets word spacing (Tw operator).
// Value is in unscaled text space units.
func (tb *TextBuilder) SetWordSpacing(
	spacing float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s Tw\n",
			formatFloat(spacing),
		),
	)

	return tb
}

// SetHorizontalScaling sets horizontal text scaling (Tz operator).
// Value is a percentage (100 = normal).
func (tb *TextBuilder) SetHorizontalScaling(
	scale float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s Tz\n",
			formatFloat(scale),
		),
	)

	return tb
}

// SetLeading sets the text leading (TL operator).
// Leading is the vertical distance between baselines.
func (tb *TextBuilder) SetLeading(
	leading float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s TL\n",
			formatFloat(leading),
		),
	)

	return tb
}

// SetTextRise sets the text rise (Ts operator).
// This moves text up (positive) or down (negative) from the baseline.
func (tb *TextBuilder) SetTextRise(
	rise float64,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf("%s Ts\n", formatFloat(rise)),
	)

	return tb
}

// SetRenderingMode sets the text rendering mode (Tr operator).
func (tb *TextBuilder) SetRenderingMode(
	mode TextRenderingMode,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf("%d Tr\n", mode),
	)

	return tb
}

// ShowText shows a simple text string (Tj operator).
// The string is automatically escaped for PDF.
func (tb *TextBuilder) ShowText(
	text string,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"(%s) Tj\n",
			escapePDFString(text),
		),
	)

	return tb
}

// ShowTextNextLine shows text and moves to next line (' operator).
// Equivalent to: T* (text) Tj
func (tb *TextBuilder) ShowTextNextLine(
	text string,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"(%s) '\n",
			escapePDFString(text),
		),
	)

	return tb
}

// ShowTextWithSpacing shows text with word and character spacing (" operator).
// Equivalent to: aw Tw ac Tc (text) '
func (tb *TextBuilder) ShowTextWithSpacing(
	wordSpacing, charSpacing float64,
	text string,
) *TextBuilder {
	tb.operators.WriteString(
		fmt.Sprintf(
			"%s %s (%s) \"\n",
			formatFloat(
				wordSpacing,
			),
			formatFloat(charSpacing),
			escapePDFString(text),
		),
	)

	return tb
}

// ShowTextArray shows text with individual positioning adjustments (TJ operator).
// Each element can be either a string or a number (negative for forward movement).
// This is used for kerning and precise text positioning.
func (tb *TextBuilder) ShowTextArray(
	elements ...interface{},
) *TextBuilder {
	tb.operators.WriteString("[")
	for i, elem := range elements {
		if i > 0 {
			tb.operators.WriteString(" ")
		}
		switch v := elem.(type) {
		case string:
			tb.operators.WriteString(fmt.Sprintf("(%s)", escapePDFString(v)))
		case float64:
			tb.operators.WriteString(formatFloat(v))
		case int:
			tb.operators.WriteString(fmt.Sprintf("%d", v))
		}
	}
	tb.operators.WriteString("] TJ\n")

	return tb
}

// ShowHexText shows text using hexadecimal encoding (Tj with hex string).
// This is useful for fonts with non-standard encodings or for Unicode text
// in CID fonts.
func (tb *TextBuilder) ShowHexText(
	data []byte,
) *TextBuilder {
	tb.operators.WriteString("<")
	for _, b := range data {
		tb.operators.WriteString(
			fmt.Sprintf("%02X", b),
		)
	}
	tb.operators.WriteString("> Tj\n")

	return tb
}

// ShowGlyphs shows text as a sequence of glyph IDs (for CID fonts).
// Each glyph ID is encoded as a 2-byte big-endian value.
func (tb *TextBuilder) ShowGlyphs(
	glyphIDs []uint16,
) *TextBuilder {
	tb.operators.WriteString("<")
	for _, gid := range glyphIDs {
		tb.operators.WriteString(
			fmt.Sprintf("%04X", gid),
		)
	}
	tb.operators.WriteString("> Tj\n")

	return tb
}

// ApplyStyle applies a TextStyle to the current text state.
func (tb *TextBuilder) ApplyStyle(
	style *TextStyle,
) *TextBuilder {
	if style == nil {
		return tb
	}

	// Apply character spacing if non-zero
	if style.CharacterSpacing != 0 {
		tb.SetCharacterSpacing(
			style.CharacterSpacing,
		)
	}

	// Apply word spacing if non-zero
	if style.WordSpacing != 0 {
		tb.SetWordSpacing(style.WordSpacing)
	}

	// Apply horizontal scaling if not 100%
	if style.HorizontalScaling != 100 &&
		style.HorizontalScaling != 0 {
		tb.SetHorizontalScaling(
			style.HorizontalScaling,
		)
	}

	// Apply text rise if non-zero
	if style.TextRise != 0 {
		tb.SetTextRise(style.TextRise)
	}

	// Apply leading if set
	if style.Leading != 0 {
		tb.SetLeading(style.Leading)
	}

	// Apply rendering mode if not fill (default)
	if style.RenderingMode != TextRenderFill {
		tb.SetRenderingMode(style.RenderingMode)
	}

	return tb
}

// String returns the accumulated text operators.
func (tb *TextBuilder) String() string {
	return tb.operators.String()
}

// Reset clears the builder for reuse.
func (tb *TextBuilder) Reset() *TextBuilder {
	tb.operators.Reset()
	tb.fontRef = nil
	tb.inText = false
	tb.currentX = 0
	tb.currentY = 0

	return tb
}

// DrawText is a convenience function that draws text at a position.
// It creates a complete BT/ET block with the specified styling.
func DrawText(
	x, y float64,
	text string,
	fontName string,
	style *TextStyle,
) string {
	if style == nil {
		style = NewTextStyle()
	}

	tb := NewTextBuilder()
	tb.BeginText()

	// Set fill color
	tb.operators.WriteString(
		style.Color.SetFillRGB() + "\n",
	)

	// Set stroke color if using stroke rendering
	if style.RenderingMode == TextRenderStroke ||
		style.RenderingMode == TextRenderFillStroke ||
		style.RenderingMode == TextRenderStrokeClip ||
		style.RenderingMode == TextRenderFillStrokeClip {
		tb.operators.WriteString(
			style.StrokeColor.SetStrokeRGB() + "\n",
		)
		tb.operators.WriteString(
			fmt.Sprintf(
				"%s w\n",
				formatFloat(style.StrokeWidth),
			),
		)
	}

	// Set font
	tb.SetFont(fontName, style.FontSize)

	// Apply other style properties
	tb.ApplyStyle(style)

	// Position and show text
	tb.SetTextPosition(x, y)
	tb.ShowText(text)

	tb.EndText()

	return tb.String()
}

// DrawTextAt is a simpler version of DrawText using default styling.
func DrawTextAt(
	x, y float64,
	text string,
	fontName string,
	fontSize float64,
) string {
	tb := NewTextBuilder()
	tb.BeginText()
	tb.SetFont(fontName, fontSize)
	tb.SetTextPosition(x, y)
	tb.ShowText(text)
	tb.EndText()

	return tb.String()
}

// TextWidth calculates the width of text in points using the given font and size.
// Returns 0 if the font is nil.
func TextWidth(
	text string,
	f *font.Font,
	fontSize float64,
) float64 {
	if f == nil {
		return 0
	}
	// Get width in font units
	widthInUnits := f.TextWidth(text)
	// Convert to points using font metrics
	return f.ScaledWidth(widthInUnits, fontSize)
}

// TextHeight calculates the height of text (ascender - descender) in points.
// Returns 0 if the font is nil.
func TextHeight(
	f *font.Font,
	fontSize float64,
) float64 {
	if f == nil {
		return fontSize // Fallback to font size
	}
	if f.Metrics.UnitsPerEm == 0 {
		return fontSize
	}
	// Height is ascender - descender (descender is typically negative)
	heightInUnits := f.Metrics.Ascender - f.Metrics.Descender

	return float64(
		heightInUnits,
	) * fontSize / float64(
		f.Metrics.UnitsPerEm,
	)
}

// TextAscender returns the ascender height in points.
func TextAscender(
	f *font.Font,
	fontSize float64,
) float64 {
	if f == nil || f.Metrics.UnitsPerEm == 0 {
		return fontSize * 0.8 // Approximate
	}

	return float64(
		f.Metrics.Ascender,
	) * fontSize / float64(
		f.Metrics.UnitsPerEm,
	)
}

// TextDescender returns the descender depth in points (typically negative).
func TextDescender(
	f *font.Font,
	fontSize float64,
) float64 {
	if f == nil || f.Metrics.UnitsPerEm == 0 {
		return fontSize * -0.2 // Approximate
	}

	return float64(
		f.Metrics.Descender,
	) * fontSize / float64(
		f.Metrics.UnitsPerEm,
	)
}

// SubsetTracker tracks which glyphs are used from a font for subsetting.
type SubsetTracker struct {
	builder *font.SubsetBuilder
	subset  *font.SubsetFont
}

// NewSubsetTracker creates a new subset tracker for the given font.
func NewSubsetTracker(
	f *font.Font,
) *SubsetTracker {
	return &SubsetTracker{
		builder: font.NewSubset(f),
	}
}

// AddText adds all characters in the text to the subset.
func (st *SubsetTracker) AddText(text string) {
	st.builder.AddString(text)
}

// AddRune adds a single character to the subset.
func (st *SubsetTracker) AddRune(r rune) {
	st.builder.AddRune(r)
}

// Build creates the font subset. Call this after adding all text.
func (st *SubsetTracker) Build() (*font.SubsetFont, error) {
	if st.subset != nil {
		return st.subset, nil
	}
	var err error
	st.subset, err = st.builder.Build()

	return st.subset, err
}

// Subset returns the built subset, or nil if Build hasn't been called.
func (st *SubsetTracker) Subset() *font.SubsetFont {
	return st.subset
}

// EncodeText encodes text using the subset glyph IDs.
// Must call Build() before using this method.
func (st *SubsetTracker) EncodeText(
	text string,
) []uint16 {
	if st.subset == nil {
		return nil
	}

	return st.subset.EncodedText(text)
}

// escapePDFString escapes special characters in a PDF string.
func escapePDFString(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for _, r := range s {
		switch r {
		case '(':
			sb.WriteString("\\(")
		case ')':
			sb.WriteString("\\)")
		case '\\':
			sb.WriteString("\\\\")
		case '\n':
			sb.WriteString("\\n")
		case '\r':
			sb.WriteString("\\r")
		case '\t':
			sb.WriteString("\\t")
		case '\b':
			sb.WriteString("\\b")
		case '\f':
			sb.WriteString("\\f")
		default:
			// Handle non-printable ASCII characters
			if r < 32 || r > 126 {
				// Use octal escape for control characters
				if r <= 0xFF {
					sb.WriteString(
						fmt.Sprintf("\\%03o", r),
					)
				} else {
					// For Unicode, just write the character
					// (PDF 1.6+ supports UTF-16BE strings with BOM)
					sb.WriteRune(r)
				}
			} else {
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}

// GenerateFontResourceName generates a unique font resource name.
// The prefix is typically "F" for regular fonts or "TT" for TrueType.
func GenerateFontResourceName(
	prefix string,
	index int,
) string {
	return fmt.Sprintf("%s%d", prefix, index)
}

// FontDescriptor contains metadata for embedding a font in PDF.
type FontDescriptor struct {
	// FontName is the PostScript name of the font.
	FontName string
	// FontFamily is the font family name.
	FontFamily string
	// Flags describes font characteristics (fixed-pitch, serif, etc.).
	Flags int
	// FontBBox is the font bounding box [llx lly urx ury].
	FontBBox [4]int
	// ItalicAngle is the italic angle in degrees.
	ItalicAngle float64
	// Ascent is the maximum height above the baseline.
	Ascent int
	// Descent is the maximum depth below the baseline (negative).
	Descent int
	// CapHeight is the height of capital letters.
	CapHeight int
	// XHeight is the height of lowercase letters.
	XHeight int
	// StemV is the dominant vertical stem width.
	StemV int
}

// FontDescriptorFromFont creates a FontDescriptor from a parsed font.
func FontDescriptorFromFont(
	f *font.Font,
) *FontDescriptor {
	if f == nil {
		return nil
	}

	// Calculate font flags
	flags := 0
	if f.Style == font.StyleItalic ||
		f.Style == font.StyleBoldItalic {
		flags |= 1 << 6 // Italic
	}
	// Assume non-symbolic fonts for now
	flags |= 1 << 5 // Nonsymbolic

	return &FontDescriptor{
		FontName:   f.Family,
		FontFamily: f.Family,
		Flags:      flags,
		FontBBox: [4]int{
			0,
			f.Metrics.Descender,
			1000,
			f.Metrics.Ascender,
		},
		ItalicAngle: 0, // Would need to extract from font
		Ascent:      f.Metrics.Ascender,
		Descent:     f.Metrics.Descender,
		CapHeight:   f.Metrics.CapHeight,
		XHeight:     f.Metrics.XHeight,
		StemV:       80, // Approximate; would need to calculate from font
	}
}

// TextLine represents a single line of text for layout purposes.
type TextLine struct {
	Text   string
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// WrapText wraps text to fit within a maximum width.
// Returns a slice of lines with their positions.
func WrapText(
	text string,
	f *font.Font,
	fontSize, maxWidth, leading float64,
) []TextLine {
	if f == nil || maxWidth <= 0 {
		return []TextLine{{Text: text}}
	}

	var lines []TextLine
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}

	spaceWidth := TextWidth(" ", f, fontSize)
	currentLine := words[0]
	currentWidth := TextWidth(
		currentLine,
		f,
		fontSize,
	)
	y := float64(0)

	for i := 1; i < len(words); i++ {
		word := words[i]
		wordWidth := TextWidth(word, f, fontSize)
		testWidth := currentWidth + spaceWidth + wordWidth

		if testWidth <= maxWidth {
			currentLine += " " + word
			currentWidth = testWidth
		} else {
			// Start new line
			lines = append(lines, TextLine{
				Text:   currentLine,
				X:      0,
				Y:      y,
				Width:  currentWidth,
				Height: TextHeight(f, fontSize),
			})
			y -= leading
			currentLine = word
			currentWidth = wordWidth
		}
	}

	// Add last line
	if currentLine != "" {
		lines = append(lines, TextLine{
			Text:   currentLine,
			X:      0,
			Y:      y,
			Width:  currentWidth,
			Height: TextHeight(f, fontSize),
		})
	}

	return lines
}
