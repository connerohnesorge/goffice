//nolint:revive // This file contains shape element implementation.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// Shape represents a shape element (p:sp).
type Shape struct {
	*openxml.CompositeElementBase
}

// NewShape creates a new Shape element.
func NewShape() *Shape {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"sp",
		PrefixP,
	)
	sp := &Shape{CompositeElementBase: elem}

	// Add required non-visual shape properties
	sp.AppendChild(NewNonVisualShapeProperties())
	// Add required shape properties (DrawingML)
	sp.AppendChild(NewShapeProperties())
	// Note: Text body is created lazily via GetOrCreateTextBody() when needed

	return sp
}

// NonVisualShapeProperties returns the non-visual shape properties.
func (sp *Shape) NonVisualShapeProperties() *NonVisualShapeProperties {
	elem := sp.GetElement(
		"nvSpPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if nvsp, ok := elem.(*NonVisualShapeProperties); ok {
		return nvsp
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &NonVisualShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// ShapeProperties returns the shape properties (a:spPr).
func (sp *Shape) ShapeProperties() *ShapeProperties {
	elem := sp.GetElement(
		"spPr",
		NamespaceDrawingML,
	)
	if elem == nil {
		return nil
	}
	if spPr, ok := elem.(*ShapeProperties); ok {
		return spPr
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateShapeProperties returns the shape properties, creating if needed.
func (sp *Shape) GetOrCreateShapeProperties() *ShapeProperties {
	spPr := sp.ShapeProperties()
	if spPr != nil {
		return spPr
	}
	spPr = NewShapeProperties()
	// Insert after nvSpPr
	if nvsp := sp.NonVisualShapeProperties(); nvsp != nil {
		sp.InsertAfter(spPr, nvsp)
	} else {
		sp.PrependChild(spPr)
	}

	return spPr
}

// TextBody returns the text body element.
func (sp *Shape) TextBody() *drawingml.TextBody {
	elem := sp.GetElement(
		"txBody",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if tb, ok := elem.(*drawingml.TextBody); ok {
		return tb
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &drawingml.TextBody{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateTextBody returns the text body, creating if needed.
func (sp *Shape) GetOrCreateTextBody() *drawingml.TextBody {
	tb := sp.TextBody()
	if tb != nil {
		return tb
	}
	tb = NewTextBody() // Use presentation NewTextBody() which creates <p:txBody>
	sp.AppendChild(tb)

	return tb
}

// Style returns the shape style element (p:style).
func (sp *Shape) Style() *ShapeStyle {
	elem := sp.GetElement(
		"style",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if style, ok := elem.(*ShapeStyle); ok {
		return style
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ShapeStyle{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetStyle sets the shape style.
func (sp *Shape) SetStyle(style *ShapeStyle) {
	// Remove existing style
	if existing := sp.Style(); existing != nil {
		sp.RemoveChild(existing)
	}
	if style != nil {
		// Insert after spPr
		if spPr := sp.ShapeProperties(); spPr != nil {
			sp.InsertAfter(style, spPr)
		} else {
			sp.AppendChild(style)
		}
	}
}

// SetText sets the text content of the shape.
func (sp *Shape) SetText(text string) {
	tb := sp.GetOrCreateTextBody()
	tb.ClearParagraphs()
	tb.AddParagraph(text)
}

// GetText returns the plain text content of the shape.
func (sp *Shape) GetText() string {
	tb := sp.TextBody()
	if tb == nil {
		return ""
	}
	var text string
	for _, p := range tb.Paragraphs() {
		if text != "" {
			text += "\n"
		}
		text += p.GetText()
	}

	return text
}

// AddParagraph adds a paragraph with the given text.
func (sp *Shape) AddParagraph(
	text string,
) *drawingml.TextParagraph {
	tb := sp.GetOrCreateTextBody()

	return tb.AddParagraph(text)
}

// SetShapeType sets the preset geometry shape type.
func (sp *Shape) SetShapeType(
	shapeType ShapeType,
) {
	spPr := sp.GetOrCreateShapeProperties()
	spPr.SetPresetGeometry(string(shapeType))
}

// SetPosition sets the position of the shape in EMUs.
func (sp *Shape) SetPosition(x, y int) {
	spPr := sp.GetOrCreateShapeProperties()
	spPr.SetOffset(x, y)
}

// SetSize sets the size of the shape in EMUs.
func (sp *Shape) SetSize(cx, cy int) {
	spPr := sp.GetOrCreateShapeProperties()
	spPr.SetExtents(cx, cy)
}

// SetSolidFill sets a solid color fill.
func (sp *Shape) SetSolidFill(hexColor string) {
	spPr := sp.GetOrCreateShapeProperties()
	spPr.SetSolidFill(hexColor)
}

// SetNoFill removes any fill from the shape.
func (sp *Shape) SetNoFill() {
	spPr := sp.GetOrCreateShapeProperties()
	spPr.SetNoFill()
}

// Clone creates a deep copy of this Shape element.
func (sp *Shape) Clone() openxml.Element {
	cloned := sp.CompositeElementBase.Clone()

	return &Shape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// NonVisualShapeProperties (p:nvSpPr)
// ===========================================================================

// NonVisualShapeProperties represents non-visual shape properties (p:nvSpPr).
type NonVisualShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewNonVisualShapeProperties creates a new NonVisualShapeProperties element.
func NewNonVisualShapeProperties() *NonVisualShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvSpPr",
		PrefixP,
	)
	nvsp := &NonVisualShapeProperties{
		CompositeElementBase: elem,
	}

	// Add required cNvPr (common non-visual properties)
	cNvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvPr",
		PrefixP,
	)
	id := nextShapeId()
	cNvPr.SetAttribute(
		openxml.NewAttribute(
			"",
			"id",
			"",
			strconv.FormatUint(uint64(id), 10),
		),
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			"Shape "+strconv.FormatUint(
				uint64(id),
				10,
			),
		),
	)
	nvsp.AppendChild(cNvPr)

	// Add required cNvSpPr (common non-visual shape properties)
	cNvSpPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvSpPr",
		PrefixP,
	)
	nvsp.AppendChild(cNvSpPr)

	// Add required nvPr (non-visual properties)
	nvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPr",
		PrefixP,
	)
	nvsp.AppendChild(nvPr)

	return nvsp
}

// Id returns the shape ID.
func (nvsp *NonVisualShapeProperties) Id() uint32 {
	cNvPr := nvsp.GetElement(
		"cNvPr",
		NamespacePresentationML,
	)
	if cNvPr == nil {
		return 0
	}
	attr, found := cNvPr.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseUint(
		attr.Value(),
		10,
		32,
	)

	return uint32(val)
}

// SetId sets the shape ID.
func (nvsp *NonVisualShapeProperties) SetId(
	id uint32,
) {
	cNvPr := nvsp.GetElement(
		"cNvPr",
		NamespacePresentationML,
	)
	if cNvPr == nil {
		return
	}
	cNvPr.SetAttribute(openxml.NewAttribute(
		"",
		"id",
		"",
		strconv.FormatUint(uint64(id), 10),
	))
}

// Name returns the shape name.
func (nvsp *NonVisualShapeProperties) Name() string {
	cNvPr := nvsp.GetElement(
		"cNvPr",
		NamespacePresentationML,
	)
	if cNvPr == nil {
		return ""
	}
	attr, found := cNvPr.GetAttribute("name", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetName sets the shape name.
func (nvsp *NonVisualShapeProperties) SetName(
	name string,
) {
	cNvPr := nvsp.GetElement(
		"cNvPr",
		NamespacePresentationML,
	)
	if cNvPr == nil {
		return
	}
	cNvPr.SetAttribute(openxml.NewAttribute(
		"",
		"name",
		"",
		name,
	))
}

// SetPlaceholder sets the placeholder type.
func (nvsp *NonVisualShapeProperties) SetPlaceholder(
	phType PlaceholderType,
	idx int,
) {
	nvPrElem := nvsp.GetElement(
		"nvPr",
		NamespacePresentationML,
	)
	if nvPrElem == nil {
		return
	}
	nvPr, ok := nvPrElem.(*openxml.CompositeElementBase)
	if !ok {
		return
	}

	// Remove existing placeholder
	if existing := nvPr.GetElement("ph", NamespacePresentationML); existing != nil {
		nvPr.RemoveChild(existing)
	}

	// Create new placeholder element
	ph := openxml.NewLeafElement(
		NamespacePresentationML,
		"ph",
		PrefixP,
	)
	if phType != "" {
		ph.SetAttribute(
			openxml.NewAttribute(
				"",
				"type",
				"",
				string(phType),
			),
		)
	}
	if idx >= 0 {
		ph.SetAttribute(
			openxml.NewAttribute(
				"",
				"idx",
				"",
				strconv.Itoa(idx),
			),
		)
	}
	nvPr.AppendChild(ph)
}

// Clone creates a deep copy of this NonVisualShapeProperties element.
func (nvsp *NonVisualShapeProperties) Clone() openxml.Element {
	cloned := nvsp.CompositeElementBase.Clone()

	return &NonVisualShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// ShapeProperties (a:spPr)
// ===========================================================================

// ShapeProperties represents shape properties (a:spPr).
type ShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewShapeProperties creates a new ShapeProperties element.
func NewShapeProperties() *ShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"spPr",
		PrefixA,
	)

	return &ShapeProperties{
		CompositeElementBase: elem,
	}
}

// SetOffset sets the position offset in EMUs.
func (spPr *ShapeProperties) SetOffset(x, y int) {
	xfrm := spPr.getOrCreateTransform()
	offElem := xfrm.GetElement(
		"off",
		NamespaceDrawingML,
	)
	var off *openxml.LeafElementBase
	if offElem != nil {
		off, _ = offElem.(*openxml.LeafElementBase)
	}
	if off == nil {
		off = openxml.NewLeafElement(
			NamespaceDrawingML,
			"off",
			PrefixA,
		)
		// Insert as first child
		if first := xfrm.FirstChild(); first != nil {
			xfrm.InsertBefore(off, first)
		} else {
			xfrm.AppendChild(off)
		}
	}
	off.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.Itoa(x),
		),
	)
	off.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.Itoa(y),
		),
	)
}

// SetExtents sets the size extents in EMUs.
func (spPr *ShapeProperties) SetExtents(
	cx, cy int,
) {
	xfrm := spPr.getOrCreateTransform()
	extElem := xfrm.GetElement(
		"ext",
		NamespaceDrawingML,
	)
	var ext *openxml.LeafElementBase
	if extElem != nil {
		ext, _ = extElem.(*openxml.LeafElementBase)
	}
	if ext == nil {
		ext = openxml.NewLeafElement(
			NamespaceDrawingML,
			"ext",
			PrefixA,
		)
		xfrm.AppendChild(ext)
	}
	ext.SetAttribute(
		openxml.NewAttribute(
			"",
			"cx",
			"",
			strconv.Itoa(cx),
		),
	)
	ext.SetAttribute(
		openxml.NewAttribute(
			"",
			"cy",
			"",
			strconv.Itoa(cy),
		),
	)
}

// SetPresetGeometry sets the preset geometry type.
func (spPr *ShapeProperties) SetPresetGeometry(
	prst string,
) {
	// Remove existing geometry
	if existing := spPr.GetElement("prstGeom", NamespaceDrawingML); existing != nil {
		spPr.RemoveChild(existing)
	}
	if existing := spPr.GetElement("custGeom", NamespaceDrawingML); existing != nil {
		spPr.RemoveChild(existing)
	}

	prstGeom := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"prstGeom",
		PrefixA,
	)
	prstGeom.SetAttribute(
		openxml.NewAttribute(
			"",
			"prst",
			"",
			prst,
		),
	)
	// Add empty avLst (adjustment list)
	avLst := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"avLst",
		PrefixA,
	)
	prstGeom.AppendChild(avLst)
	spPr.AppendChild(prstGeom)
}

// SetSolidFill sets a solid color fill.
func (spPr *ShapeProperties) SetSolidFill(
	hexColor string,
) {
	spPr.removeFill()
	solidFill := drawingml.NewSolidFillWithRgb(
		hexColor,
	)
	spPr.AppendChild(solidFill)
}

// SetNoFill sets no fill.
func (spPr *ShapeProperties) SetNoFill() {
	spPr.removeFill()
	noFill := drawingml.NewNoFill()
	spPr.AppendChild(noFill)
}

// removeFill removes any existing fill elements.
func (spPr *ShapeProperties) removeFill() {
	if nf := spPr.GetElement("noFill", NamespaceDrawingML); nf != nil {
		spPr.RemoveChild(nf)
	}
	if sf := spPr.GetElement("solidFill", NamespaceDrawingML); sf != nil {
		spPr.RemoveChild(sf)
	}
	if gf := spPr.GetElement("gradFill", NamespaceDrawingML); gf != nil {
		spPr.RemoveChild(gf)
	}
	if pf := spPr.GetElement("pattFill", NamespaceDrawingML); pf != nil {
		spPr.RemoveChild(pf)
	}
	if bf := spPr.GetElement("blipFill", NamespaceDrawingML); bf != nil {
		spPr.RemoveChild(bf)
	}
}

// getOrCreateTransform returns the transform element, creating if needed.
func (spPr *ShapeProperties) getOrCreateTransform() *openxml.CompositeElementBase {
	xfrmElem := spPr.GetElement(
		"xfrm",
		NamespaceDrawingML,
	)
	if xfrmElem != nil {
		if xfrm, ok := xfrmElem.(*openxml.CompositeElementBase); ok {
			return xfrm
		}
	}
	xfrm := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"xfrm",
		PrefixA,
	)
	// Insert as first child
	if first := spPr.FirstChild(); first != nil {
		spPr.InsertBefore(xfrm, first)
	} else {
		spPr.AppendChild(xfrm)
	}

	return xfrm
}

// Clone creates a deep copy of this ShapeProperties element.
func (spPr *ShapeProperties) Clone() openxml.Element {
	cloned := spPr.CompositeElementBase.Clone()

	return &ShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// ShapeStyle (p:style)
// ===========================================================================

// ShapeStyle represents a shape style element (p:style).
type ShapeStyle struct {
	*openxml.CompositeElementBase
}

// NewShapeStyle creates a new ShapeStyle element.
func NewShapeStyle() *ShapeStyle {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"style",
		PrefixP,
	)

	return &ShapeStyle{CompositeElementBase: elem}
}

// Clone creates a deep copy of this ShapeStyle element.
func (ss *ShapeStyle) Clone() openxml.Element {
	cloned := ss.CompositeElementBase.Clone()

	return &ShapeStyle{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
