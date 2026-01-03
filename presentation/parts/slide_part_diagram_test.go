package parts

import (
	"testing"
)

func TestSlidePart_AddDiagramPart(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	diagram, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("AddDiagramPart() error = %v", err)
	}

	if diagram.DataPart == nil {
		t.Error("AddDiagramPart() returned Diagram with nil DataPart")
	}
	if diagram.LayoutPart == nil {
		t.Error("AddDiagramPart() returned Diagram with nil LayoutPart")
	}
	if diagram.StylePart == nil {
		t.Error("AddDiagramPart() returned Diagram with nil StylePart")
	}
	if diagram.ColorsPart == nil {
		t.Error("AddDiagramPart() returned Diagram with nil ColorsPart")
	}

	if diagram.DataPart.DataModel() == nil {
		t.Error("Diagram DataPart has no DataModel root element")
	}
	if diagram.LayoutPart.LayoutDefinition() == nil {
		t.Error("Diagram LayoutPart has no LayoutDefinition root element")
	}
	if diagram.StylePart.StyleDefinition() == nil {
		t.Error("Diagram StylePart has no StyleDefinition root element")
	}
	if diagram.ColorsPart.ColorsDefinition() == nil {
		t.Error("Diagram ColorsPart has no ColorsDefinition root element")
	}

	foundData := false
	foundLayout := false
	foundStyle := false
	foundColors := false

	for part := range slidePart.Parts() {
		if _, ok := part.(*DiagramDataPart); ok {
			foundData = true
		}
		if _, ok := part.(*DiagramLayoutDefinitionPart); ok {
			foundLayout = true
		}
		if _, ok := part.(*DiagramStylePart); ok {
			foundStyle = true
		}
		if _, ok := part.(*DiagramColorsPart); ok {
			foundColors = true
		}
	}

	if !foundData {
		t.Error("DiagramDataPart not found in slide's child parts")
	}
	if !foundLayout {
		t.Error("DiagramLayoutDefinitionPart not found in slide's child parts")
	}
	if !foundStyle {
		t.Error("DiagramStylePart not found in slide's child parts")
	}
	if !foundColors {
		t.Error("DiagramColorsPart not found in slide's child parts")
	}
}

func TestSlidePart_AddDiagramPart_UniqueURIs(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	diagram1, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("First AddDiagramPart() error = %v", err)
	}

	diagram2, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("Second AddDiagramPart() error = %v", err)
	}

	if diagram1.DataPart.URI() == diagram2.DataPart.URI() {
		t.Error("Multiple diagrams have same DataPart URI")
	}
	if diagram1.LayoutPart.URI() == diagram2.LayoutPart.URI() {
		t.Error("Multiple diagrams have same LayoutPart URI")
	}
	if diagram1.StylePart.URI() == diagram2.StylePart.URI() {
		t.Error("Multiple diagrams have same StylePart URI")
	}
	if diagram1.ColorsPart.URI() == diagram2.ColorsPart.URI() {
		t.Error("Multiple diagrams have same ColorsPart URI")
	}

	expectedPrefix := "/ppt/diagrams/"
	if diagram1.DataPart.URI()[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("DataPart URI %q does not start with %q", diagram1.DataPart.URI(), expectedPrefix)
	}
	if diagram1.LayoutPart.URI()[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("LayoutPart URI %q does not start with %q", diagram1.LayoutPart.URI(), expectedPrefix)
	}
	if diagram1.StylePart.URI()[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("StylePart URI %q does not start with %q", diagram1.StylePart.URI(), expectedPrefix)
	}
	if diagram1.ColorsPart.URI()[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("ColorsPart URI %q does not start with %q", diagram1.ColorsPart.URI(), expectedPrefix)
	}
}

func TestDiagram_AllPartsRequired(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	diagram, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("AddDiagramPart() error = %v", err)
	}

	if diagram.DataPart == nil || diagram.LayoutPart == nil ||
		diagram.StylePart == nil || diagram.ColorsPart == nil {
		t.Error("Diagram must contain all 4 required parts")
	}
}

func TestDiagram_PartContentTypes(t *testing.T) {
	presPart, cleanup := createTestPresentationPart(t)
	defer cleanup()

	slidePart, err := presPart.AddSlidePart()
	if err != nil {
		t.Fatalf("AddSlidePart() error = %v", err)
	}

	diagram, err := slidePart.AddDiagramPart()
	if err != nil {
		t.Fatalf("AddDiagramPart() error = %v", err)
	}

	if diagram.DataPart.FixedContentType() != ContentTypeDiagramData {
		t.Errorf("DataPart has wrong content type: %v", diagram.DataPart.FixedContentType())
	}
	if diagram.LayoutPart.FixedContentType() != ContentTypeDiagramLayoutDef {
		t.Errorf("LayoutPart has wrong content type: %v", diagram.LayoutPart.FixedContentType())
	}
	if diagram.StylePart.FixedContentType() != ContentTypeDiagramStyle {
		t.Errorf("StylePart has wrong content type: %v", diagram.StylePart.FixedContentType())
	}
	if diagram.ColorsPart.FixedContentType() != ContentTypeDiagramColors {
		t.Errorf("ColorsPart has wrong content type: %v", diagram.ColorsPart.FixedContentType())
	}
}
