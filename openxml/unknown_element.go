package openxml

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// UnknownElement preserves unknown XML elements during roundtrip.
// It stores raw XML to avoid data loss for unsupported extensions.
// This is crucial for forward compatibility - when opening documents
// that contain newer Office extensions that are not yet supported,
// the unknown elements are preserved byte-for-byte so they can be
// written back without loss of data.
type UnknownElement struct {
	localName    string
	namespaceURI string
	prefix       string
	rawXML       []byte
	parent       Element
	feats        *features.FeatureCollection
}

// NewUnknownElement creates a new UnknownElement with the given name and raw XML.
func NewUnknownElement(
	namespaceURI, localName, prefix string,
	rawXML []byte,
) *UnknownElement {
	return &UnknownElement{
		localName:    localName,
		namespaceURI: namespaceURI,
		prefix:       prefix,
		rawXML:       rawXML,
		feats:        features.NewFeatureCollection(),
	}
}

// LocalName returns the local name of the element.
func (u *UnknownElement) LocalName() string {
	return u.localName
}

// NamespaceURI returns the namespace URI of the element.
func (u *UnknownElement) NamespaceURI() string {
	return u.namespaceURI
}

// QName returns the fully qualified name of the element.
func (u *UnknownElement) QName() OpenXmlQualifiedName {
	return NewQualifiedName(
		u.namespaceURI,
		u.localName,
	)
}

// Prefix returns the namespace prefix of the element.
func (u *UnknownElement) Prefix() string {
	return u.prefix
}

// SetPrefix sets the namespace prefix (no-op for unknown elements).
func (u *UnknownElement) SetPrefix(
	prefix string,
) {
	u.prefix = prefix
}

// Parent returns the parent element, or nil if this is a root element.
func (u *UnknownElement) Parent() Element {
	return u.parent
}

// setParent sets the parent element.
func (u *UnknownElement) setParent(
	parent Element,
) {
	u.parent = parent
	// Update feature collection parent when element parent changes
	if parent != nil {
		u.feats.SetParent(parent.Features())
	} else {
		u.feats.SetParent(nil)
	}
}

// Features returns the feature collection for this element.
func (u *UnknownElement) Features() *features.FeatureCollection {
	return u.feats
}

// Attributes returns all attributes on this element.
// For unknown elements, we don't parse attributes separately,
// they are preserved in the raw XML.
func (*UnknownElement) Attributes() []OpenXmlAttribute {
	return make([]OpenXmlAttribute, 0)
}

// GetAttribute returns the attribute with the given local name and namespace URI.
// For unknown elements, this always returns false since we don't parse attributes.
//
//nolint:revive // unused-parameter: parameters required by Element interface
func (*UnknownElement) GetAttribute(
	_, _ string,
) (OpenXmlAttribute, bool) {
	return OpenXmlAttribute{}, false
}

// SetAttribute sets or adds an attribute on this element.
// For unknown elements, this is a no-op since we preserve raw XML.
//
//nolint:revive // unused-parameter: parameter required by Element interface
func (*UnknownElement) SetAttribute(
	_ OpenXmlAttribute,
) {
	// No-op for unknown elements
}

// RemoveAttribute removes the attribute with the given local name and namespace URI.
// For unknown elements, this always returns false.
//
//nolint:revive // unused-parameter: parameters required by Element interface
func (*UnknownElement) RemoveAttribute(
	_, _ string,
) bool {
	return false
}

// OuterXml returns the complete XML representation of this element.
func (u *UnknownElement) OuterXml() string {
	return string(u.rawXML)
}

// InnerXml returns the XML representation of this element's children only.
// For unknown elements, we return empty string since we don't parse the content.
func (*UnknownElement) InnerXml() string {
	return ""
}

// WriteXML writes the raw XML to the given writer.
// This preserves the unknown element exactly as it was read.
func (u *UnknownElement) WriteXML(
	w io.Writer,
) error {
	_, err := w.Write(u.rawXML)

	return err
}

// Clone creates a deep copy of this element.
func (u *UnknownElement) Clone() Element {
	return u.CloneNode(true)
}

// CloneNode creates a copy of this element.
// The deep parameter is ignored for unknown elements since they are opaque.
//
//nolint:revive // flag-parameter: deep is standard clone API parameter
func (u *UnknownElement) CloneNode(
	deep bool,
) Element {
	// Make a copy of the raw XML
	rawXMLCopy := make([]byte, len(u.rawXML))
	copy(rawXMLCopy, u.rawXML)

	return &UnknownElement{
		localName:    u.localName,
		namespaceURI: u.namespaceURI,
		prefix:       u.prefix,
		rawXML:       rawXMLCopy,
		parent:       nil, // Clone has no parent
		feats:        features.NewFeatureCollection(),
	}
}

// RawXML returns the raw XML bytes of this unknown element.
// This is useful for debugging or custom processing.
func (u *UnknownElement) RawXML() []byte {
	return u.rawXML
}
