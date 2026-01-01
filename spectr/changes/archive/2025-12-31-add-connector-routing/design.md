# Design: Complete Connector Shape Routing and Management

## Problem Statement

ConnectionShape elements exist in goffice but lack all practical functionality:
- No way to specify which shapes a connector connects
- No intelligent routing to avoid obstacles
- No automatic updates when connected shapes move
- No connection point system for precise attachment
- PDF rendering treats connectors as static shapes, ignoring connection semantics

This blocks users from creating flowcharts, org charts, network diagrams, and any other connected visualization programmatically.

## Current Architecture

```
Presentation Slide
    ↓
ShapeTree
    ├─ Shape (basic shapes) ────── Has position, geometry
    ├─ Picture ────────────────── Has position, image
    ├─ GroupShape ─────────────── Container for shapes
    └─ ConnectionShape ────────── EXISTS but non-functional
            ↓
        Only has:
        - NewConnectionShape() constructor
        - Basic XML parse/serialize

        Missing:
        - Connection to other shapes
        - Routing algorithms
        - Path calculation
        - Update mechanisms
```

**Current ConnectionShape** (`presentation/elements/shape_tree.go:543-610`):
```go
type ConnectionShape struct {
    *openxml.CompositeElementBase
}

func NewConnectionShape() *ConnectionShape {
    // Creates XML structure
    // Adds nvCxnSpPr, cNvCxnSpPr elements
    // NO connection or routing logic
}

// END OF IMPLEMENTATION
```

## Proposed Architecture

```
Presentation Slide
    ↓
ShapeTree
    ├─ Shape
    │   ├─ ConnectionSites() ────── NEW: Attachment points
    │   ├─ AddConnectionSite() ──── NEW: Custom points
    │   └─ GetConnectionSite() ──── NEW: Retrieve by index
    │
    └─ ConnectionShape
        ├─ StartConnection() ───────── NEW: Get start shape/site
        ├─ SetStartConnection() ────── NEW: Connect to shape
        ├─ EndConnection() ─────────── NEW: Get end shape/site
        ├─ SetEndConnection() ──────── NEW: Connect to shape
        └─ Path calculation ────────── NEW: Routing algorithms
                ↓
        NEW: presentation/routing/
            ├─ Router interface
            ├─ StraightRouter ──────── Direct line
            ├─ ElbowRouter ─────────── A* orthogonal routing
            └─ CurvedRouter ────────── Bezier curves
                    ↓
            NEW: Slide.ConnectShapes()
            NEW: Slide.UpdateAllConnectors()
                    ↓
            NEW: pdf/drawing/ConnectorRenderer
                ├─ Resolve shape references
                ├─ Calculate connection points
                ├─ Render routed path
                └─ Draw arrow decorations
```

## Design Alternatives Considered

### Alternative 1: Store Direct Shape Pointers in ConnectionShape

**Approach**: Keep Go pointers to start/end shapes in ConnectionShape struct.

**Pros**:
- Fast access (no ID lookup)
- Type-safe

**Cons**:
- Breaks on serialization (pointers can't be saved to XML)
- Fragile to shape deletion (dangling pointers)
- Breaks copy/paste operations
- Doesn't match ECMA-376 schema (uses shape IDs)

**Decision**: REJECTED - Use shape ID strings per ECMA-376 standard

### Alternative 2: Automatic Routing on Every Shape Move

**Approach**: Hook into Shape.SetPosition() to auto-recalculate connected connectors.

**Pros**:
- Connectors always up-to-date
- No manual UpdateAllConnectors() call needed

**Cons**:
- Performance: Recalculating on every drag event (hundreds of calls)
- Unexpected side effects when moving shapes
- Doesn't match Office behavior (updates on mouse release)
- Complex dependency tracking

**Decision**: REJECTED - Use manual Slide.UpdateAllConnectors() for explicit control

### Alternative 3: Visibility Graph for Elbow Routing

**Approach**: Use visibility graph algorithm for obstacle avoidance.

**Pros**:
- Theoretical optimality for shortest path
- Handles complex obstacle shapes

**Cons**:
- More complex than A* for orthogonal routing
- Overkill for manhattan distance (right-angle only)
- Harder to implement and debug
- A* is industry standard for grid-based routing

**Decision**: REJECTED - Use A* pathfinding (simpler, proven, sufficient)

### Alternative 4: Embed Routing Logic in ConnectionShape

**Approach**: Put all routing algorithms directly in ConnectionShape methods.

**Pros**:
- Everything in one place
- No separate routing package

**Cons**:
- Violates single responsibility principle
- Not extensible (can't add custom routers)
- Hard to test routing algorithms in isolation
- Couples connection management with routing logic

**Decision**: REJECTED - Use Router interface with separate implementations

## Detailed Component Design

**Design Note - Dependencies on Existing Shape API**: This proposal assumes the existing Shape element (from presentation-elements) already provides basic methods for position and size access:
- `Shape.Offset() (x, y EMU)` - Returns shape position
- `Shape.Extent() (width, height EMU)` - Returns shape size
- `Shape.ID() string` - Returns shape identifier

These methods are foundational to the existing Shape API and are used throughout connector routing for:
- Calculating absolute connection point positions
- Determining shape bounding boxes for obstacle avoidance
- Resolving shape references by ID

If these methods are not yet implemented, they should be added to the base presentation-elements spec as prerequisite work.

### 1. Connection Site System

**ConnectionSite Struct**:
```go
// ConnectionSite represents a point where connectors can attach to a shape.
// Corresponds to ECMA-376 CT_ConnectionSite (a:cxn element).
type ConnectionSite struct {
    // Position is the location relative to shape bounds
    // (0, 0) = top-left corner of shape
    Position struct {
        X EMU
        Y EMU
    }

    // Angle is the direction normal to the shape edge at this point
    // Measured in 1/60000 degrees (ECMA-376 standard)
    // 0 = right, 90*60000 = down, 180*60000 = left, 270*60000 = up
    Angle int32

    // Index is the 0-based identifier for this connection site
    // Used in connector start/end connection references
    Index int
}
```

**Shape Connection Site API**:
```go
// ConnectionSites returns all connection sites on this shape.
// For rectangles, returns default sites (top, right, bottom, left).
// For custom shapes, returns user-added sites.
func (s *Shape) ConnectionSites() []*ConnectionSite {
    // Look for a:cxnLst child element in shape properties
    // If not present, generate default sites for rectangle
    // Parse each a:cxn child element into ConnectionSite
}

// AddConnectionSite adds a custom connection site to the shape.
// x, y are in EMU relative to shape bounds.
// angle is in 1/60000 degrees (direction normal to edge).
func (s *Shape) AddConnectionSite(x, y EMU, angle int32) *ConnectionSite {
    // Get or create a:cxnLst element in shape properties
    // Create new a:cxn child with ang attribute
    // Append a:pos element with x, y attributes
    // Assign next available index
    // Return ConnectionSite struct
}

// GetConnectionSite retrieves a specific connection site by index.
// Returns nil if index out of range.
func (s *Shape) GetConnectionSite(index int) *ConnectionSite {
    sites := s.ConnectionSites()
    if index < 0 || index >= len(sites) {
        return nil
    }
    return sites[index]
}
```

**Default Connection Sites for Rectangles**:
```go
func generateDefaultConnectionSites(shapeWidth, shapeHeight EMU) []*ConnectionSite {
    return []*ConnectionSite{
        // Index 0: Top center (normal pointing up)
        {
            Position: struct{ X, Y EMU }{X: shapeWidth / 2, Y: 0},
            Angle:    270 * 60000,
            Index:    0,
        },

        // Index 1: Right center (normal pointing right)
        {
            Position: struct{ X, Y EMU }{X: shapeWidth, Y: shapeHeight / 2},
            Angle:    0,
            Index:    1,
        },

        // Index 2: Bottom center (normal pointing down)
        {
            Position: struct{ X, Y EMU }{X: shapeWidth / 2, Y: shapeHeight},
            Angle:    90 * 60000,
            Index:    2,
        },

        // Index 3: Left center (normal pointing left)
        {
            Position: struct{ X, Y EMU }{X: 0, Y: shapeHeight / 2},
            Angle:    180 * 60000,
            Index:    3,
        },
    }
}
```

### 2. ConnectionShape Connection API

**XML Structure** (ECMA-376):
```xml
<p:cxnSp>
  <p:nvCxnSpPr>
    <p:cNvPr id="1" name="Connector 1"/>
    <p:cNvCxnSpPr>
      <!-- Start connection -->
      <a:stCxn id="2" idx="0"/>  <!-- Connect to shape ID 2, site index 0 -->

      <!-- End connection -->
      <a:endCxn id="3" idx="2"/>  <!-- Connect to shape ID 3, site index 2 -->
    </p:cNvCxnSpPr>
    <p:nvPr/>
  </p:nvCxnSpPr>
  <p:spPr>
    <!-- Path geometry goes here -->
  </p:spPr>
</p:cxnSp>
```

**Connection API Implementation**:
```go
// StartConnection returns the start connection information.
// shapeId is the ID of the shape this connector starts at.
// connectionSiteIndex is the index of the connection site on that shape.
// ok is false if no start connection is set.
func (cs *ConnectionShape) StartConnection() (shapeId string, connectionSiteIndex int, ok bool) {
    // Navigate to nvCxnSpPr/cNvCxnSpPr element
    nvCxnSpPr := cs.GetElement("nvCxnSpPr", NamespacePresentationML)
    if nvCxnSpPr == nil {
        return "", 0, false
    }

    cNvCxnSpPr := nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML)
    if cNvCxnSpPr == nil {
        return "", 0, false
    }

    // Look for a:stCxn element
    stCxn := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML)
    if stCxn == nil {
        return "", 0, false
    }

    // Read id and idx attributes
    idAttr, found := stCxn.GetAttribute("id", "")
    if !found {
        return "", 0, false
    }

    idxAttr, found := stCxn.GetAttribute("idx", "")
    if !found {
        return "", 0, false
    }

    idx, _ := strconv.Atoi(idxAttr.Value())
    return idAttr.Value(), idx, true
}

// SetStartConnection sets the start connection.
// Creates or updates the a:stCxn element in cNvCxnSpPr.
func (cs *ConnectionShape) SetStartConnection(shapeId string, connectionSiteIndex int) {
    // Get or create nvCxnSpPr/cNvCxnSpPr structure
    nvCxnSpPr := cs.getOrCreateElement("nvCxnSpPr")
    cNvCxnSpPr := nvCxnSpPr.getOrCreateElement("cNvCxnSpPr")

    // Remove existing stCxn if present
    if existing := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML); existing != nil {
        cNvCxnSpPr.RemoveChild(existing)
    }

    // Create new a:stCxn element
    stCxn := openxml.NewCompositeElement(NamespaceDrawingML, "stCxn", "a")
    stCxn.SetAttribute(openxml.NewAttribute("", "id", "", shapeId))
    stCxn.SetAttribute(openxml.NewAttribute("", "idx", "", strconv.Itoa(connectionSiteIndex)))

    // Insert before endCxn if present, otherwise append
    if endCxn := cNvCxnSpPr.GetElement("endCxn", NamespaceDrawingML); endCxn != nil {
        cNvCxnSpPr.InsertBefore(stCxn, endCxn)
    } else {
        cNvCxnSpPr.AppendChild(stCxn)
    }
}

// EndConnection and SetEndConnection follow same pattern
// (omitted for brevity - identical to start connection but uses endCxn element)

// ClearStartConnection removes the start connection.
func (cs *ConnectionShape) ClearStartConnection() {
    nvCxnSpPr := cs.GetElement("nvCxnSpPr", NamespacePresentationML)
    if nvCxnSpPr == nil {
        return
    }

    cNvCxnSpPr := nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML)
    if cNvCxnSpPr == nil {
        return
    }

    if stCxn := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML); stCxn != nil {
        cNvCxnSpPr.RemoveChild(stCxn)
    }
}
```

### 2.1 RouterType Persistence

**Storage Mechanism**: RouterType is stored as a custom extension element within the ConnectionShape's non-visual properties to enable path recalculation after shape moves.

**Implementation**:
```go
// SetRouterType stores the router type for later recalculation
func (cs *ConnectionShape) SetRouterType(routerType RouterType) {
    nvCxnSpPr := cs.NonVisualConnectionShapeProperties()
    if nvCxnSpPr == nil {
        return  // Invalid ConnectionShape
    }

    nvPr := nvCxnSpPr.NonVisualDrawingProperties()
    if nvPr == nil {
        nvPr = openxml.NewCompositeElement(NamespacePresentation, "nvPr", "p")
        nvCxnSpPr.AppendChild(nvPr)
    }

    // Access or create extension list (p:extLst)
    extLst := nvPr.GetElement("extLst", NamespacePresentation)
    if extLst == nil {
        extLst = openxml.NewCompositeElement(NamespacePresentation, "extLst", "p")
        nvPr.AppendChild(extLst)
    }

    // Find existing router type extension or create new
    var ext openxml.Element
    for _, child := range extLst.Children() {
        if e, ok := child.(openxml.Element); ok {
            if uri := e.GetAttribute("uri", ""); uri != nil && uri.Value() == "http://goffice.dev/connector/routerType" {
                ext = e
                break
            }
        }
    }

    if ext == nil {
        ext = openxml.NewCompositeElement(NamespacePresentation, "ext", "p")
        ext.SetAttribute(openxml.NewAttribute("", "uri", "", "http://goffice.dev/connector/routerType"))
        extLst.AppendChild(ext)
    }

    // Store router type as text content
    ext.SetTextContent(routerType.String())
}

// RouterType retrieves the stored router type, defaulting to Straight if not present
func (cs *ConnectionShape) RouterType() RouterType {
    nvCxnSpPr := cs.NonVisualConnectionShapeProperties()
    if nvCxnSpPr == nil {
        return RouterTypeStraight  // Default
    }

    nvPr := nvCxnSpPr.NonVisualDrawingProperties()
    if nvPr == nil {
        return RouterTypeStraight  // Default
    }

    extLst := nvPr.GetElement("extLst", NamespacePresentation)
    if extLst == nil {
        return RouterTypeStraight  // Default
    }

    // Find router type extension
    for _, child := range extLst.Children() {
        if e, ok := child.(openxml.Element); ok {
            if uri := e.GetAttribute("uri", ""); uri != nil && uri.Value() == "http://goffice.dev/connector/routerType" {
                // Parse router type from text content
                switch e.TextContent() {
                case "straight":
                    return RouterTypeStraight
                case "elbow":
                    return RouterTypeElbow
                case "curved":
                    return RouterTypeCurved
                default:
                    return RouterTypeStraight
                }
            }
        }
    }

    return RouterTypeStraight  // Default if extension not found
}
```

**XML Structure with RouterType Extension**:
```xml
<p:cxnSp>
  <p:nvCxnSpPr>
    <p:cNvPr id="3" name="Connector 1"/>
    <p:cNvCxnSpPr>
      <a:stCxn id="2" idx="0"/>
      <a:endCxn id="4" idx="2"/>
    </p:cNvCxnSpPr>
    <p:nvPr>
      <p:extLst>
        <p:ext uri="http://goffice.dev/connector/routerType">elbow</p:ext>
      </p:extLst>
    </p:nvPr>
  </p:nvCxnSpPr>
  <p:spPr>...</p:spPr>
  <p:style>...</p:style>
</p:cxnSp>
```

**Design Rationale**:
- Uses standard ECMA-376 extension mechanism (p:extLst/p:ext) for forward compatibility
- Custom URI namespace "http://goffice.dev/connector/routerType" avoids conflicts with Microsoft Office extensions
- Simple text content ("straight", "elbow", "curved") is human-readable and easy to serialize/parse
- Defaults to RouterTypeStraight if extension is missing, providing graceful handling of legacy documents
- Extension elements are preserved by Office applications even when they don't understand them (per ECMA-376 spec)
- Storing in nvPr (non-visual properties) is semantically correct since routing type affects path calculation, not visual appearance directly

### 3. Routing Algorithm Architecture

**Router Interface**:
```go
// Router calculates a path between two points, optionally avoiding obstacles.
type Router interface {
    // Route calculates a path from start to end.
    // Returns the path as a sequence of drawing commands, or error if routing fails.
    Route(start, end Point, context *RoutingContext) (Path, error)
}

// RoutingContext provides information needed for intelligent routing.
type RoutingContext struct {
    // StartShape is the shape the connector starts from (for angle calculation)
    StartShape *Shape

    // EndShape is the shape the connector ends at (for angle calculation)
    EndShape *Shape

    // Obstacles are shapes to avoid during routing
    Obstacles []*Shape

    // SlideWidth and SlideHeight define the routing area bounds
    SlideWidth  EMU
    SlideHeight EMU

    // Margin is minimum distance to keep from obstacles (default: 10 points)
    Margin EMU
}

// Point represents a 2D coordinate in EMU.
type Point struct {
    X EMU
    Y EMU
}

// Path represents a sequence of drawing commands (MoveTo, LineTo, CurveTo).
type Path struct {
    Commands []PathCommand
}

type PathCommand interface {
    pathCommand() // marker interface
}

type MoveTo struct {
    X, Y EMU
}

type LineTo struct {
    X, Y EMU
}

type CubicBezierTo struct {
    CP1X, CP1Y EMU  // First control point
    CP2X, CP2Y EMU  // Second control point
    X, Y       EMU  // End point
}

func (m MoveTo) pathCommand()        {}
func (l LineTo) pathCommand()        {}
func (c CubicBezierTo) pathCommand() {}
```

### 4. Straight Line Router

**Implementation**:
```go
// StraightRouter creates a direct line from start to end.
// Ignores obstacles - simplest routing algorithm.
type StraightRouter struct{}

// Route creates a straight line path.
func (r *StraightRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    // Handle coincident points (zero-length line)
    if start.X == end.X && start.Y == end.Y {
        return Path{
            Commands: []PathCommand{
                MoveTo{X: start.X, Y: start.Y},
            },
        }, nil
    }

    return Path{
        Commands: []PathCommand{
            MoveTo{X: start.X, Y: start.Y},
            LineTo{X: end.X, Y: end.Y},
        },
    }, nil
}
```

### 5. Elbow Router (A* Orthogonal Routing)

**Grid-Based Pathfinding**:
```go
// ElbowRouter uses A* pathfinding on a grid for orthogonal routing.
// Creates paths with only horizontal and vertical segments (right-angle bends).
type ElbowRouter struct {
    // GridSize is the spacing between grid points (default: 10 points)
    GridSize EMU

    // Margin is minimum distance from obstacles (default: 10 points)
    Margin EMU
}

// Route calculates an orthogonal path avoiding obstacles.
func (r *ElbowRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    // 1. Create grid covering slide area
    grid := r.createGrid(ctx.SlideWidth, ctx.SlideHeight)

    // 2. Mark obstacle cells as blocked
    r.markObstacles(grid, ctx.Obstacles, ctx.Margin)

    // 3. Snap start and end to nearest grid points
    startCell := r.snapToGrid(start)
    endCell := r.snapToGrid(end)

    // 4. Run A* search
    path := r.astar(grid, startCell, endCell)
    if path == nil {
        return Path{}, errors.New("no path found")
    }

    // 5. Convert grid path to Path commands
    return r.gridPathToPath(path, start, end), nil
}

// Grid represents the routing grid.
type Grid struct {
    Width   int
    Height  int
    Cells   [][]GridCell
    CellSize EMU
}

type GridCell struct {
    X, Y    int
    Blocked bool
}

// PriorityQueue implements a min-heap priority queue for A* pathfinding.
// Items are stored with priority values, and Pop() always returns the item with minimum priority.
type PriorityQueue struct {
    items []pqItem
}

type pqItem struct {
    cell     GridCell
    priority int
}

func NewPriorityQueue() *PriorityQueue {
    return &PriorityQueue{
        items: make([]pqItem, 0),
    }
}

func (pq *PriorityQueue) Push(cell GridCell, priority int) {
    pq.items = append(pq.items, pqItem{cell: cell, priority: priority})
    pq.bubbleUp(len(pq.items) - 1)
}

func (pq *PriorityQueue) Pop() GridCell {
    if len(pq.items) == 0 {
        return GridCell{} // Return zero value if empty
    }

    result := pq.items[0].cell
    lastIdx := len(pq.items) - 1
    pq.items[0] = pq.items[lastIdx]
    pq.items = pq.items[:lastIdx]

    if len(pq.items) > 0 {
        pq.bubbleDown(0)
    }

    return result
}

func (pq *PriorityQueue) Empty() bool {
    return len(pq.items) == 0
}

func (pq *PriorityQueue) Contains(cell GridCell) bool {
    for _, item := range pq.items {
        if item.cell.X == cell.X && item.cell.Y == cell.Y {
            return true
        }
    }
    return false
}

func (pq *PriorityQueue) bubbleUp(idx int) {
    for idx > 0 {
        parentIdx := (idx - 1) / 2
        if pq.items[idx].priority >= pq.items[parentIdx].priority {
            break
        }
        pq.items[idx], pq.items[parentIdx] = pq.items[parentIdx], pq.items[idx]
        idx = parentIdx
    }
}

func (pq *PriorityQueue) bubbleDown(idx int) {
    for {
        leftIdx := 2*idx + 1
        rightIdx := 2*idx + 2
        smallest := idx

        if leftIdx < len(pq.items) && pq.items[leftIdx].priority < pq.items[smallest].priority {
            smallest = leftIdx
        }
        if rightIdx < len(pq.items) && pq.items[rightIdx].priority < pq.items[smallest].priority {
            smallest = rightIdx
        }

        if smallest == idx {
            break
        }

        pq.items[idx], pq.items[smallest] = pq.items[smallest], pq.items[idx]
        idx = smallest
    }
}

// createGrid builds a grid covering the slide area.
func (r *ElbowRouter) createGrid(width, height EMU) *Grid {
    cellSize := r.GridSize
    if cellSize == 0 {
        cellSize = 10 * EMUPerPoint  // Default 10 points
    }

    gridWidth := int(width / cellSize) + 1
    gridHeight := int(height / cellSize) + 1

    cells := make([][]GridCell, gridHeight)
    for y := 0; y < gridHeight; y++ {
        cells[y] = make([]GridCell, gridWidth)
        for x := 0; x < gridWidth; x++ {
            cells[y][x] = GridCell{X: x, Y: y, Blocked: false}
        }
    }

    return &Grid{
        Width:    gridWidth,
        Height:   gridHeight,
        Cells:    cells,
        CellSize: cellSize,
    }
}

// markObstacles marks grid cells covered by obstacle shapes as blocked.
func (r *ElbowRouter) markObstacles(grid *Grid, obstacles []*Shape, margin EMU) {
    for _, obstacle := range obstacles {
        // Get obstacle bounding box
        bounds := r.getShapeBounds(obstacle)

        // Expand by margin
        bounds.MinX -= margin
        bounds.MinY -= margin
        bounds.MaxX += margin
        bounds.MaxY += margin

        // Mark all covered cells as blocked
        minCellX := int(bounds.MinX / grid.CellSize)
        minCellY := int(bounds.MinY / grid.CellSize)
        maxCellX := int(bounds.MaxX / grid.CellSize)
        maxCellY := int(bounds.MaxY / grid.CellSize)

        for y := minCellY; y <= maxCellY && y < grid.Height; y++ {
            for x := minCellX; x <= maxCellX && x < grid.Width; x++ {
                if y >= 0 && x >= 0 {
                    grid.Cells[y][x].Blocked = true
                }
            }
        }
    }
}

// astar implements A* pathfinding on the grid.
func (r *ElbowRouter) astar(grid *Grid, start, end GridCell) []GridCell {
    // A* data structures
    openSet := NewPriorityQueue()
    closedSet := make(map[GridCell]bool) // Track explored nodes to avoid revisiting
    cameFrom := make(map[GridCell]GridCell)
    gScore := make(map[GridCell]int)
    fScore := make(map[GridCell]int)

    // Initialize start node
    gScore[start] = 0
    fScore[start] = r.heuristic(start, end)
    openSet.Push(start, fScore[start])

    for !openSet.Empty() {
        current := openSet.Pop()

        // Skip if already explored (prevents re-exploring nodes)
        if closedSet[current] {
            continue
        }

        // Mark current as explored
        closedSet[current] = true

        // Goal reached
        if current.X == end.X && current.Y == end.Y {
            return r.reconstructPath(cameFrom, current)
        }

        // Explore neighbors (up, right, down, left - orthogonal only)
        neighbors := r.getNeighbors(grid, current)
        for _, neighbor := range neighbors {
            // Skip blocked cells
            if grid.Cells[neighbor.Y][neighbor.X].Blocked {
                continue
            }

            // Skip already explored nodes
            if closedSet[neighbor] {
                continue
            }

            // Calculate tentative gScore
            tentativeGScore := gScore[current] + 1

            // If this path is better
            if prevGScore, exists := gScore[neighbor]; !exists || tentativeGScore < prevGScore {
                cameFrom[neighbor] = current
                gScore[neighbor] = tentativeGScore
                fScore[neighbor] = tentativeGScore + r.heuristic(neighbor, end)

                if !openSet.Contains(neighbor) {
                    openSet.Push(neighbor, fScore[neighbor])
                }
            }
        }
    }

    // No path found
    return nil
}

// heuristic calculates Manhattan distance (admissible for orthogonal routing).
func (r *ElbowRouter) heuristic(a, b GridCell) int {
    dx := abs(a.X - b.X)
    dy := abs(a.Y - b.Y)
    return dx + dy
}

// getNeighbors returns the 4 orthogonal neighbors of a cell.
func (r *ElbowRouter) getNeighbors(grid *Grid, cell GridCell) []GridCell {
    neighbors := []GridCell{}

    // Up
    if cell.Y > 0 {
        neighbors = append(neighbors, GridCell{X: cell.X, Y: cell.Y - 1})
    }

    // Right
    if cell.X < grid.Width-1 {
        neighbors = append(neighbors, GridCell{X: cell.X + 1, Y: cell.Y})
    }

    // Down
    if cell.Y < grid.Height-1 {
        neighbors = append(neighbors, GridCell{X: cell.X, Y: cell.Y + 1})
    }

    // Left
    if cell.X > 0 {
        neighbors = append(neighbors, GridCell{X: cell.X - 1, Y: cell.Y})
    }

    return neighbors
}

// reconstructPath builds the path from start to end using cameFrom map.
func (r *ElbowRouter) reconstructPath(cameFrom map[GridCell]GridCell, current GridCell) []GridCell {
    path := []GridCell{current}

    for {
        prev, exists := cameFrom[current]
        if !exists {
            break
        }
        path = append([]GridCell{prev}, path...)  // Prepend
        current = prev
    }

    return path
}

// gridPathToPath converts grid cells to Path commands.
func (r *ElbowRouter) gridPathToPath(gridPath []GridCell, realStart, realEnd Point) Path {
    if len(gridPath) == 0 {
        return Path{}
    }

    commands := []PathCommand{}

    // Start at real start point
    commands = append(commands, MoveTo{X: realStart.X, Y: realStart.Y})

    // Add lines through grid points
    for _, cell := range gridPath[1:] {  // Skip first cell (already at start)
        x := EMU(cell.X) * r.GridSize
        y := EMU(cell.Y) * r.GridSize
        commands = append(commands, LineTo{X: x, Y: y})
    }

    // End at real end point
    commands = append(commands, LineTo{X: realEnd.X, Y: realEnd.Y})

    return Path{Commands: commands}
}
```

### 6. Curved Router (Bezier)

**Smooth Curve Implementation**:
```go
// CurvedRouter creates smooth Bezier curves between connection points.
// Uses connection site angles to determine curve direction.
type CurvedRouter struct {
    // ControlPointDistance is how far from start/end to place control points
    // Default: 1/3 of total distance
    ControlPointDistance float64
}

// Route creates a cubic Bezier curve path.
func (r *CurvedRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    // Get connection site angles for tangent directions
    startAngle := r.getConnectionAngle(ctx.StartShape, start)
    endAngle := r.getConnectionAngle(ctx.EndShape, end)

    // Calculate control points based on angles
    distance := r.pointDistance(start, end)
    cpDistance := distance * r.getControlPointDistance()

    // Control point 1: start + direction from start angle
    cp1 := Point{
        X: start.X + EMU(float64(cpDistance)*math.Cos(float64(startAngle))),
        Y: start.Y + EMU(float64(cpDistance)*math.Sin(float64(startAngle))),
    }

    // Control point 2: end - direction from end angle
    cp2 := Point{
        X: end.X - EMU(float64(cpDistance)*math.Cos(float64(endAngle))),
        Y: end.Y - EMU(float64(cpDistance)*math.Sin(float64(endAngle))),
    }

    return Path{
        Commands: []PathCommand{
            MoveTo{X: start.X, Y: start.Y},
            CubicBezierTo{
                CP1X: cp1.X, CP1Y: cp1.Y,
                CP2X: cp2.X, CP2Y: cp2.Y,
                X:    end.X, Y: end.Y,
            },
        },
    }, nil
}

func (r *CurvedRouter) getControlPointDistance() float64 {
    if r.ControlPointDistance > 0 {
        return r.ControlPointDistance
    }
    return 0.33  // Default: 1/3 of distance
}

func (r *CurvedRouter) getConnectionAngle(shape *Shape, point Point) float64 {
    // Find nearest connection site to determine angle
    sites := shape.ConnectionSites()
    nearest := r.findNearestSite(sites, point)
    if nearest != nil {
        // Convert from 1/60000 degrees to radians
        // Formula: (angle / 60000) converts from 1/60000 degree units to degrees
        // Then multiply by π/180 to convert degrees to radians
        return float64(nearest.Angle) / 60000.0 * math.Pi / 180.0
    }

    // Fallback: calculate angle from shape center to point
    center := r.getShapeCenter(shape)
    return math.Atan2(float64(point.Y-center.Y), float64(point.X-center.X))
}

// findNearestSite finds the connection site closest to the given point.
func (r *CurvedRouter) findNearestSite(sites []*ConnectionSite, point Point) *ConnectionSite {
    if len(sites) == 0 {
        return nil
    }

    var nearest *ConnectionSite
    minDistance := float64(math.MaxInt64)

    for _, site := range sites {
        dist := r.pointDistance(Point{X: site.Position.X, Y: site.Position.Y}, point)
        if dist < minDistance {
            minDistance = dist
            nearest = site
        }
    }

    return nearest
}

// pointDistance calculates Euclidean distance between two points.
func (r *CurvedRouter) pointDistance(a, b Point) float64 {
    dx := float64(b.X - a.X)
    dy := float64(b.Y - a.Y)
    return math.Sqrt(dx*dx + dy*dy)
}

// getShapeCenter returns the center point of a shape.
func (r *CurvedRouter) getShapeCenter(shape *Shape) Point {
    offset := shape.Offset()
    extent := shape.Extent()
    return Point{
        X: offset.X + extent.Cx/2,
        Y: offset.Y + extent.Cy/2,
    }
}
```

### 7. High-Level Connection API

**Slide.ConnectShapes()**:
```go
// ConnectShapes creates a connector between two shapes with automatic routing.
// Finds the nearest connection sites and calculates the optimal path.
// Returns nil if either shape is nil (graceful handling of invalid input).
func (s *Slide) ConnectShapes(startShape, endShape *Shape, routerType RouterType) *ConnectionShape {
    // Validate input - return nil for nil shapes (graceful handling)
    if startShape == nil || endShape == nil {
        return nil
    }

    // 1. Create new ConnectionShape
    connector := NewConnectionShape()

    // 2. Find nearest connection sites
    startSite, endSite := s.findNearestSites(startShape, endShape)

    // 3. Set connections
    connector.SetStartConnection(startShape.ID(), startSite.Index)
    connector.SetEndConnection(endShape.ID(), endSite.Index)

    // 4. Calculate connection points in absolute coordinates
    startPoint := s.calculateAbsoluteConnectionPoint(startShape, startSite)
    endPoint := s.calculateAbsoluteConnectionPoint(endShape, endSite)

    // 5. Route path
    router := s.getRouter(routerType)
    obstacles := s.getObstacles(startShape, endShape)

    ctx := &RoutingContext{
        StartShape:  startShape,
        EndShape:    endShape,
        Obstacles:   obstacles,
        SlideWidth:  s.slideWidth,
        SlideHeight: s.slideHeight,
    }

    path, err := router.Route(startPoint, endPoint, ctx)
    if err != nil {
        // Fallback to straight line
        path = Path{
            Commands: []PathCommand{
                MoveTo{X: startPoint.X, Y: startPoint.Y},
                LineTo{X: endPoint.X, Y: endPoint.Y},
            },
        }
    }

    // 6. Set path geometry on connector
    s.setConnectorPath(connector, path)

    // 7. Persist router type for later recalculation (UpdateAllConnectors)
    connector.SetRouterType(routerType)

    // 8. Add to slide
    s.ShapeTree().AppendChild(connector)

    return connector
}

// findNearestSites finds the pair of connection sites that minimize distance.
func (s *Slide) findNearestSites(startShape, endShape *Shape) (*ConnectionSite, *ConnectionSite) {
    startSites := startShape.ConnectionSites()
    endSites := endShape.ConnectionSites()

    minDistance := math.MaxFloat64
    var bestStart, bestEnd *ConnectionSite

    for _, startSite := range startSites {
        startPoint := s.calculateAbsoluteConnectionPoint(startShape, startSite)

        for _, endSite := range endSites {
            endPoint := s.calculateAbsoluteConnectionPoint(endShape, endSite)

            distance := s.pointDistance(startPoint, endPoint)
            if distance < minDistance {
                minDistance = distance
                bestStart = startSite
                bestEnd = endSite
            }
        }
    }

    return bestStart, bestEnd
}

// UpdateAllConnectors recalculates paths for all connectors on the slide.
func (s *Slide) UpdateAllConnectors() error {
    connectors := s.ShapeTree().ConnectionShapes()

    for _, connector := range connectors {
        // Get start and end connections
        startShapeId, startSiteIdx, hasStart := connector.StartConnection()
        endShapeId, endSiteIdx, hasEnd := connector.EndConnection()

        if !hasStart || !hasEnd {
            continue  // Skip unconnected connectors
        }

        // Resolve shape references
        startShape := s.findShapeById(startShapeId)
        endShape := s.findShapeById(endShapeId)

        if startShape == nil || endShape == nil {
            continue  // Shape deleted
        }

        // Get connection sites
        startSite := startShape.GetConnectionSite(startSiteIdx)
        endSite := endShape.GetConnectionSite(endSiteIdx)

        if startSite == nil || endSite == nil {
            continue  // Invalid site index
        }

        // Calculate connection points
        startPoint := s.calculateAbsoluteConnectionPoint(startShape, startSite)
        endPoint := s.calculateAbsoluteConnectionPoint(endShape, endSite)

        // Re-route using the connector's stored router type
        routerType := connector.RouterType()  // Retrieve stored router type
        router := s.getRouter(routerType)      // Get corresponding router implementation
        ctx := &RoutingContext{
            StartShape:  startShape,
            EndShape:    endShape,
            Obstacles:   s.getObstacles(startShape, endShape),
            SlideWidth:  s.slideWidth,
            SlideHeight: s.slideHeight,
        }

        path, err := router.Route(startPoint, endPoint, ctx)
        if err != nil {
            continue  // Keep old path on routing failure
        }

        // Update connector path
        s.setConnectorPath(connector, path)
    }

    return nil
}

// getRouter returns the router implementation for the given type.
func (s *Slide) getRouter(routerType RouterType) Router {
    switch routerType {
    case RouterTypeStraight:
        return &StraightRouter{}
    case RouterTypeElbow:
        return &ElbowRouter{Margin: 10000}  // Default 10000 EMU margin
    case RouterTypeCurved:
        return &CurvedRouter{ControlPointDistance: 0.33}
    default:
        return &StraightRouter{}  // Fallback
    }
}

// getObstacles returns all shapes on the slide except the two being connected.
func (s *Slide) getObstacles(startShape, endShape *Shape) []*Shape {
    obstacles := []*Shape{}
    for _, shape := range s.ShapeTree().Shapes() {
        if shape != startShape && shape != endShape {
            obstacles = append(obstacles, shape)
        }
    }
    return obstacles
}

// findShapeById searches for a shape by its ID.
func (s *Slide) findShapeById(shapeId string) *Shape {
    for _, shape := range s.ShapeTree().Shapes() {
        if shape.ID() == shapeId {
            return shape
        }
    }
    return nil
}

// calculateAbsoluteConnectionPoint calculates the absolute position of a connection site.
func (s *Slide) calculateAbsoluteConnectionPoint(shape *Shape, site *ConnectionSite) Point {
    offset := shape.Offset()
    return Point{
        X: offset.X + site.Position.X,
        Y: offset.Y + site.Position.Y,
    }
}

// setConnectorPath updates the connector's path geometry.
func (s *Slide) setConnectorPath(connector *ConnectionShape, path Path) {
    // Update the connector's spPr/a:custGeom or a:prstGeom with path data
    // Implementation details depend on how ConnectionShape stores geometry
    // For now, stub showing the concept
    spPr := connector.ShapeProperties()
    if spPr == nil {
        return
    }
    // TODO: Convert Path to DrawingML geometry (a:custGeom with a:pathLst)
}
```

### 7.1 Advanced Path Methods (Phase 1 Stubs - Full Implementation in Future)

**Path Validation**:
```go
// IsValid checks if the path has valid command sequence.
func (p *Path) IsValid() bool {
    if len(p.Commands) == 0 {
        return false  // Empty path is invalid
    }
    // First command must be MoveTo
    if _, ok := p.Commands[0].(MoveTo); !ok {
        return false
    }
    return true
}
```

**Path Transformation** (Stub - basic implementation):
```go
// Transform applies a transformation matrix to all path commands.
// Full matrix transformation will be implemented in Phase 2.
func (p *Path) Transform(matrix Matrix) {
    for i := range p.Commands {
        switch cmd := p.Commands[i].(type) {
        case MoveTo:
            x, y := matrix.TransformPoint(cmd.X, cmd.Y)
            p.Commands[i] = MoveTo{X: x, Y: y}
        case LineTo:
            x, y := matrix.TransformPoint(cmd.X, cmd.Y)
            p.Commands[i] = LineTo{X: x, Y: y}
        case CubicBezierTo:
            cp1x, cp1y := matrix.TransformPoint(cmd.CP1X, cmd.CP1Y)
            cp2x, cp2y := matrix.TransformPoint(cmd.CP2X, cmd.CP2Y)
            x, y := matrix.TransformPoint(cmd.X, cmd.Y)
            p.Commands[i] = CubicBezierTo{
                CP1X: cp1x, CP1Y: cp1y,
                CP2X: cp2x, CP2Y: cp2y,
                X:    x, Y:    y,
            }
        }
    }
}
```

**Path Reversal** (Stub):
```go
// Reverse reverses the path direction.
// Full implementation with proper Bezier control point reversal in Phase 2.
func (p *Path) Reverse() {
    if len(p.Commands) < 2 {
        return
    }
    // Simple reversal - swap start and end points
    // TODO: Proper Bezier curve control point reversal
    reversed := make([]PathCommand, len(p.Commands))
    for i := range p.Commands {
        reversed[len(p.Commands)-1-i] = p.Commands[i]
    }
    p.Commands = reversed
}
```

**Bezier Utilities** (Stubs - for future advanced curve manipulation):
```go
// EvaluateBezier calculates point on curve at parameter t (0..1).
// Stub for Phase 2 advanced curve manipulation.
func (p *Path) EvaluateBezier(cmdIndex int, t float64) (EMU, EMU) {
    // TODO: Full De Casteljau's algorithm implementation
    return 0, 0
}

// SubdivideBezier splits Bezier curve at parameter t.
// Stub for Phase 2.
func (p *Path) SubdivideBezier(cmdIndex int, t float64) (Path, Path) {
    // TODO: De Casteljau subdivision
    return Path{}, Path{}
}

// TangentAtBezier calculates tangent vector at parameter t.
// Stub for Phase 2.
func (p *Path) TangentAtBezier(cmdIndex int, t float64) (float64, float64) {
    // TODO: Bezier derivative calculation
    return 0, 0
}
```

**Design Note**: The methods above are basic implementations sufficient for Phase 1 connector routing. Full-featured path manipulation (e.g., precise Bezier subdivision, accurate tangent calculation) will be implemented in Phase 2 as needed for advanced connector features like curved arrow heads and path editing.

### 8. PDF Rendering

**ConnectorRenderer**:
```go
// ConnectorRenderer renders ConnectionShape elements to PDF.
type ConnectorRenderer struct {
    ctx *core.RenderingContext
}

func (r *ConnectorRenderer) Render(connector *ConnectionShape) error {
    // 1. Get connector properties
    spPr := connector.ShapeProperties()
    if spPr == nil {
        return nil  // No visual properties
    }

    // 2. Get line style
    ln := spPr.Line()
    if ln != nil {
        r.applyLineStyle(ln)
    }

    // 3. Get path geometry
    custGeom := spPr.CustomGeometry()
    if custGeom == nil {
        return nil  // No path
    }

    // 4. Render path
    path := custGeom.PathList()
    if path != nil {
        r.renderPath(path)
    }

    // 5. Render arrow decorations
    if ln != nil {
        if headEnd := ln.HeadEnd(); headEnd != nil {
            r.renderArrowHead(headEnd, path, true)  // Start arrow
        }
        if tailEnd := ln.TailEnd(); tailEnd != nil {
            r.renderArrowHead(tailEnd, path, false)  // End arrow
        }
    }

    return nil
}

func (r *ConnectorRenderer) renderArrowHead(arrow *drawingml.LineEndProperties, path *Path, isStart bool) {
    // Get arrow type (triangle, arrow, diamond, etc.)
    arrowType := arrow.Type()

    // Get arrow size
    arrowWidth := arrow.Width()
    arrowLength := arrow.Length()

    // Calculate arrow position and angle from path
    var point Point
    var angle float64

    if isStart {
        point = path.StartPoint()
        angle = path.StartAngle()
    } else {
        point = path.EndPoint()
        angle = path.EndAngle()
    }

    // Draw arrow based on type
    switch arrowType {
    case ArrowTypeTriangle:
        r.renderTriangleArrow(point, angle, arrowWidth, arrowLength)
    case ArrowTypeArrow:
        r.renderArrow(point, angle, arrowWidth, arrowLength)
    case ArrowTypeDiamond:
        r.renderDiamondArrow(point, angle, arrowWidth, arrowLength)
    }
}
```

## Path Geometry Utilities

**Path struct enhancements**:
```go
// AddLine appends a line segment to the path.
func (p *Path) AddLine(x, y EMU) {
    p.Commands = append(p.Commands, LineTo{X: x, Y: y})
}

// AddCubicBezier appends a cubic Bezier curve to the path.
func (p *Path) AddCubicBezier(cp1x, cp1y, cp2x, cp2y, x, y EMU) {
    p.Commands = append(p.Commands, CubicBezierTo{
        CP1X: cp1x, CP1Y: cp1y,
        CP2X: cp2x, CP2Y: cp2y,
        X: x, Y: y,
    })
}

// Bounds calculates the bounding box of the path.
func (p *Path) Bounds() (minX, minY, maxX, maxY EMU) {
    if len(p.Commands) == 0 {
        return 0, 0, 0, 0
    }

    minX, minY = math.MaxInt64, math.MaxInt64
    maxX, maxY = math.MinInt64, math.MinInt64

    for _, cmd := range p.Commands {
        switch c := cmd.(type) {
        case MoveTo:
            minX = min(minX, c.X)
            minY = min(minY, c.Y)
            maxX = max(maxX, c.X)
            maxY = max(maxY, c.Y)

        case LineTo:
            minX = min(minX, c.X)
            minY = min(minY, c.Y)
            maxX = max(maxX, c.X)
            maxY = max(maxY, c.Y)

        case CubicBezierTo:
            // Include control points and end point
            minX = min(minX, min(c.CP1X, min(c.CP2X, c.X)))
            minY = min(minY, min(c.CP1Y, min(c.CP2Y, c.Y)))
            maxX = max(maxX, max(c.CP1X, max(c.CP2X, c.X)))
            maxY = max(maxY, max(c.CP1Y, max(c.CP2Y, c.Y)))
        }
    }

    return minX, minY, maxX, maxY
}

// Length calculates the approximate length of the path.
func (p *Path) Length() float64 {
    if len(p.Commands) == 0 {
        return 0
    }

    totalLength := 0.0
    var currentX, currentY EMU

    for _, cmd := range p.Commands {
        switch c := cmd.(type) {
        case MoveTo:
            currentX, currentY = c.X, c.Y

        case LineTo:
            dx := float64(c.X - currentX)
            dy := float64(c.Y - currentY)
            totalLength += math.Sqrt(dx*dx + dy*dy)
            currentX, currentY = c.X, c.Y

        case CubicBezierTo:
            // Approximate Bezier length with linear segments
            length := p.approximateBezierLength(
                currentX, currentY,
                c.CP1X, c.CP1Y,
                c.CP2X, c.CP2Y,
                c.X, c.Y,
            )
            totalLength += length
            currentX, currentY = c.X, c.Y
        }
    }

    return totalLength
}

// approximateBezierLength uses recursive subdivision to estimate curve length.
func (p *Path) approximateBezierLength(x0, y0, x1, y1, x2, y2, x3, y3 EMU) float64 {
    // Calculate chord length (straight line from start to end)
    dx := float64(x3 - x0)
    dy := float64(y3 - y0)
    chord := math.Sqrt(dx*dx + dy*dy)

    // Calculate control polygon length
    dx01 := float64(x1 - x0)
    dy01 := float64(y1 - y0)
    dx12 := float64(x2 - x1)
    dy12 := float64(y2 - y1)
    dx23 := float64(x3 - x2)
    dy23 := float64(y3 - y2)

    controlLength := math.Sqrt(dx01*dx01 + dy01*dy01) +
                    math.Sqrt(dx12*dx12 + dy12*dy12) +
                    math.Sqrt(dx23*dx23 + dy23*dy23)

    // If close enough, use average
    // Threshold of 0.5 EMU ensures <1 EMU total error per spec requirement
    // At typical recursion depths (6-8 levels), this produces sub-EMU accuracy
    // 914400 EMU = 1 inch, so 0.5 EMU ≈ 0.0000005 inches (sub-pixel accuracy)
    if math.Abs(controlLength - chord) < 0.5 {
        return (controlLength + chord) / 2
    }

    // Otherwise, subdivide and recurse
    // De Casteljau's algorithm for subdivision at t=0.5
    midX01 := (x0 + x1) / 2
    midY01 := (y0 + y1) / 2
    midX12 := (x1 + x2) / 2
    midY12 := (y1 + y2) / 2
    midX23 := (x2 + x3) / 2
    midY23 := (y2 + y3) / 2

    midX012 := (midX01 + midX12) / 2
    midY012 := (midY01 + midY12) / 2
    midX123 := (midX12 + midX23) / 2
    midY123 := (midY12 + midY23) / 2

    midX := (midX012 + midX123) / 2
    midY := (midY012 + midY123) / 2

    // Recurse on two halves
    length1 := p.approximateBezierLength(x0, y0, midX01, midY01, midX012, midY012, midX, midY)
    length2 := p.approximateBezierLength(midX, midY, midX123, midY123, midX23, midY23, x3, y3)

    return length1 + length2
}
```

## Testing Strategy

**Unit Tests**:
```go
// Test connection site management
func TestShape_ConnectionSites(t *testing.T) {
    shape := NewShape()

    // Default sites for rectangle
    sites := shape.ConnectionSites()
    assert.Equal(t, 4, len(sites))  // Top, right, bottom, left

    // Custom site
    custom := shape.AddConnectionSite(100*EMUPerPoint, 50*EMUPerPoint, 45*60000)
    assert.Equal(t, 4, custom.Index)  // Index 4 (after default sites)

    sites = shape.ConnectionSites()
    assert.Equal(t, 5, len(sites))  // 4 default + 1 custom
}

// Test connection set/get
func TestConnectionShape_Connection(t *testing.T) {
    connector := NewConnectionShape()

    // Initially no connections
    _, _, ok := connector.StartConnection()
    assert.False(t, ok)

    // Set connection
    connector.SetStartConnection("shape123", 2)

    // Get connection
    shapeId, siteIdx, ok := connector.StartConnection()
    assert.True(t, ok)
    assert.Equal(t, "shape123", shapeId)
    assert.Equal(t, 2, siteIdx)

    // Clear connection
    connector.ClearStartConnection()
    _, _, ok = connector.StartConnection()
    assert.False(t, ok)
}

// Test straight routing
func TestStraightRouter(t *testing.T) {
    router := &StraightRouter{}

    start := Point{X: 100 * EMUPerPoint, Y: 100 * EMUPerPoint}
    end := Point{X: 300 * EMUPerPoint, Y: 200 * EMUPerPoint}

    path, err := router.Route(start, end, &RoutingContext{})
    assert.NoError(t, err)
    assert.Equal(t, 2, len(path.Commands))

    moveTo, ok := path.Commands[0].(MoveTo)
    assert.True(t, ok)
    assert.Equal(t, start.X, moveTo.X)

    lineTo, ok := path.Commands[1].(LineTo)
    assert.True(t, ok)
    assert.Equal(t, end.X, lineTo.X)
}

// Test elbow routing with obstacles
func TestElbowRouter_WithObstacles(t *testing.T) {
    router := &ElbowRouter{GridSize: 10 * EMUPerPoint}

    start := Point{X: 50 * EMUPerPoint, Y: 50 * EMUPerPoint}
    end := Point{X: 150 * EMUPerPoint, Y: 150 * EMUPerPoint}

    // Obstacle in direct path
    obstacle := NewShape()
    obstacle.SetPosition(90*EMUPerPoint, 90*EMUPerPoint, 40*EMUPerPoint, 40*EMUPerPoint)

    ctx := &RoutingContext{
        Obstacles:   []*Shape{obstacle},
        SlideWidth:  200 * EMUPerPoint,
        SlideHeight: 200 * EMUPerPoint,
        Margin:      5 * EMUPerPoint,
    }

    path, err := router.Route(start, end, ctx)
    assert.NoError(t, err)

    // Path should go around obstacle (more than 2 segments)
    assert.Greater(t, len(path.Commands), 2)

    // Verify path avoids obstacle bounds
    for _, cmd := range path.Commands {
        switch c := cmd.(type) {
        case LineTo:
            assert.False(t, isPointInBounds(c.X, c.Y, obstacle))
        }
    }
}
```

**Integration Tests**:
```go
func TestFlowchartCreation(t *testing.T) {
    pres := presentation.Create()
    slide := pres.AddSlide()

    // Create decision tree
    //     [Start]
    //        ↓
    //    [Decision]
    //     ↙    ↘
    //  [Yes]  [No]

    start := slide.AddShape()
    start.SetText("Start")
    start.SetPosition(150*EMUPerPoint, 50*EMUPerPoint, 80*EMUPerPoint, 40*EMUPerPoint)

    decision := slide.AddShape()
    decision.SetText("Decision?")
    decision.SetPosition(140*EMUPerPoint, 120*EMUPerPoint, 100*EMUPerPoint, 40*EMUPerPoint)

    yes := slide.AddShape()
    yes.SetText("Yes")
    yes.SetPosition(80*EMUPerPoint, 190*EMUPerPoint, 60*EMUPerPoint, 40*EMUPerPoint)

    no := slide.AddShape()
    no.SetText("No")
    no.SetPosition(240*EMUPerPoint, 190*EMUPerPoint, 60*EMUPerPoint, 40*EMUPerPoint)

    // Connect with elbow routing
    c1 := slide.ConnectShapes(start, decision, RouterTypeElbow)
    c2 := slide.ConnectShapes(decision, yes, RouterTypeElbow)
    c3 := slide.ConnectShapes(decision, no, RouterTypeElbow)

    // Verify connections
    startId1, _, _ := c1.StartConnection()
    assert.Equal(t, start.ID(), startId1)

    endId1, _, _ := c1.EndConnection()
    assert.Equal(t, decision.ID(), endId1)

    // Save and verify
    filename := "testdata/flowchart-test.pptx"
    err := pres.SaveAs(filename)
    assert.NoError(t, err)

    // Reopen and verify structure preserved
    pres2, err := presentation.Open(filename)
    assert.NoError(t, err)

    slide2 := pres2.Slides()[0]
    connectors := slide2.ShapeTree().ConnectionShapes()
    assert.Equal(t, 3, len(connectors))
}
```

## Performance Considerations

**Grid Size vs Routing Quality**:
- Smaller grid = more precise routing, slower performance
- Larger grid = faster routing, less optimal paths
- Default 10 points balances quality and speed
- User can override via ElbowRouter.GridSize

**A* Performance**:
- Complexity: O(b^d) where b=branching factor (4 for orthogonal), d=depth
- For typical slide (10x7.5 inches at 10pt grid): ~100x75 = 7500 cells
- Typical path depth: 50-100 cells
- Expected runtime: <10ms for single path

**Caching Strategy**:
- Don't cache routed paths (shapes move frequently)
- Do cache obstacle grids (reuse for multiple connectors)
- Invalidate grid on UpdateAllConnectors()

## Future Enhancements

**Phase 2 Improvements** (out of current scope):
- Connection site auto-generation for complex shapes (circles, polygons)
- Global routing optimization (minimize crossings)
- Connector glue points (Office 2013+ feature)
- Curved elbow routing (rounded corners instead of sharp bends)
- Multi-segment path editing (user-adjustable waypoints)
- Smart connector relayout on shape resize (not just move)
