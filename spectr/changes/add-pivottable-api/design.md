# Design: PivotTable API Completion

## Context

goffice has **existing PivotTable infrastructure** (~560 LOC) that provides:

**Existing (DO NOT REBUILD):**
- `PivotTable` struct with sheet reference, pivot part, source range, destination cell
- `AddRowField()`, `AddColumnField()`, `AddDataField()`, `AddPageField()` with fluent chaining
- `Remove*Field()` methods returning `bool`
- `RowFields()`, `ColumnFields()`, `PageFields()` returning `[]string`
- `AggregateFunction` enum with all 11 subtotal types
- `newPivotTable()` factory function
- `updateFields()` for synchronizing wrapper state to elements
- Part infrastructure: `PivotTablePart`, `PivotCacheDefinitionPart`, `PivotCacheRecordsPart`

**Missing (BUILD THESE):**
1. Cache data population from source range
2. Field validation against cache definitions
3. PivotField property exposure (Axis, ShowAll, Compact, etc.)
4. Complete `updateFields()` serialization
5. `Refresh()` implementation
6. Error returns for validation

The solution must **enhance existing code**, not replace it.

## Goals

1. **Complete cache population** - Extract data from source range into PivotCacheRecords
2. **Add field validation** - Verify field names exist in source data before adding
3. **Expose PivotField properties** - Access individual field settings (10 critical properties)
4. **Complete XML serialization** - Ensure roundtrip fidelity
5. **Add error handling** - Provide `TryAdd*Field()` variants with validation

## Non-Goals

- Replace existing `PivotTable` wrapper (it works, enhance it)
- OLAP/cube features (future phase)
- External data connections (future phase)
- Calculated fields (Phase 2)
- Field grouping (Phase 2)
- PivotFilters API (Phase 2)

## Architecture Overview

### Existing Code (DO NOT MODIFY STRUCTURE)

```
spreadsheet/
├── pivottable.go        // PivotTable struct, accessors (117 lines)
├── pivottable_fields.go // Add/Remove field methods (155 lines)
└── pivottable_update.go // updateFields() implementation (~300 lines)

spreadsheet/parts/
├── pivot_table_part.go              // PivotTablePart
├── pivot_cache_definition_part.go   // PivotCacheDefinitionPart
└── pivot_cache_records_part.go      // PivotCacheRecordsPart
```

### New Code (ADD THESE)

```
spreadsheet/
├── pivottable_cache.go   // NEW: Cache population, Refresh()
├── pivottable_field.go   // NEW: PivotField wrapper for property access
└── pivottable_errors.go  // NEW: Error types (ErrFieldNotFound, etc.)
```

### Enhancement to Existing Code

```go
// In pivottable.go - ADD methods:
func (p *PivotTable) Refresh() error
func (p *PivotTable) Field(name string) *PivotField
func (p *PivotTable) Fields() []*PivotField
func (p *PivotTable) TryAddRowField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddColumnField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddDataField(name string, agg AggregateFunction) (*PivotTable, error)
func (p *PivotTable) TryAddPageField(name string) (*PivotTable, error)

// Keep existing methods unchanged for backward compatibility:
func (p *PivotTable) AddRowField(name string) *PivotTable  // unchanged
```

## Detailed Design Decisions

### Decision 1: Cache Population Strategy

**Existing State**: Cache parts exist but no data extraction from source range.

**Design**:
```go
// spreadsheet/pivottable_cache.go
func (p *PivotTable) Refresh() error {
    // 1. Get source range reference
    source := p.SourceRange()  // existing method

    // 2. Parse range into sheet + cell reference
    sheetName, cellRef, err := parseRangeReference(source)
    if err != nil {
        return fmt.Errorf("parse source range %q: %w", source, err)
    }

    // 3. Get worksheet containing data
    sheet := p.sheet.document.Sheet(sheetName)
    if sheet == nil {
        return fmt.Errorf("source sheet %q not found", sheetName)
    }

    // 4. Extract header row for field names
    headers := extractHeaderRow(sheet, cellRef)

    // 5. Extract data rows
    rows := extractDataRows(sheet, cellRef)

    // 6. Populate cache definition with field names
    cacheDef := p.getCacheDefinitionPart()
    populateCacheFields(cacheDef, headers)

    // 7. Populate cache records with data
    cacheRecords := p.getCacheRecordsPart()
    populateCacheRecords(cacheRecords, rows, headers)

    return nil
}
```

**Rationale**:
- Explicit `Refresh()` call matches Excel behavior
- Allows batching field changes before expensive data extraction
- Lazy: doesn't extract on every field add

### Decision 2: Field Validation

**Design**: Provide `TryAdd*Field()` methods alongside existing methods:

```go
// spreadsheet/pivottable.go

// TryAddRowField validates field exists before adding
func (p *PivotTable) TryAddRowField(name string) (*PivotTable, error) {
    if err := p.validateFieldName(name); err != nil {
        return p, err
    }
    if p.containsField(name) {
        return p, ErrDuplicateField{Field: name}
    }
    p.rowFields = append(p.rowFields, name)
    p.updateFields()
    return p, nil
}

// AddRowField keeps existing behavior (no validation, for backward compat)
func (p *PivotTable) AddRowField(name string) *PivotTable {
    p.rowFields = append(p.rowFields, name)
    p.updateFields()
    return p
}

// validateFieldName checks field exists in cache definition
func (p *PivotTable) validateFieldName(name string) error {
    cacheDef := p.getCacheDefinitionPart()
    if cacheDef == nil {
        return nil  // No cache yet, allow (cache populated later)
    }
    for _, field := range cacheDef.CacheFields() {
        if field.Name() == name {
            return nil
        }
    }
    return ErrFieldNotFound{Field: name}
}
```

**Rationale**:
- Backward compatible: `Add*Field()` unchanged
- New `TryAdd*Field()` provides validation for users who want it
- Error types allow programmatic error handling

### Decision 3: PivotField Property Access

**Design**: `pivot.Field(name)` returns a `*PivotField` wrapper:

```go
// spreadsheet/pivottable_field.go

type PivotField struct {
    pivot    *PivotTable
    element  *elements.CT_PivotField
    index    uint32
}

// 10 critical properties (Phase 1)
func (f *PivotField) Name() string
func (f *PivotField) Index() uint32
func (f *PivotField) Axis() AxisType
func (f *PivotField) ShowAll() bool
func (f *PivotField) SetShowAll(v bool)
func (f *PivotField) Compact() bool
func (f *PivotField) SetCompact(v bool)
func (f *PivotField) HideNewItems() bool
func (f *PivotField) SetHideNewItems(v bool)
func (f *PivotField) ShowDropDowns() bool
func (f *PivotField) SetShowDropDowns(v bool)
func (f *PivotField) DefaultSubtotal() bool
func (f *PivotField) SetDefaultSubtotal(v bool)
func (f *PivotField) DataType() FieldDataType
func (f *PivotField) ItemCount() int

// Access from PivotTable
func (p *PivotTable) Field(name string) *PivotField {
    ptDef := p.pivotPart.PivotTableDefinition()
    if ptDef == nil {
        return nil
    }
    fields := ptDef.PivotFields()
    if fields == nil {
        return nil
    }
    for i, f := range fields.PivotField() {
        if f.Name() == name {
            return &PivotField{
                pivot:   p,
                element: f,
                index:   uint32(i),
            }
        }
    }
    return nil
}

func (p *PivotTable) Fields() []*PivotField {
    // Return all fields as wrappers
}
```

**Rationale**:
- Matches goffice pattern: `sheet.Cell()`, `row.Cell()`
- Returns nil if not found (Go convention)
- 10 properties cover 90% of use cases
- More properties can be added in Phase 2 without API change

### Decision 4: Error Types

**Design**: Define sentinel error types for specific error handling:

```go
// spreadsheet/pivottable_errors.go

// ErrFieldNotFound indicates a field name doesn't exist in the pivot cache
type ErrFieldNotFound struct {
    Field string
}

func (e ErrFieldNotFound) Error() string {
    return fmt.Sprintf("pivot field %q not found in cache", e.Field)
}

// ErrDuplicateField indicates a field is already assigned to an axis
type ErrDuplicateField struct {
    Field string
}

func (e ErrDuplicateField) Error() string {
    return fmt.Sprintf("pivot field %q already assigned to an axis", e.Field)
}

// ErrInvalidSourceRange indicates the source range is invalid
type ErrInvalidSourceRange struct {
    Range string
    Cause error
}

func (e ErrInvalidSourceRange) Error() string {
    return fmt.Sprintf("invalid source range %q: %v", e.Range, e.Cause)
}
```

**Rationale**:
- Allows `errors.As()` for programmatic handling
- Clear error messages for debugging
- Matches Go error handling patterns

## Module Organization

```
spreadsheet/
├── pivottable.go           // EXISTING: Add Refresh(), Field(), TryAdd*()
├── pivottable_fields.go    // EXISTING: Unchanged
├── pivottable_update.go    // EXISTING: Complete updateFields() implementation
├── pivottable_cache.go     // NEW: Cache population logic (~150 lines)
├── pivottable_field.go     // NEW: PivotField wrapper (~100 lines)
└── pivottable_errors.go    // NEW: Error types (~50 lines)
```

**Total new code**: ~300 LOC
**Total enhanced code**: ~200 LOC modifications

## Testing Strategy

### Unit Tests (Per-Feature, Not Batched)

**Cache Population** (8 tests):
- [ ] Refresh() extracts header row from source range
- [ ] Refresh() populates cache fields with correct names
- [ ] Refresh() extracts data rows into cache records
- [ ] Refresh() handles empty source range
- [ ] Refresh() returns error for invalid source range
- [ ] Refresh() returns error for missing source sheet
- [ ] Refresh() preserves existing field assignments
- [ ] Refresh() clears and rebuilds records on re-refresh

**Field Validation** (10 tests):
- [ ] TryAddRowField() succeeds for valid field
- [ ] TryAddRowField() returns ErrFieldNotFound for invalid field
- [ ] TryAddRowField() returns ErrDuplicateField for duplicate
- [ ] TryAddColumnField() succeeds for valid field
- [ ] TryAddDataField() succeeds with aggregation function
- [ ] TryAddPageField() validates field exists
- [ ] AddRowField() still works without validation (backward compat)
- [ ] validateFieldName() allows any field before cache populated
- [ ] validateFieldName() validates after Refresh()
- [ ] containsField() detects duplicates across all axes

**PivotField Properties** (12 tests):
- [ ] Field() returns nil for non-existent field
- [ ] Field() returns correct wrapper for existing field
- [ ] Name() returns field name
- [ ] Index() returns correct 0-based index
- [ ] Axis() returns correct axis (Row/Column/Page/Data/None)
- [ ] ShowAll() returns default value
- [ ] SetShowAll() persists value
- [ ] Compact() returns default value
- [ ] SetCompact() persists value
- [ ] DefaultSubtotal() returns default value
- [ ] SetDefaultSubtotal() persists value
- [ ] ItemCount() returns correct count

**Error Types** (6 tests):
- [ ] ErrFieldNotFound.Error() formats correctly
- [ ] ErrDuplicateField.Error() formats correctly
- [ ] ErrInvalidSourceRange.Error() includes cause
- [ ] errors.As() works for ErrFieldNotFound
- [ ] errors.As() works for ErrDuplicateField
- [ ] errors.As() works for ErrInvalidSourceRange

### Integration Tests (8 scenarios)

- [ ] Create pivot → Refresh() → add fields → save → open → verify
- [ ] Open existing Excel pivot → read field properties → verify values
- [ ] Modify existing pivot → change field properties → save → verify in Excel
- [ ] Add invalid field → verify error → add valid field → succeeds
- [ ] Multiple axes with same field → verify error
- [ ] Change data field aggregation → verify persists
- [ ] Large dataset (10k rows) → Refresh() completes <2s
- [ ] Field ordering preserved through save/open cycle

### Roundtrip Tests (4 Excel versions)

- [ ] Excel 2010 pivot file → open → modify → save → reopen in Excel
- [ ] Excel 2016 pivot file → roundtrip verification
- [ ] Excel 2019 pivot file → roundtrip verification
- [ ] Excel 365 pivot file → roundtrip verification

## API Surface Summary

### New Methods on PivotTable

```go
// Cache operations
func (p *PivotTable) Refresh() error

// Field access
func (p *PivotTable) Field(name string) *PivotField
func (p *PivotTable) Fields() []*PivotField

// Validated field operations (new)
func (p *PivotTable) TryAddRowField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddColumnField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddDataField(name string, agg AggregateFunction) (*PivotTable, error)
func (p *PivotTable) TryAddPageField(name string) (*PivotTable, error)

// Existing methods unchanged for backward compat
func (p *PivotTable) AddRowField(name string) *PivotTable
func (p *PivotTable) AddColumnField(name string) *PivotTable
func (p *PivotTable) AddDataField(name string, agg AggregateFunction) *PivotTable
func (p *PivotTable) AddPageField(name string) *PivotTable
```

### New Type: PivotField

```go
type PivotField struct {
    // Accessors (read)
    func (f *PivotField) Name() string
    func (f *PivotField) Index() uint32
    func (f *PivotField) Axis() AxisType
    func (f *PivotField) ShowAll() bool
    func (f *PivotField) Compact() bool
    func (f *PivotField) HideNewItems() bool
    func (f *PivotField) ShowDropDowns() bool
    func (f *PivotField) DefaultSubtotal() bool
    func (f *PivotField) DataType() FieldDataType
    func (f *PivotField) ItemCount() int

    // Mutators (write)
    func (f *PivotField) SetShowAll(v bool)
    func (f *PivotField) SetCompact(v bool)
    func (f *PivotField) SetHideNewItems(v bool)
    func (f *PivotField) SetShowDropDowns(v bool)
    func (f *PivotField) SetDefaultSubtotal(v bool)
}
```

### New Error Types

```go
type ErrFieldNotFound struct { Field string }
type ErrDuplicateField struct { Field string }
type ErrInvalidSourceRange struct { Range string; Cause error }
```

## Backward Compatibility

**Unchanged APIs** (backward compatible):
- `AddRowField()`, `AddColumnField()`, `AddDataField()`, `AddPageField()` - same signatures
- `Remove*Field()` - same signatures
- `RowFields()`, `ColumnFields()`, `PageFields()` - same return types
- `Name()`, `SetName()`, `SourceRange()`, `DestinationCell()` - unchanged

**New APIs** (additive):
- `TryAdd*Field()` - validation variants
- `Refresh()` - cache population
- `Field()`, `Fields()` - property access
- `PivotField` type - new wrapper

**No Breaking Changes**: Existing code continues to work unchanged.

## Success Metrics

1. **Backward compatible**: Existing tests pass without modification
2. **Cache functional**: `Refresh()` populates cache from source range
3. **Validation works**: `TryAddRowField("invalid")` returns `ErrFieldNotFound`
4. **Properties accessible**: `pivot.Field("Product").Compact()` returns value
5. **Roundtrip works**: Create → save → open → modify → save produces valid Excel
6. **Performance**: `Refresh()` on 10k rows completes <2s
7. **Test coverage**: >80% on new code paths

## Phase 2 Extensions (NOT IN THIS CHANGE)

- Filter API (`PivotFilters`, `AutoSortScope`)
- Field Grouping (`FieldGroup`, `RangeProperties`, `DateGroupItem`)
- Calculated Fields (`AddCalculatedField(name, formula)`)
- Slicer Integration
