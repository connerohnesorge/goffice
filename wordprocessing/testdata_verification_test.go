package wordprocessing

import (
	"os"
	"path/filepath"
	"testing"
)

// TestTestDataFilesExist verifies that all expected test data files are present
func TestTestDataFilesExist(t *testing.T) {
	testFiles := []string{
		"HelloWorld.docx",
		"Plain.docx",
		"Document.docx",
		"Document.dotx",
		"Complex01.docx",
		"Comments.docx",
		"Hyperlink.docx",
	}

	for _, file := range testFiles {
		path := filepath.Join("testdata", file)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Errorf(
				"Test file %s does not exist",
				path,
			)
		} else {
			t.Logf("✓ Found test file: %s", path)
		}
	}
}

// TestTestDataDirectory verifies the testdata directory structure
func TestTestDataDirectory(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf(
			"Failed to read testdata directory: %v",
			err,
		)
	}

	docxCount := 0
	dotxCount := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		switch ext {
		case ".docx":
			docxCount++
		case ".dotx":
			dotxCount++
		}
	}

	t.Logf("Found %d .docx files", docxCount)
	t.Logf("Found %d .dotx files", dotxCount)

	if docxCount != 38 {
		t.Errorf(
			"Expected 38 .docx files, found %d",
			docxCount,
		)
	}
	if dotxCount != 1 {
		t.Errorf(
			"Expected 1 .dotx file, found %d",
			dotxCount,
		)
	}
}
