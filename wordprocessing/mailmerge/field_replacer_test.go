// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

//nolint:revive,goconst // test file allows more flexibility
package mailmerge

import (
	"os"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestApplyFormatSwitch(t *testing.T) {
	mm := New(nil)

	tests := []struct {
		name     string
		value    string
		switches map[string]string
		want     string
	}{
		{
			name:     "No switches",
			value:    "Hello World",
			switches: map[string]string{},
			want:     "Hello World",
		},
		{
			name:  "Upper case",
			value: "Hello World",
			switches: map[string]string{
				"\\*": "UPPER",
			},
			want: "HELLO WORLD",
		},
		{
			name:  "Lower case",
			value: "Hello World",
			switches: map[string]string{
				"\\*": "LOWER",
			},
			want: "hello world",
		},
		{
			name:  "First cap",
			value: "hello world",
			switches: map[string]string{
				"\\*": "FIRSTCAP",
			},
			want: "Hello world",
		},
		{
			name:  "Title case (Caps)",
			value: "hello world",
			switches: map[string]string{
				"\\*": "CAPS",
			},
			want: "Hello World",
		},
		{
			name:  "Unknown switch",
			value: "Hello World",
			switches: map[string]string{
				"\\*": "UNKNOWN",
			},
			want: "Hello World",
		},
		{
			name:  "Empty value",
			value: "",
			switches: map[string]string{
				"\\*": "UPPER",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mm.applyFormatSwitch(
				tt.value,
				tt.switches,
			)
			if got != tt.want {
				t.Errorf(
					"applyFormatSwitch() = %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestReplaceSimpleField(t *testing.T) {
	mm := New(nil)

	t.Run(
		"Replace simple field with plain text",
		func(t *testing.T) {
			// Create a paragraph with a SimpleField
			para := elements.NewParagraph()
			simpleField := elements.NewSimpleField()
			simpleField.Instruction = types.NewStringValue(
				"MERGEFIELD FirstName \\* MERGEFORMAT",
			)
			para.AppendChild(simpleField)

			// Add a run inside the SimpleField (placeholder text)
			run := elements.NewRun("«FirstName»")
			simpleField.AppendChild(run)

			// Create merge field structure
			field := &MergeField{
				FieldName: "FirstName",
				FieldType: FieldTypeSimple,
				Element:   simpleField,
				FieldCode: "MERGEFIELD FirstName \\* MERGEFORMAT",
				Switches: map[string]string{
					"\\*": "MERGEFORMAT",
				},
				Paragraph: para,
			}

			// Replace with value
			err := mm.replaceSimpleField(
				field,
				"John",
			)
			if err != nil {
				t.Fatalf(
					"replaceSimpleField failed: %v",
					err,
				)
			}

			// Verify the SimpleField was removed
			found := false
			for child := range para.Children() {
				if child == simpleField {
					found = true

					break
				}
			}
			if found {
				t.Error(
					"SimpleField should have been removed from paragraph",
				)
			}

			// Verify a Run was added with the value
			foundRun := false
			for child := range para.Children() {
				if run, ok := child.(*elements.Run); ok {
					text := run.InnerText()
					if text == "John" {
						foundRun = true

						break
					}
				}
			}
			if !foundRun {
				t.Error(
					"Expected to find Run with text 'John'",
				)
			}
		},
	)
}

func TestReplaceComplexField(t *testing.T) {
	mm := New(nil)

	t.Run(
		"Replace complex field with result text",
		func(t *testing.T) {
			// Create a paragraph with a ComplexField
			para := elements.NewParagraph()

			// Begin run with fldChar begin
			beginRun := elements.NewRun("")
			beginChar := elements.NewFieldChar(
				elements.FieldCharBegin,
			)
			beginRun.AppendChild(beginChar)
			para.AppendChild(beginRun)

			// InstrText run
			instrRun := elements.NewRun("")
			instrText := elements.NewInstrText(
				"MERGEFIELD FirstName \\* MERGEFORMAT",
			)
			instrRun.AppendChild(instrText)
			para.AppendChild(instrRun)

			// Separate run with fldChar separate
			separateRun := elements.NewRun("")
			separateChar := elements.NewFieldChar(
				elements.FieldCharSeparate,
			)
			separateRun.AppendChild(separateChar)
			para.AppendChild(separateRun)

			// Result run with placeholder text
			resultRun := elements.NewRun(
				"«FirstName»",
			)
			para.AppendChild(resultRun)

			// End run with fldChar end
			endRun := elements.NewRun("")
			endChar := elements.NewFieldChar(
				elements.FieldCharEnd,
			)
			endRun.AppendChild(endChar)
			para.AppendChild(endRun)

			// Create merge field structure
			field := &MergeField{
				FieldName:   "FirstName",
				FieldType:   FieldTypeComplex,
				Element:     beginRun,
				BeginRun:    beginRun,
				SeparateRun: separateRun,
				EndRun:      endRun,
				FieldCode:   "MERGEFIELD FirstName \\* MERGEFORMAT",
				Switches: map[string]string{
					"\\*": "MERGEFORMAT",
				},
				Paragraph: para,
			}

			// Replace with value
			err := mm.replaceComplexField(
				field,
				"Jane",
			)
			if err != nil {
				t.Fatalf(
					"replaceComplexField failed: %v",
					err,
				)
			}

			// Verify all field runs were removed
			for child := range para.Children() {
				if child == beginRun ||
					child == instrRun ||
					child == separateRun ||
					child == resultRun ||
					child == endRun {
					t.Errorf(
						"Field run should have been removed: %v",
						child.LocalName(),
					)
				}
			}

			// Verify a Run was added with the value
			foundRun := false
			for child := range para.Children() {
				if run, ok := child.(*elements.Run); ok {
					text := run.InnerText()
					if text == "Jane" {
						foundRun = true

						break
					}
				}
			}
			if !foundRun {
				t.Error(
					"Expected to find Run with text 'Jane'",
				)
			}
		},
	)

	t.Run(
		"Replace complex field without separate run",
		func(t *testing.T) {
			// Create a paragraph with a ComplexField without separate/result
			para := elements.NewParagraph()

			// Begin run
			beginRun := elements.NewRun("")
			beginChar := elements.NewFieldChar(
				elements.FieldCharBegin,
			)
			beginRun.AppendChild(beginChar)
			para.AppendChild(beginRun)

			// InstrText run
			instrRun := elements.NewRun("")
			instrText := elements.NewInstrText(
				"MERGEFIELD LastName",
			)
			instrRun.AppendChild(instrText)
			para.AppendChild(instrRun)

			// End run (no separate or result)
			endRun := elements.NewRun("")
			endChar := elements.NewFieldChar(
				elements.FieldCharEnd,
			)
			endRun.AppendChild(endChar)
			para.AppendChild(endRun)

			// Create merge field structure
			field := &MergeField{
				FieldName:   "LastName",
				FieldType:   FieldTypeComplex,
				Element:     beginRun,
				BeginRun:    beginRun,
				SeparateRun: nil, // No separate
				EndRun:      endRun,
				FieldCode:   "MERGEFIELD LastName",
				Switches:    map[string]string{},
				Paragraph:   para,
			}

			// Replace with value
			err := mm.replaceComplexField(
				field,
				"Doe",
			)
			if err != nil {
				t.Fatalf(
					"replaceComplexField failed: %v",
					err,
				)
			}

			// Verify a Run was added with the value
			foundRun := false
			for child := range para.Children() {
				if run, ok := child.(*elements.Run); ok {
					text := run.InnerText()
					if text == "Doe" {
						foundRun = true

						break
					}
				}
			}
			if !foundRun {
				t.Error(
					"Expected to find Run with text 'Doe'",
				)
			}
		},
	)
}

func TestValidateFields(t *testing.T) {
	t.Run(
		"All fields present - strict mode",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{
						"FirstName": "John",
						"LastName":  "Doe",
					},
				},
			)
			_ = ds.Open()

			mm := New(nil)
			mm.options.StrictFields = true

			fields := []*MergeField{
				{FieldName: "FirstName"},
				{FieldName: "LastName"},
			}

			err := mm.validateFields(fields, ds)
			if err != nil {
				t.Errorf(
					"validateFields should pass when all fields present, got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Missing field - strict mode",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{"FirstName": "John"},
				},
			)
			_ = ds.Open()

			mm := New(nil)
			mm.options.StrictFields = true

			fields := []*MergeField{
				{FieldName: "FirstName"},
				{
					FieldName: "LastName",
				}, // Missing in data source
			}

			err := mm.validateFields(fields, ds)
			if err == nil {
				t.Error(
					"validateFields should fail when field is missing in strict mode",
				)
			}
			if !strings.Contains(
				err.Error(),
				"LastName",
			) {
				t.Errorf(
					"Error should mention missing field 'LastName', got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Missing field - non-strict mode",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{"FirstName": "John"},
				},
			)
			_ = ds.Open()

			mm := New(nil)
			mm.options.StrictFields = false

			fields := []*MergeField{
				{FieldName: "FirstName"},
				{
					FieldName: "LastName",
				}, // Missing but should be okay
			}

			err := mm.validateFields(fields, ds)
			if err != nil {
				t.Errorf(
					"validateFields should pass in non-strict mode, got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"GREETINGLINE field skipped",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{"FirstName": "John"},
				},
			)
			_ = ds.Open()

			mm := New(nil)
			mm.options.StrictFields = true

			fields := []*MergeField{
				{FieldName: "FirstName"},
				{
					FieldName: "GREETINGLINE",
				}, // Special field, should be skipped
			}

			err := mm.validateFields(fields, ds)
			if err != nil {
				t.Errorf(
					"validateFields should skip GREETINGLINE field, got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Case insensitive field matching",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{
						"firstname": "John",
					}, // lowercase in data source
				},
			)
			_ = ds.Open()

			mm := New(nil)
			mm.options.StrictFields = true

			fields := []*MergeField{
				{
					FieldName: "FirstName",
				}, // Mixed case in field
			}

			err := mm.validateFields(fields, ds)
			if err != nil {
				t.Errorf(
					"validateFields should match case-insensitively, got: %v",
					err,
				)
			}
		},
	)
}

func TestGetFieldValue(t *testing.T) {
	t.Run(
		"Get regular field value",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{"FirstName": "John"},
				},
			)
			_ = ds.Open()
			ds.Next()

			mm := New(nil)

			field := &MergeField{
				FieldName: "FirstName",
			}

			value, err := mm.getFieldValue(
				field,
				ds,
			)
			if err != nil {
				t.Fatalf(
					"getFieldValue failed: %v",
					err,
				)
			}
			if value != "John" {
				t.Errorf(
					"Expected value 'John', got %q",
					value,
				)
			}
		},
	)

	t.Run(
		"Get GREETINGLINE field",
		func(t *testing.T) {
			ds := NewMapDataSource(
				[]map[string]string{
					{"FirstName": "John"},
				},
			)
			_ = ds.Open()
			ds.Next()

			mm := New(nil)

			field := &MergeField{
				FieldName: "GREETINGLINE",
			}

			value, err := mm.getFieldValue(
				field,
				ds,
			)
			if err != nil {
				t.Fatalf(
					"getFieldValue failed: %v",
					err,
				)
			}
			// Should return a greeting (current implementation returns simple greeting)
			if value == "" {
				t.Error(
					"GREETINGLINE should return a non-empty value",
				)
			}
		},
	)
}

func TestReplaceFields(t *testing.T) {
	t.Run(
		"Replace multiple simple fields",
		func(t *testing.T) {
			mm := New(nil)

			// Create paragraph with two SimpleFields
			para := elements.NewParagraph()

			field1 := elements.NewSimpleField()
			field1.Instruction = types.NewStringValue(
				"MERGEFIELD FirstName",
			)
			para.AppendChild(field1)

			field2 := elements.NewSimpleField()
			field2.Instruction = types.NewStringValue(
				"MERGEFIELD LastName",
			)
			para.AppendChild(field2)

			// Create merge field structures
			fields := []*MergeField{
				{
					FieldName: "FirstName",
					FieldType: FieldTypeSimple,
					Element:   field1,
					FieldCode: "MERGEFIELD FirstName",
					Switches:  map[string]string{},
					Paragraph: para,
				},
				{
					FieldName: "LastName",
					FieldType: FieldTypeSimple,
					Element:   field2,
					FieldCode: "MERGEFIELD LastName",
					Switches:  map[string]string{},
					Paragraph: para,
				},
			}

			// Replace with getValue function
			getValue := func(fieldName string) (string, error) {
				if fieldName == "FirstName" {
					return "Alice", nil
				}
				if fieldName == "LastName" {
					return "Smith", nil
				}

				return "", nil
			}

			err := mm.replaceFields(
				fields,
				getValue,
			)
			if err != nil {
				t.Fatalf(
					"replaceFields failed: %v",
					err,
				)
			}

			// Verify both fields were replaced
			text := ""
			for child := range para.Children() {
				if run, ok := child.(*elements.Run); ok {
					text += run.InnerText()
				}
			}

			if !strings.Contains(text, "Alice") ||
				!strings.Contains(text, "Smith") {
				t.Errorf(
					"Expected text to contain 'Alice' and 'Smith', got: %q",
					text,
				)
			}
		},
	)

	t.Run(
		"Replace with format switch",
		func(t *testing.T) {
			mm := New(nil)

			para := elements.NewParagraph()
			field := elements.NewSimpleField()
			field.Instruction = types.NewStringValue(
				"MERGEFIELD Name \\* UPPER",
			)
			para.AppendChild(field)

			fields := []*MergeField{
				{
					FieldName: "Name",
					FieldType: FieldTypeSimple,
					Element:   field,
					FieldCode: "MERGEFIELD Name \\* UPPER",
					Switches: map[string]string{
						"\\*": "UPPER",
					},
					Paragraph: para,
				},
			}

			getValue := func(fieldName string) (string, error) {
				return "john doe", nil
			}

			err := mm.replaceFields(
				fields,
				getValue,
			)
			if err != nil {
				t.Fatalf(
					"replaceFields failed: %v",
					err,
				)
			}

			// Verify uppercase transformation
			text := ""
			for child := range para.Children() {
				if run, ok := child.(*elements.Run); ok {
					text += run.InnerText()
				}
			}

			if text != "JOHN DOE" {
				t.Errorf(
					"Expected 'JOHN DOE', got: %q",
					text,
				)
			}
		},
	)
}

func TestExecute(t *testing.T) {
	t.Run(
		"Execute with no data source",
		func(t *testing.T) {
			// Create a minimal document template
			doc, err := createMinimalDocument(t)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Create MailMerge without data source
			mm := New(doc)

			// Execute should fail with ErrNoDataSource
			_, err = mm.Execute()
			if err == nil {
				t.Error(
					"Expected error when executing without data source",
				)
			}
			if err != ErrNoDataSource {
				t.Errorf(
					"Expected ErrNoDataSource, got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Execute with zero records",
		func(t *testing.T) {
			// Create a minimal document
			doc, err := createMinimalDocument(t)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Create empty data source
			ds := NewMapDataSource(
				[]map[string]string{},
			) // Empty slice - zero records

			// Create MailMerge with empty data source
			mm := New(doc).DataSource(ds)

			// Execute should fail with ErrNoRecords
			_, err = mm.Execute()
			if err == nil {
				t.Error(
					"Expected error when data source has zero records",
				)
			}
			if err != ErrNoRecords {
				t.Errorf(
					"Expected ErrNoRecords, got: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Execute with simple template",
		func(t *testing.T) {
			// Create a document with a merge field (using temp file approach)
			tmpFile, err := os.CreateTemp(
				"",
				"test_execute_*.docx",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create temp file: %v",
					err,
				)
			}
			tmpPath := tmpFile.Name()
			_ = tmpFile.Close()
			defer func() { _ = os.Remove(tmpPath) }()

			// Create document
			doc, err := wordprocessing.New(
				tmpPath,
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Build document structure
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			// Add paragraph with merge field
			para := elements.NewParagraph()
			field := elements.NewSimpleField()
			field.Instruction = types.NewStringValue(
				"MERGEFIELD Name",
			)
			run := elements.NewRun("«Name»")
			field.AppendChild(run)
			para.AppendChild(field)
			body.AppendChild(para)

			// Create CSV file with test data
			csvFile, err := os.CreateTemp(
				"",
				"test_data_*.csv",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create CSV file: %v",
					err,
				)
			}
			csvPath := csvFile.Name()
			_, _ = csvFile.WriteString(
				"Name\nJohn Doe\n",
			)
			_ = csvFile.Close()
			defer func() { _ = os.Remove(csvPath) }()

			// Execute merge
			mm := New(doc).DataSource(
				NewCSVDataSource(csvPath),
			)
			result, err := mm.Execute()
			if err != nil {
				t.Fatalf(
					"Execute failed: %v",
					err,
				)
			}
			if result == nil {
				t.Fatal(
					"Execute returned nil document",
				)
			}

			// Verify the field was replaced
			text := extractDocumentText(result)
			if !strings.Contains(
				text,
				"John Doe",
			) {
				t.Errorf(
					"Expected text to contain 'John Doe', got: %q",
					text,
				)
			}
			if strings.Contains(text, "«Name»") {
				t.Errorf(
					"Text should not contain placeholder '«Name»', got: %q",
					text,
				)
			}
		},
	)
}

// Helper function to create a minimal document for testing
func createMinimalDocument(
	t *testing.T,
) (*wordprocessing.Document, error) {
	t.Helper()

	// Create a temporary file
	tmpFile, err := os.CreateTemp(
		"",
		"test_minimal_*.docx",
	)
	if err != nil {
		return nil, err
	}
	tmpPath := tmpFile.Name()
	_ = tmpFile.Close()

	// Create a new document
	doc, err := wordprocessing.New(
		tmpPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		return nil, err
	}

	// Clean up temp file on test cleanup
	t.Cleanup(
		func() { _ = os.Remove(tmpPath) },
	)

	return doc, nil
}

// Helper function to extract text from a document
func extractDocumentText(
	doc *wordprocessing.Document,
) string {
	mainPart := doc.MainPart()
	if mainPart == nil {
		return ""
	}

	docElem := mainPart.Document()
	if docElem == nil {
		return ""
	}

	body := docElem.GetOrCreateBody()

	var text strings.Builder
	for child := range body.Children() {
		para, ok := child.(*elements.Paragraph)
		if !ok {
			continue
		}

		for paraChild := range para.Children() {
			switch elem := paraChild.(type) {
			case *elements.Run:
				text.WriteString(elem.InnerText())
			case *elements.SimpleField:
				// Extract text from SimpleField children
				for fieldChild := range elem.Children() {
					run, ok := fieldChild.(*elements.Run)
					if ok {
						text.WriteString(
							run.InnerText(),
						)
					}
				}
			}
		}
	}

	return text.String()
}
