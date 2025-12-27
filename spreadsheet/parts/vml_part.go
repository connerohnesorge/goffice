//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// VmlDrawingPart represents a VML drawing part (xl/drawings/vmlDrawing1.vml, etc.).
// VML (Vector Markup Language) is used for legacy drawing features like
// comment callout shapes and form controls.
type VmlDrawingPart struct {
	*openxml.OpenXmlPartData
}

// newVmlDrawingPart creates a new VML drawing part.
func newVmlDrawingPart(
	worksheetPart *WorksheetPart,
	uri string,
) (*VmlDrawingPart, error) {
	packPart, relID, err := worksheetPart.addChildPart(
		uri,
		ContentTypeVml,
		RelationshipTypeVml,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeVml,
		packPart,
		worksheetPart,
	)
	partData.SetRelationshipID(relID)

	vp := &VmlDrawingPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal VML content
	vp.initializeContent()

	// Add to worksheet part's child parts
	if err := worksheetPart.AddPart(vp, relID); err != nil {
		return nil, err
	}

	return vp, nil
}

// initializeContent sets up minimal VML drawing content.
func (vp *VmlDrawingPart) initializeContent() {
	content := `<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel">
 <o:shapelayout v:ext="edit">
  <o:idmap v:ext="edit" data="1"/>
 </o:shapelayout>
 <v:shapetype id="_x0000_t202" coordsize="21600,21600" o:spt="202" path="m,l,21600r21600,l21600,xe">
  <v:stroke joinstyle="miter"/>
  <v:path gradientshapeok="t" o:connecttype="rect"/>
 </v:shapetype>
</xml>`
	vp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*VmlDrawingPart) FixedContentType() string {
	return ContentTypeVml
}

// VmlDrawing returns the raw VML content.
// Note: VML is not a standard OpenXML element, so we return the raw data.
func (vp *VmlDrawingPart) VmlDrawing() []byte {
	return vp.GetData()
}

// SetVmlDrawing sets the VML content.
func (vp *VmlDrawingPart) SetVmlDrawing(
	data []byte,
) {
	vp.SetData(data)
}

// GetStream returns a reader for the part content.
func (vp *VmlDrawingPart) GetStream() io.Reader {
	return vp.OpenXmlPartData.GetStream()
}

// Ensure VmlDrawingPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*VmlDrawingPart)(nil)

// VmlDrawingPartFactory creates a VmlDrawingPart from a URI and container.
func VmlDrawingPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeVml,
		packPart,
		container,
	)

	return &VmlDrawingPart{
		OpenXmlPartData: partData,
	}
}

// Register the VmlDrawingPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeVml,
			RelationshipType:   RelationshipTypeVml,
			Factory:            VmlDrawingPartFactory,
			DefaultURI:         "/xl/drawings/vmlDrawing1.vml",
			IsFixedContentType: true,
		},
	)
}
