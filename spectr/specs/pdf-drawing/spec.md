# Pdf Drawing Specification

## Requirements

### Requirement: Shape Geometry Rendering
The system SHALL render DrawingML shape geometries to PDF paths.

#### Scenario: Preset shape
- GIVEN a preset shape type (e.g., rect, ellipse, star5)
- WHEN rendered
- THEN correct geometric path is produced

#### Scenario: Custom geometry
- GIVEN a shape with custom path definition
- WHEN rendered
- THEN path commands are converted to PDF path operations

#### Scenario: Rounded corners
- GIVEN a shape with rounded corners
- WHEN rendered
- THEN corner radii are rendered correctly

### Requirement: Fill Rendering
The system SHALL render DrawingML fill types.

#### Scenario: Solid fill
- GIVEN a shape with solid color fill
- WHEN rendered
- THEN shape is filled with specified color

#### Scenario: Gradient fill
- GIVEN a shape with linear gradient fill
- WHEN rendered
- THEN gradient is rendered with correct colors and direction

#### Scenario: Radial gradient
- GIVEN a shape with radial gradient fill
- WHEN rendered
- THEN radial gradient is rendered from center outward

#### Scenario: Pattern fill
- GIVEN a shape with pattern fill
- WHEN rendered
- THEN pattern is tiled within shape bounds

#### Scenario: Picture fill
- GIVEN a shape with image fill
- WHEN rendered
- THEN image is placed within shape (stretched, tiled, or cropped)

### Requirement: Line/Stroke Rendering
The system SHALL render DrawingML outline properties.

#### Scenario: Solid line
- GIVEN a shape with solid outline
- WHEN rendered
- THEN outline is drawn with correct color and width

#### Scenario: Dashed line
- GIVEN a shape with dashed outline
- WHEN rendered
- THEN dash pattern matches specification

#### Scenario: Line caps
- GIVEN a line with specific cap style (flat, round, square)
- WHEN rendered
- THEN line ends are drawn with correct cap

#### Scenario: Line joins
- GIVEN a shape with specific join style (miter, round, bevel)
- WHEN rendered
- THEN corners are drawn with correct join

#### Scenario: Compound line
- GIVEN a line with compound style (double, triple)
- WHEN rendered
- THEN multiple strokes are drawn

### Requirement: Text in Shapes
The system SHALL render text content within shapes.

#### Scenario: Text body
- GIVEN a shape with text content
- WHEN rendered
- THEN text is laid out within shape bounds

#### Scenario: Text wrapping
- GIVEN a shape with text that exceeds width
- WHEN rendered
- THEN text wraps to multiple lines within shape

#### Scenario: Vertical text
- GIVEN a shape with vertical text direction
- WHEN rendered
- THEN text is rotated 90 degrees

#### Scenario: Text anchor
- GIVEN a shape with text anchor (top, middle, bottom)
- WHEN rendered
- THEN text is vertically positioned accordingly

#### Scenario: Text margins
- GIVEN a shape with internal text margins
- WHEN rendered
- THEN text is inset from shape edges

### Requirement: Picture Rendering
The system SHALL render DrawingML pictures.

#### Scenario: Inline image
- GIVEN a picture element with image reference
- WHEN rendered
- THEN image is displayed at specified size

#### Scenario: Image crop
- GIVEN a picture with crop settings
- WHEN rendered
- THEN only visible portion of image is shown

#### Scenario: Image stretch
- GIVEN a picture with stretch fill settings
- WHEN rendered
- THEN image fills frame with possible distortion

### Requirement: Transform Rendering
The system SHALL apply DrawingML transforms.

#### Scenario: Rotation
- GIVEN a shape with rotation angle
- WHEN rendered
- THEN shape is rotated around its center

#### Scenario: Flip
- GIVEN a shape with horizontal/vertical flip
- WHEN rendered
- THEN shape is mirrored appropriately

#### Scenario: Offset
- GIVEN a shape with position offset
- WHEN rendered
- THEN shape is placed at correct coordinates

### Requirement: Effect Rendering
The system SHALL render DrawingML effects.

#### Scenario: Drop shadow
- GIVEN a shape with drop shadow effect
- WHEN rendered
- THEN shadow is rendered behind shape with correct blur and offset

#### Scenario: Outer glow
- GIVEN a shape with outer glow effect
- WHEN rendered
- THEN glow is rendered around shape edges

#### Scenario: Soft edges
- GIVEN a shape with soft edges effect
- WHEN rendered
- THEN edges are blurred/feathered

#### Scenario: Reflection
- GIVEN a shape with reflection effect
- WHEN rendered
- THEN mirrored reflection appears below shape

### Requirement: Connector Rendering
The system SHALL render connector lines between shapes.

#### Scenario: Straight connector
- GIVEN a straight connector between shapes
- WHEN rendered
- THEN line connects at specified connection points

#### Scenario: Elbow connector
- GIVEN an elbow (bent) connector
- WHEN rendered
- THEN connector bends at right angles between shapes

#### Scenario: Curved connector
- GIVEN a curved connector
- WHEN rendered
- THEN smooth curve connects the shapes

### Requirement: Group Rendering
The system SHALL render grouped shapes.

#### Scenario: Shape group
- GIVEN a group containing multiple shapes
- WHEN rendered
- THEN all shapes are rendered with correct relative positions

#### Scenario: Nested groups
- GIVEN groups containing other groups
- WHEN rendered
- THEN nesting is preserved with correct transforms

### Requirement: Chart Rendering
The system SHALL render DrawingML charts.

#### Scenario: Column chart
- GIVEN a column/bar chart
- WHEN rendered
- THEN bars are drawn with correct heights and spacing

#### Scenario: Line chart
- GIVEN a line chart with data series
- WHEN rendered
- THEN lines connect data points correctly

#### Scenario: Pie chart
- GIVEN a pie chart
- WHEN rendered
- THEN segments are sized proportionally to data

#### Scenario: Chart legend
- GIVEN a chart with legend
- WHEN rendered
- THEN legend is positioned and styled correctly

#### Scenario: Chart axes
- GIVEN a chart with axes
- WHEN rendered
- THEN axes, labels, and gridlines are drawn correctly

