# Spreadsheet Cell Spec Delta

## ADDED Requirements

### Requirement: Cell Formula Evaluation

The system SHALL provide methods on Cell to evaluate formulas and retrieve computed values.

#### Scenario: Evaluate formula and get result
- GIVEN cell C1 with formula "=A1+B1"
- AND A1 contains 5
- AND B1 contains 10
- WHEN EvaluateFormula() is called on C1
- THEN result is 15
- AND no error is returned

#### Scenario: Evaluate formula with error
- GIVEN cell C1 with formula "=A1/B1"
- AND A1 contains 10
- AND B1 contains 0
- WHEN EvaluateFormula() is called on C1
- THEN result is error value "#DIV/0!"

#### Scenario: Evaluate cell without formula
- GIVEN cell A1 with value 42 (no formula)
- WHEN EvaluateFormula() is called on A1
- THEN an error is returned
- AND error indicates cell has no formula

#### Scenario: Get computed value from formula cell
- GIVEN cell C1 with formula "=SUM(A1:A10)"
- WHEN GetComputedValue() is called on C1
- THEN formula is evaluated
- AND computed sum is returned

#### Scenario: Get computed value from value cell
- GIVEN cell A1 with value 42 (no formula)
- WHEN GetComputedValue() is called on A1
- THEN value 42 is returned
- AND no evaluation occurs

#### Scenario: Get cached value if available
- GIVEN cell C1 with formula "=A1+B1" and cached value 100
- WHEN GetComputedValue() is called on C1
- THEN cached value 100 is returned
- AND formula is not re-evaluated (cached value used)

#### Scenario: Set formula and evaluate immediately
- GIVEN cell C1
- WHEN SetFormulaAndEvaluate("=A1+B1") is called
- THEN formula text is set to "=A1+B1"
- AND formula is evaluated immediately
- AND cell value is set to computed result

#### Scenario: Set formula and evaluate with error
- GIVEN cell C1
- WHEN SetFormulaAndEvaluate("=A1/0") is called
- THEN formula text is set to "=A1/0"
- AND formula is evaluated
- AND cell value is set to #DIV/0!
- AND no error is returned (error stored in cell)

#### Scenario: Evaluate formula with complex expression
- GIVEN cell D1 with formula "=SUM(A1:A10) + AVERAGE(B1:B10)"
- WHEN EvaluateFormula() is called
- THEN both SUM and AVERAGE are evaluated
- AND results are added
- AND final result is returned

#### Scenario: Evaluate formula with cross-sheet reference
- GIVEN cell Sheet1!A1 with formula "=Sheet2!B5"
- AND Sheet2!B5 contains 100
- WHEN EvaluateFormula() is called on Sheet1!A1
- THEN value 100 is resolved from Sheet2!B5

### Requirement: Auto-Recalculation Hooks

The system SHALL trigger auto-recalculation when cell values change if enabled.

#### Scenario: Auto-recalculate on value change
- GIVEN workbook with auto-recalculate enabled
- AND cell C1 with formula "=A1+B1"
- AND C1 currently has value 10
- WHEN A1 value changes from 5 to 15
- THEN C1 is automatically recalculated
- AND C1 value updates to 25

#### Scenario: Auto-recalculate disabled
- GIVEN workbook with auto-recalculate disabled
- AND cell C1 with formula "=A1+B1"
- WHEN A1 value changes
- THEN C1 is NOT automatically recalculated
- AND C1 value remains unchanged (stale)

#### Scenario: Auto-recalculate cascades to dependents
- GIVEN formulas A1=5, B1=A1+1, C1=B1+1
- AND auto-recalculate enabled
- WHEN A1 value changes to 10
- THEN B1 is recalculated to 11
- AND C1 is recalculated to 12

#### Scenario: Auto-recalculate on SetNumber
- GIVEN cell A1 with auto-recalculate enabled
- WHEN SetNumber(42) is called on A1
- THEN A1 value is set to 42
- AND all cells depending on A1 are recalculated

#### Scenario: Auto-recalculate on SetString
- GIVEN cell A1 with auto-recalculate enabled
- WHEN SetString("hello") is called on A1
- THEN A1 value is set to "hello"
- AND all cells depending on A1 are recalculated

#### Scenario: Auto-recalculate on SetBoolean
- GIVEN cell A1 with auto-recalculate enabled
- WHEN SetBoolean(true) is called on A1
- THEN A1 value is set to TRUE
- AND all cells depending on A1 are recalculated

### Requirement: Formula Text Preservation

The system SHALL preserve formula text when setting values from evaluation.

#### Scenario: Evaluation updates value but preserves formula
- GIVEN cell C1 with formula "=A1+B1"
- WHEN EvaluateFormula() is called
- THEN cell value is updated to computed result
- AND formula text "=A1+B1" is preserved
- AND subsequent GetFormula() returns "=A1+B1"

#### Scenario: SetFormulaAndEvaluate preserves formula
- GIVEN cell C1
- WHEN SetFormulaAndEvaluate("=SUM(A1:A10)") is called
- THEN formula text "=SUM(A1:A10)" is stored
- AND cell value is set to computed sum
- AND formula text is not lost

### Requirement: Cell Value Type Handling

The system SHALL set appropriate cell value types based on evaluation results.

#### Scenario: Evaluation result is number
- GIVEN cell C1 with formula "=5+3"
- WHEN EvaluateFormula() is called
- THEN cell value type is set to Number
- AND cell value is 8

#### Scenario: Evaluation result is string
- GIVEN cell C1 with formula "=CONCATENATE(\"Hello\", \" World\")"
- WHEN EvaluateFormula() is called
- THEN cell value type is set to String
- AND cell value is "Hello World"

#### Scenario: Evaluation result is boolean
- GIVEN cell C1 with formula "=A1>5"
- WHEN EvaluateFormula() is called
- THEN cell value type is set to Boolean
- AND cell value is TRUE or FALSE

#### Scenario: Evaluation result is error
- GIVEN cell C1 with formula "=A1/0"
- WHEN EvaluateFormula() is called
- THEN cell value type is set to Error
- AND cell value is "#DIV/0!"

### Requirement: Evaluation Performance

The system SHALL optimize formula evaluation for performance.

#### Scenario: Cached value reuse
- GIVEN cell C1 with formula and valid cached value
- WHEN GetComputedValue() is called multiple times
- THEN formula is evaluated only once
- AND cached value is reused on subsequent calls

#### Scenario: Invalidate cache on dependency change
- GIVEN cell C1 with formula "=A1+B1" and cached value 10
- WHEN A1 value changes
- THEN C1 cached value is invalidated
- AND next GetComputedValue() re-evaluates formula

### Requirement: Error Handling in Cell API

The system SHALL handle evaluation errors gracefully in Cell API.

#### Scenario: Invalid formula syntax
- GIVEN cell C1 with formula "=A1 +"
- WHEN EvaluateFormula() is called
- THEN an error is returned
- AND error message indicates invalid syntax

#### Scenario: Unknown function
- GIVEN cell C1 with formula "=UNKNOWN(A1)"
- WHEN EvaluateFormula() is called
- THEN cell value is set to "#NAME?"
- AND no error is returned (error stored in cell)

#### Scenario: Circular reference in evaluation
- GIVEN cell A1 with formula "=B1"
- AND cell B1 with formula "=A1"
- WHEN EvaluateFormula() is called on A1
- THEN an error is returned
- AND error indicates circular reference
