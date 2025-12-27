//nolint:revive // comments-density
package elements

import (
	"github.com/connerohnesorge/goffice/openxml"
)

// Worksheet represents the worksheet root element (x:worksheet).
type Worksheet struct {
	*openxml.CompositeElementBase
}

// NewWorksheet creates a new Worksheet element.
func NewWorksheet() *Worksheet {
	elem := openxml.NewCompositeElement(
		NamespaceSML,
		"worksheet",
		PrefixDefault,
	)

	return &Worksheet{CompositeElementBase: elem}
}

// SheetPr returns the sheet properties element, or nil if not present.
func (ws *Worksheet) SheetPr() *SheetPr {
	elem := ws.GetElement("sheetPr", NamespaceSML)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*SheetPr); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SheetPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSheetPr returns the sheet properties, creating if needed.
func (ws *Worksheet) GetOrCreateSheetPr() *SheetPr {
	sp := ws.SheetPr()
	if sp != nil {
		return sp
	}
	sp = NewSheetPr()
	// SheetPr should be first
	if first := ws.FirstChild(); first != nil {
		ws.InsertBefore(sp, first)
	} else {
		ws.AppendChild(sp)
	}

	return sp
}

// Dimension returns the dimension element, or nil if not present.
func (ws *Worksheet) Dimension() *Dimension {
	elem := ws.GetElement(
		"dimension",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if d, ok := elem.(*Dimension); ok {
		return d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Dimension{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDimension returns the dimension element, creating if needed.
func (ws *Worksheet) GetOrCreateDimension() *Dimension {
	d := ws.Dimension()
	if d != nil {
		return d
	}
	d = NewDimension()
	// Insert after sheetPr if present
	sp := ws.SheetPr()
	if sp != nil {
		ws.InsertAfter(d, sp)
	} else if first := ws.FirstChild(); first != nil {
		ws.InsertBefore(d, first)
	} else {
		ws.AppendChild(d)
	}

	return d
}

// SheetViews returns the sheet views element, or nil if not present.
func (ws *Worksheet) SheetViews() *SheetViews {
	elem := ws.GetElement(
		"sheetViews",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if sv, ok := elem.(*SheetViews); ok {
		return sv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SheetViews{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSheetViews returns the sheet views, creating if needed.
func (ws *Worksheet) GetOrCreateSheetViews() *SheetViews {
	sv := ws.SheetViews()
	if sv != nil {
		return sv
	}
	sv = NewSheetViews()
	ws.AppendChild(sv)

	return sv
}

// SheetFormatPr returns the sheet format properties, or nil if not present.
func (ws *Worksheet) SheetFormatPr() *SheetFormatPr {
	elem := ws.GetElement(
		"sheetFormatPr",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if sfp, ok := elem.(*SheetFormatPr); ok {
		return sfp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SheetFormatPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSheetFormatPr returns the sheet format properties, creating if needed.
func (ws *Worksheet) GetOrCreateSheetFormatPr() *SheetFormatPr {
	sfp := ws.SheetFormatPr()
	if sfp != nil {
		return sfp
	}
	sfp = NewSheetFormatPr()
	ws.AppendChild(sfp)

	return sfp
}

// Cols returns the columns element, or nil if not present.
func (ws *Worksheet) Cols() *Cols {
	elem := ws.GetElement("cols", NamespaceSML)
	if elem == nil {
		return nil
	}
	if c, ok := elem.(*Cols); ok {
		return c
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Cols{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCols returns the columns element, creating if needed.
func (ws *Worksheet) GetOrCreateCols() *Cols {
	c := ws.Cols()
	if c != nil {
		return c
	}
	c = NewCols()
	ws.AppendChild(c)

	return c
}

// SheetData returns the sheet data element, or nil if not present.
func (ws *Worksheet) SheetData() *SheetData {
	elem := ws.GetElement(
		"sheetData",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if sd, ok := elem.(*SheetData); ok {
		return sd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SheetData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSheetData returns the sheet data, creating if needed.
func (ws *Worksheet) GetOrCreateSheetData() *SheetData {
	sd := ws.SheetData()
	if sd != nil {
		return sd
	}
	sd = NewSheetData()
	ws.AppendChild(sd)

	return sd
}

// SheetProtection returns the sheet protection element, or nil if not present.
func (ws *Worksheet) SheetProtection() *SheetProtection {
	elem := ws.GetElement(
		"sheetProtection",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*SheetProtection); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &SheetProtection{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSheetProtection returns the sheet protection, creating if needed.
func (ws *Worksheet) GetOrCreateSheetProtection() *SheetProtection {
	sp := ws.SheetProtection()
	if sp != nil {
		return sp
	}
	sp = NewSheetProtection()
	ws.AppendChild(sp)

	return sp
}

// MergeCells returns the merge cells element, or nil if not present.
func (ws *Worksheet) MergeCells() *MergeCells {
	elem := ws.GetElement(
		"mergeCells",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if mc, ok := elem.(*MergeCells); ok {
		return mc
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &MergeCells{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateMergeCells returns the merge cells, creating if needed.
func (ws *Worksheet) GetOrCreateMergeCells() *MergeCells {
	mc := ws.MergeCells()
	if mc != nil {
		return mc
	}
	mc = NewMergeCells()
	ws.AppendChild(mc)

	return mc
}

// ConditionalFormatting returns the conditional formatting elements.
// Note: There can be multiple conditionalFormatting elements.
func (ws *Worksheet) ConditionalFormatting() []*ConditionalFormatting {
	var result []*ConditionalFormatting
	for child := range ws.Children() {
		if child.LocalName() != "conditionalFormatting" ||
			child.NamespaceURI() != NamespaceSML {
			continue
		}
		if cf, ok := child.(*ConditionalFormatting); ok {
			result = append(result, cf)
		} else if comp, ok := child.(*openxml.CompositeElementBase); ok {
			result = append(result, &ConditionalFormatting{CompositeElementBase: comp})
		}
	}

	return result
}

// AddConditionalFormatting adds a new conditional formatting element.
func (ws *Worksheet) AddConditionalFormatting() *ConditionalFormatting {
	cf := NewConditionalFormatting()
	ws.AppendChild(cf)

	return cf
}

// DataValidations returns the data validations element, or nil if not present.
func (ws *Worksheet) DataValidations() *DataValidations {
	elem := ws.GetElement(
		"dataValidations",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if dv, ok := elem.(*DataValidations); ok {
		return dv
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DataValidations{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDataValidations returns the data validations, creating if needed.
func (ws *Worksheet) GetOrCreateDataValidations() *DataValidations {
	dv := ws.DataValidations()
	if dv != nil {
		return dv
	}
	dv = NewDataValidations()
	ws.AppendChild(dv)

	return dv
}

// Hyperlinks returns the hyperlinks element, or nil if not present.
func (ws *Worksheet) Hyperlinks() *Hyperlinks {
	elem := ws.GetElement(
		"hyperlinks",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if h, ok := elem.(*Hyperlinks); ok {
		return h
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Hyperlinks{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateHyperlinks returns the hyperlinks, creating if needed.
func (ws *Worksheet) GetOrCreateHyperlinks() *Hyperlinks {
	h := ws.Hyperlinks()
	if h != nil {
		return h
	}
	h = NewHyperlinks()
	ws.AppendChild(h)

	return h
}

// PrintOptions returns the print options element, or nil if not present.
func (ws *Worksheet) PrintOptions() *PrintOptions {
	elem := ws.GetElement(
		"printOptions",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if po, ok := elem.(*PrintOptions); ok {
		return po
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PrintOptions{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePrintOptions returns the print options, creating if needed.
func (ws *Worksheet) GetOrCreatePrintOptions() *PrintOptions {
	po := ws.PrintOptions()
	if po != nil {
		return po
	}
	po = NewPrintOptions()
	ws.AppendChild(po)

	return po
}

// PageMargins returns the page margins element, or nil if not present.
func (ws *Worksheet) PageMargins() *PageMargins {
	elem := ws.GetElement(
		"pageMargins",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if pm, ok := elem.(*PageMargins); ok {
		return pm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageMargins{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageMargins returns the page margins, creating if needed.
func (ws *Worksheet) GetOrCreatePageMargins() *PageMargins {
	pm := ws.PageMargins()
	if pm != nil {
		return pm
	}
	pm = NewPageMargins()
	ws.AppendChild(pm)

	return pm
}

// PageSetup returns the page setup element, or nil if not present.
func (ws *Worksheet) PageSetup() *PageSetup {
	elem := ws.GetElement(
		"pageSetup",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ps, ok := elem.(*PageSetup); ok {
		return ps
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PageSetup{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePageSetup returns the page setup, creating if needed.
func (ws *Worksheet) GetOrCreatePageSetup() *PageSetup {
	ps := ws.PageSetup()
	if ps != nil {
		return ps
	}
	ps = NewPageSetup()
	ws.AppendChild(ps)

	return ps
}

// HeaderFooter returns the header/footer element, or nil if not present.
func (ws *Worksheet) HeaderFooter() *HeaderFooter {
	elem := ws.GetElement(
		"headerFooter",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if hf, ok := elem.(*HeaderFooter); ok {
		return hf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &HeaderFooter{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateHeaderFooter returns the header/footer, creating if needed.
func (ws *Worksheet) GetOrCreateHeaderFooter() *HeaderFooter {
	hf := ws.HeaderFooter()
	if hf != nil {
		return hf
	}
	hf = NewHeaderFooter()
	ws.AppendChild(hf)

	return hf
}

// RowBreaks returns the row breaks element, or nil if not present.
func (ws *Worksheet) RowBreaks() *RowBreaks {
	elem := ws.GetElement(
		"rowBreaks",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if rb, ok := elem.(*RowBreaks); ok {
		return rb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &RowBreaks{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateRowBreaks returns the row breaks, creating if needed.
func (ws *Worksheet) GetOrCreateRowBreaks() *RowBreaks {
	rb := ws.RowBreaks()
	if rb != nil {
		return rb
	}
	rb = NewRowBreaks()
	ws.AppendChild(rb)

	return rb
}

// ColBreaks returns the column breaks element, or nil if not present.
func (ws *Worksheet) ColBreaks() *ColBreaks {
	elem := ws.GetElement(
		"colBreaks",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if cb, ok := elem.(*ColBreaks); ok {
		return cb
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ColBreaks{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateColBreaks returns the column breaks, creating if needed.
func (ws *Worksheet) GetOrCreateColBreaks() *ColBreaks {
	cb := ws.ColBreaks()
	if cb != nil {
		return cb
	}
	cb = NewColBreaks()
	ws.AppendChild(cb)

	return cb
}

// Drawing returns the drawing element, or nil if not present.
func (ws *Worksheet) Drawing() *Drawing {
	elem := ws.GetElement("drawing", NamespaceSML)
	if elem == nil {
		return nil
	}
	if d, ok := elem.(*Drawing); ok {
		return d
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Drawing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateDrawing returns the drawing element, creating if needed.
func (ws *Worksheet) GetOrCreateDrawing() *Drawing {
	d := ws.Drawing()
	if d != nil {
		return d
	}
	d = NewDrawing()
	ws.AppendChild(d)

	return d
}

// LegacyDrawing returns the legacy drawing element, or nil if not present.
func (ws *Worksheet) LegacyDrawing() *LegacyDrawing {
	elem := ws.GetElement(
		"legacyDrawing",
		NamespaceSML,
	)
	if elem == nil {
		return nil
	}
	if ld, ok := elem.(*LegacyDrawing); ok {
		return ld
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &LegacyDrawing{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateLegacyDrawing returns the legacy drawing, creating if needed.
func (ws *Worksheet) GetOrCreateLegacyDrawing() *LegacyDrawing {
	ld := ws.LegacyDrawing()
	if ld != nil {
		return ld
	}
	ld = NewLegacyDrawing()
	ws.AppendChild(ld)

	return ld
}

// Picture returns the picture element (for background), or nil if not present.
func (ws *Worksheet) Picture() *Picture {
	elem := ws.GetElement("picture", NamespaceSML)
	if elem == nil {
		return nil
	}
	if p, ok := elem.(*Picture); ok {
		return p
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Picture{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreatePicture returns the picture element, creating if needed.
func (ws *Worksheet) GetOrCreatePicture() *Picture {
	p := ws.Picture()
	if p != nil {
		return p
	}
	p = NewPicture()
	ws.AppendChild(p)

	return p
}

// Clone creates a deep copy of this Worksheet element.
func (ws *Worksheet) Clone() openxml.Element {
	cloned := ws.CompositeElementBase.Clone()

	return &Worksheet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Worksheet element.
func (ws *Worksheet) CloneNode(
	deep bool,
) openxml.Element {
	cloned := ws.CompositeElementBase.CloneNode(
		deep,
	)

	return &Worksheet{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
