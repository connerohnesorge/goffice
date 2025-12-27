//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// CalculationChainPart represents the calculation chain part (xl/calcChain.xml).
// This part defines the order in which formulas should be calculated.
type CalculationChainPart struct {
	*openxml.OpenXmlPartData
}

// newCalculationChainPart creates a new calculation chain part.
func newCalculationChainPart(
	workbookPart *WorkbookPart,
) (*CalculationChainPart, error) {
	uri := "/xl/calcChain.xml"

	packPart, relID, err := workbookPart.addChildPart(
		uri,
		ContentTypeCalcChain,
		RelationshipTypeCalcChain,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCalcChain,
		packPart,
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	ccp := &CalculationChainPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal calculation chain content
	ccp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(ccp, relID); err != nil {
		return nil, err
	}

	return ccp, nil
}

// initializeContent sets up minimal calculation chain content.
func (ccp *CalculationChainPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<calcChain xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
</calcChain>`
	ccp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*CalculationChainPart) FixedContentType() string {
	return ContentTypeCalcChain
}

// CalculationChain returns the root CalculationChain element.
// TODO: Return a proper CalculationChain element type when elements are implemented.
func (ccp *CalculationChainPart) CalculationChain() openxml.PartRootElement {
	return ccp.RootElement()
}

// GetStream returns a reader for the part content.
func (ccp *CalculationChainPart) GetStream() io.Reader {
	return ccp.OpenXmlPartData.GetStream()
}

// Ensure CalculationChainPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*CalculationChainPart)(
	nil,
)

// CalculationChainPartFactory creates a CalculationChainPart from a URI and container.
func CalculationChainPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeCalcChain,
		packPart,
		container,
	)

	return &CalculationChainPart{
		OpenXmlPartData: partData,
	}
}

// Register the CalculationChainPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeCalcChain,
			RelationshipType:   RelationshipTypeCalcChain,
			Factory:            CalculationChainPartFactory,
			DefaultURI:         "/xl/calcChain.xml",
			IsFixedContentType: true,
		},
	)
}
