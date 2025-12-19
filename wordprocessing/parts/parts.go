// Package parts provides Word document part types.
// Parts are the individual XML files within a .docx package, including
// the main document, styles, settings, and other document components.
//
// # Part Types
//
// The package provides the following part types:
//
//   - MainPart: The main document content (word/document.xml)
//   - StylesPart: Style definitions (word/styles.xml)
//   - NumberingPart: Numbering/list definitions (word/numbering.xml)
//   - SettingsPart: Document settings (word/settings.xml)
//   - WebSettingsPart: Web-related settings (word/webSettings.xml)
//   - FontsPart: Font table (word/fontTable.xml)
//   - HeaderPart: Header content (word/header1.xml, etc.)
//   - FooterPart: Footer content (word/footer1.xml, etc.)
//   - FootnotesPart: Footnotes (word/footnotes.xml)
//   - EndnotesPart: Endnotes (word/endnotes.xml)
//   - CommentsPart: Comments (word/comments.xml)
//   - ThemePart: Theme definitions (word/theme/theme1.xml)
//   - ImagePart: Embedded images (word/media/image1.png, etc.)
//   - CustomXmlPart: Custom XML data (customXml/item1.xml, etc.)
//   - VbaProjectPart: VBA macros for macro-enabled documents (word/vbaProject.bin)
//
// # Creating Parts
//
// Parts are typically created through the MainPart methods:
//
//	doc := wordprocessing.New("document.docx", wordprocessing.DocTypeDocument)
//	mainPart := doc.MainPart()
//
//	// Add supporting parts
//	stylesPart, _ := mainPart.AddStylesPart()
//	settingsPart, _ := mainPart.AddSettingsPart()
//	fontsPart, _ := mainPart.AddFontsPart()
//
// # Part Registration
//
// All part types are automatically registered with the openxml package's
// part type registry. This allows parts to be automatically instantiated
// when opening existing documents.
package parts

// This file serves as the package documentation and ensures all part types
// are registered via their init() functions.
