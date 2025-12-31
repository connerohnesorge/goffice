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

### Requirement: Page Interface

The system SHALL provide a Page interface for abstract drawing operations.

#### Scenario: Graphics state management
- GIVEN a Page implementation
- WHEN `page.SaveGraphicsState()` is called
- THEN the current graphics state is pushed to a stack
- AND `page.RestoreGraphicsState()` pops and restores it

#### Scenario: Set fill color RGB
- GIVEN a Page implementation
- WHEN `page.SetFillColor(1.0, 0.0, 0.0)` is called
- THEN subsequent fill operations use red color
- AND the PDF content stream contains "1 0 0 rg"

#### Scenario: Set fill color with alpha
- GIVEN a Page implementation
- WHEN `page.SetFillColorRGBA(1.0, 0.0, 0.0, 0.5)` is called
- THEN subsequent fill operations use red with 50% opacity
- AND an ExtGState resource is created for transparency

#### Scenario: Set stroke color
- GIVEN a Page implementation
- WHEN `page.SetStrokeColor(0.0, 0.0, 1.0)` is called
- THEN subsequent stroke operations use blue color
- AND the PDF content stream contains "0 0 1 RG"

#### Scenario: Set line width
- GIVEN a Page implementation
- WHEN `page.SetLineWidth(2.5)` is called
- THEN subsequent strokes use 2.5 point line width
- AND the PDF content stream contains "2.5 w"

#### Scenario: Set line dash pattern
- GIVEN a Page implementation
- WHEN `page.SetLineDashPattern([]float64{3, 2}, 0)` is called
- THEN subsequent strokes use dash pattern (3 on, 2 off)
- AND the PDF content stream contains "[3 2] 0 d"

#### Scenario: Set line cap
- GIVEN a Page implementation
- WHEN `page.SetLineCap(1)` is called (1 = round cap)
- THEN subsequent strokes use round cap style
- AND the PDF content stream contains "1 J"

#### Scenario: Set line join
- GIVEN a Page implementation
- WHEN `page.SetLineJoin(2)` is called (2 = bevel join)
- THEN subsequent strokes use bevel join style
- AND the PDF content stream contains "2 j"

### Requirement: Shape Drawing

The system SHALL provide primitive shape drawing operations.

#### Scenario: Draw rectangle with fill and stroke
- GIVEN a Page implementation
- WHEN `page.DrawRectangle(100, 200, 50, 30, true, true)` is called
- THEN a rectangle is drawn and both filled and stroked
- AND the PDF content stream contains "100 200 50 30 re" followed by "B"

#### Scenario: Draw rectangle with fill only
- GIVEN a Page implementation
- WHEN `page.DrawRectangle(100, 200, 50, 30, true, false)` is called
- THEN a rectangle is drawn and filled but not stroked
- AND the PDF content stream contains "100 200 50 30 re" followed by "f"

#### Scenario: Draw rounded rectangle
- GIVEN a Page implementation
- WHEN `page.DrawRoundedRectangle(100, 200, 50, 30, 5)` is called
- THEN a rectangle with 5-point corner radius is drawn
- AND the path uses Bezier curves for corners

#### Scenario: Draw circle with fill and stroke
- GIVEN a Page implementation
- WHEN `page.DrawCircle(150, 150, 25, true, true)` is called
- THEN a circle centered at (150, 150) with radius 25 is drawn
- AND the circle is both filled and stroked
- AND the path uses Bezier approximation (4 curves)

#### Scenario: Draw ellipse
- GIVEN a Page implementation
- WHEN `page.DrawEllipse(150, 150, 30, 20)` is called
- THEN an ellipse with horizontal radius 30 and vertical radius 20 is drawn

#### Scenario: Draw line
- GIVEN a Page implementation
- WHEN `page.DrawLine(10, 10, 100, 100)` is called
- THEN a line from (10, 10) to (100, 100) is drawn
- AND the PDF content stream contains "10 10 m 100 100 l"

#### Scenario: Draw complex path
- GIVEN a Page implementation and a Path with multiple segments
- WHEN `page.DrawPath(path)` is called
- THEN the complete path is rendered to the content stream

### Requirement: Fill and Stroke Operations

The system SHALL provide fill and stroke painting operations.

#### Scenario: Fill path
- GIVEN a Page with a rectangle path
- WHEN `page.Fill()` is called
- THEN the path is filled with current fill color
- AND the PDF content stream contains "f"

#### Scenario: Stroke path
- GIVEN a Page with a rectangle path
- WHEN `page.Stroke()` is called
- THEN the path is stroked with current stroke color and line width
- AND the PDF content stream contains "S"

#### Scenario: Fill and stroke path
- GIVEN a Page with a rectangle path
- WHEN `page.FillStroke()` is called
- THEN the path is both filled and stroked
- AND the PDF content stream contains "B"

### Requirement: Text Drawing

The system SHALL provide text drawing operations.

#### Scenario: Set font
- GIVEN a Page implementation
- WHEN `page.SetFont("Helvetica", 12.0)` is called
- THEN subsequent text uses the Helvetica font at 12 points
- AND the font is registered in the page resources

#### Scenario: Draw text
- GIVEN a Page with font set
- WHEN `page.DrawText(100, 200, "Hello")` is called
- THEN "Hello" is drawn at position (100, 200)
- AND text is wrapped in BT/ET operators

#### Scenario: Draw rotated text
- GIVEN a Page with font set
- WHEN `page.DrawTextRotated(100, 200, 45.0, "Rotated")` is called
- THEN "Rotated" is drawn at (100, 200) rotated 45 degrees
- AND a transformation matrix is applied

### Requirement: Image Drawing

The system SHALL provide image drawing operations.

#### Scenario: Draw image
- GIVEN a Page implementation and an Image resource
- WHEN `page.DrawImage(img, 100, 200, 150, 100)` is called
- THEN the image is drawn at (100, 200) scaled to 150x100 points
- AND an XObject reference is written to content stream

### Requirement: Transformations

The system SHALL provide coordinate transformation operations.

#### Scenario: Translate
- GIVEN a Page implementation
- WHEN `page.Translate(50, 100)` is called
- THEN subsequent drawing is offset by (50, 100)
- AND a "cm" operator with translation matrix is written

#### Scenario: Scale
- GIVEN a Page implementation
- WHEN `page.Scale(2.0, 0.5)` is called
- THEN subsequent drawing is scaled 2x horizontally and 0.5x vertically

#### Scenario: Rotate
- GIVEN a Page implementation
- WHEN `page.Rotate(90.0)` is called
- THEN subsequent drawing is rotated 90 degrees

### Requirement: Clipping

The system SHALL provide clipping operations.

#### Scenario: Apply clipping path
- GIVEN a Page implementation with a path drawn
- WHEN `page.Clip()` is called
- THEN the current path becomes the clipping path
- AND subsequent drawing is clipped to that path
- AND the "W n" operator is used

#### Scenario: Set clip rectangle
- GIVEN a Page implementation
- WHEN `page.SetClipRect(50, 50, 200, 200)` is called
- THEN subsequent drawing is clipped to that rectangle
- AND the "W" operator is used

#### Scenario: Clear clip
- GIVEN a Page with a clip rectangle set
- WHEN `page.ClearClip()` is called
- THEN clipping is removed (via graphics state restore)

### Requirement: RenderingContext Integration

The system SHALL integrate Page with the existing RenderingContext.

#### Scenario: Context has Page field
- GIVEN a RenderingContext
- WHEN `ctx.Page` is accessed
- THEN it returns the current Page interface (or nil)

#### Scenario: Set page on context
- GIVEN a RenderingContext and a PDFPage
- WHEN `ctx.SetPage(page)` is called
- THEN `ctx.Page` returns the provided page
- AND renderers can use `ctx.Page.DrawRectangle(...)` etc.

### Requirement: Chart Rendering (Enabled)

The system SHALL render charts via the enabled chart_renderer.go.

#### Scenario: Render bar chart
- GIVEN a RenderingContext with Page set and a BarChart with data
- WHEN `RenderBarChart(ctx, chart, bounds)` is called
- THEN bars are drawn with correct heights proportional to data values
- AND bar colors match chart series colors
- AND no error is returned

#### Scenario: Render line chart
- GIVEN a RenderingContext with Page set and a LineChart with data
- WHEN `RenderLineChart(ctx, chart, bounds)` is called
- THEN lines connect data points
- AND markers are drawn at data points if specified
- AND no error is returned

#### Scenario: Render pie chart
- GIVEN a RenderingContext with Page set and a PieChart with data
- WHEN `RenderPieChart(ctx, chart, bounds)` is called
- THEN pie slices are drawn with correct angles proportional to values
- AND slice colors match series colors
- AND no error is returned

#### Scenario: Render chart title
- GIVEN a RenderingContext and a chart with title "Sales Data"
- WHEN `RenderChartTitle(ctx, "Sales Data", bounds)` is called
- THEN the title text is drawn centered above the chart area

### Requirement: Shape Rendering (Enabled)

The system SHALL render shapes via the enabled shape_renderer.go.

#### Scenario: Render shape with geometry
- GIVEN a RenderingContext with Page set and a Shape with preset geometry
- WHEN `RenderShape(ctx, shape, bounds)` is called
- THEN the shape geometry is rendered within bounds
- AND fill and stroke are applied per shape properties

### Requirement: Fill Rendering (Enabled)

The system SHALL render fills via the enabled fill_renderer.go.

#### Scenario: Render solid fill
- GIVEN a RenderingContext with Page set and a SolidFill with RGB color
- WHEN `RenderSolidFill(ctx, fill)` is called
- THEN `ctx.Page.SetFillColor()` is called with correct RGB values

#### Scenario: Render gradient fill
- GIVEN a RenderingContext with Page set and a GradientFill
- WHEN `RenderGradientFill(ctx, fill, bounds)` is called
- THEN gradient is approximated via multiple filled shapes
- AND colors transition smoothly across bounds

### Requirement: Stroke Rendering (Enabled)

The system SHALL render strokes via the enabled stroke_renderer.go.

#### Scenario: Render stroke
- GIVEN a RenderingContext with Page set and an Outline with width and color
- WHEN `RenderStroke(ctx, outline)` is called
- THEN `ctx.Page.SetStrokeColor()` and `SetLineWidth()` are called appropriately

### Requirement: Effects Rendering (Enabled)

The system SHALL render effects via the enabled effects_renderer.go.

#### Scenario: Render shadow
- GIVEN a RenderingContext with Page set and a Shadow effect
- WHEN `RenderShadow(ctx, shadow)` is called
- THEN a shadow is drawn offset from the shape
- AND shadow uses blur and transparency from effect definition

#### Scenario: Render glow
- GIVEN a RenderingContext with Page set and a Glow effect
- WHEN `RenderGlow(ctx, glow)` is called
- THEN a glow effect is rendered around the shape

### Requirement: Transform Rendering (Enabled)

The system SHALL render transforms via the enabled transform_renderer.go.

#### Scenario: Render transform
- GIVEN a RenderingContext with Page set and a Transform with rotation
- WHEN `RenderTransform(ctx, transform)` is called
- THEN `ctx.Page.Rotate()` is called with correct angle

### Requirement: Text in Shape (Enabled)

The system SHALL render text in shapes via the enabled text_in_shape.go.

#### Scenario: Render text in shape
- GIVEN a RenderingContext with Page set and a TextBody within shape bounds
- WHEN `RenderTextInShape(ctx, text, bounds)` is called
- THEN text is rendered within the shape bounds
- AND text wrapping respects bounds

### Requirement: Image Rendering (Enabled)

The system SHALL render images via the enabled image_renderer.go.

#### Scenario: Render image
- GIVEN a RenderingContext with Page set and a Blip (image reference)
- WHEN `RenderImage(ctx, blip, bounds)` is called
- THEN the image is drawn within bounds
- AND aspect ratio may be preserved per settings
