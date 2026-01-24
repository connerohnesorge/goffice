# goffice Proposals: Quick Reference Guide

## Proposal Locations

All proposals are located in `/spectr/changes/` with the following structure:

```
spectr/changes/
├── add-wordprocessing-advanced-features/     [49 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── wordprocessing-elements/
│           └── spec.md
├── add-spreadsheet-advanced-features/        [72 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── spreadsheet-cells/
│           └── spec.md
├── add-presentation-advanced-features/       [56 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── presentation-elements/
│           └── spec.md
├── enhance-drawingml-coverage/               [39 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── drawingml-core/
│           └── spec.md
├── enhance-pdf-rendering-fidelity/           [61 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── pdf-core/
│           └── spec.md
├── expand-test-coverage/                     [102 tasks]
│   ├── proposal.md
│   ├── tasks.md
│   └── specs/
│       └── (multiple)
└── add-validation-and-compliance/            [60 tasks]
    ├── proposal.md
    ├── tasks.md
    └── specs/
        └── validation/
            └── spec.md
```

---

## Quick Overview by Feature Area

### Word (WordprocessingML)

**Proposal:** `add-wordprocessing-advanced-features`

**Features:**
- Advanced text formatting (theme colors, spacing, effects)
- Table operations (merging, complex layouts)
- Paragraph formatting (outline levels, shading, borders)
- Section management (headers, footers per section)
- Comments & annotations
- Bookmarks & cross-references
- Hyperlinks
- Document protection
- Text boxes & frames
- Advanced footnotes/endnotes

**Quick Facts:**
- 49 tasks
- Affects: `wordprocessing/elements/`, `wordprocessing/document.go`
- Requires: Roundtrip tests for all features
- Time estimate: 3-4 weeks

---

### Excel (SpreadsheetML)

**Proposal:** `add-spreadsheet-advanced-features`

**Features:**
- Data validation (list, number, date, custom)
- Conditional formatting (colors, data bars, icons)
- Filtering & sorting
- Sparklines
- Named ranges
- Advanced charts (waterfall, funnel, sunburst)
- Chart formatting (axes, legends, trend lines)
- Slicers & timelines
- What-if analysis
- Array formulas
- Comments
- Hyperlinks
- Rich text
- Merged cells

**Quick Facts:**
- 72 tasks (largest feature set)
- Affects: `spreadsheet/sheet.go`, `spreadsheet/chart.go`
- Requires: Comprehensive cell/chart tests
- Time estimate: 4-5 weeks

---

### PowerPoint (PresentationML)

**Proposal:** `add-presentation-advanced-features`

**Features:**
- Master slides & layouts
- Custom slide dimensions
- Slide transitions
- Shape animations
- Hyperlinks & actions
- SmartArt
- Themes & color schemes
- Speaker notes
- Handouts
- Slide numbering
- Protection
- Sections
- OLE objects

**Quick Facts:**
- 56 tasks
- Affects: `presentation/document.go`, `presentation/elements/`
- Requires: Animation & master slide tests
- Time estimate: 3-4 weeks

---

### Drawing (DrawingML)

**Proposal:** `enhance-drawingml-coverage`

**Features:**
- 3D shapes (extrusion, bevel, materials)
- Shadows (outer, inner, perspective)
- Reflections & glows
- Gradient fills (linear, radial, path)
- Pattern fills
- Advanced strokes (dashes, caps, joins)
- Text rotation
- Connectors
- Picture effects
- Chart formatting

**Quick Facts:**
- 39 tasks
- Affects: `drawingml/`, `spreadsheet/chart.go`, `presentation/elements/`
- Requires: Effect combination tests
- Time estimate: 2-3 weeks
- Cross-cutting: Used by Word, Excel, PowerPoint

---

### PDF Rendering

**Proposal:** `enhance-pdf-rendering-fidelity`

**Features:**
- Advanced font rendering
- Bidirectional text
- Shape effects in PDF
- Gradient/pattern fills
- Chart rendering improvements
- Table merging
- Headers/footers per section
- Form field rendering
- Comments as annotations
- Watermarks
- Text rotation
- PDF links
- Font subsetting
- PDF/A accessibility

**Quick Facts:**
- 61 tasks
- Affects: `pdf/`, `spreadsheet/`, `wordprocessing/`, `presentation/` PDF files
- Requires: Visual regression tests
- Time estimate: 4-5 weeks

---

### Testing

**Proposal:** `expand-test-coverage`

**Coverage Areas:**
- Roundtrip tests (create→save→load→verify)
- Interoperability (Office 2007-365, LibreOffice, Google Docs)
- Edge cases & boundaries
- Performance benchmarks
- Fuzzing tests
- Visual regression (PDF)
- Version compatibility
- Security tests

**Quick Facts:**
- 102 test tasks
- Affects: All test files
- Target: >90% code coverage
- Time estimate: 4-5 weeks (parallel with features)

---

### Validation & Compliance

**Proposal:** `add-validation-and-compliance`

**Features:**
- ECMA-376 schema validation
- ISO/IEC 29500 compliance
- Element cardinality validation
- Type & enumeration validation
- Relationship validation
- Custom validation rules
- Document repair
- Error reporting
- All standard editions support

**Quick Facts:**
- 60 tasks
- Affects: `openxml/validation/`, all element files
- Requires: Comprehensive validation tests
- Time estimate: 3-4 weeks
- Priority: Foundation (do first)

---

## Implementation Workflow

### For Each Proposal:

1. **Read the proposal file**
   ```bash
   cat spectr/changes/[proposal-name]/proposal.md
   ```
   - Understand the "Why" and "What Changes"
   - Review the Impact section

2. **Review the tasks**
   ```bash
   cat spectr/changes/[proposal-name]/tasks.md
   ```
   - Check task breakdown
   - Estimate effort for each task
   - Identify dependencies

3. **Review the spec deltas**
   ```bash
   cat spectr/changes/[proposal-name]/specs/[capability]/spec.md
   ```
   - Understand requirements
   - Review scenarios (WHEN-THEN format)
   - Identify test cases

4. **Start implementation**
   - Create feature branch: `feature/[proposal-name]`
   - Follow task checklist
   - Implement features following spec requirements
   - Write tests as you code

5. **Submit for review**
   - Update task status (mark completed items)
   - Create PR with reference to proposal
   - Include roundtrip test results
   - Get code review approval

---

## Feature Dependencies

```
Validation & Compliance ◄── All features depend on this
        ↑
        ├─→ Wordprocessing Advanced ◄─── DrawingML, PDF
        ├─→ Spreadsheet Advanced    ◄─── DrawingML, PDF
        ├─→ Presentation Advanced   ◄─── DrawingML, PDF
        └─→ Test Coverage (parallel)
```

**Recommendation:** Implement in order:
1. Validation & Compliance (foundation)
2. DrawingML (cross-cutting)
3. Feature sets (Word, Excel, PPT) in parallel
4. PDF Rendering
5. Test Coverage (throughout)

---

## Key Statistics

| Proposal | Tasks | Time | Priority | Status |
|----------|-------|------|----------|--------|
| Wordprocessing | 49 | 3-4w | High | Awaiting approval |
| Spreadsheet | 72 | 4-5w | High | Awaiting approval |
| Presentation | 56 | 3-4w | Medium | Awaiting approval |
| DrawingML | 39 | 2-3w | Medium | Awaiting approval |
| PDF | 61 | 4-5w | High | Awaiting approval |
| Testing | 102 | 4-5w | High | Awaiting approval |
| Validation | 60 | 3-4w | Critical | Awaiting approval |
| **TOTAL** | **439** | **23-30w** | - | - |

---

## Approval & Tracking

### Checking Approval Status:
```bash
# List all proposals
cat spectr/changes/

# Read any proposal
cat spectr/changes/[proposal-name]/proposal.md
```

### Tracking Progress:
Each proposal has a `tasks.md` file with checkbox format:
```markdown
- [ ] 1.1 Task description
- [x] 1.2 Completed task
```

### Moving to Implementation:
1. Proposal approved by team ✅
2. Create feature branch
3. Mark tasks as in-progress
4. Implement features
5. Write tests
6. Update task status
7. Submit PR
8. Code review
9. Merge & update tasks

---

## Contact & Questions

For questions about:
- **Specific proposals:** Check the proposal.md file
- **Technical design:** Read the design.md file (if exists)
- **Implementation details:** Review the spec delta files
- **Approval process:** Contact team lead

---

## Additional Resources

- **Spec Instructions:** See `/spectr/AGENTS.md`
- **Project Context:** See `/spectr/project.md`
- **Existing Specs:** Browse `/spectr/specs/`
- **Summary Document:** See `/PROPOSALS_SUMMARY.md`

---

**Created:** January 23, 2026
**Status:** All proposals ready for review and approval
