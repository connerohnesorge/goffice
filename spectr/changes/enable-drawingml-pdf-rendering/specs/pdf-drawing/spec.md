# Pdf Drawing Specification (Delta)

## MODIFIED Requirements

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

## ADDED Requirements

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
