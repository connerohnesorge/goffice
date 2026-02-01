# Proposal: Implement Formula Function Calling in AST

## Summary

Implement proper function calling in the spreadsheet formula AST to enable formula evaluation with function invocation.

## Background

The formula AST (`spreadsheet/formula/ast.go`) has a TODO indicating that proper function calling needs to be implemented:

**Reference TODO:** `spreadsheet/formula/ast.go:231`
```go
// TODO: Implement proper function calling
```

This is a critical component for formula evaluation support in the spreadsheet package.

## Motivation

Formula evaluation requires:
- Parsing function names and arguments
- Resolving function implementations
- Calling functions with evaluated arguments
- Handling nested function calls
- Supporting built-in Excel functions

Without proper function calling, formula evaluation is incomplete.

## Technical Design

### AST Function Call Node

Current structure likely has a placeholder. Need to implement:

```go
// FunctionCall represents a function call expression in the AST
type FunctionCall struct {
    Name      string        // Function name (e.g., "SUM", "IF")
    Arguments []Expression  // Evaluated arguments
    Pos       Position      // Source position for errors
}

func (f *FunctionCall) Eval(ctx *EvalContext) (Value, error) {
    // Look up function implementation
    fn, err := ctx.GetFunction(f.Name)
    if err != nil {
        return nil, fmt.Errorf("unknown function %s at %v: %w", f.Name, f.Pos, err)
    }
    
    // Evaluate all arguments
    args := make([]Value, len(f.Arguments))
    for i, arg := range f.Arguments {
        val, err := arg.Eval(ctx)
        if err != nil {
            return nil, fmt.Errorf("argument %d error in %s: %w", i, f.Name, err)
        }
        args[i] = val
    }
    
    // Call the function
    return fn(args)
}
```

### Function Registry

Need a registry to map function names to implementations:

```go
type FunctionRegistry struct {
    functions map[string]Function
}

type Function func(args []Value) (Value, error)

func (r *FunctionRegistry) Register(name string, fn Function) {
    r.functions[strings.ToUpper(name)] = fn
}

func (r *FunctionRegistry) Get(name string) (Function, error) {
    fn, ok := r.functions[strings.ToUpper(name)]
    if !ok {
        return nil, fmt.Errorf("function not found: %s", name)
    }
    return fn, nil
}
```

### Built-in Functions

Initial set of functions to implement:

**Mathematical:**
- SUM, AVERAGE, MIN, MAX, COUNT
- ROUND, INT, ABS, POWER, SQRT

**Logical:**
- IF, AND, OR, NOT, TRUE, FALSE

**Text:**
- CONCATENATE, LEFT, RIGHT, MID, LEN
- UPPER, LOWER, TRIM

**Lookup/Reference:**
- VLOOKUP, HLOOKUP, INDEX, MATCH

**Date/Time:**
- NOW, TODAY, DATE, TIME
- YEAR, MONTH, DAY, HOUR, MINUTE, SECOND

## Requirements

### SHALL Requirements

#### Requirement: Function Name Resolution
The AST SHALL resolve function names case-insensitively.

##### Scenario: Uppercase Function Name
Given a formula with =SUM(A1:A10)
When evaluated
Then the SUM function SHALL be called

##### Scenario: Mixed Case Function Name
Given a formula with =Sum(A1:A10)
When evaluated
Then the SUM function SHALL be called (case-insensitive)

#### Requirement: Argument Evaluation
The AST SHALL evaluate all arguments before calling the function.

##### Scenario: Cell Reference Argument
Given =SUM(A1, A2)
When evaluated
Then A1 and A2 SHALL be resolved to values before SUM is called

##### Scenario: Expression Argument
Given =SUM(A1*2, A2+5)
When evaluated
Then expressions SHALL be evaluated before SUM is called

##### Scenario: Nested Function Argument
Given =SUM(A1, MAX(B1:B10))
When evaluated
Then MAX SHALL be evaluated before SUM is called

#### Requirement: Error Propagation
The AST SHALL propagate errors from argument evaluation.

##### Scenario: Error in Argument
Given =SUM(A1, 1/0)
When evaluated
Then a division by zero error SHALL be returned

#### Requirement: Built-in Function Support
The implementation SHALL support at least 20 common Excel functions.

##### Scenario: SUM Function
Given =SUM(1, 2, 3)
When evaluated
Then the result SHALL be 6

##### Scenario: IF Function
Given =IF(TRUE, "yes", "no")
When evaluated
Then the result SHALL be "yes"

##### Scenario: VLOOKUP Function
Given =VLOOKUP("key", A1:B10, 2, FALSE)
When evaluated with matching data
Then the corresponding value SHALL be returned

### SHOULD Requirements

#### Requirement: Custom Function Registration
The implementation SHOULD allow registering custom functions.

#### Requirement: Variable Arguments
Functions SHOULD support variable argument counts (e.g., SUM can take any number of args).

## Testing Strategy

### Unit Tests
- Test function name resolution (case sensitivity)
- Test argument evaluation order
- Test error propagation
- Test each built-in function
- Test nested function calls

### Integration Tests
- End-to-end formula evaluation with functions
- Test with real spreadsheet data
- Test error scenarios

### Performance Tests
- Test evaluation speed with many function calls
- Test with deeply nested functions

## Implementation Plan

1. Implement FunctionCall AST node with Eval method
2. Create FunctionRegistry
3. Implement built-in mathematical functions
4. Implement built-in logical functions
5. Implement built-in text functions
6. Implement built-in lookup functions
7. Implement built-in date functions
8. Add unit tests
9. Add integration tests

## Related Changes

- `spreadsheet/formula/ast.go` - FunctionCall node implementation
- `spreadsheet/formula/functions.go` - NEW: Built-in functions
- `spreadsheet/formula/registry.go` - NEW: Function registry
- `spreadsheet/formula/eval.go` - Evaluation context updates

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| Complex function implementations | Medium | Start with simple functions |
| Circular references | High | Detect and handle circular refs |
| Performance with many calls | Low | Optimize after correctness |
| Excel compatibility | Medium | Test against Excel behavior |

## Acceptance Criteria

- [ ] Function names resolved case-insensitively
- [ ] Arguments evaluated before function call
- [ ] Nested functions work correctly
- [ ] At least 20 built-in functions implemented
- [ ] Unit tests pass with >90% coverage
- [ ] Integration tests pass
- [ ] Error handling works correctly
