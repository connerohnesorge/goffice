# Presentation Document Spec Delta

## ADDED Requirements

### Requirement: Create Group from Existing Shapes

The system SHALL provide a high-level operation to group existing shapes on a slide.

#### Scenario: Group two shapes
- GIVEN a slide with 2 Shape elements at positions (100, 100) and (200, 200)
- WHEN slide.CreateGroup(shape1, shape2) is called
- THEN a new GroupShape is created
- AND both shapes are removed from the slide's shape tree
- AND both shapes are added to the GroupShape as children
- AND the GroupShape is added to the slide's shape tree
- AND the GroupShape is returned

#### Scenario: Group shapes with coordinate adjustment
- GIVEN shape1 at absolute position (100, 150)
- AND shape2 at absolute position (300, 400)
- WHEN CreateGroup(shape1, shape2) is called
- THEN the GroupShape is positioned at (100, 150) - the top-left of bounding box
- AND shape1's local coordinates within the group are (0, 0)
- AND shape2's local coordinates within the group are (200, 250)

#### Scenario: Group mixed element types
- GIVEN a slide with:
  - 1 Shape
  - 1 Picture
  - 1 ConnectionShape
- WHEN CreateGroup(shape, picture, connector) is called
- THEN all 3 elements are added to the created GroupShape
- AND all 3 elements are removed from the slide
- AND the GroupShape contains all 3 children

#### Scenario: Group includes bounding box calculation
- GIVEN 3 shapes with positions and extents:
  - Shape1: (100, 100) size (50, 50)
  - Shape2: (200, 150) size (100, 75)
  - Shape3: (50, 200) size (80, 60)
- WHEN CreateGroup(shape1, shape2, shape3) is called
- THEN the GroupShape bounding box is:
  - Offset: (50, 100) - minimum X and Y
  - Extents: (250, 160) - covers all shapes

#### Scenario: Group preserves shape properties
- GIVEN a shape with fill color red, rotation 30°, and specific size
- WHEN the shape is added to a group via CreateGroup()
- THEN the shape retains its fill color, rotation, and size
- AND only the position coordinates are adjusted (relative to group)

### Requirement: Ungroup Shape

The system SHALL provide a high-level operation to ungroup a GroupShape back to individual elements.

#### Scenario: Ungroup simple group
- GIVEN a GroupShape containing 2 shapes
- WHEN slide.UngroupShape(group) is called
- THEN the GroupShape is removed from the slide
- AND both child shapes are added to the slide's shape tree
- AND a slice containing both shapes is returned

#### Scenario: Ungroup with coordinate conversion
- GIVEN a GroupShape at position (100, 200)
- AND the group contains a shape at local position (50, 75) relative to group
- WHEN UngroupShape(group) is called
- THEN the ungrouped shape has absolute position (150, 275)

#### Scenario: Ungroup with group transformation
- GIVEN a GroupShape rotated 45 degrees at position (100, 100)
- AND the group contains a shape at local position (50, 0)
- WHEN UngroupShape(group) is called
- THEN the ungrouped shape's transformation includes:
  - The group's rotation
  - The translated local position
  - The shape's own transformation (if any)

#### Scenario: Ungroup nested group (one level only)
- GIVEN GroupA containing GroupB and Shape1
- AND GroupB contains Shape2 and Shape3
- WHEN UngroupShape(groupA) is called
- THEN groupA is removed from the slide
- AND groupB and shape1 are added to the slide
- AND groupB still contains shape2 and shape3 (not recursively ungrouped)

#### Scenario: Ungroup returns all children
- GIVEN a GroupShape with 5 child elements of mixed types
- WHEN UngroupShape(group) is called
- THEN a slice with all 5 elements is returned
- AND the slice elements match the children that were in the group

### Requirement: Group Operation Validation

The system SHALL validate inputs and handle edge cases for grouping operations.

#### Scenario: Group with no shapes
- WHEN CreateGroup() is called with zero arguments
- THEN an error is returned
- OR an empty GroupShape is created (implementation-defined)

#### Scenario: Group with nil shapes
- WHEN CreateGroup(shape1, nil, shape2) is called
- THEN nil entries are skipped
- AND a GroupShape containing shape1 and shape2 is created

#### Scenario: Group shapes from different parents
- GIVEN shape1 in slideA
- AND shape2 in slideB (different slide)
- WHEN slideA.CreateGroup(shape1, shape2) is called
- THEN both shapes are removed from their current parents
- AND both are added to the new GroupShape
- AND the GroupShape is added to slideA

#### Scenario: Ungroup empty group
- GIVEN a GroupShape with no children
- WHEN UngroupShape(group) is called
- THEN the group is removed from the slide
- AND an empty slice is returned

#### Scenario: Ungroup non-existent group
- GIVEN a GroupShape that is not in the slide
- WHEN UngroupShape(group) is called
- THEN no error occurs (idempotent)
- AND the group is not in the slide (unchanged state)

### Requirement: Group Operation Integration

The system SHALL ensure grouping operations integrate correctly with the overall document model.

#### Scenario: Group after save/load roundtrip
- GIVEN a presentation with grouped shapes
- WHEN the presentation is saved and reopened
- AND CreateGroup() is called on the reopened shapes
- THEN the grouping operation works correctly
- AND the new group can be saved again

#### Scenario: Ungroup after save/load roundtrip
- GIVEN a presentation with a GroupShape
- WHEN the presentation is saved and reopened
- AND UngroupShape() is called on the reopened group
- THEN the ungrouping operation works correctly
- AND the ungrouped shapes have correct absolute positions

#### Scenario: Multiple group operations
- GIVEN a slide with 6 shapes
- WHEN CreateGroup(shape1, shape2, shape3) is called
- AND CreateGroup(shape4, shape5) is called
- AND CreateGroup(group1, group2, shape6) is called
- THEN 3 GroupShape elements are created
- AND the third group is a parent of the first two groups
- AND all 6 original shapes are accessible through the hierarchy

#### Scenario: Group then ungroup preserves positions
- GIVEN shapes at absolute positions P1, P2, P3
- WHEN CreateGroup(shapes) creates a group
- AND UngroupShape(group) immediately ungroups
- THEN the ungrouped shapes are at positions P1, P2, P3
- AND positions are preserved (within floating-point tolerance)

### Requirement: Group Shape Tree Consistency

The system SHALL maintain ShapeTree consistency during grouping operations.

#### Scenario: Parent references updated on group
- GIVEN shape1 with parent slideA.ShapeTree()
- WHEN CreateGroup(shape1) is called
- THEN shape1.Parent() returns the GroupShape (not ShapeTree)

#### Scenario: Child removal from ShapeTree
- GIVEN a ShapeTree with 5 shapes
- WHEN CreateGroup(shape2, shape4) is called
- THEN ShapeTree.Shapes() returns 3 shapes (shape1, shape3, shape5)
- AND the GroupShape is in ShapeTree.GroupShapes()

#### Scenario: Group added to ShapeTree
- GIVEN a ShapeTree with 3 elements
- WHEN CreateGroup(shape1, shape2) is called
- THEN ShapeTree.Children() includes the new GroupShape
- AND the GroupShape appears in the correct position in child order
