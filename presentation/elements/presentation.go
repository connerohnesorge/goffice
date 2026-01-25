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
















