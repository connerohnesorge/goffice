package presentation

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/presentation"
)

// TestIntegration_RealPresentation tests rendering a real .pptx file if available.
func TestIntegration_RealPresentation(
	t *testing.T,
) {
	// Look for test fixture
	fixturePath := filepath.Join(
		"..",
		"..",
		"testdata",
		"fixtures",
		"presentation.pptx",
	)
	if _, err := os.Stat(fixturePath); os.IsNotExist(
		err,
	) {
		t.Skip("test fixture not available")
	}

	// Open presentation
	doc, err := presentation.Open(
		fixturePath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"failed to open presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Render to PDF
	renderer := NewPresentationRenderer(doc)
	outputPath := filepath.Join(
		t.TempDir(),
		"output.pdf",
	)

	if err := renderer.RenderToFile(outputPath); err != nil {
		t.Fatalf(
			"failed to render presentation: %v",
			err,
		)
	}

	// Verify output exists
	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("failed to stat output: %v", err)
	}

	if info.Size() == 0 {
		t.Error("output PDF is empty")
	}

	t.Logf(
		"Rendered PDF to %s (size: %d bytes)",
		outputPath,
		info.Size(),
	)
}

// TestIntegration_SlideRendering tests complete slide rendering workflow.
func TestIntegration_SlideRendering(
	t *testing.T,
) {
	// Create a presentation with various content
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
		slide, err := doc.AddSlide()
		if err != nil {
			t.Fatalf(
				"failed to add slide %d: %v",
				i+1,
				err,
			)
		}

		// Add content to slides
		slideElem := slide.Slide()
		if slideElem != nil {
			// Add a title shape
			titleShape := slideElem.AddShape()
			titleShape.SetText(
				fmt.Sprintf(
					"Slide %d Title",
					i+1,
				),
			)
			titleShape.SetPosition(
				914400,
				457200,
			) // 1 inch from left, 0.5 inches from top
			titleShape.SetSize(
				7315200,
				914400,
			) // 8 inches wide, 1 inch tall
			titleShape.SetShapeType("rect")
			titleShape.SetSolidFill(
				"4472C4",
			) // Blue background

			// Add a body text shape
			bodyShape := slideElem.AddShape()
			bodyShape.SetText(
				fmt.Sprintf(
					"This is the body text for slide %d.\nIt can have multiple lines.",
					i+1,
				),
			)
			bodyShape.SetPosition(
				914400,
				1828800,
			) // 1 inch from left, 2 inches from top
			bodyShape.SetSize(
				7315200,
				2743200,
			) // 8 inches wide, 3 inches tall
		}
	}

	// Test all output modes
	modes := []struct {
		name string
		opts *RenderOptions
	}{
		{
			"slides",
			&RenderOptions{
				OutputMode: OutputModeSlides,
			},
		},
		{
			"notes",
			&RenderOptions{
				OutputMode: OutputModeNotes,
			},
		},
		{
			"handouts_4",
			&RenderOptions{
				OutputMode:    OutputModeHandouts,
				HandoutLayout: Handout4,
			},
		},
	}

	for _, mode := range modes {
		t.Run(mode.name, func(t *testing.T) {
			renderer := NewPresentationRenderer(
				doc,
			)
			renderer.SetOptions(mode.opts)

			outputPath := filepath.Join(
				t.TempDir(),
				mode.name+".pdf",
			)
			if err := renderer.RenderToFile(outputPath); err != nil {
				t.Fatalf(
					"failed to render %s mode: %v",
					mode.name,
					err,
				)
			}

			info, err := os.Stat(outputPath)
			if err != nil {
				t.Fatalf(
					"failed to stat output: %v",
					err,
				)
			}

			if info.Size() == 0 {
				t.Errorf(
					"output PDF is empty for %s mode",
					mode.name,
				)
			}

			t.Logf(
				"%s mode: %d bytes",
				mode.name,
				info.Size(),
			)
		})
	}
}

// TestIntegration_MasterSlideInheritance tests master slide inheritance.
func TestIntegration_MasterSlideInheritance(
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

	// Add a slide master
	master, err := doc.AddSlideMaster()
	if err != nil {
		t.Fatalf(
			"failed to add slide master: %v",
			err,
		)
	}
	_ = master

	// Add slides
	for i := 0; i < 2; i++ {
		if _, err := doc.AddSlide(); err != nil {
			t.Fatalf(
				"failed to add slide: %v",
				err,
			)
		}
	}

	// Render
	renderer := NewPresentationRenderer(doc)
	outputPath := filepath.Join(
		t.TempDir(),
		"master_test.pdf",
	)

	if err := renderer.RenderToFile(outputPath); err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	// Just verify it doesn't crash
	t.Log(
		"Master slide inheritance test completed",
	)
}

// TestIntegration_Performance tests rendering performance.
func TestIntegration_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip(
			"skipping performance test in short mode",
		)
	}

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

	// Add many slides
	numSlides := 50
	for i := 0; i < numSlides; i++ {
		if _, err := doc.AddSlide(); err != nil {
			t.Fatalf(
				"failed to add slide %d: %v",
				i+1,
				err,
			)
		}
	}

	// Render and measure
	renderer := NewPresentationRenderer(doc)
	outputPath := filepath.Join(
		t.TempDir(),
		"performance.pdf",
	)

	if err := renderer.RenderToFile(outputPath); err != nil {
		t.Fatalf("failed to render: %v", err)
	}

	info, err := os.Stat(outputPath)
	if err != nil {
		t.Fatalf("failed to stat output: %v", err)
	}

	t.Logf(
		"Rendered %d slides: %d bytes",
		numSlides,
		info.Size(),
	)
}

// BenchmarkSlideRendering benchmarks slide rendering performance.
func BenchmarkSlideRendering(b *testing.B) {
	doc, err := presentation.New(
		filepath.Join(b.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		b.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a single slide
	if _, err := doc.AddSlide(); err != nil {
		b.Fatalf("failed to add slide: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		renderer := NewPresentationRenderer(doc)
		pdfDoc, err := renderer.Render()
		if err != nil {
			b.Fatalf("failed to render: %v", err)
		}
		pdfDoc.Close()
	}
}

// BenchmarkHandoutRendering benchmarks handout rendering performance.
func BenchmarkHandoutRendering(b *testing.B) {
	doc, err := presentation.New(
		filepath.Join(b.TempDir(), "test.pptx"),
		presentation.DocTypePresentation,
	)
	if err != nil {
		b.Fatalf(
			"failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add slides
	for i := 0; i < 6; i++ {
		if _, err := doc.AddSlide(); err != nil {
			b.Fatalf(
				"failed to add slide: %v",
				err,
			)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		renderer := NewPresentationRenderer(doc)
		renderer.SetOptions(&RenderOptions{
			OutputMode:    OutputModeHandouts,
			HandoutLayout: Handout6,
		})

		pdfDoc, err := renderer.Render()
		if err != nil {
			b.Fatalf("failed to render: %v", err)
		}
		pdfDoc.Close()
	}
}
