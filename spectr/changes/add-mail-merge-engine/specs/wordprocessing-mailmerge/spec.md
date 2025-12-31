# Wordprocessing Mail Merge Specification

## Requirements

### Requirement: DataSource Interface
The system SHALL provide a DataSource interface for accessing merge data.

#### Scenario: Open data source
- GIVEN a DataSource implementation
- WHEN Open() is called
- THEN the data source is initialized and ready for iteration

#### Scenario: Iterate records
- GIVEN an open DataSource
- WHEN Next() is called repeatedly
- THEN it advances through records and returns true until exhausted, then returns false

#### Scenario: Access field names
- GIVEN an open DataSource
- WHEN Fields() is called
- THEN all available field names are returned as a slice

#### Scenario: Get field value
- GIVEN an open DataSource positioned at a record
- WHEN Get(fieldName) is called
- THEN the value of that field for the current record is returned

#### Scenario: Case-insensitive field access
- GIVEN an open DataSource with field "FirstName"
- WHEN Get("firstname"), Get("FIRSTNAME"), or Get("FirstName") is called
- THEN the same field value is returned

#### Scenario: Missing field handling
- GIVEN an open DataSource without field "MiddleName"
- WHEN Get("MiddleName") is called
- THEN an error is returned indicating field not found

#### Scenario: Close data source
- GIVEN an open DataSource
- WHEN Close() is called
- THEN all resources are released and the source is no longer usable

#### Scenario: Safe multiple close
- GIVEN a DataSource that has been closed
- WHEN Close() is called again
- THEN no error occurs (idempotent operation)

### Requirement: CSVDataSource Implementation
The system SHALL provide a CSVDataSource implementation for CSV files.

#### Scenario: Open CSV file
- GIVEN a path to a valid CSV file with headers
- WHEN NewCSVDataSource(path) and Open() are called
- THEN the CSV is opened and headers are read as field names

#### Scenario: Parse CSV headers
- GIVEN a CSV file with first row "FirstName,LastName,Email"
- WHEN Open() is called
- THEN Fields() returns ["FirstName", "LastName", "Email"]

#### Scenario: Iterate CSV rows
- GIVEN a CSV file with 3 data rows
- WHEN Next() is called repeatedly
- THEN it returns true 3 times, then false

#### Scenario: Access CSV values
- GIVEN a CSV row "John,Doe,john@example.com"
- WHEN Get("FirstName") is called
- THEN "John" is returned

#### Scenario: Handle missing CSV columns
- GIVEN a CSV row with fewer columns than headers
- WHEN Get() is called for a missing column
- THEN empty string is returned (not error)

#### Scenario: Trim CSV whitespace
- GIVEN a CSV file with "  FirstName  " header and "  John  " value
- WHEN Fields() and Get() are called
- THEN "FirstName" and "John" are returned (trimmed)

#### Scenario: Handle CSV parse errors
- GIVEN a malformed CSV file
- WHEN Open() or Next() is called
- THEN an error is returned

### Requirement: JSONDataSource Implementation
The system SHALL provide a JSONDataSource implementation for JSON array files.

#### Scenario: Open JSON array file
- GIVEN a path to a valid JSON array file
- WHEN NewJSONDataSource(path) and Open() are called
- THEN the JSON is parsed into memory

#### Scenario: Parse JSON fields
- GIVEN a JSON array [{"FirstName": "John", "LastName": "Doe"}]
- WHEN Fields() is called
- THEN ["FirstName", "LastName"] is returned (keys from first object)

#### Scenario: Iterate JSON records
- GIVEN a JSON array with 3 objects
- WHEN Next() is called repeatedly
- THEN it returns true 3 times, then false

#### Scenario: Access JSON values
- GIVEN a JSON object {"FirstName": "John", "Age": 30}
- WHEN Get("FirstName") is called
- THEN "John" is returned

#### Scenario: Convert non-string values
- GIVEN a JSON object {"Age": 30, "Active": true}
- WHEN Get("Age") is called
- THEN "30" is returned (converted to string)

#### Scenario: Handle empty JSON array
- GIVEN a JSON file "[]"
- WHEN Open() is called
- THEN it succeeds, but Next() returns false

#### Scenario: Handle non-array JSON
- GIVEN a JSON file "{}" (object, not array)
- WHEN Open() is called
- THEN an error is returned

### Requirement: MapDataSource Implementation
The system SHALL provide a MapDataSource implementation for in-memory data.

#### Scenario: Create from map slice
- GIVEN a slice of maps []map[string]string
- WHEN NewMapDataSource(data) is called
- THEN a DataSource is created with the data in memory

#### Scenario: No file I/O
- GIVEN a MapDataSource
- WHEN Open() or Close() is called
- THEN they succeed immediately (no file operations)

#### Scenario: Collect heterogeneous fields
- GIVEN maps with different keys: [{"A": "1"}, {"B": "2"}]
- WHEN Fields() is called
- THEN ["A", "B"] is returned (union of all keys)

#### Scenario: Iterate map records
- GIVEN a MapDataSource with 3 maps
- WHEN Next() is called repeatedly
- THEN it returns true 3 times, then false

#### Scenario: Access map values
- GIVEN a current record {"FirstName": "John"}
- WHEN Get("FirstName") is called
- THEN "John" is returned

### Requirement: MailMerge Engine
The system SHALL provide a MailMerge type for executing mail merge operations.

#### Scenario: Create mail merge
- GIVEN a WordprocessingDocument template
- WHEN New(template) is called
- THEN a MailMerge instance is created

#### Scenario: Set data source
- GIVEN a MailMerge instance
- WHEN DataSource(csvSource) is called
- THEN the data source is attached and the MailMerge is returned (fluent)

#### Scenario: Set options
- GIVEN a MailMerge instance
- WHEN Options(opts) is called
- THEN merge options are set and the MailMerge is returned (fluent)

#### Scenario: Fluent chaining
- GIVEN a MailMerge instance
- WHEN New(template).DataSource(ds).Options(opts).Execute() is chained
- THEN all methods execute in sequence

### Requirement: MERGEFIELD Detection (SimpleField)
The system SHALL detect and parse SimpleField (w:fldSimple) merge fields.

#### Scenario: Find SimpleField merge field
- GIVEN a template with `<w:fldSimple w:instr="MERGEFIELD FirstName"/>`
- WHEN findMergeFields() is called
- THEN a MergeField with FieldName="FirstName" and FieldType=Simple is returned

#### Scenario: Find SimpleField with switches
- GIVEN a template with `<w:fldSimple w:instr="MERGEFIELD Address \\* MERGEFORMAT"/>`
- WHEN findMergeFields() is called
- THEN a MergeField with FieldName="Address" and Switches={"*": "MERGEFORMAT"} is returned

#### Scenario: Find SimpleField in table
- GIVEN a table cell containing a SimpleField
- WHEN findMergeFields() is called
- THEN the SimpleField is detected

#### Scenario: Find SimpleField in header
- GIVEN a header containing a SimpleField
- WHEN findMergeFields() is called
- THEN the SimpleField is detected

#### Scenario: Find SimpleField in footer
- GIVEN a footer containing a SimpleField
- WHEN findMergeFields() is called
- THEN the SimpleField is detected

### Requirement: MERGEFIELD Detection (ComplexField)
The system SHALL detect and parse ComplexField (w:fldChar + w:instrText) merge fields.

#### Scenario: Find ComplexField merge field
- GIVEN a template with ComplexField structure (begin + instrText + separate + result + end)
- WHEN findMergeFields() is called
- THEN a MergeField with FieldName extracted and FieldType=Complex is returned

#### Scenario: Reconstruct field code from multiple runs
- GIVEN a ComplexField with instrText split across 3 runs
- WHEN findMergeFields() is called
- THEN the field code is correctly reconstructed from all instrText runs

#### Scenario: Handle ComplexField without separate marker
- GIVEN a ComplexField with begin + instrText + end (no separate or result)
- WHEN findMergeFields() is called
- THEN the field is detected correctly

#### Scenario: Find ComplexField with switches
- GIVEN a ComplexField with "MERGEFIELD Name \\* UPPER"
- WHEN findMergeFields() is called
- THEN Switches={"*": "UPPER"} is returned

#### Scenario: Find multiple merge fields (mixed types)
- GIVEN a template with 2 SimpleField and 2 ComplexField merge fields
- WHEN findMergeFields() is called
- THEN 4 MergeField instances are returned with correct types

#### Scenario: Handle no merge fields
- GIVEN a template with no MERGEFIELD codes
- WHEN findMergeFields() is called
- THEN an empty slice is returned (not error)

#### Scenario: Parse field code with switches
- GIVEN field code text "MERGEFIELD CustomerName \\* MERGEFORMAT \\b \"prefix\""
- WHEN parseMergeFieldCode() is called
- THEN "CustomerName" is extracted with Switches={"*": "MERGEFORMAT", "b": "\"prefix\""}

#### Scenario: Handle malformed field code
- GIVEN field code text "MERGE FIELD" (invalid syntax)
- WHEN parseMergeFieldCode() is called
- THEN empty field name is returned (field ignored)

### Requirement: Field Replacement (SimpleField)
The system SHALL replace SimpleField elements with data values.

#### Scenario: Replace SimpleField with value
- GIVEN a SimpleField "MERGEFIELD FirstName" and data value "John"
- WHEN replaceSimpleField() is called
- THEN the entire w:fldSimple element is replaced with a Run containing "John"

#### Scenario: Preserve SimpleField formatting
- GIVEN a SimpleField with child run having bold, italic formatting
- WHEN field is replaced with text
- THEN formatting is copied to replacement run

#### Scenario: Replace SimpleField in parent
- GIVEN a SimpleField inside a paragraph
- WHEN replaceSimpleField() is called
- THEN the SimpleField is removed from paragraph and Run is inserted at same index

### Requirement: Field Replacement (ComplexField)
The system SHALL replace ComplexField spans with data values.

#### Scenario: Replace ComplexField with value
- GIVEN a ComplexField "MERGEFIELD LastName" and data value "Smith"
- WHEN replaceComplexField() is called
- THEN all runs from begin to end are removed and replaced with a Run containing "Smith"

#### Scenario: Preserve ComplexField result formatting
- GIVEN a ComplexField with result runs having font and color formatting
- WHEN field is replaced with text
- THEN formatting from result run is copied to replacement run

#### Scenario: Replace ComplexField without separate
- GIVEN a ComplexField with begin + instrText + end (no separate/result)
- WHEN replaceComplexField() is called
- THEN runs are correctly replaced even without result formatting

#### Scenario: Replace all fields in document (mixed types)
- GIVEN a document with 2 SimpleField and 2 ComplexField merge fields
- WHEN replaceFields() is called
- THEN all 4 fields are replaced with data values

#### Scenario: Handle missing field (default)
- GIVEN a template field "MiddleName" not in data source
- WHEN replaceFields() is called with default options
- THEN the field is replaced with empty string

#### Scenario: Handle missing field (strict mode)
- GIVEN a template field "MiddleName" not in data source
- WHEN replaceFields() is called with StrictFields=true
- THEN an error is returned

#### Scenario: Apply format switch
- GIVEN a field with "\\* UPPER" switch and data value "john"
- WHEN field is replaced
- THEN the value is formatted as "JOHN"

### Requirement: Execute Single Document
The system SHALL merge all records into a single output document.

#### Scenario: Execute with multiple records
- GIVEN a template and data source with 3 records
- WHEN Execute() is called
- THEN a single document is returned with all 3 records merged

#### Scenario: Page breaks between records
- GIVEN Execute() with 3 records
- WHEN the output document is examined
- THEN page breaks exist between each record

#### Scenario: Clone template per record
- GIVEN a template with formatting and content
- WHEN Execute() is called
- THEN each record has the same formatting as the template

#### Scenario: Template not modified
- GIVEN a template document
- WHEN Execute() is called
- THEN the template document is unchanged

#### Scenario: Handle zero records
- GIVEN a data source with zero records
- WHEN Execute() is called
- THEN an error is returned

#### Scenario: Close data source after execution
- GIVEN an Execute() operation
- WHEN execution completes (success or error)
- THEN the data source is closed

### Requirement: Execute Multiple Documents
The system SHALL merge each record into a separate output document.

#### Scenario: Execute to multiple documents
- GIVEN a template and data source with 3 records
- WHEN ExecuteToDocuments() is called
- THEN 3 separate WordprocessingDocument instances are returned

#### Scenario: One document per record
- GIVEN ExecuteToDocuments() with 3 records
- WHEN the output slice is examined
- THEN it contains exactly 3 documents

#### Scenario: Each document is independent
- GIVEN output documents from ExecuteToDocuments()
- WHEN one document is modified
- THEN other documents are not affected

#### Scenario: Each document has one record
- GIVEN ExecuteToDocuments() with 3 records
- WHEN each output document is examined
- THEN it contains content for only one record

#### Scenario: No page breaks within documents
- GIVEN output from ExecuteToDocuments()
- WHEN a document is examined
- THEN it has no page breaks (only one record per document)

#### Scenario: Template not modified
- GIVEN a template document
- WHEN ExecuteToDocuments() is called
- THEN the template document is unchanged

#### Scenario: Handle zero records
- GIVEN a data source with zero records
- WHEN ExecuteToDocuments() is called
- THEN an error is returned

### Requirement: Field Validation
The system SHALL validate that data sources provide required fields.

#### Scenario: Validate all fields present
- GIVEN a template with fields ["FirstName", "LastName"]
- WHEN validateFields() is called with data source having those fields
- THEN validation passes

#### Scenario: Validate field missing (strict mode)
- GIVEN a template with field "MiddleName"
- WHEN validateFields() is called in strict mode without that field
- THEN an error is returned

#### Scenario: Skip validation (non-strict mode)
- GIVEN StrictFields=false
- WHEN validateFields() is called
- THEN validation is skipped (always passes)

#### Scenario: Case-insensitive validation
- GIVEN a template with field "FirstName"
- WHEN validateFields() is called with data source field "firstname"
- THEN validation passes (case-insensitive match)

### Requirement: Merge Options
The system SHALL support configurable merge behavior through MergeOptions.

#### Scenario: Default options
- GIVEN DefaultMergeOptions()
- THEN StrictFields=false and RemoveUnusedFields=false

#### Scenario: Strict fields option
- GIVEN MergeOptions with StrictFields=true
- WHEN merge encounters missing field
- THEN an error is returned

#### Scenario: Non-strict fields option
- GIVEN MergeOptions with StrictFields=false
- WHEN merge encounters missing field
- THEN empty string is used (no error)

#### Scenario: Remove unused fields option
- GIVEN MergeOptions with RemoveUnusedFields=true
- WHEN merge encounters field not in data source
- THEN the field code is removed from output

#### Scenario: Keep unused fields option
- GIVEN MergeOptions with RemoveUnusedFields=false
- WHEN merge encounters field not in data source
- THEN the field code remains in output

### Requirement: Error Handling
The system SHALL provide clear error messages for all failure scenarios.

#### Scenario: Data source not set
- GIVEN a MailMerge without DataSource called
- WHEN Execute() is called
- THEN an error "data source not set" is returned

#### Scenario: Data source open failure
- GIVEN a CSVDataSource with invalid file path
- WHEN Execute() is called
- THEN an error indicating file open failure is returned

#### Scenario: Field finding failure
- GIVEN a template with corrupted XML
- WHEN Execute() is called
- THEN an error indicating field finding failure is returned

#### Scenario: Field validation failure
- GIVEN strict mode with missing field
- WHEN Execute() is called
- THEN an error indicating which field is missing is returned

#### Scenario: Template clone failure
- GIVEN a template that cannot be cloned
- WHEN Execute() is called
- THEN an error indicating clone failure is returned

#### Scenario: Field replacement failure
- GIVEN a field that cannot be replaced
- WHEN Execute() is called
- THEN an error indicating which record/field failed is returned

### Requirement: Special Character Handling
The system SHALL correctly handle special characters in merge data.

#### Scenario: XML special characters
- GIVEN data values containing &, <, >, ", '
- WHEN merge is executed
- THEN characters are properly XML-escaped in output

#### Scenario: Unicode characters
- GIVEN data values with Unicode (emoji, accents, etc.)
- WHEN merge is executed
- THEN Unicode is preserved in output

#### Scenario: Newlines in data
- GIVEN data values containing newline characters
- WHEN merge is executed
- THEN newlines are preserved (creating line breaks in output)

#### Scenario: Empty string values
- GIVEN data values that are empty strings
- WHEN merge is executed
- THEN fields are replaced with empty content (not error)

### Requirement: Document Compatibility
The system SHALL produce merged documents compatible with Microsoft Word and LibreOffice.

#### Scenario: Open in Microsoft Word
- GIVEN a merged document
- WHEN opened in Microsoft Word
- THEN it opens without errors or corruption warnings

#### Scenario: Open in LibreOffice Writer
- GIVEN a merged document
- WHEN opened in LibreOffice Writer
- THEN it opens and displays correctly

#### Scenario: Roundtrip test
- GIVEN a merged document
- WHEN saved, closed, and reopened
- THEN content is identical to before save

#### Scenario: Validation passes
- GIVEN a merged document
- WHEN validated against OOXML schema
- THEN validation passes (document is well-formed)

### Requirement: OOXML Metadata Integration
The system SHALL read and apply mail merge metadata from settings.xml.

#### Scenario: Load mail merge metadata
- GIVEN a template with w:mailMerge element in settings.xml
- WHEN loadMailMergeMetadata() is called
- THEN DataType, ConnectString, Query are extracted

#### Scenario: Load ODSO field mappings
- GIVEN settings.xml with w:odso containing fieldMapData elements
- WHEN loadMailMergeMetadata() is called
- THEN FieldMappings map is populated with template→data column mappings

#### Scenario: Apply field mappings during merge
- GIVEN metadata with mapping "FullName" → "CustomerName"
- WHEN merge executes with field "FullName"
- THEN data is retrieved using column "CustomerName"

#### Scenario: Handle template without metadata
- GIVEN a template without settings.xml or w:mailMerge
- WHEN loadMailMergeMetadata() is called
- THEN nil is returned (fallback to direct field matching)

#### Scenario: Update settings.xml after merge
- GIVEN UpdateSettings option is true
- WHEN Execute() completes
- THEN settings.xml is updated with data source type

### Requirement: GREETINGLINE Field
The system SHALL support GREETINGLINE merge field for generating greetings.

#### Scenario: Generate default greeting
- GIVEN GREETINGLINE field and data with Title="Mr." and LastName="Smith"
- WHEN generateGreetingLine() is called
- THEN "Dear Mr. Smith," is generated

#### Scenario: Generate greeting with FirstName + LastName
- GIVEN GREETINGLINE and data with FirstName="John" and LastName="Doe"
- WHEN generateGreetingLine() is called
- THEN "Dear John Doe," is generated

#### Scenario: Handle missing title
- GIVEN GREETINGLINE and data with only FirstName and LastName
- WHEN generateGreetingLine() is called
- THEN "Dear FirstName LastName," is generated

#### Scenario: Custom greeting format
- GIVEN GREETINGLINE with \\f switch "Hello {FirstName}!"
- WHEN generateGreetingLine() is called
- THEN custom format is used

#### Scenario: Fallback greeting
- GIVEN GREETINGLINE and data with no name fields
- WHEN generateGreetingLine() is called
- THEN "Dear Recipient," is generated

### Requirement: SKIPIF Field (Basic Conditional)
The system SHALL support SKIPIF field for conditionally skipping records.

#### Scenario: Skip record with equality condition
- GIVEN SKIPIF field "SKIPIF Status = \"Inactive\""
- WHEN current record has Status="Inactive"
- THEN record is skipped (not included in output)

#### Scenario: Include record when condition false
- GIVEN SKIPIF field "SKIPIF Status = \"Inactive\""
- WHEN current record has Status="Active"
- THEN record is included in output

#### Scenario: Skip with inequality condition
- GIVEN SKIPIF field "SKIPIF Age <> \"18\""
- WHEN current record has Age="18"
- THEN record is included (condition false)

#### Scenario: Skip with numeric comparison
- GIVEN SKIPIF field "SKIPIF Quantity < \"10\""
- WHEN current record has Quantity="5"
- THEN record is skipped

#### Scenario: SKIPIF in Execute mode
- GIVEN Execute() with 5 records, 2 match SKIPIF condition
- WHEN merge executes
- THEN output document contains 3 records

#### Scenario: SKIPIF in ExecuteToDocuments mode
- GIVEN ExecuteToDocuments() with 5 records, 2 match SKIPIF
- WHEN merge executes
- THEN 3 documents are returned
