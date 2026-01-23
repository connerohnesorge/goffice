package parts

import (
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
)

func TestDiagram_TemplatesExist(t *testing.T) {
	// Test that all template types exist
	templateTypes := []diagram.TemplateType{
		diagram.TemplateTypeList,
		diagram.TemplateTypeHierarchy,
	}

	for _, templateType := range templateTypes {
		t.Run(string(templateType), func(t *testing.T) {
			layoutXML := diagram.GetLayoutTemplate(templateType)
			if layoutXML == "" {
				t.Errorf("GetLayoutTemplate(%s) returned empty string", templateType)
			}
			if !strings.Contains(layoutXML, "<dgm:layoutDef") {
				t.Errorf("GetLayoutTemplate(%s) doesn't contain layoutDef element", templateType)
			}

			styleXML := diagram.GetStyleTemplate(templateType)
			if styleXML == "" {
				t.Errorf("GetStyleTemplate(%s) returned empty string", templateType)
			}
			if !strings.Contains(styleXML, "<dgm:styleDef") {
				t.Errorf("GetStyleTemplate(%s) doesn't contain styleDef element", templateType)
			}

			colorXML := diagram.GetColorTemplate(templateType)
			if colorXML == "" {
				t.Errorf("GetColorTemplate(%s) returned empty string", templateType)
			}
			if !strings.Contains(colorXML, "<dgm:colorsDef") {
				t.Errorf("GetColorTemplate(%s) doesn't contain colorsDef element", templateType)
			}
		})
	}
}

func TestDiagram_DefaultTemplateFallback(t *testing.T) {
	// Test that unknown template types fall back to list template
	unknownType := diagram.TemplateType("unknown")
	layoutXML := diagram.GetLayoutTemplate(unknownType)
	if layoutXML == "" {
		t.Error("GetLayoutTemplate with unknown type should return default template, not empty")
	}
	if !strings.Contains(layoutXML, "<dgm:layoutDef") {
		t.Error("GetLayoutTemplate with unknown type should return list template")
	}
}

func TestSlidePart_AddDiagram(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	// Test creating diagram with list template
	diagramResult, err := slidePart.AddDiagram(diagram.TemplateTypeList)
	if err != nil {
		t.Fatalf("AddDiagram(list) error = %v", err)
	}

	// Verify all parts are created
	if diagramResult.DataPart == nil {
		t.Error("AddDiagram() returned Diagram with nil DataPart")
	}
	if diagramResult.LayoutPart == nil {
		t.Error("AddDiagram() returned Diagram with nil LayoutPart")
	}
	if diagramResult.StylePart == nil {
		t.Error("AddDiagram() returned Diagram with nil StylePart")
	}
	if diagramResult.ColorsPart == nil {
		t.Error("AddDiagram() returned Diagram with nil ColorsPart")
	}

	// Verify parts have root elements
	if diagramResult.DataPart.DataModel() == nil {
		t.Error("Diagram DataPart has no DataModel root element")
	}
	if diagramResult.LayoutPart.LayoutDefinition() == nil {
		t.Error("Diagram LayoutPart has no LayoutDefinition root element")
	}
	if diagramResult.StylePart.StyleDefinition() == nil {
		t.Error("Diagram StylePart has no StyleDefinition root element")
	}
	if diagramResult.ColorsPart.ColorsDefinition() == nil {
		t.Error("Diagram ColorsPart has no ColorsDefinition root element")
	}
}

func TestSlidePart_AddDiagram_Hierarchy(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	// Test creating diagram with hierarchy template
	diagramResult, err := slidePart.AddDiagram(diagram.TemplateTypeHierarchy)
	if err != nil {
		t.Fatalf("AddDiagram(hierarchy) error = %v", err)
	}

	// Verify all parts are created
	if diagramResult.DataPart == nil {
		t.Error("AddDiagram() with hierarchy returned Diagram with nil DataPart")
	}
	if diagramResult.LayoutPart == nil {
		t.Error("AddDiagram() with hierarchy returned Diagram with nil LayoutPart")
	}
	if diagramResult.StylePart == nil {
		t.Error("AddDiagram() with hierarchy returned Diagram with nil StylePart")
	}
	if diagramResult.ColorsPart == nil {
		t.Error("AddDiagram() with hierarchy returned Diagram with nil ColorsPart")
	}
}

func TestDiagram_TemplateMethods(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	diag, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("AddDiagramPart() error = %v", err)
	}

	// Test template helper methods
	listLayout := diag.GetLayoutTemplate(diagram.TemplateTypeList)
	if listLayout == "" {
		t.Error("GetLayoutTemplate should return non-empty string for list template")
	}

	hierarchyLayout := diag.GetLayoutTemplate(diagram.TemplateTypeHierarchy)
	if hierarchyLayout == "" {
		t.Error("GetLayoutTemplate should return non-empty string for hierarchy template")
	}

	listStyle := diag.GetStyleTemplate(diagram.TemplateTypeList)
	if listStyle == "" {
		t.Error("GetStyleTemplate should return non-empty string for list template")
	}

	listColors := diag.GetColorTemplate(diagram.TemplateTypeList)
	if listColors == "" {
		t.Error("GetColorTemplate should return non-empty string for list template")
	}

	// Verify templates are different for different types
	if listLayout == hierarchyLayout {
		t.Error("List and hierarchy layout templates should be different")
	}
}

func TestDiagram_TemplatesAreValidXML(t *testing.T) {
	templateTypes := []diagram.TemplateType{
		diagram.TemplateTypeList,
		diagram.TemplateTypeHierarchy,
	}

	for _, templateType := range templateTypes {
		t.Run(string(templateType), func(t *testing.T) {
			layoutXML := diagram.GetLayoutTemplate(templateType)
			if !strings.HasPrefix(layoutXML, "<?xml") {
				t.Errorf("Layout template should start with XML declaration, got: %s...", layoutXML[:50])
			}

			styleXML := diagram.GetStyleTemplate(templateType)
			if !strings.HasPrefix(styleXML, "<?xml") {
				t.Errorf("Style template should start with XML declaration, got: %s...", styleXML[:50])
			}

			colorXML := diagram.GetColorTemplate(templateType)
			if !strings.HasPrefix(colorXML, "<?xml") {
				t.Errorf("Color template should start with XML declaration, got: %s...", colorXML[:50])
			}
		})
	}
}
