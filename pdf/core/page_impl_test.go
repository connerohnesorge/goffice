package core

import (
	"strings"
	"testing"
)

func TestPageImpl_DrawRectangle(t *testing.T) {
	p := NewPageImpl(612, 792)

	// Draw a filled and stroked rectangle
	p.SetFillColor(1, 0, 0)   // Red fill
	p.SetStrokeColor(0, 0, 1) // Blue stroke
	p.SetLineWidth(2)
	p.DrawRectangle(
		100,
		100,
		200,
		150,
		true,
		true,
	)

	content := p.GetContent()

	// Verify operators are present
	if !strings.Contains(content, "1 0 0 rg") {
		t.Error(
			"Expected fill color operator 'rg'",
		)
	}
	if !strings.Contains(content, "0 0 1 RG") {
		t.Error(
			"Expected stroke color operator 'RG'",
		)
	}
	if !strings.Contains(content, "2 w") {
		t.Error(
			"Expected line width operator 'w'",
		)
	}
	if !strings.Contains(
		content,
		"100 100 200 150 re",
	) {
		t.Error(
			"Expected rectangle operator 're'",
		)
	}
	if !strings.Contains(content, "B\n") {
		t.Error(
			"Expected fill and stroke operator 'B'",
		)
	}
}

func TestPageImpl_DrawCircle(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetFillColor(0, 1, 0) // Green fill
	p.DrawCircle(300, 400, 50, true, false)

	content := p.GetContent()

	// Verify bezier curve operators are present
	if !strings.Contains(content, "m\n") {
		t.Error("Expected move-to operator 'm'")
	}
	if !strings.Contains(content, "c\n") {
		t.Error("Expected curve-to operator 'c'")
	}
	if !strings.Contains(content, "f\n") {
		t.Error("Expected fill operator 'f'")
	}
	// Should have 4 curve segments
	curveCount := strings.Count(content, "c\n")
	if curveCount != 4 {
		t.Errorf(
			"Expected 4 curve segments, got %d",
			curveCount,
		)
	}
}

func TestPageImpl_DrawEllipse(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetStrokeColor(1, 0, 1) // Magenta stroke
	p.SetLineWidth(3)
	p.DrawEllipse(300, 400, 100, 50, false, true)

	content := p.GetContent()

	// Verify operators
	if !strings.Contains(content, "1 0 1 RG") {
		t.Error("Expected stroke color operator")
	}
	if !strings.Contains(content, "3 w") {
		t.Error("Expected line width operator")
	}
	if !strings.Contains(content, "S\n") {
		t.Error("Expected stroke operator 'S'")
	}
	// Should have 4 curve segments
	curveCount := strings.Count(content, "c\n")
	if curveCount != 4 {
		t.Errorf(
			"Expected 4 curve segments, got %d",
			curveCount,
		)
	}
}

func TestPageImpl_StateOptimization(
	t *testing.T,
) {
	p := NewPageImpl(612, 792)

	// Set color once
	p.SetFillColor(1, 0, 0)
	// Set same color again - should be optimized away
	p.SetFillColor(1, 0, 0)
	// Set different color
	p.SetFillColor(0, 1, 0)

	content := p.GetContent()

	// Should only have 2 color operators (not 3)
	colorCount := strings.Count(content, "rg\n")
	if colorCount != 2 {
		t.Errorf(
			"Expected 2 fill color operators (with optimization), got %d",
			colorCount,
		)
	}
}

func TestPageImpl_LineDashPattern(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetLineDashPattern([]float64{5, 3, 1, 3}, 0)

	content := p.GetContent()

	if !strings.Contains(
		content,
		"[5 3 1 3] 0 d",
	) {
		t.Error("Expected dash pattern operator")
	}
}

func TestPageImpl_LineCap(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetLineCap(LineCapRound)

	content := p.GetContent()

	if !strings.Contains(content, "1 J") {
		t.Error(
			"Expected line cap operator with value 1 (round)",
		)
	}
}

func TestPageImpl_LineJoin(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetLineJoin(LineJoinBevel)

	content := p.GetContent()

	if !strings.Contains(content, "2 j") {
		t.Error(
			"Expected line join operator with value 2 (bevel)",
		)
	}
}

func TestPageImpl_GraphicsState(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SaveGraphicsState()
	p.SetFillColor(1, 0, 0)
	p.RestoreGraphicsState()

	content := p.GetContent()

	if !strings.Contains(content, "q\n") {
		t.Error(
			"Expected save graphics state operator 'q'",
		)
	}
	if !strings.Contains(content, "Q\n") {
		t.Error(
			"Expected restore graphics state operator 'Q'",
		)
	}
}

func TestPageImpl_Transform(t *testing.T) {
	p := NewPageImpl(612, 792)

	matrix := PDFMatrix{
		A: 1,
		B: 0,
		C: 0,
		D: 1,
		E: 100,
		F: 50,
	}
	p.Transform(matrix)

	content := p.GetContent()

	if !strings.Contains(
		content,
		"1 0 0 1 100 50 cm",
	) {
		t.Error(
			"Expected transformation matrix operator 'cm'",
		)
	}
}

func TestPageImpl_DrawText(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetFont("Helvetica", 12)
	p.DrawText(100, 700, "Hello World")

	content := p.GetContent()

	if !strings.Contains(
		content,
		"/Helvetica 12 Tf",
	) {
		t.Error("Expected font operator 'Tf'")
	}
	if !strings.Contains(content, "BT\n") {
		t.Error(
			"Expected begin text operator 'BT'",
		)
	}
	if !strings.Contains(content, "100 700 Td") {
		t.Error(
			"Expected text position operator 'Td'",
		)
	}
	if !strings.Contains(
		content,
		"(Hello World) Tj",
	) {
		t.Error(
			"Expected show text operator 'Tj'",
		)
	}
	if !strings.Contains(content, "ET\n") {
		t.Error("Expected end text operator 'ET'")
	}
}

func TestPageImpl_TextEscaping(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.SetFont("Helvetica", 12)
	p.DrawText(
		100,
		700,
		"Text with (parentheses) and \\backslash",
	)

	content := p.GetContent()

	// Check that special characters are escaped
	if !strings.Contains(content, "\\(") &&
		!strings.Contains(content, "\\)") {
		t.Error(
			"Expected parentheses to be escaped",
		)
	}
	if !strings.Contains(content, "\\\\") {
		t.Error(
			"Expected backslash to be escaped",
		)
	}
}

func TestPageImpl_WriteContent(t *testing.T) {
	p := NewPageImpl(612, 792)

	p.WriteContent("0.5 g")
	p.WriteContent("100 100 m\n200 200 l\nS")

	content := p.GetContent()

	if !strings.Contains(content, "0.5 g\n") {
		t.Error(
			"Expected custom content with newline",
		)
	}
	if !strings.Contains(content, "100 100 m") {
		t.Error("Expected custom path operators")
	}
}

func TestPageImpl_EmptyContent(t *testing.T) {
	p := NewPageImpl(612, 792)

	content := p.GetContent()

	if content != "" {
		t.Errorf(
			"Expected empty content, got %q",
			content,
		)
	}
}

// mockPageImage is a test implementation of PageImage
type mockPageImage struct {
	width  int
	height int
	data   []byte
}

func (m *mockPageImage) Width() int { return m.width }

func (m *mockPageImage) Height() int { return m.height }

func (m *mockPageImage) Data() []byte { return m.data }

func TestPageImpl_AddImage(t *testing.T) {
	p := NewPageImpl(612, 792)

	img := &mockPageImage{
		width:  100,
		height: 100,
		data:   []byte{},
	}
	p.AddImage(img, 50, 50, 200, 200)

	content := p.GetContent()

	// Should have a comment (placeholder implementation)
	if !strings.Contains(content, "% Image:") {
		t.Error(
			"Expected image placeholder comment",
		)
	}
}

func TestMockPage_CallRecording(t *testing.T) {
	mock := NewMockPage()

	// Test drawing operations
	mock.DrawRectangle(
		10,
		20,
		100,
		50,
		true,
		true,
	)
	mock.DrawCircle(50, 50, 25, true, false)
	mock.DrawEllipse(
		100,
		100,
		50,
		30,
		false,
		true,
	)

	if mock.CallCount() != 3 {
		t.Errorf(
			"Expected 3 calls, got %d",
			mock.CallCount(),
		)
	}

	if !mock.HasCall(
		"DrawRectangle(10, 20, 100, 50, true, true)",
	) {
		t.Error(
			"Expected DrawRectangle call to be recorded",
		)
	}

	if !mock.HasCall(
		"DrawCircle(50, 50, 25, true, false)",
	) {
		t.Error(
			"Expected DrawCircle call to be recorded",
		)
	}

	if !mock.HasCall(
		"DrawEllipse(100, 100, 50, 30, false, true)",
	) {
		t.Error(
			"Expected DrawEllipse call to be recorded",
		)
	}
}

func TestMockPage_StateTracking(t *testing.T) {
	mock := NewMockPage()

	// Test color state
	mock.SetFillColor(0.5, 0.6, 0.7)
	if mock.FillR != 0.5 || mock.FillG != 0.6 ||
		mock.FillB != 0.7 {
		t.Errorf(
			"Fill color not tracked correctly: got (%g, %g, %g)",
			mock.FillR,
			mock.FillG,
			mock.FillB,
		)
	}

	mock.SetStrokeColor(0.1, 0.2, 0.3)
	if mock.StrokeR != 0.1 ||
		mock.StrokeG != 0.2 ||
		mock.StrokeB != 0.3 {
		t.Errorf(
			"Stroke color not tracked correctly: got (%g, %g, %g)",
			mock.StrokeR,
			mock.StrokeG,
			mock.StrokeB,
		)
	}

	// Test line width
	mock.SetLineWidth(2.5)
	if mock.LineWidthVal != 2.5 {
		t.Errorf(
			"Line width not tracked correctly: got %g",
			mock.LineWidthVal,
		)
	}

	// Test font
	mock.SetFont("Helvetica", 12)
	if mock.FontName != "Helvetica" ||
		mock.FontSize != 12 {
		t.Errorf(
			"Font not tracked correctly: got %s %g",
			mock.FontName,
			mock.FontSize,
		)
	}
}

func TestMockPage_GraphicsStateDepth(
	t *testing.T,
) {
	mock := NewMockPage()

	if mock.GraphicsStateDepth != 0 {
		t.Error(
			"Initial graphics state depth should be 0",
		)
	}

	mock.SaveGraphicsState()
	if mock.GraphicsStateDepth != 1 {
		t.Errorf(
			"Graphics state depth should be 1 after save, got %d",
			mock.GraphicsStateDepth,
		)
	}

	mock.SaveGraphicsState()
	if mock.GraphicsStateDepth != 2 {
		t.Errorf(
			"Graphics state depth should be 2 after second save, got %d",
			mock.GraphicsStateDepth,
		)
	}

	mock.RestoreGraphicsState()
	if mock.GraphicsStateDepth != 1 {
		t.Errorf(
			"Graphics state depth should be 1 after restore, got %d",
			mock.GraphicsStateDepth,
		)
	}

	mock.RestoreGraphicsState()
	if mock.GraphicsStateDepth != 0 {
		t.Errorf(
			"Graphics state depth should be 0 after second restore, got %d",
			mock.GraphicsStateDepth,
		)
	}
}

func TestMockPage_LastCall(t *testing.T) {
	mock := NewMockPage()

	if mock.LastCall() != "" {
		t.Error(
			"LastCall should be empty for new mock",
		)
	}

	mock.SetFillColor(1, 0, 0)
	if mock.LastCall() != "SetFillColor(1, 0, 0)" {
		t.Errorf(
			"LastCall incorrect: got %q",
			mock.LastCall(),
		)
	}

	mock.DrawRectangle(
		10,
		20,
		30,
		40,
		true,
		false,
	)
	if mock.LastCall() != "DrawRectangle(10, 20, 30, 40, true, false)" {
		t.Errorf(
			"LastCall incorrect: got %q",
			mock.LastCall(),
		)
	}
}

func TestMockPage_Reset(t *testing.T) {
	mock := NewMockPage()

	// Add some state
	mock.SetFillColor(1, 0, 0)
	mock.SetStrokeColor(0, 1, 0)
	mock.SetLineWidth(5)
	mock.SetFont("Times", 14)
	mock.SaveGraphicsState()
	mock.DrawRectangle(10, 20, 30, 40, true, true)

	if mock.CallCount() != 6 {
		t.Errorf(
			"Expected 6 calls before reset, got %d",
			mock.CallCount(),
		)
	}

	// Reset
	mock.Reset()

	// Verify everything is cleared
	if mock.CallCount() != 0 {
		t.Errorf(
			"Expected 0 calls after reset, got %d",
			mock.CallCount(),
		)
	}

	if mock.FillR != 0 || mock.FillG != 0 ||
		mock.FillB != 0 {
		t.Error("Fill color should be reset to 0")
	}

	if mock.StrokeR != 0 || mock.StrokeG != 0 ||
		mock.StrokeB != 0 {
		t.Error(
			"Stroke color should be reset to 0",
		)
	}

	if mock.LineWidthVal != 0 {
		t.Error("Line width should be reset to 0")
	}

	if mock.FontName != "" || mock.FontSize != 0 {
		t.Error("Font should be reset")
	}

	if mock.GraphicsStateDepth != 0 {
		t.Error(
			"Graphics state depth should be reset to 0",
		)
	}

	if mock.LastCall() != "" {
		t.Error(
			"LastCall should be empty after reset",
		)
	}
}

func TestMockPage_ComplexScenario(t *testing.T) {
	mock := NewMockPage()

	// Simulate a complex drawing scenario
	mock.SaveGraphicsState()
	mock.SetFillColor(1, 0, 0)
	mock.SetStrokeColor(0, 0, 1)
	mock.SetLineWidth(2)
	mock.DrawRectangle(
		10,
		20,
		100,
		50,
		true,
		true,
	)
	mock.RestoreGraphicsState()

	mock.SaveGraphicsState()
	mock.Transform(
		PDFMatrix{
			A: 1,
			B: 0,
			C: 0,
			D: 1,
			E: 50,
			F: 50,
		},
	)
	mock.DrawCircle(0, 0, 25, true, false)
	mock.RestoreGraphicsState()

	mock.SetFont("Helvetica", 12)
	mock.DrawText(100, 700, "Hello World")

	// Verify the sequence
	expectedCalls := []string{
		"SaveGraphicsState()",
		"SetFillColor(1, 0, 0)",
		"SetStrokeColor(0, 0, 1)",
		"SetLineWidth(2)",
		"DrawRectangle(10, 20, 100, 50, true, true)",
		"RestoreGraphicsState()",
		"SaveGraphicsState()",
		"Transform({1 0 0 1 50 50})",
		"DrawCircle(0, 0, 25, true, false)",
		"RestoreGraphicsState()",
		"SetFont(\"Helvetica\", 12)",
		"DrawText(100, 700, \"Hello World\")",
	}

	if len(mock.Calls) != len(expectedCalls) {
		t.Fatalf(
			"Expected %d calls, got %d",
			len(expectedCalls),
			len(mock.Calls),
		)
	}

	for i, expected := range expectedCalls {
		if mock.Calls[i] != expected {
			t.Errorf(
				"Call %d: expected %q, got %q",
				i,
				expected,
				mock.Calls[i],
			)
		}
	}

	// Verify final state
	if mock.GraphicsStateDepth != 0 {
		t.Error(
			"Graphics state should be balanced",
		)
	}

	if mock.FontName != "Helvetica" ||
		mock.FontSize != 12 {
		t.Error("Font state should be preserved")
	}
}
