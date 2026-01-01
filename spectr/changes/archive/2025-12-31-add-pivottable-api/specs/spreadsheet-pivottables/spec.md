# Delta Spec: PivotTable API Completion

**Status**: ENHANCEMENT to existing `spreadsheet-pivottables` capability
**Base Spec**: `spectr/specs/spreadsheet-pivottables/spec.md` (provides element definitions)
**Delta**: Completes high-level API by adding cache population, field validation, PivotField properties

## Baseline Acknowledgment

goffice has **existing PivotTable infrastructure** (~560 LOC) that this spec **enhances**:
- `PivotTable` struct with sheet reference, pivot part, source range, destination cell
- `AddRowField()`, `AddColumnField()`, `AddDataField()`, `AddPageField()` methods with fluent chaining
- `Remove*Field()` methods returning `bool`
- `AggregateFunction` enum with all 11 subtotal types
- Part infrastructure: `PivotTablePart`, `PivotCacheDefinitionPart`, `PivotCacheRecordsPart`

## ADDED Requirements

### Requirement: Cache Population from Source Range

The system SHALL populate pivot cache from worksheet source data.

#### Scenario: Refresh cache from source range
- GIVEN a PivotTable with source range "Data!A1:Z1000"
- WHEN `pivot.Refresh()` is called
- THEN header row is extracted for field names
- AND data rows are extracted into cache records
- AND PivotCacheRecords part is populated
- AND error is returned if source range is invalid

#### Scenario: Refresh handles empty source
- GIVEN a PivotTable with empty source range
- WHEN `pivot.Refresh()` is called
- THEN cache is cleared
- AND no error is returned (empty state is valid)

#### Scenario: Refresh handles missing sheet
- GIVEN a PivotTable with source referencing deleted sheet
- WHEN `pivot.Refresh()` is called
- THEN `ErrInvalidSourceRange` is returned with cause
- AND existing cache is NOT modified

### Requirement: Field Validation with Error Returns

The system SHALL provide validated field operations alongside existing methods.

#### Scenario: TryAddRowField succeeds for valid field
- GIVEN a PivotTable with populated cache containing "Product" field
- WHEN `pivot.TryAddRowField("Product")` is called
- THEN field is added to row fields slice
- AND `(*PivotTable, nil)` is returned
- AND `updateFields()` is called to sync elements

#### Scenario: TryAddRowField fails for unknown field
- GIVEN a PivotTable with populated cache NOT containing "Unknown"
- WHEN `pivot.TryAddRowField("Unknown")` is called
- THEN `(pivot, ErrFieldNotFound{Field: "Unknown"})` is returned
- AND row fields slice is NOT modified

#### Scenario: TryAddRowField fails for duplicate field
- GIVEN a PivotTable with "Product" already in row fields
- WHEN `pivot.TryAddRowField("Product")` is called
- THEN `(pivot, ErrDuplicateField{Field: "Product"})` is returned
- AND row fields slice is NOT modified

#### Scenario: Existing AddRowField unchanged (backward compat)
- GIVEN a PivotTable (cache populated or not)
- WHEN `pivot.AddRowField("Product")` is called
- THEN field is added to row fields slice unconditionally
- AND `*PivotTable` is returned (same as before)
- AND no error is returned (silent behavior preserved)

#### Scenario: Validation accepts any field before cache populated
- GIVEN a PivotTable with no cache (never refreshed)
- WHEN `pivot.TryAddRowField("AnyField")` is called
- THEN field is added (validation skipped when no cache)
- AND `(*PivotTable, nil)` is returned

### Requirement: PivotField Property Access

The system SHALL provide PivotField wrapper for accessing individual field properties.

#### Scenario: Access field by name
- GIVEN a PivotTable with populated fields
- WHEN `field := pivot.Field("Product")` is called
- THEN `*PivotField` wrapper is returned
- AND wrapper provides access to underlying CT_PivotField element
- AND nil is returned if field not found

#### Scenario: List all fields
- GIVEN a PivotTable with populated fields
- WHEN `fields := pivot.Fields()` is called
- THEN `[]*PivotField` slice is returned
- AND slice contains wrapper for each field in PivotFields collection
- AND order matches definition order

#### Scenario: Read PivotField properties
- GIVEN a PivotField wrapper
- WHEN property accessors are called:
  - `field.Name()` returns field name
  - `field.Index()` returns 0-based index in PivotFields
  - `field.Axis()` returns AxisRow/AxisColumn/AxisPage/AxisData/AxisNone
  - `field.ShowAll()` returns showAll attribute value
  - `field.Compact()` returns compact attribute value
  - `field.DefaultSubtotal()` returns defaultSubtotal attribute value
  - `field.HideNewItems()` returns hideNewItems attribute value
  - `field.ShowDropDowns()` returns showDropDowns attribute value
  - `field.DataType()` returns FieldDataString/FieldDataNumber/FieldDataDate
  - `field.ItemCount()` returns count of items in field

#### Scenario: Modify PivotField properties
- GIVEN a PivotField wrapper
- WHEN property setters are called:
  - `field.SetShowAll(true)`
  - `field.SetCompact(false)`
  - `field.SetDefaultSubtotal(true)`
  - `field.SetHideNewItems(true)`
  - `field.SetShowDropDowns(false)`
- THEN underlying CT_PivotField element attributes are modified
- AND changes persist when workbook is saved

### Requirement: Error Type Definitions

The system SHALL define specific error types for programmatic error handling.

#### Scenario: ErrFieldNotFound
- GIVEN field validation fails for unknown field
- WHEN error is returned
- THEN `errors.As(err, &ErrFieldNotFound{})` returns true
- AND `err.Error()` returns `pivot field "FieldName" not found in cache`

#### Scenario: ErrDuplicateField
- GIVEN field validation fails for duplicate
- WHEN error is returned
- THEN `errors.As(err, &ErrDuplicateField{})` returns true
- AND `err.Error()` returns `pivot field "FieldName" already assigned to an axis`

#### Scenario: ErrInvalidSourceRange
- GIVEN source range validation fails
- WHEN error is returned
- THEN `errors.As(err, &ErrInvalidSourceRange{})` returns true
- AND `err.Cause` contains underlying parse error
- AND `err.Error()` returns `invalid source range "ref": cause`

### Requirement: Aggregation Function Support

The system SHALL support all 11 ECMA-376 subtotal types (ALREADY EXISTS - validation only).

#### Scenario: All aggregation types work
- GIVEN a PivotTable
- WHEN `pivot.AddDataField(name, agg)` is called with any of:
  - AggregateSUM → "sum"
  - AggregateCount → "count"
  - AggregateAverage → "average"
  - AggregateMin → "min"
  - AggregateMax → "max"
  - AggregateProduct → "product"
  - AggregateCountNums → "countNums"
  - AggregateStdDev → "stdDev"
  - AggregateStdDevP → "stdDevP"
  - AggregateVar → "var"
  - AggregateVarP → "varP"
- THEN correct subtotal XML value is set

### Requirement: Roundtrip Persistence

The system SHALL preserve pivot structure through save/open cycles.

#### Scenario: Create → Save → Open → Verify
- GIVEN pivot created with row/column/data fields
- WHEN workbook is saved, closed, and reopened
- THEN `pivot.RowFields()` returns same field names in same order
- AND `pivot.ColumnFields()` returns same field names
- AND data field aggregations are preserved

#### Scenario: Excel compatibility
- GIVEN pivot created and saved by goffice
- WHEN file is opened in Excel 2010/2016/2019/365
- THEN no repair warnings are shown
- AND pivot displays correctly
- AND pivot can be modified in Excel

## Implementation Notes

### File Structure (Enhancement to existing)
```
spreadsheet/
├── pivottable.go           // EXISTING: Add Refresh(), Field(), TryAdd*()
├── pivottable_fields.go    // EXISTING: Unchanged
├── pivottable_update.go    // EXISTING: Complete updateFields() serialization
├── pivottable_cache.go     // NEW: Cache population logic (~150 LOC)
├── pivottable_field.go     // NEW: PivotField wrapper (~100 LOC)
└── pivottable_errors.go    // NEW: Error types (~50 LOC)
```

### API Surface Summary (Additions Only)

```go
// NEW Methods on existing PivotTable struct

// Cache operations
func (p *PivotTable) Refresh() error

// Field property access
func (p *PivotTable) Field(name string) *PivotField
func (p *PivotTable) Fields() []*PivotField

// Validated field operations (new)
func (p *PivotTable) TryAddRowField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddColumnField(name string) (*PivotTable, error)
func (p *PivotTable) TryAddDataField(name string, agg AggregateFunction) (*PivotTable, error)
func (p *PivotTable) TryAddPageField(name string) (*PivotTable, error)

// EXISTING methods remain unchanged (backward compat)
func (p *PivotTable) AddRowField(name string) *PivotTable
func (p *PivotTable) AddColumnField(name string) *PivotTable
func (p *PivotTable) AddDataField(name string, agg AggregateFunction) *PivotTable
func (p *PivotTable) AddPageField(name string) *PivotTable
func (p *PivotTable) RowFields() []string
func (p *PivotTable) ColumnFields() []string
func (p *PivotTable) PageFields() []string
func (p *PivotTable) RemoveRowField(name string) bool
// ... etc
```

### NEW Type: PivotField

```go
type PivotField struct {
    pivot   *PivotTable
    element *elements.CT_PivotField
    index   uint32
}

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
```

### NEW Types: Enums

```go
type AxisType int
const (
    AxisNone AxisType = iota
    AxisRow
    AxisColumn
    AxisPage
    AxisData
)

type FieldDataType int
const (
    FieldDataString FieldDataType = iota
    FieldDataNumber
    FieldDataDate
    FieldDataBoolean
)
```

### NEW Types: Errors

```go
type ErrFieldNotFound struct { Field string }
type ErrDuplicateField struct { Field string }
type ErrInvalidSourceRange struct { Range string; Cause error }
```

## Backward Compatibility

**Unchanged APIs** (fully backward compatible):
- `AddRowField()`, `AddColumnField()`, `AddDataField()`, `AddPageField()` - same signatures
- `Remove*Field()` - same signatures and return types
- `RowFields()`, `ColumnFields()`, `PageFields()` - return `[]string` (not FieldAxis)
- `Name()`, `SetName()`, `SourceRange()`, `DestinationCell()` - unchanged

**New APIs** (additive only):
- `TryAdd*Field()` - validation variants returning error
- `Refresh()` - cache population
- `Field()`, `Fields()` - property access
- `PivotField` type - new wrapper

**No Breaking Changes**: Existing code continues to work unchanged.

## Phase 2+ Enhancements (NOT IN THIS CHANGE)

Deferred to future proposals:
- Filter API (`PivotFilters`, `AutoSortScope`)
- Field Grouping (`FieldGroup`, `RangeProperties`, `DateGroupItem`)
- Calculated Fields (`AddCalculatedField(name, formula)`)
- Builder pattern redesign (if requested)
- Slicer Integration

## Testing Requirements

**Unit Tests (36+ cases)**: Per design.md section
- Cache population: 8 tests
- Field validation: 10 tests
- PivotField properties: 12 tests
- Error types: 6 tests

**Integration Tests (8 scenarios)**: Per design.md section
- Create/save/open/verify roundtrip
- Excel file read/modify/save
- Large dataset performance

**Roundtrip Tests (4 Excel versions)**:
- Excel 2010, 2016, 2019, 365 compatibility

See `tasks.md` for detailed test plan with specific test cases.
