# Design: Word Mail Merge Engine Architecture

## Problem Statement

The wordprocessing package has the XML elements to represent mail merge fields (MailMergeFieldType, FieldMapData, FieldCode, FieldChar) but no way to actually execute a mail merge. Users can create templates with MERGEFIELD codes but cannot connect data sources or produce merged output documents. This blocks all document automation workflows that require data-driven document generation.

## Current Architecture

```
WordprocessingDocument
    ↓
Document.Body
    ↓
Paragraph → Run → FieldCode("MERGEFIELD CustomerName")
    ↓
??? No merge execution ???
```

**Gap**: No engine to find fields, connect data, and produce merged output.

## Proposed Architecture

```
User Application
    ↓
WordprocessingDocument.MailMerge()
    ↓
MailMerge.DataSource(csvSource)
    ↓
MailMerge.Execute() / ExecuteToDocuments()
    ↓
FieldFinder → Finds all MERGEFIELD codes
    ↓
FieldMapper → Maps data source fields to merge field names
    ↓
FieldReplacer → Replaces field codes with data values
    ↓
DocumentBuilder → Constructs output document(s)
    ↓
Merged WordprocessingDocument(s)
```

## Design Alternatives Considered

### Alternative 1: Template Engine Approach (Mustache-style)
**Approach**: Replace MERGEFIELD with template syntax like {{CustomerName}}

**Pros**:
- Familiar to developers
- Simple text replacement

**Cons**:
- Breaks compatibility with Word
- Users must use custom syntax
- Loses Word's built-in field features

**Decision**: REJECTED - Must support standard MERGEFIELD codes for Word compatibility

### Alternative 2: LINQ-style Query Interface
**Approach**: Fluent query API for data selection
```go
template.MailMerge().
    From(csvSource).
    Where(record => record.Get("Status") == "Active").
    OrderBy("LastName").
    Execute()
```

**Pros**:
- Powerful data filtering
- Familiar to .NET developers

**Cons**:
- Overengineered for v1
- Complex implementation
- Users can filter data before calling merge

**Decision**: REJECTED - Too complex for initial implementation, defer to Phase 2

### Alternative 3: Callback-based Iteration
**Approach**: User provides callback for each record
```go
template.MailMerge().Execute(func(record map[string]string) {
    // User populates record
})
```

**Pros**:
- Maximum flexibility
- No DataSource abstraction needed

**Cons**:
- Awkward API for common cases
- Difficult to test
- Verbose for simple CSV/JSON cases

**Decision**: REJECTED - DataSource interface is cleaner and more testable

### Alternative 4: DataSource Interface (CHOSEN)
**Approach**: Define iterator-based DataSource interface with built-in implementations

**Pros**:
- Clean separation of data access from merge logic
- Extensible (users can implement custom sources)
- Testable (mock DataSource for unit tests)
- Simple implementations for common cases (CSV, JSON, Map)

**Cons**:
- One extra abstraction layer

**Decision**: ACCEPTED - Best balance of simplicity and extensibility

## DataSource Interface Design

### Core Interface

```go
package mailmerge

// DataSource provides access to merge data records.
type DataSource interface {
    // Open initializes the data source and prepares for iteration.
    // Must be called before Next/Get.
    Open() error

    // Close releases resources held by the data source.
    // Safe to call multiple times.
    Close() error

    // Next advances to the next record.
    // Returns true if a record is available, false if iteration is complete.
    // Must call Open before first Next call.
    Next() bool

    // Fields returns the list of available field names.
    // Field names are used to match MERGEFIELD codes in the template.
    // Valid after Open, before Close.
    Fields() []string

    // Get retrieves the value of the named field for the current record.
    // Field names are case-insensitive.
    // Returns empty string if field not found (unless strict mode enabled).
    // Valid after successful Next call.
    Get(fieldName string) (string, error)
}
```

### CSVDataSource Implementation

```go
package mailmerge

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strings"
)

// CSVDataSource provides merge data from CSV files.
type CSVDataSource struct {
    filename string
    file     *os.File
    reader   *csv.Reader
    headers  []string
    current  []string
    err      error
}

// NewCSVDataSource creates a DataSource from a CSV file.
// The first row of the CSV must contain field names (headers).
func NewCSVDataSource(filename string) *CSVDataSource {
    return &CSVDataSource{
        filename: filename,
    }
}

// Open opens the CSV file and reads the header row.
func (ds *CSVDataSource) Open() error {
    f, err := os.Open(ds.filename)
    if err != nil {
        return fmt.Errorf("failed to open CSV file: %w", err)
    }
    ds.file = f
    ds.reader = csv.NewReader(f)

    // Read header row
    headers, err := ds.reader.Read()
    if err != nil {
        ds.file.Close()
        return fmt.Errorf("failed to read CSV headers: %w", err)
    }

    // Trim whitespace from headers
    for i, h := range headers {
        headers[i] = strings.TrimSpace(h)
    }
    ds.headers = headers

    return nil
}

// Close closes the CSV file.
func (ds *CSVDataSource) Close() error {
    if ds.file != nil {
        return ds.file.Close()
    }
    return nil
}

// Next advances to the next record.
func (ds *CSVDataSource) Next() bool {
    record, err := ds.reader.Read()
    if err == io.EOF {
        return false
    }
    if err != nil {
        ds.err = err
        return false
    }

    ds.current = record
    return true
}

// Fields returns the CSV column headers.
func (ds *CSVDataSource) Fields() []string {
    return ds.headers
}

// Get retrieves the value for a field in the current record.
// Field names are matched case-insensitively.
func (ds *CSVDataSource) Get(fieldName string) (string, error) {
    // Case-insensitive lookup
    for i, header := range ds.headers {
        if strings.EqualFold(header, fieldName) {
            if i < len(ds.current) {
                return strings.TrimSpace(ds.current[i]), nil
            }
            return "", nil // Column exists but no value in this row
        }
    }
    return "", fmt.Errorf("field not found: %s", fieldName)
}

// Err returns any error encountered during iteration.
func (ds *CSVDataSource) Err() error {
    return ds.err
}
```

### JSONDataSource Implementation

```go
package mailmerge

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

// JSONDataSource provides merge data from JSON files.
// Expects a JSON array of objects: [{"Field1": "Value1", ...}, ...]
type JSONDataSource struct {
    filename string
    records  []map[string]interface{}
    index    int
    fields   []string
}

// NewJSONDataSource creates a DataSource from a JSON file.
func NewJSONDataSource(filename string) *JSONDataSource {
    return &JSONDataSource{
        filename: filename,
        index:    -1,
    }
}

// Open reads and parses the JSON file.
func (ds *JSONDataSource) Open() error {
    data, err := os.ReadFile(ds.filename)
    if err != nil {
        return fmt.Errorf("failed to read JSON file: %w", err)
    }

    err = json.Unmarshal(data, &ds.records)
    if err != nil {
        return fmt.Errorf("failed to parse JSON: %w", err)
    }

    // Collect all unique field names from first record
    if len(ds.records) > 0 {
        for key := range ds.records[0] {
            ds.fields = append(ds.fields, key)
        }
    }

    return nil
}

// Close releases resources (no-op for JSON).
func (ds *JSONDataSource) Close() error {
    return nil
}

// Next advances to the next record.
func (ds *JSONDataSource) Next() bool {
    ds.index++
    return ds.index < len(ds.records)
}

// Fields returns the JSON object keys from the first record.
func (ds *JSONDataSource) Fields() []string {
    return ds.fields
}

// Get retrieves the value for a field in the current record.
func (ds *JSONDataSource) Get(fieldName string) (string, error) {
    if ds.index < 0 || ds.index >= len(ds.records) {
        return "", fmt.Errorf("no current record")
    }

    record := ds.records[ds.index]

    // Case-insensitive lookup
    for key, value := range record {
        if strings.EqualFold(key, fieldName) {
            return fmt.Sprintf("%v", value), nil
        }
    }

    return "", fmt.Errorf("field not found: %s", fieldName)
}
```

### MapDataSource Implementation

```go
package mailmerge

import (
    "fmt"
    "strings"
)

// MapDataSource provides merge data from in-memory maps.
// Useful for testing and programmatic data population.
type MapDataSource struct {
    records []map[string]string
    index   int
    fields  []string
}

// NewMapDataSource creates a DataSource from a slice of maps.
func NewMapDataSource(records []map[string]string) *MapDataSource {
    ds := &MapDataSource{
        records: records,
        index:   -1,
    }

    // Collect all unique field names
    seen := make(map[string]bool)
    for _, record := range records {
        for key := range record {
            if !seen[key] {
                seen[key] = true
                ds.fields = append(ds.fields, key)
            }
        }
    }

    return ds
}

// Open prepares the data source (no-op for in-memory).
func (ds *MapDataSource) Open() error {
    return nil
}

// Close releases resources (no-op for in-memory).
func (ds *MapDataSource) Close() error {
    return nil
}

// Next advances to the next record.
func (ds *MapDataSource) Next() bool {
    ds.index++
    return ds.index < len(ds.records)
}

// Fields returns all unique field names across all records.
func (ds *MapDataSource) Fields() []string {
    return ds.fields
}

// Get retrieves the value for a field in the current record.
func (ds *MapDataSource) Get(fieldName string) (string, error) {
    if ds.index < 0 || ds.index >= len(ds.records) {
        return "", fmt.Errorf("no current record")
    }

    record := ds.records[ds.index]

    // Case-insensitive lookup
    for key, value := range record {
        if strings.EqualFold(key, fieldName) {
            return value, nil
        }
    }

    return "", fmt.Errorf("field not found: %s", fieldName)
}
```

## Merge Engine Design

### MailMerge Type

```go
package mailmerge

import (
    "fmt"
    "github.com/yourusername/goffice/wordprocessing"
)

// MailMerge coordinates mail merge execution.
type MailMerge struct {
    template   *wordprocessing.WordprocessingDocument
    dataSource DataSource
    options    MergeOptions
}

// MergeOptions configures merge behavior.
type MergeOptions struct {
    // StrictFields causes merge to fail if a field in the template
    // is not found in the data source.
    StrictFields bool

    // RemoveUnusedFields removes MERGEFIELD codes that don't match
    // any data source fields.
    RemoveUnusedFields bool
}

// DefaultMergeOptions returns sensible defaults.
func DefaultMergeOptions() MergeOptions {
    return MergeOptions{
        StrictFields:       false, // Empty string for missing fields
        RemoveUnusedFields: false, // Keep unmatched fields
    }
}

// New creates a MailMerge for the given template.
func New(template *wordprocessing.WordprocessingDocument) *MailMerge {
    return &MailMerge{
        template: template,
        options:  DefaultMergeOptions(),
    }
}

// DataSource sets the data source for the merge.
func (mm *MailMerge) DataSource(ds DataSource) *MailMerge {
    mm.dataSource = ds
    return mm
}

// Options sets merge options.
func (mm *MailMerge) Options(opts MergeOptions) *MailMerge {
    mm.options = opts
    return mm
}
```

### Merge Execution Algorithm

```go
// Execute performs mail merge and returns a single document with all records.
// Records are separated by page breaks.
func (mm *MailMerge) Execute() (*wordprocessing.WordprocessingDocument, error) {
    if mm.dataSource == nil {
        return nil, fmt.Errorf("data source not set")
    }

    // Open data source
    if err := mm.dataSource.Open(); err != nil {
        return nil, fmt.Errorf("failed to open data source: %w", err)
    }
    defer mm.dataSource.Close()

    // Find all merge fields in template
    fields, err := mm.findMergeFields()
    if err != nil {
        return nil, fmt.Errorf("failed to find merge fields: %w", err)
    }

    // Validate data source has required fields
    if err := mm.validateFields(fields); err != nil {
        return nil, err
    }

    // Create output document
    output, err := mm.template.Clone()
    if err != nil {
        return nil, fmt.Errorf("failed to clone template: %w", err)
    }

    body := output.MainDocumentPart().Document().Body()

    // Clear existing body content (will be replaced with merged content)
    body.RemoveAllChildren()

    // Iterate records and merge
    recordCount := 0
    for mm.dataSource.Next() {
        // Clone template body for this record
        recordBody, err := mm.template.MainDocumentPart().Document().Body().Clone()
        if err != nil {
            return nil, fmt.Errorf("failed to clone body: %w", err)
        }

        // Replace fields in this record's body
        if err := mm.replaceFi

elds(recordBody, fields); err != nil {
            return nil, fmt.Errorf("failed to replace fields in record %d: %w", recordCount, err)
        }

        // Add page break before subsequent records
        if recordCount > 0 {
            // Insert page break paragraph
            breakPara := body.AppendParagraph()
            breakRun := breakPara.AppendRun()
            breakRun.AppendBreak(wordprocessing.BreakType.Page)
        }

        // Append merged content to output
        for _, child := range recordBody.Children() {
            body.AppendChild(child)
        }

        recordCount++
    }

    if recordCount == 0 {
        return nil, fmt.Errorf("no records in data source")
    }

    return output, nil
}

// ExecuteToDocuments performs mail merge and returns one document per record.
func (mm *MailMerge) ExecuteToDocuments() ([]*wordprocessing.WordprocessingDocument, error) {
    if mm.dataSource == nil {
        return nil, fmt.Errorf("data source not set")
    }

    // Open data source
    if err := mm.dataSource.Open(); err != nil {
        return nil, fmt.Errorf("failed to open data source: %w", err)
    }
    defer mm.dataSource.Close()

    // Find all merge fields in template
    fields, err := mm.findMergeFields()
    if err != nil {
        return nil, fmt.Errorf("failed to find merge fields: %w", err)
    }

    // Validate data source has required fields
    if err := mm.validateFields(fields); err != nil {
        return nil, err
    }

    // Iterate records and create documents
    var documents []*wordprocessing.WordprocessingDocument
    recordCount := 0

    for mm.dataSource.Next() {
        // Clone template for this record
        doc, err := mm.template.Clone()
        if err != nil {
            return nil, fmt.Errorf("failed to clone template for record %d: %w", recordCount, err)
        }

        body := doc.MainDocumentPart().Document().Body()

        // Replace fields in this document
        if err := mm.replaceFields(body, fields); err != nil {
            return nil, fmt.Errorf("failed to replace fields in record %d: %w", recordCount, err)
        }

        documents = append(documents, doc)
        recordCount++
    }

    if recordCount == 0 {
        return nil, fmt.Errorf("no records in data source")
    }

    return documents, nil
}
```

### Field Finding Algorithm

```go
// MergeField represents a MERGEFIELD found in the template.
type MergeField struct {
    FieldName   string
    FieldType   FieldType // SimpleField or ComplexField
    Element     openxml.Element // For SimpleField: the w:fldSimple element
    BeginRun    *wordprocessing.Run // For ComplexField: run containing begin marker
    EndRun      *wordprocessing.Run // For ComplexField: run containing end marker
    FieldCode   string // Full reconstructed field code
    Switches    map[string]string // Parsed field switches
}

type FieldType int
const (
    FieldTypeSimple FieldType = iota
    FieldTypeComplex
)

// findMergeFields scans the document tree for MERGEFIELD codes.
// Handles BOTH SimpleField and ComplexField types.
// Searches body, tables, headers, footers, text boxes.
func (mm *MailMerge) findMergeFields() ([]MergeField, error) {
    var fields []MergeField

    // Search in main document body
    bodyFields := mm.findFieldsInElement(mm.template.MainDocumentPart().Document().Body())
    fields = append(fields, bodyFields...)

    // Search in headers
    for _, headerPart := range mm.template.HeaderParts() {
        headerFields := mm.findFieldsInElement(headerPart.Header())
        fields = append(fields, headerFields...)
    }

    // Search in footers
    for _, footerPart := range mm.template.FooterParts() {
        footerFields := mm.findFieldsInElement(footerPart.Footer())
        fields = append(fields, footerFields...)
    }

    return fields, nil
}

// findFieldsInElement recursively walks an element tree finding merge fields.
func (mm *MailMerge) findFieldsInElement(elem openxml.Element) []MergeField {
    var fields []MergeField

    // Check if this is a SimpleField element
    if simpleField, ok := elem.(*wordprocessing.SimpleField); ok {
        if instr := simpleField.Instruction(); instr != nil {
            fieldCode := instr.Value()
            if fieldName, switches := mm.parseMergeFieldCode(fieldCode); fieldName != "" {
                fields = append(fields, MergeField{
                    FieldName:   fieldName,
                    FieldType:   FieldTypeSimple,
                    Element:     simpleField,
                    FieldCode:   fieldCode,
                    Switches:    switches,
                })
            }
        }
    }

    // Check for ComplexField (w:fldChar begin...end spans)
    if para, ok := elem.(*wordprocessing.Paragraph); ok {
        complexFields := mm.findComplexFieldsInParagraph(para)
        fields = append(fields, complexFields...)
    }

    // Recursively search in tables
    if table, ok := elem.(*wordprocessing.Table); ok {
        for _, row := range table.Rows() {
            for _, cell := range row.Cells() {
                cellFields := mm.findFieldsInElement(cell)
                fields = append(fields, cellFields...)
            }
        }
    }

    // Recursively search children
    if composite, ok := elem.(openxml.CompositeElement); ok {
        for _, child := range composite.Children() {
            childFields := mm.findFieldsInElement(child)
            fields = append(fields, childFields...)
        }
    }

    return fields
}

// findComplexFieldsInParagraph finds ComplexField spans in a paragraph.
// ComplexField structure:
//   Run1: <w:fldChar w:fldCharType="begin"/>
//   Run2: <w:instrText>MERGEFIELD FirstName</w:instrText> (may span multiple runs)
//   Run3: <w:fldChar w:fldCharType="separate"/>
//   Run4-N: Field result text (placeholder)
//   RunN+1: <w:fldChar w:fldCharType="end"/>
func (mm *MailMerge) findComplexFieldsInParagraph(para *wordprocessing.Paragraph) []MergeField {
    var fields []MergeField
    runs := para.Runs()

    for i := 0; i < len(runs); i++ {
        run := runs[i]

        // Look for field begin marker
        if mm.hasFieldChar(run, "begin") {
            // Reconstruct field code from subsequent runs
            fieldCode, endIndex := mm.reconstructFieldCode(runs, i+1)

            if fieldName, switches := mm.parseMergeFieldCode(fieldCode); fieldName != "" {
                fields = append(fields, MergeField{
                    FieldName:   fieldName,
                    FieldType:   FieldTypeComplex,
                    BeginRun:    run,
                    EndRun:      runs[endIndex], // Run containing "end" marker
                    FieldCode:   fieldCode,
                    Switches:    switches,
                })
            }

            // Skip to end of field
            i = endIndex
        }
    }

    return fields
}

// reconstructFieldCode reconstructs the field code from instrText spans.
// Returns the complete field code and the index of the "end" marker run.
func (mm *MailMerge) reconstructFieldCode(runs []*wordprocessing.Run, startIndex int) (string, int) {
    var codeParts []string
    endIndex := -1

    for i := startIndex; i < len(runs); i++ {
        run := runs[i]

        // Check for instrText elements
        for _, child := range run.Children() {
            if instrText, ok := child.(*wordprocessing.FieldCode); ok {
                codeParts = append(codeParts, instrText.Text())
            }
        }

        // Check for separate marker (end of instruction, start of result)
        if mm.hasFieldChar(run, "separate") {
            // Continue to find "end" marker
            for j := i + 1; j < len(runs); j++ {
                if mm.hasFieldChar(runs[j], "end") {
                    endIndex = j
                    break
                }
            }
            break
        }

        // Check for end marker (field with no result)
        if mm.hasFieldChar(run, "end") {
            endIndex = i
            break
        }
    }

    return strings.TrimSpace(strings.Join(codeParts, " ")), endIndex
}

// hasFieldChar checks if a run contains a FieldChar with the specified type.
func (mm *MailMerge) hasFieldChar(run *wordprocessing.Run, charType string) bool {
    for _, child := range run.Children() {
        if fieldChar, ok := child.(*wordprocessing.FieldChar); ok {
            if fldCharType := fieldChar.GetAttribute(NamespaceWML, "fldCharType"); fldCharType == charType {
                return true
            }
        }
    }
    return false
}

// parseMergeFieldCode extracts the field name and switches from a field code.
// Returns field name and map of switches.
// Examples:
//   "MERGEFIELD CustomerName" -> ("CustomerName", {})
//   "MERGEFIELD Address \\* MERGEFORMAT" -> ("Address", {"*": "MERGEFORMAT"})
//   "GREETINGLINE \\f \" \\l \"Dear \"" -> ("", {"_greetingline": true, "f": " ", "l": "Dear "})
func (mm *MailMerge) parseMergeFieldCode(codeText string) (string, map[string]string) {
    switches := make(map[string]string)

    // Split by whitespace (but preserve quoted strings)
    parts := mm.tokenizeFieldCode(codeText)

    if len(parts) == 0 {
        return "", switches
    }

    // First part is field type (MERGEFIELD, GREETINGLINE, etc.)
    fieldType := strings.ToUpper(parts[0])

    switch fieldType {
    case "MERGEFIELD":
        // Format: MERGEFIELD FieldName [switches...]
        if len(parts) < 2 {
            return "", switches
        }
        fieldName := parts[1]

        // Parse switches (\\* MERGEFORMAT, \\b "text", etc.)
        for i := 2; i < len(parts); i++ {
            if strings.HasPrefix(parts[i], "\\") {
                switchName := strings.TrimPrefix(parts[i], "\\")
                switchValue := ""
                if i+1 < len(parts) && !strings.HasPrefix(parts[i+1], "\\") {
                    switchValue = parts[i+1]
                    i++
                }
                switches[switchName] = switchValue
            }
        }

        return fieldName, switches

    case "GREETINGLINE":
        // GREETINGLINE doesn't have a field name, it's a special field
        switches["_greetingline"] = "true"
        // Parse GREETINGLINE switches
        for i := 1; i < len(parts); i++ {
            if strings.HasPrefix(parts[i], "\\") {
                switchName := strings.TrimPrefix(parts[i], "\\")
                switchValue := ""
                if i+1 < len(parts) && !strings.HasPrefix(parts[i+1], "\\") {
                    switchValue = parts[i+1]
                    i++
                }
                switches[switchName] = switchValue
            }
        }
        return "", switches

    default:
        // Unknown field type
        return "", switches
    }
}

// tokenizeFieldCode splits a field code into tokens, preserving quoted strings.
func (mm *MailMerge) tokenizeFieldCode(code string) []string {
    var tokens []string
    var current strings.Builder
    inQuote := false

    for i := 0; i < len(code); i++ {
        ch := code[i]

        if ch == '"' {
            inQuote = !inQuote
            current.WriteByte(ch)
        } else if ch == ' ' && !inQuote {
            if current.Len() > 0 {
                tokens = append(tokens, current.String())
                current.Reset()
            }
        } else {
            current.WriteByte(ch)
        }
    }

    if current.Len() > 0 {
        tokens = append(tokens, current.String())
    }

    return tokens
}
```

### Field Replacement Algorithm

```go
// replaceFields replaces all merge fields with data values.
// Handles both SimpleField and ComplexField types.
func (mm *MailMerge) replaceFields(body *wordprocessing.Body, fields []MergeField) error {
    for _, field := range fields {
        // Get replacement value
        value, err := mm.getFieldValue(field)
        if err != nil {
            // Handle missing field
            if mm.options.StrictFields {
                return fmt.Errorf("field %s: %w", field.FieldName, err)
            }
            // Use empty string for missing field
            value = ""
        }

        // Replace based on field type
        switch field.FieldType {
        case FieldTypeSimple:
            if err := mm.replaceSimpleField(field, value); err != nil {
                return err
            }
        case FieldTypeComplex:
            if err := mm.replaceComplexField(field, value); err != nil {
                return err
            }
        }
    }

    return nil
}

// getFieldValue retrieves the value for a merge field.
// Handles special field types (GREETINGLINE) and standard MERGEFIELD.
func (mm *MailMerge) getFieldValue(field MergeField) (string, error) {
    // Check for special fields
    if field.Switches["_greetingline"] == "true" {
        return mm.generateGreetingLine(field.Switches)
    }

    // Standard MERGEFIELD - get from data source
    if field.FieldName == "" {
        return "", fmt.Errorf("empty field name")
    }

    value, err := mm.dataSource.Get(field.FieldName)
    if err != nil {
        return "", err
    }

    // Apply formatting switches if present
    if formatSwitch := field.Switches["*"]; formatSwitch != "" {
        value = mm.applyFormatSwitch(value, formatSwitch)
    }

    return value, nil
}

// replaceSimpleField replaces a SimpleField (w:fldSimple) element.
// Strategy: Replace the entire w:fldSimple element with a Run containing text.
func (mm *MailMerge) replaceSimpleField(field MergeField, value string) error {
    simpleField := field.Element.(*wordprocessing.SimpleField)

    // Get parent element (usually a paragraph)
    parent := simpleField.Parent()
    if parent == nil {
        return fmt.Errorf("SimpleField has no parent")
    }

    // Create replacement run with text
    run := wordprocessing.NewRun()
    run.AppendText(value)

    // Copy formatting from SimpleField to run (if SimpleField has run properties)
    // SimpleField can contain runs with formatting - preserve first run's formatting
    if len(simpleField.Children()) > 0 {
        if firstRun, ok := simpleField.Children()[0].(*wordprocessing.Run); ok {
            if runProps := firstRun.RunProperties(); runProps != nil {
                run.SetRunProperties(runProps.Clone())
            }
        }
    }

    // Replace SimpleField with Run in parent
    parentComposite, ok := parent.(openxml.CompositeElement)
    if !ok {
        return fmt.Errorf("SimpleField parent is not a composite element")
    }

    // Find index of SimpleField in parent
    children := parentComposite.Children()
    for i, child := range children {
        if child == field.Element {
            // Remove SimpleField
            parentComposite.RemoveChildAt(i)
            // Insert replacement run at same position
            parentComposite.InsertChildAt(i, run)
            return nil
        }
    }

    return fmt.Errorf("SimpleField not found in parent")
}

// replaceComplexField replaces a ComplexField (w:fldChar begin...end span).
// Strategy: Remove all runs between begin and end, insert new run with value.
func (mm *MailMerge) replaceComplexField(field MergeField, value string) error {
    // Get paragraph containing the field
    para := field.BeginRun.Parent()
    if para == nil {
        return fmt.Errorf("ComplexField BeginRun has no parent")
    }

    paraComposite, ok := para.(openxml.CompositeElement)
    if !ok {
        return fmt.Errorf("paragraph is not a composite element")
    }

    // Find indices of begin and end runs
    children := paraComposite.Children()
    beginIndex := -1
    endIndex := -1

    for i, child := range children {
        if child == field.BeginRun {
            beginIndex = i
        }
        if child == field.EndRun {
            endIndex = i
            break
        }
    }

    if beginIndex == -1 || endIndex == -1 {
        return fmt.Errorf("ComplexField runs not found in paragraph")
    }

    // Create replacement run with text
    run := wordprocessing.NewRun()
    run.AppendText(value)

    // Preserve formatting from the result run (between separate and end markers)
    // Find the first run after separate marker
    for i := beginIndex + 1; i < endIndex; i++ {
        if resultRun, ok := children[i].(*wordprocessing.Run); ok {
            // Check if this is after the separate marker
            hasSeparate := false
            for j := beginIndex + 1; j < i; j++ {
                if r, ok := children[j].(*wordprocessing.Run); ok {
                    if mm.hasFieldChar(r, "separate") {
                        hasSeparate = true
                        break
                    }
                }
            }

            if hasSeparate {
                // This run is in the result section - copy its formatting
                if runProps := resultRun.RunProperties(); runProps != nil {
                    run.SetRunProperties(runProps.Clone())
                }
                break
            }
        }
    }

    // Remove all runs from beginIndex to endIndex (inclusive)
    for i := endIndex; i >= beginIndex; i-- {
        paraComposite.RemoveChildAt(i)
    }

    // Insert replacement run at beginIndex
    paraComposite.InsertChildAt(beginIndex, run)

    return nil
}

// generateGreetingLine generates a greeting line from field switches and data.
// Example: "Dear Mr. Smith,"
func (mm *MailMerge) generateGreetingLine(switches map[string]string) (string, error) {
    // Default greeting format: "Dear [Title] [LastName],"

    title, _ := mm.dataSource.Get("Title")
    firstName, _ := mm.dataSource.Get("FirstName")
    lastName, _ := mm.dataSource.Get("LastName")

    // Check for custom format switch (\\f)
    if format := switches["f"]; format != "" {
        // Use custom format (simplified - full implementation would parse format codes)
        return format, nil
    }

    // Default greeting
    if title != "" && lastName != "" {
        return fmt.Sprintf("Dear %s %s,", title, lastName)
    } else if firstName != "" && lastName != "" {
        return fmt.Sprintf("Dear %s %s,", firstName, lastName)
    } else if lastName != "" {
        return fmt.Sprintf("Dear %s,", lastName)
    } else if firstName != "" {
        return fmt.Sprintf("Dear %s,", firstName)
    }

    return "Dear Recipient,", nil
}

// applyFormatSwitch applies a format switch to a value.
// Example: \\* Upper -> "JOHN DOE"
func (mm *MailMerge) applyFormatSwitch(value, formatSwitch string) string {
    switch strings.ToUpper(formatSwitch) {
    case "UPPER":
        return strings.ToUpper(value)
    case "LOWER":
        return strings.ToLower(value)
    case "FIRSTCAP":
        if len(value) > 0 {
            return strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
        }
    case "CAPS":
        return strings.Title(strings.ToLower(value))
    case "MERGEFORMAT":
        // Preserve formatting - no change to value
        return value
    }
    return value
}

// validateFields checks that the data source has all required fields.
func (mm *MailMerge) validateFields(fields []MergeField) error {
    if !mm.options.StrictFields {
        return nil // Validation not required
    }

    dataFields := mm.dataSource.Fields()
    dataFieldSet := make(map[string]bool)
    for _, f := range dataFields {
        dataFieldSet[strings.ToLower(f)] = true
    }

    for _, field := range fields {
        if field.FieldName == "" {
            continue // Skip special fields like GREETINGLINE
        }
        if !dataFieldSet[strings.ToLower(field.FieldName)] {
            return fmt.Errorf("template field %s not found in data source", field.FieldName)
        }
    }

    return nil
}
```

### OOXML Metadata Integration

```go
// MailMergeMetadata represents mail merge settings from settings.xml.
type MailMergeMetadata struct {
    DataType      string // Type of data source (database, spreadsheet, text, etc.)
    ConnectString string // Connection string to data source
    Query         string // SQL query or data selection
    FieldMappings map[string]string // Maps template field names to data column names
}

// loadMailMergeMetadata reads mail merge configuration from settings.xml.
func (mm *MailMerge) loadMailMergeMetadata() (*MailMergeMetadata, error) {
    // Get settings part
    settingsPart := mm.template.DocumentSettingsPart()
    if settingsPart == nil {
        return nil, nil // No settings.xml - use direct field matching
    }

    settings := settingsPart.Settings()
    if settings == nil {
        return nil, nil
    }

    // Look for w:mailMerge element
    mailMergeElem := settings.MailMerge()
    if mailMergeElem == nil {
        return nil, nil // No mail merge configuration
    }

    metadata := &MailMergeMetadata{
        FieldMappings: make(map[string]string),
    }

    // Read mailMerge properties
    if dataType := mailMergeElem.DataType(); dataType != nil {
        metadata.DataType = dataType.Value()
    }

    if connectString := mailMergeElem.ConnectString(); connectString != nil {
        metadata.ConnectString = connectString.Value()
    }

    if query := mailMergeElem.Query(); query != nil {
        metadata.Query = query.Value()
    }

    // Read ODSO (Open Data Source Object) field mappings
    if odso := mailMergeElem.ODSO(); odso != nil {
        if fieldMapData := odso.FieldMapData(); fieldMapData != nil {
            for _, mapping := range fieldMapData {
                // mapping has Name (template field) and MappedName (data column)
                if name := mapping.Name(); name != nil {
                    if mappedName := mapping.MappedName(); mappedName != nil {
                        metadata.FieldMappings[name.Value()] = mappedName.Value()
                    }
                }
            }
        }
    }

    return metadata, nil
}

// applyFieldMappings applies field name mappings to a field name.
// If a mapping exists, returns the mapped name; otherwise returns original.
func (mm *MailMerge) applyFieldMappings(fieldName string, metadata *MailMergeMetadata) string {
    if metadata == nil {
        return fieldName
    }

    if mappedName, exists := metadata.FieldMappings[fieldName]; exists {
        return mappedName
    }

    return fieldName
}

// updateMailMergeMetadata optionally updates settings.xml with data source info.
func (mm *MailMerge) updateMailMergeMetadata(dataSourceType string) error {
    // This is optional - only update if user requests it
    if !mm.options.UpdateSettings {
        return nil
    }

    settingsPart := mm.template.DocumentSettingsPart()
    if settingsPart == nil {
        // Create settings part if it doesn't exist
        settingsPart = mm.template.AddDocumentSettingsPart()
    }

    settings := settingsPart.Settings()
    if settings.MailMerge() == nil {
        mailMerge := wordprocessing.NewMailMerge()
        settings.AppendChild(mailMerge)
    }

    mailMerge := settings.MailMerge()

    // Set data type based on data source
    dataType := wordprocessing.NewMailMergeDataType()
    switch dataSourceType {
    case "csv":
        dataType.SetVal(wordprocessing.MailMergeDataValuesTextfile)
    case "json":
        dataType.SetVal(wordprocessing.MailMergeDataValuesNative)
    case "database":
        dataType.SetVal(wordprocessing.MailMergeDataValuesDatabase)
    default:
        dataType.SetVal(wordprocessing.MailMergeDataValuesNative)
    }

    mailMerge.SetDataType(dataType)

    return nil
}
```

## Integration with WordprocessingDocument

```go
// In wordprocessing/document.go

// MailMerge returns a MailMerge instance for executing mail merge operations.
func (d *WordprocessingDocument) MailMerge() *mailmerge.MailMerge {
    return mailmerge.New(d)
}
```

## Testing Strategy

### Unit Tests

**DataSource Tests**:
```go
func TestCSVDataSource(t *testing.T) {
    // Create test CSV file
    csvContent := `FirstName,LastName,Email
John,Doe,john@example.com
Jane,Smith,jane@example.com`
    tmpFile := createTempFile(t, csvContent)
    defer os.Remove(tmpFile)

    ds := NewCSVDataSource(tmpFile)
    err := ds.Open()
    assert.NoError(t, err)
    defer ds.Close()

    // Check fields
    assert.Equal(t, []string{"FirstName", "LastName", "Email"}, ds.Fields())

    // First record
    assert.True(t, ds.Next())
    assert.Equal(t, "John", ds.Get("FirstName"))
    assert.Equal(t, "Doe", ds.Get("LastName"))

    // Case-insensitive
    assert.Equal(t, "john@example.com", ds.Get("email"))

    // Second record
    assert.True(t, ds.Next())
    assert.Equal(t, "Jane", ds.Get("FirstName"))

    // EOF
    assert.False(t, ds.Next())
}
```

**Merge Engine Tests**:
```go
func TestMailMerge_Execute(t *testing.T) {
    // Create template with merge fields
    template := wordprocessing.Create("template.docx", DocumentType.Document)
    body := template.MainDocumentPart().Document().Body()
    para := body.AppendParagraph()
    run := para.AppendRun()
    run.AppendFieldCode("MERGEFIELD FirstName")
    run.AppendText(" ")
    run.AppendFieldCode("MERGEFIELD LastName")

    // Create data source
    data := []map[string]string{
        {"FirstName": "John", "LastName": "Doe"},
        {"FirstName": "Jane", "LastName": "Smith"},
    }
    ds := NewMapDataSource(data)

    // Execute merge
    merged, err := template.MailMerge().DataSource(ds).Execute()
    assert.NoError(t, err)

    // Verify merged content
    mergedBody := merged.MainDocumentPart().Document().Body()
    text := mergedBody.InnerText()
    assert.Contains(t, text, "John Doe")
    assert.Contains(t, text, "Jane Smith")
}
```

### Integration Tests

```go
func TestMailMerge_Integration_CSV(t *testing.T) {
    // Create template document
    template := createTemplateDocument(t)
    defer template.Close()

    // Create CSV data file
    csvFile := createCustomerCSV(t)
    defer os.Remove(csvFile)

    // Execute merge
    csvSource := NewCSVDataSource(csvFile)
    merged, err := template.MailMerge().DataSource(csvSource).Execute()
    assert.NoError(t, err)
    defer merged.Close()

    // Save merged document
    merged.SaveAs("merged-output.docx")

    // Roundtrip test: open and verify
    reopened, err := wordprocessing.Open("merged-output.docx", false)
    assert.NoError(t, err)
    defer reopened.Close()

    text := reopened.MainDocumentPart().Document().Body().InnerText()
    assert.Contains(t, text, "John")
    assert.Contains(t, text, "Jane")
}
```

## Performance Considerations

### Memory Optimization
- Lazy field finding (only parse field codes when needed)
- Stream CSV parsing (don't load entire file into memory)
- Reuse cloned sections instead of re-parsing XML

### Execution Time Optimization
- Cache parsed field codes
- Avoid redundant tree walks
- Clone efficiently (share immutable content where possible)

## Risks & Mitigation

### Risk: Complex Field Codes
**Impact**: Medium (some fields may not parse correctly)
**Probability**: Medium (Word supports complex field syntax)
**Mitigation**:
- Start with simple MERGEFIELD syntax
- Document supported field code formats
- Add complex field support in Phase 2

### Risk: Large Data Sources
**Impact**: Medium (memory pressure, slow execution)
**Probability**: Low (most mail merges are <10,000 records)
**Mitigation**:
- Stream-based CSV parsing
- ExecuteToDocuments mode for very large datasets
- Document memory considerations

### Risk: Special Characters in Data
**Impact**: Medium (XML encoding issues)
**Probability**: Medium (data may contain <, >, &, quotes)
**Mitigation**:
- Proper XML encoding when inserting text
- Test with special characters in data
- Document supported character sets

### Risk: Template Corruption
**Impact**: High (user loses work)
**Probability**: Low (templates are cloned, not modified)
**Mitigation**:
- Never modify original template
- All merge operations work on clones
- Validate clone success before proceeding

## Success Metrics

- [ ] CSVDataSource handles files up to 10MB efficiently
- [ ] JSONDataSource handles arrays up to 10,000 records
- [ ] Merge execution completes in <5 seconds for 1,000 records
- [ ] Memory usage stays under 100MB for typical merges
- [ ] Field finding accuracy is 100% for standard MERGEFIELD syntax
- [ ] Case-insensitive field matching works correctly
- [ ] Merged documents pass Word validation
- [ ] Merged documents open in LibreOffice without errors
- [ ] Roundtrip tests pass (merge → save → open → verify)
- [ ] Special characters (quotes, ampersands, Unicode) are handled correctly

## Future Enhancements (Not in Scope)

1. **Conditional Fields**: Support IF, COMPARE field codes
2. **Nested Regions**: Repeating sections for one-to-many data
3. **Formula Fields**: Calculated fields (SUM, COUNT, etc.)
4. **Image Fields**: Replace fields with images from data
5. **Database Connectivity**: SQL DataSource implementation
6. **Excel DataSource**: Read data from Excel workbooks
7. **Performance Optimizations**: Parallel merge execution
8. **Incremental Merge**: Merge subsets of records

These can be added incrementally without breaking the DataSource interface or merge API.
