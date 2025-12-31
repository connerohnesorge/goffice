# goffice Change Proposals Implementation Timeline

This document outlines the chronological implementation order for all change proposals in `spectr/changes/`. The order is determined by dependencies, foundational requirements, and strategic value.

## Timeline Overview

| Phase | Proposal | Category | Dependencies |
|-------|----------|----------|--------------|
| 1 | enable-drawingml-pdf-rendering | Infrastructure | None |
| 2 | enable-chart-rendering | Infrastructure | enable-drawingml-pdf-rendering |
| 3 | add-extension-schema-support | Infrastructure | None |
| 4 | add-dotnet-cross-runtime-testing | Testing | None |
| 5 | add-e2e-visual-testing | Testing | add-dotnet-cross-runtime-testing |
| 6 | add-visual-comparison | Testing | enable-drawingml-pdf-rendering |
| 7 | add-row-column-operations | Spreadsheet | None |
| 8 | add-pivottable-api | Spreadsheet | add-row-column-operations (partial) |
| 9 | add-comments-api | Word | None |
| 10 | add-track-changes-api | Word | None |
| 11 | add-form-fields-api | Word | None |
| 12 | add-mail-merge-engine | Word | add-form-fields-api (partial) |
| 13 | add-smartart-support | DrawingML | add-extension-schema-support |

---

## Phase 1: Enable DrawingML PDF Rendering

**Proposal**: `enable-drawingml-pdf-rendering`
**Category**: Core Infrastructure
**Priority**: Critical (10/10)
**Estimated Effort**: 6 weeks

### Why First
This is the **most critical blocker** in the codebase. 12 renderer files (3,543 LOC) are disabled with `.wip` extension because they depend on a non-existent `Page` abstraction API in `pdf/core/`. Without this:
- Charts cannot render in any PDF output
- Shapes with fills/strokes don't render
- Images within shapes don't display
- Visual effects (shadows, glows) are missing
- Text within shapes cannot render

### What It Enables
- All subsequent PDF rendering improvements
- Chart rendering infrastructure
- Visual comparison testing (requires PDF output)
- High-fidelity document export

### Key Deliverables
- `Page` interface in `pdf/core/`
- pdfcpu-backed implementation
- 12 `.wip` files re-enabled as `.go`
- DrawingML integration tests

---

## Phase 2: Enable Chart Rendering

**Proposal**: `enable-chart-rendering`
**Category**: PDF Infrastructure
**Priority**: High (8/10)
**Estimated Effort**: 3-4 weeks
**Depends On**: enable-drawingml-pdf-rendering

### Why Second
With the Page API from Phase 1, we can immediately enable the 3,543 lines of existing chart/shape rendering code. This is high-value work because:
- Code already exists (just blocked)
- Charts are essential for Excel and PowerPoint PDF export
- Minimal new development required

### What It Enables
- Bar, Line, Pie chart rendering to PDF
- Shape geometry rendering
- Fill rendering (solid, gradient, pattern)
- Effects rendering (shadows, glows, reflections)
- Text within shapes

### Key Deliverables
- Re-enabled chart_renderer.go, shape_renderer.go, etc.
- Integration with presentation PDF rendering
- Integration with spreadsheet PDF rendering

---

## Phase 3: Add Extension Schema Support

**Proposal**: `add-extension-schema-support`
**Category**: Core Infrastructure
**Priority**: High (8/10)
**Estimated Effort**: 3-4 weeks

### Why Third
The code generators currently load only `main.json` schemas, skipping **100+ extension schemas** that define Office 2010-2024 features. This leaves goffice unable to:
- Read documents using modern Office features
- Preserve extension elements on roundtrip
- Generate documents with Office 2010+ features

This is foundational infrastructure that:
- Enables SmartArt support (Phase 13)
- Enables Content Controls (future)
- Enables Slicers/Timelines (future)
- Prevents data loss with modern documents

### What It Enables
- 200-300 new element types
- AlternateContent handling
- Unknown element preservation
- Version-aware validation
- SmartArt schema loading (required for Phase 13)

### Key Deliverables
- Updated schema loading in all generators
- Extension namespace constants (w14, w15, x14, etc.)
- AlternateContent, Choice, Fallback elements
- UnknownElement preservation

---

## Phase 4: Add .NET Cross-Runtime Testing

**Proposal**: `add-dotnet-cross-runtime-testing`
**Category**: Testing Infrastructure
**Priority**: High (8/10)
**Estimated Effort**: 4-5 weeks

### Why Fourth
This creates foundational testing infrastructure to validate API parity with Microsoft's Open-XML-SDK. Benefits:
- Validates Go implementation matches .NET SDK behavior
- Catches regressions automatically
- Provides evidence of feature parity
- Foundation for E2E visual testing

Can run in parallel with Phases 1-3 if resources allow.

### What It Enables
- Cross-runtime document comparison
- API parity validation
- Regression detection
- E2E visual testing (Phase 5)

### Key Deliverables
- Nix environment with .NET SDK 9
- Test scenario framework (YAML-based)
- Go and C# document generation bridges
- Three-level comparison (XML, binary, visual)
- CI/CD integration

---

## Phase 5: Add E2E Visual Testing

**Proposal**: `add-e2e-visual-testing`
**Category**: Testing Infrastructure
**Priority**: Medium-High (7/10)
**Estimated Effort**: 3-4 weeks
**Depends On**: add-dotnet-cross-runtime-testing

### Why Fifth
Builds on cross-runtime testing to add visual verification. Ensures PPTX/DOCX/XLSX generation produces visually identical output to .NET SDK.

### What It Enables
- Visual regression testing
- Chart rendering accuracy validation
- Layout verification
- Automated fidelity checks

### Key Deliverables
- PDF rendering pipeline for comparison
- Pixel-level image diffing
- HTML comparison reports
- CI integration for visual checks

---

## Phase 6: Add Visual Comparison for PDF Fidelity

**Proposal**: `add-visual-comparison`
**Category**: Testing
**Priority**: Medium (6/10)
**Estimated Effort**: 2-3 weeks
**Depends On**: enable-drawingml-pdf-rendering

### Why Sixth
With PDF rendering enabled (Phase 1), we can implement automated visual comparison to validate PDF output fidelity against Microsoft Office output.

### What It Enables
- Automated PDF fidelity validation
- Visual regression detection
- Continuous rendering quality assurance

### Key Deliverables
- `pdf/comparison` package
- PDF-to-PNG conversion utilities
- Pixel-level difference detection
- Visual diff report generation

---

## Phase 7: Add Row/Column Operations

**Proposal**: `add-row-column-operations`
**Category**: Spreadsheet
**Priority**: Critical (9/10)
**Estimated Effort**: 6 weeks

### Why Seventh
This is a **Tier 1 priority gap** for spreadsheet manipulation. Users cannot:
- Insert rows/columns
- Delete rows/columns
- Group rows/columns for outlining

The hard part is cell reference updates - formulas, named ranges, merged cells, conditional formatting, charts must all be updated correctly.

### What It Enables
- Dynamic spreadsheet manipulation
- Template expansion/contraction
- Batch data processing
- PivotTable enhancements (Phase 8)

### Key Deliverables
- Cell reference parser (A1, R1C1, absolute/relative)
- Formula tokenizer and rewriter
- Sheet.InsertRows(), DeleteRows(), InsertColumns(), DeleteColumns()
- Comprehensive reference update engine
- Row/column grouping API

---

## Phase 8: Complete PivotTable API

**Proposal**: `add-pivottable-api`
**Category**: Spreadsheet
**Priority**: High (9/10)
**Estimated Effort**: 5-6 weeks
**Depends On**: add-row-column-operations (partial - for reference updates)

### Why Eighth
Existing PivotTable infrastructure (~560 LOC) provides basic field management but lacks:
- Cache population from source data
- Field validation
- PivotField property exposure
- Complete roundtrip persistence

### What It Enables
- Complete pivot table creation workflow
- Enterprise BI dashboard generation
- Financial reporting automation

### Key Deliverables
- Cache data population
- Field validation with errors
- PivotField wrapper with properties
- Complete XML serialization
- Roundtrip compatibility

---

## Phase 9: Add Comments API

**Proposal**: `add-comments-api`
**Category**: Word Processing
**Priority**: Medium-High (7/10)
**Estimated Effort**: 4-5 weeks

### Why Ninth
Existing comment infrastructure (~674 LOC) provides basic comment creation but lacks:
- High-level comment range marking API
- Comment status support (done/resolved)
- Comment deletion with cascade cleanup
- Filtering by author, date, status

### What It Enables
- Document review workflows
- Programmatic comment management
- Legal document automation

### Key Deliverables
- Paragraph.MarkCommentRange()
- Document.ApplyComment(), RemoveComment()
- Comment status (done/resolved) support
- Comment filtering methods
- CommentsExtendedPart support

---

## Phase 10: Add Track Changes API

**Proposal**: `add-track-changes-api`
**Category**: Word Processing
**Priority**: Critical (9/10)
**Estimated Effort**: 5.5 weeks

### Why Tenth
Existing revision element infrastructure (~946 LOC) provides low-level XML handling but no high-level API for:
- Enumerating all tracked changes
- Accepting/rejecting revisions
- Filtering by author, type, date

This blocks legal, editorial, and compliance workflows.

### What It Enables
- Contract review automation
- Editorial workflow processing
- Compliance audit trails
- Document finalization

### Key Deliverables
- Revision wrapper struct
- Accept()/Reject() methods
- Document.GetRevisions() iterator
- AcceptAllRevisions(), RejectAllRevisions()
- Move pair coordination

---

## Phase 11: Add Word Form Fields API

**Proposal**: `add-form-fields-api`
**Category**: Word Processing
**Priority**: High (7/10)
**Estimated Effort**: 3-4 weeks

### Why Eleventh
Basic form field elements exist but no manipulation API. Users cannot:
- Insert text fields, checkboxes, dropdowns
- Get/set form field values
- Protect documents for form filling

### What It Enables
- Business form creation
- Survey/questionnaire generation
- Data entry workflows
- Template creation with fillable fields
- Mail merge templates (Phase 12)

### Key Deliverables
- FormField wrapper class
- InsertTextField(), InsertCheckBox(), InsertDropDown()
- Document.GetFormFields(), GetFormFieldByName()
- Document protection for forms

---

## Phase 12: Add Mail Merge Engine

**Proposal**: `add-mail-merge-engine`
**Category**: Word Processing
**Priority**: High (8/10)
**Estimated Effort**: 4 weeks
**Depends On**: add-form-fields-api (partial - for field infrastructure understanding)

### Why Twelfth
Mail merge field elements exist but no execution engine. Users can create MERGEFIELD templates but cannot:
- Connect data sources
- Execute merge to produce documents
- Generate batches of personalized documents

### What It Enables
- Batch letter generation
- Certificate/badge creation
- Report automation
- Form filling from data

### Key Deliverables
- DataSource interface
- CSVDataSource, JSONDataSource, MapDataSource
- MailMerge engine with Execute(), ExecuteToDocuments()
- SimpleField and ComplexField support
- GREETINGLINE and SKIPIF field support

---

## Phase 13: Add SmartArt Support

**Proposal**: `add-smartart-support`
**Category**: DrawingML
**Priority**: Medium (6/10)
**Estimated Effort**: 2-3 weeks (Phase 1 only)
**Depends On**: add-extension-schema-support (for diagram schemas)

### Why Last
SmartArt roundtrip support requires:
- Extension schema loading (Phase 3)
- Diagram element generation from schemas
- Four interconnected diagram parts

This is foundational work for future SmartArt layout engine (separate proposal).

### What It Enables
- Safe roundtrip of documents with SmartArt
- No data loss when opening/saving
- Foundation for SmartArt creation (future)

### Key Deliverables
- Diagram element classes in `drawingml/diagram/`
- DiagramDataPart, DiagramLayoutDefinitionPart, DiagramStylePart, DiagramColorsPart
- Roundtrip preservation of SmartArt content

---

## Dependency Graph

```
                    ┌─────────────────────────────────────┐
                    │                                     │
                    ▼                                     │
┌──────────────────────────────┐                         │
│ enable-drawingml-pdf-rendering │◄────────────────────┐  │
└──────────────────────────────┘                       │  │
                    │                                   │  │
                    ▼                                   │  │
┌──────────────────────────────┐   ┌──────────────────┐│  │
│   enable-chart-rendering     │   │add-visual-comparison│  │
└──────────────────────────────┘   └──────────────────┘│  │
                                                       │  │
┌──────────────────────────────┐                       │  │
│  add-extension-schema-support │                       │  │
└──────────────────────────────┘                       │  │
                    │                                   │  │
                    ▼                                   │  │
┌──────────────────────────────┐                       │  │
│    add-smartart-support      │                       │  │
└──────────────────────────────┘                       │  │
                                                       │  │
┌──────────────────────────────┐                       │  │
│add-dotnet-cross-runtime-testing│                      │  │
└──────────────────────────────┘                       │  │
                    │                                   │  │
                    ▼                                   │  │
┌──────────────────────────────┐                       │  │
│   add-e2e-visual-testing     │───────────────────────┘  │
└──────────────────────────────┘                          │
                                                          │
┌──────────────────────────────┐                          │
│  add-row-column-operations   │                          │
└──────────────────────────────┘                          │
                    │                                      │
                    ▼                                      │
┌──────────────────────────────┐                          │
│    add-pivottable-api        │                          │
└──────────────────────────────┘                          │
                                                          │
┌──────────────────────────────┐                          │
│     add-comments-api         │ (independent)            │
└──────────────────────────────┘                          │
                                                          │
┌──────────────────────────────┐                          │
│   add-track-changes-api      │ (independent)            │
└──────────────────────────────┘                          │
                                                          │
┌──────────────────────────────┐                          │
│    add-form-fields-api       │                          │
└──────────────────────────────┘                          │
                    │                                      │
                    ▼                                      │
┌──────────────────────────────┐                          │
│   add-mail-merge-engine      │──────────────────────────┘
└──────────────────────────────┘
```

---

## Parallel Execution Opportunities

### Parallel Track A: PDF Infrastructure
1. enable-drawingml-pdf-rendering
2. enable-chart-rendering
3. add-visual-comparison

### Parallel Track B: Testing Infrastructure
1. add-dotnet-cross-runtime-testing
2. add-e2e-visual-testing

### Parallel Track C: Spreadsheet Features
1. add-row-column-operations
2. add-pivottable-api

### Parallel Track D: Word Processing Features
1. add-comments-api (independent)
2. add-track-changes-api (independent)
3. add-form-fields-api → add-mail-merge-engine

### Parallel Track E: Schema/DrawingML
1. add-extension-schema-support → add-smartart-support

**Maximum Parallelization**: Tracks A, B, C, D, E can begin simultaneously if resources permit.

---

## Summary by Priority

### Critical (Must Do First)
1. **enable-drawingml-pdf-rendering** - Unblocks all PDF rendering
2. **enable-chart-rendering** - Enables existing 3,543 LOC
3. **add-row-column-operations** - Tier 1 spreadsheet gap
4. **add-track-changes-api** - Blocks legal/editorial workflows

### High Priority
5. **add-extension-schema-support** - Modern Office compatibility
6. **add-dotnet-cross-runtime-testing** - API parity validation
7. **add-pivottable-api** - Enterprise BI requirement
8. **add-mail-merge-engine** - Document automation core

### Medium Priority
9. **add-comments-api** - Document review workflows
10. **add-form-fields-api** - Business form automation
11. **add-e2e-visual-testing** - Visual regression prevention
12. **add-visual-comparison** - PDF fidelity validation
13. **add-smartart-support** - Document compatibility

---

## Total Estimated Effort

| Category | Proposals | Estimated Weeks |
|----------|-----------|-----------------|
| PDF Infrastructure | 2 | 9-10 |
| Core Infrastructure | 1 | 3-4 |
| Testing | 3 | 9-12 |
| Spreadsheet | 2 | 11-12 |
| Word Processing | 4 | 16-18 |
| DrawingML | 1 | 2-3 |
| **Total** | **13** | **50-59 weeks** |

With maximum parallelization across tracks, timeline could compress to **20-25 weeks** with 3 parallel work streams.
