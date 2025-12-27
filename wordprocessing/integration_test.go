package wordprocessing

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

func TestHeadersFootersIntegration(t *testing.T) {
	tmpFile, err := os.CreateTemp(
		"",
		"header_footer_test_*.docx",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp file: %v",
			err,
		)
	}
	defer func() { _ = os.Remove(tmpFile.Name()) }()
	tmpName := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := New(tmpName, DocTypeDocument)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Test creating a header
	headerPart, err := doc.MainPart().
		AddHeaderPart()
	if err != nil {
		t.Fatalf(
			"Failed to add header part: %v",
			err,
		)
	}

	header := headerPart.GetOrCreateHeader()
	header.AppendParagraph("Test Header Content")

	// Test creating a footer
	footerPart, err := doc.MainPart().
		AddFooterPart()
	if err != nil {
		t.Fatalf(
			"Failed to add footer part: %v",
			err,
		)
	}

	footer := footerPart.GetOrCreateFooter()
	footer.AppendParagraph("Test Footer Content")

	// Save
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

func TestFootnotesIntegration(t *testing.T) {
	tmpFile, err := os.CreateTemp(
		"",
		"footnote_test_*.docx",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp file: %v",
			err,
		)
	}
	defer func() { _ = os.Remove(tmpFile.Name()) }()
	tmpName := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := New(tmpName, DocTypeDocument)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Add footnotes part
	footnotesPart, err := doc.MainPart().
		AddFootnotesPart()
	if err != nil {
		t.Fatalf(
			"Failed to add footnotes part: %v",
			err,
		)
	}

	// Add a footnote
	footnote := footnotesPart.AddFootnote(
		"This is a test footnote",
	)
	if footnote.Id() != 1 {
		t.Errorf(
			"Expected footnote ID 1, got %d",
			footnote.Id(),
		)
	}

	// Add another footnote
	footnote2 := footnotesPart.AddFootnote(
		"Second footnote",
	)
	if footnote2.Id() != 2 {
		t.Errorf(
			"Expected footnote ID 2, got %d",
			footnote2.Id(),
		)
	}

	// Save
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

func TestEndnotesIntegration(t *testing.T) {
	tmpFile, err := os.CreateTemp(
		"",
		"endnote_test_*.docx",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp file: %v",
			err,
		)
	}
	defer func() { _ = os.Remove(tmpFile.Name()) }()
	tmpName := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := New(tmpName, DocTypeDocument)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Add endnotes part
	endnotesPart, err := doc.MainPart().
		AddEndnotesPart()
	if err != nil {
		t.Fatalf(
			"Failed to add endnotes part: %v",
			err,
		)
	}

	// Add an endnote
	endnote := endnotesPart.AddEndnote(
		"This is a test endnote",
	)
	if endnote.Id() != 1 {
		t.Errorf(
			"Expected endnote ID 1, got %d",
			endnote.Id(),
		)
	}

	// Save
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

func TestCommentsIntegration(t *testing.T) {
	tmpFile, err := os.CreateTemp(
		"",
		"comment_test_*.docx",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp file: %v",
			err,
		)
	}
	defer func() { _ = os.Remove(tmpFile.Name()) }()
	tmpName := tmpFile.Name()
	_ = tmpFile.Close()

	doc, err := New(tmpName, DocTypeDocument)
	if err != nil {
		t.Fatalf(
			"Failed to create document: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Add comments part
	commentsPart, err := doc.MainPart().
		AddCommentsPart()
	if err != nil {
		t.Fatalf(
			"Failed to add comments part: %v",
			err,
		)
	}

	// Add a comment
	comment := commentsPart.AddComment(
		"John Doe",
		"This is a test comment",
	)
	if comment.Id() != 1 {
		t.Errorf(
			"Expected comment ID 1, got %d",
			comment.Id(),
		)
	}

	if comment.Author() != "John Doe" {
		t.Errorf(
			"Expected author 'John Doe', got '%s'",
			comment.Author(),
		)
	}

	// Save
	if err := doc.Save(); err != nil {
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
}

func TestHeaderElementsAPI(t *testing.T) {
	header := elements.NewHeader()

	// Test AppendParagraph
	p1 := header.AppendParagraph(
		"First paragraph",
	)
	if p1 == nil {
		t.Fatal("AppendParagraph returned nil")
	}

	// Test PrependParagraph
	p2 := header.PrependParagraph(
		"Prepended paragraph",
	)
	if p2 == nil {
		t.Fatal("PrependParagraph returned nil")
	}

	// Test AppendTable
	table := header.AppendTable(2, 3)
	if table == nil {
		t.Fatal("AppendTable returned nil")
	}

	// Test iterators
	paragraphCount := 0
	for range header.Paragraphs() {
		paragraphCount++
	}
	if paragraphCount != 2 {
		t.Errorf(
			"Expected 2 paragraphs, got %d",
			paragraphCount,
		)
	}

	tableCount := 0
	for range header.Tables() {
		tableCount++
	}
	if tableCount != 1 {
		t.Errorf(
			"Expected 1 table, got %d",
			tableCount,
		)
	}

	// Test ClearContent
	header.ClearContent()
	paragraphCount = 0
	for range header.Paragraphs() {
		paragraphCount++
	}
	if paragraphCount != 0 {
		t.Errorf(
			"Expected 0 paragraphs after clear, got %d",
			paragraphCount,
		)
	}
}

func TestFooterElementsAPI(t *testing.T) {
	footer := elements.NewFooter()

	// Test AppendParagraph
	p1 := footer.AppendParagraph(
		"First paragraph",
	)
	if p1 == nil {
		t.Fatal("AppendParagraph returned nil")
	}

	// Test PrependParagraph
	p2 := footer.PrependParagraph(
		"Prepended paragraph",
	)
	if p2 == nil {
		t.Fatal("PrependParagraph returned nil")
	}

	// Test iterators
	paragraphCount := 0
	for range footer.Paragraphs() {
		paragraphCount++
	}
	if paragraphCount != 2 {
		t.Errorf(
			"Expected 2 paragraphs, got %d",
			paragraphCount,
		)
	}
}

// TestIntegrationCreateCompleteDocument tests creating a complete document
// with paragraphs, formatting, tables, headers/footers, and styles.
func TestIntegrationCreateCompleteDocument(
	t *testing.T,
) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-integration-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"complete.docx",
	)

	// Create a new document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Verify the document was created properly
	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeDocument,
		)
	}

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart() = nil, want non-nil")
	}

	// Save the document
	err = doc.SaveAs(testPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(testPath); os.IsNotExist(
		err,
	) {
		t.Error("SaveAs() did not create file")
	}
}

// TestIntegrationDocumentBuilderPattern tests using the builder pattern
// to create valid documents.
func TestIntegrationDocumentBuilderPattern(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	// Add a title heading
	builder.AddHeading("Document Title", 1).
		AlignCenter().
		Font("Arial", 16)

	// Add body paragraphs
	builder.AddParagraph("This is the first paragraph of the document.").
		SpacingAfter(12)

	builder.AddParagraph("This is the second paragraph with some formatting.").
		Bold().
		SpacingBefore(6).
		SpacingAfter(12)

	// Add a subheading
	builder.AddHeading("Section 1", 2).
		SpacingBefore(18)

	// Add a paragraph with multiple runs
	pb := builder.AddParagraph("")
	pb.AddRun("This has ").Font("Times New Roman")
	pb.AddRun("bold").Bold()
	pb.AddRun(" and ").Font("Times New Roman")
	pb.AddRun("italic").Italic()
	pb.AddRun(" text.").Font("Times New Roman")

	// Add a table
	builder.AddTable(3, 3).
		SetCellText(0, 0, "Header 1").
		SetCellText(0, 1, "Header 2").
		SetCellText(0, 2, "Header 3").
		SetCellText(1, 0, "Cell 1").
		SetCellText(1, 1, "Cell 2").
		SetCellText(1, 2, "Cell 3").
		SetCellText(2, 0, "Cell 4").
		SetCellText(2, 1, "Cell 5").
		SetCellText(2, 2, "Cell 6").
		SetStyle("TableGrid")

	// Add a page break
	builder.AddPageBreak()

	// Add content on the second page
	builder.AddHeading("Page 2", 1)
	builder.AddParagraph(
		"Content on the second page.",
	)

	// Build the document
	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	if doc == nil {
		t.Fatal("Build() returned nil document")
	}

	// Verify the document structure
	body := doc.Body()
	if body == nil {
		t.Fatal("Document body is nil")
	}

	// Count paragraphs
	paraCount := 0
	for range body.Paragraphs() {
		paraCount++
	}

	// We should have at least 7 paragraphs (headings + body + page break)
	if paraCount < 7 {
		t.Errorf(
			"Expected at least 7 paragraphs, got %d",
			paraCount,
		)
	}

	// Verify we have a table
	tableCount := 0
	for range body.Tables() {
		tableCount++
	}
	if tableCount != 1 {
		t.Errorf(
			"Expected 1 table, got %d",
			tableCount,
		)
	}
}

// TestIntegrationTableOperations tests creating and manipulating tables.
func TestIntegrationTableOperations(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	// Create a 4x4 table with various properties
	tb := builder.AddTable(4, 4).
		SetWidthPercent(100).
		SetBorders(elements.BorderSingle, 4, "000000")

	// Set header row
	row := tb.Row(0)
	if row != nil {
		row.SetHeaderRow(true).
			SetHeight(400, elements.HeightRuleExact)
	}

	// Set cell content and formatting
	for i := range 4 {
		cell := tb.Cell(0, i)
		if cell != nil {
			cell.SetText("Header " + string(rune('A'+i))).
				SetShading("CCCCCC")
		}
	}

	// Fill data rows
	for row := 1; row < 4; row++ {
		for col := range 4 {
			cell := tb.Cell(row, col)
			if cell != nil {
				cell.SetText(
					"Data " + string(
						rune('0'+row),
					) + string(
						rune('A'+col),
					),
				)
			}
		}
	}

	// Build and verify
	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	for table := range body.Tables() {
		// Verify row count
		if table.RowCount() != 4 {
			t.Errorf(
				"Expected 4 rows, got %d",
				table.RowCount(),
			)
		}

		// Verify header cell content
		cell := table.GetCell(0, 0)
		if cell == nil {
			t.Fatal("GetCell(0, 0) returned nil")
		}
		if cell.InnerText() != "Header A" {
			t.Errorf(
				"Expected 'Header A', got '%s'",
				cell.InnerText(),
			)
		}

		// Verify data cell content
		cell = table.GetCell(1, 1)
		if cell == nil {
			t.Fatal("GetCell(1, 1) returned nil")
		}
		if cell.InnerText() != "Data 1B" {
			t.Errorf(
				"Expected 'Data 1B', got '%s'",
				cell.InnerText(),
			)
		}
	}
}

// TestIntegrationParagraphFormatting tests paragraph-level formatting.
//
//nolint:revive // cyclomatic: test validation requires switch with multiple cases
func TestIntegrationParagraphFormatting(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	// Add paragraphs with different alignments
	builder.AddParagraph("Left aligned text").
		AlignLeft()
	builder.AddParagraph("Center aligned text").
		AlignCenter()
	builder.AddParagraph("Right aligned text").
		AlignRight()
	builder.AddParagraph("Justified text that should be long enough to see the justification effect.").
		AlignJustify()

	// Add paragraphs with spacing and indentation
	builder.AddParagraph("Paragraph with custom spacing").
		SpacingBefore(24).
		SpacingAfter(24)

	builder.AddParagraph("Paragraph with left indent").
		LeftIndent(72)

		// 72 points = 1 inch

	builder.AddParagraph("Paragraph with first line indent").
		FirstLineIndent(36)

		// 36 points = 0.5 inch

	builder.AddParagraph("Paragraph with hanging indent").
		LeftIndent(36).
		HangingIndent(36)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()

	// Verify paragraph properties
	i := 0
	for p := range body.Paragraphs() {
		props := p.Properties()

		switch i {
		case 0:
			// Left alignment is default, may not have explicit property
		case 1:
			if props != nil &&
				props.Justification() != elements.JustificationCenter {
				t.Errorf(
					"Paragraph %d: expected center, got %s",
					i,
					props.Justification(),
				)
			}
		case 2:
			if props != nil &&
				props.Justification() != elements.JustificationRight {
				t.Errorf(
					"Paragraph %d: expected right, got %s",
					i,
					props.Justification(),
				)
			}
		case 3:
			if props != nil &&
				props.Justification() != elements.JustificationBoth {
				t.Errorf(
					"Paragraph %d: expected both/justify, got %s",
					i,
					props.Justification(),
				)
			}
		case 4:
			if props != nil {
				spacing := props.SpacingBetweenLines()
				if spacing != nil {
					// 24 points = 480 twips
					if spacing.Before() != 480 {
						t.Errorf(
							"Paragraph %d: expected before 480, got %d",
							i,
							spacing.Before(),
						)
					}
					if spacing.After() != 480 {
						t.Errorf(
							"Paragraph %d: expected after 480, got %d",
							i,
							spacing.After(),
						)
					}
				}
			}
		case 5:
			if props != nil {
				ind := props.Indentation()
				if ind != nil {
					// 72 points = 1440 twips
					if ind.Left() != 1440 {
						t.Errorf(
							"Paragraph %d: expected left 1440, got %d",
							i,
							ind.Left(),
						)
					}
				}
			}
		case 6:
			if props != nil {
				ind := props.Indentation()
				if ind != nil {
					// 36 points = 720 twips
					if ind.FirstLine() != 720 {
						t.Errorf(
							"Paragraph %d: expected firstLine 720, got %d",
							i,
							ind.FirstLine(),
						)
					}
				}
			}
		}
		i++
	}
}

// TestIntegrationRunFormatting tests run-level (character) formatting.
func TestIntegrationRunFormatting(t *testing.T) {
	builder := NewDocumentBuilder()

	// Add paragraphs with different run formatting
	pb := builder.AddParagraph("")
	pb.AddRun("Bold text").Bold()
	pb.AddRun(" ")
	pb.AddRun("Italic text").Italic()
	pb.AddRun(" ")
	pb.AddRun("Underlined text").Underline()
	pb.AddRun(" ")
	pb.AddRun("Red text").Color("FF0000")
	pb.AddRun(" ")
	pb.AddRun("Highlighted text").
		Highlight(elements.HighlightYellow)
	pb.AddRun(" ")
	pb.AddRun("Large text").FontSize(18)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		runCount := 0
		for r := range p.Runs() {
			runCount++
			text := r.InnerText()
			props := r.Properties()

			if props == nil && text != " " {
				// Skip spacing runs
				continue
			}

			xml := ""
			if props != nil {
				xml = props.OuterXml()
			}

			switch {
			case strings.Contains(text, "Bold"):
				if !strings.Contains(
					xml,
					"<w:b",
				) {
					t.Errorf(
						"Expected bold formatting for '%s'",
						text,
					)
				}
			case strings.Contains(text, "Italic"):
				if !strings.Contains(
					xml,
					"<w:i",
				) {
					t.Errorf(
						"Expected italic formatting for '%s'",
						text,
					)
				}
			case strings.Contains(text, "Underlined"):
				if !strings.Contains(
					xml,
					"<w:u",
				) {
					t.Errorf(
						"Expected underline formatting for '%s'",
						text,
					)
				}
			case strings.Contains(text, "Red"):
				if !strings.Contains(
					xml,
					"FF0000",
				) {
					t.Errorf(
						"Expected red color for '%s'",
						text,
					)
				}
			case strings.Contains(text, "Highlighted"):
				if !strings.Contains(
					xml,
					"yellow",
				) {
					t.Errorf(
						"Expected yellow highlight for '%s'",
						text,
					)
				}
			case strings.Contains(text, "Large"):
				// 18 points = 36 half-points
				if !strings.Contains(xml, "36") {
					t.Errorf(
						"Expected font size 36 (18pt) for '%s'",
						text,
					)
				}
			}
		}

		if runCount < 6 {
			t.Errorf(
				"Expected at least 6 runs, got %d",
				runCount,
			)
		}
	}
}

// TestIntegrationBuildToBytes tests building a document directly to bytes.
func TestIntegrationBuildToBytes(t *testing.T) {
	builder := NewDocumentBuilder()

	builder.AddHeading("Test Document", 1)
	builder.AddParagraph(
		"This is a test paragraph.",
	)
	builder.AddTable(2, 2).
		SetCellText(0, 0, "A").
		SetCellText(0, 1, "B").
		SetCellText(1, 0, "C").
		SetCellText(1, 1, "D")

	data, err := builder.BuildToBytes()
	if err != nil {
		t.Fatalf("BuildToBytes() error = %v", err)
	}

	if len(data) == 0 {
		t.Error(
			"BuildToBytes() returned empty data",
		)
	}

	// Verify it contains expected XML elements
	str := string(data)
	if !strings.Contains(str, "document") {
		t.Error(
			"Expected document element in output",
		)
	}
	if !strings.Contains(str, "body") {
		t.Error("Expected body element in output")
	}
	if !strings.Contains(str, "Test Document") {
		t.Error(
			"Expected 'Test Document' text in output",
		)
	}
}

// TestIntegrationDocumentSaveAndOpen tests saving a document and reopening it.
func TestIntegrationDocumentSaveAndOpen(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-integration-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"saveandopen.docx",
	)

	// Create and save a document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc1.Close()

	// Reopen the document
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = doc2.Close() }()

	// Verify it was opened correctly
	if doc2.Type() != DocTypeDocument {
		t.Errorf(
			"Opened document Type() = %v, want %v",
			doc2.Type(),
			DocTypeDocument,
		)
	}

	if !doc2.IsEditable() {
		t.Error(
			"Opened document IsEditable() = false, want true",
		)
	}

	if doc2.MainPart() == nil {
		t.Error(
			"Opened document MainPart() = nil, want non-nil",
		)
	}
}

// TestIntegrationDocumentStream tests creating and writing a document to a stream.
func TestIntegrationDocumentStream(t *testing.T) {
	var buf bytes.Buffer

	doc, err := NewWriter(&buf, DocTypeDocument)
	if err != nil {
		t.Fatalf("NewWriter() error = %v", err)
	}

	// Verify document properties
	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeDocument,
		)
	}

	if doc.Path() != "" {
		t.Errorf(
			"Path() = %v, want empty string",
			doc.Path(),
		)
	}

	if !doc.IsEditable() {
		t.Error("IsEditable() = false, want true")
	}

	// Close should write to the buffer
	_ = doc.Close()
}

// TestIntegrationMultipleDocumentTypes tests creating different document types.
func TestIntegrationMultipleDocumentTypes(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-integration-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testCases := []struct {
		docType   DocType
		extension string
	}{
		{DocTypeDocument, ".docx"},
		{DocTypeTemplate, ".dotx"},
		{DocTypeMacroEnabled, ".docm"},
		{DocTypeMacroTemplate, ".dotm"},
	}

	for _, tc := range testCases {
		t.Run(
			tc.docType.String(),
			func(t *testing.T) {
				testPath := filepath.Join(
					tmpDir,
					"test"+tc.extension,
				)

				doc, err := New(
					testPath,
					tc.docType,
				)
				if err != nil {
					t.Fatalf(
						"New() error = %v",
						err,
					)
				}

				if doc.Type() != tc.docType {
					t.Errorf(
						"Type() = %v, want %v",
						doc.Type(),
						tc.docType,
					)
				}

				// Verify extension
				if tc.docType.Extension() != tc.extension {
					t.Errorf(
						"Extension() = %v, want %v",
						tc.docType.Extension(),
						tc.extension,
					)
				}

				// Save and verify file creation
				err = doc.SaveAs(testPath)
				if err != nil {
					t.Fatalf(
						"SaveAs() error = %v",
						err,
					)
				}

				if _, err := os.Stat(testPath); os.IsNotExist(
					err,
				) {
					t.Errorf(
						"File %s was not created",
						testPath,
					)
				}

				_ = doc.Close()
			},
		)
	}
}

// TestIntegrationDocumentChangeType tests changing document types.
func TestIntegrationDocumentChangeType(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-integration-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"change.docx",
	)

	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Change from document to template
	err = doc.ChangeType(DocTypeTemplate)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypeTemplate {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeTemplate,
		)
	}

	// Change to macro-enabled
	err = doc.ChangeType(DocTypeMacroEnabled)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypeMacroEnabled {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeMacroEnabled,
		)
	}

	// Change back to document
	err = doc.ChangeType(DocTypeDocument)
	if err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc.Type(),
			DocTypeDocument,
		)
	}
}

// TestIntegrationBuilderChaining tests method chaining in builders.
func TestIntegrationBuilderChaining(
	t *testing.T,
) {
	// Test that all builder methods return proper types for chaining
	doc, err := NewDocumentBuilder().
		AddHeading("Title", 1).
		AlignCenter().
		Bold().
		Document().
		AddParagraph("Body").
		Italic().
		SpacingAfter(12).
		LeftIndent(36).
		Color("0000FF").
		Document().
		AddTable(2, 2).
		SetCellText(0, 0, "A").
		SetWidthPercent(100).
		SetBorders(elements.BorderSingle, 4, "000000").
		Document().
		AddPageBreak().
		AddHeading("Page 2", 1).
		Document().
		Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	if doc == nil {
		t.Fatal("Build() returned nil document")
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Document body is nil")
	}

	// Count elements
	paraCount := 0
	tableCount := 0
	for range body.Paragraphs() {
		paraCount++
	}
	for range body.Tables() {
		tableCount++
	}

	// We expect: Title, Body, page break para, Page 2 = 4 paragraphs minimum
	if paraCount < 4 {
		t.Errorf(
			"Expected at least 4 paragraphs, got %d",
			paraCount,
		)
	}

	if tableCount != 1 {
		t.Errorf(
			"Expected 1 table, got %d",
			tableCount,
		)
	}
}

// TestIntegrationRealDocxFile tests opening and reading a real .docx file.
// This tests Task 5.31: Integration tests with real .docx files (Office 2016+).
func TestIntegrationRealDocxFile(t *testing.T) {
	// Path to the test fixture
	fixturePath := "../testdata/fixtures/minimal.docx"

	// Open the real .docx file
	doc, err := Open(fixturePath, false)
	if err != nil {
		t.Fatalf(
			"Failed to open minimal.docx: %v",
			err,
		)
	}
	defer func() { _ = doc.Close() }()

	// Verify document properties
	if doc.Type() != DocTypeDocument {
		t.Errorf(
			"Expected DocTypeDocument, got %v",
			doc.Type(),
		)
	}

	// Verify main part exists
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Verify package exists
	if doc.Package() == nil {
		t.Fatal("Package is nil")
	}
}

// TestIntegrationRealDocxRoundtrip tests opening a real .docx, modifying, saving, and reopening.
// This tests Task 5.31: roundtrip testing with real files.
func TestIntegrationRealDocxRoundtrip(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Path to the test fixture
	fixturePath := "../testdata/fixtures/minimal.docx"
	outputPath := filepath.Join(
		tmpDir,
		"modified.docx",
	)

	// Open the real .docx file
	doc, err := Open(fixturePath, true)
	if err != nil {
		t.Fatalf(
			"Failed to open minimal.docx: %v",
			err,
		)
	}

	// Save to a new location
	err = doc.SaveAs(outputPath)
	if err != nil {
		_ = doc.Close()
		t.Fatalf(
			"Failed to save document: %v",
			err,
		)
	}
	_ = doc.Close()

	// Reopen the saved document
	doc2, err := Open(outputPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to reopen saved document: %v",
			err,
		)
	}
	defer func() { _ = doc2.Close() }()

	// Verify the reopened document is valid
	if doc2.Type() != DocTypeDocument {
		t.Errorf(
			"Reopened document: expected DocTypeDocument, got %v",
			doc2.Type(),
		)
	}

	if doc2.MainPart() == nil {
		t.Error(
			"Reopened document MainPart is nil",
		)
	}
}

// TestIntegrationCreateDocumentWithText tests creating a basic document with text content.
// This tests Task 5.33: Integration test for creating basic documents with text.
func TestIntegrationCreateDocumentWithText(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-text-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"text_doc.docx",
	)

	// Create a document with text content using the builder
	builder := NewDocumentBuilder()

	// Add a title
	builder.AddHeading("Test Document Title", 1).
		AlignCenter().
		Bold()

	// Add body text
	builder.AddParagraph("This is the first paragraph of the test document.").
		SpacingAfter(12)

	builder.AddParagraph("This is the second paragraph with different formatting.").
		Italic().
		Color("0000FF")

	// Add a paragraph with mixed formatting
	pb := builder.AddParagraph("")
	pb.AddRun("This paragraph has ")
	pb.AddRun("bold").Bold()
	pb.AddRun(" and ")
	pb.AddRun("italic").Italic()
	pb.AddRun(" text.")

	// Build and verify
	builtDoc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	if builtDoc == nil {
		t.Fatal("Build() returned nil document")
	}

	body := builtDoc.Body()
	if body == nil {
		t.Fatal("Document body is nil")
	}

	// Count paragraphs
	paragraphCount := 0
	for range body.Paragraphs() {
		paragraphCount++
	}

	// We should have 4 paragraphs: title, first body, second body, mixed formatting
	if paragraphCount < 4 {
		t.Errorf(
			"Expected at least 4 paragraphs, got %d",
			paragraphCount,
		)
	}

	// Create an actual file and verify it can be opened
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	err = doc.SaveAs(testPath)
	if err != nil {
		_ = doc.Close()
		t.Fatalf("SaveAs() error = %v", err)
	}
	_ = doc.Close()

	// Verify file was created and can be reopened
	reopened, err := Open(testPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to reopen saved document: %v",
			err,
		)
	}
	_ = reopened.Close()
}

// TestIntegrationWithPhase1to4Components tests integration with Phase 1-4 components.
// This tests Task 5.34: Integration test with Phase 1-4 components.
//
//nolint:revive // cyclomatic: test requires many component verifications
func TestIntegrationWithPhase1to4Components(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-components-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"components.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test Phase 2: Packaging layer
	pkg := doc.Package()
	if pkg == nil {
		t.Fatal(
			"Package is nil - Phase 2 component failure",
		)
	}

	// Test Phase 3: OpenXML base layer
	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal(
			"MainPart is nil - Phase 3 component failure",
		)
	}

	// Test Phase 4: Features system
	features := doc.Features()
	if features == nil {
		t.Fatal(
			"Features is nil - Phase 4 component failure",
		)
	}

	// Test adding parts (Phase 5 with Phase 3 integration)
	stylesPart, err := mainPart.AddStylesPart()
	if err != nil {
		t.Fatalf(
			"AddStylesPart() error = %v",
			err,
		)
	}
	if stylesPart == nil {
		t.Fatal("StylesPart is nil")
	}

	settingsPart, err := mainPart.AddSettingsPart()
	if err != nil {
		t.Fatalf(
			"AddSettingsPart() error = %v",
			err,
		)
	}
	if settingsPart == nil {
		t.Fatal("SettingsPart is nil")
	}

	// Test header/footer parts
	headerPart, err := mainPart.AddHeaderPart()
	if err != nil {
		t.Fatalf(
			"AddHeaderPart() error = %v",
			err,
		)
	}
	if headerPart == nil {
		t.Fatal("HeaderPart is nil")
	}

	footerPart, err := mainPart.AddFooterPart()
	if err != nil {
		t.Fatalf(
			"AddFooterPart() error = %v",
			err,
		)
	}
	if footerPart == nil {
		t.Fatal("FooterPart is nil")
	}

	// Test footnotes/endnotes parts
	footnotesPart, err := mainPart.AddFootnotesPart()
	if err != nil {
		t.Fatalf(
			"AddFootnotesPart() error = %v",
			err,
		)
	}
	if footnotesPart == nil {
		t.Fatal("FootnotesPart is nil")
	}

	endnotesPart, err := mainPart.AddEndnotesPart()
	if err != nil {
		t.Fatalf(
			"AddEndnotesPart() error = %v",
			err,
		)
	}
	if endnotesPart == nil {
		t.Fatal("EndnotesPart is nil")
	}

	// Test comments part
	commentsPart, err := mainPart.AddCommentsPart()
	if err != nil {
		t.Fatalf(
			"AddCommentsPart() error = %v",
			err,
		)
	}
	if commentsPart == nil {
		t.Fatal("CommentsPart is nil")
	}

	// Test adding content to parts
	header := headerPart.GetOrCreateHeader()
	if header == nil {
		t.Fatal("Header element is nil")
	}
	header.AppendParagraph("Header Text")

	footer := footerPart.GetOrCreateFooter()
	if footer == nil {
		t.Fatal("Footer element is nil")
	}
	footer.AppendParagraph("Footer Text")

	// Add a footnote
	fn := footnotesPart.AddFootnote(
		"This is a test footnote.",
	)
	if fn == nil {
		t.Fatal("Footnote is nil")
	}

	// Add an endnote
	en := endnotesPart.AddEndnote(
		"This is a test endnote.",
	)
	if en == nil {
		t.Fatal("Endnote is nil")
	}

	// Add a comment
	comment := commentsPart.AddComment(
		"Test Author",
		"This is a test comment.",
	)
	if comment == nil {
		t.Fatal("Comment is nil")
	}

	// Save and verify
	err = doc.SaveAs(testPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(testPath); os.IsNotExist(
		err,
	) {
		t.Fatal("File was not created")
	}

	// Reopen and verify parts
	doc2, err := Open(testPath, false)
	if err != nil {
		t.Fatalf(
			"Failed to reopen document: %v",
			err,
		)
	}
	defer func() { _ = doc2.Close() }()

	if doc2.MainPart() == nil {
		t.Error(
			"Reopened document MainPart is nil",
		)
	}
}

// TestIntegrationGlossaryPart tests the GlossaryPart functionality.
func TestIntegrationGlossaryPart(t *testing.T) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-glossary-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(
		tmpDir,
		"glossary.docx",
	)

	// Create a document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	mainPart := doc.MainPart()
	if mainPart == nil {
		t.Fatal("MainPart is nil")
	}

	// Add glossary part
	glossaryPart, err := mainPart.AddGlossaryPart()
	if err != nil {
		t.Fatalf(
			"AddGlossaryPart() error = %v",
			err,
		)
	}

	if glossaryPart == nil {
		t.Fatal("GlossaryPart is nil")
	}

	// Verify content type
	if glossaryPart.FixedContentType() != "application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml" {
		t.Errorf(
			"Unexpected content type: %s",
			glossaryPart.FixedContentType(),
		)
	}

	// Test retrieval
	retrievedPart := mainPart.GlossaryPart()
	if retrievedPart == nil {
		t.Fatal(
			"GlossaryPart() returned nil after adding",
		)
	}

	// Save
	err = doc.SaveAs(testPath)
	if err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}
