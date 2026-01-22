# Slide Comment Support Specification

## ADDED

#### Scenario: Add comments to slides
- **GIVEN** a slide with editable access
- **WHEN** adding a comment with text and author
- **THEN** create comment part with proper XML structure
- **AND** position comment on slide (x, y coordinates)
- **AND** store author name, initials, and timestamp
- **AND** support multiple comments per slide

#### Scenario: Threaded comment replies
- **GIVEN** existing comments on a slide
- **WHEN** adding replies to comments
- **THEN** maintain comment hierarchy (parent-child relationships)
- **AND** show reply count and thread expansion
- **AND** preserve reply ordering by timestamp
- **AND** support multiple nesting levels

#### Scenario: Comment formatting and text
- **GIVEN** comments with text content
- **WHEN** formatting comments
- **THEN** support basic text formatting (bold, italic)
- **AND** allow hyperlinks in comment text
- **AND** support multiple paragraphs in comments
- **AND** preserve line breaks and spacing

#### Scenario: Comment positioning and visibility
- **GIVEN** comments on slide elements
- **WHEN** managing comment display
- **THEN** attach comments to specific shapes or slide positions
- **AND** show/hide comment indicators
- **AND** support comment anchoring to slide elements
- **AND** manage comment visibility in presentation mode

#### Scenario: Comment authors and management
- **GIVEN** multiple users collaborating
- **WHEN** managing comment authors
- **THEN** track comment author information
- **AND** support initials and color coding by author
- **AND** show comment timestamps
- **AND** enable comment resolution status

#### Scenario: Export comments
- **GIVEN** presentations with comments
- **WHEN** exporting to PDF or printing
- **THEN** optionally include comments in output
- **AND** format comment printouts with author and date
- **AND** link comments to slide content
- **AND** provide comment summary pages
