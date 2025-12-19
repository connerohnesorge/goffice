package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// FootnoteType represents the type of a footnote.
type FootnoteType string

const (
	// FootnoteTypeNormal is a normal footnote.
	FootnoteTypeNormal FootnoteType = "normal"
	// FootnoteTypeSeparator is a footnote separator.
	FootnoteTypeSeparator FootnoteType = "separator"
	// FootnoteTypeContinuationSeparator is a continuation separator.
	FootnoteTypeContinuationSeparator FootnoteType = "continuationSeparator"
	// FootnoteTypeContinuationNotice is a continuation notice.
	FootnoteTypeContinuationNotice FootnoteType = "continuationNotice"
)

// Footnotes represents the root element for a footnotes part (w:footnotes).
type Footnotes struct {
	*openxml.PartRootElementBase
	nextID int
}

// NewFootnotes creates a new Footnotes element with default separators.
func NewFootnotes() *Footnotes {
	elem := openxml.NewPartRootElement(
		NamespaceWML,
		"footnotes",
		PrefixW,
	)
	fn := &Footnotes{
		PartRootElementBase: elem,
		nextID:              1,
	}

	// Add default separator footnotes
	fn.addSeparator(-1, FootnoteTypeSeparator)
	fn.addSeparator(
		0,
		FootnoteTypeContinuationSeparator,
	)

	return fn
}

// addSeparator adds a separator footnote with the given ID and type.
func (fn *Footnotes) addSeparator(
	id int,
	fnType FootnoteType,
) {
	footnote := newFootnoteWithType(id, fnType)

	// Add separator content
	p := NewParagraph()
	r := NewRun("")
	sep := openxml.NewCompositeElement(
		NamespaceWML,
		string(fnType),
		PrefixW,
	)
	r.AppendChild(sep)
	p.AppendChild(r)
	footnote.AppendChild(p)

	fn.AppendChild(footnote)
}

// Footnotes returns an iterator over all Footnote elements (excluding separators).
func (fn *Footnotes) Footnotes() iter.Seq[*Footnote] {
	return func(yield func(*Footnote) bool) {
		for child := range fn.Children() {
			if child.LocalName() == "footnote" &&
				child.NamespaceURI() == NamespaceWML {
				var f *Footnote
				if footnote, ok := child.(*Footnote); ok {
					f = footnote
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					f = &Footnote{CompositeElementBase: comp}
				}
				if f != nil {
					// Skip separator footnotes
					if f.Type() == FootnoteTypeNormal ||
						f.Type() == "" {
						if !yield(f) {
							return
						}
					}
				}
			}
		}
	}
}

// AllFootnotes returns an iterator over all Footnote elements (including separators).
func (fn *Footnotes) AllFootnotes() iter.Seq[*Footnote] {
	return func(yield func(*Footnote) bool) {
		for child := range fn.Children() {
			if child.LocalName() == "footnote" &&
				child.NamespaceURI() == NamespaceWML {
				var f *Footnote
				if footnote, ok := child.(*Footnote); ok {
					f = footnote
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					f = &Footnote{CompositeElementBase: comp}
				}
				if f != nil && !yield(f) {
					return
				}
			}
		}
	}
}

// GetFootnote returns the footnote with the specified ID, or nil if not found.
func (fn *Footnotes) GetFootnote(
	id int,
) *Footnote {
	for f := range fn.AllFootnotes() {
		if f.Id() == id {
			return f
		}
	}

	return nil
}

// AddFootnote adds a new footnote with the given text and returns it.
func (fn *Footnotes) AddFootnote(
	text string,
) *Footnote {
	f := NewFootnote(fn.nextID, text)
	fn.nextID++
	fn.AppendChild(f)

	return f
}

// NextID returns the next available footnote ID.
func (fn *Footnotes) NextID() int {
	return fn.nextID
}

// SetNextID sets the next footnote ID (used when loading existing documents).
func (fn *Footnotes) SetNextID(id int) {
	fn.nextID = id
}

// Clone creates a deep copy of this Footnotes element.
func (fn *Footnotes) Clone() openxml.Element {
	return &Footnotes{
		PartRootElementBase: fn.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
		nextID:              fn.nextID,
	}
}

// CloneNode creates a copy of this Footnotes element.
func (fn *Footnotes) CloneNode(
	deep bool,
) openxml.Element {
	return &Footnotes{
		PartRootElementBase: fn.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
		nextID:              fn.nextID,
	}
}

// Footnote represents a footnote element (w:footnote).
type Footnote struct {
	*openxml.CompositeElementBase
}

// NewFootnote creates a new Footnote element with the given ID and text.
func NewFootnote(id int, text string) *Footnote {
	f := newFootnoteWithType(
		id,
		FootnoteTypeNormal,
	)
	if text != "" {
		p := NewParagraph(text)
		f.AppendChild(p)
	}

	return f
}

// newFootnoteWithType creates a new Footnote with the specified ID and type.
func newFootnoteWithType(
	id int,
	fnType FootnoteType,
) *Footnote {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"footnote",
		PrefixW,
	)
	f := &Footnote{CompositeElementBase: elem}
	f.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
	if fnType != FootnoteTypeNormal &&
		fnType != "" {
		f.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"type",
				PrefixW,
				string(fnType),
			),
		)
	}

	return f
}

// Id returns the footnote ID.
func (f *Footnote) Id() int {
	attr, found := f.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())

	return id
}

// SetId sets the footnote ID.
func (f *Footnote) SetId(id int) {
	f.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Type returns the footnote type.
func (f *Footnote) Type() FootnoteType {
	attr, found := f.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return FootnoteTypeNormal
	}

	return FootnoteType(attr.Value())
}

// SetType sets the footnote type.
func (f *Footnote) SetType(fnType FootnoteType) {
	if fnType == FootnoteTypeNormal ||
		fnType == "" {
		f.RemoveAttribute("type", NamespaceWML)
	} else {
		f.SetAttribute(openxml.NewAttribute(NamespaceWML, "type", PrefixW, string(fnType)))
	}
}

// Paragraphs returns an iterator over all paragraphs in the footnote.
func (f *Footnote) Paragraphs() iter.Seq[*Paragraph] {
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

// AppendParagraph appends a paragraph with the given text.
func (f *Footnote) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	f.AppendChild(p)

	return p
}

// Clone creates a deep copy of this Footnote element.
func (f *Footnote) Clone() openxml.Element {
	return &Footnote{
		CompositeElementBase: f.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Footnote element.
func (f *Footnote) CloneNode(
	deep bool,
) openxml.Element {
	return &Footnote{
		CompositeElementBase: f.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// FootnoteReference represents a footnote reference element (w:footnoteReference).
// This is placed inline within a Run to reference a footnote.
type FootnoteReference struct {
	*openxml.CompositeElementBase
}

// NewFootnoteReference creates a new FootnoteReference element with the given ID.
func NewFootnoteReference(
	id int,
) *FootnoteReference {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"footnoteReference",
		PrefixW,
	)
	fr := &FootnoteReference{
		CompositeElementBase: elem,
	}
	fr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)

	return fr
}

// Id returns the referenced footnote ID.
func (fr *FootnoteReference) Id() int {
	attr, found := fr.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())

	return id
}

// SetId sets the referenced footnote ID.
func (fr *FootnoteReference) SetId(id int) {
	fr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy of this FootnoteReference element.
func (fr *FootnoteReference) Clone() openxml.Element {
	return &FootnoteReference{
		CompositeElementBase: fr.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this FootnoteReference element.
func (fr *FootnoteReference) CloneNode(
	deep bool,
) openxml.Element {
	return &FootnoteReference{
		CompositeElementBase: fr.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
