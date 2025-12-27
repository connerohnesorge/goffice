package spreadsheet

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
)

// TestCreate tests creating a new spreadsheet document.
func TestCreate(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-spreadsheet-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.xlsx")

	// Create a document
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify document properties
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}

	if doc.Path() != testPath {
		t.Errorf(
			"Path() = %v, want %v",
			doc.Path(),
			testPath,
		)
	}

	if !doc.IsEditable() {
		t.Error("IsEditable() = false, want true")
	}

	if doc.Package() == nil {
		t.Error("Package() returned nil")
	}
}

// TestCreateDifferentTypes tests creating documents of different types.
func TestCreateDifferentTypes(t *testing.T) {
	tests := []struct {
		name        string
		docType     DocType
		extension   string
		contentType string
	}{
		{
			name:        "Workbook",
			docType:     DocTypeWorkbook,
			extension:   ".xlsx",
			contentType: ContentTypeSpreadsheetMLWorkbook,
		},
		{
			name:        "Template",
			docType:     DocTypeTemplate,
			extension:   ".xltx",
			contentType: ContentTypeSpreadsheetMLTemplate,
		},
		{
			name:        "MacroEnabledWorkbook",
			docType:     DocTypeMacroEnabledWorkbook,
			extension:   ".xlsm",
			contentType: ContentTypeSpreadsheetMLMacroEnabled,
		},
		{
			name:        "MacroEnabledTemplate",
			docType:     DocTypeMacroEnabledTemplate,
			extension:   ".xltm",
			contentType: ContentTypeSpreadsheetMLMacroTemplate,
		},
		{
			name:        "AddIn",
			docType:     DocTypeAddIn,
			extension:   ".xlam",
			contentType: ContentTypeSpreadsheetMLAddIn,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test String method
			if tc.docType.String() != tc.name {
				t.Errorf(
					"String() = %v, want %v",
					tc.docType.String(),
					tc.name,
				)
			}

			// Test Extension method
			if tc.docType.Extension() != tc.extension {
				t.Errorf(
					"Extension() = %v, want %v",
					tc.docType.Extension(),
					tc.extension,
				)
			}

			// Test ContentType method
			if tc.docType.ContentType() != tc.contentType {
				t.Errorf(
					"ContentType() = %v, want %v",
					tc.docType.ContentType(),
					tc.contentType,
				)
			}
		})
	}
}

// TestOpen tests opening an existing spreadsheet document.
func TestOpen(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Open the document read-only
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify document properties
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}

	if doc.Path() != "" &&
		doc.Path() != fixturePath {
		t.Errorf(
			"Path() = %v, want %v",
			doc.Path(),
			fixturePath,
		)
	}

	if doc.IsEditable() {
		t.Error(
			"IsEditable() = true, want false for read-only open",
		)
	}
}

// TestOpenEditable tests opening a document for editing.
func TestOpenEditable(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Open the document editable
	doc, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	if !doc.IsEditable() {
		t.Error("IsEditable() = false, want true")
	}
}

// TestOpenNumbers tests opening the numbers.xlsx fixture.
func TestOpenNumbers(t *testing.T) {
	fixturePath := "../testdata/fixtures/numbers.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: numbers.xlsx fixture not found",
		)
	}

	// Open the document
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify document type
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestSaveAs tests saving a document to a new location.
func TestSaveAs(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-saveas-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Open the document
	doc, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	// Save to new location
	newPath := filepath.Join(tmpDir, "copy.xlsx")
	err = doc.SaveAs(newPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Close original
	_ = doc.Close()

	// Verify new file exists
	if _, err := os.Stat(newPath); os.IsNotExist(
		err,
	) {
		t.Error("SaveAs() did not create file")
	}

	// Open the saved file to verify it's valid
	doc2, err := Open(newPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open saved file: %v",
			err,
		)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.Type() != DocTypeWorkbook {
		t.Errorf(
			"Saved document Type() = %v, want %v",
			doc2.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestRoundtrip tests opening, saving, and reopening a document.
func TestRoundtrip(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

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

	// Open original
	doc1, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	// Save to temp
	tempPath := filepath.Join(
		tmpDir,
		"roundtrip.xlsx",
	)
	err = doc1.SaveAs(tempPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Open saved file
	doc2, err := Open(tempPath, true)
	if err != nil {
		t.Fatalf(
			"Failed to open saved file: %v",
			err,
		)
	}

	// Save again
	tempPath2 := filepath.Join(
		tmpDir,
		"roundtrip2.xlsx",
	)
	err = doc2.SaveAs(tempPath2)
	if err != nil {
		t.Fatalf(
			"Second SaveAs() error = %v",
			err,
		)
	}
	_ = doc2.Close()

	// Open final file
	doc3, err := Open(tempPath2, false)
	if err != nil {
		t.Fatalf(
			"Failed to open second saved file: %v",
			err,
		)
	}
	defer func() { _ = doc3.Close() }()

	if doc3.Type() != DocTypeWorkbook {
		t.Errorf(
			"Roundtrip document Type() = %v, want %v",
			doc3.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestDocTypeDetection tests automatic document type detection.
func TestDocTypeDetection(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Open the document
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// The fixture should be detected as a workbook
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Detected type = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}
}

// TestChangeType tests changing the document type.
func TestChangeType(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-changetype-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.xlsx")

	// Create a workbook
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify initial type
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Initial Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}

	// Change to template
	err = doc.ChangeType(DocTypeTemplate)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	// Verify new type
	if doc.Type() != DocTypeTemplate {
		t.Errorf(
			"After ChangeType, Type() = %v, want %v",
			doc.Type(),
			DocTypeTemplate,
		)
	}
}

// TestReadOnlyOperations tests that read-only documents cannot be modified.
func TestReadOnlyOperations(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Open read-only
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Save should fail
	err = doc.Save()
	if err != ErrReadOnly {
		t.Errorf(
			"Save() on read-only doc error = %v, want %v",
			err,
			ErrReadOnly,
		)
	}

	// ChangeType should fail
	err = doc.ChangeType(DocTypeTemplate)
	if err != ErrReadOnly {
		t.Errorf(
			"ChangeType() on read-only doc error = %v, want %v",
			err,
			ErrReadOnly,
		)
	}
}

// TestOpenSettings tests the OpenSettings functionality.
func TestOpenSettings(t *testing.T) {
	settings := DefaultOpenSettings()

	// Test defaults
	if settings.AutoSave {
		t.Error(
			"DefaultOpenSettings() AutoSave = true, want false",
		)
	}

	if settings.MaxCharactersInPart != 0 {
		t.Errorf(
			"DefaultOpenSettings() MaxCharactersInPart = %d, want 0",
			settings.MaxCharactersInPart,
		)
	}

	// Test fluent setters
	s2 := settings.WithAutoSave(true)
	if !s2.AutoSave {
		t.Error(
			"WithAutoSave(true) AutoSave = false, want true",
		)
	}
	if settings.AutoSave {
		t.Error(
			"WithAutoSave should not modify original",
		)
	}

	s3 := settings.WithMaxCharacters(1000)
	if s3.MaxCharactersInPart != 1000 {
		t.Errorf(
			"WithMaxCharacters(1000) MaxCharactersInPart = %d, want 1000",
			s3.MaxCharactersInPart,
		)
	}

	s4 := settings.WithTargetVersion(
		FileFormatVersionOffice2019,
	)
	if s4.MarkupCompatibilityProcessSettings.TargetFileFormatVersions != FileFormatVersionOffice2019 {
		t.Error(
			"WithTargetVersion did not set correct version",
		)
	}
}

// TestValidate tests document validation.
func TestValidate(t *testing.T) {
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
		"validate.xlsx",
	)

	// Create a document
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

// TestValidateClosedDocument tests validation on a closed document.
func TestValidateClosedDocument(t *testing.T) {
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
		"closed.xlsx",
	)

	// Create and close a document
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

// TestIsValid tests the IsValid convenience method.
func TestIsValid(t *testing.T) {
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
		"isvalid.xlsx",
	)

	// Create a document
	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
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

// TestCreateFromStream tests creating a document from a stream.
func TestCreateFromStream(t *testing.T) {
	var buf bytes.Buffer

	// Create a document to a buffer
	doc, err := CreateFromStream(
		&buf,
		DocTypeWorkbook,
	)
	if err != nil {
		t.Fatalf(
			"CreateFromStream() error = %v",
			err,
		)
	}

	// Verify document properties
	if doc.Type() != DocTypeWorkbook {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeWorkbook,
		)
	}

	if !doc.IsEditable() {
		t.Error("IsEditable() = false, want true")
	}

	if doc.Path() != "" {
		t.Errorf(
			"Path() = %v, want empty string",
			doc.Path(),
		)
	}

	// Close the document
	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

// TestFeatures tests the Features method.
func TestFeatures(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-features-*",
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
		"features.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	features := doc.Features()
	if features == nil {
		t.Error("Features() returned nil")
	}
}

// TestCloseIdempotent tests that Close can be called multiple times.
func TestCloseIdempotent(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-close-*",
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
		"close.xlsx",
	)

	doc, err := Create(testPath, DocTypeWorkbook)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	// Close multiple times should not panic or error
	err = doc.Close()
	if err != nil {
		t.Errorf("First Close() error = %v", err)
	}

	err = doc.Close()
	if err != nil {
		t.Errorf("Second Close() error = %v", err)
	}
}

// TestValidateWithRealFile tests validation on a real .xlsx file.
func TestValidateWithRealFile(t *testing.T) {
	fixturePath := "../testdata/fixtures/minimal.xlsx"

	// Check if fixture exists
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip(
			"Skipping test: minimal.xlsx fixture not found",
		)
	}

	// Open the real .xlsx file
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open minimal.xlsx: %v",
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
