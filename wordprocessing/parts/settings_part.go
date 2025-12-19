package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// SettingsPart represents the document settings part (word/settings.xml).
type SettingsPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for settings.
const (
	ContentTypeSettings      = "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"
	RelationshipTypeSettings = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"
)

// newSettingsPart creates a new document settings part.
func newSettingsPart(
	mainPart *MainPart,
) (*SettingsPart, error) {
	uri := "/word/settings.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeSettings,
		RelationshipTypeSettings,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSettings,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	sp := &SettingsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal settings content
	sp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// initializeContent sets up minimal settings content.
func (sp *SettingsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:settings xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:defaultTabStop w:val="720"/>
  <w:characterSpacingControl w:val="doNotCompress"/>
  <w:compat>
    <w:compatSetting w:name="compatibilityMode" w:uri="http://schemas.microsoft.com/office/word" w:val="15"/>
  </w:compat>
</w:settings>`
	sp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (sp *SettingsPart) FixedContentType() string {
	return ContentTypeSettings
}

// Settings returns the root Settings element.
// TODO: Return a proper Settings element type when elements are implemented.
func (sp *SettingsPart) Settings() openxml.PartRootElement {
	return sp.RootElement()
}

// GetStream returns a reader for the part content.
func (sp *SettingsPart) GetStream() io.Reader {
	return sp.OpenXmlPartData.GetStream()
}

// Ensure SettingsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SettingsPart)(nil)

// SettingsPartFactory creates a SettingsPart from a URI and container.
func SettingsPartFactory(
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
		ContentTypeSettings,
		packPart,
		container,
	)

	return &SettingsPart{
		OpenXmlPartData: partData,
	}
}

// Register the SettingsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSettings,
			RelationshipType:   RelationshipTypeSettings,
			Factory:            SettingsPartFactory,
			DefaultURI:         "/word/settings.xml",
			IsFixedContentType: true,
		},
	)
}
