package wordprocessing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"validate.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer doc.Close()

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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"validate.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer doc.Close()

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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"isvalid.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer doc.Close()

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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"closed.docx",
	)

	// Create and close a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	doc.Close()

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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"versions.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer doc.Close()

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
	defer doc.Close()

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
