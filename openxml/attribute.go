package openxml

// OpenXmlAttribute represents an XML attribute on an element.
type OpenXmlAttribute struct {
	qname  OpenXmlQualifiedName
	prefix string
	value  string
}

// NewAttribute creates a new attribute with the given namespace URI, local name, and value.
func NewAttribute(namespaceURI, localName, prefix, value string) OpenXmlAttribute {
	return OpenXmlAttribute{
		qname:  NewQualifiedName(namespaceURI, localName),
		prefix: prefix,
		value:  value,
	}
}

// NewAttributeWithQName creates a new attribute with the given qualified name and value.
func NewAttributeWithQName(qname OpenXmlQualifiedName, prefix, value string) OpenXmlAttribute {
	return OpenXmlAttribute{
		qname:  qname,
		prefix: prefix,
		value:  value,
	}
}

// NewSimpleAttribute creates an attribute with no namespace.
func NewSimpleAttribute(localName, value string) OpenXmlAttribute {
	return OpenXmlAttribute{
		qname: NewQualifiedName("", localName),
		value: value,
	}
}

// QName returns the qualified name of the attribute.
func (a OpenXmlAttribute) QName() OpenXmlQualifiedName {
	return a.qname
}

// NamespaceURI returns the namespace URI of the attribute.
func (a OpenXmlAttribute) NamespaceURI() string {
	return a.qname.NamespaceURI()
}

// LocalName returns the local name of the attribute.
func (a OpenXmlAttribute) LocalName() string {
	return a.qname.LocalName()
}

// Prefix returns the namespace prefix of the attribute.
func (a OpenXmlAttribute) Prefix() string {
	return a.prefix
}

// Value returns the string value of the attribute.
func (a OpenXmlAttribute) Value() string {
	return a.value
}

// SetValue sets the value of the attribute.
func (a *OpenXmlAttribute) SetValue(value string) {
	a.value = value
}

// String returns a string representation of the attribute in the form "prefix:localName=value".
func (a OpenXmlAttribute) String() string {
	name := a.qname.LocalName()
	if a.prefix != "" {
		name = a.prefix + ":" + name
	}
	return name + "=\"" + a.value + "\""
}

// IsEmpty returns true if the attribute has no local name.
func (a OpenXmlAttribute) IsEmpty() bool {
	return a.qname.LocalName() == ""
}

// Equals returns true if this attribute has the same qualified name as the other.
// Note: This compares by name only, not value.
func (a OpenXmlAttribute) Equals(other OpenXmlAttribute) bool {
	return a.qname.Equals(other.qname)
}
