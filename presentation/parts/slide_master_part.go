//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// Counter for generating unique slide layout filenames per master.
var masterLayoutCounter uint64

// SlideMasterPart represents a slide master part (ppt/slideMasters/slideMaster1.xml, etc.).
type SlideMasterPart struct {
	*openxml.OpenXmlPartData
}

// newSlideMasterPart creates a new slide master part.
func newSlideMasterPart(
	presentationPart *PresentationPart,
	uri string,
) (*SlideMasterPart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeSlideMaster,
		RelationshipTypeSlideMaster,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideMaster,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	smp := &SlideMasterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide master content
	smp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(smp, relID); err != nil {
		return nil, err
	}

	return smp, nil
}

// initializeContent sets up minimal slide master content.
func (smp *SlideMasterPart) initializeContent() {
	sm := elements.NewSlideMaster()
	smp.SetRootElement(sm)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlideMasterPart) FixedContentType() string {
	return ContentTypeSlideMaster
}

// SlideMaster returns the root SlideMaster element.
func (smp *SlideMasterPart) SlideMaster() *elements.SlideMaster {
	root := smp.RootElement()
	if root == nil {
		return nil
	}
	if sm, ok := root.(*elements.SlideMaster); ok {
		return sm
	}

	return nil
}

// AddSlideLayoutPart adds a new slide layout part to this slide master.
func (smp *SlideMasterPart) AddSlideLayoutPart() (*SlideLayoutPart, error) {
	num := atomic.AddUint64(
		&masterLayoutCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/slideLayouts/slideLayout%d.xml",
		num,
	)

	return newSlideLayoutPart(smp, uri)
}

// SlideLayoutParts returns all slide layout parts.
func (smp *SlideMasterPart) SlideLayoutParts() []*SlideLayoutPart {
	var layouts []*SlideLayoutPart
	for part := range smp.Parts() {
		if slp, ok := part.(*SlideLayoutPart); ok {
			layouts = append(layouts, slp)
		}
	}

	return layouts
}

// AddThemePart adds a theme part to this slide master.
func (smp *SlideMasterPart) AddThemePart() (*ThemePart, error) {
	return newThemePartForMaster(smp)
}

// ThemePart returns the theme part if present.
func (smp *SlideMasterPart) ThemePart() *ThemePart {
	for part := range smp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (smp *SlideMasterPart) GetStream() io.Reader {
	return smp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (smp *SlideMasterPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		smp,
		uri,
		contentType,
		relType,
	)
}

// Ensure SlideMasterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlideMasterPart)(
	nil,
)

// SlideMasterPartFactory creates a SlideMasterPart from a URI and container.
func SlideMasterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideMaster,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSlideMaster()
		},
	)

	return &SlideMasterPart{
		OpenXmlPartData: partData,
	}
}

// Register the SlideMasterPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlideMaster,
			RelationshipType:   RelationshipTypeSlideMaster,
			Factory:            SlideMasterPartFactory,
			DefaultURI:         "/ppt/slideMasters/slideMaster1.xml",
			IsFixedContentType: true,
		},
	)
}
