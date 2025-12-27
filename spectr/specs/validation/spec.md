# Validation Specification

## Requirements

### Requirement: Validation Framework
The system SHALL provide a validation framework for document structure and content.

#### Scenario: Validate entire document
- GIVEN a WordprocessingDocument
- WHEN Validate() is called
- THEN all parts and elements are validated
- AND a slice of ValidationError is returned

#### Scenario: Validate single element
- GIVEN an OpenXmlElement
- WHEN ValidateElement(element, context) is called
- THEN the element and its children are validated

#### Scenario: Validate with version
- GIVEN a document and target Office version
- WHEN Validate(Office2016) is called
- THEN validation uses Office 2016 schema rules

#### Scenario: Validation context
- GIVEN a validation in progress
- WHEN validators access ValidationContext
- THEN version, errors collected, and settings are available

### Requirement: Schema Validation
The system SHALL validate element structure against OOXML schema.

#### Scenario: Validate child element sequence
- GIVEN an element with specific child order requirements
- WHEN children are in wrong order
- THEN validation error with "unexpected element order" is reported

#### Scenario: Validate required children
- GIVEN an element requiring certain children
- WHEN required child is missing
- THEN validation error with "missing required element" is reported

#### Scenario: Validate element multiplicity
- GIVEN an element allowing only one child of a type
- WHEN multiple children of that type exist
- THEN validation error with "too many elements" is reported

#### Scenario: Validate allowed children
- GIVEN an element with restricted child types
- WHEN an invalid child type is present
- THEN validation error with "unexpected element" is reported

### Requirement: Attribute Validation
The system SHALL validate attribute values against schema constraints.

#### Scenario: Validate required attributes
- GIVEN an element with required attribute
- WHEN the attribute is missing
- THEN validation error with "missing required attribute" is reported

#### Scenario: Validate attribute type
- GIVEN an attribute expecting integer value
- WHEN the value is not a valid integer
- THEN validation error with "invalid attribute value" is reported

#### Scenario: Validate enumeration values
- GIVEN an attribute with enumerated values
- WHEN value is not in allowed set
- THEN validation error with "value not in enumeration" is reported

#### Scenario: Validate attribute range
- GIVEN an attribute with min/max constraints
- WHEN value is outside range
- THEN validation error with "value out of range" is reported

### Requirement: Semantic Validation
The system SHALL validate business rules beyond schema structure.

#### Scenario: Relationship existence constraint
- GIVEN an element referencing a relationship by ID
- WHEN the referenced relationship doesn't exist
- THEN validation error with "referenced relationship not found" is reported

#### Scenario: Unique value constraint
- GIVEN elements requiring unique IDs
- WHEN duplicate IDs exist
- THEN validation error with "duplicate ID" is reported

#### Scenario: Parent type constraint
- GIVEN an element that must have specific parent type
- WHEN parent is wrong type
- THEN validation error is reported

#### Scenario: Mutually exclusive attributes
- GIVEN attributes that cannot coexist
- WHEN both are present
- THEN validation error with "mutually exclusive attributes" is reported

### Requirement: Validation Error Reporting
The system SHALL provide detailed validation error information.

#### Scenario: Error contains element path
- GIVEN a validation error
- WHEN Error.Path is accessed
- THEN the XPath-like path to the element is returned

#### Scenario: Error contains description
- GIVEN a validation error
- WHEN Error.Description is accessed
- THEN a human-readable error message is returned

#### Scenario: Error contains element reference
- GIVEN a validation error
- WHEN Error.Element is accessed
- THEN the actual element with the error is returned

#### Scenario: Error severity levels
- GIVEN validation errors
- WHEN Error.Severity is accessed
- THEN Error or Warning severity is indicated

#### Scenario: Error contains error code
- GIVEN a validation error
- WHEN Error.Code is accessed
- THEN a machine-readable error code is returned

### Requirement: Office Version Support (Office 2016+ Only)
The system SHALL support validation against Office 2016 and later versions only.

#### Scenario: Office 2016 validation
- GIVEN FileFormatVersions.Office2016
- WHEN validation runs
- THEN ECMA-376 5th edition features are valid

#### Scenario: Office 2019 validation
- GIVEN FileFormatVersions.Office2019
- WHEN validation runs
- THEN Office 2019 extensions are also valid

#### Scenario: Office 2021 validation
- GIVEN FileFormatVersions.Office2021
- WHEN validation runs
- THEN Office 2021 extensions are also valid

#### Scenario: Microsoft365 validation
- GIVEN FileFormatVersions.Microsoft365
- WHEN validation runs
- THEN all current Microsoft365 features are valid

#### Scenario: Element version availability
- GIVEN an element introduced in Office 2019
- WHEN validating for Office 2016
- THEN "element not available in this version" error is reported

### Requirement: Compiled Particle Validators
The system SHALL use compiled schema particles for efficient validation.

#### Scenario: Sequence particle validation
- GIVEN a sequence particle (A, B, C)
- WHEN children are validated
- THEN they must appear in that exact order

#### Scenario: Choice particle validation
- GIVEN a choice particle (A | B | C)
- WHEN children are validated
- THEN exactly one of the choices must appear

#### Scenario: All particle validation
- GIVEN an all particle (A & B & C)
- WHEN children are validated
- THEN all must appear but order doesn't matter

#### Scenario: Repetition validation
- GIVEN a particle with minOccurs=1, maxOccurs=unbounded
- WHEN children are validated
- THEN at least one must appear, any number allowed

### Requirement: Custom Constraint Registration
The system SHALL allow registration of custom semantic constraints.

#### Scenario: Register custom constraint
- GIVEN a custom Constraint implementation
- WHEN RegisterConstraint(elementType, constraint) is called
- THEN the constraint is applied during validation

#### Scenario: Custom constraint execution
- GIVEN a registered custom constraint
- WHEN validation runs on matching element
- THEN the custom constraint's Check() method is called

### Requirement: Validation Settings
The system SHALL provide configurable validation settings.

#### Scenario: Set max errors
- GIVEN ValidationSettings with MaxErrors=10
- WHEN validation finds more than 10 errors
- THEN validation stops and returns errors found so far

#### Scenario: Continue on error
- GIVEN ValidationSettings with ContinueOnError=true
- WHEN an error is found
- THEN validation continues to find additional errors

#### Scenario: Skip semantic validation
- GIVEN ValidationSettings with SemanticValidation=false
- WHEN validation runs
- THEN only schema validation is performed

