# Delta Spec: Wordprocessing Document - Track Changes API

## ADDED Requirements

### Requirement: Revision Enumeration
The system SHALL provide methods to enumerate all tracked changes in a document.

#### Scenario: Get all revisions from document
- GIVEN a WordprocessingDocument with 5 tracked changes (3 insertions, 2 deletions)
- WHEN GetRevisions() is called
- THEN an iterator over all 5 Revision wrappers is returned
- AND each Revision provides access to metadata and operations

#### Scenario: Get revisions from empty document
- GIVEN a WordprocessingDocument with no tracked changes
- WHEN GetRevisions() is called
- THEN an empty iterator is returned
- AND iterating completes immediately without errors

#### Scenario: Revisions span all document parts
- GIVEN a document with changes in body, headers, footers, and footnotes
- WHEN GetRevisions() is called
- THEN all changes from all parts are included in the iterator
- AND changes from headers are included
- AND changes from footers are included
- AND changes from footnotes are included

### Requirement: Accept All Revisions
The system SHALL support accepting all tracked changes in a document.

#### Scenario: Accept all revisions
- GIVEN a document with 5 tracked changes
- WHEN AcceptAllRevisions() is called
- THEN all InsertedRun elements are replaced with their Run children
- AND all DeletedRun elements are removed
- AND all MoveFrom/MoveTo pairs are finalized
- AND GetRevisions() returns an empty iterator
- AND no error is returned

#### Scenario: Accept all on clean document
- GIVEN a document with no tracked changes
- WHEN AcceptAllRevisions() is called
- THEN no changes are made to the document
- AND no error is returned

#### Scenario: Accept all with move pairs
- GIVEN a document with MoveFromRun id=1 and MoveToRun id=1
- WHEN AcceptAllRevisions() is called
- THEN MoveFromRun is removed (content moved away)
- AND MoveToRun is replaced with its Run children (content stays in new location)
- AND move is finalized

### Requirement: Reject All Revisions
The system SHALL support rejecting all tracked changes in a document.

#### Scenario: Reject all revisions
- GIVEN a document with 5 tracked changes
- WHEN RejectAllRevisions() is called
- THEN all InsertedRun elements are removed
- AND all DeletedRun elements are converted back to normal text
- AND all MoveFrom/MoveTo pairs are reverted
- AND GetRevisions() returns an empty iterator
- AND no error is returned

#### Scenario: Reject all on clean document
- GIVEN a document with no tracked changes
- WHEN RejectAllRevisions() is called
- THEN no changes are made to the document
- AND no error is returned

#### Scenario: Reject all restores deleted content
- GIVEN a document with DeletedRun containing "deleted text"
- WHEN RejectAllRevisions() is called
- THEN "deleted text" appears as normal text in the document
- AND DeletedRun wrapper is removed

### Requirement: Filter Revisions by Author
The system SHALL support accepting or rejecting revisions from specific authors.

#### Scenario: Accept revisions by author
- GIVEN a document with changes from "John Doe" and "Jane Smith"
- WHEN AcceptRevisionsByAuthor("John Doe") is called
- THEN all changes authored by "John Doe" are accepted
- AND changes authored by "Jane Smith" remain as tracked changes
- AND no error is returned

#### Scenario: Reject revisions by author
- GIVEN a document with changes from "John Doe" and "Jane Smith"
- WHEN RejectRevisionsByAuthor("Jane Smith") is called
- THEN all changes authored by "Jane Smith" are rejected
- AND changes authored by "John Doe" remain as tracked changes
- AND no error is returned

#### Scenario: Filter by non-existent author
- GIVEN a document with changes from various authors
- WHEN AcceptRevisionsByAuthor("NonExistent") is called
- THEN no changes are made to the document
- AND no error is returned

### Requirement: Revision Metadata Access
The system SHALL provide access to revision metadata through the Revision wrapper.

#### Scenario: Access revision author
- GIVEN a Revision wrapping InsertedRun with author="John Doe"
- WHEN revision.Author() is called
- THEN "John Doe" is returned

#### Scenario: Access revision date
- GIVEN a Revision wrapping InsertedRun with date="2024-01-15T10:30:00Z"
- WHEN revision.Date() is called
- THEN time.Time representing 2024-01-15 10:30:00 UTC is returned

#### Scenario: Access revision ID
- GIVEN a Revision wrapping DeletedRun with id=42
- WHEN revision.Id() is called
- THEN 42 is returned

#### Scenario: Access revision type
- GIVEN a Revision wrapping InsertedRun
- WHEN revision.Type() is called
- THEN RevisionTypeInsert is returned

### Requirement: Individual Revision Operations
The system SHALL support accepting or rejecting individual revisions.

#### Scenario: Accept individual insertion
- GIVEN a Revision wrapping InsertedRun containing "new text"
- WHEN revision.Accept() is called
- THEN the Run elements inside InsertedRun are promoted to parent
- AND the InsertedRun wrapper is removed
- AND "new text" remains in the document
- AND no error is returned

#### Scenario: Reject individual insertion
- GIVEN a Revision wrapping InsertedRun containing "new text"
- WHEN revision.Reject() is called
- THEN the InsertedRun and all its content are removed
- AND "new text" does not appear in the document
- AND no error is returned

#### Scenario: Accept individual deletion
- GIVEN a Revision wrapping DeletedRun containing "removed text"
- WHEN revision.Accept() is called
- THEN the DeletedRun and its content are removed
- AND "removed text" does not appear in the document
- AND no error is returned

#### Scenario: Reject individual deletion
- GIVEN a Revision wrapping DeletedRun containing DeletedText "removed text"
- WHEN revision.Reject() is called
- THEN "removed text" is restored as normal text
- AND the DeletedRun wrapper is removed
- AND "removed text" appears in the document
- AND no error is returned

### Requirement: Move Revision Handling
The system SHALL handle MoveFrom and MoveTo revision pairs correctly.

#### Scenario: Accept move pair
- GIVEN a MoveFromRun with id=10 at position A
- AND a MoveToRun with id=10 at position B
- WHEN both revisions are accepted
- THEN content is removed from position A
- AND content appears at position B as normal text
- AND both MoveFromRun and MoveToRun wrappers are removed

#### Scenario: Reject move pair
- GIVEN a MoveFromRun with id=10 at position A
- AND a MoveToRun with id=10 at position B
- WHEN both revisions are rejected
- THEN content is restored at position A as normal text
- AND content is removed from position B
- AND both MoveFromRun and MoveToRun wrappers are removed

#### Scenario: Accept orphaned move
- GIVEN a MoveFromRun with id=10 without matching MoveToRun
- WHEN revision.Accept() is called
- THEN an error is returned indicating orphaned move
- AND the document is not modified

### Requirement: Format Change Revision Handling
The system SHALL handle RunPropertiesChange and ParagraphPropertiesChange revisions.

#### Scenario: Accept format change
- GIVEN a RunPropertiesChange showing bold formatting added
- WHEN revision.Accept() is called
- THEN the new formatting (bold) is applied
- AND the revision wrapper is removed
- AND no error is returned

#### Scenario: Reject format change
- GIVEN a RunPropertiesChange showing bold formatting added
- WHEN revision.Reject() is called
- THEN the original formatting (non-bold) is restored
- AND the revision wrapper is removed
- AND no error is returned
