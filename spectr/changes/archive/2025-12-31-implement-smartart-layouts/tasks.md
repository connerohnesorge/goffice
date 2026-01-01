# Implementation Tasks

## Phase 1: Layout Engine Foundation (Week 1-2)

### 1.1 Core Interfaces and Types
- [ ] 1.1.1 Create `presentation/smartart/layout/engine.go` with LayoutEngine interface
- [ ] 1.1.2 Define LayoutType enum (Hierarchy, List, Cycle, Pyramid, Matrix)
- [ ] 1.1.3 Create Rectangle, Position, Point helper types
- [ ] 1.1.4 Add unit tests for type definitions

### 1.2 Constraint System
- [ ] 1.2.1 Create `presentation/smartart/constraints/spacing.go` with Spacing struct
- [ ] 1.2.2 Create `presentation/smartart/constraints/sizing.go` with Sizing struct
- [ ] 1.2.3 Create `presentation/smartart/constraints/alignment.go` with Alignment types
- [ ] 1.2.4 Define DefaultConstraints with reasonable defaults
- [ ] 1.2.5 Add constraint validation logic
- [ ] 1.2.6 Add unit tests for constraints

### 1.3 Layout Engine Registry
- [ ] 1.3.1 Create `presentation/smartart/layout/registry.go` with Registry struct
- [ ] 1.3.2 Implement Register() method
- [ ] 1.3.3 Implement GetEngine() method
- [ ] 1.3.4 Create defaultRegistry singleton
- [ ] 1.3.5 Add unit tests for registry

### 1.4 SmartArt Extensions
- [ ] 1.4.1 Extend SmartArt type in `presentation/elements/smartart.go`
- [ ] 1.4.2 Add Layout() method to SmartArt
- [ ] 1.4.3 Add SetLayoutType() method
- [ ] 1.4.4 Add GetLayoutEngine() method
- [ ] 1.4.5 Add AutoLayout() enable/disable method
- [ ] 1.4.6 Add layout caching logic (layoutCached flag, layoutVersion)
- [ ] 1.4.7 Add unit tests for SmartArt extensions

### 1.5 Text Measurement
- [ ] 1.5.1 Create `presentation/smartart/layout/text_measurer.go`
- [ ] 1.5.2 Implement TextMeasurer struct using PDF font system
- [ ] 1.5.3 Implement MeasureText() method
- [ ] 1.5.4 Implement AutoSizeShape() method
- [ ] 1.5.5 Add unit tests with various fonts and sizes
- [ ] 1.5.6 Add benchmark tests for performance

## Phase 2: Hierarchy Layout (Week 3-4)

### 2.1 Tree Data Structure
- [ ] 2.1.1 Create `presentation/smartart/layout/hierarchy.go`
- [ ] 2.1.2 Define TreeNode struct (node, children, parent, position, size)
- [ ] 2.1.3 Implement buildTree() to convert diagram data to tree
- [ ] 2.1.4 Handle assistant nodes (dotted line subordinates)
- [ ] 2.1.5 Add tree validation (detect cycles, orphaned nodes)
- [ ] 2.1.6 Add unit tests for tree construction

### 2.2 Reingold-Tilford Algorithm
- [ ] 2.2.1 Implement HierarchyLayout struct
- [ ] 2.2.2 Implement calculatePositions() using Reingold-Tilford algorithm
- [ ] 2.2.3 Implement ensureSeparation() for sibling spacing
- [ ] 2.2.4 Implement shiftSubtree() for collision avoidance
- [ ] 2.2.5 Implement positionAssistants() for assistant nodes
- [ ] 2.2.6 Add unit tests for positioning algorithm

### 2.3 Orientation Support
- [ ] 2.3.1 Add orientation field (vertical, horizontal)
- [ ] 2.3.2 Implement applyOrientation() to rotate tree 90 degrees
- [ ] 2.3.3 Add unit tests for both orientations

### 2.4 Scaling and Centering
- [ ] 2.4.1 Implement scaleAndCenter() to fit tree in bounds
- [ ] 2.4.2 Handle bounds overflow (shrink tree if needed)
- [ ] 2.4.3 Add unit tests for scaling logic

### 2.5 Position Application
- [ ] 2.5.1 Implement applyPositions() to update diagram shapes
- [ ] 2.5.2 Implement routeConnectors() for parent-child lines
- [ ] 2.5.3 Add special connector rendering for assistants (dotted lines)
- [ ] 2.5.4 Add unit tests for position application

### 2.6 Integration Tests
- [ ] 2.6.1 Test simple 3-node tree
- [ ] 2.6.2 Test complex multi-level tree (5+ levels)
- [ ] 2.6.3 Test unbalanced tree (deep on one side)
- [ ] 2.6.4 Test tree with assistant nodes
- [ ] 2.6.5 Test horizontal orientation
- [ ] 2.6.6 Test auto-sizing based on text content

## Phase 3: List Layout (Week 5)

### 3.1 List Layout Implementation
- [ ] 3.1.1 Create `presentation/smartart/layout/list.go`
- [ ] 3.1.2 Implement ListLayout struct (orientation, spacing, alignment)
- [ ] 3.1.3 Implement calculatePositions() for vertical list
- [ ] 3.1.4 Implement calculatePositions() for horizontal list
- [ ] 3.1.5 Implement alignX() and alignY() helper methods
- [ ] 3.1.6 Add unit tests for list positioning

### 3.2 Connector Routing
- [ ] 3.2.1 Implement routeConnectors() for straight-line connections
- [ ] 3.2.2 Handle vertical flow (bottom → top connections)
- [ ] 3.2.3 Handle horizontal flow (right → left connections)
- [ ] 3.2.4 Add unit tests for connector routing

### 3.3 Integration Tests
- [ ] 3.3.1 Test vertical list with 5 items
- [ ] 3.3.2 Test horizontal list with 5 items
- [ ] 3.3.3 Test left/center/right alignment (vertical)
- [ ] 3.3.4 Test top/middle/bottom alignment (horizontal)
- [ ] 3.3.5 Test auto-sizing to container bounds

## Phase 4: Cycle Layout (Week 6)

### 4.1 Cycle Layout Implementation
- [ ] 4.1.1 Create `presentation/smartart/layout/cycle.go`
- [ ] 4.1.2 Implement CycleLayout struct (direction, centerText)
- [ ] 4.1.3 Implement calculateCircle() to determine radius and center
- [ ] 4.1.4 Implement positionOnCircle() for equal angular spacing
- [ ] 4.1.5 Add unit tests for circular positioning

### 4.2 Curved Arrows
- [ ] 4.2.1 Implement addArrows() for directional curved connectors
- [ ] 4.2.2 Calculate bezier control points for smooth curves
- [ ] 4.2.3 Add arrow heads at destination
- [ ] 4.2.4 Handle clockwise vs counter-clockwise direction
- [ ] 4.2.5 Add unit tests for arrow generation

### 4.3 Center Text
- [ ] 4.3.1 Implement addCenterText() for optional center label
- [ ] 4.3.2 Position text at circle center
- [ ] 4.3.3 Add unit tests for center text

### 4.4 Integration Tests
- [ ] 4.4.1 Test 4-item cycle (clockwise)
- [ ] 4.4.2 Test 6-item cycle (counter-clockwise)
- [ ] 4.4.3 Test with center text
- [ ] 4.4.4 Test radius calculation for various bounds
- [ ] 4.4.5 Test arrow rendering

## Phase 5: Pyramid & Matrix Layouts (Week 7)

### 5.1 Pyramid Layout Implementation
- [ ] 5.1.1 Create `presentation/smartart/layout/pyramid.go`
- [ ] 5.1.2 Implement PyramidLayout struct (orientation)
- [ ] 5.1.3 Implement calculateLevelSizes() for proportional sizing
- [ ] 5.1.4 Implement positionLevels() for stacked layout
- [ ] 5.1.5 Handle normal vs inverted orientation
- [ ] 5.1.6 Add unit tests for pyramid layout

### 5.2 Pyramid Integration Tests
- [ ] 5.2.1 Test 3-level pyramid (normal)
- [ ] 5.2.2 Test 5-level pyramid (inverted)
- [ ] 5.2.3 Test multiple items per level
- [ ] 5.2.4 Test proportional sizing

### 5.3 Matrix Layout Implementation
- [ ] 5.3.1 Create `presentation/smartart/layout/matrix.go`
- [ ] 5.3.2 Implement MatrixLayout struct (rows, columns)
- [ ] 5.3.3 Implement determineGridSize() for auto-sizing
- [ ] 5.3.4 Implement grid positioning logic
- [ ] 5.3.5 Add unit tests for matrix layout

### 5.4 Matrix Integration Tests
- [ ] 5.4.1 Test 2x2 matrix
- [ ] 5.4.2 Test 3x3 matrix
- [ ] 5.4.3 Test auto-determined grid size
- [ ] 5.4.4 Test cell centering

## Phase 6: PDF Rendering Integration (Week 8)

### 6.1 Diagram Renderer Updates
- [ ] 6.1.1 Modify `pdf/drawing/diagram_renderer.go`
- [ ] 6.1.2 Remove placeholder rendering logic
- [ ] 6.1.3 Implement RenderDiagram() with positioned shapes
- [ ] 6.1.4 Implement renderShape() for individual shapes
- [ ] 6.1.5 Implement renderConnector() for connectors
- [ ] 6.1.6 Handle straight-line connectors
- [ ] 6.1.7 Handle curved connectors (bezier)
- [ ] 6.1.8 Add arrow rendering

### 6.2 Shape Rendering
- [ ] 6.2.1 Draw shape background (fill color)
- [ ] 6.2.2 Draw shape border (stroke)
- [ ] 6.2.3 Draw shape text (centered)
- [ ] 6.2.4 Support various shape geometries (rectangle, ellipse, trapezoid)
- [ ] 6.2.5 Add unit tests for shape rendering

### 6.3 Connector Rendering
- [ ] 6.3.1 Calculate connection points on shapes
- [ ] 6.3.2 Draw straight lines
- [ ] 6.3.3 Draw bezier curves
- [ ] 6.3.4 Draw arrow heads
- [ ] 6.3.5 Support dotted lines (for assistant nodes)
- [ ] 6.3.6 Add unit tests for connector rendering

### 6.4 PDF Integration Tests
- [ ] 6.4.1 Test PDF rendering for hierarchy layout
- [ ] 6.4.2 Test PDF rendering for list layout
- [ ] 6.4.3 Test PDF rendering for cycle layout
- [ ] 6.4.4 Test PDF rendering for pyramid layout
- [ ] 6.4.5 Test PDF rendering for matrix layout
- [ ] 6.4.6 Visual comparison with reference PDFs

## Phase 7: High-Level API (Week 9)

### 7.1 SmartArt Creation API
- [ ] 7.1.1 Create `presentation/smartart.go`
- [ ] 7.1.2 Implement NewSmartArt(layoutType) function
- [ ] 7.1.3 Implement SmartArtFromData(data, layoutType) function
- [ ] 7.1.4 Define SmartArtData struct (items, connections)
- [ ] 7.1.5 Add unit tests for creation API

### 7.2 Data Model Helpers
- [ ] 7.2.1 Implement SmartArt.AddPoint(text) method
- [ ] 7.2.2 Implement SmartArt.AddConnection(src, dst, type) method
- [ ] 7.2.3 Implement SmartArt.RemovePoint(id) method
- [ ] 7.2.4 Implement SmartArt.RemoveConnection(id) method
- [ ] 7.2.5 Trigger layout invalidation on data changes
- [ ] 7.2.6 Add unit tests for data manipulation

### 7.3 API Integration Tests
- [ ] 7.3.1 Test NewSmartArt() for all layout types
- [ ] 7.3.2 Test SmartArtFromData() with complex data
- [ ] 7.3.3 Test data manipulation with auto-layout
- [ ] 7.3.4 Test layout invalidation and re-layout

## Phase 8: Testing & Documentation (Week 9-10)

### 8.1 Visual Regression Tests
- [ ] 8.1.1 Create test fixtures for all layout types
- [ ] 8.1.2 Generate reference images for comparison
- [ ] 8.1.3 Implement image comparison utility
- [ ] 8.1.4 Add visual regression tests for hierarchy layout
- [ ] 8.1.5 Add visual regression tests for list layout
- [ ] 8.1.6 Add visual regression tests for cycle layout
- [ ] 8.1.7 Add visual regression tests for pyramid layout
- [ ] 8.1.8 Add visual regression tests for matrix layout

### 8.2 Roundtrip Tests
- [ ] 8.2.1 Test create → layout → save → load → verify positions
- [ ] 8.2.2 Test layout preservation across save/load
- [ ] 8.2.3 Test auto-layout toggle persistence
- [ ] 8.2.4 Test layout type change roundtrip

### 8.3 Performance Tests
- [ ] 8.3.1 Benchmark hierarchy layout with 100 nodes
- [ ] 8.3.2 Benchmark list layout with 100 items
- [ ] 8.3.3 Benchmark cycle layout with 50 items
- [ ] 8.3.4 Benchmark layout caching effectiveness
- [ ] 8.3.5 Ensure layout operations complete in <100ms

### 8.4 Documentation
- [ ] 8.4.1 Write godoc for all public types and methods
- [ ] 8.4.2 Create `presentation/smartart/layout/doc.go` package documentation
- [ ] 8.4.3 Create `presentation/smartart/doc.go` API documentation
- [ ] 8.4.4 Document constraint system usage
- [ ] 8.4.5 Document extensibility (how to add new layouts)

### 8.5 Examples
- [ ] 8.5.1 Create example: Simple org chart
- [ ] 8.5.2 Create example: Process flow (list)
- [ ] 8.5.3 Create example: Circular process (cycle)
- [ ] 8.5.4 Create example: Hierarchy pyramid
- [ ] 8.5.5 Create example: Feature matrix
- [ ] 8.5.6 Create example: Custom layout configuration
- [ ] 8.5.7 Create example: SmartArt to PDF rendering

### 8.6 Integration with Existing Code
- [ ] 8.6.1 Update `CHANGELOG.md` with SmartArt layout support
- [ ] 8.6.2 Ensure no breaking changes to Phase 1 code
- [ ] 8.6.3 Verify all existing tests still pass
- [ ] 8.6.4 Run golangci-lint and fix any issues
- [ ] 8.6.5 Verify examples from Phase 1 still work

## Phase 9: Validation & Polish (Week 10)

### 9.1 Office Compatibility Testing
- [ ] 9.1.1 Create SmartArt diagrams in PowerPoint and compare layouts
- [ ] 9.1.2 Test roundtrip: goffice → PowerPoint → goffice
- [ ] 9.1.3 Verify layout fidelity with Office output
- [ ] 9.1.4 Document known differences from Office

### 9.2 Error Handling
- [ ] 9.2.1 Add validation for invalid layout types
- [ ] 9.2.2 Add validation for empty diagrams
- [ ] 9.2.3 Add validation for circular dependencies in hierarchy
- [ ] 9.2.4 Add graceful degradation for unsupported features
- [ ] 9.2.5 Add informative error messages

### 9.3 Edge Cases
- [ ] 9.3.1 Test single-node diagrams
- [ ] 9.3.2 Test diagrams with no connections
- [ ] 9.3.3 Test diagrams with very long text
- [ ] 9.3.4 Test diagrams with zero-size bounds
- [ ] 9.3.5 Test layout with extreme aspect ratios

### 9.4 Code Review & Cleanup
- [ ] 9.4.1 Review all code for clarity and consistency
- [ ] 9.4.2 Remove debug code and TODOs
- [ ] 9.4.3 Ensure consistent naming conventions
- [ ] 9.4.4 Optimize hot paths identified in benchmarks
- [ ] 9.4.5 Final golangci-lint pass

### 9.5 Release Preparation
- [ ] 9.5.1 Verify all tasks in this list are complete
- [ ] 9.5.2 Update all documentation
- [ ] 9.5.3 Create migration guide for users
- [ ] 9.5.4 Prepare release notes
- [ ] 9.5.5 Tag release version

## Notes

- **Dependencies**: Phase 1 foundation must be complete before starting Phase 2
- **Parallelization**: Phases 3-5 (List, Cycle, Pyramid/Matrix) can be implemented in parallel after Phase 2
- **Testing**: Write tests concurrently with implementation, not after
- **Documentation**: Write godoc as you implement, not at the end
- **Validation**: Run `spectr validate implement-smartart-layouts` before marking proposal complete
