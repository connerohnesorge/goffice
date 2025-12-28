package comparison

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ConvertPDFToPNG converts a PDF file to PNG images using Ghostscript.
// It returns the path to the generated PNG file (for single-page PDFs) or
// the pattern for multi-page PDFs (e.g., "output-%d.png").
//
// For multi-page PDFs, Ghostscript will create multiple files:
// output-1.png, output-2.png, etc.
//
// Requirements:
//   - Ghostscript (gs) must be installed and available in PATH
//
// Parameters:
//   - pdfPath: path to the input PDF file
//   - outputPath: path for the output PNG file(s)
//   - dpi: resolution for rendering (default: 150)
//
// Returns the output path and any error encountered.
func ConvertPDFToPNG(
	pdfPath, outputPath string,
	dpi int,
) (string, error) {
	// Check if Ghostscript is available
	gsPath, err := exec.LookPath("gs")
	if err != nil {
		return "", fmt.Errorf(
			"Ghostscript not found in PATH. Please install Ghostscript to enable visual comparison. Error: %w",
			err,
		)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf(
			"failed to create output directory: %w",
			err,
		)
	}

	// Set default DPI if not specified
	if dpi == 0 {
		dpi = 150
	}

	// Build Ghostscript command
	// -q: quiet mode
	// -dNOPAUSE: don't pause between pages
	// -dBATCH: exit after processing
	// -dSAFER: safer mode (disable file operations)
	// -sDEVICE=png16m: 24-bit RGB PNG output
	// -r{dpi}: resolution in DPI
	// -sOutputFile: output file pattern
	args := []string{
		"-q",
		"-dNOPAUSE",
		"-dBATCH",
		"-dSAFER",
		"-sDEVICE=png16m",
		fmt.Sprintf("-r%d", dpi),
		fmt.Sprintf(
			"-sOutputFile=%s",
			outputPath,
		),
		pdfPath,
	}

	// Execute Ghostscript
	cmd := exec.Command(gsPath, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"Ghostscript conversion failed: %w\nStderr: %s",
			err,
			stderr.String(),
		)
	}

	// Validate output file exists and has content
	stat, err := os.Stat(outputPath)
	if err != nil {
		return "", fmt.Errorf(
			"output PNG file not found: %w",
			err,
		)
	}

	if stat.Size() == 0 {
		return "", fmt.Errorf(
			"output PNG file is empty",
		)
	}

	return outputPath, nil
}

// ConvertPDFBytesToPNG converts PDF bytes to PNG, saving to the specified output path.
// This is a convenience wrapper around ConvertPDFToPNG for in-memory PDF data.
func ConvertPDFBytesToPNG(
	pdfBytes []byte,
	outputPath string,
	dpi int,
) (string, error) {
	// Create temporary PDF file
	tmpFile, err := os.CreateTemp(
		"",
		"goffice-pdf-*.pdf",
	)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create temporary PDF file: %w",
			err,
		)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Write PDF bytes to temp file
	if _, err := tmpFile.Write(pdfBytes); err != nil {
		return "", fmt.Errorf(
			"failed to write PDF to temporary file: %w",
			err,
		)
	}

	if err := tmpFile.Close(); err != nil {
		return "", fmt.Errorf(
			"failed to close temporary PDF file: %w",
			err,
		)
	}

	// Convert using the file-based function
	return ConvertPDFToPNG(
		tmpFile.Name(),
		outputPath,
		dpi,
	)
}

// IsGhostscriptAvailable checks if Ghostscript is available in the system PATH.
func IsGhostscriptAvailable() bool {
	_, err := exec.LookPath("gs")
	return err == nil
}
