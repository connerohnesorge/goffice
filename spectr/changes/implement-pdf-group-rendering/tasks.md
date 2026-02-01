# Tasks: Implement PDF Group Rendering Features

## Picture Rendering Implementation

- [ ] 1.1 Implement renderPicture method in GroupRenderer
- [ ] 1.2 Add getPictureBounds helper for coordinate transformation
- [ ] 1.3 Integrate with ImageRenderer for actual rendering
- [ ] 1.4 Handle picture transforms within group context
- [ ] 1.5 Test with various image types (PNG, JPEG)

## Connection Shape Implementation

- [ ] 2.1 Implement renderConnection method
- [ ] 2.2 Implement getPath for connection geometry
- [ ] 2.3 Implement transformPath for group coordinates
- [ ] 2.4 Add support for straight connectors
- [ ] 2.5 Add support for elbow connectors
- [ ] 2.6 Add support for curved connectors
- [ ] 2.7 Implement connector styling (color, width, dash)

## Graphic Frame Implementation

- [ ] 3.1 Implement renderGraphicFrame method
- [ ] 3.2 Implement content type detection (chart vs table)
- [ ] 3.3 Implement renderChartInFrame delegation
- [ ] 3.4 Implement renderTableInFrame delegation
- [ ] 3.5 Implement renderPlaceholderFrame for unsupported types
- [ ] 3.6 Add getFrameBounds helper

## Coordinate Transformation

- [ ] 4.1 Create coordinate transformation utilities
- [ ] 4.2 Implement local to group coordinate conversion
- [ ] 4.3 Implement group to page coordinate conversion
- [ ] 4.4 Handle rotation and scaling in transforms
- [ ] 4.5 Test transform accuracy

## Unit Tests

- [ ] 5.1 Create test: Picture rendering in group
- [ ] 5.2 Create test: Picture with transform in group
- [ ] 5.3 Create test: Straight connector rendering
- [ ] 5.4 Create test: Elbow connector rendering
- [ ] 5.5 Create test: Connector styling
- [ ] 5.6 Create test: Graphic frame with chart
- [ ] 5.7 Create test: Graphic frame with table
- [ ] 5.8 Create test: Coordinate transformation
- [ ] 5.9 Create test: Unsupported content placeholder

## Integration Tests

- [ ] 6.1 Create test: Group with multiple pictures
- [ ] 6.2 Create test: Group with connectors between shapes
- [ ] 6.3 Create test: Group with chart in graphic frame
- [ ] 6.4 Create test: Complex group with mixed content
- [ ] 6.5 Create test: Nested groups rendering

## Visual/Regression Tests

- [ ] 7.1 Create visual test: Picture in group vs reference
- [ ] 7.2 Create visual test: Connectors vs reference
- [ ] 7.3 Create visual test: Chart in group vs reference
- [ ] 7.4 Compare output with Office PDF export

## Test Fixtures

- [ ] 8.1 Create test document: Group with picture
- [ ] 8.2 Create test document: Group with connector
- [ ] 8.3 Create test document: Group with chart
- [ ] 8.4 Create test document: Complex group (mixed content)
- [ ] 8.5 Create test images for picture tests

## Documentation

- [ ] 9.1 Update pdf/AGENTS.md with group rendering details
- [ ] 9.2 Add code comments for coordinate transformations
- [ ] 9.3 Document supported group content types
- [ ] 9.4 Add troubleshooting guide for group rendering

## Verification

- [ ] 10.1 Run all pdf drawing tests - ensure no regressions
- [ ] 10.2 Verify test coverage >90% for changed code
- [ ] 10.3 Visual comparison with Excel/Word PDF export
- [ ] 10.4 Performance test with complex groups
