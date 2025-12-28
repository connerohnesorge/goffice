# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

<!-- spectr:START -->
# Spectr Instructions

These instructions are for AI assistants working in this project.

Always open `@/spectr/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/spectr/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

When delegating tasks from a change proposal to subagents:
- Provide the proposal path: `spectr/changes/<id>/proposal.md`
- Include task context: `spectr/changes/<id>/tasks.jsonc`
- Reference delta specs: `spectr/changes/<id>/specs/<capability>/spec.md`

<!-- spectr:END -->

## Project Overview

**goffice** is a comprehensive Go SDK for manipulating Office Open XML (OOXML) documents. It provides feature parity with Microsoft's Open-XML-SDK for .NET, enabling creation, reading, and modification of Microsoft Office documents (.docx, .xlsx, .pptx) in pure Go.

**Key Features:**
- Word documents (.docx, .dotx, .docm, .dotm)
- Excel spreadsheets (.xlsx)
- PowerPoint presentations (.pptx, .potx, .ppsx)
- PDF rendering with high fidelity
- Full compliance with ECMA-376 and ISO/IEC 29500 standards
- Zero external dependencies (only Go stdlib + golang.org/x/text)

## Development Environment

### Setup
```bash
# Enter Nix development shell (includes all tools)
nix develop

# Or use direnv for automatic shell activation
direnv allow
```

The Nix flake provides:
- Go 1.25 toolchain
- golangci-lint, gopls, air, gotestsum
- All development tools pre-configured

### Common Commands

**Building:**
```bash
# Build all packages
go build ./...

# Build a specific package
go build ./wordprocessing
go build ./spreadsheet
go build ./presentation
```

**Testing:**
```bash
# Run all tests (standard)
go test ./...

# Run all tests with formatted output (if gotestsum is available)
gotestsum --format short-verbose ./...

# Run tests for a specific package
go test ./wordprocessing
go test ./drawingml
go test ./packaging

# Run a specific test
go test ./wordprocessing -run TestDocument

# Run tests with coverage
go test ./... -cover

# Run tests with race detector
go test ./... -race

# Run benchmarks
go test ./... -bench=.

# Run tests with timeout
go test ./... -timeout=5m
```

**Linting:**
```bash
# Run golangci-lint with auto-fix
lint  # (Nix script)

# Or manually
golangci-lint run --fix

# Run on specific directory
golangci-lint run --fix ./wordprocessing/...
```

**Formatting:**
```bash
# Format all code (runs gofmt, golines, goimports, alejandra)
nix fmt

# Or manually
gofmt -w .
goimports -w .
```

**Code Generation:**
```bash
# Regenerate wordprocessing elements from schema
go run ./cmd/gen-go-wordprocessing

# Regenerate presentation elements from schema
go run ./cmd/gen-go-presentation

# Regenerate spreadsheet elements from schema
go run ./cmd/gen-go-spreadsheet
```

**Development:**
```bash
# Live reload during development
air

# Edit flake.nix
dx  # (Nix script)
```

## Architecture

### Three-Layer Design

```
Application Layer (wordprocessing/, presentation/, spreadsheet/)
    ↓ uses
OpenXML Framework Layer (openxml/)
    ↓ uses
Packaging Layer (packaging/)
```

**1. Packaging Layer** (`packaging/`)
- Implements Open Packaging Conventions (OPC)
- ZIP package handling via `archive/zip`
- Relationship management
- Content type discovery
- Core properties (author, title, dates)

**2. OpenXML Framework Layer** (`openxml/`)
- Base element system (`BaseElement`, `CompositeElement`)
- XML serialization/deserialization
- Validation framework (ECMA-376 schema validation)
- Type system for attributes (`IntValue`, `BoolValue`, `OnOffValue`, etc.)
- Feature collection (dependency injection without global state)
- XML reading with namespace support

**3. Application Layer** (`wordprocessing/`, `presentation/`, `spreadsheet/`)
- Document-level APIs
- High-level manipulation (paragraphs, tables, slides, cells)
- Parts management (main document, styles, settings, images, etc.)
- Schema-driven generated element classes

### Shared Components

**DrawingML** (`drawingml/`)
- Shared drawing specification across all Office apps
- Shapes, charts, images, effects
- Color handling (RGB, scheme colors, system colors)
- Text properties, transforms, geometry

**PDF Rendering** (`pdf/`)
- PDF document generation from Office documents
- Font handling and subsetting (`pdf/font/`)
- Text layout engine (`pdf/layout/`)
- DrawingML to PDF rendering (`pdf/drawing/`)
- Word to PDF (`pdf/word/`)
- Excel/PowerPoint rendering in progress

## Code Organization

### Generated Code
Most element classes are **schema-driven generated code**:
- `wordprocessing/elements/*.go` (~58,698 lines generated)
- `presentation/elements/*.go`
- `spreadsheet/elements/*.go`
- Generated from Microsoft's Open-XML-SDK JSON schemas
- Provides type-safe element classes matching ECMA-376 standard

**Do not manually edit generated files** - modify the generators in `cmd/gen-go-*/` instead.

### Key Design Patterns

**1. Feature Collection Pattern**
```go
// Dependency injection without global state
element.Features().Set(validation.Strict())
child.Features().Get(validation.Key)  // Inherits from parent
```

**2. Lazy Loading**
```go
// Parts load XML only on first access
part := doc.MainDocumentPart()  // Doesn't parse XML yet
body := part.Document().Body()  // Now XML is parsed
```

**3. Fluent API**
```go
// Chainable methods for document construction
para := NewParagraph()
run := para.AppendRun()
run.AppendText("Hello World")
```

**4. Type-Safe Attributes**
```go
// Wrapped primitive types for optional XML attributes
elem.SetVal(types.NewInt32Value(42))
if val := elem.Val(); val != nil {
    num := val.Value()  // int32
}
```

### Testing Conventions

**Test Patterns:**
- Table-driven tests for multiple cases
- Subtests with `t.Run()` for organization
- Roundtrip tests: create → save → open → verify
- Fixture-based tests with real Office documents in `testdata/`

**Example Test Structure:**
```go
func TestFeature(t *testing.T) {
    t.Run("SubCase1", func(t *testing.T) {
        // Arrange
        // Act
        // Assert
    })
    t.Run("SubCase2", func(t *testing.T) {
        // ...
    })
}
```

### Naming Conventions
- Types mirror Microsoft Open-XML-SDK names: `WordprocessingDocument`, `MainDocumentPart`
- Methods use Go conventions: `GetStyleById()`, `AppendChild()`
- Enum values use PascalCase: `JustificationValues.Center`
- Private helpers use camelCase

## Important Constraints

### Zero Dependencies
- **Core library has zero external dependencies** (only Go stdlib)
- Single exception: `golang.org/x/text` for Unicode text handling
- No CGO - maximum portability
- Pure Go implementation

### Standards Compliance
- Full ECMA-376 standard compliance
- ISO/IEC 29500 conformance
- Must maintain compatibility with Microsoft Office applications
- Validation against official schemas

### Performance Considerations
- Lazy loading for large documents
- Streaming XML parsing where possible
- Minimal memory allocations
- Benchmark tests for critical paths

## Working with Generated Code

### When to Regenerate
Regenerate elements when:
- Microsoft updates Open-XML-SDK schemas
- Schema mappings change in generators
- Adding support for new Office features

### Generator Structure
Each generator (`cmd/gen-go-*/`) has:
- `gen_core.go` - Core generation logic
- `gen_enum.go` - Enum type generation
- `gen_struct.go` - Element struct generation
- `gen_methods.go` - Method generation
- `loaders.go` - Schema loading
- `types.go` - Type mappings
- `utils.go` - Helper functions

## PDF Rendering

The `pdf/` module renders Office documents to PDF:
- **High fidelity**: Matches Office rendering as closely as possible
- **Font handling**: Automatic font discovery, subsetting, embedding
- **Layout engine**: Text layout with line breaking, justification
- **Current status**: Word to PDF is functional, Excel/PowerPoint in progress

See `pdf/FONTS.md` for font configuration and `pdf/FIDELITY.md` for rendering accuracy details.

## Spectr Change Management

This project uses Spectr for structured change proposals:
- **Read first**: `spectr/AGENTS.md` before making major changes
- **Proposals**: `spectr/changes/<id>/proposal.md`
- **Specs**: `spectr/specs/<capability>/spec.md`
- **Use for**: New features, breaking changes, architectural shifts

## Directory Reference

**Core Packages:**
- `packaging/` - OPC layer (128K)
- `openxml/` - OpenXML framework (488K)
- `openxml/types/` - Typed attribute values
- `openxml/validation/` - Schema validation
- `openxml/features/` - Feature collection system

**Applications:**
- `wordprocessing/` - Word API (5.8M)
- `presentation/` - PowerPoint API (5.3M)
- `spreadsheet/` - Excel API (3.4M)

**Shared:**
- `drawingml/` - Drawing specification (556K)
- `pdf/` - PDF rendering (25M)

**Development:**
- `cmd/gen-go-*/` - Code generators
- `examples/` - Example usage
- `testdata/` - Test fixtures
- `internal/` - Internal utilities

## Tips

- **Read before editing**: Always read files before modifying them
- **Test after changes**: Run tests for affected packages
- **Check generated code**: Don't manually edit files with generation headers
- **Validate roundtrips**: Ensure save → open → save produces identical XML
- **Use golangci-lint**: The project has extensive linting rules - run `lint` before committing
- **Check compatibility**: Test with real Office applications when possible
