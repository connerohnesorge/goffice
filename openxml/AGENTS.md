# OPENXML FRAMEWORK KNOWLEDGE BASE

## OVERVIEW
The core framework layer that provides base types, XML serialization, and standard services for all Office document types.
Implements the Open Packaging Conventions (OPC) base structures and XML handling.

## STRUCTURE
```
openxml/
├── element.go      # Base interfaces (Element, CompositeElement, LeafElement)
├── features/       # Feature collection (dependency injection) system
├── validation/     # ECMA-376 schema validation engine
├── types/          # Core simple types (ST_*)
├── package.go      # OpenXmlPackage implementation
└── part.go         # OpenXmlPart implementation
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| **Base Types** | `element.go` | `CompositeElement` vs `LeafElement` |
| **Parsing** | `xml_reader.go` | XML decoding logic |
| **Writing** | `xml_writer.go` | XML encoding logic |
| **Attributes** | `attribute.go` | Handling namespaced attributes |
| **Validation** | `validation/` | Validating element structure against schema |

## CONVENTIONS
*   **Element Interface:** All nodes implement `Element`. `CompositeElement` has children; `LeafElement` has text.
*   **Features:** Use `elem.Features().Get(...)` to access services (like validation or namespace resolution).
*   **Namespaces:** Strictly handled via `OpenXmlQualifiedName` and `NamespaceURI`.
*   **Lazy Loading:** Parts are loaded on demand. XML is parsed only when the element hierarchy is accessed.
*   **PartRootElement:** The top-level element of any part (e.g., `Document`, `Workbook`, `ChartSpace`).

## ANTI-PATTERNS
*   **Casting:** Avoid unnecessary type casting; use interface methods where possible.
*   **Global State:** Do not introduce package-level variables for state. Use `Features`.
