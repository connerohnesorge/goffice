package layout

import (
	"testing"

	"github.com/connerohnesorge/goffice-pdf/font"
)

func TestNewTextLayoutEngine(t *testing.T) {
	fc := font.NewFontCache(10)
	engine := NewTextLayoutEngine(fc)

	if engine == nil {
		t.Fatal(
			"NewTextLayoutEngine returned nil",
		)
	}

	if engine.FontCache() != fc {
		t.Error("FontCache not set correctly")
	}

	if engine.lineBreaker == nil {
		t.Error("LineBreaker not initialized")
	}

	if engine.DefaultFontSize() != 12.0 {
		t.Errorf(
			"DefaultFontSize = %v, want 12.0",
			engine.DefaultFontSize(),
		)
	}
}

func TestTextLayoutEngine_SetDefaultFontSize(
	t *testing.T,
) {
	engine := NewTextLayoutEngine(nil)

	engine.SetDefaultFontSize(10.5)
	if engine.DefaultFontSize() != 10.5 {
		t.Errorf(
			"DefaultFontSize = %v, want 10.5",
			engine.DefaultFontSize(),
		)
	}
}

func TestDefaultLayoutOptions(t *testing.T) {
	opts := DefaultLayoutOptions()

	if opts.MaxWidth != 500.0 {
		t.Errorf(
			"Default MaxWidth = %v, want 500.0",
			opts.MaxWidth,
		)
	}

	if opts.Hyphenation {
		t.Error(
			"Default Hyphenation should be false",
		)
	}

	if opts.Justify {
		t.Error("Default Justify should be false")
	}
}

func TestLayoutParagraph_Simple(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// "Hello World" = 61.2 pt
	p.AddRun(NewTextRun("Hello World", gm, 12.0))

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 100.0

	lines := engine.LayoutParagraph(p, opts)

	if len(lines) != 1 {
		t.Errorf(
			"Expected 1 line, got %d",
			len(lines),
		)
	}

	if lines[0].Width != 61.2 {
		t.Errorf(
			"Line width = %v, want 61.2",
			lines[0].Width,
		)
	}
}

func TestLayoutParagraph_Wrapping(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// "Hello" = 27.0 pt, " " = 3.0 pt, "World" = 31.2 pt
	// Total "Hello World" = 61.2 pt
	p.AddRun(NewTextRun("Hello World", gm, 12.0))

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 40.0 // Should break between Hello and World

	lines := engine.LayoutParagraph(p, opts)

	if len(lines) != 2 {
		t.Errorf(
			"Expected 2 lines, got %d",
			len(lines),
		)
	}

	if len(lines) >= 1 {
		// "Hello " = 27.0 + 3.0 = 30.0
		if !floatEquals(
			lines[0].Width,
			30.0,
			0.001,
		) {
			t.Errorf(
				"Line 0 width = %v, want 30.0",
				lines[0].Width,
			)
		}
		if len(lines[0].Runs) != 1 ||
			lines[0].Runs[0].Run.Text != "Hello " {
			t.Errorf(
				"Line 0 content mismatch: got %q",
				lines[0].Runs[0].Run.Text,
			)
		}
	}

	if len(lines) >= 2 {
		if !floatEquals(
			lines[1].Width,
			31.2,
			0.001,
		) { // "World"
			t.Errorf(
				"Line 1 width = %v, want 31.2",
				lines[1].Width,
			)
		}
		if len(lines[1].Runs) != 1 ||
			lines[1].Runs[0].Run.Text != "World" {
			t.Errorf(
				"Line 1 content mismatch: got %q",
				lines[1].Runs[0].Run.Text,
			)
		}
	}
}

func TestLayoutParagraph_Indentation(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	p.AddRun(NewTextRun("Hello World", gm, 12.0))
	p.Properties.Indentation.Left = 36.0 // 0.5 inch

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 100.0 // Available: 100 - 36 = 64.0. "Hello World" (61.2) should fit.

	lines := engine.LayoutParagraph(p, opts)

	if len(lines) != 1 {
		t.Errorf(
			"Expected 1 line, got %d",
			len(lines),
		)
	}

	// First line available was 64.0, "Hello World" is 61.2.
	// Note: currently createLineFromRange doesn't add the indentation X offset to the Runs themselves?
	// It starts currentX at 0.0. The caller (renderer) will probably apply the indent.
}

func TestLayoutParagraph_Justification(
	t *testing.T,
) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// "Hello" = 27.0, " " = 3.0, "World" = 31.2. Total = 61.2
	p.AddRun(NewTextRun("Hello World", gm, 12.0))
	p.Properties.Alignment = AlignJustify

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 100.0 // 100.0 available. "Hello World" is 61.2.
	// But it's the LAST line, so it should NOT be justified by default.

	lines := engine.LayoutParagraph(p, opts)
	if lines[0].Width != 61.2 {
		t.Errorf(
			"Last line of justified paragraph should not be justified, width = %v",
			lines[0].Width,
		)
	}

	// Now force multi-line so we have a non-last line
	p = NewParagraph()
	// "Hello World Hello World"
	p.AddRun(
		NewTextRun(
			"Hello World Hello World",
			gm,
			12.0,
		),
	)
	p.Properties.Alignment = AlignJustify
	opts.MaxWidth = 100.0 // "Hello World " = 64.2. Next "Hello" makes it > 100.
	// Line 0: "Hello World " (64.2)
	// Remaining: "Hello World" (61.2) - last line

	lines = engine.LayoutParagraph(p, opts)
	if len(lines) != 2 {
		t.Fatalf(
			"Expected 2 lines, got %d",
			len(lines),
		)
	}

	if !floatEquals(
		lines[0].Width,
		100.0,
		0.001,
	) {
		t.Errorf(
			"First line should be justified to 100.0, got %v",
			lines[0].Width,
		)
	}

	if !floatEquals(lines[1].Width, 31.2, 0.001) {
		t.Errorf(
			"Last line should NOT be justified, got %v",
			lines[1].Width,
		)
	}
}

func TestLayoutParagraph_Tabs(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// "Hello" = 27.0
	// "\t" jumps to 36.0 (default tab interval)
	// "World" = 31.2
	// Total line width should be 36.0 + 31.2 = 67.2
	p.AddRun(NewTextRun("Hello\tWorld", gm, 12.0))

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 200.0

	lines := engine.LayoutParagraph(p, opts)
	if len(lines) != 1 {
		t.Fatalf(
			"Expected 1 line, got %d",
			len(lines),
		)
	}

	if !floatEquals(lines[0].Width, 67.2, 0.001) {
		t.Errorf(
			"Line width with tab = %v, want 67.2",
			lines[0].Width,
		)
	}

	if len(lines[0].Runs) != 2 {
		t.Errorf(
			"Expected 2 positioned runs (split by tab), got %d",
			len(lines[0].Runs),
		)
	}

	if len(lines[0].Runs) >= 2 {
		if lines[0].Runs[1].X != 36.0 {
			t.Errorf(
				"Second run should start at 36.0, got %v",
				lines[0].Runs[1].X,
			)
		}
	}
}

func TestLayoutParagraph_Complex(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// Run 1: "Hello " (30.0)
	p.AddRun(NewTextRun("Hello ", gm, 12.0))
	// Run 2: "Big World" (Big=18.48, space=3.0, World=31.2. Total=52.68)
	// Word: B(540) i(220) g(520) = 1280 = 15.36 pt.
	// Wait, 'B' is not in my mock font? Let's check mockFontForParagraph.
	// B is not there. It will use replacement char (600 = 7.2 pt).
	// 'B' i(220) g(520) = 1340 = 16.08 pt.
	p.AddRun(NewTextRun("Big World", gm, 12.0))

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 60.0
	// Line 0 avail: 60.0. "Hello " (30.0) + "Big " (16.08 + 3.0 = 19.08) = 49.08.
	// Next "World" (31.2) would make it 80.28 > 60.

	lines := engine.LayoutParagraph(p, opts)
	if len(lines) != 2 {
		t.Fatalf(
			"Expected 2 lines, got %d",
			len(lines),
		)
	}

	if len(lines[0].Runs) != 2 {
		t.Errorf(
			"Line 0 should have 2 runs, got %d",
			len(lines[0].Runs),
		)
	}

	if lines[0].Runs[0].Run.Text != "Hello " {
		t.Errorf(
			"Line 0 run 0 text = %q, want \"Hello \"",
			lines[0].Runs[0].Run.Text,
		)
	}

	if lines[0].Runs[1].Run.Text != "Big " {
		t.Errorf(
			"Line 0 run 1 text = %q, want \"Big \"",
			lines[0].Runs[1].Run.Text,
		)
	}
}

func TestLayoutParagraph_RTL(t *testing.T) {
	f := mockFontForParagraph()
	gm := NewGlyphMetrics(f)
	fc := font.NewFontCache(1)
	fc.Put(f)

	engine := NewTextLayoutEngine(fc)

	p := NewParagraph()
	// "שלום" - all Hebrew, should be RTL base
	p.AddRun(NewTextRun("שלום", gm, 12.0))

	opts := DefaultLayoutOptions()
	opts.MaxWidth = 100.0

	lines := engine.LayoutParagraph(p, opts)
	if len(lines) != 1 {
		t.Fatalf(
			"Expected 1 line, got %d",
			len(lines),
		)
	}

	// For RTL base, it should probably be right-aligned if it's the only line
	// and no explicit alignment is set?
	// Actually, Word's default alignment is Left, but for RTL paragraphs it might be Right.
	// In our implementation, we follow p.Properties.Alignment.
}
