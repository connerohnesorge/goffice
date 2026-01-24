# Add Spreadsheet Formula Evaluation

## Overview
Implement Excel formula evaluation engine for calculating cell values. Enables reading spreadsheets with formulas and generating spreadsheets with computed values without requiring Excel installation.

## Motivation
Current goffice can read/write formula strings but cannot evaluate them. Users need calculated values for data exports, reports, and PDF rendering. Open-XML-SDK delegates to Excel; goffice needs built-in evaluation.

## Goals
- Implement formula parser for Excel formula syntax
- Implement evaluator for 50+ core functions:
  - Math: SUM, AVERAGE, COUNT, MIN, MAX, ROUND, ABS
  - Logical: IF, AND, OR, NOT
  - Text: CONCATENATE, LEFT, RIGHT, LEN, FIND
  - Lookup: VLOOKUP, HLOOKUP, INDEX, MATCH
  - Date/Time: TODAY, NOW, DATE, YEAR, MONTH, DAY
  - Financial: PMT, FV, PV, RATE, NPV, IRR
- Support cell references (A1, $A$1, Sheet1!A1)
- Support range references (A1:B10)
- Support named ranges
- Handle circular reference detection
- Support array formulas
- Enable partial evaluation (evaluate specific cells only)
- Provide dependency graph for optimization

## Non-Goals
- 100% Excel function compatibility (start with common 50)
- VBA or macro evaluation
- External data sources
- Custom functions (can be future)
- Real-time recalculation

## Dependencies
- Depends on: spreadsheet core
- Blocks: accurate spreadsheet PDF rendering
- Related: add-spreadsheet-pivot-table-support

## Technical Approach

### Parser
```go
type Formula struct {
    Expression string
    AST FormulaNode
}

func ParseFormula(expr string) (*Formula, error) {
    // Tokenize and parse to AST
}
```

### Evaluator
```go
type Evaluator struct {
    workbook *Workbook
    cache map[CellRef]Value
}

func (e *Evaluator) Evaluate(cell CellRef) (Value, error) {
    // Check cache
    // Parse formula
    // Evaluate AST
    // Detect circular refs
    // Cache result
}
```

### Functions
```go
var builtinFunctions = map[string]Function{
    "SUM": sumFunc,
    "IF": ifFunc,
    // ... 50+ functions
}
```

## Estimated Effort
16 weeks (1 engineer specialized in parsers)

## Priority
P0 - Required for accurate data processing
