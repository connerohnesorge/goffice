package framework

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ExampleReporter demonstrates how to use the Reporter to generate HTML and JSON reports
func ExampleReporter() {
	// Create a temporary directory for the report
	tmpDir := os.TempDir()
	reportDir := filepath.Join(tmpDir, "e2e-report-example")
	defer func() {
		_ = os.RemoveAll(reportDir) // Clean up
	}()

	// Create reporter
	reporter := NewReporter(reportDir)

	// Create sample test results
	results := []*TestResult{
		{
			Scenario:  "basic_text_formatting",
			Status:    StatusPassed,
			StartTime: time.Now().Add(-10 * time.Second),
			EndTime:   time.Now().Add(-5 * time.Second),
			GoOutputPath:     filepath.Join(tmpDir, "go", "test1.docx"),
			DotNetOutputPath: filepath.Join(tmpDir, "dotnet", "test1.docx"),
			Comparison: &ComparisonResult{
				XMLMatch:    true,
				BinaryMatch: true,
				VisualMatch: true,
			},
		},
		{
			Scenario:  "table_with_borders",
			Status:    StatusFailed,
			StartTime: time.Now().Add(-5 * time.Second),
			EndTime:   time.Now(),
			GoOutputPath:     filepath.Join(tmpDir, "go", "test2.docx"),
			DotNetOutputPath: filepath.Join(tmpDir, "dotnet", "test2.docx"),
			ComparisonError: "XML structure mismatch in table borders",
			Comparison: &ComparisonResult{
				XMLMatch:    false,
				BinaryMatch: false,
				VisualMatch: false,
			},
		},
	}

	// Generate reports
	if err := reporter.Generate(results); err != nil {
		fmt.Printf("Error generating reports: %v\n", err)

		return
	}

	// Verify files were created
	htmlPath := filepath.Join(reportDir, "index.html")
	jsonPath := filepath.Join(reportDir, "results.json")

	htmlExists := false
	jsonExists := false

	if _, err := os.Stat(htmlPath); err == nil {
		htmlExists = true
	}
	if _, err := os.Stat(jsonPath); err == nil {
		jsonExists = true
	}

	fmt.Printf("HTML report created: %t\n", htmlExists)
	fmt.Printf("JSON report created: %t\n", jsonExists)
	fmt.Printf("Report directory: %s\n", reportDir)

	// Output:
	// HTML report created: true
	// JSON report created: true
	// Report directory: /tmp/e2e-report-example
}

// ExampleGenerateHTMLReport demonstrates the backward-compatible convenience function
func ExampleGenerateHTMLReport() {
	// Create a temporary directory for the report
	tmpDir := os.TempDir()
	reportPath := filepath.Join(tmpDir, "test-report.html")
	defer func() {
		_ = os.Remove(reportPath) // Clean up
	}()

	// Create sample test results
	results := []*TestResult{
		{
			Scenario:  "simple_test",
			Status:    StatusPassed,
			StartTime: time.Now().Add(-2 * time.Second),
			EndTime:   time.Now(),
		},
	}

	// Generate HTML report using the convenience function
	if err := GenerateHTMLReport(results, reportPath); err != nil {
		fmt.Printf("Error: %v\n", err)

		return
	}

	// Verify the file was created
	if _, err := os.Stat(reportPath); err == nil {
		fmt.Println("Report created successfully")
	}

	// Output:
	// Report created successfully
}

// ExampleReportSummary demonstrates summary statistics calculation
func ExampleReportSummary() {
	results := []*TestResult{
		{
			Status:    StatusPassed,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Second),
		},
		{
			Status:    StatusPassed,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(2 * time.Second),
		},
		{
			Status:    StatusFailed,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Second),
		},
		{
			Status:    StatusSkipped,
			StartTime: time.Now(),
			EndTime:   time.Now(),
		},
	}

	summary := calculateSummary(results)

	fmt.Printf("Total: %d\n", summary.Total)
	fmt.Printf("Passed: %d\n", summary.Passed)
	fmt.Printf("Failed: %d\n", summary.Failed)
	fmt.Printf("Skipped: %d\n", summary.Skipped)
	fmt.Printf("Pass Rate: %.1f%%\n", summary.PassRate)

	// Output:
	// Total: 4
	// Passed: 2
	// Failed: 1
	// Skipped: 1
	// Pass Rate: 50.0%
}
