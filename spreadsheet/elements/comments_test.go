package elements

import (
	"strings"
	"testing"
)

const (
	testAuthorJohnDoe   = "John Doe"
	testAuthorJaneSmith = "Jane Smith"
)

func TestCommentsCreation(t *testing.T) {
	c := NewComments()

	if c == nil {
		t.Fatal("NewComments returned nil")
	}

	if c.LocalName() != "comments" {
		t.Errorf(
			"expected local name 'comments', got '%s'",
			c.LocalName(),
		)
	}

	if c.NamespaceURI() != NamespaceSML {
		t.Errorf(
			"expected namespace '%s', got '%s'",
			NamespaceSML,
			c.NamespaceURI(),
		)
	}

	// Initially no children
	if c.Authors() != nil {
		t.Error(
			"expected Authors() to be nil initially",
		)
	}

	if c.CommentList() != nil {
		t.Error(
			"expected CommentList() to be nil initially",
		)
	}
}

func TestCommentsWithAuthorsAndList(
	t *testing.T,
) {
	c := NewComments()

	// Create authors
	authors := c.GetOrCreateAuthors()
	if authors == nil {
		t.Fatal("GetOrCreateAuthors returned nil")
	}

	if c.Authors() != authors {
		t.Error(
			"Authors() should return the same authors element",
		)
	}

	// Create comment list
	cl := c.GetOrCreateCommentList()
	if cl == nil {
		t.Fatal(
			"GetOrCreateCommentList returned nil",
		)
	}

	if c.CommentList() != cl {
		t.Error(
			"CommentList() should return the same comment list element",
		)
	}

	// Calling GetOrCreate again should return existing elements
	if c.GetOrCreateAuthors() != authors {
		t.Error(
			"GetOrCreateAuthors should return existing element",
		)
	}

	if c.GetOrCreateCommentList() != cl {
		t.Error(
			"GetOrCreateCommentList should return existing element",
		)
	}
}

func TestAuthorsCreation(t *testing.T) {
	a := NewAuthors()

	if a == nil {
		t.Fatal("NewAuthors returned nil")
	}

	if a.LocalName() != "authors" {
		t.Errorf(
			"expected local name 'authors', got '%s'",
			a.LocalName(),
		)
	}

	// Initially empty
	if a.AuthorCount() != 0 {
		t.Errorf(
			"expected 0 authors, got %d",
			a.AuthorCount(),
		)
	}
}

func TestAuthorsAddAndGet(t *testing.T) {
	a := NewAuthors()

	// Add authors
	author1 := a.AddAuthor(testAuthorJohnDoe)
	if author1 == nil {
		t.Fatal("AddAuthor returned nil")
	}

	if author1.Name() != testAuthorJohnDoe {
		t.Errorf(
			"expected 'John Doe', got '%s'",
			author1.Name(),
		)
	}

	author2 := a.AddAuthor(testAuthorJaneSmith)
	if author2 == nil {
		t.Fatal("AddAuthor returned nil")
	}

	// Check count
	if a.AuthorCount() != 2 {
		t.Errorf(
			"expected 2 authors, got %d",
			a.AuthorCount(),
		)
	}

	// Get by index
	if a.GetAuthor(0).
		Name() !=
		testAuthorJohnDoe {
		t.Error("expected 'John Doe' at index 0")
	}

	if a.GetAuthor(1).
		Name() !=
		testAuthorJaneSmith {
		t.Error(
			"expected 'Jane Smith' at index 1",
		)
	}

	// Get by name
	if a.GetAuthorName(0) != testAuthorJohnDoe {
		t.Error(
			"expected GetAuthorName(0) to return 'John Doe'",
		)
	}

	if a.GetAuthorName(1) != testAuthorJaneSmith {
		t.Error(
			"expected GetAuthorName(1) to return 'Jane Smith'",
		)
	}

	// Invalid indices
	if a.GetAuthor(-1) != nil {
		t.Error("expected nil for negative index")
	}

	if a.GetAuthor(10) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}

	if a.GetAuthorName(10) != "" {
		t.Error(
			"expected empty string for out of range index",
		)
	}
}

func TestAuthorsFindIndex(t *testing.T) {
	a := NewAuthors()

	a.AddAuthor("Alice")
	a.AddAuthor("Bob")
	a.AddAuthor("Charlie")

	if a.FindAuthorIndex("Alice") != 0 {
		t.Error("expected index 0 for 'Alice'")
	}

	if a.FindAuthorIndex("Bob") != 1 {
		t.Error("expected index 1 for 'Bob'")
	}

	if a.FindAuthorIndex("Charlie") != 2 {
		t.Error("expected index 2 for 'Charlie'")
	}

	if a.FindAuthorIndex("Unknown") != -1 {
		t.Error("expected -1 for unknown author")
	}
}

func TestAuthorsGetOrAdd(t *testing.T) {
	a := NewAuthors()

	// First call should add
	idx1 := a.GetOrAddAuthor("First Author")
	if idx1 != 0 {
		t.Errorf("expected index 0, got %d", idx1)
	}

	// Second call should return existing
	idx2 := a.GetOrAddAuthor("First Author")
	if idx2 != 0 {
		t.Errorf(
			"expected index 0 for existing author, got %d",
			idx2,
		)
	}

	// Different author should get new index
	idx3 := a.GetOrAddAuthor("Second Author")
	if idx3 != 1 {
		t.Errorf(
			"expected index 1 for new author, got %d",
			idx3,
		)
	}

	// Count should be 2
	if a.AuthorCount() != 2 {
		t.Errorf(
			"expected 2 authors, got %d",
			a.AuthorCount(),
		)
	}
}

func TestAuthorElement(t *testing.T) {
	author := NewAuthor()

	if author == nil {
		t.Fatal("NewAuthor returned nil")
	}

	if author.LocalName() != "author" {
		t.Errorf(
			"expected local name 'author', got '%s'",
			author.LocalName(),
		)
	}

	// Initially empty
	if author.Name() != "" {
		t.Errorf(
			"expected empty name, got '%s'",
			author.Name(),
		)
	}

	// Set name
	author.SetName("Test Author")
	if author.Name() != "Test Author" {
		t.Errorf(
			"expected 'Test Author', got '%s'",
			author.Name(),
		)
	}
}

func TestCommentListCreation(t *testing.T) {
	cl := NewCommentList()

	if cl == nil {
		t.Fatal("NewCommentList returned nil")
	}

	if cl.LocalName() != "commentList" {
		t.Errorf(
			"expected local name 'commentList', got '%s'",
			cl.LocalName(),
		)
	}

	// Initially empty
	if cl.CommentCount() != 0 {
		t.Errorf(
			"expected 0 comments, got %d",
			cl.CommentCount(),
		)
	}
}

func TestCommentListAddAndGet(t *testing.T) {
	cl := NewCommentList()

	// Add comments
	c1 := cl.AddComment("A1", 0)
	if c1 == nil {
		t.Fatal("AddComment returned nil")
	}

	if c1.Ref() != "A1" {
		t.Errorf(
			"expected ref 'A1', got '%s'",
			c1.Ref(),
		)
	}

	if c1.AuthorId() != 0 {
		t.Errorf(
			"expected authorId 0, got %d",
			c1.AuthorId(),
		)
	}

	c2 := cl.AddComment("B2", 1)

	// Check count
	if cl.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments, got %d",
			cl.CommentCount(),
		)
	}

	// Get by index
	if cl.GetComment(0).Ref() != "A1" {
		t.Error("expected ref 'A1' at index 0")
	}

	if cl.GetComment(1).Ref() != "B2" {
		t.Error("expected ref 'B2' at index 1")
	}

	// Get by ref
	comment := cl.GetCommentByRef("A1")
	if comment == nil {
		t.Fatal("GetCommentByRef returned nil")
	}

	if comment.Ref() != "A1" {
		t.Error("expected ref 'A1'")
	}

	// Unknown ref
	if cl.GetCommentByRef("Z99") != nil {
		t.Error("expected nil for unknown ref")
	}

	// Invalid indices
	if cl.GetComment(-1) != nil {
		t.Error("expected nil for negative index")
	}

	if cl.GetComment(10) != nil {
		t.Error(
			"expected nil for out of range index",
		)
	}

	// Don't use c2 directly to avoid unused variable warning
	_ = c2
}

func TestCommentListRemove(t *testing.T) {
	cl := NewCommentList()

	cl.AddComment("A1", 0)
	cl.AddComment("B2", 1)
	cl.AddComment("C3", 2)

	if cl.CommentCount() != 3 {
		t.Fatalf(
			"expected 3 comments, got %d",
			cl.CommentCount(),
		)
	}

	// Remove middle comment
	removed := cl.RemoveComment("B2")
	if !removed {
		t.Error(
			"expected RemoveComment to return true",
		)
	}

	if cl.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments after remove, got %d",
			cl.CommentCount(),
		)
	}

	if cl.GetCommentByRef("B2") != nil {
		t.Error("expected comment to be removed")
	}

	// Remove non-existent
	removed = cl.RemoveComment("Z99")
	if removed {
		t.Error(
			"expected RemoveComment to return false for unknown ref",
		)
	}
}

func TestCommentElement(t *testing.T) {
	c := NewComment()

	if c == nil {
		t.Fatal("NewComment returned nil")
	}

	if c.LocalName() != "comment" {
		t.Errorf(
			"expected local name 'comment', got '%s'",
			c.LocalName(),
		)
	}

	// Set and get ref
	c.SetRef("C5")
	if c.Ref() != "C5" {
		t.Errorf(
			"expected ref 'C5', got '%s'",
			c.Ref(),
		)
	}

	// Set and get authorId
	c.SetAuthorId(3)
	if c.AuthorId() != 3 {
		t.Errorf(
			"expected authorId 3, got %d",
			c.AuthorId(),
		)
	}

	// Set and get GUID
	c.SetGUID(
		"{12345678-1234-1234-1234-123456789012}",
	)
	if c.GUID() != "{12345678-1234-1234-1234-123456789012}" {
		t.Error("expected GUID to be set")
	}

	// Clear GUID
	c.SetGUID("")
	if c.GUID() != "" {
		t.Error("expected GUID to be cleared")
	}

	// Set and get shapeId
	c.SetShapeId(100)
	if c.ShapeId() != 100 {
		t.Errorf(
			"expected shapeId 100, got %d",
			c.ShapeId(),
		)
	}

	// Clear shapeId
	c.SetShapeId(0)
	if c.ShapeId() != 0 {
		t.Error("expected shapeId to be 0")
	}
}

func TestCommentText(t *testing.T) {
	c := NewComment()

	// Initially no text
	if c.Text() != nil {
		t.Error(
			"expected Text() to be nil initially",
		)
	}

	if c.PlainText() != "" {
		t.Error(
			"expected PlainText() to be empty initially",
		)
	}

	// Set plain text
	c.SetPlainText("This is a comment")

	if c.Text() == nil {
		t.Fatal(
			"expected Text() to be non-nil after SetPlainText",
		)
	}

	if c.PlainText() != "This is a comment" {
		t.Errorf(
			"expected 'This is a comment', got '%s'",
			c.PlainText(),
		)
	}

	// Update plain text
	c.SetPlainText("Updated comment")

	if c.PlainText() != "Updated comment" {
		t.Errorf(
			"expected 'Updated comment', got '%s'",
			c.PlainText(),
		)
	}
}

func TestCommentTextElement(t *testing.T) {
	ct := NewCommentText()

	if ct == nil {
		t.Fatal("NewCommentText returned nil")
	}

	if ct.LocalName() != "text" {
		t.Errorf(
			"expected local name 'text', got '%s'",
			ct.LocalName(),
		)
	}

	// Initially empty
	if ct.PlainText() != "" {
		t.Error(
			"expected PlainText() to be empty",
		)
	}

	if ct.IsRichText() {
		t.Error(
			"expected IsRichText() to be false",
		)
	}
}

func TestCommentTextRichFormatting(t *testing.T) {
	ct := NewCommentText()

	// Add rich text runs
	run1 := ct.AddRichTextRun()
	run1.SetTextContent("Bold text")
	run1.GetOrCreateRunProperties().SetBold(true)

	run2 := ct.AddRichTextRun()
	run2.SetTextContent(" Normal text")

	// Should be rich text
	if !ct.IsRichText() {
		t.Error(
			"expected IsRichText() to be true",
		)
	}

	// Plain text should concatenate
	if ct.PlainText() != "Bold text Normal text" {
		t.Errorf(
			"expected 'Bold text Normal text', got '%s'",
			ct.PlainText(),
		)
	}

	// Count runs
	count := 0
	for range ct.RichTextRuns() {
		count++
	}

	if count != 2 {
		t.Errorf(
			"expected 2 rich text runs, got %d",
			count,
		)
	}
}

func TestCommentTextAddRichTextRunWithText(
	t *testing.T,
) {
	ct := NewCommentText()

	run := ct.AddRichTextRunWithText("Quick text")

	if run == nil {
		t.Fatal(
			"AddRichTextRunWithText returned nil",
		)
	}

	if run.TextContent() != "Quick text" {
		t.Errorf(
			"expected 'Quick text', got '%s'",
			run.TextContent(),
		)
	}

	if !ct.IsRichText() {
		t.Error(
			"expected IsRichText() to be true",
		)
	}
}

func TestCommentTextSwitchToPlain(t *testing.T) {
	ct := NewCommentText()

	// Start with rich text
	ct.AddRichTextRunWithText("Part 1")
	ct.AddRichTextRunWithText(" Part 2")

	if !ct.IsRichText() {
		t.Fatal("expected rich text")
	}

	// Switch to plain text
	ct.SetPlainText("Plain text now")

	// Rich text runs should be removed
	if ct.IsRichText() {
		t.Error(
			"expected IsRichText() to be false after SetPlainText",
		)
	}

	if ct.PlainText() != "Plain text now" {
		t.Errorf(
			"expected 'Plain text now', got '%s'",
			ct.PlainText(),
		)
	}
}

func TestCommentsClone(t *testing.T) {
	c := NewComments()

	authors := c.GetOrCreateAuthors()
	authors.AddAuthor("Test Author")

	cl := c.GetOrCreateCommentList()
	comment := cl.AddComment("A1", 0)
	comment.SetPlainText("Test comment")

	// Clone
	clonedResult := c.Clone()
	cloned, ok := clonedResult.(*Comments)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Comments",
			clonedResult,
		)
	}

	// Verify cloned data
	if cloned.Authors() == nil {
		t.Fatal("expected Authors() in clone")
	}

	if cloned.Authors().AuthorCount() != 1 {
		t.Errorf(
			"expected 1 author in clone, got %d",
			cloned.Authors().AuthorCount(),
		)
	}

	if cloned.CommentList() == nil {
		t.Fatal("expected CommentList() in clone")
	}

	if cloned.CommentList().CommentCount() != 1 {
		t.Errorf(
			"expected 1 comment in clone, got %d",
			cloned.CommentList().CommentCount(),
		)
	}

	// Modifying clone should not affect original
	cloned.GetOrCreateAuthors().
		AddAuthor("New Author")
	if c.Authors().AuthorCount() != 1 {
		t.Error("original should not be modified")
	}
}

func TestCommentClone(t *testing.T) {
	comment := NewComment()
	comment.SetRef("B5")
	comment.SetAuthorId(2)
	comment.SetPlainText("Original comment")

	clonedResult := comment.Clone()
	cloned, ok := clonedResult.(*Comment)
	if !ok {
		t.Fatalf(
			"Clone() returned %T, expected *Comment",
			clonedResult,
		)
	}

	if cloned.Ref() != "B5" {
		t.Errorf(
			"expected ref 'B5', got '%s'",
			cloned.Ref(),
		)
	}

	if cloned.AuthorId() != 2 {
		t.Errorf(
			"expected authorId 2, got %d",
			cloned.AuthorId(),
		)
	}

	if cloned.PlainText() != "Original comment" {
		t.Errorf(
			"expected 'Original comment', got '%s'",
			cloned.PlainText(),
		)
	}

	// Modifying clone should not affect original
	cloned.SetPlainText("Modified comment")
	if comment.PlainText() == "Modified comment" {
		t.Error("original should not be modified")
	}
}

func TestCommentsXML(t *testing.T) {
	c := NewComments()

	authors := c.GetOrCreateAuthors()
	authors.AddAuthor("Test Author")

	cl := c.GetOrCreateCommentList()
	comment := cl.AddComment("A1", 0)
	comment.SetPlainText("Test comment text")

	xml := c.OuterXml()

	// Check XML contains expected elements
	if !strings.Contains(xml, "comments") {
		t.Error(
			"XML should contain comments element",
		)
	}

	if !strings.Contains(xml, "authors") {
		t.Error(
			"XML should contain authors element",
		)
	}

	if !strings.Contains(
		xml,
		"Test Author",
	) {
		t.Error(
			"XML should contain author content",
		)
	}

	if !strings.Contains(xml, "commentList") {
		t.Error(
			"XML should contain commentList element",
		)
	}

	if !strings.Contains(xml, `ref="A1"`) {
		t.Error(
			"XML should contain ref attribute",
		)
	}

	if !strings.Contains(xml, `authorId="0"`) {
		t.Error(
			"XML should contain authorId attribute",
		)
	}

	if !strings.Contains(xml, "text") {
		t.Error(
			"XML should contain text element",
		)
	}

	if !strings.Contains(
		xml,
		"Test comment text",
	) {
		t.Error("XML should contain comment text")
	}
}

func TestAuthorsIteration(t *testing.T) {
	a := NewAuthors()

	expected := []string{
		"Author1",
		"Author2",
		"Author3",
	}
	for _, name := range expected {
		a.AddAuthor(name)
	}

	i := 0
	for author := range a.AuthorList() {
		if author.Name() != expected[i] {
			t.Errorf(
				"iteration %d: expected '%s', got '%s'",
				i,
				expected[i],
				author.Name(),
			)
		}
		i++
	}

	if i != len(expected) {
		t.Errorf(
			"expected %d iterations, got %d",
			len(expected),
			i,
		)
	}
}

func TestCommentsIteration(t *testing.T) {
	cl := NewCommentList()

	refs := []string{"A1", "B2", "C3"}
	for i, ref := range refs {
		cl.AddComment(ref, i)
	}

	i := 0
	for comment := range cl.Comments() {
		if comment.Ref() != refs[i] {
			t.Errorf(
				"iteration %d: expected ref '%s', got '%s'",
				i,
				refs[i],
				comment.Ref(),
			)
		}
		i++
	}

	if i != len(refs) {
		t.Errorf(
			"expected %d iterations, got %d",
			len(refs),
			i,
		)
	}
}

func TestFullCommentsWorkflow(t *testing.T) {
	// Create comments structure
	comments := NewComments()

	// Add authors
	authors := comments.GetOrCreateAuthors()
	authors.AddAuthor("Alice")
	authors.AddAuthor("Bob")

	// Get comment list
	commentList := comments.GetOrCreateCommentList()

	// Add a comment from Alice
	c1 := commentList.AddComment("A1", 0)
	c1.SetPlainText(
		"This is Alice's comment on A1",
	)

	// Add a comment from Bob with rich text
	c2 := commentList.AddComment("B2", 1)
	text := c2.GetOrCreateText()
	run1 := text.AddRichTextRun()
	run1.SetTextContent("Important: ")
	run1.GetOrCreateRunProperties().SetBold(true)
	run2 := text.AddRichTextRun()
	run2.SetTextContent("This needs review")

	// Verify structure
	if authors.AuthorCount() != 2 {
		t.Errorf(
			"expected 2 authors, got %d",
			authors.AuthorCount(),
		)
	}

	if commentList.CommentCount() != 2 {
		t.Errorf(
			"expected 2 comments, got %d",
			commentList.CommentCount(),
		)
	}

	// Verify Alice's comment
	aliceComment := commentList.GetCommentByRef(
		"A1",
	)
	if aliceComment == nil {
		t.Fatal(
			"expected to find Alice's comment",
		)
	}

	if aliceComment.AuthorId() != 0 {
		t.Error("expected authorId 0 for Alice")
	}

	if aliceComment.PlainText() != "This is Alice's comment on A1" {
		t.Errorf(
			"unexpected comment text: %s",
			aliceComment.PlainText(),
		)
	}

	// Verify Bob's comment
	bobComment := commentList.GetCommentByRef(
		"B2",
	)
	if bobComment == nil {
		t.Fatal("expected to find Bob's comment")
	}

	if bobComment.AuthorId() != 1 {
		t.Error("expected authorId 1 for Bob")
	}

	if !bobComment.Text().IsRichText() {
		t.Error(
			"expected Bob's comment to be rich text",
		)
	}

	if bobComment.PlainText() != "Important: This needs review" {
		t.Errorf(
			"unexpected comment text: %s",
			bobComment.PlainText(),
		)
	}
}
