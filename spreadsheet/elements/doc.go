//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package elements provides Excel document element types for SpreadsheetML.
//
// This package contains the core elements used in SpreadsheetML documents,
// including workbook structure, worksheets, cells, rows, columns, and
// formatting.
//
// # Workbook Structure
//
// The workbook hierarchy is:
//
//	Workbook
//	  -> Sheets
//	       -> Sheet (reference)
//	  -> DefinedNames
//	       -> DefinedName
//
//	Worksheet
//	  -> SheetData
//	       -> Row
//	            -> Cell
//	  -> MergeCells
//	  -> ConditionalFormatting
//	  -> DataValidations
//
// Creating a workbook structure:
//
//	wb := elements.NewWorkbook()
//	sheets := wb.Sheets()
//	sheet := sheets.AddSheet("Sheet1", 1)
//
// # Worksheets
//
// Worksheets contain the actual data:
//
//	ws := elements.NewWorksheet()
//	sheetData := ws.GetOrCreateSheetData()
//	row := sheetData.GetOrCreateRow(1)
//	cell := row.GetOrCreateCell("A1")
//	cell.SetValue("Hello")
//
// # Rows
//
// Rows contain cells and can have properties:
//
//	row := elements.NewRow(1)
//	row.SetHeight(24.0)
//	row.SetHidden(false)
//	row.SetCustomHeight(true)
//
// Iterate over cells:
//
//	for cell := range row.Cells() {
//		fmt.Println(cell.Reference(), cell.Value())
//	}
//
// # Cells
//
// Cells hold values, formulas, and formatting:
//
//	cell := elements.NewCell("A1")
//	cell.SetValue("Text")
//	cell.SetNumberValue(42.5)
//	cell.SetBoolValue(true)
//	cell.SetFormula("=SUM(B1:B10)")
//	cell.SetStyleIndex(1)
//
// Cell types:
//
//	elements.CellTypeString     // String value
//	elements.CellTypeNumber     // Numeric value
//	elements.CellTypeBoolean    // Boolean value
//	elements.CellTypeError      // Error value
//	elements.CellTypeFormula    // Formula
//	elements.CellTypeSharedString // Shared string reference
//
// # Shared Strings
//
// Shared strings optimize storage of repeated text:
//
//	sst := elements.NewSharedStringTable()
//	index := sst.AddString("Hello")
//	cell.SetSharedStringIndex(index)
//
// Rich text in shared strings:
//
//	si := elements.NewSharedStringItem()
//	run := si.AddRun()
//	run.SetText("Bold")
//	run.SetBold(true)
//
// # Styles
//
// Style elements define cell formatting:
//
//	stylesheet := elements.NewStylesheet()
//
//	// Fonts
//	font := elements.NewFont()
//	font.SetBold(true)
//	font.SetSize(12)
//	font.SetColor("FF0000")
//	stylesheet.AddFont(font)
//
//	// Fills
//	fill := elements.NewFill()
//	fill.SetPatternType(elements.PatternSolid)
//	fill.SetForegroundColor("FFFF00")
//	stylesheet.AddFill(fill)
//
//	// Borders
//	border := elements.NewBorder()
//	border.SetLeftStyle(elements.BorderThin)
//	border.SetRightStyle(elements.BorderThin)
//	stylesheet.AddBorder(border)
//
//	// Cell formats (xf)
//	xf := elements.NewCellFormat()
//	xf.SetFontId(0)
//	xf.SetFillId(1)
//	xf.SetBorderId(1)
//	xf.SetNumberFormatId(1)
//	stylesheet.AddCellFormat(xf)
//
// # Number Formats
//
// Custom number formats:
//
//	numFmt := elements.NewNumberFormat(164, "#,##0.00")
//	stylesheet.AddNumberFormat(numFmt)
//
// Built-in format IDs:
//
//	0  - General
//	1  - 0
//	2  - 0.00
//	9  - 0%
//	10 - 0.00%
//	14 - m/d/yyyy
//	22 - m/d/yyyy h:mm
//
// # Merge Cells
//
// Merge cell ranges:
//
//	mergeCells := ws.GetOrCreateMergeCells()
//	mergeCells.AddMergeCell("A1:D1")
//
// # Conditional Formatting
//
// Apply conditional formatting rules:
//
//	cf := elements.NewConditionalFormatting()
//	cf.SetRange("A1:A10")
//	rule := cf.AddRule(elements.CFRuleCellIs)
//	rule.SetOperator(elements.CFOperatorGreaterThan)
//	rule.SetFormula("100")
//	rule.SetStyleId(1)
//
// # Data Validations
//
// Add data validation rules:
//
//	dv := elements.NewDataValidation()
//	dv.SetRange("B1:B100")
//	dv.SetType(elements.DVTypeList)
//	dv.SetFormula1("Sheet2!$A$1:$A$10")
//
// # Defined Names
//
// Create named ranges:
//
//	dn := elements.NewDefinedName()
//	dn.SetName("MyRange")
//	dn.SetFormula("Sheet1!$A$1:$D$10")
//
// # Charts
//
// Chart elements (via DrawingML):
//
//	chart := elements.NewChart()
//	chart.SetTitle("Sales Data")
//	chart.SetType(elements.ChartTypeBar)
//
// # Measurement Units
//
// The package uses standard OOXML units for spreadsheets:
//
//   - Column width: Character units (based on default font)
//   - Row height: Points (1/72 inch)
//   - EMUs: English Metric Units for drawings (914400 EMU = 1 inch)
//
// # Namespaces
//
// Standard SpreadsheetML namespace:
//
//	elements.NamespaceSML  // spreadsheetml/2006/main namespace
//	                       // (http://schemas.openxmlformats.org/...)
//	elements.PrefixDefault // "" (no prefix, default namespace)
//
// # Iterator Pattern
//
// Collections use Go 1.25 iterator pattern (iter.Seq):
//
//	for row := range sheetData.Rows() {
//		for cell := range row.Cells() {
//			fmt.Println(cell.Value())
//		}
//	}
//
// # Cloning Elements
//
// All elements support cloning:
//
//	clone := cell.Clone().(*elements.Cell)
//	deepClone := row.CloneNode(true)
package elements
