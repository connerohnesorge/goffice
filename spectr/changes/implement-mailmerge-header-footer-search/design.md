# Proposal: Implement Mailmerge Header and Footer Field Search

## Summary

Extend the mailmerge field finder to search for merge fields in document headers and footers, not just the main body content.

## Background

Currently, the mailmerge `field_finder.go` only searches for merge fields within the main document body. This limitation means that merge fields placed in headers or footers are not detected or processed during mail merge operations.

**Reference TODO:** `wordprocessing/mailmerge/field_finder.go:49`
```go
// TODO: Search in headers and footers in later phases
```

## Motivation

Headers and footers are common locations for merge fields in professional documents:
- Letterhead information (company name, address)
- Document control information (document ID, revision date)
- Personalized headers per recipient
- Page numbering with merge context

Without this feature, users cannot create fully personalized documents that include header/footer content specific to each data record.

## Technical Design

### Current Implementation

The `findMergeFields()` function in `field_finder.go` currently:
1. Gets the document body from the main document part
2. Iterates through all body elements
3. Searches for `SimpleField` and `ComplexField` elements within the body

### Proposed Implementation

Extend `findMergeFields()` to:
1. Continue searching the main document body (existing behavior)
2. Iterate through all headers associated with the document
3. Iterate through all footers associated with the document
4. Search for merge fields within each header and footer

### API Changes

No public API changes required. The `findMergeFields()` method is internal.

### Code Structure

```go
func (m *MailMerge) findMergeFields() []*MergeField {
    var fields []*MergeField
    
    // Search in main document body (existing)
    mainPart := m.template.MainPart()
    if mainPart == nil {
        return fields
    }
    
    doc := mainPart.Document()
    if doc == nil {
        return fields
    }
    
    body := doc.Body()
    if body != nil {
        fields = append(fields, findFieldsInElement(body)...)
    }
    
    // NEW: Search in headers
    for header := range m.template.Headers() {
        fields = append(fields, findFieldsInElement(header)...)
    }
    
    // NEW: Search in footers
    for footer := range m.template.Footers() {
        fields = append(fields, findFieldsInElement(footer)...)
    }
    
    return fields
}
```

## Requirements

### SHALL Requirements

#### Requirement: Header Field Detection
The field finder SHALL detect merge fields within all header parts associated with the document.

##### Scenario: Detect MERGEFIELD in Default Header
Given a document with a default header containing `MERGEFIELD CompanyName`
When the mailmerge operation executes
Then the CompanyName field SHALL be detected and replaced with data source values

##### Scenario: Detect MERGEFIELD in First Page Header
Given a document with a first page header containing `MERGEFIELD DocumentTitle`
When the mailmerge operation executes
Then the DocumentTitle field SHALL be detected and replaced

##### Scenario: Detect Complex Fields in Headers
Given a document with complex field markers (fldChar) in a header
When the mailmerge operation executes
Then all complex merge fields SHALL be properly detected and replaced

#### Requirement: Footer Field Detection
The field finder SHALL detect merge fields within all footer parts associated with the document.

##### Scenario: Detect MERGEFIELD in Default Footer
Given a document with a default footer containing `MERGEFIELD PageLabel`
When the mailmerge operation executes
Then the PageLabel field SHALL be detected and replaced

##### Scenario: Detect Complex Fields in Footers
Given a document with complex field markers (fldChar) in a footer
When the mailmerge operation executes
Then all complex merge fields SHALL be properly detected and replaced

#### Requirement: Header/Footer with Same Field Names
The field finder SHALL correctly handle fields with the same name appearing in both body and header/footer.

##### Scenario: Same Field Name in Multiple Locations
Given a document with `MERGEFIELD CustomerName` in body and header
When the mailmerge operation executes
Then both instances SHALL be replaced with the same data source value

#### Requirement: Empty Header/Footer Handling
The implementation SHALL gracefully handle documents with no headers or footers.

##### Scenario: Document Without Headers
Given a document with no header parts
When the field finder executes
Then it SHALL complete without error and only search the body

##### Scenario: Document Without Footers
Given a document with no footer parts
When the field finder executes
Then it SHALL complete without error and only search the body

### SHOULD Requirements

#### Requirement: Performance Optimization
The implementation SHOULD minimize performance impact when documents have no header/footer merge fields.

#### Requirement: Support All Header/Footer Types
The implementation SHOULD support all Word header/footer types:
- Default (odd pages)
- Even pages
- First page

## Testing Strategy

### Unit Tests
- Test field detection in headers with simple fields
- Test field detection in headers with complex fields
- Test field detection in footers with simple fields
- Test field detection in footers with complex fields
- Test documents without headers/footers
- Test same field name in multiple locations

### Integration Tests
- End-to-end mailmerge with header fields
- End-to-end mailmerge with footer fields
- End-to-end mailmerge with fields in all locations

### Test Documents
Create test documents with:
1. Simple MERGEFIELD in header
2. Complex field (fldChar markers) in header
3. Simple MERGEFIELD in footer
4. Complex field in footer
5. GREETINGLINE in header
6. Multiple field types across body/header/footer

## Implementation Plan

1. Update `findMergeFields()` to iterate headers
2. Update `findMergeFields()` to iterate footers
3. Add unit tests
4. Add integration tests
5. Update documentation

## Related Changes

- `wordprocessing/mailmerge/field_finder.go` - Main implementation
- `wordprocessing/mailmerge/engine_test.go` - Unit tests
- Tests may require new test document fixtures

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Performance degradation with large documents | Medium | Early exit if no headers/footers; efficient iteration |
| Breaking existing body-only merge behavior | Low | Add feature flag if needed; maintain backward compatibility |
| Complex field parsing issues in headers | Low | Reuse existing field parsing logic |

## Acceptance Criteria

- [ ] All SHALL requirements implemented
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] No regression in existing mailmerge tests
- [ ] Documentation updated
