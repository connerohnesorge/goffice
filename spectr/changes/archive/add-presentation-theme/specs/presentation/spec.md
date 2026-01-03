## ADDED Requirements

### Requirement: Default Theme Generation
New Presentation documents MUST contain a `theme1.xml` part to ensure valid visual styling.

#### Scenario: New presentation has theme
- WHEN a new `presentation.Document` is created via `New()`
- THEN the document package MUST contain a `theme` part (e.g., `/ppt/theme/theme1.xml`)
- AND the `SlideMaster` MUST reference this theme via a relationship

### Requirement: Slide Master Theme Linkage
Every Slide Master MUST have a relationship to a Theme part.

#### Scenario: Master links to theme
- WHEN `AddSlideMaster()` is called
- THEN the resulting `SlideMasterPart` MUST have a relationship of type `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`
