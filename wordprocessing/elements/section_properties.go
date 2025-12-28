//nolint:revive // file-length-limit: comprehensive section properties with many related types
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

const (
	// Common attribute names
	attrNameTitlePg = "titlePg"
	attrNameType    = "type"
	attrNameID      = "id"
	attrNameVal     = "val"
	attrNameStart   = "start"
	attrNameName    = "name"

	// Common attribute values
	attrValueTrue  = "true"
	attrValueFalse = "false"
	attrValueZero  = "0"

	// Page dimensions in twips (1/20th of a point)
	defaultPageWidth    = 12240 // 8.5 inches = 12240 twips
	defaultPageHeight   = 15840 // 11 inches = 15840 twips
	defaultMargin       = 1440  // 1 inch = 1440 twips
	defaultHeaderFooter = 720   // 0.5 inches = 720 twips
)

// SectionProperties represents section properties (w:sectPr).
type SectionProperties struct {
	*openxml.CompositeElementBase
}

// NewSectionProperties creates a new SectionProperties element.
func NewSectionProperties() *SectionProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"sectPr",
		PrefixW,
	)

	return &SectionProperties{
		CompositeElementBase: elem,
	}
}

// PageSize returns the page size settings.
func (sp *SectionProperties) PageSize() *PageSize {
	elem := sp.GetElement("pgSz", NamespaceWML)
	if elem == nil {
		return nil
	}
	if ps, ok := elem.(*PageSize); ok {
		return ps
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageSize{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageSize returns the page size settings, creating if needed.
func (sp *SectionProperties) GetOrCreatePageSize() *PageSize {
	ps := sp.PageSize()
	if ps != nil {
		return ps
	}
	ps = NewPageSize()
	sp.AppendChild(ps)

	return ps
}

// PageMargins returns the page margin settings.
func (sp *SectionProperties) PageMargins() *PageMargins {
	elem := sp.GetElement("pgMar", NamespaceWML)
	if elem == nil {
		return nil
	}
	if pm, ok := elem.(*PageMargins); ok {
		return pm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageMargins{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageMargins returns the page margins, creating if needed.
func (sp *SectionProperties) GetOrCreatePageMargins() *PageMargins {
	pm := sp.PageMargins()
	if pm != nil {
		return pm
	}
	pm = NewPageMargins()
	sp.AppendChild(pm)

	return pm
}

// Columns returns the column settings.
func (sp *SectionProperties) Columns() *Columns {
	elem := sp.GetElement("cols", NamespaceWML)
	if elem == nil {
		return nil
	}
	if cols, ok := elem.(*Columns); ok {
		return cols
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Columns{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateColumns returns the column settings, creating if needed.
func (sp *SectionProperties) GetOrCreateColumns() *Columns {
	cols := sp.Columns()
	if cols != nil {
		return cols
	}
	cols = NewColumns()
	sp.AppendChild(cols)

	return cols
}

// HeaderReferences returns all header references.
func (sp *SectionProperties) HeaderReferences() []*HeaderReference {
	var refs []*HeaderReference
	for child := range sp.Children() {
		if child.LocalName() == "headerReference" &&
			child.NamespaceURI() == NamespaceWML {
			if hr, ok := child.(*HeaderReference); ok {
				refs = append(refs, hr)
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				refs = append(refs, &HeaderReference{CompositeElementBase: comp})
			}
		}
	}

	return refs
}

// FooterReferences returns all footer references.
func (sp *SectionProperties) FooterReferences() []*FooterReference {
	var refs []*FooterReference
	for child := range sp.Children() {
		if child.LocalName() == "footerReference" &&
			child.NamespaceURI() == NamespaceWML {
			if fr, ok := child.(*FooterReference); ok {
				refs = append(refs, fr)
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				refs = append(refs, &FooterReference{CompositeElementBase: comp})
			}
		}
	}

	return refs
}

// AddHeaderReference adds a header reference.
func (sp *SectionProperties) AddHeaderReference(
	relId string,
	hfType HeaderFooterValues,
) *HeaderReference {
	hr := NewHeaderReference(relId, hfType)
	sp.AppendChild(hr)

	return hr
}

// AddFooterReference adds a footer reference.
func (sp *SectionProperties) AddFooterReference(
	relId string,
	hfType HeaderFooterValues,
) *FooterReference {
	fr := NewFooterReference(relId, hfType)
	sp.AppendChild(fr)

	return fr
}

// GetHeaderReference returns the header reference of the specified type, or nil if not found.
func (sp *SectionProperties) GetHeaderReference(
	hfType HeaderFooterValues,
) *HeaderReference {
	for _, hr := range sp.HeaderReferences() {
		if hr.Type() == hfType {
			return hr
		}
	}

	return nil
}

// GetFooterReference returns the footer reference of the specified type, or nil if not found.
func (sp *SectionProperties) GetFooterReference(
	hfType HeaderFooterValues,
) *FooterReference {
	for _, fr := range sp.FooterReferences() {
		if fr.Type() == hfType {
			return fr
		}
	}

	return nil
}

// RemoveHeaderReference removes the header reference of the specified type.
// Returns true if a reference was removed.
func (sp *SectionProperties) RemoveHeaderReference(
	hfType HeaderFooterValues,
) bool {
	for _, hr := range sp.HeaderReferences() {
		if hr.Type() == hfType {
			sp.RemoveChild(hr)

			return true
		}
	}

	return false
}

// RemoveFooterReference removes the footer reference of the specified type.
// Returns true if a reference was removed.
func (sp *SectionProperties) RemoveFooterReference(
	hfType HeaderFooterValues,
) bool {
	for _, fr := range sp.FooterReferences() {
		if fr.Type() == hfType {
			sp.RemoveChild(fr)

			return true
		}
	}

	return false
}

// SetHeaderReference sets or replaces the header reference of the specified type.
func (sp *SectionProperties) SetHeaderReference(
	relId string,
	hfType HeaderFooterValues,
) *HeaderReference {
	sp.RemoveHeaderReference(hfType)

	return sp.AddHeaderReference(relId, hfType)
}

// SetFooterReference sets or replaces the footer reference of the specified type.
func (sp *SectionProperties) SetFooterReference(
	relId string,
	hfType HeaderFooterValues,
) *FooterReference {
	sp.RemoveFooterReference(hfType)

	return sp.AddFooterReference(relId, hfType)
}

// TitlePage returns whether the section has a different first page header/footer.
func (sp *SectionProperties) TitlePage() bool {
	elem := sp.GetElement(
		attrNameTitlePg,
		NamespaceWML,
	)
	if elem == nil {
		return false
	}
	attr, found := elem.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if found {
		val := attr.Value()

		return val != attrValueFalse &&
			val != attrValueZero
	}

	return true
}

// SetTitlePage sets whether the section has a different first page header/footer.
func (sp *SectionProperties) SetTitlePage(
	b bool,
) {
	if b {
		if sp.GetElement(
			attrNameTitlePg,
			NamespaceWML,
		) == nil {
			elem := NewTitlePageElement()
			sp.AppendChild(elem)
		}
	} else {
		if elem := sp.GetElement(attrNameTitlePg, NamespaceWML); elem != nil {
			sp.RemoveChild(elem)
		}
	}
}

// GetOrCreateTitlePage returns the title page element, creating if needed.
func (sp *SectionProperties) GetOrCreateTitlePage() *TitlePageElement {
	elem := sp.GetElement(
		attrNameTitlePg,
		NamespaceWML,
	)
	if elem == nil {
		tp := NewTitlePageElement()
		sp.AppendChild(tp)

		return tp
	}
	if tp, ok := elem.(*TitlePageElement); ok {
		return tp
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TitlePageElement{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// SectionTypeElement returns the section type element, or nil if not present.
func (sp *SectionProperties) SectionTypeElement() *SectionType {
	elem := sp.GetElement(
		attrNameType,
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if st, ok := elem.(*SectionType); ok {
		return st
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &SectionType{LeafElementBase: leaf}
	}

	return nil
}

// GetSectionType returns the section type value.
func (sp *SectionProperties) GetSectionType() SectionTypeValue {
	st := sp.SectionTypeElement()
	if st == nil {
		return SectionTypeNextPage
	}

	return st.Value()
}

// SetSectionType sets the section type.
func (sp *SectionProperties) SetSectionType(
	val SectionTypeValue,
) {
	if elem := sp.GetElement(attrNameType, NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	st := NewSectionType(val)
	sp.AppendChild(st)
}

// PageNumberTypeElement returns the page number type element, or nil if not present.
func (sp *SectionProperties) PageNumberTypeElement() *PageNumberType {
	elem := sp.GetElement(
		"pgNumType",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if pn, ok := elem.(*PageNumberType); ok {
		return pn
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &PageNumberType{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreatePageNumberType returns the page number type element, creating if needed.
func (sp *SectionProperties) GetOrCreatePageNumberType() *PageNumberType {
	pn := sp.PageNumberTypeElement()
	if pn != nil {
		return pn
	}
	pn = NewPageNumberType()
	sp.AppendChild(pn)

	return pn
}

// FormProtectionElement returns the form protection element, or nil if not present.
func (sp *SectionProperties) FormProtectionElement() *FormProtection {
	elem := sp.GetElement(
		"formProt",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if fp, ok := elem.(*FormProtection); ok {
		return fp
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &FormProtection{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// SetFormProtection sets the form protection.
func (sp *SectionProperties) SetFormProtection(
	enabled bool,
) {
	if elem := sp.GetElement("formProt", NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	if enabled {
		fp := NewFormProtection(true)
		sp.AppendChild(fp)
	}
}

// VerticalTextAlignmentElement returns the vertical text alignment element, or nil if not present.
func (sp *SectionProperties) VerticalTextAlignmentElement() *VerticalTextAlignment {
	elem := sp.GetElement("vAlign", NamespaceWML)
	if elem == nil {
		return nil
	}
	if vta, ok := elem.(*VerticalTextAlignment); ok {
		return vta
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &VerticalTextAlignment{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetVerticalTextAlignment returns the vertical text alignment value.
func (sp *SectionProperties) GetVerticalTextAlignment() VerticalTextAlignmentValue {
	vta := sp.VerticalTextAlignmentElement()
	if vta == nil {
		return VerticalTextAlignTop
	}

	return vta.Value()
}

// SetVerticalTextAlignment sets the vertical text alignment.
func (sp *SectionProperties) SetVerticalTextAlignment(
	val VerticalTextAlignmentValue,
) {
	if elem := sp.GetElement("vAlign", NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	if val != VerticalTextAlignTop {
		vta := NewVerticalTextAlignment(val)
		sp.AppendChild(vta)
	}
}

// NoEndnoteElement returns the no endnote element, or nil if not present.
func (sp *SectionProperties) NoEndnoteElement() *NoEndnote {
	elem := sp.GetElement(
		"noEndnote",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ne, ok := elem.(*NoEndnote); ok {
		return ne
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &NoEndnote{LeafElementBase: leaf}
	}

	return nil
}

// SetNoEndnote sets whether endnotes are suppressed in this section.
func (sp *SectionProperties) SetNoEndnote(
	suppress bool,
) {
	if elem := sp.GetElement("noEndnote", NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	if suppress {
		ne := NewNoEndnote(true)
		sp.AppendChild(ne)
	}
}

// PaperSourceElement returns the paper source element, or nil if not present.
func (sp *SectionProperties) PaperSourceElement() *PaperSource {
	elem := sp.GetElement(
		"paperSrc",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ps, ok := elem.(*PaperSource); ok {
		return ps
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &PaperSource{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreatePaperSource returns the paper source element, creating if needed.
func (sp *SectionProperties) GetOrCreatePaperSource() *PaperSource {
	ps := sp.PaperSourceElement()
	if ps != nil {
		return ps
	}
	ps = NewPaperSource()
	sp.AppendChild(ps)

	return ps
}

// LineNumberTypeElement returns the line number type element, or nil if not present.
func (sp *SectionProperties) LineNumberTypeElement() *LineNumberType {
	elem := sp.GetElement(
		"lnNumType",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ln, ok := elem.(*LineNumberType); ok {
		return ln
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &LineNumberType{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateLineNumberType returns the line number type element, creating if needed.
func (sp *SectionProperties) GetOrCreateLineNumberType() *LineNumberType {
	ln := sp.LineNumberTypeElement()
	if ln != nil {
		return ln
	}
	ln = NewLineNumberType()
	sp.AppendChild(ln)

	return ln
}

// TextDirectionElement returns the text direction element, or nil if not present.
func (sp *SectionProperties) TextDirectionElement() *TextDirection {
	elem := sp.GetElement(
		"textDirection",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if td, ok := elem.(*TextDirection); ok {
		return td
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &TextDirection{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetTextDirection returns the text direction value.
func (sp *SectionProperties) GetTextDirection() TextDirectionValue {
	td := sp.TextDirectionElement()
	if td == nil {
		return TextDirectionLrTb
	}

	return td.Value()
}

// SetTextDirection sets the text direction.
func (sp *SectionProperties) SetTextDirection(
	val TextDirectionValue,
) {
	if elem := sp.GetElement("textDirection", NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	if val != TextDirectionLrTb {
		td := NewTextDirection(val)
		sp.AppendChild(td)
	}
}

// RTLGutterElement returns the RTL gutter element, or nil if not present.
func (sp *SectionProperties) RTLGutterElement() *RTLGutter {
	elem := sp.GetElement(
		"rtlGutter",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if rg, ok := elem.(*RTLGutter); ok {
		return rg
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &RTLGutter{LeafElementBase: leaf}
	}

	return nil
}

// SetRTLGutter sets whether RTL gutter is enabled.
func (sp *SectionProperties) SetRTLGutter(
	enabled bool,
) {
	if elem := sp.GetElement("rtlGutter", NamespaceWML); elem != nil {
		sp.RemoveChild(elem)
	}
	if enabled {
		rg := NewRTLGutter(true)
		sp.AppendChild(rg)
	}
}

// DocGridElement returns the document grid element, or nil if not present.
func (sp *SectionProperties) DocGridElement() *DocGrid {
	elem := sp.GetElement("docGrid", NamespaceWML)
	if elem == nil {
		return nil
	}
	if dg, ok := elem.(*DocGrid); ok {
		return dg
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &DocGrid{LeafElementBase: leaf}
	}

	return nil
}

// GetOrCreateDocGrid returns the document grid element, creating if needed.
func (sp *SectionProperties) GetOrCreateDocGrid() *DocGrid {
	dg := sp.DocGridElement()
	if dg != nil {
		return dg
	}
	dg = NewDocGrid()
	sp.AppendChild(dg)

	return dg
}

// Clone creates a deep copy of this SectionProperties element.
func (sp *SectionProperties) Clone() openxml.Element {
	return &SectionProperties{
		CompositeElementBase: sp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SectionProperties element.
func (sp *SectionProperties) CloneNode(
	deep bool,
) openxml.Element {
	return &SectionProperties{
		CompositeElementBase: sp.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// PageSize represents page size settings (w:pgSz).
type PageSize struct {
	*openxml.CompositeElementBase
}

// NewPageSize creates a new PageSize element with default Letter size.
func NewPageSize() *PageSize {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pgSz",
		PrefixW,
	)
	ps := &PageSize{CompositeElementBase: elem}
	// Default to Letter size (8.5 x 11 inches)
	ps.SetWidth(defaultPageWidth)
	ps.SetHeight(defaultPageHeight)

	return ps
}

// Width returns the page width in twips.
func (ps *PageSize) Width() int {
	attr, found := ps.GetAttribute(
		"w",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWidth sets the page width in twips.
func (ps *PageSize) SetWidth(twips int) {
	ps.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"w",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Height returns the page height in twips.
func (ps *PageSize) Height() int {
	attr, found := ps.GetAttribute(
		"h",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetHeight sets the page height in twips.
func (ps *PageSize) SetHeight(twips int) {
	ps.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"h",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Orient returns the page orientation.
func (ps *PageSize) Orient() PageOrientation {
	attr, found := ps.GetAttribute(
		"orient",
		NamespaceWML,
	)
	if !found {
		return PageOrientationPortrait
	}

	return PageOrientation(attr.Value())
}

// SetOrient sets the page orientation.
func (ps *PageSize) SetOrient(
	orient PageOrientation,
) {
	if orient == PageOrientationPortrait {
		ps.RemoveAttribute("orient", NamespaceWML)
	} else {
		ps.SetAttribute(openxml.NewAttribute(NamespaceWML, "orient", PrefixW, string(orient)))
	}
}

// Clone creates a deep copy of this PageSize element.
func (ps *PageSize) Clone() openxml.Element {
	return &PageSize{
		CompositeElementBase: ps.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageSize element.
func (ps *PageSize) CloneNode(
	deep bool,
) openxml.Element {
	return &PageSize{
		CompositeElementBase: ps.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// PageMargins represents page margin settings (w:pgMar).
type PageMargins struct {
	*openxml.CompositeElementBase
}

// NewPageMargins creates a new PageMargins element with default margins.
func NewPageMargins() *PageMargins {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"pgMar",
		PrefixW,
	)
	pm := &PageMargins{CompositeElementBase: elem}
	// Default 1 inch margins
	pm.SetTop(defaultMargin)
	pm.SetBottom(defaultMargin)
	pm.SetLeft(defaultMargin)
	pm.SetRight(defaultMargin)
	pm.SetHeader(defaultHeaderFooter)
	pm.SetFooter(defaultHeaderFooter)

	return pm
}

// Top returns the top margin in twips.
func (pm *PageMargins) Top() int {
	return pm.getIntAttribute("top")
}

// SetTop sets the top margin in twips.
func (pm *PageMargins) SetTop(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"top",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Bottom returns the bottom margin in twips.
func (pm *PageMargins) Bottom() int {
	return pm.getIntAttribute("bottom")
}

// SetBottom sets the bottom margin in twips.
func (pm *PageMargins) SetBottom(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"bottom",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Left returns the left margin in twips.
func (pm *PageMargins) Left() int {
	return pm.getIntAttribute("left")
}

// SetLeft sets the left margin in twips.
func (pm *PageMargins) SetLeft(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"left",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Right returns the right margin in twips.
func (pm *PageMargins) Right() int {
	return pm.getIntAttribute("right")
}

// SetRight sets the right margin in twips.
func (pm *PageMargins) SetRight(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"right",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Header returns the header margin in twips.
func (pm *PageMargins) Header() int {
	return pm.getIntAttribute("header")
}

// SetHeader sets the header margin in twips.
func (pm *PageMargins) SetHeader(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"header",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Footer returns the footer margin in twips.
func (pm *PageMargins) Footer() int {
	return pm.getIntAttribute("footer")
}

// SetFooter sets the footer margin in twips.
func (pm *PageMargins) SetFooter(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"footer",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Gutter returns the gutter margin in twips.
func (pm *PageMargins) Gutter() int {
	return pm.getIntAttribute("gutter")
}

// SetGutter sets the gutter margin in twips.
func (pm *PageMargins) SetGutter(twips int) {
	pm.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"gutter",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

func (pm *PageMargins) getIntAttribute(
	name string,
) int {
	attr, found := pm.GetAttribute(
		name,
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Clone creates a deep copy of this PageMargins element.
func (pm *PageMargins) Clone() openxml.Element {
	return &PageMargins{
		CompositeElementBase: pm.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageMargins element.
func (pm *PageMargins) CloneNode(
	deep bool,
) openxml.Element {
	return &PageMargins{
		CompositeElementBase: pm.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// Columns represents column settings (w:cols).
type Columns struct {
	*openxml.CompositeElementBase
}

// NewColumns creates a new Columns element.
func NewColumns() *Columns {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"cols",
		PrefixW,
	)

	return &Columns{CompositeElementBase: elem}
}

// Num returns the number of columns.
func (c *Columns) Num() int {
	attr, found := c.GetAttribute(
		"num",
		NamespaceWML,
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetNum sets the number of columns.
func (c *Columns) SetNum(num int) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"num",
			PrefixW,
			strconv.Itoa(num),
		),
	)
}

// Space returns the space between columns in twips.
func (c *Columns) Space() int {
	attr, found := c.GetAttribute(
		"space",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetSpace sets the space between columns in twips.
func (c *Columns) SetSpace(twips int) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"space",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// EqualWidth returns whether columns have equal width.
func (c *Columns) EqualWidth() bool {
	attr, found := c.GetAttribute(
		"equalWidth",
		NamespaceWML,
	)
	if !found {
		return true // default
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEqualWidth sets whether columns have equal width.
func (c *Columns) SetEqualWidth(b bool) {
	if b {
		c.RemoveAttribute(
			"equalWidth",
			NamespaceWML,
		)
	} else {
		c.SetAttribute(openxml.NewAttribute(NamespaceWML, "equalWidth", PrefixW, attrValueFalse))
	}
}

// Clone creates a deep copy of this Columns element.
func (c *Columns) Clone() openxml.Element {
	return &Columns{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Columns element.
func (c *Columns) CloneNode(
	deep bool,
) openxml.Element {
	return &Columns{
		CompositeElementBase: c.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// HeaderReference represents a header reference (w:headerReference).
type HeaderReference struct {
	*openxml.CompositeElementBase
}

// NewHeaderReference creates a new HeaderReference element.
func NewHeaderReference(
	relId string,
	hfType HeaderFooterValues,
) *HeaderReference {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"headerReference",
		PrefixW,
	)
	hr := &HeaderReference{
		CompositeElementBase: elem,
	}
	hr.SetAttribute(
		openxml.NewAttribute(
			openxml.NamespaceRelationships,
			attrNameID,
			"r",
			relId,
		),
	)
	hr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameType,
			PrefixW,
			string(hfType),
		),
	)

	return hr
}

// RelationshipId returns the relationship ID.
func (hr *HeaderReference) RelationshipId() string {
	attr, found := hr.GetAttribute(
		attrNameID,
		openxml.NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Type returns the header type.
func (hr *HeaderReference) Type() HeaderFooterValues {
	attr, found := hr.GetAttribute(
		attrNameType,
		NamespaceWML,
	)
	if !found {
		return HeaderFooterDefault
	}

	return HeaderFooterValues(attr.Value())
}

// Clone creates a deep copy of this HeaderReference element.
func (hr *HeaderReference) Clone() openxml.Element {
	return &HeaderReference{
		CompositeElementBase: hr.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this HeaderReference element.
func (hr *HeaderReference) CloneNode(
	deep bool,
) openxml.Element {
	return &HeaderReference{
		CompositeElementBase: hr.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// FooterReference represents a footer reference (w:footerReference).
type FooterReference struct {
	*openxml.CompositeElementBase
}

// NewFooterReference creates a new FooterReference element.
func NewFooterReference(
	relId string,
	hfType HeaderFooterValues,
) *FooterReference {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"footerReference",
		PrefixW,
	)
	fr := &FooterReference{
		CompositeElementBase: elem,
	}
	fr.SetAttribute(
		openxml.NewAttribute(
			openxml.NamespaceRelationships,
			attrNameID,
			"r",
			relId,
		),
	)
	fr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameType,
			PrefixW,
			string(hfType),
		),
	)

	return fr
}

// RelationshipId returns the relationship ID.
func (fr *FooterReference) RelationshipId() string {
	attr, found := fr.GetAttribute(
		attrNameID,
		openxml.NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Type returns the footer type.
func (fr *FooterReference) Type() HeaderFooterValues {
	attr, found := fr.GetAttribute(
		attrNameType,
		NamespaceWML,
	)
	if !found {
		return HeaderFooterDefault
	}

	return HeaderFooterValues(attr.Value())
}

// Clone creates a deep copy of this FooterReference element.
func (fr *FooterReference) Clone() openxml.Element {
	return &FooterReference{
		CompositeElementBase: fr.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FooterReference element.
func (fr *FooterReference) CloneNode(
	deep bool,
) openxml.Element {
	return &FooterReference{
		CompositeElementBase: fr.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// SectionTypeValue represents section break type values.
type SectionTypeValue string

const (
	// SectionTypeNextPage starts the section on the next page.
	SectionTypeNextPage SectionTypeValue = "nextPage"
	// SectionTypeContinuous continues the section without a break.
	SectionTypeContinuous SectionTypeValue = "continuous"
	// SectionTypeEvenPage starts the section on the next even page.
	SectionTypeEvenPage SectionTypeValue = "evenPage"
	// SectionTypeOddPage starts the section on the next odd page.
	SectionTypeOddPage SectionTypeValue = "oddPage"
	// SectionTypeNextColumn starts the section in the next column.
	SectionTypeNextColumn SectionTypeValue = "nextColumn"
)

// SectionType represents section type element (w:type).
type SectionType struct {
	*openxml.LeafElementBase
}

// NewSectionType creates a new SectionType element with the specified value.
func NewSectionType(
	val SectionTypeValue,
) *SectionType {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"type",
		PrefixW,
	)
	st := &SectionType{LeafElementBase: elem}
	st.SetValue(val)

	return st
}

// Value returns the section type value.
func (st *SectionType) Value() SectionTypeValue {
	attr, found := st.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return SectionTypeNextPage
	}

	return SectionTypeValue(attr.Value())
}

// SetValue sets the section type value.
func (st *SectionType) SetValue(
	val SectionTypeValue,
) {
	st.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			string(val),
		),
	)
}

// Clone creates a deep copy of this SectionType element.
func (st *SectionType) Clone() openxml.Element {
	return &SectionType{
		LeafElementBase: st.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this SectionType element.
func (st *SectionType) CloneNode(
	deep bool,
) openxml.Element {
	return &SectionType{
		LeafElementBase: st.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// TitlePageElement represents the title page element (w:titlePg).
// This element indicates that the first page has different header/footer.
type TitlePageElement struct {
	*openxml.LeafElementBase
}

// NewTitlePageElement creates a new TitlePageElement element.
func NewTitlePageElement() *TitlePageElement {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"titlePg",
		PrefixW,
	)

	return &TitlePageElement{
		LeafElementBase: elem,
	}
}

// Clone creates a deep copy of this TitlePageElement element.
func (tp *TitlePageElement) Clone() openxml.Element {
	return &TitlePageElement{
		LeafElementBase: tp.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TitlePageElement element.
func (tp *TitlePageElement) CloneNode(
	deep bool,
) openxml.Element {
	return &TitlePageElement{
		LeafElementBase: tp.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// PageNumberFormatValue represents page number format values.
type PageNumberFormatValue string

const (
	// PageNumberFormatDecimal uses decimal numbers (1, 2, 3, ...).
	PageNumberFormatDecimal PageNumberFormatValue = "decimal"
	// PageNumberFormatUpperRoman uses uppercase Roman numerals (I, II, III, ...).
	PageNumberFormatUpperRoman PageNumberFormatValue = "upperRoman"
	// PageNumberFormatLowerRoman uses lowercase Roman numerals (i, ii, iii, ...).
	PageNumberFormatLowerRoman PageNumberFormatValue = "lowerRoman"
	// PageNumberFormatUpperLetter uses uppercase letters (A, B, C, ...).
	PageNumberFormatUpperLetter PageNumberFormatValue = "upperLetter"
	// PageNumberFormatLowerLetter uses lowercase letters (a, b, c, ...).
	PageNumberFormatLowerLetter PageNumberFormatValue = "lowerLetter"
	// PageNumberFormatCardinalText uses cardinal text (one, two, three, ...).
	PageNumberFormatCardinalText PageNumberFormatValue = "cardinalText"
	// PageNumberFormatOrdinalText uses ordinal text (first, second, third, ...).
	PageNumberFormatOrdinalText PageNumberFormatValue = "ordinalText"
	// PageNumberFormatNone indicates no page numbering.
	PageNumberFormatNone PageNumberFormatValue = "none"
)

// PageNumberType represents page numbering settings (w:pgNumType).
type PageNumberType struct {
	*openxml.LeafElementBase
}

// NewPageNumberType creates a new PageNumberType element.
func NewPageNumberType() *PageNumberType {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"pgNumType",
		PrefixW,
	)

	return &PageNumberType{LeafElementBase: elem}
}

// Format returns the page number format.
func (pn *PageNumberType) Format() PageNumberFormatValue {
	attr, found := pn.GetAttribute(
		"fmt",
		NamespaceWML,
	)
	if !found {
		return PageNumberFormatDecimal
	}

	return PageNumberFormatValue(attr.Value())
}

// SetFormat sets the page number format.
func (pn *PageNumberType) SetFormat(
	format PageNumberFormatValue,
) {
	pn.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"fmt",
			PrefixW,
			string(format),
		),
	)
}

// Start returns the starting page number.
func (pn *PageNumberType) Start() int {
	attr, found := pn.GetAttribute(
		attrNameStart,
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetStart sets the starting page number.
func (pn *PageNumberType) SetStart(start int) {
	pn.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameStart,
			PrefixW,
			strconv.Itoa(start),
		),
	)
}

// ChapStyle returns the chapter heading style level.
func (pn *PageNumberType) ChapStyle() int {
	attr, found := pn.GetAttribute(
		"chapStyle",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetChapStyle sets the chapter heading style level.
func (pn *PageNumberType) SetChapStyle(
	level int,
) {
	if level <= 0 {
		pn.RemoveAttribute(
			"chapStyle",
			NamespaceWML,
		)
	} else {
		pn.SetAttribute(openxml.NewAttribute(NamespaceWML, "chapStyle", PrefixW, strconv.Itoa(level)))
	}
}

// ChapSep returns the chapter separator character.
func (pn *PageNumberType) ChapSep() string {
	attr, found := pn.GetAttribute(
		"chapSep",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetChapSep sets the chapter separator character.
// Valid values: "hyphen", "period", "colon", "emDash", "enDash"
func (pn *PageNumberType) SetChapSep(sep string) {
	if sep == "" {
		pn.RemoveAttribute(
			"chapSep",
			NamespaceWML,
		)
	} else {
		pn.SetAttribute(openxml.NewAttribute(NamespaceWML, "chapSep", PrefixW, sep))
	}
}

// Clone creates a deep copy of this PageNumberType element.
func (pn *PageNumberType) Clone() openxml.Element {
	return &PageNumberType{
		LeafElementBase: pn.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PageNumberType element.
func (pn *PageNumberType) CloneNode(
	deep bool,
) openxml.Element {
	return &PageNumberType{
		LeafElementBase: pn.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// FormProtection represents form protection setting (w:formProt).
type FormProtection struct {
	*openxml.LeafElementBase
}

// NewFormProtection creates a new FormProtection element.
func NewFormProtection(
	enabled bool,
) *FormProtection {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"formProt",
		PrefixW,
	)
	fp := &FormProtection{LeafElementBase: elem}
	if !enabled {
		fp.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameVal,
				PrefixW,
				attrValueFalse,
			),
		)
	}

	return fp
}

// Enabled returns whether form protection is enabled.
func (fp *FormProtection) Enabled() bool {
	attr, found := fp.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return true // default is true when element is present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnabled sets whether form protection is enabled.
func (fp *FormProtection) SetEnabled(
	enabled bool,
) {
	if enabled {
		fp.RemoveAttribute(
			attrNameVal,
			NamespaceWML,
		)
	} else {
		fp.SetAttribute(openxml.NewAttribute(NamespaceWML, attrNameVal, PrefixW, attrValueFalse))
	}
}

// Clone creates a deep copy of this FormProtection element.
func (fp *FormProtection) Clone() openxml.Element {
	return &FormProtection{
		LeafElementBase: fp.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this FormProtection element.
func (fp *FormProtection) CloneNode(
	deep bool,
) openxml.Element {
	return &FormProtection{
		LeafElementBase: fp.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// VerticalTextAlignmentValue represents vertical text alignment values for pages.
type VerticalTextAlignmentValue string

const (
	// VerticalTextAlignTop aligns text to the top of the page.
	VerticalTextAlignTop VerticalTextAlignmentValue = "top"
	// VerticalTextAlignCenter centers text vertically on the page.
	VerticalTextAlignCenter VerticalTextAlignmentValue = "center"
	// VerticalTextAlignBoth justifies text to fill the page vertically.
	VerticalTextAlignBoth VerticalTextAlignmentValue = "both"
	// VerticalTextAlignBottom aligns text to the bottom of the page.
	VerticalTextAlignBottom VerticalTextAlignmentValue = "bottom"
)

// VerticalTextAlignment represents vertical alignment on page (w:vAlign).
type VerticalTextAlignment struct {
	*openxml.LeafElementBase
}

// NewVerticalTextAlignment creates a new VerticalTextAlignment element.
func NewVerticalTextAlignment(
	val VerticalTextAlignmentValue,
) *VerticalTextAlignment {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"vAlign",
		PrefixW,
	)
	vta := &VerticalTextAlignment{
		LeafElementBase: elem,
	}
	vta.SetValue(val)

	return vta
}

// Value returns the vertical text alignment value.
func (v *VerticalTextAlignment) Value() VerticalTextAlignmentValue {
	attr, found := v.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return VerticalTextAlignTop
	}

	return VerticalTextAlignmentValue(
		attr.Value(),
	)
}

// SetValue sets the vertical text alignment value.
func (v *VerticalTextAlignment) SetValue(
	val VerticalTextAlignmentValue,
) {
	v.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			string(val),
		),
	)
}

// Clone creates a deep copy of this VerticalTextAlignment element.
func (v *VerticalTextAlignment) Clone() openxml.Element {
	return &VerticalTextAlignment{
		LeafElementBase: v.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this VerticalTextAlignment element.
func (v *VerticalTextAlignment) CloneNode(
	deep bool,
) openxml.Element {
	return &VerticalTextAlignment{
		LeafElementBase: v.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// NoEndnote represents suppress endnotes in section (w:noEndnote).
type NoEndnote struct {
	*openxml.LeafElementBase
}

// NewNoEndnote creates a new NoEndnote element.
func NewNoEndnote(suppress bool) *NoEndnote {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"noEndnote",
		PrefixW,
	)
	ne := &NoEndnote{LeafElementBase: elem}
	if !suppress {
		ne.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameVal,
				PrefixW,
				attrValueFalse,
			),
		)
	}

	return ne
}

// Suppressed returns whether endnotes are suppressed.
func (ne *NoEndnote) Suppressed() bool {
	attr, found := ne.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return true // default is true when element is present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSuppressed sets whether endnotes are suppressed.
func (ne *NoEndnote) SetSuppressed(
	suppress bool,
) {
	if suppress {
		ne.RemoveAttribute(
			attrNameVal,
			NamespaceWML,
		)
	} else {
		ne.SetAttribute(openxml.NewAttribute(NamespaceWML, attrNameVal, PrefixW, attrValueFalse))
	}
}

// Clone creates a deep copy of this NoEndnote element.
func (ne *NoEndnote) Clone() openxml.Element {
	return &NoEndnote{
		LeafElementBase: ne.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this NoEndnote element.
func (ne *NoEndnote) CloneNode(
	deep bool,
) openxml.Element {
	return &NoEndnote{
		LeafElementBase: ne.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// PaperSource represents printer paper source (w:paperSrc).
type PaperSource struct {
	*openxml.LeafElementBase
}

// NewPaperSource creates a new PaperSource element.
func NewPaperSource() *PaperSource {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"paperSrc",
		PrefixW,
	)

	return &PaperSource{LeafElementBase: elem}
}

// First returns the first page paper source code.
func (ps *PaperSource) First() int {
	attr, found := ps.GetAttribute(
		"first",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetFirst sets the first page paper source code.
func (ps *PaperSource) SetFirst(code int) {
	if code <= 0 {
		ps.RemoveAttribute("first", NamespaceWML)
	} else {
		ps.SetAttribute(openxml.NewAttribute(NamespaceWML, "first", PrefixW, strconv.Itoa(code)))
	}
}

// Other returns the other pages paper source code.
func (ps *PaperSource) Other() int {
	attr, found := ps.GetAttribute(
		"other",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetOther sets the other pages paper source code.
func (ps *PaperSource) SetOther(code int) {
	if code <= 0 {
		ps.RemoveAttribute("other", NamespaceWML)
	} else {
		ps.SetAttribute(openxml.NewAttribute(NamespaceWML, "other", PrefixW, strconv.Itoa(code)))
	}
}

// Clone creates a deep copy of this PaperSource element.
func (ps *PaperSource) Clone() openxml.Element {
	return &PaperSource{
		LeafElementBase: ps.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this PaperSource element.
func (ps *PaperSource) CloneNode(
	deep bool,
) openxml.Element {
	return &PaperSource{
		LeafElementBase: ps.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// LineNumberRestartValue represents line number restart values.
type LineNumberRestartValue string

const (
	// LineNumberRestartNewPage restarts line numbering on each page.
	LineNumberRestartNewPage LineNumberRestartValue = "newPage"
	// LineNumberRestartNewSection restarts line numbering on each section.
	LineNumberRestartNewSection LineNumberRestartValue = "newSection"
	// LineNumberRestartContinuous continues line numbering from previous section.
	LineNumberRestartContinuous LineNumberRestartValue = "continuous"
)

// LineNumberType represents line numbering settings (w:lnNumType).
type LineNumberType struct {
	*openxml.LeafElementBase
}

// NewLineNumberType creates a new LineNumberType element.
func NewLineNumberType() *LineNumberType {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"lnNumType",
		PrefixW,
	)

	return &LineNumberType{LeafElementBase: elem}
}

// CountBy returns the line number increment.
func (ln *LineNumberType) CountBy() int {
	attr, found := ln.GetAttribute(
		"countBy",
		NamespaceWML,
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCountBy sets the line number increment.
func (ln *LineNumberType) SetCountBy(count int) {
	if count <= 1 {
		ln.RemoveAttribute(
			"countBy",
			NamespaceWML,
		)
	} else {
		ln.SetAttribute(openxml.NewAttribute(NamespaceWML, "countBy", PrefixW, strconv.Itoa(count)))
	}
}

// Start returns the starting line number.
func (ln *LineNumberType) Start() int {
	attr, found := ln.GetAttribute(
		attrNameStart,
		NamespaceWML,
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetStart sets the starting line number.
func (ln *LineNumberType) SetStart(start int) {
	if start <= 1 {
		ln.RemoveAttribute(
			attrNameStart,
			NamespaceWML,
		)
	} else {
		ln.SetAttribute(openxml.NewAttribute(NamespaceWML, attrNameStart, PrefixW, strconv.Itoa(start)))
	}
}

// Distance returns the distance from text to line numbers in twips.
func (ln *LineNumberType) Distance() int {
	attr, found := ln.GetAttribute(
		"distance",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetDistance sets the distance from text to line numbers in twips.
func (ln *LineNumberType) SetDistance(twips int) {
	if twips <= 0 {
		ln.RemoveAttribute(
			"distance",
			NamespaceWML,
		)
	} else {
		ln.SetAttribute(openxml.NewAttribute(NamespaceWML, "distance", PrefixW, strconv.Itoa(twips)))
	}
}

// Restart returns the line number restart value.
func (ln *LineNumberType) Restart() LineNumberRestartValue {
	attr, found := ln.GetAttribute(
		"restart",
		NamespaceWML,
	)
	if !found {
		return LineNumberRestartNewPage
	}

	return LineNumberRestartValue(attr.Value())
}

// SetRestart sets the line number restart value.
func (ln *LineNumberType) SetRestart(
	val LineNumberRestartValue,
) {
	if val == LineNumberRestartNewPage {
		ln.RemoveAttribute(
			"restart",
			NamespaceWML,
		)
	} else {
		ln.SetAttribute(openxml.NewAttribute(NamespaceWML, "restart", PrefixW, string(val)))
	}
}

// Clone creates a deep copy of this LineNumberType element.
func (ln *LineNumberType) Clone() openxml.Element {
	return &LineNumberType{
		LeafElementBase: ln.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this LineNumberType element.
func (ln *LineNumberType) CloneNode(
	deep bool,
) openxml.Element {
	return &LineNumberType{
		LeafElementBase: ln.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// TextDirectionValue represents text direction values.
type TextDirectionValue string

const (
	// TextDirectionLrTb is left-to-right, top-to-bottom (default).
	TextDirectionLrTb TextDirectionValue = "lrTb"
	// TextDirectionTbRl is top-to-bottom, right-to-left.
	TextDirectionTbRl TextDirectionValue = "tbRl"
	// TextDirectionBtLr is bottom-to-top, left-to-right.
	TextDirectionBtLr TextDirectionValue = "btLr"
	// TextDirectionLrTbV is left-to-right, top-to-bottom (rotated).
	TextDirectionLrTbV TextDirectionValue = "lrTbV"
	// TextDirectionTbRlV is top-to-bottom, right-to-left (rotated).
	TextDirectionTbRlV TextDirectionValue = "tbRlV"
	// TextDirectionTbLrV is top-to-bottom, left-to-right (rotated).
	TextDirectionTbLrV TextDirectionValue = "tbLrV"
)

// TextDirection represents text flow direction (w:textDirection).
type TextDirection struct {
	*openxml.LeafElementBase
}

// NewTextDirection creates a new TextDirection element.
func NewTextDirection(
	val TextDirectionValue,
) *TextDirection {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"textDirection",
		PrefixW,
	)
	td := &TextDirection{LeafElementBase: elem}
	td.SetValue(val)

	return td
}

// Value returns the text direction value.
func (td *TextDirection) Value() TextDirectionValue {
	attr, found := td.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return TextDirectionLrTb
	}

	return TextDirectionValue(attr.Value())
}

// SetValue sets the text direction value.
func (td *TextDirection) SetValue(
	val TextDirectionValue,
) {
	td.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			string(val),
		),
	)
}

// Clone creates a deep copy of this TextDirection element.
func (td *TextDirection) Clone() openxml.Element {
	return &TextDirection{
		LeafElementBase: td.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this TextDirection element.
func (td *TextDirection) CloneNode(
	deep bool,
) openxml.Element {
	return &TextDirection{
		LeafElementBase: td.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// RTLGutter represents right-to-left gutter (w:rtlGutter).
type RTLGutter struct {
	*openxml.LeafElementBase
}

// NewRTLGutter creates a new RTLGutter element.
func NewRTLGutter(enabled bool) *RTLGutter {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"rtlGutter",
		PrefixW,
	)
	rg := &RTLGutter{LeafElementBase: elem}
	if !enabled {
		rg.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameVal,
				PrefixW,
				attrValueFalse,
			),
		)
	}

	return rg
}

// Enabled returns whether RTL gutter is enabled.
func (rg *RTLGutter) Enabled() bool {
	attr, found := rg.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return true // default is true when element is present
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnabled sets whether RTL gutter is enabled.
func (rg *RTLGutter) SetEnabled(enabled bool) {
	if enabled {
		rg.RemoveAttribute(
			attrNameVal,
			NamespaceWML,
		)
	} else {
		rg.SetAttribute(openxml.NewAttribute(NamespaceWML, attrNameVal, PrefixW, attrValueFalse))
	}
}

// Clone creates a deep copy of this RTLGutter element.
func (rg *RTLGutter) Clone() openxml.Element {
	return &RTLGutter{
		LeafElementBase: rg.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this RTLGutter element.
func (rg *RTLGutter) CloneNode(
	deep bool,
) openxml.Element {
	return &RTLGutter{
		LeafElementBase: rg.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}

// DocGridTypeValue represents document grid type values.
type DocGridTypeValue string

const (
	// DocGridDefault uses no document grid.
	DocGridDefault DocGridTypeValue = "default"
	// DocGridLines uses line-only document grid.
	DocGridLines DocGridTypeValue = "lines"
	// DocGridLinesAndChars uses lines and characters document grid.
	DocGridLinesAndChars DocGridTypeValue = "linesAndChars"
	// DocGridSnapToChars uses snap to characters document grid.
	DocGridSnapToChars DocGridTypeValue = "snapToChars"
)

// DocGrid represents document grid settings (w:docGrid).
type DocGrid struct {
	*openxml.LeafElementBase
}

// NewDocGrid creates a new DocGrid element.
func NewDocGrid() *DocGrid {
	elem := openxml.NewLeafElement(
		NamespaceWML,
		"docGrid",
		PrefixW,
	)

	return &DocGrid{LeafElementBase: elem}
}

// Type returns the document grid type.
func (dg *DocGrid) Type() DocGridTypeValue {
	attr, found := dg.GetAttribute(
		attrNameType,
		NamespaceWML,
	)
	if !found {
		return DocGridDefault
	}

	return DocGridTypeValue(attr.Value())
}

// SetType sets the document grid type.
func (dg *DocGrid) SetType(val DocGridTypeValue) {
	if val == DocGridDefault {
		dg.RemoveAttribute(
			attrNameType,
			NamespaceWML,
		)
	} else {
		dg.SetAttribute(openxml.NewAttribute(NamespaceWML, attrNameType, PrefixW, string(val)))
	}
}

// LinePitch returns the line pitch in twips.
func (dg *DocGrid) LinePitch() int {
	attr, found := dg.GetAttribute(
		"linePitch",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetLinePitch sets the line pitch in twips.
func (dg *DocGrid) SetLinePitch(twips int) {
	if twips <= 0 {
		dg.RemoveAttribute(
			"linePitch",
			NamespaceWML,
		)
	} else {
		dg.SetAttribute(openxml.NewAttribute(NamespaceWML, "linePitch", PrefixW, strconv.Itoa(twips)))
	}
}

// CharSpace returns the character pitch adjustment.
func (dg *DocGrid) CharSpace() int {
	attr, found := dg.GetAttribute(
		"charSpace",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetCharSpace sets the character pitch adjustment.
func (dg *DocGrid) SetCharSpace(val int) {
	if val == 0 {
		dg.RemoveAttribute(
			"charSpace",
			NamespaceWML,
		)
	} else {
		dg.SetAttribute(openxml.NewAttribute(NamespaceWML, "charSpace", PrefixW, strconv.Itoa(val)))
	}
}

// Clone creates a deep copy of this DocGrid element.
func (dg *DocGrid) Clone() openxml.Element {
	return &DocGrid{
		LeafElementBase: dg.LeafElementBase.Clone().(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this DocGrid element.
func (dg *DocGrid) CloneNode(
	deep bool,
) openxml.Element {
	return &DocGrid{
		LeafElementBase: dg.LeafElementBase.CloneNode(deep).(*openxml.LeafElementBase),
	}
}
