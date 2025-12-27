//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// NotesMasterPart represents a notes master part (ppt/notesMasters/notesMaster1.xml).
type NotesMasterPart struct {
	*openxml.OpenXmlPartData
}

// newNotesMasterPart creates a new notes master part.
func newNotesMasterPart(
	presentationPart *PresentationPart,
	uri string,
) (*NotesMasterPart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeNotesMaster,
		RelationshipTypeNotesMaster,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesMaster,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	nmp := &NotesMasterPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal notes master content
	nmp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(nmp, relID); err != nil {
		return nil, err
	}

	return nmp, nil
}

// initializeContent sets up minimal notes master content.
func (nmp *NotesMasterPart) initializeContent() {
	nm := elements.NewNotesMaster()
	nmp.SetRootElement(nm)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*NotesMasterPart) FixedContentType() string {
	return ContentTypeNotesMaster
}

// NotesMaster returns the root NotesMaster element.
func (nmp *NotesMasterPart) NotesMaster() *elements.NotesMaster {
	root := nmp.RootElement()
	if root == nil {
		return nil
	}
	if nm, ok := root.(*elements.NotesMaster); ok {
		return nm
	}

	return nil
}

// AddThemePart adds a theme part to this notes master.
func (nmp *NotesMasterPart) AddThemePart() (*ThemePart, error) {
	return newThemePartForNotesMaster(nmp)
}

// ThemePart returns the theme part if present.
func (nmp *NotesMasterPart) ThemePart() *ThemePart {
	for part := range nmp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (nmp *NotesMasterPart) GetStream() io.Reader {
	return nmp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (nmp *NotesMasterPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		nmp,
		uri,
		contentType,
		relType,
	)
}

// Ensure NotesMasterPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*NotesMasterPart)(
	nil,
)

// NotesMasterPartFactory creates a NotesMasterPart from a URI and container.
func NotesMasterPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesMaster,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewNotesMaster()
		},
	)

	return &NotesMasterPart{
		OpenXmlPartData: partData,
	}
}

// Register the NotesMasterPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeNotesMaster,
			RelationshipType:   RelationshipTypeNotesMaster,
			Factory:            NotesMasterPartFactory,
			DefaultURI:         "/ppt/notesMasters/notesMaster1.xml",
			IsFixedContentType: true,
		},
	)
}
