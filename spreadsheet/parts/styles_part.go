//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// WorkbookStylesPart represents the workbook styles part (xl/styles.xml).
// This part contains cell formatting definitions including fonts, fills,
// borders, number formats, and cell styles.
type WorkbookStylesPart struct {
	*openxml.OpenXmlPartData
}

// newWorkbookStylesPart creates a new workbook styles part.
func newWorkbookStylesPart(
	workbookPart *WorkbookPart,
) (*WorkbookStylesPart, error) {
	uri := "/xl/styles.xml"

	packPart, relID, err := workbookPart.addChildPart(
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
		workbookPart,
	)
	partData.SetRelationshipID(relID)

	sp := &WorkbookStylesPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal styles content
	sp.initializeContent()

	// Add to workbook part's child parts
	if err := workbookPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// initializeContent sets up minimal styles content with default formatting.
func (sp *WorkbookStylesPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<styleSheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
  <numFmts count="0"/>
  <fonts count="1">
    <font>
      <sz val="11"/>
      <color theme="1"/>
      <name val="Calibri"/>
      <family val="2"/>
      <scheme val="minor"/>
    </font>
  </fonts>
  <fills count="2">
    <fill>
      <patternFill patternType="none"/>
    </fill>
    <fill>
      <patternFill patternType="gray125"/>
    </fill>
  </fills>
  <borders count="1">
    <border>
      <left/>
      <right/>
      <top/>
      <bottom/>
      <diagonal/>
    </border>
  </borders>
  <cellStyleXfs count="1">
    <xf numFmtId="0" fontId="0" fillId="0" borderId="0"/>
  </cellStyleXfs>
  <cellXfs count="1">
    <xf numFmtId="0" fontId="0" fillId="0" borderId="0" xfId="0"/>
  </cellXfs>
  <cellStyles count="1">
    <cellStyle name="Normal" xfId="0" builtinId="0"/>
  </cellStyles>
</styleSheet>`
	sp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*WorkbookStylesPart) FixedContentType() string {
	return ContentTypeStyles
}

// Stylesheet returns the root Stylesheet element.
// TODO: Return a proper Stylesheet element type when elements are implemented.
func (sp *WorkbookStylesPart) Stylesheet() openxml.PartRootElement {
	return sp.RootElement()
}

// GetStream returns a reader for the part content.
func (sp *WorkbookStylesPart) GetStream() io.Reader {
	return sp.OpenXmlPartData.GetStream()
}

// Ensure WorkbookStylesPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WorkbookStylesPart)(
	nil,
)

// WorkbookStylesPartFactory creates a WorkbookStylesPart from a URI and container.
func WorkbookStylesPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
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

	return &WorkbookStylesPart{
		OpenXmlPartData: partData,
	}
}

// Register the WorkbookStylesPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeStyles,
			RelationshipType:   RelationshipTypeStyles,
			Factory:            WorkbookStylesPartFactory,
			DefaultURI:         "/xl/styles.xml",
			IsFixedContentType: true,
		},
	)
}
