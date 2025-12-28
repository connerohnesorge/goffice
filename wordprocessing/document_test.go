package wordprocessing

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestDocumentValidate tests the Validate method on Document.
func TestDocumentValidate(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-validate-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"validate.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Validate against Office 2016
	errors := doc.Validate(validation.Office2016)

	// A newly created document should have minimal or no validation errors
	if errors != nil && errors.HasErrors() {
		t.Logf(
			"Validation returned %d errors (may be expected for minimal document)",
			len(errors),
		)
	}
}

// TestDocumentValidateWithSettings tests ValidateWithSettings method.
func TestDocumentValidateWithSettings(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-validate-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"validate.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test with fast settings
	fastSettings := validation.FastSettings()
	errors := doc.ValidateWithSettings(
		validation.Microsoft365,
		fastSettings,
	)

	// Just verify it runs without panic
	_ = errors

	// Test with strict settings
	strictSettings := validation.StrictSettings()
	errors = doc.ValidateWithSettings(
		validation.Microsoft365,
		strictSettings,
	)

	// Just verify it runs without panic
	_ = errors
}

// TestDocumentIsValid tests the IsValid convenience method.
func TestDocumentIsValid(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-isvalid-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"isvalid.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test IsValid method
	valid := doc.IsValid(validation.Office2016)

	// A newly created document should typically be valid
	// (or at least the method should not panic)
	t.Logf(
		"Document IsValid(Office2016) = %v",
		valid,
	)
}

// TestDocumentValidateClosedDocument tests validation on a closed document.
func TestDocumentValidateClosedDocument(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-validate-closed-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"closed.docx",
	)

	// Create and close a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	_ = doc.Close()

	// Validate closed document
	errors := doc.Validate(validation.Office2016)

	// Should have at least one error about nil package
	if len(errors) == 0 {
		t.Error(
			"Expected validation error for closed document",
		)
	}

	// Verify the error message indicates the package is nil
	found := false
	for _, e := range errors {
		if e.Description == "Document package is nil or closed" {
			found = true

			break
		}
	}
	if !found {
		t.Error(
			"Expected 'Document package is nil or closed' error",
		)
	}
}

// TestDocumentValidateVersions tests validation against different Office versions.
func TestDocumentValidateVersions(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-versions-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"versions.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test all supported versions
	versions := []validation.FileFormatVersions{
		validation.Office2016,
		validation.Office2019,
		validation.Office2021,
		validation.Microsoft365,
	}

	for _, version := range versions {
		t.Run(
			version.String(),
			func(t *testing.T) {
				errors := doc.Validate(version)
				// Just verify it runs without panic for each version
				t.Logf(
					"Version %s: %d validation messages",
					version.String(),
					len(errors),
				)
			},
		)
	}
}

// TestDocumentValidateWithRealFile tests validation on a real .docx file.
func TestDocumentValidateWithRealFile(
	t *testing.T,
) {
	fixturePath := "../testdata/fixtures/minimal.docx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.docx fixture not found",
		)
	}

	// Open the real .docx file
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open minimal.docx: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Validate against Office 2016
	errors := doc.Validate(validation.Office2016)

	// Log any errors (real documents may have some)
	if len(errors) > 0 {
		t.Logf(
			"Real document validation found %d issues:",
			len(errors),
		)
		for i, e := range errors {
			if i < 5 { // Limit output
				t.Logf("  - %s", e.Error())
			}
		}
	}

	// Also test IsValid
	valid := doc.IsValid(validation.Microsoft365)
	t.Logf(
		"Real document IsValid(Microsoft365) = %v",
		valid,
	)
}

// TestDocumentSaveToWriter_Basic tests basic SaveTo method with io.Writer.
func TestDocumentSaveToWriter_Basic(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-saveto-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document
	testPath := filepath.Join(
		tmpDir,
		"test.docx",
	)
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add some content
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() returned nil")
	}

	// Write to a bytes.Buffer
	var buf []byte
	w := &bufferWriter{buf: &buf}
	if err := doc.SaveTo(w); err != nil {
		t.Fatalf("SaveTo() error = %v", err)
	}

	// Close the document
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Verify the buffer contains valid ZIP data
	if len(buf) == 0 {
		t.Fatal("SaveTo() wrote no data")
	}

	// Try to open the data as a ZIP
	reader := bytes.NewReader(buf)
	zipReader, err := zip.NewReader(
		reader,
		int64(len(buf)),
	)
	if err != nil {
		t.Fatalf(
			"SaveTo() did not produce valid ZIP: %v",
			err,
		)
	}

	// Verify essential files are present
	hasContentTypes := false
	hasDocument := false
	hasRels := false

	for _, f := range zipReader.File {
		switch f.Name {
		case "[Content_Types].xml":
			hasContentTypes = true
		case "word/document.xml":
			hasDocument = true
		case "_rels/.rels":
			hasRels = true
		}
	}

	if !hasContentTypes {
		t.Error(
			"ZIP missing [Content_Types].xml",
		)
	}
	if !hasDocument {
		t.Error(
			"ZIP missing word/document.xml",
		)
	}
	if !hasRels {
		t.Error("ZIP missing _rels/.rels")
	}
}

// TestDocumentSaveToWriter_Reopen tests SaveTo and reopening the document.
func TestDocumentSaveToWriter_Reopen(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-saveto-reopen-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create a document with content
	testPath := filepath.Join(
		tmpDir,
		"test.docx",
	)
	doc, err := New(
		testPath,
		DocTypeDocument,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Write to buffer
	var buf []byte
	w := &bufferWriter{buf: &buf}
	if err := doc.SaveTo(w); err != nil {
		t.Fatalf(
			"SaveTo() error = %v",
			err,
		)
	}

	// Close original document
	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Close() error = %v",
			err,
		)
	}

	// Write buffer to file and reopen
	reopenPath := filepath.Join(
		tmpDir,
		"reopened.docx",
	)
	if err := os.WriteFile(reopenPath, buf, 0o600); err != nil {
		t.Fatalf(
			"WriteFile() error = %v",
			err,
		)
	}

	// Open the saved document
	doc2, err := Open(reopenPath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify main part exists
	mainPart := doc2.MainPart()
	if mainPart == nil {
		t.Error(
			"Reopened document missing main part",
		)
	}
}

// TestDocumentSaveToWriter_Closed tests SaveTo with a closed document.
func TestDocumentSaveToWriter_Closed(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-saveto-closed-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create and close a document
	testPath := filepath.Join(
		tmpDir,
		"test.docx",
	)
	doc, err := New(
		testPath,
		DocTypeDocument,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	_ = doc.Close()

	// Try to save closed document
	var buf []byte
	w := &bufferWriter{buf: &buf}
	err = doc.SaveTo(w)

	// Should get an error
	if err == nil {
		t.Error(
			"SaveTo() on closed document should return error",
		)
	}
}

// bufferWriter wraps a byte slice pointer to implement io.Writer.
type bufferWriter struct {
	buf *[]byte
}

func (w *bufferWriter) Write(
	p []byte,
) (n int, err error) {
	*w.buf = append(*w.buf, p...)

	return len(p), nil
}

// TestAddHeaderWithType tests that AddHeader correctly links headers to section properties.
func TestAddHeaderWithType(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-header-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"header.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Get the elements package reference
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() returned nil")
	}

	t.Run("Default Header", func(t *testing.T) {
		// Add a default header
		header, err := doc.AddHeader(
			elements.HeaderFooterDefault,
		)
		if err != nil {
			t.Fatalf(
				"AddHeader(Default) error = %v",
				err,
			)
		}
		if header == nil {
			t.Fatal(
				"AddHeader(Default) returned nil",
			)
		}

		// Verify section properties have the header reference
		docElem := mainPart.Document()
		if docElem == nil {
			t.Fatal("Document() returned nil")
		}

		body := docElem.Body()
		if body == nil {
			t.Fatal("Body() returned nil")
		}

		sectPr := body.SectionProperties()
		if sectPr == nil {
			t.Fatal(
				"SectionProperties() returned nil",
			)
		}

		headerRef := sectPr.GetHeaderReference(
			elements.HeaderFooterDefault,
		)
		if headerRef == nil {
			t.Error(
				"Expected default header reference in section properties",
			)
		} else {
			if headerRef.Type() != elements.HeaderFooterDefault {
				t.Errorf("Expected header type %v, got %v", elements.HeaderFooterDefault, headerRef.Type())
			}
			if headerRef.RelationshipId() == "" {
				t.Error("Expected non-empty relationship ID")
			}
		}
	})

	t.Run(
		"First Page Header",
		func(t *testing.T) {
			// Add a first page header
			header, err := doc.AddHeader(
				elements.HeaderFooterValuesFirst,
			)
			if err != nil {
				t.Fatalf(
					"AddHeader(First) error = %v",
					err,
				)
			}
			if header == nil {
				t.Fatal(
					"AddHeader(First) returned nil",
				)
			}

			// Verify section properties have the header reference
			docElem := mainPart.Document()
			body := docElem.Body()
			sectPr := body.SectionProperties()

			headerRef := sectPr.GetHeaderReference(
				elements.HeaderFooterValuesFirst,
			)
			if headerRef == nil {
				t.Error(
					"Expected first page header reference in section properties",
				)
			} else if headerRef.Type() != elements.HeaderFooterValuesFirst {
				t.Errorf("Expected header type %v, got %v", elements.HeaderFooterValuesFirst, headerRef.Type())
			}

			// Verify titlePg is set
			if !sectPr.TitlePage() {
				t.Error(
					"Expected titlePg to be set for first page header",
				)
			}
		},
	)

	t.Run("Even Page Header", func(t *testing.T) {
		// Add an even page header
		header, err := doc.AddHeader(
			elements.HeaderFooterValuesEven,
		)
		if err != nil {
			t.Fatalf(
				"AddHeader(Even) error = %v",
				err,
			)
		}
		if header == nil {
			t.Fatal(
				"AddHeader(Even) returned nil",
			)
		}

		// Verify section properties have the header reference
		docElem := mainPart.Document()
		body := docElem.Body()
		sectPr := body.SectionProperties()

		headerRef := sectPr.GetHeaderReference(
			elements.HeaderFooterValuesEven,
		)
		if headerRef == nil {
			t.Error(
				"Expected even page header reference in section properties",
			)
		} else if headerRef.Type() != elements.HeaderFooterValuesEven {
			t.Errorf("Expected header type %v, got %v", elements.HeaderFooterValuesEven, headerRef.Type())
		}
	})
}

// TestAddFooterWithType tests that AddFooter correctly links footers to section properties.
func TestAddFooterWithType(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-footer-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"footer.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() returned nil")
	}

	t.Run("Default Footer", func(t *testing.T) {
		// Add a default footer
		footer, err := doc.AddFooter(
			elements.HeaderFooterDefault,
		)
		if err != nil {
			t.Fatalf(
				"AddFooter(Default) error = %v",
				err,
			)
		}
		if footer == nil {
			t.Fatal(
				"AddFooter(Default) returned nil",
			)
		}

		// Verify section properties have the footer reference
		docElem := mainPart.Document()
		body := docElem.Body()
		sectPr := body.SectionProperties()

		footerRef := sectPr.GetFooterReference(
			elements.HeaderFooterDefault,
		)
		if footerRef == nil {
			t.Error(
				"Expected default footer reference in section properties",
			)
		} else {
			if footerRef.Type() != elements.HeaderFooterDefault {
				t.Errorf("Expected footer type %v, got %v", elements.HeaderFooterDefault, footerRef.Type())
			}
			if footerRef.RelationshipId() == "" {
				t.Error("Expected non-empty relationship ID")
			}
		}
	})

	t.Run(
		"First Page Footer",
		func(t *testing.T) {
			// Add a first page footer
			footer, err := doc.AddFooter(
				elements.HeaderFooterValuesFirst,
			)
			if err != nil {
				t.Fatalf(
					"AddFooter(First) error = %v",
					err,
				)
			}
			if footer == nil {
				t.Fatal(
					"AddFooter(First) returned nil",
				)
			}

			// Verify section properties have the footer reference
			docElem := mainPart.Document()
			body := docElem.Body()
			sectPr := body.SectionProperties()

			footerRef := sectPr.GetFooterReference(
				elements.HeaderFooterValuesFirst,
			)
			if footerRef == nil {
				t.Error(
					"Expected first page footer reference in section properties",
				)
			} else if footerRef.Type() != elements.HeaderFooterValuesFirst {
				t.Errorf("Expected footer type %v, got %v", elements.HeaderFooterValuesFirst, footerRef.Type())
			}

			// Verify titlePg is set
			if !sectPr.TitlePage() {
				t.Error(
					"Expected titlePg to be set for first page footer",
				)
			}
		},
	)

	t.Run("Even Page Footer", func(t *testing.T) {
		// Add an even page footer
		footer, err := doc.AddFooter(
			elements.HeaderFooterValuesEven,
		)
		if err != nil {
			t.Fatalf(
				"AddFooter(Even) error = %v",
				err,
			)
		}
		if footer == nil {
			t.Fatal(
				"AddFooter(Even) returned nil",
			)
		}

		// Verify section properties have the footer reference
		docElem := mainPart.Document()
		body := docElem.Body()
		sectPr := body.SectionProperties()

		footerRef := sectPr.GetFooterReference(
			elements.HeaderFooterValuesEven,
		)
		if footerRef == nil {
			t.Error(
				"Expected even page footer reference in section properties",
			)
		} else if footerRef.Type() != elements.HeaderFooterValuesEven {
			t.Errorf("Expected footer type %v, got %v", elements.HeaderFooterValuesEven, footerRef.Type())
		}
	})
}

// TestHeaderFooterRelationshipIDs tests that relationship IDs are correctly set and unique.
func TestHeaderFooterRelationshipIDs(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-hf-rel-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()

	// Add multiple headers and footers
	_, err = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(Default) error = %v",
			err,
		)
	}

	_, err = doc.AddHeader(
		elements.HeaderFooterValuesFirst,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(First) error = %v",
			err,
		)
	}

	_, err = doc.AddFooter(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"AddFooter(Default) error = %v",
			err,
		)
	}

	_, err = doc.AddFooter(
		elements.HeaderFooterValuesFirst,
	)
	if err != nil {
		t.Fatalf(
			"AddFooter(First) error = %v",
			err,
		)
	}

	// Get section properties
	docElem := mainPart.Document()
	body := docElem.Body()
	sectPr := body.SectionProperties()

	// Collect all relationship IDs
	relIDs := make(map[string]bool)

	headerRefs := sectPr.HeaderReferences()
	for _, hr := range headerRefs {
		relID := hr.RelationshipId()
		if relID == "" {
			t.Error(
				"Found empty relationship ID in header reference",
			)
		}
		if relIDs[relID] {
			t.Errorf(
				"Duplicate relationship ID: %s",
				relID,
			)
		}
		relIDs[relID] = true
	}

	footerRefs := sectPr.FooterReferences()
	for _, fr := range footerRefs {
		relID := fr.RelationshipId()
		if relID == "" {
			t.Error(
				"Found empty relationship ID in footer reference",
			)
		}
		if relIDs[relID] {
			t.Errorf(
				"Duplicate relationship ID: %s",
				relID,
			)
		}
		relIDs[relID] = true
	}

	// Verify we have 4 unique relationship IDs
	if len(relIDs) != 4 {
		t.Errorf(
			"Expected 4 unique relationship IDs, got %d",
			len(relIDs),
		)
	}
}

// TestHeaderFooterReplacement tests that adding a header/footer of the same type replaces the existing one.
func TestHeaderFooterReplacement(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-hf-replace-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()

	// Add a default header
	_, err = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(Default) error = %v",
			err,
		)
	}

	// Get the first relationship ID
	docElem := mainPart.Document()
	body := docElem.Body()
	sectPr := body.SectionProperties()
	headerRef1 := sectPr.GetHeaderReference(
		elements.HeaderFooterDefault,
	)
	if headerRef1 == nil {
		t.Fatal("Expected header reference")
	}
	relID1 := headerRef1.RelationshipId()

	// Add another default header (should replace)
	_, err = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"Second AddHeader(Default) error = %v",
			err,
		)
	}

	// Get the second relationship ID
	headerRef2 := sectPr.GetHeaderReference(
		elements.HeaderFooterDefault,
	)
	if headerRef2 == nil {
		t.Fatal(
			"Expected header reference after replacement",
		)
	}
	relID2 := headerRef2.RelationshipId()

	// Relationship IDs should be different (new part created)
	if relID1 == relID2 {
		t.Error(
			"Expected different relationship ID after replacement",
		)
	}

	// Should still have only one default header reference
	defaultCount := 0
	for _, hr := range sectPr.HeaderReferences() {
		if hr.Type() == elements.HeaderFooterDefault {
			defaultCount++
		}
	}
	if defaultCount != 1 {
		t.Errorf(
			"Expected 1 default header reference, got %d",
			defaultCount,
		)
	}
}

// TestHeaderFooterPersistence_Create tests creating and saving a document with headers and footers.
func TestHeaderFooterPersistence_Create(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-hf-persist-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"persist.docx",
	)

	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add all three types of headers
	_, err = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(Default) error = %v",
			err,
		)
	}

	_, err = doc.AddHeader(
		elements.HeaderFooterValuesFirst,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(First) error = %v",
			err,
		)
	}

	_, err = doc.AddHeader(
		elements.HeaderFooterValuesEven,
	)
	if err != nil {
		t.Fatalf(
			"AddHeader(Even) error = %v",
			err,
		)
	}

	// Add all three types of footers
	_, err = doc.AddFooter(
		elements.HeaderFooterDefault,
	)
	if err != nil {
		t.Fatalf(
			"AddFooter(Default) error = %v",
			err,
		)
	}

	_, err = doc.AddFooter(
		elements.HeaderFooterValuesFirst,
	)
	if err != nil {
		t.Fatalf(
			"AddFooter(First) error = %v",
			err,
		)
	}

	_, err = doc.AddFooter(
		elements.HeaderFooterValuesEven,
	)
	if err != nil {
		t.Fatalf(
			"AddFooter(Even) error = %v",
			err,
		)
	}

	// Save and close
	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

// TestHeaderFooterPersistence_Verify tests that header/footer references are correctly persisted and loaded.
func TestHeaderFooterPersistence_Verify(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-hf-persist-verify-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"persist.docx",
	)

	// Create a document with headers and footers
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	_, _ = doc.AddHeader(
		elements.HeaderFooterDefault,
	)
	_, _ = doc.AddHeader(
		elements.HeaderFooterValuesFirst,
	)
	_, _ = doc.AddHeader(
		elements.HeaderFooterValuesEven,
	)
	_, _ = doc.AddFooter(
		elements.HeaderFooterDefault,
	)
	_, _ = doc.AddFooter(
		elements.HeaderFooterValuesFirst,
	)
	_, _ = doc.AddFooter(
		elements.HeaderFooterValuesEven,
	)

	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Reopen and verify
	doc, err = Open(testPath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() returned nil")
	}

	// Verify we have exactly 3 header and 3 footer parts
	headerParts := mainPart.HeaderParts()
	if len(headerParts) != 3 {
		t.Errorf(
			"Expected 3 header parts, got %d",
			len(headerParts),
		)
	}

	footerParts := mainPart.FooterParts()
	if len(footerParts) != 3 {
		t.Errorf(
			"Expected 3 footer parts, got %d",
			len(footerParts),
		)
	}

	// Verify the raw XML contains the section properties with header/footer references
	rawData := string(mainPart.GetData())
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("<w:sectPr"),
	) {
		t.Error(
			"Expected sectPr element in document XML",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("<w:headerReference"),
	) {
		t.Error(
			"Expected headerReference elements in document XML",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("<w:footerReference"),
	) {
		t.Error(
			"Expected footerReference elements in document XML",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("w:type=\"default\""),
	) {
		t.Error(
			"Expected default type header/footer references",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("w:type=\"first\""),
	) {
		t.Error(
			"Expected first type header/footer references",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("w:type=\"even\""),
	) {
		t.Error(
			"Expected even type header/footer references",
		)
	}
	if !bytes.Contains(
		[]byte(rawData),
		[]byte("<w:titlePg"),
	) {
		t.Error(
			"Expected titlePg element in section properties",
		)
	}
}
