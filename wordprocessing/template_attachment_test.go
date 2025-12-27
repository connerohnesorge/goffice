package wordprocessing

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/packaging"
)

func TestNewFromTemplateWithAttachment(
	t *testing.T,
) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create a simple template
	templatePath := filepath.Join(
		tmpDir,
		"template.dotx",
	)
	template, err := New(
		templatePath,
		DocTypeTemplate,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	if err := template.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}
	if err := template.Close(); err != nil {
		t.Fatalf(
			"Failed to close template: %v",
			err,
		)
	}

	// Create a document from the template with attachment
	doc, err := NewFromTemplate(
		templatePath,
		true,
	)
	if err != nil {
		t.Fatalf(
			"NewFromTemplate() error = %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Verify the main part exists
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Verify the settings part exists
	settingsPart := mainPart.SettingsPart()
	if settingsPart == nil {
		t.Fatal(
			"SettingsPart is nil after template attachment",
		)
	}

	// Verify the settings element exists
	settings := settingsPart.Settings()
	if settings == nil {
		t.Fatal("Settings element is nil")
	}

	// Verify the attachedTemplate element exists
	attachedTemplate := settings.GetElement(
		"attachedTemplate",
		"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	)
	if attachedTemplate == nil {
		t.Fatal(
			"attachedTemplate element not found in settings",
		)
	}

	// Verify the r:id attribute is set
	idAttr, found := attachedTemplate.GetAttribute(
		"id",
		"http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	)
	if !found {
		t.Fatal(
			"r:id attribute not found on attachedTemplate element",
		)
	}
	if idAttr.Value() == "" {
		t.Fatal("r:id attribute value is empty")
	}

	// Verify the external relationship exists
	pkg := doc.Package().Package()
	if pkg == nil {
		t.Fatal("Package is nil")
	}

	rels := pkg.PartRelationships(
		settingsPart.URI(),
	)
	if rels == nil {
		t.Fatal(
			"Relationships not found for settings part",
		)
	}

	rel, err := rels.Get(idAttr.Value())
	if err != nil {
		t.Fatalf(
			"Relationship not found: %v",
			err,
		)
	}

	if rel.Type() != "http://schemas.openxmlformats.org/officeDocument/2006/relationships/attachedTemplate" {
		t.Errorf(
			"Relationship type = %q, want attachedTemplate relationship type",
			rel.Type(),
		)
	}

	if rel.TargetMode() != packaging.TargetModeExternal {
		t.Errorf(
			"Relationship target mode = %v, want External",
			rel.TargetMode(),
		)
	}

	if rel.Target() != templatePath {
		t.Errorf(
			"Relationship target = %q, want %q",
			rel.Target(),
			templatePath,
		)
	}

	// Save the document to verify it can be written
	docPath := filepath.Join(
		tmpDir,
		"document_with_template.docx",
	)
	if err := doc.SaveAs(docPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify the file was created
	if _, err := os.Stat(docPath); os.IsNotExist(
		err,
	) {
		t.Fatal("Document file was not created")
	}
}

func TestNewFromTemplateWithoutAttachment(
	t *testing.T,
) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create a simple template
	templatePath := filepath.Join(
		tmpDir,
		"template.dotx",
	)
	template, err := New(
		templatePath,
		DocTypeTemplate,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	if err := template.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}
	if err := template.Close(); err != nil {
		t.Fatalf(
			"Failed to close template: %v",
			err,
		)
	}

	// Create a document from the template without attachment
	doc, err := NewFromTemplate(
		templatePath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"NewFromTemplate() error = %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Verify the main part exists
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Settings part might not exist if we don't attach template
	settingsPart := mainPart.SettingsPart()
	if settingsPart == nil {
		return
	}
	settings := settingsPart.Settings()
	if settings == nil {
		return
	}
	// If settings exist, verify there's no attachedTemplate
	attachedTemplate := settings.GetElement(
		"attachedTemplate",
		"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	)
	if attachedTemplate != nil {
		t.Error(
			"attachedTemplate element should not exist when attachTemplate is false",
		)
	}
}

func TestAttachTemplateLinkErrors(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T) *Document
		wantErr bool
	}{
		{
			name: "nil main part",
			setup: func(_ *testing.T) *Document {
				// Create a document without a main part
				doc := &Document{
					mainPart: nil,
				}

				return doc
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := tt.setup(t)
			err := doc.attachTemplateLink(
				"/some/path/template.dotx",
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"attachTemplateLink() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestTemplateAttachmentIntegration(
	t *testing.T,
) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create and save a template with content
	templatePath := filepath.Join(
		tmpDir,
		"MyTemplate.dotx",
	)
	template, err := New(
		templatePath,
		DocTypeTemplate,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}

	// Add some content to the template
	mainPart := template.MainPart()
	if mainPart == nil {
		t.Fatal("Template MainPart is nil")
	}

	if err := template.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}
	if err := template.Close(); err != nil {
		t.Fatalf(
			"Failed to close template: %v",
			err,
		)
	}

	// Create a document from the template with attachment
	doc, err := NewFromTemplate(
		templatePath,
		true,
	)
	if err != nil {
		t.Fatalf(
			"NewFromTemplate() error = %v",
			err,
		)
	}

	// Save the document
	docPath := filepath.Join(
		tmpDir,
		"DocumentFromTemplate.docx",
	)
	if err := doc.SaveAs(docPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Reopen the document and verify the template attachment persists
	reopenedDoc, err := Open(docPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to reopen document: %v",
			err,
		)
	}
	defer func() { _ = reopenedDoc.Close() }()

	// Verify the settings part exists
	mainPart = reopenedDoc.MainPart()
	if mainPart == nil {
		t.Fatal(
			"Reopened document MainPart is nil",
		)
	}

	settingsPart := mainPart.SettingsPart()
	if settingsPart == nil {
		t.Fatal(
			"SettingsPart is nil in reopened document",
		)
	}

	// Verify the settings element exists
	settings := settingsPart.Settings()
	if settings == nil {
		t.Fatal(
			"Settings element is nil in reopened document",
		)
	}

	// Verify the attachedTemplate element exists
	attachedTemplate := settings.GetElement(
		"attachedTemplate",
		"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	)
	if attachedTemplate == nil {
		t.Fatal(
			"attachedTemplate element not found in reopened document",
		)
	}

	// Verify the r:id attribute is set
	idAttr, found := attachedTemplate.GetAttribute(
		"id",
		"http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	)
	if !found {
		t.Fatal(
			"r:id attribute not found on attachedTemplate element in reopened document",
		)
	}
	if idAttr.Value() == "" {
		t.Fatal(
			"r:id attribute value is empty in reopened document",
		)
	}

	// Verify the external relationship exists and points to the template
	pkg := reopenedDoc.Package().Package()
	if pkg == nil {
		t.Fatal(
			"Package is nil in reopened document",
		)
	}

	rels := pkg.PartRelationships(
		settingsPart.URI(),
	)
	if rels == nil {
		t.Fatal(
			"Relationships not found for settings part in reopened document",
		)
	}

	rel, err := rels.Get(idAttr.Value())
	if err != nil {
		t.Fatalf(
			"Relationship not found in reopened document: %v",
			err,
		)
	}

	if rel.Target() != templatePath {
		t.Errorf(
			"Relationship target = %q, want %q",
			rel.Target(),
			templatePath,
		)
	}

	if rel.TargetMode() != packaging.TargetModeExternal {
		t.Errorf(
			"Relationship target mode = %v, want External",
			rel.TargetMode(),
		)
	}
}
