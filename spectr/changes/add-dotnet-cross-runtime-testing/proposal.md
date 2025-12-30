# Change: Add .NET Cross-Runtime Testing Infrastructure with Nix Integration

## Why

The goffice project aims to provide feature parity with Microsoft's Open-XML-SDK for .NET. Currently, we have:
- A complete Go implementation of Office document manipulation (Word, Excel, PowerPoint)
- The official .NET SDK as a git submodule at `Open-XML-SDK/`
- Basic PDF fidelity testing infrastructure (`pdf/comparison/`)
- Separate test suites for Go and .NET with NO integration

**The Problem**: We have no automated way to verify that our Go implementation produces documents identical to the .NET SDK. This creates risks:
- Silent API divergence (methods work differently)
- Incompatible document generation (documents that don't match)
- Missing validation against the reference implementation
- No regression detection when .NET SDK updates
- No confidence in "feature parity" claims

**The Solution**: Create a comprehensive Nix-powered cross-runtime testing system that:
1. Generates documents using BOTH our Go API AND the .NET SDK from identical test scenarios
2. Compares outputs at multiple levels (XML structure, binary content, visual rendering)
3. Runs automatically in CI/CD with reproducible Nix environment
4. Provides visual diff reports for manual review when differences detected
5. Supports all three document types (Word, Excel, PowerPoint)

This system will:
- **Validate API parity**: Prove our Go API matches .NET SDK behavior
- **Catch regressions**: Detect when changes break compatibility
- **Guide development**: Show exactly where implementations diverge
- **Build confidence**: Provide evidence of feature parity with reference implementation

## What Changes

### 1. New E2E Test Infrastructure (`tests/e2e/`)
- **Location**: New top-level `tests/e2e/` directory (user-requested)
- **Test matrix framework**: Go + C# document generation with comparison
- **Test scenario definitions**: JSON/YAML definitions of documents to create
- **Visual assertion engine**: Extended `pdf/comparison` for cross-runtime use
- **Diff reporting**: HTML/PDF reports showing Go vs. .NET differences
- **Fixture management**: Automated baseline generation and update workflow

### 2. Nix Environment Extensions
- **Add .NET SDK 9**: Include `dotnet-sdk_9` in development shell
- **Cross-language scripts**:
  - `test-go` - Run Go tests only
  - `test-dotnet` - Run .NET SDK tests only
  - `test-e2e` - Run cross-runtime comparison tests
  - `test-all` - Complete test suite (Go + .NET + E2E)
  - `update-baselines` - Regenerate reference documents from .NET SDK
- **Multi-runtime builds**: Support building both Go binaries and .NET projects
- **Reproducibility**: Pin .NET SDK version in flake.lock for consistency

### 3. Cross-Runtime Test Harness
- **Scenario executor**: Runs identical test scenarios in both Go and C#
- **Document generator bridges**:
  - Go bridge: Uses goffice API to create documents
  - C# bridge: Uses Open-XML-SDK to create documents
- **Comparison pipeline**:
  - **Level 1**: XML structure comparison (elements, attributes, namespaces)
  - **Level 2**: Binary content comparison (relationships, content types, images)
  - **Level 3**: Visual rendering comparison (PDF output pixel diff)
- **Result aggregation**: Unified report across all test scenarios

### 4. Test Scenario System
- **Scenario definitions**: Declarative test descriptions (YAML/JSON)
- **Example scenarios**:
  - `basic-paragraph.yaml`: Simple text document
  - `formatted-text.yaml`: Bold, italic, colors, fonts
  - `tables.yaml`: Table with borders, shading, merged cells
  - `spreadsheet-formulas.yaml`: Excel with formulas and formatting
  - `presentation-shapes.yaml`: PowerPoint with shapes and text boxes
- **Scenario categories**:
  - Basic functionality (smoke tests)
  - Complex documents (real-world cases)
  - Edge cases (boundaries, limits, unusual combinations)
  - Regression tests (historical bugs)

### 5. Visual Comparison Extensions
- **Extend `pdf/comparison`** to support:
  - Side-by-side document rendering (Go vs .NET)
  - Multi-page diff reports with thumbnails
  - Diff statistics (% match, pixel deltas, region highlighting)
  - Interactive HTML reports (zoom, toggle overlay)
- **Support Office-native rendering** (optional):
  - LibreOffice headless conversion as fallback
  - Direct OOXML rendering via system Office (if available)

### 6. CI/CD Integration
- **GitHub Actions workflow** (or equivalent):
  - Runs on PR and main branch commits
  - Executes full E2E test suite
  - Uploads diff artifacts for failed tests
  - Comments on PRs with comparison results
- **Baseline management**:
  - Store baseline documents in git (or LFS for large files)
  - Automated baseline update PRs when .NET SDK updates
  - Approval workflow for baseline changes

## Impact

### Affected Specs
- **NEW: `e2e-testing`** - Cross-runtime end-to-end testing framework
- **NEW: `nix-environment`** - Multi-language Nix development environment
- **MODIFIED: `pdf-testing`** - Enhanced visual comparison for cross-runtime use

### Affected Code
- **NEW: `tests/e2e/`** - Complete E2E test infrastructure
  - `tests/e2e/framework/` - Test execution framework (Go)
  - `tests/e2e/bridges/` - Language-specific document generators
    - `bridges/go/` - Go document generation bridge
    - `bridges/csharp/` - C# document generation bridge (calls Open-XML-SDK)
  - `tests/e2e/scenarios/` - Test scenario definitions (YAML/JSON)
  - `tests/e2e/comparison/` - Extended comparison logic
  - `tests/e2e/reports/` - HTML/PDF diff report generation
- **MODIFIED: `flake.nix`** - Add .NET SDK and E2E test scripts
- **MODIFIED: `pdf/comparison/`** - Extend for cross-runtime use (generalize APIs)
- **NEW: `tests/e2e/README.md`** - Complete E2E test documentation

### Breaking Changes
**NONE** - This is purely additive. Existing tests and APIs unchanged.

## Key Design Decisions

### 1. Use Nix for Multi-Runtime Environment
**Decision**: Use Nix flake to provide reproducible Go + .NET environment

**Rationale**:
- Project already uses Nix for Go development
- Nix provides reproducible builds across machines and CI
- .NET SDK available in nixpkgs (`dotnet-sdk_9`)
- Single `nix develop` command gives complete environment
- Pin .NET version in flake.lock for consistency

**Alternatives considered**:
- Docker: More complex, heavier, less integrated with direnv
- Manual installation: Not reproducible, breaks on different machines
- Dev containers: VS Code-specific, doesn't help CLI/CI

### 2. Place E2E Tests in `tests/e2e/` Directory
**Decision**: Top-level `tests/e2e/` directory (user requirement)

**Rationale**:
- User explicitly requested "tests/e2e/**/ dir"
- Separates cross-runtime tests from package-level unit tests
- Clear organizational boundary (E2E vs unit vs integration)
- Follows common convention (Playwright, Cypress use `tests/` or `e2e/`)

**Structure**:
```
tests/e2e/
├── README.md              # Documentation
├── framework/             # Test execution framework (Go)
├── bridges/
│   ├── go/                # Go document generator
│   └── csharp/            # C# document generator
├── scenarios/             # Test definitions (YAML/JSON)
│   ├── wordprocessing/
│   ├── spreadsheet/
│   └── presentation/
├── baselines/             # Reference documents from .NET SDK
├── comparison/            # Comparison logic
├── reports/               # Generated diff reports
└── Makefile or flake.nix  # Test runner

```

### 3. Three-Level Comparison Strategy
**Decision**: XML → Binary → Visual pyramid approach

**Level 1: XML Structure Comparison** (Fast, Precise)
- Parse XML from both Go and .NET documents
- Compare element trees, attributes, namespaces
- Ignore formatting whitespace, attribute order
- Report: Missing elements, attribute mismatches, content differences
- **Goal**: Catch API bugs (wrong elements, missing properties)

**Level 2: Binary Content Comparison** (Medium Speed)
- Extract ZIP entries from both documents
- Compare relationships, content types, part URIs
- Byte-level comparison for binary parts (images, embedded objects)
- Report: Missing parts, relationship errors, binary mismatches
- **Goal**: Catch packaging bugs (wrong relationships, missing parts)

**Level 3: Visual Rendering Comparison** (Slow, High-Level)
- Render both documents to PDF
- Convert PDF pages to PNG (via Ghostscript)
- Pixel-level diff with configurable tolerance
- Generate annotated diff images
- Report: Visual differences, rendering errors, layout issues
- **Goal**: Catch semantic bugs (document looks wrong in Office)

**Rationale**:
- Pyramid approach: Fast tests fail early, expensive tests run only if needed
- Multiple perspectives: Different bugs caught at different levels
- Debugging aid: Specific level tells you where to look
- Performance: Most bugs caught in Level 1/2 without rendering

**Alternatives**:
- Visual-only: Too slow, misses structural bugs
- XML-only: Misses visual issues, incomplete validation
- Combined: This approach gives best of all worlds

### 4. Use Test Scenarios (Declarative)
**Decision**: Define tests as YAML/JSON scenarios, not hardcoded

**Example Scenario**:
```yaml
name: basic-paragraph
description: Single paragraph with plain text
document_type: wordprocessing
expected_file: baselines/wordprocessing/basic-paragraph.docx

operations:
  - action: create_document
    type: document

  - action: add_paragraph
    text: "Hello, World!"

  - action: save
    path: output/go/basic-paragraph.docx
```

**Rationale**:
- **Language-agnostic**: Same scenario runs in Go and C#
- **Readable**: Non-programmers can understand test intent
- **Versionable**: Scenarios tracked in git, easy to review
- **Extensible**: Add new operations without code changes
- **Reusable**: Combine scenarios into suites

**Bridges interpret scenarios**:
- **Go bridge**: Translates to goffice API calls
- **C# bridge**: Translates to Open-XML-SDK API calls
- Both produce documents from identical logical operations

### 5. Visual Comparison via Ghostscript (Already Available)
**Decision**: Reuse existing `pdf/comparison` infrastructure

**Current state**:
- `pdf/comparison/convert.go` - PDF→PNG via Ghostscript
- `pdf/comparison/diff.go` - Pixel-level comparison
- `pdf/comparison/annotate.go` - Diff visualization
- **Already working** for PDF fidelity tests

**Extension needed**:
- Generalize for cross-runtime use (not just PDF)
- Support Office documents (convert to PDF first)
- Multi-page diff reports (not just single page)
- HTML report generation (not just PNG diffs)

**Rationale**:
- Don't reinvent the wheel
- Proven approach (already in use)
- Ghostscript widely available via Nix (`pkgs.ghostscript`)

### 6. Baseline Management Strategy
**Decision**: Store .NET-generated baselines in git, update via script

**Workflow**:
1. **Initial baseline generation**:
   - Run `test-e2e --generate-baselines`
   - Uses .NET SDK to create reference documents
   - Stores in `tests/e2e/baselines/`
   - Commit to git

2. **Test execution**:
   - Generate document with Go
   - Compare against baseline (stored .NET document)
   - Report differences

3. **Baseline updates** (when .NET SDK changes):
   - Run `update-baselines` script
   - Regenerate all baselines with latest .NET SDK
   - Review diffs (did .NET behavior change?)
   - Create PR with baseline updates
   - Approve if changes expected

**Rationale**:
- **Deterministic**: Same baseline for all developers
- **Reviewable**: Baseline changes visible in git diffs
- **Traceable**: Know when and why baselines changed
- **Fast**: Don't re-run .NET for every test execution

**Alternatives**:
- Generate on-the-fly: Slower, requires .NET SDK always available
- External storage: Complexity, versioning issues
- No baselines: Can't detect regressions

## Implementation Scope

### Phase 1: Foundation (MVP for Word documents)
- Extend flake.nix with .NET SDK 9
- Create `tests/e2e/` directory structure
- Implement test framework core (scenario parsing, execution)
- Build Go bridge (goffice wordprocessing API)
- Build C# bridge (Open-XML-SDK wordprocessing API)
- Create 5-10 basic Word scenarios (paragraphs, runs, tables, formatting)
- Level 1 comparison only (XML structure)
- Command-line test runner
- **Deliverable**: Can create simple Word docs in Go + C#, compare structure

### Phase 2: Complete Comparison Pipeline
- Implement Level 2 comparison (binary content)
- Implement Level 3 comparison (visual rendering)
- Extend `pdf/comparison` for multi-page, cross-runtime use
- HTML diff report generation
- Baseline management scripts
- **Deliverable**: Full 3-level comparison with visual diffs

### Phase 3: Expand Coverage
- Add Excel scenarios (spreadsheet bridge, formulas, charts)
- Add PowerPoint scenarios (presentation bridge, slides, shapes)
- Increase scenario count to 30-50 across all document types
- Add edge cases and regression tests
- **Deliverable**: Comprehensive cross-runtime coverage

### Phase 4: CI/CD Integration
- GitHub Actions workflow for E2E tests
- Artifact upload for failed test diffs
- PR comment integration (show diffs inline)
- Automated baseline update workflow
- **Deliverable**: Fully automated testing in CI

## Success Metrics

1. **Coverage**: 50+ test scenarios across Word/Excel/PowerPoint
2. **Pass Rate**: >95% XML structure match between Go and .NET
3. **Visual Fidelity**: >99% pixel match for rendering (Level 3)
4. **Performance**: E2E test suite completes in <10 minutes
5. **Developer Experience**: `nix develop` + `test-e2e` just works
6. **Regression Detection**: Catch API divergence within 1 commit
7. **CI Integration**: Runs on every PR, reports results automatically

## Open Questions (To Be Resolved During Implementation)

1. **Scenario format**: YAML, JSON, or custom DSL? (Recommend YAML)
2. **Baseline size**: Store in git or git-lfs? (Depends on document size)
3. **Test data generation**: Manual or property-based? (Start manual)
4. **Office version targeting**: Office 2016, 2019, 365? (Start with 2016)
5. **Parallel execution**: Run tests in parallel? (Yes, for speed)
6. **Tolerance configuration**: Per-scenario tolerances? (Yes, some documents naturally differ)

## Dependencies

- **.NET SDK 9**: Available in nixpkgs as `dotnet-sdk_9`
- **Ghostscript**: Available in nixpkgs as `ghostscript`
- **Open-XML-SDK submodule**: Already present at `Open-XML-SDK/`
- **pdf/comparison package**: Already exists and working

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| .NET SDK API mismatch | Can't create equivalent documents | Research .NET SDK patterns first (DONE) |
| Visual comparison flakiness | False positives on diffs | Configurable tolerances, baseline management |
| Test suite too slow | Developers skip tests | Parallel execution, pyramid approach (fast XML first) |
| Baseline drift | Baselines out of date | Automated update workflow, CI checks |
| Nix complexity | Hard to setup | Good documentation, direnv auto-activation |
| Large diff reports | Storage/git bloat | git-lfs for large files, HTML over binary |

## Next Steps

1. **Review and approve** this proposal
2. **Create detailed design.md** with exact code patterns and APIs
3. **Write tasks.md** with step-by-step implementation checklist
4. **Implement Phase 1** (MVP with Word documents only)
5. **Iterate** based on learnings from Phase 1
