// Package formula implements Excel formula evaluation engine.
package formula

import (
	"errors"
	"fmt"

	"github.com/connerohnesorge/goffice/spreadsheet"
)

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

// CellResolver resolves cell and range references to values.
type CellResolver interface {
	// ResolveCell returns the value of a cell.
	ResolveCell(ref spreadsheet.CellRef) (Value, error)

	// ResolveRange returns all values in a range as a 2D array.
	ResolveRange(rng spreadsheet.RangeRef) ([][]Value, error)
}

// NewEvalContext creates a new evaluation context.
func NewEvalContext(workbook interface{}, currentSheet string) *EvalContext {
	return &EvalContext{
		Workbook:         workbook,
		CurrentSheet:     currentSheet,
		Resolver:         nil, // Must be set by caller
		FunctionRegistry: NewFunctionRegistry(),
	}
}

// NewEvalContextWithResolver creates a new evaluation context with a custom resolver.
func NewEvalContextWithResolver(workbook interface{}, currentSheet string, resolver CellResolver) *EvalContext {
	return &EvalContext{
		Workbook:         workbook,
		CurrentSheet:     currentSheet,
		Resolver:         resolver,
		FunctionRegistry: NewFunctionRegistry(),
	}
}

// EvaluateFormula evaluates a formula string and returns the result.
func EvaluateFormula(formula string, ctx *EvalContext) (Value, error) {
	if formula == "" {
		return NewEmptyValue(), nil
	}

	// Parse the formula
	expr, err := Parse(formula)
	if err != nil {
		return NewErrorValue(ErrValue), fmt.Errorf("parse error: %w", err)
	}

	// Evaluate the expression
	return expr.Evaluate(ctx)
}

// EvaluateCellFormula evaluates a cell's formula and returns the result.
func EvaluateCellFormula(cellValue interface{}, cellResolver CellResolver) (Value, error) {
	// Get the formula from the cell
	var formula string
	switch v := cellValue.(type) {
	case string:
		formula = v
	case nil:
		return NewEmptyValue(), nil
	default:
		return NewErrorValue(ErrValue), errors.New("cell value is not a string")
	}

	if formula == "" || formula[0] != '=' {
		// Not a formula, return the literal value
		return convertToValue(cellValue), nil
	}

	// Evaluate the formula (strip the leading '=')
	ctx := NewEvalContext(nil, "")
	ctx.Resolver = cellResolver

	result, err := EvaluateFormula(formula[1:], ctx)
	if err != nil {
		return NewErrorValue(ErrValue), err
	}

	return result, nil
}

// convertToValue converts a Go value to a formula Value.
func convertToValue(v interface{}) Value {
	switch val := v.(type) {
	case nil:
		return NewEmptyValue()
	case string:
		return NewStringValue(val)
	case float64:
		return NewNumberValue(val)
	case bool:
		return NewBooleanValue(val)
	case int:
		return NewNumberValue(float64(val))
	default:
		return NewErrorValue(ErrValue), fmt.Errorf("unsupported value type: %T", val)
	}
}

// CellValueExtractor extracts values from cells for evaluation.
type CellValueExtractor interface {
	// GetCellValue returns the value of a cell.
	GetCellValue(ref spreadsheet.CellRef) (interface{}, error)
}

// DefaultCellValueExtractor implements CellValueExtractor for spreadsheets.
type DefaultCellValueExtractor struct {
	sheet interface{}
}

// NewDefaultCellValueExtractor creates a new cell value extractor.
func NewDefaultCellValueExtractor(sheet interface{}) *DefaultCellValueExtractor {
	return &DefaultCellValueExtractor{sheet: sheet}
}

// GetCellValue returns the value of a cell from the sheet.
func (e *DefaultCellValueExtractor) GetCellValue(ref spreadsheet.CellRef) (interface{}, error) {
	// This is a placeholder - actual implementation would access the sheet
	// and return the cell's value
	return nil, errors.New("cell value extraction not implemented")
}

// CellResolver resolves cell and range references to values.
type CellResolver interface {
	// ResolveCell returns the value of a cell.
	ResolveCell(ref CellRef) (Value, error)

	// ResolveRange returns all values in a range as a 2D array.
	ResolveRange(rng RangeRef) ([][]Value, error)
}

// NewEvalContext creates a new evaluation context.
func NewEvalContext(workbook interface{}, currentSheet string) *EvalContext {
	return &EvalContext{
		Workbook:         workbook,
		CurrentSheet:     currentSheet,
		Resolver:         nil, // Must be set by caller
		FunctionRegistry: NewFunctionRegistry(),
	}
}

// NewEvalContextWithResolver creates a new evaluation context with a custom resolver.
func NewEvalContextWithResolver(workbook interface{}, currentSheet string, resolver CellResolver) *EvalContext {
	return &EvalContext{
		Workbook:         workbook,
		CurrentSheet:     currentSheet,
		Resolver:         resolver,
		FunctionRegistry: NewFunctionRegistry(),
	}
}

// EvaluateFormula evaluates a formula string and returns the result.
func EvaluateFormula(formula string, ctx *EvalContext) (Value, error) {
	if formula == "" {
		return NewEmptyValue(), nil
	}

	// Parse the formula
	expr, err := Parse(formula)
	if err != nil {
		return NewErrorValue(ErrValue), fmt.Errorf("parse error: %w", err)
	}

	// Evaluate the expression
	return expr.Evaluate(ctx)
}

// EvaluateCellFormula evaluates a cell's formula and returns the result.
func EvaluateCellFormula(cellValue interface{}, cellResolver CellResolver) (Value, error) {
	// Get the formula from the cell
	var formula string
	switch v := cellValue.(type) {
	case string:
		formula = v
	case nil:
		return NewEmptyValue(), nil
	default:
		return NewErrorValue(ErrValue), errors.New("cell value is not a string")
	}

	if formula == "" || formula[0] != '=' {
		// Not a formula, return the literal value
		return convertToValue(cellValue), nil
	}

	// Evaluate the formula (strip the leading '=')
	ctx := NewEvalContext(nil, "")
	ctx.Resolver = cellResolver

	result, err := EvaluateFormula(formula[1:], ctx)
	if err != nil {
		return NewErrorValue(ErrValue), err
	}

	return result, nil
}

// convertToValue converts a Go value to a formula Value.
func convertToValue(v interface{}) Value {
	switch val := v.(type) {
	case nil:
		return NewEmptyValue()
	case string:
		return NewStringValue(val)
	case float64:
		return NewNumberValue(val)
	case bool:
		return NewBooleanValue(val)
	case int:
		return NewNumberValue(float64(val))
	default:
		return NewErrorValue(ErrValue), fmt.Errorf("unsupported value type: %T", val)
	}
}

// CellValueExtractor extracts values from cells for evaluation.
type CellValueExtractor interface {
	// GetCellValue returns the value of a cell.
	GetCellValue(ref CellRef) (interface{}, error)
}

// DefaultCellValueExtractor implements CellValueExtractor for spreadsheets.
type DefaultCellValueExtractor struct {
	sheet interface{}
}

// NewDefaultCellValueExtractor creates a new cell value extractor.
func NewDefaultCellValueExtractor(sheet interface{}) *DefaultCellValueExtractor {
	return &DefaultCellValueExtractor{sheet: sheet}
}

// GetCellValue returns the value of a cell from the sheet.
func (e *DefaultCellValueExtractor) GetCellValue(ref CellRef) (interface{}, error) {
	// This is a placeholder - actual implementation would access the sheet
	// and return the cell's value
	return nil, errors.New("cell value extraction not implemented")
}
