# Useless Tests Report

## Criteria
- No assertions: test only logs, assigns, or comments "passes if no panic."
- Always skipped: unconditional `t.Skip` or tests that never run under default settings.
- Compile-time checks inside tests: interface conformance checks that would fail at build time regardless of test execution.
- Redundant wrappers: tests that only re-run other tests without adding assertions or new coverage.

## Findings

### Always skipped / unreachable by default
- `tests/e2e/framework/runner_test.go:75` `TestGoGenerator_Generate` - unconditional `t.Skip`, so it never executes.
- `tests/e2e/framework/runner_test.go:121` `TestCSharpGenerator_Generate` - unconditional `t.Skip`, so it never executes.
- `pdf/word/renderer_fields_test.go:546` `TestFieldsInHeadersFooters` - unconditional `t.Skip` with a note to rely on another test.
- `pdf/profiling_test.go:1` `TestProfile*` - build tag `profiling` means these never run in default `go test` (profiling harness, not validation).

### No-assert / log-only tests (always pass unless they panic)
- `drawingml/drawingml_test.go:8` `TestPackageExists` - only logs.
- `openxml/types/types_test.go:9` `TestPackageExists` - only logs.
- `wordprocessing/wordprocessing_test.go:10` `TestPackageExists` - only logs.
- `wordprocessing/parts/parts_test.go:10` `TestPackageExists` - only logs.
- `wordprocessing/elements/elements_test.go:15` `TestPackageExists` - only logs.
- `pdf/comparison/comparison_test.go:282` `TestIsGhostscriptAvailable` - only logs the result.
- `pdf/font/system_test.go:411` `TestFontRegistryScanDir` - only logs results, no expectations.
- `pdf/font/system_test.go:502` `TestFindSystemFont` - only logs, no assertion on outcome.
- `pdf/drawing/chart_renderer_test.go:122` `TestChartTitle` - no assertions; comment says "passes if no panic."
- `pdf/drawing/shape_renderer_test.go:145` `TestRoundRectangle` - no assertions.
- `pdf/drawing/shape_renderer_test.go:151` `TestEllipse` - no assertions.
- `pdf/drawing/shape_renderer_test.go:157` `TestPolygon` - no assertions in subtests.
- `pdf/drawing/shape_renderer_test.go:181` `TestStar` - no assertions in subtests.
- `pdf/layout/linebreak_word_test.go:866` `TestWordBreaking_MixedScripts` - logs only; no failure path.
- `pdf/layout/linebreak_word_test.go:1251` `TestWordBreaking_FilePaths` - logs only; no failure path.
- `pdf/layout/linebreak_word_test.go:1400` `TestWordBreaking_EdgeCases` - no assertions (only ensures no panic).
- `pdf/layout/word_breaks_test.go:347` `TestWordLineBreaker_KeepTitlesWithNames` - no assertions.
- `pdf/layout/word_breaks_test.go:528` `TestWordLineBreaker_EnDash` - no assertions.
- `pdf/layout/word_breaks_test.go:581` `TestWordLineBreaker_DisabledOptions` - no assertions.
- `pdf/layout/word_breaks_test.go:825` `TestWordLineBreaker_OnlySpaces` - no assertions.
- `pdf/layout/word_breaks_test.go:839` `TestWordLineBreaker_URLWithLongPath` - no assertions.
- `presentation/comprehensive_integration_test.go:833` `TestComprehensiveEncryptedPresentation` - logs only; no asserted outcome.

### Compile-time interface checks inside tests (no runtime signal)
- `openxml/relationship_test.go:661` `TestRelationshipInterfaceCompliance` - interface conformance checks that would fail at compile time anyway.
- `pdf/core/page_impl_test.go:8` `TestPageImpl_Interface` - compile-time check only.
- `pdf/core/page_impl_test.go:359` `TestMockPage_Interface` - compile-time check only.
- `pdf/core/page_test.go:1167` `TestPageDrawer_Interface` - compile-time check only.

### Redundant wrapper tests
- `pdf/validation_test.go:593` `TestPDFValidation_ComprehensiveBatch` - re-runs other tests via `t.Run` and only logs a summary; adds runtime cost without adding new assertions.

## Notes
- Some tests flagged by "no direct assertions" are actually useful because they delegate to helper functions that call `t.Error`/`t.Fatal` (e.g., fidelity tests). Those were excluded from this list after manual review.
- If you want these "useless" tests to add value, consider turning log-only checks into assertions or converting compile-time checks into package-level `var _ Interface = (*Type)(nil)` statements.
