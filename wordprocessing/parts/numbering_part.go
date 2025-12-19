package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// NumberingPart represents the numbering definitions part (word/numbering.xml).
type NumberingPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for numbering.
const (
	ContentTypeNumbering      = "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"
	RelationshipTypeNumbering = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"
)

// newNumberingPart creates a new numbering definitions part.
func newNumberingPart(
	mainPart *MainPart,
) (*NumberingPart, error) {
	uri := "/word/numbering.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeNumbering,
		RelationshipTypeNumbering,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNumbering,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	np := &NumberingPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal numbering content
	np.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(np, relID); err != nil {
		return nil, err
	}

	return np, nil
}

// initializeContent sets up minimal numbering content.
func (np *NumberingPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:numbering xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
</w:numbering>`
	np.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (np *NumberingPart) FixedContentType() string {
	return ContentTypeNumbering
}

// Numbering returns the root Numbering element.
// TODO: Return a proper Numbering element type when elements are implemented.
func (np *NumberingPart) Numbering() openxml.PartRootElement {
	return np.RootElement()
}

// GetAbstractNum returns an abstract numbering definition by ID.
// TODO: Implement proper AbstractNum element type.
func (np *NumberingPart) GetAbstractNum(
	id int,
) interface{} {
	// TODO: Parse numbering and find abstract num by ID
	return nil
}

// GetNumInstance returns a numbering instance by ID.
// TODO: Implement proper NumInstance element type.
func (np *NumberingPart) GetNumInstance(
	id int,
) interface{} {
	// TODO: Parse numbering and find num instance by ID
	return nil
}

// GetStream returns a reader for the part content.
func (np *NumberingPart) GetStream() io.Reader {
	return np.OpenXmlPartData.GetStream()
}

// Ensure NumberingPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*NumberingPart)(nil)

// NumberingPartFactory creates a NumberingPart from a URI and container.
func NumberingPartFactory(
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
		ContentTypeNumbering,
		packPart,
		container,
	)
	return &NumberingPart{
		OpenXmlPartData: partData,
	}
}

// Register the NumberingPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeNumbering,
			RelationshipType:   RelationshipTypeNumbering,
			Factory:            NumberingPartFactory,
			DefaultURI:         "/word/numbering.xml",
			IsFixedContentType: true,
		},
	)
}
