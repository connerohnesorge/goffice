// processing.
// This includes element types, features, and validation infrastructure shared
// across all document types (Word, Excel, PowerPoint).
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
// # Features
//
// The features subpackage provides a hierarchical feature collection system
// that allows elements to access shared services:
//
//   - IPackageFeature: Access to the underlying package
//   - IContentTypeFeature: Content type resolution
//   - INamespaceFeature: Namespace resolution
//   - IPartRelationshipsFeature: Relationship management
//
// # Element Manipulation
//
// CompositeElement provides DOM-style manipulation:
//
//	parent.AppendChild(child)
//	parent.PrependChild(child)
//	parent.InsertBefore(newChild, refChild)
//	parent.InsertAfter(newChild, refChild)
//	parent.RemoveChild(child)
//	parent.ReplaceChild(newChild, oldChild)
//
// # Traversal Helpers
//
// Generic helper functions for element traversal:
//
//	First[T](parent)           - Get first child of type T
//	All[T](parent)             - Iterate all children of type T
//	Descendants(el)            - Iterate all descendants
//	DescendantsOfType[T](el)   - Iterate descendants of type T
//	Ancestors(el)              - Iterate ancestors up to root
//
// # XML Serialization
//
// Elements can be serialized to and from XML:
//
//	elem.OuterXml()            - Complete XML including element
//	elem.InnerXml()            - XML of children only
//	elem.WriteTo(w)            - Write to io.Writer
//	ParseElement(r, factory)   - Parse XML from reader
//
// # Cloning
//
// Elements can be deep or shallow cloned:
//
//	elem.Clone()               - Deep clone with all descendants
//	elem.CloneNode(false)      - Shallow clone without children
package openxml
