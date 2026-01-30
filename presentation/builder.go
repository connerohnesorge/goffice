// Package presentation provides a fluent API for creating and modifying PowerPoint presentations.
package presentation

import (
	"errors"
	"fmt"
	"io"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// PresentationBuilder provides a fluent API for building presentations.
type PresentationBuilder struct {
	doc    *Document
	errors []error
}

// NewPresentationBuilder creates a new PresentationBuilder.
func NewPresentationBuilder() *PresentationBuilder {
	doc, err := NewWriter(io.Discard, DocTypePresentation)
	pb := &PresentationBuilder{
		doc:    doc,
		errors: make([]error, 0),
	}
	if err != nil {
		pb.errors = append(pb.errors, err)
	}

	return pb
}

// AddSlide adds a new slide to the presentation and returns a SlideBuilder.
func (pb *PresentationBuilder) AddSlide() *SlideBuilder {
	if pb.doc == nil {
		pb.errors = append(pb.errors, errors.New("document is nil"))

		return &SlideBuilder{pb: pb}
	}

	slidePart, err := pb.doc.AddSlide()
	if err != nil {
		pb.errors = append(pb.errors, err)

		return &SlideBuilder{pb: pb}
	}

	return &SlideBuilder{
		pb:        pb,
		slidePart: slidePart,
	}
}

// Build returns the constructed Document or the first error encountered.
func (pb *PresentationBuilder) Build() (*Document, error) {
	if len(pb.errors) > 0 {
		return nil, pb.errors[0]
	}

	return pb.doc, nil
}

// SlideBuilder provides a fluent API for building slides.
type SlideBuilder struct {
	pb        *PresentationBuilder
	slidePart *parts.SlidePart
}

// AddShape adds a new shape to the slide and returns a ShapeBuilder.
func (sb *SlideBuilder) AddShape() *ShapeBuilder {
	if sb.slidePart == nil {
		return &ShapeBuilder{sb: sb}
	}

	slide := sb.slidePart.Slide()
	shape := slide.AddShape()

	return &ShapeBuilder{
		sb:    sb,
		shape: shape,
	}
}

// AddTitle adds a title shape to the slide with the specified text.
func (sb *SlideBuilder) AddTitle(text string) *SlideBuilder {
	if sb.slidePart == nil {
		return sb
	}

	shape := sb.AddShape()
	shape.SetType(drawingml.ShapeTypeRectangle)
	shape.AddParagraph(text)

	return sb
}

// Presentation returns the parent PresentationBuilder.
func (sb *SlideBuilder) Presentation() *PresentationBuilder {
	return sb.pb
}

// SlidePart returns the underlying slide part.
func (sb *SlideBuilder) SlidePart() *parts.SlidePart {
	return sb.slidePart
}

// ShapeBuilder provides a fluent API for building shapes.
type ShapeBuilder struct {
	sb    *SlideBuilder
	shape *elements.Shape
}

// SetType sets the shape type.
func (sh *ShapeBuilder) SetType(shapeType drawingml.ShapeTypeValue) *ShapeBuilder {
	if sh.shape != nil {
		spPr := sh.shape.GetOrCreateShapeProperties()
		spPr.SetPresetGeometry(string(shapeType))
	}

	return sh
}

// AddParagraph adds a new paragraph to the shape and returns a ParagraphBuilder.
func (sh *ShapeBuilder) AddParagraph(text string) *ParagraphBuilder {
	if sh.shape == nil {
		return &ParagraphBuilder{sh: sh}
	}

	tb := sh.shape.GetOrCreateTextBody()
	p := tb.AddParagraph(text)

	return &ParagraphBuilder{
		sh: sh,
		p:  p,
	}
}

// Slide returns the parent SlideBuilder.
func (sh *ShapeBuilder) Slide() *SlideBuilder {
	return sh.sb
}

// Shape returns the underlying shape element.
func (sh *ShapeBuilder) Shape() *elements.Shape {
	return sh.shape
}

// ParagraphBuilder provides a fluent API for building paragraphs.
type ParagraphBuilder struct {
	sh *ShapeBuilder
	p  *drawingml.TextParagraph
}

// AddRun adds a new text run to the paragraph and returns a RunBuilder.
func (pb *ParagraphBuilder) AddRun(text string) *RunBuilder {
	if pb.p == nil {
		return &RunBuilder{pb: pb}
	}

	r := pb.p.AddRun(text)

	return &RunBuilder{
		pb: pb,
		r:  r,
	}
}

// Shape returns the parent ShapeBuilder.
func (pb *ParagraphBuilder) Shape() *ShapeBuilder {
	return pb.sh
}

// RunBuilder provides a fluent API for building text runs.
type RunBuilder struct {
	pb *ParagraphBuilder
	r  *drawingml.TextRun
}

// Bold sets the bold property of the text run.
func (rb *RunBuilder) Bold(bold bool) *RunBuilder {
	if rb.r != nil {
		rb.r.SetBold(bold)
	}

	return rb
}

// Size sets the font size of the text run.
func (rb *RunBuilder) Size(size int) *RunBuilder {
	if rb.r != nil {
		rb.r.SetFontSize(size)
	}

	return rb
}

// Paragraph returns the parent ParagraphBuilder.
func (rb *RunBuilder) Paragraph() *ParagraphBuilder {
	return rb.pb
}
