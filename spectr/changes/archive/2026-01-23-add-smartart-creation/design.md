## Context
SmartArt is complex, requiring 4 synchronized XML parts and a specific `graphicFrame` reference on the slide. The current codebase can read them but not write them from scratch.

## Goals
- Enable "Hello World" SmartArt creation: A simple list of items.
- Hide the complexity of the 4 parts from the user where possible.

## Decisions
- **Decision:** We will hardcode a few "Standard" Layouts/Styles/Colors as embedded XML strings to start.
    - *Rationale:* Parsing/generating Layout definitions (constraints, rules) is extremely complex. Users typically just want "Basic List" or "Hierarchy". We can provide the XML for "Basic Block List" as a starting point.
- **Decision:** The API will focus on the *Data Model* (Points and Connections).
    - *Rationale:* The layout engine handles the positioning (in theory, though usually the creating app calculates it; for now we might rely on the viewing app to re-layout, or provide minimal pre-calculated positions if necessary. Note: PowerPoint usually re-calculates layout on load if data changes, but we might need to set a flag or minimal bounding box).

## Risks
- **Risk:** Viewers might not render the diagram if the layout cache (`drawing` part inside the diagram) is missing or empty.
    - *Mitigation:* We'll investigate if we need to generate a "fallback" image or if PPT handles re-layouting automatically. (PPT usually does if the parts are valid).

## Open Questions
- Do we need to generate the `drawing` part (cached vector representation)?
    - *Answer:* Ideally yes for compatibility, but modern PPT might regenerate it. We will test without it first (just data/layout/style/color).
