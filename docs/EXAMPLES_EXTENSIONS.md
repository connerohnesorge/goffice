# Extension Schema Examples

This document provides practical examples of working with extension elements from Office 2010-2024.

## Word 2010+ Content Controls Example

Content Controls (SDT - Structured Document Tags) are Office 2010 features for structured data:

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/openxml/validation"
)

func CreateDocumentWithContentControl() error {
    // Create a new document
    doc, err := wordprocessing.New("output.docx", wordprocessing.DocTypeDocument)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    body := doc.MainPart().Document().Body()
    
    // Create a paragraph
    para := wordprocessing.NewParagraph()
    
    // Add a run with text
    run := para.AppendChild(wordprocessing.NewRun())
    run.AppendChild(wordprocessing.NewText("This is a content control:"))
    
    body.AppendChild(para)
    
    // Create Content Control (Office 2010 w14 namespace)
    cc := wordprocessing.NewContentControl()
    ccPara := cc.AppendChild(wordprocessing.NewParagraph())
    ccRun := ccPara.AppendChild(wordprocessing.NewRun())
    ccRun.AppendChild(wordprocessing.NewText("Controlled content"))
    
    body.AppendChild(cc)
    
    // Save
    return doc.SaveAs("output.docx")
}

func ReadDocumentWithContentControl() error {
    // Open document
    doc, err := wordprocessing.Open("document.docx", false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    // Check if document uses Office 2010+ features
    minVersion := validation.DetectMinimumVersionForDocument(doc)
    fmt.Printf("Document requires Office %v\n", minVersion)
    
    body := doc.MainPart().Document().Body()
    
    // Find Content Controls in document
    for _, child := range body.Children() {
        // Content Controls will be in the generated elements
        // Check element type and process accordingly
        fmt.Printf("Element type: %T\n", child)
    }
    
    return nil
}
```

## Excel 2010+ Slicers Example

Slicers provide filtering controls for pivot tables:

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice/spreadsheet"
    "github.com/connerohnesorge/goffice/openxml/validation"
)

func CreateSpreadsheetWithSlicer() error {
    // Create workbook
    doc, err := spreadsheet.Create("workbook.xlsx", spreadsheet.DocTypeWorkbook)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    // Slicers are Office 2010+ features (x14 namespace)
    // They would be added to worksheet parts
    
    // Save
    return doc.SaveAs("workbook.xlsx")
}

func ReadExcelExtensions() error {
    doc, err := spreadsheet.Open("workbook.xlsx", false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    // Detect version
    minVersion := validation.DetectMinimumVersionForPackage(doc)
    fmt.Printf("Workbook requires Excel %v\n", minVersion)
    
    // Iterate sheets
    for sheet := range doc.Sheets() {
        fmt.Printf("Sheet: %s\n", sheet.Name())
        
        // Sheet would contain extension elements if present
        // Process rows and cells as normal
        for row := range sheet.Rows() {
            fmt.Printf("  Row %d: %d cells\n", row.Index(), len(row.Cells()))
        }
    }
    
    return nil
}
```

## PowerPoint 2010+ Modern Animations Example

Modern animation models for Office 2010+:

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice/presentation"
    "github.com/connerohnesorge/goffice/openxml/validation"
)

func CreatePresentationWithModernAnimation() error {
    doc, err := presentation.New("presentation.pptx", presentation.DocTypePresentation)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    // Modern animations are PowerPoint 2010+ features (p14+ namespaces)
    // Animation elements would be added to slide parts
    
    return doc.SaveAs("presentation.pptx")
}

func ReadPresentationExtensions() error {
    doc, err := presentation.Open("presentation.pptx", false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    minVersion := validation.DetectMinimumVersionForPackage(doc)
    fmt.Printf("Presentation requires PowerPoint %v\n", minVersion)
    
    // Access slides
    pp := doc.PresentationPart()
    if pp == nil {
        return fmt.Errorf("no presentation part")
    }
    
    pres := pp.Presentation()
    if pres == nil {
        return fmt.Errorf("no presentation element")
    }
    
    // Iterate slides (would contain extension animations)
    if pres.SlideList() != nil {
        for _, slideRef := range pres.SlideList().SlideReferences() {
            fmt.Printf("Slide: %s\n", slideRef.ID())
        }
    }
    
    return nil
}
```

## Version Detection and Validation Example

Comprehensive example showing version handling:

```go
package main

import (
    "fmt"
    "log"
    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/openxml/validation"
)

func ProcessDocumentByVersion() error {
    // Open document
    doc, err := wordprocessing.Open("input.docx", true)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    // Detect version
    detectedVersion := validation.DetectMinimumVersionForDocument(doc)
    
    switch detectedVersion {
    case validation.Office2007:
        return processOffice2007Document(doc)
    case validation.Office2010:
        return processOffice2010Document(doc)
    case validation.Office2016:
        return processOffice2016Document(doc)
    default:
        return processModernDocument(doc, detectedVersion)
    }
}

func processOffice2007Document(doc *wordprocessing.Document) error {
    fmt.Println("Processing Office 2007 document")
    // Use only Office 2007 features
    body := doc.MainPart().Document().Body()
    
    for _, para := range body.Paragraphs() {
        fmt.Printf("Paragraph: %s\n", para.Text())
    }
    
    return doc.SaveAs("output-2007.docx")
}

func processOffice2010Document(doc *wordprocessing.Document) error {
    fmt.Println("Processing Office 2010+ document")
    // Can use Office 2010 features like Content Controls
    
    return doc.SaveAs("output-2010.docx")
}

func processOffice2016Document(doc *wordprocessing.Document) error {
    fmt.Println("Processing Office 2016+ document")
    // Can use Office 2016 features
    
    return doc.SaveAs("output-2016.docx")
}

func processModernDocument(doc *wordprocessing.Document, version validation.FileFormatVersion) error {
    fmt.Printf("Processing Office %v document\n", version)
    
    return doc.SaveAs("output-modern.docx")
}

// Validate before saving
func ValidateAndSave(doc *wordprocessing.Document, targetVersion validation.FileFormatVersion, path string) error {
    // Check compatibility
    errors := validation.ValidateWithVersion(doc, targetVersion)
    
    if len(errors) > 0 {
        fmt.Printf("Document incompatible with %v:\n", targetVersion)
        for _, err := range errors {
            fmt.Printf("  - %s: %s\n", err.ElementName, err.Description)
        }
        return fmt.Errorf("validation failed")
    }
    
    fmt.Printf("Document is compatible with %v\n", targetVersion)
    return doc.SaveAs(path)
}
```

## AlternateContent Compatibility Example

Working with AlternateContent blocks for version compatibility:

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice/openxml"
    "github.com/connerohnesorge/goffice/openxml/validation"
    "github.com/connerohnesorge/goffice/wordprocessing"
)

func CreateCompatibleContent() {
    body := // ... get document body
    
    // Create AlternateContent for Office 2010 compatibility
    ac := openxml.NewAlternateContent()
    
    // Choice for Office 2010+
    choice := openxml.NewChoice()
    choice.Requires = "w14"
    
    // Add Office 2010 feature
    cc := wordprocessing.NewContentControl()
    ccPara := cc.AppendChild(wordprocessing.NewParagraph())
    ccRun := ccPara.AppendChild(wordprocessing.NewRun())
    ccRun.AppendChild(wordprocessing.NewText("Office 2010+ feature"))
    choice.AppendChild(cc)
    
    ac.AppendChild(choice)
    
    // Fallback for Office 2007
    fallback := openxml.NewFallback()
    sdt := wordprocessing.NewStructuredDocumentTag()
    sdtPara := sdt.AppendChild(wordprocessing.NewParagraph())
    sdtRun := sdtPara.AppendChild(wordprocessing.NewRun())
    sdtRun.AppendChild(wordprocessing.NewText("Office 2007 fallback"))
    fallback.AppendChild(sdt)
    
    ac.AppendChild(fallback)
    
    // Add to document
    body.AppendChild(ac)
}

func ReadAlternateContent(elem interface{}) {
    if ac, ok := elem.(*openxml.AlternateContent); ok {
        // Select content for target version
        content := ac.SelectContent(validation.Office2010)
        
        if content != nil {
            fmt.Printf("Selected content for Office 2010: %T\n", content)
        }
    }
}
```

## Unknown Element Preservation Example

Handling future/unknown extensions:

```go
package main

import (
    "fmt"
    "github.com/connerohnesorge/goffice/openxml"
    "github.com/connerohnesorge/goffice/wordprocessing"
)

func HandleUnknownElements() error {
    // Open document that may contain unknown extensions
    doc, err := wordprocessing.Open("future-office.docx", false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    body := doc.MainPart().Document().Body()
    
    // Process all children
    for _, child := range body.Children() {
        if unknown, ok := child.(*openxml.UnknownElement); ok {
            // Handle unknown element
            fmt.Printf("Found unknown element: %s\n", unknown.QualifiedName())
            fmt.Printf("Preserving: %s\n", unknown.SerializeToString())
            
            // The unknown element will be preserved when saving
        } else {
            // Handle known elements normally
            fmt.Printf("Known element: %T\n", child)
        }
    }
    
    // Save - unknown elements are preserved unchanged
    return doc.SaveAs("output.docx")
}

func CountExtensions() error {
    doc, err := wordprocessing.Open("document.docx", false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    body := doc.MainPart().Document().Body()
    
    unknownCount := 0
    
    // Count unknown elements
    walkElements(body, func(elem interface{}) {
        if _, ok := elem.(*openxml.UnknownElement); ok {
            unknownCount++
        }
    })
    
    if unknownCount > 0 {
        fmt.Printf("Document contains %d unknown extension elements\n", unknownCount)
    }
    
    return nil
}

func walkElements(elem interface{}, fn func(interface{})) {
    fn(elem)
    
    // Recursively walk children if composite
    if composite, ok := elem.(openxml.CompositeElement); ok {
        for _, child := range composite.Children() {
            walkElements(child, fn)
        }
    }
}
```

## Complete End-to-End Example

Full example showing document processing workflow:

```go
package main

import (
    "fmt"
    "log"
    "github.com/connerohnesorge/goffice/wordprocessing"
    "github.com/connerohnesorge/goffice/openxml/validation"
)

func main() {
    if err := processDocument(); err != nil {
        log.Fatalf("Error: %v", err)
    }
}

func processDocument() error {
    // 1. Open document
    doc, err := wordprocessing.Open("input.docx", true)
    if err != nil {
        return fmt.Errorf("failed to open document: %w", err)
    }
    defer doc.Close()
    
    // 2. Detect version
    minVersion := validation.DetectMinimumVersionForDocument(doc)
    fmt.Printf("Input document requires: %v\n", minVersion)
    
    // 3. Validate for target version
    targetVersion := validation.Office2016
    validationErrors := validation.ValidateWithVersion(doc, targetVersion)
    
    if len(validationErrors) > 0 {
        fmt.Printf("Warning: Document uses features from %v\n", minVersion)
        fmt.Printf("Attempting to save for %v anyway...\n", targetVersion)
    }
    
    // 4. Process content
    body := doc.MainPart().Document().Body()
    paraCount := len(body.Paragraphs())
    fmt.Printf("Document has %d paragraphs\n", paraCount)
    
    // Add new content
    newPara := wordprocessing.NewParagraph()
    run := newPara.AppendChild(wordprocessing.NewRun())
    run.AppendChild(wordprocessing.NewText("Processed by goffice"))
    body.AppendChild(newPara)
    
    // 5. Save
    outputPath := "output.docx"
    if err := doc.SaveAs(outputPath); err != nil {
        return fmt.Errorf("failed to save document: %w", err)
    }
    
    fmt.Printf("Saved to %s\n", outputPath)
    return nil
}
```

## Troubleshooting Examples

### Check What Extensions Are Used

```go
func AnalyzeExtensions(docPath string) error {
    doc, err := wordprocessing.Open(docPath, false)
    if err != nil {
        return err
    }
    defer doc.Close()
    
    minVersion := validation.DetectMinimumVersionForDocument(doc)
    
    fmt.Printf("Document Analysis:\n")
    fmt.Printf("  Minimum Office Version: %v\n", minVersion)
    fmt.Printf("  Path: %s\n", docPath)
    
    body := doc.MainPart().Document().Body()
    elementCounts := make(map[string]int)
    
    walkElements(body, func(elem interface{}) {
        elemType := fmt.Sprintf("%T", elem)
        elementCounts[elemType]++
        
        // Check if element has version metadata
        if metaElem, ok := elem.(openxml.MetadataElement); ok {
            meta := metaElem.Metadata()
            fmt.Printf("  Found: %s (available since %v)\n",
                meta.Name, meta.AvailableSinceVersion)
        }
    })
    
    return nil
}
```

## References

- [Extension Schema Support Guide](EXTENSION_SCHEMA_SUPPORT.md)
- [API Version Support Reference](API_VERSION_SUPPORT.md)
- [Generator Development Guide](GENERATOR_DEVELOPMENT.md)
