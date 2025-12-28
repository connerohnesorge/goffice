package pdf

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/drawing"
	"github.com/connerohnesorge/goffice-pdf/font"
	"github.com/connerohnesorge/goffice-pdf/layout"
	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/spreadsheet"
	"github.com/connerohnesorge/goffice/wordprocessing"
)

// BenchmarkWordRendering benchmarks Word document rendering performance.
func BenchmarkWordRendering(b *testing.B) {
	// Load test document
	doc, err := loadTestWordDocument(b)
	if err != nil {
		b.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()
	opts.FontEmbedding = EmbedSubset

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := RenderWord(doc, &buf, opts); err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

// BenchmarkWordRenderingLarge benchmarks rendering of a large Word document.
func BenchmarkWordRenderingLarge(b *testing.B) {
	doc, err := loadTestWordDocument(b)
	if err != nil {
		b.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()
	opts.FontEmbedding = EmbedSubset
	opts.CompressContent = true

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := RenderWord(doc, &buf, opts); err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

// BenchmarkSpreadsheetRendering benchmarks Excel rendering performance.
func BenchmarkSpreadsheetRendering(b *testing.B) {
	doc, err := loadTestSpreadsheetDocument(b)
	if err != nil {
		b.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := RenderSpreadsheet(doc, &buf, opts); err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

// BenchmarkPresentationRendering benchmarks PowerPoint rendering performance.
func BenchmarkPresentationRendering(
	b *testing.B,
) {
	doc, err := loadTestPresentationDocument(b)
	if err != nil {
		b.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := RenderPresentation(doc, &buf, opts); err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

// BenchmarkFontLoading benchmarks font loading and caching.
func BenchmarkFontLoading(b *testing.B) {
	// Create a new font cache for testing
	cache := font.NewFontCache(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Benchmark font cache operations: Get/Put cycles
		// This tests both cache hits and misses
		f := &font.Font{
			Family: "Calibri",
			Style:  0, // Regular
		}
		cache.Put(f)

		// Try to retrieve from cache (should be cached now)
		_, ok := cache.Get("Calibri", 0)
		if !ok {
			b.Fatal(
				"Font not found in cache after Put",
			)
		}
	}
}

// BenchmarkFontLoadingFile benchmarks loading fonts from disk.
func BenchmarkFontLoadingFile(b *testing.B) {
	// Try to find a system font to load
	var fontPath string

	// Look for common system fonts (in order of preference)
	paths := []string{
		// Nix/NixOS fonts
		"/nix/store/dg8mhwckl2wl1b92hswj0pxsfzg8lv23-corefonts-1/share/fonts/truetype/Times_New_Roman_Bold.ttf",
		"/nix/store/85z7s2w9ps2ln74cchlq1g3b4gdqad5l-vista-fonts-1/share/fonts/truetype/constani.ttf",
		// Linux standard locations
		"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
		"/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
		// macOS
		"/System/Library/Fonts/Helvetica.ttc",
		// Windows
		"C:\\Windows\\Fonts\\arial.ttf",
		// Home directory (user fonts)
		filepath.Join(
			os.Getenv("HOME"),
			".local/share/fonts/Meslo LG S DZ Italic for Powerline.ttf",
		),
	}

	for _, path := range paths {
		if path != "" {
			if _, err := os.Stat(path); err == nil {
				fontPath = path
				break
			}
		}
	}

	if fontPath == "" {
		b.Skip(
			"No system fonts found for benchmark",
		)
		return
	}

	cache := font.NewFontCache(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Benchmark loading the same font from disk multiple times
		// The cache should serve the second load from memory
		f, err := cache.LoadFile(fontPath)
		if err != nil {
			b.Fatalf(
				"Failed to load font: %v",
				err,
			)
		}
		if f == nil {
			b.Fatal("Font is nil")
		}
	}
}

// BenchmarkFontCacheHitMiss benchmarks cache hit vs miss performance.
func BenchmarkFontCacheHitMiss(b *testing.B) {
	cache := font.NewFontCache(100)

	// Pre-populate cache with some fonts
	for i := 0; i < 50; i++ {
		f := &font.Font{
			Family: fmt.Sprintf("Font%d", i),
			Style:  0,
		}
		cache.Put(f)
	}

	b.ResetTimer()
	b.Run(
		"CacheHit",
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// These should all be cache hits
				_, ok := cache.Get("Font25", 0)
				if !ok {
					b.Fatal(
						"Cache hit expected but missed",
					)
				}
			}
		},
	)

	b.Run(
		"CacheMiss",
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// These should all be cache misses
				_, ok := cache.Get(
					fmt.Sprintf(
						"NonExistent%d",
						i,
					),
					0,
				)
				if ok {
					b.Fatal(
						"Cache miss expected but hit",
					)
				}
			}
		},
	)
}

// BenchmarkFontCacheLRU benchmarks LRU eviction performance.
func BenchmarkFontCacheLRU(b *testing.B) {
	// Create a small cache to force evictions
	cache := font.NewFontCache(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Add more fonts than cache can hold
		// This forces LRU evictions
		for j := 0; j < 20; j++ {
			f := &font.Font{
				Family: fmt.Sprintf("Font%d", j),
				Style:  0,
			}
			cache.Put(f)
		}
	}

	// Verify cache respects size limit
	if cache.Size() > cache.MaxSize() {
		b.Fatalf(
			"Cache size %d exceeds max %d",
			cache.Size(),
			cache.MaxSize(),
		)
	}
}

// BenchmarkImageDecoding benchmarks image decoding and embedding.
func BenchmarkImageDecoding(b *testing.B) {
	// Create sample image data for benchmarking
	// We'll use different formats to benchmark various decoders

	b.Run("JPEG", func(b *testing.B) {
		// Create a simple JPEG image
		img := createTestImage(100, 100)
		var jpegBuf bytes.Buffer
		if err := jpeg.Encode(&jpegBuf, img, &jpeg.Options{Quality: 90}); err != nil {
			b.Fatalf(
				"Failed to create JPEG: %v",
				err,
			)
		}
		jpegData := jpegBuf.Bytes()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := drawing.LoadJPEG(
				jpegData,
				fmt.Sprintf("Img%d", i),
			)
			if err != nil {
				b.Fatalf(
					"Failed to load JPEG: %v",
					err,
				)
			}
		}
	})

	b.Run("PNG", func(b *testing.B) {
		// Create a PNG image
		img := createTestImage(100, 100)
		var pngBuf bytes.Buffer
		if err := png.Encode(&pngBuf, img); err != nil {
			b.Fatalf(
				"Failed to create PNG: %v",
				err,
			)
		}
		pngData := pngBuf.Bytes()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := drawing.LoadPNG(
				pngData,
				fmt.Sprintf("Img%d", i),
			)
			if err != nil {
				b.Fatalf(
					"Failed to load PNG: %v",
					err,
				)
			}
		}
	})

	b.Run("PNG_WithAlpha", func(b *testing.B) {
		// Create a PNG image with alpha channel
		img := createTestImageWithAlpha(100, 100)
		var pngBuf bytes.Buffer
		if err := png.Encode(&pngBuf, img); err != nil {
			b.Fatalf(
				"Failed to create PNG: %v",
				err,
			)
		}
		pngData := pngBuf.Bytes()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := drawing.LoadPNG(
				pngData,
				fmt.Sprintf("Img%d", i),
			)
			if err != nil {
				b.Fatalf(
					"Failed to load PNG: %v",
					err,
				)
			}
		}
	})
}

// BenchmarkTextLayout benchmarks text layout engine.
func BenchmarkTextLayout(b *testing.B) {
	// Create a font cache and layout engine
	cache := font.NewFontCache(100)

	// Try to load a system font for realistic benchmarking
	var fontPath string
	paths := []string{
		// Nix/NixOS fonts
		"/nix/store/dg8mhwckl2wl1b92hswj0pxsfzg8lv23-corefonts-1/share/fonts/truetype/Times_New_Roman_Bold.ttf",
		"/nix/store/85z7s2w9ps2ln74cchlq1g3b4gdqad5l-vista-fonts-1/share/fonts/truetype/constani.ttf",
		// Linux standard locations
		"/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
		"/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf",
		// macOS
		"/System/Library/Fonts/Helvetica.ttc",
		// Windows
		"C:\\Windows\\Fonts\\arial.ttf",
		// Home directory (user fonts)
		filepath.Join(
			os.Getenv("HOME"),
			".local/share/fonts/Meslo LG S DZ Italic for Powerline.ttf",
		),
	}

	for _, path := range paths {
		if path != "" {
			if _, err := os.Stat(path); err == nil {
				fontPath = path
				break
			}
		}
	}

	if fontPath == "" {
		b.Skip(
			"No system fonts found for benchmark",
		)
		return
	}

	testFont, err := cache.LoadFile(fontPath)
	if err != nil {
		b.Skipf("Failed to load font: %v", err)
		return
	}

	engine := layout.NewTextLayoutEngine(cache)

	b.Run("ShortParagraph", func(b *testing.B) {
		// Benchmark laying out a short paragraph
		text := "The quick brown fox jumps over the lazy dog."
		para := createTestParagraph(
			text,
			testFont,
			12.0,
		)

		opts := layout.DefaultLayoutOptions()
		opts.MaxWidth = 400.0

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lines := engine.LayoutParagraph(
				para,
				opts,
			)
			if len(lines) == 0 {
				b.Fatal("No lines generated")
			}
		}
	})

	b.Run("LongParagraph", func(b *testing.B) {
		// Benchmark laying out a longer paragraph with line breaking
		text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
		para := createTestParagraph(
			text,
			testFont,
			12.0,
		)

		opts := layout.DefaultLayoutOptions()
		opts.MaxWidth = 400.0

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lines := engine.LayoutParagraph(
				para,
				opts,
			)
			if len(lines) == 0 {
				b.Fatal("No lines generated")
			}
		}
	})

	b.Run("JustifiedText", func(b *testing.B) {
		// Benchmark justified text layout
		text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
		para := createTestParagraph(
			text,
			testFont,
			12.0,
		)
		para.Properties.Alignment = layout.AlignJustify

		opts := layout.DefaultLayoutOptions()
		opts.MaxWidth = 400.0
		opts.Justify = true

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lines := engine.LayoutParagraph(
				para,
				opts,
			)
			if len(lines) == 0 {
				b.Fatal("No lines generated")
			}
		}
	})

	b.Run(
		"MultipleFontSizes",
		func(b *testing.B) {
			// Benchmark paragraph with multiple runs at different sizes
			para := layout.NewParagraph()

			runs := []*layout.TextRun{
				createTestRun(
					"Large text ",
					testFont,
					16.0,
				),
				createTestRun(
					"normal text ",
					testFont,
					12.0,
				),
				createTestRun(
					"small text",
					testFont,
					8.0,
				),
			}

			for _, run := range runs {
				para.AddRun(run)
			}

			opts := layout.DefaultLayoutOptions()
			opts.MaxWidth = 400.0

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lines := engine.LayoutParagraph(
					para,
					opts,
				)
				if len(lines) == 0 {
					b.Fatal("No lines generated")
				}
			}
		},
	)
}

// BenchmarkStreamingOutput benchmarks streaming page output.
func BenchmarkStreamingOutput(b *testing.B) {
	doc, err := loadTestWordDocument(b)
	if err != nil {
		b.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()
	pageCount := 0
	opts.OnPageComplete = func(pageNum int, data []byte) {
		pageCount++
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pageCount = 0
		var buf bytes.Buffer
		if err := RenderWord(doc, &buf, opts); err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

// Helper functions to load test documents

func loadTestWordDocument(
	b *testing.B,
) (*wordprocessing.Document, error) {
	// Try to find a test document
	paths := []string{
		filepath.Join(
			"testdata",
			"benchmark.docx",
		),
		filepath.Join(
			"..",
			"testdata",
			"benchmark.docx",
		),
		filepath.Join("testdata", "sample.docx"),
		filepath.Join(
			"..",
			"wordprocessing",
			"testdata",
			"document.docx",
		),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return wordprocessing.Open(
				path,
				false,
			)
		}
	}

	return nil, os.ErrNotExist
}

func loadTestSpreadsheetDocument(
	b *testing.B,
) (*spreadsheet.Document, error) {
	paths := []string{
		filepath.Join(
			"testdata",
			"benchmark.xlsx",
		),
		filepath.Join(
			"..",
			"testdata",
			"benchmark.xlsx",
		),
		filepath.Join("testdata", "sample.xlsx"),
		filepath.Join(
			"..",
			"spreadsheet",
			"testdata",
			"workbook.xlsx",
		),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return spreadsheet.Open(path, false)
		}
	}

	return nil, os.ErrNotExist
}

func loadTestPresentationDocument(
	b *testing.B,
) (*presentation.Document, error) {
	paths := []string{
		filepath.Join(
			"testdata",
			"benchmark.pptx",
		),
		filepath.Join(
			"..",
			"testdata",
			"benchmark.pptx",
		),
		filepath.Join("testdata", "sample.pptx"),
		filepath.Join(
			"..",
			"presentation",
			"testdata",
			"presentation.pptx",
		),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return presentation.Open(path, false)
		}
	}

	return nil, os.ErrNotExist
}

// Helper functions for creating test data

// createTestImage creates a simple test image for benchmarking
func createTestImage(
	width, height int,
) image.Image {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	// Fill with a gradient pattern
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x * 255) / width)
			g := uint8((y * 255) / height)
			b := uint8(128)
			img.Set(
				x,
				y,
				color.RGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				},
			)
		}
	}
	return img
}

// createTestImageWithAlpha creates a test image with alpha channel
func createTestImageWithAlpha(
	width, height int,
) image.Image {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)
	// Fill with a gradient pattern with alpha
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8((x * 255) / width)
			g := uint8((y * 255) / height)
			b := uint8(128)
			a := uint8(
				(x * 255) / width,
			) // Variable alpha
			img.Set(
				x,
				y,
				color.RGBA{
					R: r,
					G: g,
					B: b,
					A: a,
				},
			)
		}
	}
	return img
}

// createTestParagraph creates a test paragraph with a single run
func createTestParagraph(
	text string,
	f *font.Font,
	fontSize float64,
) *layout.Paragraph {
	para := layout.NewParagraph()
	run := createTestRun(text, f, fontSize)
	para.AddRun(run)
	return para
}

// createTestRun creates a test text run with the given font
func createTestRun(
	text string,
	f *font.Font,
	fontSize float64,
) *layout.TextRun {
	// Get glyph metrics from the font
	gm := layout.NewGlyphMetrics(f)

	// Get kerning table if available
	kt := layout.NewKerningTable(f)

	run := &layout.TextRun{
		Text:              text,
		GlyphMetrics:      gm,
		KerningTable:      kt,
		FontSize:          fontSize,
		HorizontalScaling: 1.0,
	}

	return run
}
