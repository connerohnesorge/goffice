# Wordprocessing Document Specification (Delta)

## ADDED Requirements

### Requirement: Mail Merge Accessor
The system SHALL provide access to mail merge operations from WordprocessingDocument.

#### Scenario: Access mail merge engine
- GIVEN a WordprocessingDocument (template or document)
- WHEN MailMerge() is called
- THEN a MailMerge instance is returned configured with the document as template

#### Scenario: Fluent API access
- GIVEN a WordprocessingDocument
- WHEN MailMerge().DataSource(csvSource).Execute() is called
- THEN mail merge executes using the document as template

#### Scenario: Template preservation
- GIVEN a WordprocessingDocument used as template
- WHEN MailMerge operations are performed
- THEN the original template document is not modified

#### Scenario: Use with any document type
- GIVEN a WordprocessingDocument of any type (Document, Template, MacroEnabled)
- WHEN MailMerge() is called
- THEN mail merge is available regardless of document type
