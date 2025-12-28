//nolint:revive // file-length-limit - this file contains related endnote types
package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// EndnoteType represents the type of an endnote.
type EndnoteType string

const (
	// EndnoteTypeNormal is a normal endnote.
	EndnoteTypeNormal EndnoteType = "normal"
	// EndnoteTypeSeparator is an endnote separator.
	EndnoteTypeSeparator EndnoteType = "separator"
	// EndnoteTypeContinuationSeparator is a continuation separator.
	EndnoteTypeContinuationSeparator EndnoteType = "continuationSeparator"
	// EndnoteTypeContinuationNotice is a continuation notice.
	EndnoteTypeContinuationNotice EndnoteType = "continuationNotice"
)

// Endnotes represents the root element for an endnotes part (w:endnotes).
type Endnotes struct {
	*openxml.PartRootElementBase
	nextID int
}

// NewEndnotes creates a new Endnotes element with default separators.
func NewEndnotes() *Endnotes {
	elem := openxml.NewPartRootElement(
		NamespaceWML,
		"endnotes",
		PrefixW,
	)
	en := &Endnotes{
		PartRootElementBase: elem,
		nextID:              1,
	}

	// Add default separator endnotes
	en.addSeparator(-1, EndnoteTypeSeparator)
	en.addSeparator(
		0,
		EndnoteTypeContinuationSeparator,
	)

	return en
}

// addSeparator adds a separator endnote with the given ID and type.
func (en *Endnotes) addSeparator(
	id int,
	enType EndnoteType,
) {
	endnote := newEndnoteWithType(id, enType)

	// Add separator content
	p := NewParagraph()
	r := NewRun("")
	sep := openxml.NewCompositeElement(
		NamespaceWML,
		string(enType),
		PrefixW,
	)
	r.AppendChild(sep)
	p.AppendChild(r)
	endnote.AppendChild(p)

	en.AppendChild(endnote)
}

// Endnotes returns an iterator over all Endnote elements (excluding separators).
func (en *Endnotes) Endnotes() iter.Seq[*Endnote] {
	return func(yield func(*Endnote) bool) {
		for child := range en.Children() {
			if child.LocalName() == "endnote" &&
				child.NamespaceURI() == NamespaceWML {
				var e *Endnote
				switch v := child.(type) {
				case *Endnote:
					e = v
				case *openxml.CompositeElementBase:
					e = &Endnote{CompositeElementBase: v}
				}
				if e != nil {
					// Skip separator endnotes
					if e.Type() == EndnoteTypeNormal ||
						e.Type() == "" {
						if !yield(e) {
							return
						}
					}
				}
			}
		}
	}
}

// AllEndnotes returns an iterator over all Endnote elements (including separators).
func (en *Endnotes) AllEndnotes() iter.Seq[*Endnote] {
	return func(yield func(*Endnote) bool) {
		for child := range en.Children() {
			if child.LocalName() == "endnote" &&
				child.NamespaceURI() == NamespaceWML {
				var e *Endnote
				switch v := child.(type) {
				case *Endnote:
					e = v
				case *openxml.CompositeElementBase:
					e = &Endnote{CompositeElementBase: v}
				}
				if e != nil && !yield(e) {
					return
				}
			}
		}
	}
}

// GetEndnote returns the endnote with the specified ID, or nil if not found.
func (en *Endnotes) GetEndnote(id int) *Endnote {
	for e := range en.AllEndnotes() {
		if e.Id() == id {
			return e
		}
	}

	return nil
}

// AddEndnote adds a new endnote with the given text and returns it.
func (en *Endnotes) AddEndnote(
	text string,
) *Endnote {
	e := NewEndnote(en.nextID, text)
	en.nextID++
	en.AppendChild(e)

	return e
}

// NextID returns the next available endnote ID.
func (en *Endnotes) NextID() int {
	return en.nextID
}

// SetNextID sets the next endnote ID (used when loading existing documents).
func (en *Endnotes) SetNextID(id int) {
	en.nextID = id
}

// Clone creates a deep copy of this Endnotes element.
func (en *Endnotes) Clone() openxml.Element {
	return &Endnotes{
		PartRootElementBase: en.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
		nextID:              en.nextID,
	}
}

// CloneNode creates a copy of this Endnotes element.
func (en *Endnotes) CloneNode(
	deep bool,
) openxml.Element {
	return &Endnotes{
		PartRootElementBase: en.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
		nextID:              en.nextID,
	}
}

// Endnote represents an endnote element (w:endnote).
type Endnote struct {
	*openxml.CompositeElementBase
}

// NewEndnote creates a new Endnote element with the given ID and text.
func NewEndnote(id int, text string) *Endnote {
	e := newEndnoteWithType(id, EndnoteTypeNormal)
	if text != "" {
		p := NewParagraph(text)
		e.AppendChild(p)
	}

	return e
}

// newEndnoteWithType creates a new Endnote with the specified ID and type.
func newEndnoteWithType(
	id int,
	enType EndnoteType,
) *Endnote {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"endnote",
		PrefixW,
	)
	e := &Endnote{CompositeElementBase: elem}
	e.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
	if enType != EndnoteTypeNormal &&
		enType != "" {
		e.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"type",
				PrefixW,
				string(enType),
			),
		)
	}

	return e
}

// Id returns the endnote ID.
func (e *Endnote) Id() int {
	attr, found := e.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())

	return id
}

// SetId sets the endnote ID.
func (e *Endnote) SetId(id int) {
	e.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Type returns the endnote type.
func (e *Endnote) Type() EndnoteType {
	attr, found := e.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return EndnoteTypeNormal
	}

	return EndnoteType(attr.Value())
}

// SetType sets the endnote type.
func (e *Endnote) SetType(enType EndnoteType) {
	if enType == EndnoteTypeNormal ||
		enType == "" {
		e.RemoveAttribute("type", NamespaceWML)
	} else {
		e.SetAttribute(openxml.NewAttribute(NamespaceWML, "type", PrefixW, string(enType)))
	}
}

// Paragraphs returns an iterator over all paragraphs in the endnote.
func (e *Endnote) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range e.Children() {
			if child.LocalName() == "p" &&
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

// AppendParagraph appends a paragraph with the given text.
func (e *Endnote) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	e.AppendChild(p)

	return p
}

// Clone creates a deep copy of this Endnote element.
func (e *Endnote) Clone() openxml.Element {
	return &Endnote{
		CompositeElementBase: e.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Endnote element.
func (e *Endnote) CloneNode(
	deep bool,
) openxml.Element {
	return &Endnote{
		CompositeElementBase: e.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// EndnoteReference represents an endnote reference element (w:endnoteReference).
// This is placed inline within a Run to reference an endnote.
type EndnoteReference struct {
	*openxml.CompositeElementBase
}

// NewEndnoteReference creates a new EndnoteReference element with the given ID.
func NewEndnoteReference(
	id int,
) *EndnoteReference {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"endnoteReference",
		PrefixW,
	)
	er := &EndnoteReference{
		CompositeElementBase: elem,
	}
	er.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)

	return er
}

// Id returns the referenced endnote ID.
func (er *EndnoteReference) Id() int {
	attr, found := er.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())

	return id
}

// SetId sets the referenced endnote ID.
func (er *EndnoteReference) SetId(id int) {
	er.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy of this EndnoteReference element.
func (er *EndnoteReference) Clone() openxml.Element {
	return &EndnoteReference{
		CompositeElementBase: er.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this EndnoteReference element.
func (er *EndnoteReference) CloneNode(
	deep bool,
) openxml.Element {
	return &EndnoteReference{
		CompositeElementBase: er.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
