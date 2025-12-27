package openxml

import (
	"encoding/xml"
	"io"
	"strings"
)

// ElementFactory is a function that creates an element for a given
// qualified name. If the factory returns nil, a default element is created.
type ElementFactory func(namespaceURI, localName string) Element

// DefaultElementFactory creates default elements
// (CompositeElementBase or LeafElementBase).
func DefaultElementFactory(
	namespaceURI, localName string,
) Element {
	// By default, create composite elements that can hold children
	return NewCompositeElement(
		namespaceURI,
		localName,
		"",
	)
}

// ParseElement parses XML from a reader and constructs an element tree.
// The factory function is called to create elements for each XML element
// encountered. If factory is nil, DefaultElementFactory is used.
func ParseElement(
	r io.Reader,
	factory ElementFactory,
) (Element, error) {
	if factory == nil {
		factory = DefaultElementFactory //nolint:revive // modifies-parameter is intentional here
	}

	decoder := xml.NewDecoder(r)

	return parseElementFromDecoder(
		decoder,
		factory,
	)
}

// parseElementFromDecoder recursively parses elements from the decoder.
func parseElementFromDecoder(
	decoder *xml.Decoder,
	factory ElementFactory,
) (Element, error) {
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			return parseStartElement(decoder, t, factory)
		case xml.CharData:
			// Skip whitespace between elements at root level
			continue
		case xml.Comment, xml.ProcInst, xml.Directive:
			// Skip comments, processing instructions, directives at root level
			continue
		}
	}
}

// parseStartElement parses a start element and its contents.
func parseStartElement(
	decoder *xml.Decoder,
	start xml.StartElement,
	factory ElementFactory,
) (Element, error) {
	// Create the element using the factory
	elem := factory(
		start.Name.Space,
		start.Name.Local,
	)
	if elem == nil {
		elem = NewCompositeElement(
			start.Name.Space,
			start.Name.Local,
			"",
		)
	}

	// Extract prefix from the name if present in the token
	prefix := ""
	if start.Name.Space != "" {
		prefix = guessPrefixForNamespace(
			start.Name.Space,
		)
	}

	// Set prefix if element supports it and it's not the default namespace
	if base, ok := getBaseElement(elem); ok &&
		prefix != "" {
		base.prefix = prefix
	}

	// Set attributes
	for _, attr := range start.Attr {
		attrPrefix := extractPrefix(attr.Name)
		openxmlAttr := NewAttribute(
			attr.Name.Space,
			attr.Name.Local,
			attrPrefix,
			attr.Value,
		)
		elem.SetAttribute(openxmlAttr)
	}

	// Parse children
	if comp, ok := elem.(CompositeElement); ok {
		if err := parseChildren(decoder, comp, factory); err != nil {
			return nil, err
		}
	} else if leaf, ok := elem.(LeafElement); ok {
		// For leaf elements, read text content until end element
		text, err := readTextContent(decoder)
		if err != nil {
			return nil, err
		}
		leaf.SetInnerText(text)
	}

	return elem, nil
}

// parseChildren parses child elements and adds them to the parent.
func parseChildren(
	decoder *xml.Decoder,
	parent CompositeElement,
	factory ElementFactory,
) error {
	var textBuilder strings.Builder

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			child, err := parseStartElement(decoder, t, factory)
			if err != nil {
				return err
			}
			if child != nil {
				parent.AppendChild(child)
			}

		case xml.EndElement:
			// Set inner text for leaf elements with mixed content
			setLeafInnerText(parent, &textBuilder)

			return nil

		case xml.CharData:
			textBuilder.Write(t)

		case xml.Comment, xml.ProcInst, xml.Directive:
			// Skip these
		}
	}
}

// setLeafInnerText sets the inner text on a leaf element if applicable.
func setLeafInnerText(
	parent CompositeElement,
	textBuilder *strings.Builder,
) {
	if textBuilder.Len() == 0 {
		return
	}
	leaf, ok := parent.(LeafElement)
	if !ok {
		return
	}
	leaf.SetInnerText(
		strings.TrimSpace(textBuilder.String()),
	)
}

// readTextContent reads text content until the end element is reached.
func readTextContent(
	decoder *xml.Decoder,
) (string, error) {
	var textBuilder strings.Builder

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return textBuilder.String(), nil
		}
		if err != nil {
			return "", err
		}

		switch t := token.(type) {
		case xml.EndElement:
			return strings.TrimSpace(textBuilder.String()), nil
		case xml.CharData:
			textBuilder.Write(t)
		case xml.StartElement:
			// Unexpected child element in leaf - skip it
			if err := skipElement(decoder); err != nil {
				return "", err
			}
		}
	}
}

// skipElement skips the current element and all its children.
func skipElement(decoder *xml.Decoder) error {
	depth := 1
	for depth > 0 {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch token.(type) {
		case xml.StartElement:
			depth++
		case xml.EndElement:
			depth--
		}
	}

	return nil
}

// extractPrefix extracts the namespace prefix from an XML name.
// This is a heuristic since encoding/xml doesn't preserve prefixes.
func extractPrefix(name xml.Name) string {
	// encoding/xml doesn't preserve prefixes, so we have to guess
	// based on well-known namespace URIs
	return guessPrefixForNamespace(name.Space)
}

// guessPrefixForNamespace returns a common prefix for well-known namespaces.
func guessPrefixForNamespace(
	namespace string,
) string {
	switch namespace {
	case NamespaceWordprocessingML:
		return "w"
	case NamespaceRelationships:
		return "r"
	case NamespaceDrawingML:
		return "a"
	case NamespaceDrawingMLPicture:
		return "pic"
	case NamespaceDrawingMLWordprocessing:
		return "wp"
	case NamespaceContentTypes:
		return ""
	case NamespaceMarkupCompatibility:
		return "mc"
	case NamespaceOfficeDocument:
		return "o"
	case NamespaceVML:
		return "v"
	default:
		return ""
	}
}

// getBaseElement extracts the BaseElement from an element if possible.
func getBaseElement(
	elem Element,
) (*BaseElement, bool) {
	switch e := elem.(type) {
	case *CompositeElementBase:
		return &e.BaseElement, true
	case *LeafElementBase:
		return &e.BaseElement, true
	case *PartRootElementBase:
		return &e.BaseElement, true
	default:
		return nil, false
	}
}

// SetOuterXml parses XML and replaces the element's content.
// For composite elements, this replaces all children.
// For leaf elements, this sets the inner text.
func SetOuterXml(
	elem Element,
	xmlContent string,
) error {
	reader := strings.NewReader(xmlContent)
	parsed, err := ParseElement(reader, nil)
	if err != nil {
		return err
	}

	if parsed == nil {
		return nil
	}

	// Copy attributes
	if base, ok := getBaseElement(elem); ok {
		if parsedBase, ok := getBaseElement(parsed); ok {
			base.attributes = parsedBase.attributes
		}
	}

	// Copy children for composite elements
	if comp, ok := elem.(CompositeElement); ok {
		comp.RemoveAllChildren()
		if parsedComp, ok := parsed.(CompositeElement); ok {
			for child := range parsedComp.Children() {
				comp.AppendChild(child)
			}
		}
	}

	// Copy text for leaf elements
	if leaf, ok := elem.(LeafElement); ok {
		if parsedLeaf, ok := parsed.(LeafElement); ok {
			leaf.SetInnerText(
				parsedLeaf.InnerText(),
			)
		}
	}

	return nil
}
