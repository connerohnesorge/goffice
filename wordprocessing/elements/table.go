//nolint:revive // file-length-limit - table implementation requires many methods
package elements

import (
	"iter"
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// TableWidthType represents the type of table/cell width.
type TableWidthType string

const (
	// TableWidthTypeAuto allows automatic width.
	TableWidthTypeAuto TableWidthType = "auto"
	// TableWidthTypeDxa uses twips (twentieths of a point).
	TableWidthTypeDxa TableWidthType = "dxa"
	// TableWidthTypeNil indicates no width.
	TableWidthTypeNil TableWidthType = "nil"
	// TableWidthTypePct uses percentage (fifths of a percent, e.g., 5000 = 100%).
	TableWidthTypePct TableWidthType = "pct"
)

// TableProperties returns the table properties element, or nil if not present.
func (t *Table) TableProperties() *TableProperties {
	elem := t.GetElement("tblPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tp, ok := elem.(*TableProperties); ok {
		return tp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableProperties returns the table properties, creating if needed.
func (t *Table) GetOrCreateTableProperties() *TableProperties {
	props := t.TableProperties()
	if props != nil {
		return props
	}
	props = NewTableProperties()
	// Properties should be first child
	if first := t.FirstChild(); first != nil {
		t.InsertBefore(props, first)
	} else {
		t.AppendChild(props)
	}

	return props
}

// TableGrid returns the table grid element, or nil if not present.
func (t *Table) TableGrid() *TableGrid {
	elem := t.GetElement("tblGrid", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tg, ok := elem.(*TableGrid); ok {
		return tg
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableGrid{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Rows returns an iterator over all TableRow elements in the table.
func (t *Table) Rows() iter.Seq[*TableRow] {
	return func(yield func(*TableRow) bool) {
		for child := range t.Children() {
			if child.LocalName() == "tr" &&
				child.NamespaceURI() == NamespaceWML {
				var tr *TableRow
				if row, ok := child.(*TableRow); ok {
					tr = row
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					tr = &TableRow{CompositeElementBase: comp}
				}
				if tr != nil && !yield(tr) {
					return
				}
			}
		}
	}
}

// GetRow returns the row at the specified index (0-based).
func (t *Table) GetRow(index int) *TableRow {
	i := 0
	for row := range t.Rows() {
		if i == index {
			return row
		}
		i++
	}

	return nil
}

// RowCount returns the number of rows in the table.
func (t *Table) RowCount() int {
	count := 0
	for range t.Rows() {
		count++
	}

	return count
}

// AppendRow appends a new row with the specified number of cells.
func (t *Table) AppendRow(cols int) *TableRow {
	tr := NewTableRow(cols)
	t.AppendChild(tr)

	return tr
}

// InsertRow inserts a new row at the specified index.
func (t *Table) InsertRow(
	index, cols int,
) *TableRow {
	tr := NewTableRow(cols)
	refRow := t.GetRow(index)
	if refRow != nil {
		t.InsertBefore(tr, refRow)
	} else {
		t.AppendChild(tr)
	}

	return tr
}

// DeleteRow removes the row at the specified index.
func (t *Table) DeleteRow(index int) bool {
	row := t.GetRow(index)
	if row != nil {
		return t.RemoveChild(row)
	}

	return false
}

// GetCell returns the cell at the specified row and column (0-based).
func (t *Table) GetCell(row, col int) *TableCell {
	r := t.GetRow(row)
	if r == nil {
		return nil
	}

	return r.GetCell(col)
}

// SetCellText sets the text content of a specific cell.
func (t *Table) SetCellText(
	row, col int,
	text string,
) {
	cell := t.GetCell(row, col)
	if cell != nil {
		cell.SetText(text)
	}
}

// SetStyle sets the table style.
func (t *Table) SetStyle(styleId string) *Table {
	t.GetOrCreateTableProperties().
		SetTableStyle(styleId)

	return t
}

// SetWidth sets the table width.
func (t *Table) SetWidth(
	width int,
	widthType TableWidthType,
) *Table {
	t.GetOrCreateTableProperties().
		SetTableWidth(width, widthType)

	return t
}

// SetColumnWidth sets the width of a specific column in twips.
func (t *Table) SetColumnWidth(
	col, width int,
) *Table {
	grid := t.TableGrid()
	if grid != nil {
		grid.SetColumnWidth(col, width)
	}

	return t
}

// TableProperties represents the w:tblPr element.
type TableProperties struct {
	*openxml.CompositeElementBase
}

// NewTableProperties creates a new TableProperties element.
func NewTableProperties() *TableProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblPr",
		PrefixW,
	)

	return &TableProperties{
		CompositeElementBase: elem,
	}
}

// TableStyle returns the table style ID, or empty string if not set.
func (tp *TableProperties) TableStyle() string {
	elem := tp.GetElement(
		"tblStyle",
		NamespaceWML,
	)
	if elem == nil {
		return ""
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetTableStyle sets the table style.
func (tp *TableProperties) SetTableStyle(
	styleId string,
) {
	elem := tp.GetElement(
		"tblStyle",
		NamespaceWML,
	)
	if elem == nil {
		elem = openxml.NewCompositeElement(
			NamespaceWML,
			"tblStyle",
			PrefixW,
		)
		// Insert at beginning
		if first := tp.FirstChild(); first != nil {
			tp.InsertBefore(elem, first)
		} else {
			tp.AppendChild(elem)
		}
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			styleId,
		),
	)
}

// TableWidth returns the table width element.
func (tp *TableProperties) TableWidth() *TableWidth {
	elem := tp.GetElement("tblW", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tw, ok := elem.(*TableWidth); ok {
		return tw
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableWidth{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetTableWidth sets the table width.
func (tp *TableProperties) SetTableWidth(
	width int,
	widthType TableWidthType,
) {
	tw := tp.TableWidth()
	if tw == nil {
		tw = NewTableWidth(width, widthType)
		tp.AppendChild(tw)
	} else {
		tw.SetWidth(width)
		tw.SetType(widthType)
	}
}

// TableBorders returns the table borders element.
func (tp *TableProperties) TableBorders() *TableBorders {
	elem := tp.GetElement(
		"tblBorders",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if tb, ok := elem.(*TableBorders); ok {
		return tb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableBorders{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableBorders returns the table borders, creating if needed.
func (tp *TableProperties) GetOrCreateTableBorders() *TableBorders {
	borders := tp.TableBorders()
	if borders != nil {
		return borders
	}
	borders = NewTableBorders()
	tp.AppendChild(borders)

	return borders
}

// TableLook returns the table look element.
func (tp *TableProperties) TableLook() *TableLook {
	elem := tp.GetElement("tblLook", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tl, ok := elem.(*TableLook); ok {
		return tl
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableLook{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetTableLook sets the table look properties.
func (tp *TableProperties) SetTableLook(
	firstRow, lastRow, firstColumn, lastColumn, noHBand, noVBand bool,
) {
	tl := tp.TableLook()
	if tl == nil {
		tl = NewTableLook()
		tp.AppendChild(tl)
	}
	tl.SetFirstRow(firstRow)
	tl.SetLastRow(lastRow)
	tl.SetFirstColumn(firstColumn)
	tl.SetLastColumn(lastColumn)
	tl.SetNoHBand(noHBand)
	tl.SetNoVBand(noVBand)
}

// Clone creates a deep copy of this TableProperties element.
func (tp *TableProperties) Clone() openxml.Element {
	return &TableProperties{
		CompositeElementBase: tp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableWidth represents the w:tblW element.
type TableWidth struct {
	*openxml.CompositeElementBase
}

// NewTableWidth creates a new TableWidth element.
func NewTableWidth(
	width int,
	widthType TableWidthType,
) *TableWidth {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblW",
		PrefixW,
	)
	tw := &TableWidth{CompositeElementBase: elem}
	tw.SetWidth(width)
	tw.SetType(widthType)

	return tw
}

// Width returns the width value.
func (tw *TableWidth) Width() int {
	attr, found := tw.GetAttribute(
		"w",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWidth sets the width value.
func (tw *TableWidth) SetWidth(width int) {
	tw.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"w",
			PrefixW,
			strconv.Itoa(width),
		),
	)
}

// Type returns the width type.
func (tw *TableWidth) Type() TableWidthType {
	attr, found := tw.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return TableWidthTypeAuto
	}

	return TableWidthType(attr.Value())
}

// SetType sets the width type.
func (tw *TableWidth) SetType(
	widthType TableWidthType,
) {
	tw.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"type",
			PrefixW,
			string(widthType),
		),
	)
}

// Clone creates a deep copy of this TableWidth element.
func (tw *TableWidth) Clone() openxml.Element {
	return &TableWidth{
		CompositeElementBase: tw.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableBorders represents the w:tblBorders element.
type TableBorders struct {
	*openxml.CompositeElementBase
}

// NewTableBorders creates a new TableBorders element.
func NewTableBorders() *TableBorders {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblBorders",
		PrefixW,
	)

	return &TableBorders{
		CompositeElementBase: elem,
	}
}

// SetAllBorders sets all borders to the same style.
func (tb *TableBorders) SetAllBorders(
	style BorderStyle,
	size int,
	color string,
) {
	tb.SetTop(style, size, color)
	tb.SetBottom(style, size, color)
	tb.SetLeft(style, size, color)
	tb.SetRight(style, size, color)
	tb.SetInsideH(style, size, color)
	tb.SetInsideV(style, size, color)
}

// SetTop sets the top border.
func (tb *TableBorders) SetTop(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("top", style, size, color)
}

// SetBottom sets the bottom border.
func (tb *TableBorders) SetBottom(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("bottom", style, size, color)
}

// SetLeft sets the left border.
func (tb *TableBorders) SetLeft(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("left", style, size, color)
}

// SetRight sets the right border.
func (tb *TableBorders) SetRight(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("right", style, size, color)
}

// SetInsideH sets the inside horizontal border.
func (tb *TableBorders) SetInsideH(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("insideH", style, size, color)
}

// SetInsideV sets the inside vertical border.
func (tb *TableBorders) SetInsideV(
	style BorderStyle,
	size int,
	color string,
) {
	tb.setBorder("insideV", style, size, color)
}

func (tb *TableBorders) setBorder(
	name string,
	style BorderStyle,
	size int,
	color string,
) {
	elem := tb.GetElement(name, NamespaceWML)
	if elem == nil {
		elem = openxml.NewCompositeElement(
			NamespaceWML,
			name,
			PrefixW,
		)
		tb.AppendChild(elem)
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(style),
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"sz",
			PrefixW,
			strconv.Itoa(size),
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"color",
			PrefixW,
			color,
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"space",
			PrefixW,
			"0",
		),
	)
}

// Clone creates a deep copy of this TableBorders element.
func (tb *TableBorders) Clone() openxml.Element {
	return &TableBorders{
		CompositeElementBase: tb.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableLook represents the w:tblLook element.
type TableLook struct {
	*openxml.CompositeElementBase
}

// NewTableLook creates a new TableLook element.
func NewTableLook() *TableLook {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblLook",
		PrefixW,
	)

	return &TableLook{CompositeElementBase: elem}
}

// SetFirstRow sets whether to apply first row formatting.
func (tl *TableLook) SetFirstRow(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"firstRow",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

// SetLastRow sets whether to apply last row formatting.
func (tl *TableLook) SetLastRow(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"lastRow",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

// SetFirstColumn sets whether to apply first column formatting.
func (tl *TableLook) SetFirstColumn(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"firstColumn",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

// SetLastColumn sets whether to apply last column formatting.
func (tl *TableLook) SetLastColumn(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"lastColumn",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

// SetNoHBand sets whether to suppress horizontal banding.
func (tl *TableLook) SetNoHBand(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"noHBand",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

// SetNoVBand sets whether to suppress vertical banding.
func (tl *TableLook) SetNoVBand(val bool) {
	tl.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"noVBand",
			PrefixW,
			boolToOnOff(val),
		),
	)
}

func boolToOnOff(val bool) string {
	if val {
		return "1"
	}

	return "0"
}

// Clone creates a deep copy of this TableLook element.
func (tl *TableLook) Clone() openxml.Element {
	return &TableLook{
		CompositeElementBase: tl.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableGrid represents the w:tblGrid element.
type TableGrid struct {
	*openxml.CompositeElementBase
}

// NewTableGrid creates a new TableGrid element.
func NewTableGrid() *TableGrid {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tblGrid",
		PrefixW,
	)

	return &TableGrid{CompositeElementBase: elem}
}

// GridColumns returns an iterator over grid column elements.
func (tg *TableGrid) GridColumns() iter.Seq[*GridColumn] {
	return func(yield func(*GridColumn) bool) {
		for child := range tg.Children() {
			if child.LocalName() == "gridCol" &&
				child.NamespaceURI() == NamespaceWML {
				var gc *GridColumn
				if col, ok := child.(*GridColumn); ok {
					gc = col
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					gc = &GridColumn{CompositeElementBase: comp}
				}
				if gc != nil && !yield(gc) {
					return
				}
			}
		}
	}
}

// SetColumnWidth sets the width of a specific column in twips.
func (tg *TableGrid) SetColumnWidth(
	col, width int,
) {
	i := 0
	for gc := range tg.GridColumns() {
		if i == col {
			gc.SetWidth(width)

			return
		}
		i++
	}
}

// Clone creates a deep copy of this TableGrid element.
func (tg *TableGrid) Clone() openxml.Element {
	return &TableGrid{
		CompositeElementBase: tg.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GridColumn represents the w:gridCol element.
type GridColumn struct {
	*openxml.CompositeElementBase
}

// NewGridColumn creates a new GridColumn element.
func NewGridColumn(width int) *GridColumn {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"gridCol",
		PrefixW,
	)
	gc := &GridColumn{CompositeElementBase: elem}
	if width > 0 {
		gc.SetWidth(width)
	}

	return gc
}

// Width returns the column width in twips.
func (gc *GridColumn) Width() int {
	attr, found := gc.GetAttribute(
		"w",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWidth sets the column width in twips.
func (gc *GridColumn) SetWidth(width int) {
	gc.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"w",
			PrefixW,
			strconv.Itoa(width),
		),
	)
}

// Clone creates a deep copy of this GridColumn element.
func (gc *GridColumn) Clone() openxml.Element {
	return &GridColumn{
		CompositeElementBase: gc.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableRowProperties returns the table row properties element, or nil if not present.
func (tr *TableRow) TableRowProperties() *TableRowProperties {
	elem := tr.GetElement("trPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if trp, ok := elem.(*TableRowProperties); ok {
		return trp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableRowProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableRowProperties returns the table row properties, creating if needed.
func (tr *TableRow) GetOrCreateTableRowProperties() *TableRowProperties {
	props := tr.TableRowProperties()
	if props != nil {
		return props
	}
	props = NewTableRowProperties()
	// Properties should be first child
	if first := tr.FirstChild(); first != nil {
		tr.InsertBefore(props, first)
	} else {
		tr.AppendChild(props)
	}

	return props
}

// Cells returns an iterator over all TableCell elements in this row.
func (tr *TableRow) Cells() iter.Seq[*TableCell] {
	return func(yield func(*TableCell) bool) {
		for child := range tr.Children() {
			if child.LocalName() == "tc" &&
				child.NamespaceURI() == NamespaceWML {
				var tc *TableCell
				if cell, ok := child.(*TableCell); ok {
					tc = cell
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					tc = &TableCell{CompositeElementBase: comp}
				}
				if tc != nil && !yield(tc) {
					return
				}
			}
		}
	}
}

// GetCell returns the cell at the specified index (0-based).
func (tr *TableRow) GetCell(
	index int,
) *TableCell {
	i := 0
	for cell := range tr.Cells() {
		if i == index {
			return cell
		}
		i++
	}

	return nil
}

// CellCount returns the number of cells in this row.
func (tr *TableRow) CellCount() int {
	count := 0
	for range tr.Cells() {
		count++
	}

	return count
}

// AppendCell appends a new cell to this row.
func (tr *TableRow) AppendCell() *TableCell {
	tc := NewTableCell()
	tr.AppendChild(tc)

	return tc
}

// SetHeight sets the row height in twips.
func (tr *TableRow) SetHeight(
	height int,
	rule HeightRule,
) {
	tr.GetOrCreateTableRowProperties().
		SetHeight(height, rule)
}

// SetHeaderRow marks this row as a header row.
func (tr *TableRow) SetHeaderRow(isHeader bool) {
	tr.GetOrCreateTableRowProperties().
		SetHeader(isHeader)
}

// TableRowProperties represents the w:trPr element.
type TableRowProperties struct {
	*openxml.CompositeElementBase
}

// NewTableRowProperties creates a new TableRowProperties element.
func NewTableRowProperties() *TableRowProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"trPr",
		PrefixW,
	)

	return &TableRowProperties{
		CompositeElementBase: elem,
	}
}

// HeightRule represents the height rule for a table row.
type HeightRule string

const (
	// HeightRuleAuto uses automatic height.
	HeightRuleAuto HeightRule = "auto"
	// HeightRuleExact uses exact height.
	HeightRuleExact HeightRule = "exact"
	// HeightRuleAtLeast uses at least the specified height.
	HeightRuleAtLeast HeightRule = "atLeast"
)

// SetHeight sets the row height.
func (trp *TableRowProperties) SetHeight(
	height int,
	rule HeightRule,
) {
	elem := trp.GetElement(
		"trHeight",
		NamespaceWML,
	)
	if elem == nil {
		elem = openxml.NewCompositeElement(
			NamespaceWML,
			"trHeight",
			PrefixW,
		)
		trp.AppendChild(elem)
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(height),
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"hRule",
			PrefixW,
			string(rule),
		),
	)
}

// SetHeader sets whether this row is a header row.
func (trp *TableRowProperties) SetHeader(
	isHeader bool,
) {
	elem := trp.GetElement(
		"tblHeader",
		NamespaceWML,
	)
	if isHeader {
		if elem == nil {
			elem = openxml.NewCompositeElement(
				NamespaceWML,
				"tblHeader",
				PrefixW,
			)
			trp.AppendChild(elem)
		}
	} else {
		if elem != nil {
			trp.RemoveChild(elem)
		}
	}
}

// Clone creates a deep copy of this TableRowProperties element.
func (trp *TableRowProperties) Clone() openxml.Element {
	return &TableRowProperties{
		CompositeElementBase: trp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableCellProperties returns the table cell properties element, or nil if not present.
func (tc *TableCell) TableCellProperties() *TableCellProperties {
	elem := tc.GetElement("tcPr", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tcp, ok := elem.(*TableCellProperties); ok {
		return tcp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableCellProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableCellProperties returns the table cell properties, creating if needed.
func (tc *TableCell) GetOrCreateTableCellProperties() *TableCellProperties {
	props := tc.TableCellProperties()
	if props != nil {
		return props
	}
	props = NewTableCellProperties()
	// Properties should be first child
	if first := tc.FirstChild(); first != nil {
		tc.InsertBefore(props, first)
	} else {
		tc.AppendChild(props)
	}

	return props
}

// Paragraphs returns an iterator over all Paragraph elements in this cell.
func (tc *TableCell) Paragraphs() iter.Seq[*Paragraph] {
	return func(yield func(*Paragraph) bool) {
		for child := range tc.Children() {
			if child.LocalName() == "p" &&
				child.NamespaceURI() == NamespaceWML {
				var p *Paragraph
				if para, ok := child.(*Paragraph); ok {
					p = para
				} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
					p = &Paragraph{CompositeElementBase: comp}
				}
				if p != nil && !yield(p) {
					return
				}
			}
		}
	}
}

// AppendParagraph appends a new paragraph with the given text to this cell.
func (tc *TableCell) AppendParagraph(
	text string,
) *Paragraph {
	p := NewParagraph(text)
	tc.AppendChild(p)

	return p
}

// SetText sets the text content of the cell, replacing all existing content.
func (tc *TableCell) SetText(text string) {
	// Remove all paragraphs
	var toRemove []openxml.Element
	for child := range tc.Children() {
		if child.LocalName() == "p" &&
			child.NamespaceURI() == NamespaceWML {
			toRemove = append(toRemove, child)
		}
	}
	for _, elem := range toRemove {
		tc.RemoveChild(elem)
	}
	// Add new paragraph with text
	tc.AppendParagraph(text)
}

// InnerText returns the concatenated text content of the cell.
func (tc *TableCell) InnerText() string {
	var text string
	for p := range tc.Paragraphs() {
		if text != "" {
			text += "\n"
		}
		text += p.InnerText()
	}

	return text
}

// SetWidth sets the cell width.
func (tc *TableCell) SetWidth(
	width int,
	widthType TableWidthType,
) {
	tc.GetOrCreateTableCellProperties().
		SetWidth(width, widthType)
}

// SetShading sets the cell shading/background color.
func (tc *TableCell) SetShading(
	fillColor string,
) {
	tc.GetOrCreateTableCellProperties().
		SetShading(fillColor)
}

// SetVerticalMerge sets vertical cell merge.
func (tc *TableCell) SetVerticalMerge(
	mergeType VerticalMergeType,
) {
	tc.GetOrCreateTableCellProperties().
		SetVerticalMerge(mergeType)
}

// SetHorizontalMerge sets horizontal cell span (gridSpan).
func (tc *TableCell) SetHorizontalMerge(
	span int,
) {
	tc.GetOrCreateTableCellProperties().
		SetGridSpan(span)
}

// TableCellProperties represents the w:tcPr element.
type TableCellProperties struct {
	*openxml.CompositeElementBase
}

// NewTableCellProperties creates a new TableCellProperties element.
func NewTableCellProperties() *TableCellProperties {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tcPr",
		PrefixW,
	)

	return &TableCellProperties{
		CompositeElementBase: elem,
	}
}

// TableCellWidth returns the cell width element.
func (tcp *TableCellProperties) TableCellWidth() *TableCellWidth {
	elem := tcp.GetElement("tcW", NamespaceWML)
	if elem == nil {
		return nil
	}
	if tcw, ok := elem.(*TableCellWidth); ok {
		return tcw
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableCellWidth{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetWidth sets the cell width.
func (tcp *TableCellProperties) SetWidth(
	width int,
	widthType TableWidthType,
) {
	tcw := tcp.TableCellWidth()
	if tcw == nil {
		tcw = NewTableCellWidth(width, widthType)
		// Insert at beginning
		if first := tcp.FirstChild(); first != nil {
			tcp.InsertBefore(tcw, first)
		} else {
			tcp.AppendChild(tcw)
		}
	} else {
		tcw.SetWidth(width)
		tcw.SetType(widthType)
	}
}

// VerticalMergeType represents the type of vertical merge.
type VerticalMergeType string

const (
	// VerticalMergeRestart starts a new vertical merge.
	VerticalMergeRestart VerticalMergeType = "restart"
	// VerticalMergeContinue continues an existing vertical merge.
	VerticalMergeContinue VerticalMergeType = "continue"
)

// VerticalMerge returns the vertical merge element.
func (tcp *TableCellProperties) VerticalMerge() *VerticalMerge {
	elem := tcp.GetElement("vMerge", NamespaceWML)
	if elem == nil {
		return nil
	}
	if vm, ok := elem.(*VerticalMerge); ok {
		return vm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &VerticalMerge{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetVerticalMerge sets the vertical merge type.
func (tcp *TableCellProperties) SetVerticalMerge(
	mergeType VerticalMergeType,
) {
	vm := tcp.VerticalMerge()
	if vm == nil {
		vm = NewVerticalMerge(mergeType)
		tcp.AppendChild(vm)
	} else {
		vm.SetType(mergeType)
	}
}

// GridSpan returns the grid span value.
func (tcp *TableCellProperties) GridSpan() int {
	elem := tcp.GetElement(
		"gridSpan",
		NamespaceWML,
	)
	if elem == nil {
		return 1
	}
	attr, found := elem.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return 1
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetGridSpan sets the grid span (horizontal merge).
func (tcp *TableCellProperties) SetGridSpan(
	span int,
) {
	elem := tcp.GetElement(
		"gridSpan",
		NamespaceWML,
	)
	if elem == nil {
		elem = openxml.NewCompositeElement(
			NamespaceWML,
			"gridSpan",
			PrefixW,
		)
		tcp.AppendChild(elem)
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			strconv.Itoa(span),
		),
	)
}

// Shading returns the shading element.
func (tcp *TableCellProperties) Shading() *Shading {
	elem := tcp.GetElement("shd", NamespaceWML)
	if elem == nil {
		return nil
	}
	if s, ok := elem.(*Shading); ok {
		return s
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Shading{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetShading sets the cell shading/background color.
func (tcp *TableCellProperties) SetShading(
	fillColor string,
) {
	shd := tcp.Shading()
	if shd == nil {
		shd = NewShading()
		shd.SetFill(fillColor)
		shd.SetVal(ShadingClear)
		shd.SetColor("auto")
		tcp.AppendChild(shd)
	} else {
		shd.SetFill(fillColor)
	}
}

// TableCellBorders returns the cell borders element.
func (tcp *TableCellProperties) TableCellBorders() *TableCellBorders {
	elem := tcp.GetElement(
		"tcBorders",
		NamespaceWML,
	)
	if elem == nil {
		return nil
	}
	if tcb, ok := elem.(*TableCellBorders); ok {
		return tcb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &TableCellBorders{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTableCellBorders returns the cell borders, creating if needed.
func (tcp *TableCellProperties) GetOrCreateTableCellBorders() *TableCellBorders {
	borders := tcp.TableCellBorders()
	if borders != nil {
		return borders
	}
	borders = NewTableCellBorders()
	tcp.AppendChild(borders)

	return borders
}

// Clone creates a deep copy of this TableCellProperties element.
func (tcp *TableCellProperties) Clone() openxml.Element {
	return &TableCellProperties{
		CompositeElementBase: tcp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableCellWidth represents the w:tcW element.
type TableCellWidth struct {
	*openxml.CompositeElementBase
}

// NewTableCellWidth creates a new TableCellWidth element.
func NewTableCellWidth(
	width int,
	widthType TableWidthType,
) *TableCellWidth {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tcW",
		PrefixW,
	)
	tcw := &TableCellWidth{
		CompositeElementBase: elem,
	}
	tcw.SetWidth(width)
	tcw.SetType(widthType)

	return tcw
}

// Width returns the width value.
func (tcw *TableCellWidth) Width() int {
	attr, found := tcw.GetAttribute(
		"w",
		NamespaceWML,
	)
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())

	return val
}

// SetWidth sets the width value.
func (tcw *TableCellWidth) SetWidth(width int) {
	tcw.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"w",
			PrefixW,
			strconv.Itoa(width),
		),
	)
}

// Type returns the width type.
func (tcw *TableCellWidth) Type() TableWidthType {
	attr, found := tcw.GetAttribute(
		"type",
		NamespaceWML,
	)
	if !found {
		return TableWidthTypeAuto
	}

	return TableWidthType(attr.Value())
}

// SetType sets the width type.
func (tcw *TableCellWidth) SetType(
	widthType TableWidthType,
) {
	tcw.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"type",
			PrefixW,
			string(widthType),
		),
	)
}

// Clone creates a deep copy of this TableCellWidth element.
func (tcw *TableCellWidth) Clone() openxml.Element {
	return &TableCellWidth{
		CompositeElementBase: tcw.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// VerticalMerge represents the w:vMerge element.
type VerticalMerge struct {
	*openxml.CompositeElementBase
}

// NewVerticalMerge creates a new VerticalMerge element.
func NewVerticalMerge(
	mergeType VerticalMergeType,
) *VerticalMerge {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"vMerge",
		PrefixW,
	)
	vm := &VerticalMerge{
		CompositeElementBase: elem,
	}
	if mergeType == VerticalMergeRestart {
		vm.SetType(mergeType)
	}
	// For "continue", no val attribute is needed (it's the default)
	return vm
}

// Type returns the merge type.
func (vm *VerticalMerge) Type() VerticalMergeType {
	attr, found := vm.GetAttribute(
		"val",
		NamespaceWML,
	)
	if !found {
		return VerticalMergeContinue // Default when no val attribute
	}

	return VerticalMergeType(attr.Value())
}

// SetType sets the merge type.
func (vm *VerticalMerge) SetType(
	mergeType VerticalMergeType,
) {
	if mergeType == VerticalMergeRestart {
		vm.SetAttribute(
			openxml.NewAttribute(
				NamespaceWML,
				"val",
				PrefixW,
				string(mergeType),
			),
		)
	} else {
		vm.RemoveAttribute("val", NamespaceWML)
	}
}

// Clone creates a deep copy of this VerticalMerge element.
func (vm *VerticalMerge) Clone() openxml.Element {
	return &VerticalMerge{
		CompositeElementBase: vm.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// TableCellBorders represents the w:tcBorders element.
type TableCellBorders struct {
	*openxml.CompositeElementBase
}

// NewTableCellBorders creates a new TableCellBorders element.
func NewTableCellBorders() *TableCellBorders {
	elem := openxml.NewCompositeElement(
		NamespaceWML,
		"tcBorders",
		PrefixW,
	)

	return &TableCellBorders{
		CompositeElementBase: elem,
	}
}

// SetAllBorders sets all borders to the same style.
func (tcb *TableCellBorders) SetAllBorders(
	style BorderStyle,
	size int,
	color string,
) {
	tcb.SetTop(style, size, color)
	tcb.SetBottom(style, size, color)
	tcb.SetLeft(style, size, color)
	tcb.SetRight(style, size, color)
}

// SetTop sets the top border.
func (tcb *TableCellBorders) SetTop(
	style BorderStyle,
	size int,
	color string,
) {
	tcb.setBorder("top", style, size, color)
}

// SetBottom sets the bottom border.
func (tcb *TableCellBorders) SetBottom(
	style BorderStyle,
	size int,
	color string,
) {
	tcb.setBorder("bottom", style, size, color)
}

// SetLeft sets the left border.
func (tcb *TableCellBorders) SetLeft(
	style BorderStyle,
	size int,
	color string,
) {
	tcb.setBorder("left", style, size, color)
}

// SetRight sets the right border.
func (tcb *TableCellBorders) SetRight(
	style BorderStyle,
	size int,
	color string,
) {
	tcb.setBorder("right", style, size, color)
}

func (tcb *TableCellBorders) setBorder(
	name string,
	style BorderStyle,
	size int,
	color string,
) {
	elem := tcb.GetElement(name, NamespaceWML)
	if elem == nil {
		elem = openxml.NewCompositeElement(
			NamespaceWML,
			name,
			PrefixW,
		)
		tcb.AppendChild(elem)
	}
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"val",
			PrefixW,
			string(style),
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"sz",
			PrefixW,
			strconv.Itoa(size),
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"color",
			PrefixW,
			color,
		),
	)
	elem.SetAttribute(
		openxml.NewAttribute(
			NamespaceWML,
			"space",
			PrefixW,
			"0",
		),
	)
}

// Clone creates a deep copy of this TableCellBorders element.
func (tcb *TableCellBorders) Clone() openxml.Element {
	return &TableCellBorders{
		CompositeElementBase: tcb.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
