//nolint:all // formula package is scaffolding for future formula evaluation - WIP
package formula

import (
	"fmt"
	"strings"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

const (
	// BooleanTrue is the string representation of TRUE.
	BooleanTrue = "TRUE"
	// BooleanFalse is the string representation of FALSE.
	BooleanFalse = "FALSE"
)

// Expression represents a node in the abstract syntax tree of a formula.
type Expression interface {
	// Evaluate computes the value of this expression in the given context.
	Evaluate(ctx *EvalContext) (Value, error)

	// Dependencies returns all cell references this expression depends on.
	Dependencies() []spreadsheet.CellRef

	// String returns a string representation of this expression.
	String() string
}

// Literal represents a literal value (number, string, boolean).
type Literal struct {
	Val Value
}

// Evaluate returns the literal value.
func (l *Literal) Evaluate(_ *EvalContext) (Value, error) {
	return l.Val, nil
}

// Dependencies returns an empty slice (literals have no dependencies).
func (_ *Literal) Dependencies() []spreadsheet.CellRef {
	return nil
}

// String returns the string representation of the literal.
func (l *Literal) String() string {
	switch l.Val.Type {
	case ValueTypeNumber:
		return fmt.Sprintf("%v", l.Val.Value)
	case ValueTypeString:
		return fmt.Sprintf("%q", l.Val.Value)
	case ValueTypeBoolean:
		if b, ok := l.Val.Value.(bool); ok && b {
			return BooleanTrue
		}

		return BooleanFalse
	case ValueError:
		return fmt.Sprintf("%v", l.Val.Value)
	case ValueTypeArray:
		return fmt.Sprintf("%v", l.Val.Value)
	default:
		return fmt.Sprintf("%v", l.Val.Value)
	}
}

// CellReference represents a reference to a single cell.
type CellReference struct {
	Ref spreadsheet.CellRef
}

// Evaluate retrieves the cell value from the context.
func (c *CellReference) Evaluate(ctx *EvalContext) (Value, error) {
	if ctx.Resolver == nil {
		return NewErrorValue(ErrNA), nil
	}

	val, err := ctx.Resolver.ResolveCell(c.Ref)
	if err != nil {
		return NewErrorValue(ErrRef), nil
	}

	return val, nil
}

// Dependencies returns the cell reference.
func (c *CellReference) Dependencies() []spreadsheet.CellRef {
	return []spreadsheet.CellRef{c.Ref}
}

// String returns the A1-style reference.
func (c *CellReference) String() string {
	return c.Ref.String()
}

// RangeReference represents a reference to a range of cells.
type RangeReference struct {
	Range spreadsheet.RangeRef
}

// Evaluate retrieves all cell values in the range as an array.
func (r *RangeReference) Evaluate(ctx *EvalContext) (Value, error) {
	if ctx.Resolver == nil {
		return NewErrorValue(ErrNA), nil
	}

	vals, err := ctx.Resolver.ResolveRange(r.Range)
	if err != nil {
		return NewErrorValue(ErrRef), nil
	}

	return Value{
		Type:  ValueTypeArray,
		Value: vals,
	}, nil
}

// Dependencies returns all cell references in the range.
func (r *RangeReference) Dependencies() []spreadsheet.CellRef {
	size := r.Range.Size()
	deps := make([]spreadsheet.CellRef, 0, size)
	for cell := range r.Range.Cells() {
		deps = append(deps, cell)
	}

	return deps
}

// String returns the A1-style range reference.
func (r *RangeReference) String() string {
	return r.Range.String()
}

// BinaryOp represents a binary operation.
type BinaryOp struct {
	Left     Expression
	Right    Expression
	Operator string
}

// Evaluate computes the binary operation.
func (b *BinaryOp) Evaluate(ctx *EvalContext) (Value, error) {
	left, err := b.Left.Evaluate(ctx)
	if err != nil {
		return Value{}, err
	}

	// Short-circuit error propagation
	if left.Type == ValueError {
		return left, nil
	}

	right, err := b.Right.Evaluate(ctx)
	if err != nil {
		return Value{}, err
	}

	// Short-circuit error propagation
	if right.Type == ValueError {
		return right, nil
	}

	return evaluateBinaryOp(left, right, b.Operator)
}

// Dependencies returns all dependencies from both operands.
func (b *BinaryOp) Dependencies() []spreadsheet.CellRef {
	deps := b.Left.Dependencies()
	deps = append(deps, b.Right.Dependencies()...)

	return deps
}

// String returns the string representation of the binary operation.
func (b *BinaryOp) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

// UnaryOp represents a unary operation.
type UnaryOp struct {
	Operand  Expression
	Operator string
}

// Evaluate computes the unary operation.
func (u *UnaryOp) Evaluate(ctx *EvalContext) (Value, error) {
	operand, err := u.Operand.Evaluate(ctx)
	if err != nil {
		return Value{}, err
	}

	// Short-circuit error propagation
	if operand.Type == ValueError {
		return operand, nil
	}

	return evaluateUnaryOp(operand, u.Operator)
}

// Dependencies returns all dependencies from the operand.
func (u *UnaryOp) Dependencies() []spreadsheet.CellRef {
	return u.Operand.Dependencies()
}

// String returns the string representation of the unary operation.
func (u *UnaryOp) String() string {
	return fmt.Sprintf("(%s%s)", u.Operator, u.Operand.String())
}

// FunctionCall represents a function call.
type FunctionCall struct {
	Name string
	Args []Expression
}

// Evaluate computes the function call.
func (f *FunctionCall) Evaluate(ctx *EvalContext) (Value, error) {
	if ctx.FunctionRegistry == nil {
		return NewErrorValue(ErrName), nil
	}

	fn := ctx.FunctionRegistry.Get(f.Name)
	if fn == nil {
		return NewErrorValue(ErrName), nil
	}

	// Check argument count
	if len(f.Args) < fn.MinArgs() || (fn.MaxArgs() >= 0 && len(f.Args) > fn.MaxArgs()) {
		return NewErrorValue(ErrValue), nil
	}

	// Evaluate all arguments
	args := make([]Value, len(f.Args))
	for i, arg := range f.Args {
		val, err := arg.Evaluate(ctx)
		if err != nil {
			return Value{}, err
		}
		args[i] = val
	}

	// Call the function
	return fn.Call(ctx, args)
}

// Dependencies returns all dependencies from all arguments.
func (f *FunctionCall) Dependencies() []spreadsheet.CellRef {
	var deps []spreadsheet.CellRef
	for _, arg := range f.Args {
		deps = append(deps, arg.Dependencies()...)
	}

	return deps
}

// String returns the string representation of the function call.
func (f *FunctionCall) String() string {
	args := make([]string, 0, len(f.Args))
	for _, arg := range f.Args {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%s(%s)", strings.ToUpper(f.Name), strings.Join(args, ", "))
}
