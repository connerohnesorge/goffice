## Context
The current implementation of `goffice` creates valid ZIP packages for .docx files, but the internal structure lacks the standard `styles.xml` part that defines the look and feel of the document. This leads to "raw" XML rendering in viewers, which usually fallback to hardcoded application defaults but complain about the missing part.

## Goals / Non-Goals
- **Goal:** Ensure `goffice` generated documents open without warnings in LibreOffice and Word.
- **Goal:** Provide a baseline `Normal` style that matches standard Word defaults (Calibri/Arial, 11/12pt, standard spacing).
- **Non-Goal:** Implement a full style engine or complex style inheritance logic (that is a separate capability). We just want the *default* artifact to be present.

## Decisions
- **Decision:** Embed a minimal `styles.xml` template string in the Go code (similar to how `MainPart` works).
    - *Rationale:* It's simple, performance-friendly, and ensures consistency. Parsing/generating from objects for the *default* is unnecessary overhead unless we want dynamic defaults.
- **Decision:** Initialize this part automatically in `document.go`.
    - *Rationale:* The user shouldn't have to remember to `AddStylesPart()` for a basic "Hello World" document.

## Risks / Trade-offs
- **Risk:** Overwriting styles if we modify `InitializeContent` logic incorrectly for *opened* documents.
    - *Mitigation:* Ensure this logic is only called in `New()` or `initializeDocument` paths that are strictly for new document creation.

## Open Questions
- Should we expose configuration for the default font (e.g. Arial vs Calibri) in `OpenSettings`?
    - *Answer:* Not for this iteration. Stick to a sensible default (Calibri 11pt is standard for modern Word).
