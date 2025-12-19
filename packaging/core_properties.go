// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import (
	"bytes"
	"encoding/xml"
	"io"
	"sync"
	"time"
)

// CoreProperties represents the core document properties (Dublin Core metadata).
// These are stored in docProps/core.xml in the package.
type CoreProperties struct {
	mu sync.RWMutex

	// Dublin Core elements
	title       string
	subject     string
	creator     string
	keywords    string
	description string
	language    string

	// Core properties elements
	lastModifiedBy string
	revision       string
	category       string
	contentStatus  string

	// Date/time properties
	created     *time.Time
	modified    *time.Time
	lastPrinted *time.Time
}

// NewCoreProperties creates a new CoreProperties instance.
func NewCoreProperties() *CoreProperties {
	return &CoreProperties{}
}

// Title returns the document title.
func (cp *CoreProperties) Title() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.title
}

// SetTitle sets the document title.
func (cp *CoreProperties) SetTitle(title string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.title = title
}

// Subject returns the document subject.
func (cp *CoreProperties) Subject() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.subject
}

// SetSubject sets the document subject.
func (cp *CoreProperties) SetSubject(subject string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.subject = subject
}

// Creator returns the document creator (author).
func (cp *CoreProperties) Creator() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.creator
}

// SetCreator sets the document creator (author).
func (cp *CoreProperties) SetCreator(creator string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.creator = creator
}

// Keywords returns the document keywords.
func (cp *CoreProperties) Keywords() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.keywords
}

// SetKeywords sets the document keywords.
func (cp *CoreProperties) SetKeywords(keywords string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.keywords = keywords
}

// Description returns the document description (comments).
func (cp *CoreProperties) Description() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.description
}

// SetDescription sets the document description (comments).
func (cp *CoreProperties) SetDescription(description string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.description = description
}

// Language returns the document language.
func (cp *CoreProperties) Language() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.language
}

// SetLanguage sets the document language.
func (cp *CoreProperties) SetLanguage(language string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.language = language
}

// LastModifiedBy returns the last modifier.
func (cp *CoreProperties) LastModifiedBy() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.lastModifiedBy
}

// SetLastModifiedBy sets the last modifier.
func (cp *CoreProperties) SetLastModifiedBy(name string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.lastModifiedBy = name
}

// Revision returns the document revision number.
func (cp *CoreProperties) Revision() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.revision
}

// SetRevision sets the document revision number.
func (cp *CoreProperties) SetRevision(revision string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.revision = revision
}

// Category returns the document category.
func (cp *CoreProperties) Category() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.category
}

// SetCategory sets the document category.
func (cp *CoreProperties) SetCategory(category string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.category = category
}

// ContentStatus returns the document content status.
func (cp *CoreProperties) ContentStatus() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.contentStatus
}

// SetContentStatus sets the document content status.
func (cp *CoreProperties) SetContentStatus(status string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.contentStatus = status
}

// Created returns the document creation date/time.
func (cp *CoreProperties) Created() *time.Time {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.created
}

// SetCreated sets the document creation date/time.
func (cp *CoreProperties) SetCreated(t time.Time) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.created = &t
}

// Modified returns the document modification date/time.
func (cp *CoreProperties) Modified() *time.Time {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.modified
}

// SetModified sets the document modification date/time.
func (cp *CoreProperties) SetModified(t time.Time) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.modified = &t
}

// LastPrinted returns the document last printed date/time.
func (cp *CoreProperties) LastPrinted() *time.Time {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.lastPrinted
}

// SetLastPrinted sets the document last printed date/time.
func (cp *CoreProperties) SetLastPrinted(t time.Time) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.lastPrinted = &t
}

// XML namespaces for core properties
const (
	nsCoreProperties = "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"
	nsDublinCore     = "http://purl.org/dc/elements/1.1/"
	nsDCTerms        = "http://purl.org/dc/terms/"
	nsXSI            = "http://www.w3.org/2001/XMLSchema-instance"
)

// XML types for core properties serialization
// Using fully qualified namespace URIs for proper XML namespace handling
type xmlCoreProperties struct {
	XMLName        xml.Name         `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Title          string           `xml:"http://purl.org/dc/elements/1.1/ title,omitempty"`
	Subject        string           `xml:"http://purl.org/dc/elements/1.1/ subject,omitempty"`
	Creator        string           `xml:"http://purl.org/dc/elements/1.1/ creator,omitempty"`
	Keywords       string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties keywords,omitempty"`
	Description    string           `xml:"http://purl.org/dc/elements/1.1/ description,omitempty"`
	Language       string           `xml:"http://purl.org/dc/elements/1.1/ language,omitempty"`
	LastModifiedBy string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties lastModifiedBy,omitempty"`
	Revision       string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties revision,omitempty"`
	Category       string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties category,omitempty"`
	ContentStatus  string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties contentStatus,omitempty"`
	Created        *xmlDCTermsDate  `xml:"http://purl.org/dc/terms/ created,omitempty"`
	Modified       *xmlDCTermsDate  `xml:"http://purl.org/dc/terms/ modified,omitempty"`
	LastPrinted    string           `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties lastPrinted,omitempty"`
}

type xmlDCTermsDate struct {
	Type  string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// MarshalToXML serializes the CoreProperties to XML.
func (cp *CoreProperties) MarshalToXML() ([]byte, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	xmlCP := xmlCoreProperties{
		Title:          cp.title,
		Subject:        cp.subject,
		Creator:        cp.creator,
		Keywords:       cp.keywords,
		Description:    cp.description,
		Language:       cp.language,
		LastModifiedBy: cp.lastModifiedBy,
		Revision:       cp.revision,
		Category:       cp.category,
		ContentStatus:  cp.contentStatus,
	}

	if cp.created != nil {
		xmlCP.Created = &xmlDCTermsDate{
			Type:  "dcterms:W3CDTF",
			Value: cp.created.Format(time.RFC3339),
		}
	}

	if cp.modified != nil {
		xmlCP.Modified = &xmlDCTermsDate{
			Type:  "dcterms:W3CDTF",
			Value: cp.modified.Format(time.RFC3339),
		}
	}

	if cp.lastPrinted != nil {
		xmlCP.LastPrinted = cp.lastPrinted.Format(time.RFC3339)
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)

	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlCP); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalFromXML deserializes the CoreProperties from XML.
func (cp *CoreProperties) UnmarshalFromXML(r io.Reader) error {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	var xmlCP xmlCoreProperties
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&xmlCP); err != nil {
		return err
	}

	cp.title = xmlCP.Title
	cp.subject = xmlCP.Subject
	cp.creator = xmlCP.Creator
	cp.keywords = xmlCP.Keywords
	cp.description = xmlCP.Description
	cp.language = xmlCP.Language
	cp.lastModifiedBy = xmlCP.LastModifiedBy
	cp.revision = xmlCP.Revision
	cp.category = xmlCP.Category
	cp.contentStatus = xmlCP.ContentStatus

	if xmlCP.Created != nil && xmlCP.Created.Value != "" {
		if t, err := time.Parse(time.RFC3339, xmlCP.Created.Value); err == nil {
			cp.created = &t
		}
	}

	if xmlCP.Modified != nil && xmlCP.Modified.Value != "" {
		if t, err := time.Parse(time.RFC3339, xmlCP.Modified.Value); err == nil {
			cp.modified = &t
		}
	}

	if xmlCP.LastPrinted != "" {
		if t, err := time.Parse(time.RFC3339, xmlCP.LastPrinted); err == nil {
			cp.lastPrinted = &t
		}
	}

	return nil
}

// Core properties content type and relationship type constants
const (
	CorePropertiesContentType      = "application/vnd.openxmlformats-package.core-properties+xml"
	CorePropertiesRelationshipType = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	CorePropertiesPartURI          = "/docProps/core.xml"
)
