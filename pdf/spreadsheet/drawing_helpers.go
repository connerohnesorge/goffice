// drawing_helpers.go provides simplified drawing primitives for spreadsheet rendering.
// This is a temporary solution until the drawing package is fully integrated.

package spreadsheet

import (
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice-pdf/font"
)

// drawFillRect draws a filled rectangle on a page.
func drawFillRect(
	page *core.Page,
	x, y, width, height float64,
	color core.RGB,
	alpha float64,
) {
	// Set color
	page.WriteContentString(
		fmt.Sprintf(
			"%.3f %.3f %.3f rg\n",
			color.R,
			color.G,
			color.B,
		),
	)

	// Draw rectangle and fill
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f re f\n",
			x,
			y,
			width,
			height,
		),
	)
}

// drawStrokeRect draws a stroked rectangle on a page.
func drawStrokeRect(
	page *core.Page,
	x, y, width, height float64,
	color core.RGB,
	lineWidth float64,
) {
	// Set color
	page.WriteContentString(
		fmt.Sprintf(
			"%.3f %.3f %.3f RG\n",
			color.R,
			color.G,
			color.B,
		),
	)

	// Set line width
	page.WriteContentString(
		fmt.Sprintf("%.2f w\n", lineWidth),
	)

	// Draw rectangle and stroke
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f re S\n",
			x,
			y,
			width,
			height,
		),
	)
}

// drawLine draws a line on a page.
func drawLine(
	page *core.Page,
	x1, y1, x2, y2 float64,
	color core.RGB,
	lineWidth float64,
) {
	// Set color
	page.WriteContentString(
		fmt.Sprintf(
			"%.3f %.3f %.3f RG\n",
			color.R,
			color.G,
			color.B,
		),
	)

	// Set line width
	page.WriteContentString(
		fmt.Sprintf("%.2f w\n", lineWidth),
	)

	// Draw line
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f m %.2f %.2f l S\n",
			x1,
			y1,
			x2,
			y2,
		),
	)
}

// drawText draws text on a page at the specified position.
// The renderer is needed to properly register fonts in the PDF resources.
func drawText(
	renderer *SpreadsheetRenderer,
	page *core.Page,
	fontObj *font.Font,
	text string,
	x, y, fontSize float64,
	color core.RGB,
) error {
	if text == "" {
		return nil
	}

	// Track glyph usage for font subsetting
	if fontObj != nil && fontObj.Family != "" {
		renderer.trackGlyphUsage(
			fontObj.Family,
			text,
		)
	}

	// Register font in page resources and get the resource name
	fontName := renderer.registerFont(
		page,
		fontObj,
	)

	// Set text color
	page.WriteContentString(
		fmt.Sprintf(
			"%.3f %.3f %.3f rg\n",
			color.R,
			color.G,
			color.B,
		),
	)

	// Begin text object
	page.WriteContentString("BT\n")

	// Set font and size
	page.WriteContentString(
		fmt.Sprintf(
			"/%s %.2f Tf\n",
			fontName,
			fontSize,
		),
	)

	// Set text position
	page.WriteContentString(
		fmt.Sprintf("%.2f %.2f Td\n", x, y),
	)

	// Show text (escape special characters)
	escapedText := escapeTextForPDF(text)
	page.WriteContentString(
		fmt.Sprintf("(%s) Tj\n", escapedText),
	)

	// End text object
	page.WriteContentString("ET\n")

	return nil
}

// escapeTextForPDF escapes special characters in text for PDF content streams.
func escapeTextForPDF(text string) string {
	result := make([]byte, 0, len(text))

	for i := 0; i < len(text); i++ {
		ch := text[i]
		switch ch {
		case '(':
			result = append(result, '\\', '(')
		case ')':
			result = append(result, '\\', ')')
		case '\\':
			result = append(result, '\\', '\\')
		default:
			result = append(result, ch)
		}
	}

	return string(result)
}

// drawDataBar draws a data bar within a cell.
func drawDataBar(
	page *core.Page,
	x, y, width, height float64,
	percent float64,
	color core.RGB,
) {
	barWidth := width * percent
	padding := 2.0

	if barWidth > 0 {
		// Draw filled rectangle with 70% opacity
		fillColor := core.RGB{
			R: color.R * 0.7,
			G: color.G * 0.7,
			B: color.B * 0.7,
		}

		drawFillRect(
			page,
			x+padding,
			y+padding,
			barWidth-2*padding,
			height-2*padding,
			fillColor,
			0.7,
		)
	}
}

// drawCircle draws a filled circle (used for icon sets).
func drawCircle(
	page *core.Page,
	centerX, centerY, radius float64,
	color core.RGB,
) {
	// Set fill color
	page.WriteContentString(
		fmt.Sprintf(
			"%.3f %.3f %.3f rg\n",
			color.R,
			color.G,
			color.B,
		),
	)

	// Draw circle using bezier curves (approximation)
	k := 0.5522848 // Magic constant for circle approximation

	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f m\n",
			centerX,
			centerY-radius,
		),
	)
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f %.2f %.2f c\n",
			centerX+k*radius,
			centerY-radius,
			centerX+radius,
			centerY-k*radius,
			centerX+radius,
			centerY,
		),
	)
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f %.2f %.2f c\n",
			centerX+radius,
			centerY+k*radius,
			centerX+k*radius,
			centerY+radius,
			centerX,
			centerY+radius,
		),
	)
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f %.2f %.2f c\n",
			centerX-k*radius,
			centerY+radius,
			centerX-radius,
			centerY+k*radius,
			centerX-radius,
			centerY,
		),
	)
	page.WriteContentString(
		fmt.Sprintf(
			"%.2f %.2f %.2f %.2f %.2f %.2f c\n",
			centerX-radius,
			centerY-k*radius,
			centerX-k*radius,
			centerY-radius,
			centerX,
			centerY-radius,
		),
	)
	page.WriteContentString("f\n")
}
