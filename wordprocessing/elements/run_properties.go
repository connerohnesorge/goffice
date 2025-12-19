package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// RunProperties represents run formatting properties (w:rPr).
type RunProperties struct {
	*openxml.CompositeElementBase
}

// NewRunProperties creates a new RunProperties element.
func NewRunProperties() *RunProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"rPr",
		PrefixW,
	)
	return &RunProperties{
		CompositeElementBase: elem,
	}
}

// Bold returns whether bold formatting is applied.
func (rp *RunProperties) Bold() bool {
	return rp.hasOnOffElement("b")
}

// SetBold sets bold formatting.
func (rp *RunProperties) SetBold(b bool) {
	rp.setOnOffElement("b", b)
}

// Italic returns whether italic formatting is applied.
func (rp *RunProperties) Italic() bool {
	return rp.hasOnOffElement("i")
}

// SetItalic sets italic formatting.
func (rp *RunProperties) SetItalic(b bool) {
	rp.setOnOffElement("i", b)
}

// Underline returns the underline style.
func (rp *RunProperties) Underline() UnderlineValue {
	elem := rp.GetElement("u", NamespaceWML)
	if elem == nil {
		return UnderlineNone
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return UnderlineSingle // Default if element exists but no val
	}
	return UnderlineValue(attr.Value())
}

// SetUnderline sets the underline style.
func (rp *RunProperties) SetUnderline(
	u UnderlineValue,
) {
	if u == UnderlineNone {
		rp.removeElement("u")
		return
	}
	elem := rp.getOrCreateElement("u")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(u),
		),
	)
}

// Strike returns whether strikethrough is applied.
func (rp *RunProperties) Strike() bool {
	return rp.hasOnOffElement("strike")
}

// SetStrike sets strikethrough formatting.
func (rp *RunProperties) SetStrike(b bool) {
	rp.setOnOffElement("strike", b)
}

// DoubleStrike returns whether double strikethrough is applied.
func (rp *RunProperties) DoubleStrike() bool {
	return rp.hasOnOffElement("dstrike")
}

// SetDoubleStrike sets double strikethrough formatting.
func (rp *RunProperties) SetDoubleStrike(b bool) {
	rp.setOnOffElement("dstrike", b)
}

// FontSize returns the font size in half-points.
func (rp *RunProperties) FontSize() int {
	elem := rp.GetElement("sz", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetFontSize sets the font size in half-points.
func (rp *RunProperties) SetFontSize(
	halfPoints int,
) {
	if halfPoints <= 0 {
		rp.removeElement("sz")
		return
	}
	elem := rp.getOrCreateElement("sz")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(halfPoints),
		),
	)
}

// FontSizeComplexScript returns the complex script font size in half-points.
func (rp *RunProperties) FontSizeComplexScript() int {
	elem := rp.GetElement("szCs", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetFontSizeComplexScript sets the complex script font size in half-points.
func (rp *RunProperties) SetFontSizeComplexScript(
	halfPoints int,
) {
	if halfPoints <= 0 {
		rp.removeElement("szCs")
		return
	}
	elem := rp.getOrCreateElement("szCs")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(halfPoints),
		),
	)
}

// RunFonts returns the run fonts element.
func (rp *RunProperties) RunFonts() *RunFonts {
	elem := rp.GetElement("rFonts", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rf, ok := elem.(*RunFonts); ok {
		return rf
	}
	// Wrap existing element
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunFonts{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// SetFont sets the font name for all script types.
func (rp *RunProperties) SetFont(
	fontName string,
) {
	rf := rp.getOrCreateRunFonts()
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"ascii",
			PrefixW,
			fontName,
		),
	)
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hAnsi",
			PrefixW,
			fontName,
		),
	)
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"eastAsia",
			PrefixW,
			fontName,
		),
	)
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"cs",
			PrefixW,
			fontName,
		),
	)
}

// Color returns the text color as a hex string.
func (rp *RunProperties) Color() string {
	elem := rp.GetElement("color", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetColor sets the text color from a hex string (without # prefix).
func (rp *RunProperties) SetColor(hex string) {
	if hex == "" {
		rp.removeElement("color")
		return
	}
	elem := rp.getOrCreateElement("color")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			hex,
		),
	)
}

// Highlight returns the highlight color.
func (rp *RunProperties) Highlight() HighlightColor {
	elem := rp.GetElement(
		"highlight",
		NamespaceWML,
	)
	if elem == nil {
		return HighlightNone
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return HighlightNone
	}
	return HighlightColor(attr.Value())
}

// SetHighlight sets the highlight color.
func (rp *RunProperties) SetHighlight(
	color HighlightColor,
) {
	if color == HighlightNone {
		rp.removeElement("highlight")
		return
	}
	elem := rp.getOrCreateElement("highlight")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(color),
		),
	)
}

// Caps returns whether all caps is applied.
func (rp *RunProperties) Caps() bool {
	return rp.hasOnOffElement("caps")
}

// SetCaps sets all caps formatting.
func (rp *RunProperties) SetCaps(b bool) {
	rp.setOnOffElement("caps", b)
}

// SmallCaps returns whether small caps is applied.
func (rp *RunProperties) SmallCaps() bool {
	return rp.hasOnOffElement("smallCaps")
}

// SetSmallCaps sets small caps formatting.
func (rp *RunProperties) SetSmallCaps(b bool) {
	rp.setOnOffElement("smallCaps", b)
}

// VerticalTextAlignment returns the vertical text alignment.
func (rp *RunProperties) VerticalTextAlignment() VerticalAlignValue {
	elem := rp.GetElement(
		"vertAlign",
		NamespaceWML,
	)
	if elem == nil {
		return VerticalAlignBaseline
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return VerticalAlignBaseline
	}
	return VerticalAlignValue(attr.Value())
}

// SetVerticalTextAlignment sets the vertical text alignment.
func (rp *RunProperties) SetVerticalTextAlignment(
	v VerticalAlignValue,
) {
	if v == VerticalAlignBaseline {
		rp.removeElement("vertAlign")
		return
	}
	elem := rp.getOrCreateElement("vertAlign")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(v),
		),
	)
}

// Vanish returns whether the text is hidden.
func (rp *RunProperties) Vanish() bool {
	return rp.hasOnOffElement("vanish")
}

// SetVanish sets whether the text is hidden.
func (rp *RunProperties) SetVanish(b bool) {
	rp.setOnOffElement("vanish", b)
}

// RunStyle returns the character style ID.
func (rp *RunProperties) RunStyle() string {
	elem := rp.GetElement("rStyle", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetRunStyle sets the character style ID.
func (rp *RunProperties) SetRunStyle(
	styleId string,
) {
	if styleId == "" {
		rp.removeElement("rStyle")
		return
	}
	elem := rp.getOrCreateElement("rStyle")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			styleId,
		),
	)
}

// Emboss returns whether emboss effect is applied.
func (rp *RunProperties) Emboss() bool {
	return rp.hasOnOffElement("emboss")
}

// SetEmboss sets emboss effect.
func (rp *RunProperties) SetEmboss(b bool) {
	rp.setOnOffElement("emboss", b)
}

// Imprint returns whether imprint/engrave effect is applied.
func (rp *RunProperties) Imprint() bool {
	return rp.hasOnOffElement("imprint")
}

// SetImprint sets imprint/engrave effect.
func (rp *RunProperties) SetImprint(b bool) {
	rp.setOnOffElement("imprint", b)
}

// Shadow returns whether shadow effect is applied.
func (rp *RunProperties) Shadow() bool {
	return rp.hasOnOffElement("shadow")
}

// SetShadow sets shadow effect.
func (rp *RunProperties) SetShadow(b bool) {
	rp.setOnOffElement("shadow", b)
}

// Outline returns whether outline effect is applied.
func (rp *RunProperties) Outline() bool {
	return rp.hasOnOffElement("outline")
}

// SetOutline sets outline effect.
func (rp *RunProperties) SetOutline(b bool) {
	rp.setOnOffElement("outline", b)
}

// NoProof returns whether spell/grammar checking is suppressed.
func (rp *RunProperties) NoProof() bool {
	return rp.hasOnOffElement("noProof")
}

// SetNoProof sets whether to suppress spell/grammar checking.
func (rp *RunProperties) SetNoProof(
	noProof bool,
) {
	rp.setOnOffElement("noProof", noProof)
}

// WebHidden returns whether the text is hidden in web layout.
func (rp *RunProperties) WebHidden() bool {
	return rp.hasOnOffElement("webHidden")
}

// SetWebHidden sets whether the text is hidden in web layout.
func (rp *RunProperties) SetWebHidden(
	hidden bool,
) {
	rp.setOnOffElement("webHidden", hidden)
}

// CharSpacing returns the character spacing in twips.
func (rp *RunProperties) CharSpacing() int {
	elem := rp.GetElement("spacing", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetCharSpacing sets the character spacing in twips.
func (rp *RunProperties) SetCharSpacing(
	twips int,
) {
	if twips == 0 {
		rp.removeElement("spacing")
		return
	}
	elem := rp.getOrCreateElement("spacing")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Position returns the vertical text position (raise/lower) in half-points.
func (rp *RunProperties) Position() int {
	elem := rp.GetElement(
		"position",
		NamespaceWML,
	)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetPosition sets the vertical text position (raise/lower) in half-points.
func (rp *RunProperties) SetPosition(
	halfPoints int,
) {
	if halfPoints == 0 {
		rp.removeElement("position")
		return
	}
	elem := rp.getOrCreateElement("position")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(halfPoints),
		),
	)
}

// CharacterWidth returns the character scale percentage (50-600).
func (rp *RunProperties) CharacterWidth() int {
	elem := rp.GetElement("w", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetCharacterWidth sets the character scale percentage (50-600).
func (rp *RunProperties) SetCharacterWidth(
	percent int,
) {
	if percent == 0 || percent == 100 {
		rp.removeElement("w")
		return
	}
	elem := rp.getOrCreateElement("w")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(percent),
		),
	)
}

// Kerning returns the auto-kern threshold in half-points.
func (rp *RunProperties) Kerning() int {
	elem := rp.GetElement("kern", NamespaceWML)
	if elem == nil {
		return 0
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetKerning sets the auto-kern threshold in half-points.
func (rp *RunProperties) SetKerning(
	halfPoints int,
) {
	if halfPoints <= 0 {
		rp.removeElement("kern")
		return
	}
	elem := rp.getOrCreateElement("kern")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(halfPoints),
		),
	)
}

// RunStyleId returns the run style reference ID.
// Note: This is an alias for RunStyle() for API consistency.
func (rp *RunProperties) RunStyleId() string {
	return rp.RunStyle()
}

// SetRunStyleId sets the run style reference ID.
// Note: This is an alias for SetRunStyle() for API consistency.
func (rp *RunProperties) SetRunStyleId(
	id string,
) {
	rp.SetRunStyle(id)
}

// RTL returns whether right-to-left text direction is applied.
func (rp *RunProperties) RTL() bool {
	return rp.hasOnOffElement("rtl")
}

// SetRTL sets whether right-to-left text direction is applied.
func (rp *RunProperties) SetRTL(rtl bool) {
	rp.setOnOffElement("rtl", rtl)
}

// ComplexScript returns whether complex script formatting is applied.
func (rp *RunProperties) ComplexScript() bool {
	return rp.hasOnOffElement("cs")
}

// SetComplexScript sets whether complex script formatting is applied.
func (rp *RunProperties) SetComplexScript(
	cs bool,
) {
	rp.setOnOffElement("cs", cs)
}

// EmphasisMark returns the emphasis mark type.
func (rp *RunProperties) EmphasisMark() EmphasisMarkValue {
	elem := rp.GetElement("em", NamespaceWML)
	if elem == nil {
		return EmphasisNone
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return EmphasisNone
	}
	return EmphasisMarkValue(attr.Value())
}

// SetEmphasisMark sets the emphasis mark type.
func (rp *RunProperties) SetEmphasisMark(
	em EmphasisMarkValue,
) {
	if em == EmphasisNone {
		rp.removeElement("em")
		return
	}
	elem := rp.getOrCreateElement("em")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(em),
		),
	)
}

// Helper methods

func (rp *RunProperties) hasOnOffElement(
	name string,
) bool {
	elem := rp.GetElement(name, NamespaceWML)
	if elem == nil {
		return false
	}
	// Check for explicit val="false" or val="0"
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if found {
		val := attr.Value()
		return val != "false" && val != "0" &&
			val != "off"
	}
	return true
}

func (rp *RunProperties) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		rp.getOrCreateElement(name)
	} else {
		rp.removeElement(name)
	}
}

func (rp *RunProperties) getOrCreateElement(
	name string,
) openxml.Element {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	rp.AppendChild(newElem)
	return newElem
}

func (rp *RunProperties) getOrCreateRunFonts() *RunFonts {
	elem := rp.GetElement("rFonts", NamespaceWML)
	if elem != nil {
		if rf, ok := elem.(*RunFonts); ok {
			return rf
		}
		if comp, ok := elem.(*openxml.CompositeElementBase); ok {
			return &RunFonts{
				CompositeElementBase: comp,
			}
		}
	}
	rf := NewRunFonts()
	rp.AppendChild(rf)
	return rf
}

func (rp *RunProperties) removeElement(
	name string,
) {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		rp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this RunProperties element.
func (rp *RunProperties) Clone() openxml.Element {
	return &RunProperties{
		CompositeElementBase: rp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RunProperties element.
func (rp *RunProperties) CloneNode(
	deep bool,
) openxml.Element {
	return &RunProperties{
		CompositeElementBase: rp.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// RunFonts represents font settings for a run (w:rFonts).
type RunFonts struct {
	*openxml.CompositeElementBase
}

// NewRunFonts creates a new RunFonts element.
func NewRunFonts() *RunFonts {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"rFonts",
		PrefixW,
	)
	return &RunFonts{CompositeElementBase: elem}
}

// ASCII returns the ASCII font name.
func (rf *RunFonts) ASCII() string {
	attr, found := rf.GetAttribute(
		"ascii",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetASCII sets the ASCII font name.
func (rf *RunFonts) SetASCII(fontName string) {
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"ascii",
			PrefixW,
			fontName,
		),
	)
}

// HAnsi returns the high ANSI font name.
func (rf *RunFonts) HAnsi() string {
	attr, found := rf.GetAttribute(
		"hAnsi",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetHAnsi sets the high ANSI font name.
func (rf *RunFonts) SetHAnsi(fontName string) {
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hAnsi",
			PrefixW,
			fontName,
		),
	)
}

// EastAsia returns the East Asian font name.
func (rf *RunFonts) EastAsia() string {
	attr, found := rf.GetAttribute(
		"eastAsia",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetEastAsia sets the East Asian font name.
func (rf *RunFonts) SetEastAsia(fontName string) {
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"eastAsia",
			PrefixW,
			fontName,
		),
	)
}

// ComplexScript returns the complex script font name.
func (rf *RunFonts) ComplexScript() string {
	attr, found := rf.GetAttribute(
		"cs",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetComplexScript sets the complex script font name.
func (rf *RunFonts) SetComplexScript(
	fontName string,
) {
	rf.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"cs",
			PrefixW,
			fontName,
		),
	)
}

// Clone creates a deep copy of this RunFonts element.
func (rf *RunFonts) Clone() openxml.Element {
	return &RunFonts{
		CompositeElementBase: rf.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RunFonts element.
func (rf *RunFonts) CloneNode(
	deep bool,
) openxml.Element {
	return &RunFonts{
		CompositeElementBase: rf.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
