# Word Processing Test Data

This directory contains comprehensive test files for integration testing of the Word document processing functionality.

## Source

All files were copied from the Open-XML-SDK test assets:
- Source: `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests.Assets/assets/TestFiles/`
- Files: 38 .docx files and 1 .dotx file

## File Categories

### Document Templates
- `Document.dotx` - Word template file

### Standard Documents
- `Plain.docx` - Simple plain document
- `HelloWorld.docx` - Basic hello world document
- `HelloO14.docx` - Office 2014 version document

### Complex Documents
- `Complex01.docx` - Complex document with advanced features (828KB)
- `complex0.docx` - Another complex document variant (215KB)
- `complex2010.docx` - Office 2010 complex document (231KB)
- `Document.docx` - Full-featured document (799KB)

### Feature-Specific Documents
- `Comments.docx` - Document with comments
- `Hyperlink.docx` - Document with hyperlinks
- `mailmerge.docx` - Mail merge document
- `Notes.docx` - Document with notes
- `AnnotationRef.docx` - Document with annotation references
- `simpleSdt.docx` - Document with structured document tags
- `Data-Bound-Content-Controls.docx` - Document with data-bound content controls
- `svg.docx` - Document with SVG graphics

### Office 2016 Feature Documents
- `Of16-01.docx` through `Of16-10-SymEx.docx` - Office 2016 specific features
- `Of16-09-UnknownElement.docx` - Tests unknown element handling

### Document Properties
- `DocProps.docx` - Document with properties
- `MoreDocProps.docx` - Extended document properties
- `NoDocProps.docx` - Document without properties
- `BadDocProps.docx` - Document with malformed properties
- `InvalidDocProps.docx` - Document with invalid properties
- `InvalidDocPropsct.docx` - Document with invalid property count

### Error Handling & Edge Cases
- `5Errors.docx` - Document with 5 intentional errors
- `UnknownElement.docx` - Document with unknown elements
- `EmptyRelationshipElement.docx` - Document with empty relationships
- `Strict01.docx` - Strict Open XML document

### Markup Compatibility
- `mcdoc.docx` - Markup compatibility document
- `mcinleaf.docx` - Markup compatibility in leaf elements

### Miscellaneous
- `May_12_04.docx` - Date-stamped test document

## Usage in Tests

These files are available for Go tests using the standard testdata directory convention. Access them in tests like:

```go
filepath.Join("testdata", "HelloWorld.docx")
```

## File Statistics

- Total files: 39
- .docx files: 38
- .dotx files: 1
- Total size: ~3.9MB
