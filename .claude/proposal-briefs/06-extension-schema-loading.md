# Proposal Brief: Extension Schema Loading (Office 2010-2024 Support)

## Priority: TIER 2 - High
**Impact**: 8/10  
**Effort**: 2-3 weeks  
**Status**: Generator design limitation

## Gap Analysis

### Current State
**What exists:**
- ✅ **Generators work** for main schemas (cmd/gen-go-wordprocessing/, cmd/gen-go-spreadsheet/, cmd/gen-go-presentation/)
- ✅ **Main schemas loaded**: Only `main.json` files currently processed

**What's missing:**
- ❌ **Extension schemas NOT loaded**: 100+ extension schemas exist but unused
- ❌ **Office 2010-2024 features**: Missing 200-300 element types
- ❌ **Modern features**: No support for newer Office capabilities

### Current Generator Behavior

**WordProcessing Generator** (cmd/gen-go-wordprocessing/main.go):
```go
for _, f := range files {
    if strings.Contains(f.Name(), "main.json") {
        collectTypes(filepath.Join(schemasDir, f.Name()))
    }
}
```

**Problem**: Only loads `main.json`, skips all extension schemas!

### Missing Schemas (Examples)

**WordProcessing Extensions**:
- `schemas_microsoft_com_office_word_2010_wordml.json`
- `schemas_microsoft_com_office_word_2012_wordml.json`
- `schemas_microsoft_com_office_word_2015_wordml_symex.json`
- `schemas_microsoft_com_office_word_2016_wordml_cid.json`
- `schemas_microsoft_com_office_word_2018_wordml.json`
- `schemas_microsoft_com_office_word_2020_wordml_sdtdatahash.json`
- `schemas_microsoft_com_office_word_2023_wordml_word16du.json`
- `schemas_microsoft_com_office_word_2024_wordml_sdtformatlock.json`
- Plus drawing/shape/canvas extensions (10+ more)

**Spreadsheet Extensions**:
- `schemas_microsoft_com_office_excel_2010_spreadsheetml.json`
- `schemas_microsoft_com_office_excel_2012_pivot.json`
- Through 2024 extensions

**Presentation Extensions**:
- `schemas_microsoft_com_office_powerpoint_2010_main.json`
- Through 2024 extensions

**Estimated Missing Elements**: 200-300 types across all extensions

## Implementation Scope

### Files to Modify

1. **cmd/gen-go-wordprocessing/main.go**
   - Change filter from `main.json` to include extension schemas
   - Add namespace handling for extension elements

2. **cmd/gen-go-spreadsheet/main.go**
   - Same changes as wordprocessing

3. **cmd/gen-go-presentation/main.go**
   - Same changes as presentation

4. **cmd/gen-go-drawingml/main.go** (if separate)
   - Include drawing extensions

### Core Design Decisions

**1. Which Schemas to Load?**

**Option A**: Load ALL schemas (main + all extensions)
- **Pro**: Complete coverage
- **Con**: May generate unused elements, larger codebase

**Option B**: Load specific extension years (e.g., 2010, 2013, 2016, 2019)
- **Pro**: Controlled scope
- **Con**: Arbitrary cutoff, may miss features

**Option C**: Load based on namespace patterns
- **Pro**: Systematic approach
- **Con**: Still loads everything

**Recommendation**: **Option A** - Load all schemas. Unused elements are no-ops, disk is cheap.

**2. Namespace Handling**

Extension schemas use different XML namespaces:
- Main: `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
- 2010: `http://schemas.microsoft.com/office/word/2010/wordml`
- 2012: `http://schemas.microsoft.com/office/word/2012/wordml`
- etc.

Generator must:
- Map namespace prefixes correctly
- Generate constants for each namespace
- Handle namespace in XML serialization

**3. Element Organization**

**Current**: All elements in one file (`elements/elements.go`)

**Proposed**: Keep same pattern, but optionally split by namespace:
- `elements/main.go` - Main 2006 namespace
- `elements/word2010.go` - Word 2010 extensions
- `elements/word2012.go` - Word 2012 extensions
- etc.

**Or**: Keep single file (simpler, current pattern)

**Recommendation**: **Single file** for now, split later if too large.

**4. Backward Compatibility**

Loading extensions doesn't break existing code:
- New elements are additive
- Existing elements unchanged
- Namespaces are distinct

**5. Testing**

Need test documents using modern features:
- Word 2010+ documents with extended features
- Excel 2013+ with slicers, timelines
- PowerPoint 2016+ with modern animations

## References

### OpenXML-SDK Schema Loading
Look at @OpenXML-SDK/data/schemas/ to see:
- All available schema files
- Naming patterns
- What extensions exist

### Generator Code
- `cmd/gen-go-wordprocessing/main.go` - Current implementation
- `cmd/gen-go-wordprocessing/loaders.go` - Schema loading logic
- `cmd/gen-go-wordprocessing/gen_struct.go` - Element generation

## Success Criteria

- [ ] All extension schemas are loaded during generation
- [ ] New element types are generated (200-300 additional types)
- [ ] Namespace constants are generated correctly
- [ ] Existing tests pass (no regression)
- [ ] Can open Office 2010+ documents with extended features
- [ ] Extended elements roundtrip correctly
- [ ] Code compiles after regeneration
- [ ] golangci-lint passes

## Spec Capabilities to Update

- **wordprocessing-elements**: Note extension element support
- **spreadsheet-elements**: Note extension element support
- **presentation-elements**: Note extension element support
- **framework**: Add namespace handling requirements

## Dependencies

- Existing generator framework (cmd/gen-go-*)
- OpenXML-SDK schema files (already present)

## Out of Scope

- Implementing APIs for new elements (that's separate proposals)
- Full documentation of all 200-300 new elements
- Migration guide for users

## Example Changes

### Before (main.json only):
```go
// cmd/gen-go-wordprocessing/main.go
for _, f := range files {
    if strings.Contains(f.Name(), "main.json") {
        collectTypes(filepath.Join(schemasDir, f.Name()))
    }
}
```

### After (all schemas):
```go
// Load main schema
if strings.Contains(f.Name(), "wordprocessingml") && strings.Contains(f.Name(), "main.json") {
    collectTypes(filepath.Join(schemasDir, f.Name()))
}

// Load extension schemas
if strings.Contains(f.Name(), "microsoft_com_office_word") {
    collectTypes(filepath.Join(schemasDir, f.Name()))
}
```

## Implementation Algorithm

```go
func loadSchemas(schemasDir string) error {
    files, _ := os.ReadDir(schemasDir)
    
    var schemas []string
    
    // 1. Find all relevant schema files
    for _, f := range files {
        name := f.Name()
        
        // Main WordProcessingML schema
        if strings.Contains(name, "wordprocessingml") && strings.Contains(name, "main.json") {
            schemas = append(schemas, filepath.Join(schemasDir, name))
        }
        
        // Extension schemas (Office 2010-2024)
        if strings.Contains(name, "microsoft_com_office_word") {
            schemas = append(schemas, filepath.Join(schemasDir, name))
        }
        
        // Drawing extensions (shapes, canvas, etc.)
        if strings.Contains(name, "microsoft_com_office_drawing") {
            schemas = append(schemas, filepath.Join(schemasDir, name))
        }
    }
    
    // 2. Load each schema
    for _, schemaPath := range schemas {
        if err := collectTypes(schemaPath); err != nil {
            return fmt.Errorf("failed to load %s: %w", schemaPath, err)
        }
    }
    
    return nil
}

func collectTypes(schemaPath string) error {
    // Existing logic:
    // 1. Parse JSON schema
    // 2. Extract type definitions
    // 3. Build Go struct definitions
    // 4. Register in type map
    
    // NEW: Also extract and register namespace
    schema := parseSchema(schemaPath)
    registerNamespace(schema.TargetNamespace)
    
    // Continue with existing type collection...
}
```

## Validation

**Before regeneration**:
```bash
# Count current element types
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Example output: 619 types
```

**After regeneration**:
```bash
# Should have 200-300 more
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Example output: 850 types
```

## Risks

**1. Code Size Increase**
- **Risk**: Generated files become very large (50K+ lines)
- **Mitigation**: Split into multiple files if needed, build times still acceptable

**2. Breaking Changes**
- **Risk**: Schema changes cause generation failures
- **Mitigation**: Extensive testing, can always regenerate from known-good schemas

**3. Namespace Conflicts**
- **Risk**: Extension elements conflict with main elements
- **Mitigation**: Namespaces are distinct, Go package handles this

## Estimated Effort

- **Week 1**: Update wordprocessing generator, regenerate, test
- **Week 2**: Update spreadsheet and presentation generators, regenerate, test
- **Week 3**: Integration testing, validation, documentation

**Total**: 2-3 weeks

## Follow-up Proposals

Once extension elements are generated, separate proposals needed for:
- Content Controls (SDT) API (uses Word 2010+ elements)
- Slicers/Timelines API (uses Excel 2010+ elements)  
- Modern Animations (uses PowerPoint 2010+ elements)
