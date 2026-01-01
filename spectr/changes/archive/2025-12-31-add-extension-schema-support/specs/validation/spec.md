## ADDED Requirements

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

## MODIFIED Requirements

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
