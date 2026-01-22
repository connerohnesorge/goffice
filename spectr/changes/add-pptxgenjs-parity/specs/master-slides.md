# Master Slides Specification

## ADDED

#### Scenario: Create slide master templates
- **GIVEN** a presentation with editable access
- **WHEN** creating a slide master
- **THEN** define layout for title, content, and placeholder positions
- **AND** set background colors, gradients, or images
- **AND** manage corporate branding elements (logos, watermarks)
- **AND** define default fonts and color schemes
- **AND** support multiple master slides per presentation

#### Scenario: Apply slide master to slides
- **GIVEN** a presentation with slide masters
- **WHEN** creating a new slide or changing existing slide
- **THEN** allow selection of which master to apply
- **AND** inherit layout and formatting from master
- **AND** support overriding master elements on individual slides
- **AND** preserve slide-specific content

#### Scenario: Create custom slide layouts
- **GIVEN** an existing slide master
- **WHEN** creating custom layouts
- **THEN** support standard layouts (Title, Title+Content, Section Header, etc.)
- **AND** allow custom placeholder types and positions
- **AND** manage layout inheritance from master
- **AND** set layout-specific formatting

#### Scenario: Master slide inheritance and updates
- **GIVEN** presentation with slides using masters
- **WHEN** updating master slide design
- **THEN** propagate changes to all slides using that master
- **AND** preserve slide-specific overrides
- **AND** update slide numbers and footers automatically
- **AND** handle master deletion with fallback options

#### Scenario: Theme integration with masters
- **GIVEN** a presentation with themes and masters
- **WHEN** applying theme changes
- **THEN** update master colors, fonts, and effects
- **AND** cascade changes to layouts and slides
- **AND** support theme variants (dark/light)
- **AND** manage custom color palettes
