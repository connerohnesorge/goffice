package elements

//revive:disable:file-length-limit many sheet properties

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// SheetPr represents the sheet properties element (x:sheetPr).
type SheetPr struct {
	*openxml.CompositeElementBase
}

// NewSheetPr creates a new SheetPr element.
func NewSheetPr() *SheetPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"sheetPr",
		PrefixDefault,
	)

	return &SheetPr{CompositeElementBase: elem}
}

// SyncHorizontal returns whether horizontal syncing is enabled.
func (sp *SheetPr) SyncHorizontal() bool {
	attr, found := sp.GetAttribute(
		"syncHorizontal",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSyncHorizontal sets whether horizontal syncing is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetSyncHorizontal(value bool) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"syncHorizontal",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("syncHorizontal", "")
	}
}

// SyncVertical returns whether vertical syncing is enabled.
func (sp *SheetPr) SyncVertical() bool {
	attr, found := sp.GetAttribute(
		"syncVertical",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetSyncVertical sets whether vertical syncing is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetSyncVertical(value bool) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"syncVertical",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("syncVertical", "")
	}
}

// SyncRef returns the synchronization reference.
func (sp *SheetPr) SyncRef() string {
	attr, found := sp.GetAttribute("syncRef", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetSyncRef sets the synchronization reference.
func (sp *SheetPr) SetSyncRef(ref string) {
	if ref == "" {
		sp.RemoveAttribute("syncRef", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"syncRef",
			"",
			ref,
		),
	)
}

// TransitionEvaluation returns whether transition formula evaluation is
// enabled.
func (sp *SheetPr) TransitionEvaluation() bool {
	attr, found := sp.GetAttribute(
		"transitionEvaluation",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetTransitionEvaluation sets whether transition formula evaluation is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetTransitionEvaluation(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"transitionEvaluation",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("transitionEvaluation", "")
	}
}

// TransitionEntry returns whether transition formula entry is enabled.
func (sp *SheetPr) TransitionEntry() bool {
	attr, found := sp.GetAttribute(
		"transitionEntry",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetTransitionEntry sets whether transition formula entry is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetTransitionEntry(
	value bool,
) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"transitionEntry",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("transitionEntry", "")
	}
}

// Published returns whether the sheet is published.
func (sp *SheetPr) Published() bool {
	attr, found := sp.GetAttribute(
		"published",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetPublished sets whether the sheet is published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetPublished(value bool) {
	if value {
		sp.RemoveAttribute("published", "")
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "published", "", attrValueFalse))
	}
}

// CodeName returns the VBA code name for the sheet.
func (sp *SheetPr) CodeName() string {
	attr, found := sp.GetAttribute("codeName", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCodeName sets the VBA code name for the sheet.
func (sp *SheetPr) SetCodeName(name string) {
	if name == "" {
		sp.RemoveAttribute("codeName", "")

		return
	}
	sp.SetAttribute(
		openxml.NewAttribute(
			"",
			"codeName",
			"",
			name,
		),
	)
}

// FilterMode returns whether the sheet is in filter mode.
func (sp *SheetPr) FilterMode() bool {
	attr, found := sp.GetAttribute(
		"filterMode",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFilterMode sets whether the sheet is in filter mode.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetFilterMode(value bool) {
	if value {
		sp.SetAttribute(
			openxml.NewAttribute(
				"",
				"filterMode",
				"",
				attrValueTrue,
			),
		)
	} else {
		sp.RemoveAttribute("filterMode", "")
	}
}

// EnableFormatConditionsCalculation returns whether format conditions should
// be calculated.
func (sp *SheetPr) EnableFormatConditionsCalculation() bool {
	attr, found := sp.GetAttribute(
		"enableFormatConditionsCalculation",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetEnableFormatConditionsCalculation sets whether format conditions should
// be calculated.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (sp *SheetPr) SetEnableFormatConditionsCalculation(
	value bool,
) {
	if value {
		sp.RemoveAttribute(
			"enableFormatConditionsCalculation",
			"",
		)
	} else {
		sp.SetAttribute(openxml.NewAttribute("", "enableFormatConditionsCalculation", "", attrValueFalse))
	}
}

// TabColor returns the tab color element, or nil if not present.
func (sp *SheetPr) TabColor() *TabColor {
	elem := sp.GetElement(
		"tabColor",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if tc, ok := elem.(*TabColor); ok {
		return tc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TabColor{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTabColor returns the tab color element, creating if needed.
func (sp *SheetPr) GetOrCreateTabColor() *TabColor {
	tc := sp.TabColor()
	if tc != nil {
		return tc
	}
	tc = NewTabColor()
	sp.AppendChild(tc)

	return tc
}

// OutlinePr returns the outline properties element, or nil if not present.
func (sp *SheetPr) OutlinePr() *OutlinePr {
	elem := sp.GetElement(
		"outlinePr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if op, ok := elem.(*OutlinePr); ok {
		return op
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &OutlinePr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateOutlinePr returns the outline properties, creating if needed.
func (sp *SheetPr) GetOrCreateOutlinePr() *OutlinePr {
	op := sp.OutlinePr()
	if op != nil {
		return op
	}
	op = NewOutlinePr()
	sp.AppendChild(op)

	return op
}

// PageSetUpPr returns the page setup properties element, or nil if not present.
func (sp *SheetPr) PageSetUpPr() *PageSetUpPr {
	elem := sp.GetElement(
		"pageSetUpPr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if psp, ok := elem.(*PageSetUpPr); ok {
		return psp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageSetUpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageSetUpPr returns the page setup properties, creating if needed.
func (sp *SheetPr) GetOrCreatePageSetUpPr() *PageSetUpPr {
	psp := sp.PageSetUpPr()
	if psp != nil {
		return psp
	}
	psp = NewPageSetUpPr()
	sp.AppendChild(psp)

	return psp
}

// Clone creates a deep copy of this SheetPr element.
func (sp *SheetPr) Clone() openxml.Element {
	cloned := sp.CompositeElementBase.Clone()

	return &SheetPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this SheetPr element.
func (sp *SheetPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := sp.CompositeElementBase.CloneNode(
		deep,
	)

	return &SheetPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TabColor represents the tab color element (x:tabColor).
type TabColor struct {
	*openxml.CompositeElementBase
}

// NewTabColor creates a new TabColor element.
func NewTabColor() *TabColor {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"tabColor",
		PrefixDefault,
	)

	return &TabColor{CompositeElementBase: elem}
}

// RGB returns the RGB color value (format: AARRGGBB or RRGGBB).
func (tc *TabColor) RGB() string {
	attr, found := tc.GetAttribute("rgb", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRGB sets the RGB color value.
func (tc *TabColor) SetRGB(rgb string) {
	if rgb == "" {
		tc.RemoveAttribute("rgb", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute("", "rgb", "", rgb),
	)
}

// Theme returns the theme color index.
func (tc *TabColor) Theme() string {
	attr, found := tc.GetAttribute("theme", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTheme sets the theme color index.
func (tc *TabColor) SetTheme(theme string) {
	if theme == "" {
		tc.RemoveAttribute("theme", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"theme",
			"",
			theme,
		),
	)
}

// Indexed returns the indexed color value.
func (tc *TabColor) Indexed() string {
	attr, found := tc.GetAttribute("indexed", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetIndexed sets the indexed color value.
func (tc *TabColor) SetIndexed(indexed string) {
	if indexed == "" {
		tc.RemoveAttribute("indexed", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"indexed",
			"",
			indexed,
		),
	)
}

// Tint returns the tint value for the color (-1.0 to 1.0).
func (tc *TabColor) Tint() string {
	attr, found := tc.GetAttribute("tint", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTint sets the tint value.
func (tc *TabColor) SetTint(tint string) {
	if tint == "" {
		tc.RemoveAttribute("tint", "")

		return
	}
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"tint",
			"",
			tint,
		),
	)
}

// Clone creates a deep copy of this TabColor element.
func (tc *TabColor) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &TabColor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this TabColor element.
func (tc *TabColor) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &TabColor{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// OutlinePr represents the outline properties element (x:outlinePr).
type OutlinePr struct {
	*openxml.CompositeElementBase
}

// NewOutlinePr creates a new OutlinePr element.
func NewOutlinePr() *OutlinePr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"outlinePr",
		PrefixDefault,
	)

	return &OutlinePr{CompositeElementBase: elem}
}

// ApplyStyles returns whether outline styles should be applied.
func (op *OutlinePr) ApplyStyles() bool {
	attr, found := op.GetAttribute(
		"applyStyles",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetApplyStyles sets whether outline styles should be applied.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (op *OutlinePr) SetApplyStyles(value bool) {
	if value {
		op.SetAttribute(
			openxml.NewAttribute(
				"",
				"applyStyles",
				"",
				attrValueTrue,
			),
		)
	} else {
		op.RemoveAttribute("applyStyles", "")
	}
}

// SummaryBelow returns whether summary rows appear below detail.
func (op *OutlinePr) SummaryBelow() bool {
	attr, found := op.GetAttribute(
		"summaryBelow",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSummaryBelow sets whether summary rows appear below detail.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (op *OutlinePr) SetSummaryBelow(value bool) {
	if value {
		op.RemoveAttribute("summaryBelow", "")
	} else {
		op.SetAttribute(openxml.NewAttribute("", "summaryBelow", "", attrValueFalse))
	}
}

// SummaryRight returns whether summary columns appear to the right of detail.
func (op *OutlinePr) SummaryRight() bool {
	attr, found := op.GetAttribute(
		"summaryRight",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetSummaryRight sets whether summary columns appear to the right of detail.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (op *OutlinePr) SetSummaryRight(value bool) {
	if value {
		op.RemoveAttribute("summaryRight", "")
	} else {
		op.SetAttribute(openxml.NewAttribute("", "summaryRight", "", attrValueFalse))
	}
}

// ShowOutlineSymbols returns whether outline symbols are shown.
func (op *OutlinePr) ShowOutlineSymbols() bool {
	attr, found := op.GetAttribute(
		"showOutlineSymbols",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetShowOutlineSymbols sets whether outline symbols are shown.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (op *OutlinePr) SetShowOutlineSymbols(
	value bool,
) {
	if value {
		op.RemoveAttribute(
			"showOutlineSymbols",
			"",
		)
	} else {
		op.SetAttribute(openxml.NewAttribute("", "showOutlineSymbols", "", attrValueFalse))
	}
}

// Clone creates a deep copy of this OutlinePr element.
func (op *OutlinePr) Clone() openxml.Element {
	cloned := op.CompositeElementBase.Clone()

	return &OutlinePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this OutlinePr element.
func (op *OutlinePr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := op.CompositeElementBase.CloneNode(
		deep,
	)

	return &OutlinePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// PageSetUpPr represents the page setup properties element (x:pageSetUpPr).
type PageSetUpPr struct {
	*openxml.CompositeElementBase
}

// NewPageSetUpPr creates a new PageSetUpPr element.
func NewPageSetUpPr() *PageSetUpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"pageSetUpPr",
		PrefixDefault,
	)

	return &PageSetUpPr{
		CompositeElementBase: elem,
	}
}

// AutoPageBreaks returns whether automatic page breaks are enabled.
func (p *PageSetUpPr) AutoPageBreaks() bool {
	attr, found := p.GetAttribute(
		"autoPageBreaks",
		"",
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero
}

// SetAutoPageBreaks sets whether automatic page breaks are enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PageSetUpPr) SetAutoPageBreaks(
	value bool,
) {
	if value {
		p.RemoveAttribute("autoPageBreaks", "")
	} else {
		p.SetAttribute(openxml.NewAttribute("", "autoPageBreaks", "", attrValueFalse))
	}
}

// FitToPage returns whether fit to page printing is enabled.
func (p *PageSetUpPr) FitToPage() bool {
	attr, found := p.GetAttribute(
		"fitToPage",
		"",
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetFitToPage sets whether fit to page printing is enabled.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (p *PageSetUpPr) SetFitToPage(value bool) {
	if value {
		p.SetAttribute(
			openxml.NewAttribute(
				"",
				"fitToPage",
				"",
				attrValueTrue,
			),
		)
	} else {
		p.RemoveAttribute("fitToPage", "")
	}
}

// Clone creates a deep copy of this PageSetUpPr element.
func (p *PageSetUpPr) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &PageSetUpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PageSetUpPr element.
func (p *PageSetUpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &PageSetUpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
