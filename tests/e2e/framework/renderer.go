package framework

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"
)

// Renderer converts PPTX files to PNG images for visual comparison.
// Implementations include LibreOfficeRenderer (via headless soffice) and others.
type Renderer interface {
	// Render converts a PPTX file to a list of PNG images (one per slide).
	// Returns the paths to the generated PNG files in slide order.
	// The context can be used to cancel long-running renders or enforce timeouts.
	Render(
		ctx context.Context,
		pptxPath string,
	) ([]string, error)
}

// LibreOfficeRenderer uses LibreOffice headless mode to render PPTX files.
// It performs a two-step conversion:
//  1. PPTX to PDF using `soffice --headless --convert-to pdf`
//  2. PDF to PNG using `pdftoppm` from poppler-utils
//
// This provides high-fidelity rendering that closely matches Microsoft Office output.
type LibreOfficeRenderer struct {
	// DPI for PNG output (default: 150)
	DPI int

	// OutputDir is where temporary and final files are stored
	OutputDir string

	// Timeout for LibreOffice conversion (default: 30 seconds)
	Timeout time.Duration

	// LibreOfficeBin is the path to the soffice binary
	// If empty, it will be auto-detected
	LibreOfficeBin string

	// PdftoppmBin is the path to the pdftoppm binary
	// If empty, it will be auto-detected
	PdftoppmBin string
}

// NewLibreOfficeRenderer creates a new renderer with default settings.
// Default DPI is 150, and timeout is 30 seconds per conversion operation.
// Binary paths for soffice and pdftoppm are auto-detected from common locations.
func NewLibreOfficeRenderer(
	outputDir string,
) *LibreOfficeRenderer {
	return &LibreOfficeRenderer{
		DPI:       150,
		OutputDir: outputDir,
		Timeout:   30 * time.Second,
	}
}

// Render converts a PPTX file to PNG images
func (r *LibreOfficeRenderer) Render(
	ctx context.Context,
	pptxPath string,
) ([]string, error) {
	// Ensure binaries are located
	if r.LibreOfficeBin == "" {
		bin, err := r.findLibreOffice()
		if err != nil {
			return nil, fmt.Errorf(
				"locate LibreOffice: %w",
				err,
			)
		}
		r.LibreOfficeBin = bin
	}

	if r.PdftoppmBin == "" {
		bin, err := r.findPdftoppm()
		if err != nil {
			return nil, fmt.Errorf(
				"locate pdftoppm: %w",
				err,
			)
		}
		r.PdftoppmBin = bin
	}

	// Ensure output directory exists
	if err := os.MkdirAll(r.OutputDir, 0o755); err != nil {
		return nil, fmt.Errorf(
			"create output directory: %w",
			err,
		)
	}

	// Step 1: Convert PPTX to PDF
	pdfPath, err := r.convertToPDF(ctx, pptxPath)
	if err != nil {
		return nil, fmt.Errorf(
			"convert to PDF: %w",
			err,
		)
	}

	// Step 2: Convert PDF to PNG images
	pngPaths, err := r.convertPDFToPNG(
		ctx,
		pdfPath,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"convert PDF to PNG: %w",
			err,
		)
	}

	return pngPaths, nil
}

// convertToPDF converts a PPTX file to PDF using LibreOffice headless
func (r *LibreOfficeRenderer) convertToPDF(
	ctx context.Context,
	pptxPath string,
) (string, error) {
	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(
		ctx,
		r.Timeout,
	)
	defer cancel()

	// Determine output PDF path
	baseName := filepath.Base(pptxPath)
	nameWithoutExt := baseName[:len(baseName)-len(filepath.Ext(baseName))]
	pdfPath := filepath.Join(
		r.OutputDir,
		nameWithoutExt+".pdf",
	)

	// Remove existing PDF if present (LibreOffice may not overwrite)
	if _, err := os.Stat(pdfPath); err == nil {
		if err := os.Remove(pdfPath); err != nil {
			return "", fmt.Errorf(
				"remove existing PDF: %w",
				err,
			)
		}
	}

	// Build LibreOffice command
	// soffice --headless --convert-to pdf --outdir <outputDir> <pptxPath>
	cmd := exec.CommandContext(
		timeoutCtx,
		r.LibreOfficeBin,
		"--headless",
		"--convert-to",
		"pdf",
		"--outdir",
		r.OutputDir,
		pptxPath,
	)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(
			"LibreOffice conversion failed: %w\nOutput: %s",
			err,
			string(output),
		)
	}

	// Verify PDF was created
	if _, err := os.Stat(pdfPath); os.IsNotExist(
		err,
	) {
		return "", fmt.Errorf(
			"PDF not created at expected path: %s\nOutput: %s",
			pdfPath,
			string(output),
		)
	}

	return pdfPath, nil
}

// convertPDFToPNG converts a PDF file to PNG images using pdftoppm
func (r *LibreOfficeRenderer) convertPDFToPNG(
	ctx context.Context,
	pdfPath string,
) ([]string, error) {
	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(
		ctx,
		r.Timeout,
	)
	defer cancel()

	// Determine output prefix
	baseName := filepath.Base(pdfPath)
	nameWithoutExt := baseName[:len(baseName)-len(filepath.Ext(baseName))]
	outputPrefix := filepath.Join(
		r.OutputDir,
		nameWithoutExt,
	)

	// Build pdftoppm command
	// pdftoppm -png -r <dpi> <pdfPath> <outputPrefix>
	cmd := exec.CommandContext(
		timeoutCtx,
		r.PdftoppmBin,
		"-png",
		"-r",
		fmt.Sprintf("%d", r.DPI),
		pdfPath,
		outputPrefix,
	)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf(
			"pdftoppm conversion failed: %w\nOutput: %s",
			err,
			string(output),
		)
	}

	// Collect generated PNG files
	// pdftoppm outputs: prefix-1.png, prefix-2.png, prefix-3.png, ...
	pattern := fmt.Sprintf(
		"%s-*.png",
		outputPrefix,
	)
	pngPaths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf(
			"glob PNG files: %w",
			err,
		)
	}

	if len(pngPaths) == 0 {
		return nil, fmt.Errorf(
			"no PNG files generated from PDF: %s\nOutput: %s",
			pdfPath,
			string(output),
		)
	}

	// Sort paths to ensure consistent ordering (prefix-1.png, prefix-2.png, ...)
	sort.Strings(pngPaths)

	return pngPaths, nil
}

// findLibreOffice locates the LibreOffice soffice binary
func (r *LibreOfficeRenderer) findLibreOffice() (string, error) {
	// Try common paths based on OS
	var candidates []string

	switch runtime.GOOS {
	case "linux":
		candidates = []string{
			"/usr/bin/libreoffice",
			"/usr/bin/soffice",
			"/usr/local/bin/libreoffice",
			"/usr/local/bin/soffice",
			"/opt/libreoffice/program/soffice",
		}
	case "darwin": // macOS
		candidates = []string{
			"/Applications/LibreOffice.app/Contents/MacOS/soffice",
			"/usr/local/bin/soffice",
		}
	case "windows":
		candidates = []string{
			"C:\\Program Files\\LibreOffice\\program\\soffice.exe",
			"C:\\Program Files (x86)\\LibreOffice\\program\\soffice.exe",
		}
	}

	// Check each candidate path
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	// Fallback to PATH lookup
	path, err := exec.LookPath("soffice")
	if err == nil {
		return path, nil
	}

	// Also try "libreoffice" command
	path, err = exec.LookPath("libreoffice")
	if err == nil {
		return path, nil
	}

	return "", fmt.Errorf(
		"LibreOffice not found: please install LibreOffice or set LibreOfficeBin explicitly",
	)
}

// findPdftoppm locates the pdftoppm binary (from poppler-utils)
func (r *LibreOfficeRenderer) findPdftoppm() (string, error) {
	// Try common paths based on OS
	var candidates []string

	switch runtime.GOOS {
	case "linux":
		candidates = []string{
			"/usr/bin/pdftoppm",
			"/usr/local/bin/pdftoppm",
		}
	case "darwin": // macOS
		candidates = []string{
			"/usr/local/bin/pdftoppm",
			"/opt/homebrew/bin/pdftoppm",
		}
	case "windows":
		candidates = []string{
			"C:\\Program Files\\poppler\\bin\\pdftoppm.exe",
			"C:\\Program Files (x86)\\poppler\\bin\\pdftoppm.exe",
		}
	}

	// Check each candidate path
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	// Fallback to PATH lookup
	path, err := exec.LookPath("pdftoppm")
	if err == nil {
		return path, nil
	}

	return "", fmt.Errorf(
		"pdftoppm not found: please install poppler-utils or set PdftoppmBin explicitly",
	)
}
