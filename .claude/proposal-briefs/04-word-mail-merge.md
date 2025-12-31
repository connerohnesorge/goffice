# Proposal Brief: Word Mail Merge Engine

## Priority: TIER 2 - High
**Impact**: 8/10  
**Effort**: 2 weeks  
**Status**: Elements exist, NO execution engine

## Gap Analysis

### Current State
**What exists:**
- ✅ **Mail merge field elements** (wordprocessing/elements/elements.go):
  - `MailMergeFieldType` - Field type definition
  - `FieldMapData` - Field mapping information  
- ✅ **Field infrastructure**: FieldCode, FieldChar can represent `MERGEFIELD` codes

**What's missing:**
- ❌ No mail merge data source connection
- ❌ No mail merge execution engine
- ❌ No template + data → multiple documents generation
- ❌ No `ExecuteMailMerge(dataSource)` method
- ❌ No data source abstraction (CSV, JSON, database)

### OpenXML-SDK Comparison
**Microsoft Word Object Model:**
```csharp
// Connect to data source
document.MailMerge.OpenDataSource("customers.csv")

// Execute merge
document.MailMerge.Execute()

// Or merge to new document
document.MailMerge.Destination = WdMailMergeDestination.wdSendToNewDocument
document.MailMerge.ExecuteWithRegions()
```

**Open-XML-SDK** provides low-level field manipulation but not merge execution.

## Implementation Scope

### Files to Create

1. **wordprocessing/mailmerge/data_source.go** (NEW package)
   - `DataSource` interface
   - `CSVDataSource` implementation
   - `JSONDataSource` implementation
   - `MapDataSource` implementation (in-memory)

2. **wordprocessing/mailmerge/engine.go** (NEW)
   - `MailMerge` struct with execution logic
   - `Execute()` method
   - `ExecuteToDocuments()` method (returns slice of documents)

3. **wordprocessing/mailmerge/field_mapper.go** (NEW)
   - Maps data source columns to merge field names
   - Handles field name case-insensitivity

4. **wordprocessing/document.go** (MODIFY)
   - Add `MailMerge() *mailmerge.MailMerge` accessor

### Core Design Decisions

**1. DataSource Interface**
```go
type DataSource interface {
    // Open initializes the data source
    Open() error
    
    // Close releases resources
    Close() error
    
    // Next advances to next record, returns false when done
    Next() bool
    
    // Fields returns available field names
    Fields() []string
    
    // Get returns value for field in current record
    Get(fieldName string) (string, error)
}
```

**2. Merge Field Detection**
Scan document for `MERGEFIELD` codes:
```
MERGEFIELD CustomerName
MERGEFIELD Address \* MERGEFORMAT
```

Parse field code to extract field name.

**3. Two Execution Modes**

**Mode A: Single Document (Merge All)**
```go
doc.MailMerge().
    DataSource(csvSource).
    Execute()
// Result: Single document with all records appended
```

**Mode B: Multiple Documents (Merge to Files)**
```go
docs := doc.MailMerge().
    DataSource(csvSource).
    ExecuteToDocuments()
// Result: []Document, one per record
```

**4. Field Replacement Logic**
For each record:
1. Clone template document (or section)
2. Find all `MERGEFIELD` codes
3. Replace field code with actual data value
4. Append to output document or add to documents slice

**5. Template Preservation**
Original template document is NOT modified. Merge creates new document(s).

## References

### Existing Code
- `wordprocessing/elements/elements.go` - MailMergeFieldType, FieldMapData
- `wordprocessing/document.go` - Document API
- `wordprocessing/builder.go` - Document building patterns

### Similar Implementations
- **docx-mailmerge** (Python) - Simple merge implementation
- **docx-templates** (JS) - Template filling with data

### Test Data
- Create Word template with merge fields
- Create CSV with sample data
- Test merge execution

## Success Criteria

- [ ] Can create template document with merge fields
- [ ] Can connect CSV data source
- [ ] Can execute merge and get single document with all records
- [ ] Can execute merge to multiple documents (one per record)
- [ ] Merge fields are replaced with correct data values
- [ ] Field names are case-insensitive
- [ ] Missing fields use empty string (or configurable default)
- [ ] Special characters in data are escaped/handled correctly
- [ ] Can merge to new document without modifying template
- [ ] Output documents are valid and open in Microsoft Word
- [ ] Roundtrip test passes

## Spec Capabilities to Update

- **wordprocessing-document**: Add mail merge requirements
- Create **wordprocessing-mailmerge**: NEW spec for merge engine

## Dependencies

- Existing field infrastructure
- CSV parsing (can use encoding/csv from stdlib)
- JSON parsing (encoding/json from stdlib)

## Out of Scope

- Database connectivity (ODBC, SQL) - users can implement DataSource interface
- Excel as data source (can add later)
- Conditional merge fields (IF, COMPARE) - Phase 2
- Nested merge regions (repeating sections) - Phase 2
- Email merge (sending emails) - separate feature

## Example Usage (Target API)

```go
// Create template document
doc, _ := wordprocessing.New("template.docx")
para := doc.MainDocumentPart().Document().Body().AppendParagraph()
run := para.AppendRun()
run.AppendFieldCode("MERGEFIELD FirstName")
para.AppendText(" ")
run2 := para.AppendRun()
run2.AppendFieldCode("MERGEFIELD LastName")

doc.SaveAs("letter-template.docx")

// Execute mail merge
template, _ := wordprocessing.Open("letter-template.docx")

// Option 1: From CSV
csvSource, _ := mailmerge.NewCSVDataSource("customers.csv")
merged := template.MailMerge().
    DataSource(csvSource).
    Execute()
merged.SaveAs("merged-letters.docx")

// Option 2: From in-memory data
data := []map[string]string{
    {"FirstName": "John", "LastName": "Doe"},
    {"FirstName": "Jane", "LastName": "Smith"},
}
mapSource := mailmerge.NewMapDataSource(data)
docs := template.MailMerge().
    DataSource(mapSource).
    ExecuteToDocuments()

for i, doc := range docs {
    doc.SaveAs(fmt.Sprintf("letter-%d.docx", i+1))
}
```

## Implementation Algorithm

```go
func (mm *MailMerge) Execute() (*Document, error) {
    // 1. Parse template for merge fields
    fields := mm.findMergeFields()
    
    // 2. Validate data source has required fields
    for _, field := range fields {
        if !mm.dataSource.HasField(field.Name) {
            return nil, fmt.Errorf("data source missing field: %s", field.Name)
        }
    }
    
    // 3. Create output document
    output := mm.template.Clone()
    firstRecord := true
    
    // 4. Iterate data source records
    for mm.dataSource.Next() {
        // Clone template section for this record
        section := mm.cloneTemplateSection()
        
        // Replace all merge fields with data
        for _, field := range fields {
            value, _ := mm.dataSource.Get(field.Name)
            mm.replaceField(section, field, value)
        }
        
        // Append to output
        if !firstRecord {
            section.PrependPageBreak()
        }
        output.AppendSection(section)
        firstRecord = false
    }
    
    return output, nil
}

func (mm *MailMerge) findMergeFields() []MergeField {
    var fields []MergeField
    
    // Walk document tree
    for para := range mm.template.Paragraphs() {
        for run := range para.Runs() {
            if run.FieldCode() != nil {
                code := run.FieldCode().Text()
                if strings.HasPrefix(code, "MERGEFIELD") {
                    fieldName := parseMergeFieldName(code)
                    fields = append(fields, MergeField{Name: fieldName, Run: run})
                }
            }
        }
    }
    
    return fields
}

func (mm *MailMerge) replaceField(doc *Document, field MergeField, value string) {
    // Find the run containing the field code
    run := field.Run
    
    // Replace field code with plain text value
    run.ClearFieldCode()
    run.AppendText(value)
}
```

## Data Source Implementations

### CSV Data Source
```go
type CSVDataSource struct {
    file     *os.File
    reader   *csv.Reader
    headers  []string
    current  []string
}

func (ds *CSVDataSource) Open() error {
    f, err := os.Open(ds.filename)
    if err != nil {
        return err
    }
    ds.file = f
    ds.reader = csv.NewReader(f)
    
    // Read header row
    headers, err := ds.reader.Read()
    if err != nil {
        return err
    }
    ds.headers = headers
    return nil
}

func (ds *CSVDataSource) Next() bool {
    record, err := ds.reader.Read()
    if err != nil {
        return false
    }
    ds.current = record
    return true
}

func (ds *CSVDataSource) Get(fieldName string) (string, error) {
    for i, header := range ds.headers {
        if strings.EqualFold(header, fieldName) {
            return ds.current[i], nil
        }
    }
    return "", fmt.Errorf("field not found: %s", fieldName)
}
```
