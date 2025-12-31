//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// DiagramColorsPart represents a diagram colors part (ppt/diagrams/colors1.xml, etc.).
// This part contains the SmartArt color scheme definition.
type DiagramColorsPart struct {
	*openxml.OpenXmlPartData
}

// newDiagramColorsPart creates a new diagram colors part.
//
//nolint:unused // Will be used when diagram creation API is implemented
func newDiagramColorsPart(
	parent openxml.OpenXmlPartContainer,
	uri string,
) (*DiagramColorsPart, error) {
	parentPart, ok := parent.(partWithPackage)
	if !ok {
		return nil, ErrNilPackage
	}

	packPart, relID, err := addChildPart(
		parentPart,
		uri,
		ContentTypeDiagramColors,
		RelationshipTypeDiagramColors,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramColors,
		packPart,
		parent,
	)
	partData.SetRelationshipID(relID)

	dcp := &DiagramColorsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal colors definition content
	dcp.initializeContent()

	// Add to parent's child parts
	if err := parent.AddPart(dcp, relID); err != nil {
		return nil, err
	}

	return dcp, nil
}

// initializeContent sets up minimal colors definition content.
//
//nolint:unused // Will be used when diagram creation API is implemented
func (dcp *DiagramColorsPart) initializeContent() {
	colorsDef := diagram.NewColorsDefinition()
	dcp.SetRootElement(colorsDef)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*DiagramColorsPart) FixedContentType() string {
	return ContentTypeDiagramColors
}

// ColorsDefinition returns the root ColorsDefinition element.
func (dcp *DiagramColorsPart) ColorsDefinition() *diagram.ColorsDefinition {
	root := dcp.RootElement()
	if root == nil {
		return nil
	}
	if colorsDef, ok := root.(*diagram.ColorsDefinition); ok {
		return colorsDef
	}

	return nil
}

// GetStream returns a reader for the part content.
func (dcp *DiagramColorsPart) GetStream() io.Reader {
	return dcp.OpenXmlPartData.GetStream()
}

// Ensure DiagramColorsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*DiagramColorsPart)(
	nil,
)

// DiagramColorsPartFactory creates a DiagramColorsPart from a URI and container.
func DiagramColorsPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramColors,
		packPart,
		container,
	)

	return &DiagramColorsPart{
		OpenXmlPartData: partData,
	}
}

// Register the DiagramColorsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeDiagramColors,
			RelationshipType:   RelationshipTypeDiagramColors,
			Factory:            DiagramColorsPartFactory,
			DefaultURI:         "/ppt/diagrams/colors1.xml",
			IsFixedContentType: true,
		},
	)
}
