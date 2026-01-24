# Design: Advanced DrawingML Support

## Overview
Extends DrawingML support with advanced shape properties, SmartArt, connectors, and 3D capabilities matching Open-XML-SDK.

## Core Types & Interfaces

### Shape Enhancement
```go
type ShapeEffects interface {
    ApplyShadow(*Shadow) error
    ApplyReflection(*Reflection) error
    ApplyGlow(*Glow) error
    ApplyBevel(*Bevel) error
}

type Transform3D struct {
    RotationX float64
    RotationY float64
    RotationZ float64
    Depth     int
}
```

### SmartArt Support
```go
type SmartArt struct {
    DataModel   *SmartArtDataModel
    Shapes      []Shape
    Connections []SmartArtConnection
}

type SmartArtNode struct {
    ID       string
    Text     string
    Children []*SmartArtNode
    Shape    Shape
}
```

## Architectural Decisions

### 1. Shape Hierarchy
**Decision**: Base interface with composition for effects

**Rationale**:
- Reduces duplication
- Maintains flexibility
- Matches ECMA-376 model

### 2. Effect Pipeline
**Decision**: Chain-of-responsibility pattern

**Rationale**:
- Multiple effects per shape
- Order matters
- Composable and serializable

### 3. SmartArt Processing
**Decision**: Lazy layout resolution with caching

**Rationale**:
- Expensive to parse layouts
- Documents reuse layouts
- Avoid redundant work

## Implementation Strategy

### Phase 1: Core Effects (1-2 weeks)
- [ ] Shadow, glow, reflection, bevel
- [ ] Effect pipeline
- [ ] Serialization
- [ ] PDF rendering

### Phase 2: SmartArt (2 weeks)
- [ ] Data model parsing
- [ ] Layout resolution
- [ ] Node positioning
- [ ] Connection rendering

### Phase 3: Connectors (1 week)
- [ ] Connector routing
- [ ] Arrow properties
- [ ] Callout shapes

### Phase 4: 3D Support (2 weeks)
- [ ] Transformation matrices
- [ ] Perspective projection
- [ ] Extrusion handling

## Performance Considerations

1. Effect caching by shape ID
2. Layout definition caching
3. Lazy rendering
4. Connector path caching
