# Spreadsheet Workbook Spec Delta

## ADDED Requirements

### Requirement: Workbook Recalculation

The system SHALL provide methods on Workbook to recalculate all formulas across all worksheets.

#### Scenario: Recalculate entire workbook
- GIVEN a workbook with 3 worksheets
- AND each worksheet has formulas
- WHEN Recalculate() is called on the workbook
- THEN all formulas in all worksheets are evaluated
- AND formulas are evaluated in global dependency order

#### Scenario: Recalculate workbook with cross-sheet dependencies
- GIVEN Sheet1 with A1="=Sheet2!B1"
- AND Sheet2 with B1="=Sheet3!C1"
- AND Sheet3 with C1=100
- WHEN Recalculate() is called
- THEN C1 is evaluated first (value: 100)
- AND B1 is evaluated second (value: 100)
- AND A1 is evaluated third (value: 100)

#### Scenario: Recalculate preserves worksheet independence
- GIVEN a workbook with Sheet1 and Sheet2
- AND no cross-sheet dependencies
- WHEN Recalculate() is called
- THEN both sheets are recalculated
- AND evaluation order can be Sheet1 then Sheet2, or Sheet2 then Sheet1

#### Scenario: Recalculate with circular reference across sheets
- GIVEN Sheet1!A1="=Sheet2!B1"
- AND Sheet2!B1="=Sheet1!A1"
- WHEN Recalculate() is called
- THEN an error is returned
- AND error indicates cross-sheet circular reference

#### Scenario: Recalculate workbook performance
- GIVEN a workbook with 1000 formulas across 10 worksheets
- WHEN Recalculate() is called
- THEN all formulas are evaluated
- AND operation completes in reasonable time (<100ms for typical formulas)

### Requirement: Calculation Chain Management

The system SHALL build and manage calculation chains for workbooks.

#### Scenario: Rebuild calculation chain
- GIVEN a workbook with formulas
- WHEN RebuildCalculationChain() is called
- THEN CalcChainPart is populated with all formula cells
- AND cells are ordered by dependency order
- AND calculation chain is persisted in workbook

#### Scenario: Calculation chain with cross-sheet dependencies
- GIVEN formulas across multiple sheets
- WHEN RebuildCalculationChain() is called
- THEN calculation chain includes cells from all sheets
- AND cells are ordered respecting cross-sheet dependencies

#### Scenario: Calculation chain excludes value cells
- GIVEN a workbook with 1000 cells (500 formulas, 500 values)
- WHEN RebuildCalculationChain() is called
- THEN calculation chain contains only 500 formula cells

#### Scenario: Calculation chain persistence
- GIVEN a workbook with calculation chain rebuilt
- WHEN workbook is saved to file
- AND file is reopened
- THEN calculation chain is preserved
- AND cells are still in correct order

#### Scenario: Rebuild calculation chain with circular reference
- GIVEN a workbook with circular reference
- WHEN RebuildCalculationChain() is called
- THEN an error is returned
- AND calculation chain is not updated (remains unchanged or empty)

### Requirement: Calculation Order

The system SHALL provide global calculation order for all worksheets.

#### Scenario: Get calculation order for workbook
- GIVEN formulas across 3 worksheets
- WHEN CalculationOrder() is called
- THEN all formula cells are returned in dependency order
- AND cross-sheet dependencies are respected

#### Scenario: Calculation order with no dependencies
- GIVEN formulas with no interdependencies
- WHEN CalculationOrder() is called
- THEN all cells are returned in any valid order

#### Scenario: Calculation order includes sheet names
- GIVEN formulas in Sheet1 and Sheet2
- WHEN CalculationOrder() is called
- THEN returned cell references include sheet names
- AND cells are identifiable by (SheetName, Col, Row)

### Requirement: Auto-Recalculation Settings

The system SHALL provide workbook-level auto-recalculation settings.

#### Scenario: Enable auto-recalculation
- GIVEN a workbook
- WHEN SetAutoRecalculate(true) is called
- THEN auto-recalculation is enabled
- AND cell value changes trigger recalculation

#### Scenario: Disable auto-recalculation
- GIVEN a workbook with auto-recalculate enabled
- WHEN SetAutoRecalculate(false) is called
- THEN auto-recalculation is disabled
- AND cell value changes do NOT trigger recalculation

#### Scenario: Auto-recalculation default setting
- GIVEN a newly created workbook
- THEN auto-recalculation is enabled by default

#### Scenario: Get auto-recalculation status
- GIVEN a workbook
- WHEN GetAutoRecalculate() is called
- THEN current auto-recalculation setting is returned (true or false)

#### Scenario: Auto-recalculation persists across save/load
- GIVEN a workbook with auto-recalculate disabled
- WHEN workbook is saved and reopened
- THEN auto-recalculate setting remains disabled

### Requirement: Workbook Calculation Mode

The system SHALL support different calculation modes (automatic, manual, automatic except tables).

#### Scenario: Set calculation mode to automatic
- GIVEN a workbook
- WHEN SetCalculationMode(Automatic) is called
- THEN calculation mode is set to automatic
- AND all changes trigger recalculation

#### Scenario: Set calculation mode to manual
- GIVEN a workbook
- WHEN SetCalculationMode(Manual) is called
- THEN calculation mode is set to manual
- AND changes do NOT trigger recalculation
- AND user must call Recalculate() explicitly

#### Scenario: Get calculation mode
- GIVEN a workbook with calculation mode Manual
- WHEN GetCalculationMode() is called
- THEN Manual is returned

### Requirement: Global Dependency Graph

The system SHALL provide a global dependency graph across all worksheets.

#### Scenario: Get global dependency graph
- GIVEN a workbook with formulas in multiple sheets
- WHEN DependencyGraph() is called on workbook
- THEN a global dependency graph is returned
- AND graph includes cells from all worksheets
- AND cross-sheet dependencies are tracked

#### Scenario: Global dependency graph with isolated sheets
- GIVEN Sheet1 and Sheet2 with no cross-sheet dependencies
- WHEN DependencyGraph() is called
- THEN graph has two independent subgraphs
- AND each subgraph represents one sheet

#### Scenario: Global dependency graph update on formula change
- GIVEN a workbook with dependency graph built
- WHEN a formula is added/modified/removed
- THEN dependency graph is updated
- AND affected dependencies are re-evaluated

### Requirement: Workbook-Level Circular Reference Detection

The system SHALL detect circular references across all worksheets.

#### Scenario: Detect circular reference across sheets
- GIVEN Sheet1!A1="=Sheet2!B1"
- AND Sheet2!B1="=Sheet1!A1"
- WHEN FindCircularReferences() is called on workbook
- THEN circular reference [Sheet1!A1, Sheet2!B1, Sheet1!A1] is returned

#### Scenario: Detect multiple circular references in workbook
- GIVEN multiple separate circular references across sheets
- WHEN FindCircularReferences() is called
- THEN all circular references are returned
- AND each cycle is reported separately

#### Scenario: No circular references in workbook
- GIVEN a workbook with no circular references
- WHEN FindCircularReferences() is called
- THEN empty list is returned

### Requirement: Workbook Formula Statistics

The system SHALL provide statistics about formulas across the workbook.

#### Scenario: Count total formulas in workbook
- GIVEN a workbook with 500 formula cells across 5 worksheets
- WHEN total formula count is queried
- THEN 500 is returned

#### Scenario: Count formulas per worksheet
- GIVEN a workbook with formulas in each worksheet
- WHEN per-sheet formula counts are queried
- THEN counts for each worksheet are returned

#### Scenario: Count cells by formula type
- GIVEN a workbook with various formula types
- WHEN formula type analysis is performed
- THEN cells are categorized (e.g., 200 SUM, 100 IF, 50 VLOOKUP, etc.)

### Requirement: Workbook Recalculation Performance

The system SHALL optimize workbook-level recalculation for performance.

#### Scenario: Incremental recalculation after single cell change
- GIVEN a large workbook with many formulas
- WHEN a single cell value changes
- THEN only cells dependent on changed cell are recalculated
- AND unaffected cells are not recalculated

#### Scenario: Parallel evaluation across worksheets
- GIVEN a workbook with independent worksheets
- WHEN Recalculate() is called
- THEN worksheets can be evaluated in parallel

#### Scenario: Cached dependency graph reuse
- GIVEN a workbook with dependency graph already built
- WHEN Recalculate() is called multiple times
- THEN dependency graph is reused
- AND topological sort is not recalculated unless formulas change

#### Scenario: Lazy evaluation of unchanged formulas
- GIVEN a workbook with formulas already evaluated
- WHEN Recalculate() is called without changes
- THEN no formulas are re-evaluated
- AND cached values are used

### Requirement: Workbook Error Handling

The system SHALL handle errors during workbook-level operations.

#### Scenario: Handle circular reference during recalculation
- GIVEN a workbook with circular reference
- WHEN Recalculate() is called
- THEN an error is returned
- AND error message identifies the circular reference

#### Scenario: Handle invalid formulas during recalculation
- GIVEN a workbook with some invalid formulas
- WHEN Recalculate() is called
- THEN invalid formula cells are set to #VALUE!
- AND valid formulas are still calculated
- AND no error is returned (errors stored in cells)

#### Scenario: Handle missing cross-sheet references
- GIVEN formula "=Sheet5!A1"
- AND Sheet5 does not exist
- WHEN Recalculate() is called
- THEN cell is set to #REF!
- AND error is handled gracefully

### Requirement: Calculation Chain Part API

The system SHALL provide API for manipulating CalcChainPart.

#### Scenario: Add cell to calculation chain
- GIVEN a CalcChainPart
- WHEN AddCalculationCell(sheetName, col, row, index) is called
- THEN cell is added to calculation chain at specified index

#### Scenario: Remove cell from calculation chain
- GIVEN a CalcChainPart with cells
- WHEN RemoveCalculationCell(sheetName, col, row) is called
- THEN cell is removed from calculation chain

#### Scenario: Clear calculation chain
- GIVEN a CalcChainPart with 100 cells
- WHEN Clear() is called
- THEN calculation chain is emptied
- AND no cells remain

#### Scenario: Get calculation order from chain
- GIVEN a CalcChainPart with ordered cells
- WHEN GetOrder() is called
- THEN cells are returned in stored order

### Requirement: Workbook Recalculation Events

The system SHALL provide hooks for workbook-level recalculation events.

#### Scenario: Before workbook recalculation hook
- GIVEN a workbook with registered before-recalculation hook
- WHEN Recalculate() is called
- THEN hook is called before any formula evaluation
- AND hook receives workbook instance

#### Scenario: After workbook recalculation hook
- GIVEN a workbook with registered after-recalculation hook
- WHEN Recalculate() is called
- THEN hook is called after all formulas are evaluated
- AND hook can inspect all updated values

#### Scenario: On circular reference detected hook
- GIVEN a workbook with registered circular-reference hook
- WHEN circular reference is detected
- THEN hook is called with cycle information
- AND hook can handle error or abort recalculation

### Requirement: Workbook Calculation Settings Persistence

The system SHALL persist calculation settings in the workbook file.

#### Scenario: Save calculation mode
- GIVEN a workbook with calculation mode Manual
- WHEN workbook is saved
- THEN calculation mode is stored in workbook.xml
- AND mode is restored when file is opened

#### Scenario: Save auto-recalculate setting
- GIVEN a workbook with auto-recalculate disabled
- WHEN workbook is saved
- THEN setting is stored in CalcPr element
- AND setting is restored when file is opened

#### Scenario: Default calculation settings for new workbook
- GIVEN a new workbook is created
- THEN calculation mode defaults to Automatic
- AND auto-recalculate defaults to true
- AND settings are persisted on first save

### Requirement: Workbook Formula Validation

The system SHALL validate formulas at the workbook level.

#### Scenario: Validate all formulas in workbook
- GIVEN a workbook with 100 formulas
- WHEN ValidateAllFormulas() is called
- THEN all formulas are parsed and validated
- AND list of invalid formulas is returned

#### Scenario: Validate formula syntax only
- GIVEN a workbook
- WHEN ValidateAllFormulas(syntaxOnly=true) is called
- THEN only syntax is validated
- AND circular references are not checked

#### Scenario: Validate including circular references
- GIVEN a workbook
- WHEN ValidateAllFormulas(syntaxOnly=false) is called
- THEN syntax and circular references are validated
- AND both types of errors are reported
