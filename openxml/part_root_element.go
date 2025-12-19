package openxml

import (
	"bytes"

	"github.com/connerohnesorge/goffice/openxml/features"
)

// PartRootElementBase provides the implementation for part root elements.
// A part root element is the root element of an OpenXML part and provides
// save/reload functionality.
type PartRootElementBase struct {
	CompositeElementBase
	part OpenXmlPart
}

// NewPartRootElement creates a new part root element with the given namespace URI and local name.
func NewPartRootElement(namespaceURI, localName, prefix string) *PartRootElementBase {
	elem := &PartRootElementBase{}
	InitBaseElement(&elem.BaseElement, namespaceURI, localName, prefix, nil)
	return elem
}

// NewPartRootElementWithFeatures creates a new part root element with parent features.
func NewPartRootElementWithFeatures(namespaceURI, localName, prefix string, parentFeatures *features.FeatureCollection) *PartRootElementBase {
	elem := &PartRootElementBase{}
	InitBaseElement(&elem.BaseElement, namespaceURI, localName, prefix, parentFeatures)
	return elem
}

// Part returns the part containing this root element.
func (p *PartRootElementBase) Part() OpenXmlPart {
	return p.part
}

// SetPart sets the containing part.
func (p *PartRootElementBase) SetPart(part OpenXmlPart) {
	p.part = part
}

// Save serializes this element back to the part.
func (p *PartRootElementBase) Save() error {
	if p.part == nil {
		return nil
	}

	var buf bytes.Buffer

	// Write XML declaration
	buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")

	// Write the element
	if err := p.WriteXML(&buf); err != nil {
		return err
	}

	p.part.SetData(buf.Bytes())
	return nil
}

// Reload re-parses the element from the part's stream.
// Note: This is a placeholder - actual implementation requires XML parsing.
func (p *PartRootElementBase) Reload() error {
	if p.part == nil {
		return nil
	}

	// Clear existing children
	p.RemoveAllChildren()
	p.ClearAttributes()

	// Reload would need to parse from p.part.GetStream()
	// This requires the XML reader implementation
	return nil
}

// Clone creates a deep copy of this element.
func (p *PartRootElementBase) Clone() Element {
	return p.CloneNode(true)
}

// CloneNode creates a copy of this element.
// If deep is true, children are also cloned.
// The clone is not associated with any part.
func (p *PartRootElementBase) CloneNode(deep bool) Element {
	clone := &PartRootElementBase{
		CompositeElementBase: CompositeElementBase{
			BaseElement: p.copyBaseElement(),
		},
		part: nil, // Clone is not associated with a part
	}

	if deep {
		for node := p.firstChild; node != nil; node = node.next {
			childClone := node.element.Clone()
			clone.AppendChild(childClone)
		}
	}

	return clone
}
