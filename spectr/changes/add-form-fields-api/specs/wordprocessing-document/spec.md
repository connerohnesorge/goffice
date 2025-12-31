# Wordprocessing Document Specification (Delta)

## ADDED Requirements

### Requirement: Form Field Iteration
The system SHALL provide iteration over all form fields in a document.

#### Scenario: Iterate all form fields
- GIVEN a WordprocessingDocument with multiple form fields
- WHEN GetFormFields() is called
- THEN an iterator yielding each FormField is returned

#### Scenario: Empty document iteration
- GIVEN a WordprocessingDocument with no form fields
- WHEN GetFormFields() is called
- THEN an empty iterator is returned (no yields)

#### Scenario: Early break from iteration
- GIVEN a document with 10 form fields
- WHEN GetFormFields() iterator breaks after 3 fields
- THEN only 3 fields are processed (lazy evaluation)

#### Scenario: Multiple field types in iteration
- GIVEN a document with text fields, checkboxes, and dropdowns
- WHEN GetFormFields() is called
- THEN all three types are yielded by the iterator

### Requirement: Form Field Lookup by Name
The system SHALL provide lookup of form fields by name.

#### Scenario: Find existing field by name
- GIVEN a document with a form field named "CustomerName"
- WHEN GetFormFieldByName("CustomerName") is called
- THEN the FormField with that name is returned

#### Scenario: Field not found
- GIVEN a document with no field named "MissingField"
- WHEN GetFormFieldByName("MissingField") is called
- THEN nil is returned

#### Scenario: Duplicate field names
- GIVEN a document with two fields both named "Field1"
- WHEN GetFormFieldByName("Field1") is called
- THEN the first matching field is returned

#### Scenario: Case-sensitive name matching
- GIVEN a document with a field named "TestField"
- WHEN GetFormFieldByName("testfield") is called
- THEN nil is returned (case-sensitive match)

### Requirement: Form Protection
The system SHALL support protecting documents for form filling only.

#### Scenario: Protect document for forms
- GIVEN a WordprocessingDocument with form fields
- WHEN ProtectForForms() is called
- THEN the document is protected with edit="forms" restriction

#### Scenario: Protection persists on save
- GIVEN a protected document
- WHEN Save() is called
- THEN the documentProtection element is written to settings part

#### Scenario: Unprotect document
- GIVEN a protected document
- WHEN UnprotectDocument() is called
- THEN the documentProtection element is removed from settings

#### Scenario: Protection does not affect programmatic access
- GIVEN a protected document
- WHEN form field values are set via API
- THEN values are updated successfully (protection is UI-level)

#### Scenario: Protect document without settings part
- GIVEN a document without a DocumentSettingsPart
- WHEN ProtectForForms() is called
- THEN a DocumentSettingsPart is created and protection is applied

### Requirement: Form Field Count
The system SHALL provide a count of form fields in a document.

#### Scenario: Count all form fields
- GIVEN a document with 5 form fields
- WHEN iterating GetFormFields() to count
- THEN 5 fields are counted

#### Scenario: Count by type
- GIVEN a document with 2 text fields, 3 checkboxes, 1 dropdown
- WHEN filtering GetFormFields() by type
- THEN correct counts for each type are obtained

### Requirement: Form Field Presence Check
The system SHALL provide checking if a form field exists.

#### Scenario: Check field exists by name
- GIVEN a document with a field named "TestField"
- WHEN GetFormFieldByName("TestField") != nil is checked
- THEN true is returned

#### Scenario: Check field does not exist
- GIVEN a document without a field named "MissingField"
- WHEN GetFormFieldByName("MissingField") != nil is checked
- THEN false is returned
