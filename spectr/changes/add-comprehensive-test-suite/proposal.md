# Add Comprehensive Test Suite Matching Open-XML-SDK

## Overview
Create comprehensive test suite matching the breadth and depth of Open-XML-SDK test coverage. This includes conformance tests, roundtrip tests, validation tests, and real-world document tests ported from the SDK's 10,000+ test cases.

## Motivation
Current goffice test coverage (~60-70%) is focused on core functionality and happy paths. The Open-XML-SDK has 15+ years of accumulated test cases covering:
- Edge cases and boundary conditions
- Office version compatibility (2007-365)
- Malformed and corrupt documents
- Performance and memory limits
- Interoperability scenarios

Achieving test parity is essential for:
1. **Production readiness**: Confidence that goffice handles all real-world scenarios
2. **Regression prevention**: Catch bugs before they reach users
3. **Feature parity proof**: Demonstrate equivalent behavior to SDK
4. **Documentation**: Tests serve as working examples

## Goals

### Test Coverage Targets
- **Code coverage**: >80% line coverage across all packages
- **Branch coverage**: >70% branch coverage for critical paths
- **Test count**: ~8,000-10,000 test cases (parity with SDK)
- **Test types**: Unit, integration, conformance, performance, fuzzing

### Test Categories to Implement

#### 1. Core Functionality Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Tests/`
- OpenXmlDomTest: Element creation, manipulation, serialization
- SimpleTypes: All simple type conversions and validation
- ofapiTest: API surface tests
- Extensions: Extension method tests

#### 2. Packaging Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Packaging.Tests/`
- Package creation, opening, closing
- Part addition, removal, modification
- Relationship management
- Content type handling
- Stream handling and disposal

#### 3. Validation Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Tests/Validation/`
- Schema validation tests
- Semantic validation tests (once implemented)
- Version-specific validation
- Error reporting and recovery

#### 4. Document-Specific Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Tests/Wordprocessing/`, `Spreadsheet/`, etc.
- Document structure tests
- Style handling
- Formatting preservation
- Complex elements (tables, charts, SmartArt)

#### 5. Conformance Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Tests/ConformanceTest/`
- ECMA-376 conformance
- ISO/IEC 29500 conformance
- Office compatibility

#### 6. Framework Tests (port from SDK)
**Source**: `DocumentFormat.OpenXml.Framework.Tests/`
- Feature system tests
- Metadata tests
- Schema tests

#### 7. LINQ Tests (once implemented)
**Source**: `DocumentFormat.OpenXml.Linq.Tests/`
- LINQ query tests
- XElement conversion tests

#### 8. Features Tests (once implemented)
**Source**: `DocumentFormat.OpenXml.Framework.Features.Tests/`
- Element events
- Paragraph IDs
- Random number generation

### Additional Test Types

#### 9. Roundtrip Tests
- Create document → Save → Load → Verify
- Modify document → Save → Load → Verify changes
- Test across Office versions

#### 10. Real-World Document Tests
**Source**: SDK test assets + curated documents
- Documents from Office 2007, 2010, 2013, 2016, 2019, 2021, 365
- Documents from LibreOffice, Google Docs exports
- Complex real-world documents (resumes, reports, spreadsheets, presentations)
- Documents with errors/corruption (ensure graceful handling)

#### 11. Performance Tests
- Benchmark document operations
- Memory usage profiling
- Streaming vs DOM performance
- Large document handling (>100MB)

#### 12. Fuzz Tests
- Random document generation
- Mutation-based fuzzing of valid documents
- Structure-aware fuzzing

#### 13. PDF Rendering Tests
- Pixel-perfect comparison with Office PDF export
- Font rendering accuracy
- Layout fidelity
- Complex element rendering (charts, tables, images)

## Non-Goals
- 100% line coverage (some error paths are hard to trigger)
- GUI or end-user application tests
- Network-dependent tests (unless mocked)
- Platform-specific tests (focus on cross-platform Go code)

## Dependencies
- **Requires**: All feature implementations (tests follow features)
- **Requires**: Test asset repository (SDK test documents)
- **Blocks**: v1.0.0 release
- **Enables**: Production readiness certification

## Technical Approach

### 1. Test Organization

```
goffice/
├── openxml/
│   ├── element_test.go          # Core element tests
│   ├── features/
│   │   └── features_test.go     # Feature system tests
│   ├── validation/
│   │   └── validation_test.go   # Validation tests
│   └── testdata/                # Test fixtures
├── packaging/
│   ├── package_test.go          # Package tests
│   └── testdata/
├── wordprocessing/
│   ├── document_test.go         # Word API tests
│   ├── elements/
│   │   └── paragraph_test.go    # Element-specific tests
│   ├── conformance_test.go      # ECMA-376 conformance
│   └── testdata/
├── spreadsheet/
│   └── ... (similar structure)
├── presentation/
│   └── ... (similar structure)
├── pdf/
│   ├── rendering_test.go        # PDF rendering tests
│   ├── comparison_test.go       # Pixel comparison tests
│   └── testdata/
└── tests/
    ├── e2e/                     # End-to-end tests
    │   ├── roundtrip_test.go
    │   └── realworld_test.go
    ├── conformance/             # Standards conformance
    ├── performance/             # Benchmarks
    ├── fuzzing/                 # Fuzz tests
    └── testdata/                # Shared test assets
```

### 2. Test Porting Strategy

**Automated Port (where possible)**:
```bash
# Tool to convert C# tests to Go
go run ./cmd/port-sdk-tests \
    -sdk ../Open-XML-SDK/test \
    -output ./tests/ported \
    -package DocumentFormat.OpenXml.Tests
```

**Manual Port (for complex tests)**:
- Study C# test logic
- Rewrite using Go idioms (table-driven tests)
- Adapt to goffice API differences
- Add Go-specific assertions

### 3. Test Data Management

**SDK Test Assets**:
- Download from SDK repository: `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests.Assets/`
- Store in goffice: `testdata/sdk-assets/`
- Organize by Office version and document type

**Custom Test Documents**:
- Generate programmatically where possible
- Store minimal binary files (use Git LFS for large files)
- Document provenance (how document was created)

### 4. Test Utilities

```go
package testutil

// Load test document
func LoadTestDocument(t *testing.T, path string) *wordprocessing.Document {
    doc, err := wordprocessing.Open(path)
    require.NoError(t, err)
    t.Cleanup(func() { doc.Close() })
    return doc
}

// Compare documents
func AssertDocumentsEqual(t *testing.T, expected, actual *wordprocessing.Document, opts ...CompareOption) {
    comparer := NewDocumentComparer(opts...)
    differences := comparer.Compare(expected, actual)
    assert.Empty(t, differences, "Documents differ:\n%s", formatDifferences(differences))
}

// Compare PDFs
func AssertPDFsMatch(t *testing.T, expectedPDF, actualPDF string, tolerance float64) {
    differ := pdf.NewVisualDiffer(tolerance)
    diff, err := differ.Compare(expectedPDF, actualPDF)
    require.NoError(t, err)
    assert.Less(t, diff.MaxPixelDifference, tolerance,
        "PDF rendering differs at %v (%.2f%% different)",
        diff.MaxDifferenceLocation, diff.PercentDifferent)
}

// Roundtrip helper
func RoundtripDocument(t *testing.T, doc *wordprocessing.Document) *wordprocessing.Document {
    // Save to temp file
    tmp := filepath.Join(t.TempDir(), "roundtrip.docx")
    require.NoError(t, doc.SaveAs(tmp))
    
    // Reload
    return LoadTestDocument(t, tmp)
}
```

### 5. CI/CD Integration

**.github/workflows/tests.yml**:
```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    strategy:
      matrix:
        go-version: [1.22, 1.23, 1.24, 1.25]
        os: [ubuntu-latest, windows-latest, macos-latest]
    
    runs-on: ${{ matrix.os }}
    
    steps:
      - uses: actions/checkout@v4
        with:
          lfs: true
      
      - uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
      
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          file: ./coverage.out
  
  conformance:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run conformance tests
        run: go test -v ./tests/conformance/...
  
  performance:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run benchmarks
        run: go test -bench=. -benchmem ./...
      - name: Store results
        uses: benchmark-action/github-action-benchmark@v1
        with:
          tool: 'go'
          output-file-path: benchmark.txt
```

## Implementation Plan

### Phase 1: Infrastructure (2 weeks)
- [ ] Set up test organization structure
- [ ] Create test utility package
- [ ] Set up CI/CD test matrix
- [ ] Download and organize SDK test assets
- [ ] Create test porting tool (C# → Go converter)

### Phase 2: Core Tests (4 weeks)
- [ ] Port OpenXmlDomTest (~1,000 tests)
- [ ] Port SimpleTypes tests (~500 tests)
- [ ] Port ofapiTest (~300 tests)
- [ ] Port Extension tests (~200 tests)

### Phase 3: Packaging Tests (2 weeks)
- [ ] Port packaging tests (~800 tests)
- [ ] Add goffice-specific packaging tests
- [ ] Add stream handling tests

### Phase 4: Validation Tests (3 weeks)
- [ ] Port schema validation tests (~1,000 tests)
- [ ] Add semantic validation tests (once implemented)
- [ ] Add version-specific validation tests

### Phase 5: Document-Specific Tests (6 weeks)
- [ ] Port Wordprocessing tests (~2,000 tests)
- [ ] Port Spreadsheet tests (~2,000 tests)
- [ ] Port Presentation tests (~1,000 tests)
- [ ] Port chart/drawing tests (~500 tests)

### Phase 6: Conformance Tests (2 weeks)
- [ ] Port ConformanceTest suite
- [ ] Port IsoStrictTest suite
- [ ] Add ECMA-376 conformance tests
- [ ] Add ISO 29500 conformance tests

### Phase 7: Framework Tests (2 weeks)
- [ ] Port Framework.Tests (~400 tests)
- [ ] Port Features.Tests (~200 tests)
- [ ] Port Linq.Tests (~300 tests)

### Phase 8: Additional Tests (4 weeks)
- [ ] Create roundtrip test suite (~500 tests)
- [ ] Create real-world document test suite (~200 documents)
- [ ] Create performance benchmark suite (~100 benchmarks)
- [ ] Create fuzz test suite

### Phase 9: PDF Tests (3 weeks)
- [ ] Port PDF rendering tests
- [ ] Create pixel comparison test suite (~500 comparisons)
- [ ] Add font rendering tests
- [ ] Add layout fidelity tests

### Phase 10: Integration & Polish (2 weeks)
- [ ] Fix test failures
- [ ] Achieve >80% code coverage
- [ ] Document test organization
- [ ] Create test writing guide

**Total Estimate: 30 weeks (1 engineer) OR 15 weeks (2 engineers)**

## Test Metrics & Reporting

### Coverage Metrics
```bash
# Run with coverage
go test -coverprofile=coverage.out ./...

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html

# Coverage summary
go tool cover -func=coverage.out | grep total
```

### Test Reports
```bash
# Verbose test output
go test -v ./... | tee test-output.log

# JSON output for tooling
go test -json ./... > test-results.json

# Benchmark comparison
go test -bench=. -benchmem ./... > benchmark-new.txt
benchcmp benchmark-old.txt benchmark-new.txt
```

### Quality Gates
- All tests must pass before merge
- Coverage cannot decrease (ratcheting)
- No new code without tests
- Benchmarks within 10% of baseline

## Test Examples

### Table-Driven Test (Go Idiom)
```go
func TestParagraphAlignment(t *testing.T) {
    tests := []struct {
        name      string
        alignment wordprocessing.JustificationValue
        xmlOutput string
    }{
        {
            name:      "Left alignment",
            alignment: wordprocessing.JustificationLeft,
            xmlOutput: `<w:jc w:val="left"/>`,
        },
        {
            name:      "Center alignment",
            alignment: wordprocessing.JustificationCenter,
            xmlOutput: `<w:jc w:val="center"/>`,
        },
        // ... more cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            para := wordprocessing.NewParagraph()
            para.Properties().SetAlignment(tt.alignment)
            
            xml := para.ToXML()
            assert.Contains(t, xml, tt.xmlOutput)
        })
    }
}
```

### Roundtrip Test
```go
func TestRoundtripComplexDocument(t *testing.T) {
    // Load original
    original := testutil.LoadTestDocument(t, "testdata/complex-resume.docx")
    
    // Roundtrip
    roundtripped := testutil.RoundtripDocument(t, original)
    
    // Compare
    testutil.AssertDocumentsEqual(t, original, roundtripped,
        testutil.IgnoreWhitespace(),
        testutil.IgnoreAttributeOrder(),
    )
}
```

### Conformance Test
```go
func TestECMA376Conformance(t *testing.T) {
    doc := wordprocessing.NewDocument()
    // ... build document
    
    validator := validation.NewValidator(
        validation.WithStandard(validation.ECMA376_Edition5),
    )
    
    errors := validator.Validate(doc)
    assert.Empty(t, errors, "Document violates ECMA-376")
}
```

### Performance Benchmark
```go
func BenchmarkDocumentCreation(b *testing.B) {
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        doc := wordprocessing.NewDocument()
        for j := 0; j < 100; j++ {
            para := doc.Body().AppendParagraph()
            para.AppendText("Benchmark paragraph")
        }
        doc.Close()
    }
}
```

## Risks & Mitigations

### Risk: Test porting effort is massive (10,000+ tests)
**Impact:** CRITICAL - Resource constraints

**Mitigation:**
- Automate test conversion where possible (tool-assisted porting)
- Prioritize by feature importance (P0 features get tests first)
- Incremental approach (port tests as features are implemented)
- Community contributions (open source leverage)

### Risk: Test failures reveal implementation bugs
**Impact:** HIGH - Delays feature completion

**Mitigation:**
- That's the point! Better to find bugs during testing
- Fix bugs as discovered, don't defer
- Track known failures, prioritize fixes

### Risk: Test maintenance burden
**Impact:** MEDIUM - Tests become stale or brittle

**Mitigation:**
- Keep tests simple and focused
- Use table-driven tests (easy to add cases)
- Generate tests from metadata where possible
- Regular test review and cleanup

### Risk: Test execution time becomes prohibitive
**Impact:** MEDIUM - Slow CI/CD pipeline

**Mitigation:**
- Parallel test execution (`go test -parallel`)
- Short tests in PR pipeline, full suite nightly
- Cache test results
- Skip slow tests in development (tag-based)

### Risk: Test data storage (large binary files)
**Impact:** MEDIUM - Repository bloat

**Mitigation:**
- Use Git LFS for binary test files
- Generate test documents programmatically where possible
- Host large test data externally, download on demand

## Success Criteria

1. ✅ >80% code coverage across all packages
2. ✅ >8,000 test cases (parity with SDK)
3. ✅ All critical paths have test coverage
4. ✅ Tests pass on Linux, macOS, Windows
5. ✅ Tests pass on Go 1.22, 1.23, 1.24, 1.25
6. ✅ Benchmark baseline established for performance tracking
7. ✅ Conformance tests pass for ECMA-376 and ISO 29500
8. ✅ Real-world documents from Office 2007-365 handled correctly

## References

- **Open-XML-SDK Tests**: `Open-XML-SDK/test/`
- **SDK Test Assets**: `Open-XML-SDK/test/DocumentFormat.OpenXml.Tests.Assets/`
- **Go Testing**: https://go.dev/doc/tutorial/add-a-test
- **Table-Driven Tests**: https://go.dev/wiki/TableDrivenTests
- **Benchmarking**: https://pkg.go.dev/testing#hdr-Benchmarks

## Future Enhancements

Beyond initial test suite:
1. **Mutation testing**: Verify tests catch bugs (e.g., stryker-go)
2. **Property-based testing**: Generate random valid documents (gopter)
3. **Visual regression testing**: Screenshot-based UI tests (if GUI added)
4. **Load testing**: Concurrent document generation at scale
5. **Compatibility matrix**: Test against all Office versions automatically

---

**Estimated Effort:** 30 weeks (1 engineer) OR 15 weeks (2 engineers)  
**Priority:** P0 - Cannot claim parity without test coverage  
**Status:** Draft  
**Last Updated:** 2026-01-24
