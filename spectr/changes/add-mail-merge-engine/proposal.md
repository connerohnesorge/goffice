# Change: Add Mail Merge Engine

## Why

The goffice library has **mail merge field elements** (MailMergeFieldType, FieldMapData) but **no execution engine**. Users can create templates with MERGEFIELD codes but cannot actually merge data to produce output documents. This is a critical gap for document automation use cases:

- **Batch letter generation** - Cannot generate personalized letters from customer data
- **Certificate/badge creation** - Cannot produce certificates from participant lists
- **Report automation** - Cannot populate report templates from database exports
- **Form filling** - Cannot fill template forms with structured data

**Current State**:
- Elements exist in `wordprocessing/elements/elements.go` and `wordprocessing/elements/settings.go`
- SimpleField (w:fldSimple) with w:instr attribute for simple merge fields
- FieldChar (w:fldChar) + FieldCode (w:instrText) for complex multi-run fields
- MailMerge element in settings.xml with ODSO support infrastructure
- FieldMapData for mapping field names to data columns
- No way to connect data sources (CSV, JSON, database)
- No execution method to perform merge

**Root Cause**: Missing merge execution engine and data source abstraction.

**Impact**: **High (8/10)** - Blocks major document automation workflows. Users must resort to external tools or manual population.

## What Changes

**Core Mail Merge Engine**:
- Create `wordprocessing/mailmerge/` package
- Implement DataSource interface with Open/Close/Next/Fields/Get methods
- Implement CSVDataSource, JSONDataSource, MapDataSource
- Implement MailMerge engine with field finding and replacement logic
- Support two execution modes: Execute (single doc) and ExecuteToDocuments (multi-doc)

**Document Integration**:
- Add MailMerge() accessor to WordprocessingDocument
- Integrate with existing field infrastructure
- Preserve template document (create new output documents)

**Field Processing**:
- Find all MERGEFIELD codes in document tree (both SimpleField and ComplexField types)
- Support SimpleField (w:fldSimple with w:instr attribute)
- Support ComplexField (w:fldChar begin/separate/end + w:instrText spans)
- Search in paragraphs, tables, headers, footers, text boxes
- Parse field codes to extract field names and switches
- Case-insensitive field name matching
- Replace field codes with data values
- Handle missing fields gracefully (empty string or error)
- Support GREETINGLINE and conditional fields (SKIPIF)

**Testing & Documentation**:
- Unit tests for DataSource implementations
- Integration tests with real templates and data
- Roundtrip tests (merge → save → open → verify)
- Examples showing CSV merge, JSON merge, programmatic merge
- API documentation and usage guide

**Breaking changes**: None. This adds new functionality without changing existing APIs.

## Impact

**Affected specs**:
- `wordprocessing-document` (ADDED mail merge accessor method)
- `wordprocessing-mailmerge` (NEW spec for merge engine capability)

**New capabilities**:
- Connect data sources to templates
- Execute mail merge to single document with all records
- Execute mail merge to multiple documents (one per record)
- Support CSV, JSON, and in-memory data sources
- Extensible DataSource interface for custom sources (database, API, etc.)

**Affected code**:
- `wordprocessing/mailmerge/data_source.go` - NEW: DataSource interface and implementations
- `wordprocessing/mailmerge/engine.go` - NEW: MailMerge engine with Execute methods
- `wordprocessing/mailmerge/field_finder.go` - NEW: Field finding for SimpleField and ComplexField
- `wordprocessing/mailmerge/field_replacer.go` - NEW: Field replacement handling both field types
- `wordprocessing/mailmerge/advanced_fields.go` - NEW: GREETINGLINE and conditional field support
- `wordprocessing/mailmerge/metadata.go` - NEW: settings.xml and ODSO integration
- `wordprocessing/document.go` - MODIFIED: Add MailMerge() accessor
- `wordprocessing/settings_part.go` - MODIFIED: Access to MailMerge settings

## Key Design Decisions

### 1. DataSource Interface Design
Abstract data source to allow custom implementations:
```go
type DataSource interface {
    Open() error
    Close() error
    Next() bool               // Advance to next record
    Fields() []string         // Available field names
    Get(fieldName string) (string, error)
}
```

**Rationale**:
- Decouples merge engine from specific data formats
- Users can implement database sources, API sources, etc.
- Stdlib implementations (CSV, JSON) cover common cases
- Simple iterator pattern (Next/Get) is familiar

### 2. Two Execution Modes
**Mode A - Execute**: Merge all records into single document
```go
merged := template.MailMerge().DataSource(csvSource).Execute()
// Result: One document with all records (page break between records)
```

**Mode B - ExecuteToDocuments**: Merge each record to separate document
```go
docs := template.MailMerge().DataSource(csvSource).ExecuteToDocuments()
// Result: []WordprocessingDocument, one per record
```

**Rationale**:
- Mode A for mass printing (letters, certificates)
- Mode B for individual files (contracts, invoices)
- Both modes common in Microsoft Word mail merge
- Minimal API surface (two clear methods)

### 3. Field Name Case-Insensitivity
Field names matched case-insensitively (MERGEFIELD FirstName matches "firstname" in data).

**Rationale**:
- Matches Microsoft Word behavior
- Reduces user friction (no exact case matching required)
- Common in Office automation

### 4. Missing Field Handling
Strategy for missing fields:
- **Default**: Empty string (silent substitution)
- **Optional**: Error mode (strict validation)

**Rationale**:
- Default matches Word behavior (blank for missing fields)
- Strict mode available for validation scenarios
- Configurable via MailMerge options

### 5. Template Preservation
Original template document is never modified. Merge operations clone content.

**Rationale**:
- Reusable templates
- Prevents accidental corruption
- Clear separation of template vs output

### 6. Clone vs Append Strategy
For Execute (single document):
- Clone document body for each record
- Append cloned content with page breaks
- Remove original template content after merge

**Rationale**:
- Preserves section formatting per record
- Natural page breaks between records
- Efficient (no repeated document creation)

### 7. Dual Field Type Support (SimpleField vs ComplexField)
**OOXML defines TWO field representations**:

**SimpleField (w:fldSimple)**:
- Single element with w:instr attribute containing field code
- Example: `<w:fldSimple w:instr="MERGEFIELD FirstName"/>`
- MORE common in Word-generated templates
- Simpler to find and replace

**ComplexField (w:fldChar + w:instrText)**:
- Multi-element structure spanning runs
- Begin: `<w:fldChar w:fldCharType="begin"/>`
- Instruction: `<w:instrText>MERGEFIELD FirstName</w:instrText>` (may span multiple runs)
- Separate: `<w:fldChar w:fldCharType="separate"/>`
- Result: Placeholder text (replaced during merge)
- End: `<w:fldChar w:fldCharType="end"/>`
- Used for fields with formatting switches

**Rationale**:
- Both types exist in real Word documents
- SimpleField is more common but ComplexField is used for advanced formatting
- Must support both to handle all Word templates
- Field finding must walk entire document tree (not just body paragraphs)

### 8. OOXML Mail Merge Metadata Integration
**settings.xml contains mail merge configuration**:
- `<w:mailMerge>` element with dataType, connectString, query
- `<w:odso>` (Open Data Source Object) with field mappings
- `<w:fieldMapData>` for mapping template fields to data columns
- MailMergeRecipientDataPart for storing recipient data

**Integration Strategy**:
- Read settings.xml to discover existing mail merge configuration
- Respect field mappings if present (use mapped names)
- Optionally update settings.xml when executing merge (record data source used)
- Support templates with or without settings.xml metadata

**Rationale**:
- Word stores mail merge metadata in settings.xml
- Proper OOXML compliance requires reading this metadata
- Field mappings allow template field names to differ from data column names
- Optional integration (fallback to direct field matching if no metadata)

### 9. Advanced Field Types Support
**Phase 1 (Included)**:
- MERGEFIELD - Standard field replacement
- GREETINGLINE - Formatted greeting (e.g., "Dear Mr. Smith")

**Phase 2 (Future)**:
- SKIPIF - Conditionally skip record
- IF - Conditional content
- NEXT/NEXTIF - Control record iteration
- Nested fields

**Rationale**:
- MERGEFIELD is essential (MVP)
- GREETINGLINE is common in business letters
- Conditional fields are advanced but valuable
- Start simple, expand incrementally

## Implementation Scope

### Phase 1: Core Infrastructure (1 week)
- Define DataSource interface
- Implement CSVDataSource
- Implement MapDataSource (in-memory)
- Unit tests for data sources

### Phase 2: Field Finding & Parsing (1 week)
- Implement SimpleField detection and parsing
- Implement ComplexField detection (reconstruct field code from spans)
- Walk entire document tree (body, tables, headers, footers, text boxes)
- Parse field codes to extract field names and switches
- Handle both w:fldSimple and w:fldChar structures
- Unit tests for field finding with both field types

### Phase 3: Field Replacement (1 week)
- Implement SimpleField replacement (replace entire w:fldSimple element)
- Implement ComplexField replacement (handle begin/separate/end markers)
- Preserve run formatting during replacement
- Handle field switches and formatting codes
- Implement Execute method (single document output)
- Unit tests for replacement with both field types

### Phase 4: OOXML Metadata Integration (3 days)
- Read MailMerge element from settings.xml
- Parse ODSO field mappings
- Apply field name mappings during merge
- Optionally update settings.xml with data source metadata
- Tests for metadata reading and field mapping

### Phase 5: Multiple Document Mode & Advanced Fields (1 week)
- Implement ExecuteToDocuments method
- Document cloning per record
- Implement GREETINGLINE field support
- Add basic conditional field support (SKIPIF)
- Tests for multi-document output and advanced fields

### Phase 6: Advanced Data Sources (3 days)
- Implement JSONDataSource
- Add data source validation (required fields check)
- Error handling and reporting
- Tests for JSON and validation

### Phase 7: Integration & Examples (3 days)
- Add MailMerge() to WordprocessingDocument
- Create example templates with SimpleField and ComplexField
- Integration tests with real documents
- Roundtrip tests

### Phase 8: Documentation (2 days)
- API documentation
- Usage guide with code examples
- Document SimpleField vs ComplexField handling
- Document OOXML metadata integration
- Document supported field types (MERGEFIELD, GREETINGLINE, SKIPIF)
- Migration guide for users
- Update README with mail merge capabilities

**Total: 4 weeks**

### Out of Scope (Future Enhancements)
- Database connectivity (ODBC, SQL) - users implement DataSource interface
- Excel as data source - future enhancement
- Advanced conditional fields (IF with complex expressions, COMPARE) - Phase 2 feature
- Nested merge regions (repeating sections) - Phase 2 feature
- Email merge (sending emails) - separate feature
- Formula evaluation in fields - not supported
- NEXT/NEXTIF fields - Phase 2 feature
- Image merge fields - Phase 2 feature

## Dependencies

**Existing**:
- SimpleField element (wordprocessing/elements/elements.go)
- FieldChar element (wordprocessing/elements/special_chars.go)
- FieldCode (instrText) element
- MailMerge element (wordprocessing/elements/settings.go)
- FieldMapData element for ODSO mappings
- Document cloning/copying APIs
- Settings part access (settings.xml)
- Standard library CSV parser (encoding/csv)
- Standard library JSON parser (encoding/json)

**New**: None (all required OOXML elements already exist)

## Success Criteria

- [ ] DataSource interface is clear and extensible
- [ ] CSVDataSource can parse CSV files and iterate records
- [ ] JSONDataSource can parse JSON arrays and iterate records
- [ ] MapDataSource supports in-memory data (for testing/simple cases)
- [ ] MailMerge engine finds all MERGEFIELD codes (SimpleField and ComplexField)
- [ ] SimpleField elements are detected and parsed correctly
- [ ] ComplexField spans are reconstructed from multiple runs
- [ ] Field finding searches body, tables, headers, footers, text boxes
- [ ] Field names are extracted correctly from field codes
- [ ] Field values replace SimpleField elements correctly
- [ ] Field values replace ComplexField spans correctly (preserving formatting)
- [ ] OOXML metadata from settings.xml is read and applied
- [ ] Field name mappings from ODSO are respected
- [ ] GREETINGLINE fields generate proper greetings
- [ ] SKIPIF fields conditionally skip records
- [ ] Execute() produces single document with all records merged
- [ ] ExecuteToDocuments() produces one document per record
- [ ] Template document is not modified during merge
- [ ] Missing fields use empty string (or error in strict mode)
- [ ] Field name matching is case-insensitive
- [ ] Merged documents open correctly in Microsoft Word
- [ ] Merged documents open correctly in LibreOffice
- [ ] Special characters in data (quotes, newlines) are handled correctly
- [ ] Roundtrip tests pass (merge → save → open → verify)
- [ ] Examples demonstrate both field types and advanced fields
- [ ] API documentation is complete
