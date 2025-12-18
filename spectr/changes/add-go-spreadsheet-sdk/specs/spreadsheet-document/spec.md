## ADDED Requirements

### Requirement: SpreadsheetDocument Type

The system SHALL provide a `SpreadsheetDocument` type that serves as the main entry point for creating, opening, and manipulating Excel workbooks.

#### Scenario: Create new workbook from file path
- GIVEN a valid file path
- WHEN `spreadsheet.Create(path, SpreadsheetDocumentType.Workbook)` is called
- THEN a new SpreadsheetDocument is created at the specified path
- AND the document is opened in read-write mode
- AND the document type is set to Workbook (.xlsx)

#### Scenario: Create new workbook from stream
- GIVEN a writable io.Writer
- WHEN `spreadsheet.CreateFromStream(stream, SpreadsheetDocumentType.Workbook)` is called
- THEN a new SpreadsheetDocument is created writing to the stream
- AND the document is opened in write mode

#### Scenario: Open existing workbook for reading
- GIVEN a path to an existing .xlsx file
- WHEN `spreadsheet.Open(path, false)` is called
- THEN the SpreadsheetDocument is opened in read-only mode
- AND the document type is detected from content types

#### Scenario: Open existing workbook for editing
- GIVEN a path to an existing .xlsx file
- WHEN `spreadsheet.Open(path, true)` is called
- THEN the SpreadsheetDocument is opened in read-write mode
- AND modifications can be saved

#### Scenario: Save workbook changes
- GIVEN an open SpreadsheetDocument in read-write mode
- WHEN `document.Save()` is called
- THEN all pending changes are written to the package
- AND the document remains open for further modifications

#### Scenario: Save workbook to new location
- GIVEN an open SpreadsheetDocument
- WHEN `document.SaveAs(newPath)` is called
- THEN the document is saved to the new location
- AND the original file is unchanged

#### Scenario: Close workbook
- GIVEN an open SpreadsheetDocument
- WHEN `document.Close()` is called
- THEN all resources are released
- AND the underlying package is closed
- AND the document cannot be used further

### Requirement: SpreadsheetDocumentType Enum

The system SHALL provide a `SpreadsheetDocumentType` enumeration supporting all Excel document types.

#### Scenario: Workbook type
- GIVEN a SpreadsheetDocumentType of Workbook
- THEN the content type is `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml`
- AND the file extension is `.xlsx`

#### Scenario: Template type
- GIVEN a SpreadsheetDocumentType of Template
- THEN the content type is `application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml`
- AND the file extension is `.xltx`

#### Scenario: MacroEnabledWorkbook type
- GIVEN a SpreadsheetDocumentType of MacroEnabledWorkbook
- THEN the content type is `application/vnd.ms-excel.sheet.macroEnabled.main+xml`
- AND the file extension is `.xlsm`

#### Scenario: MacroEnabledTemplate type
- GIVEN a SpreadsheetDocumentType of MacroEnabledTemplate
- THEN the content type is `application/vnd.ms-excel.template.macroEnabled.main+xml`
- AND the file extension is `.xltm`

#### Scenario: AddIn type
- GIVEN a SpreadsheetDocumentType of AddIn
- THEN the content type is `application/vnd.ms-excel.addin.macroEnabled.main+xml`
- AND the file extension is `.xlam`

### Requirement: Document Type Conversion

The system SHALL support converting between document types.

#### Scenario: Convert workbook to template
- GIVEN a SpreadsheetDocument of type Workbook
- WHEN `document.ChangeDocumentType(SpreadsheetDocumentType.Template)` is called
- THEN the document type is changed to Template
- AND the main part content type is updated

#### Scenario: Convert template to macro-enabled
- GIVEN a SpreadsheetDocument of type Template
- WHEN `document.ChangeDocumentType(SpreadsheetDocumentType.MacroEnabledTemplate)` is called
- THEN the document type is changed to MacroEnabledTemplate
- AND the main part content type is updated
- AND a VbaProjectPart can be added

### Requirement: FlatOPC Format Support

The system SHALL support creating and opening FlatOPC format documents.

#### Scenario: Export to FlatOPC
- GIVEN an open SpreadsheetDocument
- WHEN `document.ToFlatOpcDocument()` is called
- THEN a single XML document is returned containing all parts
- AND the XML follows the FlatOPC schema

#### Scenario: Import from FlatOPC
- GIVEN a FlatOPC XML document
- WHEN `spreadsheet.FromFlatOpcDocument(xml)` is called
- THEN a SpreadsheetDocument is created from the XML
- AND all parts are reconstructed

### Requirement: Template Creation

The system SHALL support creating documents from templates.

#### Scenario: Create from template file
- GIVEN a path to a .xltx template file
- WHEN `spreadsheet.CreateFromTemplate(templatePath)` is called
- THEN a new SpreadsheetDocument is created based on the template
- AND the document type is Workbook (not Template)
- AND all template content is preserved

### Requirement: Document Validation

The system SHALL provide document-level validation.

#### Scenario: Validate document structure
- GIVEN an open SpreadsheetDocument
- WHEN `document.Validate(FileFormatVersions.Office2016)` is called
- THEN all parts are validated against the specified version
- AND schema validation is performed on all XML parts
- AND semantic constraints are checked
- AND validation errors are returned

### Requirement: Semantic Validation Constraints

The system SHALL implement all 19+ semantic constraint types for comprehensive validation.

#### Scenario: AttributeValueRange constraint
- GIVEN an element with a numeric attribute
- WHEN the attribute value is outside min/max bounds
- THEN a validation error is reported

#### Scenario: AttributeValuePattern constraint
- GIVEN an element with an attribute requiring pattern matching
- WHEN the attribute value does not match the regex pattern
- THEN a validation error is reported

#### Scenario: AttributeValueSet constraint
- GIVEN an element with an enumerated attribute
- WHEN the attribute value is not in the allowed set
- THEN a validation error is reported

#### Scenario: ParentType constraint
- GIVEN an element with restricted parent types
- WHEN the element's parent is not an allowed type
- THEN a validation error is reported

#### Scenario: ChildElement constraint
- GIVEN an element with required/allowed children
- WHEN required children are missing or disallowed children are present
- THEN a validation error is reported

#### Scenario: UniqueValue constraint
- GIVEN elements requiring unique values within a scope
- WHEN duplicate values exist
- THEN a validation error is reported

#### Scenario: RelationshipExist constraint
- GIVEN an element referencing a relationship by ID
- WHEN the referenced relationship does not exist
- THEN a validation error is reported

#### Scenario: RelationshipType constraint
- GIVEN an element referencing a relationship
- WHEN the relationship type is incorrect
- THEN a validation error is reported

#### Scenario: ReferenceExist constraint
- GIVEN an element referencing another element by ID
- WHEN the referenced element does not exist
- THEN a validation error is reported

#### Scenario: IndexRange constraint
- GIVEN an element with an index attribute
- WHEN the index is outside valid bounds
- THEN a validation error is reported

#### Scenario: RootAttribute constraint
- GIVEN a root element with required attributes
- WHEN a required attribute is missing
- THEN a validation error is reported

#### Scenario: AttributeAbsent constraint
- GIVEN mutually exclusive attributes
- WHEN both attributes are present
- THEN a validation error is reported

#### Scenario: AttributeCannotOmit constraint
- GIVEN a conditionally required attribute
- WHEN the condition is met but the attribute is missing
- THEN a validation error is reported

#### Scenario: AttributeValueLength constraint
- GIVEN an attribute with length limits
- WHEN the attribute value exceeds the length limit
- THEN a validation error is reported

#### Scenario: UniqueAttributeValue constraint
- GIVEN attributes that must be unique across the document
- WHEN duplicate attribute values exist
- THEN a validation error is reported

#### Scenario: PartContainer constraint
- GIVEN a part that must be in a specific container
- WHEN the part is in the wrong container
- THEN a validation error is reported

#### Scenario: DataPart constraint
- GIVEN a data part with specific requirements
- WHEN the data part requirements are not met
- THEN a validation error is reported

#### Scenario: PartType constraint
- GIVEN a part with a required content type
- WHEN the part content type is incorrect
- THEN a validation error is reported

#### Scenario: Custom constraint extensibility
- GIVEN a custom constraint implementation
- WHEN the constraint is registered
- THEN the constraint is applied during validation
