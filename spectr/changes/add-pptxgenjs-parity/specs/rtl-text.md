# RTL Text Support Specification

## ADDED

#### Scenario: Basic RTL text rendering
- **GIVEN** a slide with text shapes
- **WHEN** adding Hebrew, Arabic, or Persian text content
- **THEN** render text right-to-left by default for RTL languages
- **AND** correctly handle bidirectional text mixing with LTR content
- **AND** support RTL paragraph direction
- **AND** maintain proper character ordering and joining

#### Scenario: RTL paragraph formatting
- **GIVEN** RTL text content in slides
- **WHEN** setting paragraph properties
- **THEN** support right-aligned, center-aligned, and justified text
- **AND** correctly indent from the right margin
- **AND** handle RTL bullet points and numbering
- **AND** manage line breaking rules for RTL languages

#### Scenario: Mixed LTR and RTL text (bidirectional)
- **GIVEN** text containing both LTR and RTL content
- **WHEN** rendering the text
- **THEN** automatically detect text direction per paragraph
- **AND** handle language-specific character reordering (Arabic shaping)
- **AND** manage embedding levels for proper text flow
- **AND** support explicit direction overrides

#### Scenario: Number and date formatting in RTL
- **GIVEN** RTL text with numbers and dates
- **WHEN** displaying numeric content
- **THEN** properly place Western numbers in RTL context
- **AND** support Arabic-Indic numerals
- **AND** format dates according to locale settings
- **AND** handle special characters and punctuation ordering

#### Scenario: Complex script languages
- **GIVEN** text in complex script languages
- **WHEN** rendering to slides
- **THEN** support contextual forms (Arabic initial, medial, final forms)
- **AND** handle ligatures correctly
- **AND** manage diacritic placement above/below letters
- **AND** support language-specific line breaking and kashida justification
