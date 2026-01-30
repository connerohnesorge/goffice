package openxml

import (
	"testing"
)

func TestRelationshipCollection(t *testing.T) {
	// Mock container
	container := NewOpenXmlPartData("/word/document.xml", "application/xml", nil, nil)
	rc := NewRelationshipCollection(container)

	t.Run("Add Internal Relationship", func(t *testing.T) {
		targetPart := NewOpenXmlPartData("/word/styles.xml", "application/xml", nil, nil)
		target := NewInternalTarget(targetPart)

		rel, err := rc.Add(RelationshipTypeStyles, target)
		if err != nil {
			t.Fatalf("Add failed: %v", err)
		}

		if rel.Type() != string(RelationshipTypeStyles) {
			t.Errorf("Type = %s, want %s", rel.Type(), RelationshipTypeStyles)
		}

		if rel.Target() != "/word/styles.xml" {
			t.Errorf("Target = %s, want /word/styles.xml", rel.Target())
		}
	})

	t.Run("Add External Relationship", func(t *testing.T) {
		target := NewExternalTarget("https://example.com")

		rel, err := rc.Add(RelationshipTypeHyperlink, target)
		if err != nil {
			t.Fatalf("Add failed: %v", err)
		}

		if rel.Type() != string(RelationshipTypeHyperlink) {
			t.Errorf("Type = %s, want %s", rel.Type(), RelationshipTypeHyperlink)
		}

		if rel.TargetMode() != TargetModeExternal {
			t.Errorf("TargetMode = %v, want External", rel.TargetMode())
		}
	})

	t.Run("GetByType", func(t *testing.T) {
		// Clear collection for this test or use a new one
		rc := NewRelationshipCollection(container)

		// Add 2 styles rels and 1 hyperlink
		targetPart := NewOpenXmlPartData("/word/styles.xml", "application/xml", nil, nil)
		_, _ = rc.Add(RelationshipTypeStyles, NewInternalTarget(targetPart))
		
		targetPart2 := NewOpenXmlPartData("/word/styles2.xml", "application/xml", nil, nil)
		_, _ = rc.Add(RelationshipTypeStyles, NewInternalTarget(targetPart2))

		_, _ = rc.Add(RelationshipTypeHyperlink, NewExternalTarget("https://google.com"))

		styles := rc.GetByType(RelationshipTypeStyles)
		if len(styles) != 2 {
			t.Errorf("GetByType(Styles) count = %d, want 2", len(styles))
		}

		links := rc.GetByType(RelationshipTypeHyperlink)
		if len(links) != 1 {
			t.Errorf("GetByType(Hyperlink) count = %d, want 1", len(links))
		}

		images := rc.GetByType(RelationshipTypeImage)
		if len(images) != 0 {
			t.Errorf("GetByType(Image) count = %d, want 0", len(images))
		}
	})

	t.Run("GetByID and Remove", func(t *testing.T) {
		rc := NewRelationshipCollection(container)
		target := NewExternalTarget("https://example.com")
		rel, _ := rc.Add(RelationshipTypeHyperlink, target)
		id := rel.ID()

		got, err := rc.GetByID(id)
		if err != nil {
			t.Errorf("GetByID failed: %v", err)
		}
		if got != rel {
			t.Error("GetByID returned different relationship")
		}

		err = rc.Remove(id)
		if err != nil {
			t.Errorf("Remove failed: %v", err)
		}

		_, err = rc.GetByID(id)
		if err == nil {
			t.Error("GetByID should fail after removal")
		}
	})
}
