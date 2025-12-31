// Package spreadsheet provides SpreadsheetML support for Excel documents.
//
//nolint:revive // file-length-limit: comprehensive sheet operations with all methods
package spreadsheet

import (
	"errors"
	"fmt"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
)

// InsertRows inserts count blank rows starting at index (1-based).
// Rows at index and below shift down. All references in workbook are updated.
// Returns an error if index or count are out of valid range.
func (s *Sheet) InsertRows(
	index, count uint32,
) error {
	// Validate parameters
	if index < MinRow || index > MaxRow {
		return ErrRowOutOfRange
	}
	if count == 0 {
		return errors.New(
			"count must be greater than zero",
		)
	}
	if index+count-1 > MaxRow {
		return errors.New(
			"inserting rows would exceed maximum row count",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// Perform physical row manipulation
	if err := s.physicalInsertRows(sheetData, index, count); err != nil {
		return err
	}

	// Update all references across the workbook
	operation := ShiftOperation{
		Type:      ShiftInsertRows,
		Index:     index,
		Count:     count,
		Worksheet: s.name,
	}

	return s.updateAllReferences(operation)
}

// DeleteRows deletes count rows starting at index (1-based).
// Rows below shift up. References to deleted cells become #REF!.
// Returns an error if index or count are out of valid range.
func (s *Sheet) DeleteRows(
	index, count uint32,
) error {
	// Validate parameters
	if index < MinRow || index > MaxRow {
		return ErrRowOutOfRange
	}
	if count == 0 {
		return errors.New(
			"count must be greater than zero",
		)
	}
	if index+count-1 > MaxRow {
		return errors.New(
			"deletion range exceeds maximum row count",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// Perform physical row manipulation
	if err := s.physicalDeleteRows(sheetData, index, count); err != nil {
		return err
	}

	// Update all references across the workbook
	operation := ShiftOperation{
		Type:      ShiftDeleteRows,
		Index:     index,
		Count:     count,
		Worksheet: s.name,
	}

	return s.updateAllReferences(operation)
}

// InsertColumns inserts count blank columns starting at index (1-based).
// Columns at index and right shift right. All references in workbook are updated.
// Returns an error if index or count are out of valid range.
func (s *Sheet) InsertColumns(
	index, count uint32,
) error {
	// Validate parameters
	if index < MinColumn || index > MaxColumn {
		return ErrColumnOutOfRange
	}
	if count == 0 {
		return errors.New(
			"count must be greater than zero",
		)
	}
	if index+count-1 > MaxColumn {
		return errors.New(
			"inserting columns would exceed maximum column count",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// Perform physical column manipulation
	if err := s.physicalInsertColumns(sheetData, index, count); err != nil {
		return err
	}

	// Update all references across the workbook
	operation := ShiftOperation{
		Type:      ShiftInsertColumns,
		Index:     index,
		Count:     count,
		Worksheet: s.name,
	}

	return s.updateAllReferences(operation)
}

// DeleteColumns deletes count columns starting at index (1-based).
// Columns right shift left. References to deleted cells become #REF!.
// Returns an error if index or count are out of valid range.
func (s *Sheet) DeleteColumns(
	index, count uint32,
) error {
	// Validate parameters
	if index < MinColumn || index > MaxColumn {
		return ErrColumnOutOfRange
	}
	if count == 0 {
		return errors.New(
			"count must be greater than zero",
		)
	}
	if index+count-1 > MaxColumn {
		return errors.New(
			"deletion range exceeds maximum column count",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// Perform physical column manipulation
	if err := s.physicalDeleteColumns(sheetData, index, count); err != nil {
		return err
	}

	// Update all references across the workbook
	operation := ShiftOperation{
		Type:      ShiftDeleteColumns,
		Index:     index,
		Count:     count,
		Worksheet: s.name,
	}

	return s.updateAllReferences(operation)
}

// physicalInsertRows shifts existing rows down to make space for new rows.
func (s *Sheet) physicalInsertRows(
	sheetData *elements.SheetData,
	index, count uint32,
) error {
	// Collect all rows that need to be shifted (in reverse order to avoid conflicts)
	var rowsToShift []*elements.Row
	for row := range sheetData.Rows() {
		if row.RowIndex() >= index {
			rowsToShift = append(rowsToShift, row)
		}
	}

	// Process rows in reverse order (highest row number first)
	// This avoids overwriting rows that haven't been processed yet
	for i := len(rowsToShift) - 1; i >= 0; i-- {
		row := rowsToShift[i]
		oldIndex := row.RowIndex()
		newIndex := oldIndex + count

		if newIndex > MaxRow {
			// Skip rows that would exceed maximum
			continue
		}

		// Update row index
		row.SetRowIndex(newIndex)

		// Update all cell references in this row
		for cell := range row.Cells() {
			ref := cell.Reference()
			if ref == "" {
				continue
			}

			cellRef, err := ParseCellRef(ref)
			if err != nil {
				continue
			}

			// Update row number
			cellRef.Row = int(newIndex)
			cell.SetReference(cellRef.String())
		}
	}

	return nil
}

// physicalDeleteRows removes rows and shifts remaining rows up.
func (s *Sheet) physicalDeleteRows(
	sheetData *elements.SheetData,
	index, count uint32,
) error {
	deleteEnd := index + count - 1

	// Remove rows in deletion range
	for i := index; i <= deleteEnd && i <= MaxRow; i++ {
		sheetData.RemoveRow(i)
	}

	// Collect all rows that need to be shifted up
	var rowsToShift []*elements.Row
	for row := range sheetData.Rows() {
		if row.RowIndex() > deleteEnd {
			rowsToShift = append(rowsToShift, row)
		}
	}

	// Process rows in forward order (lowest row number first)
	for _, row := range rowsToShift {
		oldIndex := row.RowIndex()
		newIndex := oldIndex - count

		// Update row index
		row.SetRowIndex(newIndex)

		// Update all cell references in this row
		for cell := range row.Cells() {
			ref := cell.Reference()
			if ref == "" {
				continue
			}

			cellRef, err := ParseCellRef(ref)
			if err != nil {
				continue
			}

			// Update row number
			cellRef.Row = int(newIndex)
			cell.SetReference(cellRef.String())
		}
	}

	return nil
}

// physicalInsertColumns shifts existing columns right to make space for new columns.
func (s *Sheet) physicalInsertColumns(
	sheetData *elements.SheetData,
	index, count uint32,
) error {
	// For each row, shift cells in affected columns
	for row := range sheetData.Rows() {
		if err := s.shiftRowCellsRight(row, index, count); err != nil {
			return err
		}
	}

	return nil
}

// physicalDeleteColumns removes columns and shifts remaining columns left.
func (s *Sheet) physicalDeleteColumns(
	sheetData *elements.SheetData,
	index, count uint32,
) error {
	deleteEnd := index + count - 1

	// For each row, delete cells in affected columns and shift remaining
	for row := range sheetData.Rows() {
		if err := s.deleteRowCellsAndShiftLeft(row, index, deleteEnd, count); err != nil {
			return err
		}
	}

	return nil
}

// shiftRowCellsRight shifts cells in a row to the right starting at the given column index.
func (s *Sheet) shiftRowCellsRight(
	row *elements.Row,
	index, count uint32,
) error {
	// Collect cells that need to be shifted
	var cellsToShift []*elements.Cell
	for cell := range row.Cells() {
		ref := cell.Reference()
		if ref == "" {
			continue
		}

		cellRef, err := ParseCellRef(ref)
		if err != nil {
			continue
		}

		if uint32(cellRef.Col) >= index {
			cellsToShift = append(
				cellsToShift,
				cell,
			)
		}
	}

	// Process cells in reverse order (rightmost first)
	// Sort by column index descending
	for i := 0; i < len(cellsToShift); i++ {
		for j := i + 1; j < len(cellsToShift); j++ {
			refI, _ := ParseCellRef(
				cellsToShift[i].Reference(),
			)
			refJ, _ := ParseCellRef(
				cellsToShift[j].Reference(),
			)
			if refI.Col < refJ.Col {
				cellsToShift[i], cellsToShift[j] = cellsToShift[j], cellsToShift[i]
			}
		}
	}

	// Update cell references
	for _, cell := range cellsToShift {
		ref := cell.Reference()
		cellRef, err := ParseCellRef(ref)
		if err != nil {
			continue
		}

		newCol := uint32(cellRef.Col) + count
		if newCol > MaxColumn {
			// Skip cells that would exceed maximum
			continue
		}

		cellRef.Col = int(newCol)
		cell.SetReference(cellRef.String())
	}

	return nil
}

// deleteRowCellsAndShiftLeft removes cells in deletion range and shifts remaining cells left.
func (s *Sheet) deleteRowCellsAndShiftLeft(
	row *elements.Row,
	index, deleteEnd, count uint32,
) error {
	// Collect cells to delete and cells to shift
	var cellsToDelete []*elements.Cell
	var cellsToShift []*elements.Cell

	for cell := range row.Cells() {
		ref := cell.Reference()
		if ref == "" {
			continue
		}

		cellRef, err := ParseCellRef(ref)
		if err != nil {
			continue
		}

		col := uint32(cellRef.Col)
		if col >= index && col <= deleteEnd {
			cellsToDelete = append(
				cellsToDelete,
				cell,
			)
		} else if col > deleteEnd {
			cellsToShift = append(cellsToShift, cell)
		}
	}

	// Delete cells in deletion range
	for _, cell := range cellsToDelete {
		row.RemoveChild(cell)
	}

	// Sort cells to shift by column index (ascending)
	for i := 0; i < len(cellsToShift); i++ {
		for j := i + 1; j < len(cellsToShift); j++ {
			refI, _ := ParseCellRef(
				cellsToShift[i].Reference(),
			)
			refJ, _ := ParseCellRef(
				cellsToShift[j].Reference(),
			)
			if refI.Col > refJ.Col {
				cellsToShift[i], cellsToShift[j] = cellsToShift[j], cellsToShift[i]
			}
		}
	}

	// Update cell references for shifted cells
	for _, cell := range cellsToShift {
		ref := cell.Reference()
		cellRef, err := ParseCellRef(ref)
		if err != nil {
			continue
		}

		newCol := uint32(cellRef.Col) - count
		cellRef.Col = int(newCol)
		cell.SetReference(cellRef.String())
	}

	return nil
}

// updateAllReferences updates all references across the workbook after a shift operation.
func (s *Sheet) updateAllReferences(
	operation ShiftOperation,
) error {
	// Update formulas in all worksheets
	if err := s.updateAllFormulas(operation); err != nil {
		return fmt.Errorf(
			"updating formulas: %w",
			err,
		)
	}

	// Update named ranges in workbook
	if err := s.updateNamedRanges(operation); err != nil {
		return fmt.Errorf(
			"updating named ranges: %w",
			err,
		)
	}

	// Update merged cells in this worksheet
	if err := s.updateMergedCells(operation); err != nil {
		return fmt.Errorf(
			"updating merged cells: %w",
			err,
		)
	}

	// Update conditional formatting in this worksheet
	if err := s.updateConditionalFormatting(operation); err != nil {
		return fmt.Errorf(
			"updating conditional formatting: %w",
			err,
		)
	}

	// Update data validation in this worksheet
	if err := s.updateDataValidation(operation); err != nil {
		return fmt.Errorf(
			"updating data validation: %w",
			err,
		)
	}

	// Update table ranges
	if err := s.updateTableRanges(operation); err != nil {
		return fmt.Errorf(
			"updating table ranges: %w",
			err,
		)
	}

	return nil
}

// updateAllFormulas updates formulas across all worksheets.
func (s *Sheet) updateAllFormulas(
	operation ShiftOperation,
) error {
	// Create a formula rewriter for this operation
	rewriter := NewFormulaRewriter(
		operation,
		s.name,
	)

	// Update formulas in all sheets
	for sheet := range s.doc.Sheets() {
		sheetData := sheet.SheetData()
		for row := range sheetData.Rows() {
			for cell := range row.Cells() {
				formula := cell.CellFormula()
				if formula == nil ||
					formula.Formula() == "" {
					continue
				}

				// Rewrite the formula
				oldFormula := formula.Formula()
				newFormula, err := rewriter.Rewrite(
					oldFormula,
				)
				if err != nil {
					// If rewriting fails, skip this formula
					continue
				}

				if newFormula != oldFormula {
					formula.SetFormula(newFormula)
				}
			}
		}
	}

	return nil
}

// updateNamedRanges updates workbook DefinedNames.
func (s *Sheet) updateNamedRanges(
	operation ShiftOperation,
) error {
	// Get workbook part
	wp := s.doc.WorkbookPart()
	if wp == nil {
		return nil
	}

	wb := wp.Workbook().(*elements.Workbook)
	definedNames := wb.DefinedNames()
	if definedNames == nil {
		return nil
	}

	// Create a formula rewriter for this operation
	rewriter := NewFormulaRewriter(
		operation,
		s.name,
	)

	// Update each defined name
	for definedName := range definedNames.DefinedNames() {
		oldFormula := definedName.Formula()
		if oldFormula == "" {
			continue
		}

		newFormula, err := rewriter.Rewrite(
			oldFormula,
		)
		if err != nil {
			// If rewriting fails, skip this defined name
			continue
		}

		if newFormula != oldFormula {
			definedName.SetFormula(newFormula)
		}
	}

	return nil
}

// updateMergedCells updates MergeCells ranges in this worksheet.
func (s *Sheet) updateMergedCells(
	operation ShiftOperation,
) error {
	ws := s.Worksheet()
	mergeCells := ws.MergeCells()
	if mergeCells == nil {
		return nil
	}

	// Create a cell reference updater
	updater := NewCellReferenceUpdater(
		operation,
		s.name,
	)

	// Collect merge cells to update
	mergeRefs := make([]string, 0, 10)
	for mergeCell := range mergeCells.GetMergeCells() {
		ref := mergeCell.Ref()
		if ref == "" {
			continue
		}
		mergeRefs = append(mergeRefs, ref)
	}

	// Update each merge cell range
	for _, ref := range mergeRefs {
		rangeRef, err := ParseRangeRef(ref)
		if err != nil {
			continue
		}

		updatedRange, err := updater.UpdateRange(
			rangeRef,
		)
		if err == ErrReferenceDeleted {
			// Remove this merge cell
			mergeCells.RemoveMergeCell(ref)

			continue
		}

		newRef := updatedRange.String()
		if newRef != ref {
			// Remove old and add new
			mergeCells.RemoveMergeCell(ref)
			mergeCells.AddMergeCell(newRef)
		}
	}

	return nil
}

// updateConditionalFormatting updates ConditionalFormatting ranges.
func (s *Sheet) updateConditionalFormatting(
	operation ShiftOperation,
) error {
	ws := s.Worksheet()

	// Create a cell reference updater
	updater := NewCellReferenceUpdater(
		operation,
		s.name,
	)

	// Iterate over conditional formatting elements
	cfs := ws.ConditionalFormatting()
	for _, cf := range cfs {
		// Update sqref attribute
		sqref := cf.Sqref()
		if sqref == "" {
			continue
		}

		// Parse and update range
		rangeRef, err := ParseRangeRef(sqref)
		if err != nil {
			continue
		}

		updatedRange, err := updater.UpdateRange(
			rangeRef,
		)
		if err == ErrReferenceDeleted {
			// Mark for deletion (we can't delete during iteration)
			cf.SetSqref("#REF!")

			continue
		}

		newRef := updatedRange.String()
		if newRef != sqref {
			cf.SetSqref(newRef)
		}
	}

	return nil
}

// updateDataValidation updates DataValidation ranges and formulas.
func (s *Sheet) updateDataValidation(
	operation ShiftOperation,
) error {
	ws := s.Worksheet()
	dvs := ws.DataValidations()
	if dvs == nil {
		return nil
	}

	// Create updater and rewriter
	updater := NewCellReferenceUpdater(
		operation,
		s.name,
	)
	rewriter := NewFormulaRewriter(
		operation,
		s.name,
	)

	// Iterate over data validation elements
	for dv := range dvs.GetDataValidations() {
		// Update sqref attribute (range)
		sqref := dv.Sqref()
		if sqref != "" {
			rangeRef, err := ParseRangeRef(sqref)
			if err == nil {
				updatedRange, err := updater.UpdateRange(
					rangeRef,
				)
				if err == ErrReferenceDeleted {
					dv.SetSqref("#REF!")
				} else if updatedRange.String() != sqref {
					dv.SetSqref(updatedRange.String())
				}
			}
		}

		// Update formula1 and formula2
		formula1 := dv.Formula1Text()
		if formula1 != "" {
			newFormula, err := rewriter.Rewrite(
				formula1,
			)
			if err == nil &&
				newFormula != formula1 {
				dv.AddFormula1(newFormula)
			}
		}

		formula2 := dv.Formula2Text()
		if formula2 != "" {
			newFormula, err := rewriter.Rewrite(
				formula2,
			)
			if err == nil &&
				newFormula != formula2 {
				dv.AddFormula2(newFormula)
			}
		}
	}

	return nil
}

// updateTableRanges updates Table part ranges.
func (s *Sheet) updateTableRanges(
	operation ShiftOperation,
) error {
	// Get all table definition parts for this worksheet
	tables := s.Tables()

	// Create a cell reference updater
	updater := NewCellReferenceUpdater(
		operation,
		s.name,
	)

	for _, table := range tables {
		// Update ref attribute
		ref := table.Ref()
		if ref == "" {
			continue
		}

		rangeRef, err := ParseRangeRef(ref)
		if err != nil {
			continue
		}

		updatedRange, err := updater.UpdateRange(
			rangeRef,
		)
		if err == ErrReferenceDeleted {
			// Mark table as invalid (SetRef returns error, but we ignore it here)
			_ = table.SetRef("#REF!")

			continue
		}

		newRef := updatedRange.String()
		if newRef != ref {
			_ = table.SetRef(newRef)
		}
	}

	return nil
}

// GroupRows sets outline grouping for rows from start to end (inclusive, 1-based).
// Each call increases the outline level by 1 (max 7 levels).
// Returns an error if parameters are invalid.
func (s *Sheet) GroupRows(
	start, end uint32,
) error {
	// Validate parameters
	if start < MinRow || start > MaxRow {
		return ErrRowOutOfRange
	}
	if end < MinRow || end > MaxRow {
		return ErrRowOutOfRange
	}
	if start > end {
		return errors.New(
			"start must be less than or equal to end",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// For each row in range, increment outline level
	for rowIndex := start; rowIndex <= end; rowIndex++ {
		row := sheetData.GetOrCreateRow(rowIndex)
		currentLevel := row.OutlineLevel()

		// Increment level (cap at 7)
		newLevel := currentLevel + 1
		if newLevel > 7 {
			newLevel = 7
		}

		row.SetOutlineLevel(newLevel)
	}

	return nil
}

// UngroupRows decreases outline level for rows from start to end (inclusive, 1-based).
// If level reaches 0, grouping is removed.
// Returns an error if parameters are invalid.
func (s *Sheet) UngroupRows(
	start, end uint32,
) error {
	// Validate parameters
	if start < MinRow || start > MaxRow {
		return ErrRowOutOfRange
	}
	if end < MinRow || end > MaxRow {
		return ErrRowOutOfRange
	}
	if start > end {
		return errors.New(
			"start must be less than or equal to end",
		)
	}

	// Get sheet data
	sheetData := s.SheetData()

	// For each row in range, decrement outline level
	for rowIndex := start; rowIndex <= end; rowIndex++ {
		row := sheetData.GetOrCreateRow(rowIndex)
		currentLevel := row.OutlineLevel()

		// Decrement level (min 0)
		newLevel := currentLevel - 1
		if newLevel < 0 {
			newLevel = 0
		}

		row.SetOutlineLevel(newLevel)
	}

	return nil
}

// GroupColumns sets outline grouping for columns from start to end (inclusive, 1-based).
// Each call increases the outline level by 1 (max 7 levels).
// Returns an error if parameters are invalid.
func (s *Sheet) GroupColumns(
	start, end uint32,
) error {
	// Validate parameters
	if start < MinColumn || start > MaxColumn {
		return ErrColumnOutOfRange
	}
	if end < MinColumn || end > MaxColumn {
		return ErrColumnOutOfRange
	}
	if start > end {
		return errors.New(
			"start must be less than or equal to end",
		)
	}

	// Get or create Cols element
	ws := s.Worksheet()
	cols := ws.GetOrCreateCols()

	// For each column in range, update outline level
	for colIndex := start; colIndex <= end; colIndex++ {
		col := s.getOrCreateColForIndex(
			cols,
			colIndex,
		)
		currentLevel := col.OutlineLevel()

		// Increment level (cap at 7)
		newLevel := currentLevel + 1
		if newLevel > 7 {
			newLevel = 7
		}

		col.SetOutlineLevel(newLevel)
	}

	return nil
}

// UngroupColumns decreases outline level for columns from start to end (inclusive, 1-based).
// If level reaches 0, grouping is removed.
// Returns an error if parameters are invalid.
func (s *Sheet) UngroupColumns(
	start, end uint32,
) error {
	// Validate parameters
	if start < MinColumn || start > MaxColumn {
		return ErrColumnOutOfRange
	}
	if end < MinColumn || end > MaxColumn {
		return ErrColumnOutOfRange
	}
	if start > end {
		return errors.New(
			"start must be less than or equal to end",
		)
	}

	// Get or create Cols element
	ws := s.Worksheet()
	cols := ws.GetOrCreateCols()

	// For each column in range, update outline level
	for colIndex := start; colIndex <= end; colIndex++ {
		col := s.getOrCreateColForIndex(
			cols,
			colIndex,
		)
		currentLevel := col.OutlineLevel()

		// Decrement level (min 0)
		newLevel := currentLevel - 1
		if newLevel < 0 {
			newLevel = 0
		}

		col.SetOutlineLevel(newLevel)
	}

	return nil
}

// getOrCreateColForIndex returns the Col element for the given column index,
// creating it if necessary. This may split existing Col elements.
func (s *Sheet) getOrCreateColForIndex(
	cols *elements.Cols,
	index uint32,
) *elements.Col {
	idx := int(index)

	// Check if there's an existing Col that contains this column
	for col := range cols.Cols() {
		minCol := col.Min()
		maxCol := col.Max()

		if idx >= minCol && idx <= maxCol {
			// Found a Col that contains this index
			// If it's exactly this column, return it
			if minCol == idx && maxCol == idx {
				return col
			}

			// Need to split the Col element
			// Remove the original col
			cols.RemoveChild(col)

			// Create new Col elements for the split ranges
			if minCol < idx {
				// Create Col for columns before the target
				beforeCol := elements.NewCol()
				beforeCol.SetMin(minCol)
				beforeCol.SetMax(idx - 1)
				// Copy attributes from original
				s.copyColAttributes(
					col,
					beforeCol,
				)
				cols.AppendChild(beforeCol)
			}

			// Create Col for the target column
			targetCol := elements.NewCol()
			targetCol.SetMin(idx)
			targetCol.SetMax(idx)
			// Copy attributes from original
			s.copyColAttributes(col, targetCol)
			cols.AppendChild(targetCol)

			if maxCol > idx {
				// Create Col for columns after the target
				afterCol := elements.NewCol()
				afterCol.SetMin(idx + 1)
				afterCol.SetMax(maxCol)
				// Copy attributes from original
				s.copyColAttributes(col, afterCol)
				cols.AppendChild(afterCol)
			}

			return targetCol
		}
	}

	// No existing Col found, create a new one
	newCol := elements.NewCol()
	newCol.SetMin(idx)
	newCol.SetMax(idx)
	cols.AppendChild(newCol)

	return newCol
}

// copyColAttributes copies all attributes except min and max from src to dst.
func (s *Sheet) copyColAttributes(
	src, dst *elements.Col,
) {
	if width := src.Width(); width > 0 {
		dst.SetWidth(width)
	}
	if style := src.Style(); style > 0 {
		dst.SetStyle(style)
	}
	if src.Hidden() {
		dst.SetHidden(true)
	}
	if src.BestFit() {
		dst.SetBestFit(true)
	}
	if src.CustomWidth() {
		dst.SetCustomWidth(true)
	}
	if src.Collapsed() {
		dst.SetCollapsed(true)
	}
	if level := src.OutlineLevel(); level > 0 {
		dst.SetOutlineLevel(level)
	}
	if src.Phonetic() {
		dst.SetPhonetic(true)
	}
}
