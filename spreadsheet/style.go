// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level Style API with a fluent builder pattern.
//
//nolint:revive // file-length-limit: style builder API is cohesive
package spreadsheet

// Note: imports are minimal since style registration is pending implementation

// Style represents a cell style with font, fill, border, alignment, etc.
type Style struct {
	// index is the style index in the workbook's styles part
	index uint32

	// font settings
	fontName      string
	fontSize      float64
	fontBold      bool
	fontItalic    bool
	fontUnderline bool
	fontStrike    bool
	fontColor     string

	// fill settings
	fillPattern    PatternType
	fillFgColor    string
	fillBgColor    string
	fillGradient   bool
	gradientType   GradientType
	gradientDegree float64
	gradientStops  []GradientStop

	// border settings
	borderLeft        BorderStyleType
	borderRight       BorderStyleType
	borderTop         BorderStyleType
	borderBottom      BorderStyleType
	borderDiag        BorderStyleType
	borderDiagUp      bool
	borderDiagDown    bool
	borderLeftColor   string
	borderRightColor  string
	borderTopColor    string
	borderBottomColor string
	borderDiagColor   string

	// number format
	numberFormat   string
	numberFormatID uint32

	// alignment
	horizontal   HorizontalAlign
	vertical     VerticalAlign
	wrapText     bool
	shrinkToFit  bool
	textRotation int
	indent       uint32

	// protection
	locked bool
	hidden bool
}

// PatternType represents a fill pattern type.
type PatternType string

const (
	// PatternNone indicates no fill pattern.
	PatternNone PatternType = "none"
	// PatternSolid indicates a solid fill.
	PatternSolid PatternType = "solid"
	// PatternGray125 indicates a 12.5% gray fill.
	PatternGray125 PatternType = "gray125"
	// PatternGray0625 indicates a 6.25% gray fill.
	PatternGray0625 PatternType = "gray0625"
	// PatternDarkGray indicates a dark gray fill.
	PatternDarkGray PatternType = "darkGray"
	// PatternMediumGray indicates a medium gray fill.
	PatternMediumGray PatternType = "mediumGray"
	// PatternLightGray indicates a light gray fill.
	PatternLightGray PatternType = "lightGray"
)

// GradientType represents a gradient fill type.
type GradientType string

const (
	// GradientLinear indicates a linear gradient.
	GradientLinear GradientType = "linear"
	// GradientPath indicates a path gradient.
	GradientPath GradientType = "path"
)

// GradientStop represents a color stop in a gradient.
type GradientStop struct {
	Position float64
	Color    string
}

// BorderStyleType represents a border style.
type BorderStyleType string

const (
	// BorderNone indicates no border.
	BorderNone BorderStyleType = "none"
	// BorderThin indicates a thin border.
	BorderThin BorderStyleType = "thin"
	// BorderMedium indicates a medium border.
	BorderMedium BorderStyleType = "medium"
	// BorderDashed indicates a dashed border.
	BorderDashed BorderStyleType = "dashed"
	// BorderDotted indicates a dotted border.
	BorderDotted BorderStyleType = "dotted"
	// BorderThick indicates a thick border.
	BorderThick BorderStyleType = "thick"
	// BorderDouble indicates a double border.
	BorderDouble BorderStyleType = "double"
	// BorderHair indicates a hair border.
	BorderHair BorderStyleType = "hair"
	// BorderMediumDashed indicates a medium dashed border.
	BorderMediumDashed BorderStyleType = "mediumDashed"
	// BorderDashDot indicates a dash-dot border.
	BorderDashDot BorderStyleType = "dashDot"
	// BorderMediumDashDot indicates a medium dash-dot border.
	BorderMediumDashDot BorderStyleType = "mediumDashDot"
	// BorderDashDotDot indicates a dash-dot-dot border.
	BorderDashDotDot BorderStyleType = "dashDotDot"
	// BorderMediumDashDotDot indicates a medium dash-dot-dot border.
	BorderMediumDashDotDot BorderStyleType = "mediumDashDotDot"
	// BorderSlantDashDot indicates a slant dash-dot border.
	BorderSlantDashDot BorderStyleType = "slantDashDot"
)

// HorizontalAlign represents horizontal alignment.
type HorizontalAlign string

const (
	// HorizontalGeneral indicates general alignment.
	HorizontalGeneral HorizontalAlign = "general"
	// HorizontalLeft indicates left alignment.
	HorizontalLeft HorizontalAlign = "left"
	// HorizontalCenter indicates center alignment.
	HorizontalCenter HorizontalAlign = "center"
	// HorizontalRight indicates right alignment.
	HorizontalRight HorizontalAlign = "right"
	// HorizontalFill indicates fill alignment.
	HorizontalFill HorizontalAlign = "fill"
	// HorizontalJustify indicates justified alignment.
	HorizontalJustify HorizontalAlign = "justify"
	// HorizontalCenterContinuous indicates center continuous alignment.
	HorizontalCenterContinuous HorizontalAlign = "centerContinuous"
	// HorizontalDistributed indicates distributed alignment.
	HorizontalDistributed HorizontalAlign = "distributed"
)

// VerticalAlign represents vertical alignment.
type VerticalAlign string

const (
	// VerticalTop indicates top alignment.
	VerticalTop VerticalAlign = "top"
	// VerticalCenter indicates center alignment.
	VerticalCenter VerticalAlign = "center"
	// VerticalBottom indicates bottom alignment.
	VerticalBottom VerticalAlign = "bottom"
	// VerticalJustify indicates justified alignment.
	VerticalJustify VerticalAlign = "justify"
	// VerticalDistributed indicates distributed alignment.
	VerticalDistributed VerticalAlign = "distributed"
)

// StyleBuilder provides a fluent API for building cell styles.
type StyleBuilder struct {
	doc   *Document
	style *Style
}

// NewStyle creates a new style builder.
func NewStyle(doc *Document) *StyleBuilder {
	return &StyleBuilder{
		doc: doc,
		style: &Style{
			fillPattern: PatternNone,
			locked:      true, // Cells are locked by default
		},
	}
}

// Font sets font properties for the style.
func (sb *StyleBuilder) Font(
	name string,
	size float64,
	bold, italic bool,
) *StyleBuilder {
	sb.style.fontName = name
	sb.style.fontSize = size
	sb.style.fontBold = bold
	sb.style.fontItalic = italic

	return sb
}

// FontName sets the font name.
func (sb *StyleBuilder) FontName(
	name string,
) *StyleBuilder {
	sb.style.fontName = name

	return sb
}

// FontSize sets the font size in points.
func (sb *StyleBuilder) FontSize(
	size float64,
) *StyleBuilder {
	sb.style.fontSize = size

	return sb
}

// Bold sets bold formatting.
func (sb *StyleBuilder) Bold(
	bold bool,
) *StyleBuilder {
	sb.style.fontBold = bold

	return sb
}

// Italic sets italic formatting.
func (sb *StyleBuilder) Italic(
	italic bool,
) *StyleBuilder {
	sb.style.fontItalic = italic

	return sb
}

// Underline sets underline formatting.
func (sb *StyleBuilder) Underline(
	underline bool,
) *StyleBuilder {
	sb.style.fontUnderline = underline

	return sb
}

// Strike sets strikethrough formatting.
func (sb *StyleBuilder) Strike(
	strike bool,
) *StyleBuilder {
	sb.style.fontStrike = strike

	return sb
}

// FontColor sets the font color (ARGB hex string, e.g., "FF000000" for black).
func (sb *StyleBuilder) FontColor(
	color string,
) *StyleBuilder {
	sb.style.fontColor = color

	return sb
}

// Fill sets fill properties for the style.
func (sb *StyleBuilder) Fill(
	pattern PatternType,
	fgColor, bgColor string,
) *StyleBuilder {
	sb.style.fillPattern = pattern
	sb.style.fillFgColor = fgColor
	sb.style.fillBgColor = bgColor
	sb.style.fillGradient = false

	return sb
}

// SolidFill sets a solid fill with the specified color (ARGB hex string).
func (sb *StyleBuilder) SolidFill(
	color string,
) *StyleBuilder {
	sb.style.fillPattern = PatternSolid
	sb.style.fillFgColor = color
	sb.style.fillGradient = false

	return sb
}

// GradientFill sets a gradient fill.
func (sb *StyleBuilder) GradientFill(
	gradType GradientType,
	degree float64,
	stops []GradientStop,
) *StyleBuilder {
	sb.style.fillGradient = true
	sb.style.gradientType = gradType
	sb.style.gradientDegree = degree
	sb.style.gradientStops = stops

	return sb
}

// Border sets border properties for all sides.
func (sb *StyleBuilder) Border(
	style BorderStyleType,
	color string,
) *StyleBuilder {
	sb.style.borderLeft = style
	sb.style.borderRight = style
	sb.style.borderTop = style
	sb.style.borderBottom = style
	sb.style.borderLeftColor = color
	sb.style.borderRightColor = color
	sb.style.borderTopColor = color
	sb.style.borderBottomColor = color

	return sb
}

// BorderLeft sets the left border.
func (sb *StyleBuilder) BorderLeft(
	style BorderStyleType,
	color string,
) *StyleBuilder {
	sb.style.borderLeft = style
	sb.style.borderLeftColor = color

	return sb
}

// BorderRight sets the right border.
func (sb *StyleBuilder) BorderRight(
	style BorderStyleType,
	color string,
) *StyleBuilder {
	sb.style.borderRight = style
	sb.style.borderRightColor = color

	return sb
}

// BorderTop sets the top border.
func (sb *StyleBuilder) BorderTop(
	style BorderStyleType,
	color string,
) *StyleBuilder {
	sb.style.borderTop = style
	sb.style.borderTopColor = color

	return sb
}

// BorderBottom sets the bottom border.
func (sb *StyleBuilder) BorderBottom(
	style BorderStyleType,
	color string,
) *StyleBuilder {
	sb.style.borderBottom = style
	sb.style.borderBottomColor = color

	return sb
}

// BorderDiagonal sets diagonal borders.
func (sb *StyleBuilder) BorderDiagonal(
	style BorderStyleType,
	color string,
	up, down bool,
) *StyleBuilder {
	sb.style.borderDiag = style
	sb.style.borderDiagColor = color
	sb.style.borderDiagUp = up
	sb.style.borderDiagDown = down

	return sb
}

// NumberFormat sets the number format string.
// Common formats:
//   - "General" - General format
//   - "0" - Integer
//   - "0.00" - Two decimal places
//   - "#,##0" - Thousands separator
//   - "#,##0.00" - Thousands with decimals
//   - "0%" - Percentage
//   - "0.00%" - Percentage with decimals
//   - "yyyy-mm-dd" - Date
//   - "hh:mm:ss" - Time
//   - "$#,##0.00" - Currency
func (sb *StyleBuilder) NumberFormat(
	format string,
) *StyleBuilder {
	sb.style.numberFormat = format

	return sb
}

// NumberFormatID sets a built-in number format ID.
// Common IDs:
//
//	0 - General
//	1 - 0
//	2 - 0.00
//	3 - #,##0
//	4 - #,##0.00
//	9 - 0%
//	10 - 0.00%
//	11 - 0.00E+00
//	12 - # ?/?
//	13 - # ??/??
//	14 - mm-dd-yy
//	15 - d-mmm-yy
//	16 - d-mmm
//	17 - mmm-yy
//	18 - h:mm AM/PM
//	19 - h:mm:ss AM/PM
//	20 - h:mm
//	21 - h:mm:ss
//	22 - m/d/yy h:mm
func (sb *StyleBuilder) NumberFormatID(
	id uint32,
) *StyleBuilder {
	sb.style.numberFormatID = id

	return sb
}

// Alignment sets alignment properties.
func (sb *StyleBuilder) Alignment(
	horizontal HorizontalAlign,
	vertical VerticalAlign,
) *StyleBuilder {
	sb.style.horizontal = horizontal
	sb.style.vertical = vertical

	return sb
}

// HorizontalAlignment sets the horizontal alignment.
func (sb *StyleBuilder) HorizontalAlignment(
	align HorizontalAlign,
) *StyleBuilder {
	sb.style.horizontal = align

	return sb
}

// VerticalAlignment sets the vertical alignment.
func (sb *StyleBuilder) VerticalAlignment(
	align VerticalAlign,
) *StyleBuilder {
	sb.style.vertical = align

	return sb
}

// WrapText sets text wrapping.
func (sb *StyleBuilder) WrapText(
	wrap bool,
) *StyleBuilder {
	sb.style.wrapText = wrap

	return sb
}

// ShrinkToFit sets shrink-to-fit.
func (sb *StyleBuilder) ShrinkToFit(
	shrink bool,
) *StyleBuilder {
	sb.style.shrinkToFit = shrink

	return sb
}

// TextRotation sets the text rotation angle (-90 to 90 degrees).
// 255 indicates vertical text.
func (sb *StyleBuilder) TextRotation(
	degrees int,
) *StyleBuilder {
	sb.style.textRotation = degrees

	return sb
}

// Indent sets the indentation level.
func (sb *StyleBuilder) Indent(
	level uint32,
) *StyleBuilder {
	sb.style.indent = level

	return sb
}

// Protection sets cell protection properties.
func (sb *StyleBuilder) Protection(
	locked, hidden bool,
) *StyleBuilder {
	sb.style.locked = locked
	sb.style.hidden = hidden

	return sb
}

// Locked sets whether the cell is locked when the sheet is protected.
func (sb *StyleBuilder) Locked(
	locked bool,
) *StyleBuilder {
	sb.style.locked = locked

	return sb
}

// Hidden sets whether the formula is hidden when the sheet is protected.
func (sb *StyleBuilder) Hidden(
	hidden bool,
) *StyleBuilder {
	sb.style.hidden = hidden

	return sb
}

// Build creates the style and returns it.
// Note: Full style registration to the workbook's styles part is not yet
// implemented. The returned Style contains the specified formatting options
// but is not yet persisted to the document's styles.xml.
// TODO: Implement full style registration when element tree manipulation
// is complete.
func (sb *StyleBuilder) Build() *Style {
	// For now, just return the style with a placeholder index
	// Full implementation requires proper element tree manipulation
	// which is pending completion of the elements package integration
	return sb.style
}
