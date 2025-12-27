//nolint:revive // file-length-limit: comprehensive style definitions with many related types
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// StyleType represents the type of a style.
type StyleType string

const (
	// StyleTypeParagraph is a paragraph style.
	StyleTypeParagraph StyleType = "paragraph"
	// StyleTypeCharacter is a character style.
	StyleTypeCharacter StyleType = "character"
	// StyleTypeTable is a table style.
	StyleTypeTable StyleType = "table"
	// StyleTypeNumbering is a numbering style.
	StyleTypeNumbering StyleType = "numbering"
)

// Built-in style IDs.
const (
	// StyleIdNormal is the Normal paragraph style.
	StyleIdNormal = "Normal"
	// StyleIdHeading1 is the Heading 1 style.
	StyleIdHeading1 = "Heading1"
	// StyleIdHeading2 is the Heading 2 style.
	StyleIdHeading2 = "Heading2"
	// StyleIdHeading3 is the Heading 3 style.
	StyleIdHeading3 = "Heading3"
	// StyleIdHeading4 is the Heading 4 style.
	StyleIdHeading4 = "Heading4"
	// StyleIdHeading5 is the Heading 5 style.
	StyleIdHeading5 = "Heading5"
	// StyleIdHeading6 is the Heading 6 style.
	StyleIdHeading6 = "Heading6"
	// StyleIdHeading7 is the Heading 7 style.
	StyleIdHeading7 = "Heading7"
	// StyleIdHeading8 is the Heading 8 style.
	StyleIdHeading8 = "Heading8"
	// StyleIdHeading9 is the Heading 9 style.
	StyleIdHeading9 = "Heading9"
	// StyleIdTitle is the Title style.
	StyleIdTitle = "Title"
	// StyleIdSubtitle is the Subtitle style.
	StyleIdSubtitle = "Subtitle"
	// StyleIdNoSpacing is the No Spacing style.
	StyleIdNoSpacing = "NoSpacing"
	// StyleIdQuote is the Quote style.
	StyleIdQuote = "Quote"
	// StyleIdIntenseQuote is the Intense Quote style.
	StyleIdIntenseQuote = "IntenseQuote"
	// StyleIdListParagraph is the List Paragraph style.
	StyleIdListParagraph = "ListParagraph"
	// StyleIdTOCHeading is the TOC Heading style.
	StyleIdTOCHeading = "TOCHeading"
	// StyleIdTOC1 is the TOC 1 style.
	StyleIdTOC1 = "TOC1"
	// StyleIdTOC2 is the TOC 2 style.
	StyleIdTOC2 = "TOC2"
	// StyleIdTOC3 is the TOC 3 style.
	StyleIdTOC3 = "TOC3"
	// StyleIdCaption is the Caption style.
	StyleIdCaption = "Caption"
	// StyleIdFootnoteText is the Footnote Text style.
	StyleIdFootnoteText = "FootnoteText"
	// StyleIdEndnoteText is the Endnote Text style.
	StyleIdEndnoteText = "EndnoteText"

	// Character styles
	// StyleIdDefaultParagraphFont is the Default Paragraph Font style.
	StyleIdDefaultParagraphFont = "DefaultParagraphFont"
	// StyleIdHyperlink is the Hyperlink style.
	StyleIdHyperlink = "Hyperlink"
	// StyleIdFollowedHyperlink is the Followed Hyperlink style.
	StyleIdFollowedHyperlink = "FollowedHyperlink"
	// StyleIdStrong is the Strong (bold) style.
	StyleIdStrong = "Strong"
	// StyleIdEmphasis is the Emphasis (italic) style.
	StyleIdEmphasis = "Emphasis"
	// StyleIdSubtleEmphasis is the Subtle Emphasis style.
	StyleIdSubtleEmphasis = "SubtleEmphasis"
	// StyleIdIntenseEmphasis is the Intense Emphasis style.
	StyleIdIntenseEmphasis = "IntenseEmphasis"
	// StyleIdSubtleReference is the Subtle Reference style.
	StyleIdSubtleReference = "SubtleReference"
	// StyleIdIntenseReference is the Intense Reference style.
	StyleIdIntenseReference = "IntenseReference"
	// StyleIdBookTitle is the Book Title style.
	StyleIdBookTitle = "BookTitle"
	// StyleIdFootnoteReference is the Footnote Reference style.
	StyleIdFootnoteReference = "FootnoteReference"
	// StyleIdEndnoteReference is the Endnote Reference style.
	StyleIdEndnoteReference = "EndnoteReference"

	// Table styles
	// StyleIdTableNormal is the Table Normal style.
	StyleIdTableNormal = "TableNormal"
	// StyleIdTableGrid is the Table Grid style.
	StyleIdTableGrid = "TableGrid"
)

// Style-related constants for magic numbers and attribute names.
const (
	maxOutlineLevel      = 9
	defaultStylePriority = 99
	maxHeadingLevel      = 8
	attrValueOne         = "1"
)

// Style represents an individual style definition (w:style).
type Style struct {
	*openxml.CompositeElementBase
}

// NewStyle creates a new Style element with the given ID and type.
func NewStyle(
	id string,
	styleType StyleType,
) *Style {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"style",
		PrefixW,
	)
	s := &Style{CompositeElementBase: elem}
	s.SetStyleId(id)
	s.SetType(styleType)

	return s
}

// NewParagraphStyle creates a new paragraph style with the given ID and name.
func NewParagraphStyle(id, name string) *Style {
	s := NewStyle(id, StyleTypeParagraph)
	s.SetStyleName(name)

	return s
}

// NewCharacterStyle creates a new character style with the given ID and name.
func NewCharacterStyle(id, name string) *Style {
	s := NewStyle(id, StyleTypeCharacter)
	s.SetStyleName(name)

	return s
}

// NewTableStyle creates a new table style with the given ID and name.
func NewTableStyle(id, name string) *Style {
	s := NewStyle(id, StyleTypeTable)
	s.SetStyleName(name)

	return s
}

// NewHeadingStyle creates a heading style for the specified level (1-9).
func NewHeadingStyle(level int) *Style {
	if level < 1 {
		level = 1
	}
	if level > maxOutlineLevel {
		level = maxOutlineLevel
	}

	id := "Heading" + strconv.Itoa(level)
	name := "Heading " + strconv.Itoa(level)

	s := NewParagraphStyle(id, name)
	s.SetBasedOn(StyleIdNormal)

	// Set outline level (0-indexed)
	props := s.GetOrCreateStyleParagraphProperties()
	props.SetOutlineLevel(level - 1)

	// Set quick format
	s.SetQuickFormat(true)

	// Set UI priority (headings are 9 + level)
	s.SetUIPriority(9 + level)

	return s
}

// StyleId returns the style ID.
func (s *Style) StyleId() string {
	attr, found := s.GetAttribute(
		"styleId",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetStyleId sets the style ID.
func (s *Style) SetStyleId(id string) {
	s.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"styleId",
			PrefixW,
			id,
		),
	)
}

// Type returns the style type.
func (s *Style) Type() StyleType {
	attr, found := s.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return StyleTypeParagraph
	}

	return StyleType(attr.Value())
}

// SetType sets the style type.
func (s *Style) SetType(t StyleType) {
	s.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"type",
			PrefixW,
			string(t),
		),
	)
}

// StyleName returns the human-readable style name.
func (s *Style) StyleName() string {
	elem := s.GetElement("name", NamespaceWML)
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

// SetStyleName sets the human-readable style name.
func (s *Style) SetStyleName(name string) {
	if name == "" {
		s.removeElement("name")

		return
	}
	elem := s.getOrCreateElement("name")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			name,
		),
	)
}

// BasedOn returns the parent style ID.
func (s *Style) BasedOn() string {
	elem := s.GetElement("basedOn", NamespaceWML)
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

// SetBasedOn sets the parent style ID.
func (s *Style) SetBasedOn(id string) {
	if id == "" {
		s.removeElement("basedOn")

		return
	}
	elem := s.getOrCreateElement("basedOn")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			id,
		),
	)
}

// NextParagraphStyle returns the style ID for the next paragraph.
func (s *Style) NextParagraphStyle() string {
	elem := s.GetElement("next", NamespaceWML)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetNextParagraphStyle sets the style ID for the next paragraph.
func (s *Style) SetNextParagraphStyle(id string) {
	if id == "" {
		s.removeElement("next")

		return
	}
	elem := s.getOrCreateElement("next")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			id,
		),
	)
}

// LinkedStyle returns the linked style ID.
func (s *Style) LinkedStyle() string {
	elem := s.GetElement("link", NamespaceWML)
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

// SetLinkedStyle sets the linked style ID.
func (s *Style) SetLinkedStyle(id string) {
	if id == "" {
		s.removeElement("link")

		return
	}
	elem := s.getOrCreateElement("link")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			id,
		),
	)
}

// Default returns whether this is the default style of its type.
func (s *Style) Default() bool {
	attr, found := s.GetAttribute(
		"default",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == "true" ||
		val == "on"
}

// SetDefault sets whether this is the default style of its type.
func (s *Style) SetDefault(b bool) {
	if b {
		s.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"default",
				PrefixW,
				"1",
			),
		)
	} else {
		s.RemoveAttribute("default", NamespaceWML)
	}
}

// CustomStyle returns whether this is a user-defined style.
func (s *Style) CustomStyle() bool {
	attr, found := s.GetAttribute(
		"customStyle",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == "true" ||
		val == "on"
}

// SetCustomStyle sets whether this is a user-defined style.
func (s *Style) SetCustomStyle(b bool) {
	if b {
		s.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"customStyle",
				PrefixW,
				attrValueOne,
			),
		)
	} else {
		s.RemoveAttribute("customStyle", NamespaceWML)
	}
}

// StyleParagraphProperties returns the paragraph properties for this style.
func (s *Style) StyleParagraphProperties() *StyleParagraphProperties {
	elem := s.GetElement("pPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*StyleParagraphProperties); ok {
		return pp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &StyleParagraphProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateStyleParagraphProperties returns the paragraph properties, creating if needed.
func (s *Style) GetOrCreateStyleParagraphProperties() *StyleParagraphProperties {
	pp := s.StyleParagraphProperties()
	if pp != nil {
		return pp
	}
	pp = NewStyleParagraphProperties()
	s.AppendChild(pp)

	return pp
}

// StyleRunProperties returns the run properties for this style.
func (s *Style) StyleRunProperties() *StyleRunProperties {
	elem := s.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*StyleRunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &StyleRunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateStyleRunProperties returns the run properties, creating if needed.
func (s *Style) GetOrCreateStyleRunProperties() *StyleRunProperties {
	rp := s.StyleRunProperties()
	if rp != nil {
		return rp
	}
	rp = NewStyleRunProperties()
	s.AppendChild(rp)

	return rp
}

// StyleTableProperties returns the table properties for this style.
func (s *Style) StyleTableProperties() *StyleTableProperties {
	elem := s.GetElement("tblPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tp, ok := elem.(*StyleTableProperties); ok {
		return tp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &StyleTableProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateStyleTableProperties returns the table properties, creating if needed.
func (s *Style) GetOrCreateStyleTableProperties() *StyleTableProperties {
	tp := s.StyleTableProperties()
	if tp != nil {
		return tp
	}
	tp = NewStyleTableProperties()
	s.AppendChild(tp)

	return tp
}

// UIPriority returns the sort order in the style gallery.
func (s *Style) UIPriority() int {
	elem := s.GetElement(
		"uiPriority",
		NamespaceWML,
	)
	if elem == nil {
		return defaultStylePriority // Default priority
	}
	attr, found := elem.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return defaultStylePriority
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return defaultStylePriority
	}

	return val
}

// SetUIPriority sets the sort order in the style gallery.
func (s *Style) SetUIPriority(priority int) {
	elem := s.getOrCreateElement("uiPriority")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(priority),
		),
	)
}

// QuickFormat returns whether the style appears in quick styles.
func (s *Style) QuickFormat() bool {
	return s.hasOnOffElement("qFormat")
}

// SetQuickFormat sets whether the style appears in quick styles.
func (s *Style) SetQuickFormat(b bool) {
	s.setOnOffElement("qFormat", b)
}

// SemiHidden returns whether the style is hidden from UI.
func (s *Style) SemiHidden() bool {
	return s.hasOnOffElement("semiHidden")
}

// SetSemiHidden sets whether the style is hidden from UI.
func (s *Style) SetSemiHidden(b bool) {
	s.setOnOffElement("semiHidden", b)
}

// UnhideWhenUsed returns whether the style becomes visible when used.
func (s *Style) UnhideWhenUsed() bool {
	return s.hasOnOffElement("unhideWhenUsed")
}

// SetUnhideWhenUsed sets whether the style becomes visible when used.
func (s *Style) SetUnhideWhenUsed(b bool) {
	s.setOnOffElement("unhideWhenUsed", b)
}

// Helper methods

func (s *Style) hasOnOffElement(
	name string,
) bool {
	elem := s.GetElement(name, NamespaceWML)
	if elem == nil {
		return false
	}
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

//nolint:revive // bool parameter is intentional for on/off toggle
func (s *Style) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		s.getOrCreateElement(name)
	} else {
		s.removeElement(name)
	}
}

func (s *Style) getOrCreateElement(
	name string,
) openxml.Element {
	elem := s.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	s.AppendChild(newElem)

	return newElem
}

func (s *Style) removeElement(name string) {
	elem := s.GetElement(name, NamespaceWML)
	if elem != nil {
		s.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Style element.
func (s *Style) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Style{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Style element.
func (s *Style) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &Style{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// StyleParagraphProperties represents paragraph properties within a style (w:pPr).
type StyleParagraphProperties struct {
	*openxml.CompositeElementBase
}

// NewStyleParagraphProperties creates a new StyleParagraphProperties element.
func NewStyleParagraphProperties() *StyleParagraphProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pPr",
		PrefixW,
	)

	return &StyleParagraphProperties{
		CompositeElementBase: elem,
	}
}

// SetJustification sets the paragraph alignment.
func (pp *StyleParagraphProperties) SetJustification(
	j JustificationValue,
) {
	if j == "" || j == JustificationLeft {
		pp.removeElement("jc")

		return
	}
	elem := pp.getOrCreateElement("jc")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(j),
		),
	)
}

// SetOutlineLevel sets the outline level (0-8).
func (pp *StyleParagraphProperties) SetOutlineLevel(
	level int,
) {
	if level < 0 || level > maxHeadingLevel {
		pp.removeElement("outlineLvl")

		return
	}
	elem := pp.getOrCreateElement("outlineLvl")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(level),
		),
	)
}

// SetKeepNext sets whether to keep with the next paragraph.
func (pp *StyleParagraphProperties) SetKeepNext(
	b bool,
) {
	pp.setOnOffElement("keepNext", b)
}

// SetKeepLines sets whether to keep lines together.
func (pp *StyleParagraphProperties) SetKeepLines(
	b bool,
) {
	pp.setOnOffElement("keepLines", b)
}

// SetPageBreakBefore sets whether to insert a page break before.
func (pp *StyleParagraphProperties) SetPageBreakBefore(
	b bool,
) {
	pp.setOnOffElement("pageBreakBefore", b)
}

// SetSpacingBefore sets the space before the paragraph in twips.
func (pp *StyleParagraphProperties) SetSpacingBefore(
	twips int,
) {
	elem := pp.getOrCreateElement("spacing")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"before",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// SetSpacingAfter sets the space after the paragraph in twips.
func (pp *StyleParagraphProperties) SetSpacingAfter(
	twips int,
) {
	elem := pp.getOrCreateElement("spacing")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"after",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// SetLeftIndent sets the left indentation in twips.
func (pp *StyleParagraphProperties) SetLeftIndent(
	twips int,
) {
	elem := pp.getOrCreateElement("ind")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"left",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// SetHangingIndent sets the hanging indentation in twips.
func (pp *StyleParagraphProperties) SetHangingIndent(
	twips int,
) {
	elem := pp.getOrCreateElement("ind")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hanging",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

//
//nolint:revive // bool parameter is intentional for on/off toggle
func (pp *StyleParagraphProperties) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		pp.getOrCreateElement(name)
	} else {
		pp.removeElement(name)
	}
}

func (pp *StyleParagraphProperties) getOrCreateElement(
	name string,
) openxml.Element {
	elem := pp.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	pp.AppendChild(newElem)

	return newElem
}

func (pp *StyleParagraphProperties) removeElement(
	name string,
) {
	elem := pp.GetElement(name, NamespaceWML)
	if elem != nil {
		pp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this StyleParagraphProperties element.
func (pp *StyleParagraphProperties) Clone() openxml.Element {
	return &StyleParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// StyleRunProperties represents run properties within a style (w:rPr).
type StyleRunProperties struct {
	*openxml.CompositeElementBase
}

// NewStyleRunProperties creates a new StyleRunProperties element.
func NewStyleRunProperties() *StyleRunProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"rPr",
		PrefixW,
	)

	return &StyleRunProperties{
		CompositeElementBase: elem,
	}
}

// SetBold sets bold formatting.
func (rp *StyleRunProperties) SetBold(b bool) {
	rp.setOnOffElement("b", b)
}

// SetItalic sets italic formatting.
func (rp *StyleRunProperties) SetItalic(b bool) {
	rp.setOnOffElement("i", b)
}

// SetUnderline sets the underline style.
func (rp *StyleRunProperties) SetUnderline(
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

// SetFontSize sets the font size in half-points.
func (rp *StyleRunProperties) SetFontSize(
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

// SetFontSizeComplexScript sets the complex script font size in half-points.
func (rp *StyleRunProperties) SetFontSizeComplexScript(
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

// SetFont sets the font name for all script types.
func (rp *StyleRunProperties) SetFont(
	fontName string,
) {
	rf := rp.getOrCreateElement("rFonts")
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

// SetColor sets the text color from a hex string (without # prefix).
func (rp *StyleRunProperties) SetColor(
	hex string,
) {
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

// SetCaps sets all caps formatting.
func (rp *StyleRunProperties) SetCaps(b bool) {
	rp.setOnOffElement("caps", b)
}

// SetSmallCaps sets small caps formatting.
func (rp *StyleRunProperties) SetSmallCaps(
	b bool,
) {
	rp.setOnOffElement("smallCaps", b)
}

// SetStrike sets strikethrough formatting.
func (rp *StyleRunProperties) SetStrike(b bool) {
	rp.setOnOffElement("strike", b)
}

//nolint:revive // bool parameter is intentional for on/off toggle
func (rp *StyleRunProperties) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		rp.getOrCreateElement(name)
	} else {
		rp.removeElement(name)
	}
}

func (rp *StyleRunProperties) getOrCreateElement(
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

func (rp *StyleRunProperties) removeElement(
	name string,
) {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		rp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this StyleRunProperties element.
func (rp *StyleRunProperties) Clone() openxml.Element {
	return &StyleRunProperties{
		CompositeElementBase: rp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// StyleTableProperties represents table properties within a style (w:tblPr).
type StyleTableProperties struct {
	*openxml.CompositeElementBase
}

// NewStyleTableProperties creates a new StyleTableProperties element.
func NewStyleTableProperties() *StyleTableProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblPr",
		PrefixW,
	)

	return &StyleTableProperties{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this StyleTableProperties element.
func (tp *StyleTableProperties) Clone() openxml.Element {
	return &StyleTableProperties{
		CompositeElementBase: tp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
