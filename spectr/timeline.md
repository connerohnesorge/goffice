# Spectr Change Proposals Timeline

This document lists all change proposals in dependency order. Proposals in lower tiers
should be completed before proposals in higher tiers that depend on them.

---

## Archived (Completed)

These proposals have been completed and archived. They form the foundation for all active proposals.

| ID | Title | Date | Notes |
|----|-------|------|-------|
| `2025-12-20-add-go-word-sdk` | Word SDK Foundation | 2025-12-20 | Original Word SDK - foundational |
| `2025-12-23-add-go-drawingml` | DrawingML Support | 2025-12-23 | Shared drawing specification |
| `2025-12-23-add-go-presentation-sdk` | PowerPoint SDK | 2025-12-23 | Presentation support |
| `2025-12-23-add-go-spreadsheet-sdk` | Excel SDK | 2025-12-23 | Spreadsheet support |
| `2025-12-23-add-pdf-rendering` | PDF Rendering Foundation | 2025-12-23 | PDF generation core |
| `2025-12-31-add-form-fields-api` | Form Fields API | 2025-12-31 | Word form fields |
| `2025-12-31-add-row-column-operations` | Row/Column Operations | 2025-12-31 | Excel row/column manipulation |
| `2025-12-31-add-track-changes-api` | Track Changes API | 2025-12-31 | Word revision tracking |
| `2025-12-31-enable-chart-rendering` | Chart Rendering | 2025-12-31 | Chart to PDF rendering |
| `2025-12-31-enable-drawingml-pdf-rendering` | DrawingML PDF Rendering | 2025-12-31 | DrawingML to PDF |
| `2025-12-31-add-extension-schema-support` | Extension Schema Support | 2025-12-31 | Office 2010-2025 elements, AlternateContent |

---

## Active Proposals

### Tier 1: Testing Infrastructure

No dependencies on other active proposals. These can be worked on in parallel and provide
infrastructure that benefits other proposals.

| ID | Title | Impact | Est. Time | Notes |
|----|-------|--------|-----------|-------|
| `add-visual-comparison` | Automated Visual Comparison | Medium | 2-3 weeks | PDF fidelity testing infrastructure. MVP for Word, extends to Excel/PowerPoint. |
| `add-e2e-visual-testing` | E2E Visual Testing Framework | Medium | 3-4 weeks | PPTX generation comparison with .NET SDK. Nix-powered test framework. |
| `add-dotnet-cross-runtime-testing` | .NET Cross-Runtime Testing | Medium | 4+ weeks | Three-level comparison (XML → Binary → Visual). CI/CD integration. |

---

### Tier 2: Core Features (Standalone)

No dependencies on other active proposals. Each provides independent functionality.

| ID | Title | Impact | Est. Time | Notes |
|----|-------|--------|-----------|-------|
| `add-formula-evaluation` | Formula Evaluation Engine | High | 8 weeks | Spreadsheet formula evaluation. 60+ Excel functions. Parser, AST, dependency tracking. |
| `add-comments-api` | Word Comments API | Medium | 80-105 hours | Enhance existing 674 LOC. Comment ranges, status, deletion, filtering, threading. |
| `add-pivottable-api` | PivotTable API Enhancement | High | 150-200 hours | Enhance existing 560 LOC. Cache population, field validation, PivotField wrapper. |
| `add-presentation-table-api` | PowerPoint Table API | Medium | 3-4 weeks | DrawingML table elements. High-level table creation/manipulation. PDF rendering. |
| `add-mail-merge-engine` | Mail Merge Engine | High | 4 weeks | DataSource interface, field finding/replacement, OOXML metadata integration. |
| `enhance-shape-grouping` | Shape Grouping Support | Medium | 2-3 weeks | GroupShape API, transform composition, PDF group rendering. **ENABLES**: add-connector-routing |

---

### Tier 3: Features with Soft Dependencies

These proposals benefit from or are easier to implement after certain Tier 2 proposals.

| ID | Title | Impact | Est. Time | Soft Dependency | Notes |
|----|-------|--------|-----------|-----------------|-------|
| `add-smartart-support` | SmartArt/Diagram Support (Phase 1) | High | 2-3 weeks | — | Schema generation, diagram parts, roundtrip. Foundation for Phase 2. |
| `add-connector-routing` | Connector Shape Routing | Medium | 3-4 weeks | enhance-shape-grouping | Connection points, routing algorithms (straight/elbow/curved), PDF rendering. |

---

### Tier 4: Advanced Features (Hard Dependencies)

These proposals explicitly require completion of earlier proposals.

| ID | Title | Impact | Est. Time | **Requires** | Notes |
|----|-------|--------|-----------|--------------|-------|
| `implement-smartart-layouts` | SmartArt Layout Engine (Phase 2) | High | 8-10 weeks | **add-smartart-support** | 5 core layout algorithms: Hierarchy, List, Cycle, Pyramid, Matrix. |

---

## Dependency Graph

```
                     ┌─────────────────────────────────────┐
                     │         ARCHIVED (COMPLETE)         │
                     │  Word SDK → DrawingML → PDF Core    │
                     │  Presentation SDK → Spreadsheet SDK │
                     │  Form Fields, Track Changes, Charts │
                     │  Extension Schema Support (2010-25) │
                     └─────────────────────────────────────┘
                                       │
           ┌───────────────────────────┼───────────────────────────┐
           │                           │                           │
           ▼                           ▼                           ▼
┌─────────────────────┐   ┌─────────────────────┐   ┌─────────────────────┐
│ TIER 1: TESTING     │   │ TIER 1: TESTING     │   │ TIER 1: TESTING     │
├─────────────────────┤   ├─────────────────────┤   ├─────────────────────┤
│ add-visual-         │   │ add-e2e-visual-     │   │ add-dotnet-cross-   │
│ comparison          │   │ testing             │   │ runtime-testing     │
└─────────────────────┘   └─────────────────────┘   └─────────────────────┘
           │                           │                           │
           └───────────────────────────┼───────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                        TIER 2: CORE FEATURES                            │
├─────────────────────────────────────────────────────────────────────────┤
│ add-formula-evaluation     │ add-comments-api     │ add-pivottable-api  │
│ add-presentation-table-api │ add-mail-merge-engine│ enhance-shape-      │
│                            │                      │ grouping ──────────►│
└─────────────────────────────────────────────────────────────────────────┘
                                       │                           │
                                       │                           │
           ┌───────────────────────────┘                           │
           │                                                       │
           ▼                                                       ▼
┌─────────────────────┐                               ┌─────────────────────┐
│ TIER 3: DEPENDENT   │                               │ TIER 3: DEPENDENT   │
├─────────────────────┤                               ├─────────────────────┤
│ add-smartart-       │                               │ add-connector-      │
│ support (Phase 1)   │                               │ routing             │
│                     │                               │                     │
│                     │                               │ ★ Benefits from     │
│                     │                               │   shape grouping    │
└─────────────────────┘                               └─────────────────────┘
           │
           │ REQUIRED
           ▼
┌─────────────────────┐
│ TIER 4: ADVANCED    │
├─────────────────────┤
│ implement-smartart- │
│ layouts (Phase 2)   │
│                     │
│ ★ REQUIRES Phase 1  │
└─────────────────────┘
```

---

## Recommended Implementation Order

Based on dependencies and strategic value:

### Phase A: Testing Infrastructure
1. **add-visual-comparison** - Enables automated PDF fidelity testing
2. **add-e2e-visual-testing** / **add-dotnet-cross-runtime-testing** - Cross-runtime validation

### Phase B: Core Functionality (Parallel)
Work on these in parallel based on priority:
- **add-formula-evaluation** - High-value for spreadsheet automation
- **add-mail-merge-engine** - High-value for document automation
- **add-pivottable-api** - High-value for BI/reporting
- **enhance-shape-grouping** - Enables connector routing

### Phase C: Presentation Features
- **add-presentation-table-api** - Complete table support for PowerPoint
- **add-connector-routing** - After enhance-shape-grouping

### Phase D: SmartArt (Sequential)
1. **add-smartart-support** (Phase 1) - Schema generation, roundtrip
2. **implement-smartart-layouts** (Phase 2) - Layout engine (after Phase 1)

### Phase E: Polish
- **add-comments-api** - Enhance existing comments infrastructure

---

## Notes

- **Parallel work**: Proposals in the same tier can generally be worked on in parallel
- **Hard dependencies**: Tier 4 proposals MUST wait for their requirements
- **Soft dependencies**: Tier 3 proposals benefit from but don't strictly require earlier proposals
- **Testing infrastructure**: Can be developed alongside feature work
- **Extension schemas**: Already implemented (archived 2025-12-31) - Office 2010-2025 elements available

---

*Generated: 2025-12-31*
