// Package main provides a code generator for WordprocessingML elements.
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
	// Find all Go files in the drawingml directory.
	files, _ := filepath.Glob("drawingml/*.go")
	for _, f := range files {
		data, errRead := os.ReadFile(f)
		if errRead != nil {
			continue
		}
		// Parse each file line by line to find type declarations.
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if !strings.HasPrefix(line, "type ") {
				continue
			}
			// Extract the type name from the declaration.
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
		"wordprocessing/elements/*.go",
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

// collectTypesWithVersion extracts type metadata from a schema file with version information.
// It handles both complex types and enumeration definitions, storing version metadata.
func collectTypesWithVersion(
	metadata *SchemaFileMetadata,
) {
	data, errRead := os.ReadFile(metadata.Path)
	if errRead != nil {
		msg := "failed to read schema file %s: %w"
		panic(
			fmt.Errorf(
				msg,
				metadata.Path,
				errRead,
			),
		)
	}

	var schema SchemaFile
	if errUnmarshal := json.Unmarshal(data, &schema); errUnmarshal != nil {
		msg := "failed to unmarshal schema %s: %w"
		panic(
			fmt.Errorf(
				msg,
				metadata.Path,
				errUnmarshal,
			),
		)
	}

	// Register namespace in global namespace map
	if metadata.NamespacePrefix != "" &&
		metadata.Namespace != "" {
		namespaceMap[metadata.NamespacePrefix] = metadata.Namespace
	}

	// Collect types with version metadata
	for i := range schema.Types {
		t := &schema.Types[i]
		t.TargetNamespace = schema.TargetNamespace
		if t.ClassName != "" {
			typeMap[t.Name] = TypeInfo{
				ClassName: t.ClassName,
				Namespace: t.TargetNamespace,
				Version:   metadata.Version,
			}
		}
	}

	// Collect enums with version metadata
	for i := range schema.Enums {
		e := &schema.Enums[i]
		e.TargetNamespace = schema.TargetNamespace
		enumMap[e.Name] = *e
	}
}
