package framework

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

// GenerateHTMLReport generates an HTML report of test results
func GenerateHTMLReport(
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
	var passed, failed, skipped int
	for _, result := range results {
		switch result.Status {
		case StatusPassed:
			passed++
		case StatusFailed:
			failed++
		case StatusSkipped:
			skipped++
		}
	}

	// Prepare template data
	data := struct {
		Title     string
		Generated string
		Total     int
		Passed    int
		Failed    int
		Skipped   int
		PassRate  float64
		Results   []*TestResult
	}{
		Title: "E2E Test Results",
		Generated: time.Now().
			Format("2006-01-02 15:04:05"),
		Total:    len(results),
		Passed:   passed,
		Failed:   failed,
		Skipped:  skipped,
		PassRate: 0.0,
		Results:  results,
	}

	if len(results) > 0 {
		data.PassRate = float64(
			passed,
		) / float64(
			len(results),
		) * 100.0
	}

	// Create template with custom functions
	funcMap := template.FuncMap{
		"lower": func(s TestStatus) string {
			switch s {
			case StatusPassed:
				return "passed"
			case StatusFailed:
				return "failed"
			case StatusSkipped:
				return "skipped"
			default:
				return string(s)
			}
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

// htmlTemplate is a basic HTML template for the test report
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
            max-width: 1200px;
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

        <div class="summary">
            <div class="summary-card">
                <div class="label">Total Tests</div>
                <div class="value">{{.Total}}</div>
            </div>
            <div class="summary-card passed">
                <div class="label">Passed</div>
                <div class="value">{{.Passed}}</div>
            </div>
            <div class="summary-card failed">
                <div class="label">Failed</div>
                <div class="value">{{.Failed}}</div>
            </div>
            <div class="summary-card skipped">
                <div class="label">Skipped</div>
                <div class="value">{{.Skipped}}</div>
            </div>
            <div class="summary-card">
                <div class="label">Pass Rate</div>
                <div class="value">{{printf "%.1f" .PassRate}}%</div>
            </div>
        </div>

        <div class="results">
            <h2>Test Results</h2>
            {{range .Results}}
            <div class="result-item">
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
</body>
</html>
`
