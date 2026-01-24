# Add Document Builder API

## Overview
Implement fluent builder API for constructing OpenXML documents, matching Open-XML-SDK Builder namespace. Provides ergonomic alternative to direct element construction with proper package initialization and relationship management.

## Motivation
Open-XML-SDK provides OpenXmlPackageBuilder that simplifies document creation. Current goffice requires manual setup of all package components. Builder pattern dramatically improves developer experience.

## Goals
- Implement OpenXmlPackageBuilder pattern with method chaining
- Support creating documents from templates
- Auto-initialize required parts and relationships
- Provide builders for common document types (Word, Excel, PowerPoint)
- Support streaming construction for large documents
- Enable schema tracking and validation during construction

## Non-Goals
- Visual designer or GUI builder
- Code generation from existing documents
- Support for custom schemas

## Dependencies
- Depends on: packaging core, openxml framework
- Related: add-package-initialization-helpers, add-schema-tracking-feature

## Technical Approach

### Builder Pattern
```go
doc := wordprocessing.NewBuilder().
    WithTemplate("template.docx").
    WithStyles(defaultStyles).
    Build()

doc.AddParagraph().
    WithText("Hello World").
    WithStyle("Heading1")

doc.SaveAs("output.docx")
```

### Auto-Initialization
```go
type DocumentBuilder struct {
    pkg *packaging.Package
    parts map[string]Part
}

func (b *DocumentBuilder) Build() *Document {
    // Auto-create MainDocumentPart
    // Auto-create default styles
    // Auto-create settings
    // Wire up relationships
    return &Document{...}
}
```

## Estimated Effort
4 weeks (1 engineer)

## Priority
P0 - Dramatically improves DX
