# Tasks: PivotTable API Completion

**Baseline**: ~560 LOC exists in `spreadsheet/pivottable*.go` with field management, aggregation enum, and part infrastructure. This proposal **enhances** existing code.

## 1. Phase 1: Core Completion (60-80 hours total)

### 1.1 Cache Population Implementation (15-20h)

- [ ] 1.1.1 Create `spreadsheet/pivottable_cache.go`
- [ ] 1.1.2 Implement `parseRangeReference(ref string) (sheet, cellRange, error)` helper
- [ ] 1.1.3 Implement `extractHeaderRow(sheet, cellRef) []string` to get field names
- [ ] 1.1.4 Implement `extractDataRows(sheet, cellRef) [][]interface{}` to get data
- [ ] 1.1.5 Implement `populateCacheFields(cacheDef, headers)` to set field definitions
- [ ] 1.1.6 Implement `populateCacheRecords(records, rows, headers)` to store data
- [ ] 1.1.7 Implement `PivotTable.Refresh() error` method
- [ ] 1.1.8 Add `getCacheDefinitionPart()` helper to PivotTable
- [ ] 1.1.9 Add `getCacheRecordsPart()` helper to PivotTable
- [ ] 1.1.10 Write 8 unit tests for cache population (see design.md)
- [ ] 1.1.11 Test with 10k row dataset for performance baseline

### 1.2 Field Validation Implementation (8-12h)

- [ ] 1.2.1 Create `spreadsheet/pivottable_errors.go` with error types
- [ ] 1.2.2 Implement `ErrFieldNotFound` struct with Error() method
- [ ] 1.2.3 Implement `ErrDuplicateField` struct with Error() method
- [ ] 1.2.4 Implement `ErrInvalidSourceRange` struct with Error() method
- [ ] 1.2.5 Add `validateFieldName(name string) error` to PivotTable
- [ ] 1.2.6 Add `containsField(name string) bool` to check all axes
- [ ] 1.2.7 Implement `TryAddRowField(name) (*PivotTable, error)`
- [ ] 1.2.8 Implement `TryAddColumnField(name) (*PivotTable, error)`
- [ ] 1.2.9 Implement `TryAddDataField(name, agg) (*PivotTable, error)`
- [ ] 1.2.10 Implement `TryAddPageField(name) (*PivotTable, error)`
- [ ] 1.2.11 Write 10 unit tests for validation (see design.md)

### 1.3 PivotField Wrapper Implementation (12-15h)

- [ ] 1.3.1 Create `spreadsheet/pivottable_field.go`
- [ ] 1.3.2 Define `PivotField` struct with pivot, element, index fields
- [ ] 1.3.3 Define `AxisType` enum (Row, Column, Page, Data, None)
- [ ] 1.3.4 Define `FieldDataType` enum (String, Number, Date, Boolean)
- [ ] 1.3.5 Implement `PivotField.Name() string`
- [ ] 1.3.6 Implement `PivotField.Index() uint32`
- [ ] 1.3.7 Implement `PivotField.Axis() AxisType`
- [ ] 1.3.8 Implement `PivotField.ShowAll() bool` and `SetShowAll(bool)`
- [ ] 1.3.9 Implement `PivotField.Compact() bool` and `SetCompact(bool)`
- [ ] 1.3.10 Implement `PivotField.HideNewItems() bool` and `SetHideNewItems(bool)`
- [ ] 1.3.11 Implement `PivotField.ShowDropDowns() bool` and `SetShowDropDowns(bool)`
- [ ] 1.3.12 Implement `PivotField.DefaultSubtotal() bool` and `SetDefaultSubtotal(bool)`
- [ ] 1.3.13 Implement `PivotField.DataType() FieldDataType`
- [ ] 1.3.14 Implement `PivotField.ItemCount() int`
- [ ] 1.3.15 Add `PivotTable.Field(name string) *PivotField`
- [ ] 1.3.16 Add `PivotTable.Fields() []*PivotField`
- [ ] 1.3.17 Write 12 unit tests for PivotField properties (see design.md)

### 1.4 XML Serialization Enhancement (8-12h)

- [ ] 1.4.1 Review existing `updateFields()` in `pivottable_update.go`
- [ ] 1.4.2 Ensure PivotFields collection serializes all field metadata
- [ ] 1.4.3 Ensure RowFields indices point to correct PivotField indices
- [ ] 1.4.4 Ensure ColFields indices serialize correctly
- [ ] 1.4.5 Ensure PageFields indices serialize correctly
- [ ] 1.4.6 Ensure DataFields serialize with aggregation function
- [ ] 1.4.7 Add relationship ID management for cache linking
- [ ] 1.4.8 Verify Location element serialization (ref, firstDataRow, etc.)
- [ ] 1.4.9 Test roundtrip: create → save → load → verify structure matches

### 1.5 Roundtrip & Integration Testing (8-10h)

- [ ] 1.5.1 Create test fixtures directory `testdata/pivottable/`
- [ ] 1.5.2 Create simple pivot test fixture (Product vs Year)
- [ ] 1.5.3 Create complex pivot fixture (multi-axis, multiple data fields)
- [ ] 1.5.4 Write integration test: create → save → open → verify
- [ ] 1.5.5 Write integration test: open Excel file → read properties → verify
- [ ] 1.5.6 Write integration test: modify existing → save → verify in Excel
- [ ] 1.5.7 Write performance test: 10k rows < 2s refresh
- [ ] 1.5.8 Create Excel compatibility test (Excel 2010, 2016, 2019, 365)

## 2. Phase 2: Advanced Features (60-80 hours total)

### 2.1 Filter API (15-20h)

- [ ] 2.1.1 Analyze PivotFilters element structure
- [ ] 2.1.2 Implement `PivotFilter` wrapper type
- [ ] 2.1.3 Add `PivotTable.Filters() []*PivotFilter`
- [ ] 2.1.4 Add `PivotTable.AddFilter(fieldName, filterType, values) error`
- [ ] 2.1.5 Add `PivotTable.RemoveFilter(fieldName) error`
- [ ] 2.1.6 Write 10 unit tests for filter operations

### 2.2 AutoSort API (8-12h)

- [ ] 2.2.1 Analyze AutoSortScope element structure
- [ ] 2.2.2 Implement `SortScope` wrapper type
- [ ] 2.2.3 Add `PivotField.Sort() *SortScope`
- [ ] 2.2.4 Add `PivotField.SetSort(order, dataField) error`
- [ ] 2.2.5 Write 6 unit tests for sort operations

### 2.3 Field Grouping (25-35h)

- [ ] 2.3.1 Analyze FieldGroup, RangeProperties, DateGroupItem elements
- [ ] 2.3.2 Define `DateGroupingType` enum (Years, Quarters, Months, Days, Hours)
- [ ] 2.3.3 Implement `FieldGroup` wrapper type
- [ ] 2.3.4 Add `PivotField.GroupByDate(intervals []DateGroupingType) error`
- [ ] 2.3.5 Add `PivotField.GroupByRange(start, end, step float64) error`
- [ ] 2.3.6 Add `PivotField.Ungroup() error`
- [ ] 2.3.7 Write 12 unit tests for grouping operations
- [ ] 2.3.8 Test roundtrip with grouped fields

## 3. Phase 3: Polish (30-40 hours total)

### 3.1 Examples (12-15h)

- [ ] 3.1.1 Create `examples/pivot_basic/main.go` - simple sales pivot
- [ ] 3.1.2 Create `examples/pivot_analysis/main.go` - multi-field analysis
- [ ] 3.1.3 Create `examples/pivot_modify/main.go` - open and modify existing
- [ ] 3.1.4 Add README.md for each example with expected output
- [ ] 3.1.5 Verify all examples compile and run correctly

### 3.2 Documentation (10-15h)

- [ ] 3.2.1 Add comprehensive godoc comments to all exported types
- [ ] 3.2.2 Add usage examples in doc comments
- [ ] 3.2.3 Write migration guide for OpenXML-SDK users
- [ ] 3.2.4 Document error handling patterns
- [ ] 3.2.5 Document performance considerations

### 3.3 Performance Optimization (8-10h)

- [ ] 3.3.1 Add benchmark tests for cache loading
- [ ] 3.3.2 Profile Refresh() on large datasets
- [ ] 3.3.3 Optimize field lookup with index maps if needed
- [ ] 3.3.4 Add streaming option for very large sources (>100k rows)
- [ ] 3.3.5 Document performance targets in README

## Summary

| Phase | Hours | Key Deliverables |
|-------|-------|------------------|
| Phase 1 | 60-80h | Cache population, validation, PivotField wrapper, roundtrip |
| Phase 2 | 60-80h | Filters, sorting, field grouping |
| Phase 3 | 30-40h | Examples, documentation, performance |
| **Total** | **150-200h** | **Complete PivotTable API** |

## Dependencies

- **No external dependencies**: Uses existing generated elements and parts
- **Internal dependencies**:
  - 1.2 depends on 1.1 (validation needs cache for field lookup)
  - 1.4 depends on 1.1-1.3 (serialization tests need complete API)
  - 1.5 depends on 1.1-1.4 (integration tests need all Phase 1 features)
  - Phase 2 depends on Phase 1 completion
  - Phase 3 depends on Phase 2 completion

## Parallelization

- Tasks 1.1 and 1.3 can run in parallel (cache and field wrapper are independent)
- Tasks 2.1, 2.2, 2.3 can run in parallel (filter, sort, grouping are independent)
- Phase 3 examples can start after Phase 1 (don't need Phase 2 features)
