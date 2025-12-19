package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
)

// WebSettingsPart represents the web settings part (word/webSettings.xml).
type WebSettingsPart struct {
	*openxml.OpenXmlPartData
}

// Content type and relationship type for web settings.
const (
	ContentTypeWebSettings      = "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"
	RelationshipTypeWebSettings = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"
)

// newWebSettingsPart creates a new web settings part.
func newWebSettingsPart(mainPart *MainPart) (*WebSettingsPart, error) {
	uri := "/word/webSettings.xml"

	packPart, relID, err := mainPart.addChildPart(uri, ContentTypeWebSettings, RelationshipTypeWebSettings)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeWebSettings, packPart, mainPart)
	partData.SetRelationshipID(relID)

	wsp := &WebSettingsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal web settings content
	wsp.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(wsp, relID); err != nil {
		return nil, err
	}

	return wsp, nil
}

// initializeContent sets up minimal web settings content.
func (wsp *WebSettingsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:webSettings xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:optimizeForBrowser/>
  <w:allowPNG/>
</w:webSettings>`
	wsp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (wsp *WebSettingsPart) FixedContentType() string {
	return ContentTypeWebSettings
}

// WebSettings returns the root WebSettings element.
// TODO: Return a proper WebSettings element type when elements are implemented.
func (wsp *WebSettingsPart) WebSettings() openxml.PartRootElement {
	return wsp.RootElement()
}

// GetStream returns a reader for the part content.
func (wsp *WebSettingsPart) GetStream() io.Reader {
	return wsp.OpenXmlPartData.GetStream()
}

// Ensure WebSettingsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WebSettingsPart)(nil)

// WebSettingsPartFactory creates a WebSettingsPart from a URI and container.
func WebSettingsPartFactory(uri string, container openxml.OpenXmlPartContainer) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(uri, ContentTypeWebSettings, packPart, container)
	return &WebSettingsPart{
		OpenXmlPartData: partData,
	}
}

// Register the WebSettingsPart type.
func init() {
	openxml.RegisterPartType(&openxml.PartTypeInfo{
		ContentType:        ContentTypeWebSettings,
		RelationshipType:   RelationshipTypeWebSettings,
		Factory:            WebSettingsPartFactory,
		DefaultURI:         "/word/webSettings.xml",
		IsFixedContentType: true,
	})
}
