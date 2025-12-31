package framework

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// mockGoBridge is a mock implementation of GoBridge for testing
type mockGoBridge struct {
	scenario   *TestScenario
	shouldFail bool
}

func (m *mockGoBridge) Execute(
	outputPath string,
) error {
	if m.shouldFail {
		return fmt.Errorf("mock bridge error")
	}

	// Create a dummy file
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(
		outputPath,
		[]byte("mock document"),
		0644,
	)
}

func mockGoBridgeFactory(
	scenario *TestScenario,
) GoBridge {
	return &mockGoBridge{
		scenario:   scenario,
		shouldFail: false,
	}
}

func mockGoBridgeFactoryFailure(
	scenario *TestScenario,
) GoBridge {
	return &mockGoBridge{
		scenario:   scenario,
		shouldFail: true,
	}
}

// mockComparer is a mock implementation of DocumentComparer for testing
type mockComparer struct {
	shouldFail bool
	xmlMatch   bool
}

func (m *mockComparer) CompareDocuments(
	goPath, dotnetPath string,
	tolerance ToleranceConfig,
) (*ComparisonResult, error) {
	if m.shouldFail {
		return nil, fmt.Errorf("comparison error")
	}

	return &ComparisonResult{
		XMLMatch:    m.xmlMatch,
		BinaryMatch: true,
		VisualMatch: true,
	}, nil
}

func newMockComparerPass() *mockComparer {
	return &mockComparer{
		shouldFail: false,
		xmlMatch:   true,
	}
}

func newMockComparerFail() *mockComparer {
	return &mockComparer{
		shouldFail: false,
		xmlMatch:   false,
	}
}

func TestNewExecutor(t *testing.T) {
	t.Run("CreateExecutor", func(t *testing.T) {
		// Arrange
		csharpBridge := "/path/to/bridge"
		outputDir := "/tmp/output"
		baselineDir := "/tmp/baselines"

		// Act
		executor := NewExecutor(
			csharpBridge,
			outputDir,
			baselineDir,
		)

		// Assert
		if executor == nil {
			t.Fatal(
				"expected executor to be non-nil",
			)
		}
		if executor.csharpBridge != csharpBridge {
			t.Errorf(
				"got csharpBridge %q, want %q",
				executor.csharpBridge,
				csharpBridge,
			)
		}
		if executor.outputDir != outputDir {
			t.Errorf(
				"got outputDir %q, want %q",
				executor.outputDir,
				outputDir,
			)
		}
		if executor.baselineDir != baselineDir {
			t.Errorf(
				"got baselineDir %q, want %q",
				executor.baselineDir,
				baselineDir,
			)
		}
		if executor.baselineMode {
			t.Error(
				"expected baselineMode to be false by default",
			)
		}
	})
}

func TestExecuteGoBridge(t *testing.T) {
	t.Run(
		"GeneratesDocument",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			outputPath := filepath.Join(
				tmpDir,
				"test.docx",
			)

			scenario := &TestScenario{
				Name:         "test-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				"",
			)
			executor.SetGoBridgeFactory(
				mockGoBridgeFactory,
			)
			ctx := context.Background()

			// Act
			err := executor.executeGoBridge(
				ctx,
				scenario,
				outputPath,
			)
			// Assert
			if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			// Verify file was created
			if _, err := os.Stat(outputPath); os.IsNotExist(
				err,
			) {
				t.Errorf(
					"expected output file to be created at %s",
					outputPath,
				)
			}
		},
	)

	t.Run(
		"CreatesOutputDirectory",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			outputPath := filepath.Join(
				tmpDir,
				"nested",
				"dir",
				"test.docx",
			)

			scenario := &TestScenario{
				Name:         "test-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				"",
			)
			executor.SetGoBridgeFactory(
				mockGoBridgeFactory,
			)
			ctx := context.Background()

			// Act
			err := executor.executeGoBridge(
				ctx,
				scenario,
				outputPath,
			)
			// Assert
			if err != nil {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			// Verify directory was created
			if _, err := os.Stat(filepath.Dir(outputPath)); os.IsNotExist(
				err,
			) {
				t.Errorf(
					"expected directory to be created at %s",
					filepath.Dir(outputPath),
				)
			}
		},
	)

	t.Run(
		"ErrorsWhenFactoryNotSet",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			outputPath := filepath.Join(
				tmpDir,
				"test.docx",
			)

			scenario := &TestScenario{
				Name:         "test-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				"",
			)
			// Don't set factory
			ctx := context.Background()

			// Act
			err := executor.executeGoBridge(
				ctx,
				scenario,
				outputPath,
			)

			// Assert
			if err == nil {
				t.Fatal(
					"expected error when factory not set",
				)
			}
		},
	)
}

func TestTestStatus(t *testing.T) {
	t.Run("StatusConstants", func(t *testing.T) {
		if StatusPassed != "PASSED" {
			t.Errorf(
				"got StatusPassed %q, want %q",
				StatusPassed,
				"PASSED",
			)
		}
		if StatusFailed != "FAILED" {
			t.Errorf(
				"got StatusFailed %q, want %q",
				StatusFailed,
				"FAILED",
			)
		}
		if StatusSkipped != "SKIPPED" {
			t.Errorf(
				"got StatusSkipped %q, want %q",
				StatusSkipped,
				"SKIPPED",
			)
		}
	})
}

func TestTestResult(t *testing.T) {
	t.Run("CreateTestResult", func(t *testing.T) {
		// Arrange
		startTime := time.Now()
		endTime := startTime.Add(5 * time.Second)

		// Act
		result := &TestResult{
			Scenario:         "test-scenario",
			Status:           StatusPassed,
			StartTime:        startTime,
			EndTime:          endTime,
			GoOutputPath:     "/tmp/go/test.docx",
			DotNetOutputPath: "/tmp/dotnet/test.docx",
		}

		// Assert
		if result.Scenario != "test-scenario" {
			t.Errorf(
				"got Scenario %q, want %q",
				result.Scenario,
				"test-scenario",
			)
		}
		if result.Status != StatusPassed {
			t.Errorf(
				"got Status %q, want %q",
				result.Status,
				StatusPassed,
			)
		}
		if result.GoOutputPath != "/tmp/go/test.docx" {
			t.Errorf(
				"got GoOutputPath %q, want %q",
				result.GoOutputPath,
				"/tmp/go/test.docx",
			)
		}
	})
}

func TestComparisonResult(t *testing.T) {
	t.Run(
		"CreateComparisonResult",
		func(t *testing.T) {
			// Arrange & Act
			result := &ComparisonResult{
				XMLMatch:    true,
				BinaryMatch: true,
				VisualMatch: true,
			}

			// Assert
			if !result.XMLMatch {
				t.Error(
					"expected XMLMatch to be true",
				)
			}
			if !result.BinaryMatch {
				t.Error(
					"expected BinaryMatch to be true",
				)
			}
			if !result.VisualMatch {
				t.Error(
					"expected VisualMatch to be true",
				)
			}
		},
	)
}

func TestExecuteScenario_GoBridgeError(
	t *testing.T,
) {
	t.Run(
		"HandlesBridgeError",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()

			scenario := &TestScenario{
				Name:         "error-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				tmpDir,
			)
			executor.SetGoBridgeFactory(
				mockGoBridgeFactoryFailure,
			)
			ctx := context.Background()

			// Act
			result, err := executor.ExecuteScenario(
				ctx,
				scenario,
			)
			// Assert
			if err != nil {
				t.Fatalf(
					"unexpected error from ExecuteScenario: %v",
					err,
				)
			}
			if result.Status != StatusFailed {
				t.Errorf(
					"got Status %q, want %q",
					result.Status,
					StatusFailed,
				)
			}
			if result.GoError == "" {
				t.Error(
					"expected GoError to be set",
				)
			}
		},
	)
}

func TestExecuteScenario_BaselineMode(
	t *testing.T,
) {
	t.Run("MissingBaseline", func(t *testing.T) {
		// Arrange
		tmpDir := t.TempDir()

		scenario := &TestScenario{
			Name:         "test-scenario",
			DocumentType: "wordprocessing",
			Operations:   []Operation{},
		}

		executor := NewExecutor(
			"",
			tmpDir,
			tmpDir,
		)
		executor.SetGoBridgeFactory(
			mockGoBridgeFactory,
		)
		executor.SetComparer(
			newMockComparerPass(),
		)
		executor.baselineMode = true
		ctx := context.Background()

		// Act
		result, err := executor.ExecuteScenario(
			ctx,
			scenario,
		)
		// Assert
		if err != nil {
			t.Fatalf(
				"unexpected error from ExecuteScenario: %v",
				err,
			)
		}
		if result.Status != StatusFailed {
			t.Errorf(
				"got Status %q, want %q",
				result.Status,
				StatusFailed,
			)
		}
		if result.ComparisonError == "" {
			t.Error(
				"expected ComparisonError to be set for missing baseline",
			)
		}
	})
}

func TestExecuteScenario_Success(t *testing.T) {
	t.Run(
		"PassingComparison",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()

			// Create baseline file
			baselineDir := filepath.Join(
				tmpDir,
				"baselines",
				"wordprocessing",
			)
			if err := os.MkdirAll(baselineDir, 0755); err != nil {
				t.Fatalf(
					"failed to create baseline dir: %v",
					err,
				)
			}
			baselinePath := filepath.Join(
				baselineDir,
				"test-scenario.docx",
			)
			if err := os.WriteFile(baselinePath, []byte("baseline"), 0644); err != nil {
				t.Fatalf(
					"failed to create baseline: %v",
					err,
				)
			}

			scenario := &TestScenario{
				Name:         "test-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				filepath.Join(
					tmpDir,
					"baselines",
				),
			)
			executor.SetGoBridgeFactory(
				mockGoBridgeFactory,
			)
			executor.SetComparer(
				newMockComparerPass(),
			)
			executor.baselineMode = true
			ctx := context.Background()

			// Act
			result, err := executor.ExecuteScenario(
				ctx,
				scenario,
			)
			// Assert
			if err != nil {
				t.Fatalf(
					"unexpected error from ExecuteScenario: %v",
					err,
				)
			}
			if result.Status != StatusPassed {
				t.Errorf(
					"got Status %q, want %q",
					result.Status,
					StatusPassed,
				)
			}
			if result.GoError != "" {
				t.Errorf(
					"unexpected GoError: %s",
					result.GoError,
				)
			}
			if result.ComparisonError != "" {
				t.Errorf(
					"unexpected ComparisonError: %s",
					result.ComparisonError,
				)
			}
		},
	)

	t.Run(
		"FailingComparison",
		func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()

			// Create baseline file
			baselineDir := filepath.Join(
				tmpDir,
				"baselines",
				"wordprocessing",
			)
			if err := os.MkdirAll(baselineDir, 0755); err != nil {
				t.Fatalf(
					"failed to create baseline dir: %v",
					err,
				)
			}
			baselinePath := filepath.Join(
				baselineDir,
				"test-scenario.docx",
			)
			if err := os.WriteFile(baselinePath, []byte("baseline"), 0644); err != nil {
				t.Fatalf(
					"failed to create baseline: %v",
					err,
				)
			}

			scenario := &TestScenario{
				Name:         "test-scenario",
				DocumentType: "wordprocessing",
				Operations:   []Operation{},
			}

			executor := NewExecutor(
				"",
				tmpDir,
				filepath.Join(
					tmpDir,
					"baselines",
				),
			)
			executor.SetGoBridgeFactory(
				mockGoBridgeFactory,
			)
			executor.SetComparer(
				newMockComparerFail(),
			)
			executor.baselineMode = true
			ctx := context.Background()

			// Act
			result, err := executor.ExecuteScenario(
				ctx,
				scenario,
			)
			// Assert
			if err != nil {
				t.Fatalf(
					"unexpected error from ExecuteScenario: %v",
					err,
				)
			}
			if result.Status != StatusFailed {
				t.Errorf(
					"got Status %q, want %q",
					result.Status,
					StatusFailed,
				)
			}
		},
	)
}
