# Change: Add Complete Connector Shape Routing and Management

## Why

ConnectionShape elements exist in goffice but are essentially non-functional shells with no routing, connection management, or intelligent path calculation:

**Current State**:
- ConnectionShape exists in presentation/elements/shape_tree.go with only basic construction
- No connection point management (where connectors attach to shapes)
- No routing algorithms (straight, elbow/orthogonal, curved)
- No automatic path recalculation when shapes move
- PDF rendering treats connectors as static paths
- No API to programmatically connect shapes

**Impact**:
- **Flowcharts/Org Charts**: Cannot create diagrams with smart connectors
- **Network Diagrams**: No way to show relationships between entities
- **Process Flows**: Manual connector manipulation required
- **Data Visualization**: Cannot programmatically generate connected diagrams
- **User Productivity**: Must manually adjust every connector when shapes move

**Evidence from Codebase**:
- `presentation/elements/shape_tree.go:543-610` - ConnectionShape has NewConnectionShape() only
- No methods for setting start/end connections
- No routing algorithm implementations
- PDF rendering ignores connection semantics (pdf/drawing/)

## What Changes

### 1. Connection Point System (presentation-elements capability)

**Add to Shape**:
- `ConnectionSites() []*ConnectionSite` - Get all connection points on shape
- `AddConnectionSite(x, y EMU, angle int32) *ConnectionSite` - Add custom connection point
- `GetConnectionSite(index int) *ConnectionSite` - Get specific connection site by index

**ConnectionSite struct** (from ECMA-376 CT_ConnectionSite):
- Position (x, y) in shape local coordinates
- Angle (direction normal to shape edge, in 1/60000 degrees)
- Index (0-based identifier for connection site)

**Default Connection Sites** (auto-generated for rectangles):
- Index 0: Top center
- Index 1: Right center
- Index 2: Bottom center
- Index 3: Left center

### 2. ConnectionShape Connection API (presentation-elements capability)

**Add to ConnectionShape**:
- `StartConnection() (shapeId string, connectionSiteIndex int, ok bool)` - Get start connection
- `SetStartConnection(shapeId string, connectionSiteIndex int)` - Connect start to shape
- `EndConnection() (shapeId string, connectionSiteIndex int, ok bool)` - Get end connection
- `SetEndConnection(shapeId string, connectionSiteIndex int)` - Connect end to shape
- `ClearStartConnection()` - Disconnect start
- `ClearEndConnection()` - Disconnect end

**Implementation Note**: These manipulate the cNvCxnSpPr element's stCxn and endCxn child elements per ECMA-376 CT_ConnectionShapeProperties.

### 3. Routing Algorithms (presentation-routing capability)

**New package**: `presentation/routing/`

**Router interface**:
```go
type Router interface {
    Route(start, end Point, context *RoutingContext) (Path, error)
}

type RoutingContext struct {
    StartShape  *Shape
    EndShape    *Shape
    Obstacles   []*Shape  // shapes to avoid
    SlideWidth  EMU
    SlideHeight EMU
}
```

**Straight Line Router**:
```go
type StraightRouter struct{}

func (r *StraightRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    return Path{
        MoveTo(start),
        LineTo(end),
    }, nil
}
```

**Elbow Router** (Orthogonal/Manhattan routing):
```go
type ElbowRouter struct {
    Margin EMU  // minimum distance from obstacles
}

func (r *ElbowRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    // A* pathfinding with Manhattan distance heuristic
    // Grid-based routing avoiding obstacles
    // Minimize turns + total distance
    // Returns orthogonal path with right-angle bends
}
```

**Curved Router** (Smooth Bezier):
```go
type CurvedRouter struct{}

func (r *CurvedRouter) Route(start, end Point, ctx *RoutingContext) (Path, error) {
    // Cubic Bezier curve
    // Control points based on connection site angles
    // Smooth S-curve or C-curve depending on positions
}
```

### 4. High-Level Connection API (presentation-document capability)

**Add to Slide**:
- `ConnectShapes(start *Shape, end *Shape, routerType RouterType) *ConnectionShape` - Create and route connector
- `UpdateAllConnectors()` - Recalculate all connector paths (call when shapes move)
- `GetConnectedShapes(shape *Shape) []*ConnectionShape` - Find all connectors touching a shape

**RouterType enum**:
- `RouterTypeStraight` - Direct line
- `RouterTypeElbow` - Orthogonal/Manhattan routing
- `RouterTypeCurved` - Bezier curve

**Example Usage**:
```go
// Create shapes
shape1 := slide.AddShape()
shape1.SetPosition(100*drawingml.EMUPerPoint, 100*drawingml.EMUPerPoint)

shape2 := slide.AddShape()
shape2.SetPosition(300*drawingml.EMUPerPoint, 200*drawingml.EMUPerPoint)

// Connect with elbow routing
connector := slide.ConnectShapes(shape1, shape2, presentation.RouterTypeElbow)

// Move shape - connector path auto-updates
shape2.SetPosition(400*drawingml.EMUPerPoint, 300*drawingml.EMUPerPoint)
slide.UpdateAllConnectors()
```

### 5. PDF Rendering (pdf-drawing capability)

**New file**: `pdf/drawing/connector_renderer.go`

**ConnectorRenderer**:
- Resolve start/end shape references
- Calculate connection point positions
- Render routed path
- Support arrow decorations (start/end arrow heads)
- Support line styles (solid, dashed, dotted)

### 6. Path Geometry Utilities (drawingml-core capability)

**Add to drawingml/path.go**:
- `Path.AddLine(x, y EMU)` - Append line segment
- `Path.AddCubicBezier(cp1x, cp1y, cp2x, cp2y, x, y EMU)` - Append Bezier curve
- `Path.Bounds() (minX, minY, maxX, maxY EMU)` - Calculate bounding box
- `Path.Length() float64` - Calculate total path length

### 7. Testing & Documentation

**Unit Tests**:
- Connection site management (add, get, default sites)
- Connection set/get/clear operations
- Straight routing (simple line)
- Elbow routing (A* pathfinding, obstacle avoidance)
- Curved routing (Bezier control points)

**Integration Tests**:
- Create flowchart with 5 connected shapes
- Move shapes and verify connectors update
- PDF rendering with arrow decorations
- Roundtrip: create → save → open → verify connections preserved

**Test Documents** (in `testdata/`):
- `simple-connector.pptx`: Two rectangles with straight connector
- `flowchart.pptx`: Decision tree with elbow connectors
- `org-chart.pptx`: Hierarchical diagram with curved connectors
- `network-diagram.pptx`: Complex graph with multiple routing types

**Examples** (in `examples/connectors/`):
- `create-flowchart.go`: Build process flow diagram
- `org-chart-generator.go`: Generate organizational hierarchy
- `network-topology.go`: Create network diagram
- `connector-styles.go`: Demonstrate line styles and arrows

## Impact

**Affected specs**:
- `presentation-elements` (MODIFIED - add connection site and connection APIs)
- `presentation-routing` (ADDED - routing algorithms)
- `presentation-document` (ADDED - high-level connection API)
- `pdf-drawing` (MODIFIED - connector rendering)
- `drawingml-core` (MODIFIED - path geometry utilities)

**New capabilities**:
- Programmatically connect shapes with intelligent routing
- Automatic path recalculation when shapes move
- Multiple routing algorithms (straight, elbow, curved)
- Connection site management
- PDF rendering with proper connector semantics
- Arrow decorations and line styles

**Breaking changes**: None. All changes are additive.

## Key Design Decisions

### 1. Router Interface for Extensibility

**Decision**: Use Router interface to support multiple algorithms.

**Rationale**:
- Different diagram types need different routing (flowcharts use elbow, org charts use curved)
- Allows third-party custom routers
- Clean separation of routing logic from connection management

**Alternative considered**: Hard-code routing in ConnectionShape
- **Rejected**: Not extensible, violates single responsibility

### 2. Connection Sites vs Connection Points

**Decision**: Use ECMA-376 terminology "Connection Site" (not "Connection Point").

**Rationale**:
- Matches Office Open XML standard exactly (CT_ConnectionSite)
- Connection site includes position AND angle (direction normal)
- Avoids confusion with geometric "points"

### 3. A* Pathfinding for Elbow Routing

**Decision**: Implement A* algorithm for orthogonal routing.

**Rationale**:
- Industry standard for graph pathfinding
- Guarantees shortest path given heuristic
- Handles obstacle avoidance naturally
- Performance: O(b^d) where b=branching factor, d=depth

**Alternative considered**: Visibility graph
- **Rejected**: More complex, overkill for orthogonal routing

### 4. Slide-Level UpdateAllConnectors()

**Decision**: Manual connector update, not automatic on shape move.

**Rationale**:
- Performance: Avoid recalculating on every SetPosition() call
- User control: Batch updates after moving multiple shapes
- Matches Office behavior (connectors update on mouse release, not drag)

**Alternative considered**: Automatic update on shape move
- **Rejected**: Performance penalty, unexpected side effects

### 5. Shape ID References for Connections

**Decision**: Store shape ID strings in connection elements.

**Rationale**:
- Matches ECMA-376 CT_Connection (uses shape ID attribute)
- Handles shape deletion gracefully (ID lookup fails, connector orphaned)
- Survives copy/paste operations

**Alternative considered**: Direct shape pointers
- **Rejected**: Breaks on serialization, fragile to shape deletion

## Dependencies

**Existing**:
- ConnectionShape element (presentation/elements/shape_tree.go) - exists but empty
- ShapeProperties (drawingml/) - exists
- PDF Page API (pdf/core/) - exists

**New**:
- `presentation/routing/` package - routing algorithms
- `pdf/drawing/connector_renderer.go` - connector rendering

**Schema compatibility**:
- No schema changes (cxnSp is standard ECMA-376)
- Uses existing cNvCxnSpPr/stCxn/endCxn elements

## Success Criteria

- [ ] Shapes have connection sites (default + custom)
- [ ] ConnectionShape can set/get start/end connections
- [ ] Straight routing creates direct line path
- [ ] Elbow routing creates orthogonal path avoiding obstacles
- [ ] Curved routing creates smooth Bezier path
- [ ] slide.UpdateAllConnectors() recalculates all paths
- [ ] PDF rendering shows routed connector with arrows
- [ ] Roundtrip preserves connections and routing

## Out of Scope

- Connection site auto-generation for complex shapes (only rectangles in Phase 1)
- Global router settings (per-connector routing type only)
- Connector text labels (defer to separate proposal)
- 3D connectors (defer to 3D shapes proposal)
- Animated routing (defer to animation proposal)
- Collision detection during routing (simple bounding box only)
- Connector glue points (Office 2013+ feature, defer to Phase 2)
- Rerouting optimization (minimize crossings - Phase 2)

## References

- ECMA-376 Part 1, Section 21.3.2.36 (CT_ConnectionShape)
- ECMA-376 Part 1, Section 20.1.2.2.3 (CT_ConnectionSite)
- ECMA-376 Part 1, Section 20.1.9.12 (CT_Connection)
- A* Pathfinding: Hart, Peter E.; Nilsson, Nils J.; Raphael, Bertram (1968)
- Open-XML-SDK: DocumentFormat.OpenXml.Presentation.ConnectionShape
