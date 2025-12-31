package framework

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// TestScenario represents a complete test scenario
type TestScenario struct {
	Name         string          `yaml:"name"`
	Description  string          `yaml:"description"`
	DocumentType string          `yaml:"document_type"` // wordprocessing, spreadsheet, presentation
	Category     string          `yaml:"category"`      // basic, complex, edge-case
	Tags         []string        `yaml:"tags"`
	Tolerance    ToleranceConfig `yaml:"tolerance"`
	Operations   []Operation     `yaml:"operations"`
}

// Operation represents a single document operation
type Operation struct {
	Action     string                 `yaml:"action"` // create_document, add_paragraph, add_run, etc.
	Properties map[string]interface{} `yaml:",inline"`
}

// ToleranceConfig configures comparison tolerances
type ToleranceConfig struct {
	XMLAttributeOrderSensitive bool    `yaml:"xml_attribute_order_sensitive"`
	VisualPixelTolerance       float64 `yaml:"visual_pixel_tolerance"`
	VisualDiffThreshold        float64 `yaml:"visual_diff_threshold"`
}

// LoadScenario loads a scenario from YAML file
func LoadScenario(
	path string,
) (scenario *TestScenario, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(
			"open scenario file: %w",
			err,
		)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil &&
			err == nil {
			err = fmt.Errorf(
				"close scenario file: %w",
				closeErr,
			)
		}
	}()

	var s TestScenario
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&s); err != nil {
		return nil, fmt.Errorf(
			"parse scenario YAML: %w",
			err,
		)
	}

	// Set defaults
	if s.Tolerance.VisualPixelTolerance == 0 {
		s.Tolerance.VisualPixelTolerance = 2.0
	}
	if s.Tolerance.VisualDiffThreshold == 0 {
		s.Tolerance.VisualDiffThreshold = 0.005
	}

	return &s, nil
}

// LoadScenariosFromDirectory loads all scenarios in a directory
func LoadScenariosFromDirectory(
	dirPath string,
) ([]*TestScenario, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf(
			"read scenarios directory: %w",
			err,
		)
	}

	var scenarios []*TestScenario
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(
			entry.Name(),
			".yaml",
		) &&
			!strings.HasSuffix(
				entry.Name(),
				".yml",
			) {
			continue
		}

		scenarioPath := filepath.Join(
			dirPath,
			entry.Name(),
		)
		scenario, err := LoadScenario(
			scenarioPath,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"load scenario %s: %w",
				entry.Name(),
				err,
			)
		}
		scenarios = append(scenarios, scenario)
	}

	return scenarios, nil
}
