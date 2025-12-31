# DrawingML Diagram (SmartArt) Package

This package provides support for reading and roundtripping SmartArt diagrams in Office Open XML documents.

## Overview

SmartArt diagrams use the `dgm:` namespace (`http://schemas.openxmlformats.org/drawingml/2006/diagram`) and represent structured visual diagrams with data models, layouts, styles, and colors.

## Diagram Structure

SmartArt diagrams consist of four main components, each stored in separate parts within the Office document:

### 1. Data Model (`DiagramDataPart`)
- Contains the logical structure of the diagram
- Defines points (nodes) and connections (relationships)
- Stores text content and properties for each node
- File pattern: `diagrams/data*.xml`

### 2. Layout Definition (`DiagramLayoutDefinitionPart`)
- Specifies how data points are arranged visually
- Contains layout algorithms for automatic positioning
- Defines constraints and rules for shape arrangement
- File pattern: `diagrams/layout*.xml`

### 3. Style Definition (`DiagramStylePart`)
- Defines visual styling of diagram elements
- Specifies shape types, fills, lines, and effects
- Controls text formatting within shapes
- File pattern: `diagrams/quickStyle*.xml`

### 4. Colors Definition (`DiagramColorsPart`)
- Specifies the color scheme applied to the diagram
- Maps semantic color roles to actual RGB/theme colors
- File pattern: `diagrams/colors*.xml`

## Current Capabilities (Phase 1)

**Phase 1** provides foundational support for SmartArt diagrams:

- **Roundtrip Support**: Read SmartArt diagrams from Office documents and write them back without data loss
- **Part Management**: Access and manipulate diagram parts (data, layout, style, colors)
- **Element Structure**: Strongly-typed element classes for all diagram XML structures
- **PDF Placeholder**: Minimal placeholder rendering in PDF output (shows "SmartArt Diagram" box)

### What's Supported

```go
// Open a document containing SmartArt
doc, _ := wordprocessing.Open("document_with_smartart.docx", false)

// Access diagram parts from a graphic frame
// (Implementation depends on integration with main document API)

// Read diagram data
dataRoot := diagramDataPart.DataModelRoot()
points := dataRoot.PointList.Points
for _, point := range points {
    fmt.Println("Node:", point.ModelID)
}

// Roundtrip: Save the document
doc.Save()  // SmartArt preserved unchanged
```

### What's Not Yet Supported (Coming in Phase 2)

- **Layout Engine**: Automatic diagram layout computation
- **Full PDF Rendering**: High-fidelity rendering of diagrams to PDF
- **Diagram Creation**: Creating new SmartArt diagrams from scratch
- **Diagram Modification**: Changing diagram structure, adding/removing nodes
- **Layout Algorithm Execution**: Running the layout algorithms to position shapes

## API Examples

### Reading Diagram Structure

```go
import (
    "github.com/connerohnesorge/goffice/drawingml/diagram"
    "github.com/connerohnesorge/goffice/wordprocessing"
)

// Open document
doc, err := wordprocessing.Open("document.docx", false)
if err != nil {
    panic(err)
}
defer doc.Close()

// Access diagram parts (exact API depends on document integration)
// Diagram parts are accessed through graphic frames that reference them

// Example: Read data model
dataModel := diagram.NewDataModelRoot()
// Load from XML...

// Iterate through points
for _, point := range dataModel.PointList.Points {
    if point.ModelID != nil {
        fmt.Printf("Point ID: %s\n", *point.ModelID)
    }
    if point.TextBody != nil {
        // Access text content
    }
}

// Iterate through connections
for _, conn := range dataModel.ConnectionList.Connections {
    if conn.SourceID != nil && conn.DestinationID != nil {
        fmt.Printf("Connection: %s -> %s\n", *conn.SourceID, *conn.DestinationID)
    }
}
```

### PDF Rendering (Placeholder)

```go
import (
    "github.com/connerohnesorge/goffice-pdf/core"
    "github.com/connerohnesorge/goffice-pdf/drawing"
)

// Create PDF page and rendering context
page := doc.AddPage(core.PageSizeA4)
ctx := core.NewRenderingContext(612, 792).WithPage(page)

// Create diagram renderer
renderer := drawing.NewDiagramRenderer(ctx)

// Render placeholder (Phase 1)
bounds := drawing.RenderBounds{
    X:      100,
    Y:      200,
    Width:  300,
    Height: 150,
}
renderer.RenderDiagramBounds(bounds)
// This draws a gray rectangle with "SmartArt Diagram" label
```

## File Patterns

SmartArt diagrams are stored across multiple files in the Office package:

```
word/
  diagrams/
    data1.xml              # Data model (points & connections)
    layout1.xml            # Layout definition
    quickStyle1.xml        # Style definition
    colors1.xml            # Color scheme
  diagrams/_rels/
    data1.xml.rels         # Relationships to layout, style, colors
```

## Element Types

Key element types provided by this package:

### Data Model Elements
- `DataModelRoot` - Root container for diagram data
- `PointList` - Collection of diagram points (nodes)
- `Point` - Individual node with properties and text
- `ConnectionList` - Collection of connections
- `Connection` - Relationship between points

### Layout Elements
- `LayoutDefinition` - Root container for layout
- `LayoutNode` - Node in the layout tree
- `Algorithm` - Layout algorithm specification
- `Constraint` - Layout constraint rule

### Style Elements
- `StyleDefinition` - Root container for styles
- `StyleLabel` - Named style configuration
- `ShapeStyle` - Shape appearance properties

### Color Elements
- `ColorsDefinition` - Root container for colors
- `ColorTransform` - Color transformation rules
- `FillColorList` - Fill color specifications

## Standards Compliance

This package implements the DrawingML Diagram specification from:
- **ECMA-376 Part 1** - Office Open XML File Formats
- **ISO/IEC 29500-1** - Information technology — Document description and processing languages

Supports diagram features from:
- Office 2007 (baseline)
- Office 2010 extensions
- Office 2013 extensions
- Office 2016+ extensions

## Code Generation

Most types in this package are **schema-driven generated code**:

- `elements.go` - Generated element classes (~132KB)
- `enums.go` - Generated enumeration types (~108KB)

**Do not manually edit generated files**. To regenerate:

```bash
# When generator is implemented
go run ./cmd/gen-go-diagram
```

## Future Plans (Phase 2)

Phase 2 will add full diagram support:

1. **Layout Engine**
   - Parse and execute layout algorithms
   - Compute shape positions and sizes
   - Handle constraints and rules
   - Support all standard layout types (hierarchy, cycle, matrix, pyramid, etc.)

2. **Full PDF Rendering**
   - Render computed shapes with styles
   - Apply color schemes
   - Render text within shapes
   - Draw connection lines and arrows
   - Support effects (shadows, reflections, etc.)

3. **Diagram Creation & Modification**
   - Create new SmartArt diagrams programmatically
   - Add/remove points and connections
   - Change layout types
   - Apply different styles and colors

4. **Advanced Features**
   - Custom layout definitions
   - Animation support
   - Data binding
   - Alternative text and accessibility

## Limitations (Phase 1)

Current limitations in Phase 1:

- **No layout computation**: Shapes are not positioned
- **Placeholder PDF rendering only**: Just shows a labeled box
- **Read-only**: Cannot create or modify diagrams
- **No style application**: Styles are preserved but not applied during rendering
- **No validation**: XML structure is preserved but not validated

These will be addressed in Phase 2.

## Package Documentation

For detailed API documentation:

```bash
go doc github.com/connerohnesorge/goffice/drawingml/diagram
```

## Related Packages

- `drawingml` - DrawingML core types (shapes, fills, effects)
- `wordprocessing` - Word document API
- `presentation` - PowerPoint document API
- `spreadsheet` - Excel document API
- `pdf/drawing` - PDF rendering for DrawingML elements

## Contributing

When working with diagram support:

1. **Preserve roundtrip fidelity**: Ensure all XML data is preserved
2. **Use generated types**: Work with the generated element classes
3. **Follow OpenXML standards**: Implement according to ECMA-376 spec
4. **Test with real documents**: Use actual Office files for testing
5. **Plan for Phase 2**: Keep full rendering implementation in mind

## References

- [ECMA-376 Part 1 - DrawingML Diagrams](http://www.ecma-international.org/publications/standards/Ecma-376.htm)
- [Microsoft Open XML SDK Documentation](https://docs.microsoft.com/en-us/office/open-xml/structure-of-a-presentationml-document)
- Office Open XML Format Specification
