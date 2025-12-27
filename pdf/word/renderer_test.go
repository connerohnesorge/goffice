package word

import (
	"fmt"
	"os"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestWordRenderer_Basic(t *testing.T) {
	// 1. Create a dummy Word document
	doc, err := wordprocessing.New(
		"test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("test.docx")

	// 2. Setup font cache and engine
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)

	// 3. Render
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("test.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("test.pdf")
}

func TestWordRenderer_ParagraphStyles(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"styles.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("styles.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// 1. Spacing
	p1 := body.AppendParagraph(
		"Paragraph with spacing",
	)
	p1.SetSpacingBefore(240) // 12pt
	p1.SetSpacingAfter(240)  // 12pt

	// 2. Alignment
	p2 := body.AppendParagraph(
		"Centered paragraph",
	)
	p2.SetJustification(
		elements.JustificationCenter,
	)

	p3 := body.AppendParagraph(
		"Right aligned paragraph",
	)
	p3.SetJustification(
		elements.JustificationRight,
	)

	// 3. Borders and Shading
	p4 := body.AppendParagraph(
		"Paragraph with shading",
	)
	p4.GetOrCreateProperties().
		GetOrCreateShading().
		SetFill("FFFF00")
		// Yellow

	p5 := body.AppendParagraph(
		"Paragraph with borders",
	)
	pb := p5.GetOrCreateProperties().
		GetOrCreateParagraphBorders()
	pb.SetTop(elements.BorderSingle, 4, "FF0000")
	pb.SetBottom(
		elements.BorderSingle,
		4,
		"0000FF",
	)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("styles.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("styles.pdf")
}

func TestWordRenderer_RunStyles(t *testing.T) {
	doc, _ := wordprocessing.New(
		"runstyles.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("runstyles.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	p := body.AppendParagraph("")
	p.AppendRun("Bold text").SetBold(true)
	p.AppendRun(" Italic text").SetItalic(true)
	p.AppendRun(" Red text").SetColor("FF0000")
	p.AppendRun(" Underlined").
		SetUnderline(elements.UnderlineSingle)
	p.AppendRun(" Strikethrough").SetStrike(true)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("runstyles.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("runstyles.pdf")
}

func TestWordRenderer_SubscriptSuperscript(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"subsup.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("subsup.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Test subscript
	p1 := body.AppendParagraph("")
	p1.AppendRun("H")
	p1.AppendRun("2").
		SetVerticalTextAlignment(elements.VerticalAlignSubscript)
	p1.AppendRun("O (water)")

	// Test superscript
	p2 := body.AppendParagraph("")
	p2.AppendRun("E=mc")
	p2.AppendRun("2").
		SetVerticalTextAlignment(elements.VerticalAlignSuperscript)

	// Combined with other styles
	p3 := body.AppendParagraph("")
	p3.AppendRun("X")
	r := p3.AppendRun("n")
	r.SetVerticalTextAlignment(
		elements.VerticalAlignSuperscript,
	)
	r.SetBold(true)
	r.SetColor("0000FF")

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("subsup.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("subsup.pdf")
}

func TestWordRenderer_TextHighlight(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"highlight.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("highlight.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Test paragraph shading (background)
	p1 := body.AppendParagraph(
		"This paragraph has yellow background",
	)
	p1.GetOrCreateProperties().
		GetOrCreateShading().
		SetFill("FFFF00")

	// Test different shading colors
	p2 := body.AppendParagraph(
		"This paragraph has light blue background",
	)
	p2.GetOrCreateProperties().
		GetOrCreateShading().
		SetFill("ADD8E6")

	// Test shading with text color
	p3 := body.AppendParagraph("")
	p3.GetOrCreateProperties().
		GetOrCreateShading().
		SetFill("90EE90")
	p3.AppendRun("Green background with red text").
		SetColor("FF0000")

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("highlight.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("highlight.pdf")
}

func TestWordRenderer_UnderlineStyles(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"underlines.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("underlines.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Test various underline styles
	p := body.AppendParagraph("")
	p.AppendRun("Single underline").
		SetUnderline(elements.UnderlineSingle)
	p.AppendRun(" | ")
	p.AppendRun("Double underline").
		SetUnderline(elements.UnderlineDouble)
	p.AppendRun(" | ")
	p.AppendRun("Thick underline").
		SetUnderline(elements.UnderlineThick)
	p.AppendRun(" | ")
	p.AppendRun("Dotted underline").
		SetUnderline(elements.UnderlineDotted)
	p.AppendRun(" | ")
	p.AppendRun("Dashed underline").
		SetUnderline(elements.UnderlineDash)
	p.AppendRun(" | ")
	p.AppendRun("Wave underline").
		SetUnderline(elements.UnderlineWave)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("underlines.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("underlines.pdf")
}

func TestWordRenderer_FontSizes(t *testing.T) {
	doc, _ := wordprocessing.New(
		"fontsizes.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("fontsizes.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Test various font sizes (in half-points, so 20 = 10pt, 48 = 24pt, etc.)
	sizes := []struct {
		size int
		desc string
	}{
		{16, "8pt"},
		{20, "10pt"},
		{24, "12pt"},
		{32, "16pt"},
		{48, "24pt"},
		{72, "36pt"},
	}

	for _, s := range sizes {
		p := body.AppendParagraph("")
		p.AppendRun(s.desc + " font size").
			SetFontSize(s.size)
	}

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("fontsizes.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("fontsizes.pdf")
}

func TestWordRenderer_CombinedStyles(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"combined.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("combined.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Test combinations of styles
	p1 := body.AppendParagraph("")
	r1 := p1.AppendRun("Bold + Underline")
	r1.SetBold(true)
	r1.SetUnderline(elements.UnderlineSingle)

	p2 := body.AppendParagraph("")
	r2 := p2.AppendRun(
		"Italic + Strikethrough + Color",
	)
	r2.SetItalic(true)
	r2.SetStrike(true)
	r2.SetColor("FF00FF")

	p3 := body.AppendParagraph("")
	r3 := p3.AppendRun(
		"Bold + Italic + Underline + Color + Large",
	)
	r3.SetBold(true)
	r3.SetItalic(true)
	r3.SetUnderline(elements.UnderlineDouble)
	r3.SetColor("0000FF")
	r3.SetFontSize(32)

	p4 := body.AppendParagraph("")
	r4 := p4.AppendRun(
		"Superscript + Bold + Color",
	)
	r4.SetVerticalTextAlignment(
		elements.VerticalAlignSuperscript,
	)
	r4.SetBold(true)
	r4.SetColor("FF0000")

	p5 := body.AppendParagraph("")
	p5.GetOrCreateProperties().
		GetOrCreateShading().
		SetFill("FFFF99")
	r5 := p5.AppendRun(
		"Background + Underline + Strike + Color",
	)
	r5.SetUnderline(elements.UnderlineWave)
	r5.SetStrike(true)
	r5.SetColor("00FF00")

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("combined.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("combined.pdf")
}

func TestWordRenderer_SimpleTable(t *testing.T) {
	doc, _ := wordprocessing.New(
		"table_simple.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_simple.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create a simple 3x3 table
	tbl := elements.NewTable(3, 3)

	// Set table width
	tbl.SetWidth(
		5000,
		elements.TableWidthTypeDxa,
	) // 250 points

	// Fill cells with content
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tbl.SetCellText(
				i,
				j,
				fmt.Sprintf(
					"Row %d, Col %d",
					i+1,
					j+1,
				),
			)
		}
	}

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("table_simple.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_simple.pdf")
}

func TestWordRenderer_TableWithBorders(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"table_borders.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_borders.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create table with borders
	tbl := elements.NewTable(2, 2)

	// Set table borders
	borders := tbl.GetOrCreateTableProperties().
		GetOrCreateTableBorders()
	borders.SetAllBorders(
		elements.BorderSingle,
		8,
		"000000",
	)

	// Fill cells
	tbl.SetCellText(0, 0, "Top Left")
	tbl.SetCellText(0, 1, "Top Right")
	tbl.SetCellText(1, 0, "Bottom Left")
	tbl.SetCellText(1, 1, "Bottom Right")

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("table_borders.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_borders.pdf")
}

func TestWordRenderer_TableWithShading(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"table_shading.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_shading.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create table with shading
	tbl := elements.NewTable(3, 3)

	// Color cells in a pattern
	colors := []string{
		"FFCCCC",
		"CCFFCC",
		"CCCCFF",
		"FFFFCC",
		"FFCCFF",
		"CCFFFF",
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			idx := (i*3 + j) % len(colors)
			cell := tbl.GetCell(i, j)
			if cell != nil {
				cell.SetShading(colors[idx])
				cell.SetText(
					fmt.Sprintf(
						"Cell %d",
						i*3+j+1,
					),
				)
			}
		}
	}

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("table_shading.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_shading.pdf")
}

func TestWordRenderer_TableWithMergedCells(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"table_merged.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_merged.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create table with merged cells
	tbl := elements.NewTable(3, 4)

	// Horizontal merge: merge first two cells of first row
	cell00 := tbl.GetCell(0, 0)
	if cell00 != nil {
		cell00.SetHorizontalMerge(
			2,
		) // Span 2 columns
		cell00.SetText("Merged Horizontally")
	}

	// Vertical merge: merge cells in last column
	cell03 := tbl.GetCell(0, 3)
	if cell03 != nil {
		cell03.SetVerticalMerge(
			elements.VerticalMergeRestart,
		)
		cell03.SetText("Merged Vertically")
	}
	cell13 := tbl.GetCell(1, 3)
	if cell13 != nil {
		cell13.SetVerticalMerge(
			elements.VerticalMergeContinue,
		)
	}
	cell23 := tbl.GetCell(2, 3)
	if cell23 != nil {
		cell23.SetVerticalMerge(
			elements.VerticalMergeContinue,
		)
	}

	// Fill other cells
	tbl.SetCellText(0, 2, "Col 3")
	tbl.SetCellText(1, 0, "Row 2, Col 1")
	tbl.SetCellText(1, 1, "Row 2, Col 2")
	tbl.SetCellText(1, 2, "Row 2, Col 3")
	tbl.SetCellText(2, 0, "Row 3, Col 1")
	tbl.SetCellText(2, 1, "Row 3, Col 2")
	tbl.SetCellText(2, 2, "Row 3, Col 3")

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("table_merged.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_merged.pdf")
}

func TestWordRenderer_TableWithCellBorders(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"table_cell_borders.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_cell_borders.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create table with custom cell borders
	tbl := elements.NewTable(2, 2)

	// Set different borders for each cell
	cell00 := tbl.GetCell(0, 0)
	if cell00 != nil {
		borders := cell00.GetOrCreateTableCellProperties().
			GetOrCreateTableCellBorders()
		borders.SetTop(
			elements.BorderSingle,
			12,
			"FF0000",
		)
		borders.SetLeft(
			elements.BorderSingle,
			12,
			"FF0000",
		)
		cell00.SetText("Red Border (Top/Left)")
	}

	cell01 := tbl.GetCell(0, 1)
	if cell01 != nil {
		borders := cell01.GetOrCreateTableCellProperties().
			GetOrCreateTableCellBorders()
		borders.SetTop(
			elements.BorderSingle,
			12,
			"00FF00",
		)
		borders.SetRight(
			elements.BorderSingle,
			12,
			"00FF00",
		)
		cell01.SetText("Green Border (Top/Right)")
	}

	cell10 := tbl.GetCell(1, 0)
	if cell10 != nil {
		borders := cell10.GetOrCreateTableCellProperties().
			GetOrCreateTableCellBorders()
		borders.SetBottom(
			elements.BorderSingle,
			12,
			"0000FF",
		)
		borders.SetLeft(
			elements.BorderSingle,
			12,
			"0000FF",
		)
		cell10.SetText(
			"Blue Border (Bottom/Left)",
		)
	}

	cell11 := tbl.GetCell(1, 1)
	if cell11 != nil {
		borders := cell11.GetOrCreateTableCellProperties().
			GetOrCreateTableCellBorders()
		borders.SetAllBorders(
			elements.BorderDouble,
			8,
			"FF00FF",
		)
		cell11.SetText(
			"Purple Double Border (All)",
		)
	}

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render(
		"table_cell_borders.pdf",
	)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_cell_borders.pdf")
}

func TestWordRenderer_ComplexTable(t *testing.T) {
	doc, _ := wordprocessing.New(
		"table_complex.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("table_complex.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Create a complex table with multiple features
	tbl := elements.NewTable(4, 4)

	// Set table borders
	borders := tbl.GetOrCreateTableProperties().
		GetOrCreateTableBorders()
	borders.SetAllBorders(
		elements.BorderSingle,
		6,
		"000000",
	)

	// Header row with shading
	for j := 0; j < 4; j++ {
		cell := tbl.GetCell(0, j)
		if cell != nil {
			cell.SetShading("4472C4")
			para := cell.AppendParagraph(
				fmt.Sprintf("Header %d", j+1),
			)
			for run := range para.Runs() {
				run.SetColor("FFFFFF")
				run.SetBold(true)
			}
		}
	}

	// Data rows with alternating shading
	for i := 1; i < 4; i++ {
		rowColor := "FFFFFF"
		if i%2 == 0 {
			rowColor = "F2F2F2"
		}
		for j := 0; j < 4; j++ {
			cell := tbl.GetCell(i, j)
			if cell != nil {
				cell.SetShading(rowColor)
				cell.SetText(
					fmt.Sprintf(
						"Data %d-%d",
						i,
						j+1,
					),
				)
			}
		}
	}

	body.AppendChild(tbl)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("table_complex.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("table_complex.pdf")
}

// Pagination Tests

func TestWordRenderer_PageBreakBefore(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"pagebreak.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("pagebreak.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Add first paragraph
	body.AppendParagraph("This is on page 1")

	// Add paragraph with page break before
	p2 := body.AppendParagraph(
		"This should be on page 2",
	)
	p2.SetPageBreakBefore(true)

	// Add more content
	body.AppendParagraph("Still on page 2")

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("pagebreak.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("pagebreak.pdf")

	// Manual verification: The PDF should have at least 2 pages
	t.Log(
		"Manual verification needed: Check that pagebreak.pdf has 2 pages",
	)
}

func TestWordRenderer_KeepWithNext(t *testing.T) {
	doc, _ := wordprocessing.New(
		"keepnext.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("keepnext.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Fill first page
	for i := 0; i < 40; i++ {
		body.AppendParagraph(
			fmt.Sprintf(
				"Filler paragraph %d",
				i+1,
			),
		)
	}

	// Add heading that should keep with next paragraph
	heading := body.AppendParagraph(
		"Important Heading",
	)
	heading.SetKeepNext(true)

	// Add content paragraph
	body.AppendParagraph(
		"This content should stay with the heading above",
	)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("keepnext.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("keepnext.pdf")

	t.Log(
		"Manual verification needed: Check that heading and content stay together",
	)
}

func TestWordRenderer_KeepLinesTogether(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"keeplines.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("keeplines.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Fill most of first page
	for i := 0; i < 38; i++ {
		body.AppendParagraph(
			fmt.Sprintf("Filler line %d", i+1),
		)
	}

	// Add long paragraph that should stay together
	longText := "This is a long paragraph with multiple lines. "
	for i := 0; i < 20; i++ {
		longText += "This paragraph has keep-lines-together enabled. "
	}
	p := body.AppendParagraph(longText)
	p.SetKeepLines(true)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("keeplines.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("keeplines.pdf")

	t.Log(
		"Manual verification needed: Check that long paragraph stays on one page",
	)
}

func TestWordRenderer_WidowOrphanControl(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"widoworphan.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("widoworphan.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Fill most of first page
	for i := 0; i < 39; i++ {
		body.AppendParagraph(
			fmt.Sprintf("Line %d", i+1),
		)
	}

	// Add paragraph that would create widow/orphan without control
	longText := ""
	for i := 0; i < 5; i++ {
		longText += fmt.Sprintf(
			"Line %d of multi-line paragraph. ",
			i+1,
		)
	}
	body.AppendParagraph(longText)
	// Widow/orphan control is enabled by default

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("widoworphan.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("widoworphan.pdf")

	t.Log(
		"Manual verification needed: Check that no single lines are isolated",
	)
}

func TestWordRenderer_SectionBreak(t *testing.T) {
	doc, _ := wordprocessing.New(
		"sectionbreak.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("sectionbreak.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Section 1
	p1 := body.AppendParagraph(
		"This is section 1",
	)

	// Create section break with section properties
	sectPr := elements.NewSectionProperties()
	sectPr.SetSectionType(
		elements.SectionTypeNextPage,
	)
	// Add section properties to paragraph properties
	p1.GetOrCreateProperties().AppendChild(sectPr)

	// Section 2
	body.AppendParagraph(
		"This is section 2 on a new page",
	)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("sectionbreak.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("sectionbreak.pdf")

	t.Log(
		"Manual verification needed: Check that section 2 starts on new page",
	)
}

func TestWordRenderer_MixedPaginationFeatures(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"pagination_mixed.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("pagination_mixed.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	root := mainPart.Document()
	var body *elements.Body
	for child := range root.Children() {
		if child.LocalName() == "body" {
			if b, ok := child.(*elements.Body); ok {
				body = b
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				body = &elements.Body{CompositeElementBase: comp}
			}
			break
		}
	}

	if body == nil {
		t.Fatal("Failed to find body element")
	}

	// Page 1
	body.AppendParagraph("Page 1 content")

	// Force page break
	p2 := body.AppendParagraph("Page 2 content")
	p2.SetPageBreakBefore(true)

	// Add some filler
	for i := 0; i < 30; i++ {
		body.AppendParagraph(
			fmt.Sprintf("Filler %d", i+1),
		)
	}

	// Add heading with keep-with-next
	h1 := body.AppendParagraph("Section Heading")
	h1.SetKeepNext(true)

	// Content that should stay with heading
	body.AppendParagraph(
		"Section content that stays with heading",
	)

	// Add long paragraph with keep-lines-together
	longText := ""
	for i := 0; i < 15; i++ {
		longText += "This paragraph should stay together on one page. "
	}
	pLong := body.AppendParagraph(longText)
	pLong.SetKeepLines(true)

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("pagination_mixed.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("pagination_mixed.pdf")

	t.Log(
		"Manual verification needed: Check all pagination features work correctly",
	)
}
