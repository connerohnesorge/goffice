package wordprocessing

import (
	"os"
	"testing"
)

func TestDocument_ExtendedProperties(
	t *testing.T,
) {
	// Create a temporary document
	tmpFile := "test_extended_properties.docx"
	defer func() { _ = os.Remove(tmpFile) }()

	// Create a new document
	doc, err := New(tmpFile, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Get extended properties
	ep := doc.ExtendedProperties()
	if ep == nil {
		t.Fatal(
			"ExtendedProperties() returned nil",
		)
	}

	// Set extended properties
	ep.SetApplication("Go Office")
	ep.SetApplicationVersion("1.0.0")
	ep.SetCompany("Test Company")
	ep.SetManager("Jane Smith")
	ep.SetTemplate("Normal.dotm")
	ep.SetPages(15)
	ep.SetWords(2500)
	ep.SetCharacters(15000)
	ep.SetCharactersWithSpaces(17500)
	ep.SetLines(300)
	ep.SetParagraphs(150)
	ep.SetTotalTime(180)

	// Ensure the extended properties part exists before saving
	pkg := doc.Package().Package()
	if err := pkg.EnsureExtendedPropertiesPart(); err != nil {
		t.Fatalf(
			"EnsureExtendedPropertiesPart() error = %v",
			err,
		)
	}

	// Save the document
	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Close the document
	if err := doc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Re-open the document
	doc2, err := Open(tmpFile, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify extended properties were persisted
	ep2 := doc2.ExtendedProperties()
	if ep2 == nil {
		t.Fatal(
			"ExtendedProperties() returned nil after re-opening",
		)
	}

	// Verify values
	tests := []struct {
		name string
		got  any
		want any
	}{
		{
			"Application",
			ep2.Application(),
			"Go Office",
		},
		{
			"ApplicationVersion",
			ep2.ApplicationVersion(),
			"1.0.0",
		},
		{
			"Company",
			ep2.Company(),
			"Test Company",
		},
		{"Manager", ep2.Manager(), "Jane Smith"},
		{
			"Template",
			ep2.Template(),
			"Normal.dotm",
		},
		{"Pages", ep2.Pages(), 15},
		{"Words", ep2.Words(), 2500},
		{"Characters", ep2.Characters(), 15000},
		{
			"CharactersWithSpaces",
			ep2.CharactersWithSpaces(),
			17500,
		},
		{"Lines", ep2.Lines(), 300},
		{"Paragraphs", ep2.Paragraphs(), 150},
		{"TotalTime", ep2.TotalTime(), 180},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf(
				"%s = %v, want %v",
				tt.name,
				tt.got,
				tt.want,
			)
		}
	}
}

func TestDocument_ExtendedProperties_ReadExisting(
	t *testing.T,
) {
	// Create a temporary document with extended properties
	tmpFile := "test_read_extended.docx"
	defer func() { _ = os.Remove(tmpFile) }()

	// Create and populate document
	doc, err := New(tmpFile, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	ep := doc.ExtendedProperties()
	ep.SetApplication("Test Application")
	ep.SetCompany("ACME Inc")
	ep.SetPages(42)

	pkg := doc.Package().Package()
	if err := pkg.EnsureExtendedPropertiesPart(); err != nil {
		t.Fatalf(
			"EnsureExtendedPropertiesPart() error = %v",
			err,
		)
	}

	if err := doc.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	_ = doc.Close()

	// Open as read-only
	doc2, err := Open(tmpFile, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	ep2 := doc2.ExtendedProperties()
	if ep2 == nil {
		t.Fatal(
			"ExtendedProperties() returned nil",
		)
	}

	if got := ep2.Application(); got != "Test Application" {
		t.Errorf(
			"Application() = %q, want %q",
			got,
			"Test Application",
		)
	}

	if got := ep2.Company(); got != "ACME Inc" {
		t.Errorf(
			"Company() = %q, want %q",
			got,
			"ACME Inc",
		)
	}

	if got := ep2.Pages(); got != 42 {
		t.Errorf("Pages() = %d, want %d", got, 42)
	}
}

func TestDocument_ExtendedProperties_NotNil(
	t *testing.T,
) {
	// Even for new documents, ExtendedProperties should never return nil
	tmpFile := "test_extended_not_nil.docx"
	defer func() { _ = os.Remove(tmpFile) }()

	doc, err := New(tmpFile, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	ep := doc.ExtendedProperties()
	if ep == nil {
		t.Fatal(
			"ExtendedProperties() returned nil for new document",
		)
	}

	// Should be able to set properties even without calling EnsureExtendedPropertiesPart
	ep.SetApplication("Test")
	if got := ep.Application(); got != "Test" {
		t.Errorf(
			"Application() = %q, want %q",
			got,
			"Test",
		)
	}
}
