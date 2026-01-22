# HTML to PowerPoint Conversion Specification

## ADDED

#### Scenario: Convert HTML tables to PowerPoint tables
- **GIVEN** valid HTML table markup (table, tr, td, th)
- **WHEN** converting to PowerPoint format
- **THEN** preserve table structure and cell relationships
- **AND** maintain row and column counts
- **AND** respect rowspan and colspan attributes
- **AND** convert basic styling to PowerPoint equivalents

#### Scenario: Convert HTML text formatting
- **GIVEN** HTML text with formatting tags
- **WHEN** converting for presentation
- **THEN** map `<b>`, `<strong>` to bold text runs
- **AND** map `<i>`, `<em>` to italic text runs
- **AND** map `<u>` to underline styling
- **AND** map `<s>`, `<strike>` to strikethrough
- **AND** handle nested formatting correctly

#### Scenario: Convert HTML lists to PowerPoint
- **GIVEN** HTML list markup (`<ul>`, `<ol>`, `<li>`)
- **WHEN** converting to presentation format
- **THEN** create bulleted lists for `<ul>` with proper indentation
- **AND** create numbered lists for `<ol>` with correct numbering
- **AND** preserve nested list structure
- **AND** maintain list styles and indentation levels

#### Scenario: Handle CSS inline styles
- **GIVEN** HTML elements with inline CSS styles
- **WHEN** converting to PowerPoint
- **THEN** parse and apply font-family, size, color styles
- **AND** handle background-color for tables and cells
- **AND** respect text-align (left, center, right, justify)
- **AND** apply border styles and colors
- **AND** handle padding and margin appropriately

#### Scenario: Convert HTML links and images
- **GIVEN** HTML with `<a>` tags and `<img>` tags
- **WHEN** converting to PowerPoint
- **THEN** preserve hyperlinks in converted text
- **AND** embed images referenced by `<img>` tags
- **AND** handle relative and absolute URLs
- **AND** fall back gracefully for unsupported image formats

#### Scenario: Complex HTML structures
- **GIVEN** nested HTML structures (headers, nested tables)
- **WHEN** converting to presentation
- **THEN** handle `<h1>`-`<h6>` with appropriate font sizing
- **AND** manage nested tables with proper layout
- **AND** convert `<br>` tags to line breaks
- **AND** handle special character entities

#### Scenario: Encoding and character handling
- **GIVEN** HTML content with various encodings
- **WHEN** converting to PowerPoint
- **THEN** detect and handle UTF-8, UTF-16, and other encodings
- **AND** convert HTML entities to Unicode characters
- **AND** maintain character encoding in output
- **AND** handle BOM markers appropriately
