package elements

import (
	"iter"
	"strconv"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
)

// Comments represents the root element for a comments part (w:comments).
type Comments struct {
	*openxml.PartRootElementBase
	nextID int
}

// NewComments creates a new Comments element.
func NewComments() *Comments {
	elem := openxml.NewPartRootElement(
		NamespaceWML,
		"comments",
		PrefixW,
	)
	return &Comments{
		PartRootElementBase: elem,
		nextID:              1,
	}
}

// Comments returns an iterator over all Comment elements.
func (c *Comments) Comments() iter.Seq[*Comment] {
	return func(yield func(*Comment) bool) {
		for child := range c.Children() {
			if child.LocalName() == "comment" &&
				child.NamespaceURI() == NamespaceWML {
				var comment *Comment
				if cm, ok := child.(*Comment); ok {
					comment = cm
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					comment = &Comment{CompositeElementBase: comp}
				}
				if comment != nil &&
					!yield(comment) {
					return
				}
			}
		}
	}
}

// GetComment returns the comment with the specified ID, or nil if not found.
func (c *Comments) GetComment(id int) *Comment {
	for comment := range c.Comments() {
		if comment.Id() == id {
			return comment
		}
	}
	return nil
}

// AddComment adds a new comment with the given author and text.
func (c *Comments) AddComment(
	author, text string,
) *Comment {
	comment := NewComment(c.nextID, author, text)
	c.nextID++
	c.AppendChild(comment)
	return comment
}

// NextID returns the next available comment ID.
func (c *Comments) NextID() int {
	return c.nextID
}

// SetNextID sets the next comment ID (used when loading existing documents).
func (c *Comments) SetNextID(id int) {
	c.nextID = id
}

// Clone creates a deep copy of this Comments element.
func (c *Comments) Clone() openxml.Element {
	return &Comments{
		PartRootElementBase: c.PartRootElementBase.Clone().(*openxml.PartRootElementBase),
		nextID:              c.nextID,
	}
}

// CloneNode creates a copy of this Comments element.
func (c *Comments) CloneNode(
	deep bool,
) openxml.Element {
	return &Comments{
		PartRootElementBase: c.PartRootElementBase.CloneNode(deep).(*openxml.PartRootElementBase),
		nextID:              c.nextID,
	}
}

// Comment represents a comment element (w:comment).
type Comment struct {
	*openxml.CompositeElementBase
}

// NewComment creates a new Comment element with the given ID, author, and text.
func NewComment(
	id int,
	author, text string,
) *Comment {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"comment",
		PrefixW,
	)
	c := &Comment{CompositeElementBase: elem}
	c.SetId(id)
	c.SetAuthor(author)
	c.SetDate(time.Now())

	if text != "" {
		p := NewParagraph(text)
		c.AppendChild(p)
	}
	return c
}

// Id returns the comment ID.
func (c *Comment) Id() int {
	attr, found := c.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())
	return id
}

// SetId sets the comment ID.
func (c *Comment) SetId(id int) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Author returns the comment author name.
func (c *Comment) Author() string {
	attr, found := c.GetAttribute(
		"author",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetAuthor sets the comment author name.
func (c *Comment) SetAuthor(name string) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"author",
			PrefixW,
			name,
		),
	)
}

// Date returns the comment date.
func (c *Comment) Date() time.Time {
	attr, found := c.GetAttribute(
		"date",
		NamespaceWML,
	)
	if !found {
		return time.Time{}
	}
	t, err := time.Parse(
		time.RFC3339,
		attr.Value(),
	)
	if err != nil {
		return time.Time{}
	}
	return t
}

// SetDate sets the comment date.
func (c *Comment) SetDate(t time.Time) {
	c.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"date",
			PrefixW,
			t.Format(time.RFC3339),
		),
	)
}

// Initials returns the comment author's initials.
func (c *Comment) Initials() string {
	attr, found := c.GetAttribute(
		"initials",
		NamespaceWML,
	)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetInitials sets the comment author's initials.
func (c *Comment) SetInitials(initials string) {
	if initials == "" {
		c.RemoveAttribute(
			"initials",
			NamespaceWML,
		)
	} else {
		c.SetAttribute(openxml.NewAttribute(NamespaceWML, "initials", PrefixW, initials))
	}
}

// Paragraphs returns an iterator over all paragraphs in the comment.
func (c *Comment) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range c.Children() {
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
func (c *Comment) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	c.AppendChild(p)
	return p
}

// Clone creates a deep copy of this Comment element.
func (c *Comment) Clone() openxml.Element {
	return &Comment{
		CompositeElementBase: c.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Comment element.
func (c *Comment) CloneNode(
	deep bool,
) openxml.Element {
	return &Comment{
		CompositeElementBase: c.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// CommentRangeStart represents the start of a comment range (w:commentRangeStart).
// This marks where a comment's scope begins in the document.
type CommentRangeStart struct {
	*openxml.CompositeElementBase
}

// NewCommentRangeStart creates a new CommentRangeStart element with the given comment ID.
func NewCommentRangeStart(
	id int,
) *CommentRangeStart {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"commentRangeStart",
		PrefixW,
	)
	crs := &CommentRangeStart{
		CompositeElementBase: elem,
	}
	crs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
	return crs
}

// Id returns the associated comment ID.
func (crs *CommentRangeStart) Id() int {
	attr, found := crs.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())
	return id
}

// SetId sets the associated comment ID.
func (crs *CommentRangeStart) SetId(id int) {
	crs.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy of this CommentRangeStart element.
func (crs *CommentRangeStart) Clone() openxml.Element {
	return &CommentRangeStart{
		CompositeElementBase: crs.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CommentRangeStart element.
func (crs *CommentRangeStart) CloneNode(
	deep bool,
) openxml.Element {
	return &CommentRangeStart{
		CompositeElementBase: crs.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// CommentRangeEnd represents the end of a comment range (w:commentRangeEnd).
// This marks where a comment's scope ends in the document.
type CommentRangeEnd struct {
	*openxml.CompositeElementBase
}

// NewCommentRangeEnd creates a new CommentRangeEnd element with the given comment ID.
func NewCommentRangeEnd(id int) *CommentRangeEnd {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"commentRangeEnd",
		PrefixW,
	)
	cre := &CommentRangeEnd{
		CompositeElementBase: elem,
	}
	cre.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
	return cre
}

// Id returns the associated comment ID.
func (cre *CommentRangeEnd) Id() int {
	attr, found := cre.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())
	return id
}

// SetId sets the associated comment ID.
func (cre *CommentRangeEnd) SetId(id int) {
	cre.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy of this CommentRangeEnd element.
func (cre *CommentRangeEnd) Clone() openxml.Element {
	return &CommentRangeEnd{
		CompositeElementBase: cre.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CommentRangeEnd element.
func (cre *CommentRangeEnd) CloneNode(
	deep bool,
) openxml.Element {
	return &CommentRangeEnd{
		CompositeElementBase: cre.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// CommentReference represents a comment reference element (w:commentReference).
// This is placed inline within a Run to mark the location of a comment anchor.
type CommentReference struct {
	*openxml.CompositeElementBase
}

// NewCommentReference creates a new CommentReference element with the given comment ID.
func NewCommentReference(
	id int,
) *CommentReference {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"commentReference",
		PrefixW,
	)
	cr := &CommentReference{
		CompositeElementBase: elem,
	}
	cr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
	return cr
}

// Id returns the referenced comment ID.
func (cr *CommentReference) Id() int {
	attr, found := cr.GetAttribute(
		"id",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())
	return id
}

// SetId sets the referenced comment ID.
func (cr *CommentReference) SetId(id int) {
	cr.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"id",
			PrefixW,
			strconv.Itoa(id),
		),
	)
}

// Clone creates a deep copy of this CommentReference element.
func (cr *CommentReference) Clone() openxml.Element {
	return &CommentReference{
		CompositeElementBase: cr.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CommentReference element.
func (cr *CommentReference) CloneNode(
	deep bool,
) openxml.Element {
	return &CommentReference{
		CompositeElementBase: cr.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}
