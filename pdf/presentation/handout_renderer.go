package presentation

import (
	"fmt"
	"math"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// HandoutPageRenderer renders multiple slide thumbnails on a single page.
type HandoutPageRenderer struct {
	doc      *presentation.Document
	slides   []*parts.SlidePart
	page     *core.Page
	pageSize core.PageSize
	layout   HandoutLayout
}

// NewHandoutPageRenderer creates a new handout page renderer.
func NewHandoutPageRenderer(
	doc *presentation.Document,
	slides []*parts.SlidePart,
	page *core.Page,
	pageSize core.PageSize,
	layout HandoutLayout,
) *HandoutPageRenderer {
	return &HandoutPageRenderer{
		doc:      doc,
		slides:   slides,
		page:     page,
		pageSize: pageSize,
		layout:   layout,
	}
}

// Render renders the handout page.
func (hr *HandoutPageRenderer) Render() error {
	// Write white background
	hr.writeWhiteBackground()

	// Calculate grid layout based on handout type
	rows, cols := hr.getGridDimensions()

	// Calculate thumbnail dimensions
	margin := 36.0 // 0.5 inch margins
	gap := 18.0    // Gap between thumbnails

	availWidth := hr.pageSize.Width - 2*margin - float64(
		cols-1,
	)*gap
	availHeight := hr.pageSize.Height - 2*margin - float64(
		rows-1,
	)*gap

	thumbWidth := availWidth / float64(cols)
	thumbHeight := availHeight / float64(rows)

	// Render each slide thumbnail
	for i, slide := range hr.slides {
		row := i / cols
		col := i % cols

		x := margin + float64(
			col,
		)*(thumbWidth+gap)
		// PDF uses bottom-left origin, so flip Y
		y := hr.pageSize.Height - margin - float64(
			row+1,
		)*(thumbHeight+gap)

		if err := hr.renderSlideThumbnail(slide, x, y, thumbWidth, thumbHeight, i+1); err != nil {
			return fmt.Errorf(
				"failed to render slide %d thumbnail: %w",
				i+1,
				err,
			)
		}

		// For 3-slide handout, add lines for notes on the right
		if hr.layout == Handout3 && col == 0 {
			hr.renderNotesLines(
				x+thumbWidth+gap,
				y,
				availWidth-thumbWidth-gap,
				thumbHeight,
			)
		}
	}

	return nil
}

// writeWhiteBackground writes PDF operators for white background.
func (hr *HandoutPageRenderer) writeWhiteBackground() {
	content := fmt.Sprintf(
		"q\n1 1 1 rg\n0 0 %.2f %.2f re\nf\nQ\n",
		hr.pageSize.Width,
		hr.pageSize.Height,
	)
	hr.page.WriteContentString(content)
}

// getGridDimensions returns the number of rows and columns for the handout layout.
func (hr *HandoutPageRenderer) getGridDimensions() (rows, cols int) {
	switch hr.layout {
	case Handout1:
		return 1, 1
	case Handout2:
		return 2, 1 // Vertical arrangement
	case Handout3:
		return 3, 1 // Vertical with notes lines
	case Handout4:
		return 2, 2
	case Handout6:
		return 3, 2
	case Handout9:
		return 3, 3
	default:
		return 1, 1
	}
}

// renderSlideThumbnail renders a single slide thumbnail.
func (hr *HandoutPageRenderer) renderSlideThumbnail(
	slide *parts.SlidePart,
	x, y, width, height float64,
	slideNum int,
) error {
	// Draw border
	content := fmt.Sprintf(
		"q\n0 0 0 RG\n0.5 w\n%.2f %.2f %.2f %.2f re\nS\nQ\n",
		x,
		y,
		width,
		height,
	)
	hr.page.WriteContentString(content)

	// Draw slide number
	textX := x + width/2 - 10
	textY := y + height/2

	text := fmt.Sprintf(
		"BT\n/F1 14 Tf\n%.2f %.2f Td\n(%d) Tj\nET\n",
		textX,
		textY,
		slideNum,
	)
	hr.page.WriteContentString(text)

	return nil
}

// renderNotesLines renders horizontal lines for note-taking (3-slide handout only).
func (hr *HandoutPageRenderer) renderNotesLines(
	x, y, width, height float64,
) {
	content := "q\n0.7 0.7 0.7 RG\n0.25 w\n"

	// Draw horizontal lines
	numLines := int(
		math.Floor(height / 20),
	) // Line every 20 points
	for i := 0; i < numLines; i++ {
		lineY := y + height - float64(i)*20
		content += fmt.Sprintf(
			"%.2f %.2f m\n%.2f %.2f l\nS\n",
			x,
			lineY,
			x+width,
			lineY,
		)
	}

	content += "Q\n"
	hr.page.WriteContentString(content)
}
