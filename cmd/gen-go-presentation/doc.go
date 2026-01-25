//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package gen-go-presentation provides code generation for PresentationML elements
// from XML schema definitions.
package gen-go-presentation

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type (
	name string
	format string
)

// CodeGenerationConfig holds configuration for code generation.
type CodeGenerationConfig struct {
	// FileMode is the file permission mode for generated files.
	FileMode os.FileMode
}

// newCodeGenerationConfig creates a default configuration with executable permissions.
func newCodeGenerationConfig() *CodeGenerationConfig {
	return &CodeGenerationConfig{
		FileMode: 0755,
	}
}

// main is the entry point for the generator.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config> [options]\n", os.Args[0])
		os.Exit(1)
	}

	// Determine output directory
	outDir := "."
	if len(os.Args) > 2 {
		outDir = os.Args[1]
	}

	// Read template file
	templateFile, err := gen_go_presentation_template.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading template: %v\n", err)
		os.Exit(1)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	// Parse schema definitions
	schemaFile := os.Args[0]
	schema, err := gen_go_presentation_template.LoadSchema()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing schema: %v\n", err)
		os.Exit(1)
	}

	// Create generator
	config := newCodeGenerationConfig()
	gen := gen_go_presentation.NewGenerator(config, schema)

	// Generate all element types
	if err := gen.GenerateAll(); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d element types\n", gen.Count())
}