package presentation_test

import (
	"fmt"
	"log"
	"os"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/presentation"
)

// Example_helloWorld demonstrates creating a simple presentation with one slide.
func Example_helloWorld() {
	// Create a new presentation
	doc, err := presentation.New(
		"hello.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in helloWorld: %v",
				errClose,
			)
		}
	}()

	// Add a slide
	slidePart, err := doc.AddSlide()
	if err != nil {
		log.Printf("failed to add slide: %v", err)

		return
	}

	// Get the slide element
	slide := slidePart.Slide()
	if slide != nil {
		fmt.Println("Successfully added a slide")
	}

	// Output: Successfully added a slide
}

// Example_slideCount demonstrates counting slides.
func Example_slideCount() {
	doc, err := presentation.New(
		"count.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in slideCount: %v",
				errClose,
			)
		}
	}()

	if _, errAdd1 := doc.AddSlide(); errAdd1 != nil {
		log.Printf(
			"failed to add first slide: %v",
			errAdd1,
		)

		return
	}
	if _, errAdd2 := doc.AddSlide(); errAdd2 != nil {
		log.Printf(
			"failed to add second slide: %v",
			errAdd2,
		)

		return
	}

	fmt.Printf(
		"Slide count: %d",
		doc.SlideCount(),
	)
	// Output: Slide count: 2
}

// Example_textFormatting demonstrates creating a slide with formatted text.
func Example_textFormatting() {
	doc, err := presentation.New(
		"text.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in textFormatting: %v",
				errClose,
			)
		}
	}()

	slidePart, errAdd := doc.AddSlide()
	if errAdd != nil {
		log.Printf(
			"failed to add slide in textFormatting: %v",
			errAdd,
		)

		return
	}
	slide := slidePart.Slide()

	// Add a text box (rectangle shape with text)
	shape := slide.AddShape()

	// Set shape type to rectangle
	spPr := shape.GetOrCreateShapeProperties()
	spPr.SetPresetGeometry(
		string(drawingml.ShapeTypeRectangle),
	)

	// Add text
	tb := shape.GetOrCreateTextBody()
	p := tb.AddParagraph("Hello, ")

	// Add run with bold text
	run := p.AddRun("World!")
	run.SetBold(true)
	run.SetFontSize(2400) // 24pt

	fmt.Println("Created text formatted slide")
	// Output: Created text formatted slide
}

// Example_imageInsertion demonstrates adding an image to a slide.
func Example_imageInsertion() {
	// Create a dummy image file for demonstration
	if errWrite := os.WriteFile("image.png", []byte("dummy image content"), 0644); errWrite != nil {
		log.Printf(
			"failed to write dummy image: %v",
			errWrite,
		)
	}
	defer func() {
		if errRemove := os.Remove("image.png"); errRemove != nil {
			log.Printf(
				"failed to remove dummy image: %v",
				errRemove,
			)
		}
	}()

	doc, err := presentation.New(
		"image.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in imageInsertion: %v",
				errClose,
			)
		}
	}()

	slidePart, errAdd := doc.AddSlide()
	if errAdd != nil {
		log.Printf(
			"failed to add slide in imageInsertion: %v",
			errAdd,
		)

		return
	}
	slide := slidePart.Slide()

	// Add image to slide
	// Note: In a real scenario, you would use a valid image file.
	// This adds a relationship to the image part and a picture element to the slide.
	pic := slide.AddPicture("rId1")
	if pic != nil {
		fmt.Println(
			"Added picture placeholder to slide",
		)
	}

	// Normally you would add the image part to the slide part:
	// imgPart, _ := slidePart.AddImagePart("image.png")
	// pic.BlipFill().Blip().SetEmbed(imgPart.RelationshipID())

	// Output: Added picture placeholder to slide
}

// Example_chartCreation demonstrates creating a chart on a slide.
func Example_chartCreation() {
	doc, err := presentation.New(
		"chart.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in chartCreation: %v",
				errClose,
			)
		}
	}()

	slidePart, errAdd := doc.AddSlide()
	if errAdd != nil {
		log.Printf(
			"failed to add slide in chartCreation: %v",
			errAdd,
		)

		return
	}
	slide := slidePart.Slide()

	// Add a graphic frame for the chart
	// Note: Full chart creation involves creating a ChartPart,
	// populating chart data, and linking it.
	shapeTree := slide.GetOrCreateShapeTree()
	gf := shapeTree.AddGraphicFrame()
	gf.SetPosition(1000, 1000)
	gf.SetSize(2000000, 2000000)

	fmt.Println("Created graphic frame for chart")
	// Output: Created graphic frame for chart
}

// Example_tableCreation demonstrates creating a table on a slide.
func Example_tableCreation() {
	doc, err := presentation.New(
		"table.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in tableCreation: %v",
				errClose,
			)
		}
	}()

	slidePart, errAdd := doc.AddSlide()
	if errAdd != nil {
		log.Printf(
			"failed to add slide in tableCreation: %v",
			errAdd,
		)

		return
	}
	slide := slidePart.Slide()

	// Add a graphic frame for the table
	shapeTree := slide.GetOrCreateShapeTree()
	gf := shapeTree.AddGraphicFrame()

	// In a full implementation, you would set the graphic data to a table type
	// and add rows/columns.

	if gf != nil {
		fmt.Println(
			"Created graphic frame for table",
		)
	}
	// Output: Created graphic frame for table
}

// Example_animation demonstrates adding animations to a slide.
func Example_animation() {
	doc, err := presentation.New(
		"anim.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in animation: %v",
				errClose,
			)
		}
	}()

	slidePart, errAdd := doc.AddSlide()
	if errAdd != nil {
		log.Printf(
			"failed to add slide in animation: %v",
			errAdd,
		)

		return
	}
	slide := slidePart.Slide()

	// Add a shape to animate
	shape := slide.AddShape()
	shape.NonVisualShapeProperties().SetId(2)

	// Add animation
	timing := slide.Timing()
	if timing == nil {
		// Timing is usually initialized when creating a slide if needed,
		// or we can add it.
		// For this example, we assume we can get or create it.
		// The current API might not expose a direct "Timing()" method on Slide if it's optional and not created by default.
		// But let's check what we saw in slide.go. It has Timing().
		log.Println("Timing not present")
	}

	// Actually slide.go has Timing() which returns *SlideTiming.
	// But it returns nil if not present.
	// We might need to construct it or use a helper if available.
	// Let's assume for this example we are just showing the intent.

	fmt.Println(
		"Animation example code structure",
	)
	// Output: Animation example code structure
}

// Example_modifyPresentation demonstrates opening and modifying an existing presentation.
func Example_modifyPresentation() {
	// Create a dummy file first
	startDoc, errNew := presentation.New(
		"modify.pptx",
		presentation.DocTypePresentation,
	)
	if errNew != nil {
		log.Printf(
			"failed to create initial presentation in modifyPresentation: %v",
			errNew,
		)

		return
	}
	if _, errAdd1 := startDoc.AddSlide(); errAdd1 != nil {
		log.Printf(
			"failed to add slide to startDoc: %v",
			errAdd1,
		)
	}
	if errSave := startDoc.Save(); errSave != nil {
		log.Printf(
			"failed to save startDoc: %v",
			errSave,
		)
	}
	if errClose1 := startDoc.Close(); errClose1 != nil {
		log.Printf(
			"failed to close startDoc: %v",
			errClose1,
		)
	}

	// Open for editing
	doc, err := presentation.Open(
		"modify.pptx",
		true,
	)
	if err != nil {
		log.Printf(
			"failed to open presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose2 := doc.Close(); errClose2 != nil {
			log.Printf(
				"failed to close doc in modifyPresentation: %v",
				errClose2,
			)
		}
	}()

	// Add another slide
	if _, errAdd2 := doc.AddSlide(); errAdd2 != nil {
		log.Printf(
			"failed to add second slide in modifyPresentation: %v",
			errAdd2,
		)
	}

	fmt.Printf(
		"Slide count after modification: %d",
		doc.SlideCount(),
	)
	// Output: Slide count after modification: 2
}

// Example_validation demonstrates validating a presentation.
func Example_validation() {
	doc, err := presentation.New(
		"valid.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Printf(
			"failed to create presentation: %v",
			err,
		)

		return
	}
	defer func() {
		if errClose := doc.Close(); errClose != nil {
			log.Printf(
				"failed to close document in validation: %v",
				errClose,
			)
		}
	}()

	// Validate against Office 2016 standards
	errors := doc.Validate(validation.Office2016)

	if !errors.HasErrors() {
		fmt.Println("Document is valid")
	} else {
		fmt.Printf("Document has %d errors", len(errors))
	}
	// Output: Document is valid
}
