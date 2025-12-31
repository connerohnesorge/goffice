# Spreadsheet Document Specification

## Requirements

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

### Requirement: Row Insert Operations
The system SHALL provide methods to insert rows into worksheets, shifting existing rows down and updating all cell references.

#### Scenario: Insert single row
- GIVEN a worksheet with data in rows 1-10
- WHEN `sheet.InsertRows(5, 1)` is called
- THEN a new empty row is inserted at row 5
- AND rows 5-10 shift down to rows 6-11
- AND all formulas referencing rows >= 5 are updated

#### Scenario: Insert multiple rows
- GIVEN a worksheet with data
- WHEN `sheet.InsertRows(10, 5)` is called
- THEN 5 new empty rows are inserted at row 10
- AND existing rows 10+ shift down by 5
- AND all cell references in formulas are updated accordingly

#### Scenario: Insert rows with formulas
- GIVEN a worksheet with formula `=SUM(A1:A10)` in cell B1
- WHEN `sheet.InsertRows(5, 3)` is called
- THEN the formula becomes `=SUM(A1:A13)`
- AND the range extends to include inserted rows

#### Scenario: Insert rows with absolute references
- GIVEN a formula `=$A$5+A10` in cell B1
- WHEN `sheet.InsertRows(5, 2)` is called
- THEN the formula becomes `=$A$5+A12`
- AND the absolute reference $A$5 is NOT shifted

#### Scenario: Insert rows updates named ranges
- GIVEN a named range `SalesData = Sheet1!A1:D10`
- WHEN `sheet.InsertRows(5, 3)` is called
- THEN the named range becomes `Sheet1!A1:D13`

#### Scenario: Insert rows updates merged cells
- GIVEN merged cells `A5:C5`
- WHEN `sheet.InsertRows(5, 2)` is called
- THEN the merged cells range becomes `A7:C7`

#### Scenario: Insert at first row
- GIVEN a worksheet with data
- WHEN `sheet.InsertRows(1, 3)` is called
- THEN 3 rows are inserted at the top
- AND all existing rows shift down by 3

#### Scenario: Insert rows validation
- GIVEN a worksheet
- WHEN `sheet.InsertRows(0, 1)` is called
- THEN an error is returned (index must be >= 1)

### Requirement: Row Delete Operations
The system SHALL provide methods to delete rows from worksheets, shifting remaining rows up and updating all cell references.

#### Scenario: Delete single row
- GIVEN a worksheet with data in rows 1-10
- WHEN `sheet.DeleteRows(5, 1)` is called
- THEN row 5 is deleted
- AND rows 6-10 shift up to rows 5-9
- AND all formulas referencing rows > 5 are updated

#### Scenario: Delete multiple rows
- GIVEN a worksheet with data
- WHEN `sheet.DeleteRows(5, 3)` is called
- THEN rows 5, 6, 7 are deleted
- AND rows 8+ shift up by 3
- AND all cell references are updated accordingly

#### Scenario: Delete rows with formula references
- GIVEN a formula `=SUM(A1:A10)` in cell B1
- WHEN `sheet.DeleteRows(8, 3)` is called
- THEN the formula becomes `=SUM(A1:A7)`
- AND the range shrinks to exclude deleted rows

#### Scenario: Delete rows containing formula references
- GIVEN a formula `=A5+A6+A7` in cell B1
- WHEN `sheet.DeleteRows(5, 3)` is called (deletes rows 5, 6, 7)
- THEN the formula becomes `=#REF!+#REF!+#REF!`
- AND references to deleted cells become #REF! errors

#### Scenario: Delete rows updates named ranges
- GIVEN a named range `SalesData = Sheet1!A1:D10`
- WHEN `sheet.DeleteRows(8, 3)` is called
- THEN the named range becomes `Sheet1!A1:D7`

#### Scenario: Delete rows updates merged cells
- GIVEN merged cells `A8:C8`
- WHEN `sheet.DeleteRows(5, 2)` is called
- THEN the merged cells range becomes `A6:C6`

#### Scenario: Delete rows validation
- GIVEN a worksheet
- WHEN `sheet.DeleteRows(0, 1)` is called
- THEN an error is returned (index must be >= 1)

### Requirement: Column Insert Operations
The system SHALL provide methods to insert columns into worksheets, shifting existing columns right and updating all cell references.

#### Scenario: Insert single column
- GIVEN a worksheet with data in columns A-J
- WHEN `sheet.InsertColumns(5, 1)` is called (insert at column E)
- THEN a new empty column is inserted at column E
- AND columns E-J shift right to columns F-K
- AND all formulas referencing columns >= E are updated

#### Scenario: Insert multiple columns
- GIVEN a worksheet with data
- WHEN `sheet.InsertColumns(3, 5)` is called (insert at column C)
- THEN 5 new empty columns are inserted at column C
- AND existing columns C+ shift right by 5
- AND all cell references in formulas are updated accordingly

#### Scenario: Insert columns with formulas
- GIVEN a formula `=SUM(A1:E1)` in cell F1
- WHEN `sheet.InsertColumns(3, 2)` is called (insert at column C)
- THEN the formula becomes `=SUM(A1:G1)`
- AND the range extends to include inserted columns

#### Scenario: Insert columns with absolute references
- GIVEN a formula `=$C$1+E1` in cell F1
- WHEN `sheet.InsertColumns(3, 2)` is called
- THEN the formula becomes `=$C$1+G1`
- AND the absolute reference $C$1 is NOT shifted

#### Scenario: Insert columns updates named ranges
- GIVEN a named range `QuarterData = Sheet1!A1:D100`
- WHEN `sheet.InsertColumns(3, 2)` is called
- THEN the named range becomes `Sheet1!A1:F100`

#### Scenario: Insert columns updates merged cells
- GIVEN merged cells `C5:E5`
- WHEN `sheet.InsertColumns(3, 2)` is called
- THEN the merged cells range becomes `E5:G5`

#### Scenario: Insert at first column
- GIVEN a worksheet with data
- WHEN `sheet.InsertColumns(1, 3)` is called
- THEN 3 columns are inserted at the left
- AND all existing columns shift right by 3

#### Scenario: Insert columns validation
- GIVEN a worksheet
- WHEN `sheet.InsertColumns(0, 1)` is called
- THEN an error is returned (index must be >= 1)

### Requirement: Column Delete Operations
The system SHALL provide methods to delete columns from worksheets, shifting remaining columns left and updating all cell references.

#### Scenario: Delete single column
- GIVEN a worksheet with data in columns A-J
- WHEN `sheet.DeleteColumns(5, 1)` is called (delete column E)
- THEN column E is deleted
- AND columns F-J shift left to columns E-I
- AND all formulas referencing columns > E are updated

#### Scenario: Delete multiple columns
- GIVEN a worksheet with data
- WHEN `sheet.DeleteColumns(3, 3)` is called (delete columns C, D, E)
- THEN columns C, D, E are deleted
- AND columns F+ shift left by 3
- AND all cell references are updated accordingly

#### Scenario: Delete columns with formula references
- GIVEN a formula `=SUM(A1:J1)` in cell K1
- WHEN `sheet.DeleteColumns(8, 3)` is called (delete columns H, I, J)
- THEN the formula becomes `=SUM(A1:G1)`
- AND the range shrinks to exclude deleted columns

#### Scenario: Delete columns containing formula references
- GIVEN a formula `=C1+D1+E1` in cell F1
- WHEN `sheet.DeleteColumns(3, 3)` is called (delete columns C, D, E)
- THEN the formula becomes `=#REF!+#REF!+#REF!`
- AND references to deleted cells become #REF! errors

#### Scenario: Delete columns updates named ranges
- GIVEN a named range `QuarterData = Sheet1!A1:H100`
- WHEN `sheet.DeleteColumns(6, 3)` is called
- THEN the named range becomes `Sheet1!A1:E100`

#### Scenario: Delete columns updates merged cells
- GIVEN merged cells `E5:G5`
- WHEN `sheet.DeleteColumns(3, 2)` is called
- THEN the merged cells range becomes `C5:E5`

#### Scenario: Delete columns validation
- GIVEN a worksheet
- WHEN `sheet.DeleteColumns(0, 1)` is called
- THEN an error is returned (index must be >= 1)

### Requirement: Row Grouping Operations
The system SHALL provide methods to group and ungroup rows for hierarchical outlining.

#### Scenario: Group rows
- GIVEN a worksheet
- WHEN `sheet.GroupRows(10, 20)` is called
- THEN rows 10-20 have outlineLevel set to 1
- AND Excel displays expand/collapse controls for the group

#### Scenario: Nested row groups
- GIVEN rows 10-20 already grouped (outlineLevel = 1)
- WHEN `sheet.GroupRows(12, 15)` is called
- THEN rows 12-15 have outlineLevel set to 2
- AND a nested group is created

#### Scenario: Ungroup rows
- GIVEN rows 10-20 with outlineLevel = 1
- WHEN `sheet.UngroupRows(10, 20)` is called
- THEN rows 10-20 have outlineLevel set to 0
- AND the group is removed

#### Scenario: Maximum group level
- GIVEN rows with outlineLevel = 7 (maximum)
- WHEN `sheet.GroupRows(10, 20)` is called
- THEN an error is returned (max level is 7)

#### Scenario: Group validation
- GIVEN a worksheet
- WHEN `sheet.GroupRows(20, 10)` is called (start > end)
- THEN an error is returned (start must be <= end)

### Requirement: Column Grouping Operations
The system SHALL provide methods to group and ungroup columns for hierarchical outlining.

#### Scenario: Group columns
- GIVEN a worksheet
- WHEN `sheet.GroupColumns(5, 10)` is called
- THEN columns 5-10 (E-J) have outlineLevel set to 1
- AND Excel displays expand/collapse controls for the group

#### Scenario: Nested column groups
- GIVEN columns 5-10 already grouped (outlineLevel = 1)
- WHEN `sheet.GroupColumns(6, 8)` is called
- THEN columns 6-8 have outlineLevel set to 2
- AND a nested group is created

#### Scenario: Ungroup columns
- GIVEN columns 5-10 with outlineLevel = 1
- WHEN `sheet.UngroupColumns(5, 10)` is called
- THEN columns 5-10 have outlineLevel set to 0
- AND the group is removed

#### Scenario: Maximum group level
- GIVEN columns with outlineLevel = 7 (maximum)
- WHEN `sheet.GroupColumns(5, 10)` is called
- THEN an error is returned (max level is 7)

#### Scenario: Group validation
- GIVEN a worksheet
- WHEN `sheet.GroupColumns(10, 5)` is called (start > end)
- THEN an error is returned (start must be <= end)

### Requirement: Cross-Sheet Reference Updates
The system SHALL update cross-sheet references when rows or columns are inserted or deleted.

#### Scenario: Cross-sheet formula update on insert
- GIVEN a formula `=Sheet1!A5+Sheet2!B10` in Sheet2
- WHEN `sheet.InsertRows(5, 2)` is called on Sheet1
- THEN the formula becomes `=Sheet1!A7+Sheet2!B10`
- AND only the Sheet1 reference is updated

#### Scenario: Cross-sheet formula update on delete
- GIVEN a formula `=Sheet1!A10+Sheet2!A10` in Sheet3
- WHEN `sheet.DeleteRows(5, 3)` is called on Sheet1
- THEN the formula becomes `=Sheet1!A7+Sheet2!A10`
- AND only the Sheet1 reference is updated

#### Scenario: Cross-sheet named range update
- GIVEN a named range `Data = Sheet1!A1:D10`
- WHEN `sheet.InsertRows(5, 3)` is called on Sheet1
- THEN the named range becomes `Data = Sheet1!A1:D13`

### Requirement: Conditional Formatting Range Updates
The system SHALL update conditional formatting ranges when rows or columns are inserted or deleted.

#### Scenario: Conditional formatting range insert
- GIVEN conditional formatting applied to range A1:D10
- WHEN `sheet.InsertRows(5, 3)` is called
- THEN the conditional formatting range becomes A1:D13

#### Scenario: Conditional formatting range delete
- GIVEN conditional formatting applied to range A1:D10
- WHEN `sheet.DeleteRows(8, 3)` is called
- THEN the conditional formatting range becomes A1:D7

#### Scenario: Conditional formatting formula update
- GIVEN conditional formatting with formula `=$A1>100`
- WHEN `sheet.InsertRows(1, 2)` is called
- THEN the formula becomes `=$A3>100` (relative row shifted)

### Requirement: Data Validation Range Updates
The system SHALL update data validation ranges when rows or columns are inserted or deleted.

#### Scenario: Data validation range insert
- GIVEN data validation applied to range A1:A10
- WHEN `sheet.InsertRows(5, 3)` is called
- THEN the data validation range becomes A1:A13

#### Scenario: Data validation range delete
- GIVEN data validation applied to range A1:A10
- WHEN `sheet.DeleteRows(8, 3)` is called
- THEN the data validation range becomes A1:A7

#### Scenario: Data validation formula update
- GIVEN data validation with formula `=Sheet2!$A$1:$A$100`
- WHEN `sheet.InsertRows(50, 10)` is called on Sheet2
- THEN the formula becomes `=Sheet2!$A$1:$A$110` (absolute range extended)

### Requirement: Table Range Updates
The system SHALL update table ranges when rows or columns are inserted or deleted within the table.

#### Scenario: Table expands on row insert
- GIVEN a table with range A1:D10
- WHEN `sheet.InsertRows(5, 3)` is called (within table)
- THEN the table range becomes A1:D13

#### Scenario: Table shrinks on row delete
- GIVEN a table with range A1:D10
- WHEN `sheet.DeleteRows(8, 2)` is called (within table)
- THEN the table range becomes A1:D8

#### Scenario: Table shifts on column insert
- GIVEN a table with range C1:F10
- WHEN `sheet.InsertColumns(2, 2)` is called (before table)
- THEN the table range becomes E1:H10

#### Scenario: Table autoFilter range updates
- GIVEN a table with autoFilter on range A1:D10
- WHEN `sheet.InsertRows(5, 3)` is called
- THEN the autoFilter range becomes A1:D13
