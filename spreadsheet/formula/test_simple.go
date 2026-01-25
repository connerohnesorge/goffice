//go:build ignore
// +build ignore

package main

import (
	"fmt"

	formula "github.com/connerohnesorge/goffice/spreadsheet"
)

func main() {
	// Test basic evaluation
	result, err := formula.EvaluateFormula("=1+2", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
