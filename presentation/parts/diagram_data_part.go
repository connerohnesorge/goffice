//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// DiagramDataPart represents a diagram data part (ppt/diagrams/data1.xml, etc.).
// This part contains the SmartArt diagram data and node structure.
type DiagramDataPart struct {
	*openxml.OpenXmlPartData
}

// newDiagramDataPart creates a new diagram data part.
//
//nolint:unused // Will be used when diagram creation API is implemented
func newDiagramDataPart(
	parent openxml.OpenXmlPartContainer,
	uri string,
) (*DiagramDataPart, error) {
	parentPart, ok := parent.(partWithPackage)
	if !ok {
		return nil, ErrNilPackage
	}

	packPart, relID, err := addChildPart(
		parentPart,
		uri,
		ContentTypeDiagramData,
		RelationshipTypeDiagramData,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramData,
		packPart,
		parent,
	)
	partData.SetRelationshipID(relID)

	ddp := &DiagramDataPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal diagram data content
	ddp.initializeContent()

	// Add to parent's child parts
	if err := parent.AddPart(ddp, relID); err != nil {
		return nil, err
	}

	return ddp, nil
}

// initializeContent sets up minimal diagram data content.
//
//nolint:unused // Will be used when diagram creation API is implemented
func (ddp *DiagramDataPart) initializeContent() {
	dataModel := diagram.NewDataModelRoot()
	ddp.SetRootElement(dataModel)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*DiagramDataPart) FixedContentType() string {
	return ContentTypeDiagramData
}

// DataModel returns the root DataModelRoot element.
func (ddp *DiagramDataPart) DataModel() *diagram.DataModelRoot {
	root := ddp.RootElement()
	if root == nil {
		return nil
	}
	if dataModel, ok := root.(*diagram.DataModelRoot); ok {
		return dataModel
	}

	return nil
}

// GetStream returns a reader for the part content.
func (ddp *DiagramDataPart) GetStream() io.Reader {
	return ddp.OpenXmlPartData.GetStream()
}

// Ensure DiagramDataPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*DiagramDataPart)(
	nil,
)

// DiagramDataPartFactory creates a DiagramDataPart from a URI and container.
func DiagramDataPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramData,
		packPart,
		container,
	)

	return &DiagramDataPart{
		OpenXmlPartData: partData,
	}
}

// Register the DiagramDataPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeDiagramData,
			RelationshipType:   RelationshipTypeDiagramData,
			Factory:            DiagramDataPartFactory,
			DefaultURI:         "/ppt/diagrams/data1.xml",
			IsFixedContentType: true,
		},
	)
}
