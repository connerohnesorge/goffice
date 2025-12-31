package framework

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// GoBridge defines the interface for executing Go bridge
type GoBridge interface {
	// Execute generates a document at the specified output path
	Execute(outputPath string) error
}

// DocumentComparer defines the interface for comparing documents
type DocumentComparer interface {
	// CompareDocuments compares two documents and returns comparison results
	CompareDocuments(
		goPath, dotnetPath string,
		tolerance ToleranceConfig,
	) (*ComparisonResult, error)
}

// Executor runs test scenarios through both bridges and compares results
type Executor struct {
	csharpBridge    string // Path to C# bridge executable
	outputDir       string // Base output directory
	baselineDir     string // Baseline documents directory
	baselineMode    bool   // If true, compare against baselines instead of running C# bridge
	goBridgeFactory func(*TestScenario) GoBridge
	comparer        DocumentComparer
}

// NewExecutor creates a new test executor
func NewExecutor(
	csharpBridgePath, outputDir, baselineDir string,
) *Executor {
	return &Executor{
		csharpBridge:    csharpBridgePath,
		outputDir:       outputDir,
		baselineDir:     baselineDir,
		baselineMode:    false,
		goBridgeFactory: nil, // Will be set by caller to avoid import cycle
		comparer:        nil, // Will be set by caller to avoid import cycle
	}
}

// SetGoBridgeFactory sets the factory function for creating Go bridges
func (e *Executor) SetGoBridgeFactory(
	factory func(*TestScenario) GoBridge,
) {
	e.goBridgeFactory = factory
}

// SetComparer sets the document comparer
func (e *Executor) SetComparer(
	comparer DocumentComparer,
) {
	e.comparer = comparer
}

// ExecuteScenario runs a scenario through both bridges and compares
func (e *Executor) ExecuteScenario(
	ctx context.Context,
	scenario *TestScenario,
) (*TestResult, error) {
	result := &TestResult{
		Scenario:  scenario.Name,
		StartTime: time.Now(),
	}

	// 1. Generate document with Go
	goOutputPath := filepath.Join(
		e.outputDir,
		"go",
		scenario.Name+".docx",
	)
	if err := e.executeGoBridge(ctx, scenario, goOutputPath); err != nil {
		result.GoError = err.Error()
		result.Status = StatusFailed
		result.EndTime = time.Now()

		return result, nil // Don't fail whole test, just record error
	}
	result.GoOutputPath = goOutputPath

	// 2. Determine comparison target
	var comparisonTargetPath string
	if e.baselineMode {
		// Baseline mode: compare against pre-generated baseline
		comparisonTargetPath = filepath.Join(
			e.baselineDir,
			scenario.DocumentType,
			scenario.Name+".docx",
		)
		if _, err := os.Stat(comparisonTargetPath); os.IsNotExist(
			err,
		) {
			result.ComparisonError = fmt.Sprintf(
				"baseline not found: %s",
				comparisonTargetPath,
			)
			result.Status = StatusFailed
			result.EndTime = time.Now()

			return result, nil
		}
		result.DotNetOutputPath = comparisonTargetPath
	} else {
		// Normal mode: generate document with C# bridge
		dotnetOutputPath := filepath.Join(e.outputDir, "dotnet", scenario.Name+".docx")
		if err := e.executeCSharpBridge(ctx, scenario, dotnetOutputPath); err != nil {
			result.DotNetError = err.Error()
			result.Status = StatusFailed
			result.EndTime = time.Now()

			return result, nil
		}
		result.DotNetOutputPath = dotnetOutputPath
		comparisonTargetPath = dotnetOutputPath
	}

	// 3. Compare documents (three levels)
	comparisonResult, err := e.compareDocuments(
		goOutputPath,
		comparisonTargetPath,
		scenario.Tolerance,
	)
	if err != nil {
		result.ComparisonError = err.Error()
		result.Status = StatusFailed
		result.EndTime = time.Now()

		return result, nil
	}
	result.Comparison = comparisonResult

	// 4. Determine pass/fail
	if comparisonResult.XMLMatch {
		result.Status = StatusPassed
	} else {
		result.Status = StatusFailed
	}

	result.EndTime = time.Now()

	return result, nil
}

// executeGoBridge generates document using goffice
func (e *Executor) executeGoBridge(
	_ context.Context,
	scenario *TestScenario,
	outputPath string,
) error {
	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	if e.goBridgeFactory == nil {
		return fmt.Errorf(
			"go bridge factory not set",
		)
	}

	// Create and execute Go bridge
	bridge := e.goBridgeFactory(scenario)
	if err := bridge.Execute(outputPath); err != nil {
		return fmt.Errorf(
			"go bridge execution: %w",
			err,
		)
	}

	return nil
}

// executeCSharpBridge generates document using Open-XML-SDK
func (e *Executor) executeCSharpBridge(
	ctx context.Context,
	scenario *TestScenario,
	outputPath string,
) error {
	// Add timeout context (default 30 seconds for C# bridge execution)
	ctx, cancel := context.WithTimeout(
		ctx,
		30*time.Second,
	)
	defer cancel()

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	// Write scenario to temp file for C# to read
	scenarioTempPath := filepath.Join(
		os.TempDir(),
		fmt.Sprintf(
			"scenario_%s.yaml",
			scenario.Name,
		),
	)
	scenarioData, err := yaml.Marshal(scenario)
	if err != nil {
		return fmt.Errorf(
			"marshal scenario: %w",
			err,
		)
	}
	if err := os.WriteFile(scenarioTempPath, scenarioData, 0o644); err != nil {
		return fmt.Errorf(
			"write scenario temp file: %w",
			err,
		)
	}
	defer func() {
		_ = os.Remove(scenarioTempPath)
	}()

	// Execute C# bridge with timeout
	cmd := exec.CommandContext(
		ctx,
		e.csharpBridge,
		"--scenario",
		scenarioTempPath,
		"--output",
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Check if timeout occurred
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf(
				"csharp bridge execution timeout exceeded (30s)",
			)
		}

		return fmt.Errorf(
			"csharp bridge execution: %w (output: %s)",
			err,
			string(output),
		)
	}

	return nil
}

// compareDocuments performs three-level comparison
func (e *Executor) compareDocuments(
	goPath, dotnetPath string,
	tolerance ToleranceConfig,
) (*ComparisonResult, error) {
	if e.comparer == nil {
		return nil, fmt.Errorf(
			"document comparer not set",
		)
	}

	return e.comparer.CompareDocuments(
		goPath,
		dotnetPath,
		tolerance,
	)
}

// TestResult represents the result of running a test scenario
type TestResult struct {
	Scenario         string
	Status           TestStatus
	StartTime        time.Time
	EndTime          time.Time
	GoOutputPath     string
	DotNetOutputPath string
	GoError          string
	DotNetError      string
	ComparisonError  string
	Comparison       *ComparisonResult
}

// TestStatus represents test execution status
type TestStatus string

const (
	StatusPassed  TestStatus = "PASSED"
	StatusFailed  TestStatus = "FAILED"
	StatusSkipped TestStatus = "SKIPPED"
)

// XMLDiffResult contains XML comparison results (forward declaration to avoid import cycle)
type XMLDiffResult interface{}

// ComparisonResult contains results from three-level comparison
type ComparisonResult struct {
	XMLMatch    bool
	BinaryMatch bool
	VisualMatch bool

	XMLDiff XMLDiffResult
	// BinaryDiff and VisualDiff are stubs for now
}
