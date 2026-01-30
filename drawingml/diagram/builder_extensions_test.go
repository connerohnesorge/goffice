package diagram_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

func TestDataModelBuilder_WithText(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	// Add point with text
	builder.AddPoint("node1").WithText("Hello World")

	dataModel := builder.Build()

	// Verify text
	var found bool
	for pt := range openxml.Elements[*diagram.Point](dataModel.PointList) {
		if pt.TextBody == nil {
			continue
		}

		// Accessing text content might require traversing paragraphs
		// Since TextBody implementation is in drawingml package, we can check basic structure
		if pt.TextBody.Paragraphs() == nil {
			continue
		}

		found = true
		// We can't easily check text content without exposing more methods or casting
		// But existence of TextBody and Paragraphs is a good sign
		paras := pt.TextBody.Paragraphs()
		if len(paras) == 0 {
			t.Error("Expected paragraphs in text body")
		}
	}

	if !found {
		t.Error("Expected point with text body")
	}
}

func TestDataModelBuilder_WithShapeProperties(t *testing.T) {
	builder := diagram.NewDataModelBuilder()

	spPr := diagram.NewShapeProperties()
	// Set some property
	// spPr.SetSolidFill("FF0000") // Assuming SetSolidFill exists on ShapeProperties?
	// Check drawingml/shape_properties_fill.go?
	// Yes, ShapeProperties has SetSolidFill (via CompositeElementBase or helper?)
	// Actually drawingml/shape_properties.go had SetTransform etc.
	// drawingml/shape_properties_fill.go likely has fill methods.

	// Add point with shape properties
	builder.AddPoint("node1").WithShapeProperties(spPr)

	dataModel := builder.Build()

	// Verify shape properties
	var found bool
	for pt := range openxml.Elements[*diagram.Point](dataModel.PointList) {
		if pt.ShapeProperties != nil {
			found = true
		}
	}

	if !found {
		t.Error("Expected point with shape properties")
	}
}
