//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// Counter for generating unique notes slide filenames.
var notesSlideUniqueCounter uint64

// NotesSlidePart represents a notes slide part (ppt/notesSlides/notesSlide1.xml, etc.).
type NotesSlidePart struct {
	*openxml.OpenXmlPartData
}

// newNotesSlidePart creates a new notes slide part.
func newNotesSlidePart(
	slidePart *SlidePart,
) (*NotesSlidePart, error) {
	num := atomic.AddUint64(
		&notesSlideUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/notesSlides/notesSlide%d.xml",
		num,
	)

	packPart, relID, err := slidePart.addChildPart(
		uri,
		ContentTypeNotesSlide,
		RelationshipTypeNotesSlide,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesSlide,
		packPart,
		slidePart,
	)
	partData.SetRelationshipID(relID)

	nsp := &NotesSlidePart{
		OpenXmlPartData: partData,
	}

	// Initialize with minimal notes slide content
	nsp.initializeContent()

	// Add to slide part's child parts
	if err := slidePart.AddPart(nsp, relID); err != nil {
		return nil, err
	}

	return nsp, nil
}

// initializeContent sets up minimal notes slide content.
func (nsp *NotesSlidePart) initializeContent() {
	ns := elements.NewNotesSlide()
	nsp.SetRootElement(ns)
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*NotesSlidePart) FixedContentType() string {
	return ContentTypeNotesSlide
}

// NotesSlide returns the root NotesSlide element.
func (nsp *NotesSlidePart) NotesSlide() *elements.NotesSlide {
	root := nsp.RootElement()
	if root == nil {
		return nil
	}
	if ns, ok := root.(*elements.NotesSlide); ok {
		return ns
	}

	return nil
}

// GetOrCreateCommonSlideData returns or creates the common slide data (p:cSld).
func (nsp *NotesSlidePart) GetOrCreateCommonSlideData() *elements.CommonSlideData {
	ns := nsp.NotesSlide()
	if ns == nil {
		return nil
	}
	
	// Check if cSld already exists
	for child := range ns.Children() {
		if csd, ok := child.(*elements.CommonSlideData); ok {
			return csd
		}
	}
	
	csd := elements.NewCommonSlideData()
	ns.AppendChild(csd)
	return csd
}

// SetNotes sets the plain text notes for this slide.
func (nsp *NotesSlidePart) SetNotes(text string) {
	csd := nsp.GetOrCreateCommonSlideData()
	st := csd.GetOrCreateShapeTree()
	
	// Notes usually have a body placeholder for text
	// For now, we'll just add a shape with the text
	shape := st.AddShape()
	tb := shape.GetOrCreateTextBody()
	if tb == nil {
		panic("GetOrCreateTextBody returned nil")
	}
	if tb.CompositeElementBase == nil {
		panic("TextBody.CompositeElementBase is nil")
	}
	tb.AddParagraph(text)
}

// GetOrCreateTextBody returns or creates the main text body for notes.
func (nsp *NotesSlidePart) GetOrCreateTextBody() *elements.TextBody {
	csd := nsp.GetOrCreateCommonSlideData()
	st := csd.GetOrCreateShapeTree()
	
	// Try to find existing body placeholder
	for _, shape := range st.Shapes() {
		if nvsp := shape.NonVisualShapeProperties(); nvsp != nil {
			// In notes, usually there is a placeholder with type "body"
			// But for simplicity, we'll just use the first shape with a text body or create one
			if tb := shape.TextBody(); tb != nil {
				return tb
			}
		}
	}
	
	// Create new shape with text body
	shape := st.AddShape()
	return shape.GetOrCreateTextBody()
}

// GetStream returns a reader for the part content.
func (nsp *NotesSlidePart) GetStream() io.Reader {
	return nsp.OpenXmlPartData.GetStream()
}

// Ensure NotesSlidePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*NotesSlidePart)(nil)

// NotesSlidePartFactory creates a NotesSlidePart from a URI and container.
func NotesSlidePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		ContentTypeNotesSlide,
		packPart,
		container,
	)
	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewNotesSlide()
		},
	)

	return &NotesSlidePart{
		OpenXmlPartData: partData,
	}
}

// Register the NotesSlidePart type.
func init() {
	openxml.RegisterPartType(
		&openxml.PartTypeInfo{
			ContentType:        ContentTypeNotesSlide,
			RelationshipType:   RelationshipTypeNotesSlide,
			Factory:            NotesSlidePartFactory,
			DefaultURI:         "/ppt/notesSlides/notesSlide1.xml",
			IsFixedContentType: true,
		},
	)
}
