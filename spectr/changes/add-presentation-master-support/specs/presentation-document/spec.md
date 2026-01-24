## ADDED Requirements

### Requirement: Slide Master Implementation
The system SHALL support slide master creation and management with proper hierarchy, layout definitions, and property inheritance.

#### Scenario: Create slide master
- WHEN a slide master is created
- THEN it contains layout references
- AND text styles are defined
- AND placeholder shapes are configured

#### Scenario: Apply layout to slide
- WHEN a slide layout is applied to a presentation slide
- THEN placeholder positions and sizes are set from the layout
- AND text properties inherit from master
- AND the slide reflects all master formatting

### Requirement: Theme Color and Font Scheme
The system SHALL manage theme colors and fonts with variant support and property resolution.

#### Scenario: Define theme colors
- WHEN a color scheme is created in a theme
- THEN accent colors and scheme colors are definable
- AND colors are referenceable by theme index
- AND color overrides are supported

#### Scenario: Apply theme to presentation
- WHEN a theme is applied to a presentation
- THEN all slides inherit the theme colors
- AND all shapes use theme fonts
- AND theme overrides can be applied per slide

### Requirement: Master Inheritance Chain
The system SHALL properly resolve property inheritance from slide master through layout to individual slides.

#### Scenario: Property resolution
- WHEN a property is requested on a slide
- AND the property is not set on the slide
- THEN the property is resolved from the applied layout
- THEN the property is resolved from the slide master
