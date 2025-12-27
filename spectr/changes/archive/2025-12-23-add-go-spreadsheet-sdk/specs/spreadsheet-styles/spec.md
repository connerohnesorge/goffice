## ADDED Requirements

### Requirement: Stylesheet Root Element

The system SHALL provide a `Stylesheet` element as the root of WorkbookStylesPart.

#### Scenario: Access style collections
- GIVEN a Stylesheet element
- WHEN collections are accessed
- THEN NumberFormats, Fonts, Fills, Borders, CellStyleXfs, CellXfs, CellStyles, DifferentialFormats, TableStyles, Colors are available

### Requirement: Number Formats

The system SHALL support built-in and custom number formats.

#### Scenario: Built-in number format
- GIVEN a style with NumberFormatId 2
- THEN the format "0.00" is applied

#### Scenario: Custom number format
- GIVEN a custom format "$#,##0.00"
- WHEN added to NumberFormats
- THEN a NumberFormat element with Id >= 164 is created
- AND the FormatCode contains the custom format

#### Scenario: Common built-in formats
- GIVEN NumberFormatId values 0-49
- THEN standard formats are applied:
  - 0: General
  - 1: 0
  - 2: 0.00
  - 3: #,##0
  - 4: #,##0.00
  - 9: 0%
  - 10: 0.00%
  - 11: 0.00E+00
  - 14: m/d/yy
  - 22: m/d/yy h:mm

### Requirement: Font Definitions

The system SHALL support font definitions.

#### Scenario: Font properties
- GIVEN a Font element
- WHEN properties are accessed
- THEN Name, Size, Bold, Italic, Underline, Strike, Color, Family, Scheme, Charset, Condense, Extend, Shadow, Outline are available

#### Scenario: Create bold font
- GIVEN a Fonts collection
- WHEN a new Font with Bold=true is added
- THEN the font index can be referenced in CellXf

#### Scenario: Font color
- GIVEN a Font element
- WHEN Color is set
- THEN RGB, Theme, Tint, Indexed, Auto attributes are available

### Requirement: Fill Definitions

The system SHALL support fill (background) definitions.

#### Scenario: Fill types
- GIVEN a Fill element
- WHEN PatternFill or GradientFill child is accessed
- THEN the fill type is determined

#### Scenario: Solid fill
- GIVEN a PatternFill element
- WHEN PatternType is "solid"
- THEN FgColor specifies the solid fill color

#### Scenario: Pattern fill
- GIVEN a PatternFill element
- WHEN PatternType is a pattern (e.g., "gray125")
- THEN FgColor and BgColor specify pattern colors

#### Scenario: Gradient fill
- GIVEN a GradientFill element
- WHEN Type, Degree, Stop elements are set
- THEN a gradient fill is applied

### Requirement: Border Definitions

The system SHALL support border definitions.

#### Scenario: Border sides
- GIVEN a Border element
- WHEN sides are accessed
- THEN Left, Right, Top, Bottom, Diagonal are available

#### Scenario: Border style
- GIVEN a BorderPr element (Left, Right, etc.)
- WHEN Style is set
- THEN values None, Thin, Medium, Thick, Dotted, Dashed, Double, Hair, MediumDashed, DashDot, MediumDashDot, DashDotDot, MediumDashDotDot, SlantDashDot are supported

#### Scenario: Diagonal border
- GIVEN a Border element
- WHEN DiagonalUp or DiagonalDown is true
- THEN Diagonal border style is applied in that direction

### Requirement: Cell Formatting Records (CellXfs)

The system SHALL support cell formatting through CellXfs.

#### Scenario: CellXf structure
- GIVEN a CellXf element (in CellXfs collection)
- WHEN properties are accessed
- THEN NumberFormatId, FontId, FillId, BorderId, XfId, ApplyNumberFormat, ApplyFont, ApplyFill, ApplyBorder, ApplyAlignment, ApplyProtection, Alignment, Protection are available

#### Scenario: Apply style to cell
- GIVEN a CellXf with complete formatting
- WHEN a cell's StyleIndex is set to that CellXf index
- THEN all formatting is applied to the cell

#### Scenario: Style inheritance
- GIVEN a CellXf with XfId pointing to CellStyleXfs
- THEN base formatting comes from CellStyleXfs
- AND Apply* flags indicate which properties are overridden

### Requirement: Cell Style Definitions

The system SHALL support named cell styles.

#### Scenario: CellStyle element
- GIVEN a CellStyle element
- WHEN properties are accessed
- THEN Name, XfId (into CellStyleXfs), BuiltinId, CustomBuiltin, Hidden are available

#### Scenario: Built-in styles
- GIVEN BuiltinId values
- THEN standard styles are available:
  - 0: Normal
  - 1: RowLevel_1 through RowLevel_7
  - 2: ColLevel_1 through ColLevel_7
  - 3: Comma
  - 4: Currency
  - 5: Percent

### Requirement: Alignment

The system SHALL support cell alignment properties.

#### Scenario: Alignment properties
- GIVEN an Alignment element (in CellXf)
- WHEN properties are accessed
- THEN Horizontal, Vertical, TextRotation, WrapText, Indent, ShrinkToFit, ReadingOrder, JustifyLastLine are available

#### Scenario: Horizontal alignment values
- GIVEN Horizontal attribute
- THEN values General, Left, Center, Right, Fill, Justify, CenterContinuous, Distributed are supported

#### Scenario: Vertical alignment values
- GIVEN Vertical attribute
- THEN values Top, Center, Bottom, Justify, Distributed are supported

### Requirement: Protection

The system SHALL support cell protection properties.

#### Scenario: Protection properties
- GIVEN a Protection element (in CellXf)
- WHEN properties are accessed
- THEN Locked, Hidden are available

#### Scenario: Cell protection behavior
- GIVEN Locked=true (default)
- WHEN sheet protection is enabled
- THEN the cell cannot be edited

### Requirement: Differential Formatting

The system SHALL support differential formats for conditional formatting.

#### Scenario: DifferentialFormat (Dxf)
- GIVEN a Dxf element in DifferentialFormats
- WHEN properties are accessed
- THEN Font, Fill, Border, NumberFormat, Alignment are available
- AND only changed properties are specified (differential)

#### Scenario: Use in conditional formatting
- GIVEN a conditional formatting rule
- WHEN DxfId references a Dxf
- THEN that differential format is applied when the condition is met

### Requirement: Table Styles

The system SHALL support table style definitions.

#### Scenario: TableStyles element
- GIVEN a TableStyles element
- WHEN properties are accessed
- THEN DefaultTableStyle, DefaultPivotStyle, TableStyle children are available

#### Scenario: TableStyle element
- GIVEN a TableStyle element
- WHEN properties are accessed
- THEN Name, Pivot, Count, TableStyleElement children are available

#### Scenario: TableStyleElement
- GIVEN a TableStyleElement element
- WHEN properties are accessed
- THEN Type, Size, DxfId are available
- AND Type values include WholeTable, HeaderRow, TotalRow, FirstColumn, LastColumn, etc.

### Requirement: Indexed Colors

The system SHALL support indexed color palette.

#### Scenario: Access indexed colors
- GIVEN an IndexedColors element
- WHEN RgbColor children are accessed
- THEN the 64-color palette is available

#### Scenario: Default indexed colors
- GIVEN default indexed colors
- THEN standard colors are at fixed indices:
  - 0: Black
  - 1: White
  - 2: Red
  - etc.

### Requirement: Theme Colors

The system SHALL support theme color references.

#### Scenario: Theme color in style
- GIVEN a Color element with Theme attribute
- THEN the color references the theme color scheme
- AND values 0-9 correspond to:
  - 0: Background 1 (lt1)
  - 1: Text 1 (dk1)
  - 2: Background 2 (lt2)
  - 3: Text 2 (dk2)
  - 4-9: Accent 1-6

#### Scenario: Theme color with tint
- GIVEN a Color with Theme and Tint
- THEN the theme color is modified by the tint (-1.0 to 1.0)

### Requirement: Style Builder API

The system SHALL provide a fluent API for building styles.

#### Scenario: Build complete style
- GIVEN a StyleBuilder
- WHEN font, fill, border, number format, alignment are set
- THEN a complete CellXf is created with all components

#### Scenario: Style deduplication
- GIVEN identical styles built twice
- WHEN added to the stylesheet
- THEN only one CellXf is created
- AND the same index is returned
