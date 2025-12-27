package elements

import (
	"strings"
	"testing"
	"time"
)

func TestThreadedCommentsCreation(t *testing.T) {
	tc := NewThreadedComments()

	if tc == nil {
		t.Fatal(
			"NewThreadedComments returned nil",
		)
	}

	if tc.LocalName() != "ThreadedComments" {
		t.Errorf(
			"expected local name 'ThreadedComments', got '%s'",
			tc.LocalName(),
		)
	}

	if tc.NamespaceURI() != NamespaceThreadedComments {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceThreadedComments,
			tc.NamespaceURI(),
		)
	}

	// Initially empty
	if tc.CommentCount() != 0 {
		t.Errorf(
			"expected 0 comments, got %d",
			tc.CommentCount(),
		)
	}
}

func TestThreadedCommentsAddAndGet(t *testing.T) {
	tc := NewThreadedComments()

	// Add comments
	c1 := tc.AddComment()
	c1.SetID("{comment-1}")
	c1.SetRef("A1")
	c1.SetPersonID("{person-1}")
	c1.SetTextContent("First comment")

	c2 := tc.AddComment()
	c2.SetID("{comment-2}")
	c2.SetRef("B2")
	c2.SetPersonID("{person-2}")
	c2.SetTextContent("Second comment")

	// Check count
	if tc.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments, got %d",
			tc.CommentCount(),
		)
	}

	// Get by index
	if tc.GetComment(0).ID() != "{comment-1}" {
		t.Error(
			"expected ID '{comment-1}' at index 0",
		)
	}

	if tc.GetComment(1).ID() != "{comment-2}" {
		t.Error(
			"expected ID '{comment-2}' at index 1",
		)
	}

	// Get by ID
	comment := tc.GetCommentByID("{comment-1}")
	if comment == nil {
		t.Fatal("GetCommentByID returned nil")
	}

	if comment.Ref() != "A1" {
		t.Errorf(
			"expected ref 'A1', got '%s'",
			comment.Ref(),
		)
	}

	// Unknown ID
	if tc.GetCommentByID("{unknown}") != nil {
		t.Error("expected nil for unknown ID")
	}

	// Invalid indices
	if tc.GetComment(-1) != nil {
		t.Error("expected nil for negative index")
	}

	if tc.GetComment(10) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}
}

func TestThreadedCommentsGetByRef(t *testing.T) {
	tc := NewThreadedComments()

	// Add multiple comments for same cell
	c1 := tc.AddComment()
	c1.SetID("{comment-1}")
	c1.SetRef("A1")
	c1.SetTextContent("First comment on A1")

	c2 := tc.AddComment()
	c2.SetID("{comment-2}")
	c2.SetRef("A1")
	c2.SetParentID("{comment-1}")
	c2.SetTextContent("Reply to first comment")

	c3 := tc.AddComment()
	c3.SetID("{comment-3}")
	c3.SetRef("B2")
	c3.SetTextContent("Comment on B2")

	// Get comments for A1
	comments := tc.GetCommentsByRef("A1")
	if len(comments) != 2 {
		t.Errorf(
			"expected 2 comments for A1, got %d",
			len(comments),
		)
	}

	// Get comments for B2
	comments = tc.GetCommentsByRef("B2")
	if len(comments) != 1 {
		t.Errorf(
			"expected 1 comment for B2, got %d",
			len(comments),
		)
	}

	// Get comments for unknown ref
	comments = tc.GetCommentsByRef("Z99")
	if len(comments) != 0 {
		t.Errorf(
			"expected 0 comments for Z99, got %d",
			len(comments),
		)
	}
}

func TestThreadedCommentsRootComments(
	t *testing.T,
) {
	tc := NewThreadedComments()

	// Add root comments
	root1 := tc.AddComment()
	root1.SetID("{root-1}")
	root1.SetRef("A1")
	root1.SetTextContent("Root comment 1")

	root2 := tc.AddComment()
	root2.SetID("{root-2}")
	root2.SetRef("B2")
	root2.SetTextContent("Root comment 2")

	// Add replies
	reply1 := tc.AddComment()
	reply1.SetID("{reply-1}")
	reply1.SetRef("A1")
	reply1.SetParentID("{root-1}")
	reply1.SetTextContent("Reply to root 1")

	reply2 := tc.AddComment()
	reply2.SetID("{reply-2}")
	reply2.SetRef("A1")
	reply2.SetParentID("{root-1}")
	reply2.SetTextContent(
		"Another reply to root 1",
	)

	// Get root comments
	roots := tc.GetRootComments()
	if len(roots) != 2 {
		t.Errorf(
			"expected 2 root comments, got %d",
			len(roots),
		)
	}

	// Get replies
	replies := tc.GetReplies("{root-1}")
	if len(replies) != 2 {
		t.Errorf(
			"expected 2 replies, got %d",
			len(replies),
		)
	}

	replies = tc.GetReplies("{root-2}")
	if len(replies) != 0 {
		t.Errorf(
			"expected 0 replies for root-2, got %d",
			len(replies),
		)
	}
}

func TestThreadedCommentsRemove(t *testing.T) {
	tc := NewThreadedComments()

	c1 := tc.AddComment()
	c1.SetID("{comment-1}")

	c2 := tc.AddComment()
	c2.SetID("{comment-2}")

	c3 := tc.AddComment()
	c3.SetID("{comment-3}")

	if tc.CommentCount() != 3 {
		t.Fatalf(
			"expected 3 comments, got %d",
			tc.CommentCount(),
		)
	}

	// Remove middle comment
	removed := tc.RemoveComment("{comment-2}")
	if !removed {
		t.Error(
			"expected RemoveComment to return true",
		)
	}

	if tc.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments after remove, got %d",
			tc.CommentCount(),
		)
	}

	if tc.GetCommentByID("{comment-2}") != nil {
		t.Error("expected comment to be removed")
	}

	// Remove non-existent
	removed = tc.RemoveComment("{unknown}")
	if removed {
		t.Error(
			"expected RemoveComment to return false for unknown ID",
		)
	}
}

func TestThreadedCommentElement(t *testing.T) {
	c := NewThreadedComment()

	if c == nil {
		t.Fatal("NewThreadedComment returned nil")
	}

	if c.LocalName() != "threadedComment" {
		t.Errorf(
			"expected local name 'threadedComment', got '%s'",
			c.LocalName(),
		)
	}

	// Set and get ID
	c.SetID("{test-id}")
	if c.ID() != "{test-id}" {
		t.Errorf(
			"expected ID '{test-id}', got '%s'",
			c.ID(),
		)
	}

	// Set and get ref
	c.SetRef("D4")
	if c.Ref() != "D4" {
		t.Errorf(
			"expected ref 'D4', got '%s'",
			c.Ref(),
		)
	}

	// Set and get personId
	c.SetPersonID("{person-123}")
	if c.PersonID() != "{person-123}" {
		t.Errorf(
			"expected personId '{person-123}', got '%s'",
			c.PersonID(),
		)
	}

	// Set and get parentId
	c.SetParentID("{parent-id}")
	if c.ParentID() != "{parent-id}" {
		t.Errorf(
			"expected parentId '{parent-id}', got '%s'",
			c.ParentID(),
		)
	}

	// Clear parentId
	c.SetParentID("")
	if c.ParentID() != "" {
		t.Error("expected parentId to be cleared")
	}
}

func TestThreadedCommentDateTime(t *testing.T) {
	c := NewThreadedComment()

	// Initially no date
	if !c.DT().IsZero() {
		t.Error(
			"expected DT() to be zero initially",
		)
	}

	// Set date
	testTime := time.Date(
		2024,
		1,
		15,
		10,
		30,
		0,
		0,
		time.UTC,
	)
	c.SetDT(testTime)

	got := c.DT()
	if got.Year() != 2024 || got.Month() != 1 ||
		got.Day() != 15 {
		t.Errorf(
			"expected 2024-01-15, got %v",
			got,
		)
	}

	// Clear date
	c.SetDT(time.Time{})
	if !c.DT().IsZero() {
		t.Error(
			"expected DT() to be zero after clear",
		)
	}
}

func TestThreadedCommentDone(t *testing.T) {
	c := NewThreadedComment()

	// Initially not done
	if c.Done() {
		t.Error(
			"expected Done() to be false initially",
		)
	}

	// Set done
	c.SetDone(true)
	if !c.Done() {
		t.Error("expected Done() to be true")
	}

	// Clear done
	c.SetDone(false)
	if c.Done() {
		t.Error("expected Done() to be false")
	}
}

func TestThreadedCommentText(t *testing.T) {
	c := NewThreadedComment()

	// Initially no text
	if c.Text() != nil {
		t.Error(
			"expected Text() to be nil initially",
		)
	}

	if c.TextContent() != "" {
		t.Error(
			"expected TextContent() to be empty initially",
		)
	}

	// Set text content
	c.SetTextContent(
		"This is the comment content",
	)

	if c.Text() == nil {
		t.Fatal(
			"expected Text() to be non-nil after SetTextContent",
		)
	}

	if c.TextContent() != "This is the comment content" {
		t.Errorf(
			"expected 'This is the comment content', got '%s'",
			c.TextContent(),
		)
	}

	// Update text content
	c.SetTextContent("Updated content")
	if c.TextContent() != "Updated content" {
		t.Errorf(
			"expected 'Updated content', got '%s'",
			c.TextContent(),
		)
	}
}

func TestThreadedCommentTextElement(
	t *testing.T,
) {
	text := NewThreadedCommentText()

	if text == nil {
		t.Fatal(
			"NewThreadedCommentText returned nil",
		)
	}

	if text.LocalName() != "text" {
		t.Errorf(
			"expected local name 'text', got '%s'",
			text.LocalName(),
		)
	}

	if text.NamespaceURI() != NamespaceThreadedComments {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceThreadedComments,
			text.NamespaceURI(),
		)
	}

	// Set and get text
	text.SetText("Hello")
	if text.Text() != "Hello" {
		t.Errorf(
			"expected 'Hello', got '%s'",
			text.Text(),
		)
	}
}

func TestThreadedCommentMentions(t *testing.T) {
	c := NewThreadedComment()
	c.SetTextContent("Hey @John, please review")

	// Add mention
	mention := c.AddMention("{person-john}", 4, 5)

	if mention == nil {
		t.Fatal("AddMention returned nil")
	}

	// Count mentions
	count := 0
	for range c.Mentions() {
		count++
	}

	if count != 1 {
		t.Errorf(
			"expected 1 mention, got %d",
			count,
		)
	}

	// Verify mention properties
	if mention.MentionPersonID() != "{person-john}" {
		t.Error(
			"expected mentionpersonId '{person-john}'",
		)
	}

	if mention.StartIndex() != 4 {
		t.Errorf(
			"expected startIndex 4, got %d",
			mention.StartIndex(),
		)
	}

	if mention.Length() != 5 {
		t.Errorf(
			"expected length 5, got %d",
			mention.Length(),
		)
	}
}

func TestMentionElement(t *testing.T) {
	m := NewMention()

	if m == nil {
		t.Fatal("NewMention returned nil")
	}

	if m.LocalName() != "mention" {
		t.Errorf(
			"expected local name 'mention', got '%s'",
			m.LocalName(),
		)
	}

	// Set and get properties
	m.SetMentionPersonID("{person-123}")
	if m.MentionPersonID() != "{person-123}" {
		t.Errorf(
			"expected '{person-123}', got '%s'",
			m.MentionPersonID(),
		)
	}

	m.SetStartIndex(10)
	if m.StartIndex() != 10 {
		t.Errorf(
			"expected 10, got %d",
			m.StartIndex(),
		)
	}

	m.SetLength(8)
	if m.Length() != 8 {
		t.Errorf("expected 8, got %d", m.Length())
	}
}

func TestPersonListCreation(t *testing.T) {
	pl := NewPersonList()

	if pl == nil {
		t.Fatal("NewPersonList returned nil")
	}

	if pl.LocalName() != "personList" {
		t.Errorf(
			"expected local name 'personList', got '%s'",
			pl.LocalName(),
		)
	}

	if pl.NamespaceURI() != NamespacePersons {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespacePersons,
			pl.NamespaceURI(),
		)
	}

	// Initially empty
	if pl.PersonCount() != 0 {
		t.Errorf(
			"expected 0 persons, got %d",
			pl.PersonCount(),
		)
	}
}

func TestPersonListAddAndGet(t *testing.T) {
	pl := NewPersonList()

	// Add persons
	p1 := pl.AddPerson()
	p1.SetID("{person-1}")
	p1.SetDisplayName("John Doe")
	p1.SetUserID("john@example.com")

	p2 := pl.AddPerson()
	p2.SetID("{person-2}")
	p2.SetDisplayName("Jane Smith")
	p2.SetUserID("jane@example.com")

	// Check count
	if pl.PersonCount() != 2 {
		t.Errorf(
			"expected 2 persons, got %d",
			pl.PersonCount(),
		)
	}

	// Get by index
	if pl.GetPerson(0).ID() != "{person-1}" {
		t.Error(
			"expected ID '{person-1}' at index 0",
		)
	}

	if pl.GetPerson(1).ID() != "{person-2}" {
		t.Error(
			"expected ID '{person-2}' at index 1",
		)
	}

	// Get by ID
	person := pl.GetPersonByID("{person-1}")
	if person == nil {
		t.Fatal("GetPersonByID returned nil")
	}

	if person.DisplayName() != "John Doe" {
		t.Errorf(
			"expected 'John Doe', got '%s'",
			person.DisplayName(),
		)
	}

	// Get by user ID
	person = pl.GetPersonByUserID(
		"jane@example.com",
	)
	if person == nil {
		t.Fatal("GetPersonByUserID returned nil")
	}

	if person.DisplayName() != "Jane Smith" {
		t.Errorf(
			"expected 'Jane Smith', got '%s'",
			person.DisplayName(),
		)
	}

	// Unknown ID
	if pl.GetPersonByID("{unknown}") != nil {
		t.Error("expected nil for unknown ID")
	}

	// Unknown user ID
	if pl.GetPersonByUserID(
		"unknown@example.com",
	) != nil {
		t.Error(
			"expected nil for unknown user ID",
		)
	}

	// Invalid indices
	if pl.GetPerson(-1) != nil {
		t.Error("expected nil for negative index")
	}

	if pl.GetPerson(10) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}
}

func TestPersonListRemove(t *testing.T) {
	pl := NewPersonList()

	p1 := pl.AddPerson()
	p1.SetID("{person-1}")

	p2 := pl.AddPerson()
	p2.SetID("{person-2}")

	p3 := pl.AddPerson()
	p3.SetID("{person-3}")

	if pl.PersonCount() != 3 {
		t.Fatalf(
			"expected 3 persons, got %d",
			pl.PersonCount(),
		)
	}

	// Remove middle person
	removed := pl.RemovePerson("{person-2}")
	if !removed {
		t.Error(
			"expected RemovePerson to return true",
		)
	}

	if pl.PersonCount() != 2 {
		t.Errorf(
			"expected 2 persons after remove, got %d",
			pl.PersonCount(),
		)
	}

	if pl.GetPersonByID("{person-2}") != nil {
		t.Error("expected person to be removed")
	}

	// Remove non-existent
	removed = pl.RemovePerson("{unknown}")
	if removed {
		t.Error(
			"expected RemovePerson to return false for unknown ID",
		)
	}
}

func TestPersonElement(t *testing.T) {
	p := NewPerson()

	if p == nil {
		t.Fatal("NewPerson returned nil")
	}

	if p.LocalName() != "person" {
		t.Errorf(
			"expected local name 'person', got '%s'",
			p.LocalName(),
		)
	}

	// Set and get ID
	p.SetID("{test-person}")
	if p.ID() != "{test-person}" {
		t.Errorf(
			"expected ID '{test-person}', got '%s'",
			p.ID(),
		)
	}

	// Set and get displayName
	p.SetDisplayName("Test User")
	if p.DisplayName() != "Test User" {
		t.Errorf(
			"expected 'Test User', got '%s'",
			p.DisplayName(),
		)
	}

	// Clear displayName
	p.SetDisplayName("")
	if p.DisplayName() != "" {
		t.Error(
			"expected displayName to be cleared",
		)
	}

	// Set and get userId
	p.SetUserID("test@example.com")
	if p.UserID() != "test@example.com" {
		t.Errorf(
			"expected 'test@example.com', got '%s'",
			p.UserID(),
		)
	}

	// Clear userId
	p.SetUserID("")
	if p.UserID() != "" {
		t.Error("expected userId to be cleared")
	}

	// Set and get providerId
	p.SetProviderID("Microsoft")
	if p.ProviderID() != "Microsoft" {
		t.Errorf(
			"expected 'Microsoft', got '%s'",
			p.ProviderID(),
		)
	}

	// Clear providerId
	p.SetProviderID("")
	if p.ProviderID() != "" {
		t.Error(
			"expected providerId to be cleared",
		)
	}
}

func TestThreadedCommentsClone(t *testing.T) {
	tc := NewThreadedComments()

	c1 := tc.AddComment()
	c1.SetID("{comment-1}")
	c1.SetRef("A1")
	c1.SetTextContent("Test comment")

	c2 := tc.AddComment()
	c2.SetID("{comment-2}")
	c2.SetRef("A1")
	c2.SetParentID("{comment-1}")
	c2.SetTextContent("Reply")

	// Clone
	clonedResult := tc.Clone()
	cloned, ok := clonedResult.(*ThreadedComments)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *ThreadedComments",
			clonedResult,
		)
	}

	// Verify cloned data
	if cloned.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments in clone, got %d",
			cloned.CommentCount(),
		)
	}

	// Modifying clone should not affect original
	cloned.AddComment().SetID("{new-comment}")
	if tc.CommentCount() != 2 {
		t.Error("original should not be modified")
	}
}

func TestPersonListClone(t *testing.T) {
	pl := NewPersonList()

	p1 := pl.AddPerson()
	p1.SetID("{person-1}")
	p1.SetDisplayName("Test User")

	// Clone
	clonedResult := pl.Clone()
	cloned, ok := clonedResult.(*PersonList)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *PersonList",
			clonedResult,
		)
	}

	// Verify cloned data
	if cloned.PersonCount() != 1 {
		t.Errorf(
			"expected 1 person in clone, got %d",
			cloned.PersonCount(),
		)
	}

	// Modifying clone should not affect original
	cloned.AddPerson().SetID("{new-person}")
	if pl.PersonCount() != 1 {
		t.Error("original should not be modified")
	}
}

func TestThreadedCommentsXML(t *testing.T) {
	tc := NewThreadedComments()

	c := tc.AddComment()
	c.SetID("{test-id}")
	c.SetRef("A1")
	c.SetPersonID("{person-id}")
	c.SetTextContent("Test comment content")

	xml := tc.OuterXml()

	// Check XML contains expected elements
	if !strings.Contains(
		xml,
		"ThreadedComments",
	) {
		t.Error(
			"XML should contain ThreadedComments element",
		)
	}

	if !strings.Contains(xml, "threadedComment") {
		t.Error(
			"XML should contain threadedComment element",
		)
	}

	if !strings.Contains(xml, `id="{test-id}"`) {
		t.Error("XML should contain id attribute")
	}

	if !strings.Contains(xml, `ref="A1"`) {
		t.Error(
			"XML should contain ref attribute",
		)
	}

	if !strings.Contains(
		xml,
		`personId="{person-id}"`,
	) {
		t.Error(
			"XML should contain personId attribute",
		)
	}

	if !strings.Contains(
		xml,
		"Test comment content",
	) {
		t.Error("XML should contain comment text")
	}
}

func TestPersonListXML(t *testing.T) {
	pl := NewPersonList()

	p := pl.AddPerson()
	p.SetID("{test-person}")
	p.SetDisplayName("Test User")
	p.SetUserID("test@example.com")
	p.SetProviderID("Microsoft")

	xml := pl.OuterXml()

	// Check XML contains expected elements
	if !strings.Contains(xml, "personList") {
		t.Error(
			"XML should contain personList element",
		)
	}

	if !strings.Contains(xml, "person") {
		t.Error(
			"XML should contain person element",
		)
	}

	if !strings.Contains(
		xml,
		`id="{test-person}"`,
	) {
		t.Error("XML should contain id attribute")
	}

	if !strings.Contains(
		xml,
		`displayName="Test User"`,
	) {
		t.Error(
			"XML should contain displayName attribute",
		)
	}

	if !strings.Contains(
		xml,
		`userId="test@example.com"`,
	) {
		t.Error(
			"XML should contain userId attribute",
		)
	}

	if !strings.Contains(
		xml,
		`providerId="Microsoft"`,
	) {
		t.Error(
			"XML should contain providerId attribute",
		)
	}
}

func TestFullThreadedCommentsWorkflow(
	t *testing.T,
) {
	// Create person list
	personList := NewPersonList()

	alice := personList.AddPerson()
	alice.SetID("{person-alice}")
	alice.SetDisplayName("Alice Smith")
	alice.SetUserID("alice@example.com")

	bob := personList.AddPerson()
	bob.SetID("{person-bob}")
	bob.SetDisplayName("Bob Jones")
	bob.SetUserID("bob@example.com")

	// Create threaded comments
	comments := NewThreadedComments()

	// Alice starts a thread
	rootComment := comments.AddComment()
	rootComment.SetID("{comment-1}")
	rootComment.SetRef("A1")
	rootComment.SetPersonID("{person-alice}")
	rootComment.SetDT(time.Now())
	rootComment.SetTextContent(
		"@Bob, please review this cell",
	)

	// Add mention
	rootComment.AddMention("{person-bob}", 0, 4)

	// Bob replies
	reply := comments.AddComment()
	reply.SetID("{comment-2}")
	reply.SetRef("A1")
	reply.SetPersonID("{person-bob}")
	reply.SetParentID("{comment-1}")
	reply.SetDT(time.Now())
	reply.SetTextContent("Looks good to me!")

	// Alice marks as resolved
	rootComment.SetDone(true)

	// Verify structure
	if personList.PersonCount() != 2 {
		t.Errorf(
			"expected 2 persons, got %d",
			personList.PersonCount(),
		)
	}

	if comments.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments, got %d",
			comments.CommentCount(),
		)
	}

	roots := comments.GetRootComments()
	if len(roots) != 1 {
		t.Errorf(
			"expected 1 root comment, got %d",
			len(roots),
		)
	}

	replies := comments.GetReplies("{comment-1}")
	if len(replies) != 1 {
		t.Errorf(
			"expected 1 reply, got %d",
			len(replies),
		)
	}

	if !rootComment.Done() {
		t.Error(
			"expected root comment to be done",
		)
	}

	// Verify mention
	mentionCount := 0
	for range rootComment.Mentions() {
		mentionCount++
	}

	if mentionCount != 1 {
		t.Errorf(
			"expected 1 mention, got %d",
			mentionCount,
		)
	}
}

func TestPersonsIteration(t *testing.T) {
	pl := NewPersonList()

	ids := []string{
		"{person-1}",
		"{person-2}",
		"{person-3}",
	}
	for _, id := range ids {
		p := pl.AddPerson()
		p.SetID(id)
	}

	i := 0
	for person := range pl.Persons() {
		if person.ID() != ids[i] {
			t.Errorf(
				"iteration %d: expected ID '%s', got '%s'",
				i,
				ids[i],
				person.ID(),
			)
		}
		i++
	}

	if i != len(ids) {
		t.Errorf(
			"expected %d iterations, got %d",
			len(ids),
			i,
		)
	}
}

func TestThreadedCommentsIteration(t *testing.T) {
	tc := NewThreadedComments()

	ids := []string{
		"{comment-1}",
		"{comment-2}",
		"{comment-3}",
	}
	for _, id := range ids {
		c := tc.AddComment()
		c.SetID(id)
	}

	i := 0
	for comment := range tc.Comments() {
		if comment.ID() != ids[i] {
			t.Errorf(
				"iteration %d: expected ID '%s', got '%s'",
				i,
				ids[i],
				comment.ID(),
			)
		}
		i++
	}

	if i != len(ids) {
		t.Errorf(
			"expected %d iterations, got %d",
			len(ids),
			i,
		)
	}
}
