# Change: Add Extension Schema Support (Office 2010-2025)

## Why

The code generators currently **load only main.json schemas**, skipping **100+ extension schemas** that define Office 2010-2025 features. This leaves goffice unable to read or manipulate documents using modern Office capabilities, severely limiting real-world usability.

**Current Situation:**
- Generators filter for `"main.json"` only: `if strings.Contains(f.Name(), "main.json")`
- **200-300 element types** from extension schemas are not generated
- Documents with modern features (Office 2010+) cannot be fully read or written
- Extension elements appear as unknown XML, losing data on roundtrip

**Missing Features Examples:**
- **Word**: Content Controls (SDT), Drawing Canvas (2010+), Format Lock (2024)
- **Excel**: Slicers, Timelines, Sparklines (2010+), Rich Data (2017+), Pivot Auto Refresh (2024), Pivot Data Source (2025)
- **PowerPoint**: Modern Animations, Roaming Settings (2012+), 3D Models (2018+)
- **DrawingML**: Decorative flags, SVG Support, Animation Models (2010-2024)

**Business Impact:**
- Cannot process modern Office documents correctly
- Data loss on open/save cycles for documents using extensions
- Feature parity with Open-XML-SDK blocked
- Users must avoid documents created in Office 2010 or later

**Root Cause:** Generator design limitation - hardcoded filter instead of systematic schema loading.

## What Changes

**Generator Updates (All Generators):**
- **Expand schema loading** from `main.json` only to include all relevant extension schemas
- **Add namespace mapping** for extension namespaces (e.g., `w14:`, `w15:`, `x14:`, `p14:`)
- **Generate extension elements** with proper namespace prefixes and version metadata
- **Version awareness** in element metadata (Office 2010, 2013, 2016, 2019, 2021, 2022, 2023, 2024, 2025, M365)

**Framework Enhancements:**
- **AlternateContent handling** for Choice/Fallback version compatibility patterns
- **Unknown element preservation** to avoid data loss for unsupported extensions
- **Namespace registration** for all extension namespaces during XML parsing
- **Version detection** API to identify document Office version requirements

**Validation Updates:**
- **Version-aware validation** - elements validated against their introduction version
- **Extension schema rules** loaded for Office 2010+ validation
- **AlternateContent validation** to ensure Choice/Fallback blocks are well-formed

**Code Organization:**
- Generated elements in existing files (single-file approach maintained)
- Namespace constants generated for all extension namespaces
- Version metadata on element types (AvailableInVersion field)

**Breaking Changes:** None. This is additive - new element types enable new features without changing existing APIs.

## Impact

**Affected Specs:**
- `framework` - ADDED AlternateContent support, unknown element preservation, namespace handling
- `validation` - MODIFIED to support version-aware validation
- `wordprocessing-elements` - ADDED 60-80 extension element types (Office 2010-2025)
- `spreadsheet-elements` - ADDED 80-100 extension element types
- `presentation-elements` - ADDED 50-70 extension element types

**New Capabilities:**
- Read documents with Office 2010-2025 features without data loss
- Write documents using modern Office extensions
- Validate against specific Office versions (2010, 2013, 2016, 2019, 2021, 2022, 2023, 2024, 2025, M365)
- Detect minimum Office version required for a document
- Preserve unknown extensions for forward compatibility

**Affected Code:**
- `cmd/gen-go-wordprocessing/main.go` - MODIFIED: Schema loading logic
- `cmd/gen-go-spreadsheet/main.go` - MODIFIED: Schema loading logic
- `cmd/gen-go-presentation/main.go` - MODIFIED: Schema loading logic
- `openxml/namespace.go` - MODIFIED: Add extension namespace constants
- `openxml/element.go` - MODIFIED: Add version metadata
- `openxml/xml_reader.go` - MODIFIED: Preserve unknown elements
- `openxml/validation/validator.go` - MODIFIED: Version-aware validation
- `wordprocessing/elements/` - GENERATED: 60-80 new element types
- `spreadsheet/elements/` - GENERATED: 80-100 new element types
- `presentation/elements/` - GENERATED: 50-70 new element types

**User Impact:**
- Existing code continues to work (backward compatible)
- New documents with extensions can be read without data loss
- Can now build features requiring extensions (Content Controls, Slicers, etc.)
- Validation catches version-specific errors

## Key Design Decisions

### 1. Load All Extension Schemas
**Decision:** Load ALL extension schemas (Office 2010-2025), not a filtered subset.

**Rationale:**
- Disk space is cheap, unused elements are no-ops
- Avoids arbitrary version cutoffs
- Future-proof as new Office versions release
- Matches Open-XML-SDK behavior (complete coverage)

**Alternatives Considered:**
- Load only specific years (2010, 2013, 2016, 2019) - REJECTED: Arbitrary, incomplete
- Load on-demand at runtime - REJECTED: Complexity, runtime overhead
- Manual curation - REJECTED: Maintenance burden

### 2. Single-File Element Generation
**Decision:** Keep all elements in single file per package (e.g., `wordprocessing/elements/elements.go`).

**Rationale:**
- Consistent with current generator pattern
- Simpler build process
- File splitting can be done later if size becomes issue
- 50K-80K lines is manageable for Go compiler

**Alternatives Considered:**
- Split by namespace (elements/main.go, elements/w14.go, etc.) - DEFERRED: Can do later if needed
- Split by feature area - REJECTED: Unclear boundaries

### 3. AlternateContent as Framework Feature
**Decision:** Handle AlternateContent (Choice/Fallback) at framework level, not in each element type.

**Rationale:**
- Pattern applies uniformly across all element types
- Single implementation point reduces bugs
- Matches Open-XML-SDK architecture
- Enables automatic fallback selection based on target version

**Implementation:** AlternateContent, Choice, Fallback as special composite elements in openxml/

### 4. Unknown Element Preservation
**Decision:** Preserve unknown elements as opaque XML during roundtrip.

**Rationale:**
- Forward compatibility - handle future Office versions gracefully
- Avoid data loss when opening documents with unsupported extensions
- Users can still read/write documents even if not all features are supported

**Implementation:** UnknownElement type stores raw XML, re-serializes unchanged

### 5. Version Metadata on Elements
**Decision:** Add `AvailableInVersion() FileFormatVersions` method to element metadata.

**Rationale:**
- Enables version-aware validation
- Allows runtime version detection (scan document, find newest element)
- Generator can extract version from schema file names (e.g., "2010" from filename)

**Implementation:** Generator parses schema filename, embeds version in generated metadata

### 6. Namespace Prefix Convention
**Decision:** Use standard Office namespace prefixes (w14, w15, x14, x15, p14, p15, a14, etc.).

**Rationale:**
- Matches Microsoft Office output
- Familiar to developers reading OOXML specs
- Consistent with existing main namespace prefixes (w, x, p, a)

**Implementation:** Namespace constants generated from schema target namespaces

## Implementation Scope

### Phase 1: Generator Infrastructure (1 week)
- Update schema loading logic in all generators
- Add namespace detection and mapping
- Add version metadata extraction from schema filenames
- Test generators compile and run (don't generate yet)

### Phase 2: Framework Extensions (1 week)
- Implement AlternateContent, Choice, Fallback elements
- Add UnknownElement for preservation
- Update XML reader to preserve unknown elements
- Add namespace registration for extensions
- Unit tests for new framework features

### Phase 3: Code Generation (3 days)
- Run generators to produce extension elements
- Verify generated code compiles
- Fix any generator bugs discovered
- Run golangci-lint and fix issues

### Phase 4: Validation Integration (3 days)
- Add version parameter to validation
- Load extension schema rules
- Implement version availability checking
- AlternateContent validation rules
- Unit tests for version-aware validation

### Phase 5: Integration Testing (1 week)
- Create test documents with Office 2010-2024 features
- Roundtrip tests (open, save, verify)
- Validation tests for version compatibility
- Unknown element preservation tests
- Compatibility testing with Microsoft Office

### Phase 6: Documentation (2 days)
- Document extension schema support
- Version compatibility matrix
- Migration guide (none needed - backward compatible)
- Generator developer notes

**Total: 3-4 weeks**

### Out of Scope
- **High-level APIs for extension elements** - This proposal only generates element types; using them (e.g., Content Control API) is separate proposals
- **Full semantic validation** - Only schema structure validation, not business rules
- **Automatic version downgrading** - Converting Office 2019 document to 2010 format (complex transformation)
- **Extension schema documentation** - Too many elements to document individually

## Dependencies

**Existing:**
- OpenXML-SDK schema files (already present in `Open-XML-SDK/data/schemas/`)
- Generator framework (cmd/gen-go-*)
- OpenXML framework (openxml/)
- Validation framework (openxml/validation/)

**New:** None

**Blocks:**
- Content Controls API (needs Word extension elements)
- Slicers/Timelines API (needs Excel extension elements)
- Modern Animations API (needs PowerPoint extension elements)

## Success Criteria

- [ ] All generators load extension schemas (not just main.json)
- [ ] 200-300+ new element types generated across Word/Excel/PowerPoint
- [ ] Extension namespace constants generated (w14, w15, x14, x15, p14, a14, etc.)
- [ ] AlternateContent, Choice, Fallback elements implemented
- [ ] UnknownElement preserves unsupported extensions
- [ ] Version-aware validation works for Office 2010-2025
- [ ] Existing tests pass (no regression)
- [ ] Roundtrip tests pass with extension elements
- [ ] golangci-lint passes on generated code
- [ ] Can open Office 2010+ documents without data loss
- [ ] Can detect minimum Office version for a document

## Validation Approach

**Before Generation:**
```bash
# Count current element types
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Expected: ~619 types (baseline)
```

**After Generation:**
```bash
# Should have 200-300 more across all packages
grep "^type.*struct" wordprocessing/elements/elements.go | wc -l
# Expected: ~680-700 types (60-80 new Word extensions)

grep "^type.*struct" spreadsheet/elements/elements.go | wc -l
# Expected: +80-100 new Excel extensions

grep "^type.*struct" presentation/elements/elements.go | wc -l
# Expected: +50-70 new PowerPoint extensions
```

**Functional Validation:**
```bash
# Create test document with Word 2010 features, roundtrip
go test ./wordprocessing/... -run TestExtensionRoundtrip

# Validate document against Office 2016
go test ./openxml/validation/... -run TestVersionValidation
```

## Risks & Mitigation

### Risk: Code Size Increase
**Impact:** High - Generated files could become very large (80K+ lines)
**Probability:** High
**Mitigation:**
- Go compiler handles large files well (tested up to 100K lines)
- Can split files later if needed (backward compatible change)
- Build times acceptable (<30s for full rebuild)

### Risk: Namespace Conflicts
**Impact:** Medium - Extension elements might conflict with main elements
**Probability:** Low
**Mitigation:**
- Namespaces are distinct in XML
- Go struct names include namespace prefix if needed (W14ContentControl vs ContentControl)
- Generator has conflict resolution logic

### Risk: Schema Changes
**Impact:** Medium - Microsoft might change extension schemas
**Probability:** Low (schemas are stable once published)
**Mitigation:**
- Lock to specific Open-XML-SDK version
- Regenerate when updating OpenXML-SDK
- Version-aware validation catches incompatibilities

### Risk: Breaking Changes in Generated Code
**Impact:** High - Regeneration could break existing code
**Probability:** Low
**Mitigation:**
- Existing elements unchanged (only additions)
- Extensive testing before/after generation
- Can always revert to previous generation

### Risk: Validation Complexity
**Impact:** Medium - Version-aware validation adds complexity
**Probability:** Medium
**Mitigation:**
- Start with simple version checking
- Extensive unit tests for edge cases
- Clear error messages for version conflicts

## Follow-up Proposals

Once extension elements are generated:
- **Content Controls API** (Word extensions, SDT elements)
- **Slicers and Timelines API** (Excel extensions)
- **Modern Animations API** (PowerPoint extensions)
- **Rich Data Types** (Excel extensions, linked data)
- **Drawing Canvas** (Word extensions, advanced graphics)
- **3D Models Support** (Office 365 extensions)
