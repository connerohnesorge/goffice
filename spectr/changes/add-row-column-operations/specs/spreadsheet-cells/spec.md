# Spreadsheet Cells Specification (Delta)

## ADDED Requirements

### Requirement: Cell Reference Parsing
The system SHALL provide utilities for parsing all cell reference formats used in Excel.

#### Scenario: Parse simple A1 reference
- GIVEN a string "A1"
- WHEN `ParseCellRef("A1")` is called
- THEN a CellReference with Column=1, Row=1, ColumnAbs=false, RowAbs=false is returned

#### Scenario: Parse absolute reference
- GIVEN a string "$A$1"
- WHEN `ParseCellRef("$A$1")` is called
- THEN a CellReference with Column=1, Row=1, ColumnAbs=true, RowAbs=true is returned

#### Scenario: Parse mixed absolute reference (column)
- GIVEN a string "$A1"
- WHEN `ParseCellRef("$A1")` is called
- THEN a CellReference with ColumnAbs=true, RowAbs=false is returned

#### Scenario: Parse mixed absolute reference (row)
- GIVEN a string "A$1"
- WHEN `ParseCellRef("A$1")` is called
- THEN a CellReference with ColumnAbs=false, RowAbs=true is returned

#### Scenario: Parse extended column reference
- GIVEN a string "XFD1048576" (maximum Excel cell)
- WHEN `ParseCellRef("XFD1048576")` is called
- THEN a CellReference with Column=16384, Row=1048576 is returned

#### Scenario: Parse cross-sheet reference
- GIVEN a string "Sheet2!A1"
- WHEN `ParseCellRef("Sheet2!A1")` is called
- THEN a CellReference with Sheet="Sheet2", Column=1, Row=1 is returned

#### Scenario: Parse quoted sheet name reference
- GIVEN a string "'Sheet Name'!A1"
- WHEN `ParseCellRef("'Sheet Name'!A1")` is called
- THEN a CellReference with Sheet="Sheet Name", Column=1, Row=1 is returned

#### Scenario: Parse quoted sheet with special characters
- GIVEN a string "'Sales (2024)'!A1"
- WHEN `ParseCellRef("'Sales (2024)'!A1")` is called
- THEN a CellReference with Sheet="Sales (2024)" is returned

#### Scenario: Parse R1C1 absolute reference
- GIVEN a string "R5C10"
- WHEN `ParseCellRef("R5C10")` is called
- THEN a CellReference with Row=5, Column=10, IsR1C1=true is returned

#### Scenario: Parse R1C1 relative reference
- GIVEN a string "R[-1]C[2]"
- WHEN `ParseCellRef("R[-1]C[2]")` is called
- THEN a CellReference with R1C1RelRow=-1, R1C1RelCol=2, IsR1C1=true is returned

#### Scenario: Parse R1C1 mixed reference
- GIVEN a string "R5C[-2]"
- WHEN `ParseCellRef("R5C[-2]")` is called
- THEN a CellReference with Row=5 (absolute), R1C1RelCol=-2 (relative) is returned

#### Scenario: Invalid reference returns error
- GIVEN a string "@@@" (invalid reference)
- WHEN `ParseCellRef("@@@")` is called
- THEN an error is returned

#### Scenario: Out of bounds reference returns error
- GIVEN a string "XFE1" (column > 16384)
- WHEN `ParseCellRef("XFE1")` is called
- THEN an error is returned

### Requirement: Range Reference Parsing
The system SHALL provide utilities for parsing range references.

#### Scenario: Parse simple range
- GIVEN a string "A1:D10"
- WHEN `ParseRangeRef("A1:D10")` is called
- THEN a RangeReference with Start=(A1) and End=(D10) is returned

#### Scenario: Parse single cell as range
- GIVEN a string "A1"
- WHEN `ParseRangeRef("A1")` is called
- THEN a RangeReference with Start=End=(A1) is returned

#### Scenario: Parse absolute range
- GIVEN a string "$A$1:$D$10"
- WHEN `ParseRangeRef("$A$1:$D$10")` is called
- THEN a RangeReference with absolute start and end is returned

#### Scenario: Parse cross-sheet range
- GIVEN a string "Sheet1!A1:D10"
- WHEN `ParseRangeRef("Sheet1!A1:D10")` is called
- THEN a RangeReference with Sheet="Sheet1" for both start and end is returned

#### Scenario: Parse whole column range
- GIVEN a string "A:A"
- WHEN `ParseRangeRef("A:A")` is called
- THEN a RangeReference representing entire column A is returned

#### Scenario: Parse whole row range
- GIVEN a string "5:5"
- WHEN `ParseRangeRef("5:5")` is called
- THEN a RangeReference representing entire row 5 is returned

#### Scenario: Parse multi-column range
- GIVEN a string "C:F"
- WHEN `ParseRangeRef("C:F")` is called
- THEN a RangeReference representing columns C through F is returned

#### Scenario: Invalid range returns error
- GIVEN a string "A1:@@@"
- WHEN `ParseRangeRef("A1:@@@")` is called
- THEN an error is returned

### Requirement: Cell Reference String Conversion
The system SHALL provide utilities to convert cell references back to string format.

#### Scenario: Convert simple reference to string
- GIVEN a CellReference with Column=1, Row=1
- WHEN `ref.String()` is called
- THEN "A1" is returned

#### Scenario: Convert absolute reference to string
- GIVEN a CellReference with Column=1, Row=1, ColumnAbs=true, RowAbs=true
- WHEN `ref.String()` is called
- THEN "$A$1" is returned

#### Scenario: Convert mixed absolute reference to string
- GIVEN a CellReference with Column=1, Row=1, ColumnAbs=true, RowAbs=false
- WHEN `ref.String()` is called
- THEN "$A1" is returned

#### Scenario: Convert cross-sheet reference to string
- GIVEN a CellReference with Sheet="Sheet2", Column=1, Row=1
- WHEN `ref.String()` is called
- THEN "Sheet2!A1" is returned

#### Scenario: Convert reference with quoted sheet name to string
- GIVEN a CellReference with Sheet="Sheet Name", Column=1, Row=1
- WHEN `ref.String()` is called
- THEN "'Sheet Name'!A1" is returned

#### Scenario: Convert range to string
- GIVEN a RangeReference with Start=(A1) and End=(D10)
- WHEN `rangeRef.String()` is called
- THEN "A1:D10" is returned

#### Scenario: Roundtrip parsing
- GIVEN a cell reference string "Sheet1!$A$5"
- WHEN parsed and converted back to string
- THEN the result is "Sheet1!$A$5" (identical)

### Requirement: Column Name Utilities
The system SHALL provide utilities for converting between column indices and names.

#### Scenario: Convert column index to name
- GIVEN column index 1
- WHEN `ColumnName(1)` is called
- THEN "A" is returned

#### Scenario: Convert extended column index to name
- GIVEN column index 27
- WHEN `ColumnName(27)` is called
- THEN "AA" is returned

#### Scenario: Convert maximum column index to name
- GIVEN column index 16384
- WHEN `ColumnName(16384)` is called
- THEN "XFD" is returned

#### Scenario: Convert column name to index
- GIVEN column name "A"
- WHEN `ColumnIndex("A")` is called
- THEN 1 is returned

#### Scenario: Convert extended column name to index
- GIVEN column name "AA"
- WHEN `ColumnIndex("AA")` is called
- THEN 27 is returned

#### Scenario: Convert maximum column name to index
- GIVEN column name "XFD"
- WHEN `ColumnIndex("XFD")` is called
- THEN 16384 is returned

#### Scenario: Case insensitive column name
- GIVEN column name "aa" (lowercase)
- WHEN `ColumnIndex("aa")` is called
- THEN 27 is returned (same as "AA")

#### Scenario: Invalid column name returns error
- GIVEN column name "XFE" (out of bounds)
- WHEN `ColumnIndex("XFE")` is called
- THEN an error is returned

### Requirement: Cell Reference Shifting
The system SHALL provide utilities for shifting cell references based on row/column operations.

#### Scenario: Shift reference for row insert
- GIVEN a CellReference A10
- WHEN shifted for InsertRows(5, 3) operation
- THEN the reference becomes A13 (shifted down by 3)

#### Scenario: Preserve absolute row reference on row insert
- GIVEN a CellReference $A$10
- WHEN shifted for InsertRows(5, 3) operation
- THEN the reference remains $A$10 (absolute row not shifted)

#### Scenario: Shift reference for row delete
- GIVEN a CellReference A10
- WHEN shifted for DeleteRows(5, 3) operation (delete rows 5, 6, 7)
- THEN the reference becomes A7 (shifted up by 3)

#### Scenario: Invalidate reference in deleted rows
- GIVEN a CellReference A6
- WHEN shifted for DeleteRows(5, 3) operation (delete rows 5, 6, 7)
- THEN the reference becomes invalid (A6 is deleted)
- AND error marker #REF! is used

#### Scenario: Shift reference for column insert
- GIVEN a CellReference E1
- WHEN shifted for InsertColumns(3, 2) operation
- THEN the reference becomes G1 (shifted right by 2)

#### Scenario: Preserve absolute column reference on column insert
- GIVEN a CellReference $E$1
- WHEN shifted for InsertColumns(3, 2) operation
- THEN the reference remains $E$1 (absolute column not shifted)

#### Scenario: Shift reference for column delete
- GIVEN a CellReference E1
- WHEN shifted for DeleteColumns(2, 2) operation (delete columns B, C)
- THEN the reference becomes C1 (shifted left by 2)

#### Scenario: Invalidate reference in deleted columns
- GIVEN a CellReference C1
- WHEN shifted for DeleteColumns(2, 3) operation (delete columns B, C, D)
- THEN the reference becomes invalid (C1 is deleted)
- AND error marker #REF! is used

#### Scenario: Shift mixed absolute reference
- GIVEN a CellReference $A5 (absolute column, relative row)
- WHEN shifted for InsertRows(3, 2) operation
- THEN the reference becomes $A7 (only row shifted)

#### Scenario: No shift for different worksheet
- GIVEN a CellReference Sheet2!A5
- WHEN shifted for InsertRows(3, 2) operation on Sheet1
- THEN the reference remains Sheet2!A5 (not affected)

#### Scenario: Shift cross-sheet reference for target worksheet
- GIVEN a CellReference Sheet1!A5
- WHEN shifted for InsertRows(3, 2) operation on Sheet1
- THEN the reference becomes Sheet1!A7 (shifted)

### Requirement: Range Reference Shifting
The system SHALL provide utilities for shifting range references.

#### Scenario: Shift range for row insert
- GIVEN a RangeReference A1:D10
- WHEN shifted for InsertRows(5, 3) operation
- THEN the range becomes A1:D13 (end row extended by 3)

#### Scenario: Shift range entirely after insert point
- GIVEN a RangeReference A10:D15
- WHEN shifted for InsertRows(5, 2) operation
- THEN the range becomes A12:D17 (both start and end shifted)

#### Scenario: Shift range for row delete
- GIVEN a RangeReference A1:D10
- WHEN shifted for DeleteRows(8, 3) operation
- THEN the range becomes A1:D7 (end row shrinks by 3)

#### Scenario: Shift range entirely after delete point
- GIVEN a RangeReference A10:D15
- WHEN shifted for DeleteRows(5, 3) operation
- THEN the range becomes A7:D12 (both start and end shifted)

#### Scenario: Shift range for column insert
- GIVEN a RangeReference A1:E10
- WHEN shifted for InsertColumns(3, 2) operation
- THEN the range becomes A1:G10 (end column extended by 2)

#### Scenario: Shift range for column delete
- GIVEN a RangeReference A1:H10
- WHEN shifted for DeleteColumns(6, 3) operation
- THEN the range becomes A1:E10 (end column shrinks by 3)

#### Scenario: Preserve absolute range boundaries
- GIVEN a RangeReference $A$1:$D$10
- WHEN shifted for InsertRows(5, 3) operation
- THEN the range remains $A$1:$D$10 (absolute references not shifted)

#### Scenario: Invalidate range with deleted boundaries
- GIVEN a RangeReference A5:D7
- WHEN shifted for DeleteRows(5, 3) operation (delete rows 5, 6, 7)
- THEN the range becomes invalid (both boundaries deleted)
- AND error marker #REF! is used

### Requirement: Cell Reference Comparison
The system SHALL provide utilities for comparing cell references.

#### Scenario: Compare equal references
- GIVEN two CellReference objects both representing A1
- WHEN `ref1.Equals(ref2)` is called
- THEN true is returned

#### Scenario: Compare different references
- GIVEN CellReference A1 and CellReference B2
- WHEN `ref1.Equals(ref2)` is called
- THEN false is returned

#### Scenario: Compare references with different absolute flags
- GIVEN CellReference A1 and CellReference $A$1
- WHEN `ref1.Equals(ref2)` is called
- THEN false is returned (different absolute flags)

#### Scenario: Compare references with different sheets
- GIVEN CellReference Sheet1!A1 and CellReference Sheet2!A1
- WHEN `ref1.Equals(ref2)` is called
- THEN false is returned (different sheets)

### Requirement: Cell Reference Validation
The system SHALL validate cell references against Excel constraints.

#### Scenario: Validate within bounds
- GIVEN a CellReference A1
- WHEN `ref.Validate()` is called
- THEN no error is returned

#### Scenario: Validate maximum cell
- GIVEN a CellReference XFD1048576
- WHEN `ref.Validate()` is called
- THEN no error is returned

#### Scenario: Validate column out of bounds
- GIVEN a CellReference with Column=16385 (> maximum 16384)
- WHEN `ref.Validate()` is called
- THEN an error is returned

#### Scenario: Validate row out of bounds
- GIVEN a CellReference with Row=1048577 (> maximum 1048576)
- WHEN `ref.Validate()` is called
- THEN an error is returned

#### Scenario: Validate zero indices
- GIVEN a CellReference with Column=0 or Row=0
- WHEN `ref.Validate()` is called
- THEN an error is returned (indices are 1-based)
