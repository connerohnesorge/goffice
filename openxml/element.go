package openxml

import (
	"io"
	"iter"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// Element is the base interface for all OpenXML elements.
// It provides access to the element's identity, attributes, parent, and features.
type Element interface {
	// LocalName returns the local name of the element (without namespace prefix).
	LocalName() string

	// NamespaceURI returns the namespace URI of the element.
	NamespaceURI() string

	// QName returns the fully qualified name of the element.
	QName() OpenXmlQualifiedName

	// Prefix returns the namespace prefix of the element.
	Prefix() string

	// Parent returns the parent element, or nil if this is a root element.
	Parent() Element

	// Features returns the feature collection for this element.
	// Features provide access to shared services like namespace resolution.
	Features() *features.FeatureCollection

	// Attributes returns all attributes on this element.
	Attributes() []OpenXmlAttribute

	// GetAttribute returns the attribute with the given local name and namespace URI.
	// Returns the attribute and true if found, or an empty attribute and false if not.
	GetAttribute(localName, namespaceURI string) (OpenXmlAttribute, bool)

	// SetAttribute sets or adds an attribute on this element.
	SetAttribute(attr OpenXmlAttribute)

	// RemoveAttribute removes the attribute with the given local name and namespace URI.
	// Returns true if an attribute was removed.
	RemoveAttribute(localName, namespaceURI string) bool

	// OuterXml returns the complete XML representation of this element including children.
	OuterXml() string

	// InnerXml returns the XML representation of this element's children only.
	InnerXml() string

	// WriteXML writes the XML representation to the given writer.
	WriteXML(w io.Writer) error

	// Clone creates a deep copy of this element.
	Clone() Element

	// CloneNode creates a copy of this element.
	// If deep is true, children are also cloned.
	CloneNode(deep bool) Element

	// setParent sets the parent element (internal use).
	setParent(parent Element)
}

// CompositeElement is an element that can contain child elements.
type CompositeElement interface {
	Element

	// Children returns an iterator over all direct child elements.
	Children() iter.Seq[Element]

	// FirstChild returns the first child element, or nil if empty.
	FirstChild() Element

	// LastChild returns the last child element, or nil if empty.
	LastChild() Element

	// GetElement returns the first child element of type T, or nil if not found.
	GetElement(localName, namespaceURI string) Element

	// AppendChild adds a child element at the end of the children list.
	// The child's parent is set to this element.
	AppendChild(child Element)

	// PrependChild adds a child element at the beginning of the children list.
	// The child's parent is set to this element.
	PrependChild(child Element)

	// InsertBefore inserts newChild immediately before refChild.
	// If refChild is nil, newChild is appended.
	InsertBefore(newChild, refChild Element)

	// InsertAfter inserts newChild immediately after refChild.
	// If refChild is nil, newChild is prepended.
	InsertAfter(newChild, refChild Element)

	// RemoveChild removes the specified child from this element.
	// Returns true if the child was found and removed.
	RemoveChild(child Element) bool

	// RemoveAllChildren removes all children from this element.
	RemoveAllChildren()

	// ReplaceChild replaces oldChild with newChild.
	// Returns true if oldChild was found and replaced.
	ReplaceChild(newChild, oldChild Element) bool

	// ChildCount returns the number of direct children.
	ChildCount() int
}

// LeafElement is an element that contains only text content (no child elements).
type LeafElement interface {
	Element

	// InnerText returns the text content of this element.
	InnerText() string

	// SetInnerText sets the text content of this element.
	SetInnerText(value string)
}

// PartRootElement is the root element of a part.
// It provides access to the containing part and save/reload functionality.
type PartRootElement interface {
	CompositeElement

	// Part returns the part containing this root element.
	Part() OpenXmlPart

	// SetPart sets the containing part.
	SetPart(part OpenXmlPart)

	// Save serializes this element back to the part.
	Save() error

	// Reload re-parses the element from the part's stream.
	Reload() error
}

// OpenXmlPart is an interface representing a part in an OpenXML package.
// This is defined here to avoid circular imports with the packaging package.
type OpenXmlPart interface {
	// URI returns the URI of this part.
	URI() string

	// ContentType returns the content type of this part.
	ContentType() string

	// GetStream returns a reader for the part's content.
	GetStream() io.Reader

	// SetData sets the part's raw content.
	SetData(data []byte)
}

// ElementWithSiblings provides sibling navigation.
type ElementWithSiblings interface {
	Element

	// NextSibling returns the next sibling element, or nil if this is the last child.
	NextSibling() Element

	// PreviousSibling returns the previous sibling element, or nil if this is the first child.
	PreviousSibling() Element
}
