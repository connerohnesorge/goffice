package presentation

import (
	"testing"
)

func TestPresentationBuilder(t *testing.T) {
	pb := NewPresentationBuilder()
	doc, err := pb.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if doc == nil {
		t.Fatal("Build returned nil document")
	}
}

func TestSlideBuilder(t *testing.T) {
	pb := NewPresentationBuilder()
	pb.AddSlide().
		AddTitle("My Title")

	doc, err := pb.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	if doc.SlideCount() != 1 {
		t.Errorf("Expected 1 slide, got %d", doc.SlideCount())
	}
}
