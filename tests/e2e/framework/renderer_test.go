package framework

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLibreOfficeRenderer_Render(
	t *testing.T,
) {
	// Create a temporary output directory
	tempDir, err := os.MkdirTemp(
		"",
		"renderer_test_*",
	)
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Use an existing test PPTX file from the repository
	testPPTX := filepath.Join(
		"..",
		"..",
		"..",
		"Open-XML-SDK",
		"samples",
		"Linq",
		"SvgExample",
		"Content",
		"EmptySlide.pptx",
	)

	// Check if the file exists
	if _, err := os.Stat(testPPTX); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test PPTX file not found: %s",
			testPPTX,
		)
	}

	// Create renderer
	renderer := NewLibreOfficeRenderer(tempDir)

	// Verify binaries can be found
	if _, err := renderer.findLibreOffice(); err != nil {
		t.Skipf(
			"LibreOffice not available: %v",
			err,
		)
	}
	if _, err := renderer.findPdftoppm(); err != nil {
		t.Skipf("pdftoppm not available: %v", err)
	}

	// Test rendering
	ctx := context.Background()
	pngPaths, err := renderer.Render(
		ctx,
		testPPTX,
	)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	// Verify PNG files were created
	if len(pngPaths) == 0 {
		t.Fatalf("No PNG files generated")
	}

	t.Logf(
		"Generated %d PNG files:",
		len(pngPaths),
	)
	for i, path := range pngPaths {
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf(
				"PNG file %d not found: %s",
				i,
				path,
			)

			continue
		}
		t.Logf(
			"  [%d] %s (size: %d bytes)",
			i+1,
			filepath.Base(path),
			info.Size(),
		)
	}

	// Verify each PNG file exists and has content
	for i, path := range pngPaths {
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf(
				"PNG file %d not accessible: %v",
				i,
				err,
			)

			continue
		}
		if info.Size() == 0 {
			t.Errorf(
				"PNG file %d is empty: %s",
				i,
				path,
			)
		}
	}
}

func TestLibreOfficeRenderer_findLibreOffice(
	t *testing.T,
) {
	renderer := NewLibreOfficeRenderer("")

	path, err := renderer.findLibreOffice()
	if err != nil {
		t.Skipf("LibreOffice not found: %v", err)
	}

	t.Logf("Found LibreOffice at: %s", path)

	// Verify the binary exists
	if _, err := os.Stat(path); err != nil {
		t.Errorf(
			"LibreOffice binary not accessible: %v",
			err,
		)
	}
}

func TestLibreOfficeRenderer_findPdftoppm(
	t *testing.T,
) {
	renderer := NewLibreOfficeRenderer("")

	path, err := renderer.findPdftoppm()
	if err != nil {
		t.Skipf("pdftoppm not found: %v", err)
	}

	t.Logf("Found pdftoppm at: %s", path)

	// Verify the binary exists
	if _, err := os.Stat(path); err != nil {
		t.Errorf(
			"pdftoppm binary not accessible: %v",
			err,
		)
	}
}

func TestLibreOfficeRenderer_convertToPDF(
	t *testing.T,
) {
	// Create a temporary output directory
	tempDir, err := os.MkdirTemp(
		"",
		"renderer_pdf_test_*",
	)
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Use an existing test PPTX file
	testPPTX := filepath.Join(
		"..",
		"..",
		"..",
		"Open-XML-SDK",
		"samples",
		"Linq",
		"SvgExample",
		"Content",
		"EmptySlide.pptx",
	)

	if _, err := os.Stat(testPPTX); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test PPTX file not found: %s",
			testPPTX,
		)
	}

	// Create renderer
	renderer := NewLibreOfficeRenderer(tempDir)

	// Verify LibreOffice is available and set the binary path
	bin, err := renderer.findLibreOffice()
	if err != nil {
		t.Skipf(
			"LibreOffice not available: %v",
			err,
		)
	}
	renderer.LibreOfficeBin = bin

	// Test PDF conversion
	ctx := context.Background()
	pdfPath, err := renderer.convertToPDF(
		ctx,
		testPPTX,
	)
	if err != nil {
		t.Fatalf("convertToPDF failed: %v", err)
	}

	// Verify PDF was created
	info, err := os.Stat(pdfPath)
	if err != nil {
		t.Fatalf("PDF file not found: %v", err)
	}

	if info.Size() == 0 {
		t.Fatalf("PDF file is empty")
	}

	t.Logf(
		"Created PDF: %s (size: %d bytes)",
		pdfPath,
		info.Size(),
	)
}

func TestLibreOfficeRenderer_convertPDFToPNG(
	t *testing.T,
) {
	// This test requires first creating a PDF, so we'll skip if we can't do that
	tempDir, err := os.MkdirTemp(
		"",
		"renderer_png_test_*",
	)
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testPPTX := filepath.Join(
		"..",
		"..",
		"..",
		"Open-XML-SDK",
		"samples",
		"Linq",
		"SvgExample",
		"Content",
		"EmptySlide.pptx",
	)

	if _, err := os.Stat(testPPTX); os.IsNotExist(
		err,
	) {
		t.Skipf(
			"Test PPTX file not found: %s",
			testPPTX,
		)
	}

	renderer := NewLibreOfficeRenderer(tempDir)

	// Verify tools are available and set the binary paths
	bin, err := renderer.findLibreOffice()
	if err != nil {
		t.Skipf(
			"LibreOffice not available: %v",
			err,
		)
	}
	renderer.LibreOfficeBin = bin

	pdfBin, err := renderer.findPdftoppm()
	if err != nil {
		t.Skipf("pdftoppm not available: %v", err)
	}
	renderer.PdftoppmBin = pdfBin

	// First create a PDF
	ctx := context.Background()
	pdfPath, err := renderer.convertToPDF(
		ctx,
		testPPTX,
	)
	if err != nil {
		t.Fatalf("convertToPDF failed: %v", err)
	}

	// Now test PNG conversion
	pngPaths, err := renderer.convertPDFToPNG(
		ctx,
		pdfPath,
	)
	if err != nil {
		t.Fatalf(
			"convertPDFToPNG failed: %v",
			err,
		)
	}

	if len(pngPaths) == 0 {
		t.Fatalf("No PNG files generated")
	}

	t.Logf(
		"Generated %d PNG files from PDF",
		len(pngPaths),
	)
	for i, path := range pngPaths {
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf(
				"PNG file %d not found: %s",
				i,
				path,
			)

			continue
		}
		if info.Size() == 0 {
			t.Errorf("PNG file %d is empty", i)
		}
		t.Logf(
			"  [%d] %s (size: %d bytes)",
			i+1,
			filepath.Base(path),
			info.Size(),
		)
	}
}
