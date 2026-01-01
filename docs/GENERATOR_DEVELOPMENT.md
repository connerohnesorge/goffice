# Code Generator Development Guide

This guide describes how the goffice code generators work and how they support extension schemas.

## Overview

goffice includes code generators that create Go structs and methods from Office Open XML schemas. These generators:

1. Load schema files from the Open-XML-SDK
2. Extract element and attribute definitions
3. Generate type-safe Go code
4. Include version metadata for each element

**Location:** `cmd/gen-go-wordprocessing/`, `cmd/gen-go-spreadsheet/`, `cmd/gen-go-presentation/`

## Generator Architecture

### Three-Phase Generation

1. **Schema Loading** - Load JSON schema files and build in-memory type definitions
2. **Type Collection** - Extract types from schemas, handle conflicts, add metadata
3. **Code Generation** - Template-based generation of Go code

### Schema Loading with Extensions

Starting with version X.X, generators load all relevant schemas:

```
Open-XML-SDK/data/schemas/
├── wordprocessingml
│   ├── main.json                                 (Office 2007)
│   ├── schemas_microsoft_com_office_word_2010_wordml.json
│   ├── schemas_microsoft_com_office_word_2012_wordml.json
│   ├── schemas_microsoft_com_office_word_2015_wordml_symex.json
│   ├── schemas_microsoft_com_office_word_2016_wordml_cid.json
│   └── ... (through 2024)
└── drawingml/
    ├── schemas_microsoft_com_office_drawing_2010_main.json
    ├── schemas_microsoft_com_office_drawing_2012_main.json
    └── ... (through 2024)
```

### Schema File Metadata Extraction

Each schema filename encodes metadata:

```
schemas_microsoft_com_office_word_2010_wordml.json
         ↑              ↑    ↑    ↑ ↑
         +--+--+--+--+--+    |    | +- Namespace/suffix
             Company         App  Year

Parsed as:
- App: "word"
- Year: 2010
- Namespace: http://schemas.microsoft.com/office/word/2010/wordml
- VersionEnum: Office2010
- NamespacePrefix: "w14" (Office 2010 = version 14)
```

## Running Generators

### Regenerate All Schemas

```bash
# Regenerate Word elements with extensions
go run ./cmd/gen-go-wordprocessing

# Regenerate Excel elements
go run ./cmd/gen-go-spreadsheet

# Regenerate PowerPoint elements
go run ./cmd/gen-go-presentation
```

### Output Structure

```
wordprocessing/elements/
├── elements.go         (Main generated file with all types)
├── namespaces.go       (Generated namespace constants)
└── metadata.go         (Generated element metadata)
```

## Generator Code Organization

### cmd/gen-go-wordprocessing/main.go

Main entry point that orchestrates generation:

```go
func main() {
    // 1. Load schemas (main + extensions)
    schemas := loadSchemas("Open-XML-SDK/data/schemas")
    
    // 2. Collect types from all schemas
    types := collectTypesWithVersion(schemas)
    
    // 3. Generate Go code
    generateCode(types, "wordprocessing/elements/elements.go")
    
    // 4. Generate namespace constants
    generateNamespaces(schemas, "wordprocessing/elements/namespaces.go")
}
```

### SchemaFile Structure

```go
type SchemaFile struct {
    Path            string                  // File path
    Namespace       string                  // XML namespace URL
    NamespacePrefix string                  // Prefix (w14, x15, a14, etc)
    Version         FileFormatVersion       // Office version
    TargetApp       string                  // "word", "excel", "powerpoint", "drawing"
}
```

### Version Extraction Logic

```go
// Extract year from filename
// "schemas_microsoft_com_office_word_2010_wordml.json" → 2010
// "schemas_microsoft_com_office_powerpoint_2022_03_main.json" → 2022

func parseSchemaFile(filename string) SchemaFile {
    // Extract version year using regex
    year := extractYearFromFilename(filename)
    
    // Map to FileFormatVersion enum
    version := yearToVersion(year)
    
    // Determine namespace prefix
    prefix := derivePrefix(namespace, version)
    
    return SchemaFile{...}
}

func yearToVersion(year int) FileFormatVersion {
    switch year {
    case 2010: return Office2010
    case 2013: return Office2013
    case 2016: return Office2016
    case 2019: return Office2019
    case 2021: return Office2021
    case 2022: return Office2022
    case 2023: return Office2023
    case 2024: return Office2024
    case 2025: return Office2025
    default:
        // Round down for in-between years
        if year >= 2025 { return Office2025 }
        if year >= 2024 { return Office2024 }
        // ... more mappings
        return Office2007
    }
}
```

### Namespace Prefix Derivation

```go
func derivePrefix(namespace string, version FileFormatVersion) string {
    // Extract base prefix from namespace type
    base := "w"  // default for word
    if strings.Contains(namespace, "spreadsheetml") {
        base = "x"
    } else if strings.Contains(namespace, "powerpoint") {
        base = "p"
    } else if strings.Contains(namespace, "drawing") {
        base = "a"
    }
    
    // Add version suffix
    suffix := map[FileFormatVersion]string{
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
    }[version]
    
    return base + suffix  // "w14", "x15", "p14", "a14", etc
}
```

## Type Collection with Version Metadata

### Handling Version-Specific Types

When a type appears in multiple version schemas:

1. First occurrence: Create type definition
2. Subsequent occurrences: Update metadata with version info
3. Conflicts: Track all versions where type exists

```go
func collectTypesWithVersion(schemas []SchemaFile) map[string]*TypeDef {
    types := make(map[string]*TypeDef)
    
    for _, schema := range schemas {
        schemaTypes := loadSchema(schema.Path)
        
        for name, typeDef := range schemaTypes {
            if existing, ok := types[name]; ok {
                // Type exists - track all versions
                existing.AvailableVersions = append(
                    existing.AvailableVersions,
                    schema.Version,
                )
                // Use the earliest version as primary
                if schema.Version < existing.AvailableIn {
                    existing.AvailableIn = schema.Version
                }
            } else {
                // New type
                typeDef.AvailableIn = schema.Version
                typeDef.AvailableVersions = []FileFormatVersion{schema.Version}
                types[name] = typeDef
            }
        }
    }
    
    return types
}
```

## Generated Code Structure

### Element Type Definition

```go
// Generated in wordprocessing/elements/elements.go

// ContentControl is a Word 2010 extension element (w14)
type ContentControl struct {
    XMLName xml.Name
    
    // Attributes
    Alias *string  `xml:"alias,attr"`
    Tag   *string  `xml:"tag,attr"`
    
    // Child elements
    SdtProperties *SdtProperties
    Children      []interface{}
    
    // Version metadata (generated)
    _metadata *ElementMetadata
}

func (cc *ContentControl) Metadata() *ElementMetadata {
    if cc._metadata == nil {
        cc._metadata = &ElementMetadata{
            Name:                "ContentControl",
            AvailableInVersion:  Office2010,
            Namespace:           W14Namespace,
            NamespacePrefix:     "w14",
        }
    }
    return cc._metadata
}
```

### Namespace Constants

```go
// Generated in wordprocessing/elements/namespaces.go

const (
    // Main namespace (Office 2007)
    WordprocessingMLNamespace = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
    
    // Word 2010 extensions
    W14Namespace = "http://schemas.microsoft.com/office/word/2010/wordml"
    
    // Word 2012 extensions
    W15Namespace = "http://schemas.microsoft.com/office/word/2012/wordml"
    
    // Word 2016 extensions
    W16Namespace = "http://schemas.microsoft.com/office/word/2015/wordml_cid"
    
    // ... more namespaces through Office 2024
)
```

## Common Generation Tasks

### Add Support for New Office Version

1. Update `Open-XML-SDK` submodule with new schemas
2. Run generators - they automatically detect new schemas
3. Verify new types are generated
4. Update `FileFormatVersion` enum if needed

```bash
cd Open-XML-SDK
git pull origin master
cd ..
go run ./cmd/gen-go-wordprocessing
```

### Handle Type Conflicts

If two schemas define the same type name:

1. Generator detects conflict
2. Creates disambiguated names: `TypeName_Office2010`, `TypeName_Office2016`
3. Logs warning for manual resolution

### Extract Enum Values

Generators extract enum restrictions from schemas:

```json
// In schema
{
  "type": "string",
  "enum": ["left", "center", "right", "justify"]
}
```

Generated as:

```go
type JustificationValue string

const (
    JustificationLeft     JustificationValue = "left"
    JustificationCenter   JustificationValue = "center"
    JustificationRight    JustificationValue = "right"
    JustificationJustify  JustificationValue = "justify"
)
```

## Debugging Generators

### Enable Verbose Output

Modify `main.go` to add debug logging:

```go
import "log"

func loadSchemas(dir string) []SchemaFile {
    var schemas []SchemaFile
    files, _ := os.ReadDir(dir)
    
    for _, f := range files {
        sf := parseSchemaFile(f.Name())
        log.Printf("Loaded schema: %s → %s (version %v)\n",
            f.Name(), sf.NamespacePrefix, sf.Version)
        schemas = append(schemas, sf)
    }
    
    return schemas
}
```

### Verify Type Counts

Count generated types to ensure extension schemas loaded:

```bash
# Before extension support (Office 2007 only)
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Expected: ~619

# After extension support (2007-2024)
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Expected: ~700-750 (depending on overlap)
```

### Check Namespace Generation

```bash
# Verify namespace constants exist
grep "^const.*Namespace" wordprocessing/elements/namespaces.go | head -20

# Expected output includes:
# const W14Namespace = "..."
# const W15Namespace = "..."
# ... more through W24Namespace
```

## Performance Optimization

### Schema Loading Optimization

For faster generation with many schemas:

```go
// Parallel schema loading
var wg sync.WaitGroup
schemasChannel := make(chan SchemaFile)

for _, f := range files {
    wg.Add(1)
    go func(filename string) {
        defer wg.Done()
        sf := parseSchemaFile(filename)
        schemasChannel <- sf
    }(f.Name())
}

go func() {
    wg.Wait()
    close(schemasChannel)
}()
```

### Type Deduplication

Cache schema loads to avoid redundant parsing:

```go
var schemaCache = make(map[string]interface{})

func loadSchema(path string) interface{} {
    if cached, ok := schemaCache[path]; ok {
        return cached
    }
    
    // Load from disk
    data := loadFromDisk(path)
    schemaCache[path] = data
    return data
}
```

## Testing Generators

### Unit Tests for Version Extraction

```go
func TestVersionExtraction(t *testing.T) {
    tests := []struct {
        filename string
        expected FileFormatVersion
    }{
        {"schemas_microsoft_com_office_word_2010_wordml.json", Office2010},
        {"schemas_microsoft_com_office_word_2012_wordml.json", Office2013},
        {"schemas_microsoft_com_office_word_2024_wordml_sdtformatlock.json", Office2024},
    }
    
    for _, tt := range tests {
        sf := parseSchemaFile(tt.filename)
        if sf.Version != tt.expected {
            t.Errorf("parseSchemaFile(%s) = %v, want %v",
                tt.filename, sf.Version, tt.expected)
        }
    }
}
```

### Integration Tests

```go
func TestGeneratorProducesValidCode(t *testing.T) {
    // Run generator
    output := runGenerator()
    
    // Verify it compiles
    cmd := exec.Command("go", "build", "-o", "/dev/null", ".")
    cmd.Dir = "wordprocessing/elements"
    if err := cmd.Run(); err != nil {
        t.Fatalf("Generated code does not compile: %v", err)
    }
    
    // Verify contains extension types
    if !strings.Contains(output, "ContentControl") {
        t.Error("Generated code missing ContentControl type")
    }
}
```

## Troubleshooting

### Generator Hangs

1. Check for infinite loops in schema loading
2. Add timeout: `ctx, cancel := context.WithTimeout()`
3. Verify schema files are valid JSON

### Missing Types

1. Verify schema files exist in `Open-XML-SDK/data/schemas/`
2. Check filename patterns match expected format
3. Run `ls Open-XML-SDK/data/schemas/ | grep word` to verify

### Compilation Errors

1. Check for reserved word usage in generated names
2. Verify enum value naming doesn't clash
3. Run `golangci-lint` on generated code

## References

- [Open-XML-SDK Schemas](https://github.com/OfficeDev/Open-XML-SDK/tree/main/src/DocumentFormat.OpenXml/schemas)
- [ECMA-376 Specification](http://www.ecma-international.org/publications/standards/Ecma-376.htm)
- [ISO/IEC 29500 Standard](https://www.iso.org/standard/71691.html)
