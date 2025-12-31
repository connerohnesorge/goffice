# Design: Extension Schema Loading Architecture

## Problem Statement

The code generators for WordprocessingML, SpreadsheetML, and PresentationML currently filter schemas to load only `main.json` files, skipping over 100 extension schemas that define Office 2010-2024 features. This design limitation prevents goffice from supporting modern Office documents, causing data loss and limiting real-world applicability.

## Current Architecture

### Generator Schema Loading (Problematic)

```go
// cmd/gen-go-wordprocessing/main.go (current)
schemasDir := "Open-XML-SDK/data/schemas"
files, _ := os.ReadDir(schemasDir)

for _, f := range files {
    if strings.Contains(f.Name(), "main.json") {
        collectTypes(filepath.Join(schemasDir, f.Name()))
    }
    // All extension schemas skipped!
}
```

**Problems:**
- Hardcoded filter for "main.json" only
- No mechanism to load extension schemas
- No namespace handling for extensions
- No version metadata extracted

### Schema File Naming Patterns

Extension schemas follow predictable patterns:

**WordProcessing Extensions:**
- `schemas_microsoft_com_office_word_2010_wordml.json`
- `schemas_microsoft_com_office_word_2012_wordml.json`
- `schemas_microsoft_com_office_word_2015_wordml_symex.json`
- `schemas_microsoft_com_office_word_2016_wordml_cid.json`
- `schemas_microsoft_com_office_word_2018_wordml.json`
- `schemas_microsoft_com_office_word_2020_wordml_sdtdatahash.json`
- `schemas_microsoft_com_office_word_2023_wordml_word16du.json`
- `schemas_microsoft_com_office_word_2024_wordml_sdtformatlock.json`

**Drawing Extensions:**
- `schemas_microsoft_com_office_drawing_2008_diagram.json`
- `schemas_microsoft_com_office_drawing_2010_main.json`
- `schemas_microsoft_com_office_drawing_2013_main_command.json`
- `schemas_microsoft_com_office_drawing_2017_decorative.json`
- `schemas_microsoft_com_office_drawing_2018_animation_model3d.json`

**Spreadsheet Extensions:**
- `schemas_microsoft_com_office_spreadsheetml_2010_11_main.json`
- `schemas_microsoft_com_office_spreadsheetml_2015_02_main.json`
- `schemas_microsoft_com_office_spreadsheetml_2016_*.json` (multiple)
- `schemas_microsoft_com_office_spreadsheetml_2017_richdata.json`
- `schemas_microsoft_com_office_spreadsheetml_2024_pivotAutoRefresh.json`

**Presentation Extensions:**
- `schemas_microsoft_com_office_powerpoint_2010_main.json`
- `schemas_microsoft_com_office_powerpoint_2012_roamingSettings.json`
- `schemas_microsoft_com_office_powerpoint_2015_main.json`
- `schemas_microsoft_com_office_powerpoint_2020_02_main.json`
- `schemas_microsoft_com_office_powerpoint_2021_06_main.json`
- `schemas_microsoft_com_office_powerpoint_2022_03_main.json`
- `schemas_microsoft_com_office_powerpoint_2022_08_main.json`
- `schemas_microsoft_com_office_powerpoint_2023_02_main.json`

**Note on Version Numbering:**
- Word and Excel extensions typically use year-only schema names (2010, 2013, 2024, 2025)
- PowerPoint extensions often use year-month schema names (2020_02, 2022_03, 2023_02)
- Monthly versions within a year (e.g., 2022_03 and 2022_08) map to the same Office version constant (Office2022)
- This approach ensures all features from a given Office release year are included under one version enum value

## Proposed Architecture

### Enhanced Generator Schema Loading

```go
// cmd/gen-go-wordprocessing/main.go (proposed)
schemasDir := "Open-XML-SDK/data/schemas"
files, _ := os.ReadDir(schemasDir)

var schemas []SchemaFile

for _, f := range files {
    name := f.Name()

    // Load main WordProcessingML schema
    if strings.Contains(name, "wordprocessingml") &&
       strings.Contains(name, "main.json") {
        schemas = append(schemas, parseSchemaFile(name))
    }

    // Load Word extension schemas (Office 2010-2024)
    if strings.Contains(name, "microsoft_com_office_word") &&
       strings.HasSuffix(name, ".json") {
        schemas = append(schemas, parseSchemaFile(name))
    }

    // Load drawing extension schemas (shared across apps)
    if strings.Contains(name, "microsoft_com_office_drawing") &&
       strings.HasSuffix(name, ".json") {
        schemas = append(schemas, parseSchemaFile(name))
    }
}

// Sort by version (main first, then 2010, 2013, etc.)
sort.Slice(schemas, func(i, j int) bool {
    return schemas[i].Version < schemas[j].Version
})

// Load types from all schemas
for _, schema := range schemas {
    collectTypesWithVersion(schema)
}
```

### SchemaFile Metadata Structure

```go
type SchemaFile struct {
    Path            string
    Namespace       string
    NamespacePrefix string
    Version         FileFormatVersion
    TargetApp       string // "word", "excel", "powerpoint", "drawing"
}

func parseSchemaFile(filename string) SchemaFile {
    sf := SchemaFile{Path: filename}

    // Extract version from filename
    // "schemas_microsoft_com_office_word_2010_wordml.json" -> 2010
    // "schemas_microsoft_com_office_powerpoint_2022_03_main.json" -> 2022
    // Monthly versions (YYYY_MM) map to the year's Office version
    if matches := versionRegex.FindStringSubmatch(filename); len(matches) > 1 {
        year, _ := strconv.Atoi(matches[1])
        sf.Version = yearToVersion(year)
    }

    // Determine target app from filename
    if strings.Contains(filename, "_word_") {
        sf.TargetApp = "word"
    } else if strings.Contains(filename, "_spreadsheetml_") {
        sf.TargetApp = "excel"
    } else if strings.Contains(filename, "_powerpoint_") {
        sf.TargetApp = "powerpoint"
    } else if strings.Contains(filename, "_drawing_") {
        sf.TargetApp = "drawing"
    }

    // Load schema to get actual namespace
    schema := loadSchemaJSON(filename)
    sf.Namespace = schema.TargetNamespace
    sf.NamespacePrefix = derivePrefix(sf.Namespace, sf.Version)

    return sf
}

func derivePrefix(namespace string, version FileFormatVersion) string {
    // "http://schemas.microsoft.com/office/word/2010/wordml" -> "w14"
    // "http://schemas.microsoft.com/office/drawing/2010/main" -> "a14"

    // Extract year and base prefix
    // Office 2010 = 14, 2013 = 15, 2016 = 16, etc.
    versionSuffix := map[FileFormatVersion]string{
        Office2010: "14",
        Office2013: "15",
        Office2016: "16",
        Office2019: "19",
        Office2021: "21",
        Office2022: "22",
        Office2023: "23",
        Office2024: "24",
        Office2025: "25",
        Microsoft365: "365",
    }

    basePrefix := "w" // or "x", "p", "a" based on namespace
    if strings.Contains(namespace, "/spreadsheetml/") {
        basePrefix = "x"
    } else if strings.Contains(namespace, "/powerpoint/") {
        basePrefix = "p"
    } else if strings.Contains(namespace, "/drawing/") {
        basePrefix = "a"
    }

    return basePrefix + versionSuffix[version]
}

func yearToVersion(year int) FileFormatVersion {
    // Map schema year to FileFormatVersion
    // Note: For monthly versions (e.g., 2022_03), extract the year part
    switch year {
    case 2007:
        return Office2007
    case 2010:
        return Office2010
    case 2013:
        return Office2013
    case 2016:
        return Office2016
    case 2019:
        return Office2019
    case 2021:
        return Office2021
    case 2022:
        return Office2022
    case 2023:
        return Office2023
    case 2024:
        return Office2024
    case 2025:
        return Office2025
    default:
        // For schemas between major releases, round down
        // e.g., 2020 -> Office2019, 2018 -> Office2016, etc.
        if year >= 2025 {
            return Office2025
        } else if year >= 2024 {
            return Office2024
        } else if year >= 2023 {
            return Office2023
        } else if year >= 2022 {
            return Office2022
        } else if year >= 2021 {
            return Office2021
        } else if year >= 2019 {
            return Office2019
        } else if year >= 2016 {
            return Office2016
        } else if year >= 2013 {
            return Office2013
        } else if year >= 2010 {
            return Office2010
        }
        return Office2007
    }
}
```

## Design Alternatives Considered

### Alternative 1: Load All JSON Files Indiscriminately
**Approach:** Load every .json file in schemas directory

**Pros:**
- Simplest code
- Catches any schema additions automatically

**Cons:**
- Loads irrelevant schemas (metadata, content types, VML, etc.)
- Namespace conflicts from unrelated schemas
- Harder to debug generation issues

**Decision:** REJECTED - Too broad, generates noise

### Alternative 2: Hardcoded Schema List
**Approach:** Maintain explicit list of schemas to load

```go
var wordSchemas = []string{
    "main.json",
    "schemas_microsoft_com_office_word_2010_wordml.json",
    "schemas_microsoft_com_office_word_2012_wordml.json",
    // ... 20 more lines
}
```

**Pros:**
- Explicit control
- No surprises

**Cons:**
- Maintenance burden (update for each new Office version)
- Easy to forget schemas
- Duplicate logic across generators

**Decision:** REJECTED - High maintenance, error-prone

### Alternative 3: Pattern-Based Loading (CHOSEN)
**Approach:** Use filename patterns to identify relevant schemas

**Pros:**
- Automatic discovery of extension schemas
- Self-documenting (patterns show intent)
- Low maintenance (works for future Office versions)
- Scoped to relevant schemas only

**Cons:**
- Slightly more complex than hardcoded list
- Depends on Microsoft's naming conventions

**Decision:** ACCEPTED - Best balance of automation and control

### Alternative 4: Runtime Schema Loading
**Approach:** Load schemas at runtime when opening documents

**Pros:**
- Smaller generated code
- Only load needed schemas

**Cons:**
- Runtime overhead
- Requires bundling schema files with application
- More complex error handling
- Breaks pure-Go compilation

**Decision:** REJECTED - Violates goffice design principles (compile-time generation)

## Namespace Handling

### Extension Namespace Mapping

Extension elements use different XML namespaces than main schemas:

**Main Namespaces:**
- WordprocessingML: `http://schemas.openxmlformats.org/wordprocessingml/2006/main` (prefix: `w`)
- SpreadsheetML: `http://schemas.openxmlformats.org/spreadsheetml/2006/main` (prefix: `x`)
- PresentationML: `http://schemas.openxmlformats.org/presentationml/2006/main` (prefix: `p`)
- DrawingML: `http://schemas.openxmlformats.org/drawingml/2006/main` (prefix: `a`)

**Extension Namespaces (Word examples):**
- Office 2010: `http://schemas.microsoft.com/office/word/2010/wordml` (prefix: `w14`)
- Office 2013: `http://schemas.microsoft.com/office/word/2012/wordml` (prefix: `w15`)
- Office 2016: `http://schemas.microsoft.com/office/word/2015/wordml/symex` (prefix: `w16`)

**Note:** Version suffix doesn't match year directly (2010=14, 2013=15, etc.) because Microsoft skipped version 13.

### Generated Namespace Constants

```go
// openxml/namespace.go (additions)

const (
    // Existing main namespaces
    NamespaceWordprocessingML = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"

    // Word extension namespaces
    NamespaceWord2010 = "http://schemas.microsoft.com/office/word/2010/wordml"
    NamespaceWord2012 = "http://schemas.microsoft.com/office/word/2012/wordml"
    NamespaceWord2015 = "http://schemas.microsoft.com/office/word/2015/wordml/symex"
    NamespaceWord2016 = "http://schemas.microsoft.com/office/word/2016/wordml/cid"
    NamespaceWord2018 = "http://schemas.microsoft.com/office/word/2018/wordml"
    NamespaceWord2020 = "http://schemas.microsoft.com/office/word/2020/wordml/sdtdatahash"
    NamespaceWord2023 = "http://schemas.microsoft.com/office/word/2023/wordml/word16du"
    NamespaceWord2024 = "http://schemas.microsoft.com/office/word/2024/wordml/sdtformatlock"

    // Drawing extension namespaces
    NamespaceDrawing2010 = "http://schemas.microsoft.com/office/drawing/2010/main"
    NamespaceDrawing2013 = "http://schemas.microsoft.com/office/drawing/2013/main/command"
    NamespaceDrawing2017 = "http://schemas.microsoft.com/office/drawing/2017/decorative"
    NamespaceDrawing2018 = "http://schemas.microsoft.com/office/drawing/2018/animation/model3d"

    // Excel extension namespaces
    NamespaceExcel2010 = "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"
    NamespaceExcel2015 = "http://schemas.microsoft.com/office/spreadsheetml/2015/02/main"
    NamespaceExcel2017RichData = "http://schemas.microsoft.com/office/spreadsheetml/2017/richdata"
    NamespaceExcel2024PivotRefresh = "http://schemas.microsoft.com/office/spreadsheetml/2024/pivotAutoRefresh"

    // PowerPoint extension namespaces
    NamespacePowerPoint2010 = "http://schemas.microsoft.com/office/powerpoint/2010/main"
    NamespacePowerPoint2012 = "http://schemas.microsoft.com/office/powerpoint/2012/roamingSettings"
    NamespacePowerPoint2015 = "http://schemas.microsoft.com/office/powerpoint/2015/main"
)

var NamespacePrefixMap = map[string]string{
    NamespaceWordprocessingML: "w",
    NamespaceWord2010:         "w14",
    NamespaceWord2012:         "w15",
    NamespaceWord2015:         "w16",
    // ... etc
}
```

## AlternateContent Handling

### What is AlternateContent?

AlternateContent is Office's version compatibility mechanism. It provides a Choice (modern version) and Fallback (older version) for the same content:

```xml
<mc:AlternateContent xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006">
  <mc:Choice Requires="w14">
    <w14:contentControl><!-- Office 2010+ feature --></w14:contentControl>
  </mc:Choice>
  <mc:Fallback>
    <w:sdt><!-- Office 2007 compatible fallback --></w:sdt>
  </mc:Fallback>
</mc:AlternateContent>
```

### AlternateContent Element Implementation

```go
// openxml/alternate_content.go (new file)

// AlternateContent provides version compatibility via Choice/Fallback
type AlternateContent struct {
    *CompositeElement
}

func NewAlternateContent() *AlternateContent {
    return &AlternateContent{
        CompositeElement: NewCompositeElement(
            "AlternateContent",
            NamespaceMarkupCompatibility,
        ),
    }
}

// Choice returns the first Choice child (modern version)
func (a *AlternateContent) Choice() *Choice {
    return GetElement[*Choice](a)
}

// Fallback returns the Fallback child (compatible version)
func (a *AlternateContent) Fallback() *Fallback {
    return GetElement[*Fallback](a)
}

// SelectContent returns appropriate content based on target version
func (a *AlternateContent) SelectContent(targetVersion FileFormatVersion) Element {
    choice := a.Choice()
    if choice != nil && choice.IsCompatibleWith(targetVersion) {
        return choice.FirstChild()
    }

    fallback := a.Fallback()
    if fallback != nil {
        return fallback.FirstChild()
    }

    return nil
}

// Choice represents the modern version option
type Choice struct {
    *CompositeElement
    Requires string // Namespace prefix required (e.g., "w14")
}

func (c *Choice) IsCompatibleWith(version FileFormatVersion) bool {
    // Check if version supports required namespace
    requiredVersion := prefixToVersion(c.Requires)
    return version >= requiredVersion
}

// Fallback represents the backward-compatible option
type Fallback struct {
    *CompositeElement
}
```

### AlternateContent Processing

When reading documents:
```go
// openxml/xml_reader.go (modified)

func (r *Reader) ReadElement() (Element, error) {
    // ... existing code ...

    // Special handling for AlternateContent
    if r.current.LocalName == "AlternateContent" &&
       r.current.NamespaceURI == NamespaceMarkupCompatibility {
        return r.readAlternateContent()
    }

    // ... existing code ...
}

func (r *Reader) readAlternateContent() (*AlternateContent, error) {
    ac := NewAlternateContent()

    // Read Choice and Fallback children
    for r.Read() {
        if r.current.Type == xml.EndElement {
            break
        }

        child, err := r.ReadElement()
        if err != nil {
            return nil, err
        }
        ac.AppendChild(child)
    }

    return ac, nil
}
```

When writing documents with version targeting:
```go
// openxml/xml_writer.go (modified)

func (w *Writer) WriteElement(elem Element, targetVersion FileFormatVersion) error {
    // Handle AlternateContent specially
    if ac, ok := elem.(*AlternateContent); ok {
        // Select appropriate content based on target version
        content := ac.SelectContent(targetVersion)
        if content != nil {
            return w.WriteElement(content, targetVersion)
        }
        return nil
    }

    // ... existing code ...
}
```

## Unknown Element Preservation

### UnknownElement Type

```go
// openxml/unknown_element.go (new file)

// UnknownElement preserves XML elements not recognized by goffice
type UnknownElement struct {
    *BaseElement
    localName    string
    namespaceURI string
    rawXML       []byte
}

func NewUnknownElement(localName, namespaceURI string, rawXML []byte) *UnknownElement {
    return &UnknownElement{
        BaseElement:  NewBaseElement(),
        localName:    localName,
        namespaceURI: namespaceURI,
        rawXML:       rawXML,
    }
}

func (u *UnknownElement) LocalName() string {
    return u.localName
}

func (u *UnknownElement) NamespaceURI() string {
    return u.namespaceURI
}

func (u *UnknownElement) WriteTo(w io.Writer) error {
    // Write preserved raw XML unchanged
    _, err := w.Write(u.rawXML)
    return err
}
```

### Unknown Element Detection

```go
// openxml/xml_reader.go (modified)

func (r *Reader) ReadElement() (Element, error) {
    qname := QualifiedName{
        Namespace: r.current.NamespaceURI,
        LocalName: r.current.LocalName,
    }

    // Try to find registered element type
    factory := r.registry.Get(qname)
    if factory == nil {
        // Unknown element - preserve as UnknownElement
        return r.readUnknownElement()
    }

    elem := factory()
    // ... read element normally ...
    return elem, nil
}

func (r *Reader) readUnknownElement() (*UnknownElement, error) {
    // Capture entire element subtree as raw XML
    startPos := r.decoder.InputOffset()
    depth := 1

    for depth > 0 && r.Read() {
        if r.current.Type == xml.StartElement {
            depth++
        } else if r.current.Type == xml.EndElement {
            depth--
        }
    }

    endPos := r.decoder.InputOffset()
    rawXML := r.input[startPos:endPos]

    return NewUnknownElement(
        r.current.LocalName,
        r.current.NamespaceURI,
        rawXML,
    ), nil
}
```

## Version Metadata System

### FileFormatVersion Enum

```go
// openxml/version.go (new file)

type FileFormatVersion int

const (
    Office2007 FileFormatVersion = iota
    Office2010
    Office2013
    Office2016
    Office2019
    Office2021
    Office2022
    Office2023
    Office2024
    Office2025
    Microsoft365
)

func (v FileFormatVersion) String() string {
    names := map[FileFormatVersion]string{
        Office2007:   "Office 2007",
        Office2010:   "Office 2010",
        Office2013:   "Office 2013",
        Office2016:   "Office 2016",
        Office2019:   "Office 2019",
        Office2021:   "Office 2021",
        Office2022:   "Office 2022",
        Office2023:   "Office 2023",
        Office2024:   "Office 2024",
        Office2025:   "Office 2025",
        Microsoft365: "Microsoft 365",
    }
    return names[v]
}
```

### Element Metadata with Version

```go
// openxml/metadata.go (modified)

type IElementMetadata interface {
    QName() QualifiedName
    Validators() []Validator
    ChildInfo() ParticleConstraint
    AvailableInVersion() FileFormatVersion  // NEW
}

// Generated element metadata includes version
type elementMetadata struct {
    qname              QualifiedName
    validators         []Validator
    childInfo          ParticleConstraint
    availableInVersion FileFormatVersion
}

func (m *elementMetadata) AvailableInVersion() FileFormatVersion {
    return m.availableInVersion
}
```

### Generated Element Example

```go
// wordprocessing/elements/elements.go (generated)

// ContentControl is a content control (Office 2010+)
type ContentControl struct {
    *CompositeElement
    // ... fields ...
}

func NewContentControl() *ContentControl {
    return &ContentControl{
        CompositeElement: NewCompositeElement(
            "contentControl",
            NamespaceWord2010,
        ),
    }
}

func (c *ContentControl) Metadata() IElementMetadata {
    return &elementMetadata{
        qname: QualifiedName{
            Namespace: NamespaceWord2010,
            LocalName: "contentControl",
        },
        availableInVersion: Office2010,  // Extracted from schema
        validators:         contentControlValidators,
        childInfo:          contentControlChildren,
    }
}
```

## Version-Aware Validation

### Validation with Version Target

```go
// openxml/validation/validator.go (modified)

type ValidationContext struct {
    TargetVersion FileFormatVersion  // NEW
    Errors        []ValidationError
    Settings      ValidationSettings
}

func Validate(doc Document, targetVersion FileFormatVersion) []ValidationError {
    ctx := &ValidationContext{
        TargetVersion: targetVersion,
        Settings:      DefaultSettings,
    }

    ValidateElement(doc.RootElement(), ctx)
    return ctx.Errors
}

func ValidateElement(elem Element, ctx *ValidationContext) {
    // Check if element is available in target version
    metadata := elem.Metadata()
    if metadata.AvailableInVersion() > ctx.TargetVersion {
        ctx.Errors = append(ctx.Errors, ValidationError{
            Element:     elem,
            Code:        "ElementNotAvailableInVersion",
            Description: fmt.Sprintf(
                "%s not available in %s (introduced in %s)",
                elem.LocalName(),
                ctx.TargetVersion,
                metadata.AvailableInVersion(),
            ),
        })
    }

    // Continue with normal validation
    // ... existing validation logic ...
}
```

### Document Version Detection

```go
// openxml/version_detector.go (new file)

// DetectMinimumVersion scans document to find newest element used
func DetectMinimumVersion(doc Document) FileFormatVersion {
    var maxVersion FileFormatVersion = Office2007

    walkElements(doc.RootElement(), func(elem Element) {
        metadata := elem.Metadata()
        if metadata.AvailableInVersion() > maxVersion {
            maxVersion = metadata.AvailableInVersion()
        }
    })

    return maxVersion
}

func walkElements(elem Element, fn func(Element)) {
    fn(elem)

    if composite, ok := elem.(CompositeElement); ok {
        for _, child := range composite.Children() {
            walkElements(child, fn)
        }
    }
}
```

## Testing Strategy

### Generator Testing

```go
// cmd/gen-go-wordprocessing/main_test.go

func TestExtensionSchemaLoading(t *testing.T) {
    schemas := loadSchemas("../../Open-XML-SDK/data/schemas")

    // Verify we loaded extension schemas
    extensionCount := 0
    for _, s := range schemas {
        if s.Version > Office2007 {
            extensionCount++
        }
    }

    assert.Greater(t, extensionCount, 50, "Should load 50+ extension schemas")
}

func TestNamespaceMapping(t *testing.T) {
    tests := []struct {
        namespace string
        expected  string
    }{
        {"http://schemas.microsoft.com/office/word/2010/wordml", "w14"},
        {"http://schemas.microsoft.com/office/word/2012/wordml", "w15"},
        {"http://schemas.microsoft.com/office/drawing/2010/main", "a14"},
    }

    for _, tt := range tests {
        prefix := derivePrefix(tt.namespace, versionFromNamespace(tt.namespace))
        assert.Equal(t, tt.expected, prefix)
    }
}
```

### AlternateContent Testing

```go
// openxml/alternate_content_test.go

func TestAlternateContentSelection(t *testing.T) {
    ac := NewAlternateContent()

    choice := NewChoice()
    choice.Requires = "w14"
    choice.AppendChild(NewContentControl())
    ac.AppendChild(choice)

    fallback := NewFallback()
    fallback.AppendChild(NewStructuredDocumentTag())
    ac.AppendChild(fallback)

    // Office 2010+ should get choice
    elem := ac.SelectContent(Office2010)
    assert.IsType(t, &ContentControl{}, elem)

    // Office 2007 should get fallback
    elem = ac.SelectContent(Office2007)
    assert.IsType(t, &StructuredDocumentTag{}, elem)
}
```

### Unknown Element Testing

```go
// openxml/unknown_element_test.go

func TestUnknownElementPreservation(t *testing.T) {
    xml := `<unknown:element xmlns:unknown="http://example.com/unknown">
        <unknown:child>data</unknown:child>
    </unknown:element>`

    // Read unknown element
    elem, err := ParseElement(strings.NewReader(xml))
    require.NoError(t, err)
    assert.IsType(t, &UnknownElement{}, elem)

    // Write back - should be identical
    var buf bytes.Buffer
    err = elem.WriteTo(&buf)
    require.NoError(t, err)
    assert.Equal(t, xml, buf.String())
}
```

### Version Validation Testing

```go
// openxml/validation/version_test.go

func TestVersionValidation(t *testing.T) {
    doc := NewWordprocessingDocument()
    body := doc.MainDocumentPart().Document().Body()

    // Add Office 2010 element
    body.AppendChild(NewContentControl())

    // Validate against Office 2007 - should error
    errors := Validate(doc, Office2007)
    assert.Len(t, errors, 1)
    assert.Contains(t, errors[0].Description, "not available in Office 2007")

    // Validate against Office 2010 - should pass
    errors = Validate(doc, Office2010)
    assert.Len(t, errors, 0)
}
```

### Roundtrip Testing

```go
// wordprocessing/roundtrip_extension_test.go

func TestExtensionElementRoundtrip(t *testing.T) {
    // Create document with Office 2010 elements
    doc := NewWordprocessingDocument()
    body := doc.MainDocumentPart().Document().Body()

    cc := NewContentControl()
    cc.SdtProperties().SetAlias("MyControl")
    body.AppendChild(cc)

    // Save
    var buf bytes.Buffer
    err := doc.SaveAs(&buf)
    require.NoError(t, err)

    // Reopen
    doc2, err := OpenWordprocessingDocument(bytes.NewReader(buf.Bytes()))
    require.NoError(t, err)

    // Verify extension element preserved
    cc2 := GetElement[*ContentControl](doc2.MainDocumentPart().Document().Body())
    require.NotNil(t, cc2)
    assert.Equal(t, "MyControl", cc2.SdtProperties().Alias())
}
```

## Migration Path

### Phase 1: Generator Updates
1. Update schema loading logic
2. Add version extraction
3. Add namespace mapping
4. Test generators (don't generate yet)

### Phase 2: Framework Infrastructure
1. Implement AlternateContent elements
2. Implement UnknownElement
3. Add version metadata to Element interface
4. Update XML reader for unknown elements

### Phase 3: Code Generation
1. Run generators with extension schemas
2. Fix any generation errors
3. Verify compilation
4. Run golangci-lint

### Phase 4: Validation Integration
1. Add version parameter to validation API
2. Implement version checking in validators
3. Add version detection API
4. Test version validation

### Phase 5: Testing & Validation
1. Create test documents with extensions
2. Roundtrip tests
3. Version validation tests
4. Unknown element preservation tests

## Performance Considerations

### Code Generation Time
- Loading 100+ schemas instead of 3-5 will increase generation time
- Expected: 5-10 seconds → 30-60 seconds (acceptable for infrequent operation)
- Mitigation: Parallel schema loading, caching

### Generated Code Size
- Current: ~50K lines per package
- Expected: ~70-80K lines per package
- Impact: Slightly longer compile times (acceptable)
- Mitigation: Can split files later if needed

### Runtime Performance
- No impact - generated code identical in structure
- Version checking is fast (integer comparison)
- Unknown element preservation adds minimal overhead (only for unknown elements)

## Success Metrics

- [ ] All 3 generators load 100+ schemas (not just main.json)
- [ ] 200-300 new element types generated
- [ ] All extension namespaces defined in constants
- [ ] AlternateContent/Choice/Fallback implemented and tested
- [ ] UnknownElement preserves forward compatibility
- [ ] Version-aware validation passes tests
- [ ] Roundtrip tests pass with extension elements
- [ ] Zero regression in existing tests
- [ ] golangci-lint passes
- [ ] Documentation complete

## Future Enhancements

1. **Automatic Fallback Generation**: Given Office 2019 feature, auto-generate Office 2007 equivalent
2. **Version Downgrade API**: Convert document from Office 2019 to Office 2010 format
3. **Extension Discovery API**: List all extensions used in a document
4. **Namespace-Aware Pretty Printing**: Assign minimal necessary namespace prefixes when writing
5. **Schema Evolution Tracking**: Track when elements/attributes changed across Office versions
