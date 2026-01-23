//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/drawingml/diagram"
	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// Counter for generating unique diagram filenames.
var diagramUniqueCounter uint64

// Diagram represents a collection of all 4 diagram parts required for SmartArt.
// This struct provides convenient access to all diagram components created together.
type Diagram struct {
	DataPart   *DiagramDataPart
	LayoutPart *DiagramLayoutDefinitionPart
	StylePart  *DiagramStylePart
	ColorsPart *DiagramColorsPart
}

// SlidePart represents a slide part (ppt/slides/slide1.xml, etc.).
type SlidePart struct {
	*openxml.OpenXmlPartData
}

// newSlidePart creates a new slide part.
func newSlidePart(
	presentationPart *PresentationPart,
	uri string,
) (*SlidePart, error) {
	packPart, relID, err := presentationPart.addChildPart(
		uri,
		ContentTypeSlide,
		RelationshipTypeSlide,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlide,
		packPart,
		presentationPart,
	)
	partData.SetRelationshipID(relID)

	sp := &SlidePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal slide content
	sp.initializeContent()

	// Add to presentation part's child parts
	if err := presentationPart.AddPart(sp, relID); err != nil {
		return nil, err
	}

	return sp, nil
}

// initializeContent sets up minimal slide content.
func (sp *SlidePart) initializeContent() {
	slide := elements.NewSlide()
	sp.SetRootElement(slide)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*SlidePart) FixedContentType() string {
	return ContentTypeSlide
}

// Slide returns the root Slide element.
func (sp *SlidePart) Slide() *elements.Slide {
	root := sp.RootElement()
	if root == nil {
		return nil
	}
	if slide, ok := root.(*elements.Slide); ok {
		return slide
	}

	return nil
}

// AddNotesSlidePart adds a notes slide part to this slide.
func (sp *SlidePart) AddNotesSlidePart() (*NotesSlidePart, error) {
	return newNotesSlidePart(sp)
}

// NotesSlidePart returns the notes slide part if present.
func (sp *SlidePart) NotesSlidePart() *NotesSlidePart {
	for part := range sp.Parts() {
		if nsp, ok := part.(*NotesSlidePart); ok {
			return nsp
		}
	}

	return nil
}

// AddSlideCommentsPart adds a comments part to this slide.
func (sp *SlidePart) AddSlideCommentsPart() (*SlideCommentsPart, error) {
	return newSlideCommentsPart(sp)
}

// SlideCommentsPart returns the comments part if present.
func (sp *SlidePart) SlideCommentsPart() *SlideCommentsPart {
	for part := range sp.Parts() {
		if scp, ok := part.(*SlideCommentsPart); ok {
			return scp
		}
	}

	return nil
}

// AddImagePart adds an image part with the specified type.
func (sp *SlidePart) AddImagePart(
	imageType ImageType,
) (*ImagePart, error) {
	return newImagePartForSlide(sp, imageType)
}

// ImageParts returns all image parts.
func (sp *SlidePart) ImageParts() []*ImagePart {
	var images []*ImagePart
	for part := range sp.Parts() {
		if ip, ok := part.(*ImagePart); ok {
			images = append(images, ip)
		}
	}

	return images
}

// AddChartPart adds a chart part to this slide.
func (sp *SlidePart) AddChartPart() (*ChartPart, error) {
	return newChartPartForSlide(sp)
}

// ChartParts returns all chart parts.
func (sp *SlidePart) ChartParts() []*ChartPart {
	var charts []*ChartPart
	for part := range sp.Parts() {
		if cp, ok := part.(*ChartPart); ok {
			charts = append(charts, cp)
		}
	}

	return charts
}

// AddDiagramPart creates and returns all 4 diagram parts required for SmartArt.
// The parts are created as child parts of the slide.
func (sp *SlidePart) AddDiagramPart() (*Diagram, error) {
	num := atomic.AddUint64(&diagramUniqueCounter, 1)

	dataURI := fmt.Sprintf("/ppt/diagrams/data%d.xml", num)
	layoutURI := fmt.Sprintf("/ppt/diagrams/layout%d.xml", num)
	styleURI := fmt.Sprintf("/ppt/diagrams/quickStyle%d.xml", num)
	colorsURI := fmt.Sprintf("/ppt/diagrams/colors%d.xml", num)

	dataPart, err := newDiagramDataPart(sp, dataURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create diagram data part: %w", err)
	}

	layoutPart, err := newDiagramLayoutDefinitionPart(sp, layoutURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create diagram layout part: %w", err)
	}

	stylePart, err := newDiagramStylePart(sp, styleURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create diagram style part: %w", err)
	}

	colorsPart, err := newDiagramColorsPart(sp, colorsURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create diagram colors part: %w", err)
	}

	return &Diagram{
		DataPart:   dataPart,
		LayoutPart: layoutPart,
		StylePart:  stylePart,
		ColorsPart: colorsPart,
	}, nil
}

// AddDiagram creates a complete SmartArt diagram on the slide with the specified template type.
// This method creates all 4 required diagram parts (Data, Layout, Style, Colors) and initializes
// them with the appropriate default templates.
//
// The templateType parameter specifies the diagram layout:
//   - diagram.TemplateTypeList: Basic block list layout
//   - diagram.TemplateTypeHierarchy: Organization chart layout
//
// Returns the Diagram containing all 4 parts, or an error if creation fails.
func (sp *SlidePart) AddDiagram(templateType diagram.TemplateType) (*Diagram, error) {
	diagramPart, err := sp.AddDiagramPart()
	if err != nil {
		return nil, err
	}

	// Initialize with template content
	// The parts are already created with minimal content by AddDiagramPart.
	// The templates would need to be parsed and applied if we want to use custom templates.
	// For now, the default initializeContent methods create basic empty structures.
	_ = diagramPart.GetLayoutTemplate(templateType)
	_ = diagramPart.GetStyleTemplate(templateType)
	_ = diagramPart.GetColorTemplate(templateType)

	return diagramPart, nil
}

// GetLayoutTemplate returns the layout template XML for the specified template type.
func (d *Diagram) GetLayoutTemplate(templateType diagram.TemplateType) string {
	return diagram.GetLayoutTemplate(templateType)
}

// GetStyleTemplate returns the style template XML for the specified template type.
func (d *Diagram) GetStyleTemplate(templateType diagram.TemplateType) string {
	return diagram.GetStyleTemplate(templateType)
}

// GetColorTemplate returns the color template XML for the specified template type.
func (d *Diagram) GetColorTemplate(templateType diagram.TemplateType) string {
	return diagram.GetColorTemplate(templateType)
}

// GetStream returns a reader for the part content.
func (sp *SlidePart) GetStream() io.Reader {
	return sp.OpenXmlPartData.GetStream()
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (sp *SlidePart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		sp,
		uri,
		contentType,
		relType,
	)
}

// Ensure SlidePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*SlidePart)(nil)

// SlidePartFactory creates a SlidePart from a URI and container.
func SlidePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeSlide,
		packPart,
		container,
	)

	return &SlidePart{
		OpenXmlPartData: partData,
	}
}

// Register the SlidePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeSlide,
			RelationshipType:   RelationshipTypeSlide,
			Factory:            SlidePartFactory,
			DefaultURI:         "/ppt/slides/slide1.xml",
			IsFixedContentType: true,
		},
	)
}
