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

### Requirement: Formula Tokenization
The system SHALL provide utilities for tokenizing Excel formulas into constituent parts.

#### Scenario: Tokenize simple formula
- GIVEN a formula "=A1+B1"
- WHEN `tokenizeFormula("=A1+B1")` is called
- THEN tokens [Operator "=", Reference "A1", Operator "+", Reference "B1"] are returned

#### Scenario: Tokenize formula with function
- GIVEN a formula "=SUM(A1:A10)"
- WHEN `tokenizeFormula("=SUM(A1:A10)")` is called
- THEN tokens [Operator "=", Function "SUM", Paren "(", Reference "A1:A10", Paren ")"] are returned

#### Scenario: Tokenize formula with string literal
- GIVEN a formula "=\"Value: \"&A1"
- WHEN `tokenizeFormula("=\"Value: \"&A1")` is called
- THEN the string literal "Value: " is tokenized as Literal (not parsed for references)
- AND the reference A1 is tokenized as Reference

#### Scenario: Tokenize formula with numeric literal
- GIVEN a formula "=A1*2+100"
- WHEN `tokenizeFormula("=A1*2+100")` is called
- THEN numeric literals "2" and "100" are tokenized as Literal type

#### Scenario: Tokenize complex nested formula
- GIVEN a formula "=IF(A1>0,SUM(B1:B10),0)"
- WHEN `tokenizeFormula("=IF(A1>0,SUM(B1:B10),0)")` is called
- THEN all nested function calls, references, and operators are tokenized correctly

#### Scenario: Tokenize formula with cross-sheet reference
- GIVEN a formula "=Sheet2!A1+Sheet3!B2"
- WHEN `tokenizeFormula("=Sheet2!A1+Sheet3!B2")` is called
- THEN "Sheet2!A1" and "Sheet3!B2" are tokenized as Reference type

#### Scenario: Tokenize formula without leading equals
- GIVEN a formula "A1+B1" (no leading =)
- WHEN `tokenizeFormula("A1+B1")` is called
- THEN tokens are returned without Operator "=" at start

#### Scenario: Tokenize formula with all operator types
- GIVEN a formula "=A1+B1-C1*D1/E1&F1=G1<>H1<=I1>=J1"
- WHEN tokenized
- THEN operators +, -, *, /, &, =, <>, <=, >= are all tokenized correctly

#### Scenario: Tokenize formula with parentheses
- GIVEN a formula "=(A1+B1)*(C1+D1)"
- WHEN tokenized
- THEN opening and closing parentheses are tokenized as Paren type

#### Scenario: Tokenize formula with commas
- GIVEN a formula "=SUM(A1,B1,C1)"
- WHEN tokenized
- THEN commas are tokenized as Comma type

### Requirement: Formula Reference Extraction
The system SHALL provide utilities for extracting cell references from formulas.

#### Scenario: Extract simple references
- GIVEN a formula "=A1+B2+C3"
- WHEN references are extracted
- THEN references [A1, B2, C3] are returned

#### Scenario: Extract range references
- GIVEN a formula "=SUM(A1:A10)+SUM(B1:B10)"
- WHEN references are extracted
- THEN range references [A1:A10, B1:B10] are returned

#### Scenario: Extract cross-sheet references
- GIVEN a formula "=Sheet1!A1+Sheet2!B2"
- WHEN references are extracted
- THEN references [Sheet1!A1, Sheet2!B2] are returned

#### Scenario: Ignore references in string literals
- GIVEN a formula "=\"Cell A1 contains\"&A1"
- WHEN references are extracted
- THEN only [A1] is extracted (not "A1" from string)

#### Scenario: Extract absolute references
- GIVEN a formula "=$A$1+$A1+A$1"
- WHEN references are extracted
- THEN references [$A$1, $A1, A$1] are returned with absolute flags preserved

#### Scenario: Extract references from nested functions
- GIVEN a formula "=IF(A1>0,SUM(B1:B10),VLOOKUP(C1,D1:E10,2,FALSE))"
- WHEN references are extracted
- THEN all references [A1, B1:B10, C1, D1:E10] are returned

### Requirement: Formula Rewriting
The system SHALL provide utilities for rewriting formulas after row/column operations.

#### Scenario: Rewrite formula after row insert
- GIVEN a formula "=SUM(A1:A10)" in cell B1
- WHEN rewritten for InsertRows(5, 3) operation
- THEN the formula becomes "=SUM(A1:A13)"

#### Scenario: Rewrite formula after row delete
- GIVEN a formula "=SUM(A1:A10)" in cell B1
- WHEN rewritten for DeleteRows(8, 3) operation
- THEN the formula becomes "=SUM(A1:A7)"

#### Scenario: Rewrite formula after column insert
- GIVEN a formula "=SUM(A1:E1)" in cell F1
- WHEN rewritten for InsertColumns(3, 2) operation
- THEN the formula becomes "=SUM(A1:G1)"

#### Scenario: Rewrite formula after column delete
- GIVEN a formula "=SUM(A1:H1)" in cell I1
- WHEN rewritten for DeleteColumns(6, 3) operation
- THEN the formula becomes "=SUM(A1:E1)"

#### Scenario: Preserve absolute references in rewrite
- GIVEN a formula "=$A$1+A10"
- WHEN rewritten for InsertRows(5, 3) operation
- THEN the formula becomes "=$A$1+A13" (absolute ref unchanged, relative shifted)

#### Scenario: Preserve string literals in rewrite
- GIVEN a formula "=\"Total: \"&SUM(A1:A10)"
- WHEN rewritten for InsertRows(5, 3) operation
- THEN the formula becomes "=\"Total: \"&SUM(A1:A13)" (string unchanged, range shifted)

#### Scenario: Rewrite complex nested formula
- GIVEN a formula "=IF(A1>0,SUM(B1:B10),AVERAGE(C1:C10))"
- WHEN rewritten for InsertRows(5, 3) operation
- THEN the formula becomes "=IF(A1>0,SUM(B1:B13),AVERAGE(C1:C13))"

#### Scenario: Rewrite formula with deleted reference
- GIVEN a formula "=A5+A6+A7"
- WHEN rewritten for DeleteRows(5, 3) operation (delete rows 5, 6, 7)
- THEN the formula becomes "=#REF!+#REF!+#REF!" (all refs deleted)

#### Scenario: Rewrite formula with partially deleted range
- GIVEN a formula "=SUM(A5:A10)"
- WHEN rewritten for DeleteRows(7, 2) operation (delete rows 7, 8)
- THEN the formula becomes "=SUM(A5:A8)" (range shrinks, references adjusted)

#### Scenario: Rewrite cross-sheet formula
- GIVEN a formula "=Sheet1!A5+Sheet2!A10" in Sheet3
- WHEN rewritten for InsertRows(3, 2) operation on Sheet1
- THEN the formula becomes "=Sheet1!A7+Sheet2!A10" (only Sheet1 ref updated)

#### Scenario: Rewrite formula with mixed absolute references
- GIVEN a formula "=$A5+A$5+$A$5"
- WHEN rewritten for InsertRows(3, 2) operation
- THEN the formula becomes "=$A7+A$5+$A$5" (only relative components shifted)

#### Scenario: Preserve function names during rewrite
- GIVEN a formula "=SUM(A1:A10)+AVERAGE(B1:B10)"
- WHEN rewritten for InsertRows(5, 3) operation
- THEN function names SUM and AVERAGE are preserved
- AND only the range references are updated

### Requirement: Formula Reconstruction
The system SHALL provide utilities for reconstructing formulas from tokens.

#### Scenario: Reconstruct simple formula
- GIVEN tokens [Operator "=", Reference "A1", Operator "+", Reference "B1"]
- WHEN `reconstructFormula(tokens)` is called
- THEN "=A1+B1" is returned

#### Scenario: Reconstruct formula with function
- GIVEN tokens for =SUM(A1:A10)
- WHEN reconstructed
- THEN "=SUM(A1:A10)" is returned

#### Scenario: Reconstruct formula with spacing
- GIVEN tokens for a complex formula
- WHEN reconstructed
- THEN appropriate spacing is added (e.g., after commas, around operators)

#### Scenario: Ensure leading equals sign
- GIVEN tokens without leading equals operator
- WHEN reconstructed
- THEN a leading "=" is added to make it a valid formula

#### Scenario: Reconstruct formula preserves string literals
- GIVEN tokens including a string literal "Value: "
- WHEN reconstructed
- THEN the string literal is preserved with correct quoting

#### Scenario: Reconstruct nested formula
- GIVEN tokens for =IF(A1>0,SUM(B1:B10),0)
- WHEN reconstructed
- THEN "=IF(A1>0,SUM(B1:B10),0)" is returned with correct nesting

### Requirement: Named Range Reference Updates
The system SHALL update named range definitions when rows or columns are inserted or deleted.

#### Scenario: Update named range after row insert
- GIVEN a named range "SalesData = Sheet1!A1:D10"
- WHEN InsertRows(5, 3) is called on Sheet1
- THEN the named range becomes "SalesData = Sheet1!A1:D13"

#### Scenario: Update named range after row delete
- GIVEN a named range "SalesData = Sheet1!A1:D10"
- WHEN DeleteRows(8, 3) is called on Sheet1
- THEN the named range becomes "SalesData = Sheet1!A1:D7"

#### Scenario: Update named range after column insert
- GIVEN a named range "QuarterData = Sheet1!A1:D100"
- WHEN InsertColumns(3, 2) is called on Sheet1
- THEN the named range becomes "QuarterData = Sheet1!A1:F100"

#### Scenario: Update named range after column delete
- GIVEN a named range "QuarterData = Sheet1!A1:H100"
- WHEN DeleteColumns(6, 3) is called on Sheet1
- THEN the named range becomes "QuarterData = Sheet1!A1:E100"

#### Scenario: Preserve absolute named range boundaries
- GIVEN a named range "Data = Sheet1!$A$1:$D$10"
- WHEN InsertRows(5, 3) is called on Sheet1
- THEN the named range remains "Data = Sheet1!$A$1:$D$10" (absolute refs)

#### Scenario: Update formula-based named range
- GIVEN a named range "Total = SUM(Sheet1!A1:A10)"
- WHEN InsertRows(5, 3) is called on Sheet1
- THEN the named range becomes "Total = SUM(Sheet1!A1:A13)"

#### Scenario: Update local-scoped named range
- GIVEN a sheet-scoped named range "Data = Sheet1!A1:D10" (LocalSheetId=0)
- WHEN InsertRows(5, 3) is called on Sheet1
- THEN the named range is updated to "Sheet1!A1:D13"

#### Scenario: Preserve named range for different worksheet
- GIVEN a named range "Data = Sheet2!A1:D10"
- WHEN InsertRows(5, 3) is called on Sheet1
- THEN the named range remains unchanged (different sheet)

### Requirement: Shared Formula Updates
The system SHALL update shared formulas when rows or columns are inserted or deleted.

**Background**: Shared formulas in Excel use a master cell formula that is applied to a range of cells. The master cell contains the actual formula text (e.g., "=A1+1") and the `ref` attribute specifies the range where the formula applies (e.g., "B1:B100"). Child cells reference the master using the `si` (shared index) attribute without duplicating the formula text.

**ECMA-376 attributes**:
- `t="shared"` - Indicates this is a shared formula
- `si="N"` - Shared index (unique identifier for the shared formula group)
- `ref="A1:A10"` - Range of cells sharing this formula (only on master cell)

**Master cell** (contains formula text + ref + si):
```xml
<c r="B1">
  <f t="shared" ref="B1:B100" si="0">=A1+1</f>
  <v>42</v>
</c>
```

**Child cells** (reference master via si):
```xml
<c r="B2">
  <f t="shared" si="0"/>
  <v>43</v>
</c>
```

#### Scenario: Identify shared formula master cell
- GIVEN a formula element with t="shared" and ref="B1:B100"
- WHEN checking if it's a master cell
- THEN it's identified as master because ref attribute is present

#### Scenario: Identify shared formula child cell
- GIVEN a formula element with t="shared" and si="0" but no ref attribute
- WHEN checking if it's a child cell
- THEN it's identified as child because ref attribute is absent

#### Scenario: Update shared formula master cell
- GIVEN a shared formula "=A1+1" in range B1:B100 (master cell B1 with ref="B1:B100" si="0")
- WHEN InsertRows(50, 10) is called
- THEN the master cell formula becomes "=A1+1" (formula text unchanged, A1 is before insert point)
- AND the master cell ref attribute becomes "B1:B110" (range extended)
- AND the si attribute remains "0" (unchanged)

#### Scenario: Update shared formula reference range
- GIVEN a shared formula with reference range B1:B100
- WHEN InsertRows(50, 10) is called
- THEN the reference range becomes B1:B110

#### Scenario: Adjust shared formula for deleted rows
- GIVEN a shared formula with reference range B10:B100
- WHEN DeleteRows(50, 10) is called
- THEN the reference range becomes B10:B90

#### Scenario: Update shared formula master when formula contains affected reference
- GIVEN a shared formula "=SUM(A1:A100)" in range B1:B100 (master cell B1)
- WHEN InsertRows(50, 10) is called
- THEN the formula text becomes "=SUM(A1:A110)" (range inside formula updated)
- AND the ref attribute becomes "B1:B110" (master range extended)

#### Scenario: Preserve shared index during updates
- GIVEN multiple shared formula groups (si="0", si="1", si="2")
- WHEN rows are inserted or deleted
- THEN all si values are preserved (shared index never changes during row/column operations)

#### Scenario: Update shared formula when master cell is in affected range
- GIVEN a shared formula with master cell B50 and range B50:B100
- WHEN DeleteRows(50, 1) is called (delete row 50)
- THEN the master cell moves to B50 (was B51)
- AND the ref becomes "B50:B99"
- AND the formula text is rewritten for the new position

#### Scenario: Handle shared formula child cells
- GIVEN child cells B2:B100 with t="shared" si="0" (no ref attribute)
- WHEN InsertRows(50, 10) is called
- THEN child cells remain t="shared" si="0" (no changes to child cell formulas)
- AND their positions shift per normal row shift logic

### Requirement: Array Formula Updates
The system SHALL update array formulas when rows or columns are inserted or deleted.

**Background**: Array formulas in Excel come in two varieties:
1. **CSE (Ctrl+Shift+Enter) array formulas**: Legacy array formulas that occupy a fixed range. All cells in the range have the same formula with `t="array"` and `ref` attributes. Only the top-left cell contains the formula text.
2. **Dynamic array formulas**: Modern Excel (Office 365+) formulas that automatically spill to adjacent cells. The formula is only in the anchor cell.

**ECMA-376 attributes for CSE arrays**:
- `t="array"` - Indicates this is an array formula
- `ref="C1:C10"` - Range occupied by the array formula (all cells in range)
- Formula text only in top-left cell, other cells have empty formula elements

**CSE array formula** (top-left cell):
```xml
<c r="C1">
  <f t="array" ref="C1:C10">=A1:A10*B1:B10</f>
  <v>42</v>
</c>
```

**CSE array formula** (other cells in range):
```xml
<c r="C2">
  <f t="array" ref="C1:C10"/>
  <v>43</v>
</c>
```

**Dynamic array formula** (anchor cell only, Excel 365+):
```xml
<c r="B1">
  <f>=SORT(A1:A10)</f>
  <v>spilled value</v>
</c>
```

#### Scenario: Identify CSE array formula anchor cell
- GIVEN a formula element with t="array", ref="C1:C10", and formula text "=A1:A10*B1:B10"
- WHEN checking if it's the anchor cell
- THEN it's identified as anchor because it contains formula text

#### Scenario: Identify CSE array formula member cell
- GIVEN a formula element with t="array", ref="C1:C10", but no formula text
- WHEN checking if it's a member cell
- THEN it's identified as member because formula text is absent

#### Scenario: Update CSE array formula range
- GIVEN an array formula "=A1:A10*B1:B10" with ref="C1:C10"
- WHEN InsertRows(5, 3) is called
- THEN the formula becomes "=A1:A13*B1:B13" (references in formula updated)
- AND the ref attribute becomes "C1:C13" (array range extended)
- AND all member cells update their ref to "C1:C13"

#### Scenario: Update dynamic array formula
- GIVEN a dynamic array formula "=SORT(A1:A10)" in cell B1 (no t or ref attributes)
- WHEN InsertRows(5, 3) is called
- THEN the formula becomes "=SORT(A1:A13)"
- AND no ref attribute is added (dynamic arrays don't use ref)

#### Scenario: Preserve array formula type during update
- GIVEN an array formula with t="array"
- WHEN updated for row/column operations
- THEN the t attribute remains "array"
- AND the formula type is preserved

#### Scenario: Update array formula when anchor cell is in affected range
- GIVEN an array formula with anchor C50 and ref="C50:C100"
- WHEN DeleteRows(50, 1) is called (delete row 50)
- THEN the anchor cell moves to C50 (was C51)
- AND the ref becomes "C50:C99"
- AND all member cells update their ref

#### Scenario: Handle array formula deletion when range partially deleted
- GIVEN an array formula with ref="C50:C60"
- WHEN DeleteRows(55, 3) is called (delete rows 55, 56, 57)
- THEN the ref becomes "C50:C57" (3 rows removed from range)
- AND the array formula continues to function

#### Scenario: Handle array formula deletion when anchor deleted
- GIVEN an array formula with anchor C50 and ref="C50:C60"
- WHEN DeleteRows(50, 1) is called (delete anchor row 50)
- THEN the new anchor becomes C50 (was C51)
- AND the ref becomes "C50:C59"
- AND the formula text moves to the new anchor cell

### Requirement: Data Table Formula Updates
The system SHALL update data table formulas when rows or columns are inserted or deleted.

#### Scenario: Update one-variable data table row input
- GIVEN a data table with RowInput="A1" and range B1:B10
- WHEN InsertRows(1, 2) is called
- THEN RowInput becomes "A3"

#### Scenario: Update two-variable data table inputs
- GIVEN a data table with RowInput="A1" and ColumnInput="B1"
- WHEN InsertRows(1, 2) is called
- THEN RowInput becomes "A3" and ColumnInput becomes "B3"

#### Scenario: Update data table range
- GIVEN a data table with range B2:D10
- WHEN InsertRows(5, 3) is called
- THEN the data table range becomes B2:D13

### Requirement: Error Reference Handling
The system SHALL generate appropriate error references when cells are deleted.

#### Scenario: Generate REF error for deleted cell
- GIVEN a formula "=A5"
- WHEN DeleteRows(5, 1) is called (delete row 5)
- THEN the formula becomes "=#REF!"

#### Scenario: Generate REF error for deleted range
- GIVEN a formula "=SUM(A5:A7)"
- WHEN DeleteRows(5, 3) is called (delete rows 5, 6, 7)
- THEN the formula becomes "=SUM(#REF!)"

#### Scenario: Partial REF error in formula
- GIVEN a formula "=A4+A5+A8"
- WHEN DeleteRows(5, 1) is called (delete row 5)
- THEN the formula becomes "=A4+#REF!+A7"

#### Scenario: Preserve non-deleted references
- GIVEN a formula "=A1+A10"
- WHEN DeleteRows(5, 3) is called (delete rows 5, 6, 7)
- THEN the formula becomes "=A1+A7" (A1 unchanged, A10 shifted to A7)

### Requirement: Formula Validation After Updates
The system SHALL validate formulas after rewriting to ensure correctness.

#### Scenario: Validate rewritten formula syntax
- GIVEN a formula rewritten after row insert
- WHEN validation is performed
- THEN the formula syntax is correct (balanced parentheses, valid operators)

#### Scenario: Validate reference bounds after update
- GIVEN a formula with references updated
- WHEN validation is performed
- THEN all references are within Excel bounds (columns <= XFD, rows <= 1048576)

#### Scenario: Detect circular references after update
- GIVEN formulas that may create circular references after update
- WHEN validation is performed
- THEN circular references are detected and reported (optional validation)
