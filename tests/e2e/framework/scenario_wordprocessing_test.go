package framework

import (
	"path/filepath"
	"testing"
)

func TestLoadWordprocessingScenarios(
	t *testing.T,
) {
	scenarioDir := filepath.Join(
		"..",
		"scenarios",
		"wordprocessing",
	)

	scenarios, err := LoadScenariosFromDirectory(
		scenarioDir,
	)
	if err != nil {
		t.Fatalf(
			"Failed to load wordprocessing scenarios: %v",
			err,
		)
	}

	if len(scenarios) == 0 {
		t.Fatal(
			"No scenarios were loaded from wordprocessing directory",
		)
	}

	t.Logf(
		"Successfully loaded %d wordprocessing scenario(s)",
		len(scenarios),
	)

	// Verify each scenario
	expectedScenarios := map[string]struct{}{
		"basic-paragraph":     {},
		"formatted-text":      {},
		"multiple-paragraphs": {},
		"simple-table":        {},
		"mixed-content":       {},
	}

	for _, s := range scenarios {
		t.Logf(
			"  - %s: %s",
			s.Name,
			s.Description,
		)

		// Verify scenario name is expected
		if _, ok := expectedScenarios[s.Name]; !ok {
			t.Errorf(
				"Unexpected scenario name: %s",
				s.Name,
			)
		}

		// Verify required fields
		if s.Name == "" {
			t.Errorf("Scenario has empty name")
		}
		if s.Description == "" {
			t.Errorf(
				"Scenario %s has empty description",
				s.Name,
			)
		}
		if s.DocumentType != "wordprocessing" {
			t.Errorf(
				"Scenario %s has wrong document_type: %s",
				s.Name,
				s.DocumentType,
			)
		}
		if s.Category == "" {
			t.Errorf(
				"Scenario %s has empty category",
				s.Name,
			)
		}
		if len(s.Tags) == 0 {
			t.Errorf(
				"Scenario %s has no tags",
				s.Name,
			)
		}
		if len(s.Operations) == 0 {
			t.Errorf(
				"Scenario %s has no operations",
				s.Name,
			)
		}

		// Verify first operation is create_document
		if s.Operations[0].Action != "create_document" {
			t.Errorf(
				"Scenario %s first operation is not create_document: %s",
				s.Name,
				s.Operations[0].Action,
			)
		}

		// Verify tolerance defaults are set
		if s.Tolerance.VisualPixelTolerance <= 0 {
			t.Errorf(
				"Scenario %s has invalid visual_pixel_tolerance: %f",
				s.Name,
				s.Tolerance.VisualPixelTolerance,
			)
		}
		if s.Tolerance.VisualDiffThreshold <= 0 {
			t.Errorf(
				"Scenario %s has invalid visual_diff_threshold: %f",
				s.Name,
				s.Tolerance.VisualDiffThreshold,
			)
		}
	}

	// Verify we found all expected scenarios
	if len(scenarios) != len(expectedScenarios) {
		t.Errorf(
			"Expected %d scenarios, got %d",
			len(expectedScenarios),
			len(scenarios),
		)
	}
}
