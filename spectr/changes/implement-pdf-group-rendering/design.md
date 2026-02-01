# Proposal: Implement PDF Group Rendering Features

## Summary

Implement picture rendering, connection shape rendering, and graphic frame rendering within the PDF group renderer.

## Background

The PDF group renderer (`pdf/drawing/group_renderer.go`) has three TODOs for missing rendering capabilities:

**Reference TODOs:**
- `pdf/drawing/group_renderer.go:159` - "TODO: Implement picture rendering (this requires image handling)"
- `pdf/drawing/group_renderer.go:170` - "TODO: Implement connection shape rendering"
- `pdf/drawing/group_renderer.go:181` - "TODO: Implement graphic frame rendering (charts, tables)"

These features are needed to fully render grouped shapes in PDF output.

## Motivation

Drawing groups can contain various element types:
- **Pictures**: Images that are part of a group
- **Connection Shapes**: Connectors between shapes (lines with special endpoints)
- **Graphic Frames**: Containers for charts, tables, and other embedded content

Without these features, grouped content is incomplete in PDF output.

## Technical Design

### 1. Picture Rendering in Groups

**Purpose:** Render images that are part of a shape group.

**Implementation:**
```go
func (r *GroupRenderer) renderPicture(pic *elements.Picture) error {
    if pic == nil {
        return nil
    }
    
    // Get the blip fill (image reference)
    blipFill := pic.GetBlipFill()
    if blipFill == nil {
        return nil
    }
    
    // Get picture bounds within the group
    bounds := r.getPictureBounds(pic)
    
    // Use ImageRenderer to render the picture
    return r.imageRenderer.RenderPicture(blipFill, bounds)
}
```

**Requirements:**
- Access the group's coordinate transformation
- Calculate picture bounds within group context
- Delegate to existing ImageRenderer

### 2. Connection Shape Rendering

**Purpose:** Render connector lines between shapes.

**Implementation:**
```go
func (r *GroupRenderer) renderConnection(conn *elements.ConnectionShape) error {
    if conn == nil {
        return nil
    }
    
    // Get connection path
    path := conn.GetPath()
    if path == nil {
        return nil
    }
    
    // Get connection style (line properties)
    style := conn.GetStyle()
    
    // Transform coordinates to group context
    transformedPath := r.transformPath(path)
    
    // Render the connector line
    return r.renderConnectorPath(transformedPath, style)
}
```

**Connector Types to Support:**
- Straight connectors
- Elbow connectors (with angles)
- Curved connectors

**Requirements:**
- Parse connection path geometry
- Handle different connector styles
- Draw line with proper styling

### 3. Graphic Frame Rendering

**Purpose:** Render charts and tables within groups.

**Implementation:**
```go
func (r *GroupRenderer) renderGraphicFrame(frame *elements.GraphicFrame) error {
    if frame == nil {
        return nil
    }
    
    // Determine frame content type
    contentType := frame.GetContentType()
    
    switch contentType {
    case ContentTypeChart:
        return r.renderChartInFrame(frame)
    case ContentTypeTable:
        return r.renderTableInFrame(frame)
    default:
        // Unknown content type, skip or render placeholder
        return r.renderPlaceholderFrame(frame)
    }
}

func (r *GroupRenderer) renderChartInFrame(frame *elements.GraphicFrame) error {
    // Get chart data
    chart := frame.GetChart()
    if chart == nil {
        return nil
    }
    
    // Get frame bounds in group context
    bounds := r.getFrameBounds(frame)
    
    // Use ChartRenderer to render within bounds
    return r.chartRenderer.Render(chart, bounds)
}
```

**Requirements:**
- Detect content type (chart vs table)
- Delegate to appropriate renderer
- Handle missing content gracefully

## Requirements

### SHALL Requirements

#### Requirement: Picture Rendering in Groups
The group renderer SHALL render pictures that are part of shape groups.

##### Scenario: Picture in Group
Given a group containing a picture element
When the group is rendered to PDF
Then the picture SHALL be displayed within the group

##### Scenario: Picture with Transform
Given a picture with position transform within a group
When rendered
Then the picture SHALL be positioned according to group coordinates

#### Requirement: Connection Shape Rendering
The group renderer SHALL render connection shapes (connectors).

##### Scenario: Straight Connector
Given a straight connector between two shapes in a group
When rendered
Then a straight line SHALL be drawn between the connection points

##### Scenario: Styled Connector
Given a connector with line style (color, width, dash)
When rendered
Then the connector SHALL be drawn with the specified style

#### Requirement: Graphic Frame Rendering
The group renderer SHALL render charts and tables within graphic frames.

##### Scenario: Chart in Frame
Given a graphic frame containing a chart
When rendered
Then the chart SHALL be displayed within the frame bounds

##### Scenario: Table in Frame
Given a graphic frame containing a table
When rendered
Then the table SHALL be displayed within the frame bounds

### SHOULD Requirements

#### Requirement: Fallback Rendering
The implementation SHOULD render a placeholder for unsupported content types.

#### Requirement: Performance
The implementation SHOULD cache renderers to avoid repeated initialization.

## Testing Strategy

### Unit Tests
- Test picture rendering with various image types
- Test connection shape rendering with different styles
- Test graphic frame content type detection
- Test coordinate transformations

### Integration Tests
- End-to-end rendering of groups with pictures
- End-to-end rendering of groups with connectors
- End-to-end rendering of groups with charts/tables

### Visual Tests
- Compare grouped content rendering with reference
- Verify proper positioning within groups

## Implementation Plan

1. Implement picture rendering in groups
2. Implement connection shape rendering
3. Implement graphic frame rendering
4. Add coordinate transformation utilities
5. Add unit tests
6. Add integration tests
7. Visual regression testing

## Related Changes

- `pdf/drawing/group_renderer.go` - Main implementation
- `pdf/drawing/image_renderer.go` - May need updates for group context
- `pdf/drawing/chart_renderer.go` - May need updates for frame rendering

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Complex coordinate transformations | Medium | Thorough testing; validation |
| Missing chart/table renderers | High | Ensure those renderers exist first |
| Performance with many group elements | Low | Optimize if needed |

## Acceptance Criteria

- [ ] Pictures render correctly within groups
- [ ] Connection shapes render with proper styling
- [ ] Charts render within graphic frames
- [ ] Tables render within graphic frames
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] Visual fidelity matches Office
