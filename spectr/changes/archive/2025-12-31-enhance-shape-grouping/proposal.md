# Change: Enhance Shape Grouping Support

## Why

While `GroupShape` elements exist in the codebase (`presentation/elements/shape_tree.go`, `wordprocessing/elements/elements.go`), they lack the functionality needed for practical use:

**Current State**:
- `GroupShape` type exists but has no methods to manipulate child shapes
- No API to add shapes to a group or remove shapes from a group
- No iteration over child shapes within a group
- No recursive group traversal (groups within groups)
- PDF rendering completely ignores group shapes
- No coordinate transformation utilities for group hierarchies

**Impact**:
- **PowerPoint**: Documents with grouped shapes can be read but groups cannot be manipulated programmatically. PDF rendering breaks (shapes render in wrong positions because group transforms are ignored).
- **Word**: Inline diagram groups cannot be accessed or modified.
- **Excel**: Chart annotation groups are inaccessible.
- **User Experience**: Users

 cannot create, edit, or properly render documents with grouped shapes, forcing manual workarounds in Office applications.

**Evidence from Codebase**:
- `presentation/elements/shape_tree.go:320`: `GroupShape` struct exists but only has `NewGroupShape()` and `Clone()` methods
- `presentation/elements/shape_tree.go:120-135`: `GroupShapes()` returns groups but no way to access shapes *within* a group
- `pdf/` package: No references to `GroupShape` or `grpSp` (verified via grep) - PDF rendering does not handle groups
- `pdf/drawing/transform_renderer.go`: Handles individual shape transforms but not group-level coordinate spaces

## What Changes

### 1. GroupShape API Enhancement (presentation-elements capability)

**Add methods to GroupShape**:
- `Shapes() []*Shape` - Return all child shapes
- `Pictures() []*Picture` - Return all child pictures
- `GroupShapes() []*GroupShape` - Return nested child groups
- `ConnectionShapes() []*ConnectionShape` - Return child connectors
- `GraphicFrames() []*GraphicFrame` - Return child graphic frames
- `AddShape() *Shape` - Add a new shape to the group
- `AddPicture(relId string) *Picture` - Add a picture to the group
- `AddGroupShape() *GroupShape` - Add a nested group
- `AddConnectionShape() *ConnectionShape` - Add a connector to the group
- `RemoveChild(child openxml.Element)` - Remove a shape from the group
- `Clear()` - Remove all children
- `GroupShapeProperties() *GroupShapeProperties` - Access group-level properties (returns existing or nil)
- `NonVisualGroupShapeProperties() *NonVisualGroupShapeProperties` - Access non-visual properties (returns existing or nil)

**Design Note**: These convenience methods are a goffice enhancement. Microsoft's OpenXML-SDK uses `ChildElements` property with LINQ queries instead of typed accessors. goffice provides these methods for ergonomic API design, consistent with existing `ShapeTree.Shapes()` pattern.

**Add to GroupShapeProperties**:
- `Transform() *drawingml.Transform2D` - Get/set group-level transformation (returns existing or nil)
- `SetTransform(xfrm *drawingml.Transform2D)` - Set group transformation

**Add to Transform2D** (for group support):
- `ChildOffset() (x, y EMU, present bool)` - Get child coordinate space origin (a:chOff element). Returns (0, 0, false) if not present.
- `SetChildOffset(x, y EMU)` - Set child offset (creates a:chOff element if needed)
- `ChildExtent() (cx, cy EMU, present bool)` - Get child coordinate space size (a:chExt element). Returns (0, 0, false) if not present.
- `SetChildExtent(cx, cy EMU)` - Set child extent (creates a:chExt element if needed)

**Design Note**: These methods return a `present` boolean flag to distinguish between "not present" (returns false) and "present but zero" (returns true with 0, 0). This matches Go's map lookup pattern and avoids pointer allocations while maintaining API consistency with existing Offset()/Extent() methods that return value types.

**Note**: Groups use the same Transform2D element type as regular shapes (CT_Transform2D wraps CT_GroupTransform2D in ECMA-376). For groups, Transform2D includes additional child elements:
- `Offset` (a:off) - group position
- `Extent` (a:ext) - group size
- `ChildOffset` (a:chOff) - origin of child coordinate space within group (optional, defaults to group offset)
- `ChildExtent` (a:chExt) - size of child coordinate space (optional, defines viewport scaling)
- `Rotation`, `FlipH`, `FlipV` attributes - standard transform properties

### 2. Transform Composition Utilities (drawingml-core capability)

**New package**: `drawingml/transform/`

**Coordinate Transformation**:
- `ComputeAbsoluteTransform(shape Element, root Element) Matrix` - Traverse parent chain, compose transforms
- `LocalToAbsolute(point Point, shape Element, root Element) Point` - Convert local coordinates to absolute
- `AbsoluteToLocal(point Point, shape Element, root Element) Point` - Convert absolute to local coordinates
- `DecomposeTransform(matrix Matrix) (translation, rotation, scale)` - Extract components from matrix

**Matrix Operations**:
- `MultiplyMatrices(m1, m2 Matrix) Matrix` - Matrix multiplication
- `InvertMatrix(m Matrix) Matrix` - Matrix inversion
- `TransformPoint(point Point, matrix Matrix) Point` - Apply matrix to point

### 3. PDF Group Rendering (pdf-drawing capability)

**New file**: `pdf/drawing/group_renderer.go`

**Rendering Strategy**:
1. Enter group: Push PDF graphics state (`q` operator)
2. Apply group-level transform to PDF CTM (Current Transformation Matrix)
3. Recursively render child shapes (they inherit group's coordinate space)
4. Exit group: Pop PDF graphics state (`Q` operator)

**Implementation**:
```go
type GroupRenderer struct {
    ctx *core.RenderingContext
}

func (r *GroupRenderer) RenderGroup(group *GroupShape) error {
    // Push graphics state
    r.ctx.Page.SaveGraphicsState()
    defer r.ctx.Page.RestoreGraphicsState()

    // Apply group transformation
    if props := group.GroupShapeProperties(); props != nil {
        if xfrm := props.Transform(); xfrm != nil {
            matrix := transformToMatrix(xfrm)
            r.ctx.Page.Transform(matrix)
        }
    }

    // Render all children
    for child := range group.Children() {
        switch c := child.(type) {
        case *Shape:
            shapeRenderer.Render(c)
        case *Picture:
            pictureRenderer.Render(c)
        case *GroupShape:
            r.RenderGroup(c) // Recursive!
        case *ConnectionShape:
            connectorRenderer.Render(c)
        }
    }

    return nil
}
```

### 4. High-Level Grouping Operations (presentation-document capability)

**Add to Slide**:
- `CreateGroup(shapes ...interface{}) *GroupShape` - Group existing shapes (removes them from slide, adds to new group, adds group to slide)
- `UngroupShape(group *GroupShape) []interface{}` - Flatten group (removes group, adds children back to slide with absolute transforms)

**Example Usage**:
```go
// Create individual shapes
rect1 := slide.ShapeTree().AddShape()
rect2 := slide.ShapeTree().AddShape()
circle := slide.ShapeTree().AddShape()

// Group them
group := slide.CreateGroup(rect1, rect2, circle)

// Transform the entire group
groupProps := group.GroupShapeProperties()
xfrm := groupProps.Transform()
xfrm.SetRotation(45 * 60000) // Rotate 45 degrees (rotation stored as attribute on xfrm)

// Add another shape to the group
rect3 := group.AddShape()

// Ungroup later
shapes := slide.UngroupShape(group)
```

### 5. Testing & Documentation

**Unit Tests**:
- GroupShape method tests (add/remove/iterate children)
- Transform composition tests (multiply 3 transforms, verify result)
- Coordinate conversion tests (local → absolute for 3-level nesting)
- Matrix inversion tests

**Integration Tests**:
- Roundtrip test: create grouped shapes → save → open → verify structure
- PDF rendering test: grouped shapes with rotation → render PDF → verify positions
- Nested groups test: 4 levels deep → operations work correctly
- Ungroup test: group → ungroup → shapes have correct absolute positions

**Test Documents** (in `testdata/`):
- `simple-group.pptx`: 3 rectangles in one group
- `nested-groups.pptx`: Groups within groups (3 levels)
- `rotated-group.pptx`: Group rotated 45°, child shapes rotated 30°
- `flipped-group.pptx`: Group with horizontal flip
- `complex-diagram.pptx`: Mixed shapes, groups, and connectors

**Documentation**:
- Godoc for all new methods
- `examples/group-shapes/` directory with:
  - `create-group.go`: Demonstrates grouping shapes
  - `nested-groups.go`: Shows recursive groups
  - `transform-group.go`: Applies transformations
  - `pdf-group-rendering.go`: Generates PDF with groups

## Impact

**Affected specs**:
- `presentation-elements` (MODIFIED - enhanced GroupShape API)
- `drawingml-core` (ADDED - transformation utilities)
- `pdf-drawing` (ADDED - group rendering)
- `presentation-document` (ADDED - high-level grouping operations)

**New capabilities**:
- Programmatically create and manipulate grouped shapes
- Iterate over shapes within groups
- Apply transformations to entire groups
- Render grouped shapes correctly in PDF output
- Support recursive groups (groups within groups)
- Convert between local and absolute coordinates
- Group/ungroup operations with correct transform preservation

**Affected code**:
- `presentation/elements/shape_tree.go` - MODIFIED: Add methods to GroupShape
- `drawingml/transform/` - NEW PACKAGE: Transformation utilities
- `pdf/drawing/group_renderer.go` - NEW: Group rendering
- `pdf/presentation/renderer.go` - MODIFIED: Call group renderer
- `presentation/slide.go` - MODIFIED: Add CreateGroup/UngroupShape methods

**Breaking changes**: None. All changes are additive.

## Key Design Decisions

### 1. Recursive Rendering with Graphics State Stack

**Decision**: Use PDF graphics state stack (`q`/`Q`) operators for group transforms.

**Rationale**:
- PDF natively supports graphics state stacking
- Clean separation: enter group → push state → transform → render children → pop state
- Automatic handling of nested transforms (stack composition)
- No need to manually track/undo transforms

**Alternative considered**: Manual transform tracking (accumulate transforms, apply inverse when exiting group)
- **Rejected**: Error-prone, complex for deep nesting, PDF state stack is purpose-built for this

### 2. Transform Composition in Separate Package

**Decision**: Create `drawingml/transform/` package for transformation math.

**Rationale**:
- Reusable across presentation, word, spreadsheet
- Testable in isolation
- Keeps GroupShape API focused on element manipulation
- Aligns with single-responsibility principle

**Alternative considered**: Embed transform methods in GroupShape
- **Rejected**: Couples geometry math to element structure, harder to test

### 3. High-Level Grouping Operations on Slide

**Decision**: `slide.CreateGroup()` and `slide.UngroupShape()` as convenience methods.

**Rationale**:
- Common operations should be simple (80/20 rule)
- Matches user mental model ("group these shapes")
- Handles complex coordinate adjustments automatically
- Consistent with existing `slide.AddShape()` pattern

**Alternative considered**: Only low-level `group.AddShape()` API
- **Rejected**: Tedious for users (manual coordinate adjustment, removal from slide, adding group to slide)

### 4. Coordinate Space Handling

**Decision**: Child shapes use coordinates relative to group origin.

**Rationale**:
- Matches Office Open XML spec exactly
- Allows moving/rotating groups without recalculating all children
- Simplifies user code (group transform automatically affects children)

**Edge case handling**:
- Ungrouping: Convert child local coordinates to absolute before removing from group
- PDF rendering: Compose all parent transforms before rendering child

### 5. Maximum Nesting Depth

**Decision**: No artificial limit, but recommend ≤10 levels in documentation.

**Rationale**:
- Office applications support arbitrary nesting
- Artificial limit would break compatibility
- Deep nesting is rare in practice (performance concern only for extreme cases)

**Mitigation**: Document performance implications, add depth field to test suite

## Dependencies

**Existing**:
- `GroupShape` element (presentation/elements/shape_tree.go) - exists but needs methods
- `Transform2D` element (drawingml/) - exists
- PDF Page API (`pdf/core/`) - exists (provides SaveGraphicsState/RestoreGraphicsState/Transform)
- Shape renderers (`pdf/drawing/shape_renderer.go`, etc.) - exist

**New**:
- `drawingml/transform/` package - coordinate transformation utilities
- `pdf/drawing/group_renderer.go` - group rendering

**Schema compatibility**:
- No schema changes required (grpSp is standard ECMA-376 element)
- Existing XML parsing/serialization works as-is

## Success Criteria

- [ ] `GroupShape` has methods to add/remove/iterate child shapes
- [ ] Can create nested groups (3+ levels deep)
- [ ] `CreateGroup()` correctly moves shapes from slide to group
- [ ] `UngroupShape()` preserves absolute positions after flattening
- [ ] PDF rendering handles grouped shapes (no position errors)
- [ ] PDF rendering handles nested groups correctly
- [ ] Transform composition math is correct (verified with test matrices)
- [ ] Roundtrip test: group → save → open → group structure preserved
- [ ] LibreOffice can edit documents with groups created by goffice
- [ ] Performance: 1000-shape document with 10 groups renders PDF in <2s

## Out of Scope

- ContentPart support (Office 2010+ feature, defer to Phase 2 - requires investigation of Office 2010 compatibility layer)
- SmartArt diagrams (different element type: `dgm:grpSp`, not `p:grpSp`)
- 3D transformations (defer to 3D shapes proposal)
- Group-level visual effects (shadows/glows applied to entire group - future enhancement)
- Group locking/protection (security feature, separate proposal)
- Animation of groups (animation infrastructure not yet implemented)
- Word/Excel group shape enhancements (focus on PowerPoint first; Word/Excel can follow same pattern in future)
- Connector routing around groups (connector proposal scope)
- Extension lists (p:extLst) - preserved during round-trip but not actively manipulated in Phase 1

## References

- ECMA-376 Part 1, Section 21.3.2.14 (CT_GroupShape)
- ECMA-376 Part 1, Section 20.1.7.5 (CT_GroupShapeProperties)
- Open-XML-SDK: `DocumentFormat.OpenXml.Drawing.GroupShape`
- Open-XML-SDK: `DocumentFormat.OpenXml.Presentation.GroupShape`
