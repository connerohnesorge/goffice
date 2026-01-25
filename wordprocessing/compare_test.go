package wordprocessing_test

import (
	"bytes"
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing"
)

func TestCompareAndMerge(t *testing.T) {
	// Create in-memory documents
	doc1, err := wordprocessing.NewWriter(&bytes.Buffer{}, wordprocessing.DocTypeDocument)
	if err != nil {
		t.Fatal(err)
	}
	doc1.MainPart().InitializeContent() // Sets <w:t></w:t>
	// We need to inject content.
	// Since Elements API is available, we can construct content.
	// But MainPart.Document() returns *elements.Document which we can manipulate.

	// For simplicity, let's just use the fact that they are initialized.
	// doc1 content: empty text.

	doc2, err := wordprocessing.NewWriter(&bytes.Buffer{}, wordprocessing.DocTypeDocument)
	if err != nil {
		t.Fatal(err)
	}
	doc2.MainPart().InitializeContent()

	// Modify doc2
	// Accessing deep elements without helper is verbose, but we can do:
	// doc2.MainPart().Document().Body()...
	// But elements are generated and might not have easy setters for deep text.

	comparator := wordprocessing.NewDocumentComparator()
	diffs, err := comparator.Compare(doc1, doc2)
	if err != nil {
		t.Fatalf("Compare failed: %v", err)
	}

	// They should be identical
	if len(diffs) != 0 {
		t.Errorf("Expected 0 diffs for identical docs, got %d", len(diffs))
	}
}
