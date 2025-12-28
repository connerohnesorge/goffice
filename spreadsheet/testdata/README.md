# Spreadsheet Test Data

This directory contains comprehensive test files for integration testing of the Excel spreadsheet processing functionality.

## Source

All files were copied from the Open-XML-SDK test assets:
- Source: `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests.Assets/assets/TestFiles/`
- Files: 13 .xlsx files and 1 .xltx file

## File Categories

### Spreadsheet Templates
- `Spreadsheet.xltx` - Excel template file (8.6KB)

### Standard Spreadsheets
- `Spreadsheet.xlsx` - Full-featured spreadsheet file (51KB)
- `basicspreadsheet.xlsx` - Basic spreadsheet for simple tests (26KB)

### Complex & Feature-Rich
- `Complex01.xlsx` - Complex spreadsheet with advanced features (114KB)
- `excel14.xlsx` - Excel 2014 format features (87KB)

### Comments & Collaboration
- `Comments.xlsx` - Spreadsheet with comments (11KB)
- `Revision_NameCommentChange.xlsx` - Revision tracking with name/comment changes (8.1KB)

### Calculation Features
- `missingcalcchainpart.xlsx` - Test for missing calculation chain part (51KB)

### Extension Lists
- `extlst.xlsx` - Extension list testing (15KB)

### Media & References
- `Youtube.xlsx` - Media reference (YouTube) testing (23KB)

### Markup Compatibility
- `MCExecl.xlsx` - Markup compatibility Excel (6.8KB)

### VML Drawing
- `vmldrawingroot.xlsx` - VML drawing root testing (9.4KB)

### Malformed URIs
- `malformed_uri.xlsx` - Malformed URI handling (10KB)
- `malformed_uri_long.xlsx` - Long malformed URI handling (11KB)

## Usage in Tests

These files are available for Go tests using the standard testdata directory convention. Access them in tests like:

```go
filepath.Join("testdata", "Spreadsheet.xlsx")
```

## File Statistics

- Total files: 14
- .xlsx files: 13
- .xltx files: 1
- Total size: ~456KB
- Largest file: Complex01.xlsx (114KB)
- Smallest file: MCExecl.xlsx (6.8KB)

## Feature Coverage

The test files cover:
- Basic spreadsheet structure
- Worksheets and cells
- Formulas and calculations
- Comments and revisions
- Templates (.xltx)
- Complex formulas and features
- Extension lists
- VML drawings
- Media references
- Markup compatibility
- Malformed URI handling
- Calculation chain parts
- Excel 2014 features
