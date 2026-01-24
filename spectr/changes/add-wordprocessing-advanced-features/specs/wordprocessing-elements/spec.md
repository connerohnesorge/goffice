## ADDED Requirements

### Requirement: Advanced Text Formatting
The system SHALL provide comprehensive text formatting capabilities beyond basic bold/italic/underline.

#### Scenario: Apply theme color to text
- GIVEN a text run element
- WHEN SetThemeColor(ThemeColorScheme.Accent1, shade=25000) is called
- THEN the text renders with the specified theme color
- AND the theme color persists in roundtrip save/load

#### Scenario: Set character spacing
- GIVEN a text run element
- WHEN SetCharacterSpacing(100) is called (in 1/20th of a point)
- THEN spacing attribute is set to 100
- AND visual spacing increases accordingly

#### Scenario: Apply font scaling
- GIVEN a text run element
- WHEN SetFontScale(80) is called (percentage)
- THEN the font width scales to 80% of normal
- AND height remains unchanged

#### Scenario: Apply text effects (all caps, small caps)
- GIVEN a text run element
- WHEN SetAllCaps(true) or SetSmallCaps(true) is called
- THEN the appropriate text transformation is applied
- AND persistence is maintained in roundtrip

#### Scenario: Set superscript and subscript
- GIVEN a text run element
- WHEN SetSuperscript() or SetSubscript() is called
- THEN position attribute is set accordingly
- AND text renders at correct vertical position

#### Scenario: Apply text fill and outline
- GIVEN a text run element
- WHEN SetSolidFill(color) is called
- THEN text fill is set to solid color
- AND outline properties can be applied separately

### Requirement: Table Cell Merging
The system SHALL support merging table cells horizontally and vertically.

#### Scenario: Merge cells horizontally
- GIVEN a table row with 4 cells
- WHEN cells[0].MergeHorizontally(cells[1]) and cells[2].MergeHorizontally(cells[3]) are called
- THEN visual table shows 2 merged cells per row
- AND gridSpan is properly set

#### Scenario: Merge cells vertically
- GIVEN a table with 3 rows and 2 columns
- WHEN row[0].cells[0].MergeVertically(row[1].cells[0]) and row[1].cells[0].MergeVertically(row[2].cells[0]) are called
- THEN cells appear merged vertically across 3 rows
- AND vMerge attributes are correctly set

#### Scenario: Check merge status
- GIVEN a table cell that is part of a merge
- WHEN cell.IsMerged() is called
- THEN true is returned if cell is in merge, false otherwise
- AND cell.MergeStart() returns the first cell in merge range

#### Scenario: Unmerge cells
- GIVEN previously merged cells
- WHEN cell.Unmerge() is called
- THEN cells become independent
- AND merge attributes are removed

### Requirement: Table Styling
The system SHALL support applying and managing table styles.

#### Scenario: Apply table style
- GIVEN a table element
- WHEN table.SetStyle("TableGrid") is called
- THEN the style reference is set in table properties
- AND formatting from the style is applied

#### Scenario: Apply conditional formatting
- GIVEN a table with SetConditionalFormatting()
- WHEN firstRow=true, lastRow=true, firstCol=true, lastCol=true are set
- THEN table formatting respects these row/column conditions
- AND style rules apply conditionally

### Requirement: Advanced Paragraph Formatting
The system SHALL support outline levels, orphan/widow control, and advanced spacing.

#### Scenario: Set paragraph outline level
- GIVEN a paragraph element
- WHEN paragraph.SetOutlineLevel(1) is called (0-8)
- THEN ilvl is set to 1
- AND paragraph is treated as heading level 2

#### Scenario: Enable orphan/widow control
- GIVEN a paragraph element
- WHEN paragraph.SetOrphanWidowControl(true) is called
- THEN orphanControl element is set
- AND paragraph keeps at least 2 lines together

#### Scenario: Set paragraph shading
- GIVEN a paragraph element
- WHEN paragraph.SetShading(ShadingPattern.Clear, color.White, color.Blue) is called
- THEN background shading is applied
- AND pattern type is respected

#### Scenario: Set paragraph borders
- GIVEN a paragraph element
- WHEN paragraph.SetBorder(BorderType.Top, BorderStyle.Single, 12, color.Black) is called
- THEN border is drawn on specified side
- AND color and width are applied

#### Scenario: Set line spacing
- GIVEN a paragraph element
- WHEN paragraph.SetLineSpacing(LineSpacingType.AtLeast, 360) is called (in twips)
- THEN spacing between lines is set
- AND spacing type (single, 1.5, double, at least, exactly) is applied

#### Scenario: Set paragraph spacing (before/after)
- GIVEN a paragraph element
- WHEN paragraph.SetSpacingBefore(240) and SetSpacingAfter(120) are called (in twips)
- THEN spacing attributes are set
- AND spacing is applied before and after paragraph

#### Scenario: Set keep with next and keep lines together
- GIVEN a paragraph element
- WHEN paragraph.SetKeepWithNext(true) is called
- THEN keepNext attribute is set
- AND paragraph stays with next paragraph on page break

### Requirement: Section Management
The system SHALL support creating, accessing, and modifying sections with distinct properties.

#### Scenario: Create new section
- GIVEN a document with paragraphs
- WHEN document.InsertSection(atParagraph) is called
- THEN a new section is created
- AND section properties are initialized with defaults

#### Scenario: Set section type
- GIVEN a section element
- WHEN section.SetSectionType(SectionType.NextPage) is called
- THEN page break behavior is applied
- AND continuous, odd/even page options are supported

#### Scenario: Set different header per section
- GIVEN a document with multiple sections
- WHEN section.SetHeaderReference(headerPartReference) is called
- THEN header is unique to this section
- AND other sections maintain their headers

#### Scenario: Set different footer per section
- GIVEN a document with multiple sections
- WHEN section.SetFooterReference(footerPartReference) is called
- THEN footer is unique to this section
- AND footer content differs from other sections

#### Scenario: Set section page dimensions
- GIVEN a section element
- WHEN section.SetPageDimensions(width=12240, height=15840, marginTop=1440) are called (in twips)
- THEN page size and margins are set for this section only
- AND other sections retain their dimensions

#### Scenario: Set section column layout
- GIVEN a section element
- WHEN section.SetColumns(ColumnLayout.TwoColumn, sep=true, space=720) is called
- THEN text columns are configured
- AND separator line is drawn if enabled

### Requirement: Comments and Annotations
The system SHALL support creating, reading, and managing document comments.

#### Scenario: Add comment to text range
- GIVEN a document with selected text range
- WHEN document.AddComment("Author Name", "Comment text", date, initials) is called
- THEN comment is created and attached to text
- AND comment markers are inserted in document

#### Scenario: Get all comments
- GIVEN a document with multiple comments
- WHEN document.GetComments() is called
- THEN all comments are returned with metadata
- AND comments include author, date, text

#### Scenario: Reply to comment
- GIVEN an existing comment
- WHEN comment.Reply("Reply author", "Reply text") is called
- THEN a reply comment is created
- AND reply is threaded under original comment

#### Scenario: Access comment metadata
- GIVEN a comment element
- WHEN comment.Author(), comment.Date(), comment.Initials() are called
- THEN metadata fields are returned
- AND date is parsed to time.Time

#### Scenario: Delete comment
- GIVEN a comment in document
- WHEN comment.Delete() is called
- THEN comment is removed
- AND comment markers are removed from document

### Requirement: Bookmarks and Cross-References
The system SHALL support creating and referencing bookmarks within documents.

#### Scenario: Create bookmark
- GIVEN a text range in document
- WHEN document.AddBookmark("BookmarkName", startPara, endPara) is called
- THEN bookmark elements are inserted
- AND bookmark is accessible by name

#### Scenario: Get bookmark by name
- GIVEN a document with bookmarks
- WHEN document.GetBookmark("BookmarkName") is called
- THEN bookmark reference is returned
- AND bookmark range can be accessed

#### Scenario: Create cross-reference field
- GIVEN a paragraph element
- WHEN paragraph.AddCrossReferenceField(CrossRefType.Bookmark, "BookmarkName") is called
- THEN field is inserted
- AND field code references the bookmark

#### Scenario: Create reference fields (REF, STYLEREF, PAGEREF)
- GIVEN a paragraph element
- WHEN paragraph.AddRefField("BookmarkName") is called
- THEN REF field is inserted
- AND field updates to show bookmark content or page number

### Requirement: Hyperlinks
The system SHALL support creating and managing hyperlinks with advanced options.

#### Scenario: Create external hyperlink
- GIVEN a text run element
- WHEN run.AddHyperlink("https://example.com", "Link text") is called
- THEN hyperlink relationship is created
- AND link is rendered as blue underlined text

#### Scenario: Create internal bookmark hyperlink
- GIVEN a text run element
- WHEN run.AddBookmarkLink("BookmarkName", "Link text") is called
- THEN hyperlink points to internal bookmark
- AND link is internal (no external relationship)

#### Scenario: Create email hyperlink
- GIVEN a text run element
- WHEN run.AddEmailLink("user@example.com", "Send email") is called
- THEN hyperlink is created with mailto: scheme
- AND clicking opens email client

#### Scenario: Add hyperlink tooltip
- GIVEN a hyperlink element
- WHEN hyperlink.SetTooltip("This is a helpful tooltip") is called
- THEN tooltip attribute is set
- AND tooltip displays on mouse hover

#### Scenario: Modify hyperlink target
- GIVEN an existing hyperlink
- WHEN hyperlink.SetTarget("https://newsite.com") is called
- THEN hyperlink relationship is updated
- AND hyperlink points to new target

### Requirement: Document Protection
The system SHALL support document protection with various permission types.

#### Scenario: Protect document for comments only
- GIVEN a document
- WHEN document.Protect(ProtectionType.Comments, "password") is called
- THEN documentProtection element is set
- AND only comments can be edited, not document content

#### Scenario: Protect document for forms only
- GIVEN a document
- WHEN document.Protect(ProtectionType.Forms, "password") is called
- THEN only form fields can be edited
- AND protection persists in roundtrip

#### Scenario: Protect document for tracked changes only
- GIVEN a document
- WHEN document.Protect(ProtectionType.TrackedChanges, "password") is called
- THEN only tracked changes can be made
- AND all edits are tracked automatically

#### Scenario: Check protection status
- GIVEN a protected document
- WHEN document.IsProtected() is called
- THEN true is returned
- AND document.ProtectionType() returns the protection type

#### Scenario: Unprotect document
- GIVEN a protected document
- WHEN document.Unprotect("password") is called
- THEN protection is removed
- AND document is fully editable

### Requirement: Text Boxes and Frames
The system SHALL support creating and managing text boxes and floating frames.

#### Scenario: Create text box
- GIVEN a paragraph element
- WHEN paragraph.AddTextBox(width=1440, height=1440) is called
- THEN text box element is created
- AND text box can contain paragraphs and text

#### Scenario: Create anchored frame
- GIVEN a paragraph element
- WHEN paragraph.AddFrame(width=2880, height=1440, name="Frame1") is called
- THEN frame element is created with anchor
- AND frame positioning is set relative to paragraph

#### Scenario: Set frame positioning
- GIVEN a frame element
- WHEN frame.SetPosition(x=720, y=720, xAnchor=AnchorType.Margin, yAnchor=AnchorType.Paragraph) are called
- THEN positioning attributes are set
- AND frame position is fixed relative to anchor

#### Scenario: Set text box text wrapping
- GIVEN a text box element
- WHEN textbox.SetTextWrapping(TextWrapping.Square) is called
- THEN wrap attribute is set
- AND text flows around frame/textbox

### Requirement: Footnote and Endnote Configuration
The system SHALL support advanced footnote and endnote settings.

#### Scenario: Set footnote numbering format
- GIVEN a document with footnote properties part
- WHEN document.SetFootnoteNumberingFormat(NumberFormat.Arabic, startingValue=1) is called
- THEN footnote numbering format is set
- AND format applies to all footnotes

#### Scenario: Set custom footnote separator
- GIVEN a document
- WHEN document.SetFootnoteSeparator(customText) is called
- THEN footnote separator is customized
- AND separator appears in footnotes area

#### Scenario: Configure footnote position
- GIVEN a document with footnote settings
- WHEN document.SetFootnotePosition(FootnotePosition.PageBottom) is called
- THEN footnotes are positioned at page bottom
- AND alternative positions (section end, etc.) are supported

#### Scenario: Enable endnotes instead of footnotes
- GIVEN a document
- WHEN document.UseEndnotes(true) is called
- THEN footnotes are converted to endnotes
- AND endnotes appear at document/section end

## MODIFIED Requirements

### Requirement: Create WordprocessingDocument
The system SHALL support creating new Word documents with extended builder options.

#### Scenario: Create with advanced builder
- GIVEN CreateDefaultBuilder().WithContentControls(true).WithTrackingEnabled(true).Build()
- WHEN Create(path, type) is called
- THEN document includes optional advanced features
- AND features are configured via builder pattern
