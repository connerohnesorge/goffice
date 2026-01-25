//go:build ignore
// +build ignore

// Package main demonstrates PowerPoint presentation to PDF rendering with DrawingML shapes.
//
// This example shows how to:
// 1. Create a PowerPoint presentation with various shapes
// 2. Render the presentation to PDF using DrawingML shape rendering
// 3. Generate a professional-looking PDF with vector graphics
//
// The DrawingML rendering system handles:
// - Shape rendering (rectangles, ellipses, triangles, stars, arrows)
// - Fill colors and gradients
// - Stroke/outline rendering
// - Text within shapes
// - Professional layout and positioning
package main

import (
	"fmt"
	"log"

	pdfpresentation "github.com/connerohnesorge/goffice-pdf/presentation"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

const (
	pptxFilename = "powerpoint_shapes_example.pptx"
	pdfFilename  = "powerpoint_shapes_example.pdf"
)

func main() {
	fmt.Println(
		"Creating PowerPoint presentation with shapes...",
	)

	// Create a new presentation
	doc, err := presentation.New(
		pptxFilename,
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Fatalf(
			"Failed to create presentation: %v",
			err,
		)
	}
	defer doc.Close()

	// Create the slides and content
	if err := createPowerPointPresentation(doc); err != nil {
		log.Fatalf(
			"Failed to create PowerPoint presentation: %v",
			err,
		)
	}

	fmt.Printf("Created %s\n", pptxFilename)
	fmt.Println("Rendering to PDF...")

	// Render to PDF (using the open document, before closing)
	if err := renderToPDF(doc); err != nil {
		log.Fatalf(
			"Failed to render PDF: %v",
			err,
		)
	}

	fmt.Printf(
		"Successfully rendered to %s\n",
		pdfFilename,
	)
	fmt.Println(
		"\nThe PDF demonstrates DrawingML shape rendering:",
	)
	fmt.Println(
		"  - Basic shapes (rectangles, ellipses, triangles)",
	)
	fmt.Println(
		"  - Advanced shapes (stars, arrows, diamonds)",
	)
	fmt.Println("  - Text within shapes")
	fmt.Println("  - Various fill colors")
	fmt.Println(
		"  - Professional layout and positioning",
	)
}

// createPowerPointPresentation creates a PowerPoint presentation with various shapes.
func createPowerPointPresentation(doc *presentation.Document) error {
	// Slide 1: Basic Shapes
	if err := createBasicShapesSlide(doc); err != nil {
		return fmt.Errorf(
			"creating basic shapes slide: %w",
			err,
		)
	}

	// Slide 2: Advanced Shapes
	if err := createAdvancedShapesSlide(doc); err != nil {
		return fmt.Errorf(
			"creating advanced shapes slide: %w",
			err,
		)
	}

	// Slide 3: Shapes with Text
	if err := createTextShapesSlide(doc); err != nil {
		return fmt.Errorf(
			"creating text shapes slide: %w",
			err,
		)
	}

	if err := doc.Save(); err != nil {
		return fmt.Errorf(
			"saving presentation: %w",
			err,
		)
	}

	return nil
}

// createBasicShapesSlide creates a slide with basic shapes.
func createBasicShapesSlide(
	doc *presentation.Document,
) error {
	slide, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("adding slide: %w", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		return fmt.Errorf(
			"failed to get slide element",
		)
	}

	// Title
	title := slideElem.AddShape()
	title.SetText(
		"Basic Shapes - DrawingML Rendering",
	)
	title.SetPosition(
		457200,
		457200,
	) // 0.5 inches from left and top
	title.SetSize(
		8229600,
		914400,
	) // 9 inches wide, 1 inch tall
	title.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	title.SetSolidFill("2B579A") // Dark blue

	// Rectangle (Blue)
	rect := slideElem.AddShape()
	rect.SetText("Rectangle")
	rect.SetPosition(
		914400,
		1828800,
	) // 1 inch from left, 2 inches from top
	rect.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	rect.SetShapeType(elements.ShapeTypeRectangle)
	rect.SetSolidFill("4472C4") // Blue

	// Ellipse (Green)
	ellipse := slideElem.AddShape()
	ellipse.SetText("Ellipse")
	ellipse.SetPosition(
		4572000,
		1828800,
	) // 5 inches from left, 2 inches from top
	ellipse.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	ellipse.SetShapeType(
		elements.ShapeTypeEllipse,
	)
	ellipse.SetSolidFill("70AD47") // Green

	// Triangle (Red)
	triangle := slideElem.AddShape()
	triangle.SetText("Triangle")
	triangle.SetPosition(
		2743200,
		3657600,
	) // 3 inches from left, 4 inches from top
	triangle.SetSize(
		2743200,
		1371600,
	) // 3 inches wide, 1.5 inches tall
	triangle.SetShapeType(
		elements.ShapeTypeTriangle,
	)
	triangle.SetSolidFill("E74C3C") // Red

	return nil
}

// createAdvancedShapesSlide creates a slide with advanced shapes.
func createAdvancedShapesSlide(
	doc *presentation.Document,
) error {
	slide, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("adding slide: %w", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		return fmt.Errorf(
			"failed to get slide element",
		)
	}

	// Title
	title := slideElem.AddShape()
	title.SetText(
		"Advanced Shapes - Diamonds and Lines",
	)
	title.SetPosition(457200, 457200)
	title.SetSize(8229600, 914400)
	title.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	title.SetSolidFill("2B579A")

	// Diamond (Purple)
	diamond := slideElem.AddShape()
	diamond.SetText("Diamond")
	diamond.SetPosition(914400, 1828800)
	diamond.SetSize(
		2286000,
		2286000,
	) // Square for diamond
	diamond.SetShapeType(
		elements.ShapeTypeDiamond,
	)
	diamond.SetSolidFill("9B59B6") // Purple

	// Circle (Orange)
	circle := slideElem.AddShape()
	circle.SetText("Circle")
	circle.SetPosition(3657600, 1828800)
	circle.SetSize(2286000, 2286000)
	circle.SetShapeType(elements.ShapeTypeEllipse)
	circle.SetSolidFill("E67E22") // Orange

	// Line (Yellow)
	line := slideElem.AddShape()
	line.SetText("Line")
	line.SetPosition(6400800, 1828800)
	line.SetSize(2286000, 100000)
	line.SetShapeType(elements.ShapeTypeLine)
	line.SetSolidFill("F39C12") // Yellow

	// Rounded Rectangle (Teal)
	roundRect := slideElem.AddShape()
	roundRect.SetText("Rounded Rect")
	roundRect.SetPosition(2286000, 4571000)
	roundRect.SetSize(4572000, 1371600)
	roundRect.SetShapeType(
		elements.ShapeTypeRoundRectangle,
	)
	roundRect.SetSolidFill("1ABC9C") // Teal

	return nil
}

// createTextShapesSlide creates a slide with shapes containing formatted text.
func createTextShapesSlide(
	doc *presentation.Document,
) error {
	slide, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("adding slide: %w", err)
	}

	slideElem := slide.Slide()
	if slideElem == nil {
		return fmt.Errorf(
			"failed to get slide element",
		)
	}

	// Title
	title := slideElem.AddShape()
	title.SetText("Shapes with Text Content")
	title.SetPosition(457200, 457200)
	title.SetSize(8229600, 914400)
	title.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	title.SetSolidFill("2B579A")

	// Text Box 1
	textBox1 := slideElem.AddShape()
	textBox1.SetText(
		"Important Information\n\nThis is a text box with multiple lines.",
	)
	textBox1.SetPosition(914400, 1828800)
	textBox1.SetSize(3657600, 1828800)
	textBox1.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	textBox1.SetSolidFill("ECF0F1") // Light gray

	// Text Box 2
	textBox2 := slideElem.AddShape()
	textBox2.SetText(
		"Key Takeaways:\n• Point 1\n• Point 2\n• Point 3",
	)
	textBox2.SetPosition(5029200, 1828800)
	textBox2.SetSize(3657600, 1828800)
	textBox2.SetShapeType(
		elements.ShapeTypeRoundRectangle,
	)
	textBox2.SetSolidFill("D5DBDB") // Gray

	// Callout Box
	callout := slideElem.AddShape()
	callout.SetText(
		"Note: DrawingML shapes render with high fidelity in PDF!",
	)
	callout.SetPosition(2286000, 4114200)
	callout.SetSize(5486400, 1371600)
	callout.SetShapeType(
		elements.ShapeTypeRectangle,
	)
	callout.SetSolidFill("FFEB3B") // Yellow

	return nil
}

// renderToPDF renders the PowerPoint presentation to PDF.
func renderToPDF(doc *presentation.Document) error {
	// Create renderer using the already-open document
	renderer := pdfpresentation.NewPresentationRenderer(
		doc,
	)

	// Render to PDF
	if err := renderer.RenderToFile(pdfFilename); err != nil {
		return fmt.Errorf(
			"rendering PDF: %w",
			err,
		)
	}

	return nil
}
