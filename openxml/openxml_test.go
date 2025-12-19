package openxml

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/packaging"
)

func TestQualifiedName(t *testing.T) {
	t.Run("NewQualifiedName", func(t *testing.T) {
		qn := NewQualifiedName("http://example.com", "element")
		if qn.NamespaceURI() != "http://example.com" {
			t.Errorf("expected namespace URI http://example.com, got %s", qn.NamespaceURI())
		}
		if qn.LocalName() != "element" {
			t.Errorf("expected local name element, got %s", qn.LocalName())
		}
	})

	t.Run("String with namespace", func(t *testing.T) {
		qn := NewQualifiedName("http://example.com", "element")
		expected := "{http://example.com}element"
		if qn.String() != expected {
			t.Errorf("expected %s, got %s", expected, qn.String())
		}
	})

	t.Run("String without namespace", func(t *testing.T) {
		qn := NewQualifiedName("", "element")
		if qn.String() != "element" {
			t.Errorf("expected element, got %s", qn.String())
		}
	})

	t.Run("ParseQualifiedName with namespace", func(t *testing.T) {
		qn := ParseQualifiedName("{http://example.com}element")
		if qn.NamespaceURI() != "http://example.com" {
			t.Errorf("expected namespace URI http://example.com, got %s", qn.NamespaceURI())
		}
		if qn.LocalName() != "element" {
			t.Errorf("expected local name element, got %s", qn.LocalName())
		}
	})

	t.Run("ParseQualifiedName without namespace", func(t *testing.T) {
		qn := ParseQualifiedName("element")
		if qn.NamespaceURI() != "" {
			t.Errorf("expected empty namespace URI, got %s", qn.NamespaceURI())
		}
		if qn.LocalName() != "element" {
			t.Errorf("expected local name element, got %s", qn.LocalName())
		}
	})
}

func TestAttribute(t *testing.T) {
	t.Run("NewAttribute", func(t *testing.T) {
		attr := NewAttribute("http://example.com", "name", "ex", "value")
		if attr.LocalName() != "name" {
			t.Errorf("expected local name 'name', got %s", attr.LocalName())
		}
		if attr.NamespaceURI() != "http://example.com" {
			t.Errorf("expected namespace URI, got %s", attr.NamespaceURI())
		}
		if attr.Prefix() != "ex" {
			t.Errorf("expected prefix 'ex', got %s", attr.Prefix())
		}
		if attr.Value() != "value" {
			t.Errorf("expected value 'value', got %s", attr.Value())
		}
	})

	t.Run("NewSimpleAttribute", func(t *testing.T) {
		attr := NewSimpleAttribute("name", "value")
		if attr.LocalName() != "name" {
			t.Errorf("expected local name 'name', got %s", attr.LocalName())
		}
		if attr.NamespaceURI() != "" {
			t.Errorf("expected empty namespace URI, got %s", attr.NamespaceURI())
		}
		if attr.Value() != "value" {
			t.Errorf("expected value 'value', got %s", attr.Value())
		}
	})
}

func TestCompositeElement(t *testing.T) {
	t.Run("Create composite element", func(t *testing.T) {
		elem := NewCompositeElement(NamespaceWordprocessingML, "document", "w")
		if elem.LocalName() != "document" {
			t.Errorf("expected local name 'document', got %s", elem.LocalName())
		}
		if elem.NamespaceURI() != NamespaceWordprocessingML {
			t.Errorf("expected namespace URI, got %s", elem.NamespaceURI())
		}
		if elem.Prefix() != "w" {
			t.Errorf("expected prefix 'w', got %s", elem.Prefix())
		}
	})

	t.Run("AppendChild", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")

		parent.AppendChild(child1)
		parent.AppendChild(child2)

		if parent.ChildCount() != 2 {
			t.Errorf("expected 2 children, got %d", parent.ChildCount())
		}

		if parent.FirstChild() != Element(child1) {
			t.Error("expected child1 as first child")
		}

		if parent.LastChild() != Element(child2) {
			t.Error("expected child2 as last child")
		}

		if child1.Parent() != Element(parent) {
			t.Error("expected parent as child1's parent")
		}
	})

	t.Run("PrependChild", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")

		parent.AppendChild(child1)
		parent.PrependChild(child2)

		if parent.FirstChild() != Element(child2) {
			t.Error("expected child2 as first child after prepend")
		}
	})

	t.Run("InsertBefore", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")
		child3 := NewCompositeElement("", "child3", "")

		parent.AppendChild(child1)
		parent.AppendChild(child3)
		parent.InsertBefore(child2, child3)

		count := 0
		var order []string
		for child := range parent.Children() {
			order = append(order, child.LocalName())
			count++
		}

		if count != 3 {
			t.Errorf("expected 3 children, got %d", count)
		}

		expected := []string{"child1", "child2", "child3"}
		for i, name := range expected {
			if order[i] != name {
				t.Errorf("expected %s at position %d, got %s", name, i, order[i])
			}
		}
	})

	t.Run("InsertAfter", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")
		child3 := NewCompositeElement("", "child3", "")

		parent.AppendChild(child1)
		parent.AppendChild(child3)
		parent.InsertAfter(child2, child1)

		var order []string
		for child := range parent.Children() {
			order = append(order, child.LocalName())
		}

		expected := []string{"child1", "child2", "child3"}
		for i, name := range expected {
			if order[i] != name {
				t.Errorf("expected %s at position %d, got %s", name, i, order[i])
			}
		}
	})

	t.Run("RemoveChild", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")

		parent.AppendChild(child1)
		parent.AppendChild(child2)
		parent.RemoveChild(child1)

		if parent.ChildCount() != 1 {
			t.Errorf("expected 1 child after remove, got %d", parent.ChildCount())
		}

		if child1.Parent() != nil {
			t.Error("expected nil parent after remove")
		}
	})

	t.Run("ReplaceChild", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")

		parent.AppendChild(child1)
		parent.ReplaceChild(child2, child1)

		if parent.ChildCount() != 1 {
			t.Errorf("expected 1 child, got %d", parent.ChildCount())
		}

		if parent.FirstChild().LocalName() != "child2" {
			t.Error("expected child2 as first child")
		}

		if child1.Parent() != nil {
			t.Error("expected nil parent for replaced child")
		}
	})

	t.Run("RemoveAllChildren", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("", "child1", "")
		child2 := NewCompositeElement("", "child2", "")

		parent.AppendChild(child1)
		parent.AppendChild(child2)
		parent.RemoveAllChildren()

		if parent.ChildCount() != 0 {
			t.Errorf("expected 0 children, got %d", parent.ChildCount())
		}
	})

	t.Run("GetElement", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewCompositeElement("http://ns1.com", "child", "")
		child2 := NewCompositeElement("http://ns2.com", "child", "")

		parent.AppendChild(child1)
		parent.AppendChild(child2)

		found := parent.GetElement("child", "http://ns2.com")
		if found != Element(child2) {
			t.Error("expected to find child2")
		}
	})
}

func TestLeafElement(t *testing.T) {
	t.Run("Create leaf element", func(t *testing.T) {
		elem := NewLeafElement(NamespaceWordprocessingML, "t", "w")
		if elem.LocalName() != "t" {
			t.Errorf("expected local name 't', got %s", elem.LocalName())
		}
	})

	t.Run("InnerText", func(t *testing.T) {
		elem := NewLeafElementWithText(NamespaceWordprocessingML, "t", "w", "Hello World")
		if elem.InnerText() != "Hello World" {
			t.Errorf("expected 'Hello World', got %s", elem.InnerText())
		}
	})

	t.Run("SetInnerText", func(t *testing.T) {
		elem := NewLeafElement("", "t", "")
		elem.SetInnerText("New Text")
		if elem.InnerText() != "New Text" {
			t.Errorf("expected 'New Text', got %s", elem.InnerText())
		}
	})
}

func TestAttributes(t *testing.T) {
	t.Run("SetAttribute and GetAttribute", func(t *testing.T) {
		elem := NewCompositeElement("", "element", "")
		attr := NewSimpleAttribute("name", "value")
		elem.SetAttribute(attr)

		got, ok := elem.GetAttribute("name", "")
		if !ok {
			t.Error("expected to find attribute")
		}
		if got.Value() != "value" {
			t.Errorf("expected value 'value', got %s", got.Value())
		}
	})

	t.Run("RemoveAttribute", func(t *testing.T) {
		elem := NewCompositeElement("", "element", "")
		attr := NewSimpleAttribute("name", "value")
		elem.SetAttribute(attr)

		removed := elem.RemoveAttribute("name", "")
		if !removed {
			t.Error("expected attribute to be removed")
		}

		_, ok := elem.GetAttribute("name", "")
		if ok {
			t.Error("expected attribute to not be found")
		}
	})

	t.Run("Multiple attributes", func(t *testing.T) {
		elem := NewCompositeElement("", "element", "")
		elem.SetAttribute(NewSimpleAttribute("attr1", "val1"))
		elem.SetAttribute(NewSimpleAttribute("attr2", "val2"))
		elem.SetAttribute(NewSimpleAttribute("attr3", "val3"))

		attrs := elem.Attributes()
		if len(attrs) != 3 {
			t.Errorf("expected 3 attributes, got %d", len(attrs))
		}
	})
}

func TestXmlSerialization(t *testing.T) {
	t.Run("Empty element OuterXml", func(t *testing.T) {
		elem := NewCompositeElement("", "element", "")
		xml := elem.OuterXml()
		if xml != "<element/>" {
			t.Errorf("expected '<element/>', got %s", xml)
		}
	})

	t.Run("Element with children OuterXml", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child := NewLeafElementWithText("", "child", "", "text")
		parent.AppendChild(child)

		xml := parent.OuterXml()
		if !strings.Contains(xml, "<parent>") {
			t.Error("expected opening tag")
		}
		if !strings.Contains(xml, "</parent>") {
			t.Error("expected closing tag")
		}
		if !strings.Contains(xml, "<child>text</child>") {
			t.Error("expected child element")
		}
	})

	t.Run("Element with attributes", func(t *testing.T) {
		elem := NewCompositeElement("", "element", "")
		elem.SetAttribute(NewSimpleAttribute("id", "123"))
		xml := elem.OuterXml()
		if !strings.Contains(xml, `id="123"`) {
			t.Errorf("expected attribute in XML, got %s", xml)
		}
	})

	t.Run("InnerXml", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child1 := NewLeafElementWithText("", "child1", "", "text1")
		child2 := NewLeafElementWithText("", "child2", "", "text2")
		parent.AppendChild(child1)
		parent.AppendChild(child2)

		innerXml := parent.InnerXml()
		if strings.Contains(innerXml, "<parent>") {
			t.Error("InnerXml should not contain parent element")
		}
		if !strings.Contains(innerXml, "<child1>text1</child1>") {
			t.Error("expected child1 in InnerXml")
		}
		if !strings.Contains(innerXml, "<child2>text2</child2>") {
			t.Error("expected child2 in InnerXml")
		}
	})

	t.Run("WriteXML", func(t *testing.T) {
		elem := NewLeafElementWithText("", "test", "", "content")
		var buf bytes.Buffer
		err := elem.WriteXML(&buf)
		if err != nil {
			t.Errorf("WriteXML failed: %v", err)
		}
		if buf.String() != "<test>content</test>" {
			t.Errorf("unexpected output: %s", buf.String())
		}
	})
}

func TestCloning(t *testing.T) {
	t.Run("Clone composite element", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		parent.SetAttribute(NewSimpleAttribute("id", "1"))
		child := NewLeafElementWithText("", "child", "", "text")
		parent.AppendChild(child)

		clone := parent.Clone()

		// Check clone is independent
		if clone == Element(parent) {
			t.Error("clone should be a different instance")
		}

		if clone.Parent() != nil {
			t.Error("clone should have no parent")
		}

		cloneComp := clone.(CompositeElement)
		if cloneComp.ChildCount() != 1 {
			t.Error("clone should have same number of children")
		}

		// Modify clone and verify original is unchanged
		cloneComp.RemoveAllChildren()
		if parent.ChildCount() != 1 {
			t.Error("original should be unchanged after modifying clone")
		}
	})

	t.Run("CloneNode shallow", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		child := NewLeafElementWithText("", "child", "", "text")
		parent.AppendChild(child)

		clone := parent.CloneNode(false)
		cloneComp := clone.(CompositeElement)

		if cloneComp.ChildCount() != 0 {
			t.Error("shallow clone should have no children")
		}
	})
}

func TestHelpers(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		leaf1 := NewLeafElementWithText("", "leaf", "", "first")
		comp := NewCompositeElement("", "comp", "")
		leaf2 := NewLeafElementWithText("", "leaf", "", "second")

		parent.AppendChild(leaf1)
		parent.AppendChild(comp)
		parent.AppendChild(leaf2)

		first := First[*LeafElementBase](parent)
		if first == nil {
			t.Error("expected to find leaf element")
		}
		if first.InnerText() != "first" {
			t.Error("expected first leaf element")
		}
	})

	t.Run("All", func(t *testing.T) {
		parent := NewCompositeElement("", "parent", "")
		leaf1 := NewLeafElementWithText("", "leaf", "", "1")
		leaf2 := NewLeafElementWithText("", "leaf", "", "2")
		comp := NewCompositeElement("", "comp", "")

		parent.AppendChild(leaf1)
		parent.AppendChild(comp)
		parent.AppendChild(leaf2)

		count := 0
		for range All[*LeafElementBase](parent) {
			count++
		}
		if count != 2 {
			t.Errorf("expected 2 leaf elements, got %d", count)
		}
	})

	t.Run("Descendants", func(t *testing.T) {
		root := NewCompositeElement("", "root", "")
		child := NewCompositeElement("", "child", "")
		grandchild := NewLeafElementWithText("", "grandchild", "", "text")

		root.AppendChild(child)
		child.AppendChild(grandchild)

		count := 0
		for range Descendants(root) {
			count++
		}
		if count != 2 {
			t.Errorf("expected 2 descendants, got %d", count)
		}
	})

	t.Run("Ancestors", func(t *testing.T) {
		root := NewCompositeElement("", "root", "")
		child := NewCompositeElement("", "child", "")
		grandchild := NewLeafElementWithText("", "grandchild", "", "text")

		root.AppendChild(child)
		child.AppendChild(grandchild)

		count := 0
		for range Ancestors(grandchild) {
			count++
		}
		if count != 2 {
			t.Errorf("expected 2 ancestors, got %d", count)
		}
	})
}

func TestXmlParsing(t *testing.T) {
	t.Run("Parse simple element", func(t *testing.T) {
		xml := `<root><child>text</child></root>`
		elem, err := ParseElement(strings.NewReader(xml), nil)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if elem.LocalName() != "root" {
			t.Errorf("expected root element, got %s", elem.LocalName())
		}

		comp := elem.(CompositeElement)
		if comp.ChildCount() != 1 {
			t.Errorf("expected 1 child, got %d", comp.ChildCount())
		}
	})

	t.Run("Parse element with attributes", func(t *testing.T) {
		xml := `<element id="123" name="test"/>`
		elem, err := ParseElement(strings.NewReader(xml), nil)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		attrs := elem.Attributes()
		if len(attrs) != 2 {
			t.Errorf("expected 2 attributes, got %d", len(attrs))
		}
	})

	t.Run("Parse nested elements", func(t *testing.T) {
		xml := `<a><b><c>text</c></b></a>`
		elem, err := ParseElement(strings.NewReader(xml), nil)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		count := 0
		for range Descendants(elem.(CompositeElement)) {
			count++
		}
		if count != 2 {
			t.Errorf("expected 2 descendants, got %d", count)
		}
	})
}

func TestNamespaces(t *testing.T) {
	t.Run("GetPrefixForNamespace", func(t *testing.T) {
		prefix := GetPrefixForNamespace(NamespaceWordprocessingML)
		if prefix != "w" {
			t.Errorf("expected 'w', got %s", prefix)
		}
	})

	t.Run("Unknown namespace", func(t *testing.T) {
		prefix := GetPrefixForNamespace("http://unknown.com")
		if prefix != "" {
			t.Errorf("expected empty string, got %s", prefix)
		}
	})
}

// Integration tests - Task 3.59
// These tests combine the openxml package with the packaging layer

func TestIntegrationOpenXmlWithPackaging(t *testing.T) {
	t.Run("create document from scratch", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "integration_test.docx")

		// Create package using packaging layer
		pkg, err := packaging.Create(tmpPath)
		if err != nil {
			t.Fatalf("Failed to create package: %v", err)
		}

		// Wrap with OpenXmlPackage
		oxPkg := NewOpenXmlPackage(pkg)

		// Create main document part
		docPart, err := oxPkg.AddNewPart(
			"/word/document.xml",
			ContentTypeWordprocessingMLDocument,
			RelationshipTypeOfficeDocument,
		)
		if err != nil {
			t.Fatalf("Failed to add document part: %v", err)
		}

		// Set as main part
		oxPkg.SetMainPart(docPart)

		// Create a root element for the document
		root := NewPartRootElement(NamespaceWordprocessingML, "document", "w")

		// Add body element
		body := NewCompositeElement(NamespaceWordprocessingML, "body", "w")
		root.AppendChild(body)

		// Add paragraph
		para := NewCompositeElement(NamespaceWordprocessingML, "p", "w")
		body.AppendChild(para)

		// Add run
		run := NewCompositeElement(NamespaceWordprocessingML, "r", "w")
		para.AppendChild(run)

		// Add text
		text := NewLeafElementWithText(NamespaceWordprocessingML, "t", "w", "Hello, World!")
		run.AppendChild(text)

		// Set root element on part
		if partData, ok := docPart.(*OpenXmlPartData); ok {
			partData.SetRootElement(root)
		}

		// Save
		err = oxPkg.Save()
		if err != nil {
			t.Fatalf("Failed to save: %v", err)
		}

		oxPkg.Close()

		// Verify by reopening
		pkg2, err := packaging.Open(tmpPath, true)
		if err != nil {
			t.Fatalf("Failed to reopen package: %v", err)
		}
		defer pkg2.Close()

		// Verify document part exists
		docPart2, err := pkg2.Part("/word/document.xml")
		if err != nil {
			t.Fatalf("Document part not found: %v", err)
		}

		// Verify content type
		ct, _ := pkg2.ContentTypes().GetContentType("/word/document.xml")
		if ct != ContentTypeWordprocessingMLDocument {
			t.Errorf("Content type = %q, want %q", ct, ContentTypeWordprocessingMLDocument)
		}

		// Verify data was saved
		data := docPart2.GetData()
		if len(data) == 0 {
			t.Error("Document part has no data")
		}
		if !strings.Contains(string(data), "Hello, World!") {
			t.Error("Document content not preserved")
		}
	})

	t.Run("open existing and modify", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "modify_test.docx")

		// Create initial document
		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		docPart, _ := oxPkg.AddNewPart("/word/document.xml", ContentTypeWordprocessingMLDocument, RelationshipTypeOfficeDocument)
		docPart.SetData([]byte("<w:document><w:body><w:p><w:r><w:t>Original</w:t></w:r></w:p></w:body></w:document>"))
		oxPkg.Save()
		oxPkg.Close()

		// Reopen and modify
		oxPkg2, err := OpenPackage(tmpPath, false)
		if err != nil {
			t.Fatalf("Failed to reopen: %v", err)
		}

		// Add a new part
		_, err = oxPkg2.AddNewPart("/word/styles.xml", ContentTypeStyles, RelationshipTypeStyles)
		if err != nil {
			t.Fatalf("Failed to add styles part: %v", err)
		}

		// Save modifications
		err = oxPkg2.Save()
		if err != nil {
			t.Fatalf("Failed to save modifications: %v", err)
		}
		oxPkg2.Close()

		// Verify modifications
		oxPkg3, _ := OpenPackage(tmpPath, true)
		defer oxPkg3.Close()

		_, err = oxPkg3.GetPartByURI("/word/styles.xml")
		if err != nil {
			t.Error("Styles part should exist after modification")
		}
	})

	t.Run("part relationships", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "relationships_test.docx")

		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Create document part
		docPart, _ := oxPkg.AddNewPart("/word/document.xml", ContentTypeWordprocessingMLDocument, RelationshipTypeOfficeDocument)
		oxPkg.SetMainPart(docPart)

		// Create styles part as child of document
		stylesPart := NewOpenXmlPartData("/word/styles.xml", ContentTypeStyles, nil, nil)

		// Add styles as child part of document
		if docData, ok := docPart.(*OpenXmlPartData); ok {
			err := docData.AddPart(stylesPart, "rId1")
			if err != nil {
				t.Fatalf("Failed to add child part: %v", err)
			}

			// Verify child part is accessible
			retrieved, err := docData.GetPartById("rId1")
			if err != nil {
				t.Fatalf("GetPartById failed: %v", err)
			}
			if retrieved != stylesPart {
				t.Error("Wrong child part returned")
			}

			// Verify iteration
			count := 0
			for range docData.Parts() {
				count++
			}
			if count != 1 {
				t.Errorf("Expected 1 child part, got %d", count)
			}
		}

		oxPkg.Close()
	})

	t.Run("feature inheritance", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "features_test.docx")

		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Create part with package as container
		part := NewOpenXmlPartData("/word/document.xml", "application/xml", nil, oxPkg)

		// Verify features are inherited
		if part.Features() == nil {
			t.Error("Part should have features")
		}

		// Part's features should have access to package features through parent
		// This is tested indirectly - if features are nil, many operations would fail

		oxPkg.Close()
	})

	t.Run("content type management", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "content_types_test.docx")

		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Add parts with different content types
		oxPkg.AddNewPart("/word/document.xml", ContentTypeWordprocessingMLDocument, "")
		oxPkg.AddNewPart("/word/styles.xml", ContentTypeStyles, "")
		oxPkg.AddNewPart("/word/numbering.xml", ContentTypeNumbering, "")

		// Verify content types are set
		ct := pkg.ContentTypes()

		docCT, err := ct.GetContentType("/word/document.xml")
		if err != nil || docCT != ContentTypeWordprocessingMLDocument {
			t.Error("Document content type not set correctly")
		}

		stylesCT, err := ct.GetContentType("/word/styles.xml")
		if err != nil || stylesCT != ContentTypeStyles {
			t.Error("Styles content type not set correctly")
		}

		oxPkg.Close()
	})

	t.Run("element to part serialization roundtrip", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "roundtrip_test.docx")

		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Create part with packaging part
		packPart, _ := pkg.CreatePart("/word/document.xml", ContentTypeWordprocessingMLDocument)
		pkg.CreateRelationship("/word/document.xml", RelationshipTypeOfficeDocument, "rId1")

		part := NewOpenXmlPartData("/word/document.xml", ContentTypeWordprocessingMLDocument, packPart, oxPkg)

		// Create element tree
		root := NewPartRootElement(NamespaceWordprocessingML, "document", "w")
		root.SetAttribute(NewAttribute(NamespaceRelationships, "conformance", "xmlns:w", "strict"))

		body := NewCompositeElement(NamespaceWordprocessingML, "body", "w")
		root.AppendChild(body)

		for i := 0; i < 3; i++ {
			para := NewCompositeElement(NamespaceWordprocessingML, "p", "w")
			run := NewCompositeElement(NamespaceWordprocessingML, "r", "w")
			text := NewLeafElementWithText(NamespaceWordprocessingML, "t", "w", "Paragraph text")
			run.AppendChild(text)
			para.AppendChild(run)
			body.AppendChild(para)
		}

		// Set root and save
		part.SetRootElement(root)
		err := part.Save()
		if err != nil {
			t.Fatalf("Save failed: %v", err)
		}

		// Verify data
		data := part.GetData()
		if !strings.Contains(string(data), "Paragraph text") {
			t.Error("Content not serialized correctly")
		}
		if !strings.Contains(string(data), "<?xml") {
			t.Error("XML declaration missing")
		}

		oxPkg.Close()
	})

	t.Run("multiple part types", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "multi_parts_test.docx")

		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Add various parts
		parts := []struct {
			uri         string
			contentType string
			relType     string
		}{
			{"/word/document.xml", ContentTypeWordprocessingMLDocument, RelationshipTypeOfficeDocument},
			{"/word/styles.xml", ContentTypeStyles, RelationshipTypeStyles},
			{"/word/numbering.xml", ContentTypeNumbering, RelationshipTypeNumbering},
			{"/word/settings.xml", ContentTypeSettings, RelationshipTypeSettings},
			{"/word/fontTable.xml", ContentTypeFontTable, RelationshipTypeFontTable},
		}

		for _, p := range parts {
			_, err := oxPkg.AddNewPart(p.uri, p.contentType, p.relType)
			if err != nil {
				t.Errorf("Failed to add part %s: %v", p.uri, err)
			}
		}

		// Verify all parts exist
		count := 0
		for range oxPkg.Parts() {
			count++
		}
		if count != len(parts) {
			t.Errorf("Expected %d parts, got %d", len(parts), count)
		}

		// Test GetPartsOfType
		stylesCount := 0
		for range oxPkg.GetPartsOfType(ContentTypeStyles) {
			stylesCount++
		}
		if stylesCount != 1 {
			t.Errorf("Expected 1 styles part, got %d", stylesCount)
		}

		oxPkg.Save()
		oxPkg.Close()

		// Verify persistence
		oxPkg2, _ := OpenPackage(tmpPath, true)
		defer oxPkg2.Close()

		count2 := 0
		for range oxPkg2.Parts() {
			count2++
		}
		if count2 != len(parts) {
			t.Errorf("After reload: expected %d parts, got %d", len(parts), count2)
		}
	})
}

func TestIntegrationPartRootElement(t *testing.T) {
	t.Run("PartRootElement with part association", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "root_element_test.docx")
		pkg, _ := packaging.Create(tmpPath)
		defer pkg.Close()

		packPart, _ := pkg.CreatePart("/word/document.xml", ContentTypeWordprocessingMLDocument)
		part := NewOpenXmlPartData("/word/document.xml", ContentTypeWordprocessingMLDocument, packPart, nil)

		// Create and associate root element
		root := NewPartRootElement(NamespaceWordprocessingML, "document", "w")
		part.SetRootElement(root)

		// Verify bidirectional association
		if root.Part() != part {
			t.Error("Root element should reference its part")
		}

		if part.RootElement() != root {
			t.Error("Part should reference its root element")
		}
	})

	t.Run("PartRootElement Save and Reload", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "save_reload_test.docx")
		pkg, _ := packaging.Create(tmpPath)

		packPart, _ := pkg.CreatePart("/word/document.xml", ContentTypeWordprocessingMLDocument)
		part := NewOpenXmlPartData("/word/document.xml", ContentTypeWordprocessingMLDocument, packPart, nil)

		// Create element structure
		root := NewPartRootElement(NamespaceWordprocessingML, "document", "w")
		child := NewLeafElementWithText(NamespaceWordprocessingML, "test", "w", "Test content")
		root.AppendChild(child)
		part.SetRootElement(root)

		// Save
		err := part.Save()
		if err != nil {
			t.Fatalf("Save failed: %v", err)
		}

		// Verify content was serialized
		data := part.GetData()
		if !strings.Contains(string(data), "Test content") {
			t.Error("Content not saved correctly")
		}

		pkg.Close()
	})

	t.Run("PartRootElement Clone does not copy part", func(t *testing.T) {
		root := NewPartRootElement(NamespaceWordprocessingML, "document", "w")
		part := NewOpenXmlPartData("/word/document.xml", ContentTypeWordprocessingMLDocument, nil, nil)
		part.SetRootElement(root)

		// Clone the root
		clone := root.Clone()
		cloneRoot, ok := clone.(*PartRootElementBase)
		if !ok {
			t.Fatal("Clone should be a PartRootElementBase")
		}

		// Clone should not be associated with the part
		if cloneRoot.Part() != nil {
			t.Error("Cloned root element should not be associated with a part")
		}
	})
}

func TestIntegrationRelationshipManagement(t *testing.T) {
	t.Run("relationship ID generation across parts", func(t *testing.T) {
		tmpPath := filepath.Join(t.TempDir(), "rel_ids_test.docx")
		pkg, _ := packaging.Create(tmpPath)
		oxPkg := NewOpenXmlPackage(pkg)

		// Add multiple parts and verify unique IDs
		ids := make(map[string]bool)

		for i := 0; i < 10; i++ {
			uri := "/word/part" + string(rune('0'+i)) + ".xml"
			part, err := oxPkg.AddNewPart(uri, "application/xml", "")
			if err != nil {
				t.Fatalf("Failed to add part: %v", err)
			}

			if partData, ok := part.(*OpenXmlPartData); ok {
				id := partData.RelationshipID()
				if ids[id] {
					t.Errorf("Duplicate ID generated: %s", id)
				}
				ids[id] = true
			}
		}

		oxPkg.Close()
	})

	t.Run("external relationships", func(t *testing.T) {
		// Test creating external relationship types
		externalRel := NewExternalRelationship("rId1", RelationshipTypeHyperlink, "https://example.com", nil)

		if externalRel.TargetMode() != TargetModeExternal {
			t.Error("External relationship should have external target mode")
		}

		if externalRel.Target() != "https://example.com" {
			t.Error("Target should be the external URL")
		}
	})

	t.Run("hyperlink relationships", func(t *testing.T) {
		// External hyperlink
		extHyperlink := NewHyperlinkRelationship("rId1", "https://example.com", true, nil)
		if !extHyperlink.IsExternal() {
			t.Error("External hyperlink should be external")
		}

		// Internal hyperlink (bookmark)
		intHyperlink := NewHyperlinkRelationship("rId2", "#bookmark", false, nil)
		if intHyperlink.IsExternal() {
			t.Error("Internal hyperlink should not be external")
		}
	})
}
