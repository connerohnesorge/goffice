package presentation

import (
	"os"
	"path/filepath"
	"testing"
)

// TestTestDataFilesExist verifies that all expected test data files are present
func TestTestDataFilesExist(t *testing.T) {
	testFiles := []string{
		"Presentation.pptx",
		"Presentation.potx",
		"animation.pptx",
		"3dtestdash.pptx",
		"3dtestdot.pptx",
		"Algn_tab_TabAlignment.pptx",
		"mediareference.pptx",
		"encrypted_pptx.pptx",
		"Of16-01.pptx",
		"Of16-02.pptx",
		"Of16-03.pptx",
		"mcppt.pptx",
		"autosave.pptx",
		"o09_Performance_typical.pptx",
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

	pptxCount := 0
	potxCount := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		switch ext {
		case ".pptx":
			pptxCount++
		case ".potx":
			potxCount++
		}
	}

	t.Logf("Found %d .pptx files", pptxCount)
	t.Logf("Found %d .potx files", potxCount)

	if pptxCount != 13 {
		t.Errorf(
			"Expected 13 .pptx files, found %d",
			pptxCount,
		)
	}
	if potxCount != 1 {
		t.Errorf(
			"Expected 1 .potx file, found %d",
			potxCount,
		)
	}
}

// TestTestDataFilesAccessible verifies that test files can be opened
func TestTestDataFilesAccessible(t *testing.T) {
	// Test a few representative files to ensure they're readable
	testFiles := []string{
		"Presentation.pptx",
		"Presentation.potx",
		"mcppt.pptx",
	}

	for _, file := range testFiles {
		path := filepath.Join("testdata", file)
		f, err := os.Open(path)
		if err != nil {
			t.Errorf(
				"Failed to open test file %s: %v",
				path,
				err,
			)

			continue
		}

		// Verify file has content
		stat, err := f.Stat()
		if err != nil {
			_ = f.Close()
			t.Errorf(
				"Failed to stat test file %s: %v",
				path,
				err,
			)

			continue
		}

		if stat.Size() == 0 {
			_ = f.Close()
			t.Errorf(
				"Test file %s is empty",
				path,
			)
		} else {
			_ = f.Close()
			t.Logf("✓ File %s is accessible and has size: %d bytes", file, stat.Size())
		}
	}
}
