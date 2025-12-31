# Implementation Tasks: Track Changes API

This implementation adds high-level API for revision management to the existing revision element infrastructure (~946 LOC in `wordprocessing/elements/revision.go`).

## Phase 1: Core Revision Wrapper (1 week, 8 tasks) ✅ COMPLETED

### 1.1 Define Core Types
- [x] 1.1.1 Create `wordprocessing/revisions.go`
  - **Verification**: File exists and compiles ✅
- [x] 1.1.2 Define `RevisionType` enum (Insert, Delete, MoveFrom, MoveTo, FormatChange)
  - **Verification**: All 5 enum values defined with godoc comments ✅
- [x] 1.1.3 Define `Revision` struct with element, typ, parent, document fields
  - **Verification**: Struct compiles with all fields properly typed ✅
- [x] 1.1.4 Add godoc comments to all types
  - **Verification**: `go doc wordprocessing.Revision` shows complete documentation ✅

### 1.2 Metadata Accessors
- [x] 1.2.1 Implement `Revision.Author() string` extracting from underlying element
  - **Verification**: Returns correct author for InsertedRun, DeletedRun, MoveFromRun, MoveToRun ✅
- [x] 1.2.2 Implement `Revision.Date() time.Time` parsing RFC3339 timestamps
  - **Verification**: Returns parsed time.Time for all revision types ✅
- [x] 1.2.3 Implement `Revision.Id() int` extracting revision ID
  - **Verification**: Returns correct int ID for all revision types ✅
- [x] 1.2.4 Implement `Revision.Type() RevisionType` returning enum value
  - **Verification**: Returns correct type for each underlying element type ✅
- [x] 1.2.5 Implement `Revision.Content() string` extracting text content
  - **Verification**: Returns InnerText for InsertedRun, DeletedText content for DeletedRun ✅
- [x] 1.2.6 Write unit tests for metadata accessors (5 tests minimum)
  - **Verification**: Tests pass, cover all revision types ✅

## Phase 2: Accept/Reject Logic (1.5 weeks, 12 tasks)

### 2.1 InsertedRun Operations
- [ ] 2.1.1 Implement `acceptInsertedRun()` helper
  - Gets parent element
  - Extracts Run children from InsertedRun
  - Replaces InsertedRun with Run children in parent
  - **Verification**: Unit test shows InsertedRun replaced, content promoted
- [ ] 2.1.2 Implement `rejectInsertedRun()` helper
  - Gets parent element
  - Removes InsertedRun entirely from parent
  - **Verification**: Unit test shows InsertedRun and content removed
- [ ] 2.1.3 Write 3 unit tests for InsertedRun accept/reject
  - Accept success case
  - Reject success case
  - Error when no parent element
  - **Verification**: All tests pass

### 2.2 DeletedRun Operations
- [ ] 2.2.1 Implement `acceptDeletedRun()` helper
  - Gets parent element
  - Removes DeletedRun entirely (content stays deleted)
  - **Verification**: Unit test shows DeletedRun removed
- [ ] 2.2.2 Implement `rejectDeletedRun()` helper
  - Gets parent element
  - Converts DeletedText to regular Text
  - Extracts Runs from DeletedRun
  - Replaces DeletedRun with Runs in parent
  - **Verification**: Unit test shows content restored as normal text
- [ ] 2.2.3 Write 3 unit tests for DeletedRun accept/reject
  - Accept success case
  - Reject success case (verify content restored)
  - Error when no parent element
  - **Verification**: All tests pass

### 2.3 Move Operations
- [ ] 2.3.1 Implement `findMovePair()` helper
  - Searches document for MoveFrom/MoveTo with matching ID
  - Returns paired Revision or nil
  - **Verification**: Unit test finds matching pair, returns nil for orphan
- [ ] 2.3.2 Implement `acceptMovePair()` helper
  - Removes MoveFromRun from origin
  - Replaces MoveToRun with content at destination
  - **Verification**: Unit test shows content moved correctly
- [ ] 2.3.3 Implement `rejectMovePair()` helper
  - Restores MoveFromRun content at origin
  - Removes MoveToRun from destination
  - **Verification**: Unit test shows move reverted correctly
- [ ] 2.3.4 Write 4 unit tests for move operations
  - Accept move pair
  - Reject move pair
  - Error on orphaned MoveFrom
  - Error on orphaned MoveTo
  - **Verification**: All tests pass

### 2.4 Format Change Operations
- [ ] 2.4.1 Implement `acceptFormatChange()` helper
  - Removes RunPropertiesChange/ParagraphPropertiesChange wrapper
  - Keeps current properties
  - **Verification**: Unit test shows wrapper removed, new formatting kept
- [ ] 2.4.2 Implement `rejectFormatChange()` helper
  - Extracts previous properties from revision element
  - Applies previous properties to parent
  - Removes wrapper
  - **Verification**: Unit test shows original formatting restored
- [ ] 2.4.3 Write 2 unit tests for format change operations
  - Accept format change
  - Reject format change
  - **Verification**: All tests pass

### 2.5 Main Accept/Reject Methods
- [ ] 2.5.1 Implement `Revision.Accept() error`
  - Switches on RevisionType
  - Calls appropriate helper
  - Returns error for orphaned moves
  - **Verification**: Unit test for each revision type
- [ ] 2.5.2 Implement `Revision.Reject() error`
  - Switches on RevisionType
  - Calls appropriate helper
  - Returns error for orphaned moves
  - **Verification**: Unit test for each revision type

## Phase 3: Document Traversal (1 week, 10 tasks)

### 3.1 Tree Traversal Implementation
- [ ] 3.1.1 Implement `collectRevisionsFromElement()` helper
  - Recursively walks element tree
  - Identifies revision elements by type
  - Wraps in Revision struct
  - **Verification**: Unit test finds revisions in nested structures
- [ ] 3.1.2 Handle InsertedRun element detection
  - **Verification**: Test finds InsertedRun at various nesting levels
- [ ] 3.1.3 Handle DeletedRun element detection
  - **Verification**: Test finds DeletedRun at various nesting levels
- [ ] 3.1.4 Handle MoveFromRun element detection
  - **Verification**: Test finds MoveFromRun elements
- [ ] 3.1.5 Handle MoveToRun element detection
  - **Verification**: Test finds MoveToRun elements
- [ ] 3.1.6 Handle RunPropertiesChange element detection
  - **Verification**: Test finds RunPropertiesChange elements
- [ ] 3.1.7 Handle ParagraphPropertiesChange element detection
  - **Verification**: Test finds ParagraphPropertiesChange elements

### 3.2 Document Part Traversal
- [ ] 3.2.1 Implement traversal for main document Body
  - **Verification**: Test finds all revisions in Body paragraphs
- [ ] 3.2.2 Implement traversal for all Headers
  - Iterate HeaderPart for first, even, odd, default
  - **Verification**: Test finds revisions in header content
- [ ] 3.2.3 Implement traversal for all Footers
  - Iterate FooterPart for first, even, odd, default
  - **Verification**: Test finds revisions in footer content
- [ ] 3.2.4 Implement traversal for Comments
  - **Verification**: Test finds revisions in comment text
- [ ] 3.2.5 Implement traversal for Footnotes
  - **Verification**: Test finds revisions in footnote content
- [ ] 3.2.6 Implement traversal for Endnotes
  - **Verification**: Test finds revisions in endnote content
- [ ] 3.2.7 Write 6 unit tests for part traversal
  - Body traversal
  - Header traversal
  - Footer traversal
  - Comments traversal
  - Footnotes traversal
  - Endnotes traversal
  - **Verification**: All tests pass

## Phase 4: Document-Level API (0.5 weeks, 6 tasks)

### 4.1 Iterator Implementation
- [ ] 4.1.1 Implement `WordprocessingDocument.GetRevisions() iter.Seq[*Revision]`
  - Uses Go 1.23+ iter.Seq pattern
  - Calls traversal logic
  - Yields Revision wrappers
  - **Verification**: Test can iterate all revisions in document
- [ ] 4.1.2 Write 2 unit tests for GetRevisions iterator
  - Iterate all revisions
  - Early break stops enumeration
  - **Verification**: Tests pass

### 4.2 Bulk Operations
- [ ] 4.2.1 Implement `AcceptAllRevisions() error`
  - Iterates GetRevisions()
  - Calls Accept() on each
  - Returns first error or nil
  - **Verification**: Test shows all revisions accepted
- [ ] 4.2.2 Implement `RejectAllRevisions() error`
  - Iterates GetRevisions()
  - Calls Reject() on each
  - Returns first error or nil
  - **Verification**: Test shows all revisions rejected
- [ ] 4.2.3 Implement `AcceptRevisionsByAuthor(author string) error`
  - Iterates GetRevisions()
  - Filters by author
  - Calls Accept() on matches
  - **Verification**: Test shows only author's revisions accepted
- [ ] 4.2.4 Implement `RejectRevisionsByAuthor(author string) error`
  - Iterates GetRevisions()
  - Filters by author
  - Calls Reject() on matches
  - **Verification**: Test shows only author's revisions rejected
- [ ] 4.2.5 Write 4 unit tests for bulk operations
  - AcceptAllRevisions
  - RejectAllRevisions
  - AcceptRevisionsByAuthor
  - RejectRevisionsByAuthor
  - **Verification**: All tests pass

## Phase 5: Integration & Testing (1 week, 8 tasks)

### 5.1 Roundtrip Tests
- [ ] 5.1.1 Create test document with InsertedRun and DeletedRun
  - **Verification**: Document file created in testdata/
- [ ] 5.1.2 Write roundtrip test: open → enumerate → verify
  - Opens document with track changes
  - Enumerates all revisions
  - Verifies count, authors, dates, types
  - **Verification**: Test passes
- [ ] 5.1.3 Write roundtrip test: accept → save → open → verify
  - Opens document
  - Accepts specific revisions
  - Saves document
  - Reopens and verifies changes accepted
  - **Verification**: Test passes
- [ ] 5.1.4 Write roundtrip test: reject → save → open → verify
  - Opens document
  - Rejects specific revisions
  - Saves document
  - Reopens and verifies changes rejected
  - **Verification**: Test passes

### 5.2 Edge Case Tests
- [ ] 5.2.1 Write test for nested revisions (InsertedRun inside DeletedRun)
  - **Verification**: Test handles nested cases without error
- [ ] 5.2.2 Write test for orphaned move elements
  - **Verification**: Test returns error for orphaned moves
- [ ] 5.2.3 Write test for revisions in table cells
  - **Verification**: Test finds revisions in tables
- [ ] 5.2.4 Write test for revisions in headers/footers
  - **Verification**: Test finds revisions in all header/footer types

### 5.3 Performance & Compatibility Tests
- [ ] 5.3.1 Write performance test: 10,000 revisions enumerated < 1 second
  - **Verification**: Benchmark shows acceptable performance
- [ ] 5.3.2 Create test documents from Word 2007, 2010, 2016, 365
  - **Verification**: All Word versions supported
- [ ] 5.3.3 Validate output opens correctly in Microsoft Word
  - **Verification**: Manual verification or automated check
- [ ] 5.3.4 Validate output opens correctly in LibreOffice
  - **Verification**: Manual verification or automated check

## Phase 6: Documentation (0.5 weeks, 6 tasks)

### 6.1 API Documentation
- [ ] 6.1.1 Add godoc comments to all exported types
  - **Verification**: `go doc wordprocessing` shows complete docs
- [ ] 6.1.2 Add godoc comments to all exported methods
  - **Verification**: Each method has meaningful documentation
- [ ] 6.1.3 Add usage examples in package-level godoc
  - **Verification**: Examples compile and run correctly

### 6.2 Code Examples
- [ ] 6.2.1 Create `examples/track-changes/accept-all/main.go`
  - Opens document
  - Accepts all revisions
  - Saves result
  - **Verification**: Example compiles and runs
- [ ] 6.2.2 Create `examples/track-changes/accept-by-author/main.go`
  - Opens document
  - Accepts revisions from specific author
  - Saves result
  - **Verification**: Example compiles and runs
- [ ] 6.2.3 Create `examples/track-changes/enumerate/main.go`
  - Opens document
  - Enumerates all revisions
  - Prints metadata (author, date, type, content)
  - **Verification**: Example compiles and runs

## Validation Checkpoints

**After Phase 1**: ✅ COMPLETED
- [x] All metadata accessors work correctly
- [x] Revision wrapper compiles and has tests
- [x] All 5 revision types can be detected

**After Phase 2**:
- [ ] Accept/Reject work for all revision types
- [ ] Move pairs coordinated correctly
- [ ] Orphaned moves return errors
- [ ] Unit tests have >90% coverage

**After Phase 3**:
- [ ] GetRevisions() finds all revisions in document
- [ ] All document parts searched (body, headers, footers, etc.)
- [ ] Iterator pattern works correctly

**After Phase 4**:
- [ ] Document-level API complete
- [ ] AcceptAllRevisions() works
- [ ] RejectAllRevisions() works
- [ ] Filter by author works

**After Phase 5**:
- [ ] Roundtrip tests pass
- [ ] Edge cases handled
- [ ] Performance acceptable
- [ ] Compatibility verified

**After Phase 6**:
- [ ] Documentation complete
- [ ] Examples work
- [ ] Ready for review

## Task Count Summary

- Phase 1: 8 tasks
- Phase 2: 12 tasks
- Phase 3: 10 tasks
- Phase 4: 6 tasks
- Phase 5: 8 tasks
- Phase 6: 6 tasks
- **Total: 50 tasks** (exceeds requested 30-40 for granularity)

## Dependencies

- Phase 2 depends on Phase 1 (needs Revision struct)
- Phase 3 independent (can run parallel with Phase 2)
- Phase 4 depends on Phase 1 and Phase 3 (needs Revision and traversal)
- Phase 5 depends on Phase 1-4 (integration testing)
- Phase 6 can start after Phase 4 (examples need API complete)

## Parallelization Opportunities

- Phase 2 and Phase 3 can run in parallel (different files)
- Within Phase 2, different revision type implementations can run in parallel
- Within Phase 3, different part traversal implementations can run in parallel
- Phase 6 tasks are independent and can run in parallel
