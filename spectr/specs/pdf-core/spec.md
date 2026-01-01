# Pdf Core Specification

## Requirements

### Requirement: PDF Document Creation
The system SHALL provide a PDF document writer that produces valid PDF 1.7 output.

#### Scenario: Create empty PDF
- GIVEN a new PDF document
- WHEN the document is saved
- THEN a valid PDF file is produced with correct header and trailer

#### Scenario: Create multi-page PDF
- GIVEN a PDF document with multiple pages
- WHEN pages are added with content
- THEN all pages are included in correct order in the output

### Requirement: Font Embedding
The system SHALL embed fonts in PDF output to ensure consistent rendering across systems.

#### Scenario: Embed TrueType font
- GIVEN a document using a TrueType font
- WHEN rendered to PDF
- THEN the font is embedded as a subset containing only used glyphs

#### Scenario: Font fallback
- GIVEN a document using an unavailable font
- WHEN rendered to PDF
- THEN a suitable fallback font is substituted with a warning

### Requirement: Font Metrics Extraction
The system SHALL extract font metrics from TrueType/OpenType fonts for text layout.

#### Scenario: Extract glyph widths
- GIVEN a TrueType font file
- WHEN metrics are requested
- THEN accurate glyph widths are returned for all characters

#### Scenario: Extract vertical metrics
- GIVEN a font file
- WHEN metrics are requested
- THEN ascender, descender, line gap, and cap height are returned

### Requirement: Image Embedding
The system SHALL embed images in PDF output with appropriate compression.

#### Scenario: Embed JPEG image
- GIVEN a JPEG image in the document
- WHEN rendered to PDF
- THEN the image is embedded using DCT encoding (pass-through)

#### Scenario: Embed PNG image
- GIVEN a PNG image in the document
- WHEN rendered to PDF
- THEN the image is embedded with appropriate encoding and alpha handling

### Requirement: Coordinate Transformation
The system SHALL correctly transform between OOXML EMU coordinates and PDF points.

#### Scenario: EMU to points conversion
- GIVEN a dimension in EMU (914400 = 1 inch)
- WHEN converted to PDF points
- THEN the result is 72 points per inch

#### Scenario: Origin transformation
- GIVEN OOXML coordinates (origin top-left)
- WHEN rendered to PDF
- THEN coordinates are transformed to PDF origin (bottom-left)

### Requirement: Color Handling
The system SHALL correctly render colors from OOXML color specifications.

#### Scenario: RGB color
- GIVEN an RGB color value
- WHEN rendered to PDF
- THEN the color is applied correctly

#### Scenario: Theme color
- GIVEN a theme color reference (e.g., accent1)
- WHEN rendered to PDF
- THEN the color is resolved from the document theme and applied

#### Scenario: Transparency
- GIVEN an element with transparency/opacity
- WHEN rendered to PDF
- THEN appropriate alpha channel or blend mode is applied

### Requirement: Render Options
The system SHALL support configuration options for PDF rendering.

#### Scenario: Set output quality
- GIVEN render options specifying image quality
- WHEN rendering with high quality
- THEN images are rendered at higher DPI

#### Scenario: Enable accessibility
- GIVEN render options enabling accessibility
- WHEN rendering
- THEN tagged PDF output is produced

### Requirement: Page Drawing Abstraction
The system SHALL provide a Page interface for drawing graphics primitives onto PDF pages.

#### Scenario: Draw filled rectangle
- GIVEN a Page instance
- WHEN DrawRectangle is called with (x, y, width, height, fill=true, stroke=false)
- THEN a filled rectangle is drawn using the current fill color

#### Scenario: Draw stroked rectangle
- GIVEN a Page instance with stroke color set
- WHEN DrawRectangle is called with (x, y, width, height, fill=false, stroke=true)
- THEN a rectangle outline is drawn with the current stroke color

#### Scenario: Draw filled and stroked rectangle
- GIVEN a Page instance with fill and stroke colors set
- WHEN DrawRectangle is called with (x, y, width, height, fill=true, stroke=true)
- THEN a rectangle is drawn with both fill and stroke

#### Scenario: Draw circle
- GIVEN a Page instance
- WHEN DrawCircle is called with (cx, cy, radius, fill=true, stroke=true)
- THEN a circle is drawn with both fill and stroke

#### Scenario: Draw ellipse
- GIVEN a Page instance
- WHEN DrawEllipse is called with (cx, cy, rx, ry, fill=true, stroke=false)
- THEN an ellipse is drawn using current fill color

#### Scenario: Draw arbitrary path
- GIVEN a PathBuilder with move, line, and curve commands
- WHEN DrawPath is called on Page
- THEN the path is rendered to PDF using current fill/stroke

#### Scenario: Write raw PDF content
- GIVEN a PathBuilder that generates PDF operators
- WHEN WriteContent is called with path.Stroke()
- THEN the raw PDF operators are written to the content stream

### Requirement: Graphics State Management
The system SHALL manage PDF graphics state through a stack-based API.

#### Scenario: Save graphics state
- GIVEN a Page with current fill color red
- WHEN SaveGraphicsState is called
- THEN current state is saved to stack (PDF 'q' operator)

#### Scenario: Restore graphics state
- GIVEN a Page with saved state (fill=red) and current state (fill=blue)
- WHEN RestoreGraphicsState is called
- THEN fill color reverts to red (PDF 'Q' operator)

#### Scenario: Nested state stack
- GIVEN multiple SaveGraphicsState calls
- WHEN RestoreGraphicsState is called multiple times
- THEN states are restored in LIFO order

#### Scenario: Transform matrix
- GIVEN a Page instance
- WHEN Transform is called with rotation matrix
- THEN subsequent drawing operations are rotated

### Requirement: Color and Stroke Configuration
The system SHALL allow setting fill color, stroke color, and line properties.

#### Scenario: Set fill color
- GIVEN a Page instance
- WHEN SetFillColor is called with (r=0.5, g=0.5, b=0.5)
- THEN subsequent filled shapes use that RGB color

#### Scenario: Set stroke color
- GIVEN a Page instance
- WHEN SetStrokeColor is called with (r=1.0, g=0.0, b=0.0)
- THEN subsequent stroked shapes use red color

#### Scenario: Set line width
- GIVEN a Page instance
- WHEN SetLineWidth is called with width=2.0 points
- THEN subsequent stroked shapes use that line width

#### Scenario: Set dash pattern
- GIVEN a Page instance
- WHEN SetLineDashPattern is called with pattern=[5, 3] and phase=0
- THEN subsequent lines are drawn with 5-point dashes and 3-point gaps

#### Scenario: Clear dash pattern (solid line)
- GIVEN a Page with dashed lines
- WHEN SetLineDashPattern is called with empty pattern=[]
- THEN subsequent lines are drawn solid

#### Scenario: Set line cap style
- GIVEN a Page instance
- WHEN SetLineCap is called with LineCap.Round (value=1)
- THEN subsequent stroked lines use rounded end caps

#### Scenario: Set line join style
- GIVEN a Page instance
- WHEN SetLineJoin is called with LineJoin.Bevel (value=2)
- THEN subsequent stroked lines use beveled joins

### Requirement: Content Embedding
The system SHALL support embedding and drawing images and text.

#### Scenario: Add image
- GIVEN a Page instance and an image object
- WHEN AddImage is called with (img, x, y, width, height)
- THEN image is embedded in PDF and drawn at specified location

#### Scenario: Set font
- GIVEN a Page instance
- WHEN SetFont is called with (name="Helvetica", size=12.0)
- THEN the font is set for subsequent DrawText calls

#### Scenario: Draw text
- GIVEN a Page instance with font set
- WHEN DrawText is called with (text="Hello", x=100.0, y=200.0)
- THEN text is rendered at specified location with current font

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

#### Scenario: Coordinate conversion edge cases - negative coordinates
- GIVEN a Page with coordinate conversion
- WHEN DrawRectangle is called with negative x or y
- THEN the shape is rendered correctly (supports shapes extending beyond page)

#### Scenario: Coordinate conversion edge cases - very large coordinates
- GIVEN a Page with coordinate conversion
- WHEN DrawRectangle is called with very large EMU values (>1 billion)
- THEN the conversion produces valid PDF point values without overflow

#### Scenario: Coordinate conversion edge cases - zero dimensions
- GIVEN a Page with coordinate conversion
- WHEN DrawRectangle is called with width=0 or height=0
- THEN the degenerate shape is handled gracefully (renders as line or point)
