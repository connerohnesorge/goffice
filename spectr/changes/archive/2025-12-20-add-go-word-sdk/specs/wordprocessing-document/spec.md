## ADDED Requirements

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
