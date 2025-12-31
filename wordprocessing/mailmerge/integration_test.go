// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"os"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const testNameField = "Name\nJohn Doe\n"

// Helper: Create a template document with a SimpleField
func createTemplateWithSimpleField(
	fieldName string,
) (*wordprocessing.Document, error) {
	// Create a new document in memory
	tmpFile, err := os.CreateTemp(
		"",
		"template*.docx",
	)
	if err != nil {
		return nil, err
	}
	tmpPath := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := wordprocessing.New(
		tmpPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		return nil, err
	}

	// Get the body
	mainPart := doc.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	docElem := mainPart.Document()
	if docElem == nil {
		return nil, ErrNoDocument
	}

	body := docElem.GetOrCreateBody()

	// Create a paragraph with a SimpleField: <w:fldSimple w:instr="MERGEFIELD Name">
	para := elements.NewParagraph()

	field := elements.NewSimpleField()
	field.Instruction = types.NewStringValue(
		"MERGEFIELD " + fieldName,
	)

	// Add placeholder run inside field
	run := elements.NewRun("«" + fieldName + "»")
	field.AppendChild(run)

	para.AppendChild(field)
	body.AppendChild(para)

	// Don't save/reopen - just return the document as-is
	return doc, nil
}

// Helper: Create a template with multiple fields
func createTemplateWithFields(
	fields []string,
) (*wordprocessing.Document, error) {
	tmpFile, err := os.CreateTemp(
		"",
		"template*.docx",
	)
	if err != nil {
		return nil, err
	}
	tmpPath := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := wordprocessing.New(
		tmpPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		return nil, err
	}

	mainPart := doc.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	docElem := mainPart.Document()
	if docElem == nil {
		return nil, ErrNoDocument
	}

	body := docElem.GetOrCreateBody()

	// Create paragraphs with SimpleFields for each field
	for _, fieldName := range fields {
		para := elements.NewParagraph()

		field := elements.NewSimpleField()
		field.Instruction = types.NewStringValue(
			"MERGEFIELD " + fieldName,
		)

		run := elements.NewRun(
			"«" + fieldName + "»",
		)
		field.AppendChild(run)

		para.AppendChild(field)
		body.AppendChild(para)
	}

	// Don't save/reopen - just return as-is
	return doc, nil
}

// Helper: Create test CSV file
func createTestCSV(
	t *testing.T,
	content string,
) string {
	tmpFile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatalf(
			"Failed to create temp CSV: %v",
			err,
		)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf(
			"Failed to write CSV content: %v",
			err,
		)
	}
	_ = tmpFile.Close()

	t.Cleanup(
		func() { _ = os.Remove(tmpFile.Name()) },
	)

	return tmpFile.Name()
}

// Helper: Create test JSON file
func createTestJSON(
	t *testing.T,
	content string,
) string {
	tmpFile, err := os.CreateTemp(
		"",
		"test*.json",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp JSON: %v",
			err,
		)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf(
			"Failed to write JSON content: %v",
			err,
		)
	}
	_ = tmpFile.Close()

	t.Cleanup(
		func() { _ = os.Remove(tmpFile.Name()) },
	)

	return tmpFile.Name()
}

// Helper: Extract text content from document
func extractTextContent(
	doc *wordprocessing.Document,
) (string, error) {
	mainPart := doc.MainPart()
	if mainPart == nil {
		return "", ErrNoMainPart
	}

	docElem := mainPart.Document()
	if docElem == nil {
		return "", ErrNoDocument
	}

	body := docElem.GetOrCreateBody()

	var textContent strings.Builder

	// Walk through all children and extract text
	for child := range body.Children() {
		para, ok := child.(*elements.Paragraph)
		if !ok {
			continue
		}

		for paraChild := range para.Children() {
			switch elem := paraChild.(type) {
			case *elements.Run:
				for runChild := range elem.Children() {
					if text, ok := runChild.(*elements.Text); ok {
						textContent.WriteString(text.InnerText())
					}
				}
			case *elements.SimpleField:
				// Extract text from SimpleField children
				for fieldChild := range elem.Children() {
					run, ok := fieldChild.(*elements.Run)
					if !ok {
						continue
					}
					for runChild := range run.Children() {
						if text, ok := runChild.(*elements.Text); ok {
							textContent.WriteString(text.InnerText())
						}
					}
				}
			}
		}
		textContent.WriteString("\n")
	}

	return textContent.String(), nil
}

// Test 7.2.1: CSV merge end-to-end
func TestIntegration_CSVMerge(t *testing.T) {
	// Create template with SimpleField
	doc, err := createTemplateWithSimpleField(
		"Name",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Create CSV file with test data
	csvContent := testNameField
	csvPath := createTestCSV(t, csvContent)

	// Execute merge
	merged, err := New(doc).
		DataSource(NewCSVDataSource(csvPath)).
		Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if merged == nil {
		t.Fatal("Merged document is nil")
	}

	// Verify merged content
	content, err := extractTextContent(merged)
	if err != nil {
		t.Fatalf(
			"Failed to extract text: %v",
			err,
		)
	}
	if !strings.Contains(content, "John Doe") {
		t.Errorf(
			"Expected content to contain 'John Doe', got: %s",
			content,
		)
	}
	if strings.Contains(content, "«Name»") {
		t.Errorf(
			"Content should not contain placeholder '«Name»', got: %s",
			content,
		)
	}
}

// Test 7.2.2: JSON merge end-to-end
func TestIntegration_JSONMerge(t *testing.T) {
	// Create template with MERGEFIELD
	doc, err := createTemplateWithFields(
		[]string{"FirstName", "LastName"},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Create JSON file
	jsonContent := `[
		{"FirstName": "Alice", "LastName": "Smith"},
		{"FirstName": "Bob", "LastName": "Jones"}
	]`
	jsonPath := createTestJSON(t, jsonContent)

	// Execute merge (processes first record)
	merged, err := New(doc).
		DataSource(NewJSONDataSource(jsonPath)).
		Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if merged == nil {
		t.Fatal("Merged document is nil")
	}

	// Verify output contains first record data
	content, err := extractTextContent(merged)
	if err != nil {
		t.Fatalf(
			"Failed to extract text: %v",
			err,
		)
	}
	if !strings.Contains(content, "Alice") {
		t.Errorf(
			"Expected content to contain 'Alice', got: %s",
			content,
		)
	}
	if !strings.Contains(content, "Smith") {
		t.Errorf(
			"Expected content to contain 'Smith', got: %s",
			content,
		)
	}
}

// Test 7.2.3: Missing fields with strict mode
func TestIntegration_StrictFieldsError(
	t *testing.T,
) {
	// Create template with field that won't be in data
	doc, err := createTemplateWithFields(
		[]string{"Name", "MissingField"},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Create data source without MissingField
	csvContent := testNameField
	csvPath := createTestCSV(t, csvContent)

	// Execute merge with strict fields enabled
	opts := &MergeOptions{
		StrictFields:       true,
		RemoveUnusedFields: false,
	}

	_, err = New(doc).
		DataSource(NewCSVDataSource(csvPath)).
		Options(opts).
		Execute()

	// Should error because MissingField is not in data source
	if err == nil {
		t.Error(
			"Expected error for missing field in strict mode",
		)
	} else if !strings.Contains(err.Error(), "MissingField") {
		t.Errorf("Expected error to mention 'MissingField', got: %v", err)
	}
}

// Test 7.2.4: Missing fields with non-strict mode
func TestIntegration_NonStrictFieldsRemoval(
	t *testing.T,
) {
	// Create template with field that won't be in data
	doc, err := createTemplateWithFields(
		[]string{"Name", "OptionalField"},
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Create data source without OptionalField
	csvContent := testNameField
	csvPath := createTestCSV(t, csvContent)

	// Execute merge with strict fields disabled and remove unused fields enabled
	opts := &MergeOptions{
		StrictFields:       false,
		RemoveUnusedFields: true,
	}

	merged, err := New(doc).
		DataSource(NewCSVDataSource(csvPath)).
		Options(opts).
		Execute()
	if err != nil {
		t.Fatalf(
			"Execute should succeed in non-strict mode: %v",
			err,
		)
	}

	// Verify that OptionalField placeholder is removed
	content, err := extractTextContent(merged)
	if err != nil {
		t.Fatalf(
			"Failed to extract text: %v",
			err,
		)
	}
	if !strings.Contains(content, "John Doe") {
		t.Errorf(
			"Expected 'John Doe', got: %s",
			content,
		)
	}
	if strings.Contains(
		content,
		"«OptionalField»",
	) {
		t.Errorf(
			"Optional field placeholder should be removed, got: %s",
			content,
		)
	}
}

// Test 7.3: Test using fluent API
func TestIntegration_FluentAPIUsage(
	t *testing.T,
) {
	// Create template
	doc, err := createTemplateWithSimpleField(
		"Name",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create template: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Create CSV data
	csvContent := "Name\nWorld\n"
	csvPath := createTestCSV(t, csvContent)

	// Execute merge using fluent API
	result, err := New(doc).
		DataSource(NewCSVDataSource(csvPath)).
		Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if result == nil {
		t.Fatal("Result is nil")
	}

	// Verify merged content
	content, err := extractTextContent(result)
	if err != nil {
		t.Fatalf(
			"Failed to extract text: %v",
			err,
		)
	}
	if !strings.Contains(content, "World") {
		t.Errorf(
			"Expected 'World', got: %s",
			content,
		)
	}
	if strings.Contains(content, "«Name»") {
		t.Errorf(
			"Should not contain '«Name»', got: %s",
			content,
		)
	}
}
