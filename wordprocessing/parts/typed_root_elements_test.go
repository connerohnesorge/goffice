package parts

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestMainPartDocument tests that MainPart.Document() returns a typed Document element.
func TestMainPartDocument(t *testing.T) {
	// Create a MainPart
	mp := &MainPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/document.xml",
			ContentTypeDocument,
			nil,
			nil,
		),
		contentType: ContentTypeDocument,
	}

	// Set up the root factory
	mp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewDocument()
		},
	)

	// Set a document as root element
	doc := elements.NewDocument()
	mp.SetRootElement(doc)

	// Get the document
	result := mp.Document()
	if result == nil {
		t.Fatal("Document() returned nil")
	}

	// Verify it's a Document type
	if _, ok := any(result).(*elements.Document); !ok {
		t.Errorf(
			"Document() did not return *elements.Document, got %T",
			result,
		)
	}
}

// TestSettingsPartSettings tests that SettingsPart.Settings() returns a typed Settings element.
func TestSettingsPartSettings(t *testing.T) {
	// Create a SettingsPart
	sp := &SettingsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/settings.xml",
			ContentTypeSettings,
			nil,
			nil,
		),
	}

	// Set up the root factory
	sp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSettings()
		},
	)

	// Set a settings element as root
	settings := elements.NewSettings()
	sp.SetRootElement(settings)

	// Get the settings
	result := sp.Settings()
	if result == nil {
		t.Fatal("Settings() returned nil")
	}

	// Verify it's a Settings type
	if _, ok := any(result).(*elements.Settings); !ok {
		t.Errorf(
			"Settings() did not return *elements.Settings, got %T",
			result,
		)
	}
}

// TestWebSettingsPartWebSettings tests that WebSettingsPart.WebSettings() returns a typed WebSettings element.
func TestWebSettingsPartWebSettings(
	t *testing.T,
) {
	// Create a WebSettingsPart
	wsp := &WebSettingsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/webSettings.xml",
			ContentTypeWebSettings,
			nil,
			nil,
		),
	}

	// Set up the root factory
	wsp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWebSettings()
		},
	)

	// Set a web settings element as root
	webSettings := elements.NewWebSettings()
	wsp.SetRootElement(webSettings)

	// Get the web settings
	result := wsp.WebSettings()
	if result == nil {
		t.Fatal("WebSettings() returned nil")
	}

	// Verify it's a WebSettings type
	if _, ok := any(result).(*elements.WebSettings); !ok {
		t.Errorf(
			"WebSettings() did not return *elements.WebSettings, got %T",
			result,
		)
	}
}

// TestGlossaryPartGlossaryDocument tests that GlossaryPart.GlossaryDocument() returns a typed GlossaryDocument element.
func TestGlossaryPartGlossaryDocument(
	t *testing.T,
) {
	// Create a GlossaryPart
	gp := &GlossaryPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/glossary/document.xml",
			ContentTypeGlossary,
			nil,
			nil,
		),
	}

	// Set up the root factory
	gp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewGlossaryDocument()
		},
	)

	// Set a glossary document element as root
	glossaryDoc := elements.NewGlossaryDocument()
	gp.SetRootElement(glossaryDoc)

	// Get the glossary document
	result := gp.GlossaryDocument()
	if result == nil {
		t.Fatal("GlossaryDocument() returned nil")
	}

	// Verify it's a GlossaryDocument type
	if _, ok := any(result).(*elements.GlossaryDocument); !ok {
		t.Errorf(
			"GlossaryDocument() did not return *elements.GlossaryDocument, got %T",
			result,
		)
	}
}

// TestTypedElementsAreUsable tests that the typed elements can be used to access their properties.
func TestTypedElementsAreUsable(t *testing.T) {
	// Test Document element methods
	mp := &MainPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/document.xml",
			ContentTypeDocument,
			nil,
			nil,
		),
		contentType: ContentTypeDocument,
	}
	mp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewDocument()
		},
	)

	doc := elements.NewDocument()
	mp.SetRootElement(doc)

	result := mp.Document()
	if result == nil {
		t.Fatal("Document() returned nil")
	}

	// Access Body (method specific to Document type)
	body := result.Body()
	if body == nil {
		t.Error(
			"Document.Body() returned nil, expected non-nil body",
		)
	}

	// Test Settings-specific methods
	sp := &SettingsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/settings.xml",
			ContentTypeSettings,
			nil,
			nil,
		),
	}
	sp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewSettings()
		},
	)

	settings := elements.NewSettings()
	sp.SetRootElement(settings)

	resultSettings := sp.Settings()
	if resultSettings == nil {
		t.Fatal("Settings() returned nil")
	}

	// Test Settings-specific method
	resultSettings.SetZoom(150)
	zoom := resultSettings.Zoom()
	if zoom == nil {
		t.Error(
			"Settings.Zoom() returned nil after SetZoom",
		)
	}

	// Test WebSettings-specific methods
	wsp := &WebSettingsPart{
		OpenXmlPartData: openxml.NewOpenXmlPartData(
			"/word/webSettings.xml",
			ContentTypeWebSettings,
			nil,
			nil,
		),
	}
	wsp.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewWebSettings()
		},
	)

	webSettings := elements.NewWebSettings()
	wsp.SetRootElement(webSettings)

	resultWebSettings := wsp.WebSettings()
	if resultWebSettings == nil {
		t.Fatal("WebSettings() returned nil")
	}

	// Test WebSettings-specific method
	resultWebSettings.SetOptimizeForBrowser(true)
	if !resultWebSettings.OptimizeForBrowser() {
		t.Error(
			"WebSettings.OptimizeForBrowser() returned false after setting to true",
		)
	}
}
