## ADDED Requirements

### Requirement: Advanced Relationship Querying
The system SHALL support filtering and querying relationships by type, target, and ID with efficient lookup.

#### Scenario: Query relationships by type
- WHEN relationships of a specific type are queried
- THEN all relationships matching the type are returned
- AND the results are iterable without materializing all relationships

### Requirement: External Relationship Support
The system SHALL manage external relationships with URI resolution and validation.

#### Scenario: Create external relationship
- WHEN an external relationship is created
- THEN the target URI is stored
- AND the relationship ID is generated
- AND the relationship is tracked in the package

### Requirement: Copy with Relationship Rewriting
The system SHALL rewrite relationship IDs and targets when copying parts across documents.

#### Scenario: Clone part
- WHEN a part is cloned
- THEN internal relationship IDs are rewritten
- AND external relationships are preserved or mapped
- AND the cloned part is fully functional
