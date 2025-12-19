package testutil

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateMinimalDocx(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	docxPath := filepath.Join(tmpDir, "test.docx")

	// Create the minimal docx
	err := CreateMinimalDocx(docxPath)
	if err != nil {
		t.Fatalf("CreateMinimalDocx failed: %v", err)
	}

	// Verify the file exists
	if _, err := os.Stat(docxPath); os.IsNotExist(err) {
		t.Fatal("docx file was not created")
	}

	// Open and verify the ZIP structure
	r, err := zip.OpenReader(docxPath)
	if err != nil {
		t.Fatalf("Failed to open docx as ZIP: %v", err)
	}
	defer r.Close()

	// Check that required files exist
	requiredFiles := map[string]bool{
		"[Content_Types].xml":          false,
		"_rels/.rels":                  false,
		"word/document.xml":            false,
		"word/_rels/document.xml.rels": false,
	}

	for _, f := range r.File {
		if _, ok := requiredFiles[f.Name]; ok {
			requiredFiles[f.Name] = true
		}
	}

	for name, found := range requiredFiles {
		if !found {
			t.Errorf("Required file %q not found in docx", name)
		}
	}
}

func TestMinimalDocxFilesContents(t *testing.T) {
	// Verify each minimal docx file has content
	for name, content := range MinimalDocxFiles {
		if content == "" {
			t.Errorf("MinimalDocxFiles[%q] is empty", name)
		}
	}
}
