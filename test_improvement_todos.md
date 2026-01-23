# XML Output Test Improvement Todos

## Task: Find and remove or improve minimal XML output tests

### Files with minimal XML output tests to review:

#### Priority 1 (drawingml files):
- [x] drawingml/effects_test.go - TestSoftEdgeXmlOutput, TestReflectionXmlOutput
- [x] drawingml/geometry_test.go - Multiple tests with strings.Contains

#### Priority 2 (spreadsheet files):
- [x] spreadsheet/elements/worksheet_test.go
- [ ] spreadsheet/elements/workbook_test.go
- [ ] spreadsheet/elements/stylesheet_test.go
- [ ] spreadsheet/elements/drawing_test.go
- [ ] spreadsheet/elements/data_validation_test.go
- [ ] spreadsheet/elements/comments_test.go
- [ ] spreadsheet/elements/conditional_formatting_test.go
- [ ] spreadsheet/elements/filter_test.go
- [ ] spreadsheet/elements/cell_test.go
- [ ] spreadsheet/elements/shared_strings_test.go
- [ ] spreadsheet/elements/threaded_comments_test.go

#### Priority 3 (wordprocessing files):
- [ ] wordprocessing/elements/elements_test.go
- [ ] wordprocessing/elements/revision_test.go
- [ ] wordprocessing/elements/drawing_test.go
- [ ] wordprocessing/integration_test.go
- [ ] wordprocessing/alternate_content_integration_test.go
- [ ] wordprocessing/mailmerge/field_replacer_test.go
- [ ] wordprocessing/mailmerge/integration_test.go
- [ ] wordprocessing/mailmerge/data_source_test.go
- [ ] wordprocessing/mailmerge/advanced_fields_test.go
- [ ] wordprocessing/comprehensive_integration_test.go
- [ ] wordprocessing/roundtrip_test.go
- [ ] wordprocessing/builder_test.go

#### Priority 4 (presentation files):
- [ ] presentation/parts/smartart_integration_test.go
- [ ] presentation/comprehensive_integration_test.go
- [ ] presentation/alternate_content_integration_test.go

#### Priority 5 (other files):
- [ ] openxml/openxml_test.go
- [ ] openxml/alternate_content_test.go
- [ ] packaging/extended_properties_test.go
- [ ] pdf files (multiple)

### Analysis of problematic patterns found:

1. **drawingml/effects_test.go**:
   - TestSoftEdgeXmlOutput: Only checks if XML contains "softEdge" and "rad=\"120000\""
   - TestReflectionXmlOutput: Checks for element name and attributes but not structure

2. **drawingml/geometry_test.go**:
   - Multiple tests checking only for element presence (avLst, gd, prstGeom, etc.)
   - Tests that verify XML contains specific strings but not proper structure
   - Lines 44-49, 59, 74-78, 85-90, 133-135, 185-191, 198-201, etc.

3. **spreadsheet/elements/worksheet_test.go**:
   - Lines 672-694: Checking for element names only (worksheet, dimension, sheetViews, etc.)

### Approach for each file:
1. Read the file and identify tests that use strings.Contains with XML
2. Determine if the test provides value:
   - If it's just checking element names without structure → Remove
   - If it's checking attributes but not structure → Consider improving
   - If it's part of a roundtrip test → Keep but improve validation
3. For tests to improve:
   - Replace strings.Contains with proper XML parsing/validation
   - Check actual structure, not just string presence
   - Use proper XML comparison or schema validation
4. Run tests after changes to ensure nothing breaks

### Implementation Plan:
1. Start with drawingml/effects_test.go (smaller file, easier to fix)
2. Then drawingml/geometry_test.go (more tests but similar pattern)
3. Then spreadsheet files
4. Then wordprocessing files
5. Finally presentation and other files