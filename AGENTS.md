# PROJECT KNOWLEDGE BASE

**Generated:** 2026-01-02
**Project:** goffice

## OVERVIEW
Pure Go SDK for creating and editing Microsoft Office documents (Word, Excel, PowerPoint) and rendering them to PDF.
Achieves feature parity with C# Open-XML-SDK without external dependencies (except `golang.org/x/text`).
Strict adherence to ECMA-376 and ISO/IEC 29500 standards.

## STRUCTURE
```
goffice/
├── wordprocessing/     # Word (.docx) API & elements
├── spreadsheet/        # Excel (.xlsx) API & elements
├── presentation/       # PowerPoint (.pptx) API & elements
├── pdf/                # High-fidelity PDF rendering engine
│   ├── word/           # Word -> PDF renderer
│   ├── spreadsheet/    # Excel -> PDF renderer
│   └── presentation/   # PPT -> PDF renderer
├── drawingml/          # Shared drawing specs (charts, shapes, colors)
├── openxml/            # Core framework (XML, validation, base types)
├── packaging/          # OPC layer (ZIP, relationships, content types)
├── cmd/                # Code generators (schema -> Go structs)
├── Open-XML-SDK/       # Reference implementation / Schema source
└── testdata/           # Office document fixtures
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Core Logic** | `openxml/` | Base elements, XML marshalling, Features |
| **Doc Editing** | `wordprocessing/document.go` | Main entry (`Document`, `DocType`) |
| **Sheet Editing** | `spreadsheet/document.go` | Main entry (`Workbook`, `DocType`) |
| **PDF Rendering** | `pdf/pdf.go` | `RenderWord`, `RenderSpreadsheet`, `RenderOptions` |
| **Schema Defs** | `*/elements/` | GENERATED CODE - Do not edit manually |
| **Generators** | `cmd/gen-go-*/` | Modify these to change element code |
| **Validation** | `openxml/validation/` | ECMA-376 schema validation logic |

## CODE MAP
| Symbol | Type | Location | Role |
|--------|------|----------|------|
| `Document` | Struct | `*/document.go` | Root of the respective Office document |
| `CommonWrapper` | Interface | `openxml/element.go` | Base interface for all XML elements |
| `Element` | Interface | `openxml/element.go` | Base interface for OpenXML elements |
| `RenderOptions` | Struct | `pdf/pdf.go` | Config for PDF output (DPI, fonts, compliance) |
| `Features` | Method | `openxml/features/` | Dependency injection container |

## CONVENTIONS
*   **Zero Dependencies:** Only stdlib + `x/text`. No CGO.
*   **Generated Code:** Files in `elements/` are machine-generated. Edit `cmd/gen-go-*` instead.
*   **Feature Collection:** Use `element.Features()` for DI (no global state).
*   **Lazy Loading:** Parts/XML parsed only on access.
*   **Fluent API:** Method chaining preferred for construction (`para.AppendRun().AppendText(...)`).
*   **PDF Fidelity:** `pdf/` attempts to match Word's rendering exactly. Font metrics matter.

## ANTI-PATTERNS (THIS PROJECT)
*   **Manual XML:** Don't write XML strings directly; use the object model.
*   **Editing Generated Files:** Never modify `elements/*.go` files.
*   **Global State:** Avoid global vars; use `Features()` pattern.
*   **Hard Deps:** Do not add new external library dependencies without approval.

## COMMANDS
```bash
# Build
go build ./...

# Test (Standard)
go test ./...

# Test (Specific)
go test ./wordprocessing -run TestDocument

# Lint
golangci-lint run --fix

# Regenerate Elements (Example)
go run ./cmd/gen-go-wordprocessing
```

## NOTES
*   **PDF Rendering:** Supports Word, Excel, and PowerPoint. Configurable via `RenderOptions`.
*   **Open-XML-SDK:** The `Open-XML-SDK` folder is a reference/submodule, not the main Go code.
