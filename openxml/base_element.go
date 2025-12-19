package openxml

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// BaseElement provides the common implementation for all element types.
// It is embedded by CompositeElementBase and LeafElementBase.
type BaseElement struct {
	qname      OpenXmlQualifiedName
	prefix     string
	parent     Element
	attributes []OpenXmlAttribute
	features   *features.FeatureCollection
}

// InitBaseElement initializes a BaseElement with the given namespace URI, local name, and prefix.
func InitBaseElement(
	b *BaseElement,
	namespaceURI, localName, prefix string,
	parentFeatures *features.FeatureCollection,
) {
	b.qname = NewQualifiedName(
		namespaceURI,
		localName,
	)
	b.prefix = prefix
	b.attributes = nil
	if parentFeatures != nil {
		b.features = features.NewFeatureCollectionWithParent(
			parentFeatures,
		)
	} else {
		b.features = features.NewFeatureCollection()
	}
}

// LocalName returns the local name of the element.
func (b *BaseElement) LocalName() string {
	return b.qname.LocalName()
}

// NamespaceURI returns the namespace URI of the element.
func (b *BaseElement) NamespaceURI() string {
	return b.qname.NamespaceURI()
}

// QName returns the fully qualified name of the element.
func (b *BaseElement) QName() OpenXmlQualifiedName {
	return b.qname
}

// Prefix returns the namespace prefix of the element.
func (b *BaseElement) Prefix() string {
	return b.prefix
}

// SetPrefix sets the namespace prefix.
func (b *BaseElement) SetPrefix(prefix string) {
	b.prefix = prefix
}

// Parent returns the parent element, or nil if this is a root element.
func (b *BaseElement) Parent() Element {
	return b.parent
}

// setParent sets the parent element.
func (b *BaseElement) setParent(parent Element) {
	b.parent = parent
	// Update feature collection parent when element parent changes
	if parent != nil {
		b.features.SetParent(parent.Features())
	} else {
		b.features.SetParent(nil)
	}
}

// Features returns the feature collection for this element.
func (b *BaseElement) Features() *features.FeatureCollection {
	return b.features
}

// Attributes returns all attributes on this element.
func (b *BaseElement) Attributes() []OpenXmlAttribute {
	if b.attributes == nil {
		return []OpenXmlAttribute{}
	}
	// Return a copy to prevent external modification
	result := make(
		[]OpenXmlAttribute,
		len(b.attributes),
	)
	copy(result, b.attributes)

	return result
}

// GetAttribute returns the attribute with the given local name and namespace URI.
func (b *BaseElement) GetAttribute(
	localName, namespaceURI string,
) (OpenXmlAttribute, bool) {
	for _, attr := range b.attributes {
		if attr.LocalName() == localName &&
			attr.NamespaceURI() == namespaceURI {
			return attr, true
		}
	}

	return OpenXmlAttribute{}, false
}

// SetAttribute sets or adds an attribute on this element.
func (b *BaseElement) SetAttribute(
	attr OpenXmlAttribute,
) {
	// Look for existing attribute with same name
	for i, existing := range b.attributes {
		if existing.LocalName() == attr.LocalName() &&
			existing.NamespaceURI() == attr.NamespaceURI() {
			b.attributes[i] = attr

			return
		}
	}
	// Add new attribute
	b.attributes = append(b.attributes, attr)
}

// RemoveAttribute removes the attribute with the given local name and namespace URI.
func (b *BaseElement) RemoveAttribute(
	localName, namespaceURI string,
) bool {
	for i, attr := range b.attributes {
		if attr.LocalName() == localName &&
			attr.NamespaceURI() == namespaceURI {
			b.attributes = append(
				b.attributes[:i],
				b.attributes[i+1:]...)

			return true
		}
	}

	return false
}

// ClearAttributes removes all attributes from this element.
func (b *BaseElement) ClearAttributes() {
	b.attributes = nil
}

// AttributeCount returns the number of attributes on this element.
func (b *BaseElement) AttributeCount() int {
	return len(b.attributes)
}

// xmlStartElement returns an xml.StartElement for this element.
func (b *BaseElement) xmlStartElement() xml.StartElement {
	name := xml.Name{
		Space: b.qname.NamespaceURI(),
		Local: b.qname.LocalName(),
	}

	attrs := make([]xml.Attr, len(b.attributes))
	for i, attr := range b.attributes {
		attrs[i] = xml.Attr{
			Name: xml.Name{
				Space: attr.NamespaceURI(),
				Local: attr.LocalName(),
			},
			Value: attr.Value(),
		}
	}

	return xml.StartElement{
		Name: name,
		Attr: attrs,
	}
}

// writeStartElement writes the opening tag to the writer.
func (b *BaseElement) writeStartElement(
	w io.Writer,
	selfClose bool,
) error {
	var buf bytes.Buffer
	buf.WriteByte('<')

	if b.prefix != "" {
		buf.WriteString(b.prefix)
		buf.WriteByte(':')
	}
	buf.WriteString(b.qname.LocalName())

	// Write namespace declaration if this is a root element or has a different namespace
	if b.prefix != "" &&
		b.qname.NamespaceURI() != "" {
		buf.WriteString(" xmlns:")
		buf.WriteString(b.prefix)
		buf.WriteString("=\"")
		buf.WriteString(
			escapeXmlAttr(b.qname.NamespaceURI()),
		)
		buf.WriteByte('"')
	}

	// Write attributes
	for _, attr := range b.attributes {
		buf.WriteByte(' ')
		if attr.Prefix() != "" {
			buf.WriteString(attr.Prefix())
			buf.WriteByte(':')
		}
		buf.WriteString(attr.LocalName())
		buf.WriteString("=\"")
		buf.WriteString(
			escapeXmlAttr(attr.Value()),
		)
		buf.WriteByte('"')
	}

	if selfClose {
		buf.WriteString("/>")
	} else {
		buf.WriteByte('>')
	}

	_, err := w.Write(buf.Bytes())

	return err
}

// writeEndElement writes the closing tag to the writer.
func (b *BaseElement) writeEndElement(
	w io.Writer,
) error {
	var buf bytes.Buffer
	buf.WriteString("</")
	if b.prefix != "" {
		buf.WriteString(b.prefix)
		buf.WriteByte(':')
	}
	buf.WriteString(b.qname.LocalName())
	buf.WriteByte('>')
	_, err := w.Write(buf.Bytes())

	return err
}

// escapeXmlAttr escapes special characters in an XML attribute value.
func escapeXmlAttr(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		switch r {
		case '<':
			buf.WriteString("&lt;")
		case '>':
			buf.WriteString("&gt;")
		case '&':
			buf.WriteString("&amp;")
		case '"':
			buf.WriteString("&quot;")
		case '\'':
			buf.WriteString("&apos;")
		default:
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

// escapeXmlText escapes special characters in XML text content.
func escapeXmlText(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		switch r {
		case '<':
			buf.WriteString("&lt;")
		case '>':
			buf.WriteString("&gt;")
		case '&':
			buf.WriteString("&amp;")
		default:
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

// copyBaseElement copies the base element fields.
func (b *BaseElement) copyBaseElement() BaseElement {
	attrs := make(
		[]OpenXmlAttribute,
		len(b.attributes),
	)
	copy(attrs, b.attributes)

	return BaseElement{
		qname:      b.qname,
		prefix:     b.prefix,
		parent:     nil, // Clone has no parent
		attributes: attrs,
		features:   features.NewFeatureCollection(),
	}
}
