package elements

import (
	"strings"
	"testing"
)

func TestSheetData(t *testing.T) {
	t.Run("NewSheetData", func(t *testing.T) {
		sd := NewSheetData()
		if sd == nil {
			t.Fatal("expected non-nil SheetData")
		}
		if sd.LocalName() != "sheetData" {
			t.Errorf(
				"expected localName sheetData, got %s",
				sd.LocalName(),
			)
		}
		if sd.NamespaceURI() != NamespaceSML {
			t.Errorf(
				"expected namespace %s, got %s",
				NamespaceSML,
				sd.NamespaceURI(),
			)
		}
	})

	t.Run("AddRow", func(t *testing.T) {
		sd := NewSheetData()
		row1 := sd.AddRow(1)
		if row1 == nil {
			t.Fatal("expected non-nil row")
		}
		if row1.RowIndex() != 1 {
			t.Errorf(
				"expected row index 1, got %d",
				row1.RowIndex(),
			)
		}

		row3 := sd.AddRow(3)
		if row3.RowIndex() != 3 {
			t.Errorf(
				"expected row index 3, got %d",
				row3.RowIndex(),
			)
		}

		// Add row 2 - should be inserted between 1 and 3
		row2 := sd.AddRow(2)
		if row2.RowIndex() != 2 {
			t.Errorf(
				"expected row index 2, got %d",
				row2.RowIndex(),
			)
		}

		// Verify order
		expectedOrder := []uint32{1, 2, 3}
		i := 0
		for row := range sd.Rows() {
			if row.RowIndex() != expectedOrder[i] {
				t.Errorf(
					"expected row %d at position %d, got %d",
					expectedOrder[i],
					i,
					row.RowIndex(),
				)
			}
			i++
		}
	})

	t.Run("GetRow", func(t *testing.T) {
		sd := NewSheetData()
		sd.AddRow(1)
		sd.AddRow(5)

		if sd.GetRow(1) == nil {
			t.Error("expected to find row 1")
		}
		if sd.GetRow(5) == nil {
			t.Error("expected to find row 5")
		}
		if sd.GetRow(3) != nil {
			t.Error("expected row 3 to be nil")
		}
	})

	t.Run("GetOrCreateRow", func(t *testing.T) {
		sd := NewSheetData()
		row := sd.GetOrCreateRow(1)
		if row == nil {
			t.Fatal("expected non-nil row")
		}

		// Getting same row again should return same instance
		row2 := sd.GetOrCreateRow(1)
		if row.RowIndex() != row2.RowIndex() {
			t.Error("expected same row instance")
		}
	})

	t.Run("RemoveRow", func(t *testing.T) {
		sd := NewSheetData()
		sd.AddRow(1)
		sd.AddRow(2)

		if !sd.RemoveRow(1) {
			t.Error(
				"expected RemoveRow to return true",
			)
		}
		if sd.GetRow(1) != nil {
			t.Error(
				"expected row 1 to be removed",
			)
		}
		if sd.GetRow(2) == nil {
			t.Error(
				"expected row 2 to still exist",
			)
		}
	})

	t.Run("RowCount", func(t *testing.T) {
		sd := NewSheetData()
		if sd.RowCount() != 0 {
			t.Errorf(
				"expected 0 rows, got %d",
				sd.RowCount(),
			)
		}

		sd.AddRow(1)
		sd.AddRow(2)
		if sd.RowCount() != 2 {
			t.Errorf(
				"expected 2 rows, got %d",
				sd.RowCount(),
			)
		}
	})
}

func TestRow(t *testing.T) {
	t.Run("NewRow", func(t *testing.T) {
		row := NewRow()
		if row == nil {
			t.Fatal("expected non-nil Row")
		}
		if row.LocalName() != "row" {
			t.Errorf(
				"expected localName row, got %s",
				row.LocalName(),
			)
		}
	})

	t.Run("RowIndex", func(t *testing.T) {
		row := NewRow()
		row.SetRowIndex(5)
		if row.RowIndex() != 5 {
			t.Errorf(
				"expected row index 5, got %d",
				row.RowIndex(),
			)
		}
	})

	t.Run("Spans", func(t *testing.T) {
		row := NewRow()
		row.SetSpans("1:10")
		if row.Spans() != "1:10" {
			t.Errorf(
				"expected spans 1:10, got %s",
				row.Spans(),
			)
		}
	})

	t.Run("StyleIndex", func(t *testing.T) {
		row := NewRow()
		row.SetStyleIndex(5)
		if row.StyleIndex() != 5 {
			t.Errorf(
				"expected style index 5, got %d",
				row.StyleIndex(),
			)
		}

		// Setting to 0 should remove attribute
		row.SetStyleIndex(0)
		if row.StyleIndex() != 0 {
			t.Errorf(
				"expected style index 0, got %d",
				row.StyleIndex(),
			)
		}
	})

	t.Run("Height", func(t *testing.T) {
		row := NewRow()
		row.SetHeight(15.5)
		if row.Height() != 15.5 {
			t.Errorf(
				"expected height 15.5, got %f",
				row.Height(),
			)
		}
	})

	t.Run("Hidden", func(t *testing.T) {
		row := NewRow()
		if row.Hidden() {
			t.Error(
				"expected Hidden to be false by default",
			)
		}

		row.SetHidden(true)
		if !row.Hidden() {
			t.Error("expected Hidden to be true")
		}

		row.SetHidden(false)
		if row.Hidden() {
			t.Error("expected Hidden to be false")
		}
	})

	t.Run("CustomHeight", func(t *testing.T) {
		row := NewRow()
		row.SetCustomHeight(true)
		if !row.CustomHeight() {
			t.Error(
				"expected CustomHeight to be true",
			)
		}
	})

	t.Run("OutlineLevel", func(t *testing.T) {
		row := NewRow()
		row.SetOutlineLevel(3)
		if row.OutlineLevel() != 3 {
			t.Errorf(
				"expected outline level 3, got %d",
				row.OutlineLevel(),
			)
		}
	})

	t.Run("AddCell", func(t *testing.T) {
		row := NewRow()
		row.SetRowIndex(1)

		cellA := row.AddCell("A1")
		if cellA == nil {
			t.Fatal("expected non-nil cell")
		}
		if cellA.Reference() != "A1" {
			t.Errorf(
				"expected reference A1, got %s",
				cellA.Reference(),
			)
		}

		// Add cell C1
		row.AddCell("C1")

		// Add cell B1 - should be inserted between A1 and C1
		row.AddCell("B1")

		// Verify order
		expectedOrder := []string{
			"A1",
			"B1",
			"C1",
		}
		i := 0
		for cell := range row.Cells() {
			if cell.Reference() != expectedOrder[i] {
				t.Errorf(
					"expected cell %s at position %d, got %s",
					expectedOrder[i],
					i,
					cell.Reference(),
				)
			}
			i++
		}
	})

	t.Run("GetCell", func(t *testing.T) {
		row := NewRow()
		row.AddCell("A1")
		row.AddCell("B1")

		if row.GetCell("A1") == nil {
			t.Error("expected to find cell A1")
		}
		if row.GetCell("C1") != nil {
			t.Error("expected cell C1 to be nil")
		}
	})

	t.Run("CellCount", func(t *testing.T) {
		row := NewRow()
		if row.CellCount() != 0 {
			t.Errorf(
				"expected 0 cells, got %d",
				row.CellCount(),
			)
		}

		row.AddCell("A1")
		row.AddCell("B1")
		if row.CellCount() != 2 {
			t.Errorf(
				"expected 2 cells, got %d",
				row.CellCount(),
			)
		}
	})
}

//nolint:revive // cyclomatic: comprehensive test coverage requires many cases
func TestCell(t *testing.T) {
	t.Run("NewCell", func(t *testing.T) {
		cell := CreateCell()
		if cell == nil {
			t.Fatal("expected non-nil Cell")
		}
		if cell.LocalName() != "c" {
			t.Errorf(
				"expected localName c, got %s",
				cell.LocalName(),
			)
		}
	})

	t.Run("Reference", func(t *testing.T) {
		cell := CreateCell()
		cell.SetReference("A1")
		if cell.Reference() != "A1" {
			t.Errorf(
				"expected reference A1, got %s",
				cell.Reference(),
			)
		}
	})

	t.Run("StyleIndex", func(t *testing.T) {
		cell := CreateCell()
		cell.SetStyleIndex(5)
		if cell.StyleIndex() != 5 {
			t.Errorf(
				"expected style index 5, got %d",
				cell.StyleIndex(),
			)
		}
	})

	t.Run("DataType", func(t *testing.T) {
		cell := CreateCell()
		// Default is number
		if cell.DataType() != CellTypeNumber {
			t.Errorf(
				"expected default type number, got %s",
				cell.DataType(),
			)
		}

		cell.SetDataType(CellTypeSharedString)
		if cell.DataType() != CellTypeSharedString {
			t.Errorf(
				"expected type s, got %s",
				cell.DataType(),
			)
		}

		cell.SetDataType(CellTypeBoolean)
		if cell.DataType() != CellTypeBoolean {
			t.Errorf(
				"expected type b, got %s",
				cell.DataType(),
			)
		}
	})

	t.Run("SetNumberValue", func(t *testing.T) {
		cell := CreateCell()
		cell.SetNumberValue(123.45)
		if cell.DataType() != CellTypeNumber {
			t.Errorf(
				"expected type n, got %s",
				cell.DataType(),
			)
		}
		if cell.Value() != "123.45" {
			t.Errorf(
				"expected value 123.45, got %s",
				cell.Value(),
			)
		}
	})

	t.Run("SetBoolValue", func(t *testing.T) {
		cell := CreateCell()
		cell.SetBoolValue(true)
		if cell.DataType() != CellTypeBoolean {
			t.Errorf(
				"expected type b, got %s",
				cell.DataType(),
			)
		}
		if cell.Value() != "1" {
			t.Errorf(
				"expected value 1, got %s",
				cell.Value(),
			)
		}

		cell.SetBoolValue(false)
		if cell.Value() != "0" {
			t.Errorf(
				"expected value 0, got %s",
				cell.Value(),
			)
		}
	})

	t.Run(
		"SetSharedStringIndex",
		func(t *testing.T) {
			cell := CreateCell()
			cell.SetSharedStringIndex(42)
			if cell.DataType() != CellTypeSharedString {
				t.Errorf(
					"expected type s, got %s",
					cell.DataType(),
				)
			}
			if cell.Value() != "42" {
				t.Errorf(
					"expected value 42, got %s",
					cell.Value(),
				)
			}
		},
	)

	t.Run("CellValue", func(t *testing.T) {
		cell := CreateCell()
		if cell.CellValue() != nil {
			t.Error(
				"expected nil CellValue initially",
			)
		}

		cv := cell.GetOrCreateCellValue()
		if cv == nil {
			t.Fatal("expected non-nil CellValue")
		}
		cv.SetValue("test")

		cv2 := cell.CellValue()
		if cv2 == nil {
			t.Fatal("expected non-nil CellValue")
		}
		if cv2.Value() != "test" {
			t.Errorf(
				"expected value test, got %s",
				cv2.Value(),
			)
		}
	})

	t.Run("CellFormula", func(t *testing.T) {
		cell := CreateCell()
		if cell.CellFormula() != nil {
			t.Error(
				"expected nil CellFormula initially",
			)
		}

		cell.SetFormula("SUM(A1:A10)")
		if !cell.HasFormula() {
			t.Error(
				"expected HasFormula to be true",
			)
		}
		if cell.Formula() != "SUM(A1:A10)" {
			t.Errorf(
				"expected formula SUM(A1:A10), got %s",
				cell.Formula(),
			)
		}

		cell.RemoveFormula()
		if cell.HasFormula() {
			t.Error(
				"expected HasFormula to be false after removal",
			)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		cell := CreateCell()
		cell.SetValue("test")
		cell.SetFormula("=A1")

		cell.Clear()

		if cell.CellValue() != nil {
			t.Error(
				"expected nil CellValue after clear",
			)
		}
		if cell.CellFormula() != nil {
			t.Error(
				"expected nil CellFormula after clear",
			)
		}
	})
}

func TestCellValue(t *testing.T) {
	t.Run("NewCellValue", func(t *testing.T) {
		cv := NewCellValue()
		if cv == nil {
			t.Fatal("expected non-nil CellValue")
		}
		if cv.LocalName() != "v" {
			t.Errorf(
				"expected localName v, got %s",
				cv.LocalName(),
			)
		}
	})

	t.Run(
		"NewCellValueWithText",
		func(t *testing.T) {
			cv := NewCellValueWithText("123")
			if cv.Value() != "123" {
				t.Errorf(
					"expected value 123, got %s",
					cv.Value(),
				)
			}
		},
	)

	t.Run("Value", func(t *testing.T) {
		cv := NewCellValue()
		cv.SetValue("test value")
		if cv.Value() != "test value" {
			t.Errorf(
				"expected value 'test value', got %s",
				cv.Value(),
			)
		}
	})
}

func TestCellFormula(t *testing.T) {
	t.Run("NewCellFormula", func(t *testing.T) {
		cf := NewCellFormula()
		if cf == nil {
			t.Fatal(
				"expected non-nil CellFormula",
			)
		}
		if cf.LocalName() != "f" {
			t.Errorf(
				"expected localName f, got %s",
				cf.LocalName(),
			)
		}
	})

	t.Run("Formula", func(t *testing.T) {
		cf := NewCellFormula()
		cf.SetFormula("SUM(A1:A10)")
		if cf.Formula() != "SUM(A1:A10)" {
			t.Errorf(
				"expected formula SUM(A1:A10), got %s",
				cf.Formula(),
			)
		}
	})

	t.Run("FormulaType", func(t *testing.T) {
		cf := NewCellFormula()
		if cf.FormulaType() != FormulaTypeNormal {
			t.Errorf(
				"expected default type normal, got %s",
				cf.FormulaType(),
			)
		}

		cf.SetFormulaType(FormulaTypeArray)
		if cf.FormulaType() != FormulaTypeArray {
			t.Errorf(
				"expected type array, got %s",
				cf.FormulaType(),
			)
		}

		cf.SetFormulaType(FormulaTypeShared)
		if cf.FormulaType() != FormulaTypeShared {
			t.Errorf(
				"expected type shared, got %s",
				cf.FormulaType(),
			)
		}
	})

	t.Run("Ref", func(t *testing.T) {
		cf := NewCellFormula()
		cf.SetRef("A1:B10")
		if cf.Ref() != "A1:B10" {
			t.Errorf(
				"expected ref A1:B10, got %s",
				cf.Ref(),
			)
		}
	})

	t.Run("SharedIndex", func(t *testing.T) {
		cf := NewCellFormula()
		cf.SetSharedIndex(5)
		if cf.SharedIndex() != 5 {
			t.Errorf(
				"expected shared index 5, got %d",
				cf.SharedIndex(),
			)
		}
	})
}

func TestInlineString(t *testing.T) {
	t.Run("NewInlineString", func(t *testing.T) {
		is := NewInlineString()
		if is == nil {
			t.Fatal(
				"expected non-nil InlineString",
			)
		}
		if is.LocalName() != "is" {
			t.Errorf(
				"expected localName is, got %s",
				is.LocalName(),
			)
		}
	})

	t.Run("SetPlainText", func(t *testing.T) {
		is := NewInlineString()
		is.SetPlainText("Hello World")
		if is.PlainText() != "Hello World" {
			t.Errorf(
				"expected text 'Hello World', got %s",
				is.PlainText(),
			)
		}
	})

	t.Run("RichTextRuns", func(t *testing.T) {
		is := NewInlineString()
		rtr1 := is.AddRichTextRun()
		rtr1.SetTextContent("Bold")

		_ = is.AddRichTextRunWithText("Normal")

		count := 0
		for range is.RichTextRuns() {
			count++
		}
		if count != 2 {
			t.Errorf(
				"expected 2 rich text runs, got %d",
				count,
			)
		}

		if is.PlainText() != "BoldNormal" {
			t.Errorf(
				"expected concatenated text 'BoldNormal', got %s",
				is.PlainText(),
			)
		}
	})

	t.Run("HasRichText", func(t *testing.T) {
		is := NewInlineString()
		if is.HasRichText() {
			t.Error(
				"expected HasRichText to be false initially",
			)
		}

		is.AddRichTextRun()
		if !is.HasRichText() {
			t.Error(
				"expected HasRichText to be true",
			)
		}
	})
}

func TestRichTextRun(t *testing.T) {
	t.Run("NewRichTextRun", func(t *testing.T) {
		rtr := NewRichTextRun()
		if rtr == nil {
			t.Fatal(
				"expected non-nil RichTextRun",
			)
		}
		if rtr.LocalName() != "r" {
			t.Errorf(
				"expected localName r, got %s",
				rtr.LocalName(),
			)
		}
	})

	t.Run("TextContent", func(t *testing.T) {
		rtr := NewRichTextRun()
		rtr.SetTextContent("Hello")
		if rtr.TextContent() != "Hello" {
			t.Errorf(
				"expected text 'Hello', got %s",
				rtr.TextContent(),
			)
		}
	})

	t.Run("RunProperties", func(t *testing.T) {
		rtr := NewRichTextRun()
		if rtr.RunProperties() != nil {
			t.Error(
				"expected nil RunProperties initially",
			)
		}

		rp := rtr.GetOrCreateRunProperties()
		if rp == nil {
			t.Fatal(
				"expected non-nil RunProperties",
			)
		}

		rp.SetBold(true)
		rp.SetItalic(true)

		rp2 := rtr.RunProperties()
		if !rp2.IsBold() {
			t.Error("expected bold to be true")
		}
		if !rp2.IsItalic() {
			t.Error("expected italic to be true")
		}
	})
}

func TestRunProperties(t *testing.T) {
	t.Run("Bold", func(t *testing.T) {
		rp := NewRunProperties()
		if rp.IsBold() {
			t.Error(
				"expected IsBold to be false initially",
			)
		}

		rp.SetBold(true)
		if !rp.IsBold() {
			t.Error("expected IsBold to be true")
		}

		rp.SetBold(false)
		if rp.IsBold() {
			t.Error(
				"expected IsBold to be false after setting false",
			)
		}
	})

	t.Run("Italic", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetItalic(true)
		if !rp.IsItalic() {
			t.Error(
				"expected IsItalic to be true",
			)
		}
	})

	t.Run("Strikethrough", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetStrikethrough(true)
		if !rp.IsStrikethrough() {
			t.Error(
				"expected IsStrikethrough to be true",
			)
		}
	})

	t.Run("Underline", func(t *testing.T) {
		rp := NewRunProperties()
		if rp.UnderlineStyle() != UnderlineStyleNone {
			t.Errorf(
				"expected no underline, got %s",
				rp.UnderlineStyle(),
			)
		}

		rp.SetUnderline(UnderlineStyleSingle)
		if rp.UnderlineStyle() != UnderlineStyleSingle {
			t.Errorf(
				"expected single underline, got %s",
				rp.UnderlineStyle(),
			)
		}

		rp.SetUnderline(UnderlineStyleDouble)
		if rp.UnderlineStyle() != UnderlineStyleDouble {
			t.Errorf(
				"expected double underline, got %s",
				rp.UnderlineStyle(),
			)
		}
	})

	t.Run(
		"VerticalAlignment",
		func(t *testing.T) {
			rp := NewRunProperties()
			if rp.VerticalAlignment() != VerticalAlignRunBaseline {
				t.Errorf(
					"expected baseline, got %s",
					rp.VerticalAlignment(),
				)
			}

			rp.SetVerticalAlignment(
				VerticalAlignRunSuperscript,
			)
			if rp.VerticalAlignment() != VerticalAlignRunSuperscript {
				t.Errorf(
					"expected superscript, got %s",
					rp.VerticalAlignment(),
				)
			}
		},
	)

	t.Run("FontSize", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetSize(12.5)
		if rp.Size() != 12.5 {
			t.Errorf(
				"expected size 12.5, got %f",
				rp.Size(),
			)
		}
	})

	t.Run("FontName", func(t *testing.T) {
		rp := NewRunProperties()
		rp.SetFont("Arial")
		if rp.Font() != "Arial" {
			t.Errorf(
				"expected font Arial, got %s",
				rp.Font(),
			)
		}
	})

	t.Run("FontColor", func(t *testing.T) {
		rp := NewRunProperties()
		color := rp.GetOrCreateFontColor()
		color.SetRGB("FF0000")
		if color.RGB() != "FF0000" {
			t.Errorf(
				"expected RGB FF0000, got %s",
				color.RGB(),
			)
		}
	})
}

func TestMergeCells(t *testing.T) {
	t.Run("NewMergeCells", func(t *testing.T) {
		mc := NewMergeCells()
		if mc == nil {
			t.Fatal("expected non-nil MergeCells")
		}
		if mc.LocalName() != "mergeCells" {
			t.Errorf(
				"expected localName mergeCells, got %s",
				mc.LocalName(),
			)
		}
	})

	t.Run("AddMergeCell", func(t *testing.T) {
		mc := NewMergeCells()
		m := mc.AddMergeCell("A1:D1")
		if m == nil {
			t.Fatal("expected non-nil MergeCell")
		}
		if m.Ref() != "A1:D1" {
			t.Errorf(
				"expected ref A1:D1, got %s",
				m.Ref(),
			)
		}
		if mc.Count() != 1 {
			t.Errorf(
				"expected count 1, got %d",
				mc.Count(),
			)
		}
	})

	t.Run("GetMergeCells", func(t *testing.T) {
		mc := NewMergeCells()
		mc.AddMergeCell("A1:B1")
		mc.AddMergeCell("C1:D1")

		count := 0
		for range mc.GetMergeCells() {
			count++
		}
		if count != 2 {
			t.Errorf(
				"expected 2 merge cells, got %d",
				count,
			)
		}
	})

	t.Run(
		"GetMergeCellByRef",
		func(t *testing.T) {
			mc := NewMergeCells()
			mc.AddMergeCell("A1:B1")
			mc.AddMergeCell("C1:D1")

			m := mc.GetMergeCellByRef("C1:D1")
			if m == nil {
				t.Fatal(
					"expected non-nil MergeCell",
				)
			}
			if m.Ref() != "C1:D1" {
				t.Errorf(
					"expected ref C1:D1, got %s",
					m.Ref(),
				)
			}

			m2 := mc.GetMergeCellByRef("E1:F1")
			if m2 != nil {
				t.Error(
					"expected nil for non-existent merge cell",
				)
			}
		},
	)

	t.Run("RemoveMergeCell", func(t *testing.T) {
		mc := NewMergeCells()
		mc.AddMergeCell("A1:B1")
		mc.AddMergeCell("C1:D1")

		if !mc.RemoveMergeCell("A1:B1") {
			t.Error(
				"expected RemoveMergeCell to return true",
			)
		}
		if mc.Count() != 1 {
			t.Errorf(
				"expected count 1, got %d",
				mc.Count(),
			)
		}
		if mc.HasMergeCell("A1:B1") {
			t.Error(
				"expected A1:B1 to be removed",
			)
		}
	})
}

func TestHyperlinks(t *testing.T) {
	t.Run("NewHyperlinks", func(t *testing.T) {
		h := NewHyperlinks()
		if h == nil {
			t.Fatal("expected non-nil Hyperlinks")
		}
		if h.LocalName() != "hyperlinks" {
			t.Errorf(
				"expected localName hyperlinks, got %s",
				h.LocalName(),
			)
		}
	})

	t.Run("AddHyperlink", func(t *testing.T) {
		h := NewHyperlinks()
		hl := h.AddHyperlink("A1", "rId1")
		if hl == nil {
			t.Fatal("expected non-nil Hyperlink")
		}
		if hl.Ref() != "A1" {
			t.Errorf(
				"expected ref A1, got %s",
				hl.Ref(),
			)
		}
		if hl.RelationshipId() != "rId1" {
			t.Errorf(
				"expected r:id rId1, got %s",
				hl.RelationshipId(),
			)
		}
		if !hl.IsExternal() {
			t.Error(
				"expected IsExternal to be true",
			)
		}
	})

	t.Run(
		"AddInternalHyperlink",
		func(t *testing.T) {
			h := NewHyperlinks()
			hl := h.AddInternalHyperlink(
				"B1",
				"Sheet2!A1",
			)
			if hl.Location() != "Sheet2!A1" {
				t.Errorf(
					"expected location Sheet2!A1, got %s",
					hl.Location(),
				)
			}
			if !hl.IsInternal() {
				t.Error(
					"expected IsInternal to be true",
				)
			}
			if hl.IsExternal() {
				t.Error(
					"expected IsExternal to be false",
				)
			}
		},
	)

	t.Run("HyperlinkCount", func(t *testing.T) {
		h := NewHyperlinks()
		h.AddHyperlink("A1", "rId1")
		h.AddHyperlink("B1", "rId2")
		if h.HyperlinkCount() != 2 {
			t.Errorf(
				"expected 2 hyperlinks, got %d",
				h.HyperlinkCount(),
			)
		}
	})

	t.Run(
		"Display and Tooltip",
		func(t *testing.T) {
			hl := NewHyperlink()
			hl.SetDisplay("Click here")
			hl.SetTooltip("Go to website")

			if hl.Display() != "Click here" {
				t.Errorf(
					"expected display 'Click here', got %s",
					hl.Display(),
				)
			}
			if hl.Tooltip() != "Go to website" {
				t.Errorf(
					"expected tooltip 'Go to website', got %s",
					hl.Tooltip(),
				)
			}
		},
	)

	t.Run("RemoveHyperlink", func(t *testing.T) {
		h := NewHyperlinks()
		h.AddHyperlink("A1", "rId1")
		h.AddHyperlink("B1", "rId2")

		if !h.RemoveHyperlink("A1") {
			t.Error(
				"expected RemoveHyperlink to return true",
			)
		}
		if h.HasHyperlink("A1") {
			t.Error(
				"expected A1 hyperlink to be removed",
			)
		}
		if !h.HasHyperlink("B1") {
			t.Error(
				"expected B1 hyperlink to still exist",
			)
		}
	})
}

func TestColumnComparison(t *testing.T) {
	tests := []struct {
		col1     string
		col2     string
		expected int
	}{
		{"A", "A", 0},
		{"A", "B", -1},
		{"B", "A", 1},
		{"A", "Z", -1},
		{"Z", "AA", -1},
		{"AA", "Z", 1},
		{"AA", "AB", -1},
		{"AZ", "BA", -1},
	}

	for _, tt := range tests {
		result := compareColumns(tt.col1, tt.col2)
		var failed bool
		switch {
		case result < 0 && tt.expected >= 0:
			failed = true
		case result > 0 && tt.expected <= 0:
			failed = true
		case result == 0 && tt.expected != 0:
			failed = true
		}
		if failed {
			t.Errorf(
				"compareColumns(%s, %s) = %d, expected %d",
				tt.col1,
				tt.col2,
				result,
				tt.expected,
			)
		}
	}
}

func TestColumnFromReference(t *testing.T) {
	tests := []struct {
		ref      string
		expected string
	}{
		{"A1", "A"},
		{"B10", "B"},
		{"AA100", "AA"},
		{"ZZ999", "ZZ"},
		{"ABC1234", "ABC"},
	}

	for _, tt := range tests {
		result := columnFromReference(tt.ref)
		if result != tt.expected {
			t.Errorf(
				"columnFromReference(%s) = %s, expected %s",
				tt.ref,
				result,
				tt.expected,
			)
		}
	}
}

func TestXMLOutput(t *testing.T) {
	t.Run("SheetData XML", func(t *testing.T) {
		sd := NewSheetData()
		row := sd.AddRow(1)
		cell := row.AddCell("A1")
		cell.SetNumberValue(42)

		xml := sd.OuterXml()
		if !strings.Contains(xml, "sheetData") {
			t.Error(
				"expected XML to contain sheetData",
			)
		}
		if !strings.Contains(xml, ":row") &&
			!strings.Contains(xml, "<row") {
			t.Error(
				"expected XML to contain row element",
			)
		}
		if !strings.Contains(xml, "r=\"1\"") {
			t.Error(
				"expected XML to contain r=\"1\"",
			)
		}
		if !strings.Contains(xml, ":c ") &&
			!strings.Contains(xml, "<c ") {
			t.Error(
				"expected XML to contain c element",
			)
		}
		if !strings.Contains(xml, "r=\"A1\"") {
			t.Error(
				"expected XML to contain r=\"A1\"",
			)
		}
		if !strings.Contains(xml, ">42<") {
			t.Error(
				"expected XML to contain value 42",
			)
		}
	})

	t.Run("MergeCells XML", func(t *testing.T) {
		mc := NewMergeCells()
		mc.AddMergeCell("A1:D1")

		xml := mc.OuterXml()
		if !strings.Contains(xml, "mergeCells") {
			t.Error(
				"expected XML to contain mergeCells",
			)
		}
		if !strings.Contains(xml, "count=\"1\"") {
			t.Error(
				"expected XML to contain count=\"1\"",
			)
		}
		if !strings.Contains(
			xml,
			"ref=\"A1:D1\"",
		) {
			t.Error(
				"expected XML to contain ref=\"A1:D1\"",
			)
		}
	})
}
