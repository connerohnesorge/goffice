# Wordprocessing Styles Specification

## Requirements

### Requirement: Styles Root Element
The system SHALL provide a Styles element as the root of the styles part.

#### Scenario: Access document defaults
- GIVEN a Styles element
- WHEN DocDefaults() is called
- THEN the default paragraph and run formatting is returned

#### Scenario: Access latent styles
- GIVEN a Styles element
- WHEN LatentStyles() is called
- THEN latent style definitions are returned

#### Scenario: Enumerate all styles
- GIVEN a Styles element
- WHEN Elements[Style]() is called
- THEN all Style children are returned

#### Scenario: Get style by ID
- GIVEN a Styles element
- WHEN GetStyleById(styleId) is called
- THEN the Style with that ID is returned

#### Scenario: Get style by name
- GIVEN a Styles element
- WHEN GetStyleByName(name) is called
- THEN the Style with that name is returned

#### Scenario: Get styles by type
- GIVEN a Styles element
- WHEN GetStylesByType(StyleType.Paragraph) is called
- THEN all paragraph styles are returned

#### Scenario: Add style
- GIVEN a Styles element
- WHEN AppendChild(NewStyle(id, name, type)) is called
- THEN a new style is added

### Requirement: Style Element
The system SHALL provide a Style element for style definitions.

#### Scenario: Style ID
- GIVEN a Style element
- WHEN StyleId() is called
- THEN the unique style identifier is returned

#### Scenario: Style name
- GIVEN a Style element
- WHEN StyleName() is called
- THEN the human-readable style name is returned

#### Scenario: Style type
- GIVEN a Style element
- WHEN Type() is called
- THEN Paragraph, Character, Table, or Numbering is returned

#### Scenario: Based on style
- GIVEN a Style element
- WHEN BasedOn() is called
- THEN the parent style ID is returned

#### Scenario: Set based on
- GIVEN a Style element
- WHEN SetBasedOn("Normal") is called
- THEN the style inherits from Normal

#### Scenario: Next style
- GIVEN a Style element
- WHEN NextParagraphStyle() is called
- THEN the style for next paragraph is returned

#### Scenario: Linked style
- GIVEN a Style element
- WHEN LinkedStyle() is called
- THEN the linked character/paragraph style ID is returned

#### Scenario: Default style flag
- GIVEN a Style element
- WHEN Default() is called
- THEN whether this is the default style of its type is returned

#### Scenario: Custom style flag
- GIVEN a Style element
- WHEN CustomStyle() is called
- THEN whether this is a user-defined style is returned

#### Scenario: Style paragraph properties
- GIVEN a Style element of type Paragraph
- WHEN StyleParagraphProperties() is called
- THEN paragraph formatting for this style is returned

#### Scenario: Style run properties
- GIVEN a Style element
- WHEN StyleRunProperties() is called
- THEN character formatting for this style is returned

#### Scenario: Style table properties
- GIVEN a Style element of type Table
- WHEN StyleTableProperties() is called
- THEN table formatting for this style is returned

#### Scenario: Style table row properties
- GIVEN a Style element of type Table
- WHEN TableStyleRowProperties() is called
- THEN row formatting for this style is returned

#### Scenario: Style table cell properties
- GIVEN a Style element of type Table
- WHEN TableStyleCellProperties() is called
- THEN cell formatting for this style is returned

#### Scenario: UI priority
- GIVEN a Style element
- WHEN UIPriority() is called
- THEN the sort order in style gallery is returned

#### Scenario: Quick format flag
- GIVEN a Style element
- WHEN QuickFormat() is called
- THEN whether style appears in quick styles is returned

#### Scenario: Hidden flag
- GIVEN a Style element
- WHEN SemiHidden() is called
- THEN whether style is hidden from UI is returned

#### Scenario: Unhide when used
- GIVEN a Style element
- WHEN UnhideWhenUsed() is called
- THEN whether style becomes visible when used is returned

### Requirement: DocDefaults Element
The system SHALL provide DocDefaults for document-wide default formatting.

#### Scenario: Run properties default
- GIVEN a DocDefaults element
- WHEN RunPropertiesDefault() is called
- THEN default character formatting is returned

#### Scenario: Paragraph properties default
- GIVEN a DocDefaults element
- WHEN ParagraphPropertiesDefault() is called
- THEN default paragraph formatting is returned

### Requirement: LatentStyles Element
The system SHALL provide LatentStyles for built-in style behavior.

#### Scenario: Default latent style settings
- GIVEN a LatentStyles element
- WHEN DefSemiHidden(), DefUnhideWhenUsed(), DefQFormat() are accessed
- THEN default behavior for latent styles is returned

#### Scenario: Exception styles
- GIVEN a LatentStyles element
- WHEN Elements[LatentStyleException]() is called
- THEN specific style overrides are returned

### Requirement: Table Style Conditional Formatting
The system SHALL support conditional formatting in table styles.

#### Scenario: Access conditional formatting
- GIVEN a table Style element
- WHEN TableStyleProperties() is called
- THEN conditional formatting rules are returned

#### Scenario: Conditional type
- GIVEN a TableStyleProperties element
- WHEN Type() is called
- THEN FirstRow, LastRow, Band1Horz, etc. is returned

#### Scenario: Conditional run properties
- GIVEN a TableStyleProperties element
- WHEN RunProperties() is called
- THEN character formatting for condition is returned

#### Scenario: Conditional paragraph properties
- GIVEN a TableStyleProperties element
- WHEN ParagraphProperties() is called
- THEN paragraph formatting for condition is returned

#### Scenario: Conditional table cell properties
- GIVEN a TableStyleProperties element
- WHEN TableCellProperties() is called
- THEN cell formatting for condition is returned

### Requirement: Style Inheritance Resolution
The system SHALL support resolving inherited style properties.

#### Scenario: Resolve paragraph formatting
- GIVEN a style with basedOn chain
- WHEN ResolveProperties() is called
- THEN merged properties from inheritance chain is returned

#### Scenario: Property override
- GIVEN a style inheriting from another
- WHEN both define the same property
- THEN the child style's value takes precedence

#### Scenario: Default fallback
- GIVEN a style without explicit property
- WHEN property is accessed
- THEN DocDefaults value is used as fallback

### Requirement: Built-in Style IDs
The system SHALL define constants for built-in style IDs.

#### Scenario: Common style IDs
- GIVEN need to reference built-in styles
- WHEN using style constants
- THEN Normal, Heading1-9, Title, Subtitle, NoSpacing, etc. are available

#### Scenario: Table style IDs
- GIVEN need for table styles
- WHEN using style constants
- THEN TableGrid, TableNormal, etc. are available

#### Scenario: Character style IDs
- GIVEN need for character styles
- WHEN using style constants
- THEN DefaultParagraphFont, Hyperlink, etc. are available

### Requirement: Create Common Styles
The system SHALL provide factory methods for common styles.

#### Scenario: Create paragraph style
- GIVEN style ID, name, and formatting
- WHEN NewParagraphStyle(id, name, props) is called
- THEN a paragraph style is created

#### Scenario: Create character style
- GIVEN style ID, name, and formatting
- WHEN NewCharacterStyle(id, name, props) is called
- THEN a character style is created

#### Scenario: Create table style
- GIVEN style ID, name, and formatting
- WHEN NewTableStyle(id, name, props) is called
- THEN a table style is created

#### Scenario: Create heading style
- GIVEN heading level
- WHEN NewHeadingStyle(level) is called
- THEN a properly configured heading style is created

