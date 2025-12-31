// Go Generator for E2E Visual Testing
//
// This generator reads a test case JSON file and generates a PowerPoint presentation
// using the goffice library. It serves as one side of the visual comparison test.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

func main() {
	// Parse command line flags
	testCaseFile := flag.String(
		"test",
		"",
		"Path to test case JSON file",
	)
	outputFile := flag.String(
		"output",
		"",
		"Path to output PPTX file",
	)
	verbose := flag.Bool(
		"verbose",
		false,
		"Enable verbose logging",
	)
	flag.Parse()

	if *testCaseFile == "" || *outputFile == "" {
		log.Fatal(
			"Usage: e2e-go-generator -test <testcase.json> -output <output.pptx>",
		)
	}

	// Enable verbose logging if requested
	if *verbose {
		log.SetFlags(
			log.Ldate | log.Ltime | log.Lshortfile,
		)
		log.Printf(
			"Loading test case from: %s",
			*testCaseFile,
		)
	}

	// Load test case from JSON
	data, err := os.ReadFile(*testCaseFile)
	if err != nil {
		log.Fatalf(
			"Failed to read test case file: %v",
			err,
		)
	}

	tc, err := framework.FromJSON(data)
	if err != nil {
		log.Fatalf(
			"Failed to parse test case JSON: %v",
			err,
		)
	}

	if *verbose {
		log.Printf(
			"Loaded test case: %s (%s)",
			tc.Name,
			tc.ID,
		)
		log.Printf(
			"Category: %s, Tags: %v",
			tc.Category,
			tc.Tags,
		)
		log.Printf(
			"Slides: %d",
			tc.Spec.SlideCount,
		)
	}

	// Generate presentation
	if err := generatePresentation(tc, *outputFile, *verbose); err != nil {
		log.Fatalf(
			"Failed to generate presentation: %v",
			err,
		)
	}

	fmt.Printf(
		"✓ Successfully generated: %s\n",
		*outputFile,
	)
}

// generatePresentation creates the PPTX from test case specification
func generatePresentation(
	tc *framework.TestCase,
	outputPath string,
	verbose bool,
) error {
	// Create new presentation document
	doc, err := presentation.New(
		outputPath,
		presentation.DocTypePresentation,
	)
	if err != nil {
		return fmt.Errorf(
			"create presentation document: %w",
			err,
		)
	}
	defer doc.Close()

	if verbose {
		log.Printf(
			"Created presentation document",
		)
	}

	// Apply slide size if specified
	if tc.Spec.SlideSize.Width > 0 &&
		tc.Spec.SlideSize.Height > 0 {
		if err := applySlideSize(doc, tc.Spec.SlideSize); err != nil {
			return fmt.Errorf(
				"apply slide size: %w",
				err,
			)
		}
		if verbose {
			log.Printf(
				"Set slide size: %d x %d EMUs",
				tc.Spec.SlideSize.Width,
				tc.Spec.SlideSize.Height,
			)
		}
	}

	// Generate each slide
	for i, slideSpec := range tc.Spec.Slides {
		if verbose {
			log.Printf(
				"Generating slide %d/%d (index %d)",
				i+1,
				len(tc.Spec.Slides),
				slideSpec.Index,
			)
		}
		if err := generateSlide(doc, slideSpec, verbose); err != nil {
			return fmt.Errorf(
				"generate slide %d: %w",
				slideSpec.Index,
				err,
			)
		}
	}

	// Save presentation
	if err := doc.Save(); err != nil {
		return fmt.Errorf(
			"save presentation: %w",
			err,
		)
	}

	if verbose {
		log.Printf(
			"Presentation saved to: %s",
			outputPath,
		)
	}

	return nil
}

// applySlideSize sets the slide dimensions
func applySlideSize(
	doc *presentation.Document,
	size framework.SlideSize,
) error {
	presPart := doc.PresentationPart()
	if presPart == nil {
		return fmt.Errorf(
			"presentation part is nil",
		)
	}

	pres := presPart.Presentation()
	if pres == nil {
		return fmt.Errorf("presentation is nil")
	}

	slideSize := pres.GetOrCreateSlideSize()
	slideSize.SetCx(int(size.Width))
	slideSize.SetCy(int(size.Height))

	return nil
}

// generateSlide creates a single slide with all its elements
func generateSlide(
	doc *presentation.Document,
	spec framework.SlideSpec,
	verbose bool,
) error {
	// Add new slide
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("add slide: %w", err)
	}

	slide := slidePart.Slide()
	if slide == nil {
		return fmt.Errorf("slide is nil")
	}

	// Apply background if specified
	if spec.Background != nil {
		if verbose {
			log.Printf("  Applying background")
		}
		if err := applyBackground(slide, spec.Background); err != nil {
			return fmt.Errorf(
				"apply background: %w",
				err,
			)
		}
	}

	// Generate each element on the slide
	for i, elemSpec := range spec.Elements {
		if verbose {
			log.Printf(
				"  Generating element %d/%d (type: %s)",
				i+1,
				len(spec.Elements),
				elemSpec.Type,
			)
		}
		if err := generateElement(slidePart, slide, elemSpec, verbose); err != nil {
			return fmt.Errorf(
				"generate element %d (type %s): %w",
				i,
				elemSpec.Type,
				err,
			)
		}
	}

	return nil
}

// applyBackground applies background styling to a slide
func applyBackground(
	slide *elements.Slide,
	spec *framework.BackgroundSpec,
) error {
	// TODO: Implement background application
	// For now, return nil (blank background)
	return nil
}

// generateElement creates a single element (chart, shape, text, image, or table)
func generateElement(
	slidePart *parts.SlidePart,
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	switch spec.Type {
	case framework.ElementTypeChart:
		return generateChart(
			slidePart,
			slide,
			spec,
			verbose,
		)
	case framework.ElementTypeShape:
		return generateShape(slide, spec, verbose)
	case framework.ElementTypeText:
		return generateText(slide, spec, verbose)
	case framework.ElementTypeImage:
		return generateImage(
			slidePart,
			slide,
			spec,
			verbose,
		)
	case framework.ElementTypeTable:
		return generateTable(slide, spec, verbose)
	default:
		return fmt.Errorf(
			"unknown element type: %s",
			spec.Type,
		)
	}
}

// generateChart creates a chart element on the slide
func generateChart(
	slidePart *parts.SlidePart,
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	if spec.Chart == nil {
		return fmt.Errorf("chart spec is nil")
	}

	// Create chart generator
	chartGen := NewChartGenerator(
		slidePart,
		spec.Chart,
		verbose,
	)

	// Generate the chart
	chartPart, err := chartGen.Generate()
	if err != nil {
		return fmt.Errorf(
			"generate chart: %w",
			err,
		)
	}

	// Add chart to slide
	// Get the relationship ID for the chart part
	relID := chartPart.RelationshipID()
	if relID == "" {
		return fmt.Errorf(
			"chart part has no relationship ID",
		)
	}

	// Create graphicFrame to hold the chart
	if err := addChartGraphicFrame(
		slide,
		relID,
		spec.Position,
		spec.Size,
	); err != nil {
		return fmt.Errorf(
			"add chart graphic frame: %w",
			err,
		)
	}

	return nil
}

// addChartGraphicFrame adds a graphic frame containing the chart to the slide
func addChartGraphicFrame(
	slide *elements.Slide,
	chartRelID string,
	pos framework.Position,
	size framework.Size,
) error {
	// Get or create the shape tree
	shapeTree := slide.GetOrCreateShapeTree()

	// Add a new graphic frame
	gf := shapeTree.AddGraphicFrame()
	if gf == nil {
		return fmt.Errorf(
			"failed to add graphic frame",
		)
	}

	// Set position and size
	gf.SetPosition(int(pos.X), int(pos.Y))
	gf.SetSize(int(size.Width), int(size.Height))

	// Link the graphic frame to the chart
	elements.LinkGraphicFrameToChart(
		gf,
		chartRelID,
	)

	return nil
}

func generateShape(
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	if spec.Shape == nil {
		return fmt.Errorf("shape spec is nil")
	}

	// Create shape generator
	shapeGen := NewShapeGenerator(
		spec.Shape,
		verbose,
	)

	// Generate the shape
	_, err := shapeGen.Generate(
		slide,
		spec.Position,
		spec.Size,
	)
	if err != nil {
		return fmt.Errorf(
			"generate shape: %w",
			err,
		)
	}

	return nil
}

func generateText(
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	if spec.Text == nil {
		return fmt.Errorf("text spec is nil")
	}

	// Create text generator
	textGen := NewTextGenerator(
		spec.Text,
		verbose,
	)

	// Generate the text
	_, err := textGen.Generate(
		slide,
		spec.Position,
		spec.Size,
	)
	if err != nil {
		return fmt.Errorf(
			"generate text: %w",
			err,
		)
	}

	return nil
}

func generateImage(
	slidePart *parts.SlidePart,
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	// TODO: Implement image generation (future phase)
	if verbose {
		log.Printf(
			"    Image generation not yet implemented",
		)
	}

	return fmt.Errorf(
		"image generation not yet implemented",
	)
}

func generateTable(
	slide *elements.Slide,
	spec framework.ElementSpec,
	verbose bool,
) error {
	// TODO: Implement table generation (future phase)
	if verbose {
		log.Printf(
			"    Table generation not yet implemented",
		)
	}

	return fmt.Errorf(
		"table generation not yet implemented",
	)
}
