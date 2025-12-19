package openxml

import (
	"bytes"
	"io"
	"sort"
)

// XMLWriter provides namespace-aware XML serialization.
type XMLWriter struct {
	w                io.Writer
	namespaces       map[string]string // namespace URI -> prefix
	declaredPrefixes map[string]bool   // prefixes declared in current scope
	indent           string
	currentIndent    string
	pretty           bool
}

// NewXMLWriter creates a new XML writer that writes to w.
func NewXMLWriter(w io.Writer) *XMLWriter {
	return &XMLWriter{
		w:                w,
		namespaces:       make(map[string]string),
		declaredPrefixes: make(map[string]bool),
	}
}

// NewPrettyXMLWriter creates a new XML writer with indentation.
func NewPrettyXMLWriter(w io.Writer, indent string) *XMLWriter {
	return &XMLWriter{
		w:                w,
		namespaces:       make(map[string]string),
		declaredPrefixes: make(map[string]bool),
		indent:           indent,
		pretty:           true,
	}
}

// RegisterNamespace registers a prefix for a namespace URI.
func (xw *XMLWriter) RegisterNamespace(prefix, namespaceURI string) {
	xw.namespaces[namespaceURI] = prefix
}

// WriteDeclaration writes the XML declaration.
func (xw *XMLWriter) WriteDeclaration() error {
	_, err := io.WriteString(xw.w, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>")
	if xw.pretty {
		_, err = io.WriteString(xw.w, "\n")
	}
	return err
}

// WriteElement writes an element and its children.
func (xw *XMLWriter) WriteElement(elem Element) error {
	return elem.WriteXML(xw.w)
}

// WriteElementWithNamespaces writes an element with namespace declarations at the root.
func (xw *XMLWriter) WriteElementWithNamespaces(elem Element, namespaces map[string]string) error {
	var buf bytes.Buffer

	// Collect all namespaces used in the element tree
	usedNamespaces := collectNamespaces(elem, namespaces)

	// Write start tag with namespace declarations
	buf.WriteByte('<')

	prefix := elem.Prefix()
	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteByte(':')
	}
	buf.WriteString(elem.LocalName())

	// Write namespace declarations
	writeNamespaceDeclarations(&buf, usedNamespaces)

	// Write attributes
	for _, attr := range elem.Attributes() {
		buf.WriteByte(' ')
		if attr.Prefix() != "" {
			buf.WriteString(attr.Prefix())
			buf.WriteByte(':')
		}
		buf.WriteString(attr.LocalName())
		buf.WriteString("=\"")
		buf.WriteString(escapeXmlAttr(attr.Value()))
		buf.WriteByte('"')
	}

	// Check if element has children
	if comp, ok := elem.(CompositeElement); ok && comp.ChildCount() > 0 {
		buf.WriteByte('>')

		// Write children
		for child := range comp.Children() {
			if err := child.WriteXML(&buf); err != nil {
				return err
			}
		}

		// Write end tag
		buf.WriteString("</")
		if prefix != "" {
			buf.WriteString(prefix)
			buf.WriteByte(':')
		}
		buf.WriteString(elem.LocalName())
		buf.WriteByte('>')
	} else if leaf, ok := elem.(LeafElement); ok && leaf.InnerText() != "" {
		buf.WriteByte('>')
		buf.WriteString(escapeXmlText(leaf.InnerText()))
		buf.WriteString("</")
		if prefix != "" {
			buf.WriteString(prefix)
			buf.WriteByte(':')
		}
		buf.WriteString(elem.LocalName())
		buf.WriteByte('>')
	} else {
		buf.WriteString("/>")
	}

	_, err := xw.w.Write(buf.Bytes())
	return err
}

// collectNamespaces collects all namespaces used in an element tree.
func collectNamespaces(elem Element, prefixes map[string]string) map[string]string {
	result := make(map[string]string)

	// Add element's namespace
	if ns := elem.NamespaceURI(); ns != "" {
		if prefix, ok := prefixes[ns]; ok {
			result[ns] = prefix
		} else if elem.Prefix() != "" {
			result[ns] = elem.Prefix()
		}
	}

	// Add attribute namespaces
	for _, attr := range elem.Attributes() {
		if ns := attr.NamespaceURI(); ns != "" {
			if prefix, ok := prefixes[ns]; ok {
				result[ns] = prefix
			} else if attr.Prefix() != "" {
				result[ns] = attr.Prefix()
			}
		}
	}

	// Recurse into children
	if comp, ok := elem.(CompositeElement); ok {
		for child := range comp.Children() {
			childNs := collectNamespaces(child, prefixes)
			for ns, prefix := range childNs {
				result[ns] = prefix
			}
		}
	}

	return result
}

// writeNamespaceDeclarations writes xmlns declarations to the buffer.
func writeNamespaceDeclarations(buf *bytes.Buffer, namespaces map[string]string) {
	// Sort namespaces for deterministic output
	var uris []string
	for uri := range namespaces {
		uris = append(uris, uri)
	}
	sort.Strings(uris)

	for _, uri := range uris {
		prefix := namespaces[uri]
		buf.WriteByte(' ')
		if prefix != "" {
			buf.WriteString("xmlns:")
			buf.WriteString(prefix)
		} else {
			buf.WriteString("xmlns")
		}
		buf.WriteString("=\"")
		buf.WriteString(escapeXmlAttr(uri))
		buf.WriteByte('"')
	}
}

// WriteDocumentElement writes an element as a document root with XML declaration.
func WriteDocumentElement(w io.Writer, elem Element, namespaces map[string]string) error {
	// Write XML declaration
	if _, err := io.WriteString(w, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n"); err != nil {
		return err
	}

	// Write element with namespaces
	xw := NewXMLWriter(w)
	return xw.WriteElementWithNamespaces(elem, namespaces)
}

// OuterXmlWithNamespaces returns the XML representation with namespace declarations.
func OuterXmlWithNamespaces(elem Element, namespaces map[string]string) string {
	var buf bytes.Buffer
	xw := NewXMLWriter(&buf)
	xw.WriteElementWithNamespaces(elem, namespaces)
	return buf.String()
}
