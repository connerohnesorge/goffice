package compare_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/compare"
)

func TestRoundtrip(t *testing.T) {
	// Create base document structure
	base := openxml.NewCompositeElement("http://example.com", "root", "ex")
	base.SetAttribute(openxml.NewAttribute("", "attr1", "", "val1"))
	child1 := openxml.NewLeafElement("http://example.com", "child", "ex")
	child1.SetInnerText("text1")
	base.AppendChild(child1)

	// Create target document structure (modified)
	target := openxml.NewCompositeElement("http://example.com", "root", "ex")
	target.SetAttribute(openxml.NewAttribute("", "attr1", "", "val2")) // Mod
	target.SetAttribute(openxml.NewAttribute("", "attr2", "", "new"))  // Add
	child1_mod := openxml.NewLeafElement("http://example.com", "child", "ex")
	child1_mod.SetInnerText("text1-mod") // Mod text
	target.AppendChild(child1_mod)
	child2 := openxml.NewLeafElement("http://example.com", "child2", "ex") // Add child
	child2.SetInnerText("text2")
	target.AppendChild(child2)

	// Merge target into base
	merger := compare.NewElementMerger()
	if err := merger.Merge(base, target); err != nil {
		t.Fatalf("Merge failed: %v", err)
	}

	// Compare base (now merged) with target
	comparator := compare.NewElementComparator()
	diffs := comparator.Compare(base, target)

	if len(diffs) > 0 {
		t.Errorf("Roundtrip failed, expected 0 diffs, got %d", len(diffs))
		for _, d := range diffs {
			t.Logf("Diff: %v", d)
		}
	}
}
