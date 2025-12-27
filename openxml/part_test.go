package openxml

import (
	"bytes"
	"testing"

	"github.com/connerohnesorge/goffice/packaging"
)

// Test OpenXmlPartData creation and initialization

func TestNewOpenXmlPartData(t *testing.T) {
	t.Run(
		"create with nil container",
		func(t *testing.T) {
			part := NewOpenXmlPartData(
				"/word/document.xml",
				"application/xml",
				nil,
				nil,
			)

			if part.URI() != "/word/document.xml" {
				t.Errorf(
					"URI() = %q, want /word/document.xml",
					part.URI(),
				)
			}

			if part.ContentType() != "application/xml" {
				t.Errorf(
					"ContentType() = %q, want application/xml",
					part.ContentType(),
				)
			}

			if part.RelationshipID() != "" {
				t.Errorf(
					"RelationshipID() = %q, want empty",
					part.RelationshipID(),
				)
			}

			if part.Container() != nil {
				t.Error(
					"Container() should be nil",
				)
			}

			if part.Features() == nil {
				t.Error(
					"Features() should not be nil",
				)
			}
		},
	)

	t.Run(
		"create with packaging part",
		func(t *testing.T) {
			// Create a mock package first
			tmpPath := t.TempDir() + "/test.docx"
			pkg, err := packaging.Create(tmpPath)
			if err != nil {
				t.Fatalf(
					"Failed to create package: %v",
					err,
				)
			}
			defer func() { _ = pkg.Close() }()

			packPart, err := pkg.CreatePart(
				"/word/document.xml",
				"application/xml",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create packaging part: %v",
					err,
				)
			}

			part := NewOpenXmlPartData(
				"/word/document.xml",
				"application/xml",
				packPart,
				nil,
			)

			if part.PackagingPart() != packPart {
				t.Error(
					"PackagingPart() should return the packaging part",
				)
			}
		},
	)
}

// Test Part loading and content management

func TestOpenXmlPartDataContent(t *testing.T) {
	tmpPath := t.TempDir() + "/test.docx" //nolint:lll
	pkg, err := packaging.Create(tmpPath)
	if err != nil {
		t.Fatalf(
			"Failed to create package: %v",
			err,
		)
	}
	defer func() { _ = pkg.Close() }()

	packPart, err := pkg.CreatePart(
		"/word/document.xml",
		"application/xml",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create packaging part: %v",
			err,
		)
	}

	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		packPart,
		nil,
	)

	t.Run(
		"SetData and GetData",
		func(t *testing.T) {
			testData := []byte(
				"<document>Test content</document>",
			)
			part.SetData(testData)

			retrieved := part.GetData()
			if !bytes.Equal(retrieved, testData) {
				t.Errorf(
					"GetData() = %q, want %q",
					string(retrieved),
					string(testData),
				)
			}

			if !part.IsDirty() {
				t.Error(
					"Part should be dirty after SetData",
				)
			}
		},
	)

	t.Run(
		"GetStream",
		func(t *testing.T) { //nolint:lll
			testData := []byte(
				"<document>Stream test</document>",
			)
			part.SetData(testData)

			reader := part.GetStream()
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(reader)

			if buf.String() != string(testData) {
				t.Errorf(
					"GetStream() returned %q, want %q",
					buf.String(),
					string(testData),
				)
			}
		},
	)

	t.Run(
		"GetData with nil packaging part",
		func(t *testing.T) {
			partNoPkg := NewOpenXmlPartData(
				"/test.xml",
				"application/xml",
				nil,
				nil,
			)
			data := partNoPkg.GetData()
			if data != nil {
				t.Error(
					"GetData() should return nil when no packaging part",
				)
			}
		},
	)

	t.Run(
		"GetStream with nil packaging part",
		func(t *testing.T) {
			partNoPkg := NewOpenXmlPartData(
				"/test.xml",
				"application/xml",
				nil,
				nil,
			)
			reader := partNoPkg.GetStream()
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(reader)
			if buf.Len() != 0 {
				t.Error(
					"GetStream() should return empty reader when no packaging part",
				)
			}
		},
	)
}

// Test PartRootElement functionality

func TestOpenXmlPartDataRootElement(
	t *testing.T,
) {
	tmpPath := t.TempDir() + "/test.docx"
	pkg, err := packaging.Create(tmpPath)
	if err != nil {
		t.Fatalf(
			"Failed to create package: %v",
			err,
		)
	}
	defer func() { _ = pkg.Close() }()

	packPart, err := pkg.CreatePart(
		"/word/document.xml",
		"application/xml",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create packaging part: %v",
			err,
		)
	}

	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		packPart,
		nil,
	)

	t.Run("SetRootElement", func(t *testing.T) {
		root := NewPartRootElement(
			NamespaceWordprocessingML,
			"document",
			"w",
		)
		part.SetRootElement(root)

		if part.RootElement() == nil {
			t.Error(
				"RootElement() should not be nil after SetRootElement",
			)
		}

		if root.Part() != part {
			t.Error(
				"Root element's Part() should point back to the part",
			)
		}

		if !part.IsDirty() {
			t.Error(
				"Part should be dirty after SetRootElement",
			)
		}
	})

	t.Run(
		"RootFactory lazy loading",
		func(t *testing.T) {
			part2 := NewOpenXmlPartData(
				"/word/styles.xml",
				"application/xml",
				nil,
				nil,
			)

			factoryCalled := false
			part2.SetRootFactory(
				func() PartRootElement {
					factoryCalled = true

					return NewPartRootElement(
						NamespaceWordprocessingML,
						"styles",
						"w",
					)
				},
			)

			// First access should trigger factory
			root := part2.RootElement()
			if !factoryCalled {
				t.Error(
					"Root factory should have been called",
				)
			}
			if root == nil {
				t.Error(
					"RootElement() should not be nil after factory",
				)
			}

			// Second access should return cached value (without setting new factory)
			factoryCalled = false
			_ = part2.RootElement()
			if factoryCalled {
				t.Error(
					"Root factory should not be called for cached value",
				)
			}
		},
	)

	t.Run(
		"SetRootElement nil",
		func(t *testing.T) {
			part3 := NewOpenXmlPartData(
				"/word/test.xml",
				"application/xml",
				nil,
				nil,
			)
			part3.SetRootElement(nil)
			if part3.RootElement() != nil {
				t.Error(
					"RootElement() should be nil after SetRootElement(nil)",
				)
			}
		},
	)
}

// Test GetPartsOfType filtering

func TestOpenXmlPartDataGetPartsOfType(
	t *testing.T,
) {
	tmpPath := t.TempDir() + "/test.docx"
	pkg, err := packaging.Create(tmpPath)
	if err != nil {
		t.Fatalf(
			"Failed to create package: %v",
			err,
		)
	}
	defer func() { _ = pkg.Close() }()

	packPart, _ := pkg.CreatePart(
		"/word/document.xml",
		"application/xml",
	)
	parent := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		packPart,
		nil,
	)

	// Add child parts of different content types
	child1 := NewOpenXmlPartData(
		"/word/styles.xml",
		ContentTypeStyles,
		nil,
		parent,
	)
	child2 := NewOpenXmlPartData(
		"/word/numbering.xml",
		ContentTypeNumbering,
		nil,
		parent,
	)
	child3 := NewOpenXmlPartData(
		"/word/settings.xml",
		ContentTypeSettings,
		nil,
		parent,
	)
	// Same content type as child1 for testing GetPartsOfType
	child4 := NewOpenXmlPartData(
		"/word/fontTable.xml",
		ContentTypeStyles,
		nil,
		parent,
	)

	_ = parent.AddPart(child1, "rId1")
	_ = parent.AddPart(child2, "rId2")
	_ = parent.AddPart(child3, "rId3")
	_ = parent.AddPart(child4, "rId4")

	t.Run(
		"filter by content type",
		func(t *testing.T) {
			count := 0
			for range parent.GetPartsOfType(ContentTypeStyles) {
				count++
			}
			if count != 2 {
				t.Errorf(
					"GetPartsOfType(styles) count = %d, want 2",
					count,
				)
			}
		},
	)

	t.Run(
		"filter returns empty for unknown type",
		func(t *testing.T) {
			count := 0
			for range parent.GetPartsOfType("unknown/type") {
				count++
			}
			if count != 0 {
				t.Errorf(
					"GetPartsOfType(unknown) count = %d, want 0",
					count,
				)
			}
		},
	)
}

// Test child part management

func TestOpenXmlPartDataChildParts(t *testing.T) {
	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		nil,
		nil,
	)

	child1 := NewOpenXmlPartData(
		"/word/styles.xml",
		"application/xml",
		nil,
		part,
	)
	child2 := NewOpenXmlPartData(
		"/word/numbering.xml",
		"application/xml",
		nil,
		part,
	)

	t.Run(
		"AddPart with explicit ID",
		func(t *testing.T) {
			err := part.AddPart(child1, "rId1")
			if err != nil {
				t.Fatalf(
					"AddPart() error = %v",
					err,
				)
			}

			retrieved, err := part.GetPartById(
				"rId1",
			)
			if err != nil {
				t.Fatalf(
					"GetPartById() error = %v",
					err,
				)
			}
			if retrieved != child1 {
				t.Error(
					"GetPartById() returned wrong part",
				)
			}

			if child1.RelationshipID() != "rId1" {
				t.Errorf(
					"Child RelationshipID() = %q, want rId1",
					child1.RelationshipID(),
				)
			}
		},
	)

	t.Run(
		"AddPart with auto-generated ID",
		func(t *testing.T) {
			err := part.AddPart(child2, "")
			if err != nil {
				t.Fatalf(
					"AddPart() error = %v",
					err,
				)
			}

			if child2.RelationshipID() == "" {
				t.Error(
					"Child should have auto-generated relationship ID",
				)
			}
		},
	)

	t.Run("DeletePart", func(t *testing.T) {
		// Add and then delete
		child3 := NewOpenXmlPartData(
			"/word/settings.xml",
			"application/xml",
			nil,
			part,
		)
		_ = part.AddPart(child3, "rId99")

		err := part.DeletePart("rId99")
		if err != nil {
			t.Fatalf(
				"DeletePart() error = %v",
				err,
			)
		}

		_, err = part.GetPartById("rId99")
		if err != ErrPartNotFound {
			t.Errorf(
				"GetPartById after delete error = %v, want ErrPartNotFound",
				err,
			)
		}
	})

	t.Run(
		"DeletePart not found",
		func(t *testing.T) {
			err := part.DeletePart("nonexistent")
			if err != ErrPartNotFound {
				t.Errorf(
					"DeletePart(nonexistent) error = %v, want ErrPartNotFound",
					err,
				)
			}
		},
	)

	t.Run("Parts iterator", func(t *testing.T) {
		count := 0
		for range part.Parts() {
			count++
		}
		// Should have child1 and child2 (child3 was deleted)
		if count != 2 {
			t.Errorf(
				"Parts() count = %d, want 2",
				count,
			)
		}
	})

	t.Run(
		"GetPartById not found",
		func(t *testing.T) {
			_, err := part.GetPartById(
				"nonexistent",
			)
			if err != ErrPartNotFound {
				t.Errorf(
					"GetPartById(nonexistent) error = %v, want ErrPartNotFound",
					err,
				)
			}
		},
	)
}

// Test dirty flag management

func TestOpenXmlPartDataDirtyFlag(t *testing.T) {
	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		nil,
		nil,
	)

	t.Run(
		"initial state is not dirty",
		func(t *testing.T) {
			if part.IsDirty() {
				t.Error(
					"New part should not be dirty",
				)
			}
		},
	)

	t.Run("MarkDirty", func(t *testing.T) {
		part.MarkDirty()
		if !part.IsDirty() {
			t.Error(
				"Part should be dirty after MarkDirty",
			)
		}
	})

	t.Run("ClearDirty", func(t *testing.T) {
		part.ClearDirty()
		if part.IsDirty() {
			t.Error(
				"Part should not be dirty after ClearDirty",
			)
		}
	})

	t.Run(
		"AddPart sets dirty",
		func(t *testing.T) {
			part.ClearDirty()
			child := NewOpenXmlPartData(
				"/word/test.xml",
				"application/xml",
				nil,
				part,
			)
			_ = part.AddPart(child, "rIdTest")
			if !part.IsDirty() {
				t.Error(
					"Part should be dirty after AddPart",
				)
			}
		},
	)
}

// Test part type registry

func TestPartTypeRegistry(t *testing.T) {
	t.Run(
		"RegisterPartType and lookup",
		func(t *testing.T) {
			info := &PartTypeInfo{
				ContentType:      "application/test+xml",
				RelationshipType: "http://test.com/relationship",
				Factory: func(uri string, container OpenXmlPartContainer) OpenXmlPart {
					return NewOpenXmlPartData(
						uri,
						"application/test+xml",
						nil,
						container,
					)
				},
			}

			RegisterPartType(info)

			// Lookup by content type
			retrieved, ok := GetPartTypeByContentType(
				"application/test+xml",
			)
			if !ok {
				t.Error(
					"GetPartTypeByContentType() should find registered type",
				)
			}
			if retrieved.ContentType != "application/test+xml" {
				t.Error(
					"Wrong content type returned",
				)
			}

			// Lookup by relationship type
			retrieved, ok = GetPartTypeByRelationship(
				"http://test.com/relationship",
			)
			if !ok {
				t.Error(
					"GetPartTypeByRelationship() should find registered type",
				)
			}
			if retrieved.RelationshipType != "http://test.com/relationship" {
				t.Error(
					"Wrong relationship type returned",
				)
			}
		},
	)

	t.Run(
		"CreatePartByContentType with registered type",
		func(t *testing.T) {
			part := CreatePartByContentType(
				"application/test+xml",
				"/test.xml",
				nil,
				nil,
			)
			if part == nil {
				t.Error(
					"CreatePartByContentType() should return a part",
				)
			}
			if part.ContentType() != "application/test+xml" {
				t.Errorf(
					"Part content type = %q",
					part.ContentType(),
				)
			}
		},
	)

	t.Run(
		"CreatePartByContentType with unknown type",
		func(t *testing.T) {
			part := CreatePartByContentType(
				"unknown/type",
				"/test.xml",
				nil,
				nil,
			)
			if part == nil {
				t.Error(
					"CreatePartByContentType() should return generic part for unknown type",
				)
			}
		},
	)

	t.Run(
		"CreatePartByRelationship with registered type",
		func(t *testing.T) {
			part := CreatePartByRelationship(
				"http://test.com/relationship",
				"/test.xml",
				nil,
				nil,
			)
			if part == nil {
				t.Error(
					"CreatePartByRelationship() should return a part",
				)
			}
		},
	)

	t.Run(
		"CreatePartByRelationship with unknown type",
		func(t *testing.T) {
			part := CreatePartByRelationship(
				"unknown/relationship",
				"/test.xml",
				nil,
				nil,
			)
			if part == nil {
				t.Error(
					"CreatePartByRelationship() should return generic part for unknown type",
				)
			}
		},
	)
}

// Test generic GetPartsOfType and FirstPartOfType helpers

func TestGetPartsOfTypeGeneric(t *testing.T) {
	// Create a container and add mixed part types
	pkg, err := packaging.Create(
		t.TempDir() + "/test.docx",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create package: %v",
			err,
		)
	}
	defer func() { _ = pkg.Close() }()

	oxPkg := NewOpenXmlPackage(pkg)

	// Add some parts
	part1 := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		nil,
		oxPkg,
	)
	part2 := NewOpenXmlPartData(
		"/word/styles.xml",
		"application/xml",
		nil,
		oxPkg,
	)

	_ = oxPkg.AddPart(part1, "rId1")
	_ = oxPkg.AddPart(part2, "rId2")

	t.Run(
		"GetPartsOfType generic",
		func(t *testing.T) {
			count := 0
			for range GetPartsOfType[*OpenXmlPartData](oxPkg) {
				count++
			}
			if count < 2 {
				t.Errorf(
					"GetPartsOfType count = %d, want at least 2",
					count,
				)
			}
		},
	)

	t.Run(
		"FirstPartOfType generic",
		func(t *testing.T) {
			first := FirstPartOfType[*OpenXmlPartData](
				oxPkg,
			)
			if first == nil {
				t.Error(
					"FirstPartOfType() should return a part",
				)
			}
		},
	)
}

// Test Part errors

func TestPartErrors(t *testing.T) {
	t.Run(
		"ErrPartNotFound message",
		func(t *testing.T) {
			if ErrPartNotFound.Error() != "part not found" {
				t.Errorf(
					"ErrPartNotFound.Error() = %q",
					ErrPartNotFound.Error(),
				)
			}
		},
	)

	t.Run(
		"ErrPartExists message",
		func(t *testing.T) {
			if ErrPartExists.Error() != "part already exists" {
				t.Errorf(
					"ErrPartExists.Error() = %q",
					ErrPartExists.Error(),
				)
			}
		},
	)

	t.Run(
		"ErrInvalidPartURI message",
		func(t *testing.T) {
			if ErrInvalidPartURI.Error() != "invalid part URI" {
				t.Errorf(
					"ErrInvalidPartURI.Error() = %q",
					ErrInvalidPartURI.Error(),
				)
			}
		},
	)

	t.Run(
		"ErrContentTypeMismatch message",
		func(t *testing.T) {
			if ErrContentTypeMismatch.Error() != "content type mismatch" {
				t.Errorf(
					"ErrContentTypeMismatch.Error() = %q",
					ErrContentTypeMismatch.Error(),
				)
			}
		},
	)
}

// Test relationship ID on parts

func TestOpenXmlPartDataRelationshipID(
	t *testing.T,
) {
	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		nil,
		nil,
	)

	t.Run(
		"SetRelationshipID",
		func(t *testing.T) {
			part.SetRelationshipID("rId42")
			if part.RelationshipID() != "rId42" {
				t.Errorf(
					"RelationshipID() = %q, want rId42",
					part.RelationshipID(),
				)
			}
		},
	)
}

// Test Save and Reload

func TestOpenXmlPartDataSaveReload(t *testing.T) {
	tmpPath := t.TempDir() + "/test.docx"
	pkg, err := packaging.Create(tmpPath)
	if err != nil {
		t.Fatalf(
			"Failed to create package: %v",
			err,
		)
	}
	defer func() { _ = pkg.Close() }()

	packPart, _ := pkg.CreatePart(
		"/word/document.xml",
		"application/xml",
	)
	part := NewOpenXmlPartData(
		"/word/document.xml",
		"application/xml",
		packPart,
		nil,
	)

	t.Run(
		"Save with root element",
		func(t *testing.T) {
			root := NewPartRootElement(
				NamespaceWordprocessingML,
				"document",
				"w",
			)
			child := NewLeafElementWithText(
				NamespaceWordprocessingML,
				"t",
				"w",
				"Hello World",
			)
			root.AppendChild(child)
			part.SetRootElement(root)

			err := part.Save()
			if err != nil {
				t.Fatalf("Save() error = %v", err)
			}

			if part.IsDirty() {
				t.Error(
					"Part should not be dirty after Save",
				)
			}

			// Verify data was written
			data := part.GetData()
			if len(data) == 0 {
				t.Error(
					"Part should have data after Save",
				)
			}
		},
	)

	t.Run(
		"Save with nil root element",
		func(t *testing.T) {
			part2 := NewOpenXmlPartData(
				"/word/styles.xml",
				"application/xml",
				nil,
				nil,
			)
			err := part2.Save()
			if err != nil {
				t.Errorf(
					"Save() with nil root should not error: %v",
					err,
				)
			}
		},
	)

	t.Run(
		"Reload with nil root element",
		func(t *testing.T) {
			part2 := NewOpenXmlPartData(
				"/word/styles.xml",
				"application/xml",
				nil,
				nil,
			)
			err := part2.Reload()
			if err != nil {
				t.Errorf(
					"Reload() with nil root should not error: %v",
					err,
				)
			}
		},
	)
}

// Test Package accessor

func TestOpenXmlPartDataPackage(t *testing.T) {
	t.Run(
		"Package with container",
		func(t *testing.T) {
			tmpPath := t.TempDir() + "/test.docx"
			pkg, _ := packaging.Create(tmpPath)
			defer func() { _ = pkg.Close() }()

			oxPkg := NewOpenXmlPackage(pkg)
			part := NewOpenXmlPartData(
				"/word/document.xml",
				"application/xml",
				nil,
				oxPkg,
			)

			if part.Package() != pkg {
				t.Error(
					"Package() should return the underlying package",
				)
			}
		},
	)

	t.Run(
		"Package with nil container",
		func(t *testing.T) {
			part := NewOpenXmlPartData(
				"/word/document.xml",
				"application/xml",
				nil,
				nil,
			)
			if part.Package() != nil {
				t.Error(
					"Package() should be nil when container is nil",
				)
			}
		},
	)
}

// Test GetPackagingPart delegation

func TestOpenXmlPartDataGetPackagingPart(
	t *testing.T,
) {
	tmpPath := t.TempDir() + "/test.docx"
	pkg, _ := packaging.Create(tmpPath)
	defer func() { _ = pkg.Close() }()

	packPart, _ := pkg.CreatePart(
		"/word/document.xml",
		"application/xml",
	)

	oxPkg := NewOpenXmlPackage(pkg)

	t.Run(
		"GetPackagingPart with valid URI",
		func(t *testing.T) {
			part := NewOpenXmlPartData(
				"/word/test.xml",
				"application/xml",
				nil,
				oxPkg,
			)
			retrieved := part.GetPackagingPart(
				"/word/document.xml",
			)
			if retrieved != packPart {
				t.Error(
					"GetPackagingPart() should return the packaging part",
				)
			}
		},
	)

	t.Run(
		"GetPackagingPart with nil container",
		func(t *testing.T) {
			part := NewOpenXmlPartData(
				"/word/test.xml",
				"application/xml",
				nil,
				nil,
			)
			retrieved := part.GetPackagingPart(
				"/word/document.xml",
			)
			if retrieved != nil {
				t.Error(
					"GetPackagingPart() should return nil when container is nil",
				)
			}
		},
	)
}
