# Spreadsheet Pivottables Specification

## Requirements

### Requirement: PivotTableDefinition Root Element

The system SHALL provide a `PivotTableDefinition` element as the root of PivotTablePart.

#### Scenario: Access pivot table structure
- GIVEN a PivotTableDefinition element
- WHEN properties are accessed
- THEN Name, CacheId, DataOnRows, DataPosition, AutoFormatId, ApplyNumberFormats, ApplyBorderFormats, ApplyFontFormats, ApplyPatternFormats, ApplyAlignmentFormats, ApplyWidthHeightFormats, DataCaption, GrandTotalCaption, ErrorCaption, ShowError, MissingCaption, ShowMissing, PageStyle, PivotTableStyle, VacatedStyle, Tag, UpdatedVersion, MinRefreshableVersion, AsteriskTotals, ShowItems, EditData, DisableFieldList, ShowCalcMbrs, VisualTotals, ShowMultipleLabel, ShowDataDropDown, ShowDrill, PrintDrill, ShowMemberPropertyTips, ShowDataTips, EnableWizard, EnableDrill, EnableFieldProperties, PreserveFormatting, UseAutoFormatting, PageWrap, PageOverThenDown, SubtotalHiddenItems, RowGrandTotals, ColGrandTotals, FieldPrintTitles, ItemPrintTitles, MergeItem, ShowDropZones, CreatedVersion, Indent, ShowEmptyRow, ShowEmptyCol, ShowHeaders, Compact, Outline, OutlineData, CompactData, Published, GridDropZones, Immersive, MultipleFieldFilters, ChartFormat, RowHeaderCaption, ColHeaderCaption, FieldListSortAscending, MdxSubqueries, CustomListSort are available

### Requirement: Pivot Cache Definition

The system SHALL provide a `PivotCacheDefinition` element as the root of PivotTableCacheDefinitionPart.

#### Scenario: Access cache definition
- GIVEN a PivotCacheDefinition element
- WHEN properties are accessed
- THEN Id (r:id to records), Invalid, SaveData, RefreshOnLoad, OptimizeMemory, EnableRefresh, RefreshedBy, RefreshedDate, BackgroundQuery, MissingItemsLimit, CreatedVersion, RefreshedVersion, MinRefreshableVersion, RecordCount, UpgradeOnRefresh, TupleCache, SupportSubquery, SupportAdvancedDrill, CacheSource, CacheFields, CacheHierarchies, Kpis, TupleCache, CalculatedItems, CalculatedMembers, Dimensions, MeasureGroups, Maps are available

#### Scenario: Cache source types
- GIVEN a CacheSource element
- WHEN Type is accessed
- THEN values Worksheet, External, Consolidation, Scenario are supported

#### Scenario: Worksheet source
- GIVEN a CacheSource with Type Worksheet
- WHEN WorksheetSource child is accessed
- THEN Ref (range), Sheet, Name (defined name) are available

### Requirement: Cache Fields

The system SHALL provide cache field definitions.

#### Scenario: CacheField properties
- GIVEN a CacheField element
- WHEN properties are accessed
- THEN Name, Caption, PropertyName, ServerField, UniqueList, NumFmtId, Formula, SqlType, Hierarchy, Level, DatabaseField, MappingCount, MemberPropertyField, SharedItems, FieldGroup, MemberPropertiesMap are available

#### Scenario: Shared items
- GIVEN a SharedItems element in CacheField
- WHEN items are accessed
- THEN Boolean, DateTime, Error, Missing, Number, String items with Value attribute are available
- AND ContainsSemiMixedTypes, ContainsNonDate, ContainsDate, ContainsString, ContainsBlank, ContainsMixedTypes, ContainsNumber, ContainsInteger, MinValue, MaxValue, MinDate, MaxDate, Count, LongText are available

### Requirement: Pivot Cache Records

The system SHALL provide a `PivotCacheRecords` element for cache data.

#### Scenario: Access cache records
- GIVEN a PivotCacheRecords element
- WHEN records are accessed
- THEN Record elements containing Boolean, DateTime, Error, Missing, Number, String, Index values are available

### Requirement: Pivot Fields

The system SHALL provide pivot field definitions.

#### Scenario: PivotField properties
- GIVEN a PivotField element in PivotFields
- WHEN properties are accessed
- THEN Name, Axis (axisRow/axisCol/axisPage/axisValues), DataField, SubtotalCaption, ShowDropDowns, HiddenLevel, UniqueMemberProperty, Compact, AllDrilled, NumFmtId, Outline, SubtotalTop, DragToRow, DragToCol, MultipleItemSelectionAllowed, DragToPage, DragToData, DragOff, ShowAll, InsertBlankRow, ServerField, InsertPageBreak, AutoShow, TopAutoShow, HideNewItems, MeasureFilter, IncludeNewItemsInFilter, ItemPageCount, SortType, DataSourceSort, NonAutoSortDefault, RankBy, DefaultSubtotal, SumSubtotal, CountASubtotal, AvgSubtotal, MaxSubtotal, MinSubtotal, ProductSubtotal, CountSubtotal, StdDevSubtotal, StdDevPSubtotal, VarSubtotal, VarPSubtotal, ShowPropCell, ShowPropTip, ShowPropAsCaption, DefaultAttributeDrillState, Items, AutoSortScope are available

#### Scenario: Field items
- GIVEN a PivotField with Items
- WHEN Item elements are accessed
- THEN Type, Hidden, Character, Child, CalculatedMember, MissingValue, StringValue, HideDetail, Detail, Expanded, DrillAcross are available

### Requirement: Row and Column Fields

The system SHALL provide row/column field organization.

#### Scenario: RowFields element
- GIVEN a RowFields element
- WHEN Field elements are accessed
- THEN Index (into PivotFields) is available for each field

#### Scenario: ColFields element
- GIVEN a ColFields element
- WHEN Field elements are accessed
- THEN Index (into PivotFields) is available for each field

#### Scenario: Special field values
- GIVEN a Field element
- WHEN Index is -2
- THEN it represents the "Values" field (data fields)

### Requirement: Page Fields

The system SHALL provide page (filter) field definitions.

#### Scenario: PageFields element
- GIVEN a PageFields element
- WHEN PageField elements are accessed
- THEN Field (index into PivotFields), Item (selected item index), Hierarchy, Name, Caption are available

### Requirement: Data Fields

The system SHALL provide data (value) field definitions.

#### Scenario: DataField properties
- GIVEN a DataField element
- WHEN properties are accessed
- THEN Name, Field (index into PivotFields), Subtotal (sum/count/average/max/min/product/countNums/stdDev/stdDevp/var/varp), ShowDataAs (normal/difference/percent/percentDiff/runTotal/percentOfRow/percentOfCol/percentOfTotal/index), BaseField, BaseItem, NumFmtId are available

### Requirement: Calculated Fields

The system SHALL support calculated fields.

#### Scenario: Calculated field in cache
- GIVEN a CalculatedItem element in PivotCacheDefinition
- WHEN properties are accessed
- THEN Field, Formula are available
- AND the formula uses field references

### Requirement: Pivot Table Location

The system SHALL specify pivot table location.

#### Scenario: Location element
- GIVEN a Location element in PivotTableDefinition
- WHEN properties are accessed
- THEN Ref (cell range), FirstHeaderRow, FirstDataRow, FirstDataCol, RowPageCount, ColPageCount are available

### Requirement: Pivot Table Styles

The system SHALL support pivot table styling.

#### Scenario: PivotTableStyleInfo
- GIVEN a PivotTableStyleInfo element
- WHEN properties are accessed
- THEN Name (table style name), ShowRowHeaders, ShowColHeaders, ShowRowStripes, ShowColStripes, ShowLastColumn are available

### Requirement: Slicers

The system SHALL support slicer controls for pivot tables.

#### Scenario: Slicer definition
- GIVEN a Slicer element in SlicersPart
- WHEN properties are accessed
- THEN Name, Cache (r:id to SlicerCachePart), Caption, StartItem, ColumnCount, ShowCaption, Level, Style, LockedPosition, RowHeight are available

#### Scenario: Slicer cache
- GIVEN a SlicerCacheDefinition element
- WHEN properties are accessed
- THEN Name, SourceName, PivotTables (references), Data (SlicerCacheData) are available

### Requirement: Timelines

The system SHALL support timeline controls.

#### Scenario: Timeline definition
- GIVEN a Timeline element in TimeLinePart
- WHEN properties are accessed
- THEN Name, Cache (r:id), Caption, ShowHeader, ShowSelectionLabel, ShowTimeLevel, ShowHorizontalScrollbar, Level, SelectionLevel, ScrollPosition, Style are available

#### Scenario: Timeline cache
- GIVEN a TimeLineCacheDefinition element
- WHEN properties are accessed
- THEN Name, SourceName, PivotTables, PivotFilter, BoundsData, SelectionData are available

### Requirement: Pivot Table API

The system SHALL provide a high-level API for pivot table creation.

#### Scenario: Create pivot table from range
- GIVEN a source data range and destination cell
- WHEN `sheet.AddPivotTable("A1:D100", "Sheet2!A1")` is called
- THEN a PivotTablePart and PivotTableCacheDefinitionPart are created
- AND the pivot table references the source range

#### Scenario: Add row field
- GIVEN a pivot table
- WHEN `pivot.AddRowField("Category")` is called
- THEN the field is added to RowFields
- AND the PivotField is configured appropriately

#### Scenario: Add data field
- GIVEN a pivot table
- WHEN `pivot.AddDataField("Sales", AggregateSum)` is called
- THEN a DataField is added with Subtotal="sum"
- AND the field name is "Sum of Sales"

