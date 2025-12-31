# Proposal Brief: Word Form Fields Implementation

## Priority: TIER 2 - High
**Impact**: 7/10  
**Effort**: 3-4 weeks  
**Status**: Basic elements exist, NO manipulation API

## Gap Analysis

### Current State
**What exists:**
- ✅ **Basic field elements** (wordprocessing/elements/elements.go):
  - `FieldCode` - Contains field instructions
  - `FieldChar` - Field begin/end/separate markers
  - `SimpleField` - Simple field wrapper
  - `FormFieldData` - Form field properties
  - `DropDownListFormField` - Dropdown field type
  - `TextBoxFormFieldType` - Text input field type

**What's missing:**
- ❌ No `InsertTextField(name, value)` method
- ❌ No `InsertCheckBox(name, checked)` method
- ❌ No `InsertDropDown(name, items, selected)` method
- ❌ No form field value setting/getting
- ❌ No form field protection (lock/unlock)
- ❌ No form field validation
- ❌ No high-level FormField wrapper class

### OpenXML-SDK Comparison
**Microsoft Word Object Model:**
```csharp
// Insert form fields
document.FormFields.Add(range, WdFieldType.wdFieldFormTextInput)
formField.Name = "CustomerName"
formField.Result = "John Doe"

// Checkboxes
formField.CheckBox.Value = true
formField.CheckBox.Size = 11

// Dropdowns
formField.DropDown.ListEntries.Add("Option 1")
formField.DropDown.Value = 2
```

## Implementation Scope

### Files to Create

1. **wordprocessing/form_fields.go** (NEW)
   - `FormField` wrapper class
   - `FormFieldType` enum (TextField, CheckBox, DropDown)
   - High-level methods for field manipulation

2. **wordprocessing/document.go** (MODIFY)
   - Add `GetFormFields() iter.Seq[*FormField]`
   - Add `GetFormFieldByName(name string) *FormField`

3. **wordprocessing/paragraph.go** (MODIFY)
   - Add `InsertTextField(name, defaultValue string) *FormField`
   - Add `InsertCheckBox(name string, checked bool) *FormField`
   - Add `InsertDropDown(name string, items []string, selectedIndex int) *FormField`

### Core Design Decisions

**1. Form Field Structure in OOXML**
Legacy form fields (Word 97-2003 compatibility) use `fldSimple` or `fldChar` + `ffData`:
```xml
<w:p>
  <w:fldSimple w:instr="FORMTEXT">
    <w:ffData>
      <w:name w:val="CustomerName"/>
      <w:textInput>
        <w:default w:val="Enter name"/>
      </w:textInput>
    </w:ffData>
    <w:r><w:t>John Doe</w:t></w:r>
  </w:fldSimple>
</w:p>
```

Modern form fields (Word 2007+) use content controls (SDT elements) - that's a separate feature (Content Controls).

**2. Focus on Legacy Form Fields**
Phase 1: Legacy form fields (backward compatible, simpler)
Phase 2: Modern content controls (separate proposal)

**3. FormField Wrapper Design**
```go
type FormField struct {
    simpleField *elements.SimpleField
    ffData      *elements.FormFieldData
    paragraph   *Paragraph
    
    name        string
    typ         FormFieldType
}

func (f *FormField) Name() string { ... }
func (f *FormField) SetName(name string) { ... }
func (f *FormField) Type() FormFieldType { ... }

// TextField specific
func (f *FormField) GetTextValue() string { ... }
func (f *FormField) SetTextValue(value string) { ... }
func (f *FormField) SetMaxLength(length int) { ... }

// CheckBox specific  
func (f *FormField) IsChecked() bool { ... }
func (f *FormField) SetChecked(checked bool) { ... }
func (f *FormField) SetCheckBoxSize(size int) { ... }

// DropDown specific
func (f *FormField) GetDropDownItems() []string { ... }
func (f *FormField) SetDropDownItems(items []string) { ... }
func (f *FormField) GetSelectedIndex() int { ... }
func (f *FormField) SetSelectedIndex(index int) { ... }
```

**4. Field Code Construction**
Text field: `FORMTEXT`
Checkbox: `FORMCHECKBOX`
Dropdown: `FORMDROPDOWN`

**5. Protection/Editing**
Form fields can be protected (read-only except when filling form):
```go
func (d *Document) ProtectForForms() { ... }
func (d *Document) UnprotectDocument() { ... }
```

## References

### Existing Code
- `wordprocessing/elements/elements.go` - Field elements (FieldCode, FieldChar, SimpleField, FormFieldData)
- `wordprocessing/paragraph.go` - Paragraph API (add field insertion here)
- `wordprocessing/run.go` - Run API (fields use runs)

### OpenXML-SDK References
Look at @OpenXML-SDK/ for:
- `FormFieldData` structure
- `TextInput`, `CheckBox`, `DropDownList` field types
- Field code formatting

### Test Data
- Create Word documents with form fields (Word 97-2003 format for legacy fields)
- Test roundtrip: create form → save → open in Word → verify works

## Success Criteria

- [ ] Can call `paragraph.InsertTextField("name", "default")` and field appears
- [ ] Can call `field.SetTextValue("new value")` and it persists
- [ ] Can insert checkbox and toggle checked state
- [ ] Can insert dropdown with multiple items
- [ ] Can set selected dropdown item by index
- [ ] Can iterate all form fields in document
- [ ] Can get form field by name
- [ ] Form protection works (document.ProtectForForms())
- [ ] Fields work correctly when opened in Microsoft Word
- [ ] Roundtrip test passes

## Spec Capabilities to Update

- **wordprocessing-document**: Add form field requirements
- **wordprocessing-elements**: Note form field API additions

## Dependencies

- Existing field elements (FieldCode, FieldChar, SimpleField)
- Existing paragraph/run infrastructure

## Out of Scope

- Modern content controls (SDT-based forms) - separate proposal
- Form field macros/calculations
- Form field help text/status bar text (can add in Phase 2)
- Advanced validation rules (can add in Phase 2)

## Example Usage (Target API)

```go
doc, _ := wordprocessing.New("form.docx", wordprocessing.DocTypeDocument)

para := doc.MainDocumentPart().Document().Body().AppendParagraph()

// Insert text field
nameField := para.InsertTextField("CustomerName", "Enter your name")
nameField.SetMaxLength(50)

// Insert checkbox
agreeField := para.InsertCheckBox("AgreeTerms", false)
agreeField.SetCheckBoxSize(12)

// Insert dropdown
countryField := para.InsertDropDown("Country", []string{"USA", "Canada", "UK"}, 0)

// Protect document for form filling only
doc.ProtectForForms()

doc.Save()

// Later: fill the form
doc2, _ := wordprocessing.Open("form.docx")
for field := range doc2.GetFormFields() {
    if field.Name() == "CustomerName" {
        field.SetTextValue("John Doe")
    }
    if field.Name() == "AgreeTerms" {
        field.SetChecked(true)
    }
}
doc2.Save()
```

## Technical Details

### FormFieldData XML Structure
```xml
<w:ffData>
  <w:name w:val="FieldName"/>
  <w:enabled/>  <!-- Can be disabled -->
  <w:calcOnExit w:val="0"/>  <!-- Recalc on exit -->
  <w:helpText w:type="text" w:val="Help text"/>
  <w:statusText w:type="text" w:val="Status text"/>
  
  <!-- For text fields -->
  <w:textInput>
    <w:type w:val="regular"/>  <!-- regular, number, date, currentDate, currentTime -->
    <w:default w:val="Default value"/>
    <w:maxLength w:val="255"/>
    <w:format w:val=""/>  <!-- Format string -->
  </w:textInput>
  
  <!-- For checkboxes -->
  <w:checkBox>
    <w:size w:val="10"/>  <!-- Half-points -->
    <w:sizeAuto/>  <!-- Auto size -->
    <w:default w:val="0"/>  <!-- 0=unchecked, 1=checked -->
    <w:checked w:val="1"/>
  </w:checkBox>
  
  <!-- For dropdowns -->
  <w:ddList>
    <w:result w:val="0"/>  <!-- Selected index -->
    <w:listEntry w:val="Item 1"/>
    <w:listEntry w:val="Item 2"/>
    <w:listEntry w:val="Item 3"/>
  </w:ddList>
</w:ffData>
```

### Field Result Storage
The visible text is stored in run elements after the field code:
```xml
<w:fldSimple w:instr="FORMTEXT">
  <w:ffData>...</w:ffData>
  <w:r>
    <w:t>Current field value displayed here</w:t>
  </w:r>
</w:fldSimple>
```

So setting field value requires:
1. Update ffData default/checked/result
2. Update the run text to match
