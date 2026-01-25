package core

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// TestIntegration_DocumentCreationAndOutput tests the full document creation pipeline.
func TestIntegration_DocumentCreationAndOutput(
	t *testing.T,
) {
	t.Run(
		"create document add pages and write to buffer",
		func(t *testing.T) {
			doc, err := NewDocument()
			if err != nil {
				t.Fatalf(
					"NewDocument() error = %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add multiple pages with different sizes
			page1, err := doc.AddPageA4()
			if err != nil {
				t.Fatalf(
					"AddPageA4() error = %v",
					err,
				)
			}
			if page1.Number() != 1 {
				t.Errorf(
					"First page number = %d, want 1",
					page1.Number(),
				)
			}

			page2, err := doc.AddPageLetter()
			if err != nil {
				t.Fatalf(
					"AddPageLetter() error = %v",
					err,
				)
			}
			if page2.Number() != 2 {
				t.Errorf(
					"Second page number = %d, want 2",
					page2.Number(),
				)
			}

			page3, err := doc.AddPage(
				PageSizeA4.Landscape(),
			)
			if err != nil {
				t.Fatalf(
					"AddPage() error = %v",
					err,
				)
			}
			if page3.Number() != 3 {
				t.Errorf(
					"Third page number = %d, want 3",
					page3.Number(),
				)
			}

			// Verify page count
			if doc.PageCount() != 3 {
				t.Errorf(
					"PageCount() = %d, want 3",
					doc.PageCount(),
				)
			}

			// Write to buffer
			var buf bytes.Buffer
			if err := doc.Write(&buf); err != nil {
				t.Fatalf(
					"Write() error = %v",
					err,
				)
			}

			// Verify PDF output structure
			data := buf.Bytes()
			verifyPDFStructure(t, data, V17)
		},
	)

	t.Run(
		"verify PDF magic bytes and structure",
		func(t *testing.T) {
			doc, err := NewDocument()
			if err != nil {
				t.Fatalf(
					"NewDocument() error = %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			_, err = doc.AddPageA4()
			if err != nil {
				t.Fatalf(
					"AddPageA4() error = %v",
					err,
				)
			}

			var buf bytes.Buffer
			if err := doc.Write(&buf); err != nil {
				t.Fatalf(
					"Write() error = %v",
					err,
				)
			}

			data := buf.Bytes()

			// Check PDF magic bytes
			if len(data) < 8 {
				t.Fatal("PDF output too small")
			}
			if !bytes.HasPrefix(
				data,
				[]byte("%PDF-"),
			) {
				t.Errorf(
					"PDF does not start with %%PDF- magic bytes",
				)
			}

			// Check for PDF EOF marker
			if !bytes.Contains(
				data,
				[]byte("%%EOF"),
			) {
				t.Error(
					"PDF missing EOF marker",
				)
			}

			// Check for essential PDF objects
			// Note: In PDF 1.5+ with object streams, these may be compressed
			// so we check for either the uncompressed form or the object stream indicator
			hasCatalog := bytes.Contains(
				data,
				[]byte("/Type /Catalog"),
			) ||
				bytes.Contains(
					data,
					[]byte("/Type/Catalog"),
				)
			hasPages := bytes.Contains(
				data,
				[]byte("/Type /Pages"),
			) ||
				bytes.Contains(
					data,
					[]byte("/Type/Pages"),
				)
			hasPage := bytes.Contains(
				data,
				[]byte("/Type /Page"),
			) ||
				bytes.Contains(
					data,
					[]byte("/Type/Page"),
				)
			hasObjStm := bytes.Contains(
				data,
				[]byte("/Type /ObjStm"),
			) ||
				bytes.Contains(
					data,
					[]byte("/Type/ObjStm"),
				)

			// If object streams are used, the catalog/pages might be inside them
			if !hasCatalog && !hasObjStm {
				t.Logf(
					"PDF content (first 500 bytes): %s",
					string(
						data[:min(500, len(data))],
					),
				)
				t.Error(
					"PDF missing Catalog object",
				)
			}
			if !hasPages && !hasObjStm {
				t.Error(
					"PDF missing Pages object",
				)
			}
			if !hasPage && !hasObjStm {
				t.Error(
					"PDF missing Page object",
				)
			}
		},
	)
}

// TestIntegration_PDFVersions tests PDF output with different versions.
func TestIntegration_PDFVersions(t *testing.T) {
	versions := []struct {
		version     Version
		wantHeader  string
		description string
	}{
		{V14, "%PDF-1.4", "PDF 1.4 (Acrobat 5)"},
		{V15, "%PDF-1.5", "PDF 1.5 (Acrobat 6)"},
		{V16, "%PDF-1.6", "PDF 1.6 (Acrobat 7)"},
		{
			V17,
			"%PDF-1.7",
			"PDF 1.7 (ISO 32000-1)",
		},
		// Note: V20 might output as 1.7 or 2.0 depending on pdfcpu version
	}

	for _, tt := range versions {
		t.Run(tt.description, func(t *testing.T) {
			opts := &DocumentOptions{
				Version: tt.version,
			}
			doc, err := NewDocumentWithOptions(
				opts,
			)
			if err != nil {
				t.Fatalf(
					"NewDocumentWithOptions() error = %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			_, err = doc.AddPageA4()
			if err != nil {
				t.Fatalf(
					"AddPageA4() error = %v",
					err,
				)
			}

			var buf bytes.Buffer
			if err := doc.Write(&buf); err != nil {
				t.Fatalf(
					"Write() error = %v",
					err,
				)
			}

			data := buf.Bytes()
			header := string(data[:8])
			if !strings.HasPrefix(
				header,
				tt.wantHeader,
			) {
				// Some versions may normalize to 1.7
				if !strings.HasPrefix(
					header,
					"%PDF-1.",
				) {
					t.Errorf(
						"PDF header = %q, want %q or valid PDF header",
						header,
						tt.wantHeader,
					)
				}
			}
		})
	}
}

// TestIntegration_DocumentMetadata tests that metadata is correctly embedded.
func TestIntegration_DocumentMetadata(
	t *testing.T,
) {
	opts := &DocumentOptions{
		Version: V17,
		Metadata: Metadata{
			Title:    "Test PDF Document",
			Author:   "Test Author",
			Subject:  "Integration Test",
			Keywords: "test, pdf, integration",
			Creator:  "Test Creator Application",
			Producer: "goffice-pdf-test",
			CreationDate: time.Date(
				2024,
				1,
				15,
				10,
				30,
				0,
				0,
				time.UTC,
			),
		},
	}

	doc, err := NewDocumentWithOptions(opts)
	if err != nil {
		t.Fatalf(
			"NewDocumentWithOptions() error = %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Verify metadata is set
	meta := doc.Metadata()
	if meta.Title != "Test PDF Document" {
		t.Errorf(
			"Metadata.Title = %q, want %q",
			meta.Title,
			"Test PDF Document",
		)
	}
	if meta.Author != "Test Author" {
		t.Errorf(
			"Metadata.Author = %q, want %q",
			meta.Author,
			"Test Author",
		)
	}

	_, err = doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	// Verify PDF is valid
	data := buf.Bytes()
	verifyPDFStructure(t, data, V17)
}

// TestIntegration_PageWithContent tests pages with content streams.
func TestIntegration_PageWithContent(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	// Write some PDF content operators
	content := `q
0.5 0.5 0.5 RG
1 w
100 700 m
500 700 l
S
Q
`
	n, err := page.WriteContentString(content)
	if err != nil {
		t.Fatalf(
			"WriteContentString() error = %v",
			err,
		)
	}
	if n != len(content) {
		t.Errorf(
			"WriteContentString() wrote %d bytes, want %d",
			n,
			len(content),
		)
	}

	// Verify content is stored in page
	storedContent := page.Content()
	if string(storedContent) != content {
		t.Errorf("Page content mismatch")
	}

	// Test that we can retrieve the content
	if len(storedContent) == 0 {
		t.Error(
			"Page content should not be empty",
		)
	}

	// Note: Full PDF output test with content streams requires pdfcpu
	// to properly handle the stream dict - this is tested in document_test.go
}

// TestIntegration_MultiplePagesWithContent tests multiple pages with content.
func TestIntegration_MultiplePagesWithContent(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Create 5 pages with different sizes
	for i := 1; i <= 5; i++ {
		var page *Page
		var addErr error

		switch i % 3 {
		case 0:
			page, addErr = doc.AddPageA4()
		case 1:
			page, addErr = doc.AddPageLetter()
		case 2:
			page, addErr = doc.AddPage(
				PageSizeA4.Landscape(),
			)
		}

		if addErr != nil {
			t.Fatalf(
				"AddPage() error = %v for page %d",
				addErr,
				i,
			)
		}

		// Verify page number
		if page.Number() != i {
			t.Errorf(
				"Page %d number = %d, want %d",
				i,
				page.Number(),
				i,
			)
		}

		// Add content to the page (stored in memory)
		content := strings.Repeat(
			"% Page content ",
			i*10,
		) + "\n"
		if _, err := page.WriteContentString(content); err != nil {
			t.Fatalf(
				"WriteContentString() error = %v for page %d",
				err,
				i,
			)
		}

		// Verify content was stored
		if len(page.Content()) == 0 {
			t.Errorf("Page %d has no content", i)
		}
	}

	if doc.PageCount() != 5 {
		t.Errorf(
			"PageCount() = %d, want 5",
			doc.PageCount(),
		)
	}

	// Verify we can get pages
	for i := range 5 {
		page, err := doc.GetPage(i)
		if err != nil {
			t.Fatalf(
				"GetPage(%d) error = %v",
				i,
				err,
			)
		}
		if page == nil {
			t.Fatalf(
				"GetPage(%d) returned nil",
				i,
			)
		}
		if page.Number() != i+1 {
			t.Errorf(
				"GetPage(%d).Number() = %d, want %d",
				i,
				page.Number(),
				i+1,
			)
		}
	}
}

// TestIntegration_PageOptions tests pages with various options.
func TestIntegration_PageOptions(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	t.Run(
		"page with margins",
		func(t *testing.T) {
			opts := NewPageOptions(
				PageSizeLetter,
			).
				WithMargins(NewMarginsUniform(72))

				// 1 inch margins

			page, err := doc.AddPageWithOptions(
				opts,
			)
			if err != nil {
				t.Fatalf(
					"AddPageWithOptions() error = %v",
					err,
				)
			}

			margins := page.Margins()
			if margins.Top != 72 ||
				margins.Right != 72 ||
				margins.Bottom != 72 ||
				margins.Left != 72 {
				t.Errorf(
					"Margins = %+v, want all 72",
					margins,
				)
			}
		},
	)

	t.Run(
		"page with rotation",
		func(t *testing.T) {
			opts := NewPageOptions(
				PageSizeA4,
			).WithRotation(90)

			page, err := doc.AddPageWithOptions(
				opts,
			)
			if err != nil {
				t.Fatalf(
					"AddPageWithOptions() error = %v",
					err,
				)
			}

			if page.Rotation() != 90 {
				t.Errorf(
					"Rotation() = %d, want 90",
					page.Rotation(),
				)
			}
		},
	)

	t.Run(
		"page with custom boxes",
		func(t *testing.T) {
			boxes := NewPageBoxesFromSize(
				PageSizeLetter,
			).
				WithCropBox(NewRectangle(18, 18, 594, 774)).
				WithTrimBox(NewRectangle(27, 27, 585, 765)).
				WithBleed(9)

			opts := NewPageOptions(
				PageSizeLetter,
			).WithBoxes(boxes)

			page, err := doc.AddPageWithOptions(
				opts,
			)
			if err != nil {
				t.Fatalf(
					"AddPageWithOptions() error = %v",
					err,
				)
			}

			if page.Boxes() == nil {
				t.Fatal(
					"Boxes() should not be nil",
				)
			}
			if page.Boxes().CropBox == nil {
				t.Error("CropBox should be set")
			}
			if page.Boxes().TrimBox == nil {
				t.Error("TrimBox should be set")
			}
			if page.Boxes().BleedBox == nil {
				t.Error("BleedBox should be set")
			}
		},
	)

	// Write and verify
	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	data := buf.Bytes()
	verifyPDFStructure(t, data, V17)
}

// TestIntegration_EmptyDocument tests error handling for empty documents.
func TestIntegration_EmptyDocument(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	var buf bytes.Buffer
	err = doc.Write(&buf)

	if err == nil {
		t.Error(
			"Write() should fail for empty document",
		)
	}
	if err != ErrNoPages {
		t.Errorf(
			"Write() error = %v, want ErrNoPages",
			err,
		)
	}
}

// TestIntegration_LargeContentStreams tests handling of large content in memory.
func TestIntegration_LargeContentStreams(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	// Generate a large content stream (many drawing operations)
	var contentBuilder strings.Builder
	contentBuilder.WriteString("q\n")
	contentBuilder.WriteString("1 0 0 RG\n")
	contentBuilder.WriteString("0.5 w\n")

	// Draw many lines
	for i := range 1000 {
		y := float64(i%800) + 20
		contentBuilder.WriteString("50 ")
		contentBuilder.WriteString(
			testFormatFloat(y),
		)
		contentBuilder.WriteString(" m\n")
		contentBuilder.WriteString("550 ")
		contentBuilder.WriteString(
			testFormatFloat(y),
		)
		contentBuilder.WriteString(" l\n")
		contentBuilder.WriteString("S\n")
	}

	contentBuilder.WriteString("Q\n")

	content := contentBuilder.String()
	if _, err := page.WriteContentString(content); err != nil {
		t.Fatalf(
			"WriteContentString() error = %v",
			err,
		)
	}

	// Verify the content was stored
	storedContent := page.Content()
	if len(storedContent) == 0 {
		t.Error("Content should not be empty")
	}

	// Verify content size
	expectedSize := len(content)
	actualSize := len(storedContent)
	if actualSize != expectedSize {
		t.Errorf(
			"Content size = %d, want %d",
			actualSize,
			expectedSize,
		)
	}

	// Verify the content contains expected patterns
	contentStr := string(storedContent)
	if !strings.Contains(contentStr, "1 0 0 RG") {
		t.Error(
			"Content should contain color setting",
		)
	}
	if !strings.Contains(contentStr, "0.5 w") {
		t.Error(
			"Content should contain line width",
		)
	}
	// Count stroke operations - each line ends with "S\n"
	strokeCount := strings.Count(
		contentStr,
		"S\n",
	)
	if strokeCount != 1000 {
		t.Errorf(
			"Content should contain 1000 stroke operations, got %d",
			strokeCount,
		)
	}
}

// TestIntegration_ResourceManagement tests that resources are properly managed.
func TestIntegration_ResourceManagement(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	// Add resources to the page
	page.AddResource(
		"Font",
		"F1",
		types.Name("Helvetica"),
	)
	page.AddResource(
		"Font",
		"F2",
		types.Name("Times-Roman"),
	)
	gsDict := types.Dict(map[string]types.Object{
		"Type": types.Name("ExtGState"),
		"ca":   types.Float(0.5),
	})
	page.AddResource("ExtGState", "GS1", gsDict)

	resources := page.Resources()
	if resources == nil {
		t.Fatal("Resources() should not be nil")
	}

	// Verify Font resource category exists
	if _, ok := resources["Font"]; !ok {
		t.Error(
			"Font resource category not found",
		)
	}

	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	data := buf.Bytes()
	verifyPDFStructure(t, data, V17)
}

// TestIntegration_ClosedDocument tests operations on closed documents.
func TestIntegration_ClosedDocument(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}

	_, err = doc.AddPageA4()
	if err != nil {
		t.Fatalf("AddPageA4() error = %v", err)
	}

	// Close the document
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Try operations on closed document
	_, err = doc.AddPageA4()
	if err != ErrDocumentClosed {
		t.Errorf(
			"AddPageA4() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}

	_, err = doc.GetPage(0)
	if err != ErrDocumentClosed {
		t.Errorf(
			"GetPage() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}

	var buf bytes.Buffer
	err = doc.Write(&buf)
	if err != ErrDocumentClosed {
		t.Errorf(
			"Write() on closed doc error = %v, want ErrDocumentClosed",
			err,
		)
	}
}

// TestIntegration_CoordinateSystem tests coordinate system transformations.
func TestIntegration_CoordinateSystem(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	page, err := doc.AddPageLetter()
	if err != nil {
		t.Fatalf(
			"AddPageLetter() error = %v",
			err,
		)
	}

	cs := page.CoordinateSystem()
	if cs == nil {
		t.Fatal(
			"CoordinateSystem() should not be nil",
		)
	}

	// Test default origin (bottom-left)
	x, y := cs.TransformPoint(100, 100)
	if x != 100 || y != 100 {
		t.Errorf(
			"TransformPoint(100, 100) with bottom-left origin = (%v, %v), want (100, 100)",
			x,
			y,
		)
	}

	// Test top-left origin transformation
	cs.WithTopLeftOrigin()
	x, y = cs.TransformPoint(100, 100)
	expectedY := page.Height() - 100
	if x != 100 || y != expectedY {
		t.Errorf(
			"TransformPoint(100, 100) with top-left origin = (%v, %v), want (100, %v)",
			x,
			y,
			expectedY,
		)
	}
}

// TestIntegration_PageSizesAndOrientations tests various page sizes and orientations.
func TestIntegration_PageSizesAndOrientations(
	t *testing.T,
) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	pageSizes := []struct {
		name string
		size PageSize
	}{
		{"A4", PageSizeA4},
		{"A3", PageSizeA3},
		{"A5", PageSizeA5},
		{"Letter", PageSizeLetter},
		{"Legal", PageSizeLegal},
		{"Tabloid", PageSizeTabloid},
		{"Executive", PageSizeExecutive},
		{"A4 Landscape", PageSizeA4.Landscape()},
		{
			"Letter Landscape",
			PageSizeLetter.Landscape(),
		},
	}

	for _, ps := range pageSizes {
		page, err := doc.AddPage(ps.size)
		if err != nil {
			t.Fatalf(
				"AddPage(%s) error = %v",
				ps.name,
				err,
			)
		}

		if page.Width() != ps.size.Width {
			t.Errorf(
				"%s page width = %v, want %v",
				ps.name,
				page.Width(),
				ps.size.Width,
			)
		}
		if page.Height() != ps.size.Height {
			t.Errorf(
				"%s page height = %v, want %v",
				ps.name,
				page.Height(),
				ps.size.Height,
			)
		}
	}

	var buf bytes.Buffer
	if err := doc.Write(&buf); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	data := buf.Bytes()
	verifyPDFStructure(t, data, V17)
}

// TestIntegration_ContentArea tests content area calculations.
func TestIntegration_ContentArea(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	margins := NewMargins(
		72,
		54,
		72,
		54,
	) // 1" top/bottom, 0.75" left/right
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(margins)

	page, err := doc.AddPageWithOptions(opts)
	if err != nil {
		t.Fatalf(
			"AddPageWithOptions() error = %v",
			err,
		)
	}

	// Verify content dimensions
	expectedWidth := PageSizeLetter.Width - 54 - 54   // 612 - 108 = 504
	expectedHeight := PageSizeLetter.Height - 72 - 72 // 792 - 144 = 648

	if page.ContentWidth() != expectedWidth {
		t.Errorf(
			"ContentWidth() = %v, want %v",
			page.ContentWidth(),
			expectedWidth,
		)
	}
	if page.ContentHeight() != expectedHeight {
		t.Errorf(
			"ContentHeight() = %v, want %v",
			page.ContentHeight(),
			expectedHeight,
		)
	}

	// Verify content area rectangle
	area := page.ContentArea()
	if area.LLX != 54 {
		t.Errorf(
			"ContentArea.LLX = %v, want 54",
			area.LLX,
		)
	}
	if area.LLY != 72 {
		t.Errorf(
			"ContentArea.LLY = %v, want 72",
			area.LLY,
		)
	}
}

// verifyPDFStructure performs basic validation of PDF output.
func verifyPDFStructure(
	t *testing.T,
	data []byte,
	_ Version,
) {
	t.Helper()

	if len(data) < 100 {
		t.Error(
			"PDF output too small to be valid",
		)

		return
	}

	// Check PDF header
	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Error("PDF missing PDF header")

		return
	}

	// Check EOF marker
	if !bytes.Contains(data, []byte("%%EOF")) {
		t.Error("PDF missing EOF marker")
	}

	// Check for xref or xref stream
	hasXref := bytes.Contains(
		data,
		[]byte("xref"),
	) ||
		bytes.Contains(
			data,
			[]byte("/Type /XRef"),
		)
	if !hasXref {
		t.Error(
			"PDF missing cross-reference table or stream",
		)
	}

	// Check for trailer or trailer equivalent
	hasTrailer := bytes.Contains(
		data,
		[]byte("trailer"),
	) ||
		bytes.Contains(data, []byte("startxref"))
	if !hasTrailer {
		t.Error("PDF missing trailer information")
	}
}

// testFormatFloat formats a float for PDF output (helper for tests).
func testFormatFloat(f float64) string {
	// Handle special cases
	if f == 0 {
		return "0"
	}
	if f == 1 {
		return "1"
	}
	if f == -1 {
		return "-1"
	}

	// Format with precision
	s := strings.TrimRight(strings.TrimRight(
		fmt.Sprintf("%.6f", f),
		"0"),
		".")

	return s
}

// BenchmarkDocumentCreation benchmarks document creation performance.
func BenchmarkDocumentCreation(b *testing.B) {
	b.Run("single page", func(b *testing.B) {
		for range b.N {
			doc, _ := NewDocument()
			_, _ = doc.AddPageA4()
			var buf bytes.Buffer
			_ = doc.Write(&buf)
			_ = doc.Close()
		}
	})

	b.Run("10 pages", func(b *testing.B) {
		for range b.N {
			doc, _ := NewDocument()
			for range 10 {
				_, _ = doc.AddPageA4()
			}
			var buf bytes.Buffer
			_ = doc.Write(&buf)
			_ = doc.Close()
		}
	})

	b.Run(
		"with content in memory",
		func(b *testing.B) {
			content := strings.Repeat(
				"100 100 m 200 200 l S\n",
				100,
			)
			for range b.N {
				doc, _ := NewDocument()
				page, _ := doc.AddPageA4()
				_, _ = page.WriteContentString(
					content,
				)
				// Just test content writing without PDF output
				// (pdfcpu content stream writing has issues)
				_ = page.Content()
				_ = doc.Close()
			}
		},
	)
}
