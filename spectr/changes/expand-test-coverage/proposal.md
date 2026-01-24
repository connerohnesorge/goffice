# Change: Expand Test Coverage to Match Open-XML-SDK Reference Implementation

## Why
goffice has good test coverage for core functionality but lacks comprehensive test scenarios that exist in Open-XML-SDK. To achieve feature parity and ensure reliability:
- Roundtrip tests for all new features (create → save → load → verify)
- Interoperability tests with Microsoft Office 2007 through 365 documents
- Edge case and boundary tests for all data types and structures
- Performance benchmarks for large documents
- Stress tests for deeply nested structures
- Fuzzing tests for malformed documents
- Compatibility tests with LibreOffice, Google Docs, and other readers
- Regression tests for reported issues
- Table-driven comprehensive scenario tests
- Visual regression tests for PDF rendering

## What Changes
- Add comprehensive roundtrip tests for all features (word, excel, presentation)
- Add interoperability tests with Office version-specific documents
- Add edge case tests (empty elements, maximum nesting, extreme values)
- Add performance benchmarks (memory, time, throughput)
- Add fuzzing tests for malformed input
- Add visual regression tests for PDF rendering
- Add cross-version compatibility tests (ECMA-376 1st, 2nd, 3rd editions)
- Add document corruption recovery tests
- Add concurrent access tests (thread-safety)
- Add encoding/decoding tests for special characters
- Add formula parsing and recalculation tests
- Add OLE object handling tests
- BREAKING: Test infrastructure may require additional test fixtures

## Impact
- Affected specs: All specs (framework, wordprocessing, spreadsheet, presentation, drawingml, pdf)
- Affected code: All test files (*_test.go), test infrastructure, test data
- New dependencies: Testing frameworks may be required (table-driven test generators)
- Test coverage: Aim for >90% code coverage across all packages
