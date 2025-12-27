//nolint:revive // file-length-limit: comprehensive document settings with many element types
package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// ZoomViewType represents zoom view type values for document display.
type ZoomViewType string

const (
	// ZoomViewNone indicates no specific zoom view type.
	ZoomViewNone ZoomViewType = "none"
	// ZoomViewFullPage displays the full page.
	ZoomViewFullPage ZoomViewType = "fullPage"
	// ZoomViewBestFit fits the document to the window.
	ZoomViewBestFit ZoomViewType = "bestFit"
	// ZoomViewTextFit fits text to the window width.
	ZoomViewTextFit ZoomViewType = "textFit"
)

const (
	// defaultZoomPercent is the default zoom percentage for documents.
	defaultZoomPercent = 100
)

// Settings attribute name and value constants.
const (
	attrValueOff       = "off"
	attrNameFormatting = "formatting"
)

// DocumentProtectionType represents document protection edit restriction types.
type DocumentProtectionType string

const (
	// DocumentProtectionNone indicates no protection.
	DocumentProtectionNone DocumentProtectionType = "none"
	// DocumentProtectionReadOnly allows read-only access.
	DocumentProtectionReadOnly DocumentProtectionType = "readOnly"
	// DocumentProtectionComments allows only comments.
	DocumentProtectionComments DocumentProtectionType = "comments"
	// DocumentProtectionTrackedChanges allows only tracked changes.
	DocumentProtectionTrackedChanges DocumentProtectionType = "trackedChanges"
	// DocumentProtectionForms allows only form field editing.
	DocumentProtectionForms DocumentProtectionType = "forms"
)

// ProofStateValue represents proofing state values.
type ProofStateValue string

const (
	// ProofStateClean indicates proofing is complete.
	ProofStateClean ProofStateValue = "clean"
	// ProofStateDirty indicates proofing is needed.
	ProofStateDirty ProofStateValue = "dirty"
)

// Settings represents the root element for the document settings part
// (w:settings).
type Settings struct {
	*openxml.CompositeElementBase
}

// NewSettings creates a new Settings element.
func NewSettings() *Settings {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"settings",
		PrefixW,
	)

	return &Settings{CompositeElementBase: elem}
}

// Zoom returns the zoom settings element, or nil if not present.
func (s *Settings) Zoom() *Zoom {
	elem := s.GetElement("zoom", NamespaceWML)
	if elem == nil {
		return nil
	}
	if z, ok := elem.(*Zoom); ok {
		return z
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Zoom{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateZoom returns the zoom settings, creating if needed.
func (s *Settings) GetOrCreateZoom() *Zoom {
	z := s.Zoom()
	if z != nil {
		return z
	}
	z = NewZoom()
	s.AppendChild(z)

	return z
}

// SetZoom sets the zoom percentage.
func (s *Settings) SetZoom(percent int) {
	z := s.GetOrCreateZoom()
	z.SetPercent(percent)
}

// DefaultTabStop returns the default tab stop element, or nil if not present.
func (s *Settings) DefaultTabStop() *DefaultTabStop {
	elem := s.GetElement(
		"defaultTabStop",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if dt, ok := elem.(*DefaultTabStop); ok {
		return dt
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DefaultTabStop{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDefaultTabStop returns the default tab stop, creating if needed.
func (s *Settings) GetOrCreateDefaultTabStop() *DefaultTabStop {
	dt := s.DefaultTabStop()
	if dt != nil {
		return dt
	}
	dt = NewDefaultTabStop()
	s.AppendChild(dt)

	return dt
}

// DocumentProtection returns the document protection element, or nil if not
// present.
func (s *Settings) DocumentProtection() *DocumentProtection {
	elem := s.GetElement(
		"documentProtection",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if dp, ok := elem.(*DocumentProtection); ok {
		return dp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DocumentProtection{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDocumentProtection returns the document protection, creating if
// needed.
func (s *Settings) GetOrCreateDocumentProtection() *DocumentProtection {
	dp := s.DocumentProtection()
	if dp != nil {
		return dp
	}
	dp = NewDocumentProtection()
	s.AppendChild(dp)

	return dp
}

// TrackRevisions returns whether revision tracking is enabled.
func (s *Settings) TrackRevisions() bool {
	return s.hasOnOffElement("trackRevisions")
}

// SetTrackRevisions sets whether revision tracking is enabled.
func (s *Settings) SetTrackRevisions(b bool) {
	s.setOnOffElement("trackRevisions", b)
}

// MirrorMargins returns whether margins mirror for binding.
func (s *Settings) MirrorMargins() bool {
	return s.hasOnOffElement("mirrorMargins")
}

// SetMirrorMargins sets whether margins mirror for binding.
func (s *Settings) SetMirrorMargins(b bool) {
	s.setOnOffElement("mirrorMargins", b)
}

// EvenAndOddHeaders returns whether different even/odd headers are used.
func (s *Settings) EvenAndOddHeaders() bool {
	return s.hasOnOffElement("evenAndOddHeaders")
}

// SetEvenAndOddHeaders sets whether different even/odd headers are used.
func (s *Settings) SetEvenAndOddHeaders(b bool) {
	s.setOnOffElement("evenAndOddHeaders", b)
}

// DisplayBackgroundShape returns whether the background shape is displayed.
func (s *Settings) DisplayBackgroundShape() bool {
	return s.hasOnOffElement(
		"displayBackgroundShape",
	)
}

// SetDisplayBackgroundShape sets whether the background shape is displayed.
func (s *Settings) SetDisplayBackgroundShape(
	b bool,
) {
	s.setOnOffElement("displayBackgroundShape", b)
}

// HideSpellingErrors returns whether spelling errors are hidden.
func (s *Settings) HideSpellingErrors() bool {
	return s.hasOnOffElement("hideSpellingErrors")
}

// SetHideSpellingErrors sets whether spelling errors are hidden.
func (s *Settings) SetHideSpellingErrors(b bool) {
	s.setOnOffElement("hideSpellingErrors", b)
}

// HideGrammaticalErrors returns whether grammatical errors are hidden.
func (s *Settings) HideGrammaticalErrors() bool {
	return s.hasOnOffElement(
		"hideGrammaticalErrors",
	)
}

// SetHideGrammaticalErrors sets whether grammatical errors are hidden.
func (s *Settings) SetHideGrammaticalErrors(
	b bool,
) {
	s.setOnOffElement("hideGrammaticalErrors", b)
}

// Compatibility returns the compatibility settings element, or nil if not
// present.
func (s *Settings) Compatibility() *Compatibility {
	elem := s.GetElement("compat", NamespaceWML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Compatibility); ok {
		return c
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Compatibility{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCompatibility returns the compatibility settings, creating if
// needed.
func (s *Settings) GetOrCreateCompatibility() *Compatibility {
	c := s.Compatibility()
	if c != nil {
		return c
	}
	c = NewCompatibility()
	s.AppendChild(c)

	return c
}

// MailMerge returns the mail merge settings element, or nil if not present.
func (s *Settings) MailMerge() *MailMerge {
	elem := s.GetElement(
		"mailMerge",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if mm, ok := elem.(*MailMerge); ok {
		return mm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &MailMerge{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateMailMerge returns the mail merge settings, creating if needed.
func (s *Settings) GetOrCreateMailMerge() *MailMerge {
	mm := s.MailMerge()
	if mm != nil {
		return mm
	}
	mm = NewMailMerge()
	s.AppendChild(mm)

	return mm
}

// WriteProtection returns the write protection element, or nil if not present.
func (s *Settings) WriteProtection() *WriteProtection {
	elem := s.GetElement(
		"writeProtection",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if wp, ok := elem.(*WriteProtection); ok {
		return wp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &WriteProtection{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateWriteProtection returns the write protection, creating if needed.
func (s *Settings) GetOrCreateWriteProtection() *WriteProtection {
	wp := s.WriteProtection()
	if wp != nil {
		return wp
	}
	wp = NewWriteProtection()
	s.AppendChild(wp)

	return wp
}

// RsidRoot returns the original document RSID.
func (s *Settings) RsidRoot() string {
	elem := s.GetElement("rsidRoot", NamespaceWML)
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

// SetRsidRoot sets the original document RSID.
func (s *Settings) SetRsidRoot(rsid string) {
	if rsid == "" {
		s.removeElement("rsidRoot")

		return
	}
	elem := s.getOrCreateElement("rsidRoot")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			rsid,
		),
	)
}

// DocumentVariables returns the document variables element, or nil if not
// present.
func (s *Settings) DocumentVariables() *DocumentVariables {
	elem := s.GetElement("docVars", NamespaceWML)
	if elem == nil {
		return nil
	}
	if dv, ok := elem.(*DocumentVariables); ok {
		return dv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DocumentVariables{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDocumentVariables returns the document variables, creating if
// needed.
func (s *Settings) GetOrCreateDocumentVariables() *DocumentVariables {
	dv := s.DocumentVariables()
	if dv != nil {
		return dv
	}
	dv = NewDocumentVariables()
	s.AppendChild(dv)

	return dv
}

// ProofState returns the proof state element, or nil if not present.
func (s *Settings) ProofState() *ProofState {
	elem := s.GetElement(
		"proofState",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if ps, ok := elem.(*ProofState); ok {
		return ps
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ProofState{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateProofState returns the proof state, creating if needed.
func (s *Settings) GetOrCreateProofState() *ProofState {
	ps := s.ProofState()
	if ps != nil {
		return ps
	}
	ps = NewProofState()
	s.AppendChild(ps)

	return ps
}

// RevisionView returns the revision view element, or nil if not present.
func (s *Settings) RevisionView() *RevisionView {
	elem := s.GetElement(
		"revisionView",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if rv, ok := elem.(*RevisionView); ok {
		return rv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RevisionView{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRevisionView returns the revision view, creating if needed.
func (s *Settings) GetOrCreateRevisionView() *RevisionView {
	rv := s.RevisionView()
	if rv != nil {
		return rv
	}
	rv = NewRevisionView()
	s.AppendChild(rv)

	return rv
}

// ThemeFontLang returns the theme font language element, or nil if not present.
func (s *Settings) ThemeFontLang() *ThemeFontLang {
	elem := s.GetElement(
		"themeFontLang",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if tfl, ok := elem.(*ThemeFontLang); ok {
		return tfl
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ThemeFontLang{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateThemeFontLang returns the theme font language, creating if needed.
func (s *Settings) GetOrCreateThemeFontLang() *ThemeFontLang {
	tfl := s.ThemeFontLang()
	if tfl != nil {
		return tfl
	}
	tfl = NewThemeFontLang()
	s.AppendChild(tfl)

	return tfl
}

// Helper methods

func (s *Settings) hasOnOffElement(
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

//
//nolint:revive // flag-parameter: bool setter required for on/off element
func (s *Settings) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		s.getOrCreateElement(name)
	} else {
		s.removeElement(name)
	}
}

func (s *Settings) getOrCreateElement(
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

func (s *Settings) removeElement(name string) {
	elem := s.GetElement(name, NamespaceWML)
	if elem != nil {
		s.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Settings element.
func (s *Settings) Clone() openxml.Element {
	cloned := s.CompositeElementBase.Clone()

	return &Settings{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Settings element.
func (s *Settings) CloneNode(
	deep bool,
) openxml.Element {
	cloned := s.CompositeElementBase.CloneNode(
		deep,
	)

	return &Settings{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Zoom represents the zoom settings element (w:zoom).
type Zoom struct {
	*openxml.CompositeElementBase
}

// NewZoom creates a new Zoom element.
func NewZoom() *Zoom {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"zoom",
		PrefixW,
	)

	return &Zoom{CompositeElementBase: elem}
}

// Percent returns the zoom percentage.
func (z *Zoom) Percent() int {
	attr, found := z.GetAttribute(
		"percent",
		NamespaceWML,
	)
	if !found {
		return defaultZoomPercent // Default zoom
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetPercent sets the zoom percentage.
func (z *Zoom) SetPercent(p int) {
	z.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"percent",
			PrefixW,
			strconv.Itoa(p),
		),
	)
}

// Val returns the zoom view type.
func (z *Zoom) Val() ZoomViewType {
	attr, found := z.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return ZoomViewNone
	}

	return ZoomViewType(attr.Value())
}

// SetVal sets the zoom view type.
func (z *Zoom) SetVal(t ZoomViewType) {
	if t == ZoomViewNone {
		z.RemoveAttribute(
			attrNameVal,
			NamespaceWML,
		)

		return
	}
	z.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			string(t),
		),
	)
}

// Clone creates a deep copy of this Zoom element.
func (z *Zoom) Clone() openxml.Element {
	cloned := z.CompositeElementBase.Clone()

	return &Zoom{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Zoom element.
func (z *Zoom) CloneNode(
	deep bool,
) openxml.Element {
	cloned := z.CompositeElementBase.CloneNode(
		deep,
	)

	return &Zoom{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DefaultTabStop represents the default tab stop width element
// (w:defaultTabStop).
type DefaultTabStop struct {
	*openxml.CompositeElementBase
}

// NewDefaultTabStop creates a new DefaultTabStop element.
func NewDefaultTabStop() *DefaultTabStop {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"defaultTabStop",
		PrefixW,
	)

	return &DefaultTabStop{
		CompositeElementBase: elem,
	}
}

// Val returns the default tab stop width in twips.
func (dt *DefaultTabStop) Val() int {
	attr, found := dt.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return defaultHeaderFooter // Default 0.5 inch
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetVal sets the default tab stop width in twips.
func (dt *DefaultTabStop) SetVal(twips int) {
	dt.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(twips),
		),
	)
}

// Clone creates a deep copy of this DefaultTabStop element.
func (dt *DefaultTabStop) Clone() openxml.Element {
	cloned := dt.CompositeElementBase.Clone()

	return &DefaultTabStop{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DefaultTabStop element.
func (dt *DefaultTabStop) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dt.CompositeElementBase.CloneNode(
		deep,
	)

	return &DefaultTabStop{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DocumentProtection represents document protection settings
// (w:documentProtection).
type DocumentProtection struct {
	*openxml.CompositeElementBase
}

// NewDocumentProtection creates a new DocumentProtection element.
func NewDocumentProtection() *DocumentProtection {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"documentProtection",
		PrefixW,
	)

	return &DocumentProtection{
		CompositeElementBase: elem,
	}
}

// Edit returns the protection type.
func (dp *DocumentProtection) Edit() DocumentProtectionType {
	attr, found := dp.GetAttribute(
		"edit",
		NamespaceWML,
	)
	if !found {
		return DocumentProtectionNone
	}

	return DocumentProtectionType(attr.Value())
}

// SetEdit sets the protection type.
func (dp *DocumentProtection) SetEdit(
	t DocumentProtectionType,
) {
	if t == DocumentProtectionNone {
		dp.RemoveAttribute("edit", NamespaceWML)

		return
	}
	dp.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"edit",
			PrefixW,
			string(t),
		),
	)
}

// Enforcement returns whether protection is enforced.
func (dp *DocumentProtection) Enforcement() bool {
	attr, found := dp.GetAttribute(
		"enforcement",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == "true" ||
		val == "on"
}

// SetEnforcement sets whether protection is enforced.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (dp *DocumentProtection) SetEnforcement(
	b bool,
) {
	if b {
		dp.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"enforcement",
				PrefixW,
				attrValueOne,
			),
		)
	} else {
		dp.RemoveAttribute("enforcement", NamespaceWML)
	}
}

// Formatting returns whether formatting is restricted.
func (dp *DocumentProtection) Formatting() bool {
	attr, found := dp.GetAttribute(
		attrNameFormatting,
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == "true" ||
		val == "on"
}

// SetFormatting sets whether formatting is restricted.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (dp *DocumentProtection) SetFormatting(
	b bool,
) {
	if b {
		dp.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				attrNameFormatting,
				PrefixW,
				attrValueOne,
			),
		)
	} else {
		dp.RemoveAttribute(attrNameFormatting, NamespaceWML)
	}
}

// Hash returns the password hash.
func (dp *DocumentProtection) Hash() string {
	attr, found := dp.GetAttribute(
		"hash",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Salt returns the password salt.
func (dp *DocumentProtection) Salt() string {
	attr, found := dp.GetAttribute(
		"salt",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// Clone creates a deep copy of this DocumentProtection element.
func (dp *DocumentProtection) Clone() openxml.Element {
	cloned := dp.CompositeElementBase.Clone()

	return &DocumentProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DocumentProtection element.
func (dp *DocumentProtection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dp.CompositeElementBase.CloneNode(
		deep,
	)

	return &DocumentProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Compatibility represents compatibility settings (w:compat).
type Compatibility struct {
	*openxml.CompositeElementBase
}

// NewCompatibility creates a new Compatibility element.
func NewCompatibility() *Compatibility {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"compat",
		PrefixW,
	)

	return &Compatibility{
		CompositeElementBase: elem,
	}
}

// UseFELayout returns whether Far East layout is used.
func (c *Compatibility) UseFELayout() bool {
	return c.hasOnOffElement("useFELayout")
}

// SetUseFELayout sets whether Far East layout is used.
func (c *Compatibility) SetUseFELayout(b bool) {
	c.setOnOffElement("useFELayout", b)
}

// UseWord2003TableStyleRules returns whether Word 2003 table style rules are used.
func (c *Compatibility) UseWord2003TableStyleRules() bool {
	return c.hasOnOffElement(
		"useWord2003TableStyleRules",
	)
}

// SetUseWord2003TableStyleRules sets whether Word 2003 table style rules are used.
func (c *Compatibility) SetUseWord2003TableStyleRules(
	b bool,
) {
	c.setOnOffElement(
		"useWord2003TableStyleRules",
		b,
	)
}

// GrowAutofit returns whether autofit grows.
func (c *Compatibility) GrowAutofit() bool {
	return c.hasOnOffElement("growAutofit")
}

// SetGrowAutofit sets whether autofit grows.
func (c *Compatibility) SetGrowAutofit(b bool) {
	c.setOnOffElement("growAutofit", b)
}

// CompatSettings returns an iterator over compatibility settings.
func (c *Compatibility) CompatSettings() iter.Seq[*CompatSetting] {
	return func(yield func(*CompatSetting) bool) {
		for child := range c.Children() {
			if child.LocalName() != "compatSetting" ||
				child.NamespaceURI() != NamespaceWML {
				continue
			}
			var cs *CompatSetting
			if setting, ok := child.(*CompatSetting); ok {
				cs = setting
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				cs = &CompatSetting{CompositeElementBase: comp}
			}
			if cs != nil && !yield(cs) {
				return
			}
		}
	}
}

// AddCompatSetting adds a compatibility setting.
func (c *Compatibility) AddCompatSetting(
	name, uri, val string,
) *CompatSetting {
	cs := NewCompatSetting()
	cs.SetName(name)
	cs.SetUri(uri)
	cs.SetVal(val)
	c.AppendChild(cs)

	return cs
}

// Helper methods

func (c *Compatibility) hasOnOffElement(
	name string,
) bool {
	elem := c.GetElement(name, NamespaceWML)
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

//
//nolint:revive // flag-parameter: bool setter required for on/off element
func (c *Compatibility) setOnOffElement(
	name string,
	value bool,
) {
	if value {
		c.getOrCreateElement(name)
	} else {
		c.removeElement(name)
	}
}

func (c *Compatibility) getOrCreateElement(
	name string,
) openxml.Element {
	elem := c.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	c.AppendChild(newElem)

	return newElem
}

func (c *Compatibility) removeElement(
	name string,
) {
	elem := c.GetElement(name, NamespaceWML)
	if elem != nil {
		c.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Compatibility element.
func (c *Compatibility) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Compatibility{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Compatibility element.
func (c *Compatibility) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Compatibility{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CompatSetting represents a compatibility setting (w:compatSetting).
type CompatSetting struct {
	*openxml.CompositeElementBase
}

// NewCompatSetting creates a new CompatSetting element.
func NewCompatSetting() *CompatSetting {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"compatSetting",
		PrefixW,
	)

	return &CompatSetting{
		CompositeElementBase: elem,
	}
}

// Name returns the setting name.
func (cs *CompatSetting) Name() string {
	attr, found := cs.GetAttribute(
		"name",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the setting name.
func (cs *CompatSetting) SetName(name string) {
	cs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"name",
			PrefixW,
			name,
		),
	)
}

// Uri returns the setting URI.
func (cs *CompatSetting) Uri() string {
	attr, found := cs.GetAttribute(
		"uri",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUri sets the setting URI.
func (cs *CompatSetting) SetUri(uri string) {
	cs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"uri",
			PrefixW,
			uri,
		),
	)
}

// Val returns the setting value.
func (cs *CompatSetting) Val() string {
	attr, found := cs.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the setting value.
func (cs *CompatSetting) SetVal(val string) {
	cs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			val,
		),
	)
}

// Clone creates a deep copy of this CompatSetting element.
func (cs *CompatSetting) Clone() openxml.Element {
	cloned := cs.CompositeElementBase.Clone()

	return &CompatSetting{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CompatSetting element.
func (cs *CompatSetting) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cs.CompositeElementBase.CloneNode(
		deep,
	)

	return &CompatSetting{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DocumentVariables represents document variables (w:docVars).
type DocumentVariables struct {
	*openxml.CompositeElementBase
}

// NewDocumentVariables creates a new DocumentVariables element.
func NewDocumentVariables() *DocumentVariables {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"docVars",
		PrefixW,
	)

	return &DocumentVariables{
		CompositeElementBase: elem,
	}
}

// Variables returns an iterator over document variables.
func (dv *DocumentVariables) Variables() iter.Seq[*DocumentVariable] {
	return func(yield func(*DocumentVariable) bool) {
		for child := range dv.Children() {
			if child.LocalName() != "docVar" ||
				child.NamespaceURI() != NamespaceWML {
				continue
			}
			var v *DocumentVariable
			if variable, ok := child.(*DocumentVariable); ok {
				v = variable
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				v = &DocumentVariable{CompositeElementBase: comp}
			}
			if v != nil && !yield(v) {
				return
			}
		}
	}
}

// GetVariable returns the value of a variable by name.
func (dv *DocumentVariables) GetVariable(
	name string,
) string {
	for v := range dv.Variables() {
		if v.Name() == name {
			return v.Val()
		}
	}

	return ""
}

// SetVariable sets or adds a variable.
func (dv *DocumentVariables) SetVariable(
	name, value string,
) {
	// Look for existing variable
	for v := range dv.Variables() {
		if v.Name() == name {
			v.SetVal(value)

			return
		}
	}
	// Create new variable
	v := NewDocumentVariable()
	v.SetName(name)
	v.SetVal(value)
	dv.AppendChild(v)
}

// RemoveVariable removes a variable by name.
func (dv *DocumentVariables) RemoveVariable(
	name string,
) {
	for child := range dv.Children() {
		if child.LocalName() != "docVar" ||
			child.NamespaceURI() != NamespaceWML {
			continue
		}
		attr, found := child.GetAttribute(
			attrNameName,
			NamespaceWML,
		)
		if found && attr.Value() == name {
			dv.RemoveChild(child)

			return
		}
	}
}

// Clone creates a deep copy of this DocumentVariables element.
func (dv *DocumentVariables) Clone() openxml.Element {
	cloned := dv.CompositeElementBase.Clone()

	return &DocumentVariables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DocumentVariables element.
func (dv *DocumentVariables) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dv.CompositeElementBase.CloneNode(
		deep,
	)

	return &DocumentVariables{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// DocumentVariable represents a single document variable (w:docVar).
type DocumentVariable struct {
	*openxml.CompositeElementBase
}

// NewDocumentVariable creates a new DocumentVariable element.
func NewDocumentVariable() *DocumentVariable {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"docVar",
		PrefixW,
	)

	return &DocumentVariable{
		CompositeElementBase: elem,
	}
}

// Name returns the variable name.
func (v *DocumentVariable) Name() string {
	attr, found := v.GetAttribute(
		attrNameName,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the variable name.
func (v *DocumentVariable) SetName(name string) {
	v.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameName,
			PrefixW,
			name,
		),
	)
}

// Val returns the variable value.
func (v *DocumentVariable) Val() string {
	attr, found := v.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the variable value.
func (v *DocumentVariable) SetVal(val string) {
	v.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			val,
		),
	)
}

// Clone creates a deep copy of this DocumentVariable element.
func (v *DocumentVariable) Clone() openxml.Element {
	cloned := v.CompositeElementBase.Clone()

	return &DocumentVariable{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DocumentVariable element.
func (v *DocumentVariable) CloneNode(
	deep bool,
) openxml.Element {
	cloned := v.CompositeElementBase.CloneNode(
		deep,
	)

	return &DocumentVariable{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// WriteProtection represents write protection settings (w:writeProtection).
type WriteProtection struct {
	*openxml.CompositeElementBase
}

// NewWriteProtection creates a new WriteProtection element.
func NewWriteProtection() *WriteProtection {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"writeProtection",
		PrefixW,
	)

	return &WriteProtection{
		CompositeElementBase: elem,
	}
}

// Recommended returns whether read-only is recommended.
func (wp *WriteProtection) Recommended() bool {
	attr, found := wp.GetAttribute(
		"recommended",
		NamespaceWML,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" || val == "true" ||
		val == "on"
}

// SetRecommended sets whether read-only is recommended.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (wp *WriteProtection) SetRecommended(
	b bool,
) {
	if b {
		wp.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"recommended",
				PrefixW,
				attrValueOne,
			),
		)
	} else {
		wp.RemoveAttribute("recommended", NamespaceWML)
	}
}

// Clone creates a deep copy of this WriteProtection element.
func (wp *WriteProtection) Clone() openxml.Element {
	cloned := wp.CompositeElementBase.Clone()

	return &WriteProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this WriteProtection element.
func (wp *WriteProtection) CloneNode(
	deep bool,
) openxml.Element {
	cloned := wp.CompositeElementBase.CloneNode(
		deep,
	)

	return &WriteProtection{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ProofState represents proofing state settings (w:proofState).
type ProofState struct {
	*openxml.CompositeElementBase
}

// NewProofState creates a new ProofState element.
func NewProofState() *ProofState {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"proofState",
		PrefixW,
	)

	return &ProofState{CompositeElementBase: elem}
}

// Spelling returns the spelling proof state.
func (ps *ProofState) Spelling() ProofStateValue {
	attr, found := ps.GetAttribute(
		"spelling",
		NamespaceWML,
	)
	if !found {
		return ProofStateDirty
	}

	return ProofStateValue(attr.Value())
}

// SetSpelling sets the spelling proof state.
func (ps *ProofState) SetSpelling(
	v ProofStateValue,
) {
	ps.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"spelling",
			PrefixW,
			string(v),
		),
	)
}

// Grammar returns the grammar proof state.
func (ps *ProofState) Grammar() ProofStateValue {
	attr, found := ps.GetAttribute(
		"grammar",
		NamespaceWML,
	)
	if !found {
		return ProofStateDirty
	}

	return ProofStateValue(attr.Value())
}

// SetGrammar sets the grammar proof state.
func (ps *ProofState) SetGrammar(
	v ProofStateValue,
) {
	ps.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"grammar",
			PrefixW,
			string(v),
		),
	)
}

// Clone creates a deep copy of this ProofState element.
func (ps *ProofState) Clone() openxml.Element {
	cloned := ps.CompositeElementBase.Clone()

	return &ProofState{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ProofState element.
func (ps *ProofState) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ps.CompositeElementBase.CloneNode(
		deep,
	)

	return &ProofState{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// RevisionView represents revision view settings (w:revisionView).
type RevisionView struct {
	*openxml.CompositeElementBase
}

// NewRevisionView creates a new RevisionView element.
func NewRevisionView() *RevisionView {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"revisionView",
		PrefixW,
	)

	return &RevisionView{
		CompositeElementBase: elem,
	}
}

// Markup returns whether markup is shown.
func (rv *RevisionView) Markup() bool {
	attr, found := rv.GetAttribute(
		"markup",
		NamespaceWML,
	)
	if !found {
		return true // Default is true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetMarkup sets whether markup is shown.
func (rv *RevisionView) SetMarkup(b bool) {
	if b {
		rv.RemoveAttribute("markup", NamespaceWML)
	} else {
		rv.SetAttribute(openxml.NewAttribute(NamespaceWML, "markup", PrefixW, attrValueZero))
	}
}

// Comments returns whether comments are shown.
func (rv *RevisionView) Comments() bool {
	attr, found := rv.GetAttribute(
		"comments",
		NamespaceWML,
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetComments sets whether comments are shown.
func (rv *RevisionView) SetComments(b bool) {
	if b {
		rv.RemoveAttribute(
			"comments",
			NamespaceWML,
		)
	} else {
		rv.SetAttribute(openxml.NewAttribute(NamespaceWML, "comments", PrefixW, attrValueZero))
	}
}

// InsertionsAndDeletions returns whether insertions and deletions are shown.
func (rv *RevisionView) InsertionsAndDeletions() bool {
	attr, found := rv.GetAttribute(
		"insDel",
		NamespaceWML,
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetInsertionsAndDeletions sets whether insertions and deletions are shown.
func (rv *RevisionView) SetInsertionsAndDeletions(
	b bool,
) {
	if b {
		rv.RemoveAttribute("insDel", NamespaceWML)
	} else {
		rv.SetAttribute(openxml.NewAttribute(NamespaceWML, "insDel", PrefixW, attrValueZero))
	}
}

// Formatting returns whether formatting changes are shown.
func (rv *RevisionView) Formatting() bool {
	attr, found := rv.GetAttribute(
		attrNameFormatting,
		NamespaceWML,
	)
	if !found {
		return true
	}
	val := attr.Value()

	return val != attrValueFalse &&
		val != attrValueZero &&
		val != attrValueOff
}

// SetFormatting sets whether formatting changes are shown.
func (rv *RevisionView) SetFormatting(b bool) {
	if b {
		rv.RemoveAttribute(
			attrNameFormatting,
			NamespaceWML,
		)
	} else {
		attr := openxml.NewAttribute(
			NamespaceWML,
			attrNameFormatting,
			PrefixW,
			attrValueZero,
		)
		rv.SetAttribute(attr)
	}
}

// Clone creates a deep copy of this RevisionView element.
func (rv *RevisionView) Clone() openxml.Element {
	cloned := rv.CompositeElementBase.Clone()

	return &RevisionView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this RevisionView element.
func (rv *RevisionView) CloneNode(
	deep bool,
) openxml.Element {
	cloned := rv.CompositeElementBase.CloneNode(
		deep,
	)

	return &RevisionView{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// MailMerge represents mail merge settings (w:mailMerge).
type MailMerge struct {
	*openxml.CompositeElementBase
}

// NewMailMerge creates a new MailMerge element.
func NewMailMerge() *MailMerge {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"mailMerge",
		PrefixW,
	)

	return &MailMerge{CompositeElementBase: elem}
}

// MainDocumentType returns the main document type.
func (mm *MailMerge) MainDocumentType() string {
	elem := mm.GetElement(
		"mainDocumentType",
		NamespaceWML,
	)
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

// SetMainDocumentType sets the main document type.
func (mm *MailMerge) SetMainDocumentType(
	docType string,
) {
	elem := mm.getOrCreateElement(
		"mainDocumentType",
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			docType,
		),
	)
}

// DataType returns the data source type.
func (mm *MailMerge) DataType() string {
	elem := mm.GetElement(
		"dataType",
		NamespaceWML,
	)
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

// SetDataType sets the data source type.
func (mm *MailMerge) SetDataType(
	dataType string,
) {
	elem := mm.getOrCreateElement("dataType")
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			dataType,
		),
	)
}

func (mm *MailMerge) getOrCreateElement(
	name string,
) openxml.Element {
	elem := mm.GetElement(name, NamespaceWML)
	if elem != nil {
		return elem
	}
	newElem := openxml.NewCompositeElement(
		NamespaceWML,
		name,
		PrefixW,
	)
	mm.AppendChild(newElem)

	return newElem
}

// Clone creates a deep copy of this MailMerge element.
func (mm *MailMerge) Clone() openxml.Element {
	cloned := mm.CompositeElementBase.Clone()

	return &MailMerge{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this MailMerge element.
func (mm *MailMerge) CloneNode(
	deep bool,
) openxml.Element {
	cloned := mm.CompositeElementBase.CloneNode(
		deep,
	)

	return &MailMerge{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ThemeFontLang represents theme font language settings (w:themeFontLang).
type ThemeFontLang struct {
	*openxml.CompositeElementBase
}

// NewThemeFontLang creates a new ThemeFontLang element.
func NewThemeFontLang() *ThemeFontLang {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"themeFontLang",
		PrefixW,
	)

	return &ThemeFontLang{
		CompositeElementBase: elem,
	}
}

// Val returns the primary language.
func (t *ThemeFontLang) Val() string {
	attr, found := t.GetAttribute(
		attrNameVal,
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetVal sets the primary language.
func (t *ThemeFontLang) SetVal(lang string) {
	if lang == "" {
		t.RemoveAttribute(
			attrNameVal,
			NamespaceWML,
		)

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			attrNameVal,
			PrefixW,
			lang,
		),
	)
}

// EastAsia returns the East Asian language.
func (t *ThemeFontLang) EastAsia() string {
	attr, found := t.GetAttribute(
		"eastAsia",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetEastAsia sets the East Asian language.
func (t *ThemeFontLang) SetEastAsia(
	lang string,
) {
	if lang == "" {
		t.RemoveAttribute(
			"eastAsia",
			NamespaceWML,
		)

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"eastAsia",
			PrefixW,
			lang,
		),
	)
}

// Bidi returns the bidirectional language.
func (t *ThemeFontLang) Bidi() string {
	attr, found := t.GetAttribute(
		"bidi",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetBidi sets the bidirectional language.
func (t *ThemeFontLang) SetBidi(lang string) {
	if lang == "" {
		t.RemoveAttribute("bidi", NamespaceWML)

		return
	}
	t.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"bidi",
			PrefixW,
			lang,
		),
	)
}

// Clone creates a deep copy of this ThemeFontLang element.
func (t *ThemeFontLang) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &ThemeFontLang{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ThemeFontLang element.
func (t *ThemeFontLang) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &ThemeFontLang{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
