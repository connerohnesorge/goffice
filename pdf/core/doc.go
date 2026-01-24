//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package core provides PDF document creation and low-level PDF primitives.
//
// This package contains the fundamental building blocks for PDF generation:
//   - PDF document structure (pages, objects, streams)
//   - Content stream operators (text, graphics, images)
//   - Resource management (fonts, images, color spaces)
//   - Object serialization and cross-reference tables
//
// The core package abstracts over the underlying PDF library (pdfcpu) to provide
// a clean API for the higher-level renderers in the word, spreadsheet, and
// presentation packages.
//
// # Document Creation
//
// A PDF document consists of a catalog, page tree, and resources. This package
// handles the creation and management of these structures:
//
//	// Create a new document with default options (PDF 1.7)
//	doc, err := core.NewDocument()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer doc.Close()
//
//	// Add a page
//	page, err := doc.AddPage(core.PageSizeA4)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Write to file
//	if err := doc.WriteToFile("output.pdf"); err != nil {
//	    log.Fatal(err)
//	}
//
// # Document Options
//
// Documents can be created with custom options:
//
//	opts := &core.DocumentOptions{
//	    Version: core.V17,
//	    Metadata: core.Metadata{
//	        Title:   "My Document",
//	        Author:  "John Doe",
//	        Creator: "My Application",
//	    },
//	    CompressContent: true,
//	}
//	doc, err := core.NewDocumentWithOptions(opts)
//
// # PDF Versions
//
// The package supports PDF versions 1.4 through 2.0:
//   - V14: PDF 1.4 (Acrobat 5)
//   - V15: PDF 1.5 (Acrobat 6) - adds object streams
//   - V16: PDF 1.6 (Acrobat 7)
//   - V17: PDF 1.7 (ISO 32000-1:2008) - recommended default
//   - V20: PDF 2.0 (ISO 32000-2:2017)
//
// # Page Sizes
//
// Common page sizes are provided as constants:
//
// ISO A Series:
//   - PageSizeA0, PageSizeA1, PageSizeA2, PageSizeA3, PageSizeA4, PageSizeA5, PageSizeA6
//
// ISO B Series:
//   - PageSizeB4, PageSizeB5
//
// US Sizes:
//   - PageSizeLetter: 8.5in x 11in
//   - PageSizeLegal: 8.5in x 14in
//   - PageSizeTabloid: 11in x 17in
//   - PageSizeLedger: 17in x 11in (landscape Tabloid)
//   - PageSizeExecutive: 7.25in x 10.5in
//   - PageSizeStatement: 5.5in x 8.5in
//   - PageSizeFolio: 8.5in x 13in
//   - PageSizeQuarto: 8.5in x 10.83in
//
// Custom page sizes can be created in various ways:
//
//	// Direct point values
//	customSize := core.PageSize{Width: 500, Height: 700}
//
//	// From dimensions with units
//	customSize := core.NewPageSizeFromDimensions(core.Inches(8.5), core.Inches(11))
//
//	// From a specific unit
//	customSize := core.NewPageSizeFromUnit(210, 297, core.UnitMM)
//
// Landscape and portrait orientations:
//
//	landscapeA4 := core.PageSizeA4.Landscape()
//	portraitA4 := landscapeA4.Portrait()
//
// # Page Options
//
// Pages can be created with comprehensive options including margins, rotation,
// and page boxes:
//
//	opts := core.NewPageOptions(core.PageSizeLetter).
//	    WithMargins(core.DefaultMargins()).
//	    WithRotation(90).
//	    Landscape()
//	page, err := doc.AddPageWithOptions(opts)
//
// # Page Boxes
//
// PDF supports multiple page boundary boxes (per PDF Reference 14.11.2):
//   - MediaBox: physical medium boundaries (required)
//   - CropBox: visible region (defaults to MediaBox)
//   - BleedBox: production clipping region
//   - TrimBox: finished page dimensions
//   - ArtBox: meaningful content area
//
// Example:
//
//	boxes := core.NewPageBoxesFromSize(core.PageSizeLetter).
//	    WithCropBox(core.NewRectangle(18, 18, 594, 774)).
//	    WithBleed(9)
//	opts := core.NewPageOptions(core.PageSizeLetter).WithBoxes(boxes)
//
// # Margins
//
// Standard margin presets are available:
//   - DefaultMargins(): 1 inch on all sides
//   - NarrowMargins(): 0.5 inch on all sides
//   - WideMargins(): 1.5 inch horizontal, 1 inch vertical
//   - NoMargins(): zero margins
//
// Custom margins:
//
//	margins := core.NewMarginsFromDimensions(
//	    core.Inches(1), core.Inches(0.75),
//	    core.Inches(1), core.Inches(0.75))
//
// # Units
//
// All coordinates in this package use PDF points (1/72 inch) as the base unit.
// Conversion utilities are provided for common units:
//
//	// Convert from other units to points
//	points := core.InchesToPoints(1.5)
//	points := core.MMToPoints(25.4)
//	points := core.CMToPoints(2.54)
//	points := core.EMUToPoints(914400)  // OOXML EMU
//	points := core.TwipsToPoints(1440)  // OOXML Twips
//
//	// Using the Dimension type
//	dim := core.Inches(1.5)
//	points := dim.ToPoints()
//
//	// Using the Unit type
//	points := core.UnitMM.ToPoints(25.4)
//
// # Coordinate System
//
// PDF uses a bottom-left origin coordinate system. For compatibility with
// OOXML (which uses top-left origin), use the CoordinateSystem type:
//
//	cs := page.CoordinateSystem().WithTopLeftOrigin()
//	pdfY := cs.TransformY(ooxmlY)
//
// # Thread Safety
//
// The Document and Page types are safe for concurrent use. Multiple goroutines
// can add pages and write content simultaneously, though the final page order
// depends on which AddPage calls complete first.
//
// # Error Handling
//
// The package defines several error types:
//   - DocumentError: general document operation errors
//   - PageError: errors related to specific pages
//   - WriteError: errors during PDF writing
//
// Common sentinel errors:
//   - ErrDocumentClosed: operation on a closed document
//   - ErrInvalidPageSize: page dimensions are invalid
//   - ErrNoPages: attempting to write a document with no pages
package core
