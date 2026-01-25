//
//nolint:revive // file-length-limit: form fields implementation requires comprehensive coverage
package wordprocessing

import (
	"errors"
	"fmt"
	"strings"

	"github.com/connerohnesorge/goffice/openxml/types"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const (
	// maxDropDownItems is the maximum number of items allowed in a dropdown per Word specification.
	maxDropDownItems = 25
	// intFormatSpecifier is the format specifier for integer conversion.
	intFormatSpecifier = "%d"
	// fldSimpleLocalName is the local name for SimpleField elements.
	fldSimpleLocalName = "fldSimple"
	// ffDataLocalName is the local name for FormFieldData elements.
	ffDataLocalName = "ffData"
)

// FormFieldType represents the type of form field.
type FormFieldType int

const (
	// FormFieldTypeText represents a text input form field (FORMTEXT).
	FormFieldTypeText FormFieldType = iota
	// FormFieldTypeCheckBox represents a checkbox form field (FORMCHECKBOX).
	FormFieldTypeCheckBox
	// FormFieldTypeDropDown represents a dropdown list form field (FORMDROPDOWN).
	FormFieldTypeDropDown
)

// String returns the string representation of the form field type.
func (t FormFieldType) String() string {
	switch t {
	case FormFieldTypeText:
		return "Text"
	case FormFieldTypeCheckBox:
		return "CheckBox"
	case FormFieldTypeDropDown:
		return "DropDown"
	default:
		return "Unknown"
	}
}

// FormField wraps a form field element in a Word document.
// It provides a high-level API for working with legacy form fields
// (Word 97-2003 compatible using fldSimple/ffData).
type FormField struct {
	simpleField *elements.SimpleField
	ffData      *elements.FormFieldData
	paragraph   *elements.Paragraph
	fieldType   FormFieldType
}

// newFormField creates a new FormField wrapper.
func newFormField(
	simpleField *elements.SimpleField,
	ffData *elements.FormFieldData,
	paragraph *elements.Paragraph,
	fieldType FormFieldType,
) *FormField {
	return &FormField{
		simpleField: simpleField,
		ffData:      ffData,
		paragraph:   paragraph,
		fieldType:   fieldType,
	}
}

// Name returns the name of the form field.
func (f *FormField) Name() string {
	if f.ffData == nil ||
		f.ffData.FormFieldName == nil {
		return ""
	}
	if val := f.ffData.FormFieldName.Val; val != nil {
		return val.Value()
	}

	return ""
}

// SetName sets the name of the form field.
func (f *FormField) SetName(name string) {
	if f.ffData == nil {
		return
	}
	if f.ffData.FormFieldName == nil {
		f.ffData.FormFieldName = elements.NewFormFieldName()
	}
	f.ffData.FormFieldName.Val = types.NewStringValue(
		name,
	)
}

// Type returns the type of the form field.
func (f *FormField) Type() FormFieldType {
	return f.fieldType
}

// Enabled returns whether the form field is enabled.
func (f *FormField) Enabled() bool {
	if f.ffData == nil {
		return true // Default is enabled
	}
	// Enabled element uses presence/absence to indicate state
	// If present, check its inner text value (true/false/1/0)
	if f.ffData.Enabled != nil {
		text := f.ffData.Enabled.InnerText()
		// Empty or "true" or "1" means enabled
		return text == "" || text == "true" ||
			text == "1"
	}

	return true
}

// SetEnabled sets whether the form field is enabled.
//
//nolint:revive // flag-parameter: standard setter for boolean value
func (f *FormField) SetEnabled(isEnabled bool) {
	if f.ffData == nil {
		return
	}
	if isEnabled {
		// Remove the Enabled element if present (default is enabled)
		if f.ffData.Enabled != nil {
			f.ffData.Enabled = nil
		}
	} else {
		// Add Enabled element with "false" value
		if f.ffData.Enabled == nil {
			f.ffData.Enabled = elements.NewEnabled()
		}
		f.ffData.Enabled.SetInnerText("false")
	}
}

// Remove removes the form field from its paragraph.
func (f *FormField) Remove() {
	if f.paragraph != nil &&
		f.simpleField != nil {
		f.paragraph.RemoveChild(f.simpleField)
	}
}

// TextField Methods

// GetTextValue returns the current value of a text field.
// Returns empty string if this is not a text field.
func (f *FormField) GetTextValue() string {
	if f.fieldType != FormFieldTypeText ||
		f.simpleField == nil {
		return ""
	}
	// The text value is stored in the run content within the simple field
	if run := f.simpleField.Run; run != nil {
		return run.InnerText()
	}

	return ""
}

// SetTextValue sets the value of a text field.
// Does nothing if this is not a text field.
func (f *FormField) SetTextValue(value string) {
	if f.fieldType != FormFieldTypeText ||
		f.simpleField == nil {
		return
	}
	// Set the text in the run content
	if f.simpleField.Run == nil {
		f.simpleField.Run = elements.NewRun("")
	}
	// Clear existing text elements and add new one
	for child := range f.simpleField.Run.Children() {
		if child.LocalName() == "t" {
			f.simpleField.Run.RemoveChild(child)
		}
	}
	f.simpleField.Run.AppendChild(
		elements.NewText(value),
	)
}

// GetMaxLength returns the maximum length for a text field.
// Returns 0 if no maximum or if this is not a text field.
func (f *FormField) GetMaxLength() int {
	if f.fieldType != FormFieldTypeText ||
		f.ffData == nil ||
		f.ffData.TextInput == nil {
		return 0
	}
	if maxLen := f.ffData.TextInput.MaxLength; maxLen != nil {
		if val := maxLen.Val; val != nil {
			// Parse string to int
			var length int
			_, _ = fmt.Sscanf(
				val.Value(),
				intFormatSpecifier,
				&length,
			)

			return length
		}
	}

	return 0
}

// SetMaxLength sets the maximum length for a text field.
// Does nothing if this is not a text field.
func (f *FormField) SetMaxLength(length int) {
	if f.fieldType != FormFieldTypeText ||
		f.ffData == nil {
		return
	}
	if f.ffData.TextInput == nil {
		f.ffData.TextInput = elements.NewTextInput()
	}
	if f.ffData.TextInput.MaxLength == nil {
		f.ffData.TextInput.MaxLength = elements.NewMaxLength()
	}
	f.ffData.TextInput.MaxLength.Val = types.NewStringValue(
		fmt.Sprintf(intFormatSpecifier, length),
	)
}

// GetDefaultTextValue returns the default value for a text field.
// Returns empty string if this is not a text field.
func (f *FormField) GetDefaultTextValue() string {
	if f.fieldType != FormFieldTypeText ||
		f.ffData == nil ||
		f.ffData.TextInput == nil {
		return ""
	}
	if defaultVal := f.ffData.TextInput.DefaultTextBoxFormFieldString; defaultVal != nil {
		return defaultVal.InnerText()
	}

	return ""
}

// SetDefaultTextValue sets the default value for a text field.
// Does nothing if this is not a text field.
func (f *FormField) SetDefaultTextValue(
	value string,
) {
	if f.fieldType != FormFieldTypeText ||
		f.ffData == nil {
		return
	}
	if f.ffData.TextInput == nil {
		f.ffData.TextInput = elements.NewTextInput()
	}
	if f.ffData.TextInput.DefaultTextBoxFormFieldString == nil {
		f.ffData.TextInput.DefaultTextBoxFormFieldString = elements.NewDefaultTextBoxFormFieldString()
	}
	f.ffData.TextInput.DefaultTextBoxFormFieldString.SetInnerText(
		value,
	)
}

// CheckBox Methods

// IsChecked returns whether the checkbox is checked.
// Returns false if this is not a checkbox field.
func (f *FormField) IsChecked() bool {
	if f.fieldType != FormFieldTypeCheckBox ||
		f.ffData == nil ||
		f.ffData.CheckBox == nil {
		return false
	}
	// Checked element presence indicates checked state
	return f.ffData.CheckBox.Checked != nil
}

// SetChecked sets whether the checkbox is checked.
// Does nothing if this is not a checkbox field.
//
//nolint:revive // flag-parameter: standard setter for boolean value
func (f *FormField) SetChecked(isChecked bool) {
	if f.fieldType != FormFieldTypeCheckBox ||
		f.ffData == nil {
		return
	}
	if f.ffData.CheckBox == nil {
		f.ffData.CheckBox = elements.NewCheckBox()
	}
	if isChecked {
		// Add Checked element if not present
		if f.ffData.CheckBox.Checked == nil {
			f.ffData.CheckBox.Checked = elements.NewChecked()
		}
	} else {
		// Remove Checked element
		f.ffData.CheckBox.Checked = nil
	}

	// Update the visual representation in the run
	// Checkboxes use Wingdings font with specific symbols:
	// F052 (checked), F06F (unchecked)
	if f.simpleField == nil ||
		f.simpleField.Run == nil {
		return
	}

	symbol := "\uF06F" // unchecked
	if isChecked {
		symbol = "\uF052" // checked
	}
	// Clear existing text and add checkbox symbol
	for child := range f.simpleField.Run.Children() {
		if child.LocalName() == "t" {
			f.simpleField.Run.RemoveChild(child)
		}
	}
	f.simpleField.Run.AppendChild(
		elements.NewText(symbol),
	)
}

// GetCheckBoxSize returns the size of the checkbox in half-points.
// Returns 0 if no size is set or if this is not a checkbox field.
func (f *FormField) GetCheckBoxSize() int {
	if f.fieldType != FormFieldTypeCheckBox ||
		f.ffData == nil ||
		f.ffData.CheckBox == nil {
		return 0
	}
	if size := f.ffData.CheckBox.FormFieldSize; size != nil {
		// Size is stored in inner text
		text := size.InnerText()
		if text != "" {
			var val int
			_, _ = fmt.Sscanf(
				text,
				intFormatSpecifier,
				&val,
			)

			return val
		}
	}

	return 0
}

// SetCheckBoxSize sets the size of the checkbox in half-points.
// Does nothing if this is not a checkbox field.
func (f *FormField) SetCheckBoxSize(size int) {
	if f.fieldType != FormFieldTypeCheckBox ||
		f.ffData == nil {
		return
	}
	if f.ffData.CheckBox == nil {
		f.ffData.CheckBox = elements.NewCheckBox()
	}
	if f.ffData.CheckBox.FormFieldSize == nil {
		f.ffData.CheckBox.FormFieldSize = elements.NewFormFieldSize()
	}
	f.ffData.CheckBox.FormFieldSize.SetInnerText(
		fmt.Sprintf(intFormatSpecifier, size),
	)
}

// DropDown Methods

// GetDropDownItems returns the list of items in a dropdown field.
// Returns empty slice if this is not a dropdown field.
func (f *FormField) GetDropDownItems() []string {
	if f.fieldType != FormFieldTypeDropDown ||
		f.ffData == nil ||
		f.ffData.DropDownListFormField == nil {
		return nil
	}

	var items []string
	ddList := f.ffData.DropDownListFormField

	// Iterate through child elements to find all listEntry elements
	for child := range ddList.Children() {
		listEntry, ok := child.(*elements.ListEntryFormField)
		if !ok {
			continue
		}

		text := listEntry.InnerText()
		if text != "" {
			items = append(items, text)
		}
	}

	return items
}

// SetDropDownItems sets the list of items in a dropdown field.
// Maximum 25 items are allowed per Word specification.
// Does nothing if this is not a dropdown field.
func (f *FormField) SetDropDownItems(
	items []string,
) {
	if f.fieldType != FormFieldTypeDropDown ||
		f.ffData == nil {
		return
	}
	if f.ffData.DropDownListFormField == nil {
		f.ffData.DropDownListFormField = elements.NewDropDownListFormField()
	}

	// Enforce 25-item limit per Word specification
	itemsToAdd := items
	if len(itemsToAdd) > maxDropDownItems {
		itemsToAdd = itemsToAdd[:maxDropDownItems]
	}

	ddList := f.ffData.DropDownListFormField

	// Remove existing list entries
	for child := range ddList.Children() {
		if _, ok := child.(*elements.ListEntryFormField); ok {
			ddList.RemoveChild(child)
		}
	}

	// Add new list entries
	for _, item := range itemsToAdd {
		listEntry := elements.NewListEntryFormField()
		listEntry.SetInnerText(item)
		ddList.AppendChild(listEntry)
	}
}

// GetSelectedIndex returns the index of the selected item in a dropdown field.
// Returns -1 if no item is selected or if this is not a dropdown field.
func (f *FormField) GetSelectedIndex() int {
	if f.fieldType != FormFieldTypeDropDown ||
		f.ffData == nil ||
		f.ffData.DropDownListFormField == nil {
		return -1
	}
	if result := f.ffData.DropDownListFormField.DropDownListSelection; result != nil {
		// Selection index is stored in inner text
		text := result.InnerText()
		if text != "" {
			var val int
			_, _ = fmt.Sscanf(
				text,
				intFormatSpecifier,
				&val,
			)

			return val
		}
	}

	return -1
}

// SetSelectedIndex sets the index of the selected item in a dropdown field.
// Does nothing if this is not a dropdown field.
func (f *FormField) SetSelectedIndex(index int) {
	if f.fieldType != FormFieldTypeDropDown ||
		f.ffData == nil {
		return
	}
	if f.ffData.DropDownListFormField == nil {
		f.ffData.DropDownListFormField = elements.NewDropDownListFormField()
	}
	if f.ffData.DropDownListFormField.DropDownListSelection == nil {
		f.ffData.DropDownListFormField.DropDownListSelection = elements.NewDropDownListSelection()
	}
	f.ffData.DropDownListFormField.DropDownListSelection.SetInnerText(
		fmt.Sprintf(intFormatSpecifier, index),
	)

	// Update the visual representation in the run
	if f.simpleField == nil ||
		f.simpleField.Run == nil {
		return
	}

	items := f.GetDropDownItems()
	if index < 0 || index >= len(items) {
		return
	}

	// Clear existing text and add selected item
	for child := range f.simpleField.Run.Children() {
		if child.LocalName() == "t" {
			f.simpleField.Run.RemoveChild(child)
		}
	}
	f.simpleField.Run.AppendChild(
		elements.NewText(items[index]),
	)
}

// GetSelectedValue returns the value of the selected item in a dropdown field.
// Returns empty string if no item is selected or if this is not a dropdown field.
func (f *FormField) GetSelectedValue() string {
	items := f.GetDropDownItems()
	index := f.GetSelectedIndex()
	if index >= 0 && index < len(items) {
		return items[index]
	}

	return ""
}

// detectFormFieldType detects the type of form field from the instruction.
func detectFormFieldType(
	instruction string,
) (FormFieldType, error) {
	if instruction == "" {
		return FormFieldTypeText, errors.New(
			"empty instruction",
		)
	}

	// Trim spaces and convert to uppercase for comparison
	normalized := strings.ToUpper(
		strings.TrimSpace(instruction),
	)

	// Simple detection based on instruction content
	if strings.HasPrefix(normalized, "FORMTEXT") {
		return FormFieldTypeText, nil
	}
	if strings.HasPrefix(
		normalized,
		"FORMCHECKBOX",
	) {
		return FormFieldTypeCheckBox, nil
	}
	if strings.HasPrefix(
		normalized,
		"FORMDROPDOWN",
	) {
		return FormFieldTypeDropDown, nil
	}

	return FormFieldTypeText, fmt.Errorf(
		"unknown form field instruction: %s",
		instruction,
	)
}

// InsertTextField inserts a text form field into the paragraph.
// Creates a SimpleField with FORMTEXT instruction, FormFieldData, and default value.
func InsertTextField(
	p *elements.Paragraph,
	name, defaultValue string,
) *FormField {
	// Create SimpleField
	simpleField := elements.NewSimpleField()
	simpleField.Instruction = types.NewStringValue(
		" FORMTEXT ",
	)

	// Create FormFieldData
	ffData := elements.NewFormFieldData()
	ffData.FormFieldName = elements.NewFormFieldName()
	ffData.FormFieldName.Val = types.NewStringValue(
		name,
	)
	ffData.Enabled = elements.NewEnabled()

	// Create TextInput with default value
	ffData.TextInput = elements.NewTextInput()
	if defaultValue != "" {
		ffData.TextInput.DefaultTextBoxFormFieldString = elements.NewDefaultTextBoxFormFieldString()
		ffData.TextInput.DefaultTextBoxFormFieldString.SetInnerText(
			defaultValue,
		)
	}

	// Append FormFieldData to SimpleField
	simpleField.AppendChild(ffData)

	// Create Run with text content
	run := elements.NewRun("")
	run.AppendChild(
		elements.NewText(defaultValue),
	)
	simpleField.AppendChild(run)

	// Append SimpleField to paragraph
	p.AppendChild(simpleField)

	return newFormField(
		simpleField,
		ffData,
		p,
		FormFieldTypeText,
	)
}

// InsertCheckBox inserts a checkbox form field into the paragraph.
// Creates a SimpleField with FORMCHECKBOX instruction, FormFieldData, and checkbox symbol.
//
//nolint:revive // flag-parameter: standard API for checkbox state
func InsertCheckBox(
	p *elements.Paragraph,
	name string,
	checked bool,
) *FormField {
	// Create SimpleField
	simpleField := elements.NewSimpleField()
	simpleField.Instruction = types.NewStringValue(
		" FORMCHECKBOX ",
	)

	// Create FormFieldData
	ffData := elements.NewFormFieldData()
	ffData.FormFieldName = elements.NewFormFieldName()
	ffData.FormFieldName.Val = types.NewStringValue(
		name,
	)
	ffData.Enabled = elements.NewEnabled()

	// Create CheckBox
	ffData.CheckBox = elements.NewCheckBox()
	ffData.CheckBox.FormFieldSize = elements.NewFormFieldSize()
	ffData.CheckBox.FormFieldSize.SetInnerText(
		"20",
	)
	ffData.CheckBox.DefaultCheckBoxFormFieldState = elements.NewDefaultCheckBoxFormFieldState()
	ffData.CheckBox.DefaultCheckBoxFormFieldState.SetInnerText(
		"0",
	)

	if checked {
		ffData.CheckBox.Checked = elements.NewChecked()
	}

	// Append FormFieldData to SimpleField
	simpleField.AppendChild(ffData)

	// Create Run with checkbox symbol (Wingdings font)
	run := elements.NewRun("")
	sym := elements.NewSymbolChar()
	sym.Font = types.NewStringValue("Wingdings")
	if checked {
		sym.Char = types.NewStringValue(
			"F052",
		) // checked
	} else {
		sym.Char = types.NewStringValue("F06F") // unchecked
	}
	run.AppendChild(sym)
	simpleField.AppendChild(run)

	// Append SimpleField to paragraph
	p.AppendChild(simpleField)

	return newFormField(
		simpleField,
		ffData,
		p,
		FormFieldTypeCheckBox,
	)
}

// InsertDropDown inserts a dropdown form field into the paragraph.
// Creates a SimpleField with FORMDROPDOWN instruction, FormFieldData, and list items.
// The selectedIndex parameter specifies which item is initially selected (0-based).
func InsertDropDown(
	p *elements.Paragraph,
	name string,
	items []string,
	selectedIndex int,
) *FormField {
	// Enforce 25-item limit per Word specification
	itemsToAdd := items
	if len(itemsToAdd) > maxDropDownItems {
		itemsToAdd = itemsToAdd[:maxDropDownItems]
	}

	// Create SimpleField
	simpleField := elements.NewSimpleField()
	simpleField.Instruction = types.NewStringValue(
		" FORMDROPDOWN ",
	)

	// Create FormFieldData
	ffData := elements.NewFormFieldData()
	ffData.FormFieldName = elements.NewFormFieldName()
	ffData.FormFieldName.Val = types.NewStringValue(
		name,
	)
	ffData.Enabled = elements.NewEnabled()

	// Create DropDownListFormField
	ffData.DropDownListFormField = elements.NewDropDownListFormField()

	// Set selected index
	if selectedIndex >= 0 &&
		selectedIndex < len(itemsToAdd) {
		ffData.DropDownListFormField.DropDownListSelection = elements.NewDropDownListSelection()
		ffData.DropDownListFormField.DropDownListSelection.SetInnerText(
			fmt.Sprintf(
				intFormatSpecifier,
				selectedIndex,
			),
		)
	}

	// Add list items
	for _, item := range itemsToAdd {
		listEntry := elements.NewListEntryFormField()
		listEntry.SetInnerText(item)
		ffData.DropDownListFormField.AppendChild(
			listEntry,
		)
	}

	// Append FormFieldData to SimpleField
	simpleField.AppendChild(ffData)

	// Create Run with selected item text
	run := elements.NewRun("")
	if selectedIndex >= 0 &&
		selectedIndex < len(itemsToAdd) {
		run.AppendChild(
			elements.NewText(
				itemsToAdd[selectedIndex],
			),
		)
	} else if len(itemsToAdd) > 0 {
		run.AppendChild(elements.NewText(itemsToAdd[0]))
	}
	simpleField.AppendChild(run)

	// Append SimpleField to paragraph
	p.AppendChild(simpleField)

	return newFormField(
		simpleField,
		ffData,
		p,
		FormFieldTypeDropDown,
	)
}

// GetFormFields returns all form fields in the document.
// It traverses the document body and collects all SimpleField elements that contain FormFieldData.
//
//nolint:revive // cognitive-complexity: complex traversal logic necessary for form field detection
func (d *Document) GetFormFields() []*FormField {
	var formFields []*FormField

	mainPart := d.MainPart()
	if mainPart == nil {
		return formFields
	}

	doc := mainPart.Document()
	if doc == nil {
		return formFields
	}

	body := doc.Body()
	if body == nil {
		return formFields
	}

	// Traverse all paragraphs in the body
	for para := range body.Paragraphs() {
		// Look for SimpleField elements in each paragraph
		for child := range para.Children() {
			if child.LocalName() != fldSimpleLocalName ||
				child.NamespaceURI() != "http://schemas.openxmlformats.org/wordprocessingml/2006/main" {
				continue
			}

			// Try to cast to SimpleField
			var simpleField *elements.SimpleField
			switch v := child.(type) {
			case *elements.SimpleField:
				simpleField = v
			default:
				// Skip if not a SimpleField
				continue
			}

			// Check if it has an instruction attribute
			if simpleField.Instruction == nil {
				continue
			}

			instruction := simpleField.Instruction.Value()

			// Detect form field type
			fieldType, err := detectFormFieldType(
				instruction,
			)
			if err != nil {
				continue
			}

			// Find FormFieldData child
			var ffData *elements.FormFieldData
			for ffChild := range simpleField.Children() {
				if ffChild.LocalName() != ffDataLocalName {
					continue
				}

				if ffd, ok := ffChild.(*elements.FormFieldData); ok {
					ffData = ffd

					break
				}
			}

			if ffData != nil {
				formFields = append(
					formFields,
					newFormField(
						simpleField,
						ffData,
						para,
						fieldType,
					),
				)
			}
		}
	}

	return formFields
}

// GetFormFieldByName finds and returns a form field by its name.
// Returns nil if no form field with the given name is found.
func (d *Document) GetFormFieldByName(
	name string,
) *FormField {
	formFields := d.GetFormFields()

	for _, ff := range formFields {
		if ff.Name() == name {
			return ff
		}
	}

	return nil
}
