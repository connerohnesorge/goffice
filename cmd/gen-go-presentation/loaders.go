//nolint:revive,gocritic // code generator with complex loading logic
// Package main provides a code generator for PresentationML elements.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
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
		"presentation/elements/*.go",
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
			className := resolveStructName(t)
			t.ClassName = className
			typeMap[t.Name] = TypeInfo{
				ClassName: className,
				Namespace: t.TargetNamespace,
				Version:   metadata.Version,
			}
		}
	}

	// Collect enums with version metadata
	for i := range schema.Enums {
		e := &schema.Enums[i]
		e.TargetNamespace = schema.TargetNamespace
		enumMap[e.Name] = e
	}
}

// collectEnumNames collects enum definitions before type processing.
func collectEnumNames(
	metadata *SchemaFileMetadata,
) {
	data, errRead := os.ReadFile(metadata.Path)
	if errRead != nil {
		return
	}

	var schema SchemaFile
	if errUnmarshal := json.Unmarshal(data, &schema); errUnmarshal != nil {
		return
	}

	for i := range schema.Enums {
		e := &schema.Enums[i]
		e.TargetNamespace = schema.TargetNamespace
		enumMap[e.Name] = e
	}
}

// collectTypeNames collects class names for constructor conflict detection.
func collectTypeNames(
	metadata *SchemaFileMetadata,
) {
	data, errRead := os.ReadFile(metadata.Path)
	if errRead != nil {
		return
	}

	var schema SchemaFile
	if errUnmarshal := json.Unmarshal(data, &schema); errUnmarshal != nil {
		return
	}

	for i := range schema.Types {
		t := &schema.Types[i]
		if t.ClassName != "" {
			typeNameMap[t.ClassName] = true
		}
	}
}

// loadExistingEnumNames adds existing enum type names and constant names from enums.go to enumNameMap.
// This prevents struct types from conflicting with existing enum types and constants.
func loadExistingEnumNames() {
	data, errRead := os.ReadFile("presentation/elements/enums.go")
	if errRead != nil {
		return
	}

	count := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Look for "type <Name> string" pattern
		if strings.HasPrefix(line, "type ") && strings.Contains(line, " string") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				name := parts[1]
				enumNameMap[name] = true
				count++
			}
		}

		// Look for constant values like "ConstantName <Type> = ..."
		// Only inside const blocks - check if the line has an assignment
		if strings.Contains(trimmed, " = ") && !strings.HasPrefix(trimmed, "//") &&
			!strings.HasPrefix(trimmed, "/*") && !strings.HasPrefix(trimmed, "*") {
			// Extract the constant name (first field before the type)
			parts := strings.Fields(trimmed)
			if len(parts) >= 1 && parts[0] != "" {
				constantName := parts[0]
				// Avoid false positives - constants start with uppercase
				if constantName != "" && constantName[0] >= 'A' && constantName[0] <= 'Z' {
					enumNameMap[constantName] = true
					count++
				}
			}
		}
	}
	fmt.Printf("Loaded %d existing enum names/constants\n", count)
}

// buildEnumNameMap populates enumNameMap from all enums to detect naming conflicts.
// Call this after all schemas are loaded with collectTypesWithVersion().
func buildEnumNameMap() {
	// First, add existing enum names from enums.go
	loadExistingEnumNames()

	// Then add enums from schemas
	for name, enum := range enumMap {
		enumNameMap[name] = true
		if enum == nil {
			continue
		}
		enumNameMap[enum.ClassName] = true
		for _, facet := range enum.Facets {
			constantName := enumConstantName(
				enum.Name,
				facet.Value,
			)
			enumNameMap[constantName] = true
		}
	}
}

// buildConstructorNameMap builds reserved constructor names from schema types.
func buildConstructorNameMap() {
	for name := range typeNameMap {
		constructorNameMap["New"+name] = true
	}
}

func enumConstantName(enumName, facetValue string) string {
	name := toPascalCase(facetValue)
	if name == "" {
		name = "None"
	}
	if unicode.IsDigit(rune(name[0])) {
		name = enumName + name
	}

	return enumName + name
}

func resolveStructName(t *SchemaType) string {
	original := t.ClassName
	if original == "" {
		return ""
	}

	if existingTypes[original] {
		usedStructNames[original] = true
		constructorNameMap["New"+original] = true

		return original
	}

	prefix := ""
	candidate := original
	if hasNameConflict(candidate, original) {
		prefix = deriveStructPrefix(t.TargetNamespace)
		candidate = prefix + original
	}

	suffix := 1
	for hasNameConflict(candidate, original) {
		candidate = fmt.Sprintf(
			"%s%s%d",
			prefix,
			original,
			suffix,
		)
		suffix++
	}

	usedStructNames[candidate] = true
	constructorNameMap["New"+candidate] = true
	if candidate != original {
		renamedTypes[t.Name] = candidate
	}

	return candidate
}

func hasNameConflict(candidate, original string) bool {
	if enumNameMap[candidate] ||
		constructorNameMap[candidate] ||
		usedStructNames[candidate] {
		return true
	}
	if candidate != original && existingTypes[candidate] {
		return true
	}
	if typeNameMap["New"+candidate] ||
		existingTypes["New"+candidate] ||
		enumNameMap["New"+candidate] ||
		usedStructNames["New"+candidate] {
		return true
	}

	return false
}
