package presentation

import (
	"fmt"
	"io"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

type PresentationBuilder struct {
	doc    *Document
	errors []error
}

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

func (pb *PresentationBuilder) AddSlide() *SlideBuilder {
	if pb.doc == nil {
		pb.errors = append(pb.errors, fmt.Errorf("document is nil"))

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

func (pb *PresentationBuilder) Build() (*Document, error) {
	if len(pb.errors) > 0 {
		return nil, pb.errors[0]
	}

	return pb.doc, nil
}

type SlideBuilder struct {
	pb        *PresentationBuilder
	slidePart *parts.SlidePart
}

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

func (sb *SlideBuilder) AddTitle(text string) *SlideBuilder {
	if sb.slidePart == nil {
		return sb
	}

	shape := sb.AddShape()
	shape.SetType(drawingml.ShapeTypeRectangle)
	shape.AddParagraph(text)

	return sb
}

func (sb *SlideBuilder) Presentation() *PresentationBuilder {
	return sb.pb
}

func (sb *SlideBuilder) SlidePart() *parts.SlidePart {
	return sb.slidePart
}

type ShapeBuilder struct {
	sb    *SlideBuilder
	shape *elements.Shape
}

func (shb *ShapeBuilder) SetType(shapeType drawingml.ShapeTypeValue) *ShapeBuilder {
	if shb.shape != nil {
		spPr := shb.shape.GetOrCreateShapeProperties()
		spPr.SetPresetGeometry(string(shapeType))
	}

	return shb
}

func (shb *ShapeBuilder) AddParagraph(text string) *ParagraphBuilder {
	if shb.shape == nil {
		return &ParagraphBuilder{shb: shb}
	}

	tb := shb.shape.GetOrCreateTextBody()
	p := tb.AddParagraph(text)

	return &ParagraphBuilder{
		shb: shb,
		p:   p,
	}
}

func (shb *ShapeBuilder) Slide() *SlideBuilder {
	return shb.sb
}

func (shb *ShapeBuilder) Shape() *elements.Shape {
	return shb.shape
}

type ParagraphBuilder struct {
	shb *ShapeBuilder
	p   *drawingml.TextParagraph
}

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

func (pb *ParagraphBuilder) Shape() *ShapeBuilder {
	return pb.shb
}

type RunBuilder struct {
	pb *ParagraphBuilder
	r  *drawingml.TextRun
}

func (rb *RunBuilder) Bold(bold bool) *RunBuilder {
	if rb.r != nil {
		rb.r.SetBold(bold)
	}

	return rb
}

func (rb *RunBuilder) Size(size int) *RunBuilder {
	if rb.r != nil {
		rb.r.SetFontSize(size)
	}

	return rb
}

func (rb *RunBuilder) Paragraph() *ParagraphBuilder {
	return rb.pb
}
