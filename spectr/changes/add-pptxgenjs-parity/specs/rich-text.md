# Rich Text and Text Formatting Enhancement Specification

## MODIFIED

#### Scenario: Multiple formatting runs in single text block
- **GIVEN** a text shape or paragraph
- **WHEN** adding text with varying formatting
- **THEN** support breaking text into multiple runs with different styles
- **AND** allow intermixing bold, italic, colors, and fonts within one paragraph
- **AND** maintain proper text flow and spacing between runs
- **AND** support hypertext runs within mixed formatting

#### Scenario: Superscript and subscript support
- **GIVEN** text content requiring special positioning
- **WHEN** applying superscript or subscript formatting
- **THEN** support `a²` (superscript) for mathematical expressions
- **AND** support `H₂O` (subscript) for chemical formulas
- **AND** maintain proper font sizing and baseline offset
- **AND** mix with other formatting options

#### Scenario: Adjustable character spacing and kerning
- **GIVEN** text requiring character-level spacing adjustments
- **WHEN** applying character formatting
- **THEN** support increased/decreased character spacing
- **AND** enable kerning for better visual appearance
- **AND** apply ligature settings for supported fonts
- **AND** handle spacing for different font sizes appropriately

#### Scenario: Text stroke and outline effects
- **GIVEN** text needing outline effects
- **WHEN** applying text formatting
- **THEN** support colored text outlines
- **AND** adjust outline thickness and dash styles
- **AND** create hollow text effects
- **AND** combine with fill colors for double-stroke effects

#### Scenario: Text effects and transformations
- **GIVEN** text shapes in slides
- **WHEN** applying visual effects
- **THEN** support text shadows with adjustable blur and offset
- **AND** apply reflections below text
- **AND** enable glow effects with custom colors
- **AND** rotate text to any angle while maintaining editability

#### Scenario: Paragraph borders and fills
- **GIVEN** paragraph-level text blocks
- **WHEN** applying paragraph formatting
- **THEN** support border styling (solid, dashed, dotted)
- **AND** apply rounded corners to paragraph boxes
- **AND** add gradient fills within text areas
- **AND** add patterns and textures to paragraph backgrounds

#### Scenario: Advanced list formatting
- **GIVEN** lists in text content
- **WHEN** formatting list appearance
- **THEN** support custom bullet characters and images
- **AND** enable multilevel numbering schemes
- **AND** adjust spacing between bullets and text
- **AND** change bullet size and color independently from text
