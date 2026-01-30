package compare_test

import (
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/compare"
)

func TestElementMerger_Merge(t *testing.T) {
	merger := compare.NewElementMerger()

	t.Run("Merge Attribute Change", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")
		base.SetAttribute(openxml.NewAttribute("", "attr", "", "base"))

		other := openxml.NewCompositeElement("http://example.com", "root", "ex")
		other.SetAttribute(openxml.NewAttribute("", "attr", "", "new"))

		err := merger.Merge(base, other)
		if err != nil {
			t.Fatalf("Merge failed: %v", err)
		}

		attr, found := base.GetAttribute("attr", "")
		if !found {
			t.Fatal("Attribute not found after merge")
		}
		if attr.Value() != "new" {
			t.Errorf("expected attribute value 'new', got '%s'", attr.Value())
		}
	})

	t.Run("Selective Merge - Ignore Attributes", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")
		base.SetAttribute(openxml.NewAttribute("", "attr", "", "base"))

		other := openxml.NewCompositeElement("http://example.com", "root", "ex")
		other.SetAttribute(openxml.NewAttribute("", "attr", "", "new"))

		merger.SetOptions(compare.MergeOptions{IgnoreAttributes: true})
		err := merger.Merge(base, other)
		if err != nil {
			t.Fatalf("Merge failed: %v", err)
		}

		attr, found := base.GetAttribute("attr", "")
		if !found {
			t.Fatal("Attribute lost during merge (should have been ignored)")
		}
		if attr.Value() != "base" {
			t.Errorf("expected 'base', got '%s'", attr.Value())
		}
		merger.SetOptions(compare.MergeOptions{}) // Reset
	})
}

func TestElementMerger_ThreeWayMerge(t *testing.T) {
	merger := compare.NewElementMerger()

	t.Run("Conflict Attribute", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")
		base.SetAttribute(openxml.NewAttribute("", "attr", "", "base"))

		ours := openxml.NewCompositeElement("http://example.com", "root", "ex")
		ours.SetAttribute(openxml.NewAttribute("", "attr", "", "ours"))

		theirs := openxml.NewCompositeElement("http://example.com", "root", "ex")
		theirs.SetAttribute(openxml.NewAttribute("", "attr", "", "theirs"))

		merged, conflicts, err := merger.ThreeWayMerge(base, ours, theirs)
		if err != nil {
			t.Fatalf("ThreeWayMerge failed: %v", err)
		}

		if len(conflicts) != 1 {
			t.Fatalf("expected 1 conflict, got %d", len(conflicts))
		}
		if conflicts[0].Type != compare.ConflictAttribute {
			t.Errorf("expected ConflictAttribute, got %v", conflicts[0].Type)
		}

		// Merged result should reflect 'ours' in case of conflict (based on current implementation)
		// But let's verify what it holds.
		attr, _ := merged.GetAttribute("attr", "")
		if attr.Value() != "ours" {
			t.Errorf("expected merged to keep 'ours' value 'ours', got '%s'", attr.Value())
		}
	})

	t.Run("No Conflict Different Attributes", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")

		ours := openxml.NewCompositeElement("http://example.com", "root", "ex")
		ours.SetAttribute(openxml.NewAttribute("", "attr1", "", "ours"))

		theirs := openxml.NewCompositeElement("http://example.com", "root", "ex")
		theirs.SetAttribute(openxml.NewAttribute("", "attr2", "", "theirs"))

		merged, conflicts, err := merger.ThreeWayMerge(base, ours, theirs)
		if err != nil {
			t.Fatalf("ThreeWayMerge failed: %v", err)
		}

		if len(conflicts) != 0 {
			t.Errorf("expected 0 conflicts, got %d", len(conflicts))
		}

		attr1, _ := merged.GetAttribute("attr1", "")
		if attr1.Value() != "ours" {
			t.Errorf("expected attr1='ours', got '%s'", attr1.Value())
		}

		attr2, _ := merged.GetAttribute("attr2", "")
		if attr2.Value() != "theirs" {
			t.Errorf("expected attr2='theirs', got '%s'", attr2.Value())
		}
	})

	t.Run("Strategy Theirs Wins", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")
		base.SetAttribute(openxml.NewAttribute("", "attr", "", "base"))

		ours := openxml.NewCompositeElement("http://example.com", "root", "ex")
		ours.SetAttribute(openxml.NewAttribute("", "attr", "", "ours"))

		theirs := openxml.NewCompositeElement("http://example.com", "root", "ex")
		theirs.SetAttribute(openxml.NewAttribute("", "attr", "", "theirs"))

		merger.SetStrategy(&compare.StrategyTheirsWins{})
		merged, conflicts, err := merger.ThreeWayMerge(base, ours, theirs)
		if err != nil {
			t.Fatalf("ThreeWayMerge failed: %v", err)
		}

		if len(conflicts) != 1 {
			t.Fatalf("expected 1 conflict, got %d", len(conflicts))
		}

		attr, _ := merged.GetAttribute("attr", "")
		if attr.Value() != "theirs" {
			t.Errorf("expected merged to keep 'theirs' value 'theirs', got '%s'", attr.Value())
		}

		// Reset strategy
		merger.SetStrategy(&compare.StrategyOursWins{})
	})

	t.Run("Strategy Custom", func(t *testing.T) {
		base := openxml.NewCompositeElement("http://example.com", "root", "ex")
		base.SetAttribute(openxml.NewAttribute("", "attr", "", "base"))

		ours := openxml.NewCompositeElement("http://example.com", "root", "ex")
		ours.SetAttribute(openxml.NewAttribute("", "attr", "", "ours"))

		theirs := openxml.NewCompositeElement("http://example.com", "root", "ex")
		theirs.SetAttribute(openxml.NewAttribute("", "attr", "", "theirs"))

		// Custom strategy: prefer 'base' (reject changes)
		custom := &compare.StrategyCustom{
			Resolver: func(c compare.Conflict) (interface{}, error) {
				return c.BaseValue, nil
			},
		}

		merger.SetStrategy(custom)
		merged, _, err := merger.ThreeWayMerge(base, ours, theirs)
		if err != nil {
			t.Fatalf("ThreeWayMerge failed: %v", err)
		}

		attr, _ := merged.GetAttribute("attr", "")
		if attr.Value() != "base" {
			t.Errorf("expected merged to keep 'base' value 'base', got '%s'", attr.Value())
		}

		merger.SetStrategy(&compare.StrategyOursWins{})
	})

	t.Run("Strategy Combined", func(t *testing.T) {
		base := openxml.NewLeafElement("http://example.com", "text", "ex")
		base.SetInnerText("Original")

		ours := openxml.NewLeafElement("http://example.com", "text", "ex")
		ours.SetInnerText("Ours")

		theirs := openxml.NewLeafElement("http://example.com", "text", "ex")
		theirs.SetInnerText("Theirs")

		merger.SetStrategy(&compare.StrategyCombined{Separator: " | "})
		merged, _, err := merger.ThreeWayMerge(base, ours, theirs)
		if err != nil {
			t.Fatalf("ThreeWayMerge failed: %v", err)
		}

		// Text conflicts might not be reported as ConflictText yet?
		// My diff detection for text uses Modified type.
		// My detectConflicts creates ConflictAttribute for everything currently.
		// I need to fix detectConflicts to set proper ConflictType for text.

		// Let's verify what detectConflicts does.
		// It sets Type: ConflictAttribute.
		// I need to update detectConflicts logic first.

		// But let's write the test expectation assuming I fix it.
		// If I fix it, I expect "Ours | Theirs".
		leaf, ok := merged.(openxml.LeafElement)
		if !ok {
			t.Fatal("merged is not leaf")
		}

		if leaf.InnerText() != "Ours | Theirs" {
			t.Errorf("expected 'Ours | Theirs', got '%s'", leaf.InnerText())
		}

		merger.SetStrategy(&compare.StrategyOursWins{})
	})
}
