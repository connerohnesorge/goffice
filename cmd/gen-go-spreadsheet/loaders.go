// Package main provides a code generator for SpreadsheetML elements.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// collectDrawingMLTypes scans DrawingML files to identify implemented types.
// It populates drawingMLTypes map with type names found in drawingml/*.go.
// This prevents the generator from creating duplicates.
func collectDrawingMLTypes() {
	files, _ := filepath.Glob("drawingml/*.go")
	for _, f := range files {
		data, errRead := os.ReadFile(f)
		if errRead != nil {
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if !strings.HasPrefix(line, "type ") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				name := parts[1]
				drawingMLTypes[name] = true
			}
		}
	}
	fmt.Printf(
		"Found %d DrawingML types\n",
		len(drawingMLTypes),
	)
}

// collectExistingTypes identifies types already in hand-written files.
// It populates existingTypes map to avoid overwriting manual implementations.
// It skips elements.go and enums.go as they are generated.
func collectExistingTypes() {
	files, _ := filepath.Glob(
		"spreadsheet/elements/*.go",
	)
	for _, f := range files {
		base := filepath.Base(f)
		if base == "elements.go" ||
			base == "enums.go" {
			continue
		}
		data, errRead := os.ReadFile(f)
		if errRead != nil {
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			// Mark existing types
			if !strings.HasPrefix(line, "type ") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				name := parts[1]
				existingTypes[name] = true
			}
		}
	}
	fmt.Printf(
		"Found %d existing types to skip\n",
		len(existingTypes),
	)
}

// loadNamespaces reads the canonical namespace mappings from JSON.
// It populates namespaceMap used for resolving prefixes to URIs.
func loadNamespaces() {
	path := "Open-XML-SDK/data/namespaces.json"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(
			fmt.Errorf(
				"failed to read namespaces.json: %w",
				err,
			),
		)
	}
	var entries []NamespaceEntry
	if errUnmarshal := json.Unmarshal(data, &entries); errUnmarshal != nil {
		panic(
			fmt.Errorf(
				"failed to unmarshal namespaces: %w",
				errUnmarshal,
			),
		)
	}
	for _, e := range entries {
		namespaceMap[e.Prefix] = e.Uri
	}
}

// collectTypes extracts type metadata from a schema main.json file.
// It handles both complex types and enumeration definitions.
func collectTypes(path string) {
	data, errRead := os.ReadFile(path)
	if errRead != nil {
		msg := "failed to read schema file %s: %w"
		panic(fmt.Errorf(msg, path, errRead))
	}

	var schema SchemaFile
	if errUnmarshal := json.Unmarshal(data, &schema); errUnmarshal != nil {
		msg := "failed to unmarshal schema %s: %w"
		panic(fmt.Errorf(msg, path, errUnmarshal))
	}

	for _, t := range schema.Types {
		t.TargetNamespace = schema.TargetNamespace
		if t.ClassName != "" {
			typeMap[t.Name] = TypeInfo{
				ClassName: t.ClassName,
				Namespace: t.TargetNamespace,
			}
		}
	}

	for _, e := range schema.Enums {
		e.TargetNamespace = schema.TargetNamespace
		enumMap[e.Name] = e
	}
}
