# Design: Streaming API Implementation

## Overview
Memory-efficient streaming APIs for large document processing without loading full DOM.

## Core Types & Interfaces

### Streaming Readers
```go
type StreamingPartReader interface {
    ReadElements(ctx context.Context) (<-chan ElementEvent, error)
    Close() error
}

type ElementEvent struct {
    Type  EventType
    Name  xml.Name
    Attrs []xml.Attr
    Data  []byte
}

type StreamingElementIterator interface {
    Next(context.Context) (Element, error)
    HasNext() bool
    Close() error
}
```

### Streaming Writers
```go
type StreamingPartWriter interface {
    WriteElements(context.Context, <-chan Element) error
    Close() error
}

type StreamingWriteOptions struct {
    ChunkSize      int
    BufferSize     int
    ValidateSchema bool
}
```

### Progress & Cancellation
```go
type ProgressCallback func(current, total int64)

type CancellationToken interface {
    IsCancelled() bool
    Cancel()
    Done() <-chan struct{}
}
```

## Architectural Decisions

### 1. Event-Driven Architecture
**Decision**: SAX-style event stream

**Rationale**:
- Constant memory usage
- Process elements as they arrive
- Support filtering/transformation
- Pipe-and-filter compatible

### 2. Chunked Processing
**Decision**: Configurable chunk size

**Rationale**:
- Balance memory and throughput
- Customizable per use case
- Parallel processing enabled
- Progress reporting

### 3. Validation During Streaming
**Decision**: Incremental validation

**Rationale**:
- Detect errors early
- Memory-efficient
- Constraint checking on arrival

## Implementation Strategy

### Phase 1: Core Streaming (1.5 weeks)
- [ ] ElementEvent types
- [ ] StreamingPartReader
- [ ] Event emitters
- [ ] Basic iteration

### Phase 2: Writing (1 week)
- [ ] StreamingPartWriter
- [ ] Element-to-Event conversion
- [ ] Encoding
- [ ] Validation

### Phase 3: Advanced (1.5 weeks)
- [ ] Progress reporting
- [ ] Cancellation support
- [ ] Filtering pipeline
- [ ] Transformations

### Phase 4: Integration (1 week)
- [ ] Document API hooks
- [ ] Packaging support
- [ ] Examples

## Performance Characteristics

- Memory: O(chunk size) not O(document size)
- Processing can start before full load
- Comparable throughput to DOM at chunk size 1000
