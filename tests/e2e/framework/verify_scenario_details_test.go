package framework

import (
	"path/filepath"
	"testing"
)

func TestFormattedTextScenarioDetails(
	t *testing.T,
) {
	scenarioPath := filepath.Join(
		"..",
		"scenarios",
		"wordprocessing",
		"formatted-text.yaml",
	)

	scenario, err := LoadScenario(scenarioPath)
	if err != nil {
		t.Fatalf(
			"Failed to load formatted-text scenario: %v",
			err,
		)
	}

	// Verify basic metadata
	if scenario.Name != "formatted-text" {
		t.Errorf(
			"Expected name 'formatted-text', got %s",
			scenario.Name,
		)
	}
	if scenario.DocumentType != "wordprocessing" {
		t.Errorf(
			"Expected document_type 'wordprocessing', got %s",
			scenario.DocumentType,
		)
	}
	if scenario.Category != "basic" {
		t.Errorf(
			"Expected category 'basic', got %s",
			scenario.Category,
		)
	}

	// Verify tags
	expectedTags := []string{
		"text-formatting",
		"runs",
		"properties",
	}
	if len(scenario.Tags) != len(expectedTags) {
		t.Errorf(
			"Expected %d tags, got %d",
			len(expectedTags),
			len(scenario.Tags),
		)
	}

	// Verify custom tolerance values
	if scenario.Tolerance.VisualPixelTolerance != 2.0 {
		t.Errorf(
			"Expected visual_pixel_tolerance 2.0, got %f",
			scenario.Tolerance.VisualPixelTolerance,
		)
	}
	if scenario.Tolerance.VisualDiffThreshold != 0.005 {
		t.Errorf(
			"Expected visual_diff_threshold 0.005, got %f",
			scenario.Tolerance.VisualDiffThreshold,
		)
	}

	// Verify operations
	if len(scenario.Operations) != 2 {
		t.Fatalf(
			"Expected 2 operations, got %d",
			len(scenario.Operations),
		)
	}

	// First operation: create_document
	if scenario.Operations[0].Action != "create_document" {
		t.Errorf(
			"Expected first operation 'create_document', got %s",
			scenario.Operations[0].Action,
		)
	}

	// Second operation: add_paragraph with nested operations
	if scenario.Operations[1].Action != "add_paragraph" {
		t.Errorf(
			"Expected second operation 'add_paragraph', got %s",
			scenario.Operations[1].Action,
		)
	}

	// Verify nested operations exist
	nestedOps, ok := scenario.Operations[1].Properties["operations"].([]interface{})
	if !ok {
		t.Fatal(
			"Expected 'operations' property in add_paragraph",
		)
	}

	if len(nestedOps) != 4 {
		t.Errorf(
			"Expected 4 nested operations (runs), got %d",
			len(nestedOps),
		)
	}

	// Verify first run (normal text)
	run1, ok := nestedOps[0].(map[string]interface{})
	if !ok {
		t.Fatal(
			"Expected first nested operation to be a map",
		)
	}
	if action, _ := run1["action"].(string); action != "add_run" {
		t.Errorf(
			"Expected action 'add_run', got %s",
			action,
		)
	}
	if text, _ := run1["text"].(string); text != "Normal text " {
		t.Errorf(
			"Expected text 'Normal text ', got %s",
			text,
		)
	}

	// Verify second run (bold)
	run2, ok := nestedOps[1].(map[string]interface{})
	if !ok {
		t.Fatal(
			"Expected second nested operation to be a map",
		)
	}
	if text, _ := run2["text"].(string); text != "bold text " {
		t.Errorf(
			"Expected text 'bold text ', got %s",
			text,
		)
	}
	props2, ok := run2["properties"].(map[string]interface{})
	if !ok {
		t.Fatal(
			"Expected properties in second run",
		)
	}
	if bold, _ := props2["bold"].(bool); !bold {
		t.Error(
			"Expected bold property to be true",
		)
	}

	// Verify fourth run (colored and bold)
	run4, ok := nestedOps[3].(map[string]interface{})
	if !ok {
		t.Fatal(
			"Expected fourth nested operation to be a map",
		)
	}
	if text, _ := run4["text"].(string); text != "colored" {
		t.Errorf(
			"Expected text 'colored', got %s",
			text,
		)
	}
	props4, ok := run4["properties"].(map[string]interface{})
	if !ok {
		t.Fatal(
			"Expected properties in fourth run",
		)
	}
	if color, _ := props4["color"].(string); color != "FF0000" {
		t.Errorf(
			"Expected color 'FF0000', got %s",
			color,
		)
	}
	if bold, _ := props4["bold"].(bool); !bold {
		t.Error(
			"Expected bold property to be true",
		)
	}

	t.Log(
		"formatted-text scenario parsed correctly with all nested operations",
	)
}

func TestSimpleTableScenarioDetails(
	t *testing.T,
) {
	scenarioPath := filepath.Join(
		"..",
		"scenarios",
		"wordprocessing",
		"simple-table.yaml",
	)

	scenario, err := LoadScenario(scenarioPath)
	if err != nil {
		t.Fatalf(
			"Failed to load simple-table scenario: %v",
			err,
		)
	}

	// Verify operations
	if len(scenario.Operations) != 2 {
		t.Fatalf(
			"Expected 2 operations, got %d",
			len(scenario.Operations),
		)
	}

	// Second operation: add_table
	if scenario.Operations[1].Action != "add_table" {
		t.Errorf(
			"Expected second operation 'add_table', got %s",
			scenario.Operations[1].Action,
		)
	}

	// Verify table dimensions
	rows, ok := scenario.Operations[1].Properties["rows"].(int)
	if !ok {
		t.Fatal(
			"Expected 'rows' property in add_table",
		)
	}
	if rows != 2 {
		t.Errorf("Expected 2 rows, got %d", rows)
	}

	cols, ok := scenario.Operations[1].Properties["cols"].(int)
	if !ok {
		t.Fatal(
			"Expected 'cols' property in add_table",
		)
	}
	if cols != 3 {
		t.Errorf("Expected 3 cols, got %d", cols)
	}

	// Verify cells
	cells, ok := scenario.Operations[1].Properties["cells"].([]interface{})
	if !ok {
		t.Fatal(
			"Expected 'cells' property in add_table",
		)
	}
	if len(cells) != 6 {
		t.Errorf(
			"Expected 6 cells, got %d",
			len(cells),
		)
	}

	// Verify first cell
	cell1, ok := cells[0].(map[string]interface{})
	if !ok {
		t.Fatal("Expected first cell to be a map")
	}
	if row, _ := cell1["row"].(int); row != 0 {
		t.Errorf(
			"Expected cell row 0, got %d",
			row,
		)
	}
	if col, _ := cell1["col"].(int); col != 0 {
		t.Errorf(
			"Expected cell col 0, got %d",
			col,
		)
	}
	if text, _ := cell1["text"].(string); text != "Header 1" {
		t.Errorf(
			"Expected cell text 'Header 1', got %s",
			text,
		)
	}

	t.Log(
		"simple-table scenario parsed correctly with all cell definitions",
	)
}
