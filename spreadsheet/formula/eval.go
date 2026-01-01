//nolint:all // formula package is scaffolding for future formula evaluation - WIP
// Package formula provides formula parsing and evaluation for spreadsheets.
package formula

import (
	"fmt"
	"math"
	"strconv"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

// ValueType represents the type of a value.
type ValueType int

const (
	// ValueTypeNumber represents a numeric value.
	ValueTypeNumber ValueType = iota
	// ValueTypeString represents a string value.
	ValueTypeString
	// ValueTypeBoolean represents a boolean value.
	ValueTypeBoolean
	// ValueError represents an error value.
	ValueError
	// ValueTypeArray represents an array of values.
	ValueTypeArray
)

// Value represents a value in formula evaluation.
type Value struct {
	Type  ValueType
	Value interface{}
}

// Excel error constants.
const (
	ErrDiv0  = "#DIV/0!"
	ErrValue = "#VALUE!"
	ErrRef   = "#REF!"
	ErrName  = "#NAME?"
	ErrNA    = "#N/A"
	ErrNull  = "#NULL!"
	ErrNum   = "#NUM!"
)

// NewNumberValue creates a new number value.
func NewNumberValue(v float64) Value {
	return Value{Type: ValueTypeNumber, Value: v}
}

// NewStringValue creates a new string value.
func NewStringValue(v string) Value {
	return Value{Type: ValueTypeString, Value: v}
}

// NewBooleanValue creates a new boolean value.
func NewBooleanValue(v bool) Value {
	return Value{Type: ValueTypeBoolean, Value: v}
}

// NewErrorValue creates a new error value.
func NewErrorValue(err string) Value {
	return Value{Type: ValueError, Value: err}
}

// NewArrayValue creates a new array value.
func NewArrayValue(vals [][]Value) Value {
	return Value{Type: ValueTypeArray, Value: vals}
}

// AsNumber converts a value to a number.
// Returns an error value if conversion fails.
func (v Value) AsNumber() (float64, error) {
	switch v.Type {
	case ValueTypeNumber:
		if n, ok := v.Value.(float64); ok {
			return n, nil
		}

		return 0, fmt.Errorf("invalid number value")
	case ValueTypeString:
		s, ok := v.Value.(string)
		if !ok {
			return 0, fmt.Errorf("invalid string value")
		}
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}

		return n, nil
	case ValueTypeBoolean:
		if b, ok := v.Value.(bool); ok && b {
			return 1, nil
		}

		return 0, nil
	case ValueError:
		return 0, fmt.Errorf("error value: %v", v.Value)
	case ValueTypeArray:
		return 0, fmt.Errorf("cannot convert array to number")
	default:
		return 0, fmt.Errorf("cannot convert to number")
	}
}

// AsString converts a value to a string.
func (v Value) AsString() string {
	switch v.Type {
	case ValueTypeString:
		if s, ok := v.Value.(string); ok {
			return s
		}

		return ""
	case ValueTypeNumber:
		if n, ok := v.Value.(float64); ok {
			return strconv.FormatFloat(n, 'f', -1, 64)
		}

		return ""
	case ValueTypeBoolean:
		if b, ok := v.Value.(bool); ok && b {
			return BooleanTrue
		}

		return BooleanFalse
	case ValueError:
		return fmt.Sprintf("%v", v.Value)
	case ValueTypeArray:
		return fmt.Sprintf("%v", v.Value)
	default:
		return fmt.Sprintf("%v", v.Value)
	}
}

// AsBoolean converts a value to a boolean.
func (v Value) AsBoolean() (bool, error) {
	switch v.Type {
	case ValueTypeBoolean:
		if b, ok := v.Value.(bool); ok {
			return b, nil
		}

		return false, fmt.Errorf("invalid boolean value")
	case ValueTypeNumber:
		n, err := v.AsNumber()
		if err != nil {
			return false, err
		}

		return n != 0, nil
	case ValueTypeString:
		s := v.AsString()
		switch s {
		case "TRUE", "True", "true":
			return true, nil
		case "FALSE", "False", "false":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert string to boolean")
		}
	case ValueError:
		return false, fmt.Errorf("error value: %v", v.Value)
	default:
		return false, fmt.Errorf("cannot convert to boolean")
	}
}

// EvalContext provides context for formula evaluation.
type EvalContext struct {
	// Workbook is the current workbook (can be nil for standalone evaluation).
	Workbook interface{}

	// CurrentSheet is the name of the current sheet.
	CurrentSheet string

	// Resolver resolves cell and range references.
	Resolver CellResolver

	// FunctionRegistry provides access to functions.
	FunctionRegistry *FunctionRegistry
}

// NewEvalContext creates a new evaluation context.
func NewEvalContext(workbook interface{}, currentSheet string) *EvalContext {
	return &EvalContext{
		Workbook:         workbook,
		CurrentSheet:     currentSheet,
		FunctionRegistry: NewFunctionRegistry(),
	}
}

// CellResolver resolves cell and range references to values.
type CellResolver interface {
	// ResolveCell returns the value of a cell.
	ResolveCell(ref spreadsheet.CellRef) (Value, error)

	// ResolveRange returns all values in a range as a 2D array.
	ResolveRange(rng spreadsheet.RangeRef) ([][]Value, error)
}

// evaluateBinaryOp evaluates a binary operation.
//
//nolint:revive // cyclomatic: binary operations have many cases
func evaluateBinaryOp(left, right Value, op string) (Value, error) {
	switch op {
	case "+":
		l, err := left.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		r, err := right.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}

		return NewNumberValue(l + r), nil

	case "-":
		l, err := left.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		r, err := right.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}

		return NewNumberValue(l - r), nil

	case "*":
		l, err := left.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		r, err := right.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}

		return NewNumberValue(l * r), nil

	case "/":
		l, err := left.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		r, err := right.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if r == 0 {
			return NewErrorValue(ErrDiv0), nil
		}

		return NewNumberValue(l / r), nil

	case "^":
		l, err := left.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		r, err := right.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		result := math.Pow(l, r)
		if math.IsNaN(result) || math.IsInf(result, 0) {
			return NewErrorValue(ErrNum), nil
		}

		return NewNumberValue(result), nil

	case "&":
		// String concatenation
		return NewStringValue(left.AsString() + right.AsString()), nil

	case "=":
		return NewBooleanValue(compareValues(left, right) == 0), nil

	case "<>":
		return NewBooleanValue(compareValues(left, right) != 0), nil

	case "<":
		return NewBooleanValue(compareValues(left, right) < 0), nil

	case "<=":
		return NewBooleanValue(compareValues(left, right) <= 0), nil

	case ">":
		return NewBooleanValue(compareValues(left, right) > 0), nil

	case ">=":
		return NewBooleanValue(compareValues(left, right) >= 0), nil

	default:
		return NewErrorValue(ErrValue), nil
	}
}

// evaluateUnaryOp evaluates a unary operation.
func evaluateUnaryOp(operand Value, op string) (Value, error) {
	switch op {
	case "-":
		n, err := operand.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}

		return NewNumberValue(-n), nil

	case "+":
		n, err := operand.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}

		return NewNumberValue(n), nil

	default:
		return NewErrorValue(ErrValue), nil
	}
}

// compareValues compares two values.
// Returns -1 if left < right, 0 if left == right, 1 if left > right.
func compareValues(left, right Value) int {
	// Type priority: number > string > boolean
	if left.Type != right.Type {
		return int(left.Type) - int(right.Type)
	}

	switch left.Type {
	case ValueTypeNumber:
		l, _ := left.AsNumber()
		r, _ := right.AsNumber()
		if l < r {
			return -1
		} else if l > r {
			return 1
		}

		return 0

	case ValueTypeString:
		l := left.AsString()
		r := right.AsString()
		if l < r {
			return -1
		} else if l > r {
			return 1
		}

		return 0

	case ValueTypeBoolean:
		l, _ := left.AsBoolean()
		r, _ := right.AsBoolean()
		if !l && r {
			return -1
		} else if l && !r {
			return 1
		}

		return 0

	default:
		return 0
	}
}
