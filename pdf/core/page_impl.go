package core

import (
	"fmt"
	"strings"
)

// PageImpl implements the PageDrawer interface using PDF content stream operators.
// It accumulates PDF drawing commands that can be written to a page's content stream.
type PageImpl struct {
	content    strings.Builder // Accumulates PDF operators
	pageWidth  float64
	pageHeight float64

	// Current graphics state (for optimization - avoid redundant operators)
	fillR, fillG, fillB       float64
	strokeR, strokeG, strokeB float64
	lineWidth                 float64

	// Font state
	currentFont     string
	currentFontSize float64

	// Track if values have been initialized
	fillColorSet   bool
	strokeColorSet bool
	lineWidthSet   bool
	fontSet        bool
}

// Compile-time check to ensure PageImpl implements PageDrawer.
var _ PageDrawer = (*PageImpl)(nil)

// NewPageImpl creates a new PageImpl with the specified page dimensions.
func NewPageImpl(
	pageWidth, pageHeight float64,
) *PageImpl {
	return &PageImpl{
		pageWidth:  pageWidth,
		pageHeight: pageHeight,
		// Initialize to invalid values so first set always writes
		fillR:     -1,
		fillG:     -1,
		fillB:     -1,
		strokeR:   -1,
		strokeG:   -1,
		strokeB:   -1,
		lineWidth: -1,
	}
}

// DrawRectangle draws a rectangle at (x, y) with the given width and height.
// Uses the PDF 're' (rectangle) operator for efficient rendering.
func (p *PageImpl) DrawRectangle(
	x, y, width, height float64,
	fill, stroke bool,
) {
	// Rectangle path operator: x y width height re
	p.content.WriteString(
		fmt.Sprintf("%s %s %s %s re\n",
			FormatFloat(x),
			FormatFloat(y),
			FormatFloat(width),
			FormatFloat(height)),
	)

	p.paintPath(fill, stroke)
}

// kappa is the magic number for approximating a circle with bezier curves.
// It's the distance of control points from the circle's edge for a perfect circular arc.
const kappa = 0.5522847498

// DrawCircle draws a circle centered at (cx, cy) with the given radius.
// Approximates the circle using 4 cubic bezier curves.
func (p *PageImpl) DrawCircle(
	cx, cy, radius float64,
	fill, stroke bool,
) {
	p.DrawEllipse(
		cx,
		cy,
		radius,
		radius,
		fill,
		stroke,
	)
}

// DrawEllipse draws an ellipse centered at (cx, cy) with radii rx and ry.
// Approximates the ellipse using 4 cubic bezier curves.
func (p *PageImpl) DrawEllipse(
	cx, cy, rx, ry float64,
	fill, stroke bool,
) {
	// Calculate control point offset for bezier approximation
	ox := rx * kappa
	oy := ry * kappa

	// Start at the rightmost point
	p.content.WriteString(fmt.Sprintf("%s %s m\n",
		FormatFloat(cx+rx),
		FormatFloat(cy)))

	// First curve (right to top)
	p.content.WriteString(
		fmt.Sprintf(
			"%s %s %s %s %s %s c\n",
			FormatFloat(
				cx+rx,
			),
			FormatFloat(cy+oy),
			FormatFloat(
				cx+ox,
			),
			FormatFloat(cy+ry),
			FormatFloat(cx),
			FormatFloat(cy+ry),
		),
	)

	// Second curve (top to left)
	p.content.WriteString(
		fmt.Sprintf(
			"%s %s %s %s %s %s c\n",
			FormatFloat(
				cx-ox,
			),
			FormatFloat(cy+ry),
			FormatFloat(
				cx-rx,
			),
			FormatFloat(cy+oy),
			FormatFloat(cx-rx),
			FormatFloat(cy),
		),
	)

	// Third curve (left to bottom)
	p.content.WriteString(
		fmt.Sprintf(
			"%s %s %s %s %s %s c\n",
			FormatFloat(
				cx-rx,
			),
			FormatFloat(cy-oy),
			FormatFloat(
				cx-ox,
			),
			FormatFloat(cy-ry),
			FormatFloat(cx),
			FormatFloat(cy-ry),
		),
	)

	// Fourth curve (bottom to right)
	p.content.WriteString(
		fmt.Sprintf(
			"%s %s %s %s %s %s c\n",
			FormatFloat(
				cx+ox,
			),
			FormatFloat(cy-ry),
			FormatFloat(
				cx+rx,
			),
			FormatFloat(cy-oy),
			FormatFloat(cx+rx),
			FormatFloat(cy),
		),
	)

	p.paintPath(fill, stroke)
}

// paintPath applies the appropriate painting operator based on fill/stroke flags.
func (p *PageImpl) paintPath(fill, stroke bool) {
	if fill && stroke {
		p.content.WriteString(
			"B\n",
		) // Fill and stroke
	} else if fill {
		p.content.WriteString("f\n") // Fill only
	} else if stroke {
		p.content.WriteString("S\n") // Stroke only
	}
}

// SetFillColor sets the fill color in RGB.
// Uses the PDF 'rg' operator (lowercase for fill color).
func (p *PageImpl) SetFillColor(r, g, b float64) {
	// Optimization: skip if color hasn't changed
	if p.fillColorSet && p.fillR == r &&
		p.fillG == g &&
		p.fillB == b {
		return
	}

	p.content.WriteString(
		fmt.Sprintf("%s %s %s rg\n",
			FormatFloat(r),
			FormatFloat(g),
			FormatFloat(b)),
	)

	p.fillR = r
	p.fillG = g
	p.fillB = b
	p.fillColorSet = true
}

// SetStrokeColor sets the stroke color in RGB.
// Uses the PDF 'RG' operator (uppercase for stroke color).
func (p *PageImpl) SetStrokeColor(
	r, g, b float64,
) {
	// Optimization: skip if color hasn't changed
	if p.strokeColorSet && p.strokeR == r &&
		p.strokeG == g &&
		p.strokeB == b {
		return
	}

	p.content.WriteString(
		fmt.Sprintf("%s %s %s RG\n",
			FormatFloat(r),
			FormatFloat(g),
			FormatFloat(b)),
	)

	p.strokeR = r
	p.strokeG = g
	p.strokeB = b
	p.strokeColorSet = true
}

// SetLineWidth sets the line width for stroking.
// Uses the PDF 'w' operator.
func (p *PageImpl) SetLineWidth(width float64) {
	// Optimization: skip if width hasn't changed
	if p.lineWidthSet && p.lineWidth == width {
		return
	}

	p.content.WriteString(
		fmt.Sprintf("%s w\n", FormatFloat(width)),
	)
	p.lineWidth = width
	p.lineWidthSet = true
}

// SetLineDashPattern sets the line dash pattern for stroking.
// Uses the PDF 'd' operator: [array] phase d
func (p *PageImpl) SetLineDashPattern(
	pattern []float64,
	phase float64,
) {
	p.content.WriteString("[")
	for i, v := range pattern {
		if i > 0 {
			p.content.WriteString(" ")
		}
		p.content.WriteString(FormatFloat(v))
	}
	p.content.WriteString(
		fmt.Sprintf(
			"] %s d\n",
			FormatFloat(phase),
		),
	)
}

// SetLineCap sets the line cap style.
// Uses the PDF 'J' operator (uppercase).
// 0 = butt cap, 1 = round cap, 2 = square cap
func (p *PageImpl) SetLineCap(lineCap LineCap) {
	p.content.WriteString(
		fmt.Sprintf("%d J\n", int(lineCap)),
	)
}

// SetLineJoin sets the line join style.
// Uses the PDF 'j' operator (lowercase).
// 0 = miter join, 1 = round join, 2 = bevel join
func (p *PageImpl) SetLineJoin(
	lineJoin LineJoin,
) {
	p.content.WriteString(
		fmt.Sprintf("%d j\n", int(lineJoin)),
	)
}

// SaveGraphicsState saves the current graphics state to the stack.
// Uses the PDF 'q' operator.
func (p *PageImpl) SaveGraphicsState() {
	p.content.WriteString("q\n")
}

// RestoreGraphicsState restores the graphics state from the stack.
// Uses the PDF 'Q' operator.
// Note: After restore, the tracked state values may no longer be accurate.
// Callers should re-set any state they need after restore.
func (p *PageImpl) RestoreGraphicsState() {
	p.content.WriteString("Q\n")
	// Reset tracking flags since state is now unknown
	p.fillColorSet = false
	p.strokeColorSet = false
	p.lineWidthSet = false
	p.fontSet = false
}

// Transform applies a transformation matrix to the current transformation matrix.
// Uses the PDF 'cm' operator (concat matrix).
// The matrix is [a b c d e f] which transforms (x, y) to:
//
//	x' = a*x + c*y + e
//	y' = b*x + d*y + f
func (p *PageImpl) Transform(matrix PDFMatrix) {
	p.content.WriteString(
		fmt.Sprintf("%s %s %s %s %s %s cm\n",
			FormatFloat(matrix.A),
			FormatFloat(matrix.B),
			FormatFloat(matrix.C),
			FormatFloat(matrix.D),
			FormatFloat(matrix.E),
			FormatFloat(matrix.F)),
	)
}

// AddImage adds an image to the page at the specified position and size.
// This is a placeholder implementation that logs the operation.
// Image handling requires XObject registration, which is complex and
// will be implemented in a future iteration.
func (p *PageImpl) AddImage(
	img PageImage,
	x, y, width, height float64,
) {
	// Placeholder: Image handling requires XObject resources
	// For now, just add a comment to the content stream
	p.content.WriteString(
		fmt.Sprintf(
			"%% Image: %dx%d at (%s, %s) size %sx%s\n",
			img.Width(),
			img.Height(),
			FormatFloat(x),
			FormatFloat(y),
			FormatFloat(width),
			FormatFloat(height),
		),
	)
}

// DrawText draws text at the specified position.
// Uses the PDF text operators: BT (begin text), Td (move), Tj (show), ET (end text).
func (p *PageImpl) DrawText(
	x, y float64,
	text string,
) {
	p.content.WriteString("BT\n")
	p.content.WriteString(
		fmt.Sprintf("%s %s Td\n",
			FormatFloat(x),
			FormatFloat(y)),
	)
	// Escape special characters in text
	escaped := escapeString(text)
	p.content.WriteString(
		fmt.Sprintf("(%s) Tj\n", escaped),
	)
	p.content.WriteString("ET\n")
}

// SetFont sets the current font and size.
// Uses the PDF 'Tf' operator.
// The font name should be a resource name from the page's font resources.
func (p *PageImpl) SetFont(
	name string,
	size float64,
) {
	// Optimization: skip if font hasn't changed
	if p.fontSet && p.currentFont == name &&
		p.currentFontSize == size {
		return
	}

	p.content.WriteString(
		fmt.Sprintf("/%s %s Tf\n",
			name,
			FormatFloat(size)),
	)

	p.currentFont = name
	p.currentFontSize = size
	p.fontSet = true
}

// WriteContent appends raw PDF content to the stream.
// This allows direct writing of PDF operators for advanced use cases.
func (p *PageImpl) WriteContent(content string) {
	p.content.WriteString(content)
	if !strings.HasSuffix(content, "\n") {
		p.content.WriteString("\n")
	}
}

// GetContent returns the accumulated PDF content stream.
func (p *PageImpl) GetContent() string {
	return p.content.String()
}

// Reset clears the content buffer and resets the graphics state.
// This is useful when reusing a PageImpl for multiple shapes.
func (p *PageImpl) Reset() {
	p.content.Reset()
	p.fillColorSet = false
	p.strokeColorSet = false
	p.lineWidthSet = false
	p.fontSet = false
}

// escapeString escapes special characters for PDF string literals.
func escapeString(s string) string {
	// Replace backslash first to avoid double-escaping
	s = strings.ReplaceAll(s, "\\", "\\\\")
	// Escape parentheses
	s = strings.ReplaceAll(s, "(", "\\(")
	s = strings.ReplaceAll(s, ")", "\\)")
	// Escape other special characters if needed
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\t", "\\t")

	return s
}
