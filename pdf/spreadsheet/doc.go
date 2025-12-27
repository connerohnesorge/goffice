// Package spreadsheet provides PDF rendering for SpreadsheetML (XLSX)
// documents.
//
// This package converts parsed Excel workbooks from the goffice/spreadsheet
// package into high-fidelity PDF output. It handles all Excel-specific layout
// and formatting:
//
//   - Cell formatting (alignment, borders, fills, number formats)
//   - Row and column sizing (auto-fit, manual widths)
//   - Merged cells and cell spanning
//   - Conditional formatting (data bars, color scales, icons)
//   - Charts (bar, line, pie, scatter, etc.)
//   - Images and shapes
//   - Print areas and page breaks
//   - Headers and footers with page numbers
//   - Frozen panes and split views (for visual reference)
//
// # Usage
//
//	import (
//		"github.com/connerohnesorge/goffice/spreadsheet"
//		pdfspreadsheet "github.com/connerohnesorge/goffice-pdf/spreadsheet"
//	)
//
//	workbook, err := spreadsheet.Open("workbook.xlsx")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	renderer := pdfspreadsheet.NewRenderer()
//	err = renderer.Render(workbook, "output.pdf")
//
// # Multi-Sheet Rendering
//
// By default, each worksheet becomes a separate page (or set of pages) in the
// PDF. Options are available to render only specific sheets or to control
// page breaking within large worksheets.
//
// # Print Settings
//
// The renderer respects Excel's print settings including:
//   - Print area
//   - Page orientation and paper size
//   - Scaling (fit to page, percentage)
//   - Margins
//   - Gridlines and headings visibility
package spreadsheet
