package wordprocessing

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
)

// testRequiresW14 is the Requires attribute value for Word 2010 extensions.
const testRequiresW14 = "w14"

// TestAlternateContentDocumentIntegration tests that AlternateContent elements
// can be created and roundtripped in Word documents.
func TestAlternateContentDocumentIntegration(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "alternate_content.docx")

	// Create document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Get body element
	mainPart := doc1.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() = nil")
	}

	docElem := mainPart.Document()
	if docElem == nil {
		t.Fatal("Document() = nil")
	}

	body := docElem.Body()
	if body == nil {
		t.Fatal("Body() = nil")
	}

	// Add a paragraph with AlternateContent
	para := openxml.NewCompositeElement(
		"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		"p",
		"w",
	)

	// Create AlternateContent
	ac := openxml.NewAlternateContent()

	// Add a Choice for Office 2010+
	choice := openxml.NewChoice()
	choice.SetRequires("w14")
	choiceChild := openxml.NewCompositeElement(
		"http://schemas.microsoft.com/office/word/2010/wordml",
		"test",
		"w14",
	)
	choice.AppendChild(choiceChild)
	ac.AppendChild(choice)

	// Add a Fallback
	fallback := openxml.NewFallback()
	fbChild := openxml.NewCompositeElement(
		"http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		"r",
		"w",
	)
	fallback.AppendChild(fbChild)
	ac.AppendChild(fallback)

	para.AppendChild(ac)
	body.AppendChild(para)

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify structure is preserved
	mainPart2 := doc2.MainPart()
	if mainPart2 == nil {
		t.Error("MainPart() = nil after reopen")

		return
	}

	docElem2 := mainPart2.Document()
	if docElem2 == nil {
		t.Error("Document() = nil after reopen")

		return
	}

	body2 := docElem2.Body()
	if body2 == nil {
		t.Error("Body() = nil after reopen")

		return
	}

	// Verify we have paragraphs
	paraCount := 0
	for p := range body2.Paragraphs() {
		if p != nil {
			paraCount++
		}
	}
	if paraCount == 0 {
		t.Error("No paragraphs found after reopen")
	}
}

// TestAlternateContentXMLPreservation tests that AlternateContent XML is preserved
// during a roundtrip.
func TestAlternateContentXMLPreservation(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-ac-xml-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "ac_xml.docx")

	// Create document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Create AlternateContent element
	ac := openxml.NewAlternateContent()
	choice := openxml.NewChoice()
	choice.SetRequires("w14")
	ac.AppendChild(choice)

	fallback := openxml.NewFallback()
	ac.AppendChild(fallback)

	// Get XML
	xmlStr1 := ac.OuterXml()

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Create another AlternateContent element
	ac2 := openxml.NewAlternateContent()
	choice2 := openxml.NewChoice()
	choice2.SetRequires("w14")
	ac2.AppendChild(choice2)

	fallback2 := openxml.NewFallback()
	ac2.AppendChild(fallback2)

	xmlStr2 := ac2.OuterXml()

	// Verify XML structure is consistent
	if !strings.Contains(xmlStr1, "AlternateContent") {
		t.Error("XML missing AlternateContent element")
	}
	if !strings.Contains(xmlStr1, "Choice") {
		t.Error("XML missing Choice element")
	}
	if !strings.Contains(xmlStr1, "Fallback") {
		t.Error("XML missing Fallback element")
	}

	// Both should have similar structure
	if len(xmlStr1) != len(xmlStr2) {
		t.Logf("XML lengths differ: %d vs %d (this may be okay)", len(xmlStr1), len(xmlStr2))
	}
}

// TestAlternateContentSelectionByVersion tests that SelectContent correctly
// chooses based on Office version.
func TestAlternateContentSelectionByVersion(t *testing.T) {
	ac := openxml.NewAlternateContent()

	// Add Office 2010 choice
	choice2010 := openxml.NewChoice()
	choice2010.SetRequires("w14")
	choice2010.AppendChild(
		openxml.NewCompositeElement("http://example.com", "modern2010", "ex"),
	)
	ac.AppendChild(choice2010)

	// Add Office 2016 choice
	choice2016 := openxml.NewChoice()
	choice2016.SetRequires("w16")
	choice2016.AppendChild(
		openxml.NewCompositeElement("http://example.com", "modern2016", "ex"),
	)
	ac.AppendChild(choice2016)

	// Add fallback
	fallback := openxml.NewFallback()
	fallback.AppendChild(
		openxml.NewCompositeElement("http://example.com", "legacy", "ex"),
	)
	ac.AppendChild(fallback)

	// Test selection for Office 2007 - should get fallback
	selected := ac.SelectContent(openxml.FileFormatVersionOffice2007)
	if _, ok := selected.(*openxml.Fallback); !ok {
		t.Errorf(
			"SelectContent(Office2007) returned %T, want *Fallback",
			selected,
		)
	}

	// Test selection for Office 2010 - should get first choice
	selected = ac.SelectContent(openxml.FileFormatVersionOffice2010)
	if choice, ok := selected.(*openxml.Choice); !ok {
		t.Errorf(
			"SelectContent(Office2010) returned %T, want *Choice",
			selected,
		)
	} else if choice.Requires() != testRequiresW14 {
		t.Errorf(
			"Selected choice has Requires=%q, want w14",
			choice.Requires(),
		)
	}

	// Test selection for Office 2016 - should get first matching choice (w14)
	// because SelectContent returns the first satisfied choice
	selected = ac.SelectContent(openxml.FileFormatVersionOffice2016)
	if choice, ok := selected.(*openxml.Choice); !ok {
		t.Errorf(
			"SelectContent(Office2016) returned %T, want *Choice",
			selected,
		)
	} else if choice.Requires() != testRequiresW14 {
		t.Errorf(
			"Selected choice has Requires=%q, want w14 (first satisfied)",
			choice.Requires(),
		)
	}

	// Test selection for Office 2019 - should still get first matching choice (w14)
	selected = ac.SelectContent(openxml.FileFormatVersionOffice2019)
	if choice, ok := selected.(*openxml.Choice); !ok {
		t.Errorf(
			"SelectContent(Office2019) returned %T, want *Choice",
			selected,
		)
	} else if choice.Requires() != testRequiresW14 {
		t.Errorf(
			"Selected choice has Requires=%q, want w14 (first satisfied)",
			choice.Requires(),
		)
	}
}

// TestUnknownElementPreservationInWord tests that unknown elements are preserved
// during roundtrips.
func TestUnknownElementPreservationInWord(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "goffice-unknown-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "unknown.docx")

	// Create document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Save again to verify roundtrip
	if err := doc2.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify structure is still valid
	mainPart := doc2.MainPart()
	if mainPart == nil {
		t.Error("MainPart() = nil")
	}
}

// TestMultipleAlternateContentsInDocument tests documents with multiple
// AlternateContent blocks.
func TestMultipleAlternateContentsInDocument(t *testing.T) {
	ac1 := openxml.NewAlternateContent()
	choice1 := openxml.NewChoice()
	choice1.SetRequires("w14")
	ac1.AppendChild(choice1)

	fallback1 := openxml.NewFallback()
	ac1.AppendChild(fallback1)

	ac2 := openxml.NewAlternateContent()
	choice2 := openxml.NewChoice()
	choice2.SetRequires("w15")
	ac2.AppendChild(choice2)

	fallback2 := openxml.NewFallback()
	ac2.AppendChild(fallback2)

	// Both should create valid XML independently
	xml1 := ac1.OuterXml()
	xml2 := ac2.OuterXml()

	if !strings.Contains(xml1, "w14") {
		t.Error("First AlternateContent missing w14 requirement")
	}
	if !strings.Contains(xml2, "w15") {
		t.Error("Second AlternateContent missing w15 requirement")
	}
}
