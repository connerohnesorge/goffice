package e2e

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

func TestE2E(t *testing.T) {
	// Skip if not in E2E mode
	if testing.Short() {
		t.Skip("skipping E2E tests in short mode")
	}

	// Discover test cases
	testCases := discoverTestCases(t)
	if len(testCases) == 0 {
		t.Skip("no test cases found")
	}

	t.Logf(
		"Discovered %d test case(s)",
		len(testCases),
	)

	// Create runner
	config := framework.RunConfig{
		OutputDir: filepath.Join(
			"output",
			"e2e",
		),
		Parallel:      0, // Sequential for now
		RetryCount:    1,
		Timeout:       2 * time.Minute,
		GenerateDiffs: true,
		KeepArtifacts: true,
		GoGeneratorPath: getGoGeneratorPath(
			t,
		),
		CSharpGeneratorPath: getCSharpGeneratorPath(
			t,
		),
	}

	runner, err := framework.NewRunner(config)
	if err != nil {
		t.Fatalf("create runner: %v", err)
	}

	// Run tests
	result, err := runner.Run(
		context.Background(),
		testCases,
	)
	if err != nil {
		t.Fatalf("run tests: %v", err)
	}

	// Log summary
	t.Logf(
		"Test Summary: %d passed, %d failed, %d skipped (total: %d)",
		result.Passed,
		result.Failed,
		result.Skipped,
		len(result.Tests),
	)
	t.Logf(
		"Duration: %v",
		result.EndTime.Sub(result.StartTime),
	)
	t.Logf(
		"Report: %s",
		filepath.Join(
			config.OutputDir,
			"report.html",
		),
	)

	// Assert outcomes
	for _, tr := range result.Tests {
		t.Run(
			tr.TestCase.Name,
			func(t *testing.T) {
				switch tr.Status {
				case framework.StatusFailed:
					t.Errorf(
						"test failed: %v",
						tr.Error,
					)

					// Log slide-level details
					for _, sr := range tr.SlideResults {
						if sr.DiffResult != nil &&
							!sr.DiffResult.Passed {
							t.Logf(
								"  Slide %d: %.2f%% difference (threshold: %.2f%%)",
								sr.SlideIndex+1,
								sr.DiffResult.DiffPercent,
								sr.DiffResult.Threshold*100,
							)
							t.Logf(
								"    Go PNG:     %s",
								sr.GoPNGPath,
							)
							t.Logf(
								"    C# PNG:     %s",
								sr.CSharpPNGPath,
							)
							t.Logf(
								"    Diff PNG:   %s",
								sr.DiffPath,
							)
						}
					}
				case framework.StatusPassed:
					t.Logf(
						"test passed in %v (Go: %v, C#: %v, Render: %v, Compare: %v)",
						tr.GoGenTime+tr.CSharpGenTime+tr.RenderTime+tr.CompareTime,
						tr.GoGenTime,
						tr.CSharpGenTime,
						tr.RenderTime,
						tr.CompareTime,
					)
				}
			},
		)
	}

	// Fail if any tests failed
	if result.Failed > 0 {
		t.Errorf(
			"%d test(s) failed",
			result.Failed,
		)
	}
}

// discoverTestCases loads test cases from testcases/ directory
func discoverTestCases(
	t *testing.T,
) []*framework.TestCase {
	t.Helper()

	testCasesDir := "testcases"

	// Check if testcases directory exists
	if _, err := os.Stat(testCasesDir); os.IsNotExist(
		err,
	) {
		t.Logf(
			"testcases directory not found: %s",
			testCasesDir,
		)

		return nil
	}

	// Find all JSON files in testcases directory
	var testCases []*framework.TestCase

	err := filepath.Walk(
		testCasesDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			// Only process JSON files
			if filepath.Ext(path) != ".json" {
				return nil
			}

			// Load test case
			data, err := os.ReadFile(path)
			if err != nil {
				t.Logf(
					"Failed to read test case %s: %v",
					path,
					err,
				)

				return nil // Continue with other files
			}

			tc, err := framework.FromJSON(data)
			if err != nil {
				t.Logf(
					"Failed to parse test case %s: %v",
					path,
					err,
				)

				return nil // Continue with other files
			}

			testCases = append(testCases, tc)
			t.Logf(
				"Loaded test case: %s (%s) from %s",
				tc.Name,
				tc.ID,
				path,
			)

			return nil
		},
	)
	if err != nil {
		t.Fatalf("discover test cases: %v", err)
	}

	return testCases
}

// getGoGeneratorPath returns the path to the Go generator binary
func getGoGeneratorPath(t *testing.T) string {
	t.Helper()

	// Check environment variable first
	if path := os.Getenv("GO_GENERATOR_PATH"); path != "" {
		return path
	}

	// Try to build the generator
	generatorDir := filepath.Join(
		"generators",
		"go",
	)
	if _, err := os.Stat(generatorDir); os.IsNotExist(
		err,
	) {
		t.Logf(
			"Go generator directory not found: %s",
			generatorDir,
		)

		return ""
	}

	// Build the generator binary
	binaryPath := filepath.Join(
		generatorDir,
		"e2e-go-generator",
	)
	t.Logf(
		"Building Go generator: %s",
		binaryPath,
	)

	// Use go build to compile the generator
	// Note: This will be done externally, for now just return expected path
	return binaryPath
}

// getCSharpGeneratorPath returns the path to the C# generator binary
func getCSharpGeneratorPath(t *testing.T) string {
	t.Helper()

	// Check environment variable first
	if path := os.Getenv("CSHARP_GENERATOR_PATH"); path != "" {
		return path
	}

	// Try to find the C# generator binary
	generatorDir := filepath.Join(
		"generators",
		"csharp",
	)
	if _, err := os.Stat(generatorDir); os.IsNotExist(
		err,
	) {
		t.Logf(
			"C# generator directory not found: %s",
			generatorDir,
		)

		return ""
	}

	// Look for compiled binary
	binaryPath := filepath.Join(
		generatorDir,
		"bin",
		"Debug",
		"net8.0",
		"e2e-csharp-generator",
	)
	if _, err := os.Stat(binaryPath); os.IsNotExist(
		err,
	) {
		t.Logf(
			"C# generator binary not found: %s",
			binaryPath,
		)

		return ""
	}

	return binaryPath
}
