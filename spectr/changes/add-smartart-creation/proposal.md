# Change: Add SmartArt Creation Support to Presentation

## Why
Currently, the `drawingml/diagram` package supports roundtripping existing SmartArt but lacks the ability to create new diagrams from scratch. Users cannot programmatically generate rich visual diagrams like lists, processes, or hierarchies in their presentations.

## What Changes
- Add `AddDiagram` methods to `presentation.SlidePart` to create and link the necessary 4 diagram parts (Data, Layout, Style, Colors).
- Implement a `DiagramBuilder` or helper API to simplify adding nodes ("Points") and text to the data model.
- Provide standard defaults for Layout (`layout1.xml` - Basic Block List), Style (`quickStyle1.xml` - Simple Fill), and Colors (`colors1.xml` - Colorful Accent).

## Impact
- **Affected specs:** `presentation`, `drawingml/diagram`
- **Affected code:**
    - `presentation/parts/slide_part.go`
    - `drawingml/diagram/` (new builder/helper logic)
- **Breaking changes:** None. Additive.
