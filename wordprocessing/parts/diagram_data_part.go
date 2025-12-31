//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// Content type for diagram data.
const (
	ContentTypeDiagramData      = "application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml"
	ContentTypeDiagramLayoutDef = "application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml"
	ContentTypeDiagramStyle     = "application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml"
	ContentTypeDiagramColors    = "application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml"
)

// Relationship types for diagrams.
const (
	RelationshipTypeDiagramData       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramData"
	RelationshipTypeDiagramLayout     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramLayout"
	RelationshipTypeDiagramStyle      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramStyle"
	RelationshipTypeDiagramColors     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramColors"
	RelationshipTypeDiagramQuickStyle = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramQuickStyle"
)

// DiagramDataPart represents a diagram data part (word/diagrams/data1.xml, etc.).
// This part contains the SmartArt diagram data and node structure.
type DiagramDataPart struct {
	*openxml.OpenXmlPartData
}

// newDiagramDataPart creates a new diagram data part.
//
//nolint:unused // Will be used when diagram creation API is implemented
func newDiagramDataPart(
	mainPart *MainPart,
	uri string,
) (*DiagramDataPart, error) {
	packPart, relID, err := mainPart.addChildPart(
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
		mainPart,
	)
	partData.SetRelationshipID(relID)

	ddp := &DiagramDataPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal diagram data content
	ddp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(ddp, relID); err != nil {
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
			DefaultURI:         "/word/diagrams/data1.xml",
			IsFixedContentType: true,
		},
	)
}
