# DrawingML Document Anchors

Delta specification for document-specific drawing anchors including spreadsheet (xdr:) and word processing (wp:) drawing containers.

## Summary

This spec defines the DrawingML anchor systems that position shapes and images within specific document types. Each document type (Spreadsheet, Word) has its own anchor mechanism tailored to how drawings are positioned relative to document content.

## Motivation

While DrawingML core types are shared, each Office document type positions drawings differently:

- **Spreadsheets**: Drawings are anchored to cells with offsets, allowing them to move/resize with cells
- **Word Documents**: Drawings are either inline (within text flow) or anchored (floating relative to page/paragraph/character)
- **Presentations**: Use presentation-specific positioning (handled in presentation SDK)

These anchor systems wrap the shared DrawingML shapes/pictures while providing document-specific positioning.

## ADDED Requirements

### Spreadsheet Drawing Anchors

### Requirement: Worksheet Drawing Root Element

The system SHALL provide a `WorksheetDrawing` element (xdr:wsDr) as the root of spreadsheet drawing parts.

#### Scenario: WorksheetDrawing structure
- GIVEN a WorksheetDrawing element
- WHEN children are accessed
- THEN TwoCellAnchor (xdr:twoCellAnchor), OneCellAnchor (xdr:oneCellAnchor), AbsoluteAnchor (xdr:absoluteAnchor) elements are available

### Requirement: TwoCellAnchor Element

The system SHALL provide a `TwoCellAnchor` element (xdr:twoCellAnchor) for objects anchored to two cells.

#### Scenario: TwoCellAnchor structure
- GIVEN a TwoCellAnchor element
- WHEN children are accessed
- THEN From (xdr:from), To (xdr:to), and one of Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, plus ClientData (xdr:clientData) are available

#### Scenario: From and To markers
- GIVEN a TwoCellAnchor element
- WHEN From (xdr:from) marker is accessed
- THEN Col (xdr:col), ColOff (xdr:colOff), Row (xdr:row), RowOff (xdr:rowOff) are available
- AND Col/Row are zero-based cell indices
- AND ColOff/RowOff are EMU offsets within the cell

#### Scenario: EditAs attribute
- GIVEN a TwoCellAnchor element
- WHEN editAs attribute is accessed
- THEN values twoCell, oneCell, absolute are available
- AND twoCell: object resizes with cells
- AND oneCell: object moves but does not resize
- AND absolute: object does not move or resize

### Requirement: OneCellAnchor Element

The system SHALL provide a `OneCellAnchor` element (xdr:oneCellAnchor) for objects anchored to one cell.

#### Scenario: OneCellAnchor structure
- GIVEN a OneCellAnchor element
- WHEN children are accessed
- THEN From (xdr:from), Extent (xdr:ext), and one of Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, plus ClientData (xdr:clientData) are available

#### Scenario: Extent element
- GIVEN a OneCellAnchor element
- WHEN Extent (xdr:ext) is accessed
- THEN Cx (width in EMUs) and Cy (height in EMUs) are available

### Requirement: AbsoluteAnchor Element

The system SHALL provide an `AbsoluteAnchor` element (xdr:absoluteAnchor) for objects at absolute positions.

#### Scenario: AbsoluteAnchor structure
- GIVEN an AbsoluteAnchor element
- WHEN children are accessed
- THEN Position (xdr:pos), Extent (xdr:ext), and one of Shape/Picture/ConnectionShape/GraphicFrame/GroupShape, plus ClientData (xdr:clientData) are available

#### Scenario: Position element
- GIVEN an AbsoluteAnchor element
- WHEN Position (xdr:pos) is accessed
- THEN X and Y in EMUs from the worksheet origin are available

### Requirement: Spreadsheet Shape Element

The system SHALL provide a `Shape` element (xdr:sp) for basic shapes in spreadsheets.

#### Scenario: Shape structure
- GIVEN a Shape element
- WHEN properties are accessed
- THEN macro, textlink, fLocksText, fPublished attributes are available
- AND NonVisualShapeProperties (xdr:nvSpPr), ShapeProperties (xdr:spPr), Style (xdr:style), TextBody (xdr:txBody) are available

#### Scenario: NonVisualShapeProperties
- GIVEN a Shape element
- WHEN NonVisualShapeProperties (xdr:nvSpPr) is accessed
- THEN NonVisualDrawingProperties (xdr:cNvPr with id, name, descr, hidden), NonVisualShapeDrawingProperties (xdr:cNvSpPr with txBox, locks) are available

#### Scenario: Shape style reference
- GIVEN a Shape element
- WHEN Style (xdr:style) is accessed
- THEN LineReference (a:lnRef), FillReference (a:fillRef), EffectReference (a:effectRef), FontReference (a:fontRef) from theme are available

### Requirement: Spreadsheet Picture Element

The system SHALL provide a `Picture` element (xdr:pic) for images in spreadsheets.

#### Scenario: Picture structure
- GIVEN a Picture element
- WHEN properties are accessed
- THEN macro, fPublished attributes are available
- AND NonVisualPictureProperties (xdr:nvPicPr), BlipFill (xdr:blipFill), ShapeProperties (xdr:spPr) are available

#### Scenario: Blip fill reference
- GIVEN a Picture element
- WHEN BlipFill (xdr:blipFill) is accessed
- THEN Blip (a:blip) with r:embed or r:link to ImagePart is available
- AND SourceRect (a:srcRect), Stretch (a:stretch), Tile (a:tile) are available

### Requirement: Spreadsheet GraphicFrame Element

The system SHALL provide a `GraphicFrame` element (xdr:graphicFrame) for embedded objects like charts.

#### Scenario: GraphicFrame structure
- GIVEN a GraphicFrame element
- WHEN properties are accessed
- THEN macro, fPublished attributes are available
- AND NonVisualGraphicFrameProperties (xdr:nvGraphicFramePr), Transform (xdr:xfrm), Graphic (a:graphic) are available

#### Scenario: Embedded chart reference
- GIVEN a GraphicFrame with embedded chart
- WHEN Graphic (a:graphic) is accessed
- THEN GraphicData (a:graphicData) with uri "http://schemas.openxmlformats.org/drawingml/2006/chart" is available
- AND c:chart with r:id references the ChartPart

### Requirement: Spreadsheet GroupShape Element

The system SHALL provide a `GroupShape` element (xdr:grpSp) for grouped objects.

#### Scenario: GroupShape structure
- GIVEN a GroupShape element
- WHEN properties are accessed
- THEN NonVisualGroupShapeProperties (xdr:nvGrpSpPr), GroupShapeProperties (xdr:grpSpPr), and child Shape/Picture/ConnectionShape/GraphicFrame/GroupShape elements are available

#### Scenario: Group transform
- GIVEN a GroupShape element
- WHEN GroupShapeProperties transform is accessed
- THEN ChildOffset (a:chOff) and ChildExtents (a:chExt) define the child coordinate space

### Requirement: Spreadsheet ConnectionShape Element

The system SHALL provide a `ConnectionShape` element (xdr:cxnSp) for connectors.

#### Scenario: ConnectionShape structure
- GIVEN a ConnectionShape element
- WHEN properties are accessed
- THEN macro, fPublished attributes are available
- AND NonVisualConnectionShapeProperties (xdr:nvCxnSpPr), ShapeProperties (xdr:spPr), Style (xdr:style) are available

#### Scenario: Connection endpoints
- GIVEN a ConnectionShape element
- WHEN NonVisualConnectionShapeProperties is accessed
- THEN StartConnection (a:stCxn) and EndConnection (a:endCxn) with id, idx reference shapes and connection points

### Requirement: ClientData Element

The system SHALL provide a `ClientData` element (xdr:clientData) for client-specific settings.

#### Scenario: ClientData properties
- GIVEN a ClientData element
- WHEN properties are accessed
- THEN fLocksWithSheet, fPrintsWithSheet boolean attributes are available

### Word Processing Drawing Anchors

### Requirement: Drawing Element

The system SHALL provide a `Drawing` element (w:drawing) as the container for drawings in Word documents.

#### Scenario: Drawing structure
- GIVEN a Drawing element in Word document
- WHEN children are accessed
- THEN Inline (wp:inline) or Anchor (wp:anchor) is available

### Requirement: Inline Element

The system SHALL provide an `Inline` element (wp:inline) for drawings inline with text.

#### Scenario: Inline structure
- GIVEN an Inline element
- WHEN properties are accessed
- THEN distT, distB, distL, distR (distance from text in EMUs) are available
- AND Extent (wp:extent), EffectExtent (wp:effectExtent), DocProperties (wp:docPr), NonVisualGraphicFrameProperties (wp:cNvGraphicFramePr), Graphic (a:graphic) are available

#### Scenario: Inline extent
- GIVEN an Inline element
- WHEN Extent (wp:extent) is accessed
- THEN cx (width) and cy (height) in EMUs are available

#### Scenario: Effect extent
- GIVEN an Inline element
- WHEN EffectExtent (wp:effectExtent) is accessed
- THEN l, t, r, b margins in EMUs for effects extending beyond extent are available

### Requirement: Anchor Element

The system SHALL provide an `Anchor` element (wp:anchor) for floating drawings.

#### Scenario: Anchor structure
- GIVEN an Anchor element
- WHEN properties are accessed
- THEN distT, distB, distL, distR, simplePos, relativeHeight, behindDoc, locked, layoutInCell, hidden, allowOverlap are available
- AND SimplePosition (wp:simplePos), HorizontalPosition (wp:positionH), VerticalPosition (wp:positionV), Extent (wp:extent), EffectExtent (wp:effectExtent), WrapNone/WrapSquare/WrapTight/WrapThrough/WrapTopAndBottom, DocProperties (wp:docPr), NonVisualGraphicFrameProperties (wp:cNvGraphicFramePr), Graphic (a:graphic) are available

#### Scenario: Relative height
- GIVEN an Anchor element
- WHEN relativeHeight attribute is accessed
- THEN z-order value determines stacking order of overlapping anchored objects

#### Scenario: Behind document
- GIVEN an Anchor element
- WHEN behindDoc attribute is true
- THEN the drawing appears behind the document text

### Requirement: Horizontal Position Element

The system SHALL provide a `HorizontalPosition` element (wp:positionH) for horizontal positioning.

#### Scenario: HorizontalPosition structure
- GIVEN a HorizontalPosition element
- WHEN properties are accessed
- THEN relativeFrom attribute with values character, column, insideMargin, leftMargin, margin, outsideMargin, page, rightMargin is available
- AND one of Align (wp:align) or PositionOffset (wp:posOffset) is available

#### Scenario: Horizontal alignment
- GIVEN a HorizontalPosition element
- WHEN Align (wp:align) is accessed
- THEN values center, inside, left, outside, right are available

#### Scenario: Horizontal offset
- GIVEN a HorizontalPosition element
- WHEN PositionOffset (wp:posOffset) is accessed
- THEN offset value in EMUs from the relative reference is available

### Requirement: Vertical Position Element

The system SHALL provide a `VerticalPosition` element (wp:positionV) for vertical positioning.

#### Scenario: VerticalPosition structure
- GIVEN a VerticalPosition element
- WHEN properties are accessed
- THEN relativeFrom attribute with values bottomMargin, insideMargin, line, margin, outsideMargin, page, paragraph, topMargin is available
- AND one of Align (wp:align) or PositionOffset (wp:posOffset) is available

#### Scenario: Vertical alignment
- GIVEN a VerticalPosition element
- WHEN Align (wp:align) is accessed
- THEN values bottom, center, inside, outside, top are available

### Requirement: Text Wrapping Elements

The system SHALL provide text wrapping elements for anchored drawings.

#### Scenario: No wrap
- GIVEN a WrapNone element (wp:wrapNone)
- THEN no text wrapping occurs around the drawing

#### Scenario: Square wrap
- GIVEN a WrapSquare element (wp:wrapSquare)
- WHEN properties are accessed
- THEN wrapText attribute with values bothSides, largest, left, right is available
- AND EffectExtent (wp:effectExtent) is available

#### Scenario: Tight wrap
- GIVEN a WrapTight element (wp:wrapTight)
- WHEN properties are accessed
- THEN wrapText attribute is available
- AND WrapPolygon (wp:wrapPolygon) with Start (wp:start) and LineTo (wp:lineTo) points is available

#### Scenario: Through wrap
- GIVEN a WrapThrough element (wp:wrapThrough)
- WHEN properties are accessed
- THEN same as WrapTight with wrapText and WrapPolygon is available
- AND text can flow through transparent areas

#### Scenario: Top and bottom wrap
- GIVEN a WrapTopAndBottom element (wp:wrapTopAndBottom)
- WHEN properties are accessed
- THEN EffectExtent (wp:effectExtent) is available
- AND text appears only above and below the drawing

#### Scenario: Wrap polygon
- GIVEN a WrapPolygon element
- WHEN points are accessed
- THEN edited attribute indicates if polygon was user-edited
- AND Start (wp:start) with x, y in EMUs defines starting point
- AND LineTo (wp:lineTo) elements define polygon vertices

### Requirement: DocProperties Element

The system SHALL provide a `DocProperties` element (wp:docPr) for drawing properties.

#### Scenario: DocProperties attributes
- GIVEN a DocProperties element
- WHEN attributes are accessed
- THEN id (unique identifier), name (drawing name), descr (description), hidden, title are available

#### Scenario: Hyperlink in DocProperties
- GIVEN a DocProperties element with hyperlink
- WHEN HyperlinkClick (a:hlinkClick) is accessed
- THEN r:id references relationship to target URL

### Requirement: Picture Element for Word

The system SHALL integrate with Picture element (pic:pic) for images in Word.

#### Scenario: Picture in Word drawing
- GIVEN a Graphic element in Word drawing
- WHEN GraphicData contains picture
- THEN uri "http://schemas.openxmlformats.org/drawingml/2006/picture" is used
- AND pic:pic element contains NonVisualPictureProperties (pic:nvPicPr), BlipFill (pic:blipFill), ShapeProperties (pic:spPr)

### Requirement: EMU Positioning Utilities

The system SHALL provide utilities for EMU-based positioning.

#### Scenario: EMU conversion constants
- WHEN EMU values are used
- THEN 1 inch = 914400 EMUs
- AND 1 cm = 360000 EMUs
- AND 1 point = 12700 EMUs
- AND 1 pixel (96 DPI) = 9525 EMUs

#### Scenario: Cell position calculation (spreadsheet)
- GIVEN cell coordinates and column widths/row heights
- WHEN EMU position is calculated
- THEN position accounts for cumulative widths/heights of preceding cells

## Design

### Package Structure

```
drawingml/
  spreadsheet/         # Spreadsheet drawing (xdr: namespace)
    wsdrawing.go       # WorksheetDrawing root element
    twocellanchor.go   # TwoCellAnchor element
    onecellanchor.go   # OneCellAnchor element
    absoluteanchor.go  # AbsoluteAnchor element
    shape.go           # Shape element (xdr:sp)
    picture.go         # Picture element (xdr:pic)
    graphicframe.go    # GraphicFrame element
    groupshape.go      # GroupShape element
    connectionshape.go # ConnectionShape element
    clientdata.go      # ClientData element
    marker.go          # From/To marker types

  word/                # Word drawing (wp: namespace)
    inline.go          # Inline element
    anchor.go          # Anchor element
    position.go        # HorizontalPosition, VerticalPosition
    wrap.go            # WrapNone, WrapSquare, WrapTight, etc.
    docproperties.go   # DocProperties element

  picture/             # Picture types (pic: namespace)
    picture.go         # pic:pic element
    nvpicpr.go         # Non-visual picture properties
```

### Type Naming Convention

| XML Element | Go Type |
|-------------|---------|
| xdr:wsDr | WsDr |
| xdr:twoCellAnchor | TwoCellAnchor |
| xdr:oneCellAnchor | OneCellAnchor |
| xdr:absoluteAnchor | AbsoluteAnchor |
| xdr:sp | Sp |
| xdr:pic | Pic |
| xdr:graphicFrame | GraphicFrame |
| wp:inline | Inline |
| wp:anchor | Anchor |
| wp:positionH | PositionH |
| wp:positionV | PositionV |
| wp:wrapSquare | WrapSquare |
| pic:pic | Pic |

## API

### Spreadsheet Drawing

```go
// Create worksheet drawing
wsDr := spreadsheet.NewWsDr()

// Add image with two-cell anchor
anchor := spreadsheet.NewTwoCellAnchor(
    spreadsheet.WithFrom(0, 0, 0, 0),     // Column 0, Row 0, no offset
    spreadsheet.WithTo(3, 5, 0, 0),       // Column 3, Row 5, no offset
    spreadsheet.WithEditAs("oneCell"),
)

pic := spreadsheet.NewPic(
    spreadsheet.WithBlipEmbed("rId1"),    // relationship to image
)
anchor.SetContent(pic)
wsDr.AddAnchor(anchor)

// Add chart with one-cell anchor
chartAnchor := spreadsheet.NewOneCellAnchor(
    spreadsheet.WithFrom(5, 0, 0, 0),
    spreadsheet.WithExtent(main.InchesToEMU(4), main.InchesToEMU(3)),
)

graphicFrame := spreadsheet.NewGraphicFrame(
    spreadsheet.WithChartReference("rId2"),
)
chartAnchor.SetContent(graphicFrame)
wsDr.AddAnchor(chartAnchor)
```

### Word Drawing - Inline

```go
// Create inline image
inline := word.NewInline(
    word.WithExtent(main.InchesToEMU(2), main.InchesToEMU(1.5)),
    word.WithDocProperties(1, "Image1", "Product photo"),
)

// Set picture content
pic := picture.NewPic(
    picture.WithBlipEmbed("rId1"),
    picture.WithShapeProperties(main.NewSpPr()),
)
inline.SetGraphic(pic)
```

### Word Drawing - Anchor

```go
// Create floating image
anchor := word.NewAnchor(
    word.WithRelativeHeight(1000),
    word.WithBehindDoc(false),
    word.WithExtent(main.InchesToEMU(3), main.InchesToEMU(2)),
    word.WithHorizontalPosition("page", word.PositionOffset(main.InchesToEMU(1))),
    word.WithVerticalPosition("paragraph", word.Align("top")),
    word.WithWrapSquare("bothSides"),
)

pic := picture.NewPic(
    picture.WithBlipEmbed("rId1"),
)
anchor.SetGraphic(pic)
```

### Text Wrapping

```go
// Tight wrap with custom polygon
anchor := word.NewAnchor(
    word.WithWrapTight("bothSides",
        word.WrapPolygon(
            word.Point(0, 0),
            word.Point(21600, 0),
            word.Point(21600, 21600),
            word.Point(0, 21600),
            word.Point(0, 0),
        ),
    ),
)

// Top and bottom wrap
anchor := word.NewAnchor(
    word.WithWrapTopAndBottom(),
)
```

## Testing

### Unit Tests

- WorksheetDrawing creation and anchor management
- TwoCellAnchor with From/To markers
- OneCellAnchor with extent
- AbsoluteAnchor with position
- All shape types in spreadsheet drawings
- Inline element with extent
- Anchor element with positioning
- All wrap types with configurations
- Position calculations from cell references

### Integration Tests

- Complete worksheet drawing round-trip
- Image embedding with relationships
- Chart embedding with relationships
- Word document drawing round-trip
- Text wrapping behavior verification

### Validation Tests

- Invalid cell indices
- Missing required elements
- Invalid wrap text values
- Invalid position relative values

## Dependencies

- `encoding/xml` - Standard library XML support
- `drawingml/main` - Core DrawingML types
- `drawingml/picture` - Picture element types
- No external dependencies

## Migration

N/A - New capability with no existing implementation.
