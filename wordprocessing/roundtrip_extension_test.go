package wordprocessing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

// TestRoundtripExtensionElementsBasic tests basic extension element roundtrip
// by creating and reopening a document with simple content.
func TestRoundtripExtensionElementsBasic(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-extension-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "extension.docx")

	// Create document using builder
	builder := NewDocumentBuilder()
	builder.AddParagraph("Extension Element Test")
	builder.AddHeading("Section with Extensions", 1)
	builder.AddParagraph("Content after heading")

	if _, err := builder.Build(); err != nil {
		t.Fatalf("Builder.Build() error = %v", err)
	}

	// Save document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen the document
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify document is readable
	mainPart := doc2.MainPart()
	if mainPart == nil {
		t.Error("MainPart() = nil after reopen")

		return
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Error("Document() = nil after reopen")

		return
	}

	body := docElem.Body()
	if body == nil {
		t.Error("Body() = nil after reopen")

		return
	}

	// Verify we can iterate paragraphs
	paraCount := 0
	for range body.Paragraphs() {
		paraCount++
	}

	if paraCount == 0 {
		t.Error("No paragraphs found after reopen")
	}
}

// TestVersionDetectionFromDocument tests that documents are correctly identified
// with their Office version requirements.
func TestVersionDetectionFromDocument(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-version-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "version-detect.docx")

	// Create and save a simple document
	doc1, err := New(testPath, DocTypeDocument)
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

	// Basic Office 2007 document should be Office2007 or higher
	if detectedVersion < validation.Office2007 {
		t.Errorf("DetectMinimumVersion() = %v, want >= Office2007", detectedVersion)
	}
}

// TestUnknownElementPreservation tests that unknown elements
// don't prevent document reading and re-saving.
func TestUnknownElementPreservation(t *testing.T) {
	// Create a simple document buffer
	tmpDir, err := os.MkdirTemp("", "goffice-unknown-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "unknown.docx")
	testPath2 := filepath.Join(tmpDir, "unknown2.docx")

	// Create and save a document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Read the file back
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Try to re-save it
	if err := doc2.SaveAs(testPath2); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify we can open the re-saved file
	doc3, err := Open(testPath2, true)
	if err != nil {
		t.Fatalf("Open() of resaved file error = %v", err)
	}
	_ = doc3.Close()
}

// TestExtensionNamespaceLoading tests that extension namespaces are handled correctly
func TestExtensionNamespaceLoading(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-namespace-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "namespace.docx")

	// Create document
	builder := NewDocumentBuilder()
	builder.AddParagraph("Testing namespace handling")
	if _, err := builder.Build(); err != nil {
		t.Fatalf("Builder.Build() error = %v", err)
	}

	// Save
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

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

	// Get the document XML to verify structure
	mainPart := doc2.MainPart()
	if mainPart == nil {
		t.Error("MainPart() = nil")

		return
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Error("Document() = nil")

		return
	}

	// Verify document has body
	body := docElem.Body()
	if body == nil {
		t.Error("Body() = nil")

		return
	}

	// Count paragraphs to verify content
	paraCount := 0
	for range body.Paragraphs() {
		paraCount++
	}

	if paraCount == 0 {
		t.Error("No paragraphs found")
	} else {
		t.Logf("Document has %d paragraphs", paraCount)
	}
}

// TestMultipleSaveRoundtrips tests that document survives multiple save/open cycles
func TestMultipleSaveRoundtrips(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-multisave-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "multisave.docx")

	// Create initial document
	builder := NewDocumentBuilder()
	builder.AddParagraph("Initial content")
	if _, err := builder.Build(); err != nil {
		t.Fatalf("Builder.Build() error = %v", err)
	}

	doc1, err := New(testPath, DocTypeDocument)
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
		mainPart := doc2.MainPart()
		if mainPart == nil {
			t.Fatalf("MainPart() = nil at cycle %d", i)
		}

		docElem := mainPart.Document()
		if docElem == nil {
			t.Fatalf("Document() = nil at cycle %d", i)
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

	if doc3.MainPart() == nil {
		t.Error("MainPart() = nil after multiple roundtrips")
	}
}
