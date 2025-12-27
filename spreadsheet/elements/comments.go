package elements

//revive:disable:file-length-limit many comment properties

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Comments represents the comments root element (x:comments).
// This element is the root of a worksheet comments part.
type Comments struct {
	*openxml.PartRootElementBase
}

// NewComments creates a new Comments element.
func NewComments() *Comments {
	elem := openxml.NewPartRootElement(
		NamespaceSML,
		"comments",
		PrefixDefault,
	)

	return &Comments{
		PartRootElementBase: elem,
	}
}

// Authors returns the authors element (x:authors), or nil if not present.
func (c *Comments) Authors() *Authors {
	elem := c.GetElement("authors", NamespaceSML)
	if elem == nil {
		return nil
	}
	if a, ok := elem.(*Authors); ok {
		return a
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Authors{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateAuthors returns the authors element, creating if needed.
func (c *Comments) GetOrCreateAuthors() *Authors {
	a := c.Authors()
	if a != nil {
		return a
	}
	a = NewAuthors()
	// Authors should come before commentList
	cl := c.CommentList()
	if cl != nil {
		c.InsertBefore(a, cl)
	} else {
		c.AppendChild(a)
	}

	return a
}

// CommentList returns the comment list element (x:commentList), or nil if not
// present.
func (c *Comments) CommentList() *CommentList {
	elem := c.GetElement(
		"commentList",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cl, ok := elem.(*CommentList); ok {
		return cl
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CommentList{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCommentList returns the comment list element, creating if needed.
func (c *Comments) GetOrCreateCommentList() *CommentList {
	cl := c.CommentList()
	if cl != nil {
		return cl
	}
	cl = NewCommentList()
	c.AppendChild(cl)

	return cl
}

// Clone creates a deep copy of this Comments element.
func (c *Comments) Clone() openxml.Element {
	cloned := c.PartRootElementBase.Clone()

	return &Comments{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// CloneNode creates a copy of this Comments element.
func (c *Comments) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.PartRootElementBase.CloneNode(
		deep,
	)

	return &Comments{
		PartRootElementBase: cloned.(*openxml.PartRootElementBase),
	}
}

// Authors represents the authors element (x:authors).
// This element contains a list of comment authors.
type Authors struct {
	*openxml.CompositeElementBase
}

// NewAuthors creates a new Authors element.
func NewAuthors() *Authors {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"authors",
		PrefixDefault,
	)

	return &Authors{
		CompositeElementBase: elem,
	}
}

// AuthorList returns an iterator over all Author elements.
func (a *Authors) AuthorList() iter.Seq[*Author] {
	return func(yield func(*Author) bool) {
		for child := range a.Children() {
			if child.LocalName() != "author" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var author *Author
			if au, ok := child.(*Author); ok {
				author = au
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				author = &Author{LeafElementBase: leaf}
			}
			if author != nil && !yield(author) {
				return
			}
		}
	}
}

// AuthorCount returns the number of authors.
func (a *Authors) AuthorCount() int {
	count := 0
	for range a.AuthorList() {
		count++
	}

	return count
}

// GetAuthor returns the author at the specified index, or nil if out of range.
func (a *Authors) GetAuthor(index int) *Author {
	if index < 0 {
		return nil
	}
	i := 0
	for author := range a.AuthorList() {
		if i == index {
			return author
		}
		i++
	}

	return nil
}

// GetAuthorName returns the author name at the specified index.
func (a *Authors) GetAuthorName(
	index int,
) string {
	author := a.GetAuthor(index)
	if author == nil {
		return ""
	}

	return author.Name()
}

// AddAuthor adds a new author and returns it.
func (a *Authors) AddAuthor(name string) *Author {
	author := NewAuthor()
	author.SetName(name)
	a.AppendChild(author)

	return author
}

// FindAuthorIndex returns the index of the author with the given name, or -1
// if not found.
func (a *Authors) FindAuthorIndex(
	name string,
) int {
	i := 0
	for author := range a.AuthorList() {
		if author.Name() == name {
			return i
		}
		i++
	}

	return -1
}

// GetOrAddAuthor returns the index of the author with the given name,
// adding a new author if not found.
func (a *Authors) GetOrAddAuthor(
	name string,
) int {
	idx := a.FindAuthorIndex(name)
	if idx >= 0 {
		return idx
	}
	a.AddAuthor(name)

	return a.AuthorCount() - 1
}

// Clone creates a deep copy of this Authors element.
func (a *Authors) Clone() openxml.Element {
	cloned := a.CompositeElementBase.Clone()

	return &Authors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Authors element.
func (a *Authors) CloneNode(
	deep bool,
) openxml.Element {
	cloned := a.CompositeElementBase.CloneNode(
		deep,
	)

	return &Authors{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Author represents an author element (x:author).
// This element contains the name of a comment author.
type Author struct {
	*openxml.LeafElementBase
}

// NewAuthor creates a new Author element.
func NewAuthor() *Author {
	elem := openxml.NewLeafElement(
		NamespaceSML,
		"author",
		PrefixDefault,
	)

	return &Author{LeafElementBase: elem}
}

// Name returns the author name.
func (a *Author) Name() string {
	return a.InnerText()
}

// SetName sets the author name.
func (a *Author) SetName(name string) {
	a.SetInnerText(name)
}

// Clone creates a deep copy of this Author element.
func (a *Author) Clone() openxml.Element {
	cloned := a.LeafElementBase.Clone()

	return &Author{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this Author element.
func (a *Author) CloneNode(
	deep bool,
) openxml.Element {
	cloned := a.LeafElementBase.CloneNode(deep)

	return &Author{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CommentList represents the comment list element (x:commentList).
// This element contains all comments for a worksheet.
type CommentList struct {
	*openxml.CompositeElementBase
}

// NewCommentList creates a new CommentList element.
func NewCommentList() *CommentList {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"commentList",
		PrefixDefault,
	)

	return &CommentList{
		CompositeElementBase: elem,
	}
}

// Comments returns an iterator over all Comment elements.
func (cl *CommentList) Comments() iter.Seq[*Comment] {
	return func(yield func(*Comment) bool) {
		for child := range cl.Children() {
			if child.LocalName() != "comment" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var comment *Comment
			if c, ok := child.(*Comment); ok {
				comment = c
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				comment = &Comment{CompositeElementBase: comp}
			}
			if comment != nil && !yield(comment) {
				return
			}
		}
	}
}

// CommentCount returns the number of comments.
func (cl *CommentList) CommentCount() int {
	count := 0
	for range cl.Comments() {
		count++
	}

	return count
}

// GetComment returns the comment at the specified index, or nil if out of
// range.
func (cl *CommentList) GetComment(
	index int,
) *Comment {
	if index < 0 {
		return nil
	}
	i := 0
	for comment := range cl.Comments() {
		if i == index {
			return comment
		}
		i++
	}

	return nil
}

// GetCommentByRef returns the comment for the given cell reference, or nil if
// not found.
func (cl *CommentList) GetCommentByRef(
	ref string,
) *Comment {
	for comment := range cl.Comments() {
		if comment.Ref() == ref {
			return comment
		}
	}

	return nil
}

// AddComment adds a new comment and returns it.
func (cl *CommentList) AddComment(
	ref string,
	authorId int,
) *Comment {
	comment := NewComment()
	comment.SetRef(ref)
	comment.SetAuthorId(authorId)
	cl.AppendChild(comment)

	return comment
}

// RemoveComment removes a comment by cell reference.
// Returns true if a comment was removed, false if no comment was found.
func (cl *CommentList) RemoveComment(
	ref string,
) bool {
	comment := cl.GetCommentByRef(ref)
	if comment != nil {
		cl.RemoveChild(comment)

		return true
	}

	return false
}

// Clone creates a deep copy of this CommentList element.
func (cl *CommentList) Clone() openxml.Element {
	cloned := cl.CompositeElementBase.Clone()

	return &CommentList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CommentList element.
func (cl *CommentList) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cl.CompositeElementBase.CloneNode(
		deep,
	)

	return &CommentList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Comment represents a comment element (x:comment).
// This element contains a single cell comment.
type Comment struct {
	*openxml.CompositeElementBase
}

// NewComment creates a new Comment element.
func NewComment() *Comment {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"comment",
		PrefixDefault,
	)

	return &Comment{
		CompositeElementBase: elem,
	}
}

// Ref returns the cell reference for this comment (e.g., "A1").
func (c *Comment) Ref() string {
	attr, found := c.GetAttribute(elemNameRef, "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell reference for this comment.
func (c *Comment) SetRef(ref string) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			elemNameRef,
			"",
			ref,
		),
	)
}

// AuthorId returns the author index for this comment.
func (c *Comment) AuthorId() int {
	attr, found := c.GetAttribute("authorId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetAuthorId sets the author index for this comment.
func (c *Comment) SetAuthorId(authorId int) {
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"authorId",
			"",
			strconv.Itoa(authorId),
		),
	)
}

// GUID returns the globally unique identifier for this comment.
func (c *Comment) GUID() string {
	attr, found := c.GetAttribute("guid", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetGUID sets the globally unique identifier for this comment.
func (c *Comment) SetGUID(guid string) {
	if guid == "" {
		c.RemoveAttribute("guid", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute("", "guid", "", guid),
		)
	}
}

// ShapeId returns the VML shape ID for this comment.
func (c *Comment) ShapeId() int {
	attr, found := c.GetAttribute("shapeId", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetShapeId sets the VML shape ID for this comment.
func (c *Comment) SetShapeId(shapeId int) {
	if shapeId == 0 {
		c.RemoveAttribute("shapeId", "")
	} else {
		c.SetAttribute(
			openxml.NewAttribute(
				"",
				"shapeId",
				"",
				strconv.Itoa(shapeId),
			),
		)
	}
}

// Text returns the comment text element (x:text), or nil if not present.
func (c *Comment) Text() *CommentText {
	elem := c.GetElement("text", NamespaceSML)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*CommentText); ok {
		return t
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CommentText{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateText returns the comment text element, creating if needed.
func (c *Comment) GetOrCreateText() *CommentText {
	t := c.Text()
	if t != nil {
		return t
	}
	t = NewCommentText()
	c.AppendChild(t)

	return t
}

// PlainText returns the plain text content of the comment.
func (c *Comment) PlainText() string {
	t := c.Text()
	if t == nil {
		return ""
	}

	return t.PlainText()
}

// SetPlainText sets the comment content as plain text.
func (c *Comment) SetPlainText(text string) {
	t := c.GetOrCreateText()
	t.SetPlainText(text)
}

// Clone creates a deep copy of this Comment element.
func (c *Comment) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &Comment{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Comment element.
func (c *Comment) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &Comment{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CommentText represents the comment text element (x:text).
// This element contains the rich text content of a comment.
type CommentText struct {
	*openxml.CompositeElementBase
}

// NewCommentText creates a new CommentText element.
func NewCommentText() *CommentText {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"text",
		PrefixDefault,
	)

	return &CommentText{
		CompositeElementBase: elem,
	}
}

// PlainText returns the plain text content.
// If the text contains rich text runs, concatenates all text from runs.
// If the text contains a simple text element, returns that text.
func (ct *CommentText) PlainText() string {
	// First check for simple text element
	elem := ct.GetElement("t", NamespaceSML)
	if elem != nil {
		if t, ok := elem.(*Text); ok {
			return t.Text()
		}
		if leaf, ok := elem.(*openxml.LeafElementBase); ok {
			return leaf.InnerText()
		}
	}

	// Concatenate text from rich text runs
	var result string
	for rtr := range ct.RichTextRuns() {
		if text := rtr.Text(); text != nil {
			result += text.Text()
		}
	}

	return result
}

// SetPlainText sets the content as simple text.
// This removes any rich text runs and uses a single text element.
func (ct *CommentText) SetPlainText(text string) {
	// Remove all rich text runs
	for rtr := range ct.RichTextRuns() {
		ct.RemoveChild(rtr)
	}

	// Get or create text element
	elem := ct.GetElement("t", NamespaceSML)
	if elem == nil {
		t := NewText()
		t.SetText(text)
		ct.AppendChild(t)
	} else if t, ok := elem.(*Text); ok {
		t.SetText(text)
	} else if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		leaf.SetInnerText(text)
	}
}

// RichTextRuns returns an iterator over all rich text run elements (x:r).
func (ct *CommentText) RichTextRuns() iter.Seq[*RichTextRun] {
	return func(yield func(*RichTextRun) bool) {
		for child := range ct.Children() {
			if child.LocalName() != "r" ||
				child.NamespaceURI() != NamespaceSML {
				continue
			}
			var rtr *RichTextRun
			if r, ok := child.(*RichTextRun); ok {
				rtr = r
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				rtr = &RichTextRun{CompositeElementBase: comp}
			}
			if rtr != nil && !yield(rtr) {
				return
			}
		}
	}
}

// AddRichTextRun adds a new rich text run and returns it.
// If the text currently contains simple text, it removes it.
func (ct *CommentText) AddRichTextRun() *RichTextRun {
	// Remove simple text element if present
	elem := ct.GetElement("t", NamespaceSML)
	if elem != nil {
		ct.RemoveChild(elem)
	}

	rtr := NewRichTextRun()
	ct.AppendChild(rtr)

	return rtr
}

// AddRichTextRunWithText adds a new rich text run with the given text.
func (ct *CommentText) AddRichTextRunWithText(
	text string,
) *RichTextRun {
	rtr := ct.AddRichTextRun()
	rtr.GetOrCreateText().SetText(text)

	return rtr
}

// IsRichText returns whether the comment text contains rich text runs.
func (ct *CommentText) IsRichText() bool {
	for range ct.RichTextRuns() {
		return true
	}

	return false
}

// Clone creates a deep copy of this CommentText element.
func (ct *CommentText) Clone() openxml.Element {
	cloned := ct.CompositeElementBase.Clone()

	return &CommentText{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CommentText element.
func (ct *CommentText) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ct.CompositeElementBase.CloneNode(
		deep,
	)

	return &CommentText{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
