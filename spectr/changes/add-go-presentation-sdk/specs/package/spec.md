# Packaging

Delta specification for OPC (Open Packaging Conventions) implementation in Go.

## ADDED Requirements

### Requirement: IPackage Abstraction

The system SHALL provide an `IPackage` interface that abstracts the underlying ZIP-based OPC package.

#### Scenario: Package creation
- WHEN creating a new package
- THEN an empty OPC structure is initialized
- AND `[Content_Types].xml` is created
- AND `_rels/.rels` is created

#### Scenario: Package opening from file
- WHEN opening an existing .pptx file
- THEN the package is loaded
- AND parts are accessible
- AND relationships are parsed

#### Scenario: Package opening from stream
- WHEN opening a package from io.Reader
- THEN the package is loaded into memory
- AND modifications can be made
- AND the original stream is not modified until save

#### Scenario: Package save
- WHEN `Save()` is called
- THEN all modified parts are written
- AND content types are updated
- AND relationships are serialized

#### Scenario: Package close
- WHEN `Close()` is called
- THEN resources are released
- AND pending changes are discarded if not saved
- AND subsequent operations return error

### Requirement: OpenXmlPackage Base Type

The system SHALL provide `OpenXmlPackage` as the base type for document packages.

#### Scenario: Root part access
- WHEN accessing `RootPart()` on a package
- THEN the main document part is returned
- AND nil is returned if not yet created

#### Scenario: File access mode
- WHEN checking `FileOpenAccess()`
- THEN Read, Write, or ReadWrite is returned
- AND operations respect the access mode

#### Scenario: Auto-save behavior
- WHEN `AutoSave()` is true and package is disposed
- THEN changes are automatically saved
- AND when false, explicit save is required

#### Scenario: Data parts collection
- WHEN accessing `DataParts()`
- THEN all media/data parts are returned
- AND new data parts can be created

#### Scenario: Save capability check
- WHEN checking `CanSave()`
- THEN true is returned if package supports saving
- AND false for read-only packages

### Requirement: OpenXmlPart Base Type

The system SHALL provide `OpenXmlPart` as the base type for all document parts.

#### Scenario: Part URI access
- WHEN accessing `Uri()` on a part
- THEN the part's URI within the package is returned

#### Scenario: Content type access
- WHEN accessing `ContentType()` on a part
- THEN the MIME content type is returned

#### Scenario: Part stream access
- WHEN calling `GetStream()` on a part
- THEN an io.ReadWriteCloser is returned
- AND content can be read or written

#### Scenario: Feed data to part
- WHEN calling `FeedData(reader)` on a part
- THEN the part content is replaced
- AND the previous content is discarded

#### Scenario: Root element access
- WHEN accessing `RootElement()` on a typed part
- THEN the parsed root element is returned
- AND lazy loading is performed if needed

#### Scenario: Part parent package
- WHEN accessing `OpenXmlPackage()` on a part
- THEN the owning package is returned

### Requirement: OpenXmlPartContainer

The system SHALL provide part container functionality for packages and parts.

#### Scenario: Add typed part
- WHEN calling `AddNewPart[T](contentType, id)`
- THEN a new part of type T is created
- AND a relationship is created with the specified ID
- AND the part is returned

#### Scenario: Add existing part
- WHEN calling `AddPart[T](part)` with an existing part
- THEN the part is added as a child
- AND the relationship is created
- AND the original part is returned

#### Scenario: Delete part by reference
- WHEN calling `DeletePart(part)`
- THEN the part is removed from the container
- AND the relationship is deleted
- AND the part content is removed from package

#### Scenario: Delete part by ID
- WHEN calling `DeletePart(relationshipId)`
- THEN the part with that relationship ID is removed

#### Scenario: Get part by ID
- WHEN calling `GetPartById(id)`
- THEN the part with that relationship ID is returned
- AND error if not found

#### Scenario: Get parts by type
- WHEN calling `GetPartsOfType[T]()`
- THEN all child parts of type T are returned

#### Scenario: Get ID of part
- WHEN calling `GetIdOfPart(part)`
- THEN the relationship ID is returned

#### Scenario: Change part ID
- WHEN calling `ChangeIdOfPart(part, newId)`
- THEN the relationship ID is updated
- AND references in content must be manually updated

### Requirement: Relationship Management

The system SHALL provide comprehensive relationship management.

#### Scenario: Internal relationship creation
- WHEN adding a part to a container
- THEN an internal relationship is created
- AND the target is a relative URI

#### Scenario: External relationship creation
- WHEN calling `AddExternalRelationship(type, uri, id)`
- THEN an external relationship is created
- AND the target is an absolute URI

#### Scenario: Hyperlink relationship creation
- WHEN calling `AddHyperlinkRelationship(uri, isExternal, id)`
- THEN a hyperlink relationship is created
- AND external links target absolute URIs

#### Scenario: Relationship enumeration
- WHEN accessing relationships on a container
- THEN all relationships are returned
- AND type, ID, target, and mode are accessible

#### Scenario: Relationship deletion
- WHEN deleting a relationship
- THEN it is removed from `_rels/*.rels`
- AND the target part may become orphaned

### Requirement: Content Type Management

The system SHALL manage content types via `[Content_Types].xml`.

#### Scenario: Default content type by extension
- WHEN a part has a standard extension (.xml, .rels)
- THEN the default content type is used
- AND no override is created

#### Scenario: Override content type
- WHEN a part has a specific content type
- THEN an override entry is created
- AND the full part URI is specified

#### Scenario: Content type validation
- WHEN opening a part with wrong content type
- THEN validation fails
- AND appropriate error is returned

### Requirement: Part URI Generation

The system SHALL generate unique part URIs following OPC conventions.

#### Scenario: URI from target path
- WHEN creating a new slide part
- THEN URI follows pattern `/ppt/slides/slideN.xml`
- AND N is the lowest unused number

#### Scenario: URI uniqueness
- WHEN a URI would conflict with existing part
- THEN the number is incremented
- AND uniqueness is guaranteed

#### Scenario: URI validation
- WHEN an invalid URI is specified
- THEN an error is returned
- AND the part is not created

### Requirement: Package Properties

The system SHALL support core, extended, and custom properties.

#### Scenario: Core properties access
- WHEN accessing `PackageProperties()`
- THEN title, author, created, modified, etc. are accessible

#### Scenario: Core properties modification
- WHEN modifying title or author
- THEN changes are persisted on save
- AND `docProps/core.xml` is updated

#### Scenario: Extended properties
- WHEN accessing extended properties
- THEN application, company, version are accessible
- AND `docProps/app.xml` is the source

#### Scenario: Custom properties
- WHEN adding custom properties
- THEN arbitrary key-value pairs can be stored
- AND `docProps/custom.xml` is updated

### Requirement: DataPart and MediaDataPart

The system SHALL support binary data parts for media content.

#### Scenario: Media data part creation
- WHEN calling `CreateMediaDataPart(contentType)`
- THEN a data part is created
- AND it can store binary content

#### Scenario: Media data part with extension
- WHEN calling `CreateMediaDataPart(contentType, extension)`
- THEN the part uses the specified extension
- AND content type is registered

#### Scenario: Media data part by type
- WHEN calling `CreateMediaDataPart(MediaDataPartType.Mp4)`
- THEN appropriate content type and extension are used

#### Scenario: Data part deletion
- WHEN calling `DeletePart(dataPart)`
- THEN the part is removed if not referenced
- AND error if still referenced by relationships
