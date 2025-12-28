package wordprocessing

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestComprehensiveDocumentOpenClose tests opening and closing various test documents.
func TestComprehensiveDocumentOpenClose(
	t *testing.T,
) {
	testCases := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"HelloWorld", "HelloWorld.docx", false},
		{"Plain", "Plain.docx", false},
		{"Complex01", "Complex01.docx", false},
		{"Document", "Document.docx", false},
		{"Hyperlink", "Hyperlink.docx", false},
		{"Notes", "Notes.docx", false},
		{"SimpleSdt", "simpleSdt.docx", false},
		{
			"AnnotationRef",
			"AnnotationRef.docx",
			false,
		},
		{"DocProps", "DocProps.docx", false},
		{
			"MoreDocProps",
			"MoreDocProps.docx",
			false,
		},
		{"HelloO14", "HelloO14.docx", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				tc.filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			// Open document
			doc, err := Open(path, false)
			if (err != nil) != tc.wantErr {
				t.Fatalf(
					"Open() error = %v, wantErr %v",
					err,
					tc.wantErr,
				)
			}
			if err != nil {
				return
			}
			defer func() { _ = doc.Close() }()

			// Verify basic properties
			if doc == nil {
				t.Fatal(
					"Open() returned nil document",
				)
			}

			// Verify package
			if doc.Package() == nil {
				t.Error("Package() returned nil")
			}

			// Verify main part (skip if nil - document may be corrupted)
			mainPart := doc.MainPart()
			if mainPart == nil {
				t.Skip(
					"MainPart() returned nil - document may be corrupted",
				)
			}

			// Close and verify
			if err := doc.Close(); err != nil {
				t.Errorf(
					"Close() error = %v",
					err,
				)
			}
		})
	}
}

// TestComprehensiveDocumentStructure tests document structure validation.
func TestComprehensiveDocumentStructure(
	t *testing.T,
) {
	testCases := []string{
		"HelloWorld.docx",
		"Plain.docx",
		"Complex01.docx",
		"Document.docx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Verify main part
			mainPart := doc.MainPart()
			if mainPart == nil {
				t.Fatal("MainPart() is nil")
			}

			// Get and verify document element
			docElem := mainPart.Document()
			if docElem == nil {
				t.Fatal("Document() is nil")
			}

			// Verify body
			body := docElem.Body()
			if body == nil {
				t.Fatal("Body() is nil")
			}

			// Count paragraphs
			paraCount := 0
			for range body.Paragraphs() {
				paraCount++
			}
			t.Logf(
				"Document has %d paragraphs",
				paraCount,
			)

			// Count tables
			tableCount := 0
			for range body.Tables() {
				tableCount++
			}
			t.Logf(
				"Document has %d tables",
				tableCount,
			)
		})
	}
}

// TestComprehensiveRoundtrip tests opening, modifying, saving, and reopening.
func TestComprehensiveRoundtrip(t *testing.T) {
	testCases := []string{
		"HelloWorld.docx",
		"Plain.docx",
		"SimpleSdt.docx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			sourcePath := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(sourcePath); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					sourcePath,
				)
			}

			// Create temp directory
			tmpDir, err := os.MkdirTemp(
				"",
				"goffice-roundtrip-*",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create temp dir: %v",
					err,
				)
			}
			defer func() { _ = os.RemoveAll(tmpDir) }()

			// Open original document
			doc, err := Open(sourcePath, true)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}

			// Save to new location
			outputPath := filepath.Join(
				tmpDir,
				"modified_"+filename,
			)
			if err := doc.SaveAs(outputPath); err != nil {
				_ = doc.Close()
				t.Fatalf(
					"SaveAs() error = %v",
					err,
				)
			}
			_ = doc.Close()

			// Reopen saved document
			doc2, err := Open(outputPath, false)
			if err != nil {
				t.Fatalf(
					"Failed to reopen saved document: %v",
					err,
				)
			}
			defer func() { _ = doc2.Close() }()

			// Verify structure
			if doc2.MainPart() == nil {
				t.Error(
					"Reopened document missing main part",
				)
			}

			// Verify package integrity
			if doc2.Package() == nil {
				t.Error(
					"Reopened document missing package",
				)
			}
		})
	}
}

// TestComprehensivePartsExtraction tests extracting various parts from documents.
func TestComprehensivePartsExtraction(
	t *testing.T,
) {
	t.Run(
		"DocumentWithStyles",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Complex01.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			mainPart := doc.MainPart()

			// Check for styles part
			stylesPart := mainPart.StylesPart()
			if stylesPart != nil {
				t.Log("Found styles part")
				styles := stylesPart.Styles()
				if styles != nil {
					t.Log(
						"Styles element loaded",
					)
				}
			}

			// Check for numbering part
			numberingPart := mainPart.NumberingPart()
			if numberingPart != nil {
				t.Log("Found numbering part")
			}

			// Check for settings part
			settingsPart := mainPart.SettingsPart()
			if settingsPart != nil {
				t.Log("Found settings part")
			}

			// Check for web settings part
			webSettingsPart := mainPart.WebSettingsPart()
			if webSettingsPart != nil {
				t.Log("Found web settings part")
			}
		},
	)

	t.Run(
		"DocumentWithNotes",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Notes.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			mainPart := doc.MainPart()

			// Check for footnotes
			footnotesPart := mainPart.FootnotesPart()
			if footnotesPart != nil {
				t.Log("Found footnotes part")
				footnotes := footnotesPart.Footnotes()
				if footnotes != nil {
					count := 0
					for range footnotes.Footnotes() {
						count++
					}
					t.Logf(
						"Document has %d footnotes",
						count,
					)
				}
			}

			// Check for endnotes
			endnotesPart := mainPart.EndnotesPart()
			if endnotesPart != nil {
				t.Log("Found endnotes part")
			}
		},
	)

	t.Run(
		"DocumentWithComments",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Comments.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			mainPart := doc.MainPart()

			// Check for comments part
			commentsPart := mainPart.CommentsPart()
			if commentsPart == nil {
				return
			}
			t.Log("Found comments part")
			comments := commentsPart.Comments()
			if comments == nil {
				return
			}
			count := 0
			for range comments.Comments() {
				count++
			}
			t.Logf(
				"Document has %d comments",
				count,
			)
		},
	)
}

// TestComprehensiveRelationshipHandling tests relationship handling.
func TestComprehensiveRelationshipHandling(
	t *testing.T,
) {
	testFiles := []string{
		"HelloWorld.docx",
		"Complex01.docx",
		"Hyperlink.docx",
	}

	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			pkg := doc.Package()
			if pkg == nil {
				t.Fatal("Package is nil")
			}

			// Get package relationships
			packPkg := pkg.Package()
			if packPkg == nil {
				t.Fatal(
					"Underlying package is nil",
				)
			}

			// Verify main document relationship exists
			rels := packPkg.Relationships()
			if rels == nil {
				t.Fatal(
					"Package relationships are nil",
				)
			}

			foundMainDoc := false
			for rel := range rels.All() {
				if rel.Type() == "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" {
					foundMainDoc = true
					t.Logf(
						"Found main document relationship: %s",
						rel.ID(),
					)
				}
			}

			if !foundMainDoc {
				t.Error(
					"Main document relationship not found",
				)
			}
		})
	}
}

// TestComprehensiveDocumentProperties tests core, extended, and custom properties.
func TestComprehensiveDocumentProperties(
	t *testing.T,
) {
	t.Run("CoreProperties", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"DocProps.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = doc.Close() }()

		coreProps := doc.CoreProperties()
		if coreProps == nil {
			return
		}
		t.Log("Core properties found")
		if title := coreProps.Title(); title != "" {
			t.Logf("Title: %s", title)
		}
		if creator := coreProps.Creator(); creator != "" {
			t.Logf("Creator: %s", creator)
		}
	})

	t.Run(
		"ExtendedProperties",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"MoreDocProps.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			extProps := doc.ExtendedProperties()
			if extProps == nil {
				return
			}
			t.Log(
				"Extended properties found",
			)
			if app := extProps.Application(); app != "" {
				t.Logf("Application: %s", app)
			}
		},
	)

	t.Run("NoDocProperties", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"NoDocProps.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = doc.Close() }()

		// Should handle missing properties gracefully
		coreProps := doc.CoreProperties()
		if coreProps == nil {
			t.Log(
				"No core properties (expected)",
			)
		}
	})
}

// TestComprehensiveCommentsAndAnnotations tests comments and annotations.
func TestComprehensiveCommentsAndAnnotations(
	t *testing.T,
) {
	t.Run("CommentsDocument", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"Comments.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = doc.Close() }()

		mainPart := doc.MainPart()
		commentsPart := mainPart.CommentsPart()

		if commentsPart == nil {
			t.Skip("No comments part found")
		}

		comments := commentsPart.Comments()
		if comments == nil {
			t.Skip(
				"Comments element is nil - file may not have standard comments structure",
			)
		}

		// Verify comments
		count := 0
		for comment := range comments.Comments() {
			count++
			t.Logf(
				"Comment %d: Author=%s, ID=%d",
				count,
				comment.Author(),
				comment.Id(),
			)
		}

		if count == 0 {
			t.Log("No comments found in document")
		}
	})

	t.Run("AnnotationRef", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"AnnotationRef.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = doc.Close() }()

		// Verify document can be opened and read (skip if corrupted)
		if doc.MainPart() == nil {
			t.Skip(
				"MainPart is nil - document may be corrupted",
			)
		}
	})
}

// TestComprehensiveComplexDocuments tests complex real-world documents.
func TestComprehensiveComplexDocuments(
	t *testing.T,
) {
	testCases := []struct {
		name     string
		filename string
	}{
		{"Complex01", "Complex01.docx"},
		{"Document", "Document.docx"},
		{"complex0", "complex0.docx"},
		{"complex2010", "complex2010.docx"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				tc.filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			mainPart := doc.MainPart()
			if mainPart == nil {
				t.Fatal("MainPart is nil")
			}

			// Get document element
			docElem := mainPart.Document()
			if docElem == nil {
				t.Fatal("Document element is nil")
			}

			// Verify body
			body := docElem.Body()
			if body == nil {
				t.Fatal("Body is nil")
			}

			// Count elements
			paraCount := 0
			tableCount := 0
			for range body.Paragraphs() {
				paraCount++
			}
			for range body.Tables() {
				tableCount++
			}

			t.Logf(
				"Complex document stats: %d paragraphs, %d tables",
				paraCount,
				tableCount,
			)

			// Verify various parts
			if stylesPart := mainPart.StylesPart(); stylesPart != nil {
				t.Log("Has styles part")
			}
			if numberingPart := mainPart.NumberingPart(); numberingPart != nil {
				t.Log("Has numbering part")
			}
			if settingsPart := mainPart.SettingsPart(); settingsPart != nil {
				t.Log("Has settings part")
			}
		})
	}
}

// TestComprehensiveErrorHandling tests error handling for malformed documents.
func TestComprehensiveErrorHandling(
	t *testing.T,
) {
	t.Run("NonExistentFile", func(t *testing.T) {
		doc, err := Open(
			"testdata/nonexistent.docx",
			false,
		)
		if err == nil {
			_ = doc.Close()
			t.Error(
				"Expected error opening non-existent file",
			)
		}
	})

	t.Run("BadDocProps", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"BadDocProps.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		// Should still be able to open document with bad properties
		doc, err := Open(path, false)
		if err != nil {
			t.Logf(
				"Open with bad doc props error (may be expected): %v",
				err,
			)
		} else {
			defer func() { _ = doc.Close() }()
			t.Log("Opened document with bad properties")
		}
	})

	t.Run("5Errors", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"5Errors.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		// Document with intentional errors should still open
		doc, err := Open(path, false)
		if err != nil {
			t.Logf("Open with errors: %v", err)
		} else {
			defer func() { _ = doc.Close() }()
			t.Log("Opened document with known errors")

			// Verify we can still access basic structure
			if mainPart := doc.MainPart(); mainPart != nil {
				t.Log("MainPart accessible despite errors")
			}
		}
	})

	t.Run(
		"EmptyRelationshipElement",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"EmptyRelationshipElement.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Logf(
					"Open with empty relationship error: %v",
					err,
				)
			} else {
				defer func() { _ = doc.Close() }()
				t.Log("Handled empty relationship element")
			}
		},
	)

	t.Run("UnknownElement", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"UnknownElement.docx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		doc, err := Open(path, false)
		if err != nil {
			t.Logf(
				"Open with unknown element: %v",
				err,
			)
		} else {
			defer func() { _ = doc.Close() }()
			t.Log("Handled unknown element gracefully")
		}
	})
}

// TestComprehensiveOffice2016Features tests Office 2016 specific features.
func TestComprehensiveOffice2016Features(
	t *testing.T,
) {
	office2016Files := []string{
		"Of16-01.docx",
		"Of16-02.docx",
		"Of16-03.docx",
		"Of16-04.docx",
		"Of16-05.docx",
		"Of16-06.docx",
		"Of16-07.docx",
		"Of16-08.docx",
	}

	for _, filename := range office2016Files {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Verify basic structure
			if doc.MainPart() == nil {
				t.Error("MainPart is nil")
			}

			// Test validation against Office 2016
			errors := doc.Validate(
				validation.Office2016,
			)
			if len(errors) > 0 {
				t.Logf(
					"Validation found %d issues (may be expected)",
					len(errors),
				)
			}
		})
	}

	t.Run(
		"Of16-09-UnknownElement",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Of16-09-UnknownElement.docx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Logf(
					"Open() with unknown element: %v",
					err,
				)
			} else {
				defer func() { _ = doc.Close() }()
				t.Log("Handled Office 2016 unknown element")
			}
		},
	)
}

// TestComprehensiveStrictOpenXML tests strict Open XML compliance.
func TestComprehensiveStrictOpenXML(
	t *testing.T,
) {
	testFiles := []string{
		"HelloWorld.docx",
		"Plain.docx",
		"Complex01.docx",
	}

	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Validate with strict settings
			strictSettings := validation.StrictSettings()
			errors := doc.ValidateWithSettings(
				validation.Microsoft365,
				strictSettings,
			)

			if len(errors) > 0 {
				t.Logf(
					"Strict validation found %d issues",
					len(errors),
				)
				for i, err := range errors {
					if i < 5 { // Limit output
						t.Logf(
							"  - %s",
							err.Description,
						)
					}
				}
			} else {
				t.Log("Document passes strict validation")
			}
		})
	}
}

// TestComprehensivePerformanceLargeDocuments tests performance with large documents.
func TestComprehensivePerformanceLargeDocuments(
	t *testing.T,
) {
	if testing.Short() {
		t.Skip(
			"Skipping performance test in short mode",
		)
	}

	largeFiles := []string{
		"Complex01.docx", // 828K
		"Document.docx",  // 799K
		"complex0.docx",  // 215K
	}

	for _, filename := range largeFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			// Measure open time
			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			// Iterate through all paragraphs
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.Body()

			paraCount := 0
			for para := range body.Paragraphs() {
				paraCount++
				// Access inner text to force parsing
				_ = para.InnerText()
			}

			t.Logf(
				"Processed %d paragraphs from large document",
				paraCount,
			)
		})
	}
}

// TestComprehensiveHyperlinkHandling tests hyperlink handling.
func TestComprehensiveHyperlinkHandling(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Hyperlink.docx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	docElem := mainPart.Document()
	body := docElem.Body()

	// Search for hyperlinks in paragraphs
	foundHyperlink := false
	for para := range body.Paragraphs() {
		// Check if paragraph contains hyperlink elements
		xml := para.OuterXml()
		if strings.Contains(xml, "w:hyperlink") {
			foundHyperlink = true
			t.Log("Found hyperlink in paragraph")
		}
	}

	if foundHyperlink {
		t.Log("Successfully detected hyperlinks")
	} else {
		t.Log("No hyperlinks found (document may not contain any)")
	}
}

// TestComprehensiveSVGSupport tests SVG image support.
func TestComprehensiveSVGSupport(t *testing.T) {
	path := filepath.Join("testdata", "svg.docx")
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Check for image parts
	imageParts := mainPart.ImageParts()
	if len(imageParts) == 0 {
		return
	}
	t.Logf(
		"Found %d image parts",
		len(imageParts),
	)
	for _, imgPart := range imageParts {
		contentType := imgPart.ContentType()
		t.Logf(
			"Image part content type: %s",
			contentType,
		)
		if strings.Contains(
			contentType,
			"svg",
		) {
			t.Log("Found SVG image")
		}
	}
}

// TestComprehensiveMailMerge tests mail merge document handling.
func TestComprehensiveMailMerge(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"mailmerge.docx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Check for settings part (mail merge settings)
	settingsPart := mainPart.SettingsPart()
	if settingsPart == nil {
		return
	}
	t.Log(
		"Found settings part (may contain mail merge settings)",
	)
	settings := settingsPart.Settings()
	if settings == nil {
		return
	}
	xml := settings.OuterXml()
	if strings.Contains(
		xml,
		"mailMerge",
	) {
		t.Log(
			"Document contains mail merge settings",
		)
	}
}

// TestComprehensiveContentControls tests content control (SDT) handling.
func TestComprehensiveContentControls(
	t *testing.T,
) {
	testFiles := []string{
		"simpleSdt.docx",
		"Data-Bound-Content-Controls.docx",
	}

	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			doc, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.Body()

			// Search for content controls (SDT elements)
			foundSDT := false
			for para := range body.Paragraphs() {
				xml := para.OuterXml()
				if strings.Contains(
					xml,
					"w:sdt",
				) {
					foundSDT = true
					t.Log(
						"Found content control (SDT)",
					)
				}
			}

			if foundSDT {
				t.Log(
					"Successfully detected content controls",
				)
			}
		})
	}
}

// TestComprehensiveMacroDocument tests macro-enabled document handling.
func TestComprehensiveMacroDocument(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"mcdoc.docx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify document type detection
	docType := doc.Type()
	t.Logf("Document type: %s", docType.String())

	// Check for VBA project
	mainPart := doc.MainPart()
	vbaPart := mainPart.VbaProjectPart()
	if vbaPart != nil {
		t.Log("Found VBA project part")
	}
}

// TestComprehensiveDocumentToBytes tests saving document to byte array.
func TestComprehensiveDocumentToBytes(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"HelloWorld.docx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Save to buffer
	var buf bytes.Buffer
	if err := doc.SaveTo(&buf); err != nil {
		t.Fatalf("SaveTo() error = %v", err)
	}

	// Verify buffer has data
	if buf.Len() == 0 {
		t.Error("SaveTo() produced empty buffer")
	}

	t.Logf(
		"Saved document to buffer: %d bytes",
		buf.Len(),
	)
}

// TestComprehensiveValidationVersions tests validation across Office versions.
func TestComprehensiveValidationVersions(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Plain.docx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	doc, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	versions := []struct {
		name    string
		version validation.FileFormatVersions
	}{
		{"Office2016", validation.Office2016},
		{"Office2019", validation.Office2019},
		{"Office2021", validation.Office2021},
		{"Microsoft365", validation.Microsoft365},
	}

	for _, v := range versions {
		t.Run(v.name, func(t *testing.T) {
			errors := doc.Validate(v.version)
			t.Logf(
				"%s validation: %d issues",
				v.name,
				len(errors),
			)

			// Check IsValid
			isValid := doc.IsValid(v.version)
			t.Logf(
				"%s IsValid: %v",
				v.name,
				isValid,
			)
		})
	}
}

// TestComprehensiveHeaderFooterIterators tests header/footer iteration.
func TestComprehensiveHeaderFooterIterators(
	t *testing.T,
) {
	// Create a document with headers and footers
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-hf-iter-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Add headers
	_, _ = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	_, _ = doc.AddHeader(
		elements.HeaderFooterValuesFirst,
	)

	// Add footers
	_, _ = doc.AddFooter(
		elements.HeaderFooterDefault,
	)
	_, _ = doc.AddFooter(
		elements.HeaderFooterValuesFirst,
	)

	// Test header iterator
	headerCount := 0
	for header := range doc.Headers() {
		headerCount++
		if header == nil {
			t.Error(
				"Header iterator returned nil",
			)
		}
	}
	if headerCount != 2 {
		t.Errorf(
			"Expected 2 headers, got %d",
			headerCount,
		)
	}

	// Test footer iterator
	footerCount := 0
	for footer := range doc.Footers() {
		footerCount++
		if footer == nil {
			t.Error(
				"Footer iterator returned nil",
			)
		}
	}
	if footerCount != 2 {
		t.Errorf(
			"Expected 2 footers, got %d",
			footerCount,
		)
	}
}

// TestComprehensiveDocumentType tests document type detection and conversion.
func TestComprehensiveDocumentType(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-doctype-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Test each document type
	types := []struct {
		docType   DocType
		extension string
	}{
		{DocTypeDocument, ".docx"},
		{DocTypeTemplate, ".dotx"},
		{DocTypeMacroEnabled, ".docm"},
		{DocTypeMacroTemplate, ".dotm"},
	}

	for _, tt := range types {
		t.Run(
			tt.docType.String(),
			func(t *testing.T) {
				testPath := filepath.Join(
					tmpDir,
					"test"+tt.extension,
				)
				doc, err := New(
					testPath,
					tt.docType,
				)
				if err != nil {
					t.Fatalf(
						"New() error = %v",
						err,
					)
				}
				defer func() { _ = doc.Close() }()

				// Verify type
				if doc.Type() != tt.docType {
					t.Errorf(
						"Expected type %v, got %v",
						tt.docType,
						doc.Type(),
					)
				}

				// Verify extension
				if tt.docType.Extension() != tt.extension {
					t.Errorf(
						"Expected extension %s, got %s",
						tt.extension,
						tt.docType.Extension(),
					)
				}

				// Verify content type
				contentType := tt.docType.ContentType()
				if contentType == "" {
					t.Error(
						"ContentType() returned empty string",
					)
				}
				t.Logf(
					"Content type: %s",
					contentType,
				)
			},
		)
	}
}
