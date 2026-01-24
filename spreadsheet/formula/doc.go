//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package formula provides formula parsing and evaluation for Excel-compatible formulas.
//
// # Overview
//
// This package implements a complete formula parser and evaluator for Excel formulas,
// supporting basic arithmetic, comparison operations, cell/range references, and
// common functions.
//
// # Basic Usage
//
//	// Create a parser
//	parser := formula.NewParser()
//
//	// Parse a formula
//	expr, err := parser.Parse("=SUM(A1:A10)+B1*2")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Create an evaluation context
//	ctx := formula.NewEvalContext(workbook, "Sheet1")
//	ctx.Resolver = myResolver  // Implement CellResolver interface
//
//	// Evaluate the formula
//	value, err := expr.Evaluate(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get the result
//	result, _ := value.AsNumber()
//	fmt.Printf("Result: %v\n", result)
//
// # Supported Operations
//
// Arithmetic: +, -, *, /, ^ (exponentiation)
//
// Comparison: =, <>, <, <=, >, >=
//
// String: & (concatenation)
//
// # Supported Functions
//
// Math: SUM, AVERAGE, COUNT, MIN, MAX
//
// Logical: IF, AND, OR, NOT
//
// Text: CONCATENATE, LEN
//
// # Operator Precedence
//
// From highest to lowest:
//
// 1. Parentheses: ()
//
// 2. Unary: -, +
//
// 3. Exponentiation: ^ (right-associative)
//
// 4. Multiplication/Division: *, /
//
// 5. Addition/Subtraction: +, -
//
// 6. Concatenation: &
//
// 7. Comparison: =, <>, <, <=, >, >=
//
// # Cell Resolver
//
// To evaluate formulas that reference cells, implement the CellResolver interface:
//
//	type MyCellResolver struct {
//	    sheet *spreadsheet.Sheet
//	}
//
//	func (r *MyCellResolver) ResolveCell(ref spreadsheet.CellRef) (formula.Value, error) {
//	    cell := r.sheet.Cell(ref.Row, ref.Col)
//	    if cell == nil {
//	        return formula.NewNumberValue(0), nil
//	    }
//	    return formula.NewNumberValue(cell.GetNumber()), nil
//	}
//
//	func (r *MyCellResolver) ResolveRange(rng spreadsheet.RangeRef) ([][]formula.Value, error) {
//	    // Return 2D array of values for the range
//	    // ...
//	}
//
// # Error Handling
//
// The evaluator propagates Excel-compatible errors:
//
//	#DIV/0! - Division by zero
//	#VALUE! - Invalid type conversion
//	#REF!   - Invalid cell reference
//	#NAME?  - Unknown function name
//	#N/A    - Value not available
//	#NULL!  - Null intersection
//	#NUM!   - Invalid numeric value
package formula
