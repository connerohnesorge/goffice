# Change: Add Word Form Fields API

## Why

The goffice library has **basic form field elements** but **no manipulation API**, blocking users from creating or modifying form-based Word documents programmatically. This is a common requirement for:

- **Business forms**: Purchase orders, invoices, contracts with fillable fields
- **Surveys and questionnaires**: Forms with text inputs, checkboxes, dropdowns
- **Templates**: Document templates with pre-defined fillable fields
- **Data entry workflows**: Forms that get filled programmatically from databases

**Current State**:
- Elements exist: `FieldCode`, `FieldChar`, `SimpleField`, `FormFieldData`, `DropDownListFormField`, `TextBoxFormFieldType`
- Missing: No methods to insert, modify, or retrieve form field values
- Gap: Cannot call `paragraph.InsertTextField("name", "default")` or `field.SetTextValue("value")`

**Impact**: **High (7/10)** - Common business use case, moderate implementation effort (3-4 weeks)

This proposal focuses on **legacy form fields** (Word 97-2003 compatible, using `fldSimple`/`ffData`) rather than modern content controls (SDT elements), which will be a separate future enhancement. Legacy fields are simpler, more backward compatible, and cover the majority of form use cases.

## What Changes

**New API Components**:
- `FormField` wrapper class providing high-level form field manipulation
- `InsertTextField`, `InsertCheckBox`, `InsertDropDown` methods on Paragraph
- `GetFormFields`, `GetFormFieldByName` methods on Document
- Form protection API (`ProtectForForms`, `UnprotectDocument`) on Document
- Field value setting/getting methods on FormField

**Core Design**:
- FormField wraps underlying SimpleField and FormFieldData elements
- Separate methods for field-type-specific operations (text max length, checkbox size, dropdown items)
- Field code construction helpers (FORMTEXT, FORMCHECKBOX, FORMDROPDOWN)
- Automatic synchronization between ffData and run text values

**Implementation Files**:
- `wordprocessing/form_fields.go` (NEW) - FormField wrapper and type enum
- `wordprocessing/document.go` (MODIFIED) - Add form field iteration/lookup
- `wordprocessing/paragraph.go` (MODIFIED) - Add form field insertion methods
- `wordprocessing/protection.go` (NEW) - Document protection for forms

**Testing**:
- Unit tests for each form field type
- Roundtrip tests (create → save → open → verify)
- Compatibility tests with Microsoft Word
- Value setting/getting validation

**Breaking changes**: None. This adds new functionality without changing existing APIs.

## Impact

**Affected specs**:
- `wordprocessing-document` (ADDED form field requirements)
- `wordprocessing-elements` (MODIFIED to note form field API, not just raw elements)

**New capabilities**:
- Insert text fields with default values and max length
- Insert checkboxes with checked/unchecked state and size
- Insert dropdown fields with multiple options and selected index
- Get/set form field values programmatically
- Iterate all form fields in document by name
- Protect documents for form filling only

**Affected code**:
- `wordprocessing/form_fields.go` - NEW: FormField wrapper, FormFieldType enum
- `wordprocessing/document.go` - MODIFIED: Add GetFormFields(), GetFormFieldByName()
- `wordprocessing/paragraph.go` - MODIFIED: Add InsertTextField(), InsertCheckBox(), InsertDropDown()
- `wordprocessing/protection.go` - NEW: ProtectForForms(), UnprotectDocument()
- `wordprocessing/elements/elements.go` - Used (not modified): Existing field elements
- `wordprocessing/form_fields_test.go` - NEW: Comprehensive test suite

**User-facing API example**:
```go
// Create form
doc, _ := wordprocessing.New("form.docx", wordprocessing.DocTypeDocument)
para := doc.MainDocumentPart().Document().Body().AppendParagraph()

nameField := para.InsertTextField("CustomerName", "Enter your name")
nameField.SetMaxLength(50)

agreeField := para.InsertCheckBox("AgreeTerms", false)
countryField := para.InsertDropDown("Country", []string{"USA", "Canada", "UK"}, 0)
countryField.SetDefaultDropDownIndex(0)  // Set default selection for form reset

doc.ProtectForForms()
doc.Save()

// Fill form
doc2, _ := wordprocessing.Open("form.docx")
for field := range doc2.GetFormFields() {
    if field.Name() == "CustomerName" {
        field.SetTextValue("John Doe")
    }
}
doc2.Save()
```

**Key constraints**:
- Dropdown fields limited to maximum 25 items (OpenXML schema constraint)
- Checkbox symbols MUST only be set via `SetChecked()` to ensure synchronization
- FormField uses `FormFieldName` element (not `FFName`)

## Dependencies

**Existing**:
- Field elements (FieldCode, FieldChar, SimpleField, FormFieldData) - already implemented
- Paragraph/Run infrastructure - already implemented
- Document save/open functionality - already implemented

**New**: None - this builds entirely on existing infrastructure

## Out of Scope

**Deferred to future proposals**:
- Modern content controls (SDT-based forms) - separate, more complex feature
- Form field macros/calculations - requires macro engine
- Help text and status bar text - can add in Phase 2
- Advanced validation rules - can add in Phase 2
- Form field events/callbacks - requires event system

## Success Criteria

- [ ] Can insert text fields with default values
- [ ] Can set/get text field values and max length
- [ ] Can insert checkboxes with checked state
- [ ] Can toggle checkbox state
- [ ] Can insert dropdowns with multiple items
- [ ] Can set selected dropdown item by index
- [ ] Can iterate all form fields in document
- [ ] Can get form field by name
- [ ] Can protect document for form filling only
- [ ] Fields work correctly when opened in Microsoft Word
- [ ] Roundtrip tests pass (create → save → open → verify identical)
- [ ] All unit tests pass with >90% coverage
- [ ] golangci-lint passes with no errors
