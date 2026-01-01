//nolint:all // formula package is scaffolding for future formula evaluation - WIP
// Package formula provides formula parsing and evaluation for spreadsheets.
package formula

import (
	"math"
	"strings"
)

// Function represents a spreadsheet function.
type Function interface {
	// Name returns the function name.
	Name() string

	// MinArgs returns the minimum number of arguments.
	MinArgs() int

	// MaxArgs returns the maximum number of arguments (-1 for unlimited).
	MaxArgs() int

	// Call evaluates the function with the given arguments.
	Call(ctx *EvalContext, args []Value) (Value, error)
}

// FunctionRegistry manages available functions.
type FunctionRegistry struct {
	functions map[string]Function
}

// NewFunctionRegistry creates a new function registry with built-in functions.
func NewFunctionRegistry() *FunctionRegistry {
	reg := &FunctionRegistry{
		functions: make(map[string]Function),
	}

	// Register built-in functions
	reg.Register(&SumFunction{})
	reg.Register(&AverageFunction{})
	reg.Register(&CountFunction{})
	reg.Register(&MinFunction{})
	reg.Register(&MaxFunction{})
	reg.Register(&IfFunction{})
	reg.Register(&AndFunction{})
	reg.Register(&OrFunction{})
	reg.Register(&NotFunction{})
	reg.Register(&ConcatenateFunction{})
	reg.Register(&LenFunction{})

	return reg
}

// Register adds a function to the registry.
func (r *FunctionRegistry) Register(fn Function) {
	r.functions[strings.ToUpper(fn.Name())] = fn
}

// Get retrieves a function by name (case-insensitive).
func (r *FunctionRegistry) Get(name string) Function {
	return r.functions[strings.ToUpper(name)]
}

// SumFunction implements the SUM function.
type SumFunction struct{}

func (f *SumFunction) Name() string { return "SUM" }
func (f *SumFunction) MinArgs() int { return 1 }
func (f *SumFunction) MaxArgs() int { return -1 }
func (f *SumFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	sum := 0.0
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		switch arg.Type {
		case ValueTypeArray:
			arr, ok := arg.Value.([][]Value)
			if !ok {
				return NewErrorValue(ErrValue), nil
			}
			for _, row := range arr {
				for _, cell := range row {
					if cell.Type == ValueError {
						return cell, nil
					}
					if cell.Type == ValueTypeNumber {
						n, _ := cell.AsNumber()
						sum += n
					}
				}
			}
		case ValueTypeNumber:
			n, _ := arg.AsNumber()
			sum += n
		}
	}

	return NewNumberValue(sum), nil
}

// AverageFunction implements the AVERAGE function.
type AverageFunction struct{}

func (f *AverageFunction) Name() string { return "AVERAGE" }
func (f *AverageFunction) MinArgs() int { return 1 }
func (f *AverageFunction) MaxArgs() int { return -1 }
func (f *AverageFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	sum := 0.0
	count := 0
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		switch arg.Type {
		case ValueTypeArray:
			arr, ok := arg.Value.([][]Value)
			if !ok {
				return NewErrorValue(ErrValue), nil
			}
			for _, row := range arr {
				for _, cell := range row {
					if cell.Type == ValueError {
						return cell, nil
					}
					if cell.Type == ValueTypeNumber {
						n, _ := cell.AsNumber()
						sum += n
						count++
					}
				}
			}
		case ValueTypeNumber:
			n, _ := arg.AsNumber()
			sum += n
			count++
		}
	}
	if count == 0 {
		return NewErrorValue(ErrDiv0), nil
	}

	return NewNumberValue(sum / float64(count)), nil
}

// CountFunction implements the COUNT function.
type CountFunction struct{}

func (f *CountFunction) Name() string { return "COUNT" }
func (f *CountFunction) MinArgs() int { return 1 }
func (f *CountFunction) MaxArgs() int { return -1 }
func (f *CountFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	count := 0
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		switch arg.Type {
		case ValueTypeArray:
			arr, ok := arg.Value.([][]Value)
			if !ok {
				return NewErrorValue(ErrValue), nil
			}
			for _, row := range arr {
				for _, cell := range row {
					if cell.Type == ValueError {
						return cell, nil
					}
					if cell.Type == ValueTypeNumber {
						count++
					}
				}
			}
		case ValueTypeNumber:
			count++
		}
	}

	return NewNumberValue(float64(count)), nil
}

// MinFunction implements the MIN function.
type MinFunction struct{}

func (f *MinFunction) Name() string { return "MIN" }
func (f *MinFunction) MinArgs() int { return 1 }
func (f *MinFunction) MaxArgs() int { return -1 }
func (f *MinFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	min := math.Inf(1)
	found := false
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		switch arg.Type {
		case ValueTypeArray:
			arr, ok := arg.Value.([][]Value)
			if !ok {
				return NewErrorValue(ErrValue), nil
			}
			for _, row := range arr {
				for _, cell := range row {
					if cell.Type == ValueError {
						return cell, nil
					}
					if cell.Type == ValueTypeNumber {
						n, _ := cell.AsNumber()
						if n < min {
							min = n
						}
						found = true
					}
				}
			}
		case ValueTypeNumber:
			n, _ := arg.AsNumber()
			if n < min {
				min = n
			}
			found = true
		}
	}
	if !found {
		return NewNumberValue(0), nil
	}

	return NewNumberValue(min), nil
}

// MaxFunction implements the MAX function.
type MaxFunction struct{}

func (f *MaxFunction) Name() string { return "MAX" }
func (f *MaxFunction) MinArgs() int { return 1 }
func (f *MaxFunction) MaxArgs() int { return -1 }
func (f *MaxFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	max := math.Inf(-1)
	found := false
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		switch arg.Type {
		case ValueTypeArray:
			arr, ok := arg.Value.([][]Value)
			if !ok {
				return NewErrorValue(ErrValue), nil
			}
			for _, row := range arr {
				for _, cell := range row {
					if cell.Type == ValueError {
						return cell, nil
					}
					if cell.Type == ValueTypeNumber {
						n, _ := cell.AsNumber()
						if n > max {
							max = n
						}
						found = true
					}
				}
			}
		case ValueTypeNumber:
			n, _ := arg.AsNumber()
			if n > max {
				max = n
			}
			found = true
		}
	}
	if !found {
		return NewNumberValue(0), nil
	}

	return NewNumberValue(max), nil
}

// IfFunction implements the IF function.
type IfFunction struct{}

func (f *IfFunction) Name() string { return "IF" }
func (f *IfFunction) MinArgs() int { return 2 }
func (f *IfFunction) MaxArgs() int { return 3 }
func (f *IfFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	condition := args[0]
	if condition.Type == ValueError {
		return condition, nil
	}

	cond, err := condition.AsBoolean()
	if err != nil {
		return NewErrorValue(ErrValue), nil
	}

	if cond {
		return args[1], nil
	}

	if len(args) == 3 {
		return args[2], nil
	}

	return NewBooleanValue(false), nil
}

// AndFunction implements the AND function.
type AndFunction struct{}

func (f *AndFunction) Name() string { return "AND" }
func (f *AndFunction) MinArgs() int { return 1 }
func (f *AndFunction) MaxArgs() int { return -1 }
func (f *AndFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		cond, err := arg.AsBoolean()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if !cond {
			return NewBooleanValue(false), nil
		}
	}

	return NewBooleanValue(true), nil
}

// OrFunction implements the OR function.
type OrFunction struct{}

func (f *OrFunction) Name() string { return "OR" }
func (f *OrFunction) MinArgs() int { return 1 }
func (f *OrFunction) MaxArgs() int { return -1 }
func (f *OrFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		cond, err := arg.AsBoolean()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if cond {
			return NewBooleanValue(true), nil
		}
	}

	return NewBooleanValue(false), nil
}

// NotFunction implements the NOT function.
type NotFunction struct{}

func (f *NotFunction) Name() string { return "NOT" }
func (f *NotFunction) MinArgs() int { return 1 }
func (f *NotFunction) MaxArgs() int { return 1 }
func (f *NotFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if args[0].Type == ValueError {
		return args[0], nil
	}
	cond, err := args[0].AsBoolean()
	if err != nil {
		return NewErrorValue(ErrValue), nil
	}

	return NewBooleanValue(!cond), nil
}

// ConcatenateFunction implements the CONCATENATE function.
type ConcatenateFunction struct{}

func (f *ConcatenateFunction) Name() string { return "CONCATENATE" }
func (f *ConcatenateFunction) MinArgs() int { return 1 }
func (f *ConcatenateFunction) MaxArgs() int { return -1 }
func (f *ConcatenateFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	var result strings.Builder
	for _, arg := range args {
		if arg.Type == ValueError {
			return arg, nil
		}
		result.WriteString(arg.AsString())
	}

	return NewStringValue(result.String()), nil
}

// LenFunction implements the LEN function.
type LenFunction struct{}

func (f *LenFunction) Name() string { return "LEN" }
func (f *LenFunction) MinArgs() int { return 1 }
func (f *LenFunction) MaxArgs() int { return 1 }
func (f *LenFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if args[0].Type == ValueError {
		return args[0], nil
	}

	return NewNumberValue(float64(len(args[0].AsString()))), nil
}
