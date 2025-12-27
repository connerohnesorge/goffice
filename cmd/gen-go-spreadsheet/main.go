// Package main provides a code generator for SpreadsheetML elements.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	collectDrawingMLTypes()
	collectExistingTypes()
	loadNamespaces()

	schemasDir := "Open-XML-SDK/data/schemas"
	files, err := os.ReadDir(schemasDir)
	if err != nil {
		panic(
			fmt.Errorf(
				"failed to read schemas dir: %w",
				err,
			),
		)
	}

	// First pass: Collect types from all main.json schemas
	for _, f := range files {
		if strings.Contains(
			f.Name(),
			"main.json",
		) {
			collectTypes(
				filepath.Join(
					schemasDir,
					f.Name(),
				),
			)
		}
	}

	// Generate Code
	generateCode()
}
