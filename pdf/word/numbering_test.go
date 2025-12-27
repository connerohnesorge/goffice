package word

import (
	"os"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestNumberingState_BulletList(t *testing.T) {
	// Create a document with a bullet list
	doc, err := wordprocessing.New(
		"bullet_test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("bullet_test.docx")

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

	// Create numbering part
	numPart, err := mainPart.AddNumberingPart()
	if err != nil {
		t.Fatalf(
			"Failed to add numbering part: %v",
			err,
		)
	}

	// Load the numbering
	_ = numPart.Reload()

	// Create numbering directly via the part's root element
	numbering := elements.NewNumbering()

	// Create a bullet list abstract numbering
	abstractNum := elements.NewStandardBulletList()
	abstractId := numbering.AddAbstractNum(
		abstractNum,
	)

	// Create a numbering instance
	numInstance := numbering.CreateNumberingInstance(
		abstractId,
	)
	numId := numInstance.NumId()

	// Save the numbering back to the part
	numPart.SetData([]byte(numbering.OuterXml()))

	// Add list items
	items := []string{
		"First bullet point",
		"Second bullet point",
		"Third bullet point",
	}

	for _, item := range items {
		p := body.AppendParagraph(item)
		p.SetNumbering(numId, 0)
	}

	// Render to PDF
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("bullet_list.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("bullet_list.pdf")

	// Verify the PDF was created
	if _, err := os.Stat("bullet_list.pdf"); os.IsNotExist(
		err,
	) {
		t.Fatal("PDF file was not created")
	}
}

func TestNumberingState_DecimalList(
	t *testing.T,
) {
	// Create a document with a numbered list
	doc, err := wordprocessing.New(
		"decimal_test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("decimal_test.docx")

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

	// Create numbering part
	numPart, err := mainPart.AddNumberingPart()
	if err != nil {
		t.Fatalf(
			"Failed to add numbering part: %v",
			err,
		)
	}

	// Load the numbering
	_ = numPart.Reload()

	// Create numbering directly
	numbering := elements.NewNumbering()

	// Create a decimal list abstract numbering
	abstractNum := elements.NewDecimalList()
	abstractId := numbering.AddAbstractNum(
		abstractNum,
	)

	// Create a numbering instance
	numInstance := numbering.CreateNumberingInstance(
		abstractId,
	)
	numId := numInstance.NumId()

	// Save the numbering back to the part
	numPart.SetData([]byte(numbering.OuterXml()))

	// Add list items
	items := []string{
		"First numbered item",
		"Second numbered item",
		"Third numbered item",
		"Fourth numbered item",
		"Fifth numbered item",
	}

	for _, item := range items {
		p := body.AppendParagraph(item)
		p.SetNumbering(numId, 0)
	}

	// Render to PDF
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("decimal_list.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("decimal_list.pdf")

	// Verify the PDF was created
	if _, err := os.Stat("decimal_list.pdf"); os.IsNotExist(
		err,
	) {
		t.Fatal("PDF file was not created")
	}
}

func TestNumberingState_MultiLevelList(
	t *testing.T,
) {
	// Create a document with a multi-level list
	doc, err := wordprocessing.New(
		"multilevel_test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("multilevel_test.docx")

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

	// Create numbering part
	numPart, err := mainPart.AddNumberingPart()
	if err != nil {
		t.Fatalf(
			"Failed to add numbering part: %v",
			err,
		)
	}

	// Load the numbering
	_ = numPart.Reload()

	// Create numbering directly
	numbering := elements.NewNumbering()

	// Create an outline list abstract numbering
	abstractNum := elements.NewOutlineList()
	abstractId := numbering.AddAbstractNum(
		abstractNum,
	)

	// Create a numbering instance
	numInstance := numbering.CreateNumberingInstance(
		abstractId,
	)
	numId := numInstance.NumId()

	// Save the numbering back to the part
	numPart.SetData([]byte(numbering.OuterXml()))

	// Add multi-level list items
	type listItem struct {
		text  string
		level int
	}

	items := []listItem{
		{"Chapter 1", 0},
		{"Section 1.1", 1},
		{"Section 1.2", 1},
		{"Subsection 1.2.1", 2},
		{"Subsection 1.2.2", 2},
		{"Section 1.3", 1},
		{"Chapter 2", 0},
		{"Section 2.1", 1},
		{"Section 2.2", 1},
	}

	for _, item := range items {
		p := body.AppendParagraph(item.text)
		p.SetNumbering(numId, item.level)
	}

	// Render to PDF
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("multilevel_list.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("multilevel_list.pdf")

	// Verify the PDF was created
	if _, err := os.Stat("multilevel_list.pdf"); os.IsNotExist(
		err,
	) {
		t.Fatal("PDF file was not created")
	}
}

func TestNumberingState_AlphabeticList(
	t *testing.T,
) {
	// Create a document with alphabetic lists
	doc, err := wordprocessing.New(
		"alpha_test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("alpha_test.docx")

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

	// Create numbering part
	numPart, err := mainPart.AddNumberingPart()
	if err != nil {
		t.Fatalf(
			"Failed to add numbering part: %v",
			err,
		)
	}

	// Load the numbering
	_ = numPart.Reload()

	// Create numbering directly
	numbering := elements.NewNumbering()

	// Create lowercase alphabetic list
	abstractNum := elements.NewAlphabeticList(
		true,
	)
	abstractId := numbering.AddAbstractNum(
		abstractNum,
	)
	numInstance := numbering.CreateNumberingInstance(
		abstractId,
	)
	numId := numInstance.NumId()

	// Save the numbering back to the part
	numPart.SetData([]byte(numbering.OuterXml()))

	// Add title
	body.AppendParagraph(
		"Lowercase alphabetic list:",
	)

	// Add list items
	items := []string{
		"First item",
		"Second item",
		"Third item",
	}
	for _, item := range items {
		p := body.AppendParagraph(item)
		p.SetNumbering(numId, 0)
	}

	// Render to PDF
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("alpha_list.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("alpha_list.pdf")

	// Verify the PDF was created
	if _, err := os.Stat("alpha_list.pdf"); os.IsNotExist(
		err,
	) {
		t.Fatal("PDF file was not created")
	}
}

func TestNumberingState_RomanNumeralList(
	t *testing.T,
) {
	// Create a document with Roman numeral lists
	doc, err := wordprocessing.New(
		"roman_test.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("roman_test.docx")

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

	// Create numbering part
	numPart, err := mainPart.AddNumberingPart()
	if err != nil {
		t.Fatalf(
			"Failed to add numbering part: %v",
			err,
		)
	}

	// Load the numbering
	_ = numPart.Reload()

	// Create numbering directly
	numbering := elements.NewNumbering()

	// Create uppercase Roman numeral list
	abstractNum := elements.NewRomanNumeralList(
		false,
	)
	abstractId := numbering.AddAbstractNum(
		abstractNum,
	)
	numInstance := numbering.CreateNumberingInstance(
		abstractId,
	)
	numId := numInstance.NumId()

	// Save the numbering back to the part
	numPart.SetData([]byte(numbering.OuterXml()))

	// Add title
	body.AppendParagraph(
		"Uppercase Roman numeral list:",
	)

	// Add list items (test up to 10 to see various numerals)
	items := []string{
		"First item",
		"Second item",
		"Third item",
		"Fourth item",
		"Fifth item",
		"Sixth item",
		"Seventh item",
		"Eighth item",
		"Ninth item",
		"Tenth item",
	}
	for _, item := range items {
		p := body.AppendParagraph(item)
		p.SetNumbering(numId, 0)
	}

	// Render to PDF
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	err = renderer.Render("roman_list.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("roman_list.pdf")

	// Verify the PDF was created
	if _, err := os.Stat("roman_list.pdf"); os.IsNotExist(
		err,
	) {
		t.Fatal("PDF file was not created")
	}
}

// Test number formatting functions
func TestNumberFormatting(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		format   elements.NumberFormatValue
		expected string
	}{
		{
			"Decimal 1",
			1,
			elements.NumberFormatDecimal,
			"1",
		},
		{
			"Decimal 42",
			42,
			elements.NumberFormatDecimal,
			"42",
		},
		{
			"Upper Roman 1",
			1,
			elements.NumberFormatUpperRoman,
			"I",
		},
		{
			"Upper Roman 4",
			4,
			elements.NumberFormatUpperRoman,
			"IV",
		},
		{
			"Upper Roman 9",
			9,
			elements.NumberFormatUpperRoman,
			"IX",
		},
		{
			"Upper Roman 27",
			27,
			elements.NumberFormatUpperRoman,
			"XXVII",
		},
		{
			"Lower Roman 3",
			3,
			elements.NumberFormatLowerRoman,
			"iii",
		},
		{
			"Upper Letter 1",
			1,
			elements.NumberFormatUpperLetter,
			"A",
		},
		{
			"Upper Letter 26",
			26,
			elements.NumberFormatUpperLetter,
			"Z",
		},
		{
			"Lower Letter 1",
			1,
			elements.NumberFormatLowerLetter,
			"a",
		},
		{
			"Lower Letter 3",
			3,
			elements.NumberFormatLowerLetter,
			"c",
		},
		{
			"Ordinal 1",
			1,
			elements.NumberFormatOrdinal,
			"1st",
		},
		{
			"Ordinal 2",
			2,
			elements.NumberFormatOrdinal,
			"2nd",
		},
		{
			"Ordinal 3",
			3,
			elements.NumberFormatOrdinal,
			"3rd",
		},
		{
			"Ordinal 4",
			4,
			elements.NumberFormatOrdinal,
			"4th",
		},
		{
			"Ordinal 11",
			11,
			elements.NumberFormatOrdinal,
			"11th",
		},
		{
			"Ordinal 21",
			21,
			elements.NumberFormatOrdinal,
			"21st",
		},
	}

	ns := NewNumberingState(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ns.formatNumber(
				tt.value,
				tt.format,
			)
			if result != tt.expected {
				t.Errorf(
					"formatNumber(%d, %v) = %q, want %q",
					tt.value,
					tt.format,
					result,
					tt.expected,
				)
			}
		})
	}
}
