# Test Files Guide

Quick reference guide for choosing the right test files for different testing scenarios.

## Basic Functionality Testing

### Simple Documents (Good for Basic Read/Write Tests)
- `Plain.docx` - Minimal structure, ideal for basic tests
- `HelloWorld.docx` - Simple text content
- `HelloO14.docx` - Office 2014 format variant

### Templates
- `Document.dotx` - Word template file for template testing

## Advanced Feature Testing

### Document Properties
```go
// Testing various document property scenarios
"DocProps.docx"           // Standard properties
"MoreDocProps.docx"       // Extended properties
"NoDocProps.docx"         // Missing properties
"BadDocProps.docx"        // Malformed properties
"InvalidDocProps.docx"    // Invalid properties
"InvalidDocPropsct.docx"  // Invalid property count
```

### Content Features
```go
// Different content types and features
"Comments.docx"                        // Comment functionality
"Hyperlink.docx"                       // Hyperlink handling
"Notes.docx"                           // Notes/footnotes
"AnnotationRef.docx"                   // Annotation references
"svg.docx"                             // SVG graphics
"Data-Bound-Content-Controls.docx"    // Data binding
"simpleSdt.docx"                       // Structured document tags
```

### Mail Merge
```go
"mailmerge.docx" // Mail merge functionality
```

## Complex Document Testing

### Real-World Complex Documents
```go
// Large, feature-rich documents
"Complex01.docx"    // 828KB - Very comprehensive
"complex0.docx"     // 215KB - Complex variant
"complex2010.docx"  // 231KB - Office 2010 complex
"Document.docx"     // 799KB - Full-featured
```

### Office 2016 Features
```go
// Modern Office 2016 specific features
"Of16-01.docx" through "Of16-08.docx"  // Various O16 features
"Of16-09-UnknownElement.docx"          // Unknown element handling
"Of16-10-SymEx.docx"                   // Symbol extensions
```

## Error Handling & Edge Cases

### Intentional Errors
```go
"5Errors.docx"                   // Document with 5 errors
"UnknownElement.docx"            // Unknown XML elements
"EmptyRelationshipElement.docx"  // Empty relationship elements
```

### Format Compliance
```go
"Strict01.docx" // Strict Open XML compliance testing
```

### Markup Compatibility
```go
"mcdoc.docx"     // Markup compatibility document
"mcinleaf.docx"  // MC in leaf elements
```

## Example Test Usage

```go
// Basic read test
func TestBasicRead(t *testing.T) {
    doc, err := Open(filepath.Join("testdata", "Plain.docx"))
    // ...
}

// Complex document test
func TestComplexDocument(t *testing.T) {
    doc, err := Open(filepath.Join("testdata", "Complex01.docx"))
    // ...
}

// Error handling test
func TestErrorHandling(t *testing.T) {
    doc, err := Open(filepath.Join("testdata", "5Errors.docx"))
    // Should handle errors gracefully
}

// Properties test
func TestDocumentProperties(t *testing.T) {
    testCases := []struct{
        file string
        expectProps bool
    }{
        {"DocProps.docx", true},
        {"NoDocProps.docx", false},
        {"InvalidDocProps.docx", false},
    }
    // ...
}

// Office version compatibility
func TestOffice2016Features(t *testing.T) {
    files := []string{
        "Of16-01.docx",
        "Of16-02.docx",
        // ...
    }
    for _, file := range files {
        doc, err := Open(filepath.Join("testdata", file))
        // ...
    }
}
```

## Recommended Test Progression

1. **Start Simple**: Use `Plain.docx` and `HelloWorld.docx`
2. **Add Features**: Test `Comments.docx`, `Hyperlink.docx`
3. **Test Properties**: Use the DocProps series
4. **Go Complex**: Test with `Complex01.docx`, `Document.docx`
5. **Edge Cases**: Use error documents and strict compliance
6. **Modern Features**: Test Office 2016 files

## File Size Reference

| Category | Files | Size Range |
|----------|-------|------------|
| Simple | Plain, HelloWorld | 10-15KB |
| Standard | Most feature files | 10-20KB |
| Medium | Mail merge, May_12_04 | 15-80KB |
| Complex | complex0, complex2010 | 215-231KB |
| Very Complex | Complex01, Document | 799-828KB |

## Quick Stats

- Total test files: 39
- .docx files: 38
- .dotx files: 1
- Total size: ~3.9MB
- Office versions covered: 2007, 2010, 2014, 2016
- Format types: Transitional and Strict Open XML
