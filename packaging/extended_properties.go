package packaging

import (
	"bytes"
	"encoding/xml"
	"io"
	"sync"
)

// ExtendedProperties represents the extended/app document properties.
// These are stored in docProps/app.xml in the package.
// Extended properties include application-specific metadata like
// application name, version, company, manager, template, document
// statistics, etc.
type ExtendedProperties struct {
	mu sync.RWMutex

	// Application properties
	application        string
	applicationVersion string
	template           string
	manager            string
	company            string

	// Document statistics
	pages                int
	words                int
	characters           int
	charactersWithSpaces int
	lines                int
	paragraphs           int

	// Other properties
	totalTime          int // in minutes
	hyperlinkBase      string
	presentationFormat string
	documentSecurity   int

	// Presentation-specific
	slides          int
	hiddenSlides    int
	notes           int
	multimediaClips int

	// Boolean flags
	scaleCrop         bool
	linksUpToDate     bool
	sharedDocument    bool
	hyperlinksChanged bool

	// Digital signature
	digitalSignature string
}

// NewExtendedProperties creates a new ExtendedProperties instance.
func NewExtendedProperties() *ExtendedProperties {
	return &ExtendedProperties{}
}

// Application returns the name of the application that created the document.
func (ep *ExtendedProperties) Application() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.application
}

// SetApplication sets the name of the application that created the document.
func (ep *ExtendedProperties) SetApplication(
	app string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.application = app
}

// ApplicationVersion returns the version of the application.
func (ep *ExtendedProperties) ApplicationVersion() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.applicationVersion
}

// SetApplicationVersion sets the version of the application.
func (ep *ExtendedProperties) SetApplicationVersion(
	version string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.applicationVersion = version
}

// Template returns the template used for the document.
func (ep *ExtendedProperties) Template() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.template
}

// SetTemplate sets the template used for the document.
func (ep *ExtendedProperties) SetTemplate(
	template string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.template = template
}

// Manager returns the manager name.
func (ep *ExtendedProperties) Manager() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.manager
}

// SetManager sets the manager name.
func (ep *ExtendedProperties) SetManager(
	manager string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.manager = manager
}

// Company returns the company name.
func (ep *ExtendedProperties) Company() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.company
}

// SetCompany sets the company name.
func (ep *ExtendedProperties) SetCompany(
	company string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.company = company
}

// Pages returns the total number of pages.
func (ep *ExtendedProperties) Pages() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.pages
}

// SetPages sets the total number of pages.
func (ep *ExtendedProperties) SetPages(
	pages int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.pages = pages
}

// Words returns the total word count.
func (ep *ExtendedProperties) Words() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.words
}

// SetWords sets the total word count.
func (ep *ExtendedProperties) SetWords(
	words int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.words = words
}

// Characters returns the total character count.
func (ep *ExtendedProperties) Characters() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.characters
}

// SetCharacters sets the total character count.
func (ep *ExtendedProperties) SetCharacters(
	characters int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.characters = characters
}

// CharactersWithSpaces returns the character count including spaces.
func (ep *ExtendedProperties) CharactersWithSpaces() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.charactersWithSpaces
}

// SetCharactersWithSpaces sets the character count including spaces.
func (ep *ExtendedProperties) SetCharactersWithSpaces(
	count int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.charactersWithSpaces = count
}

// Lines returns the total line count.
func (ep *ExtendedProperties) Lines() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.lines
}

// SetLines sets the total line count.
func (ep *ExtendedProperties) SetLines(
	lines int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.lines = lines
}

// Paragraphs returns the total paragraph count.
func (ep *ExtendedProperties) Paragraphs() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.paragraphs
}

// SetParagraphs sets the total paragraph count.
func (ep *ExtendedProperties) SetParagraphs(
	paragraphs int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.paragraphs = paragraphs
}

// TotalTime returns the total editing time in minutes.
func (ep *ExtendedProperties) TotalTime() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.totalTime
}

// SetTotalTime sets the total editing time in minutes.
func (ep *ExtendedProperties) SetTotalTime(
	minutes int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.totalTime = minutes
}

// HyperlinkBase returns the base path for hyperlinks.
func (ep *ExtendedProperties) HyperlinkBase() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.hyperlinkBase
}

// SetHyperlinkBase sets the base path for hyperlinks.
func (ep *ExtendedProperties) SetHyperlinkBase(
	base string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.hyperlinkBase = base
}

// PresentationFormat returns the presentation format.
func (ep *ExtendedProperties) PresentationFormat() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.presentationFormat
}

// SetPresentationFormat sets the presentation format.
func (ep *ExtendedProperties) SetPresentationFormat(
	format string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.presentationFormat = format
}

// DocumentSecurity returns the document security level.
func (ep *ExtendedProperties) DocumentSecurity() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.documentSecurity
}

// SetDocumentSecurity sets the document security level.
func (ep *ExtendedProperties) SetDocumentSecurity(
	level int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.documentSecurity = level
}

// Slides returns the number of slides (for presentations).
func (ep *ExtendedProperties) Slides() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.slides
}

// SetSlides sets the number of slides.
func (ep *ExtendedProperties) SetSlides(
	slides int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.slides = slides
}

// HiddenSlides returns the number of hidden slides.
func (ep *ExtendedProperties) HiddenSlides() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.hiddenSlides
}

// SetHiddenSlides sets the number of hidden slides.
func (ep *ExtendedProperties) SetHiddenSlides(
	count int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.hiddenSlides = count
}

// Notes returns the number of notes.
func (ep *ExtendedProperties) Notes() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.notes
}

// SetNotes sets the number of notes.
func (ep *ExtendedProperties) SetNotes(
	notes int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.notes = notes
}

// MultimediaClips returns the number of multimedia clips.
func (ep *ExtendedProperties) MultimediaClips() int {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.multimediaClips
}

// SetMultimediaClips sets the number of multimedia clips.
func (ep *ExtendedProperties) SetMultimediaClips(
	count int,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.multimediaClips = count
}

// ScaleCrop returns whether to scale or crop the thumbnail.
func (ep *ExtendedProperties) ScaleCrop() bool {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.scaleCrop
}

// SetScaleCrop sets whether to scale or crop the thumbnail.
func (ep *ExtendedProperties) SetScaleCrop(
	scale bool,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.scaleCrop = scale
}

// LinksUpToDate returns whether all links are up to date.
func (ep *ExtendedProperties) LinksUpToDate() bool {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.linksUpToDate
}

// SetLinksUpToDate sets whether all links are up to date.
func (ep *ExtendedProperties) SetLinksUpToDate(
	upToDate bool,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.linksUpToDate = upToDate
}

// SharedDocument returns whether the document is shared.
func (ep *ExtendedProperties) SharedDocument() bool {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.sharedDocument
}

// SetSharedDocument sets whether the document is shared.
func (ep *ExtendedProperties) SetSharedDocument(
	shared bool,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.sharedDocument = shared
}

// HyperlinksChanged returns whether hyperlinks have changed.
func (ep *ExtendedProperties) HyperlinksChanged() bool {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.hyperlinksChanged
}

// SetHyperlinksChanged sets whether hyperlinks have changed.
func (ep *ExtendedProperties) SetHyperlinksChanged(
	changed bool,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.hyperlinksChanged = changed
}

// DigitalSignature returns the digital signature.
func (ep *ExtendedProperties) DigitalSignature() string {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	return ep.digitalSignature
}

// SetDigitalSignature sets the digital signature.
func (ep *ExtendedProperties) SetDigitalSignature(
	signature string,
) {
	ep.mu.Lock()
	defer ep.mu.Unlock()
	ep.digitalSignature = signature
}

// XML types for extended properties serialization.
//
//nolint:revive,lll // XML struct tags contain long namespace URLs
type xmlExtendedProperties struct {
	XMLName              xml.Name `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Properties"`
	Template             string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Template,omitempty"`
	Manager              string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Manager,omitempty"`
	Company              string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Company,omitempty"`
	Pages                int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Pages,omitempty"`
	Words                int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Words,omitempty"`
	Characters           int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Characters,omitempty"`
	PresentationFormat   string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties PresentationFormat,omitempty"`
	Lines                int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Lines,omitempty"`
	Paragraphs           int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Paragraphs,omitempty"`
	Slides               int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Slides,omitempty"`
	Notes                int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Notes,omitempty"`
	TotalTime            int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties TotalTime,omitempty"`
	HiddenSlides         int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties HiddenSlides,omitempty"`
	MultimediaClips      int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties MMClips,omitempty"`
	ScaleCrop            bool     `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties ScaleCrop,omitempty"`
	LinksUpToDate        bool     `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties LinksUpToDate,omitempty"`
	CharactersWithSpaces int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties CharactersWithSpaces,omitempty"`
	SharedDocument       bool     `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties SharedDoc,omitempty"`
	HyperlinkBase        string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties HyperlinkBase,omitempty"`
	HyperlinksChanged    bool     `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties HyperlinksChanged,omitempty"`
	DigitalSignature     string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties DigSig,omitempty"`
	Application          string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Application,omitempty"`
	ApplicationVersion   string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties AppVersion,omitempty"`
	DocumentSecurity     int      `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties DocSecurity,omitempty"`
}

// MarshalToXML serializes the ExtendedProperties to XML.
func (ep *ExtendedProperties) MarshalToXML() ([]byte, error) {
	ep.mu.RLock()
	defer ep.mu.RUnlock()

	xmlEP := xmlExtendedProperties{
		Template:             ep.template,
		Manager:              ep.manager,
		Company:              ep.company,
		Pages:                ep.pages,
		Words:                ep.words,
		Characters:           ep.characters,
		PresentationFormat:   ep.presentationFormat,
		Lines:                ep.lines,
		Paragraphs:           ep.paragraphs,
		Slides:               ep.slides,
		Notes:                ep.notes,
		TotalTime:            ep.totalTime,
		HiddenSlides:         ep.hiddenSlides,
		MultimediaClips:      ep.multimediaClips,
		ScaleCrop:            ep.scaleCrop,
		LinksUpToDate:        ep.linksUpToDate,
		CharactersWithSpaces: ep.charactersWithSpaces,
		SharedDocument:       ep.sharedDocument,
		HyperlinkBase:        ep.hyperlinkBase,
		HyperlinksChanged:    ep.hyperlinksChanged,
		DigitalSignature:     ep.digitalSignature,
		Application:          ep.application,
		ApplicationVersion:   ep.applicationVersion,
		DocumentSecurity:     ep.documentSecurity,
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)

	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlEP); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// assignFieldsFromXML assigns fields from the XML representation to
// the ExtendedProperties instance.
// This method assumes the caller holds the appropriate lock.
func (ep *ExtendedProperties) assignFieldsFromXML(
	xmlEP *xmlExtendedProperties,
) {
	ep.template = xmlEP.Template
	ep.manager = xmlEP.Manager
	ep.company = xmlEP.Company
	ep.pages = xmlEP.Pages
	ep.words = xmlEP.Words
	ep.characters = xmlEP.Characters
	ep.presentationFormat = xmlEP.PresentationFormat
	ep.lines = xmlEP.Lines
	ep.paragraphs = xmlEP.Paragraphs
	ep.slides = xmlEP.Slides
	ep.notes = xmlEP.Notes
	ep.totalTime = xmlEP.TotalTime
	ep.hiddenSlides = xmlEP.HiddenSlides
	ep.multimediaClips = xmlEP.MultimediaClips
	ep.scaleCrop = xmlEP.ScaleCrop
	ep.linksUpToDate = xmlEP.LinksUpToDate
	ep.charactersWithSpaces = xmlEP.CharactersWithSpaces
	ep.sharedDocument = xmlEP.SharedDocument
	ep.hyperlinkBase = xmlEP.HyperlinkBase
	ep.hyperlinksChanged = xmlEP.HyperlinksChanged
	ep.digitalSignature = xmlEP.DigitalSignature
	ep.application = xmlEP.Application
	ep.applicationVersion = xmlEP.ApplicationVersion
	ep.documentSecurity = xmlEP.DocumentSecurity
}

// UnmarshalFromXML deserializes the ExtendedProperties from XML.
func (ep *ExtendedProperties) UnmarshalFromXML(
	r io.Reader,
) error {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	var xmlEP xmlExtendedProperties
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&xmlEP); err != nil {
		return err
	}

	ep.assignFieldsFromXML(&xmlEP)

	return nil
}

// Extended properties content type and relationship type constants.
//
//nolint:revive,lll // Long URLs cannot be broken
const (
	ExtendedPropertiesContentType      = "application/vnd.openxmlformats-officedocument.extended-properties+xml"
	ExtendedPropertiesRelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	ExtendedPropertiesPartURI          = "/docProps/app.xml"
)
