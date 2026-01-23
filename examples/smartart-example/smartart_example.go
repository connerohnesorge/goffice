// Package main provides an example of creating PowerPoint presentations with SmartArt diagrams.
//nolint:revive // Example code may have longer functions and repeated string literals for clarity
package main

import (
	"fmt"
	"log"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/presentation"
)

// createBasicListDiagram creates a slide with a basic list SmartArt diagram.
// This demonstrates creating a simple bullet-point style diagram.
func createBasicListDiagram(doc *presentation.Document) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("failed to add slide: %w", err)
	}

	// Create a SmartArt diagram with list layout
	dgm, err := slidePart.AddDiagram(diagram.TemplateTypeList)
	if err != nil {
		return fmt.Errorf("failed to create diagram: %w", err)
	}

	// Use the DataModelBuilder to construct the diagram content
	builder := diagram.NewDataModelBuilder()

	// Add points (nodes) to the diagram
	// Note: The current implementation supports the diagram infrastructure.
	// Points and connections are added to the data model for structure.
	builder.AddPoint("root")
	builder.AddPoint("item1")

	// Add connection between points
	builder.AddParentOfConnection("root", "item1")

	// Build the data model and set it on the data part
	dataModel := builder.Build()
	dgm.DataPart.SetRootElement(dataModel)

	fmt.Println("Created basic list diagram")

	return nil
}

// createHierarchyDiagram creates a slide with an organization chart style diagram.
// This demonstrates creating a hierarchical tree structure.
func createHierarchyDiagram(doc *presentation.Document) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("failed to add slide: %w", err)
	}

	// Create a SmartArt diagram with hierarchy layout (org chart)
	dgm, err := slidePart.AddDiagram(diagram.TemplateTypeHierarchy)
	if err != nil {
		return fmt.Errorf("failed to create diagram: %w", err)
	}

	// Build an organization chart structure
	builder := diagram.NewDataModelBuilder()

	// Add organizational nodes
	builder.AddPoint("CEO", diagram.PointValuesNode)
	builder.AddPoint("VP Engineering", diagram.PointValuesNode)

	// Add assistant point (e.g., executive assistant)
	builder.AddAssistantPoint("Executive Assistant")

	// Create hierarchical connections
	builder.AddParentOfConnection("CEO", "VP Engineering")
	builder.AddParentOfConnection("CEO", "Executive Assistant")

	// Build and set the data model
	dataModel := builder.Build()
	dgm.DataPart.SetRootElement(dataModel)

	fmt.Println("Created hierarchy (org chart) diagram")

	return nil
}

// createProcessDiagram creates a slide with a process flow diagram.
// This demonstrates a process workflow structure.
func createProcessDiagram(doc *presentation.Document) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("failed to add slide: %w", err)
	}

	// Create a diagram (using list layout as base)
	dgm, err := slidePart.AddDiagram(diagram.TemplateTypeList)
	if err != nil {
		return fmt.Errorf("failed to create diagram: %w", err)
	}

	// Build a process workflow structure
	builder := diagram.NewDataModelBuilder()

	// Process stages
	builder.AddPoint("Start")
	builder.AddPoint("Planning")
	builder.AddPoint("Execution")

	// Connect stages in sequence
	builder.AddParentOfConnection("Start", "Planning")
	builder.AddParentOfConnection("Planning", "Execution")

	// Build and set the data model
	dataModel := builder.Build()
	dgm.DataPart.SetRootElement(dataModel)

	fmt.Println("Created process flow diagram")

	return nil
}

// createComparisonDiagram creates a slide with a comparison diagram.
// This demonstrates a side-by-side comparison structure.
func createComparisonDiagram(doc *presentation.Document) error {
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("failed to add slide: %w", err)
	}

	// Create a diagram
	dgm, err := slidePart.AddDiagram(diagram.TemplateTypeList)
	if err != nil {
		return fmt.Errorf("failed to create diagram: %w", err)
	}

	// Build a comparison structure (Option A vs Option B)
	builder := diagram.NewDataModelBuilder()

	// Main comparison topic
	builder.AddPoint("Comparison Topic")

	// Option A branch
	builder.AddPoint("Option A: Traditional")
	builder.AddPoint("A: Feature 1")

	// Option B branch
	builder.AddPoint("Option B: Modern")
	builder.AddPoint("B: Feature 1")

	// Create connections showing the comparison structure
	builder.AddParentOfConnection("Comparison Topic", "Option A: Traditional")
	builder.AddParentOfConnection("Comparison Topic", "Option B: Modern")
	builder.AddParentOfConnection("Option A: Traditional", "A: Feature 1")
	builder.AddParentOfConnection("Option B: Modern", "B: Feature 1")

	// Build and set the data model
	dataModel := builder.Build()
	dgm.DataPart.SetRootElement(dataModel)

	fmt.Println("Created comparison diagram")

	return nil
}

func main() {
	fmt.Println("Creating SmartArt example presentation...")

	// Create a new presentation document
	doc, err := presentation.New(
		"smartart_example.pptx",
		presentation.DocTypePresentation,
	)
	if err != nil {
		log.Fatalf("Failed to create presentation: %v", err)
	}
	defer func() {
		if err := doc.Close(); err != nil {
			log.Printf("Failed to close document: %v", err)
		}
	}()

	// Create slides with different SmartArt diagrams
	if err := createBasicListDiagram(doc); err != nil {
		log.Printf("Failed to create basic list diagram: %v", err)

		return
	}

	if err := createHierarchyDiagram(doc); err != nil {
		log.Printf("Failed to create hierarchy diagram: %v", err)

		return
	}

	if err := createProcessDiagram(doc); err != nil {
		log.Printf("Failed to create process diagram: %v", err)

		return
	}

	if err := createComparisonDiagram(doc); err != nil {
		log.Printf("Failed to create comparison diagram: %v", err)

		return
	}

	// Save the presentation
	if err := doc.Save(); err != nil {
		log.Printf("Failed to save presentation: %v", err)

		return
	}

	fmt.Println("✅ SmartArt example presentation created successfully!")
	fmt.Println("📄 Output file: smartart_example.pptx")
	fmt.Println("")
	fmt.Println("Diagrams created:")
	fmt.Println("  1. Basic List - Simple bullet-point style diagram")
	fmt.Println("  2. Hierarchy - Organization chart structure")
	fmt.Println("  3. Process Flow - Linear workflow diagram")
	fmt.Println("  4. Comparison - Side-by-side comparison structure")
	fmt.Println("")
	fmt.Println("Open smartart_example.pptx in PowerPoint to view the diagrams.")
	fmt.Println("")
	fmt.Println("Note: This example demonstrates the SmartArt creation API.")
	fmt.Println("The diagram infrastructure (parts, templates, builder) is in place.")
	fmt.Println("Full point/connection serialization requires additional element support.")
}
