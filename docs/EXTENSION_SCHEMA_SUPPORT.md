# Extension Schema Support (Office 2010-2025)

This document describes how goffice handles extension schemas and modern Office features.

## Overview

Starting with Office 2007, Microsoft extended the Office Open XML format with additional capabilities in each new version (2010, 2013, 2016, 2019, 2021, 2022, 2023, 2024, 2025). goffice now supports these extensions, enabling you to:

- Read documents created in modern versions of Office without data loss
- Write documents using extension elements (e.g., Content Controls, Slicers, Modern Animations)
- Validate documents against specific Office version requirements
- Detect the minimum Office version required for a document
- Preserve unknown extensions for forward compatibility

## Office Version Support

goffice supports the following Office versions through the `FileFormatVersion` enum:

```go
type FileFormatVersion int

const (
    Office2007 FileFormatVersion = iota
    Office2010
    Office2013
    Office2016
    Office2019
    Office2021
    Office2022
    Office2023
    Office2024
    Office2025
    Microsoft365
)
```

## Extension Namespaces

Extension elements use versioned namespaces:

- **Word 2010-2024**: `w14`, `w15`, `w16`, etc.
- **Excel 2010-2024**: `x14`, `x15`, `x16`, etc.
- **PowerPoint 2010-2024**: `p14`, `p15`, `p16`, etc.
- **Drawing 2010-2024**: `a14`, `a15`, `a16`, etc.

These namespace prefixes are constants in the `openxml` package:

```go
// Word extensions
const W14Namespace = "http://schemas.microsoft.com/office/word/2010/wordml"
const W15Namespace = "http://schemas.microsoft.com/office/word/2012/wordml"
// ... more extensions

// Excel extensions
const X14Namespace = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"
// ... more extensions
```

## Version-Aware APIs

### Detecting Document Version

Use `DetectMinimumVersion()` to find the minimum Office version required:

```go
import "github.com/connerohnesorge/goffice/openxml/validation"

doc, _ := wordprocessing.Open("document.docx", false)
version := validation.DetectMinimumVersionForDocument(doc)

switch version {
case validation.Office2007:
    fmt.Println("Document uses only Office 2007 features")
case validation.Office2010:
    fmt.Println("Document requires Office 2010+")
case validation.Office2024:
    fmt.Println("Document requires Office 2024+")
}
```

### Version-Aware Validation

Validate documents against a target Office version:

```go
// Validate that document works with Office 2016
errors := validation.ValidateWithVersion(doc, validation.Office2016)

for _, err := range errors {
    if err.ErrorLevel == validation.ErrorVersionNotAvailable {
        fmt.Printf("Element %s not available in Office 2016: %s\n",
            err.ElementName, err.Description)
    }
}
```

### Checking Element Version Availability

Check if an element is available in a specific Office version:

```go
elem := doc.MainPart().Document().Body().Children()[0]
metadata := elem.(openxml.MetadataElement)

availableIn := metadata.AvailableInVersion()
if availableIn > validation.Office2016 {
    fmt.Printf("Element requires Office %d+\n", availableIn)
}
```

## AlternateContent (Office Version Fallbacks)

When Office saves documents in compatibility mode, it may include `AlternateContent` blocks with multiple versions:

```xml
<w:AlternateContent>
    <w:Choice Requires="w14">
        <!-- Office 2010+ version -->
        <w14:ContentControl .../>
    </w:Choice>
    <w:Fallback>
        <!-- Office 2007 fallback -->
        <w:StructuredDocumentTag .../>
    </w:Fallback>
</w:AlternateContent>
```

goffice provides utilities to work with AlternateContent:

```go
import "github.com/connerohnesorge/goffice/openxml"

elem := // ... a w:AlternateContent element

// Select content for a target version
selectedElem := elem.SelectContent(validation.Office2010)

// This returns the w:Choice content if available, otherwise the w:Fallback
```

## Unknown Element Preservation

If goffice encounters an element from an extension schema it doesn't recognize (e.g., from a future Office version), it preserves it as an `UnknownElement`:

```go
// When opening a document with unknown extensions
doc, _ := wordprocessing.Open("future-office.docx", false)

// Unknown elements are preserved in the document tree
body := doc.MainPart().Document().Body()
for _, child := range body.Children() {
    if unknown, ok := child.(*openxml.UnknownElement); ok {
        fmt.Printf("Preserved unknown element: %s\n", unknown.Name)
        // The raw XML is available
        xml := unknown.SerializeToString()
    }
}

// When you save the document, unknown elements are written back unchanged
doc.SaveAs("output.docx")
```

This ensures forward compatibility and prevents data loss when working with documents from newer Office versions.

## Common Extension Elements

### Word 2010+ Extensions (Content Controls)

Content Controls (SDT - Structured Document Tags) in the Office 2010 namespace:

```go
// Content Control (word 2010 w14 namespace)
cc := wordprocessing.NewContentControl()
// Set properties via w14 namespace
```

### Word 2012+ Extensions (Format Lock)

Format Lock allows protecting formatting:

```go
// Access Office 2012+ features via w15 namespace
// Format lock control is available as w15:FormattingLock
```

### Excel 2010+ Extensions (Slicers)

Excel Slicers for filtering:

```go
// Slicer elements in Excel 2010 x14 namespace
// Use x14:slicer for creating filter controls
```

### PowerPoint 2010+ Extensions (Modern Animations)

Modern animation models:

```go
// PowerPoint modern animations in p14+ namespaces
// Use p14:*, p15:*, p16:* for modern animation features
```

## Best Practices

### 1. Version Detection Before Feature Use

Always detect document version before using version-specific features:

```go
minVersion := validation.DetectMinimumVersionForDocument(doc)

if minVersion >= validation.Office2010 {
    // Safe to use Office 2010 features
    cc := wordprocessing.NewContentControl()
    body.AppendChild(cc)
}
```

### 2. Preserve Unknown Extensions

When you don't recognize an extension, let goffice preserve it:

```go
// Don't try to parse unknown elements
// goffice automatically wraps them in UnknownElement
doc, _ := wordprocessing.Open("document.docx", false)
doc.SaveAs("output.docx")
// Unknown elements are preserved
```

### 3. Validate Against Target Version

Always validate before targeting a specific Office version:

```go
targetVersion := validation.Office2016

errors := validation.ValidateWithVersion(doc, targetVersion)
if len(errors) > 0 {
    fmt.Printf("Document incompatible with Office 2016:\n")
    for _, err := range errors {
        fmt.Printf("  - %s\n", err.Description)
    }
    return
}

// Safe to save for Office 2016
doc.SaveAs("office2016.docx")
```

### 4. Use AlternateContent for Compatibility

When supporting multiple Office versions, use AlternateContent:

```go
// Create compatibility layer
ac := openxml.NewAlternateContent()

// Choice for Office 2010+
choice := openxml.NewChoice()
choice.Requires = "w14"
choice.AppendChild(wordprocessing.NewContentControl())
ac.AppendChild(choice)

// Fallback for Office 2007
fallback := openxml.NewFallback()
fallback.AppendChild(wordprocessing.NewStructuredDocumentTag())
ac.AppendChild(fallback)

// Use the alternation
body.AppendChild(ac)
```

## Troubleshooting

### "Element not available in Office X"

This validation error means the document uses features from a newer Office version than the target:

```go
// Check minimum version
minVersion := validation.DetectMinimumVersionForDocument(doc)
if minVersion > targetVersion {
    fmt.Printf("Document requires at least Office %v\n", minVersion)
}
```

### Unknown Elements in Reopened Document

If reopening a document shows `UnknownElement` entries:

1. Check the Office version the document was created in
2. Ensure extension schemas are loaded (automatic in goffice)
3. Unknown elements are intentionally preserved - this is not an error

### Missing Extension Elements

If expected elements are missing after generation:

1. Verify the generator loaded the extension schemas
2. Check the Open-XML-SDK schema files exist in `Open-XML-SDK/data/schemas/`
3. Regenerate with `go run ./cmd/gen-go-wordprocessing`

## Migration Guide

### From Previous goffice (Without Extensions)

No breaking changes! Extension support is fully backward compatible:

1. Regenerate elements: `go run ./cmd/gen-go-wordprocessing`
2. Code using only Office 2007 features continues to work unchanged
3. New extension elements are automatically available
4. Unknown element preservation prevents data loss

### Updating Code to Use Extensions

To use new extension features:

1. Detect document version with `DetectMinimumVersionForDocument()`
2. Validate with `ValidateWithVersion()` before using new features
3. Use new element types (now available after regeneration)
4. Test with Office to ensure compatibility

## References

- [ECMA-376 Office Open XML Standard](http://www.ecma-international.org/publications/standards/Ecma-376.htm)
- [ISO/IEC 29500 Open XML Format](https://www.iso.org/standard/71691.html)
- [Microsoft Office Open XML Developer Documentation](https://docs.microsoft.com/en-us/office/open-xml/)
- [Open-XML-SDK Repository](https://github.com/OfficeDev/Open-XML-SDK)
