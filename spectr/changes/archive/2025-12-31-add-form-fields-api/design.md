# Design: Word Form Fields API

## Problem Statement

Users need to create fillable forms in Word documents with text fields, checkboxes, and dropdown lists. While goffice has the underlying XML elements (`SimpleField`, `FormFieldData`, etc.), there's no high-level API to insert, configure, or fill form fields.

This design addresses legacy form fields (Word 97-2003 compatible) using `fldSimple` elements with `ffData` child elements. Modern content controls (Word 2007+ SDT elements) are intentionally out of scope for this phase.

## Current Architecture

```
Paragraph
    ↓
Run (contains text)
    ↓
Text element

[No form field API exists]
```

**Gap**: Cannot insert form fields or manipulate their values.

## Proposed Architecture

```
Document
    ↓ GetFormFields() iter.Seq[*FormField]
    ↓ GetFormFieldByName(name) *FormField
    ↓
Paragraph
    ↓ InsertTextField(name, default) *FormField
    ↓ InsertCheckBox(name, checked) *FormField
    ↓ InsertDropDown(name, items, selected) *FormField
    ↓
FormField (wrapper)
    ↓ wraps SimpleField + FormFieldData
    ↓ SetTextValue() / GetTextValue()
    ↓ SetChecked() / IsChecked()
    ↓ SetDropDownItems() / GetSelectedIndex()
    ↓
SimpleField element
    ↓ contains field instruction
    ↓
FormFieldData element
    ↓ contains field configuration
    ↓
Run element
    ↓ contains field result (displayed value)
```

## Design Alternatives Considered

### Alternative 1: Direct Element Manipulation
**Approach**: Users directly create SimpleField, FormFieldData, Run elements

**Pros**:
- No new abstraction layer
- Maximum control

**Cons**:
- Complex for users (must understand XML structure)
- Easy to create invalid combinations
- Requires manual synchronization between ffData and run text
- Violates "simple things should be simple" principle

**Decision**: REJECTED - Too low-level for common use case

### Alternative 2: Separate Insert Methods per Type
**Approach**: InsertTextFormField(), InsertCheckBoxFormField(), InsertDropDownFormField()

**Pros**:
- Very explicit
- Type-safe at insertion

**Cons**:
- Verbose method names
- Inconsistent with existing API patterns
- Harder to discover (3 methods instead of 3 well-named ones)

**Decision**: REJECTED - Prefer shorter, clearer names

### Alternative 3: FormField with Type-Specific Methods (CHOSEN)
**Approach**: FormField wrapper with IsTextField(), AsTextField(), etc.

**Pros**:
- Clean insertion API
- Type-safe accessors for field-specific properties
- Single FormField type simplifies iteration
- Matches Microsoft Word Object Model pattern

**Cons**:
- Runtime type checking needed

**Decision**: ACCEPTED - Best balance of usability and safety

### Alternative 4: Separate TextField, CheckBoxField, DropDownField Types
**Approach**: Three distinct Go types

**Pros**:
- Maximum type safety
- No runtime checks

**Cons**:
- Cannot iterate "all form fields" easily (need union type or separate iterators)
- More complex API surface
- Doesn't match Word's unified FormFields collection

**Decision**: REJECTED - Type safety doesn't outweigh usability loss

## FormField Wrapper Design

### Core Structure

```go
// FormFieldType identifies the type of form field
type FormFieldType int

const (
    FormFieldTypeText FormFieldType = iota
    FormFieldTypeCheckBox
    FormFieldTypeDropDown
)

// FormField wraps a legacy form field (fldSimple with ffData)
type FormField struct {
    simpleField *elements.SimpleField  // The w:fldSimple element
    ffData      *elements.FormFieldData // The w:ffData child
    paragraph   *Paragraph              // Parent paragraph (for context)

    fieldType   FormFieldType           // Cached type
    name        string                  // Cached name
}
```

### Type Detection

Field type is determined by the field instruction in `w:fldSimple/@w:instr`:
- `FORMTEXT` → TextField
- `FORMCHECKBOX` → CheckBox
- `FORMDROPDOWN` → DropDown

### Universal Methods

```go
// Name returns the form field name
func (f *FormField) Name() string

// SetName sets the form field name
func (f *FormField) SetName(name string)

// Type returns the form field type
func (f *FormField) Type() FormFieldType

// Enabled returns whether the field is enabled for editing
func (f *FormField) Enabled() bool

// SetEnabled sets whether the field is enabled
func (f *FormField) SetEnabled(enabled bool)

// Remove removes the form field from the document
func (f *FormField) Remove()
```

### Text Field Methods

```go
// GetTextValue returns the current text value (only for text fields)
func (f *FormField) GetTextValue() string

// SetTextValue sets the text value (only for text fields)
func (f *FormField) SetTextValue(value string)

// GetMaxLength returns the maximum text length (0 = unlimited)
func (f *FormField) GetMaxLength() int

// SetMaxLength sets the maximum text length
func (f *FormField) SetMaxLength(length int)

// GetTextFormat returns the text format (e.g., "Uppercase", "Lowercase")
func (f *FormField) GetTextFormat() string

// SetTextFormat sets the text format
func (f *FormField) SetTextFormat(format string)

// GetDefaultTextValue returns the default text value
func (f *FormField) GetDefaultTextValue() string

// SetDefaultTextValue sets the default text value
func (f *FormField) SetDefaultTextValue(value string)
```

### CheckBox Methods

```go
// IsChecked returns whether the checkbox is checked (only for checkboxes)
func (f *FormField) IsChecked() bool

// SetChecked sets the checkbox state (only for checkboxes)
func (f *FormField) SetChecked(checked bool)

// GetCheckBoxSize returns the checkbox size in half-points
func (f *FormField) GetCheckBoxSize() int

// SetCheckBoxSize sets the checkbox size (in half-points, e.g., 20 = 10pt)
func (f *FormField) SetCheckBoxSize(size int)

// GetAutoCheckBoxSize returns whether size is automatic
func (f *FormField) GetAutoCheckBoxSize() bool

// SetAutoCheckBoxSize sets whether size is automatic
func (f *FormField) SetAutoCheckBoxSize(auto bool)

// GetDefaultChecked returns the default checked state
func (f *FormField) GetDefaultChecked() bool

// SetDefaultChecked sets the default checked state
func (f *FormField) SetDefaultChecked(checked bool)
```

### DropDown Methods

```go
// GetDropDownItems returns the list of dropdown items (only for dropdowns)
func (f *FormField) GetDropDownItems() []string

// SetDropDownItems sets the list of dropdown items
// CONSTRAINT: OpenXML schema limits dropdown fields to maximum 25 items
// This method will panic if len(items) > 25
func (f *FormField) SetDropDownItems(items []string)

// GetSelectedIndex returns the selected item index (0-based)
func (f *FormField) GetSelectedIndex() int

// SetSelectedIndex sets the selected item index
func (f *FormField) SetSelectedIndex(index int)

// GetSelectedValue returns the selected item value
func (f *FormField) GetSelectedValue() string

// GetDefaultDropDownIndex returns the default selected item index
func (f *FormField) GetDefaultDropDownIndex() int

// SetDefaultDropDownIndex sets the default selected item index
func (f *FormField) SetDefaultDropDownIndex(index int)
```

## XML Structure Details

### Text Field XML

```xml
<w:p>
  <w:fldSimple w:instr="FORMTEXT">
    <w:ffData>
      <w:name w:val="CustomerName"/>
      <w:enabled/>
      <w:calcOnExit w:val="0"/>
      <w:textInput>
        <w:type w:val="regular"/>
        <w:default w:val="Enter name here"/>
        <w:maxLength w:val="50"/>
        <w:format w:val=""/>
      </w:textInput>
    </w:ffData>
    <w:r>
      <w:t>John Doe</w:t>
    </w:r>
  </w:fldSimple>
</w:p>
```

**Key elements**:
- `w:fldSimple/@w:instr="FORMTEXT"` - Field code
- `w:ffData/w:name/@w:val` - Field name
- `w:ffData/w:textInput/w:default/@w:val` - Default value
- `w:ffData/w:textInput/w:maxLength/@w:val` - Maximum length
- `w:r/w:t` - Current displayed value (must sync with ffData)

### CheckBox Field XML

```xml
<w:p>
  <w:fldSimple w:instr="FORMCHECKBOX">
    <w:ffData>
      <w:name w:val="AgreeTerms"/>
      <w:enabled/>
      <w:calcOnExit w:val="0"/>
      <w:checkBox>
        <w:size w:val="20"/>
        <w:sizeAuto w:val="0"/>
        <w:default w:val="0"/>
        <w:checked w:val="1"/>
      </w:checkBox>
    </w:ffData>
    <w:r>
      <w:sym w:font="MS Gothic" w:char="F052"/>
    </w:r>
  </w:fldSimple>
</w:p>
```

**Key elements**:
- `w:fldSimple/@w:instr="FORMCHECKBOX"` - Field code
- `w:ffData/w:checkBox/w:size/@w:val` - Size in half-points
- `w:ffData/w:checkBox/w:sizeAuto/@w:val` - Auto size (0=false, 1=true)
- `w:ffData/w:checkBox/w:default/@w:val` - Default state (0=unchecked, 1=checked)
- `w:ffData/w:checkBox/w:checked/@w:val` - Current state
- `w:r/w:sym` - Symbol character (checkbox visual)

**Checkbox symbols** (using Symbol font or MS Gothic):
- Unchecked: `&#xF06F;` (empty box) or `&#x2610;` (ballot box)
- Checked: `&#xF052;` (checked box) or `&#x2612;` (ballot box with X)

### DropDown Field XML

```xml
<w:p>
  <w:fldSimple w:instr="FORMDROPDOWN">
    <w:ffData>
      <w:name w:val="Country"/>
      <w:enabled/>
      <w:calcOnExit w:val="0"/>
      <w:ddList>
        <w:default w:val="0"/>
        <w:result w:val="1"/>
        <w:listEntry w:val="USA"/>
        <w:listEntry w:val="Canada"/>
        <w:listEntry w:val="UK"/>
        <w:listEntry w:val="Australia"/>
      </w:ddList>
    </w:ffData>
    <w:r>
      <w:t>Canada</w:t>
    </w:r>
  </w:fldSimple>
</w:p>
```

**Key elements**:
- `w:fldSimple/@w:instr="FORMDROPDOWN"` - Field code
- `w:ffData/w:ddList/w:default/@w:val` - Default selected index (0-based, used when form is reset)
- `w:ffData/w:ddList/w:result/@w:val` - Current selected index (0-based, differs from default when user changes selection)
- `w:ffData/w:ddList/w:listEntry/@w:val` - Each dropdown option
- `w:r/w:t` - Current selected value text (must match items[result])

## Field Code Construction

Field codes are simple string constants:

```go
const (
    fieldCodeText     = "FORMTEXT"
    fieldCodeCheckBox = "FORMCHECKBOX"
    fieldCodeDropDown = "FORMDROPDOWN"
)
```

No complex field code parsing needed - legacy form fields use fixed instruction strings.

## Value Synchronization Strategy

**Critical requirement**: Setting a form field value requires updating **TWO** places:

1. **ffData element** - Stores the structural value (checked state, selected index, default text)
2. **Run element** - Stores the displayed text (what user sees)

### Text Field Synchronization

```go
func (f *FormField) SetTextValue(value string) {
    // 1. Update ffData default value
    textInput := f.ffData.TextInput()
    if textInput == nil {
        textInput = elements.NewTextInput()
        f.ffData.AppendChild(textInput)
    }
    textInput.SetDefault(types.NewStringValue(value))

    // 2. Update or create run with text
    runs := f.simpleField.Elements[elements.Run]()
    var run *elements.Run
    if len(runs) > 0 {
        run = runs[0]
    } else {
        run = elements.NewRun()
        f.simpleField.AppendChild(run)
    }

    // Update text element
    texts := run.Elements[elements.Text]()
    var text *elements.Text
    if len(texts) > 0 {
        text = texts[0]
    } else {
        text = elements.NewText()
        run.AppendChild(text)
    }
    text.SetText(value)
}
```

### CheckBox Synchronization

```go
func (f *FormField) SetChecked(checked bool) {
    // 1. Update ffData checked state
    checkBox := f.ffData.CheckBox()
    if checkBox == nil {
        checkBox = elements.NewCheckBox()
        f.ffData.AppendChild(checkBox)
    }

    checkedVal := 0
    if checked {
        checkedVal = 1
    }
    checkBox.SetChecked(types.NewInt32Value(checkedVal))

    // 2. Update symbol character in run
    runs := f.simpleField.Elements[elements.Run]()
    var run *elements.Run
    if len(runs) > 0 {
        run = runs[0]
    } else {
        run = elements.NewRun()
        f.simpleField.AppendChild(run)
    }

    // Update or create symbol
    syms := run.Elements[elements.SymbolChar]()
    var sym *elements.SymbolChar
    if len(syms) > 0 {
        sym = syms[0]
    } else {
        sym = elements.NewSymbolChar()
        run.AppendChild(sym)
    }

    // Set appropriate checkbox symbol
    // IMPORTANT: Symbols MUST only be set via SetChecked() to ensure proper synchronization
    // Symbol characters (F052 for checked, F06F for unchecked) must match the checked state
    // Manual manipulation of symbol elements outside SetChecked() may create invalid state
    sym.SetFont(types.NewStringValue("MS Gothic"))
    if checked {
        sym.SetChar(types.NewStringValue("F052")) // Checked box
    } else {
        sym.SetChar(types.NewStringValue("F06F")) // Empty box
    }
}
```

### DropDown Synchronization

```go
func (f *FormField) SetSelectedIndex(index int) {
    // 1. Update ffData result
    ddList := f.ffData.DropDownList()
    if ddList == nil {
        return // Cannot set if no items exist
    }

    ddList.SetResult(types.NewInt32Value(index))

    // 2. Update run text to show selected item
    items := f.GetDropDownItems()
    if index < 0 || index >= len(items) {
        return // Invalid index
    }

    selectedText := items[index]

    runs := f.simpleField.Elements[elements.Run]()
    var run *elements.Run
    if len(runs) > 0 {
        run = runs[0]
    } else {
        run = elements.NewRun()
        f.simpleField.AppendChild(run)
    }

    texts := run.Elements[elements.Text]()
    var text *elements.Text
    if len(texts) > 0 {
        text = texts[0]
    } else {
        text = elements.NewText()
        run.AppendChild(text)
    }
    text.SetText(selectedText)
}
```

## Insertion Implementation

### Paragraph.InsertTextField

```go
func (p *Paragraph) InsertTextField(name, defaultValue string) *FormField {
    // Create SimpleField element
    simpleField := elements.NewSimpleField()
    simpleField.SetInstr(types.NewStringValue(fieldCodeText))

    // Create FormFieldData
    ffData := elements.NewFormFieldData()

    // Set name
    nameElem := elements.NewFormFieldName()
    nameElem.SetVal(types.NewStringValue(name))
    ffData.AppendChild(nameElem)

    // Set enabled
    enabled := elements.NewEnabled()
    ffData.AppendChild(enabled)

    // Create TextInput
    textInput := elements.NewTextInput()
    textInput.SetType(types.NewStringValue("regular"))
    textInput.SetDefault(types.NewStringValue(defaultValue))
    ffData.AppendChild(textInput)

    // Add ffData to simpleField
    simpleField.AppendChild(ffData)

    // Create run with default text
    run := elements.NewRun()
    text := elements.NewText()
    text.SetText(defaultValue)
    run.AppendChild(text)
    simpleField.AppendChild(run)

    // Add simpleField to paragraph
    p.element.AppendChild(simpleField)

    // Return FormField wrapper
    return &FormField{
        simpleField: simpleField,
        ffData:      ffData,
        paragraph:   p,
        fieldType:   FormFieldTypeText,
        name:        name,
    }
}
```

### Paragraph.InsertCheckBox

```go
func (p *Paragraph) InsertCheckBox(name string, checked bool) *FormField {
    // Create SimpleField
    simpleField := elements.NewSimpleField()
    simpleField.SetInstr(types.NewStringValue(fieldCodeCheckBox))

    // Create FormFieldData
    ffData := elements.NewFormFieldData()

    // Set name
    nameElem := elements.NewFormFieldName()
    nameElem.SetVal(types.NewStringValue(name))
    ffData.AppendChild(nameElem)

    // Set enabled
    enabled := elements.NewEnabled()
    ffData.AppendChild(enabled)

    // Create CheckBox
    checkBox := elements.NewCheckBox()
    checkBox.SetSize(types.NewInt32Value(20)) // 10pt default
    checkBox.SetSizeAuto(types.NewOnOffValue(false))

    checkedVal := 0
    if checked {
        checkedVal = 1
    }
    checkBox.SetDefault(types.NewInt32Value(checkedVal))
    checkBox.SetChecked(types.NewInt32Value(checkedVal))
    ffData.AppendChild(checkBox)

    // Add ffData to simpleField
    simpleField.AppendChild(ffData)

    // Create run with checkbox symbol
    run := elements.NewRun()
    sym := elements.NewSymbolChar()
    sym.SetFont(types.NewStringValue("MS Gothic"))
    if checked {
        sym.SetChar(types.NewStringValue("F052"))
    } else {
        sym.SetChar(types.NewStringValue("F06F"))
    }
    run.AppendChild(sym)
    simpleField.AppendChild(run)

    // Add to paragraph
    p.element.AppendChild(simpleField)

    return &FormField{
        simpleField: simpleField,
        ffData:      ffData,
        paragraph:   p,
        fieldType:   FormFieldTypeCheckBox,
        name:        name,
    }
}
```

### Paragraph.InsertDropDown

```go
func (p *Paragraph) InsertDropDown(name string, items []string, selectedIndex int) *FormField {
    // Validate dropdown item count (OpenXML schema limits to 25 items maximum)
    if len(items) > 25 {
        panic("dropdown form fields cannot have more than 25 items (OpenXML schema constraint)")
    }

    // Create SimpleField
    simpleField := elements.NewSimpleField()
    simpleField.SetInstr(types.NewStringValue(fieldCodeDropDown))

    // Create FormFieldData
    ffData := elements.NewFormFieldData()

    // Set name
    nameElem := elements.NewFormFieldName()
    nameElem.SetVal(types.NewStringValue(name))
    ffData.AppendChild(nameElem)

    // Set enabled
    enabled := elements.NewEnabled()
    ffData.AppendChild(enabled)

    // Create DropDownList
    ddList := elements.NewDropDownList()
    ddList.SetResult(types.NewInt32Value(selectedIndex))

    // Add items
    for _, item := range items {
        listEntry := elements.NewListEntry()
        listEntry.SetVal(types.NewStringValue(item))
        ddList.AppendChild(listEntry)
    }
    ffData.AppendChild(ddList)

    // Add ffData to simpleField
    simpleField.AppendChild(ffData)

    // Create run with selected item text
    run := elements.NewRun()
    text := elements.NewText()
    if selectedIndex >= 0 && selectedIndex < len(items) {
        text.SetText(items[selectedIndex])
    }
    run.AppendChild(text)
    simpleField.AppendChild(run)

    // Add to paragraph
    p.element.AppendChild(simpleField)

    return &FormField{
        simpleField: simpleField,
        ffData:      ffData,
        paragraph:   p,
        fieldType:   FormFieldTypeDropDown,
        name:        name,
    }
}
```

## Document-Level API

### GetFormFields Iterator

```go
func (d *Document) GetFormFields() iter.Seq[*FormField] {
    return func(yield func(*FormField) bool) {
        // Iterate all paragraphs in body
        body := d.MainDocumentPart().Document().Body()
        for para := range body.Descendants[*Paragraph]() {
            // Find SimpleField elements with ffData
            for simpleField := range para.Descendants[*elements.SimpleField]() {
                // Check if it has FormFieldData child
                ffDataList := simpleField.Elements[elements.FormFieldData]()
                if len(ffDataList) == 0 {
                    continue
                }

                ffData := ffDataList[0]

                // Determine type from instruction
                instr := simpleField.Instr()
                if instr == nil {
                    continue
                }

                instrText := instr.Value()
                var fieldType FormFieldType
                switch instrText {
                case fieldCodeText:
                    fieldType = FormFieldTypeText
                case fieldCodeCheckBox:
                    fieldType = FormFieldTypeCheckBox
                case fieldCodeDropDown:
                    fieldType = FormFieldTypeDropDown
                default:
                    continue // Not a form field
                }

                // Get name
                var name string
                nameElems := ffData.Elements[elements.FormFieldName]()
                if len(nameElems) > 0 {
                    if nameVal := nameElems[0].Val(); nameVal != nil {
                        name = nameVal.Value()
                    }
                }

                // Create FormField wrapper
                field := &FormField{
                    simpleField: simpleField,
                    ffData:      ffData,
                    paragraph:   para,
                    fieldType:   fieldType,
                    name:        name,
                }

                if !yield(field) {
                    return
                }
            }
        }
    }
}
```

### GetFormFieldByName

```go
func (d *Document) GetFormFieldByName(name string) *FormField {
    for field := range d.GetFormFields() {
        if field.Name() == name {
            return field
        }
    }
    return nil
}
```

## Document Protection API

Form fields are typically used with document protection that allows only form filling (not general editing).

### Protection XML Structure

```xml
<w:documentProtection
    w:edit="forms"
    w:enforcement="1"
    w:cryptProviderType="rsaFull"
    w:cryptAlgorithmClass="hash"
    w:cryptAlgorithmType="typeAny"
    w:cryptAlgorithmSid="4"
    w:cryptSpinCount="100000"
    w:hash="..."
    w:salt="..."/>
```

### ProtectForForms Method

```go
func (d *Document) ProtectForForms() {
    settings := d.MainDocumentPart().DocumentSettingsPart()
    if settings == nil {
        settings = d.MainDocumentPart().AddDocumentSettingsPart()
    }

    settingsElem := settings.Settings()

    // Create or update documentProtection element
    protections := settingsElem.Elements[elements.DocumentProtection]()
    var protection *elements.DocumentProtection
    if len(protections) > 0 {
        protection = protections[0]
    } else {
        protection = elements.NewDocumentProtection()
        settingsElem.AppendChild(protection)
    }

    // Set to forms-only editing
    protection.SetEdit(types.NewStringValue("forms"))
    protection.SetEnforcement(types.NewOnOffValue(true))
}
```

### UnprotectDocument Method

```go
func (d *Document) UnprotectDocument() {
    settings := d.MainDocumentPart().DocumentSettingsPart()
    if settings == nil {
        return
    }

    settingsElem := settings.Settings()

    // Remove documentProtection element
    protections := settingsElem.Elements[elements.DocumentProtection]()
    for _, protection := range protections {
        protection.Remove()
    }
}
```

## Testing Strategy

### Unit Tests (Fast)

```go
func TestInsertTextField(t *testing.T) {
    doc := createTestDocument()
    para := doc.MainDocumentPart().Document().Body().AppendParagraph()

    field := para.InsertTextField("TestField", "Default Value")

    assert.Equal(t, "TestField", field.Name())
    assert.Equal(t, FormFieldTypeText, field.Type())
    assert.Equal(t, "Default Value", field.GetTextValue())
}

func TestSetTextValue(t *testing.T) {
    doc := createTestDocument()
    para := doc.MainDocumentPart().Document().Body().AppendParagraph()
    field := para.InsertTextField("TestField", "Default")

    field.SetTextValue("New Value")

    assert.Equal(t, "New Value", field.GetTextValue())

    // Verify XML structure
    simpleField := field.simpleField
    runs := simpleField.Elements[elements.Run]()
    require.Len(t, runs, 1)
    texts := runs[0].Elements[elements.Text]()
    require.Len(t, texts, 1)
    assert.Equal(t, "New Value", texts[0].InnerText())
}

func TestCheckBoxToggle(t *testing.T) {
    doc := createTestDocument()
    para := doc.MainDocumentPart().Document().Body().AppendParagraph()
    field := para.InsertCheckBox("TestCheck", false)

    assert.False(t, field.IsChecked())

    field.SetChecked(true)
    assert.True(t, field.IsChecked())

    field.SetChecked(false)
    assert.False(t, field.IsChecked())
}

func TestDropDownSelection(t *testing.T) {
    doc := createTestDocument()
    para := doc.MainDocumentPart().Document().Body().AppendParagraph()
    items := []string{"Option A", "Option B", "Option C"}
    field := para.InsertDropDown("TestDD", items, 1)

    assert.Equal(t, 1, field.GetSelectedIndex())
    assert.Equal(t, "Option B", field.GetSelectedValue())

    field.SetSelectedIndex(2)
    assert.Equal(t, 2, field.GetSelectedIndex())
    assert.Equal(t, "Option C", field.GetSelectedValue())
}
```

### Roundtrip Tests (Integration)

```go
func TestFormFieldRoundtrip(t *testing.T) {
    // Create document with form fields
    doc1, _ := wordprocessing.Create("test.docx", wordprocessing.DocTypeDocument)
    para := doc1.MainDocumentPart().Document().Body().AppendParagraph()

    para.InsertTextField("Name", "John Doe")
    para.InsertCheckBox("Agree", true)
    para.InsertDropDown("Country", []string{"USA", "UK"}, 0)

    doc1.Save()
    doc1.Close()

    // Open and verify
    doc2, _ := wordprocessing.Open("test.docx")

    nameField := doc2.GetFormFieldByName("Name")
    require.NotNil(t, nameField)
    assert.Equal(t, "John Doe", nameField.GetTextValue())

    agreeField := doc2.GetFormFieldByName("Agree")
    require.NotNil(t, agreeField)
    assert.True(t, agreeField.IsChecked())

    countryField := doc2.GetFormFieldByName("Country")
    require.NotNil(t, countryField)
    assert.Equal(t, "USA", countryField.GetSelectedValue())

    doc2.Close()
}
```

### Compatibility Tests (Microsoft Word)

```go
func TestMicrosoftWordCompatibility(t *testing.T) {
    // Create form field document
    doc, _ := wordprocessing.Create("word_compat.docx", wordprocessing.DocTypeDocument)
    para := doc.MainDocumentPart().Document().Body().AppendParagraph()

    field := para.InsertTextField("TestField", "Test Value")
    field.SetMaxLength(50)

    doc.ProtectForForms()
    doc.Save()
    doc.Close()

    // Manual verification step:
    // 1. Open word_compat.docx in Microsoft Word
    // 2. Verify field appears and is editable
    // 3. Verify document is protected for forms
    // 4. Fill field and save
    // 5. Re-open with goffice and verify changes

    // This test requires manual verification but validates real-world usage
}
```

## Performance Considerations

### Iterator Performance

GetFormFields uses Go 1.23 iterators for memory efficiency:
- No need to allocate array of all fields upfront
- Caller can break early
- Lazy evaluation

### Caching Strategy

FormField caches name and type to avoid repeated XML traversal:

```go
type FormField struct {
    // ... other fields ...

    // Cached values (immutable after creation)
    fieldType FormFieldType
    name      string
}
```

Name and type don't change during field lifetime, safe to cache.

## Migration Path

### Phase 1: Core Implementation (2 weeks)
- FormField wrapper class
- Insertion methods (TextField, CheckBox, DropDown)
- Value get/set methods
- Basic tests

### Phase 2: Document-Level API (1 week)
- GetFormFields iterator
- GetFormFieldByName lookup
- Protection API
- Integration tests

### Phase 3: Advanced Features (1 week)
- Max length for text fields
- Checkbox size customization
- Help text / status text (optional)
- Advanced validation

### Phase 4: Polish & Documentation (0.5 weeks)
- Roundtrip tests
- Compatibility validation
- godoc comments
- Example code

**Total: 3-4 weeks**

## Risks & Mitigation

### Risk: XML Structure Misunderstanding
**Impact**: High (fields don't work in Word)
**Probability**: Medium (complex XML schema)
**Mitigation**:
- Reference Microsoft's Open-XML-SDK implementation
- Test with real Word documents from different versions
- Validate with Word 2007, 2013, 2016, 2019, Microsoft 365

### Risk: Symbol Font Compatibility
**Impact**: Medium (checkboxes render as squares/question marks)
**Probability**: Medium (font availability varies)
**Mitigation**:
- Test with multiple symbol fonts (Wingdings, MS Gothic, Symbol)
- Fallback to Unicode checkbox characters (U+2610, U+2612)
- Document font requirements

### Risk: Protection Breaking Editing
**Impact**: Medium (users can't edit protected docs)
**Probability**: Low (protection is well-defined)
**Mitigation**:
- Provide UnprotectDocument method
- Document protection behavior clearly
- Test protection roundtrips

### Risk: Field Code Parsing Edge Cases
**Impact**: Low (only affects advanced scenarios)
**Probability**: Low (we use simple fixed codes)
**Mitigation**:
- Start with simple field codes (no switches)
- Defer complex field code parsing to future enhancement

## Success Metrics

- [ ] All unit tests pass with >90% coverage
- [ ] Roundtrip tests pass (create → save → open → identical)
- [ ] Documents open correctly in Word 2007, 2013, 2016, 2019, Microsoft 365
- [ ] Form protection works in Microsoft Word
- [ ] Fields are editable when document is protected for forms
- [ ] golangci-lint passes with no errors
- [ ] godoc documentation complete and clear

## Future Enhancements (Not in Scope)

1. **Help Text**: `w:ffData/w:helpText` for field tooltips
2. **Status Bar Text**: `w:ffData/w:statusText` for status bar display
3. **Field Switches**: Advanced field code options (e.g., `FORMTEXT \* Upper`)
4. **Calculate on Exit**: `w:ffData/w:calcOnExit` for field calculations
5. **Modern Content Controls**: SDT-based forms (separate major feature)
6. **Field Events**: Callbacks when field values change
7. **Custom Validation**: Regular expressions, min/max values
8. **Field Linking**: Link form fields to document properties

These can be added incrementally without breaking the core API.
