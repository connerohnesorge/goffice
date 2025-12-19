package openxml

import (
	"bytes"
	"io"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// LeafElementBase provides the implementation for leaf elements.
// Leaf elements contain only text content and cannot have child elements.
type LeafElementBase struct {
	BaseElement
	innerText string
}

// NewLeafElement creates a new leaf element with the given namespace URI and local name.
func NewLeafElement(namespaceURI, localName, prefix string) *LeafElementBase {
	elem := &LeafElementBase{}
	InitBaseElement(&elem.BaseElement, namespaceURI, localName, prefix, nil)
	return elem
}

// NewLeafElementWithFeatures creates a new leaf element with parent features.
func NewLeafElementWithFeatures(namespaceURI, localName, prefix string, parentFeatures *features.FeatureCollection) *LeafElementBase {
	elem := &LeafElementBase{}
	InitBaseElement(&elem.BaseElement, namespaceURI, localName, prefix, parentFeatures)
	return elem
}

// NewLeafElementWithText creates a new leaf element with initial text content.
func NewLeafElementWithText(namespaceURI, localName, prefix, text string) *LeafElementBase {
	elem := NewLeafElement(namespaceURI, localName, prefix)
	elem.innerText = text
	return elem
}

// InnerText returns the text content of this element.
func (l *LeafElementBase) InnerText() string {
	return l.innerText
}

// SetInnerText sets the text content of this element.
func (l *LeafElementBase) SetInnerText(value string) {
	l.innerText = value
}

// NextSibling returns the next sibling element, or nil if this is the last child.
func (l *LeafElementBase) NextSibling() Element {
	if l.parent == nil {
		return nil
	}

	if parentComp, ok := l.parent.(*CompositeElementBase); ok {
		for node := parentComp.firstChild; node != nil; node = node.next {
			if node.element == Element(l) && node.next != nil {
				return node.next.element
			}
		}
	}
	return nil
}

// PreviousSibling returns the previous sibling element, or nil if this is the first child.
func (l *LeafElementBase) PreviousSibling() Element {
	if l.parent == nil {
		return nil
	}

	if parentComp, ok := l.parent.(*CompositeElementBase); ok {
		for node := parentComp.firstChild; node != nil; node = node.next {
			if node.element == Element(l) && node.prev != nil {
				return node.prev.element
			}
		}
	}
	return nil
}

// OuterXml returns the complete XML representation of this element.
func (l *LeafElementBase) OuterXml() string {
	var buf bytes.Buffer
	_ = l.WriteXML(&buf)
	return buf.String()
}

// InnerXml returns the text content of this element.
// For leaf elements, InnerXml is the same as InnerText (escaped).
func (l *LeafElementBase) InnerXml() string {
	return escapeXmlText(l.innerText)
}

// WriteXML writes the XML representation to the given writer.
func (l *LeafElementBase) WriteXML(w io.Writer) error {
	if l.innerText == "" && len(l.attributes) == 0 {
		return l.writeStartElement(w, true)
	}

	if l.innerText == "" {
		return l.writeStartElement(w, true)
	}

	if err := l.writeStartElement(w, false); err != nil {
		return err
	}

	if _, err := io.WriteString(w, escapeXmlText(l.innerText)); err != nil {
		return err
	}

	return l.writeEndElement(w)
}

// Clone creates a deep copy of this element.
func (l *LeafElementBase) Clone() Element {
	return l.CloneNode(true)
}

// CloneNode creates a copy of this element.
// For leaf elements, deep has no effect since there are no children.
func (l *LeafElementBase) CloneNode(deep bool) Element {
	return &LeafElementBase{
		BaseElement: l.copyBaseElement(),
		innerText:   l.innerText,
	}
}

// IsEmpty returns true if the element has no text content and no attributes.
func (l *LeafElementBase) IsEmpty() bool {
	return l.innerText == "" && len(l.attributes) == 0
}
