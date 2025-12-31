# goffice Implementation Timeline

This document outlines the chronological implementation order for all Spectr change proposals, organized by dependency requirements and optimized for parallel development where possible.

## Executive Summary

**Total Proposals:** 8 major feature proposals (graded and validated)
**Total Estimated Duration:** 18-22 weeks with parallelization
**Critical Path:** enable-drawingml-pdf-rendering → enable-chart-rendering → add-smartart-support

---

## Dependency Graph

```
                      ┌─────────────────────────────┐
                      │  add-extension-schema-support│
                      │         (3-4 weeks)          │
                      └─────────────────────────────┘
                                    │
                                    │ enables
                                    ▼
        ┌───────────────────────────────────────────────────┐
        │                 Office 2010-2025                   │
        │            Element Types Available                 │
        └───────────────────────────────────────────────────┘

                      ┌─────────────────────────────┐
                      │enable-drawingml-pdf-rendering│
                      │         (6 weeks)            │
                      └─────────────────────────────┘
                                    │
                                    │ provides Page interface
                                    ▼
                      ┌─────────────────────────────┐
                      │   enable-chart-rendering     │
                      │         (4 weeks)            │
                      └─────────────────────────────┘
                                    │
                                    │ uses chart renderer
                                    ▼
                      ┌─────────────────────────────┐
                      │   add-smartart-support       │
                      │         (3 weeks)            │
                      │    (Phase 2: +6-12 weeks)    │
                      └─────────────────────────────┘

        ┌─────────────────────────────────────────────┐
        │           INDEPENDENT PROPOSALS              │
        │   (Can run in parallel with PDF track)       │
        ├─────────────────────────────────────────────┤
        │ • add-track-changes-api      (5.5 weeks)    │
        │ • add-row-column-operations  (5 weeks)      │
        │ • add-form-fields-api        (4 weeks)      │
        │ • add-mail-merge-engine      (4 weeks)      │
        └─────────────────────────────────────────────┘
```

---

## Phase 1: Foundation (Weeks 1-4)

### Week 1-4: Infrastructure Track (Critical Path)

#### 1. enable-drawingml-pdf-rendering (START)
**Priority:** CRITICAL - Blocks chart and SmartArt rendering
**Duration:** 6 weeks (Weeks 1-6)
**Dependencies:** None (foundational)
**Description:** Creates the Page interface abstraction for PDF drawing primitives, enabling all 12 WIP renderer files.

**Key Deliverables:**
- Page interface with all 25+ methods
- pdfcpu adapter implementation
- Coordinate conversion (EMU → PDF points)
- Graphics state management (save/restore)
- Re-enable all `.wip` → `.go` files

#### 2. add-extension-schema-support (START - Parallel)
**Priority:** HIGH - Enables Office 2010-2025 features
**Duration:** 3-4 weeks (Weeks 1-4)
**Dependencies:** None (foundational)
**Description:** Updates code generators to load extension schemas, adding 200-300 new element types.

**Key Deliverables:**
- Generator schema loading expansion
- AlternateContent/Choice/Fallback support
- Unknown element preservation
- Version-aware validation
- 200-300 new generated element classes

---

## Phase 2: Parallel Development (Weeks 3-8)

### Word Processing Track (Weeks 3-8)

#### 3. add-track-changes-api (START Week 3)
**Priority:** HIGH
**Duration:** 5.5 weeks (Weeks 3-8.5)
**Dependencies:** None
**Description:** Full revision tracking for Word documents.

**Key Deliverables:**
- RevisionCollection with filtering
- Accept/Reject operations
- Move pair coordination
- Format change handling
- Phase 1: 6 revision types (85% coverage)

#### 4. add-form-fields-api (START Week 3)
**Priority:** MEDIUM
**Duration:** 4 weeks (Weeks 3-7)
**Dependencies:** None
**Description:** Legacy form field support (text, checkbox, dropdown).

**Key Deliverables:**
- TextField, CheckBox, DropDown wrappers
- Value synchronization (ffData ↔ run text)
- Iterator for field enumeration
- FormFieldName, DefaultDropDownIndex support

#### 5. add-mail-merge-engine (START Week 7)
**Priority:** MEDIUM
**Duration:** 4 weeks (Weeks 7-11)
**Dependencies:** Soft dependency on add-form-fields-api understanding
**Description:** Mail merge field substitution engine.

**Key Deliverables:**
- SimpleField and ComplexField support
- DataSource abstraction (CSV, JSON, custom)
- MERGEFIELD, GREETINGLINE, SKIPIF support
- Settings.xml metadata integration

### Spreadsheet Track (Weeks 3-8)

#### 6. add-row-column-operations (START Week 3)
**Priority:** HIGH
**Duration:** 5 weeks (Weeks 3-8)
**Dependencies:** None
**Description:** Row/column insert/delete with formula updates.

**Key Deliverables:**
- InsertRows, DeleteRows, InsertColumns, DeleteColumns
- Formula reference shifting
- Shared formula and array formula handling
- Named range updates
- 25+ element type updates (merged cells, charts, etc.)

---

## Phase 3: PDF Rendering Completion (Weeks 6-13)

#### 7. enable-chart-rendering (START Week 7)
**Priority:** HIGH
**Duration:** 4 weeks (Weeks 7-11)
**Dependencies:** enable-drawingml-pdf-rendering MUST be complete
**Description:** Re-enables chart rendering in PDF output.

**Key Deliverables:**
- Bar, line, pie, scatter, area, radar charts
- Axis rendering with labels
- Legend rendering
- Data series styling
- Integration with RenderingContext

#### 8. add-smartart-support Phase 1 (START Week 11)
**Priority:** MEDIUM
**Duration:** 3 weeks (Weeks 11-14)
**Dependencies:** 
  - enable-drawingml-pdf-rendering (for PDF placeholders)
  - drawingml-core infrastructure
**Description:** SmartArt roundtrip support (no layout engine).

**Key Deliverables:**
- Schema generation from 5 diagram schema files
- DiagramDataPart, DiagramLayoutDefinitionPart
- DiagramStylePart, DiagramColorsPart
- PDF placeholder rendering (bounding box + label)
- Roundtrip preservation for all SmartArt types

---

## Timeline Summary

| Week | Proposal | Track | Status |
|------|----------|-------|--------|
| 1-6 | enable-drawingml-pdf-rendering | PDF/Drawing | CRITICAL PATH |
| 1-4 | add-extension-schema-support | Infrastructure | Parallel |
| 3-8.5 | add-track-changes-api | Word | Parallel |
| 3-7 | add-form-fields-api | Word | Parallel |
| 3-8 | add-row-column-operations | Excel | Parallel |
| 7-11 | add-mail-merge-engine | Word | Sequential |
| 7-11 | enable-chart-rendering | PDF/Drawing | Depends on Week 6 |
| 11-14 | add-smartart-support | PDF/Drawing | Depends on Week 11 |

---

## Resource Allocation Recommendation

### Optimal Team Configuration (3 developers)

**Developer A: PDF/Drawing Track (Weeks 1-14)**
- Week 1-6: enable-drawingml-pdf-rendering
- Week 7-11: enable-chart-rendering
- Week 11-14: add-smartart-support

**Developer B: Word Track (Weeks 3-11)**
- Week 3-8.5: add-track-changes-api
- Week 7-11: add-mail-merge-engine (overlap with track changes completion)

**Developer C: Infrastructure + Excel (Weeks 1-8)**
- Week 1-4: add-extension-schema-support
- Week 3-8: add-row-column-operations
- Week 3-7: add-form-fields-api (parallel with row/column ops)

### Single Developer Timeline

If single developer: **22-24 weeks total**
1. enable-drawingml-pdf-rendering (6 weeks)
2. add-extension-schema-support (4 weeks)
3. add-track-changes-api (5.5 weeks)
4. add-row-column-operations (5 weeks) - overlap with track changes
5. add-form-fields-api (4 weeks)
6. enable-chart-rendering (4 weeks)
7. add-mail-merge-engine (4 weeks)
8. add-smartart-support (3 weeks)

---

## Critical Path Analysis

**Longest Path:** enable-drawingml-pdf-rendering → enable-chart-rendering → add-smartart-support
**Duration:** 6 + 4 + 3 = 13 weeks

**Parallelizable Work:**
- add-extension-schema-support (independent)
- add-track-changes-api (independent)
- add-row-column-operations (independent)
- add-form-fields-api (independent)
- add-mail-merge-engine (soft dependency on form fields)

---

## Risk Factors

### High Risk
1. **enable-drawingml-pdf-rendering**: Foundation for all PDF features - delays cascade
2. **Chart rendering complexity**: May uncover issues in WIP files

### Medium Risk
1. **Formula shifting edge cases**: Array formulas, shared formulas
2. **Track changes move pairs**: Cross-part coordination complexity
3. **Mail merge field types**: SimpleField vs ComplexField handling

### Low Risk
1. **Form fields**: Well-understood, limited scope
2. **Extension schemas**: Generator changes only, low runtime impact
3. **SmartArt Phase 1**: Roundtrip only, no layout engine

---

## Success Milestones

| Week | Milestone | Verification |
|------|-----------|--------------|
| 4 | Extension schemas generating | 200+ new element types |
| 6 | PDF Page interface complete | All 12 WIP files compile |
| 8 | Track changes API functional | Accept/Reject roundtrips |
| 8 | Row/column operations working | Formulas shift correctly |
| 11 | Chart rendering enabled | PDF includes charts |
| 14 | SmartArt roundtrips | All SmartArt types preserved |

---

## Pre-Existing Proposals (Not Graded in This Session)

The following proposals were found in spectr/changes/ but were not part of this grading session:

1. **add-pivottable-api** - Excel PivotTable support
2. **add-visual-comparison** - Visual regression testing
3. **add-comments-api** - Document comments
4. **add-dotnet-cross-runtime-testing** - .NET test infrastructure
5. **add-e2e-visual-testing** - End-to-end visual tests

These should be scheduled after the core 8 proposals are complete.

---

## Version History

| Date | Change |
|------|--------|
| 2025-12-30 | Initial timeline created after grading all 8 proposals |

