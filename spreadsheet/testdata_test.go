package spreadsheet

import (
	"os"
	"path/filepath"
	"testing"
)

// TestTestDataFilesExist verifies that all test data files are accessible
func TestTestDataFilesExist(t *testing.T) {
	testFiles := []struct {
		name        string
		description string
	}{
		{
			"basicspreadsheet.xlsx",
			"Basic spreadsheet for simple tests",
		},
		{
			"Comments.xlsx",
			"Spreadsheet with comments",
		},
		{
			"Complex01.xlsx",
			"Complex spreadsheet with advanced features",
		},
		{
			"excel14.xlsx",
			"Excel 2014 format features",
		},
		{"extlst.xlsx", "Extension list testing"},
		{
			"malformed_uri_long.xlsx",
			"Long malformed URI handling",
		},
		{
			"malformed_uri.xlsx",
			"Malformed URI handling",
		},
		{
			"MCExecl.xlsx",
			"Markup compatibility Excel",
		},
		{
			"missingcalcchainpart.xlsx",
			"Missing calculation chain part",
		},
		{
			"Revision_NameCommentChange.xlsx",
			"Revision tracking",
		},
		{
			"Spreadsheet.xlsx",
			"Standard spreadsheet file",
		},
		{
			"Spreadsheet.xltx",
			"Excel template file",
		},
		{
			"vmldrawingroot.xlsx",
			"VML drawing root testing",
		},
		{
			"Youtube.xlsx",
			"Media reference testing",
		},
	}

	for _, tf := range testFiles {
		t.Run(tf.name, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				tf.name,
			)
			info, err := os.Stat(path)
			if err != nil {
				t.Fatalf(
					"File %s does not exist: %v",
					tf.name,
					err,
				)
			}
			if info.IsDir() {
				t.Fatalf(
					"%s is a directory, expected a file",
					tf.name,
				)
			}
			if info.Size() == 0 {
				t.Fatalf("%s is empty", tf.name)
			}
			t.Logf(
				"✓ %s (%d bytes) - %s",
				tf.name,
				info.Size(),
				tf.description,
			)
		})
	}
}

// TestTestDataFileCount verifies the expected number of test files
func TestTestDataFileCount(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf(
			"Failed to read testdata directory: %v",
			err,
		)
	}

	var xlsxCount, xltxCount, otherCount int
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		switch ext {
		case ".xlsx":
			xlsxCount++
		case ".xltx":
			xltxCount++
		case ".md":
			// Documentation files, skip
		default:
			if ext != "" {
				otherCount++
			}
		}
	}

	expectedXlsx := 13
	expectedXltx := 1

	if xlsxCount != expectedXlsx {
		t.Errorf(
			"Expected %d .xlsx files, found %d",
			expectedXlsx,
			xlsxCount,
		)
	}
	if xltxCount != expectedXltx {
		t.Errorf(
			"Expected %d .xltx files, found %d",
			expectedXltx,
			xltxCount,
		)
	}
	if otherCount > 0 {
		t.Logf(
			"Warning: Found %d files with unexpected extensions",
			otherCount,
		)
	}

	t.Logf(
		"Test data summary: %d .xlsx files, %d .xltx files",
		xlsxCount,
		xltxCount,
	)
}

// TestTestDataCategories verifies files can be categorized correctly
func TestTestDataCategories(t *testing.T) {
	categories := map[string][]string{
		"templates": {
			"Spreadsheet.xltx",
		},
		"basic": {
			"basicspreadsheet.xlsx",
			"Spreadsheet.xlsx",
			"MCExecl.xlsx",
		},
		"complex": {
			"Complex01.xlsx",
			"excel14.xlsx",
		},
		"comments_revisions": {
			"Comments.xlsx",
			"Revision_NameCommentChange.xlsx",
		},
		"calculations": {
			"missingcalcchainpart.xlsx",
		},
		"extensions": {
			"extlst.xlsx",
		},
		"media": {
			"Youtube.xlsx",
		},
		"vml_drawing": {
			"vmldrawingroot.xlsx",
		},
		"error_handling": {
			"malformed_uri.xlsx",
			"malformed_uri_long.xlsx",
		},
	}

	for category, files := range categories {
		t.Run(category, func(t *testing.T) {
			for _, file := range files {
				path := filepath.Join(
					"testdata",
					file,
				)
				if _, err := os.Stat(path); err != nil {
					t.Errorf(
						"File %s in category %s does not exist: %v",
						file,
						category,
						err,
					)
				}
			}
			t.Logf(
				"Category '%s' has %d files",
				category,
				len(files),
			)
		})
	}
}

// TestTestDataReadable verifies that all test files can be opened
func TestTestDataReadable(t *testing.T) {
	testFiles := []string{
		"basicspreadsheet.xlsx",
		"Comments.xlsx",
		"Complex01.xlsx",
		"excel14.xlsx",
		"extlst.xlsx",
		"malformed_uri_long.xlsx",
		"malformed_uri.xlsx",
		"MCExecl.xlsx",
		"missingcalcchainpart.xlsx",
		"Revision_NameCommentChange.xlsx",
		"Spreadsheet.xlsx",
		"Spreadsheet.xltx",
		"vmldrawingroot.xlsx",
		"Youtube.xlsx",
	}

	for _, file := range testFiles {
		t.Run(file, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				file,
			)
			f, err := os.Open(path)
			if err != nil {
				t.Fatalf(
					"Cannot open %s: %v",
					file,
					err,
				)
			}
			defer func() { _ = f.Close() }()

			// Read first few bytes to verify it's a valid ZIP file (Office Open XML format)
			header := make([]byte, 4)
			n, err := f.Read(header)
			if err != nil || n != 4 {
				t.Fatalf(
					"Cannot read header from %s: %v",
					file,
					err,
				)
			}

			// Check for ZIP magic number (PK\x03\x04)
			if header[0] != 0x50 ||
				header[1] != 0x4B ||
				header[2] != 0x03 ||
				header[3] != 0x04 {
				t.Errorf(
					"%s does not appear to be a valid Office Open XML file (invalid ZIP header)",
					file,
				)
			}

			t.Logf(
				"✓ %s is readable and has valid ZIP header",
				file,
			)
		})
	}
}

// TestTestDataDocumentation verifies that documentation files exist
func TestTestDataDocumentation(t *testing.T) {
	docFiles := []string{
		"README.md",
		"TEST_FILES_GUIDE.md",
	}

	for _, file := range docFiles {
		t.Run(file, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				file,
			)
			info, err := os.Stat(path)
			if err != nil {
				t.Fatalf(
					"Documentation file %s does not exist: %v",
					file,
					err,
				)
			}
			if info.Size() == 0 {
				t.Fatalf("%s is empty", file)
			}
			t.Logf(
				"✓ %s exists (%d bytes)",
				file,
				info.Size(),
			)
		})
	}
}
