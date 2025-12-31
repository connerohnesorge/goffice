package wordprocessing_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const (
	commentRangeStartName = "commentRangeStart"
	commentRangeEndName   = "commentRangeEnd"
)

func TestParagraphMarkCommentRange(t *testing.T) {
	t.Run("ValidRange", func(t *testing.T) {
		// Create a paragraph with multiple runs
		para := elements.NewParagraph()
		para.AppendRun("First ")
		para.AppendRun("Second ")
		para.AppendRun("Third")

		// Mark comment range from run 0 to run 2
		err := para.MarkCommentRange(0, 2, 1)
		if err != nil {
			t.Fatalf(
				"MarkCommentRange failed: %v",
				err,
			)
		}

		// Verify CommentRangeStart was inserted before first run
		children := make(
			[]string,
			0,
		)
		for child := range para.Children() {
			children = append(
				children,
				child.LocalName(),
			)
		}

		// Expected: [commentRangeStart, r, r, r, commentRangeEnd]
		if len(children) < 5 {
			t.Fatalf(
				"Expected at least 5 children, got %d",
				len(children),
			)
		}

		if children[0] != commentRangeStartName {
			t.Errorf(
				"Expected first child to be commentRangeStart, got %s",
				children[0],
			)
		}

		if children[4] != commentRangeEndName {
			t.Errorf(
				"Expected fifth child to be commentRangeEnd, got %s",
				children[4],
			)
		}

		// Verify CommentReference was added to the last run
		runs := make([]*elements.Run, 0)
		for run := range para.Runs() {
			runs = append(runs, run)
		}

		lastRun := runs[len(runs)-1]
		hasCommentRef := false
		for child := range lastRun.Children() {
			if child.LocalName() == "commentReference" {
				hasCommentRef = true

				break
			}
		}

		if !hasCommentRef {
			t.Error(
				"Expected last run to contain commentReference",
			)
		}
	})

	t.Run("SingleRun", func(t *testing.T) {
		// Create a paragraph with a single run
		para := elements.NewParagraph(
			"Single run",
		)

		// Mark comment range on the single run
		err := para.MarkCommentRange(0, 0, 2)
		if err != nil {
			t.Fatalf(
				"MarkCommentRange failed: %v",
				err,
			)
		}

		// Count children
		childCount := 0
		for range para.Children() {
			childCount++
		}

		// Expected: [commentRangeStart, r, commentRangeEnd]
		if childCount != 3 {
			t.Errorf(
				"Expected 3 children, got %d",
				childCount,
			)
		}
	})

	t.Run(
		"InvalidIndices_Negative",
		func(t *testing.T) {
			para := elements.NewParagraph("Test")

			err := para.MarkCommentRange(
				-1,
				0,
				1,
			)
			if err == nil {
				t.Error(
					"Expected error for negative startRunIdx",
				)
			}
		},
	)

	t.Run(
		"InvalidIndices_OutOfBounds",
		func(t *testing.T) {
			para := elements.NewParagraph("Test")

			err := para.MarkCommentRange(0, 5, 1)
			if err == nil {
				t.Error(
					"Expected error for out of bounds endRunIdx",
				)
			}
		},
	)

	t.Run(
		"InvalidIndices_StartGreaterThanEnd",
		func(t *testing.T) {
			para := elements.NewParagraph()
			para.AppendRun("First")
			para.AppendRun("Second")

			err := para.MarkCommentRange(1, 0, 1)
			if err == nil {
				t.Error(
					"Expected error for startRunIdx > endRunIdx",
				)
			}
		},
	)

	t.Run("NoRuns", func(t *testing.T) {
		para := elements.NewParagraph()

		err := para.MarkCommentRange(0, 0, 1)
		if err == nil {
			t.Error(
				"Expected error for paragraph with no runs",
			)
		}
	})
}

//nolint:revive // cyclomatic: test function with multiple subtests
func TestDocumentApplyComment(t *testing.T) {
	t.Run(
		"CreateCommentAndMarkRange",
		func(t *testing.T) {
			// Create a new document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Get the main part and add a paragraph
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			para := elements.NewParagraph()
			para.AppendRun("First ")
			para.AppendRun("Second ")
			para.AppendRun("Third")
			body.AppendChild(para)

			// Apply a comment
			comment, err := doc.ApplyComment(
				para,
				0,
				2,
				"John Doe",
				"This is a test comment",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment failed: %v",
					err,
				)
			}

			if comment == nil {
				t.Fatal(
					"Expected comment to be created",
				)
			}

			// Verify comment properties
			if comment.Author() != "John Doe" {
				t.Errorf(
					"Expected author 'John Doe', got '%s'",
					comment.Author(),
				)
			}

			// Verify comment ID
			commentID := comment.Id()
			if commentID == 0 {
				t.Error(
					"Expected non-zero comment ID",
				)
			}

			// Verify the comment part was created
			commentsPart, err := doc.CommentsPart()
			if err != nil {
				t.Fatalf(
					"Failed to get comments part: %v",
					err,
				)
			}

			if commentsPart == nil {
				t.Fatal(
					"Expected comments part to be created",
				)
			}

			// Verify the comment is in the comments part
			retrievedComment := commentsPart.GetComment(
				commentID,
			)
			if retrievedComment == nil {
				t.Error(
					"Expected comment to be in comments part",
				)
			}
		},
	)

	t.Run("SingleRunComment", func(t *testing.T) {
		// Create a new document
		doc, err := wordprocessing.New(
			t.TempDir()+"/test2.docx",
			wordprocessing.DocTypeDocument,
		)
		if err != nil {
			t.Fatalf(
				"Failed to create document: %v",
				err,
			)
		}
		defer func() { _ = doc.Close() }()

		// Get the main part and add a paragraph
		mainPart := doc.MainPart()
		docElem := mainPart.Document()
		body := docElem.GetOrCreateBody()

		para := elements.NewParagraph(
			"Single run",
		)
		body.AppendChild(para)

		// Apply a comment to the single run
		comment, err := doc.ApplyComment(
			para,
			0,
			0,
			"Jane Smith",
			"Comment on single run",
		)
		if err != nil {
			t.Fatalf(
				"ApplyComment failed: %v",
				err,
			)
		}

		if comment == nil {
			t.Fatal(
				"Expected comment to be created",
			)
		}
	})

	t.Run(
		"CreatesCommentsPartIfNotExists",
		func(t *testing.T) {
			// Create a new document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test3.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Verify comments part doesn't exist yet
			mainPart := doc.MainPart()
			if mainPart.CommentsPart() != nil {
				t.Fatal(
					"Expected no comments part initially",
				)
			}

			// Add a paragraph
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()
			para := elements.NewParagraph("Test")
			body.AppendChild(para)

			// Apply a comment - should create the comments part
			_, err = doc.ApplyComment(
				para,
				0,
				0,
				"Author",
				"Comment",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment failed: %v",
					err,
				)
			}

			// Verify comments part was created
			if mainPart.CommentsPart() == nil {
				t.Error(
					"Expected comments part to be created",
				)
			}
		},
	)

	t.Run(
		"MultipleCommentsOnSameParagraph",
		func(t *testing.T) {
			// Create a new document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test4.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add a paragraph with multiple runs
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			para := elements.NewParagraph()
			para.AppendRun("First ")
			para.AppendRun("Second ")
			para.AppendRun("Third ")
			para.AppendRun("Fourth")
			body.AppendChild(para)

			// Apply first comment to runs 0-1
			comment1, err := doc.ApplyComment(
				para,
				0,
				1,
				"Author1",
				"Comment on first and second",
			)
			if err != nil {
				t.Fatalf(
					"First ApplyComment failed: %v",
					err,
				)
			}

			// Apply second comment to runs 2-3
			comment2, err := doc.ApplyComment(
				para,
				2,
				3,
				"Author2",
				"Comment on third and fourth",
			)
			if err != nil {
				t.Fatalf(
					"Second ApplyComment failed: %v",
					err,
				)
			}

			// Verify both comments have different IDs
			if comment1.Id() == comment2.Id() {
				t.Error(
					"Expected different comment IDs",
				)
			}

			// Verify both comments are in the comments part
			commentsPart := mainPart.CommentsPart()
			if commentsPart == nil {
				t.Fatal(
					"Expected comments part to exist",
				)
			}

			if commentsPart.GetComment(
				comment1.Id(),
			) == nil {
				t.Error(
					"Expected first comment in comments part",
				)
			}
			if commentsPart.GetComment(
				comment2.Id(),
			) == nil {
				t.Error(
					"Expected second comment in comments part",
				)
			}
		},
	)

	t.Run(
		"InvalidRange_RollsBackComment",
		func(t *testing.T) {
			// Create a new document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test5.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add a paragraph
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()
			para := elements.NewParagraph("Test")
			body.AppendChild(para)

			// Try to apply comment with invalid range
			comment, err := doc.ApplyComment(
				para,
				0,
				5,
				"Author",
				"Comment",
			)
			if err == nil {
				t.Error(
					"Expected error for invalid range",
				)
			}
			if comment != nil {
				t.Error(
					"Expected no comment to be returned on error",
				)
			}

			// Verify the comment was rolled back (not in comments part)
			commentsPart := mainPart.CommentsPart()
			if commentsPart != nil {
				// If the comments part was created, verify it has no comments
				commentCount := 0
				comments := commentsPart.GetOrCreateComments()
				for range comments.Comments() {
					commentCount++
				}
				if commentCount > 0 {
					t.Errorf(
						"Expected 0 comments after rollback, got %d",
						commentCount,
					)
				}
			}
		},
	)
}

func TestParagraphInsertChildBefore(
	t *testing.T,
) {
	t.Run("ValidInsert", func(t *testing.T) {
		para := elements.NewParagraph()
		run1 := para.AppendRun("First")
		para.AppendRun("Second")

		// Create a new run to insert
		newRun := elements.NewRun("Inserted")

		// Insert before run1
		err := para.InsertChildBefore(
			newRun,
			run1,
		)
		if err != nil {
			t.Fatalf(
				"InsertChildBefore failed: %v",
				err,
			)
		}

		// Verify order
		runs := make([]*elements.Run, 0)
		for run := range para.Runs() {
			runs = append(runs, run)
		}

		if len(runs) != 3 {
			t.Fatalf(
				"Expected 3 runs, got %d",
				len(runs),
			)
		}

		if runs[0].InnerText() != "Inserted" {
			t.Errorf(
				"Expected first run to be 'Inserted', got '%s'",
				runs[0].InnerText(),
			)
		}
	})

	t.Run("InvalidRefChild", func(t *testing.T) {
		para := elements.NewParagraph("Test")

		// Create a run that's not a child of para
		foreignRun := elements.NewRun("Foreign")

		// Try to insert before a non-child
		newRun := elements.NewRun("New")
		err := para.InsertChildBefore(
			newRun,
			foreignRun,
		)
		if err == nil {
			t.Error(
				"Expected error for non-child reference",
			)
		}
	})
}

func TestParagraphInsertChildAfter(t *testing.T) {
	t.Run("ValidInsert", func(t *testing.T) {
		para := elements.NewParagraph()
		run1 := para.AppendRun("First")
		para.AppendRun("Second")

		// Create a new run to insert
		newRun := elements.NewRun("Inserted")

		// Insert after run1
		err := para.InsertChildAfter(newRun, run1)
		if err != nil {
			t.Fatalf(
				"InsertChildAfter failed: %v",
				err,
			)
		}

		// Verify order
		runs := make([]*elements.Run, 0)
		for run := range para.Runs() {
			runs = append(runs, run)
		}

		if len(runs) != 3 {
			t.Fatalf(
				"Expected 3 runs, got %d",
				len(runs),
			)
		}

		if runs[1].InnerText() != "Inserted" {
			t.Errorf(
				"Expected second run to be 'Inserted', got '%s'",
				runs[1].InnerText(),
			)
		}
	})

	t.Run("InvalidRefChild", func(t *testing.T) {
		para := elements.NewParagraph("Test")

		// Create a run that's not a child of para
		foreignRun := elements.NewRun("Foreign")

		// Try to insert after a non-child
		newRun := elements.NewRun("New")
		err := para.InsertChildAfter(
			newRun,
			foreignRun,
		)
		if err == nil {
			t.Error(
				"Expected error for non-child reference",
			)
		}
	})
}

//nolint:revive // cyclomatic: test function with multiple subtests
func TestCommentsRemoveComment(t *testing.T) {
	t.Run(
		"RemoveExistingComment",
		func(t *testing.T) {
			// Create a Comments element with multiple comments
			comments := elements.NewComments()
			comment1 := comments.AddComment(
				"Author1",
				"Comment 1",
			)
			comment2 := comments.AddComment(
				"Author2",
				"Comment 2",
			)
			comment3 := comments.AddComment(
				"Author3",
				"Comment 3",
			)

			id2 := comment2.Id()

			// Remove the second comment
			removed := comments.RemoveComment(id2)
			if !removed {
				t.Error(
					"Expected RemoveComment to return true",
				)
			}

			// Verify the comment was removed
			if comments.GetComment(id2) != nil {
				t.Error(
					"Comment should have been removed",
				)
			}

			// Verify other comments still exist
			if comments.GetComment(
				comment1.Id(),
			) == nil {
				t.Error(
					"Comment 1 should still exist",
				)
			}
			if comments.GetComment(
				comment3.Id(),
			) == nil {
				t.Error(
					"Comment 3 should still exist",
				)
			}
		},
	)

	t.Run(
		"RemoveNonExistentComment",
		func(t *testing.T) {
			comments := elements.NewComments()
			comments.AddComment(
				"Author",
				"Comment",
			)

			// Try to remove a comment that doesn't exist
			removed := comments.RemoveComment(999)
			if removed {
				t.Error(
					"Expected RemoveComment to return false for non-existent ID",
				)
			}
		},
	)
}

func TestCommentsPartRemoveComment(t *testing.T) {
	t.Run(
		"RemoveExistingComment",
		func(t *testing.T) {
			// Create a document with comments
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add comments
			commentsPart, err := doc.CommentsPart()
			if err != nil {
				t.Fatalf(
					"Failed to get comments part: %v",
					err,
				)
			}

			comment1 := commentsPart.AddComment(
				"Author1",
				"Comment 1",
			)
			comment2 := commentsPart.AddComment(
				"Author2",
				"Comment 2",
			)

			id1 := comment1.Id()

			// Remove first comment
			removed := commentsPart.RemoveComment(
				id1,
			)
			if !removed {
				t.Error(
					"Expected RemoveComment to return true",
				)
			}

			// Verify it was removed
			if commentsPart.GetComment(
				id1,
			) != nil {
				t.Error(
					"Comment should have been removed",
				)
			}

			// Verify second comment still exists
			if commentsPart.GetComment(
				comment2.Id(),
			) == nil {
				t.Error(
					"Comment 2 should still exist",
				)
			}
		},
	)

	t.Run(
		"RemoveFromEmptyPart",
		func(t *testing.T) {
			// Create a document without comments part
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Try to remove from non-existent part (this will create it)
			commentsPart, _ := doc.CommentsPart()
			removed := commentsPart.RemoveComment(
				1,
			)
			if removed {
				t.Error(
					"Expected RemoveComment to return false when no comments exist",
				)
			}
		},
	)
}

//nolint:revive // cyclomatic: test function with multiple subtests
func TestDocumentRemoveComment(t *testing.T) {
	t.Run(
		"RemoveCommentWithMarkers",
		func(t *testing.T) {
			// Create a document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add a paragraph with runs
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			para := elements.NewParagraph()
			para.AppendRun("First ")
			para.AppendRun("Second ")
			para.AppendRun("Third")
			body.AppendChild(para)

			// Apply a comment
			comment, err := doc.ApplyComment(
				para,
				0,
				2,
				"Author",
				"Test comment",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment failed: %v",
					err,
				)
			}

			commentID := comment.Id()

			// Verify markers exist
			hasStart := false
			hasEnd := false
			for child := range para.Children() {
				if child.LocalName() == commentRangeStartName {
					hasStart = true
				}
				if child.LocalName() == commentRangeEndName {
					hasEnd = true
				}
			}
			if !hasStart || !hasEnd {
				t.Fatal(
					"Comment markers should exist before removal",
				)
			}

			// Remove the comment
			err = doc.RemoveComment(commentID)
			if err != nil {
				t.Fatalf(
					"RemoveComment failed: %v",
					err,
				)
			}

			// Verify comment was removed from comments part
			commentsPart := mainPart.CommentsPart()
			if commentsPart.GetComment(
				commentID,
			) != nil {
				t.Error(
					"Comment should have been removed from comments part",
				)
			}

			// Verify markers were removed from paragraph
			hasStart = false
			hasEnd = false
			hasRef := false
			for child := range para.Children() {
				if child.LocalName() == commentRangeStartName {
					hasStart = true
				}
				if child.LocalName() == commentRangeEndName {
					hasEnd = true
				}
				if run, ok := child.(*elements.Run); ok {
					for runChild := range run.Children() {
						if runChild.LocalName() == "commentReference" {
							hasRef = true
						}
					}
				}
			}
			if hasStart {
				t.Error(
					"CommentRangeStart should have been removed",
				)
			}
			if hasEnd {
				t.Error(
					"CommentRangeEnd should have been removed",
				)
			}
			if hasRef {
				t.Error(
					"CommentReference should have been removed",
				)
			}
		},
	)

	t.Run(
		"RemoveNonExistentComment",
		func(t *testing.T) {
			// Create a document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Create comments part
			_, err = doc.CommentsPart()
			if err != nil {
				t.Fatalf(
					"Failed to get comments part: %v",
					err,
				)
			}

			// Try to remove non-existent comment
			err = doc.RemoveComment(999)
			if err == nil {
				t.Error(
					"Expected error when removing non-existent comment",
				)
			}
		},
	)

	t.Run(
		"MultipleComments_RemoveOne",
		func(t *testing.T) {
			// Create a document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add paragraphs with runs
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			para1 := elements.NewParagraph()
			para1.AppendRun("First paragraph")
			body.AppendChild(para1)

			para2 := elements.NewParagraph()
			para2.AppendRun("Second paragraph")
			body.AppendChild(para2)

			// Apply two comments
			comment1, err := doc.ApplyComment(
				para1,
				0,
				0,
				"Author1",
				"Comment 1",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment 1 failed: %v",
					err,
				)
			}

			comment2, err := doc.ApplyComment(
				para2,
				0,
				0,
				"Author2",
				"Comment 2",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment 2 failed: %v",
					err,
				)
			}

			id1 := comment1.Id()
			id2 := comment2.Id()

			// Remove first comment
			err = doc.RemoveComment(id1)
			if err != nil {
				t.Fatalf(
					"RemoveComment failed: %v",
					err,
				)
			}

			// Verify first comment was removed
			commentsPart := mainPart.CommentsPart()
			if commentsPart.GetComment(
				id1,
			) != nil {
				t.Error(
					"Comment 1 should have been removed",
				)
			}

			// Verify second comment still exists
			if commentsPart.GetComment(
				id2,
			) == nil {
				t.Error(
					"Comment 2 should still exist",
				)
			}

			// Verify markers for comment 1 were removed from para1
			hasMarkers1 := false
			for child := range para1.Children() {
				if child.LocalName() == commentRangeStartName ||
					child.LocalName() == commentRangeEndName {
					hasMarkers1 = true
				}
			}
			if hasMarkers1 {
				t.Error(
					"Markers for comment 1 should have been removed",
				)
			}

			// Verify markers for comment 2 still exist in para2
			hasMarkers2 := false
			for child := range para2.Children() {
				if child.LocalName() == commentRangeStartName ||
					child.LocalName() == commentRangeEndName {
					hasMarkers2 = true
				}
			}
			if !hasMarkers2 {
				t.Error(
					"Markers for comment 2 should still exist",
				)
			}
		},
	)

	t.Run(
		"CascadeIntoNestedElements",
		func(t *testing.T) {
			// Create a document
			doc, err := wordprocessing.New(
				t.TempDir()+"/test.docx",
				wordprocessing.DocTypeDocument,
			)
			if err != nil {
				t.Fatalf(
					"Failed to create document: %v",
					err,
				)
			}
			defer func() { _ = doc.Close() }()

			// Add a table with a paragraph
			mainPart := doc.MainPart()
			docElem := mainPart.Document()
			body := docElem.GetOrCreateBody()

			table := elements.NewTable(1, 1)
			row := elements.NewTableRow(1)
			cell := elements.NewTableCell()
			para := elements.NewParagraph()
			para.AppendRun("Cell text")
			cell.AppendChild(para)
			row.AppendChild(cell)
			table.AppendChild(row)
			body.AppendChild(table)

			// Apply a comment to the paragraph in the table cell
			comment, err := doc.ApplyComment(
				para,
				0,
				0,
				"Author",
				"Table comment",
			)
			if err != nil {
				t.Fatalf(
					"ApplyComment failed: %v",
					err,
				)
			}

			commentID := comment.Id()

			// Verify markers exist
			hasMarkers := false
			for child := range para.Children() {
				if child.LocalName() == commentRangeStartName ||
					child.LocalName() == commentRangeEndName {
					hasMarkers = true
				}
			}
			if !hasMarkers {
				t.Fatal(
					"Comment markers should exist in table cell paragraph",
				)
			}

			// Remove the comment
			err = doc.RemoveComment(commentID)
			if err != nil {
				t.Fatalf(
					"RemoveComment failed: %v",
					err,
				)
			}

			// Verify markers were removed even from nested paragraph
			hasMarkers = false
			for child := range para.Children() {
				if child.LocalName() == commentRangeStartName ||
					child.LocalName() == commentRangeEndName {
					hasMarkers = true
				}
			}
			if hasMarkers {
				t.Error(
					"Comment markers should have been removed from nested paragraph",
				)
			}
		},
	)
}
