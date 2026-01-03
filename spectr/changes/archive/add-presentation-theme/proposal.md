# Change: Add Default Theme to Presentation Documents

## Why
Generated PowerPoint presentations currently lack a default theme (`theme1.xml`). This results in warnings during import in LibreOffice ("PowerPointImport::mpThemePtr is NULL") and likely causes rendering issues or fallback behaviors in other viewers. A valid PPTX typically requires at least one theme definition linked from the presentation or master.

## What Changes
- Modify `presentation.Document.initializeDocument` (or `PresentationPart.InitializeContent`) to automatically add a `ThemePart`.
- Ensure the `Presentation` element in `presentation.xml` (or the relevant Master) correctly references the new Theme part. Note: Usually themes are linked from the `SlideMaster` or `HandoutMaster`.
- Update `AddSlideMasterPart` logic to ensure new masters are linked to a theme (either the default one or a new one).

## Impact
- **Affected specs:** `presentation`
- **Affected code:**
    - `presentation/document.go`
    - `presentation/parts/presentation_part.go`
    - `presentation/parts/slide_master_part.go`
- **Breaking changes:** None. Additive.
