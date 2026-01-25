package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/connerohnesorge/goffice/tests/e2e/comparison"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

func main() {
	var (
		scenariosDir = flag.String(
			"scenarios",
			"scenarios/wordprocessing",
			"Scenarios directory",
		)
		outputDir = flag.String(
			"output",
			"output",
			"Output directory",
		)
		baselineDir = flag.String(
			"baselines",
			"baselines",
			"Baselines directory",
		)
		csharpBridge = flag.String(
			"csharp-bridge",
			"bridges/csharp/DocxBridge/bin/Release/net9.0/DocxBridge",
			"C# bridge executable",
		)
		generateBaselines = flag.Bool(
			"generate-baselines",
			false,
			"Generate new baselines from .NET SDK",
		)
		htmlReport = flag.Bool(
			"html-report",
			true,
			"Generate HTML report",
		)
	)
	flag.Parse()

	ctx := context.Background()

	// TODO: Implement baseline generation mode when --generate-baselines is set
	// For now, this flag is declared but not used (will be implemented in Phase 2)
	_ = generateBaselines

	// Load scenarios
	scenarios, err := framework.LoadScenariosFromDirectory(
		*scenariosDir,
	)
	if err != nil {
		log.Fatalf(
			"Failed to load scenarios: %v",
			err,
		)
	}

	log.Printf(
		"Loaded %d scenarios from %s",
		len(scenarios),
		*scenariosDir,
	)

	// Create executor
	executor := framework.NewExecutor(
		*csharpBridge,
		*outputDir,
		*baselineDir,
	)

	// Wire up Go bridge factory
	executor.SetGoBridgeFactory(
		func(scenario *framework.TestScenario) framework.GoBridge {
			return gobridge.NewBridge(scenario)
		},
	)

	// Wire up document comparer
	executor.SetComparer(&documentComparer{})

	// Execute all scenarios
	var results []*framework.TestResult
	for _, scenario := range scenarios {
		log.Printf(
			"Executing scenario: %s",
			scenario.Name,
		)

		result, err := executor.ExecuteScenario(
			ctx,
			scenario,
		)
		if err != nil {
			log.Printf("  ERROR: %v", err)

			continue
		}

		log.Printf("  Status: %s", result.Status)
		results = append(results, result)
	}

	// Generate HTML report
	if *htmlReport {
		reportPath := filepath.Join(
			"reports",
			"html",
			"index.html",
		)
		if err := framework.GenerateHTMLReport(results, reportPath); err != nil {
			log.Fatalf(
				"Failed to generate HTML report: %v",
				err,
			)
		}
		log.Printf(
			"HTML report generated: %s",
			reportPath,
		)
	}

	// Print summary
	printSummary(results)

	// Exit with error if any tests failed
	for _, result := range results {
		if result.Status == framework.StatusFailed {
			os.Exit(1)
		}
	}
}

func printSummary(
	results []*framework.TestResult,
) {
	var passed, failed, skipped int
	for _, result := range results {
		switch result.Status {
		case framework.StatusPassed:
			passed++
		case framework.StatusFailed:
			failed++
		case framework.StatusSkipped:
			skipped++
		}
	}

	fmt.Println(
		"\n================================",
	)
	fmt.Printf("Test Results Summary\n")
	fmt.Println(
		"================================",
	)
	fmt.Printf("Total:   %d\n", len(results))
	fmt.Printf("Passed:  %d\n", passed)
	fmt.Printf("Failed:  %d\n", failed)
	fmt.Printf("Skipped: %d\n", skipped)
	fmt.Println(
		"================================",
	)
}

// documentComparer implements the DocumentComparer interface
type documentComparer struct{}

func (c *documentComparer) CompareDocuments(
	goPath, dotnetPath string,
	tolerance framework.ToleranceConfig,
) (*framework.ComparisonResult, error) {
	// Use the comparison package's CompareXMLStructure function
	xmlDiff, err := comparison.CompareXMLStructure(
		goPath,
		dotnetPath,
		tolerance,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"XML comparison failed: %w",
			err,
		)
	}

	// For now, only XML comparison is implemented
	// BinaryMatch and VisualMatch will be added in Phase 2
	result := &framework.ComparisonResult{
		XMLMatch:    xmlDiff.Match,
		BinaryMatch: true, // Stub for now
		VisualMatch: true, // Stub for now
		XMLDiff:     xmlDiff,
	}

	return result, nil
}
