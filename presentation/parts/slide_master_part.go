//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
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
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sldMaster xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:bg>
      <p:bgRef idx="1001">
        <a:schemeClr val="bg1"/>
      </p:bgRef>
    </p:bg>
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr/>
    </p:spTree>
  </p:cSld>
  <p:clrMap bg1="lt1" tx1="dk1" bg2="lt2" tx2="dk2" accent1="accent1" accent2="accent2" accent3="accent3" accent4="accent4" accent5="accent5" accent6="accent6" hlink="hlink" folHlink="folHlink"/>
  <p:sldLayoutIdLst/>
</p:sldMaster>`
	smp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlideMasterPart) FixedContentType() string {
	return ContentTypeSlideMaster
}

// SlideMaster returns the root SlideMaster element.
// TODO: Return a proper SlideMaster element type when elements are implemented.
func (smp *SlideMasterPart) SlideMaster() openxml.PartRootElement {
	return smp.RootElement()
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
