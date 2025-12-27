//nolint:revive // This file contains slide element implementation.
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
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

// Clone creates a deep copy of this Slide element.
func (s *Slide) Clone() openxml.Element {
	cloned := s.PartRootElementBase.Clone()

	return &Slide{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
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
		attr.Value() == "true"
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
