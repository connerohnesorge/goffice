//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// GlossaryPart represents the glossary document part (word/glossary/document.xml).
// This part contains reusable building blocks like AutoText and Quick Parts.
type GlossaryPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for glossary.
const (
	ContentTypeGlossary      = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml"
	RelationshipTypeGlossary = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/glossaryDocument"
)

// newGlossaryPart creates a new glossary document part.
func newGlossaryPart(
	mainPart *MainPart,
) (*GlossaryPart, error) {
	uri := "/word/glossary/document.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeGlossary,
		RelationshipTypeGlossary,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeGlossary,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	gp := &GlossaryPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal glossary content
	gp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(gp, relID); err != nil {
		return nil, err
	}

	return gp, nil
}

// initializeContent sets up minimal glossary document content.
func (gp *GlossaryPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:glossaryDocument xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:docParts/>
</w:glossaryDocument>`
	gp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (*GlossaryPart) FixedContentType() string {
	return ContentTypeGlossary
}

// GlossaryDocument returns the root GlossaryDocument element.
// TODO: Return a proper GlossaryDocument element type when elements are implemented.
func (gp *GlossaryPart) GlossaryDocument() openxml.PartRootElement {
	return gp.RootElement()
}

// GetStream returns a reader for the part content.
func (gp *GlossaryPart) GetStream() io.Reader {
	return gp.OpenXmlPartData.GetStream()
}

// Ensure GlossaryPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*GlossaryPart)(nil)

// GlossaryPartFactory creates a GlossaryPart from a URI and container.
func GlossaryPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Use GetPackagingPart instead of Package() to avoid locking issues
	// during initialization when the package lock is already held
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeGlossary,
		packPart,
		container,
	)

	return &GlossaryPart{
		OpenXmlPartData: partData,
	}
}

// Register the GlossaryPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeGlossary,
			RelationshipType:   RelationshipTypeGlossary,
			Factory:            GlossaryPartFactory,
			DefaultURI:         "/word/glossary/document.xml",
			IsFixedContentType: true,
		},
	)
}
