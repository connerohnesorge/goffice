//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"github.com/connerohnesorge/goffice/packaging"
)

// Content types for PresentationML parts.
const (
	// Presentation content types
	ContentTypePresentation             = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"
	ContentTypePresentationTemplate     = "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml"
	ContentTypePresentationMacroEnabled = "application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml"
	ContentTypeMacroTemplate            = "application/vnd.ms-powerpoint.template.macroEnabled.main+xml"
	ContentTypeSlideshow                = "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml"
	ContentTypeSlideshowMacroEnabled    = "application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml"
	ContentTypeAddIn                    = "application/vnd.ms-powerpoint.addin.macroEnabled.main+xml"

	// Slide content types
	ContentTypeSlide       = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	ContentTypeSlideLayout = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
	ContentTypeSlideMaster = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"

	// Notes content types
	ContentTypeNotesSlide  = "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml"
	ContentTypeNotesMaster = "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml"

	// Handout content types
	ContentTypeHandoutMaster = "application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml"

	// Theme content types
	ContentTypeTheme = "application/vnd.openxmlformats-officedocument.theme+xml"

	// Comment content types
	ContentTypeCommentAuthors = "application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml"
	ContentTypeComments       = "application/vnd.openxmlformats-officedocument.presentationml.comments+xml"

	// Drawing content types
	ContentTypeDrawing = "application/vnd.openxmlformats-officedocument.drawing+xml"
	ContentTypeChart   = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"

	// VBA content types
	ContentTypeVbaProject = "application/vnd.ms-office.vbaProject"

	// Table style content type
	ContentTypeTableStyles = "application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml"

	// View properties content type
	ContentTypeViewProps = "application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml"

	// Presentation properties content type
	ContentTypePresProps = "application/vnd.openxmlformats-officedocument.presentationml.presProps+xml"

	// Tags content type
	ContentTypeTags = "application/vnd.openxmlformats-officedocument.presentationml.tags+xml"
)

// Relationship types for PresentationML parts.
const (
	// Core relationships
	RelationshipTypeOfficeDocument = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	RelationshipTypeSlide          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	RelationshipTypeSlideLayout    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
	RelationshipTypeSlideMaster    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"

	// Notes relationships
	RelationshipTypeNotesSlide  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesSlide"
	RelationshipTypeNotesMaster = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"

	// Handout relationships
	RelationshipTypeHandoutMaster = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"

	// Theme relationships
	RelationshipTypeTheme = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"

	// Comment relationships
	RelationshipTypeCommentAuthors = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/commentAuthors"
	RelationshipTypeComments       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"

	// Drawing relationships
	RelationshipTypeDrawing = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	RelationshipTypeChart   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	RelationshipTypeImage   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"

	// VBA relationships
	RelationshipTypeVbaProject = "http://schemas.microsoft.com/office/2006/relationships/vbaProject"

	// Table style relationships
	RelationshipTypeTableStyles = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"

	// View properties relationships
	RelationshipTypeViewProps = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"

	// Presentation properties relationships
	RelationshipTypePresProps = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"

	// Tags relationships
	RelationshipTypeTags = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tags"

	// Hyperlink
	RelationshipTypeHyperlink = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
)

// ErrNilPackage is returned when the package is nil.
var ErrNilPackage = partError("package is nil")

type partError string

func (e partError) Error() string {
	return string(e)
}

// addChildPart is a helper to add a child part with the appropriate relationship.
func addChildPart(
	parent partWithPackage,
	uri, contentType, relType string,
) (*packaging.Part, string, error) {
	pkg := parent.Package()
	if pkg == nil {
		return nil, "", ErrNilPackage
	}

	// Create the underlying packaging part
	packPart, err := pkg.CreatePart(
		uri,
		contentType,
	)
	if err != nil {
		return nil, "", err
	}

	// Create relationship from parent to this part
	rel, err := pkg.CreatePartRelationship(
		parent.URI(),
		uri,
		relType,
		"",
	)
	if err != nil {
		return nil, "", err
	}

	return packPart, rel.ID(), nil
}

// partWithPackage is an interface for parts that can access the package.
type partWithPackage interface {
	Package() *packaging.Package
	URI() string
}
