# WORDPROCESSING KNOWLEDGE BASE

## OVERVIEW
Implementation of the Word (.docx) format. Handles Paragraphs, Runs, Tables, Styles, and Sections.
Corresponds to `WordprocessingDocument` in Open-XML-SDK.

## STRUCTURE
```
wordprocessing/
├── elements/     # GENERATED schema definitions (Body, P, R, Tbl)
├── parts/        # Headers, Footers, Settings
├── document.go   # Main document container & DocType definitions
├── builder.go    # Helpers for constructing complex elements
└── style.go      # Paragraph and Character styles
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Text** | `document.go` | `Document` struct, `AddParagraph`, `AddTable` |
| **Creation** | `document.go` | `New`, `Open`, `NewFromTemplate` |
| **Formatting** | `elements/r_pr.go` | Run properties (Bold, Italic) |
| **Tables** | `elements/tbl.go` | Low-level table construction |
| **Styles** | `style.go` | defining doc-wide styles |

## CONVENTIONS
*   **DocType:** Supports `DocTypeDocument`, `DocTypeTemplate`, `DocTypeMacroEnabled`.
*   **Hierarchy:** Document -> Body -> Paragraph -> Run -> Text.
*   **Properties:** Formatting is usually in `*Pr` elements (e.g., `PPr`, `RPr`).
*   **Images:** Added as relationships, referenced by ID in `drawing` elements.