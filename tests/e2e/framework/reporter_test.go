package framework

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestReporter_Generate(t *testing.T) {
	// Create temporary output directory
	tmpDir := t.TempDir()

	// Create test reporter
	reporter := NewReporter(tmpDir)

	// Create sample test results
	results := []*TestResult{
		{
			Scenario:         "test_scenario_1",
			Status:           StatusPassed,
			StartTime:        time.Now().Add(-5 * time.Second),
			EndTime:          time.Now(),
			GoOutputPath:     "/tmp/go/test1.docx",
			DotNetOutputPath: "/tmp/dotnet/test1.docx",
			Comparison: &ComparisonResult{
				XMLMatch:    true,
				BinaryMatch: false,
				VisualMatch: true,
			},
		},
		{
			Scenario:         "test_scenario_2",
			Status:           StatusFailed,
			StartTime:        time.Now().Add(-3 * time.Second),
			EndTime:          time.Now(),
			GoOutputPath:     "/tmp/go/test2.docx",
			DotNetOutputPath: "/tmp/dotnet/test2.docx",
			ComparisonError:  "XML structure mismatch",
			Comparison: &ComparisonResult{
				XMLMatch:    false,
				BinaryMatch: false,
				VisualMatch: false,
			},
		},
		{
			Scenario:  "test_scenario_3",
			Status:    StatusSkipped,
			StartTime: time.Now().Add(-1 * time.Second),
			EndTime:   time.Now(),
		},
	}

	// Generate reports
	err := reporter.Generate(results)
	if err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	// Verify HTML report exists
	htmlPath := filepath.Join(tmpDir, "index.html")
	if _, err := os.Stat(htmlPath); os.IsNotExist(err) {
		t.Errorf("HTML report not created at %s", htmlPath)
	}

	// Verify JSON report exists
	jsonPath := filepath.Join(tmpDir, "results.json")
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		t.Errorf("JSON report not created at %s", jsonPath)
	}

	// Verify HTML content contains expected elements
	htmlContent, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Failed to read HTML report: %v", err)
	}

	htmlStr := string(htmlContent)

	// Check for title
	if !containsString(htmlStr, "E2E Visual Testing Report") {
		t.Error("HTML report missing title")
	}

	// Check for test scenarios
	if !containsString(htmlStr, "test_scenario_1") {
		t.Error("HTML report missing test_scenario_1")
	}
	if !containsString(htmlStr, "test_scenario_2") {
		t.Error("HTML report missing test_scenario_2")
	}
	if !containsString(htmlStr, "test_scenario_3") {
		t.Error("HTML report missing test_scenario_3")
	}

	// Check for status indicators
	if !containsString(htmlStr, "PASSED") {
		t.Error("HTML report missing PASSED status")
	}
	if !containsString(htmlStr, "FAILED") {
		t.Error("HTML report missing FAILED status")
	}
	if !containsString(htmlStr, "SKIPPED") {
		t.Error("HTML report missing SKIPPED status")
	}

	// Check for JavaScript functionality
	if !containsString(htmlStr, "searchInput") {
		t.Error("HTML report missing search functionality")
	}
	if !containsString(htmlStr, "statusFilter") {
		t.Error("HTML report missing filter functionality")
	}
}

func TestCalculateSummary(t *testing.T) {
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
			EndTime:   time.Now().Add(3 * time.Second),
		},
		{
			Status:    StatusSkipped,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Second),
		},
	}

	summary := calculateSummary(results)

	if summary.Total != 4 {
		t.Errorf("Expected Total=4, got %d", summary.Total)
	}
	if summary.Passed != 2 {
		t.Errorf("Expected Passed=2, got %d", summary.Passed)
	}
	if summary.Failed != 1 {
		t.Errorf("Expected Failed=1, got %d", summary.Failed)
	}
	if summary.Skipped != 1 {
		t.Errorf("Expected Skipped=1, got %d", summary.Skipped)
	}

	expectedPassRate := 50.0
	if summary.PassRate != expectedPassRate {
		t.Errorf("Expected PassRate=%.1f%%, got %.1f%%", expectedPassRate, summary.PassRate)
	}
}

func TestGenerateHTMLReport_BackwardCompatibility(t *testing.T) {
	tmpDir := t.TempDir()
	reportPath := filepath.Join(tmpDir, "report.html")

	results := []*TestResult{
		{
			Scenario:  "test",
			Status:    StatusPassed,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(1 * time.Second),
		},
	}

	err := GenerateHTMLReport(results, reportPath)
	if err != nil {
		t.Fatalf("GenerateHTMLReport() failed: %v", err)
	}

	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Errorf("Report not created at %s", reportPath)
	}
}

func TestJSONReport_Structure(t *testing.T) {
	tmpDir := t.TempDir()
	reporter := NewReporter(tmpDir)

	results := []*TestResult{
		{
			Scenario:     "json_test",
			Status:       StatusPassed,
			StartTime:    time.Now().Add(-2 * time.Second),
			EndTime:      time.Now(),
			GoOutputPath: "/tmp/test.docx",
			Comparison: &ComparisonResult{
				XMLMatch:    true,
				BinaryMatch: true,
				VisualMatch: true,
			},
		},
	}

	err := reporter.Generate(results)
	if err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	jsonPath := filepath.Join(tmpDir, "results.json")
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("Failed to read JSON report: %v", err)
	}

	jsonStr := string(jsonContent)

	// Verify JSON structure
	if !containsString(jsonStr, `"version"`) {
		t.Error("JSON report missing version field")
	}
	if !containsString(jsonStr, `"generated_at"`) {
		t.Error("JSON report missing generated_at field")
	}
	if !containsString(jsonStr, `"summary"`) {
		t.Error("JSON report missing summary field")
	}
	if !containsString(jsonStr, `"results"`) {
		t.Error("JSON report missing results field")
	}
	if !containsString(jsonStr, `"json_test"`) {
		t.Error("JSON report missing test scenario")
	}
	if !containsString(jsonStr, `"comparison"`) {
		t.Error("JSON report missing comparison data")
	}
}

// Helper function to check if string contains substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		func() bool {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}())
}
