# Wordprocessing Properties Specification

## Requirements

### Requirement: ParagraphProperties Element
The system SHALL provide ParagraphProperties for paragraph formatting.

#### Scenario: Access paragraph style
- GIVEN a ParagraphProperties element
- WHEN ParagraphStyleId() is called
- THEN the style ID (e.g., "Heading1") is returned

#### Scenario: Set paragraph style
- GIVEN a ParagraphProperties element
- WHEN SetParagraphStyleId("Heading1") is called
- THEN the style reference is set

#### Scenario: Justification (alignment)
- GIVEN a ParagraphProperties element
- WHEN Justification() is called
- THEN the alignment (Left, Center, Right, Both) is returned

#### Scenario: Set justification
- GIVEN a ParagraphProperties element
- WHEN SetJustification(JustificationValues.Center) is called
- THEN center alignment is set

#### Scenario: Indentation
- GIVEN a ParagraphProperties element
- WHEN Indentation() is called
- THEN left, right, firstLine, and hanging indentation values are returned

#### Scenario: Set indentation
- GIVEN a ParagraphProperties element
- WHEN Indentation().Left.SetValue(720) is called
- THEN 720 twips (0.5 inch) left indent is set

#### Scenario: Spacing
- GIVEN a ParagraphProperties element
- WHEN SpacingBetweenLines() is called
- THEN before, after, and line spacing are returned

#### Scenario: Set spacing
- GIVEN a ParagraphProperties element
- WHEN SpacingBetweenLines().Before.SetValue(240) is called
- THEN 12pt space before paragraph is set

#### Scenario: Keep with next
- GIVEN a ParagraphProperties element
- WHEN KeepNext() is called
- THEN whether paragraph stays with next is returned

#### Scenario: Keep lines together
- GIVEN a ParagraphProperties element
- WHEN KeepLines() is called
- THEN whether lines stay together is returned

#### Scenario: Page break before
- GIVEN a ParagraphProperties element
- WHEN PageBreakBefore() is called
- THEN whether page break precedes paragraph is returned

#### Scenario: Widow/orphan control
- GIVEN a ParagraphProperties element
- WHEN WidowControl() is called
- THEN widow/orphan setting is returned

#### Scenario: Numbering properties
- GIVEN a ParagraphProperties element
- WHEN NumberingProperties() is called
- THEN the list numbering reference is returned

#### Scenario: Tab stops
- GIVEN a ParagraphProperties element
- WHEN Tabs() is called
- THEN custom tab stop definitions are returned

#### Scenario: Borders
- GIVEN a ParagraphProperties element
- WHEN ParagraphBorders() is called
- THEN paragraph border settings are returned

#### Scenario: Shading
- GIVEN a ParagraphProperties element
- WHEN Shading() is called
- THEN background shading settings are returned

#### Scenario: Outline level
- GIVEN a ParagraphProperties element
- WHEN OutlineLevel() is called
- THEN the outline/heading level is returned

#### Scenario: Run properties in paragraph
- GIVEN a ParagraphProperties element
- WHEN ParagraphMarkRunProperties() is called
- THEN formatting for paragraph mark is returned

### Requirement: RunProperties Element
The system SHALL provide RunProperties for character formatting.

#### Scenario: Bold
- GIVEN a RunProperties element
- WHEN Bold() is called
- THEN whether text is bold is returned

#### Scenario: Set bold
- GIVEN a RunProperties element
- WHEN SetBold(true) is called
- THEN bold formatting is applied

#### Scenario: Italic
- GIVEN a RunProperties element
- WHEN Italic() is called
- THEN whether text is italic is returned

#### Scenario: Set italic
- GIVEN a RunProperties element
- WHEN SetItalic(true) is called
- THEN italic formatting is applied

#### Scenario: Underline
- GIVEN a RunProperties element
- WHEN Underline() is called
- THEN underline style is returned

#### Scenario: Set underline
- GIVEN a RunProperties element
- WHEN SetUnderline(UnderlineValues.Single) is called
- THEN single underline is applied

#### Scenario: Strike through
- GIVEN a RunProperties element
- WHEN Strike() is called
- THEN strikethrough status is returned

#### Scenario: Double strike through
- GIVEN a RunProperties element
- WHEN DoubleStrike() is called
- THEN double strikethrough status is returned

#### Scenario: Font size
- GIVEN a RunProperties element
- WHEN FontSize() is called
- THEN font size in half-points is returned

#### Scenario: Set font size
- GIVEN a RunProperties element
- WHEN SetFontSize(24) is called
- THEN 12pt font size is set (24 half-points)

#### Scenario: Font name
- GIVEN a RunProperties element
- WHEN RunFonts() is called
- THEN font family names (ASCII, eastAsia, etc.) are returned

#### Scenario: Set font
- GIVEN a RunProperties element
- WHEN SetFont("Arial") is called
- THEN Arial font is applied

#### Scenario: Color
- GIVEN a RunProperties element
- WHEN Color() is called
- THEN text color value is returned

#### Scenario: Set color
- GIVEN a RunProperties element
- WHEN SetColor("FF0000") is called
- THEN red color is applied

#### Scenario: Highlight
- GIVEN a RunProperties element
- WHEN Highlight() is called
- THEN highlight color is returned

#### Scenario: Set highlight
- GIVEN a RunProperties element
- WHEN SetHighlight(HighlightColorValues.Yellow) is called
- THEN yellow highlight is applied

#### Scenario: All caps
- GIVEN a RunProperties element
- WHEN Caps() is called
- THEN whether all caps is enabled is returned

#### Scenario: Small caps
- GIVEN a RunProperties element
- WHEN SmallCaps() is called
- THEN whether small caps is enabled is returned

#### Scenario: Subscript/superscript
- GIVEN a RunProperties element
- WHEN VerticalTextAlignment() is called
- THEN Subscript, Superscript, or Baseline is returned

#### Scenario: Character spacing
- GIVEN a RunProperties element
- WHEN Spacing() is called
- THEN character spacing adjustment is returned

#### Scenario: Text position (raise/lower)
- GIVEN a RunProperties element
- WHEN Position() is called
- THEN vertical position offset is returned

#### Scenario: Scale (character width)
- GIVEN a RunProperties element
- WHEN CharacterScale() is called
- THEN width scaling percentage is returned

#### Scenario: Emboss and imprint
- GIVEN a RunProperties element
- WHEN Emboss() and Imprint() are called
- THEN emboss/imprint effects are returned

#### Scenario: Shadow
- GIVEN a RunProperties element
- WHEN Shadow() is called
- THEN shadow effect status is returned

#### Scenario: Outline
- GIVEN a RunProperties element
- WHEN Outline() is called
- THEN outline effect status is returned

#### Scenario: Hidden text
- GIVEN a RunProperties element
- WHEN Vanish() is called
- THEN hidden text status is returned

#### Scenario: Language
- GIVEN a RunProperties element
- WHEN Languages() is called
- THEN language settings for proofing are returned

#### Scenario: Run style
- GIVEN a RunProperties element
- WHEN RunStyle() is called
- THEN the character style ID is returned

### Requirement: NumberingProperties Element
The system SHALL provide NumberingProperties for list formatting.

#### Scenario: Numbering ID
- GIVEN a NumberingProperties element
- WHEN NumberingId() is called
- THEN the numbering definition ID is returned

#### Scenario: Set numbering ID
- GIVEN a NumberingProperties element
- WHEN SetNumberingId(numId) is called
- THEN the paragraph is linked to that numbering

#### Scenario: Numbering level
- GIVEN a NumberingProperties element
- WHEN NumberingLevelReference() is called
- THEN the list level (0-8) is returned

#### Scenario: Set numbering level
- GIVEN a NumberingProperties element
- WHEN SetNumberingLevel(0) is called
- THEN the first level is applied

### Requirement: Indentation Element
The system SHALL provide Indentation for paragraph indents.

#### Scenario: Left indent
- GIVEN an Indentation element
- WHEN Left() is called
- THEN the left indent in twips is returned

#### Scenario: Right indent
- GIVEN an Indentation element
- WHEN Right() is called
- THEN the right indent in twips is returned

#### Scenario: First line indent
- GIVEN an Indentation element
- WHEN FirstLine() is called
- THEN the first line indent in twips is returned

#### Scenario: Hanging indent
- GIVEN an Indentation element
- WHEN Hanging() is called
- THEN the hanging indent in twips is returned

### Requirement: SpacingBetweenLines Element
The system SHALL provide SpacingBetweenLines for paragraph spacing.

#### Scenario: Space before
- GIVEN a SpacingBetweenLines element
- WHEN Before() is called
- THEN space before in twips is returned

#### Scenario: Space after
- GIVEN a SpacingBetweenLines element
- WHEN After() is called
- THEN space after in twips is returned

#### Scenario: Line spacing
- GIVEN a SpacingBetweenLines element
- WHEN Line() is called
- THEN line spacing value is returned

#### Scenario: Line spacing rule
- GIVEN a SpacingBetweenLines element
- WHEN LineRule() is called
- THEN Auto, Exact, or AtLeast is returned

### Requirement: TabStop Element
The system SHALL provide TabStop for custom tab stops.

#### Scenario: Tab position
- GIVEN a TabStop element
- WHEN Position() is called
- THEN tab position in twips is returned

#### Scenario: Tab alignment
- GIVEN a TabStop element
- WHEN Alignment() is called
- THEN Left, Center, Right, or Decimal is returned

#### Scenario: Tab leader
- GIVEN a TabStop element
- WHEN Leader() is called
- THEN None, Dot, Hyphen, Underscore, etc. is returned

### Requirement: Border Elements
The system SHALL provide Border elements for paragraph and text borders.

#### Scenario: Border value
- GIVEN a Border element (top, bottom, left, right)
- WHEN Value() is called
- THEN the border style (Single, Double, etc.) is returned

#### Scenario: Border size
- GIVEN a Border element
- WHEN Size() is called
- THEN the border width in eighths of a point is returned

#### Scenario: Border color
- GIVEN a Border element
- WHEN Color() is called
- THEN the border color is returned

#### Scenario: Border space
- GIVEN a Border element
- WHEN Space() is called
- THEN spacing between border and content is returned

### Requirement: Shading Element
The system SHALL provide Shading for background formatting.

#### Scenario: Fill color
- GIVEN a Shading element
- WHEN Fill() is called
- THEN the background fill color is returned

#### Scenario: Pattern color
- GIVEN a Shading element
- WHEN Color() is called
- THEN the pattern color is returned

#### Scenario: Shading pattern
- GIVEN a Shading element
- WHEN Val() is called
- THEN the shading pattern (Clear, Solid, percent patterns) is returned

