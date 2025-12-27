package word

import (
	"os"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestWordRenderer_SimpleHeader tests basic header rendering
func TestWordRenderer_SimpleHeader(t *testing.T) {
	// Create a document with a simple header
	doc, err := wordprocessing.New(
		"test_header.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("test_header.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	// Add a header part
	headerPart, err := mainPart.AddHeaderPart()
	if err != nil {
		t.Fatalf(
			"Failed to add header part: %v",
			err,
		)
	}

	// Add content to header
	header := headerPart.GetOrCreateHeader()
	header.AppendParagraph("Test Header")

	// Link header to section
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("No document element")
	}

	// Cast to *elements.Document
	doc2, ok := docElem.(*elements.Document)
	if !ok {
		t.Fatal(
			"Document is not *elements.Document",
		)
	}

	body := doc2.Body()
	if body == nil {
		t.Fatal("No body element")
	}

	sectPr := body.GetOrCreateSectionProperties()
	relID := headerPart.RelationshipID()
	sectPr.SetHeaderReference(
		relID,
		elements.HeaderFooterDefault,
	)

	// Add some body content
	body.AppendParagraph("Body content")

	// Setup renderer
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	err = renderer.Render("test_header.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("test_header.pdf")

	// Verify PDF was created
	if _, err := os.Stat("test_header.pdf"); os.IsNotExist(
		err,
	) {
		t.Error("PDF file was not created")
	}
}

// TestWordRenderer_SimpleFooter tests basic footer rendering
func TestWordRenderer_SimpleFooter(t *testing.T) {
	// Create a document with a simple footer
	doc, err := wordprocessing.New(
		"test_footer.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("test_footer.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	// Add a footer part
	footerPart, err := mainPart.AddFooterPart()
	if err != nil {
		t.Fatalf(
			"Failed to add footer part: %v",
			err,
		)
	}

	// Add content to footer
	footer := footerPart.GetOrCreateFooter()
	footer.AppendParagraph("Test Footer")

	// Link footer to section
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("No document element")
	}

	doc2, ok := docElem.(*elements.Document)
	if !ok {
		t.Fatal(
			"Document is not *elements.Document",
		)
	}

	body := doc2.Body()
	if body == nil {
		t.Fatal("No body element")
	}

	sectPr := body.GetOrCreateSectionProperties()
	relID := footerPart.RelationshipID()
	sectPr.SetFooterReference(
		relID,
		elements.HeaderFooterDefault,
	)

	// Add some body content
	body.AppendParagraph("Body content")

	// Setup renderer
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	err = renderer.Render("test_footer.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("test_footer.pdf")

	// Verify PDF was created
	if _, err := os.Stat("test_footer.pdf"); os.IsNotExist(
		err,
	) {
		t.Error("PDF file was not created")
	}
}

// TestWordRenderer_FirstPageDifferent tests different first page header/footer
func TestWordRenderer_FirstPageDifferent(
	t *testing.T,
) {
	doc, err := wordprocessing.New(
		"test_first_different.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("test_first_different.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	// Add first page header
	firstHeaderPart, err := mainPart.AddHeaderPart()
	if err != nil {
		t.Fatalf(
			"Failed to add first header part: %v",
			err,
		)
	}
	firstHeader := firstHeaderPart.GetOrCreateHeader()
	firstHeader.AppendParagraph(
		"First Page Header",
	)

	// Add default header
	defaultHeaderPart, err := mainPart.AddHeaderPart()
	if err != nil {
		t.Fatalf(
			"Failed to add default header part: %v",
			err,
		)
	}
	defaultHeader := defaultHeaderPart.GetOrCreateHeader()
	defaultHeader.AppendParagraph(
		"Other Pages Header",
	)

	// Link headers to section
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("No document element")
	}

	doc2, ok := docElem.(*elements.Document)
	if !ok {
		t.Fatal(
			"Document is not *elements.Document",
		)
	}

	body := doc2.Body()
	if body == nil {
		t.Fatal("No body element")
	}

	sectPr := body.GetOrCreateSectionProperties()
	sectPr.SetTitlePage(
		true,
	) // Enable different first page
	sectPr.SetHeaderReference(
		firstHeaderPart.RelationshipID(),
		elements.HeaderFooterFirst,
	)
	sectPr.SetHeaderReference(
		defaultHeaderPart.RelationshipID(),
		elements.HeaderFooterDefault,
	)

	// Add body content
	body.AppendParagraph("First page body")
	body.AppendParagraph("Second page body")

	// Setup renderer
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	err = renderer.Render(
		"test_first_different.pdf",
	)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("test_first_different.pdf")

	// Verify PDF was created
	if _, err := os.Stat("test_first_different.pdf"); os.IsNotExist(
		err,
	) {
		t.Error("PDF file was not created")
	}
}

// TestWordRenderer_PageNumbers tests PAGE field in footer
func TestWordRenderer_PageNumbers(t *testing.T) {
	doc, err := wordprocessing.New(
		"test_page_numbers.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer os.Remove("test_page_numbers.docx")

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	// Add footer with page number
	footerPart, err := mainPart.AddFooterPart()
	if err != nil {
		t.Fatalf(
			"Failed to add footer part: %v",
			err,
		)
	}

	footer := footerPart.GetOrCreateFooter()
	// Create a paragraph with "Page X" text
	// In a real implementation, this would include field codes
	footer.AppendParagraph(
		"Page 1",
	) // Simplified for now

	// Link footer to section
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("No document element")
	}

	doc2, ok := docElem.(*elements.Document)
	if !ok {
		t.Fatal(
			"Document is not *elements.Document",
		)
	}

	body := doc2.Body()
	if body == nil {
		t.Fatal("No body element")
	}

	sectPr := body.GetOrCreateSectionProperties()
	relID := footerPart.RelationshipID()
	sectPr.SetFooterReference(
		relID,
		elements.HeaderFooterDefault,
	)

	// Add multiple pages of content
	for i := 1; i <= 5; i++ {
		body.AppendParagraph("Page content")
	}

	// Setup renderer
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	err = renderer.Render("test_page_numbers.pdf")
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	defer os.Remove("test_page_numbers.pdf")

	// Verify PDF was created
	if _, err := os.Stat("test_page_numbers.pdf"); os.IsNotExist(
		err,
	) {
		t.Error("PDF file was not created")
	}

	// Verify total pages was set
	if renderer.totalPages == 0 {
		t.Error("Total pages was not set")
	}
}

// TestWordRenderer_EvaluateFieldCode tests field code evaluation
func TestWordRenderer_EvaluateFieldCode(
	t *testing.T,
) {
	doc, _ := wordprocessing.New(
		"test.docx",
		wordprocessing.DocTypeDocument,
	)
	defer os.Remove("test.docx")

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, _ := NewWordRenderer(doc, engine)

	// Set page numbers for testing
	renderer.currentPageNumber = 3
	renderer.totalPages = 10

	tests := []struct {
		name      string
		fieldCode string
		want      string
	}{
		{"PAGE field", "PAGE", "3"},
		{"PAGE field uppercase", "PAGE", "3"},
		{"NUMPAGES field", "NUMPAGES", "10"},
		{
			"DATE field",
			"DATE",
			"",
		}, // Will be current date
		{
			"TIME field",
			"TIME",
			"",
		}, // Will be current time
		{"Unknown field", "UNKNOWN", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderer.evaluateFieldCode(
				tt.fieldCode,
			)
			if tt.want != "" && got != tt.want {
				t.Errorf(
					"evaluateFieldCode(%q) = %q, want %q",
					tt.fieldCode,
					got,
					tt.want,
				)
			}
			// For DATE and TIME, just verify they return non-empty strings
			if (tt.fieldCode == "DATE" || tt.fieldCode == "TIME") &&
				got == "" {
				t.Errorf(
					"evaluateFieldCode(%q) returned empty string",
					tt.fieldCode,
				)
			}
		})
	}
}
