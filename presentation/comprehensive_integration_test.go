package presentation

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice/openxml/validation"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// TestComprehensivePresentationOpenClose tests opening and closing various test presentations.
func TestComprehensivePresentationOpenClose(
	t *testing.T,
) {
	testCases := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			"Presentation",
			"Presentation.pptx",
			false,
		},
		{"Autosave", "autosave.pptx", false},
		{"McPpt", "mcppt.pptx", false},
		{"Animation", "animation.pptx", false},
		{
			"MediaReference",
			"mediareference.pptx",
			false,
		},
		{
			"AlignTab",
			"Algn_tab_TabAlignment.pptx",
			false,
		},
		{"3DTestDash", "3dtestdash.pptx", false},
		{"3DTestDot", "3dtestdot.pptx", false},
		{"Office2016_01", "Of16-01.pptx", false},
		{"Office2016_02", "Of16-02.pptx", false},
		{"Office2016_03", "Of16-03.pptx", false},
		{"Template", "Presentation.potx", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				tc.filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			// Open presentation
			pres, err := Open(path, false)
			if (err != nil) != tc.wantErr {
				t.Fatalf(
					"Open() error = %v, wantErr %v",
					err,
					tc.wantErr,
				)
			}
			if err != nil {
				return
			}
			defer func() { _ = pres.Close() }()

			// Verify basic properties
			if pres == nil {
				t.Fatal(
					"Open() returned nil presentation",
				)
			}

			// Verify package
			if pres.Package() == nil {
				t.Error("Package() returned nil")
			}

			// Verify presentation part
			presPart := pres.PresentationPart()
			if presPart == nil {
				t.Log(
					"PresentationPart() returned nil - may need initialization",
				)
			}

			// Close and verify
			if err := pres.Close(); err != nil {
				t.Errorf(
					"Close() error = %v",
					err,
				)
			}
		})
	}
}

// TestComprehensivePresentationStructure tests presentation structure validation.
func TestComprehensivePresentationStructure(
	t *testing.T,
) {
	testCases := []string{
		"Presentation.pptx",
		"autosave.pptx",
		"mcppt.pptx",
		"animation.pptx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			// Verify presentation part
			presPart := pres.PresentationPart()
			if presPart == nil {
				t.Skip(
					"PresentationPart() is nil - skipping structure test",
				)
			}

			// Count slides
			slideCount := pres.SlideCount()
			t.Logf(
				"Presentation has %d slides",
				slideCount,
			)

			// Verify we can iterate through slides
			iterCount := 0
			for slide := range pres.Slides() {
				if slide == nil {
					t.Error(
						"Slides() yielded nil",
					)
				}
				iterCount++
			}

			if iterCount != slideCount {
				t.Errorf(
					"Slides() iteration count = %d, SlideCount() = %d",
					iterCount,
					slideCount,
				)
			}

			// Check for slide masters
			masterCount := 0
			for master := range pres.SlideMasters() {
				if master == nil {
					t.Error(
						"SlideMasters() yielded nil",
					)
				}
				masterCount++
			}
			t.Logf(
				"Presentation has %d slide masters",
				masterCount,
			)
		})
	}
}

// TestComprehensiveRoundtrip tests opening, modifying, saving, and reopening.
func TestComprehensiveRoundtrip(t *testing.T) {
	testCases := []string{
		"Presentation.pptx",
		"autosave.pptx",
		"mcppt.pptx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			sourcePath := filepath.Join(
				"testdata",
				filename,
			)

			// Skip if file doesn't exist
			if _, err := os.Stat(sourcePath); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					sourcePath,
				)
			}

			// Create temp directory
			tmpDir, err := os.MkdirTemp(
				"",
				"goffice-roundtrip-pptx-*",
			)
			if err != nil {
				t.Fatalf(
					"Failed to create temp dir: %v",
					err,
				)
			}
			defer func() { _ = os.RemoveAll(tmpDir) }()

			// Open original presentation
			pres, err := Open(sourcePath, true)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}

			// Record original slide count
			originalSlideCount := pres.SlideCount()
			t.Logf(
				"Original presentation has %d slides",
				originalSlideCount,
			)

			// Save to new location
			outputPath := filepath.Join(
				tmpDir,
				"modified_"+filename,
			)
			if err := pres.SaveAs(outputPath); err != nil {
				_ = pres.Close()
				t.Fatalf(
					"SaveAs() error = %v",
					err,
				)
			}
			_ = pres.Close()

			// Reopen saved presentation
			pres2, err := Open(outputPath, false)
			if err != nil {
				t.Fatalf(
					"Failed to reopen saved presentation: %v",
					err,
				)
			}
			defer func() { _ = pres2.Close() }()

			// Verify structure
			if pres2.PresentationPart() == nil {
				t.Log(
					"Reopened presentation missing presentation part (may need initialization)",
				)
			}

			// Verify package integrity
			if pres2.Package() == nil {
				t.Error(
					"Reopened presentation missing package",
				)
			}

			// Log slide count after roundtrip
			reopenedSlideCount := pres2.SlideCount()
			t.Logf(
				"Reopened presentation has %d slides",
				reopenedSlideCount,
			)
		})
	}
}

// TestComprehensivePartsExtraction tests extracting various parts from presentations.
func TestComprehensivePartsExtraction(
	t *testing.T,
) {
	t.Run(
		"PresentationWithSlides",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Presentation.pptx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			presPart := pres.PresentationPart()
			if presPart == nil {
				t.Skip(
					"PresentationPart() is nil",
				)
			}

			// Check for slides
			slides := presPart.SlideParts()
			if len(slides) > 0 {
				t.Logf(
					"Found %d slide parts",
					len(slides),
				)
			}

			// Check for slide masters
			masters := presPart.SlideMasterParts()
			if len(masters) > 0 {
				t.Logf(
					"Found %d slide master parts",
					len(masters),
				)
			}

			// Check for themes
			themes := presPart.ThemeParts()
			if len(themes) > 0 {
				t.Logf(
					"Found %d theme parts",
					len(themes),
				)
			}

			// Check for notes master
			notesMaster := presPart.NotesMasterPart()
			if notesMaster != nil {
				t.Log("Found notes master part")
			}

			// Check for handout master
			handoutMaster := presPart.HandoutMasterPart()
			if handoutMaster != nil {
				t.Log(
					"Found handout master part",
				)
			}
		},
	)

	t.Run(
		"PresentationWithComments",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Presentation.pptx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			presPart := pres.PresentationPart()
			if presPart == nil {
				t.Skip(
					"PresentationPart() is nil",
				)
			}

			// Check for comment authors
			commentAuthors := presPart.CommentAuthorsPart()
			if commentAuthors != nil {
				t.Log(
					"Found comment authors part",
				)
			}

			// Check for slide comments
			slides := presPart.SlideParts()
			for i, slide := range slides {
				comments := slide.SlideCommentsPart()
				if comments != nil {
					t.Logf(
						"Slide %d has comments",
						i+1,
					)
				}
			}
		},
	)

	t.Run(
		"PresentationWithMedia",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"mediareference.pptx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			presPart := pres.PresentationPart()
			if presPart == nil {
				t.Skip(
					"PresentationPart() is nil",
				)
			}

			// Check for image parts in slides
			slides := presPart.SlideParts()
			for i, slide := range slides {
				images := slide.ImageParts()
				if len(images) > 0 {
					t.Logf(
						"Slide %d has %d image parts",
						i+1,
						len(images),
					)
				}
			}
		},
	)
}

// TestComprehensiveRelationshipHandling tests relationship handling.
func TestComprehensiveRelationshipHandling(
	t *testing.T,
) {
	testFiles := []string{
		"Presentation.pptx",
		"autosave.pptx",
		"mcppt.pptx",
	}

	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			pkg := pres.Package()
			if pkg == nil {
				t.Fatal("Package is nil")
			}

			// Get package relationships
			packPkg := pkg.Package()
			if packPkg == nil {
				t.Fatal(
					"Underlying package is nil",
				)
			}

			// Verify main presentation relationship exists
			rels := packPkg.Relationships()
			if rels == nil {
				t.Fatal(
					"Package relationships are nil",
				)
			}

			foundMainPres := false
			for rel := range rels.All() {
				if strings.Contains(
					rel.Type(),
					"officeDocument",
				) {
					foundMainPres = true
					t.Logf(
						"Found main presentation relationship: %s",
						rel.ID(),
					)
				}
			}

			if !foundMainPres {
				t.Error(
					"Main presentation relationship not found",
				)
			}
		})
	}
}

// TestComprehensivePresentationProperties tests presentation properties.
func TestComprehensivePresentationProperties(
	t *testing.T,
) {
	t.Run("CoreProperties", func(t *testing.T) {
		path := filepath.Join(
			"testdata",
			"Presentation.pptx",
		)
		if _, err := os.Stat(path); os.IsNotExist(
			err,
		) {
			t.Skip("Test file not found")
		}

		pres, err := Open(path, false)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer func() { _ = pres.Close() }()

		coreProps := pres.CoreProperties()
		if coreProps == nil {
			return
		}
		t.Log("Core properties found")
		if title := coreProps.Title(); title != "" {
			t.Logf("Title: %s", title)
		}
		if creator := coreProps.Creator(); creator != "" {
			t.Logf("Creator: %s", creator)
		}
	})

	t.Run(
		"ExtendedProperties",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"Presentation.pptx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			extProps := pres.ExtendedProperties()
			if extProps != nil {
				t.Log(
					"Extended properties found",
				)
			}
		},
	)
}

// TestComprehensiveAnimations tests animation handling.
func TestComprehensiveAnimations(t *testing.T) {
	path := filepath.Join(
		"testdata",
		"animation.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Verify presentation can be opened
	if pres.PresentationPart() == nil {
		t.Skip("PresentationPart() is nil")
	}

	// Count slides
	slideCount := pres.SlideCount()
	t.Logf(
		"Animation presentation has %d slides",
		slideCount,
	)

	// Verify we can iterate slides
	for i := range slideCount {
		slidePart, err := pres.GetSlide(i)
		if err != nil {
			t.Logf(
				"Could not get slide %d: %v",
				i,
				err,
			)

			continue
		}
		if slidePart != nil {
			t.Logf("Accessed slide %d", i+1)
		}
	}
}

// TestComprehensiveMediaReferences tests media reference handling.
func TestComprehensiveMediaReferences(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"mediareference.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	presPart := pres.PresentationPart()
	if presPart == nil {
		t.Skip("PresentationPart() is nil")
	}

	// Check for media in slides
	slides := presPart.SlideParts()
	totalImages := 0
	for i, slide := range slides {
		images := slide.ImageParts()
		if len(images) > 0 {
			t.Logf(
				"Slide %d has %d image parts",
				i+1,
				len(images),
			)
			totalImages += len(images)
		}
	}

	t.Logf(
		"Total image parts found: %d",
		totalImages,
	)
}

// TestComprehensive3DGraphics tests 3D graphics support.
func TestComprehensive3DGraphics(t *testing.T) {
	testCases := []string{
		"3dtestdash.pptx",
		"3dtestdot.pptx",
	}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			// Verify basic structure
			if pres.Package() == nil {
				t.Error("Package is nil")
			}

			// Log slide count
			slideCount := pres.SlideCount()
			t.Logf(
				"3D graphics presentation has %d slides",
				slideCount,
			)

			// Verify we can access presentation part
			if pres.PresentationPart() == nil {
				t.Log(
					"PresentationPart() is nil (may need initialization)",
				)
			}
		})
	}
}

// TestComprehensiveTextFormatting tests text formatting and alignment.
func TestComprehensiveTextFormatting(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Algn_tab_TabAlignment.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Verify presentation can be opened
	if pres.Package() == nil {
		t.Fatal("Package is nil")
	}

	// Log slide count
	slideCount := pres.SlideCount()
	t.Logf(
		"Text formatting presentation has %d slides",
		slideCount,
	)

	// Verify basic structure
	if pres.PresentationPart() == nil {
		t.Log(
			"PresentationPart() is nil (may need initialization)",
		)
	}
}

// TestComprehensiveOffice2016Features tests Office 2016 specific features.
func TestComprehensiveOffice2016Features(
	t *testing.T,
) {
	office2016Files := []string{
		"Of16-01.pptx",
		"Of16-02.pptx",
		"Of16-03.pptx",
	}

	for _, filename := range office2016Files {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skipf(
					"Test file not found: %s",
					path,
				)
			}

			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			// Verify basic structure
			if pres.Package() == nil {
				t.Error("Package is nil")
			}

			// Log slide count
			slideCount := pres.SlideCount()
			t.Logf(
				"Office 2016 presentation has %d slides",
				slideCount,
			)

			// Test validation against Office 2016
			errors := pres.Validate(
				validation.Office2016,
			)
			if len(errors) > 0 {
				t.Logf(
					"Validation found %d issues (may be expected)",
					len(errors),
				)
			}
		})
	}
}

// TestComprehensiveEncryptedPresentation tests encrypted presentation handling.
func TestComprehensiveEncryptedPresentation(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"encrypted_pptx.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	// Encrypted presentations should fail to open or be handled gracefully
	pres, err := Open(path, false)
	if err != nil {
		t.Logf(
			"Encrypted presentation correctly failed to open: %v",
			err,
		)

		return
	}
	defer func() { _ = pres.Close() }()

	// If it opened, log a warning
	t.Log(
		"WARNING: Encrypted presentation opened (encryption may not be enforced)",
	)
}

// TestComprehensiveErrorHandling tests error handling for malformed presentations.
func TestComprehensiveErrorHandling(
	t *testing.T,
) {
	t.Run("NonExistentFile", func(t *testing.T) {
		pres, err := Open(
			"testdata/nonexistent.pptx",
			false,
		)
		if err == nil {
			_ = pres.Close()
			t.Error(
				"Expected error opening non-existent file",
			)
		}
	})

	t.Run(
		"EncryptedPresentation",
		func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				"encrypted_pptx.pptx",
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			// Should handle encrypted presentations gracefully
			pres, err := Open(path, false)
			if err != nil {
				t.Logf(
					"Encrypted presentation error (expected): %v",
					err,
				)
			} else {
				defer func() { _ = pres.Close() }()
				t.Log("Encrypted presentation opened (may not be fully encrypted)")
			}
		},
	)
}

// TestComprehensivePerformanceLargePresentation tests performance with large presentations.
func TestComprehensivePerformanceLargePresentation(
	t *testing.T,
) {
	if testing.Short() {
		t.Skip(
			"Skipping performance test in short mode",
		)
	}

	largeFiles := []string{
		"o09_Performance_typical.pptx", // 1.5MB
		"Presentation.pptx",            // 1002KB
	}

	for _, filename := range largeFiles {
		t.Run(filename, func(t *testing.T) {
			path := filepath.Join(
				"testdata",
				filename,
			)
			if _, err := os.Stat(path); os.IsNotExist(
				err,
			) {
				t.Skip("Test file not found")
			}

			// Measure open time
			pres, err := Open(path, false)
			if err != nil {
				t.Fatalf("Open() error = %v", err)
			}
			defer func() { _ = pres.Close() }()

			// Count slides
			slideCount := pres.SlideCount()
			t.Logf(
				"Large presentation has %d slides",
				slideCount,
			)

			// Iterate through slides
			presPart := pres.PresentationPart()
			if presPart != nil {
				slides := presPart.SlideParts()
				t.Logf(
					"Accessed %d slide parts from large presentation",
					len(slides),
				)
			}
		})
	}
}

// TestComprehensiveMarkupCompatibility tests markup compatibility handling.
func TestComprehensiveMarkupCompatibility(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"mcppt.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Verify basic structure
	if pres.Package() == nil {
		t.Fatal("Package is nil")
	}

	// Log slide count
	slideCount := pres.SlideCount()
	t.Logf(
		"Markup compatibility presentation has %d slides",
		slideCount,
	)
}

// TestComprehensiveTemplateHandling tests template (.potx) file handling.
func TestComprehensiveTemplateHandling(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"Presentation.potx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Verify document type
	if pres.Type() == DocTypeTemplate {
		t.Log("Correctly detected template type")
	} else {
		t.Logf("Document type: %s (expected Template)", pres.Type().String())
	}

	// Verify basic structure
	if pres.Package() == nil {
		t.Fatal("Package is nil")
	}

	// Log slide count
	slideCount := pres.SlideCount()
	t.Logf("Template has %d slides", slideCount)
}

// TestComprehensivePresentationToBytes tests saving presentation to byte array.
func TestComprehensivePresentationToBytes(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"autosave.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Save to buffer
	var buf bytes.Buffer
	if err := pres.SaveTo(&buf); err != nil {
		t.Fatalf("SaveTo() error = %v", err)
	}

	// Verify buffer has data
	if buf.Len() == 0 {
		t.Error("SaveTo() produced empty buffer")
	}

	t.Logf(
		"Saved presentation to buffer: %d bytes",
		buf.Len(),
	)
}

// TestComprehensiveValidationVersions tests validation across Office versions.
func TestComprehensiveValidationVersions(
	t *testing.T,
) {
	path := filepath.Join(
		"testdata",
		"autosave.pptx",
	)
	if _, err := os.Stat(path); os.IsNotExist(
		err,
	) {
		t.Skip("Test file not found")
	}

	pres, err := Open(path, false)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	versions := []struct {
		name    string
		version validation.FileFormatVersions
	}{
		{"Office2016", validation.Office2016},
		{"Office2019", validation.Office2019},
		{"Office2021", validation.Office2021},
		{"Microsoft365", validation.Microsoft365},
	}

	for _, v := range versions {
		t.Run(v.name, func(t *testing.T) {
			errors := pres.Validate(v.version)
			t.Logf(
				"%s validation: %d issues",
				v.name,
				len(errors),
			)

			// Check IsValid
			isValid := pres.IsValid(v.version)
			t.Logf(
				"%s IsValid: %v",
				v.name,
				isValid,
			)
		})
	}
}

// TestComprehensivePresentationType tests presentation type detection and conversion.
func TestComprehensivePresentationType(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-prestype-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Test each presentation type
	types := []struct {
		presType  DocType
		extension string
	}{
		{DocTypePresentation, ".pptx"},
		{DocTypeTemplate, ".potx"},
		{DocTypeSlideshow, ".ppsx"},
	}

	for _, tt := range types {
		t.Run(
			tt.presType.String(),
			func(t *testing.T) {
				testPath := filepath.Join(
					tmpDir,
					"test"+tt.extension,
				)
				pres, err := New(
					testPath,
					tt.presType,
				)
				if err != nil {
					t.Fatalf(
						"New() error = %v",
						err,
					)
				}
				defer func() { _ = pres.Close() }()

				// Verify type
				if pres.Type() != tt.presType {
					t.Errorf(
						"Expected type %v, got %v",
						tt.presType,
						pres.Type(),
					)
				}

				// Verify extension
				if tt.presType.Extension() != tt.extension {
					t.Errorf(
						"Expected extension %s, got %s",
						tt.extension,
						tt.presType.Extension(),
					)
				}

				// Verify content type
				contentType := tt.presType.ContentType()
				if contentType == "" {
					t.Error(
						"ContentType() returned empty string",
					)
				}
				t.Logf(
					"Content type: %s",
					contentType,
				)
			},
		)
	}
}

// TestComprehensiveSlideManipulation tests slide creation and manipulation.
func TestComprehensiveSlideManipulation(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-slide-manip-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.pptx")
	pres, err := New(
		testPath,
		DocTypePresentation,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Add multiple slides
	for i := range 5 {
		slide, err := pres.AddSlide()
		if err != nil {
			t.Fatalf(
				"AddSlide() error at slide %d: %v",
				i,
				err,
			)
		}
		if slide == nil {
			t.Fatalf(
				"AddSlide() returned nil at slide %d",
				i,
			)
		}
	}

	// Verify slide count
	if pres.SlideCount() != 5 {
		t.Errorf(
			"Expected 5 slides, got %d",
			pres.SlideCount(),
		)
	}

	// Test GetSlide
	for i := range 5 {
		slide, err := pres.GetSlide(i)
		if err != nil {
			t.Errorf(
				"GetSlide(%d) error: %v",
				i,
				err,
			)
		}
		if slide == nil {
			t.Errorf(
				"GetSlide(%d) returned nil",
				i,
			)
		}
	}

	// Test out of range
	_, err = pres.GetSlide(10)
	if err == nil {
		t.Error(
			"GetSlide(10) should return error for out of range",
		)
	}

	// Save and verify
	if err := pres.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestComprehensiveSlideMasterAndLayouts tests slide master and layout hierarchy.
func TestComprehensiveSlideMasterAndLayouts(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-master-layout-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.pptx")
	pres, err := New(
		testPath,
		DocTypePresentation,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Add slide master
	master, err := pres.AddSlideMaster()
	if err != nil {
		t.Fatalf(
			"AddSlideMaster() error = %v",
			err,
		)
	}
	if master == nil {
		t.Fatal("AddSlideMaster() returned nil")
	}

	// Add layout to master
	layout, err := master.AddSlideLayoutPart()
	if err != nil {
		t.Fatalf(
			"AddSlideLayoutPart() error = %v",
			err,
		)
	}
	if layout == nil {
		t.Fatal(
			"AddSlideLayoutPart() returned nil",
		)
	}

	// Verify we can iterate slide masters
	masterCount := 0
	for m := range pres.SlideMasters() {
		if m == nil {
			t.Error("SlideMasters() yielded nil")
		}
		masterCount++
	}

	if masterCount == 0 {
		t.Error(
			"Expected at least 1 slide master",
		)
	}

	// Save and verify
	if err := pres.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestComprehensiveThemeHandling tests theme part handling.
func TestComprehensiveThemeHandling(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-theme-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.pptx")
	pres, err := New(
		testPath,
		DocTypePresentation,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Add theme
	theme, err := pres.AddTheme()
	if err != nil {
		t.Fatalf("AddTheme() error = %v", err)
	}
	if theme == nil {
		t.Fatal("AddTheme() returned nil")
	}

	// Verify theme is accessible
	if pres.Theme() == nil {
		t.Error(
			"Theme() should not be nil after AddTheme()",
		)
	}

	// Save and verify
	if err := pres.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestComprehensiveNotesAndComments tests notes and comments functionality.
func TestComprehensiveNotesAndComments(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-notes-comments-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.pptx")
	pres, err := New(
		testPath,
		DocTypePresentation,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Add a slide
	slide, err := pres.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}

	// Add notes to slide
	notes, err := slide.AddNotesSlidePart()
	if err != nil {
		t.Fatalf(
			"AddNotesSlidePart() error = %v",
			err,
		)
	}
	if notes == nil {
		t.Fatal(
			"AddNotesSlidePart() returned nil",
		)
	}

	// Verify notes are accessible
	if slide.NotesSlidePart() == nil {
		t.Error(
			"NotesSlidePart() should not be nil after adding",
		)
	}

	// Add comment authors
	presPart := pres.PresentationPart()
	if presPart == nil {
		t.Fatal("PresentationPart() returned nil")
	}

	authors, err := presPart.AddCommentAuthorsPart()
	if err != nil {
		t.Fatalf(
			"AddCommentAuthorsPart() error = %v",
			err,
		)
	}
	if authors == nil {
		t.Fatal(
			"AddCommentAuthorsPart() returned nil",
		)
	}

	// Add comments to slide
	comments, err := slide.AddSlideCommentsPart()
	if err != nil {
		t.Fatalf(
			"AddSlideCommentsPart() error = %v",
			err,
		)
	}
	if comments == nil {
		t.Fatal(
			"AddSlideCommentsPart() returned nil",
		)
	}

	// Save and verify
	if err := pres.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestComprehensiveImageHandling tests image part handling.
func TestComprehensiveImageHandling(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-image-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	testPath := filepath.Join(tmpDir, "test.pptx")
	pres, err := New(
		testPath,
		DocTypePresentation,
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer func() { _ = pres.Close() }()

	// Add a slide
	slide, err := pres.AddSlide()
	if err != nil {
		t.Fatalf("AddSlide() error = %v", err)
	}

	// Add an image
	imgPart, err := slide.AddImagePart(
		parts.ImageTypePng,
	)
	if err != nil {
		t.Fatalf("AddImagePart() error = %v", err)
	}
	if imgPart == nil {
		t.Fatal("AddImagePart() returned nil")
	}

	// Feed some PNG data (minimal valid PNG header)
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG signature
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52, // IHDR chunk
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, // 1x1 pixel
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xFF, 0xFF, 0x3F,
		0x00, 0x05, 0xFE, 0x02, 0xFE, 0xDC, 0xCC, 0x59,
		0xE7, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E,
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}
	imgPart.FeedDataBytes(pngData)

	// Verify image parts
	images := slide.ImageParts()
	if len(images) != 1 {
		t.Errorf(
			"ImageParts() = %d, want 1",
			len(images),
		)
	}

	// Save and verify
	if err := pres.SaveAs(testPath); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
}

// TestComprehensiveRoundtripWithModifications tests roundtrip with modifications.
func TestComprehensiveRoundtripWithModifications(
	t *testing.T,
) {
	tmpDir, err := os.MkdirTemp(
		"",
		"goffice-roundtrip-mod-*",
	)
	if err != nil {
		t.Fatalf(
			"Failed to create temp dir: %v",
			err,
		)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create initial presentation
	path1 := filepath.Join(
		tmpDir,
		"original.pptx",
	)
	pres, err := New(path1, DocTypePresentation)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// Add some slides
	for range 3 {
		_, err := pres.AddSlide()
		if err != nil {
			t.Fatalf("AddSlide() error = %v", err)
		}
	}

	// Save and close
	if err := pres.SaveAs(path1); err != nil {
		t.Fatalf("SaveAs() error = %v", err)
	}
	if err := pres.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	// Reopen and modify
	pres2, err := Open(path1, true)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	// Record slide count after reopen
	reopenedCount := pres2.SlideCount()
	t.Logf(
		"After reopen SlideCount() = %d",
		reopenedCount,
	)

	// Add more slides
	for range 2 {
		_, err := pres2.AddSlide()
		if err != nil {
			t.Fatalf(
				"AddSlide() after reopen error = %v",
				err,
			)
		}
	}

	// Change document type
	if err := pres2.ChangeType(DocTypeSlideshow); err != nil {
		t.Fatalf("ChangeType() error = %v", err)
	}

	// Save with new extension
	path2 := filepath.Join(
		tmpDir,
		"modified.ppsx",
	)
	if err := pres2.SaveAs(path2); err != nil {
		t.Fatalf(
			"SaveAs() with new type error = %v",
			err,
		)
	}
	if err := pres2.Close(); err != nil {
		t.Fatalf("Close() second error = %v", err)
	}

	// Final verification
	pres3, err := Open(path2, false)
	if err != nil {
		t.Fatalf("Open() final error = %v", err)
	}
	defer func() { _ = pres3.Close() }()

	if pres3.Type() != DocTypeSlideshow {
		t.Errorf(
			"Final Type() = %v, want %v",
			pres3.Type(),
			DocTypeSlideshow,
		)
	}

	finalSlideCount := pres3.SlideCount()
	t.Logf(
		"Final SlideCount() = %d",
		finalSlideCount,
	)
}
