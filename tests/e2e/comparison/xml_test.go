package comparison

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
)

// TestCompareXMLStructure_Identical tests comparison of identical documents
func TestCompareXMLStructure_Identical(
	t *testing.T,
) {
	// Create two identical test documents
	doc1Path := createTestDocument(
		t,
		"identical1.docx",
		sampleDocumentXML,
	)
	doc2Path := createTestDocument(
		t,
		"identical2.docx",
		sampleDocumentXML,
	)
	defer func() {
		_ = os.Remove(doc1Path)
	}()
	defer func() {
		_ = os.Remove(doc2Path)
	}()

	tolerance := DefaultTolerance()
	result, err := CompareXMLStructure(
		doc1Path,
		doc2Path,
		tolerance,
	)
	if err != nil {
		t.Fatalf(
			"CompareXMLStructure failed: %v",
			err,
		)
	}

	if !result.Match {
		t.Error(
			"Expected Match=true for identical documents, got false",
		)
		t.Logf(
			"Element mismatches: %d",
			len(result.ElementMismatches),
		)
		t.Logf(
			"Attribute mismatches: %d",
			len(result.AttributeMismatches),
		)
		t.Logf(
			"Content mismatches: %d",
			len(result.ContentMismatches),
		)
	}

	if len(result.ElementMismatches) != 0 {
		t.Errorf(
			"Expected 0 element mismatches, got %d",
			len(result.ElementMismatches),
		)
		for _, m := range result.ElementMismatches {
			t.Logf(
				"  Element mismatch at %s: expected %s, got %s",
				m.Path,
				m.Expected,
				m.Actual,
			)
		}
	}

	if len(result.AttributeMismatches) != 0 {
		t.Errorf(
			"Expected 0 attribute mismatches, got %d",
			len(result.AttributeMismatches),
		)
		for _, m := range result.AttributeMismatches {
			t.Logf(
				"  Attribute mismatch at %s[%s]: expected %s, got %s",
				m.Path,
				m.Attribute,
				m.Expected,
				m.Actual,
			)
		}
	}

	if len(result.ContentMismatches) != 0 {
		t.Errorf(
			"Expected 0 content mismatches, got %d",
			len(result.ContentMismatches),
		)
		for _, m := range result.ContentMismatches {
			t.Logf(
				"  Content mismatch at %s: expected %q, got %q",
				m.Path,
				m.Expected,
				m.Actual,
			)
		}
	}
}

// TestCompareXMLStructure_ElementMismatch tests detection of element differences
func TestCompareXMLStructure_ElementMismatch(
	t *testing.T,
) {
	doc1Path := createTestDocument(
		t,
		"elem1.docx",
		sampleDocumentXML,
	)
	doc2Path := createTestDocument(
		t,
		"elem2.docx",
		documentWithDifferentElement,
	)
	defer func() {
		_ = os.Remove(doc1Path)
	}()
	defer func() {
		_ = os.Remove(doc2Path)
	}()

	tolerance := DefaultTolerance()
	result, err := CompareXMLStructure(
		doc1Path,
		doc2Path,
		tolerance,
	)
	if err != nil {
		t.Fatalf(
			"CompareXMLStructure failed: %v",
			err,
		)
	}

	if result.Match {
		t.Error(
			"Expected Match=false for documents with different elements, got true",
		)
	}

	if len(result.ElementMismatches) == 0 {
		t.Error(
			"Expected element mismatches, got none",
		)
	} else {
		t.Logf("Found %d element mismatch(es) as expected", len(result.ElementMismatches))
		for _, m := range result.ElementMismatches {
			t.Logf("  Element mismatch at %s: expected %s, got %s", m.Path, m.Expected, m.Actual)
		}
	}
}

// TestCompareXMLStructure_AttributeMismatch tests detection of attribute differences
func TestCompareXMLStructure_AttributeMismatch(
	t *testing.T,
) {
	doc1Path := createTestDocument(
		t,
		"attr1.docx",
		sampleDocumentXML,
	)
	doc2Path := createTestDocument(
		t,
		"attr2.docx",
		documentWithDifferentAttribute,
	)
	defer func() {
		_ = os.Remove(doc1Path)
	}()
	defer func() {
		_ = os.Remove(doc2Path)
	}()

	tolerance := DefaultTolerance()
	result, err := CompareXMLStructure(
		doc1Path,
		doc2Path,
		tolerance,
	)
	if err != nil {
		t.Fatalf(
			"CompareXMLStructure failed: %v",
			err,
		)
	}

	if result.Match {
		t.Error(
			"Expected Match=false for documents with different attributes, got true",
		)
	}

	if len(result.AttributeMismatches) == 0 {
		t.Error(
			"Expected attribute mismatches, got none",
		)
	} else {
		t.Logf("Found %d attribute mismatch(es) as expected", len(result.AttributeMismatches))
		for _, m := range result.AttributeMismatches {
			t.Logf("  Attribute mismatch at %s[%s]: expected %s, got %s", m.Path, m.Attribute, m.Expected, m.Actual)
		}
	}
}

// TestCompareXMLStructure_ContentMismatch tests detection of text content differences
func TestCompareXMLStructure_ContentMismatch(
	t *testing.T,
) {
	doc1Path := createTestDocument(
		t,
		"content1.docx",
		sampleDocumentXML,
	)
	doc2Path := createTestDocument(
		t,
		"content2.docx",
		documentWithDifferentContent,
	)
	defer func() {
		_ = os.Remove(doc1Path)
	}()
	defer func() {
		_ = os.Remove(doc2Path)
	}()

	tolerance := DefaultTolerance()
	result, err := CompareXMLStructure(
		doc1Path,
		doc2Path,
		tolerance,
	)
	if err != nil {
		t.Fatalf(
			"CompareXMLStructure failed: %v",
			err,
		)
	}

	if result.Match {
		t.Error(
			"Expected Match=false for documents with different content, got true",
		)
	}

	if len(result.ContentMismatches) == 0 {
		t.Errorf(
			"Expected content mismatches, got none",
		)
	} else {
		t.Logf("Found %d content mismatch(es) as expected", len(result.ContentMismatches))
		for _, m := range result.ContentMismatches {
			t.Logf("  Content mismatch at %s: expected %q, got %q", m.Path, m.Expected, m.Actual)
		}
	}
}

// TestCompareXMLStructure_ToleranceHandling tests attribute order sensitivity tolerance
func TestCompareXMLStructure_ToleranceHandling(
	t *testing.T,
) {
	doc1Path := createTestDocument(
		t,
		"tol1.docx",
		sampleDocumentXML,
	)
	doc2Path := createTestDocument(
		t,
		"tol2.docx",
		sampleDocumentXML,
	)
	defer func() {
		_ = os.Remove(doc1Path)
	}()
	defer func() {
		_ = os.Remove(doc2Path)
	}()

	t.Run("DefaultTolerance", func(t *testing.T) {
		tolerance := DefaultTolerance()
		if tolerance.XMLAttributeOrderSensitive {
			t.Errorf(
				"Expected XMLAttributeOrderSensitive=false in default tolerance, got true",
			)
		}
		if tolerance.VisualPixelTolerance != 2.0 {
			t.Errorf(
				"Expected VisualPixelTolerance=2.0, got %f",
				tolerance.VisualPixelTolerance,
			)
		}
		if tolerance.VisualDiffThreshold != 0.005 {
			t.Errorf(
				"Expected VisualDiffThreshold=0.005, got %f",
				tolerance.VisualDiffThreshold,
			)
		}
	})

	t.Run("StrictTolerance", func(t *testing.T) {
		tolerance := StrictTolerance()
		if !tolerance.XMLAttributeOrderSensitive {
			t.Errorf(
				"Expected XMLAttributeOrderSensitive=true in strict tolerance, got false",
			)
		}
		if tolerance.VisualPixelTolerance != 0.0 {
			t.Errorf(
				"Expected VisualPixelTolerance=0.0, got %f",
				tolerance.VisualPixelTolerance,
			)
		}
		if tolerance.VisualDiffThreshold != 0.0 {
			t.Errorf(
				"Expected VisualDiffThreshold=0.0, got %f",
				tolerance.VisualDiffThreshold,
			)
		}
	})
}

// TestParseXMLToTree tests XML parsing functionality
func TestParseXMLToTree(t *testing.T) {
	xmlData := []byte(
		`<?xml version="1.0" encoding="UTF-8"?>
<root xmlns="http://example.com">
	<parent attr1="value1" attr2="value2">
		<child>Text content</child>
		<child>More text</child>
	</parent>
</root>`,
	)

	tree, err := parseXMLToTree(xmlData)
	if err != nil {
		t.Fatalf("parseXMLToTree failed: %v", err)
	}

	if tree == nil {
		t.Fatal("Expected non-nil tree")
	}

	if tree.Name.Local != "root" {
		t.Errorf(
			"Expected root element 'root', got '%s'",
			tree.Name.Local,
		)
	}

	if len(tree.Children) != 1 {
		t.Fatalf(
			"Expected 1 child of root, got %d",
			len(tree.Children),
		)
	}

	parent := tree.Children[0]
	if parent.Name.Local != "parent" {
		t.Errorf(
			"Expected child element 'parent', got '%s'",
			parent.Name.Local,
		)
	}

	if len(parent.Attributes) != 2 {
		t.Errorf(
			"Expected 2 attributes on parent, got %d",
			len(parent.Attributes),
		)
	}

	if parent.Attributes["attr1"] != "value1" {
		t.Errorf(
			"Expected attr1='value1', got '%s'",
			parent.Attributes["attr1"],
		)
	}

	if len(parent.Children) != 2 {
		t.Fatalf(
			"Expected 2 children of parent, got %d",
			len(parent.Children),
		)
	}

	if parent.Children[0].Content != "Text content" {
		t.Errorf(
			"Expected first child content 'Text content', got '%s'",
			parent.Children[0].Content,
		)
	}
}

// TestExtractXMLPart tests XML part extraction from ZIP
func TestExtractXMLPart(t *testing.T) {
	docPath := createTestDocument(
		t,
		"extract.docx",
		sampleDocumentXML,
	)
	defer func() {
		_ = os.Remove(docPath)
	}()

	zipReader, err := zip.OpenReader(docPath)
	if err != nil {
		t.Fatalf("Failed to open ZIP: %v", err)
	}
	defer func() {
		_ = zipReader.Close()
	}()

	data, err := extractXMLPart(
		&zipReader.Reader,
		"word/document.xml",
	)
	if err != nil {
		t.Fatalf("extractXMLPart failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("Expected non-empty XML data")
	}

	// Verify it's valid XML
	var doc struct{}
	if err := xml.Unmarshal(data, &doc); err != nil {
		t.Errorf(
			"Extracted data is not valid XML: %v",
			err,
		)
	}
}

// TestExtractXMLPart_NotFound tests error handling for missing parts
func TestExtractXMLPart_NotFound(t *testing.T) {
	docPath := createTestDocument(
		t,
		"notfound.docx",
		sampleDocumentXML,
	)
	defer func() {
		_ = os.Remove(docPath)
	}()

	zipReader, err := zip.OpenReader(docPath)
	if err != nil {
		t.Fatalf("Failed to open ZIP: %v", err)
	}
	defer func() {
		_ = zipReader.Close()
	}()

	_, err = extractXMLPart(
		&zipReader.Reader,
		"word/nonexistent.xml",
	)
	if err == nil {
		t.Error(
			"Expected error for non-existent part, got nil",
		)
	}
}

// Helper function to create a test DOCX file with given XML content
func createTestDocument(
	t *testing.T,
	filename, documentXML string,
) string {
	t.Helper()

	tmpDir := t.TempDir()
	docPath := filepath.Join(tmpDir, filename)

	// Create a minimal DOCX ZIP structure
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Add [Content_Types].xml
	contentTypesXML := `<?xml version="1.0" encoding="UTF-8"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
	<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
	<Default Extension="xml" ContentType="application/xml"/>
	<Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
</Types>`

	w, err := zipWriter.Create(
		"[Content_Types].xml",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create [Content_Types].xml: %v",
			err,
		)
	}
	if _, err := w.Write([]byte(contentTypesXML)); err != nil {
		t.Fatalf(
			"Failed to write [Content_Types].xml: %v",
			err,
		)
	}

	// Add _rels/.rels
	relsXML := `<?xml version="1.0" encoding="UTF-8"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
	<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
</Relationships>`

	w, err = zipWriter.Create("_rels/.rels")
	if err != nil {
		t.Fatalf(
			"Failed to create _rels/.rels: %v",
			err,
		)
	}
	if _, err := w.Write([]byte(relsXML)); err != nil {
		t.Fatalf(
			"Failed to write _rels/.rels: %v",
			err,
		)
	}

	// Add word/document.xml
	w, err = zipWriter.Create("word/document.xml")
	if err != nil {
		t.Fatalf(
			"Failed to create word/document.xml: %v",
			err,
		)
	}
	if _, err := w.Write([]byte(documentXML)); err != nil {
		t.Fatalf(
			"Failed to write word/document.xml: %v",
			err,
		)
	}

	if err := zipWriter.Close(); err != nil {
		t.Fatalf(
			"Failed to close zip writer: %v",
			err,
		)
	}

	// Write to file
	if err := os.WriteFile(docPath, buf.Bytes(), 0o644); err != nil {
		t.Fatalf(
			"Failed to write document file: %v",
			err,
		)
	}

	return docPath
}

// Sample document XML for testing
const sampleDocumentXML = `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:body>
		<w:p>
			<w:r>
				<w:t>Hello World</w:t>
			</w:r>
		</w:p>
	</w:body>
</w:document>`

// Document with a different element name
const documentWithDifferentElement = `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:body>
		<w:p>
			<w:run>
				<w:t>Hello World</w:t>
			</w:run>
		</w:p>
	</w:body>
</w:document>`

// Document with a different attribute value
const documentWithDifferentAttribute = `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:body>
		<w:p w:rsidP="00000000">
			<w:r>
				<w:t>Hello World</w:t>
			</w:r>
		</w:p>
	</w:body>
</w:document>`

// Document with different text content
const documentWithDifferentContent = `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
	<w:body>
		<w:p>
			<w:r>
				<w:t>Different Text</w:t>
			</w:r>
		</w:p>
	</w:body>
</w:document>`
