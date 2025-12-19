package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Document represents the root document element (w:document).
type Document struct {
	*openxml.CompositeElementBase
}

// NewDocument creates a new Document element with an empty body.
func NewDocument() *Document {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"document",
		PrefixW,
	)
	d := &Document{CompositeElementBase: elem}
	d.AppendChild(NewBody())

	return d
}

// Body returns the document body element.
func (d *Document) Body() *Body {
	elem := d.GetElement("body", NamespaceWML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Body); ok {
		return b
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Body{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateBody returns the body element, creating if needed.
func (d *Document) GetOrCreateBody() *Body {
	body := d.Body()
	if body != nil {
		return body
	}
	body = NewBody()
	d.AppendChild(body)

	return body
}

// AppendParagraph is a convenience method to append a paragraph to the body.
func (d *Document) AppendParagraph(
	text string,
) *Paragraph {
	return d.GetOrCreateBody().
		AppendParagraph(text)
}

// AppendTable is a convenience method to append a table to the body.
func (d *Document) AppendTable(
	rows, cols int,
) *Table {
	return d.GetOrCreateBody().
		AppendTable(rows, cols)
}

// Clone creates a deep copy of this Document element.
func (d *Document) Clone() openxml.Element {
	return &Document{
		CompositeElementBase: d.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Document element.
func (d *Document) CloneNode(
	deep bool,
) openxml.Element {
	return &Document{
		CompositeElementBase: d.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
