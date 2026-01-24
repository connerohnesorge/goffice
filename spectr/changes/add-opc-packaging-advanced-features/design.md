# Design: Advanced OPC Packaging

## Overview
Implementation of advanced Open Packaging Conventions (OPC) features, including digital signatures, optimized compression, relationship validation, and package cloning.

## Core Types & Interfaces

### Digital Signatures
```go
type PackageSigner struct {
    Cert *x509.Certificate
    Key  crypto.PrivateKey
}

type Signature struct {
    Parts     []string
    Timestamp time.Time
}
```

### Optimization
```go
type CompressionStrategy int

const (
    Store CompressionStrategy = iota
    Deflate
    DeflateFast
    DeflateBest
)

type PackageOptimizer struct {
    Strategy CompressionStrategy
}
```

## Architectural Decisions

### 1. Signing Layer
**Decision**: XML-DSIG implementation wrapper

**Rationale**:
- Standard for Office documents
- Integration with Go's crypto/x509
- Verifiable by Office apps

### 2. ZIP Abstraction
**Decision**: Enhanced ZIP writer wrapper

**Rationale**:
- Control over compression levels per part (mimetype = Store)
- Streaming support
- Deterministic output for signing

## Implementation Strategy

### Phase 1: Optimization & Cloning
- [ ] Implement DeepClone() for packages
- [ ] Add compression level controls
- [ ] Optimize mimetype storage (must be uncompressed)

### Phase 2: Relationship Validation
- [ ] Validate .rels targets
- [ ] Check for orphaned parts
- [ ] Detect circular dependencies

### Phase 3: Digital Signatures
- [ ] Implement XML-DSIG generation
- [ ] Add signature parts (_xmlsignatures)
- [ ] Verify existing signatures

## Performance Considerations

- Streaming ZIP writing to avoid memory buffering
- Parallel compression for large parts
- caching of checksums