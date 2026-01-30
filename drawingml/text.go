// and effects.
//
//nolint:revive // This file contains many public types for OOXML text elements.
package drawingml

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
	"golang.org/x/text/unicode/bidi"
)

// TextAlignValue represents text alignment.
type TextAlignValue string

// Text alignment values.
const (
	TextAlignLeft        TextAlignValue = "l"
	TextAlignCenter      TextAlignValue = "ctr"
	TextAlignRight       TextAlignValue = "r"
	TextAlignJustify     TextAlignValue = "just"
	TextAlignJustifyLow  TextAlignValue = "justLow"
	TextAlignDistributed TextAlignValue = "dist"
	TextAlignThaiDist    TextAlignValue = "thaiDist"
)

// TextAnchorValue represents vertical anchor (anchoring type).
type TextAnchorValue string

// Text anchor values.
const (
	TextAnchorTop         TextAnchorValue = "t"
	TextAnchorCenter      TextAnchorValue = "ctr"
	TextAnchorBottom      TextAnchorValue = "b"
	TextAnchorJustified   TextAnchorValue = "just"
	TextAnchorDistributed TextAnchorValue = "dist"
)

// TextVerticalValue represents vertical text orientation.
type TextVerticalValue string

// Text vertical values.
const (
	TextVerticalHorizontal TextVerticalValue = "horz"
	TextVerticalRotate90   TextVerticalValue = "vert"
	TextVerticalRotate270  TextVerticalValue = "vert270"
	TextVerticalWordArt    TextVerticalValue = "wordArtVert"
	TextVerticalEAVert     TextVerticalValue = "eaVert"
	TextVerticalMongolian  TextVerticalValue = "mongolianVert"
	TextVerticalWordArtRtl TextVerticalValue = "wordArtVertRtl"
)

// TextWrapValue represents text wrapping type.
type TextWrapValue string

// Text wrap values.
const (
	TextWrapNone   TextWrapValue = "none"
	TextWrapSquare TextWrapValue = "square"
)

// UnderlineValue represents underline style.
type UnderlineValue string

// Underline values.
const (
	UnderlineNone            UnderlineValue = "none"
	UnderlineSingle          UnderlineValue = "sng"
	UnderlineDouble          UnderlineValue = "dbl"
	UnderlineHeavy           UnderlineValue = "heavy"
	UnderlineDotted          UnderlineValue = "dotted"
	UnderlineDottedHeavy     UnderlineValue = "dottedHeavy"
	UnderlineDash            UnderlineValue = "dash"
	UnderlineDashHeavy       UnderlineValue = "dashHeavy"
	UnderlineDashLong        UnderlineValue = "dashLong"
	UnderlineDashLongHeavy   UnderlineValue = "dashLongHeavy"
	UnderlineDotDash         UnderlineValue = "dotDash"
	UnderlineDotDashHeavy    UnderlineValue = "dotDashHeavy"
	UnderlineDotDotDash      UnderlineValue = "dotDotDash"
	UnderlineDotDotDashHeavy UnderlineValue = "dotDotDashHeavy"
	UnderlineWavy            UnderlineValue = "wavy"
	UnderlineWavyHeavy       UnderlineValue = "wavyHeavy"
	UnderlineWavyDouble      UnderlineValue = "wavyDbl"
	UnderlineWords           UnderlineValue = "words"
)

// StrikeValue represents strikethrough style.
type StrikeValue string

// Strike values.
const (
	StrikeNoStrike     StrikeValue = "noStrike"
	StrikeSingleStrike StrikeValue = "sngStrike"
	StrikeDoubleStrike StrikeValue = "dblStrike"
)

// CapValue represents text capitalization.
type CapValue string

// Cap values.
const (
	CapNone  CapValue = "none"
	CapSmall CapValue = "small"
	CapAll   CapValue = "all"
)

// FontAlignValue represents font alignment for text.
type FontAlignValue string

// Font alignment values.
const (
	FontAlignAuto   FontAlignValue = "auto"
	FontAlignTop    FontAlignValue = "t"
	FontAlignCenter FontAlignValue = "ctr"
	FontAlignBase   FontAlignValue = "base"
	FontAlignBottom FontAlignValue = "b"
)

// AutoNumberSchemeValue represents bullet auto-numbering schemes.
type AutoNumberSchemeValue string

// Auto number scheme values.
const (
	AutoNumArabicParenBoth  AutoNumberSchemeValue = "arabicParenBoth"
	AutoNumArabicParenR     AutoNumberSchemeValue = "arabicParenR"
	AutoNumArabicPeriod     AutoNumberSchemeValue = "arabicPeriod"
	AutoNumArabicPlain      AutoNumberSchemeValue = "arabicPlain"
	AutoNumRomanLcParenBoth AutoNumberSchemeValue = "romanLcParenBoth"
	AutoNumRomanLcParenR    AutoNumberSchemeValue = "romanLcParenR"
	AutoNumRomanLcPeriod    AutoNumberSchemeValue = "romanLcPeriod"
	AutoNumRomanUcParenBoth AutoNumberSchemeValue = "romanUcParenBoth"
	AutoNumRomanUcParenR    AutoNumberSchemeValue = "romanUcParenR"
	AutoNumRomanUcPeriod    AutoNumberSchemeValue = "romanUcPeriod"
	AutoNumAlphaLcParenBoth AutoNumberSchemeValue = "alphaLcParenBoth"
	AutoNumAlphaLcParenR    AutoNumberSchemeValue = "alphaLcParenR"
	AutoNumAlphaLcPeriod    AutoNumberSchemeValue = "alphaLcPeriod"
	AutoNumAlphaUcParenBoth AutoNumberSchemeValue = "alphaUcParenBoth"
	AutoNumAlphaUcParenR    AutoNumberSchemeValue = "alphaUcParenR"
	AutoNumAlphaUcPeriod    AutoNumberSchemeValue = "alphaUcPeriod"
	AutoNumCircleNumDbPlain AutoNumberSchemeValue = "circleNumDbPlain"
	AutoNumCircleNumWdBlack AutoNumberSchemeValue = "circleNumWdBlackPlain"
	AutoNumCircleNumWdWhite AutoNumberSchemeValue = "circleNumWdWhitePlain"
)

// TabAlignValue represents tab stop alignment.
type TabAlignValue string

// Tab alignment values.
const (
	TabAlignLeft    TabAlignValue = "l"
	TabAlignCenter  TabAlignValue = "ctr"
	TabAlignRight   TabAlignValue = "r"
	TabAlignDecimal TabAlignValue = "dec"
)

// ===========================================================================
// TextBody (a:txBody) - Container for text in a shape
// ===========================================================================

// TextBody represents a text body element (a:txBody).
// This is the container for all text within a shape.
type TextBody struct {
	*openxml.CompositeElementBase
}

// NewTextBody creates a new text body element.
func NewTextBody() *TextBody {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"txBody",
		PrefixMain,
	)
	tb := &TextBody{CompositeElementBase: elem}
	// Add default body properties
	tb.AppendChild(NewTextBodyProperties())

	return tb
}

// BodyProperties returns the text body properties.
func (tb *TextBody) BodyProperties() *TextBodyProperties {
	elem := tb.GetElement("bodyPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if bp, ok := elem.(*TextBodyProperties); ok {
		return bp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextBodyProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetBodyProperties sets the text body properties.
func (tb *TextBody) SetBodyProperties(
	bp *TextBodyProperties,
) {
	// Remove existing body properties
	if existing := tb.GetElement("bodyPr", NamespaceMain); existing != nil {
		tb.RemoveChild(existing)
	}
	if bp != nil {
		tb.PrependChild(bp)
	}
}

// AddParagraph adds a new paragraph with the given text.
func (tb *TextBody) AddParagraph(
	text string,
) *TextParagraph {
	p := NewTextParagraph()
	if text != "" {
		r := NewTextRun(text)
		p.AppendChild(r)
	}
	tb.AppendChild(p)

	return p
}

// AddEmptyParagraph adds an empty paragraph.
func (tb *TextBody) AddEmptyParagraph() *TextParagraph {
	p := NewTextParagraph()
	tb.AppendChild(p)

	return p
}

// Paragraphs returns all paragraphs in the text body.
func (tb *TextBody) Paragraphs() []*TextParagraph {
	var paragraphs []*TextParagraph
	for child := range tb.Children() {
		if child.LocalName() == "p" &&
			child.NamespaceURI() == NamespaceMain {
			switch v := child.(type) {
			case *TextParagraph:
				paragraphs = append(paragraphs, v)
			case *openxml.CompositeElementBase:
				paragraphs = append(paragraphs, &TextParagraph{CompositeElementBase: v})
			}
		}
	}

	return paragraphs
}

// ClearParagraphs removes all paragraphs.
func (tb *TextBody) ClearParagraphs() {
	for _, p := range tb.Paragraphs() {
		tb.RemoveChild(p)
	}
}

// AutoDetectRTL automatically detects RTL for all paragraphs in the text body.
func (tb *TextBody) AutoDetectRTL() {
	for _, p := range tb.Paragraphs() {
		p.AutoDetectRTL()
	}
}

// Clone creates a deep copy of this TextBody element.
func (tb *TextBody) Clone() openxml.Element {
	return &TextBody{
		CompositeElementBase: tb.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextBodyProperties (a:bodyPr) - Text body properties
// ===========================================================================

// TextBodyProperties represents text body properties (a:bodyPr).
type TextBodyProperties struct {
	*openxml.CompositeElementBase
}

// NewTextBodyProperties creates a new text body properties element.
func NewTextBodyProperties() *TextBodyProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"bodyPr",
		PrefixMain,
	)

	return &TextBodyProperties{
		CompositeElementBase: elem,
	}
}

// Rotation returns the rotation in 60000ths of a degree.
func (bp *TextBodyProperties) Rotation() int {
	attr, found := bp.GetAttribute("rot", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetRotation sets the rotation in 60000ths of a degree.
func (bp *TextBodyProperties) SetRotation(
	rot int,
) {
	if rot == 0 {
		bp.RemoveAttribute("rot", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"rot",
			"",
			strconv.Itoa(rot),
		),
	)
}

// RotationDegrees returns the rotation in degrees.
func (bp *TextBodyProperties) RotationDegrees() float64 {
	return AngleUnitsToDegrees(bp.Rotation())
}

// SetRotationDegrees sets the rotation in degrees.
func (bp *TextBodyProperties) SetRotationDegrees(
	degrees float64,
) {
	bp.SetRotation(DegreesToAngleUnits(degrees))
}

// Vertical returns the vertical text orientation.
func (bp *TextBodyProperties) Vertical() TextVerticalValue {
	attr, found := bp.GetAttribute("vert", "")
	if !found {
		return TextVerticalHorizontal
	}

	return TextVerticalValue(attr.Value())
}

// SetVertical sets the vertical text orientation.
func (bp *TextBodyProperties) SetVertical(
	vert TextVerticalValue,
) {
	if vert == "" ||
		vert == TextVerticalHorizontal {
		bp.RemoveAttribute("vert", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"vert",
			"",
			string(vert),
		),
	)
}

// Wrap returns the text wrapping type.
func (bp *TextBodyProperties) Wrap() TextWrapValue {
	attr, found := bp.GetAttribute("wrap", "")
	if !found {
		return TextWrapSquare
	}

	return TextWrapValue(attr.Value())
}

// SetWrap sets the text wrapping type.
func (bp *TextBodyProperties) SetWrap(
	wrap TextWrapValue,
) {
	if wrap == "" || wrap == TextWrapSquare {
		bp.RemoveAttribute("wrap", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"wrap",
			"",
			string(wrap),
		),
	)
}

// Anchor returns the vertical anchor.
func (bp *TextBodyProperties) Anchor() TextAnchorValue {
	attr, found := bp.GetAttribute("anchor", "")
	if !found {
		return TextAnchorTop
	}

	return TextAnchorValue(attr.Value())
}

// SetAnchor sets the vertical anchor.
func (bp *TextBodyProperties) SetAnchor(
	anchor TextAnchorValue,
) {
	if anchor == "" || anchor == TextAnchorTop {
		bp.RemoveAttribute("anchor", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"anchor",
			"",
			string(anchor),
		),
	)
}

// AnchorCenter returns whether the text is centered horizontally in the text box.
func (bp *TextBodyProperties) AnchorCenter() bool {
	attr, found := bp.GetAttribute(
		"anchorCtr",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetAnchorCenter sets whether the text is centered horizontally.
func (bp *TextBodyProperties) SetAnchorCenter(
	center bool,
) {
	if !center {
		bp.RemoveAttribute("anchorCtr", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"anchorCtr",
			"",
			"1",
		),
	)
}

// LeftInset returns the left inset in EMUs.
func (bp *TextBodyProperties) LeftInset() int {
	attr, found := bp.GetAttribute("lIns", "")
	if !found {
		return 91440 // default: 0.1 inch
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLeftInset sets the left inset in EMUs.
func (bp *TextBodyProperties) SetLeftInset(
	inset int,
) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"lIns",
			"",
			strconv.Itoa(inset),
		),
	)
}

// TopInset returns the top inset in EMUs.
func (bp *TextBodyProperties) TopInset() int {
	attr, found := bp.GetAttribute("tIns", "")
	if !found {
		return 45720 // default: 0.05 inch
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetTopInset sets the top inset in EMUs.
func (bp *TextBodyProperties) SetTopInset(
	inset int,
) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"tIns",
			"",
			strconv.Itoa(inset),
		),
	)
}

// RightInset returns the right inset in EMUs.
func (bp *TextBodyProperties) RightInset() int {
	attr, found := bp.GetAttribute("rIns", "")
	if !found {
		return 91440 // default: 0.1 inch
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetRightInset sets the right inset in EMUs.
func (bp *TextBodyProperties) SetRightInset(
	inset int,
) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"rIns",
			"",
			strconv.Itoa(inset),
		),
	)
}

// BottomInset returns the bottom inset in EMUs.
func (bp *TextBodyProperties) BottomInset() int {
	attr, found := bp.GetAttribute("bIns", "")
	if !found {
		return 45720 // default: 0.05 inch
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetBottomInset sets the bottom inset in EMUs.
func (bp *TextBodyProperties) SetBottomInset(
	inset int,
) {
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"bIns",
			"",
			strconv.Itoa(inset),
		),
	)
}

// SetInsets sets all insets in EMUs.
func (bp *TextBodyProperties) SetInsets(
	left, top, right, bottom int,
) {
	bp.SetLeftInset(left)
	bp.SetTopInset(top)
	bp.SetRightInset(right)
	bp.SetBottomInset(bottom)
}

// ColumnCount returns the number of columns.
func (bp *TextBodyProperties) ColumnCount() int {
	attr, found := bp.GetAttribute("numCol", "")
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetColumnCount sets the number of columns (1-16).
func (bp *TextBodyProperties) SetColumnCount(
	count int,
) {
	if count <= 1 {
		bp.RemoveAttribute("numCol", "")

		return
	}
	if count > 16 {
		count = 16
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"numCol",
			"",
			strconv.Itoa(count),
		),
	)
}

// ColumnSpacing returns the space between columns in EMUs.
func (bp *TextBodyProperties) ColumnSpacing() int {
	attr, found := bp.GetAttribute("spcCol", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetColumnSpacing sets the space between columns in EMUs.
func (bp *TextBodyProperties) SetColumnSpacing(
	spacing int,
) {
	if spacing <= 0 {
		bp.RemoveAttribute("spcCol", "")

		return
	}
	bp.SetAttribute(
		openxml.NewAttribute(
			"",
			"spcCol",
			"",
			strconv.Itoa(spacing),
		),
	)
}

// Clone creates a deep copy of this TextBodyProperties element.
func (bp *TextBodyProperties) Clone() openxml.Element {
	return &TextBodyProperties{
		CompositeElementBase: bp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextParagraph (a:p) - Individual paragraph
// ===========================================================================

// TextParagraph represents a text paragraph element (a:p).
type TextParagraph struct {
	*openxml.CompositeElementBase
}

// NewTextParagraph creates a new text paragraph element.
func NewTextParagraph() *TextParagraph {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"p",
		PrefixMain,
	)

	return &TextParagraph{
		CompositeElementBase: elem,
	}
}

// Properties returns the paragraph properties.
func (p *TextParagraph) Properties() *TextParagraphProperties {
	elem := p.GetElement("pPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*TextParagraphProperties); ok {
		return pp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextParagraphProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// EnsureProperties returns the paragraph properties, creating them if needed.
func (p *TextParagraph) EnsureProperties() *TextParagraphProperties {
	if pp := p.Properties(); pp != nil {
		return pp
	}
	pp := NewTextParagraphProperties()
	p.PrependChild(pp)

	return pp
}

// SetProperties sets the paragraph properties.
func (p *TextParagraph) SetProperties(
	pp *TextParagraphProperties,
) {
	// Remove existing properties
	if existing := p.GetElement("pPr", NamespaceMain); existing != nil {
		p.RemoveChild(existing)
	}
	if pp != nil {
		p.PrependChild(pp)
	}
}

// AddRun adds a text run to the paragraph.
func (p *TextParagraph) AddRun(
	text string,
) *TextRun {
	r := NewTextRun(text)
	// Insert before endParaRPr if present
	if endRpr := p.GetElement("endParaRPr", NamespaceMain); endRpr != nil {
		p.InsertBefore(r, endRpr)
	} else {
		p.AppendChild(r)
	}

	return r
}

// AddLineBreak adds a line break to the paragraph.
func (p *TextParagraph) AddLineBreak() *TextLineBreak {
	br := NewTextLineBreak()
	// Insert before endParaRPr if present
	if endRpr := p.GetElement("endParaRPr", NamespaceMain); endRpr != nil {
		p.InsertBefore(br, endRpr)
	} else {
		p.AppendChild(br)
	}

	return br
}

// Runs returns all text runs in the paragraph.
func (p *TextParagraph) Runs() []*TextRun {
	var runs []*TextRun
	for child := range p.Children() {
		if child.LocalName() == "r" &&
			child.NamespaceURI() == NamespaceMain {
			switch v := child.(type) {
			case *TextRun:
				runs = append(runs, v)
			case *openxml.CompositeElementBase:
				runs = append(runs, &TextRun{CompositeElementBase: v})
			}
		}
	}

	return runs
}

// EndParagraphRunProperties returns the end paragraph run properties.
func (p *TextParagraph) EndParagraphRunProperties() *TextCharacterProperties {
	elem := p.GetElement(
		"endParaRPr",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*TextCharacterProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextCharacterProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// EnsureEndParagraphRunProperties returns end paragraph run properties, creating if needed.
func (p *TextParagraph) EnsureEndParagraphRunProperties() *TextCharacterProperties {
	if rp := p.EndParagraphRunProperties(); rp != nil {
		return rp
	}
	rp := NewEndParagraphRunProperties()
	p.AppendChild(rp)

	return rp
}

// SetAlignment sets the paragraph alignment.
func (p *TextParagraph) SetAlignment(
	align TextAlignValue,
) {
	pp := p.EnsureProperties()
	pp.SetAlignment(align)
}

// GetText returns the combined text content of all runs in the paragraph.
func (p *TextParagraph) GetText() string {
	var text string
	for _, r := range p.Runs() {
		text += r.Text()
	}

	return text
}

// AutoDetectRTL automatically detects if the paragraph should be RTL
// based on its text content.
func (p *TextParagraph) AutoDetectRTL() {
	text := p.GetText()
	if text == "" {
		return
	}

	p.EnsureProperties().SetRightToLeft(IsRTL(text))
}

// IsRTL returns true if the text contains predominantly RTL characters.
func IsRTL(text string) bool {
	p := bidi.Paragraph{}
	_, _ = p.SetString(text)
	direction := p.Direction()

	return direction == bidi.RightToLeft
}

// SetLevel sets the paragraph level (0-8).
func (p *TextParagraph) SetLevel(level int) {
	pp := p.EnsureProperties()
	pp.SetLevel(level)
}

// SetCharacterBullet sets the paragraph to use a character bullet.
func (p *TextParagraph) SetCharacterBullet(char string) {
	p.EnsureProperties().SetCharacterBullet(char)
}

// SetAutoNumberedBullet sets the paragraph to use auto-numbered bullets.
func (p *TextParagraph) SetAutoNumberedBullet(scheme AutoNumberSchemeValue, startAt int) {
	p.EnsureProperties().SetAutoNumberedBullet(scheme, startAt)
}

// SetBulletFont sets the bullet font.
func (p *TextParagraph) SetBulletFont(typeface string) {
	p.EnsureProperties().SetBulletFont(typeface)
}

// SetRTLAlignment sets both RTL direction and right alignment.
func (p *TextParagraph) SetRTLAlignment() {
	pp := p.EnsureProperties()
	pp.SetRightToLeft(true)
	pp.SetAlignment(TextAlignRight)
}

// Clone creates a deep copy of this TextParagraph element.
func (p *TextParagraph) Clone() openxml.Element {
	return &TextParagraph{
		CompositeElementBase: p.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Hyperlink represents a hyperlink element (a:hlinkClick).
type Hyperlink struct {
	*openxml.CompositeElementBase
}

// NewHyperlink creates a new Hyperlink element with relationship ID.
func NewHyperlink(relId string) *Hyperlink {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"hlinkClick",
		PrefixMain,
	)
	h := &Hyperlink{CompositeElementBase: elem}
	h.SetAttribute(
		openxml.NewAttribute(
			"http://schemas.openxmlformats.org/officeDocument/2006/relationships",
			"id",
			"r",
			relId,
		),
	)

	return h
}

// Clone creates a deep copy of this Hyperlink element.
func (h *Hyperlink) Clone() openxml.Element {
	return &Hyperlink{
		CompositeElementBase: h.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextParagraphProperties (a:pPr) - Paragraph properties
// ===========================================================================

// TextParagraphProperties represents text paragraph properties (a:pPr).
type TextParagraphProperties struct {
	*openxml.CompositeElementBase
}

// NewTextParagraphProperties creates a new text paragraph properties element.
func NewTextParagraphProperties() *TextParagraphProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"pPr",
		PrefixMain,
	)

	return &TextParagraphProperties{
		CompositeElementBase: elem,
	}
}

// Alignment returns the text alignment.
func (pp *TextParagraphProperties) Alignment() TextAlignValue {
	attr, found := pp.GetAttribute("algn", "")
	if !found {
		return ""
	}

	return TextAlignValue(attr.Value())
}

// SetAlignment sets the text alignment.
func (pp *TextParagraphProperties) SetAlignment(
	align TextAlignValue,
) {
	if align == "" {
		pp.RemoveAttribute("algn", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"algn",
			"",
			string(align),
		),
	)
}

// Level returns the paragraph level (0-8).
func (pp *TextParagraphProperties) Level() int {
	attr, found := pp.GetAttribute("lvl", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLevel sets the paragraph level (0-8).
func (pp *TextParagraphProperties) SetLevel(
	level int,
) {
	if level <= 0 {
		pp.RemoveAttribute("lvl", "")

		return
	}
	if level > 8 {
		level = 8
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"lvl",
			"",
			strconv.Itoa(level),
		),
	)
}

// LeftMargin returns the left margin in EMUs.
func (pp *TextParagraphProperties) LeftMargin() int {
	attr, found := pp.GetAttribute("marL", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLeftMargin sets the left margin in EMUs.
func (pp *TextParagraphProperties) SetLeftMargin(
	margin int,
) {
	if margin <= 0 {
		pp.RemoveAttribute("marL", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marL",
			"",
			strconv.Itoa(margin),
		),
	)
}

// RightMargin returns the right margin in EMUs.
func (pp *TextParagraphProperties) RightMargin() int {
	attr, found := pp.GetAttribute("marR", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetRightMargin sets the right margin in EMUs.
func (pp *TextParagraphProperties) SetRightMargin(
	margin int,
) {
	if margin <= 0 {
		pp.RemoveAttribute("marR", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marR",
			"",
			strconv.Itoa(margin),
		),
	)
}

// Indent returns the indent in EMUs.
func (pp *TextParagraphProperties) Indent() int {
	attr, found := pp.GetAttribute("indent", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetIndent sets the indent in EMUs.
func (pp *TextParagraphProperties) SetIndent(
	indent int,
) {
	if indent == 0 {
		pp.RemoveAttribute("indent", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"indent",
			"",
			strconv.Itoa(indent),
		),
	)
}

// DefaultTabSize returns the default tab size in EMUs.
func (pp *TextParagraphProperties) DefaultTabSize() int {
	attr, found := pp.GetAttribute("defTabSz", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDefaultTabSize sets the default tab size in EMUs.
func (pp *TextParagraphProperties) SetDefaultTabSize(
	size int,
) {
	if size <= 0 {
		pp.RemoveAttribute("defTabSz", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"defTabSz",
			"",
			strconv.Itoa(size),
		),
	)
}

// RightToLeft returns whether the paragraph is right-to-left.
func (pp *TextParagraphProperties) RightToLeft() bool {
	attr, found := pp.GetAttribute("rtl", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetRightToLeft sets whether the paragraph is right-to-left.
func (pp *TextParagraphProperties) SetRightToLeft(
	rtl bool,
) {
	if !rtl {
		pp.RemoveAttribute("rtl", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute("", "rtl", "", "1"),
	)
}

// FontAlignment returns the font alignment.
func (pp *TextParagraphProperties) FontAlignment() FontAlignValue {
	attr, found := pp.GetAttribute("fontAlgn", "")
	if !found {
		return FontAlignAuto
	}

	return FontAlignValue(attr.Value())
}

// SetFontAlignment sets the font alignment.
func (pp *TextParagraphProperties) SetFontAlignment(
	align FontAlignValue,
) {
	if align == "" || align == FontAlignAuto {
		pp.RemoveAttribute("fontAlgn", "")

		return
	}
	pp.SetAttribute(
		openxml.NewAttribute(
			"",
			"fontAlgn",
			"",
			string(align),
		),
	)
}

// LineSpacing returns the line spacing element.
func (pp *TextParagraphProperties) LineSpacing() *TextSpacing {
	elem := pp.GetElement("lnSpc", NamespaceMain)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*TextSpacing); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextSpacing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetLineSpacingPercent sets line spacing as a percentage (e.g., 100000 = 100%).
func (pp *TextParagraphProperties) SetLineSpacingPercent(
	percent int,
) {
	pp.removeSpacingElement("lnSpc")
	sp := NewTextSpacing("lnSpc")
	sp.SetPercent(percent)
	pp.AppendChild(sp)
}

// SetLineSpacingPoints sets line spacing in hundredths of a point.
func (pp *TextParagraphProperties) SetLineSpacingPoints(
	points int,
) {
	pp.removeSpacingElement("lnSpc")
	sp := NewTextSpacing("lnSpc")
	sp.SetPoints(points)
	pp.AppendChild(sp)
}

// SpaceBefore returns the space before element.
func (pp *TextParagraphProperties) SpaceBefore() *TextSpacing {
	elem := pp.GetElement("spcBef", NamespaceMain)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*TextSpacing); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextSpacing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetSpaceBeforePercent sets space before as a percentage.
func (pp *TextParagraphProperties) SetSpaceBeforePercent(
	percent int,
) {
	pp.removeSpacingElement("spcBef")
	sp := NewTextSpacing("spcBef")
	sp.SetPercent(percent)
	pp.AppendChild(sp)
}

// SetSpaceBeforePoints sets space before in hundredths of a point.
func (pp *TextParagraphProperties) SetSpaceBeforePoints(
	points int,
) {
	pp.removeSpacingElement("spcBef")
	sp := NewTextSpacing("spcBef")
	sp.SetPoints(points)
	pp.AppendChild(sp)
}

// SpaceAfter returns the space after element.
func (pp *TextParagraphProperties) SpaceAfter() *TextSpacing {
	elem := pp.GetElement("spcAft", NamespaceMain)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*TextSpacing); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextSpacing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetSpaceAfterPercent sets space after as a percentage.
func (pp *TextParagraphProperties) SetSpaceAfterPercent(
	percent int,
) {
	pp.removeSpacingElement("spcAft")
	sp := NewTextSpacing("spcAft")
	sp.SetPercent(percent)
	pp.AppendChild(sp)
}

// SetSpaceAfterPoints sets space after in hundredths of a point.
func (pp *TextParagraphProperties) SetSpaceAfterPoints(
	points int,
) {
	pp.removeSpacingElement("spcAft")
	sp := NewTextSpacing("spcAft")
	sp.SetPoints(points)
	pp.AppendChild(sp)
}

// removeSpacingElement removes the specified spacing element.
func (pp *TextParagraphProperties) removeSpacingElement(
	name string,
) {
	if elem := pp.GetElement(name, NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
}

// SetNoBullet sets the paragraph to have no bullet.
func (pp *TextParagraphProperties) SetNoBullet() {
	pp.removeBullet()
	buNone := openxml.NewLeafElement(
		NamespaceMain,
		"buNone",
		PrefixMain,
	)
	pp.AppendChild(buNone)
}

// SetCharacterBullet sets the paragraph to use a character bullet.
func (pp *TextParagraphProperties) SetCharacterBullet(
	char string,
) {
	pp.removeBullet()
	buChar := openxml.NewLeafElement(
		NamespaceMain,
		"buChar",
		PrefixMain,
	)
	buChar.SetAttribute(
		openxml.NewAttribute(
			"",
			"char",
			"",
			char,
		),
	)
	pp.AppendChild(buChar)
}

// SetAutoNumberedBullet sets the paragraph to use auto-numbered bullets.
func (pp *TextParagraphProperties) SetAutoNumberedBullet(
	scheme AutoNumberSchemeValue,
	startAt int,
) {
	pp.removeBullet()
	buAutoNum := openxml.NewLeafElement(
		NamespaceMain,
		"buAutoNum",
		PrefixMain,
	)
	buAutoNum.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			string(scheme),
		),
	)
	if startAt > 1 {
		buAutoNum.SetAttribute(
			openxml.NewAttribute(
				"",
				"startAt",
				"",
				strconv.Itoa(startAt),
			),
		)
	}
	pp.AppendChild(buAutoNum)
}

// SetBulletFont sets the bullet font.
func (pp *TextParagraphProperties) SetBulletFont(
	typeface string,
) {
	// Remove existing bullet font
	if existing := pp.GetElement("buFont", NamespaceMain); existing != nil {
		pp.RemoveChild(existing)
	}
	buFont := NewTextFont("buFont")
	buFont.SetTypeface(typeface)
	pp.AppendChild(buFont)
}

// SetBulletColor sets the bullet color with an RGB value.
func (pp *TextParagraphProperties) SetBulletColor(
	hexColor string,
) {
	// Remove existing bullet color
	pp.removeBulletColor()
	buClr := openxml.NewCompositeElement(
		NamespaceMain,
		"buClr",
		PrefixMain,
	)
	rgb := NewRgbColor(hexColor)
	buClr.AppendChild(rgb)
	pp.AppendChild(buClr)
}

// removeBullet removes any existing bullet type elements.
func (pp *TextParagraphProperties) removeBullet() {
	if elem := pp.GetElement("buNone", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
	if elem := pp.GetElement("buChar", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
	if elem := pp.GetElement("buAutoNum", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
	if elem := pp.GetElement("buBlip", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
}

// removeBulletColor removes the bullet color element.
func (pp *TextParagraphProperties) removeBulletColor() {
	if elem := pp.GetElement("buClr", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
	if elem := pp.GetElement("buClrTx", NamespaceMain); elem != nil {
		pp.RemoveChild(elem)
	}
}

// DefaultRunProperties returns the default run properties.
func (pp *TextParagraphProperties) DefaultRunProperties() *TextCharacterProperties {
	elem := pp.GetElement("defRPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*TextCharacterProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextCharacterProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// EnsureDefaultRunProperties returns the default run properties, creating if needed.
func (pp *TextParagraphProperties) EnsureDefaultRunProperties() *TextCharacterProperties {
	if rp := pp.DefaultRunProperties(); rp != nil {
		return rp
	}
	rp := NewDefaultRunProperties()
	pp.AppendChild(rp)

	return rp
}

// Clone creates a deep copy of this TextParagraphProperties element.
func (pp *TextParagraphProperties) Clone() openxml.Element {
	return &TextParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextSpacing - Spacing elements (a:lnSpc, a:spcBef, a:spcAft)
// ===========================================================================

// TextSpacing represents text spacing elements.
type TextSpacing struct {
	*openxml.CompositeElementBase
}

// NewTextSpacing creates a new text spacing element.
func NewTextSpacing(
	localName string,
) *TextSpacing {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		localName,
		PrefixMain,
	)

	return &TextSpacing{
		CompositeElementBase: elem,
	}
}

// SetPercent sets the spacing as a percentage (in 1000ths of a percent).
func (sp *TextSpacing) SetPercent(percent int) {
	sp.RemoveAllChildren()
	spcPct := openxml.NewLeafElement(
		NamespaceMain,
		"spcPct",
		PrefixMain,
	)
	spcPct.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.Itoa(percent),
		),
	)
	sp.AppendChild(spcPct)
}

// SetPoints sets the spacing in hundredths of a point.
func (sp *TextSpacing) SetPoints(points int) {
	sp.RemoveAllChildren()
	spcPts := openxml.NewLeafElement(
		NamespaceMain,
		"spcPts",
		PrefixMain,
	)
	spcPts.SetAttribute(
		openxml.NewAttribute(
			"",
			"val",
			"",
			strconv.Itoa(points),
		),
	)
	sp.AppendChild(spcPts)
}

// IsPercent returns true if the spacing is specified as a percentage.
func (sp *TextSpacing) IsPercent() bool {
	return sp.GetElement(
		"spcPct",
		NamespaceMain,
	) != nil
}

// Percent returns the percentage value (in 1000ths of a percent).
func (sp *TextSpacing) Percent() int {
	elem := sp.GetElement("spcPct", NamespaceMain)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Points returns the point value (in hundredths of a point).
func (sp *TextSpacing) Points() int {
	elem := sp.GetElement("spcPts", NamespaceMain)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute("val", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this TextSpacing element.
func (sp *TextSpacing) Clone() openxml.Element {
	return &TextSpacing{
		CompositeElementBase: sp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextRun (a:r) - Text run with formatting
// ===========================================================================

// TextRun represents a text run element (a:r).
type TextRun struct {
	*openxml.CompositeElementBase
}

// NewTextRun creates a new text run element with the given text.
func NewTextRun(text string) *TextRun {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"r",
		PrefixMain,
	)
	r := &TextRun{CompositeElementBase: elem}
	r.SetText(text)

	return r
}

// Properties returns the run properties.
func (r *TextRun) Properties() *TextCharacterProperties {
	elem := r.GetElement("rPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*TextCharacterProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextCharacterProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// EnsureProperties returns the run properties, creating them if needed.
func (r *TextRun) EnsureProperties() *TextCharacterProperties {
	if rp := r.Properties(); rp != nil {
		return rp
	}
	rp := NewRunProperties()
	r.PrependChild(rp)

	return rp
}

// SetProperties sets the run properties.
func (r *TextRun) SetProperties(
	rp *TextCharacterProperties,
) {
	// Remove existing properties
	if existing := r.GetElement("rPr", NamespaceMain); existing != nil {
		r.RemoveChild(existing)
	}
	if rp != nil {
		r.PrependChild(rp)
	}
}

// Text returns the text content.
func (r *TextRun) Text() string {
	elem := r.GetElement("t", NamespaceMain)
	if elem == nil {
		return ""
	}
	// Text element should be a leaf with text content
	if te, ok := elem.(*TextElement); ok {
		return te.Text()
	}
	// Check for CompositeElement with InnerXml
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return comp.InnerXml()
	}
	// Check for text wrapper
	if tw, ok := elem.(*textWrapper); ok {
		return tw.text
	}

	return ""
}

// SetText sets the text content.
func (r *TextRun) SetText(text string) {
	// Remove existing text element
	if existing := r.GetElement("t", NamespaceMain); existing != nil {
		r.RemoveChild(existing)
	}
	t := NewTextElement(text)
	r.AppendChild(t)
}

// SetBold sets whether the text is bold.
func (r *TextRun) SetBold(bold bool) {
	rp := r.EnsureProperties()
	rp.SetBold(bold)
}

// SetRightToLeft sets whether the text is right-to-left.
func (r *TextRun) SetRightToLeft(rtl bool) {
	rp := r.EnsureProperties()
	rp.SetRightToLeft(rtl)
}

// SetHyperlink sets the hyperlink relationship ID.
func (r *TextRun) SetHyperlink(relId string) {
	r.EnsureProperties().SetHyperlink(relId)
}

// SetItalic sets whether the text is italic.
func (r *TextRun) SetItalic(italic bool) {
	rp := r.EnsureProperties()
	rp.SetItalic(italic)
}

// SetUnderline sets the underline style.
func (r *TextRun) SetUnderline(
	style UnderlineValue,
) {
	rp := r.EnsureProperties()
	rp.SetUnderline(style)
}

// SetFontSize sets the font size in hundredths of a point.
func (r *TextRun) SetFontSize(size int) {
	rp := r.EnsureProperties()
	rp.SetFontSize(size)
}

// SetFontSizePoints sets the font size in points.
func (r *TextRun) SetFontSizePoints(
	points float64,
) {
	rp := r.EnsureProperties()
	rp.SetFontSizePoints(points)
}

// SetColor sets the text color with an RGB value.
func (r *TextRun) SetColor(hexColor string) {
	rp := r.EnsureProperties()
	rp.SetSolidFill(hexColor)
}

// SetLatinFont sets the Latin (Western) font.
func (r *TextRun) SetLatinFont(typeface string) {
	rp := r.EnsureProperties()
	rp.SetLatinFont(typeface)
}

// Clone creates a deep copy of this TextRun element.
func (r *TextRun) Clone() openxml.Element {
	return &TextRun{
		CompositeElementBase: r.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextElement (a:t) - Text content element
// ===========================================================================

// TextElement represents a text element (a:t).
type TextElement struct {
	*textWrapper
}

// textWrapper is an internal type to hold text content.
type textWrapper struct {
	*openxml.LeafElementBase
	text string
}

// NewTextElement creates a new text element with the given content.
func NewTextElement(text string) *TextElement {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"t",
		PrefixMain,
	)
	elem.SetInnerText(text)
	tw := &textWrapper{
		LeafElementBase: elem,
		text:            text,
	}

	return &TextElement{textWrapper: tw}
}

// Text returns the text content.
func (t *TextElement) Text() string {
	return t.text
}

// SetText sets the text content.
func (t *TextElement) SetText(text string) {
	t.text = text
	t.SetInnerText(text)
}

// Clone creates a deep copy of this TextElement.
func (t *TextElement) Clone() openxml.Element {
	return NewTextElement(t.text)
}

// ===========================================================================
// TextLineBreak (a:br) - Line break
// ===========================================================================

// TextLineBreak represents a line break element (a:br).
type TextLineBreak struct {
	*openxml.CompositeElementBase
}

// NewTextLineBreak creates a new line break element.
func NewTextLineBreak() *TextLineBreak {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"br",
		PrefixMain,
	)

	return &TextLineBreak{
		CompositeElementBase: elem,
	}
}

// Properties returns the run properties for the line break.
func (br *TextLineBreak) Properties() *TextCharacterProperties {
	elem := br.GetElement("rPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*TextCharacterProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TextCharacterProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this TextLineBreak element.
func (br *TextLineBreak) Clone() openxml.Element {
	return &TextLineBreak{
		CompositeElementBase: br.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextField (a:fld) - Field element
// ===========================================================================

// TextField represents a field element (a:fld).
type TextField struct {
	*openxml.CompositeElementBase
}

// NewTextField creates a new field element.
func NewTextField(
	id, fieldType string,
) *TextField {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"fld",
		PrefixMain,
	)
	f := &TextField{CompositeElementBase: elem}
	f.SetAttribute(
		openxml.NewAttribute("", "id", "", id),
	)
	if fieldType != "" {
		f.SetAttribute(
			openxml.NewAttribute(
				"",
				"type",
				"",
				fieldType,
			),
		)
	}

	return f
}

// ID returns the field ID.
func (f *TextField) ID() string {
	attr, found := f.GetAttribute("id", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetID sets the field ID.
func (f *TextField) SetID(id string) {
	f.SetAttribute(
		openxml.NewAttribute("", "id", "", id),
	)
}

// FieldType returns the field type.
func (f *TextField) FieldType() string {
	attr, found := f.GetAttribute("type", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFieldType sets the field type.
func (f *TextField) SetFieldType(
	fieldType string,
) {
	if fieldType == "" {
		f.RemoveAttribute("type", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"type",
			"",
			fieldType,
		),
	)
}

// Text returns the text content.
func (f *TextField) Text() string {
	elem := f.GetElement("t", NamespaceMain)
	if elem == nil {
		return ""
	}
	if te, ok := elem.(*TextElement); ok {
		return te.Text()
	}

	return ""
}

// SetText sets the text content.
func (f *TextField) SetText(text string) {
	// Remove existing text element
	if existing := f.GetElement("t", NamespaceMain); existing != nil {
		f.RemoveChild(existing)
	}
	t := NewTextElement(text)
	f.AppendChild(t)
}

// Clone creates a deep copy of this TextField element.
func (f *TextField) Clone() openxml.Element {
	return &TextField{
		CompositeElementBase: f.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextCharacterProperties (a:rPr, a:defRPr, a:endParaRPr) - Run properties
// ===========================================================================

// TextCharacterProperties represents text character properties.
type TextCharacterProperties struct {
	*openxml.CompositeElementBase
}

// NewRunProperties creates a new run properties element (a:rPr).
func NewRunProperties() *TextCharacterProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"rPr",
		PrefixMain,
	)

	return &TextCharacterProperties{
		CompositeElementBase: elem,
	}
}

// NewDefaultRunProperties creates a new default run properties element (a:defRPr).
func NewDefaultRunProperties() *TextCharacterProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"defRPr",
		PrefixMain,
	)

	return &TextCharacterProperties{
		CompositeElementBase: elem,
	}
}

// NewEndParagraphRunProperties creates a new end paragraph run properties element (a:endParaRPr).
func NewEndParagraphRunProperties() *TextCharacterProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"endParaRPr",
		PrefixMain,
	)

	return &TextCharacterProperties{
		CompositeElementBase: elem,
	}
}

// FontSize returns the font size in hundredths of a point.
func (rp *TextCharacterProperties) FontSize() int {
	attr, found := rp.GetAttribute("sz", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFontSize sets the font size in hundredths of a point.
// Valid range: 100 to 400000 (1pt to 4000pt).
func (rp *TextCharacterProperties) SetFontSize(
	size int,
) {
	if size <= 0 {
		rp.RemoveAttribute("sz", "")

		return
	}
	if size < 100 {
		size = 100
	}
	if size > 400000 {
		size = 400000
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"sz",
			"",
			strconv.Itoa(size),
		),
	)
}

// FontSizePoints returns the font size in points.
func (rp *TextCharacterProperties) FontSizePoints() float64 {
	return float64(rp.FontSize()) / 100.0
}

// SetFontSizePoints sets the font size in points.
func (rp *TextCharacterProperties) SetFontSizePoints(
	points float64,
) {
	rp.SetFontSize(int(points * 100))
}

// Bold returns whether the text is bold.
func (rp *TextCharacterProperties) Bold() bool {
	attr, found := rp.GetAttribute("b", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetBold sets whether the text is bold.
func (rp *TextCharacterProperties) SetBold(
	bold bool,
) {
	if !bold {
		rp.RemoveAttribute("b", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute("", "b", "", "1"),
	)
}

// RightToLeft returns whether the text is right-to-left.
func (rp *TextCharacterProperties) RightToLeft() bool {
	attr, found := rp.GetAttribute("rtl", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == "true"
}

// SetRightToLeft sets whether the text is right-to-left.
func (rp *TextCharacterProperties) SetRightToLeft(
	rtl bool,
) {
	if !rtl {
		rp.RemoveAttribute("rtl", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute("", "rtl", "", "1"),
	)
}

// SetHyperlink sets the hyperlink relationship ID.
func (rp *TextCharacterProperties) SetHyperlink(relId string) {
	// Remove existing hyperlink if any
	if existing := rp.GetElement("hlinkClick", NamespaceMain); existing != nil {
		rp.RemoveChild(existing)
	}
	if relId != "" {
		rp.AppendChild(NewHyperlink(relId))
	}
}

// Italic returns whether the text is italic.
func (rp *TextCharacterProperties) Italic() bool {
	attr, found := rp.GetAttribute("i", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetItalic sets whether the text is italic.
func (rp *TextCharacterProperties) SetItalic(
	italic bool,
) {
	if !italic {
		rp.RemoveAttribute("i", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute("", "i", "", "1"),
	)
}

// Underline returns the underline style.
func (rp *TextCharacterProperties) Underline() UnderlineValue {
	attr, found := rp.GetAttribute("u", "")
	if !found {
		return UnderlineNone
	}

	return UnderlineValue(attr.Value())
}

// SetUnderline sets the underline style.
func (rp *TextCharacterProperties) SetUnderline(
	style UnderlineValue,
) {
	if style == "" || style == UnderlineNone {
		rp.RemoveAttribute("u", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"u",
			"",
			string(style),
		),
	)
}

// Strike returns the strikethrough style.
func (rp *TextCharacterProperties) Strike() StrikeValue {
	attr, found := rp.GetAttribute("strike", "")
	if !found {
		return StrikeNoStrike
	}

	return StrikeValue(attr.Value())
}

// SetStrike sets the strikethrough style.
func (rp *TextCharacterProperties) SetStrike(
	strike StrikeValue,
) {
	if strike == "" || strike == StrikeNoStrike {
		rp.RemoveAttribute("strike", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"strike",
			"",
			string(strike),
		),
	)
}

// Capitalization returns the capitalization style.
func (rp *TextCharacterProperties) Capitalization() CapValue {
	attr, found := rp.GetAttribute("cap", "")
	if !found {
		return CapNone
	}

	return CapValue(attr.Value())
}

// SetCapitalization sets the capitalization style.
func (rp *TextCharacterProperties) SetCapitalization(
	capValue CapValue,
) {
	if capValue == "" || capValue == CapNone {
		rp.RemoveAttribute("cap", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"cap",
			"",
			string(capValue),
		),
	)
}

// Kerning returns the kerning value in hundredths of a point.
func (rp *TextCharacterProperties) Kerning() int {
	attr, found := rp.GetAttribute("kern", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetKerning sets the kerning value in hundredths of a point.
func (rp *TextCharacterProperties) SetKerning(
	kern int,
) {
	if kern <= 0 {
		rp.RemoveAttribute("kern", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"kern",
			"",
			strconv.Itoa(kern),
		),
	)
}

// Spacing returns the character spacing in hundredths of a point.
func (rp *TextCharacterProperties) Spacing() int {
	attr, found := rp.GetAttribute("spc", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSpacing sets the character spacing in hundredths of a point.
func (rp *TextCharacterProperties) SetSpacing(
	spacing int,
) {
	if spacing == 0 {
		rp.RemoveAttribute("spc", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"spc",
			"",
			strconv.Itoa(spacing),
		),
	)
}

// Baseline returns the baseline offset in percentage (e.g., 30000 = 30%).
func (rp *TextCharacterProperties) Baseline() int {
	attr, found := rp.GetAttribute("baseline", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetBaseline sets the baseline offset for superscript/subscript.
// Positive values are superscript, negative values are subscript.
func (rp *TextCharacterProperties) SetBaseline(
	baseline int,
) {
	if baseline == 0 {
		rp.RemoveAttribute("baseline", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"baseline",
			"",
			strconv.Itoa(baseline),
		),
	)
}

// SetSuperscript sets the text as superscript.
func (rp *TextCharacterProperties) SetSuperscript() {
	rp.SetBaseline(30000) // 30% baseline offset
}

// SetSubscript sets the text as subscript.
func (rp *TextCharacterProperties) SetSubscript() {
	rp.SetBaseline(-25000) // -25% baseline offset
}

// Language returns the language code.
func (rp *TextCharacterProperties) Language() string {
	attr, found := rp.GetAttribute("lang", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLanguage sets the language code (e.g., "en-US").
func (rp *TextCharacterProperties) SetLanguage(
	lang string,
) {
	if lang == "" {
		rp.RemoveAttribute("lang", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"lang",
			"",
			lang,
		),
	)
}

// Dirty returns whether the text has been changed since last run.
func (rp *TextCharacterProperties) Dirty() bool {
	attr, found := rp.GetAttribute("dirty", "")
	if !found {
		return false
	}

	return attr.Value() == "1" ||
		attr.Value() == attrTrue
}

// SetDirty sets the dirty flag.
func (rp *TextCharacterProperties) SetDirty(
	dirty bool,
) {
	if !dirty {
		rp.RemoveAttribute("dirty", "")

		return
	}
	rp.SetAttribute(
		openxml.NewAttribute(
			"",
			"dirty",
			"",
			"1",
		),
	)
}

// SetSolidFill sets a solid color fill for the text.
func (rp *TextCharacterProperties) SetSolidFill(
	hexColor string,
) {
	rp.removeFill()
	sf := NewSolidFillWithRgb(hexColor)
	rp.PrependChild(sf)
}

// SetSolidFillSchemeColor sets a solid scheme color fill.
func (rp *TextCharacterProperties) SetSolidFillSchemeColor(
	color SchemeColorValue,
) {
	rp.removeFill()
	sf := NewSolidFillWithSchemeColor(color)
	rp.PrependChild(sf)
}

// SetNoFill sets the text to have no fill.
func (rp *TextCharacterProperties) SetNoFill() {
	rp.removeFill()
	nf := NewNoFill()
	rp.PrependChild(nf)
}

// removeFill removes any existing fill element.
func (rp *TextCharacterProperties) removeFill() {
	if nf := rp.GetElement("noFill", NamespaceMain); nf != nil {
		rp.RemoveChild(nf)
	}
	if sf := rp.GetElement("solidFill", NamespaceMain); sf != nil {
		rp.RemoveChild(sf)
	}
	if gf := rp.GetElement("gradFill", NamespaceMain); gf != nil {
		rp.RemoveChild(gf)
	}
	if pf := rp.GetElement("pattFill", NamespaceMain); pf != nil {
		rp.RemoveChild(pf)
	}
}

// SolidFill returns the solid fill if set, or nil.
func (rp *TextCharacterProperties) SolidFill() *SolidFill {
	elem := rp.GetElement(
		"solidFill",
		NamespaceMain,
	)
	if elem == nil {
		return nil
	}
	if sf, ok := elem.(*SolidFill); ok {
		return sf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SolidFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetLatinFont sets the Latin (Western) font.
func (rp *TextCharacterProperties) SetLatinFont(
	typeface string,
) {
	if existing := rp.GetElement("latin", NamespaceMain); existing != nil {
		rp.RemoveChild(existing)
	}
	font := NewTextFont("latin")
	font.SetTypeface(typeface)
	rp.AppendChild(font)
}

// LatinFont returns the Latin font.
func (rp *TextCharacterProperties) LatinFont() *TextFont {
	elem := rp.GetElement("latin", NamespaceMain)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*TextFont); ok {
		return f
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TextFont{LeafElementBase: leaf}
	}

	return nil
}

// SetEastAsianFont sets the East Asian font.
func (rp *TextCharacterProperties) SetEastAsianFont(
	typeface string,
) {
	if existing := rp.GetElement("ea", NamespaceMain); existing != nil {
		rp.RemoveChild(existing)
	}
	font := NewTextFont("ea")
	font.SetTypeface(typeface)
	rp.AppendChild(font)
}

// EastAsianFont returns the East Asian font.
func (rp *TextCharacterProperties) EastAsianFont() *TextFont {
	elem := rp.GetElement("ea", NamespaceMain)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*TextFont); ok {
		return f
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TextFont{LeafElementBase: leaf}
	}

	return nil
}

// SetComplexScriptFont sets the complex script (CS) font.
func (rp *TextCharacterProperties) SetComplexScriptFont(
	typeface string,
) {
	if existing := rp.GetElement("cs", NamespaceMain); existing != nil {
		rp.RemoveChild(existing)
	}
	font := NewTextFont("cs")
	font.SetTypeface(typeface)
	rp.AppendChild(font)
}

// ComplexScriptFont returns the complex script font.
func (rp *TextCharacterProperties) ComplexScriptFont() *TextFont {
	elem := rp.GetElement("cs", NamespaceMain)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*TextFont); ok {
		return f
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TextFont{LeafElementBase: leaf}
	}

	return nil
}

// SetSymbolFont sets the symbol font.
func (rp *TextCharacterProperties) SetSymbolFont(
	typeface string,
) {
	if existing := rp.GetElement("sym", NamespaceMain); existing != nil {
		rp.RemoveChild(existing)
	}
	font := NewTextFont("sym")
	font.SetTypeface(typeface)
	rp.AppendChild(font)
}

// SymbolFont returns the symbol font.
func (rp *TextCharacterProperties) SymbolFont() *TextFont {
	elem := rp.GetElement("sym", NamespaceMain)
	if elem == nil {
		return nil
	}
	if f, ok := elem.(*TextFont); ok {
		return f
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TextFont{LeafElementBase: leaf}
	}

	return nil
}

// Clone creates a deep copy of this TextCharacterProperties element.
func (rp *TextCharacterProperties) Clone() openxml.Element {
	return &TextCharacterProperties{
		CompositeElementBase: rp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// TextFont (a:latin, a:ea, a:cs, a:sym, a:buFont) - Font reference
// ===========================================================================

// TextFont represents a text font element.
type TextFont struct {
	*openxml.LeafElementBase
}

// NewTextFont creates a new text font element.
func NewTextFont(localName string) *TextFont {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		localName,
		PrefixMain,
	)

	return &TextFont{LeafElementBase: elem}
}

// NewLatinFont creates a new Latin font element.
func NewLatinFont(typeface string) *TextFont {
	f := NewTextFont("latin")
	f.SetTypeface(typeface)

	return f
}

// NewEastAsianFont creates a new East Asian font element.
func NewEastAsianFont(typeface string) *TextFont {
	f := NewTextFont("ea")
	f.SetTypeface(typeface)

	return f
}

// NewComplexScriptFont creates a new complex script font element.
func NewComplexScriptFont(
	typeface string,
) *TextFont {
	f := NewTextFont("cs")
	f.SetTypeface(typeface)

	return f
}

// NewSymbolFont creates a new symbol font element.
func NewSymbolFont(typeface string) *TextFont {
	f := NewTextFont("sym")
	f.SetTypeface(typeface)

	return f
}

// NewBulletFont creates a new bullet font element.
func NewBulletFont(typeface string) *TextFont {
	f := NewTextFont("buFont")
	f.SetTypeface(typeface)

	return f
}

// Typeface returns the font typeface name.
func (f *TextFont) Typeface() string {
	attr, found := f.GetAttribute("typeface", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTypeface sets the font typeface name.
func (f *TextFont) SetTypeface(typeface string) {
	if typeface == "" {
		f.RemoveAttribute("typeface", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"typeface",
			"",
			typeface,
		),
	)
}

// Panose returns the Panose-1 classification number.
func (f *TextFont) Panose() string {
	attr, found := f.GetAttribute("panose", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPanose sets the Panose-1 classification number.
func (f *TextFont) SetPanose(panose string) {
	if panose == "" {
		f.RemoveAttribute("panose", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"panose",
			"",
			panose,
		),
	)
}

// PitchFamily returns the pitch family value.
func (f *TextFont) PitchFamily() int {
	attr, found := f.GetAttribute(
		"pitchFamily",
		"",
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPitchFamily sets the pitch family value.
func (f *TextFont) SetPitchFamily(
	pitchFamily int,
) {
	if pitchFamily == 0 {
		f.RemoveAttribute("pitchFamily", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"pitchFamily",
			"",
			strconv.Itoa(pitchFamily),
		),
	)
}

// CharacterSet returns the character set value.
func (f *TextFont) CharacterSet() int {
	attr, found := f.GetAttribute("charset", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCharacterSet sets the character set value.
func (f *TextFont) SetCharacterSet(charset int) {
	if charset == 0 {
		f.RemoveAttribute("charset", "")

		return
	}
	f.SetAttribute(
		openxml.NewAttribute(
			"",
			"charset",
			"",
			strconv.Itoa(charset),
		),
	)
}

// Clone creates a deep copy of this TextFont element.
func (f *TextFont) Clone() openxml.Element {
	return &TextFont{
		LeafElementBase: f.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// ===========================================================================
// TabStop (a:tab) - Tab stop
// ===========================================================================

// TabStop represents a tab stop element (a:tab).
type TabStop struct {
	*openxml.LeafElementBase
}

// NewTabStop creates a new tab stop element.
func NewTabStop(
	position int,
	alignment TabAlignValue,
) *TabStop {
	elem := openxml.NewLeafElement(
		NamespaceMain,
		"tab",
		PrefixMain,
	)
	t := &TabStop{LeafElementBase: elem}
	t.SetPosition(position)
	if alignment != "" {
		t.SetAlignment(alignment)
	}

	return t
}

// Position returns the tab position in EMUs.
func (t *TabStop) Position() int {
	attr, found := t.GetAttribute("pos", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPosition sets the tab position in EMUs.
func (t *TabStop) SetPosition(pos int) {
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"pos",
			"",
			strconv.Itoa(pos),
		),
	)
}

// Alignment returns the tab alignment.
func (t *TabStop) Alignment() TabAlignValue {
	attr, found := t.GetAttribute("algn", "")
	if !found {
		return TabAlignLeft
	}

	return TabAlignValue(attr.Value())
}

// SetAlignment sets the tab alignment.
func (t *TabStop) SetAlignment(
	align TabAlignValue,
) {
	if align == "" || align == TabAlignLeft {
		t.RemoveAttribute("algn", "")

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			"",
			"algn",
			"",
			string(align),
		),
	)
}

// Clone creates a deep copy of this TabStop element.
func (t *TabStop) Clone() openxml.Element {
	return &TabStop{
		LeafElementBase: t.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}
