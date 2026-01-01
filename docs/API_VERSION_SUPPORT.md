# Version Support APIs

This document describes the APIs for working with Office versions and extension schemas.

## FileFormatVersion Enum

The `FileFormatVersion` enum represents Office versions supported by goffice.

```go
package validation

type FileFormatVersion int

const (
    Office2007 FileFormatVersion = iota  // Original Office Open XML (2007)
    Office2010                           // Office 2010 extensions
    Office2013                           // Office 2013 extensions
    Office2016                           // Office 2016 extensions
    Office2019                           // Office 2019 extensions
    Office2021                           // Office 2021 extensions
    Office2022                           // Office 2022 extensions
    Office2023                           // Office 2023 extensions
    Office2024                           // Office 2024 extensions
    Office2025                           // Office 2025 extensions
    Microsoft365                         // Microsoft 365 (latest)
)
```

### String Representation

```go
version := validation.Office2016
fmt.Println(version.String()) // "Office2016"
```

## Version Detection API

### DetectMinimumVersionForDocument

Scan a Word document to find the minimum Office version required.

```go
func DetectMinimumVersionForDocument(doc *wordprocessing.Document) FileFormatVersion
```

**Example:**

```go
doc, _ := wordprocessing.Open("document.docx", false)
minVersion := validation.DetectMinimumVersionForDocument(doc)

if minVersion >= validation.Office2010 {
    fmt.Println("Document uses Office 2010+ features")
}
```

### DetectMinimumVersionForPackage

Scan a spreadsheet or presentation package to find minimum Office version.

```go
func DetectMinimumVersionForPackage(pkg *openxml.OpenXmlPackage) FileFormatVersion
```

**Example:**

```go
doc, _ := spreadsheet.Open("workbook.xlsx", false)
minVersion := validation.DetectMinimumVersionForPackage(doc)
```

## Validation APIs

### ValidateWithVersion

Validate a document against a specific Office version.

```go
func ValidateWithVersion(doc Document, version FileFormatVersion) []ValidationError
```

**Example:**

```go
errors := validation.ValidateWithVersion(doc, validation.Office2016)

for _, err := range errors {
    if err.ErrorLevel == validation.ErrorVersionNotAvailable {
        fmt.Printf("Element not available: %s\n", err.ElementName)
    }
}
```

### ValidationError Structure

```go
type ValidationError struct {
    ElementName   string
    Description   string
    ErrorLevel    ErrorLevel
    ElementPath   string  // XPath to element in document
}
```

**Error Levels:**

- `ErrorVersionNotAvailable`: Element requires a newer Office version
- `ErrorAlternateContentStructure`: AlternateContent block is malformed
- `ErrorValidationFailed`: Standard schema validation failed

## Element Metadata API

### MetadataElement Interface

Elements that support version metadata implement this interface:

```go
type MetadataElement interface {
    Metadata() *ElementMetadata
}

type ElementMetadata struct {
    Name                string
    AvailableSinceVersion FileFormatVersion
    Deprecated          bool
    DeprecatedSince     FileFormatVersion
}
```

**Example:**

```go
elem := doc.MainPart().Document().Body().Children()[0]

if metadata, ok := elem.(openxml.MetadataElement); ok {
    fmt.Printf("Element %s available since %s\n",
        metadata.Metadata().Name,
        metadata.Metadata().AvailableSinceVersion)
}
```

## AlternateContent APIs

### AlternateContent Element

The `AlternateContent` element allows specifying different content for different Office versions.

```go
type AlternateContent struct {
    XMLName xml.Name
    Choices []*Choice  // Version-specific content
    Fallback *Fallback // Fallback for unsupported versions
}

type Choice struct {
    Requires string      // Namespace requirement (e.g., "w14")
    Content  interface{} // Child element
}

type Fallback struct {
    Content interface{} // Fallback child element
}
```

### SelectContent Method

Select the appropriate content for a target version.

```go
func (ac *AlternateContent) SelectContent(version FileFormatVersion) interface{}
```

**Example:**

```go
ac := // ... AlternateContent element

// Get content for Office 2010
content := ac.SelectContent(validation.Office2010)

// Returns Choice content if available, otherwise Fallback
```

## UnknownElement API

### UnknownElement Structure

Preserves unknown elements from unrecognized extension schemas.

```go
type UnknownElement struct {
    Name       string
    Attributes map[string]string
    Children   []interface{}
    Content    string  // Raw text content
}
```

### Methods

```go
// Get the raw XML string
func (ue *UnknownElement) SerializeToString() string

// Check if this is a known element type
func (ue *UnknownElement) IsKnown() bool

// Get element name with namespace
func (ue *UnknownElement) QualifiedName() string
```

**Example:**

```go
// Preserve unknown extensions
for _, child := range body.Children() {
    if unknown, ok := child.(*openxml.UnknownElement); ok {
        fmt.Printf("Preserved: %s\n", unknown.QualifiedName())
        
        // Save the raw XML
        xmlStr := unknown.SerializeToString()
        fmt.Println(xmlStr)
    }
}

// When saving, unknown elements are preserved unchanged
doc.SaveAs("output.docx")
```

## Namespace Constants

Extension namespaces are available as constants in the `openxml` package:

```go
// Word extensions
const (
    W14Namespace = "http://schemas.microsoft.com/office/word/2010/wordml"
    W15Namespace = "http://schemas.microsoft.com/office/word/2012/wordml"
    W16Namespace = "http://schemas.microsoft.com/office/word/2015/wordml_cid"
    W16CNamespace = "http://schemas.microsoft.com/office/word/2015/wordml/cxnSp"
    // ... more W namespaces
)

// Excel extensions
const (
    X14Namespace = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"
    X15Namespace = "http://schemas.microsoft.com/office/spreadsheetml/2012/main"
    // ... more X namespaces
)

// PowerPoint extensions
const (
    P14Namespace = "http://schemas.microsoft.com/office/powerpoint/2010/main"
    P15Namespace = "http://schemas.microsoft.com/office/powerpoint/2012/main"
    // ... more P namespaces
)

// Drawing extensions
const (
    A14Namespace = "http://schemas.microsoft.com/office/drawing/2010/main"
    A15Namespace = "http://schemas.microsoft.com/office/drawing/2012/main"
    // ... more A namespaces
)
```

## Version Comparison

FileFormatVersion supports standard comparison operators:

```go
version := validation.Office2016
targetVersion := validation.Office2019

if version < targetVersion {
    fmt.Println("Version is older than 2019")
}

if version >= validation.Office2013 {
    fmt.Println("Version is 2013 or newer")
}
```

## Common Patterns

### Check if Feature is Available

```go
if detectedVersion >= validation.Office2010 {
    // Safe to use Office 2010 features
} else {
    // Use Office 2007 fallbacks only
}
```

### Conditional Element Creation

```go
func CreateContentControl(doc *wordprocessing.Document, text string) openxml.Element {
    body := doc.MainPart().Document().Body()
    minVersion := validation.DetectMinimumVersionForDocument(doc)
    
    if minVersion >= validation.Office2010 {
        // Use native Content Control (w14)
        cc := wordprocessing.NewContentControl()
        para := cc.AppendChild(wordprocessing.NewParagraph())
        para.AppendChild(wordprocessing.NewRun(text))
        return cc
    } else {
        // Use Office 2007 fallback (SDT)
        sdt := wordprocessing.NewStructuredDocumentTag()
        para := sdt.AppendChild(wordprocessing.NewParagraph())
        para.AppendChild(wordprocessing.NewRun(text))
        return sdt
    }
}
```

### Validate Before Saving

```go
func SaveForOfficeVersion(doc *wordprocessing.Document, targetVersion validation.FileFormatVersion, path string) error {
    errors := validation.ValidateWithVersion(doc, targetVersion)
    
    if len(errors) > 0 {
        return fmt.Errorf("document incompatible with target version: %v", errors)
    }
    
    return doc.SaveAs(path)
}
```

## Error Handling

Always check for validation errors when working with versions:

```go
errors := validation.ValidateWithVersion(doc, validation.Office2016)

if len(errors) > 0 {
    for _, err := range errors {
        log.Printf("Validation error at %s: %s", err.ElementPath, err.Description)
    }
    return fmt.Errorf("validation failed")
}
```

## Performance Considerations

- `DetectMinimumVersion()` walks the entire document tree - cache the result if checking multiple times
- Version checking is O(1) - safe to use in loops
- Unknown element preservation has minimal overhead

## See Also

- [Extension Schema Support Guide](EXTENSION_SCHEMA_SUPPORT.md)
- [Word Elements API Reference](API_WORDPROCESSING_ELEMENTS.md)
- [Excel Elements API Reference](API_SPREADSHEET_ELEMENTS.md)
- [PowerPoint Elements API Reference](API_PRESENTATION_ELEMENTS.md)
