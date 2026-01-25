package compare_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/compare"
)

func TestElementComparator_Compare(t *testing.T) {
	cmp := compare.NewElementComparator()

	t.Run("Identical Elements", func(t *testing.T) {
		el1 := openxml.NewLeafElement("http://example.com", "test", "ex")
		el1.SetInnerText("hello")
		el2 := openxml.NewLeafElement("http://example.com", "test", "ex")
		el2.SetInnerText("hello")

		diffs := cmp.Compare(el1, el2)
		if len(diffs) != 0 {
			t.Errorf("expected 0 diffs, got %d", len(diffs))
		}
	})

	t.Run("Text Change", func(t *testing.T) {
		el1 := openxml.NewLeafElement("http://example.com", "test", "ex")
		el1.SetInnerText("hello")
		el2 := openxml.NewLeafElement("http://example.com", "test", "ex")
		el2.SetInnerText("world")

		diffs := cmp.Compare(el1, el2)
		if len(diffs) != 1 {
			t.Errorf("expected 1 diff, got %d", len(diffs))
		}
		if diffs[0].Type != compare.Modified {
			t.Errorf("expected Modified, got %v", diffs[0].Type)
		}
	})

	t.Run("Attribute Added", func(t *testing.T) {
		el1 := openxml.NewCompositeElement("http://example.com", "test", "ex")
		el2 := openxml.NewCompositeElement("http://example.com", "test", "ex")
		el2.SetAttribute(openxml.NewAttribute("", "attr", "", "val"))

		diffs := cmp.Compare(el1, el2)
		if len(diffs) != 1 {
			t.Errorf("expected 1 diff, got %d", len(diffs))
		}
		if diffs[0].Type != compare.Added {
			t.Errorf("expected Added, got %v", diffs[0].Type)
		}
	})

	t.Run("Child Added", func(t *testing.T) {
		el1 := openxml.NewCompositeElement("http://example.com", "root", "ex")
		el2 := openxml.NewCompositeElement("http://example.com", "root", "ex")
		child := openxml.NewLeafElement("http://example.com", "child", "ex")
		el2.AppendChild(child)

		diffs := cmp.Compare(el1, el2)
		if len(diffs) != 1 {
			t.Errorf("expected 1 diff, got %d", len(diffs))
		}
		if diffs[0].Type != compare.Modified {
			t.Errorf("expected Modified parent diff, got %v", diffs[0].Type)
		}
		if len(diffs[0].ChildDiffs) != 1 {
			t.Fatalf("expected 1 child diff, got %d", len(diffs[0].ChildDiffs))
		}
		if diffs[0].ChildDiffs[0].Type != compare.Added {
			t.Errorf("expected Added child diff, got %v", diffs[0].ChildDiffs[0].Type)
		}
	})
}
