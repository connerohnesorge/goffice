## ADDED Requirements

### Requirement: WorksheetDrawing Root Element

The system SHALL provide a `WorksheetDrawing` element (xdr:wsDr) as the root of DrawingsPart.

#### Scenario: Access drawing anchors
- GIVEN a WorksheetDrawing element
- WHEN children are accessed
- THEN TwoCellAnchor, OneCellAnchor, AbsoluteAnchor elements are available

### Requirement: TwoCellAnchor Element

The system SHALL provide a `TwoCellAnchor` element for objects anchored to two cells.

#### Scenario: TwoCellAnchor properties
- GIVEN a TwoCellAnchor element
- WHEN properties are accessed
- THEN EditAs (twoCell/oneCell/absolute), From, To, Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, ClientData are available

#### Scenario: From and To markers
- GIVEN a TwoCellAnchor element
- WHEN From and To markers are accessed
- THEN Col, ColOff, Row, RowOff are available
- AND Col/Row are zero-based cell indices
- AND ColOff/RowOff are EMU offsets within the cell

#### Scenario: EditAs behavior
- GIVEN a TwoCellAnchor with different EditAs values
- WHEN cells are resized
- THEN twoCell: object resizes with cells
- AND oneCell: object moves but doesn't resize
- AND absolute: object doesn't move or resize

### Requirement: OneCellAnchor Element

The system SHALL provide a `OneCellAnchor` element for objects anchored to one cell.

#### Scenario: OneCellAnchor properties
- GIVEN a OneCellAnchor element
- WHEN properties are accessed
- THEN From, Extent, Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, ClientData are available

#### Scenario: Extent size
- GIVEN a OneCellAnchor element
- WHEN Extent (xdr:ext) is accessed
- THEN Cx (width in EMUs) and Cy (height in EMUs) are available

### Requirement: AbsoluteAnchor Element

The system SHALL provide an `AbsoluteAnchor` element for objects at absolute positions.

#### Scenario: AbsoluteAnchor properties
- GIVEN an AbsoluteAnchor element
- WHEN properties are accessed
- THEN Position (xdr:pos), Extent, Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, ClientData are available

#### Scenario: Absolute position
- GIVEN an AbsoluteAnchor element
- WHEN Position is accessed
- THEN X and Y in EMUs from the worksheet origin are available

### Requirement: Shape Element

The system SHALL provide a `Shape` element (xdr:sp) for basic shapes.

#### Scenario: Shape properties
- GIVEN a Shape element
- WHEN properties are accessed
- THEN Macro, TextLink, FLocksText, FPublished, NonVisualProperties, ShapeProperties, Style, TextBody are available

#### Scenario: NonVisualShapeProperties
- GIVEN a Shape element
- WHEN NonVisualProperties (xdr:nvSpPr) is accessed
- THEN cNvPr (id, name, description, hidden), cNvSpPr (text box, locks) are available

#### Scenario: Shape style
- GIVEN a Shape element
- WHEN Style (xdr:style) is accessed
- THEN LineReference, FillReference, EffectReference, FontReference from theme are available

### Requirement: Picture Element

The system SHALL provide a `Picture` element (xdr:pic) for images.

#### Scenario: Picture properties
- GIVEN a Picture element
- WHEN properties are accessed
- THEN Macro, FPublished, NonVisualPictureProperties, BlipFill, ShapeProperties are available

#### Scenario: Blip fill
- GIVEN a Picture element
- WHEN BlipFill (xdr:blipFill) is accessed
- THEN Blip (a:blip with r:embed to ImagePart), SourceRect, Tile/Stretch are available

#### Scenario: Image relationship
- GIVEN a Picture with embedded image
- WHEN Blip r:embed is resolved
- THEN the relationship points to an ImagePart containing the image data

### Requirement: GraphicFrame Element

The system SHALL provide a `GraphicFrame` element (xdr:graphicFrame) for embedded objects like charts.

#### Scenario: GraphicFrame properties
- GIVEN a GraphicFrame element
- WHEN properties are accessed
- THEN Macro, FPublished, NonVisualGraphicFrameProperties, Transform, Graphic are available

#### Scenario: Embedded chart
- GIVEN a GraphicFrame with embedded chart
- WHEN Graphic (a:graphic) is accessed
- THEN GraphicData with uri "http://schemas.openxmlformats.org/drawingml/2006/chart" contains c:chart with r:id to ChartPart

### Requirement: GroupShape Element

The system SHALL provide a `GroupShape` element (xdr:grpSp) for grouped objects.

#### Scenario: GroupShape properties
- GIVEN a GroupShape element
- WHEN properties are accessed
- THEN NonVisualGroupShapeProperties, GroupShapeProperties, Shape/Picture/ConnectionShape/GraphicFrame/GroupShape children are available

#### Scenario: Group transform
- GIVEN a GroupShape element
- WHEN GroupShapeProperties.Transform is accessed
- THEN ChildOffset, ChildExtents define the child coordinate space

### Requirement: ConnectionShape Element

The system SHALL provide a `ConnectionShape` element (xdr:cxnSp) for connectors.

#### Scenario: ConnectionShape properties
- GIVEN a ConnectionShape element
- WHEN properties are accessed
- THEN Macro, FPublished, NonVisualConnectionShapeProperties, ShapeProperties, Style are available

#### Scenario: Connection points
- GIVEN a ConnectionShape element
- WHEN NonVisualConnectionShapeProperties is accessed
- THEN StartConnection and EndConnection (id, idx) reference shapes and connection points

### Requirement: ClientData Element

The system SHALL provide a `ClientData` element for client-specific settings.

#### Scenario: ClientData properties
- GIVEN a ClientData element
- WHEN properties are accessed
- THEN FLocksWithSheet, FPrintsWithSheet are available

### Requirement: Drawing Relationship to Worksheet

The system SHALL establish proper relationships between worksheet and drawings.

#### Scenario: Worksheet drawing reference
- GIVEN a Worksheet with drawings
- WHEN the Drawing element in Worksheet is accessed
- THEN r:id references the DrawingsPart

#### Scenario: Add drawing to worksheet
- GIVEN a WorksheetPart without drawings
- WHEN `worksheet.AddDrawingsPart()` is called
- THEN a DrawingsPart is created
- AND a Drawing element with r:id is added to the Worksheet

### Requirement: Image Part Integration

The system SHALL support embedding images in worksheets.

#### Scenario: Add image to drawings
- GIVEN a DrawingsPart
- WHEN an image is added
- THEN an ImagePart is created with the image data
- AND a Picture element is added to WorksheetDrawing
- AND a TwoCellAnchor or OneCellAnchor contains the Picture

#### Scenario: Supported image formats
- GIVEN an image file
- WHEN the content type is determined
- THEN PNG, JPEG, GIF, BMP, TIFF, WMF, EMF are supported

### Requirement: EMU (English Metric Units)

The system SHALL use EMUs for precise positioning.

#### Scenario: EMU conversion
- GIVEN dimensions in various units
- WHEN converted to EMUs
- THEN 1 inch = 914400 EMUs
- AND 1 cm = 360000 EMUs
- AND 1 point = 12700 EMUs
- AND 1 pixel (96 dpi) = 9525 EMUs

#### Scenario: Position from cell
- GIVEN a cell reference and offset
- WHEN converted to EMUs
- THEN the position accounts for column widths and row heights

### Requirement: Drawing API

The system SHALL provide a high-level API for adding drawings.

#### Scenario: Add image at cell
- GIVEN a worksheet
- WHEN `sheet.AddImage("A1", imageBytes, "image/png")` is called
- THEN an ImagePart is created
- AND a OneCellAnchor with Picture is added at cell A1

#### Scenario: Add image spanning cells
- GIVEN a worksheet
- WHEN `sheet.AddImageBetween("A1", "C5", imageBytes, "image/png")` is called
- THEN a TwoCellAnchor with Picture is added
- AND From is A1 and To is C5

#### Scenario: Add chart at position
- GIVEN a worksheet with data
- WHEN `sheet.AddChart(ChartTypeColumn, "A1:B10", "D1", "I15")` is called
- THEN a ChartPart is created
- AND a TwoCellAnchor with GraphicFrame is added
- AND the chart references the data range
