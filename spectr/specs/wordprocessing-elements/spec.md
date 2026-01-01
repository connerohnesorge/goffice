# Wordprocessing Elements Specification

## Requirements

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
The system SHALL provide elements for field codes AND a high-level API for form field manipulation.

#### Scenario: Simple field (unchanged)
- GIVEN a SimpleField element
- WHEN Instruction() is accessed
- THEN the field code (e.g., "PAGE", "DATE") is returned

#### Scenario: Complex field (unchanged)
- GIVEN FieldChar elements (Begin, Separate, End) with FieldCode
- WHEN processed
- THEN the complex field instruction is understood

#### Scenario: Create page number field (unchanged)
- GIVEN need for page number
- WHEN NewSimpleField("PAGE") is called
- THEN a page number field is created

#### Scenario: Form field API access
- GIVEN a SimpleField with FormFieldData child
- WHEN wrapped in FormField object
- THEN high-level get/set methods are available

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
The system SHALL provide elements for tracked changes with high-level wrapper API for document-level operations.

**Note**: This modifies the existing "Revision Elements" requirement in `spectr/specs/wordprocessing-elements/spec.md` to add the high-level wrapper context.

#### Scenario: InsertedRun (w:ins)
- GIVEN an InsertedRun element
- WHEN Author(), Date(), and content are accessed
- THEN the insertion revision information is returned
- AND the element can be wrapped in a high-level Revision object
- AND the Revision wrapper provides Accept() and Reject() methods

#### Scenario: DeletedRun (w:del)
- GIVEN a DeletedRun element
- WHEN content is accessed
- THEN the deleted content and revision info are returned
- AND the element can be wrapped in a high-level Revision object
- AND the Revision wrapper provides Accept() and Reject() methods

#### Scenario: DeletedText
- GIVEN a DeletedText element
- WHEN InnerText() is called
- THEN the deleted text content is returned
- AND the text can be restored to normal Text during reject operations

#### Scenario: MoveFromRun
- GIVEN a MoveFromRun element with id=10
- WHEN wrapped in a Revision object
- THEN the revision metadata (id, author, date) is accessible
- AND the Revision can be paired with matching MoveToRun id=10
- AND Accept/Reject operations coordinate with the move pair

#### Scenario: MoveToRun
- GIVEN a MoveToRun element with id=10
- WHEN wrapped in a Revision object
- THEN the revision metadata (id, author, date) is accessible
- AND the Revision can be paired with matching MoveFromRun id=10
- AND Accept/Reject operations coordinate with the move pair

#### Scenario: RunPropertiesChange
- GIVEN a RunPropertiesChange element
- WHEN wrapped in a Revision object
- THEN the previous properties and new properties are accessible
- AND Accept applies new properties and removes wrapper
- AND Reject restores previous properties and removes wrapper

#### Scenario: ParagraphPropertiesChange
- GIVEN a ParagraphPropertiesChange element
- WHEN wrapped in a Revision object
- THEN the previous properties and new properties are accessible
- AND Accept applies new properties and removes wrapper
- AND Reject restores previous properties and removes wrapper

### Requirement: Form Field Wrapper
The system SHALL provide a FormField wrapper for legacy form fields.

#### Scenario: FormField wraps SimpleField
- GIVEN a SimpleField with FormFieldData
- WHEN FormField wrapper is created
- THEN the wrapper provides access to field properties

#### Scenario: Field type detection
- GIVEN a SimpleField with instr="FORMTEXT"
- WHEN FormField.Type() is called
- THEN FormFieldTypeText is returned

#### Scenario: Field name access
- GIVEN a FormField with name "CustomerName"
- WHEN FormField.Name() is called
- THEN "CustomerName" is returned

#### Scenario: Field name modification
- GIVEN a FormField
- WHEN FormField.SetName("NewName") is called
- THEN the ffData/name element is updated

### Requirement: Text Field Manipulation
The system SHALL provide methods for text form field manipulation.

#### Scenario: Get text field value
- GIVEN a text FormField with value "John Doe"
- WHEN GetTextValue() is called
- THEN "John Doe" is returned

#### Scenario: Set text field value
- GIVEN a text FormField
- WHEN SetTextValue("Jane Smith") is called
- THEN both ffData/textInput/default AND run/text are updated to "Jane Smith"

#### Scenario: Set maximum length
- GIVEN a text FormField
- WHEN SetMaxLength(50) is called
- THEN ffData/textInput/maxLength is set to 50

#### Scenario: Get maximum length
- GIVEN a text FormField with maxLength=50
- WHEN GetMaxLength() is called
- THEN 50 is returned

#### Scenario: Set default text value
- GIVEN a text FormField
- WHEN SetDefaultTextValue("Default Name") is called
- THEN ffData/textInput/default is set to "Default Name"

#### Scenario: Text value synchronization
- GIVEN a text FormField
- WHEN SetTextValue("New Value") is called
- THEN the run text element matches "New Value" exactly

### Requirement: CheckBox Field Manipulation
The system SHALL provide methods for checkbox form field manipulation.

#### Scenario: Check checkbox state
- GIVEN a checkbox FormField that is checked
- WHEN IsChecked() is called
- THEN true is returned

#### Scenario: Toggle checkbox to checked
- GIVEN an unchecked checkbox FormField
- WHEN SetChecked(true) is called
- THEN ffData/checkBox/checked is set to 1 AND symbol character is updated to F052

#### Scenario: Toggle checkbox to unchecked
- GIVEN a checked checkbox FormField
- WHEN SetChecked(false) is called
- THEN ffData/checkBox/checked is set to 0 AND symbol character is updated to F06F

#### Scenario: Set checkbox size
- GIVEN a checkbox FormField
- WHEN SetCheckBoxSize(20) is called
- THEN ffData/checkBox/size is set to 20 (10 points)

#### Scenario: Get checkbox size
- GIVEN a checkbox FormField with size=20
- WHEN GetCheckBoxSize() is called
- THEN 20 is returned

#### Scenario: Auto checkbox size
- GIVEN a checkbox FormField
- WHEN SetAutoCheckBoxSize(true) is called
- THEN ffData/checkBox/sizeAuto is set to 1

#### Scenario: Checkbox symbol synchronization
- GIVEN a checkbox FormField
- WHEN SetChecked(true) is called
- THEN the run/sym element has char="F052" (checked box symbol)

#### Scenario: Checkbox symbol must be set via SetChecked
- GIVEN a checkbox FormField
- WHEN SetChecked() is used to change state
- THEN both ffData/checkBox/checked AND run/sym/char are synchronized
- NOTE: Symbols MUST only be set via SetChecked() to ensure proper synchronization
- NOTE: Manual manipulation of symbol elements may create invalid state

### Requirement: DropDown Field Manipulation
The system SHALL provide methods for dropdown form field manipulation.

#### Scenario: Get dropdown items
- GIVEN a dropdown FormField with items ["USA", "Canada", "UK"]
- WHEN GetDropDownItems() is called
- THEN ["USA", "Canada", "UK"] is returned

#### Scenario: Set dropdown items
- GIVEN a dropdown FormField
- WHEN SetDropDownItems(["Option A", "Option B", "Option C"]) is called
- THEN ffData/ddList contains three listEntry elements with those values

#### Scenario: Get selected index
- GIVEN a dropdown FormField with selectedIndex=1
- WHEN GetSelectedIndex() is called
- THEN 1 is returned

#### Scenario: Set selected index
- GIVEN a dropdown FormField with items ["A", "B", "C"]
- WHEN SetSelectedIndex(2) is called
- THEN ffData/ddList/result is 2 AND run/text is "C"

#### Scenario: Get selected value
- GIVEN a dropdown FormField with items ["USA", "Canada"] and selectedIndex=0
- WHEN GetSelectedValue() is called
- THEN "USA" is returned

#### Scenario: DropDown value synchronization
- GIVEN a dropdown FormField with items ["X", "Y", "Z"]
- WHEN SetSelectedIndex(1) is called
- THEN the run text element contains "Y" exactly

#### Scenario: Invalid selected index handling
- GIVEN a dropdown FormField with 3 items
- WHEN SetSelectedIndex(5) is called (out of range)
- THEN the operation is ignored or error is returned

#### Scenario: Get default dropdown index
- GIVEN a dropdown FormField with default=2
- WHEN GetDefaultDropDownIndex() is called
- THEN 2 is returned

#### Scenario: Set default dropdown index
- GIVEN a dropdown FormField with items ["A", "B", "C"]
- WHEN SetDefaultDropDownIndex(1) is called
- THEN ffData/ddList/default is set to 1

#### Scenario: DropDown max items validation
- GIVEN a Paragraph
- WHEN SetDropDownItems() is called with 26 items
- THEN the operation panics with "cannot have more than 25 items"

#### Scenario: DropDown insert max items validation
- GIVEN a Paragraph
- WHEN InsertDropDown() is called with 26 items
- THEN the operation panics with schema constraint error

### Requirement: Form Field Insertion in Paragraphs
The system SHALL provide methods to insert form fields into paragraphs.

#### Scenario: Insert text field
- GIVEN a Paragraph
- WHEN InsertTextField("CustomerName", "Enter name") is called
- THEN a SimpleField with FORMTEXT instruction, FormFieldData, and run is added

#### Scenario: Insert checkbox
- GIVEN a Paragraph
- WHEN InsertCheckBox("AgreeTerms", false) is called
- THEN a SimpleField with FORMCHECKBOX instruction, FormFieldData with checkBox, and symbol run is added

#### Scenario: Insert dropdown
- GIVEN a Paragraph
- WHEN InsertDropDown("Country", ["USA", "UK"], 0) is called
- THEN a SimpleField with FORMDROPDOWN instruction, FormFieldData with ddList, and run with selected text is added

#### Scenario: Multiple fields in one paragraph
- GIVEN a Paragraph
- WHEN InsertTextField(), InsertCheckBox(), and InsertDropDown() are called sequentially
- THEN all three SimpleField elements are added to the paragraph

#### Scenario: Inserted field is immediately accessible
- GIVEN a Paragraph
- WHEN field := InsertTextField("Test", "Value") is called
- THEN field.GetTextValue() returns "Value" immediately

### Requirement: Form Field Removal
The system SHALL support removing form fields from paragraphs.

#### Scenario: Remove field from paragraph
- GIVEN a FormField in a Paragraph
- WHEN Remove() is called
- THEN the SimpleField element is removed from the paragraph

#### Scenario: Remove and re-iterate
- GIVEN a document with 3 form fields
- WHEN one field is removed
- THEN GetFormFields() yields only 2 fields

### Requirement: Form Field Enabled State
The system SHALL support enabling and disabling form fields.

#### Scenario: Check if field is enabled
- GIVEN a FormField with enabled element present
- WHEN Enabled() is called
- THEN true is returned

#### Scenario: Disable field
- GIVEN a FormField
- WHEN SetEnabled(false) is called
- THEN the ffData/enabled element is removed

#### Scenario: Enable field
- GIVEN a FormField
- WHEN SetEnabled(true) is called
- THEN the ffData/enabled element is added (if not present)

#### Scenario: Disabled field persists on save
- GIVEN a disabled FormField
- WHEN the document is saved and reopened
- THEN the field is still disabled

### Requirement: Form Field XML Structure Compliance
The system SHALL generate form field XML structures compatible with Microsoft Word.

#### Scenario: Text field XML structure
- GIVEN a created text FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMTEXT"]/ffData/textInput/default

#### Scenario: CheckBox field XML structure
- GIVEN a created checkbox FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMCHECKBOX"]/ffData/checkBox/checked

#### Scenario: DropDown field XML structure
- GIVEN a created dropdown FormField
- WHEN serialized to XML
- THEN the structure matches: fldSimple[@instr="FORMDROPDOWN"]/ffData/ddList/listEntry

#### Scenario: Field name in ffData
- GIVEN any FormField with name "TestField"
- WHEN serialized to XML
- THEN ffData contains: name[@val="TestField"]

#### Scenario: Enabled attribute
- GIVEN an enabled FormField
- WHEN serialized to XML
- THEN ffData contains: enabled element (empty)

### Requirement: Form Field Roundtrip Compatibility
The system SHALL maintain form field integrity across save/load cycles.

#### Scenario: Text field value roundtrip
- GIVEN a text FormField with value "Test Value"
- WHEN saved, closed, and reopened
- THEN GetTextValue() returns "Test Value"

#### Scenario: CheckBox state roundtrip
- GIVEN a checked checkbox FormField
- WHEN saved, closed, and reopened
- THEN IsChecked() returns true

#### Scenario: DropDown selection roundtrip
- GIVEN a dropdown FormField with selectedIndex=2
- WHEN saved, closed, and reopened
- THEN GetSelectedIndex() returns 2

#### Scenario: Field name roundtrip
- GIVEN a FormField with name "UniqueField"
- WHEN saved, closed, and reopened
- THEN Name() returns "UniqueField"

#### Scenario: Multiple fields roundtrip
- GIVEN a document with text, checkbox, and dropdown fields
- WHEN saved, closed, and reopened
- THEN all three fields exist with correct types and values

### Requirement: Revision Type Enumeration
The system SHALL provide type identification for all revision elements.

#### Scenario: Identify insertion type
- GIVEN an InsertedRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeInsert

#### Scenario: Identify deletion type
- GIVEN a DeletedRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeDelete

#### Scenario: Identify move-from type
- GIVEN a MoveFromRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeMoveFrom

#### Scenario: Identify move-to type
- GIVEN a MoveToRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeMoveTo

#### Scenario: Identify format change type
- GIVEN a RunPropertiesChange or ParagraphPropertiesChange element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeFormatChange

### Requirement: Revision Content Extraction
The system SHALL provide unified content access across revision types.

#### Scenario: Extract insertion content
- GIVEN an InsertedRun containing runs with text "new content"
- WHEN wrapped as Revision and Content() is called
- THEN "new content" is returned

#### Scenario: Extract deletion content
- GIVEN a DeletedRun containing DeletedText "removed content"
- WHEN wrapped as Revision and Content() is called
- THEN "removed content" is returned

#### Scenario: Extract move content
- GIVEN a MoveFromRun containing runs with text "moved text"
- WHEN wrapped as Revision and Content() is called
- THEN "moved text" is returned

#### Scenario: Extract format change description
- GIVEN a RunPropertiesChange element
- WHEN wrapped as Revision and Content() is called
- THEN a description of the property change is returned

### Requirement: Revision Element Traversal
The system SHALL support finding all revision elements in a document tree.

#### Scenario: Find revisions in paragraph
- GIVEN a Paragraph containing 2 InsertedRun and 1 DeletedRun
- WHEN the paragraph is traversed for revisions
- THEN all 3 revision elements are found

#### Scenario: Find revisions in nested elements
- GIVEN a Table containing cells with InsertedRun elements
- WHEN the table is traversed for revisions
- THEN all InsertedRun elements in all cells are found

#### Scenario: Find revisions in headers
- GIVEN a HeaderPart containing tracked changes
- WHEN the header is traversed for revisions
- THEN all revision elements in the header are found

#### Scenario: Find revisions in footers
- GIVEN a FooterPart containing tracked changes
- WHEN the footer is traversed for revisions
- THEN all revision elements in the footer are found

#### Scenario: Find revisions in comments
- GIVEN a CommentsPart with tracked changes in comment text
- WHEN the comments are traversed for revisions
- THEN all revision elements in comments are found

### Requirement: Revision Iterator Pattern
The system SHALL use Go 1.23+ iter.Seq for memory-efficient revision enumeration.

#### Scenario: Iterate revisions without loading all
- GIVEN a document with 10,000 tracked changes
- WHEN GetRevisions() returns iter.Seq[*Revision]
- THEN memory usage scales with iteration, not total count
- AND early break stops further enumeration

#### Scenario: Iterate and filter
- GIVEN an iterator over revisions
- WHEN filtering by author during iteration
- THEN only matching revisions are processed
- AND non-matching revisions are skipped efficiently

#### Scenario: Empty revision iterator
- GIVEN a document with no tracked changes
- WHEN GetRevisions() is called
- THEN the iterator completes immediately
- AND no allocations occur for empty results

### Requirement: Office 2010 Extension Elements
The system SHALL provide element types for Word 2010 extensions.

#### Scenario: ContentControl element
- GIVEN Office 2010 w14:contentControl element
- WHEN parsed from XML
- THEN ContentControl element instance is created in w14 namespace

#### Scenario: CustomXmlConflictInsertionRangeEnd element
- GIVEN Office 2010 w14:customXmlConflictInsRangeEnd element
- WHEN accessed via API
- THEN strongly-typed CustomXmlConflictInsertionRangeEnd element is available

#### Scenario: DocId element
- GIVEN Office 2010 w14:docId element
- WHEN present in document
- THEN DocId element with Val attribute is accessible

#### Scenario: Drawing Canvas elements
- GIVEN Office 2010 drawing canvas elements (wpc namespace)
- WHEN parsed
- THEN WordprocessingCanvas, WordprocessingGroup, WordprocessingShape elements available

### Requirement: Office 2012-2013 Extension Elements
The system SHALL provide element types for Word 2012-2013 extensions.

#### Scenario: WebExtension elements
- GIVEN Office 2013 webextension elements (we namespace)
- WHEN parsed
- THEN WebExtension, WebExtensionReference, WebExtensionProperty elements available

#### Scenario: Chart Style elements
- GIVEN Office 2013 chart style elements (w15 namespace)
- WHEN parsed
- THEN ChartStyle, ChartColor elements available

### Requirement: Office 2015-2016 Extension Elements
The system SHALL provide element types for Word 2015-2016 extensions.

#### Scenario: Appearance elements
- GIVEN Office 2015 appearance elements (w15 symex namespace)
- WHEN parsed
- THEN SdtAppearance element with Val attribute available

#### Scenario: Comments Extended elements
- GIVEN Office 2016 comments elements (w16cex namespace)
- WHEN parsed
- THEN CommentExtensible element available

#### Scenario: Comments ID elements
- GIVEN Office 2016 comment ID elements (w16cid namespace)
- WHEN parsed
- THEN CommentId, ParagraphId elements available

### Requirement: Office 2018-2020 Extension Elements
The system SHALL provide element types for Word 2018-2020 extensions.

#### Scenario: 2018 extension elements
- GIVEN Office 2018 elements (w18 namespace)
- WHEN parsed
- THEN Word 2018 extension elements available

#### Scenario: SDT Data Hash elements
- GIVEN Office 2020 SDT data hash elements (w20sdtdh namespace)
- WHEN parsed
- THEN SdtDataHash element available

### Requirement: Office 2023-2024 Extension Elements
The system SHALL provide element types for Word 2023-2024 extensions.

#### Scenario: Word 16 Document Undo elements
- GIVEN Office 2023 elements (w16du namespace)
- WHEN parsed
- THEN Word 16 document undo elements available

#### Scenario: SDT Format Lock elements
- GIVEN Office 2024 SDT format lock elements (w24sdtfl namespace)
- WHEN parsed
- THEN SdtFormatLock element with formatting lock properties available

### Requirement: Extension Element Attributes
The system SHALL provide typed attributes for extension elements.

#### Scenario: Extension string attributes
- GIVEN an extension element with string attribute
- WHEN attribute is accessed
- THEN StringValue wrapper is returned

#### Scenario: Extension enum attributes
- GIVEN an extension element with enumeration attribute
- WHEN attribute is accessed
- THEN enum value from extension schema is returned

#### Scenario: Extension boolean attributes
- GIVEN an extension element with boolean attribute
- WHEN attribute is accessed
- THEN BoolValue wrapper is returned

### Requirement: Extension Element Namespace Handling
The system SHALL correctly handle namespaces for extension elements.

#### Scenario: Element qualified name
- GIVEN a ContentControl element (w14 namespace)
- WHEN QName() is called
- THEN QualifiedName with namespace "http://schemas.microsoft.com/office/word/2010/wordml" is returned

#### Scenario: Element local name
- GIVEN a ContentControl element
- WHEN LocalName() is called
- THEN "contentControl" is returned

#### Scenario: Namespace prefix in XML
- GIVEN a ContentControl element
- WHEN written to XML
- THEN element is serialized as `<w14:contentControl>`

### Requirement: Extension Element Version Metadata
The system SHALL provide version information for extension elements.

#### Scenario: Office 2010 element version
- GIVEN a ContentControl element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Office 2024 element version
- GIVEN an SdtFormatLock element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2024 is returned

### Requirement: Element Generation from Extension Schemas
The system SHALL generate element types from extension schemas in addition to main schemas.

#### Scenario: Generate from main schema
- GIVEN wordprocessingml main.json schema
- WHEN generator runs
- THEN core WordprocessingML elements are generated

#### Scenario: Generate from extension schemas
- GIVEN Word 2010-2024 extension schemas
- WHEN generator runs
- THEN extension element types are generated with correct namespaces and version metadata

#### Scenario: Total element count
- GIVEN all schemas loaded
- WHEN generation completes
- THEN 680-700+ element types exist (baseline ~619 + 60-80 extensions)

#### Scenario: Generated element namespaces
- GIVEN generated extension elements
- WHEN inspecting code
- THEN elements use correct namespace constants (NamespaceWord2010, NamespaceWord2012, etc.)

### Requirement: Extension Element Type Registration
The system SHALL register all extension element types for XML parsing.

#### Scenario: Register extension elements
- GIVEN generated extension element types
- WHEN element registry is initialized
- THEN all extension elements are registered with their qualified names

#### Scenario: Parse registered extension element
- GIVEN XML with w14:contentControl element
- WHEN parsing document
- THEN ContentControl instance is created (not UnknownElement)

#### Scenario: Parse unregistered element
- GIVEN XML with unknown namespace element
- WHEN parsing document
- THEN UnknownElement instance is created preserving raw XML
