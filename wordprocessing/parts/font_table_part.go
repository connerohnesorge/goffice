package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// FontsPart represents the font table part (word/fontTable.xml).
type FontsPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for fonts.
const (
	ContentTypeFontTable      = "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"
	RelationshipTypeFontTable = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"
)

// newFontsPart creates a new font table part.
func newFontsPart(mainPart *MainPart) (*FontsPart, error) {
	uri := "/word/fontTable.xml"

	packPart, relID, err := mainPart.addChildPart(uri, ContentTypeFontTable, RelationshipTypeFontTable)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeFontTable, packPart, mainPart)
	partData.SetRelationshipID(relID)

	fp := &FontsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal font table content
	fp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(fp, relID); err != nil {
		return nil, err
	}

	return fp, nil
}

// initializeContent sets up minimal font table content.
func (fp *FontsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:fonts xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:font w:name="Calibri">
    <w:panose1 w:val="020F0502020204030204"/>
    <w:charset w:val="00"/>
    <w:family w:val="swiss"/>
    <w:pitch w:val="variable"/>
  </w:font>
  <w:font w:name="Times New Roman">
    <w:panose1 w:val="02020603050405020304"/>
    <w:charset w:val="00"/>
    <w:family w:val="roman"/>
    <w:pitch w:val="variable"/>
  </w:font>
</w:fonts>`
	fp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (fp *FontsPart) FixedContentType() string {
	return ContentTypeFontTable
}

// Fonts returns the root Fonts element.
// TODO: Return a proper Fonts element type when elements are implemented.
func (fp *FontsPart) Fonts() openxml.PartRootElement {
	return fp.RootElement()
}

// GetFont returns a font definition by name.
// TODO: Implement proper Font element type.
func (fp *FontsPart) GetFont(name string) interface{} {
	// TODO: Parse fonts and find by name
	return nil
}

// GetStream returns a reader for the part content.
func (fp *FontsPart) GetStream() io.Reader {
	return fp.OpenXmlPartData.GetStream()
}

// Ensure FontsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*FontsPart)(nil)

// FontsPartFactory creates a FontsPart from a URI and container.
func FontsPartFactory(uri string, container openxml.OpenXmlPartContainer) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeFontTable, packPart, container)
	return &FontsPart{
		OpenXmlPartData: partData,
	}
}

// Register the FontsPart type.
func init() {
	openxml.RegisterPartType(&openxml.PartTypeInfo{
		ContentType:        ContentTypeFontTable,
		RelationshipType:   RelationshipTypeFontTable,
		Factory:            FontsPartFactory,
		DefaultURI:         "/word/fontTable.xml",
		IsFixedContentType: true,
	})
}
