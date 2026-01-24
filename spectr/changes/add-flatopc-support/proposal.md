# Add FlatOPC Support

## Overview
Implement Flat OPC (single XML file) format support for Word, Excel, and PowerPoint documents, matching Open-XML-SDK FlatOpc functionality. Enables single-file representation without ZIP packaging.

## Motivation
Open-XML-SDK provides .FlatOpc methods for reading/writing documents as single XML files. Useful for version control, text-based diff, and debugging. Currently missing in goffice.

## Goals
- Implement FlatOpc loading for Word, Excel, PowerPoint
- Implement FlatOpc saving with proper namespace handling
- Support conversion between package and FlatOpc formats
- Preserve all parts, relationships, and content types in XML
- Support streaming FlatOpc generation for large documents
- Maintain compatibility with Office FlatOpc support

## Non-Goals
- Custom FlatOpc schema variations
- Optimization for FlatOpc as primary storage
- Incremental FlatOpc updates

## Dependencies
- Depends on: packaging, openxml core
- Related: add-package-clone-support

## Technical Approach

### FlatOPC Structure
```xml
<?xml version="1.0"?>
<pkg:package xmlns:pkg="...">
  <pkg:part pkg:name="/word/document.xml" pkg:contentType="...">
    <pkg:xmlData>
      <w:document>...</w:document>
    </pkg:xmlData>
  </pkg:part>
  <pkg:part pkg:name="/word/styles.xml">
    <pkg:xmlData>...</pkg:xmlData>
  </pkg:part>
</pkg:package>
```

### Load FlatOPC
```go
func OpenFlatOPC(path string) (*Document, error) {
    // Parse single XML file
    // Extract parts
    // Reconstruct package
}
```

### Save FlatOPC
```go
func (d *Document) SaveAsFlatOPC(path string) error {
    // Walk all parts
    // Serialize to single XML
}
```

## Estimated Effort
3 weeks (1 engineer)

## Priority
P0 - Critical for version control scenarios
