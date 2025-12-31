# Mail Merge Package

The `mailmerge` package provides comprehensive mail merge functionality for Word documents in the goffice library. It allows you to create personalized documents by combining a template document with data from various sources.

## Features

- **Multiple Data Source Types**: CSV files, JSON files, and in-memory maps
- **Custom Data Sources**: Implement the `DataSource` interface for custom data providers
- **Two Field Formats**: Supports both SimpleField and ComplexField (Word's two field representations)
- **Advanced Field Types**: GREETINGLINE for personalized greetings, SKIPIF for conditional records
- **Field Formatting**: Supports formatting switches (\* Upper, \* Lower, \* FirstCap, \* Caps)
- **OOXML Metadata**: Automatic integration with Word's mail merge settings
- **Flexible Options**: Configure strict field validation and unused field removal

## Installation

```bash
go get github.com/connerohnesorge/goffice/wordprocessing/mailmerge
```

## Quick Start

### Basic CSV Mail Merge

```go
import (
    "log"
    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/wordprocessing/mailmerge"
)

func main() {
    // Open the template document
    doc, err := wordprocessing.Open("template.docx", true)
    if err != nil {
        log.Fatal(err)
    }

    // Create CSV data source
    ds := mailmerge.NewCSVDataSource("contacts.csv")

    // Execute merge
    result, err := mailmerge.New(doc).
        DataSource(ds).
        Execute()
    if err != nil {
        log.Fatal(err)
    }

    // Save the merged document
    result.SaveAs("merged.docx")
}
```

## Data Sources

### CSV Data Source

CSV files must have headers in the first row. Field names are case-insensitive.

**contacts.csv:**
```csv
FirstName,LastName,Email,Company
John,Doe,john@example.com,Acme Corp
Jane,Smith,jane@example.com,Tech Inc
```

**Usage:**
```go
ds := mailmerge.NewCSVDataSource("contacts.csv")
```

### JSON Data Source

JSON files must contain an array of objects. Field values are automatically converted to strings.

**contacts.json:**
```json
[
    {
        "FirstName": "John",
        "LastName": "Doe",
        "Email": "john@example.com",
        "Age": 30
    },
    {
        "FirstName": "Jane",
        "LastName": "Smith",
        "Email": "jane@example.com",
        "Age": 25
    }
]
```

**Usage:**
```go
ds := mailmerge.NewJSONDataSource("contacts.json")
```

### Map Data Source

For programmatic data generation or when you already have data in memory.

```go
records := []map[string]string{
    {
        "FirstName": "John",
        "LastName": "Doe",
        "Email": "john@example.com",
    },
    {
        "FirstName": "Jane",
        "LastName": "Smith",
        "Email": "jane@example.com",
    },
}

ds := mailmerge.NewMapDataSource(records)
```

### Custom Data Source

Implement the `DataSource` interface for custom data providers:

```go
type DataSource interface {
    Open() error
    Close() error
    Next() bool
    Fields() []string
    Get(fieldName string) (string, error)
}
```

## Field Types

Word documents can contain merge fields in two formats:

### SimpleField

Single XML element with field code as an attribute:

```xml
<w:fldSimple w:instr="MERGEFIELD FirstName \* MERGEFORMAT">
    <w:r><w:t>«FirstName»</w:t></w:r>
</w:fldSimple>
```

### ComplexField

Span of elements with begin/separate/end markers:

```xml
<w:r><w:fldChar w:fldCharType="begin"/></w:r>
<w:r><w:instrText>MERGEFIELD FirstName \* MERGEFORMAT</w:instrText></w:r>
<w:r><w:fldChar w:fldCharType="separate"/></w:r>
<w:r><w:t>«FirstName»</w:t></w:r>
<w:r><w:fldChar w:fldCharType="end"/></w:r>
```

Both formats are automatically detected and replaced during merge.

## Advanced Fields

### GREETINGLINE

Generates personalized greetings based on available name fields.

**Field in template:**
```
GREETINGLINE \f "Dear "
```

**Output examples:**
- With Title and LastName: "Dear Mr. Smith,"
- With FirstName and LastName: "Dear John Smith,"
- With FirstName only: "Dear John,"
- With no name fields: "Dear Sir or Madam,"

### SKIPIF

Conditionally skips records based on field value comparisons. Only works with `ExecuteToDocuments()`.

**Syntax:**
```
SKIPIF fieldname operator value
```

**Supported operators:**
- `=` or `==` - Equal (case-insensitive)
- `<>` or `!=` - Not equal (case-insensitive)
- `<` - Less than (numeric or lexicographic)
- `>` - Greater than (numeric or lexicographic)
- `<=` - Less than or equal
- `>=` - Greater than or equal

**Examples:**
```
SKIPIF Status = "Inactive"       // Skip inactive records
SKIPIF Age < 18                   // Skip minors
SKIPIF Country <> "USA"           // Skip non-USA records
```

## Merge Options

Configure merge behavior with `MergeOptions`:

```go
opts := &mailmerge.MergeOptions{
    StrictFields:       true,  // Error if field not found (default: false)
    RemoveUnusedFields: true,  // Remove unreplaced fields (default: true)
}

result, err := mailmerge.New(doc).
    Options(opts).
    DataSource(ds).
    Execute()
```

### StrictFields

- **`true`**: Merge fails with error if a MERGEFIELD references a field not in the data source
- **`false`** (default): Missing fields are replaced with empty strings

### RemoveUnusedFields

- **`true`** (default): Merge fields that couldn't be replaced are removed from output
- **`false`**: Field codes remain in the document if they couldn't be replaced

## Execution Modes

### Execute() - Single Merged Document

Combines all records into a single document with page breaks between records:

```go
result, err := mailmerge.New(doc).
    DataSource(ds).
    Execute()
```

**Use when:** You want one document containing all personalized records (e.g., batch printing).

### ExecuteToDocuments() - Individual Documents

Creates a separate document for each record:

```go
documents, err := mailmerge.New(doc).
    DataSource(ds).
    ExecuteToDocuments()

for i, doc := range documents {
    doc.SaveAs(fmt.Sprintf("letter_%d.docx", i))
}
```

**Use when:** You need individual files for each record (e.g., email attachments, separate PDFs).

**Note:** SKIPIF fields only work with `ExecuteToDocuments()`.

## Field Formatting Switches

The `\*` switch controls text formatting:

| Switch | Description | Example |
|--------|-------------|---------|
| `\* UPPER` | All uppercase | JOHN DOE |
| `\* LOWER` | All lowercase | john doe |
| `\* FIRSTCAP` | First letter capitalized | John doe |
| `\* CAPS` | Title case (first letter of each word) | John Doe |

**In template:**
```
MERGEFIELD FirstName \* UPPER
MERGEFIELD LastName \* FIRSTCAP
```

## OOXML Metadata Integration

The package automatically reads mail merge metadata from the document's `settings.xml` file (the `w:mailMerge` element). This provides compatibility with templates created in Microsoft Word.

### Field Name Mappings

Word's mail merge UI allows mapping document field names to different data source column names. The package respects these mappings automatically.

**Example:**
- Document has MERGEFIELD "FirstName"
- Settings.xml maps "FirstName" → "fname" (data source column)
- Package automatically uses "fname" when reading from data source

This happens transparently - you don't need to do anything special.

### Metadata Fields

The `OoxmlMailMergeMetadata` struct contains:
- **DataType**: Type of data source (textFile, database, etc.)
- **ConnectString**: Connection information
- **Query**: Query or file path
- **FieldMappings**: Document field name → data source field name mappings

## Error Handling

Common errors and how to handle them:

### No Data Source
```go
result, err := mm.Execute()
if err == mailmerge.ErrNoDataSource {
    log.Fatal("Must configure a data source before executing")
}
```

### Missing Fields (StrictFields = true)
```go
opts := &mailmerge.MergeOptions{StrictFields: true}
result, err := mailmerge.New(doc).Options(opts).DataSource(ds).Execute()
if err != nil {
    // Error message includes list of missing fields
    log.Printf("Missing fields: %v", err)
}
```

### No Records
```go
result, err := mm.Execute()
if err == mailmerge.ErrNoRecords {
    log.Fatal("Data source is empty")
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/wordprocessing/mailmerge"
)

func main() {
    // Open template
    doc, err := wordprocessing.Open("invitation_template.docx", true)
    if err != nil {
        log.Fatal(err)
    }

    // Prepare data
    guests := []map[string]string{
        {
            "Title":     "Mr.",
            "FirstName": "John",
            "LastName":  "Doe",
            "Email":     "john@example.com",
            "Status":    "VIP",
        },
        {
            "FirstName": "Jane",
            "LastName":  "Smith",
            "Email":     "jane@example.com",
            "Status":    "Regular",
        },
    }

    // Configure merge
    ds := mailmerge.NewMapDataSource(guests)
    opts := &mailmerge.MergeOptions{
        StrictFields:       false, // Allow missing fields
        RemoveUnusedFields: true,  // Clean output
    }

    // Execute merge - individual documents
    documents, err := mailmerge.New(doc).
        Options(opts).
        DataSource(ds).
        ExecuteToDocuments()
    if err != nil {
        log.Fatal(err)
    }

    // Save each document
    for i, doc := range documents {
        filename := fmt.Sprintf("invitation_%d.docx", i+1)
        if err := doc.SaveAs(filename); err != nil {
            log.Printf("Failed to save %s: %v", filename, err)
        } else {
            fmt.Printf("Created %s\n", filename)
        }
    }
}
```

## API Reference

For detailed API documentation, see the [godoc](https://pkg.go.dev/github.com/connerohnesorge/goffice/wordprocessing/mailmerge).

## Testing

Run the test suite:

```bash
go test ./wordprocessing/mailmerge/... -v
```

Run specific tests:

```bash
go test ./wordprocessing/mailmerge -run TestCSVDataSource -v
go test ./wordprocessing/mailmerge -run TestGreetingLine -v
```

## License

Copyright 2024 goffice authors. All rights reserved.
Use of this source code is governed by a BSD-style license.
