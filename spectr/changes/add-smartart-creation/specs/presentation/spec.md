## ADDED Requirements

### Requirement: Diagram Creation
The system MUST support creating new SmartArt diagrams on a slide.

#### Scenario: Create basic diagram
- WHEN `slide.AddDiagram(layout, style, color)` is called
- THEN a `graphicFrame` is added to the slide referencing the diagram
- AND 4 new parts are added to the package: Data, Layout, Style, Colors
- AND the Data part is initialized with a root node

### Requirement: Diagram Data Manipulation
The system MUST provide an API to add nodes and text to the diagram.

#### Scenario: Add node to diagram
- WHEN `diagram.AddNode("Text content")` is called
- THEN a new `dgm:pt` (Point) is added to the `PointList`
- AND the text content is correctly serialized
