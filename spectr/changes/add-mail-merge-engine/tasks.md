# Implementation Tasks

## Phase 1: Core DataSource Infrastructure (1 week)

### 1.1 Define DataSource Interface
- [ ] Create `wordprocessing/mailmerge/data_source.go`
- [ ] Define DataSource interface with Open/Close/Next/Fields/Get methods
- [ ] Add comprehensive godoc comments
- [ ] Include usage examples in package documentation
- [ ] Validation: Interface compiles, godoc looks professional

### 1.2 Implement CSVDataSource
- [ ] Implement CSVDataSource struct in `data_source.go`
- [ ] Implement Open method (open file, read headers)
- [ ] Implement Next method (iterate CSV rows)
- [ ] Implement Fields method (return headers)
- [ ] Implement Get method (retrieve field value, case-insensitive)
- [ ] Implement Close method (close file handle)
- [ ] Add error handling for malformed CSV
- [ ] Trim whitespace from headers and values
- [ ] Validation: CSVDataSource compiles and satisfies DataSource interface

### 1.3 Implement MapDataSource
- [ ] Implement MapDataSource struct in `data_source.go`
- [ ] Implement Open method (no-op for in-memory)
- [ ] Implement Next method (iterate slice)
- [ ] Implement Fields method (collect unique keys)
- [ ] Implement Get method (lookup in current map, case-insensitive)
- [ ] Implement Close method (no-op)
- [ ] Validation: MapDataSource compiles and satisfies DataSource interface

### 1.4 Write DataSource Unit Tests
- [ ] Create `wordprocessing/mailmerge/data_source_test.go`
- [ ] Test CSVDataSource with valid CSV file
- [ ] Test CSVDataSource with headers-only CSV (zero records)
- [ ] Test CSVDataSource with empty file (error case)
- [ ] Test CSVDataSource case-insensitive field lookup
- [ ] Test CSVDataSource with missing columns in row
- [ ] Test MapDataSource with multiple records
- [ ] Test MapDataSource with empty slice (zero records)
- [ ] Test MapDataSource case-insensitive field lookup
- [ ] Test MapDataSource with heterogeneous keys across records
- [ ] Validation: All tests pass, >90% coverage for DataSource implementations

## Phase 2: Field Finding & Parsing (1 week)

### 2.1 Create MailMerge Type
- [ ] Create `wordprocessing/mailmerge/engine.go`
- [ ] Define MailMerge struct with template, dataSource, options, metadata fields
- [ ] Define MergeOptions struct (StrictFields, RemoveUnusedFields, UpdateSettings)
- [ ] Implement New(template) constructor
- [ ] Implement DataSource() fluent method
- [ ] Implement Options() fluent method
- [ ] Add DefaultMergeOptions() function
- [ ] Validation: MailMerge type compiles, API is fluent

### 2.2 Define Field Types and Structures
- [ ] Create `wordprocessing/mailmerge/field_types.go`
- [ ] Define FieldType enum (FieldTypeSimple, FieldTypeComplex)
- [ ] Define MergeField struct with FieldName, FieldType, Element, BeginRun, EndRun, FieldCode, Switches
- [ ] Add comprehensive godoc comments explaining SimpleField vs ComplexField
- [ ] Validation: Field type definitions compile

### 2.3 Implement SimpleField Detection
- [ ] Create `wordprocessing/mailmerge/field_finder.go`
- [ ] Implement findMergeFields() method (entry point)
- [ ] Implement findFieldsInElement() recursive walker
- [ ] Detect SimpleField elements (w:fldSimple)
- [ ] Extract w:instr attribute from SimpleField
- [ ] Parse SimpleField instruction text
- [ ] Search in body, headers, footers, tables
- [ ] Validation: SimpleField detection finds all w:fldSimple elements

### 2.4 Implement ComplexField Detection
- [ ] Implement findComplexFieldsInParagraph() method
- [ ] Detect FieldChar begin markers (w:fldChar fldCharType="begin")
- [ ] Implement reconstructFieldCode() to reassemble instrText spans
- [ ] Handle instrText split across multiple runs
- [ ] Detect separate markers (w:fldChar fldCharType="separate")
- [ ] Detect end markers (w:fldChar fldCharType="end")
- [ ] Implement hasFieldChar() helper method
- [ ] Handle ComplexField with no separate marker (no result)
- [ ] Validation: ComplexField detection reconstructs field codes correctly

### 2.5 Implement Field Code Parsing
- [ ] Implement parseMergeFieldCode() method (returns fieldName and switches)
- [ ] Implement tokenizeFieldCode() to handle quoted strings
- [ ] Parse MERGEFIELD syntax ("MERGEFIELD FieldName \\switches")
- [ ] Parse GREETINGLINE syntax ("GREETINGLINE \\switches")
- [ ] Extract field switches (\\*, \\b, \\f, etc.)
- [ ] Handle quoted switch values
- [ ] Return map of switches
- [ ] Validation: Field code parsing handles all switch types correctly

### 2.6 Write Field Finding Tests
- [ ] Create `wordprocessing/mailmerge/field_finder_test.go`
- [ ] Test findMergeFields with SimpleField elements
- [ ] Test findMergeFields with ComplexField spans
- [ ] Test field finding in paragraphs, tables, headers, footers
- [ ] Test ComplexField with multiple instrText runs
- [ ] Test ComplexField with separate marker
- [ ] Test ComplexField without separate marker
- [ ] Test field code parsing with switches
- [ ] Test tokenizeFieldCode with quoted strings
- [ ] Validation: All field finding tests pass, >90% coverage

## Phase 3: Field Replacement (1 week)

### 3.1 Implement Field Replacement Core
- [ ] Create `wordprocessing/mailmerge/field_replacer.go`
- [ ] Implement replaceFields() method (entry point)
- [ ] Implement getFieldValue() method (handles MERGEFIELD and special fields)
- [ ] Implement applyFormatSwitch() for formatting switches (\\* Upper, \\* Lower, etc.)
- [ ] Handle missing fields (empty string or error based on StrictFields)
- [ ] Validation: Field replacement core logic compiles

### 3.2 Implement SimpleField Replacement
- [ ] Implement replaceSimpleField() method
- [ ] Get parent element of SimpleField
- [ ] Create replacement Run with text value
- [ ] Preserve formatting from SimpleField's child runs
- [ ] Find index of SimpleField in parent
- [ ] Remove SimpleField element from parent
- [ ] Insert replacement Run at same index
- [ ] Validation: SimpleField replacement works in tests

### 3.3 Implement ComplexField Replacement
- [ ] Implement replaceComplexField() method
- [ ] Get paragraph containing BeginRun and EndRun
- [ ] Find indices of begin and end markers
- [ ] Create replacement Run with text value
- [ ] Preserve formatting from result runs (between separate and end)
- [ ] Remove all runs from begin to end (inclusive)
- [ ] Insert replacement Run at begin index
- [ ] Handle ComplexField with no separate marker
- [ ] Validation: ComplexField replacement preserves formatting

### 3.4 Implement Field Validation
- [ ] Update validateFields() method to handle both field types
- [ ] Skip validation for special fields (GREETINGLINE)
- [ ] Case-insensitive field name matching
- [ ] Return clear error messages for missing fields
- [ ] Validation: Field validation catches all missing fields

### 3.5 Implement Execute Method
- [ ] Update Execute() method to use new field finding/replacement
- [ ] Load mail merge metadata from settings.xml
- [ ] Apply field mappings if metadata exists
- [ ] Find all merge fields (SimpleField and ComplexField)
- [ ] Validate fields against data source
- [ ] Clone template document
- [ ] Iterate records and replace fields
- [ ] Insert page breaks between records
- [ ] Handle zero records case (error)
- [ ] Validation: Execute produces single document with all records

### 3.6 Write Field Replacement Tests
- [ ] Create `wordprocessing/mailmerge/field_replacer_test.go`
- [ ] Test replaceSimpleField with plain text
- [ ] Test replaceSimpleField preserving formatting
- [ ] Test replaceComplexField with result text
- [ ] Test replaceComplexField preserving formatting
- [ ] Test getFieldValue with MERGEFIELD
- [ ] Test applyFormatSwitch (Upper, Lower, FirstCap, Caps)
- [ ] Test Execute with SimpleField template
- [ ] Test Execute with ComplexField template
- [ ] Test Execute with mixed field types
- [ ] Validation: All replacement tests pass, >85% coverage

## Phase 4: OOXML Metadata Integration (3 days)

### 4.1 Implement Metadata Types
- [ ] Create `wordprocessing/mailmerge/metadata.go`
- [ ] Define MailMergeMetadata struct (DataType, ConnectString, Query, FieldMappings)
- [ ] Add godoc comments explaining OOXML mail merge metadata
- [ ] Validation: Metadata types compile

### 4.2 Implement Metadata Loading
- [ ] Implement loadMailMergeMetadata() method
- [ ] Access DocumentSettingsPart from template
- [ ] Read w:mailMerge element from settings.xml
- [ ] Extract DataType, ConnectString, Query properties
- [ ] Read ODSO (Open Data Source Object)
- [ ] Parse FieldMapData elements for field name mappings
- [ ] Return MailMergeMetadata or nil if no settings
- [ ] Validation: Metadata loading reads settings.xml correctly

### 4.3 Implement Field Mapping Application
- [ ] Implement applyFieldMappings() method
- [ ] Check if metadata contains mapping for field name
- [ ] Return mapped name if exists, otherwise return original
- [ ] Integrate into getFieldValue() method
- [ ] Case-insensitive mapping lookup
- [ ] Validation: Field mappings are applied during merge

### 4.4 Implement Metadata Update (Optional)
- [ ] Implement updateMailMergeMetadata() method
- [ ] Create DocumentSettingsPart if it doesn't exist
- [ ] Create w:mailMerge element if it doesn't exist
- [ ] Set DataType based on data source type (CSV, JSON, etc.)
- [ ] Only update if UpdateSettings option is true
- [ ] Validation: Metadata update creates proper settings.xml

### 4.5 Write Metadata Tests
- [ ] Create `wordprocessing/mailmerge/metadata_test.go`
- [ ] Test loadMailMergeMetadata with existing settings
- [ ] Test loadMailMergeMetadata with no settings (returns nil)
- [ ] Test loadMailMergeMetadata with ODSO field mappings
- [ ] Test applyFieldMappings with mapping present
- [ ] Test applyFieldMappings with no mapping
- [ ] Test updateMailMergeMetadata creates settings
- [ ] Test Execute with field mappings from metadata
- [ ] Validation: All metadata tests pass, >85% coverage

## Phase 5: Multiple Document Mode & Advanced Fields (1 week)

### 5.1 Implement ExecuteToDocuments Method
- [ ] Implement ExecuteToDocuments() method signature
- [ ] Validate data source is set
- [ ] Open data source
- [ ] Load metadata (if present)
- [ ] Find all merge fields
- [ ] Validate fields against data source
- [ ] Iterate data source records
- [ ] For each record: clone template document (full clone, not just body)
- [ ] Replace fields in cloned document
- [ ] Append cloned document to output slice
- [ ] Close data source
- [ ] Return slice of merged documents
- [ ] Handle zero records case (error)
- [ ] Validation: ExecuteToDocuments produces one document per record

### 5.2 Implement GREETINGLINE Field
- [ ] Create `wordprocessing/mailmerge/advanced_fields.go`
- [ ] Implement generateGreetingLine() method
- [ ] Read Title, FirstName, LastName from data source
- [ ] Parse \\f switch for custom format
- [ ] Generate default greeting format ("Dear [Title] [LastName],")
- [ ] Handle missing name fields gracefully
- [ ] Support common greeting variations
- [ ] Validation: GREETINGLINE generates proper greetings

### 5.3 Implement SKIPIF Field (Basic)
- [ ] Extend parseMergeFieldCode() to recognize SKIPIF
- [ ] Parse SKIPIF condition (field, operator, value)
- [ ] Implement evaluateSKIPIF() method
- [ ] Support basic comparisons (=, <>, <, >, <=, >=)
- [ ] Return true if record should be skipped
- [ ] Integrate into Execute and ExecuteToDocuments
- [ ] Skip record if SKIPIF condition is true
- [ ] Validation: SKIPIF conditionally skips records

### 5.4 Write ExecuteToDocuments Tests
- [ ] Create tests for ExecuteToDocuments with MapDataSource
- [ ] Test ExecuteToDocuments with multiple records
- [ ] Test ExecuteToDocuments returns correct number of documents
- [ ] Test each document has correct merged content
- [ ] Test documents are independent (modifying one doesn't affect others)
- [ ] Test ExecuteToDocuments with zero records (error)
- [ ] Validation: All tests pass

### 5.5 Write Advanced Field Tests
- [ ] Create `wordprocessing/mailmerge/advanced_fields_test.go`
- [ ] Test GREETINGLINE with full name (Title + LastName)
- [ ] Test GREETINGLINE with partial name (FirstName + LastName)
- [ ] Test GREETINGLINE with missing fields
- [ ] Test GREETINGLINE with custom format switch
- [ ] Test SKIPIF with equality condition (=)
- [ ] Test SKIPIF with inequality condition (<>)
- [ ] Test SKIPIF with numeric comparisons (<, >, <=, >=)
- [ ] Test Execute skipping records based on SKIPIF
- [ ] Validation: All advanced field tests pass, >80% coverage

## Phase 6: Advanced Data Sources (3 days)

### 6.1 Implement JSONDataSource
- [ ] Implement JSONDataSource struct in `data_source.go`
- [ ] Implement Open method (read file, parse JSON array)
- [ ] Implement Next method (iterate array)
- [ ] Implement Fields method (extract keys from first record)
- [ ] Implement Get method (lookup in current object, case-insensitive)
- [ ] Implement Close method (no-op)
- [ ] Handle non-array JSON (error)
- [ ] Handle empty array (zero records)
- [ ] Convert non-string values to strings (fmt.Sprintf)
- [ ] Validation: JSONDataSource compiles and satisfies interface

### 6.2 Write JSONDataSource Tests
- [ ] Create test with valid JSON array
- [ ] Test with array of objects with matching keys
- [ ] Test with heterogeneous objects (different keys per record)
- [ ] Test with empty array (zero records)
- [ ] Test with non-array JSON (error)
- [ ] Test case-insensitive field lookup
- [ ] Test non-string values (numbers, booleans, nulls)
- [ ] Validation: All tests pass, >90% coverage

### 6.3 Add Field Validation
- [ ] Implement validateFields() method in engine
- [ ] Get field names from data source
- [ ] Create case-insensitive lookup set
- [ ] Check all merge fields exist in data source
- [ ] Return error if field missing (when StrictFields enabled)
- [ ] Skip validation if StrictFields is false
- [ ] Validation: Validation catches missing fields in strict mode

### 6.4 Write Validation Tests
- [ ] Test validateFields with all fields present
- [ ] Test validateFields with missing field (strict mode: error)
- [ ] Test validateFields with missing field (non-strict: pass)
- [ ] Test validateFields with case-insensitive match
- [ ] Validation: All validation tests pass

## Phase 7: Integration & Examples (3 days)

### 7.1 Integrate with WordprocessingDocument
- [ ] Add MailMerge() method to `wordprocessing/document.go`
- [ ] Return mailmerge.New(d) from MailMerge()
- [ ] Update imports
- [ ] Validation: Method compiles and is accessible

### 7.2 Create CSV Merge Example
- [ ] Create `examples/mail-merge-csv/`
- [ ] Create template document with SimpleField merge fields
- [ ] Create template document with ComplexField merge fields
- [ ] Create sample CSV file (customers.csv)
- [ ] Create main.go that executes merge
- [ ] Save merged output to file
- [ ] Add README.md explaining both field types
- [ ] Validation: Example runs and produces correct output

### 7.3 Create JSON Merge Example
- [ ] Create `examples/mail-merge-json/`
- [ ] Create template document with GREETINGLINE field
- [ ] Create sample JSON file (employees.json)
- [ ] Create main.go that executes merge
- [ ] Save merged output
- [ ] Add README.md
- [ ] Validation: Example runs and produces correct output

### 7.4 Create Multi-Document Example
- [ ] Create `examples/mail-merge-documents/`
- [ ] Create template with both SimpleField and ComplexField
- [ ] Create data file with field mappings
- [ ] Create main.go using ExecuteToDocuments
- [ ] Save each document to separate file
- [ ] Add README.md
- [ ] Validation: Example produces multiple output files

### 7.5 Create Advanced Fields Example
- [ ] Create `examples/mail-merge-advanced/`
- [ ] Create template with GREETINGLINE and SKIPIF fields
- [ ] Create data file with varied records
- [ ] Create main.go demonstrating advanced fields
- [ ] Show SKIPIF conditionally skipping records
- [ ] Add README.md explaining advanced field types
- [ ] Validation: Example demonstrates all advanced features

### 7.6 Write Integration Tests
- [ ] Create `wordprocessing/mailmerge/integration_test.go`
- [ ] Test CSV merge end-to-end (template → CSV → merged doc)
- [ ] Test JSON merge end-to-end
- [ ] Test ExecuteToDocuments end-to-end
- [ ] Test with special characters in data (quotes, ampersands, Unicode)
- [ ] Test with missing fields (both strict and non-strict)
- [ ] Roundtrip test: merge → save → open → verify content
- [ ] Validation: All integration tests pass

### 7.7 Test SimpleField Merged Documents
- [ ] Open merged document (SimpleField template) in Microsoft Word
- [ ] Verify field replacements display correctly
- [ ] Verify no corruption errors
- [ ] Verify special characters render correctly
- [ ] Test same document in LibreOffice Writer
- [ ] Validation: Documents open without errors

### 7.8 Test ComplexField Merged Documents
- [ ] Open merged document (ComplexField template) in Microsoft Word
- [ ] Verify field replacements preserved formatting
- [ ] Verify no field remnants visible
- [ ] Verify special characters render correctly
- [ ] Test same document in LibreOffice Writer
- [ ] Validation: Documents open and formatting is preserved

### 7.9 Test Documents with OOXML Metadata
- [ ] Create template with settings.xml mail merge metadata
- [ ] Create template with ODSO field mappings
- [ ] Execute merge with field mappings
- [ ] Verify field mappings were applied correctly
- [ ] Open merged document in Microsoft Word
- [ ] Validation: Metadata integration works correctly

## Phase 8: Documentation (2 days)

### 8.1 Write API Documentation
- [ ] Add godoc comments to all exported types
- [ ] Add godoc comments to all exported methods
- [ ] Add package-level documentation to `mailmerge` package
- [ ] Include code examples in godoc comments
- [ ] Document DataSource interface with examples
- [ ] Document MergeOptions fields (including UpdateSettings)
- [ ] Document FieldType enum and MergeField struct
- [ ] Explain SimpleField vs ComplexField in package docs
- [ ] Run `go doc wordprocessing/mailmerge` and verify output
- [ ] Validation: godoc output is professional and complete

### 8.2 Create Usage Guide
- [ ] Create `wordprocessing/mailmerge/README.md`
- [ ] Introduction: What is mail merge?
- [ ] Quick Start: Simplest merge example
- [ ] Field Types: Explain SimpleField vs ComplexField
- [ ] DataSource: Explain interface and implementations
- [ ] Merge Modes: Execute vs ExecuteToDocuments
- [ ] Options: StrictFields, RemoveUnusedFields, UpdateSettings
- [ ] OOXML Metadata: Explain settings.xml and ODSO integration
- [ ] Advanced Fields: Document GREETINGLINE and SKIPIF
- [ ] Field Switches: Document supported formatting switches
- [ ] Custom DataSources: How to implement
- [ ] Troubleshooting: Common issues and solutions
- [ ] Validation: Guide is clear and helpful

### 8.3 Update Main README
- [ ] Add mail merge to feature list in `README.md`
- [ ] Add mail merge code example showing both field types
- [ ] Mention GREETINGLINE and SKIPIF support
- [ ] Link to mailmerge package documentation
- [ ] Link to examples
- [ ] Validation: README accurately reflects new capability

### 8.4 Create Migration Guide
- [ ] Create `wordprocessing/mailmerge/MIGRATION.md`
- [ ] Explain how to migrate from external tools
- [ ] Show equivalent of Word VBA mail merge
- [ ] Show equivalent of Python docx-mailmerge
- [ ] Explain converting Word templates (SimpleField vs ComplexField)
- [ ] Code comparison examples
- [ ] Validation: Guide helps users migrate

### 8.5 Write CHANGELOG Entry
- [ ] Add entry to `CHANGELOG.md`
- [ ] Describe new mail merge feature
- [ ] List all new types and methods
- [ ] Mention SimpleField and ComplexField support
- [ ] Mention GREETINGLINE and SKIPIF support
- [ ] Mention OOXML metadata integration
- [ ] Note any breaking changes (none expected)
- [ ] Validation: CHANGELOG is up to date

## Validation Checkpoints

### After Phase 1 (Core Infrastructure):
- [ ] DataSource interface is clean and extensible
- [ ] CSVDataSource can parse CSV files correctly
- [ ] MapDataSource supports in-memory data
- [ ] All data source tests pass

### After Phase 2 (Field Finding & Parsing):
- [ ] MailMerge type has fluent API
- [ ] SimpleField detection finds all w:fldSimple elements
- [ ] ComplexField detection reconstructs field codes from spans
- [ ] Field finding searches body, headers, footers, tables
- [ ] Field code parsing handles switches correctly
- [ ] All field finding tests pass

### After Phase 3 (Field Replacement):
- [ ] SimpleField replacement works correctly
- [ ] ComplexField replacement preserves formatting
- [ ] Format switches are applied (Upper, Lower, etc.)
- [ ] Execute() produces single merged document
- [ ] All replacement tests pass

### After Phase 4 (OOXML Metadata Integration):
- [ ] Metadata loading reads settings.xml correctly
- [ ] ODSO field mappings are parsed
- [ ] Field mappings are applied during merge
- [ ] Metadata update creates proper settings.xml
- [ ] All metadata tests pass

### After Phase 5 (Multiple Document Mode & Advanced Fields):
- [ ] ExecuteToDocuments() produces one document per record
- [ ] Documents are independent
- [ ] GREETINGLINE generates proper greetings
- [ ] SKIPIF conditionally skips records
- [ ] All advanced field tests pass

### After Phase 6 (Advanced Data Sources):
- [ ] JSONDataSource parses JSON arrays
- [ ] Field validation catches missing fields
- [ ] StrictFields option works correctly
- [ ] Tests cover all data source types

### After Phase 7 (Integration & Examples):
- [ ] MailMerge() accessible from WordprocessingDocument
- [ ] CSV example demonstrates both field types
- [ ] JSON example demonstrates GREETINGLINE
- [ ] Multi-document example works
- [ ] Advanced fields example demonstrates SKIPIF
- [ ] Integration tests pass
- [ ] SimpleField documents open in Word and LibreOffice
- [ ] ComplexField documents open in Word and LibreOffice
- [ ] Metadata integration works in Word

### After Phase 8 (Documentation):
- [ ] API documentation is complete
- [ ] Usage guide explains SimpleField vs ComplexField
- [ ] Usage guide explains OOXML metadata integration
- [ ] Usage guide explains advanced field types
- [ ] README is updated with new capabilities
- [ ] Migration guide assists users
- [ ] CHANGELOG is accurate

## Dependencies

**Blocked by**: None (all required APIs exist)

**Blocks**:
- Conditional merge fields (IF, COMPARE) - Phase 2 feature
- Nested merge regions - Phase 2 feature
- Formula fields in merge - Phase 2 feature

## Parallelizable Work

- Phase 1: Data source implementations can be developed in parallel (CSV, Map)
- Phase 4: JSONDataSource can be developed parallel to validation work
- Phase 5: Examples can be created in parallel
- Phase 6: Documentation tasks can be parallelized

## Risk Mitigation

- **Risk**: Complex field codes not parsed correctly
  - **Mitigation**: Start with simple "MERGEFIELD Name" syntax, add complex parsing later
  - **Test**: Create test cases with various field code formats

- **Risk**: Memory issues with large CSV files
  - **Mitigation**: Use streaming CSV parser (encoding/csv does this)
  - **Test**: Test with 10MB+ CSV files

- **Risk**: XML encoding issues with special characters
  - **Mitigation**: Use proper XML escaping in text insertion
  - **Test**: Include special characters in test data (&, <, >, ", ')

- **Risk**: Template cloning errors
  - **Mitigation**: Validate clone success before proceeding
  - **Test**: Test cloning with complex templates (tables, images, etc.)

- **Risk**: Word/LibreOffice compatibility
  - **Mitigation**: Test merged documents in both applications
  - **Test**: Open all test outputs in Word and LibreOffice

## Success Criteria

All tasks above completed AND:
- [ ] DataSource interface has 3+ implementations (CSV, JSON, Map)
- [ ] MailMerge engine supports Execute and ExecuteToDocuments modes
- [ ] SimpleField detection has 100% accuracy
- [ ] ComplexField detection properly reconstructs field codes from spans
- [ ] Field finding searches body, tables, headers, footers, text boxes
- [ ] SimpleField replacement correctly replaces w:fldSimple elements
- [ ] ComplexField replacement correctly handles begin/separate/end markers
- [ ] Field replacement preserves formatting
- [ ] OOXML metadata (settings.xml) is read and applied
- [ ] ODSO field mappings are respected during merge
- [ ] GREETINGLINE field generates proper greetings
- [ ] SKIPIF field conditionally skips records
- [ ] Format switches (\\* Upper, \\* Lower, etc.) work correctly
- [ ] Case-insensitive field matching works
- [ ] Merged documents (SimpleField) pass validation in Word and LibreOffice
- [ ] Merged documents (ComplexField) pass validation in Word and LibreOffice
- [ ] Integration tests cover all major workflows
- [ ] Examples demonstrate both field types and advanced fields
- [ ] API documentation explains SimpleField vs ComplexField clearly
- [ ] API documentation is complete and clear
- [ ] Roundtrip tests pass (merge → save → open → verify)
- [ ] Performance is acceptable (1000 records in <5 seconds)
- [ ] Memory usage is reasonable (<100MB for typical merges)
