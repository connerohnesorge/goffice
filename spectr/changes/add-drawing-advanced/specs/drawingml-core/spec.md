## ADDED Requirements

### Requirement: 3D Shape Support
The system SHALL support 3D shape properties including bevels, depth effects, and lighting.

#### Scenario: Create 3D shape
- WHEN a shape with 3D properties is created
- THEN bevel, depth, and contour are configurable
- AND lighting effects can be applied
- AND the shape renders with 3D appearance in PDF

### Requirement: Shape Effects
The system SHALL support shadow, glow, reflection, blur, and other visual effects on shapes.

#### Scenario: Apply shadow effect
- WHEN a shadow effect is applied to a shape
- THEN shadow direction, distance, and blur are configurable
- AND the shadow renders in presentations and PDF

### Requirement: Picture Manipulation
The system SHALL support picture cropping, rotation, and other transformations.

#### Scenario: Crop picture
- WHEN a picture is cropped
- THEN the crop area is defined
- AND the crop is lossless and reversible
- AND the picture displays with crop applied

### Requirement: Shape Connectors
The system SHALL support connector shapes with routing and attachment point management.

#### Scenario: Create connector
- WHEN a connector is created between shapes
- THEN attachment points are specified
- AND routing mode is selectable
- AND the connector route adjusts with shape movement

### Requirement: Chart Enhancement
The system SHALL support chart labels, legends, and axis properties for data visualization.

#### Scenario: Add data labels
- WHEN data labels are added to chart series
- THEN label position and format are configurable
- AND labels display in presentations and PDF
