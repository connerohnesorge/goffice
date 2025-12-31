package gobridge

import (
	"fmt"

	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// Bridge translates test scenarios into goffice API calls
type Bridge struct {
	scenario *framework.TestScenario
}

// NewBridge creates a new Go bridge
func NewBridge(
	scenario *framework.TestScenario,
) *Bridge {
	return &Bridge{scenario: scenario}
}

// Execute runs the scenario and produces a document at outputPath
func (b *Bridge) Execute(
	outputPath string,
) error {
	switch b.scenario.DocumentType {
	case "wordprocessing":
		return b.executeWordprocessing(outputPath)
	case "spreadsheet":
		return b.executeSpreadsheet(outputPath)
	case "presentation":
		return b.executePresentation(outputPath)
	default:
		return fmt.Errorf(
			"unknown document type: %s",
			b.scenario.DocumentType,
		)
	}
}
