# Package Specification

## Requirements

### Requirement: OPC Package Implementation
The system SHALL implement the Open Packaging Conventions (OPC) specification (ECMA-376 Part 2) for reading and writing Office Open XML packages.

#### Scenario: Create new package
- GIVEN a file path or io.Writer
- WHEN Package.Create() is called
- THEN a new empty OPC package is created with [Content_Types].xml initialized

#### Scenario: Open existing package for reading
- GIVEN a path to an existing .docx file
- WHEN Package.Open(path, ReadOnly) is called
- THEN the package is opened and parts are accessible
- AND the underlying ZIP is not held open exclusively

#### Scenario: Open existing package for editing
- GIVEN a path to an existing .docx file
- WHEN Package.Open(path, ReadWrite) is called
- THEN the package is opened for modification
- AND changes can be saved back to the file

#### Scenario: Open package from stream
- GIVEN an io.ReaderAt containing a valid OPC package
- WHEN Package.OpenReader(reader, size) is called
- THEN the package is opened from the stream
- AND the original stream is used (not copied)

#### Scenario: Save package
- GIVEN an open package with modifications
- WHEN Package.Save() is called
- THEN all modified parts are written to the underlying storage
- AND the [Content_Types].xml is updated

#### Scenario: Save package to new location
- GIVEN an open package
- WHEN Package.SaveAs(newPath) is called
- THEN a complete copy is written to the new location
- AND the original file is unchanged

#### Scenario: Close package
- GIVEN an open package
- WHEN Package.Close() is called
- THEN all resources are released
- AND any unsaved changes are discarded

### Requirement: Content Types Management
The system SHALL manage content type mappings in [Content_Types].xml according to OPC specification.

#### Scenario: Default content type by extension
- GIVEN a package with parts
- WHEN a part with extension .xml is added without explicit content type
- THEN the default content type for .xml extension is used

#### Scenario: Override content type for specific part
- GIVEN a package with [Content_Types].xml
- WHEN a part is added with specific content type
- THEN an override entry is created for that part URI

#### Scenario: Query content type for part
- GIVEN a package with parts
- WHEN GetContentType(partURI) is called
- THEN the correct content type is returned (override takes precedence over default)

### Requirement: Package Part Access
The system SHALL provide access to parts within the package by URI.

#### Scenario: Get part by URI
- GIVEN an open package with parts
- WHEN GetPart(uri) is called with valid part URI
- THEN the corresponding PackagePart is returned

#### Scenario: Part not found
- GIVEN an open package
- WHEN GetPart(uri) is called with non-existent URI
- THEN ErrPartNotFound is returned

#### Scenario: Create new part
- GIVEN an open writable package
- WHEN CreatePart(uri, contentType) is called
- THEN a new part is created at the specified URI
- AND the content type is registered

#### Scenario: Delete part
- GIVEN an open writable package with an existing part
- WHEN DeletePart(uri) is called
- THEN the part is removed from the package
- AND its content type entry is removed
- AND relationships to the part are removed

#### Scenario: Enumerate all parts
- GIVEN an open package with multiple parts
- WHEN Parts() is called
- THEN all parts in the package are returned (excluding relationship parts)

### Requirement: Package Part Streams
The system SHALL provide stream access to part content.

#### Scenario: Read part content
- GIVEN an open package with a part
- WHEN part.GetStream() is called
- THEN an io.Reader for the part content is returned

#### Scenario: Write part content
- GIVEN an open writable package with a part
- WHEN part.GetStream() is called for writing
- THEN an io.Writer for the part content is returned

#### Scenario: Stream mode options
- GIVEN a part in a package
- WHEN GetStream(mode) is called with Read, Write, or ReadWrite
- THEN the appropriate stream type is returned based on mode

### Requirement: Package Relationship Management
The system SHALL manage package-level relationships in _rels/.rels.

#### Scenario: Create package relationship
- GIVEN an open writable package
- WHEN CreateRelationship(targetUri, relationshipType, id) is called
- THEN a relationship is added to _rels/.rels

#### Scenario: Get relationships by type
- GIVEN a package with relationships
- WHEN GetRelationshipsByType(relType) is called
- THEN all relationships of that type are returned

#### Scenario: Get relationship by ID
- GIVEN a package with relationships
- WHEN GetRelationship(id) is called
- THEN the relationship with that ID is returned

#### Scenario: Delete relationship
- GIVEN a package with a relationship
- WHEN DeleteRelationship(id) is called
- THEN the relationship is removed from _rels/.rels

### Requirement: Part URI Handling
The system SHALL correctly handle OPC part URIs according to specification.

#### Scenario: Normalize part URI
- GIVEN a part URI "/word/document.xml"
- WHEN the URI is processed
- THEN it is normalized to lowercase with forward slashes

#### Scenario: Resolve relative URI
- GIVEN a base part URI "/word/document.xml" and relative URI "media/image1.png"
- WHEN ResolvePartUri(base, relative) is called
- THEN "/word/media/image1.png" is returned

#### Scenario: Validate part URI
- GIVEN a string that may be a part URI
- WHEN ValidatePartUri(uri) is called
- THEN true is returned if valid OPC part URI, false otherwise

### Requirement: Package Properties Access
The system SHALL provide access to OPC core properties.

#### Scenario: Read core properties
- GIVEN a package with core properties part
- WHEN package.CoreProperties() is called
- THEN creator, title, subject, description, etc. are accessible

#### Scenario: Write core properties
- GIVEN an open writable package
- WHEN CoreProperties().SetCreator("name") is called
- THEN the core properties part is updated on save

### Requirement: OpenXmlPart Base Type
The system SHALL provide an abstract OpenXmlPart type that serves as the base for all document parts.

#### Scenario: Part identity
- GIVEN an OpenXmlPart instance
- WHEN URI(), ContentType(), and RelationshipId() are called
- THEN the part's unique identifiers are returned

#### Scenario: Part root element access
- GIVEN an OpenXmlPart with XML content
- WHEN RootElement() is called
- THEN the parsed root element is returned (lazily loaded)

#### Scenario: Part stream access
- GIVEN an OpenXmlPart
- WHEN GetStream(mode) is called
- THEN access to the underlying part stream is provided

#### Scenario: Part features
- GIVEN an OpenXmlPart
- WHEN Features() is called
- THEN the part's feature collection is returned (inheriting from package)

### Requirement: OpenXmlPartContainer
The system SHALL provide a container type for managing child parts and relationships.

#### Scenario: Access child parts
- GIVEN an OpenXmlPartContainer with child parts
- WHEN Parts() is called
- THEN all direct child parts are returned with their relationship IDs

#### Scenario: Get part by relationship ID
- GIVEN an OpenXmlPartContainer with child parts
- WHEN GetPartById(id) is called
- THEN the part with that relationship ID is returned

#### Scenario: Get parts by relationship type
- GIVEN an OpenXmlPartContainer with parts of various types
- WHEN GetPartsOfType[T]() is called
- THEN all parts of type T are returned

#### Scenario: Add child part
- GIVEN a writable OpenXmlPartContainer
- WHEN AddPart(part, id) is called
- THEN the part is added as a child with a relationship

#### Scenario: Delete child part
- GIVEN an OpenXmlPartContainer with a child part
- WHEN DeletePart(id) is called
- THEN the part and its relationship are removed

### Requirement: Typed Part Registration
The system SHALL maintain a registry of part types by content type and relationship type.

#### Scenario: Register part type
- GIVEN a part type with known content type
- WHEN the system initializes
- THEN the part type is registered for automatic instantiation

#### Scenario: Create part by content type
- GIVEN a registered content type
- WHEN a part with that content type is encountered during load
- THEN the correct concrete part type is instantiated

#### Scenario: Unknown part type handling
- GIVEN an unknown content type
- WHEN loading a part with that type
- THEN a generic OpenXmlPart is created with raw access

### Requirement: Part Lifecycle Management
The system SHALL manage part creation, loading, and saving.

#### Scenario: Lazy loading
- GIVEN an opened package with many parts
- WHEN parts are enumerated
- THEN part XML is not parsed until RootElement() is accessed

#### Scenario: Save modified part
- GIVEN a part with modified root element
- WHEN the package is saved
- THEN the part's XML is serialized from the element tree

#### Scenario: Track part modifications
- GIVEN a part
- WHEN its root element is modified
- THEN the part is marked as modified for saving

### Requirement: Fixed Content Type Parts
The system SHALL support parts with invariant content types (IFixedContentTypePart pattern).

#### Scenario: Fixed content type enforcement
- GIVEN a part type with fixed content type (e.g., MainDocumentPart)
- WHEN the part is created
- THEN only the fixed content type is allowed

#### Scenario: Content type validation on load
- GIVEN a package being loaded
- WHEN a fixed content type part has wrong content type
- THEN a validation error is reported

### Requirement: Part URI Generation
The system SHALL generate valid URIs for new parts.

#### Scenario: Default URI generation
- GIVEN a part type being added
- WHEN no explicit URI is provided
- THEN a default URI is generated based on part type conventions

#### Scenario: Unique URI generation
- GIVEN multiple parts of the same type
- WHEN URIs are generated
- THEN unique URIs are created (e.g., /word/media/image1.png, /word/media/image2.png)

#### Scenario: Custom URI
- GIVEN a part being added with explicit URI
- WHEN AddPart is called with custom URI
- THEN the specified URI is used if valid

