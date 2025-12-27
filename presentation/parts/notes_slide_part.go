//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
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
	content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:notes xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr/>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="2" name="Slide Image Placeholder 1"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1" noRot="1" noChangeAspect="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="sldImg"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr/>
      </p:sp>
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="3" name="Notes Placeholder 2"/>
          <p:cNvSpPr>
            <a:spLocks noGrp="1"/>
          </p:cNvSpPr>
          <p:nvPr>
            <p:ph type="body" idx="1"/>
          </p:nvPr>
        </p:nvSpPr>
        <p:spPr/>
        <p:txBody>
          <a:bodyPr/>
          <a:lstStyle/>
          <a:p>
            <a:endParaRPr lang="en-US"/>
          </a:p>
        </p:txBody>
      </p:sp>
    </p:spTree>
  </p:cSld>
  <p:clrMapOvr>
    <a:masterClrMapping/>
  </p:clrMapOvr>
</p:notes>`
	nsp.SetData([]byte(content))
}

// FixedContentType returns the content type for this part.
//
//nolint:revive // unused-receiver: interface implementation returns constant
func (*NotesSlidePart) FixedContentType() string {
	return ContentTypeNotesSlide
}

// NotesSlide returns the root NotesSlide element.
// TODO: Return a proper NotesSlide element type when elements are implemented.
func (nsp *NotesSlidePart) NotesSlide() openxml.PartRootElement {
	return nsp.RootElement()
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
