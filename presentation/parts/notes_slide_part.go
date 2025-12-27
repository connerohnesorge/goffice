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
