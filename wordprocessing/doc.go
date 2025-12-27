// Package wordprocessing provides WordprocessingML support for Word documents.
//
// This package implements the document-level API for creating, reading, and
// modifying .docx files (and related formats like .dotx, .docm, .dotm).
//
// # Creating Documents
//
// Create a new document at a file path:
//
//	doc, err := wordprocessing.New("document.docx", wordprocessing.DocTypeDocument)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer doc.Close()
//
// Create different document types:
//
//	doc, _ := wordprocessing.New("template.dotx", wordprocessing.DocTypeTemplate)
//	doc, _ := wordprocessing.New("macro.docm", wordprocessing.DocTypeMacroEnabled)
//
// Write to an io.Writer:
//
//	var buf bytes.Buffer
//	doc, err := wordprocessing.NewWriter(&buf, wordprocessing.DocTypeDocument)
//
// # Opening Documents
//
// Open an existing document:
//
//	doc, err := wordprocessing.Open("document.docx", true) // editable=true
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer doc.Close()
//
// Open read-only:
//
//	doc, err := wordprocessing.Open("document.docx", false)
//
// Open from an io.ReaderAt:
//
//	doc, err := wordprocessing.OpenReader(reader, size, true)
//
// # Using the Builder Pattern
//
// For creating documents programmatically, use the DocumentBuilder:
//
//	doc, err := wordprocessing.NewDocumentBuilder().
//		AddHeading("My Document", 1).
//		AlignCenter().
//		Document().
//		AddParagraph("Introduction text.").
//		SpacingAfter(12).
//		Document().
//		AddTable(3, 4).
//		SetCellText(0, 0, "Header").
//		Document().
//		Build()
//
// # Accessing Document Content
//
// Access the main document part:
//
//	mainPart := doc.MainPart()
//	docElement := mainPart.Document()
//	body := docElement.Body()
//
// Iterate over paragraphs:
//
//	for p := range body.Paragraphs() {
//		text := p.InnerText()
//		fmt.Println(text)
//	}
//
// Iterate over tables:
//
//	for table := range body.Tables() {
//		for row := range table.Rows() {
//			for cell := range row.Cells() {
//				fmt.Println(cell.InnerText())
//			}
//		}
//	}
//
// # Headers and Footers
//
// Add headers and footers:
//
//	header, err := doc.AddHeader(elements.HeaderFooterDefault)
//	header.AppendParagraph("Header text")
//
//	footer, err := doc.AddFooter(elements.HeaderFooterDefault)
//	footer.AppendParagraph("Page ")
//
// Iterate over headers/footers:
//
//	for header := range doc.Headers() {
//		// process header
//	}
//
// # Comments, Footnotes, Endnotes
//
// Add comments:
//
//	comment, err := doc.AddComment("Author Name", "Comment text")
//
// Access footnotes and endnotes:
//
//	fnPart, err := doc.FootnotesPart()
//	footnote := fnPart.AddFootnote("Footnote text")
//
//	enPart, err := doc.EndnotesPart()
//	endnote := enPart.AddEndnote("Endnote text")
//
// # Saving Documents
//
// Save to the original location:
//
//	err := doc.Save()
//
// Save to a new location:
//
//	err := doc.SaveAs("copy.docx")
//
// # Document Types
//
// The package supports different Word document types:
//
//   - DocTypeDocument: Standard document (.docx)
//   - DocTypeTemplate: Template (.dotx)
//   - DocTypeMacroEnabled: Macro-enabled document (.docm)
//   - DocTypeMacroTemplate: Macro-enabled template (.dotm)
//
// Change document type:
//
//	err := doc.ChangeType(wordprocessing.DocTypeTemplate)
//
// # Document Settings
//
// Configure document open settings:
//
//	settings := &wordprocessing.OpenSettings{
//		AutoSave: true,
//		MaxCharactersInPart: 0, // unlimited
//	}
//	doc, err := wordprocessing.OpenWithSettings("doc.docx", true, settings)
//
// # Thread Safety
//
// Document operations are protected by sync.RWMutex. Read operations can
// proceed concurrently, while write operations require exclusive access.
//
// # Subpackages
//
// The wordprocessing package is organized into subpackages:
//
//   - wordprocessing/elements: WordprocessingML element types
//   - wordprocessing/parts: Document part types (MainPart, StylesPart, etc.)
//
// See those packages for lower-level element manipulation.
//
//nolint:revive // line-length-limit: documentation examples contain long code samples
package wordprocessing
