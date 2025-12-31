package wordprocessing

import (
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestFormFieldType(t *testing.T) {
	tests := []struct {
		name     string
		fType    FormFieldType
		expected string
	}{
		{"Text", FormFieldTypeText, "Text"},
		{
			"CheckBox",
			FormFieldTypeCheckBox,
			"CheckBox",
		},
		{
			"DropDown",
			FormFieldTypeDropDown,
			"DropDown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fType.String() != tt.expected {
				t.Errorf(
					"FormFieldType.String() = %v, want %v",
					tt.fType.String(),
					tt.expected,
				)
			}
		})
	}
}

func TestFormField_Name(t *testing.T) {
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// Test getting name when not set
	if name := ff.Name(); name != "" {
		t.Errorf(
			"Name() = %v, want empty string",
			name,
		)
	}

	// Test setting and getting name
	ff.SetName("TestField")
	if name := ff.Name(); name != "TestField" {
		t.Errorf(
			"Name() = %v, want TestField",
			name,
		)
	}
}

func TestFormField_Type(t *testing.T) {
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	tests := []struct {
		name      string
		fieldType FormFieldType
	}{
		{"Text", FormFieldTypeText},
		{"CheckBox", FormFieldTypeCheckBox},
		{"DropDown", FormFieldTypeDropDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ff := newFormField(
				simpleField,
				ffData,
				para,
				tt.fieldType,
			)
			if ff.Type() != tt.fieldType {
				t.Errorf(
					"Type() = %v, want %v",
					ff.Type(),
					tt.fieldType,
				)
			}
		})
	}
}

func TestFormField_Enabled(t *testing.T) {
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// Test default enabled state
	if !ff.Enabled() {
		t.Error(
			"Enabled() = false, want true (default)",
		)
	}

	// Test disabling
	ff.SetEnabled(false)
	if ff.Enabled() {
		t.Error("Enabled() = true, want false")
	}

	// Test enabling again
	ff.SetEnabled(true)
	if !ff.Enabled() {
		t.Error("Enabled() = false, want true")
	}
}

func TestFormField_TextField(t *testing.T) {
	simpleField := elements.NewSimpleField()
	simpleField.Run = elements.NewRun("")
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// Test setting and getting text value
	ff.SetTextValue("Hello World")
	if val := ff.GetTextValue(); val != "Hello World" {
		t.Errorf(
			"GetTextValue() = %v, want Hello World",
			val,
		)
	}

	// Test max length
	ff.SetMaxLength(50)
	if length := ff.GetMaxLength(); length != 50 {
		t.Errorf(
			"GetMaxLength() = %v, want 50",
			length,
		)
	}

	// Test default value
	ff.SetDefaultTextValue("Default Text")
	if val := ff.GetDefaultTextValue(); val != "Default Text" {
		t.Errorf(
			"GetDefaultTextValue() = %v, want Default Text",
			val,
		)
	}
}

func TestFormField_CheckBox(t *testing.T) {
	simpleField := elements.NewSimpleField()
	simpleField.Run = elements.NewRun("")
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeCheckBox,
	)

	// Test default unchecked state
	if ff.IsChecked() {
		t.Error(
			"IsChecked() = true, want false (default)",
		)
	}

	// Test checking
	ff.SetChecked(true)
	if !ff.IsChecked() {
		t.Error("IsChecked() = false, want true")
	}

	// Test unchecking
	ff.SetChecked(false)
	if ff.IsChecked() {
		t.Error("IsChecked() = true, want false")
	}

	// Test checkbox size
	ff.SetCheckBoxSize(24)
	if size := ff.GetCheckBoxSize(); size != 24 {
		t.Errorf(
			"GetCheckBoxSize() = %v, want 24",
			size,
		)
	}
}

func TestFormField_DropDown(t *testing.T) {
	simpleField := elements.NewSimpleField()
	simpleField.Run = elements.NewRun("")
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeDropDown,
	)

	// Test setting items
	items := []string{
		"Option 1",
		"Option 2",
		"Option 3",
	}
	ff.SetDropDownItems(items)

	// Test getting items
	gotItems := ff.GetDropDownItems()
	if len(gotItems) != len(items) {
		t.Errorf(
			"GetDropDownItems() length = %v, want %v",
			len(gotItems),
			len(items),
		)
	}
	for i, item := range items {
		if i >= len(gotItems) ||
			gotItems[i] != item {
			t.Errorf(
				"GetDropDownItems()[%d] = %v, want %v",
				i,
				gotItems[i],
				item,
			)
		}
	}

	// Test selected index
	ff.SetSelectedIndex(1)
	if idx := ff.GetSelectedIndex(); idx != 1 {
		t.Errorf(
			"GetSelectedIndex() = %v, want 1",
			idx,
		)
	}

	// Test selected value
	if val := ff.GetSelectedValue(); val != "Option 2" {
		t.Errorf(
			"GetSelectedValue() = %v, want Option 2",
			val,
		)
	}

	// Test invalid index
	ff.SetSelectedIndex(-1)
	if val := ff.GetSelectedValue(); val != "" {
		t.Errorf(
			"GetSelectedValue() with invalid index = %v, want empty",
			val,
		)
	}
}

func TestFormField_DropDown_25ItemLimit(
	t *testing.T,
) {
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeDropDown,
	)

	// Create 30 items (exceeds 25-item limit)
	items := make([]string, 30)
	for i := range 30 {
		items[i] = string(rune('A' + i))
	}

	ff.SetDropDownItems(items)

	// Should only have 25 items
	gotItems := ff.GetDropDownItems()
	if len(gotItems) != 25 {
		t.Errorf(
			"GetDropDownItems() length = %v, want 25 (enforced limit)",
			len(gotItems),
		)
	}
}

func TestFormField_Remove(t *testing.T) {
	para := elements.NewParagraph()
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()

	para.AppendChild(simpleField)

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// Verify field is in paragraph
	found := false
	for child := range para.Children() {
		if child == simpleField {
			found = true

			break
		}
	}
	if !found {
		t.Error(
			"SimpleField not found in paragraph before Remove()",
		)
	}

	// Remove the field
	ff.Remove()

	// Verify field is removed from paragraph
	found = false
	for child := range para.Children() {
		if child == simpleField {
			found = true

			break
		}
	}
	if found {
		t.Error(
			"SimpleField still in paragraph after Remove()",
		)
	}
}

func TestDetectFormFieldType(t *testing.T) {
	tests := []struct {
		name        string
		instruction string
		expected    FormFieldType
		wantErr     bool
	}{
		{
			"Text",
			"FORMTEXT",
			FormFieldTypeText,
			false,
		},
		{
			"CheckBox",
			"FORMCHECKBOX",
			FormFieldTypeCheckBox,
			false,
		},
		{
			"DropDown",
			"FORMDROPDOWN",
			FormFieldTypeDropDown,
			false,
		},
		{"Empty", "", FormFieldTypeText, true},
		{
			"Unknown",
			"UNKNOWN",
			FormFieldTypeText,
			true,
		},
		{
			"Partial",
			"FORM",
			FormFieldTypeText,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := detectFormFieldType(
				tt.instruction,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"detectFormFieldType() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)

				return
			}
			if !tt.wantErr && got != tt.expected {
				t.Errorf(
					"detectFormFieldType() = %v, want %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestFormField_TextFieldMethods_WrongType(
	t *testing.T,
) {
	// Test that text field methods do nothing on non-text fields
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeCheckBox,
	)

	// These should do nothing
	ff.SetTextValue("test")
	ff.SetMaxLength(10)
	ff.SetDefaultTextValue("default")

	if ff.GetTextValue() != "" {
		t.Error(
			"GetTextValue() on checkbox should return empty",
		)
	}
	if ff.GetMaxLength() != 0 {
		t.Error(
			"GetMaxLength() on checkbox should return 0",
		)
	}
	if ff.GetDefaultTextValue() != "" {
		t.Error(
			"GetDefaultTextValue() on checkbox should return empty",
		)
	}
}

func TestFormField_CheckBoxMethods_WrongType(
	t *testing.T,
) {
	// Test that checkbox methods do nothing on non-checkbox fields
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// These should do nothing
	ff.SetChecked(true)
	ff.SetCheckBoxSize(24)

	if ff.IsChecked() {
		t.Error(
			"IsChecked() on text field should return false",
		)
	}
	if ff.GetCheckBoxSize() != 0 {
		t.Error(
			"GetCheckBoxSize() on text field should return 0",
		)
	}
}

func TestFormField_DropDownMethods_WrongType(
	t *testing.T,
) {
	// Test that dropdown methods do nothing on non-dropdown fields
	simpleField := elements.NewSimpleField()
	ffData := elements.NewFormFieldData()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		ffData,
		para,
		FormFieldTypeText,
	)

	// These should do nothing
	ff.SetDropDownItems([]string{"A", "B"})
	ff.SetSelectedIndex(0)

	if items := ff.GetDropDownItems(); items != nil {
		t.Error(
			"GetDropDownItems() on text field should return nil",
		)
	}
	if idx := ff.GetSelectedIndex(); idx != -1 {
		t.Error(
			"GetSelectedIndex() on text field should return -1",
		)
	}
	if val := ff.GetSelectedValue(); val != "" {
		t.Error(
			"GetSelectedValue() on text field should return empty",
		)
	}
}

func TestFormField_NilFFData(t *testing.T) {
	// Test that methods handle nil ffData gracefully
	simpleField := elements.NewSimpleField()
	para := elements.NewParagraph()

	ff := newFormField(
		simpleField,
		nil,
		para,
		FormFieldTypeText,
	)

	// These should not panic
	ff.SetName("test")
	ff.SetEnabled(false)
	ff.SetMaxLength(10)
	ff.SetDefaultTextValue("test")
	ff.SetChecked(true)
	ff.SetCheckBoxSize(24)
	ff.SetDropDownItems([]string{"A"})
	ff.SetSelectedIndex(0)

	// Getters should return default values
	if name := ff.Name(); name != "" {
		t.Error(
			"Name() with nil ffData should return empty",
		)
	}
	if enabled := ff.Enabled(); !enabled {
		t.Error(
			"Enabled() with nil ffData should return true (default)",
		)
	}
}

func TestParagraph_InsertTextField(t *testing.T) {
	para := elements.NewParagraph()

	// Insert a text field
	ff := InsertTextField(
		para,
		"TextField1",
		"Default Value",
	)

	// Verify field was created
	if ff == nil {
		t.Fatal("InsertTextField returned nil")
	}

	// Verify field type
	if ff.Type() != FormFieldTypeText {
		t.Errorf(
			"Type() = %v, want %v",
			ff.Type(),
			FormFieldTypeText,
		)
	}

	// Verify field name
	if ff.Name() != "TextField1" {
		t.Errorf(
			"Name() = %v, want TextField1",
			ff.Name(),
		)
	}

	// Verify default value
	if ff.GetDefaultTextValue() != "Default Value" {
		t.Errorf(
			"GetDefaultTextValue() = %v, want Default Value",
			ff.GetDefaultTextValue(),
		)
	}

	// Verify the SimpleField was added to the paragraph
	found := false
	for child := range para.Children() {
		if child.LocalName() == fldSimpleLocalName {
			found = true

			break
		}
	}
	if !found {
		t.Error(
			"SimpleField not found in paragraph",
		)
	}
}

func TestParagraph_InsertCheckBox(t *testing.T) {
	tests := []struct {
		name    string
		checked bool
	}{
		{"Unchecked", false},
		{"Checked", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			para := elements.NewParagraph()

			// Insert a checkbox
			ff := InsertCheckBox(
				para,
				"CheckBox1",
				tt.checked,
			)

			// Verify field was created
			if ff == nil {
				t.Fatal(
					"InsertCheckBox returned nil",
				)
			}

			// Verify field type
			if ff.Type() != FormFieldTypeCheckBox {
				t.Errorf(
					"Type() = %v, want %v",
					ff.Type(),
					FormFieldTypeCheckBox,
				)
			}

			// Verify field name
			if ff.Name() != "CheckBox1" {
				t.Errorf(
					"Name() = %v, want CheckBox1",
					ff.Name(),
				)
			}

			// Verify checked state
			if ff.IsChecked() != tt.checked {
				t.Errorf(
					"IsChecked() = %v, want %v",
					ff.IsChecked(),
					tt.checked,
				)
			}

			// Verify the SimpleField was added to the paragraph
			found := false
			for child := range para.Children() {
				if child.LocalName() == fldSimpleLocalName {
					found = true

					break
				}
			}
			if !found {
				t.Error(
					"SimpleField not found in paragraph",
				)
			}
		})
	}
}

func TestParagraph_InsertDropDown(t *testing.T) {
	para := elements.NewParagraph()
	items := []string{
		"Option 1",
		"Option 2",
		"Option 3",
	}

	// Insert a dropdown
	ff := InsertDropDown(
		para,
		"DropDown1",
		items,
		1,
	)

	// Verify field was created
	if ff == nil {
		t.Fatal("InsertDropDown returned nil")
	}

	// Verify field type
	if ff.Type() != FormFieldTypeDropDown {
		t.Errorf(
			"Type() = %v, want %v",
			ff.Type(),
			FormFieldTypeDropDown,
		)
	}

	// Verify field name
	if ff.Name() != "DropDown1" {
		t.Errorf(
			"Name() = %v, want DropDown1",
			ff.Name(),
		)
	}

	// Verify items
	gotItems := ff.GetDropDownItems()
	if len(gotItems) != len(items) {
		t.Errorf(
			"GetDropDownItems() length = %v, want %v",
			len(gotItems),
			len(items),
		)
	}

	// Verify selected index
	if ff.GetSelectedIndex() != 1 {
		t.Errorf(
			"GetSelectedIndex() = %v, want 1",
			ff.GetSelectedIndex(),
		)
	}

	// Verify selected value
	if ff.GetSelectedValue() != "Option 2" {
		t.Errorf(
			"GetSelectedValue() = %v, want Option 2",
			ff.GetSelectedValue(),
		)
	}

	// Verify the SimpleField was added to the paragraph
	found := false
	for child := range para.Children() {
		if child.LocalName() == fldSimpleLocalName {
			found = true

			break
		}
	}
	if !found {
		t.Error(
			"SimpleField not found in paragraph",
		)
	}
}

func TestDocument_GetFormFields(t *testing.T) {
	// Create a new document
	doc, err := New(
		t.TempDir()+"/test.docx",
		DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Document element is nil")
	}

	body := docElem.GetOrCreateBody()

	// Add paragraphs with different form fields
	para1 := elements.NewParagraph()
	InsertTextField(para1, "Field1", "Value1")
	body.AppendChild(para1)

	para2 := elements.NewParagraph()
	InsertCheckBox(para2, "Field2", true)
	body.AppendChild(para2)

	para3 := elements.NewParagraph()
	InsertDropDown(
		para3,
		"Field3",
		[]string{"A", "B", "C"},
		0,
	)
	body.AppendChild(para3)

	// Get all form fields
	fields := doc.GetFormFields()

	// Verify count
	if len(fields) != 3 {
		t.Errorf(
			"GetFormFields() returned %d fields, want 3",
			len(fields),
		)
	}

	// Verify field types and names
	expectedFields := []struct {
		name      string
		fieldType FormFieldType
	}{
		{"Field1", FormFieldTypeText},
		{"Field2", FormFieldTypeCheckBox},
		{"Field3", FormFieldTypeDropDown},
	}

	for i, expected := range expectedFields {
		if i >= len(fields) {
			t.Errorf(
				"Missing field at index %d",
				i,
			)

			continue
		}

		if fields[i].Name() != expected.name {
			t.Errorf(
				"Field[%d].Name() = %v, want %v",
				i,
				fields[i].Name(),
				expected.name,
			)
		}

		if fields[i].Type() != expected.fieldType {
			t.Errorf(
				"Field[%d].Type() = %v, want %v",
				i,
				fields[i].Type(),
				expected.fieldType,
			)
		}
	}
}

func TestDocument_GetFormFieldByName(
	t *testing.T,
) {
	// Create a new document
	doc, err := New(
		t.TempDir()+"/test.docx",
		DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Document element is nil")
	}

	body := docElem.GetOrCreateBody()

	// Add paragraphs with form fields
	para1 := elements.NewParagraph()
	InsertTextField(para1, "MyTextField", "Test")
	body.AppendChild(para1)

	para2 := elements.NewParagraph()
	InsertCheckBox(para2, "MyCheckBox", false)
	body.AppendChild(para2)

	// Test finding existing field
	field := doc.GetFormFieldByName("MyTextField")
	if field == nil {
		t.Fatal(
			"GetFormFieldByName(\"MyTextField\") returned nil",
		)
	}

	if field.Type() != FormFieldTypeText {
		t.Errorf(
			"Field type = %v, want %v",
			field.Type(),
			FormFieldTypeText,
		)
	}

	// Test finding another field
	field2 := doc.GetFormFieldByName("MyCheckBox")
	if field2 == nil {
		t.Fatal(
			"GetFormFieldByName(\"MyCheckBox\") returned nil",
		)
	}

	if field2.Type() != FormFieldTypeCheckBox {
		t.Errorf(
			"Field type = %v, want %v",
			field2.Type(),
			FormFieldTypeCheckBox,
		)
	}

	// Test finding non-existent field
	field3 := doc.GetFormFieldByName(
		"NonExistent",
	)
	if field3 != nil {
		t.Error(
			"GetFormFieldByName(\"NonExistent\") should return nil",
		)
	}
}
