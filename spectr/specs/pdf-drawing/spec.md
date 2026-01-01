# Pdf Drawing Specification

## Requirements

### Requirement: Shape Geometry Rendering
The system SHALL render DrawingML shape geometries to PDF paths using the Page abstraction.

#### Scenario: Preset shape with Page API
- GIVEN a preset shape type (e.g., rect, ellipse)
- WHEN rendered using ctx.Page methods
- THEN shape is drawn using DrawRectangle/DrawEllipse/DrawPath

### Requirement: Fill Rendering  
The system SHALL render DrawingML fill types using the Page abstraction.

#### Scenario: Solid fill with Page API
- GIVEN a shape with solid color fill
- WHEN rendered  
- THEN Page.SetFillColor is called followed by Page.DrawPath

#### Scenario: Linear gradient approximation
- GIVEN a shape with linear gradient fill
- WHEN rendered
- THEN gradient is approximated using multiple DrawRectangle calls with interpolated colors

### Requirement: Line/Stroke Rendering
The system SHALL render DrawingML outline properties using the Page abstraction.

#### Scenario: Solid line with Page API
- GIVEN a shape with solid outline
- WHEN rendered
- THEN Page.SetStrokeColor and Page.SetLineWidth are called before drawing

#### Scenario: Dashed line with Page API
- GIVEN a shape with dashed outline
- WHEN rendered
- THEN Page.SetLineDashPattern is called with appropriate pattern

### Requirement: Text in Shapes
The system SHALL render text content within shapes using the Page abstraction.

#### Scenario: Text body with Page API
- GIVEN a shape with text content
- WHEN rendered
- THEN Page.DrawText is called for each text run with correct positioning

### Requirement: Picture Rendering
The system SHALL render DrawingML pictures using the Page abstraction.

#### Scenario: Inline image with Page API
- GIVEN a picture element with image reference
- WHEN rendered
- THEN Page.AddImage is called with correct dimensions and positioning

### Requirement: Transform Rendering
The system SHALL apply DrawingML transforms using the Page abstraction.

#### Scenario: Rotation with Page API
- GIVEN a shape with rotation angle
- WHEN rendered
- THEN Page.PushState, Page.Transform with rotation matrix, Page.PopState are called

### Requirement: Effect Rendering
The system SHALL render DrawingML effects using the Page abstraction.

#### Scenario: Drop shadow with Page API
- GIVEN a shape with drop shadow effect
- WHEN rendered
- THEN shadow is drawn first using Page.SetFillColor with transparency, followed by shape

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
The system SHALL render DrawingML charts using the Page abstraction.

#### Scenario: Bar chart with Page API
- GIVEN a bar chart
- WHEN rendered
- THEN bars are drawn using Page.DrawRectangle for each data point

#### Scenario: Line chart with Page API
- GIVEN a line chart
- WHEN rendered
- THEN lines are drawn using Page.DrawPath connecting data points

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

### Requirement: Renderer File Re-enablement
The system SHALL have all DrawingML renderer files active (no .wip extensions).

#### Scenario: Chart renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN chart_renderer.go exists (not .wip)

#### Scenario: Fill renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN fill_renderer.go exists (not .wip)

#### Scenario: Shape renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN shape_renderer.go exists (not .wip)

#### Scenario: Stroke renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN stroke_renderer.go exists (not .wip)

#### Scenario: Image renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN image_renderer.go exists (not .wip)

#### Scenario: Text in shape renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN text_in_shape.go exists (not .wip)

#### Scenario: Transform renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN transform_renderer.go exists (not .wip)

#### Scenario: Effects renderer active
- GIVEN the pdf/drawing package
- WHEN inspecting files
- THEN effects_renderer.go exists (not .wip)

### Requirement: Integration with Document Renderers
The system SHALL integrate DrawingML rendering with Word, Excel, and PowerPoint PDF renderers.

#### Scenario: Excel chart rendering
- GIVEN an Excel document with a column chart
- WHEN rendered to PDF
- THEN chart appears in PDF output at correct location

#### Scenario: PowerPoint shape rendering
- GIVEN a PowerPoint slide with shapes
- WHEN rendered to PDF  
- THEN shapes appear with correct fills and strokes

#### Scenario: Word inline image rendering
- GIVEN a Word document with inline images
- WHEN rendered to PDF
- THEN images appear at correct positions with correct sizing
