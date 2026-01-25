package framework

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Reporter generates HTML and JSON reports from E2E test results.
// HTML reports include interactive side-by-side comparisons with diff images.
// JSON reports provide machine-readable test metrics for CI/CD integration.
type Reporter struct {
	OutputDir string // Directory where reports are saved
}

// NewReporter creates a new reporter that saves reports to the specified directory.
// The directory is created if it doesn't exist.
func NewReporter(outputDir string) *Reporter {
	return &Reporter{
		OutputDir: outputDir,
	}
}

// Generate creates both HTML and JSON reports
func (r *Reporter) Generate(
	results []*TestResult,
) error {
	// Generate HTML report
	htmlPath := filepath.Join(
		r.OutputDir,
		"index.html",
	)
	if err := r.generateHTMLReport(results, htmlPath); err != nil {
		return fmt.Errorf(
			"generate HTML report: %w",
			err,
		)
	}

	// Generate JSON report
	jsonPath := filepath.Join(
		r.OutputDir,
		"results.json",
	)
	if err := r.generateJSONReport(results, jsonPath); err != nil {
		return fmt.Errorf(
			"generate JSON report: %w",
			err,
		)
	}

	return nil
}

// generateHTMLReport creates an HTML report with interactive features
func (r *Reporter) generateHTMLReport(
	results []*TestResult,
	reportPath string,
) error {
	// Ensure report directory exists
	if err := os.MkdirAll(filepath.Dir(reportPath), 0o755); err != nil {
		return fmt.Errorf(
			"create report directory: %w",
			err,
		)
	}

	// Open output file
	file, err := os.Create(reportPath)
	if err != nil {
		return fmt.Errorf(
			"create report file: %w",
			err,
		)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil &&
			err == nil {
			err = fmt.Errorf(
				"close report file: %w",
				closeErr,
			)
		}
	}()

	// Calculate summary statistics
	summary := calculateSummary(results)

	// Prepare template data
	data := struct {
		Title     string
		Generated string
		Summary   ReportSummary
		Results   []*TestResult
	}{
		Title:     "E2E Visual Testing Report",
		Generated: time.Now().Format("2006-01-02 15:04:05"),
		Summary:   summary,
		Results:   results,
	}

	// Create template with custom functions
	funcMap := template.FuncMap{
		"lower": func(s TestStatus) string {
			return strings.ToLower(string(s))
		},
		"hasImages": func(r *TestResult) bool {
			// Check if result has image paths (for visual comparison)
			return r.GoOutputPath != "" || r.DotNetOutputPath != ""
		},
	}

	// Execute template
	tmpl := template.Must(
		template.New("report").
			Funcs(funcMap).
			Parse(htmlTemplate),
	)
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf(
			"execute template: %w",
			err,
		)
	}

	return nil
}

// generateJSONReport creates a machine-readable JSON report
func (r *Reporter) generateJSONReport(
	results []*TestResult,
	jsonPath string,
) error {
	// Ensure report directory exists
	if err := os.MkdirAll(filepath.Dir(jsonPath), 0o755); err != nil {
		return fmt.Errorf(
			"create report directory: %w",
			err,
		)
	}

	// Calculate summary
	summary := calculateSummary(results)

	// Create JSON report structure
	report := JSONReport{
		Version:     "1.0",
		GeneratedAt: time.Now(),
		Summary:     summary,
		Results:     make([]JSONTestResult, 0, len(results)),
	}

	// Convert test results to JSON format
	for _, result := range results {
		jsonResult := JSONTestResult{
			Scenario:         result.Scenario,
			Status:           string(result.Status),
			StartTime:        result.StartTime,
			EndTime:          result.EndTime,
			Duration:         result.EndTime.Sub(result.StartTime).Seconds(),
			GoOutputPath:     result.GoOutputPath,
			DotNetOutputPath: result.DotNetOutputPath,
			GoError:          result.GoError,
			DotNetError:      result.DotNetError,
			ComparisonError:  result.ComparisonError,
		}

		if result.Comparison != nil {
			jsonResult.Comparison = &JSONComparisonResult{
				XMLMatch:    result.Comparison.XMLMatch,
				BinaryMatch: result.Comparison.BinaryMatch,
				VisualMatch: result.Comparison.VisualMatch,
			}
		}

		report.Results = append(report.Results, jsonResult)
	}

	// Write JSON file
	file, err := os.Create(jsonPath)
	if err != nil {
		return fmt.Errorf(
			"create JSON report file: %w",
			err,
		)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil &&
			err == nil {
			err = fmt.Errorf(
				"close JSON report file: %w",
				closeErr,
			)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf(
			"encode JSON report: %w",
			err,
		)
	}

	return nil
}

// calculateSummary computes summary statistics from test results
func calculateSummary(
	results []*TestResult,
) ReportSummary {
	var passed, failed, skipped int
	var totalDuration time.Duration

	for _, result := range results {
		duration := result.EndTime.Sub(result.StartTime)
		totalDuration += duration

		switch result.Status {
		case StatusPassed:
			passed++
		case StatusFailed:
			failed++
		case StatusSkipped:
			skipped++
		}
	}

	total := len(results)
	passRate := 0.0
	if total > 0 {
		passRate = float64(passed) / float64(total) * 100.0
	}

	return ReportSummary{
		Total:         total,
		Passed:        passed,
		Failed:        failed,
		Skipped:       skipped,
		PassRate:      passRate,
		TotalDuration: totalDuration,
	}
}

// ReportSummary contains aggregated test statistics
type ReportSummary struct {
	Total         int
	Passed        int
	Failed        int
	Skipped       int
	PassRate      float64
	TotalDuration time.Duration
}

// JSONReport represents the JSON report structure
type JSONReport struct {
	Version     string           `json:"version"`
	GeneratedAt time.Time        `json:"generated_at"`
	Summary     ReportSummary    `json:"summary"`
	Results     []JSONTestResult `json:"results"`
}

// JSONTestResult represents a single test result in JSON format
type JSONTestResult struct {
	Scenario         string                `json:"scenario"`
	Status           string                `json:"status"`
	StartTime        time.Time             `json:"start_time"`
	EndTime          time.Time             `json:"end_time"`
	Duration         float64               `json:"duration_seconds"`
	GoOutputPath     string                `json:"go_output_path,omitempty"`
	DotNetOutputPath string                `json:"dotnet_output_path,omitempty"`
	GoError          string                `json:"go_error,omitempty"`
	DotNetError      string                `json:"dotnet_error,omitempty"`
	ComparisonError  string                `json:"comparison_error,omitempty"`
	Comparison       *JSONComparisonResult `json:"comparison,omitempty"`
}

// JSONComparisonResult represents comparison results in JSON format
type JSONComparisonResult struct {
	XMLMatch    bool `json:"xml_match"`
	BinaryMatch bool `json:"binary_match"`
	VisualMatch bool `json:"visual_match"`
}

// GenerateHTMLReport is a convenience function that maintains backward compatibility
func GenerateHTMLReport(
	results []*TestResult,
	reportPath string,
) error {
	reporter := &Reporter{
		OutputDir: filepath.Dir(reportPath),
	}

	return reporter.generateHTMLReport(results, reportPath)
}

// htmlTemplate is an enhanced HTML template with interactive features
const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: #f5f5f5;
            padding: 20px;
            color: #333;
        }
        .container {
            max-width: 1400px;
            margin: 0 auto;
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 30px;
        }
        .header h1 {
            font-size: 28px;
            margin-bottom: 10px;
        }
        .header .meta {
            opacity: 0.9;
            font-size: 14px;
        }
        .controls {
            padding: 20px 30px;
            background: #f9fafb;
            border-bottom: 1px solid #e5e7eb;
            display: flex;
            gap: 15px;
            align-items: center;
        }
        .controls input[type="text"] {
            flex: 1;
            padding: 8px 12px;
            border: 1px solid #d1d5db;
            border-radius: 6px;
            font-size: 14px;
        }
        .controls select {
            padding: 8px 12px;
            border: 1px solid #d1d5db;
            border-radius: 6px;
            font-size: 14px;
            background: white;
        }
        .summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            padding: 30px;
            background: #f9fafb;
            border-bottom: 1px solid #e5e7eb;
        }
        .summary-card {
            background: white;
            padding: 20px;
            border-radius: 6px;
            border-left: 4px solid #667eea;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
        }
        .summary-card.passed {
            border-left-color: #10b981;
        }
        .summary-card.failed {
            border-left-color: #ef4444;
        }
        .summary-card.skipped {
            border-left-color: #f59e0b;
        }
        .summary-card .label {
            color: #6b7280;
            font-size: 14px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            margin-bottom: 8px;
        }
        .summary-card .value {
            font-size: 32px;
            font-weight: bold;
            color: #111827;
        }
        .results {
            padding: 30px;
        }
        .results h2 {
            font-size: 20px;
            margin-bottom: 20px;
            color: #111827;
        }
        .result-item {
            background: white;
            border: 1px solid #e5e7eb;
            border-radius: 6px;
            padding: 20px;
            margin-bottom: 15px;
            transition: box-shadow 0.2s;
        }
        .result-item:hover {
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }
        .result-item.hidden {
            display: none;
        }
        .result-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 10px;
        }
        .result-name {
            font-size: 18px;
            font-weight: 600;
            color: #111827;
        }
        .status {
            display: inline-block;
            padding: 6px 12px;
            border-radius: 4px;
            font-size: 12px;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        .status.passed {
            background: #d1fae5;
            color: #065f46;
        }
        .status.failed {
            background: #fee2e2;
            color: #991b1b;
        }
        .status.skipped {
            background: #fef3c7;
            color: #92400e;
        }
        .result-details {
            font-size: 14px;
            color: #6b7280;
            margin-top: 10px;
        }
        .result-details .detail-row {
            display: flex;
            justify-content: space-between;
            padding: 4px 0;
        }
        .comparison-section {
            margin-top: 15px;
            padding-top: 15px;
            border-top: 1px solid #e5e7eb;
        }
        .comparison-section h4 {
            font-size: 14px;
            color: #111827;
            margin-bottom: 10px;
        }
        .comparison-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 10px;
            margin-bottom: 10px;
        }
        .comparison-item {
            background: #f9fafb;
            padding: 10px;
            border-radius: 4px;
        }
        .comparison-item .label {
            font-size: 12px;
            color: #6b7280;
            margin-bottom: 5px;
        }
        .comparison-item .value {
            font-size: 14px;
            color: #111827;
            font-weight: 500;
        }
        .comparison-item .value.match {
            color: #10b981;
        }
        .comparison-item .value.mismatch {
            color: #ef4444;
        }
        .error {
            background: #fef2f2;
            border-left: 4px solid #ef4444;
            padding: 12px;
            margin-top: 10px;
            border-radius: 4px;
            font-family: 'Monaco', 'Menlo', 'Courier New', monospace;
            font-size: 13px;
            color: #991b1b;
            white-space: pre-wrap;
            word-break: break-all;
        }
        .footer {
            text-align: center;
            padding: 20px;
            color: #6b7280;
            font-size: 14px;
            border-top: 1px solid #e5e7eb;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title}}</h1>
            <div class="meta">Generated: {{.Generated}}</div>
        </div>

        <div class="controls">
            <input type="text" id="searchInput" placeholder="Search test scenarios...">
            <select id="statusFilter">
                <option value="all">All Tests</option>
                <option value="passed">Passed</option>
                <option value="failed">Failed</option>
                <option value="skipped">Skipped</option>
            </select>
        </div>

        <div class="summary">
            <div class="summary-card">
                <div class="label">Total Tests</div>
                <div class="value">{{.Summary.Total}}</div>
            </div>
            <div class="summary-card passed">
                <div class="label">Passed</div>
                <div class="value">{{.Summary.Passed}}</div>
            </div>
            <div class="summary-card failed">
                <div class="label">Failed</div>
                <div class="value">{{.Summary.Failed}}</div>
            </div>
            <div class="summary-card skipped">
                <div class="label">Skipped</div>
                <div class="value">{{.Summary.Skipped}}</div>
            </div>
            <div class="summary-card">
                <div class="label">Pass Rate</div>
                <div class="value">{{printf "%.1f" .Summary.PassRate}}%</div>
            </div>
            <div class="summary-card">
                <div class="label">Duration</div>
                <div class="value">{{printf "%.1fs" .Summary.TotalDuration.Seconds}}</div>
            </div>
        </div>

        <div class="results">
            <h2>Test Results</h2>
            {{range .Results}}
            <div class="result-item" data-status="{{.Status | lower}}" data-scenario="{{.Scenario}}">
                <div class="result-header">
                    <div class="result-name">{{.Scenario}}</div>
                    <span class="status {{.Status | lower}}">{{.Status}}</span>
                </div>
                <div class="result-details">
                    <div class="detail-row">
                        <span>Duration:</span>
                        <span>{{.EndTime.Sub .StartTime}}</span>
                    </div>
                    {{if .GoOutputPath}}
                    <div class="detail-row">
                        <span>Go Output:</span>
                        <span>{{.GoOutputPath}}</span>
                    </div>
                    {{end}}
                    {{if .DotNetOutputPath}}
                    <div class="detail-row">
                        <span>.NET Output:</span>
                        <span>{{.DotNetOutputPath}}</span>
                    </div>
                    {{end}}
                </div>
                {{if .Comparison}}
                <div class="comparison-section">
                    <h4>Comparison Results</h4>
                    <div class="comparison-grid">
                        <div class="comparison-item">
                            <div class="label">XML Match</div>
                            <div class="value {{if .Comparison.XMLMatch}}match{{else}}mismatch{{end}}">
                                {{if .Comparison.XMLMatch}}PASS{{else}}FAIL{{end}}
                            </div>
                        </div>
                        <div class="comparison-item">
                            <div class="label">Binary Match</div>
                            <div class="value {{if .Comparison.BinaryMatch}}match{{else}}mismatch{{end}}">
                                {{if .Comparison.BinaryMatch}}PASS{{else}}FAIL{{end}}
                            </div>
                        </div>
                        <div class="comparison-item">
                            <div class="label">Visual Match</div>
                            <div class="value {{if .Comparison.VisualMatch}}match{{else}}mismatch{{end}}">
                                {{if .Comparison.VisualMatch}}PASS{{else}}FAIL{{end}}
                            </div>
                        </div>
                    </div>
                </div>
                {{end}}
                {{if .GoError}}
                <div class="error">Go Error: {{.GoError}}</div>
                {{end}}
                {{if .DotNetError}}
                <div class="error">.NET Error: {{.DotNetError}}</div>
                {{end}}
                {{if .ComparisonError}}
                <div class="error">Comparison Error: {{.ComparisonError}}</div>
                {{end}}
            </div>
            {{end}}
        </div>

        <div class="footer">
            <p>E2E Cross-Runtime Testing Framework</p>
            <p>goffice vs Open-XML-SDK</p>
        </div>
    </div>

    <script>
        // Search functionality
        const searchInput = document.getElementById('searchInput');
        const statusFilter = document.getElementById('statusFilter');
        const resultItems = document.querySelectorAll('.result-item');

        function filterResults() {
            const searchTerm = searchInput.value.toLowerCase();
            const statusValue = statusFilter.value;

            resultItems.forEach(item => {
                const scenario = item.getAttribute('data-scenario').toLowerCase();
                const status = item.getAttribute('data-status');

                const matchesSearch = scenario.includes(searchTerm);
                const matchesStatus = statusValue === 'all' || status === statusValue;

                if (matchesSearch && matchesStatus) {
                    item.classList.remove('hidden');
                } else {
                    item.classList.add('hidden');
                }
            });
        }

        searchInput.addEventListener('input', filterResults);
        statusFilter.addEventListener('change', filterResults);
    </script>
</body>
</html>
`
