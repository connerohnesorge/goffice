# goffice Context

## Purpose
goffice is a Go SDK for Office Open XML (OOXML) document manipulation, providing feature parity with Microsoft's Open-XML-SDK for .NET. The initial focus is on Word document processing (WordprocessingML) with support for .docx, .dotx, .docm, and .dotm files.

## Tech Stack
- Go 1.25+
- Standard library only for core functionality:
  - `archive/zip` - OPC package handling
  - `encoding/xml` - XML serialization
  - `io`, `os` - File and stream operations
- No CGO dependencies
- No external runtime dependencies

## Project Conventions

### Code Style
- Follow official Go style guide and `gofmt`
- Use meaningful variable names; avoid single-letter names except for iterators
- Document all exported types, functions, and methods
- Use interfaces for abstraction; concrete types embed interface implementations
- Prefer composition over inheritance patterns
- Use generics where type safety improves API usability

### Package Structure
```
goffice/
├── packaging/           # OPC (Open Packaging Conventions) layer
├── openxml/             # Core OpenXML framework
│   ├── simpletypes/     # Typed attribute values
│   ├── features/        # Feature collection system
│   └── validation/      # Validation framework
├── wordprocessing/      # WordprocessingML implementation
│   ├── elements/        # Word document elements
│   └── parts/           # Word document parts
└── schema/
    └── wml/             # WordprocessingML schema types
```

### Naming Conventions
- Types mirror Open-XML-SDK names for familiarity (e.g., `WordprocessingDocument`, `MainDocumentPart`)
- Methods use Go conventions (e.g., `GetStyleById` not `GetStyleByID`)
- Enum values use PascalCase (e.g., `JustificationValues.Center`)
- Interface names use `I` prefix when matching C# SDK (e.g., `IFeatureCollection`)

### Architecture Patterns
- **Three-layer architecture**: Packaging → OpenXML → WordProcessing
- **Feature collection**: Configuration injection without global state
- **Lazy loading**: Parts load XML on first access, not on open
- **Immutable simple types**: Attribute value wrappers are value types
- **DOM-style API**: Element tree manipulation similar to XML DOM

### Testing Strategy
- Unit tests for all packages using standard `go test`
- Table-driven tests for edge cases
- Roundtrip tests: create → save → open → verify
- Fixture-based tests with real .docx files from different Office versions
- Benchmark tests for performance-critical paths
- No mocking of standard library types

### Git Workflow
- Main branch is protected; requires PR review
- Feature branches: `feature/description`
- Bug fixes: `fix/description`
- Conventional commits: `feat:`, `fix:`, `docs:`, `test:`, `refactor:`
- Squash merge for feature branches

## Domain Context

### Office Open XML (OOXML)
- ECMA-376 and ISO/IEC 29500 standardized format
- ZIP-based package containing XML parts
- Relationship-based linking between parts
- Multiple namespace versions (transitional vs. strict)

### Key Terminology
- **Package**: ZIP container for document parts
- **Part**: Individual XML or binary content within package
- **Relationship**: Link between parts (internal) or to external resources
- **Element**: XML node in a part's content
- **Content Type**: MIME type identifying part format

### Office Versions
- Office 2007 (ECMA-376 1st edition)
- Office 2010+ (extensions and strict mode)
- Microsoft 365 (latest features)

## Important Constraints

### Technical Constraints
- Must work on all Go-supported platforms
- No external dependencies for core functionality
- Must support streaming for large documents
- Thread-safe for concurrent read operations

### Compatibility Constraints
- Documents must open correctly in Microsoft Word
- Documents must open correctly in LibreOffice
- Must handle documents from Office 2007 through Microsoft 365
- Must handle malformed documents gracefully

### Performance Constraints
- Opening a 1MB document should complete in <100ms
- Memory usage should not exceed 10x document size
- Validation should complete in <1s for typical documents

## External Dependencies
- Reference implementation: [Open-XML-SDK](https://github.com/dotnet/Open-XML-SDK)
- Specifications: ECMA-376, ISO/IEC 29500
- Test documents: Various Office versions and third-party generators
