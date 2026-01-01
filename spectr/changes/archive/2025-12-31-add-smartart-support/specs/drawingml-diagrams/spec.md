# Drawingml Diagrams Specification (Delta)

## ADDED Requirements

### Requirement: Data Model Elements

The system SHALL provide element classes for SmartArt data model structures.

#### Scenario: DataModelRoot element
- GIVEN a diagram data XML file
- WHEN parsed
- THEN DataModelRoot (dgm:dataModel) element is available (root element of DiagramDataPart)
- AND PointList (dgm:ptLst) and ConnectionList (dgm:cxnLst) children are accessible

#### Scenario: Point element
- GIVEN a PointList element
- WHEN points are accessed
- THEN Point (dgm:pt) elements are available
- AND each Point has modelId, type, cxnId attributes
- AND PropertySet (dgm:prSet), ShapeProperties (dgm:spPr), and Text (dgm:t) children are available

#### Scenario: Connection element
- GIVEN a ConnectionList element
- WHEN connections are accessed
- THEN Connection (dgm:cxn) elements are available
- AND each Connection has modelId, type, srcId, destId, srcOrd, destOrd attributes
- AND srcId and destId reference Point elements by modelId

#### Scenario: PropertySet element
- GIVEN a Point element
- WHEN PropertySet is accessed
- THEN properties like presLayoutVars, presName, presStyleLbl, presStyleIdx are available

### Requirement: Layout Definition Elements

The system SHALL provide element classes for SmartArt layout algorithms.

#### Scenario: LayoutDefinition element
- GIVEN a diagram layout XML file
- WHEN parsed
- THEN LayoutDefinition (dgm:layoutDef) element is available (root element of DiagramLayoutDefinitionPart)
- AND uniqueId, minVer, defStyle attributes are accessible
- AND Title (dgm:title), Description (dgm:desc), CategoryList (dgm:catLst) children are available

#### Scenario: LayoutNode element
- GIVEN a LayoutDefinition element
- WHEN layout tree is accessed
- THEN root LayoutNode (dgm:layoutNode) is available
- AND name, styleLbl, chOrder attributes are accessible
- AND VariableList (dgm:varLst), Algorithm (dgm:alg), Shape (dgm:shape), ForEach (dgm:forEach), Choose (dgm:choose) children are available
- AND LayoutNode children can be nested recursively

#### Scenario: Algorithm element
- GIVEN a LayoutNode element
- WHEN Algorithm is accessed
- THEN type attribute contains algorithm identifier (e.g., "composite", "snake", "cycle", "hierChild")
- AND algorithm parameters are available as child elements

#### Scenario: ForEach element
- GIVEN a LayoutNode element
- WHEN ForEach is accessed
- THEN name, ref, axis, ptType, hideLastTrans attributes are available
- AND child LayoutNode elements define iteration template

#### Scenario: Choose element
- GIVEN a LayoutNode element
- WHEN Choose is accessed
- THEN name attribute is available
- AND If (dgm:if) and Else (dgm:else) children provide conditional branches
- AND each branch contains LayoutNode elements

### Requirement: Style Definition Elements

The system SHALL provide element classes for SmartArt visual styling.

#### Scenario: StyleData element
- GIVEN a diagram style XML file
- WHEN parsed
- THEN StyleData (dgm:styleData) element is available (root element of DiagramStylePart)
- AND DataModelRoot (dgm:dataModel) child may contain sample data for previews

#### Scenario: StyleLabel element
- GIVEN a StyleData element
- WHEN style labels are accessed
- THEN StyleLabel (dgm:styleLbl) elements are available
- AND each has name attribute
- AND Scene3D (dgm:scene3d), Style3D (dgm:sp3d), TextProperties (dgm:txPr), ShapeProperties (dgm:spPr), LineProperties (dgm:lnRef), FillReference (dgm:fillRef), EffectReference (dgm:effectRef), FontReference (dgm:fontRef) children define styling

#### Scenario: StyleDefinition element
- GIVEN a StyleData element
- WHEN StyleDefinition (dgm:styleDef) is accessed
- THEN Title (dgm:title), Description (dgm:desc), CategoryList (dgm:catLst), Scene3D (dgm:scene3d), StyleLabel list (dgm:styleLbl) children are available

### Requirement: Color Definition Elements

The system SHALL provide element classes for SmartArt color schemes.

#### Scenario: ColorsDefinition element
- GIVEN a diagram colors XML file
- WHEN parsed
- THEN ColorsDefinition (dgm:colorsDef) element is available (root element of DiagramColorsPart)
- AND uniqueId, minVer attributes are accessible
- AND Title (dgm:title), Description (dgm:desc), CategoryList (dgm:catLst) children are available

#### Scenario: ColorData element
- GIVEN a ColorsDefinition element
- WHEN color data is accessed
- THEN StyleLabel (dgm:styleLbl) elements are available for each category
- AND each StyleLabel contains FillReference (dgm:fillRef), LineReference (dgm:lnRef), EffectReference (dgm:effectRef), FontReference (dgm:fontRef) with color scheme colors

#### Scenario: ColorTransform element
- GIVEN a ColorData element
- WHEN color transformations are accessed
- THEN Alpha (dgm:alpha), HueOff (dgm:hueOff), Sat (dgm:sat), SatMod (dgm:satMod), Lum (dgm:lum), LumMod (dgm:lumMod), Tint (dgm:tint), Shade (dgm:shade) transformations are available

### Requirement: Diagram Parts Infrastructure

The system SHALL provide part classes for loading and saving SmartArt diagram XML files.

#### Scenario: DiagramDataPart
- GIVEN a document with SmartArt
- WHEN DiagramDataPart is accessed
- THEN part loads data*.xml file
- AND DataModel() method returns parsed diagram data model (*diagram.DataModelRoot)
- AND Save() method writes data model back to XML

#### Scenario: DiagramLayoutDefinitionPart
- GIVEN a DiagramDataPart
- WHEN layout part relationship is followed
- THEN DiagramLayoutDefinitionPart is accessible
- AND LayoutDefinition() method returns parsed layout definition (*diagram.LayoutDefinition)
- AND Save() method writes layout back to XML

#### Scenario: DiagramStylePart
- GIVEN a DiagramDataPart
- WHEN style part relationship is followed
- THEN DiagramStylePart is accessible
- AND StyleData() method returns parsed style data (*diagram.StyleData)
- AND Save() method writes style back to XML

#### Scenario: DiagramColorsPart
- GIVEN a DiagramDataPart
- WHEN colors part relationship is followed
- THEN DiagramColorsPart is accessible
- AND ColorsDefinition() method returns parsed colors definition (*diagram.ColorsDefinition)
- AND Save() method writes colors back to XML

### Requirement: Diagram Part Discovery

The system SHALL provide methods to discover diagram parts in documents.

#### Scenario: PowerPoint diagram discovery
- GIVEN a PresentationDocument
- WHEN DiagramDataParts() is called
- THEN all diagram data parts in the presentation are returned

#### Scenario: Word diagram discovery
- GIVEN a WordprocessingDocument
- WHEN DiagramDataParts() is called
- THEN all diagram data parts in the document are returned

#### Scenario: Excel diagram discovery
- GIVEN a SpreadsheetDocument
- WHEN DiagramDataParts() is called
- THEN all diagram data parts in the workbook are returned

### Requirement: Diagram Relationship Management

The system SHALL manage relationships between diagram parts.

#### Scenario: Data to layout relationship
- GIVEN a DiagramDataPart
- WHEN layout relationship is accessed
- THEN relationship type is "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramLayout"
- AND target points to DiagramLayoutDefinitionPart

#### Scenario: Data to colors relationship
- GIVEN a DiagramDataPart
- WHEN colors relationship is accessed
- THEN relationship type is "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramColors"
- AND target points to DiagramColorsPart

#### Scenario: Data to style relationship
- GIVEN a DiagramDataPart
- WHEN style relationship is accessed
- THEN relationship type is "http://schemas.openxmlformats.org/officeDocument/2006/relationships/diagramQuickStyle"
- AND target points to DiagramStylePart

### Requirement: Diagram Roundtrip Preservation

The system SHALL preserve SmartArt content when opening and saving documents.

#### Scenario: PowerPoint SmartArt roundtrip
- GIVEN a PowerPoint file with organizational chart SmartArt
- WHEN opened with goffice
- AND saved without modifications
- AND reopened
- THEN diagram data model is identical to original
- AND all four diagram parts are preserved
- AND relationships are maintained

#### Scenario: Word SmartArt roundtrip
- GIVEN a Word document with hierarchy SmartArt
- WHEN opened with goffice
- AND saved without modifications
- AND reopened
- THEN SmartArt content is preserved identically

#### Scenario: Excel SmartArt roundtrip
- GIVEN an Excel workbook with list SmartArt
- WHEN opened with goffice
- AND saved without modifications
- AND reopened
- THEN SmartArt content is preserved identically

### Requirement: Diagram Content Type Support

The system SHALL recognize diagram content types.

#### Scenario: Diagram data content type
- GIVEN a package with diagram data part
- WHEN content type is checked
- THEN type is "application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml"

#### Scenario: Diagram layout content type
- GIVEN a package with diagram layout part
- WHEN content type is checked
- THEN type is "application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml"

#### Scenario: Diagram style content type
- GIVEN a package with diagram style part
- WHEN content type is checked
- THEN type is "application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml"

#### Scenario: Diagram colors content type
- GIVEN a package with diagram colors part
- WHEN content type is checked
- THEN type is "application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml"

### Requirement: Lazy Loading of Diagram Parts

The system SHALL load diagram part XML only when accessed.

#### Scenario: Deferred data model loading
- GIVEN a DiagramDataPart
- WHEN the part is created
- THEN XML is not parsed immediately
- WHEN DataModel() is called
- THEN XML is parsed and returned
- WHEN DataModel() is called again
- THEN cached data model is returned without re-parsing

#### Scenario: Avoid parsing unused diagrams
- GIVEN a document with 10 SmartArt diagrams
- WHEN document is opened
- AND only 1 diagram is accessed
- THEN only 1 diagram's XML is parsed
- AND memory is not consumed by unused diagrams

### Requirement: DiagramPersistLayoutPart Not Supported (Phase 1)

The system SHALL NOT implement DiagramPersistLayoutPart in Phase 1.

#### Scenario: DiagramPersistLayoutPart deferred
- GIVEN the Office 2010+ DiagramPersistLayoutPart (cached layout)
- WHEN implementing Phase 1
- THEN DiagramPersistLayoutPart is explicitly out of scope
- AND it will be addressed in a future proposal when layout engine is implemented
- AND documents with persist layout parts can still be opened
- AND persist layout parts are preserved on roundtrip but not parsed

### Requirement: Multi-Version Schema Support

The system SHALL support diagram elements from all Office versions.

#### Scenario: Office 2007 diagrams
- GIVEN a diagram using 2006 schema elements
- WHEN parsed
- THEN all elements are recognized and preserved

#### Scenario: Office 2010 diagrams
- GIVEN a diagram using 2008 schema extensions
- WHEN parsed
- THEN extended elements are recognized and preserved

#### Scenario: Office 2013 diagrams
- GIVEN a diagram using 2010 schema extensions
- WHEN parsed
- THEN extended elements are recognized and preserved

#### Scenario: Office 2016/2019/365 diagrams
- GIVEN a diagram using 2016 schema extensions
- WHEN parsed
- THEN extended elements are recognized and preserved
