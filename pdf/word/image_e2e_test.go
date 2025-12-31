// image_e2e_test.go contains end-to-end tests for Word image to PDF rendering.

package word

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

// TestImageRendering_BasicInlineImage tests rendering a Word document with an inline image to PDF.
// This is an end-to-end test that validates the image rendering integration.
func TestImageRendering_BasicInlineImage(
	t *testing.T,
) {
	// Create a test document with an inline image
	tmpDir := t.TempDir()
	docxPath := filepath.Join(
		tmpDir,
		"image_test.docx",
	)

	doc, err := wordprocessing.New(
		docxPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			t.Fatalf(
				"Failed to add main part: %v",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		t.Fatalf(
			"Failed to reload main part: %v",
			err,
		)
	}

	// Get the body element
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Failed to get document element")
	}
	body := docElem.Body()
	if body == nil {
		t.Fatal("Failed to get body element")
	}

	// Add a paragraph with text
	body.AppendParagraph(
		"Document with inline image:",
	)

	// Create a paragraph with an inline drawing
	p := body.AppendParagraph("")
	r := p.AppendRun("")

	// Create a simple 1x1 PNG image (smallest valid PNG)
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG signature
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52, // IHDR chunk
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, // 1x1
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41, // IDAT chunk
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0xB4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, // IEND chunk
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}

	// Add the image part
	imagePart, err := mainPart.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		t.Fatalf(
			"Failed to add image part: %v",
			err,
		)
	}

	// Write image data
	imagePart.FeedDataBytes(pngData)

	// Get the relationship ID
	relID := imagePart.RelationshipID()

	// Create inline drawing with 2 inch x 2 inch size
	width := int64(2 * elements.EMUsPerInch)
	height := int64(2 * elements.EMUsPerInch)
	drawing := elements.NewInlineDrawing(
		width,
		height,
		relID,
	)
	r.AppendChild(drawing)

	// Add another paragraph after the image
	body.AppendParagraph("Text after the image.")

	// Save the document
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	// Close and reopen for rendering
	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Reopen the document for rendering
	doc2, err := wordprocessing.Open(
		docxPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	// Setup renderer
	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc2, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	// Render to PDF
	pdfPath := filepath.Join(
		tmpDir,
		"image_test.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}

	// Verify PDF has content (non-zero size)
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	// Basic size check - PDF should be at least a few hundred bytes
	// (even a minimal PDF with empty page is around 600+ bytes)
	if stat.Size() < 500 {
		t.Errorf(
			"PDF file is suspiciously small (%d bytes), expected at least 500 bytes",
			stat.Size(),
		)
	}

	t.Logf("PDF size: %d bytes", stat.Size())
}

// TestImageRendering_MultipleImages tests rendering a document with multiple images.
// This validates that image rendering works correctly with multiple images.
func TestImageRendering_MultipleImages(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	docxPath := filepath.Join(
		tmpDir,
		"multi_image.docx",
	)

	doc, err := wordprocessing.New(
		docxPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			t.Fatalf(
				"Failed to add main part: %v",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		t.Fatalf(
			"Failed to reload main part: %v",
			err,
		)
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Failed to get document element")
	}
	body := docElem.Body()
	if body == nil {
		t.Fatal("Failed to get body element")
	}

	// Simple 1x1 PNG image data
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0xB4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E,
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}

	// Add 3 images with different sizes
	sizes := []struct {
		width  float64
		height float64
		text   string
	}{
		{1.0, 1.0, "Small image (1x1 inch)"},
		{2.0, 1.5, "Medium image (2x1.5 inches)"},
		{3.0, 2.0, "Large image (3x2 inches)"},
	}

	for i, size := range sizes {
		// Add descriptive text
		body.AppendParagraph(size.text)

		// Create paragraph with image
		p := body.AppendParagraph("")
		r := p.AppendRun("")

		// Add image part
		imagePart, err := mainPart.AddImagePart(
			parts.ImageTypePng,
		)
		if err != nil {
			t.Fatalf(
				"Failed to add image part %d: %v",
				i+1,
				err,
			)
		}

		imagePart.FeedDataBytes(pngData)

		relID := imagePart.RelationshipID()
		width := int64(
			size.width * float64(
				elements.EMUsPerInch,
			),
		)
		height := int64(
			size.height * float64(
				elements.EMUsPerInch,
			),
		)
		drawing := elements.NewInlineDrawing(
			width,
			height,
			relID,
		)
		r.AppendChild(drawing)

		// Add spacing between images
		body.AppendParagraph("")
	}

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := wordprocessing.Open(
		docxPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc2, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(
		tmpDir,
		"multi_image.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	t.Logf(
		"PDF size for 3 images: %d bytes",
		stat.Size(),
	)
}

// TestImageRendering_EmptyDocument tests rendering an empty document.
// This validates graceful handling when no images are present.
func TestImageRendering_EmptyDocument(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	docxPath := filepath.Join(
		tmpDir,
		"empty.docx",
	)

	doc, err := wordprocessing.New(
		docxPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			t.Fatalf(
				"Failed to add main part: %v",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		t.Fatalf(
			"Failed to reload main part: %v",
			err,
		)
	}

	// Add just a simple paragraph
	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Failed to get document element")
	}
	body := docElem.Body()
	if body == nil {
		t.Fatal("Failed to get body element")
	}

	body.AppendParagraph(
		"Empty document with no images.",
	)

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := wordprocessing.Open(
		docxPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc2, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "empty.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

// TestImageRendering_TextAndImagesMixed tests rendering a document with mixed text and images.
// This validates that image rendering works correctly alongside text content.
func TestImageRendering_TextAndImagesMixed(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	docxPath := filepath.Join(
		tmpDir,
		"mixed.docx",
	)

	doc, err := wordprocessing.New(
		docxPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			t.Fatalf(
				"Failed to add main part: %v",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		t.Fatalf(
			"Failed to reload main part: %v",
			err,
		)
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Failed to get document element")
	}
	body := docElem.Body()
	if body == nil {
		t.Fatal("Failed to get body element")
	}

	// Simple PNG data
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0xB4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E,
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}

	// Add title
	p1 := body.AppendParagraph("")
	r1 := p1.AppendRun("Document Title")
	r1.SetBold(true)
	r1.SetFontSize(
		28,
	) // Font size in half-points (28 = 14pt)

	// Add introduction paragraph
	body.AppendParagraph(
		"This is a document that contains both text and images. The images should render correctly alongside the text content.",
	)

	// Add first image
	p2 := body.AppendParagraph("")
	r2 := p2.AppendRun("")
	imagePart1, err := mainPart.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		t.Fatalf(
			"Failed to add image part 1: %v",
			err,
		)
	}
	imagePart1.FeedDataBytes(pngData)
	relID1 := imagePart1.RelationshipID()
	drawing1 := elements.NewInlineDrawing(
		int64(1.5*float64(elements.EMUsPerInch)),
		int64(1.5*float64(elements.EMUsPerInch)),
		relID1,
	)
	r2.AppendChild(drawing1)

	// Add middle paragraph
	body.AppendParagraph(
		"Here is some text in between two images. This tests that the layout engine properly handles mixed content.",
	)

	// Add second image
	p3 := body.AppendParagraph("")
	r3 := p3.AppendRun("")
	imagePart2, err := mainPart.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		t.Fatalf(
			"Failed to add image part 2: %v",
			err,
		)
	}
	imagePart2.FeedDataBytes(pngData)
	relID2 := imagePart2.RelationshipID()
	drawing2 := elements.NewInlineDrawing(
		int64(2.0*float64(elements.EMUsPerInch)),
		int64(1.0*float64(elements.EMUsPerInch)),
		relID2,
	)
	r3.AppendChild(drawing2)

	// Add conclusion
	body.AppendParagraph(
		"This is the end of the document. All content should be rendered correctly in the PDF.",
	)

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := wordprocessing.Open(
		docxPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc2, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(tmpDir, "mixed.pdf")
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}
	if stat.Size() < 500 {
		t.Errorf(
			"PDF file is suspiciously small (%d bytes), expected at least 500 bytes",
			stat.Size(),
		)
	}

	t.Logf(
		"PDF size for mixed content: %d bytes",
		stat.Size(),
	)
}

// TestImageRendering_LargeImage tests image rendering with a larger image size.
// This validates that the renderer can handle images of various sizes.
func TestImageRendering_LargeImage(t *testing.T) {
	tmpDir := t.TempDir()
	docxPath := filepath.Join(
		tmpDir,
		"large_image.docx",
	)

	doc, err := wordprocessing.New(
		docxPath,
		wordprocessing.DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			t.Errorf(
				"Failed to close document: %v",
				err,
			)
		}
	}()

	mainPart := doc.MainPart()
	if mainPart == nil {
		mainPart, err = doc.AddMainPart()
		if err != nil {
			t.Fatalf(
				"Failed to add main part: %v",
				err,
			)
		}
	}
	mainPart.InitializeContent()
	if err := mainPart.Reload(); err != nil {
		t.Fatalf(
			"Failed to reload main part: %v",
			err,
		)
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Failed to get document element")
	}
	body := docElem.Body()
	if body == nil {
		t.Fatal("Failed to get body element")
	}

	// Simple PNG data
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0xB4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E,
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}

	body.AppendParagraph(
		"Document with large image (5x4 inches)",
	)

	// Add a large image (5 inches wide, 4 inches tall)
	p := body.AppendParagraph("")
	r := p.AppendRun("")
	imagePart, err := mainPart.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		t.Fatalf(
			"Failed to add image part: %v",
			err,
		)
	}
	imagePart.FeedDataBytes(pngData)
	relID := imagePart.RelationshipID()
	drawing := elements.NewInlineDrawing(
		int64(5.0*float64(elements.EMUsPerInch)),
		int64(4.0*float64(elements.EMUsPerInch)),
		relID,
	)
	r.AppendChild(drawing)

	body.AppendParagraph(
		"Text after the large image.",
	)

	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}

	if err := doc.Close(); err != nil {
		t.Fatalf(
			"Failed to close document: %v",
			err,
		)
	}

	// Render to PDF
	doc2, err := wordprocessing.Open(
		docxPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open document: %v",
			err,
		)
	}
	defer func() {
		if err := doc2.Close(); err != nil {
			t.Errorf(
				"Failed to close reopened document: %v",
				err,
			)
		}
	}()

	fc := font.NewFontCache(10)
	engine := layout.NewTextLayoutEngine(fc)
	renderer, err := NewWordRenderer(doc2, engine)
	if err != nil {
		t.Fatalf(
			"Failed to create renderer: %v",
			err,
		)
	}

	pdfPath := filepath.Join(
		tmpDir,
		"large_image.pdf",
	)
	if err := renderer.Render(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	t.Logf(
		"PDF size with large image: %d bytes",
		stat.Size(),
	)
}
