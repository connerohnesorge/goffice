# Spreadsheet Worksheet Spec Delta

## ADDED Requirements

### Requirement: Worksheet Recalculation

The system SHALL provide methods on Worksheet to recalculate all formulas in the worksheet.

#### Scenario: Recalculate all formulas in worksheet
- GIVEN a worksheet with formulas in cells A1, B1, C1
- AND A1="=5", B1="=A1+1", C1="=B1+1"
- WHEN Recalculate() is called on the worksheet
- THEN all formulas are evaluated in dependency order
- AND A1 has value 5
- AND B1 has value 6
- AND C1 has value 7

#### Scenario: Recalculate with no formulas
- GIVEN a worksheet with no formula cells (all values)
- WHEN Recalculate() is called
- THEN no errors occur
- AND operation completes successfully
- AND no cells are modified

#### Scenario: Recalculate preserves formula text
- GIVEN a worksheet with formulas
- WHEN Recalculate() is called
- THEN cell values are updated to computed results
- AND formula text is preserved in all cells

#### Scenario: Recalculate only affects current worksheet
- GIVEN a workbook with Sheet1 and Sheet2
- AND both sheets have formulas
- WHEN Recalculate() is called on Sheet1
- THEN only Sheet1 formulas are recalculated
- AND Sheet2 formulas are not affected

#### Scenario: Recalculate with cross-sheet dependencies
- GIVEN Sheet1 with A1="=Sheet2!B5"
- AND Sheet2 with B5=100
- WHEN Recalculate() is called on Sheet1
- THEN A1 is updated to 100 (resolved from Sheet2)

#### Scenario: Recalculate with circular reference error
- GIVEN a worksheet with circular reference A1=B1, B1=A1
- WHEN Recalculate() is called
- THEN an error is returned
- AND error message indicates circular reference detected

#### Scenario: Recalculate updates all dependent cells
- GIVEN formulas A1=5, B1=A1+1, C1=A1+5, D1=B1+C1
- WHEN Recalculate() is called
- THEN A1=5, B1=6, C1=10, D1=16
- AND all dependencies are correctly evaluated

### Requirement: Worksheet Dependency Graph

The system SHALL provide access to the dependency graph for the worksheet.

#### Scenario: Get dependency graph for worksheet
- GIVEN a worksheet with formulas
- WHEN DependencyGraph() is called
- THEN a DependencyGraph is returned
- AND graph contains all formula cells in worksheet
- AND graph tracks dependencies between cells

#### Scenario: Dependency graph excludes value cells
- GIVEN a worksheet with 10 cells (5 formulas, 5 values)
- WHEN DependencyGraph() is called
- THEN graph contains only the 5 formula cells
- AND value cells are not included as nodes

#### Scenario: Dependency graph tracks cross-cell dependencies
- GIVEN formulas B1="=A1+1", C1="=B1+1"
- WHEN DependencyGraph() is called
- THEN graph shows B1 depends on A1
- AND graph shows C1 depends on B1

#### Scenario: Dependency graph for worksheet with no formulas
- GIVEN a worksheet with no formula cells
- WHEN DependencyGraph() is called
- THEN an empty graph is returned
- AND graph has no nodes or edges

### Requirement: Circular Reference Detection

The system SHALL detect and report circular references in worksheet formulas.

#### Scenario: Find direct circular reference
- GIVEN formulas A1="=B1", B1="=A1"
- WHEN FindCircularReferences() is called
- THEN a cycle [A1, B1, A1] is returned
- AND no error occurs (returns list of cycles)

#### Scenario: Find indirect circular reference
- GIVEN formulas A1="=B1", B1="=C1", C1="=A1"
- WHEN FindCircularReferences() is called
- THEN a cycle [A1, B1, C1, A1] is returned

#### Scenario: Find multiple circular references
- GIVEN two separate circular references
- AND cycle 1: A1=B1, B1=A1
- AND cycle 2: D1=E1, E1=D1
- WHEN FindCircularReferences() is called
- THEN two cycles are returned
- AND each cycle is reported independently

#### Scenario: No circular references
- GIVEN formulas A1=5, B1=A1+1, C1=B1+1 (no cycles)
- WHEN FindCircularReferences() is called
- THEN empty list is returned
- AND no error occurs

#### Scenario: Self-referencing cell
- GIVEN formula A1="=A1+1"
- WHEN FindCircularReferences() is called
- THEN a cycle [A1, A1] is returned

### Requirement: Worksheet Calculation Order

The system SHALL provide the calculation order for worksheet formulas.

#### Scenario: Get calculation order
- GIVEN formulas A1=5, B1=A1+1, C1=B1+1
- WHEN CalculationOrder() is called
- THEN cells are returned in order [A1, B1, C1]
- AND order respects dependencies

#### Scenario: Calculation order with independent cells
- GIVEN formulas A1=5, B1=10, C1=A1+B1
- WHEN CalculationOrder() is called
- THEN A1 and B1 come before C1
- AND A1 and B1 can be in any order (independent)

#### Scenario: Calculation order with circular reference
- GIVEN formulas with circular reference
- WHEN CalculationOrder() is called
- THEN an error is returned
- AND error indicates circular reference prevents ordering

### Requirement: Worksheet Formula Validation

The system SHALL validate formulas when added to worksheet.

#### Scenario: Validate formula syntax on add
- GIVEN a cell
- WHEN SetFormula("=A1+") is called
- THEN formula is stored (validation may be deferred)
- AND next Recalculate() detects syntax error

#### Scenario: Detect circular reference on add
- GIVEN cell A1 with formula "=B1"
- WHEN cell B1 formula is set to "=A1"
- THEN circular reference is not immediately detected
- AND next Recalculate() or FindCircularReferences() detects it

### Requirement: Worksheet Performance Optimization

The system SHALL optimize worksheet-level recalculation for performance.

#### Scenario: Incremental recalculation after cell change
- GIVEN worksheet with many formulas
- WHEN single cell value changes
- THEN only affected cells are recalculated
- AND unrelated cells are not recalculated

#### Scenario: Parallel evaluation of independent formulas
- GIVEN worksheet with independent formula cells
- WHEN Recalculate() is called
- THEN independent cells in same dependency level can be evaluated in parallel

#### Scenario: Cached results reuse
- GIVEN worksheet with formulas already evaluated
- WHEN Recalculate() is called without changes
- THEN cached results are reused
- AND formulas are not re-evaluated

### Requirement: Worksheet Formula Statistics

The system SHALL provide statistics about formulas in the worksheet.

#### Scenario: Count formula cells
- GIVEN a worksheet with 100 cells (20 formulas, 80 values)
- WHEN formula count is queried
- THEN 20 is returned

#### Scenario: Count cells by formula complexity
- GIVEN a worksheet with various formulas
- WHEN complexity analysis is performed
- THEN cells are categorized by dependency depth

### Requirement: Worksheet Error Handling

The system SHALL handle errors during worksheet-level operations.

#### Scenario: Handle invalid formula during recalculation
- GIVEN worksheet with one invalid formula "=A1 +"
- WHEN Recalculate() is called
- THEN invalid formula cell is set to #VALUE!
- AND other valid formulas are still calculated
- AND no error is returned (errors stored in cells)

#### Scenario: Handle missing dependency during recalculation
- GIVEN formula "=Sheet2!A1"
- AND Sheet2 does not exist
- WHEN Recalculate() is called
- THEN cell is set to #REF!
- AND error is handled gracefully

### Requirement: Worksheet Recalculation Events

The system SHALL provide hooks for recalculation events.

#### Scenario: Before recalculation hook
- GIVEN a worksheet with registered before-recalculation hook
- WHEN Recalculate() is called
- THEN hook is called before any formula evaluation
- AND hook can inspect current state

#### Scenario: After recalculation hook
- GIVEN a worksheet with registered after-recalculation hook
- WHEN Recalculate() is called
- THEN hook is called after all formulas are evaluated
- AND hook can inspect updated values
