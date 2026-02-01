# Tasks: Implement Formula Function Calling in AST

## AST Implementation

- [ ] 1.1 Add FunctionCall struct to ast.go
- [ ] 1.2 Implement Eval method for FunctionCall
- [ ] 1.3 Add Position field for error reporting
- [ ] 1.4 Implement String method for debugging
- [ ] 1.5 Add visitor support for FunctionCall

## Function Registry

- [ ] 2.1 Create FunctionRegistry type
- [ ] 2.2 Implement Register method
- [ ] 2.3 Implement Get method (case-insensitive)
- [ ] 2.4 Add built-in function initialization
- [ ] 2.5 Add custom function registration support

## Built-in Functions - Mathematical

- [ ] 3.1 Implement SUM function
- [ ] 3.2 Implement AVERAGE function
- [ ] 3.3 Implement MIN function
- [ ] 3.4 Implement MAX function
- [ ] 3.5 Implement COUNT function
- [ ] 3.6 Implement ROUND function
- [ ] 3.7 Implement INT function
- [ ] 3.8 Implement ABS function
- [ ] 3.9 Implement POWER function
- [ ] 3.10 Implement SQRT function

## Built-in Functions - Logical

- [ ] 4.1 Implement IF function
- [ ] 4.2 Implement AND function
- [ ] 4.3 Implement OR function
- [ ] 4.4 Implement NOT function
- [ ] 4.5 Implement TRUE function
- [ ] 4.6 Implement FALSE function

## Built-in Functions - Text

- [ ] 5.1 Implement CONCATENATE function
- [ ] 5.2 Implement LEFT function
- [ ] 5.3 Implement RIGHT function
- [ ] 5.4 Implement MID function
- [ ] 5.5 Implement LEN function
- [ ] 5.6 Implement UPPER function
- [ ] 5.7 Implement LOWER function
- [ ] 5.8 Implement TRIM function

## Built-in Functions - Lookup/Reference

- [ ] 6.1 Implement VLOOKUP function
- [ ] 6.2 Implement HLOOKUP function
- [ ] 6.3 Implement INDEX function
- [ ] 6.4 Implement MATCH function

## Built-in Functions - Date/Time

- [ ] 7.1 Implement NOW function
- [ ] 7.2 Implement TODAY function
- [ ] 7.3 Implement DATE function
- [ ] 7.4 Implement TIME function
- [ ] 7.5 Implement YEAR function
- [ ] 7.6 Implement MONTH function
- [ ] 7.7 Implement DAY function

## Error Handling

- [ ] 8.1 Implement unknown function error
- [ ] 8.2 Implement wrong argument count error
- [ ] 8.3 Implement argument type error
- [ ] 8.4 Implement circular reference detection
- [ ] 8.5 Add error position information

## Unit Tests

- [ ] 9.1 Create test: Function name case insensitivity
- [ ] 9.2 Create test: Argument evaluation order
- [ ] 9.3 Create test: Nested function calls
- [ ] 9.4 Create test: Error propagation from arguments
- [ ] 9.5 Create test: Each mathematical function
- [ ] 9.6 Create test: Each logical function
- [ ] 9.7 Create test: Each text function
- [ ] 9.8 Create test: Each lookup function
- [ ] 9.9 Create test: Each date function
- [ ] 9.10 Create test: Custom function registration

## Integration Tests

- [ ] 10.1 Create test: Formula with single function
- [ ] 10.2 Create test: Formula with nested functions
- [ ] 10.3 Create test: Formula with cell references
- [ ] 10.4 Create test: Formula with ranges
- [ ] 10.5 Create test: Complex real-world formula

## Performance Tests

- [ ] 11.1 Benchmark: 100 simple function calls
- [ ] 11.2 Benchmark: 100 nested function calls
- [ ] 11.3 Benchmark: 1000 function calls with caching

## Documentation

- [ ] 12.1 Update formula/AGENTS.md with function calling
- [ ] 12.2 Document built-in function list
- [ ] 12.3 Document custom function registration
- [ ] 12.4 Add code comments to Eval methods

## Verification

- [ ] 13.1 Run all formula tests - ensure no regressions
- [ ] 13.2 Verify test coverage >90% for changed code
- [ ] 13.3 Test against Excel behavior for compatibility
- [ ] 13.4 Verify at least 20 functions implemented
