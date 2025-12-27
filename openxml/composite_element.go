//nolint:revive // file-length-limit: composite element requires comprehensive implementation
package openxml

import (
	"bytes"
	"io"
	"iter"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// childNode is a node in a doubly-linked list of children.
type childNode struct {
	element Element
	prev    *childNode
	next    *childNode
}

// CompositeElementBase provides the implementation for composite elements.
// It uses a doubly-linked list for efficient child manipulation.
type CompositeElementBase struct {
	BaseElement
	firstChild *childNode
	lastChild  *childNode
	childCount int
}

// NewCompositeElement creates a new composite element with the given
// namespace URI and local name.
func NewCompositeElement(
	namespaceURI, localName, prefix string,
) *CompositeElementBase {
	elem := &CompositeElementBase{}
	InitBaseElement(
		&elem.BaseElement,
		namespaceURI,
		localName,
		prefix,
		nil,
	)

	return elem
}

// NewCompositeElementWithFeatures creates a new composite element
// with parent features.
func NewCompositeElementWithFeatures(
	namespaceURI, localName, prefix string,
	parentFeatures *features.FeatureCollection,
) *CompositeElementBase {
	elem := &CompositeElementBase{}
	InitBaseElement(
		&elem.BaseElement,
		namespaceURI,
		localName,
		prefix,
		parentFeatures,
	)

	return elem
}

// Children returns an iterator over all direct child elements.
func (c *CompositeElementBase) Children() iter.Seq[Element] {
	return func(yield func(Element) bool) {
		for node := c.firstChild; node != nil; node = node.next {
			if !yield(node.element) {
				return
			}
		}
	}
}

// FirstChild returns the first child element, or nil if empty.
func (c *CompositeElementBase) FirstChild() Element {
	if c.firstChild == nil {
		return nil
	}

	return c.firstChild.element
}

// LastChild returns the last child element, or nil if empty.
func (c *CompositeElementBase) LastChild() Element {
	if c.lastChild == nil {
		return nil
	}

	return c.lastChild.element
}

// GetElement returns the first child element with the given local name
// and namespace URI.
func (c *CompositeElementBase) GetElement(
	localName, namespaceURI string,
) Element {
	for node := c.firstChild; node != nil; node = node.next {
		if node.element.LocalName() == localName &&
			node.element.NamespaceURI() == namespaceURI {
			return node.element
		}
	}

	return nil
}

// findNode finds the node containing the given element.
func (c *CompositeElementBase) findNode(
	element Element,
) *childNode {
	for node := c.firstChild; node != nil; node = node.next {
		if node.element == element {
			return node
		}
	}

	return nil
}

// AppendChild adds a child element at the end of the children list.
func (c *CompositeElementBase) AppendChild(
	child Element,
) {
	if child == nil {
		return
	}

	// Remove from previous parent if any
	if parent := child.Parent(); parent != nil {
		if comp, ok := parent.(CompositeElement); ok {
			comp.RemoveChild(child)
		}
	}

	node := &childNode{element: child}
	child.setParent(c)

	if c.lastChild == nil {
		// Empty list
		c.firstChild = node
		c.lastChild = node
	} else {
		// Add to end
		node.prev = c.lastChild
		c.lastChild.next = node
		c.lastChild = node
	}
	c.childCount++
}

// PrependChild adds a child element at the beginning of the children list.
func (c *CompositeElementBase) PrependChild(
	child Element,
) {
	if child == nil {
		return
	}

	// Remove from previous parent if any
	if parent := child.Parent(); parent != nil {
		if comp, ok := parent.(CompositeElement); ok {
			comp.RemoveChild(child)
		}
	}

	node := &childNode{element: child}
	child.setParent(c)

	if c.firstChild == nil {
		// Empty list
		c.firstChild = node
		c.lastChild = node
	} else {
		// Add to beginning
		node.next = c.firstChild
		c.firstChild.prev = node
		c.firstChild = node
	}
	c.childCount++
}

// InsertBefore inserts newChild immediately before refChild.
// If refChild is nil, newChild is appended.
func (c *CompositeElementBase) InsertBefore(
	newChild, refChild Element,
) {
	if newChild == nil {
		return
	}

	if refChild == nil {
		c.AppendChild(newChild)

		return
	}

	refNode := c.findNode(refChild)
	if refNode == nil {
		// refChild not found, append
		c.AppendChild(newChild)

		return
	}

	// Remove from previous parent if any
	if parent := newChild.Parent(); parent != nil {
		if comp, ok := parent.(CompositeElement); ok {
			comp.RemoveChild(newChild)
		}
	}

	newNode := &childNode{element: newChild}
	newChild.setParent(c)

	newNode.next = refNode
	newNode.prev = refNode.prev

	if refNode.prev != nil {
		refNode.prev.next = newNode
	} else {
		c.firstChild = newNode
	}
	refNode.prev = newNode
	c.childCount++
}

// InsertAfter inserts newChild immediately after refChild.
// If refChild is nil, newChild is prepended.
func (c *CompositeElementBase) InsertAfter(
	newChild, refChild Element,
) {
	if newChild == nil {
		return
	}

	if refChild == nil {
		c.PrependChild(newChild)

		return
	}

	refNode := c.findNode(refChild)
	if refNode == nil {
		// refChild not found, append
		c.AppendChild(newChild)

		return
	}

	// Remove from previous parent if any
	if parent := newChild.Parent(); parent != nil {
		if comp, ok := parent.(CompositeElement); ok {
			comp.RemoveChild(newChild)
		}
	}

	newNode := &childNode{element: newChild}
	newChild.setParent(c)

	newNode.prev = refNode
	newNode.next = refNode.next

	if refNode.next != nil {
		refNode.next.prev = newNode
	} else {
		c.lastChild = newNode
	}
	refNode.next = newNode
	c.childCount++
}

// RemoveChild removes the specified child from this element.
func (c *CompositeElementBase) RemoveChild(
	child Element,
) bool {
	if child == nil {
		return false
	}

	node := c.findNode(child)
	if node == nil {
		return false
	}

	// Unlink the node
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.firstChild = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.lastChild = node.prev
	}

	child.setParent(nil)
	c.childCount--

	return true
}

// RemoveAllChildren removes all children from this element.
func (c *CompositeElementBase) RemoveAllChildren() {
	for node := c.firstChild; node != nil; node = node.next {
		node.element.setParent(nil)
	}
	c.firstChild = nil
	c.lastChild = nil
	c.childCount = 0
}

// ReplaceChild replaces oldChild with newChild.
func (c *CompositeElementBase) ReplaceChild(
	newChild, oldChild Element,
) bool {
	if newChild == nil || oldChild == nil {
		return false
	}

	oldNode := c.findNode(oldChild)
	if oldNode == nil {
		return false
	}

	// Remove newChild from its previous parent if any
	if parent := newChild.Parent(); parent != nil {
		if comp, ok := parent.(CompositeElement); ok {
			comp.RemoveChild(newChild)
		}
	}

	// Replace the element in the node
	oldChild.setParent(nil)
	oldNode.element = newChild
	newChild.setParent(c)

	return true
}

// ChildCount returns the number of direct children.
func (c *CompositeElementBase) ChildCount() int {
	return c.childCount
}

// NextSibling returns the next sibling element, or nil if this is the
// last child.
func (c *CompositeElementBase) NextSibling() Element {
	if c.parent == nil {
		return nil
	}

	if parentComp, ok := c.parent.(*CompositeElementBase); ok {
		for node := parentComp.firstChild; node != nil; node = node.next {
			if node.next != nil &&
				node.element == Element(c) {
				return node.next.element
			}
		}
	}

	return nil
}

// PreviousSibling returns the previous sibling element, or nil if this is
// the first child.
func (c *CompositeElementBase) PreviousSibling() Element {
	if c.parent == nil {
		return nil
	}

	if parentComp, ok := c.parent.(*CompositeElementBase); ok {
		for node := parentComp.firstChild; node != nil; node = node.next {
			if node.prev != nil &&
				node.element == Element(c) {
				return node.prev.element
			}
		}
	}

	return nil
}

// OuterXml returns the complete XML representation of this element.
func (c *CompositeElementBase) OuterXml() string {
	var buf bytes.Buffer
	_ = c.WriteXML(&buf)

	return buf.String()
}

// InnerXml returns the XML representation of this element's children only.
func (c *CompositeElementBase) InnerXml() string {
	var buf bytes.Buffer
	for node := c.firstChild; node != nil; node = node.next {
		_ = node.element.WriteXML(&buf)
	}

	return buf.String()
}

// WriteXML writes the XML representation to the given writer.
func (c *CompositeElementBase) WriteXML(
	w io.Writer,
) error {
	// If it's a root element (no parent), ensure it has the xmlns attribute
	// if it doesn't already have one and has a namespace URI.
	if c.parent == nil && c.NamespaceURI() != "" {
		found := false
		xmlnsLocal := "xmlns"
		if c.prefix != "" {
			xmlnsLocal = c.prefix
		}

		for _, attr := range c.attributes {
			if attr.LocalName() == xmlnsLocal &&
				(attr.Prefix() == "xmlns" || (c.prefix == "" && attr.Prefix() == "")) {
				found = true

				break
			}
		}
		if !found {
			if c.prefix == "" {
				c.SetAttribute(
					NewAttribute(
						"http://www.w3.org/2000/xmlns/",
						"xmlns",
						"",
						c.NamespaceURI(),
					),
				)
			} else {
				c.SetAttribute(NewAttribute("http://www.w3.org/2000/xmlns/", c.prefix, "xmlns", c.NamespaceURI()))
			}
		}
	}

	if c.childCount == 0 {
		return c.writeStartElement(w, true)
	}

	if err := c.writeStartElement(w, false); err != nil {
		return err
	}

	for node := c.firstChild; node != nil; node = node.next {
		if err := node.element.WriteXML(w); err != nil {
			return err
		}
	}

	return c.writeEndElement(w)
}

// Clone creates a deep copy of this element.
func (c *CompositeElementBase) Clone() Element {
	return c.CloneNode(true)
}

// CloneNode creates a copy of this element.
// If deep is true, children are also cloned.
//
//nolint:revive // flag-parameter: deep is standard clone API parameter
func (c *CompositeElementBase) CloneNode(
	deep bool,
) Element {
	clone := &CompositeElementBase{
		BaseElement: c.copyBaseElement(),
	}

	if deep {
		for node := c.firstChild; node != nil; node = node.next {
			childClone := node.element.Clone()
			clone.AppendChild(childClone)
		}
	}

	return clone
}

// Elements returns an iterator over child elements of a specific type.
func Elements[T Element](
	parent CompositeElement,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for child := range parent.Children() {
			typed, ok := child.(T)
			if !ok {
				continue
			}
			if !yield(typed) {
				return
			}
		}
	}
}

// GetElementTyped returns the first child element of type T.
func GetElementTyped[T Element](
	parent CompositeElement,
) (result T, found bool) {
	for child := range parent.Children() {
		if typed, ok := child.(T); ok {
			return typed, true
		}
	}

	return result, false
}
