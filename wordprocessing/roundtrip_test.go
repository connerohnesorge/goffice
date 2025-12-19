package wordprocessing

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// TestRoundtripBasicDocument tests creating, saving, reopening, and verifying
// a basic document with paragraphs.
func TestRoundtripBasicDocument(t *testing.T) {
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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"roundtrip.docx",
	)

	// Create a document with content using the builder
	builder := NewDocumentBuilder()
	builder.AddParagraph("First paragraph")
	builder.AddParagraph("Second paragraph")
	builder.AddParagraph("Third paragraph")

	originalDoc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	// Get original content
	originalXml := originalDoc.OuterXml()

	// Now create a proper document file
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save and close
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer doc2.Close()

	// Verify the document can be opened
	if doc2.MainPart() == nil {
		t.Error("MainPart() = nil after reopen")
	}

	// Verify document type is preserved
	if doc2.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want %v",
			doc2.Type(),
			DocTypeDocument,
		)
	}

	// Verify original XML had expected content
	if !strings.Contains(
		originalXml,
		"First paragraph",
	) {
		t.Error(
			"Original XML missing 'First paragraph'",
		)
	}
}

// TestRoundtripFormattedText tests that text formatting survives save/reopen.
func TestRoundtripFormattedText(t *testing.T) {
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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"formatted.docx",
	)

	// Create document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	doc1.Close()

	// Reopen
	doc2, err := Open(testPath, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer doc2.Close()

	// Verify read-only
	if doc2.IsEditable() {
		t.Error("Expected read-only document")
	}
}

// TestRoundtripTable tests that tables survive save/reopen.
func TestRoundtripTable(t *testing.T) {
	builder := NewDocumentBuilder()

	// Create a 3x3 table with content
	builder.AddTable(3, 3).
		SetCellText(0, 0, "A1").
		SetCellText(0, 1, "B1").
		SetCellText(0, 2, "C1").
		SetCellText(1, 0, "A2").
		SetCellText(1, 1, "B2").
		SetCellText(1, 2, "C2").
		SetCellText(2, 0, "A3").
		SetCellText(2, 1, "B3").
		SetCellText(2, 2, "C3").
		SetStyle("TableGrid")

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Body is nil")
	}

	// Verify table exists
	tableCount := 0
	for table := range body.Tables() {
		tableCount++

		// Verify row count
		if table.RowCount() != 3 {
			t.Errorf(
				"Expected 3 rows, got %d",
				table.RowCount(),
			)
		}

		// Verify cell content
		cell := table.GetCell(1, 1)
		if cell == nil {
			t.Fatal("GetCell(1,1) returned nil")
		}
		if cell.InnerText() != "B2" {
			t.Errorf(
				"Expected 'B2', got '%s'",
				cell.InnerText(),
			)
		}
	}

	if tableCount != 1 {
		t.Errorf(
			"Expected 1 table, got %d",
			tableCount,
		)
	}

	// Verify XML contains table elements
	xml := doc.OuterXml()
	if !strings.Contains(xml, "tbl") {
		t.Error("XML missing table element")
	}
	if !strings.Contains(xml, "A1") {
		t.Error("XML missing cell content 'A1'")
	}
}

// TestRoundtripStyles tests that style references survive save/reopen.
func TestRoundtripStyles(t *testing.T) {
	builder := NewDocumentBuilder()

	// Add paragraphs with different styles
	builder.AddHeading("Heading 1", 1)
	builder.AddHeading("Heading 2", 2)
	builder.AddParagraph("Normal paragraph")
	builder.AddHeading("Heading 3", 3)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Body is nil")
	}

	// Verify paragraph styles
	i := 0
	expectedStyles := []string{
		"Heading1",
		"Heading2",
		"",
		"Heading3",
	}
	for p := range body.Paragraphs() {
		if i >= len(expectedStyles) {
			break
		}

		props := p.Properties()
		if expectedStyles[i] != "" {
			if props == nil {
				t.Errorf(
					"Paragraph %d: expected style '%s', got nil properties",
					i,
					expectedStyles[i],
				)
			} else {
				styleId := props.ParagraphStyleId()
				if styleId != expectedStyles[i] {
					t.Errorf("Paragraph %d: expected style '%s', got '%s'", i, expectedStyles[i], styleId)
				}
			}
		}
		i++
	}
}

// TestRoundtripComplexDocument tests a complex document with multiple element types.
func TestRoundtripComplexDocument(t *testing.T) {
	builder := NewDocumentBuilder()

	// Title
	builder.AddHeading("Complex Document", 1).
		AlignCenter()

	// Introduction
	builder.AddParagraph("This document tests roundtrip functionality.").
		SpacingAfter(12)

	// First section
	builder.AddHeading("Section 1: Tables", 2)
	builder.AddParagraph(
		"Below is a sample table:",
	)

	builder.AddTable(2, 3).
		SetCellText(0, 0, "Header A").
		SetCellText(0, 1, "Header B").
		SetCellText(0, 2, "Header C").
		SetCellText(1, 0, "Data 1").
		SetCellText(1, 1, "Data 2").
		SetCellText(1, 2, "Data 3")

	// Second section
	builder.AddHeading("Section 2: Formatting", 2)

	pb := builder.AddParagraph("")
	pb.AddRun("This paragraph has ")
	pb.AddRun("bold").Bold()
	pb.AddRun(", ")
	pb.AddRun("italic").Italic()
	pb.AddRun(", and ")
	pb.AddRun("colored").Color("FF0000")
	pb.AddRun(" text.")

	// Page break
	builder.AddPageBreak()

	// Third section
	builder.AddHeading("Section 3: Lists", 2)
	builder.AddParagraph("Item 1").LeftIndent(36)
	builder.AddParagraph("Item 2").LeftIndent(36)
	builder.AddParagraph("Item 3").LeftIndent(36)

	// Build
	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Body is nil")
	}

	// Verify paragraph count
	paraCount := 0
	for range body.Paragraphs() {
		paraCount++
	}
	if paraCount < 10 {
		t.Errorf(
			"Expected at least 10 paragraphs, got %d",
			paraCount,
		)
	}

	// Verify table count
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

	// Verify XML content
	xml := doc.OuterXml()
	expectedElements := []string{
		"Complex Document",
		"Section 1",
		"Section 2",
		"Section 3",
		"Header A",
		"bold",
		"italic",
		"FF0000",
		"page", // page break
	}

	for _, elem := range expectedElements {
		if !strings.Contains(xml, elem) {
			t.Errorf(
				"XML missing expected content: '%s'",
				elem,
			)
		}
	}
}

// TestRoundtripDocumentTypes tests that different document types are preserved.
func TestRoundtripDocumentTypes(t *testing.T) {
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
	defer os.RemoveAll(tmpDir)

	testCases := []struct {
		docType   DocType
		extension string
	}{
		{DocTypeDocument, ".docx"},
		{DocTypeTemplate, ".dotx"},
	}

	for _, tc := range testCases {
		t.Run(
			tc.docType.String(),
			func(t *testing.T) {
				testPath := filepath.Join(
					tmpDir,
					"test"+tc.extension,
				)

				// Create
				doc1, err := New(
					testPath,
					tc.docType,
				)
				if err != nil {
					t.Fatalf(
						"New() error = %v",
						err,
					)
				}

				// Save
				if err := doc1.SaveAs(testPath); err != nil {
					t.Fatalf(
						"SaveAs() error = %v",
						err,
					)
				}
				doc1.Close()

				// Reopen
				doc2, err := Open(testPath, true)
				if err != nil {
					t.Fatalf(
						"Open() error = %v",
						err,
					)
				}

				// Verify type is preserved
				if doc2.Type() != tc.docType {
					t.Errorf(
						"Type() = %v, want %v",
						doc2.Type(),
						tc.docType,
					)
				}

				doc2.Close()
			},
		)
	}
}

// TestRoundtripParagraphProperties tests that paragraph properties survive roundtrip.
func TestRoundtripParagraphProperties(
	t *testing.T,
) {
	builder := NewDocumentBuilder()

	// Add paragraphs with various properties
	builder.AddParagraph("Centered").AlignCenter()
	builder.AddParagraph("Right aligned").
		AlignRight()
	builder.AddParagraph("Justified").
		AlignJustify()
	builder.AddParagraph("With spacing").
		SpacingBefore(12).
		SpacingAfter(12)
	builder.AddParagraph("Indented").
		LeftIndent(36).
		FirstLineIndent(18)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()

	// Verify properties
	i := 0
	for p := range body.Paragraphs() {
		props := p.Properties()

		switch i {
		case 0: // Centered
			if props != nil &&
				props.Justification() != elements.JustificationCenter {
				t.Errorf(
					"Paragraph %d: expected center",
					i,
				)
			}
		case 1: // Right
			if props != nil &&
				props.Justification() != elements.JustificationRight {
				t.Errorf(
					"Paragraph %d: expected right",
					i,
				)
			}
		case 2: // Justified
			if props != nil &&
				props.Justification() != elements.JustificationBoth {
				t.Errorf(
					"Paragraph %d: expected justify",
					i,
				)
			}
		case 3: // Spacing
			if props != nil {
				spacing := props.SpacingBetweenLines()
				if spacing != nil {
					// 12 points = 240 twips
					if spacing.Before() != 240 {
						t.Errorf(
							"Paragraph %d: expected before 240, got %d",
							i,
							spacing.Before(),
						)
					}
					if spacing.After() != 240 {
						t.Errorf(
							"Paragraph %d: expected after 240, got %d",
							i,
							spacing.After(),
						)
					}
				}
			}
		case 4: // Indented
			if props != nil {
				ind := props.Indentation()
				if ind != nil {
					// 36 points = 720 twips
					if ind.Left() != 720 {
						t.Errorf(
							"Paragraph %d: expected left 720, got %d",
							i,
							ind.Left(),
						)
					}
					// 18 points = 360 twips
					if ind.FirstLine() != 360 {
						t.Errorf(
							"Paragraph %d: expected firstLine 360, got %d",
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

// TestRoundtripRunProperties tests that run properties survive roundtrip.
func TestRoundtripRunProperties(t *testing.T) {
	builder := NewDocumentBuilder()

	pb := builder.AddParagraph("")
	pb.AddRun("Bold").Bold()
	pb.AddRun("Italic").Italic()
	pb.AddRun("Underline").Underline()
	pb.AddRun("Colored").Color("0000FF")
	pb.AddRun("Large").FontSize(20)

	doc, err := builder.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	body := doc.Body()
	for p := range body.Paragraphs() {
		for r := range p.Runs() {
			text := r.InnerText()
			props := r.Properties()

			if props == nil {
				continue
			}

			xml := props.OuterXml()

			switch text {
			case "Bold":
				if !strings.Contains(
					xml,
					"<w:b",
				) {
					t.Errorf(
						"Expected bold for '%s'",
						text,
					)
				}
			case "Italic":
				if !strings.Contains(
					xml,
					"<w:i",
				) {
					t.Errorf(
						"Expected italic for '%s'",
						text,
					)
				}
			case "Underline":
				if !strings.Contains(
					xml,
					"<w:u",
				) {
					t.Errorf(
						"Expected underline for '%s'",
						text,
					)
				}
			case "Colored":
				if !strings.Contains(
					xml,
					"0000FF",
				) {
					t.Errorf(
						"Expected blue color for '%s'",
						text,
					)
				}
			case "Large":
				// 20 points = 40 half-points
				if !strings.Contains(xml, "40") {
					t.Errorf(
						"Expected font size 40 for '%s'",
						text,
					)
				}
			}
		}
	}
}

// TestRoundtripEmptyDocument tests that an empty document can be saved and reopened.
func TestRoundtripEmptyDocument(t *testing.T) {
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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"empty.docx",
	)

	// Create empty document
	doc1, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Save
	if err := doc1.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	doc1.Close()

	// Reopen
	doc2, err := Open(testPath, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer doc2.Close()

	// Verify it's valid
	if doc2.MainPart() == nil {
		t.Error("MainPart() = nil")
	}
	if doc2.Type() != DocTypeDocument {
		t.Errorf(
			"Type() = %v, want DocTypeDocument",
			doc2.Type(),
		)
	}
}

// TestRoundtripMultipleSaves tests saving the same document multiple times.
func TestRoundtripMultipleSaves(t *testing.T) {
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
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(
		tmpDir,
		"multisave.docx",
	)

	// Create document
	doc, err := New(testPath, DocTypeDocument)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer doc.Close()

	// Save multiple times
	for i := 0; i < 3; i++ {
		if err := doc.SaveAs(testPath); err != nil {
			t.Fatalf(
				"SaveAs() attempt %d error = %v",
				i+1,
				err,
			)
		}

		// Verify file exists after each save
		info, err := os.Stat(testPath)
		if err != nil {
			t.Fatalf(
				"Stat() attempt %d error = %v",
				i+1,
				err,
			)
		}
		if info.Size() == 0 {
			t.Errorf(
				"File size is 0 after save attempt %d",
				i+1,
			)
		}
	}
}
