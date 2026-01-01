# Proposal: Add Formula Evaluation Engine

## Summary

Add a comprehensive formula evaluation engine to the goffice spreadsheet package, enabling calculation of cell values from formulas without requiring Excel. This includes an expression parser, Abstract Syntax Tree (AST) representation, built-in function registry with 100+ Excel-compatible functions, dependency tracking, and incremental recalculation.

## Problem Statement

### Current State

goffice can create, read, and manipulate Excel spreadsheets (.xlsx) with formulas, but **cannot evaluate them**:

**What EXISTS:**
- Formula storage (`CellFormula` element with formula text)
- Formula tokenization (operators, functions, references, literals)
- Cell/range reference parsing (A1, R1C1, relative/absolute, cross-sheet)
- Formula rewriting when rows/columns are inserted/deleted
- Calculation chain infrastructure (structural only)
- Cell value types and error codes (#DIV/0!, #VALUE!, #REF!, etc.)

**What is MISSING:**
- Formula evaluation engine (no computation of results)
- Expression parser and AST builder
- Built-in function implementations (SUM, AVERAGE, IF, VLOOKUP, etc.)
- Cell resolver for fetching dependencies
- Dependency graph and topological sort
- Circular reference detection
- Error propagation during evaluation

### Impact

**Without formula evaluation:**
- ❌ Cannot compute cell values from formulas in Go
- ❌ Cannot validate formula correctness before saving
- ❌ Cannot build calculation engines or automation tools
- ❌ Cannot preview results without opening in Excel
- ❌ Cannot detect formula errors at creation time
- ❌ Reliant on Excel for all calculations

**Use cases blocked:**
- Server-side spreadsheet generation with computed values
- Automated report generation from templates
- Formula validation and error checking
- Batch processing of spreadsheets
- Testing spreadsheet logic
- Building calculation-heavy applications (financial models, simulations)

### Evidence from Codebase

From `spreadsheet/formula_rewriter.go:486` - Tokenization exists but stops there:
```go
func tokenizeFormula(formula string) []formulaToken {
    // Tokenizes "=SUM(A1:A10)+B1*2" into tokens
    // But no evaluation follows
}
```

From `spreadsheet/cell.go:340` - Formulas stored as strings:
```go
func (c *Cell) SetFormula(formula string) {
    // Stores formula text, no evaluation
}

func (c *Cell) GetFormula() string {
    // Returns formula text, no computed value
}
```

From `spreadsheet/parts/calc_chain_part.go:130` - Calculation chain empty:
```go
// CalcChainPart exists structurally
// But no mechanism to populate it with dependencies
```

## Proposed Solution

### Phase 1: Expression Parser & AST

**Add package:** `spreadsheet/formula/`

**Core Components:**

1. **AST Node Types** (`ast.go`):
```go
type Expression interface {
    Evaluate(ctx *EvalContext) (Value, error)
}

type BinaryOp struct {
    Left, Right Expression
    Operator    TokenType  // +, -, *, /, ^, &, =, <>, <=, >=, <, >
}

type FunctionCall struct {
    Name string
    Args []Expression
}

type CellReference struct {
    SheetName string
    Col       int
    Row       int
    ColAbs    bool
    RowAbs    bool
}

type RangeReference struct {
    SheetName  string
    StartCol   int
    StartRow   int
    EndCol     int
    EndRow     int
}

type Literal struct {
    Value interface{}  // number, string, boolean
}
```

2. **Parser** (`parser.go`):
```go
type Parser struct {
    tokens []formulaToken
    pos    int
}

func (p *Parser) Parse(formula string) (Expression, error) {
    tokens := tokenizeFormula(formula)
    return p.parseExpression(tokens)
}

func (p *Parser) parseExpression() Expression {
    // Recursive descent parser
    // Handles operator precedence
    // Builds AST from tokens
}
```

**Operator Precedence (highest to lowest):**
1. Parentheses `()`
2. Exponentiation `^`
3. Multiplication `*`, Division `/`
4. Addition `+`, Subtraction `-`
5. Concatenation `&`
6. Comparison `=`, `<>`, `<=`, `>=`, `<`, `>`

### Phase 2: Evaluation Engine

**Evaluation Context** (`context.go`):
```go
type EvalContext struct {
    Workbook      *Workbook
    CurrentSheet  *Worksheet
    Resolver      CellResolver
    FunctionRegistry *FunctionRegistry
}

type CellResolver interface {
    GetCellValue(sheet string, col, row int) (Value, error)
}

type Value struct {
    Type  ValueType  // Number, String, Boolean, Error, Array
    Value interface{}
}
```

**Evaluator** (`eval.go`):
```go
func (expr *BinaryOp) Evaluate(ctx *EvalContext) (Value, error) {
    left, err := expr.Left.Evaluate(ctx)
    if err != nil {
        return Value{}, err
    }
    right, err := expr.Right.Evaluate(ctx)
    if err != nil {
        return Value{}, err
    }

    switch expr.Operator {
    case TokenPlus:
        return add(left, right)
    case TokenMinus:
        return subtract(left, right)
    case TokenMult:
        return multiply(left, right)
    case TokenDiv:
        return divide(left, right)
    // ... other operators
    }
}

func (expr *FunctionCall) Evaluate(ctx *EvalContext) (Value, error) {
    fn := ctx.FunctionRegistry.Get(expr.Name)
    if fn == nil {
        return ErrorValue("#NAME?"), nil
    }

    args := make([]Value, len(expr.Args))
    for i, argExpr := range expr.Args {
        val, err := argExpr.Evaluate(ctx)
        if err != nil {
            return Value{}, err
        }
        args[i] = val
    }

    return fn.Call(args, ctx)
}
```

### Phase 3: Built-in Functions

**Function Registry** (`functions/registry.go`):
```go
type Function interface {
    Name() string
    MinArgs() int
    MaxArgs() int  // -1 for unlimited
    Call(args []Value, ctx *EvalContext) (Value, error)
}

type FunctionRegistry struct {
    functions map[string]Function
}

func (r *FunctionRegistry) Register(fn Function) {
    r.functions[strings.ToUpper(fn.Name())] = fn
}
```

**Priority 1: Mathematical Functions** (`functions/math.go`):
- `SUM(number1, [number2], ...)` - Sum of values
- `AVERAGE(number1, [number2], ...)` - Average
- `COUNT(value1, [value2], ...)` - Count of numbers
- `COUNTA(value1, [value2], ...)` - Count non-empty
- `MIN(number1, [number2], ...)` - Minimum value
- `MAX(number1, [number2], ...)` - Maximum value
- `ROUND(number, digits)` - Round to specified digits
- `ROUNDUP(number, digits)` - Round up
- `ROUNDDOWN(number, digits)` - Round down
- `ABS(number)` - Absolute value
- `SQRT(number)` - Square root
- `POWER(number, power)` - Exponentiation
- `MOD(number, divisor)` - Modulo
- `PRODUCT(number1, [number2], ...)` - Product
- `SUMIF(range, criteria, [sum_range])` - Conditional sum
- `COUNTIF(range, criteria)` - Conditional count
- `AVERAGEIF(range, criteria, [average_range])` - Conditional average

**Priority 2: Logical Functions** (`functions/logical.go`):
- `IF(condition, value_if_true, value_if_false)` - Conditional
- `AND(logical1, [logical2], ...)` - Logical AND
- `OR(logical1, [logical2], ...)` - Logical OR
- `NOT(logical)` - Logical NOT
- `IFERROR(value, value_if_error)` - Error handling
- `IFNA(value, value_if_na)` - #N/A handling
- `IFS(condition1, value1, [condition2, value2], ...)` - Multiple conditions

**Priority 3: Text Functions** (`functions/text.go`):
- `CONCATENATE(text1, [text2], ...)` - Join strings
- `LEFT(text, [num_chars])` - Left substring
- `RIGHT(text, [num_chars])` - Right substring
- `MID(text, start, num_chars)` - Middle substring
- `LEN(text)` - String length
- `UPPER(text)` - Convert to uppercase
- `LOWER(text)` - Convert to lowercase
- `TRIM(text)` - Remove extra spaces
- `FIND(find_text, within_text, [start_num])` - Find substring
- `SUBSTITUTE(text, old_text, new_text, [instance_num])` - Replace substring
- `TEXT(value, format_text)` - Format number as text

**Priority 4: Lookup Functions** (`functions/lookup.go`):
- `VLOOKUP(lookup_value, table_array, col_index, [range_lookup])` - Vertical lookup
- `HLOOKUP(lookup_value, table_array, row_index, [range_lookup])` - Horizontal lookup
- `INDEX(array, row_num, [column_num])` - Get value at index
- `MATCH(lookup_value, lookup_array, [match_type])` - Find index
- `XLOOKUP(lookup_value, lookup_array, return_array)` - Modern lookup
- `CHOOSE(index_num, value1, [value2], ...)` - Choose from list

**Priority 5: Date/Time Functions** (`functions/datetime.go`):
- `TODAY()` - Current date
- `NOW()` - Current date and time
- `DATE(year, month, day)` - Create date
- `TIME(hour, minute, second)` - Create time
- `YEAR(date)` - Extract year
- `MONTH(date)` - Extract month
- `DAY(date)` - Extract day
- `HOUR(time)` - Extract hour
- `MINUTE(time)` - Extract minute
- `SECOND(time)` - Extract second
- `DATEDIF(start_date, end_date, unit)` - Date difference

**Priority 6: Statistical Functions** (`functions/stats.go`):
- `MEDIAN(number1, [number2], ...)` - Median value
- `MODE(number1, [number2], ...)` - Most common value
- `STDEV(number1, [number2], ...)` - Standard deviation
- `VAR(number1, [number2], ...)` - Variance
- `PERCENTILE(array, k)` - Percentile value
- `QUARTILE(array, quart)` - Quartile value

**Total: 60+ functions in Phase 3, extensible to 400+ Excel functions**

### Phase 4: Dependency Tracking & Recalculation

**Dependency Graph** (`dependency.go`):
```go
type DependencyGraph struct {
    nodes map[CellRef]*DependencyNode
    edges map[CellRef][]CellRef  // Cell -> Cells that depend on it
}

type DependencyNode struct {
    Cell         CellRef
    Formula      Expression
    Dependencies []CellRef
    Dependents   []CellRef
}

func (g *DependencyGraph) AddFormula(cell CellRef, formula Expression) {
    deps := extractDependencies(formula)
    g.nodes[cell] = &DependencyNode{
        Cell:         cell,
        Formula:      formula,
        Dependencies: deps,
    }
    for _, dep := range deps {
        g.edges[dep] = append(g.edges[dep], cell)
    }
}

func (g *DependencyGraph) TopologicalSort() ([]CellRef, error) {
    // Kahn's algorithm for topological sort
    // Detects circular references
}
```

**Calculation Chain Builder** (`calc_chain.go`):
```go
func (w *Workbook) RebuildCalculationChain() error {
    graph := NewDependencyGraph()

    // 1. Find all formula cells
    for _, sheet := range w.Worksheets() {
        for _, cell := range sheet.Cells() {
            if formula := cell.GetFormula(); formula != "" {
                expr, err := ParseFormula(formula)
                if err != nil {
                    return err
                }
                graph.AddFormula(cell.Reference(), expr)
            }
        }
    }

    // 2. Topological sort
    order, err := graph.TopologicalSort()
    if err != nil {
        return err  // Circular reference detected
    }

    // 3. Update CalcChainPart
    calcChain := w.CalcChainPart()
    calcChain.Clear()
    for _, cellRef := range order {
        calcChain.AddCell(cellRef)
    }

    return nil
}
```

**Recalculation** (`recalc.go`):
```go
func (w *Workbook) Recalculate() error {
    order, err := w.CalculationOrder()
    if err != nil {
        return err
    }

    ctx := &EvalContext{
        Workbook:         w,
        FunctionRegistry: DefaultFunctionRegistry(),
        Resolver:         NewWorkbookResolver(w),
    }

    for _, cellRef := range order {
        cell := w.GetCell(cellRef)
        formula := cell.GetFormula()
        if formula == "" {
            continue
        }

        expr, err := ParseFormula(formula)
        if err != nil {
            cell.SetError("#VALUE!")
            continue
        }

        value, err := expr.Evaluate(ctx)
        if err != nil {
            cell.SetError("#VALUE!")
            continue
        }

        // Store computed value
        switch value.Type {
        case ValueTypeNumber:
            cell.SetNumber(value.Value.(float64))
        case ValueTypeString:
            cell.SetString(value.Value.(string))
        case ValueTypeBoolean:
            cell.SetBoolean(value.Value.(bool))
        case ValueTypeError:
            cell.SetError(value.Value.(string))
        }
    }

    return nil
}
```

**Incremental Recalculation:**
```go
func (w *Workbook) RecalculateCell(cellRef CellRef) error {
    // Recalculate only this cell and its dependents
    graph := w.DependencyGraph()
    affected := graph.GetDependents(cellRef)

    order := graph.TopologicalSort(affected)

    ctx := &EvalContext{...}
    for _, ref := range order {
        // Evaluate and update
    }
}
```

### Phase 5: Array Formulas & Advanced Features

**Array Formula Support:**
```go
type ArrayFormula struct {
    Expression Expression
    Range      RangeRef
}

func (a *ArrayFormula) Evaluate(ctx *EvalContext) ([][]Value, error) {
    // Evaluate once, distribute results across range
}
```

**Shared Formula Expansion:**
```go
func (c *Cell) ExpandSharedFormula() (Expression, error) {
    // Convert shared formula reference to actual formula
    // Adjust references based on cell offset
}
```

**Volatile Functions:**
```go
type VolatileFunction interface {
    Function
    IsVolatile() bool
}

// Volatile functions: NOW(), TODAY(), RAND(), RANDBETWEEN()
// Always recalculate, even if dependencies unchanged
```

### Phase 6: High-Level API

**Add to Cell:**
```go
// Evaluate formula and get result
func (c *Cell) EvaluateFormula() (interface{}, error)

// Get computed value (from cache or evaluate)
func (c *Cell) GetComputedValue() (interface{}, error)

// Set formula and immediately evaluate
func (c *Cell) SetFormulaAndEvaluate(formula string) error
```

**Add to Worksheet:**
```go
// Recalculate all formulas in this sheet
func (ws *Worksheet) Recalculate() error

// Get dependency graph for this sheet
func (ws *Worksheet) DependencyGraph() *DependencyGraph

// Find circular references
func (ws *Worksheet) FindCircularReferences() ([][]CellRef, error)
```

**Add to Workbook:**
```go
// Recalculate entire workbook
func (wb *Workbook) Recalculate() error

// Rebuild calculation chain
func (wb *Workbook) RebuildCalculationChain() error

// Get calculation order
func (wb *Workbook) CalculationOrder() ([]CellRef, error)

// Enable/disable automatic recalculation
func (wb *Workbook) SetAutoRecalculate(enabled bool)
```

## Architecture

### Package Structure

```
spreadsheet/formula/
├── ast.go             # AST node types (Expression, BinaryOp, FunctionCall, etc.)
├── parser.go          # Parser (tokenizeFormula -> AST)
├── eval.go            # Evaluator (AST -> Value)
├── context.go         # EvalContext, CellResolver, Value types
├── dependency.go      # DependencyGraph, TopologicalSort
├── calc_chain.go      # Calculation chain builder
├── recalc.go          # Recalculation engine
├── functions/
│   ├── registry.go    # Function registry
│   ├── math.go        # SUM, AVERAGE, COUNT, MIN, MAX, etc.
│   ├── logical.go     # IF, AND, OR, NOT, IFERROR, etc.
│   ├── text.go        # CONCATENATE, LEFT, RIGHT, MID, etc.
│   ├── lookup.go      # VLOOKUP, HLOOKUP, INDEX, MATCH, etc.
│   ├── datetime.go    # TODAY, NOW, DATE, YEAR, MONTH, etc.
│   └── stats.go       # MEDIAN, MODE, STDEV, VAR, etc.
└── functions_test.go  # Function tests
```

### Integration Points

1. **Cell API** (`spreadsheet/cell.go`):
   - Add `EvaluateFormula()`, `GetComputedValue()`, `SetFormulaAndEvaluate()`

2. **Worksheet API** (`spreadsheet/worksheet.go`):
   - Add `Recalculate()`, `DependencyGraph()`, `FindCircularReferences()`

3. **Workbook API** (`spreadsheet/workbook.go`):
   - Add `Recalculate()`, `RebuildCalculationChain()`, `SetAutoRecalculate()`

4. **Calculation Chain Part** (`spreadsheet/parts/calc_chain_part.go`):
   - Add `AddCell()`, `RemoveCell()`, `Clear()`, `GetOrder()`

5. **Formula Rewriter** (`spreadsheet/formula_rewriter.go`):
   - Hook into AST for smarter reference updates

## Dependencies

**Existing:**
- Formula tokenization (`formula_rewriter.go`)
- Cell/range reference parsing (`cellref.go`, `rangeref.go`)
- Cell value types (`cellvalue.go`)
- Calculation chain infrastructure (`parts/calc_chain_part.go`)

**New:**
- `spreadsheet/formula/` package (evaluation engine)
- `spreadsheet/formula/functions/` package (function implementations)

**External:**
- No new dependencies (pure Go stdlib)

## Success Criteria

- [ ] Parser converts formula strings to AST
- [ ] Evaluator computes results from AST
- [ ] 60+ built-in functions implemented (Math, Logical, Text, Lookup, Date, Stats)
- [ ] Cell resolver fetches dependencies correctly
- [ ] Dependency graph tracks formula relationships
- [ ] Topological sort produces correct calculation order
- [ ] Circular reference detection works
- [ ] Recalculation updates all dependent cells
- [ ] Array formulas evaluate correctly
- [ ] Shared formulas expand correctly
- [ ] Error propagation works (#DIV/0!, #VALUE!, #REF!, #NAME?, etc.)
- [ ] Roundtrip preserves formula text and computed values
- [ ] Performance: Evaluate 1000 formulas in <100ms
- [ ] Excel compatibility: Results match Excel for all test cases

## Out of Scope

- Advanced array formulas (FILTER, SORT, UNIQUE - defer to Excel 365 proposal)
- External data sources (database connections - defer to External Data proposal)
- User-defined functions (UDF - defer to VBA/Macros proposal)
- Formula auditing (trace precedents/dependents - defer to Formula Auditing proposal)
- Structured references (Table[@Column] - defer to Table Enhancement proposal)
- Dynamic arrays and spilling (defer to Excel 365 proposal)
- Calculation mode UI (manual/automatic - this is API only)

## Timeline Estimate

**Phase 1: Parser & AST** - ~1 week
**Phase 2: Evaluator** - ~1 week
**Phase 3: Functions (60+)** - ~3 weeks
**Phase 4: Dependencies** - ~1 week
**Phase 5: Advanced Features** - ~1 week
**Phase 6: High-Level API & Testing** - ~1 week

**Total: 8 weeks** for full implementation

## Backward Compatibility

**Fully backward compatible:**
- Existing formula storage/retrieval unchanged
- Formula rewriting still works
- Evaluation is opt-in via `Recalculate()` calls
- Files without formulas unaffected
- Formula text preserved exactly

## Related Proposals

- **None** - This is the first formula evaluation proposal

## References

- ECMA-376 §18.17: Formula calculation
- Excel function documentation: https://support.microsoft.com/en-us/office/excel-functions-by-category
- Microsoft Open-XML-SDK formula evaluation (not implemented in their SDK either)
- LibreOffice Calc formula engine (reference implementation)

## Implementation Notes

This proposal enables goffice to become a **full-featured spreadsheet engine**, not just a file format library. The evaluation engine will be:

1. **Excel-compatible**: All functions match Excel behavior exactly
2. **High-performance**: Minimal allocations, efficient dependency tracking
3. **Extensible**: Easy to add new functions via registry
4. **Robust**: Comprehensive error handling and circular reference detection
5. **Well-tested**: Every function has test cases matching Excel results
