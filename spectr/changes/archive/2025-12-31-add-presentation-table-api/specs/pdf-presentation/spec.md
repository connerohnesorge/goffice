# PDF Presentation Spec Delta

## ADDED Requirements

### Requirement: PowerPoint Table Rendering to PDF
The system SHALL render PowerPoint tables to PDF with accurate positioning and formatting.

#### Scenario: Render simple table to PDF
- GIVEN a presentation with a 3x3 table
- WHEN presentation is rendered to PDF
- THEN table is drawn at correct position on PDF page
- AND table grid (borders) is rendered
- AND cell text is rendered

#### Scenario: Render table with custom position
- GIVEN a table positioned at (100pt, 150pt) on slide
- WHEN rendered to PDF
- THEN table appears at (100pt, 150pt) on PDF page

#### Scenario: Render table with custom size
- GIVEN a table sized to 400pt x 200pt
- WHEN rendered to PDF
- THEN table dimensions in PDF are 400pt x 200pt

#### Scenario: Render multiple tables on single slide
- GIVEN a slide with 2 tables at different positions
- WHEN rendered to PDF
- THEN both tables are rendered
- AND tables are at correct positions
- AND tables do not overlap unless positioned to overlap

### Requirement: Table Grid Rendering
The system SHALL render table grid lines (borders) in PDF.

#### Scenario: Render horizontal grid lines
- GIVEN a table with 4 rows
- WHEN table grid is rendered
- THEN 5 horizontal lines are drawn (top, 3 internal, bottom)
- AND lines span full table width

#### Scenario: Render vertical grid lines
- GIVEN a table with 3 columns
- WHEN table grid is rendered
- THEN 4 vertical lines are drawn (left, 2 internal, right)
- AND lines span full table height

#### Scenario: Grid line color and width
- WHEN table grid is rendered
- THEN grid lines are black (0, 0, 0)
- AND line width is 0.5 points

#### Scenario: Grid lines align with column widths
- GIVEN a table with unequal column widths (100pt, 150pt, 200pt)
- WHEN vertical grid lines are rendered
- THEN lines are positioned at correct x-coordinates based on column widths

### Requirement: Cell Background Rendering
The system SHALL render cell background fills in PDF.

#### Scenario: Render cell with solid fill
- GIVEN a cell with background color RGB(255, 200, 100)
- WHEN cell is rendered
- THEN filled rectangle is drawn with color (1.0, 0.78, 0.39)
- AND rectangle covers cell area

#### Scenario: Render cell without background
- GIVEN a cell with no background fill
- WHEN cell is rendered
- THEN no fill rectangle is drawn
- AND cell background is transparent (page background shows through)

#### Scenario: Render multiple cells with different backgrounds
- GIVEN row 0 cells with backgrounds (red, green, blue)
- WHEN cells are rendered
- THEN each cell has its specified background color

### Requirement: Cell Text Rendering
The system SHALL render cell text content in PDF.

#### Scenario: Render cell text top-aligned
- GIVEN a cell with text "Header" and top alignment
- WHEN cell is rendered
- THEN text is drawn near top of cell with padding

#### Scenario: Render cell text center-aligned
- GIVEN a cell with text "Data" and center alignment
- WHEN cell is rendered
- THEN text is drawn vertically centered in cell

#### Scenario: Render cell text bottom-aligned
- GIVEN a cell with text "Total" and bottom alignment
- WHEN cell is rendered
- THEN text is drawn near bottom of cell with padding

#### Scenario: Render empty cell
- GIVEN a cell with no text content
- WHEN cell is rendered
- THEN no text is drawn
- AND no error occurs

#### Scenario: Render cell with long text
- GIVEN a cell with text longer than cell width
- WHEN cell is rendered
- THEN text is clipped or wrapped within cell bounds

### Requirement: Table Style Rendering
The system SHALL render table styles (banded rows, special rows) in PDF.

#### Scenario: Render table with banded rows
- GIVEN a table with BandRow=true
- WHEN table is rendered
- THEN alternating rows have different background colors
- AND pattern matches Office banded row appearance

#### Scenario: Render table with first row bold
- GIVEN a table with FirstRow=true
- WHEN table is rendered
- THEN first row text is rendered in bold font
- AND first row may have different background

#### Scenario: Render table with last row bold
- GIVEN a table with LastRow=true
- WHEN table is rendered
- THEN last row text is rendered in bold font
- AND last row may have different background

### Requirement: Cell Margin Rendering
The system SHALL respect cell margins when rendering text.

#### Scenario: Render text with cell margins
- GIVEN a cell with margins (left:10pt, right:10pt, top:5pt, bottom:5pt)
- WHEN cell text is rendered
- THEN text is inset 10pt from left/right edges
- AND text is inset 5pt from top/bottom edges

#### Scenario: Render text with zero margins
- GIVEN a cell with zero margins
- WHEN cell text is rendered
- THEN text is drawn with minimal default padding

### Requirement: Table Column Width Rendering
The system SHALL use table column widths from TableGrid when rendering.

#### Scenario: Render table with equal columns
- GIVEN a table with 4 equal columns (100pt each)
- WHEN table is rendered
- THEN each column occupies exactly 100pt width

#### Scenario: Render table with unequal columns
- GIVEN a table with columns (50pt, 150pt, 200pt, 100pt)
- WHEN table is rendered
- THEN columns are rendered at specified widths
- AND total width matches table width

#### Scenario: Render table with auto-calculated columns
- GIVEN a table with SetEqualColumnWidths() applied
- WHEN table is rendered
- THEN columns are evenly distributed across table width

### Requirement: Table Row Height Rendering
The system SHALL calculate row heights for rendering.

#### Scenario: Render table with uniform row heights
- GIVEN a table where all rows have equal height
- WHEN table is rendered
- THEN rows are evenly distributed across table height

#### Scenario: Render table with custom row heights
- GIVEN rows with heights (30pt, 50pt, 40pt)
- WHEN table is rendered
- THEN each row has its specified height

#### Scenario: Render table with auto-calculated row heights
- GIVEN a table where row heights are unspecified
- WHEN table is rendered
- THEN row heights are calculated as (table height / row count)

### Requirement: Table Rendering Performance
The system SHALL render tables efficiently for PDF generation.

#### Scenario: Render small table quickly
- GIVEN a 3x3 table
- WHEN table is rendered to PDF
- THEN rendering completes in under 10 milliseconds

#### Scenario: Render large table
- GIVEN a 50x20 table (1000 cells)
- WHEN table is rendered to PDF
- THEN all cells are rendered
- AND rendering completes in reasonable time (< 500ms)

### Requirement: Table Rendering Accuracy
The system SHALL render tables with visual fidelity matching PowerPoint display.

#### Scenario: Position accuracy
- GIVEN a table at precise position (123.45pt, 67.89pt)
- WHEN rendered to PDF
- THEN table is positioned at exactly (123.45, 67.89)

#### Scenario: Size accuracy
- GIVEN a table with size (432.1pt x 234.5pt)
- WHEN rendered to PDF
- THEN table dimensions are exactly (432.1, 234.5)

#### Scenario: Grid alignment accuracy
- GIVEN a table with column widths specified
- WHEN rendered to PDF
- THEN grid lines align perfectly with column boundaries
- AND no gaps or overlaps occur

#### Scenario: Text positioning accuracy
- GIVEN cell text with specific alignment
- WHEN rendered to PDF
- THEN text position matches expected alignment
- AND text is pixel-perfect centered/aligned

### Requirement: Table Font Rendering
The system SHALL render table cell text with appropriate fonts.

#### Scenario: Render table with default font
- GIVEN a table with no explicit font specified
- WHEN table is rendered
- THEN Arial font at 11pt is used for cell text

#### Scenario: Render table with custom font
- GIVEN a table where cells specify font "Calibri" 14pt
- WHEN table is rendered
- THEN Calibri font at 14pt is used

#### Scenario: Render first row with bold font
- GIVEN a table with FirstRow=true
- WHEN first row is rendered
- THEN first row text uses bold weight

### Requirement: Table Border Rendering
The system SHALL render cell borders if specified in cell properties.

#### Scenario: Render cell with custom borders
- GIVEN a cell with left border (2pt, red) and right border (1pt, blue)
- WHEN cell is rendered
- THEN left border is drawn 2pt wide in red
- AND right border is drawn 1pt wide in blue

#### Scenario: Render cell without borders
- GIVEN a cell with no border specifications
- WHEN cell is rendered
- THEN only table grid lines are drawn
- AND no additional cell borders appear

### Requirement: Table PDF Integration
The system SHALL integrate table rendering with existing PDF presentation rendering.

#### Scenario: Render slide with table and other shapes
- GIVEN a slide with table, image, and text box
- WHEN slide is rendered to PDF
- THEN all elements are rendered
- AND z-order is respected
- AND table appears in correct layer

#### Scenario: Render multi-slide presentation with tables
- GIVEN a 5-slide presentation where slides 2 and 4 have tables
- WHEN presentation is rendered to PDF
- THEN slides 2 and 4 have tables rendered
- AND slides without tables render normally
