## ADDED Requirements

### Requirement: Formula Parser
The system SHALL provide comprehensive formula parsing capabilities for Excel formulas.

#### Scenario: Parse arithmetic formula
- GIVEN a formula "=A1+B1*C1/D1"
- WHEN parsed through FormulaParser
- THEN a parse tree with correct operator precedence is generated
- AND operations are ordered: multiplication/division before addition/subtraction

#### Scenario: Parse function formula
- GIVEN a formula "=SUM(A1:A10, B1:B10)"
- WHEN parsed through FormulaParser
- THEN a function node with name SUM is created
- AND two range arguments are correctly parsed

#### Scenario: Parse nested function formula
- GIVEN a formula "=IF(A1>0, SUM(B1:B10), AVERAGE(C1:C10))"
- WHEN parsed through FormulaParser
- THEN nested function structure is correctly represented
- AND IF function has three arguments: condition, true-value, false-value

#### Scenario: Parse string literal formula
- GIVEN a formula "=\"Total: \"&A1"
- WHEN parsed through FormulaParser
- THEN string literal "Total: " is correctly identified
- AND concatenation operator & is parsed correctly

#### Scenario: Parse formula with error handling
- GIVEN a malformed formula "=SUM(A1:A10"
- WHEN parsed through FormulaParser
- THEN appropriate parse error is returned
- AND error location in formula is identified

### Requirement: Cell Evaluation Engine
The system SHALL evaluate parsed formulas to calculate cell values.

#### Scenario: Evaluate arithmetic expression
- GIVEN parsed formula "=A1+B1" with A1=10, B1=20
- WHEN evaluated
- THEN result is 30
- AND result type is numeric

#### Scenario: Evaluate function call
- GIVEN parsed formula "=SUM(A1:A3)" with A1=10, A2=20, A3=30
- WHEN evaluated
- THEN result is 60
- AND range references are correctly resolved

#### Scenario: Evaluate volatile function
- GIVEN parsed formula "=NOW()" or "=RAND()"
- WHEN evaluated multiple times
- THEN results may differ between evaluations
- AND function is marked as volatile for recalculation

#### Scenario: Evaluate with circular reference
- GIVEN A1 formula "=B1+1" and B1 formula "=A1+1"
- WHEN evaluation is attempted
- THEN circular reference is detected
- AND appropriate error is returned

#### Scenario: Evaluate with dependency tracking
- GIVEN multiple cells with interdependent formulas
- WHEN any dependent cell value changes
- THEN affected cells are identified for recalculation
- AND dependency graph is maintained

### Requirement: Calculation Mode Support
The system SHALL support different calculation modes as defined by Excel.

#### Scenario: Automatic calculation mode
- GIVEN workbook with CalcMode.Auto
- WHEN any cell value changes
- THEN dependent formulas are automatically recalculated
- AND calculation chain is updated

#### Scenario: Manual calculation mode
- GIVEN workbook with CalcMode.Manual
- WHEN cell values change
- THEN formulas are NOT automatically recalculated
- AND calculation only occurs when explicitly triggered

#### Scenario: Automatic except tables mode
- GIVEN workbook with CalcMode.AutoNoTable
- WHEN data table cells change
- THEN data table formulas are not recalculated
- AND other formulas are recalculated normally

### Requirement: Circular Reference Detection
The system SHALL detect and report circular references in formulas.

#### Scenario: Direct circular reference detection
- GIVEN cell A1 formula "=A1+1" (self-reference)
- WHEN evaluation is attempted
- THEN circular reference is immediately detected
- AND error indicates the problematic cell

#### Scenario: Indirect circular reference detection
- GIVEN A1="=B1+1", B1="=C1+1", C1="=A1+1" (cycle)
- WHEN evaluation is attempted
- THEN circular reference is detected through dependency traversal
- AND cycle path is reported

#### Scenario: Circular reference in complex dependency
- GIVEN large dependency graph with one cycle
- WHEN evaluation is attempted
- THEN only cells involved in cycle report circular reference error
- AND other cells evaluate normally

### Requirement: Named Range Resolution
The system SHALL resolve named ranges in formulas to actual cell references.

#### Scenario: Resolve global named range
- GIVEN workbook defines "SalesData" = "Sheet1!$A$1:$D$100"
- WHEN formula "=SUM(SalesData)" is evaluated
- THEN named range is resolved to actual range reference
- AND evaluation proceeds with resolved range

#### Scenario: Resolve local named range
- GIVEN sheet defines "TaxRate" = "0.0875" with LocalSheetId=0
- WHEN formula on Sheet1 uses "=A1*TaxRate"
- THEN local named range is resolved
- AND sheet scope is respected

#### Scenario: Resolve named range with formula
- GIVEN named range "QuarterTotal" = "=SUM(Q1:Q4)"
- WHEN formula "=QuarterTotal*0.1" is evaluated
- THEN named range formula is evaluated first
- AND result is used in outer formula

### Requirement: External Link Resolution
The system SHALL handle references to external workbooks.

#### Scenario: Parse external reference
- GIVEN formula "='[Budget.xlsx]Sheet1'!$A$1"
- WHEN parsed
- THEN external workbook reference is correctly identified
- AND sheet name and cell reference are extracted

#### Scenario: Resolve external link
- GIVEN external reference to external workbook
- WHEN evaluation requires external value
- THEN system attempts to resolve external link
- AND appropriate handling if link is unavailable

### Requirement: Volatile Function Tracking
The system SHALL track volatile functions for proper recalculation.

#### Scenario: Identify volatile functions
- GIVEN formula contains NOW(), TODAY(), RAND(), OFFSET(), etc.
- WHEN parsed
- THEN formula is marked as volatile
- AND special recalculation rules apply

#### Scenario: Force recalculation of volatile
- GIVEN volatile formulas in workbook
- WHEN workbook opens or time passes
- THEN all volatile functions are recalculated
- AND dependent cells are updated

### Requirement: Array Formula Support
The system SHALL support array formula evaluation.

#### Scenario: Evaluate CSE array formula
- GIVEN array formula "=A1:A3*B1:B3" entered with Ctrl+Shift+Enter
- WHEN evaluated
- THEN results are calculated for each cell in array range
- AND results are properly distributed to array cells

#### Scenario: Evaluate dynamic array formula
- GIVEN modern Excel formula "=SORT(A1:B10)"
- WHEN evaluated
- THEN dynamic spilling behavior is simulated
- AND affected range is calculated

### Requirement: Calculation Chain Optimization
The system SHALL optimize calculation order for performance.

#### Scenario: Build calculation chain
- GIVEN workbook with multiple dependent formulas
- WHEN calculation chain is built
- THEN cells are ordered by dependency
- AND minimal recalculation paths are identified

#### Scenario: Incremental recalculation
- GIVEN single cell value change
- WHEN incremental recalculation is performed
- THEN only directly dependent cells are recalculated
- AND unrelated cells are skipped