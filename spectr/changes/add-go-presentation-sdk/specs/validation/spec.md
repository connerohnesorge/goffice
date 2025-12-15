# Validation

Delta specification for OpenXML document validation in Go.

## ADDED Requirements

### Requirement: OpenXmlValidator Type

The system SHALL provide `OpenXmlValidator` for validating documents against ECMA-376 schemas.

#### Scenario: Validator creation
- WHEN `NewOpenXmlValidator(fileFormatVersion)` is called
- THEN a validator configured for the specified version is returned
- AND version can be Office2007 through Office2021

#### Scenario: Validate element
- WHEN `Validate(element)` is called
- THEN the element and descendants are validated
- AND a slice of ValidationError is returned

#### Scenario: Validate part
- WHEN `Validate(part)` is called
- THEN the part's root element is validated
- AND part-specific rules are checked

#### Scenario: Validate document
- WHEN `Validate(document)` is called
- THEN all parts are validated
- AND inter-part relationships are verified
- AND document-level rules are checked

#### Scenario: Max errors limit
- WHEN `SetMaxNumberOfErrors(n)` is called
- THEN validation stops after n errors
- AND performance is improved for invalid documents

### Requirement: ValidationError Type

The system SHALL provide detailed validation error information.

#### Scenario: Error properties
- WHEN a validation error is created
- THEN `ErrorType()` returns the category
- AND `Description()` returns human-readable message
- AND `Path()` returns XPath-like location
- AND `Part()` returns the containing part if applicable

#### Scenario: Error types
- WHEN validation fails
- THEN error type is one of:
  - Schema (element/attribute structure)
  - Semantic (business rules)
  - Package (OPC structure)
  - Relationship (reference integrity)

### Requirement: FileFormatVersions

The system SHALL support targeting specific Office versions.

#### Scenario: Office 2007 validation
- WHEN targeting FileFormatVersions.Office2007
- THEN only Office 2007 schema elements are valid
- AND newer elements are flagged as errors

#### Scenario: Office 2021 validation
- WHEN targeting FileFormatVersions.Office2021
- THEN Office 2021 and earlier elements are valid

#### Scenario: Multiple version validation
- WHEN targeting multiple versions (Office2019 | Office2021)
- THEN elements valid in any target version pass

### Requirement: Schema Validation

The system SHALL validate element structure against ECMA-376 schemas.

#### Scenario: Unknown element detection
- WHEN an element has unknown qualified name
- THEN a schema error is reported
- AND the element path is included

#### Scenario: Required attribute validation
- WHEN a required attribute is missing
- THEN a schema error is reported
- AND the attribute name is included

#### Scenario: Attribute value validation
- WHEN an attribute value violates type constraints
- THEN a schema error is reported
- AND the invalid value is included

#### Scenario: Child element sequence validation
- WHEN child elements are in wrong order
- THEN a schema error is reported
- AND expected sequence is indicated

#### Scenario: Child element cardinality validation
- WHEN a required child element is missing
- THEN a schema error is reported
- WHEN too many child elements exist
- THEN a schema error is reported

#### Scenario: Content model validation
- WHEN element content violates xs:choice or xs:sequence
- THEN appropriate error is reported

### Requirement: Semantic Validation

The system SHALL validate business rules beyond schema structure.

#### Scenario: Relationship target validation
- WHEN a relationship targets non-existent part
- THEN a semantic error is reported

#### Scenario: ID uniqueness validation
- WHEN duplicate IDs exist in document
- THEN a semantic error is reported

#### Scenario: Cross-reference validation
- WHEN an element references non-existent ID
- THEN a semantic error is reported

#### Scenario: Presentation-specific rules
- WHEN slide references invalid layout
- THEN a semantic error is reported
- WHEN animation targets non-existent shape
- THEN a semantic error is reported

### Requirement: Package Validation

The system SHALL validate OPC package structure.

#### Scenario: Content types validation
- WHEN [Content_Types].xml is malformed
- THEN a package error is reported

#### Scenario: Relationship file validation
- WHEN _rels/*.rels is malformed
- THEN a package error is reported

#### Scenario: Part existence validation
- WHEN content types reference non-existent parts
- THEN a package error is reported

#### Scenario: Duplicate part URI detection
- WHEN package contains duplicate URIs
- THEN a package error is reported
