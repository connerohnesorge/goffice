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

