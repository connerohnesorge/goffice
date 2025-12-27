package openxml

import "strings"

// OpenXmlQualifiedName represents an XML qualified name
// (namespace URI + local name). This is used to uniquely identify
// element and attribute names in the OpenXML schema.
type OpenXmlQualifiedName struct {
	namespaceURI string
	localName    string
}

// NewQualifiedName creates a new qualified name with the given
// namespace URI and local name.
func NewQualifiedName(
	namespaceURI, localName string,
) OpenXmlQualifiedName {
	return OpenXmlQualifiedName{
		namespaceURI: namespaceURI,
		localName:    localName,
	}
}

// NamespaceURI returns the namespace URI component of the qualified name.
func (qn OpenXmlQualifiedName) NamespaceURI() string {
	return qn.namespaceURI
}

// LocalName returns the local name component of the qualified name.
func (qn OpenXmlQualifiedName) LocalName() string {
	return qn.localName
}

// String returns the string representation of the qualified name.
// Format: "{namespaceURI}localName" if namespace is present,
// otherwise just "localName".
func (qn OpenXmlQualifiedName) String() string {
	if qn.namespaceURI == "" {
		return qn.localName
	}

	return "{" + qn.namespaceURI + "}" + qn.localName
}

// IsEmpty returns true if both the namespace URI and local name are empty.
func (qn OpenXmlQualifiedName) IsEmpty() bool {
	return qn.namespaceURI == "" &&
		qn.localName == ""
}

// Equals returns true if this qualified name equals the other.
func (qn OpenXmlQualifiedName) Equals(
	other OpenXmlQualifiedName,
) bool {
	return qn.namespaceURI == other.namespaceURI &&
		qn.localName == other.localName
}

// ParseQualifiedName parses a qualified name from string format.
// Accepts either "{namespaceURI}localName" or just "localName".
func ParseQualifiedName(
	s string,
) OpenXmlQualifiedName {
	if !strings.HasPrefix(s, "{") {
		return OpenXmlQualifiedName{localName: s}
	}

	idx := strings.Index(s, "}")
	if idx == -1 {
		return OpenXmlQualifiedName{localName: s}
	}

	return OpenXmlQualifiedName{
		namespaceURI: s[1:idx],
		localName:    s[idx+1:],
	}
}
