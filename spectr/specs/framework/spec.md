# Framework Specification

## Requirements

### Requirement: OpenXmlElement Base Interface
The system SHALL provide an Element interface as the base for all XML elements in the document.

#### Scenario: Element identity
- GIVEN an Element instance
- WHEN LocalName() and NamespaceURI() are called
- THEN the element's XML qualified name components are returned

#### Scenario: Element qualified name
- GIVEN an Element instance
- WHEN QName() is called
- THEN an OpenXmlQualifiedName with namespace and local name is returned

#### Scenario: Parent access
- GIVEN an element in a document tree
- WHEN Parent() is called
- THEN the containing parent element is returned (or nil for root)

#### Scenario: Features access
- GIVEN an Element
- WHEN Features() is called
- THEN the element's feature collection is returned

### Requirement: OpenXmlCompositeElement
The system SHALL provide a CompositeElement type for elements that contain child elements.

#### Scenario: Access child elements
- GIVEN a composite element with children
- WHEN Children() is called
- THEN all direct child elements are returned in document order

#### Scenario: Typed child access
- GIVEN a composite element
- WHEN GetElement[T]() is called
- THEN the first child element of type T is returned

#### Scenario: Get all children of type
- GIVEN a composite element with multiple children of same type
- WHEN Elements[T]() is called
- THEN all children of type T are returned in order

#### Scenario: First and last child
- GIVEN a composite element with children
- WHEN FirstChild() and LastChild() are called
- THEN the first and last child elements are returned

#### Scenario: Sibling navigation
- GIVEN an element with siblings
- WHEN NextSibling() and PreviousSibling() are called
- THEN the adjacent sibling elements are returned

### Requirement: Child Element Manipulation
The system SHALL support DOM-style manipulation of child elements.

#### Scenario: Append child
- GIVEN a composite element
- WHEN AppendChild(child) is called
- THEN the child is added at the end of children list
- AND the child's parent is set to this element

#### Scenario: Prepend child
- GIVEN a composite element
- WHEN PrependChild(child) is called
- THEN the child is added at the beginning of children list

#### Scenario: Insert before
- GIVEN a composite element with existing child
- WHEN InsertBefore(newChild, referenceChild) is called
- THEN newChild is inserted immediately before referenceChild

#### Scenario: Insert after
- GIVEN a composite element with existing child
- WHEN InsertAfter(newChild, referenceChild) is called
- THEN newChild is inserted immediately after referenceChild

#### Scenario: Remove child
- GIVEN a composite element with a child
- WHEN RemoveChild(child) is called
- THEN the child is removed from children list
- AND the child's parent is set to nil

#### Scenario: Remove all children
- GIVEN a composite element with multiple children
- WHEN RemoveAllChildren() is called
- THEN all children are removed

#### Scenario: Replace child
- GIVEN a composite element with an existing child
- WHEN ReplaceChild(newChild, oldChild) is called
- THEN oldChild is replaced by newChild at the same position

### Requirement: OpenXmlLeafElement
The system SHALL provide a LeafElement type for elements that contain only text content.

#### Scenario: Get inner text
- GIVEN a leaf element with text content
- WHEN InnerText() is called
- THEN the text content is returned

#### Scenario: Set inner text
- GIVEN a leaf element
- WHEN SetInnerText(value) is called
- THEN the element's text content is updated

### Requirement: Attribute Management
The system SHALL provide attribute access and manipulation on elements.

#### Scenario: Get all attributes
- GIVEN an element with attributes
- WHEN Attributes() is called
- THEN all attributes are returned as OpenXmlAttribute slice

#### Scenario: Get specific attribute
- GIVEN an element with attributes
- WHEN GetAttribute(localName, namespaceUri) is called
- THEN the matching attribute is returned if exists

#### Scenario: Set attribute
- GIVEN an element
- WHEN SetAttribute(attr) is called
- THEN the attribute is added or updated

#### Scenario: Remove attribute
- GIVEN an element with an attribute
- WHEN RemoveAttribute(localName, namespaceUri) is called
- THEN the attribute is removed

#### Scenario: Typed attribute access
- GIVEN an element with typed attribute fields
- WHEN accessing the field (e.g., paragraph.Properties.Justification)
- THEN the strongly-typed value is returned

### Requirement: XML Serialization
The system SHALL serialize elements to and from XML.

#### Scenario: Serialize to XML string
- GIVEN an element tree
- WHEN OuterXml() is called
- THEN the complete XML representation is returned

#### Scenario: Get inner XML
- GIVEN a composite element
- WHEN InnerXml() is called
- THEN the XML of all children (without the element itself) is returned

#### Scenario: Set from XML string
- GIVEN an XML string
- WHEN SetOuterXml(xml) is called
- THEN the element tree is replaced with parsed XML

#### Scenario: Write to writer
- GIVEN an element tree
- WHEN WriteTo(writer) is called
- THEN XML is written to the io.Writer with proper namespace handling

#### Scenario: Parse from reader
- GIVEN an io.Reader with XML content
- WHEN ParseElement(reader) is called
- THEN the element tree is constructed

### Requirement: Element Cloning
The system SHALL support deep and shallow cloning of elements.

#### Scenario: Deep clone
- GIVEN an element with children
- WHEN Clone() is called
- THEN a complete copy including all descendants is created
- AND the clone has no parent

#### Scenario: Clone node shallow
- GIVEN a composite element with children
- WHEN CloneNode(deep=false) is called
- THEN only the element itself is copied (no children)

#### Scenario: Clone independence
- GIVEN a cloned element
- WHEN the clone is modified
- THEN the original element is unchanged

### Requirement: Element Metadata
The system SHALL provide access to element schema metadata.

#### Scenario: Get element metadata
- GIVEN an element type
- WHEN Metadata() is called
- THEN IElementMetadata with schema information is returned

#### Scenario: Metadata contains validators
- GIVEN element metadata
- WHEN Validators() is accessed
- THEN schema validators for this element type are returned

#### Scenario: Metadata contains child info
- GIVEN element metadata
- WHEN Children() is accessed
- THEN allowed child element types and constraints are returned

#### Scenario: Version availability
- GIVEN element metadata
- WHEN Availability() is accessed
- THEN the Office versions supporting this element are returned

### Requirement: OpenXmlPartRootElement
The system SHALL provide a root element type for elements that are part roots.

#### Scenario: Part reference
- GIVEN a part root element
- WHEN Part() is called
- THEN the containing OpenXmlPart is returned

#### Scenario: Save to part
- GIVEN a modified part root element
- WHEN Save() is called
- THEN the element tree is serialized back to the part

#### Scenario: Reload from part
- GIVEN a part root element
- WHEN Reload() is called
- THEN the element tree is re-parsed from the part stream

