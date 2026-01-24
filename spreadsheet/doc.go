//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
// This package implements the document-level API for creating, reading, and
// modifying .xlsx files (and related formats like .xlsm, .xltx, .xltm, .xlam).
//
// # Creating Workbooks
//
// Create a new workbook at a file path:
//
//	wb, err := spreadsheet.New("workbook.xlsx", spreadsheet.DocTypeWorkbook)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer wb.Close()
//
// Create different document types:
//
//	wb, _ := spreadsheet.New("template.xltx", spreadsheet.DocTypeTemplate)
//	wb, _ := spreadsheet.New("macro.xlsm", spreadsheet.DocTypeMacroEnabled)
//
// Write to an io.Writer:
//
//	var buf bytes.Buffer
//	wb, err := spreadsheet.NewWriter(&buf, spreadsheet.DocTypeWorkbook)
//
// # Opening Workbooks
//
// Open an existing workbook:
//
//	wb, err := spreadsheet.Open("workbook.xlsx", true) // editable=true
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer wb.Close()
//
// Open read-only:
//
//	wb, err := spreadsheet.Open("workbook.xlsx", false)
//
// Open from an io.ReaderAt:
//
//	wb, err := spreadsheet.OpenReader(reader, size, true)
//
// # Working with Worksheets
//
// Access and create worksheets:
//
//	// Get the first worksheet
//	sheet := wb.Worksheet(0)
//
//	// Add a new worksheet
//	sheet, err := wb.AddWorksheet("Sheet2")
//
//	// Iterate over worksheets
//	for sheet := range wb.Worksheets() {
//		fmt.Println(sheet.Name())
//	}
//
// # Working with Cells
//
// Access and modify cell values:
//
//	sheet.SetCellValue("A1", "Hello")
//	sheet.SetCellValue("B1", 42)
//	sheet.SetCellValue("C1", 3.14159)
//	sheet.SetCellValue("D1", true)
//
//	value := sheet.CellValue("A1")
//
// Using cell references:
//
//	cell := sheet.Cell("A1")
//	cell.SetValue("Hello")
//	cell.SetFormula("=SUM(B1:B10)")
//	cell.SetStyle(styleId)
//
// # Rows and Columns
//
// Work with rows and columns:
//
//	row := sheet.Row(1)
//	row.SetHeight(24.0)
//
//	col := sheet.Column("A")
//	col.SetWidth(15.0)
//	col.SetHidden(true)
//
// # Formulas
//
// Set cell formulas:
//
//	sheet.SetCellFormula("E1", "=SUM(A1:D1)")
//	sheet.SetCellFormula("F1", "=AVERAGE(A1:D1)")
//
// # Styles
//
// Apply cell formatting:
//
//	style := wb.CreateStyle()
//	style.SetFontBold(true)
//	style.SetFontSize(14)
//	style.SetFillColor("FFFF00")
//	style.SetBorder(spreadsheet.BorderThin)
//
//	styleId := wb.AddStyle(style)
//	sheet.SetCellStyle("A1", styleId)
//
// # Number Formats
//
// Apply number formatting:
//
//	style := wb.CreateStyle()
//	style.SetNumberFormat("#,##0.00")
//	style.SetNumberFormat("$#,##0.00")
//	style.SetNumberFormat("0%")
//	style.SetNumberFormat("yyyy-mm-dd")
//
// # Saving Workbooks
//
// Save to the original location:
//
//	err := wb.Save()
//
// Save to a new location:
//
//	err := wb.SaveAs("copy.xlsx")
//
// # Document Types
//
// The package supports different Excel document types:
//
//   - DocTypeWorkbook: Standard workbook (.xlsx)
//   - DocTypeTemplate: Template (.xltx)
//   - DocTypeMacroEnabled: Macro-enabled workbook (.xlsm)
//   - DocTypeMacroTemplate: Macro-enabled template (.xltm)
//   - DocTypeAddIn: Add-in (.xlam)
//
// Change document type:
//
//	err := wb.ChangeType(spreadsheet.DocTypeTemplate)
//
// # Workbook Settings
//
// Configure workbook open settings:
//
//	settings := &spreadsheet.OpenSettings{
//		AutoSave: true,
//		MaxCharactersInPart: 0, // unlimited
//	}
//	wb, err := spreadsheet.OpenWithSettings("workbook.xlsx", true, settings)
//
// # Thread Safety
//
// Workbook operations are protected by sync.RWMutex. Read operations can
// proceed concurrently, while write operations require exclusive access.
//
// # Subpackages
//
// The spreadsheet package is organized into subpackages:
//
//   - spreadsheet/elements: SpreadsheetML element types
//   - spreadsheet/parts: Workbook part types (WorkbookPart, WorksheetPart, etc.)
//
// See those packages for lower-level element manipulation.
//
//nolint:revive // line-length-limit: documentation examples contain long code samples
package spreadsheet
