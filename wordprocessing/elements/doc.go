// Package elements provides Word document element types for WordprocessingML.
//
// This package contains the core elements used in WordprocessingML documents,
// including document structure, paragraphs, runs, tables, and formatting.
//
// # Document Structure
//
// The document hierarchy is:
//
//	Document
//	  -> Body
//	       -> Paragraph
//	            -> Run
//	                 -> Text
//	       -> Table
//	            -> TableRow
//	                 -> TableCell
//	                      -> Paragraph
//
// Creating a document structure:
//
//	doc := elements.NewDocument()
//	body := doc.Body()
//	p := body.AppendParagraph("Hello, World!")
//	r := p.AppendRun("More text")
//
// # Paragraphs
//
// Paragraphs are the primary block-level element:
//
//	p := elements.NewParagraph()
//	p.AppendRun("Text content")
//	p.SetStyle("Heading1")
//	p.SetJustification(elements.JustificationCenter)
//
// Paragraph properties:
//
//	props := p.GetOrCreateProperties()
//	props.SetSpacingBefore(240)  // 12pt in twips
//	props.SetSpacingAfter(240)
//	props.SetLeftIndent(720)     // 0.5 inch
//
// # Runs
//
// Runs contain inline content with consistent formatting:
//
//	r := elements.NewRun()
//	r.AppendText("Hello")
//	r.SetBold(true)
//	r.SetItalic(true)
//	r.SetUnderline(elements.UnderlineSingle)
//	r.SetColor("FF0000")
//	r.SetFontSize(24) // 12pt in half-points
//
// # Tables
//
// Create tables with rows and cells:
//
//	table := elements.NewTable(3, 4) // 3 rows, 4 columns
//	table.SetCellText(0, 0, "Header")
//	table.SetStyle("TableGrid")
//	table.SetWidth(5000, elements.TableWidthTypePct) // 100%
//
// Table navigation:
//
//	for row := range table.Rows() {
//		for cell := range row.Cells() {
//			fmt.Println(cell.InnerText())
//		}
//	}
//
// Cell manipulation:
//
//	cell := table.GetCell(0, 0)
//	cell.SetText("Content")
//	cell.SetShading("CCCCCC")
//	cell.SetHorizontalMerge(2) // Span 2 columns
//
// # Headers and Footers
//
// Header and footer elements contain paragraphs and tables:
//
//	header := elements.NewHeader()
//	header.AppendParagraph("Document Title")
//	header.AppendTable(1, 3)
//
//	footer := elements.NewFooter()
//	footer.AppendParagraph("Page ")
//
// # Styles
//
// Style elements define formatting presets:
//
//	style := elements.NewStyle(elements.StyleTypeParagraph)
//	style.SetStyleId("MyStyle")
//	style.SetName("My Custom Style")
//	style.SetBasedOn("Normal")
//
// # Numbering
//
// Numbering definitions for lists:
//
//	numbering := elements.NewNumbering()
//	abstractNum := elements.NewAbstractNum()
//	level := elements.NewLevel(0)
//	level.SetNumberFormat(elements.NumberFormatDecimal)
//	level.SetText("%1.")
//	abstractNum.AddLevel(level)
//	numbering.AddAbstractNum(abstractNum)
//
// # Breaks
//
// Insert various break types:
//
//	p.AppendBreak(elements.BreakLine)   // Line break
//	p.AppendBreak(elements.BreakPage)   // Page break
//	p.AppendBreak(elements.BreakColumn) // Column break
//
// # Drawings
//
// Inline drawings for images:
//
//	drawing := elements.NewInlineDrawing(width, height, relId)
//	r.AppendChild(drawing)
//
// # Formatting Values
//
// Common formatting enumerations:
//
//	// Justification
//	elements.JustificationLeft
//	elements.JustificationCenter
//	elements.JustificationRight
//	elements.JustificationBoth
//
//	// Underline styles
//	elements.UnderlineSingle
//	elements.UnderlineDouble
//	elements.UnderlineWave
//
//	// Border styles
//	elements.BorderSingle
//	elements.BorderDouble
//	elements.BorderDotted
//	elements.BorderDashed
//
//	// Highlight colors
//	elements.HighlightYellow
//	elements.HighlightGreen
//	elements.HighlightCyan
//
// # Measurement Units
//
// The package uses standard OOXML units:
//
//   - Twips: 1/20 of a point, 1440 twips = 1 inch
//   - Half-points: Font sizes, 24 half-points = 12pt
//   - EMUs: English Metric Units, 914400 EMU = 1 inch
//   - Fiftieths of percent: Table widths, 5000 = 100%
//
// Conversion constants:
//
//	elements.TwipsPerInch   // 1440
//	elements.EMUsPerInch    // 914400
//	elements.EMUsPerCm      // 360000
//
// # Namespaces
//
// Standard WordprocessingML namespaces:
//
//	elements.NamespaceWML  // Main namespace
//	elements.PrefixW       // "w:" prefix
//
// # Iterator Pattern
//
// Collections use Go 1.25 iterator pattern (iter.Seq):
//
//	for p := range body.Paragraphs() {
//		for r := range p.Runs() {
//			fmt.Println(r.InnerText())
//		}
//	}
//
// # Cloning Elements
//
// All elements support cloning:
//
//	clone := p.Clone().(*elements.Paragraph)
//	deepClone := p.CloneNode(true)
package elements
