package formula

import (
	"math"
)

// AndFunction implements AND function
type AndFunction struct{}

func (f *AndFunction) Name() string {
	return "AND"
}

func (f *AndFunction) MinArgs() int {
	return 1
}

func (f *AndFunction) MaxArgs() int {
	return 255 // Excel allows up to 255 arguments for AND
}

func (f *AndFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	// Excel's AND returns FALSE if any argument is FALSE
	for _, arg := range args {
		val, err := arg.AsBoolean()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if !val {
			return NewBooleanValue(false), nil
		}
	}
	
	return NewBooleanValue(true), nil
}

// OrFunction implements OR function
type OrFunction struct{}

func (f *OrFunction) Name() string {
	return "OR"
}

func (f *OrFunction) MinArgs() int {
	return 1
}

func (f *OrFunction) MaxArgs() int {
	return 255 // Excel allows up to 255 arguments for OR
}

func (f *OrFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	// Excel's OR returns TRUE if any argument is TRUE
	for _, arg := range args {
		val, err := arg.AsBoolean()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if val {
			return NewBooleanValue(true), nil
		}
	}
	
	return NewBooleanValue(false), nil
}

// NotFunction implements NOT function
type NotFunction struct{}

func (f *NotFunction) Name() string {
	return "NOT"
}

func (f *NotFunction) MinArgs() int {
	return 1
}

func (f *NotFunction) MaxArgs() int {
	return 1
}

func (f *NotFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) != 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	arg := args[0]
	val, err := arg.AsBoolean()
	if err != nil {
		return NewErrorValue(ErrValue), nil
	}
	
	return NewBooleanValue(!val), nil
}

// IfFunction implements IF function
type IfFunction struct{}

func (f *IfFunction) Name() string {
	return "IF"
}

func (f *IfFunction) MinArgs() int {
	return 2
}

func (f *IfFunction) MaxArgs() int {
	return 3
}

func (f *IfFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 2 || len(args) > 3 {
		return NewErrorValue(ErrValue), nil
	}
	
	condition := args[0]
	condVal, err := condition.AsBoolean()
	if err != nil {
		return NewErrorValue(ErrValue), nil
	}
	
	if condVal {
		return args[1], nil // TRUE branch
	}
	
	if len(args) >= 3 {
		return args[2], nil // FALSE branch with else
	}
	
	return NewBooleanValue(false), nil // FALSE branch
}

// SumFunction implements SUM function
type SumFunction struct{}

func (f *SumFunction) Name() string {
	return "SUM"
}

func (f *SumFunction) MinArgs() int {
	return 1
}

func (f *SumFunction) MaxArgs() int {
	return 255 // Excel allows up to 255 arguments for SUM
}

func (f *SumFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	sum := 0.0
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		sum += numVal
	}
	
	return NewNumberValue(sum), nil
}

// AverageFunction implements AVERAGE function
type AverageFunction struct{}

func (f *AverageFunction) Name() string {
	return "AVERAGE"
}

func (f *AverageFunction) MinArgs() int {
	return 1
}

func (f *AverageFunction) MaxArgs() int {
	return 255
}

func (f *AverageFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	sum := 0.0
	count := 0
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		sum += numVal
		count++
	}
	
	if count == 0 {
		return NewErrorValue(ErrDiv0), nil
	}
	
	return NewNumberValue(sum / float64(count)), nil
}

// CountFunction implements COUNT function
type CountFunction struct{}

func (f *CountFunction) Name() string {
	return "COUNT"
}

func (f *CountFunction) MinArgs() int {
	return 1
}

func (f *CountFunction) MaxArgs() int {
	return 255
}

func (f *CountFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	count := 0
	for _, arg := range args {
		if arg.Type == ValueNumber {
			count++
		}
	}
	
	return NewNumberValue(float64(count)), nil
}

// MaxFunction implements MAX function
type MaxFunction struct{}

func (f *MaxFunction) Name() string {
	return "MAX"
}

func (f *MaxFunction) MinArgs() int {
	return 1
}

func (f *MaxFunction) MaxArgs() int {
	return 255
}

func (f *MaxFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	max := math.Inf(-1)
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if numVal > max {
			max = numVal
		}
	}
	
	if max == math.Inf(-1) {
		return NewErrorValue(ErrNA), nil
	}
	
	return NewNumberValue(max), nil
}

// MinFunction implements MIN function
type MinFunction struct{}

func (f *MinFunction) Name() string {
	return "MIN"
}

func (f *MinFunction) MinArgs() int {
	return 1
}

func (f *MinFunction) MaxArgs() int {
	return 255
}

func (f *MinFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	min := math.Inf(1)
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if numVal < min {
			min = numVal
		}
	}
	
	if min == math.Inf(1) {
		return NewErrorValue(ErrNA), nil
	}
	
	return NewNumberValue(min), nil
}

// ConcatenateFunction implements CONCATENATE function
type ConcatenateFunction struct{}

func (f *ConcatenateFunction) Name() string {
	return "CONCATENATE"
}

func (f *ConcatenateFunction) MinArgs() int {
	return 1
}

func (f *ConcatenateFunction) MaxArgs() int {
	return 255
}

func (f *ConcatenateFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	result := ""
	for _, arg := range args {
		result += arg.AsString()
	}
	
	return NewStringValue(result), nil
}

// LenFunction implements LEN function
type LenFunction struct{}

func (f *LenFunction) Name() string {
	return "LEN"
}

func (f *LenFunction) MinArgs() int {
	return 1
}

func (f *LenFunction) MaxArgs() int {
	return 1
}

func (f *LenFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) != 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	arg := args[0]
	str := arg.AsString()
	
	return NewNumberValue(float64(len(str))), nil
}

