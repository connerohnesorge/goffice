//nolint:revive // Package contains long URL lines
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
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
func newWebSettingsPart(
	mainPart *MainPart,
) (*WebSettingsPart, error) {
	uri := "/word/webSettings.xml"

	packPart, relID, err := mainPart.addChildPart(
		uri,
		ContentTypeWebSettings,
		RelationshipTypeWebSettings,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeWebSettings,
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWebSettings()
		},
	)

	ws := &WebSettingsPart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal web settings content
	ws.initializeContent()

	// Add to main part's child parts
	if err := mainPart.AddPart(ws, relID); err != nil {
		return nil, err
	}

	return ws, nil
}

// initializeContent sets up minimal web settings content.
func (ws *WebSettingsPart) initializeContent() {
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:webSettings xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:optimizeForBrowser/>
  <w:allowPNG/>
</w:webSettings>`
	ws.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
func (ws *WebSettingsPart) FixedContentType() string {
	return ContentTypeWebSettings
}

// WebSettings returns the root WebSettings element.
func (ws *WebSettingsPart) WebSettings() *elements.WebSettings {
	root := ws.RootElement()
	if root == nil {
		return nil
	}
	if s, ok := root.(*elements.WebSettings); ok {
		return s
	}

	return nil
}

// GetStream returns a reader for the part content.
func (ws *WebSettingsPart) GetStream() io.Reader {
	return ws.OpenXmlPartData.GetStream()
}

// Ensure WebSettingsPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*WebSettingsPart)(
	nil,
)

// WebSettingsPartFactory creates a WebSettingsPart from a URI and container.
func WebSettingsPartFactory(
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
		ContentTypeWebSettings,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWebSettings()
		},
	)

	return &WebSettingsPart{
		OpenXmlPartData: partData,
	}
}

// Register the WebSettingsPart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeWebSettings,
			RelationshipType:   RelationshipTypeWebSettings,
			Factory:            WebSettingsPartFactory,
			DefaultURI:         "/word/webSettings.xml",
			IsFixedContentType: true,
		},
	)
}
