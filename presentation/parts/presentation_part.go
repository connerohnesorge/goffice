//nolint:revive // line-length-limit: OOXML content types are long strings
package parts

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/elements"
)

// PresentationPart represents the main presentation part (ppt/presentation.xml).
// This is the primary part containing the presentation definition and slide references.
type PresentationPart struct {
	*openxml.OpenXmlPartData

	// contentType is the specific content type for this presentation type
	contentType string
}

// NewPresentationPart creates a new presentation part.
func NewPresentationPart(
	uri, contentType string,
	container openxml.OpenXmlPartContainer,
) (*PresentationPart, error) {
	// Create the underlying packaging part
	pkg := container.Package()
	if pkg == nil {
		return nil, ErrNilPackage
	}

	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)

	partData.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewPresentation()
		},
	)

	pp := &PresentationPart{
		OpenXmlPartData: partData,
		contentType:     contentType,
	}

	// Create package-level relationship
	_, err = pkg.CreateRelationship(
		uri,
		RelationshipTypeOfficeDocument,
		"",
	)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

// NewPresentationPartFromData wraps an existing OpenXmlPartData as a PresentationPart.
func NewPresentationPartFromData(
	data *openxml.OpenXmlPartData,
	contentType string,
) *PresentationPart {
	data.SetRootFactory(
		func() openxml.PartRootElement {
			return elements.NewPresentation()
		},
	)

	return &PresentationPart{
		OpenXmlPartData: data,
		contentType:     contentType,
	}
}

// FixedContentType returns the content type for this part.
func (pp *PresentationPart) FixedContentType() string {
	return pp.contentType
}

// InitializeContent sets up minimal presentation content for a new presentation.
func (pp *PresentationPart) InitializeContent() {
	pres := elements.NewPresentation()
	pp.SetRootElement(pres)
}

// Presentation returns the root Presentation element.
func (pp *PresentationPart) Presentation() *elements.Presentation {
	root := pp.RootElement()
	if root == nil {
		return nil
	}
	if pres, ok := root.(*elements.Presentation); ok {
		return pres
	}

	return nil
}

// Counters for generating unique filenames
var (
	slideCounter       uint64
	slideMasterCounter uint64
	notesMasterCounter uint64
	handoutCounter     uint64
)

// Counters for generating slide IDs (PowerPoint starts from 256)
var (
	slideIdCounter       uint32 = 255
	slideMasterIdCounter uint32 = 2147483647 // PowerPoint convention for master IDs
)

// AddSlidePart adds a new slide part to this presentation.
func (pp *PresentationPart) AddSlidePart() (*SlidePart, error) {
	num := atomic.AddUint64(&slideCounter, 1)
	uri := fmt.Sprintf(
		"/ppt/slides/slide%d.xml",
		num,
	)

	slidePart, err := newSlidePart(pp, uri)
	if err != nil {
		return nil, err
	}

	// Register the slide in presentation.xml
	slideID := atomic.AddUint32(
		&slideIdCounter,
		1,
	)
	pres := pp.Presentation()
	if pres != nil {
		pres.AddSlideId(
			slideID,
			slidePart.RelationshipID(),
		)
	}

	return slidePart, nil
}

// SlideParts returns all slide parts.
func (pp *PresentationPart) SlideParts() []*SlidePart {
	var slides []*SlidePart
	for part := range pp.Parts() {
		if sp, ok := part.(*SlidePart); ok {
			slides = append(slides, sp)
		}
	}

	return slides
}

// AddSlideMasterPart adds a new slide master part to this presentation.
func (pp *PresentationPart) AddSlideMasterPart() (*SlideMasterPart, error) {
	num := atomic.AddUint64(
		&slideMasterCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/slideMasters/slideMaster%d.xml",
		num,
	)

	masterPart, err := newSlideMasterPart(pp, uri)
	if err != nil {
		return nil, err
	}

	// Register the slide master in presentation.xml
	masterID := atomic.AddUint32(
		&slideMasterIdCounter,
		1,
	)
	pres := pp.Presentation()
	if pres != nil {
		pres.AddSlideMasterId(
			masterID,
			masterPart.RelationshipID(),
		)
	}

	return masterPart, nil
}

// SlideMasterParts returns all slide master parts.
func (pp *PresentationPart) SlideMasterParts() []*SlideMasterPart {
	var masters []*SlideMasterPart
	for part := range pp.Parts() {
		if smp, ok := part.(*SlideMasterPart); ok {
			masters = append(masters, smp)
		}
	}

	return masters
}

// AddNotesMasterPart adds a notes master part to this presentation.
func (pp *PresentationPart) AddNotesMasterPart() (*NotesMasterPart, error) {
	num := atomic.AddUint64(
		&notesMasterCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/notesMasters/notesMaster%d.xml",
		num,
	)

	return newNotesMasterPart(pp, uri)
}

// NotesMasterPart returns the notes master part if present.
func (pp *PresentationPart) NotesMasterPart() *NotesMasterPart {
	for part := range pp.Parts() {
		if nmp, ok := part.(*NotesMasterPart); ok {
			return nmp
		}
	}

	return nil
}

// AddHandoutMasterPart adds a handout master part to this presentation.
func (pp *PresentationPart) AddHandoutMasterPart() (*HandoutMasterPart, error) {
	num := atomic.AddUint64(&handoutCounter, 1)
	uri := fmt.Sprintf(
		"/ppt/handoutMasters/handoutMaster%d.xml",
		num,
	)

	return newHandoutMasterPart(pp, uri)
}

// HandoutMasterPart returns the handout master part if present.
func (pp *PresentationPart) HandoutMasterPart() *HandoutMasterPart {
	for part := range pp.Parts() {
		if hmp, ok := part.(*HandoutMasterPart); ok {
			return hmp
		}
	}

	return nil
}

// AddThemePart adds a theme part to this presentation.
func (pp *PresentationPart) AddThemePart() (*ThemePart, error) {
	return newThemePart(pp)
}

// ThemePart returns the first theme part if present.
func (pp *PresentationPart) ThemePart() *ThemePart {
	for part := range pp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			return tp
		}
	}

	return nil
}

// ThemeParts returns all theme parts.
func (pp *PresentationPart) ThemeParts() []*ThemePart {
	var themes []*ThemePart
	for part := range pp.Parts() {
		if tp, ok := part.(*ThemePart); ok {
			themes = append(themes, tp)
		}
	}

	return themes
}

// AddCommentAuthorsPart adds a comment authors part to this presentation.
func (pp *PresentationPart) AddCommentAuthorsPart() (*CommentAuthorsPart, error) {
	return newCommentAuthorsPart(pp)
}

// CommentAuthorsPart returns the comment authors part if present.
func (pp *PresentationPart) CommentAuthorsPart() *CommentAuthorsPart {
	for part := range pp.Parts() {
		if commentAuthorsPart, ok := part.(*CommentAuthorsPart); ok {
			return commentAuthorsPart
		}
	}

	return nil
}

// GetStream returns a reader for the part content.
func (pp *PresentationPart) GetStream() io.Reader {
	return pp.OpenXmlPartData.GetStream()
}

// Ensure PresentationPart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*PresentationPart)(
	nil,
)

// PresentationPartFactory creates a PresentationPart from a URI and container.
func PresentationPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// Get the packaging part from the container directly
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		packPart.ContentType(),
		packPart,
		container,
	)

	return NewPresentationPartFromData(
		partData,
		packPart.ContentType(),
	)
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func (pp *PresentationPart) addChildPart(
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	return addChildPart(
		pp,
		uri,
		contentType,
		relType,
	)
}

// Register the PresentationPart type.
func init() {
	// Register for all PowerPoint presentation content types
	contentTypes := []string{
		ContentTypePresentation,
		ContentTypePresentationTemplate,
		ContentTypePresentationMacroEnabled,
		ContentTypeMacroTemplate,
		ContentTypeSlideshow,
		ContentTypeSlideshowMacroEnabled,
		ContentTypeAddIn,
	}

	for _, ct := range contentTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:        ct,
				RelationshipType:   RelationshipTypeOfficeDocument,
				Factory:            PresentationPartFactory,
				DefaultURI:         "/ppt/presentation.xml",
				IsFixedContentType: false, // Content type varies by presentation type
			},
		)
	}
}
