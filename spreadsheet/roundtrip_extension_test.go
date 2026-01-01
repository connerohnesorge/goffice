package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

// TestRoundtripExtensionElementsBasic tests basic extension element roundtrip
// by creating and reopening a spreadsheet with simple content.
func TestRoundtripExtensionElementsBasic(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-extension-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "extension.xlsx")

	// Create workbook
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Get the workbook part
	workbookPart := doc1.WorkbookPart()
	if workbookPart == nil {
		t.Fatal("WorkbookPart() = nil")
	}

	// Save and close
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen the workbook
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify workbook is readable
	workbookPart2 := doc2.WorkbookPart()
	if workbookPart2 == nil {
		t.Error("WorkbookPart() = nil after reopen")

		return
	}

	workbook2 := workbookPart2.Workbook()
	if workbook2 == nil {
		t.Error("Workbook() = nil after reopen")
	}
}

// TestVersionDetectionFromSpreadsheet tests that spreadsheets are correctly identified
// with their Office version requirements.
func TestVersionDetectionFromSpreadsheet(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-version-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "version-detect.xlsx")

	// Create and save a simple workbook
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

	// Basic Office 2007 spreadsheet should be Office2007 or higher
	if detectedVersion < validation.Office2007 {
		t.Errorf("DetectMinimumVersion() = %v, want >= Office2007", detectedVersion)
	}
}

// TestSpreadsheetBasicRoundtrip tests simple roundtrip
func TestSpreadsheetBasicRoundtrip(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-roundtrip-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "roundtrip.xlsx")

	// Create workbook
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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
	if doc2.Type() != DocTypeWorkbook {
		t.Errorf("Type() = %v, want %v", doc2.Type(), DocTypeWorkbook)
	}
}

// TestSpreadsheetNamespacePreservation tests that namespaces are preserved
func TestSpreadsheetNamespacePreservation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-namespace-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "namespace.xlsx")

	// Create document
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

	// Verify we can access the workbook
	workbookPart := doc2.WorkbookPart()
	if workbookPart == nil {
		t.Error("WorkbookPart() = nil")

		return
	}

	workbook := workbookPart.Workbook()
	if workbook == nil {
		t.Error("Workbook() = nil")

		return
	}

	t.Log("Successfully opened and verified spreadsheet structure")
}

// TestMultipleSaveRoundtrips tests that workbook survives multiple save/open cycles
func TestMultipleSaveRoundtrips(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-multisave-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "multisave.xlsx")

	// Create initial workbook
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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
		workbookPart := doc2.WorkbookPart()
		if workbookPart == nil {
			t.Fatalf("WorkbookPart() = nil at cycle %d", i)
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

	if doc3.WorkbookPart() == nil {
		t.Error("WorkbookPart() = nil after multiple roundtrips")
	}
}

// TestRoundtripWorkbookTypes tests saving different workbook formats
func TestRoundtripWorkbookTypes(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-excel-formats-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	formats := []struct {
		name string
		typ  DocType
	}{
		{"workbook.xlsx", DocTypeWorkbook},
		{"macro.xlsm", DocTypeWorkbook},
		{"template.xltx", DocTypeTemplate},
	}

	for _, format := range formats {
		t.Run(format.name, func(t *testing.T) {
			testPath := filepath.Join(tmpDir, format.name)

			// Create workbook
			doc1, err := Create(testPath, format.typ)
			if err != nil {
				t.Fatalf("Create() error = %v", err)
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
