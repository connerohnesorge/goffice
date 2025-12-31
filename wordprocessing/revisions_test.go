package wordprocessing

import (
	"testing"
	"time"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

const (
	localNameDel = "del"
	localNameR   = "r"
)

func TestRevisionType_String(t *testing.T) {
	tests := []struct {
		name string
		typ  RevisionType
		want string
	}{
		{
			name: "Insert",
			typ:  RevisionTypeInsert,
			want: "Insert",
		},
		{
			name: "Delete",
			typ:  RevisionTypeDelete,
			want: "Delete",
		},
		{
			name: "MoveFrom",
			typ:  RevisionTypeMoveFrom,
			want: "MoveFrom",
		},
		{
			name: "MoveTo",
			typ:  RevisionTypeMoveTo,
			want: "MoveTo",
		},
		{
			name: "FormatChange",
			typ:  RevisionTypeFormatChange,
			want: "FormatChange",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.typ.String(); got != tt.want {
				t.Errorf(
					"RevisionType.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestRevisionTypeDetection(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
		want    RevisionType
	}{
		{
			name: "InsertedRun",
			element: elements.NewInsertedRun(
				1,
				"Test Author",
				testDate,
			),
			want: RevisionTypeInsert,
		},
		{
			name: "DeletedRun",
			element: elements.NewDeletedRun(
				2,
				"Test Author",
				testDate,
			),
			want: RevisionTypeDelete,
		},
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				3,
				"Test Author",
				testDate,
			),
			want: RevisionTypeMoveFrom,
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				4,
				"Test Author",
				testDate,
			),
			want: RevisionTypeMoveTo,
		},
		{
			name: "RunPropertiesChange",
			element: elements.NewRunPropertiesChange(
				5,
				"Test Author",
				testDate,
			),
			want: RevisionTypeFormatChange,
		},
		{
			name: "ParagraphPropertiesChange",
			element: elements.NewParagraphPropertiesChange(
				6,
				"Test Author",
				testDate,
			),
			want: RevisionTypeFormatChange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			if rev.Type() != tt.want {
				t.Errorf(
					"newRevision().Type() = %v, want %v",
					rev.Type(),
					tt.want,
				)
			}
		})
	}
}

func TestRevision_Author(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
		want    string
	}{
		{
			name: "InsertedRun",
			element: elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			),
			want: "Alice",
		},
		{
			name: "DeletedRun",
			element: elements.NewDeletedRun(
				2,
				"Bob",
				testDate,
			),
			want: "Bob",
		},
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				3,
				"Charlie",
				testDate,
			),
			want: "Charlie",
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				4,
				"Diana",
				testDate,
			),
			want: "Diana",
		},
		{
			name: "RunPropertiesChange",
			element: elements.NewRunPropertiesChange(
				5,
				"Eve",
				testDate,
			),
			want: "Eve",
		},
		{
			name: "ParagraphPropertiesChange",
			element: elements.NewParagraphPropertiesChange(
				6,
				"Frank",
				testDate,
			),
			want: "Frank",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			if got := rev.Author(); got != tt.want {
				t.Errorf(
					"Revision.Author() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestRevision_Date(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
	}{
		{
			name: "InsertedRun",
			element: elements.NewInsertedRun(
				1,
				"Author",
				testDate,
			),
		},
		{
			name: "DeletedRun",
			element: elements.NewDeletedRun(
				2,
				"Author",
				testDate,
			),
		},
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				3,
				"Author",
				testDate,
			),
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				4,
				"Author",
				testDate,
			),
		},
		{
			name: "RunPropertiesChange",
			element: elements.NewRunPropertiesChange(
				5,
				"Author",
				testDate,
			),
		},
		{
			name: "ParagraphPropertiesChange",
			element: elements.NewParagraphPropertiesChange(
				6,
				"Author",
				testDate,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			got := rev.Date()
			// Compare timestamps with some tolerance for serialization
			if !got.Equal(testDate) {
				t.Errorf(
					"Revision.Date() = %v, want %v",
					got,
					testDate,
				)
			}
		})
	}
}

func TestRevision_Id(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
		want    int
	}{
		{
			name: "InsertedRun",
			element: elements.NewInsertedRun(
				101,
				"Author",
				testDate,
			),
			want: 101,
		},
		{
			name: "DeletedRun",
			element: elements.NewDeletedRun(
				202,
				"Author",
				testDate,
			),
			want: 202,
		},
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				303,
				"Author",
				testDate,
			),
			want: 303,
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				404,
				"Author",
				testDate,
			),
			want: 404,
		},
		{
			name: "RunPropertiesChange",
			element: elements.NewRunPropertiesChange(
				505,
				"Author",
				testDate,
			),
			want: 505,
		},
		{
			name: "ParagraphPropertiesChange",
			element: elements.NewParagraphPropertiesChange(
				606,
				"Author",
				testDate,
			),
			want: 606,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			if got := rev.Id(); got != tt.want {
				t.Errorf(
					"Revision.Id() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestRevision_Content(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name  string
		setup func() any
		want  string
	}{
		{
			name: "InsertedRun with text",
			setup: func() any {
				ins := elements.NewInsertedRun(
					1,
					"Author",
					testDate,
				)
				ins.AppendRun("Hello World")

				return ins
			},
			want: "Hello World",
		},
		{
			name: "DeletedRun with text",
			setup: func() any {
				del := elements.NewDeletedRun(
					2,
					"Author",
					testDate,
				)
				del.AppendDeletedRun(
					"Deleted text",
				)

				return del
			},
			want: "Deleted text",
		},
		{
			name: "MoveFromRun (no InnerText method)",
			setup: func() any {
				return elements.NewMoveFromRun(
					3,
					"Author",
					testDate,
				)
			},
			want: "",
		},
		{
			name: "MoveToRun (no InnerText method)",
			setup: func() any {
				return elements.NewMoveToRun(
					4,
					"Author",
					testDate,
				)
			},
			want: "",
		},
		{
			name: "RunPropertiesChange (no content)",
			setup: func() any {
				return elements.NewRunPropertiesChange(
					5,
					"Author",
					testDate,
				)
			},
			want: "",
		},
		{
			name: "ParagraphPropertiesChange (no content)",
			setup: func() any {
				return elements.NewParagraphPropertiesChange(
					6,
					"Author",
					testDate,
				)
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := tt.setup()
			rev := newRevision(element, nil, nil)
			if got := rev.Content(); got != tt.want {
				t.Errorf(
					"Revision.Content() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestRevision_Element(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)
	element := elements.NewInsertedRun(
		1,
		"Author",
		testDate,
	)

	rev := newRevision(element, nil, nil)
	if got := rev.Element(); got != element {
		t.Error(
			"Revision.Element() returned different element",
		)
	}
}

func TestRevision_Parent(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)
	element := elements.NewInsertedRun(
		1,
		"Author",
		testDate,
	)
	parent := elements.NewParagraph()

	rev := newRevision(element, parent, nil)
	if got := rev.Parent(); got != parent {
		t.Error(
			"Revision.Parent() returned different parent",
		)
	}
}

func TestRevision_Document(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)
	element := elements.NewInsertedRun(
		1,
		"Author",
		testDate,
	)

	// Create a temporary document for testing
	doc, err := NewWriter(nil, DocTypeDocument)
	if err != nil {
		// We can't create a document without a writer, so just test with nil
		doc = nil
	}

	rev := newRevision(element, nil, doc)
	if got := rev.Document(); got != doc {
		t.Error(
			"Revision.Document() returned different document",
		)
	}
}

// TestRevision_Accept_InsertedRun tests accepting an insertion revision.
func TestRevision_Accept_InsertedRun(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a paragraph with an InsertedRun
	para := elements.NewParagraph()
	ins := elements.NewInsertedRun(
		1,
		"Author",
		testDate,
	)
	ins.AppendRun("Inserted text")
	ins.AppendRun("More text")
	para.AppendChild(ins)

	// Create revision wrapper
	rev := newRevision(ins, para, nil)

	// Accept the revision
	err := rev.Accept()
	if err != nil {
		t.Fatalf(
			"Accept() returned error: %v",
			err,
		)
	}

	// Verify: InsertedRun should be removed, runs should be direct children of paragraph
	var found bool
	for child := range para.Children() {
		if child == ins {
			t.Error(
				"InsertedRun should have been removed from paragraph",
			)
		}
		if child.LocalName() == "r" {
			found = true
		}
	}

	if !found {
		t.Error(
			"Expected to find Run elements as direct children of paragraph",
		)
	}

	// Verify content is preserved
	var runCount int
	for child := range para.Children() {
		if child.LocalName() == "r" {
			runCount++
		}
	}

	if runCount != 2 {
		t.Errorf(
			"Expected 2 run elements, got %d",
			runCount,
		)
	}
}

// TestRevision_Reject_InsertedRun tests rejecting an insertion revision.
func TestRevision_Reject_InsertedRun(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a paragraph with an InsertedRun
	para := elements.NewParagraph()
	ins := elements.NewInsertedRun(
		1,
		"Author",
		testDate,
	)
	ins.AppendRun("Inserted text")
	para.AppendChild(ins)

	// Create revision wrapper
	rev := newRevision(ins, para, nil)

	// Reject the revision
	err := rev.Reject()
	if err != nil {
		t.Fatalf(
			"Reject() returned error: %v",
			err,
		)
	}

	// Verify: InsertedRun and its content should be completely removed
	for child := range para.Children() {
		if child.LocalName() == "ins" ||
			child.LocalName() == "r" {
			t.Error(
				"InsertedRun and runs should have been removed from paragraph",
			)
		}
	}

	// Paragraph should be empty (except maybe properties)
	hasNonPropsChild := false
	for child := range para.Children() {
		if child.LocalName() != "pPr" {
			hasNonPropsChild = true
		}
	}

	if hasNonPropsChild {
		t.Error(
			"Paragraph should only have properties after rejecting insertion",
		)
	}
}

// TestRevision_Accept_DeletedRun tests accepting a deletion revision.
func TestRevision_Accept_DeletedRun(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a paragraph with a DeletedRun
	para := elements.NewParagraph()
	del := elements.NewDeletedRun(
		1,
		"Author",
		testDate,
	)
	del.AppendDeletedRun("Deleted text")
	para.AppendChild(del)

	// Create revision wrapper
	rev := newRevision(del, para, nil)

	// Accept the revision (confirms deletion)
	err := rev.Accept()
	if err != nil {
		t.Fatalf(
			"Accept() returned error: %v",
			err,
		)
	}

	// Verify: DeletedRun should be removed (content stays deleted)
	for child := range para.Children() {
		if child.LocalName() == localNameDel {
			t.Error(
				"DeletedRun should have been removed from paragraph",
			)
		}
	}

	// No regular runs should exist
	for child := range para.Children() {
		if child.LocalName() == localNameR {
			t.Error(
				"No regular runs should exist after accepting deletion",
			)
		}
	}
}

// TestRevision_Reject_DeletedRun tests rejecting a deletion revision.
func TestRevision_Reject_DeletedRun(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a paragraph with a DeletedRun
	para := elements.NewParagraph()
	del := elements.NewDeletedRun(
		1,
		"Author",
		testDate,
	)
	del.AppendDeletedRun("Deleted text")
	para.AppendChild(del)

	// Create revision wrapper
	rev := newRevision(del, para, nil)

	// Reject the revision (restores deleted content)
	err := rev.Reject()
	if err != nil {
		t.Fatalf(
			"Reject() returned error: %v",
			err,
		)
	}

	// Verify: DeletedRun should be removed
	for child := range para.Children() {
		if child.LocalName() == localNameDel {
			t.Error(
				"DeletedRun should have been removed from paragraph",
			)
		}
	}

	// Regular runs with Text elements should exist
	var runCount int
	for child := range para.Children() {
		if child.LocalName() != localNameR {
			continue
		}
		runCount++
		// Verify it contains Text, not DeletedText
		run, ok := child.(*elements.Run)
		if !ok {
			continue
		}
		hasText := false
		hasDelText := false
		for runChild := range run.Children() {
			if runChild.LocalName() == "t" {
				hasText = true
			}
			if runChild.LocalName() == "delText" {
				hasDelText = true
			}
		}
		if !hasText {
			t.Error(
				"Run should contain Text element",
			)
		}
		if hasDelText {
			t.Error(
				"Run should not contain DeletedText after rejecting deletion",
			)
		}
	}

	if runCount != 1 {
		t.Errorf(
			"Expected 1 run element after rejecting deletion, got %d",
			runCount,
		)
	}
}

// TestRevision_Accept_FormatChange_RunProperties tests accepting a run properties change.
func TestRevision_Accept_FormatChange(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a run with properties change
	run := elements.NewRun("Text")
	runProps := run.GetOrCreateProperties()

	// Create a format change tracking element
	change := elements.NewRunPropertiesChange(
		1,
		"Author",
		testDate,
	)

	// Add previous properties to the change
	prevProps := elements.NewRunProperties()
	change.SetPreviousRunProperties(prevProps)

	// Add the change to run properties
	runProps.AppendChild(change)

	// Create revision wrapper
	rev := newRevision(change, runProps, nil)

	// Accept the revision (keeps current properties)
	err := rev.Accept()
	if err != nil {
		t.Fatalf(
			"Accept() returned error: %v",
			err,
		)
	}

	// Verify: Change wrapper should be removed
	for child := range runProps.Children() {
		if child.LocalName() == "rPrChange" {
			t.Error(
				"RunPropertiesChange should have been removed",
			)
		}
	}
}

// TestRevision_Reject_FormatChange tests rejecting a format change.
func TestRevision_Reject_FormatChange(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create a run with properties
	run := elements.NewRun("Text")
	runProps := run.GetOrCreateProperties()

	// Add a current property (e.g., bold)
	currentBold := elements.NewBold()
	runProps.AppendChild(currentBold)

	// Create a format change tracking element with previous properties
	change := elements.NewRunPropertiesChange(
		1,
		"Author",
		testDate,
	)
	prevProps := elements.NewRunProperties()
	// Previous state had italic instead of bold
	prevItalic := elements.NewItalic()
	prevProps.AppendChild(prevItalic)
	change.SetPreviousRunProperties(prevProps)

	// Add the change to run properties
	runProps.AppendChild(change)

	// Create revision wrapper
	rev := newRevision(change, runProps, nil)

	// Reject the revision (restores previous properties)
	err := rev.Reject()
	if err != nil {
		t.Fatalf(
			"Reject() returned error: %v",
			err,
		)
	}

	// Verify: Change wrapper should be removed
	hasChange := false
	for child := range runProps.Children() {
		if child.LocalName() == "rPrChange" {
			hasChange = true
		}
	}
	if hasChange {
		t.Error(
			"RunPropertiesChange should have been removed",
		)
	}

	// Verify: Bold should be removed, Italic should be present
	hasBold := false
	hasItalic := false
	for child := range runProps.Children() {
		if child.LocalName() == "b" {
			hasBold = true
		}
		if child.LocalName() == "i" {
			hasItalic = true
		}
	}

	if hasBold {
		t.Error(
			"Bold should have been removed (was current property)",
		)
	}
	if !hasItalic {
		t.Error(
			"Italic should be present (was previous property)",
		)
	}
}

// TestRevision_Accept_MoveRun tests that accepting move operations returns error.
func TestRevision_Accept_MoveRun(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
	}{
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				1,
				"Author",
				testDate,
			),
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				2,
				"Author",
				testDate,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			err := rev.Accept()
			if err == nil {
				t.Error(
					"Accept() should return error for move operations",
				)
			}
			if err != ErrMoveCoordinationNotImplemented {
				t.Errorf(
					"Expected ErrMoveCoordinationNotImplemented, got %v",
					err,
				)
			}
		})
	}
}

// TestRevision_Reject_MoveRun tests that rejecting move operations returns error.
func TestRevision_Reject_MoveRun(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	tests := []struct {
		name    string
		element any
	}{
		{
			name: "MoveFromRun",
			element: elements.NewMoveFromRun(
				1,
				"Author",
				testDate,
			),
		},
		{
			name: "MoveToRun",
			element: elements.NewMoveToRun(
				2,
				"Author",
				testDate,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := newRevision(
				tt.element,
				nil,
				nil,
			)
			err := rev.Reject()
			if err == nil {
				t.Error(
					"Reject() should return error for move operations",
				)
			}
			if err != ErrMoveCoordinationNotImplemented {
				t.Errorf(
					"Expected ErrMoveCoordinationNotImplemented, got %v",
					err,
				)
			}
		})
	}
}

// TestDocument_GetRevisions tests finding all revisions in a document.
func TestDocument_GetRevisions(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	t.Run("Empty document", func(t *testing.T) {
		doc := createTestDocument(t)
		defer func() { _ = doc.Close() }()

		revisions := doc.GetRevisions()
		if len(revisions) != 0 {
			t.Errorf(
				"Expected 0 revisions, got %d",
				len(revisions),
			)
		}
	})

	t.Run(
		"Document with insertions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			// Add paragraphs with insertions
			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para1 := elements.NewParagraph()
			ins1 := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins1.AppendRun("First insertion")
			para1.AppendChild(ins1)
			body.AppendChild(para1)

			para2 := elements.NewParagraph()
			ins2 := elements.NewInsertedRun(
				2,
				"Bob",
				testDate,
			)
			ins2.AppendRun("Second insertion")
			para2.AppendChild(ins2)
			body.AppendChild(para2)

			revisions := doc.GetRevisions()
			if len(revisions) != 2 {
				t.Errorf(
					"Expected 2 revisions, got %d",
					len(revisions),
				)
			}

			// Verify revision types
			for _, rev := range revisions {
				if rev.Type() != RevisionTypeInsert {
					t.Errorf(
						"Expected RevisionTypeInsert, got %v",
						rev.Type(),
					)
				}
			}
		},
	)

	t.Run(
		"Document with deletions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()
			del := elements.NewDeletedRun(
				1,
				"Alice",
				testDate,
			)
			del.AppendDeletedRun("Deleted text")
			para.AppendChild(del)
			body.AppendChild(para)

			revisions := doc.GetRevisions()
			if len(revisions) != 1 {
				t.Errorf(
					"Expected 1 revision, got %d",
					len(revisions),
				)
			}

			if revisions[0].Type() != RevisionTypeDelete {
				t.Errorf(
					"Expected RevisionTypeDelete, got %v",
					revisions[0].Type(),
				)
			}
		},
	)

	t.Run(
		"Document with multiple revision types",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()

			// Add insertion
			ins := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins.AppendRun("Inserted")
			para.AppendChild(ins)

			// Add deletion
			del := elements.NewDeletedRun(
				2,
				"Bob",
				testDate,
			)
			del.AppendDeletedRun("Deleted")
			para.AppendChild(del)

			// Add run with format change
			run := elements.NewRun("Text")
			runProps := run.GetOrCreateProperties()
			change := elements.NewRunPropertiesChange(
				3,
				"Charlie",
				testDate,
			)
			runProps.AppendChild(change)
			para.AppendChild(run)

			body.AppendChild(para)

			revisions := doc.GetRevisions()
			if len(revisions) != 3 {
				t.Errorf(
					"Expected 3 revisions, got %d",
					len(revisions),
				)
			}

			// Count revision types
			typeCount := make(
				map[RevisionType]int,
			)
			for _, rev := range revisions {
				typeCount[rev.Type()]++
			}

			if typeCount[RevisionTypeInsert] != 1 {
				t.Errorf(
					"Expected 1 insert revision, got %d",
					typeCount[RevisionTypeInsert],
				)
			}
			if typeCount[RevisionTypeDelete] != 1 {
				t.Errorf(
					"Expected 1 delete revision, got %d",
					typeCount[RevisionTypeDelete],
				)
			}
			if typeCount[RevisionTypeFormatChange] != 1 {
				t.Errorf(
					"Expected 1 format change revision, got %d",
					typeCount[RevisionTypeFormatChange],
				)
			}
		},
	)

	t.Run(
		"Document with nested revisions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()

			// Add nested structure: paragraph > table > row > cell > paragraph > insertion
			table := elements.NewTable(1, 1)
			row := elements.NewTableRow(1)
			cell := elements.NewTableCell()
			cellPara := elements.NewParagraph()

			ins := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins.AppendRun("Nested insertion")
			cellPara.AppendChild(ins)

			cell.AppendChild(cellPara)
			row.AppendChild(cell)
			table.AppendChild(row)
			para.AppendChild(table)
			body.AppendChild(para)

			revisions := doc.GetRevisions()
			if len(revisions) != 1 {
				t.Errorf(
					"Expected 1 revision (nested), got %d",
					len(revisions),
				)
			}
		},
	)
}

// TestDocument_AcceptAllRevisions tests accepting all revisions in a document.
func TestDocument_AcceptAllRevisions(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	t.Run(
		"Accept all insertions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()
			ins := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins.AppendRun("Inserted text")
			para.AppendChild(ins)
			body.AppendChild(para)

			// Accept all revisions
			err := doc.AcceptAllRevisions()
			if err != nil {
				t.Fatalf(
					"AcceptAllRevisions() returned error: %v",
					err,
				)
			}

			// Verify no revisions remain
			revisions := doc.GetRevisions()
			if len(revisions) != 0 {
				t.Errorf(
					"Expected 0 revisions after accepting, got %d",
					len(revisions),
				)
			}

			// Verify content is preserved as regular runs
			var runCount int
			for child := range para.Children() {
				if child.LocalName() == localNameR {
					runCount++
				}
			}

			if runCount != 1 {
				t.Errorf(
					"Expected 1 run after accepting insertion, got %d",
					runCount,
				)
			}
		},
	)

	t.Run(
		"Accept all deletions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()
			del := elements.NewDeletedRun(
				1,
				"Alice",
				testDate,
			)
			del.AppendDeletedRun("Deleted text")
			para.AppendChild(del)
			body.AppendChild(para)

			// Accept all revisions
			err := doc.AcceptAllRevisions()
			if err != nil {
				t.Fatalf(
					"AcceptAllRevisions() returned error: %v",
					err,
				)
			}

			// Verify no revisions remain
			revisions := doc.GetRevisions()
			if len(revisions) != 0 {
				t.Errorf(
					"Expected 0 revisions after accepting, got %d",
					len(revisions),
				)
			}

			// Verify deleted content is removed
			for child := range para.Children() {
				if child.LocalName() == localNameDel ||
					child.LocalName() == localNameR {
					t.Error(
						"Deleted content should be removed after accepting deletion",
					)
				}
			}
		},
	)

	t.Run(
		"Accept mixed revisions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()

			ins := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins.AppendRun("Inserted")
			para.AppendChild(ins)

			del := elements.NewDeletedRun(
				2,
				"Bob",
				testDate,
			)
			del.AppendDeletedRun("Deleted")
			para.AppendChild(del)

			body.AppendChild(para)

			// Accept all revisions
			err := doc.AcceptAllRevisions()
			if err != nil {
				t.Fatalf(
					"AcceptAllRevisions() returned error: %v",
					err,
				)
			}

			// Verify no revisions remain
			revisions := doc.GetRevisions()
			if len(revisions) != 0 {
				t.Errorf(
					"Expected 0 revisions after accepting, got %d",
					len(revisions),
				)
			}
		},
	)
}

// TestDocument_RejectAllRevisions tests rejecting all revisions in a document.
func TestDocument_RejectAllRevisions(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	t.Run(
		"Reject all insertions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()
			ins := elements.NewInsertedRun(
				1,
				"Alice",
				testDate,
			)
			ins.AppendRun("Inserted text")
			para.AppendChild(ins)
			body.AppendChild(para)

			// Reject all revisions
			err := doc.RejectAllRevisions()
			if err != nil {
				t.Fatalf(
					"RejectAllRevisions() returned error: %v",
					err,
				)
			}

			// Verify no revisions remain
			revisions := doc.GetRevisions()
			if len(revisions) != 0 {
				t.Errorf(
					"Expected 0 revisions after rejecting, got %d",
					len(revisions),
				)
			}

			// Verify inserted content is removed
			for child := range para.Children() {
				if child.LocalName() == "ins" ||
					child.LocalName() == localNameR {
					t.Error(
						"Inserted content should be removed after rejecting insertion",
					)
				}
			}
		},
	)

	t.Run(
		"Reject all deletions",
		func(t *testing.T) {
			doc := createTestDocument(t)
			defer func() { _ = doc.Close() }()

			body := doc.MainPart().
				Document().
				GetOrCreateBody()

			para := elements.NewParagraph()
			del := elements.NewDeletedRun(
				1,
				"Alice",
				testDate,
			)
			del.AppendDeletedRun("Deleted text")
			para.AppendChild(del)
			body.AppendChild(para)

			// Reject all revisions
			err := doc.RejectAllRevisions()
			if err != nil {
				t.Fatalf(
					"RejectAllRevisions() returned error: %v",
					err,
				)
			}

			// Verify no revisions remain
			revisions := doc.GetRevisions()
			if len(revisions) != 0 {
				t.Errorf(
					"Expected 0 revisions after rejecting, got %d",
					len(revisions),
				)
			}

			// Verify deleted content is restored as regular runs
			var runCount int
			for child := range para.Children() {
				if child.LocalName() == localNameR {
					runCount++
				}
			}

			if runCount != 1 {
				t.Errorf(
					"Expected 1 run after rejecting deletion, got %d",
					runCount,
				)
			}
		},
	)
}

// TestDocument_AcceptRevisionsByAuthor tests accepting revisions by author.
func TestDocument_AcceptRevisionsByAuthor(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	doc := createTestDocument(t)
	defer func() { _ = doc.Close() }()

	body := doc.MainPart().
		Document().
		GetOrCreateBody()

	para := elements.NewParagraph()

	// Add insertions by different authors
	ins1 := elements.NewInsertedRun(
		1,
		"Alice",
		testDate,
	)
	ins1.AppendRun("Alice's insertion")
	para.AppendChild(ins1)

	ins2 := elements.NewInsertedRun(
		2,
		"Bob",
		testDate,
	)
	ins2.AppendRun("Bob's insertion")
	para.AppendChild(ins2)

	body.AppendChild(para)

	// Accept only Alice's revisions
	err := doc.AcceptRevisionsByAuthor("Alice")
	if err != nil {
		t.Fatalf(
			"AcceptRevisionsByAuthor() returned error: %v",
			err,
		)
	}

	// Verify only Bob's revision remains
	revisions := doc.GetRevisions()
	if len(revisions) != 1 {
		t.Errorf(
			"Expected 1 revision (Bob's), got %d",
			len(revisions),
		)
	}

	if len(revisions) > 0 &&
		revisions[0].Author() != "Bob" {
		t.Errorf(
			"Expected Bob's revision to remain, got %s",
			revisions[0].Author(),
		)
	}
}

// TestDocument_RejectRevisionsByAuthor tests rejecting revisions by author.
func TestDocument_RejectRevisionsByAuthor(
	t *testing.T,
) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	doc := createTestDocument(t)
	defer func() { _ = doc.Close() }()

	body := doc.MainPart().
		Document().
		GetOrCreateBody()

	para := elements.NewParagraph()

	// Add deletions by different authors
	del1 := elements.NewDeletedRun(
		1,
		"Alice",
		testDate,
	)
	del1.AppendDeletedRun("Alice's deletion")
	para.AppendChild(del1)

	del2 := elements.NewDeletedRun(
		2,
		"Bob",
		testDate,
	)
	del2.AppendDeletedRun("Bob's deletion")
	para.AppendChild(del2)

	body.AppendChild(para)

	// Reject only Alice's revisions
	err := doc.RejectRevisionsByAuthor("Alice")
	if err != nil {
		t.Fatalf(
			"RejectRevisionsByAuthor() returned error: %v",
			err,
		)
	}

	// Verify only Bob's revision remains
	revisions := doc.GetRevisions()
	if len(revisions) != 1 {
		t.Errorf(
			"Expected 1 revision (Bob's), got %d",
			len(revisions),
		)
	}

	if len(revisions) > 0 &&
		revisions[0].Author() != "Bob" {
		t.Errorf(
			"Expected Bob's revision to remain, got %s",
			revisions[0].Author(),
		)
	}
}

// TestRevisionCollection tests the RevisionCollection type.
func TestRevisionCollection(t *testing.T) {
	testDate := time.Date(
		2025,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC,
	)

	// Create test revisions
	ins1 := elements.NewInsertedRun(
		1,
		"Alice",
		testDate,
	)
	ins2 := elements.NewInsertedRun(
		2,
		"Bob",
		testDate,
	)
	del1 := elements.NewDeletedRun(
		3,
		"Alice",
		testDate,
	)
	change1 := elements.NewRunPropertiesChange(
		4,
		"Charlie",
		testDate,
	)

	revisions := []*Revision{
		newRevision(ins1, nil, nil),
		newRevision(ins2, nil, nil),
		newRevision(del1, nil, nil),
		newRevision(change1, nil, nil),
	}

	collection := NewRevisionCollection(revisions)

	t.Run("Count", func(t *testing.T) {
		if collection.Count() != 4 {
			t.Errorf(
				"Expected count 4, got %d",
				collection.Count(),
			)
		}
	})

	t.Run("ByAuthor", func(t *testing.T) {
		aliceRevs := collection.ByAuthor("Alice")
		if len(aliceRevs) != 2 {
			t.Errorf(
				"Expected 2 Alice revisions, got %d",
				len(aliceRevs),
			)
		}

		bobRevs := collection.ByAuthor("Bob")
		if len(bobRevs) != 1 {
			t.Errorf(
				"Expected 1 Bob revision, got %d",
				len(bobRevs),
			)
		}

		charlieRevs := collection.ByAuthor(
			"Charlie",
		)
		if len(charlieRevs) != 1 {
			t.Errorf(
				"Expected 1 Charlie revision, got %d",
				len(charlieRevs),
			)
		}

		unknownRevs := collection.ByAuthor(
			"Unknown",
		)
		if len(unknownRevs) != 0 {
			t.Errorf(
				"Expected 0 Unknown revisions, got %d",
				len(unknownRevs),
			)
		}
	})

	t.Run("ByType", func(t *testing.T) {
		insertRevs := collection.ByType(
			RevisionTypeInsert,
		)
		if len(insertRevs) != 2 {
			t.Errorf(
				"Expected 2 insert revisions, got %d",
				len(insertRevs),
			)
		}

		deleteRevs := collection.ByType(
			RevisionTypeDelete,
		)
		if len(deleteRevs) != 1 {
			t.Errorf(
				"Expected 1 delete revision, got %d",
				len(deleteRevs),
			)
		}

		formatRevs := collection.ByType(
			RevisionTypeFormatChange,
		)
		if len(formatRevs) != 1 {
			t.Errorf(
				"Expected 1 format change revision, got %d",
				len(formatRevs),
			)
		}

		moveRevs := collection.ByType(
			RevisionTypeMoveFrom,
		)
		if len(moveRevs) != 0 {
			t.Errorf(
				"Expected 0 move revisions, got %d",
				len(moveRevs),
			)
		}
	})

	t.Run("All", func(t *testing.T) {
		allRevs := collection.All()
		if len(allRevs) != 4 {
			t.Errorf(
				"Expected 4 revisions from All(), got %d",
				len(allRevs),
			)
		}
	})
}

// createTestDocument creates a minimal test document for testing.
func createTestDocument(t *testing.T) *Document {
	t.Helper()

	// Create a new document in memory
	doc, err := NewWriter(
		&discardWriter{},
		DocTypeDocument,
	)
	if err != nil {
		t.Fatalf(
			"Failed to create test document: %v",
			err,
		)
	}

	return doc
}

// discardWriter is a writer that discards all data.
type discardWriter struct{}

func (*discardWriter) Write(
	p []byte,
) (n int, err error) {
	return len(p), nil
}
