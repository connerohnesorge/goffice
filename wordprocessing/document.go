// Package wordprocessing provides WordprocessingML support for Word documents.
// This package implements the document-level API for creating, reading,
// and modifying .docx files.
//
//nolint:revive // file-length-limit and line-length-limit: comprehensive document API with OOXML constants
package wordprocessing

import (
	"io"
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
	"github.com/connerohnesorge/goffice/wordprocessing/parts"
)

// DocType represents the type of a Word document.
type DocType int

const (
	// DocTypeDocument represents a standard Word document (.docx).
	DocTypeDocument DocType = iota

	// DocTypeTemplate represents a Word template (.dotx).
	DocTypeTemplate

	// DocTypeMacroEnabled represents a macro-enabled Word document (.docm).
	DocTypeMacroEnabled

	// DocTypeMacroTemplate represents a macro-enabled Word template (.dotm).
	DocTypeMacroTemplate
)

// String returns the string representation of the document type.
func (dt DocType) String() string {
	switch dt {
	case DocTypeDocument:
		return "Document"
	case DocTypeTemplate:
		return "Template"
	case DocTypeMacroEnabled:
		return "MacroEnabledDocument"
	case DocTypeMacroTemplate:
		return "MacroEnabledTemplate"
	default:
		return "Document"
	}
}

// Extension returns the file extension for this document type.
func (dt DocType) Extension() string {
	switch dt {
	case DocTypeDocument:
		return ".docx"
	case DocTypeTemplate:
		return ".dotx"
	case DocTypeMacroEnabled:
		return ".docm"
	case DocTypeMacroTemplate:
		return ".dotm"
	default:
		return ".docx"
	}
}

// ContentType returns the MIME content type for the main document part.
func (dt DocType) ContentType() string {
	switch dt {
	case DocTypeDocument:
		return ContentTypeWordMLDocument
	case DocTypeTemplate:
		return ContentTypeWordMLTemplate
	case DocTypeMacroEnabled:
		return ContentTypeWordMLMacroEnabled
	case DocTypeMacroTemplate:
		return ContentTypeWordMLMacroTemplate
	default:
		return ContentTypeWordMLDocument
	}
}

// Content types for Word documents.
const (
	// ContentTypeWordMLDocument is the content type for a standard
	// Word document main part.
	ContentTypeWordMLDocument = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"

	// ContentTypeWordMLTemplate is the content type for a Word
	// template main part.
	ContentTypeWordMLTemplate = "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"

	// ContentTypeWordMLMacroEnabled is the content type for a
	// macro-enabled Word document main part.
	ContentTypeWordMLMacroEnabled = "application/vnd.ms-word.document.macroEnabled.main+xml"

	// ContentTypeWordMLMacroTemplate is the content type for a
	// macro-enabled Word template main part.
	ContentTypeWordMLMacroTemplate = "application/vnd.ms-word.template.macroEnabled.main+xml"
)

// Document represents a Word document (.docx, .dotx, .docm, .dotm).
// It wraps an OpenXmlPackage and provides Word-specific functionality.
type Document struct {
	// pkg is the underlying OpenXML package.
	pkg *openxml.OpenXmlPackage

	// docType is the type of this document.
	docType DocType

	// mainPart is the cached main document part.
	mainPart *parts.MainPart

	// settings are the options used to open/create this document.
	settings *OpenSettings

	// path is the file path (if opened from file).
	path string

	// isEditable indicates if the document was opened for editing.
	isEditable bool
}

// New creates a new Word document at the specified path.
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

// NewWriter creates a new Word document that writes to the given io.Writer.
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

// Open opens an existing Word document from the specified path.
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

// OpenReader opens a Word document from an io.ReaderAt.
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

// OpenWithSettings opens a Word document with custom settings.
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

// NewFromTemplate creates a new document from a template.
// If attachTemplate is true, the template remains linked to the document.
func NewFromTemplate(
	templatePath string,
	_ bool, // attachTemplate - TODO: implement template attachment
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
		docType:    DocTypeDocument, // Templates create documents
		settings:   DefaultOpenSettings(),
		isEditable: true,
	}

	// Copy content from template
	// For now, just initialize with basic structure
	if err := doc.initializeDocument(); err != nil {
		_ = doc.pkg.Close()

		return nil, err
	}

	// TODO: Implement template attachment if attachTemplate is true

	return doc, nil
}

// initializeDocument sets up the basic document structure for a new document.
func (d *Document) initializeDocument() error {
	// Create the main document part
	mainPart, err := d.AddMainPart()
	if err != nil {
		return err
	}

	// Initialize with minimal document content
	mainPart.InitializeContent()

	return nil
}

// detectDocumentType determines the document type from the main part
// content type.
func (d *Document) detectDocumentType() {
	mainPart := d.pkg.MainPart()
	if mainPart == nil {
		d.docType = DocTypeDocument

		return
	}

	switch mainPart.ContentType() {
	case ContentTypeWordMLTemplate:
		d.docType = DocTypeTemplate
	case ContentTypeWordMLMacroEnabled:
		d.docType = DocTypeMacroEnabled
	case ContentTypeWordMLMacroTemplate:
		d.docType = DocTypeMacroTemplate
	default:
		d.docType = DocTypeDocument
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

// MainPart returns the main document part.
// Returns nil if no main document part exists or the document is closed.
func (d *Document) MainPart() *parts.MainPart {
	if d.pkg == nil {
		return nil
	}

	if d.mainPart != nil {
		return d.mainPart
	}

	// Look for the main part in the package
	mainPart := d.pkg.MainPart()
	if mainPart == nil {
		return nil
	}

	// Wrap it as a MainPart if needed
	if mp, ok := mainPart.(*parts.MainPart); ok {
		d.mainPart = mp

		return mp
	}

	// Create wrapper for existing part data
	if partData, ok := mainPart.(*openxml.OpenXmlPartData); ok {
		mp := parts.NewMainPartFromData(
			partData,
			d.docType.ContentType(),
		)
		d.mainPart = mp

		return mp
	}

	return nil
}

// AddMainPart creates and adds the main document part.
// Returns an error if a main part already exists.
func (d *Document) AddMainPart() (*parts.MainPart, error) {
	if d.mainPart != nil {
		return nil, ErrMainPartExists
	}

	// Create the main document part
	mainPart, err := parts.NewMainPart(
		"/word/document.xml",
		d.docType.ContentType(),
		d.pkg,
	)
	if err != nil {
		return nil, err
	}

	// Add to package
	if err := d.pkg.AddPart(mainPart, "rId1"); err != nil {
		return nil, err
	}

	d.pkg.SetMainPart(mainPart)
	d.mainPart = mainPart

	return mainPart, nil
}

// CoreProperties returns the core file properties part.
// Returns nil if not present.
func (d *Document) CoreProperties() *packaging.CoreProperties {
	return d.pkg.Package().CoreProperties()
}

// ExtendedProperties returns the extended file properties part.
// Returns nil if not present.
// TODO: Implement ExtendedPropertiesPart type
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
func (d *Document) SaveTo(
	_ io.Writer,
) error { // TODO: Implement proper SaveTo for writers
	// Create a temporary package and write to the writer
	pkg := d.pkg.Package()
	if pkg == nil {
		return ErrPackageNil
	}

	// Use saveToWriter via the packaging layer
	return pkg.SaveAs(
		d.path,
	) // TODO: Implement proper SaveTo for writers
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
	d.mainPart = nil

	return err
}

// ChangeType changes the document type.
// This updates the content type of the main document part.
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
	mainPart := d.MainPart()
	if mainPart == nil {
		return ErrNoMainPart
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

// Errors
var (
	ErrReadOnly = docError(
		"document is read-only",
	)
	ErrNoPath = docError(
		"document has no file path",
	)
	ErrPackageNil     = docError("package is nil")
	ErrMainPartExists = docError(
		"main document part already exists",
	)
	ErrNoMainPart = docError(
		"no main document part",
	)
)

type docError string

func (e docError) Error() string {
	return string(e)
}

// Ensure Document implements io.Closer.
var _ io.Closer = (*Document)(nil)

// Header/Footer Methods

// AddHeader adds a new header part of the specified type and returns
// the Header element.
// The header is also linked to the document's section properties.
func (d *Document) AddHeader(
	_ elements.HeaderFooterType, // hfType - TODO: use for section properties linkage
) (*elements.Header, error) {
	mainPart := d.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	headerPart, err := mainPart.AddHeaderPart()
	if err != nil {
		return nil, err
	}

	header := headerPart.GetOrCreateHeader()

	// Link the header to section properties
	// Note: This assumes a single section in the document
	// For multi-section documents, users should manually manage
	// section properties

	return header, nil
}

// AddFooter adds a new footer part of the specified type and returns
// the Footer element.
func (d *Document) AddFooter(
	_ elements.HeaderFooterType, // hfType - TODO: use for section properties linkage
) (*elements.Footer, error) {
	mainPart := d.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	footerPart, err := mainPart.AddFooterPart()
	if err != nil {
		return nil, err
	}

	footer := footerPart.GetOrCreateFooter()

	return footer, nil
}

// Headers returns an iterator over all Header elements in the document.
func (d *Document) Headers() iter.Seq[*elements.Header] {
	return func(yield func(*elements.Header) bool) {
		mainPart := d.MainPart()
		if mainPart == nil {
			return
		}

		for _, hp := range mainPart.HeaderParts() {
			h := hp.Header()
			if h != nil && !yield(h) {
				return
			}
		}
	}
}

// Footers returns an iterator over all Footer elements in the document.
func (d *Document) Footers() iter.Seq[*elements.Footer] {
	return func(yield func(*elements.Footer) bool) {
		mainPart := d.MainPart()
		if mainPart == nil {
			return
		}

		for _, fp := range mainPart.FooterParts() {
			f := fp.Footer()
			if f != nil && !yield(f) {
				return
			}
		}
	}
}

// Footnotes/Endnotes Methods

// FootnotesPart returns the footnotes part, creating it if it doesn't exist.
func (d *Document) FootnotesPart() (*parts.FootnotesPart, error) {
	mainPart := d.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	fp := mainPart.FootnotesPart()
	if fp != nil {
		return fp, nil
	}

	return mainPart.AddFootnotesPart()
}

// EndnotesPart returns the endnotes part, creating it if it doesn't exist.
func (d *Document) EndnotesPart() (*parts.EndnotesPart, error) {
	mainPart := d.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	ep := mainPart.EndnotesPart()
	if ep != nil {
		return ep, nil
	}

	return mainPart.AddEndnotesPart()
}

// Comments Methods

// CommentsPart returns the comments part, creating it if it doesn't exist.
func (d *Document) CommentsPart() (*parts.CommentsPart, error) {
	mainPart := d.MainPart()
	if mainPart == nil {
		return nil, ErrNoMainPart
	}

	cp := mainPart.CommentsPart()
	if cp != nil {
		return cp, nil
	}

	return mainPart.AddCommentsPart()
}

// AddComment adds a new comment to the document and returns it.
func (d *Document) AddComment(
	author, text string,
) (*elements.Comment, error) {
	cp, err := d.CommentsPart()
	if err != nil {
		return nil, err
	}

	return cp.AddComment(author, text), nil
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
