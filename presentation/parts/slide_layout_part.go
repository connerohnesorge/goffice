//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// Counter for generating unique slide layout filenames.
var slideLayoutUniqueCounter uint64

// SlideLayoutPart represents a slide layout part (ppt/slideLayouts/slideLayout1.xml, etc.).
type SlideLayoutPart struct {
	*openxml.OpenXmlPartData
}

// newSlideLayoutPart creates a new slide layout part.
func newSlideLayoutPart(
	slideMasterPart *SlideMasterPart,
	uri string,
) (*SlideLayoutPart, error) {
	packPart, relID, err := slideMasterPart.addChildPart(
		uri,
		ContentTypeSlideLayout,
		RelationshipTypeSlideLayout,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		slideMasterPart,
	)
	partData.SetRelationshipID(relID)

	slp := &SlideLayoutPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide layout content
	slp.initializeContent()

	// Add to slide master part's child parts
	if err := slideMasterPart.AddPart(slp, relID); err != nil {
		return nil, err
	}

	return slp, nil
}

// NewSlideLayoutPart creates a new slide layout part from a presentation part.
// This is used when creating layouts directly from the presentation.
func NewSlideLayoutPart(
	presPart *PresentationPart,
) (*SlideLayoutPart, error) {
	num := atomic.AddUint64(
		&slideLayoutUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/slideLayouts/slideLayout%d.xml",
		num,
	)

	packPart, relID, err := presPart.addChildPart(
		uri,
		ContentTypeSlideLayout,
		RelationshipTypeSlideLayout,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		presPart,
	)
	partData.SetRelationshipID(relID)

	slp := &SlideLayoutPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide layout content
	slp.initializeContent()

	// Add to presentation part's child parts
	if err := presPart.AddPart(slp, relID); err != nil {
		return nil, err
	}

	return slp, nil
}

// initializeContent sets up minimal slide layout content.
func (slp *SlideLayoutPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sldLayout xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" type="blank" preserve="1">
  <p:cSld name="Blank">
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr/>
    </p:spTree>
  </p:cSld>
  <p:clrMapOvr>
    <a:masterClrMapping/>
  </p:clrMapOvr>
</p:sldLayout>`
	slp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlideLayoutPart) FixedContentType() string {
	return ContentTypeSlideLayout
}

// SlideLayout returns the root SlideLayout element.
// TODO: Return a proper SlideLayout element type when elements are implemented.
func (slp *SlideLayoutPart) SlideLayout() openxml.PartRootElement {
	return slp.RootElement()
}

// GetStream returns a reader for the part content.
func (slp *SlideLayoutPart) GetStream() io.Reader {
	return slp.OpenXmlPartData.GetStream()
}

// Ensure SlideLayoutPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlideLayoutPart)(
	nil,
)

// SlideLayoutPartFactory creates a SlideLayoutPart from a URI and container.
func SlideLayoutPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlideLayout,
		packPart,
		container,
	)

	return &SlideLayoutPart{
		OpenXmlPartData: partData,
	}
}

// Register the SlideLayoutPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlideLayout,
			RelationshipType:   RelationshipTypeSlideLayout,
			Factory:            SlideLayoutPartFactory,
			DefaultURI:         "/ppt/slideLayouts/slideLayout1.xml",
			IsFixedContentType: true,
		},
	)
}
