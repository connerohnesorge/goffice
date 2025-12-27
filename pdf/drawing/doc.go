// Package drawing provides shared rendering for DrawingML elements in PDF
// output.
//
// DrawingML is the Office Open XML markup language for graphics, used across
// Word, Excel, and PowerPoint documents. This package provides common rendering
// functions for:
//   - Shapes (rectangles, ellipses, custom geometry)
//   - Images (JPEG, PNG, with transformations)
//   - Fills (solid, gradient, pattern, picture)
//   - Lines and outlines (styles, widths, dash patterns)
//   - Effects (shadows, reflections, glow)
//   - Text in shapes (text boxes, word art)
//
// # Path Drawing
//
// The PathBuilder type provides a fluent API for constructing PDF paths:
//
//	pb := drawing.NewPathBuilder()
//	pb.MoveTo(100, 100).
//	    LineTo(200, 100).
//	    LineTo(200, 200).
//	    LineTo(100, 200).
//	    ClosePath()
//	content := pb.String()
//
// # Coordinate Transformations
//
// DrawingML uses EMU (English Metric Units) for all measurements. This package
// provides conversion utilities:
//
//	points := core.EMUToPoints(914400)  // Returns 72.0 (1 inch)
//
// # Transform Matrix
//
// The Transform type provides 2D affine transformation operations:
//
//	t := drawing.NewTransform().
//	    Translate(50, 50).
//	    Rotate(45).
//	    Scale(2, 2)
//	content := t.ToCMOperator()
//
// # Shape Rendering
//
// Shapes are rendered using their preset geometry or custom path definitions:
//
//	shape := doc.GetShape()
//	drawing.RenderShape(page, shape, transform)
//
// # Fill Rendering
//
// Fills support solid colors, gradients, patterns, and pictures:
//
//	drawing.RenderFill(page, shape.Fill(), bounds)
//
// This package is used by the word, spreadsheet, and presentation packages
// to render graphical content consistently across document types.
package drawing
