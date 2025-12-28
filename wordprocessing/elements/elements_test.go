package elements

import (
	"strings"
	"testing"
)

const (
	testTextHello  = "Hello"
	testColorRed   = "FF0000"
	testFontArial  = "Arial"
	testAuthorName = "Author"
)

func TestPackageExists(t *testing.T) {
	// Simple test to verify test infrastructure works for elements package
	t.Log(
		"elements package test infrastructure is working",
	)
}

func TestNewDocument(t *testing.T) {
	doc := NewDocument()

	if doc == nil {
		t.Fatal("NewDocument returned nil")
	}

	if doc.LocalName() != "document" {
		t.Errorf(
			"Expected LocalName 'document', got %q",
			doc.LocalName(),
		)
	}

	if doc.NamespaceURI() != NamespaceWML {
		t.Errorf(
			"Expected NamespaceURI %q, got %q",
			NamespaceWML,
			doc.NamespaceURI(),
		)
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("Document should have a body")
	}
}

func TestNewBody(t *testing.T) {
	body := NewBody()

	if body == nil {
		t.Fatal("NewBody returned nil")
	}

	if body.LocalName() != "body" {
		t.Errorf(
			"Expected LocalName 'body', got %q",
			body.LocalName(),
		)
	}
}

func TestNewParagraph(t *testing.T) {
	t.Run("empty paragraph", func(t *testing.T) {
		p := NewParagraph()
		if p == nil {
			t.Fatal("NewParagraph returned nil")
		}
		if p.LocalName() != "p" {
			t.Errorf(
				"Expected LocalName 'p', got %q",
				p.LocalName(),
			)
		}
		if p.InnerText() != "" {
			t.Errorf(
				"Expected empty text, got %q",
				p.InnerText(),
			)
		}
	})

	t.Run(
		"paragraph with text",
		func(t *testing.T) {
			p := NewParagraph("Hello, World!")
			if p.InnerText() != "Hello, World!" {
				t.Errorf(
					"Expected 'Hello, World!', got %q",
					p.InnerText(),
				)
			}
		},
	)
}

func TestNewRun(t *testing.T) {
	t.Run("empty run", func(t *testing.T) {
		r := NewRun("")
		if r == nil {
			t.Fatal("NewRun returned nil")
		}
		if r.LocalName() != "r" {
			t.Errorf(
				"Expected LocalName 'r', got %q",
				r.LocalName(),
			)
		}
	})

	t.Run("run with text", func(t *testing.T) {
		r := NewRun(testTextHello)
		if r.InnerText() != testTextHello {
			t.Errorf(
				"Expected %q, got %q",
				testTextHello,
				r.InnerText(),
			)
		}
	})
}

func TestNewText(t *testing.T) {
	text := NewText(testTextHello)
	if text == nil {
		t.Fatal("NewText returned nil")
	}
	if text.LocalName() != "t" {
		t.Errorf(
			"Expected LocalName 't', got %q",
			text.LocalName(),
		)
	}
	if text.InnerText() != testTextHello {
		t.Errorf(
			"Expected %q, got %q",
			testTextHello,
			text.InnerText(),
		)
	}
}

func TestTextSpacePreservation(t *testing.T) {
	t.Run(
		"no space preserve for normal text",
		func(t *testing.T) {
			text := NewText(testTextHello)
			if text.Space() != "" {
				t.Errorf(
					"Expected no space attribute, got %q",
					text.Space(),
				)
			}
		},
	)

	t.Run(
		"space preserve for leading whitespace",
		func(t *testing.T) {
			text := NewText("  Hello")
			if text.Space() != string(
				SpaceProcessingModePreserve,
			) {
				t.Errorf(
					"Expected space='preserve', got %q",
					text.Space(),
				)
			}
		},
	)

	t.Run(
		"space preserve for trailing whitespace",
		func(t *testing.T) {
			text := NewText("Hello  ")
			if text.Space() != string(
				SpaceProcessingModePreserve,
			) {
				t.Errorf(
					"Expected space='preserve', got %q",
					text.Space(),
				)
			}
		},
	)
}

func TestRunProperties(t *testing.T) {
	t.Run("bold", func(t *testing.T) {
		rp := NewRunProperties()
		if rp.Bold() {
			t.Error(
				"Expected Bold to be false initially",
			)
		}
		rp.SetBold(true)
		if !rp.Bold() {
			t.Error(
				"Expected Bold to be true after setting",
			)
		}
	})

	t.Run("italic", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetItalic(true)
		if !rp.Italic() {
			t.Error(
				"Expected Italic to be true after setting",
			)
		}
	})

	t.Run("font size", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetFontSize(24) // 12pt
		if rp.FontSize() != 24 {
			t.Errorf(
				"Expected FontSize 24, got %d",
				rp.FontSize(),
			)
		}
	})

	t.Run("color", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetColor(testColorRed)
		if rp.Color() != testColorRed {
			t.Errorf(
				"Expected Color %q, got %q",
				testColorRed,
				rp.Color(),
			)
		}
	})

	t.Run("underline", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetUnderline(UnderlineSingle)
		if rp.Underline() != UnderlineSingle {
			t.Errorf(
				"Expected Underline Single, got %q",
				rp.Underline(),
			)
		}
	})

	t.Run("highlight", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetHighlight(HighlightYellow)
		if rp.Highlight() != HighlightYellow {
			t.Errorf(
				"Expected Highlight Yellow, got %q",
				rp.Highlight(),
			)
		}
	})
}

func TestParagraphProperties(t *testing.T) {
	t.Run("justification", func(t *testing.T) {
		pp := NewParagraphProperties()
		pp.SetJustification(JustificationCenter)
		if pp.Justification() != JustificationCenter {
			t.Errorf(
				"Expected Justification Center, got %q",
				pp.Justification(),
			)
		}
	})

	t.Run("indentation", func(t *testing.T) {
		pp := NewParagraphProperties()
		ind := pp.GetOrCreateIndentation()
		ind.SetLeft(720)
		ind.SetFirstLine(360)
		if ind.Left() != 720 {
			t.Errorf(
				"Expected Left 720, got %d",
				ind.Left(),
			)
		}
		if ind.FirstLine() != 360 {
			t.Errorf(
				"Expected FirstLine 360, got %d",
				ind.FirstLine(),
			)
		}
	})

	t.Run("spacing", func(t *testing.T) {
		pp := NewParagraphProperties()
		sp := pp.GetOrCreateSpacingBetweenLines()
		sp.SetBefore(240)
		sp.SetAfter(120)
		if sp.Before() != 240 {
			t.Errorf(
				"Expected Before 240, got %d",
				sp.Before(),
			)
		}
		if sp.After() != 120 {
			t.Errorf(
				"Expected After 120, got %d",
				sp.After(),
			)
		}
	})

	t.Run("keep next", func(t *testing.T) {
		pp := NewParagraphProperties()
		pp.SetKeepNext(true)
		if !pp.KeepNext() {
			t.Error(
				"Expected KeepNext to be true",
			)
		}
	})

	t.Run(
		"page break before",
		func(t *testing.T) {
			pp := NewParagraphProperties()
			pp.SetPageBreakBefore(true)
			if !pp.PageBreakBefore() {
				t.Error(
					"Expected PageBreakBefore to be true",
				)
			}
		},
	)
}

func TestBodyAppendParagraph(t *testing.T) {
	body := NewBody()
	p := body.AppendParagraph("Test paragraph")

	if p == nil {
		t.Fatal("AppendParagraph returned nil")
	}

	// Check paragraph was added
	count := 0
	for range body.Paragraphs() {
		count++
	}
	if count != 1 {
		t.Errorf(
			"Expected 1 paragraph, got %d",
			count,
		)
	}
}

func TestParagraphAppendRun(t *testing.T) {
	p := NewParagraph()
	r := p.AppendRun("Hello")

	if r == nil {
		t.Fatal("AppendRun returned nil")
	}

	// Check run was added
	count := 0
	for range p.Runs() {
		count++
	}
	if count != 1 {
		t.Errorf("Expected 1 run, got %d", count)
	}

	if p.InnerText() != testTextHello {
		t.Errorf(
			"Expected %q, got %q",
			testTextHello,
			p.InnerText(),
		)
	}
}

func TestBreak(t *testing.T) {
	t.Run("line break", func(t *testing.T) {
		br := NewLineBreak()
		if br.Type() != BreakLine {
			t.Errorf(
				"Expected BreakLine, got %q",
				br.Type(),
			)
		}
	})

	t.Run("page break", func(t *testing.T) {
		br := NewPageBreak()
		if br.Type() != BreakPage {
			t.Errorf(
				"Expected BreakPage, got %q",
				br.Type(),
			)
		}
	})

	t.Run("column break", func(t *testing.T) {
		br := NewColumnBreak()
		if br.Type() != BreakColumn {
			t.Errorf(
				"Expected BreakColumn, got %q",
				br.Type(),
			)
		}
	})
}

func TestTab(t *testing.T) {
	tab := NewTab()
	if tab == nil {
		t.Fatal("NewTab returned nil")
	}
	if tab.LocalName() != "tab" {
		t.Errorf(
			"Expected LocalName 'tab', got %q",
			tab.LocalName(),
		)
	}
}

func TestBookmarks(t *testing.T) {
	start, end := CreateBookmarkPair(
		1,
		"TestBookmark",
	)

	if start.Id() != 1 {
		t.Errorf(
			"Expected Id 1, got %d",
			start.Id(),
		)
	}
	if start.Name() != "TestBookmark" {
		t.Errorf(
			"Expected Name 'TestBookmark', got %q",
			start.Name(),
		)
	}
	if end.Id() != 1 {
		t.Errorf(
			"Expected Id 1, got %d",
			end.Id(),
		)
	}
}

func TestHyperlink(t *testing.T) {
	t.Run(
		"external hyperlink",
		func(t *testing.T) {
			h := NewHyperlink(
				"Click here",
				"rId1",
			)
			if h.RelationshipId() != "rId1" {
				t.Errorf(
					"Expected RelationshipId 'rId1', got %q",
					h.RelationshipId(),
				)
			}
			if !h.IsExternal() {
				t.Error(
					"Expected IsExternal to be true",
				)
			}
		},
	)

	t.Run(
		"internal hyperlink",
		func(t *testing.T) {
			h := NewInternalHyperlink(
				"Go to section",
				"Section1",
			)
			if h.Anchor() != "Section1" {
				t.Errorf(
					"Expected Anchor 'Section1', got %q",
					h.Anchor(),
				)
			}
			if !h.IsInternal() {
				t.Error(
					"Expected IsInternal to be true",
				)
			}
		},
	)
}

func TestSectionProperties(t *testing.T) {
	t.Run("page size", func(t *testing.T) {
		sp := NewSectionProperties()
		ps := sp.GetOrCreatePageSize()
		ps.SetWidth(15840) // Legal width
		ps.SetHeight(20160)
		if ps.Width() != 15840 {
			t.Errorf(
				"Expected Width 15840, got %d",
				ps.Width(),
			)
		}
		if ps.Height() != 20160 {
			t.Errorf(
				"Expected Height 20160, got %d",
				ps.Height(),
			)
		}
	})

	t.Run("page margins", func(t *testing.T) {
		sp := NewSectionProperties()
		pm := sp.GetOrCreatePageMargins()
		pm.SetTop(1440)
		pm.SetBottom(1440)
		pm.SetLeft(1800)
		pm.SetRight(1800)
		if pm.Top() != 1440 {
			t.Errorf(
				"Expected Top 1440, got %d",
				pm.Top(),
			)
		}
		if pm.Left() != 1800 {
			t.Errorf(
				"Expected Left 1800, got %d",
				pm.Left(),
			)
		}
	})

	t.Run("columns", func(t *testing.T) {
		sp := NewSectionProperties()
		cols := sp.GetOrCreateColumns()
		cols.SetNum(2)
		cols.SetSpace(720)
		if cols.Num() != 2 {
			t.Errorf(
				"Expected Num 2, got %d",
				cols.Num(),
			)
		}
		if cols.Space() != 720 {
			t.Errorf(
				"Expected Space 720, got %d",
				cols.Space(),
			)
		}
	})

	t.Run("orientation", func(t *testing.T) {
		sp := NewSectionProperties()
		ps := sp.GetOrCreatePageSize()
		ps.SetOrient(PageOrientationLandscape)
		if ps.Orient() != PageOrientationLandscape {
			t.Errorf(
				"Expected Landscape, got %q",
				ps.Orient(),
			)
		}
	})
}

func TestTable(t *testing.T) {
	table := NewTable(3, 4)

	if table == nil {
		t.Fatal("NewTable returned nil")
	}

	if table.LocalName() != "tbl" {
		t.Errorf(
			"Expected LocalName 'tbl', got %q",
			table.LocalName(),
		)
	}
}

func TestDocumentXMLOutput(t *testing.T) {
	doc := NewDocument()
	body := doc.Body()
	p := body.AppendParagraph("Hello, World!")

	// Apply some formatting
	p.SetJustification(JustificationCenter)
	r := p.AppendRun("Bold text")
	r.SetBold(true)

	xml := doc.OuterXml()

	// Verify XML contains expected elements
	if !strings.Contains(xml, "w:document") {
		t.Error("XML should contain w:document")
	}
	if !strings.Contains(xml, "w:body") {
		t.Error("XML should contain w:body")
	}
	if !strings.Contains(xml, "w:p") {
		t.Error("XML should contain w:p")
	}
	if !strings.Contains(xml, "Hello, World!") {
		t.Error(
			"XML should contain 'Hello, World!'",
		)
	}
	if !strings.Contains(xml, "Bold text") {
		t.Error("XML should contain 'Bold text'")
	}
}

func TestEnums(t *testing.T) {
	// Test that enum values are correct strings
	if string(JustificationCenter) != "center" {
		t.Error(
			"JustificationCenter should be 'center'",
		)
	}

	if string(UnderlineSingle) != "single" {
		t.Error(
			"UnderlineSingle should be 'single'",
		)
	}

	if string(HighlightYellow) != "yellow" {
		t.Error(
			"HighlightYellow should be 'yellow'",
		)
	}

	if string(BreakPage) != "page" {
		t.Error("BreakPage should be 'page'")
	}

	if string(BorderSingle) != "single" {
		t.Error(
			"BorderSingle should be 'single'",
		)
	}
}

func TestClone(t *testing.T) {
	p := NewParagraph("Original text")
	p.SetStyle("Heading1")

	clone, ok := p.Clone().(*Paragraph)
	if !ok {
		t.Fatal("Clone did not return *Paragraph")
	}

	// Modify original
	p.SetText("Modified text")

	// Clone should still have original text
	if clone.InnerText() != "Original text" {
		t.Errorf(
			"Clone text should be 'Original text', got %q",
			clone.InnerText(),
		)
	}
}

func TestRunConvenienceMethods(t *testing.T) {
	r := NewRun("Test")
	r.SetBold(true).
		SetItalic(true).
		SetFontSize(28).
		SetColor(testColorRed)

	props := r.Properties()
	if props == nil {
		t.Fatal("Properties should not be nil")
	}

	if !props.Bold() {
		t.Error("Expected Bold to be true")
	}
	if !props.Italic() {
		t.Error("Expected Italic to be true")
	}
	if props.FontSize() != 28 {
		t.Errorf(
			"Expected FontSize 28, got %d",
			props.FontSize(),
		)
	}
	if props.Color() != testColorRed {
		t.Errorf(
			"Expected Color %q, got %q",
			testColorRed,
			props.Color(),
		)
	}
}

func TestParagraphConvenienceMethods(
	t *testing.T,
) {
	p := NewParagraph("Test")
	p.SetStyle("Heading1").
		SetJustification(JustificationRight).
		SetSpacingBefore(240).
		SetLeftIndent(720)

	props := p.Properties()
	if props == nil {
		t.Fatal("Properties should not be nil")
	}

	if props.ParagraphStyleId() != "Heading1" {
		t.Errorf(
			"Expected style 'Heading1', got %q",
			props.ParagraphStyleId(),
		)
	}
	if props.Justification() != JustificationRight {
		t.Errorf(
			"Expected justification Right, got %q",
			props.Justification(),
		)
	}
}
