package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Drawing namespace constants
const (
	// NamespaceDrawingMLWordprocessing is the DrawingML WordprocessingDrawing namespace.
	NamespaceDrawingMLWordprocessing = openxml.NamespaceDrawingMLWordprocessing
	// PrefixWP is the prefix for WordprocessingDrawing elements.
	PrefixWP = "wp"
	// NamespaceDrawingML is the main DrawingML namespace.
	NamespaceDrawingML = openxml.NamespaceDrawingML
	// PrefixA is the prefix for DrawingML elements.
	PrefixA = "a"
	// NamespaceDrawingMLPicture is the DrawingML Picture namespace.
	NamespaceDrawingMLPicture = openxml.NamespaceDrawingMLPicture
	// PrefixPic is the prefix for Picture elements.
	PrefixPic = "pic"
	// NamespaceRelationships is the relationships namespace.
	NamespaceRelationships = openxml.NamespaceRelationships
	// PrefixR is the prefix for relationship elements.
	PrefixR = "r"
)

// EMU (English Metric Units) constants
const (
	// EMUsPerInch is the number of EMUs per inch.
	EMUsPerInch int64 = 914400
	// EMUsPerPoint is the number of EMUs per point.
	EMUsPerPoint int64 = 12700
	// EMUsPerCm is the number of EMUs per centimeter.
	EMUsPerCm int64 = 360000
	// EMUsPerMm is the number of EMUs per millimeter.
	EMUsPerMm int64 = 36000
	// EMUsPerPixel is the number of EMUs per pixel at 96 DPI.
	EMUsPerPixel int64 = 9525
)

// Drawing represents the w:drawing element that contains inline or anchor drawings.
type Drawing struct {
	*openxml.CompositeElementBase
}

// NewDrawing creates a new empty Drawing element.
func NewDrawing() *Drawing {
	elem := openxml.NewCompositeElement(NamespaceWML, "drawing", PrefixW)
	return &Drawing{CompositeElementBase: elem}
}

// NewInlineDrawing creates a new Drawing with an inline drawing.
func NewInlineDrawing(width, height int64, relId string) *Drawing {
	d := NewDrawing()
	inline := NewInlineDrawingElement(width, height, relId)
	d.AppendChild(inline)
	return d
}

// NewAnchorDrawing creates a new Drawing with an anchor drawing.
func NewAnchorDrawing(width, height int64, relId string) *Drawing {
	d := NewDrawing()
	anchor := NewAnchorDrawingElement(width, height, relId)
	d.AppendChild(anchor)
	return d
}

// Inline returns the inline drawing element, or nil if not present.
func (d *Drawing) Inline() *InlineDrawing {
	elem := d.GetElement("inline", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if inline, ok := elem.(*InlineDrawing); ok {
		return inline
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &InlineDrawing{CompositeElementBase: comp}
	}
	return nil
}

// Anchor returns the anchor drawing element, or nil if not present.
func (d *Drawing) Anchor() *AnchorDrawing {
	elem := d.GetElement("anchor", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if anchor, ok := elem.(*AnchorDrawing); ok {
		return anchor
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &AnchorDrawing{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this Drawing element.
func (d *Drawing) Clone() openxml.Element {
	return &Drawing{
		CompositeElementBase: d.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Drawing element.
func (d *Drawing) CloneNode(deep bool) openxml.Element {
	return &Drawing{
		CompositeElementBase: d.CompositeElementBase.CloneNode(deep).(*openxml.CompositeElementBase),
	}
}

// InlineDrawing represents the wp:inline element for inline drawings.
type InlineDrawing struct {
	*openxml.CompositeElementBase
}

// NewInlineDrawingElement creates a new InlineDrawing element.
func NewInlineDrawingElement(width, height int64, relId string) *InlineDrawing {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "inline", PrefixWP)
	inline := &InlineDrawing{CompositeElementBase: elem}

	// Set distance from text attributes (default 0)
	inline.SetAttribute(openxml.NewAttribute("", "distT", "", "0"))
	inline.SetAttribute(openxml.NewAttribute("", "distB", "", "0"))
	inline.SetAttribute(openxml.NewAttribute("", "distL", "", "0"))
	inline.SetAttribute(openxml.NewAttribute("", "distR", "", "0"))

	// Add extent element
	extent := NewExtent(width, height)
	inline.AppendChild(extent)

	// Add effect extent (empty/default)
	effectExtent := NewEffectExtent(0, 0, 0, 0)
	inline.AppendChild(effectExtent)

	// Add doc properties
	docPr := NewDocProperties(1, "Picture")
	inline.AppendChild(docPr)

	// Add non-visual graphic frame properties
	cNvGraphicFramePr := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "cNvGraphicFramePr", PrefixWP)
	graphicFrameLocks := openxml.NewCompositeElement(NamespaceDrawingML, "graphicFrameLocks", PrefixA)
	graphicFrameLocks.SetAttribute(openxml.NewAttribute("", "noChangeAspect", "", "1"))
	cNvGraphicFramePr.AppendChild(graphicFrameLocks)
	inline.AppendChild(cNvGraphicFramePr)

	// Add graphic with picture
	graphic := NewGraphicWithPicture(relId, width, height)
	inline.AppendChild(graphic)

	return inline
}

// Width returns the width in EMUs.
func (i *InlineDrawing) Width() int64 {
	extent := i.Extent()
	if extent != nil {
		return extent.Width()
	}
	return 0
}

// Height returns the height in EMUs.
func (i *InlineDrawing) Height() int64 {
	extent := i.Extent()
	if extent != nil {
		return extent.Height()
	}
	return 0
}

// Extent returns the extent element.
func (i *InlineDrawing) Extent() *Extent {
	elem := i.GetElement("extent", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if ext, ok := elem.(*Extent); ok {
		return ext
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Extent{CompositeElementBase: comp}
	}
	return nil
}

// DocProperties returns the doc properties element.
func (i *InlineDrawing) DocProperties() *DocProperties {
	elem := i.GetElement("docPr", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if dp, ok := elem.(*DocProperties); ok {
		return dp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DocProperties{CompositeElementBase: comp}
	}
	return nil
}

// Graphic returns the graphic element.
func (i *InlineDrawing) Graphic() *Graphic {
	elem := i.GetElement("graphic", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if g, ok := elem.(*Graphic); ok {
		return g
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Graphic{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this InlineDrawing element.
func (i *InlineDrawing) Clone() openxml.Element {
	return &InlineDrawing{
		CompositeElementBase: i.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// AnchorDrawing represents the wp:anchor element for floating drawings.
type AnchorDrawing struct {
	*openxml.CompositeElementBase
}

// WrapType represents the text wrapping type for anchor drawings.
type WrapType string

const (
	// WrapTypeNone indicates no text wrapping.
	WrapTypeNone WrapType = "none"
	// WrapTypeSquare indicates square text wrapping.
	WrapTypeSquare WrapType = "square"
	// WrapTypeTight indicates tight text wrapping.
	WrapTypeTight WrapType = "tight"
	// WrapTypeThrough indicates through text wrapping.
	WrapTypeThrough WrapType = "through"
	// WrapTypeTopAndBottom indicates top and bottom text wrapping.
	WrapTypeTopAndBottom WrapType = "topAndBottom"
)

// RelativeFromH represents horizontal relative positioning.
type RelativeFromH string

const (
	// RelativeFromHCharacter positions relative to character.
	RelativeFromHCharacter RelativeFromH = "character"
	// RelativeFromHColumn positions relative to column.
	RelativeFromHColumn RelativeFromH = "column"
	// RelativeFromHInsideMargin positions relative to inside margin.
	RelativeFromHInsideMargin RelativeFromH = "insideMargin"
	// RelativeFromHLeftMargin positions relative to left margin.
	RelativeFromHLeftMargin RelativeFromH = "leftMargin"
	// RelativeFromHMargin positions relative to margin.
	RelativeFromHMargin RelativeFromH = "margin"
	// RelativeFromHOutsideMargin positions relative to outside margin.
	RelativeFromHOutsideMargin RelativeFromH = "outsideMargin"
	// RelativeFromHPage positions relative to page.
	RelativeFromHPage RelativeFromH = "page"
	// RelativeFromHRightMargin positions relative to right margin.
	RelativeFromHRightMargin RelativeFromH = "rightMargin"
)

// RelativeFromV represents vertical relative positioning.
type RelativeFromV string

const (
	// RelativeFromVBottomMargin positions relative to bottom margin.
	RelativeFromVBottomMargin RelativeFromV = "bottomMargin"
	// RelativeFromVInsideMargin positions relative to inside margin.
	RelativeFromVInsideMargin RelativeFromV = "insideMargin"
	// RelativeFromVLine positions relative to line.
	RelativeFromVLine RelativeFromV = "line"
	// RelativeFromVMargin positions relative to margin.
	RelativeFromVMargin RelativeFromV = "margin"
	// RelativeFromVOutsideMargin positions relative to outside margin.
	RelativeFromVOutsideMargin RelativeFromV = "outsideMargin"
	// RelativeFromVPage positions relative to page.
	RelativeFromVPage RelativeFromV = "page"
	// RelativeFromVParagraph positions relative to paragraph.
	RelativeFromVParagraph RelativeFromV = "paragraph"
	// RelativeFromVTopMargin positions relative to top margin.
	RelativeFromVTopMargin RelativeFromV = "topMargin"
)

// NewAnchorDrawingElement creates a new AnchorDrawing element.
func NewAnchorDrawingElement(width, height int64, relId string) *AnchorDrawing {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "anchor", PrefixWP)
	anchor := &AnchorDrawing{CompositeElementBase: elem}

	// Set default attributes
	anchor.SetAttribute(openxml.NewAttribute("", "distT", "", "0"))
	anchor.SetAttribute(openxml.NewAttribute("", "distB", "", "0"))
	anchor.SetAttribute(openxml.NewAttribute("", "distL", "", "114300"))
	anchor.SetAttribute(openxml.NewAttribute("", "distR", "", "114300"))
	anchor.SetAttribute(openxml.NewAttribute("", "simplePos", "", "0"))
	anchor.SetAttribute(openxml.NewAttribute("", "relativeHeight", "", "251658240"))
	anchor.SetAttribute(openxml.NewAttribute("", "behindDoc", "", "0"))
	anchor.SetAttribute(openxml.NewAttribute("", "locked", "", "0"))
	anchor.SetAttribute(openxml.NewAttribute("", "layoutInCell", "", "1"))
	anchor.SetAttribute(openxml.NewAttribute("", "allowOverlap", "", "1"))

	// Add simple position element
	simplePos := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "simplePos", PrefixWP)
	simplePos.SetAttribute(openxml.NewAttribute("", "x", "", "0"))
	simplePos.SetAttribute(openxml.NewAttribute("", "y", "", "0"))
	anchor.AppendChild(simplePos)

	// Add horizontal position
	positionH := NewPositionH(RelativeFromHColumn, 0)
	anchor.AppendChild(positionH)

	// Add vertical position
	positionV := NewPositionV(RelativeFromVParagraph, 0)
	anchor.AppendChild(positionV)

	// Add extent element
	extent := NewExtent(width, height)
	anchor.AppendChild(extent)

	// Add effect extent
	effectExtent := NewEffectExtent(0, 0, 0, 0)
	anchor.AppendChild(effectExtent)

	// Add wrap none by default
	wrapNone := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapNone", PrefixWP)
	anchor.AppendChild(wrapNone)

	// Add doc properties
	docPr := NewDocProperties(1, "Picture")
	anchor.AppendChild(docPr)

	// Add non-visual graphic frame properties
	cNvGraphicFramePr := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "cNvGraphicFramePr", PrefixWP)
	graphicFrameLocks := openxml.NewCompositeElement(NamespaceDrawingML, "graphicFrameLocks", PrefixA)
	graphicFrameLocks.SetAttribute(openxml.NewAttribute("", "noChangeAspect", "", "1"))
	cNvGraphicFramePr.AppendChild(graphicFrameLocks)
	anchor.AppendChild(cNvGraphicFramePr)

	// Add graphic with picture
	graphic := NewGraphicWithPicture(relId, width, height)
	anchor.AppendChild(graphic)

	return anchor
}

// Width returns the width in EMUs.
func (a *AnchorDrawing) Width() int64 {
	extent := a.Extent()
	if extent != nil {
		return extent.Width()
	}
	return 0
}

// Height returns the height in EMUs.
func (a *AnchorDrawing) Height() int64 {
	extent := a.Extent()
	if extent != nil {
		return extent.Height()
	}
	return 0
}

// Extent returns the extent element.
func (a *AnchorDrawing) Extent() *Extent {
	elem := a.GetElement("extent", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if ext, ok := elem.(*Extent); ok {
		return ext
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Extent{CompositeElementBase: comp}
	}
	return nil
}

// SetSimplePos sets whether simple positioning is used.
func (a *AnchorDrawing) SetSimplePos(useSimplePos bool) {
	val := "0"
	if useSimplePos {
		val = "1"
	}
	a.SetAttribute(openxml.NewAttribute("", "simplePos", "", val))
}

// SimplePos returns whether simple positioning is used.
func (a *AnchorDrawing) SimplePos() bool {
	attr, found := a.GetAttribute("simplePos", "")
	if !found {
		return false
	}
	return attr.Value() == "1"
}

// PositionH returns the horizontal position element.
func (a *AnchorDrawing) PositionH() *PositionH {
	elem := a.GetElement("positionH", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if pos, ok := elem.(*PositionH); ok {
		return pos
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PositionH{CompositeElementBase: comp}
	}
	return nil
}

// PositionV returns the vertical position element.
func (a *AnchorDrawing) PositionV() *PositionV {
	elem := a.GetElement("positionV", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if pos, ok := elem.(*PositionV); ok {
		return pos
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PositionV{CompositeElementBase: comp}
	}
	return nil
}

// SetBehindDoc sets whether the drawing is behind the document text.
func (a *AnchorDrawing) SetBehindDoc(behind bool) {
	val := "0"
	if behind {
		val = "1"
	}
	a.SetAttribute(openxml.NewAttribute("", "behindDoc", "", val))
}

// BehindDoc returns whether the drawing is behind the document text.
func (a *AnchorDrawing) BehindDoc() bool {
	attr, found := a.GetAttribute("behindDoc", "")
	if !found {
		return false
	}
	return attr.Value() == "1"
}

// SetLocked sets whether the anchor is locked.
func (a *AnchorDrawing) SetLocked(locked bool) {
	val := "0"
	if locked {
		val = "1"
	}
	a.SetAttribute(openxml.NewAttribute("", "locked", "", val))
}

// SetLayoutInCell sets whether the drawing can be laid out in a table cell.
func (a *AnchorDrawing) SetLayoutInCell(inCell bool) {
	val := "0"
	if inCell {
		val = "1"
	}
	a.SetAttribute(openxml.NewAttribute("", "layoutInCell", "", val))
}

// SetAllowOverlap sets whether the drawing can overlap other drawings.
func (a *AnchorDrawing) SetAllowOverlap(allow bool) {
	val := "0"
	if allow {
		val = "1"
	}
	a.SetAttribute(openxml.NewAttribute("", "allowOverlap", "", val))
}

// WrapNone sets no text wrapping.
func (a *AnchorDrawing) WrapNone() {
	a.removeWrapElements()
	wrapNone := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapNone", PrefixWP)
	a.insertWrapElement(wrapNone)
}

// WrapSquare sets square text wrapping.
func (a *AnchorDrawing) WrapSquare() {
	a.removeWrapElements()
	wrapSquare := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapSquare", PrefixWP)
	wrapSquare.SetAttribute(openxml.NewAttribute("", "wrapText", "", "bothSides"))
	a.insertWrapElement(wrapSquare)
}

// WrapTight sets tight text wrapping.
func (a *AnchorDrawing) WrapTight() {
	a.removeWrapElements()
	wrapTight := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapTight", PrefixWP)
	wrapTight.SetAttribute(openxml.NewAttribute("", "wrapText", "", "bothSides"))
	a.insertWrapElement(wrapTight)
}

// WrapThrough sets through text wrapping.
func (a *AnchorDrawing) WrapThrough() {
	a.removeWrapElements()
	wrapThrough := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapThrough", PrefixWP)
	wrapThrough.SetAttribute(openxml.NewAttribute("", "wrapText", "", "bothSides"))
	a.insertWrapElement(wrapThrough)
}

// WrapTopAndBottom sets top and bottom text wrapping.
func (a *AnchorDrawing) WrapTopAndBottom() {
	a.removeWrapElements()
	wrapTopAndBottom := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "wrapTopAndBottom", PrefixWP)
	a.insertWrapElement(wrapTopAndBottom)
}

// removeWrapElements removes all existing wrap elements.
func (a *AnchorDrawing) removeWrapElements() {
	wrapElements := []string{"wrapNone", "wrapSquare", "wrapTight", "wrapThrough", "wrapTopAndBottom"}
	var toRemove []openxml.Element
	for child := range a.Children() {
		for _, wrapName := range wrapElements {
			if child.LocalName() == wrapName && child.NamespaceURI() == NamespaceDrawingMLWordprocessing {
				toRemove = append(toRemove, child)
				break
			}
		}
	}
	for _, elem := range toRemove {
		a.RemoveChild(elem)
	}
}

// insertWrapElement inserts a wrap element in the correct position (after effectExtent).
func (a *AnchorDrawing) insertWrapElement(wrap openxml.Element) {
	// Find effectExtent and insert after it
	effectExtent := a.GetElement("effectExtent", NamespaceDrawingMLWordprocessing)
	if effectExtent != nil {
		a.InsertAfter(wrap, effectExtent)
	} else {
		// Find extent and insert after it
		extent := a.GetElement("extent", NamespaceDrawingMLWordprocessing)
		if extent != nil {
			a.InsertAfter(wrap, extent)
		} else {
			a.AppendChild(wrap)
		}
	}
}

// DocProperties returns the doc properties element.
func (a *AnchorDrawing) DocProperties() *DocProperties {
	elem := a.GetElement("docPr", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return nil
	}
	if dp, ok := elem.(*DocProperties); ok {
		return dp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &DocProperties{CompositeElementBase: comp}
	}
	return nil
}

// Graphic returns the graphic element.
func (a *AnchorDrawing) Graphic() *Graphic {
	elem := a.GetElement("graphic", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if g, ok := elem.(*Graphic); ok {
		return g
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Graphic{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this AnchorDrawing element.
func (a *AnchorDrawing) Clone() openxml.Element {
	return &AnchorDrawing{
		CompositeElementBase: a.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Extent represents the wp:extent element for drawing dimensions.
type Extent struct {
	*openxml.CompositeElementBase
}

// NewExtent creates a new Extent element.
func NewExtent(width, height int64) *Extent {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "extent", PrefixWP)
	ext := &Extent{CompositeElementBase: elem}
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", strconv.FormatInt(width, 10)))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", strconv.FormatInt(height, 10)))
	return ext
}

// Width returns the width in EMUs.
func (e *Extent) Width() int64 {
	attr, found := e.GetAttribute("cx", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return val
}

// SetWidth sets the width in EMUs.
func (e *Extent) SetWidth(cx int64) {
	e.SetAttribute(openxml.NewAttribute("", "cx", "", strconv.FormatInt(cx, 10)))
}

// Height returns the height in EMUs.
func (e *Extent) Height() int64 {
	attr, found := e.GetAttribute("cy", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(attr.Value(), 10, 64)
	return val
}

// SetHeight sets the height in EMUs.
func (e *Extent) SetHeight(cy int64) {
	e.SetAttribute(openxml.NewAttribute("", "cy", "", strconv.FormatInt(cy, 10)))
}

// Clone creates a deep copy of this Extent element.
func (e *Extent) Clone() openxml.Element {
	return &Extent{
		CompositeElementBase: e.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// EffectExtent represents the wp:effectExtent element.
type EffectExtent struct {
	*openxml.CompositeElementBase
}

// NewEffectExtent creates a new EffectExtent element.
func NewEffectExtent(left, top, right, bottom int64) *EffectExtent {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "effectExtent", PrefixWP)
	ee := &EffectExtent{CompositeElementBase: elem}
	ee.SetAttribute(openxml.NewAttribute("", "l", "", strconv.FormatInt(left, 10)))
	ee.SetAttribute(openxml.NewAttribute("", "t", "", strconv.FormatInt(top, 10)))
	ee.SetAttribute(openxml.NewAttribute("", "r", "", strconv.FormatInt(right, 10)))
	ee.SetAttribute(openxml.NewAttribute("", "b", "", strconv.FormatInt(bottom, 10)))
	return ee
}

// Clone creates a deep copy of this EffectExtent element.
func (ee *EffectExtent) Clone() openxml.Element {
	return &EffectExtent{
		CompositeElementBase: ee.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// DocProperties represents the wp:docPr element.
type DocProperties struct {
	*openxml.CompositeElementBase
}

// NewDocProperties creates a new DocProperties element.
func NewDocProperties(id int, name string) *DocProperties {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "docPr", PrefixWP)
	dp := &DocProperties{CompositeElementBase: elem}
	dp.SetAttribute(openxml.NewAttribute("", "id", "", strconv.Itoa(id)))
	dp.SetAttribute(openxml.NewAttribute("", "name", "", name))
	return dp
}

// Id returns the id attribute.
func (dp *DocProperties) Id() int {
	attr, found := dp.GetAttribute("id", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetId sets the id attribute.
func (dp *DocProperties) SetId(id int) {
	dp.SetAttribute(openxml.NewAttribute("", "id", "", strconv.Itoa(id)))
}

// Name returns the name attribute.
func (dp *DocProperties) Name() string {
	attr, found := dp.GetAttribute("name", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// SetName sets the name attribute.
func (dp *DocProperties) SetName(name string) {
	dp.SetAttribute(openxml.NewAttribute("", "name", "", name))
}

// Description returns the description attribute.
func (dp *DocProperties) Description() string {
	attr, found := dp.GetAttribute("descr", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// SetDescription sets the description attribute.
func (dp *DocProperties) SetDescription(descr string) {
	dp.SetAttribute(openxml.NewAttribute("", "descr", "", descr))
}

// Clone creates a deep copy of this DocProperties element.
func (dp *DocProperties) Clone() openxml.Element {
	return &DocProperties{
		CompositeElementBase: dp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PositionH represents the wp:positionH element.
type PositionH struct {
	*openxml.CompositeElementBase
}

// NewPositionH creates a new PositionH element.
func NewPositionH(relativeFrom RelativeFromH, offset int64) *PositionH {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "positionH", PrefixWP)
	pos := &PositionH{CompositeElementBase: elem}
	pos.SetAttribute(openxml.NewAttribute("", "relativeFrom", "", string(relativeFrom)))

	// Add posOffset child
	posOffset := openxml.NewLeafElementWithText(NamespaceDrawingMLWordprocessing, "posOffset", PrefixWP, strconv.FormatInt(offset, 10))
	pos.AppendChild(posOffset)

	return pos
}

// RelativeFrom returns the relativeFrom attribute.
func (ph *PositionH) RelativeFrom() RelativeFromH {
	attr, found := ph.GetAttribute("relativeFrom", "")
	if !found {
		return RelativeFromHColumn
	}
	return RelativeFromH(attr.Value())
}

// SetRelativeFrom sets the relativeFrom attribute.
func (ph *PositionH) SetRelativeFrom(relativeFrom RelativeFromH) {
	ph.SetAttribute(openxml.NewAttribute("", "relativeFrom", "", string(relativeFrom)))
}

// Offset returns the position offset in EMUs.
func (ph *PositionH) Offset() int64 {
	elem := ph.GetElement("posOffset", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(leaf.InnerText(), 10, 64)
		return val
	}
	return 0
}

// SetOffset sets the position offset in EMUs.
func (ph *PositionH) SetOffset(offset int64) {
	elem := ph.GetElement("posOffset", NamespaceDrawingMLWordprocessing)
	if elem != nil {
		if leaf, ok := elem.(*openxml.LeafElementBase); ok {
			leaf.SetInnerText(strconv.FormatInt(offset, 10))
		}
	}
}

// Clone creates a deep copy of this PositionH element.
func (ph *PositionH) Clone() openxml.Element {
	return &PositionH{
		CompositeElementBase: ph.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PositionV represents the wp:positionV element.
type PositionV struct {
	*openxml.CompositeElementBase
}

// NewPositionV creates a new PositionV element.
func NewPositionV(relativeFrom RelativeFromV, offset int64) *PositionV {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLWordprocessing, "positionV", PrefixWP)
	pos := &PositionV{CompositeElementBase: elem}
	pos.SetAttribute(openxml.NewAttribute("", "relativeFrom", "", string(relativeFrom)))

	// Add posOffset child
	posOffset := openxml.NewLeafElementWithText(NamespaceDrawingMLWordprocessing, "posOffset", PrefixWP, strconv.FormatInt(offset, 10))
	pos.AppendChild(posOffset)

	return pos
}

// RelativeFrom returns the relativeFrom attribute.
func (pv *PositionV) RelativeFrom() RelativeFromV {
	attr, found := pv.GetAttribute("relativeFrom", "")
	if !found {
		return RelativeFromVParagraph
	}
	return RelativeFromV(attr.Value())
}

// SetRelativeFrom sets the relativeFrom attribute.
func (pv *PositionV) SetRelativeFrom(relativeFrom RelativeFromV) {
	pv.SetAttribute(openxml.NewAttribute("", "relativeFrom", "", string(relativeFrom)))
}

// Offset returns the position offset in EMUs.
func (pv *PositionV) Offset() int64 {
	elem := pv.GetElement("posOffset", NamespaceDrawingMLWordprocessing)
	if elem == nil {
		return 0
	}
	if leaf, ok := elem.(*openxml.LeafElementBase); ok {
		val, _ := strconv.ParseInt(leaf.InnerText(), 10, 64)
		return val
	}
	return 0
}

// SetOffset sets the position offset in EMUs.
func (pv *PositionV) SetOffset(offset int64) {
	elem := pv.GetElement("posOffset", NamespaceDrawingMLWordprocessing)
	if elem != nil {
		if leaf, ok := elem.(*openxml.LeafElementBase); ok {
			leaf.SetInnerText(strconv.FormatInt(offset, 10))
		}
	}
}

// Clone creates a deep copy of this PositionV element.
func (pv *PositionV) Clone() openxml.Element {
	return &PositionV{
		CompositeElementBase: pv.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
