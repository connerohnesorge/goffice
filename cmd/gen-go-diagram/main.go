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
	const initialCapacity = 10
	schemas := make(
		[]*SchemaFileMetadata,
		0,
		initialCapacity,
	)
	for _, f := range files {
		name := f.Name()

		// Filter for Diagram schemas and DrawingML main (for shared enums)
		if !isDiagramSchema(name) &&
			!isDrawingMLMainSchema(name) {
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
		"Found %d diagram schemas to process\n",
		len(schemas),
	)

	// Collect types from all schemas
	for _, schema := range schemas {
		collectTypesWithVersion(schema)
	}

	// Generate Code
	generateCode()
}

// isDiagramSchema checks if a schema file is a Diagram schema.
func isDiagramSchema(filename string) bool {
	// Main Diagram schema
	if filename == "schemas_openxmlformats_org_drawingml_2006_diagram.json" {
		return true
	}
	// Diagram extension schemas (diagram namespace, various versions)
	if strings.Contains(filename, "diagram") {
		return true
	}

	return false
}

// isDrawingMLMainSchema checks if a schema file is the DrawingML main schema.
// We need this to get shared enum definitions used by diagram elements.
func isDrawingMLMainSchema(filename string) bool {
	return filename == "schemas_openxmlformats_org_drawingml_2006_main.json"
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
