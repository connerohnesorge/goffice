package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Footer represents the root element for a footer part (w:ftr).
// Footers can contain paragraphs, tables, and other block-level content.
type Footer struct {
	*openxml.PartRootElementBase
}

// NewFooter creates a new Footer element.
func NewFooter() *Footer {
	elem := openxml.NewPartRootElement(
		NamespaceWML,
		"ftr",
		PrefixW,
	)

	return &Footer{PartRootElementBase: elem}
}

// Paragraphs returns an iterator over all Paragraph elements in the footer.
func (f *Footer) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range f.Children() {
			if child.LocalName() == "p" &&
				child.NamespaceURI() == NamespaceWML {
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

// Tables returns an iterator over all Table elements in the footer.
func (f *Footer) Tables() iter.Seq[*Table] {
	return func(yield func(*Table) bool) {
		for child := range f.Children() {
			if child.LocalName() == "tbl" &&
				child.NamespaceURI() == NamespaceWML {
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

// AppendParagraph appends a new paragraph with the given text to the footer.
func (f *Footer) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	f.AppendChild(p)

	return p
}

// PrependParagraph prepends a new paragraph with the given text to the footer.
func (f *Footer) PrependParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	if first := f.FirstChild(); first != nil {
		f.InsertBefore(p, first)
	} else {
		f.AppendChild(p)
	}

	return p
}

// AppendTable appends a new table with the specified dimensions to the footer.
func (f *Footer) AppendTable(
	rows, cols int,
) *Table {
	t := NewTable(rows, cols)
	f.AppendChild(t)

	return t
}

// ClearContent removes all content from the footer.
func (f *Footer) ClearContent() {
	f.RemoveAllChildren()
}

// Clone creates a deep copy of this Footer element.
func (f *Footer) Clone() openxml.Element {
	return &Footer{
		PartRootElementBase: f.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Footer element.
func (f *Footer) CloneNode(
	deep bool,
) openxml.Element {
	return &Footer{
		PartRootElementBase: f.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
	}
}
