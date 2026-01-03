## Context
PowerPoint files rely heavily on Themes (`theme1.xml`) to define color schemes, font schemes, and format schemes. A "blank" presentation is not truly blank; it relies on the "Office Theme". Missing this part causes corruption warnings.

## Goals
- Ensure every `goffice` generated PPTX has a valid theme structure.
- Eliminate "mpThemePtr is NULL" warnings in LibreOffice.

## Decisions
- **Decision:** The `ThemePart` will be added to the `PresentationPart` (as a shared resource) or directly to the `SlideMasterPart`?
    - *Analysis:* In typical OOXML, the `SlideMaster` has an explicit relationship to a `ThemePart`. The `Presentation` part *can* have a relationship to a default theme, but the Master-Theme link is critical for the slide rendering.
    - *Decision:* We will ensure `SlideMasterPart` adds and links a `ThemePart` when it is created. If `AddSlideMaster` is called, it should probably also create a Theme if one doesn't exist, or link to an existing one. Ideally, `AddSlideMasterPart` in `PresentationPart` should orchestrate this.

## Risks
- **Risk:** Bloating the file with duplicate themes if we create a new `theme#.xml` for every master.
    - *Mitigation:* Reuse the same `ThemePart` if possible, or accept that one-theme-per-master is a valid (and simple) starting point.

## Open Questions
- Does `presentation.xml` need a `defaultTextStyle`?
    - *Answer:* Probably, but the Theme is the priority error we are seeing.
