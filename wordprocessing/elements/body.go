package elements

import (
	"iter"

	"github.com/connerohnesorge/goffice/openxml"
)

// Body represents the document body element (w:body).
type Body struct {
	*openxml.CompositeElementBase
}

// NewBody creates a new Body element.
func NewBody() *Body {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"body",
		PrefixW,
	)

	return &Body{CompositeElementBase: elem}
}

// Paragraphs returns an iterator over all Paragraph elements in the body.
func (b *Body) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range b.Children() {
			if child.LocalName() == "p" && //nolint:revive // early-return: iterator pattern
				child.NamespaceURI() == NamespaceWML {
				var p *Paragraph
				switch v := child.(type) {
				case *Paragraph:
					p = v
				case *openxml.CompositeElementBase:
					p = &Paragraph{CompositeElementBase: v}
				}
				if p != nil && !yield(p) {
					return
				}
			}
		}
	}
}

const tableLocalName = "tbl"

// Tables returns an iterator over all Table elements in the body.
func (b *Body) Tables() iter.Seq[*Table] {
	return func(yield func(*Table) bool) {
		for child := range b.Children() {
			// TODO: Re-enable PlaceholderValuesTbl when enum is regenerated
			if child.LocalName() != tableLocalName ||
				child.NamespaceURI() != NamespaceWML {
				continue
			}
			var t *Table
			switch v := child.(type) {
			case *Table:
				t = v
			case *openxml.CompositeElementBase:
				t = &Table{CompositeElementBase: v}
			}
			if t != nil && !yield(t) {
				return
			}
		}
	}
}

// SectionProperties returns the section properties for the document.
func (b *Body) SectionProperties() *SectionProperties {
	elem := b.GetElement("sectPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*SectionProperties); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SectionProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSectionProperties returns the section properties,
// creating if needed.
func (b *Body) GetOrCreateSectionProperties() *SectionProperties {
	sp := b.SectionProperties()
	if sp != nil {
		return sp
	}
	sp = NewSectionProperties()
	// Section properties should be the last child
	b.AppendChild(sp)

	return sp
}

// AppendParagraph appends a new paragraph with the given text.
func (b *Body) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	// Insert before section properties if present
	sectPr := b.SectionProperties()
	if sectPr != nil {
		b.InsertBefore(p, sectPr)
	} else {
		b.AppendChild(p)
	}

	return p
}

// PrependParagraph prepends a new paragraph with the given text.
func (b *Body) PrependParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	if first := b.FirstChild(); first != nil {
		b.InsertBefore(p, first)
	} else {
		b.AppendChild(p)
	}

	return p
}

// InsertParagraphBefore inserts a new paragraph before the specified element.
func (b *Body) InsertParagraphBefore(
	text string,
	ref openxml.Element,
) *Paragraph {
	p := NewParagraph(text)
	b.InsertBefore(p, ref)

	return p
}

// InsertParagraphAfter inserts a new paragraph after the specified element.
func (b *Body) InsertParagraphAfter(
	text string,
	ref openxml.Element,
) *Paragraph {
	p := NewParagraph(text)
	b.InsertAfter(p, ref)

	return p
}

// AppendTable appends a new table with the specified dimensions.
func (b *Body) AppendTable(
	rows, cols int,
) *Table {
	t := NewTable(rows, cols)
	// Insert before section properties if present
	sectPr := b.SectionProperties()
	if sectPr != nil {
		b.InsertBefore(t, sectPr)
	} else {
		b.AppendChild(t)
	}

	return t
}

// ClearContent removes all content from the body (except section properties).
func (b *Body) ClearContent() {
	var toRemove []openxml.Element
	for child := range b.Children() {
		if child.LocalName() != "sectPr" ||
			child.NamespaceURI() != NamespaceWML {
			toRemove = append(toRemove, child)
		}
	}
	for _, elem := range toRemove {
		b.RemoveChild(elem)
	}
}

// Clone creates a deep copy of this Body element.
func (b *Body) Clone() openxml.Element {
	cloned := b.CompositeElementBase.Clone()

	return &Body{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Body element.
func (b *Body) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.CompositeElementBase.CloneNode(
		deep,
	)

	return &Body{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Table represents a table element (w:tbl).
// This is a placeholder - full implementation would be in a separate file.
type Table struct {
	*openxml.CompositeElementBase
}

// NewTable creates a new Table element with the specified dimensions.
func NewTable(rows, cols int) *Table {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tbl",
		PrefixW,
	)
	t := &Table{CompositeElementBase: elem}

	// Add table properties
	tblPr := openxml.NewCompositeElement(
		NamespaceWML,
		"tblPr",
		PrefixW,
	)
	t.AppendChild(tblPr)

	// Add table grid
	tblGrid := openxml.NewCompositeElement(
		NamespaceWML,
		"tblGrid",
		PrefixW,
	)
	for range cols {
		gridCol := openxml.NewCompositeElement(
			NamespaceWML,
			"gridCol",
			PrefixW,
		)
		tblGrid.AppendChild(gridCol)
	}
	t.AppendChild(tblGrid)

	// Add rows
	for range rows {
		tr := NewTableRow(cols)
		t.AppendChild(tr)
	}

	return t
}

// Clone creates a deep copy of this Table element.
func (t *Table) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &Table{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TableRow represents a table row element (w:tr).
type TableRow struct {
	*openxml.CompositeElementBase
}

// NewTableRow creates a new TableRow element with the specified
// number of cells.
func NewTableRow(cols int) *TableRow {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tr",
		PrefixW,
	)
	tr := &TableRow{CompositeElementBase: elem}

	for range cols {
		tc := NewTableCell()
		tr.AppendChild(tc)
	}

	return tr
}

// Clone creates a deep copy of this TableRow element.
func (tr *TableRow) Clone() openxml.Element {
	cloned := tr.CompositeElementBase.Clone()

	return &TableRow{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// TableCell represents a table cell element (w:tc).
type TableCell struct {
	*openxml.CompositeElementBase
}

// NewTableCell creates a new TableCell element.
func NewTableCell() *TableCell {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tc",
		PrefixW,
	)
	tc := &TableCell{CompositeElementBase: elem}

	// Add an empty paragraph (required by OOXML)
	tc.AppendChild(NewParagraph())

	return tc
}

// Clone creates a deep copy of this TableCell element.
func (tc *TableCell) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &TableCell{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
