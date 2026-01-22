# Speaker Notes Specification

## ADDED

#### Scenario: Add speaker notes to slides
- **GIVEN** a slide part with editable access
- **WHEN** adding speaker notes with text content
- **THEN** the system creates a notes slide part
- **AND** links it to the parent slide
- **AND** stores formatted text content
- **AND** supports multi-paragraph notes

#### Scenario: Rich formatting in speaker notes
- **GIVEN** speaker notes exist for a slide
- **WHEN** applying formatting to note text
- **THEN** support bold, italic, underline, and strikethrough
- **AND** support font size, color, and family changes
- **AND** support bullet points and numbered lists
- **AND** allow hyperlinks within notes

#### Scenario: Notes master template
- **GIVEN** a presentation with editable access
- **WHEN** creating a notes master template
- **THEN** define layout for all notes slides
- **AND** set corporate branding (logo, headers, footers)
- **AND** manage slide thumbnail positioning
- **AND** control page numbering and date formatting

#### Scenario: Export notes to PDF
- **GIVEN** a presentation with speaker notes
- **WHEN** exporting to PDF with notes option
- **THEN** include notes content below slide images
- **AND** apply notes master formatting
- **AND** paginate correctly across multiple pages
- **AND** maintain all formatting and layout
