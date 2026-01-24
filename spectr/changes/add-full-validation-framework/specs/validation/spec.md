## ADDED Requirements

### Requirement: Schema-Based Validation
The system SHALL validate document structure against ECMA-376 schema with support for transitional and strict modes.

#### Scenario: Validate element type
- WHEN an element is validated
- THEN its type is checked against schema
- AND attributes are validated for type and format
- AND required attributes are verified

#### Scenario: Validate cardinality
- WHEN child elements are validated
- THEN the count of each child type is verified
- AND minOccurs and maxOccurs constraints are enforced
- AND ordering constraints are checked

### Requirement: Semantic Validation
The system SHALL validate element relationships and constraints beyond schema structure.

#### Scenario: Validate relationship
- WHEN elements are related
- THEN the relationship type is validated
- AND both elements are verified to support the relationship
- AND circular dependencies are detected

### Requirement: Validation Error Reporting
The system SHALL collect and report validation errors with detailed context and severity levels.

#### Scenario: Collect validation errors
- WHEN validation runs
- THEN all errors are collected
- AND each error includes element location, type, and message
- AND severity levels distinguish critical vs. advisory issues

### Requirement: Repair Mode
The system SHALL automatically fix common validation issues and suggest manual repairs.

#### Scenario: Auto-repair
- WHEN repair mode is enabled
- THEN missing required elements are added with defaults
- THEN malformed content is corrected
- THEN corrected document is verified

#### Scenario: Repair suggestions
- WHEN validation fails
- THEN suggestions for repair are provided
- AND user can choose to apply fixes
- AND results are reported after repair application
