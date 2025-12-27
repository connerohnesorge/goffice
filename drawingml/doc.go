// Package drawingml provides shared DrawingML types for shapes, images,
// charts, diagrams, and effects used across Office Open XML documents.
//
// DrawingML (Drawing Markup Language) is the common drawing specification used
// in Word, Excel, and PowerPoint documents. This package provides the core
// types and utilities shared by all Office applications, including:
//
//   - Coordinate and positioning types using EMU (English Metric Units)
//   - Transform2D for geometric transformations
//   - Common namespace definitions for DrawingML schemas
//   - Unit conversion utilities
//
// The drawingml package is designed to be used by application-specific packages
// such as wordprocessing, spreadsheet, and presentation, providing a consistent
// foundation for graphical content across all document types.
//
// # English Metric Units (EMU)
//
// DrawingML uses English Metric Units (EMU) as its primary unit of measurement.
// One EMU equals 1/914400 of an inch or 1/360000 of a centimeter.
// This high precision allows for exact representation of measurements in both
// imperial and metric systems without rounding errors.
//
// Use the unit conversion functions in this package to convert between EMUs
// and more familiar units like inches, centimeters, points, and pixels.
//
// # Namespaces
//
// DrawingML defines several XML namespaces for different aspects of drawing:
//
//   - Main DrawingML (a:) - Core drawing elements
//   - Picture (pic:) - Image content
//   - Chart (c:) - Charts and graphs
//   - Diagram (dgm:) - SmartArt and diagrams
//   - WordprocessingDrawing (wp:) - Drawing anchoring in Word
//   - SpreadsheetDrawing (xdr:) - Drawing positioning in Excel
//   - ChartDrawing (cdr:) - Drawing in charts
package drawingml
