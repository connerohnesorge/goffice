// Package pdf provides PDF rendering for Microsoft Office documents.
//
// # Overview
//
// goffice-pdf renders Word (.docx), Excel (.xlsx), and PowerPoint (.pptx)
// documents to PDF format with high fidelity matching Microsoft Office output.
//
// This is a separate module from the core goffice package, allowing users to
// opt-in to PDF rendering functionality and its dependencies.
//
// # Quick Start
//
// Word document to PDF:
//
//	doc, _ := wordprocessing.Open("input.docx", false)
//	defer doc.Close()
//
//	f, _ := os.Create("output.pdf")
//	defer f.Close()
//
//	pdf.RenderWord(doc, f, nil)
//
// Excel spreadsheet to PDF:
//
//	doc, _ := spreadsheet.Open("input.xlsx", false)
//	defer doc.Close()
//
//	pdf.RenderSpreadsheet(doc, f, nil)
//
// PowerPoint presentation to PDF:
//
//	doc, _ := presentation.Open("input.pptx", false)
//	defer doc.Close()
//
//	pdf.RenderPresentation(doc, f, nil)
//
// # Render Options
//
// Customize rendering with RenderOptions:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.ImageDPI = 300                    // High quality images
//	opts.FontEmbedding = pdf.EmbedSubset   // Smaller file size
//	opts.ACompliance = pdf.PDFA1b          // Archival PDF/A
//	opts.TaggedPDF = true                  // Accessible PDF
//	opts.CompressContent = true            // Compress content streams
//
//	// Set metadata
//	opts.Title = "My Document"
//	opts.Author = "John Doe"
//	opts.Subject = "Report"
//	opts.Keywords = "analysis, 2024"
//
//	pdf.RenderWord(doc, f, opts)
//
// # Streaming Output
//
// Monitor rendering progress with page callbacks:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.OnPageComplete = func(pageNum int, data []byte) {
//	    fmt.Printf("Completed page %d (%d bytes)\n", pageNum, len(data))
//	}
//
//	pdf.RenderWord(doc, f, opts)
//
// # Font Handling
//
// The renderer automatically discovers and uses system fonts. You can control
// font embedding:
//
//	// Embed only used glyphs (default, smallest files)
//	opts.FontEmbedding = pdf.EmbedSubset
//
//	// Embed complete fonts (larger files, allows PDF editing)
//	opts.FontEmbedding = pdf.EmbedFull
//
//	// Don't embed fonts (smallest files, may not render correctly everywhere)
//	opts.FontEmbedding = pdf.NoEmbed
//
// # Image Handling
//
// Control image quality and memory usage:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.ImageDPI = 300                        // Print quality (default: 150)
//	opts.ImageCacheLimit = 200 * 1024 * 1024   // 200MB cache (default: 100MB)
//
// # PDF/A Compliance
//
// Generate archival-compliant PDFs:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.ACompliance = pdf.PDFA1b  // PDF/A-1b (most compatible)
//	opts.ACompliance = pdf.PDFA2b  // PDF/A-2b (JPEG2000 support)
//	opts.ACompliance = pdf.PDFA3b  // PDF/A-3b (embedded files)
//
// PDF/A compliance requires:
//   - All fonts embedded
//   - All images in supported formats
//   - No encryption
//   - Metadata in XMP format
//
// # Accessibility
//
// Generate accessible PDFs with structure tags:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.TaggedPDF = true  // Enable structure tags for screen readers
//
// Tagged PDFs include:
//   - Document structure tree
//   - Heading hierarchy
//   - Table structure
//   - List structure
//   - Alternative text for images
//
// # Performance
//
// For large documents, consider:
//
//	opts := pdf.DefaultRenderOptions()
//	opts.ImageCacheLimit = 500 * 1024 * 1024  // Larger cache for big docs
//	opts.CompressContent = true               // Reduce file size
//	opts.FontEmbedding = pdf.EmbedSubset      // Smaller embedded fonts
//
//	// Monitor progress
//	opts.OnPageComplete = func(pageNum int, data []byte) {
//	    // Update progress bar, log, etc.
//	}
//
// # Error Handling
//
// All rendering functions return errors that should be checked:
//
//	if err := pdf.RenderWord(doc, f, opts); err != nil {
//	    switch {
//	    case errors.Is(err, pdf.ErrFontNotFound):
//	        // Handle missing font
//	    case errors.Is(err, pdf.ErrNilDocument):
//	        // Handle nil document
//	    default:
//	        // Handle other errors
//	    }
//	}
//
// # Package Structure
//
//	pdf/                   Main API and types
//	pdf/core/              PDF document abstraction
//	pdf/drawing/           DrawingML rendering (shapes, images, charts)
//	pdf/font/              Font handling and subsetting
//	pdf/layout/            Text layout engine
//	pdf/word/              Word document renderer
//	pdf/spreadsheet/       Excel spreadsheet renderer (planned)
//	pdf/presentation/      PowerPoint renderer (planned)
//
// # Fidelity
//
// goffice-pdf aims for visual equivalence with Microsoft Office output.
// See FIDELITY.md for known differences and limitations.
//
// Supported features:
//   - Text formatting (fonts, sizes, colors, effects)
//   - Paragraph formatting (alignment, spacing, indents)
//   - Tables (borders, shading, cell merging)
//   - Lists (bullets, numbering, multi-level)
//   - Headers and footers
//   - Page breaks and section breaks
//   - Images and shapes
//   - Charts
//   - Page layout (margins, columns, orientation)
//
// Not supported:
//   - Track changes
//   - Comments
//   - Form fields (rendered as static text)
//   - Macros/VBA
//   - Embedded objects (OLE)
//   - Animations (PowerPoint)
//   - Video/audio
//
// # Examples
//
// See examples/ directory for complete examples:
//   - examples/word_to_pdf.go
//   - examples/excel_to_pdf.go
//   - examples/powerpoint_to_pdf.go
//   - examples/batch_convert.go
//   - examples/custom_options.go
//
// # Benchmarks
//
// Run benchmarks to measure performance:
//
//	go test -bench=. -benchmem
//
// # License
//
// goffice-pdf is released under the MIT License.
package pdf
