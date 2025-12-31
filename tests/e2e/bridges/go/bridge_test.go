package gobridge

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/tests/e2e/framework"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestBridge_Execute(t *testing.T) {
	t.Run(
		"UnknownDocumentType",
		func(t *testing.T) {
			scenario := &framework.TestScenario{
				Name:         "unknown-type",
				DocumentType: "unknown",
			}

			bridge := NewBridge(scenario)
			err := bridge.Execute("test.docx")

			if err == nil {
				t.Fatal(
					"expected error for unknown document type, got nil",
				)
			}
			if err.Error() != "unknown document type: unknown" {
				t.Errorf(
					"unexpected error message: %s",
					err.Error(),
				)
			}
		},
	)

	t.Run(
		"SpreadsheetNotImplemented",
		func(t *testing.T) {
			scenario := &framework.TestScenario{
				Name:         "test-spreadsheet",
				DocumentType: "spreadsheet",
			}

			bridge := NewBridge(scenario)
			err := bridge.Execute("test.xlsx")

			if err == nil {
				t.Fatal(
					"expected error for spreadsheet (not implemented), got nil",
				)
			}
		},
	)

	t.Run(
		"PresentationNotImplemented",
		func(t *testing.T) {
			scenario := &framework.TestScenario{
				Name:         "test-presentation",
				DocumentType: "presentation",
			}

			bridge := NewBridge(scenario)
			err := bridge.Execute("test.pptx")

			if err == nil {
				t.Fatal(
					"expected error for presentation (not implemented), got nil",
				)
			}
		},
	)
}

func TestBridge_BasicParagraph(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "basic-paragraph",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "create_document",
				Properties: map[string]interface{}{
					"doc_type": "document",
				},
			},
			{
				Action: "add_paragraph",
				Properties: map[string]interface{}{
					"text": "Hello World",
				},
			},
		},
	}

	// tmpDir := t.TempDir()
	outputPath := "/tmp/test-basic-paragraph.docx"

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)
	if err != nil {
		t.Fatalf(
			"failed to execute scenario: %v",
			err,
		)
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Fatalf(
			"output file was not created: %s",
			outputPath,
		)
	}

	// Verify file is a valid .docx
	doc, err := wordprocessing.Open(
		outputPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"failed to open generated document: %v",
			err,
		)
	}
	defer doc.Close()

	// Verify content
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("no main part in document")
	}

	mainDoc := mainPart.Document()
	if mainDoc == nil {
		t.Fatal("no document in main part")
	}

	// Force reload from XML to work around goffice bug
	if err := mainDoc.Reload(); err != nil {
		t.Fatalf(
			"failed to reload document: %v",
			err,
		)
	}

	body := mainDoc.Body()
	if body == nil {
		t.Fatal("no body in document")
	}

	// Debug: check if body exists and what it contains
	if body == nil {
		t.Fatal("body is nil")
	}

	// Count all children
	childCount := 0
	for child := range body.Children() {
		childCount++
		t.Logf(
			"Body child %d: %s:%s",
			childCount,
			child.NamespaceURI(),
			child.LocalName(),
		)
	}
	t.Logf("Total body children: %d", childCount)

	// Count paragraphs
	paraCount := 0
	for para := range body.Paragraphs() {
		paraCount++
		// Check first paragraph text
		if paraCount == 1 {
			text := para.InnerText()
			if text != "Hello World" {
				t.Errorf(
					"expected paragraph text 'Hello World', got '%s'",
					text,
				)
			}
		}
	}

	if paraCount == 0 {
		t.Error("no paragraphs found in document")
	}
}

func TestBridge_FormattedText(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "formatted-text",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "create_document",
				Properties: map[string]interface{}{
					"doc_type": "document",
				},
			},
			{
				Action: "add_paragraph",
				Properties: map[string]interface{}{
					"operations": []interface{}{
						map[string]interface{}{
							"action": "add_run",
							"text":   "Normal text ",
						},
						map[string]interface{}{
							"action": "add_run",
							"text":   "bold text ",
							"properties": map[string]interface{}{
								"bold": true,
							},
						},
						map[string]interface{}{
							"action": "add_run",
							"text":   "italic text ",
							"properties": map[string]interface{}{
								"italic": true,
							},
						},
						map[string]interface{}{
							"action": "add_run",
							"text":   "colored",
							"properties": map[string]interface{}{
								"color": "FF0000",
								"bold":  true,
								"font_size": float64(
									12,
								),
								"font_name": "Calibri",
							},
						},
					},
				},
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"formatted-text.docx",
	)

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)
	if err != nil {
		t.Fatalf(
			"failed to execute scenario: %v",
			err,
		)
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Fatalf(
			"output file was not created: %s",
			outputPath,
		)
	}

	// Verify file is a valid .docx
	doc, err := wordprocessing.Open(
		outputPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"failed to open generated document: %v",
			err,
		)
	}
	defer doc.Close()

	// Verify content
	mainPart := doc.MainPart()
	body := mainPart.Document().Body()

	paraCount := 0
	for para := range body.Paragraphs() {
		paraCount++
		if paraCount == 1 {
			// Verify runs
			runCount := 0
			for run := range para.Runs() {
				runCount++
				props := run.Properties()

				switch runCount {
				case 1:
					// Normal text - no properties expected
					if props != nil &&
						props.Bold() {
						t.Error(
							"first run should not be bold",
						)
					}
				case 2:
					// Bold text
					if props == nil ||
						!props.Bold() {
						t.Error(
							"second run should be bold",
						)
					}
				case 3:
					// Italic text
					if props == nil ||
						!props.Italic() {
						t.Error(
							"third run should be italic",
						)
					}
				case 4:
					// Colored, bold, sized, named font
					if props == nil {
						t.Fatal(
							"fourth run should have properties",
						)
					}
					if !props.Bold() {
						t.Error(
							"fourth run should be bold",
						)
					}
					if props.Color() != "FF0000" {
						t.Errorf(
							"expected color FF0000, got %s",
							props.Color(),
						)
					}
					if props.FontSize() != 24 { // 12 points = 24 half-points
						t.Errorf(
							"expected font size 24 half-points, got %d",
							props.FontSize(),
						)
					}
				}
			}

			if runCount != 4 {
				t.Errorf(
					"expected 4 runs, got %d",
					runCount,
				)
			}
		}
	}

	if paraCount == 0 {
		t.Error("no paragraphs found in document")
	}
}

func TestBridge_Table(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "simple-table",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "create_document",
				Properties: map[string]interface{}{
					"doc_type": "document",
				},
			},
			{
				Action: "add_table",
				Properties: map[string]interface{}{
					"rows": float64(2),
					"cols": float64(3),
					"cells": []interface{}{
						map[string]interface{}{
							"row":  float64(0),
							"col":  float64(0),
							"text": "Cell 1,1",
						},
						map[string]interface{}{
							"row":  float64(0),
							"col":  float64(1),
							"text": "Cell 1,2",
						},
						map[string]interface{}{
							"row":  float64(1),
							"col":  float64(0),
							"text": "Cell 2,1",
						},
					},
				},
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"simple-table.docx",
	)

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)
	if err != nil {
		t.Fatalf(
			"failed to execute scenario: %v",
			err,
		)
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(
		err,
	) {
		t.Fatalf(
			"output file was not created: %s",
			outputPath,
		)
	}

	// Verify file is a valid .docx
	doc, err := wordprocessing.Open(
		outputPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"failed to open generated document: %v",
			err,
		)
	}
	defer doc.Close()

	// Verify content
	mainPart := doc.MainPart()
	body := mainPart.Document().Body()

	tableCount := 0
	for table := range body.Tables() {
		tableCount++

		if tableCount == 1 {
			// Verify table dimensions
			rowCount := table.RowCount()
			if rowCount != 2 {
				t.Errorf(
					"expected 2 rows, got %d",
					rowCount,
				)
			}

			// Verify cell content
			cell := table.GetCell(0, 0)
			if cell == nil {
				t.Fatal("cell (0,0) is nil")
			}
			cellText := cell.InnerText()
			if cellText != "Cell 1,1" {
				t.Errorf(
					"expected cell (0,0) text 'Cell 1,1', got '%s'",
					cellText,
				)
			}

			cell = table.GetCell(0, 1)
			if cell == nil {
				t.Fatal("cell (0,1) is nil")
			}
			cellText = cell.InnerText()
			if cellText != "Cell 1,2" {
				t.Errorf(
					"expected cell (0,1) text 'Cell 1,2', got '%s'",
					cellText,
				)
			}

			cell = table.GetCell(1, 0)
			if cell == nil {
				t.Fatal("cell (1,0) is nil")
			}
			cellText = cell.InnerText()
			if cellText != "Cell 2,1" {
				t.Errorf(
					"expected cell (1,0) text 'Cell 2,1', got '%s'",
					cellText,
				)
			}
		}
	}

	if tableCount == 0 {
		t.Error("no tables found in document")
	}
}

func TestBridge_ParagraphAlignment(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "paragraph-alignment",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "create_document",
				Properties: map[string]interface{}{
					"doc_type": "document",
				},
			},
			{
				Action: "add_paragraph",
				Properties: map[string]interface{}{
					"text": "Centered text",
					"properties": map[string]interface{}{
						"alignment": "center",
					},
				},
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"paragraph-alignment.docx",
	)

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)
	if err != nil {
		t.Fatalf(
			"failed to execute scenario: %v",
			err,
		)
	}

	// Verify file was created and is valid
	doc, err := wordprocessing.Open(
		outputPath,
		false,
	)
	if err != nil {
		t.Fatalf(
			"failed to open generated document: %v",
			err,
		)
	}
	defer doc.Close()

	// Verify alignment
	mainPart := doc.MainPart()
	body := mainPart.Document().Body()

	paraCount := 0
	for para := range body.Paragraphs() {
		paraCount++
		if paraCount == 1 {
			props := para.Properties()
			if props == nil {
				t.Fatal(
					"paragraph should have properties",
				)
			}
			jc := props.Justification()
			if jc != elements.JustificationValuesCenter {
				t.Errorf(
					"expected center alignment, got %s",
					jc,
				)
			}
		}
	}
}

func TestBridge_MissingDocument(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "missing-document",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "add_paragraph",
				Properties: map[string]interface{}{
					"text": "This should fail",
				},
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"missing-document.docx",
	)

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)

	if err == nil {
		t.Fatal(
			"expected error when adding paragraph without creating document first",
		)
	}
}

func TestBridge_UnknownAction(t *testing.T) {
	scenario := &framework.TestScenario{
		Name:         "unknown-action",
		DocumentType: "wordprocessing",
		Operations: []framework.Operation{
			{
				Action: "create_document",
				Properties: map[string]interface{}{
					"doc_type": "document",
				},
			},
			{
				Action: "unknown_action",
				Properties: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
	}

	tmpDir := t.TempDir()
	outputPath := filepath.Join(
		tmpDir,
		"unknown-action.docx",
	)

	bridge := NewBridge(scenario)
	err := bridge.Execute(outputPath)

	if err == nil {
		t.Fatal(
			"expected error for unknown action",
		)
	}
	if err.Error() != "unknown action: unknown_action" {
		t.Errorf(
			"unexpected error message: %s",
			err.Error(),
		)
	}
}
