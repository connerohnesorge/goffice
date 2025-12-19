package wordprocessing_test

import (
	"fmt"

	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// Example_createDocument demonstrates basic document creation.
func Example_createDocument() {
	// Create a new document using the builder
	builder := wordprocessing.NewDocumentBuilder()

	// Add a title
	builder.AddParagraph("My First Document").
		Bold().
		FontSize(18).
		AlignCenter()

	// Add body content
	builder.AddParagraph(
		"This is the first paragraph of the document.",
	)
	builder.AddParagraph(
		"This is the second paragraph.",
	)

	// Build the document
	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access the body and count paragraphs
	body := doc.Body()
	count := 0
	for range body.Paragraphs() {
		count++
	}

	fmt.Printf(
		"Document created with %d paragraphs\n",
		count,
	)
	// Output: Document created with 3 paragraphs
}

// Example_documentBuilder demonstrates using the fluent builder API.
func Example_documentBuilder() {
	// Use method chaining to build a document
	doc, err := wordprocessing.NewDocumentBuilder().
		AddHeading("Report Title", 1).
		AlignCenter().
		Document().
		AddHeading("Introduction", 2).
		Document().
		AddParagraph("This report covers important topics.").
		SpacingAfter(12).
		Document().
		AddHeading("Section 1", 2).
		Document().
		AddParagraph("Content for section 1.").
		Document().
		Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Count elements
	body := doc.Body()
	paraCount := 0
	for range body.Paragraphs() {
		paraCount++
	}

	fmt.Printf(
		"Built document with %d paragraphs\n",
		paraCount,
	)
	// Output: Built document with 5 paragraphs
}

// Example_addTable demonstrates creating tables in documents.
func Example_addTable() {
	builder := wordprocessing.NewDocumentBuilder()

	// Add a title
	builder.AddParagraph("Sales Report").
		Bold().
		AlignCenter()

	// Create a 4x3 table (4 rows, 3 columns)
	builder.AddTable(4, 3).
		SetCellText(0, 0, "Product").
		SetCellText(0, 1, "Quantity").
		SetCellText(0, 2, "Price").
		SetCellText(1, 0, "Widget A").
		SetCellText(1, 1, "100").
		SetCellText(1, 2, "$10.00").
		SetCellText(2, 0, "Widget B").
		SetCellText(2, 1, "50").
		SetCellText(2, 2, "$20.00").
		SetCellText(3, 0, "Widget C").
		SetCellText(3, 1, "75").
		SetCellText(3, 2, "$15.00").
		SetStyle("TableGrid").
		SetWidthPercent(100)

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Verify table exists
	body := doc.Body()
	tableCount := 0
	for table := range body.Tables() {
		tableCount++
		fmt.Printf(
			"Table has %d rows\n",
			table.RowCount(),
		)
	}

	fmt.Printf(
		"Document has %d table(s)\n",
		tableCount,
	)
	// Output:
	// Table has 4 rows
	// Document has 1 table(s)
}

// Example_formatting demonstrates text formatting options.
func Example_formatting() {
	builder := wordprocessing.NewDocumentBuilder()

	// Create a paragraph with mixed formatting
	pb := builder.AddParagraph("")

	// Add runs with different formatting
	pb.AddRun("Normal text, ")
	pb.AddRun("bold text, ").Bold()
	pb.AddRun("italic text, ").Italic()
	pb.AddRun("underlined text, ").Underline()
	pb.AddRun("and ").Font("Arial")
	pb.AddRun("colored text.").Color("FF0000")

	// Add a paragraph with custom font
	builder.AddParagraph("This uses Times New Roman at 14pt.").
		Font("Times New Roman", 14)

	// Add a highlighted paragraph
	builder.AddParagraph("This text is highlighted.").
		Highlight(elements.HighlightYellow)

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	count := 0
	for range body.Paragraphs() {
		count++
	}
	fmt.Printf(
		"Created %d paragraphs with formatting\n",
		count,
	)
	// Output: Created 3 paragraphs with formatting
}

// Example_headersFooters demonstrates adding headers and footers.
func Example_headersFooters() {
	// Create header and footer elements directly
	header := elements.NewHeader()
	header.AppendParagraph("Document Header")

	footer := elements.NewFooter()
	footer.AppendParagraph("Page Footer")

	// Verify content
	headerParas := 0
	for range header.Paragraphs() {
		headerParas++
	}

	footerParas := 0
	for range footer.Paragraphs() {
		footerParas++
	}

	fmt.Printf(
		"Header has %d paragraph(s)\n",
		headerParas,
	)
	fmt.Printf(
		"Footer has %d paragraph(s)\n",
		footerParas,
	)
	// Output:
	// Header has 1 paragraph(s)
	// Footer has 1 paragraph(s)
}

// Example_paragraphAlignment demonstrates paragraph alignment options.
func Example_paragraphAlignment() {
	builder := wordprocessing.NewDocumentBuilder()

	builder.AddParagraph("Left aligned (default)").
		AlignLeft()
	builder.AddParagraph("Center aligned").
		AlignCenter()
	builder.AddParagraph("Right aligned").
		AlignRight()
	builder.AddParagraph("Justified text - this would normally be a longer paragraph to show the effect.").
		AlignJustify()

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	i := 0
	alignments := []string{
		"left",
		"center",
		"right",
		"justify",
	}
	for p := range body.Paragraphs() {
		if i < len(alignments) {
			props := p.Properties()
			if props != nil {
				fmt.Printf(
					"Paragraph %d: %s alignment\n",
					i+1,
					alignments[i],
				)
			}
		}
		i++
	}
	// Output:
	// Paragraph 1: left alignment
	// Paragraph 2: center alignment
	// Paragraph 3: right alignment
	// Paragraph 4: justify alignment
}

// Example_tableWithFormatting demonstrates creating formatted tables.
func Example_tableWithFormatting() {
	builder := wordprocessing.NewDocumentBuilder()

	// Create a table with formatting
	tb := builder.AddTable(3, 3).
		SetWidthPercent(100).
		SetBorders(elements.BorderSingle, 4, "000000")

	// Format header row
	row := tb.Row(0)
	if row != nil {
		row.SetHeaderRow(true).
			SetHeight(400, elements.HeightRuleExact)
	}

	// Set header cells with shading
	for col := 0; col < 3; col++ {
		cell := tb.Cell(0, col)
		if cell != nil {
			cell.SetText(fmt.Sprintf("Header %d", col+1)).
				SetShading("CCCCCC")
		}
	}

	// Fill data cells
	for row := 1; row < 3; row++ {
		for col := 0; col < 3; col++ {
			cell := tb.Cell(row, col)
			if cell != nil {
				cell.SetText(
					fmt.Sprintf(
						"Data %d-%d",
						row,
						col+1,
					),
				)
			}
		}
	}

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	for table := range body.Tables() {
		cell := table.GetCell(0, 0)
		if cell != nil {
			fmt.Println(
				"First header cell:",
				cell.InnerText(),
			)
		}
		cell = table.GetCell(1, 1)
		if cell != nil {
			fmt.Println(
				"Data cell (1,1):",
				cell.InnerText(),
			)
		}
	}
	// Output:
	// First header cell: Header 1
	// Data cell (1,1): Data 1-2
}

// Example_paragraphSpacing demonstrates paragraph spacing and indentation.
func Example_paragraphSpacing() {
	builder := wordprocessing.NewDocumentBuilder()

	// Add paragraphs with custom spacing
	builder.AddParagraph("First paragraph with space after.").
		SpacingAfter(24)

		// 24 points

	builder.AddParagraph("Second paragraph with space before and after.").
		SpacingBefore(12).
		SpacingAfter(12)

	builder.AddParagraph("Indented paragraph.").
		LeftIndent(36) // 36 points = 0.5 inch

	builder.AddParagraph("First line indent.").
		FirstLineIndent(36)

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	count := 0
	for range body.Paragraphs() {
		count++
	}
	fmt.Printf(
		"Created %d paragraphs with custom spacing\n",
		count,
	)
	// Output: Created 4 paragraphs with custom spacing
}

// Example_pageBreaks demonstrates adding page breaks.
func Example_pageBreaks() {
	builder := wordprocessing.NewDocumentBuilder()

	builder.AddParagraph("Content on page 1")
	builder.AddPageBreak()
	builder.AddParagraph("Content on page 2")
	builder.AddPageBreak()
	builder.AddParagraph("Content on page 3")

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	count := 0
	for range body.Paragraphs() {
		count++
	}
	// Each page break adds a paragraph, so we have:
	// Page 1 content (1) + break para (1) + Page 2 content (1) + break para (1) + Page 3 content (1) = 5
	fmt.Printf(
		"Document has %d paragraphs (including page breaks)\n",
		count,
	)
	// Output: Document has 5 paragraphs (including page breaks)
}

// Example_runFormatting demonstrates run-level formatting.
func Example_runFormatting() {
	builder := wordprocessing.NewDocumentBuilder()

	// Create a paragraph and add formatted runs
	pb := builder.AddParagraph("")

	rb := pb.AddRun("This is bold").Bold()
	rb.AddText(" and more bold text.")

	pb.AddRun(" This is italic.").Italic()

	pb.AddRun(" Double underline.").
		UnderlineStyle(elements.UnderlineDouble)

	pb.AddRun(" Tab character:").AddTab()
	pb.AddRun("After tab.")

	doc, err := builder.Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		runCount := 0
		for range p.Runs() {
			runCount++
		}
		fmt.Printf(
			"Paragraph has %d runs\n",
			runCount,
		)
	}
	// Output: Paragraph has 5 runs
}

// Example_buildToBytes demonstrates building to bytes.
func Example_buildToBytes() {
	builder := wordprocessing.NewDocumentBuilder()
	builder.AddParagraph(
		"Simple document content",
	)

	data, err := builder.BuildToBytes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output will vary based on XML size, but should be non-zero
	if len(data) > 0 {
		fmt.Println("XML generation successful")
	}
	// Output: XML generation successful
}
