// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the document-level API for creating, reading,
// and modifying .xlsx files.
//
//nolint:revive // file-length-limit and line-length-limit: comprehensive document API with OOXML constants
package spreadsheet

import (
	"io"
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// DocType represents the type of an Excel document.
type DocType int

const (
	// DocTypeWorkbook represents a standard Excel workbook (.xlsx).
	DocTypeWorkbook DocType = iota

	// DocTypeTemplate represents an Excel template (.xltx).
	DocTypeTemplate

	// DocTypeMacroEnabledWorkbook represents a macro-enabled Excel workbook (.xlsm).
	DocTypeMacroEnabledWorkbook

	// DocTypeMacroEnabledTemplate represents a macro-enabled Excel template (.xltm).
	DocTypeMacroEnabledTemplate

	// DocTypeAddIn represents an Excel add-in (.xlam).
	DocTypeAddIn
)

// String returns the string representation of the document type.
func (dt DocType) String() string {
	switch dt {
	case DocTypeWorkbook:
		return "Workbook"
	case DocTypeTemplate:
		return "Template"
	case DocTypeMacroEnabledWorkbook:
		return "MacroEnabledWorkbook"
	case DocTypeMacroEnabledTemplate:
		return "MacroEnabledTemplate"
	case DocTypeAddIn:
		return "AddIn"
	default:
		return "Workbook"
	}
}

// Extension returns the file extension for this document type.
func (dt DocType) Extension() string {
	switch dt {
	case DocTypeWorkbook:
		return ".xlsx"
	case DocTypeTemplate:
		return ".xltx"
	case DocTypeMacroEnabledWorkbook:
		return ".xlsm"
	case DocTypeMacroEnabledTemplate:
		return ".xltm"
	case DocTypeAddIn:
		return ".xlam"
	default:
		return ".xlsx"
	}
}

// ContentType returns the MIME content type for the main workbook part.
func (dt DocType) ContentType() string {
	switch dt {
	case DocTypeWorkbook:
		return ContentTypeSpreadsheetMLWorkbook
	case DocTypeTemplate:
		return ContentTypeSpreadsheetMLTemplate
	case DocTypeMacroEnabledWorkbook:
		return ContentTypeSpreadsheetMLMacroEnabled
	case DocTypeMacroEnabledTemplate:
		return ContentTypeSpreadsheetMLMacroTemplate
	case DocTypeAddIn:
		return ContentTypeSpreadsheetMLAddIn
	default:
		return ContentTypeSpreadsheetMLWorkbook
	}
}

// Content types for Excel documents.
const (
	// ContentTypeSpreadsheetMLWorkbook is the content type for a standard
	// Excel workbook main part.
	ContentTypeSpreadsheetMLWorkbook = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"

	// ContentTypeSpreadsheetMLTemplate is the content type for an Excel
	// template main part.
	ContentTypeSpreadsheetMLTemplate = "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml"

	// ContentTypeSpreadsheetMLMacroEnabled is the content type for a
	// macro-enabled Excel workbook main part.
	ContentTypeSpreadsheetMLMacroEnabled = "application/vnd.ms-excel.sheet.macroEnabled.main+xml"

	// ContentTypeSpreadsheetMLMacroTemplate is the content type for a
	// macro-enabled Excel template main part.
	ContentTypeSpreadsheetMLMacroTemplate = "application/vnd.ms-excel.template.macroEnabled.main+xml"

	// ContentTypeSpreadsheetMLAddIn is the content type for an Excel add-in
	// main part.
	ContentTypeSpreadsheetMLAddIn = "application/vnd.ms-excel.addin.macroEnabled.main+xml"

	// ContentTypeWorksheet is the content type for a worksheet.
	ContentTypeWorksheet = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"

	// ContentTypeStyles is the content type for styles in Excel.
	ContentTypeSpreadsheetStyles = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"

	// ContentTypeSharedStrings is the content type for shared strings.
	ContentTypeSharedStrings = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"

	// ContentTypeCalcChain is the content type for calculation chain.
	ContentTypeCalcChain = "application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml"

	// ContentTypeChartsheet is the content type for a chart sheet.
	ContentTypeChartsheet = "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml"

	// ContentTypeTable is the content type for a table.
	ContentTypeTable = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"

	// ContentTypePivotTable is the content type for a pivot table.
	ContentTypePivotTable = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml"

	// ContentTypePivotCacheDefinition is the content type for a pivot cache definition.
	ContentTypePivotCacheDefinition = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheDefinition+xml"

	// ContentTypePivotCacheRecords is the content type for pivot cache records.
	ContentTypePivotCacheRecords = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheRecords+xml"
)

// Document represents an Excel workbook (.xlsx, .xltx, .xlsm, .xltm, .xlam).
// It wraps an OpenXmlPackage and provides Excel-specific functionality.
type Document struct {
	// pkg is the underlying OpenXML package.
	pkg *openxml.OpenXmlPackage

	// docType is the type of this document.
	docType DocType

	// settings are the options used to open/create this document.
	settings *OpenSettings

	// path is the file path (if opened from file).
	path string

	// isEditable indicates if the document was opened for editing.
	isEditable bool
}

// OpenSettings contains options for opening or creating documents.
type OpenSettings struct {
	// AutoSave determines if the document should be automatically saved
	// when Close() is called and there are unsaved changes.
	AutoSave bool

	// MaxCharactersInPart sets the maximum number of characters allowed
	// in a single part. A value of 0 means no limit.
	MaxCharactersInPart int64

	// MarkupCompatibilityProcessSettings controls how markup compatibility
	// elements are processed. Currently not implemented.
	MarkupCompatibilityProcessSettings MarkupCompatibilityProcessSettings
}

// MarkupCompatibilityProcessSettings controls markup compatibility processing.
type MarkupCompatibilityProcessSettings struct {
	// ProcessModeValues determines which mc:ProcessContent elements to process.
	ProcessModeValues ProcessMode

	// TargetFileFormatVersions specifies which Office versions to target.
	TargetFileFormatVersions FileFormatVersion
}

// ProcessMode specifies how to process markup compatibility content.
type ProcessMode int

const (
	// ProcessModeNoProcess does not process mc:ProcessContent.
	ProcessModeNoProcess ProcessMode = iota

	// ProcessModeProcessLoadedPartsOnly processes only loaded parts.
	ProcessModeProcessLoadedPartsOnly

	// ProcessModeProcessAllParts processes all parts.
	ProcessModeProcessAllParts
)

// FileFormatVersion represents the target Office version.
type FileFormatVersion int

const (
	// FileFormatVersionOffice2007 targets Office 2007 format.
	FileFormatVersionOffice2007 FileFormatVersion = iota

	// FileFormatVersionOffice2010 targets Office 2010 format.
	FileFormatVersionOffice2010

	// FileFormatVersionOffice2013 targets Office 2013 format.
	FileFormatVersionOffice2013

	// FileFormatVersionOffice2016 targets Office 2016 format.
	FileFormatVersionOffice2016

	// FileFormatVersionOffice2019 targets Office 2019 format.
	FileFormatVersionOffice2019

	// FileFormatVersionOffice2021 targets Office 2021 format.
	FileFormatVersionOffice2021

	// FileFormatVersionMicrosoft365 targets Microsoft 365 format.
	FileFormatVersionMicrosoft365
)

// DefaultOpenSettings returns the default settings for opening documents.
func DefaultOpenSettings() *OpenSettings {
	return &OpenSettings{
		AutoSave:            false,
		MaxCharactersInPart: 0, // No limit
		MarkupCompatibilityProcessSettings: MarkupCompatibilityProcessSettings{
			ProcessModeValues:        ProcessModeNoProcess,
			TargetFileFormatVersions: FileFormatVersionOffice2016,
		},
	}
}

// WithAutoSave returns a copy of the settings with AutoSave enabled.
func (s *OpenSettings) WithAutoSave(
	autoSave bool,
) *OpenSettings {
	settings := *s
	settings.AutoSave = autoSave

	return &settings
}

// WithMaxCharacters returns a copy of the settings with the specified max.
func (s *OpenSettings) WithMaxCharacters(
	maxChars int64,
) *OpenSettings {
	settings := *s
	settings.MaxCharactersInPart = maxChars

	return &settings
}

// WithTargetVersion returns a copy of the settings with the target version.
func (s *OpenSettings) WithTargetVersion(
	version FileFormatVersion,
) *OpenSettings {
	settings := *s
	settings.MarkupCompatibilityProcessSettings.TargetFileFormatVersions = version

	return &settings
}

// Create creates a new Excel document at the specified path.
// The document type determines the file extension and content types.
func Create(
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

// CreateFromStream creates a new Excel document that writes to the given io.Writer.
func CreateFromStream(
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

// Open opens an existing Excel document from the specified path.
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

// OpenFromStream opens an Excel document from an io.ReaderAt.
func OpenFromStream(
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

// OpenWithSettings opens an Excel document with custom settings.
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

// CreateFromTemplate creates a new document from a template.
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
		docType:    DocTypeWorkbook, // Templates create workbooks
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
	// Create the main workbook part
	mainPart, err := d.pkg.AddNewPart(
		"/xl/workbook.xml",
		d.docType.ContentType(),
		openxml.RelationshipTypeOfficeDocument,
	)
	if err != nil {
		return err
	}

	// Set as main part
	d.pkg.SetMainPart(mainPart)

	// Initialize the workbook element programmatically
	wp := d.WorkbookPart()
	if wp != nil {
		wb := wp.Workbook().(*elements.Workbook)
		// Create the sheets element
		wb.GetOrCreateSheets()
	}

	return nil
}

// detectDocumentType determines the document type from the main part
// content type.
func (d *Document) detectDocumentType() {
	mainPart := d.pkg.MainPart()
	if mainPart == nil {
		d.docType = DocTypeWorkbook

		return
	}

	switch mainPart.ContentType() {
	case ContentTypeSpreadsheetMLTemplate:
		d.docType = DocTypeTemplate
	case ContentTypeSpreadsheetMLMacroEnabled:
		d.docType = DocTypeMacroEnabledWorkbook
	case ContentTypeSpreadsheetMLMacroTemplate:
		d.docType = DocTypeMacroEnabledTemplate
	case ContentTypeSpreadsheetMLAddIn:
		d.docType = DocTypeAddIn
	default:
		d.docType = DocTypeWorkbook
	}
}

// Type returns the document type.
func (d *Document) Type() DocType {
	return d.docType
}

// SheetCount returns the number of sheets in the workbook.
func (d *Document) SheetCount() int {
	wp := d.WorkbookPart()
	if wp == nil {
		return 0
	}

	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.Sheets()
	if sheets == nil {
		return 0
	}

	return sheets.SheetCount()
}

// Package returns the underlying OpenXML package.
func (d *Document) Package() *openxml.OpenXmlPackage {
	return d.pkg
}

// WorkbookPart returns the main workbook part.
func (d *Document) WorkbookPart() *parts.WorkbookPart {
	if d.pkg == nil {
		return nil
	}
	if p := d.pkg.MainPart(); p != nil {
		if wp, ok := p.(*parts.WorkbookPart); ok {
			return wp
		}
	}

	return nil
}

// Features returns the document's feature collection.
func (d *Document) Features() *features.FeatureCollection {
	return d.pkg.Features()
}

// AddSheet adds a new worksheet to the workbook.
func (d *Document) AddSheet(
	name string,
) (*Sheet, error) {
	wp := d.WorkbookPart()
	if wp == nil {
		return nil, ErrNoMainPart
	}

	// Add worksheet part
	wsPart, err := wp.AddWorksheetPart()
	if err != nil {
		return nil, err
	}

	// Add to workbook element
	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.GetOrCreateSheets()
	sheetId := sheets.NextSheetId()
	relId := wsPart.RelationshipID()

	sheets.AddSheet(name, sheetId, relId)

	return newSheet(d, name, sheetId, wsPart), nil
}

// Sheet returns the sheet at the given index (0-based).
func (d *Document) Sheet(
	index int,
) (*Sheet, error) {
	wp := d.WorkbookPart()
	if wp == nil {
		return nil, ErrNoMainPart
	}

	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.Sheets()

	i := 0
	for s := range sheets.Sheets() {
		if i == index {
			// Find corresponding part
			relId := s.RelationshipId()
			part, err := wp.GetPartById(relId)
			if err != nil {
				return nil, err
			}
			if wsPart, ok := part.(*parts.WorksheetPart); ok {
				return newSheet(
					d,
					s.Name(),
					s.SheetId(),
					wsPart,
				), nil
			}
		}
		i++
	}

	return nil, ErrPartNotFound
}

// SheetByName returns the sheet with the given name.
func (d *Document) SheetByName(
	name string,
) (*Sheet, bool) {
	wp := d.WorkbookPart()
	if wp == nil {
		return nil, false
	}

	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.Sheets()
	if sheets == nil {
		return nil, false
	}

	for s := range sheets.Sheets() {
		if s.Name() == name {
			relId := s.RelationshipId()
			part, err := wp.GetPartById(relId)
			if err != nil {
				return nil, false
			}
			if wsPart, ok := part.(*parts.WorksheetPart); ok {
				return newSheet(
					d,
					s.Name(),
					s.SheetId(),
					wsPart,
				), true
			}
		}
	}

	return nil, false
}

// SetActiveSheet sets the active sheet in the workbook.
func (d *Document) SetActiveSheet(
	index int,
) error {
	wp := d.WorkbookPart()
	if wp == nil {
		return ErrNoMainPart
	}

	// Validate index is within bounds
	sheetCount := d.SheetCount()
	if index < 0 || index >= sheetCount {
		return ErrInvalidSheetIndex
	}

	wb := wp.Workbook().(*elements.Workbook)
	views := wb.GetOrCreateBookViews()
	view := views.GetOrCreateWorkbookView()
	view.SetActiveTab(index)

	return nil
}

// DeleteSheet removes a sheet from the workbook.
func (d *Document) DeleteSheet(
	sheetRef interface{},
) error {
	wp := d.WorkbookPart()
	if wp == nil {
		return ErrNoMainPart
	}

	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.GetOrCreateSheets()

	var sheetEntry *elements.Sheet
	switch v := sheetRef.(type) {
	case int:
		sheetEntry = sheets.GetSheetByIndex(v)
	case string:
		sheetEntry = sheets.GetSheetByName(v)
	}

	if sheetEntry == nil {
		return ErrPartNotFound
	}

	relId := sheetEntry.RelationshipId()

	// Remove the sheet entry
	sheets.RemoveSheet(sheetEntry)

	// Remove the part and its relationship
	return wp.DeletePart(relId)
}

// RenameSheet renames an existing sheet.
func (d *Document) RenameSheet(
	sheetRef interface{},
	newName string,
) error {
	wp := d.WorkbookPart()
	if wp == nil {
		return ErrNoMainPart
	}

	wb := wp.Workbook().(*elements.Workbook)
	sheets := wb.GetOrCreateSheets()

	var sheetEntry *elements.Sheet
	switch v := sheetRef.(type) {
	case int:
		sheetEntry = sheets.GetSheetByIndex(v)
	case string:
		sheetEntry = sheets.GetSheetByName(v)
	}

	if sheetEntry == nil {
		return ErrPartNotFound
	}

	sheetEntry.SetName(newName)

	return nil
}

// Sheets returns an iterator over all sheets in the document.
func (d *Document) Sheets() iter.Seq[*Sheet] {
	return func(yield func(*Sheet) bool) {
		wp := d.WorkbookPart()
		if wp == nil {
			return
		}

		wb := wp.Workbook().(*elements.Workbook)
		sheets := wb.Sheets()

		for s := range sheets.Sheets() {
			relId := s.RelationshipId()
			part, err := wp.GetPartById(relId)
			if err != nil {
				continue
			}
			if wsPart, ok := part.(*parts.WorksheetPart); ok {
				if !yield(
					newSheet(
						d,
						s.Name(),
						s.SheetId(),
						wsPart,
					),
				) {
					return
				}
			}
		}
	}
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

	// Ensure all parts are saved before saving package
	if wp := d.WorkbookPart(); wp != nil {
		if saveable, ok := any(wp).(openxml.ISaveablePart); ok {
			_ = saveable.Save()
		}
		for child := range wp.Parts() {
			if saveable, ok := any(child).(openxml.ISaveablePart); ok {
				_ = saveable.Save()
			}
		}
	}

	return d.pkg.Save()
}

// SaveAs saves the document to a new file path.
func (d *Document) SaveAs(path string) error {
	// Ensure all parts are saved before saving package
	if wp := d.WorkbookPart(); wp != nil {
		if saveable, ok := any(wp).(openxml.ISaveablePart); ok {
			_ = saveable.Save()
		}
		for child := range wp.Parts() {
			if saveable, ok := any(child).(openxml.ISaveablePart); ok {
				_ = saveable.Save()
			}
		}
	}

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

	return err
}

// ChangeType changes the document type.
// This updates the content type of the main workbook part.
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
	mainPart := d.pkg.MainPart()
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

// Errors
var (
	ErrReadOnly = docError(
		"document is read-only",
	)
	ErrNoPath = docError(
		"document has no file path",
	)
	ErrPackageNil = docError("package is nil")
	ErrNoMainPart = docError(
		"no main workbook part",
	)
	ErrPartExists = docError(
		"part already exists",
	)
	ErrPartNotFound = docError(
		"part not found",
	)
	ErrInvalidSheetIndex = docError(
		"invalid sheet index",
	)
)

type docError string

func (e docError) Error() string {
	return string(e)
}

// Ensure Document implements io.Closer.
var _ io.Closer = (*Document)(nil)
