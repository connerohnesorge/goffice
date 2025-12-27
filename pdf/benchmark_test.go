package pdf

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

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
	// TODO: Implement font loading benchmark
	b.Skip(
		"Font loading benchmark not yet implemented",
	)
}

// BenchmarkImageDecoding benchmarks image decoding and embedding.
func BenchmarkImageDecoding(b *testing.B) {
	// TODO: Implement image decoding benchmark
	b.Skip(
		"Image decoding benchmark not yet implemented",
	)
}

// BenchmarkTextLayout benchmarks text layout engine.
func BenchmarkTextLayout(b *testing.B) {
	// TODO: Implement text layout benchmark
	b.Skip(
		"Text layout benchmark not yet implemented",
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
