package packaging

import (
	"bytes"
	"strings"
	"testing"
)

func TestExtendedProperties_Basic(t *testing.T) {
	ep := NewExtendedProperties()

	// Test application properties
	ep.SetApplication("TestApp")
	if got := ep.Application(); got != "TestApp" {
		t.Errorf(
			"Application() = %q, want %q",
			got,
			"TestApp",
		)
	}

	ep.SetApplicationVersion("1.0.0")
	if got := ep.ApplicationVersion(); got != "1.0.0" {
		t.Errorf(
			"ApplicationVersion() = %q, want %q",
			got,
			"1.0.0",
		)
	}

	ep.SetCompany("Test Company")
	if got := ep.Company(); got != "Test Company" {
		t.Errorf(
			"Company() = %q, want %q",
			got,
			"Test Company",
		)
	}

	ep.SetManager("Test Manager")
	if got := ep.Manager(); got != "Test Manager" {
		t.Errorf(
			"Manager() = %q, want %q",
			got,
			"Test Manager",
		)
	}

	ep.SetTemplate("Test Template")
	if got := ep.Template(); got != "Test Template" {
		t.Errorf(
			"Template() = %q, want %q",
			got,
			"Test Template",
		)
	}

	// Test document statistics
	ep.SetPages(10)
	if got := ep.Pages(); got != 10 {
		t.Errorf("Pages() = %d, want %d", got, 10)
	}

	ep.SetWords(1000)
	if got := ep.Words(); got != 1000 {
		t.Errorf(
			"Words() = %d, want %d",
			got,
			1000,
		)
	}

	ep.SetCharacters(5000)
	if got := ep.Characters(); got != 5000 {
		t.Errorf(
			"Characters() = %d, want %d",
			got,
			5000,
		)
	}

	ep.SetTotalTime(120)
	if got := ep.TotalTime(); got != 120 {
		t.Errorf(
			"TotalTime() = %d, want %d",
			got,
			120,
		)
	}
}

func TestExtendedProperties_MarshalUnmarshal(
	t *testing.T,
) {
	ep := NewExtendedProperties()
	ep.SetApplication("Microsoft Office Word")
	ep.SetApplicationVersion("16.0000")
	ep.SetCompany("Acme Corp")
	ep.SetManager("John Doe")
	ep.SetTemplate("Normal.dotm")
	ep.SetPages(5)
	ep.SetWords(250)
	ep.SetCharacters(1500)
	ep.SetCharactersWithSpaces(1750)
	ep.SetLines(50)
	ep.SetParagraphs(25)
	ep.SetTotalTime(60)
	ep.SetScaleCrop(false)
	ep.SetLinksUpToDate(true)

	// Marshal to XML
	data, err := ep.MarshalToXML()
	if err != nil {
		t.Fatalf("MarshalToXML() error = %v", err)
	}

	// Verify it contains expected data
	xmlStr := string(data)
	if !strings.Contains(
		xmlStr,
		"Microsoft Office Word",
	) {
		t.Error(
			"XML does not contain Application",
		)
	}
	if !strings.Contains(xmlStr, "Acme Corp") {
		t.Error("XML does not contain Company")
	}

	// Unmarshal into new instance
	ep2 := NewExtendedProperties()
	if err := ep2.UnmarshalFromXML(bytes.NewReader(data)); err != nil {
		t.Fatalf(
			"UnmarshalFromXML() error = %v",
			err,
		)
	}

	// Verify values
	if ep2.Application() != "Microsoft Office Word" {
		t.Errorf(
			"Application() = %q, want %q",
			ep2.Application(),
			"Microsoft Office Word",
		)
	}
	if ep2.Company() != "Acme Corp" {
		t.Errorf(
			"Company() = %q, want %q",
			ep2.Company(),
			"Acme Corp",
		)
	}
	if ep2.Pages() != 5 {
		t.Errorf(
			"Pages() = %d, want %d",
			ep2.Pages(),
			5,
		)
	}
	if ep2.Words() != 250 {
		t.Errorf(
			"Words() = %d, want %d",
			ep2.Words(),
			250,
		)
	}
	if ep2.TotalTime() != 60 {
		t.Errorf(
			"TotalTime() = %d, want %d",
			ep2.TotalTime(),
			60,
		)
	}
	if !ep2.LinksUpToDate() {
		t.Errorf(
			"LinksUpToDate() = %v, want %v",
			ep2.LinksUpToDate(),
			true,
		)
	}
}

func TestPackage_ExtendedProperties(
	t *testing.T,
) {
	pkg, err := Create("test_extended.docx")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer func() { _ = pkg.Close() }()

	// Get extended properties
	ep := pkg.ExtendedProperties()
	if ep == nil {
		t.Fatal(
			"ExtendedProperties() returned nil",
		)
	}

	// Set some properties
	ep.SetApplication("Go Office Test")
	ep.SetCompany("Test Company")
	ep.SetPages(10)
	ep.SetWords(500)

	// Ensure the part exists
	if err := pkg.EnsureExtendedPropertiesPart(); err != nil {
		t.Fatalf(
			"EnsureExtendedPropertiesPart() error = %v",
			err,
		)
	}

	// Verify the part was created
	uri := NormalizeURI(ExtendedPropertiesPartURI)
	if _, exists := pkg.parts[uri]; !exists {
		t.Error(
			"Extended properties part was not created",
		)
	}
}
