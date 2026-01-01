# Pdf Drawing Specification (Delta)

## ADDED Requirements

### Requirement: Diagram Placeholder Rendering

The system SHALL render SmartArt diagram placeholders in PDF output (Phase 1: full rendering deferred to future).

#### Scenario: Diagram bounding box placeholder
- GIVEN a SmartArt diagram in a document
- WHEN rendered to PDF
- THEN a dashed rectangle is drawn at the diagram's position and size
- AND the rectangle uses gray stroke color
- AND the dash pattern is [5, 3] (5 points line, 3 points space)

#### Scenario: Diagram label text
- GIVEN a SmartArt diagram placeholder
- WHEN rendered to PDF
- THEN "SmartArt Diagram" text is drawn inside the bounding box
- AND text is positioned at top-left with 10 point offset
- AND text uses default font at 10 point size

#### Scenario: Diagram state isolation
- GIVEN a SmartArt diagram being rendered
- WHEN placeholder is drawn
- THEN graphics state is pushed before rendering
- AND graphics state is popped after rendering
- AND diagram rendering does not affect subsequent elements

### Requirement: Diagram Renderer Integration

The system SHALL integrate diagram placeholder rendering into document PDF renderers.

#### Scenario: PowerPoint diagram in PDF
- GIVEN a PowerPoint slide with SmartArt diagram
- WHEN rendered to PDF
- THEN diagram placeholder appears at correct location on slide
- AND placeholder size matches original diagram bounds

#### Scenario: Word diagram in PDF
- GIVEN a Word document with SmartArt diagram
- WHEN rendered to PDF
- THEN diagram placeholder appears inline or anchored as specified
- AND text flow around diagram is preserved

#### Scenario: Excel diagram in PDF
- GIVEN an Excel worksheet with SmartArt diagram
- WHEN rendered to PDF
- THEN diagram placeholder appears at correct position on sheet
- AND diagram does not overlap cell content

### Requirement: Diagram Bounds Calculation

The system SHALL calculate diagram bounding boxes for placeholder rendering.

#### Scenario: Diagram bounds from anchor
- GIVEN a SmartArt diagram with anchor/position information
- WHEN bounding box is calculated
- THEN x, y position is extracted from anchor offset
- AND width, height are extracted from anchor extents
- AND coordinates are converted from EMU to PDF points

#### Scenario: Diagram bounds from shape
- GIVEN a SmartArt diagram embedded in a shape
- WHEN bounding box is calculated
- THEN bounds are derived from containing shape's transform (a:xfrm)
- AND offset (a:off) and extents (a:ext) are used

### Requirement: Diagram Visibility in PDF

The system SHALL ensure diagrams are represented in PDF output.

#### Scenario: Prevent blank space
- GIVEN a document with SmartArt diagram
- WHEN rendered to PDF without layout engine
- THEN placeholder prevents blank/white space in output
- AND users can identify where diagrams are located

#### Scenario: Diagram indication for users
- GIVEN a PDF with diagram placeholders
- WHEN viewed by user
- THEN placeholders clearly indicate SmartArt presence
- AND users understand full rendering requires layout engine (future)

### Requirement: Diagram Renderer Testability

The system SHALL enable testing of diagram placeholder rendering.

#### Scenario: Mock Page testing
- GIVEN a DiagramRenderer
- WHEN tested with MockPage
- THEN SetStrokeColor is called with gray color
- AND SetLineDashPattern is called with [5, 3] pattern
- AND DrawRectangle is called with correct bounds
- AND DrawText is called with "SmartArt Diagram" label

#### Scenario: Integration testing
- GIVEN a document with SmartArt
- WHEN rendered to PDF
- THEN generated PDF contains diagram placeholder
- AND PDF validates correctly with PDF tools (pdfinfo, qpdf)
