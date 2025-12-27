// Package elements provides Word document element types.
//
// This package contains the core elements used in
// WordprocessingML documents:
//
// # Document Structure
//
//   - Document: Root element for the main document part
//   - Body: Container for document content
//   - Paragraph: Text block element
//   - Run: Inline content element
//   - Text: Text content element
//
// # Formatting Properties
//
//   - ParagraphProperties: Paragraph-level formatting
//   - RunProperties: Character-level formatting
//   - Indentation: Paragraph indentation settings
//   - SpacingBetweenLines: Line and paragraph spacing
//   - NumberingProperties: List formatting
//
// # Section Layout
//
//   - SectionProperties: Page layout settings
//   - PageSize: Page dimensions
//   - PageMargins: Page margins
//   - Columns: Column layout
//
// # Special Elements
//
//   - Break: Line, page, and column breaks
//   - Tab: Tab characters
//   - Hyperlink: Links to external or internal targets
//   - BookmarkStart/BookmarkEnd: Named document locations
//
// # Tables
//
//   - Table: Table container
//   - TableRow: Table row
//   - TableCell: Table cell
//
// # Usage Example
//
//	doc := elements.NewDocument()
//	body := doc.Body()
//
//	// Add a paragraph
//	p := body.AppendParagraph("Hello, World!")
//	p.SetStyle("Heading1")
//
//	// Add formatted text
//	p2 := body.AppendParagraph("")
//	r := p2.AppendRun("Bold text")
//	r.SetBold(true)
//
//nolint:revive // max-public-structs: elements package defines many public types
package elements
