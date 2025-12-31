# SmartArt Layout Engine - Design

## Architecture Overview

The SmartArt layout engine uses a **pluggable algorithm architecture** where each layout type implements the `LayoutEngine` interface. Layouts are registered in a central registry and selected based on the diagram's layout type.

```
┌─────────────────────────────────────────┐
│         SmartArt Diagram                │
│  (data model: nodes, connections)       │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│       Layout Engine Registry            │
│  SelectLayout(layoutType) → Engine      │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│      LayoutEngine Interface              │
│  - Layout(diagram, bounds) error        │
│  - CalculateBounds(diagram) Rectangle   │
└──────────────┬──────────────────────────┘
               │
       ┌───────┴────────┬──────────┬───────────┬─────────┐
       ▼                ▼          ▼           ▼         ▼
┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐
│Hierarchy │  │   List   │  │  Cycle   │  │ Pyramid  │  │  Matrix  │
│ Layout   │  │  Layout  │  │  Layout  │  │  Layout  │  │  Layout  │
└──────────┘  └──────────┘  └──────────┘  └──────────┘  └──────────┘
```

## Core Types

### LayoutEngine Interface

```go
package layout

// LayoutEngine applies automatic layout to SmartArt diagrams
type LayoutEngine interface {
    // Layout applies the layout algorithm to the diagram within the specified bounds.
    // It updates shape positions and sizes in the diagram's drawing data.
    Layout(diagram *SmartArt, bounds Rectangle) error

    // CalculateBounds calculates the minimum bounding rectangle needed for the diagram.
    CalculateBounds(diagram *SmartArt) Rectangle

    // SupportedTypes returns the layout types this engine supports.
    SupportedTypes() []LayoutType
}

// LayoutType identifies the layout algorithm
type LayoutType int

const (
    LayoutHierarchy LayoutType = iota  // Org chart, tree
    LayoutList                          // Vertical/horizontal list
    LayoutCycle                         // Circular process
    LayoutPyramid                       // Stacked levels
    LayoutMatrix                        // Grid layout
)
```

### SmartArt Extensions

```go
package presentation

// SmartArt represents a SmartArt diagram in a presentation.
// This extends the diagram elements generated in Phase 1.
type SmartArt struct {
    data       *diagram.DataModel         // From Phase 1
    layoutDef  *diagram.LayoutDefinition  // From Phase 1
    style      *diagram.StyleDefinition   // From Phase 1
    colors     *diagram.ColorTransform    // From Phase 1

    // Phase 2 additions
    layoutEngine layout.LayoutEngine
    autoLayout   bool
    bounds       Rectangle
}

// Layout applies the layout algorithm to position shapes.
// This is the main entry point for automatic layout.
func (sa *SmartArt) Layout() error {
    if sa.layoutEngine == nil {
        return fmt.Errorf("no layout engine set")
    }
    return sa.layoutEngine.Layout(sa, sa.bounds)
}

// SetLayoutType changes the layout algorithm.
func (sa *SmartArt) SetLayoutType(layoutType layout.LayoutType) error {
    engine, err := layout.GetEngine(layoutType)
    if err != nil {
        return err
    }
    sa.layoutEngine = engine
    if sa.autoLayout {
        return sa.Layout()
    }
    return nil
}

// GetLayoutEngine returns the current layout engine.
func (sa *SmartArt) GetLayoutEngine() layout.LayoutEngine {
    return sa.layoutEngine
}

// AutoLayout enables or disables automatic re-layout when data changes.
func (sa *SmartArt) AutoLayout(enabled bool) {
    sa.autoLayout = enabled
}
```

### High-Level Creation API

```go
package presentation

// NewSmartArt creates a new SmartArt diagram with the specified layout type.
func NewSmartArt(layoutType layout.LayoutType) (*SmartArt, error) {
    sa := &SmartArt{
        data:       diagram.NewDataModel(),
        layoutDef:  diagram.NewLayoutDefinition(),
        style:      diagram.NewStyleDefinition(),
        colors:     diagram.NewColorTransform(),
        autoLayout: true,
    }

    err := sa.SetLayoutType(layoutType)
    if err != nil {
        return nil, err
    }

    return sa, nil
}

// SmartArtFromData creates a SmartArt diagram from a data structure.
func SmartArtFromData(data SmartArtData, layoutType layout.LayoutType) (*SmartArt, error) {
    sa, err := NewSmartArt(layoutType)
    if err != nil {
        return nil, err
    }

    // Populate data model from input data
    for _, item := range data.Items {
        node := sa.data.AddPoint(item.ID, item.Text)
        node.SetType(item.Type)
    }

    for _, conn := range data.Connections {
        sa.data.AddConnection(conn.SourceID, conn.DestID, conn.Type)
    }

    // Trigger layout
    if sa.autoLayout {
        return sa, sa.Layout()
    }
    return sa, nil
}

// SmartArtData is a simple data structure for creating SmartArt diagrams.
type SmartArtData struct {
    Items       []SmartArtItem
    Connections []SmartArtConnection
}

type SmartArtItem struct {
    ID   string
    Text string
    Type PointType  // node, assistant, etc.
}

type SmartArtConnection struct {
    SourceID string
    DestID   string
    Type     ConnectionType  // parent-child, sibling, etc.
}
```

## Layout Algorithms

### 1. Hierarchy Layout (Org Chart)

**Algorithm: Reingold-Tilford Tree Layout**

```go
package layout

type HierarchyLayout struct {
    orientation Orientation  // vertical or horizontal
    spacing     Spacing
}

func (hl *HierarchyLayout) Layout(diagram *SmartArt, bounds Rectangle) error {
    // 1. Build tree structure from data model
    root := hl.buildTree(diagram.Data())

    // 2. Calculate positions using Reingold-Tilford algorithm
    hl.calculatePositions(root)

    // 3. Scale and center within bounds
    hl.scaleAndCenter(root, bounds)

    // 4. Update shape positions in drawing data
    hl.applyPositions(diagram, root)

    // 5. Route connectors
    hl.routeConnectors(diagram, root)

    return nil
}

func (hl *HierarchyLayout) calculatePositions(node *TreeNode) {
    // Reingold-Tilford Algorithm:
    // 1. Post-order traversal: layout children first
    for _, child := range node.Children {
        hl.calculatePositions(child)
    }

    // 2. Position this node relative to children
    if len(node.Children) == 0 {
        // Leaf node: position at origin
        node.X = 0
        node.Y = 0
    } else if len(node.Children) == 1 {
        // Single child: center above child
        child := node.Children[0]
        node.X = child.X
        node.Y = child.Y - hl.spacing.Vertical - node.Height
    } else {
        // Multiple children: center above children's midpoint
        leftmost := node.Children[0]
        rightmost := node.Children[len(node.Children)-1]
        midpoint := (leftmost.X + rightmost.X) / 2
        node.X = midpoint
        node.Y = leftmost.Y - hl.spacing.Vertical - node.Height
    }

    // 3. Ensure sibling separation (horizontal spacing)
    hl.ensureSeparation(node)

    // 4. Handle assistant nodes (dotted line subordinates)
    hl.positionAssistants(node)
}

func (hl *HierarchyLayout) ensureSeparation(node *TreeNode) {
    // Ensure minimum spacing between siblings
    for i := 1; i < len(node.Children); i++ {
        left := node.Children[i-1]
        right := node.Children[i]

        minSep := left.Width/2 + right.Width/2 + hl.spacing.Horizontal
        if right.X - left.X < minSep {
            shift := minSep - (right.X - left.X)
            hl.shiftSubtree(right, shift)
        }
    }
}

func (hl *HierarchyLayout) shiftSubtree(node *TreeNode, dx float64) {
    // Shift node and all descendants horizontally
    node.X += dx
    for _, child := range node.Children {
        hl.shiftSubtree(child, dx)
    }
}

// Horizontal variant: rotate 90 degrees
func (hl *HierarchyLayout) applyOrientation(node *TreeNode) {
    if hl.orientation == OrientationHorizontal {
        // Swap X and Y for all nodes
        hl.rotateTree(node)
    }
}
```

**Features**:
- Balanced tree layout (minimizes tree width)
- Configurable spacing (horizontal, vertical)
- Auto-size based on text content
- Assistant node support (dotted line connections)
- Horizontal and vertical orientations

### 2. List Layout

**Algorithm: Linear Flow with Equal Spacing**

```go
package layout

type ListLayout struct {
    orientation Orientation  // vertical or horizontal
    spacing     Spacing
    alignment   Alignment    // left, center, right
}

func (ll *ListLayout) Layout(diagram *SmartArt, bounds Rectangle) error {
    // 1. Get items from data model
    items := diagram.Data().GetPoints()

    // 2. Measure text and calculate sizes
    ll.measureItems(items)

    // 3. Calculate positions
    positions := ll.calculatePositions(items, bounds)

    // 4. Apply positions to shapes
    ll.applyPositions(diagram, items, positions)

    // 5. Route connectors
    ll.routeConnectors(diagram, items)

    return nil
}

func (ll *ListLayout) calculatePositions(items []*Point, bounds Rectangle) []Position {
    positions := make([]Position, len(items))

    if ll.orientation == OrientationVertical {
        // Vertical list: stack items top to bottom
        totalHeight := 0.0
        for _, item := range items {
            totalHeight += item.Height
        }
        totalHeight += float64(len(items)-1) * ll.spacing.Vertical

        // Start position (centered or aligned)
        y := bounds.Y
        if ll.alignment == AlignmentCenter {
            y = bounds.Y + (bounds.Height - totalHeight) / 2
        }

        // Position each item
        for i, item := range items {
            x := ll.alignX(item, bounds)
            positions[i] = Position{X: x, Y: y}
            y += item.Height + ll.spacing.Vertical
        }
    } else {
        // Horizontal list: arrange left to right
        totalWidth := 0.0
        for _, item := range items {
            totalWidth += item.Width
        }
        totalWidth += float64(len(items)-1) * ll.spacing.Horizontal

        // Start position
        x := bounds.X
        if ll.alignment == AlignmentCenter {
            x = bounds.X + (bounds.Width - totalWidth) / 2
        }

        // Position each item
        for i, item := range items {
            y := ll.alignY(item, bounds)
            positions[i] = Position{X: x, Y: y}
            x += item.Width + ll.spacing.Horizontal
        }
    }

    return positions
}

func (ll *ListLayout) routeConnectors(diagram *SmartArt, items []*Point) {
    // Simple straight-line connectors between consecutive items
    for i := 0; i < len(items)-1; i++ {
        src := items[i]
        dst := items[i+1]

        // Calculate connector endpoints
        var startPoint, endPoint Point
        if ll.orientation == OrientationVertical {
            // Vertical: connect bottom of item[i] to top of item[i+1]
            startPoint = Point{X: src.X + src.Width/2, Y: src.Y + src.Height}
            endPoint = Point{X: dst.X + dst.Width/2, Y: dst.Y}
        } else {
            // Horizontal: connect right of item[i] to left of item[i+1]
            startPoint = Point{X: src.X + src.Width, Y: src.Y + src.Height/2}
            endPoint = Point{X: dst.X, Y: dst.Y + dst.Height/2}
        }

        // Add connector to diagram
        diagram.AddConnector(startPoint, endPoint)
    }
}
```

**Features**:
- Vertical or horizontal flow
- Equal spacing between items
- Configurable alignment (left, center, right for vertical; top, center, bottom for horizontal)
- Auto-size to container width/height
- Straight-line connectors

### 3. Cycle Layout

**Algorithm: Circular Arrangement**

```go
package layout

type CycleLayout struct {
    direction   Direction  // clockwise or counter-clockwise
    centerText  bool
}

func (cl *CycleLayout) Layout(diagram *SmartArt, bounds Rectangle) error {
    // 1. Get items
    items := diagram.Data().GetPoints()

    // 2. Measure items
    cl.measureItems(items)

    // 3. Calculate circle parameters
    center, radius := cl.calculateCircle(items, bounds)

    // 4. Position items on circle
    cl.positionOnCircle(items, center, radius)

    // 5. Apply positions
    cl.applyPositions(diagram, items)

    // 6. Add directional arrows
    cl.addArrows(diagram, items, center, radius)

    // 7. Add center text (if enabled)
    if cl.centerText {
        cl.addCenterText(diagram, center)
    }

    return nil
}

func (cl *CycleLayout) calculateCircle(items []*Point, bounds Rectangle) (center Point, radius float64) {
    // Center of bounds
    center = Point{
        X: bounds.X + bounds.Width/2,
        Y: bounds.Y + bounds.Height/2,
    }

    // Radius: fit all items on perimeter with spacing
    maxItemSize := 0.0
    for _, item := range items {
        size := math.Max(item.Width, item.Height)
        if size > maxItemSize {
            maxItemSize = size
        }
    }

    // Approximate radius needed to fit items on perimeter
    n := float64(len(items))
    circumference := n * (maxItemSize + 20) // 20pt spacing
    radius = circumference / (2 * math.Pi)

    // Ensure fits within bounds
    maxRadius := math.Min(bounds.Width, bounds.Height) / 2 * 0.8
    if radius > maxRadius {
        radius = maxRadius
    }

    return center, radius
}

func (cl *CycleLayout) positionOnCircle(items []*Point, center Point, radius float64) {
    n := len(items)
    angleStep := 2 * math.Pi / float64(n)

    // Start angle (top of circle)
    startAngle := -math.Pi / 2
    if cl.direction == DirectionCounterClockwise {
        angleStep = -angleStep
    }

    for i, item := range items {
        angle := startAngle + float64(i)*angleStep

        // Position on circle perimeter
        item.X = center.X + radius*math.Cos(angle) - item.Width/2
        item.Y = center.Y + radius*math.Sin(angle) - item.Height/2
    }
}

func (cl *CycleLayout) addArrows(diagram *SmartArt, items []*Point, center Point, radius float64) {
    n := len(items)
    for i := 0; i < n; i++ {
        src := items[i]
        dst := items[(i+1)%n]  // Wrap around to first item

        // Calculate arrow path (curved along circle)
        // Use quadratic bezier curve with control point at circle perimeter
        angle1 := math.Atan2(src.Y+src.Height/2-center.Y, src.X+src.Width/2-center.X)
        angle2 := math.Atan2(dst.Y+dst.Height/2-center.Y, dst.X+dst.Width/2-center.X)
        angleMid := (angle1 + angle2) / 2

        startPoint := Point{
            X: src.X + src.Width/2,
            Y: src.Y + src.Height/2,
        }
        endPoint := Point{
            X: dst.X + dst.Width/2,
            Y: dst.Y + dst.Height/2,
        }
        controlPoint := Point{
            X: center.X + radius*1.1*math.Cos(angleMid),
            Y: center.Y + radius*1.1*math.Sin(angleMid),
        }

        // Add curved connector with arrow
        diagram.AddCurvedConnector(startPoint, controlPoint, endPoint, true /* arrow */)
    }
}
```

**Features**:
- Circular arrangement on perimeter
- Equal angular spacing
- Clockwise or counter-clockwise direction
- Curved directional arrows
- Optional center text

### 4. Pyramid Layout

**Algorithm: Stacked Levels with Proportional Sizing**

```go
package layout

type PyramidLayout struct {
    orientation Orientation  // normal (apex top) or inverted (apex bottom)
}

func (pl *PyramidLayout) Layout(diagram *SmartArt, bounds Rectangle) error {
    // 1. Get levels from data model
    levels := diagram.Data().GetLevels()

    // 2. Calculate level sizes (proportional)
    sizes := pl.calculateLevelSizes(levels, bounds)

    // 3. Position levels
    pl.positionLevels(levels, sizes, bounds)

    // 4. Apply positions
    pl.applyPositions(diagram, levels)

    return nil
}

func (pl *PyramidLayout) calculateLevelSizes(levels [][]*Point, bounds Rectangle) []LevelSize {
    n := len(levels)
    sizes := make([]LevelSize, n)

    levelHeight := bounds.Height / float64(n)

    for i, level := range levels {
        // Width decreases linearly from base to apex
        var widthRatio float64
        if pl.orientation == OrientationNormal {
            // Normal pyramid: top is narrowest
            widthRatio = float64(n-i) / float64(n)
        } else {
            // Inverted pyramid: top is widest
            widthRatio = float64(i+1) / float64(n)
        }

        sizes[i] = LevelSize{
            Width:  bounds.Width * widthRatio,
            Height: levelHeight,
        }
    }

    return sizes
}

func (pl *PyramidLayout) positionLevels(levels [][]*Point, sizes []LevelSize, bounds Rectangle) {
    y := bounds.Y

    for i, level := range levels {
        size := sizes[i]

        // Center level horizontally
        x := bounds.X + (bounds.Width-size.Width)/2

        // Position items within level
        if len(level) == 1 {
            // Single item: center in level
            item := level[0]
            item.X = x + (size.Width-item.Width)/2
            item.Y = y + (size.Height-item.Height)/2
        } else {
            // Multiple items: distribute evenly
            itemWidth := size.Width / float64(len(level))
            for j, item := range level {
                item.X = x + float64(j)*itemWidth + (itemWidth-item.Width)/2
                item.Y = y + (size.Height-item.Height)/2
            }
        }

        y += size.Height
    }
}
```

**Features**:
- Stacked levels with proportional sizing
- Normal or inverted orientation
- Trapezoid shapes for levels
- Text centered in each level
- Multiple items per level supported

### 5. Matrix Layout

**Algorithm: Grid Positioning**

```go
package layout

type MatrixLayout struct {
    rows    int
    columns int
}

func (ml *MatrixLayout) Layout(diagram *SmartArt, bounds Rectangle) error {
    // 1. Get items
    items := diagram.Data().GetPoints()

    // 2. Determine grid size (if not specified)
    if ml.rows == 0 || ml.columns == 0 {
        ml.determineGridSize(len(items))
    }

    // 3. Calculate cell size
    cellWidth := bounds.Width / float64(ml.columns)
    cellHeight := bounds.Height / float64(ml.rows)

    // 4. Position items in grid
    for i, item := range items {
        row := i / ml.columns
        col := i % ml.columns

        // Center item within cell
        x := bounds.X + float64(col)*cellWidth + (cellWidth-item.Width)/2
        y := bounds.Y + float64(row)*cellHeight + (cellHeight-item.Height)/2

        item.X = x
        item.Y = y
    }

    // 5. Apply positions
    ml.applyPositions(diagram, items)

    return nil
}

func (ml *MatrixLayout) determineGridSize(itemCount int) {
    // Default to square-ish grid
    ml.columns = int(math.Ceil(math.Sqrt(float64(itemCount))))
    ml.rows = int(math.Ceil(float64(itemCount) / float64(ml.columns)))
}
```

**Features**:
- Grid positioning (rows × columns)
- Auto-determine grid size for square layout
- Evenly sized cells
- Items centered in cells
- Support for merged cells (future enhancement)

## Constraint System

Layouts use a constraint system for spacing, sizing, and alignment:

```go
package constraints

// Spacing defines spacing between layout elements
type Spacing struct {
    Horizontal float64  // Horizontal spacing in points
    Vertical   float64  // Vertical spacing in points
    Padding    float64  // Padding around shapes
}

// Sizing defines how shapes are sized
type Sizing struct {
    MinWidth   float64
    MinHeight  float64
    MaxWidth   float64
    MaxHeight  float64
    AspectRatio float64  // Width/Height ratio (0 = unconstrained)
}

// Alignment defines how shapes are aligned
type Alignment int

const (
    AlignmentLeft Alignment = iota
    AlignmentCenter
    AlignmentRight
    AlignmentTop
    AlignmentMiddle
    AlignmentBottom
)

// LayoutConstraints bundles all constraints for a layout
type LayoutConstraints struct {
    Spacing   Spacing
    Sizing    Sizing
    Alignment Alignment
}

// Default constraints
var DefaultConstraints = LayoutConstraints{
    Spacing: Spacing{
        Horizontal: 20,  // 20pt
        Vertical:   20,  // 20pt
        Padding:    5,   // 5pt
    },
    Sizing: Sizing{
        MinWidth:    50,   // 50pt
        MinHeight:   30,   // 30pt
        MaxWidth:    200,  // 200pt
        MaxHeight:   100,  // 100pt
        AspectRatio: 0,    // Unconstrained
    },
    Alignment: AlignmentCenter,
}
```

## Text Measurement

Layouts use the PDF text layout engine to measure text and auto-size shapes:

```go
package layout

import (
    "goffice/pdf/layout"
    "goffice/pdf/font"
)

// TextMeasurer measures text dimensions for shape auto-sizing
type TextMeasurer struct {
    fontCache *font.FontCache
}

func NewTextMeasurer() *TextMeasurer {
    return &TextMeasurer{
        fontCache: font.NewCache(),
    }
}

func (tm *TextMeasurer) MeasureText(text string, fontName string, fontSize float64) (width, height float64) {
    // Load font
    f, err := tm.fontCache.LoadFont(fontName)
    if err != nil {
        // Fallback to default metrics
        return float64(len(text)) * fontSize * 0.5, fontSize
    }

    // Measure text
    width = f.MeasureString(text, fontSize)
    height = fontSize

    return width, height
}

// AutoSizeShape calculates shape size based on text content
func (tm *TextMeasurer) AutoSizeShape(item *Point, constraints LayoutConstraints) {
    // Measure text
    width, height := tm.MeasureText(item.Text, "Arial", 12)

    // Add padding
    width += constraints.Spacing.Padding * 2
    height += constraints.Spacing.Padding * 2

    // Apply constraints
    if width < constraints.Sizing.MinWidth {
        width = constraints.Sizing.MinWidth
    }
    if width > constraints.Sizing.MaxWidth {
        width = constraints.Sizing.MaxWidth
    }
    if height < constraints.Sizing.MinHeight {
        height = constraints.Sizing.MinHeight
    }
    if height > constraints.Sizing.MaxHeight {
        height = constraints.Sizing.MaxHeight
    }

    // Apply aspect ratio (if specified)
    if constraints.Sizing.AspectRatio > 0 {
        currentRatio := width / height
        if currentRatio > constraints.Sizing.AspectRatio {
            width = height * constraints.Sizing.AspectRatio
        } else {
            height = width / constraints.Sizing.AspectRatio
        }
    }

    item.Width = width
    item.Height = height
}
```

## PDF Rendering Integration

The layout engine positions shapes, and the PDF renderer draws them:

```go
package drawing

// RenderDiagram renders a SmartArt diagram to PDF
func (dr *DiagramRenderer) RenderDiagram(diagram *SmartArt) error {
    // Phase 1 rendered placeholder box
    // Phase 2 renders positioned shapes

    // 1. Apply layout (if not already laid out)
    if !diagram.IsLaidOut() {
        if err := diagram.Layout(); err != nil {
            return err
        }
    }

    // 2. Render shapes
    for _, point := range diagram.Data().GetPoints() {
        dr.renderShape(point)
    }

    // 3. Render connectors
    for _, conn := range diagram.Data().GetConnections() {
        dr.renderConnector(conn)
    }

    return nil
}

func (dr *DiagramRenderer) renderShape(point *Point) {
    // Get shape geometry (rectangle, ellipse, etc.)
    shape := point.GetShape()

    // Draw shape background
    dr.setFillColor(shape.FillColor)
    dr.drawShape(shape.Geometry, point.X, point.Y, point.Width, point.Height)

    // Draw shape border
    dr.setStrokeColor(shape.BorderColor)
    dr.setLineWidth(shape.BorderWidth)
    dr.strokeShape(shape.Geometry, point.X, point.Y, point.Width, point.Height)

    // Draw text
    dr.drawText(point.Text, point.X, point.Y, point.Width, point.Height, shape.TextStyle)
}

func (dr *DiagramRenderer) renderConnector(conn *Connection) {
    // Get connector endpoints
    srcPoint := conn.Source.GetConnectionPoint(conn.SourceLocation)
    dstPoint := conn.Dest.GetConnectionPoint(conn.DestLocation)

    // Draw connector line
    dr.setStrokeColor(conn.LineColor)
    dr.setLineWidth(conn.LineWidth)

    if conn.Curved {
        // Draw curved line (bezier)
        dr.drawBezier(srcPoint, conn.ControlPoint, dstPoint)
    } else {
        // Draw straight line
        dr.drawLine(srcPoint, dstPoint)
    }

    // Draw arrow (if specified)
    if conn.Arrow {
        dr.drawArrow(dstPoint, conn.ArrowSize)
    }
}
```

## Layout Engine Registry

```go
package layout

// Registry holds all registered layout engines
type Registry struct {
    engines map[LayoutType]LayoutEngine
}

var defaultRegistry *Registry

func init() {
    defaultRegistry = &Registry{
        engines: make(map[LayoutType]LayoutEngine),
    }

    // Register built-in layouts
    defaultRegistry.Register(LayoutHierarchy, &HierarchyLayout{
        orientation: OrientationVertical,
        spacing:     DefaultConstraints.Spacing,
    })
    defaultRegistry.Register(LayoutList, &ListLayout{
        orientation: OrientationVertical,
        spacing:     DefaultConstraints.Spacing,
        alignment:   AlignmentCenter,
    })
    defaultRegistry.Register(LayoutCycle, &CycleLayout{
        direction:  DirectionClockwise,
        centerText: false,
    })
    defaultRegistry.Register(LayoutPyramid, &PyramidLayout{
        orientation: OrientationNormal,
    })
    defaultRegistry.Register(LayoutMatrix, &MatrixLayout{
        rows:    0,  // Auto-determine
        columns: 0,  // Auto-determine
    })
}

// Register adds a layout engine to the registry
func (r *Registry) Register(layoutType LayoutType, engine LayoutEngine) {
    r.engines[layoutType] = engine
}

// GetEngine retrieves a layout engine by type
func (r *Registry) GetEngine(layoutType LayoutType) (LayoutEngine, error) {
    engine, ok := r.engines[layoutType]
    if !ok {
        return nil, fmt.Errorf("no layout engine registered for type %d", layoutType)
    }
    return engine, nil
}

// GetEngine retrieves from default registry
func GetEngine(layoutType LayoutType) (LayoutEngine, error) {
    return defaultRegistry.GetEngine(layoutType)
}
```

## Testing Strategy

### Unit Tests

Each layout algorithm has comprehensive unit tests:

```go
func TestHierarchyLayout(t *testing.T) {
    tests := []struct {
        name     string
        data     SmartArtData
        bounds   Rectangle
        expected []Position
    }{
        {
            name: "Simple 3-node tree",
            data: SmartArtData{
                Items: []SmartArtItem{
                    {ID: "1", Text: "Root"},
                    {ID: "2", Text: "Child 1"},
                    {ID: "3", Text: "Child 2"},
                },
                Connections: []SmartArtConnection{
                    {SourceID: "1", DestID: "2", Type: ConnectionParentChild},
                    {SourceID: "1", DestID: "3", Type: ConnectionParentChild},
                },
            },
            bounds: Rectangle{X: 0, Y: 0, Width: 200, Height: 100},
            expected: []Position{
                {X: 100, Y: 0},    // Root centered top
                {X: 50, Y: 50},    // Child 1 left
                {X: 150, Y: 50},   // Child 2 right
            },
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sa, err := SmartArtFromData(tt.data, LayoutHierarchy)
            if err != nil {
                t.Fatal(err)
            }

            sa.SetBounds(tt.bounds)
            if err := sa.Layout(); err != nil {
                t.Fatal(err)
            }

            // Verify positions
            for i, item := range sa.Data().GetPoints() {
                if !positionsEqual(item.Position(), tt.expected[i]) {
                    t.Errorf("Item %d position = %v, expected %v", i, item.Position(), tt.expected[i])
                }
            }
        })
    }
}
```

### Integration Tests

Test complete workflow from creation to PDF rendering:

```go
func TestSmartArtCreationAndRendering(t *testing.T) {
    // Create SmartArt diagram
    sa, err := NewSmartArt(LayoutHierarchy)
    if err != nil {
        t.Fatal(err)
    }

    // Add data
    root := sa.AddPoint("CEO")
    mgr1 := sa.AddPoint("Manager 1")
    mgr2 := sa.AddPoint("Manager 2")
    sa.AddConnection(root, mgr1, ConnectionParentChild)
    sa.AddConnection(root, mgr2, ConnectionParentChild)

    // Apply layout
    if err := sa.Layout(); err != nil {
        t.Fatal(err)
    }

    // Add to presentation
    pres := presentation.NewPresentation()
    slide := pres.AddSlide()
    slide.AddSmartArt(sa)

    // Save
    if err := pres.SaveToFile("test_smartart.pptx"); err != nil {
        t.Fatal(err)
    }

    // Render to PDF
    pdfDoc := pdf.NewDocument()
    if err := pdfDoc.RenderPresentation(pres); err != nil {
        t.Fatal(err)
    }
    if err := pdfDoc.SaveToFile("test_smartart.pdf"); err != nil {
        t.Fatal(err)
    }

    // Verify PDF contains positioned shapes
    // (visual regression test - compare to reference PDF)
}
```

### Visual Regression Tests

Compare rendered output to reference images:

```go
func TestVisualRegression(t *testing.T) {
    tests := []string{
        "hierarchy_simple",
        "hierarchy_complex",
        "list_vertical",
        "list_horizontal",
        "cycle_4items",
        "pyramid_3levels",
        "matrix_2x2",
    }

    for _, testName := range tests {
        t.Run(testName, func(t *testing.T) {
            // Load test data
            data := loadTestData(testName)

            // Create and layout
            sa, _ := SmartArtFromData(data, getLayoutType(testName))
            sa.Layout()

            // Render to image
            img := renderToImage(sa)

            // Compare to reference
            refImg := loadReferenceImage(testName)
            if !imagesEqual(img, refImg) {
                // Save diff image for inspection
                saveDiffImage(testName, img, refImg)
                t.Errorf("Visual regression: %s differs from reference", testName)
            }
        })
    }
}
```

## Performance Considerations

### Caching

Layout calculations are cached to avoid redundant work:

```go
type SmartArt struct {
    // ... other fields
    layoutCached bool
    layoutVersion int  // Incremented when data changes
}

func (sa *SmartArt) Layout() error {
    // Skip if already laid out and data hasn't changed
    if sa.layoutCached && sa.data.Version() == sa.layoutVersion {
        return nil
    }

    // Apply layout
    if err := sa.layoutEngine.Layout(sa, sa.bounds); err != nil {
        return err
    }

    // Mark as cached
    sa.layoutCached = true
    sa.layoutVersion = sa.data.Version()

    return nil
}

func (sa *SmartArt) invalidateLayout() {
    sa.layoutCached = false
}
```

### Complexity Analysis

- **Hierarchy Layout**: O(n) where n = number of nodes (tree traversal)
- **List Layout**: O(n) where n = number of items
- **Cycle Layout**: O(n) where n = number of items
- **Pyramid Layout**: O(n) where n = number of items
- **Matrix Layout**: O(n) where n = number of items

All algorithms are linear time, suitable for diagrams with hundreds of items.

## Open Questions

1. **Connector Routing**: Should we implement advanced connector routing (orthogonal, A*) in Phase 2 or defer to Phase 3?
   - **Recommendation**: Start with simple straight-line and bezier connectors, defer advanced routing

2. **Text Wrapping**: How should we handle text that doesn't fit in shapes?
   - **Recommendation**: Wrap text within shape bounds, use ellipsis if still doesn't fit

3. **Shape Overlap**: Should layouts prevent shape overlap or allow it?
   - **Recommendation**: Prevent overlap in initial layouts, add validation

4. **Animation Support**: Should layout changes be animatable?
   - **Recommendation**: Defer to future proposal, layout engine returns static positions

## Future Enhancements

**Phase 3 (Future Proposal)**:
- Advanced layout types (picture, relationship, venn, radial)
- Custom layout definitions
- Advanced connector routing (orthogonal, A*)
- Shape overlap detection and resolution
- Layout animation
- SmartArt styles and themes integration
