package wordprocessing

import (
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestDocumentBuilder(t *testing.T) {
	builder := NewDocumentBuilder()

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	if doc == nil {
		t.Fatal("Build returned nil document")
	}
}

func TestDocumentBuilderAddParagraph(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Hello, World!")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Document body is nil")
	}

	count := 0
	var text string
	for p := range body.Paragraphs() {
		count++
		text = p.InnerText()
	}

	if count != 1 {
		t.Errorf(
			"Expected 1 paragraph, got %d",
			count,
		)
	}
	if text != "Hello, World!" {
		t.Errorf(
			"Expected 'Hello, World!', got '%s'",
			text,
		)
	}
}

func TestDocumentBuilderAddHeading(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddHeading("My Heading", 1)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		props := p.Properties()
		if props == nil {
			t.Fatal("Paragraph properties is nil")
		}
		styleId := props.ParagraphStyleId()
		if styleId != "Heading1" {
			t.Errorf(
				"Expected style 'Heading1', got '%s'",
				styleId,
			)
		}
	}
}

func TestDocumentBuilderAddTable(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddTable(2, 3).
		SetCellText(0, 0, "Cell 00").
		SetCellText(1, 2, "Cell 12")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	tableCount := 0
	for table := range body.Tables() {
		tableCount++
		if table.RowCount() != 2 {
			t.Errorf(
				"Expected 2 rows, got %d",
				table.RowCount(),
			)
		}
		cell00 := table.GetCell(0, 0)
		if cell00.InnerText() != "Cell 00" {
			t.Errorf(
				"Expected 'Cell 00', got '%s'",
				cell00.InnerText(),
			)
		}
	}

	if tableCount != 1 {
		t.Errorf(
			"Expected 1 table, got %d",
			tableCount,
		)
	}
}

func TestParagraphBuilderFormatting(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Bold and Italic").
		Bold().
		Italic()

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		for r := range p.Runs() {
			props := r.Properties()
			if props == nil {
				t.Fatal("Run properties is nil")
			}
			xml := props.OuterXml()
			if !strings.Contains(xml, "<w:b") {
				t.Error(
					"Expected bold formatting",
				)
			}
			if !strings.Contains(xml, "<w:i") {
				t.Error(
					"Expected italic formatting",
				)
			}
		}
	}
}

func TestParagraphBuilderAlignment(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Centered").AlignCenter()
	builder.AddParagraph("Right").AlignRight()
	builder.AddParagraph("Justified").
		AlignJustify()

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	alignments := []elements.JustificationValue{
		elements.JustificationCenter,
		elements.JustificationRight,
		elements.JustificationBoth,
	}

	i := 0
	for p := range body.Paragraphs() {
		if i >= len(alignments) {
			break
		}
		props := p.Properties()
		if props == nil {
			t.Fatal("Paragraph properties is nil")
		}
		if props.Justification() != alignments[i] {
			t.Errorf(
				"Paragraph %d: expected %s, got %s",
				i,
				alignments[i],
				props.Justification(),
			)
		}
		i++
	}
}

func TestParagraphBuilderFont(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Custom Font").
		Font("Arial", 14)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		for r := range p.Runs() {
			props := r.Properties()
			if props == nil {
				t.Fatal("Run properties is nil")
			}
			xml := props.OuterXml()
			if !strings.Contains(xml, "Arial") {
				t.Error("Expected Arial font")
			}
			// Font size is in half-points, so 14pt = 28
			if !strings.Contains(xml, "28") {
				t.Error(
					"Expected font size 28 (14pt)",
				)
			}
		}
	}
}

func TestParagraphBuilderSpacing(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("With Spacing").
		SpacingBefore(12).
		SpacingAfter(12)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		props := p.Properties()
		if props == nil {
			t.Fatal("Paragraph properties is nil")
		}
		spacing := props.SpacingBetweenLines()
		if spacing == nil {
			t.Fatal("Spacing is nil")
		}
		// 12 points = 240 twips
		if spacing.Before() != 240 {
			t.Errorf(
				"Expected before 240, got %d",
				spacing.Before(),
			)
		}
		if spacing.After() != 240 {
			t.Errorf(
				"Expected after 240, got %d",
				spacing.After(),
			)
		}
	}
}

func TestParagraphBuilderIndent(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Indented").
		LeftIndent(36). // 36 points
		FirstLineIndent(18)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		props := p.Properties()
		if props == nil {
			t.Fatal("Paragraph properties is nil")
		}
		ind := props.Indentation()
		if ind == nil {
			t.Fatal("Indentation is nil")
		}
		// 36 points = 720 twips
		if ind.Left() != 720 {
			t.Errorf(
				"Expected left 720, got %d",
				ind.Left(),
			)
		}
		// 18 points = 360 twips
		if ind.FirstLine() != 360 {
			t.Errorf(
				"Expected firstLine 360, got %d",
				ind.FirstLine(),
			)
		}
	}
}

func TestRunBuilder(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("").
		AddRun("Bold").Bold().
		Paragraph().
		AddRun(" Normal ").
		Paragraph().
		AddRun("Italic").Italic()

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		runCount := 0
		for range p.Runs() {
			runCount++
		}
		if runCount != 3 {
			t.Errorf(
				"Expected 3 runs, got %d",
				runCount,
			)
		}
	}
}

func TestTableBuilderStyle(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddTable(2, 2).
		SetStyle("TableGrid").
		SetWidthPercent(100)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for table := range body.Tables() {
		props := table.TableProperties()
		if props == nil {
			t.Fatal("Table properties is nil")
		}
		if props.TableStyle() != "TableGrid" {
			t.Errorf(
				"Expected style 'TableGrid', got '%s'",
				props.TableStyle(),
			)
		}
	}
}

func TestTableBuilderBorders(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddTable(2, 2).
		SetBorders(elements.BorderSingle, 4, "000000")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for table := range body.Tables() {
		props := table.TableProperties()
		if props == nil {
			t.Fatal("Table properties is nil")
		}
		borders := props.TableBorders()
		if borders == nil {
			t.Fatal("Table borders is nil")
		}
		xml := borders.OuterXml()
		if !strings.Contains(xml, "single") {
			t.Error(
				"Expected single border style",
			)
		}
	}
}

func TestTableCellBuilder(t *testing.T) {
	builder := NewDocumentBuilder()

	tb := builder.AddTable(2, 2)
	cell := tb.Cell(0, 0)
	if cell != nil {
		cell.SetText("Hello").
			SetShading("FFFF00").
			SetHorizontalMerge(2)
	}

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for table := range body.Tables() {
		c := table.GetCell(0, 0)
		if c.InnerText() != "Hello" {
			t.Errorf(
				"Expected 'Hello', got '%s'",
				c.InnerText(),
			)
		}
		props := c.TableCellProperties()
		if props == nil {
			t.Fatal("Cell properties is nil")
		}
		if props.GridSpan() != 2 {
			t.Errorf(
				"Expected gridSpan 2, got %d",
				props.GridSpan(),
			)
		}
	}
}

func TestTableRowBuilder(t *testing.T) {
	builder := NewDocumentBuilder()

	tb := builder.AddTable(2, 2)
	row := tb.Row(0)
	if row != nil {
		row.SetHeight(500, elements.HeightRuleExact).
			SetHeaderRow(true)
	}

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for table := range body.Tables() {
		r := table.GetRow(0)
		props := r.TableRowProperties()
		if props == nil {
			t.Fatal("Row properties is nil")
		}
		xml := props.OuterXml()
		if !strings.Contains(xml, "trHeight") {
			t.Error("Expected row height")
		}
		if !strings.Contains(xml, "tblHeader") {
			t.Error("Expected header row")
		}
	}
}

func TestDocumentBuilderPageBreak(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Page 1").
		Document().
		AddPageBreak().
		AddParagraph("Page 2")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	xml := doc.OuterXml()
	if !strings.Contains(xml, "page") {
		t.Error("Expected page break in document")
	}
}

func TestImageBuilderDimensions(t *testing.T) {
	builder := NewDocumentBuilder()

	img := builder.AddImage(
		[]byte{0x89, 0x50, 0x4E, 0x47},
		"image/png",
	)

	// Test EMU dimensions
	img.Width(914400).Height(914400)
	if img.width != 914400 {
		t.Errorf(
			"Expected width 914400, got %d",
			img.width,
		)
	}
	if img.height != 914400 {
		t.Errorf(
			"Expected height 914400, got %d",
			img.height,
		)
	}

	// Test inch dimensions
	img.WidthInches(2.0).HeightInches(1.5)
	expectedWidth := int64(
		2.0 * float64(elements.EMUsPerInch),
	)
	expectedHeight := int64(
		1.5 * float64(elements.EMUsPerInch),
	)
	if img.width != expectedWidth {
		t.Errorf(
			"Expected width %d, got %d",
			expectedWidth,
			img.width,
		)
	}
	if img.height != expectedHeight {
		t.Errorf(
			"Expected height %d, got %d",
			expectedHeight,
			img.height,
		)
	}

	// Test cm dimensions
	img.WidthCm(5.0).HeightCm(3.0)
	expectedWidth = int64(
		5.0 * float64(elements.EMUsPerCm),
	)
	expectedHeight = int64(
		3.0 * float64(elements.EMUsPerCm),
	)
	if img.width != expectedWidth {
		t.Errorf(
			"Expected width %d, got %d",
			expectedWidth,
			img.width,
		)
	}
	if img.height != expectedHeight {
		t.Errorf(
			"Expected height %d, got %d",
			expectedHeight,
			img.height,
		)
	}
}

func TestParagraphBuilderColor(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Red Text").
		Color("FF0000")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		for r := range p.Runs() {
			props := r.Properties()
			if props == nil {
				t.Fatal("Run properties is nil")
			}
			xml := props.OuterXml()
			if !strings.Contains(xml, "FF0000") {
				t.Error("Expected color FF0000")
			}
		}
	}
}

func TestParagraphBuilderHighlight(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddParagraph("Highlighted").
		Highlight(elements.HighlightYellow)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		for r := range p.Runs() {
			props := r.Properties()
			if props == nil {
				t.Fatal("Run properties is nil")
			}
			xml := props.OuterXml()
			if !strings.Contains(xml, "yellow") {
				t.Error(
					"Expected yellow highlight",
				)
			}
		}
	}
}

func TestBuildToBytes(t *testing.T) {
	builder := NewDocumentBuilder()
	builder.AddParagraph("Test")

	bytes, err := builder.BuildToBytes()
	if err != nil {
		t.Fatalf("BuildToBytes failed: %v", err)
	}

	if len(bytes) == 0 {
		t.Error(
			"BuildToBytes returned empty bytes",
		)
	}

	// Should contain XML declaration and document element
	s := string(bytes)
	if !strings.Contains(s, "document") {
		t.Error(
			"Expected document element in output",
		)
	}
}
