# SPREADSHEET KNOWLEDGE BASE

## OVERVIEW
Implementation of the Excel (.xlsx) format. Handles Workbooks, Sheets, Cells, Styles, and Formulas.
Corresponds to `SpreadsheetDocument` in Open-XML-SDK.

## STRUCTURE
```
spreadsheet/
├── elements/     # GENERATED schema definitions (Rows, Cells, MergeCells)
├── formula/      # Formula parsing and calculation (tokenization)
├── parts/        # Opaque parts (charts, drawings integration)
├── cell.go       # Cell manipulation and typing
├── sheet.go      # Worksheet level operations
├── style.go      # Styling (Fonts, Fills, Borders)
└── document.go   # Entry point (Workbook management)
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Workbook** | `document.go` | `Create`, `Open`, `CreateFromTemplate` |
| **Rows/Cols** | `sheet.go` | Access rows, columns, iterators |
| **Cell Values** | `cell.go` | SetString, SetNumber, etc. |
| **Formulas** | `formula/` | Logic for evaluating Excel formulas |
| **Styles** | `style.go` | Managing the central StylesPart |

## CONVENTIONS
*   **DocType:** Supports `DocTypeWorkbook`, `DocTypeTemplate`, `DocTypeMacroEnabledWorkbook`.
*   **Cell References:** Use `A1` style strings or `row, col` indices.
*   **Shared Strings:** Text is automatically interned into the SharedStringsTable (usually invisible to API user).
*   **Dates:** Stored as numbers (OADate). Use helpers in `cell.go` to converting.

## ANTI-PATTERNS
*   **Direct Element Manipulation:** Prefer high-level helpers in `sheet.go`/`cell.go` over raw `elements/` struct access when possible.