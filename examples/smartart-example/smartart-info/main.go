// Package main demonstrates reading SmartArt diagram structure from Office documents.
//
// This example shows how to:
// 1. Open a Word document containing SmartArt diagrams
// 2. Access diagram parts (data, layout, style, colors)
// 3. Print diagram structure (nodes and connections)
// 4. Display diagram metadata
//
// Note: This is a Phase 1 example showing roundtrip support.
// Full diagram creation and modification will be added in Phase 2.
package main

import (
	"fmt"
	"os"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
)

//nolint:unused // Used in printDiagramInfo helper function
const unknownValue = "unknown"

func main() {
	fmt.Println(
		"SmartArt Diagram Information Example",
	)
	fmt.Println(
		"====================================",
	)

	// In Phase 1, we demonstrate the diagram element types
	// Full document integration will be added as part of the main APIs

	demonstrateDiagramStructure()
	demonstrateDiagramTypes()
	demonstratePDFPlaceholder()

	fmt.Println(
		"\nNote: This example demonstrates Phase 1 capabilities:",
	)
	fmt.Println(
		"  - Diagram element types and structure",
	)
	fmt.Println(
		"  - Roundtrip support (preserving diagrams in documents)",
	)
	fmt.Println(
		"  - Basic PDF placeholder rendering",
	)
	fmt.Println("\nPhase 2 will add:")
	fmt.Println(
		"  - Full diagram creation and modification",
	)
	fmt.Println(
		"  - Layout engine and automatic positioning",
	)
	fmt.Println("  - High-fidelity PDF rendering")
}

// demonstrateDiagramStructure shows the structure of a SmartArt data model
func demonstrateDiagramStructure() {
	fmt.Println(
		"1. SmartArt Data Model Structure",
	)
	fmt.Println(
		"---------------------------------",
	)

	// Create a simple data model to demonstrate the structure
	dataModel := diagram.NewDataModelRoot()

	// In a real application, this would be loaded from an Office document
	// For demonstration, we show what the structure looks like

	fmt.Println("Data Model Components:")
	fmt.Println(
		"  - Points (Nodes): Individual items in the diagram",
	)
	fmt.Println(
		"  - Connections: Relationships between points",
	)
	fmt.Println(
		"  - Background: Optional background formatting",
	)
	fmt.Println(
		"  - ExtensionList: Optional extensions for custom data",
	)

	if dataModel.PointList != nil {
		pointCount := 0
		if dataModel.PointList.Point != nil {
			pointCount = 1
		}
		fmt.Printf(
			"  Point List initialized: %d points",
			pointCount,
		)
	}
	if dataModel.ConnectionList != nil {
		connCount := 0
		if dataModel.ConnectionList.Connection != nil {
			connCount = 1
		}
		fmt.Printf(
			"  Connection List initialized: %d connections",
			connCount,
		)
	}

	fmt.Println()
}

// demonstrateDiagramTypes shows the different diagram part types
func demonstrateDiagramTypes() {
	fmt.Println("2. SmartArt Part Types")
	fmt.Println("----------------------")

	fmt.Println(
		"A SmartArt diagram consists of four separate parts:",
	)

	fmt.Println("a) DiagramDataPart (data*.xml)")
	fmt.Println(
		"   - Stores the logical diagram structure",
	)
	fmt.Println(
		"   - Contains points and connections",
	)
	fmt.Println(
		"   - Holds text content for each node",
	)

	fmt.Println(
		"b) DiagramLayoutDefinitionPart (layout*.xml)",
	)
	fmt.Println(
		"   - Defines how points are arranged",
	)
	fmt.Println("   - Contains layout algorithms")
	fmt.Println(
		"   - Specifies positioning rules and constraints",
	)

	fmt.Println(
		"c) DiagramStylePart (quickStyle*.xml)",
	)
	fmt.Println("   - Defines visual appearance")
	fmt.Println(
		"   - Specifies shape styles, fills, and effects",
	)
	fmt.Println("   - Controls text formatting")

	fmt.Println(
		"d) DiagramColorsPart (colors*.xml)",
	)
	fmt.Println("   - Defines the color scheme")
	fmt.Println(
		"   - Maps semantic colors to actual values",
	)
	fmt.Println(
		"   - Supports theme color integration",
	)
}

// demonstratePDFPlaceholder shows the PDF rendering capabilities
func demonstratePDFPlaceholder() {
	fmt.Println("3. PDF Rendering (Phase 1)")
	fmt.Println("--------------------------")

	fmt.Println(
		"Current PDF rendering capabilities:",
	)

	fmt.Println(
		"Phase 1 provides placeholder rendering:",
	)
	fmt.Println(
		"  - Draws a light gray rectangle",
	)
	fmt.Println("  - Shows dashed border")
	fmt.Println(
		"  - Displays 'SmartArt Diagram' label",
	)
	fmt.Println(
		"  - Preserves diagram space in layout",
	)

	fmt.Println("Example PDF rendering code:")
	fmt.Println(`
	// Create PDF page and rendering context
	page := doc.AddPage(core.PageSizeA4)
	ctx := core.NewRenderingContext(612, 792).WithPage(page)

	// Create diagram renderer
	renderer := drawing.NewDiagramRenderer(ctx)

	// Render placeholder
	bounds := drawing.RenderBounds{
		X:      100,
		Y:      200,
		Width:  300,
		Height: 150,
	}
	renderer.RenderDiagramBounds(bounds)
	`)

	fmt.Println(
		"\nPhase 2 will implement full rendering:",
	)
	fmt.Println("  - Execute layout algorithms")
	fmt.Println(
		"  - Render actual shapes and connections",
	)
	fmt.Println("  - Apply styles and colors")
	fmt.Println("  - Render text within shapes")
	fmt.Println(
		"  - Support effects (shadows, reflections, etc.)",
	)
}

// printDiagramInfo prints information about a diagram (helper function)
//
//nolint:unused,revive // Will be used when full document integration is available
func printDiagramInfo(
	dataModel *diagram.DataModelRoot,
) {
	if dataModel == nil {
		fmt.Println("No data model provided")

		return
	}

	fmt.Println("Diagram Information:")
	fmt.Println("-------------------")

	// Print points
	if dataModel.PointList != nil &&
		dataModel.PointList.Point != nil {
		point := dataModel.PointList.Point
		modelID := unknownValue
		if point.ModelId != nil {
			modelID = point.ModelId.Value()
		}

		pointType := "none"
		if point.Type != nil {
			pointType = string(point.Type.Value())
		}

		fmt.Printf("Points: 1\n")
		fmt.Printf(
			"  1. ID: %s, Type: %s",
			modelID,
			pointType,
		)

		// Print text if available
		if point.TextBody != nil {
			fmt.Println("     Has text content")
		}
	} else {
		fmt.Println("Points: 0")
	}

	// Print connections
	if dataModel.ConnectionList != nil &&
		dataModel.ConnectionList.Connection != nil {
		conn := dataModel.ConnectionList.Connection
		sourceID := unknownValue
		destID := unknownValue

		if conn.SourceId != nil {
			sourceID = conn.SourceId.Value()
		}
		if conn.DestinationId != nil {
			destID = conn.DestinationId.Value()
		}

		connType := "none"
		if conn.Type != nil {
			connType = string(conn.Type.Value())
		}

		fmt.Printf("\nConnections: 1\n")
		fmt.Printf(
			"  1. %s -> %s (Type: %s)",
			sourceID,
			destID,
			connType,
		)
	} else {
		fmt.Println("\nConnections: 0")
	}
}

// Example of how to use this with a real document (commented out until integration is complete)
//
//nolint:unused // Kept as example template for future integration
func exampleWithRealDocument() {
	// This will work once diagram parts are integrated with the main document APIs
	//
	// doc, err := wordprocessing.Open("document_with_smartart.docx", false)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer doc.Close()
	//
	// // Get diagram parts from document
	// // (Exact API to be determined during integration)
	// diagramParts := doc.GetDiagramParts()
	//
	// for i, dataPart := range diagramParts {
	//     fmt.Printf("\nDiagram %d:\n", i+1)
	//     dataModel := dataPart.DataModelRoot()
	//     printDiagramInfo(dataModel)
	//
	//     // Access other parts
	//     layoutPart := dataPart.LayoutDefinitionPart()
	//     stylePart := dataPart.StylePart()
	//     colorsPart := dataPart.ColorsPart()
	//
	//     fmt.Printf("  Has Layout: %v\n", layoutPart != nil)
	//     fmt.Printf("  Has Style: %v\n", stylePart != nil)
	//     fmt.Printf("  Has Colors: %v\n", colorsPart != nil)
	// }

	fmt.Fprintln(
		os.Stderr,
		"Full document integration coming in Phase 2",
	)
}
