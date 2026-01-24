//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package diagram provides types and functions for working with SmartArt diagrams
// in Office Open XML documents.
//
// SmartArt diagrams use the dgm: namespace (http://schemas.openxmlformats.org/drawingml/2006/diagram)
// and represent structured visual diagrams with data models, layouts, styles, and colors.
//
// # Diagram Structure
//
// SmartArt diagrams consist of four main components:
//
//   - Data Model (dgm:dataModel): Defines the logical structure of the diagram with points
//     and connections representing nodes and relationships.
//   - Layout Definition (dgm:layoutDef): Specifies how data points are arranged visually,
//     including algorithms for automatic layout.
//   - Style Definition (dgm:styleDef): Defines the visual styling of diagram elements
//     including shapes, text, and effects.
//   - Colors Definition (dgm:colorsDef): Specifies the color scheme applied to the diagram.
//
// # Key Element Types
//
// The package includes generated element classes for all diagram-related elements:
//
//   - DataModelRoot: Root element containing diagram data points and connections
//   - Point: Individual nodes in the diagram data structure
//   - Connection: Relationships between points in the diagram
//   - LayoutDefinition: Root element for layout algorithms
//   - StyleDefinition: Root element for style definitions
//   - ColorsDefinition: Root element for color scheme definitions
//
// # Usage
//
// This package is typically used in conjunction with the higher-level presentation,
// spreadsheet, or wordprocessing packages that contain SmartArt diagrams:
//
//	import "github.com/connerohnesorge/goffice/drawingml/diagram"
//
//	// Work with diagram data model
//	dataModel := diagram.NewDataModelRoot()
//	point := diagram.NewPoint()
//	dataModel.PointList.Points = append(dataModel.PointList.Points, point)
//
// # Standards Compliance
//
// This package implements the DrawingML Diagram specification from:
//   - ECMA-376 Part 1 (Office Open XML File Formats)
//   - ISO/IEC 29500-1
//
// It includes support for Office 2007-2019+ diagram extensions.
//
// # Code Generation
//
// Most types in this package are generated from the official Open-XML-SDK schemas
// using the cmd/gen-go-diagram generator. Do not manually edit generated files
// (elements.go, enums.go).
package diagram
