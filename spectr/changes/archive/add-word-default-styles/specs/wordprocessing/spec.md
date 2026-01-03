## ADDED Requirements

### Requirement: Default Styles Generation
New Wordprocessing documents MUST contain a `styles.xml` part with essential default styles.

#### Scenario: Create new document has styles
- WHEN a new `wordprocessing.Document` is created via `New()`
- THEN the document package MUST contain a `styles.xml` part
- AND the `styles.xml` part MUST contain `w:docDefaults`
- AND the `styles.xml` part MUST contain a `w:style` with `w:styleId="Normal"`

### Requirement: Styles Part Initialization
The Styles part MUST be capable of initializing itself with standard defaults.

#### Scenario: Initialize default content
- WHEN `StylesPart.InitializeDefault()` is called
- THEN the part data is populated with valid `w:styles` XML
- AND the XML includes `w:docDefaults` for run and paragraph properties
- AND the XML includes a default "Normal" paragraph style
