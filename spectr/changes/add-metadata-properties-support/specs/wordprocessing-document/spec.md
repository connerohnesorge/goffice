## ADDED Requirements

### Requirement: Core Document Properties Access
The system SHALL provide read/write access to core document properties including title, subject, creator, created date, modified date, category, keywords, and description.

#### Scenario: Set and retrieve core properties
- WHEN a Word document is opened
- THEN core properties can be read and modified
- AND changes persist across save/reload cycles

#### Scenario: Automatic timestamps
- WHEN a document is created or modified
- THEN created and modified timestamps are automatically set
- AND lastModifiedBy is tracked

### Requirement: Extended Properties Management
The system SHALL support extended properties including application info, statistics (page count, word count), and document-level metadata.

#### Scenario: Read document statistics
- WHEN a document with statistics is opened
- THEN word count, page count, and character count can be read
- AND values remain in sync with content

#### Scenario: Application metadata
- WHEN properties are accessed
- THEN application name, version, and creation context are available
- AND compatible with Microsoft Office tracking

### Requirement: Custom Properties
The system SHALL support custom typed properties (string, integer, float, boolean, date) with full validation.

#### Scenario: Create and manage custom properties
- WHEN custom properties are added
- THEN they are stored with type information
- AND can be retrieved with correct type conversion

#### Scenario: Property type validation
- WHEN custom properties are set
- THEN type constraints are enforced
- AND invalid types are rejected with clear errors
