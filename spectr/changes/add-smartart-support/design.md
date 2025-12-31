# Design: SmartArt/Diagram Support Architecture

## Problem Statement

goffice has no support for SmartArt diagrams, which are present in many PowerPoint presentations, Word documents, and Excel worksheets. Currently:
- Documents with SmartArt cannot be roundtripped (content is lost)
- No diagram element classes exist
- No diagram parts infrastructure
- No way to read or parse SmartArt structure

This design addresses Phase 1: enabling roundtrip support by generating diagram elements and implementing part infrastructure.

## Current Architecture

```
DrawingML Package
├── Main elements (shapes, fills, effects)
├── Chart elements (c: namespace)
├── Picture elements (pic: namespace)
└── Diagram namespace constant ONLY (dgm:)
    └── No elements, no parts, nothing implemented
```

**Gap**: Diagram namespace defined but completely unused.

## Proposed Architecture

```
DrawingML Package
├── Main elements (shapes, fills, effects)
├── Chart elements (c: namespace)
├── Picture elements (pic: namespace)
└── Diagram elements (dgm: namespace) - NEW
    ├── Data model elements (pt, cxn, ptLst, cxnLst)
    ├── Layout elements (layoutDef, layoutNode, forEach, choose)
    ├── Style elements (styleData, styleLbl)
    └── Color elements (colorsDef, colorData)

Document Parts (Presentation/Word/Excel)
├── MainDocumentPart
├── ChartPart
└── Diagram Parts - NEW
    ├── DiagramDataPart (data model XML)
    ├── DiagramLayoutDefinitionPart (layout definition XML)
    ├── DiagramStylePart (style definition XML)
    └── DiagramColorsPart (color transform XML)
    └── DiagramPersistLayoutPart (cached layout - out of scope Phase 1)
```

## SmartArt Structure Overview

### Four-Part Architecture
SmartArt is split across four XML parts, each serving a distinct purpose:

```
┌─────────────────────┐
│  DiagramDataPart    │  Logical structure (nodes, connections)
│  data*.xml          │  "What": data model, hierarchy
└─────────────────────┘
          │
          ├────────────────────────────────────┐
          │                                    │
┌──────────────────────────────┐  ┌─────────────────────┐
│ DiagramLayoutDefinitionPart  │  │  DiagramColorsPart  │
│ layout*.xml                  │  │  colors*.xml        │
│ "How": algorithm             │  │  "Color": scheme    │
└──────────────────────────────┘  └─────────────────────┘
          │                                    │
          └────────────────┬───────────────────┘
                           │
                ┌─────────────────────┐
                │  DiagramStylePart   │
                │  quickStyle*.xml    │
                │  "Look": visual     │
                └─────────────────────┘
```

**Data Part** (dgm:dataModel - DataModelRoot):
- Nodes: Individual points in diagram (e.g., employees in org chart)
- Connections: Relationships between nodes (e.g., reporting structure)
- Properties: Text, metadata per node
- Generated Go Type: DataModelRoot

**Layout Part** (dgm:layoutDef - LayoutDefinition):
- Algorithm: Rules for positioning nodes
- Constraints: Size, spacing, alignment rules
- Variables: Calculated values for layout
- Generated Go Type: LayoutDefinition

**Style Part** (dgm:styleData - StyleData):
- Fill/stroke definitions
- Font properties
- Effect styles
- Shape styling
- Generated Go Type: StyleData

**Colors Part** (dgm:colorsDef - ColorsDefinition):
- Color scheme selection
- Color transformations (tint, shade, etc.)
- Category-based color assignment
- Generated Go Type: ColorsDefinition

### Element Hierarchy

**Data Model Elements** (dgm:dataModel - DataModelRoot):
```
dgm:dataModel → DataModelRoot (generated Go type)
├── dgm:ptLst (point list) → PointList
│   └── dgm:pt (point/node) → Point
│       ├── dgm:prSet (property set) → PropertySet
│       ├── dgm:spPr (shape properties) → ShapeProperties
│       └── dgm:t (text) → Text
└── dgm:cxnLst (connection list) → ConnectionList
    └── dgm:cxn (connection) → Connection
        ├── dgm:srcId (source point ID)
        ├── dgm:destId (destination point ID)
        └── dgm:presOf (presentation of)
```

**Layout Definition Elements** (dgm:layoutDef - LayoutDefinition):
```
dgm:layoutDef → LayoutDefinition (generated Go type)
├── dgm:title (layout name) → Title
├── dgm:desc (description) → Description
├── dgm:catLst (category list) → CategoryList
├── dgm:sampData (sample data) → SampleData
├── dgm:styleData (embedded style data) → StyleData
└── dgm:layoutNode (layout tree) → LayoutNode
    ├── dgm:varLst (variable list) → VariableList
    ├── dgm:forEach (iteration) → ForEach
    ├── dgm:choose (conditional) → Choose
    ├── dgm:alg (algorithm) → Algorithm
    ├── dgm:shape (shape template) → Shape
    └── dgm:layoutNode (recursive) → LayoutNode
```

## Design Alternatives Considered

### Alternative 1: Single Unified DiagramPart
**Approach**: Combine all four parts into one DiagramPart

**Pros**:
- Simpler API (one part instead of four)
- Fewer relationship management

**Cons**:
- Violates Office Open XML specification (uses 4 separate parts)
- Cannot roundtrip correctly (Office expects separate files)
- Makes layout/style reuse impossible (Office allows sharing styles across diagrams)

**Decision**: REJECTED - Must match Office spec exactly for compatibility

### Alternative 2: Implement Layout Algorithms in Phase 1
**Approach**: Include SmartArt rendering in initial proposal

**Pros**:
- Complete feature delivery
- Users can create SmartArt immediately

**Cons**:
- 6-12 weeks additional effort
- Complex algorithms (50+ layout types)
- Proposal becomes too large to review
- Delays delivery of roundtrip support (which is simpler and valuable alone)

**Decision**: REJECTED - Split into two phases for manageable scope

### Alternative 3: Manual Element Implementation
**Approach**: Hand-write diagram element classes instead of generating

**Pros**:
- Full control over API design
- Can optimize for Go idioms

**Cons**:
- Weeks of tedious work (100+ element types)
- Error-prone (complex nested structures)
- Generator already exists and works well
- Hard to maintain (schema updates require manual changes)

**Decision**: REJECTED - Use existing generator infrastructure

### Alternative 4: Generate Elements, Implement Parts (CHOSEN)
**Approach**: Auto-generate elements from schemas, implement four part types

**Pros**:
- Leverages existing generator (proven approach)
- Matches Office spec exactly (four parts)
- Enables roundtrip immediately
- Foundation for future layout engine
- Manageable scope (2-3 weeks)

**Cons**:
- Doesn't enable SmartArt creation (Phase 2)

**Decision**: ACCEPTED - Best balance of effort and value

## Schema Generation Details

### Schemas to Load

**Core Diagram Schema (2006)**:
- `schemas_openxmlformats_org_drawingml_2006_diagram.json`
- Elements: dataModel, layoutDef, colorsDef, styleData
- Namespaces: `http://schemas.openxmlformats.org/drawingml/2006/diagram`

**Office 2010 Extensions (2008)**:
- `schemas_microsoft_com_office_drawing_2008_diagram.json`
- Extensions: Enhanced layout capabilities
- Namespace: `http://schemas.microsoft.com/office/drawing/2008/diagram`

**Office 2013 Extensions (2010)**:
- `schemas_microsoft_com_office_drawing_2010_diagram.json`
- Extensions: Additional algorithm types
- Namespace: `http://schemas.microsoft.com/office/drawing/2010/diagram`

**Office 2016 Extensions (2016/11)**:
- `schemas_microsoft_com_office_drawing_2016_11_diagram.json`
- Extensions: Modern layout features
- Namespace: `http://schemas.microsoft.com/office/drawing/2016/11/diagram`

**Office 2019/365 Extensions (2016/12)**:
- `schemas_microsoft_com_office_drawing_2016_12_diagram.json`
- Extensions: Latest features
- Namespace: `http://schemas.microsoft.com/office/drawing/2016/12/diagram`

### Generator Modifications

**In `cmd/gen-go-drawingml/main.go`**:
```go
// Add to schema loading section
func loadDiagramSchemas() ([]*Schema, error) {
    schemaFiles := []string{
        "schemas_openxmlformats_org_drawingml_2006_diagram.json",
        "schemas_microsoft_com_office_drawing_2008_diagram.json",
        "schemas_microsoft_com_office_drawing_2010_diagram.json",
        "schemas_microsoft_com_office_drawing_2016_11_diagram.json",
        "schemas_microsoft_com_office_drawing_2016_12_diagram.json",
    }

    var schemas []*Schema
    for _, file := range schemaFiles {
        schema, err := loadSchema(file)
        if err != nil {
            return nil, fmt.Errorf("load %s: %w", file, err)
        }
        schemas = append(schemas, schema)
    }

    return schemas, nil
}

// In main generation loop
diagramSchemas, err := loadDiagramSchemas()
if err != nil {
    log.Fatalf("Failed to load diagram schemas: %v", err)
}

// Generate diagram elements
if err := generatePackage("diagram", diagramSchemas, "drawingml/diagram"); err != nil {
    log.Fatalf("Failed to generate diagram package: %v", err)
}
```

**Output**: `drawingml/diagram/` package with generated element files

### Expected Generated Elements (Sample)

**Data Model Elements**:
- `DataModelRoot` (dgm:dataModel) - root element for DiagramDataPart
- `PointList` (dgm:ptLst)
- `Point` (dgm:pt)
- `ConnectionList` (dgm:cxnLst)
- `Connection` (dgm:cxn)
- `PropertySet` (dgm:prSet)

**Layout Elements**:
- `LayoutDefinition` (dgm:layoutDef) - root element for DiagramLayoutDefinitionPart
- `LayoutNode` (dgm:layoutNode)
- `Algorithm` (dgm:alg)
- `ForEach` (dgm:forEach)
- `Choose` (dgm:choose)
- `VariableList` (dgm:varLst)

**Style Elements**:
- `StyleData` (dgm:styleData) - root element for DiagramStylePart
- `StyleLabel` (dgm:styleLbl)
- `StyleDefinition` (dgm:styleDef)

**Color Elements**:
- `ColorsDefinition` (dgm:colorsDef) - root element for DiagramColorsPart
- `ColorData` (dgm:colorData)
- `ColorTransform` (dgm:colorsTransform)

## Part Implementation Design

### Part Interface Pattern

All diagram parts follow the standard goffice part pattern:

```go
// DiagramDataPart represents the data model for a SmartArt diagram.
type DiagramDataPart struct {
    *packaging.Part
    dataModel *diagram.DataModelRoot // Lazy loaded (root element is DataModelRoot)
}

// NewDiagramDataPart creates a new diagram data part.
func NewDiagramDataPart(pkg *packaging.Package, uri string) (*DiagramDataPart, error) {
    part, err := pkg.CreatePart(uri, contentType.DiagramData)
    if err != nil {
        return nil, err
    }

    return &DiagramDataPart{Part: part}, nil
}

// DataModel returns the parsed data model, loading it if necessary.
func (p *DiagramDataPart) DataModel() (*diagram.DataModelRoot, error) {
    if p.dataModel == nil {
        reader, err := p.GetStream()
        if err != nil {
            return nil, err
        }
        defer reader.Close()

        p.dataModel = &diagram.DataModelRoot{}
        if err := xml.NewDecoder(reader).Decode(p.dataModel); err != nil {
            return nil, err
        }
    }

    return p.dataModel, nil
}

// Save writes the data model back to the part.
func (p *DiagramDataPart) Save() error {
    if p.dataModel == nil {
        return nil // Nothing to save
    }

    writer, err := p.GetStreamWriter()
    if err != nil {
        return err
    }
    defer writer.Close()

    encoder := xml.NewEncoder(writer)
    encoder.Indent("", "  ")
    return encoder.Encode(p.dataModel)
}
```

### Part Types Implementation

**DiagramDataPart** (`data*.xml`):
- Content Type: `application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml`
- Root Element: `dgm:dataModel`
- Generated Go Type: `DataModelRoot`
- Purpose: Store node/connection structure

**DiagramLayoutDefinitionPart** (`layout*.xml`):
- Content Type: `application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml`
- Root Element: `dgm:layoutDef`
- Generated Go Type: `LayoutDefinition`
- Purpose: Store layout algorithm definition

**DiagramStylePart** (`quickStyle*.xml`):
- Content Type: `application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml`
- Root Element: `dgm:styleData`
- Generated Go Type: `StyleData`
- Purpose: Store visual styling

**DiagramColorsPart** (`colors*.xml`):
- Content Type: `application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml`
- Root Element: `dgm:colorsDef`
- Generated Go Type: `ColorsDefinition`
- Purpose: Store color scheme

**DiagramPersistLayoutPart** (`drawing*.xml` - Office 2010+):
- Content Type: `application/vnd.ms-office.drawingml.diagramDrawing+xml`
- Root Element: `dsp:drawing`
- Purpose: Store cached layout (persisted positioned shapes)
- Status: OUT OF SCOPE for Phase 1 - deferred to future proposal

### Relationship Management

**From MainDocumentPart to DiagramDataPart**:
```xml
<Relationship Id="rId5"
              Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramData"
              Target="diagrams/data1.xml" />
```

**From DiagramDataPart to Other Diagram Parts**:
```xml
<!-- Layout relationship -->
<Relationship Id="rId1"
              Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramLayout"
              Target="layout1.xml" />

<!-- Colors relationship -->
<Relationship Id="rId2"
              Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramColors"
              Target="colors1.xml" />

<!-- Style relationship -->
<Relationship Id="rId3"
              Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramQuickStyle"
              Target="quickStyle1.xml" />
```

**Relationship Constants** (add to `packaging/relationship_types.go`):
```go
const (
    // Diagram relationship types
    RelationshipTypeDiagramData       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramData"
    RelationshipTypeDiagramLayout     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramLayout"
    RelationshipTypeDiagramColors     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramColors"
    RelationshipTypeDiagramQuickStyle = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramQuickStyle"
)
```

### Part Discovery in Documents

**In PresentationDocument**:
```go
// DiagramDataParts returns all diagram data parts in the presentation.
func (d *PresentationDocument) DiagramDataParts() ([]*DiagramDataPart, error) {
    var parts []*DiagramDataPart

    for _, slide := range d.PresentationPart.Slides() {
        rels := slide.Part.Relationships()
        for _, rel := range rels {
            if rel.Type == RelationshipTypeDiagramData {
                part, err := NewDiagramDataPart(d.Package, rel.Target)
                if err != nil {
                    return nil, err
                }
                parts = append(parts, part)
            }
        }
    }

    return parts, nil
}
```

Similar methods in `WordprocessingDocument` and `SpreadsheetDocument`.

## PDF Rendering (Phase 1 - Minimal)

### Placeholder Renderer

For Phase 1, render a simple placeholder to indicate diagram presence:

```go
// DiagramRenderer renders SmartArt diagrams to PDF (Phase 1: placeholder only).
type DiagramRenderer struct{}

// Render draws a placeholder bounding box for the diagram.
func (r *DiagramRenderer) Render(ctx *core.RenderingContext, diagram *DiagramDataPart) error {
    // Get diagram bounding box from first shape reference
    // (actual bounds would come from parent shape/anchor)
    bounds := r.getBounds(diagram)

    page := ctx.Page

    // Draw dashed border
    page.PushState()
    page.SetStrokeColor(color.Gray{128})
    page.SetLineWidth(1.0)
    page.SetLineDashPattern([]float64{5, 3}, 0)
    page.DrawRectangle(bounds.X, bounds.Y, bounds.Width, bounds.Height)
    page.PopState()

    // Draw "SmartArt" label
    page.DrawText("SmartArt Diagram", bounds.X+10, bounds.Y+10, defaultFont, 10)

    return nil
}
```

**Rationale**:
- Prevents blank space in PDFs
- Indicates diagram location for users
- Full rendering deferred to Phase 2 (requires layout engine)

## Testing Strategy

### Unit Tests (Fast, Isolated)

**Element Parsing Tests** (`drawingml/diagram/diagram_test.go`):
```go
func TestDataModelParsing(t *testing.T) {
    xml := `
    <dgm:dataModel xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram">
        <dgm:ptLst>
            <dgm:pt modelId="{12345}">
                <dgm:prSet />
                <dgm:t>
                    <a:p>
                        <a:r>
                            <a:t>Node 1</a:t>
                        </a:r>
                    </a:p>
                </dgm:t>
            </dgm:pt>
        </dgm:ptLst>
        <dgm:cxnLst />
    </dgm:dataModel>
    `

    var model diagram.DataModelRoot
    err := xml.Unmarshal([]byte(xml), &model)
    assert.NoError(t, err)
    assert.Len(t, model.PointList.Points, 1)
    assert.Equal(t, "{12345}", model.PointList.Points[0].ModelID)
}
```

**Part Roundtrip Tests** (`presentation/parts/diagram_part_test.go`):
```go
func TestDiagramDataPartRoundtrip(t *testing.T) {
    // Create part with sample data
    pkg := packaging.NewPackage()
    part, err := NewDiagramDataPart(pkg, "/diagrams/data1.xml")
    assert.NoError(t, err)

    // Set data model
    model := &diagram.DataModelRoot{ /* ... */ }
    part.dataModel = model

    // Save
    err = part.Save()
    assert.NoError(t, err)

    // Reload
    part2, err := NewDiagramDataPart(pkg, "/diagrams/data1.xml")
    assert.NoError(t, err)

    model2, err := part2.DataModel()
    assert.NoError(t, err)

    // Compare
    assert.Equal(t, model, model2)
}
```

### Integration Tests (Realistic)

**Document Roundtrip Tests**:
```go
func TestPowerPointSmartArtRoundtrip(t *testing.T) {
    // Open PowerPoint with SmartArt org chart
    doc, err := presentation.Open("testdata/smartart/org-chart.pptx")
    assert.NoError(t, err)

    // Verify diagram parts exist
    diagrams, err := doc.DiagramDataParts()
    assert.NoError(t, err)
    assert.Len(t, diagrams, 1)

    // Verify data model parses
    model, err := diagrams[0].DataModel()
    assert.NoError(t, err)
    assert.NotNil(t, model)

    // Save to new file
    err = doc.SaveAs("testdata/smartart/org-chart-output.pptx")
    assert.NoError(t, err)

    // Reopen and verify
    doc2, err := presentation.Open("testdata/smartart/org-chart-output.pptx")
    assert.NoError(t, err)

    diagrams2, err := doc2.DiagramDataParts()
    assert.NoError(t, err)
    assert.Len(t, diagrams2, 1)

    // Compare XML (should be identical)
    compareXML(t, diagrams[0], diagrams2[0])
}
```

**Test Documents** (create in `testdata/smartart/`):
- `org-chart.pptx` - Organizational chart
- `process-flow.pptx` - Process diagram
- `cycle-diagram.pptx` - Cycle diagram
- `hierarchy.docx` - Hierarchical diagram in Word
- `list-diagram.xlsx` - List diagram in Excel

### Office Compatibility Validation

**Roundtrip Validation**:
1. Open test document in goffice
2. Save without modifications
3. Open saved document in Microsoft Office
4. Verify SmartArt appears and is editable
5. Compare original and saved XML (byte-for-byte where possible)

**Test Matrix**:
| Office Version | Diagram Type | Application | Status |
|---------------|--------------|-------------|--------|
| Office 2010   | Org Chart    | PowerPoint  | Test   |
| Office 2013   | Process      | PowerPoint  | Test   |
| Office 2016   | Cycle        | Word        | Test   |
| Office 2019   | Hierarchy    | Excel       | Test   |
| Office 365    | List         | PowerPoint  | Test   |

## Performance Considerations

### Lazy Loading
- Parts load XML only on first access (following existing pattern)
- Avoids parsing unused diagram parts
- Reduces memory footprint for documents with multiple SmartArt diagrams

### Element Generation Optimization
- Generator creates optimized struct tags for XML marshaling
- Minimal allocations during parsing
- Reuse of existing DrawingML color/style parsing code

### Roundtrip Efficiency
- XML preserved as-is when not modified
- No re-serialization if DataModel() never called
- Efficient relationship tracking

## Migration Path

### For Users

**Immediate (Phase 1)**:
- Open documents with SmartArt without errors
- Save documents preserving SmartArt content
- Read diagram structure programmatically

**Future (Phase 2)**:
- Create SmartArt programmatically
- Modify existing SmartArt
- Render SmartArt to PDF with full fidelity

### For Developers

**Phase 1 Foundation**:
- Diagram element classes available
- Part infrastructure in place
- Relationship management working

**Phase 2 Builds On**:
- Layout engine uses existing elements
- Creation API leverages parts
- PDF rendering uses existing drawing infrastructure

**No Breaking Changes**:
- Phase 2 is purely additive
- Phase 1 code remains unchanged
- Existing tests continue to pass

## Risks & Mitigation

### Risk: Schema Generation Failures
**Impact**: High (no elements generated)
**Probability**: Low (schemas follow standard format)
**Mitigation**:
- Test generator with diagram schemas early (Week 1, Day 1)
- Diagram schemas similar to chart schemas (already working)
- Fallback: Manual adjustments to generator if needed

### Risk: Part Relationship Complexity
**Impact**: Medium (roundtrip failures)
**Probability**: Medium (4 interconnected parts)
**Mitigation**:
- Follow proven pattern from ChartPart (already handles multiple relationships)
- Comprehensive relationship tests
- Test with real Office documents early

### Risk: Office Version Incompatibilities
**Impact**: Medium (some diagrams don't roundtrip)
**Probability**: Low (using official schemas)
**Mitigation**:
- Load all schema versions (2006, 2008, 2010, 2016)
- Test with documents from each Office version
- Validate with Office compatibility checker

### Risk: XML Namespace Conflicts
**Impact**: Low (parsing errors)
**Probability**: Low (namespaces well-defined)
**Mitigation**:
- Use proper namespace constants
- Test namespace prefixes with variations (dgm, dgm14, etc.)

## Success Metrics

- [ ] All diagram schemas load without errors
- [ ] 100+ diagram element classes generated
- [ ] All four part types implemented
- [ ] 10+ roundtrip tests passing
- [ ] Documents from Office 2010/2013/2016/2019/365 roundtrip correctly
- [ ] No XML loss on save (byte-for-byte comparison where possible)
- [ ] Zero regressions in existing tests
- [ ] Godoc coverage 100% for new code

## Future Enhancements (Phase 2)

These are explicitly OUT OF SCOPE for this proposal:

1. **Layout Engine**: Implement algorithms for positioning nodes
2. **SmartArt Creation API**: Programmatic diagram construction
3. **Full PDF Rendering**: Render positioned shapes to PDF
4. **Style Customization**: API for modifying diagram styles
5. **Custom Layouts**: Support for user-defined layout algorithms

Each of these requires a separate proposal with detailed design.
