# Implementation Tasks: .NET Cross-Runtime Testing

**Related Documents**:
- Requirements and rationale: See `proposal.md`
- Technical design details: See `design.md`
- E2E testing scenarios: See `specs/e2e-testing/spec.md`
- Nix environment scenarios: See `specs/nix-environment/spec.md`

## Phase 1: Foundation & MVP (Word Documents Only)

### 1. Nix Environment Setup
- [x] 1.1 Add dotnet-sdk_9 to flake.nix devShells.default packages
- [x] 1.2 Add libreoffice to flake.nix for Office→PDF conversion
- [ ] 1.3 Add ghostscript to flake.nix (already may exist)
- [ ] 1.4 Create test-go script in flake.nix scripts section
- [ ] 1.5 Create test-dotnet script in flake.nix scripts section
- [ ] 1.6 Create test-e2e script in flake.nix scripts section
- [ ] 1.7 Create test-all script in flake.nix scripts section
- [ ] 1.8 Create update-baselines script in flake.nix scripts section
- [ ] 1.9 Run `nix flake update` to update flake.lock
- [ ] 1.10 Test `nix develop` enters shell successfully with all tools

### 2. Directory Structure
- [ ] 2.1 Create `tests/e2e/` directory
- [ ] 2.2 Create `tests/e2e/framework/` for test framework code
- [ ] 2.3 Create `tests/e2e/bridges/go/` for Go bridge
- [ ] 2.4 Create `tests/e2e/bridges/csharp/DocxBridge/` for C# bridge project
- [ ] 2.5 Create `tests/e2e/scenarios/wordprocessing/` for Word test scenarios
- [ ] 2.6 Create `tests/e2e/baselines/wordprocessing/` for baseline documents
- [ ] 2.7 Create `tests/e2e/comparison/` for comparison logic
- [ ] 2.8 Create `tests/e2e/output/go/` and `tests/e2e/output/dotnet/` (gitignored)
- [ ] 2.9 Create `tests/e2e/reports/` (gitignored)
- [ ] 2.10 Create `tests/e2e/cmd/e2e-runner/` for CLI entry point
- [ ] 2.11 Create `tests/e2e/.gitignore` (ignore output/, reports/)
- [ ] 2.12 Create `tests/e2e/go.mod` as separate Go module

### 3. Scenario Definition System
- [ ] 3.1 Create `tests/e2e/framework/scenario.go` with TestScenario struct
- [ ] 3.2 Implement LoadScenario() to parse YAML files
- [ ] 3.3 Implement LoadScenariosFromDirectory() for batch loading
- [ ] 3.4 Create ToleranceConfig struct with default values
- [ ] 3.5 Write unit tests for scenario parsing
  - [ ] 3.5.1 Create `framework/scenario_test.go` for framework self-testing
  - [ ] 3.5.2 Test valid YAML parsing
  - [ ] 3.5.3 Test invalid YAML error handling
  - [ ] 3.5.4 Test default tolerance values
  - [ ] 3.5.5 Test tolerance overrides
- [ ] 3.6 Create 5 basic Word scenario YAML files:
  - [ ] 3.6.1 `basic-paragraph.yaml` - single paragraph with plain text
  - [ ] 3.6.2 `formatted-text.yaml` - bold, italic, color formatting
  - [ ] 3.6.3 `multiple-paragraphs.yaml` - multiple paragraphs with spacing
  - [ ] 3.6.4 `simple-table.yaml` - 2x3 table with text
  - [ ] 3.6.5 `mixed-content.yaml` - paragraphs + table
- [ ] 3.7 Create operation-to-API mapping specification document
  - [ ] 3.7.1 Document how each YAML operation maps to goffice API calls
  - [ ] 3.7.2 Document how each YAML operation maps to Open-XML-SDK API calls
  - [ ] 3.7.3 Document how scenario properties map to method parameters
  - [ ] 3.7.4 Include code examples for complex operations (nested runs, tables)

### 4. Go Bridge Implementation
- [ ] 4.1 Create `tests/e2e/bridges/go/bridge.go` with Bridge struct
- [ ] 4.2 Implement NewBridge() constructor
- [ ] 4.3 Implement Execute() to dispatch by document type
- [ ] 4.4 Implement executeWordprocessing() for Word documents
- [ ] 4.5 Implement addParagraph() operation handler
- [ ] 4.6 Implement applyParagraphProperties() for paragraph formatting
- [ ] 4.7 Implement applyRunProperties() for run formatting
- [ ] 4.8 Implement addTable() operation handler
- [ ] 4.9 Write unit tests for Go bridge with mock scenarios
- [ ] 4.10 Test Go bridge generates valid .docx files

### 5. C# Bridge Implementation
- [ ] 5.1 Create `tests/e2e/bridges/csharp/DocxBridge/DocxBridge.csproj`
- [ ] 5.2 Add Open-XML-SDK NuGet package reference
- [ ] 5.3 Add YamlDotNet NuGet package reference for YAML parsing
- [ ] 5.4 Add System.CommandLine NuGet package for CLI
- [ ] 5.5 Create `Program.cs` with CLI entry point
- [ ] 5.6 Create `TestScenario.cs` with C# model classes matching YAML structure
- [ ] 5.7 Create `WordprocessingBridge.cs` with scenario executor
- [ ] 5.8 Implement Execute() method for document generation
- [ ] 5.9 Implement AddParagraph() operation handler
- [ ] 5.10 Implement ApplyParagraphProperties() method
- [ ] 5.11 Implement ApplyRunProperties() method
- [ ] 5.12 Implement AddTable() operation handler
- [ ] 5.13 Create `bridges/csharp/build.sh` build script
- [ ] 5.14 Test C# bridge builds successfully with `dotnet build`
- [ ] 5.15 Test C# bridge generates valid .docx from scenario
- [ ] 5.16 Manually verify C# bridge output opens in Microsoft Word without errors
- [ ] 5.17 Validate C# bridge output using Open-XML-SDK OpenSettings.Validate()
- [ ] 5.18 Test C# bridge handles invalid scenario YAML gracefully

### 6. Level 1: XML Structure Comparison
- [ ] 6.1 Create `tests/e2e/comparison/xml.go`
- [ ] 6.2 Implement XMLDiffResult struct with mismatch lists
- [ ] 6.3 Implement CompareXMLStructure() function
- [ ] 6.4 Implement extractXMLPart() to extract XML from DOCX ZIP
- [ ] 6.5 Implement parseXMLToTree() for XML parsing to tree
- [ ] 6.6 Implement XMLNode struct
- [ ] 6.7 Implement compareNodes() recursive comparison
- [ ] 6.8 Handle element name mismatches
- [ ] 6.9 Handle attribute mismatches
- [ ] 6.10 Handle text content mismatches
- [ ] 6.11 Handle child count mismatches
- [ ] 6.12 Respect XMLAttributeOrderSensitive tolerance
- [ ] 6.13 Write unit tests for XML comparison
  - [ ] 6.13.1 Create `comparison/xml_test.go` for self-testing
  - [ ] 6.13.2 Test identical documents (should pass)
  - [ ] 6.13.3 Test element mismatches
  - [ ] 6.13.4 Test attribute differences
  - [ ] 6.13.5 Test tolerance handling

### 7. Test Execution Framework
- [ ] 7.1 Create `tests/e2e/framework/executor.go`
- [ ] 7.2 Implement Executor struct
- [ ] 7.3 Implement NewExecutor() constructor
- [ ] 7.4 Implement ExecuteScenario() orchestration method
- [ ] 7.5 Implement executeGoBridge() to call Go bridge
- [ ] 7.6 Implement executeCSharpBridge() to call C# bridge via subprocess
- [ ] 7.7 Create TestResult struct with all result fields
- [ ] 7.8 Create ComparisonResult struct (Level 1 only for now)
- [ ] 7.9 Implement CompareDocuments() function (Level 1 only)
- [ ] 7.10 Define ComparisonResult aggregation logic:
  - [ ] 7.10.1 Specify pass/fail criteria (XML must match exactly, tolerances for visual)
  - [ ] 7.10.2 Define how XML, Binary, and Visual results combine into single verdict
  - [ ] 7.10.3 Document priority rules (e.g., XML mismatch = fail even if visual matches)
- [ ] 7.11 Write unit tests for executor with mock bridges

### 8. CLI Test Runner
- [ ] 8.1 Create `tests/e2e/cmd/e2e-runner/main.go`
- [ ] 8.2 Implement command-line flags (scenarios, output, baselines, etc.)
- [ ] 8.3 Implement scenario loading
- [ ] 8.4 Implement test execution loop
- [ ] 8.5 Implement printSummary() for test results
- [ ] 8.6 Handle exit codes (0 for pass, 1 for fail)
- [ ] 8.7 Test CLI runs successfully
- [ ] 8.8 Implement cleanup mechanism for temporary files
- [ ] 8.9 Add --clean flag to remove old test outputs

### 9. Basic Reporting
- [ ] 9.1 Create `tests/e2e/framework/reporter.go`
- [ ] 9.2 Implement console output for test results
- [ ] 9.3 Implement JSON report generation (for CI)
- [ ] 9.4 Create simple text report with pass/fail summary
- [ ] 9.5 Implement error handling for missing scenario files
- [ ] 9.6 Implement error handling for corrupt or invalid documents
- [ ] 9.7 Test framework behavior with intentionally broken scenarios

### 10. MVP Testing & Validation
- [ ] 10.1 Run `test-e2e` script from Nix shell
- [ ] 10.2 Verify all 5 scenarios execute successfully
- [ ] 10.3 Verify Go and C# both generate .docx files
- [ ] 10.4 Verify XML comparison runs and reports differences
- [ ] 10.5 Manually inspect generated documents in Microsoft Word
- [ ] 10.6 Fix any bugs discovered during testing:
  - [ ] 10.6.1 Track each bug with issue number or description
  - [ ] 10.6.2 Prioritize bugs (critical/major/minor)
  - [ ] 10.6.3 Fix critical bugs before moving to Phase 2
  - [ ] 10.6.4 Create regression test for each fixed bug
- [ ] 10.7 Document usage in `tests/e2e/README.md`
- [ ] 10.8 Create initial README documentation (don't wait until Phase 23)
- [ ] 10.9 Test Nix environment works on Linux
- [ ] 10.10 Test Nix environment works on macOS

---

## Phase 2: Complete Comparison Pipeline

### 11. Level 2: Binary Content Comparison
- [ ] 11.1 Create `tests/e2e/comparison/binary.go`
- [ ] 11.2 Implement BinaryDiffResult struct
- [ ] 11.3 Implement CompareBinaryContent() function
- [ ] 11.4 Build part maps from ZIP files
- [ ] 11.5 Detect missing parts
- [ ] 11.6 Detect extra parts
- [ ] 11.7 Implement PartDifference struct
- [ ] 11.8 Implement compareBinaryBytes() for non-XML parts
- [ ] 11.9 Integrate binary comparison into CompareDocuments()
- [ ] 11.10 Write unit tests for binary comparison

### 12. Level 3: Visual Rendering Comparison
- [ ] 12.1 Create `tests/e2e/comparison/visual.go`
- [ ] 12.2 Implement VisualDiffResult struct
- [ ] 12.3 Implement convertOfficeToPDF() using LibreOffice
- [ ] 12.4 Test LibreOffice headless conversion works in Nix shell
- [ ] 12.5 Implement CompareVisualRendering() function
- [ ] 12.6 Reuse `pdf/comparison` package for PNG conversion
- [ ] 12.7 Reuse `pdf/comparison` package for pixel diff
- [ ] 12.8 Implement diff image generation
- [ ] 12.9 Handle multi-page documents
- [ ] 12.10 Integrate visual comparison into CompareDocuments()
- [ ] 12.11 Write unit tests for visual comparison

### 13. Tolerance Configuration
- [ ] 13.1 Create `tests/e2e/comparison/tolerance.go`
- [ ] 13.2 Implement ToleranceConfig struct
- [ ] 13.3 Implement default tolerances
- [ ] 13.4 Implement per-scenario tolerance overrides
- [ ] 13.5 Document tolerance configuration in README

### 14. HTML Report Generation
- [ ] 14.1 Create `tests/e2e/framework/html_reporter.go`
- [ ] 14.2 Implement GenerateHTMLReport() function
- [ ] 14.3 Create HTML template with CSS
- [ ] 14.4 Implement side-by-side document comparison view
- [ ] 14.5 Embed diff images in HTML report
- [ ] 14.6 Add interactive features (zoom, toggle overlay):
  - [ ] 14.6.1 Implement JavaScript zoom functionality OR
  - [ ] 14.6.2 Mark as "Future Enhancement" if deferred
- [ ] 14.7 Generate summary statistics (pass rate, pixel match %, DiffPercentage, MaxColorDelta, AvgColorDelta)
- [ ] 14.8 Test HTML report in browser

### 15. Baseline Management
- [ ] 14.9 Decide baseline storage strategy (git vs git-lfs) based on file sizes
- [ ] 15.1 Create `tests/e2e/scripts/generate-baselines.sh`
- [ ] 15.2 Implement baseline generation using C# bridge only
- [ ] 15.3 Store baselines in `tests/e2e/baselines/wordprocessing/`
- [ ] 15.4 Create `tests/e2e/scripts/update-baselines.sh`
- [ ] 15.5 Implement baseline update workflow
- [ ] 15.6 Add baseline comparison mode to executor
- [ ] 15.7 Document baseline workflow in README

---

## Phase 3: Expand Coverage (Excel & PowerPoint)

### 16. Excel (Spreadsheet) Support
- [ ] 16.1 Create `tests/e2e/scenarios/spreadsheet/` directory
- [ ] 16.2 Create `tests/e2e/bridges/go/spreadsheet.go`
- [ ] 16.3 Implement executeSpreadsheet() in Go bridge
- [ ] 16.4 Create `SpreadsheetBridge.cs` in C# bridge project
- [ ] 16.5 Implement Excel operations:
  - [ ] 16.5.1 create_workbook
  - [ ] 16.5.2 add_sheet
  - [ ] 16.5.3 set_cell_value
  - [ ] 16.5.4 set_cell_formula
  - [ ] 16.5.5 merge_cells
- [ ] 16.6 Create 10 Excel scenario YAML files
- [ ] 16.7 Test Excel scenarios end-to-end

### 17. PowerPoint (Presentation) Support
- [ ] 17.1 Create `tests/e2e/scenarios/presentation/` directory
- [ ] 17.2 Create `tests/e2e/bridges/go/presentation.go`
- [ ] 17.3 Implement executePresentation() in Go bridge
- [ ] 17.4 Create `PresentationBridge.cs` in C# bridge project
- [ ] 17.5 Implement PowerPoint operations:
  - [ ] 17.5.1 create_presentation
  - [ ] 17.5.2 add_slide
  - [ ] 17.5.3 add_shape
  - [ ] 17.5.4 add_text_box
  - [ ] 17.5.5 set_shape_properties
- [ ] 17.6 Create 10 PowerPoint scenario YAML files
- [ ] 17.7 Test PowerPoint scenarios end-to-end

### 18. Expand Word Scenario Coverage
- [ ] 18.1 Create advanced Word scenarios:
  - [ ] 18.1.1 headers-footers.yaml
  - [ ] 18.1.2 lists-numbered.yaml
  - [ ] 18.1.3 lists-bulleted.yaml
  - [ ] 18.1.4 styles.yaml
  - [ ] 18.1.5 hyperlinks.yaml
  - [ ] 18.1.6 images.yaml
  - [ ] 18.1.7 footnotes-endnotes.yaml
  - [ ] 18.1.8 comments.yaml
  - [ ] 18.1.9 tables-complex.yaml (merged cells, nested)
  - [ ] 18.1.10 sections.yaml (multi-section with different headers)
- [ ] 18.2 Test all advanced scenarios

### 19. Edge Cases & Regression Tests
- [ ] 19.1 Create edge case scenarios:
  - [ ] 19.1.1 empty-document.yaml
  - [ ] 19.1.2 max-table-size.yaml
  - [ ] 19.1.3 unicode-text.yaml
  - [ ] 19.1.4 special-characters.yaml
  - [ ] 19.1.5 large-document.yaml (performance test)
- [ ] 19.2 Add regression tests for historical bugs

---

## Phase 4: CI/CD Integration

### 20. GitHub Actions Workflow
- [ ] 20.1 Create `.github/workflows/e2e-tests.yml`
- [ ] 20.2 Configure workflow to run on PR and main branch
- [ ] 20.3 Set up Nix installation in CI
- [ ] 20.4 Run `nix develop --command test-e2e`
- [ ] 20.5 Upload test reports as artifacts
- [ ] 20.6 Upload diff images as artifacts for failed tests
- [ ] 20.7 Configure workflow to fail on test failures
- [ ] 20.8 Test workflow runs successfully

### 21. PR Integration
- [ ] 21.1 Add GitHub Actions bot comment with test results
- [ ] 21.2 Include pass/fail summary
- [ ] 21.3 Link to uploaded artifacts
- [ ] 21.4 Show diff images inline (if possible)

### 22. Baseline Update Automation
- [ ] 22.1 Create `.github/workflows/update-baselines.yml`
- [ ] 22.2 Trigger on manual dispatch or schedule
- [ ] 22.3 Run `update-baselines` script
- [ ] 22.4 Create PR with updated baselines
- [ ] 22.5 Require manual approval for baseline updates

---

## Documentation

### 23. README and Documentation
- [ ] 23.1 Create `tests/e2e/README.md` with:
  - [ ] 23.1.1 Overview of E2E testing
  - [ ] 23.1.2 How to run tests locally
  - [ ] 23.1.3 How to add new scenarios
  - [ ] 23.1.4 Scenario YAML format reference
  - [ ] 23.1.5 Tolerance configuration guide
  - [ ] 23.1.6 Baseline management workflow
  - [ ] 23.1.7 Troubleshooting guide
- [ ] 23.2 Update root `CLAUDE.md` with E2E testing section
- [ ] 23.3 Update root `README.md` with testing information

### 24. Examples and Templates
- [ ] 24.1 Create scenario template with comments
- [ ] 24.2 Create example custom scenario walkthrough
- [ ] 24.3 Document common patterns and anti-patterns

---

## Final Validation

### 25. End-to-End Validation
- [ ] 25.1 Run complete test suite (`test-all`)
- [ ] 25.2 Verify all tests pass
- [ ] 25.3 Verify HTML reports generate correctly
- [ ] 25.4 Verify CI/CD workflow runs successfully
- [ ] 25.5 Manual testing with Microsoft Office
- [ ] 25.6 Performance testing (test suite < 10 minutes):
  - [ ] 25.6.1 Measure baseline test execution time for all 40+ scenarios
  - [ ] 25.6.2 Identify performance bottlenecks if execution >10 minutes
  - [ ] 25.6.3 Implement optimizations (parallel execution, caching) if needed
  - [ ] 25.6.4 Verify final wall-clock execution time <10 minutes in CI
- [ ] 25.7 Code review and cleanup
- [ ] 25.8 Update Spectr proposal with actual results
- [ ] 25.9 Mark Spectr change as complete
- [ ] 25.10 Archive Spectr change proposal

---

## Development Strategy

### Branch and PR Strategy
- **Phase 1 (MVP)**: Single feature branch `e2e-testing-mvp`, single PR after complete
- **Phase 2 (Comparison)**: Branch `e2e-testing-phase2`, PR after visual comparison working
- **Phase 3 (Coverage)**: Separate branches per document type:
  - `e2e-testing-excel` for spreadsheet support
  - `e2e-testing-powerpoint` for presentation support
  - `e2e-testing-word-advanced` for advanced Word scenarios
- **Phase 4 (CI/CD)**: Branch `e2e-testing-ci`, PR after GitHub Actions working

### Commit Strategy
- Commit at end of each numbered subsection (e.g., after 1.10, 2.12, 3.7, etc.)
- Use descriptive commit messages: "Phase 1.3: Implement scenario parsing" not "WIP"
- Squash-merge PRs to keep main branch history clean

### Testing Strategy Throughout
- Write tests alongside implementation (not after)
- Run `nix develop --command lint` before each commit
- Run `nix develop --command test-e2e` before pushing
- Manual testing in Microsoft Office for each major milestone
- **Test Coverage Target**: >80% for non-generated code in `tests/e2e/framework/`, `tests/e2e/comparison/`, and `tests/e2e/bridges/go/`
  - Generated code (scenario structs) may have lower coverage
  - Focus coverage on business logic, comparison algorithms, and error handling

### Example Test Structure (Table-Driven with Subtests)
```go
func TestScenarioLoader(t *testing.T) {
    t.Run("LoadValidScenario", func(t *testing.T) {
        // Arrange
        scenarioPath := "testdata/valid-scenario.yaml"

        // Act
        scenario, err := LoadScenario(scenarioPath)

        // Assert
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if scenario.Name != "expected-name" {
            t.Errorf("got name %q, want %q", scenario.Name, "expected-name")
        }
    })

    t.Run("LoadInvalidYAML", func(t *testing.T) {
        // ... test error handling ...
    })
}
```

## Notes

- Dependencies between tasks are implicit (earlier tasks must complete first within each phase)
- Phases can be implemented sequentially or with some parallelization
- Each checkbox should be marked complete only when fully tested and validated
- Create git commits at logical boundaries (end of each major task group)
- Keep PRs focused (one phase per PR recommended)
