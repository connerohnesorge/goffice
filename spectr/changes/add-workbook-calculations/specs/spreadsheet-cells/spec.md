## MODIFIED Requirements

### Requirement: Cell Value Reading
The system SHALL support reading cell values with type detection and formula evaluation.

#### Scenario: Read calculated value
- GIVEN a Cell with formula "=A1+B1" and A1=10, B1=20
- WHEN `cell.GetValue()` is called
- THEN calculated value 30 is returned
- AND formula is evaluated if needed

#### Scenario: Read cached formula result
- GIVEN a Cell with formula and cached result
- WHEN `cell.GetValue()` is called
- AND no dependent cells have changed
- THEN cached result is returned without recalculation

#### Scenario: Read formula string
- GIVEN a Cell with formula "=A1+B1"
- WHEN `cell.GetFormula()` is called
- THEN formula string "=A1+B1" is returned
- AND no evaluation occurs

#### Scenario: Force formula recalculation
- GIVEN a Cell with formula
- WHEN `cell.Recalculate()` is called
- THEN formula is re-evaluated
- AND cached result is updated

## ADDED Requirements

### Requirement: Cell Value Setting with Formula Support
The system SHALL support setting cell values and formulas with proper dependency tracking.

#### Scenario: Set cell formula
- GIVEN a Cell element
- WHEN `cell.SetFormula("=A1+B1")` is called
- THEN CellFormula element is created with the formula text
- AND cell is marked for recalculation
- AND dependent cells are tracked

#### Scenario: Set cell value with formula invalidation
- GIVEN a Cell with formula "=SUM(A1:A10)"
- WHEN `cell.SetValue(100)` is called
- THEN existing formula is removed
- AND dependent cells are invalidated for recalculation

#### Scenario: Set array formula
- GIVEN a range of cells
- WHEN `sheet.SetArrayFormula("=A1:A10*B1:B10", "C1:C10")` is called
- THEN array formula is applied to range
- AND calculation dependencies are established

### Requirement: Cell Dependencies Tracking
The system SHALL track dependencies between cells with formulas.

#### Scenario: Track formula dependencies
- GIVEN cell A1 formula "=B1+C1"
- WHEN formula is set
- THEN B1 and C1 are tracked as dependencies of A1
- AND A1 is tracked as a dependent of B1 and C1

#### Scenario: Update dependencies on formula change
- GIVEN cell A1 formula "=B1+C1"
- WHEN formula changes to "=B1+D1"
- THEN dependency graph is updated
- AND C1 dependency is removed, D1 dependency is added

#### Scenario: Invalidate dependent cells
- GIVEN B1 has dependent A1 with formula "=B1+1"
- WHEN B1 value changes
- THEN A1 is marked as needing recalculation
- AND notification cascade may occur

### Requirement: Cell Reference System with Formula Integration
The system SHALL provide cell reference utilities with formula evaluation support.

#### Scenario: Parse and resolve reference in formula
- GIVEN formula "=Sheet2!A1+Sheet3!B2"
- WHEN references are parsed during evaluation
- THEN cross-sheet references are correctly resolved
- AND values from referenced sheets are obtained

#### Scenario: Handle deleted references in formulas
- GIVEN formula "=A5+B5" and row 5 is deleted
- WHEN formula is evaluated
- THEN deleted references become #REF!
- AND appropriate error handling occurs

#### Scenario: Update formula references on structural changes
- GIVEN formula "=SUM(A1:A10)" and row 5 is inserted
- WHEN formula is updated
- THEN range becomes "=SUM(A1:A11)"
- AND formula semantics are preserved

### Requirement: Cell Error Values
The system SHALL support Excel error values in cells.

#### Scenario: Set calculation error
- GIVEN formula evaluation fails (e.g., division by zero)
- WHEN evaluation occurs
- THEN cell value is set to appropriate error (#DIV/0!)
- AND error type is preserved

#### Scenario: Set circular reference error
- GIVEN circular reference detection
- WHEN evaluation occurs
- THEN cell value is set to #CIRCULAR!
- AND error indicates circular reference

#### Scenario: Propagate formula errors
- GIVEN A1 has error #DIV/0! and B1 formula "=A1*2"
- WHEN B1 is evaluated
- THEN B1 also returns #DIV/0!
- AND error propagation follows Excel rules