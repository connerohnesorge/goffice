//go:generate gomarkdoc -u -o CLAUDE.md .

//go:generate gomarkdoc -u -o AGENTS.md .

// Package parts provides Excel document part types.
//
// Parts are the individual XML files within an .xlsx package. Each part
// represents a distinct aspect of the workbook such as the workbook itself,
// worksheets, styles, shared strings, and more.
//
// # Part Types
//
// The package provides the following part types:
//
// Workbook Structure:
//   - WorkbookPart: Workbook definition (xl/workbook.xml)
//   - WorksheetPart: Worksheet data (xl/worksheets/sheet1.xml, etc.)
//   - ChartsheetPart: Chart sheet (xl/chartsheets/sheet1.xml, etc.)
//
// Styles and Formatting:
//   - StylesPart: Style definitions (xl/styles.xml)
//   - ThemePart: Theme definitions (xl/theme/theme1.xml)
//
// Data:
//   - SharedStringsPart: Shared string table (xl/sharedStrings.xml)
//   - TablePart: Table definitions (xl/tables/table1.xml, etc.)
//   - PivotTablePart: Pivot table definitions (xl/pivotTables/pivotTable1.xml)
//   - PivotCachePart: Pivot cache (xl/pivotCache/pivotCacheDefinition1.xml)
//
// Settings:
//   - CalcChainPart: Calculation chain (xl/calcChain.xml)
//   - WorkbookSettingsPart: Workbook settings
//
// Media and Extensions:
//   - DrawingPart: Drawings (xl/drawings/drawing1.xml)
//   - ChartPart: Charts (xl/charts/chart1.xml)
//   - ImagePart: Embedded images (xl/media/image1.png, etc.)
//   - VbaProjectPart: VBA macros (xl/vbaProject.bin)
//   - CustomXmlPart: Custom XML data (customXml/item1.xml, etc.)
//
// # WorkbookPart
//
// The WorkbookPart is the central part of any Excel document:
//
//	wbPart := wb.WorkbookPart()
//	workbook := wbPart.Workbook()
//	sheets := workbook.Sheets()
//
// Add supporting parts through WorkbookPart:
//
//	stylesPart, err := wbPart.AddStylesPart()
//	ssPart, err := wbPart.AddSharedStringsPart()
//
// # Worksheets
//
// Add and access worksheet parts:
//
//	wsPart, err := wbPart.AddWorksheetPart()
//	worksheet := wsPart.Worksheet()
//	sheetData := worksheet.GetOrCreateSheetData()
//
// Access existing worksheets:
//
//	for _, wsp := range wbPart.WorksheetParts() {
//		ws := wsp.Worksheet()
//		// process worksheet
//	}
//
// # Styles Part
//
// Access and modify workbook styles:
//
//	stylesPart, err := wbPart.AddStylesPart()
//	stylesheet := stylesPart.Stylesheet()
//
//	// Add a font
//	font := elements.NewFont()
//	font.SetBold(true)
//	stylesheet.AddFont(font)
//
//	// Add a cell format
//	xf := elements.NewCellFormat()
//	xf.SetFontId(0)
//	stylesheet.AddCellFormat(xf)
//
// # Shared Strings Part
//
// Manage shared strings for text optimization:
//
//	ssPart, err := wbPart.AddSharedStringsPart()
//	sst := ssPart.SharedStringTable()
//
//	// Add a string and get its index
//	index := sst.AddString("Hello")
//
//	// Use index in cell
//	cell.SetSharedStringIndex(index)
//
// # Drawings
//
// Add drawings for charts and images:
//
//	drawingPart, err := wsPart.AddDrawingPart()
//	drawing := drawingPart.Drawing()
//
//	// Add chart
//	chartPart, err := drawingPart.AddChartPart()
//	anchor := drawing.AddTwoCellAnchor("A1", "F10")
//	anchor.SetChartRelId(chartPart.RelationshipId())
//
// # Images
//
// Add image parts:
//
//	imagePart, relId, err := drawingPart.AddImagePart("image/png")
//	imagePart.SetData(imageBytes)
//	// Use relId in drawing anchor
//
// # Tables
//
// Add table definitions:
//
//	tablePart, err := wsPart.AddTablePart()
//	table := tablePart.Table()
//	table.SetRef("A1:D10")
//	table.SetDisplayName("SalesData")
//
//	// Add table columns
//	table.AddColumn("Name")
//	table.AddColumn("Amount")
//
// # Part URIs
//
// Standard part URIs follow OPC conventions:
//
//   - /xl/workbook.xml - Workbook
//   - /xl/worksheets/sheet1.xml - First worksheet
//   - /xl/styles.xml - Styles
//   - /xl/sharedStrings.xml - Shared strings
//   - /xl/theme/theme1.xml - Theme
//   - /xl/drawings/drawing1.xml - First drawing
//   - /xl/charts/chart1.xml - First chart
//   - /xl/tables/table1.xml - First table
//   - /xl/media/image1.png - First image
//
// # Content Types
//
// Each part type has a specific content type:
//
//	parts.ContentTypeWorkbook      // Workbook
//	parts.ContentTypeWorksheet     // Worksheet
//	parts.ContentTypeStyles        // Styles
//	parts.ContentTypeSharedStrings // Shared strings
//	parts.ContentTypeChart         // Charts
//	parts.ContentTypeTable         // Tables
//
// # Relationship Types
//
// Parts are connected via relationships:
//
//	parts.RelationshipTypeWorkbook      // Workbook
//	parts.RelationshipTypeWorksheet     // Worksheet
//	parts.RelationshipTypeStyles        // Styles
//	parts.RelationshipTypeSharedStrings // Shared strings
//	parts.RelationshipTypeTheme         // Theme
//	parts.RelationshipTypeDrawing       // Drawing
//	parts.RelationshipTypeChart         // Chart
//
// # Part Registration
//
// All part types are automatically registered with the openxml package's
// part type registry via init() functions. This enables automatic part
// instantiation when opening existing documents.
//
// # Thread Safety
//
// Part operations should be performed through the parent workbook's
// thread-safe API. Direct part manipulation is not thread-safe.
//
//nolint:revive // line-length-limit: documentation examples contain long method chains
package parts
