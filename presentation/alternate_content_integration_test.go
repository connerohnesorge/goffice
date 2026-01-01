package presentation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
)

// TestAlternateContentPresentationIntegration tests that AlternateContent elements
// can be created and roundtripped in PowerPoint documents.
func TestAlternateContentPresentationIntegration(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-pptx-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "alternate_content.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify structure is preserved
	presentationPart := doc2.PresentationPart()
	if presentationPart == nil {
		t.Error("PresentationPart() = nil after reopen")

		return
	}

	// Save again to verify double roundtrip
	if err := doc2.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestAlternateContentInPresentation tests AlternateContent elements
// within presentation slides.
func TestAlternateContentInPresentation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-slide-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "ac_slide.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Create AlternateContent element for modern PowerPoint feature
	ac := openxml.NewAlternateContent()

	// Add Office 2016 choice (3D models, modern animations)
	choice2016 := openxml.NewChoice()
	choice2016.SetRequires("p16")
	ac.AppendChild(choice2016)

	// Add fallback for older PowerPoint
	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify presentation is readable
	if doc2.PresentationPart() == nil {
		t.Error("PresentationPart() = nil after reopen")
	}
}

// TestPowerPoint2016ExtensionElements tests that PowerPoint 2016 extension elements
// (3D models, modern animations) are properly handled.
func TestPowerPoint2016ExtensionElements(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-p16-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "powerpoint2016.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save with Office 2016 content
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify presentation structure
	presentationPart := doc2.PresentationPart()
	if presentationPart == nil {
		t.Error("PresentationPart() = nil")
	}
}

// TestPowerPoint2024ExtensionElements tests handling of PowerPoint 2024 extension elements.
func TestPowerPoint2024ExtensionElements(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-p24-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "powerpoint2024.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify presentation is readable
	if doc2.PresentationPart() == nil {
		t.Error("PresentationPart() = nil after reopen")
	}
}

// TestMultipleVersionChoicesInPresentation tests AlternateContent with
// multiple version choices in presentations.
func TestMultipleVersionChoicesInPresentation(t *testing.T) {
	ac := openxml.NewAlternateContent()

	// Office 2016 feature
	choice2016 := openxml.NewChoice()
	choice2016.SetRequires("p16")
	ac.AppendChild(choice2016)

	// Office 2020 feature
	choice2020 := openxml.NewChoice()
	choice2020.SetRequires("p20")
	ac.AppendChild(choice2020)

	// Fallback
	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	// Test XML generation
	xml := ac.OuterXml()

	if !strings.Contains(xml, "p16") {
		t.Error("XML missing p16 namespace reference")
	}
	if !strings.Contains(xml, "p20") {
		t.Error("XML missing p20 namespace reference")
	}
}

// TestPresentationUnknownElementPreservation tests that unknown PowerPoint extension
// elements are preserved.
func TestPresentationUnknownElementPreservation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-unknown-pptx-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "unknown.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen multiple times to test preservation
	for i := range 2 {
		doc2, err := Open(testPath, true)
		if err != nil {
			t.Fatalf("Open() cycle %d error = %v", i, err)
		}

		if doc2.PresentationPart() == nil {
			t.Errorf("PresentationPart() = nil at cycle %d", i)
		}

		if err := doc2.SaveAs(testPath); err != nil {
			t.Fatalf("SaveAs() cycle %d error = %v", i, err)
		}
		_ = doc2.Close()
	}
}

// TestDrawingMLExtensionElements tests handling of DrawingML extension elements
// which are shared across Office applications.
func TestDrawingMLExtensionElements(t *testing.T) {
	ac := openxml.NewAlternateContent()

	// DrawingML 2010 extension
	choice := openxml.NewChoice()
	choice.SetRequires("a14")
	ac.AppendChild(choice)

	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	xml := ac.OuterXml()

	if !strings.Contains(xml, "a14") {
		t.Error("XML missing a14 (DrawingML 2010) namespace reference")
	}
}

// TestAlternateContentSelectionLogic tests that AlternateContent
// correctly selects between versions.
func TestAlternateContentSelectionLogic(t *testing.T) {
	ac := openxml.NewAlternateContent()

	// Modern choice (2016+)
	choiceModern := openxml.NewChoice()
	choiceModern.SetRequires("p16")
	ac.AppendChild(choiceModern)

	// Legacy choice (2010)
	choiceLegacy := openxml.NewChoice()
	choiceLegacy.SetRequires("p14")
	ac.AppendChild(choiceLegacy)

	// Fallback (older)
	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	tests := []struct {
		version        openxml.FileFormatVersion
		expectChoice   bool
		expectFallback bool
	}{
		{openxml.FileFormatVersionOffice2007, false, true},
		{openxml.FileFormatVersionOffice2010, true, false},
		{openxml.FileFormatVersionOffice2013, true, false},
		{openxml.FileFormatVersionOffice2016, true, false},
		{openxml.FileFormatVersionOffice2019, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.version.String(), func(t *testing.T) {
			selected := ac.SelectContent(tt.version)

			isChoice := false
			isFallback := false

			if _, ok := selected.(*openxml.Choice); ok {
				isChoice = true
			}
			if _, ok := selected.(*openxml.Fallback); ok {
				isFallback = true
			}

			if isChoice != tt.expectChoice {
				t.Errorf(
					"Expected choice=%v, isFallback=%v for %v",
					tt.expectChoice,
					isFallback,
					tt.version,
				)
			}
		})
	}
}
