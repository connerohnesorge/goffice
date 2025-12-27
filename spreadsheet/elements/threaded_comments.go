package elements

//revive:disable:file-length-limit many threaded comment types

import (
	"iter"
	"strconv"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
)

// Threaded comments namespace constants.
const (
	// NamespaceThreadedComments is the threaded comments namespace.
	//nolint:revive // line-length-limit
	NamespaceThreadedComments = "http://schemas.microsoft.com/office/spreadsheetml/2018/threadedcomments"

	// PrefixThreadedComments is the prefix for threaded comments elements.
	PrefixThreadedComments = "tc"
)

// ThreadedComments represents the threaded comments root element. This element
// is the root of a threaded comments part (introduced in Excel 2019).
type ThreadedComments struct {
	*openxml.CompositeElementBase
}

// NewThreadedComments creates a new ThreadedComments element.
func NewThreadedComments() *ThreadedComments {
	elem := openxml.NewCompositeElement(
		NamespaceThreadedComments,
		"ThreadedComments",
		PrefixThreadedComments,
	)

	return &ThreadedComments{
		CompositeElementBase: elem,
	}
}

// Comments returns an iterator over all ThreadedComment elements.
func (tc *ThreadedComments) Comments() iter.Seq[*ThreadedComment] {
	return func(yield func(*ThreadedComment) bool) {
		for child := range tc.Children() {
			if child.LocalName() != "threadedComment" ||
				child.NamespaceURI() != NamespaceThreadedComments {
				continue
			}
			var comment *ThreadedComment
			if c, ok := child.(*ThreadedComment); ok {
				comment = c
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				comment = &ThreadedComment{CompositeElementBase: comp}
			}
			if comment != nil && !yield(comment) {
				return
			}
		}
	}
}

// CommentCount returns the number of threaded comments.
func (tc *ThreadedComments) CommentCount() int {
	count := 0
	for range tc.Comments() {
		count++
	}

	return count
}

// GetComment returns the comment at the specified index, nil if out of range.
func (tc *ThreadedComments) GetComment(
	index int,
) *ThreadedComment {
	if index < 0 {
		return nil
	}
	i := 0
	for comment := range tc.Comments() {
		if i == index {
			return comment
		}
		i++
	}

	return nil
}

// GetCommentByID returns the comment with the given ID, or nil if not found.
func (tc *ThreadedComments) GetCommentByID(
	id string,
) *ThreadedComment {
	for comment := range tc.Comments() {
		if comment.ID() == id {
			return comment
		}
	}

	return nil
}

// GetCommentsByRef returns all comments for the given cell reference.
func (tc *ThreadedComments) GetCommentsByRef(
	ref string,
) []*ThreadedComment {
	var comments []*ThreadedComment
	for comment := range tc.Comments() {
		if comment.Ref() == ref {
			comments = append(comments, comment)
		}
	}

	return comments
}

// GetRootComments returns all root comments (comments that are not replies).
func (tc *ThreadedComments) GetRootComments() []*ThreadedComment {
	var roots []*ThreadedComment
	for comment := range tc.Comments() {
		if comment.ParentID() == "" {
			roots = append(roots, comment)
		}
	}

	return roots
}

// GetReplies returns all replies to the comment with the given ID.
func (tc *ThreadedComments) GetReplies(
	parentID string,
) []*ThreadedComment {
	var replies []*ThreadedComment
	for comment := range tc.Comments() {
		if comment.ParentID() == parentID {
			replies = append(replies, comment)
		}
	}

	return replies
}

// AddComment adds a new threaded comment and returns it.
func (tc *ThreadedComments) AddComment() *ThreadedComment {
	comment := NewThreadedComment()
	tc.AppendChild(comment)

	return comment
}

// RemoveComment removes a comment by ID.
// Returns true if a comment was removed, false if no comment was found.
func (tc *ThreadedComments) RemoveComment(
	id string,
) bool {
	comment := tc.GetCommentByID(id)
	if comment != nil {
		tc.RemoveChild(comment)

		return true
	}

	return false
}

// Clone creates a deep copy of this ThreadedComments element.
func (tc *ThreadedComments) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &ThreadedComments{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ThreadedComments element.
func (tc *ThreadedComments) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &ThreadedComments{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ThreadedComment represents a threaded comment element.
// This element contains a single threaded comment.
type ThreadedComment struct {
	*openxml.CompositeElementBase
}

// NewThreadedComment creates a new ThreadedComment element.
func NewThreadedComment() *ThreadedComment {
	elem := openxml.NewCompositeElement(
		NamespaceThreadedComments,
		"threadedComment",
		PrefixThreadedComments,
	)

	return &ThreadedComment{
		CompositeElementBase: elem,
	}
}

// ID returns the unique identifier for this comment.
func (tc *ThreadedComment) ID() string {
	attr, found := tc.GetAttribute("id", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetID sets the unique identifier for this comment.
func (tc *ThreadedComment) SetID(id string) {
	tc.SetAttribute(
		openxml.NewAttribute("", "id", "", id),
	)
}

// Ref returns the cell reference for this comment (e.g., "A1").
func (tc *ThreadedComment) Ref() string {
	attr, found := tc.GetAttribute("ref", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRef sets the cell reference for this comment.
func (tc *ThreadedComment) SetRef(ref string) {
	tc.SetAttribute(
		openxml.NewAttribute("", "ref", "", ref),
	)
}

// PersonID returns the person ID of the comment author.
func (tc *ThreadedComment) PersonID() string {
	attr, found := tc.GetAttribute("personId", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetPersonID sets the person ID of the comment author.
func (tc *ThreadedComment) SetPersonID(
	personId string,
) {
	tc.SetAttribute(
		openxml.NewAttribute(
			"",
			"personId",
			"",
			personId,
		),
	)
}

// ParentID returns the parent comment ID if this is a reply.
func (tc *ThreadedComment) ParentID() string {
	attr, found := tc.GetAttribute("parentId", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetParentID sets the parent comment ID (for replies).
func (tc *ThreadedComment) SetParentID(
	parentId string,
) {
	if parentId == "" {
		tc.RemoveAttribute("parentId", "")
	} else {
		tc.SetAttribute(
			openxml.NewAttribute("", "parentId", "", parentId),
		)
	}
}

// DT returns the date/time when the comment was created.
func (tc *ThreadedComment) DT() time.Time {
	attr, found := tc.GetAttribute("dT", "")
	if !found {
		return time.Time{}
	}
	// Parse ISO 8601 date time
	t, err := time.Parse(
		time.RFC3339,
		attr.Value(),
	)
	if err != nil {
		// Try alternative formats
		t, err = time.Parse(
			"2006-01-02T15:04:05",
			attr.Value(),
		)
		if err != nil {
			return time.Time{}
		}
	}

	return t
}

// SetDT sets the date/time when the comment was created.
func (tc *ThreadedComment) SetDT(dt time.Time) {
	if dt.IsZero() {
		tc.RemoveAttribute("dT", "")
	} else {
		tc.SetAttribute(
			openxml.NewAttribute(
				"",
				"dT",
				"",
				dt.Format(time.RFC3339),
			),
		)
	}
}

// Done returns whether this comment thread is marked as resolved.
func (tc *ThreadedComment) Done() bool {
	attr, found := tc.GetAttribute("done", "")
	if !found {
		return false
	}
	val := attr.Value()

	return val == attrValueTrue ||
		val == attrValueOne
}

// SetDone sets whether this comment thread is marked as resolved.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (tc *ThreadedComment) SetDone(done bool) {
	if done {
		tc.SetAttribute(
			openxml.NewAttribute(
				"",
				"done",
				"",
				attrValueTrue,
			),
		)
	} else {
		tc.RemoveAttribute("done", "")
	}
}

// Text returns the comment text element, or nil if not present.
func (tc *ThreadedComment) Text() *ThreadedCommentText {
	elem := tc.GetElement(
		"text",
		NamespaceThreadedComments,
	)
	if elem == nil {
		return nil
	}
	if t, ok := elem.(*ThreadedCommentText); ok {
		return t
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		return &ThreadedCommentText{
			LeafElementBase: leaf,
		}
	}

	return nil
}

// GetOrCreateText returns the comment text element, creating if needed.
func (tc *ThreadedComment) GetOrCreateText() *ThreadedCommentText {
	t := tc.Text()
	if t != nil {
		return t
	}
	t = NewThreadedCommentText()
	tc.AppendChild(t)

	return t
}

// TextContent returns the text content of the comment.
func (tc *ThreadedComment) TextContent() string {
	t := tc.Text()
	if t == nil {
		return ""
	}

	return t.Text()
}

// SetTextContent sets the text content of the comment.
func (tc *ThreadedComment) SetTextContent(
	text string,
) {
	t := tc.GetOrCreateText()
	t.SetText(text)
}

// Mentions returns an iterator over all mention elements.
func (tc *ThreadedComment) Mentions() iter.Seq[*Mention] {
	return func(yield func(*Mention) bool) {
		for child := range tc.Children() {
			if child.LocalName() != "mention" ||
				child.NamespaceURI() != NamespaceThreadedComments {
				continue
			}
			var mention *Mention
			if m, ok := child.(*Mention); ok {
				mention = m
			} else if leaf, ok := child.(*openxml.LeafElementBase); ok {
				mention = &Mention{LeafElementBase: leaf}
			}
			if mention != nil && !yield(mention) {
				return
			}
		}
	}
}

// AddMention adds a mention to this comment.
func (tc *ThreadedComment) AddMention(
	mentionPersonID string,
	startIndex, length int,
) *Mention {
	mention := NewMention()
	mention.SetMentionPersonID(mentionPersonID)
	mention.SetStartIndex(startIndex)
	mention.SetLength(length)
	tc.AppendChild(mention)

	return mention
}

// Clone creates a deep copy of this ThreadedComment element.
func (tc *ThreadedComment) Clone() openxml.Element {
	cloned := tc.CompositeElementBase.Clone()

	return &ThreadedComment{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ThreadedComment element.
func (tc *ThreadedComment) CloneNode(
	deep bool,
) openxml.Element {
	cloned := tc.CompositeElementBase.CloneNode(
		deep,
	)

	return &ThreadedComment{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ThreadedCommentText represents the text element of a threaded comment.
type ThreadedCommentText struct {
	*openxml.LeafElementBase
}

// NewThreadedCommentText creates a new ThreadedCommentText element.
func NewThreadedCommentText() *ThreadedCommentText {
	elem := openxml.NewLeafElement(
		NamespaceThreadedComments,
		"text",
		PrefixThreadedComments,
	)

	return &ThreadedCommentText{
		LeafElementBase: elem,
	}
}

// Text returns the text content.
func (t *ThreadedCommentText) Text() string {
	return t.InnerText()
}

// SetText sets the text content.
func (t *ThreadedCommentText) SetText(
	text string,
) {
	t.SetInnerText(text)
}

// Clone creates a deep copy of this element.
func (t *ThreadedCommentText) Clone() openxml.Element {
	cloned := t.LeafElementBase.Clone()

	return &ThreadedCommentText{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (t *ThreadedCommentText) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.LeafElementBase.CloneNode(deep)

	return &ThreadedCommentText{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// Mention represents a mention element in a threaded comment.
// This is used to @mention people in comment threads.
type Mention struct {
	*openxml.LeafElementBase
}

// NewMention creates a new Mention element.
func NewMention() *Mention {
	elem := openxml.NewLeafElement(
		NamespaceThreadedComments,
		"mention",
		PrefixThreadedComments,
	)

	return &Mention{LeafElementBase: elem}
}

// MentionPersonID returns the person ID of the mentioned person.
func (m *Mention) MentionPersonID() string {
	attr, found := m.GetAttribute(
		"mentionpersonId",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMentionPersonID sets the person ID of the mentioned person.
func (m *Mention) SetMentionPersonID(
	personId string,
) {
	m.SetAttribute(
		openxml.NewAttribute(
			"",
			"mentionpersonId",
			"",
			personId,
		),
	)
}

// StartIndex returns the start index of the mention in the text.
func (m *Mention) StartIndex() int {
	attr, found := m.GetAttribute(
		"startIndex",
		"",
	)
	if !found {
		return 0
	}
	val, _ := parseInt(attr.Value())

	return val
}

// SetStartIndex sets the start index of the mention in the text.
func (m *Mention) SetStartIndex(index int) {
	m.SetAttribute(
		openxml.NewAttribute(
			"",
			"startIndex",
			"",
			formatInt(index),
		),
	)
}

// Length returns the length of the mention text.
func (m *Mention) Length() int {
	attr, found := m.GetAttribute("length", "")
	if !found {
		return 0
	}
	val, _ := parseInt(attr.Value())

	return val
}

// SetLength sets the length of the mention text.
func (m *Mention) SetLength(length int) {
	m.SetAttribute(
		openxml.NewAttribute(
			"",
			"length",
			"",
			formatInt(length),
		),
	)
}

// Clone creates a deep copy of this element.
func (m *Mention) Clone() openxml.Element {
	cloned := m.LeafElementBase.Clone()

	return &Mention{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// CloneNode creates a copy of this element.
func (m *Mention) CloneNode(
	deep bool,
) openxml.Element {
	cloned := m.LeafElementBase.CloneNode(deep)

	return &Mention{
		LeafElementBase: cloned.(*openxml.LeafElementBase),
	}
}

// PersonList represents the persons root element. This element contains the
// list of people who can be mentioned in threaded comments.
type PersonList struct {
	*openxml.CompositeElementBase
}

// Persons namespace constant.
const (
	// NamespacePersons is the persons namespace.
	//nolint:revive // line-length-limit
	NamespacePersons = "http://schemas.microsoft.com/office/spreadsheetml/2018/threadedcomments"

	// PrefixPersons is the prefix for persons elements.
	PrefixPersons = "tc"
)

// NewPersonList creates a new PersonList element.
func NewPersonList() *PersonList {
	elem := openxml.NewCompositeElement(
		NamespacePersons,
		"personList",
		PrefixPersons,
	)

	return &PersonList{
		CompositeElementBase: elem,
	}
}

// Persons returns an iterator over all Person elements.
func (pl *PersonList) Persons() iter.Seq[*Person] {
	return func(yield func(*Person) bool) {
		for child := range pl.Children() {
			if child.LocalName() != "person" ||
				child.NamespaceURI() != NamespacePersons {
				continue
			}
			var person *Person
			if p, ok := child.(*Person); ok {
				person = p
			} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
				person = &Person{CompositeElementBase: comp}
			}
			if person != nil && !yield(person) {
				return
			}
		}
	}
}

// PersonCount returns the number of persons.
func (pl *PersonList) PersonCount() int {
	count := 0
	for range pl.Persons() {
		count++
	}

	return count
}

// GetPerson returns the person at the specified index, or nil if out of range.
func (pl *PersonList) GetPerson(
	index int,
) *Person {
	if index < 0 {
		return nil
	}
	i := 0
	for person := range pl.Persons() {
		if i == index {
			return person
		}
		i++
	}

	return nil
}

// GetPersonByID returns the person with the given ID, or nil if not found.
func (pl *PersonList) GetPersonByID(
	id string,
) *Person {
	for person := range pl.Persons() {
		if person.ID() == id {
			return person
		}
	}

	return nil
}

// GetPersonByUserID returns the person with the given user ID, or nil if not
// found.
func (pl *PersonList) GetPersonByUserID(
	userID string,
) *Person {
	for person := range pl.Persons() {
		if person.UserID() == userID {
			return person
		}
	}

	return nil
}

// AddPerson adds a new person and returns it.
func (pl *PersonList) AddPerson() *Person {
	person := NewPerson()
	pl.AppendChild(person)

	return person
}

// RemovePerson removes a person by ID.
// Returns true if a person was removed, false if no person was found.
func (pl *PersonList) RemovePerson(
	id string,
) bool {
	person := pl.GetPersonByID(id)
	if person != nil {
		pl.RemoveChild(person)

		return true
	}

	return false
}

// Clone creates a deep copy of this PersonList element.
func (pl *PersonList) Clone() openxml.Element {
	cloned := pl.CompositeElementBase.Clone()

	return &PersonList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this PersonList element.
func (pl *PersonList) CloneNode(
	deep bool,
) openxml.Element {
	cloned := pl.CompositeElementBase.CloneNode(
		deep,
	)

	return &PersonList{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Person represents a person element.
// This element contains information about a person who can be mentioned.
type Person struct {
	*openxml.CompositeElementBase
}

// NewPerson creates a new Person element.
func NewPerson() *Person {
	elem := openxml.NewCompositeElement(
		NamespacePersons,
		"person",
		PrefixPersons,
	)

	return &Person{
		CompositeElementBase: elem,
	}
}

// ID returns the unique identifier for this person.
func (p *Person) ID() string {
	attr, found := p.GetAttribute("id", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetID sets the unique identifier for this person.
func (p *Person) SetID(id string) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"id", //nolint:revive // add-constant
			"",
			id,
		),
	)
}

// DisplayName returns the display name of this person.
func (p *Person) DisplayName() string {
	attr, found := p.GetAttribute(
		"displayName",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetDisplayName sets the display name of this person.
func (p *Person) SetDisplayName(name string) {
	if name == "" {
		p.RemoveAttribute("displayName", "")
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "displayName", "", name),
		)
	}
}

// UserID returns the user ID of this person.
func (p *Person) UserID() string {
	attr, found := p.GetAttribute("userId", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetUserID sets the user ID of this person.
func (p *Person) SetUserID(userID string) {
	if userID == "" {
		p.RemoveAttribute("userId", "")
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "userId", "", userID),
		)
	}
}

// ProviderID returns the provider ID for this person.
func (p *Person) ProviderID() string {
	attr, found := p.GetAttribute(
		"providerId",
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetProviderID sets the provider ID for this person.
func (p *Person) SetProviderID(
	providerID string,
) {
	if providerID == "" {
		p.RemoveAttribute("providerId", "")
	} else {
		p.SetAttribute(
			openxml.NewAttribute("", "providerId", "", providerID),
		)
	}
}

// Clone creates a deep copy of this Person element.
func (p *Person) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &Person{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Person element.
func (p *Person) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &Person{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Helper functions for parsing/formatting integers.
func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func formatInt(i int) string {
	return strconv.Itoa(i)
}
