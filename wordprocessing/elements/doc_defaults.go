package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// DocDefaults represents document-wide default formatting (w:docDefaults).
type DocDefaults struct {
	*openxml.CompositeElementBase
}

// NewDocDefaults creates a new DocDefaults element.
func NewDocDefaults() *DocDefaults {
	elem := openxml.NewCompositeElement(NamespaceWML, "docDefaults", PrefixW)
	return &DocDefaults{CompositeElementBase: elem}
}

// RunPropertiesDefault returns the default run properties element.
func (dd *DocDefaults) RunPropertiesDefault() *RunPropertiesDefault {
	elem := dd.GetElement("rPrDefault", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rpd, ok := elem.(*RunPropertiesDefault); ok {
		return rpd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RunPropertiesDefault{CompositeElementBase: comp}
	}
	return nil
}

// GetOrCreateRunPropertiesDefault returns the default run properties, creating if needed.
func (dd *DocDefaults) GetOrCreateRunPropertiesDefault() *RunPropertiesDefault {
	rpd := dd.RunPropertiesDefault()
	if rpd != nil {
		return rpd
	}
	rpd = NewRunPropertiesDefault()
	// Run properties default should be first
	if first := dd.FirstChild(); first != nil {
		dd.InsertBefore(rpd, first)
	} else {
		dd.AppendChild(rpd)
	}
	return rpd
}

// ParagraphPropertiesDefault returns the default paragraph properties element.
func (dd *DocDefaults) ParagraphPropertiesDefault() *ParagraphPropertiesDefault {
	elem := dd.GetElement("pPrDefault", NamespaceWML)
	if elem == nil {
		return nil
	}
	if ppd, ok := elem.(*ParagraphPropertiesDefault); ok {
		return ppd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ParagraphPropertiesDefault{CompositeElementBase: comp}
	}
	return nil
}

// GetOrCreateParagraphPropertiesDefault returns the default paragraph properties, creating if needed.
func (dd *DocDefaults) GetOrCreateParagraphPropertiesDefault() *ParagraphPropertiesDefault {
	ppd := dd.ParagraphPropertiesDefault()
	if ppd != nil {
		return ppd
	}
	ppd = NewParagraphPropertiesDefault()
	dd.AppendChild(ppd)
	return ppd
}

// Clone creates a deep copy of this DocDefaults element.
func (dd *DocDefaults) Clone() openxml.Element {
	return &DocDefaults{
		CompositeElementBase: dd.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DocDefaults element.
func (dd *DocDefaults) CloneNode(deep bool) openxml.Element {
	return &DocDefaults{
		CompositeElementBase: dd.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// RunPropertiesDefault represents default run properties (w:rPrDefault).
type RunPropertiesDefault struct {
	*openxml.CompositeElementBase
}

// NewRunPropertiesDefault creates a new RunPropertiesDefault element.
func NewRunPropertiesDefault() *RunPropertiesDefault {
	elem := openxml.NewCompositeElement(NamespaceWML, "rPrDefault", PrefixW)
	return &RunPropertiesDefault{CompositeElementBase: elem}
}

// RunProperties returns the default run properties.
func (rpd *RunPropertiesDefault) RunProperties() *DefaultRunProperties {
	elem := rpd.GetElement("rPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if rp, ok := elem.(*DefaultRunProperties); ok {
		return rp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DefaultRunProperties{CompositeElementBase: comp}
	}
	return nil
}

// GetOrCreateRunProperties returns the default run properties, creating if needed.
func (rpd *RunPropertiesDefault) GetOrCreateRunProperties() *DefaultRunProperties {
	rp := rpd.RunProperties()
	if rp != nil {
		return rp
	}
	rp = NewDefaultRunProperties()
	rpd.AppendChild(rp)
	return rp
}

// Clone creates a deep copy of this RunPropertiesDefault element.
func (rpd *RunPropertiesDefault) Clone() openxml.Element {
	return &RunPropertiesDefault{
		CompositeElementBase: rpd.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ParagraphPropertiesDefault represents default paragraph properties (w:pPrDefault).
type ParagraphPropertiesDefault struct {
	*openxml.CompositeElementBase
}

// NewParagraphPropertiesDefault creates a new ParagraphPropertiesDefault element.
func NewParagraphPropertiesDefault() *ParagraphPropertiesDefault {
	elem := openxml.NewCompositeElement(NamespaceWML, "pPrDefault", PrefixW)
	return &ParagraphPropertiesDefault{CompositeElementBase: elem}
}

// ParagraphProperties returns the default paragraph properties.
func (ppd *ParagraphPropertiesDefault) ParagraphProperties() *DefaultParagraphProperties {
	elem := ppd.GetElement("pPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pp, ok := elem.(*DefaultParagraphProperties); ok {
		return pp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DefaultParagraphProperties{CompositeElementBase: comp}
	}
	return nil
}

// GetOrCreateParagraphProperties returns the default paragraph properties, creating if needed.
func (ppd *ParagraphPropertiesDefault) GetOrCreateParagraphProperties() *DefaultParagraphProperties {
	pp := ppd.ParagraphProperties()
	if pp != nil {
		return pp
	}
	pp = NewDefaultParagraphProperties()
	ppd.AppendChild(pp)
	return pp
}

// Clone creates a deep copy of this ParagraphPropertiesDefault element.
func (ppd *ParagraphPropertiesDefault) Clone() openxml.Element {
	return &ParagraphPropertiesDefault{
		CompositeElementBase: ppd.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// DefaultRunProperties represents default run formatting (w:rPr within rPrDefault).
type DefaultRunProperties struct {
	*openxml.CompositeElementBase
}

// NewDefaultRunProperties creates a new DefaultRunProperties element.
func NewDefaultRunProperties() *DefaultRunProperties {
	elem := openxml.NewCompositeElement(NamespaceWML, "rPr", PrefixW)
	return &DefaultRunProperties{CompositeElementBase: elem}
}

// SetFontSize sets the default font size in half-points.
func (rp *DefaultRunProperties) SetFontSize(halfPoints int) {
	if halfPoints <= 0 {
		rp.removeElement("sz")
		return
	}
	elem := rp.getOrCreateElement("sz")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, strconv.Itoa(halfPoints)))
}

// SetFontSizeComplexScript sets the default complex script font size in half-points.
func (rp *DefaultRunProperties) SetFontSizeComplexScript(halfPoints int) {
	if halfPoints <= 0 {
		rp.removeElement("szCs")
		return
	}
	elem := rp.getOrCreateElement("szCs")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, strconv.Itoa(halfPoints)))
}

// SetFont sets the default font name for all script types.
func (rp *DefaultRunProperties) SetFont(fontName string) {
	rf := rp.getOrCreateElement("rFonts")
	rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "asciiTheme", PrefixW, "minorHAnsi"))
	rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "hAnsiTheme", PrefixW, "minorHAnsi"))
	rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "eastAsiaTheme", PrefixW, "minorHAnsi"))
	rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "cstheme", PrefixW, "minorBidi"))
	if fontName != "" {
		rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "ascii", PrefixW, fontName))
		rf.SetAttribute(openxml.NewAttribute(NamespaceWML, "hAnsi", PrefixW, fontName))
	}
}

// SetLanguage sets the default language.
func (rp *DefaultRunProperties) SetLanguage(lang string) {
	elem := rp.getOrCreateElement("lang")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "val", PrefixW, lang))
}

func (rp *DefaultRunProperties) getOrCreateElement(name string) openxml.Element {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(NamespaceWML, name, PrefixW)
	rp.AppendChild(newElem)
	return newElem
}

func (rp *DefaultRunProperties) removeElement(name string) {
	elem := rp.GetElement(name, NamespaceWML)
	if elem != nil {
		rp.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this DefaultRunProperties element.
func (rp *DefaultRunProperties) Clone() openxml.Element {
	return &DefaultRunProperties{
		CompositeElementBase: rp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// DefaultParagraphProperties represents default paragraph formatting (w:pPr within pPrDefault).
type DefaultParagraphProperties struct {
	*openxml.CompositeElementBase
}

// NewDefaultParagraphProperties creates a new DefaultParagraphProperties element.
func NewDefaultParagraphProperties() *DefaultParagraphProperties {
	elem := openxml.NewCompositeElement(NamespaceWML, "pPr", PrefixW)
	return &DefaultParagraphProperties{CompositeElementBase: elem}
}

// SetSpacingBefore sets the default space before paragraphs in twips.
func (pp *DefaultParagraphProperties) SetSpacingBefore(twips int) {
	elem := pp.getOrCreateElement("spacing")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "before", PrefixW, strconv.Itoa(twips)))
}

// SetSpacingAfter sets the default space after paragraphs in twips.
func (pp *DefaultParagraphProperties) SetSpacingAfter(twips int) {
	elem := pp.getOrCreateElement("spacing")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "after", PrefixW, strconv.Itoa(twips)))
}

// SetLineSpacing sets the default line spacing value.
func (pp *DefaultParagraphProperties) SetLineSpacing(value int, rule LineSpacingRule) {
	elem := pp.getOrCreateElement("spacing")
	elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "line", PrefixW, strconv.Itoa(value)))
	if rule != "" {
		elem.SetAttribute(openxml.NewAttribute(NamespaceWML, "lineRule", PrefixW, string(rule)))
	}
}

func (pp *DefaultParagraphProperties) getOrCreateElement(name string) openxml.Element {
	elem := pp.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(NamespaceWML, name, PrefixW)
	pp.AppendChild(newElem)
	return newElem
}

// Clone creates a deep copy of this DefaultParagraphProperties element.
func (pp *DefaultParagraphProperties) Clone() openxml.Element {
	return &DefaultParagraphProperties{
		CompositeElementBase: pp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
