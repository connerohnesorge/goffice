## ADDED Requirements

### Requirement: Style Resolution Engine
The system SHALL resolve style properties through inheritance chains with proper cascading of defaults.

#### Scenario: Resolve paragraph style
- WHEN a paragraph style property is requested
- AND the property is not set on the style
- THEN the property is resolved from the base style (basedOn)
- THEN from the default style if set
- THEN from built-in defaults

#### Scenario: Apply style to element
- WHEN a style is applied to a paragraph
- THEN all inherited properties are available
- AND local property overrides take precedence
- AND theme references are resolved

### Requirement: Linked Styles
The system SHALL support paragraph styles linked to character styles for partial style application.

#### Scenario: Apply linked style
- WHEN a linked paragraph style is applied
- THEN paragraph properties apply to the paragraph
- THEN the linked character style applies to new runs
- AND existing character styling is preserved

### Requirement: Default Style Handling
The system SHALL provide built-in default styles with proper hierarchy (Normal, Heading 1-9, etc.).

#### Scenario: Access default style
- WHEN the Normal style is referenced
- THEN built-in defaults are loaded if not defined
- AND the style is available even without explicit definition

### Requirement: Theme Color Resolution
The system SHALL resolve theme color references to actual RGB values with tint/shade adjustments.

#### Scenario: Resolve theme color
- WHEN a theme color reference is encountered
- THEN the color index is mapped to the theme
- THEN tint or shade modifiers are applied
- THEN the final RGB color is computed
