# Drawingml Core Specification (Delta)

## ADDED Requirements

### Requirement: Diagram Element Support

The system SHALL provide generated element classes for SmartArt diagram elements from the dgm: namespace.

#### Scenario: Diagram namespace defined
- GIVEN the DrawingML package
- WHEN namespace constants are accessed
- THEN NamespaceDiagram constant is available with value "http://schemas.openxmlformats.org/drawingml/2006/diagram"

#### Scenario: Diagram package exists
- GIVEN the goffice project
- WHEN looking for diagram support
- THEN drawingml/diagram package exists with generated element classes

#### Scenario: Generated diagram elements compile
- GIVEN the drawingml/diagram package
- WHEN the package is built
- THEN all generated element classes compile without errors
- AND golangci-lint passes on generated code

### Requirement: Diagram Namespace Integration

The system SHALL integrate diagram namespace into the DrawingML namespace mapping system.

#### Scenario: Diagram prefix mapping
- GIVEN a diagram namespace URI
- WHEN GetPrefixForNamespace is called
- THEN the prefix "dgm" is returned

#### Scenario: Namespace reverse mapping
- GIVEN the prefix "dgm"
- WHEN GetNamespaceForPrefix is called
- THEN the diagram namespace URI is returned
