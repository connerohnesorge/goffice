# Implementation Tasks

## Phase 1: GroupShape API Enhancement

- [ ] 1.1 Add child accessor methods to GroupShape (Shapes, Pictures, GroupShapes, ConnectionShapes, GraphicFrames)
- [ ] 1.2 Add child manipulation methods to GroupShape (AddShape, AddPicture, AddGroupShape, AddConnectionShape, RemoveChild, Clear)
- [ ] 1.3 Add property accessor methods to GroupShape (GroupShapeProperties, NonVisualGroupShapeProperties)
- [ ] 1.4 Add Transform accessor to GroupShapeProperties (Transform, SetTransform methods)

## Phase 2: Transform Utilities

- [ ] 2.1 Create transform package with Matrix type and basic constructors (Identity, Translate, Rotate, Scale)
- [ ] 2.2 Implement matrix operations (Multiply, Invert, TransformPoint)
- [ ] 2.3 Implement Transform2D conversion (FromTransform2D function)
- [ ] 2.4 Implement hierarchical transform composition (ComputeAbsoluteTransform, LocalToAbsolute, AbsoluteToLocal)
- [ ] 2.5 Implement transform decomposition (DecomposeTransform extracts translation, rotation, scale)

## Phase 3: PDF Group Rendering

- [ ] 3.1 Create GroupRenderer struct in pdf/drawing/group_renderer.go
- [ ] 3.2 Implement basic group rendering logic (SaveGraphicsState, apply transform, render children, RestoreGraphicsState)
- [ ] 3.3 Add recursive group rendering support (handle nested GroupShapes)
- [ ] 3.4 Integrate GroupRenderer into presentation renderer (pdf/presentation/renderer.go)
- [ ] 3.5 Add error handling and edge cases (nil transforms, empty groups, invalid children)

## Phase 4: High-Level Grouping API

- [ ] 4.1 Implement bounding box calculation utility (computeBoundingBox function)
- [ ] 4.2 Implement slide.CreateGroup() method (groups shapes, adjusts coordinates, adds to slide)
- [ ] 4.3 Implement coordinate adjustment for grouped shapes (adjustCoordinatesRelativeToGroup utility)
- [ ] 4.4 Implement slide.UngroupShape() method (flattens group, preserves absolute positions)
- [ ] 4.5 Add roundtrip test for group/ungroup operations (CreateGroup → UngroupShape preserves positions)

## Phase 5: Testing & Documentation

- [ ] 5.1 Create test documents with groups (simple-group, nested-3-levels, rotated-group, flipped-group, mixed-shapes-group, complex-org-chart)
- [ ] 5.2 Add integration tests with real documents (presentation/integration_test.go)
- [ ] 5.3 Add PDF rendering visual validation tests (automated pixel comparison with PowerPoint reference)
- [ ] 5.4 Write examples demonstrating grouping (create-group, nested-groups, transform-group, pdf-group-rendering)
- [ ] 5.5 Write API documentation (godoc comments for all public APIs)
- [ ] 5.6 Update README with grouping capabilities section

## Validation Checkpoints

- [ ] After Phase 1: All GroupShape API tests pass (go test ./presentation/elements -v)
- [ ] After Phase 2: All transformation tests pass (go test ./drawingml/transform -v)
- [ ] After Phase 3: All PDF rendering tests pass, sample PDFs generated (go test ./pdf/... -v)
- [ ] After Phase 4: CreateGroup/UngroupShape tests pass (go test ./presentation -v)
- [ ] After Phase 5: All tests pass, examples compile, documentation complete (go test ./... -v)

## Dependencies

**Sequential**:
- Phase 1 → Phase 4 (CreateGroup needs GroupShape API)
- Phase 2 → Phase 3 (PDF rendering needs transform package)
- Phase 2 → Phase 4 (UngroupShape needs transform utilities)
- Phase 3 → Phase 5 (testing needs all implementations)

**Parallel opportunities**:
- Phase 1 and Phase 2 (independent)
- Tasks 1.1, 1.3, 1.4 within Phase 1
- Tasks 2.1, 2.5 within Phase 2 (2.5 only depends on 2.2)
- Tasks 5.1, 5.4 (document creation)