package word

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestFieldRendering tests PAGE and NUMPAGES field rendering in document body.
func TestFieldRendering(t *testing.T) {
	// Create a new document with PAGE and NUMPAGES fields
	doc, err := wordprocessing.New(
		"test_fields.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_fields.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add first paragraph with PAGE field
	p1 := elements.NewParagraph()
	run1 := elements.NewRun("This is page ")
	p1.AppendChild(run1)

	// Add a run with instrText for PAGE field
	runField := elements.NewRun("")
	instrElem := elements.NewText("PAGE")
	// Note: In actual field implementation, we would set instrText attribute
	// For now, we just add the text node
	runField.AppendChild(instrElem)
	p1.AppendChild(runField)

	run1b := elements.NewRun(" of ")
	p1.AppendChild(run1b)

	// Add a run with instrText for NUMPAGES field
	runFieldNum := elements.NewRun("")
	instrElemNum := elements.NewText("NUMPAGES")
	runFieldNum.AppendChild(instrElemNum)
	p1.AppendChild(runFieldNum)

	body.AppendChild(p1)

	// Add page break
	p2 := elements.NewParagraph()
	p2.SetPageBreakBefore(true)
	p2.AppendRun("Second page content")
	body.AppendChild(p2)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"fields.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Check that totalPages was set
	if renderer.totalPages == 0 {
		t.Errorf(
			"Expected totalPages to be set, got 0",
		)
	}
}

// TestHyperlinkRendering tests hyperlink rendering with external and internal links.
func TestHyperlinkRendering(t *testing.T) {
	doc, err := wordprocessing.New(
		"test_hyperlinks.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_hyperlinks.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add paragraph with external hyperlink
	p1 := elements.NewParagraph()
	hyperlink := elements.NewHyperlink(
		"Visit Example",
		"rId1",
	)
	p1.AppendChild(hyperlink)
	body.AppendChild(p1)

	// Add paragraph with internal hyperlink (anchor)
	p2 := elements.NewParagraph()
	internalLink := elements.NewInternalHyperlink(
		"Go to Section 1",
		"section1",
	)
	p2.AppendChild(internalLink)
	body.AppendChild(p2)

	// Add bookmarked section
	p3 := elements.NewParagraph()
	bookmarkStart := elements.NewBookmarkStart(
		1,
		"section1",
	)
	p3.AppendChild(bookmarkStart)
	p3.AppendRun("This is Section 1")
	bookmarkEnd := elements.NewBookmarkEnd(1)
	p3.AppendChild(bookmarkEnd)
	body.AppendChild(p3)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"hyperlinks.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Verify bookmark was tracked
	if _, found := renderer.bookmarks["section1"]; !found {
		t.Errorf(
			"Expected bookmark 'section1' to be tracked",
		)
	}
}

// TestBookmarkRendering tests bookmark start and end markers.
func TestBookmarkRendering(t *testing.T) {
	doc, err := wordprocessing.New(
		"test_bookmarks.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_bookmarks.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add paragraph with bookmarks
	p1 := elements.NewParagraph()

	// Bookmark 1
	bs1 := elements.NewBookmarkStart(
		1,
		"bookmark1",
	)
	p1.AppendChild(bs1)
	p1.AppendRun("Bookmarked text 1")
	be1 := elements.NewBookmarkEnd(1)
	p1.AppendChild(be1)

	body.AppendChild(p1)

	// Add another paragraph with different bookmark
	p2 := elements.NewParagraph()
	bs2 := elements.NewBookmarkStart(
		2,
		"bookmark2",
	)
	p2.AppendChild(bs2)
	p2.AppendRun("Bookmarked text 2")
	be2 := elements.NewBookmarkEnd(2)
	p2.AppendChild(be2)
	body.AppendChild(p2)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"bookmarks.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Verify both bookmarks were tracked
	if len(renderer.bookmarks) != 2 {
		t.Errorf(
			"Expected 2 bookmarks, got %d",
			len(renderer.bookmarks),
		)
	}

	if _, found := renderer.bookmarks["bookmark1"]; !found {
		t.Errorf(
			"Expected bookmark 'bookmark1' to be tracked",
		)
	}

	if _, found := renderer.bookmarks["bookmark2"]; !found {
		t.Errorf(
			"Expected bookmark 'bookmark2' to be tracked",
		)
	}
}

// TestFootnoteRendering tests footnote reference and content rendering.
func TestFootnoteRendering(t *testing.T) {
	// This test requires a fully constructed document with footnotes part
	// For now, we'll test the basic structure

	doc, err := wordprocessing.New(
		"test_footnotes.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_footnotes.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	// Create footnotes part (in real usage, this would be done by the document loader)
	// For this test, we'll verify the renderer can handle footnote references

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add paragraph with footnote reference
	p1 := elements.NewParagraph()
	p1.AppendRun("This text has a footnote")

	// Add footnote reference
	run := elements.NewRun("")
	footnoteRef := elements.NewFootnoteReference(
		1,
	)
	run.AppendChild(footnoteRef)
	p1.AppendChild(run)

	body.AppendChild(p1)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"footnotes.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Note: Actual footnote content rendering requires a valid footnotes part
	// which would be populated when loading a real document
}

// TestEndnoteRendering tests endnote reference and content rendering.
func TestEndnoteRendering(t *testing.T) {
	doc, err := wordprocessing.New(
		"test_endnotes.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_endnotes.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	docElem := mainPart.Document()
	body := docElem.Body()

	// Add paragraph with endnote reference
	p1 := elements.NewParagraph()
	p1.AppendRun("This text has an endnote")

	// Add endnote reference
	run := elements.NewRun("")
	endnoteRef := elements.NewEndnoteReference(1)
	run.AppendChild(endnoteRef)
	p1.AppendChild(run)

	body.AppendChild(p1)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"endnotes.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Note: Actual endnote content rendering requires a valid endnotes part
	// which would be populated when loading a real document
}

// TestEvaluateFieldCode tests the field code evaluation function.
func TestEvaluateFieldCode(t *testing.T) {
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	doc, err := wordprocessing.New(
		"test_eval.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_eval.docx") }()

	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Set page number for testing
	renderer.currentPageNumber = 3
	renderer.totalPages = 10

	tests := []struct {
		name      string
		fieldCode string
		want      string
	}{
		{
			name:      "PAGE field",
			fieldCode: "PAGE",
			want:      "3",
		},
		{
			name:      "PAGE field with formatting",
			fieldCode: "PAGE \\* MERGEFORMAT",
			want:      "3",
		},
		{
			name:      "NUMPAGES field",
			fieldCode: "NUMPAGES",
			want:      "10",
		},
		{
			name:      "DATE field",
			fieldCode: "DATE",
			want:      "", // Will return current date, hard to test exactly
		},
		{
			name:      "TIME field",
			fieldCode: "TIME",
			want:      "", // Will return current time, hard to test exactly
		},
		{
			name:      "Unknown field",
			fieldCode: "UNKNOWN",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderer.evaluateFieldCode(
				tt.fieldCode,
			)
			if tt.name == "DATE field" ||
				tt.name == "TIME field" {
				// Just verify it's not empty
				if got == "" {
					t.Errorf(
						"evaluateFieldCode(%q) returned empty string",
						tt.fieldCode,
					)
				}
			} else if got != tt.want {
				t.Errorf("evaluateFieldCode(%q) = %q, want %q", tt.fieldCode, got, tt.want)
			}
		})
	}
}

// TestFieldsInHeadersFooters verifies that fields work in headers and footers.
func TestFieldsInHeadersFooters(t *testing.T) {
	// This test verifies that PAGE and NUMPAGES fields work in headers/footers
	// The actual implementation was already tested in renderer_header_footer_test.go
	// This is just a placeholder to document the requirement

	t.Skip(
		"Fields in headers/footers are tested in renderer_header_footer_test.go",
	)
}

// TestComplexFieldScenarios tests complex scenarios with multiple field types.
func TestComplexFieldScenarios(t *testing.T) {
	doc, err := wordprocessing.New(
		"test_complex.docx",
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = os.Remove("test_complex.docx") }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, _ = doc.AddMainPart()
	}
	mainPart.InitializeContent()
	_ = mainPart.Reload()

	docElem := mainPart.Document()
	body := docElem.Body()

	// Create a complex document with multiple special content types
	p1 := elements.NewParagraph()
	p1.AppendRun(
		"Document with multiple features: ",
	)

	// Add hyperlink
	hyperlink := elements.NewHyperlink(
		"Link",
		"rId1",
	)
	p1.AppendChild(hyperlink)

	// Add bookmark
	bs := elements.NewBookmarkStart(1, "section1")
	p1.AppendChild(bs)
	p1.AppendRun(" bookmarked ")
	be := elements.NewBookmarkEnd(1)
	p1.AppendChild(be)

	body.AppendChild(p1)

	// Add second page with fields
	p2 := elements.NewParagraph()
	p2.SetPageBreakBefore(true)
	p2.AppendRun("Page ")

	// Add PAGE field
	runField := elements.NewRun("")
	instrElem := elements.NewText("PAGE")
	runField.AppendChild(instrElem)
	p2.AppendChild(runField)

	body.AppendChild(p2)

	// Create renderer
	cache := font.NewFontCache(100)
	engine := layout.NewTextLayoutEngine(cache)
	renderer, err := NewWordRenderer(doc, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"complex.pdf",
	)
	err = renderer.Render(outputPath)
	if err != nil {
		t.Fatalf(
			"Failed to render document: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Errorf("Output PDF was not created")
	}

	// Verify bookmark was tracked
	if _, found := renderer.bookmarks["section1"]; !found {
		t.Errorf(
			"Expected bookmark 'section1' to be tracked",
		)
	}

	// Verify page count - with the page break fixed, we should have at least 2 pages
	if renderer.totalPages < 2 {
		t.Errorf(
			"Expected at least 2 pages due to page break, got %d",
			renderer.totalPages,
		)
	}
}
