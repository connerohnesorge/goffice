package compare_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/compare"
)

func TestLargeMerge(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large merge test in short mode")
	}

	// Create large base document (1000 paragraphs)
	base := openxml.NewCompositeElement("http://example.com", "root", "ex")
	for i := 0; i < 1000; i++ {
		p := openxml.NewCompositeElement("http://example.com", "p", "ex")
		p.SetAttribute(openxml.NewAttribute("", "id", "", fmt.Sprintf("p%d", i)))
		tEl := openxml.NewLeafElement("http://example.com", "t", "ex")
		tEl.SetInnerText(fmt.Sprintf("Paragraph %d", i))
		p.AppendChild(tEl)
		base.AppendChild(p)
	}

	// Create other document with some modifications
	other := base.Clone().(openxml.CompositeElement)

	// Modify 100 random paragraphs
	// Add 50 new paragraphs
	// Delete 50 paragraphs
	// Since we don't have random access by index easily on CompositeElement without ToSlice,
	// we will rebuild 'other' or use ToSlice and rebuild.

	// Let's iterate and modify.
	// Actually, just append 500 new paragraphs to test merge performance.
	for i := 1000; i < 1500; i++ {
		p := openxml.NewCompositeElement("http://example.com", "p", "ex")
		p.SetAttribute(openxml.NewAttribute("", "id", "", fmt.Sprintf("p%d", i)))
		tEl := openxml.NewLeafElement("http://example.com", "t", "ex")
		tEl.SetInnerText(fmt.Sprintf("New Paragraph %d", i))
		p.AppendChild(tEl)
		other.AppendChild(p)
	}

	start := time.Now()
	merger := compare.NewElementMerger()
	err := merger.Merge(base, other)
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Merge failed: %v", err)
	}

	t.Logf("Merged 1000+500 items in %v", duration)

	if base.ChildCount() != 1500 {
		t.Errorf("Expected 1500 children, got %d", base.ChildCount())
	}
}
