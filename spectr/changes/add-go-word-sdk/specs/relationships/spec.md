## ADDED Requirements

### Requirement: Part Relationship Management
The system SHALL manage relationships between parts (internal relationships).

#### Scenario: Create internal relationship
- GIVEN an OpenXmlPart
- WHEN AddRelationship(targetPart, relationshipType) is called
- THEN a relationship is created in the part's .rels file
- AND a unique relationship ID is generated

#### Scenario: Create relationship with specific ID
- GIVEN an OpenXmlPart
- WHEN AddRelationship(targetPart, relationshipType, id) is called
- THEN a relationship with the specified ID is created

#### Scenario: Get all relationships
- GIVEN a part with relationships
- WHEN Relationships() is called
- THEN all relationships for this part are returned

#### Scenario: Get relationships by type
- GIVEN a part with various relationships
- WHEN GetRelationshipsByType(relType) is called
- THEN only relationships of that type are returned

#### Scenario: Get relationship by ID
- GIVEN a part with relationships
- WHEN GetRelationship(id) is called
- THEN the relationship with that ID is returned

#### Scenario: Delete relationship
- GIVEN a part with a relationship
- WHEN DeleteRelationship(id) is called
- THEN the relationship is removed from the .rels file

### Requirement: External Relationship Management
The system SHALL support external relationships (hyperlinks, external files).

#### Scenario: Add hyperlink relationship
- GIVEN an OpenXmlPart
- WHEN AddHyperlinkRelationship(uri, isExternal) is called
- THEN an external relationship to the URI is created

#### Scenario: Add external relationship
- GIVEN an OpenXmlPart
- WHEN AddExternalRelationship(relType, externalUri) is called
- THEN an external relationship is created with TargetMode="External"

#### Scenario: Enumerate external relationships
- GIVEN a part with mixed relationships
- WHEN ExternalRelationships() is called
- THEN only external relationships are returned

#### Scenario: Enumerate hyperlink relationships
- GIVEN a part with hyperlink relationships
- WHEN HyperlinkRelationships() is called
- THEN all hyperlink relationships are returned

### Requirement: Data Part Relationships
The system SHALL support relationships to media/data parts.

#### Scenario: Add image relationship
- GIVEN a MainDocumentPart
- WHEN AddImagePart(imageType) is called
- THEN an ImagePart is created with appropriate relationship

#### Scenario: Audio relationship
- GIVEN a part supporting audio
- WHEN AddAudioReferenceRelationship(audioPart) is called
- THEN an audio reference relationship is created

#### Scenario: Video relationship
- GIVEN a part supporting video
- WHEN AddVideoReferenceRelationship(videoPart) is called
- THEN a video reference relationship is created

### Requirement: Relationship Type Constants
The system SHALL define standard relationship type URIs.

#### Scenario: Standard relationship types available
- GIVEN the need to create a relationship
- WHEN referencing relationship types
- THEN constants for all standard types are available:
  - OfficeDocument
  - Styles
  - Settings
  - WebSettings
  - FontTable
  - Numbering
  - Footnotes
  - Endnotes
  - Comments
  - Header
  - Footer
  - Image
  - Hyperlink
  - Theme
  - CustomXml

### Requirement: Relationship Persistence
The system SHALL correctly persist relationships in .rels files.

#### Scenario: Write relationships file
- GIVEN a part with relationships
- WHEN the package is saved
- THEN a properly formatted .rels file is written at the correct location

#### Scenario: Read relationships file
- GIVEN a package with .rels files
- WHEN the package is opened
- THEN relationships are parsed and accessible

#### Scenario: Relationship file location
- GIVEN a part at /word/document.xml
- WHEN relationships exist
- THEN the .rels file is at /word/_rels/document.xml.rels

#### Scenario: Package-level relationships
- GIVEN a package
- WHEN package relationships exist
- THEN they are stored in /_rels/.rels

### Requirement: Relationship ID Uniqueness
The system SHALL ensure relationship IDs are unique within their container.

#### Scenario: Auto-generated IDs are unique
- GIVEN a part
- WHEN multiple relationships are added without explicit IDs
- THEN each receives a unique ID (e.g., rId1, rId2, rId3)

#### Scenario: Duplicate ID prevention
- GIVEN a part with relationship rId1
- WHEN AddRelationship with id="rId1" is attempted
- THEN an error is returned

#### Scenario: ID reuse after deletion
- GIVEN a part where relationship rId1 was deleted
- WHEN a new relationship is added
- THEN rId1 MAY be reused

### Requirement: Relationship Target Resolution
The system SHALL resolve relationship targets to parts or URIs.

#### Scenario: Resolve internal target
- GIVEN an internal relationship
- WHEN the target is resolved
- THEN the target OpenXmlPart is returned

#### Scenario: Resolve external target
- GIVEN an external relationship
- WHEN the target is resolved
- THEN the external URI is returned

#### Scenario: Resolve relative URI
- GIVEN a relationship with relative target URI
- WHEN resolved against the source part
- THEN the absolute part URI is computed correctly
