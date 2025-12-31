// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// createTestTemplate creates a simple Word document template for testing.
func createTestTemplate(
	t *testing.T,
	fields ...string,
) *wordprocessing.Document {
	t.Helper()

	tmpDir := t.TempDir()
	templatePath := filepath.Join(
		tmpDir,
		"template.docx",
	)

	doc, err := wordprocessing.New(
		templatePath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("Template has no main part")
	}

	document := mainPart.Document()
	if document == nil {
		t.Fatal("Template has no document")
	}

	body := document.GetOrCreateBody()

	// Add a paragraph with merge fields
	para := elements.NewParagraph()
	body.AppendChild(para)

	for i, fieldName := range fields {
		if i > 0 {
			// Add space between fields
			run := elements.NewRun(" ")
			para.AppendChild(run)
		}

		// Create a simple field for this merge field
		simpleField := elements.NewSimpleField()
		simpleField.Instruction = types.NewStringValue(
			"MERGEFIELD " + fieldName + " \\* MERGEFORMAT",
		)

		// Add placeholder text
		run := elements.NewRun(
			"«" + fieldName + "»",
		)
		simpleField.AppendChild(run)

		para.AppendChild(simpleField)
	}

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}

	return doc
}

func TestExecuteToDocuments_MultipleRecords(
	t *testing.T,
) {
	// Create template with Name and Age fields
	template := createTestTemplate(
		t,
		"Name",
		"Age",
	)

	// Create data source with 3 records
	ds := NewMapDataSource([]map[string]string{
		{"Name": "Alice", "Age": "25"},
		{"Name": "Bob", "Age": "30"},
		{"Name": "Charlie", "Age": "35"},
	})

	// Create mail merge engine
	mm := New(template).DataSource(ds)

	// Execute to documents
	docs, err := mm.ExecuteToDocuments()
	if err != nil {
		t.Fatalf(
			"ExecuteToDocuments failed: %v",
			err,
		)
	}

	// Should have 3 documents
	if len(docs) != 3 {
		t.Errorf(
			"Expected 3 documents, got %d",
			len(docs),
		)
	}

	// Verify each document is valid
	for i, doc := range docs {
		if doc == nil {
			t.Errorf("Document %d is nil", i)

			continue
		}

		mainPart := doc.MainPart()
		if mainPart == nil {
			t.Errorf(
				"Document %d has no main part",
				i,
			)

			continue
		}

		document := mainPart.Document()
		if document == nil {
			t.Errorf(
				"Document %d has no document element",
				i,
			)

			continue
		}

		// Clean up
		_ = doc.Close()
	}
}

func TestExecuteToDocuments_NoRecords(
	t *testing.T,
) {
	// Create template
	template := createTestTemplate(t, "Name")

	// Create empty data source
	ds := NewMapDataSource(nil)

	// Create mail merge engine
	mm := New(template).DataSource(ds)

	// Execute to documents
	docs, err := mm.ExecuteToDocuments()

	// Should return error for no records
	if err != ErrNoRecords {
		t.Errorf(
			"Expected ErrNoRecords, got %v",
			err,
		)
	}

	if docs != nil {
		t.Errorf(
			"Expected nil docs for no records, got %d docs",
			len(docs),
		)
	}
}

func TestExecuteToDocuments_WithSkipIf(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	templatePath := filepath.Join(
		tmpDir,
		"template.docx",
	)

	// Create a template with SKIPIF field
	doc, err := wordprocessing.New(
		templatePath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}

	mainPart := doc.MainPart()
	document := mainPart.Document()
	body := document.GetOrCreateBody()

	// Add SKIPIF field
	skipPara := elements.NewParagraph()
	body.AppendChild(skipPara)

	skipField := elements.NewSimpleField()
	skipField.Instruction = types.NewStringValue(
		"SKIPIF Status = Inactive",
	)
	skipPara.AppendChild(skipField)

	// Add Name field
	namePara := elements.NewParagraph()
	body.AppendChild(namePara)

	nameField := elements.NewSimpleField()
	nameField.Instruction = types.NewStringValue(
		"MERGEFIELD Name \\* MERGEFORMAT",
	)
	nameRun := elements.NewRun("«Name»")
	nameField.AppendChild(nameRun)
	namePara.AppendChild(nameField)

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}

	// Create data source with 3 records, one with Inactive status
	ds := NewMapDataSource([]map[string]string{
		{"Name": "Alice", "Status": "Active"},
		{
			"Name":   "Bob",
			"Status": "Inactive",
		}, // This should be skipped
		{"Name": "Charlie", "Status": "Active"},
	})

	// Create mail merge engine
	mm := New(doc).DataSource(ds)

	// Execute to documents
	docs, err := mm.ExecuteToDocuments()
	if err != nil {
		t.Fatalf(
			"ExecuteToDocuments failed: %v",
			err,
		)
	}

	// Should have 2 documents (Bob should be skipped)
	if len(docs) != 2 {
		t.Errorf(
			"Expected 2 documents (one skipped), got %d",
			len(docs),
		)
	}

	// Clean up
	for _, d := range docs {
		_ = d.Close()
	}
}

func TestExecuteToDocuments_StrictFields(
	t *testing.T,
) {
	// Create template with Name field
	template := createTestTemplate(t, "Name")

	// Create data source without Name field
	ds := NewMapDataSource([]map[string]string{
		{"Email": "alice@example.com"},
	})

	// Create mail merge engine with strict fields
	opts := &MergeOptions{
		StrictFields:       true,
		RemoveUnusedFields: false,
	}
	mm := New(
		template,
	).DataSource(ds).
		Options(opts)

	// Execute to documents
	_, err := mm.ExecuteToDocuments()

	// Should return error for missing field
	if err == nil {
		t.Error(
			"Expected error for missing field in strict mode",
		)
	}
}

func TestExecute_WithGreetingLine(t *testing.T) {
	tmpDir := t.TempDir()
	templatePath := filepath.Join(
		tmpDir,
		"template.docx",
	)

	// Create a template with GREETINGLINE field
	doc, err := wordprocessing.New(
		templatePath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}

	mainPart := doc.MainPart()
	document := mainPart.Document()
	body := document.GetOrCreateBody()

	// Add GREETINGLINE field
	para := elements.NewParagraph()
	body.AppendChild(para)

	greetingField := elements.NewSimpleField()
	greetingField.Instruction = types.NewStringValue(
		"GREETINGLINE",
	)
	greetingRun := elements.NewRun(
		"«GreetingLine»",
	)
	greetingField.AppendChild(greetingRun)
	para.AppendChild(greetingField)

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save template: %v",
			err,
		)
	}

	// Create data source with name fields
	ds := NewMapDataSource([]map[string]string{
		{
			"Title":     "Dr.",
			"FirstName": "Jane",
			"LastName":  "Smith",
		},
	})

	// Create mail merge engine
	mm := New(doc).DataSource(ds)

	// Execute
	merged, err := mm.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verify the document was created
	if merged == nil {
		t.Fatal(
			"Expected merged document, got nil",
		)
	}

	// Save the merged document to verify it's valid
	mergedPath := filepath.Join(
		tmpDir,
		"merged.docx",
	)
	if err := merged.SaveAs(mergedPath); err != nil {
		t.Errorf(
			"Failed to save merged document: %v",
			err,
		)
	}

	// Verify file was created
	if _, err := os.Stat(mergedPath); os.IsNotExist(
		err,
	) {
		t.Error(
			"Merged document file was not created",
		)
	}

	// Clean up
	_ = merged.Close()
}

func TestExecuteToDocuments_NoDataSource(
	t *testing.T,
) {
	// Create template
	template := createTestTemplate(t, "Name")

	// Create mail merge engine without data source
	mm := New(template)

	// Execute to documents
	_, err := mm.ExecuteToDocuments()

	// Should return error for no data source
	if err != ErrNoDataSource {
		t.Errorf(
			"Expected ErrNoDataSource, got %v",
			err,
		)
	}
}

func TestShouldSkipRecord_NoSkipIfFields(
	t *testing.T,
) {
	ds := NewMapDataSource([]map[string]string{
		{"Name": "Alice"},
	})
	_ = ds.Open()
	ds.Next()

	template := createTestTemplate(t, "Name")
	mm := New(template).DataSource(ds)

	// Find fields (should only have Name field, no SKIPIF)
	fields := mm.findMergeFields()

	// Check if record should be skipped
	skip, err := mm.shouldSkipRecord(fields)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if skip {
		t.Error(
			"Expected record not to be skipped (no SKIPIF fields)",
		)
	}
}

func TestExecuteToDocuments_SaveOutput(
	t *testing.T,
) {
	// Create template
	template := createTestTemplate(
		t,
		"Name",
		"Email",
	)

	// Create data source
	ds := NewMapDataSource([]map[string]string{
		{
			"Name":  "Alice",
			"Email": "alice@example.com",
		},
		{
			"Name":  "Bob",
			"Email": "bob@example.com",
		},
	})

	// Create mail merge engine
	mm := New(template).DataSource(ds)

	// Execute to documents
	docs, err := mm.ExecuteToDocuments()
	if err != nil {
		t.Fatalf(
			"ExecuteToDocuments failed: %v",
			err,
		)
	}

	// Save each document
	tmpDir := t.TempDir()
	for i, doc := range docs {
		outputPath := filepath.Join(
			tmpDir,
			"output_"+string(rune('0'+i))+".docx",
		)
		if err := doc.SaveAs(outputPath); err != nil {
			t.Errorf(
				"Failed to save document %d: %v",
				i,
				err,
			)
		}

		// Verify file exists
		if _, err := os.Stat(outputPath); os.IsNotExist(
			err,
		) {
			t.Errorf(
				"Output file %d was not created",
				i,
			)
		}

		// Clean up
		_ = doc.Close()
	}
}
