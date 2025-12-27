// Package presentation provides PresentationML to PDF rendering capabilities.
//
// This package converts PowerPoint presentations (.pptx files) to PDF format,
// handling all aspects of slide rendering including:
//
//   - Slide layouts and dimensions (4:3, 16:9, custom sizes)
//   - Master slide and layout slide inheritance
//   - Background rendering (solid colors, gradients, images)
//   - Shape rendering (rectangles, circles, custom paths)
//   - Text frames with formatting and alignment
//   - Pictures and images
//   - Tables with styling
//   - Charts (leveraging DrawingML renderer)
//   - Multiple output modes: slides, notes pages, handouts
//
// # Architecture
//
// The rendering system follows a hierarchical approach:
//
//  1. PresentationRenderer - Entry point for rendering entire presentations
//  2. SlideRenderer - Renders individual slides with master inheritance
//  3. NotesPageRenderer - Renders notes pages (slide + speaker notes)
//  4. HandoutPageRenderer - Renders multiple slides per page
//
// # Master Slide Inheritance
//
// Slides inherit properties from slide layouts, which in turn inherit from
// slide masters. The inheritance chain is:
//
//	Slide → Slide Layout → Slide Master
//
// Properties defined at a lower level (e.g., slide) override properties from
// higher levels (e.g., master). This includes:
//
//   - Background fills
//   - Placeholder content
//   - Color schemes (via theme)
//   - Text styles
//
// # Coordinate System
//
// PowerPoint uses EMU (English Metric Units) where 914400 EMUs = 1 inch.
// This package converts EMU to PDF points (72 points = 1 inch) using utilities
// from the drawingml package.
//
// # Usage Example
//
// Basic slide rendering:
//
//	import (
//	    "github.com/connerohnesorge/goffice/presentation"
//	    pdfpres "github.com/connerohnesorge/goffice-pdf/presentation"
//	)
//
//	// Open presentation
//	doc, err := presentation.Open("presentation.pptx", false)
//	if err != nil {
//	    panic(err)
//	}
//	defer doc.Close()
//
//	// Create renderer
//	renderer := pdfpres.NewPresentationRenderer(doc)
//
//	// Render to PDF file
//	err = renderer.RenderToFile("output.pdf")
//	if err != nil {
//	    panic(err)
//	}
//
// # Rendering modes:
//
// Standard slides (default):
//
//	renderer.SetOptions(&pdfpres.RenderOptions{
//	    OutputMode: pdfpres.OutputModeSlides,
//	})
//
// Notes pages (slide thumbnail + speaker notes):
//
//	renderer.SetOptions(&pdfpres.RenderOptions{
//	    OutputMode: pdfpres.OutputModeNotes,
//	})
//
// Handouts (multiple slides per page):
//
//	renderer.SetOptions(&pdfpres.RenderOptions{
//	    OutputMode: pdfpres.OutputModeHandouts,
//	    HandoutLayout: pdfpres.Handout6, // 6 slides per page
//	})
//
// Render specific slides:
//
//	renderer.SetOptions(&pdfpres.RenderOptions{
//	    SlideRange: []int{1, 3, 5}, // Only slides 1, 3, and 5
//	})
//
// # Dependencies
//
// This package leverages:
//
//   - github.com/connerohnesorge/goffice/presentation - PresentationML parsing
//   - github.com/connerohnesorge/goffice/drawingml - Shape and graphics support
//   - github.com/connerohnesorge/goffice-pdf/core - PDF document creation
//   - github.com/connerohnesorge/goffice-pdf/drawing - DrawingML to PDF rendering
//
// # Limitations
//
// The current implementation has the following limitations:
//
//   - Animations are not rendered (static snapshots only)
//   - Video and audio content is not embedded
//   - SmartArt is rendered as decomposed shapes (basic support)
//   - Some advanced effects may be approximated
//   - Slide transitions are not preserved
//
// # Performance
//
// Rendering performance scales linearly with the number of slides and complexity
// of content. For large presentations (100+ slides), consider:
//
//   - Using slide ranges to render subsets
//   - Rendering slides in parallel (future enhancement)
//   - Pre-loading images and fonts
//
// # Thread Safety
//
// Renderers are not thread-safe. Create separate renderer instances for
// concurrent rendering operations.
package presentation
