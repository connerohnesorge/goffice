# Implementation Tasks

## Phase 1: FormField Wrapper & Core Types (1 week)

### 1.1 Define FormField Type and Enum
- [ ] Create `wordprocessing/form_fields.go` with FormFieldType enum
- [ ] Define TextField, CheckBox, DropDown constants
- [ ] Create FormField struct with simpleField, ffData, paragraph, fieldType, name fields
- [ ] Add field code constants (FORMTEXT, FORMCHECKBOX, FORMDROPDOWN)
- [ ] Validation: Code compiles, types are exported correctly

### 1.2 Implement Universal FormField Methods
- [ ] Implement `Name() string` method
- [ ] Implement `SetName(name string)` method
- [ ] Implement `Type() FormFieldType` method
- [ ] Implement `Enabled() bool` method
- [ ] Implement `SetEnabled(enabled bool)` method
- [ ] Implement `Remove()` method to remove field from paragraph
- [ ] Validation: All methods compile and can be called

### 1.3 Implement TextField Methods
- [ ] Implement `GetTextValue() string` - read from run text
- [ ] Implement `SetTextValue(value string)` - update ffData AND run text
- [ ] Implement `GetMaxLength() int` - read from textInput/maxLength
- [ ] Implement `SetMaxLength(length int)` - update ffData textInput
- [ ] Implement `GetDefaultTextValue() string` - read from textInput/default
- [ ] Implement `SetDefaultTextValue(value string)` - update ffData default
- [ ] Validation: Text field values can be get/set correctly

### 1.4 Implement CheckBox Methods
- [ ] Implement `IsChecked() bool` - read from checkBox/checked
- [ ] Implement `SetChecked(checked bool)` - update ffData AND symbol character
- [ ] Implement `GetCheckBoxSize() int` - read from checkBox/size
- [ ] Implement `SetCheckBoxSize(size int)` - update checkBox size
- [ ] Implement `GetAutoCheckBoxSize() bool` - read from checkBox/sizeAuto
- [ ] Implement `SetAutoCheckBoxSize(auto bool)` - update sizeAuto
- [ ] Implement `GetDefaultChecked() bool` - read from checkBox/default
- [ ] Implement `SetDefaultChecked(checked bool)` - update default state
- [ ] Add checkbox symbol handling (F052 for checked, F06F for unchecked)
- [ ] Validation: Checkbox state can be toggled and symbols update

### 1.5 Implement DropDown Methods
- [ ] Implement `GetDropDownItems() []string` - read from ddList/listEntry elements
- [ ] Implement `SetDropDownItems(items []string)` - replace listEntry children with 25-item limit validation
- [ ] Implement `GetSelectedIndex() int` - read from ddList/result
- [ ] Implement `SetSelectedIndex(index int)` - update result AND run text
- [ ] Implement `GetSelectedValue() string` - return items[selectedIndex]
- [ ] Implement `GetDefaultDropDownIndex() int` - read from ddList/default
- [ ] Implement `SetDefaultDropDownIndex(index int)` - update ddList/default element
- [ ] Add validation: SetDropDownItems() rejects lists with more than 25 items
- [ ] Validation: Dropdown items and selection work correctly

### 1.6 Write FormField Unit Tests
- [ ] Test FormField creation with each type
- [ ] Test Name() and SetName() methods
- [ ] Test Type() detection from field code
- [ ] Test Enabled() and SetEnabled() methods
- [ ] Test GetTextValue() and SetTextValue() for text fields
- [ ] Test SetMaxLength() and verify XML structure
- [ ] Test IsChecked() and SetChecked() for checkboxes
- [ ] Test checkbox symbol updates when toggling state
- [ ] Test SetCheckBoxSize() and verify size in XML
- [ ] Test GetDropDownItems() and SetDropDownItems()
- [ ] Test SetDropDownItems() rejects 26+ items (validates 25-item limit)
- [ ] Test GetDefaultDropDownIndex() and SetDefaultDropDownIndex()
- [ ] Test SetSelectedIndex() and verify run text updates
- [ ] Test checkbox symbol validation: symbols match checked state
- [ ] Test Remove() method removes field from paragraph
- [ ] Validation: All tests pass, coverage >90%

## Phase 2: Paragraph Insertion Methods (1 week)

### 2.1 Implement Paragraph.InsertTextField
- [ ] Create method signature: `InsertTextField(name, defaultValue string) *FormField`
- [ ] Create SimpleField element with FORMTEXT instruction
- [ ] Create FormFieldData element with name and enabled
- [ ] Create TextInput element with type="regular" and default value
- [ ] Create Run element with Text containing default value
- [ ] Append all elements correctly (ffData in simpleField, run in simpleField)
- [ ] Add simpleField to paragraph
- [ ] Return FormField wrapper
- [ ] Validation: Method creates valid text field structure

### 2.2 Implement Paragraph.InsertCheckBox
- [ ] Create method signature: `InsertCheckBox(name string, checked bool) *FormField`
- [ ] Create SimpleField element with FORMCHECKBOX instruction
- [ ] Create FormFieldData element with name and enabled
- [ ] Create CheckBox element with size=20, sizeAuto=false
- [ ] Set default and checked values based on checked parameter
- [ ] Create Run element with SymbolChar (MS Gothic font, appropriate character)
- [ ] Set symbol to F052 (checked) or F06F (unchecked) based on state
- [ ] Append elements and add to paragraph
- [ ] Return FormField wrapper
- [ ] Validation: Method creates valid checkbox structure

### 2.3 Implement Paragraph.InsertDropDown
- [ ] Create method signature: `InsertDropDown(name string, items []string, selectedIndex int) *FormField`
- [ ] Add validation: panic if len(items) > 25 (OpenXML schema constraint)
- [ ] Create SimpleField element with FORMDROPDOWN instruction
- [ ] Create FormFieldData element with name and enabled
- [ ] Create DropDownList element with default=selectedIndex and result=selectedIndex
- [ ] Create ListEntry elements for each item in items array
- [ ] Create Run element with Text containing selected item text
- [ ] Append elements and add to paragraph
- [ ] Return FormField wrapper
- [ ] Validation: Method creates valid dropdown structure with both default and result elements

### 2.4 Write Insertion Unit Tests
- [ ] Test InsertTextField creates correct XML structure
- [ ] Test InsertTextField returns valid FormField wrapper
- [ ] Test InsertTextField with various default values (empty, long text)
- [ ] Test InsertCheckBox creates correct checkbox XML
- [ ] Test InsertCheckBox with checked=true and checked=false
- [ ] Test InsertCheckBox symbol characters are correct
- [ ] Test InsertDropDown creates correct dropdown XML with both default and result
- [ ] Test InsertDropDown with various item counts (1, 5, 20, 25 items)
- [ ] Test InsertDropDown panics with 26 items (validates max constraint)
- [ ] Test InsertDropDown with different selected indices
- [ ] Test multiple fields can be inserted in same paragraph
- [ ] Validation: All insertion tests pass

## Phase 3: Document-Level API (1 week)

### 3.1 Implement Document.GetFormFields Iterator
- [ ] Create method signature: `GetFormFields() iter.Seq[*FormField]`
- [ ] Implement iterator function using Go 1.23 iter.Seq pattern
- [ ] Iterate all paragraphs in document body
- [ ] Find SimpleField elements with FormFieldData children
- [ ] Detect field type from instruction attribute
- [ ] Extract field name from ffData/name element
- [ ] Create FormField wrapper for each found field
- [ ] Yield FormField to caller, respect break if !yield returns false
- [ ] Validation: Iterator finds all form fields in document

### 3.2 Implement Document.GetFormFieldByName
- [ ] Create method signature: `GetFormFieldByName(name string) *FormField`
- [ ] Call GetFormFields() iterator
- [ ] Check each field's Name() against target name
- [ ] Return first matching field
- [ ] Return nil if no match found
- [ ] Validation: Method finds field by name correctly

### 3.3 Implement Document Protection API
- [ ] Create `wordprocessing/protection.go` file
- [ ] Implement `ProtectForForms()` method on Document
- [ ] Access or create DocumentSettingsPart
- [ ] Create or update DocumentProtection element
- [ ] Set edit="forms" and enforcement="1"
- [ ] Implement `UnprotectDocument()` method
- [ ] Remove DocumentProtection element from settings
- [ ] Validation: Protection methods work correctly

### 3.4 Write Document-Level Unit Tests
- [ ] Test GetFormFields returns empty iterator for document with no fields
- [ ] Test GetFormFields finds all fields in document with multiple types
- [ ] Test GetFormFields iterator can break early
- [ ] Test GetFormFieldByName finds existing field
- [ ] Test GetFormFieldByName returns nil for non-existent name
- [ ] Test GetFormFieldByName with duplicate names (returns first)
- [ ] Test ProtectForForms creates documentProtection element
- [ ] Test ProtectForForms sets correct edit and enforcement attributes
- [ ] Test UnprotectDocument removes protection
- [ ] Test protection roundtrip (protect → save → open → verify)
- [ ] Validation: All document-level tests pass

## Phase 4: Integration & Roundtrip Tests (1 week)

### 4.1 Implement TextField Roundtrip Tests
- [ ] Create document with text field
- [ ] Set default value and max length
- [ ] Save document
- [ ] Close document
- [ ] Re-open document
- [ ] Verify field exists with correct name
- [ ] Verify default value matches
- [ ] Verify max length matches
- [ ] Change value and save again
- [ ] Re-open and verify change persisted
- [ ] Validation: Text field roundtrip works perfectly

### 4.2 Implement CheckBox Roundtrip Tests
- [ ] Create document with checkbox (checked)
- [ ] Save and close
- [ ] Re-open and verify checked state is true
- [ ] Create document with checkbox (unchecked)
- [ ] Save and close
- [ ] Re-open and verify checked state is false
- [ ] Toggle state, save, re-open, verify toggle persisted
- [ ] Verify checkbox symbol characters are correct after roundtrip
- [ ] Validation: Checkbox roundtrip preserves state

### 4.3 Implement DropDown Roundtrip Tests
- [ ] Create document with dropdown (5 items, default=0, result=2)
- [ ] Save and close
- [ ] Re-open and verify all items present
- [ ] Verify default index is 0
- [ ] Verify selected index is 2
- [ ] Verify selected value text matches items[2]
- [ ] Change selected index, save, re-open, verify change persisted
- [ ] Change default index, save, re-open, verify default persisted
- [ ] Test with empty dropdown (0 items) - should handle gracefully
- [ ] Validation: Dropdown roundtrip works correctly for both default and result

### 4.4 Implement Multi-Field Integration Tests
- [ ] Create document with all three field types in same paragraph
- [ ] Save, close, re-open
- [ ] Verify all fields exist and have correct types
- [ ] Iterate all fields with GetFormFields()
- [ ] Verify iterator returns all three fields
- [ ] Get each field by name with GetFormFieldByName()
- [ ] Modify each field's value
- [ ] Save, close, re-open, verify all modifications persisted
- [ ] Validation: Multiple fields coexist correctly

### 4.5 Implement Protection Integration Tests
- [ ] Create document with form fields
- [ ] Call ProtectForForms()
- [ ] Save and close
- [ ] Re-open document
- [ ] Verify documentProtection element exists in settings
- [ ] Verify edit="forms" attribute
- [ ] Verify enforcement="1" attribute
- [ ] Call UnprotectDocument()
- [ ] Save, close, re-open
- [ ] Verify documentProtection element is removed
- [ ] Validation: Protection persists across saves

### 4.6 Value Synchronization Tests
- [ ] Create text field and verify run text matches default value
- [ ] Call SetTextValue() and verify both ffData AND run text update
- [ ] Create checkbox and verify symbol character matches checked state
- [ ] Call SetChecked() and verify both ffData AND symbol update
- [ ] Create dropdown and verify run text shows selected item
- [ ] Call SetSelectedIndex() and verify both ffData result AND run text update
- [ ] Test edge case: empty string text value
- [ ] Test edge case: very long text value (500 characters)
- [ ] Validation: ffData and run text stay synchronized

## Phase 5: Microsoft Word Compatibility Testing (0.5 weeks)

### 5.1 Manual Word Compatibility Tests
- [ ] Create form document with goffice (all three field types)
- [ ] Open in Microsoft Word 2016/2019
- [ ] Verify all fields render correctly
- [ ] Verify fields are editable
- [ ] Fill fields manually in Word
- [ ] Save from Word
- [ ] Re-open with goffice
- [ ] Verify field values updated correctly
- [ ] Validation: Word can edit goffice-created forms

### 5.2 Protection Compatibility Tests
- [ ] Create protected form document with goffice
- [ ] Open in Microsoft Word
- [ ] Verify document is protected (only form fields editable)
- [ ] Verify paragraph text is NOT editable
- [ ] Verify form fields ARE editable
- [ ] Fill fields and save
- [ ] Re-open with goffice and verify changes
- [ ] Validation: Protection works in Microsoft Word

### 5.3 Symbol Font Compatibility
- [ ] Create checkbox fields with goffice
- [ ] Open in Word on Windows
- [ ] Verify checkboxes render as checkbox symbols (not squares/question marks)
- [ ] Open in Word on macOS
- [ ] Verify checkboxes render correctly
- [ ] Open in LibreOffice Writer
- [ ] Verify checkboxes render (may differ but should be recognizable)
- [ ] Validation: Checkbox symbols work across platforms

## Phase 6: Documentation & Examples (0.5 weeks)

### 6.1 Write godoc Documentation
- [ ] Add package-level documentation to wordprocessing/form_fields.go
- [ ] Document FormField type and all methods with examples
- [ ] Document FormFieldType enum values
- [ ] Document Paragraph insertion methods with usage examples
- [ ] Document Document methods (GetFormFields, GetFormFieldByName)
- [ ] Document protection methods with security notes
- [ ] Validation: `go doc wordprocessing.FormField` looks professional

### 6.2 Create Example Program
- [ ] Create `examples/form-fields/main.go`
- [ ] Demonstrate creating all three field types
- [ ] Demonstrate setting field values
- [ ] Demonstrate iterating all fields
- [ ] Demonstrate form protection
- [ ] Add comments explaining each step
- [ ] Validation: Example runs and creates valid form document

### 6.3 Create README for Examples
- [ ] Create `examples/form-fields/README.md`
- [ ] Explain what the example does
- [ ] Show sample output XML structure
- [ ] Link to relevant godoc
- [ ] Validation: README is helpful and clear

### 6.4 Update Main Repository Documentation
- [ ] Add form fields to feature list in main README
- [ ] Add code snippet showing form field usage
- [ ] Link to form fields example
- [ ] Validation: Main README mentions form fields

## Phase 7: Final Validation & Cleanup (0.5 weeks)

### 7.1 Run Full Test Suite
- [ ] Run `go test ./wordprocessing/...` and ensure all tests pass
- [ ] Run `go test -race ./wordprocessing/...` to check for race conditions
- [ ] Run `go test -cover ./wordprocessing/...` and verify >90% coverage
- [ ] Run `gotestsum --format short-verbose ./wordprocessing/...` for formatted output
- [ ] Validation: All tests pass, no race conditions, high coverage

### 7.2 Run Linting
- [ ] Run `golangci-lint run --fix ./wordprocessing/...`
- [ ] Fix any remaining linting errors
- [ ] Verify no golangci-lint errors remain
- [ ] Run `nix develop -c lint` from project root
- [ ] Validation: Linting passes with zero errors

### 7.3 Verify All Task Checklist Items
- [ ] Review ALL tasks in this file
- [ ] Ensure every checkbox is checked
- [ ] Re-run key tests to confirm completion
- [ ] Validation: All tasks genuinely complete

### 7.4 Update Spec Status
- [ ] Mark all tasks as completed in this file
- [ ] Update proposal status (if using status field)
- [ ] Validation: Documentation reflects completion

## Validation Checkpoints

**After Phase 1**:
- [ ] FormField type exists and compiles
- [ ] All get/set methods implemented
- [ ] Unit tests pass with >90% coverage

**After Phase 2**:
- [ ] Paragraph insertion methods work for all three types
- [ ] Insertion creates valid XML structures
- [ ] Insertion tests pass

**After Phase 3**:
- [ ] Document.GetFormFields() iterator works
- [ ] Document.GetFormFieldByName() finds fields
- [ ] Protection API works
- [ ] Document-level tests pass

**After Phase 4**:
- [ ] Roundtrip tests pass for all field types
- [ ] Multi-field integration works
- [ ] Protection persists across saves
- [ ] Value synchronization verified

**After Phase 5**:
- [ ] Forms open correctly in Microsoft Word
- [ ] Protection works in Microsoft Word
- [ ] Checkbox symbols render correctly

**After Phase 6**:
- [ ] godoc documentation complete
- [ ] Example program works
- [ ] README updated

**After Phase 7**:
- [ ] All tests pass
- [ ] Linting clean
- [ ] All tasks complete
- [ ] Ready for review/merge

## Estimated Timeline

- Phase 1: FormField Core - **1 week** (5 tasks, 6 subtasks)
- Phase 2: Insertion Methods - **1 week** (4 tasks, 11 subtasks)
- Phase 3: Document API - **1 week** (4 tasks, 10 subtasks)
- Phase 4: Integration Tests - **1 week** (6 tasks, 32 subtasks)
- Phase 5: Word Compatibility - **0.5 weeks** (3 tasks, 9 subtasks)
- Phase 6: Documentation - **0.5 weeks** (4 tasks, 4 subtasks)
- Phase 7: Final Validation - **0.5 weeks** (4 tasks, 4 subtasks)

**Total: 4.5 weeks** (conservative estimate)

## Success Criteria Summary

Before marking this change as complete, verify:
- [ ] All 76+ subtasks above are checked off
- [ ] All unit tests pass
- [ ] All integration tests pass
- [ ] All roundtrip tests pass
- [ ] golangci-lint reports zero errors
- [ ] Forms work correctly in Microsoft Word 2016/2019/Microsoft 365
- [ ] Forms work correctly in LibreOffice Writer (basic compatibility)
- [ ] godoc documentation is complete and helpful
- [ ] Example code runs and demonstrates all features
- [ ] Test coverage >90% for wordprocessing/form_fields.go
