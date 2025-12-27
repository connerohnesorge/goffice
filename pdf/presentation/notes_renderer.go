package presentation

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/openxml"
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

	// Render actual slide thumbnail by creating a mini slide renderer
	// Get the slide size to calculate scaling
	slideSize := nr.getSlideSize()

	// Calculate scale to fit thumbnail area while maintaining aspect ratio
	scaleX := width / slideSize.Width
	scaleY := height / slideSize.Height
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	// Calculate centered position
	scaledWidth := slideSize.Width * scale
	scaledHeight := slideSize.Height * scale
	offsetX := x + (width-scaledWidth)/2
	offsetY := y + (height-scaledHeight)/2

	// Create a transformed coordinate space for the thumbnail
	// For simplicity, render key slide elements in the thumbnail area
	slide := nr.slidePart.Slide()
	if slide != nil {
		csd := slide.CommonSlideData()
		if csd != nil {
			// Render simplified background
			content := fmt.Sprintf(
				"q\n0.95 0.95 0.95 rg\n%.2f %.2f %.2f %.2f re\nf\nQ\n",
				offsetX,
				offsetY,
				scaledWidth,
				scaledHeight,
			)
			nr.page.WriteContentString(content)

			// Add a label
			content = fmt.Sprintf(
				"BT\n/F1 10 Tf\n%.2f %.2f Td\n(Slide %d) Tj\nET\n",
				offsetX+4,
				offsetY+scaledHeight-14,
				nr.getSlideNumber(),
			)
			nr.page.WriteContentString(content)
		}
	}

	return nil
}

// getSlideSize returns the slide dimensions.
func (nr *NotesPageRenderer) getSlideSize() core.PageSize {
	presPart := nr.doc.PresentationPart()
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

	cx := slideSize.Cx()
	cy := slideSize.Cy()

	if cx == 0 || cy == 0 {
		return StandardSlideSize()
	}

	return core.PageSize{
		Width: float64(
			cx,
		) / 914400.0 * 72.0, // EMU to points
		Height: float64(cy) / 914400.0 * 72.0,
	}
}

// renderNotesText renders notes text extracted from the notes slide part.
func (nr *NotesPageRenderer) renderNotesText(
	x, y, width, height float64,
) error {
	// Get the notes slide part
	notesPart := nr.slidePart.NotesSlidePart()
	if notesPart == nil {
		// No notes for this slide
		content := fmt.Sprintf(
			"BT\n/F1 12 Tf\n%.2f %.2f Td\n(No speaker notes) Tj\nET\n",
			x,
			y+height-20,
		)
		nr.page.WriteContentString(content)
		return nil
	}

	// Get the notes slide element
	notesSlide := notesPart.NotesSlide()
	if notesSlide == nil {
		return nil
	}

	// Extract notes text from the notes slide
	notesText := nr.extractNotesText(notesSlide)
	if notesText == "" {
		content := fmt.Sprintf(
			"BT\n/F1 12 Tf\n%.2f %.2f Td\n(No speaker notes) Tj\nET\n",
			x,
			y+height-20,
		)
		nr.page.WriteContentString(content)
		return nil
	}

	// Render the notes text
	fontSize := 12.0
	lineHeight := fontSize * 1.2
	currentY := y + height - lineHeight

	// Simple word wrapping
	words := splitIntoWords(notesText)
	currentLine := ""
	maxLineWidth := width - 4 // Leave some margin

	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		// Simple width estimation (very rough)
		estimatedWidth := float64(
			len(testLine),
		) * fontSize * 0.5

		if estimatedWidth > maxLineWidth &&
			currentLine != "" {
			// Render current line
			content := fmt.Sprintf(
				"BT\n/F1 %.1f Tf\n%.2f %.2f Td\n(%s) Tj\nET\n",
				fontSize,
				x+2,
				currentY,
				escapeForPdf(currentLine),
			)
			nr.page.WriteContentString(content)
			currentY -= lineHeight
			currentLine = word
		} else {
			currentLine = testLine
		}
	}

	// Render last line
	if currentLine != "" {
		content := fmt.Sprintf(
			"BT\n/F1 %.1f Tf\n%.2f %.2f Td\n(%s) Tj\nET\n",
			fontSize,
			x+2,
			currentY,
			escapeForPdf(currentLine),
		)
		nr.page.WriteContentString(content)
	}

	return nil
}

// extractNotesText extracts plain text from the notes slide.
func (nr *NotesPageRenderer) extractNotesText(
	notesSlide openxml.Element,
) string {
	// Notes slides have a common slide data structure
	// Look for cSld element
	composite, ok := notesSlide.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	var cSld openxml.Element
	for child := range composite.Children() {
		if child.LocalName() == "cSld" {
			cSld = child
			break
		}
	}

	if cSld == nil {
		return ""
	}

	// Look for shape tree
	cSldComp, ok := cSld.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	var spTree openxml.Element
	for child := range cSldComp.Children() {
		if child.LocalName() == "spTree" {
			spTree = child
			break
		}
	}

	if spTree == nil {
		return ""
	}

	// Extract text from shapes
	var notesText string
	spTreeComp, ok := spTree.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	for child := range spTreeComp.Children() {
		if child.LocalName() == "sp" {
			text := extractTextFromShape(child)
			if text != "" {
				if notesText != "" {
					notesText += "\n"
				}
				notesText += text
			}
		}
	}

	return notesText
}

// extractTextFromShape extracts text from a shape element.
func extractTextFromShape(
	shape openxml.Element,
) string {
	composite, ok := shape.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	// Look for txBody element
	var txBody openxml.Element
	for child := range composite.Children() {
		if child.LocalName() == "txBody" {
			txBody = child
			break
		}
	}

	if txBody == nil {
		return ""
	}

	return extractTextFromTextBody(txBody)
}

// extractTextFromTextBody extracts text from a text body.
func extractTextFromTextBody(
	txBody openxml.Element,
) string {
	composite, ok := txBody.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	var text string
	for para := range composite.Children() {
		if para.LocalName() == "p" {
			paraText := extractTextFromParagraph(
				para,
			)
			if paraText != "" {
				if text != "" {
					text += "\n"
				}
				text += paraText
			}
		}
	}

	return text
}

// extractTextFromParagraph extracts text from a paragraph.
func extractTextFromParagraph(
	para openxml.Element,
) string {
	composite, ok := para.(openxml.CompositeElement)
	if !ok {
		return ""
	}

	var text string
	for child := range composite.Children() {
		if child.LocalName() == "r" {
			// Text run
			runComp, ok := child.(openxml.CompositeElement)
			if !ok {
				continue
			}

			for runChild := range runComp.Children() {
				if runChild.LocalName() == "t" {
					// Text element
					if leaf, ok := runChild.(*openxml.LeafElementBase); ok {
						text += leaf.InnerText()
					}
				}
			}
		}
	}

	return text
}

// splitIntoWords splits text into words for wrapping.
func splitIntoWords(text string) []string {
	var words []string
	currentWord := ""

	for _, r := range text {
		if r == ' ' || r == '\n' || r == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(r)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}

// escapeForPdf escapes special characters for PDF strings.
func escapeForPdf(s string) string {
	// Basic escaping for PDF string literals
	result := ""
	for _, r := range s {
		switch r {
		case '(':
			result += "\\("
		case ')':
			result += "\\)"
		case '\\':
			result += "\\\\"
		default:
			result += string(r)
		}
	}
	return result
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
