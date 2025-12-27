// Package pdf provides the main API entry points for PDF rendering.
package pdf

import (
	"errors"
	"io"

	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/spreadsheet"
	"github.com/connerohnesorge/goffice/wordprocessing"
)

// Version is the current version of the goffice-pdf module.
const Version = "0.1.0"

// FontEmbedMode specifies how fonts should be embedded in the PDF output.
type FontEmbedMode int

const (
	// EmbedSubset embeds only the glyphs actually used in the document.
	// This is the default and recommended mode as it produces smaller files
	// while maintaining full visual fidelity.
	EmbedSubset FontEmbedMode = iota

	// EmbedFull embeds complete font files in the PDF output.
	// This produces larger files but ensures all glyphs are available
	// if the PDF is later edited.
	EmbedFull

	// NoEmbed does not embed fonts, instead referencing them by name.
	// This produces the smallest files but may result in different
	// rendering if the fonts are not available on the viewing system.
	// Not recommended for documents requiring visual fidelity.
	NoEmbed
)

// String returns the string representation of the font embedding mode.
func (m FontEmbedMode) String() string {
	switch m {
	case EmbedSubset:
		return "EmbedSubset"
	case EmbedFull:
		return "EmbedFull"
	case NoEmbed:
		return "NoEmbed"
	default:
		return "EmbedSubset"
	}
}

// ACompliance specifies the PDF/A conformance level for archival output.
type ACompliance int

const (
	// PDFANone disables PDF/A compliance (standard PDF 1.7 output).
	PDFANone ACompliance = iota

	// PDFA1b enables PDF/A-1b compliance (basic archival, visual appearance).
	PDFA1b

	// PDFA2b enables PDF/A-2b compliance (improved archival with JPEG2000
	// support).
	PDFA2b

	// PDFA3b enables PDF/A-3b compliance (archival with embedded files
	// support).
	PDFA3b
)

// String returns the string representation of the PDF/A compliance level.
func (c ACompliance) String() string {
	switch c {
	case PDFANone:
		return "None"
	case PDFA1b:
		return "PDF/A-1b"
	case PDFA2b:
		return "PDF/A-2b"
	case PDFA3b:
		return "PDF/A-3b"
	default:
		return "None"
	}
}

// RenderOptions configures the PDF rendering behavior.
type RenderOptions struct {
	// OnPageComplete is called after each page is rendered.
	// The callback receives the page number (1-based) and the raw PDF data
	// for that page. This can be used to track rendering progress or
	// to stream pages as they are completed.
	OnPageComplete func(pageNum int, data []byte)

	// ImageDPI specifies the target resolution for images in the PDF.
	// Higher values produce better quality but larger files.
	// Default is 150 DPI. Use 300 for print quality.
	ImageDPI int

	// ImageCacheLimit sets the maximum memory (in bytes) for caching
	// decoded images during rendering. This helps control memory usage
	// for documents with many large images.
	// Default is 100MB (100 * 1024 * 1024).
	ImageCacheLimit int64

	// FontEmbedding controls how fonts are embedded in the output PDF.
	// Default is EmbedSubset.
	FontEmbedding FontEmbedMode

	// ACompliance enables PDF/A output for archival purposes.
	// Default is PDFANone (standard PDF output).
	ACompliance ACompliance

	// TaggedPDF enables accessible PDF output with structure tags.
	// When enabled, the renderer generates tagged PDF content that
	// can be read by screen readers and other accessibility tools.
	// Default is false.
	TaggedPDF bool

	// Title sets the PDF document title metadata.
	// If empty, the source document title is used.
	Title string

	// Author sets the PDF document author metadata.
	// If empty, the source document author is used.
	Author string

	// Subject sets the PDF document subject metadata.
	Subject string

	// Keywords sets the PDF document keywords metadata.
	Keywords string

	// Creator sets the PDF creator metadata.
	// Default is "goffice-pdf".
	Creator string

	// CompressContent enables content stream compression.
	// Default is true.
	CompressContent bool
}

const (
	// DefaultImageDPI is the default resolution for images in PDF (150 DPI).
	DefaultImageDPI = 150
	// DefaultImageCacheLimit is the default memory limit for image caching (100MB).
	DefaultImageCacheLimit = 100 * 1024 * 1024
)

// DefaultRenderOptions returns a new RenderOptions with sensible defaults.
func DefaultRenderOptions() *RenderOptions {
	return &RenderOptions{
		ImageDPI:        DefaultImageDPI,
		ImageCacheLimit: DefaultImageCacheLimit,
		FontEmbedding:   EmbedSubset,
		ACompliance:     PDFANone,
		TaggedPDF:       false,
		Creator:         "goffice-pdf",
		CompressContent: true,
	}
}

// WithImageDPI returns a copy of the options with the specified image DPI.
func (o *RenderOptions) WithImageDPI(
	dpi int,
) *RenderOptions {
	opts := *o
	opts.ImageDPI = dpi

	return &opts
}

// WithFontEmbedding returns a copy of the options with the specified
// font embedding mode.
func (o *RenderOptions) WithFontEmbedding(
	mode FontEmbedMode,
) *RenderOptions {
	opts := *o
	opts.FontEmbedding = mode

	return &opts
}

// WithACompliance returns a copy of the options with the specified
// PDF/A compliance level.
func (o *RenderOptions) WithACompliance(
	compliance ACompliance,
) *RenderOptions {
	opts := *o
	opts.ACompliance = compliance

	return &opts
}

// WithTaggedPDF returns a copy of the options with tagged PDF enabled
// or disabled.
func (o *RenderOptions) WithTaggedPDF(
	tagged bool,
) *RenderOptions {
	opts := *o
	opts.TaggedPDF = tagged

	return &opts
}

// WithOnPageComplete returns a copy of the options with the page
// completion callback set.
func (o *RenderOptions) WithOnPageComplete(
	callback func(pageNum int, data []byte),
) *RenderOptions {
	opts := *o
	opts.OnPageComplete = callback

	return &opts
}

// WithImageCacheLimit returns a copy of the options with the specified
// image cache limit in bytes.
func (o *RenderOptions) WithImageCacheLimit(
	limit int64,
) *RenderOptions {
	opts := *o
	opts.ImageCacheLimit = limit

	return &opts
}

// WithMetadata returns a copy of the options with the specified PDF
// metadata fields.
func (o *RenderOptions) WithMetadata(
	title, author, subject, keywords string,
) *RenderOptions {
	opts := *o
	opts.Title = title
	opts.Author = author
	opts.Subject = subject
	opts.Keywords = keywords

	return &opts
}

// Errors returned by rendering functions.
var (
	// ErrNotImplemented is returned when a rendering function has not
	// yet been implemented.
	ErrNotImplemented = errors.New(
		"pdf: rendering not yet implemented",
	)

	// ErrNilDocument is returned when a nil document is passed to a
	// render function.
	ErrNilDocument = errors.New(
		"pdf: document is nil",
	)

	// ErrNilWriter is returned when a nil writer is passed to a
	// render function.
	ErrNilWriter = errors.New(
		"pdf: writer is nil",
	)

	// ErrFontNotFound is returned when a required font cannot be located.
	ErrFontNotFound = errors.New(
		"pdf: required font not found",
	)

	// ErrInvalidOptions is returned when render options contain invalid values.
	ErrInvalidOptions = errors.New(
		"pdf: invalid render options",
	)
)

// RenderWord renders a Word document to PDF format.
//
// The document must be a valid wordprocessing.Document that has been
// opened or created using the goffice wordprocessing package.
//
// The output is written to the provided io.Writer. If options is nil,
// DefaultRenderOptions() is used.
//
// Example:
//
//	doc, err := wordprocessing.Open("input.docx", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer doc.Close()
//
//	f, err := os.Create("output.pdf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer f.Close()
//
//	if err := pdf.RenderWord(doc, f, nil); err != nil {
//	    log.Fatal(err)
//	}
func RenderWord(
	doc *wordprocessing.Document,
	output io.Writer,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if output == nil {
		return ErrNilWriter
	}
	opts := options
	if opts == nil {
		opts = DefaultRenderOptions()
	}
	_ = opts

	// TODO: Implement Word document rendering
	// This will be implemented in task 2.x (word/renderer.go)
	return ErrNotImplemented
}

// RenderSpreadsheet renders an Excel spreadsheet to PDF format.
//
// The document must be a valid spreadsheet.Document that has been
// opened or created using the goffice spreadsheet package.
//
// By default, all sheets are rendered. Use RenderOptions to customize
// which sheets to include.
//
// The output is written to the provided io.Writer. If options is nil,
// DefaultRenderOptions() is used.
//
// Example:
//
//	doc, err := spreadsheet.Open("input.xlsx", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer doc.Close()
//
//	f, err := os.Create("output.pdf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer f.Close()
//
//	if err := pdf.RenderSpreadsheet(doc, f, nil); err != nil {
//	    log.Fatal(err)
//	}
func RenderSpreadsheet(
	doc *spreadsheet.Document,
	output io.Writer,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if output == nil {
		return ErrNilWriter
	}
	opts := options
	if opts == nil {
		opts = DefaultRenderOptions()
	}
	_ = opts

	// TODO: Implement spreadsheet rendering
	// This will be implemented in task 4.x (spreadsheet/renderer.go)
	return ErrNotImplemented
}

// RenderPresentation renders a PowerPoint presentation to PDF format.
//
// The document must be a valid presentation.Document that has been
// opened or created using the goffice presentation package.
//
// Each slide is rendered as a separate page in the PDF output.
// Animations are rendered as static snapshots.
//
// The output is written to the provided io.Writer. If options is nil,
// DefaultRenderOptions() is used.
//
// Example:
//
//	doc, err := presentation.Open("input.pptx", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer doc.Close()
//
//	f, err := os.Create("output.pdf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer f.Close()
//
//	if err := pdf.RenderPresentation(doc, f, nil); err != nil {
//	    log.Fatal(err)
//	}
func RenderPresentation(
	doc *presentation.Document,
	output io.Writer,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if output == nil {
		return ErrNilWriter
	}
	opts := options
	if opts == nil {
		opts = DefaultRenderOptions()
	}
	_ = opts

	// TODO: Implement presentation rendering
	// This will be implemented in task 5.x (presentation/renderer.go)
	return ErrNotImplemented
}

// RenderWordToFile is a convenience function that renders a Word document
// to a PDF file at the specified path.
//
// This is equivalent to creating a file and calling RenderWord.
func RenderWordToFile(
	doc *wordprocessing.Document,
	outputPath string,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if outputPath == "" {
		return errors.New(
			"pdf: output path is empty",
		)
	}
	_ = options

	// TODO: Implement file creation and rendering
	return ErrNotImplemented
}

// RenderSpreadsheetToFile is a convenience function that renders a
// spreadsheet to a PDF file at the specified path.
//
// This is equivalent to creating a file and calling RenderSpreadsheet.
func RenderSpreadsheetToFile(
	doc *spreadsheet.Document,
	outputPath string,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if outputPath == "" {
		return errors.New(
			"pdf: output path is empty",
		)
	}
	_ = options

	// TODO: Implement file creation and rendering
	return ErrNotImplemented
}

// RenderPresentationToFile is a convenience function that renders a
// presentation to a PDF file at the specified path.
//
// This is equivalent to creating a file and calling RenderPresentation.
func RenderPresentationToFile(
	doc *presentation.Document,
	outputPath string,
	options *RenderOptions,
) error {
	if doc == nil {
		return ErrNilDocument
	}
	if outputPath == "" {
		return errors.New(
			"pdf: output path is empty",
		)
	}
	_ = options

	// TODO: Implement file creation and rendering
	return ErrNotImplemented
}
