# PDF Drawing Spec Delta

## ADDED Requirements

### Requirement: Connector Rendering in PDF

The system SHALL render ConnectionShape elements to PDF with routed paths and decorations.

#### Scenario: Render straight line connector
- GIVEN a ConnectionShape with straight line path from (100, 100) to (500, 500)
- WHEN the connector is rendered to PDF
- THEN a PDF line is drawn from (100, 100) to (500, 500)
- AND the line uses the connector's line style (color, width, dash pattern)
- AND the line appears in the PDF output

#### Scenario: Render elbow connector with orthogonal segments
- GIVEN a ConnectionShape with elbow path having 5 segments
- AND path is [MoveTo(100,100), LineTo(100,300), LineTo(300,300), LineTo(300,500), LineTo(500,500)]
- WHEN the connector is rendered to PDF
- THEN 4 line segments are drawn in PDF
- AND segments form right-angle bends
- AND all segments connect correctly

#### Scenario: Render curved connector with Bezier curve
- GIVEN a ConnectionShape with cubic Bezier path
- AND path has MoveTo(100,100) and CubicBezierTo(200,100, 300,200, 400,200)
- WHEN the connector is rendered to PDF
- THEN a PDF Bezier curve is drawn
- AND curve uses control points (200,100) and (300,200)
- AND curve ends at (400,200)
- AND curve is smooth in PDF output

#### Scenario: Connector line style application
- GIVEN a ConnectionShape with line color RGB(255,0,0) and width 2 points
- WHEN the connector is rendered to PDF
- THEN PDF line color is set to RGB(255,0,0)
- AND PDF line width is set to 2 points
- AND the line appears red with correct width in PDF viewer

#### Scenario: Connector with dashed line style
- GIVEN a ConnectionShape with dash pattern [5, 3] (5 on, 3 off)
- WHEN the connector is rendered to PDF
- THEN PDF line dash pattern is set to [5, 3]
- AND the rendered line appears dashed in PDF viewer

#### Scenario: Connector with solid line style
- GIVEN a ConnectionShape with solid line style (no dash pattern)
- WHEN the connector is rendered to PDF
- THEN PDF line is drawn with solid stroke
- AND no dash pattern is applied

### Requirement: Arrow Head Rendering

The system SHALL render arrow decorations at connection points in PDF output.

#### Scenario: Render start arrow head
- GIVEN a ConnectionShape with start arrow decoration
- AND arrow type is "arrow" (standard triangular arrow)
- AND path starts at (100, 100) with direction toward (200, 200)
- WHEN the connector is rendered to PDF
- THEN an arrow head is drawn at (100, 100)
- AND arrow points in direction of line (toward (200, 200))
- AND arrow size is proportional to line width

#### Scenario: Render end arrow head
- GIVEN a ConnectionShape with end arrow decoration
- AND arrow type is "arrow"
- AND path ends at (500, 500) with direction from (400, 400)
- WHEN the connector is rendered to PDF
- THEN an arrow head is drawn at (500, 500)
- AND arrow points in direction of line (from (400, 400))
- AND arrow tip is at (500, 500)

#### Scenario: Render both start and end arrows
- GIVEN a ConnectionShape with both start and end arrow decorations
- WHEN the connector is rendered to PDF
- THEN arrow head is drawn at start point
- AND arrow head is drawn at end point
- AND both arrows point in correct directions
- AND line connects both arrow bases

#### Scenario: Arrow head direction calculation for straight line
- GIVEN a straight line connector from (100, 100) to (500, 500)
- AND end arrow decoration
- WHEN arrow direction is calculated
- THEN arrow angle equals atan2(500-100, 500-100) = 45 degrees
- AND arrow points along 45-degree line

#### Scenario: Arrow head direction for elbow connector
- GIVEN an elbow connector with final segment from (400, 500) to (500, 500) (horizontal)
- AND end arrow decoration
- WHEN arrow direction is calculated
- THEN arrow angle equals 0 degrees (pointing right)
- AND arrow is oriented along final segment direction

#### Scenario: Arrow head direction for Bezier curve
- GIVEN a Bezier curve connector ending at (500, 500)
- AND control point 2 at (450, 480)
- AND end arrow decoration
- WHEN arrow direction is calculated
- THEN arrow angle is calculated from tangent at t=1 (curve endpoint)
- AND tangent direction is from control point 2 toward endpoint
- AND arrow points along tangent direction

#### Scenario: Arrow head size proportional to line width
- GIVEN a ConnectionShape with line width 4 points
- AND arrow decoration with size "medium"
- WHEN arrow head is rendered
- THEN arrow head length is approximately 4 * arrow_size_factor
- AND arrow head width is approximately 3 * arrow_size_factor
- AND arrow scales with line width

#### Scenario: Arrow types
- GIVEN a ConnectionShape with arrow type "arrow" (standard triangular)
- WHEN the arrow is rendered
- THEN a filled triangle is drawn
- AND triangle base is perpendicular to line direction
- AND triangle tip is at connection point

#### Scenario: Diamond arrow type
- GIVEN a ConnectionShape with arrow type "diamond"
- WHEN the arrow is rendered
- THEN a diamond shape is drawn at connection point
- AND diamond is oriented along line direction

#### Scenario: No arrow decoration
- GIVEN a ConnectionShape with no start or end arrow
- WHEN the connector is rendered to PDF
- THEN no arrow heads are drawn
- AND line extends fully to connection points

### Requirement: Shape Reference Resolution

The system SHALL resolve shape references to calculate connection point positions for PDF rendering.

#### Scenario: Resolve start shape reference
- GIVEN a ConnectionShape with start connection to shape ID "shape1"
- AND shape ID "shape1" exists on the slide
- WHEN connector is rendered to PDF
- THEN the shape with ID "shape1" is located
- AND the shape's position and size are retrieved
- AND connection point position is calculated

#### Scenario: Connection point position calculation
- GIVEN a ConnectionShape connected to shape "shape1" at connection site index 0
- AND shape "shape1" has offset (100, 200) and extent (400, 300)
- AND connection site 0 is top center with position (200, 0) in shape coordinates
- WHEN connection point absolute position is calculated
- THEN absolute X equals shape offset X + site X = 100 + 200 = 300
- AND absolute Y equals shape offset Y + site Y = 200 + 0 = 200
- AND position is used as path start point

#### Scenario: Resolve end shape reference
- GIVEN a ConnectionShape with end connection to shape ID "shape2"
- AND shape ID "shape2" exists on the slide
- WHEN connector is rendered to PDF
- THEN the shape with ID "shape2" is located
- AND connection site on shape2 is accessed
- AND absolute connection point position is calculated

#### Scenario: Shape reference resolution with missing shape
- GIVEN a ConnectionShape with start connection to shape ID "missingShape"
- AND shape ID "missingShape" does not exist on the slide
- WHEN connector is rendered to PDF
- THEN shape resolution fails gracefully
- AND connector is skipped or rendered with default position
- AND no error or panic occurs
- AND other connectors render normally

#### Scenario: Connection site index out of bounds
- GIVEN a ConnectionShape connected to shape at connection site index 10
- AND the shape has only 4 connection sites
- WHEN connection point position is calculated
- THEN default position (e.g., shape center) is used
- OR error is handled gracefully
- AND rendering continues without panic

### Requirement: Path Command Rendering

The system SHALL render Path commands to PDF drawing operations.

#### Scenario: Render MoveTo command
- GIVEN a Path with MoveTo(100, 200) command
- WHEN the command is rendered to PDF
- THEN PDF current point is moved to (100, 200)
- AND no stroke is drawn
- AND subsequent LineTo commands start from (100, 200)

#### Scenario: Render LineTo command
- GIVEN a Path with MoveTo(100, 100) then LineTo(200, 200)
- WHEN the commands are rendered to PDF
- THEN PDF moves to (100, 100)
- AND PDF draws line to (200, 200)
- AND line is stroked with current line style

#### Scenario: Render CubicBezierTo command
- GIVEN a Path with MoveTo(100,100) then CubicBezierTo(150,100, 150,200, 200,200)
- WHEN the commands are rendered to PDF
- THEN PDF moves to (100, 100)
- AND PDF draws cubic Bezier curve with control points (150,100) and (150,200) to endpoint (200,200)
- AND curve is stroked with current line style

#### Scenario: Render path with multiple segments
- GIVEN a Path with [MoveTo(0,0), LineTo(100,0), LineTo(100,100), LineTo(0,100)]
- WHEN the path is rendered to PDF
- THEN PDF draws 3 line segments forming a U-shape
- AND all segments are connected
- AND path is stroked as continuous line

#### Scenario: Coordinate transformation for PDF
- GIVEN a Path with coordinates in EMU
- AND Path has MoveTo(914400, 914400) (1 inch, 1 inch in EMU)
- WHEN the path is rendered to PDF
- THEN coordinates are converted from EMU to PDF points (72 points per inch)
- AND PDF MoveTo is called with (72, 72) in PDF coordinate space
- AND coordinate system transformation is applied correctly

#### Scenario: PDF coordinate system origin
- GIVEN PDF coordinate system with origin at bottom-left
- AND PowerPoint coordinate system with origin at top-left
- WHEN connector path is rendered
- THEN Y coordinates are transformed to flip vertical axis
- AND Y_pdf = page_height - Y_powerpoint
- AND connector appears in correct position in PDF

### Requirement: Connector Renderer Integration

The system SHALL integrate connector rendering into presentation PDF rendering pipeline.

#### Scenario: Render connectors during slide rendering
- GIVEN a Slide with 5 shapes and 3 connectors
- WHEN the slide is rendered to PDF
- THEN shapes are rendered first (in back layer)
- AND connectors are rendered after shapes (in front layer)
- AND connectors appear on top of shapes in PDF output

#### Scenario: Connector Z-order
- GIVEN a Slide with connector1 added before connector2
- WHEN the slide is rendered to PDF
- THEN connector1 is rendered first
- AND connector2 is rendered second (on top of connector1)
- AND Z-order matches ShapeTree element order

#### Scenario: Skip disconnected connectors
- GIVEN a Slide with connector that has no start or end connection
- WHEN the slide is rendered to PDF
- THEN the disconnected connector is skipped
- AND no error or panic occurs
- AND connected connectors render normally

#### Scenario: Render connector with invalid shape reference
- GIVEN a connector referencing deleted shape
- WHEN the slide is rendered to PDF
- THEN the connector is skipped or rendered with fallback position
- AND rendering continues without error
- AND other valid connectors render normally

### Requirement: Line Style Support

The system SHALL support various line styles for connector rendering in PDF.

#### Scenario: Solid line
- GIVEN a ConnectionShape with solid line style
- WHEN rendered to PDF
- THEN PDF stroke is solid (no dash pattern)

#### Scenario: Dashed line
- GIVEN a ConnectionShape with dash style "dash"
- WHEN rendered to PDF
- THEN PDF dash pattern is set to [4, 2] or similar
- AND line appears dashed in PDF viewer

#### Scenario: Dotted line
- GIVEN a ConnectionShape with dash style "dot"
- WHEN rendered to PDF
- THEN PDF dash pattern is set to [1, 1] or similar
- AND line appears dotted in PDF viewer

#### Scenario: Dash-dot line
- GIVEN a ConnectionShape with dash style "dashDot"
- WHEN rendered to PDF
- THEN PDF dash pattern is set to [4, 2, 1, 2] or similar
- AND line appears as dash-dot pattern in PDF viewer

#### Scenario: Line cap style
- GIVEN a ConnectionShape with line cap style "round"
- WHEN rendered to PDF
- THEN PDF line cap is set to round
- AND line ends appear rounded in PDF viewer

#### Scenario: Line join style
- GIVEN a ConnectionShape with elbow path and line join style "miter"
- WHEN rendered to PDF
- THEN PDF line join is set to miter
- AND corners appear sharp at right-angle bends

### Requirement: Color and Transparency

The system SHALL render connector colors and transparency in PDF.

#### Scenario: RGB color
- GIVEN a ConnectionShape with line color RGB(255, 0, 0) (red)
- WHEN rendered to PDF
- THEN PDF stroke color is set to RGB(1.0, 0.0, 0.0) (normalized to 0-1 range)
- AND line appears red in PDF viewer

#### Scenario: Transparency
- GIVEN a ConnectionShape with line color RGBA(255, 0, 0, 128) (50% transparent red)
- WHEN rendered to PDF
- THEN PDF stroke color is RGB(1.0, 0.0, 0.0)
- AND PDF graphics state alpha is set to 0.5
- AND line appears semi-transparent in PDF viewer

#### Scenario: Scheme color
- GIVEN a ConnectionShape with line color referencing scheme color "accent1"
- AND presentation theme defines accent1 as RGB(0, 128, 255)
- WHEN rendered to PDF
- THEN scheme color is resolved to RGB(0, 128, 255)
- AND PDF stroke color is set accordingly
- AND line appears in theme color

#### Scenario: No fill for connectors
- GIVEN a ConnectionShape (connectors are line elements, not filled shapes)
- WHEN rendered to PDF
- THEN only stroke is applied (no fill)
- AND PDF fill color is not set
- AND connector appears as line, not filled polygon
