# Test Files Guide

Quick reference guide for choosing the right test files for different testing scenarios.

## Basic Functionality Testing

### Simple Spreadsheets (Good for Basic Read/Write Tests)
- `basicspreadsheet.xlsx` - Basic spreadsheet for simple tests (26KB)
- `MCExecl.xlsx` - Minimal markup compatibility Excel (6.8KB)
- `Spreadsheet.xlsx` - Standard spreadsheet file (51KB)

### Templates
- `Spreadsheet.xltx` - Excel template file for template testing (8.6KB)

## Advanced Feature Testing

### Complex Features
```go
// Testing complex formulas and features
"Complex01.xlsx"     // Complex spreadsheet with advanced features (114KB)
"excel14.xlsx"       // Excel 2014 format features (87KB)
```

### Comments & Revisions
```go
// Comments and collaboration features
"Comments.xlsx"                     // Spreadsheet with comments
"Revision_NameCommentChange.xlsx"   // Revision tracking with name/comment changes
```

### Calculations
```go
// Formula and calculation testing
"missingcalcchainpart.xlsx"  // Test for missing calculation chain part
```

### Extension Lists
```go
// Extension list testing
"extlst.xlsx"  // Extension list features
```

### Media & References
```go
// Media references and embedded content
"Youtube.xlsx"  // YouTube media reference handling
```

### VML Drawing
```go
// VML drawing testing
"vmldrawingroot.xlsx"  // VML drawing root features
```

## Error Handling Testing

### Malformed URIs
```go
// Testing malformed URI handling
"malformed_uri.xlsx"       // Standard malformed URI
"malformed_uri_long.xlsx"  // Long malformed URI
```

## Markup Compatibility Testing

```go
"MCExecl.xlsx"  // Markup compatibility Excel
```

## Example Test Usage

```go
// Basic read test
func TestBasicRead(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Spreadsheet.xlsx"))
    // ...
}

// Template test
func TestTemplateRead(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Spreadsheet.xltx"))
    // ...
}

// Comments test
func TestComments(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Comments.xlsx"))
    // Test comment reading and writing
}

// Complex features test
func TestComplexFeatures(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Complex01.xlsx"))
    // Test advanced formulas, formatting, etc.
}

// Calculation chain test
func TestCalculationChain(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "missingcalcchainpart.xlsx"))
    // Test handling of missing calculation chain
}

// Extension list test
func TestExtensionList(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "extlst.xlsx"))
    // Test extension list features
}

// VML drawing test
func TestVMLDrawing(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "vmldrawingroot.xlsx"))
    // Test VML drawing features
}

// Media reference test
func TestMediaReferences(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Youtube.xlsx"))
    // Test media reference handling
}

// Malformed URI test
func TestMalformedURI(t *testing.T) {
    testCases := []string{
        "malformed_uri.xlsx",
        "malformed_uri_long.xlsx",
    }
    for _, file := range testCases {
        wb, err := Open(filepath.Join("testdata", file))
        // Should handle malformed URIs gracefully
    }
}

// Revision tracking test
func TestRevisions(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "Revision_NameCommentChange.xlsx"))
    // Test revision tracking features
}

// Excel 2014 features test
func TestExcel2014Features(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "excel14.xlsx"))
    // Test Excel 2014 specific features
}

// Markup compatibility test
func TestMarkupCompatibility(t *testing.T) {
    wb, err := Open(filepath.Join("testdata", "MCExecl.xlsx"))
    // Test markup compatibility
}
```

## Recommended Test Progression

1. **Start Simple**: Use `MCExecl.xlsx` and `basicspreadsheet.xlsx`
2. **Add Features**: Test `Comments.xlsx`, `extlst.xlsx`
3. **Test Calculations**: Use `missingcalcchainpart.xlsx`
4. **Test Drawing**: Use `vmldrawingroot.xlsx`
5. **Go Complex**: Test with `Complex01.xlsx`, `excel14.xlsx`
6. **Test Media**: Use `Youtube.xlsx`
7. **Error Handling**: Test `malformed_uri.xlsx`, `malformed_uri_long.xlsx`
8. **Revisions**: Test `Revision_NameCommentChange.xlsx`

## File Size Reference

| Category | Files | Size Range |
|----------|-------|------------|
| Small | MCExecl, Spreadsheet.xltx, Revision | 6.8-8.6KB |
| Medium | malformed_uri, Comments, vmldrawingroot | 9.4-11KB |
| Large | extlst, Youtube, basicspreadsheet | 15-26KB |
| Very Large | Spreadsheet, missingcalcchainpart | 51KB |
| Extra Large | excel14, Complex01 | 87-114KB |

## Quick Stats

- Total test files: 14
- .xlsx files: 13
- .xltx files: 1
- Total size: ~456KB
- Largest file: Complex01.xlsx (114KB)

## Feature Coverage Matrix

| Feature | Test Files |
|---------|------------|
| Comments | Comments.xlsx |
| Revisions | Revision_NameCommentChange.xlsx |
| Calculations | missingcalcchainpart.xlsx |
| Extension Lists | extlst.xlsx |
| Media References | Youtube.xlsx |
| Templates | Spreadsheet.xltx |
| VML Drawing | vmldrawingroot.xlsx |
| Markup Compatibility | MCExecl.xlsx |
| Malformed URIs | malformed_uri.xlsx, malformed_uri_long.xlsx |
| Excel 2014 | excel14.xlsx |
| Complex Features | Complex01.xlsx |
| Basic Structure | basicspreadsheet.xlsx, Spreadsheet.xlsx |

## Special Considerations

### Malformed URIs
- `malformed_uri.xlsx`, `malformed_uri_long.xlsx` - Test error handling for invalid URIs
- Should gracefully handle and recover from malformed URI references

### Missing Parts
- `missingcalcchainpart.xlsx` - Test handling of missing calculation chain parts
- Good for testing robustness and error recovery

### VML Drawing
- `vmldrawingroot.xlsx` - Test legacy VML drawing features
- Important for backward compatibility with older Excel files

### Media References
- `Youtube.xlsx` - Test external media reference handling
- Good for testing media link preservation and handling

### Complex Features
- `Complex01.xlsx` - Comprehensive test for advanced Excel features
- Use for stress testing and feature completeness validation

### Excel Version Compatibility
- `excel14.xlsx` - Test Excel 2014 specific features
- Ensure compatibility across different Excel versions

## Testing Best Practices

### Start with Basic Files
Begin your test suite with simple files to validate core functionality:
```go
testFiles := []string{
    "MCExecl.xlsx",
    "basicspreadsheet.xlsx",
}
```

### Progressive Complexity
Gradually add more complex test cases:
```go
testFiles := []string{
    "basicspreadsheet.xlsx",  // Basic
    "Comments.xlsx",          // + Comments
    "extlst.xlsx",            // + Extensions
    "Complex01.xlsx",         // Full complexity
}
```

### Error Handling Tests
Always include error handling tests:
```go
errorTestFiles := []string{
    "malformed_uri.xlsx",
    "malformed_uri_long.xlsx",
    "missingcalcchainpart.xlsx",
}
```

### Feature-Specific Tests
Group tests by feature:
```go
// Drawing tests
drawingFiles := []string{
    "vmldrawingroot.xlsx",
}

// Media tests
mediaFiles := []string{
    "Youtube.xlsx",
}

// Collaboration tests
collabFiles := []string{
    "Comments.xlsx",
    "Revision_NameCommentChange.xlsx",
}
```

## Common Test Scenarios

### Roundtrip Testing
Test reading and writing files without data loss:
```go
func TestRoundtrip(t *testing.T) {
    testFiles := []string{
        "basicspreadsheet.xlsx",
        "Spreadsheet.xlsx",
        "Complex01.xlsx",
    }
    for _, file := range testFiles {
        // Read, write, read again, compare
    }
}
```

### Template Processing
Test template file handling:
```go
func TestTemplateProcessing(t *testing.T) {
    tmpl, err := Open(filepath.Join("testdata", "Spreadsheet.xltx"))
    // Process template and create new workbook
}
```

### Data Extraction
Test extracting specific data:
```go
func TestDataExtraction(t *testing.T) {
    files := map[string]string{
        "Comments.xlsx": "comments",
        "Youtube.xlsx": "media",
        "extlst.xlsx": "extensions",
    }
    for file, feature := range files {
        // Extract and validate specific features
    }
}
```
