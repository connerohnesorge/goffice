package wordprocessing

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestDocTypeString(t *testing.T) {
	tests := []struct {
		docType  DocType
		expected string
	}{
		{DocTypeDocument, "Document"},
		{DocTypeTemplate, "Template"},
		{
			DocTypeMacroEnabled,
			"MacroEnabledDocument",
		},
		{
			DocTypeMacroTemplate,
			"MacroEnabledTemplate",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.docType.String(); got != tt.expected {
				t.Errorf(
					"DocType.String() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestDocTypeExtension(t *testing.T) {
	tests := []struct {
		docType  DocType
		expected string
	}{
		{DocTypeDocument, ".docx"},
		{DocTypeTemplate, ".dotx"},
		{DocTypeMacroEnabled, ".docm"},
		{DocTypeMacroTemplate, ".dotm"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.docType.Extension(); got != tt.expected {
				t.Errorf(
					"DocType.Extension() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestDocTypeContentType(t *testing.T) {
	tests := []struct {
		docType  DocType
		expected string
	}{
		{
			DocTypeDocument,
			ContentTypeWordMLDocument,
		},
		{
			DocTypeTemplate,
			ContentTypeWordMLTemplate,
		},
		{
			DocTypeMacroEnabled,
			ContentTypeWordMLMacroEnabled,
		},
		{
			DocTypeMacroTemplate,
			ContentTypeWordMLMacroTemplate,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.docType.String(),
			func(t *testing.T) {
				if got := tt.docType.ContentType(); got != tt.expected {
					t.Errorf(
						"DocType.ContentType() = %v, want %v",
						got,
						tt.expected,
					)
				}
			},
		)
	}
}

func TestNewDocument(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")

	// Create a new document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify document properties
	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeDocument,
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

	// Verify main part exists
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Error("MainPart() = nil, want non-nil")
	}

	// Verify package and features exist
	if doc.Package() == nil {
		t.Error("Package() = nil, want non-nil")
	}

	if doc.Features() == nil {
		t.Error("Features() = nil, want non-nil")
	}
}

func TestNewDocumentWithDifferentTypes(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	docTypes := []DocType{
		DocTypeDocument,
		DocTypeTemplate,
		DocTypeMacroEnabled,
		DocTypeMacroTemplate,
	}

	for _, dt := range docTypes {
		t.Run(dt.String(), func(t *testing.T) {
			testPath := filepath.Join(
				tmpDir,
				"test"+dt.Extension(),
			)
			doc, err := New(testPath, dt)
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			defer func() { _ = doc.Close() }()

			if doc.Type() != dt {
				t.Errorf(
					"Type() = %v, want %v",
					doc.Type(),
					dt,
				)
			}
		})
	}
}

func TestDocumentSaveAs(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")
	savePath := filepath.Join(
		tmpDir,
		"saved.docx",
	)

	// Create and save a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = doc.SaveAs(savePath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	_ = doc.Close()

	// Verify file was created
	if _, err := os.Stat(savePath); os.IsNotExist(
		err,
	) {
		t.Error("SaveAs() did not create file")
	}
}

func TestDocumentChangeType(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
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

	// Change to template
	err = doc.ChangeType(DocTypeTemplate)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypeTemplate {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeTemplate,
		)
	}

	// Change to same type should be no-op
	err = doc.ChangeType(DocTypeTemplate)
	if err != nil {
		t.Errorf(
			"ChangeType() to same type error = %v",
			err,
		)
	}
}

func TestNewWriter(t *testing.T) {
	var buf bytes.Buffer

	doc, err := NewWriter(&buf, DocTypeDocument)
	if err != nil {
		t.Fatalf("NewWriter() error = %v", err)
	}

	// Verify document properties
	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeDocument,
		)
	}

	if doc.Path() != "" {
		t.Errorf(
			"Path() = %v, want empty string",
			doc.Path(),
		)
	}

	if !doc.IsEditable() {
		t.Error("IsEditable() = false, want true")
	}

	_ = doc.Close()
}

func TestOpenSettings(t *testing.T) {
	settings := DefaultOpenSettings()

	if settings.AutoSave {
		t.Error(
			"Default AutoSave should be false",
		)
	}

	if settings.MaxCharactersInPart != 0 {
		t.Error(
			"Default MaxCharactersInPart should be 0",
		)
	}

	// Test WithAutoSave
	newSettings := settings.WithAutoSave(true)
	if !newSettings.AutoSave {
		t.Error(
			"WithAutoSave(true) should set AutoSave to true",
		)
	}
	// Original should be unchanged
	if settings.AutoSave {
		t.Error(
			"Original settings should be unchanged",
		)
	}

	// Test WithMaxCharacters
	newSettings = settings.WithMaxCharacters(1000)
	if newSettings.MaxCharactersInPart != 1000 {
		t.Error(
			"WithMaxCharacters(1000) should set MaxCharactersInPart to 1000",
		)
	}

	// Test WithTargetVersion
	newSettings = settings.WithTargetVersion(
		FileFormatVersionOffice2019,
	)
	if newSettings.MarkupCompatibilityProcessSettings.TargetFileFormatVersions != FileFormatVersionOffice2019 {
		t.Error(
			"WithTargetVersion should set TargetFileFormatVersions",
		)
	}
}

func TestDocumentCloseReleasesResources(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
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

	// Close the document
	err = doc.Close()
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// After close, MainPart should be nil
	if doc.MainPart() != nil {
		t.Error(
			"MainPart() should be nil after close",
		)
	}

	// Second close should not error
	err = doc.Close()
	if err != nil {
		t.Errorf("Second Close() error = %v", err)
	}
}

func TestAddMainPartError(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
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

	// Trying to add main part again should fail
	_, err = doc.AddMainPart()
	if err != ErrMainPartExists {
		t.Errorf(
			"AddMainPart() error = %v, want %v",
			err,
			ErrMainPartExists,
		)
	}
}

func TestDocumentOpenAndSave(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")

	// Create and save a document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = doc1.SaveAs(testPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Open the saved document
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify it loaded correctly
	if doc2.Type() != DocTypeDocument {
		t.Errorf(
			"Opened document Type() = %v, want %v",
			doc2.Type(),
			DocTypeDocument,
		)
	}

	if !doc2.IsEditable() {
		t.Error(
			"Opened document IsEditable() = false, want true",
		)
	}
}

func TestDocumentOpenReadOnly(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-test-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.docx")

	// Create and save a document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	err = doc1.SaveAs(testPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Open read-only
	doc2, err := Open(testPath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.IsEditable() {
		t.Error(
			"Read-only document IsEditable() = true, want false",
		)
	}

	// Save should fail on read-only document
	err = doc2.Save()
	if err != ErrReadOnly {
		t.Errorf(
			"Save() on read-only error = %v, want %v",
			err,
			ErrReadOnly,
		)
	}

	// ChangeType should fail on read-only document
	err = doc2.ChangeType(DocTypeTemplate)
	if err != ErrReadOnly {
		t.Errorf(
			"ChangeType() on read-only error = %v, want %v",
			err,
			ErrReadOnly,
		)
	}
}
