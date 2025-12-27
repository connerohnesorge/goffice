# Wordprocessing Numbering Specification

## Requirements

### Requirement: Numbering Root Element
The system SHALL provide a Numbering element as the root of the numbering part.

#### Scenario: Access abstract numbering definitions
- GIVEN a Numbering element
- WHEN Elements[AbstractNum]() is called
- THEN all abstract numbering definitions are returned

#### Scenario: Access numbering instances
- GIVEN a Numbering element
- WHEN Elements[NumberingInstance]() is called
- THEN all numbering instances are returned

#### Scenario: Get abstract num by ID
- GIVEN a Numbering element
- WHEN GetAbstractNum(abstractNumId) is called
- THEN the AbstractNum with that ID is returned

#### Scenario: Get numbering instance by ID
- GIVEN a Numbering element
- WHEN GetNumInstance(numId) is called
- THEN the NumberingInstance with that ID is returned

#### Scenario: Add abstract numbering
- GIVEN a Numbering element
- WHEN AddAbstractNum(abstractNum) is called
- THEN the definition is added with unique ID

#### Scenario: Add numbering instance
- GIVEN a Numbering element
- WHEN AddNumInstance(numInstance) is called
- THEN the instance is added with unique ID

### Requirement: AbstractNum Element
The system SHALL provide AbstractNum for abstract numbering definitions.

#### Scenario: Abstract num ID
- GIVEN an AbstractNum element
- WHEN AbstractNumId() is called
- THEN the unique identifier is returned

#### Scenario: Access levels
- GIVEN an AbstractNum element
- WHEN Elements[Level]() is called
- THEN all Level definitions (0-8) are returned

#### Scenario: Get level by index
- GIVEN an AbstractNum element
- WHEN GetLevel(levelIndex) is called
- THEN the Level at that index is returned

#### Scenario: Multilevel type
- GIVEN an AbstractNum element
- WHEN MultiLevelType() is called
- THEN SingleLevel, Multilevel, or HybridMultilevel is returned

#### Scenario: Numbering style link
- GIVEN an AbstractNum element
- WHEN NumberingStyleLink() is called
- THEN the linked numbering style ID is returned

#### Scenario: Style link
- GIVEN an AbstractNum element
- WHEN StyleLink() is called
- THEN the linked style ID is returned

### Requirement: Level Element
The system SHALL provide Level for individual list level definitions.

#### Scenario: Level index
- GIVEN a Level element
- WHEN LevelIndex() is called
- THEN the level number (0-8) is returned

#### Scenario: Start value
- GIVEN a Level element
- WHEN Start() is called
- THEN the starting number is returned

#### Scenario: Set start value
- GIVEN a Level element
- WHEN Start().Val.SetValue(1) is called
- THEN the level starts at 1

#### Scenario: Number format
- GIVEN a Level element
- WHEN NumberFormat() is called
- THEN Decimal, LowerLetter, UpperLetter, LowerRoman, UpperRoman, Bullet, etc. is returned

#### Scenario: Set number format
- GIVEN a Level element
- WHEN SetNumberFormat(NumberFormatValues.Decimal) is called
- THEN decimal numbering is used

#### Scenario: Level text
- GIVEN a Level element
- WHEN LevelText() is called
- THEN the format string (e.g., "%1.", "%1.%2") is returned

#### Scenario: Set level text
- GIVEN a Level element
- WHEN LevelText().Val.SetValue("%1)") is called
- THEN "1)" format is used

#### Scenario: Justification
- GIVEN a Level element
- WHEN LevelJustification() is called
- THEN Left, Center, or Right alignment is returned

#### Scenario: Paragraph properties
- GIVEN a Level element
- WHEN PreviousParagraphProperties() is called
- THEN indentation and other paragraph settings are returned

#### Scenario: Run properties
- GIVEN a Level element
- WHEN NumberingSymbolRunProperties() is called
- THEN formatting for the number/bullet is returned

#### Scenario: Level restart
- GIVEN a Level element
- WHEN LevelRestart() is called
- THEN the level at which this level restarts is returned

#### Scenario: Suffix
- GIVEN a Level element
- WHEN LevelSuffix() is called
- THEN Tab, Space, or Nothing after number is returned

#### Scenario: Is legal numbering
- GIVEN a Level element
- WHEN IsLegalNumberingStyle() is called
- THEN whether legal numbering format is used is returned

### Requirement: NumberingInstance Element
The system SHALL provide NumberingInstance for concrete numbering usage.

#### Scenario: Numbering ID
- GIVEN a NumberingInstance element
- WHEN NumId() is called
- THEN the unique ID referenced by paragraphs is returned

#### Scenario: Abstract num reference
- GIVEN a NumberingInstance element
- WHEN AbstractNumId() is called
- THEN the referenced abstract numbering ID is returned

#### Scenario: Level overrides
- GIVEN a NumberingInstance element
- WHEN Elements[LevelOverride]() is called
- THEN level-specific overrides are returned

### Requirement: LevelOverride Element
The system SHALL provide LevelOverride for instance-specific level changes.

#### Scenario: Override level index
- GIVEN a LevelOverride element
- WHEN LevelIndex() is called
- THEN the level being overridden is returned

#### Scenario: Start override
- GIVEN a LevelOverride element
- WHEN StartOverride() is called
- THEN the overridden start value is returned

#### Scenario: Level override
- GIVEN a LevelOverride element
- WHEN Level() is called
- THEN the complete level override definition is returned

### Requirement: Bullet List Creation
The system SHALL provide convenience methods for bullet lists.

#### Scenario: Create bullet list definition
- GIVEN bullet character and font
- WHEN NewBulletList(char, font) is called
- THEN an AbstractNum with bullet formatting is created

#### Scenario: Standard bullet list
- GIVEN need for standard bullets
- WHEN NewStandardBulletList() is called
- THEN a typical bullet list with indentation is created

#### Scenario: Multi-level bullets
- GIVEN need for nested bullets
- WHEN NewMultiLevelBulletList() is called
- THEN a multi-level bullet definition is created

### Requirement: Numbered List Creation
The system SHALL provide convenience methods for numbered lists.

#### Scenario: Create decimal list
- GIVEN need for numbered list
- WHEN NewDecimalList() is called
- THEN an AbstractNum with decimal numbering is created

#### Scenario: Create letter list
- GIVEN need for lettered list
- WHEN NewAlphabeticList(lowercase) is called
- THEN an AbstractNum with a, b, c or A, B, C is created

#### Scenario: Create roman numeral list
- GIVEN need for roman numerals
- WHEN NewRomanNumeralList(lowercase) is called
- THEN an AbstractNum with i, ii, iii or I, II, III is created

#### Scenario: Create outline numbering
- GIVEN need for outline (1.1.1 style)
- WHEN NewOutlineList() is called
- THEN a multi-level outline numbering is created

### Requirement: Apply Numbering to Paragraphs
The system SHALL support applying numbering to paragraphs.

#### Scenario: Apply numbering to paragraph
- GIVEN a Paragraph and numbering definition
- WHEN ApplyNumbering(numId, level) is called
- THEN the paragraph uses that numbering at specified level

#### Scenario: Remove numbering
- GIVEN a numbered Paragraph
- WHEN RemoveNumbering() is called
- THEN the paragraph's numbering is removed

#### Scenario: Change numbering level
- GIVEN a numbered Paragraph
- WHEN SetNumberingLevel(newLevel) is called
- THEN the paragraph's list level is changed

### Requirement: Numbering ID Generation
The system SHALL manage unique IDs for numbering definitions.

#### Scenario: Generate abstract num ID
- GIVEN a Numbering part
- WHEN NewAbstractNum() is called
- THEN a unique abstractNumId is assigned

#### Scenario: Generate num ID
- GIVEN a Numbering part
- WHEN NewNumInstance(abstractNumId) is called
- THEN a unique numId is assigned

#### Scenario: ID uniqueness
- GIVEN existing numbering definitions
- WHEN new definitions are added
- THEN IDs do not conflict with existing ones

