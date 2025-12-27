// Package elements provides PresentationML element types.
//
// This package contains the core elements used in
// PresentationML documents (.pptx):
//
// # Presentation Structure
//
// The presentation hierarchy is:
//
//	Presentation
//	  -> SlideMasterIdList
//	  -> SlideIdList
//	  -> SlideSize
//	  -> NotesSize
//
// # Slide Elements
//
//   - Slide: Individual slide content
//   - CommonSlideData: Shared slide content container
//   - ShapeTree: Container for shapes and visual elements
//
// # Shape Elements
//
//   - Shape: Text shapes and placeholders
//   - Picture: Image elements
//   - GroupShape: Grouped shapes
//   - GraphicFrame: Charts, tables, and other objects
//
// # Text Elements
//
// Text elements leverage the shared DrawingML text types:
//
//   - TextBody: Container for text content
//   - TextParagraph: Individual paragraphs
//   - TextRun: Formatted text runs
//
// # Usage Example
//
//	pres := elements.NewPresentation()
//	slide := elements.NewSlide()
//	shape := elements.NewShape()
//	shape.TextBody().AddParagraph("Hello, World!")
//
// # Namespaces
//
// Standard PresentationML namespaces:
//
//	elements.NamespacePresentationML  // Main namespace
//	elements.PrefixP                  // "p:" prefix
//
//nolint:revive // max-public-structs: elements package defines many public types
package elements
