# Pdf Presentation Specification

## Requirements

### Requirement: Presentation to PDF Conversion
The system SHALL convert PresentationML documents to PDF.

#### Scenario: Basic presentation conversion
- GIVEN a PresentationML document (.pptx)
- WHEN converted to PDF
- THEN all slides are rendered in order

#### Scenario: Slide selection
- GIVEN a presentation and a slide range
- WHEN converted with slide selection
- THEN only specified slides are included in output

### Requirement: Slide Rendering
The system SHALL render slides with full content.

#### Scenario: Slide dimensions
- GIVEN a presentation with specific slide size
- WHEN rendered
- THEN PDF pages match slide dimensions

#### Scenario: Slide background
- GIVEN a slide with background (solid, gradient, image, pattern)
- WHEN rendered
- THEN background is applied correctly

### Requirement: Shape Rendering
The system SHALL render presentation shapes.

#### Scenario: Text box
- GIVEN a text box shape with content
- WHEN rendered
- THEN text is displayed with correct formatting and position

#### Scenario: Rectangle shape
- GIVEN a rectangle with fill and outline
- WHEN rendered
- THEN shape is drawn with correct geometry and styling

#### Scenario: Preset shapes
- GIVEN a preset shape (arrow, star, callout, etc.)
- WHEN rendered
- THEN shape geometry matches PowerPoint rendering

#### Scenario: Shape positioning
- GIVEN shapes with specific positions and sizes
- WHEN rendered
- THEN shapes are placed at correct coordinates

### Requirement: Text Rendering in Shapes
The system SHALL render text within shapes correctly.

#### Scenario: Text alignment
- GIVEN a shape with text alignment settings
- WHEN rendered
- THEN text is aligned horizontally and vertically as specified

#### Scenario: Text autofit
- GIVEN a shape with autofit text (shrink to fit)
- WHEN rendered
- THEN text is scaled appropriately

#### Scenario: Bullet lists in shapes
- GIVEN a shape with bulleted text
- WHEN rendered
- THEN bullets and indentation are displayed correctly

### Requirement: Image Handling
The system SHALL render images on slides.

#### Scenario: Picture shape
- GIVEN an image on a slide
- WHEN rendered
- THEN image is displayed at correct position and size

#### Scenario: Cropped image
- GIVEN an image with cropping applied
- WHEN rendered
- THEN only the visible portion is shown

#### Scenario: Image effects
- GIVEN an image with effects (shadow, reflection)
- WHEN rendered
- THEN effects are approximated in output

### Requirement: Master Slide Application
The system SHALL apply master slide formatting.

#### Scenario: Master background
- GIVEN a slide using a master with background
- WHEN rendered
- THEN master background is applied under slide content

#### Scenario: Placeholder inheritance
- GIVEN a slide with placeholders inheriting master styling
- WHEN rendered
- THEN inherited styles are correctly applied

#### Scenario: Master shapes
- GIVEN a master slide with decorative shapes
- WHEN rendered
- THEN master shapes appear behind slide content

### Requirement: Layout Slide Application
The system SHALL apply layout slide elements.

#### Scenario: Layout placeholders
- GIVEN a slide using a layout with placeholders
- WHEN rendered
- THEN placeholder positions and styles are inherited

### Requirement: Table Rendering
The system SHALL render tables on slides.

#### Scenario: Slide table
- GIVEN a slide with a table
- WHEN rendered
- THEN table is rendered with cells, borders, and content

#### Scenario: Table styling
- GIVEN a table with style (banded rows, header row)
- WHEN rendered
- THEN table styling is applied correctly

### Requirement: Chart Rendering
The system SHALL render charts on slides.

#### Scenario: Embedded chart
- GIVEN a slide with an embedded chart
- WHEN rendered
- THEN chart is displayed with correct data and formatting

### Requirement: SmartArt Rendering
The system SHALL render SmartArt graphics.

#### Scenario: Process SmartArt
- GIVEN a slide with SmartArt (e.g., process diagram)
- WHEN rendered
- THEN SmartArt is decomposed and rendered as shapes

### Requirement: Notes Pages
The system SHALL optionally render notes pages.

#### Scenario: Notes page output
- GIVEN a presentation with speaker notes
- WHEN rendered with notes option enabled
- THEN each slide is rendered with its notes on a notes page layout

### Requirement: Handout Layout
The system SHALL support handout-style output.

#### Scenario: Multiple slides per page
- GIVEN a presentation and handout option (2, 3, 4, 6, 9 per page)
- WHEN rendered as handouts
- THEN multiple slides are arranged on each page

