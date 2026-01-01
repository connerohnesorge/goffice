# Implementation Tasks

## Phase 1: AST Type System

- [ ] 1.1 Create Expression interface with Evaluate(), Dependencies(), String() methods
- [ ] 1.2 Implement Literal expression for numbers, strings, booleans
- [ ] 1.3 Implement CellReference expression with sheet, col, row, absolute flags
- [ ] 1.4 Implement RangeReference expression with sheet, start/end cols/rows
- [ ] 1.5 Implement BinaryOp expression with left, right, operator
- [ ] 1.6 Implement UnaryOp expression with operand, operator
- [ ] 1.7 Implement FunctionCall expression with name, args
- [ ] 1.8 Add unit tests for all AST node types

## Phase 2: Parser

- [ ] 2.1 Create Parser struct with tokens, pos fields and helper methods
- [ ] 2.2 Implement parseExpression() (top-level)
- [ ] 2.3 Implement parseComparison() for =, <>, <, <=, >, >=
- [ ] 2.4 Implement parseConcatenation() for &
- [ ] 2.5 Implement parseAddition() for +, -
- [ ] 2.6 Implement parseMultiplication() for *, /
- [ ] 2.7 Implement parseExponentiation() for ^ (right-associative)
- [ ] 2.8 Implement parseUnary() for unary +, -
- [ ] 2.9 Implement parsePrimary() for literals, references, functions, parentheses
- [ ] 2.10 Add operator precedence tests with complex expressions

## Phase 3: Evaluation Engine

- [ ] 3.1 Create Value type system (ValueType enum, constructors)
- [ ] 3.2 Create EvalContext with Workbook, CurrentSheet, Resolver, FunctionRegistry
- [ ] 3.3 Implement WorkbookResolver for cell value resolution
- [ ] 3.4 Implement binary operation evaluators (add, subtract, multiply, divide, power, etc.)
- [ ] 3.5 Implement type conversion helpers (toNumber, toString, toBoolean)
- [ ] 3.6 Implement error propagation in all operators
- [ ] 3.7 Add unit tests for all operators and type conversions

## Phase 4: Function Registry

- [ ] 4.1 Create Function interface with Name(), MinArgs(), MaxArgs(), Call() methods
- [ ] 4.2 Create BaseFunction struct with default implementations
- [ ] 4.3 Create FunctionRegistry with Register(), Get(), Has() methods (thread-safe)
- [ ] 4.4 Implement DefaultFunctionRegistry() that registers all built-in functions
- [ ] 4.5 Create helper functions (flattenToNumbers, flattenToValues, matchCriteria)

## Phase 5: Mathematical Functions (18 functions)

- [ ] 5.1 Implement SUM function
- [ ] 5.2 Implement AVERAGE function
- [ ] 5.3 Implement COUNT and COUNTA functions
- [ ] 5.4 Implement MIN and MAX functions
- [ ] 5.5 Implement ROUND, ROUNDUP, ROUNDDOWN functions
- [ ] 5.6 Implement ABS, SQRT, POWER functions
- [ ] 5.7 Implement MOD and PRODUCT functions
- [ ] 5.8 Implement SUMIF, COUNTIF, AVERAGEIF functions
- [ ] 5.9 Add unit tests for all math functions with Excel reference values

## Phase 6: Logical Functions (7 functions)

- [ ] 6.1 Implement IF function
- [ ] 6.2 Implement AND, OR, NOT functions
- [ ] 6.3 Implement IFERROR and IFNA functions
- [ ] 6.4 Implement IFS function
- [ ] 6.5 Add unit tests for all logical functions

## Phase 7: Text Functions (11 functions)

- [ ] 7.1 Implement CONCATENATE function
- [ ] 7.2 Implement LEFT, RIGHT, MID functions
- [ ] 7.3 Implement LEN function
- [ ] 7.4 Implement UPPER, LOWER, TRIM functions
- [ ] 7.5 Implement FIND and SUBSTITUTE functions
- [ ] 7.6 Implement TEXT function with number formatting
- [ ] 7.7 Add unit tests for all text functions

## Phase 8: Lookup Functions (6 functions)

- [ ] 8.1 Implement VLOOKUP function with exact and approximate match
- [ ] 8.2 Implement HLOOKUP function
- [ ] 8.3 Implement INDEX function
- [ ] 8.4 Implement MATCH function with match_type support
- [ ] 8.5 Implement XLOOKUP function
- [ ] 8.6 Implement CHOOSE function
- [ ] 8.7 Add unit tests for all lookup functions

## Phase 9: Date/Time Functions (13 functions)

- [ ] 9.1 Implement Excel date serial number system (excelDateFromTime, timeFromExcelDate)
- [ ] 9.2 Implement TODAY and NOW functions
- [ ] 9.3 Implement DATE and TIME functions
- [ ] 9.4 Implement YEAR, MONTH, DAY functions
- [ ] 9.5 Implement HOUR, MINUTE, SECOND functions
- [ ] 9.6 Implement DATEDIF function with unit support (Y, M, D, MD, YM, YD)
- [ ] 9.7 Add unit tests for all date/time functions

## Phase 10: Statistical Functions (6 functions)

- [ ] 10.1 Implement MEDIAN function
- [ ] 10.2 Implement MODE function
- [ ] 10.3 Implement STDEV and VAR functions
- [ ] 10.4 Implement PERCENTILE and QUARTILE functions
- [ ] 10.5 Add unit tests for all statistical functions

## Phase 11: Dependency Graph

- [ ] 11.1 Create DependencyGraph struct with nodes, edges maps
- [ ] 11.2 Create DependencyNode with Cell, Formula, Dependencies, Dependents
- [ ] 11.3 Create CellKey with SheetName, Col, Row
- [ ] 11.4 Implement AddFormula() to extract dependencies and update graph
- [ ] 11.5 Implement RemoveFormula() to clean up graph
- [ ] 11.6 Implement TopologicalSort() using Kahn's algorithm
- [ ] 11.7 Implement circular reference detection (findCycle using DFS)
- [ ] 11.8 Implement GetDependents() to return cells depending on given cell
- [ ] 11.9 Add unit tests for dependency tracking and cycle detection

## Phase 12: Recalculation Engine

- [ ] 12.1 Implement RecalculateWorkbook() with full dependency graph and evaluation
- [ ] 12.2 Implement RecalculateWorksheet() scoped to single sheet
- [ ] 12.3 Implement RecalculateCell() for incremental recalculation
- [ ] 12.4 Implement RebuildCalculationChain() to update CalcChainPart
- [ ] 12.5 Add CalcChainPart API methods (AddCalculationCell, RemoveCalculationCell, Clear, GetOrder)
- [ ] 12.6 Add integration tests for recalculation scenarios

## Phase 13: High-Level API

- [ ] 13.1 Add Cell.EvaluateFormula() method
- [ ] 13.2 Add Cell.GetComputedValue() method
- [ ] 13.3 Add Cell.SetFormulaAndEvaluate() method
- [ ] 13.4 Add Worksheet.Recalculate() method
- [ ] 13.5 Add Worksheet.DependencyGraph() method
- [ ] 13.6 Add Worksheet.FindCircularReferences() method
- [ ] 13.7 Add Workbook.Recalculate() method
- [ ] 13.8 Add Workbook.RebuildCalculationChain() method
- [ ] 13.9 Add Workbook.CalculationOrder() method
- [ ] 13.10 Add Workbook.SetAutoRecalculate() method
- [ ] 13.11 Add auto-recalculation hooks to Cell.SetValue(), SetNumber(), SetString(), etc.
- [ ] 13.12 Add unit tests for high-level API

## Phase 14: Testing & Documentation

- [ ] 14.1 Write unit tests for all 60+ functions with Excel compatibility verification
- [ ] 14.2 Write integration tests (complex formulas, dependencies, circular references, recalculation)
- [ ] 14.3 Write performance benchmarks (parser, evaluator, 1000 formulas target <100ms)
- [ ] 14.4 Create Excel compatibility test spreadsheet with known formulas and results
- [ ] 14.5 Write formula/README.md with architecture and function documentation
- [ ] 14.6 Create example programs (simple evaluation, complex dependencies, circular detection, custom functions)
