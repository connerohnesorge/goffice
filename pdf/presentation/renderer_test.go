package presentation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/presentation"
)

// TestPresentationRenderer_StandardSlides tests rendering in standard slide mode.
func TestPresentationRenderer_StandardSlides(
	t *testing.T,
) {
	// Create a simple presentation
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a slide
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	// Create renderer
	renderer := NewPresentationRenderer(doc)

	// Render to PDF
	pdfDoc, err := renderer.Render()
	if err != nil {
		t.Fatalf(
			"failed to render presentation: %v",
			err,
		)
	}
	defer pdfDoc.Close()

	// Verify PDF has pages
	if pdfDoc.PageCount() == 0 {
		t.Error(
			"expected PDF to have at least one page",
		)
	}
}

// TestPresentationRenderer_MultipleSlides tests rendering multiple slides.
func TestPresentationRenderer_MultipleSlides(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add multiple slides
	for i := 0; i < 3; i++ {
		if _, err := doc.AddSlide(); err != nil {
			t.Fatalf(
				"failed to add slide %d: %v",
				i+1,
				err,
			)
		}
	}

	renderer := NewPresentationRenderer(doc)
	pdfDoc, err := renderer.Render()
	if err != nil {
		t.Fatalf(
			"failed to render presentation: %v",
			err,
		)
	}
	defer pdfDoc.Close()

	// Verify page count
	if pdfDoc.PageCount() != 3 {
		t.Errorf(
			"expected 3 pages, got %d",
			pdfDoc.PageCount(),
		)
	}
}

// TestPresentationRenderer_SlideRange tests rendering a subset of slides.
func TestPresentationRenderer_SlideRange(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add 5 slides
	for i := 0; i < 5; i++ {
		if _, err := doc.AddSlide(); err != nil {
			t.Fatalf(
				"failed to add slide %d: %v",
				i+1,
				err,
			)
		}
	}

	// Render only slides 1, 3, 5
	renderer := NewPresentationRenderer(doc)
	renderer.SetOptions(&RenderOptions{
		OutputMode: OutputModeSlides,
		SlideRange: []int{1, 3, 5},
	})

	pdfDoc, err := renderer.Render()
	if err != nil {
		t.Fatalf(
			"failed to render presentation: %v",
			err,
		)
	}
	defer pdfDoc.Close()

	// Verify page count
	if pdfDoc.PageCount() != 3 {
		t.Errorf(
			"expected 3 pages, got %d",
			pdfDoc.PageCount(),
		)
	}
}

// TestPresentationRenderer_NotesPages tests notes page rendering.
func TestPresentationRenderer_NotesPages(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a slide
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	// Render notes pages
	renderer := NewPresentationRenderer(doc)
	renderer.SetOptions(&RenderOptions{
		OutputMode: OutputModeNotes,
	})

	pdfDoc, err := renderer.Render()
	if err != nil {
		t.Fatalf(
			"failed to render notes pages: %v",
			err,
		)
	}
	defer pdfDoc.Close()

	// Verify PDF has pages
	if pdfDoc.PageCount() == 0 {
		t.Error(
			"expected PDF to have at least one page",
		)
	}
}

// TestPresentationRenderer_Handouts tests handout rendering.
func TestPresentationRenderer_Handouts(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add 6 slides
	for i := 0; i < 6; i++ {
		if _, err := doc.AddSlide(); err != nil {
			t.Fatalf(
				"failed to add slide %d: %v",
				i+1,
				err,
			)
		}
	}

	tests := []struct {
		name          string
		layout        HandoutLayout
		expectedPages int
	}{
		{"1 per page", Handout1, 6},
		{"2 per page", Handout2, 3},
		{"3 per page", Handout3, 2},
		{"4 per page", Handout4, 2},
		{"6 per page", Handout6, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := NewPresentationRenderer(
				doc,
			)
			renderer.SetOptions(&RenderOptions{
				OutputMode:    OutputModeHandouts,
				HandoutLayout: tt.layout,
			})

			pdfDoc, err := renderer.Render()
			if err != nil {
				t.Fatalf(
					"failed to render handouts: %v",
					err,
				)
			}
			defer pdfDoc.Close()

			if pdfDoc.PageCount() != tt.expectedPages {
				t.Errorf(
					"expected %d pages, got %d",
					tt.expectedPages,
					pdfDoc.PageCount(),
				)
			}
		})
	}
}

// TestPresentationRenderer_RenderToFile tests rendering directly to file.
func TestPresentationRenderer_RenderToFile(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a slide
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	// Render to file
	outputPath := filepath.Join(
		t.TempDir(),
		"output.pdf",
	)
	renderer := NewPresentationRenderer(doc)

	if err := renderer.RenderToFile(outputPath); err != nil {
		t.Fatalf(
			"failed to render to file: %v",
			err,
		)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Error("expected output file to exist")
	}
}

// TestPresentationRenderer_EmptyPresentation tests rendering an empty presentation.
func TestPresentationRenderer_EmptyPresentation(
	t *testing.T,
) {
	doc, err := presentation.New(
		filepath.Join(t.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Don't add any slides
	renderer := NewPresentationRenderer(doc)

	_, err = renderer.Render()
	if err == nil {
		t.Error(
			"expected error when rendering empty presentation",
		)
	}
}

// TestSlideSizes tests various slide size configurations.
func TestSlideSizes(t *testing.T) {
	tests := []struct {
		name string
		size core.PageSize
	}{
		{"Standard 4:3", StandardSlideSize()},
		{
			"Widescreen 16:9",
			WidescreenSlideSize(),
		},
		{"Letter", LetterSlideSize()},
		{"A4", A4SlideSize()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.size.Width <= 0 ||
				tt.size.Height <= 0 {
				t.Errorf(
					"invalid slide size: %v x %v",
					tt.size.Width,
					tt.size.Height,
				)
			}
		})
	}
}

// TestOutputModeString tests OutputMode string representation.
func TestOutputModeString(t *testing.T) {
	tests := []struct {
		mode     OutputMode
		expected string
	}{
		{OutputModeSlides, "slides"},
		{OutputModeNotes, "notes"},
		{OutputModeHandouts, "handouts"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.mode.String() != tt.expected {
				t.Errorf(
					"expected %q, got %q",
					tt.expected,
					tt.mode.String(),
				)
			}
		})
	}
}
