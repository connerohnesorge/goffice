package openxml

import (
	"testing"
)

// Test TargetMode

func TestTargetMode(t *testing.T) {
	t.Run(
		"TargetModeInternal String",
		func(t *testing.T) {
			if TargetModeInternal.String() != "Internal" {
				t.Errorf(
					"TargetModeInternal.String() = %q, want Internal",
					TargetModeInternal.String(),
				)
			}
		},
	)

	t.Run(
		"TargetModeExternal String",
		func(t *testing.T) {
			if TargetModeExternal.String() != "External" {
				t.Errorf(
					"TargetModeExternal.String() = %q, want External",
					TargetModeExternal.String(),
				)
			}
		},
	)
}

// Test baseRelationship

func TestBaseRelationship(t *testing.T) {
	rel := &baseRelationship{
		id:         "rId1",
		relType:    "http://example.com/relationship",
		target:     "/word/document.xml",
		targetMode: TargetModeInternal,
		container:  nil,
	}

	t.Run("ID", func(t *testing.T) {
		if rel.ID() != "rId1" {
			t.Errorf(
				"ID() = %q, want rId1",
				rel.ID(),
			)
		}
	})

	t.Run("Type", func(t *testing.T) {
		if rel.Type() != "http://example.com/relationship" {
			t.Errorf("Type() = %q", rel.Type())
		}
	})

	t.Run("Target", func(t *testing.T) {
		if rel.Target() != "/word/document.xml" {
			t.Errorf(
				"Target() = %q",
				rel.Target(),
			)
		}
	})

	t.Run("TargetMode", func(t *testing.T) {
		if rel.TargetMode() != TargetModeInternal {
			t.Errorf(
				"TargetMode() = %v, want Internal",
				rel.TargetMode(),
			)
		}
	})

	t.Run("Container", func(t *testing.T) {
		if rel.Container() != nil {
			t.Error("Container() should be nil")
		}
	})
}

// Test PartRelationship

func TestPartRelationship(t *testing.T) {
	targetPart := NewOpenXmlPartData(
		"/word/styles.xml",
		"application/xml",
		nil,
		nil,
	)

	rel := NewPartRelationship(
		"rId1",
		RelationshipTypeStyles,
		targetPart,
		nil,
	)

	t.Run("properties", func(t *testing.T) {
		if rel.ID() != "rId1" {
			t.Errorf(
				"ID() = %q, want rId1",
				rel.ID(),
			)
		}
		if rel.Type() != RelationshipTypeStyles {
			t.Errorf("Type() = %q", rel.Type())
		}
		if rel.Target() != "/word/styles.xml" {
			t.Errorf(
				"Target() = %q, want /word/styles.xml",
				rel.Target(),
			)
		}
		if rel.TargetMode() != TargetModeInternal {
			t.Errorf(
				"TargetMode() = %v, want Internal",
				rel.TargetMode(),
			)
		}
	})

	t.Run("TargetPart", func(t *testing.T) {
		if rel.TargetPart() != targetPart {
			t.Error(
				"TargetPart() should return the target part",
			)
		}
	})
}

// Test ExternalRelationship

func TestExternalRelationship(t *testing.T) {
	rel := NewExternalRelationship(
		"rId1",
		RelationshipTypeHyperlink,
		"https://example.com",
		nil,
	)

	t.Run("properties", func(t *testing.T) {
		if rel.ID() != "rId1" {
			t.Errorf(
				"ID() = %q, want rId1",
				rel.ID(),
			)
		}
		if rel.Type() != RelationshipTypeHyperlink {
			t.Errorf("Type() = %q", rel.Type())
		}
		if rel.Target() != "https://example.com" {
			t.Errorf(
				"Target() = %q",
				rel.Target(),
			)
		}
		if rel.TargetMode() != TargetModeExternal {
			t.Errorf(
				"TargetMode() = %v, want External",
				rel.TargetMode(),
			)
		}
	})
}

// Test HyperlinkRelationship

func TestHyperlinkRelationship(t *testing.T) {
	t.Run(
		"external hyperlink",
		func(t *testing.T) {
			rel := NewHyperlinkRelationship(
				"rId1",
				"https://example.com",
				true,
				nil,
			)

			if rel.ID() != "rId1" {
				t.Errorf(
					"ID() = %q, want rId1",
					rel.ID(),
				)
			}
			if rel.Type() != RelationshipTypeHyperlink {
				t.Errorf(
					"Type() = %q, want hyperlink type",
					rel.Type(),
				)
			}
			if !rel.IsExternal() {
				t.Error(
					"IsExternal() should be true",
				)
			}
			if rel.TargetMode() != TargetModeExternal {
				t.Errorf(
					"TargetMode() = %v, want External",
					rel.TargetMode(),
				)
			}
		},
	)

	t.Run(
		"internal hyperlink",
		func(t *testing.T) {
			rel := NewHyperlinkRelationship(
				"rId2",
				"#bookmark1",
				false,
				nil,
			)

			if rel.IsExternal() {
				t.Error(
					"IsExternal() should be false",
				)
			}
			if rel.TargetMode() != TargetModeInternal {
				t.Errorf(
					"TargetMode() = %v, want Internal",
					rel.TargetMode(),
				)
			}
		},
	)
}

// Test DataPartReferenceRelationship

func TestDataPartReferenceRelationship(
	t *testing.T,
) {
	rel := NewDataPartReferenceRelationship(
		"rId1",
		RelationshipTypeImage,
		"/word/media/image1.png",
		nil,
	)

	t.Run("properties", func(t *testing.T) {
		if rel.ID() != "rId1" {
			t.Errorf(
				"ID() = %q, want rId1",
				rel.ID(),
			)
		}
		if rel.Type() != RelationshipTypeImage {
			t.Errorf("Type() = %q", rel.Type())
		}
		if rel.Target() != "/word/media/image1.png" {
			t.Errorf(
				"Target() = %q",
				rel.Target(),
			)
		}
		if rel.TargetMode() != TargetModeInternal {
			t.Errorf(
				"TargetMode() = %v, want Internal",
				rel.TargetMode(),
			)
		}
	})
}

// Test RelationshipIDGenerator

func TestRelationshipIDGenerator(t *testing.T) {
	t.Run(
		"Next generates sequential IDs",
		func(t *testing.T) {
			gen := NewRelationshipIDGenerator()

			id1 := gen.Next()
			id2 := gen.Next()
			id3 := gen.Next()

			if id1 != "rId1" {
				t.Errorf(
					"First ID = %q, want rId1",
					id1,
				)
			}
			if id2 != "rId2" {
				t.Errorf(
					"Second ID = %q, want rId2",
					id2,
				)
			}
			if id3 != "rId3" {
				t.Errorf(
					"Third ID = %q, want rId3",
					id3,
				)
			}
		},
	)

	t.Run(
		"Reserve marks ID as used",
		func(t *testing.T) {
			gen := NewRelationshipIDGenerator()

			gen.Reserve("rId5")

			// Next should skip reserved IDs
			id := gen.Next()
			if id == "rId5" {
				t.Error(
					"Next() should not return reserved ID",
				)
			}
		},
	)

	t.Run(
		"Reserve updates counter for higher IDs",
		func(t *testing.T) {
			gen := NewRelationshipIDGenerator()

			gen.Reserve("rId100")

			// Next should continue from 101
			id := gen.Next()
			if id != "rId101" {
				t.Errorf(
					"After reserving rId100, Next() = %q, want rId101",
					id,
				)
			}
		},
	)

	t.Run(
		"Reserve non-standard format",
		func(t *testing.T) {
			gen := NewRelationshipIDGenerator()

			// Reserve a non-standard ID format
			gen.Reserve("customId")

			// Should still be reserved
			gen.Reserve(
				"customId",
			) // Re-reserving should be safe

			// Should be able to generate rId1
			id := gen.Next()
			if id != "rId1" {
				t.Errorf(
					"Next() = %q, want rId1",
					id,
				)
			}
		},
	)

	t.Run("uniqueness", func(t *testing.T) {
		gen := NewRelationshipIDGenerator()

		ids := make(map[string]bool)
		for i := 0; i < 100; i++ {
			id := gen.Next()
			if ids[id] {
				t.Errorf(
					"Duplicate ID generated: %q",
					id,
				)
			}
			ids[id] = true
		}
	})

	t.Run(
		"skip reserved during generation",
		func(t *testing.T) {
			gen := NewRelationshipIDGenerator()

			// Reserve first few IDs
			gen.Reserve("rId1")
			gen.Reserve("rId2")
			gen.Reserve("rId3")

			// Next should skip to rId4
			id := gen.Next()
			if id != "rId4" {
				t.Errorf(
					"Next() = %q, want rId4",
					id,
				)
			}
		},
	)
}

// Test GenerateUniqueID global function

func TestGenerateUniqueID(t *testing.T) {
	t.Run(
		"generates unique IDs",
		func(t *testing.T) {
			ids := make(map[string]bool)
			for i := 0; i < 100; i++ {
				id := GenerateUniqueID()
				if ids[id] {
					t.Errorf(
						"Duplicate global ID generated: %q",
						id,
					)
				}
				ids[id] = true
			}
		},
	)

	t.Run("format is rIdN", func(t *testing.T) {
		id := GenerateUniqueID()
		if len(id) < 4 || id[:3] != "rId" {
			t.Errorf(
				"ID format unexpected: %q",
				id,
			)
		}
	})
}

// Test ResolveTargetURI

func TestResolveTargetURI(t *testing.T) {
	tests := []struct {
		sourceURI      string
		relativeTarget string
		want           string
	}{
		{
			"/word/document.xml",
			"styles.xml",
			"/word/styles.xml",
		},
		{
			"/word/document.xml",
			"../media/image1.png",
			"/media/image1.png",
		},
		{
			"/word/document.xml",
			"/absolute/path.xml",
			"/absolute/path.xml",
		},
		{
			"/",
			"word/document.xml",
			"/word/document.xml",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.relativeTarget,
			func(t *testing.T) {
				got := ResolveTargetURI(
					tt.sourceURI,
					tt.relativeTarget,
				)
				if got != tt.want {
					t.Errorf(
						"ResolveTargetURI(%q, %q) = %q, want %q",
						tt.sourceURI,
						tt.relativeTarget,
						got,
						tt.want,
					)
				}
			},
		)
	}
}

// Test RelativeTargetURI

func TestRelativeTargetURI(t *testing.T) {
	tests := []struct {
		sourceURI string
		targetURI string
		want      string
	}{
		{
			"/word/document.xml",
			"/word/styles.xml",
			"styles.xml",
		},
		{
			"/word/document.xml",
			"/word/media/image1.png",
			"media/image1.png",
		},
		{
			"/word/document.xml",
			"/docProps/core.xml",
			"/docProps/core.xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.targetURI, func(t *testing.T) {
			got := RelativeTargetURI(
				tt.sourceURI,
				tt.targetURI,
			)
			if got != tt.want {
				t.Errorf(
					"RelativeTargetURI(%q, %q) = %q, want %q",
					tt.sourceURI,
					tt.targetURI,
					got,
					tt.want,
				)
			}
		})
	}
}

// Test RelationshipTypeInfo

func TestGetRelationshipTypeInfo(t *testing.T) {
	t.Run(
		"known relationship type",
		func(t *testing.T) {
			info, ok := GetRelationshipTypeInfo(
				RelationshipTypeOfficeDocument,
			)
			if !ok {
				t.Error(
					"GetRelationshipTypeInfo() should find office document type",
				)
			}
			if info.DefaultPartURI != "/word/document.xml" {
				t.Errorf(
					"DefaultPartURI = %q",
					info.DefaultPartURI,
				)
			}
		},
	)

	t.Run(
		"unknown relationship type",
		func(t *testing.T) {
			_, ok := GetRelationshipTypeInfo(
				"unknown/type",
			)
			if ok {
				t.Error(
					"GetRelationshipTypeInfo() should not find unknown type",
				)
			}
		},
	)
}

// Test RegisterRelationshipType

func TestRegisterRelationshipType(t *testing.T) {
	t.Run(
		"register custom type",
		func(t *testing.T) {
			info := RelationshipTypeInfo{
				Type:           "http://test.custom/relationship",
				DefaultPartURI: "/custom/part.xml",
				ContentType:    "application/custom+xml",
			}

			RegisterRelationshipType(info)

			retrieved, ok := GetRelationshipTypeInfo(
				"http://test.custom/relationship",
			)
			if !ok {
				t.Error(
					"GetRelationshipTypeInfo() should find registered type",
				)
			}
			if retrieved.DefaultPartURI != "/custom/part.xml" {
				t.Errorf(
					"DefaultPartURI = %q, want /custom/part.xml",
					retrieved.DefaultPartURI,
				)
			}
			if retrieved.ContentType != "application/custom+xml" {
				t.Errorf(
					"ContentType = %q",
					retrieved.ContentType,
				)
			}
		},
	)
}

// Test relationship type constants

func TestRelationshipTypeConstants(t *testing.T) {
	// Verify a few key constants are defined correctly
	tests := []struct {
		name  string
		value string
	}{
		{
			"RelationshipTypeDocument",
			RelationshipTypeDocument,
		},
		{
			"RelationshipTypeStyles",
			RelationshipTypeStyles,
		},
		{
			"RelationshipTypeHyperlink",
			RelationshipTypeHyperlink,
		},
		{
			"RelationshipTypeImage",
			RelationshipTypeImage,
		},
		{
			"RelationshipTypeNumbering",
			RelationshipTypeNumbering,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == "" {
				t.Errorf(
					"%s should not be empty",
					tt.name,
				)
			}
			if len(tt.value) < 10 {
				t.Errorf(
					"%s seems too short: %q",
					tt.name,
					tt.value,
				)
			}
		})
	}
}

// Test OpenXmlRelationship interface compliance

func TestRelationshipInterfaceCompliance(
	t *testing.T,
) {
	t.Run(
		"PartRelationship implements OpenXmlRelationship",
		func(t *testing.T) {
			part := NewOpenXmlPartData(
				"/word/styles.xml",
				"application/xml",
				nil,
				nil,
			)
			var _ OpenXmlRelationship = NewPartRelationship("rId1", RelationshipTypeStyles, part, nil)
		},
	)

	t.Run(
		"ExternalRelationship implements OpenXmlRelationship",
		func(t *testing.T) {
			var _ OpenXmlRelationship = NewExternalRelationship("rId1", RelationshipTypeHyperlink, "https://example.com", nil)
		},
	)

	t.Run(
		"HyperlinkRelationship implements OpenXmlRelationship",
		func(t *testing.T) {
			var _ OpenXmlRelationship = NewHyperlinkRelationship("rId1", "https://example.com", true, nil)
		},
	)

	t.Run(
		"DataPartReferenceRelationship implements OpenXmlRelationship",
		func(t *testing.T) {
			var _ OpenXmlRelationship = NewDataPartReferenceRelationship("rId1", RelationshipTypeImage, "/media/image1.png", nil)
		},
	)
}

// Test concurrent access to RelationshipIDGenerator

func TestRelationshipIDGeneratorConcurrent(
	t *testing.T,
) {
	gen := NewRelationshipIDGenerator()

	done := make(chan bool)
	ids := make(chan string, 1000)

	// Spawn multiple goroutines generating IDs
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				ids <- gen.Next()
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
	close(ids)

	// Check uniqueness
	seen := make(map[string]bool)
	for id := range ids {
		if seen[id] {
			t.Errorf(
				"Duplicate ID generated in concurrent test: %q",
				id,
			)
		}
		seen[id] = true
	}

	if len(seen) != 1000 {
		t.Errorf(
			"Expected 1000 unique IDs, got %d",
			len(seen),
		)
	}
}
