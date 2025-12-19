package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// VbaProjectPart represents a VBA project part (word/vbaProject.bin).
// This part contains the VBA macros for macro-enabled documents.
type VbaProjectPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for VBA projects.
const (
	ContentTypeVbaProject      = "application/vnd.ms-office.vbaProject"
	RelationshipTypeVbaProject = "http://schemas.microsoft.com/office/2006/relationships/vbaProject"
)

// newVbaProjectPart creates a new VBA project part.
func newVbaProjectPart(mainPart *MainPart) (*VbaProjectPart, error) {
	uri := "/word/vbaProject.bin"

	packPart, relID, err := mainPart.addChildPart(uri, ContentTypeVbaProject, RelationshipTypeVbaProject)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeVbaProject, packPart, mainPart)
	partData.SetRelationshipID(relID)

	vp := &VbaProjectPart{
		OpenXmlPartData: partData,
	}

	// Note: VBA projects are binary and cannot be easily initialized
	// The caller should use SetVbaData to set the binary content

	// Add to main part's child parts
	if err := mainPart.AddPart(vp, relID); err != nil {
		return nil, err
	}

	return vp, nil
}

// FixedContentType returns the content type for this part.
func (vp *VbaProjectPart) FixedContentType() string {
	return ContentTypeVbaProject
}

// SetVbaData sets the VBA project binary data.
func (vp *VbaProjectPart) SetVbaData(data []byte) {
	vp.SetData(data)
}

// GetVbaData returns the VBA project binary data.
func (vp *VbaProjectPart) GetVbaData() []byte {
	return vp.GetData()
}

// GetStream returns a reader for the part content.
func (vp *VbaProjectPart) GetStream() io.Reader {
	return vp.OpenXmlPartData.GetStream()
}

// Ensure VbaProjectPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*VbaProjectPart)(nil)

// VbaProjectPartFactory creates a VbaProjectPart from a URI and container.
func VbaProjectPartFactory(uri string, container openxml.OpenXmlPartContainer) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeVbaProject, packPart, container)
	return &VbaProjectPart{
		OpenXmlPartData: partData,
	}
}

// Register the VbaProjectPart type.
func init() {
	openxml.RegisterPartType(&openxml.PartTypeInfo{
		ContentType:        ContentTypeVbaProject,
		RelationshipType:   RelationshipTypeVbaProject,
		Factory:            VbaProjectPartFactory,
		DefaultURI:         "/word/vbaProject.bin",
		IsFixedContentType: true,
	})
}
