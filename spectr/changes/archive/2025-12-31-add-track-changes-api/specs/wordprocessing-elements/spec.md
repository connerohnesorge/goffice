# Delta Spec: Wordprocessing Elements - Track Changes High-Level API

## MODIFIED Requirements

### Requirement: Revision Elements
The system SHALL provide elements for tracked changes with high-level wrapper API for document-level operations.

**Note**: This modifies the existing "Revision Elements" requirement in `spectr/specs/wordprocessing-elements/spec.md` to add the high-level wrapper context.

#### Scenario: InsertedRun (w:ins)
- GIVEN an InsertedRun element
- WHEN Author(), Date(), and content are accessed
- THEN the insertion revision information is returned
- AND the element can be wrapped in a high-level Revision object
- AND the Revision wrapper provides Accept() and Reject() methods

#### Scenario: DeletedRun (w:del)
- GIVEN a DeletedRun element
- WHEN content is accessed
- THEN the deleted content and revision info are returned
- AND the element can be wrapped in a high-level Revision object
- AND the Revision wrapper provides Accept() and Reject() methods

#### Scenario: DeletedText
- GIVEN a DeletedText element
- WHEN InnerText() is called
- THEN the deleted text content is returned
- AND the text can be restored to normal Text during reject operations

#### Scenario: MoveFromRun
- GIVEN a MoveFromRun element with id=10
- WHEN wrapped in a Revision object
- THEN the revision metadata (id, author, date) is accessible
- AND the Revision can be paired with matching MoveToRun id=10
- AND Accept/Reject operations coordinate with the move pair

#### Scenario: MoveToRun
- GIVEN a MoveToRun element with id=10
- WHEN wrapped in a Revision object
- THEN the revision metadata (id, author, date) is accessible
- AND the Revision can be paired with matching MoveFromRun id=10
- AND Accept/Reject operations coordinate with the move pair

#### Scenario: RunPropertiesChange
- GIVEN a RunPropertiesChange element
- WHEN wrapped in a Revision object
- THEN the previous properties and new properties are accessible
- AND Accept applies new properties and removes wrapper
- AND Reject restores previous properties and removes wrapper

#### Scenario: ParagraphPropertiesChange
- GIVEN a ParagraphPropertiesChange element
- WHEN wrapped in a Revision object
- THEN the previous properties and new properties are accessible
- AND Accept applies new properties and removes wrapper
- AND Reject restores previous properties and removes wrapper

## ADDED Requirements

### Requirement: Revision Type Enumeration
The system SHALL provide type identification for all revision elements.

#### Scenario: Identify insertion type
- GIVEN an InsertedRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeInsert

#### Scenario: Identify deletion type
- GIVEN a DeletedRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeDelete

#### Scenario: Identify move-from type
- GIVEN a MoveFromRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeMoveFrom

#### Scenario: Identify move-to type
- GIVEN a MoveToRun element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeMoveTo

#### Scenario: Identify format change type
- GIVEN a RunPropertiesChange or ParagraphPropertiesChange element
- WHEN wrapped as Revision
- THEN Type() returns RevisionTypeFormatChange

### Requirement: Revision Content Extraction
The system SHALL provide unified content access across revision types.

#### Scenario: Extract insertion content
- GIVEN an InsertedRun containing runs with text "new content"
- WHEN wrapped as Revision and Content() is called
- THEN "new content" is returned

#### Scenario: Extract deletion content
- GIVEN a DeletedRun containing DeletedText "removed content"
- WHEN wrapped as Revision and Content() is called
- THEN "removed content" is returned

#### Scenario: Extract move content
- GIVEN a MoveFromRun containing runs with text "moved text"
- WHEN wrapped as Revision and Content() is called
- THEN "moved text" is returned

#### Scenario: Extract format change description
- GIVEN a RunPropertiesChange element
- WHEN wrapped as Revision and Content() is called
- THEN a description of the property change is returned

### Requirement: Revision Element Traversal
The system SHALL support finding all revision elements in a document tree.

#### Scenario: Find revisions in paragraph
- GIVEN a Paragraph containing 2 InsertedRun and 1 DeletedRun
- WHEN the paragraph is traversed for revisions
- THEN all 3 revision elements are found

#### Scenario: Find revisions in nested elements
- GIVEN a Table containing cells with InsertedRun elements
- WHEN the table is traversed for revisions
- THEN all InsertedRun elements in all cells are found

#### Scenario: Find revisions in headers
- GIVEN a HeaderPart containing tracked changes
- WHEN the header is traversed for revisions
- THEN all revision elements in the header are found

#### Scenario: Find revisions in footers
- GIVEN a FooterPart containing tracked changes
- WHEN the footer is traversed for revisions
- THEN all revision elements in the footer are found

#### Scenario: Find revisions in comments
- GIVEN a CommentsPart with tracked changes in comment text
- WHEN the comments are traversed for revisions
- THEN all revision elements in comments are found

### Requirement: Revision Iterator Pattern
The system SHALL use Go 1.23+ iter.Seq for memory-efficient revision enumeration.

#### Scenario: Iterate revisions without loading all
- GIVEN a document with 10,000 tracked changes
- WHEN GetRevisions() returns iter.Seq[*Revision]
- THEN memory usage scales with iteration, not total count
- AND early break stops further enumeration

#### Scenario: Iterate and filter
- GIVEN an iterator over revisions
- WHEN filtering by author during iteration
- THEN only matching revisions are processed
- AND non-matching revisions are skipped efficiently

#### Scenario: Empty revision iterator
- GIVEN a document with no tracked changes
- WHEN GetRevisions() is called
- THEN the iterator completes immediately
- AND no allocations occur for empty results
