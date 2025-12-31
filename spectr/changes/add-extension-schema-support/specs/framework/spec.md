## ADDED Requirements

### Requirement: AlternateContent Support
The system SHALL support AlternateContent elements for version compatibility.

#### Scenario: Parse AlternateContent block
- GIVEN an XML document with mc:AlternateContent element
- WHEN the document is parsed
- THEN an AlternateContent element with Choice and Fallback children is created

#### Scenario: Select Choice content
- GIVEN an AlternateContent element with Choice requiring Office 2010
- WHEN SelectContent(Office2013) is called
- THEN the Choice child content is returned

#### Scenario: Select Fallback content
- GIVEN an AlternateContent element with Choice requiring Office 2013
- WHEN SelectContent(Office2010) is called
- THEN the Fallback child content is returned

#### Scenario: Write AlternateContent
- GIVEN an AlternateContent element
- WHEN WriteTo(writer, targetVersion) is called
- THEN only the appropriate content (Choice or Fallback) is written based on target version

### Requirement: Choice Element
The system SHALL provide a Choice element for modern version content in AlternateContent blocks.

#### Scenario: Parse Choice with Requires
- GIVEN XML `<mc:Choice Requires="w14">`
- WHEN parsed
- THEN Choice element with Requires="w14" is created

#### Scenario: Check version compatibility
- GIVEN a Choice with Requires="w14" (Office 2010)
- WHEN IsCompatibleWith(Office2013) is called
- THEN true is returned

#### Scenario: Reject incompatible version
- GIVEN a Choice with Requires="w15" (Office 2013)
- WHEN IsCompatibleWith(Office2010) is called
- THEN false is returned

### Requirement: Fallback Element
The system SHALL provide a Fallback element for backward-compatible content.

#### Scenario: Parse Fallback content
- GIVEN XML `<mc:Fallback><w:p/></mc:Fallback>`
- WHEN parsed
- THEN Fallback element containing paragraph is created

#### Scenario: Access Fallback child
- GIVEN a Fallback element with child content
- WHEN FirstChild() is called
- THEN the fallback child element is returned

### Requirement: Unknown Element Preservation
The system SHALL preserve unknown XML elements during roundtrip operations.

#### Scenario: Encounter unknown element
- GIVEN XML with element in unrecognized namespace
- WHEN document is parsed
- THEN an UnknownElement instance is created preserving raw XML

#### Scenario: Write unknown element
- GIVEN an UnknownElement with preserved XML
- WHEN WriteTo(writer) is called
- THEN the original XML is written unchanged

#### Scenario: Unknown element in element tree
- GIVEN a composite element containing an UnknownElement
- WHEN iterating children
- THEN the UnknownElement is included in children list

#### Scenario: Roundtrip unknown elements
- GIVEN a document with unknown namespace elements
- WHEN opened, saved, and reopened
- THEN the unknown elements are byte-for-byte identical

### Requirement: Extension Namespace Support
The system SHALL support Office extension namespaces (Office 2010-2024).

#### Scenario: Register extension namespace
- GIVEN a Word 2010 extension namespace "http://schemas.microsoft.com/office/word/2010/wordml"
- WHEN parsing document
- THEN namespace is registered with prefix "w14"

#### Scenario: Namespace prefix lookup
- GIVEN the NamespacePrefixMap
- WHEN looking up Word 2010 namespace
- THEN prefix "w14" is returned

#### Scenario: Element with extension namespace
- GIVEN an element in namespace "http://schemas.microsoft.com/office/word/2010/wordml"
- WHEN NamespaceURI() is called
- THEN the full extension namespace URI is returned

#### Scenario: Write element with extension prefix
- GIVEN an element in Word 2010 namespace
- WHEN written to XML
- THEN prefix "w14" is used (e.g., `<w14:contentControl>`)

### Requirement: Element Version Metadata
The system SHALL provide version availability information for elements.

#### Scenario: Get element version
- GIVEN an element introduced in Office 2010
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Main schema element version
- GIVEN a Paragraph element (Office 2007)
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2007 is returned

#### Scenario: Extension element version
- GIVEN a ContentControl element (Office 2010)
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Version comparison
- GIVEN two elements with different versions
- WHEN comparing AvailableInVersion() values
- THEN Office2007 < Office2010 < Office2013 < Office2016 < Office2019 < Office2021 < Office2022 < Office2023 < Office2024 < Office2025 < Microsoft365

### Requirement: Namespace Management
The system SHALL manage XML namespaces for main and extension schemas.

#### Scenario: Main namespaces defined
- GIVEN the openxml/namespace.go file
- WHEN accessing NamespaceWordprocessingML constant
- THEN "http://schemas.openxmlformats.org/wordprocessingml/2006/main" is returned

#### Scenario: Extension namespaces defined
- GIVEN the openxml/namespace.go file
- WHEN accessing NamespaceWord2010 constant
- THEN "http://schemas.microsoft.com/office/word/2010/wordml" is returned

#### Scenario: Namespace prefix mapping
- GIVEN the NamespacePrefixMap
- WHEN looking up any registered namespace
- THEN the correct prefix is returned (w, w14, w15, x, x14, p, p14, a, a14, etc.)

#### Scenario: Markup compatibility namespace
- GIVEN the markup compatibility namespace constant
- WHEN accessed
- THEN "http://schemas.openxmlformats.org/markup-compatibility/2006" is returned

## MODIFIED Requirements

### Requirement: Element Metadata
The system SHALL provide comprehensive element metadata including version information.

#### Scenario: Metadata includes version
- GIVEN any Element instance
- WHEN Metadata() is called
- THEN IElementMetadata with AvailableInVersion() method is returned

#### Scenario: Metadata includes namespace
- GIVEN any Element instance
- WHEN Metadata().QName() is called
- THEN QualifiedName with full namespace URI is returned

#### Scenario: Metadata includes validators
- GIVEN any Element instance
- WHEN Metadata().Validators() is called
- THEN slice of Validator including version validation is returned
