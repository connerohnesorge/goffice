//go:build profiling
// +build profiling

package pdf

import (
	"bytes"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

// TestProfileWordRendering profiles Word document rendering.
// Run with: go test -tags=profiling -run=TestProfileWordRendering -cpuprofile=cpu.prof -memprofile=mem.prof
func TestProfileWordRendering(t *testing.T) {
	doc, err := loadTestWordDocument(&testing.B{})
	if err != nil {
		t.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()

	// Start CPU profiling
	f, err := os.Create("pdf_word_cpu.prof")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Render multiple times to get good profile data
	for i := 0; i < 10; i++ {
		var buf bytes.Buffer
		if err := RenderWord(doc, &buf, opts); err != nil {
			t.Fatalf("Render failed: %v", err)
		}
	}

	// Write memory profile
	mf, err := os.Create("pdf_word_mem.prof")
	if err != nil {
		t.Fatal(err)
	}
	defer mf.Close()

	runtime.GC() // Force GC to get accurate memory stats
	if err := pprof.WriteHeapProfile(mf); err != nil {
		t.Fatal(err)
	}
}

// TestProfileSpreadsheetRendering profiles Excel rendering.
func TestProfileSpreadsheetRendering(
	t *testing.T,
) {
	doc, err := loadTestSpreadsheetDocument(
		&testing.B{},
	)
	if err != nil {
		t.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()

	f, err := os.Create(
		"pdf_spreadsheet_cpu.prof",
	)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		var buf bytes.Buffer
		if err := RenderSpreadsheet(doc, &buf, opts); err != nil {
			t.Fatalf("Render failed: %v", err)
		}
	}

	mf, err := os.Create(
		"pdf_spreadsheet_mem.prof",
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mf.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(mf); err != nil {
		t.Fatal(err)
	}
}

// TestProfilePresentationRendering profiles PowerPoint rendering.
func TestProfilePresentationRendering(
	t *testing.T,
) {
	doc, err := loadTestPresentationDocument(
		&testing.B{},
	)
	if err != nil {
		t.Skipf(
			"Test document not available: %v",
			err,
		)
		return
	}
	defer doc.Close()

	opts := DefaultRenderOptions()

	f, err := os.Create(
		"pdf_presentation_cpu.prof",
	)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		var buf bytes.Buffer
		if err := RenderPresentation(doc, &buf, opts); err != nil {
			t.Fatalf("Render failed: %v", err)
		}
	}

	mf, err := os.Create(
		"pdf_presentation_mem.prof",
	)
	if err != nil {
		t.Fatal(err)
	}
	defer mf.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(mf); err != nil {
		t.Fatal(err)
	}
}
