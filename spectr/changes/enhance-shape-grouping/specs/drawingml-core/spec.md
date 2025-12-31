# DrawingML Core Spec Delta

## ADDED Requirements

### Requirement: Affine Transform Matrix

The system SHALL provide an affine transformation matrix type for 2D geometric transforms.

#### Scenario: Identity matrix
- WHEN Identity() is called
- THEN a matrix with A=1, B=0, C=0, D=1, E=0, F=0 is returned
- AND applying the matrix to any point returns the same point

#### Scenario: Translation matrix
- WHEN Translate(dx=100, dy=200) is called
- THEN a matrix is returned
- AND applying it to point (0, 0) yields (100, 200)
- AND applying it to point (50, 75) yields (150, 275)

#### Scenario: Rotation matrix
- WHEN Rotate(angle=π/2) is called (90 degrees)
- THEN a matrix is returned
- AND applying it to point (1, 0) yields approximately (0, 1)
- AND applying it to point (0, 1) yields approximately (-1, 0)

#### Scenario: Scaling matrix
- WHEN Scale(sx=2.0, sy=3.0) is called
- THEN a matrix is returned
- AND applying it to point (10, 20) yields (20, 60)

#### Scenario: Flip horizontal
- WHEN Scale(sx=-1.0, sy=1.0) is called
- THEN a matrix is returned
- AND applying it to point (100, 50) yields (-100, 50)

### Requirement: Matrix Operations

The system SHALL provide operations for combining and manipulating transformation matrices.

#### Scenario: Matrix multiplication
- GIVEN translate matrix T(100, 200)
- AND rotate matrix R(45°)
- WHEN Multiply(T, R) is called
- THEN a composite matrix is returned
- AND the composite applies R first to points, then T (standard matrix algebra: rightmost-first)
- AND for point P: result = T * R * P (matrix notation: multiply right-to-left)

#### Scenario: Matrix multiplication order
- GIVEN matrices M1 and M2
- WHEN Multiply(M1, M2) is called
- THEN the result equals M1 * M2 in mathematical notation
- AND when applied to a point P: (M1 * M2) * P means M2 is applied to P first, then M1
- AND Multiply(M2, M1) may yield different result (order matters)
- AND this follows standard matrix algebra convention: rightmost matrix in product applies first to points

#### Scenario: Matrix inversion
- GIVEN a matrix M with non-zero determinant
- WHEN M.Invert() is called
- THEN the inverse matrix M⁻¹ is returned
- AND Multiply(M, M⁻¹) equals the identity matrix (within floating-point tolerance)

#### Scenario: Singular matrix inversion
- GIVEN a matrix with zero determinant (e.g., Scale(0, 1))
- WHEN Invert() is called
- THEN the identity matrix is returned as fallback
- AND no error or panic occurs

#### Scenario: Point transformation
- GIVEN a matrix M
- AND a point P(x, y)
- WHEN M.TransformPoint(x, y) is called
- THEN the transformed point (x', y') is returned
- WHERE x' = M.A*x + M.C*y + M.E
- AND y' = M.B*x + M.D*y + M.F

### Requirement: Transform2D to Matrix Conversion

The system SHALL provide conversion from DrawingML Transform2D elements to affine transformation matrices.

#### Scenario: Convert simple transform
- GIVEN a Transform2D with offset (100, 200) and no rotation/flip
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix equivalent to Translate(100, 200) is returned

#### Scenario: Convert transform with rotation
- GIVEN a Transform2D with rotation 5400000 (90 degrees in 1/60000 units)
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix with rotation component is returned
- AND the rotation angle is π/2 radians

#### Scenario: Convert transform with flip
- GIVEN a Transform2D with flipH=true
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix with scale -1.0 in X direction is returned

#### Scenario: Convert composite transform
- GIVEN a Transform2D with offset (100, 200), rotation 45°, and flipV=true
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix combining translation, rotation, and vertical flip is returned
- AND the operations are applied in correct order

#### Scenario: Handle nil transform
- WHEN FromTransform2D(nil) is called
- THEN the identity matrix is returned

### Requirement: Transform2D with Child Coordinate Space

The system SHALL provide accessor methods on Transform2D to access child coordinate space elements (for groups).

#### Scenario: Access child offset
- GIVEN a Transform2D element with a:chOff child element with x=100, y=200
- WHEN x, y, present := ChildOffset() is called
- THEN x == 100 and y == 200
- AND present == true

#### Scenario: Access child offset when not present
- GIVEN a Transform2D element without a:chOff child element (regular shape)
- WHEN x, y, present := ChildOffset() is called
- THEN x == 0 and y == 0
- AND present == false

#### Scenario: Access child extent
- GIVEN a Transform2D element with a:chExt child element with cx=1000, cy=800
- WHEN cx, cy, present := ChildExtent() is called
- THEN cx == 1000 and cy == 800
- AND present == true

#### Scenario: Access child extent when not present
- GIVEN a Transform2D element without a:chExt child element (regular shape)
- WHEN cx, cy, present := ChildExtent() is called
- THEN cx == 0 and cy == 0
- AND present == false

#### Scenario: Set child offset
- GIVEN a Transform2D element
- WHEN SetChildOffset(100, 200) is called
- THEN a:chOff child element is created or updated
- AND subsequent ChildOffset() returns (100, 200, true)

#### Scenario: Set child extent
- GIVEN a Transform2D element
- WHEN SetChildExtent(1000, 800) is called
- THEN a:chExt child element is created or updated
- AND subsequent ChildExtent() returns (1000, 800, true)

### Requirement: Transform2D to Matrix Conversion with Viewport

The system SHALL extend FromTransform2D to handle viewport transformation when ChildOffset/ChildExtent are present.

#### Scenario: Convert Transform2D with viewport
- GIVEN a Transform2D with:
  - Offset (100, 200) and Extent (1000, 800) [group position and size]
  - ChildOffset (0, 0) and ChildExtent (2000, 1600) [child coordinate space]
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix is returned
- AND the matrix includes translation to Offset
- AND the matrix includes viewport scaling from ChildExtent to Extent (scale 0.5 in both axes)

#### Scenario: Convert Transform2D with rotated viewport
- GIVEN a Transform2D with rotation 45° and viewport transformation
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix is returned
- AND the matrix applies group rotation before viewport scaling
- AND child shapes rendered through this matrix appear correctly rotated and scaled

#### Scenario: Convert Transform2D with non-zero ChildOffset
- GIVEN a Transform2D with:
  - ChildOffset (100, 100) [child space origin offset]
  - ChildExtent (1000, 1000) and Extent (500, 500) [2x scale down]
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix is returned
- AND the matrix translates by ChildOffset before applying viewport scale
- AND a child shape at (0, 0) in child coordinates is properly mapped to group coordinates

#### Scenario: Convert Transform2D with flipping and viewport
- GIVEN a Transform2D with flipH=true and viewport transformation
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix is returned
- AND the matrix applies horizontal flip to the entire group
- AND viewport transformation is applied correctly

#### Scenario: Convert Transform2D without viewport (regular shape)
- GIVEN a Transform2D with Offset and Extent but no ChildOffset/ChildExtent
- WHEN FromTransform2D(xfrm) is called
- THEN a matrix with translation and rotation/flip is returned
- AND no viewport scaling is applied (ChildOffset/ChildExtent are nil)

### Requirement: Hierarchical Transform Composition

The system SHALL provide utilities to compute absolute transformations in element hierarchies.

#### Scenario: Single-level transform
- GIVEN a Shape with Transform2D in a ShapeTree (no group)
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN the shape's local transform is returned

#### Scenario: Two-level nesting
- GIVEN a Shape with transform T_shape
- AND the Shape is in a GroupShape with transform T_group
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN Multiply(T_group, T_shape) is returned

#### Scenario: Three-level nesting
- GIVEN Shape in GroupB in GroupA
- WHERE GroupA has transform T_A
- AND GroupB has transform T_B
- AND Shape has transform T_S
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN Multiply(Multiply(T_A, T_B), T_S) is returned

#### Scenario: Element at root
- GIVEN a Shape directly in ShapeTree (root parameter)
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN the shape's local transform is returned (no parent transforms)

#### Scenario: Shape in group with viewport transformation
- GIVEN a Shape with local transform T_shape at child coordinates (100, 100)
- AND the Shape is in a GroupShape with TransformGroup including:
  - Offset (200, 300), Extents (1000, 800)
  - ChildOffset (0, 0), ChildExtents (2000, 1600) [0.5x scale]
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN the result includes group's position transform
- AND the result includes viewport scaling (0.5x in both axes)
- AND the result includes shape's local transform
- AND composition order is: group_position * group_rotation * group_flip * viewport * shape_transform

#### Scenario: Nested groups with multiple viewports
- GIVEN Shape in GroupB in GroupA
- WHERE GroupA has TransformGroup with viewport scaling 0.5x
- AND GroupB has TransformGroup with viewport scaling 2.0x
- AND Shape has Transform2D
- WHEN ComputeAbsoluteTransform(shape, shapeTree) is called
- THEN the result correctly composes all three transforms
- AND viewport scalings are multiplied (0.5 * 2.0 = 1.0 net scale)
- AND the final absolute coordinates are correct

### Requirement: Coordinate Conversion

The system SHALL provide utilities to convert coordinates between local and absolute coordinate spaces.

#### Scenario: Local to absolute conversion
- GIVEN a point (10, 20) in a shape's local coordinates
- AND the shape is in a group with offset (100, 200)
- WHEN LocalToAbsolute(point, shape, root) is called
- THEN the absolute coordinates (110, 220) are returned

#### Scenario: Absolute to local conversion
- GIVEN an absolute point (110, 220)
- AND a shape in a group with offset (100, 200)
- WHEN AbsoluteToLocal(point, shape, root) is called
- THEN the local coordinates (10, 20) are returned

#### Scenario: Coordinate conversion with rotation
- GIVEN a point (100, 0) in local coordinates
- AND the shape is in a group rotated 90 degrees
- WHEN LocalToAbsolute(point, shape, root) is called
- THEN the transformed point reflects the rotation

### Requirement: Transform Decomposition

The system SHALL provide utilities to extract components from transformation matrices.

#### Scenario: Decompose translate-only matrix
- GIVEN a matrix from Translate(100, 200)
- WHEN DecomposeTransform(matrix) is called
- THEN translation component is (100, 200)
- AND rotation component is 0
- AND scale component is (1.0, 1.0)

#### Scenario: Decompose composite matrix
- GIVEN a matrix from Multiply(Translate(100, 200), Rotate(π/4))
- WHEN DecomposeTransform(matrix) is called
- THEN translation component is (100, 200)
- AND rotation component is approximately π/4
- AND scale component is (1.0, 1.0)

#### Scenario: Decompose with negative scale
- GIVEN a matrix with flipH (scale -1.0 in X)
- WHEN DecomposeTransform(matrix) is called
- THEN scale component reflects the negative X scale
