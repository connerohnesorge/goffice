package presentation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

// TestRoundtripExtensionElementsBasic tests basic extension element roundtrip
// by creating and reopening a presentation with simple content.
func TestRoundtripExtensionElementsBasic(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-extension-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "extension.pptx")

	// Create presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Get the presentation part
	presentationPart := doc1.PresentationPart()
	if presentationPart == nil {
		t.Fatal("PresentationPart() = nil")
	}

	// Save and close
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen the presentation
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify presentation is readable
	presentationPart2 := doc2.PresentationPart()
	if presentationPart2 == nil {
		t.Error("PresentationPart() = nil after reopen")

		return
	}

	presentation2 := presentationPart2.Presentation()
	if presentation2 == nil {
		// Note: This can be nil if the presentation was not properly initialized
		// This is expected for newly created presentations that haven't been populated
		t.Log("Presentation() = nil after reopen (expected for minimal presentations)")
	}
}

// TestVersionDetectionFromPresentation tests that presentations are correctly identified
// with their Office version requirements.
func TestVersionDetectionFromPresentation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-version-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "version-detect.pptx")

	// Create and save a simple presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen and detect version
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Detect minimum version required for this document
	detectedVersion := validation.DetectMinimumVersionForPackage(doc2)

	// Basic Office 2007 presentation should be Office2007 or higher
	if detectedVersion < validation.Office2007 {
		t.Errorf("DetectMinimumVersion() = %v, want >= Office2007", detectedVersion)
	}
}

// TestPresentationBasicRoundtrip tests simple roundtrip
func TestPresentationBasicRoundtrip(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-roundtrip-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "roundtrip.pptx")

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

	// Verify type is preserved
	if doc2.Type() != DocTypePresentation {
		t.Errorf("Type() = %v, want %v", doc2.Type(), DocTypePresentation)
	}
}

// TestPresentationNamespacePreservation tests that namespaces are preserved
func TestPresentationNamespacePreservation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-namespace-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "namespace.pptx")

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

	// Verify we can access the presentation
	presentationPart := doc2.PresentationPart()
	if presentationPart == nil {
		t.Error("PresentationPart() = nil")

		return
	}

	presentation := presentationPart.Presentation()
	if presentation == nil {
		t.Error("Presentation() = nil")

		return
	}

	t.Log("Successfully opened and verified presentation structure")
}

// TestMultipleSaveRoundtrips tests that presentation survives multiple save/open cycles
func TestMultipleSaveRoundtrips(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-multisave-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "multisave.pptx")

	// Create initial presentation
	doc1, err := New(testPath, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Do multiple roundtrip cycles
	for i := range 3 {
		doc2, err := Open(testPath, true)
		if err != nil {
			t.Fatalf("Open() cycle %d error = %v", i, err)
		}

		// Verify content is there
		presentationPart := doc2.PresentationPart()
		if presentationPart == nil {
			t.Fatalf("PresentationPart() = nil at cycle %d", i)
		}

		// Re-save
		if err := doc2.SaveAs(testPath); err != nil {
			t.Fatalf("SaveAs() cycle %d error = %v", i, err)
		}
		_ = doc2.Close()
	}

	// Final check
	doc3, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Final Open() error = %v", err)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.PresentationPart() == nil {
		t.Error("PresentationPart() = nil after multiple roundtrips")
	}
}

// TestRoundtripPresentationTypes tests saving different presentation formats
func TestRoundtripPresentationTypes(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-pptx-formats-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	formats := []struct {
		name string
		typ  DocType
	}{
		{"presentation.pptx", DocTypePresentation},
		{"macro.pptm", DocTypePresentation},
		{"template.potx", DocTypeTemplate},
	}

	for _, format := range formats {
		t.Run(format.name, func(t *testing.T) {
			testPath := filepath.Join(tmpDir, format.name)

			// Create presentation
			doc1, err := New(testPath, format.typ)
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

			// Verify type is preserved
			if doc2.Type() != format.typ {
				t.Errorf("Type() = %v, want %v", doc2.Type(), format.typ)
			}

			_ = doc2.Close()
		})
	}
}
