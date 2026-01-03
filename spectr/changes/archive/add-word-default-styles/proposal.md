# Change: Add Default Styles to Wordprocessing Documents

## Why
Generated Word documents currently lack a `styles.xml` part. This causes warnings in LibreOffice ("no style sheet found") and results in documents that may not render correctly or adhere to expected default formatting (e.g., "Normal" style). To ensure visual fidelity and valid OOXML structure, every new document should have a basic set of styles.

## What Changes
- Modify `wordprocessing.Document.initializeDocument` (or `MainPart.InitializeContent`) to automatically add a `StylesPart`.
- Implement a `StylesPart.InitializeDefault()` method that populates the part with a minimal valid `styles.xml` containing at least the `Normal` style and standard document defaults (`docDefaults`).
- Ensure the `MainPart` correctly references the `StylesPart`.

## Impact
- **Affected specs:** `wordprocessing`
- **Affected code:**
    - `wordprocessing/document.go`
    - `wordprocessing/parts/main_part.go`
    - `wordprocessing/parts/styles_part.go` (new method needed)
- **Breaking changes:** None. This is an additive change to document creation logic. Existing documents opened for reading will not be affected unless they are saved back, in which case we should ensure we don't overwrite existing styles if they exist.
