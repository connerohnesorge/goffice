//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// VbaProjectPart represents a VBA project part (xl/vbaProject.bin).
// This part contains the VBA macros for macro-enabled workbooks.
type VbaProjectPart struct {
	*openxml.OpenXmlPartData
}

// newVbaProjectPart creates a new VBA project part.
func newVbaProjectPart(
	workbookPart *WorkbookPart,
) (*VbaProjectPart, error) {
	uri := "/xl/vbaProject.bin"

	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeVbaProject,
		RelationshipTypeVbaProject,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeVbaProject,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	vp := &VbaProjectPart{
		OpenXmlPartData: partData,
	}

	// Note: VBA projects are binary and cannot be easily initialized
	// The caller should use SetVbaData to set the binary content

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(vp, relID); err != nil {
		return nil, err
	}

	return vp, nil
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*VbaProjectPart) FixedContentType() string {
	return ContentTypeVbaProject
}

// SetVbaData sets the VBA project binary data.
func (vp *VbaProjectPart) SetVbaData(
	data []byte,
) {
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
func VbaProjectPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeVbaProject,
		packPart,
		container,
	)

	return &VbaProjectPart{
		OpenXmlPartData: partData,
	}
}

// Register the VbaProjectPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeVbaProject,
			RelationshipType:   RelationshipTypeVbaProject,
			Factory:            VbaProjectPartFactory,
			DefaultURI:         "/xl/vbaProject.bin",
			IsFixedContentType: true,
		},
	)
}
