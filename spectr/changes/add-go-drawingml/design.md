# Design: Go DrawingML Shared Package

## Context

DrawingML is the shared graphics layer in Office Open XML documents, providing common types for shapes, images, effects, and charts across Word, PowerPoint, and Excel. This design document captures architectural decisions for implementing a Go package that mirrors the DrawingML portions of Microsoft's Open-XML-SDK.

**Stakeholders:**
- Word SDK (goffice/wordprocessing) - uses DrawingML for images and shapes
- Presentation SDK (goffice/presentation) - uses DrawingML extensively for all visual elements
- Spreadsheet SDK (goffice/spreadsheet) - uses DrawingML for charts and drawings

**Constraints:**
- Must support ECMA-376 5th edition (Office 2016+)
- Must integrate seamlessly with openxml/ framework package
- Must work without external dependencies
- Types must be usable by all three document SDKs

**Decided Design Choices:**
- Code generation from JSON schemas (same source as C# SDK) - pre-generated, committed to repo
- Idiomatic Go naming conventions
- Strict validation by default
- sync.RWMutex for thread safety (concurrent reads, exclusive writes)
- Functional options pattern for element construction
- Generic methods only for child access: `First[T]()`, `All[T]()`, `OfType[T]()`

## Goals / Non-Goals

### Goals
- Provide all DrawingML types needed by Word, Presentation, and Spreadsheet SDKs
- Maintain feature parity with C# Open-XML-SDK DrawingML types
- Idiomatic Go API design
- Strong typing for all elements and attributes
- Comprehensive validation
- Thread-safe read operations
- EMU conversion utilities

### Non-Goals
- Rendering or display of DrawingML content
- Image format conversion
- Chart data calculation/evaluation
- SVG import/export
- 3D scene rendering

## Architecture Overview

```
goffice/
├── drawingml/
│   ├── main/           # Core DrawingML types (a: namespace)
│   │   ├── shapes.go       # Preset and custom geometries
│   │   ├── fills.go        # Solid, gradient, pattern, blip fills
│   │   ├── effects.go      # Shadows, glows, reflections
│   │   ├── text.go         # TextBody, paragraphs, runs
│   │   ├── transforms.go   # Transform2D, offsets, extents
│   │   ├── lines.go        # Line properties, styles
│   │   ├── colors.go       # RGB, scheme, system, HSL colors
│   │   ├── blips.go        # Image references
│   │   └── enums.go        # All DrawingML enumerations
│   │
│   ├── chart/          # Chart types (c: namespace)
│   │   ├── chartspace.go   # ChartSpace root element
│   │   ├── chart.go        # Chart definition
│   │   ├── plotarea.go     # Plot area and chart types
│   │   ├── series.go       # Data series
│   │   ├── axis.go         # Category, value, date axes
│   │   ├── legend.go       # Legend formatting
│   │   └── enums.go        # Chart enumerations
│   │
│   ├── spreadsheet/    # Spreadsheet drawing (xdr: namespace)
│   │   ├── wsdrawing.go    # WorksheetDrawing root
│   │   ├── anchors.go      # TwoCellAnchor, OneCellAnchor, AbsoluteAnchor
│   │   ├── shapes.go       # Shape, Picture, GraphicFrame
│   │   └── markers.go      # Cell markers (From, To)
│   │
│   ├── word/           # Word drawing (wp: namespace)
│   │   ├── inline.go       # Inline positioning
│   │   ├── anchor.go       # Floating anchor positioning
│   │   ├── wrapping.go     # Text wrapping modes
│   │   └── position.go     # Position controls
│   │
│   ├── picture/        # Picture types (pic: namespace)
│   │   ├── picture.go      # pic:pic element
│   │   ├── nvprops.go      # Non-visual properties
│   │   └── blipfill.go     # Picture blip fill
│   │
│   └── units/          # Unit conversion utilities
│       └── emu.go          # EMU <-> inches, cm, points, pixels
│
├── openxml/            # Core framework (used by DrawingML)
│   ├── element.go
│   ├── composite.go
│   └── ...
│
└── packaging/          # OPC layer
    └── ...
```

## Decisions

### D1: Package Structure

**Decision:** Organize DrawingML types into subpackages by XML namespace.

```
drawingml/
├── main/           # a: namespace - core DrawingML
├── chart/          # c: namespace - charts
├── spreadsheet/    # xdr: namespace - spreadsheet drawings
├── word/           # wp: namespace - word drawings
├── picture/        # pic: namespace - pictures
└── units/          # EMU conversion utilities
```

**Rationale:**
- Clear namespace separation matching XML structure
- Import paths reflect usage: `import "goffice/drawingml/main"`
- Avoids name collisions between namespaces
- Easy to identify which namespace a type belongs to

### D2: Element Implementation

**Decision:** All DrawingML elements embed openxml framework types and are generated from JSON schemas.

```go
// Example: Transform2D (a:xfrm)
type Xfrm struct {
    openxml.CompositeElement

    // Attributes
    Rot   *types.Int64Value  `xml:"rot,attr,omitempty"`   // Rotation in 60000ths of a degree
    FlipH *types.BoolValue   `xml:"flipH,attr,omitempty"` // Horizontal flip
    FlipV *types.BoolValue   `xml:"flipV,attr,omitempty"` // Vertical flip

    // Child elements
    Off *Offset  `xml:"off,omitempty"` // Position offset (a:off)
    Ext *Extent  `xml:"ext,omitempty"` // Size extents (a:ext)
}

// Offset represents a:off
type Offset struct {
    openxml.LeafElement
    X int64 `xml:"x,attr"` // X coordinate in EMUs
    Y int64 `xml:"y,attr"` // Y coordinate in EMUs
}

// Extent represents a:ext
type Extent struct {
    openxml.LeafElement
    Cx int64 `xml:"cx,attr"` // Width in EMUs
    Cy int64 `xml:"cy,attr"` // Height in EMUs
}
```

**Rationale:** Consistent with Word and Presentation SDKs; embeddes framework types provide common functionality.

### D3: Fill Type Hierarchy

**Decision:** Implement fill types as concrete structs implementing a common Fill interface.

```go
// Fill is implemented by all fill types
type Fill interface {
    openxml.Element
    isFill() // Marker method
}

// SolidFill represents a:solidFill
type SolidFill struct {
    openxml.CompositeElement
    Color Color // Color can be sRGB, scheme, system, etc.
}

func (*SolidFill) isFill() {}

// GradFill represents a:gradFill
type GradFill struct {
    openxml.CompositeElement
    RotWithShape *types.BoolValue   `xml:"rotWithShape,attr,omitempty"`
    Flip         *FlipModeValue     `xml:"flip,attr,omitempty"`
    GsLst        *GradientStopList  `xml:"gsLst,omitempty"`
    Lin          *LinearShade       `xml:"lin,omitempty"`
    Path         *PathShade         `xml:"path,omitempty"`
}

func (*GradFill) isFill() {}

// Additional fill types: PatternFill, BlipFill, GroupFill, NoFill
```

**Rationale:** Go interfaces provide polymorphism for fill types in shape properties.

### D4: Color System

**Decision:** Implement color types as concrete structs with a common Color interface.

```go
// Color is implemented by all color types
type Color interface {
    openxml.Element
    isColor()
    // ToRGBA returns the resolved RGBA color value (may require theme lookup)
    ToRGBA(theme Theme) color.RGBA
}

// SrgbClr represents a:srgbClr (RGB hex color)
type SrgbClr struct {
    openxml.CompositeElement
    Val string `xml:"val,attr"` // 6-digit hex: "FF0000"
    // Transform children
    Alpha  *Alpha  `xml:"alpha,omitempty"`
    Tint   *Tint   `xml:"tint,omitempty"`
    Shade  *Shade  `xml:"shade,omitempty"`
}

func (*SrgbClr) isColor() {}

// SchemeClr represents a:schemeClr (theme color reference)
type SchemeClr struct {
    openxml.CompositeElement
    Val SchemeColorValue `xml:"val,attr"` // dk1, lt1, accent1, etc.
    // Transform children inherited from base
}

func (*SchemeClr) isColor() {}

// Scheme color values
type SchemeColorValue string

const (
    SchemeColorDk1      SchemeColorValue = "dk1"      // Dark 1 (text)
    SchemeColorLt1      SchemeColorValue = "lt1"      // Light 1 (background)
    SchemeColorDk2      SchemeColorValue = "dk2"      // Dark 2
    SchemeColorLt2      SchemeColorValue = "lt2"      // Light 2
    SchemeColorAccent1  SchemeColorValue = "accent1"  // Accent 1
    SchemeColorAccent2  SchemeColorValue = "accent2"  // Accent 2
    SchemeColorAccent3  SchemeColorValue = "accent3"  // Accent 3
    SchemeColorAccent4  SchemeColorValue = "accent4"  // Accent 4
    SchemeColorAccent5  SchemeColorValue = "accent5"  // Accent 5
    SchemeColorAccent6  SchemeColorValue = "accent6"  // Accent 6
    SchemeColorHlink    SchemeColorValue = "hlink"    // Hyperlink
    SchemeColorFolHlink SchemeColorValue = "folHlink" // Followed hyperlink
)
```

**Rationale:** Strongly typed colors with transforms; scheme colors require theme context for resolution.

### D5: Preset Geometry Registry

**Decision:** Define preset geometry names as typed constants with validation.

```go
// PresetGeomType represents preset geometry names (prstGeom/@prst)
type PresetGeomType string

const (
    // Basic shapes
    PresetGeomRect           PresetGeomType = "rect"
    PresetGeomEllipse        PresetGeomType = "ellipse"
    PresetGeomRoundRect      PresetGeomType = "roundRect"
    PresetGeomTriangle       PresetGeomType = "triangle"
    PresetGeomRtTriangle     PresetGeomType = "rtTriangle"
    PresetGeomParallelogram  PresetGeomType = "parallelogram"
    PresetGeomTrapezoid      PresetGeomType = "trapezoid"
    PresetGeomDiamond        PresetGeomType = "diamond"
    PresetGeomPentagon       PresetGeomType = "pentagon"
    PresetGeomHexagon        PresetGeomType = "hexagon"
    PresetGeomHeptagon       PresetGeomType = "heptagon"
    PresetGeomOctagon        PresetGeomType = "octagon"
    PresetGeomDecagon        PresetGeomType = "decagon"
    PresetGeomDodecagon      PresetGeomType = "dodecagon"

    // Arrows
    PresetGeomRightArrow     PresetGeomType = "rightArrow"
    PresetGeomLeftArrow      PresetGeomType = "leftArrow"
    PresetGeomUpArrow        PresetGeomType = "upArrow"
    PresetGeomDownArrow      PresetGeomType = "downArrow"
    PresetGeomLeftRightArrow PresetGeomType = "leftRightArrow"
    PresetGeomUpDownArrow    PresetGeomType = "upDownArrow"

    // Stars
    PresetGeomStar4          PresetGeomType = "star4"
    PresetGeomStar5          PresetGeomType = "star5"
    PresetGeomStar6          PresetGeomType = "star6"
    PresetGeomStar8          PresetGeomType = "star8"
    PresetGeomStar10         PresetGeomType = "star10"
    PresetGeomStar12         PresetGeomType = "star12"
    PresetGeomStar16         PresetGeomType = "star16"
    PresetGeomStar24         PresetGeomType = "star24"
    PresetGeomStar32         PresetGeomType = "star32"

    // Flowchart
    PresetGeomFlowChartProcess    PresetGeomType = "flowChartProcess"
    PresetGeomFlowChartDecision   PresetGeomType = "flowChartDecision"
    PresetGeomFlowChartTerminator PresetGeomType = "flowChartTerminator"
    // ... 100+ more presets
)

// IsValid checks if the preset geometry type is valid
func (p PresetGeomType) IsValid() bool {
    // Generated validation logic
}

// PrstGeom represents a:prstGeom (preset geometry)
type PrstGeom struct {
    openxml.CompositeElement
    Prst  PresetGeomType `xml:"prst,attr"` // Preset type
    AvLst *AdjustValueList `xml:"avLst,omitempty"` // Adjust values
}
```

**Rationale:** Type-safe preset names; validation catches invalid values early.

### D6: Text Body Structure

**Decision:** Model text body as hierarchical elements matching XML structure.

```go
// TxBody represents a:txBody or p:txBody (text body container)
type TxBody struct {
    openxml.CompositeElement
    BodyPr   *BodyPr   `xml:"bodyPr,omitempty"`   // Body properties
    LstStyle *LstStyle `xml:"lstStyle,omitempty"` // List style
    P        []*Para   `xml:"p"`                  // Paragraphs
}

// BodyPr represents a:bodyPr (body properties)
type BodyPr struct {
    openxml.LeafElement
    // Attributes
    Wrap       *TextWrappingType `xml:"wrap,attr,omitempty"`
    Anchor     *TextAnchorType   `xml:"anchor,attr,omitempty"`
    AnchorCtr  *types.BoolValue  `xml:"anchorCtr,attr,omitempty"`
    Rot        *types.Int64Value `xml:"rot,attr,omitempty"`
    Vert       *TextVerticalType `xml:"vert,attr,omitempty"`
    LIns       *types.Int64Value `xml:"lIns,attr,omitempty"` // Left inset EMU
    TIns       *types.Int64Value `xml:"tIns,attr,omitempty"` // Top inset EMU
    RIns       *types.Int64Value `xml:"rIns,attr,omitempty"` // Right inset EMU
    BIns       *types.Int64Value `xml:"bIns,attr,omitempty"` // Bottom inset EMU
    NumCol     *types.Int32Value `xml:"numCol,attr,omitempty"`
    SpcCol     *types.Int64Value `xml:"spcCol,attr,omitempty"`
    RTLCol     *types.BoolValue  `xml:"rtlCol,attr,omitempty"`
}

// Para represents a:p (paragraph)
type Para struct {
    openxml.CompositeElement
    PPr      *ParaPr    `xml:"pPr,omitempty"` // Paragraph properties
    Children []ParaChild // Runs, breaks, fields (a:r, a:br, a:fld)
}

// ParaChild is implemented by paragraph child elements
type ParaChild interface {
    openxml.Element
    isParaChild()
}

// Run represents a:r (text run)
type Run struct {
    openxml.CompositeElement
    RPr *RunPr `xml:"rPr,omitempty"` // Run properties
    T   *Text  `xml:"t,omitempty"`   // Text content
}

func (*Run) isParaChild() {}

// RunPr represents a:rPr (run properties)
type RunPr struct {
    openxml.CompositeElement
    Lang     *types.StringValue `xml:"lang,attr,omitempty"`
    Sz       *types.Int32Value  `xml:"sz,attr,omitempty"` // Font size in hundredths of point
    B        *types.BoolValue   `xml:"b,attr,omitempty"`  // Bold
    I        *types.BoolValue   `xml:"i,attr,omitempty"`  // Italic
    U        *TextUnderlineType `xml:"u,attr,omitempty"`  // Underline
    Strike   *TextStrikeType    `xml:"strike,attr,omitempty"`
    SolidFill *SolidFill        `xml:"solidFill,omitempty"` // Text color
    // ... more properties
}

// Text represents a:t (text content)
type Text struct {
    openxml.LeafTextElement
}
```

**Rationale:** Matches XML structure; enables full text formatting capabilities.

### D7: EMU Conversion Utilities

**Decision:** Provide type-safe EMU wrapper with conversion methods.

```go
// Package units provides EMU (English Metric Unit) utilities

// EMU represents a value in English Metric Units
type EMU int64

// Conversion constants
const (
    EMUsPerInch  EMU = 914400
    EMUsPerCm    EMU = 360000
    EMUsPerPoint EMU = 12700
    EMUsPerPixel EMU = 9525 // Assuming 96 DPI
)

// Inches creates EMU from inches
func Inches(in float64) EMU {
    return EMU(in * float64(EMUsPerInch))
}

// Cm creates EMU from centimeters
func Cm(cm float64) EMU {
    return EMU(cm * float64(EMUsPerCm))
}

// Points creates EMU from points
func Points(pt float64) EMU {
    return EMU(pt * float64(EMUsPerPoint))
}

// Pixels creates EMU from pixels (96 DPI)
func Pixels(px int) EMU {
    return EMU(px) * EMUsPerPixel
}

// ToInches converts EMU to inches
func (e EMU) ToInches() float64 {
    return float64(e) / float64(EMUsPerInch)
}

// ToCm converts EMU to centimeters
func (e EMU) ToCm() float64 {
    return float64(e) / float64(EMUsPerCm)
}

// ToPoints converts EMU to points
func (e EMU) ToPoints() float64 {
    return float64(e) / float64(EMUsPerPoint)
}

// ToPixels converts EMU to pixels (96 DPI)
func (e EMU) ToPixels() int {
    return int(e / EMUsPerPixel)
}

// Rotation angles
const DegreesPerRotationUnit = 60000

// Degrees creates rotation units from degrees
func Degrees(deg float64) int64 {
    return int64(deg * DegreesPerRotationUnit)
}

// ToDegrees converts rotation units to degrees
func ToDegrees(rot int64) float64 {
    return float64(rot) / DegreesPerRotationUnit
}
```

**Rationale:** Type safety prevents unit confusion; methods are ergonomic.

### D8: Chart Architecture

**Decision:** Model charts with hierarchical structure matching C:namespace.

```go
// ChartSpace represents c:chartSpace (root chart element)
type ChartSpace struct {
    openxml.CompositeElement
    Date1904     *types.BoolValue  `xml:"date1904,omitempty"`
    Lang         *types.StringValue `xml:"lang,omitempty"`
    RoundedCorners *types.BoolValue `xml:"roundedCorners,omitempty"`
    Style        *types.Int32Value `xml:"style,omitempty"`
    Chart        *Chart            `xml:"chart"`
    // ... more properties
}

// Chart represents c:chart
type Chart struct {
    openxml.CompositeElement
    Title       *Title      `xml:"title,omitempty"`
    AutoTitleDeleted *types.BoolValue `xml:"autoTitleDeleted,omitempty"`
    PlotArea    *PlotArea   `xml:"plotArea"`
    Legend      *Legend     `xml:"legend,omitempty"`
    PlotVisOnly *types.BoolValue `xml:"plotVisOnly,omitempty"`
    DispBlanksAs *DispBlanksAsValue `xml:"dispBlanksAs,omitempty"`
}

// PlotArea represents c:plotArea
type PlotArea struct {
    openxml.CompositeElement
    Layout *Layout `xml:"layout,omitempty"`
    // Chart types (choice - only one per plot area for simple charts)
    BarChart      *BarChart      `xml:"barChart,omitempty"`
    Bar3DChart    *Bar3DChart    `xml:"bar3DChart,omitempty"`
    LineChart     *LineChart     `xml:"lineChart,omitempty"`
    Line3DChart   *Line3DChart   `xml:"line3DChart,omitempty"`
    PieChart      *PieChart      `xml:"pieChart,omitempty"`
    Pie3DChart    *Pie3DChart    `xml:"pie3DChart,omitempty"`
    DoughnutChart *DoughnutChart `xml:"doughnutChart,omitempty"`
    AreaChart     *AreaChart     `xml:"areaChart,omitempty"`
    Area3DChart   *Area3DChart   `xml:"area3DChart,omitempty"`
    ScatterChart  *ScatterChart  `xml:"scatterChart,omitempty"`
    BubbleChart   *BubbleChart   `xml:"bubbleChart,omitempty"`
    RadarChart    *RadarChart    `xml:"radarChart,omitempty"`
    SurfaceChart  *SurfaceChart  `xml:"surfaceChart,omitempty"`
    StockChart    *StockChart    `xml:"stockChart,omitempty"`
    // Axes
    CatAx   []*CatAx   `xml:"catAx,omitempty"`   // Category axes
    ValAx   []*ValAx   `xml:"valAx,omitempty"`   // Value axes
    DateAx  []*DateAx  `xml:"dateAx,omitempty"`  // Date axes
    SerAx   []*SerAx   `xml:"serAx,omitempty"`   // Series axes
}

// BarChart represents c:barChart
type BarChart struct {
    openxml.CompositeElement
    BarDir    *BarDirValue    `xml:"barDir"`    // bar or col
    Grouping  *BarGroupingValue `xml:"grouping,omitempty"`
    VaryColors *types.BoolValue `xml:"varyColors,omitempty"`
    Ser       []*BarChartSeries `xml:"ser,omitempty"` // Data series
    GapWidth  *types.Int32Value `xml:"gapWidth,omitempty"`
    Overlap   *types.Int32Value `xml:"overlap,omitempty"`
    AxId      []uint32          `xml:"axId"` // Axis IDs
}
```

**Rationale:** Full chart model enables reading and writing any chart type.

### D9: Spreadsheet Drawing Anchors

**Decision:** Model anchor types for cell-relative positioning.

```go
// WsDr represents xdr:wsDr (worksheet drawing root)
type WsDr struct {
    openxml.CompositeElement
    TwoCellAnchors   []*TwoCellAnchor   `xml:"twoCellAnchor,omitempty"`
    OneCellAnchors   []*OneCellAnchor   `xml:"oneCellAnchor,omitempty"`
    AbsoluteAnchors  []*AbsoluteAnchor  `xml:"absoluteAnchor,omitempty"`
}

// TwoCellAnchor represents xdr:twoCellAnchor
type TwoCellAnchor struct {
    openxml.CompositeElement
    EditAs     *EditAsType `xml:"editAs,attr,omitempty"` // twoCell, oneCell, absolute
    From       *Marker     `xml:"from"`
    To         *Marker     `xml:"to"`
    // Shape content (choice)
    Sp         *Sp         `xml:"sp,omitempty"`
    Pic        *Pic        `xml:"pic,omitempty"`
    GraphicFrame *GraphicFrame `xml:"graphicFrame,omitempty"`
    GrpSp      *GrpSp      `xml:"grpSp,omitempty"`
    CxnSp      *CxnSp      `xml:"cxnSp,omitempty"`
    ClientData *ClientData `xml:"clientData,omitempty"`
}

// Marker represents xdr:from or xdr:to (cell marker)
type Marker struct {
    openxml.CompositeElement
    Col    int   `xml:"col"`    // Column (0-based)
    ColOff int64 `xml:"colOff"` // Column offset in EMUs
    Row    int   `xml:"row"`    // Row (0-based)
    RowOff int64 `xml:"rowOff"` // Row offset in EMUs
}

// OneCellAnchor represents xdr:oneCellAnchor
type OneCellAnchor struct {
    openxml.CompositeElement
    From       *Marker   `xml:"from"`
    Ext        *Ext      `xml:"ext"` // Size in EMUs
    // Shape content (same choices as TwoCellAnchor)
    ClientData *ClientData `xml:"clientData,omitempty"`
}
```

**Rationale:** Accurate cell-based positioning for Excel drawings.

### D10: Word Drawing Positioning

**Decision:** Model inline and anchor positioning for Word documents.

```go
// Inline represents wp:inline (inline with text positioning)
type Inline struct {
    openxml.CompositeElement
    DistT   *types.Int64Value `xml:"distT,attr,omitempty"` // Distance from text (EMU)
    DistB   *types.Int64Value `xml:"distB,attr,omitempty"`
    DistL   *types.Int64Value `xml:"distL,attr,omitempty"`
    DistR   *types.Int64Value `xml:"distR,attr,omitempty"`
    Extent  *Extent           `xml:"extent"`
    EffectExtent *EffectExtent `xml:"effectExtent,omitempty"`
    DocPr   *DocPr            `xml:"docPr"`
    CNvGraphicFramePr *CNvGraphicFramePr `xml:"cNvGraphicFramePr,omitempty"`
    Graphic *Graphic          `xml:"graphic"`
}

// Anchor represents wp:anchor (floating positioning)
type Anchor struct {
    openxml.CompositeElement
    // Attributes
    DistT          *types.Int64Value `xml:"distT,attr,omitempty"`
    DistB          *types.Int64Value `xml:"distB,attr,omitempty"`
    DistL          *types.Int64Value `xml:"distL,attr,omitempty"`
    DistR          *types.Int64Value `xml:"distR,attr,omitempty"`
    SimplePos      *types.BoolValue  `xml:"simplePos,attr,omitempty"`
    RelativeHeight uint32            `xml:"relativeHeight,attr"`
    BehindDoc      *types.BoolValue  `xml:"behindDoc,attr,omitempty"`
    Locked         *types.BoolValue  `xml:"locked,attr,omitempty"`
    LayoutInCell   *types.BoolValue  `xml:"layoutInCell,attr,omitempty"`
    Hidden         *types.BoolValue  `xml:"hidden,attr,omitempty"`
    AllowOverlap   *types.BoolValue  `xml:"allowOverlap,attr,omitempty"`
    // Child elements
    SimplePosXY    *SimplePosXY      `xml:"simplePos,omitempty"`
    PositionH      *PosH             `xml:"positionH"`
    PositionV      *PosV             `xml:"positionV"`
    Extent         *Extent           `xml:"extent"`
    EffectExtent   *EffectExtent     `xml:"effectExtent,omitempty"`
    WrapNone       *WrapNone         `xml:"wrapNone,omitempty"`
    WrapSquare     *WrapSquare       `xml:"wrapSquare,omitempty"`
    WrapTight      *WrapTight        `xml:"wrapTight,omitempty"`
    WrapThrough    *WrapThrough      `xml:"wrapThrough,omitempty"`
    WrapTopAndBottom *WrapTopAndBottom `xml:"wrapTopAndBottom,omitempty"`
    DocPr          *DocPr            `xml:"docPr"`
    Graphic        *Graphic          `xml:"graphic"`
}

// Wrapping types
type WrapSquare struct {
    openxml.LeafElement
    WrapText WrapTextType `xml:"wrapText,attr"` // bothSides, left, right, largest
}

type WrapTight struct {
    openxml.CompositeElement
    WrapText    WrapTextType `xml:"wrapText,attr"`
    WrapPolygon *WrapPolygon `xml:"wrapPolygon,omitempty"`
}
```

**Rationale:** Complete Word drawing positioning with all wrapping modes.

### D11: Thread Safety

**Decision:** Use sync.RWMutex for document-level locking.

```go
// Thread-safe access follows same pattern as Word/Presentation SDKs
// Read operations are concurrent, write operations are exclusive
// Locking is at the document level, not individual elements

// Example usage in parent SDK:
type PresentationDocument struct {
    pkg     *packaging.Package
    mu      sync.RWMutex
    // ...
}

func (d *PresentationDocument) GetSlide(idx int) (*presentation.Slide, error) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    // Safe concurrent read
    return d.slides[idx], nil
}

func (d *PresentationDocument) AddSlide(slide *presentation.Slide) error {
    d.mu.Lock()
    defer d.mu.Unlock()
    // Exclusive write
    d.slides = append(d.slides, slide)
    return nil
}
```

**Rationale:** Consistent with Word and Presentation SDK design.

### D12: Validation Integration

**Decision:** DrawingML elements integrate with openxml/validation framework.

```go
// Each DrawingML element has embedded validation metadata
// Generated during code generation from JSON schemas

// Example validation for Transform2D rotation
func (x *Xfrm) Validate(ctx *validation.Context) []validation.Error {
    var errors []validation.Error

    // Rotation must be 0-21599999 (0-360 degrees)
    if x.Rot != nil && x.Rot.HasValue() {
        rot := x.Rot.Value()
        if rot < 0 || rot > 21599999 {
            errors = append(errors, validation.Error{
                Element:     x,
                Description: "rotation must be between 0 and 21599999",
                Path:        ctx.Path + "/xfrm/@rot",
            })
        }
    }

    // Validate children
    if x.Off != nil {
        errors = append(errors, x.Off.Validate(ctx.Child("off"))...)
    }
    if x.Ext != nil {
        errors = append(errors, x.Ext.Validate(ctx.Child("ext"))...)
    }

    return errors
}
```

**Rationale:** Full validation support matching Word and Presentation SDKs.

## Key Type Mappings

### C# to Go Type Mapping

| C# DrawingML Type | Go Type | Package |
|-------------------|---------|---------|
| `Transform2D` | `Xfrm` | `drawingml/main` |
| `PresetGeometry` | `PrstGeom` | `drawingml/main` |
| `CustomGeometry` | `CustGeom` | `drawingml/main` |
| `SolidFill` | `SolidFill` | `drawingml/main` |
| `GradientFill` | `GradFill` | `drawingml/main` |
| `BlipFill` | `BlipFill` | `drawingml/main` |
| `PatternFill` | `PattFill` | `drawingml/main` |
| `LineProperties` | `Ln` | `drawingml/main` |
| `EffectList` | `EffectLst` | `drawingml/main` |
| `TextBody` | `TxBody` | `drawingml/main` |
| `Paragraph` | `Para` | `drawingml/main` |
| `Run` | `Run` | `drawingml/main` |
| `RunProperties` | `RunPr` | `drawingml/main` |
| `ChartSpace` | `ChartSpace` | `drawingml/chart` |
| `WorksheetDrawing` | `WsDr` | `drawingml/spreadsheet` |
| `TwoCellAnchor` | `TwoCellAnchor` | `drawingml/spreadsheet` |
| `Inline` | `Inline` | `drawingml/word` |
| `Anchor` | `Anchor` | `drawingml/word` |
| `Picture` | `Pic` | `drawingml/picture` |

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| Type explosion | High - 600+ types | Code generation handles scale |
| Namespace complexity | Medium - 5+ namespaces | Subpackage organization |
| Cross-SDK coupling | Medium - all SDKs depend on drawingml | Stable API, semantic versioning |
| Schema evolution | Low - ECMA-376 stable | Version-specific types if needed |

## Resolved Design Questions

### Q1: Package Organization
**Decision:** Subpackages by namespace (main, chart, spreadsheet, word, picture).

### Q2: Element Naming
**Decision:** Match XML element names with Go conventions (Xfrm, PrstGeom, TxBody).

### Q3: Fill/Color Interface Pattern
**Decision:** Marker interface pattern with concrete implementations.

### Q4: EMU Handling
**Decision:** Type-safe EMU type with conversion methods in units subpackage.

### Q5: Thread Safety
**Decision:** sync.RWMutex at document level (handled by parent SDK, not drawingml package).

## Appendix: Namespace URIs

| Prefix | URI |
|--------|-----|
| `a` | `http://schemas.openxmlformats.org/drawingml/2006/main` |
| `c` | `http://schemas.openxmlformats.org/drawingml/2006/chart` |
| `xdr` | `http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing` |
| `wp` | `http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing` |
| `pic` | `http://schemas.openxmlformats.org/drawingml/2006/picture` |
| `dgm` | `http://schemas.openxmlformats.org/drawingml/2006/diagram` |
| `r` | `http://schemas.openxmlformats.org/officeDocument/2006/relationships` |

## Appendix: Preset Geometry Categories

| Category | Count | Examples |
|----------|-------|----------|
| Basic Shapes | 32 | rect, ellipse, triangle, diamond |
| Block Arrows | 27 | rightArrow, leftArrow, upDownArrow |
| Callouts | 20 | wedgeRoundRectCallout, cloudCallout |
| Flowchart | 28 | flowChartProcess, flowChartDecision |
| Stars & Banners | 16 | star5, ribbon2, wave |
| Math Shapes | 8 | mathPlus, mathMinus, mathEqual |
| Action Buttons | 14 | actionButtonHome, actionButtonEnd |
| Connectors | 9 | straightConnector1, bentConnector3 |
| **Total** | **154+** | |
