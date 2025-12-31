# Presentation Elements Spec Delta

## ADDED Requirements

### Requirement: GroupShape Child Access

The system SHALL provide methods to access child elements within a GroupShape.

#### Scenario: Access child shapes
- GIVEN a GroupShape with 3 child Shape elements
- WHEN Shapes() is called
- THEN a slice containing all 3 Shape elements is returned
- AND the shapes are in document order

#### Scenario: Access child pictures
- GIVEN a GroupShape with 2 child Picture elements
- WHEN Pictures() is called
- THEN a slice containing both Picture elements is returned

#### Scenario: Access nested groups
- GIVEN a GroupShape containing 2 child GroupShape elements
- WHEN GroupShapes() is called
- THEN a slice containing both child GroupShape elements is returned
- AND recursive iteration is possible (child groups can be queried for their children)

#### Scenario: Access child connection shapes
- GIVEN a GroupShape with 1 child ConnectionShape element
- WHEN ConnectionShapes() is called
- THEN a slice containing the ConnectionShape is returned

#### Scenario: Access child graphic frames
- GIVEN a GroupShape with 1 child GraphicFrame element
- WHEN GraphicFrames() is called
- THEN a slice containing the GraphicFrame is returned

### Requirement: GroupShape Child Manipulation

The system SHALL provide methods to add and remove child elements from a GroupShape.

#### Scenario: Add shape to group
- GIVEN an empty GroupShape
- WHEN AddShape() is called
- THEN a new Shape element is created
- AND the Shape is added as a child of the GroupShape
- AND the Shape is returned

#### Scenario: Add picture to group
- GIVEN an empty GroupShape
- WHEN AddPicture(relId) is called with relationship ID "rId1"
- THEN a new Picture element is created with the relationship
- AND the Picture is added as a child of the GroupShape
- AND the Picture is returned

#### Scenario: Add nested group
- GIVEN an existing GroupShape
- WHEN AddGroupShape() is called
- THEN a new GroupShape element is created
- AND the new GroupShape is added as a child
- AND the new GroupShape is returned
- AND the new GroupShape can itself contain children

#### Scenario: Add connection shape to group
- GIVEN an existing GroupShape
- WHEN AddConnectionShape() is called
- THEN a new ConnectionShape element is created
- AND the ConnectionShape is added as a child
- AND the ConnectionShape is returned

#### Scenario: Remove child from group
- GIVEN a GroupShape with 3 child shapes
- WHEN RemoveChild(shape) is called on the second shape
- THEN the shape is removed from the GroupShape
- AND Shapes() returns 2 shapes
- AND the remaining shapes are the first and third shapes

#### Scenario: Clear all children
- GIVEN a GroupShape with 5 child elements of mixed types
- WHEN Clear() is called
- THEN all child elements are removed
- AND Shapes(), Pictures(), GroupShapes() all return empty slices

### Requirement: GroupShape Properties Access

The system SHALL provide accessors for group-level properties.

#### Scenario: Access group shape properties
- GIVEN a GroupShape
- WHEN GroupShapeProperties() is called
- THEN a GroupShapeProperties element is returned
- AND the element contains transform and other visual properties

#### Scenario: Access non-visual properties
- GIVEN a GroupShape
- WHEN NonVisualGroupShapeProperties() is called
- THEN a NonVisualGroupShapeProperties element is returned
- AND the element contains ID, name, and other metadata

### Requirement: GroupShapeProperties Transform Access

The system SHALL provide access to the group-level transformation via GroupShapeProperties.

#### Scenario: Get group transform
- GIVEN a GroupShapeProperties element with a:xfrm child element
- WHEN Transform() is called
- THEN a Transform2D element is returned
- AND the transform contains offset, extent, rotation, and flip properties
- AND for groups, the transform may also contain ChildOffset and ChildExtent elements

#### Scenario: Get group transform when not present
- GIVEN a GroupShapeProperties element without a:xfrm child element
- WHEN Transform() is called
- THEN nil is returned

#### Scenario: Set group transform
- GIVEN a GroupShapeProperties element
- AND a Transform2D with rotation 45 degrees
- WHEN SetTransform(xfrm) is called
- THEN the transform is set on the GroupShapeProperties
- AND subsequent Transform() calls return the same transform

#### Scenario: Access child coordinate space
- GIVEN a GroupShapeProperties with Transform2D that has ChildOffset (100, 200) and ChildExtent (1000, 800)
- WHEN transform := props.Transform() is called
- AND chOffX, chOffY, hasChOff := transform.ChildOffset() is called
- AND chExtCx, chExtCy, hasChExt := transform.ChildExtent() is called
- THEN chOffX == 100, chOffY == 200, hasChOff == true
- AND chExtCx == 1000, chExtCy == 800, hasChExt == true
- AND these define viewport scaling for child shapes
