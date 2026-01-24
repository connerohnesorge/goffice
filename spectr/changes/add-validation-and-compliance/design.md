# Design: Validation and Compliance

## Overview
Comprehensive validation framework for ISO/IEC 29500 compliance, including schema validation, semantic constraints, and document repair capabilities.

## Core Types & Interfaces

### Validator
```go
type ValidationLevel int

const (
    LevelInfo ValidationLevel = iota
    LevelWarning
    LevelError
    LevelCritical
)

type ValidationError struct {
    Level   ValidationLevel
    Path    string // XPath or part URI
    Message string
    RuleID  string
}

type Validator interface {
    Validate(pkg *Package) ([]ValidationError, error)
}
```

### Compliance Profiles
```go
type Profile interface {
    Name() string
    Rules() []Rule
}

// e.g., Strict, Transitional, Office2010, etc.
```

## Architectural Decisions

### 1. Rule Engine
**Decision**: Pluggable rule engine pattern

**Rationale**:
- Extensible for new constraints
- Configurable severity levels
- Separation of structural vs semantic rules

### 2. Schema Validation
**Decision**: Generated XSD validators vs Manual checks

**Rationale**:
- Use manual checks for critical structural constraints (performance)
- Use lightweight schema validation for simple types
- Avoid full XSD engine overhead if possible

## Implementation Strategy

### Phase 1: Structural Validation
- [ ] Part existence and relationships
- [ ] Content-Type verification
- [ ] XML well-formedness checks

### Phase 2: Semantic Validation
- [ ] ST_ types validation (patterns, enums)
- [ ] Parent-child containment rules
- [ ] Cross-reference integrity

### Phase 3: Repair
- [ ] Auto-fixer for common issues
- [ ] Orphan part cleanup
- [ ] Relationship healing

## Performance Considerations

- Fail-fast mode for critical errors
- Lazy validation (validate only accessed parts)
- Caching validation results if document immutable