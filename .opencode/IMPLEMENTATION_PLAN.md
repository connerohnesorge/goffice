# Add Extension Schema Support - Complete Implementation Plan

## Overview
This document provides comprehensive guidance for completing the `add-extension-schema-support` change proposal.

**Status**: Phase 1 complete, Phase 2 partial, Phases 3-6 need implementation

**Change**: `spectr/changes/add-extension-schema-support/`

**Proposal**: `spectr/changes/add-extension-schema-support/proposal.md`

**Design**: `spectr/changes/add-extension-schema-support/design.md`

**Tasks**: `spectr/changes/add-extension-schema-support/tasks.md`

## Phase 0: Critical Blocker - Generator Naming Conflicts

### Problem
Generators produce naming conflicts:
- `Boolean` enum (from main schema)
- `BooleanFalse` struct (from drawing 2014 extension)

### Solution
Implement namespace-aware prefixing in struct generation.

### Implementation

#### File: cmd/gen-go-wordprocessing/gen_core.go
After `collectTypesWithVersion()`, add:
```go
// Build map of enum names to detect conflicts
enumNameMap := make(map[string]bool)

// Scan main schema for enums
mainSchemaPath := "Open-XML-SDK/data/schemas/schemas_openxmlformats_org_wordprocessingml_2006_main.json"
data, _ := os.ReadFile(mainSchemaPath)
var mainSchema SchemaFile
json.Unmarshal(data, &mainSchema)
for _, enum := range mainSchema.Enums {
    enumNameMap[enum.Name] = true
    enumNameMap[enum.ClassName] = true
}
```

#### File: cmd/gen-go-wordprocessing/gen_struct.go
Modify `generateStruct()`:
```go
func generateStruct(f *os.File, t *SchemaType) {
    className := t.ClassName
    
    // Check for conflicts with enums
    if enumNameMap[className] {
        prefix := deriveStructPrefix(t.TargetNamespace)
        className = prefix + t.ClassName
    }
    
    if existingTypes[className] ||
        generatedTypes[className] {
        return
    }
    generatedTypes[className] = true
    
    // Use className instead of t.ClassName in rest of function
    // ... rest of generateStruct
}
```

#### File: cmd/gen-go-wordprocessing/utils.go
Add function:
```go
func deriveStructPrefix(namespace string) string {
    // Returns namespace-aware prefix (A14, W14, X14, P14, etc.)
    // Based on namespace and version
    
    if strings.Contains(namespace, "/drawing/") {
        if strings.Contains(namespace, "2010") { return "A14" }
        if strings.Contains(namespace, "2012") { return "A15" }
        if strings.Contains(namespace, "2014") { return "A16" }
        return "ADraw"
    }
    if strings.Contains(namespace, "/word/") {
        if strings.Contains(namespace, "2010") { return "W14" }
        if strings.Contains(namespace, "2012") { return "W15" }
        return "WExt"
    }
    if strings.Contains(namespace, "/spreadsheetml/") {
        if strings.Contains(namespace, "2010") { return "X14" }
        return "XExt"
    }
    if strings.Contains(namespace, "/powerpoint/") {
        if strings.Contains(namespace, "2010") { return "P14" }
        return "PExt"
    }
    return "Ext"
}
```

#### Apply to Other Generators
Copy the same fix to:
- `cmd/gen-go-spreadsheet/gen_struct.go`
- `cmd/gen-go-spreadsheet/utils.go`
- `cmd/gen-go-presentation/gen_struct.go`
- `cmd/gen-go-presentation/utils.go`

### Verify
```bash
go build ./wordprocessing ./spreadsheet ./presentation
# Should compile without "redeclared in this block" errors
```

---

## Phase 2: Complete Framework Infrastructure

### 2.1 Update Element Interface Metadata

**File**: `openxml/element.go`

Add optional method to Element interface:
```go
// AvailableInVersion returns the minimum Office version required for this element.
// Returns nil for Office 2007 elements (no version metadata).
AvailableInVersion() FileFormatVersion
```

Alternatively, if avoiding changing the interface, create a new interface:
```go
type VersionedElement interface {
    Element
    AvailableInVersion() FileFormatVersion
}
```

**Implementation**: Generated structs will implement this method (generator will be updated in later phase).

### 2.2 XML Reader Unknown Element Handling

**File**: `openxml/openxml.go` or wherever XML reading happens

Locate the element factory/creation logic and add fallback:
```go
// When creating element from XML, if type not found:
if factory == nil {
    // Fallback to UnknownElement
    return NewUnknownElement(qname, attributes, rawXML)
}
```

Verify `UnknownElement` exists in `openxml/unknown_element.go` (already implemented).

### 2.3 Namespace Map Integration

**File**: `openxml/namespaces.go`

Ensure XML reader uses `NamespacePrefixMap`:
```go
// During XML parsing initialization:
for prefix, uri := range NamespacePrefixMap {
    decoder.RegisterNamespace(prefix, uri)
}
```

Verify `NamespacePrefixMap` is populated (already exists, check during testing).

### Tests
```bash
go test ./openxml -v
# Verify:
# - Unknown elements don't cause errors
# - Round-trip with unknown elements preserves content
# - Namespace prefixes resolve correctly
```

---

## Phase 3: Code Generation

### After fixing generator naming conflicts:

```bash
cd /home/connerohnesorge/Documents/001Repos/goffice

# 3.1 Regenerate WordProcessing
go run ./cmd/gen-go-wordprocessing

# 3.2 Regenerate Spreadsheet
go run ./cmd/gen-go-spreadsheet

# 3.3 Regenerate Presentation
go run ./cmd/gen-go-presentation

# Verify compilation
go build ./wordprocessing ./spreadsheet ./presentation

# Count generated types
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Expected: 1900+ (was ~619)

# Fix linting issues
golangci-lint run --fix ./...

# Run tests
go test ./... -v
```

### Success Criteria
- All three generators complete without errors
- 200+ new types per app
- Code compiles
- golangci-lint passes
- No test regressions

---

## Phase 4: Validation Integration

### 4.1 Version-Aware Validation

**File**: `openxml/validation/context.go`

Add field:
```go
type ValidationContext struct {
    // ... existing fields
    TargetVersion FileFormatVersion
}
```

**File**: `openxml/validation/validator.go`

Add version checking:
```go
// Before existing validation logic:
if versioned, ok := elem.(interface { AvailableInVersion() FileFormatVersion }); ok {
    elemVersion := versioned.AvailableInVersion()
    if elemVersion > ctx.TargetVersion {
        return NewValidationError(
            elem,
            fmt.Sprintf(
                "Element %s not available in %s (introduced in %s)",
                elem.LocalName(),
                ctx.TargetVersion.String(),
                elemVersion.String(),
            ),
        )
    }
}
```

### 4.2 AlternateContent Validation

**File**: Create `openxml/validation/alternate_content_validator.go`

```go
// Validate that AlternateContent has exactly one Choice and Fallback
// Validate that Choice.Requires references valid namespace
func ValidateAlternateContent(elem *AlternateContent, ctx *ValidationContext) []ValidationError
```

### 4.3 Version Detection API

**File**: Create `openxml/version_detector.go`

```go
// DetectMinimumVersion scans document and returns minimum Office version required
func DetectMinimumVersion(doc interface{}) FileFormatVersion {
    // Walk all elements
    // Find element with highest version
    // Return that version
}
```

### 4.4 Tests

Create `openxml/validation/version_validation_test.go`:
- Test version mismatch (Office 2010 element in 2007 context)
- Test version match (Office 2010 element in 2010+ context)
- Test multiple violations
- Test detection API

---

## Phase 5: Integration Testing

### 5.1 Test Documents
Obtain or create:
- `testdata/word2010-features.docx` - With Word 2010 extensions
- `testdata/excel2013-features.xlsx` - With Excel 2013 extensions
- `testdata/powerpoint2016-features.pptx` - With PowerPoint 2016 extensions

### 5.2 Roundtrip Tests

Create `wordprocessing/roundtrip_extension_test.go`:
```go
func TestExtensionElementRoundtrip(t *testing.T) {
    // Open document with extensions
    // Read and verify elements
    // Save document
    // Re-open and verify unchanged
}
```

Similar files for spreadsheet and presentation.

### 5.3 Unknown Element Tests
- Open document with unknown namespace elements
- Verify UnknownElement instances created
- Verify XML preserved on round-trip

### 5.4 AlternateContent Integration Tests
- Create document with Choice/Fallback
- Test SelectContent() logic
- Verify correct version selection

### 5.5 Version Detection Integration Tests
```go
func TestVersionDetection(t *testing.T) {
    // Document with Office 2007 elements only → Office2007
    // Document with Office 2013 elements → Office2013
    // Mixed document → highest version
}
```

### 5.6 Compatibility Testing
- Save goffice-created extension documents
- Open in Microsoft Office (Word, Excel, PowerPoint)
- Verify features work
- Test with LibreOffice for basic compatibility

---

## Phase 6: Documentation

### 6.1 API Documentation

**openxml/version.go**:
```go
// FileFormatVersion represents Office file format versions.
// Each version includes features introduced in that Office release.
//
// Version Progression:
// - Office2007: Original ECMA-376 standard
// - Office2010: Drawing extensions, Word 2010 features
// - Office2013: Drawing 2012 features, advanced drawing
// - Office2016: Drawing 2014+ features, SVG support
// - Office2019: Modern features
// - Office2021-Office2025: Annual releases
// - Microsoft365: Cloud features
type FileFormatVersion int
```

**openxml/alternate_content.go**:
- Document AlternateContent pattern
- Explain Choice.Requires
- Show example: Choice for Office 2010+, Fallback for 2007

**openxml/unknown_element.go**:
- Document forward compatibility
- Explain how unknown extensions are preserved
- Show preservation guarantee

**openxml/validation/**:
- Update package docs
- Document TargetVersion parameter
- Show example: `Validate(doc, Office2013)`

### 6.2 User Guide

Create documentation explaining:
- What extension schemas are
- Supported Office versions (2010-2025)
- List of key features by version
- How to use extension elements
- Version detection API usage
- Validation with version checking

### 6.3 Developer Guide

Update developer documentation:
- How generator loads schemas
- Schema loading patterns
- Namespace prefix derivation
- Version metadata extraction
- Adding new schema patterns

### 6.4 Examples

Create example files:
- `examples/word-content-controls.go` - Word 2010 extensions
- `examples/excel-slicers.go` - Excel 2010 extensions
- `examples/version-detection.go` - DetectMinimumVersion API
- `examples/validation-version.go` - Version-aware validation

---

## Order of Implementation

1. **Phase 0**: Fix generator (blocks Phase 3)
2. **Phase 2**: Complete framework (enables Phase 3)
3. **Phase 3**: Generate code (enables Phase 4-5)
4. **Phase 4**: Validation (enables Phase 5)
5. **Phase 5**: Testing (validates everything works)
6. **Phase 6**: Documentation (final polish)

---

## Testing Commands

```bash
# After Phase 0 (generator fix)
go build ./wordprocessing ./spreadsheet ./presentation

# After Phase 2 (framework)
go test ./openxml -v

# After Phase 3 (generation)
go test ./... -v
golangci-lint run ./...

# After Phase 4 (validation)
go test ./openxml/validation -v

# After Phase 5 (integration)
go test ./wordprocessing -v
go test ./spreadsheet -v
go test ./presentation -v

# Final verification
go test ./...
golangci-lint run ./...
```

---

## Success Criteria

- [x] Phase 1: Generator infrastructure (DONE)
- [ ] Phase 0: Generator naming conflicts fixed
- [ ] Phase 2: Framework complete
- [ ] Phase 3: All generators produce 200+ new types, no errors
- [ ] Phase 4: Version-aware validation working
- [ ] Phase 5: All integration tests passing
- [ ] Phase 6: Documentation complete
- [ ] Final: All tests pass, golangci-lint passes, no regressions

---

## Key Design Points

1. **Non-breaking**: All changes are additive, existing code unchanged
2. **Framework-driven**: Extension elements generated, not hand-written
3. **Version-aware**: Elements include version metadata
4. **Forward-compatible**: Unknown elements preserved for future Office versions
5. **Minimal changes**: Only touches generator, validation, and framework

---

## References

- Proposal: `spectr/changes/add-extension-schema-support/proposal.md`
- Design: `spectr/changes/add-extension-schema-support/design.md`
- Tasks: `spectr/changes/add-extension-schema-support/tasks.md`
