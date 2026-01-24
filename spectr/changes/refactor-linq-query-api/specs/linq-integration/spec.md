## ADDED Requirements

### Requirement: Iterator-Based Element Queries
The system SHALL provide iterator-based access to element hierarchies with lazy evaluation of query results.

#### Scenario: Query all descendants
- WHEN a Descendants query is created on an element
- THEN an iterator is returned that yields descendants
- AND the iterator is evaluated lazily as iterated
- AND the full tree is not materialized in memory

#### Scenario: Filter elements by type
- WHEN a Where filter is applied to an iterator
- THEN only matching elements are yielded
- AND non-matching elements are skipped
- AND the filter is evaluated during iteration

### Requirement: Axis-Based Navigation
The system SHALL support XPath-style axis navigation for element relationships.

#### Scenario: Navigate to parent
- WHEN the Parent axis is accessed
- THEN the parent element is returned or nil
- AND the element chain is preserved

#### Scenario: Navigate to siblings
- WHEN the FollowingSiblings axis is accessed
- THEN all following siblings are yielded
- AND siblings are in document order
- AND the iteration is efficient

### Requirement: Query Composition
The system SHALL support chaining multiple query operations in a fluent style.

#### Scenario: Complex query chain
- WHEN multiple query operations are chained
- THEN operations are composed lazily
- AND the final result is computed efficiently
- AND intermediate allocations are minimized
