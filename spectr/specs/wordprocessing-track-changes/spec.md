# Wordprocessing Track Changes Specification

## Requirements

### Requirement: TrackChange Interface

The system SHALL provide a unified interface over all track change element types.

#### Scenario: InsertedRun implements TrackChange
- GIVEN an InsertedRun element with id="1", author="John", date="2024-01-15T10:30:00Z"
- WHEN wrapped as TrackChange
- THEN `change.Id()` returns "1"
- AND `change.Author()` returns "John"
- AND `change.Date()` returns parsed time.Time
- AND `change.Type()` returns TrackChangeInsertion
- AND `change.Content()` returns InnerText of contained runs

#### Scenario: DeletedRun implements TrackChange
- GIVEN a DeletedRun element with deleted text content
- WHEN wrapped as TrackChange
- THEN `change.Type()` returns TrackChangeDeletion
- AND `change.Content()` returns concatenated DeletedText values

#### Scenario: MoveFromRun implements TrackChange
- GIVEN a MoveFromRun element
- WHEN wrapped as TrackChange
- THEN `change.Type()` returns TrackChangeMoveFrom

#### Scenario: MoveToRun implements TrackChange
- GIVEN a MoveToRun element
- WHEN wrapped as TrackChange
- THEN `change.Type()` returns TrackChangeMoveTo

#### Scenario: Access underlying element
- GIVEN a TrackChange wrapping an InsertedRun
- WHEN `change.Element()` is called
- THEN the original `*InsertedRun` is returned
- AND can be cast back to specific type

### Requirement: Accept Operation

The system SHALL apply track changes, removing tracking wrappers.

#### Scenario: Accept InsertedRun
- GIVEN a paragraph with InsertedRun containing "new text"
- WHEN `change.Accept()` is called
- THEN the Run elements inside InsertedRun are promoted to paragraph
- AND the InsertedRun wrapper is removed
- AND the text "new text" remains in the paragraph
- AND no error is returned

#### Scenario: Accept DeletedRun
- GIVEN a paragraph with DeletedRun containing deleted text
- WHEN `change.Accept()` is called
- THEN the DeletedRun and its content are removed
- AND the deleted text does NOT appear in paragraph
- AND no error is returned

#### Scenario: Accept orphan element returns error
- GIVEN a TrackChange with no parent element
- WHEN `change.Accept()` is called
- THEN error is returned indicating orphan element

### Requirement: Reject Operation

The system SHALL revert track changes, restoring original state.

#### Scenario: Reject InsertedRun
- GIVEN a paragraph with InsertedRun containing "new text"
- WHEN `change.Reject()` is called
- THEN the InsertedRun and all its content are removed
- AND the text "new text" does NOT appear in paragraph
- AND no error is returned

#### Scenario: Reject DeletedRun
- GIVEN a paragraph with DeletedRun containing DeletedText "old text"
- WHEN `change.Reject()` is called
- THEN the DeletedText is converted to regular Run with Text
- AND the DeletedRun wrapper is removed
- AND the text "old text" appears in paragraph as normal text
- AND no error is returned

### Requirement: TrackChangeCollection

The system SHALL provide collection-based access to all track changes.

#### Scenario: Enumerate all changes
- GIVEN a document with 3 InsertedRuns and 2 DeletedRuns
- WHEN `doc.TrackChanges()` is called
- THEN a TrackChangeCollection is returned
- AND `collection.Count()` returns 5
- AND `collection.All()` returns slice of 5 TrackChange items

#### Scenario: Filter by author
- GIVEN a document with changes from authors "John" and "Jane"
- WHEN `doc.TrackChanges().ByAuthor("John")` is called
- THEN a new TrackChangeCollection is returned
- AND collection contains only changes by John

#### Scenario: Filter by date range
- GIVEN a document with changes from January and February
- WHEN `doc.TrackChanges().ByDateRange(jan1, jan31)` is called
- THEN a new TrackChangeCollection is returned
- AND collection contains only January changes

#### Scenario: Filter by type - insertions
- GIVEN a document with insertions and deletions
- WHEN `doc.TrackChanges().Insertions()` is called
- THEN collection contains only InsertedRun changes

#### Scenario: Filter by type - deletions
- GIVEN a document with insertions and deletions
- WHEN `doc.TrackChanges().Deletions()` is called
- THEN collection contains only DeletedRun changes

#### Scenario: Chain filters
- GIVEN a document with various changes
- WHEN `doc.TrackChanges().ByAuthor("John").Insertions()` is called
- THEN collection contains only insertions by John

#### Scenario: Get change by ID
- GIVEN a document with change id="42"
- WHEN `doc.TrackChanges().ById("42")` is called
- THEN the specific TrackChange is returned
- AND nil is returned if ID not found

### Requirement: Bulk Operations

The system SHALL support accepting/rejecting all changes in a collection.

#### Scenario: Accept all changes
- GIVEN a document with 5 track changes
- WHEN `doc.TrackChanges().AcceptAll()` is called
- THEN all 5 changes are accepted
- AND document no longer has track change wrappers
- AND no error is returned

#### Scenario: Reject all changes
- GIVEN a document with 5 track changes
- WHEN `doc.TrackChanges().RejectAll()` is called
- THEN all 5 changes are rejected
- AND document no longer has track change wrappers
- AND document content matches pre-change state

#### Scenario: Bulk operation with filtered collection
- GIVEN a document with changes from John and Jane
- WHEN `doc.TrackChanges().ByAuthor("John").AcceptAll()` is called
- THEN only John's changes are accepted
- AND Jane's changes remain as track changes

### Requirement: Paragraph Integration

The system SHALL provide convenience methods for creating tracked changes.

#### Scenario: Append inserted run
- GIVEN a paragraph
- WHEN `para.AppendInsertedRun("John", "new text")` is called
- THEN an InsertedRun is created with author="John"
- AND InsertedRun contains Run with "new text"
- AND InsertedRun has auto-generated id
- AND InsertedRun has current timestamp as date
- AND InsertedRun is returned for further customization

#### Scenario: Append deleted run
- GIVEN a paragraph
- WHEN `para.AppendDeletedRun("John", "removed text")` is called
- THEN a DeletedRun is created with author="John"
- AND DeletedRun contains DeletedText with "removed text"
- AND DeletedRun has auto-generated id and current timestamp
- AND DeletedRun is returned

#### Scenario: Created changes appear in enumeration
- GIVEN a paragraph where `AppendInsertedRun` was called
- WHEN `doc.TrackChanges()` is enumerated
- THEN the created InsertedRun appears in the collection

### Requirement: Body-Level Enumeration

The system SHALL support enumerating changes at body level.

#### Scenario: Body track changes
- GIVEN a document body with track changes
- WHEN `body.TrackChanges()` is called
- THEN a TrackChangeCollection is returned
- AND contains all changes within body
- AND excludes changes in headers/footers

### Requirement: Document-Level Enumeration

The system SHALL support enumerating all changes in document.

#### Scenario: Document track changes
- GIVEN a document with changes in body, headers, and footers
- WHEN `doc.TrackChanges()` is called
- THEN a TrackChangeCollection is returned
- AND contains all changes across all document parts

#### Scenario: Empty document
- GIVEN a document with no track changes
- WHEN `doc.TrackChanges()` is called
- THEN an empty TrackChangeCollection is returned
- AND `collection.Count()` returns 0

