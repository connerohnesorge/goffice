## 1. Implementation
- [x] 1.1 Add `AddDiagramPart` methods to `presentation/parts/slide_part.go` that initialize the relationship.
- [x] 1.2 Implement `drawingml/diagram/builder.go` to help construct `DataModelRoot` with `PointList` and `ConnectionList`.
- [x] 1.3 Add standard layout/style/color templates to `drawingml/diagram` (embedded XML strings for common defaults).
- [x] 1.4 Update `AddDiagram` to instantiate all 4 required parts with these defaults.
- [x] 1.5 Add an integration test creating a PPTX with a new SmartArt diagram.
