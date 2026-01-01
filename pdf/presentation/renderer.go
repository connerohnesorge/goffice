// Package presentation provides PresentationML to PDF rendering capabilities.
// It converts PowerPoint presentations (.pptx) to PDF format, handling slide layouts,
// master slide inheritance, shapes, tables, charts, and various output modes.
package presentation

import (
	"fmt"
	"io"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// OutputMode specifies how slides should be rendered in the PDF.
type OutputMode int

const (
	// OutputModeSlides renders each slide as a full-page PDF page (default).
	OutputModeSlides OutputMode = iota
	// OutputModeNotes renders notes pages with slide thumbnail and speaker notes.
	OutputModeNotes
	// OutputModeHandouts renders multiple slides per page in handout layout.
	OutputModeHandouts
)

// String returns the string representation of the output mode.
func (m OutputMode) String() string {
	switch m {
	case OutputModeSlides:
		return "slides"
	case OutputModeNotes:
		return "notes"
	case OutputModeHandouts:
		return "handouts"
	default:
		return "slides"
	}
}

// HandoutLayout specifies the number of slides per handout page.
type HandoutLayout int

const (
	// Handout1 renders 1 slide per page.
	Handout1 HandoutLayout = 1
	// Handout2 renders 2 slides per page.
	Handout2 HandoutLayout = 2
	// Handout3 renders 3 slides per page (with lines for notes).
	Handout3 HandoutLayout = 3
	// Handout4 renders 4 slides per page.
	Handout4 HandoutLayout = 4
	// Handout6 renders 6 slides per page.
	Handout6 HandoutLayout = 6
	// Handout9 renders 9 slides per page.
	Handout9 HandoutLayout = 9
)

// RenderOptions specifies options for rendering a presentation to PDF.
type RenderOptions struct {
	// OutputMode specifies how slides should be rendered.
	OutputMode OutputMode

	// HandoutLayout specifies the layout for handout mode.
	// Only used when OutputMode is OutputModeHandouts.
	HandoutLayout HandoutLayout

	// SlideRange specifies which slides to render (1-based, inclusive).
	// If nil or empty, all slides are rendered.
	// Example: []int{1, 3, 5} renders only slides 1, 3, and 5.
	SlideRange []int

	// IncludeHiddenSlides includes slides marked as hidden.
	IncludeHiddenSlides bool

	// SkipAnimations renders static snapshots (animations not supported).
	// This is always true for PDF output, but kept for future compatibility.
	SkipAnimations bool
}

// DefaultRenderOptions returns default rendering options.
func DefaultRenderOptions() *RenderOptions {
	return &RenderOptions{
		OutputMode:          OutputModeSlides,
		HandoutLayout:       Handout1,
		IncludeHiddenSlides: false,
		SkipAnimations:      true,
	}
}

// PresentationRenderer renders PowerPoint presentations to PDF.
type PresentationRenderer struct {
	doc     *presentation.Document
	pdfDoc  *core.Document
	options *RenderOptions
}

// NewPresentationRenderer creates a new presentation renderer.
func NewPresentationRenderer(
	doc *presentation.Document,
) *PresentationRenderer {
	return &PresentationRenderer{
		doc:     doc,
		options: DefaultRenderOptions(),
	}
}

// SetOptions sets the rendering options.
func (r *PresentationRenderer) SetOptions(
	opts *RenderOptions,
) {
	if opts != nil {
		r.options = opts
	}
}

// Render renders the presentation to PDF and returns the PDF document.
// The caller is responsible for closing the PDF document.
func (r *PresentationRenderer) Render() (*core.Document, error) {
	// Create PDF document
	pdfDoc, err := core.NewDocument()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create PDF document: %w",
			err,
		)
	}
	r.pdfDoc = pdfDoc

	// Set PDF metadata from presentation properties
	if err := r.setMetadata(); err != nil {
		return nil, fmt.Errorf(
			"failed to set metadata: %w",
			err,
		)
	}

	// Get slides to render
	slides, err := r.getSlidesToRender()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get slides: %w",
			err,
		)
	}

	if len(slides) == 0 {
		return nil, fmt.Errorf(
			"no slides to render",
		)
	}

	// Render based on output mode
	switch r.options.OutputMode {
	case OutputModeSlides:
		err = r.renderSlides(slides)
	case OutputModeNotes:
		err = r.renderNotesPages(slides)
	case OutputModeHandouts:
		err = r.renderHandouts(slides)
	default:
		err = r.renderSlides(slides)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"failed to render presentation: %w",
			err,
		)
	}

	return pdfDoc, nil
}

// RenderToWriter renders the presentation to a writer.
func (r *PresentationRenderer) RenderToWriter(
	w io.Writer,
) error {
	pdfDoc, err := r.Render()
	if err != nil {
		return err
	}
	defer pdfDoc.Close()

	return pdfDoc.Write(w)
}

// RenderToFile renders the presentation to a file.
func (r *PresentationRenderer) RenderToFile(
	path string,
) error {
	pdfDoc, err := r.Render()
	if err != nil {
		return err
	}
	defer pdfDoc.Close()

	return pdfDoc.WriteToFile(path)
}

// setMetadata sets PDF metadata from presentation properties.
func (r *PresentationRenderer) setMetadata() error {
	coreProps := r.doc.CoreProperties()
	if coreProps == nil {
		return nil
	}

	metadata := core.Metadata{
		Title:    coreProps.Title(),
		Author:   coreProps.Creator(),
		Subject:  coreProps.Subject(),
		Keywords: coreProps.Keywords(),
		Creator:  "goffice-pdf presentation renderer",
		Producer: "goffice-pdf",
	}

	r.pdfDoc.SetMetadata(metadata)
	return nil
}

// getSlidesToRender returns the slides to render based on options.
func (r *PresentationRenderer) getSlidesToRender() ([]*parts.SlidePart, error) {
	presPart := r.doc.PresentationPart()
	if presPart == nil {
		return nil, fmt.Errorf(
			"presentation part not found",
		)
	}

	allSlides := presPart.SlideParts()
	if len(allSlides) == 0 {
		return nil, fmt.Errorf(
			"no slides found in presentation",
		)
	}

	// If slide range is specified, filter slides
	if len(r.options.SlideRange) > 0 {
		var filtered []*parts.SlidePart
		for _, idx := range r.options.SlideRange {
			if idx < 1 || idx > len(allSlides) {
				continue // Skip invalid indices
			}
			filtered = append(
				filtered,
				allSlides[idx-1],
			)
		}
		return filtered, nil
	}

	// Filter hidden slides if IncludeHiddenSlides is false
	if !r.options.IncludeHiddenSlides {
		var filtered []*parts.SlidePart
		for _, slidePart := range allSlides {
			slide := slidePart.Slide()
			if slide == nil {
				continue
			}
			// Check for show attribute (default is true/1)
			// If show="0" or show="false", the slide is hidden
			attr, found := slide.GetAttribute(
				"show",
				"",
			)
			if !found {
				// No show attribute means visible
				filtered = append(
					filtered,
					slidePart,
				)
				continue
			}
			attrVal := attr.Value()
			if attrVal != "0" &&
				attrVal != "false" {
				// show="1" or show="true" or any other value means visible
				filtered = append(
					filtered,
					slidePart,
				)
			}
		}
		return filtered, nil
	}

	return allSlides, nil
}

// getSlideSize returns the slide size from the presentation.
// Returns standard 4:3 size if not specified.
func (r *PresentationRenderer) getSlideSize() core.PageSize {
	presPart := r.doc.PresentationPart()
	if presPart == nil {
		return StandardSlideSize()
	}

	pres := presPart.Presentation()
	if pres == nil {
		return StandardSlideSize()
	}

	slideSize := pres.SlideSize()
	if slideSize == nil {
		return StandardSlideSize()
	}

	// Get slide dimensions in EMU and convert to points
	cx := slideSize.Cx()
	cy := slideSize.Cy()

	if cx == 0 || cy == 0 {
		return StandardSlideSize()
	}

	// Convert EMU to points
	width := drawingml.EmuToPoints(
		drawingml.EMU(cx),
	)
	height := drawingml.EmuToPoints(
		drawingml.EMU(cy),
	)

	return core.PageSize{
		Width:  width,
		Height: height,
	}
}

// renderSlides renders slides in standard slide mode.
func (r *PresentationRenderer) renderSlides(
	slides []*parts.SlidePart,
) error {
	slideSize := r.getSlideSize()

	for i, slidePart := range slides {
		// Create a new page for each slide
		page, err := r.pdfDoc.AddPage(slideSize)
		if err != nil {
			return fmt.Errorf(
				"failed to add page for slide %d: %w",
				i+1,
				err,
			)
		}

		// Create slide renderer
		slideRenderer := NewSlideRenderer(
			r.doc,
			r.pdfDoc,
			slidePart,
			page,
			slideSize,
		)

		// Render the slide
		if err := slideRenderer.Render(); err != nil {
			return fmt.Errorf(
				"failed to render slide %d: %w",
				i+1,
				err,
			)
		}
	}

	return nil
}

// renderNotesPages renders slides with speaker notes.
func (r *PresentationRenderer) renderNotesPages(
	slides []*parts.SlidePart,
) error {
	// Notes pages use standard page size (typically Letter or A4)
	pageSize := core.PageSizeLetter

	for i, slidePart := range slides {
		// Create a new page for each notes page
		page, err := r.pdfDoc.AddPage(pageSize)
		if err != nil {
			return fmt.Errorf(
				"failed to add notes page for slide %d: %w",
				i+1,
				err,
			)
		}

		// Create notes page renderer
		notesRenderer := NewNotesPageRenderer(
			r.doc,
			slidePart,
			page,
			pageSize,
		)

		// Render the notes page
		if err := notesRenderer.Render(); err != nil {
			return fmt.Errorf(
				"failed to render notes page %d: %w",
				i+1,
				err,
			)
		}
	}

	return nil
}

// renderHandouts renders multiple slides per page in handout layout.
func (r *PresentationRenderer) renderHandouts(
	slides []*parts.SlidePart,
) error {
	// Handouts typically use standard page size
	pageSize := core.PageSizeLetter
	layout := r.options.HandoutLayout
	if layout < 1 {
		layout = Handout1
	}

	// Calculate how many pages needed
	slidesPerPage := int(layout)
	numPages := (len(slides) + slidesPerPage - 1) / slidesPerPage

	for pageIdx := 0; pageIdx < numPages; pageIdx++ {
		// Create a new page
		page, err := r.pdfDoc.AddPage(pageSize)
		if err != nil {
			return fmt.Errorf(
				"failed to add handout page %d: %w",
				pageIdx+1,
				err,
			)
		}

		// Get slides for this page
		startIdx := pageIdx * slidesPerPage
		endIdx := startIdx + slidesPerPage
		if endIdx > len(slides) {
			endIdx = len(slides)
		}
		pageSlides := slides[startIdx:endIdx]

		// Create handout page renderer
		handoutRenderer := NewHandoutPageRenderer(
			r.doc,
			pageSlides,
			page,
			pageSize,
			layout,
		)

		// Render the handout page
		if err := handoutRenderer.Render(); err != nil {
			return fmt.Errorf(
				"failed to render handout page %d: %w",
				pageIdx+1,
				err,
			)
		}
	}

	return nil
}

// Standard PowerPoint slide sizes in points (72 points per inch).
// These are derived from the standard OOXML slide sizes in EMU.

// StandardSlideSize returns the standard 4:3 slide size (10" x 7.5").
func StandardSlideSize() core.PageSize {
	return core.PageSize{
		Width:  720, // 10 inches
		Height: 540, // 7.5 inches
	}
}

// WidescreenSlideSize returns the widescreen 16:9 slide size (10" x 5.625").
func WidescreenSlideSize() core.PageSize {
	return core.PageSize{
		Width:  720, // 10 inches
		Height: 405, // 5.625 inches
	}
}

// LetterSlideSize returns the letter-sized slide (8.5" x 11").
func LetterSlideSize() core.PageSize {
	return core.PageSize{
		Width:  612, // 8.5 inches
		Height: 792, // 11 inches
	}
}

// A4SlideSize returns the A4-sized slide (210mm x 297mm).
func A4SlideSize() core.PageSize {
	return core.PageSize{
		Width:  595.27, // 210mm
		Height: 841.89, // 297mm
	}
}
