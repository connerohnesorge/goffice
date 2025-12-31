package framework

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// FrameworkConfig holds global framework configuration
type FrameworkConfig struct {
	// Paths
	OutputDir   string `json:"output_dir"`   // Base output directory
	ReportsDir  string `json:"reports_dir"`  // Reports directory
	FixturesDir string `json:"fixtures_dir"` // Test fixtures directory

	// Generator paths
	GoGeneratorPath     string `json:"go_generator_path"`     // Path to Go generator binary
	CSharpGeneratorPath string `json:"csharp_generator_path"` // Path to C# generator binary

	// Renderer configuration
	LibreOfficePath string `json:"libreoffice_path"` // Path to LibreOffice binary
	PowerPointPath  string `json:"powerpoint_path"`  // Path to PowerPoint (Windows only)
	PdfToPpmPath    string `json:"pdftoppm_path"`    // Path to pdftoppm (poppler)

	// Image comparison
	ImageMagickComparePath string `json:"imagemagick_compare_path"` // Path to ImageMagick compare

	// Default test configuration
	DefaultTestConfig TestConfig `json:"default_test_config"`

	// Concurrency
	MaxParallelTests int `json:"max_parallel_tests"` // Maximum number of tests to run in parallel

	// Logging
	Verbose  bool   `json:"verbose"`   // Enable verbose logging
	LogFile  string `json:"log_file"`  // Path to log file (empty = stdout)
	LogLevel string `json:"log_level"` // debug, info, warn, error
}

// DefaultFrameworkConfig returns default framework configuration
func DefaultFrameworkConfig() *FrameworkConfig {
	return &FrameworkConfig{
		OutputDir:   "output",
		ReportsDir:  "reports",
		FixturesDir: "fixtures",

		GoGeneratorPath:     "generators/go/e2e-go-generator",
		CSharpGeneratorPath: "generators/csharp/bin/Release/net9.0/e2e-csharp-generator",

		LibreOfficePath: findLibreOffice(),
		PowerPointPath:  findPowerPoint(),
		PdfToPpmPath:    findPdfToPpm(),

		ImageMagickComparePath: findImageMagickCompare(),

		DefaultTestConfig: DefaultTestConfig(),

		MaxParallelTests: 4,

		Verbose:  false,
		LogFile:  "",
		LogLevel: "info",
	}
}

// LoadConfigFromFile loads framework configuration from JSON file
func LoadConfigFromFile(
	path string,
) (*FrameworkConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read config file: %w",
			err,
		)
	}

	var config FrameworkConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf(
			"failed to parse config file: %w",
			err,
		)
	}

	return &config, nil
}

// SaveConfigToFile saves framework configuration to JSON file
func (c *FrameworkConfig) SaveConfigToFile(
	path string,
) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf(
			"failed to marshal config: %w",
			err,
		)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf(
			"failed to write config file: %w",
			err,
		)
	}

	return nil
}

// Validate checks that all required paths exist
func (c *FrameworkConfig) Validate() error {
	// Check output directories exist or can be created
	for _, dir := range []string{c.OutputDir, c.ReportsDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf(
				"failed to create directory %s: %w",
				dir,
				err,
			)
		}
	}

	// Check fixtures directory exists
	if _, err := os.Stat(c.FixturesDir); os.IsNotExist(
		err,
	) {
		return fmt.Errorf(
			"fixtures directory does not exist: %s",
			c.FixturesDir,
		)
	}

	// Check renderer is available
	if c.DefaultTestConfig.RenderBackend == RenderBackendLibreOffice {
		if c.LibreOfficePath == "" {
			return fmt.Errorf(
				"LibreOffice not found, please install it or set libreoffice_path in config",
			)
		}
		if _, err := os.Stat(c.LibreOfficePath); os.IsNotExist(
			err,
		) {
			return fmt.Errorf(
				"LibreOffice binary not found at: %s",
				c.LibreOfficePath,
			)
		}
	}

	return nil
}

// Helper functions to find tools in PATH or common locations

func findLibreOffice() string {
	// Try common locations
	candidates := []string{
		"/usr/bin/libreoffice",
		"/usr/local/bin/libreoffice",
		"/Applications/LibreOffice.app/Contents/MacOS/soffice",
		"C:\\Program Files\\LibreOffice\\program\\soffice.exe",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Try PATH
	if path, err := findInPath("libreoffice"); err == nil {
		return path
	}
	if path, err := findInPath("soffice"); err == nil {
		return path
	}

	return ""
}

func findPowerPoint() string {
	// Windows only
	candidates := []string{
		"C:\\Program Files\\Microsoft Office\\root\\Office16\\POWERPNT.EXE",
		"C:\\Program Files (x86)\\Microsoft Office\\root\\Office16\\POWERPNT.EXE",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

func findPdfToPpm() string {
	// Try common locations
	candidates := []string{
		"/usr/bin/pdftoppm",
		"/usr/local/bin/pdftoppm",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Try PATH
	if path, err := findInPath("pdftoppm"); err == nil {
		return path
	}

	return ""
}

func findImageMagickCompare() string {
	// Try common locations
	candidates := []string{
		"/usr/bin/compare",
		"/usr/local/bin/compare",
		"/opt/homebrew/bin/compare",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Try PATH
	if path, err := findInPath("compare"); err == nil {
		return path
	}

	return ""
}

func findInPath(name string) (string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH not set")
	}

	for _, dir := range filepath.SplitList(path) {
		fullPath := filepath.Join(dir, name)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}

	return "", fmt.Errorf(
		"%s not found in PATH",
		name,
	)
}

// GetAbsOutputPath returns absolute path for output subdirectory
func (c *FrameworkConfig) GetAbsOutputPath(
	subdir string,
) string {
	return filepath.Join(c.OutputDir, subdir)
}

// GetAbsReportPath returns absolute path for report subdirectory
func (c *FrameworkConfig) GetAbsReportPath(
	subdir string,
) string {
	return filepath.Join(c.ReportsDir, subdir)
}

// GetAbsFixturePath returns absolute path for fixture subdirectory
func (c *FrameworkConfig) GetAbsFixturePath(
	subdir string,
) string {
	return filepath.Join(c.FixturesDir, subdir)
}
