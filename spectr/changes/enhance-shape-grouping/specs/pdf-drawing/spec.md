# PDF Drawing Spec Delta

## ADDED Requirements

### Requirement: Group Shape Rendering

The system SHALL render GroupShape elements to PDF with correct coordinate transformations.

#### Scenario: Render simple group
- GIVEN a GroupShape containing 2 Shape elements
- WHEN the GroupShape is rendered to PDF
- THEN PDF graphics state is saved before rendering (q operator)
- AND the group's transformation matrix is applied to PDF CTM
- AND both child shapes are rendered
- AND PDF graphics state is restored after rendering (Q operator)

#### Scenario: Render group with transformation
- GIVEN a GroupShape with rotation 45 degrees and offset (100, 200)
- AND the group contains 1 rectangle shape
- WHEN the GroupShape is rendered to PDF
- THEN the group's transform is applied via PDF cm operator
- AND the rectangle renders in the group's coordinate space
- AND the final position reflects both group transform and shape's local position

#### Scenario: Render nested groups
- GIVEN GroupShape A containing GroupShape B
- AND GroupShape B contains Shape C
- WHERE GroupA has transform T_A
- AND GroupB has transform T_B
- WHEN GroupA is rendered to PDF
- THEN graphics state is saved for GroupA
- AND transform T_A is applied
- AND GroupB is recursively rendered (saves state, applies T_B)
- AND Shape C is rendered
- AND graphics states are restored in reverse order (B then A)

#### Scenario: Render group with mixed children
- GIVEN a GroupShape containing:
  - 1 Shape
  - 1 Picture
  - 1 ConnectionShape
  - 1 nested GroupShape
- WHEN the GroupShape is rendered to PDF
- THEN all 4 children are rendered in document order
- AND each child renderer is invoked appropriately

#### Scenario: Render empty group
- GIVEN a GroupShape with no children
- WHEN the GroupShape is rendered to PDF
- THEN graphics state is saved and restored
- AND no child rendering occurs
- AND no error is raised

### Requirement: Graphics State Management for Groups

The system SHALL correctly manage PDF graphics state stack during group rendering.

#### Scenario: Graphics state isolation
- GIVEN a GroupShape with a red fill color set
- AND the group contains a shape that sets blue fill
- WHEN the group is rendered
- AND a subsequent shape outside the group is rendered
- THEN the subsequent shape does not inherit the blue fill
- AND the graphics state is properly isolated

#### Scenario: Stack depth for nested groups
- GIVEN 5 levels of nested GroupShapes
- WHEN the top-level group is rendered
- THEN the graphics state stack depth reaches 5
- AND all states are properly restored
- AND the final stack depth returns to 0

#### Scenario: Transform composition via CTM
- GIVEN GroupA with translate (100, 100)
- AND GroupB (nested in A) with rotate 45°
- AND Shape C (in B) at local position (10, 10)
- WHEN GroupA is rendered
- THEN Shape C's final position is:
  - Translated by (100, 100)
  - Then rotated 45°
  - Then offset by rotated (10, 10)

### Requirement: Coordinate System Handling

The system SHALL correctly handle coordinate system conversions during group rendering.

#### Scenario: EMU to PDF points conversion
- GIVEN a GroupShape with offset in EMU units (914400, 1828800)
- WHEN the transform is applied to PDF
- THEN the offset is converted to PDF points (72, 144)

#### Scenario: Origin conversion
- GIVEN DrawingML coordinates with top-left origin
- AND PDF coordinates with bottom-left origin
- WHEN a group transform is applied
- THEN Y-coordinates are correctly flipped relative to page height

#### Scenario: Rotation direction
- GIVEN a GroupShape with rotation 5400000 (90° clockwise in DrawingML)
- WHEN the transform is applied to PDF
- THEN the rotation is correctly mapped to PDF coordinate system

### Requirement: Group Rendering Error Handling

The system SHALL handle errors during group rendering gracefully.

#### Scenario: Missing transform
- GIVEN a GroupShape with nil GroupShapeProperties or nil Transform
- WHEN the GroupShape is rendered
- THEN rendering proceeds with identity transform (no transformation applied)
- AND child shapes are rendered normally

#### Scenario: Invalid child element
- GIVEN a GroupShape with an unrecognized child element type
- WHEN the GroupShape is rendered
- THEN recognized children are rendered successfully
- AND the unrecognized element is skipped
- AND rendering continues without error

#### Scenario: Rendering error in child
- GIVEN a GroupShape with 3 child shapes
- AND the second shape fails to render (returns error)
- WHEN the GroupShape is rendered
- THEN the first shape renders successfully
- AND an error is returned after attempting all children
- AND the third shape still attempts to render

### Requirement: Group Renderer Integration

The system SHALL integrate group rendering into the overall PDF rendering pipeline.

#### Scenario: Slide with mixed content
- GIVEN a slide containing:
  - 2 Shape elements
  - 1 Picture element
  - 1 GroupShape element
  - 1 ConnectionShape element
- WHEN the slide is rendered to PDF
- THEN all elements are rendered in document order
- AND the GroupRenderer is invoked for the GroupShape
- AND appropriate renderers are used for other element types

#### Scenario: Group renderer receives correct context
- GIVEN a GroupShape to be rendered
- WHEN the GroupRenderer is invoked
- THEN the RenderingContext is passed correctly
- AND the Page API is accessible
- AND other renderers (Shape, Picture, etc.) can be invoked for children

#### Scenario: Recursive rendering context
- GIVEN nested GroupShapes
- WHEN rendering occurs
- THEN the same RenderingContext is used throughout the hierarchy
- AND transform state is maintained via PDF graphics state stack
- AND font, color, and other state persist correctly
