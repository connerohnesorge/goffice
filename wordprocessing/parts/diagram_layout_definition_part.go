//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// DiagramLayoutDefinitionPart represents a diagram layout definition part (word/diagrams/layout1.xml, etc.).
// This part contains the SmartArt layout algorithm and rules.
type DiagramLayoutDefinitionPart struct {
	*openxml.OpenXmlPartData
}

// newDiagramLayoutDefinitionPart creates a new diagram layout definition part.
//
//nolint:unused // Will be used when diagram creation API is implemented
func newDiagramLayoutDefinitionPart(
	mainPart *MainPart,
	uri string,
) (*DiagramLayoutDefinitionPart, error) {
	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeDiagramLayoutDef,
		RelationshipTypeDiagramLayout,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramLayoutDef,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	dldp := &DiagramLayoutDefinitionPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal layout definition content
	dldp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(dldp, relID); err != nil {
		return nil, err
	}

	return dldp, nil
}

// initializeContent sets up minimal layout definition content.
//
//nolint:unused // Will be used when diagram creation API is implemented
func (dldp *DiagramLayoutDefinitionPart) initializeContent() {
	layoutDef := diagram.NewLayoutDefinition()
	dldp.SetRootElement(layoutDef)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*DiagramLayoutDefinitionPart) FixedContentType() string {
	return ContentTypeDiagramLayoutDef
}

// LayoutDefinition returns the root LayoutDefinition element.
func (dldp *DiagramLayoutDefinitionPart) LayoutDefinition() *diagram.LayoutDefinition {
	root := dldp.RootElement()
	if root == nil {
		return nil
	}
	if layoutDef, ok := root.(*diagram.LayoutDefinition); ok {
		return layoutDef
	}

	return nil
}

// GetStream returns a reader for the part content.
func (dldp *DiagramLayoutDefinitionPart) GetStream() io.Reader {
	return dldp.OpenXmlPartData.GetStream()
}

// Ensure DiagramLayoutDefinitionPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*DiagramLayoutDefinitionPart)(
	nil,
)

// DiagramLayoutDefinitionPartFactory creates a DiagramLayoutDefinitionPart from a URI and container.
func DiagramLayoutDefinitionPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramLayoutDef,
		packPart,
		container,
	)

	return &DiagramLayoutDefinitionPart{
		OpenXmlPartData: partData,
	}
}

// Register the DiagramLayoutDefinitionPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeDiagramLayoutDef,
			RelationshipType:   RelationshipTypeDiagramLayout,
			Factory:            DiagramLayoutDefinitionPartFactory,
			DefaultURI:         "/word/diagrams/layout1.xml",
			IsFixedContentType: true,
		},
	)
}
