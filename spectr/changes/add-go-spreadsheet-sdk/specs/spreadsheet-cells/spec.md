## ADDED Requirements

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
