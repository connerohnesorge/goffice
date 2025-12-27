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

// NewPartRootElement creates a new part root element with
// the given namespace URI and local name.
func NewPartRootElement(
	namespaceURI, localName, prefix string,
) *PartRootElementBase {
	elem := &PartRootElementBase{}
	InitBaseElement(
		&elem.BaseElement,
		namespaceURI,
		localName,
		prefix,
		nil,
	)

	return elem
}

// NewPartRootElementWithFeatures creates a new part root element
// with parent features.
func NewPartRootElementWithFeatures(
	namespaceURI, localName, prefix string,
	parentFeatures *features.FeatureCollection,
) *PartRootElementBase {
	elem := &PartRootElementBase{}
	InitBaseElement(
		&elem.BaseElement,
		namespaceURI,
		localName,
		prefix,
		parentFeatures,
	)

	return elem
}

// Part returns the part containing this root element.
func (p *PartRootElementBase) Part() OpenXmlPart {
	return p.part
}

// SetPart sets the containing part.
func (p *PartRootElementBase) SetPart(
	part OpenXmlPart,
) {
	p.part = part
}

// Save serializes this element back to the part.
func (p *PartRootElementBase) Save() error {
	if p.part == nil {
		return nil
	}

	var buf bytes.Buffer

	// Write XML declaration
	buf.WriteString(
		"<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n",
	)

	// Write the element
	if err := p.WriteXML(&buf); err != nil {
		return err
	}

	p.part.SetData(buf.Bytes())

	return nil
}

// Reload re-parses the element from the part's stream.
func (p *PartRootElementBase) Reload() error {
	if p.part == nil {
		return nil
	}

	stream := p.part.GetStream()
	if stream == nil {
		return nil
	}

	// Create a factory that handles the root element properly
	isRoot := true
	factory := func(ns, local string) Element {
		if isRoot {
			isRoot = false
			// Reuse this element if it matches
			if p.LocalName() == local &&
				p.NamespaceURI() == ns {
				p.RemoveAllChildren()
				p.ClearAttributes()

				return p
			}
			// Otherwise create new PartRootElement
			return NewPartRootElement(
				ns,
				local,
				"",
			)
		}
		// Create normal elements for children
		return NewCompositeElement(ns, local, "")
	}

	_, err := ParseElement(stream, factory)

	return err
}

// Clone creates a deep copy of this element.
func (p *PartRootElementBase) Clone() Element {
	return p.CloneNode(true)
}

// CloneNode creates a copy of this element.
// If deep is true, children are also cloned.
// The clone is not associated with any part.
//
//nolint:revive // deep is a standard clone parameter
func (p *PartRootElementBase) CloneNode(
	deep bool,
) Element {
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
