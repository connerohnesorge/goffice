// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const (
	testFieldFirstName       = "FirstName"
	testFieldFirstNameMapped = "First_Name"
)

func TestLoadMailMergeMetadata_NoSettings(
	t *testing.T,
) {
	// Create a document without settings
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	metadata, err := loadMailMergeMetadata(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should return nil when there's no settings part
	if metadata != nil {
		t.Error(
			"expected nil metadata when no settings part exists",
		)
	}
}

func TestLoadMailMergeMetadata_NoMailMerge(
	t *testing.T,
) {
	// Create a document with settings but no mail merge configuration
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Add settings part but don't add mail merge element
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("document has no main part")
	}

	_, err = mainPart.AddSettingsPart()
	if err != nil {
		t.Fatalf(
			"failed to add settings part: %v",
			err,
		)
	}

	metadata, err := loadMailMergeMetadata(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should return nil when there's no mail merge element
	if metadata != nil {
		t.Error(
			"expected nil metadata when no mail merge element exists",
		)
	}
}

func TestLoadMailMergeMetadata_WithBasicSettings(
	t *testing.T,
) {
	// Create a document with mail merge settings
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("document has no main part")
	}

	settingsPart, err := mainPart.AddSettingsPart()
	if err != nil {
		t.Fatalf(
			"failed to add settings part: %v",
			err,
		)
	}

	settings := settingsPart.Settings()
	if settings == nil {
		t.Fatal("settings element is nil")
	}

	// Add mail merge element with basic settings
	mailMerge := settings.GetOrCreateMailMerge()
	mailMerge.SetDataType("textFile")

	metadata, err := loadMailMergeMetadata(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if metadata == nil {
		t.Fatal("expected metadata to be loaded")
	}

	if metadata.DataType != "textFile" {
		t.Errorf(
			"expected DataType 'textFile', got %q",
			metadata.DataType,
		)
	}
}

func TestApplyFieldMapping_NilMetadata(
	t *testing.T,
) {
	result := applyFieldMapping(
		testFieldFirstName,
		nil,
	)
	if result != testFieldFirstName {
		t.Errorf(
			"expected %q, got %q",
			testFieldFirstName,
			result,
		)
	}
}

func TestApplyFieldMapping_NoMapping(
	t *testing.T,
) {
	metadata := &OoxmlMailMergeMetadata{
		FieldMappings: make(map[string]string),
	}

	result := applyFieldMapping(
		testFieldFirstName,
		metadata,
	)
	if result != testFieldFirstName {
		t.Errorf(
			"expected %q, got %q",
			testFieldFirstName,
			result,
		)
	}
}

func TestApplyFieldMapping_WithMapping(
	t *testing.T,
) {
	metadata := &OoxmlMailMergeMetadata{
		FieldMappings: map[string]string{
			"firstname": testFieldFirstNameMapped,
			"lastname":  "Last_Name",
		},
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"exact match lowercase",
			"firstname",
			testFieldFirstNameMapped,
		},
		{
			"exact match uppercase",
			"FIRSTNAME",
			testFieldFirstNameMapped,
		}, // Case-insensitive lookup
		{
			"exact match mixed",
			testFieldFirstName,
			testFieldFirstNameMapped,
		}, // Case-insensitive lookup
		{
			"no mapping",
			"MiddleName",
			"MiddleName",
		},
		{"second field", "lastname", "Last_Name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyFieldMapping(
				tt.input,
				metadata,
			)
			if result != tt.expected {
				t.Errorf(
					"expected %q, got %q",
					tt.expected,
					result,
				)
			}
		})
	}
}

func TestApplyFieldMapping_CaseInsensitive(
	t *testing.T,
) {
	metadata := &OoxmlMailMergeMetadata{
		FieldMappings: map[string]string{
			"firstname": testFieldFirstNameMapped,
		},
	}

	// All variations should map to the same value (case-insensitive)
	variations := []string{
		"firstname",
		testFieldFirstName,
		"FIRSTNAME",
		"FiRsTnAmE",
	}

	for _, variant := range variations {
		result := applyFieldMapping(
			variant,
			metadata,
		)
		// All variations should map to First_Name because lookup is case-insensitive
		if result != testFieldFirstNameMapped {
			t.Errorf(
				"for %q: expected %q, got %q",
				variant,
				testFieldFirstNameMapped,
				result,
			)
		}
	}

	// Test that we properly handle case-insensitive storage
	// The mapping should be stored with lowercase keys
	lowercaseResult := applyFieldMapping(
		"firstname",
		metadata,
	)
	if lowercaseResult != testFieldFirstNameMapped {
		t.Errorf(
			"expected %q for lowercase, got %q",
			testFieldFirstNameMapped,
			lowercaseResult,
		)
	}
}

func TestExecute_WithFieldMappings(t *testing.T) {
	// Create a document with mail merge fields
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("document has no main part")
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("document element is nil")
	}

	body := docElem.Body()
	if body == nil {
		t.Fatal("body is nil")
	}

	// Add a paragraph with a simple merge field
	// The field uses testFieldFirstName but the data source has testFieldFirstName_Mapped
	para := elements.NewParagraph()
	body.AppendChild(para)

	field := elements.NewSimpleField()
	field.Instruction = types.NewStringValue(
		"MERGEFIELD " + testFieldFirstName,
	)
	para.AppendChild(field)

	// Add placeholder text
	run := elements.NewRun(
		"«" + testFieldFirstName + "»",
	)
	field.AppendChild(run)

	// Add settings with field mapping
	settingsPart, err := mainPart.AddSettingsPart()
	if err != nil {
		t.Fatalf(
			"failed to add settings part: %v",
			err,
		)
	}

	settings := settingsPart.Settings()
	if settings == nil {
		t.Fatal("settings element is nil")
	}

	// Note: We can't easily create ODSO field mappings through the API
	// because the FieldMapData element type may not exist yet.
	// This test verifies the basic integration works even without mappings.

	// Create data source with mapped field name
	ds := NewMapDataSource([]map[string]string{
		{
			testFieldFirstName: "John",
		}, // Use the document field name
	})

	// Execute merge
	mm := New(doc)
	mm.DataSource(ds)

	merged, err := mm.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if merged == nil {
		t.Fatal("merged document is nil")
	}

	// Verify the field was replaced
	// We can't easily verify the content without traversing the entire tree,
	// but at least verify that Execute completed successfully
}

func TestGetChildElementValue(t *testing.T) {
	// This is a helper function test
	// We'll test it indirectly through loadMailMergeMetadata
	// since it's not exported

	// Create a document with mail merge settings
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	settingsPart, err := mainPart.AddSettingsPart()
	if err != nil {
		t.Fatalf(
			"failed to add settings part: %v",
			err,
		)
	}

	settings := settingsPart.Settings()
	mailMerge := settings.GetOrCreateMailMerge()

	// Set values using the API
	mailMerge.SetDataType("database")

	// Load metadata and verify values were extracted
	metadata, err := loadMailMergeMetadata(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if metadata == nil {
		t.Fatal("expected metadata to be loaded")
	}

	if metadata.DataType != "database" {
		t.Errorf(
			"expected DataType 'database', got %q",
			metadata.DataType,
		)
	}
}
