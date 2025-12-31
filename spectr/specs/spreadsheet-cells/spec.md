# Spreadsheet Cells Specification

## Requirements

### Requirement: Cell Reference System

The system SHALL provide utilities for parsing and creating cell references.

#### Scenario: Parse A1-style reference
- GIVEN a string "A1"
- WHEN `ParseCellRef("A1")` is called
- THEN a CellRef with Col=1, Row=1 is returned

#### Scenario: Parse absolute reference
- GIVEN a string "$A$1"
- WHEN `ParseCellRef("$A$1")` is called
- THEN a CellRef with Col=1, Row=1, AbsCol=true, AbsRow=true is returned

#### Scenario: Parse extended column reference
- GIVEN a string "XFD1"
- WHEN `ParseCellRef("XFD1")` is called
- THEN a CellRef with Col=16384 (max column), Row=1 is returned

#### Scenario: Convert column number to name
- GIVEN column number 1
- WHEN `ColumnName(1)` is called
- THEN "A" is returned

#### Scenario: Convert column number to extended name
- GIVEN column number 27
- WHEN `ColumnName(27)` is called
- THEN "AA" is returned

#### Scenario: Convert column name to number
- GIVEN column name "AA"
- WHEN `ColumnIndex("AA")` is called
- THEN 27 is returned

### Requirement: Range Reference System

The system SHALL provide utilities for parsing and creating range references.

#### Scenario: Parse range reference
- GIVEN a string "A1:C10"
- WHEN `ParseRangeRef("A1:C10")` is called
- THEN a RangeRef with Start=(1,1) and End=(3,10) is returned

#### Scenario: Parse single cell as range
- GIVEN a string "A1"
- WHEN `ParseRangeRef("A1")` is called
- THEN a RangeRef with Start=End=(1,1) is returned

#### Scenario: Iterate cells in range
- GIVEN a RangeRef "A1:B2"
- WHEN the range is iterated
- THEN cells A1, B1, A2, B2 are yielded in order

### Requirement: Cell Value Types

The system SHALL support all Excel cell value types.

#### Scenario: Set numeric value
- GIVEN a Cell element
- WHEN `cell.SetNumber(42.5)` is called
- THEN CellValue contains "42.5"
- AND DataType is not set (defaults to number)

#### Scenario: Set string value using shared strings
- GIVEN a Cell element and shared string table
- WHEN `cell.SetString("Hello")` is called
- THEN the string is added to shared string table
- AND CellValue contains the string index
- AND DataType is "s" (shared string)

#### Scenario: Set inline string value
- GIVEN a Cell element
- WHEN `cell.SetInlineString("Hello")` is called
- THEN InlineString element is set
- AND DataType is "inlineStr"
- AND CellValue is not used

#### Scenario: Set boolean value
- GIVEN a Cell element
- WHEN `cell.SetBoolean(true)` is called
- THEN CellValue contains "1"
- AND DataType is "b"

#### Scenario: Set date value
- GIVEN a Cell element
- WHEN `cell.SetDate(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))` is called
- THEN CellValue contains the Excel serial date number
- AND a date format should be applied via style

#### Scenario: Set error value
- GIVEN a Cell element
- WHEN `cell.SetError("#DIV/0!")` is called
- THEN CellValue contains "#DIV/0!"
- AND DataType is "e"

### Requirement: Cell Value Reading

The system SHALL support reading cell values with type detection.

#### Scenario: Read numeric value
- GIVEN a Cell with numeric value "42.5"
- WHEN `cell.GetNumber()` is called
- THEN float64 42.5 is returned

#### Scenario: Read shared string value
- GIVEN a Cell with DataType "s" and CellValue "5"
- WHEN `cell.GetString(sharedStrings)` is called
- THEN the string at index 5 from shared string table is returned

#### Scenario: Read boolean value
- GIVEN a Cell with DataType "b" and CellValue "1"
- WHEN `cell.GetBoolean()` is called
- THEN true is returned

#### Scenario: Read date value
- GIVEN a Cell with numeric value and date format
- WHEN `cell.GetDate()` is called
- THEN the time.Time value is returned

#### Scenario: Detect cell value type
- GIVEN a Cell element
- WHEN `cell.ValueType()` is called
- THEN the appropriate CellType enum value is returned

### Requirement: Rich Text in Cells

The system SHALL support rich text (multiple formats) within cells.

#### Scenario: Create rich text cell
- GIVEN a Cell element
- WHEN inline string with multiple Run elements is set
- THEN each Run can have different RunProperties (font, color, etc.)

#### Scenario: Access rich text runs
- GIVEN a Cell with rich text
- WHEN `cell.InlineString.Runs()` is accessed
- THEN all Run elements with Text and RunProperties are returned

### Requirement: Cell Hyperlinks

The system SHALL support hyperlinks in cells.

#### Scenario: Add internal hyperlink
- GIVEN a Worksheet
- WHEN `sheet.AddHyperlink("A1", "#'Sheet2'!A1")` is called
- THEN a Hyperlink element is added with Location attribute

#### Scenario: Add external hyperlink
- GIVEN a Worksheet
- WHEN `sheet.AddHyperlink("A1", "https://example.com")` is called
- THEN a Hyperlink element is added with r:id attribute
- AND an external relationship is created

### Requirement: Cell Comments

The system SHALL support comments attached to cells.

#### Scenario: Add comment to cell
- GIVEN a WorksheetPart
- WHEN `sheet.AddComment("A1", "Author", "Comment text")` is called
- THEN a Comment element is added to CommentsPart
- AND a VML shape is created for the comment balloon

#### Scenario: Access cell comment
- GIVEN a Cell with a comment
- WHEN `sheet.GetComment("A1")` is called
- THEN the Comment element with author and text is returned

### Requirement: Threaded Comments

The system SHALL support threaded comments (Office 2019+).

#### Scenario: Add threaded comment
- GIVEN a WorksheetPart
- WHEN `sheet.AddThreadedComment("A1", "Author", "Reply text", parentId)` is called
- THEN a ThreadedComment element is added
- AND it references the parent comment if specified

### Requirement: Cell Metadata

The system SHALL support cell metadata for rich data types.

#### Scenario: Access cell metadata
- GIVEN a Cell with value metadata
- WHEN `cell.ValueMetadata` is accessed
- THEN the metadata index referencing CellMetadataPart is returned

### Requirement: Shared String Table

The system SHALL provide efficient shared string management.

#### Scenario: Add string to table
- GIVEN a SharedStringTablePart
- WHEN `sst.AddString("Hello")` is called
- THEN the string is added if not exists
- AND the index is returned

#### Scenario: Get string by index
- GIVEN a SharedStringTablePart with strings
- WHEN `sst.String(5)` is called
- THEN the string at index 5 is returned

#### Scenario: Check string exists
- GIVEN a SharedStringTablePart
- WHEN `sst.IndexOf("Hello")` is called
- THEN the index is returned if exists, or -1 if not

#### Scenario: Deduplicate strings
- GIVEN a SharedStringTablePart
- WHEN the same string is added twice
- THEN only one entry exists
- AND both calls return the same index

### Requirement: Excel Date Serial Numbers

The system SHALL correctly handle Excel date serial numbers.

#### Scenario: Convert date to serial
- GIVEN a time.Time value
- WHEN `ExcelDateSerial(t)` is called
- THEN the Excel serial number is returned
- AND the 1900 date system is used by default

#### Scenario: Convert serial to date
- GIVEN an Excel serial number
- WHEN `DateFromExcelSerial(serial)` is called
- THEN the time.Time value is returned

#### Scenario: Handle 1904 date system
- GIVEN a workbook using 1904 date system
- WHEN date conversion is performed
- THEN the 1904 base date is used instead of 1900

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
