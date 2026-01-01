# Spreadsheet Formula Spec Delta

## ADDED Requirements

### Requirement: Expression AST Interface

The system SHALL provide an Expression interface for representing formula abstract syntax trees.

#### Scenario: Expression interface definition
- GIVEN the spreadsheet/formula package
- WHEN the Expression interface is defined
- THEN the interface has method Evaluate(ctx *EvalContext) (Value, error)
- AND the interface has method Dependencies() []Dependency
- AND the interface has method String() string
- AND all formula components implement this interface

#### Scenario: Literal expression evaluation
- GIVEN a Literal expression with value 42.0
- WHEN Evaluate(ctx) is called
- THEN a NumberValue(42.0) is returned
- AND no error is returned
- AND Dependencies() returns empty slice

#### Scenario: Cell reference expression evaluation
- GIVEN a CellReference for cell A1
- AND cell A1 contains value 100
- WHEN Evaluate(ctx) is called
- THEN a NumberValue(100) is returned
- AND Dependencies() returns single cell dependency for A1

#### Scenario: Range reference expression evaluation
- GIVEN a RangeReference for range A1:A3
- AND cells contain values [10, 20, 30]
- WHEN Evaluate(ctx) is called
- THEN an ArrayValue([[10], [20], [30]]) is returned
- AND Dependencies() returns range dependency for A1:A3

#### Scenario: Binary operation expression evaluation
- GIVEN a BinaryOp with left=Literal(5), operator=Plus, right=Literal(3)
- WHEN Evaluate(ctx) is called
- THEN a NumberValue(8) is returned
- AND Dependencies() combines left and right dependencies (empty)

#### Scenario: Function call expression evaluation
- GIVEN a FunctionCall with name="SUM" and args=[RangeReference("A1:A3")]
- AND range contains values [10, 20, 30]
- WHEN Evaluate(ctx) is called
- THEN a NumberValue(60) is returned
- AND Dependencies() contains range dependency for A1:A3

### Requirement: Formula Parser

The system SHALL provide a recursive descent parser that converts formula strings to Expression AST.

#### Scenario: Parse simple number literal
- GIVEN formula string "42"
- WHEN Parse(formula) is called
- THEN a Literal expression with value 42.0 is returned
- AND no error is returned

#### Scenario: Parse simple cell reference
- GIVEN formula string "A1"
- WHEN Parse(formula) is called
- THEN a CellReference expression for A1 is returned
- AND no error is returned

#### Scenario: Parse absolute cell reference
- GIVEN formula string "$A$1"
- WHEN Parse(formula) is called
- THEN a CellReference with ColAbs=true and RowAbs=true is returned

#### Scenario: Parse range reference
- GIVEN formula string "A1:C10"
- WHEN Parse(formula) is called
- THEN a RangeReference expression is returned
- AND range has StartCol=1, StartRow=1, EndCol=3, EndRow=10

#### Scenario: Parse addition expression
- GIVEN formula string "5 + 3"
- WHEN Parse(formula) is called
- THEN a BinaryOp with operator TokenPlus is returned
- AND left operand is Literal(5)
- AND right operand is Literal(3)

#### Scenario: Parse multiplication with precedence
- GIVEN formula string "2 + 3 * 4"
- WHEN Parse(formula) is called
- THEN AST represents (2 + (3 * 4))
- AND evaluation yields 14, not 20

#### Scenario: Parse parenthesized expression
- GIVEN formula string "(2 + 3) * 4"
- WHEN Parse(formula) is called
- THEN AST represents ((2 + 3) * 4)
- AND evaluation yields 20, not 14

#### Scenario: Parse function call with no arguments
- GIVEN formula string "TODAY()"
- WHEN Parse(formula) is called
- THEN a FunctionCall with name="TODAY" and args=[] is returned

#### Scenario: Parse function call with multiple arguments
- GIVEN formula string "IF(A1>5, \"Yes\", \"No\")"
- WHEN Parse(formula) is called
- THEN a FunctionCall with name="IF" and 3 arguments is returned
- AND arg 0 is BinaryOp(CellReference, TokenGT, Literal(5))
- AND arg 1 is Literal("Yes")
- AND arg 2 is Literal("No")

#### Scenario: Parse nested function calls
- GIVEN formula string "SUM(A1:A10) + AVERAGE(B1:B10)"
- WHEN Parse(formula) is called
- THEN a BinaryOp with two FunctionCall operands is returned

#### Scenario: Parse with leading equals sign
- GIVEN formula string "=A1+B1"
- WHEN Parse(formula) is called
- THEN leading "=" is stripped
- AND BinaryOp(CellReference, TokenPlus, CellReference) is returned

#### Scenario: Parser operator precedence
- GIVEN formula string "2 ^ 3 ^ 2"
- WHEN Parse(formula) is called
- THEN AST is right-associative: (2 ^ (3 ^ 2))
- AND evaluation yields 512, not 64

#### Scenario: Parser error handling
- GIVEN formula string "A1 +"
- WHEN Parse(formula) is called
- THEN an error is returned
- AND error message indicates unexpected end of input

### Requirement: Value Type System

The system SHALL provide a Value type representing evaluated formula results.

#### Scenario: NumberValue construction
- GIVEN NumberValue(42.5)
- THEN Value.Type equals ValueTypeNumber
- AND Value.Value equals 42.5

#### Scenario: StringValue construction
- GIVEN StringValue("Hello")
- THEN Value.Type equals ValueTypeString
- AND Value.Value equals "Hello"

#### Scenario: BooleanValue construction
- GIVEN BooleanValue(true)
- THEN Value.Type equals ValueTypeBoolean
- AND Value.Value equals true

#### Scenario: ErrorValue construction
- GIVEN ErrorValue("#DIV/0!")
- THEN Value.Type equals ValueTypeError
- AND Value.Value equals "#DIV/0!"

#### Scenario: ArrayValue construction
- GIVEN ArrayValue([[1, 2], [3, 4]])
- THEN Value.Type equals ValueTypeArray
- AND Value.Value is 2D array [[1, 2], [3, 4]]

#### Scenario: EmptyValue construction
- GIVEN EmptyValue()
- THEN Value.Type equals ValueTypeEmpty
- AND Value.Value equals nil

### Requirement: Binary Operators

The system SHALL evaluate binary operations with Excel-compatible type coercion and error handling.

#### Scenario: Addition of numbers
- GIVEN left=NumberValue(5), right=NumberValue(3)
- WHEN add(left, right) is called
- THEN NumberValue(8) is returned

#### Scenario: Addition with type coercion
- GIVEN left=StringValue("5"), right=NumberValue(3)
- WHEN add(left, right) is called
- THEN NumberValue(8) is returned (string coerced to number)

#### Scenario: Division by zero
- GIVEN left=NumberValue(10), right=NumberValue(0)
- WHEN divide(left, right) is called
- THEN ErrorValue("#DIV/0!") is returned

#### Scenario: Exponentiation with negative result
- GIVEN left=NumberValue(-1), right=NumberValue(0.5)
- WHEN power(left, right) is called
- THEN ErrorValue("#NUM!") is returned (sqrt of negative)

#### Scenario: String concatenation
- GIVEN left=StringValue("Hello"), right=StringValue("World")
- WHEN concatenate(left, right) is called
- THEN StringValue("HelloWorld") is returned

#### Scenario: Equality comparison
- GIVEN left=NumberValue(5), right=NumberValue(5)
- WHEN equal(left, right) is called
- THEN BooleanValue(true) is returned

#### Scenario: String comparison case-insensitive
- GIVEN left=StringValue("hello"), right=StringValue("HELLO")
- WHEN equal(left, right) is called
- THEN BooleanValue(true) is returned (Excel is case-insensitive)

#### Scenario: Less than comparison
- GIVEN left=NumberValue(3), right=NumberValue(5)
- WHEN lessThan(left, right) is called
- THEN BooleanValue(true) is returned

### Requirement: Type Conversion

The system SHALL provide type conversion functions matching Excel behavior.

#### Scenario: String to number conversion
- GIVEN StringValue("42.5")
- WHEN toNumber(value) is called
- THEN NumberValue(42.5) is returned

#### Scenario: Boolean to number conversion
- GIVEN BooleanValue(true)
- WHEN toNumber(value) is called
- THEN NumberValue(1.0) is returned

#### Scenario: Boolean false to number conversion
- GIVEN BooleanValue(false)
- WHEN toNumber(value) is called
- THEN NumberValue(0.0) is returned

#### Scenario: Empty to number conversion
- GIVEN EmptyValue()
- WHEN toNumber(value) is called
- THEN NumberValue(0.0) is returned

#### Scenario: Invalid string to number conversion
- GIVEN StringValue("hello")
- WHEN toNumber(value) is called
- THEN an error is returned

#### Scenario: Number to string conversion
- GIVEN NumberValue(42.5)
- WHEN toString(value) is called
- THEN "42.5" is returned

#### Scenario: Boolean to string conversion
- GIVEN BooleanValue(true)
- WHEN toString(value) is called
- THEN "TRUE" is returned

#### Scenario: Error to string conversion
- GIVEN ErrorValue("#DIV/0!")
- WHEN toString(value) is called
- THEN "#DIV/0!" is returned

### Requirement: Error Propagation

The system SHALL propagate errors through formula evaluation following Excel rules.

#### Scenario: Error in left operand
- GIVEN BinaryOp with left=ErrorValue("#DIV/0!"), right=NumberValue(5)
- WHEN Evaluate(ctx) is called
- THEN ErrorValue("#DIV/0!") is returned immediately
- AND right operand is not evaluated (short-circuit)

#### Scenario: Error in right operand
- GIVEN BinaryOp with left=NumberValue(5), right=ErrorValue("#VALUE!")
- WHEN Evaluate(ctx) is called
- THEN ErrorValue("#VALUE!") is returned

#### Scenario: Error in function argument
- GIVEN FunctionCall("SUM", [ErrorValue("#REF!")])
- WHEN Evaluate(ctx) is called
- THEN ErrorValue("#REF!") is returned

#### Scenario: Error priority
- GIVEN multiple errors in expression
- WHEN errors propagate
- THEN higher priority error wins
- AND priority order: #NULL! > #DIV/0! > #VALUE! > #REF! > #NAME? > #NUM! > #N/A

### Requirement: Dependency Graph

The system SHALL track formula dependencies for calculating evaluation order.

#### Scenario: Add formula to dependency graph
- GIVEN a DependencyGraph
- AND formula "=A1+B1" in cell C1
- WHEN AddFormula(C1, expr) is called
- THEN C1 node is added to graph
- AND C1 has dependencies on A1 and B1
- AND A1 and B1 have C1 as dependent

#### Scenario: Remove formula from dependency graph
- GIVEN a DependencyGraph with formula in C1 depending on A1, B1
- WHEN RemoveFormula(C1) is called
- THEN C1 node is removed from graph
- AND A1 and B1 no longer have C1 as dependent

#### Scenario: Topological sort with no cycles
- GIVEN formulas: A1=5, B1=A1+1, C1=B1+1
- WHEN TopologicalSort() is called
- THEN cells are returned in order: [A1, B1, C1]
- AND no error is returned

#### Scenario: Topological sort with complex dependencies
- GIVEN formulas: A1=5, B1=10, C1=A1+B1, D1=C1*2, E1=A1+D1
- WHEN TopologicalSort() is called
- THEN A1 and B1 come before C1
- AND C1 comes before D1
- AND A1 and D1 come before E1

#### Scenario: Circular reference detection
- GIVEN formulas: A1=B1+1, B1=A1+1
- WHEN TopologicalSort() is called
- THEN an error is returned
- AND error indicates circular reference

#### Scenario: Find cycle path
- GIVEN circular reference A1→B1→C1→A1
- WHEN findCycle() is called
- THEN cycle path [A1, B1, C1, A1] is returned

#### Scenario: Get dependents
- GIVEN formulas: B1=A1+1, C1=A1+5, D1=B1+C1
- WHEN GetDependents(A1) is called
- THEN [B1, C1] are returned (direct dependents only)

### Requirement: Recalculation Engine

The system SHALL recalculate formulas in dependency order when cell values change.

#### Scenario: Recalculate workbook
- GIVEN a workbook with formulas: A1=5, B1=A1+1, C1=B1+1
- WHEN RecalculateWorkbook(wb) is called
- THEN A1 is evaluated first (result: 5)
- AND B1 is evaluated second (result: 6)
- AND C1 is evaluated third (result: 7)
- AND all cells have correct values

#### Scenario: Recalculate worksheet
- GIVEN a worksheet with formulas
- WHEN RecalculateWorksheet(ws) is called
- THEN all formulas in worksheet are recalculated in dependency order
- AND formulas in other worksheets are not affected

#### Scenario: Recalculate cell incrementally
- GIVEN formulas: A1=5, B1=A1+1, C1=B1+1, D1=10 (unrelated)
- WHEN A1 value changes to 10
- AND RecalculateCell(wb, "Sheet1", A1) is called
- THEN A1, B1, C1 are recalculated
- AND D1 is not recalculated (not dependent)

#### Scenario: Recalculation with circular reference error
- GIVEN formulas with circular reference: A1=B1, B1=A1
- WHEN RecalculateWorkbook(wb) is called
- THEN an error is returned
- AND error message indicates circular reference detected

#### Scenario: Recalculation preserves formula text
- GIVEN cell C1 with formula "=A1+B1"
- WHEN RecalculateWorkbook(wb) is called
- THEN C1 value is updated to computed result
- AND C1 formula text remains "=A1+B1" (not changed)

### Requirement: Evaluation Context

The system SHALL provide an EvalContext for formula evaluation with cell resolution and function registry.

#### Scenario: EvalContext creation
- GIVEN a Workbook, CurrentSheet, Resolver, FunctionRegistry
- WHEN EvalContext is created
- THEN context has access to all components
- AND expressions can evaluate using context

#### Scenario: CellResolver resolves cell value
- GIVEN a WorkbookResolver
- AND cell A1 contains 42
- WHEN GetCellValue("Sheet1", 1, 1) is called
- THEN NumberValue(42) is returned

#### Scenario: CellResolver resolves empty cell
- GIVEN a WorkbookResolver
- AND cell B5 is empty
- WHEN GetCellValue("Sheet1", 2, 5) is called
- THEN nil is returned

#### Scenario: CellResolver resolves formula cached value
- GIVEN cell C1 with formula "=A1+B1" and cached value 100
- WHEN GetCellValue("Sheet1", 3, 1) is called
- THEN NumberValue(100) is returned (from cache)

#### Scenario: Cross-sheet reference resolution
- GIVEN cell A1 on Sheet2 contains 50
- AND formula "=Sheet2!A1" in Sheet1
- WHEN formula is evaluated
- THEN value 50 is resolved from Sheet2.A1

### Requirement: Calculation Chain Builder

The system SHALL build and persist calculation chains representing formula evaluation order.

#### Scenario: Rebuild calculation chain
- GIVEN a workbook with formulas: A1=5, B1=A1+1, C1=B1+1
- WHEN RebuildCalculationChain(wb) is called
- THEN CalcChainPart is updated with cells [A1, B1, C1] in order
- AND calculation chain is persisted in workbook

#### Scenario: Calculation chain with multiple sheets
- GIVEN Sheet1 with A1=5, B1=A1+1
- AND Sheet2 with C1=Sheet1!B1+1
- WHEN RebuildCalculationChain(wb) is called
- THEN chain contains [Sheet1.A1, Sheet1.B1, Sheet2.C1] in order

#### Scenario: Calculation chain skips non-formula cells
- GIVEN cells: A1=5 (value), B1=A1+1 (formula), C1=10 (value)
- WHEN RebuildCalculationChain(wb) is called
- THEN chain contains only [B1] (formula cells only)

#### Scenario: Calculation chain with circular reference error
- GIVEN formulas with circular reference
- WHEN RebuildCalculationChain(wb) is called
- THEN an error is returned
- AND calculation chain is not updated

### Requirement: Parser Operator Precedence

The system SHALL parse formulas with correct operator precedence and associativity.

#### Scenario: Precedence comparison vs concatenation
- GIVEN formula "A1 = B1 & C1"
- WHEN Parse(formula) is called
- THEN AST is (A1 = (B1 & C1))
- AND concatenation evaluated before comparison

#### Scenario: Precedence concatenation vs addition
- GIVEN formula "A1 & B1 + C1"
- WHEN Parse(formula) is called
- THEN AST is (A1 & (B1 + C1))
- AND addition evaluated before concatenation

#### Scenario: Precedence addition vs multiplication
- GIVEN formula "A1 + B1 * C1"
- WHEN Parse(formula) is called
- THEN AST is (A1 + (B1 * C1))
- AND multiplication evaluated before addition

#### Scenario: Precedence multiplication vs exponentiation
- GIVEN formula "A1 * B1 ^ C1"
- WHEN Parse(formula) is called
- THEN AST is (A1 * (B1 ^ C1))
- AND exponentiation evaluated before multiplication

#### Scenario: Exponentiation right-associativity
- GIVEN formula "2 ^ 3 ^ 4"
- WHEN Parse(formula) is called
- THEN AST is (2 ^ (3 ^ 4))
- AND result is 2^81 = 2417851639229258349412352

#### Scenario: Parentheses override precedence
- GIVEN formula "(A1 + B1) * C1"
- WHEN Parse(formula) is called
- THEN AST is ((A1 + B1) * C1)
- AND addition evaluated first due to parentheses

### Requirement: Formula Validation

The system SHALL validate formula syntax and report errors.

#### Scenario: Valid formula validation
- GIVEN formula "=A1+B1"
- WHEN Parse(formula) is called
- THEN no error is returned
- AND AST is valid

#### Scenario: Invalid function syntax
- GIVEN formula "=SUM("
- WHEN Parse(formula) is called
- THEN an error is returned
- AND error message indicates missing closing parenthesis

#### Scenario: Invalid operator usage
- GIVEN formula "=A1 + + B1"
- WHEN Parse(formula) is called
- THEN an error is returned
- AND error message indicates unexpected operator

#### Scenario: Mismatched parentheses
- GIVEN formula "=(A1+B1"
- WHEN Parse(formula) is called
- THEN an error is returned
- AND error message indicates missing closing parenthesis

#### Scenario: Unknown function name
- GIVEN formula "=UNKNOWN(A1)"
- WHEN Parse(formula) is called
- THEN AST is created successfully (parser doesn't validate function names)
- WHEN Evaluate(ctx) is called
- THEN ErrorValue("#NAME?") is returned

### Requirement: Dependency Extraction

The system SHALL extract all cell and range dependencies from formula expressions.

#### Scenario: Extract cell reference dependency
- GIVEN expression CellReference("A1")
- WHEN Dependencies() is called
- THEN [Dependency{Type: Cell, CellRef: A1}] is returned

#### Scenario: Extract range reference dependency
- GIVEN expression RangeReference("A1:C10")
- WHEN Dependencies() is called
- THEN [Dependency{Type: Range, RangeRef: A1:C10}] is returned

#### Scenario: Extract binary operation dependencies
- GIVEN expression BinaryOp(CellRef("A1"), Plus, CellRef("B1"))
- WHEN Dependencies() is called
- THEN [Dependency(A1), Dependency(B1)] is returned

#### Scenario: Extract function call dependencies
- GIVEN expression FunctionCall("SUM", [RangeRef("A1:A10"), CellRef("B1")])
- WHEN Dependencies() is called
- THEN [Dependency(A1:A10), Dependency(B1)] is returned

#### Scenario: Extract nested expression dependencies
- GIVEN expression "=SUM(A1:A10) + AVERAGE(B1:B10)"
- WHEN Dependencies() is called
- THEN [Dependency(A1:A10), Dependency(B1:B10)] is returned

### Requirement: Array Formula Support

The system SHALL support array formulas that produce multiple results.

#### Scenario: Array formula evaluation
- GIVEN array formula "{=A1:A3+B1:B3}" in range C1:C3
- AND A1:A3 contains [1, 2, 3]
- AND B1:B3 contains [10, 20, 30]
- WHEN formula is evaluated
- THEN C1:C3 contains [11, 22, 33]

#### Scenario: Array formula single formula flag
- GIVEN array formula in range C1:C3
- WHEN formula is stored
- THEN C1 contains formula text
- AND C2:C3 reference C1 as master cell

#### Scenario: Array size mismatch handling
- GIVEN array formula with mismatched array sizes
- WHEN formula is evaluated
- THEN #N/A errors fill unmatched cells

### Requirement: Shared Formula Expansion

The system SHALL expand shared formulas to actual cell formulas with adjusted references.

#### Scenario: Shared formula expansion
- GIVEN shared formula "=A1+B1" in C1
- AND shared formula referenced by C2:C10
- WHEN formula in C5 is expanded
- THEN expanded formula is "=A5+B5" (references adjusted)

#### Scenario: Shared formula with absolute references
- GIVEN shared formula "=$A$1+B1" in C1
- WHEN formula in C5 is expanded
- THEN expanded formula is "=$A$1+B5" (absolute ref unchanged)

### Requirement: Performance Optimization

The system SHALL optimize formula evaluation for performance.

#### Scenario: Value caching
- GIVEN evaluated cell values cached
- WHEN same cell is resolved multiple times
- THEN cached value is returned without re-evaluation

#### Scenario: Incremental recalculation
- GIVEN change to cell A1
- WHEN incremental recalculation is triggered
- THEN only A1 and its dependents are recalculated
- AND unrelated cells are not recalculated

#### Scenario: Parallel evaluation of independent cells
- GIVEN cells with no interdependencies
- WHEN evaluation occurs
- THEN cells in same dependency level can be evaluated in parallel

#### Scenario: Lazy evaluation
- GIVEN formula cell with cached value
- WHEN cache is valid
- THEN formula is not re-evaluated
- AND cached value is used
