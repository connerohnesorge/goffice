//nolint:revive // This file contains slide element implementation.
package elements

import (
	"encoding/xml"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/types"
)

const (
	// attrValueTrue represents the string value "true" for XML attributes.
	attrValueTrue = "true"
)

// Slide represents a slide element (p:sld).
type Slide struct {
	*openxml.PartRootElementBase
}

// NewSlide creates a new Slide element.
func NewSlide() *Slide {
	elem := openxml.NewPartRootElement(
		NamespacePresentationML,
		"sld",
		PrefixP,
	)
	s := &Slide{PartRootElementBase: elem}

	// Add namespace declarations for all namespaces used in slides
	// This is required for Office/LibreOffice to correctly parse the XML
	s.SetAttribute(openxml.NewAttribute(
		"",
		PrefixA,
		"xmlns",
		NamespaceDrawingML,
	))
	s.SetAttribute(openxml.NewAttribute(
		"",
		PrefixR,
		"xmlns",
		NamespaceRelationships,
	))
	s.SetAttribute(openxml.NewAttribute(
		"",
		PrefixC,
		"xmlns",
		NamespaceDrawingMLChart,
	))

	// Add required common slide data
	s.AppendChild(NewCommonSlideData())

	return s
}

// CommonSlideData returns the common slide data element.
func (s *Slide) CommonSlideData() *CommonSlideData {
	elem := s.GetElement(
		"cSld",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if csd, ok := elem.(*CommonSlideData); ok {
		return csd
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &CommonSlideData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCommonSlideData returns the common slide data, creating if needed.
func (s *Slide) GetOrCreateCommonSlideData() *CommonSlideData {
	csd := s.CommonSlideData()
	if csd != nil {
		return csd
	}
	csd = NewCommonSlideData()
	// Insert as first child
	if first := s.FirstChild(); first != nil {
		s.InsertBefore(csd, first)
	} else {
		s.AppendChild(csd)
	}

	return csd
}

// ColorMapOverride returns the color map override element.
func (s *Slide) ColorMapOverride() *ColorMapOverride {
	elem := s.GetElement(
		"clrMapOvr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if cmo, ok := elem.(*ColorMapOverride); ok {
		return cmo
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ColorMapOverride{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Transition returns the slide transition element.
func (s *Slide) Transition() *SlideTransition {
	elem := s.GetElement(
		"transition",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if tr, ok := elem.(*SlideTransition); ok {
		return tr
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTransition{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetTransition sets the slide transition.
func (s *Slide) SetTransition(
	tr *SlideTransition,
) {
	// Remove existing transition
	if existing := s.Transition(); existing != nil {
		s.RemoveChild(existing)
	}
	if tr != nil {
		// Insert after cSld and clrMapOvr
		if cmo := s.ColorMapOverride(); cmo != nil {
			s.InsertAfter(tr, cmo)
		} else if csd := s.CommonSlideData(); csd != nil {
			s.InsertAfter(tr, csd)
		} else {
			s.AppendChild(tr)
		}
	}
}

// Timing returns the slide timing element.
func (s *Slide) Timing() *SlideTiming {
	elem := s.GetElement(
		"timing",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*SlideTiming); ok {
		return t
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTiming{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// ShapeTree returns the shape tree from the common slide data.
func (s *Slide) ShapeTree() *ShapeTree {
	csd := s.CommonSlideData()
	if csd == nil {
		return nil
	}

	return csd.ShapeTree()
}

// GetOrCreateShapeTree returns the shape tree, creating necessary parent elements if needed.
func (s *Slide) GetOrCreateShapeTree() *ShapeTree {
	csd := s.GetOrCreateCommonSlideData()

	return csd.GetOrCreateShapeTree()
}

// AddShape adds a new shape to the slide.
func (s *Slide) AddShape() *Shape {
	st := s.GetOrCreateShapeTree()

	return st.AddShape()
}

// AddPicture adds a new picture to the slide.
func (s *Slide) AddPicture(
	relId string,
) *Picture {
	st := s.GetOrCreateShapeTree()

	return st.AddPicture(relId)
}

// AddVideo adds a new video to the slide.
func (s *Slide) AddVideo(
	videoRelId string,
	posterRelId string,
) *Picture {
	pic := s.AddPicture(posterRelId)
	pic.NonVisualPictureProperties().SetVideoFile(videoRelId)

	return pic
}

// AddAudio adds a new audio to the slide.
func (s *Slide) AddAudio(
	audioRelId string,
	posterRelId string,
) *Picture {
	pic := s.AddPicture(posterRelId)
	pic.NonVisualPictureProperties().SetAudioFile(audioRelId)

	return pic
}

// ===========================================================================
// NotesSlide (p:notesSld) - Notes slide
// ===========================================================================

// HeaderFooter returns the header/footer element.
func (ns *NotesSlide) HeaderFooter() *ExtHeaderFooter {
	elem := ns.GetElement(
		"hf",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if hf, ok := elem.(*ExtHeaderFooter); ok {
		return hf
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ExtHeaderFooter{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// ===========================================================================
// ColorMapOverride (p:clrMapOvr)
// ===========================================================================

// ColorMapOverride represents the color map override element (p:clrMapOvr).
type ColorMapOverride struct {
	*openxml.CompositeElementBase
}

// NewColorMapOverride creates a new ColorMapOverride element.
func NewColorMapOverride() *ColorMapOverride {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"clrMapOvr",
		PrefixP,
	)

	return &ColorMapOverride{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this ColorMapOverride element.
func (cmo *ColorMapOverride) Clone() openxml.Element {
	cloned := cmo.CompositeElementBase.Clone()

	return &ColorMapOverride{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideTransition (p:transition)
// ===========================================================================

// SlideTransition represents the slide transition element (p:transition).
type SlideTransition struct {
	*openxml.CompositeElementBase
}

// NewSlideTransition creates a new SlideTransition element.
func NewSlideTransition() *SlideTransition {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"transition",
		PrefixP,
	)

	return &SlideTransition{
		CompositeElementBase: elem,
	}
}

// Speed returns the transition speed.
func (tr *SlideTransition) Speed() TransitionSpeed {
	attr, found := tr.GetAttribute("spd", "")
	if !found {
		return TransitionSpeedFast
	}

	return TransitionSpeed(attr.Value())
}

// SetSpeed sets the transition speed.
func (tr *SlideTransition) SetSpeed(
	speed TransitionSpeed,
) {
	if speed == "" ||
		speed == TransitionSpeedFast {
		tr.RemoveAttribute("spd", "")

		return
	}
	tr.SetAttribute(openxml.NewAttribute(
		"",
		"spd",
		"",
		string(speed),
	))
}

// AdvanceOnClick returns whether the slide advances on click.
func (tr *SlideTransition) AdvanceOnClick() bool {
	attr, found := tr.GetAttribute("advClick", "")
	if !found {
		return true // default is true
	}

	return attr.Value() == "1" ||
		attr.Value() == attrValueTrue
}

// SetAdvanceOnClick sets whether the slide advances on click.
func (tr *SlideTransition) SetAdvanceOnClick(
	advance bool,
) {
	if advance {
		tr.RemoveAttribute("advClick", "")

		return
	}
	tr.SetAttribute(openxml.NewAttribute(
		"",
		"advClick",
		"",
		"0",
	))
}

// Clone creates a deep copy of this SlideTransition element.
func (tr *SlideTransition) Clone() openxml.Element {
	cloned := tr.CompositeElementBase.Clone()

	return &SlideTransition{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideTiming (p:timing)
// ===========================================================================

// SlideTiming represents the slide timing element (p:timing).
type SlideTiming struct {
	*openxml.CompositeElementBase
}

// NewSlideTiming creates a new SlideTiming element.
func NewSlideTiming() *SlideTiming {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"timing",
		PrefixP,
	)

	return &SlideTiming{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this SlideTiming element.
func (t *SlideTiming) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &SlideTiming{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// SlideLayout (p:sldLayout) - Slide layout for reusable slide structures
// ===========================================================================

// SlideLayout represents a slide layout element (p:sldLayout).
type SlideLayout struct {
	*openxml.PartRootElementBase
}

// NewSlideLayout creates a new SlideLayout element.
func NewSlideLayout() *SlideLayout {
	elem := openxml.NewPartRootElement(
		NamespacePresentationML,
		"sldLayout",
		PrefixP,
	)
	sl := &SlideLayout{PartRootElementBase: elem}

	// Add required common slide data
	sl.AppendChild(NewCommonSlideData())

	return sl
}

// CommonSlideData returns the common slide data element.
func (sl *SlideLayout) CommonSlideData() *CommonSlideData {
	elem := sl.GetElement(
		"cSld",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if csd, ok := elem.(*CommonSlideData); ok {
		return csd
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &CommonSlideData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCommonSlideData returns or creates the common slide data.
func (sl *SlideLayout) GetOrCreateCommonSlideData() *CommonSlideData {
	csd := sl.CommonSlideData()
	if csd != nil {
		return csd
	}
	csd = NewCommonSlideData()
	sl.AppendChild(csd)
	return csd
}

// Name returns the slide layout name.
func (sl *SlideLayout) Name() string {
	csd := sl.CommonSlideData()
	if csd == nil {
		return ""
	}

	return csd.Name()
}

// SetName sets the slide layout name.
func (sl *SlideLayout) SetName(name string) {
	sl.GetOrCreateCommonSlideData().SetName(name)
}

// AddPlaceholder adds a new placeholder shape to the layout.
func (sl *SlideLayout) AddPlaceholder(phType PlaceholderType, idx int) *Shape {
	csd := sl.GetOrCreateCommonSlideData()
	st := csd.GetOrCreateShapeTree()
	shape := st.AddShape()
	shape.NonVisualShapeProperties().SetPlaceholder(phType, idx)
	return shape
}

// ColorMapOverride returns the color map override element.
func (sl *SlideLayout) ColorMapOverride() *ColorMapOverride {
	elem := sl.GetElement(
		"clrMapOvr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if cmo, ok := elem.(*ColorMapOverride); ok {
		return cmo
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ColorMapOverride{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// Transition returns the slide transition.
func (sl *SlideLayout) Transition() *SlideTransition {
	elem := sl.GetElement(
		"transition",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if tr, ok := elem.(*SlideTransition); ok {
		return tr
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTransition{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// Timing returns the slide timing.
func (sl *SlideLayout) Timing() *SlideTiming {
	elem := sl.GetElement(
		"timing",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*SlideTiming); ok {
		return t
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTiming{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// HeaderFooter returns the header/footer element.
func (sl *SlideLayout) HeaderFooter() *ExtHeaderFooter {
	elem := sl.GetElement(
		"hf",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if hf, ok := elem.(*ExtHeaderFooter); ok {
		return hf
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ExtHeaderFooter{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// ExtensionList returns the extension list.
func (sl *SlideLayout) ExtensionList() *ExtensionListModify {
	elem := sl.GetElement(
		"extLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if el, ok := elem.(*ExtensionListModify); ok {
		return el
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ExtensionListModify{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// TextStyles returns the text styles element.
func (sl *SlideLayout) TextStyles() *TextStyles {
	elem := sl.GetElement(
		"txStyles",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ts, ok := elem.(*TextStyles); ok {
		return ts
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &TextStyles{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// Clone creates a deep copy of this SlideLayout element.
func (sl *SlideLayout) Clone() openxml.Element {
	cloned := sl.PartRootElementBase.Clone()

	return &SlideLayout{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// ===========================================================================
// SlideMaster (p:sldMaster) - Master slide for consistent styling
// ===========================================================================

// SlideMaster represents a slide master element (p:sldMaster).
type SlideMaster struct {
	*openxml.PartRootElementBase
}

// NewSlideMaster creates a new SlideMaster element.
func NewSlideMaster() *SlideMaster {
	elem := openxml.NewPartRootElement(
		NamespacePresentationML,
		"sldMaster",
		PrefixP,
	)
	sm := &SlideMaster{PartRootElementBase: elem}

	// Add required common slide data
	sm.AppendChild(NewCommonSlideData())

	return sm
}

// CommonSlideData returns the common slide data element.
func (sm *SlideMaster) CommonSlideData() *CommonSlideData {
	elem := sm.GetElement(
		"cSld",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if csd, ok := elem.(*CommonSlideData); ok {
		return csd
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &CommonSlideData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Clone creates a deep copy of this SlideMaster element.
func (sm *SlideMaster) Clone() openxml.Element {
	cloned := sm.PartRootElementBase.Clone()

	return &SlideMaster{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// ColorMap returns the color map element.
func (sm *SlideMaster) ColorMap() *PresentationColorMap {
	elem := sm.GetElement(
		"clrMap",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if cm, ok := elem.(*PresentationColorMap); ok {
		return cm
	}
	// Wrap generic element
	if comp := wrapCompositeElement(elem); comp != nil {
		// Use the existing element's data but wrapped in our type
		// Note: We might need a better way to wrap if strict validation is needed
		cm := NewPresentationColorMap()
		cm.CompositeElementBase = comp
		return cm
	}
	return nil
}

// SlideLayoutIdList returns the slide layout ID list.
func (sm *SlideMaster) SlideLayoutIdList() *SlideLayoutIdList {
	elem := sm.GetElement(
		"sldLayoutIdLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if slil, ok := elem.(*SlideLayoutIdList); ok {
		return slil
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideLayoutIdList{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// Transition returns the slide transition.
func (sm *SlideMaster) Transition() *SlideTransition {
	elem := sm.GetElement(
		"transition",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if tr, ok := elem.(*SlideTransition); ok {
		return tr
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTransition{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// Timing returns the slide timing.
func (sm *SlideMaster) Timing() *SlideTiming {
	elem := sm.GetElement(
		"timing",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*SlideTiming); ok {
		return t
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &SlideTiming{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// HeaderFooter returns the header/footer element.
func (sm *SlideMaster) HeaderFooter() *ExtHeaderFooter {
	elem := sm.GetElement(
		"hf",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if hf, ok := elem.(*ExtHeaderFooter); ok {
		return hf
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ExtHeaderFooter{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// TextStyles returns the text styles element.
func (sm *SlideMaster) TextStyles() *TextStyles {
	elem := sm.GetElement(
		"txStyles",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if ts, ok := elem.(*TextStyles); ok {
		return ts
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &TextStyles{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// ExtensionList returns the extension list.
func (sm *SlideMaster) ExtensionList() *ExtensionListModify {
	elem := sm.GetElement(
		"extLst",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if el, ok := elem.(*ExtensionListModify); ok {
		return el
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ExtensionListModify{
			CompositeElementBase: comp,
		}
	}
	return nil
}

// GetOrCreateSlideLayoutIdList returns the slide layout ID list, creating if needed.
func (sm *SlideMaster) GetOrCreateSlideLayoutIdList() *SlideLayoutIdList {
	slil := sm.SlideLayoutIdList()
	if slil != nil {
		return slil
	}
	slil = NewSlideLayoutIdList()
	// Insert after clrMap
	if cm := sm.ColorMap(); cm != nil {
		sm.InsertAfter(slil, cm)
	} else if csd := sm.CommonSlideData(); csd != nil {
		sm.InsertAfter(slil, csd)
	} else {
		sm.AppendChild(slil)
	}

	return slil
}

// AddSlideLayoutId adds a slide layout ID entry.
func (sm *SlideMaster) AddSlideLayoutId(
	id uint32,
	relId string,
) *SlideLayoutId {
	return sm.GetOrCreateSlideLayoutIdList().AddSlideLayoutId(id, relId)
}

// ===========================================================================
// SlideLayoutIdList (p:sldLayoutIdLst)
// ===========================================================================

// AddSlideLayoutId adds a slide layout ID entry.
func (slil *SlideLayoutIdList) AddSlideLayoutId(
	id uint32,
	relId string,
) *SlideLayoutId {
	slid := NewSlideLayoutId()
	slid.SetId(id)
	slid.SetRelId(relId)
	slil.AppendChild(slid)

	return slid
}

// ===========================================================================
// SlideLayoutId (p:sldLayoutId)
// ===========================================================================

// SetId sets the slide layout ID.
func (sid *SlideLayoutId) SetId(id uint32) {
	sid.Id = types.NewUInt32Value(id)
}

// SetRelId sets the relationship ID.
func (sid *SlideLayoutId) SetRelId(relId string) {
	sid.RelationshipId = types.NewStringValue(relId)
}

// ===========================================================================
// HeaderFooter (p:hf)
// ===========================================================================

// SetSlideNumber sets whether to show the slide number.
func (hf *ExtHeaderFooter) SetSlideNumber(show bool) {
	hf.SlideNumber = types.NewBooleanValue(show)
}

// SetHeader sets whether to show the header.
func (hf *ExtHeaderFooter) SetHeader(show bool) {
	hf.Header = types.NewBooleanValue(show)
}

// SetFooter sets whether to show the footer.
func (hf *ExtHeaderFooter) SetFooter(show bool) {
	hf.Footer = types.NewBooleanValue(show)
}

// SetDateTime sets whether to show the date and time.
func (hf *ExtHeaderFooter) SetDateTime(show bool) {
	hf.DateTime = types.NewBooleanValue(show)
}

// ===========================================================================
// PresentationColorMap (p:clrMap)
// ===========================================================================

// PresentationColorMap represents the color map element (p:clrMap) in PresentationML.
// Note: We use this wrapper because the generated ColorMap is for DrawingML (pic:clrMap).
type PresentationColorMap struct {
	*ColorMappingType
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main clrMap"`
}

// NewPresentationColorMap creates a new PresentationColorMap.
func NewPresentationColorMap() *PresentationColorMap {
	cmt := NewColorMappingType()
	// Override the underlying element name/namespace to match p:clrMap
	cmt.CompositeElementBase = openxml.NewCompositeElement(
		NamespacePresentationML,
		"clrMap",
		PrefixP,
	)

	return &PresentationColorMap{
		ColorMappingType: cmt,
	}
}

// Clone creates a deep copy of this PresentationColorMap.
func (m *PresentationColorMap) Clone() openxml.Element {
	cloned := m.ColorMappingType.Clone()
	return &PresentationColorMap{
		ColorMappingType: cloned.(*ColorMappingType),
	}
}
