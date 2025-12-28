// Package presentation provides PresentationML support for PowerPoint documents.
// This file implements the document-level API for creating, reading,
// and modifying .pptx files.
//
//nolint:revive // file-length-limit and line-length-limit: comprehensive document API with OOXML constants
package presentation

import (
	"io"
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// DocType represents the type of a PowerPoint document.
type DocType int

const (
	// DocTypePresentation represents a standard PowerPoint presentation (.pptx).
	DocTypePresentation DocType = iota

	// DocTypeTemplate represents a PowerPoint template (.potx).
	DocTypeTemplate

	// DocTypeSlideshow represents a PowerPoint slideshow (.ppsx).
	DocTypeSlideshow

	// DocTypeMacroEnabledPresentation represents a macro-enabled PowerPoint presentation (.pptm).
	DocTypeMacroEnabledPresentation

	// DocTypeMacroEnabledTemplate represents a macro-enabled PowerPoint template (.potm).
	DocTypeMacroEnabledTemplate

	// DocTypeMacroEnabledSlideshow represents a macro-enabled PowerPoint slideshow (.ppsm).
	DocTypeMacroEnabledSlideshow

	// DocTypeAddIn represents a PowerPoint add-in (.ppam).
	DocTypeAddIn
)

// String returns the string representation of the document type.
func (dt DocType) String() string {
	switch dt {
	case DocTypePresentation:
		return "Presentation"
	case DocTypeTemplate:
		return "Template"
	case DocTypeSlideshow:
		return "Slideshow"
	case DocTypeMacroEnabledPresentation:
		return "MacroEnabledPresentation"
	case DocTypeMacroEnabledTemplate:
		return "MacroEnabledTemplate"
	case DocTypeMacroEnabledSlideshow:
		return "MacroEnabledSlideshow"
	case DocTypeAddIn:
		return "AddIn"
	default:
		return "Presentation"
	}
}

// Extension returns the file extension for this document type.
func (dt DocType) Extension() string {
	switch dt {
	case DocTypePresentation:
		return ".pptx"
	case DocTypeTemplate:
		return ".potx"
	case DocTypeSlideshow:
		return ".ppsx"
	case DocTypeMacroEnabledPresentation:
		return ".pptm"
	case DocTypeMacroEnabledTemplate:
		return ".potm"
	case DocTypeMacroEnabledSlideshow:
		return ".ppsm"
	case DocTypeAddIn:
		return ".ppam"
	default:
		return ".pptx"
	}
}

// ContentType returns the MIME content type for the main presentation part.
func (dt DocType) ContentType() string {
	switch dt {
	case DocTypePresentation:
		return ContentTypePresentationMLPresentation
	case DocTypeTemplate:
		return ContentTypePresentationMLTemplate
	case DocTypeSlideshow:
		return ContentTypePresentationMLSlideshow
	case DocTypeMacroEnabledPresentation:
		return ContentTypePresentationMLMacroEnabled
	case DocTypeMacroEnabledTemplate:
		return ContentTypePresentationMLMacroTemplate
	case DocTypeMacroEnabledSlideshow:
		return ContentTypePresentationMLMacroSlideshow
	case DocTypeAddIn:
		return ContentTypePresentationMLAddIn
	default:
		return ContentTypePresentationMLPresentation
	}
}

// Content types for PowerPoint documents.
const (
	// ContentTypePresentationMLPresentation is the content type for a standard
	// PowerPoint presentation main part.
	ContentTypePresentationMLPresentation = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"

	// ContentTypePresentationMLTemplate is the content type for a PowerPoint
	// template main part.
	ContentTypePresentationMLTemplate = "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml"

	// ContentTypePresentationMLSlideshow is the content type for a PowerPoint
	// slideshow main part.
	ContentTypePresentationMLSlideshow = "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml"

	// ContentTypePresentationMLMacroEnabled is the content type for a
	// macro-enabled PowerPoint presentation main part.
	ContentTypePresentationMLMacroEnabled = "application/vnd.ms-powerpoint.presentation.macroEnabled.main+xml"

	// ContentTypePresentationMLMacroTemplate is the content type for a
	// macro-enabled PowerPoint template main part.
	ContentTypePresentationMLMacroTemplate = "application/vnd.ms-powerpoint.template.macroEnabled.main+xml"

	// ContentTypePresentationMLMacroSlideshow is the content type for a
	// macro-enabled PowerPoint slideshow main part.
	ContentTypePresentationMLMacroSlideshow = "application/vnd.ms-powerpoint.slideshow.macroEnabled.main+xml"

	// ContentTypePresentationMLAddIn is the content type for a PowerPoint add-in
	// main part.
	ContentTypePresentationMLAddIn = "application/vnd.ms-powerpoint.addin.macroEnabled.main+xml"

	// ContentTypeSlide is the content type for a slide.
	ContentTypeSlide = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"

	// ContentTypeSlideLayout is the content type for a slide layout.
	ContentTypeSlideLayout = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"

	// ContentTypeSlideMaster is the content type for a slide master.
	ContentTypeSlideMaster = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"

	// ContentTypeNotesMaster is the content type for a notes master.
	ContentTypeNotesMaster = "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml"

	// ContentTypeNotesSlide is the content type for a notes slide.
	ContentTypeNotesSlide = "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml"

	// ContentTypeHandoutMaster is the content type for a handout master.
	ContentTypeHandoutMaster = "application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml"

	// ContentTypePresProps is the content type for presentation properties.
	ContentTypePresProps = "application/vnd.openxmlformats-officedocument.presentationml.presProps+xml"

	// ContentTypeViewProps is the content type for view properties.
	ContentTypeViewProps = "application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml"

	// ContentTypeTableStyles is the content type for table styles.
	ContentTypeTableStyles = "application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml"

	// ContentTypeCommentAuthors is the content type for comment authors.
	ContentTypeCommentAuthors = "application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml"

	// ContentTypeComments is the content type for comments.
	ContentTypePresentationComments = "application/vnd.openxmlformats-officedocument.presentationml.comments+xml"

	// ContentTypeTags is the content type for tags.
	ContentTypeTags = "application/vnd.openxmlformats-officedocument.presentationml.tags+xml"
)

// Relationship types for PowerPoint documents.
const (
	// RelationshipTypeSlide is the relationship type for slides.
	RelationshipTypeSlide = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"

	// RelationshipTypeSlideLayout is the relationship type for slide layouts.
	RelationshipTypeSlideLayout = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"

	// RelationshipTypeSlideMaster is the relationship type for slide masters.
	RelationshipTypeSlideMaster = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"

	// RelationshipTypeNotesMaster is the relationship type for notes master.
	RelationshipTypeNotesMaster = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"

	// RelationshipTypeNotesSlide is the relationship type for notes slides.
	RelationshipTypeNotesSlide = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesSlide"

	// RelationshipTypeHandoutMaster is the relationship type for handout master.
	RelationshipTypeHandoutMaster = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"

	// RelationshipTypePresProps is the relationship type for presentation properties.
	RelationshipTypePresProps = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"

	// RelationshipTypeViewProps is the relationship type for view properties.
	RelationshipTypeViewProps = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"

	// RelationshipTypeTableStyles is the relationship type for table styles.
	RelationshipTypeTableStyles = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"

	// RelationshipTypeCommentAuthors is the relationship type for comment authors.
	RelationshipTypeCommentAuthors = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/commentAuthors"

	// RelationshipTypeTags is the relationship type for tags.
	RelationshipTypeTags = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tags"
)

// Document represents a PowerPoint presentation (.pptx, .potx, .ppsx, .pptm, .potm, .ppsm, .ppam).
// It wraps an OpenXmlPackage and provides PowerPoint-specific functionality.
type Document struct {
	// pkg is the underlying OpenXML package.
	pkg *openxml.OpenXmlPackage

	// docType is the type of this document.
	docType DocType

	// presentationPart is the cached main presentation part.
	presentationPart *parts.PresentationPart

	// settings are the options used to open/create this document.
	settings *OpenSettings

	// path is the file path (if opened from file).
	path string

	// isEditable indicates if the document was opened for editing.
	isEditable bool
}

// New creates a new PowerPoint presentation at the specified path.
// The document type determines the file extension and content types.
func New(
	path string,
	docType DocType,
) (*Document, error) {
	// Create the underlying package
	pkg, err := openxml.CreatePackage(path)
	if err != nil {
		return nil, err
	}

	doc := &Document{
		pkg:        pkg,
		docType:    docType,
		settings:   DefaultOpenSettings(),
		path:       path,
		isEditable: true,
	}

	// Initialize the document structure
	if err := doc.initializeDocument(); err != nil {
		_ = pkg.Close()

		return nil, err
	}

	return doc, nil
}

// NewWriter creates a new PowerPoint presentation that writes to the given io.Writer.
func NewWriter(
	w io.Writer,
	docType DocType,
) (*Document, error) {
	// Create the underlying package
	pkg, err := packaging.CreateWriter(w)
	if err != nil {
		return nil, err
	}

	doc := &Document{
		pkg: openxml.NewOpenXmlPackage(
			pkg,
		),
		docType:    docType,
		settings:   DefaultOpenSettings(),
		isEditable: true,
	}

	// Initialize the document structure
	if err := doc.initializeDocument(); err != nil {
		_ = doc.pkg.Close()

		return nil, err
	}

	return doc, nil
}

// Open opens an existing PowerPoint presentation from the specified path.
// If editable is true, the document can be modified and saved.
func Open(
	path string,
	editable bool,
) (*Document, error) {
	return OpenWithSettings(
		path,
		editable,
		DefaultOpenSettings(),
	)
}

// OpenReader opens a PowerPoint presentation from an io.ReaderAt.
func OpenReader(
	r io.ReaderAt,
	size int64,
	editable bool,
) (*Document, error) {
	pkg, err := openxml.OpenPackageFromReader(
		r,
		size,
	)
	if err != nil {
		return nil, err
	}

	doc := &Document{
		pkg:        pkg,
		settings:   DefaultOpenSettings(),
		isEditable: editable,
	}

	// Detect document type from content type
	doc.detectDocumentType()

	return doc, nil
}

// OpenWithSettings opens a PowerPoint presentation with custom settings.
func OpenWithSettings(
	path string,
	editable bool,
	settings *OpenSettings,
) (*Document, error) {
	pkg, err := openxml.OpenPackage(
		path,
		!editable,
	)
	if err != nil {
		return nil, err
	}

	doc := &Document{
		pkg:        pkg,
		settings:   settings,
		path:       path,
		isEditable: editable,
	}

	// Detect document type from content type
	doc.detectDocumentType()

	return doc, nil
}

// CreateFromTemplate creates a new presentation from a template.
func CreateFromTemplate(
	templatePath string,
) (*Document, error) {
	// Open the template as read-only
	templateDoc, err := Open(templatePath, false)
	if err != nil {
		return nil, err
	}
	defer func() { _ = templateDoc.Close() }()

	// Create a new document
	// Note: We don't save to a path yet - caller should use SaveAs
	pkg, err := packaging.CreateWriter(io.Discard)
	if err != nil {
		return nil, err
	}

	doc := &Document{
		pkg: openxml.NewOpenXmlPackage(
			pkg,
		),
		docType:    DocTypePresentation, // Templates create presentations
		settings:   DefaultOpenSettings(),
		isEditable: true,
	}

	// Copy content from template
	// For now, just initialize with basic structure
	if err := doc.initializeDocument(); err != nil {
		_ = doc.pkg.Close()

		return nil, err
	}

	return doc, nil
}

// initializeDocument sets up the basic document structure for a new document.
func (d *Document) initializeDocument() error {
	// Create the main presentation part using the parts package
	presPart, err := parts.NewPresentationPart(
		"/ppt/presentation.xml",
		d.docType.ContentType(),
		d.pkg,
	)
	if err != nil {
		return err
	}

	// Initialize with minimal presentation content
	presPart.InitializeContent()

	// Set as main part in the package
	d.pkg.SetMainPart(presPart)
	d.presentationPart = presPart

	return nil
}

// detectDocumentType determines the document type from the main part
// content type and wraps it as a PresentationPart.
func (d *Document) detectDocumentType() {
	mainPart := d.pkg.MainPart()
	if mainPart == nil {
		d.docType = DocTypePresentation

		return
	}

	// Determine document type from content type
	switch mainPart.ContentType() {
	case ContentTypePresentationMLTemplate:
		d.docType = DocTypeTemplate
	case ContentTypePresentationMLSlideshow:
		d.docType = DocTypeSlideshow
	case ContentTypePresentationMLMacroEnabled:
		d.docType = DocTypeMacroEnabledPresentation
	case ContentTypePresentationMLMacroTemplate:
		d.docType = DocTypeMacroEnabledTemplate
	case ContentTypePresentationMLMacroSlideshow:
		d.docType = DocTypeMacroEnabledSlideshow
	case ContentTypePresentationMLAddIn:
		d.docType = DocTypeAddIn
	default:
		d.docType = DocTypePresentation
	}

	// Wrap the main part as a PresentationPart if it's the right type
	switch mp := mainPart.(type) {
	case *parts.PresentationPart:
		d.presentationPart = mp
	case *openxml.OpenXmlPartData:
		// Wrap the existing part data as a PresentationPart
		d.presentationPart = parts.NewPresentationPartFromData(
			mp,
			mainPart.ContentType(),
		)
	}
}

// Type returns the document type.
func (d *Document) Type() DocType {
	return d.docType
}

// Package returns the underlying OpenXML package.
func (d *Document) Package() *openxml.OpenXmlPackage {
	return d.pkg
}

// Features returns the document's feature collection.
func (d *Document) Features() *features.FeatureCollection {
	return d.pkg.Features()
}

// IsEditable returns true if the document was opened for editing.
func (d *Document) IsEditable() bool {
	return d.isEditable
}

// Path returns the file path of the document.
// Returns empty string if the document was created from a stream.
func (d *Document) Path() string {
	return d.path
}

// Settings returns the open settings for this document.
func (d *Document) Settings() *OpenSettings {
	return d.settings
}

// PresentationPart returns the main presentation part.
// Returns nil if no main presentation part exists or the document is closed.
func (d *Document) PresentationPart() *parts.PresentationPart {
	if d.pkg == nil {
		return nil
	}

	if d.presentationPart != nil {
		return d.presentationPart
	}

	// Look for the main part in the package and wrap it
	mainPart := d.pkg.MainPart()
	if mainPart == nil {
		return nil
	}

	// Try to get an existing PresentationPart or wrap the main part
	switch mp := mainPart.(type) {
	case *parts.PresentationPart:
		d.presentationPart = mp
	case *openxml.OpenXmlPartData:
		d.presentationPart = parts.NewPresentationPartFromData(
			mp,
			mainPart.ContentType(),
		)
	}

	return d.presentationPart
}

// AddPresentationPart creates and adds the main presentation part.
// Returns an error if a main part already exists.
func (d *Document) AddPresentationPart() (*parts.PresentationPart, error) {
	if d.presentationPart != nil {
		return nil, ErrPresentationPartExists
	}

	// Create the main presentation part using the parts package
	presPart, err := parts.NewPresentationPart(
		"/ppt/presentation.xml",
		d.docType.ContentType(),
		d.pkg,
	)
	if err != nil {
		return nil, err
	}

	// Initialize with minimal presentation content
	presPart.InitializeContent()

	// Set as main part
	d.pkg.SetMainPart(presPart)
	d.presentationPart = presPart

	return presPart, nil
}

// CoreProperties returns the core file properties part.
// Returns nil if not present.
func (d *Document) CoreProperties() *packaging.CoreProperties {
	return d.pkg.Package().CoreProperties()
}

// ExtendedProperties returns the extended file properties part.
// Returns nil if not present.
func (d *Document) ExtendedProperties() openxml.OpenXmlPart {
	for part := range d.pkg.GetPartsOfType(openxml.ContentTypeExtendedProperties) {
		return part
	}

	return nil
}

// Save saves the document to its original location.
// Returns an error if the document is read-only or was created from a stream.
func (d *Document) Save() error {
	if !d.isEditable {
		return ErrReadOnly
	}

	if d.path == "" {
		return ErrNoPath
	}

	return d.pkg.Save()
}

// SaveAs saves the document to a new file path.
func (d *Document) SaveAs(path string) error {
	return d.pkg.SaveAs(path)
}

// SaveTo writes the document to the given io.Writer.
// The document is written as a ZIP (Office Open XML format) containing
// all parts, relationships, and content types.
func (d *Document) SaveTo(
	w io.Writer,
) error {
	if d.pkg == nil {
		return ErrPackageNil
	}

	pkg := d.pkg.Package()
	if pkg == nil {
		return ErrPackageNil
	}

	// Use the packaging layer's saveToWriter method to write the entire
	// document package (all parts, relationships, content types) to the writer
	return pkg.SaveToWriter(w)
}

// Close closes the document and releases all resources.
// If AutoSave is enabled and the document has unsaved changes,
// it will be saved.
func (d *Document) Close() error {
	if d.pkg == nil {
		return nil
	}

	// Auto-save if enabled
	if d.settings.AutoSave && d.isEditable &&
		d.path != "" &&
		d.pkg.IsDirty() {
		if err := d.Save(); err != nil {
			// Log but don't fail on auto-save errors
			_ = err
		}
	}

	err := d.pkg.Close()
	d.pkg = nil
	d.presentationPart = nil

	return err
}

// ChangeType changes the document type.
// This updates the content type of the main presentation part.
func (d *Document) ChangeType(
	newType DocType,
) error {
	if !d.isEditable {
		return ErrReadOnly
	}

	if newType == d.docType {
		return nil
	}

	// Update the content type of the main part
	mainPart := d.PresentationPart()
	if mainPart == nil {
		return ErrNoPresentationPart
	}

	// Update content type in the package
	ct := d.pkg.Package().ContentTypes()
	if ct != nil {
		ct.SetOverride(
			mainPart.URI(),
			newType.ContentType(),
		)
	}

	d.docType = newType

	return nil
}

// Validate validates the document structure against the specified
// Office version. It checks for schema compliance and semantic
// constraints. Returns validation errors if any issues are found,
// or an empty slice if valid.
func (d *Document) Validate(
	version validation.FileFormatVersions,
) validation.ValidationErrors {
	return d.ValidateWithSettings(
		version,
		validation.DefaultSettings(),
	)
}

// ValidateWithSettings validates the document with custom validation
// settings. This allows control over validation behavior such as
// maximum errors to collect, whether to validate semantics,
// and strict mode.
func (d *Document) ValidateWithSettings(
	version validation.FileFormatVersions,
	settings *validation.ValidationSettings,
) validation.ValidationErrors {
	if d.pkg == nil {
		return validation.ValidationErrors{
			validation.NewValidationError(
				validation.SchemaMissingRequiredElement,
				"Document package is nil or closed",
				"/",
				nil,
			),
		}
	}

	// Use the validation framework to validate the package
	return validation.Validate(
		d.pkg,
		version,
		settings,
	)
}

// IsValid returns true if the document has no validation errors for
// the specified version. This is a convenience method that calls
// Validate and checks if there are any errors.
func (d *Document) IsValid(
	version validation.FileFormatVersions,
) bool {
	errors := d.Validate(version)

	return !errors.HasErrors()
}

// Slide Methods

// AddSlide adds a new slide to the presentation.
// Returns the newly created SlidePart.
func (d *Document) AddSlide() (*parts.SlidePart, error) {
	if !d.isEditable {
		return nil, ErrReadOnly
	}

	presPart := d.PresentationPart()
	if presPart == nil {
		return nil, ErrNoPresentationPart
	}

	return presPart.AddSlidePart()
}

// Slides returns an iterator over all slides in the presentation.
func (d *Document) Slides() iter.Seq[*parts.SlidePart] {
	return func(yield func(*parts.SlidePart) bool) {
		presPart := d.PresentationPart()
		if presPart == nil {
			return
		}

		for _, slide := range presPart.SlideParts() {
			if !yield(slide) {
				return
			}
		}
	}
}

// SlideCount returns the number of slides in the presentation.
func (d *Document) SlideCount() int {
	presPart := d.PresentationPart()
	if presPart == nil {
		return 0
	}

	return len(presPart.SlideParts())
}

// GetSlide returns the slide at the specified index (0-based).
// Returns an error if the index is out of range.
func (d *Document) GetSlide(
	index int,
) (*parts.SlidePart, error) {
	presPart := d.PresentationPart()
	if presPart == nil {
		return nil, ErrNoPresentationPart
	}

	slides := presPart.SlideParts()
	if index < 0 || index >= len(slides) {
		return nil, ErrSlideIndexOutOfRange
	}

	return slides[index], nil
}

// Slide Master and Layout Methods

// AddSlideMaster adds a new slide master to the presentation.
// Returns the newly created SlideMasterPart.
func (d *Document) AddSlideMaster() (*parts.SlideMasterPart, error) {
	if !d.isEditable {
		return nil, ErrReadOnly
	}

	presPart := d.PresentationPart()
	if presPart == nil {
		return nil, ErrNoPresentationPart
	}

	return presPart.AddSlideMasterPart()
}

// SlideMasters returns an iterator over all slide masters in the presentation.
func (d *Document) SlideMasters() iter.Seq[*parts.SlideMasterPart] {
	return func(yield func(*parts.SlideMasterPart) bool) {
		presPart := d.PresentationPart()
		if presPart == nil {
			return
		}

		for _, master := range presPart.SlideMasterParts() {
			if !yield(master) {
				return
			}
		}
	}
}

// Theme Methods

// AddTheme adds a theme to the presentation.
// Returns the newly created ThemePart.
func (d *Document) AddTheme() (*parts.ThemePart, error) {
	if !d.isEditable {
		return nil, ErrReadOnly
	}

	presPart := d.PresentationPart()
	if presPart == nil {
		return nil, ErrNoPresentationPart
	}

	return presPart.AddThemePart()
}

// Theme returns the first theme part of the presentation.
// Returns nil if no theme is present.
func (d *Document) Theme() *parts.ThemePart {
	presPart := d.PresentationPart()
	if presPart == nil {
		return nil
	}

	return presPart.ThemePart()
}

// Errors
var (
	ErrReadOnly = docError(
		"document is read-only",
	)
	ErrNoPath = docError(
		"document has no file path",
	)
	ErrPackageNil = docError(
		"package is nil",
	)
	ErrPresentationPartExists = docError(
		"presentation part already exists",
	)
	ErrNoPresentationPart = docError(
		"no presentation part",
	)
	ErrPartExists = docError(
		"part already exists",
	)
	ErrPartNotFound = docError(
		"part not found",
	)
	ErrSlideIndexOutOfRange = docError(
		"slide index out of range",
	)
)

type docError string

func (e docError) Error() string {
	return string(e)
}

// Ensure Document implements io.Closer.
var _ io.Closer = (*Document)(nil)
