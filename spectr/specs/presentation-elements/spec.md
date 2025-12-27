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

