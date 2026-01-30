//nolint:revive // file-length-limit: this file contains all comment-related types
package elements

import (
	"iter"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
)

// CommentMention represents a user mention within comment content.
type CommentMention struct {
	Username string `json:"username"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
}

// CommentContent represents structured comment content with mentions.
type CommentContent struct {
	Text     string           `json:"text"`
	Mentions []CommentMention `json:"mentions"`
}

// CommentRun represents a run of text with optional mention formatting.
type CommentRun struct {
	Text      string `json:"text"`
	IsMention bool   `json:"isMention"`
}

// parseMentions extracts @username patterns from text and returns CommentMention slices.
func parseMentions(text string) []CommentMention {
	var mentions []CommentMention
	re := regexp.MustCompile(`@(\w+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			startPos := strings.Index(text, match[0])
			mentions = append(mentions, CommentMention{
				Username: match[1],
				Start:    startPos,
				End:      startPos + len(match[0]),
			})
		}
	}

	return mentions
}

// splitTextWithMentions splits text into runs, marking mentions separately.
func splitTextWithMentions(text string) []CommentRun {
	var runs []CommentRun
	re := regexp.MustCompile(`@(\w+)`)
	lastIndex := 0

	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		// Add text before mention
		if len(match) < 2 {
			continue
		}

		mentionText := match[0]
		mentionIndex := strings.Index(text, mentionText)

		if mentionIndex > lastIndex {
			runs = append(runs, CommentRun{
				Text:      text[lastIndex:mentionIndex],
				IsMention: false,
			})
		}

		// Add the mention itself
		runs = append(runs, CommentRun{
			Text:      mentionText,
			IsMention: true,
		})

		lastIndex = mentionIndex + len(mentionText)
	}

	// Add remaining text
	if lastIndex < len(text) {
		runs = append(runs, CommentRun{
			Text:      text[lastIndex:],
			IsMention: false,
		})
	}

	return runs
}

// CommentThread represents a threaded conversation with a root comment and replies.
type CommentThread struct {
	Root     *Comment   `json:"root"`
	Replies  []*Comment `json:"replies"`
	Resolved bool       `json:"resolved"`
}

// NewCommentThread creates a new CommentThread with the given root comment.
func NewCommentThread(root *Comment) *CommentThread {
	return &CommentThread{
		Root:     root,
		Replies:  []*Comment{},
		Resolved: root.Done(),
	}
}

// AddReply adds a reply comment to this thread.
func (ct *CommentThread) AddReply(reply *Comment) {
	if reply.ParentId() != ct.Root.Id() {
		// Set the parent ID if not already set
		reply.SetParentId(ct.Root.Id())
	}
	ct.Replies = append(ct.Replies, reply)
	ct.Resolved = reply.Done() || ct.Resolved
}

// GetAllComments returns all comments in this thread (root + replies).
func (ct *CommentThread) GetAllComments() []*Comment {
	all := make([]*Comment, 0, len(ct.Replies)+1)
	all = append(all, ct.Root)
	all = append(all, ct.Replies...)

	return all
}

// GetReplyCount returns the number of replies in this thread.
func (ct *CommentThread) GetReplyCount() int {
	return len(ct.Replies)
}

// MarkResolved marks the entire thread as resolved.
func (ct *CommentThread) MarkResolved() {
	ct.Resolved = true
	ct.Root.SetDone(true)
	for _, reply := range ct.Replies {
		reply.SetDone(true)
	}
}

// IsEmpty returns true if this thread has no replies.
func (ct *CommentThread) IsEmpty() bool {
	return len(ct.Replies) == 0
}

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
				switch v := child.(type) {
				case *Comment:
					comment = v
				case *openxml.CompositeElementBase:
					comment = &Comment{CompositeElementBase: v}
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

// RemoveComment removes the comment with the specified ID from the collection.
// Returns true if the comment was found and removed, false if not found.
func (c *Comments) RemoveComment(id int) bool {
	comment := c.GetComment(id)
	if comment == nil {
		return false
	}
	c.RemoveChild(comment)

	return true
}

// ByAuthor returns all comments by the specified author (case-sensitive).
// Returns an empty slice if no matching comments are found.
func (c *Comments) ByAuthor(
	author string,
) []*Comment {
	result := make([]*Comment, 0)
	for comment := range c.Comments() {
		if comment.Author() == author {
			result = append(result, comment)
		}
	}

	return result
}

// ByDateRange returns all comments within the specified date range (inclusive).
// If from is zero, no lower bound is applied. If to is zero, no upper bound is applied.
// Returns an empty slice if no matching comments are found.
func (c *Comments) ByDateRange(
	from, to time.Time,
) []*Comment {
	result := make([]*Comment, 0)
	for comment := range c.Comments() {
		date := comment.Date()
		// Skip zero dates
		if date.IsZero() {
			continue
		}

		// Apply lower bound if from is not zero
		if !from.IsZero() && date.Before(from) {
			continue
		}

		// Apply upper bound if to is not zero
		if !to.IsZero() && date.After(to) {
			continue
		}

		result = append(result, comment)
	}

	return result
}

// Count returns the total number of comments.
func (c *Comments) Count() int {
	count := 0
	for range c.Comments() {
		count++
	}

	return count
}

// AddW15Namespace ensures the w15 namespace is declared on this Comments element.
// This is required when using Word 2013+ features like the done attribute.
func (c *Comments) AddW15Namespace() {
	// Check if the namespace is already declared by looking for the xmlns:w15 attribute
	_, found := c.GetAttribute(
		"w15",
		"http://www.w3.org/2000/xmlns/",
	)
	if found {
		return
	}
	// Add the namespace declaration as an attribute
	c.SetAttribute(
		openxml.NewAttribute(
			"http://www.w3.org/2000/xmlns/",
			"w15",
			"xmlns",
			NamespaceW15,
		),
	)
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

// Done returns whether the comment is marked as done/resolved (w15:done attribute).
// Returns true if the attribute is set to "1" or "true", false otherwise.
func (c *Comment) Done() bool {
	attr, found := c.GetAttribute(
		"done",
		NamespaceW15,
	)
	if !found {
		return false
	}
	val := attr.Value()

	return val == "1" ||
		val == string(BooleanTrue)
}

// SetDone sets the comment's done/resolved status (w15:done attribute).
// When set to true, adds w15:done="1". When set to false, removes the attribute.
// This requires the w15 namespace to be declared on the comments root element.
func (c *Comment) SetDone(done bool) {
	if done {
		c.SetAttribute(
			openxml.NewAttribute(
				NamespaceW15,
				"done",
				PrefixW15,
				"1",
			),
		)
		// Ensure w15 namespace is declared on the root element
		c.ensureW15Namespace()
	} else {
		c.RemoveAttribute(
			"done",
			NamespaceW15,
		)
	}
}

// ParentId returns the parent comment ID for threaded replies.
// Returns 0 if no parent (root-level comment).
func (c *Comment) ParentId() int {
	attr, found := c.GetAttribute(
		"parentId",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	id, _ := strconv.Atoi(attr.Value())

	return id
}

// SetParentId sets the parent comment ID for threaded replies.
// Set to 0 for root-level comments (no parent).
func (c *Comment) SetParentId(parentId int) {
	if parentId == 0 {
		c.RemoveAttribute(
			"parentId",
			NamespaceWML,
		)
	} else {
		c.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"parentId",
				PrefixW,
				strconv.Itoa(parentId),
			),
		)
	}
}

// ensureW15Namespace ensures the w15 namespace is declared on the Comments root element.
func (c *Comment) ensureW15Namespace() {
	// Walk up to find the root Comments element
	current := c.Parent()
	for current != nil {
		// Check if this is the Comments root element by name and namespace
		if current.LocalName() == "comments" &&
			current.NamespaceURI() == NamespaceWML {
			// The parent is the Comments element, but we get it as the embedded base type
			// We need to add the w15 namespace attribute directly to this element
			_, found := current.GetAttribute(
				"w15",
				"http://www.w3.org/2000/xmlns/",
			)
			if !found {
				current.SetAttribute(
					openxml.NewAttribute(
						"http://www.w3.org/2000/xmlns/",
						"w15",
						"xmlns",
						NamespaceW15,
					),
				)
			}

			return
		}
		current = current.Parent()
	}
}

// Paragraphs returns an iterator over all paragraphs in the comment.
func (c *Comment) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range c.Children() {
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
func (c *Comment) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	c.AppendChild(p)

	return p
}

// ContentWithMentions returns the comment text with @mention parsing for user references.
// This scans paragraph content for @username patterns and returns a structured view
// of the content with extracted mentions.
func (c *Comment) ContentWithMentions() *CommentContent {
	var allText strings.Builder
	content := &CommentContent{
		Text:     "",
		Mentions: make([]CommentMention, 0),
	}

	// Parse @mentions from paragraph content
	for p := range c.Paragraphs() {
		for r := range p.Runs() {
			text := r.InnerText()
			content.Mentions = append(content.Mentions, parseMentions(text)...)
			allText.WriteString(text)
		}
	}

	content.Text = allText.String()

	return content
}

// SetExtendedContent sets comment content with support for rich text and mentions.
// Automatically handles @mention parsing and creates appropriate paragraph structure.
func (c *Comment) SetExtendedContent(text string) *CommentContent {
	content := &CommentContent{
		Text:     text,
		Mentions: parseMentions(text),
	}

	// Clear existing content
	for child := range c.Children() {
		c.RemoveChild(child)
	}

	// Create new paragraph with mention-aware runs
	p := c.AppendParagraph("")
	runs := splitTextWithMentions(text)

	for _, runData := range runs {
		run := p.AppendRun(runData.Text)

		// Apply formatting for mentions
		if runData.IsMention {
			run.SetColor("0563C1") // Blue color for mentions
			run.SetBold(true)
		}
	}

	return content
}

// IsReply returns true if this comment is a reply (has a parent).
func (c *Comment) IsReply() bool {
	return c.ParentId() > 0
}

// GetReplies returns all comments that are direct replies to this comment.
// This should be called on the Comments collection to find replies.
func (c *Comment) GetReplies(comments *Comments) []*Comment {
	var replies []*Comment
	for comment := range comments.Comments() {
		if comment.ParentId() == c.Id() {
			replies = append(replies, comment)
		}
	}

	return replies
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
