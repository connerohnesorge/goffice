# Design: OpenXML Part Reader Improvements

## Overview
Enhancing the robustness and performance of the OpenXML part reader system, focusing on error recovery, stream processing, and memory efficiency.

## Core Types & Interfaces

### Reader Interface
```go
type PartReader interface {
    Read(r io.Reader) (*Part, error)
    ReadStream(r io.Reader) (PartStream, error)
}

type PartStream interface {
    Next() (Element, error)
    Close() error
}
```

### Error Handling
```go
type ReadError struct {
    PartURI string
    Offset  int64
    Err     error
    Fatal   bool
}
```

## Architectural Decisions

### 1. Stream Processing
**Decision**: Use `encoding/xml` decoder with custom stream wrapper

**Rationale**:
- Low memory footprint for large parts
- Early exit on specific elements
- Handles partial reads

### 2. Error Recovery
**Decision**: Permissive parsing mode

**Rationale**:
- Many real-world documents are malformed
- Try to recover usable data
- Log warnings instead of hard failing when possible

## Implementation Strategy

### Phase 1: Robustness
- [ ] Implement permissive XML decoder
- [ ] Add namespace fallback handling
- [ ] Validate relationships during read

### Phase 2: Performance
- [ ] Buffer pooling for readers
- [ ] Lazy attribute parsing
- [ ] String interning for common repeated values (styles, themes)

## Testing Strategy
- Fuzz testing with malformed inputs
- Benchmark comparisons with previous reader
- Memory profile analysis on large documents