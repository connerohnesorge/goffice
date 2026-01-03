# Wordprocessing Document Specification

## Requirements

### Requirement: WordprocessingDocument Type
The system SHALL provide a WordprocessingDocument type as the entry point for Word document manipulation.

#### Scenario: Document type identity
- GIVEN a WordprocessingDocument
- WHEN DocumentType() is called
- THEN the type (Document, Template, MacroEnabledDocument, MacroEnabledTemplate) is returned

#### Scenario: Access main document part
- GIVEN a WordprocessingDocument
- WHEN MainDocumentPart() is called
- THEN the MainDocumentPart is returned (or nil if not present)

#### Scenario: Access document features
- GIVEN a WordprocessingDocument
- WHEN Features() is called
- THEN the document's feature collection is returned

### Requirement: Create WordprocessingDocument
The system SHALL support creating new Word documents.

#### Scenario: Create document from path
- GIVEN a file path and document type
- WHEN WordprocessingDocument.Create(path, DocumentType.Document) is called
- THEN a new .docx file is created with basic structure

#### Scenario: Create document to stream
- GIVEN an io.Writer and document type
- WHEN WordprocessingDocument.Create(writer, DocumentType.Document) is called
- THEN a new document is created writing to the stream

#### Scenario: Create with auto-save
- GIVEN create options with AutoSave=true
- WHEN the document is modified
- THEN changes are saved automatically on close

#### Scenario: Create template
- GIVEN DocumentType.Template
- WHEN Create() is called
- THEN a .dotx template document is created

#### Scenario: Create macro-enabled document
- GIVEN DocumentType.MacroEnabledDocument
- WHEN Create() is called
- THEN a .docm document is created with appropriate content type

#### Scenario: Create macro-enabled template
- GIVEN DocumentType.MacroEnabledTemplate
- WHEN Create() is called
- THEN a .dotm template is created

### Requirement: Open WordprocessingDocument
The system SHALL support opening existing Word documents.

#### Scenario: Open document from path (read-only)
- GIVEN a path to an existing .docx file
- WHEN WordprocessingDocument.Open(path, false) is called
- THEN the document is opened in read-only mode

#### Scenario: Open document from path (editable)
- GIVEN a path to an existing .docx file
- WHEN WordprocessingDocument.Open(path, true) is called
- THEN the document is opened for editing

#### Scenario: Open document from stream
- GIVEN an io.ReaderAt with document content
- WHEN WordprocessingDocument.OpenReader(reader, size, isEditable) is called
- THEN the document is opened from the stream

#### Scenario: Open with settings
- GIVEN OpenSettings with custom configuration
- WHEN Open(path, isEditable, settings) is called
- THEN the settings (AutoSave, MaxCharacters, etc.) are applied

#### Scenario: Open template
- GIVEN a .dotx file path
- WHEN Open() is called
- THEN the template is opened with DocumentType.Template

#### Scenario: Open macro-enabled documents
- GIVEN a .docm or .dotm file
- WHEN Open() is called
- THEN the appropriate document type is detected

### Requirement: Create from Template
The system SHALL support creating documents from templates.

#### Scenario: Create from template (attached)
- GIVEN a template path and isTemplate=true
- WHEN CreateFromTemplate(templatePath, true) is called
- THEN a new document is created with template content
- AND the template is attached to the document

#### Scenario: Create from template (detached)
- GIVEN a template path and isTemplate=false
- WHEN CreateFromTemplate(templatePath, false) is called
- THEN a new document is created with template content
- AND no template attachment exists

### Requirement: Save WordprocessingDocument
The system SHALL support saving Word documents.

#### Scenario: Save in place
- GIVEN an opened editable document
- WHEN Save() is called
- THEN all changes are persisted to the original location

#### Scenario: Save as new file
- GIVEN an opened document
- WHEN SaveAs(newPath) is called
- THEN a copy is saved to the new location
- AND the original file is unchanged

#### Scenario: Save to stream
- GIVEN an opened document
- WHEN SaveTo(writer) is called
- THEN the document is written to the stream

#### Scenario: Auto-save on close
- GIVEN a document opened with AutoSave=true
- WHEN Close() is called
- THEN unsaved changes are automatically saved

### Requirement: Close WordprocessingDocument
The system SHALL properly close and release document resources.

#### Scenario: Close releases resources
- GIVEN an open document
- WHEN Close() is called
- THEN all file handles and resources are released

#### Scenario: Close with pending changes (no auto-save)
- GIVEN a document with unsaved changes and AutoSave=false
- WHEN Close() is called
- THEN changes are discarded
- AND resources are released

#### Scenario: Implement io.Closer
- GIVEN WordprocessingDocument
- WHEN used in defer doc.Close() pattern
- THEN proper cleanup occurs

### Requirement: Change Document Type
The system SHALL support changing between document types.

#### Scenario: Document to template
- GIVEN a WordprocessingDocument of type Document
- WHEN ChangeDocumentType(DocumentType.Template) is called
- THEN the document type and content types are updated

#### Scenario: Template to document
- GIVEN a WordprocessingDocument of type Template
- WHEN ChangeDocumentType(DocumentType.Document) is called
- THEN the document type and content types are updated

#### Scenario: Enable macros
- GIVEN a Document
- WHEN ChangeDocumentType(DocumentType.MacroEnabledDocument) is called
- THEN the document can contain macros

### Requirement: Document Part Access
The system SHALL provide access to all document parts.

#### Scenario: Add main document part
- GIVEN a new WordprocessingDocument without main part
- WHEN AddMainDocumentPart() is called
- THEN a MainDocumentPart is created and added

#### Scenario: Access core properties
- GIVEN a WordprocessingDocument
- WHEN CoreFilePropertiesPart() is called
- THEN the core properties part (title, author, etc.) is returned

#### Scenario: Access extended properties
- GIVEN a WordprocessingDocument
- WHEN ExtendedFilePropertiesPart() is called
- THEN the extended properties part (application, company, etc.) is returned

#### Scenario: Access custom properties
- GIVEN a WordprocessingDocument
- WHEN CustomFilePropertiesPart() is called
- THEN the custom properties part is returned

#### Scenario: Access thumbnail
- GIVEN a WordprocessingDocument with thumbnail
- WHEN ThumbnailPart() is called
- THEN the thumbnail image part is returned

### Requirement: Document-Level Validation
The system SHALL support validating the entire document.

#### Scenario: Validate document
- GIVEN a WordprocessingDocument
- WHEN Validate() is called
- THEN all parts and elements are validated against schema

#### Scenario: Validate for specific version
- GIVEN a WordprocessingDocument
- WHEN Validate(FileFormatVersions.Office2016) is called
- THEN validation uses Office 2016 schema rules

### Requirement: Document Settings
The system SHALL support document-level settings and preferences.

#### Scenario: Get/Set document ID
- GIVEN a WordprocessingDocument
- WHEN DocumentId() and SetDocumentId() are called
- THEN the unique document identifier is accessed/modified

#### Scenario: Default file format version
- GIVEN document settings
- WHEN TargetFileFormatVersion is accessed
- THEN the target Office version for saving is returned

### Requirement: Builder Pattern Support
The system SHALL support the builder pattern for document creation.

#### Scenario: Create with builder
- GIVEN WordprocessingDocument.CreateBuilder()
- WHEN UseSettings() and Build() are chained
- THEN a configured document factory is created

#### Scenario: Default builder
- GIVEN CreateDefaultBuilder()
- WHEN Build().Create(path, type) is called
- THEN a document with default settings is created

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

### Requirement: Revision Enumeration
The system SHALL provide methods to enumerate all tracked changes in a document.

#### Scenario: Get all revisions from document
- GIVEN a WordprocessingDocument with 5 tracked changes (3 insertions, 2 deletions)
- WHEN GetRevisions() is called
- THEN an iterator over all 5 Revision wrappers is returned
- AND each Revision provides access to metadata and operations

#### Scenario: Get revisions from empty document
- GIVEN a WordprocessingDocument with no tracked changes
- WHEN GetRevisions() is called
- THEN an empty iterator is returned
- AND iterating completes immediately without errors

#### Scenario: Revisions span all document parts
- GIVEN a document with changes in body, headers, footers, and footnotes
- WHEN GetRevisions() is called
- THEN all changes from all parts are included in the iterator
- AND changes from headers are included
- AND changes from footers are included
- AND changes from footnotes are included

### Requirement: Accept All Revisions
The system SHALL support accepting all tracked changes in a document.

#### Scenario: Accept all revisions
- GIVEN a document with 5 tracked changes
- WHEN AcceptAllRevisions() is called
- THEN all InsertedRun elements are replaced with their Run children
- AND all DeletedRun elements are removed
- AND all MoveFrom/MoveTo pairs are finalized
- AND GetRevisions() returns an empty iterator
- AND no error is returned

#### Scenario: Accept all on clean document
- GIVEN a document with no tracked changes
- WHEN AcceptAllRevisions() is called
- THEN no changes are made to the document
- AND no error is returned

#### Scenario: Accept all with move pairs
- GIVEN a document with MoveFromRun id=1 and MoveToRun id=1
- WHEN AcceptAllRevisions() is called
- THEN MoveFromRun is removed (content moved away)
- AND MoveToRun is replaced with its Run children (content stays in new location)
- AND move is finalized

### Requirement: Reject All Revisions
The system SHALL support rejecting all tracked changes in a document.

#### Scenario: Reject all revisions
- GIVEN a document with 5 tracked changes
- WHEN RejectAllRevisions() is called
- THEN all InsertedRun elements are removed
- AND all DeletedRun elements are converted back to normal text
- AND all MoveFrom/MoveTo pairs are reverted
- AND GetRevisions() returns an empty iterator
- AND no error is returned

#### Scenario: Reject all on clean document
- GIVEN a document with no tracked changes
- WHEN RejectAllRevisions() is called
- THEN no changes are made to the document
- AND no error is returned

#### Scenario: Reject all restores deleted content
- GIVEN a document with DeletedRun containing "deleted text"
- WHEN RejectAllRevisions() is called
- THEN "deleted text" appears as normal text in the document
- AND DeletedRun wrapper is removed

### Requirement: Filter Revisions by Author
The system SHALL support accepting or rejecting revisions from specific authors.

#### Scenario: Accept revisions by author
- GIVEN a document with changes from "John Doe" and "Jane Smith"
- WHEN AcceptRevisionsByAuthor("John Doe") is called
- THEN all changes authored by "John Doe" are accepted
- AND changes authored by "Jane Smith" remain as tracked changes
- AND no error is returned

#### Scenario: Reject revisions by author
- GIVEN a document with changes from "John Doe" and "Jane Smith"
- WHEN RejectRevisionsByAuthor("Jane Smith") is called
- THEN all changes authored by "Jane Smith" are rejected
- AND changes authored by "John Doe" remain as tracked changes
- AND no error is returned

#### Scenario: Filter by non-existent author
- GIVEN a document with changes from various authors
- WHEN AcceptRevisionsByAuthor("NonExistent") is called
- THEN no changes are made to the document
- AND no error is returned

### Requirement: Revision Metadata Access
The system SHALL provide access to revision metadata through the Revision wrapper.

#### Scenario: Access revision author
- GIVEN a Revision wrapping InsertedRun with author="John Doe"
- WHEN revision.Author() is called
- THEN "John Doe" is returned

#### Scenario: Access revision date
- GIVEN a Revision wrapping InsertedRun with date="2024-01-15T10:30:00Z"
- WHEN revision.Date() is called
- THEN time.Time representing 2024-01-15 10:30:00 UTC is returned

#### Scenario: Access revision ID
- GIVEN a Revision wrapping DeletedRun with id=42
- WHEN revision.Id() is called
- THEN 42 is returned

#### Scenario: Access revision type
- GIVEN a Revision wrapping InsertedRun
- WHEN revision.Type() is called
- THEN RevisionTypeInsert is returned

### Requirement: Individual Revision Operations
The system SHALL support accepting or rejecting individual revisions.

#### Scenario: Accept individual insertion
- GIVEN a Revision wrapping InsertedRun containing "new text"
- WHEN revision.Accept() is called
- THEN the Run elements inside InsertedRun are promoted to parent
- AND the InsertedRun wrapper is removed
- AND "new text" remains in the document
- AND no error is returned

#### Scenario: Reject individual insertion
- GIVEN a Revision wrapping InsertedRun containing "new text"
- WHEN revision.Reject() is called
- THEN the InsertedRun and all its content are removed
- AND "new text" does not appear in the document
- AND no error is returned

#### Scenario: Accept individual deletion
- GIVEN a Revision wrapping DeletedRun containing "removed text"
- WHEN revision.Accept() is called
- THEN the DeletedRun and its content are removed
- AND "removed text" does not appear in the document
- AND no error is returned

#### Scenario: Reject individual deletion
- GIVEN a Revision wrapping DeletedRun containing DeletedText "removed text"
- WHEN revision.Reject() is called
- THEN "removed text" is restored as normal text
- AND the DeletedRun wrapper is removed
- AND "removed text" appears in the document
- AND no error is returned

### Requirement: Move Revision Handling
The system SHALL handle MoveFrom and MoveTo revision pairs correctly.

#### Scenario: Accept move pair
- GIVEN a MoveFromRun with id=10 at position A
- AND a MoveToRun with id=10 at position B
- WHEN both revisions are accepted
- THEN content is removed from position A
- AND content appears at position B as normal text
- AND both MoveFromRun and MoveToRun wrappers are removed

#### Scenario: Reject move pair
- GIVEN a MoveFromRun with id=10 at position A
- AND a MoveToRun with id=10 at position B
- WHEN both revisions are rejected
- THEN content is restored at position A as normal text
- AND content is removed from position B
- AND both MoveFromRun and MoveToRun wrappers are removed

#### Scenario: Accept orphaned move
- GIVEN a MoveFromRun with id=10 without matching MoveToRun
- WHEN revision.Accept() is called
- THEN an error is returned indicating orphaned move
- AND the document is not modified

### Requirement: Format Change Revision Handling
The system SHALL handle RunPropertiesChange and ParagraphPropertiesChange revisions.

#### Scenario: Accept format change
- GIVEN a RunPropertiesChange showing bold formatting added
- WHEN revision.Accept() is called
- THEN the new formatting (bold) is applied
- AND the revision wrapper is removed
- AND no error is returned

#### Scenario: Reject format change
- GIVEN a RunPropertiesChange showing bold formatting added
- WHEN revision.Reject() is called
- THEN the original formatting (non-bold) is restored
- AND the revision wrapper is removed
- AND no error is returned

### Requirement: Default Styles Generation
New Wordprocessing documents MUST contain a `styles.xml` part with essential default styles.

#### Scenario: Create new document has styles
- WHEN a new `wordprocessing.Document` is created via `New()`
- THEN the document package MUST contain a `styles.xml` part
- AND the `styles.xml` part MUST contain `w:docDefaults`
- AND the `styles.xml` part MUST contain a `w:style` with `w:styleId="Normal"`
