# Proposal Brief: SmartArt/Diagram Support

## Priority: TIER 2 - Medium
**Impact**: 7/10  
**Effort**: 5-6 weeks  
**Status**: NOT implemented (0% coverage)

## Gap Analysis

### Current State
**What exists:**
- ✅ **Namespace defined**: `NamespaceDiagram` constant in drawingml
- ❌ **Everything else**: No elements, no API, nothing

**What's missing:**
- ❌ No diagram elements generated
- ❌ No SmartArt layout engine
- ❌ No SmartArt builder API
- ❌ No diagram styles/colors
- ❌ No conversion from data to diagram layout

### OpenXML-SDK Comparison
**Microsoft PowerPoint/Word:**
```csharp
// Insert SmartArt
shape.InsertSmartArt("org-chart", SmartArtLayout.OrgChart)
shape.SmartArt.Nodes.Add("CEO")
shape.SmartArt.Nodes[0].Nodes.Add("VP Sales")
shape.SmartArt.Nodes[0].Nodes.Add("VP Engineering")
```

**Open-XML-SDK** provides diagram element classes but NOT layout algorithms.

### Why This Is Hard
**SmartArt is complex:**
1. **Layout algorithms**: Must position shapes based on data structure and layout type
2. **Style system**: Color schemes, effects, connector styles
3. **Many layout types**: Org chart, process, cycle, hierarchy, relationship, matrix, pyramid, picture (50+ types)
4. **Data-driven rendering**: Abstract data model → concrete shape positions

### SmartArt Architecture
```
Data Model (dgm:dataModel)
  ↓
Layout Definition (dgm:layoutDef) - Algorithm
  ↓  
Color/Style (dgm:styleData, dgm:colors)
  ↓
Rendered Shapes (DrawingML)
```

## Implementation Scope

### Phase 1: Schema Generation (2 weeks)
**Goal**: Generate SmartArt element classes

1. **Update generators** to load diagram schemas:
   - `schemas_openxmlformats_org_drawingml_2006_diagram.json`
   - `schemas_microsoft_com_office_drawing_2008_diagram.json`
   - `schemas_microsoft_com_office_drawing_2010_diagram.json`

2. **Generate elements**:
   - Data model elements (dgm:dataModel, dgm:pt, dgm:cxn)
   - Layout definition elements (dgm:layoutDef)
   - Style elements (dgm:styleData, dgm:colors)

3. **Add to DrawingML**:
   - `drawingml/diagram/` new package

### Phase 2: Basic Diagram Support (2 weeks)
**Goal**: Read and preserve existing SmartArt in documents

1. **Diagram Part** support:
   - DiagramDataPart (data model XML)
   - DiagramLayoutPart (layout definition XML)
   - DiagramStylePart (style definition XML)
   - DiagramColorsPart (color transform XML)

2. **Parse and roundtrip**:
   - Open document with SmartArt → parse → save → verify preserved

### Phase 3: Simple Layout Algorithm (1-2 weeks)
**Goal**: Implement ONE simple layout (e.g., basic list)

1. **Data Model API**:
   - Create nodes, add connections
   - Build hierarchy

2. **Layout Engine** (simplified):
   - Position shapes vertically for list layout
   - Draw connectors between shapes

3. **Rendering**:
   - Convert diagram to DrawingML shapes

### Out of Scope (Future Phases)
- Full 50+ layout algorithms (incremental)
- Advanced styling and effects
- SmartArt styles catalog
- Picture-based SmartArt
- Custom layouts

## Success Criteria (Phase 1-2)

- [ ] Diagram schemas loaded in generators
- [ ] Diagram elements generated in drawingml/diagram/
- [ ] Can open PowerPoint/Word with SmartArt without errors
- [ ] SmartArt is preserved on save (roundtrip works)
- [ ] Can parse data model from existing SmartArt
- [ ] Can read diagram parts (data, layout, style, colors)

## Success Criteria (Phase 3 - if included)

- [ ] Can create simple list SmartArt programmatically
- [ ] Can add nodes to data model
- [ ] Layout algorithm positions shapes correctly
- [ ] Output renders correctly in PowerPoint/Word

## References

### OpenXML-SDK References
Look at @OpenXML-SDK/ for:
- `DocumentFormat.OpenXml.Drawing.Diagrams` namespace
- DiagramData, DiagramLayout element structures
- Point, Connection elements

### Diagram Schemas
Found in Open-XML-SDK data/schemas/:
- `schemas_openxmlformats_org_drawingml_2006_diagram.json`
- `schemas_microsoft_com_office_drawing_2008_diagram.json`
- `schemas_microsoft_com_office_drawing_2010_diagram.json`
- `schemas_microsoft_com_office_drawing_2016_11_diagram.json`
- `schemas_microsoft_com_office_drawing_2016_12_diagram.json`

## Files to Create/Modify

### Generation Phase
1. **cmd/gen-go-drawingml/main.go** (MODIFY)
   - Load diagram schemas in addition to main DrawingML schemas

2. **drawingml/diagram/** (NEW PACKAGE)
   - Generated element files

### API Phase (if doing Phase 2-3)
3. **drawingml/diagram/data_model.go** (NEW)
   - DataModel, Point, Connection wrappers

4. **drawingml/diagram/layout_engine.go** (NEW)
   - Basic layout algorithms

5. **presentation/parts/diagram_part.go** (NEW)
   - DiagramDataPart, DiagramLayoutPart, etc.

## Dependencies

- DrawingML infrastructure (exists)
- Shape rendering (exists for basic shapes)
- Generator framework (exists)

## Example Usage (Phase 3 - Future)

```go
// Create simple list SmartArt
slide := pres.AddSlide()
diagram := slide.AddDiagram(DiagramTypeList)

// Add data nodes
diagram.DataModel().AddNode("Item 1")
diagram.DataModel().AddNode("Item 2")
diagram.DataModel().AddNode("Item 3")

// Apply layout
diagram.ApplyLayout("basicList")

// Apply colors
diagram.ApplyColorScheme("colorful")

// Render to shapes
diagram.Render()

pres.Save()
```

## Technical Challenges

**1. Layout Algorithms**
Each SmartArt type has custom layout logic. Full implementation requires:
- Tree traversal for org charts
- Circular positioning for cycles
- Matrix positioning for 2D grids
- This is months of work for all 50+ layouts

**Mitigation**: Start with simplest layouts (list, basic hierarchy)

**2. Style System**
SmartArt styles are complex transformations. Full system requires:
- Color transformations (tint, shade, saturation)
- Effect cascading (shadows, glows inherit)
- Font sizing algorithms

**Mitigation**: Basic style support only in Phase 1-2

**3. Testing**
Hard to test without layout rendering.

**Mitigation**: Focus on roundtrip tests initially

## Recommendation

**For this proposal**: Focus on **Phase 1-2 only** (schema generation + roundtrip)
- 2-3 weeks effort
- Enables reading documents with SmartArt
- Lays groundwork for future layout implementation
- Separate proposal for "SmartArt Layout Engine" later

**DO NOT** attempt full SmartArt in one proposal - too large.

## Spec Capabilities to Update

- **drawingml-core**: Add diagram element requirements
- Create **drawingml-diagrams**: NEW spec for SmartArt

## Estimated Effort

**Phase 1-2 (Recommended Scope)**: 2-3 weeks
- Schema loading: 1 week
- Element generation: 1 week  
- Roundtrip testing: 1 week

**Phase 3 (Optional/Future)**: 2-3 weeks per layout type
- Simple list layout: 2 weeks
- Org chart layout: 3 weeks
- Cycle layout: 3 weeks
- etc.

**Full SmartArt Support**: 6-12 months (all layouts + styles + effects)
