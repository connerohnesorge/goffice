# Implementation Tasks

## Phase 1: Generator Infrastructure (1 week)

### 1.1 Schema File Metadata
- [ ] Create `SchemaFile` struct in `cmd/gen-go-wordprocessing/schema_file.go`
  - Fields: Path, Namespace, NamespacePrefix, Version, TargetApp
  - Validation: Struct compiles, has correct types

- [ ] Implement `parseSchemaFile()` function
  - Extract version from filename regex
  - Determine target app from filename patterns
  - Load JSON schema to get actual namespace
  - Derive namespace prefix (w14, w15, x14, a14, etc.)
  - Validation: Unit tests pass for various filenames

- [ ] Add version mapping utilities
  - `yearToVersion()` converts 2010 → Office2010
  - `versionToSuffix()` converts Office2010 → "14"
  - `prefixToVersion()` converts "w14" → Office2010
  - Validation: Mapping tests pass

### 1.2 Update WordProcessing Generator
- [ ] Modify `cmd/gen-go-wordprocessing/main.go` schema loading
  - Replace `if strings.Contains(f.Name(), "main.json")` filter
  - Add pattern matching for Word extension schemas
  - Add pattern matching for Drawing extension schemas
  - Sort schemas by version (main first, then chronological)
  - Validation: Generator finds 60+ schemas (not just 1)

- [ ] Update `collectTypesWithVersion()` function
  - Accept SchemaFile parameter (not just path)
  - Extract and store version metadata
  - Register namespace in global namespace map
  - Validation: Types collected include version metadata

- [ ] Add namespace constant generation
  - Generate `NamespaceWord2010` through `NamespaceWord2024`
  - Generate `NamespaceDrawing2010` through `NamespaceDrawing2018`
  - Generate `NamespacePrefixMap` for all namespaces
  - Validation: Constants file compiles, has correct URLs

### 1.3 Update Spreadsheet Generator
- [ ] Apply same changes to `cmd/gen-go-spreadsheet/main.go`
  - Pattern matching for Excel extension schemas
  - Include shared Drawing extensions
  - Version metadata extraction
  - Validation: Finds 80+ schemas

- [ ] Generate Excel namespace constants
  - `NamespaceExcel2010` through `NamespaceExcel2024`
  - Validation: Constants correct for spreadsheetml

### 1.4 Update Presentation Generator
- [ ] Apply same changes to `cmd/gen-go-presentation/main.go`
  - Pattern matching for PowerPoint extension schemas
  - Include shared Drawing extensions
  - Version metadata extraction
  - Validation: Finds 50+ schemas

- [ ] Generate PowerPoint namespace constants
  - `NamespacePowerPoint2010` through `NamespacePowerPoint2022`
  - Validation: Constants correct for presentationml

### 1.5 Generator Testing
- [ ] Write generator schema loading tests
  - Test schema file parsing from various filenames
  - Test version extraction accuracy
  - Test namespace prefix derivation
  - Test filtering (only relevant schemas loaded)
  - Validation: All generator tests pass

## Phase 2: Framework Extensions (1 week)

### 2.1 Version System
- [x] Create `openxml/version.go`
  - Define `FileFormatVersion` enum
  - Office2007, Office2010, Office2013, Office2016, Office2019, Office2021, Office2022, Office2023, Office2024, Office2025, Microsoft365
  - Implement `String()` method
  - Validation: Version enum compiles and prints correctly

- [ ] Update `openxml/metadata.go`
  - Add `AvailableInVersion() FileFormatVersion` to `IElementMetadata` interface
  - Add `availableInVersion` field to `elementMetadata` struct
  - Implement getter method
  - Validation: Metadata interface extends without breaking existing code

### 2.2 AlternateContent Support
- [x] Create `openxml/alternate_content.go`
  - Implement `AlternateContent` composite element
  - Implement `Choice` element with `Requires` attribute
  - Implement `Fallback` element
  - Add `SelectContent(targetVersion)` method
  - Validation: Elements compile, satisfy Element interface

- [x] Write AlternateContent tests
  - Test Choice/Fallback structure parsing
  - Test `SelectContent()` with different target versions
  - Test version requirement checking
  - Validation: All AlternateContent tests pass

### 2.3 Unknown Element Preservation
- [x] Create `openxml/unknown_element.go`
  - Implement `UnknownElement` type with raw XML storage
  - Implement `LocalName()`, `NamespaceURI()` methods
  - Implement `WriteXML()` to output raw XML unchanged
  - Validation: UnknownElement compiles and satisfies Element

- [ ] Update `openxml/xml_reader.go`
  - Add `readUnknownElement()` method
  - Capture entire unknown element subtree as raw XML
  - Return `UnknownElement` when element type not registered
  - Validation: Reader handles unknown elements without errors

- [x] Write unknown element tests
  - Test preservation of completely unknown namespaces
  - Test roundtrip: read unknown, write back, verify identical
  - Test unknown elements nested in known elements
  - Validation: Preservation tests pass

### 2.4 Namespace Registration
- [x] Update `openxml/namespaces.go`
  - Add constants for all extension namespaces (100+ constants)
  - Add `NamespacePrefixMap` global map
  - Add `RegisterNamespace()` function
  - Validation: Namespace constants compile, map correct

- [ ] Update XML reader to use namespace map
  - Look up prefixes from `NamespacePrefixMap`
  - Register extension namespaces during parse
  - Validation: Reader uses correct prefixes for extensions

### 2.5 Framework Testing
- [x] Write version system tests
  - Test version enum values and ordering
  - Test version string formatting
  - Validation: Version tests pass

- [x] Write namespace registration tests
  - Test prefix lookup for all namespaces
  - Test namespace constant values
  - Validation: Namespace tests pass

## Phase 3: Code Generation (3 days)

### 3.1 Regenerate WordProcessing Elements
- [ ] Run `go run ./cmd/gen-go-wordprocessing`
  - Generate with extension schemas included
  - Validation: Generation completes without errors

- [ ] Verify generated code
  - Check `wordprocessing/elements/elements.go` compiles
  - Count new element types (should be 60-80 more)
  - Spot-check extension elements (ContentControl, etc.)
  - Validation: Code compiles, element count increased

- [ ] Fix any generation issues
  - Handle name conflicts
  - Fix import statements
  - Adjust code templates if needed
  - Validation: All fixes applied, code compiles

### 3.2 Regenerate Spreadsheet Elements
- [ ] Run `go run ./cmd/gen-go-spreadsheet`
  - Generate with extension schemas
  - Validation: Generation completes

- [ ] Verify generated code
  - Check compilation
  - Count new types (80-100 more)
  - Validation: Code compiles

### 3.3 Regenerate Presentation Elements
- [ ] Run `go run ./cmd/gen-go-presentation`
  - Generate with extension schemas
  - Validation: Generation completes

- [ ] Verify generated code
  - Check compilation
  - Count new types (50-70 more)
  - Validation: Code compiles

### 3.4 Code Quality
- [ ] Run `golangci-lint run --fix ./...`
  - Fix all linting issues in generated code
  - Adjust generators if linting issues are systematic
  - Validation: golangci-lint passes with zero errors

- [ ] Run `go test ./...`
  - Ensure all existing tests still pass
  - Fix any test breakages
  - Validation: Zero test failures

## Phase 4: Validation Integration (3 days)

### 4.1 Version-Aware Validation
- [ ] Update `openxml/validation/context.go`
  - Add `TargetVersion FileFormatVersion` field to `ValidationContext`
  - Update `NewValidationContext()` to accept version parameter
  - Validation: Context struct compiles

- [ ] Update `openxml/validation/validator.go`
  - Modify `Validate()` to accept target version
  - Add element version checking before other validation
  - Generate version mismatch errors
  - Validation: Validator compiles

- [ ] Create version validation error messages
  - "Element X not available in Office 2010 (introduced in Office 2013)"
  - Include element path and version information
  - Validation: Error messages are clear and helpful

### 4.2 AlternateContent Validation
- [ ] Create `openxml/validation/alternate_content_validator.go`
  - Validate AlternateContent has exactly one Choice and one Fallback
  - Validate Choice has `Requires` attribute
  - Validate `Requires` references valid namespace prefix
  - Validation: Validator compiles and has tests

- [ ] Integrate into validation pipeline
  - Register AlternateContent validators
  - Run on all AlternateContent elements
  - Validation: AlternateContent validation works

### 4.3 Version Detection API
- [ ] Create `openxml/version_detector.go`
  - Implement `DetectMinimumVersion(doc)` function
  - Walk all elements, find highest version requirement
  - Return minimum Office version needed
  - Validation: Version detection works correctly

- [ ] Write version detection tests
  - Test document with only Office 2007 elements → Office2007
  - Test document with Office 2013 elements → Office2013
  - Test mixed version document → highest version
  - Validation: Detection tests pass

### 4.4 Validation Testing
- [ ] Write version validation tests
  - Test Office 2010 element in Office 2007 validation → error
  - Test Office 2010 element in Office 2013 validation → pass
  - Test multiple version violations → multiple errors
  - Validation: All validation tests pass

- [ ] Write AlternateContent validation tests
  - Test valid AlternateContent structure → pass
  - Test missing Choice or Fallback → error
  - Test invalid `Requires` value → error
  - Validation: AlternateContent tests pass

## Phase 5: Integration Testing (1 week)

### 5.1 Test Document Creation
- [ ] Create Word test document with Office 2010 features
  - Include Content Controls (w14)
  - Include Drawing Canvas (w14)
  - Save as `testdata/word2010-features.docx`
  - Validation: Document opens in Microsoft Word 2010+

- [ ] Create Excel test document with Office 2013 features
  - Include Slicers (x14)
  - Include Timelines (x15)
  - Save as `testdata/excel2013-features.xlsx`
  - Validation: Document opens in Microsoft Excel 2013+

- [ ] Create Excel test document with Office 2025 features
  - Include Pivot Data Source (x25) if available
  - Save as `testdata/excel2025-features.xlsx`
  - Validation: Document opens in Microsoft Excel 2025+

- [ ] Create PowerPoint test document with Office 2016 features
  - Include modern animations
  - Save as `testdata/powerpoint2016-features.pptx`
  - Validation: Document opens in Microsoft PowerPoint 2016+

### 5.2 Roundtrip Tests
- [ ] Write `wordprocessing/roundtrip_extension_test.go`
  - Open Word document with extensions
  - Verify extension elements read correctly
  - Save document
  - Re-open and verify elements unchanged
  - Validation: Roundtrip test passes

- [ ] Write `spreadsheet/roundtrip_extension_test.go`
  - Same for Excel extensions
  - Validation: Roundtrip test passes

- [ ] Write `presentation/roundtrip_extension_test.go`
  - Same for PowerPoint extensions
  - Validation: Roundtrip test passes

### 5.3 Unknown Element Tests
- [ ] Create document with future/unknown extensions
  - Manually add XML with unknown namespace
  - Validation: Document prepared

- [ ] Test unknown element preservation
  - Open document with unknown elements
  - Verify UnknownElement instances created
  - Save document
  - Verify unknown XML preserved byte-for-byte
  - Validation: Unknown preservation test passes

### 5.4 AlternateContent Integration Tests
- [ ] Create document with AlternateContent blocks
  - Include Choice (Office 2013) and Fallback (Office 2007)
  - Validation: Test document ready

- [ ] Test AlternateContent reading
  - Open document
  - Verify AlternateContent structure parsed
  - Test `SelectContent()` with different versions
  - Validation: AlternateContent reading test passes

- [ ] Test AlternateContent writing
  - Create document with AlternateContent programmatically
  - Save and verify XML structure correct
  - Open in Microsoft Office, verify works
  - Validation: Writing test passes

### 5.5 Version Validation Integration Tests
- [ ] Test document version detection
  - Create document with Office 2010 elements
  - Call `DetectMinimumVersion()`
  - Verify returns Office2010
  - Validation: Detection integration test passes

- [ ] Test validation with version constraints
  - Create document with Office 2013 elements
  - Validate against Office 2010 → should have errors
  - Validate against Office 2013 → should pass
  - Validation: Version validation integration passes

### 5.6 Compatibility Testing
- [ ] Test with Microsoft Office
  - Save document with goffice using extensions
  - Open in Microsoft Word/Excel/PowerPoint
  - Verify features work correctly
  - Validation: Microsoft Office compatibility confirmed

- [ ] Test with LibreOffice
  - Open goffice-created documents in LibreOffice
  - Verify basic compatibility (may not support all extensions)
  - Validation: No crashes, data preserved

## Phase 6: Documentation (2 days)

### 6.1 API Documentation
- [ ] Document `FileFormatVersion` enum
  - Godoc for each version constant
  - Usage examples
  - Validation: Godoc renders correctly

- [ ] Document `AlternateContent` API
  - Package documentation for `openxml/alternate_content.go`
  - Method documentation
  - Code examples
  - Validation: Documentation complete

- [ ] Document `UnknownElement` API
  - Explain preservation behavior
  - Forward compatibility notes
  - Validation: Documentation clear

- [ ] Document validation version parameter
  - Update `openxml/validation` package docs
  - Examples of version validation
  - Validation: Validation docs updated

### 6.2 User Guide
- [ ] Create extension schema support guide
  - Explain what extensions are supported
  - Office version compatibility matrix
  - How to use extension elements
  - Validation: Guide is helpful for users

- [ ] Document version detection
  - How to detect document minimum version
  - How to validate against target version
  - Best practices
  - Validation: Version detection documented

- [ ] Document unknown element handling
  - Explain forward compatibility
  - How UnknownElement works
  - Migration path for future Office versions
  - Validation: Unknown element docs complete

### 6.3 Developer Guide
- [ ] Update generator documentation
  - Explain schema loading changes
  - How to add new schema patterns
  - Version metadata extraction
  - Validation: Generator docs updated

- [ ] Document namespace handling
  - Extension namespace constants
  - Prefix derivation rules
  - Adding new namespaces
  - Validation: Namespace docs complete

- [ ] Create migration notes
  - Backward compatibility guaranteed
  - No breaking changes
  - New capabilities enabled
  - Validation: Migration notes clear

### 6.4 Example Updates
- [ ] Add extension element examples
  - Example using Content Controls (Word 2010)
  - Example using Slicers (Excel 2010)
  - Validation: Examples work

- [ ] Add version validation example
  - Show how to validate for specific Office version
  - Show version detection
  - Validation: Examples functional

## Validation Checkpoints

### After Phase 1
- [ ] All generators find 100+ schemas (not just main.json)
- [ ] Version metadata extraction works correctly
- [ ] Namespace prefix derivation is accurate
- [ ] Generator tests pass

### After Phase 2
- [ ] AlternateContent elements work correctly
- [ ] UnknownElement preserves arbitrary XML
- [ ] Version system is complete (Office2007, 2010, 2013, 2016, 2019, 2021, 2022, 2023, 2024, 2025, Microsoft365)
- [ ] Framework tests pass

### After Phase 3
- [ ] 200-300 new element types generated
- [ ] All generated code compiles
- [ ] golangci-lint passes
- [ ] Existing tests still pass

### After Phase 4
- [ ] Version-aware validation works
- [ ] AlternateContent validation works
- [ ] Version detection API functional
- [ ] Validation tests pass

### After Phase 5
- [ ] Extension elements roundtrip correctly
- [ ] Unknown elements preserved
- [ ] AlternateContent integration works
- [ ] Microsoft Office compatibility confirmed

### After Phase 6
- [ ] All documentation complete
- [ ] Examples work and are helpful
- [ ] Developer guide updated
- [ ] Users can use new features

## Dependencies

**Sequential Dependencies:**
- Phase 2 depends on Phase 1 (need version system for framework)
- Phase 3 depends on Phase 2 (need framework for generation)
- Phase 4 depends on Phase 3 (need generated elements for validation)
- Phase 5 depends on Phase 4 (need validation for integration tests)
- Phase 6 can happen in parallel with Phase 5

**Parallel Work:**
- Within Phase 1: All three generators can be updated simultaneously
- Within Phase 2: Framework components are independent
- Within Phase 3: All three regenerations can happen in parallel (after framework ready)
- Within Phase 5: Different document type tests can run in parallel

## Success Criteria

All tasks completed AND:
- [ ] All generators load extension schemas (verified by schema count)
- [ ] 200-300 new element types generated (verified by type count)
- [ ] Extension namespace constants defined (verified in openxml/namespace.go)
- [ ] AlternateContent/Choice/Fallback implemented and tested
- [ ] UnknownElement preserves unknown extensions
- [ ] Version-aware validation works for Office 2010-2025
- [ ] Roundtrip tests pass with extension elements
- [ ] No regression (all existing tests pass)
- [ ] golangci-lint passes (zero errors/warnings)
- [ ] Can open Office 2010+ documents without data loss
- [ ] Can detect document minimum Office version
- [ ] Documentation complete and accurate
- [ ] Microsoft Office compatibility verified
