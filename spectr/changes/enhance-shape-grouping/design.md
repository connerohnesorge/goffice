# Design: Enhanced Shape Grouping Architecture

## Problem Statement

GroupShape elements exist in goffice but are essentially non-functional shells:
- No methods to access or manipulate child shapes
- No coordinate transformation utilities for group hierarchies
- PDF rendering completely ignores groups (treats them as invisible)
- No high-level API for common grouping operations

This blocks users from working with real-world presentations that heavily rely on grouped shapes for complex diagrams, org charts, and infographics.

## Current Architecture

```
Presentation Slide
    ↓
ShapeTree (p:spTree)
    ↓
├─ Shape (p:sp) ────────── PDF renders correctly
├─ Picture (p:pic) ────────── PDF renders correctly
├─ ConnectionShape (p:cxnSp) ── PDF renders correctly
└─ GroupShape (p:grpSp) ────── PDF IGNORES (no renderer)
        ↓
    ??? No API to access children ???
    ??? Group transform ignored ???
```

**Current GroupShape Implementation** (`presentation/elements/shape_tree.go:321-351`):
```go
type GroupShape struct {
    *openxml.CompositeElementBase
}

func NewGroupShape() *GroupShape {
    elem := openxml.NewCompositeElement(NamespacePresentationML, "grpSp", PrefixP)
    gs := &GroupShape{CompositeElementBase: elem}
    gs.AppendChild(NewNonVisualGroupShapeProperties())
    gs.AppendChild(NewGroupShapeProperties())
    return gs
}

func (gs *GroupShape) Clone() openxml.Element {
    cloned := gs.CompositeElementBase.Clone()
    return &GroupShape{CompositeElementBase: cloned.(*openxml.CompositeElementBase)}
}

// END OF IMPLEMENTATION - No other methods!
```

**Gap**: GroupShape can be created and cloned, but that's it. No way to add shapes, iterate children, or access properties.

## Proposed Architecture

```
Presentation Slide
    ↓
ShapeTree (p:spTree)
    ↓
├─ Shape (p:sp)
├─ Picture (p:pic)
└─ GroupShape (p:grpSp) ── NEW: GroupRenderer
        ↓                      ├─ Push PDF graphics state (q)
        │                      ├─ Apply group transform to CTM
    NEW GroupShape API         ├─ Recursively render children
        ↓                      └─ Pop PDF graphics state (Q)
    ├─ Shapes() []*Shape
    ├─ Pictures() []*Picture
    ├─ GroupShapes() []*GroupShape  ← Recursive!
    ├─ AddShape() *Shape
    ├─ AddGroupShape() *GroupShape
    └─ GroupShapeProperties()
            ↓
        Transform2D (a:xfrm)
            ↓
        NEW drawingml/transform/ ── Transform composition utilities
            ├─ ComputeAbsoluteTransform()
            ├─ LocalToAbsolute()
            ├─ MultiplyMatrices()
            └─ DecomposeTransform()
```

## Design Alternatives Considered

### Alternative 1: Flatten Groups Before PDF Rendering

**Approach**: Convert all groups to individual shapes with absolute coordinates before PDF rendering.

**Pros**:
- Simpler PDF renderer (no group handling needed)
- No graphics state stack management

**Cons**:
- Loss of group structure in intermediate representation
- Cannot preserve groups for other outputs (SVG, future formats)
- Extra computation (transform every shape)
- Breaks abstraction (PDF renderer should match document structure)

**Decision**: REJECTED - Preserve group hierarchy throughout rendering pipeline

### Alternative 2: Manual Transform Tracking

**Approach**: Track cumulative transform matrix, apply to each shape, manually reverse when exiting group.

**Pros**:
- More explicit control
- Could optimize by caching transforms

**Cons**:
- Error-prone (easy to forget to reverse transform)
- Complex for nested groups (must track stack manually)
- Reinvents PDF's graphics state stack (which is purpose-built for this)

**Decision**: REJECTED - Use PDF's native graphics state stack

### Alternative 3: Single-Level Groups Only

**Approach**: Disallow groups within groups.

**Pros**:
- Simpler implementation
- Easier to reason about transforms

**Cons**:
- Breaks compatibility with Office (allows arbitrary nesting)
- Real documents use nested groups heavily (org charts, complex diagrams)
- Would fail to round-trip many real documents

**Decision**: REJECTED - Must support recursive groups for compatibility

### Alternative 4: Copy ShapeTree API for GroupShape (CHOSEN)

**Approach**: GroupShape has same child manipulation methods as ShapeTree.

**Pros**:
- Consistent API pattern (users already know ShapeTree)
- Reuses existing knowledge
- Both are containers of shapes → same interface makes sense

**Cons**:
- Some code duplication (mitigated by shared helper functions)

**Decision**: ACCEPTED - Best balance of consistency and usability

## GroupShape API Design

### Method Naming Rationale

Follow ShapeTree's existing pattern exactly:

| ShapeTree Method | GroupShape Method | Rationale |
|------------------|-------------------|-----------|
| `Shapes() []*Shape` | `Shapes() []*Shape` | Same semantics: return child shapes |
| `Pictures() []*Picture` | `Pictures() []*Picture` | Same semantics: return child pictures |
| `GroupShapes() []*GroupShape` | `GroupShapes() []*GroupShape` | Return nested groups (recursive!) |
| `AddShape() *Shape` | `AddShape() *Shape` | Create and add shape to group |
| `AddPicture(relId) *Picture` | `AddPicture(relId) *Picture` | Create and add picture |
| `AddGroupShape() *GroupShape` | `AddGroupShape() *GroupShape` | Create and add nested group |

**Why not `Children()` or `GetShapes()`?**
- `Children()` is too generic (returns all element types, not just shapes)
- `GetShapes()` is redundant (`Get` prefix not used elsewhere in codebase)
- `Shapes()` matches existing API style

### Property Access

```go
// Access group-level transform
props := group.GroupShapeProperties()
xfrm := props.Transform()  // Returns Transform2D (groups have child coordinate space elements)

// Access child coordinate space (specific to groups)
if chOffX, chOffY, hasChOff := xfrm.ChildOffset(); hasChOff {
    // Child space origin is offset from (0,0)
    _ = chOffX
    _ = chOffY
}
if chExtCx, chExtCy, hasChExt := xfrm.ChildExtent(); hasChExt {
    // Child space has different size than group (viewport scaling)
    _ = chExtCx
    _ = chExtCy
}

// vs alternative: Direct accessor
xfrm := group.Transform() // REJECTED - breaks pattern

// Rationale: Consistent with Shape.ShapeProperties(), maintains separation
// between element structure (GroupShape) and visual properties (GroupShapeProperties)
```

**Critical Note**: Groups use the same `Transform2D` element type as regular shapes, but the XML for groups (CT_GroupTransform2D in ECMA-376) includes additional child elements:
- `ChildOffset` (a:chOff) - defines the origin of the child coordinate space (optional child element)
- `ChildExtent` (a:chExt) - defines the size of the child coordinate space (optional child element)

Transform2D provides accessors for these optional child elements: `ChildOffset()` and `ChildExtent()` return nil if not present.

These properties create a **viewport transformation** that maps child shapes' coordinates to the group's coordinate space. This is critical for correct rendering!

### Implementation Strategy

**Option A**: Copy-paste ShapeTree methods, replace namespace constant
```go
func (gs *GroupShape) Shapes() []*Shape {
    var shapes []*Shape
    for child := range gs.Children() {
        if child.LocalName() == "sp" &&
            child.NamespaceURI() == NamespacePresentationML {
            // ... same logic as ShapeTree.Shapes()
        }
    }
    return shapes
}
```

**Option B**: Extract shared helper function
```go
func getShapesFromContainer(container *openxml.CompositeElementBase) []*Shape {
    // shared logic
}

func (gs *GroupShape) Shapes() []*Shape {
    return getShapesFromContainer(gs.CompositeElementBase)
}

func (st *ShapeTree) Shapes() []*Shape {
    return getShapesFromContainer(st.CompositeElementBase)
}
```

**Decision**: Option A for now (simpler, less indirection). If duplication becomes problematic (3+ container types), refactor to Option B in separate change.

## Coordinate Transformation Design

### Transform Hierarchy

```
Slide (absolute coordinates)
    ↓
Group A (p:grpSp)
    Transform2D (a:xfrm):
        - Offset (a:off): (100, 200)  [group position on slide]
        - Extent (a:ext): (1000, 800)  [group bounding box size]
        - ChildOffset (a:chOff): (0, 0)  [origin of child coordinate space]
        - ChildExtent (a:chExt): (2000, 1600)  [size of child coordinate space]
        - Rotation (rot attr): 45°
    ↓
    Group B (p:grpSp) - nested in A
        Transform2D (a:xfrm):
            - Offset (a:off): (500, 400)  [position in Group A's child space]
            - Extent (a:ext): (600, 500)
            - ChildOffset (a:chOff): (0, 0)
            - ChildExtent (a:chExt): (600, 500)
            - Rotation (rot attr): 30°
        ↓
        Shape C (p:sp) - in B
            Transform2D (a:xfrm):
                - Offset (a:off): (10, 10)  [position in Group B's child space]
                - Extent (a:ext): (50, 50)
            ↓
        Absolute position calculation:
            1. Shape C's position in Group B's child space: (10, 10)
            2. Apply Group B's viewport transform:
               scale_x = B.Extents.cx / B.ChildExtents.cx = 600/600 = 1.0
               scale_y = B.Extents.cy / B.ChildExtents.cy = 500/500 = 1.0
               translate = B.ChildOffset = (0, 0)
               → (10, 10) in Group B's local space
            3. Apply Group B's transform (rotate 30°, position at B.Offset)
               → position relative to Group A's child space
            4. Apply Group A's viewport transform:
               scale_x = A.Extents.cx / A.ChildExtents.cx = 1000/2000 = 0.5
               scale_y = A.Extents.cy / A.ChildExtents.cy = 800/1600 = 0.5
               → scales coordinates by 0.5x
            5. Apply Group A's transform (rotate 45°, position at A.Offset)
               → final absolute position on slide
```

**Critical Insight**: The ChildExtents/Extents ratio defines a **scaling factor** for child coordinates. If ChildExtents > Extents, children are scaled down. This allows groups to have a larger coordinate space for precision while rendering at a smaller size.

### Matrix Representation

**DrawingML Transform2D** (a:xfrm):
```xml
<a:xfrm>
    <a:off x="914400" y="1828800"/>  <!-- EMU units -->
    <a:ext cx="1828800" cy="914400"/>
    <a:rot val="5400000"/>           <!-- 1/60000 degree units -->
    <a:flipH val="0"/>
    <a:flipV val="0"/>
</a:xfrm>
```

**Conversion to Affine Transform Matrix**:
```
[a  b  0]   where:
[c  d  0]   a = scaleX * cos(rotation)
[e  f  1]   b = scaleX * sin(rotation)
            c = scaleY * -sin(rotation)
            d = scaleY * cos(rotation)
            e = translateX
            f = translateY

flipH: scaleX *= -1
flipV: scaleY *= -1
```

### Implementation: Transform Package

```go
package transform

import "math"

// Matrix represents a 2D affine transformation matrix.
type Matrix struct {
    A, B, C, D, E, F float64
}

// Identity returns the identity matrix (no transformation).
func Identity() Matrix {
    return Matrix{A: 1, D: 1}
}

// Translate creates a translation matrix.
func Translate(dx, dy float64) Matrix {
    return Matrix{A: 1, D: 1, E: dx, F: dy}
}

// Rotate creates a rotation matrix (angle in radians).
func Rotate(angle float64) Matrix {
    cos := math.Cos(angle)
    sin := math.Sin(angle)
    return Matrix{A: cos, B: sin, C: -sin, D: cos}
}

// Scale creates a scaling matrix.
func Scale(sx, sy float64) Matrix {
    return Matrix{A: sx, D: sy}
}

// Multiply returns m1 * m2 (note: order matters!).
func Multiply(m1, m2 Matrix) Matrix {
    return Matrix{
        A: m1.A*m2.A + m1.B*m2.C,
        B: m1.A*m2.B + m1.B*m2.D,
        C: m1.C*m2.A + m1.D*m2.C,
        D: m1.C*m2.B + m1.D*m2.D,
        E: m1.E*m2.A + m1.F*m2.C + m2.E,
        F: m1.E*m2.B + m1.F*m2.D + m2.F,
    }
}

// Invert returns the inverse matrix (m * m.Invert() = Identity).
func (m Matrix) Invert() Matrix {
    det := m.A*m.D - m.B*m.C
    if det == 0 {
        return Identity() // Singular matrix - cannot invert
    }
    return Matrix{
        A: m.D / det,
        B: -m.B / det,
        C: -m.C / det,
        D: m.A / det,
        E: (m.C*m.F - m.D*m.E) / det,
        F: (m.B*m.E - m.A*m.F) / det,
    }
}

// TransformPoint applies the matrix to a point.
func (m Matrix) TransformPoint(x, y float64) (float64, float64) {
    return m.A*x + m.C*y + m.E, m.B*x + m.D*y + m.F
}

// FromTransform2D converts DrawingML Transform2D to Matrix.
// For groups (when ChildOffset/ChildExtent are present), this includes viewport transformation.
func FromTransform2D(xfrm *drawingml.Transform2D) Matrix {
    if xfrm == nil {
        return Identity()
    }

    // Get offset (translation)
    off := xfrm.Offset()
    tx := 0.0
    ty := 0.0
    if off.X != 0 || off.Y != 0 {
        tx = float64(off.X)
        ty = float64(off.Y)
    }

    // Get rotation (convert from 1/60000 degrees to radians)
    // Transform2D stores rotation as an attribute
    rot := xfrm.GetAttribute("rot", "")
    angle := 0.0
    if rot != nil {
        rotVal, _ := strconv.ParseInt(rot.Value(), 10, 64)
        angle = float64(rotVal) / 60000.0 * math.Pi / 180.0
    }

    // Get flip
    flipH := false
    flipV := false
    if flipHAttr := xfrm.GetAttribute("flipH", ""); flipHAttr != nil {
        flipH = flipHAttr.Value() == "1" || flipHAttr.Value() == "true"
    }
    if flipVAttr := xfrm.GetAttribute("flipV", ""); flipVAttr != nil {
        flipV = flipVAttr.Value() == "1" || flipVAttr.Value() == "true"
    }

    flipScaleX := 1.0
    flipScaleY := 1.0
    if flipH {
        flipScaleX = -1.0
    }
    if flipV {
        flipScaleY = -1.0
    }

    // Get viewport transformation (child coordinate space - for groups only)
    // ChildOffset and ChildExtent are optional child elements
    viewportMatrix := Identity()

    ext := xfrm.Extent()
    chExtCx, chExtCy, hasChExt := xfrm.ChildExtent()
    chOffX, chOffY, hasChOff := xfrm.ChildOffset()

    if hasChExt && ext.Cx != 0 && ext.Cy != 0 && chExtCx != 0 && chExtCy != 0 {
        // This is a group transform with viewport
        // Scale factor: map ChildExtent to Extent
        scaleX := float64(ext.Cx) / float64(chExtCx)
        scaleY := float64(ext.Cy) / float64(chExtCy)

        // Translation: ChildOffset defines child space origin
        transX := 0.0
        transY := 0.0
        if hasChOff {
            transX = float64(chOffX)
            transY = float64(chOffY)
        }

        // Viewport transform: translate by negative child offset, then scale
        // This maps child coordinates to group's local space
        // P_group = (P_child - chOff) * scale
        // Matrix multiplication order: rightmost matrix is applied FIRST to points
        // So to do "translate then scale", we need Multiply(Scale, Translate)
        viewportMatrix = Multiply(Scale(scaleX, scaleY), Translate(-transX, -transY))
    }

    // Build composite transform: position * rotation * flip * viewport
    // Viewport is applied FIRST (innermost) because it operates on child coordinates
    // Matrix multiplication order: rightmost matrix is applied FIRST to points
    // Composition: viewport → flip → rotate → translate
    trans := Translate(tx, ty)
    rotate := Rotate(angle)
    flip := Scale(flipScaleX, flipScaleY)

    return Multiply(trans, Multiply(rotate, Multiply(flip, viewportMatrix)))
}

// ComputeAbsoluteTransform traverses parent chain and composes transforms.
// This handles both the element's own transform and all parent group transforms.
func ComputeAbsoluteTransform(elem openxml.Element, root openxml.Element) Matrix {
    matrices := []Matrix{}

    // First, get the element's own transform (if it's a Shape or GroupShape)
    if shape, ok := elem.(*Shape); ok {
        if props := shape.ShapeProperties(); props != nil {
            if xfrm := props.Transform(); xfrm != nil {
                matrices = append(matrices, FromTransform2D(xfrm))
            }
        }
    } else if group, ok := elem.(*GroupShape); ok {
        if props := group.GroupShapeProperties(); props != nil {
            if xfrm := props.Transform(); xfrm != nil {
                matrices = append(matrices, FromTransform2D(xfrm))
            }
        }
    }

    // Walk up parent chain collecting GROUP transforms only
    current := elem.Parent()
    for current != nil && current != root {
        if group, ok := current.(*GroupShape); ok {
            if props := group.GroupShapeProperties(); props != nil {
                if xfrm := props.Transform(); xfrm != nil {
                    // Prepend parent transform (parents come before children in multiplication)
                    matrices = append([]Matrix{FromTransform2D(xfrm)}, matrices...)
                }
            }
        }
        current = current.Parent()
    }

    // Multiply all transforms (parent → child order)
    // Result: T_parent * T_child means parent transform applies first
    result := Identity()
    for _, m := range matrices {
        result = Multiply(result, m)
    }

    return result
}
```

### Edge Case: Singular Matrices

**Problem**: If a group has zero scale (sx=0 or sy=0), the transform matrix becomes singular (determinant = 0) and cannot be inverted.

**Occurrence**: Rare but possible if user explicitly sets scale to zero.

**Handling**:
```go
func (m Matrix) Invert() Matrix {
    det := m.A*m.D - m.B*m.C
    if math.Abs(det) < 1e-10 { // Near-zero determinant
        // Cannot invert - return identity as fallback
        // Log warning in production
        return Identity()
    }
    // ... normal inversion
}
```

**Testing**: Include test case with zero-scale group, verify graceful handling.

## PDF Rendering Design

### Graphics State Stack Usage

PDF maintains a graphics state stack via `q` (push) and `Q` (pop) operators:

```pdf
q                     % Save state
1 0 0 1 100 200 cm    % Apply transform (translate 100, 200)
% ... render child shapes using transformed coordinates ...
Q                     % Restore state (transform is gone)
```

**Mapping to Groups**:
```go
func (r *GroupRenderer) RenderGroup(group *GroupShape) error {
    // 1. Push state
    r.ctx.Page.SaveGraphicsState()
    defer r.ctx.Page.RestoreGraphicsState()

    // 2. Apply group transform
    if props := group.GroupShapeProperties(); props != nil {
        if xfrm := props.Transform(); xfrm != nil {
            matrix := transform.FromTransform2D(xfrm)
            pdfMatrix := core.Matrix{
                A: matrix.A, B: matrix.B,
                C: matrix.C, D: matrix.D,
                E: matrix.E, F: matrix.F,
            }
            r.ctx.Page.Transform(pdfMatrix)
        }
    }

    // 3. Render children
    for child := range group.Children() {
        switch c := child.(type) {
        case *Shape:
            r.shapeRenderer.Render(c)
        case *Picture:
            r.pictureRenderer.Render(c)
        case *GroupShape:
            r.RenderGroup(c) // Recursive!
        case *ConnectionShape:
            r.connectorRenderer.Render(c)
        case *GraphicFrame:
            r.graphicFrameRenderer.Render(c)
        }
    }

    // 4. Pop state (defer handles this)
    return nil
}
```

### Coordinate Conversion

**DrawingML → PDF**:
- DrawingML: EMU units (914,400 EMU = 1 inch), top-left origin
- PDF: Points (72 points = 1 inch), bottom-left origin

**Conversion already exists** in `pdf/core/coord_transform.go` (verified via grep):
```go
func EMUToPoints(emu int64) float64 {
    return float64(emu) / 12700.0
}

func ConvertY(yEMU int64, pageHeight float64) float64 {
    yPoints := EMUToPoints(yEMU)
    return pageHeight - yPoints // Flip Y-axis
}
```

**Group rendering uses relative coordinates**:
- Child shapes have coordinates relative to group origin
- PDF CTM handles conversion automatically (group transform moves origin)
- No special Y-flip needed inside group (group's transform handles it)

### Renderer Integration

**Modify `pdf/presentation/renderer.go`**:
```go
func (r *PresentationRenderer) renderSlide(slide *Slide) error {
    shapeTree := slide.ShapeTree()

    // Create group renderer
    groupRenderer := drawing.NewGroupRenderer(r.ctx)

    for child := range shapeTree.Children() {
        switch c := child.(type) {
        case *Shape:
            r.shapeRenderer.Render(c)
        case *Picture:
            r.pictureRenderer.Render(c)
        case *GroupShape:
            groupRenderer.RenderGroup(c) // NEW!
        case *ConnectionShape:
            r.connectorRenderer.Render(c)
        case *GraphicFrame:
            r.graphicFrameRenderer.Render(c)
        }
    }

    return nil
}
```

### Performance Considerations

**Graphics State Stack Depth**:
- PDF specification recommends limiting stack depth to 28 (Adobe Acrobat limit)
- Each group adds 1 level to stack
- Maximum nesting depth of 10 groups → stack depth ≤ 10 (well within limit)

**Transform Multiplication Cost**:
- Each group adds 1 matrix multiplication (6 float multiplications + 6 additions = ~20 ops)
- For 10-level nesting: 200 operations total (negligible)
- Modern CPUs: ~1 nanosecond per operation = 200 ns total

**Rendering Order**:
- Depth-first traversal (render children before siblings)
- Matches Office rendering order
- Critical for overlapping shapes (z-order)

## High-Level Grouping Operations Design

### CreateGroup API

**User Story**: User has 3 shapes on a slide and wants to group them.

**Before**:
```go
slide := presentation.Slides()[0]
shapes := slide.ShapeTree().Shapes()
// ... now what? How to group them?
```

**After**:
```go
slide := presentation.Slides()[0]
shapes := slide.ShapeTree().Shapes()
rect1 := shapes[0]
rect2 := shapes[1]
circle := shapes[2]

// Simple grouping
group := slide.CreateGroup(rect1, rect2, circle)

// Result:
// - rect1, rect2, circle removed from slide
// - New GroupShape created
// - Shapes added to group (with adjusted coordinates)
// - Group added to slide
```

### Implementation Details

```go
func (s *Slide) CreateGroup(shapes ...interface{}) *GroupShape {
    // 1. Create new group
    group := elements.NewGroupShape()

    // 2. Calculate bounding box of all input shapes
    minX, minY, maxX, maxY := computeBoundingBox(shapes)

    // 3. Set group position to bounding box top-left
    groupProps := group.GroupShapeProperties()
    xfrm := drawingml.NewTransform2D()
    off := drawingml.NewOffset()
    off.SetXAttr(minX)
    off.SetYAttr(minY)
    xfrm.SetOffset(off)

    ext := drawingml.NewExtents()
    ext.SetCxAttr(maxX - minX)
    ext.SetCyAttr(maxY - minY)
    xfrm.SetExtents(ext)

    groupProps.SetTransform(xfrm)

    // 4. Move each shape into group, adjust coordinates
    for _, shape := range shapes {
        // Remove from current parent
        shape.Parent().RemoveChild(shape)

        // Adjust coordinates to be relative to group origin
        adjustCoordinatesRelativeToGroup(shape, minX, minY)

        // Add to group
        group.AppendChild(shape)
    }

    // 5. Add group to slide
    s.ShapeTree().AppendChild(group)

    return group
}
```

### UngroupShape API

**User Story**: User wants to flatten a group back to individual shapes.

```go
group := slide.ShapeTree().GroupShapes()[0]

// Ungroup
shapes := slide.UngroupShape(group)

// Result:
// - Group removed from slide
// - All child shapes added back to slide
// - Shapes have absolute coordinates (group transform baked in)
```

**Implementation**:
```go
func (s *Slide) UngroupShape(group *GroupShape) []openxml.Element {
    var ungrouped []openxml.Element

    // 1. Get group's absolute transform
    groupMatrix := transform.ComputeAbsoluteTransform(group, s.ShapeTree())

    // 2. For each child, compute absolute position
    for child := range group.Children() {
        if isShapeElement(child) {
            // Apply group transform to child's local position
            applyAbsoluteTransform(child, groupMatrix)

            // Remove from group
            group.RemoveChild(child)

            // Add to slide
            s.ShapeTree().AppendChild(child)

            ungrouped = append(ungrouped, child)
        }
    }

    // 3. Remove group from slide
    s.ShapeTree().RemoveChild(group)

    return ungrouped
}
```

### Edge Case: Ungrouping Nested Groups

**Scenario**: Group A contains Group B which contains shapes.

**Option 1**: Flatten recursively (ungroup B, then ungroup A)
```go
UngroupShape(groupA) // Also ungroups groupB automatically
```

**Option 2**: Only ungroup one level (groupB stays grouped)
```go
UngroupShape(groupA) // groupB becomes a top-level group
```

**Decision**: Option 2 (one level only)
- **Rationale**: Matches PowerPoint behavior, gives user control, simpler implementation
- **User can recursively ungroup manually if needed**:
```go
func ungroupRecursively(slide *Slide, group *GroupShape) {
    elements := slide.UngroupShape(group)
    for _, elem := range elements {
        if subgroup, ok := elem.(*GroupShape); ok {
            ungroupRecursively(slide, subgroup)
        }
    }
}
```

## Testing Strategy

### Unit Tests

**GroupShape API Tests** (`presentation/elements/shape_tree_test.go`):
```go
func TestGroupShape_AddShape(t *testing.T) {
    group := NewGroupShape()

    shape1 := NewShape()
    group.AppendChild(shape1)

    shapes := group.Shapes()
    assert.Len(t, shapes, 1)
    assert.Equal(t, shape1, shapes[0])
}

func TestGroupShape_NestedGroups(t *testing.T) {
    outerGroup := NewGroupShape()
    innerGroup := NewGroupShape()
    outerGroup.AppendChild(innerGroup)

    shape := NewShape()
    innerGroup.AppendChild(shape)

    // Verify hierarchy
    assert.Len(t, outerGroup.GroupShapes(), 1)
    assert.Len(t, innerGroup.Shapes(), 1)
}
```

**Transform Composition Tests** (`drawingml/transform/transform_test.go`):
```go
func TestMultiplyMatrices(t *testing.T) {
    // Translate then rotate
    translate := Translate(100, 200)
    rotate := Rotate(math.Pi / 4) // 45 degrees

    result := Multiply(translate, rotate)

    // Apply to point (0, 0)
    x, y := result.TransformPoint(0, 0)

    // Translated (100, 200), then rotated 45°
    // Expected: (~70.7, ~282.8) - calculated externally
    assert.InDelta(t, 70.7, x, 0.1)
    assert.InDelta(t, 282.8, y, 0.1)
}

func TestMatrixInversion(t *testing.T) {
    m := Matrix{A: 2, B: 0, C: 0, D: 3, E: 10, F: 20}
    inv := m.Invert()

    // m * inv should equal identity
    result := Multiply(m, inv)

    assert.InDelta(t, 1.0, result.A, 1e-10)
    assert.InDelta(t, 0.0, result.B, 1e-10)
    assert.InDelta(t, 0.0, result.C, 1e-10)
    assert.InDelta(t, 1.0, result.D, 1e-10)
    assert.InDelta(t, 0.0, result.E, 1e-10)
    assert.InDelta(t, 0.0, result.F, 1e-10)
}

func TestSingularMatrixInversion(t *testing.T) {
    // Zero determinant (scale 0)
    m := Matrix{A: 0, B: 0, C: 0, D: 1, E: 0, F: 0}
    inv := m.Invert()

    // Should return identity as fallback
    assert.Equal(t, Identity(), inv)
}
```

### Integration Tests

**Roundtrip Tests** (`presentation/integration_test.go`):
```go
func TestGroupShapeRoundtrip(t *testing.T) {
    // Create presentation with grouped shapes
    pres := presentation.New()
    slide := pres.AddSlide()

    shape1 := slide.ShapeTree().AddShape()
    shape2 := slide.ShapeTree().AddShape()
    group := slide.CreateGroup(shape1, shape2)

    // Save
    buf := new(bytes.Buffer)
    err := pres.Save(buf)
    require.NoError(t, err)

    // Reopen
    pres2, err := presentation.Open(bytes.NewReader(buf.Bytes()))
    require.NoError(t, err)

    // Verify group exists
    slide2 := pres2.Slides()[0]
    groups := slide2.ShapeTree().GroupShapes()
    assert.Len(t, groups, 1)

    // Verify children
    assert.Len(t, groups[0].Shapes(), 2)
}
```

**PDF Rendering Tests** (`pdf/presentation/group_rendering_test.go`):
```go
func TestPDFGroupRendering(t *testing.T) {
    // Create presentation with rotated group
    pres := presentation.New()
    slide := pres.AddSlide()

    rect1 := slide.ShapeTree().AddShape()
    setPosition(rect1, 100, 100, 200, 100)

    rect2 := slide.ShapeTree().AddShape()
    setPosition(rect2, 100, 200, 200, 100)

    group := slide.CreateGroup(rect1, rect2)

    // Rotate group 45 degrees
    groupProps := group.GroupShapeProperties()
    xfrm := groupProps.Transform()
    xfrm.SetRotation(45 * 60000)

    // Render to PDF
    pdfBytes, err := pdf.RenderPresentation(pres)
    require.NoError(t, err)

    // Validate PDF structure
    assert.Contains(t, string(pdfBytes), "q\n") // Graphics state save
    assert.Contains(t, string(pdfBytes), "Q\n") // Graphics state restore
    // TODO: Add visual validation (compare with reference PDF)
}
```

### Visual Validation Tests

**Automated Visual Comparison**:
1. Render test document to PDF using goffice
2. Render same document using Microsoft PowerPoint → PDF
3. Convert both PDFs to images (using pdf2image or similar)
4. Compute pixel difference percentage
5. Assert difference < 2% (allows for minor rendering variations)

**Test Documents** (in `testdata/grouping/`):
- `simple-group.pptx`: 3 rectangles in one group
- `nested-3-levels.pptx`: Group → Group → Group → Shapes (3 levels deep)
- `rotated-group.pptx`: Group rotated 45°, children also rotated
- `flipped-group.pptx`: Group with flipH=true
- `mixed-shapes-group.pptx`: Group containing shapes, pictures, connectors
- `complex-org-chart.pptx`: Real-world org chart (10+ groups, 50+ shapes)

## Migration Path

### Phase 1: GroupShape API (Week 1)
- Add Shapes(), Pictures(), GroupShapes(), AddShape(), etc. methods to GroupShape
- Add Transform() accessor to GroupShapeProperties
- Unit tests for all new methods

### Phase 2: Transform Utilities (Week 1)
- Create `drawingml/transform/` package
- Implement Matrix, Multiply, Invert, TransformPoint, etc.
- Unit tests for transform math (including edge cases)

### Phase 3: PDF Rendering (Week 2)
- Implement GroupRenderer in `pdf/drawing/group_renderer.go`
- Integrate with `pdf/presentation/renderer.go`
- Unit tests with mock Page
- Integration tests with real PDF generation

### Phase 4: High-Level API (Week 1)
- Implement `slide.CreateGroup()` and `slide.UngroupShape()`
- Handle coordinate adjustments
- Unit and integration tests

### Phase 5: Testing & Validation (Week 1)
- Visual comparison tests
- LibreOffice/PowerPoint compatibility tests
- Performance benchmarks
- Documentation and examples

**Total: 6 weeks**

## Risks & Mitigation

### Risk 1: Transform Math Errors

**Impact**: High (shapes render in wrong positions)
**Probability**: Medium (matrix math is error-prone)

**Mitigation**:
- Extensive unit tests with known transform values
- Cross-validate with external matrix library (e.g., gonum/mat) during development
- Visual tests catch position errors

### Risk 2: Deep Nesting Performance

**Impact**: Low (only affects extreme cases)
**Probability**: Low (real documents rarely exceed 5 levels)

**Mitigation**:
- Document recommended depth limit (≤10)
- Add depth counter in tests (fail if >20 levels)
- Benchmark with pathological cases (100 levels)

### Risk 3: Office Compatibility Issues

**Impact**: Medium (documents don't round-trip correctly)
**Probability**: Low (using standard ECMA-376 elements)

**Mitigation**:
- Test with real Office-created documents
- Roundtrip tests with PowerPoint 2016, 2019, 365
- LibreOffice compatibility validation

### Risk 4: Singular Matrix Handling

**Impact**: Low (gracefully degrades)
**Probability**: Very low (requires intentional zero scale)

**Mitigation**:
- Explicit check in Invert() function
- Return identity matrix as fallback
- Log warning for debugging
- Test case verifies behavior

## Success Metrics

- [ ] GroupShape has 10+ new methods (Shapes, Pictures, AddShape, etc.)
- [ ] Can create nested groups (tested to 5 levels deep)
- [ ] Transform composition math passes 20+ unit test cases
- [ ] PDF rendering handles groups (integration test with real PDF)
- [ ] CreateGroup/UngroupShape work correctly (roundtrip test passes)
- [ ] Visual comparison shows <2% difference from PowerPoint output
- [ ] LibreOffice can open and edit documents with goffice-created groups
- [ ] Performance: 1000-shape document with 10 groups renders in <2s
- [ ] Zero regressions in existing presentation tests

## Future Enhancements (Not in Scope)

1. **Group-level effects**: Apply shadow/glow to entire group (not just individual shapes)
2. **3D group transforms**: Rotation in 3D space
3. **Group animation**: Animate entire group as single unit
4. **Connector routing around groups**: Connectors avoid group bounding boxes
5. **Word/Excel group enhancement**: Currently focused on PowerPoint; extend to other apps later
6. **Performance optimization**: Cache absolute transforms, lazy computation
