## ADDED Requirements

### Requirement: RTL Text Rendering
The system MUST support right-to-left text rendering for appropriate languages.

#### Scenario: Render Hebrew, Arabic, Persian text
- **GIVEN** text in RTL languages
- **WHEN** adding text to slides
- **THEN** text is rendered right-to-left by default
- **AND** bidirectional text mixing with LTR content is handled correctly
- **AND** RTL paragraph direction is supported

#### Scenario: Complex Script Support
- **GIVEN** text requiring complex script handling
- **WHEN** rendering Arabic or similar scripts
- **THEN** contextual forms (initial, medial, final) are applied
- **AND** ligatures are handled correctly
- **AND** diacritic placement is managed properly