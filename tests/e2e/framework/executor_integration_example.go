package framework

// This file provides example usage patterns for the Executor
// It is not meant to be run directly, but serves as documentation

/*
// Example 1: Basic executor setup and usage
func ExampleExecutorBasicUsage() {
	// Import the necessary packages
	import (
		"context"
		"github.com/connerohnesorge/goffice/tests/e2e/bridges/go"
		"github.com/connerohnesorge/goffice/tests/e2e/comparison"
	)

	// Create executor
	executor := NewExecutor(
		"/path/to/csharp/bridge",
		"./output",
		"./baselines",
	)

	// Set Go bridge factory
	executor.SetGoBridgeFactory(func(scenario *TestScenario) GoBridge {
		return gobridge.NewBridge(scenario)
	})

	// Set document comparer
	executor.SetComparer(&comparison.DefaultComparer{})

	// Load scenario
	scenario, err := LoadScenario("scenarios/wordprocessing/basic-paragraph.yaml")
	if err != nil {
		panic(err)
	}

	// Execute scenario
	ctx := context.Background()
	result, err := executor.ExecuteScenario(ctx, scenario)
	if err != nil {
		panic(err)
	}

	// Check results
	if result.Status == StatusPassed {
		println("Test passed!")
	} else {
		println("Test failed:", result.ComparisonError)
	}
}

// Example 2: Baseline mode
func ExampleExecutorBaselineMode() {
	executor := NewExecutor(
		"",  // No C# bridge needed in baseline mode
		"./output",
		"./baselines",
	)

	// Set Go bridge factory
	executor.SetGoBridgeFactory(func(scenario *TestScenario) GoBridge {
		return gobridge.NewBridge(scenario)
	})

	// Set comparer
	executor.SetComparer(&comparison.DefaultComparer{})

	// Enable baseline mode
	executor.baselineMode = true

	// Execute against baseline
	scenario, _ := LoadScenario("scenarios/wordprocessing/basic-paragraph.yaml")
	ctx := context.Background()
	result, _ := executor.ExecuteScenario(ctx, scenario)

	// In baseline mode, comparison is against pre-generated baseline
	if result.Status == StatusPassed {
		println("Matches baseline!")
	}
}

// Example 3: Batch execution of multiple scenarios
func ExampleExecutorBatchExecution() {
	executor := NewExecutor(
		"/path/to/csharp/bridge",
		"./output",
		"./baselines",
	)

	executor.SetGoBridgeFactory(func(scenario *TestScenario) GoBridge {
		return gobridge.NewBridge(scenario)
	})
	executor.SetComparer(&comparison.DefaultComparer{})

	// Load all scenarios from directory
	scenarios, err := LoadScenariosFromDirectory("scenarios/wordprocessing")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	var results []*TestResult

	// Execute each scenario
	for _, scenario := range scenarios {
		result, err := executor.ExecuteScenario(ctx, scenario)
		if err != nil {
			println("Error executing scenario:", scenario.Name, err)
			continue
		}
		results = append(results, result)
	}

	// Print summary
	passed := 0
	for _, result := range results {
		if result.Status == StatusPassed {
			passed++
		}
	}
	println("Passed:", passed, "/", len(results))
}

// Example 4: Custom comparer with tolerance
func ExampleExecutorCustomTolerance() {
	// Create a custom comparer that respects tolerance settings
	type CustomComparer struct{}

	func (c *CustomComparer) CompareDocuments(goPath, dotnetPath string, tolerance ToleranceConfig) (*ComparisonResult, error) {
		// Use tolerance settings for comparison
		if tolerance.XMLAttributeOrderSensitive {
			// Strict comparison
		} else {
			// Relaxed comparison
		}
		return &ComparisonResult{XMLMatch: true}, nil
	}

	executor := NewExecutor(
		"/path/to/csharp/bridge",
		"./output",
		"./baselines",
	)

	executor.SetGoBridgeFactory(func(scenario *TestScenario) GoBridge {
		return gobridge.NewBridge(scenario)
	})
	executor.SetComparer(&CustomComparer{})

	// Scenario with custom tolerance
	scenario := &TestScenario{
		Name: "strict-comparison",
		Tolerance: ToleranceConfig{
			XMLAttributeOrderSensitive: true,
			VisualPixelTolerance:       0.0,
			VisualDiffThreshold:        0.0,
		},
	}

	ctx := context.Background()
	result, _ := executor.ExecuteScenario(ctx, scenario)
	_ = result
}
*/
