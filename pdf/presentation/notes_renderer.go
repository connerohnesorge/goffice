package presentation

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// NotesPageRenderer renders a notes page with slide thumbnail and speaker notes.
type NotesPageRenderer struct {
	doc       *presentation.Document
	slidePart *parts.SlidePart
	page      *core.Page
	pageSize  core.PageSize
}

// NewNotesPageRenderer creates a new notes page renderer.
func NewNotesPageRenderer(
	doc *presentation.Document,
	slidePart *parts.SlidePart,
	page *core.Page,
	pageSize core.PageSize,
) *NotesPageRenderer {
	return &NotesPageRenderer{
		doc:       doc,
		slidePart: slidePart,
		page:      page,
		pageSize:  pageSize,
	}
}

// Render renders the notes page.
func (nr *NotesPageRenderer) Render() error {
	// Write white background
	nr.writeWhiteBackground()

	// Calculate layout dimensions
	margin := 36.0 // 0.5 inch margins
	availWidth := nr.pageSize.Width - 2*margin
	availHeight := nr.pageSize.Height - 2*margin

	// Slide thumbnail area (top half)
	thumbnailHeight := availHeight * 0.4
	thumbnailWidth := availWidth

	// Notes area (bottom half)
	notesTop := margin + thumbnailHeight + 18 // 18pt gap
	notesHeight := availHeight - thumbnailHeight - 18

	// Render slide thumbnail
	if err := nr.renderSlideThumbnail(margin, notesTop+notesHeight, thumbnailWidth, thumbnailHeight); err != nil {
		return fmt.Errorf(
			"failed to render slide thumbnail: %w",
			err,
		)
	}

	// Render notes text
	if err := nr.renderNotesText(margin, notesTop, availWidth, notesHeight); err != nil {
		return fmt.Errorf(
			"failed to render notes text: %w",
			err,
		)
	}

	return nil
}

// writeWhiteBackground writes PDF operators for white background.
func (nr *NotesPageRenderer) writeWhiteBackground() {
	content := fmt.Sprintf(
		"q\n1 1 1 rg\n0 0 %.2f %.2f re\nf\nQ\n",
		nr.pageSize.Width,
		nr.pageSize.Height,
	)
	nr.page.WriteContentString(content)
}

// renderSlideThumbnail renders a placeholder for the slide thumbnail.
func (nr *NotesPageRenderer) renderSlideThumbnail(
	x, y, width, height float64,
) error {
	// Draw border
	content := fmt.Sprintf(
		"q\n0 0 0 RG\n0.5 w\n%.2f %.2f %.2f %.2f re\nS\nQ\n",
		x,
		y,
		width,
		height,
	)
	nr.page.WriteContentString(content)

	// TODO: Render actual slide thumbnail

	return nil
}

// renderNotesText renders placeholder text for notes.
func (nr *NotesPageRenderer) renderNotesText(
	x, y, width, height float64,
) error {
	// TODO: Extract and render notes text from notesPart
	// For now, just draw placeholder text

	content := fmt.Sprintf(
		"BT\n/F1 12 Tf\n%.2f %.2f Td\n(Speaker notes would appear here) Tj\nET\n",
		x,
		y+height-20,
	)
	nr.page.WriteContentString(content)

	return nil
}

// getSlideNumber returns the 1-based slide number.
func (nr *NotesPageRenderer) getSlideNumber() int {
	presPart := nr.doc.PresentationPart()
	if presPart == nil {
		return 1
	}

	slides := presPart.SlideParts()
	for i, slide := range slides {
		if slide == nr.slidePart {
			return i + 1
		}
	}

	return 1
}
