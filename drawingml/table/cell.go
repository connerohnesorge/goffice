// Package table provides DrawingML table support for presentations, spreadsheets, and diagrams.
// It implements the core table structure (a:tbl) including grids, rows, cells, and properties.
package table

import (
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// base10Cell is the numeric base for parsing/formatting EMU values in cell properties.
const base10Cell = 10

// TableCell represents a table cell (a:tc).
type TableCell struct {
	*openxml.CompositeElementBase
}

// NewTableCell creates a new table cell with an empty text body.
func NewTableCell() *TableCell {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tc",
		PrefixMain,
	)
	cell := &TableCell{CompositeElementBase: elem}

	// Add empty text body
	tb := drawingml.NewTextBody()
	tb.AddEmptyParagraph()
	cell.AppendChild(tb)

	// Add default cell properties
	cell.AppendChild(NewCellProperties())

	return cell
}

// TextBody returns the text body of this cell.
func (tc *TableCell) TextBody() *drawingml.TextBody {
	elem := tc.GetElement("txBody", NamespaceMain)
	if elem == nil {
		return nil
	}
	if tb, ok := elem.(*drawingml.TextBody); ok {
		return tb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &drawingml.TextBody{CompositeElementBase: comp}
	}

	return nil
}

// SetTextBody sets the text body of this cell.
func (tc *TableCell) SetTextBody(tb *drawingml.TextBody) {
	if existing := tc.GetElement("txBody", NamespaceMain); existing != nil {
		tc.RemoveChild(existing)
	}
	if tb == nil {
		return
	}
	// Insert before properties
	if props := tc.Properties(); props != nil {
		tc.InsertBefore(tb, props)
	} else {
		tc.AppendChild(tb)
	}
}

// Properties returns the cell properties.
func (tc *TableCell) Properties() *CellProperties {
	elem := tc.GetElement("tcPr", NamespaceMain)
	if elem == nil {
		return nil
	}
	if props, ok := elem.(*CellProperties); ok {
		return props
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CellProperties{CompositeElementBase: comp}
	}

	return nil
}

// SetProperties sets the cell properties.
func (tc *TableCell) SetProperties(props *CellProperties) {
	if existing := tc.GetElement("tcPr", NamespaceMain); existing != nil {
		tc.RemoveChild(existing)
	}
	if props != nil {
		tc.AppendChild(props)
	}
}

// SetText is a convenience method to set the cell text.
// This clears any existing paragraphs and creates a single paragraph with the given text.
func (tc *TableCell) SetText(text string) {
	tb := tc.TextBody()
	if tb == nil {
		tb = drawingml.NewTextBody()
		tc.SetTextBody(tb)
	}
	tb.ClearParagraphs()
	tb.AddParagraph(text)
}

// GetText is a convenience method to get the cell text.
// This concatenates all text runs from all paragraphs.
func (tc *TableCell) GetText() string {
	tb := tc.TextBody()
	if tb == nil {
		return ""
	}

	var result string
	for i, para := range tb.Paragraphs() {
		if i > 0 {
			result += "\n"
		}
		for run := range para.Children() {
			if run.LocalName() != "r" || run.NamespaceURI() != NamespaceMain {
				continue
			}
			if textRun, ok := run.(*drawingml.TextRun); ok {
				result += textRun.Text()
			}
		}
	}

	return result
}

// Clone creates a deep copy of this TableCell element.
func (tc *TableCell) Clone() openxml.Element {
	return &TableCell{
		CompositeElementBase: tc.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CellProperties represents cell properties (a:tcPr).
type CellProperties struct {
	*openxml.CompositeElementBase
}

// NewCellProperties creates a new cell properties element.
func NewCellProperties() *CellProperties {
	elem := openxml.NewCompositeElement(
		NamespaceMain,
		"tcPr",
		PrefixMain,
	)

	return &CellProperties{CompositeElementBase: elem}
}

// MarginLeft returns the left margin in EMUs.
func (cp *CellProperties) MarginLeft() int64 {
	attr, found := cp.GetAttribute("marL", "")
	if !found {
		return 0
	}
	margin, _ := strconv.ParseInt(attr.Value(), base10Cell, 64) //nolint:revive // 64-bit integer is standard

	return margin
}

// SetMarginLeft sets the left margin in EMUs.
func (cp *CellProperties) SetMarginLeft(marginEMU int64) {
	if marginEMU == 0 {
		cp.RemoveAttribute("marL", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marL",
			"",
			strconv.FormatInt(marginEMU, base10Cell),
		),
	)
}

// MarginRight returns the right margin in EMUs.
func (cp *CellProperties) MarginRight() int64 {
	attr, found := cp.GetAttribute("marR", "")
	if !found {
		return 0
	}
	margin, _ := strconv.ParseInt(attr.Value(), base10Cell, 64) //nolint:revive // 64-bit integer is standard

	return margin
}

// SetMarginRight sets the right margin in EMUs.
func (cp *CellProperties) SetMarginRight(marginEMU int64) {
	if marginEMU == 0 {
		cp.RemoveAttribute("marR", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marR",
			"",
			strconv.FormatInt(marginEMU, base10Cell),
		),
	)
}

// MarginTop returns the top margin in EMUs.
func (cp *CellProperties) MarginTop() int64 {
	attr, found := cp.GetAttribute("marT", "")
	if !found {
		return 0
	}
	margin, _ := strconv.ParseInt(attr.Value(), base10Cell, 64) //nolint:revive // 64-bit integer is standard

	return margin
}

// SetMarginTop sets the top margin in EMUs.
func (cp *CellProperties) SetMarginTop(marginEMU int64) {
	if marginEMU == 0 {
		cp.RemoveAttribute("marT", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marT",
			"",
			strconv.FormatInt(marginEMU, base10Cell),
		),
	)
}

// MarginBottom returns the bottom margin in EMUs.
func (cp *CellProperties) MarginBottom() int64 {
	attr, found := cp.GetAttribute("marB", "")
	if !found {
		return 0
	}
	margin, _ := strconv.ParseInt(attr.Value(), base10Cell, 64) //nolint:revive // 64-bit integer is standard

	return margin
}

// SetMarginBottom sets the bottom margin in EMUs.
func (cp *CellProperties) SetMarginBottom(marginEMU int64) {
	if marginEMU == 0 {
		cp.RemoveAttribute("marB", "")

		return
	}
	cp.SetAttribute(
		openxml.NewAttribute(
			"",
			"marB",
			"",
			strconv.FormatInt(marginEMU, base10Cell),
		),
	)
}

// Clone creates a deep copy of this CellProperties element.
func (cp *CellProperties) Clone() openxml.Element {
	return &CellProperties{
		CompositeElementBase: cp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
