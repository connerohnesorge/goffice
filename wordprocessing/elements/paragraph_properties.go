//nolint:revive // file-length-limit - this file contains all paragraph properties
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ParagraphProperties represents paragraph formatting properties (w:pPr).
type ParagraphProperties struct {
	*openxml.CompositeElementBase
}

// NewParagraphProperties creates a new ParagraphProperties element.
func NewParagraphProperties() *ParagraphProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pPr",
		PrefixW,
	)

	return &ParagraphProperties{
		CompositeElementBase: elem,
	}
}

// ParagraphStyleId returns the paragraph style ID.
func (pp *ParagraphProperties) ParagraphStyleId() string {
	elem := pp.GetElement("pStyle", NamespaceWML)
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

// SetParagraphStyleId sets the paragraph style ID.
func (pp *ParagraphProperties) SetParagraphStyleId(
	id string,
) {
	if id == "" {
		pp.removeElement("pStyle")

		return
	}
	elem := pp.getOrCreateElement("pStyle")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			id,
		),
	)
}

// Justification returns the paragraph alignment.
func (pp *ParagraphProperties) Justification() JustificationValue {
	elem := pp.GetElement("jc", NamespaceWML)
	if elem == nil {
		return JustificationLeft
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return JustificationLeft
	}

	return JustificationValue(attr.Value())
}

// SetJustification sets the paragraph alignment.
func (pp *ParagraphProperties) SetJustification(
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

// Indentation returns the indentation settings.
func (pp *ParagraphProperties) Indentation() *Indentation {
	elem := pp.GetElement("ind", NamespaceWML)
	if elem == nil {
		return nil
	}
	if ind, ok := elem.(*Indentation); ok {
		return ind
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Indentation{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetIndentation sets the indentation settings.
func (pp *ParagraphProperties) SetIndentation(
	ind *Indentation,
) {
	// Remove existing
	pp.removeElement("ind")
	if ind != nil {
		pp.AppendChild(ind)
	}
}

// GetOrCreateIndentation returns the indentation settings, creating if needed.
func (pp *ParagraphProperties) GetOrCreateIndentation() *Indentation {
	ind := pp.Indentation()
	if ind != nil {
		return ind
	}
	ind = NewIndentation()
	pp.AppendChild(ind)

	return ind
}

// SpacingBetweenLines returns the spacing settings.
func (pp *ParagraphProperties) SpacingBetweenLines() *SpacingBetweenLines {
	elem := pp.GetElement("spacing", NamespaceWML)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*SpacingBetweenLines); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SpacingBetweenLines{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSpacingBetweenLines returns the spacing settings, creating if needed.
func (pp *ParagraphProperties) GetOrCreateSpacingBetweenLines() *SpacingBetweenLines {
	sp := pp.SpacingBetweenLines()
	if sp != nil {
		return sp
	}
	sp = NewSpacingBetweenLines()
	pp.AppendChild(sp)

	return sp
}

// KeepNext returns whether the paragraph keeps with the next paragraph.
func (pp *ParagraphProperties) KeepNext() bool {
	return pp.hasOnOffElement("keepNext")
}

// SetKeepNext sets whether the paragraph keeps with the next paragraph.
func (pp *ParagraphProperties) SetKeepNext(
	b bool,
) {
	pp.setOnOffElement("keepNext", b)
}

// KeepLines returns whether all lines in the paragraph stay together.
func (pp *ParagraphProperties) KeepLines() bool {
	return pp.hasOnOffElement("keepLines")
}

// SetKeepLines sets whether all lines in the paragraph stay together.
func (pp *ParagraphProperties) SetKeepLines(
	b bool,
) {
	pp.setOnOffElement("keepLines", b)
}

// PageBreakBefore returns whether a page break precedes the paragraph.
func (pp *ParagraphProperties) PageBreakBefore() bool {
	return pp.hasOnOffElement("pageBreakBefore")
}

// SetPageBreakBefore sets whether a page break precedes the paragraph.
func (pp *ParagraphProperties) SetPageBreakBefore(
	b bool,
) {
	pp.setOnOffElement("pageBreakBefore", b)
}

// WidowControl returns whether widow/orphan control is enabled.
func (pp *ParagraphProperties) WidowControl() bool {
	// Default is true, so we check for explicit false
	elem := pp.GetElement(
		"widowControl",
		NamespaceWML,
	)
	if elem == nil {
		return true // default
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return true // element present with no val means true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetWidowControl sets whether widow/orphan control is enabled.
func (pp *ParagraphProperties) SetWidowControl(
	b bool,
) {
	if b {
		pp.removeElement(
			"widowControl",
		) // true is default
	} else {
		elem := pp.getOrCreateElement("widowControl")
		elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, attrValueFalse))
	}
}

// NumberingProperties returns the numbering (list) properties.
func (pp *ParagraphProperties) NumberingProperties() *NumberingProperties {
	elem := pp.GetElement("numPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if np, ok := elem.(*NumberingProperties); ok {
		return np
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NumberingProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNumberingProperties returns the numbering properties, creating if needed.
func (pp *ParagraphProperties) GetOrCreateNumberingProperties() *NumberingProperties {
	np := pp.NumberingProperties()
	if np != nil {
		return np
	}
	np = NewNumberingProperties()
	pp.AppendChild(np)

	return np
}

// Tabs returns the tab stop settings.
func (pp *ParagraphProperties) Tabs() *Tabs {
	elem := pp.GetElement("tabs", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tabs, ok := elem.(*Tabs); ok {
		return tabs
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Tabs{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateTabs returns the tab stop settings, creating if needed.
func (pp *ParagraphProperties) GetOrCreateTabs() *Tabs {
	tabs := pp.Tabs()
	if tabs != nil {
		return tabs
	}
	tabs = NewTabs()
	pp.AppendChild(tabs)

	return tabs
}

// ParagraphBorders returns the paragraph border settings.
func (pp *ParagraphProperties) ParagraphBorders() *ParagraphBorders {
	elem := pp.GetElement("pBdr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pb, ok := elem.(*ParagraphBorders); ok {
		return pb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ParagraphBorders{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateParagraphBorders returns the paragraph borders, creating if needed.
func (pp *ParagraphProperties) GetOrCreateParagraphBorders() *ParagraphBorders {
	pb := pp.ParagraphBorders()
	if pb != nil {
		return pb
	}
	pb = NewParagraphBorders()
	pp.AppendChild(pb)

	return pb
}

// Shading returns the paragraph shading settings.
func (pp *ParagraphProperties) Shading() *Shading {
	elem := pp.GetElement("shd", NamespaceWML)
	if elem == nil {
		return nil
	}
	if shd, ok := elem.(*Shading); ok {
		return shd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Shading{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateShading returns the paragraph shading, creating if needed.
func (pp *ParagraphProperties) GetOrCreateShading() *Shading {
	shd := pp.Shading()
	if shd != nil {
		return shd
	}
	shd = NewShading()
	pp.AppendChild(shd)

	return shd
}

// SectionProperties returns the section properties element if present (for section breaks).
func (pp *ParagraphProperties) SectionProperties() *SectionProperties {
	elem := pp.GetElement("sectPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*SectionProperties); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SectionProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// OutlineLevel returns the outline level (0-8, or -1 if not set).
func (pp *ParagraphProperties) OutlineLevel() int {
	elem := pp.GetElement(
		"outlineLvl",
		NamespaceWML,
	)
	if elem == nil {
		return -1
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return -1
	}
	val, err := strconv.Atoi(attr.Value())
	if err != nil {
		return -1
	}

	return val
}

// SetOutlineLevel sets the outline level (0-8).
func (pp *ParagraphProperties) SetOutlineLevel(
	level int,
) {
	if level < 0 || level > 8 {
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

// ParagraphMarkRunProperties returns the run properties for the paragraph mark.
func (pp *ParagraphProperties) ParagraphMarkRunProperties() *RunProperties {
	elem := pp.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*RunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SuppressAutoHyphens returns whether automatic hyphenation is suppressed.
func (pp *ParagraphProperties) SuppressAutoHyphens() bool {
	return pp.hasOnOffElement(
		"suppressAutoHyphens",
	)
}

// SetSuppressAutoHyphens sets whether automatic hyphenation is suppressed.
func (pp *ParagraphProperties) SetSuppressAutoHyphens(
	suppress bool,
) {
	pp.setOnOffElement(
		"suppressAutoHyphens",
		suppress,
	)
}

// Bidi returns whether the paragraph contains bidirectional text.
func (pp *ParagraphProperties) Bidi() bool {
	return pp.hasOnOffElement("bidi")
}

// SetBidi sets whether the paragraph contains bidirectional text.
func (pp *ParagraphProperties) SetBidi(
	bidi bool,
) {
	pp.setOnOffElement("bidi", bidi)
}

// ContextualSpacing returns whether spacing above/below is ignored for identical styles.
func (pp *ParagraphProperties) ContextualSpacing() bool {
	return pp.hasOnOffElement("contextualSpacing")
}

// SetContextualSpacing sets whether spacing above/below is ignored for identical styles.
func (pp *ParagraphProperties) SetContextualSpacing(
	contextual bool,
) {
	pp.setOnOffElement(
		"contextualSpacing",
		contextual,
	)
}

// MirrorIndents returns whether indents are mirrored on facing pages.
func (pp *ParagraphProperties) MirrorIndents() bool {
	return pp.hasOnOffElement("mirrorIndents")
}

// SetMirrorIndents sets whether indents are mirrored on facing pages.
func (pp *ParagraphProperties) SetMirrorIndents(
	mirror bool,
) {
	pp.setOnOffElement("mirrorIndents", mirror)
}

// SnapToGrid returns whether the paragraph snaps to the document grid.
func (pp *ParagraphProperties) SnapToGrid() bool {
	// Default is true, so we check for explicit false
	elem := pp.GetElement(
		"snapToGrid",
		NamespaceWML,
	)
	if elem == nil {
		return true // default
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return true // element present with no val means true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetSnapToGrid sets whether the paragraph snaps to the document grid.
func (pp *ParagraphProperties) SetSnapToGrid(
	snap bool,
) {
	if snap {
		pp.removeElement(
			"snapToGrid",
		) // true is default
	} else {
		elem := pp.getOrCreateElement("snapToGrid")
		elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, attrValueFalse))
	}
}

// TextAlignment returns the vertical text alignment within the line.
func (pp *ParagraphProperties) TextAlignment() TextAlignmentValue {
	elem := pp.GetElement(
		"textAlignment",
		NamespaceWML,
	)
	if elem == nil {
		return TextAlignmentAuto
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return TextAlignmentAuto
	}

	return TextAlignmentValue(attr.Value())
}

// SetTextAlignment sets the vertical text alignment within the line.
func (pp *ParagraphProperties) SetTextAlignment(
	align TextAlignmentValue,
) {
	if align == "" || align == TextAlignmentAuto {
		pp.removeElement("textAlignment")

		return
	}
	elem := pp.getOrCreateElement("textAlignment")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(align),
		),
	)
}

// WordWrap returns whether Latin text is allowed to wrap in the middle of a word.
func (pp *ParagraphProperties) WordWrap() bool {
	return pp.hasOnOffElement("wordWrap")
}

// SetWordWrap sets whether Latin text is allowed to wrap in the middle of a word.
func (pp *ParagraphProperties) SetWordWrap(
	wrap bool,
) {
	pp.setOnOffElement("wordWrap", wrap)
}

// OverflowPunct returns whether punctuation is allowed to overflow the line.
func (pp *ParagraphProperties) OverflowPunct() bool {
	// Default is true, so we check for explicit false
	elem := pp.GetElement(
		"overflowPunct",
		NamespaceWML,
	)
	if elem == nil {
		return true // default
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return true // element present with no val means true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetOverflowPunct sets whether punctuation is allowed to overflow the line.
func (pp *ParagraphProperties) SetOverflowPunct(
	overflow bool,
) {
	if overflow {
		pp.removeElement(
			"overflowPunct",
		) // true is default
	} else {
		elem := pp.getOrCreateElement("overflowPunct")
		elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, attrValueFalse))
	}
}

// Helper methods

func (pp *ParagraphProperties) hasOnOffElement(
	name string,
) bool {
	elem := pp.GetElement(name, NamespaceWML)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if found {
		val := attr.Value()

		return val != attrValueFalse &&
			val != attrValueZero &&
			val != attrValueOff
	}

	return true
}

func (pp *ParagraphProperties) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		pp.getOrCreateElement(name)
	} else {
		pp.removeElement(name)
	}
}

func (pp *ParagraphProperties) getOrCreateElement(
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

func (pp *ParagraphProperties) removeElement(
	name string,
) {
	elem := pp.GetElement(name, NamespaceWML)
	if elem != nil {
		pp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this ParagraphProperties element.
func (pp *ParagraphProperties) Clone() openxml.Element {
	return &ParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ParagraphProperties element.
func (pp *ParagraphProperties) CloneNode(
	deep bool,
) openxml.Element {
	return &ParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// Indentation represents paragraph indentation settings (w:ind).
type Indentation struct {
	*openxml.CompositeElementBase
}

// NewIndentation creates a new Indentation element.
func NewIndentation() *Indentation {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"ind",
		PrefixW,
	)

	return &Indentation{
		CompositeElementBase: elem,
	}
}

// Left returns the left indentation in twips.
func (ind *Indentation) Left() int {
	return ind.getIntAttribute("left")
}

// SetLeft sets the left indentation in twips.
func (ind *Indentation) SetLeft(twips int) {
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"left",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Right returns the right indentation in twips.
func (ind *Indentation) Right() int {
	return ind.getIntAttribute("right")
}

// SetRight sets the right indentation in twips.
func (ind *Indentation) SetRight(twips int) {
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"right",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// FirstLine returns the first line indentation in twips.
func (ind *Indentation) FirstLine() int {
	return ind.getIntAttribute("firstLine")
}

// SetFirstLine sets the first line indentation in twips.
func (ind *Indentation) SetFirstLine(twips int) {
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"firstLine",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Hanging returns the hanging indentation in twips.
func (ind *Indentation) Hanging() int {
	return ind.getIntAttribute("hanging")
}

// SetHanging sets the hanging indentation in twips.
func (ind *Indentation) SetHanging(twips int) {
	ind.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hanging",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

func (ind *Indentation) getIntAttribute(
	name string,
) int {
	attr, found := ind.GetAttribute(
		name,
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this Indentation element.
func (ind *Indentation) Clone() openxml.Element {
	return &Indentation{
		CompositeElementBase: ind.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// SpacingBetweenLines represents paragraph spacing settings (w:spacing).
type SpacingBetweenLines struct {
	*openxml.CompositeElementBase
}

// NewSpacingBetweenLines creates a new SpacingBetweenLines element.
func NewSpacingBetweenLines() *SpacingBetweenLines {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"spacing",
		PrefixW,
	)

	return &SpacingBetweenLines{
		CompositeElementBase: elem,
	}
}

// Before returns the space before paragraph in twips.
func (sp *SpacingBetweenLines) Before() int {
	return sp.getIntAttribute("before")
}

// SetBefore sets the space before paragraph in twips.
func (sp *SpacingBetweenLines) SetBefore(
	twips int,
) {
	sp.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"before",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// After returns the space after paragraph in twips.
func (sp *SpacingBetweenLines) After() int {
	return sp.getIntAttribute("after")
}

// SetAfter sets the space after paragraph in twips.
func (sp *SpacingBetweenLines) SetAfter(
	twips int,
) {
	sp.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"after",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Line returns the line spacing value.
func (sp *SpacingBetweenLines) Line() int {
	return sp.getIntAttribute("line")
}

// SetLine sets the line spacing value.
func (sp *SpacingBetweenLines) SetLine(
	value int,
) {
	sp.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"line",
			PrefixW,
			strconv.Itoa(value),
		),
	)
}

// LineRule returns the line spacing rule.
func (sp *SpacingBetweenLines) LineRule() LineSpacingRule {
	attr, found := sp.GetAttribute(
		"lineRule",
		NamespaceWML,
	)
	if !found {
		return LineSpacingAuto
	}

	return LineSpacingRule(attr.Value())
}

// SetLineRule sets the line spacing rule.
func (sp *SpacingBetweenLines) SetLineRule(
	rule LineSpacingRule,
) {
	sp.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"lineRule",
			PrefixW,
			string(rule),
		),
	)
}

func (sp *SpacingBetweenLines) getIntAttribute(
	name string,
) int {
	attr, found := sp.GetAttribute(
		name,
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this SpacingBetweenLines element.
func (sp *SpacingBetweenLines) Clone() openxml.Element {
	return &SpacingBetweenLines{
		CompositeElementBase: sp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// NumberingProperties represents paragraph numbering/list settings (w:numPr).
type NumberingProperties struct {
	*openxml.CompositeElementBase
}

// NewNumberingProperties creates a new NumberingProperties element.
func NewNumberingProperties() *NumberingProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"numPr",
		PrefixW,
	)

	return &NumberingProperties{
		CompositeElementBase: elem,
	}
}

// NumberingId returns the numbering definition ID.
func (np *NumberingProperties) NumberingId() int {
	elem := np.GetElement("numId", NamespaceWML)
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

// SetNumberingId sets the numbering definition ID.
func (np *NumberingProperties) SetNumberingId(
	numId int,
) {
	elem := np.getOrCreateElement("numId")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(numId),
		),
	)
}

// NumberingLevelReference returns the list level (0-8).
func (np *NumberingProperties) NumberingLevelReference() int {
	elem := np.GetElement("ilvl", NamespaceWML)
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

// SetNumberingLevel sets the list level (0-8).
func (np *NumberingProperties) SetNumberingLevel(
	level int,
) {
	elem := np.getOrCreateElement("ilvl")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(level),
		),
	)
}

func (np *NumberingProperties) getOrCreateElement(
	name string,
) openxml.Element {
	elem := np.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	np.AppendChild(newElem)

	return newElem
}

// Clone creates a deep copy of this NumberingProperties element.
func (np *NumberingProperties) Clone() openxml.Element {
	return &NumberingProperties{
		CompositeElementBase: np.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Tabs represents tab stop collection (w:tabs).
type Tabs struct {
	*openxml.CompositeElementBase
}

// NewTabs creates a new Tabs element.
func NewTabs() *Tabs {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tabs",
		PrefixW,
	)

	return &Tabs{CompositeElementBase: elem}
}

// AddTab adds a tab stop.
func (t *Tabs) AddTab(
	position int,
	align TabAlignment,
	leader TabLeader,
) *TabStop {
	tab := NewTabStop(position, align, leader)
	t.AppendChild(tab)

	return tab
}

// Clone creates a deep copy of this Tabs element.
func (t *Tabs) Clone() openxml.Element {
	return &Tabs{
		CompositeElementBase: t.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TabStop represents a single tab stop (w:tab).
type TabStop struct {
	*openxml.CompositeElementBase
}

// NewTabStop creates a new TabStop element.
func NewTabStop(
	position int,
	align TabAlignment,
	leader TabLeader,
) *TabStop {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tab",
		PrefixW,
	)
	tab := &TabStop{CompositeElementBase: elem}
	tab.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"pos",
			PrefixW,
			strconv.Itoa(position),
		),
	)
	tab.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(align),
		),
	)
	if leader != TabLeaderNone && leader != "" {
		tab.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"leader",
				PrefixW,
				string(leader),
			),
		)
	}

	return tab
}

// Position returns the tab position in twips.
func (ts *TabStop) Position() int {
	attr, found := ts.GetAttribute(
		"pos",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Alignment returns the tab alignment.
func (ts *TabStop) Alignment() TabAlignment {
	attr, found := ts.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return TabAlignLeft
	}

	return TabAlignment(attr.Value())
}

// Leader returns the tab leader character.
func (ts *TabStop) Leader() TabLeader {
	attr, found := ts.GetAttribute(
		"leader",
		NamespaceWML,
	)
	if !found {
		return TabLeaderNone
	}

	return TabLeader(attr.Value())
}

// Clone creates a deep copy of this TabStop element.
func (ts *TabStop) Clone() openxml.Element {
	return &TabStop{
		CompositeElementBase: ts.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ParagraphBorders represents paragraph border settings (w:pBdr).
type ParagraphBorders struct {
	*openxml.CompositeElementBase
}

// NewParagraphBorders creates a new ParagraphBorders element.
func NewParagraphBorders() *ParagraphBorders {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pBdr",
		PrefixW,
	)

	return &ParagraphBorders{
		CompositeElementBase: elem,
	}
}

// Top returns the top border.
func (pb *ParagraphBorders) Top() *Border {
	return pb.getBorder("top")
}

// SetTop sets the top border.
func (pb *ParagraphBorders) SetTop(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("top", style, size, color)
}

// Bottom returns the bottom border.
func (pb *ParagraphBorders) Bottom() *Border {
	return pb.getBorder("bottom")
}

// SetBottom sets the bottom border.
func (pb *ParagraphBorders) SetBottom(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("bottom", style, size, color)
}

// Left returns the left border.
func (pb *ParagraphBorders) Left() *Border {
	return pb.getBorder("left")
}

// SetLeft sets the left border.
func (pb *ParagraphBorders) SetLeft(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("left", style, size, color)
}

// Right returns the right border.
func (pb *ParagraphBorders) Right() *Border {
	return pb.getBorder("right")
}

// SetRight sets the right border.
func (pb *ParagraphBorders) SetRight(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("right", style, size, color)
}

// Between returns the border between paragraphs.
func (pb *ParagraphBorders) Between() *Border {
	return pb.getBorder("between")
}

// SetBetween sets the border between paragraphs.
func (pb *ParagraphBorders) SetBetween(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("between", style, size, color)
}

// Bar returns the bar border.
func (pb *ParagraphBorders) Bar() *Border {
	return pb.getBorder("bar")
}

// SetBar sets the bar border.
func (pb *ParagraphBorders) SetBar(
	style BorderStyle,
	size int,
	color string,
) {
	pb.setBorder("bar", style, size, color)
}

func (pb *ParagraphBorders) getBorder(
	name string,
) *Border {
	elem := pb.GetElement(name, NamespaceWML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Border); ok {
		return b
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Border{CompositeElementBase: comp}
	}

	return nil
}

func (pb *ParagraphBorders) setBorder(
	name string,
	style BorderStyle,
	size int,
	color string,
) {
	border := NewBorder(name, style, size, color)
	// Remove existing
	if existing := pb.GetElement(name, NamespaceWML); existing != nil {
		pb.RemoveChild(existing)
	}
	pb.AppendChild(border)
}

// Clone creates a deep copy of this ParagraphBorders element.
func (pb *ParagraphBorders) Clone() openxml.Element {
	return &ParagraphBorders{
		CompositeElementBase: pb.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Border represents a single border element.
type Border struct {
	*openxml.CompositeElementBase
}

// NewBorder creates a new Border element.
func NewBorder(
	name string,
	style BorderStyle,
	size int,
	color string,
) *Border {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	b := &Border{CompositeElementBase: elem}
	b.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(style),
		),
	)
	if size > 0 {
		b.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"sz",
				PrefixW,
				strconv.Itoa(size),
			),
		)
	}
	if color != "" {
		b.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"color",
				PrefixW,
				color,
			),
		)
	}

	return b
}

// Value returns the border style.
func (b *Border) Value() BorderStyle {
	attr, found := b.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return BorderNone
	}

	return BorderStyle(attr.Value())
}

// Size returns the border size in eighths of a point.
func (b *Border) Size() int {
	attr, found := b.GetAttribute(
		"sz",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Color returns the border color.
func (b *Border) Color() string {
	attr, found := b.GetAttribute(
		"color",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Space returns the spacing between border and content.
func (b *Border) Space() int {
	attr, found := b.GetAttribute(
		"space",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this Border element.
func (b *Border) Clone() openxml.Element {
	return &Border{
		CompositeElementBase: b.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Shading represents shading settings (w:shd).
type Shading struct {
	*openxml.CompositeElementBase
}

// NewShading creates a new Shading element.
func NewShading() *Shading {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"shd",
		PrefixW,
	)

	return &Shading{CompositeElementBase: elem}
}

// Fill returns the background fill color.
func (s *Shading) Fill() string {
	attr, found := s.GetAttribute(
		"fill",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetFill sets the background fill color.
func (s *Shading) SetFill(color string) {
	s.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"fill",
			PrefixW,
			color,
		),
	)
}

// Color returns the pattern color.
func (s *Shading) Color() string {
	attr, found := s.GetAttribute(
		"color",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetColor sets the pattern color.
func (s *Shading) SetColor(color string) {
	s.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"color",
			PrefixW,
			color,
		),
	)
}

// Val returns the shading pattern.
func (s *Shading) Val() ShadingPattern {
	attr, found := s.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ShadingClear
	}

	return ShadingPattern(attr.Value())
}

// SetVal sets the shading pattern.
func (s *Shading) SetVal(pattern ShadingPattern) {
	s.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(pattern),
		),
	)
}

// Clone creates a deep copy of this Shading element.
func (s *Shading) Clone() openxml.Element {
	return &Shading{
		CompositeElementBase: s.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
