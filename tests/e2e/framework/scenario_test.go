package framework

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadScenario(t *testing.T) {
	t.Run(
		"Valid scenario file",
		func(t *testing.T) {
			// Create temp directory for test
			tmpDir := t.TempDir()

			// Create a test scenario YAML file
			scenarioYAML := `name: test-scenario
description: Test scenario for unit testing
document_type: wordprocessing
category: basic
tags:
  - test
  - unit
tolerance:
  xml_attribute_order_sensitive: false
  visual_pixel_tolerance: 2.5
  visual_diff_threshold: 0.01
operations:
  - action: create_document
    doc_type: document
  - action: add_paragraph
    text: Hello World
`
			scenarioPath := filepath.Join(
				tmpDir,
				"test.yaml",
			)
			if err := os.WriteFile(scenarioPath, []byte(scenarioYAML), 0644); err != nil {
				t.Fatalf(
					"Failed to write test scenario: %v",
					err,
				)
			}

			// Load the scenario
			scenario, err := LoadScenario(
				scenarioPath,
			)
			if err != nil {
				t.Fatalf(
					"LoadScenario failed: %v",
					err,
				)
			}

			// Verify fields
			if scenario.Name != "test-scenario" {
				t.Errorf(
					"Expected name 'test-scenario', got %q",
					scenario.Name,
				)
			}
			if scenario.Description != "Test scenario for unit testing" {
				t.Errorf(
					"Expected description 'Test scenario for unit testing', got %q",
					scenario.Description,
				)
			}
			if scenario.DocumentType != "wordprocessing" {
				t.Errorf(
					"Expected document_type 'wordprocessing', got %q",
					scenario.DocumentType,
				)
			}
			if scenario.Category != "basic" {
				t.Errorf(
					"Expected category 'basic', got %q",
					scenario.Category,
				)
			}
			if len(scenario.Tags) != 2 {
				t.Errorf(
					"Expected 2 tags, got %d",
					len(scenario.Tags),
				)
			}
			if scenario.Tolerance.XMLAttributeOrderSensitive != false {
				t.Errorf(
					"Expected xml_attribute_order_sensitive false, got true",
				)
			}
			if scenario.Tolerance.VisualPixelTolerance != 2.5 {
				t.Errorf(
					"Expected visual_pixel_tolerance 2.5, got %f",
					scenario.Tolerance.VisualPixelTolerance,
				)
			}
			if scenario.Tolerance.VisualDiffThreshold != 0.01 {
				t.Errorf(
					"Expected visual_diff_threshold 0.01, got %f",
					scenario.Tolerance.VisualDiffThreshold,
				)
			}
			if len(scenario.Operations) != 2 {
				t.Errorf(
					"Expected 2 operations, got %d",
					len(scenario.Operations),
				)
			}
			if scenario.Operations[0].Action != "create_document" {
				t.Errorf(
					"Expected first operation action 'create_document', got %q",
					scenario.Operations[0].Action,
				)
			}
			if scenario.Operations[1].Action != "add_paragraph" {
				t.Errorf(
					"Expected second operation action 'add_paragraph', got %q",
					scenario.Operations[1].Action,
				)
			}
		},
	)

	t.Run(
		"Default tolerance values",
		func(t *testing.T) {
			tmpDir := t.TempDir()

			// Create scenario without tolerance config
			scenarioYAML := `name: defaults-test
description: Test default tolerance values
document_type: wordprocessing
category: basic
tags: [test]
operations:
  - action: create_document
`
			scenarioPath := filepath.Join(
				tmpDir,
				"defaults.yaml",
			)
			if err := os.WriteFile(scenarioPath, []byte(scenarioYAML), 0644); err != nil {
				t.Fatalf(
					"Failed to write test scenario: %v",
					err,
				)
			}

			scenario, err := LoadScenario(
				scenarioPath,
			)
			if err != nil {
				t.Fatalf(
					"LoadScenario failed: %v",
					err,
				)
			}

			// Verify defaults are set
			if scenario.Tolerance.VisualPixelTolerance != 2.0 {
				t.Errorf(
					"Expected default visual_pixel_tolerance 2.0, got %f",
					scenario.Tolerance.VisualPixelTolerance,
				)
			}
			if scenario.Tolerance.VisualDiffThreshold != 0.005 {
				t.Errorf(
					"Expected default visual_diff_threshold 0.005, got %f",
					scenario.Tolerance.VisualDiffThreshold,
				)
			}
		},
	)

	t.Run(
		"Non-existent file",
		func(t *testing.T) {
			_, err := LoadScenario(
				"/nonexistent/file.yaml",
			)
			if err == nil {
				t.Error(
					"Expected error for non-existent file, got nil",
				)
			}
		},
	)

	t.Run("Invalid YAML", func(t *testing.T) {
		tmpDir := t.TempDir()

		invalidYAML := `this is not: valid: yaml: content`
		scenarioPath := filepath.Join(
			tmpDir,
			"invalid.yaml",
		)
		if err := os.WriteFile(scenarioPath, []byte(invalidYAML), 0644); err != nil {
			t.Fatalf(
				"Failed to write invalid YAML: %v",
				err,
			)
		}

		_, err := LoadScenario(scenarioPath)
		if err == nil {
			t.Error(
				"Expected error for invalid YAML, got nil",
			)
		}
	})
}

func TestLoadScenariosFromDirectory(
	t *testing.T,
) {
	t.Run(
		"Multiple scenario files",
		func(t *testing.T) {
			tmpDir := t.TempDir()

			// Create multiple scenario files
			scenarios := []struct {
				filename string
				name     string
			}{
				{
					"scenario1.yaml",
					"scenario-one",
				},
				{
					"scenario2.yaml",
					"scenario-two",
				},
				{
					"scenario3.yml",
					"scenario-three",
				}, // Test .yml extension
			}

			for _, s := range scenarios {
				content := `name: ` + s.name + `
description: Test scenario
document_type: wordprocessing
category: basic
tags: [test]
operations:
  - action: create_document
`
				path := filepath.Join(
					tmpDir,
					s.filename,
				)
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf(
						"Failed to write scenario %s: %v",
						s.filename,
						err,
					)
				}
			}

			// Create a non-YAML file (should be ignored)
			if err := os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("not a scenario"), 0644); err != nil {
				t.Fatalf(
					"Failed to write non-YAML file: %v",
					err,
				)
			}

			// Load scenarios from directory
			loaded, err := LoadScenariosFromDirectory(
				tmpDir,
			)
			if err != nil {
				t.Fatalf(
					"LoadScenariosFromDirectory failed: %v",
					err,
				)
			}

			// Should have loaded 3 scenarios (ignoring readme.txt)
			if len(loaded) != 3 {
				t.Errorf(
					"Expected 3 scenarios, got %d",
					len(loaded),
				)
			}

			// Verify scenario names
			names := make(map[string]bool)
			for _, scenario := range loaded {
				names[scenario.Name] = true
			}

			for _, s := range scenarios {
				if !names[s.name] {
					t.Errorf(
						"Expected to find scenario %q, but it was not loaded",
						s.name,
					)
				}
			}
		},
	)

	t.Run("Empty directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		scenarios, err := LoadScenariosFromDirectory(
			tmpDir,
		)
		if err != nil {
			t.Fatalf(
				"LoadScenariosFromDirectory failed: %v",
				err,
			)
		}

		if len(scenarios) != 0 {
			t.Errorf(
				"Expected 0 scenarios from empty directory, got %d",
				len(scenarios),
			)
		}
	})

	t.Run(
		"Non-existent directory",
		func(t *testing.T) {
			_, err := LoadScenariosFromDirectory(
				"/nonexistent/directory",
			)
			if err == nil {
				t.Error(
					"Expected error for non-existent directory, got nil",
				)
			}
		},
	)

	t.Run(
		"Directory with subdirectories",
		func(t *testing.T) {
			tmpDir := t.TempDir()

			// Create a scenario file
			scenarioYAML := `name: root-scenario
description: Scenario in root
document_type: wordprocessing
category: basic
tags: [test]
operations:
  - action: create_document
`
			if err := os.WriteFile(filepath.Join(tmpDir, "root.yaml"), []byte(scenarioYAML), 0644); err != nil {
				t.Fatalf(
					"Failed to write root scenario: %v",
					err,
				)
			}

			// Create a subdirectory with another scenario (should be ignored)
			subDir := filepath.Join(
				tmpDir,
				"subdir",
			)
			if err := os.Mkdir(subDir, 0755); err != nil {
				t.Fatalf(
					"Failed to create subdirectory: %v",
					err,
				)
			}
			if err := os.WriteFile(filepath.Join(subDir, "sub.yaml"), []byte(scenarioYAML), 0644); err != nil {
				t.Fatalf(
					"Failed to write subdirectory scenario: %v",
					err,
				)
			}

			scenarios, err := LoadScenariosFromDirectory(
				tmpDir,
			)
			if err != nil {
				t.Fatalf(
					"LoadScenariosFromDirectory failed: %v",
					err,
				)
			}

			// Should only load scenario from root directory
			if len(scenarios) != 1 {
				t.Errorf(
					"Expected 1 scenario (subdirectories should be ignored), got %d",
					len(scenarios),
				)
			}
			if scenarios[0].Name != "root-scenario" {
				t.Errorf(
					"Expected scenario name 'root-scenario', got %q",
					scenarios[0].Name,
				)
			}
		},
	)
}

func TestOperationProperties(t *testing.T) {
	t.Run(
		"Nested operation properties",
		func(t *testing.T) {
			tmpDir := t.TempDir()

			// Create scenario with nested operations (runs within paragraph)
			scenarioYAML := `name: nested-ops
description: Test nested operations
document_type: wordprocessing
category: basic
tags: [test]
operations:
  - action: add_paragraph
    operations:
      - action: add_run
        text: "Bold text"
        properties:
          bold: true
          color: "FF0000"
      - action: add_run
        text: "Normal text"
`
			scenarioPath := filepath.Join(
				tmpDir,
				"nested.yaml",
			)
			if err := os.WriteFile(scenarioPath, []byte(scenarioYAML), 0644); err != nil {
				t.Fatalf(
					"Failed to write test scenario: %v",
					err,
				)
			}

			scenario, err := LoadScenario(
				scenarioPath,
			)
			if err != nil {
				t.Fatalf(
					"LoadScenario failed: %v",
					err,
				)
			}

			// Verify the nested structure
			if len(scenario.Operations) != 1 {
				t.Fatalf(
					"Expected 1 top-level operation, got %d",
					len(scenario.Operations),
				)
			}

			paraOp := scenario.Operations[0]
			if paraOp.Action != "add_paragraph" {
				t.Errorf(
					"Expected action 'add_paragraph', got %q",
					paraOp.Action,
				)
			}

			// Check nested operations
			nestedOps, ok := paraOp.Properties["operations"].([]interface{})
			if !ok {
				t.Fatalf(
					"Expected 'operations' property to be a slice",
				)
			}

			if len(nestedOps) != 2 {
				t.Errorf(
					"Expected 2 nested operations, got %d",
					len(nestedOps),
				)
			}

			// Verify first nested operation
			firstOp, ok := nestedOps[0].(map[string]interface{})
			if !ok {
				t.Fatalf(
					"Expected first nested operation to be a map",
				)
			}

			if action, _ := firstOp["action"].(string); action != "add_run" {
				t.Errorf(
					"Expected nested action 'add_run', got %q",
					action,
				)
			}

			if text, _ := firstOp["text"].(string); text != "Bold text" {
				t.Errorf(
					"Expected text 'Bold text', got %q",
					text,
				)
			}

			props, ok := firstOp["properties"].(map[string]interface{})
			if !ok {
				t.Fatalf(
					"Expected 'properties' to be a map",
				)
			}

			if bold, _ := props["bold"].(bool); !bold {
				t.Error(
					"Expected bold property to be true",
				)
			}

			if color, _ := props["color"].(string); color != "FF0000" {
				t.Errorf(
					"Expected color 'FF0000', got %q",
					color,
				)
			}
		},
	)
}
