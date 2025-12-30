# Pdf Core Specification (Delta)

## ADDED Requirements

### Requirement: Page Drawing Abstraction
The system SHALL provide a Page interface for drawing graphics primitives onto PDF pages.

#### Scenario: Draw filled rectangle
- GIVEN a Page instance
- WHEN DrawRectangle is called with coordinates and dimensions
- THEN a filled rectangle is drawn using the current fill color

#### Scenario: Draw stroked rectangle
- GIVEN a Page instance with stroke color set
- WHEN DrawRectangle is called after setting no fill
- THEN a rectangle outline is drawn with the current stroke color

#### Scenario: Draw ellipse
- GIVEN a Page instance  
- WHEN DrawEllipse is called with center and radii
- THEN an ellipse is drawn using current fill/stroke settings

#### Scenario: Draw arbitrary path
- GIVEN a PathBuilder with move, line, and curve commands
- WHEN DrawPath is called on Page
- THEN the path is rendered to PDF using current fill/stroke

### Requirement: Graphics State Management
The system SHALL manage PDF graphics state through a stack-based API.

#### Scenario: Push graphics state
- GIVEN a Page with current fill color red
- WHEN PushState is called
- THEN current state is saved to stack

#### Scenario: Pop graphics state
- GIVEN a Page with pushed state (fill=red) and current state (fill=blue)
- WHEN PopState is called
- THEN fill color reverts to red

#### Scenario: Nested state stack
- GIVEN multiple PushState calls
- WHEN PopState is called multiple times
- THEN states are restored in LIFO order

#### Scenario: Transform matrix
- GIVEN a Page instance
- WHEN Transform is called with rotation matrix
- THEN subsequent drawing operations are rotated

### Requirement: Color and Stroke Configuration
The system SHALL allow setting fill color, stroke color, and line properties.

#### Scenario: Set fill color
- GIVEN a Page instance
- WHEN SetFillColor is called with RGB color
- THEN subsequent filled shapes use that color

#### Scenario: Set stroke color
- GIVEN a Page instance
- WHEN SetStrokeColor is called with RGB color
- THEN subsequent stroked shapes use that color

#### Scenario: Set line width
- GIVEN a Page instance
- WHEN SetLineWidth is called with width in points
- THEN subsequent stroked shapes use that line width

#### Scenario: Set dash pattern
- GIVEN a Page instance
- WHEN SetLineDashPattern is called with pattern array [5, 3] and phase 0
- THEN subsequent lines are drawn with 5-point dashes and 3-point gaps

#### Scenario: Clear dash pattern (solid line)
- GIVEN a Page with dashed lines
- WHEN SetLineDashPattern is called with empty array
- THEN subsequent lines are drawn solid

### Requirement: Content Embedding
The system SHALL support embedding and drawing images and text.

#### Scenario: Add image
- GIVEN a Page instance and an image object
- WHEN AddImage is called with coordinates and dimensions
- THEN image is embedded in PDF and drawn at specified location

#### Scenario: Draw text
- GIVEN a Page instance
- WHEN DrawText is called with text, position, font, and size
- THEN text is rendered at specified location with specified font

### Requirement: RenderingContext Integration
The system SHALL integrate Page with the existing RenderingContext.

#### Scenario: Access Page from context
- GIVEN a RenderingContext
- WHEN Page field is accessed
- THEN a valid Page instance is returned

#### Scenario: Page lifecycle
- GIVEN a RenderingContext for a PDF document
- WHEN rendering begins
- THEN a new Page instance is created for each PDF page

#### Scenario: Coordinate system consistency
- GIVEN a RenderingContext with page dimensions
- WHEN Page methods are called with EMU coordinates
- THEN coordinates are correctly converted to PDF points with origin transformation
