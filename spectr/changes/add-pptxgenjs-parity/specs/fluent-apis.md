# Fluent API Enhancement Specification

## MODIFIED

#### Scenario: Chainable slide creation
- **GIVEN** a presentation document
- **WHEN** creating slides and adding content
- **THEN** provide fluent methods that return self or context for chaining
- **AND** slide.AddText().SetFont().SetColor() pattern
- **AND** slide.AddShape().SetSize().SetPosition() pattern
- **AND** reduce need for temporary variables
- **AND** enable single-statement slide construction

#### Scenario: Builder pattern for complex objects
- **GIVEN** complex objects like charts or SmartArt
- **WHEN** configuring multiple properties
- **THEN** provide builder APIs with Set* methods returning builder
- **AND** support Build() method to create final object
- **AND** allow incremental configuration
- **AND** support validation during build process

#### Scenario: Sensible defaults and auto-configuration
- **GIVEN** common slide creation scenarios
- **WHEN** creating slides without explicit configuration
- **THEN** auto-position shapes in flow layout
- **AND** select appropriate default fonts and colors
- **AND** automatically size text boxes to content
- **AND** provide convention-over-configuration patterns

#### Scenario: Helper methods for common patterns
- **GIVEN** frequently used slide patterns
- **WHEN** creating standard slide types
- **THEN** provide AddTitleSlide(title, subtitle) methods
- **AND** provide AddContentSlide(title, bulletPoints) methods
- **AND** support AddImageSlide(title, imagePath) patterns
- **AND** enable quick chart slides with AddChartSlide()

#### Scenario: Collection manipulation helpers
- **GIVEN** collections of slides or shapes
- **WHEN** performing bulk operations
- **THEN** provide ForEach, Map, Filter operations
- **AND** support RemoveWhere, AddRange methods
- **AND** enable LINQ-style query operations
- **AND** maintain proper collection change notifications
