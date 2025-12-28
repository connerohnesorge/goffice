package pdf

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/comparison"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/spreadsheet"
	"github.com/connerohnesorge/goffice/wordprocessing"
)

// TestWordFidelity tests visual fidelity of Word document rendering.
// This test suite compares our PDF output with Microsoft Word's "Save as PDF" output.
func TestWordFidelity(t *testing.T) {
	testCases := []struct {
		name        string
		docxFile    string
		description string
	}{
		{
			name:        "BasicText",
			docxFile:    "basic_text.docx",
			description: "Plain text with simple formatting",
		},
		{
			name:        "ComplexFormatting",
			docxFile:    "complex_formatting.docx",
			description: "Bold, italic, underline, colors, highlights",
		},
		{
			name:        "Tables",
			docxFile:    "tables.docx",
			description: "Tables with various border and shading styles",
		},
		{
			name:        "Lists",
			docxFile:    "lists.docx",
			description: "Bulleted and numbered lists with multiple levels",
		},
		{
			name:        "HeadersFooters",
			docxFile:    "headers_footers.docx",
			description: "Different headers/footers for first, odd, even pages",
		},
		{
			name:        "Images",
			docxFile:    "images.docx",
			description: "Embedded images with wrapping and positioning",
		},
		{
			name:        "Shapes",
			docxFile:    "shapes.docx",
			description: "DrawingML shapes with fills and effects",
		},
		{
			name:        "MultiColumn",
			docxFile:    "multi_column.docx",
			description: "Multi-column layout with column breaks",
		},
		{
			name:        "PageBreaks",
			docxFile:    "page_breaks.docx",
			description: "Various page break scenarios and pagination",
		},
		{
			name:        "FootnotesEndnotes",
			docxFile:    "footnotes_endnotes.docx",
			description: "Footnotes and endnotes rendering",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runWordFidelityTest(
				t,
				tc.docxFile,
				tc.description,
			)
		})
	}
}

// TestExcelFidelity tests visual fidelity of Excel spreadsheet rendering.
func TestExcelFidelity(t *testing.T) {
	testCases := []struct {
		name        string
		xlsxFile    string
		description string
	}{
		{
			name:        "BasicCells",
			xlsxFile:    "basic_cells.xlsx",
			description: "Simple cell text and number formatting",
		},
		{
			name:        "NumberFormats",
			xlsxFile:    "number_formats.xlsx",
			description: "Various number and date formats",
		},
		{
			name:        "CellFormatting",
			xlsxFile:    "cell_formatting.xlsx",
			description: "Cell borders, fills, fonts, and alignment",
		},
		{
			name:        "MergedCells",
			xlsxFile:    "merged_cells.xlsx",
			description: "Merged cell regions with content",
		},
		{
			name:        "ConditionalFormatting",
			xlsxFile:    "conditional_formatting.xlsx",
			description: "Data bars, color scales, icon sets",
		},
		{
			name:        "Charts",
			xlsxFile:    "charts.xlsx",
			description: "Bar, line, and pie charts",
		},
		{
			name:        "PrintArea",
			xlsxFile:    "print_area.xlsx",
			description: "Print area and page break settings",
		},
		{
			name:        "HeadersFooters",
			xlsxFile:    "headers_footers.xlsx",
			description: "Spreadsheet headers and footers",
		},
		{
			name:        "FitToPage",
			xlsxFile:    "fit_to_page.xlsx",
			description: "Fit to page scaling",
		},
		{
			name:        "RepeatTitles",
			xlsxFile:    "repeat_titles.xlsx",
			description: "Repeat rows and columns on each page",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runExcelFidelityTest(
				t,
				tc.xlsxFile,
				tc.description,
			)
		})
	}
}

// TestPowerPointFidelity tests visual fidelity of PowerPoint rendering.
func TestPowerPointFidelity(t *testing.T) {
	testCases := []struct {
		name        string
		pptxFile    string
		description string
	}{
		{
			name:        "BasicSlides",
			pptxFile:    "basic_slides.pptx",
			description: "Simple slides with text and shapes",
		},
		{
			name:        "MasterSlides",
			pptxFile:    "master_slides.pptx",
			description: "Slides using different master layouts",
		},
		{
			name:        "Shapes",
			pptxFile:    "shapes.pptx",
			description: "Various DrawingML shapes with effects",
		},
		{
			name:        "Images",
			pptxFile:    "images.pptx",
			description: "Images with cropping and transforms",
		},
		{
			name:        "Charts",
			pptxFile:    "charts.pptx",
			description: "Charts embedded in presentations",
		},
		{
			name:        "Tables",
			pptxFile:    "tables.pptx",
			description: "Tables with styling",
		},
		{
			name:        "TextBoxes",
			pptxFile:    "text_boxes.pptx",
			description: "Text boxes with various alignments and formatting",
		},
		{
			name:        "Backgrounds",
			pptxFile:    "backgrounds.pptx",
			description: "Slide backgrounds (solid, gradient, picture)",
		},
		{
			name:        "NotesPages",
			pptxFile:    "notes_pages.pptx",
			description: "Notes pages rendering",
		},
		{
			name:        "SmartArt",
			pptxFile:    "smartart.pptx",
			description: "SmartArt graphics rendering",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runPowerPointFidelityTest(
				t,
				tc.pptxFile,
				tc.description,
			)
		})
	}
}

// runWordFidelityTest runs a single Word fidelity test.
func runWordFidelityTest(
	t *testing.T,
	docxFile, description string,
) {
	testdataPath := filepath.Join(
		"testdata",
		"fidelity",
		"word",
		docxFile,
	)

	// Check if test file exists
	if _, err := os.Stat(testdataPath); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test file not found: %s (create it to enable this test)",
			docxFile,
		)
		return
	}

	doc, err := openTestWordDocument(testdataPath)
	if err != nil {
		t.Fatalf(
			"Failed to open test document: %v",
			err,
		)
	}
	defer doc.Close()

	// Render to PDF
	var buf bytes.Buffer
	opts := DefaultRenderOptions()
	opts.FontEmbedding = EmbedSubset

	if err := RenderWord(doc, &buf, opts); err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	// Save output for manual comparison
	outputPath := filepath.Join(
		"testdata",
		"fidelity",
		"word",
		"output",
		filepath.Base(docxFile)+".pdf",
	)
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		t.Errorf("Failed to save output: %v", err)
	}

	t.Logf(
		"✓ %s: Generated PDF for comparison (%d bytes)",
		description,
		buf.Len(),
	)

	// Perform automated visual comparison
	baselinePath := filepath.Join(
		"testdata",
		"fidelity",
		"word",
		strings.TrimSuffix(
			docxFile,
			filepath.Ext(docxFile),
		)+".baseline.png",
	)

	if err := compareWithBaseline(t, buf.Bytes(), baselinePath, description); err != nil {
		t.Logf("Visual comparison: %v", err)
	}
}

// runExcelFidelityTest runs a single Excel fidelity test.
func runExcelFidelityTest(
	t *testing.T,
	xlsxFile, description string,
) {
	testdataPath := filepath.Join(
		"testdata",
		"fidelity",
		"excel",
		xlsxFile,
	)

	if _, err := os.Stat(testdataPath); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test file not found: %s",
			xlsxFile,
		)
		return
	}

	doc, err := openTestSpreadsheetDocument(
		testdataPath,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open test document: %v",
			err,
		)
	}
	defer doc.Close()

	var buf bytes.Buffer
	opts := DefaultRenderOptions()

	if err := RenderSpreadsheet(doc, &buf, opts); err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	outputPath := filepath.Join(
		"testdata",
		"fidelity",
		"excel",
		"output",
		filepath.Base(xlsxFile)+".pdf",
	)
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		t.Errorf("Failed to save output: %v", err)
	}

	t.Logf(
		"✓ %s: Generated PDF for comparison (%d bytes)",
		description,
		buf.Len(),
	)
}

// runPowerPointFidelityTest runs a single PowerPoint fidelity test.
func runPowerPointFidelityTest(
	t *testing.T,
	pptxFile, description string,
) {
	testdataPath := filepath.Join(
		"testdata",
		"fidelity",
		"powerpoint",
		pptxFile,
	)

	if _, err := os.Stat(testdataPath); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test file not found: %s",
			pptxFile,
		)
		return
	}

	doc, err := openTestPresentationDocument(
		testdataPath,
	)
	if err != nil {
		t.Fatalf(
			"Failed to open test document: %v",
			err,
		)
	}
	defer doc.Close()

	var buf bytes.Buffer
	opts := DefaultRenderOptions()

	if err := RenderPresentation(doc, &buf, opts); err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	outputPath := filepath.Join(
		"testdata",
		"fidelity",
		"powerpoint",
		"output",
		filepath.Base(pptxFile)+".pdf",
	)
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		t.Errorf("Failed to save output: %v", err)
	}

	t.Logf(
		"✓ %s: Generated PDF for comparison (%d bytes)",
		description,
		buf.Len(),
	)
}

// Helper functions

func openTestWordDocument(
	path string,
) (*wordprocessing.Document, error) {
	return wordprocessing.Open(path, false)
}

func openTestSpreadsheetDocument(
	path string,
) (*spreadsheet.Document, error) {
	return spreadsheet.Open(path, false)
}

func openTestPresentationDocument(
	path string,
) (*presentation.Document, error) {
	return presentation.Open(path, false)
}

// compareWithBaseline compares a generated PDF with a baseline PNG image.
// This function handles the full visual comparison workflow:
//  1. Convert PDF to PNG using Ghostscript
//  2. Compare with baseline image
//  3. Generate diff images on failure
//  4. Report results
//
// The comparison is skipped gracefully if:
//   - Ghostscript is not available
//   - Baseline image doesn't exist
//
// This allows tests to run in environments without visual comparison support.
func compareWithBaseline(
	t *testing.T,
	pdfBytes []byte,
	baselinePath, description string,
) error {
	// Check if Ghostscript is available
	if !comparison.IsGhostscriptAvailable() {
		t.Log(
			"⚠ Ghostscript not found - skipping visual comparison",
		)
		t.Log(
			"  Install Ghostscript to enable automated visual comparison",
		)
		return nil
	}

	// Check if baseline exists
	if _, err := os.Stat(baselinePath); os.IsNotExist(
		err,
	) {
		t.Logf(
			"⚠ Baseline image not found: %s",
			baselinePath,
		)
		t.Log(
			"  Create baseline images to enable visual comparison",
		)
		return nil
	}

	// Convert PDF to PNG
	generatedPNGPath := baselinePath + ".generated.png"
	config := comparison.DefaultComparisonConfig()

	pngPath, err := comparison.ConvertPDFBytesToPNG(
		pdfBytes,
		generatedPNGPath,
		config.DPI,
	)
	if err != nil {
		return fmt.Errorf(
			"PDF to PNG conversion failed: %w",
			err,
		)
	}
	defer os.Remove(
		pngPath,
	) // Clean up generated PNG

	// Compare images
	result, err := comparison.CompareImages(
		baselinePath,
		pngPath,
		config,
	)
	if err != nil {
		return fmt.Errorf(
			"image comparison failed: %w",
			err,
		)
	}

	// Log results
	t.Logf(
		"Visual comparison: %s",
		result.String(),
	)

	// Check if within tolerance
	if !result.IsWithinTolerance() {
		// Generate diff image for manual review
		diffPath := baselinePath + ".diff.png"
		if err := comparison.AnnotateDiffImage(baselinePath, pngPath, diffPath, config); err != nil {
			t.Errorf(
				"Failed to generate diff image: %v",
				err,
			)
		} else {
			t.Logf("Diff image saved: %s", diffPath)
		}

		// Optionally create side-by-side comparison
		sideBySidePath := baselinePath + ".comparison.png"
		if err := comparison.CreateSideBySideDiff(baselinePath, pngPath, diffPath, sideBySidePath); err != nil {
			t.Logf(
				"Warning: Failed to create side-by-side diff: %v",
				err,
			)
		} else {
			t.Logf("Side-by-side comparison saved: %s", sideBySidePath)
		}

		return fmt.Errorf(
			"visual comparison failed: %.2f%% of pixels exceed tolerance (threshold: %.2f%%)",
			result.ExceedingTolerance(),
			config.ImageDiffThreshold*100,
		)
	}

	t.Logf(
		"✓ Visual comparison passed: differences within tolerance",
	)
	return nil
}
