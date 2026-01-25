// This file implements the high-level Style API with a fluent builder pattern.
//
//nolint:revive // file-length-limit: style builder API is cohesive
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

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

// Build creates the style and registers it in the workbook's stylesheet.
// This method:
// 1. Gets or creates the workbook's styles part
// 2. Registers font, fill, border, and number format in the stylesheet
// 3. Creates a cellXf entry that references these elements
// 4. Returns the Style with its assigned index
//
// The returned Style can be applied to cells using its index.
func (sb *StyleBuilder) Build() *Style {
	// Get the workbook part
	wp := sb.doc.WorkbookPart()
	if wp == nil {
		// Cannot register style without workbook part
		return sb.style
	}

	// Get or create the styles part
	stylesPart := wp.StylesPart()
	if stylesPart == nil {
		var err error
		stylesPart, err = wp.AddStylesPart()
		if err != nil {
			// Failed to create styles part
			return sb.style
		}
	}

	// Get the stylesheet root element
	stylesheet := stylesPart.Stylesheet()
	if stylesheet == nil {
		// No stylesheet element available
		return sb.style
	}

	// Register the style components and get their IDs
	fontID := sb.registerFont(stylesheet)
	fillID := sb.registerFill(stylesheet)
	borderID := sb.registerBorder(stylesheet)
	numFmtID := sb.registerNumberFormat(
		stylesheet,
	)

	// Create the cellXf entry
	cellXfs := stylesheet.GetOrCreateCellXfs()
	xf := cellXfs.AddXf()

	// Set the component references
	xf.SetFontId(fontID)
	xf.SetFillId(fillID)
	xf.SetBorderId(borderID)
	xf.SetNumFmtId(numFmtID)
	xf.SetXfId(
		0,
	) // Reference to default cellStyleXf

	// Set apply flags for components that differ from defaults
	if fontID > 0 || sb.hasFontFormatting() {
		xf.SetApplyFont(true)
	}
	if fillID > 0 || sb.hasFillFormatting() {
		xf.SetApplyFill(true)
	}
	if borderID > 0 || sb.hasBorderFormatting() {
		xf.SetApplyBorder(true)
	}
	if numFmtID > 0 {
		xf.SetApplyNumberFormat(true)
	}

	// Register alignment if specified
	if sb.hasAlignmentFormatting() {
		sb.registerAlignment(xf)
		xf.SetApplyAlignment(true)
	}

	// Register protection if specified
	if sb.hasProtectionFormatting() {
		sb.registerProtection(xf)
		xf.SetApplyProtection(true)
	}

	// The index of this style is the count before we added it (zero-based)
	sb.style.index = cellXfs.Count() - 1

	return sb.style
}

// registerFont registers the font in the stylesheet and returns its index.
func (sb *StyleBuilder) registerFont(
	stylesheet *elements.Stylesheet,
) uint32 {
	fonts := stylesheet.GetOrCreateFonts()

	// Check if a matching font already exists
	var i uint32
	for font := range fonts.Fonts() {
		if sb.fontMatches(font) {
			return i
		}
		i++
	}

	// Create new font
	font := fonts.AddFont()

	if sb.style.fontName != "" {
		font.SetFontName(sb.style.fontName)
	}
	if sb.style.fontSize > 0 {
		font.SetFontSize(sb.style.fontSize)
	}
	if sb.style.fontBold {
		font.SetBold(true)
	}
	if sb.style.fontItalic {
		font.SetItalic(true)
	}
	if sb.style.fontUnderline {
		font.SetUnderline(
			elements.UnderlineStyleSingle,
		)
	}
	if sb.style.fontStrike {
		font.SetStrikethrough(true)
	}
	if sb.style.fontColor != "" {
		color := font.GetOrCreateColor()
		color.SetRGB(sb.style.fontColor)
	}

	return fonts.Count() - 1
}

// registerFill registers the fill in the stylesheet and returns its index.
func (sb *StyleBuilder) registerFill(
	stylesheet *elements.Stylesheet,
) uint32 {
	fills := stylesheet.GetOrCreateFills()

	// Check if a matching fill already exists
	var i uint32
	for fill := range fills.Fills() {
		if sb.fillMatches(fill) {
			return i
		}
		i++
	}

	// Create new fill
	fill := fills.AddFill()

	if sb.style.fillGradient {
		gf := fill.GetOrCreateGradientFill()
		if sb.style.gradientType == GradientLinear {
			gf.SetType(
				elements.GradientTypeLinear,
			)
		} else {
			gf.SetType(elements.GradientTypePath)
		}
		gf.SetDegree(sb.style.gradientDegree)

		// Add gradient stops
		for _, stop := range sb.style.gradientStops {
			gs := gf.AddStop(stop.Position)
			color := gs.GetOrCreateColor()
			color.SetRGB(stop.Color)
		}
	} else {
		pf := fill.GetOrCreatePatternFill()

		// Map style pattern type to element pattern type
		var patternType elements.PatternType
		switch sb.style.fillPattern {
		case PatternNone:
			patternType = elements.PatternTypeNone
		case PatternSolid:
			patternType = elements.PatternTypeSolid
		case PatternGray125:
			patternType = elements.PatternTypeGray125
		case PatternGray0625:
			patternType = elements.PatternTypeGray0625
		case PatternDarkGray:
			patternType = elements.PatternTypeDarkGray
		case PatternMediumGray:
			patternType = elements.PatternTypeMediumGray
		case PatternLightGray:
			patternType = elements.PatternTypeLightGray
		default:
			patternType = elements.PatternTypeNone
		}
		pf.SetPatternType(patternType)

		if sb.style.fillFgColor != "" {
			fg := pf.GetOrCreateFgColor()
			fg.SetRGB(sb.style.fillFgColor)
		}
		if sb.style.fillBgColor != "" {
			bg := pf.GetOrCreateBgColor()
			bg.SetRGB(sb.style.fillBgColor)
		}
	}

	return fills.Count() - 1
}

// registerBorder registers the border in the stylesheet and returns its index.
func (sb *StyleBuilder) registerBorder(
	stylesheet *elements.Stylesheet,
) uint32 {
	borders := stylesheet.GetOrCreateBorders()

	// Check if a matching border already exists
	var i uint32
	for border := range borders.Borders() {
		if sb.borderMatches(border) {
			return i
		}
		i++
	}

	// Create new border
	border := borders.AddBorder()

	// Map style border types to element border types
	mapBorderStyle := func(style BorderStyleType) elements.BorderStyle {
		switch style {
		case BorderNone:
			return elements.BorderStyleNone
		case BorderThin:
			return elements.BorderStyleThin
		case BorderMedium:
			return elements.BorderStyleMedium
		case BorderDashed:
			return elements.BorderStyleDashed
		case BorderDotted:
			return elements.BorderStyleDotted
		case BorderThick:
			return elements.BorderStyleThick
		case BorderDouble:
			return elements.BorderStyleDouble
		case BorderHair:
			return elements.BorderStyleHair
		case BorderMediumDashed:
			return elements.BorderStyleMediumDashed
		case BorderDashDot:
			return elements.BorderStyleDashDot
		case BorderMediumDashDot:
			return elements.BorderStyleMediumDashDot
		case BorderDashDotDot:
			return elements.BorderStyleDashDotDot
		case BorderMediumDashDotDot:
			return elements.BorderStyleMediumDashDotDot
		case BorderSlantDashDot:
			return elements.BorderStyleSlantDashDot
		default:
			return elements.BorderStyleNone
		}
	}

	if sb.style.borderLeft != BorderNone {
		left := border.GetOrCreateLeft()
		left.SetStyle(
			mapBorderStyle(sb.style.borderLeft),
		)
		if sb.style.borderLeftColor != "" {
			color := left.GetOrCreateColor()
			color.SetRGB(sb.style.borderLeftColor)
		}
	}
	if sb.style.borderRight != BorderNone {
		right := border.GetOrCreateRight()
		right.SetStyle(
			mapBorderStyle(sb.style.borderRight),
		)
		if sb.style.borderRightColor != "" {
			color := right.GetOrCreateColor()
			color.SetRGB(
				sb.style.borderRightColor,
			)
		}
	}
	if sb.style.borderTop != BorderNone {
		top := border.GetOrCreateTop()
		top.SetStyle(
			mapBorderStyle(sb.style.borderTop),
		)
		if sb.style.borderTopColor != "" {
			color := top.GetOrCreateColor()
			color.SetRGB(sb.style.borderTopColor)
		}
	}
	if sb.style.borderBottom != BorderNone {
		bottom := border.GetOrCreateBottom()
		bottom.SetStyle(
			mapBorderStyle(sb.style.borderBottom),
		)
		if sb.style.borderBottomColor != "" {
			color := bottom.GetOrCreateColor()
			color.SetRGB(
				sb.style.borderBottomColor,
			)
		}
	}
	if sb.style.borderDiag != BorderNone {
		diag := border.GetOrCreateDiagonal()
		diag.SetStyle(
			mapBorderStyle(sb.style.borderDiag),
		)
		if sb.style.borderDiagColor != "" {
			color := diag.GetOrCreateColor()
			color.SetRGB(sb.style.borderDiagColor)
		}
		border.SetDiagonalUp(
			sb.style.borderDiagUp,
		)
		border.SetDiagonalDown(
			sb.style.borderDiagDown,
		)
	}

	return borders.Count() - 1
}

// registerNumberFormat registers the number format and returns its ID.
func (sb *StyleBuilder) registerNumberFormat(
	stylesheet *elements.Stylesheet,
) uint32 {
	// If a format ID was explicitly set, use it
	if sb.style.numberFormatID > 0 {
		return sb.style.numberFormatID
	}

	// If no custom format string, return 0 (General)
	if sb.style.numberFormat == "" {
		return 0
	}

	// Check if it's a built-in format
	for id, code := range elements.BuiltInNumFmtCodes {
		if code == sb.style.numberFormat {
			return id
		}
	}

	// Need to create a custom number format
	numFmts := stylesheet.GetOrCreateNumFmts()

	// Check if this format code already exists
	for numFmt := range numFmts.NumFmts() {
		if numFmt.FormatCode() == sb.style.numberFormat {
			return numFmt.NumFmtId()
		}
	}

	// Find next available custom format ID (starting from 164)
	nextID := uint32(164)
	for numFmt := range numFmts.NumFmts() {
		if numFmt.NumFmtId() >= nextID {
			nextID = numFmt.NumFmtId() + 1
		}
	}

	// Add the custom format
	numFmts.AddNumFmt(
		nextID,
		sb.style.numberFormat,
	)

	return nextID
}

// registerAlignment registers alignment settings in the cellXf.
func (sb *StyleBuilder) registerAlignment(
	xf *elements.Xf,
) {
	alignment := xf.GetOrCreateAlignment()

	// Map style alignment types to element alignment types
	if sb.style.horizontal != "" {
		var h elements.HorizontalAlignment
		switch sb.style.horizontal {
		case HorizontalGeneral:
			h = elements.HorizontalAlignmentGeneral
		case HorizontalLeft:
			h = elements.HorizontalAlignmentLeft
		case HorizontalCenter:
			h = elements.HorizontalAlignmentCenter
		case HorizontalRight:
			h = elements.HorizontalAlignmentRight
		case HorizontalFill:
			h = elements.HorizontalAlignmentFill
		case HorizontalJustify:
			h = elements.HorizontalAlignmentJustify
		case HorizontalCenterContinuous:
			h = elements.HorizontalAlignmentCenterContinuous
		case HorizontalDistributed:
			h = elements.HorizontalAlignmentDistributed
		default:
			h = elements.HorizontalAlignmentGeneral
		}
		alignment.SetHorizontal(h)
	}

	if sb.style.vertical != "" {
		var v elements.VerticalAlignment
		switch sb.style.vertical {
		case VerticalTop:
			v = elements.VerticalAlignmentTop
		case VerticalCenter:
			v = elements.VerticalAlignmentCenter
		case VerticalBottom:
			v = elements.VerticalAlignmentBottom
		case VerticalJustify:
			v = elements.VerticalAlignmentJustify
		case VerticalDistributed:
			v = elements.VerticalAlignmentDistributed
		}
		alignment.SetVertical(v)
	}

	if sb.style.wrapText {
		alignment.SetWrapText(true)
	}
	if sb.style.shrinkToFit {
		alignment.SetShrinkToFit(true)
	}
	if sb.style.textRotation != 0 {
		alignment.SetTextRotation(
			uint32(sb.style.textRotation),
		)
	}
	if sb.style.indent > 0 {
		alignment.SetIndent(sb.style.indent)
	}
}

// registerProtection registers protection settings in the cellXf.
func (sb *StyleBuilder) registerProtection(
	xf *elements.Xf,
) {
	protection := xf.GetOrCreateProtection()
	protection.SetLocked(sb.style.locked)
	protection.SetHidden(sb.style.hidden)
}

// Helper methods to check if formatting is specified

func (sb *StyleBuilder) hasFontFormatting() bool {
	return sb.style.fontName != "" ||
		sb.style.fontSize > 0 ||
		sb.style.fontBold ||
		sb.style.fontItalic ||
		sb.style.fontUnderline ||
		sb.style.fontStrike ||
		sb.style.fontColor != ""
}

func (sb *StyleBuilder) hasFillFormatting() bool {
	return sb.style.fillPattern != PatternNone ||
		sb.style.fillGradient
}

func (sb *StyleBuilder) hasBorderFormatting() bool {
	return sb.style.borderLeft != BorderNone ||
		sb.style.borderRight != BorderNone ||
		sb.style.borderTop != BorderNone ||
		sb.style.borderBottom != BorderNone ||
		sb.style.borderDiag != BorderNone
}

func (sb *StyleBuilder) hasAlignmentFormatting() bool {
	return sb.style.horizontal != "" ||
		sb.style.vertical != "" ||
		sb.style.wrapText ||
		sb.style.shrinkToFit ||
		sb.style.textRotation != 0 ||
		sb.style.indent > 0
}

func (sb *StyleBuilder) hasProtectionFormatting() bool {
	// Protection has defaults (locked=true, hidden=false),
	// so we register it if either differs from defaults
	return !sb.style.locked || sb.style.hidden
}

// Helper methods to check if elements match existing ones

func (sb *StyleBuilder) fontMatches(
	font *elements.Font,
) bool {
	if sb.style.fontName != "" &&
		font.FontNameValue() != sb.style.fontName {
		return false
	}
	if sb.style.fontSize > 0 &&
		font.FontSizeValue() != sb.style.fontSize {
		return false
	}
	if sb.style.fontBold != font.IsBold() {
		return false
	}
	if sb.style.fontItalic != font.IsItalic() {
		return false
	}
	if sb.style.fontUnderline != (font.UnderlineStyle() == elements.UnderlineStyleSingle) {
		return false
	}
	if sb.style.fontStrike != font.IsStrikethrough() {
		return false
	}
	if sb.style.fontColor != "" {
		color := font.Color()
		if color == nil ||
			color.RGB() != sb.style.fontColor {
			return false
		}
	}

	return true
}

func (sb *StyleBuilder) fillMatches(
	fill *elements.Fill,
) bool {
	if sb.style.fillGradient {
		gf := fill.GradientFill()
		if gf == nil {
			return false
		}
		// Compare gradient properties
		if sb.style.gradientType == GradientLinear &&
			gf.Type() != elements.GradientTypeLinear {
			return false
		}
		if sb.style.gradientType == GradientPath &&
			gf.Type() != elements.GradientTypePath {
			return false
		}
		if gf.Degree() != sb.style.gradientDegree {
			return false
		}
		// For simplicity, if we have gradient stops, consider it non-matching
		// (full stop comparison would be complex)
		if len(sb.style.gradientStops) > 0 {
			return false
		}

		return true
	}

	pf := fill.PatternFill()
	if pf == nil {
		return false
	}

	// Map and compare pattern types
	var patternType elements.PatternType
	switch sb.style.fillPattern {
	case PatternNone:
		patternType = elements.PatternTypeNone
	case PatternSolid:
		patternType = elements.PatternTypeSolid
	case PatternGray125:
		patternType = elements.PatternTypeGray125
	case PatternGray0625:
		patternType = elements.PatternTypeGray0625
	case PatternDarkGray:
		patternType = elements.PatternTypeDarkGray
	case PatternMediumGray:
		patternType = elements.PatternTypeMediumGray
	case PatternLightGray:
		patternType = elements.PatternTypeLightGray
	default:
		patternType = elements.PatternTypeNone
	}

	if pf.PatternType() != patternType {
		return false
	}

	if sb.style.fillFgColor != "" {
		fg := pf.FgColor()
		if fg == nil ||
			fg.RGB() != sb.style.fillFgColor {
			return false
		}
	}
	if sb.style.fillBgColor != "" {
		bg := pf.BgColor()
		if bg == nil ||
			bg.RGB() != sb.style.fillBgColor {
			return false
		}
	}

	return true
}

func (sb *StyleBuilder) borderMatches(
	border *elements.Border,
) bool {
	// Helper to compare border side
	compareSide := func(
		styleType BorderStyleType,
		styleColor string,
		borderPr *elements.BorderPr,
	) bool {
		if styleType == BorderNone {
			return borderPr == nil ||
				borderPr.Style() == elements.BorderStyleNone
		}
		if borderPr == nil {
			return false
		}

		// Map style type to element type
		var elemStyle elements.BorderStyle
		switch styleType {
		case BorderNone:
			elemStyle = elements.BorderStyleNone
		case BorderThin:
			elemStyle = elements.BorderStyleThin
		case BorderMedium:
			elemStyle = elements.BorderStyleMedium
		case BorderDashed:
			elemStyle = elements.BorderStyleDashed
		case BorderDotted:
			elemStyle = elements.BorderStyleDotted
		case BorderThick:
			elemStyle = elements.BorderStyleThick
		case BorderDouble:
			elemStyle = elements.BorderStyleDouble
		case BorderHair:
			elemStyle = elements.BorderStyleHair
		case BorderMediumDashed:
			elemStyle = elements.BorderStyleMediumDashed
		case BorderDashDot:
			elemStyle = elements.BorderStyleDashDot
		case BorderMediumDashDot:
			elemStyle = elements.BorderStyleMediumDashDot
		case BorderDashDotDot:
			elemStyle = elements.BorderStyleDashDotDot
		case BorderMediumDashDotDot:
			elemStyle = elements.BorderStyleMediumDashDotDot
		case BorderSlantDashDot:
			elemStyle = elements.BorderStyleSlantDashDot
		default:
			elemStyle = elements.BorderStyleNone
		}

		if borderPr.Style() != elemStyle {
			return false
		}

		if styleColor != "" {
			color := borderPr.Color()
			if color == nil ||
				color.RGB() != styleColor {
				return false
			}
		}

		return true
	}

	if !compareSide(
		sb.style.borderLeft,
		sb.style.borderLeftColor,
		border.Left(),
	) {
		return false
	}
	if !compareSide(
		sb.style.borderRight,
		sb.style.borderRightColor,
		border.Right(),
	) {
		return false
	}
	if !compareSide(
		sb.style.borderTop,
		sb.style.borderTopColor,
		border.Top(),
	) {
		return false
	}
	if !compareSide(
		sb.style.borderBottom,
		sb.style.borderBottomColor,
		border.Bottom(),
	) {
		return false
	}
	if !compareSide(
		sb.style.borderDiag,
		sb.style.borderDiagColor,
		border.Diagonal(),
	) {
		return false
	}

	if sb.style.borderDiag != BorderNone {
		if border.DiagonalUp() != sb.style.borderDiagUp ||
			border.DiagonalDown() != sb.style.borderDiagDown {
			return false
		}
	}

	return true
}
