package core

import "fmt"

// MockPage is a test implementation of PageDrawer that records all method calls.
type MockPage struct {
	Calls []string

	// State tracking for verification
	FillR, FillG, FillB       float64
	StrokeR, StrokeG, StrokeB float64
	LineWidthVal              float64
	FontName                  string
	FontSize                  float64
	GraphicsStateDepth        int
	LineDashPatternVal        []float64
	LineDashPhaseVal          float64
	LineCapVal                LineCap
	LineJoinVal               LineJoin
}

// Compile-time check to ensure MockPage implements PageDrawer.
var _ PageDrawer = (*MockPage)(nil)

// NewMockPage creates a new MockPage for testing.
func NewMockPage() *MockPage {
	return &MockPage{
		Calls: make([]string, 0),
	}
}

// DrawRectangle records a rectangle drawing operation.
func (m *MockPage) DrawRectangle(
	x, y, width, height float64,
	fill, stroke bool,
) {
	call := fmt.Sprintf(
		"DrawRectangle(%g, %g, %g, %g, %t, %t)",
		x,
		y,
		width,
		height,
		fill,
		stroke,
	)
	m.Calls = append(m.Calls, call)
}

// DrawCircle records a circle drawing operation.
func (m *MockPage) DrawCircle(
	cx, cy, radius float64,
	fill, stroke bool,
) {
	call := fmt.Sprintf(
		"DrawCircle(%g, %g, %g, %t, %t)",
		cx,
		cy,
		radius,
		fill,
		stroke,
	)
	m.Calls = append(m.Calls, call)
}

// DrawEllipse records an ellipse drawing operation.
func (m *MockPage) DrawEllipse(
	cx, cy, rx, ry float64,
	fill, stroke bool,
) {
	call := fmt.Sprintf(
		"DrawEllipse(%g, %g, %g, %g, %t, %t)",
		cx,
		cy,
		rx,
		ry,
		fill,
		stroke,
	)
	m.Calls = append(m.Calls, call)
}

// SetFillColor records a fill color setting and updates state.
func (m *MockPage) SetFillColor(r, g, b float64) {
	m.FillR = r
	m.FillG = g
	m.FillB = b
	call := fmt.Sprintf(
		"SetFillColor(%g, %g, %g)",
		r,
		g,
		b,
	)
	m.Calls = append(m.Calls, call)
}

// SetStrokeColor records a stroke color setting and updates state.
func (m *MockPage) SetStrokeColor(
	r, g, b float64,
) {
	m.StrokeR = r
	m.StrokeG = g
	m.StrokeB = b
	call := fmt.Sprintf(
		"SetStrokeColor(%g, %g, %g)",
		r,
		g,
		b,
	)
	m.Calls = append(m.Calls, call)
}

// SetLineWidth records a line width setting and updates state.
func (m *MockPage) SetLineWidth(width float64) {
	m.LineWidthVal = width
	call := fmt.Sprintf("SetLineWidth(%g)", width)
	m.Calls = append(m.Calls, call)
}

// SetLineDashPattern records a line dash pattern setting and updates state.
func (m *MockPage) SetLineDashPattern(
	pattern []float64,
	phase float64,
) {
	m.LineDashPatternVal = pattern
	m.LineDashPhaseVal = phase
	call := fmt.Sprintf(
		"SetLineDashPattern(%v, %g)",
		pattern,
		phase,
	)
	m.Calls = append(m.Calls, call)
}

// SetLineCap records a line cap setting and updates state.
func (m *MockPage) SetLineCap(lineCap LineCap) {
	m.LineCapVal = lineCap
	call := fmt.Sprintf("SetLineCap(%d)", lineCap)
	m.Calls = append(m.Calls, call)
}

// SetLineJoin records a line join setting and updates state.
func (m *MockPage) SetLineJoin(
	lineJoin LineJoin,
) {
	m.LineJoinVal = lineJoin
	call := fmt.Sprintf(
		"SetLineJoin(%d)",
		lineJoin,
	)
	m.Calls = append(m.Calls, call)
}

// SaveGraphicsState records a graphics state save and increments depth.
func (m *MockPage) SaveGraphicsState() {
	m.GraphicsStateDepth++
	call := "SaveGraphicsState()"
	m.Calls = append(m.Calls, call)
}

// RestoreGraphicsState records a graphics state restore and decrements depth.
func (m *MockPage) RestoreGraphicsState() {
	m.GraphicsStateDepth--
	call := "RestoreGraphicsState()"
	m.Calls = append(m.Calls, call)
}

// Transform records a transformation matrix operation.
func (m *MockPage) Transform(matrix PDFMatrix) {
	call := fmt.Sprintf("Transform(%v)", matrix)
	m.Calls = append(m.Calls, call)
}

// AddImage records an image addition operation.
func (m *MockPage) AddImage(
	img PageImage,
	x, y, width, height float64,
) {
	call := fmt.Sprintf(
		"AddImage(%p, %g, %g, %g, %g)",
		img,
		x,
		y,
		width,
		height,
	)
	m.Calls = append(m.Calls, call)
}

// DrawText records a text drawing operation.
func (m *MockPage) DrawText(
	x, y float64,
	text string,
) {
	call := fmt.Sprintf(
		"DrawText(%g, %g, %q)",
		x,
		y,
		text,
	)
	m.Calls = append(m.Calls, call)
}

// SetFont records a font setting and updates state.
func (m *MockPage) SetFont(
	name string,
	size float64,
) {
	m.FontName = name
	m.FontSize = size
	call := fmt.Sprintf(
		"SetFont(%q, %g)",
		name,
		size,
	)
	m.Calls = append(m.Calls, call)
}

// WriteContent records raw content stream writing.
func (m *MockPage) WriteContent(content string) {
	call := fmt.Sprintf(
		"WriteContent(%q)",
		content,
	)
	m.Calls = append(m.Calls, call)
}

// HasCall checks if a specific call was made.
func (m *MockPage) HasCall(call string) bool {
	for _, c := range m.Calls {
		if c == call {
			return true
		}
	}

	return false
}

// CallCount returns the total number of calls recorded.
func (m *MockPage) CallCount() int {
	return len(m.Calls)
}

// Reset clears all recorded calls and resets state.
func (m *MockPage) Reset() {
	m.Calls = make([]string, 0)
	m.FillR = 0
	m.FillG = 0
	m.FillB = 0
	m.StrokeR = 0
	m.StrokeG = 0
	m.StrokeB = 0
	m.LineWidthVal = 0
	m.FontName = ""
	m.FontSize = 0
	m.GraphicsStateDepth = 0
	m.LineDashPatternVal = nil
	m.LineDashPhaseVal = 0
	m.LineCapVal = LineCapButt
	m.LineJoinVal = LineJoinMiter
}

// LastCall returns the most recent call, or empty string if no calls.
func (m *MockPage) LastCall() string {
	if len(m.Calls) == 0 {
		return ""
	}

	return m.Calls[len(m.Calls)-1]
}
