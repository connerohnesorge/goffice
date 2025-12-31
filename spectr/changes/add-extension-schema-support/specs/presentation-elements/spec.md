## ADDED Requirements

### Requirement: Office 2010 Presentation Extension Elements
The system SHALL provide element types for PowerPoint 2010 extensions.

#### Scenario: Section elements
- GIVEN Office 2010 section elements (p14 namespace)
- WHEN parsed from XML
- THEN Section, SectionList, SectionProperties elements are created

#### Scenario: Media elements
- GIVEN Office 2010 media elements (p14 namespace)
- WHEN parsed
- THEN Media, MediaBookmark elements available

#### Scenario: Creative Commons elements
- GIVEN Office 2010 creative commons elements (p14 namespace)
- WHEN parsed
- THEN CreativeCommons element available

### Requirement: Office 2012-2013 Presentation Extension Elements
The system SHALL provide element types for PowerPoint 2012-2013 extensions.

#### Scenario: Roaming settings elements
- GIVEN Office 2012 roaming settings elements (p14rs namespace)
- WHEN parsed
- THEN RoamingSettings element available

#### Scenario: Chart style elements
- GIVEN Office 2013 chart style elements (p15 namespace)
- WHEN parsed
- THEN ChartStyle elements available

### Requirement: Office 2015-2019 Presentation Extension Elements
The system SHALL provide element types for PowerPoint 2015-2019 extensions.

#### Scenario: Threading info elements
- GIVEN Office 2015 threading elements (p15 namespace)
- WHEN parsed
- THEN ThreadingInfo elements available

#### Scenario: Presentation comment elements
- GIVEN Office 2019 presentation comment elements (p19 namespace)
- WHEN parsed
- THEN extended comment elements available

### Requirement: Office 2020-2023 Presentation Extension Elements
The system SHALL provide element types for PowerPoint 2020-2023 extensions.

**Note:** PowerPoint extension schemas use monthly versioning (YYYY_MM format), where all versions within a year map to that year's Office version constant (e.g., 2022_03 and 2022_08 both map to Office2022).

#### Scenario: 2020 main extensions
- GIVEN Office 2020 main extensions (2020_02 schema)
- WHEN parsed
- THEN PowerPoint 2020 extension elements available

#### Scenario: 2021 main extensions
- GIVEN Office 2021 main extensions (2021_06 schema)
- WHEN parsed
- THEN PowerPoint 2021 extension elements available

#### Scenario: 2022 main extensions
- GIVEN Office 2022 main extensions (2022_03, 2022_08 schemas)
- WHEN parsed
- THEN PowerPoint 2022 extension elements available

#### Scenario: 2023 main extensions
- GIVEN Office 2023 main extensions (2023_02 schema)
- WHEN parsed
- THEN PowerPoint 2023 extension elements available

### Requirement: Drawing Extension Elements for Presentations
The system SHALL provide DrawingML extension elements used in presentations.

#### Scenario: Drawing 2010 elements
- GIVEN Office 2010 drawing extensions (a14 namespace)
- WHEN parsed in presentation
- THEN drawing 2010 elements available

#### Scenario: Drawing decorative elements
- GIVEN Office 2017 drawing decorative elements (adec namespace)
- WHEN parsed
- THEN Decorative element with isDecorative attribute available

#### Scenario: 3D model animation elements
- GIVEN Office 2018 3D model animation elements (am3d namespace)
- WHEN parsed
- THEN Model3D animation elements available

### Requirement: Presentation Extension Element Attributes
The system SHALL provide typed attributes for presentation extension elements.

#### Scenario: Section name attribute
- GIVEN a Section element with name attribute
- WHEN Name is accessed
- THEN StringValue wrapper with section name is returned

#### Scenario: Media bookmark time attribute
- GIVEN a MediaBookmark element with time attribute
- WHEN Time is accessed
- THEN UInt32Value wrapper is returned

#### Scenario: Decorative flag attribute
- GIVEN a Decorative element with isDecorative attribute
- WHEN IsDecorative is accessed
- THEN BoolValue wrapper is returned

### Requirement: Presentation Extension Element Version Metadata
The system SHALL provide version information for presentation extension elements.

#### Scenario: Office 2010 element version
- GIVEN a Section element
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2010 is returned

#### Scenario: Office 2022 element version
- GIVEN a PowerPoint 2022 extension element (from 2022_03 or 2022_08 schema)
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2022 is returned

#### Scenario: Office 2023 element version
- GIVEN a PowerPoint 2023 extension element (from 2023_02 schema)
- WHEN Metadata().AvailableInVersion() is called
- THEN Office2023 is returned

### Requirement: Presentation Extension Element Namespace Handling
The system SHALL correctly handle namespaces for presentation extension elements.

#### Scenario: PowerPoint 2010 namespace
- GIVEN a Section element
- WHEN NamespaceURI() is called
- THEN "http://schemas.microsoft.com/office/powerpoint/2010/main" is returned

#### Scenario: Namespace prefix in XML
- GIVEN a Section element
- WHEN written to XML
- THEN element is serialized as `<p14:section>`

#### Scenario: Drawing extension namespace
- GIVEN a Decorative element
- WHEN NamespaceURI() is called
- THEN "http://schemas.microsoft.com/office/drawing/2017/decorative" is returned

### Requirement: Presentation Extension Element Generation
The system SHALL generate presentation element types from extension schemas in addition to main schemas.

#### Scenario: Generate from main schema
- GIVEN presentationml main.json schema
- WHEN generator runs
- THEN core PresentationML elements are generated

#### Scenario: Generate from extension schemas
- GIVEN PowerPoint 2010-2023 extension schemas (including monthly versions)
- WHEN generator runs
- THEN extension element types are generated with correct namespaces

#### Scenario: Generate from drawing extensions
- GIVEN DrawingML extension schemas
- WHEN presentation generator runs
- THEN drawing extension elements used in presentations are available

#### Scenario: Total element count increase
- GIVEN all schemas loaded
- WHEN generation completes
- THEN 50-70 additional element types exist beyond baseline

### Requirement: Presentation Extension Element Registration
The system SHALL register all presentation extension element types.

#### Scenario: Register section elements
- GIVEN Section, SectionList element types
- WHEN element registry is initialized
- THEN section elements are registered with p14 namespace

#### Scenario: Parse registered extension element
- GIVEN XML with p14:section element
- WHEN parsing presentation
- THEN Section instance is created (not UnknownElement)

#### Scenario: Parse decorative element
- GIVEN XML with adec:decorative element
- WHEN parsing presentation
- THEN Decorative instance is created
