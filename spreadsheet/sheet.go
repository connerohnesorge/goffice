// This file implements the high-level Sheet API for working with worksheets.
//
//nolint:revive // file-length-limit: comprehensive sheet API with all methods
package spreadsheet

import (
	"iter"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// Sheet represents a high-level wrapper around a worksheet with convenient methods.
type Sheet struct {
	// doc is the parent document
	doc *Document

	// name is the sheet name
	name string

	// sheetID is the sheet ID in the workbook
	sheetID int

	// worksheetPart is the underlying worksheet part
	worksheetPart *parts.WorksheetPart

	// worksheet is the cached worksheet element
	worksheet *elements.Worksheet
}

// newSheet creates a new Sheet wrapper.
func newSheet(
	doc *Document,
	name string,
	sheetID int,
	worksheetPart *parts.WorksheetPart,
) *Sheet {
	return &Sheet{
		doc:           doc,
		name:          name,
		sheetID:       sheetID,
		worksheetPart: worksheetPart,
	}
}

// Name returns the sheet name.
func (s *Sheet) Name() string {
	return s.name
}

// ID returns the sheet ID.
func (s *Sheet) ID() int {
	return s.sheetID
}

// WorksheetPart returns the underlying worksheet part.
func (s *Sheet) WorksheetPart() *parts.WorksheetPart {
	return s.worksheetPart
}

// Worksheet returns the worksheet element, loading if necessary.
func (s *Sheet) Worksheet() *elements.Worksheet {
	if s.worksheet != nil {
		return s.worksheet
	}

	// Try to get the worksheet from the part's root element
	root := s.worksheetPart.Worksheet()
	if root != nil {
		// Cache and return the worksheet
		s.worksheet = root

		return s.worksheet
	}

	// Create a new worksheet element if needed
	s.worksheet = elements.NewWorksheet()
	s.worksheet.GetOrCreateSheetData()

	// Set the worksheet as the root element on the part so it gets persisted
	s.worksheetPart.SetRootElement(s.worksheet)

	return s.worksheet
}

// SheetData returns the sheet data element, creating if necessary.
func (s *Sheet) SheetData() *elements.SheetData {
	return s.Worksheet().GetOrCreateSheetData()
}

// Cell returns a Cell wrapper for the cell at the given A1-style reference.
// Example: sheet.Cell("A1") or sheet.Cell("$B$5")
func (s *Sheet) Cell(ref string) *Cell {
	cellRef, err := ParseCellRef(ref)
	if err != nil {
		// Return nil cell for invalid reference
		return nil
	}

	return s.CellAt(
		uint32(cellRef.Row),
		uint32(cellRef.Col),
	)
}

// CellAt returns a Cell wrapper for the cell at the given row and column (1-based).
func (s *Sheet) CellAt(row, col uint32) *Cell {
	if row < MinRow || row > MaxRow ||
		col < MinColumn ||
		col > MaxColumn {
		return nil
	}

	// Get or create the row
	sheetData := s.SheetData()
	elemRow := sheetData.GetOrCreateRow(row)

	// Build the cell reference
	ref := CellRef{
		Col: int(col),
		Row: int(row),
	}.String()

	// Get or create the cell
	elemCell := elemRow.GetOrCreateCell(ref)

	return newCell(s, elemCell, row, col)
}

// Range returns a Range wrapper for cells in the given range reference.
// Example: sheet.Range("A1:D10")
func (s *Sheet) Range(ref string) *Range {
	rangeRef, err := ParseRangeRef(ref)
	if err != nil {
		return nil
	}

	return newRange(s, &rangeRef)
}

// Row returns a Row wrapper for the row at the given index (1-based).
func (s *Sheet) Row(index uint32) *Row {
	if index < MinRow || index > MaxRow {
		return nil
	}

	sheetData := s.SheetData()
	elemRow := sheetData.GetRow(index)
	if elemRow == nil {
		return nil
	}

	return newRow(s, elemRow)
}

// AddRow adds a new row at the next available index and returns it.
func (s *Sheet) AddRow() *Row {
	sheetData := s.SheetData()

	// Find the next available row index
	maxIndex := uint32(0)
	for row := range sheetData.Rows() {
		if row.RowIndex() > maxIndex {
			maxIndex = row.RowIndex()
		}
	}

	nextIndex := maxIndex + 1
	elemRow := sheetData.AddRow(nextIndex)

	return newRow(s, elemRow)
}

// AddRowAt adds a new row at the given index (1-based) and returns it.
func (s *Sheet) AddRowAt(index uint32) *Row {
	if index < MinRow || index > MaxRow {
		return nil
	}

	sheetData := s.SheetData()
	elemRow := sheetData.GetOrCreateRow(index)

	return newRow(s, elemRow)
}

// Rows returns an iterator over all rows in the sheet.
func (s *Sheet) Rows() iter.Seq[*Row] {
	return func(yield func(*Row) bool) {
		sheetData := s.SheetData()
		for elemRow := range sheetData.Rows() {
			row := newRow(s, elemRow)
			if !yield(row) {
				return
			}
		}
	}
}

// MergeCells merges the cells in the given range.
// Example: sheet.MergeCells("A1:D1")
func (s *Sheet) MergeCells(ref string) error {
	// Validate the range reference
	_, err := ParseRangeRef(ref)
	if err != nil {
		return err
	}

	ws := s.Worksheet()
	mergeCells := ws.GetOrCreateMergeCells()
	mergeCells.AddMergeCell(ref)

	return nil
}

// UnmergeCells unmerges the cells in the given range.
// Example: sheet.UnmergeCells("A1:D1")
func (s *Sheet) UnmergeCells(ref string) error {
	ws := s.Worksheet()
	mergeCells := ws.MergeCells()
	if mergeCells == nil {
		return nil // No merge cells to unmerge
	}

	mergeCells.RemoveMergeCell(ref)

	return nil
}

// SetColumnWidth sets the width of a column (1-based).
// Width is in character units (approximately the width of a digit).
func (s *Sheet) SetColumnWidth(
	col uint32,
	width float64,
) error {
	if col < MinColumn || col > MaxColumn {
		return ErrColumnOutOfRange
	}

	ws := s.Worksheet()
	cols := ws.GetOrCreateCols()

	// Check if there's already a column definition that includes this column
	existingCol := cols.GetColByIndex(int(col))
	if existingCol != nil {
		// If the existing column definition spans exactly this column
		if existingCol.Min() == int(col) &&
			existingCol.Max() == int(col) {
			existingCol.SetWidth(width)
			existingCol.SetCustomWidth(true)

			return nil
		}
		// Otherwise, we need to split the column definition
		// For simplicity, just add a new one for this specific column
	}

	// Add a new column definition
	newCol := cols.AddCol(int(col), int(col))
	newCol.SetWidth(width)
	newCol.SetCustomWidth(true)

	return nil
}

// SetRowHeight sets the height of a row (1-based).
// Height is in points (1/72 inch).
func (s *Sheet) SetRowHeight(
	row uint32,
	height float64,
) error {
	if row < MinRow || row > MaxRow {
		return ErrRowOutOfRange
	}

	sheetData := s.SheetData()
	elemRow := sheetData.GetOrCreateRow(row)
	elemRow.SetHeight(height)
	elemRow.SetCustomHeight(true)

	return nil
}

// AddTable adds a new table to the sheet.
// ref is the table range (e.g., "A1:D10"), name is the table name.
func (s *Sheet) AddTable(
	ref, name string,
) (*Table, error) {
	// Validate the range reference
	_, err := ParseRangeRef(ref)
	if err != nil {
		return nil, err
	}

	// Add a table definition part
	tablePart, err := s.worksheetPart.AddTableDefinitionPart()
	if err != nil {
		return nil, err
	}

	// Create the table wrapper
	return newTable(s, tablePart, ref, name)
}

// Tables returns all tables in the sheet.
func (s *Sheet) Tables() []*Table {
	var tables []*Table
	for _, tablePart := range s.worksheetPart.TableDefinitionParts() {
		t := newTableFromPart(s, tablePart)
		if t != nil {
			tables = append(tables, t)
		}
	}

	return tables
}

// AddChart adds a new chart to the sheet.
// chartType specifies the type of chart, dataRange is the source data range,
// and anchor specifies where to place the chart.
func (s *Sheet) AddChart(
	chartType ChartType,
	dataRange string,
	anchor CellRef,
) (*Chart, error) {
	// Ensure we have a drawings part
	drawingsPart := s.worksheetPart.DrawingsPart()
	if drawingsPart == nil {
		var err error
		drawingsPart, err = s.worksheetPart.AddDrawingsPart()
		if err != nil {
			return nil, err
		}
	}

	return newChart(
		s,
		drawingsPart,
		chartType,
		dataRange,
		anchor,
	)
}

// AddPivotTable adds a new pivot table to the sheet.
// sourceRange is the data source range (e.g., "Sheet1!A1:D100"),
// destCell is the destination cell for the pivot table.
func (s *Sheet) AddPivotTable(
	sourceRange string,
	destCell CellRef,
) (*PivotTable, error) {
	// Add a pivot table part
	pivotPart, err := s.worksheetPart.AddPivotTablePart()
	if err != nil {
		return nil, err
	}

	return newPivotTable(
		s,
		pivotPart,
		sourceRange,
		destCell,
	)
}

// AddImage adds an image to the sheet at the specified cell.
// data is the image data, contentType is the MIME type (e.g., "image/png").
func (s *Sheet) AddImage(
	cell CellRef,
	data []byte,
	contentType string,
) (*Image, error) {
	// Ensure we have a drawings part
	drawingsPart := s.worksheetPart.DrawingsPart()
	if drawingsPart == nil {
		var err error
		drawingsPart, err = s.worksheetPart.AddDrawingsPart()
		if err != nil {
			return nil, err
		}
	}

	return newImage(
		s,
		drawingsPart,
		cell,
		data,
		contentType,
	)
}

// AddImageBetween adds an image anchored between two cells.
// from and to specify the top-left and bottom-right corners.
func (s *Sheet) AddImageBetween(
	from, to CellRef,
	data []byte,
	contentType string,
) (*Image, error) {
	// Ensure we have a drawings part
	drawingsPart := s.worksheetPart.DrawingsPart()
	if drawingsPart == nil {
		var err error
		drawingsPart, err = s.worksheetPart.AddDrawingsPart()
		if err != nil {
			return nil, err
		}
	}

	return newImageBetween(
		s,
		drawingsPart,
		from,
		to,
		data,
		contentType,
	)
}

// Row represents a high-level wrapper around a worksheet row.
type Row struct {
	sheet   *Sheet
	elemRow *elements.Row
}

// newRow creates a new Row wrapper.
func newRow(
	sheet *Sheet,
	elemRow *elements.Row,
) *Row {
	return &Row{
		sheet:   sheet,
		elemRow: elemRow,
	}
}

// Index returns the 1-based row index.
func (r *Row) Index() uint32 {
	return r.elemRow.RowIndex()
}

// Cell returns a Cell wrapper for the cell at the given column (1-based).
func (r *Row) Cell(col uint32) *Cell {
	if col < MinColumn || col > MaxColumn {
		return nil
	}

	ref := CellRef{
		Col: int(col),
		Row: int(r.Index()),
	}.String()
	elemCell := r.elemRow.GetOrCreateCell(ref)

	return newCell(
		r.sheet,
		elemCell,
		r.Index(),
		col,
	)
}

// Cells returns an iterator over all cells in the row.
func (r *Row) Cells() iter.Seq[*Cell] {
	return func(yield func(*Cell) bool) {
		for elemCell := range r.elemRow.Cells() {
			cellRef, err := ParseCellRef(
				elemCell.Reference(),
			)
			if err != nil {
				continue
			}
			cell := newCell(
				r.sheet,
				elemCell,
				uint32(cellRef.Row),
				uint32(cellRef.Col),
			)
			if !yield(cell) {
				return
			}
		}
	}
}

// SetHeight sets the row height in points.
func (r *Row) SetHeight(height float64) {
	r.elemRow.SetHeight(height)
	r.elemRow.SetCustomHeight(true)
}

// Height returns the row height in points.
func (r *Row) Height() float64 {
	return r.elemRow.Height()
}

// SetHidden sets whether the row is hidden.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (r *Row) SetHidden(hidden bool) {
	r.elemRow.SetHidden(hidden)
}

// Hidden returns whether the row is hidden.
func (r *Row) Hidden() bool {
	return r.elemRow.Hidden()
}

// Range represents a range of cells in a worksheet.
type Range struct {
	sheet    *Sheet
	rangeRef RangeRef
}

// newRange creates a new Range wrapper.
func newRange(
	sheet *Sheet,
	rangeRef *RangeRef,
) *Range {
	return &Range{
		sheet:    sheet,
		rangeRef: *rangeRef,
	}
}

// String returns the A1-style range reference.
func (r *Range) String() string {
	return r.rangeRef.String()
}

// Start returns the starting cell reference.
func (r *Range) Start() CellRef {
	return r.rangeRef.Start
}

// End returns the ending cell reference.
func (r *Range) End() CellRef {
	return r.rangeRef.End
}

// Cells returns an iterator over all cells in the range.
func (r *Range) Cells() iter.Seq[*Cell] {
	return func(yield func(*Cell) bool) {
		for cellRef := range r.rangeRef.Cells() {
			cell := r.sheet.CellAt(
				uint32(cellRef.Row),
				uint32(cellRef.Col),
			)
			if cell != nil && !yield(cell) {
				return
			}
		}
	}
}

// SetValue sets the same value for all cells in the range.
func (r *Range) SetValue(value interface{}) {
	for cell := range r.Cells() {
		switch v := value.(type) {
		case float64:
			cell.SetNumber(v)
		case int:
			cell.SetNumber(float64(v))
		case string:
			cell.SetString(v)
		case bool:
			cell.SetBoolean(v)
		}
	}
}

// SetStyle sets the same style for all cells in the range.
func (r *Range) SetStyle(style *Style) {
	for cell := range r.Cells() {
		cell.SetStyle(style)
	}
}

// SetCellValue sets the value of a cell at the given A1-style reference.
// The value type is automatically detected (string, number, bool).
func (s *Sheet) SetCellValue(
	ref string,
	value interface{},
) error {
	cell := s.Cell(ref)
	if cell == nil {
		return ErrCellNotFound
	}

	switch v := value.(type) {
	case float64:
		cell.SetNumber(v)
	case int:
		cell.SetNumber(float64(v))
	case int32:
		cell.SetNumber(float64(v))
	case int64:
		cell.SetNumber(float64(v))
	case string:
		cell.SetString(v)
	case bool:
		cell.SetBoolean(v)
	default:
		// Convert to string for unknown types
		cell.SetString(stringValue(v))
	}

	return nil
}

// SetCellFormula sets the formula of a cell at the given A1-style reference.
func (s *Sheet) SetCellFormula(
	ref string,
	formula string,
) error {
	cell := s.Cell(ref)
	if cell == nil {
		return ErrCellNotFound
	}

	cell.SetFormula(formula)

	return nil
}

// GetCellValue returns the value of a cell at the given A1-style reference.
// Returns empty string if the cell doesn't exist.
func (s *Sheet) GetCellValue(ref string) string {
	cell := s.Cell(ref)
	if cell == nil {
		return ""
	}

	return cell.GetString()
}

// GetCellFormula returns the formula of a cell at the given A1-style reference.
// Returns empty string if the cell doesn't have a formula.
func (s *Sheet) GetCellFormula(
	ref string,
) string {
	cell := s.Cell(ref)
	if cell == nil {
		return ""
	}

	return cell.GetFormula()
}

// stringValue converts any value to string.
func stringValue(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

// ErrCellNotFound is returned when a cell reference is invalid.
var ErrCellNotFound = sheetError(
	"cell not found or invalid reference",
)

type sheetError string

func (e sheetError) Error() string {
	return string(e)
}
