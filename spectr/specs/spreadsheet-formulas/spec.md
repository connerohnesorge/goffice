# Spreadsheet Formulas Specification

## Requirements

### Requirement: Formula Storage

The system SHALL store formulas without evaluation.

#### Scenario: Set simple formula
- GIVEN a Cell element
- WHEN `cell.SetFormula("=A1+B1")` is called
- THEN a CellFormula element is created with the formula text
- AND FormulaType defaults to Normal

#### Scenario: Set formula with result
- GIVEN a Cell with a formula
- WHEN the formula result is known
- THEN CellValue can be set to cache the result
- AND Excel will recalculate on open if needed

#### Scenario: Read formula
- GIVEN a Cell with a formula
- WHEN `cell.Formula()` is called
- THEN the formula text is returned

### Requirement: Shared Formulas

The system SHALL support shared formulas for optimization.

#### Scenario: Create shared formula
- GIVEN a range of cells with similar formulas
- WHEN `sheet.SetSharedFormula("=A1+1", "B1:B100")` is called
- THEN the master cell (B1) has CellFormula with SharedIndex and Reference
- AND other cells reference the SharedIndex

#### Scenario: Read shared formula
- GIVEN a Cell with shared formula reference
- WHEN `cell.Formula()` is called
- THEN the resolved formula for that cell is returned

### Requirement: Array Formulas

The system SHALL support array formulas (CSE formulas).

#### Scenario: Create CSE array formula
- GIVEN a range of cells
- WHEN `sheet.SetArrayFormula("=A1:A10*B1:B10", "C1:C10")` is called
- THEN the master cell has CellFormula with FormulaType Array and Reference
- AND other cells in the range have DataType "str" (calculated)

#### Scenario: Create dynamic array formula
- GIVEN a Cell (Office 365+)
- WHEN `cell.SetFormula("=SORT(A1:A10)")` is called
- THEN the formula is stored as normal
- AND Excel handles the spilling behavior

### Requirement: Data Table Formulas

The system SHALL support data table formulas for what-if analysis.

#### Scenario: Create one-variable data table
- GIVEN a data table range
- WHEN `sheet.SetDataTableFormula("B1:B10", rowInput: "A1")` is called
- THEN CellFormula with FormulaType DataTable is created
- AND RowInput or ColumnInput attributes are set

#### Scenario: Create two-variable data table
- GIVEN a data table range
- WHEN `sheet.SetDataTableFormula("B2:D4", rowInput: "A1", colInput: "B1")` is called
- THEN both RowInput and ColumnInput are set

### Requirement: Defined Names

The system SHALL support named ranges and formulas.

#### Scenario: Create global defined name
- GIVEN a Workbook element
- WHEN `workbook.AddDefinedName("SalesData", "Sheet1!$A$1:$D$100")` is called
- THEN a DefinedName element is added to DefinedNames
- AND LocalSheetId is not set (global scope)

#### Scenario: Create local defined name
- GIVEN a Workbook element
- WHEN `workbook.AddDefinedName("_FilterDatabase", "Sheet1!$A$1:$D$100", localSheetId: 0)` is called
- THEN LocalSheetId is set to 0
- AND the name is scoped to Sheet1

#### Scenario: Create formula-based defined name
- GIVEN a Workbook element
- WHEN `workbook.AddDefinedName("TaxRate", "=0.0875")` is called
- THEN the DefinedName contains the formula
- AND it can be used in other formulas

#### Scenario: Access defined name
- GIVEN a Workbook with defined names
- WHEN `workbook.DefinedName("SalesData")` is called
- THEN the DefinedName element is returned

### Requirement: External References

The system SHALL support references to external workbooks.

#### Scenario: Parse external reference
- GIVEN a formula "[Budget.xlsx]Sheet1!$A$1"
- WHEN the formula is stored
- THEN an ExternalWorkbookPart may be created
- AND the external reference is preserved in the formula

#### Scenario: External link relationship
- GIVEN a workbook with external references
- WHEN the ExternalWorkbookPart is accessed
- THEN external workbook paths and sheet names are available

### Requirement: Structured References

The system SHALL support structured references to Excel tables.

#### Scenario: Table column reference
- GIVEN a formula "=SUM(Table1[Sales])"
- WHEN the formula is stored
- THEN the structured reference is preserved
- AND resolves to the table column range

#### Scenario: Table special references
- GIVEN formulas with [#All], [#Data], [#Headers], [#Totals], [@Column]
- WHEN the formulas are stored
- THEN the special references are preserved

### Requirement: Formula Types

The system SHALL support all formula types through the FormulaType enum.

#### Scenario: Normal formula type
- GIVEN a CellFormula element
- WHEN FormulaType is Normal or not set
- THEN it represents a standard formula

#### Scenario: Array formula type
- GIVEN a CellFormula element
- WHEN FormulaType is Array
- THEN Reference attribute specifies the array range

#### Scenario: DataTable formula type
- GIVEN a CellFormula element
- WHEN FormulaType is DataTable
- THEN RowInput, ColumnInput, Input2Deleted, DataTable2D attributes may be set

#### Scenario: Shared formula type
- GIVEN a CellFormula element
- WHEN FormulaType is Shared
- THEN SharedIndex references the formula group

### Requirement: Calculation Properties

The system SHALL support workbook calculation properties.

#### Scenario: Set calculation mode
- GIVEN a Workbook element
- WHEN `CalcPr.CalcMode` is set
- THEN values Auto, Manual, AutoNoTable are supported

#### Scenario: Set calculation iteration
- GIVEN circular reference handling is needed
- WHEN `CalcPr.Iterate` is set to true
- THEN IterateCount and IterateDelta can be configured

#### Scenario: Force full calculation
- GIVEN a workbook that needs recalculation
- WHEN `CalcPr.FullCalcOnLoad` is set to true
- THEN Excel recalculates all formulas when opening

### Requirement: Calculation Chain

The system SHALL support calculation chain management.

#### Scenario: Access calculation chain
- GIVEN a workbook with formulas
- WHEN CalculationChainPart is accessed
- THEN the ordered list of formula cells is available

#### Scenario: Calculation chain entry
- GIVEN a CalculationChain element
- WHEN C (cell) elements are enumerated
- THEN each has CellReference, SheetId, and optional flags (NewThread, Array, etc.)

