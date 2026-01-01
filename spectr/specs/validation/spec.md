# Validation Specification

## Requirements

### Requirement: Validation Framework
The system SHALL provide a validation framework for document structure and content with version awareness.

#### Scenario: Validate entire document
- GIVEN a WordprocessingDocument
- WHEN Validate(doc) is called
- THEN all parts and elements are validated
- AND version-aware validation is performed
- AND a slice of ValidationError is returned

#### Scenario: Validate single element
- GIVEN an OpenXmlElement
- WHEN ValidateElement(element, context) is called
- THEN the element and its children are validated
- AND version compatibility is checked if context has target version

#### Scenario: Validate with version
- GIVEN a document and target Office version
- WHEN Validate(doc, Office2016) is called
- THEN validation uses Office 2016 schema rules
- AND elements not available in Office 2016 generate errors

#### Scenario: Validation context with version
- GIVEN a validation in progress with target version
- WHEN validators access ValidationContext
- THEN TargetVersion, errors collected, and settings are available

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
The system SHALL provide detailed validation error information including version conflicts.

#### Scenario: Error for version mismatch
- GIVEN a validation error for element version incompatibility
- WHEN Error.Description is accessed
- THEN message includes element name, target version, and required version
- AND format is: "ElementName not available in Office 2010 (introduced in Office 2013)"

#### Scenario: Error contains version info
- GIVEN a validation error
- WHEN Error.Code is accessed
- THEN version-related errors have code "ElementNotAvailableInVersion"

#### Scenario: Version error severity
- GIVEN a version compatibility validation error
- WHEN Error.Severity is accessed
- THEN Error severity is returned (not Warning)

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
The system SHALL provide configurable validation settings including version targeting.

#### Scenario: Set target version
- GIVEN ValidationSettings
- WHEN TargetVersion is set to Office2013
- THEN validation enforces Office 2013 compatibility

#### Scenario: Skip version validation
- GIVEN ValidationSettings with VersionValidation=false
- WHEN validation runs
- THEN version compatibility checks are skipped
- AND only structural validation is performed

#### Scenario: Strict version validation
- GIVEN ValidationSettings with StrictVersionValidation=true
- WHEN validation runs
- THEN even optional extension elements are checked for version compatibility

### Requirement: Version-Aware Validation
The system SHALL validate documents against specific Office version requirements.

#### Scenario: Validate with target version
- GIVEN a WordprocessingDocument
- WHEN Validate(doc, Office2010) is called
- THEN only elements available in Office 2010 or earlier are considered valid

#### Scenario: Element not available in target version
- GIVEN a document with Office 2013 element
- WHEN validating against Office 2010
- THEN ValidationError with "element not available in this version" is reported

#### Scenario: Element available in target version
- GIVEN a document with Office 2010 element
- WHEN validating against Office 2013
- THEN no version-related errors are reported

#### Scenario: Multiple version violations
- GIVEN a document with Office 2013 and Office 2016 elements
- WHEN validating against Office 2010
- THEN ValidationErrors for both incompatible elements are reported

### Requirement: Version Detection
The system SHALL detect the minimum Office version required for a document.

#### Scenario: Detect Office 2007 document
- GIVEN a document using only Office 2007 elements
- WHEN DetectMinimumVersion(doc) is called
- THEN Office2007 is returned

#### Scenario: Detect Office 2010 document
- GIVEN a document containing at least one Office 2010 element
- WHEN DetectMinimumVersion(doc) is called
- THEN Office2010 is returned

#### Scenario: Detect mixed version document
- GIVEN a document with Office 2007, 2010, and 2013 elements
- WHEN DetectMinimumVersion(doc) is called
- THEN Office2013 (the highest version) is returned

#### Scenario: Detect Microsoft 365 document
- GIVEN a document using Microsoft 365 exclusive features
- WHEN DetectMinimumVersion(doc) is called
- THEN Microsoft365 is returned

### Requirement: AlternateContent Validation
The system SHALL validate AlternateContent structure and requirements.

#### Scenario: Valid AlternateContent structure
- GIVEN an AlternateContent with one Choice and one Fallback
- WHEN validated
- THEN no errors are reported

#### Scenario: Missing Choice element
- GIVEN an AlternateContent with only Fallback
- WHEN validated
- THEN ValidationError "AlternateContent must have Choice element" is reported

#### Scenario: Missing Fallback element
- GIVEN an AlternateContent with only Choice
- WHEN validated
- THEN ValidationError "AlternateContent must have Fallback element" is reported

#### Scenario: Multiple Choice elements
- GIVEN an AlternateContent with two Choice elements
- WHEN validated
- THEN ValidationError "AlternateContent may have only one Choice" is reported

#### Scenario: Choice missing Requires attribute
- GIVEN a Choice element without Requires attribute
- WHEN validated
- THEN ValidationError "Choice must have Requires attribute" is reported

#### Scenario: Invalid Requires value
- GIVEN a Choice with Requires="invalid_prefix"
- WHEN validated
- THEN ValidationError "unknown namespace prefix in Requires" is reported

#### Scenario: Valid namespace prefix in Requires
- GIVEN a Choice with Requires="w14"
- WHEN validated
- THEN no errors are reported (w14 is valid Word 2010 prefix)

### Requirement: Extension Schema Validation
The system SHALL validate elements against extension schema rules.

#### Scenario: Validate Office 2010 element structure
- GIVEN a ContentControl element (Office 2010)
- WHEN validated
- THEN Office 2010 schema particle rules are applied

#### Scenario: Validate extension element attributes
- GIVEN an element with Office 2013 attributes
- WHEN validated
- THEN Office 2013 schema attribute rules are enforced

#### Scenario: Validate extension element children
- GIVEN an Office 2016 element with children
- WHEN validated
- THEN Office 2016 schema child sequence rules are checked
