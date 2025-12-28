package parts

import (
	"testing"

	"github.com/connerohnesorge/goffice/presentation/elements"
)

func TestChartPartCreation(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Create a slide part
	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	// Add a chart part
	chartPart, err := slidePart.AddChartPart()
	if err != nil {
		t.Fatalf(
			"Failed to add chart part: %v",
			err,
		)
	}

	// Verify chart part was created
	if chartPart == nil {
		t.Fatal("Chart part is nil")
	}

	// Verify chart part has correct content type
	if chartPart.FixedContentType() != ContentTypeChart {
		t.Errorf(
			"Expected content type %s, got %s",
			ContentTypeChart,
			chartPart.FixedContentType(),
		)
	}

	// Verify chart part has a ChartSpace
	chartSpace := chartPart.ChartSpace()
	if chartSpace == nil {
		t.Fatal("ChartSpace is nil")
	}

	// Verify chart part has a relationship ID
	relID := chartPart.RelationshipID()
	if relID == "" {
		t.Error(
			"Chart part has no relationship ID",
		)
	}

	// Verify chart part is in slide part's child parts
	chartParts := slidePart.ChartParts()
	if len(chartParts) != 1 {
		t.Errorf(
			"Expected 1 chart part, got %d",
			len(chartParts),
		)
	}
}

func TestMultipleChartParts(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Create a slide part
	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	// Add multiple chart parts
	const numCharts = 3
	chartParts := make([]*ChartPart, numCharts)
	for i := range numCharts {
		chartPart, err := slidePart.AddChartPart()
		if err != nil {
			t.Fatalf(
				"Failed to add chart part %d: %v",
				i+1,
				err,
			)
		}
		chartParts[i] = chartPart
	}

	// Verify all chart parts were created with unique URIs
	uris := make(map[string]bool)
	for i, chartPart := range chartParts {
		uri := chartPart.URI()
		if uris[uri] {
			t.Errorf(
				"Chart part %d has duplicate URI: %s",
				i+1,
				uri,
			)
		}
		uris[uri] = true
	}

	// Verify slide part has all chart parts
	retrievedChartParts := slidePart.ChartParts()
	if len(retrievedChartParts) != numCharts {
		t.Errorf(
			"Expected %d chart parts, got %d",
			numCharts,
			len(retrievedChartParts),
		)
	}
}

func TestLinkGraphicFrameToChart(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(
		t,
	)
	defer cleanup()

	// Create a slide part
	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf(
			"Failed to create slide part: %v",
			err,
		)
	}

	// Add a chart part
	chartPart, err := slidePart.AddChartPart()
	if err != nil {
		t.Fatalf(
			"Failed to add chart part: %v",
			err,
		)
	}

	// Get the slide
	slide := slidePart.Slide()
	if slide == nil {
		t.Fatal("Slide is nil")
	}

	// Create a shape tree
	shapeTree := slide.GetOrCreateShapeTree()
	if shapeTree == nil {
		t.Fatal("ShapeTree is nil")
	}

	// Add a graphic frame
	gf := shapeTree.AddGraphicFrame()
	if gf == nil {
		t.Fatal("GraphicFrame is nil")
	}

	// Link the graphic frame to the chart part
	elements.LinkGraphicFrameToChart(
		gf,
		chartPart.RelationshipID(),
	)

	// Verify the graphic frame has the proper structure
	graphic := gf.GetElement(
		"graphic",
		elements.NamespaceDrawingML,
	)
	if graphic == nil {
		t.Fatal(
			"GraphicFrame does not have a:graphic element",
		)
	}
}
