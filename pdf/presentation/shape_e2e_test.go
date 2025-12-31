// shape_e2e_test.go contains end-to-end tests for PowerPoint shape to PDF rendering.

package presentation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// TestShapeRendering_BasicShapes tests rendering a PowerPoint with basic shapes to PDF.
// This is an end-to-end test that validates the shape rendering integration.
func TestShapeRendering_BasicShapes(
	t *testing.T,
) {
	// Create a test presentation with basic shapes
	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"shapes_test.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add a slide with basic shapes
	slide, err := doc.AddSlide()
	if err != nil {
		t.Fatalf("Failed to add slide: %v", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		t.Fatal("Failed to get slide element")
	}

	// Add a rectangle shape with blue fill
	rect := slideElem.AddShape()
	rect.SetText("Rectangle Shape")
	rect.SetPosition(
		914400,
		457200,
	) // 1 inch from left, 0.5 inches from top
	rect.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	rect.SetShapeType(elements.ShapeTypeRectangle)
	rect.SetSolidFill("4472C4") // Blue

	// Add an ellipse shape with green fill
	ellipse := slideElem.AddShape()
	ellipse.SetText("Ellipse")
	ellipse.SetPosition(
		914400,
		2285700,
	) // 1 inch from left, 2.5 inches from top
	ellipse.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	ellipse.SetShapeType(
		elements.ShapeTypeEllipse,
	)
	ellipse.SetSolidFill("70AD47") // Green

	// Add a triangle shape with red fill
	triangle := slideElem.AddShape()
	triangle.SetText("Triangle")
	triangle.SetPosition(
		4572000,
		457200,
	) // 5 inches from left, 0.5 inches from top
	triangle.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	triangle.SetShapeType(
		elements.ShapeTypeTriangle,
	)
	triangle.SetSolidFill("E74C3C") // Red

	// Render to PDF
	renderer := NewPresentationRenderer(doc)

	pdfPath := filepath.Join(
		tmpDir,
		"shapes_test.pdf",
	)
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}

	// Verify PDF has content (non-zero size)
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	// Basic size check - PDF should be at least a few hundred bytes
	// (even a minimal PDF with empty page is around 600+ bytes)
	if stat.Size() < 500 {
		t.Errorf(
			"PDF file is suspiciously small (%d bytes), expected at least 500 bytes",
			stat.Size(),
		)
	}

	t.Logf("PDF size: %d bytes", stat.Size())
}

// TestShapeRendering_MultipleSlides tests rendering a presentation with shapes across multiple slides.
// This validates that shape rendering works correctly across different slides.
func TestShapeRendering_MultipleSlides(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"multi_shapes.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Create 3 slides with different shapes
	shapeTypes := []elements.ShapeType{
		elements.ShapeTypeRectangle,
		elements.ShapeTypeEllipse,
		elements.ShapeTypeDiamond,
	}
	colors := []string{
		"4472C4",
		"70AD47",
		"E74C3C",
	}

	for i := 0; i < 3; i++ {
		slide, err := doc.AddSlide()
		if err != nil {
			t.Fatalf(
				"Failed to add slide %d: %v",
				i+1,
				err,
			)
		}

		slideElem := slide.Slide()
		if slideElem == nil {
			t.Fatalf(
				"Failed to get slide element for slide %d",
				i+1,
			)
		}

		// Add a shape to this slide
		shape := slideElem.AddShape()
		shape.SetText(
			"Shape " + string(rune('A'+i)),
		)
		shape.SetPosition(
			2286000,
			1828800,
		) // Centered approximately
		shape.SetSize(
			3657600,
			1828800,
		) // 4 inches wide, 2 inches tall
		shape.SetShapeType(shapeTypes[i])
		shape.SetSolidFill(colors[i])
	}

	// Render to PDF
	renderer := NewPresentationRenderer(doc)

	pdfPath := filepath.Join(
		tmpDir,
		"multi_shapes.pdf",
	)
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	t.Logf(
		"PDF size for 3 slides: %d bytes",
		stat.Size(),
	)
}

// TestShapeRendering_EmptyPresentation tests rendering an empty presentation.
// This validates graceful handling when no shapes are present.
func TestShapeRendering_EmptyPresentation(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"empty.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add an empty slide
	_, err = doc.AddSlide()
	if err != nil {
		t.Fatalf("Failed to add slide: %v", err)
	}

	// Render to PDF
	renderer := NewPresentationRenderer(doc)

	pdfPath := filepath.Join(tmpDir, "empty.pdf")
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		t.Errorf("PDF file was not created")
	}
}

// TestShapeRendering_WithOptions tests shape rendering with custom options.
func TestShapeRendering_WithOptions(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"options.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Add 3 slides with shapes
	for i := 0; i < 3; i++ {
		slide, err := doc.AddSlide()
		if err != nil {
			t.Fatalf(
				"Failed to add slide: %v",
				err,
			)
		}

		slideElem := slide.Slide()
		if slideElem == nil {
			continue
		}

		// Add a simple shape
		shape := slideElem.AddShape()
		shape.SetText(
			"Slide " + string(rune('1'+i)),
		)
		shape.SetPosition(914400, 914400)
		shape.SetSize(7315200, 4572000)
		shape.SetShapeType(
			elements.ShapeTypeRectangle,
		)
		shape.SetSolidFill("4472C4")
	}

	// Render with custom options (render only slides 1 and 3)
	renderer := NewPresentationRenderer(doc)
	opts := DefaultRenderOptions()
	opts.SlideRange = []int{1, 3}
	renderer.SetOptions(opts)

	pdfPath := filepath.Join(
		tmpDir,
		"options.pdf",
	)
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	t.Logf(
		"PDF size with custom options: %d bytes",
		stat.Size(),
	)
}

// TestShapeRendering_VariousShapeTypes tests rendering various shape types.
// This validates that different preset geometries render correctly.
func TestShapeRendering_VariousShapeTypes(
	t *testing.T,
) {
	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"various_shapes.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Test various shape types
	shapeTypes := []struct {
		name      string
		shapeType elements.ShapeType
		color     string
	}{
		{
			"Rectangle",
			elements.ShapeTypeRectangle,
			"4472C4",
		},
		{
			"Ellipse",
			elements.ShapeTypeEllipse,
			"70AD47",
		},
		{
			"Triangle",
			elements.ShapeTypeTriangle,
			"E74C3C",
		},
		{
			"Diamond",
			elements.ShapeTypeDiamond,
			"F39C12",
		},
	}

	slide, err := doc.AddSlide()
	if err != nil {
		t.Fatalf("Failed to add slide: %v", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		t.Fatal("Failed to get slide element")
	}

	// Arrange shapes in a grid (2 columns, 2 rows)
	colWidth := 3657600  // 4 inches
	rowHeight := 1828800 // 2 inches
	marginX := 914400    // 1 inch
	marginY := 457200    // 0.5 inches
	spacing := 457200    // 0.5 inches

	for i, st := range shapeTypes {
		col := i % 2
		row := i / 2

		x := marginX + (colWidth+spacing)*col
		y := marginY + (rowHeight+spacing)*row

		shape := slideElem.AddShape()
		shape.SetText(st.name)
		shape.SetPosition(x, y)
		shape.SetSize(colWidth, rowHeight)
		shape.SetShapeType(st.shapeType)
		shape.SetSolidFill(st.color)
	}

	// Render to PDF
	renderer := NewPresentationRenderer(doc)

	pdfPath := filepath.Join(
		tmpDir,
		"various_shapes.pdf",
	)
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}
	if stat.Size() < 500 {
		t.Errorf(
			"PDF file is suspiciously small (%d bytes), expected at least 500 bytes",
			stat.Size(),
		)
	}

	t.Logf(
		"PDF size with various shapes: %d bytes",
		stat.Size(),
	)
}

// TestShapeRendering_ComplexSlide tests rendering a slide with multiple shapes.
// This validates that complex slide layouts render correctly.
func TestShapeRendering_ComplexSlide(
	t *testing.T,
) {
	if testing.Short() {
		t.Skip(
			"Skipping complex slide test in short mode",
		)
	}

	tmpDir := t.TempDir()
	pptxPath := filepath.Join(
		tmpDir,
		"complex.pptx",
	)

	doc, err := presentation.New(
		pptxPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	slide, err := doc.AddSlide()
	if err != nil {
		t.Fatalf("Failed to add slide: %v", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		t.Fatal("Failed to get slide element")
	}

	// Add a title bar
	titleBar := slideElem.AddShape()
	titleBar.SetText("Complex Slide Layout")
	titleBar.SetPosition(0, 0)
	titleBar.SetSize(
		9144000,
		914400,
	) // Full width, 1 inch tall
	titleBar.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	titleBar.SetSolidFill("2C3E50")

	// Add content shapes in various positions
	shapes := []struct {
		text      string
		x, y      int
		w, h      int
		shapeType elements.ShapeType
		color     string
	}{
		{
			"Header",
			914400,
			1371600,
			7315200,
			914400,
			elements.ShapeTypeRectangle,
			"3498DB",
		},
		{
			"Content 1",
			914400,
			2743200,
			3200400,
			1371600,
			elements.ShapeTypeRectangle,
			"E74C3C",
		},
		{
			"Content 2",
			4572000,
			2743200,
			3200400,
			1371600,
			elements.ShapeTypeEllipse,
			"2ECC71",
		},
		{
			"Footer",
			914400,
			4572000,
			7315200,
			685800,
			elements.ShapeTypeRectangle,
			"95A5A6",
		},
	}

	for _, s := range shapes {
		shape := slideElem.AddShape()
		shape.SetText(s.text)
		shape.SetPosition(s.x, s.y)
		shape.SetSize(s.w, s.h)
		shape.SetShapeType(s.shapeType)
		shape.SetSolidFill(s.color)
	}

	// Render to PDF
	renderer := NewPresentationRenderer(doc)

	pdfPath := filepath.Join(
		tmpDir,
		"complex.pdf",
	)
	if err := renderer.RenderToFile(pdfPath); err != nil {
		t.Fatalf("Failed to render PDF: %v", err)
	}

	// Verify PDF was created and has substantial content
	stat, err := os.Stat(pdfPath)
	if os.IsNotExist(err) {
		t.Errorf("PDF file was not created")
	}
	if err != nil {
		t.Fatalf(
			"Failed to stat PDF file: %v",
			err,
		)
	}
	if stat.Size() == 0 {
		t.Errorf(
			"PDF file is empty (0 bytes)",
		)
	}

	// Log the size for informational purposes
	t.Logf(
		"PDF size for complex slide: %d bytes",
		stat.Size(),
	)
}
