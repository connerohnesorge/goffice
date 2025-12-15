# Core Framework

Delta specification for the core OpenXML element framework in Go.

## ADDED Requirements

### Requirement: OpenXmlElement Base Type

The system SHALL provide an `OpenXmlElement` interface and `BaseElement` struct that serve as the foundation for all OpenXML elements, supporting DOM-style navigation, attribute management, and XML serialization.

#### Scenario: Element creation with qualified name
- WHEN a new element is created with namespace URI and local name
- THEN the element's `NamespaceUri()` returns the specified namespace
- AND the element's `LocalName()` returns the specified local name
- AND the element's `Prefix()` returns the default prefix for the namespace

#### Scenario: Parent-child navigation
- WHEN an element is added as a child of another element
- THEN the child's `Parent()` returns the parent element
- AND the parent's `FirstChild()` returns the first child
- AND the parent's `LastChild()` returns the last child
- AND the parent's `HasChildren()` returns true

#### Scenario: Sibling navigation
- WHEN multiple elements are children of the same parent
- THEN `NextSibling()` returns the following sibling or nil
- AND `PreviousSibling()` returns the preceding sibling or nil

#### Scenario: Ancestor traversal
- WHEN calling `Ancestors()` on a deeply nested element
- THEN an iterator yields all ancestors from parent to root
- AND calling `Ancestors[T]()` filters to specific types

#### Scenario: Descendant traversal
- WHEN calling `Descendants()` on a composite element
- THEN an iterator yields all descendants in document order
- AND calling `Descendants[T]()` filters to specific types

### Requirement: OpenXmlCompositeElement Type

The system SHALL provide a `CompositeElement` struct that supports child element management with linked-list storage for efficient insertion and removal.

#### Scenario: Append child element
- WHEN `AppendChild(child)` is called
- THEN the child is added as the last child
- AND the child's parent is set to the composite element
- AND the method returns the appended child

#### Scenario: Prepend child element
- WHEN `PrependChild(child)` is called
- THEN the child is added as the first child
- AND existing children's order is preserved

#### Scenario: Insert before reference
- WHEN `InsertBefore(newChild, referenceChild)` is called
- THEN the new child is inserted immediately before the reference
- AND if referenceChild is nil, behaves like AppendChild

#### Scenario: Insert after reference
- WHEN `InsertAfter(newChild, referenceChild)` is called
- THEN the new child is inserted immediately after the reference
- AND if referenceChild is nil, behaves like PrependChild

#### Scenario: Remove child element
- WHEN `RemoveChild(child)` is called
- THEN the child is removed from the children list
- AND the child's parent is set to nil
- AND the removed child is returned

#### Scenario: Remove all children
- WHEN `RemoveAllChildren()` is called
- THEN all children are removed
- AND `HasChildren()` returns false

#### Scenario: Remove children by type
- WHEN `RemoveAllChildren[T]()` is called
- THEN only children of type T are removed
- AND other children remain

### Requirement: OpenXmlLeafElement Type

The system SHALL provide a `LeafElement` struct for elements that cannot have child elements.

#### Scenario: Leaf element child rejection
- WHEN `AppendChild()` is called on a leaf element
- THEN an error is returned
- AND no child is added

#### Scenario: Leaf element HasChildren
- WHEN checking `HasChildren()` on a leaf element
- THEN it always returns false

### Requirement: OpenXmlLeafTextElement Type

The system SHALL provide a `LeafTextElement` struct for elements that contain only text content.

#### Scenario: Text content access
- WHEN setting `Text` property on a leaf text element
- THEN `InnerText()` returns the same value
- AND the XML serialization includes the text content

#### Scenario: Empty text element
- WHEN `Text` is set to empty string
- THEN `InnerText()` returns empty string
- AND `HasValue()` returns true

### Requirement: Element Cloning

The system SHALL support cloning elements with optional deep copy of descendants.

#### Scenario: Shallow clone
- WHEN `CloneNode(false)` is called
- THEN a copy of the element is returned
- AND attributes are copied
- AND children are not copied

#### Scenario: Deep clone
- WHEN `CloneNode(true)` is called
- THEN a copy of the element is returned
- AND attributes are copied
- AND all descendants are recursively cloned

### Requirement: XML Serialization

The system SHALL support XML serialization and deserialization of elements.

#### Scenario: OuterXml generation
- WHEN accessing `OuterXml()` on an element
- THEN the complete XML representation is returned
- AND namespace declarations are included
- AND attributes are properly serialized

#### Scenario: InnerXml generation
- WHEN accessing `InnerXml()` on a composite element
- THEN only the children's XML is returned
- AND the element's own tags are not included

#### Scenario: WriteTo XML writer
- WHEN `WriteTo(writer)` is called
- THEN the element is serialized to the XML writer
- AND namespaces are properly declared
- AND child elements are recursively written

#### Scenario: Lazy parsing
- WHEN an element is loaded with raw XML
- THEN parsing is deferred until content is accessed
- AND accessing properties triggers parsing
- AND `OuterXml()` returns the raw XML if not yet parsed

### Requirement: Attribute Management

The system SHALL provide comprehensive attribute management for fixed (schema-defined) and extended attributes.

#### Scenario: Get attribute by name
- WHEN `GetAttribute(localName, namespaceUri)` is called
- THEN the attribute value is returned if present
- AND an error is returned if not found

#### Scenario: Set fixed attribute
- WHEN `SetAttribute(attr)` is called with a known attribute
- THEN the fixed attribute value is updated
- AND the element is marked as modified

#### Scenario: Set extended attribute
- WHEN `SetAttribute(attr)` is called with an unknown attribute
- THEN the attribute is added to extended attributes
- AND it is serialized in the output XML

#### Scenario: Remove attribute
- WHEN `RemoveAttribute(localName, namespaceUri)` is called
- THEN the attribute is removed
- AND subsequent `GetAttribute()` returns error

#### Scenario: Clear all attributes
- WHEN `ClearAllAttributes()` is called
- THEN all fixed attributes are unset
- AND all extended attributes are removed

### Requirement: Namespace Declaration Management

The system SHALL support adding and removing namespace declarations on elements.

#### Scenario: Add namespace declaration
- WHEN `AddNamespaceDeclaration(prefix, uri)` is called
- THEN the namespace is available for child elements
- AND `xmlns:prefix="uri"` appears in serialization

#### Scenario: Remove namespace declaration
- WHEN `RemoveNamespaceDeclaration(prefix)` is called
- THEN the namespace declaration is removed
- AND child elements may need re-declaration

#### Scenario: Lookup namespace
- WHEN `LookupNamespace(prefix)` is called
- THEN the namespace URI is returned
- AND ancestor declarations are searched

#### Scenario: Lookup prefix
- WHEN `LookupPrefix(namespaceUri)` is called
- THEN the prefix is returned
- AND ancestor declarations are searched

### Requirement: Markup Compatibility Support

The system SHALL support markup compatibility attributes for version-tolerant document handling.

#### Scenario: MCAttributes property
- WHEN `MCAttributes()` is accessed
- THEN markup compatibility attributes are returned
- AND they can be modified

#### Scenario: Ignorable namespace handling
- WHEN mc:Ignorable is specified
- THEN elements from ignorable namespaces are skipped during processing

#### Scenario: MustUnderstand validation
- WHEN mc:MustUnderstand specifies unknown namespaces
- THEN validation fails with appropriate error

### Requirement: Feature Collection

The system SHALL provide a feature collection pattern for extensibility.

#### Scenario: Get feature
- WHEN `Features().Get[T]()` is called
- THEN the registered feature of type T is returned
- AND nil is returned if not registered

#### Scenario: Set feature
- WHEN `Features().Set[T](feature)` is called
- THEN the feature is registered
- AND subsequent Get[T] returns it

#### Scenario: Feature inheritance
- WHEN a feature is not set on an element
- THEN the parent's feature collection is searched
- AND the package's feature collection is searched
