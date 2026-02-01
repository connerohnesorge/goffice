# Element Type Coverage Implementation Notes

## Schema Files Identified

### VML Namespaces
- `schemas-microsoft-com_vml.json` - Main VML namespace (urn:schemas-microsoft-com:vml)
- `schemas-microsoft-com_office_office.json` - Office extensions (urn:schemas-microsoft-com:office:office)
- `schemas-microsoft-com_office_excel.json` - Excel VML extensions (xvml:)
- `schemas-microsoft-com_office_powerpoint.json` - PowerPoint VML extensions (pvml:)

### Current Element Coverage
- Wordprocessing: 47 files
- Spreadsheet: 100 files
- Presentation: 18 files

### Missing Elements to Add
1. VML Elements (v:*) - shapes, shapetype, path, fill, stroke, etc.
2. Office VML Extensions (o:*) - shapedefaults, shapelayout, extrusion, etc.
3. Excel VML (xvml:*) - ClientData for form controls
4. PowerPoint VML (pvml:*) - textdata, iscomment
5. Alternate Content (mc:*) - needs manual implementation
6. Custom XML (w:customXml*)
7. SDT elements (w:sdt*)
8. Extended Properties
9. ID/Tracking elements

## Code Generator Updates Needed

### Files to Modify
1. `cmd/gen-go-wordprocessing/main.go` - Add VML schema support
2. `cmd/gen-go-spreadsheet/main.go` - Add VML schema support
3. `cmd/gen-go-presentation/main.go` - Add VML schema support

### Key Changes Required
1. Add `isVMLSchema()` function to detect VML-related schemas
2. Add namespace mappings for VML URIs
3. Update schema filtering to include VML files
4. Ensure VML elements are generated in appropriate packages

## Implementation Strategy

### Phase 1: VML Elements (In Progress)
- Update generators to handle VML schemas
- Run generators to create VML element files
- Verify compilation

### Phase 2: Alternate Content
- Manually implement mc:AlternateContent, mc:Choice, mc:Fallback
- These are structural elements for backward compatibility

### Phase 3: Remaining Elements
- Custom XML elements
- SDT elements
- Extended Properties
- ID/Tracking elements

## Notes
- NEVER manually edit files in `*/elements/` directories
- Always use code generators
- VML elements should be shared across wordprocessing, spreadsheet, and presentation
- Consider creating a shared `vml/` package for common VML elements

## Wave 1 Execution Status - $(date)

### Launched Tasks (All Running in Parallel)
1. **add-validation-semantic-constraints** (bg_f2ec30d2) - deep agent - 18s elapsed
2. **add-comprehensive-element-type-coverage** (bg_c193f70d) - unspecified-high agent - 11s elapsed  
3. **add-style-inheritance-and-defaults** (bg_027c326f) - quick agent - 5s elapsed

### Expected Completion
- style-inheritance: ~1-2 weeks (fastest)
- element-type-coverage: ~2-3 weeks
- validation-semantic-constraints: ~8 weeks (longest, critical path)

### Wave 2 Blocked Until
- validation-semantic-constraints completes (blocks strict-namespace-support, content-controls, change-tracking, data-validation)
- element-type-coverage completes (blocks linq-support, formula-evaluation)

### Next Actions
- Monitor background tasks for completion
- Launch Wave 2 tasks as soon as dependencies are satisfied
- Update todo list as tasks complete
