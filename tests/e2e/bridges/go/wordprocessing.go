package gobridge

import (
	"fmt"
	"strconv"

	"github.com/connerohnesorge/goffice/tests/e2e/framework"
	"github.com/connerohnesorge/goffice/wordprocessing"
	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// executeWordprocessing handles Word document scenarios
func (b *Bridge) executeWordprocessing(
	outputPath string,
) error {
	var doc *wordprocessing.Document

	for _, op := range b.scenario.Operations {
		switch op.Action {
		case "create_document":
			docType, _ := op.Properties["doc_type"].(string)
			var dt wordprocessing.DocType
			switch docType {
			case "document":
				dt = wordprocessing.DocTypeDocument
			case "template":
				dt = wordprocessing.DocTypeTemplate
			default:
				dt = wordprocessing.DocTypeDocument
			}

			var err error
			doc, err = wordprocessing.New(
				outputPath,
				dt,
			)
			if err != nil {
				return fmt.Errorf(
					"create document: %w",
					err,
				)
			}

		case "add_paragraph":
			if doc == nil {
				return fmt.Errorf(
					"document not created before add_paragraph",
				)
			}
			if err := b.addParagraph(doc, op); err != nil {
				return fmt.Errorf(
					"add paragraph: %w",
					err,
				)
			}

		case "add_table":
			if doc == nil {
				return fmt.Errorf(
					"document not created before add_table",
				)
			}
			if err := b.addTable(doc, op); err != nil {
				return fmt.Errorf(
					"add table: %w",
					err,
				)
			}

		default:
			return fmt.Errorf(
				"unknown action: %s",
				op.Action,
			)
		}
	}

	// Mark main part as dirty so it gets saved
	if doc != nil {
		mainPart := doc.MainPart()
		if mainPart != nil {
			mainPart.MarkDirty()
		}

		if err := doc.Save(); err != nil {
			return fmt.Errorf(
				"save document: %w",
				err,
			)
		}

		return doc.Close()
	}

	return nil
}

// addParagraph translates add_paragraph operation
func (b *Bridge) addParagraph(
	doc *wordprocessing.Document,
	op framework.Operation,
) error {
	mainPart := doc.MainPart()
	if mainPart == nil {
		return fmt.Errorf(
			"no main part in document",
		)
	}

	mainDoc := mainPart.Document()
	if mainDoc == nil {
		return fmt.Errorf(
			"no document in main part",
		)
	}

	body := mainDoc.Body()
	if body == nil {
		return fmt.Errorf("no body in document")
	}

	// Get paragraph text (if simple case)
	text, hasText := op.Properties["text"].(string)

	if hasText {
		// Simple case: single text string
		para := body.AppendParagraph(text)

		// Apply paragraph-level properties
		if props, ok := op.Properties["properties"].(map[string]interface{}); ok {
			b.applyParagraphProperties(
				para,
				props,
			)
		}

		return nil
	}

	// Complex case: nested operations (runs)
	operations, hasOps := op.Properties["operations"].([]interface{})
	if !hasOps {
		return fmt.Errorf(
			"paragraph must have either text or operations",
		)
	}

	para := body.AppendParagraph("")

	for _, opRaw := range operations {
		opMap, ok := opRaw.(map[string]interface{})
		if !ok {
			continue
		}

		action, _ := opMap["action"].(string)
		switch action {
		case "add_run":
			runText, _ := opMap["text"].(string)
			run := para.AppendRun(runText)

			// Apply run properties
			if props, ok := opMap["properties"].(map[string]interface{}); ok {
				b.applyRunProperties(run, props)
			}
		}
	}

	// Apply paragraph-level properties
	if props, ok := op.Properties["properties"].(map[string]interface{}); ok {
		b.applyParagraphProperties(para, props)
	}

	return nil
}

// applyParagraphProperties applies properties to paragraph
func (b *Bridge) applyParagraphProperties(
	para *elements.Paragraph,
	props map[string]interface{},
) {
	if len(props) == 0 {
		return
	}

	paraProps := para.GetOrCreateProperties()

	if align, ok := props["alignment"].(string); ok {
		var jc elements.JustificationValues
		switch align {
		case "left":
			jc = elements.JustificationValuesLeft
		case "center":
			jc = elements.JustificationValuesCenter
		case "right":
			jc = elements.JustificationValuesRight
		case "justify":
			jc = elements.JustificationValuesBoth
		default:
			jc = elements.JustificationValuesLeft
		}
		paraProps.SetJustification(jc)
	}

	if spacingAfter, ok := props["spacing_after"].(float64); ok {
		// Spacing is in points, convert to twips (1 point = 20 twips)
		spacing := paraProps.GetOrCreateSpacingBetweenLines()
		spacing.SetAfter(int(spacingAfter * 20))
	}

	if spacingBefore, ok := props["spacing_before"].(float64); ok {
		// Spacing is in points, convert to twips (1 point = 20 twips)
		spacing := paraProps.GetOrCreateSpacingBetweenLines()
		spacing.SetBefore(int(spacingBefore * 20))
	}

	if leftIndent, ok := props["left_indent"].(float64); ok {
		// Indent is in points, convert to twips (1 point = 20 twips)
		indentation := paraProps.GetOrCreateIndentation()
		indentation.SetLeft(int(leftIndent * 20))
	}
}

// applyRunProperties applies properties to run
func (b *Bridge) applyRunProperties(
	run *elements.Run,
	props map[string]interface{},
) {
	if len(props) == 0 {
		return
	}

	runProps := run.GetOrCreateProperties()

	if bold, ok := props["bold"].(bool); ok {
		runProps.SetBold(bold)
	}

	if italic, ok := props["italic"].(bool); ok {
		runProps.SetItalic(italic)
	}

	if underline, ok := props["underline"].(bool); ok &&
		underline {
		runProps.SetUnderline(
			elements.UnderlineSingle,
		)
	}

	if color, ok := props["color"].(string); ok {
		runProps.SetColor(color)
	}

	if fontSize, ok := props["font_size"].(float64); ok {
		// Font size in YAML is in points, API expects half-points
		runProps.SetFontSize(int(fontSize * 2))
	}

	if fontName, ok := props["font_name"].(string); ok {
		runProps.SetFont(fontName)
	}
}

// addTable translates add_table operation
func (b *Bridge) addTable(
	doc *wordprocessing.Document,
	op framework.Operation,
) error {
	mainPart := doc.MainPart()
	if mainPart == nil {
		return fmt.Errorf(
			"no main part in document",
		)
	}

	mainDoc := mainPart.Document()
	if mainDoc == nil {
		return fmt.Errorf(
			"no document in main part",
		)
	}

	body := mainDoc.Body()
	if body == nil {
		return fmt.Errorf("no body in document")
	}

	rows, _ := op.Properties["rows"].(float64)
	cols, _ := op.Properties["cols"].(float64)

	table := body.AppendTable(
		int(rows),
		int(cols),
	)

	// Set table properties
	if style, ok := op.Properties["style"].(string); ok {
		tableProps := table.GetOrCreateTableProperties()
		tableProps.SetTableStyle(style)
	}

	if widthPercent, ok := op.Properties["width_percent"].(float64); ok {
		tableProps := table.GetOrCreateTableProperties()
		// Width in percent: 5000 = 100%, so multiply by 50
		tableProps.SetTableWidth(
			int(widthPercent*50),
			elements.TableWidthTypePct,
		)
	}

	// Set cell values
	if cells, ok := op.Properties["cells"].([]interface{}); ok {
		for _, cellRaw := range cells {
			cellMap, ok := cellRaw.(map[string]interface{})
			if !ok {
				continue
			}

			row, _ := cellMap["row"].(float64)
			col, _ := cellMap["col"].(float64)
			text, _ := cellMap["text"].(string)

			// Convert float64 to int for indexing
			rowInt := int(row)
			colInt := int(col)

			// Get the cell and set its text
			cell := table.GetCell(rowInt, colInt)
			if cell != nil {
				// Clear existing content and add new paragraph with text
				cell.RemoveAllChildren()
				para := elements.NewParagraph(
					text,
				)
				cell.AppendChild(para)
			}
		}
	}

	return nil
}

// Helper function to convert interface{} to int
func toInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", v)
	}
}
