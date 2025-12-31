package framework

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Generator defines the interface for PPTX generation (Go or C#)
type Generator interface {
	// Generate creates a PPTX file from a test case
	Generate(
		ctx context.Context,
		tc *TestCase,
		outputPath string,
	) error
}

// GoGenerator generates PPTX using goffice (via external command)
type GoGenerator struct {
	BinaryPath string // Path to e2e-go-generator binary
}

// Generate creates PPTX using Go generator
func (g *GoGenerator) Generate(
	ctx context.Context,
	tc *TestCase,
	outputPath string,
) error {
	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	// Create temporary test case JSON file
	tempDir := os.TempDir()
	tcPath := filepath.Join(
		tempDir,
		fmt.Sprintf(
			"testcase_%s_%d.json",
			tc.ID,
			time.Now().UnixNano(),
		),
	)

	tcData, err := tc.ToJSON()
	if err != nil {
		return fmt.Errorf(
			"serialize test case: %w",
			err,
		)
	}

	if err := os.WriteFile(tcPath, tcData, 0o644); err != nil {
		return fmt.Errorf(
			"write test case file: %w",
			err,
		)
	}
	defer func() {
		_ = os.Remove(tcPath)
	}()

	// Execute Go generator
	cmd := exec.CommandContext(ctx, g.BinaryPath,
		"-test", tcPath,
		"-output", outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"go generator execution: %w\nOutput: %s",
			err,
			string(output),
		)
	}

	// Verify output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		return fmt.Errorf(
			"output file not created: %s",
			outputPath,
		)
	}

	return nil
}

// CSharpGenerator generates PPTX using Open-XML-SDK (via external command)
type CSharpGenerator struct {
	BinaryPath string // Path to C# generator executable
}

// Generate creates PPTX using C# generator
func (c *CSharpGenerator) Generate(
	ctx context.Context,
	tc *TestCase,
	outputPath string,
) error {
	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	// Create temporary test case JSON file
	tempDir := os.TempDir()
	tcPath := filepath.Join(
		tempDir,
		fmt.Sprintf(
			"testcase_%s_%d.json",
			tc.ID,
			time.Now().UnixNano(),
		),
	)

	tcData, err := tc.ToJSON()
	if err != nil {
		return fmt.Errorf(
			"serialize test case: %w",
			err,
		)
	}

	if err := os.WriteFile(tcPath, tcData, 0o644); err != nil {
		return fmt.Errorf(
			"write test case file: %w",
			err,
		)
	}
	defer func() {
		_ = os.Remove(tcPath)
	}()

	// Execute C# generator
	cmd := exec.CommandContext(ctx, c.BinaryPath,
		"-test", tcPath,
		"-output", outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"C# generator execution: %w\nOutput: %s",
			err,
			string(output),
		)
	}

	// Verify output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		return fmt.Errorf(
			"output file not created: %s",
			outputPath,
		)
	}

	return nil
}

// Reporter generates HTML reports from test results
type Reporter struct {
	OutputDir string
}

// Runner orchestrates the full test execution pipeline
type Runner struct {
	GoGenerator     Generator     // Go PPTX generator
	CSharpGenerator Generator     // C# PPTX generator (via exec)
	Renderer        Renderer      // PPTX to PNG
	Differ          Differ        // PNG comparison
	Reporter        *Reporter     // HTML report generation
	OutputDir       string        // Base output directory
	Parallel        int           // Number of parallel tests (0 = sequential)
	RetryCount      int           // Number of retries on failure
	Timeout         time.Duration // Test timeout
}

// RunConfig holds runner configuration
type RunConfig struct {
	OutputDir           string
	TestFilter          string // Glob pattern to filter tests
	Parallel            int
	RetryCount          int
	Timeout             time.Duration
	GenerateDiffs       bool
	KeepArtifacts       bool
	GoGeneratorPath     string // Path to Go generator binary
	CSharpGeneratorPath string // Path to C# generator binary
}

// TestSuiteResult holds results from running all tests
type TestSuiteResult struct {
	Tests     []*TestCaseResult
	StartTime time.Time
	EndTime   time.Time
	Passed    int
	Failed    int
	Skipped   int
}

// TestCaseResult holds result for a single test case
type TestCaseResult struct {
	TestCase      *TestCase
	Status        TestStatus
	GoGenTime     time.Duration
	CSharpGenTime time.Duration
	RenderTime    time.Duration
	CompareTime   time.Duration
	SlideResults  []*SlideResult
	Error         error
	RetryCount    int
}

// SlideResult holds comparison result for a single slide
type SlideResult struct {
	SlideIndex    int
	GoPNGPath     string
	CSharpPNGPath string
	DiffPath      string
	DiffResult    *DiffResult
}

// NewRunner creates a new test runner
func NewRunner(
	config RunConfig,
) (*Runner, error) {
	// Validate configuration
	if config.OutputDir == "" {
		config.OutputDir = "output"
	}
	if config.Timeout == 0 {
		config.Timeout = 2 * time.Minute
	}
	if config.RetryCount < 0 {
		config.RetryCount = 0
	}

	// Create output directory
	if err := os.MkdirAll(config.OutputDir, 0o755); err != nil {
		return nil, fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	// Create Go generator
	goGenerator := &GoGenerator{
		BinaryPath: config.GoGeneratorPath,
	}
	if goGenerator.BinaryPath == "" {
		// Try to find binary in PATH
		path, err := exec.LookPath(
			"e2e-go-generator",
		)
		if err != nil {
			return nil, fmt.Errorf(
				"go generator not found: %w (set GoGeneratorPath in config)",
				err,
			)
		}
		goGenerator.BinaryPath = path
	}

	// Create C# generator
	csharpGenerator := &CSharpGenerator{
		BinaryPath: config.CSharpGeneratorPath,
	}
	if csharpGenerator.BinaryPath == "" {
		// Try to find binary in PATH
		path, err := exec.LookPath(
			"e2e-csharp-generator",
		)
		if err != nil {
			return nil, fmt.Errorf(
				"C# generator not found: %w (set CSharpGeneratorPath in config)",
				err,
			)
		}
		csharpGenerator.BinaryPath = path
	}

	// Create renderer
	renderer := NewLibreOfficeRenderer(
		filepath.Join(config.OutputDir, "render"),
	)

	// Create differ
	differ := NewGoDiffer()

	// Create reporter
	reporter := &Reporter{
		OutputDir: config.OutputDir,
	}

	return &Runner{
		GoGenerator:     goGenerator,
		CSharpGenerator: csharpGenerator,
		Renderer:        renderer,
		Differ:          differ,
		Reporter:        reporter,
		OutputDir:       config.OutputDir,
		Parallel:        config.Parallel,
		RetryCount:      config.RetryCount,
		Timeout:         config.Timeout,
	}, nil
}

// Run executes all test cases matching the filter
func (r *Runner) Run(
	ctx context.Context,
	testCases []*TestCase,
) (*TestSuiteResult, error) {
	result := &TestSuiteResult{
		Tests: make(
			[]*TestCaseResult,
			0,
			len(testCases),
		),
		StartTime: time.Now(),
	}

	// Run tests sequentially or in parallel
	if r.Parallel <= 0 || r.Parallel == 1 {
		// Sequential execution
		for _, tc := range testCases {
			tcResult := r.runTest(ctx, tc)
			result.Tests = append(
				result.Tests,
				tcResult,
			)

			// Update counters
			switch tcResult.Status {
			case StatusPassed:
				result.Passed++
			case StatusFailed:
				result.Failed++
			case StatusSkipped:
				result.Skipped++
			}
		}
	} else {
		// Parallel execution
		resultChan := make(chan *TestCaseResult, len(testCases))
		semaphore := make(chan struct{}, r.Parallel)
		var wg sync.WaitGroup

		for _, tc := range testCases {
			wg.Add(1)
			go func(testCase *TestCase) {
				defer wg.Done()

				// Acquire semaphore
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				// Run test
				tcResult := r.runTest(ctx, testCase)
				resultChan <- tcResult
			}(tc)
		}

		// Wait for all tests to complete
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Collect results
		for tcResult := range resultChan {
			result.Tests = append(result.Tests, tcResult)

			// Update counters
			switch tcResult.Status {
			case StatusPassed:
				result.Passed++
			case StatusFailed:
				result.Failed++
			case StatusSkipped:
				result.Skipped++
			}
		}
	}

	result.EndTime = time.Now()

	// Generate HTML report
	reportPath := filepath.Join(
		r.OutputDir,
		"report.html",
	)
	if err := r.generateReport(result, reportPath); err != nil {
		return result, fmt.Errorf(
			"generate report: %w",
			err,
		)
	}

	return result, nil
}

// runTest executes a single test case with retry logic
func (r *Runner) runTest(
	ctx context.Context,
	tc *TestCase,
) *TestCaseResult {
	result := &TestCaseResult{
		TestCase:     tc,
		Status:       StatusFailed,
		SlideResults: []*SlideResult{},
	}

	// Apply timeout to test
	testCtx, cancel := context.WithTimeout(
		ctx,
		r.Timeout,
	)
	defer cancel()

	// Retry loop
	for attempt := 0; attempt <= r.RetryCount; attempt++ {
		if attempt > 0 {
			result.RetryCount = attempt
			// Exponential backoff: 1s, 2s, 4s
			backoff := time.Duration(
				1<<uint(attempt-1),
			) * time.Second
			select {
			case <-time.After(backoff):
			case <-testCtx.Done():
				result.Error = fmt.Errorf(
					"test timeout exceeded",
				)
				result.Status = StatusFailed

				return result
			}
		}

		// Execute test
		if err := r.executeTest(testCtx, tc, result); err != nil {
			result.Error = err
			result.Status = StatusFailed

			// Check if timeout occurred
			if testCtx.Err() == context.DeadlineExceeded {
				result.Error = fmt.Errorf(
					"test timeout exceeded (%v)",
					r.Timeout,
				)

				break // Don't retry on timeout
			}

			// Continue to next retry attempt
			continue
		}

		// Test passed
		result.Status = StatusPassed
		result.Error = nil

		break
	}

	return result
}

// executeTest performs the actual test execution
func (r *Runner) executeTest(
	ctx context.Context,
	tc *TestCase,
	result *TestCaseResult,
) error {
	// Step 1: Generate PPTX with Go
	goStart := time.Now()
	goPPTXPath := filepath.Join(
		r.OutputDir,
		"go",
		fmt.Sprintf("%s.pptx", tc.ID),
	)
	if err := r.GoGenerator.Generate(ctx, tc, goPPTXPath); err != nil {
		return fmt.Errorf(
			"generate Go PPTX: %w",
			err,
		)
	}
	result.GoGenTime = time.Since(goStart)

	// Step 2: Generate PPTX with C#
	csharpStart := time.Now()
	csharpPPTXPath := filepath.Join(
		r.OutputDir,
		"csharp",
		fmt.Sprintf("%s.pptx", tc.ID),
	)
	if err := r.CSharpGenerator.Generate(ctx, tc, csharpPPTXPath); err != nil {
		return fmt.Errorf(
			"generate C# PPTX: %w",
			err,
		)
	}
	result.CSharpGenTime = time.Since(csharpStart)

	// Step 3: Render PNGs
	renderStart := time.Now()
	goPNGs, csharpPNGs, err := r.renderPPTXFiles(
		ctx,
		goPPTXPath,
		csharpPPTXPath,
	)
	if err != nil {
		return fmt.Errorf("render PNGs: %w", err)
	}
	result.RenderTime = time.Since(renderStart)

	// Step 4: Compare PNGs
	compareStart := time.Now()
	slideResults, err := r.compareSlides(
		ctx,
		goPNGs,
		csharpPNGs,
		tc,
	)
	if err != nil {
		return fmt.Errorf(
			"compare slides: %w",
			err,
		)
	}
	result.CompareTime = time.Since(compareStart)
	result.SlideResults = slideResults

	// Check if all slides passed
	for _, sr := range slideResults {
		if sr.DiffResult != nil &&
			!sr.DiffResult.Passed {
			return fmt.Errorf(
				"slide %d comparison failed: %.2f%% difference (threshold: %.2f%%)",
				sr.SlideIndex,
				sr.DiffResult.DiffPercent,
				sr.DiffResult.Threshold*100,
			)
		}
	}

	return nil
}

// renderPPTXFiles renders both PPTX files to PNG images
func (r *Runner) renderPPTXFiles(
	ctx context.Context,
	goPPTX, csharpPPTX string,
) (goPNGs, csharpPNGs []string, err error) {
	// Render Go PPTX
	goPNGs, err = r.Renderer.Render(ctx, goPPTX)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"render Go PPTX: %w",
			err,
		)
	}

	// Render C# PPTX
	csharpPNGs, err = r.Renderer.Render(
		ctx,
		csharpPPTX,
	)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"render C# PPTX: %w",
			err,
		)
	}

	// Verify same number of slides
	if len(goPNGs) != len(csharpPNGs) {
		return nil, nil, fmt.Errorf(
			"slide count mismatch: Go=%d, C#=%d",
			len(goPNGs),
			len(csharpPNGs),
		)
	}

	return goPNGs, csharpPNGs, nil
}

// compareSlides compares PNG images between Go and C#
func (r *Runner) compareSlides(
	ctx context.Context,
	goPNGs, csharpPNGs []string,
	tc *TestCase,
) ([]*SlideResult, error) {
	results := make([]*SlideResult, len(goPNGs))

	for i := range goPNGs {
		// Create diff output path
		diffPath := filepath.Join(
			r.OutputDir,
			"diff",
			fmt.Sprintf(
				"%s_slide_%d_diff.png",
				tc.ID,
				i+1,
			),
		)
		if err := os.MkdirAll(filepath.Dir(diffPath), 0o755); err != nil {
			return nil, fmt.Errorf(
				"create diff directory: %w",
				err,
			)
		}

		// Compare images
		diffResult, err := r.Differ.Compare(
			ctx,
			goPNGs[i],
			csharpPNGs[i],
			diffPath,
			tc.Config,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"compare slide %d: %w",
				i,
				err,
			)
		}

		results[i] = &SlideResult{
			SlideIndex:    i,
			GoPNGPath:     goPNGs[i],
			CSharpPNGPath: csharpPNGs[i],
			DiffPath:      diffPath,
			DiffResult:    diffResult,
		}
	}

	return results, nil
}

// generateReport creates an HTML report from test results
func (r *Runner) generateReport(
	result *TestSuiteResult,
	reportPath string,
) error {
	// Convert TestSuiteResult to []*TestResult for reporter
	var testResults []*TestResult
	for _, tcr := range result.Tests {
		tr := &TestResult{
			Scenario:  tcr.TestCase.Name,
			Status:    tcr.Status,
			StartTime: result.StartTime,
			EndTime:   result.EndTime,
		}

		if tcr.Error != nil {
			tr.ComparisonError = tcr.Error.Error()
		}

		// Add slide comparison details
		if len(tcr.SlideResults) > 0 {
			var details []string
			for _, sr := range tcr.SlideResults {
				if sr.DiffResult != nil {
					details = append(
						details,
						fmt.Sprintf(
							"Slide %d: %.2f%% diff",
							sr.SlideIndex+1,
							sr.DiffResult.DiffPercent,
						),
					)
				}
			}
			if len(details) > 0 {
				if tr.ComparisonError != "" {
					tr.ComparisonError += "\n"
				}
				tr.ComparisonError += strings.Join(
					details,
					"\n",
				)
			}
		}

		testResults = append(testResults, tr)
	}

	return GenerateHTMLReport(
		testResults,
		reportPath,
	)
}
