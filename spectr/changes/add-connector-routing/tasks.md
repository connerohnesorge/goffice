# Implementation Tasks

## Phase 1: Connection Site System

- [ ] 1.1 Add ConnectionSite struct to drawingml package (position, angle, index)
- [ ] 1.2 Add Shape.ConnectionSites() method to enumerate connection sites
- [ ] 1.3 Add Shape.AddConnectionSite() method to add custom connection sites
- [ ] 1.4 Add Shape.GetConnectionSite() method to retrieve specific site by index
- [ ] 1.5 Implement default connection site generation for rectangles (top, right, bottom, left)
- [ ] 1.6 Add unit tests for connection site management

## Phase 2: ConnectionShape Connection API

- [ ] 2.1 Add ConnectionShape.StartConnection() method to get start connection info
- [ ] 2.2 Add ConnectionShape.SetStartConnection() method to connect to shape
- [ ] 2.3 Add ConnectionShape.EndConnection() method to get end connection info
- [ ] 2.4 Add ConnectionShape.SetEndConnection() method to connect to shape
- [ ] 2.5 Add ConnectionShape.ClearStartConnection() and ClearEndConnection() methods
- [ ] 2.6 Implement cNvCxnSpPr/stCxn/endCxn XML manipulation
- [ ] 2.7 Add unit tests for connection set/get/clear operations

## Phase 3: Routing Algorithm Infrastructure

- [ ] 3.1 Create presentation/routing/ package
- [ ] 3.2 Define Router interface (Route method)
- [ ] 3.3 Define RoutingContext struct (shapes, obstacles, slide bounds)
- [ ] 3.4 Define RouterType enum (Straight, Elbow, Curved)
- [ ] 3.5 Create Path struct utilities (AddLine, AddCubicBezier, Bounds, Length)

## Phase 4: Straight Line Router

- [ ] 4.1 Implement StraightRouter struct
- [ ] 4.2 Implement Route() method (direct line from start to end)
- [ ] 4.3 Add unit tests for straight routing
- [ ] 4.4 Add integration test creating simple connector

## Phase 5: Elbow Router (Orthogonal)

- [ ] 5.1 Implement ElbowRouter struct with A* pathfinding
- [ ] 5.2 Implement grid generation from slide bounds
- [ ] 5.3 Implement obstacle marking on grid (shape bounding boxes)
- [ ] 5.4 Implement Manhattan distance heuristic
- [ ] 5.5 Implement A* search algorithm
- [ ] 5.6 Implement path reconstruction from A* result
- [ ] 5.7 Add unit tests for elbow routing (with and without obstacles)
- [ ] 5.8 Add integration test creating flowchart with elbow connectors

## Phase 6: Curved Router (Bezier)

- [ ] 6.1 Implement CurvedRouter struct
- [ ] 6.2 Implement control point calculation based on connection site angles
- [ ] 6.3 Implement cubic Bezier curve generation
- [ ] 6.4 Add unit tests for curved routing
- [ ] 6.5 Add integration test creating org chart with curved connectors

## Phase 7: High-Level Connection API

- [ ] 7.1 Add Slide.ConnectShapes() method
- [ ] 7.2 Implement automatic connection site selection (nearest sites)
- [ ] 7.3 Implement router selection and path calculation
- [ ] 7.4 Add Slide.UpdateAllConnectors() method
- [ ] 7.5 Implement connector discovery (find all connectors on slide)
- [ ] 7.6 Implement path recalculation for all connectors
- [ ] 7.7 Add Slide.GetConnectedShapes() method
- [ ] 7.8 Add unit tests for high-level connection API

## Phase 8: PDF Rendering

- [ ] 8.1 Create pdf/drawing/connector_renderer.go
- [ ] 8.2 Implement ConnectorRenderer struct
- [ ] 8.3 Implement shape reference resolution (lookup by ID)
- [ ] 8.4 Implement connection point position calculation
- [ ] 8.5 Implement routed path rendering
- [ ] 8.6 Add arrow head rendering (start/end decorations)
- [ ] 8.7 Add line style support (solid, dashed, dotted)
- [ ] 8.8 Integrate connector renderer into presentation renderer
- [ ] 8.9 Add PDF rendering tests with visual validation

## Phase 9: Testing & Documentation

- [ ] 9.1 Create test documents (simple-connector, flowchart, org-chart, network-diagram)
- [ ] 9.2 Add roundtrip tests (create → save → open → verify)
- [ ] 9.3 Add integration test for shape move + connector update
- [ ] 9.4 Write examples/connectors/create-flowchart.go
- [ ] 9.5 Write examples/connectors/org-chart-generator.go
- [ ] 9.6 Write examples/connectors/network-topology.go
- [ ] 9.7 Write examples/connectors/connector-styles.go
- [ ] 9.8 Write API documentation (godoc for all public APIs)
- [ ] 9.9 Update README with connector capabilities section

## Validation Checkpoints

- [ ] After Phase 2: Connection set/get operations work (go test ./presentation/elements -v)
- [ ] After Phase 4: Straight routing creates valid paths (go test ./presentation/routing -v)
- [ ] After Phase 5: Elbow routing avoids obstacles (go test ./presentation/routing -v)
- [ ] After Phase 6: Curved routing creates smooth paths (go test ./presentation/routing -v)
- [ ] After Phase 8: PDF rendering shows connectors correctly (visual inspection)
- [ ] After Phase 9: All tests pass, examples compile (go test ./... -v && go build ./examples/connectors/...)

## Dependencies

**Sequential**:
- Phase 1 → Phase 2 (connection sites needed for connections)
- Phase 3 → Phase 4, 5, 6 (router interface needed for implementations)
- Phase 2, 3 → Phase 7 (connection API and routers needed for high-level API)
- Phase 7 → Phase 8 (routing needed for PDF rendering)

**Parallel opportunities**:
- Phase 4, 5, 6 can be developed in parallel (independent router implementations)
- Phase 9 test documents can be created while Phase 8 is in development
