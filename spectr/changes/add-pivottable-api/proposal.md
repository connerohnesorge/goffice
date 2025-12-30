# Change: Complete PivotTable API Enhancement

## Why

goffice has **existing PivotTable infrastructure** (~560 LOC in `spreadsheet/pivottable*.go`) providing:
- `PivotTable` wrapper struct with field management
- `AddRowField()`, `AddColumnField()`, `AddDataField()`, `AddPageField()` methods with fluent chaining
- `Remove*Field()` methods returning bool
- `AggregateFunction` enum with all 11 subtotal types
- Part relationship wiring (`PivotTablePart`, `PivotCacheDefinitionPart`, `PivotCacheRecordsPart`)

**However, the current implementation lacks:**
1. **Cache population** - No mechanism to extract data from source ranges into cache records
2. **Field validation** - No verification that field names exist in source data
3. **PivotField property exposure** - Can't access Axis, ShowAll, Compact, etc.
4. **Roundtrip persistence** - XML serialization incomplete for full pivot fidelity
5. **Filter/Sort APIs** - No PivotFilters or AutoSortScope support
6. **Error handling** - Methods return `*PivotTable` or `bool`, not errors

**Impact**: 9/10 - Enterprises require pivot tables for BI dashboards, financial reporting, and data analysis. Without complete API, users must create pivots manually in Excel.

## What Changes

### Phase 1: Core Completion (60-80 hours)
Enhance existing `PivotTable` wrapper rather than replacing it:

1. **Cache Data Population** (NEW - 15-20h)
   - Extract field names/types from source range header row
   - Populate `PivotCacheRecords` with actual cell values
   - Implement `PivotTable.Refresh() error` to rebuild cache from source

2. **Field Validation** (NEW - 8-12h)
   - `validateFieldName()` against cache field definitions
   - Add error returns to field operations (optional: provide both `Add*Field` and `TryAdd*Field`)
   - Define `ErrFieldNotFound`, `ErrDuplicateField` sentinel errors

3. **PivotField Wrapper** (NEW - 12-15h)
   - `pivot.Field(name) *PivotField` accessor
   - Expose 10 critical properties: `Axis()`, `Name()`, `ShowAll()`, `Compact()`, `HideNewItems()`, `ShowDropDowns()`, `DefaultSubtotal()`, `DataType()`, `Index()`, `ItemCount()`
   - Read/modify individual field settings via `SetCompact(bool)`, etc.

4. **XML Serialization Enhancement** (8-12h)
   - Complete `updateFields()` implementation in `pivottable_update.go`
   - Ensure all PivotTableDefinition child elements serialize correctly
   - Add relationship ID management for cache linking

5. **Roundtrip Testing** (8-10h)
   - Create → save → open → verify structure tests
   - Excel compatibility testing (open in Excel, no repair warnings)
   - Byte-level XML comparison tests

### Phase 2: Advanced Features (60-80 hours)
6. **Basic Filter API** (15-20h) - Read/apply pivot field filters via `PivotFilters` element
7. **AutoSort** (8-12h) - Sort configuration via `AutoSortScope` element
8. **Field Grouping** (25-35h) - Date/numeric grouping (RangeProperties, DateGroupItem, FieldGroup)

### Phase 3: Polish (30-40 hours)
9. **Examples** (12-15h) - 3 complete working examples
10. **Documentation** (10-15h) - API docs and migration guide
11. **Performance** (8-10h) - Benchmarks for large datasets

**Breaking changes**: Consider providing both `AddRowField(name) *PivotTable` (existing) and `TryAddRowField(name) (*PivotTable, error)` (new validated version) for backward compatibility.

## Impact

- **Affected specs**: `spreadsheet-pivottables` (enhanced with completion requirements)
- **New capabilities**: Cache population, field validation, PivotField property access
- **Affected code**:
  - `spreadsheet/pivottable.go` - add Refresh(), Field() accessor
  - `spreadsheet/pivottable_cache.go` - NEW: cache population logic
  - `spreadsheet/pivottable_field.go` - NEW: PivotField wrapper
  - `spreadsheet/pivottable_update.go` - enhance XML serialization
  - `spreadsheet/parts/pivot_cache_records_part.go` - enhance cache records API
  - `examples/pivot_*` - usage demonstrations

## Key Design Decisions

1. **Enhance vs Replace**: Enhance existing `PivotTable` wrapper (560 LOC). Don't create new package.
   - Rationale: Existing code follows goffice patterns. Replacement discards working code.

2. **Error Handling Strategy**: Provide both patterns for backward compatibility
   - `AddRowField(name) *PivotTable` - existing, silently fails if field not found
   - `TryAddRowField(name) (*PivotTable, error)` - new, validates field exists
   - Rationale: Avoids breaking existing users while enabling validation.

3. **Cache Population**: Lazy population on first `Refresh()` call
   - Rationale: Matches Excel behavior; avoids unnecessary data extraction for read-only access.

4. **PivotField Access**: `pivot.Field(name) *PivotField` accessor pattern
   - Rationale: Matches `sheet.Cell()`, `row.Cell()` accessor patterns in goffice.
   - Returns nil if field not found (caller checks nil, matches Go conventions).

5. **Property Subset**: Start with 10 critical properties, add more in Phase 2
   - Axis, Name, ShowAll, Compact, HideNewItems, ShowDropDowns, DefaultSubtotal, DataType, Index, ItemCount
   - Rationale: Cover 90% of use cases. Full property set (50+) would overwhelm.

## Validation Against OpenXML-SDK

| Feature | OpenXML-SDK | goffice Current | Proposal | Alignment |
|---------|-------------|-----------------|----------|-----------|
| PivotTable wrapper | `PivotTable` class | ✅ `PivotTable` struct | Enhance | ✅ |
| Field add/remove | Collection methods | ✅ `Add*Field()` methods | Add validation | ✅ |
| Aggregation | `DataFieldValues` enum | ✅ `AggregateFunction` enum | N/A (complete) | ✅ |
| Cache refresh | `Refresh()` method | ❌ Missing | Add `Refresh()` | ✅ |
| PivotField wrapper | `PivotField` class | ❌ Missing | Add `PivotField` struct | ✅ |
| Field properties | 50+ properties | ❌ Missing | Add 10 critical properties | ⚠️ Subset |
| Filter API | `PivotFilters` | ❌ Missing | Phase 2 | ✅ Planned |
| Grouping | `GroupBy()` methods | ❌ Missing | Phase 2 | ✅ Planned |

## Success Criteria

1. ✅ `pivot.Refresh()` populates cache from source range
2. ✅ `TryAddRowField(name)` returns error if field not in cache
3. ✅ `pivot.Field(name).Compact()` returns field compactness setting
4. ✅ Roundtrip: create → save → open → modify → save produces valid Excel files
5. ✅ Excel 2010-2021 opens generated pivots without repair warnings
6. ✅ Unit test coverage >80% for pivot table code paths
7. ✅ Performance: create 50-field pivot <500ms, refresh 100k rows <2s
8. ✅ Error messages specify which field/operation failed
9. ✅ `go doc` shows complete API surface
10. ✅ 3 complete examples demonstrating common use cases

## Risk Assessment

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Backward compatibility concerns | Medium | High | Provide both old and new method signatures |
| Cache population performance | Medium | Medium | Lazy loading, streaming for large sources |
| XML serialization edge cases | High | Medium | Extensive roundtrip testing, Excel compatibility suite |
| PivotField property count | Low | Low | Start with 10, add more based on user requests |
| Excel version compatibility | Medium | High | Test suite across Excel 2010, 2016, 2019, 365 |

## Specs Changed

- `spectr/changes/add-pivottable-api/specs/spreadsheet-pivottables/spec.md` (delta)
