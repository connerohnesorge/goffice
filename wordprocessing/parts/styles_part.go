//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// StylesPart represents the styles definitions part (word/styles.xml).
type StylesPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for styles.
const (
	ContentTypeStyles      = "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"
	RelationshipTypeStyles = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
)

// newStylesPart creates a new styles definitions part.
func newStylesPart(
	mainPart *MainPart,
) (*StylesPart, error) {
	uri := "/word/styles.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeStyles,
		RelationshipTypeStyles,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeStyles,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	sp := &StylesPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal styles content
	sp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// initializeContent sets up minimal styles content.
func (sp *StylesPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:styles xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:docDefaults>
    <w:rPrDefault>
      <w:rPr>
        <w:rFonts w:ascii="Calibri" w:eastAsia="Calibri" w:hAnsi="Calibri" w:cs="Times New Roman"/>
        <w:sz w:val="22"/>
        <w:szCs w:val="22"/>
        <w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/>
      </w:rPr>
    </w:rPrDefault>
    <w:pPrDefault/>
  </w:docDefaults>
  <w:style w:type="paragraph" w:default="1" w:styleId="Normal">
    <w:name w:val="Normal"/>
    <w:qFormat/>
  </w:style>
</w:styles>`
	sp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*StylesPart) FixedContentType() string {
	return ContentTypeStyles
}

// Styles returns the root Styles element.
// TODO: Return a proper Styles element type when elements are implemented.
func (sp *StylesPart) Styles() openxml.PartRootElement {
	return sp.RootElement()
}

// GetStyleById returns a style by its ID.
// TODO: Implement proper Style element type.
//
//nolint:revive // unused-receiver: TODO stub implementation
func (*StylesPart) GetStyleById(
	_ string,
) any {
	// TODO: Parse styles and find by ID
	return nil
}

// GetStyleByName returns a style by its name.
// TODO: Implement proper Style element type.
//
//nolint:revive // unused-receiver: TODO stub implementation
func (*StylesPart) GetStyleByName(
	_ string,
) any {
	// TODO: Parse styles and find by name
	return nil
}

// GetStream returns a reader for the part content.
func (sp *StylesPart) GetStream() io.Reader {
	return sp.OpenXmlPartData.GetStream()
}

// Ensure StylesPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*StylesPart)(nil)

// StylesPartFactory creates a StylesPart from a URI and container.
func StylesPartFactory(
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
		ContentTypeStyles,
		packPart,
		container,
	)

	return &StylesPart{
		OpenXmlPartData: partData,
	}
}

// Register the StylesPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeStyles,
			RelationshipType:   RelationshipTypeStyles,
			Factory:            StylesPartFactory,
			DefaultURI:         "/word/styles.xml",
			IsFixedContentType: true,
		},
	)
}
