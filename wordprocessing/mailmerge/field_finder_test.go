// Copyright 2024 goffice authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package mailmerge

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestParseMergeFieldCode(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		expectedField  string
		expectedSwitch map[string]string
	}{
		{
			name:          "Simple MERGEFIELD",
			code:          "MERGEFIELD FirstName",
			expectedField: "FirstName",
			expectedSwitch: make(
				map[string]string,
			),
		},
		{
			name:          "MERGEFIELD with MERGEFORMAT",
			code:          "MERGEFIELD FirstName \\* MERGEFORMAT",
			expectedField: "FirstName",
			expectedSwitch: map[string]string{
				"\\*": "MERGEFORMAT",
			},
		},
		{
			name:          "MERGEFIELD with multiple switches",
			code:          "MERGEFIELD Amount \\* MERGEFORMAT \\# \"$#,##0.00\"",
			expectedField: "Amount",
			expectedSwitch: map[string]string{
				"\\*": "MERGEFORMAT",
				"\\#": "$#,##0.00",
			},
		},
		{
			name:          "GREETINGLINE",
			code:          "GREETINGLINE",
			expectedField: "GREETINGLINE",
			expectedSwitch: make(
				map[string]string,
			),
		},
		{
			name:          "GREETINGLINE with text",
			code:          "GREETINGLINE \\f \"Dear \"",
			expectedField: "GREETINGLINE",
			expectedSwitch: map[string]string{
				"\\f": "Dear ",
			},
		},
		{
			name:          "Not a merge field",
			code:          "PAGE",
			expectedField: "",
			expectedSwitch: make(
				map[string]string,
			),
		},
		{
			name:          "Empty code",
			code:          "",
			expectedField: "",
			expectedSwitch: make(
				map[string]string,
			),
		},
		{
			name:          "MERGEFIELD only (no field name)",
			code:          "MERGEFIELD",
			expectedField: "",
			expectedSwitch: make(
				map[string]string,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldName, switches := parseMergeFieldCode(
				tt.code,
			)

			if fieldName != tt.expectedField {
				t.Errorf(
					"Expected field name %q, got %q",
					tt.expectedField,
					fieldName,
				)
			}

			if len(
				switches,
			) != len(
				tt.expectedSwitch,
			) {
				t.Errorf(
					"Expected %d switches, got %d",
					len(tt.expectedSwitch),
					len(switches),
				)
			}

			for key, expectedValue := range tt.expectedSwitch {
				if actualValue, ok := switches[key]; !ok {
					t.Errorf(
						"Expected switch %q not found",
						key,
					)
				} else if actualValue != expectedValue {
					t.Errorf("Switch %q: expected value %q, got %q", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestTokenizeFieldCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected []string
	}{
		{
			name: "Simple field",
			code: "MERGEFIELD Name",
			expected: []string{
				"MERGEFIELD",
				"Name",
			},
		},
		{
			name: "Field with switch",
			code: "MERGEFIELD Name \\* MERGEFORMAT",
			expected: []string{
				"MERGEFIELD",
				"Name",
				"\\*",
				"MERGEFORMAT",
			},
		},
		{
			name: "Field with quoted string",
			code: "MERGEFIELD Name \\f \"Dear \"",
			expected: []string{
				"MERGEFIELD",
				"Name",
				"\\f",
				"Dear ",
			},
		},
		{
			name: "Field with quoted string containing spaces",
			code: "GREETINGLINE \\f \"Dear Sir or Madam,\"",
			expected: []string{
				"GREETINGLINE",
				"\\f",
				"Dear Sir or Madam,",
			},
		},
		{
			name: "Field with multiple quoted strings",
			code: "MERGEFIELD Name \\f \"Mr. \" \\l \"Esq.\"",
			expected: []string{
				"MERGEFIELD",
				"Name",
				"\\f",
				"Mr. ",
				"\\l",
				"Esq.",
			},
		},
		{
			name: "Extra whitespace",
			code: "MERGEFIELD   Name   \\*   MERGEFORMAT",
			expected: []string{
				"MERGEFIELD",
				"Name",
				"\\*",
				"MERGEFORMAT",
			},
		},
		{
			name:     "Empty string",
			code:     "",
			expected: make([]string, 0),
		},
		{
			name: "Field with number format",
			code: "MERGEFIELD Amount \\# \"$#,##0.00\"",
			expected: []string{
				"MERGEFIELD",
				"Amount",
				"\\#",
				"$#,##0.00",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizeFieldCode(tt.code)

			if len(tokens) != len(tt.expected) {
				t.Errorf(
					"Expected %d tokens, got %d\nExpected: %v\nGot: %v",
					len(
						tt.expected,
					),
					len(tokens),
					tt.expected,
					tokens,
				)

				return
			}

			for i, expected := range tt.expected {
				if tokens[i] != expected {
					t.Errorf(
						"Token %d: expected %q, got %q",
						i,
						expected,
						tokens[i],
					)
				}
			}
		})
	}
}

func TestParseFieldSwitches(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected map[string]string
	}{
		{
			name: "Single switch with value",
			tokens: []string{
				"\\*",
				"MERGEFORMAT",
			},
			expected: map[string]string{
				"\\*": "MERGEFORMAT",
			},
		},
		{
			name: "Multiple switches",
			tokens: []string{
				"\\*",
				"MERGEFORMAT",
				"\\#",
				"$#,##0.00",
			},
			expected: map[string]string{
				"\\*": "MERGEFORMAT",
				"\\#": "$#,##0.00",
			},
		},
		{
			name: "Switch without value",
			tokens: []string{
				"\\*",
				"\\#",
				"0.00",
			},
			expected: map[string]string{
				"\\*": "",
				"\\#": "0.00",
			},
		},
		{
			name:     "No switches",
			tokens:   make([]string, 0),
			expected: make(map[string]string),
		},
		{
			name: "Non-switch tokens only",
			tokens: []string{
				"MERGEFORMAT",
				"Name",
			},
			expected: make(map[string]string),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switches := parseFieldSwitches(
				tt.tokens,
			)

			if len(switches) != len(tt.expected) {
				t.Errorf(
					"Expected %d switches, got %d",
					len(tt.expected),
					len(switches),
				)
			}

			for key, expectedValue := range tt.expected {
				if actualValue, ok := switches[key]; !ok {
					t.Errorf(
						"Expected switch %q not found",
						key,
					)
				} else if actualValue != expectedValue {
					t.Errorf("Switch %q: expected value %q, got %q", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestFindMergeFields_SimpleField(
	t *testing.T,
) {
	// Create a new document
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Get the main document part
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("Main document part is nil")
	}

	docElem := mainPart.Document()
	body := docElem.Body()

	// Create a paragraph with a SimpleField
	para := elements.NewParagraph()
	body.AppendChild(para)

	// Create a SimpleField
	simpleField := elements.NewSimpleField()
	simpleField.Instruction = &types.StringValue{}
	simpleField.Instruction.SetValue(
		"MERGEFIELD FirstName \\* MERGEFORMAT",
	)
	para.AppendChild(simpleField)

	// Add result text to the simple field
	run := elements.NewRun("")
	text := elements.NewText("«FirstName»")
	run.AppendChild(text)
	simpleField.AppendChild(run)

	// Create a MailMerge instance and find fields
	mm := New(doc)
	fields := mm.findMergeFields()

	// Verify we found the field
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 field, got %d",
			len(fields),
		)

		return
	}

	field := fields[0]
	if field.FieldName != "FirstName" {
		t.Errorf(
			"Expected field name 'FirstName', got %q",
			field.FieldName,
		)
	}

	if field.FieldType != FieldTypeSimple {
		t.Errorf(
			"Expected FieldTypeSimple, got %v",
			field.FieldType,
		)
	}

	if field.FieldCode != "MERGEFIELD FirstName \\* MERGEFORMAT" {
		t.Errorf(
			"Expected field code 'MERGEFIELD FirstName \\* MERGEFORMAT', got %q",
			field.FieldCode,
		)
	}

	if field.Element != simpleField {
		t.Error(
			"Field.Element does not match the SimpleField element",
		)
	}

	if field.Paragraph != para {
		t.Error(
			"Field.Paragraph does not match the containing paragraph",
		)
	}
}

func TestFindMergeFields_ComplexField(
	t *testing.T,
) {
	// Create a new document
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Get the main document part
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("Main document part is nil")
	}

	docElem := mainPart.Document()
	body := docElem.Body()

	// Create a paragraph with a ComplexField
	para := elements.NewParagraph()
	body.AppendChild(para)

	// Create begin run with fldChar
	beginRun := elements.NewRun("")
	beginChar := elements.NewFieldChar(
		elements.FieldCharBegin,
	)
	beginRun.AppendChild(beginChar)
	para.AppendChild(beginRun)

	// Create instrText run
	instrRun := elements.NewRun("")
	instrText := elements.NewInstrText(
		"MERGEFIELD LastName \\* MERGEFORMAT",
	)
	instrRun.AppendChild(instrText)
	para.AppendChild(instrRun)

	// Create separate run with fldChar
	separateRun := elements.NewRun("")
	separateChar := elements.NewFieldChar(
		elements.FieldCharSeparate,
	)
	separateRun.AppendChild(separateChar)
	para.AppendChild(separateRun)

	// Create result run
	resultRun := elements.NewRun("")
	resultText := elements.NewText("«LastName»")
	resultRun.AppendChild(resultText)
	para.AppendChild(resultRun)

	// Create end run with fldChar
	endRun := elements.NewRun("")
	endChar := elements.NewFieldChar(
		elements.FieldCharEnd,
	)
	endRun.AppendChild(endChar)
	para.AppendChild(endRun)

	// Create a MailMerge instance and find fields
	mm := New(doc)
	fields := mm.findMergeFields()

	// Verify we found the field
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 field, got %d",
			len(fields),
		)

		return
	}

	field := fields[0]
	if field.FieldName != "LastName" {
		t.Errorf(
			"Expected field name 'LastName', got %q",
			field.FieldName,
		)
	}

	if field.FieldType != FieldTypeComplex {
		t.Errorf(
			"Expected FieldTypeComplex, got %v",
			field.FieldType,
		)
	}

	if field.FieldCode != "MERGEFIELD LastName \\* MERGEFORMAT" {
		t.Errorf(
			"Expected field code 'MERGEFIELD LastName \\* MERGEFORMAT', got %q",
			field.FieldCode,
		)
	}

	if field.BeginRun != beginRun {
		t.Error("Field.BeginRun does not match")
	}

	if field.SeparateRun != separateRun {
		t.Error(
			"Field.SeparateRun does not match",
		)
	}

	if field.EndRun != endRun {
		t.Error("Field.EndRun does not match")
	}

	if field.Paragraph != para {
		t.Error(
			"Field.Paragraph does not match the containing paragraph",
		)
	}
}

func TestFindMergeFields_ComplexFieldMultipleInstrText(
	t *testing.T,
) {
	// Create a new document
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Get the main document part
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("Main document part is nil")
	}

	docElem := mainPart.Document()
	body := docElem.Body()

	// Create a paragraph with a ComplexField with split instrText
	para := elements.NewParagraph()
	body.AppendChild(para)

	// Create begin run with fldChar
	beginRun := elements.NewRun("")
	beginChar := elements.NewFieldChar(
		elements.FieldCharBegin,
	)
	beginRun.AppendChild(beginChar)
	para.AppendChild(beginRun)

	// Create first instrText run (split across multiple runs)
	instrRun1 := elements.NewRun("")
	instrText1 := elements.NewInstrText(
		"MERGEFIELD",
	)
	instrRun1.AppendChild(instrText1)
	para.AppendChild(instrRun1)

	// Create second instrText run
	instrRun2 := elements.NewRun("")
	instrText2 := elements.NewInstrText(
		"CompanyName",
	)
	instrRun2.AppendChild(instrText2)
	para.AppendChild(instrRun2)

	// Create third instrText run
	instrRun3 := elements.NewRun("")
	instrText3 := elements.NewInstrText(
		"\\* MERGEFORMAT",
	)
	instrRun3.AppendChild(instrText3)
	para.AppendChild(instrRun3)

	// Create separate run with fldChar
	separateRun := elements.NewRun("")
	separateChar := elements.NewFieldChar(
		elements.FieldCharSeparate,
	)
	separateRun.AppendChild(separateChar)
	para.AppendChild(separateRun)

	// Create result run
	resultRun := elements.NewRun("")
	resultText := elements.NewText(
		"«CompanyName»",
	)
	resultRun.AppendChild(resultText)
	para.AppendChild(resultRun)

	// Create end run with fldChar
	endRun := elements.NewRun("")
	endChar := elements.NewFieldChar(
		elements.FieldCharEnd,
	)
	endRun.AppendChild(endChar)
	para.AppendChild(endRun)

	// Create a MailMerge instance and find fields
	mm := New(doc)
	fields := mm.findMergeFields()

	// Verify we found the field
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 field, got %d",
			len(fields),
		)

		return
	}

	field := fields[0]
	if field.FieldName != "CompanyName" {
		t.Errorf(
			"Expected field name 'CompanyName', got %q",
			field.FieldName,
		)
	}

	if field.FieldCode != "MERGEFIELD CompanyName \\* MERGEFORMAT" {
		t.Errorf(
			"Expected field code 'MERGEFIELD CompanyName \\* MERGEFORMAT', got %q",
			field.FieldCode,
		)
	}
}

func TestFindMergeFields_ComplexFieldWithoutSeparate(
	t *testing.T,
) {
	// Create a new document
	doc, err := wordprocessing.New(
		"",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Get the main document part
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("Main document part is nil")
	}

	docElem := mainPart.Document()
	body := docElem.Body()

	// Create a paragraph with a ComplexField without separate marker
	para := elements.NewParagraph()
	body.AppendChild(para)

	// Create begin run with fldChar
	beginRun := elements.NewRun("")
	beginChar := elements.NewFieldChar(
		elements.FieldCharBegin,
	)
	beginRun.AppendChild(beginChar)
	para.AppendChild(beginRun)

	// Create instrText run
	instrRun := elements.NewRun("")
	instrText := elements.NewInstrText(
		"MERGEFIELD Email",
	)
	instrRun.AppendChild(instrText)
	para.AppendChild(instrRun)

	// Create end run with fldChar (no separate)
	endRun := elements.NewRun("")
	endChar := elements.NewFieldChar(
		elements.FieldCharEnd,
	)
	endRun.AppendChild(endChar)
	para.AppendChild(endRun)

	// Create a MailMerge instance and find fields
	mm := New(doc)
	fields := mm.findMergeFields()

	// Verify we found the field
	if len(fields) != 1 {
		t.Errorf(
			"Expected 1 field, got %d",
			len(fields),
		)

		return
	}

	field := fields[0]
	if field.FieldName != "Email" {
		t.Errorf(
			"Expected field name 'Email', got %q",
			field.FieldName,
		)
	}

	if field.SeparateRun != nil {
		t.Error("Expected SeparateRun to be nil")
	}
}

func TestReconstructFieldCode(t *testing.T) {
	tests := []struct {
		name     string
		parts    []string
		expected string
	}{
		{
			name:     "Single part",
			parts:    []string{"MERGEFIELD Name"},
			expected: "MERGEFIELD Name",
		},
		{
			name: "Multiple parts",
			parts: []string{
				"MERGEFIELD",
				"Name",
				"\\* MERGEFORMAT",
			},
			expected: "MERGEFIELD Name \\* MERGEFORMAT",
		},
		{
			name: "Parts with extra whitespace",
			parts: []string{
				"MERGEFIELD  ",
				"  Name  ",
				"  \\* MERGEFORMAT",
			},
			expected: "MERGEFIELD Name \\* MERGEFORMAT",
		},
		{
			name:     "Empty parts",
			parts:    make([]string, 0),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reconstructFieldCode(
				tt.parts,
			)
			if result != tt.expected {
				t.Errorf(
					"Expected %q, got %q",
					tt.expected,
					result,
				)
			}
		})
	}
}
