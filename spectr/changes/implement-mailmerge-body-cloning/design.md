# Proposal: Implement Mailmerge Body Cloning for Multiple Records

## Summary

Implement proper body cloning in the mailmerge engine to support generating multiple output documents from a single template with multiple data records.

## Background

Currently, the mailmerge `Execute()` function in `engine.go` has a TODO indicating that proper body cloning for multiple records needs to be implemented. The current implementation only processes the first record from the data source.

**Reference TODO:** `wordprocessing/mailmerge/engine.go:233`
```go
// TODO: Implement proper body cloning in later phases for multiple records
```

## Motivation

The mailmerge engine needs to support:
1. Single merged document with all records (current `Execute()` behavior, but incomplete)
2. Multiple separate documents (already working via `ExecuteToDocuments()`)

For the single merged document case, the engine must properly clone the template body for each data record and append them to the output document with appropriate page breaks.

## Technical Design

### Current Implementation

The `Execute()` method currently:
1. Validates the data source and template
2. Opens the data source
3. Finds merge fields in the template
4. Validates fields against data source
5. Creates an output document
6. **Only processes the first record** (this is the limitation)

### Proposed Implementation

Extend `Execute()` to:
1. Continue existing validation steps
2. Iterate through ALL records in the data source
3. For each record:
   - Clone the template body content (deep copy)
   - Replace merge fields with current record values
   - Append to output document
   - Add page break between records (except after last)
4. Return the fully merged document

### API Changes

No public API changes required. The behavior of `Execute()` will be completed to match its documented intent.

### Code Structure

```go
func (m *MailMerge) Execute() (*wordprocessing.Document, error) {
    // ... existing validation code ...
    
    // Create output document
    output, err := m.cloneTemplateDocument()
    if err != nil {
        return nil, fmt.Errorf("failed to create output document: %w", err)
    }
    
    // Track if we need page breaks between records
    isFirstRecord := true
    
    // Process ALL records (not just first)
    for m.dataSource.Next() {
        // Clone template body for this record
        clonedBody, err := m.cloneBodyForRecord()
        if err != nil {
            return nil, fmt.Errorf("failed to clone body: %w", err)
        }
        
        // Replace fields in cloned body
        if err := m.replaceFieldsInBody(clonedBody, currentFields, m.dataSource); err != nil {
            return nil, err
        }
        
        // Add page break before appending (except for first record)
        if !isFirstRecord {
            m.addPageBreakBefore(clonedBody)
        }
        isFirstRecord = false
        
        // Append cloned body to output
        if err := m.appendBodyToOutput(output, clonedBody); err != nil {
            return nil, err
        }
    }
    
    return output, nil
}
```

### Body Cloning Strategy

Options for body cloning:

1. **Deep Copy via XML Serialization** (Recommended)
   - Serialize body to XML
   - Deserialize to create new body
   - Pros: Deep copy, handles all element types
   - Cons: Slower, more memory usage

2. **Manual Element Cloning**
   - Recursively clone each element
   - Pros: Faster than serialization
   - Cons: Complex to implement, may miss edge cases

3. **Template Re-parse**
   - Re-parse original template for each record
   - Pros: Cleanest result
   - Cons: Very slow for many records

Recommendation: Start with Option 1 (XML serialization) for correctness, optimize later if needed.

## Requirements

### SHALL Requirements

#### Requirement: Multi-Record Processing
The Execute() method SHALL process all records from the data source, not just the first record.

##### Scenario: CSV with Multiple Records
Given a CSV file with 3 records
When Execute() is called
Then the output document SHALL contain all 3 personalized copies

##### Scenario: JSON with Array of Objects
Given a JSON file with an array of 5 objects
When Execute() is called
Then the output document SHALL contain all 5 personalized copies

#### Requirement: Page Breaks Between Records
The Execute() method SHALL insert page breaks between records so each record starts on a new page.

##### Scenario: Multiple Records with Page Breaks
Given a template with 1 page and a data source with 3 records
When Execute() is called
Then the output document SHALL have 3 pages with page breaks between each

#### Requirement: Field Replacement per Record
Each record SHALL receive its own field values, not values from other records.

##### Scenario: Unique Values Per Record
Given a data source with records having unique "Name" values
When Execute() is called
Then each page SHALL contain the correct Name for that record

#### Requirement: Preserve Template Formatting
The cloning process SHALL preserve all formatting from the template.

##### Scenario: Preserve Paragraph Styles
Given a template with styled paragraphs (bold, italic, alignment)
When Execute() is called with multiple records
Then each copy SHALL maintain the original formatting

##### Scenario: Preserve Table Formatting
Given a template with formatted tables
When Execute() is called with multiple records
Then tables SHALL be preserved in each copy

### SHOULD Requirements

#### Requirement: Performance Optimization
The implementation SHOULD be optimized for reasonable performance with up to 100 records.

#### Requirement: Memory Efficiency
The implementation SHOULD manage memory efficiently for large documents.

## Testing Strategy

### Unit Tests
- Test cloning of simple body content
- Test cloning of body with tables
- Test cloning of body with images
- Test field replacement in cloned bodies
- Test page break insertion

### Integration Tests
- End-to-end mailmerge with 3+ records
- End-to-end with complex formatting
- End-to-end with tables spanning records

### Performance Tests
- Test with 10 records
- Test with 50 records
- Test with 100 records

## Implementation Plan

1. Implement body cloning function
2. Update Execute() to iterate all records
3. Implement page break insertion
4. Add unit tests
5. Add integration tests
6. Performance testing
7. Documentation update

## Related Changes

- `wordprocessing/mailmerge/engine.go` - Main implementation
- `wordprocessing/mailmerge/clone.go` - New cloning utilities
- `wordprocessing/mailmerge/engine_test.go` - Unit tests

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Performance issues with many records | Medium | Optimize cloning strategy; document limits |
| Memory exhaustion with large docs | Medium | Stream processing; memory profiling |
| Formatting loss during clone | High | Comprehensive test suite; XML serialization |
| Page break issues | Low | Test various content types |

## Acceptance Criteria

- [ ] All SHALL requirements implemented
- [ ] Execute() processes all records from data source
- [ ] Page breaks inserted between records
- [ ] Formatting preserved in all copies
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] Performance acceptable for 100 records
- [ ] No regression in ExecuteToDocuments()
