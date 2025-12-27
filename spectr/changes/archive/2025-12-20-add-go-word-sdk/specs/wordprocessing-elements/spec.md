## ADDED Requirements

### Requirement: Document Element
The system SHALL provide a Document element as the root of the main document part.

#### Scenario: Access document body
- GIVEN a Document element
- WHEN Body() is called
- THEN the Body element containing document content is returned

#### Scenario: Create document with body
- GIVEN a new Document element
- WHEN AppendChild(NewBody()) is called
- THEN a Body is added to the document

#### Scenario: Document namespace
- GIVEN a Document element
- WHEN serialized to XML
- THEN the w: namespace prefix and proper namespace declarations are present

### Requirement: Body Element
The system SHALL provide a Body element as the container for document content.

#### Scenario: Access paragraphs
- GIVEN a Body element
- WHEN Elements[Paragraph]() is called
- THEN all paragraph children are returned

#### Scenario: Access tables
- GIVEN a Body element
- WHEN Elements[Table]() is called
- THEN all table children are returned

#### Scenario: Append paragraph
- GIVEN a Body element
- WHEN AppendChild(NewParagraph()) is called
- THEN a paragraph is added to the body

#### Scenario: Section properties
- GIVEN a Body element
- WHEN SectionProperties() is called
- THEN the document's section properties are returned

### Requirement: Paragraph Element
The system SHALL provide a Paragraph element for text blocks.

#### Scenario: Access paragraph properties
- GIVEN a Paragraph element
- WHEN ParagraphProperties() is called
- THEN the ParagraphProperties element is returned

#### Scenario: Access runs
- GIVEN a Paragraph element
- WHEN Elements[Run]() is called
- THEN all Run children are returned

#### Scenario: Get paragraph text
- GIVEN a Paragraph element with runs
- WHEN InnerText() is called
- THEN the concatenated text of all runs is returned

#### Scenario: Append run
- GIVEN a Paragraph element
- WHEN AppendChild(NewRun(text)) is called
- THEN a run with the text is added

#### Scenario: Create paragraph with text
- GIVEN need for simple paragraph
- WHEN NewParagraph(text) is called
- THEN a paragraph with a run containing the text is created

### Requirement: Run Element
The system SHALL provide a Run element for inline content.

#### Scenario: Access run properties
- GIVEN a Run element
- WHEN RunProperties() is called
- THEN the RunProperties element is returned

#### Scenario: Access text
- GIVEN a Run element
- WHEN Text() is called
- THEN the Text element child is returned

#### Scenario: Set run text
- GIVEN a Run element
- WHEN SetText(value) is called
- THEN the text content is set

#### Scenario: Create run with text
- GIVEN need for run
- WHEN NewRun(text) is called
- THEN a run with Text child containing the text is created

### Requirement: Text Element
The system SHALL provide a Text element for text content.

#### Scenario: Get text value
- GIVEN a Text element
- WHEN InnerText() is called
- THEN the text content is returned

#### Scenario: Set text value
- GIVEN a Text element
- WHEN SetText(value) is called
- THEN the text content is updated

#### Scenario: Preserve space attribute
- GIVEN a Text element with whitespace
- WHEN Space attribute is set to "preserve"
- THEN whitespace is preserved in the text

### Requirement: Break Element
The system SHALL provide Break elements for line and page breaks.

#### Scenario: Line break
- GIVEN a Break element with Type=Line or no type
- WHEN rendered
- THEN a line break is represented

#### Scenario: Page break
- GIVEN a Break element with Type=Page
- WHEN rendered
- THEN a page break is represented

#### Scenario: Column break
- GIVEN a Break element with Type=Column
- WHEN rendered
- THEN a column break is represented

#### Scenario: Create line break
- GIVEN need for line break
- WHEN NewBreak() is called
- THEN a line break element is created

#### Scenario: Create page break
- GIVEN need for page break
- WHEN NewBreak(BreakType.Page) is called
- THEN a page break element is created

### Requirement: Tab Element
The system SHALL provide Tab elements for tab characters.

#### Scenario: Create tab
- GIVEN need for tab character
- WHEN NewTab() is called
- THEN a tab element is created

### Requirement: SectionProperties Element
The system SHALL provide SectionProperties for section formatting.

#### Scenario: Page size
- GIVEN a SectionProperties element
- WHEN PageSize() is called
- THEN the page width and height settings are returned

#### Scenario: Page margins
- GIVEN a SectionProperties element
- WHEN PageMargins() is called
- THEN the margin settings (top, bottom, left, right) are returned

#### Scenario: Page orientation
- GIVEN a SectionProperties element
- WHEN PageSize().Orient is accessed
- THEN portrait or landscape is returned

#### Scenario: Columns
- GIVEN a SectionProperties element
- WHEN Columns() is called
- THEN the column layout settings are returned

#### Scenario: Headers and footers references
- GIVEN a SectionProperties element
- WHEN HeaderReference() and FooterReference() are called
- THEN references to header/footer parts are returned

### Requirement: Bookmark Elements
The system SHALL provide elements for bookmarks.

#### Scenario: BookmarkStart
- GIVEN a BookmarkStart element
- WHEN Id() and Name() are accessed
- THEN the bookmark identifier and name are returned

#### Scenario: BookmarkEnd
- GIVEN a BookmarkEnd element
- WHEN Id() is accessed
- THEN the matching bookmark ID is returned

#### Scenario: Create bookmark pair
- GIVEN need to bookmark content
- WHEN NewBookmarkStart(id, name) and NewBookmarkEnd(id) are called
- THEN matching bookmark elements are created

### Requirement: Hyperlink Element
The system SHALL provide Hyperlink elements for links.

#### Scenario: External hyperlink
- GIVEN a Hyperlink element with r:id attribute
- WHEN the relationship is resolved
- THEN the target URL is available

#### Scenario: Internal hyperlink (anchor)
- GIVEN a Hyperlink element with anchor attribute
- WHEN Anchor() is accessed
- THEN the bookmark name is returned

#### Scenario: Create hyperlink
- GIVEN text and URL
- WHEN NewHyperlink(text, url, relationshipId) is called
- THEN a hyperlink element with run is created

### Requirement: Field Elements
The system SHALL provide elements for field codes.

#### Scenario: Simple field
- GIVEN a SimpleField element
- WHEN Instruction() is accessed
- THEN the field code (e.g., "PAGE", "DATE") is returned

#### Scenario: Complex field
- GIVEN FieldChar elements (Begin, Separate, End) with FieldCode
- WHEN processed
- THEN the complex field instruction is understood

#### Scenario: Create page number field
- GIVEN need for page number
- WHEN NewSimpleField("PAGE") is called
- THEN a page number field is created

### Requirement: Drawing Element
The system SHALL provide Drawing elements for graphics.

#### Scenario: Inline drawing
- GIVEN a Drawing element with Inline child
- WHEN Inline() is accessed
- THEN the inline graphic positioning is returned

#### Scenario: Anchor drawing
- GIVEN a Drawing element with Anchor child
- WHEN Anchor() is accessed
- THEN the anchored graphic positioning is returned

### Requirement: Symbol and Special Characters
The system SHALL provide elements for symbols.

#### Scenario: SymbolChar element
- GIVEN a SymbolChar element
- WHEN Font() and Char() are accessed
- THEN the font and character code are returned

#### Scenario: Create symbol
- GIVEN font name and character code
- WHEN NewSymbolChar(font, char) is called
- THEN a symbol element is created

### Requirement: Comment Range Elements
The system SHALL provide elements for comment ranges.

#### Scenario: CommentRangeStart
- GIVEN a CommentRangeStart element
- WHEN Id() is accessed
- THEN the comment ID being started is returned

#### Scenario: CommentRangeEnd
- GIVEN a CommentRangeEnd element
- WHEN Id() is accessed
- THEN the comment ID being ended is returned

#### Scenario: CommentReference
- GIVEN a CommentReference element in a run
- WHEN Id() is accessed
- THEN the referenced comment ID is returned

### Requirement: Structured Document Tags (Content Controls)
The system SHALL provide SDT elements for content controls.

#### Scenario: SdtBlock
- GIVEN an SdtBlock element
- WHEN SdtProperties() and SdtContent() are accessed
- THEN the content control properties and content are returned

#### Scenario: SdtRun
- GIVEN an SdtRun element
- WHEN SdtProperties() and SdtContent() are accessed
- THEN the inline content control is accessible

#### Scenario: SdtCell
- GIVEN an SdtCell element in a table
- WHEN processed
- THEN the table cell content control is handled

### Requirement: Revision Elements
The system SHALL provide elements for tracked changes.

#### Scenario: InsertedRun (w:ins)
- GIVEN an InsertedRun element
- WHEN Author(), Date(), and content are accessed
- THEN the insertion revision information is returned

#### Scenario: DeletedRun (w:del)
- GIVEN a DeletedRun element
- WHEN content is accessed
- THEN the deleted content and revision info are returned

#### Scenario: DeletedText
- GIVEN a DeletedText element
- WHEN InnerText() is called
- THEN the deleted text content is returned
