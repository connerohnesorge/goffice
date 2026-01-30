//nolint:revive // This file contains presentation element implementation.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// NewSlideMasterIdList creates a new slide master ID list element.
func NewSlideMasterIdList() openxml.Element {
	return openxml.NewCompositeElement(
		NamespacePresentationML,
		"sldMasterIdLst",
		PrefixP,
	)
}

// NewSlideIdList creates a new slide ID list element.
func NewSlideIdList() openxml.Element {
	return openxml.NewCompositeElement(
		NamespacePresentationML,
		"sldIdLst",
		PrefixP,
	)
}

// NewSlideSize creates a new slide size element.
func NewSlideSize(cx, cy int, sizeType SlideSizeType) openxml.Element {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"sldSz",
		PrefixP,
	)
	elem.SetAttribute(openxml.NewAttribute("", "cx", "", strconv.Itoa(cx)))
	elem.SetAttribute(openxml.NewAttribute("", "cy", "", strconv.Itoa(cy)))
	if sizeType != "" {
		elem.SetAttribute(openxml.NewAttribute("", "type", "", string(sizeType)))
	}

	return elem
}

// NewNotesSize creates a new notes size element.
func NewNotesSize(cx, cy int) openxml.Element {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"notesSz",
		PrefixP,
	)
	elem.SetAttribute(openxml.NewAttribute("", "cx", "", strconv.Itoa(cx)))
	elem.SetAttribute(openxml.NewAttribute("", "cy", "", strconv.Itoa(cy)))

	return elem
}

// NotesSizeWrapper wraps a notes size element.
type NotesSizeWrapper struct {
	elem openxml.Element
}

// WrapNotesSize wraps an existing notes size element.
func WrapNotesSize(elem openxml.Element) *NotesSizeWrapper {
	return &NotesSizeWrapper{elem: elem}
}

// Cx returns the width in EMUs.
func (n *NotesSizeWrapper) Cx() int {
	attr, found := n.elem.GetAttribute("cx", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Cy returns the height in EMUs.
func (n *NotesSizeWrapper) Cy() int {
	attr, found := n.elem.GetAttribute("cy", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// Presentation represents the root presentation element (p:presentation).
type Presentation struct {
	*openxml.PartRootElementBase
}

// NewPresentation creates a new Presentation element.
func NewPresentation() *Presentation {
	elem := openxml.NewPartRootElement(
		NamespacePresentationML,
		"presentation",
		PrefixP,
	)
	p := &Presentation{PartRootElementBase: elem}

	// Add required child elements
	p.AppendChild(NewSlideMasterIdList())
	p.AppendChild(NewSlideIdList())
	p.AppendChild(
		NewSlideSize(
			Screen4x3Width,
			Screen4x3Height,
			SlideSizeScreen4x3,
		),
	)
	p.AppendChild(
		NewNotesSize(
			Screen4x3Height,
			Screen4x3Width,
		),
	)

	return p
}

// FirstSlideNumber returns the starting slide number.
func (p *Presentation) FirstSlideNumber() uint32 {
	attr, found := p.GetAttribute("firstSldNum", "")
	if !found {
		return 1 // Default is 1
	}
	val, _ := strconv.ParseUint(attr.Value(), 10, 32)

	return uint32(val)
}

// SetFirstSlideNumber sets the starting slide number.
func (p *Presentation) SetFirstSlideNumber(num uint32) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"firstSldNum",
			"",
			strconv.FormatUint(uint64(num), 10),
		),
	)
}

// AddSlideId adds a slide ID to the presentation.
func (p *Presentation) AddSlideId(id uint32, rId string) {
	slideIdList := p.GetElement("sldIdLst", NamespacePresentationML)
	if slideIdList == nil {
		return
	}
	slideId := openxml.NewCompositeElement(NamespacePresentationML, "sldId", PrefixP)
	slideId.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
	slideId.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))
	if comp, ok := slideIdList.(openxml.CompositeElement); ok {
		comp.AppendChild(slideId)
	}
}

// AddSlideMasterId adds a slide master ID to the presentation.
func (p *Presentation) AddSlideMasterId(id uint32, rId string) {
	slideMasterIdList := p.GetElement("sldMasterIdLst", NamespacePresentationML)
	if slideMasterIdList == nil {
		return
	}
	slideMasterId := openxml.NewCompositeElement(NamespacePresentationML, "sldMasterId", PrefixP)
	slideMasterId.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
	slideMasterId.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))
	if comp, ok := slideMasterIdList.(openxml.CompositeElement); ok {
		comp.AppendChild(slideMasterId)
	}
}

// SlideId represents a slide ID element wrapper.
type SlideId struct {
	elem openxml.Element
}

// NewSlideId creates a new slide ID element.
func NewSlideId(id uint32, rId string) *SlideId {
	elem := openxml.NewCompositeElement(NamespacePresentationML, "sldId", PrefixP)
	elem.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
	elem.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))

	return &SlideId{elem: elem}
}

// Id returns the slide ID.
func (s *SlideId) Id() uint32 {
	attr, found := s.elem.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(attr.Value(), 10, 32)

	return uint32(val)
}

// SetId sets the slide ID.
func (s *SlideId) SetId(id uint32) {
	s.elem.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
}

// RelId returns the relationship ID.
func (s *SlideId) RelId() string {
	attr, found := s.elem.GetAttribute("id", NamespaceRelationships)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelId sets the relationship ID.
func (s *SlideId) SetRelId(rId string) {
	s.elem.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))
}

// SlideIdListWrapper wraps a slide ID list element.
type SlideIdListWrapper struct {
	elem openxml.Element
}

// WrapSlideIdList wraps an existing slide ID list element.
func WrapSlideIdList(elem openxml.Element) *SlideIdListWrapper {
	return &SlideIdListWrapper{elem: elem}
}

// Count returns the number of slide IDs.
func (s *SlideIdListWrapper) Count() int {
	comp, ok := s.elem.(openxml.CompositeElement)
	if !ok {
		return 0
	}
	count := 0
	for range comp.Children() {
		count++
	}

	return count
}

// AddSlideId adds a slide ID to the list.
func (s *SlideIdListWrapper) AddSlideId(id uint32, rId string) {
	slideId := openxml.NewCompositeElement(NamespacePresentationML, "sldId", PrefixP)
	slideId.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
	slideId.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))
	if comp, ok := s.elem.(openxml.CompositeElement); ok {
		comp.AppendChild(slideId)
	}
}

// SlideIds returns all slide ID elements.
func (s *SlideIdListWrapper) SlideIds() []openxml.Element {
	comp, ok := s.elem.(openxml.CompositeElement)
	if !ok {
		return nil
	}
	var result []openxml.Element
	for child := range comp.Children() {
		result = append(result, child)
	}

	return result
}

// SlideMasterIdListWrapper wraps a slide master ID list element.
type SlideMasterIdListWrapper struct {
	elem openxml.Element
}

// WrapSlideMasterIdList wraps an existing slide master ID list element.
func WrapSlideMasterIdList(elem openxml.Element) *SlideMasterIdListWrapper {
	return &SlideMasterIdListWrapper{elem: elem}
}

// AddSlideMasterId adds a slide master ID to the list.
func (s *SlideMasterIdListWrapper) AddSlideMasterId(id uint32, rId string) {
	slideMasterId := openxml.NewCompositeElement(NamespacePresentationML, "sldMasterId", PrefixP)
	slideMasterId.SetAttribute(openxml.NewAttribute("", "id", "", strconv.FormatUint(uint64(id), 10)))
	slideMasterId.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "id", "r", rId))
	if comp, ok := s.elem.(openxml.CompositeElement); ok {
		comp.AppendChild(slideMasterId)
	}
}

// SlideMasterIds returns all slide master ID elements.
func (s *SlideMasterIdListWrapper) SlideMasterIds() []openxml.Element {
	comp, ok := s.elem.(openxml.CompositeElement)
	if !ok {
		return nil
	}
	var result []openxml.Element
	for child := range comp.Children() {
		result = append(result, child)
	}

	return result
}
