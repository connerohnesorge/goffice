# LINQ-Style Integration Specification

## Requirements

### Requirement: Element Query API
The system SHALL provide LINQ-style query capabilities for OpenXML elements matching DocumentFormat.OpenXml.Linq functionality.

#### Scenario: Query descendants by type
- GIVEN a composite element with nested descendants
- WHEN Descendants[T]() is called
- THEN all descendant elements of type T are returned in document order

#### Scenario: Query ancestors by type
- GIVEN an element in a document tree
- WHEN Ancestors[T]() is called
- THEN all ancestor elements of type T from parent to root are returned

#### Scenario: Descendants and self
- GIVEN a composite element
- WHEN DescendantsAndSelf[T]() is called
- THEN the element itself (if matches T) and all descendants are returned

#### Scenario: Elements after self
- GIVEN an element with following siblings
- WHEN ElementsAfter() is called
- THEN all sibling elements after this one are returned in order

#### Scenario: Elements before self
- GIVEN an element with preceding siblings
- WHEN ElementsBefore() is called
- THEN all sibling elements before this one are returned in reverse order

### Requirement: XElement Integration
The system SHALL provide bidirectional conversion between OpenXML elements and Go's xml.Token/xml.StartElement types for advanced XML manipulation.

#### Scenario: Convert element to generic XML
- GIVEN an OpenXML element tree
- WHEN ToXML() is called
- THEN a generic XML representation is returned

#### Scenario: Convert generic XML to OpenXML
- GIVEN a generic XML tree with OpenXML namespaces
- WHEN FromXML() is called
- THEN strongly-typed OpenXML elements are created

#### Scenario: Preserve mixed content
- GIVEN XML with mixed text and element content
- WHEN converting to/from generic XML
- THEN all content is preserved including whitespace

### Requirement: Namespace Constants Package
The system SHALL provide a package with typed namespace constants for all OpenXML namespaces matching DocumentFormat.OpenXml.Linq namespace classes (W, X, P, A, etc.).

#### Scenario: Word namespace constants
- GIVEN the linq package
- WHEN accessing W.document, W.body, W.p, W.r, W.t
- THEN xml.Name values with correct namespace and local name are returned

#### Scenario: Excel namespace constants
- GIVEN the linq package
- WHEN accessing X.workbook, X.worksheet, X.sheetData, X.row, X.c
- THEN xml.Name values with correct namespace and local name are returned

#### Scenario: PowerPoint namespace constants
- GIVEN the linq package
- WHEN accessing P.presentation, P.sld, P.cSld, P.spTree, P.sp
- THEN xml.Name values with correct namespace and local name are returned

#### Scenario: DrawingML namespace constants
- GIVEN the linq package
- WHEN accessing A.graphic, A.solidFill, A.ln, A.effectLst
- THEN xml.Name values with correct namespace and local name are returned

### Requirement: Advanced Element Filtering
The system SHALL provide predicate-based filtering capabilities.

#### Scenario: Filter elements by predicate
- GIVEN a collection of elements
- WHEN Where(predicate func(Element) bool) is called
- THEN only elements matching the predicate are returned

#### Scenario: Find first matching element
- GIVEN a composite element
- WHEN FirstOrDefault(predicate) is called
- THEN the first child matching predicate is returned, or nil

#### Scenario: Check if any elements match
- GIVEN a composite element
- WHEN Any(predicate) is called
- THEN true is returned if at least one child matches

#### Scenario: Check if all elements match
- GIVEN a composite element
- WHEN All(predicate) is called
- THEN true is returned if all children match predicate

### Requirement: Element Attribute Queries
The system SHALL provide fluent attribute querying APIs.

#### Scenario: Get attribute value
- GIVEN an element with attributes
- WHEN AttributeValue(name) is called
- THEN the attribute's string value is returned

#### Scenario: Get typed attribute value
- GIVEN an element with typed attributes
- WHEN AttributeValueAs[T](name) is called
- THEN the attribute value converted to type T is returned

#### Scenario: Attribute exists check
- GIVEN an element
- WHEN HasAttribute(name) is called
- THEN true is returned if attribute exists

### Requirement: Element Navigation Extensions
The system SHALL provide convenient navigation helper methods.

#### Scenario: Get single child element
- GIVEN a composite element with exactly one child of type T
- WHEN Element[T]() is called
- THEN the single child is returned

#### Scenario: Reject multiple children
- GIVEN a composite element with multiple children of type T
- WHEN Element[T]() is called
- THEN an error is returned indicating multiple matches

#### Scenario: Element at index
- GIVEN a composite element with multiple children
- WHEN ElementAt(index) is called
- THEN the child at the specified position is returned

### Requirement: Functional Element Construction
The system SHALL support functional construction patterns for element trees.

#### Scenario: Build element tree functionally
- GIVEN element constructors accepting varargs of child elements
- WHEN NewDocument(NewBody(NewParagraph(NewRun(NewText("Hello"))))) is called
- THEN a complete document tree is constructed

#### Scenario: Conditional element inclusion
- GIVEN a constructor with optional child elements
- WHEN passing nil elements
- THEN they are excluded from the tree

#### Scenario: Array expansion
- GIVEN a slice of child elements
- WHEN passed to constructor with varargs
- THEN all elements in slice are added as children

### Requirement: Text Content Aggregation
The system SHALL provide convenient text aggregation across element trees.

#### Scenario: Get inner text recursively
- GIVEN a composite element with nested text elements
- WHEN InnerTextRecursive() is called
- THEN all text content concatenated in document order is returned

#### Scenario: Get text content only
- GIVEN a composite element with mixed content
- WHEN TextOnly() is called
- THEN only direct text nodes are returned, excluding child elements

#### Scenario: Set text content batch
- GIVEN multiple text elements
- WHEN SetTextBatch(values) is called
- THEN all elements are updated with corresponding values

### Requirement: Part Extension for LINQ
The system SHALL provide LINQ-style query methods on OpenXmlPart.

#### Scenario: Query part descendants
- GIVEN a document part with nested elements
- WHEN part.Descendants[T]() is called
- THEN all descendant elements of type T in the part are returned

#### Scenario: Find elements in part by path
- GIVEN a document part
- WHEN part.ElementsByXPath("//w:p/w:r/w:t") is called
- THEN all matching text elements are returned

### Requirement: Package-Level LINQ Queries
The system SHALL provide package-wide query capabilities.

#### Scenario: Query all parts of type
- GIVEN a package with multiple parts
- WHEN package.Parts[T]() is called
- THEN all parts of type T are returned

#### Scenario: Query elements across all parts
- GIVEN a package with multiple document parts
- WHEN package.AllDescendants[T]() is called
- THEN all elements of type T from all parts are returned

#### Scenario: Find parts by relationship type
- GIVEN a package with various relationships
- WHEN package.PartsByRelationshipType(relType) is called
- THEN all parts with that relationship type are returned

### Requirement: Relationship Query API
The system SHALL provide LINQ-style relationship querying.

#### Scenario: Query relationships by type
- GIVEN a part with multiple relationships
- WHEN part.RelationshipsByType(relType) is called
- THEN all relationships of that type are returned

#### Scenario: Query target parts
- GIVEN a part with relationships
- WHEN part.GetRelatedParts[T]() is called
- THEN all related parts of type T are returned

#### Scenario: Find relationship by ID
- GIVEN a part with relationships
- WHEN part.GetRelationship(id) is called
- THEN the relationship with matching ID is returned

### Requirement: Performance Optimization
The system SHALL use lazy evaluation and iterators for large documents.

#### Scenario: Lazy descendant iteration
- GIVEN a large document with thousands of elements
- WHEN Descendants() is called
- THEN an iterator is returned that evaluates elements on demand

#### Scenario: Early termination
- GIVEN a descendant query with predicate
- WHEN FirstOrDefault() is used
- THEN iteration stops after first match without traversing entire tree

#### Scenario: Memory-efficient queries
- GIVEN a package-wide query
- WHEN iterating results
- THEN elements are processed without loading entire document into memory

### Requirement: Error Handling for Queries
The system SHALL provide clear error messages for query operations.

#### Scenario: Type mismatch error
- GIVEN a query for element type that doesn't exist
- WHEN executing query
- THEN descriptive error explaining the type mismatch is returned

#### Scenario: Invalid XPath
- GIVEN an invalid XPath expression
- WHEN ElementsByXPath() is called
- THEN error with XPath syntax details is returned

#### Scenario: Nil element handling
- GIVEN a query on nil element
- WHEN any query method is called
- THEN empty result is returned without panic

### Requirement: Compatibility with Standard Library
The system SHALL integrate with Go's standard library patterns.

#### Scenario: Iterator protocol
- GIVEN query result collections
- WHEN used in range loops
- THEN they work with Go's for...range syntax

#### Scenario: Functional options
- GIVEN query methods
- WHEN functional options are provided
- THEN query behavior is customized (e.g., depth limits, filters)

#### Scenario: Context-aware queries
- GIVEN long-running queries
- WHEN context.Context is provided
- THEN queries respect cancellation and timeouts
