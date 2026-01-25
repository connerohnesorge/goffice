package spreadsheet

import (
	"fmt"
	"io"
)

type WorkbookBuilder struct {
	doc    *Document
	errors []error
}

func NewWorkbookBuilder() *WorkbookBuilder {
	doc, err := CreateFromStream(io.Discard, DocTypeWorkbook)
	wb := &WorkbookBuilder{
		doc:    doc,
		errors: make([]error, 0),
	}
	if err != nil {
		wb.errors = append(wb.errors, err)
	}
	return wb
}

func (wb *WorkbookBuilder) AddSheet(name string) *SheetBuilder {
	if wb.doc == nil {
		wb.errors = append(wb.errors, fmt.Errorf("document is nil"))
		return &SheetBuilder{wb: wb}
	}

	sheet, err := wb.doc.AddSheet(name)
	if err != nil {
		wb.errors = append(wb.errors, err)
		return &SheetBuilder{wb: wb}
	}

	return &SheetBuilder{
		wb:    wb,
		sheet: sheet,
	}
}

func (wb *WorkbookBuilder) Build() (*Document, error) {
	if len(wb.errors) > 0 {
		return nil, wb.errors[0]
	}
	return wb.doc, nil
}

type SheetBuilder struct {
	wb    *WorkbookBuilder
	sheet *Sheet
}

func (sb *SheetBuilder) AddRow() *RowBuilder {
	if sb.sheet == nil {
		return &RowBuilder{sb: sb}
	}

	row := sb.sheet.AddRow()
	return &RowBuilder{
		sb:  sb,
		row: row,
	}
}

func (sb *SheetBuilder) Workbook() *WorkbookBuilder {
	return sb.wb
}

func (sb *SheetBuilder) Sheet() *Sheet {
	return sb.sheet
}

type RowBuilder struct {
	sb  *SheetBuilder
	row *Row
}

func (rb *RowBuilder) Cell(col uint32) *CellBuilder {
	if rb.row == nil {
		return &CellBuilder{rb: rb}
	}

	cell := rb.row.Cell(col)
	return &CellBuilder{
		rb:   rb,
		cell: cell,
	}
}

func (rb *RowBuilder) Sheet() *SheetBuilder {
	return rb.sb
}

func (rb *RowBuilder) Row() *Row {
	return rb.row
}

type CellBuilder struct {
	rb   *RowBuilder
	cell *Cell
}

func (cb *CellBuilder) SetString(value string) *CellBuilder {
	if cb.cell != nil {
		cb.cell.SetString(value)
	}
	return cb
}

func (cb *CellBuilder) SetNumber(value float64) *CellBuilder {
	if cb.cell != nil {
		cb.cell.SetNumber(value)
	}
	return cb
}

func (cb *CellBuilder) SetInt(value int) *CellBuilder {
	if cb.cell != nil {
		cb.cell.SetNumber(float64(value))
	}
	return cb
}

func (cb *CellBuilder) SetBool(value bool) *CellBuilder {
	if cb.cell != nil {
		cb.cell.SetBoolean(value)
	}
	return cb
}

func (cb *CellBuilder) SetStyle(style *Style) *CellBuilder {
	if cb.cell != nil {
		cb.cell.SetStyle(style)
	}
	return cb
}

func (cb *CellBuilder) Row() *RowBuilder {
	return cb.rb
}

func (cb *CellBuilder) Cell() *Cell {
	return cb.cell
}

func (cb *CellBuilder) Sheet() *SheetBuilder {
	return cb.rb.sb
}

func (cb *CellBuilder) Workbook() *WorkbookBuilder {
	return cb.rb.sb.wb
}

type RangeBuilder struct {
	sb   *SheetBuilder
	rnge *Range
}

func (sb *SheetBuilder) Range(ref string) *RangeBuilder {
	if sb.sheet == nil {
		return &RangeBuilder{sb: sb}
	}

	r := sb.sheet.Range(ref)
	if r == nil {
		return &RangeBuilder{sb: sb}
	}

	return &RangeBuilder{
		sb:   sb,
		rnge: r,
	}
}

func (rb *RangeBuilder) SetStyle(style *Style) *RangeBuilder {
	if rb.rnge != nil {
		rb.rnge.SetStyle(style)
	}
	return rb
}

func (rb *RangeBuilder) SetValue(value interface{}) *RangeBuilder {
	if rb.rnge != nil {
		rb.rnge.SetValue(value)
	}
	return rb
}

func (rb *RangeBuilder) Merge() *RangeBuilder {
	if rb.rnge != nil {
		_ = rb.sb.sheet.MergeCells(rb.rnge.String())
	}
	return rb
}

func (rb *RangeBuilder) Sheet() *SheetBuilder {
	return rb.sb
}
