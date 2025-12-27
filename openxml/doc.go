// Package openxml provides the core framework for Office Open XML document
// processing.
//
// This package implements the shared infrastructure used by document-specific
// packages (wordprocessing, spreadsheet, presentation). It provides element
// types, attribute handling, part management, and validation infrastructure.
//
// # Element Types
//
// The package provides a hierarchy of element types:
//
//   - Element: Base interface for all XML elements
//   - CompositeElement: Elements that can contain child elements
//   - LeafElement: Elements that contain only text content
//   - PartRootElement: Root elements of document parts
//
// Creating elements:
//
//	// Composite element (can have children)
//	elem := openxml.NewCompositeElement(namespace, localName, prefix)
//	elem.AppendChild(childElement)
//
//	// Leaf element (text content only)
//	leaf := openxml.NewLeafElement(namespace, localName, prefix)
//	leaf.SetText("content")
//
// # Attributes
//
// Elements support attribute manipulation:
//
//	// Create an attribute
//	attr := openxml.NewAttribute(namespace, name, prefix, value)
//
//	// Set/get attributes on elements
//	elem.SetAttribute(attr)
//	attr, found := elem.GetAttribute(name, namespace)
//
//	// Iterate attributes
//	for attr := range elem.Attributes() {
//		fmt.Println(attr.Name(), attr.Value())
//	}
//
// # Part Management
//
// OpenXmlPackage wraps the packaging layer with document-specific
// functionality:
//
//	pkg, err := openxml.CreatePackage("document.docx")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer pkg.Close()
//
//	// Add parts
//	mainPart := openxml.NewOpenXmlPartData(
//		uri, contentType, packPart, container)
//
//	// Get parts by type
//	for part := range pkg.GetPartsOfType(contentType) {
//		// process part
//	}
//
// # Namespaces
//
// Standard Office Open XML namespaces are defined as constants:
//
//   - NamespaceWordprocessingML: WordprocessingML namespace
//   - NamespaceDrawingML: DrawingML namespace
//   - NamespaceRelationships: Relationships namespace
//   - NamespaceContentTypes: Content types namespace
//
// # XML Serialization
//
// Elements can be serialized to XML:
//
//	xml := elem.OuterXml()  // Include element tags
//	xml := elem.InnerXml()  // Children only
//	text := elem.InnerText() // Text content only
//
// Write to an io.Writer:
//
//	err := elem.WriteTo(writer)
//
// # Child Element Access
//
// Navigate the element tree using iterator patterns (Go 1.25+ iter.Seq):
//
//	for child := range elem.Children() {
//		fmt.Println(child.LocalName())
//	}
//
//	// Get specific child element
//	child := elem.GetElement("name", namespace)
//
//	// Get all children with specific name
//	for child := range elem.Elements("name", namespace) {
//		// process child
//	}
//
// # Cloning
//
// Elements support deep and shallow cloning:
//
//	clone := elem.Clone()           // Deep clone
//	shallow := elem.CloneNode(false) // Shallow clone
//	deep := elem.CloneNode(true)     // Deep clone
//
// # Features System
//
// The features system provides extensible functionality:
//
//	features := pkg.Features()
//	features.Set(myFeature)
//	feature := features.Get(featureType)
//
// # Validation
//
// Element and document validation support is provided through the validation
// subpackage. See openxml/validation for details.
//
// # Thread Safety
//
// OpenXmlPackage is thread-safe for concurrent access. Individual elements
// should not be modified concurrently from multiple goroutines.
package openxml
