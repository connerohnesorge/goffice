//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// HandoutMasterPart represents a handout master part (ppt/handoutMasters/handoutMaster1.xml).
type HandoutMasterPart struct {
	*openxml.OpenXmlPartData
}

// newHandoutMasterPart creates a new handout master part.
func newHandoutMasterPart(
	presentationPart *PresentationPart,
	uri string,
) (*HandoutMasterPart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeHandoutMaster,
		RelationshipTypeHandoutMaster,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHandoutMaster,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	hmp := &HandoutMasterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal handout master content
	hmp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(hmp, relID); err != nil {
		return nil, err
	}

	return hmp, nil
}

// initializeContent sets up minimal handout master content.
func (hmp *HandoutMasterPart) initializeContent() {
	hm := elements.NewHandoutMaster()
	hmp.SetRootElement(hm)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*HandoutMasterPart) FixedContentType() string {
	return ContentTypeHandoutMaster
}

// HandoutMaster returns the root HandoutMaster element.
func (hmp *HandoutMasterPart) HandoutMaster() *elements.HandoutMaster {
	root := hmp.RootElement()
	if root == nil {
		return nil
	}
	if hm, ok := root.(*elements.HandoutMaster); ok {
		return hm
	}

	return nil
}

// AddThemePart adds a theme part to this handout master.
func (hmp *HandoutMasterPart) AddThemePart() (*ThemePart, error) {
	return newThemePartForHandoutMaster(hmp)
}

// ThemePart returns the theme part if present.
func (hmp *HandoutMasterPart) ThemePart() *ThemePart {
	for part := range hmp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (hmp *HandoutMasterPart) GetStream() io.Reader {
	return hmp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (hmp *HandoutMasterPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		hmp,
		uri,
		contentType,
		relType,
	)
}

// Ensure HandoutMasterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*HandoutMasterPart)(
	nil,
)

// HandoutMasterPartFactory creates a HandoutMasterPart from a URI and container.
func HandoutMasterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeHandoutMaster,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewHandoutMaster()
		},
	)

	return &HandoutMasterPart{
		OpenXmlPartData: partData,
	}
}

// Register the HandoutMasterPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeHandoutMaster,
			RelationshipType:   RelationshipTypeHandoutMaster,
			Factory:            HandoutMasterPartFactory,
			DefaultURI:         "/ppt/handoutMasters/handoutMaster1.xml",
			IsFixedContentType: true,
		},
	)
}
