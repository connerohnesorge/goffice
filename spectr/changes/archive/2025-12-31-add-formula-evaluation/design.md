# Formula Evaluation Engine - Complete Design

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [AST Type System](#ast-type-system)
3. [Parser Implementation](#parser-implementation)
4. [Evaluation Engine](#evaluation-engine)
5. [Function Registry](#function-registry)
6. [Built-in Functions](#built-in-functions)
7. [Dependency Tracking](#dependency-tracking)
8. [Recalculation Engine](#recalculation-engine)
9. [Error Handling](#error-handling)
10. [Performance Optimizations](#performance-optimizations)

---

## Architecture Overview

### Three-Layer Design

```
High-Level API Layer (Cell, Worksheet, Workbook)
          ↓
Evaluation Engine Layer (Parser, Evaluator, Functions)
          ↓
Foundation Layer (Tokenizer, CellRef, RangeRef)
```

**Foundation Layer** (existing in `spreadsheet/`):
- `tokenizeFormula()` - Splits formula into tokens
- `ParseCellRef()` - Parses A1, R1C1 references
- `ParseRangeRef()` - Parses range references
- Cell value types (Number, String, Boolean, Error)

**Evaluation Engine Layer** (new `spreadsheet/formula/`):
- `Parser` - Converts tokens to AST
- `Expression` - AST node interface
- `Evaluator` - Walks AST and computes result
- `FunctionRegistry` - Registry of built-in functions
- `DependencyGraph` - Tracks formula dependencies

**High-Level API Layer** (enhanced `spreadsheet/`):
- `Cell.EvaluateFormula()` - Evaluate single cell
- `Worksheet.Recalculate()` - Recalculate sheet
- `Workbook.Recalculate()` - Recalculate workbook

### Package Structure

```
spreadsheet/formula/
├── ast.go                 # Expression interface and AST nodes
├── parser.go              # Recursive descent parser
├── eval.go                # Expression evaluation
├── context.go             # Evaluation context and value types
├── dependency.go          # Dependency graph and topological sort
├── calc_chain.go          # Calculation chain builder
├── recalc.go              # Recalculation engine
├── errors.go              # Error types and handling
├── functions/
│   ├── registry.go        # Function registry
│   ├── function.go        # Function interface
│   ├── math.go            # SUM, AVERAGE, COUNT, MIN, MAX, etc. (18 functions)
│   ├── logical.go         # IF, AND, OR, NOT, IFERROR, etc. (7 functions)
│   ├── text.go            # CONCATENATE, LEFT, RIGHT, MID, etc. (11 functions)
│   ├── lookup.go          # VLOOKUP, HLOOKUP, INDEX, MATCH, etc. (6 functions)
│   ├── datetime.go        # TODAY, NOW, DATE, YEAR, MONTH, etc. (13 functions)
│   ├── stats.go           # MEDIAN, MODE, STDEV, VAR, etc. (6 functions)
│   └── helpers.go         # Shared helper functions
└── formula_test.go        # Comprehensive test suite
```

---

## AST Type System

### Expression Interface

All AST nodes implement the `Expression` interface:

```go
// spreadsheet/formula/ast.go

package formula

import (
    "github.com/unidoc/goffice/spreadsheet"
)

// Expression represents any evaluable formula expression
type Expression interface {
    // Evaluate computes the value of this expression
    Evaluate(ctx *EvalContext) (Value, error)

    // Dependencies returns cell/range references this expression depends on
    Dependencies() []Dependency

    // String returns a human-readable representation (for debugging)
    String() string
}

// Dependency represents a cell or range dependency
type Dependency struct {
    Type      DependencyType
    SheetName string
    CellRef   *spreadsheet.CellRef
    RangeRef  *spreadsheet.RangeRef
}

type DependencyType int

const (
    DependencyTypeCell DependencyType = iota
    DependencyTypeRange
)
```

### AST Node Types

#### 1. Literal

```go
// Literal represents a constant value (number, string, boolean)
type Literal struct {
    Value interface{}  // float64, string, or bool
    Type  ValueType
}

func NewNumberLiteral(n float64) *Literal {
    return &Literal{Value: n, Type: ValueTypeNumber}
}

func NewStringLiteral(s string) *Literal {
    return &Literal{Value: s, Type: ValueTypeString}
}

func NewBooleanLiteral(b bool) *Literal {
    return &Literal{Value: b, Type: ValueTypeBoolean}
}

func (l *Literal) Evaluate(ctx *EvalContext) (Value, error) {
    return Value{Type: l.Type, Value: l.Value}, nil
}

func (l *Literal) Dependencies() []Dependency {
    return nil  // Literals have no dependencies
}

func (l *Literal) String() string {
    switch l.Type {
    case ValueTypeNumber:
        return fmt.Sprintf("%v", l.Value)
    case ValueTypeString:
        return fmt.Sprintf("\"%s\"", l.Value)
    case ValueTypeBoolean:
        return fmt.Sprintf("%v", l.Value)
    default:
        return "UNKNOWN"
    }
}
```

#### 2. Cell Reference

```go
// CellReference represents a reference to a single cell (e.g., A1, Sheet1!B5, $A$1)
type CellReference struct {
    SheetName string
    Col       int
    Row       int
    ColAbs    bool  // true for $A
    RowAbs    bool  // true for $1
}

func NewCellReference(ref string) (*CellReference, error) {
    // Parse "A1", "$A$1", "Sheet1!B5", etc.
    cellRef, err := spreadsheet.ParseCellRef(ref)
    if err != nil {
        return nil, err
    }

    return &CellReference{
        SheetName: cellRef.SheetName,
        Col:       cellRef.Col,
        Row:       cellRef.Row,
        ColAbs:    cellRef.ColAbs,
        RowAbs:    cellRef.RowAbs,
    }, nil
}

func (cr *CellReference) Evaluate(ctx *EvalContext) (Value, error) {
    sheetName := cr.SheetName
    if sheetName == "" {
        sheetName = ctx.CurrentSheet.Name()
    }

    value, err := ctx.Resolver.GetCellValue(sheetName, cr.Col, cr.Row)
    if err != nil {
        return Value{}, err
    }

    // If empty, return 0 for numeric context
    if value.Type == ValueTypeEmpty {
        return NumberValue(0.0), nil
    }

    return value, nil
}

func (cr *CellReference) Dependencies() []Dependency {
    cellRef := &spreadsheet.CellRef{
        SheetName: cr.SheetName,
        Col:       cr.Col,
        Row:       cr.Row,
        ColAbs:    cr.ColAbs,
        RowAbs:    cr.RowAbs,
    }
    return []Dependency{{
        Type:    DependencyTypeCell,
        CellRef: cellRef,
    }}
}

func (cr *CellReference) String() string {
    colStr := ""
    if cr.ColAbs {
        colStr = "$"
    }
    colStr += spreadsheet.ColumnName(cr.Col)

    rowStr := ""
    if cr.RowAbs {
        rowStr = "$"
    }
    rowStr += fmt.Sprintf("%d", cr.Row)

    if cr.SheetName != "" {
        return fmt.Sprintf("%s!%s%s", cr.SheetName, colStr, rowStr)
    }
    return colStr + rowStr
}
```

#### 3. Range Reference

```go
// RangeReference represents a reference to a range of cells (e.g., A1:C10, Sheet1!B2:D20)
type RangeReference struct {
    SheetName string
    StartCol  int
    StartRow  int
    EndCol    int
    EndRow    int
}

func NewRangeReference(ref string) (*RangeReference, error) {
    rangeRef, err := spreadsheet.ParseRangeRef(ref)
    if err != nil {
        return nil, err
    }

    return &RangeReference{
        SheetName: rangeRef.SheetName,
        StartCol:  rangeRef.StartCol,
        StartRow:  rangeRef.StartRow,
        EndCol:    rangeRef.EndCol,
        EndRow:    rangeRef.EndRow,
    }, nil
}

func (rr *RangeReference) Evaluate(ctx *EvalContext) (Value, error) {
    // Ranges evaluate to array of values
    sheetName := rr.SheetName
    if sheetName == "" {
        sheetName = ctx.CurrentSheet.Name()
    }

    values := make([][]Value, 0)
    for row := rr.StartRow; row <= rr.EndRow; row++ {
        rowValues := make([]Value, 0)
        for col := rr.StartCol; col <= rr.EndCol; col++ {
            cell, err := ctx.Resolver.GetCellValue(sheetName, col, row)
            if err != nil {
                // Error resolving cell (e.g., sheet doesn't exist) - return error value
                rowValues = append(rowValues, ErrorValue(ErrorRef))
            } else if cell.Type == ValueTypeEmpty {
                // Empty cell - use 0
                rowValues = append(rowValues, NumberValue(0.0))
            } else {
                rowValues = append(rowValues, cell)
            }
        }
        values = append(values, rowValues)
    }

    return Value{Type: ValueTypeArray, Value: values}, nil
}

func (rr *RangeReference) Dependencies() []Dependency {
    rangeRef := &spreadsheet.RangeRef{
        SheetName: rr.SheetName,
        StartCol:  rr.StartCol,
        StartRow:  rr.StartRow,
        EndCol:    rr.EndCol,
        EndRow:    rr.EndRow,
    }
    return []Dependency{{
        Type:     DependencyTypeRange,
        RangeRef: rangeRef,
    }}
}

func (rr *RangeReference) String() string {
    start := fmt.Sprintf("%s%d", spreadsheet.ColumnName(rr.StartCol), rr.StartRow)
    end := fmt.Sprintf("%s%d", spreadsheet.ColumnName(rr.EndCol), rr.EndRow)
    if rr.SheetName != "" {
        return fmt.Sprintf("%s!%s:%s", rr.SheetName, start, end)
    }
    return fmt.Sprintf("%s:%s", start, end)
}
```

#### 4. Binary Operation

```go
// BinaryOp represents a binary operation (e.g., A1 + B1, 5 * 2, "Hello" & "World")
type BinaryOp struct {
    Left     Expression
    Right    Expression
    Operator TokenType
}

func NewBinaryOp(left Expression, op TokenType, right Expression) *BinaryOp {
    return &BinaryOp{
        Left:     left,
        Right:    right,
        Operator: op,
    }
}

func (bo *BinaryOp) Evaluate(ctx *EvalContext) (Value, error) {
    left, err := bo.Left.Evaluate(ctx)
    if err != nil {
        return Value{}, err
    }

    // Short-circuit error propagation
    if left.Type == ValueTypeError {
        return left, nil
    }

    right, err := bo.Right.Evaluate(ctx)
    if err != nil {
        return Value{}, err
    }

    if right.Type == ValueTypeError {
        return right, nil
    }

    // Array operations: If either operand is an array, use array evaluation
    // (element-wise for array-array, broadcasting for array-scalar)
    if left.Type == ValueTypeArray || right.Type == ValueTypeArray {
        return evaluateArrayBinaryOp(left, right, bo.Operator)
    }

    // Scalar operations
    switch bo.Operator {
    case TokenPlus:
        return add(left, right)
    case TokenMinus:
        return subtract(left, right)
    case TokenMult:
        return multiply(left, right)
    case TokenDiv:
        return divide(left, right)
    case TokenPower:
        return power(left, right)
    case TokenConcat:
        return concatenate(left, right)
    case TokenEQ:
        return equal(left, right)
    case TokenNE:
        return notEqual(left, right)
    case TokenLT:
        return lessThan(left, right)
    case TokenLE:
        return lessEqual(left, right)
    case TokenGT:
        return greaterThan(left, right)
    case TokenGE:
        return greaterEqual(left, right)
    default:
        return ErrorValue("#VALUE!"), nil
    }
}

func (bo *BinaryOp) Dependencies() []Dependency {
    leftDeps := bo.Left.Dependencies()
    rightDeps := bo.Right.Dependencies()
    return append(leftDeps, rightDeps...)
}

func (bo *BinaryOp) String() string {
    opStr := ""
    switch bo.Operator {
    case TokenPlus:
        opStr = "+"
    case TokenMinus:
        opStr = "-"
    case TokenMult:
        opStr = "*"
    case TokenDiv:
        opStr = "/"
    case TokenPower:
        opStr = "^"
    case TokenConcat:
        opStr = "&"
    case TokenEQ:
        opStr = "="
    case TokenNE:
        opStr = "<>"
    case TokenLT:
        opStr = "<"
    case TokenLE:
        opStr = "<="
    case TokenGT:
        opStr = ">"
    case TokenGE:
        opStr = ">="
    }
    return fmt.Sprintf("(%s %s %s)", bo.Left.String(), opStr, bo.Right.String())
}
```

#### 5. Unary Operation

```go
// UnaryOp represents a unary operation (e.g., -A1, +5)
type UnaryOp struct {
    Operand  Expression
    Operator TokenType
}

func NewUnaryOp(op TokenType, operand Expression) *UnaryOp {
    return &UnaryOp{
        Operator: op,
        Operand:  operand,
    }
}

func (uo *UnaryOp) Evaluate(ctx *EvalContext) (Value, error) {
    val, err := uo.Operand.Evaluate(ctx)
    if err != nil {
        return Value{}, err
    }

    if val.Type == ValueTypeError {
        return val, nil
    }

    switch uo.Operator {
    case TokenPlus:
        // Unary plus: convert to number
        return toNumber(val)
    case TokenMinus:
        // Unary minus: negate number
        num, err := toNumber(val)
        if err != nil {
            return ErrorValue("#VALUE!"), nil
        }
        return Value{Type: ValueTypeNumber, Value: -num.Value.(float64)}, nil
    default:
        return ErrorValue("#VALUE!"), nil
    }
}

func (uo *UnaryOp) Dependencies() []Dependency {
    return uo.Operand.Dependencies()
}

func (uo *UnaryOp) String() string {
    opStr := ""
    switch uo.Operator {
    case TokenPlus:
        opStr = "+"
    case TokenMinus:
        opStr = "-"
    }
    return fmt.Sprintf("(%s%s)", opStr, uo.Operand.String())
}
```

#### 6. Function Call

```go
// FunctionCall represents a function invocation (e.g., SUM(A1:A10), IF(B1>5, "Yes", "No"))
type FunctionCall struct {
    Name string
    Args []Expression
}

func NewFunctionCall(name string, args []Expression) *FunctionCall {
    return &FunctionCall{
        Name: strings.ToUpper(name),
        Args: args,
    }
}

func (fc *FunctionCall) Evaluate(ctx *EvalContext) (Value, error) {
    fn := ctx.FunctionRegistry.Get(fc.Name)
    if fn == nil {
        // Unknown function
        return ErrorValue("#NAME?"), nil
    }

    // Check argument count
    if len(fc.Args) < fn.MinArgs() {
        return ErrorValue("#VALUE!"), nil  // Too few arguments
    }
    if fn.MaxArgs() >= 0 && len(fc.Args) > fn.MaxArgs() {
        return ErrorValue("#VALUE!"), nil  // Too many arguments
    }

    // Evaluate all arguments
    args := make([]Value, len(fc.Args))
    for i, argExpr := range fc.Args {
        val, err := argExpr.Evaluate(ctx)
        if err != nil {
            return Value{}, err
        }
        args[i] = val
    }

    // Call function
    return fn.Call(args, ctx)
}

func (fc *FunctionCall) Dependencies() []Dependency {
    var deps []Dependency
    for _, arg := range fc.Args {
        deps = append(deps, arg.Dependencies()...)
    }
    return deps
}

func (fc *FunctionCall) String() string {
    argStrs := make([]string, len(fc.Args))
    for i, arg := range fc.Args {
        argStrs[i] = arg.String()
    }
    return fmt.Sprintf("%s(%s)", fc.Name, strings.Join(argStrs, ", "))
}
```

---

## Parser Implementation

### Token Type System

**Design Note - Token Type Integration**: The existing `spreadsheet` package has a generic `TokenType` enum with types like `TokenOperator`, `TokenFunction`, etc. For formula parsing with operator precedence, we need specific token types like `TokenPlus`, `TokenMinus`, etc. We extend the existing system as follows:

```go
// spreadsheet/formula/token.go

package formula

// TokenType represents specific operator/literal types for parsing
type TokenType int

const (
    TokenEOF TokenType = iota

    // Operators
    TokenPlus        // +
    TokenMinus       // -
    TokenMult        // *
    TokenDiv         // /
    TokenPower       // ^
    TokenConcat      // &

    // Comparison
    TokenEQ          // =
    TokenNE          // <>
    TokenLT          // <
    TokenLE          // <=
    TokenGT          // >
    TokenGE          // >=

    // Delimiters
    TokenLParen      // (
    TokenRParen      // )
    TokenComma       // ,
    TokenColon       // :

    // Literals and identifiers
    TokenNumber      // 123, 45.67
    TokenString      // "text"
    TokenBoolean     // TRUE, FALSE
    TokenReference   // A1, $A$1, Sheet1!B5
    TokenFunction    // SUM, IF, VLOOKUP
)

// Token represents a parsed token with type and value
type Token struct {
    Type  TokenType
    Value string
}

// tokenizeFormula converts a formula string to tokens with specific types
func tokenizeFormula(formula string) []Token {
    // Use existing spreadsheet.TokenizeFormula() to get generic tokens
    genericTokens := spreadsheet.TokenizeFormula(formula)

    // Map generic tokens to specific operator types
    tokens := make([]Token, 0, len(genericTokens))
    for _, gt := range genericTokens {
        token := Token{Value: gt.Value}

        switch gt.Type {
        case spreadsheet.TokenOperator:
            // Map operator value to specific type
            switch gt.Value {
            case "+":
                token.Type = TokenPlus
            case "-":
                token.Type = TokenMinus
            case "*":
                token.Type = TokenMult
            case "/":
                token.Type = TokenDiv
            case "^":
                token.Type = TokenPower
            case "&":
                token.Type = TokenConcat
            case "=":
                token.Type = TokenEQ
            case "<>":
                token.Type = TokenNE
            case "<":
                token.Type = TokenLT
            case "<=":
                token.Type = TokenLE
            case ">":
                token.Type = TokenGT
            case ">=":
                token.Type = TokenGE
            }
        case spreadsheet.TokenLParen:
            token.Type = TokenLParen
        case spreadsheet.TokenRParen:
            token.Type = TokenRParen
        case spreadsheet.TokenComma:
            token.Type = TokenComma
        case spreadsheet.TokenColon:
            token.Type = TokenColon
        case spreadsheet.TokenLiteral:
            // Determine if number or string
            if _, err := strconv.ParseFloat(gt.Value, 64); err == nil {
                token.Type = TokenNumber
            } else {
                token.Type = TokenString
            }
        case spreadsheet.TokenBoolean:
            token.Type = TokenBoolean
        case spreadsheet.TokenReference:
            token.Type = TokenReference
        case spreadsheet.TokenFunction:
            token.Type = TokenFunction
        }

        tokens = append(tokens, token)
    }

    return tokens
}
```

### Recursive Descent Parser

The parser converts tokens into an AST using recursive descent with operator precedence.

```go
// spreadsheet/formula/parser.go

package formula

import (
    "fmt"
    "strings"
)

// Parser converts formula tokens into an Abstract Syntax Tree
type Parser struct {
    tokens []Token
    pos    int
}

// Parse parses a formula string into an Expression
func Parse(formula string) (Expression, error) {
    // Remove leading "=" if present
    formula = strings.TrimPrefix(formula, "=")

    // Tokenize with specific operator types
    tokens := tokenizeFormula(formula)

    p := &Parser{
        tokens: tokens,
        pos:    0,
    }

    return p.parseExpression()
}

// parseExpression parses the top-level expression (lowest precedence: comparison)
func (p *Parser) parseExpression() (Expression, error) {
    return p.parseComparison()
}

// parseComparison parses comparison operators: =, <>, <, <=, >, >=
func (p *Parser) parseComparison() (Expression, error) {
    left, err := p.parseConcatenation()
    if err != nil {
        return nil, err
    }

    for {
        if p.match(TokenEQ, TokenNE, TokenLT,
            TokenLE, TokenGT, TokenGE) {
            op := p.previous().Type
            right, err := p.parseConcatenation()
            if err != nil {
                return nil, err
            }
            left = NewBinaryOp(left, op, right)
        } else {
            break
        }
    }

    return left, nil
}

// parseConcatenation parses concatenation operator: &
func (p *Parser) parseConcatenation() (Expression, error) {
    left, err := p.parseAddition()
    if err != nil {
        return nil, err
    }

    for p.match(TokenConcat) {
        right, err := p.parseAddition()
        if err != nil {
            return nil, err
        }
        left = NewBinaryOp(left, TokenConcat, right)
    }

    return left, nil
}

// parseAddition parses addition and subtraction: +, -
func (p *Parser) parseAddition() (Expression, error) {
    left, err := p.parseMultiplication()
    if err != nil {
        return nil, err
    }

    for p.match(TokenPlus, TokenMinus) {
        op := p.previous().Type
        right, err := p.parseMultiplication()
        if err != nil {
            return nil, err
        }
        left = NewBinaryOp(left, op, right)
    }

    return left, nil
}

// parseMultiplication parses multiplication and division: *, /
func (p *Parser) parseMultiplication() (Expression, error) {
    left, err := p.parseExponentiation()
    if err != nil {
        return nil, err
    }

    for p.match(TokenMult, TokenDiv) {
        op := p.previous().Type
        right, err := p.parseExponentiation()
        if err != nil {
            return nil, err
        }
        left = NewBinaryOp(left, op, right)
    }

    return left, nil
}

// parseExponentiation parses exponentiation: ^
func (p *Parser) parseExponentiation() (Expression, error) {
    left, err := p.parseUnary()
    if err != nil {
        return nil, err
    }

    // Right-associative
    if p.match(TokenPower) {
        right, err := p.parseExponentiation()
        if err != nil {
            return nil, err
        }
        left = NewBinaryOp(left, TokenPower, right)
    }

    return left, nil
}

// parseUnary parses unary operators: +, -
func (p *Parser) parseUnary() (Expression, error) {
    if p.match(TokenPlus, TokenMinus) {
        op := p.previous().Type
        expr, err := p.parseUnary()
        if err != nil {
            return nil, err
        }
        return NewUnaryOp(op, expr), nil
    }

    return p.parsePrimary()
}

// parsePrimary parses primary expressions: literals, references, functions, parentheses
func (p *Parser) parsePrimary() (Expression, error) {
    // Numbers and strings (from TokenNumber or TokenString)
    if p.match(TokenNumber) {
        token := p.previous()
        num, err := parseNumber(token.Value)
        if err != nil {
            return nil, fmt.Errorf("invalid number: %s", token.Value)
        }
        return NewNumberLiteral(num), nil
    }

    // Strings (quoted)
    if p.match(TokenString) {
        token := p.previous()
        // Remove quotes
        str := strings.Trim(token.Value, "\"")
        return NewStringLiteral(str), nil
    }

    // Booleans
    if p.match(TokenBoolean) {
        token := p.previous()
        b := strings.ToUpper(token.Value) == "TRUE"
        return NewBooleanLiteral(b), nil
    }

    // Cell references
    if p.match(TokenReference) {
        token := p.previous()
        // Check if it's a range (contains ":")
        if strings.Contains(token.Value, ":") {
            return NewRangeReference(token.Value)
        }
        return NewCellReference(token.Value)
    }

    // Functions
    if p.match(TokenFunction) {
        funcName := p.previous().Value

        if !p.match(TokenLParen) {
            return nil, fmt.Errorf("expected '(' after function name")
        }

        args := make([]Expression, 0)

        // Empty argument list
        if p.check(TokenRParen) {
            p.advance()
            return NewFunctionCall(funcName, args), nil
        }

        // Parse arguments
        for {
            arg, err := p.parseExpression()
            if err != nil {
                return nil, err
            }
            args = append(args, arg)

            if !p.match(TokenComma) {
                break
            }
        }

        if !p.match(TokenRParen) {
            return nil, fmt.Errorf("expected ')' after function arguments")
        }

        return NewFunctionCall(funcName, args), nil
    }

    // Parenthesized expressions
    if p.match(TokenLParen) {
        expr, err := p.parseExpression()
        if err != nil {
            return nil, err
        }

        if !p.match(TokenRParen) {
            return nil, fmt.Errorf("expected ')' after expression")
        }

        return expr, nil
    }

    return nil, fmt.Errorf("unexpected token: %v", p.peek())
}

// Helper methods

func (p *Parser) match(types ...TokenType) bool {
    for _, t := range types {
        if p.check(t) {
            p.advance()
            return true
        }
    }
    return false
}

func (p *Parser) check(t TokenType) bool {
    if p.isAtEnd() {
        return false
    }
    return p.peek().Type == t
}

func (p *Parser) advance() Token {
    if !p.isAtEnd() {
        p.pos++
    }
    return p.previous()
}

func (p *Parser) isAtEnd() bool {
    return p.pos >= len(p.tokens)
}

func (p *Parser) peek() Token {
    if p.isAtEnd() {
        return Token{Type: TokenEOF}
    }
    return p.tokens[p.pos]
}

func (p *Parser) previous() Token {
    return p.tokens[p.pos-1]
}

func parseNumber(s string) (float64, error) {
    return strconv.ParseFloat(s, 64)
}
```

**Operator Precedence (lowest to highest):**

1. Comparison: `=`, `<>`, `<`, `<=`, `>`, `>=`
2. Concatenation: `&`
3. Addition/Subtraction: `+`, `-`
4. Multiplication/Division: `*`, `/`
5. Exponentiation: `^` (right-associative)
6. Unary: `+`, `-`
7. Primary: literals, references, functions, `()`

---

## Evaluation Engine

### Evaluation Context

```go
// spreadsheet/formula/context.go

package formula

import (
    "github.com/unidoc/goffice/spreadsheet"
)

// EvalContext provides the context for formula evaluation
type EvalContext struct {
    Workbook         *spreadsheet.Workbook
    CurrentSheet     *spreadsheet.Worksheet
    Resolver         CellResolver
    FunctionRegistry *FunctionRegistry
}

// CellResolver resolves cell values by reference
type CellResolver interface {
    GetCellValue(sheetName string, col, row int) (Value, error)
}

// WorkbookResolver resolves cells from a workbook
type WorkbookResolver struct {
    workbook *spreadsheet.Workbook
}

func NewWorkbookResolver(wb *spreadsheet.Workbook) *WorkbookResolver {
    return &WorkbookResolver{workbook: wb}
}

func (wr *WorkbookResolver) GetCellValue(sheetName string, col, row int) (Value, error) {
    sheet := wr.workbook.SheetByName(sheetName)
    if sheet == nil {
        // Sheet doesn't exist - return #REF! error
        return ErrorValue(ErrorRef), nil
    }

    cell := sheet.Cell(row, col)
    if cell == nil {
        // Empty cell returns 0 for numeric context
        return NumberValue(0.0), nil
    }

    // Get cell value based on type
    switch cell.Type() {
    case spreadsheet.CellTypeNumber:
        num, _ := cell.GetNumber()
        return Value{Type: ValueTypeNumber, Value: num}, nil
    case spreadsheet.CellTypeString:
        str, _ := cell.GetString()
        return Value{Type: ValueTypeString, Value: str}, nil
    case spreadsheet.CellTypeBoolean:
        b, _ := cell.GetBoolean()
        return Value{Type: ValueTypeBoolean, Value: b}, nil
    case spreadsheet.CellTypeError:
        err, _ := cell.GetError()
        return Value{Type: ValueTypeError, Value: err}, nil
    case spreadsheet.CellTypeFormula:
        // Return cached value if available
        if cell.HasCachedValue() {
            return wr.getCachedValue(cell)
        }
        // Otherwise return empty (triggers re-evaluation)
        return EmptyValue(), nil
    default:
        return EmptyValue(), nil
    }
}

func (wr *WorkbookResolver) getCachedValue(cell *spreadsheet.Cell) (Value, error) {
    // Extract cached value from cell based on cell type
    switch cell.Type() {
    case spreadsheet.CellTypeNumber:
        num, _ := cell.GetNumber()
        return NumberValue(num), nil
    case spreadsheet.CellTypeString, spreadsheet.CellTypeInlineString:
        str := cell.GetString()
        return StringValue(str), nil
    case spreadsheet.CellTypeBoolean:
        b, _ := cell.GetBoolean()
        return BooleanValue(b), nil
    case spreadsheet.CellTypeError:
        errStr := cell.GetError()
        return ErrorValue(errStr), nil
    case spreadsheet.CellTypeFormula:
        // Formula cell - check if it has a cached result
        if cell.HasCachedValue() {
            return wr.getCachedValue(cell)  // Recursive to get cached result type
        }
        // No cached value - return empty to trigger evaluation
        return EmptyValue(), nil
    default:
        // Empty cell
        return EmptyValue(), nil
    }
}
```

### Value Type System

```go
// Value represents a computed value
type Value struct {
    Type  ValueType
    Value interface{}  // float64, string, bool, string (error), [][]Value (array)
}

type ValueType int

const (
    ValueTypeNumber ValueType = iota
    ValueTypeString
    ValueTypeBoolean
    ValueTypeError
    ValueTypeArray
    ValueTypeEmpty
)

// Constructor helpers
func NumberValue(n float64) Value {
    return Value{Type: ValueTypeNumber, Value: n}
}

func StringValue(s string) Value {
    return Value{Type: ValueTypeString, Value: s}
}

func BooleanValue(b bool) Value {
    return Value{Type: ValueTypeBoolean, Value: b}
}

func ErrorValue(err string) Value {
    return Value{Type: ValueTypeError, Value: err}
}

func ArrayValue(arr [][]Value) Value {
    return Value{Type: ValueTypeArray, Value: arr}
}

func EmptyValue() Value {
    return Value{Type: ValueTypeEmpty, Value: nil}
}

// Error constants
const (
    ErrorNull        = "#NULL!"
    ErrorDiv0        = "#DIV/0!"
    ErrorValue       = "#VALUE!"
    ErrorRef         = "#REF!"
    ErrorName        = "#NAME?"
    ErrorNum         = "#NUM!"
    ErrorNA          = "#N/A"
    ErrorGettingData = "#GETTING_DATA"
    ErrorSpill       = "#SPILL!"
    ErrorCalc        = "#CALC!"
)
```

### Binary Operation Evaluators

```go
// spreadsheet/formula/eval.go

package formula

import (
    "math"
    "strconv"
    "strings"
)

// errorPriority returns priority level for error values (lower = higher priority)
// Excel error priority: #NULL! > #DIV/0! > #VALUE! > #REF! > #NAME? > #NUM! > #N/A
func errorPriority(errValue string) int {
    priorities := map[string]int{
        ErrorNull:        0,
        ErrorDiv0:        1,
        ErrorValue:       2,
        ErrorRef:         3,
        ErrorName:        4,
        ErrorNum:         5,
        ErrorNA:          6,
        ErrorGettingData: 7,
        ErrorSpill:       8,
        ErrorCalc:        9,
    }
    if p, ok := priorities[errValue]; ok {
        return p
    }
    return 100 // Unknown errors have lowest priority
}

// propagateError returns the highest priority error
func propagateError(left, right Value) Value {
    // If only one is error, return it
    if left.Type == ValueTypeError && right.Type != ValueTypeError {
        return left
    }
    if right.Type == ValueTypeError && left.Type != ValueTypeError {
        return right
    }

    // Both are errors - return highest priority
    if left.Type == ValueTypeError && right.Type == ValueTypeError {
        leftErr := left.Value.(string)
        rightErr := right.Value.(string)
        if errorPriority(leftErr) < errorPriority(rightErr) {
            return left
        }
        return right
    }

    // Neither is error
    return Value{}
}

// add evaluates left + right
func add(left, right Value) (Value, error) {
    // Check for error propagation
    if errVal := propagateError(left, right); errVal.Type == ValueTypeError {
        return errVal, nil
    }

    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return NumberValue(l.Value.(float64) + r.Value.(float64)), nil
}

// subtract evaluates left - right
func subtract(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return NumberValue(l.Value.(float64) - r.Value.(float64)), nil
}

// multiply evaluates left * right
func multiply(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return NumberValue(l.Value.(float64) * r.Value.(float64)), nil
}

// divide evaluates left / right
func divide(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    rightNum := r.Value.(float64)
    if rightNum == 0 {
        return ErrorValue(ErrorDiv0), nil
    }

    return NumberValue(l.Value.(float64) / rightNum), nil
}

// power evaluates left ^ right
func power(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    result := math.Pow(l.Value.(float64), r.Value.(float64))
    if math.IsNaN(result) || math.IsInf(result, 0) {
        return ErrorValue(ErrorNum), nil
    }

    return NumberValue(result), nil
}

// concatenate evaluates left & right
func concatenate(left, right Value) (Value, error) {
    l := toString(left)
    r := toString(right)
    return StringValue(l + r), nil
}

// equal evaluates left = right
func equal(left, right Value) (Value, error) {
    // Type coercion: numbers first, then strings
    if left.Type == ValueTypeNumber && right.Type == ValueTypeNumber {
        return BooleanValue(left.Value.(float64) == right.Value.(float64)), nil
    }

    leftStr := toString(left)
    rightStr := toString(right)
    // Case-insensitive comparison for strings
    return BooleanValue(strings.EqualFold(leftStr, rightStr)), nil
}

// notEqual evaluates left <> right
func notEqual(left, right Value) (Value, error) {
    eq, err := equal(left, right)
    if err != nil {
        return Value{}, err
    }
    return BooleanValue(!eq.Value.(bool)), nil
}

// lessThan evaluates left < right
func lessThan(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return BooleanValue(l.Value.(float64) < r.Value.(float64)), nil
}

// lessEqual evaluates left <= right
func lessEqual(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return BooleanValue(l.Value.(float64) <= r.Value.(float64)), nil
}

// greaterThan evaluates left > right
func greaterThan(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return BooleanValue(l.Value.(float64) > r.Value.(float64)), nil
}

// greaterEqual evaluates left >= right
func greaterEqual(left, right Value) (Value, error) {
    l, err := toNumber(left)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }
    r, err := toNumber(right)
    if err != nil {
        return ErrorValue(ErrorValue), nil
    }

    return BooleanValue(l.Value.(float64) >= r.Value.(float64)), nil
}

// Type conversion helpers

func toNumber(v Value) (Value, error) {
    switch v.Type {
    case ValueTypeNumber:
        return v, nil
    case ValueTypeBoolean:
        if v.Value.(bool) {
            return NumberValue(1.0), nil
        }
        return NumberValue(0.0), nil
    case ValueTypeString:
        // Try to parse string as number
        num, err := strconv.ParseFloat(v.Value.(string), 64)
        if err != nil {
            return Value{}, fmt.Errorf("cannot convert to number")
        }
        return NumberValue(num), nil
    case ValueTypeEmpty:
        return NumberValue(0.0), nil
    case ValueTypeError:
        return v, nil
    default:
        return Value{}, fmt.Errorf("cannot convert to number")
    }
}

func toString(v Value) string {
    switch v.Type {
    case ValueTypeString:
        return v.Value.(string)
    case ValueTypeNumber:
        return fmt.Sprintf("%v", v.Value)
    case ValueTypeBoolean:
        if v.Value.(bool) {
            return "TRUE"
        }
        return "FALSE"
    case ValueTypeError:
        return v.Value.(string)
    case ValueTypeEmpty:
        return ""
    default:
        return ""
    }
}

func toBoolean(v Value) (Value, error) {
    switch v.Type {
    case ValueTypeBoolean:
        return v, nil
    case ValueTypeNumber:
        return BooleanValue(v.Value.(float64) != 0), nil
    case ValueTypeString:
        s := strings.ToUpper(v.Value.(string))
        if s == "TRUE" {
            return BooleanValue(true), nil
        }
        if s == "FALSE" {
            return BooleanValue(false), nil
        }
        return Value{}, fmt.Errorf("cannot convert to boolean")
    case ValueTypeError:
        return v, nil
    default:
        return Value{}, fmt.Errorf("cannot convert to boolean")
    }
}

// containsVolatileFunction checks if an expression contains any volatile functions
// Volatile functions: NOW(), TODAY(), RAND(), RANDBETWEEN() - always recalculate
func containsVolatileFunction(expr Expression) bool {
    switch e := expr.(type) {
    case *FunctionCall:
        // Check if the function itself is volatile (common volatile function names)
        volatileFuncs := map[string]bool{
            "NOW":         true,
            "TODAY":       true,
            "RAND":        true,
            "RANDBETWEEN": true,
        }
        if volatileFuncs[strings.ToUpper(e.Name)] {
            return true
        }
        // Check arguments recursively
        for _, arg := range e.Args {
            if containsVolatileFunction(arg) {
                return true
            }
        }
    case *BinaryOp:
        return containsVolatileFunction(e.Left) || containsVolatileFunction(e.Right)
    case *UnaryOp:
        return containsVolatileFunction(e.Operand)
    }
    return false
}

// Helper functions for conditional aggregation (SUMIF, COUNTIF, AVERAGEIF)
// These are used by functions package but defined here to avoid import cycles

// matchCriteria checks if a value matches a criteria string
// Criteria can be:
// - Number: 5, 10.5
// - String: "text", "hello"
// - Comparison: ">5", ">=10", "<100", "<=50", "<>0", "=text"
func matchCriteria(value Value, criteriaStr string) bool {
    // Parse criteria string
    if len(criteriaStr) == 0 {
        return false
    }

    // Check for comparison operators
    operators := []string{"<>", ">=", "<=", ">", "<", "="}
    for _, op := range operators {
        if strings.HasPrefix(criteriaStr, op) {
            // Extract comparison value
            compareStr := strings.TrimPrefix(criteriaStr, op)
            compareVal := parseCriteriaValue(compareStr)

            return evaluateComparison(value, op, compareVal)
        }
    }

    // No operator - exact match
    criteriaVal := parseCriteriaValue(criteriaStr)
    return exactMatch(value, criteriaVal)
}

// parseCriteriaValue parses a criteria string into a Value
func parseCriteriaValue(s string) Value {
    // Try number
    if num, err := strconv.ParseFloat(s, 64); err == nil {
        return NumberValue(num)
    }

    // Try boolean
    if strings.EqualFold(s, "TRUE") {
        return BooleanValue(true)
    }
    if strings.EqualFold(s, "FALSE") {
        return BooleanValue(false)
    }

    // Default to string
    return StringValue(s)
}

// evaluateComparison evaluates value op compareVal
func evaluateComparison(value Value, op string, compareVal Value) bool {
    switch op {
    case ">":
        cmp := compareForApprox(value, compareVal)
        return cmp > 0
    case ">=":
        cmp := compareForApprox(value, compareVal)
        return cmp >= 0
    case "<":
        cmp := compareForApprox(value, compareVal)
        return cmp < 0
    case "<=":
        cmp := compareForApprox(value, compareVal)
        return cmp <= 0
    case "=":
        return exactMatch(value, compareVal)
    case "<>":
        return !exactMatch(value, compareVal)
    default:
        return false
    }
}
```

---

## Function Registry

### Function Interface

```go
// spreadsheet/formula/functions/function.go

package functions

import (
    "github.com/unidoc/goffice/spreadsheet/formula"
)

// Function represents a built-in Excel function
type Function interface {
    // Name returns the function name (uppercase)
    Name() string

    // MinArgs returns the minimum number of arguments
    MinArgs() int

    // MaxArgs returns the maximum number of arguments (-1 for unlimited)
    MaxArgs() int

    // Call executes the function with given arguments
    Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error)
}

// BaseFunction provides common function implementation
type BaseFunction struct {
    name    string
    minArgs int
    maxArgs int
}

func (bf *BaseFunction) Name() string {
    return bf.name
}

func (bf *BaseFunction) MinArgs() int {
    return bf.minArgs
}

func (bf *BaseFunction) MaxArgs() int {
    return bf.maxArgs
}

// Volatile Functions Support
//
// Volatile functions recalculate on every calculation pass, regardless of whether their
// dependencies have changed. Examples: NOW(), TODAY(), RAND(), RANDBETWEEN()

// VolatileFunction marker interface for functions that always recalculate
type VolatileFunction interface {
    Function
    IsVolatile() bool
}

// VolatileBaseFunction extends BaseFunction with volatile support
type VolatileBaseFunction struct {
    BaseFunction
}

func (vbf *VolatileBaseFunction) IsVolatile() bool {
    return true
}
```

### Registry Implementation

```go
// spreadsheet/formula/functions/registry.go

package functions

import (
    "strings"
    "sync"
)

// FunctionRegistry manages built-in functions
type FunctionRegistry struct {
    mu        sync.RWMutex
    functions map[string]Function
}

func NewFunctionRegistry() *FunctionRegistry {
    return &FunctionRegistry{
        functions: make(map[string]Function),
    }
}

func (fr *FunctionRegistry) Register(fn Function) {
    fr.mu.Lock()
    defer fr.mu.Unlock()
    fr.functions[strings.ToUpper(fn.Name())] = fn
}

func (fr *FunctionRegistry) Get(name string) Function {
    fr.mu.RLock()
    defer fr.mu.RUnlock()
    return fr.functions[strings.ToUpper(name)]
}

func (fr *FunctionRegistry) Has(name string) bool {
    fr.mu.RLock()
    defer fr.mu.RUnlock()
    _, exists := fr.functions[strings.ToUpper(name)]
    return exists
}

// DefaultFunctionRegistry returns a registry with all built-in functions
func DefaultFunctionRegistry() *FunctionRegistry {
    registry := NewFunctionRegistry()

    // Math functions
    registry.Register(&SumFunction{})
    registry.Register(&AverageFunction{})
    registry.Register(&CountFunction{})
    registry.Register(&CountAFunction{})
    registry.Register(&MinFunction{})
    registry.Register(&MaxFunction{})
    registry.Register(&RoundFunction{})
    registry.Register(&RoundUpFunction{})
    registry.Register(&RoundDownFunction{})
    registry.Register(&AbsFunction{})
    registry.Register(&SqrtFunction{})
    registry.Register(&PowerFunction{})
    registry.Register(&ModFunction{})
    registry.Register(&ProductFunction{})
    registry.Register(&SumIfFunction{})
    registry.Register(&CountIfFunction{})
    registry.Register(&AverageIfFunction{})

    // Logical functions
    registry.Register(&IfFunction{})
    registry.Register(&AndFunction{})
    registry.Register(&OrFunction{})
    registry.Register(&NotFunction{})
    registry.Register(&IfErrorFunction{})
    registry.Register(&IfNAFunction{})
    registry.Register(&IfsFunction{})

    // Text functions
    registry.Register(&ConcatenateFunction{})
    registry.Register(&LeftFunction{})
    registry.Register(&RightFunction{})
    registry.Register(&MidFunction{})
    registry.Register(&LenFunction{})
    registry.Register(&UpperFunction{})
    registry.Register(&LowerFunction{})
    registry.Register(&TrimFunction{})
    registry.Register(&FindFunction{})
    registry.Register(&SubstituteFunction{})
    registry.Register(&TextFunction{})

    // Lookup functions
    registry.Register(&VLookupFunction{})
    registry.Register(&HLookupFunction{})
    registry.Register(&IndexFunction{})
    registry.Register(&MatchFunction{})
    registry.Register(&XLookupFunction{})
    registry.Register(&ChooseFunction{})

    // Date/Time functions
    registry.Register(&TodayFunction{})
    registry.Register(&NowFunction{})
    registry.Register(&DateFunction{})
    registry.Register(&TimeFunction{})
    registry.Register(&YearFunction{})
    registry.Register(&MonthFunction{})
    registry.Register(&DayFunction{})
    registry.Register(&HourFunction{})
    registry.Register(&MinuteFunction{})
    registry.Register(&SecondFunction{})
    registry.Register(&DateDifFunction{})

    // Statistical functions
    registry.Register(&MedianFunction{})
    registry.Register(&ModeFunction{})
    registry.Register(&StDevFunction{})
    registry.Register(&VarFunction{})
    registry.Register(&PercentileFunction{})
    registry.Register(&QuartileFunction{})

    return registry
}
```

---

## Built-in Functions

### Mathematical Functions

#### SUM Function

```go
// spreadsheet/formula/functions/math.go

package functions

import (
    "github.com/unidoc/goffice/spreadsheet/formula"
)

// SumFunction implements SUM(number1, [number2], ...)
type SumFunction struct {
    BaseFunction
}

func init() {
    // Constructor
}

func (f *SumFunction) Name() string {
    return "SUM"
}

func (f *SumFunction) MinArgs() int {
    return 1
}

func (f *SumFunction) MaxArgs() int {
    return -1  // Unlimited
}

func (f *SumFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    sum := 0.0

    for _, arg := range args {
        if arg.Type == formula.ValueTypeError {
            return arg, nil  // Propagate error
        }

        if arg.Type == formula.ValueTypeArray {
            // Flatten array and sum all numbers
            arr := arg.Value.([][]formula.Value)
            for _, row := range arr {
                for _, cell := range row {
                    if cell.Type == formula.ValueTypeNumber {
                        sum += cell.Value.(float64)
                    }
                    // Ignore non-numeric values
                }
            }
        } else if arg.Type == formula.ValueTypeNumber {
            sum += arg.Value.(float64)
        }
        // Ignore strings, booleans, empty
    }

    return formula.NumberValue(sum), nil
}

// AverageFunction implements AVERAGE(number1, [number2], ...)
type AverageFunction struct {
    BaseFunction
}

func (f *AverageFunction) Name() string {
    return "AVERAGE"
}

func (f *AverageFunction) MinArgs() int {
    return 1
}

func (f *AverageFunction) MaxArgs() int {
    return -1
}

func (f *AverageFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    sum := 0.0
    count := 0

    for _, arg := range args {
        if arg.Type == formula.ValueTypeError {
            return arg, nil
        }

        if arg.Type == formula.ValueTypeArray {
            arr := arg.Value.([][]formula.Value)
            for _, row := range arr {
                for _, cell := range row {
                    if cell.Type == formula.ValueTypeNumber {
                        sum += cell.Value.(float64)
                        count++
                    }
                }
            }
        } else if arg.Type == formula.ValueTypeNumber {
            sum += arg.Value.(float64)
            count++
        }
    }

    if count == 0 {
        return formula.ErrorValue(formula.ErrorDiv0), nil
    }

    return formula.NumberValue(sum / float64(count)), nil
}

// COUNT, COUNTA, MIN, MAX, ROUND, etc. follow similar pattern
// Full implementations in math.go
```

### Logical Functions

#### IF Function

```go
// IfFunction implements IF(condition, value_if_true, value_if_false)
type IfFunction struct {
    BaseFunction
}

func (f *IfFunction) Name() string {
    return "IF"
}

func (f *IfFunction) MinArgs() int {
    return 2
}

func (f *IfFunction) MaxArgs() int {
    return 3
}

func (f *IfFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    condition := args[0]
    valueIfTrue := args[1]
    valueIfFalse := formula.BooleanValue(false)
    if len(args) == 3 {
        valueIfFalse = args[2]
    }

    // Convert condition to boolean
    condBool, err := toBoolean(condition)
    if err != nil {
        return formula.ErrorValue(formula.ErrorValue), nil
    }

    if condBool.Value.(bool) {
        return valueIfTrue, nil
    }
    return valueIfFalse, nil
}

// AND, OR, NOT, IFERROR, etc. follow similar pattern
```

### Text Functions

#### CONCATENATE Function

```go
// ConcatenateFunction implements CONCATENATE(text1, [text2], ...)
type ConcatenateFunction struct {
    BaseFunction
}

func (f *ConcatenateFunction) Name() string {
    return "CONCATENATE"
}

func (f *ConcatenateFunction) MinArgs() int {
    return 1
}

func (f *ConcatenateFunction) MaxArgs() int {
    return -1
}

func (f *ConcatenateFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    result := ""

    for _, arg := range args {
        if arg.Type == formula.ValueTypeError {
            return arg, nil
        }

        result += toString(arg)
    }

    return formula.StringValue(result), nil
}

// LEFT, RIGHT, MID, LEN, UPPER, LOWER, etc. follow similar pattern
```

### Lookup Functions

#### VLOOKUP Function

```go
// VLookupFunction implements VLOOKUP(lookup_value, table_array, col_index_num, [range_lookup])
type VLookupFunction struct {
    BaseFunction
}

func (f *VLookupFunction) Name() string {
    return "VLOOKUP"
}

func (f *VLookupFunction) MinArgs() int {
    return 3
}

func (f *VLookupFunction) MaxArgs() int {
    return 4
}

func (f *VLookupFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    lookupValue := args[0]
    tableArray := args[1]
    colIndexNum := args[2]
    rangeLookup := formula.BooleanValue(true)  // Default: approximate match
    if len(args) == 4 {
        rangeLookup = args[3]
    }

    // Validate table array
    if tableArray.Type != formula.ValueTypeArray {
        return formula.ErrorValue(formula.ErrorRef), nil
    }

    // Validate column index
    if colIndexNum.Type != formula.ValueTypeNumber {
        return formula.ErrorValue(formula.ErrorValue), nil
    }
    colIdx := int(colIndexNum.Value.(float64))
    if colIdx < 1 {
        return formula.ErrorValue(formula.ErrorValue), nil
    }

    table := tableArray.Value.([][]formula.Value)
    if len(table) == 0 {
        return formula.ErrorValue(formula.ErrorNA), nil
    }
    if colIdx > len(table[0]) {
        return formula.ErrorValue(formula.ErrorRef), nil
    }

    // Search for lookup value in first column
    if rangeLookup.Value.(bool) {
        // Approximate match: binary search for largest value <= lookupValue
        // Table MUST be sorted in ascending order
        rowIdx := binarySearchVLookup(table, lookupValue)
        if rowIdx < 0 {
            return formula.ErrorValue(formula.ErrorNA), nil
        }
        if colIdx-1 < len(table[rowIdx]) {
            return table[rowIdx][colIdx-1], nil
        }
        return formula.ErrorValue(formula.ErrorRef), nil
    }

    // Exact match: linear search
    for _, row := range table {
        if len(row) == 0 {
            continue
        }

        firstCell := row[0]
        if exactMatch(firstCell, lookupValue) {
            // Found match, return value from specified column
            if colIdx-1 < len(row) {
                return row[colIdx-1], nil
            }
            return formula.ErrorValue(formula.ErrorRef), nil
        }
    }

    return formula.ErrorValue(formula.ErrorNA), nil
}

// binarySearchVLookup performs binary search for approximate match VLOOKUP
// Returns index of largest value <= lookupValue, or -1 if not found
func binarySearchVLookup(table [][]formula.Value, lookupValue formula.Value) int {
    if len(table) == 0 {
        return -1
    }

    left := 0
    right := len(table) - 1
    result := -1

    for left <= right {
        mid := (left + right) / 2
        if len(table[mid]) == 0 {
            right = mid - 1
            continue
        }

        midValue := table[mid][0]
        cmp := compareForApprox(midValue, lookupValue)

        if cmp == 0 {
            // Exact match found
            return mid
        } else if cmp < 0 {
            // midValue < lookupValue, could be result
            result = mid
            left = mid + 1
        } else {
            // midValue > lookupValue, search left
            right = mid - 1
        }
    }

    return result
}

// compareForApprox compares values for approximate match
// Returns -1 if a < b, 0 if a == b, 1 if a > b
// Type ordering: numbers < strings < booleans < errors
// Empty values are treated as 0 for numbers
func compareForApprox(a, b formula.Value) int {
    // Handle empty values - treat as 0
    if a.Type == formula.ValueTypeEmpty {
        a = formula.NumberValue(0.0)
    }
    if b.Type == formula.ValueTypeEmpty {
        b = formula.NumberValue(0.0)
    }

    // Type priority (numbers are lowest, then strings, then booleans, errors highest)
    typePriority := func(v formula.Value) int {
        switch v.Type {
        case formula.ValueTypeNumber:
            return 0
        case formula.ValueTypeString:
            return 1
        case formula.ValueTypeBoolean:
            return 2
        case formula.ValueTypeError:
            return 4  // Errors sort last
        case formula.ValueTypeArray:
            return 5  // Arrays not supported in comparison
        default:
            return 6
        }
    }

    aPrio := typePriority(a)
    bPrio := typePriority(b)

    if aPrio != bPrio {
        if aPrio < bPrio {
            return -1
        }
        return 1
    }

    // Same type - compare values
    switch a.Type {
    case formula.ValueTypeNumber:
        aNum := a.Value.(float64)
        bNum := b.Value.(float64)
        if aNum < bNum {
            return -1
        } else if aNum > bNum {
            return 1
        }
        return 0
    case formula.ValueTypeString:
        return strings.Compare(strings.ToLower(a.Value.(string)), strings.ToLower(b.Value.(string)))
    case formula.ValueTypeBoolean:
        aBool := a.Value.(bool)
        bBool := b.Value.(bool)
        if !aBool && bBool {
            return -1
        } else if aBool && !bBool {
            return 1
        }
        return 0
    case formula.ValueTypeError:
        // Errors are equal to each other
        return 0
    default:
        return 0
    }
}

// exactMatch checks for exact value equality
func exactMatch(a, b formula.Value) bool {
    if a.Type != b.Type {
        return false
    }

    switch a.Type {
    case formula.ValueTypeNumber:
        return a.Value.(float64) == b.Value.(float64)
    case formula.ValueTypeString:
        return strings.EqualFold(a.Value.(string), b.Value.(string))
    case formula.ValueTypeBoolean:
        return a.Value.(bool) == b.Value.(bool)
    default:
        return false
    }
}

// HLOOKUP, INDEX, MATCH, XLOOKUP, CHOOSE follow similar pattern
```

### Date/Time Functions

#### TODAY Function

```go
// TodayFunction implements TODAY()
type TodayFunction struct {
    VolatileBaseFunction
}

func (f *TodayFunction) Name() string {
    return "TODAY"
}

func (f *TodayFunction) MinArgs() int {
    return 0
}

func (f *TodayFunction) MaxArgs() int {
    return 0
}

func (f *TodayFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    // Excel date serial number: days since 1900-01-01
    now := time.Now()
    excelDate := excelDateFromTime(now)
    return formula.NumberValue(excelDate), nil
}

func excelDateFromTime(t time.Time) float64 {
    // Excel epoch: 1900-01-01 (serial 1)
    //
    // EXCEL BUG COMPATIBILITY: Excel incorrectly treats 1900 as a leap year.
    // In reality, 1900 is NOT a leap year (not divisible by 400).
    // However, Excel includes February 29, 1900 as serial number 60.
    //
    // For compatibility with Excel:
    // - Serials 1-59: Jan 1, 1900 - Feb 28, 1900 (correct)
    // - Serial 60: Feb 29, 1900 (INCORRECT - this day never existed!)
    // - Serials 61+: Mar 1, 1900 onwards (off by 1 day from reality)
    //
    // To match Excel, we use epoch of 1899-12-30 (serial 0) and add 1 to dates
    // on or after March 1, 1900 to account for the phantom Feb 29.
    epoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
    duration := t.Sub(epoch)
    days := duration.Hours() / 24

    // Add 1 for dates on/after March 1, 1900 to match Excel's bug
    mar1_1900 := time.Date(1900, 3, 1, 0, 0, 0, 0, time.UTC)
    if t.After(mar1_1900) || t.Equal(mar1_1900) {
        days += 1
    }

    return days
}

// NOW, DATE, TIME, YEAR, MONTH, DAY, HOUR, MINUTE, SECOND, DATEDIF follow similar pattern
```

### Statistical Functions

#### MEDIAN Function

```go
// MedianFunction implements MEDIAN(number1, [number2], ...)
type MedianFunction struct {
    BaseFunction
}

func (f *MedianFunction) Name() string {
    return "MEDIAN"
}

func (f *MedianFunction) MinArgs() int {
    return 1
}

func (f *MedianFunction) MaxArgs() int {
    return -1
}

func (f *MedianFunction) Call(args []formula.Value, ctx *formula.EvalContext) (formula.Value, error) {
    var numbers []float64

    for _, arg := range args {
        if arg.Type == formula.ValueTypeError {
            return arg, nil
        }

        if arg.Type == formula.ValueTypeArray {
            arr := arg.Value.([][]formula.Value)
            for _, row := range arr {
                for _, cell := range row {
                    if cell.Type == formula.ValueTypeNumber {
                        numbers = append(numbers, cell.Value.(float64))
                    }
                }
            }
        } else if arg.Type == formula.ValueTypeNumber {
            numbers = append(numbers, arg.Value.(float64))
        }
    }

    if len(numbers) == 0 {
        return formula.ErrorValue(formula.ErrorNum), nil
    }

    // Sort numbers
    sort.Float64s(numbers)

    // Calculate median
    n := len(numbers)
    if n%2 == 1 {
        // Odd count: middle value
        return formula.NumberValue(numbers[n/2]), nil
    }
    // Even count: average of two middle values
    median := (numbers[n/2-1] + numbers[n/2]) / 2
    return formula.NumberValue(median), nil
}

// MODE, STDEV, VAR, PERCENTILE, QUARTILE follow similar pattern
```

---

## Array Formula Support

### Array Formula Evaluation

Array formulas operate on ranges and produce multiple results that fill a range of cells.

```go
// spreadsheet/formula/array.go

package formula

// ArrayFormula represents an array formula that produces multiple values
type ArrayFormula struct {
    Expression Expression
    Range      RangeRef  // Target range for results
}

// Evaluate evaluates the array formula and returns 2D array of results
func (af *ArrayFormula) Evaluate(ctx *EvalContext) ([][]Value, error) {
    // Evaluate the expression (may return array or single value)
    result, err := af.Expression.Evaluate(ctx)
    if err != nil {
        return nil, err
    }

    // If result is already an array, use it
    if result.Type == ValueTypeArray {
        resultArray := result.Value.([][]Value)
        return af.distributeResults(resultArray), nil
    }

    // Single value - broadcast to entire range
    rows := af.Range.EndRow - af.Range.StartRow + 1
    cols := af.Range.EndCol - af.Range.StartCol + 1

    array := make([][]Value, rows)
    for i := 0; i < rows; i++ {
        array[i] = make([]Value, cols)
        for j := 0; j < cols; j++ {
            array[i][j] = result
        }
    }

    return array, nil
}

// distributeResults distributes array results across target range
// Handles size mismatches by filling with #N/A
func (af *ArrayFormula) distributeResults(results [][]Value) [][]Value {
    targetRows := af.Range.EndRow - af.Range.StartRow + 1
    targetCols := af.Range.EndCol - af.Range.StartCol + 1

    output := make([][]Value, targetRows)
    for i := 0; i < targetRows; i++ {
        output[i] = make([]Value, targetCols)
        for j := 0; j < targetCols; j++ {
            if i < len(results) && j < len(results[i]) {
                // Result available
                output[i][j] = results[i][j]
            } else {
                // Size mismatch - fill with #N/A
                output[i][j] = ErrorValue(ErrorNA)
            }
        }
    }

    return output
}

// SetArrayFormula sets an array formula in a range
func (ws *Worksheet) SetArrayFormula(rangeRef string, formula string) error {
    // Parse range
    rng, err := ParseRangeRef(rangeRef)
    if err != nil {
        return err
    }

    // Parse formula
    expr, err := Parse(formula)
    if err != nil {
        return err
    }

    // Create array formula
    af := &ArrayFormula{
        Expression: expr,
        Range:      *rng,
    }

    // Evaluate to get results
    ctx := &EvalContext{
        Workbook:         ws.Workbook(),
        CurrentSheet:     ws,
        FunctionRegistry: DefaultFunctionRegistry(),
        Resolver:         NewWorkbookResolver(ws.Workbook()),
    }

    results, err := af.Evaluate(ctx)
    if err != nil {
        return err
    }

    // Set formula in all cells of range
    // First cell (top-left) contains the formula
    // Other cells reference the first cell as master
    for i := 0; i < len(results); i++ {
        for j := 0; j < len(results[i]); j++ {
            row := rng.StartRow + i
            col := rng.StartCol + j

            cell := ws.Cell(row, col)

            if i == 0 && j == 0 {
                // Master cell - set formula
                cell.SetFormula(formula)
                cell.SetArrayFormula(true)
                cell.SetArrayRange(rangeRef)
            } else {
                // Slave cell - reference master
                cell.SetArrayFormulaMaster(rng.StartCol, rng.StartRow)
            }

            // Set computed value
            result := results[i][j]
            switch result.Type {
            case ValueTypeNumber:
                cell.SetNumber(result.Value.(float64))
            case ValueTypeString:
                cell.SetString(result.Value.(string))
            case ValueTypeBoolean:
                cell.SetBoolean(result.Value.(bool))
            case ValueTypeError:
                cell.SetError(result.Value.(string))
            }
        }
    }

    return nil
}
```

### Array Operations

Binary operations on arrays perform element-wise operations. See BinaryOp.Evaluate() in AST section (lines 341-396) for array handling logic.

```go
// evaluateArrayBinaryOp handles array-array and array-scalar operations
func evaluateArrayBinaryOp(left, right Value, op TokenType) (Value, error) {
    // If both are arrays, perform element-wise operation
    if left.Type == ValueTypeArray && right.Type == ValueTypeArray {
        leftArray := left.Value.([][]Value)
        rightArray := right.Value.([][]Value)

        // Arrays must have same dimensions
        if len(leftArray) != len(rightArray) {
            return ErrorValue("#VALUE!"), nil
        }

        result := make([][]Value, len(leftArray))
        for i := 0; i < len(leftArray); i++ {
            if len(leftArray[i]) != len(rightArray[i]) {
                return ErrorValue("#VALUE!"), nil
            }

            result[i] = make([]Value, len(leftArray[i]))
            for j := 0; j < len(leftArray[i]); j++ {
                // Apply operator element-wise
                switch op {
                case TokenPlus:
                    result[i][j], _ = add(leftArray[i][j], rightArray[i][j])
                case TokenMinus:
                    result[i][j], _ = subtract(leftArray[i][j], rightArray[i][j])
                case TokenMult:
                    result[i][j], _ = multiply(leftArray[i][j], rightArray[i][j])
                case TokenDiv:
                    result[i][j], _ = divide(leftArray[i][j], rightArray[i][j])
                case TokenPower:
                    result[i][j], _ = power(leftArray[i][j], rightArray[i][j])
                case TokenConcat:
                    result[i][j], _ = concatenate(leftArray[i][j], rightArray[i][j])
                case TokenEQ:
                    result[i][j], _ = equal(leftArray[i][j], rightArray[i][j])
                case TokenNE:
                    result[i][j], _ = notEqual(leftArray[i][j], rightArray[i][j])
                case TokenLT:
                    result[i][j], _ = lessThan(leftArray[i][j], rightArray[i][j])
                case TokenLE:
                    result[i][j], _ = lessEqual(leftArray[i][j], rightArray[i][j])
                case TokenGT:
                    result[i][j], _ = greaterThan(leftArray[i][j], rightArray[i][j])
                case TokenGE:
                    result[i][j], _ = greaterEqual(leftArray[i][j], rightArray[i][j])
                default:
                    result[i][j] = ErrorValue("#VALUE!")
                }
            }
        }

        return ArrayValue(result), nil
    }

    // If one is array and one is scalar, broadcast scalar
    if left.Type == ValueTypeArray {
        return evaluateArrayScalarOp(left, right, op, false)
    }
    if right.Type == ValueTypeArray {
        return evaluateArrayScalarOp(right, left, op, true)
    }

    // Should not reach here - caller checks for arrays
    return ErrorValue("#VALUE!"), nil
}

func evaluateArrayScalarOp(arrayVal, scalarVal Value, op TokenType, reverse bool) (Value, error) {
    array := arrayVal.Value.([][]Value)
    result := make([][]Value, len(array))

    for i := 0; i < len(array); i++ {
        result[i] = make([]Value, len(array[i]))
        for j := 0; j < len(array[i]); j++ {
            // Apply operator: array[i][j] op scalar (or reverse)
            var val Value
            var err error

            if reverse {
                // scalar op array[i][j]
                switch op {
                case TokenPlus:
                    val, err = add(scalarVal, array[i][j])
                case TokenMinus:
                    val, err = subtract(scalarVal, array[i][j])
                case TokenMult:
                    val, err = multiply(scalarVal, array[i][j])
                case TokenDiv:
                    val, err = divide(scalarVal, array[i][j])
                case TokenPower:
                    val, err = power(scalarVal, array[i][j])
                case TokenConcat:
                    val, err = concatenate(scalarVal, array[i][j])
                case TokenEQ:
                    val, err = equal(scalarVal, array[i][j])
                case TokenNE:
                    val, err = notEqual(scalarVal, array[i][j])
                case TokenLT:
                    val, err = lessThan(scalarVal, array[i][j])
                case TokenLE:
                    val, err = lessEqual(scalarVal, array[i][j])
                case TokenGT:
                    val, err = greaterThan(scalarVal, array[i][j])
                case TokenGE:
                    val, err = greaterEqual(scalarVal, array[i][j])
                default:
                    val = ErrorValue("#VALUE!")
                }
            } else {
                // array[i][j] op scalar
                switch op {
                case TokenPlus:
                    val, err = add(array[i][j], scalarVal)
                case TokenMinus:
                    val, err = subtract(array[i][j], scalarVal)
                case TokenMult:
                    val, err = multiply(array[i][j], scalarVal)
                case TokenDiv:
                    val, err = divide(array[i][j], scalarVal)
                case TokenPower:
                    val, err = power(array[i][j], scalarVal)
                case TokenConcat:
                    val, err = concatenate(array[i][j], scalarVal)
                case TokenEQ:
                    val, err = equal(array[i][j], scalarVal)
                case TokenNE:
                    val, err = notEqual(array[i][j], scalarVal)
                case TokenLT:
                    val, err = lessThan(array[i][j], scalarVal)
                case TokenLE:
                    val, err = lessEqual(array[i][j], scalarVal)
                case TokenGT:
                    val, err = greaterThan(array[i][j], scalarVal)
                case TokenGE:
                    val, err = greaterEqual(array[i][j], scalarVal)
                default:
                    val = ErrorValue("#VALUE!")
                }
            }

            if err != nil {
                result[i][j] = ErrorValue("#VALUE!")
            } else {
                result[i][j] = val
            }
        }
    }

    return ArrayValue(result), nil
}
```

---

## Dependency Tracking

### Dependency Graph

```go
// spreadsheet/formula/dependency.go

package formula

import (
    "fmt"
    "github.com/unidoc/goffice/spreadsheet"
)

// DependencyGraph tracks formula dependencies for recalculation order
type DependencyGraph struct {
    nodes map[CellKey]*DependencyNode
    edges map[CellKey][]CellKey  // Cell -> Cells that depend on it (dependents)
}

// CellKey uniquely identifies a cell
type CellKey struct {
    SheetName string
    Col       int
    Row       int
}

func NewCellKey(sheetName string, col, row int) CellKey {
    return CellKey{SheetName: sheetName, Col: col, Row: row}
}

// DependencyNode represents a cell with formula
type DependencyNode struct {
    Cell         CellKey
    Formula      Expression
    Dependencies []CellKey  // Cells this formula depends on
    Dependents   []CellKey  // Cells that depend on this cell
}

func NewDependencyGraph() *DependencyGraph {
    return &DependencyGraph{
        nodes: make(map[CellKey]*DependencyNode),
        edges: make(map[CellKey][]CellKey),
    }
}

// AddFormula adds a formula cell to the dependency graph
func (g *DependencyGraph) AddFormula(cell CellKey, formula Expression) {
    // Extract dependencies from formula
    deps := extractDependencies(formula)

    node := &DependencyNode{
        Cell:         cell,
        Formula:      formula,
        Dependencies: deps,
    }

    g.nodes[cell] = node

    // Update edges: for each dependency, add this cell as a dependent
    for _, dep := range deps {
        g.edges[dep] = append(g.edges[dep], cell)
    }

    // Update dependent lists
    for _, dep := range deps {
        if depNode, exists := g.nodes[dep]; exists {
            depNode.Dependents = append(depNode.Dependents, cell)
        }
    }
}

// RemoveFormula removes a formula cell from the graph
func (g *DependencyGraph) RemoveFormula(cell CellKey) {
    node, exists := g.nodes[cell]
    if !exists {
        return
    }

    // Remove from dependency edges
    for _, dep := range node.Dependencies {
        g.edges[dep] = removeFromSlice(g.edges[dep], cell)
    }

    // Remove node
    delete(g.nodes, cell)
    delete(g.edges, cell)
}

// TopologicalSort returns cells in dependency order (Kahn's algorithm)
func (g *DependencyGraph) TopologicalSort() ([]CellKey, error) {
    // Build in-degree map
    inDegree := make(map[CellKey]int)
    for cell := range g.nodes {
        inDegree[cell] = 0
    }

    for _, node := range g.nodes {
        for _, dep := range node.Dependencies {
            inDegree[node.Cell]++
        }
    }

    // Queue of cells with no dependencies
    queue := make([]CellKey, 0)
    for cell, degree := range inDegree {
        if degree == 0 {
            queue = append(queue, cell)
        }
    }

    // Process queue
    result := make([]CellKey, 0)
    for len(queue) > 0 {
        // Dequeue
        current := queue[0]
        queue = queue[1:]

        result = append(result, current)

        // Reduce in-degree of dependents
        for _, dependent := range g.edges[current] {
            inDegree[dependent]--
            if inDegree[dependent] == 0 {
                queue = append(queue, dependent)
            }
        }
    }

    // Check for circular references
    if len(result) != len(g.nodes) {
        // Circular reference detected
        cycle := g.findCycle()
        return nil, fmt.Errorf("circular reference detected: %v", cycle)
    }

    return result, nil
}

// findCycle finds a circular reference using DFS with path tracking
func (g *DependencyGraph) findCycle() []CellKey {
    visited := make(map[CellKey]bool)
    recStack := make(map[CellKey]bool)

    var dfs func(cell CellKey, path []CellKey) []CellKey

    dfs = func(cell CellKey, path []CellKey) []CellKey {
        visited[cell] = true
        recStack[cell] = true
        path = append(path, cell)

        for _, dependent := range g.edges[cell] {
            if !visited[dependent] {
                if cycle := dfs(dependent, path); cycle != nil {
                    return cycle
                }
            } else if recStack[dependent] {
                // Cycle found - extract cycle from path
                cycleStart := -1
                for i, c := range path {
                    if c == dependent {
                        cycleStart = i
                        break
                    }
                }
                if cycleStart >= 0 {
                    // Create cycle: cells from cycleStart to end + dependent to close
                    cycle := make([]CellKey, 0, len(path)-cycleStart+1)
                    cycle = append(cycle, path[cycleStart:]...)
                    cycle = append(cycle, dependent) // Close the cycle
                    return cycle
                }
            }
        }

        recStack[cell] = false
        return nil
    }

    for cell := range g.nodes {
        if !visited[cell] {
            if cycle := dfs(cell, make([]CellKey, 0)); cycle != nil {
                return cycle
            }
        }
    }

    return nil
}

// GetDependents returns all cells that depend on the given cell
func (g *DependencyGraph) GetDependents(cell CellKey) []CellKey {
    return g.edges[cell]
}

// Helper: Extract dependencies from expression
func extractDependencies(expr Expression) []CellKey {
    deps := expr.Dependencies()
    keys := make([]CellKey, 0)

    for _, dep := range deps {
        switch dep.Type {
        case DependencyTypeCell:
            ref := dep.CellRef
            keys = append(keys, NewCellKey(ref.SheetName, ref.Col, ref.Row))
        case DependencyTypeRange:
            ref := dep.RangeRef
            // Add all cells in range
            for row := ref.StartRow; row <= ref.EndRow; row++ {
                for col := ref.StartCol; col <= ref.EndCol; col++ {
                    keys = append(keys, NewCellKey(ref.SheetName, col, row))
                }
            }
        }
    }

    return keys
}

func removeFromSlice(slice []CellKey, item CellKey) []CellKey {
    result := make([]CellKey, 0)
    for _, v := range slice {
        if v != item {
            result = append(result, v)
        }
    }
    return result
}
```

---

## Recalculation Engine

### Workbook Recalculation

```go
// spreadsheet/formula/recalc.go

package formula

import (
    "github.com/unidoc/goffice/spreadsheet"
)

// RecalculateWorkbook recalculates all formulas in the workbook
func RecalculateWorkbook(wb *spreadsheet.Workbook) error {
    // 1. Build dependency graph
    graph := NewDependencyGraph()

    for _, sheet := range wb.Worksheets() {
        for _, row := range sheet.Rows() {
            for _, cell := range row.Cells() {
                formulaText := cell.GetFormula()
                if formulaText == "" {
                    continue
                }

                // Parse formula
                expr, err := Parse(formulaText)
                if err != nil {
                    // Invalid formula, set error
                    cell.SetError(ErrorValue)
                    continue
                }

                // Add to dependency graph
                cellKey := NewCellKey(sheet.Name(), cell.Column(), cell.Row())
                graph.AddFormula(cellKey, expr)
            }
        }
    }

    // 2. Topological sort to get calculation order
    order, err := graph.TopologicalSort()
    if err != nil {
        return err  // Circular reference
    }

    // 3. Create evaluation context
    ctx := &EvalContext{
        Workbook:         wb,
        FunctionRegistry: DefaultFunctionRegistry(),
        Resolver:         NewWorkbookResolver(wb),
    }

    // 4. Evaluate formulas in order
    for _, cellKey := range order {
        sheet := wb.SheetByName(cellKey.SheetName)
        if sheet == nil {
            continue
        }

        cell := sheet.Cell(cellKey.Row, cellKey.Col)
        if cell == nil {
            continue
        }

        node := graph.nodes[cellKey]
        if node == nil {
            continue
        }

        // Evaluate formula
        // NOTE: When caching is implemented (lazy evaluation), check if formula contains
        // volatile functions. If so, always re-evaluate regardless of cache status.
        // Volatile functions: NOW(), TODAY(), RAND(), RANDBETWEEN() - always recalculate
        //
        // Example cache check:
        //   if !containsVolatileFunction(node.Formula) && cache.Has(cellKey) {
        //       value = cache.Get(cellKey)
        //   } else {
        //       value, err = node.Formula.Evaluate(ctx)
        //       cache.Set(cellKey, value)
        //   }
        value, err := node.Formula.Evaluate(ctx)
        if err != nil {
            cell.SetError(ErrorValue)
            continue
        }

        // Set cell value based on result type
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

// RecalculateWorksheet recalculates all formulas in a worksheet
func RecalculateWorksheet(ws *spreadsheet.Worksheet) error {
    // Build dependency graph for formulas in this worksheet only
    graph := NewDependencyGraph()

    for _, row := range ws.Rows() {
        for _, cell := range row.Cells() {
            formulaText := cell.GetFormula()
            if formulaText == "" {
                continue
            }

            expr, err := Parse(formulaText)
            if err != nil {
                cell.SetError("#VALUE!")
                continue
            }

            cellKey := NewCellKey(ws.Name(), cell.Column(), cell.Row())
            graph.AddFormula(cellKey, expr)
        }
    }

    // Topological sort
    order, err := graph.TopologicalSort()
    if err != nil {
        return err  // Circular reference
    }

    // Create evaluation context
    ctx := &EvalContext{
        Workbook:         ws.Workbook(),
        CurrentSheet:     ws,
        FunctionRegistry: DefaultFunctionRegistry(),
        Resolver:         NewWorkbookResolver(ws.Workbook()),
    }

    // Evaluate in order
    for _, cellKey := range order {
        cell := ws.Cell(cellKey.Row, cellKey.Col)
        if cell == nil {
            continue
        }

        node := graph.nodes[cellKey]
        if node == nil {
            continue
        }

        value, err := node.Formula.Evaluate(ctx)
        if err != nil {
            cell.SetError("#VALUE!")
            continue
        }

        // Set cell value
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

// RecalculateCell recalculates a single cell and its dependents
func RecalculateCell(wb *spreadsheet.Workbook, sheetName string, col, row int) error {
    // 1. Build full dependency graph (need all sheets to find dependents)
    graph := NewDependencyGraph()

    for _, sheet := range wb.Worksheets() {
        for _, rowIter := range sheet.Rows() {
            for _, cell := range rowIter.Cells() {
                formulaText := cell.GetFormula()
                if formulaText == "" {
                    continue
                }

                expr, err := Parse(formulaText)
                if err != nil {
                    continue
                }

                cellKey := NewCellKey(sheet.Name(), cell.Column(), cell.Row())
                graph.AddFormula(cellKey, expr)
            }
        }
    }

    // 2. Find cell and its dependents
    cellKey := NewCellKey(sheetName, col, row)
    affectedCells := []CellKey{cellKey}
    affectedCells = append(affectedCells, graph.GetDependents(cellKey)...)

    // 3. Topological sort of affected cells
    subgraph := buildSubgraph(graph, affectedCells)
    order, err := subgraph.TopologicalSort()
    if err != nil {
        return err
    }

    // 4. Evaluate affected cells
    ctx := &EvalContext{
        Workbook:         wb,
        FunctionRegistry: DefaultFunctionRegistry(),
        Resolver:         NewWorkbookResolver(wb),
    }

    for _, key := range order {
        // Evaluate and update cell
        sheet := wb.SheetByName(key.SheetName)
        if sheet == nil {
            continue
        }

        cell := sheet.Cell(key.Row, key.Col)
        if cell == nil {
            continue
        }

        node := graph.nodes[key]
        if node == nil {
            continue
        }

        // Set current sheet context
        ctx.CurrentSheet = sheet

        // Evaluate formula
        value, err := node.Formula.Evaluate(ctx)
        if err != nil {
            cell.SetError("#VALUE!")
            continue
        }

        // Set cell value based on result type
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

func buildSubgraph(full *DependencyGraph, cells []CellKey) *DependencyGraph {
    sub := NewDependencyGraph()
    for _, cell := range cells {
        if node, exists := full.nodes[cell]; exists {
            sub.AddFormula(cell, node.Formula)
        }
    }
    return sub
}
```

### Calculation Chain Builder

```go
// spreadsheet/formula/calc_chain.go

package formula

import (
    "github.com/unidoc/goffice/spreadsheet"
)

// RebuildCalculationChain rebuilds the calculation chain for a workbook
func RebuildCalculationChain(wb *spreadsheet.Workbook) error {
    // 1. Build dependency graph
    graph := NewDependencyGraph()

    for _, sheet := range wb.Worksheets() {
        for _, row := range sheet.Rows() {
            for _, cell := range row.Cells() {
                formulaText := cell.GetFormula()
                if formulaText == "" {
                    continue
                }

                expr, err := Parse(formulaText)
                if err != nil {
                    continue  // Skip invalid formulas
                }

                cellKey := NewCellKey(sheet.Name(), cell.Column(), cell.Row())
                graph.AddFormula(cellKey, expr)
            }
        }
    }

    // 2. Topological sort
    order, err := graph.TopologicalSort()
    if err != nil {
        return err
    }

    // 3. Update CalcChainPart
    calcChainPart := wb.CalcChainPart()
    if calcChainPart == nil {
        // Create calc chain part if not exists
        calcChainPart = wb.AddCalcChainPart()
    }

    // Clear existing calculation chain
    calcChainPart.Clear()

    // Add cells in dependency order
    for i, cellKey := range order {
        calcChainPart.AddCalculationCell(cellKey.SheetName, cellKey.Col, cellKey.Row, i)
    }

    return nil
}
```

---

## Error Handling

### Error Propagation

All formula errors follow Excel's error propagation rules:

1. **Immediate Error Return**: If any operand is an error, the error propagates immediately
2. **Error Priority**: If multiple errors, highest priority error wins:
   - `#NULL!` - Intersection of ranges that don't intersect
   - `#DIV/0!` - Division by zero
   - `#VALUE!` - Wrong type of operand
   - `#REF!` - Invalid cell reference
   - `#NAME?` - Unrecognized function/name
   - `#NUM!` - Invalid numeric value (e.g., SQRT(-1))
   - `#N/A` - Value not available

3. **Error Codes**:
```go
const (
    ErrorNull        = "#NULL!"
    ErrorDiv0        = "#DIV/0!"
    ErrorValue       = "#VALUE!"
    ErrorRef         = "#REF!"
    ErrorName        = "#NAME?"
    ErrorNum         = "#NUM!"
    ErrorNA          = "#N/A"
    ErrorGettingData = "#GETTING_DATA"
    ErrorSpill       = "#SPILL!"
    ErrorCalc        = "#CALC!"
)
```

---

## Performance Optimizations

### 1. Lazy Evaluation

Only evaluate cells when needed:
```go
// Cache evaluated values
type CellValueCache struct {
    values map[CellKey]Value
    dirty  map[CellKey]bool
}

func (c *CellValueCache) Get(key CellKey) (Value, bool) {
    if c.dirty[key] {
        return Value{}, false
    }
    val, ok := c.values[key]
    return val, ok
}

func (c *CellValueCache) Set(key CellKey, val Value) {
    c.values[key] = val
    c.dirty[key] = false
}

func (c *CellValueCache) MarkDirty(key CellKey) {
    c.dirty[key] = true
}
```

### 2. Incremental Recalculation

Only recalculate affected cells:
```go
func (wb *Workbook) OnCellChange(sheetName string, col, row int) {
    cellKey := NewCellKey(sheetName, col, row)

    // Mark cell and dependents as dirty
    graph := wb.DependencyGraph()
    affectedCells := []CellKey{cellKey}
    affectedCells = append(affectedCells, graph.GetDependents(cellKey)...)

    for _, cell := range affectedCells {
        wb.valueCache.MarkDirty(cell)
    }

    // Recalculate only affected cells
    RecalculateCell(wb, sheetName, col, row)
}
```

### 3. Parallel Evaluation

Evaluate independent cells in parallel:
```go
func (g *DependencyGraph) ParallelEvaluate(ctx *EvalContext) error {
    levels := g.topoLevels()

    for _, level := range levels {
        // All cells in a level have no dependencies on each other
        var wg sync.WaitGroup
        for _, cellKey := range level {
            wg.Add(1)
            go func(key CellKey) {
                defer wg.Done()
                // Evaluate cell
                node := g.nodes[key]
                node.Formula.Evaluate(ctx)
            }(cellKey)
        }
        wg.Wait()
    }

    return nil
}

// topoLevels groups cells by dependency level for parallel evaluation
// Level 0: cells with no dependencies
// Level 1: cells depending only on level 0
// Level N: cells depending on levels 0..N-1
func (g *DependencyGraph) topoLevels() [][]CellKey {
    // Calculate in-degrees for each node
    inDegree := make(map[CellKey]int)
    for cell := range g.nodes {
        inDegree[cell] = 0
    }

    for _, node := range g.nodes {
        for _, dep := range node.Dependencies {
            inDegree[node.Cell]++
        }
    }

    // Build levels using modified Kahn's algorithm
    levels := make([][]CellKey, 0)
    remaining := make(map[CellKey]bool)
    for cell := range g.nodes {
        remaining[cell] = true
    }

    for len(remaining) > 0 {
        // Find all cells with in-degree 0 (current level)
        currentLevel := make([]CellKey, 0)
        for cell := range remaining {
            if inDegree[cell] == 0 {
                currentLevel = append(currentLevel, cell)
            }
        }

        if len(currentLevel) == 0 {
            // Circular reference - remaining nodes all have dependencies
            break
        }

        levels = append(levels, currentLevel)

        // Remove current level nodes and update in-degrees
        for _, cell := range currentLevel {
            delete(remaining, cell)

            // Decrease in-degree for all dependents
            node := g.nodes[cell]
            for _, dependent := range node.Dependents {
                inDegree[dependent]--
            }
        }
    }

    return levels
}
```

### 4. Memory Pooling

Reuse value objects:
```go
var valuePool = sync.Pool{
    New: func() interface{} {
        return &Value{}
    },
}

func AcquireValue() *Value {
    return valuePool.Get().(*Value)
}

func ReleaseValue(v *Value) {
    v.Type = ValueTypeEmpty
    v.Value = nil
    valuePool.Put(v)
}
```

---

## Complete Implementation Checklist

- [x] AST type system (Expression, Literal, CellReference, RangeReference, BinaryOp, UnaryOp, FunctionCall)
- [x] Parser (recursive descent, operator precedence)
- [x] Evaluator (binary ops, type conversion, error propagation)
- [x] Value type system (Number, String, Boolean, Error, Array, Empty)
- [x] Evaluation context (CellResolver, FunctionRegistry)
- [x] Function interface and registry
- [x] 60+ built-in functions (Math, Logical, Text, Lookup, Date/Time, Statistical)
- [x] Dependency graph (DependencyNode, edges, topological sort)
- [x] Circular reference detection (DFS cycle finding)
- [x] Recalculation engine (Workbook, Worksheet, Cell)
- [x] Calculation chain builder
- [x] Error handling and propagation
- [x] Performance optimizations (caching, incremental, parallel, pooling)

---

## Design Summary

This design provides a **complete, production-ready formula evaluation engine** for goffice with:

1. **Complete AST representation** with 7 expression types
2. **Recursive descent parser** with full operator precedence
3. **Comprehensive evaluator** with type coercion and error handling
4. **60+ Excel-compatible functions** across 6 categories
5. **Dependency tracking** with topological sort and circular reference detection
6. **Recalculation engine** supporting full workbook, worksheet, and incremental cell recalculation
7. **Performance optimizations** including caching, parallel evaluation, and memory pooling

**Total estimated implementation:** ~6000 lines of Go code across 15 files.
