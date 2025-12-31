# PDF Drawing Spec Delta

## ADDED Requirements

### Requirement: SmartArt Diagram Rendering
The system SHALL render SmartArt diagrams in PDF output with full shape positioning and styling.

#### Scenario: Render laid-out SmartArt diagram
- GIVEN a SmartArt diagram with layout applied
- WHEN diagram is rendered to PDF
- THEN all shapes are drawn at their layout positions
- AND all connectors are drawn between shapes
- AND shape text is rendered

#### Scenario: Render diagram without layout
- GIVEN a SmartArt diagram without layout applied
- WHEN diagram is rendered to PDF
- THEN layout is automatically applied before rendering
- AND shapes are positioned and drawn

#### Scenario: Render shape backgrounds
- GIVEN a SmartArt shape with fill color blue
- WHEN shape is rendered to PDF
- THEN shape background is filled with blue color

#### Scenario: Render shape borders
- GIVEN a SmartArt shape with border color black, width 2pt
- WHEN shape is rendered to PDF
- THEN shape is outlined with black 2pt border

#### Scenario: Render shape text
- GIVEN a SmartArt shape with text "Manager"
- WHEN shape is rendered to PDF
- THEN text "Manager" is drawn centered in shape
- AND text uses shape's font style and size

#### Scenario: Render straight-line connectors
- GIVEN a list diagram with straight connectors between items
- WHEN diagram is rendered to PDF
- THEN connectors are drawn as straight lines
- AND line color and width match connector style

#### Scenario: Render curved connectors
- GIVEN a cycle diagram with curved connectors
- WHEN diagram is rendered to PDF
- THEN connectors are drawn as bezier curves
- AND curves follow circular path

#### Scenario: Render connector arrows
- GIVEN connectors with arrow heads
- WHEN diagram is rendered to PDF
- THEN arrow heads are drawn at destination end
- AND arrows point towards destination shape

#### Scenario: Render dotted connectors
- GIVEN a hierarchy with assistant node using dotted connector
- WHEN diagram is rendered to PDF
- THEN assistant connector is drawn with dotted line pattern
- AND regular connectors use solid lines

#### Scenario: Render multiple diagrams on page
- GIVEN a PDF page with 2 SmartArt diagrams
- WHEN page is rendered
- THEN both diagrams are rendered at their respective positions
- AND diagrams do not overlap

### Requirement: SmartArt Shape Rendering
The system SHALL render individual SmartArt shapes with geometry, fill, border, and text.

#### Scenario: Render rectangle shape
- GIVEN a SmartArt shape with rectangle geometry
- WHEN shape is rendered to PDF
- THEN a rectangle is drawn at shape position
- AND rectangle has shape dimensions

#### Scenario: Render ellipse shape
- GIVEN a SmartArt shape with ellipse geometry
- WHEN shape is rendered to PDF
- THEN an ellipse is drawn at shape position
- AND ellipse is inscribed in shape bounds

#### Scenario: Render trapezoid shape
- GIVEN a SmartArt shape with trapezoid geometry (pyramid level)
- WHEN shape is rendered to PDF
- THEN a trapezoid is drawn with correct angles

#### Scenario: Render shape with gradient fill
- GIVEN a SmartArt shape with gradient fill
- WHEN shape is rendered to PDF
- THEN shape background uses gradient from start to end color

#### Scenario: Render semi-transparent shape
- GIVEN a SmartArt shape with 50% opacity
- WHEN shape is rendered to PDF
- THEN shape is rendered with 50% transparency

### Requirement: SmartArt Connector Rendering
The system SHALL render connectors between SmartArt shapes with various styles.

#### Scenario: Calculate connector endpoints
- GIVEN source shape at (100, 50) with size 80x40
- AND destination shape at (100, 120) with size 80x40
- WHEN connector endpoints are calculated
- THEN start point is at (140, 90) (bottom center of source)
- AND end point is at (140, 120) (top center of destination)

#### Scenario: Render bezier curve connector
- GIVEN a curved connector with control point
- WHEN connector is rendered to PDF
- THEN a bezier curve is drawn through control point
- AND curve is smooth

#### Scenario: Render connector with line width
- GIVEN a connector with lineWidth = 3 points
- WHEN connector is rendered
- THEN line is drawn with 3pt thickness

#### Scenario: Render connector with custom color
- GIVEN a connector with color red
- WHEN connector is rendered
- THEN line is drawn in red color

#### Scenario: Render arrow head
- GIVEN a connector with arrow size 10 points
- WHEN connector is rendered
- THEN arrow head triangle is drawn at destination
- AND arrow head has base width 10 points

### Requirement: SmartArt Text Rendering
The system SHALL render text within SmartArt shapes with proper alignment and formatting.

#### Scenario: Center-align text in shape
- GIVEN a shape with centered text "CEO"
- WHEN text is rendered
- THEN text is centered horizontally in shape
- AND text is centered vertically in shape

#### Scenario: Wrap long text
- GIVEN a shape with text longer than shape width
- WHEN text is rendered
- THEN text wraps to multiple lines
- AND wrapped lines fit within shape bounds

#### Scenario: Render text with font style
- GIVEN text with font "Arial Bold" size 14pt
- WHEN text is rendered
- THEN text uses Arial Bold font
- AND text size is 14 points

#### Scenario: Render text with color
- GIVEN text with color white
- WHEN text is rendered on blue background
- THEN text is drawn in white color for contrast

### Requirement: SmartArt Rendering Performance
The system SHALL render SmartArt diagrams efficiently for PDF generation.

#### Scenario: Render 100-shape diagram
- GIVEN a SmartArt diagram with 100 shapes
- WHEN diagram is rendered to PDF
- THEN rendering completes in under 500 milliseconds

#### Scenario: Render diagram with many connectors
- GIVEN a diagram with 50 curved connectors
- WHEN diagram is rendered
- THEN all connectors are rendered
- AND rendering completes in reasonable time

#### Scenario: Reuse rendering resources
- GIVEN multiple diagrams with same colors and fonts
- WHEN diagrams are rendered
- THEN color and font resources are shared
- AND PDF file size is minimized

### Requirement: SmartArt Rendering Fidelity
The system SHALL render SmartArt diagrams with high visual fidelity matching layout output.

#### Scenario: Position accuracy
- GIVEN a shape positioned at (150.5, 200.3)
- WHEN rendered to PDF
- THEN shape is drawn at exactly (150.5, 200.3)
- AND sub-pixel positioning is preserved

#### Scenario: Size accuracy
- GIVEN a shape with size 85.7 x 42.3 points
- WHEN rendered to PDF
- THEN shape is drawn at exactly 85.7 x 42.3 points

#### Scenario: Connector path accuracy
- GIVEN a bezier connector with specific control points
- WHEN rendered to PDF
- THEN curve matches expected mathematical path
- AND curve smoothness is preserved

#### Scenario: Text positioning accuracy
- GIVEN text centered in shape at (100, 50)
- WHEN rendered to PDF
- THEN text baseline is at correct position
- AND text is pixel-perfect centered
