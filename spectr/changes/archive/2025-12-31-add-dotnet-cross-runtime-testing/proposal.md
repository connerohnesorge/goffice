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

**Related Documents**:
- Detailed implementation: See `design.md`
- Task breakdown: See `tasks.md`
- E2E testing specification: See `specs/e2e-testing/spec.md`
- Nix environment specification: See `specs/nix-environment/spec.md`

### 1. New E2E Test Infrastructure (`tests/e2e/`)
- **Location**: New top-level `tests/e2e/` directory (user-requested)
- **Test matrix framework**: Go + C# bridge document generation with comparison
- **Test scenario definitions**: YAML definitions of documents to create
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
- **Scenario executor**: Runs identical test scenarios in both Go and C# bridge
- **Document generator bridges**:
  - Go bridge: Uses goffice API to create documents
  - C# bridge (csharp): Uses Open-XML-SDK to create documents
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

### 5. Visual Comparison Infrastructure
- **Reuse existing `pdf/comparison` package**:
  - CompareImages() - Pixel-level diff with tolerance
  - ConvertPDFToPNG() - PDF to PNG conversion via Ghostscript
  - No modifications to pdf/comparison needed
- **New E2E-specific features** (in `tests/e2e/comparison/`):
  - Office→PDF conversion via LibreOffice headless
  - Side-by-side document rendering (Go vs .NET)
  - Multi-page diff reports with thumbnails
  - Diff statistics (% match, pixel deltas, region highlighting)
- **HTML reporting** (in `tests/e2e/framework/`):
  - Interactive HTML reports (zoom, toggle overlay)
  - Embedded diff images and statistics
- **Why LibreOffice instead of goffice PDF renderer**:
  - Avoids circular testing (using goffice to test goffice)
  - LibreOffice provides neutral third-party rendering
  - Same conversion process for both Go and .NET documents
  - Validates that both implementations produce Office-compatible documents

### 6. CI/CD Integration
- **GitHub Actions workflow** (`.github/workflows/e2e-tests.yml`):
  - **Triggers**: Pull requests (all branches) + push to main
  - **Jobs**:
    1. nix-setup (install Nix with determinate-systems/nix-installer-action)
    2. build-bridges (dotnet build C# bridge, verify Go bridge compiles)
    3. run-e2e (nix develop --command test-e2e)
    4. upload-artifacts (diff images, HTML reports for failed tests)
  - **Timeout**: 15 minutes maximum
  - **Artifacts**: Stored for 7 days, downloadable from PR page
- **PR Comments** (via GitHub Actions bot):
  - **Format**: Pass/fail summary table + artifact download links
  - **Implementation**: actions/github-script or gh CLI
  - **Example**: "5/10 passed, 3 failed (see artifacts), 2 skipped"
- **Baseline Updates**:
  - **Trigger**: Manual dispatch workflow OR weekly cron (Monday 00:00 UTC)
  - **Process**: Run update-baselines script → create PR → assign for manual review
  - **Requires**: Human approval before merge (prevents accidental regressions)

## Impact

### Affected Specs
- **NEW: `e2e-testing`** - Cross-runtime end-to-end testing framework
- **NEW: `nix-environment`** - Multi-language Nix development environment

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
- **REUSED: `pdf/comparison/`** - Existing visual comparison utilities (no modifications needed)
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
├── scenarios/             # Test definitions (YAML)
│   ├── wordprocessing/
│   ├── spreadsheet/
│   └── presentation/
├── baselines/             # Reference documents from .NET SDK
│   ├── wordprocessing/
│   ├── spreadsheet/
│   └── presentation/
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

### 7. Test Data Privacy and Security
**Decision**: All test scenarios use synthetic, public-safe data only

**Rationale**:
- Test baselines committed to git (public repository)
- No real user data, no copyrighted content
- Scenarios use Lorem Ipsum text, geometric shapes, synthetic numbers
- Compliance with open-source project security best practices

**Guidelines for scenario authors**:
- **Text**: Use "Sample text", "Test document", placeholder names (e.g., "Jane Doe")
- **Numbers**: Use obviously fake data (Phone: "555-1234", SSN: "000-00-0000")
- **Images**: Simple geometric shapes or public domain images only
- **Dates**: Use clearly fictional dates (e.g., "2099-01-01")
- **Addresses**: Use "123 Test St, Example City, EX 12345"
- **No embedded files** from real Office documents or proprietary sources

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

1. **Coverage**: 40+ test scenarios across Word/Excel/PowerPoint (5 basic Word, 10 advanced Word, 10 Excel, 10 PowerPoint, 5 edge cases)
2. **Pass Rate**: >95% of scenarios have zero XML structure differences
3. **Visual Fidelity**: >99.5% pixel match for passing scenarios (0.5% diff threshold, avg across all tests)
4. **Performance**: E2E test suite completes in <10 minutes (wall-clock time, measured in CI)
   - **Expected baseline**: ~5-15 seconds per scenario (includes Go generation, C# generation, 3-level comparison)
   - **Phase 1 (5 scenarios)**: ~1 minute total
   - **Phase 3 (40+ scenarios)**: ~5-10 minutes total
   - **If exceeding 10 minutes**: Implement parallel execution or optimize bottlenecks
5. **Developer Experience**: `nix develop` + `test-e2e` succeeds on fresh clone
6. **Regression Detection**: Catch API divergence within 1 commit (via bisection when bugs found)
7. **CI Integration**: Runs on every PR, reports results automatically

## Resolved Design Decisions

1. **Scenario format**: YAML (human-readable, language-agnostic, extensive tooling support)
2. **Baseline storage**: Git for Phase 1-3 (small documents <5MB total), evaluate git-lfs in Phase 4 if total size >50MB
3. **Test data generation**: Manual scenarios for MVP (Phase 1-3), property-based testing deferred to future work
4. **Office version compatibility**: OOXML standard compliance (ECMA-376) ensures compatibility across Office 2016/2019/365
5. **Parallel execution**: Sequential in Phase 1-2 (simpler implementation), add parallel execution in Phase 4 if needed for performance
6. **Tolerance configuration**: Per-scenario overrides supported via YAML (defaults: VisualPixelTolerance=2.0, VisualDiffThreshold=0.005)
7. **C# bridge naming**: Named "DocxBridge" in Phase 1 for Word documents; rename to "OfficeBridge" in Phase 3 when adding Excel/PowerPoint support (breaking change acceptable since internal testing tool)

## Dependencies

**NOTE**: All versions listed below must match exactly across all documents (proposal.md, design.md, tasks.md). When updating a version, update ALL references.

- **.NET SDK 9.0.x**: Available in nixpkgs as `dotnet-sdk_9` (latest 9.0 patch version, pinned in flake.lock)
- **LibreOffice**: Available in nixpkgs as `libreoffice` (version pinned in flake.lock for reproducible Office→PDF conversion)
- **Ghostscript**: Available in nixpkgs as `ghostscript` (version pinned in flake.lock for PDF→PNG conversion)
- **Go 1.25**: Already specified in flake.nix and flake.lock
- **Open-XML-SDK submodule**: Already present at `Open-XML-SDK/` (commit hash in git submodule)
- **Go dependencies** (for tests/e2e/go.mod):
  - `gopkg.in/yaml.v3 v3.0.1` - YAML parsing for test scenarios
  - `github.com/connerohnesorge/goffice` - Via replace directive for local development
- **C# dependencies** (NuGet packages in DocxBridge.csproj):
  - `DocumentFormat.OpenXml 3.2.0` - Open-XML-SDK for .NET
  - `YamlDotNet 16.2.1` - YAML parsing for C#
  - `System.CommandLine 2.0.0-beta4.22272.1` - CLI argument parsing
- **pdf/comparison package**: Already exists in goffice codebase (no external dependency)

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| .NET SDK API mismatch | Can't create equivalent documents | Research .NET SDK patterns first (DONE) |
| Visual comparison flakiness | False positives on diffs | 1. Per-scenario tolerance tuning (pixel threshold + diff percentage)<br>2. Baseline regeneration workflow (update when .NET SDK changes)<br>3. Deterministic rendering (LibreOffice headless with fixed fonts)<br>4. Diff percentage threshold (ignore <0.5% differences) |
| Test suite too slow | Developers skip tests | 1. Pyramid approach: Fast XML tests fail first (seconds)<br>2. Parallel execution in Phase 4 if needed<br>3. Performance budget: <10 minutes for 40+ scenarios<br>4. CI timeout enforcement prevents runaway tests |
| Baseline drift | Baselines out of date | 1. Automated weekly check for .NET SDK updates<br>2. CI flags when baselines haven't updated in >90 days<br>3. Update script with git diff preview before commit |
| Nix complexity | Hard to setup | 1. direnv auto-activation (one-time `direnv allow`)<br>2. Comprehensive README with troubleshooting<br>3. CI validates Nix environment on every PR |
| Large diff reports | Storage/git bloat | 1. git for Phase 1-3 (baselines <5MB)<br>2. Evaluate git-lfs if total >50MB<br>3. HTML reports (text) preferred over binary screenshots |

## Next Steps

1. **Review and approve** this proposal
2. **Create detailed design.md** with exact code patterns and APIs
3. **Write tasks.md** with step-by-step implementation checklist
4. **Implement Phase 1** (MVP with Word documents only)
5. **Iterate** based on learnings from Phase 1
