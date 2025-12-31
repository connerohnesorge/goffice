package framework

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	// Create temporary output directory
	tempDir := t.TempDir()

	// Create config
	config := RunConfig{
		OutputDir:           tempDir,
		Parallel:            0,
		RetryCount:          1,
		Timeout:             2 * time.Minute,
		GoGeneratorPath:     "/fake/go-generator",
		CSharpGeneratorPath: "/fake/csharp-generator",
	}

	// Create runner
	runner, err := NewRunner(config)
	if err != nil {
		t.Fatalf("NewRunner failed: %v", err)
	}

	// Verify runner fields
	if runner.OutputDir != tempDir {
		t.Errorf(
			"OutputDir = %s, want %s",
			runner.OutputDir,
			tempDir,
		)
	}
	if runner.Parallel != 0 {
		t.Errorf(
			"Parallel = %d, want 0",
			runner.Parallel,
		)
	}
	if runner.RetryCount != 1 {
		t.Errorf(
			"RetryCount = %d, want 1",
			runner.RetryCount,
		)
	}
	if runner.Timeout != 2*time.Minute {
		t.Errorf(
			"Timeout = %v, want %v",
			runner.Timeout,
			2*time.Minute,
		)
	}
	if runner.GoGenerator == nil {
		t.Error("GoGenerator is nil")
	}
	if runner.CSharpGenerator == nil {
		t.Error("CSharpGenerator is nil")
	}
	if runner.Renderer == nil {
		t.Error("Renderer is nil")
	}
	if runner.Differ == nil {
		t.Error("Differ is nil")
	}
	if runner.Reporter == nil {
		t.Error("Reporter is nil")
	}
}

func TestGoGenerator_Generate(t *testing.T) {
	// Skip if go generator binary doesn't exist
	t.Skip(
		"skipping generator test (requires binary)",
	)

	tempDir := t.TempDir()

	// Create test case
	tc := NewTestCase(
		"test_01",
		"Test Case 1",
		CategoryChart,
	)
	tc.Spec = TestSpec{
		SlideCount: 1,
		SlideSize:  SlideSize16x9,
		Slides: []SlideSpec{
			{
				Index:    0,
				Layout:   "Blank",
				Elements: []ElementSpec{},
			},
		},
	}

	// Create generator
	gen := &GoGenerator{
		BinaryPath: "/path/to/e2e-go-generator",
	}

	// Generate PPTX
	outputPath := filepath.Join(
		tempDir,
		"test.pptx",
	)
	ctx := context.Background()
	err := gen.Generate(ctx, tc, outputPath)
	if err != nil {
		t.Logf(
			"Generate failed (expected without binary): %v",
			err,
		)
	}
}

func TestCSharpGenerator_Generate(t *testing.T) {
	// Skip if csharp generator binary doesn't exist
	t.Skip(
		"skipping generator test (requires binary)",
	)

	tempDir := t.TempDir()

	// Create test case
	tc := NewTestCase(
		"test_01",
		"Test Case 1",
		CategoryChart,
	)
	tc.Spec = TestSpec{
		SlideCount: 1,
		SlideSize:  SlideSize16x9,
		Slides: []SlideSpec{
			{
				Index:    0,
				Layout:   "Blank",
				Elements: []ElementSpec{},
			},
		},
	}

	// Create generator
	gen := &CSharpGenerator{
		BinaryPath: "/path/to/e2e-csharp-generator",
	}

	// Generate PPTX
	outputPath := filepath.Join(
		tempDir,
		"test.pptx",
	)
	ctx := context.Background()
	err := gen.Generate(ctx, tc, outputPath)
	if err != nil {
		t.Logf(
			"Generate failed (expected without binary): %v",
			err,
		)
	}
}

func TestRunner_Run_NoTestCases(t *testing.T) {
	// Create temporary output directory
	tempDir := t.TempDir()

	// Create config
	config := RunConfig{
		OutputDir:           tempDir,
		Parallel:            0,
		RetryCount:          1,
		Timeout:             2 * time.Minute,
		GoGeneratorPath:     "/fake/go-generator",
		CSharpGeneratorPath: "/fake/csharp-generator",
	}

	// Create runner
	runner, err := NewRunner(config)
	if err != nil {
		t.Fatalf("NewRunner failed: %v", err)
	}

	// Run with no test cases
	result, err := runner.Run(
		context.Background(),
		[]*TestCase{},
	)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	// Verify result
	if len(result.Tests) != 0 {
		t.Errorf(
			"Tests count = %d, want 0",
			len(result.Tests),
		)
	}
	if result.Passed != 0 {
		t.Errorf(
			"Passed = %d, want 0",
			result.Passed,
		)
	}
	if result.Failed != 0 {
		t.Errorf(
			"Failed = %d, want 0",
			result.Failed,
		)
	}
	if result.Skipped != 0 {
		t.Errorf(
			"Skipped = %d, want 0",
			result.Skipped,
		)
	}

	// Verify report was generated
	reportPath := filepath.Join(
		tempDir,
		"report.html",
	)
	if _, err := os.Stat(reportPath); os.IsNotExist(
		err,
	) {
		t.Errorf(
			"Report not generated: %s",
			reportPath,
		)
	}
}

func TestTestSuiteResult_Counters(t *testing.T) {
	result := &TestSuiteResult{
		Tests:     []*TestCaseResult{},
		StartTime: time.Now(),
		EndTime: time.Now().
			Add(1 * time.Minute),
	}

	// Add test results
	result.Tests = append(
		result.Tests,
		&TestCaseResult{
			Status: StatusPassed,
		},
	)
	result.Passed++

	result.Tests = append(
		result.Tests,
		&TestCaseResult{
			Status: StatusFailed,
		},
	)
	result.Failed++

	result.Tests = append(
		result.Tests,
		&TestCaseResult{
			Status: StatusSkipped,
		},
	)
	result.Skipped++

	// Verify counters
	if result.Passed != 1 {
		t.Errorf(
			"Passed = %d, want 1",
			result.Passed,
		)
	}
	if result.Failed != 1 {
		t.Errorf(
			"Failed = %d, want 1",
			result.Failed,
		)
	}
	if result.Skipped != 1 {
		t.Errorf(
			"Skipped = %d, want 1",
			result.Skipped,
		)
	}
	if len(result.Tests) != 3 {
		t.Errorf(
			"Tests count = %d, want 3",
			len(result.Tests),
		)
	}
}

func TestSlideResult_Structure(t *testing.T) {
	sr := &SlideResult{
		SlideIndex:    0,
		GoPNGPath:     "/path/to/go.png",
		CSharpPNGPath: "/path/to/csharp.png",
		DiffPath:      "/path/to/diff.png",
		DiffResult: &DiffResult{
			DiffPercent: 1.5,
			Passed:      true,
		},
	}

	if sr.SlideIndex != 0 {
		t.Errorf(
			"SlideIndex = %d, want 0",
			sr.SlideIndex,
		)
	}
	if sr.DiffResult == nil {
		t.Error("DiffResult is nil")
	} else if sr.DiffResult.DiffPercent != 1.5 {
		t.Errorf("DiffPercent = %.2f, want 1.5", sr.DiffResult.DiffPercent)
	}
}
