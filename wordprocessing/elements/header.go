package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Header represents the root element for a header part (w:hdr).
// Headers can contain paragraphs, tables, and other block-level content.
type Header struct {
	*openxml.PartRootElementBase
}

// NewHeader creates a new Header element.
func NewHeader() *Header {
	elem := openxml.NewPartRootElement(NamespaceWML, "hdr", PrefixW)
	return &Header{PartRootElementBase: elem}
}

// Paragraphs returns an iterator over all Paragraph elements in the header.
func (h *Header) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range h.Children() {
			if child.LocalName() == "p" && child.NamespaceURI() == NamespaceWML {
				var p *Paragraph
				if para, ok := child.(*Paragraph); ok {
					p = para
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					p = &Paragraph{CompositeElementBase: comp}
				}
				if p != nil && !yield(p) {
					return
				}
			}
		}
	}
}

// Tables returns an iterator over all Table elements in the header.
func (h *Header) Tables() iter.Seq[*Table] {
	return func(yield func(*Table) bool) {
		for child := range h.Children() {
			if child.LocalName() == "tbl" && child.NamespaceURI() == NamespaceWML {
				var t *Table
				if tbl, ok := child.(*Table); ok {
					t = tbl
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					t = &Table{CompositeElementBase: comp}
				}
				if t != nil && !yield(t) {
					return
				}
			}
		}
	}
}

// AppendParagraph appends a new paragraph with the given text to the header.
func (h *Header) AppendParagraph(text string) *Paragraph {
	p := NewParagraph(text)
	h.AppendChild(p)
	return p
}

// PrependParagraph prepends a new paragraph with the given text to the header.
func (h *Header) PrependParagraph(text string) *Paragraph {
	p := NewParagraph(text)
	if first := h.FirstChild(); first != nil {
		h.InsertBefore(p, first)
	} else {
		h.AppendChild(p)
	}
	return p
}

// AppendTable appends a new table with the specified dimensions to the header.
func (h *Header) AppendTable(rows, cols int) *Table {
	t := NewTable(rows, cols)
	h.AppendChild(t)
	return t
}

// ClearContent removes all content from the header.
func (h *Header) ClearContent() {
	h.RemoveAllChildren()
}

// Clone creates a deep copy of this Header element.
func (h *Header) Clone() openxml.Element {
	return &Header{
		PartRootElementBase: h.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Header element.
func (h *Header) CloneNode(deep bool) openxml.Element {
	return &Header{
		PartRootElementBase: h.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
	}
}
