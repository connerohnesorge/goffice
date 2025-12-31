//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// DiagramStylePart represents a diagram style part (ppt/diagrams/quickStyle1.xml, etc.).
// This part contains the SmartArt quick style definition.
type DiagramStylePart struct {
	*openxml.OpenXmlPartData
}

// newDiagramStylePart creates a new diagram style part.
//
//nolint:unused // Will be used when diagram creation API is implemented
func newDiagramStylePart(
	parent openxml.OpenXmlPartContainer,
	uri string,
) (*DiagramStylePart, error) {
	parentPart, ok := parent.(partWithPackage)
	if !ok {
		return nil, ErrNilPackage
	}

	packPart, relID, err := addChildPart(
		parentPart,
		uri,
		ContentTypeDiagramStyle,
		RelationshipTypeDiagramStyle,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramStyle,
		packPart,
		parent,
	)
	partData.SetRelationshipID(relID)

	dsp := &DiagramStylePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal style definition content
	dsp.initializeContent()

	// Add to parent's child parts
	if err := parent.AddPart(dsp, relID); err != nil {
		return nil, err
	}

	return dsp, nil
}

// initializeContent sets up minimal style definition content.
//
//nolint:unused // Will be used when diagram creation API is implemented
func (dsp *DiagramStylePart) initializeContent() {
	styleDef := diagram.NewStyleDefinition()
	dsp.SetRootElement(styleDef)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*DiagramStylePart) FixedContentType() string {
	return ContentTypeDiagramStyle
}

// StyleDefinition returns the root StyleDefinition element.
func (dsp *DiagramStylePart) StyleDefinition() *diagram.StyleDefinition {
	root := dsp.RootElement()
	if root == nil {
		return nil
	}
	if styleDef, ok := root.(*diagram.StyleDefinition); ok {
		return styleDef
	}

	return nil
}

// GetStream returns a reader for the part content.
func (dsp *DiagramStylePart) GetStream() io.Reader {
	return dsp.OpenXmlPartData.GetStream()
}

// Ensure DiagramStylePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*DiagramStylePart)(
	nil,
)

// DiagramStylePartFactory creates a DiagramStylePart from a URI and container.
func DiagramStylePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeDiagramStyle,
		packPart,
		container,
	)

	return &DiagramStylePart{
		OpenXmlPartData: partData,
	}
}

// Register the DiagramStylePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeDiagramStyle,
			RelationshipType:   RelationshipTypeDiagramStyle,
			Factory:            DiagramStylePartFactory,
			DefaultURI:         "/ppt/diagrams/quickStyle1.xml",
			IsFixedContentType: true,
		},
	)
}
