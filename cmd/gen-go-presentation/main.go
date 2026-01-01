//nolint:revive // code generator entry point
// Package main provides a code generator for PresentationML elements.
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

	// Collect schemas with metadata
	schemas := make(
		[]*SchemaFileMetadata,
		0,
		len(files),
	)
	for _, f := range files {
		name := f.Name()

		// Filter for Presentation and Drawing schemas
		if !isPresentationSchema(name) &&
			!isDrawingSchema(name) {
			continue
		}

		path := filepath.Join(schemasDir, name)
		metadata, parseErr := parseSchemaFile(
			path,
		)
		if parseErr != nil {
			fmt.Printf(
				"Warning: failed to parse schema %s: %v\n",
				name,
				parseErr,
			)

			continue
		}
		schemas = append(schemas, metadata)
	}

	// Sort schemas: main schemas first (Office2007), then by version
	sortSchemas(schemas)

	fmt.Printf(
		"Found %d schemas to process\n",
		len(schemas),
	)

	// Collect enums and type names before resolving type names.
	for _, schema := range schemas {
		collectEnumNames(schema)
		collectTypeNames(schema)
	}

	// Build enum name map to detect struct naming conflicts
	// This must happen BEFORE collectTypesWithVersion to avoid duplicate type names
	buildEnumNameMap()
	buildConstructorNameMap()

	// Collect types from all schemas
	for _, schema := range schemas {
		collectTypesWithVersion(schema)
	}

	// Generate Code
	generateCode()
}

// isPresentationSchema checks if a schema file is a Presentation schema.
func isPresentationSchema(filename string) bool {
	// Main PresentationML schema
	if filename == "schemas_openxmlformats_org_presentationml_2006_main.json" {
		return true
	}
	// PowerPoint extension schemas (p14, p15, p16, etc.)
	if strings.Contains(
		filename,
		"office_powerpoint",
	) {
		return true
	}
	// PresentationML-specific namespaces
	if strings.Contains(
		filename,
		"presentationml",
	) {
		return true
	}

	return false
}

// isDrawingSchema checks if a schema file is a Drawing schema relevant to Presentation.
func isDrawingSchema(filename string) bool {
	// Main DrawingML schema
	if strings.Contains(
		filename,
		"drawingml_2006",
	) {
		return true
	}
	// Drawing extension schemas (a14, a15, a16, etc.)
	if strings.Contains(
		filename,
		"office_drawing",
	) {
		return true
	}

	return false
}

// sortSchemas sorts schemas by version: main first, then chronological.
func sortSchemas(schemas []*SchemaFileMetadata) {
	// Sort by version (Office2007 comes first, then ascending)
	for i := 0; i < len(schemas); i++ {
		for j := i + 1; j < len(schemas); j++ {
			if schemas[i].Version > schemas[j].Version {
				schemas[i], schemas[j] = schemas[j], schemas[i]
			}
		}
	}
}
