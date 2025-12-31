## ADDED Requirements

### Requirement: Office 2010 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2010 extensions.

#### Scenario: Slicer elements
- GIVEN Office 2010 slicer elements (x14 namespace)
- WHEN parsed from XML
- THEN Slicer, SlicerCache, SlicerCacheDefinition elements are created

#### Scenario: Conditional formatting extensions
- GIVEN Office 2010 conditional formatting extensions (x14 namespace)
- WHEN parsed
- THEN ConditionalFormatting, DataBar, IconSet, ColorScale elements available

#### Scenario: Sparkline elements
- GIVEN Office 2010 sparkline elements (x14 namespace)
- WHEN parsed
- THEN Sparkline, SparklineGroup elements available

#### Scenario: PivotTable extensions
- GIVEN Office 2010 pivot table extensions (x14 namespace)
- WHEN parsed
- THEN PivotTableDefinition extensions available

### Requirement: Office 2013-2015 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2013-2015 extensions.

#### Scenario: Timeline elements
- GIVEN Office 2013 timeline elements (x15 namespace)
- WHEN parsed
- THEN Timeline, TimelineCacheDefinition, TimelineStyle elements available

#### Scenario: Chart elements
- GIVEN Office 2015 chart extensions (x15 namespace)
- WHEN parsed
- THEN extended chart elements available

#### Scenario: Data model elements
- GIVEN Office 2015 data model extensions (x15 namespace)
- WHEN parsed
- THEN DataModel, ModelTable, ModelRelationship elements available

### Requirement: Office 2016 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2016 extensions.

#### Scenario: Revision elements (x16r5)
- GIVEN Office 2016 revision elements (x16r5 namespace)
- WHEN parsed
- THEN revision tracking elements available

#### Scenario: Pivot default layout elements
- GIVEN Office 2016 pivot default layout elements (x16pdl namespace)
- WHEN parsed
- THEN PivotDefaultLayout element available

#### Scenario: SVG support elements
- GIVEN Office 2016 SVG elements (x16svg namespace)
- WHEN parsed
- THEN SVG image reference elements available

### Requirement: Office 2017-2020 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2017-2020 extensions.

#### Scenario: Rich data elements
- GIVEN Office 2017 rich data elements (xlrd namespace)
- WHEN parsed
- THEN RichValue, RichValueStructure elements available

#### Scenario: Rich data web image elements
- GIVEN Office 2020 rich data web image elements (xlrdwi namespace)
- WHEN parsed
- THEN WebImageElement available

### Requirement: Office 2021-2024 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2021-2024 extensions.

#### Scenario: Rich value relation elements
- GIVEN Office 2022 rich value relation elements (xlrvrel namespace)
- WHEN parsed
- THEN RichValueRelation elements available

#### Scenario: Pivot auto refresh elements
- GIVEN Office 2024 pivot auto refresh elements (xpar namespace)
- WHEN parsed
- THEN PivotAutoRefresh element available

### Requirement: Office 2025 Spreadsheet Extension Elements
The system SHALL provide element types for Excel 2025 extensions.

#### Scenario: Pivot data source elements
- GIVEN Office 2025 pivot data source elements (x25pds namespace)
- WHEN parsed
- THEN PivotDataSource elements available

#### Scenario: External code service elements
- GIVEN Office 2025 external code service elements (x25ecs namespace)
- WHEN parsed
- THEN ExternalCodeService elements available

### Requirement: Spreadsheet Extension Element Attributes
The system SHALL provide typed attributes for spreadsheet extension elements.

#### Scenario: Slicer name attribute
- GIVEN a Slicer element with name attribute
- WHEN Name is accessed
- THEN StringValue wrapper with slicer name is returned

#### Scenario: Timeline cache id attribute
- GIVEN a Timeline element with cache id
- WHEN CacheId is accessed
- THEN UInt32Value wrapper is returned

#### Scenario: Rich value type attribute
- GIVEN a RichValue element with type attribute
- WHEN Type is accessed
- THEN enum value from rich data schema is returned

### Requirement: Spreadsheet Extension Element Version Metadata
The system SHALL provide version information for spreadsheet extension elements.

#### Scenario: Office 2010 element version
- GIVEN a Slicer element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Office 2024 element version
- GIVEN a PivotAutoRefresh element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2024 is returned

#### Scenario: Office 2025 element version
- GIVEN a PivotDataSource element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2025 is returned

### Requirement: Spreadsheet Extension Element Namespace Handling
The system SHALL correctly handle namespaces for spreadsheet extension elements.

#### Scenario: Excel 2010 namespace
- GIVEN a Slicer element
- WHEN NamespaceURI() is called
- THEN "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" is returned

#### Scenario: Namespace prefix in XML
- GIVEN a Slicer element
- WHEN written to XML
- THEN element is serialized as `<x14:slicer>`

#### Scenario: Rich data namespace
- GIVEN a RichValue element
- WHEN NamespaceURI() is called
- THEN rich data namespace URI is returned

### Requirement: Spreadsheet Extension Element Generation
The system SHALL generate spreadsheet element types from extension schemas in addition to main schemas.

#### Scenario: Generate from main schema
- GIVEN spreadsheetml main.json schema
- WHEN generator runs
- THEN core SpreadsheetML elements are generated

#### Scenario: Generate from extension schemas
- GIVEN Excel 2010-2025 extension schemas
- WHEN generator runs
- THEN extension element types are generated with correct namespaces

#### Scenario: Total element count increase
- GIVEN all schemas loaded
- WHEN generation completes
- THEN 80-100 additional element types exist beyond baseline

#### Scenario: Generated element organization
- GIVEN generated spreadsheet elements
- WHEN inspecting spreadsheet/elements/ package
- THEN extension elements are organized with main elements

### Requirement: Spreadsheet Extension Element Registration
The system SHALL register all spreadsheet extension element types.

#### Scenario: Register slicer elements
- GIVEN Slicer, SlicerCache element types
- WHEN element registry is initialized
- THEN slicer elements are registered with x14 namespace

#### Scenario: Parse registered extension element
- GIVEN XML with x14:slicer element
- WHEN parsing workbook
- THEN Slicer instance is created (not UnknownElement)

#### Scenario: Parse timeline element
- GIVEN XML with x15:timeline element
- WHEN parsing workbook
- THEN Timeline instance is created
