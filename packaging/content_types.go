// Package packaging provides the OPC (Open Packaging Conventions) layer.
package packaging

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
	"sync"
)

// ContentTypes manages content type mappings for an OPC package.
// It handles both default content types (by file extension) and
// override content types (for specific part URIs).
type ContentTypes struct {
	mu        sync.RWMutex
	defaults  map[string]string // extension (without dot) -> content type
	overrides map[string]string // normalized URI -> content type
}

// NewContentTypes creates a new ContentTypes instance with standard defaults.
func NewContentTypes() *ContentTypes {
	ct := &ContentTypes{
		defaults:  make(map[string]string),
		overrides: make(map[string]string),
	}

	// Register standard OPC defaults
	ct.defaults["xml"] = "application/xml"
	ct.defaults["rels"] = "application/vnd.openxmlformats-package.relationships+xml"

	return ct
}

// SetDefault registers a default content type for a file extension.
// The extension should not include the leading dot.
func (ct *ContentTypes) SetDefault(
	extension, contentType string,
) {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	ext := strings.ToLower(
		strings.TrimPrefix(extension, "."),
	)
	ct.defaults[ext] = contentType
}

// GetDefault returns the default content type for an extension.
// Returns empty string if no default is registered.
func (ct *ContentTypes) GetDefault(
	extension string,
) string {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	ext := strings.ToLower(
		strings.TrimPrefix(extension, "."),
	)
	return ct.defaults[ext]
}

// SetOverride registers an override content type for a specific part URI.
func (ct *ContentTypes) SetOverride(
	uri, contentType string,
) {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	normalizedURI := NormalizeURI(uri)
	ct.overrides[normalizedURI] = contentType
}

// RemoveOverride removes an override content type for a specific part URI.
func (ct *ContentTypes) RemoveOverride(
	uri string,
) {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	normalizedURI := NormalizeURI(uri)
	delete(ct.overrides, normalizedURI)
}

// GetContentType returns the content type for a part URI.
// It first checks for an override, then falls back to the default by extension.
// Returns empty string and ErrContentTypeNotFound if no content type is found.
func (ct *ContentTypes) GetContentType(
	uri string,
) (string, error) {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	normalizedURI := NormalizeURI(uri)

	// Check override first
	if contentType, ok := ct.overrides[normalizedURI]; ok {
		return contentType, nil
	}

	// Fall back to default by extension
	ext := strings.ToLower(
		strings.TrimPrefix(
			URIExtension(uri),
			".",
		),
	)
	if contentType, ok := ct.defaults[ext]; ok {
		return contentType, nil
	}

	return "", ErrContentTypeNotFound
}

// Defaults returns a copy of all default content type mappings.
func (ct *ContentTypes) Defaults() map[string]string {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	result := make(
		map[string]string,
		len(ct.defaults),
	)
	for k, v := range ct.defaults {
		result[k] = v
	}
	return result
}

// Overrides returns a copy of all override content type mappings.
func (ct *ContentTypes) Overrides() map[string]string {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	result := make(
		map[string]string,
		len(ct.overrides),
	)
	for k, v := range ct.overrides {
		result[k] = v
	}
	return result
}

// XML types for [Content_Types].xml serialization

type xmlContentTypes struct {
	XMLName   xml.Name      `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Defaults  []xmlDefault  `xml:"Default"`
	Overrides []xmlOverride `xml:"Override"`
}

type xmlDefault struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type xmlOverride struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

// MarshalToXML serializes the ContentTypes to XML.
func (ct *ContentTypes) MarshalToXML() ([]byte, error) {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	xmlCT := xmlContentTypes{}

	// Add defaults
	for ext, contentType := range ct.defaults {
		xmlCT.Defaults = append(
			xmlCT.Defaults,
			xmlDefault{
				Extension:   ext,
				ContentType: contentType,
			},
		)
	}

	// Add overrides
	for uri, contentType := range ct.overrides {
		xmlCT.Overrides = append(
			xmlCT.Overrides,
			xmlOverride{
				PartName:    uri,
				ContentType: contentType,
			},
		)
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)

	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlCT); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnmarshalFromXML deserializes the ContentTypes from XML.
func (ct *ContentTypes) UnmarshalFromXML(
	r io.Reader,
) error {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	var xmlCT xmlContentTypes
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&xmlCT); err != nil {
		return err
	}

	// Clear existing and repopulate
	ct.defaults = make(map[string]string)
	ct.overrides = make(map[string]string)

	for _, def := range xmlCT.Defaults {
		ct.defaults[strings.ToLower(def.Extension)] = def.ContentType
	}

	for _, ovr := range xmlCT.Overrides {
		ct.overrides[NormalizeURI(ovr.PartName)] = ovr.ContentType
	}

	return nil
}
