## ADDED Requirements

### Requirement: Office 2010 Extension Elements
The system SHALL provide element types for Word 2010 extensions.

#### Scenario: ContentControl element
- GIVEN Office 2010 w14:contentControl element
- WHEN parsed from XML
- THEN ContentControl element instance is created in w14 namespace

#### Scenario: CustomXmlConflictInsertionRangeEnd element
- GIVEN Office 2010 w14:customXmlConflictInsRangeEnd element
- WHEN accessed via API
- THEN strongly-typed CustomXmlConflictInsertionRangeEnd element is available

#### Scenario: DocId element
- GIVEN Office 2010 w14:docId element
- WHEN present in document
- THEN DocId element with Val attribute is accessible

#### Scenario: Drawing Canvas elements
- GIVEN Office 2010 drawing canvas elements (wpc namespace)
- WHEN parsed
- THEN WordprocessingCanvas, WordprocessingGroup, WordprocessingShape elements available

### Requirement: Office 2012-2013 Extension Elements
The system SHALL provide element types for Word 2012-2013 extensions.

#### Scenario: WebExtension elements
- GIVEN Office 2013 webextension elements (we namespace)
- WHEN parsed
- THEN WebExtension, WebExtensionReference, WebExtensionProperty elements available

#### Scenario: Chart Style elements
- GIVEN Office 2013 chart style elements (w15 namespace)
- WHEN parsed
- THEN ChartStyle, ChartColor elements available

### Requirement: Office 2015-2016 Extension Elements
The system SHALL provide element types for Word 2015-2016 extensions.

#### Scenario: Appearance elements
- GIVEN Office 2015 appearance elements (w15 symex namespace)
- WHEN parsed
- THEN SdtAppearance element with Val attribute available

#### Scenario: Comments Extended elements
- GIVEN Office 2016 comments elements (w16cex namespace)
- WHEN parsed
- THEN CommentExtensible element available

#### Scenario: Comments ID elements
- GIVEN Office 2016 comment ID elements (w16cid namespace)
- WHEN parsed
- THEN CommentId, ParagraphId elements available

### Requirement: Office 2018-2020 Extension Elements
The system SHALL provide element types for Word 2018-2020 extensions.

#### Scenario: 2018 extension elements
- GIVEN Office 2018 elements (w18 namespace)
- WHEN parsed
- THEN Word 2018 extension elements available

#### Scenario: SDT Data Hash elements
- GIVEN Office 2020 SDT data hash elements (w20sdtdh namespace)
- WHEN parsed
- THEN SdtDataHash element available

### Requirement: Office 2023-2024 Extension Elements
The system SHALL provide element types for Word 2023-2024 extensions.

#### Scenario: Word 16 Document Undo elements
- GIVEN Office 2023 elements (w16du namespace)
- WHEN parsed
- THEN Word 16 document undo elements available

#### Scenario: SDT Format Lock elements
- GIVEN Office 2024 SDT format lock elements (w24sdtfl namespace)
- WHEN parsed
- THEN SdtFormatLock element with formatting lock properties available

### Requirement: Extension Element Attributes
The system SHALL provide typed attributes for extension elements.

#### Scenario: Extension string attributes
- GIVEN an extension element with string attribute
- WHEN attribute is accessed
- THEN StringValue wrapper is returned

#### Scenario: Extension enum attributes
- GIVEN an extension element with enumeration attribute
- WHEN attribute is accessed
- THEN enum value from extension schema is returned

#### Scenario: Extension boolean attributes
- GIVEN an extension element with boolean attribute
- WHEN attribute is accessed
- THEN BoolValue wrapper is returned

### Requirement: Extension Element Namespace Handling
The system SHALL correctly handle namespaces for extension elements.

#### Scenario: Element qualified name
- GIVEN a ContentControl element (w14 namespace)
- WHEN QName() is called
- THEN QualifiedName with namespace "http://schemas.microsoft.com/office/word/2010/wordml" is returned

#### Scenario: Element local name
- GIVEN a ContentControl element
- WHEN LocalName() is called
- THEN "contentControl" is returned

#### Scenario: Namespace prefix in XML
- GIVEN a ContentControl element
- WHEN written to XML
- THEN element is serialized as `<w14:contentControl>`

### Requirement: Extension Element Version Metadata
The system SHALL provide version information for extension elements.

#### Scenario: Office 2010 element version
- GIVEN a ContentControl element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Office 2024 element version
- GIVEN an SdtFormatLock element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2024 is returned

### Requirement: Element Generation from Extension Schemas
The system SHALL generate element types from extension schemas in addition to main schemas.

#### Scenario: Generate from main schema
- GIVEN wordprocessingml main.json schema
- WHEN generator runs
- THEN core WordprocessingML elements are generated

#### Scenario: Generate from extension schemas
- GIVEN Word 2010-2024 extension schemas
- WHEN generator runs
- THEN extension element types are generated with correct namespaces and version metadata

#### Scenario: Total element count
- GIVEN all schemas loaded
- WHEN generation completes
- THEN 680-700+ element types exist (baseline ~619 + 60-80 extensions)

#### Scenario: Generated element namespaces
- GIVEN generated extension elements
- WHEN inspecting code
- THEN elements use correct namespace constants (NamespaceWord2010, NamespaceWord2012, etc.)

### Requirement: Extension Element Type Registration
The system SHALL register all extension element types for XML parsing.

#### Scenario: Register extension elements
- GIVEN generated extension element types
- WHEN element registry is initialized
- THEN all extension elements are registered with their qualified names

#### Scenario: Parse registered extension element
- GIVEN XML with w14:contentControl element
- WHEN parsing document
- THEN ContentControl instance is created (not UnknownElement)

#### Scenario: Parse unregistered element
- GIVEN XML with unknown namespace element
- WHEN parsing document
- THEN UnknownElement instance is created preserving raw XML
