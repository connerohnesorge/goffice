//nolint:revive // line-length-limit: test names can be long
package parts

import (
	"testing"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
)

// TestDiagramDataPart_ContentType tests that DiagramDataPart has the correct content type.
func TestDiagramDataPart_ContentType(
	t *testing.T,
) {
	part := &DiagramDataPart{}
	ct := part.FixedContentType()
	if ct != ContentTypeDiagramData {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			ct,
			ContentTypeDiagramData,
		)
	}
}

// TestDiagramDataPart_DataModel tests that DataModel returns nil when no root is set.
func TestDiagramDataPart_DataModel(t *testing.T) {
	// Create a part with proper OpenXmlPartData initialization
	part := &DiagramDataPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}
	dataModel := part.DataModel()
	if dataModel != nil {
		t.Errorf(
			"DataModel() = %v, want nil (no root set)",
			dataModel,
		)
	}
}

// TestDiagramDataPart_DataModelWithRoot tests DataModel with a root element set.
func TestDiagramDataPart_DataModelWithRoot(
	t *testing.T,
) {
	// Create a minimal part structure for testing
	part := &DiagramDataPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}

	// Create and set a data model root
	dm := diagram.NewDataModelRoot()
	part.SetRootElement(dm)

	// Retrieve and verify
	retrieved := part.DataModel()
	if retrieved == nil {
		t.Fatal(
			"DataModel() returned nil after SetRootElement",
		)
	}
	if retrieved != dm {
		t.Error(
			"DataModel() returned different element than what was set",
		)
	}
}

// TestDiagramLayoutDefinitionPart_ContentType tests layout definition part content type.
func TestDiagramLayoutDefinitionPart_ContentType(
	t *testing.T,
) {
	part := &DiagramLayoutDefinitionPart{}
	ct := part.FixedContentType()
	if ct != ContentTypeDiagramLayoutDef {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			ct,
			ContentTypeDiagramLayoutDef,
		)
	}
}

// TestDiagramLayoutDefinitionPart_LayoutDefinition tests LayoutDefinition method.
func TestDiagramLayoutDefinitionPart_LayoutDefinition(
	t *testing.T,
) {
	part := &DiagramLayoutDefinitionPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}
	layout := part.LayoutDefinition()
	if layout != nil {
		t.Errorf(
			"LayoutDefinition() = %v, want nil (no root set)",
			layout,
		)
	}
}

// TestDiagramLayoutDefinitionPart_LayoutDefinitionWithRoot tests with root set.
func TestDiagramLayoutDefinitionPart_LayoutDefinitionWithRoot(
	t *testing.T,
) {
	part := &DiagramLayoutDefinitionPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}

	// Create and set a layout definition
	layout := diagram.NewLayoutDefinition()
	part.SetRootElement(layout)

	// Retrieve and verify
	retrieved := part.LayoutDefinition()
	if retrieved == nil {
		t.Fatal(
			"LayoutDefinition() returned nil after SetRootElement",
		)
	}
	if retrieved != layout {
		t.Error(
			"LayoutDefinition() returned different element than what was set",
		)
	}
}

// TestDiagramStylePart_ContentType tests style part content type.
func TestDiagramStylePart_ContentType(
	t *testing.T,
) {
	part := &DiagramStylePart{}
	ct := part.FixedContentType()
	if ct != ContentTypeDiagramStyle {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			ct,
			ContentTypeDiagramStyle,
		)
	}
}

// TestDiagramStylePart_StyleDefinition tests StyleDefinition method.
func TestDiagramStylePart_StyleDefinition(
	t *testing.T,
) {
	part := &DiagramStylePart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}
	style := part.StyleDefinition()
	if style != nil {
		t.Errorf(
			"StyleDefinition() = %v, want nil (no root set)",
			style,
		)
	}
}

// TestDiagramStylePart_StyleDefinitionWithRoot tests with root set.
func TestDiagramStylePart_StyleDefinitionWithRoot(
	t *testing.T,
) {
	part := &DiagramStylePart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}

	// Create and set a style definition
	style := diagram.NewStyleDefinition()
	part.SetRootElement(style)

	// Retrieve and verify
	retrieved := part.StyleDefinition()
	if retrieved == nil {
		t.Fatal(
			"StyleDefinition() returned nil after SetRootElement",
		)
	}
	if retrieved != style {
		t.Error(
			"StyleDefinition() returned different element than what was set",
		)
	}
}

// TestDiagramColorsPart_ContentType tests colors part content type.
func TestDiagramColorsPart_ContentType(
	t *testing.T,
) {
	part := &DiagramColorsPart{}
	ct := part.FixedContentType()
	if ct != ContentTypeDiagramColors {
		t.Errorf(
			"FixedContentType() = %v, want %v",
			ct,
			ContentTypeDiagramColors,
		)
	}
}

// TestDiagramColorsPart_ColorsDefinition tests ColorsDefinition method.
func TestDiagramColorsPart_ColorsDefinition(
	t *testing.T,
) {
	part := &DiagramColorsPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}
	colors := part.ColorsDefinition()
	if colors != nil {
		t.Errorf(
			"ColorsDefinition() = %v, want nil (no root set)",
			colors,
		)
	}
}

// TestDiagramColorsPart_ColorsDefinitionWithRoot tests with root set.
func TestDiagramColorsPart_ColorsDefinitionWithRoot(
	t *testing.T,
) {
	part := &DiagramColorsPart{
		OpenXmlPartData: &openxml.OpenXmlPartData{},
	}

	// Create and set a colors definition
	colors := diagram.NewColorsDefinition()
	part.SetRootElement(colors)

	// Retrieve and verify
	retrieved := part.ColorsDefinition()
	if retrieved == nil {
		t.Fatal(
			"ColorsDefinition() returned nil after SetRootElement",
		)
	}
	if retrieved != colors {
		t.Error(
			"ColorsDefinition() returned different element than what was set",
		)
	}
}

// Note: Factory functions are tested indirectly through integration tests
// with actual containers. Testing them with nil containers would cause panics
// which is expected behavior when the factory is called incorrectly.
