//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
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
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewStyles()
		},
	)
	partData.SetRelationshipID(relID)

	sp := &StylesPart{
		OpenXmlPartData: partData,
	}

	// Add to main part's child parts
	if err := mainPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// InitializeDefault sets up minimal styles content.
func (sp *StylesPart) InitializeDefault() {
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
func (sp *StylesPart) Styles() openxml.PartRootElement {
	return sp.RootElement()
}

// GetStyleById returns a style by its ID.
func (sp *StylesPart) GetStyleById(
	id string,
) *elements.Style {
	root := sp.RootElement()
	if root == nil {
		return nil
	}

	// Iterate through child elements to find the style with the matching styleId
	for child := range root.Children() {
		if child.LocalName() == "style" &&
			child.NamespaceURI() == elements.NamespaceWML {
			// Check if this style has the ID we're looking for
			attr, found := child.GetAttribute(
				"styleId",
				elements.NamespaceWML,
			)
			if found && attr.Value() == id {
				// Convert to *elements.Style
				if style, ok := child.(*elements.Style); ok {
					return style
				}
				// Try to wrap if it's a CompositeElementBase
				if comp, ok := child.(*openxml.CompositeElementBase); ok {
					return &elements.Style{
						CompositeElementBase: comp,
					}
				}
			}
		}
	}

	return nil
}

// GetStyleByName returns a style by its name.
func (sp *StylesPart) GetStyleByName(
	name string,
) *elements.Style {
	root := sp.RootElement()
	if root == nil {
		return nil
	}

	// Iterate through child elements to find the style with the matching name
	for child := range root.Children() {
		if child.LocalName() == "style" &&
			child.NamespaceURI() == elements.NamespaceWML {
			// Check if child is a CompositeElement to access GetElement
			if compChild, ok := child.(openxml.CompositeElement); ok {
				// Check for the name child element
				nameElem := compChild.GetElement(
					"name",
					elements.NamespaceWML,
				)
				if nameElem != nil {
					attr, found := nameElem.GetAttribute(
						"val",
						elements.NamespaceWML,
					)
					if found &&
						attr.Value() == name {
						// Convert to *elements.Style
						if style, ok := child.(*elements.Style); ok {
							return style
						}
						// Try to wrap if it's a CompositeElementBase
						if comp, ok := child.(*openxml.CompositeElementBase); ok {
							return &elements.Style{
								CompositeElementBase: comp,
							}
						}
					}
				}
			}
		}
	}

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
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewStyles()
		},
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
