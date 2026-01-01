# Presentation Elements Specification

## Requirements

### Requirement: Presentation Element

The system SHALL provide the `Presentation` root element type (p:presentation).

#### Scenario: Presentation element structure
- WHEN creating a Presentation element
- THEN it contains SlideIdList, SlideMasterIdList, SlideSize, NotesSize
- AND optional elements like DefaultTextStyle, CustomShowList

#### Scenario: Slide ID management
- WHEN accessing `SlideIdList()`
- THEN slide references with IDs and relationship IDs are returned
- AND slides can be reordered by modifying the list

#### Scenario: Slide master ID management
- WHEN accessing `SlideMasterIdList()`
- THEN master slide references are returned

#### Scenario: Slide size configuration
- WHEN accessing `SlideSize()`
- THEN width and height (EMUs) are available
- AND standard sizes (4:3, 16:9, etc.) can be set

### Requirement: Slide Element

The system SHALL provide the `Slide` element type (p:sld).

#### Scenario: Slide element structure
- WHEN creating a Slide element
- THEN it contains CommonSlideData (cSld) as required child
- AND optional ColorMapOverride, Timing, Transition

#### Scenario: Common slide data access
- WHEN accessing `CommonSlideData()`
- THEN the shape tree and background are accessible

#### Scenario: Slide show attribute
- WHEN accessing `Show` attribute
- THEN it indicates if slide is shown in slideshow
- AND defaults to true

### Requirement: CommonSlideData Element

The system SHALL provide `CommonSlideData` (p:cSld) for shared slide content.

#### Scenario: Shape tree access
- WHEN accessing `ShapeTree()`
- THEN the sp tree containing all shapes is returned

#### Scenario: Background access
- WHEN accessing `Background()`
- THEN slide background settings are returned

#### Scenario: Name attribute
- WHEN accessing `Name` attribute
- THEN the slide name is returned

### Requirement: ShapeTree Element

The system SHALL provide `ShapeTree` (p:spTree) for containing shapes.

#### Scenario: Shape tree structure
- WHEN creating a ShapeTree
- THEN NonVisualGroupShapeProperties and GroupShapeProperties are required
- AND child shapes follow

#### Scenario: Shape enumeration
- WHEN accessing child shapes
- THEN Shape, Picture, GroupShape, GraphicFrame, ConnectionShape are returned

### Requirement: Shape Element

The system SHALL provide `Shape` (p:sp) for basic shapes.

#### Scenario: Shape element structure
- WHEN creating a Shape
- THEN it contains NonVisualShapeProperties, ShapeProperties
- AND optional TextBody for text content

#### Scenario: Non-visual properties
- WHEN accessing `NonVisualShapeProperties()`
- THEN name, ID, and placeholder info are accessible

#### Scenario: Shape properties
- WHEN accessing `ShapeProperties()`
- THEN transform, preset geometry, fill, outline are accessible

#### Scenario: Text body access
- WHEN accessing `TextBody()`
- THEN text content with paragraphs is accessible
- AND nil if shape has no text

### Requirement: TextBody Element

The system SHALL provide `TextBody` (p:txBody) for text content.

#### Scenario: Text body structure
- WHEN creating a TextBody
- THEN it contains BodyProperties, ListStyle (optional), and Paragraphs

#### Scenario: Body properties
- WHEN accessing `BodyProperties()`
- THEN anchor, wrap, columns, rotation settings are available

#### Scenario: Paragraph enumeration
- WHEN accessing `Paragraphs()`
- THEN all p:p elements are returned

### Requirement: Paragraph Element

The system SHALL provide `Paragraph` (a:p) for text paragraphs.

#### Scenario: Paragraph structure
- WHEN creating a Paragraph
- THEN it can contain ParagraphProperties, Runs, Line Breaks, Fields

#### Scenario: Paragraph properties
- WHEN accessing `ParagraphProperties()`
- THEN alignment, spacing, bullet settings are available

#### Scenario: Run enumeration
- WHEN accessing `Runs()` or `Elements()`
- THEN text runs and other inline elements are returned

### Requirement: Run Element

The system SHALL provide `Run` (a:r) for text runs.

#### Scenario: Run structure
- WHEN creating a Run
- THEN it contains RunProperties and Text

#### Scenario: Run properties
- WHEN accessing `RunProperties()`
- THEN font, size, bold, italic, color are available

#### Scenario: Text content
- WHEN accessing `Text()`
- THEN the a:t element content is returned

### Requirement: Picture Element

The system SHALL provide `Picture` (p:pic) for embedded images.

#### Scenario: Picture structure
- WHEN creating a Picture
- THEN it contains NonVisualPictureProperties, BlipFill, ShapeProperties

#### Scenario: Blip fill
- WHEN accessing `BlipFill()`
- THEN image reference (r:embed) is accessible

#### Scenario: Picture transform
- WHEN accessing shape properties transform
- THEN position and size are available

### Requirement: GraphicFrame Element

The system SHALL provide `GraphicFrame` (p:graphicFrame) for charts, tables, diagrams.

#### Scenario: Graphic frame structure
- WHEN creating a GraphicFrame
- THEN it contains NonVisualGraphicFrameProperties, Transform, Graphic

#### Scenario: Graphic content
- WHEN accessing `Graphic()`
- THEN the a:graphic element with chart/table/diagram data is returned

### Requirement: Table Elements

The system SHALL provide table elements (a:tbl, a:tr, a:tc).

#### Scenario: Table structure
- WHEN creating a Table
- THEN it contains TableProperties, TableGrid, TableRows

#### Scenario: Table row structure
- WHEN creating a TableRow
- THEN it contains TableCells with height

#### Scenario: Table cell structure
- WHEN creating a TableCell
- THEN it contains TextBody for cell content

### Requirement: GroupShape Element

The system SHALL provide `GroupShape` (p:grpSp) for grouped shapes.

#### Scenario: Group shape structure
- WHEN creating a GroupShape
- THEN it contains NonVisualGroupShapeProperties, GroupShapeProperties
- AND child shapes

#### Scenario: Child shape management
- WHEN adding shapes to a group
- THEN they become children of the group
- AND transforms are relative to group

### Requirement: ConnectionShape Element

The system SHALL provide `ConnectionShape` (p:cxnSp) for connectors.

#### Scenario: Connection shape structure
- WHEN creating a ConnectionShape
- THEN it contains NonVisualConnectionShapeProperties, ShapeProperties

#### Scenario: Connection points
- WHEN configuring start/end connections
- THEN shape references and connection site indices are stored

### Requirement: Transition Element

The system SHALL provide `Transition` (p:transition) for slide transitions.

#### Scenario: Transition type
- WHEN setting transition effect
- THEN various effects (fade, push, wipe, etc.) are available

#### Scenario: Transition timing
- WHEN configuring transition
- THEN duration, advance on click, advance after time are available

#### Scenario: Transition sound
- WHEN setting transition sound
- THEN sound reference can be included

### Requirement: Timing Element

The system SHALL provide `Timing` (p:timing) for animations.

#### Scenario: Timing structure
- WHEN creating Timing
- THEN it contains TimeNodeList with animation sequences

#### Scenario: Build list
- WHEN accessing build animations
- THEN paragraph and shape build effects are available

### Requirement: Animation Elements

The system SHALL provide animation elements for slide animations.

#### Scenario: Parallel time node
- WHEN creating animation sequences
- THEN parallel/sequential containers are used

#### Scenario: Animation effects
- WHEN adding animations
- THEN entrance, emphasis, exit, motion path effects are available

#### Scenario: Animation triggers
- WHEN configuring animation start
- THEN on click, with previous, after previous triggers are available

### Requirement: SlideLayout Element

The system SHALL provide `SlideLayout` (p:sldLayout) for layout definitions.

#### Scenario: Layout element structure
- WHEN creating a SlideLayout
- THEN it contains CommonSlideData, ColorMapOverride
- AND type attribute for layout type

#### Scenario: Layout types
- WHEN setting layout type
- THEN title, object, section header, and other types are available

### Requirement: SlideMaster Element

The system SHALL provide `SlideMaster` (p:sldMaster) for master definitions.

#### Scenario: Master element structure
- WHEN creating a SlideMaster
- THEN it contains CommonSlideData, ColorMap, SlideLayoutIdList

#### Scenario: Color map
- WHEN accessing `ColorMap()`
- THEN theme color mappings are available

### Requirement: Theme Element

The system SHALL provide `Theme` (a:theme) for theme definitions.

#### Scenario: Theme element structure
- WHEN creating a Theme
- THEN it contains ThemeElements with color scheme, font scheme, format scheme

#### Scenario: Color scheme
- WHEN accessing color scheme
- THEN dk1, lt1, dk2, lt2, accent1-6, hlink, folHlink colors are available

#### Scenario: Font scheme
- WHEN accessing font scheme
- THEN major and minor font definitions are available

### Requirement: DrawingML Common Elements

The system SHALL provide DrawingML elements for graphics and formatting.

#### Scenario: Transform2D element
- WHEN creating Transform2D (a:xfrm)
- THEN offset and extents define position and size

#### Scenario: Preset geometry
- WHEN setting shape geometry
- THEN all preset shapes (rect, ellipse, etc.) are available

#### Scenario: Fill types
- WHEN setting fill
- THEN solid, gradient, picture, pattern fills are available

#### Scenario: Outline properties
- WHEN setting outline
- THEN width, color, dash style are available

#### Scenario: Effect properties
- WHEN setting effects
- THEN shadow, glow, reflection, blur are available

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
