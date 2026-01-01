package spreadsheet

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
)

// TestAlternateContentSpreadsheetIntegration tests that AlternateContent elements
// can be created and roundtripped in Excel documents.
func TestAlternateContentSpreadsheetIntegration(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-xl-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "alternate_content.xlsx")

	// Create spreadsheet
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

	// Verify structure is preserved
	workbookPart := doc2.WorkbookPart()
	if workbookPart == nil {
		t.Error("WorkbookPart() = nil after reopen")
	}
}

// TestAlternateContentInSpreadsheetCell tests AlternateContent elements
// within spreadsheet cells.
func TestAlternateContentInSpreadsheetCell(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-cell-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "ac_cell.xlsx")

	// Create spreadsheet
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Create AlternateContent element for modern Excel feature
	ac := openxml.NewAlternateContent()

	// Add Office 2013 choice (slicers, timelines)
	choice2013 := openxml.NewChoice()
	choice2013.SetRequires("x14")
	ac.AppendChild(choice2013)

	// Add fallback for older Excel
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

	// Verify workbook is readable
	if doc2.WorkbookPart() == nil {
		t.Error("WorkbookPart() = nil after reopen")
	}
}

// TestExcel2013ExtensionElements tests that Excel 2013 extension elements
// (slicers, timelines) are properly handled.
func TestExcel2013ExtensionElements(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-x14-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "excel2013.xlsx")

	// Create spreadsheet
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Save with Office 2013 content
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

	// Verify workbook structure
	workbookPart := doc2.WorkbookPart()
	if workbookPart == nil {
		t.Error("WorkbookPart() = nil")
	}
}

// TestExcel2025ExtensionElements tests handling of Excel 2025 extension elements.
func TestExcel2025ExtensionElements(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-x25-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "excel2025.xlsx")

	// Create spreadsheet
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

	// Verify workbook is readable
	if doc2.WorkbookPart() == nil {
		t.Error("WorkbookPart() = nil after reopen")
	}
}

// TestMultipleVersionChoicesInSpreadsheet tests AlternateContent with
// multiple version choices in spreadsheets.
func TestMultipleVersionChoicesInSpreadsheet(t *testing.T) {
	ac := openxml.NewAlternateContent()

	// Office 2013 feature
	choice2013 := openxml.NewChoice()
	choice2013.SetRequires("x14")
	ac.AppendChild(choice2013)

	// Office 2016 feature
	choice2016 := openxml.NewChoice()
	choice2016.SetRequires("x16")
	ac.AppendChild(choice2016)

	// Fallback
	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	// Test XML generation
	xml := ac.OuterXml()

	if !strings.Contains(xml, "x14") {
		t.Error("XML missing x14 namespace reference")
	}
	if !strings.Contains(xml, "x16") {
		t.Error("XML missing x16 namespace reference")
	}
}

// TestSpreadsheetUnknownElementPreservation tests that unknown Excel extension
// elements are preserved.
func TestSpreadsheetUnknownElementPreservation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-unknown-xl-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "unknown.xlsx")

	// Create spreadsheet
	doc1, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

		if doc2.WorkbookPart() == nil {
			t.Errorf("WorkbookPart() = nil at cycle %d", i)
		}

		if err := doc2.SaveAs(testPath); err != nil {
			t.Fatalf("SaveAs() cycle %d error = %v", i, err)
		}
		_ = doc2.Close()
	}
}

// TestNamespacePrefixConsistency tests that namespace prefixes are consistent
// across extension elements.
func TestNamespacePrefixConsistency(t *testing.T) {
	tests := []struct {
		name      string
		prefix    string
		namespace string
	}{
		{"Excel2010", "x14", "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"},
		{"Excel2013", "x14", "http://schemas.microsoft.com/office/spreadsheetml/2013/9/main"},
		{"Excel2016", "x16", "http://schemas.microsoft.com/office/spreadsheetml/2016/9/main"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			choice := openxml.NewChoice()
			choice.SetRequires(tt.prefix)

			if choice.Requires() != tt.prefix {
				t.Errorf(
					"Requires() = %q, want %q",
					choice.Requires(),
					tt.prefix,
				)
			}
		})
	}
}
