// Package parts provides Word document part types.
//
// Parts are the individual XML files within a .docx package. Each part
// represents a distinct aspect of the document such as the main content,
// styles, settings, headers, footers, and more.
//
// # Part Types
//
// The package provides the following part types:
//
// Document Content:
//   - MainPart: Main document content (word/document.xml)
//   - HeaderPart: Header content (word/header1.xml, etc.)
//   - FooterPart: Footer content (word/footer1.xml, etc.)
//
// Styles and Formatting:
//   - StylesPart: Style definitions (word/styles.xml)
//   - NumberingPart: Numbering/list definitions (word/numbering.xml)
//   - FontsPart: Font table (word/fontTable.xml)
//   - ThemePart: Theme definitions (word/theme/theme1.xml)
//
// Settings:
//   - SettingsPart: Document settings (word/settings.xml)
//   - WebSettingsPart: Web-related settings (word/webSettings.xml)
//
// Annotations:
//   - FootnotesPart: Footnotes (word/footnotes.xml)
//   - EndnotesPart: Endnotes (word/endnotes.xml)
//   - CommentsPart: Comments (word/comments.xml)
//
// Media and Extensions:
//   - ImagePart: Embedded images (word/media/image1.png, etc.)
//   - CustomXmlPart: Custom XML data (customXml/item1.xml, etc.)
//   - VbaProjectPart: VBA macros (word/vbaProject.bin)
//
// # MainPart
//
// The MainPart is the central part of any Word document:
//
//	mainPart := doc.MainPart()
//	document := mainPart.Document()
//	body := document.Body()
//
// Add supporting parts through MainPart:
//
//	stylesPart, err := mainPart.AddStylesPart()
//	settingsPart, err := mainPart.AddSettingsPart()
//	numberingPart, err := mainPart.AddNumberingPart()
//
// # Headers and Footers
//
// Add header and footer parts:
//
//	headerPart, err := mainPart.AddHeaderPart()
//	header := headerPart.GetOrCreateHeader()
//	header.AppendParagraph("Header text")
//
//	footerPart, err := mainPart.AddFooterPart()
//	footer := footerPart.GetOrCreateFooter()
//	footer.AppendParagraph("Footer text")
//
// Access existing headers/footers:
//
//	for _, hp := range mainPart.HeaderParts() {
//		header := hp.Header()
//		// process header
//	}
//
// # Styles Part
//
// Access and modify document styles:
//
//	stylesPart, err := mainPart.AddStylesPart()
//	styles := stylesPart.Styles()
//
//	// Get a specific style
//	style := styles.GetStyle("Heading1")
//
//	// Add a new style
//	newStyle := elements.NewStyle(elements.StyleTypeParagraph)
//	newStyle.SetStyleId("CustomStyle")
//	styles.AddStyle(newStyle)
//
// # Numbering Part
//
// Manage numbered and bulleted lists:
//
//	numberingPart, err := mainPart.AddNumberingPart()
//	numbering := numberingPart.Numbering()
//
//	// Create abstract numbering definition
//	abstractNum := elements.NewAbstractNum()
//	numbering.AddAbstractNum(abstractNum)
//
//	// Create numbering instance
//	numInstance := numbering.CreateNumberingInstance(abstractNum.AbstractNumId())
//
// # Footnotes and Endnotes
//
// Add footnotes and endnotes:
//
//	fnPart, err := mainPart.AddFootnotesPart()
//	footnote := fnPart.AddFootnote("Footnote text")
//	id := footnote.Id() // Use this ID to reference in document
//
//	enPart, err := mainPart.AddEndnotesPart()
//	endnote := enPart.AddEndnote("Endnote text")
//
// # Comments
//
// Add document comments:
//
//	commentsPart, err := mainPart.AddCommentsPart()
//	comment := commentsPart.AddComment("Author", "Comment text")
//	id := comment.Id() // Use this ID to reference in document
//
// # Images
//
// Add image parts:
//
//	imagePart, relId, err := mainPart.AddImagePart("image/png")
//	imagePart.SetData(imageBytes)
//	// Use relId in Drawing element
//
// # Part URIs
//
// Standard part URIs follow OPC conventions:
//
//   - /word/document.xml - Main document
//   - /word/styles.xml - Styles
//   - /word/settings.xml - Settings
//   - /word/numbering.xml - Numbering
//   - /word/header1.xml - First header
//   - /word/footer1.xml - First footer
//   - /word/footnotes.xml - Footnotes
//   - /word/endnotes.xml - Endnotes
//   - /word/comments.xml - Comments
//   - /word/media/image1.png - First image
//
// # Content Types
//
// Each part type has a specific content type:
//
//	parts.ContentTypeMain     // Main document
//	parts.ContentTypeStyles   // Styles
//	parts.ContentTypeHeader   // Headers
//	parts.ContentTypeFooter   // Footers
//
// # Relationship Types
//
// Parts are connected via relationships:
//
//	parts.RelationshipTypeDocument  // Main document
//	parts.RelationshipTypeStyles    // Styles
//	parts.RelationshipTypeHeader    // Headers
//	parts.RelationshipTypeFooter    // Footers
//
// # Part Registration
//
// All part types are automatically registered with the openxml package's
// part type registry via init() functions. This enables automatic part
// instantiation when opening existing documents.
//
// # Thread Safety
//
// Part operations should be performed through the parent document's
// thread-safe API. Direct part manipulation is not thread-safe.
package parts
