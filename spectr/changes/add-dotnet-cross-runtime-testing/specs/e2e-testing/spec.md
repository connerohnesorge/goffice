# E2E Testing Specification

**Related Documents**:
- Full proposal context: See `../../proposal.md`
- Implementation design: See `../../design.md`
- Task breakdown: See `../../tasks.md`
- Nix environment spec: See `../nix-environment/spec.md`

## ADDED Requirements

### Requirement: Cross-Runtime Test Scenarios
The system SHALL support declarative test scenarios that execute identically in both Go (goffice) and C# (Open-XML-SDK) to validate API parity.

**Default Tolerance Values**:
- `VisualPixelTolerance`: 2.0 (Euclidean color distance per pixel)
- `VisualDiffThreshold`: 0.005 (0.5% of total pixels may differ)
- `XMLAttributeOrderSensitive`: false

**Supported Operations (Phase 1 - Word Documents)**:
- `create_document` - Create new document
- `add_paragraph` - Add paragraph with optional nested runs
- `add_run` - Add text run with formatting
- `add_table` - Add table with rows and columns

**Future Operations (Phase 3+)**:
- Spreadsheet operations: `create_workbook`, `add_sheet`, `set_cell_value`, `set_cell_formula`
- Presentation operations: `create_presentation`, `add_slide`, `add_shape`, `add_text_box`
- Advanced Word operations: `add_header`, `add_footer`, `add_image`, `add_hyperlink`

#### Scenario: Load YAML test scenario
- GIVEN a YAML file defining document operations
- WHEN the scenario is loaded via LoadScenario()
- THEN a TestScenario struct is returned with all operations parsed

#### Scenario: Execute scenario in Go bridge
- GIVEN a valid TestScenario
- WHEN Execute() is called on GoBridge
- THEN a document is generated using goffice API
- AND the document is saved to the specified output path

#### Scenario: Execute scenario in C# bridge
- GIVEN a valid TestScenario
- WHEN the C# bridge CLI is invoked with --scenario and --output
- THEN a document is generated using Open-XML-SDK API
- AND the document is saved to the specified output path

#### Scenario: Scenario supports tolerance configuration
- GIVEN a scenario with tolerance overrides in YAML
- WHEN the scenario is loaded
- THEN ToleranceConfig is populated with custom values
- AND defaults are used for unspecified tolerances (VisualPixelTolerance=2.0, VisualDiffThreshold=0.005)

#### Scenario: XML attribute order tolerance applied
- GIVEN tolerance config with XMLAttributeOrderSensitive = true
- WHEN XML comparison encounters different attribute order
- THEN mismatch is reported as error

#### Scenario: Visual pixel tolerance applied
- GIVEN tolerance with VisualPixelTolerance = 5.0
- WHEN visual comparison finds pixels differing by 3.0
- THEN difference is within tolerance and not reported

#### Scenario: Visual diff threshold applied
- GIVEN tolerance with VisualDiffThreshold = 0.01 (1%)
- WHEN visual comparison finds 0.5% pixel differences
- THEN comparison passes as within threshold

---

### Requirement: Three-Level Document Comparison
The system SHALL compare generated documents at three levels (XML structure, binary content, visual rendering) with configurable tolerances.

#### Scenario: Compare XML structure
- GIVEN two DOCX files (Go-generated and .NET-generated)
- WHEN CompareXMLStructure() is called
- THEN XML trees are parsed from both documents
- AND element names, attributes, and content are compared recursively
- AND mismatches are recorded in XMLDiffResult

#### Scenario: Detect element mismatches
- GIVEN Go document with extra paragraph element
- WHEN XML comparison runs
- THEN ElementMismatch is recorded with XPath and expected/actual values

#### Scenario: Detect attribute mismatches
- GIVEN Go document with different font size attribute value
- WHEN XML comparison runs
- THEN AttributeMismatch is recorded with attribute name and expected/actual values

#### Scenario: Respect attribute order tolerance
- GIVEN tolerance config with XMLAttributeOrderSensitive = false
- WHEN XML comparison runs
- THEN attributes in different order are not reported as mismatches

#### Scenario: Handle multi-page visual diff
- GIVEN two documents with multiple pages that differ
- WHEN CompareVisualRendering() is called
- THEN each page is converted to PNG separately
- AND per-page diff images are generated
- AND overall match status reflects all pages

#### Scenario: Compare binary content
- GIVEN two DOCX files
- WHEN CompareBinaryContent() is called
- THEN ZIP parts are extracted from both documents
- AND missing, extra, and differing parts are identified
- AND BinaryDiffResult is returned with part differences

#### Scenario: Compare visual rendering
- GIVEN two DOCX files
- WHEN CompareVisualRendering() is called
- THEN both documents are converted to PDF via LibreOffice
- AND PDFs are rendered to PNG via Ghostscript
- AND pixel-level diff is calculated
- AND VisualDiffResult is returned with diff percentage and images

#### Scenario: Generate diff image on mismatch
- GIVEN two documents that differ visually
- WHEN visual comparison completes
- THEN an annotated diff PNG is generated
- AND exceeding-tolerance pixels are highlighted in red

---

### Requirement: Test Execution Framework
The system SHALL provide a test executor that orchestrates scenario execution through both bridges and aggregates comparison results.

#### Scenario: Execute single scenario
- GIVEN a TestScenario and Executor
- WHEN ExecuteScenario() is called
- THEN Go bridge generates document
- AND C# bridge generates document
- AND three-level comparison runs
- AND TestResult is returned with pass/fail status

#### Scenario: Handle Go bridge error
- GIVEN a scenario that fails in Go bridge
- WHEN ExecuteScenario() runs
- THEN TestResult.Status is StatusFailed
- AND TestResult.GoError contains error message
- AND comparison is skipped

#### Scenario: Handle C# bridge error
- GIVEN a scenario that fails in C# bridge
- WHEN ExecuteScenario() runs
- THEN TestResult.Status is StatusFailed
- AND TestResult.DotNetError contains error message
- AND comparison is skipped

#### Scenario: Handle file write permission error
- GIVEN output directory is read-only
- WHEN ExecuteScenario() attempts to save document
- THEN TestResult.Status is StatusFailed
- AND error message indicates permission denied

#### Scenario: Handle bridge timeout
- GIVEN C# bridge takes longer than timeout limit
- WHEN ExecuteScenario() runs with timeout=30s
- THEN execution is cancelled after timeout
- AND TestResult.DotNetError contains "timeout exceeded"

#### Scenario: Handle invalid scenario operation
- GIVEN scenario with unknown action "add_invalid"
- WHEN bridge executes operations
- THEN error is reported with "unknown action"
- AND TestResult.Status is StatusFailed

#### Scenario: Handle comparison error
- GIVEN both bridges succeed in generating documents
- WHEN comparison itself fails with internal error
- THEN TestResult.ComparisonError contains error message
- AND TestResult.Status is StatusFailed

#### Scenario: Determine pass/fail from comparison
- GIVEN both bridges succeed
- WHEN all three comparison levels match
- THEN TestResult.Status is StatusPassed

#### Scenario: Execute multiple scenarios
- GIVEN 10 test scenarios
- WHEN all scenarios are executed
- THEN 10 TestResults are returned
- AND each result has independent status

---

### Requirement: CLI Test Runner
The system SHALL provide a command-line interface for running E2E tests with flexible options.

#### Scenario: Run tests from scenarios directory
- GIVEN scenarios in tests/e2e/scenarios/wordprocessing/
- WHEN e2e-runner is invoked with --scenarios flag
- THEN all YAML files are loaded as scenarios
- AND all scenarios are executed

#### Scenario: Generate HTML report
- GIVEN test execution completes
- WHEN --html-report flag is true
- THEN an HTML report is generated at reports/html/index.html
- AND report contains pass/fail summary
- AND report contains diff images for failures

#### Scenario: Exit with error code on failure
- GIVEN at least one test fails
- WHEN test suite completes
- THEN exit code is 1

#### Scenario: Exit with success code on all pass
- GIVEN all tests pass
- WHEN test suite completes
- THEN exit code is 0

---

### Requirement: Baseline Management
The system SHALL support generating and updating reference documents from the .NET SDK for regression testing.

#### Scenario: Generate baselines from .NET SDK
- GIVEN test scenarios
- WHEN --generate-baselines flag is set
- THEN C# bridge generates documents for all scenarios
- AND documents are saved to tests/e2e/baselines/

#### Scenario: Compare against baselines
- GIVEN baseline documents exist
- WHEN tests run without --generate-baselines
- THEN Go-generated documents are compared against baselines
- AND .NET bridge is not invoked (faster execution)

#### Scenario: Baseline mode execution
- GIVEN baseline documents exist in baselines directory
- WHEN ExecuteScenario() runs with baseline mode enabled
- THEN Go bridge generates document
- AND document is compared against baseline only
- AND .NET bridge is not invoked
- AND TestResult indicates baseline comparison mode

#### Scenario: Missing baseline fallback
- GIVEN no baseline exists for scenario
- WHEN baseline mode is requested
- THEN error is reported indicating baseline not found
- AND test is marked as skipped or failed

#### Scenario: Update baselines workflow
- GIVEN .NET SDK version has changed
- WHEN update-baselines script runs
- THEN all baselines are regenerated
- AND git diff shows baseline changes
- AND user reviews and approves changes

---

### Requirement: Go Bridge Implementation
The system SHALL translate test scenarios into goffice API calls to generate documents.

#### Scenario: Handle create_document operation
- GIVEN scenario operation with action "create_document"
- WHEN Go bridge executes operation
- THEN wordprocessing.New() is called with correct DocType

#### Scenario: Handle add_paragraph operation
- GIVEN scenario operation with action "add_paragraph"
- WHEN Go bridge executes operation
- THEN DocumentBuilder.AddParagraph() is called

#### Scenario: Apply run properties
- GIVEN scenario operation with run properties (bold, italic, color)
- WHEN Go bridge Execute() is called
- THEN RunBuilder methods are called (Bold(), Italic(), Color())

#### Scenario: Handle add_table operation
- GIVEN scenario operation with rows and cols
- WHEN Go bridge Execute() is called
- THEN DocumentBuilder.AddTable() is called with correct dimensions

---

### Requirement: C# Bridge Implementation
The system SHALL translate test scenarios into Open-XML-SDK API calls to generate documents.

#### Scenario: Parse scenario from YAML
- GIVEN scenario YAML file path
- WHEN C# bridge is invoked with --scenario
- THEN YamlDotNet deserializes YAML to TestScenario object

#### Scenario: Handle create_document operation
- GIVEN scenario operation "create_document"
- WHEN C# bridge executes operation
- THEN WordprocessingDocument.Create() is called

#### Scenario: Handle add_paragraph operation
- GIVEN scenario operation "add_paragraph"
- WHEN C# bridge executes operation
- THEN new Paragraph() is created and appended to Body

#### Scenario: Apply run properties
- GIVEN run properties in scenario
- WHEN C# bridge applies properties
- THEN RunProperties element is created with Bold, Italic, Color children

#### Scenario: CLI returns success exit code
- GIVEN document generation succeeds
- WHEN C# bridge completes
- THEN exit code is 0

#### Scenario: CLI returns error exit code on failure
- GIVEN document generation fails
- WHEN C# bridge encounters error
- THEN error message is printed to stderr
- AND exit code is 1

---

### Requirement: Reporting
The system SHALL generate comprehensive reports of test results in multiple formats.

#### Scenario: Generate console summary
- GIVEN test execution completes
- WHEN printSummary() is called
- THEN console output shows total, passed, failed, skipped counts
- AND output format matches:
  ```
  ================================
  Test Results Summary
  ================================
  Total:   <N>
  Passed:  <N>
  Failed:  <N>
  Skipped: <N>
  ================================
  ```

#### Scenario: Generate JSON report
- GIVEN test results
- WHEN JSON report is requested
- THEN JSON file is written with all TestResults serialized

#### Scenario: Generate HTML report with diff images
- GIVEN test results with failures
- WHEN HTML report is generated
- THEN HTML includes embedded diff images
- AND side-by-side comparison views
- AND diff statistics are displayed (DiffPercentage, MaxColorDelta, AvgColorDelta)

#### Scenario: Report includes execution time
- GIVEN TestResult with StartTime and EndTime
- WHEN report is generated
- THEN execution duration is calculated as (EndTime - StartTime)
- AND duration is displayed in human-readable format (e.g., "2.5s", "1m 30s")
