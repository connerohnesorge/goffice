//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
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
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*NumberingPart) FixedContentType() string {
	return ContentTypeNumbering
}

// Numbering returns the root Numbering element.
func (np *NumberingPart) Numbering() openxml.PartRootElement {
	return np.RootElement()
}

// GetAbstractNum returns an abstract numbering definition by ID.
func (np *NumberingPart) GetAbstractNum(
	id int,
) *elements.AbstractNum {
	root := np.RootElement()
	if root == nil {
		return nil
	}

	// Iterate through child elements to find the abstractNum with the matching abstractNumId
	for child := range root.Children() {
		if child.LocalName() == "abstractNum" &&
			child.NamespaceURI() == elements.NamespaceWML {
			// Check if this abstractNum has the ID we're looking for
			attr, found := child.GetAttribute(
				"abstractNumId",
				elements.NamespaceWML,
			)
			if found {
				// Parse the ID and compare
				var childID int
				if _, err := fmt.Sscanf(attr.Value(), "%d", &childID); err == nil &&
					childID == id {
					// Convert to *elements.AbstractNum
					if absNum, ok := child.(*elements.AbstractNum); ok {
						return absNum
					}
					// Try to wrap if it's a CompositeElementBase
					if comp, ok := child.(*openxml.CompositeElementBase); ok {
						return &elements.AbstractNum{
							CompositeElementBase: comp,
						}
					}
				}
			}
		}
	}

	return nil
}

// GetNumInstance returns a numbering instance by ID.
func (np *NumberingPart) GetNumInstance(
	id int,
) *elements.NumberingInstance {
	root := np.RootElement()
	if root == nil {
		return nil
	}

	// Iterate through child elements to find the num with the matching numId
	for child := range root.Children() {
		if child.LocalName() == "num" &&
			child.NamespaceURI() == elements.NamespaceWML {
			// Check if this num has the ID we're looking for
			attr, found := child.GetAttribute(
				"numId",
				elements.NamespaceWML,
			)
			if found {
				// Parse the ID and compare
				var childID int
				if _, err := fmt.Sscanf(attr.Value(), "%d", &childID); err == nil &&
					childID == id {
					// Convert to *elements.NumberingInstance
					if numInst, ok := child.(*elements.NumberingInstance); ok {
						return numInst
					}
					// Try to wrap if it's a CompositeElementBase
					if comp, ok := child.(*openxml.CompositeElementBase); ok {
						return &elements.NumberingInstance{
							CompositeElementBase: comp,
						}
					}
				}
			}
		}
	}

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
	// Use GetPackagingPart instead of Package() to avoid locking issues
	// during initialization when the package lock is already held
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
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
