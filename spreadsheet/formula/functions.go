package formula

import (
	"math"
)

// Excel function argument limits.
const (
	excelMaxArgsLogical = 255 // Excel allows up to 255 arguments for AND/OR
	excelMaxArgsMath    = 255 // Excel allows up to 255 arguments for SUM etc
	excelMaxArgsIf      = 3   // IF function takes 2-3 arguments
)

// AndFunction implements AND function.
type AndFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *AndFunction) Name() string {
	return "AND"
}

//nolint:revive // unused-receiver: standard function interface
func (f *AndFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *AndFunction) MaxArgs() int {
	return excelMaxArgsLogical
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

// OrFunction implements OR function.
type OrFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *OrFunction) Name() string {
	return "OR"
}

//nolint:revive // unused-receiver: standard function interface
func (f *OrFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *OrFunction) MaxArgs() int {
	return excelMaxArgsLogical
}

//nolint:revive // unused-receiver: standard function interface
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

// NotFunction implements NOT function.
type NotFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *NotFunction) Name() string {
	return "NOT"
}

//nolint:revive // unused-receiver: standard function interface
func (f *NotFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *NotFunction) MaxArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
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

// IfFunction implements IF function.
type IfFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *IfFunction) Name() string {
	return "IF"
}

//nolint:revive // unused-receiver: standard function interface
func (f *IfFunction) MinArgs() int {
	return 2
}

//nolint:revive // unused-receiver: standard function interface
func (f *IfFunction) MaxArgs() int {
	return excelMaxArgsIf
}

//nolint:revive // unused-receiver: standard function interface
func (f *IfFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 2 || len(args) > excelMaxArgsIf {
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
	
	if len(args) >= excelMaxArgsIf {
		return args[2], nil // FALSE branch with else
	}
	
	return NewBooleanValue(false), nil // FALSE branch
}

// SumFunction implements SUM function.
type SumFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *SumFunction) Name() string {
	return "SUM"
}

//nolint:revive // unused-receiver: standard function interface
func (f *SumFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *SumFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
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

// AverageFunction implements AVERAGE function.
type AverageFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *AverageFunction) Name() string {
	return "AVERAGE"
}

//nolint:revive // unused-receiver: standard function interface
func (f *AverageFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *AverageFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
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

// CountFunction implements COUNT function.
type CountFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *CountFunction) Name() string {
	return "COUNT"
}

//nolint:revive // unused-receiver: standard function interface
func (f *CountFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *CountFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
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

// MaxFunction implements MAX function.
type MaxFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *MaxFunction) Name() string {
	return "MAX"
}

//nolint:revive // unused-receiver: standard function interface
func (f *MaxFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *MaxFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
func (f *MaxFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}

	maxVal := math.Inf(-1)
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if numVal > maxVal {
			maxVal = numVal
		}
	}

	if maxVal == math.Inf(-1) {
		return NewErrorValue(ErrNA), nil
	}

	return NewNumberValue(maxVal), nil
}

// MinFunction implements MIN function.
type MinFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *MinFunction) Name() string {
	return "MIN"
}

//nolint:revive // unused-receiver: standard function interface
func (f *MinFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *MinFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
func (f *MinFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) < 1 {
		return NewErrorValue(ErrValue), nil
	}

	minVal := math.Inf(1)
	for _, arg := range args {
		numVal, err := arg.AsNumber()
		if err != nil {
			return NewErrorValue(ErrValue), nil
		}
		if numVal < minVal {
			minVal = numVal
		}
	}

	if minVal == math.Inf(1) {
		return NewErrorValue(ErrNA), nil
	}

	return NewNumberValue(minVal), nil
}

// ConcatenateFunction implements CONCATENATE function.
type ConcatenateFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *ConcatenateFunction) Name() string {
	return "CONCATENATE"
}

//nolint:revive // unused-receiver: standard function interface
func (f *ConcatenateFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *ConcatenateFunction) MaxArgs() int {
	return excelMaxArgsMath
}

//nolint:revive // unused-receiver: standard function interface
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

// LenFunction implements LEN function.
type LenFunction struct{}

//nolint:revive // unused-receiver: standard function interface
func (f *LenFunction) Name() string {
	return "LEN"
}

//nolint:revive // unused-receiver: standard function interface
func (f *LenFunction) MinArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *LenFunction) MaxArgs() int {
	return 1
}

//nolint:revive // unused-receiver: standard function interface
func (f *LenFunction) Call(ctx *EvalContext, args []Value) (Value, error) {
	if len(args) != 1 {
		return NewErrorValue(ErrValue), nil
	}
	
	arg := args[0]
	str := arg.AsString()
	
	return NewNumberValue(float64(len(str))), nil
}

