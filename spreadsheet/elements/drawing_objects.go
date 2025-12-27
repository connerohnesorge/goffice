package elements

//revive:disable:file-length-limit many drawing object types
//revive:disable:max-public-structs many drawing object types

import (
	"fmt"

	"github.com/connerohnesorge/goffice/openxml"
)

// DrawingPicture represents the picture element (xdr:pic).
// This is a picture/image object within a spreadsheet drawing.
// Named DrawingPicture to avoid conflict with the existing Picture type
// which represents the x:picture element for background images.
type DrawingPicture struct {
	*openxml.CompositeElementBase
}

// NewDrawingPicture creates a new DrawingPicture element.
func NewDrawingPicture() *DrawingPicture {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"pic",
		PrefixXDR,
	)

	return &DrawingPicture{
		CompositeElementBase: elem,
	}
}

// Macro returns the macro attribute value.
func (dp *DrawingPicture) Macro() string {
	attr, found := dp.GetAttribute(
		"macro",
		"",
	) //nolint:revive // add-constant: attribute name
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMacro sets the macro attribute.
func (dp *DrawingPicture) SetMacro(macro string) {
	if macro == "" {
		dp.RemoveAttribute(
			"macro",
			"",
		) //nolint:revive // add-constant: attribute name

		return
	}
	dp.SetAttribute(
		openxml.NewAttribute(
			"",
			"macro", //nolint:revive // add-constant: attribute name
			"",
			macro,
		),
	)
}

// FPublished returns whether the picture is published.
func (dp *DrawingPicture) FPublished() bool {
	attr, found := dp.GetAttribute(
		"fPublished",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetFPublished sets whether the picture is published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (dp *DrawingPicture) SetFPublished(
	value bool,
) {
	if value {
		dp.SetAttribute(
			openxml.NewAttribute(
				"",
				"fPublished",
				"",
				attrValueTrue,
			),
		)
	} else {
		dp.RemoveAttribute("fPublished", "")
	}
}

// NvPicPr returns the non-visual picture properties element,
// or nil if not present.
func (dp *DrawingPicture) NvPicPr() *NvPicPr {
	elem := dp.GetElement(
		"nvPicPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if nvPicPr, ok := elem.(*NvPicPr); ok {
		return nvPicPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NvPicPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNvPicPr returns the non-visual picture properties,
// creating if needed.
func (dp *DrawingPicture) GetOrCreateNvPicPr() *NvPicPr {
	nvPicPr := dp.NvPicPr()
	if nvPicPr != nil {
		return nvPicPr
	}
	nvPicPr = NewNvPicPr()
	if first := dp.FirstChild(); first != nil {
		dp.InsertBefore(nvPicPr, first)
	} else {
		dp.AppendChild(nvPicPr)
	}

	return nvPicPr
}

// BlipFill returns the blip fill element, or nil if not present.
func (dp *DrawingPicture) BlipFill() *BlipFill {
	elem := dp.GetElement(
		"blipFill",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if blipFill, ok := elem.(*BlipFill); ok {
		return blipFill
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BlipFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateBlipFill returns the blip fill element, creating if needed.
func (dp *DrawingPicture) GetOrCreateBlipFill() *BlipFill {
	blipFill := dp.BlipFill()
	if blipFill != nil {
		return blipFill
	}
	blipFill = NewBlipFill()
	nvPicPr := dp.NvPicPr()
	if nvPicPr != nil {
		dp.InsertAfter(blipFill, nvPicPr)
	} else {
		dp.AppendChild(blipFill)
	}

	return blipFill
}

// SpPr returns the shape properties element, or nil if not present.
func (dp *DrawingPicture) SpPr() *ShapeProperties {
	elem := dp.GetElement(
		"spPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if spPr, ok := elem.(*ShapeProperties); ok {
		return spPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSpPr returns the shape properties, creating if needed.
func (dp *DrawingPicture) GetOrCreateSpPr() *ShapeProperties {
	spPr := dp.SpPr()
	if spPr != nil {
		return spPr
	}
	spPr = NewShapeProperties()
	blipFill := dp.BlipFill()
	if blipFill != nil {
		dp.InsertAfter(spPr, blipFill)
	} else {
		dp.AppendChild(spPr)
	}

	return spPr
}

// SetPictureId is a convenience method to set shape ID and a picture name.
func (dp *DrawingPicture) SetPictureId(
	id uint32,
	name string,
) {
	nvPicPr := dp.GetOrCreateNvPicPr()
	cNvPr := nvPicPr.GetOrCreateCNvPr()
	cNvPr.SetId(id)
	cNvPr.SetName(fmt.Sprintf("Picture %d", id))
	if name != "" {
		cNvPr.SetDescr(name)
	}
}

// SetImageRelationship sets the image relationship ID.
func (dp *DrawingPicture) SetImageRelationship(
	relId string,
) {
	blipFill := dp.GetOrCreateBlipFill()
	blip := blipFill.GetOrCreateBlip()
	blip.SetEmbed(relId)
}

// Clone creates a deep copy of this DrawingPicture element.
func (dp *DrawingPicture) Clone() openxml.Element {
	cloned := dp.CompositeElementBase.Clone()

	return &DrawingPicture{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this DrawingPicture element.
func (dp *DrawingPicture) CloneNode(
	deep bool,
) openxml.Element {
	cloned := dp.CompositeElementBase.CloneNode(
		deep,
	)

	return &DrawingPicture{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NvPicPr represents non-visual picture properties (xdr:nvPicPr).
type NvPicPr struct {
	*openxml.CompositeElementBase
}

// NewNvPicPr creates a new NvPicPr element.
func NewNvPicPr() *NvPicPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"nvPicPr",
		PrefixXDR,
	)

	return &NvPicPr{CompositeElementBase: elem}
}

// CNvPr returns the common non-visual properties element,
// or nil if not present.
func (n *NvPicPr) CNvPr() *CNvPr {
	elem := n.GetElement(
		"cNvPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPr, ok := elem.(*CNvPr); ok {
		return cNvPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCNvPr returns the common non-visual properties,
// creating if needed.
func (n *NvPicPr) GetOrCreateCNvPr() *CNvPr {
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		return cNvPr
	}
	cNvPr = NewCNvPr()
	if first := n.FirstChild(); first != nil {
		n.InsertBefore(cNvPr, first)
	} else {
		n.AppendChild(cNvPr)
	}

	return cNvPr
}

// CNvPicPr returns the non-visual picture drawing properties element.
func (n *NvPicPr) CNvPicPr() *CNvPicPr {
	elem := n.GetElement(
		"cNvPicPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPicPr, ok := elem.(*CNvPicPr); ok {
		return cNvPicPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPicPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCNvPicPr returns the non-visual picture drawing properties,
// creating if needed.
func (n *NvPicPr) GetOrCreateCNvPicPr() *CNvPicPr {
	cNvPicPr := n.CNvPicPr()
	if cNvPicPr != nil {
		return cNvPicPr
	}
	cNvPicPr = NewCNvPicPr()
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		n.InsertAfter(cNvPicPr, cNvPr)
	} else {
		n.AppendChild(cNvPicPr)
	}

	return cNvPicPr
}

// Clone creates a deep copy of this NvPicPr element.
func (n *NvPicPr) Clone() openxml.Element {
	cloned := n.CompositeElementBase.Clone()

	return &NvPicPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NvPicPr element.
func (n *NvPicPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.CompositeElementBase.CloneNode(
		deep,
	)

	return &NvPicPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvPicPr represents connection non-visual picture properties (xdr:cNvPicPr).
type CNvPicPr struct {
	*openxml.CompositeElementBase
}

// NewCNvPicPr creates a new CNvPicPr element.
func NewCNvPicPr() *CNvPicPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvPicPr",
		PrefixXDR,
	)

	return &CNvPicPr{CompositeElementBase: elem}
}

// PreferRelativeResize returns whether relative resize is preferred.
func (c *CNvPicPr) PreferRelativeResize() bool {
	attr, found := c.GetAttribute(
		"preferRelativeResize",
		"",
	)
	if !found {
		return true // Default
	}

	return attr.Value() != attrValueFalse &&
		attr.Value() != attrValueZero
}

// SetPreferRelativeResize sets whether relative resize is preferred.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (c *CNvPicPr) SetPreferRelativeResize(
	value bool,
) {
	if value {
		c.RemoveAttribute(
			"preferRelativeResize",
			"",
		) // true is default

		return
	}
	c.SetAttribute(
		openxml.NewAttribute(
			"",
			"preferRelativeResize",
			"",
			attrValueFalse,
		),
	)
}

// Clone creates a deep copy of this CNvPicPr element.
func (c *CNvPicPr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvPicPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvPicPr element.
func (c *CNvPicPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvPicPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// BlipFill represents the blip fill element (xdr:blipFill).
type BlipFill struct {
	*openxml.CompositeElementBase
}

// NewBlipFill creates a new BlipFill element.
func NewBlipFill() *BlipFill {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"blipFill",
		PrefixXDR,
	)

	return &BlipFill{CompositeElementBase: elem}
}

// RotWithShape returns whether the fill rotates with shape.
func (bf *BlipFill) RotWithShape() bool {
	attr, found := bf.GetAttribute(
		"rotWithShape",
		"",
	)
	if !found {
		return true // Default
	}

	return attr.Value() != attrValueFalse &&
		attr.Value() != attrValueZero
}

// SetRotWithShape sets whether the fill rotates with shape.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (bf *BlipFill) SetRotWithShape(value bool) {
	if value {
		bf.RemoveAttribute(
			"rotWithShape",
			"",
		) // true is default

		return
	}
	bf.SetAttribute(
		openxml.NewAttribute(
			"",
			"rotWithShape",
			"",
			attrValueFalse,
		),
	)
}

// Blip returns the Blip element, or nil if not present.
func (bf *BlipFill) Blip() *Blip {
	elem := bf.GetElement(
		"blip",
		NamespaceDrawingML,
	)
	if elem == nil {
		return nil
	}
	if blip, ok := elem.(*Blip); ok {
		return blip
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Blip{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateBlip returns the Blip element, creating if needed.
func (bf *BlipFill) GetOrCreateBlip() *Blip {
	blip := bf.Blip()
	if blip != nil {
		return blip
	}
	blip = NewBlip()
	if first := bf.FirstChild(); first != nil {
		bf.InsertBefore(blip, first)
	} else {
		bf.AppendChild(blip)
	}

	return blip
}

// Clone creates a deep copy of this BlipFill element.
func (bf *BlipFill) Clone() openxml.Element {
	cloned := bf.CompositeElementBase.Clone()

	return &BlipFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this BlipFill element.
func (bf *BlipFill) CloneNode(
	deep bool,
) openxml.Element {
	cloned := bf.CompositeElementBase.CloneNode(
		deep,
	)

	return &BlipFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Blip represents the blip element (a:blip).
// This references an image via relationship ID.
type Blip struct {
	*openxml.CompositeElementBase
}

// NewBlip creates a new Blip element.
func NewBlip() *Blip {
	elem := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"blip",
		PrefixA,
	)

	return &Blip{CompositeElementBase: elem}
}

// Embed returns the embedded relationship ID.
func (b *Blip) Embed() string {
	attr, found := b.GetAttribute(
		"embed",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetEmbed sets the embedded relationship ID.
func (b *Blip) SetEmbed(id string) {
	if id == "" {
		b.RemoveAttribute(
			"embed",
			NamespaceRelationships,
		)

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"embed",
			PrefixR,
			id,
		),
	)
}

// Link returns the linked relationship ID.
func (b *Blip) Link() string {
	attr, found := b.GetAttribute(
		"link",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetLink sets the linked relationship ID.
func (b *Blip) SetLink(id string) {
	if id == "" {
		b.RemoveAttribute(
			"link",
			NamespaceRelationships,
		)

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			NamespaceRelationships,
			"link",
			PrefixR,
			id,
		),
	)
}

// Cstate returns the compression state.
func (b *Blip) Cstate() string {
	attr, found := b.GetAttribute("cstate", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetCstate sets the compression state.
func (b *Blip) SetCstate(cstate string) {
	if cstate == "" {
		b.RemoveAttribute("cstate", "")

		return
	}
	b.SetAttribute(
		openxml.NewAttribute(
			"",
			"cstate",
			"",
			cstate,
		),
	)
}

// Clone creates a deep copy of this Blip element.
func (b *Blip) Clone() openxml.Element {
	cloned := b.CompositeElementBase.Clone()

	return &Blip{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Blip element.
func (b *Blip) CloneNode(
	deep bool,
) openxml.Element {
	cloned := b.CompositeElementBase.CloneNode(
		deep,
	)

	return &Blip{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// GraphicFrame represents the graphic frame element (xdr:graphicFrame).
// This is used to contain charts and other graphic objects.
type GraphicFrame struct {
	*openxml.CompositeElementBase
}

// NewGraphicFrame creates a new GraphicFrame element.
func NewGraphicFrame() *GraphicFrame {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"graphicFrame",
		PrefixXDR,
	)

	return &GraphicFrame{
		CompositeElementBase: elem,
	}
}

// Macro returns the macro attribute value.
func (gf *GraphicFrame) Macro() string {
	attr, found := gf.GetAttribute(
		"macro", //nolint:revive // add-constant
		"",
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMacro sets the macro attribute.
func (gf *GraphicFrame) SetMacro(macro string) {
	if macro == "" {
		gf.RemoveAttribute(
			"macro",
			"",
		) //nolint:revive // add-constant: attribute name

		return
	}
	gf.SetAttribute(
		openxml.NewAttribute(
			"",
			"macro", //nolint:revive // add-constant: attribute name
			"",
			macro,
		),
	)
}

// FPublished returns whether the graphic frame is published.
//
//nolint:revive // add-constant: attribute name standard pattern
func (gf *GraphicFrame) FPublished() bool {
	attr, found := gf.GetAttribute(
		"fPublished",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetFPublished sets whether the graphic frame is published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (gf *GraphicFrame) SetFPublished(
	value bool,
) {
	if value {
		gf.SetAttribute(
			openxml.NewAttribute(
				"",
				"fPublished",
				"",
				attrValueTrue,
			),
		)
	} else {
		gf.RemoveAttribute("fPublished", "")
	}
}

// NvGraphicFramePr returns the non-visual graphic frame properties, or nil if
// not present.
func (gf *GraphicFrame) NvGraphicFramePr() *NvGraphicFramePr {
	elem := gf.GetElement(
		"nvGraphicFramePr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if nvGfPr, ok := elem.(*NvGraphicFramePr); ok {
		return nvGfPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NvGraphicFramePr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNvGraphicFramePr returns the non-visual graphic frame properties,
// creating if needed.
func (gf *GraphicFrame) GetOrCreateNvGraphicFramePr() *NvGraphicFramePr {
	nvGfPr := gf.NvGraphicFramePr()
	if nvGfPr != nil {
		return nvGfPr
	}
	nvGfPr = NewNvGraphicFramePr()
	if first := gf.FirstChild(); first != nil {
		gf.InsertBefore(nvGfPr, first)
	} else {
		gf.AppendChild(nvGfPr)
	}

	return nvGfPr
}

// Xfrm returns the transform element, or nil if not present.
func (gf *GraphicFrame) Xfrm() *Transform {
	elem := gf.GetElement(
		"xfrm",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if xfrm, ok := elem.(*Transform); ok {
		return xfrm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Transform{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateXfrm returns the transform element, creating if needed.
func (gf *GraphicFrame) GetOrCreateXfrm() *Transform {
	xfrm := gf.Xfrm()
	if xfrm != nil {
		return xfrm
	}
	xfrm = NewTransform()
	nvGfPr := gf.NvGraphicFramePr()
	if nvGfPr != nil {
		gf.InsertAfter(xfrm, nvGfPr)
	} else {
		gf.AppendChild(xfrm)
	}

	return xfrm
}

// Graphic returns the graphic element, or nil if not present.
func (gf *GraphicFrame) Graphic() *Graphic {
	elem := gf.GetElement(
		"graphic",
		NamespaceDrawingML,
	)
	if elem == nil {
		return nil
	}
	if g, ok := elem.(*Graphic); ok {
		return g
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Graphic{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateGraphic returns the graphic element, creating if needed.
func (gf *GraphicFrame) GetOrCreateGraphic() *Graphic {
	g := gf.Graphic()
	if g != nil {
		return g
	}
	g = NewGraphic()
	gf.AppendChild(g)

	return g
}

// SetChartInfo is a convenience method to set up a chart graphic frame.
func (gf *GraphicFrame) SetChartInfo(
	id uint32,
	name string,
) {
	nvGfPr := gf.GetOrCreateNvGraphicFramePr()
	cNvPr := nvGfPr.GetOrCreateCNvPr()
	cNvPr.SetId(id)
	cNvPr.SetName(fmt.Sprintf("Chart %d", id))
	if name != "" {
		cNvPr.SetDescr(name)
	}
}

// Clone creates a deep copy of this GraphicFrame element.
func (gf *GraphicFrame) Clone() openxml.Element {
	cloned := gf.CompositeElementBase.Clone()

	return &GraphicFrame{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GraphicFrame element.
func (gf *GraphicFrame) CloneNode(
	deep bool,
) openxml.Element {
	cloned := gf.CompositeElementBase.CloneNode(
		deep,
	)

	return &GraphicFrame{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NvGraphicFramePr represents non-visual graphic frame properties
// (xdr:nvGraphicFramePr).
type NvGraphicFramePr struct {
	*openxml.CompositeElementBase
}

// NewNvGraphicFramePr creates a new NvGraphicFramePr element.
func NewNvGraphicFramePr() *NvGraphicFramePr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"nvGraphicFramePr",
		PrefixXDR,
	)

	return &NvGraphicFramePr{
		CompositeElementBase: elem,
	}
}

// CNvPr returns the common non-visual properties element,
// or nil if not present.
func (n *NvGraphicFramePr) CNvPr() *CNvPr {
	elem := n.GetElement(
		"cNvPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPr, ok := elem.(*CNvPr); ok {
		return cNvPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCNvPr returns the common non-visual properties,
// creating if needed.
func (n *NvGraphicFramePr) GetOrCreateCNvPr() *CNvPr {
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		return cNvPr
	}
	cNvPr = NewCNvPr()
	if first := n.FirstChild(); first != nil {
		n.InsertBefore(cNvPr, first)
	} else {
		n.AppendChild(cNvPr)
	}

	return cNvPr
}

// CNvGraphicFramePr returns the non-visual graphic frame drawing properties.
func (n *NvGraphicFramePr) CNvGraphicFramePr() *CNvGraphicFramePr {
	elem := n.GetElement(
		"cNvGraphicFramePr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvGfPr, ok := elem.(*CNvGraphicFramePr); ok {
		return cNvGfPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvGraphicFramePr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCNvGraphicFramePr returns the non-visual graphic frame
// properties, creating if needed.
func (n *NvGraphicFramePr) GetOrCreateCNvGraphicFramePr() *CNvGraphicFramePr {
	cNvGfPr := n.CNvGraphicFramePr()
	if cNvGfPr != nil {
		return cNvGfPr
	}
	cNvGfPr = NewCNvGraphicFramePr()
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		n.InsertAfter(cNvGfPr, cNvPr)
	} else {
		n.AppendChild(cNvGfPr)
	}

	return cNvGfPr
}

// Clone creates a deep copy of this NvGraphicFramePr element.
func (n *NvGraphicFramePr) Clone() openxml.Element {
	cloned := n.CompositeElementBase.Clone()

	return &NvGraphicFramePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NvGraphicFramePr element.
func (n *NvGraphicFramePr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.CompositeElementBase.CloneNode(
		deep,
	)

	return &NvGraphicFramePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvGraphicFramePr represents non-visual graphic frame drawing properties
// (xdr:cNvGraphicFramePr).
type CNvGraphicFramePr struct {
	*openxml.CompositeElementBase
}

// NewCNvGraphicFramePr creates a new CNvGraphicFramePr element.
func NewCNvGraphicFramePr() *CNvGraphicFramePr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvGraphicFramePr",
		PrefixXDR,
	)

	return &CNvGraphicFramePr{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this CNvGraphicFramePr element.
func (c *CNvGraphicFramePr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvGraphicFramePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvGraphicFramePr element.
func (c *CNvGraphicFramePr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvGraphicFramePr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Transform represents the transform element (xdr:xfrm).
type Transform struct {
	*openxml.CompositeElementBase
}

// NewTransform creates a new Transform element.
func NewTransform() *Transform {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"xfrm",
		PrefixXDR,
	)

	return &Transform{CompositeElementBase: elem}
}

// Clone creates a deep copy of this Transform element.
func (t *Transform) Clone() openxml.Element {
	cloned := t.CompositeElementBase.Clone()

	return &Transform{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Transform element.
func (t *Transform) CloneNode(
	deep bool,
) openxml.Element {
	cloned := t.CompositeElementBase.CloneNode(
		deep,
	)

	return &Transform{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Graphic represents the graphic element (a:graphic).
type Graphic struct {
	*openxml.CompositeElementBase
}

// NewGraphic creates a new Graphic element.
func NewGraphic() *Graphic {
	elem := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphic",
		PrefixA,
	)

	return &Graphic{CompositeElementBase: elem}
}

// GraphicData returns the graphic data element, or nil if not present.
func (g *Graphic) GraphicData() *GraphicData {
	elem := g.GetElement(
		"graphicData",
		NamespaceDrawingML,
	)
	if elem == nil {
		return nil
	}
	if gd, ok := elem.(*GraphicData); ok {
		return gd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GraphicData{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateGraphicData returns the graphic data element, creating if needed.
func (g *Graphic) GetOrCreateGraphicData() *GraphicData {
	gd := g.GraphicData()
	if gd != nil {
		return gd
	}
	gd = NewGraphicData()
	g.AppendChild(gd)

	return gd
}

// Clone creates a deep copy of this Graphic element.
func (g *Graphic) Clone() openxml.Element {
	cloned := g.CompositeElementBase.Clone()

	return &Graphic{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Graphic element.
func (g *Graphic) CloneNode(
	deep bool,
) openxml.Element {
	cloned := g.CompositeElementBase.CloneNode(
		deep,
	)

	return &Graphic{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// GraphicData represents the graphic data element (a:graphicData).
type GraphicData struct {
	*openxml.CompositeElementBase
}

// NewGraphicData creates a new GraphicData element.
func NewGraphicData() *GraphicData {
	elem := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"graphicData",
		PrefixA,
	)

	return &GraphicData{
		CompositeElementBase: elem,
	}
}

// URI returns the URI identifying the type of graphic data.
func (gd *GraphicData) URI() string {
	attr, found := gd.GetAttribute("uri", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetURI sets the URI identifying the type of graphic data.
func (gd *GraphicData) SetURI(uri string) {
	gd.SetAttribute(
		openxml.NewAttribute("", "uri", "", uri),
	)
}

// Chart namespace URI constant for graphic data.
const NamespaceChart = "http://schemas.openxmlformats.org/drawingml/2006/chart"

// SetChartURI sets the URI to the chart namespace.
func (gd *GraphicData) SetChartURI() {
	gd.SetURI(NamespaceChart)
}

// Clone creates a deep copy of this GraphicData element.
func (gd *GraphicData) Clone() openxml.Element {
	cloned := gd.CompositeElementBase.Clone()

	return &GraphicData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GraphicData element.
func (gd *GraphicData) CloneNode(
	deep bool,
) openxml.Element {
	cloned := gd.CompositeElementBase.CloneNode(
		deep,
	)

	return &GraphicData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ClientData represents the client data element (xdr:clientData).
// This contains client-specific data for the drawing object.
type ClientData struct {
	*openxml.CompositeElementBase
}

// NewClientData creates a new ClientData element.
func NewClientData() *ClientData {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"clientData",
		PrefixXDR,
	)

	return &ClientData{CompositeElementBase: elem}
}

// FLocksWithSheet returns whether the object is locked when the sheet is
// protected.
func (cd *ClientData) FLocksWithSheet() bool {
	attr, found := cd.GetAttribute(
		"fLocksWithSheet",
		"",
	)
	if !found {
		return true // Default is true per spec
	}

	return attr.Value() != attrValueFalse &&
		attr.Value() != attrValueZero
}

// SetFLocksWithSheet sets whether the object is locked when the sheet is protected.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cd *ClientData) SetFLocksWithSheet(
	value bool,
) {
	if value {
		// Default value, can be omitted
		cd.RemoveAttribute("fLocksWithSheet", "")

		return
	}
	cd.SetAttribute(
		openxml.NewAttribute(
			"",
			"fLocksWithSheet",
			"",
			attrValueFalse,
		),
	)
}

// FPrintsWithSheet returns whether the object prints when the sheet is printed.
func (cd *ClientData) FPrintsWithSheet() bool {
	attr, found := cd.GetAttribute(
		"fPrintsWithSheet",
		"",
	)
	if !found {
		return true // Default is true per spec
	}

	return attr.Value() != attrValueFalse &&
		attr.Value() != attrValueZero
}

// SetFPrintsWithSheet sets whether the object prints when the sheet is printed.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cd *ClientData) SetFPrintsWithSheet(
	value bool,
) {
	if value {
		// Default value, can be omitted
		cd.RemoveAttribute("fPrintsWithSheet", "")

		return
	}
	cd.SetAttribute(
		openxml.NewAttribute(
			"",
			"fPrintsWithSheet",
			"",
			attrValueFalse,
		),
	)
}

// Clone creates a deep copy of this ClientData element.
func (cd *ClientData) Clone() openxml.Element {
	cloned := cd.CompositeElementBase.Clone()

	return &ClientData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ClientData element.
func (cd *ClientData) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cd.CompositeElementBase.CloneNode(
		deep,
	)

	return &ClientData{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ConnectionShape represents the connection shape element (xdr:cxnSp).
// This is used for connector lines between shapes.
type ConnectionShape struct {
	*openxml.CompositeElementBase
}

// NewConnectionShape creates a new ConnectionShape element.
func NewConnectionShape() *ConnectionShape {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cxnSp",
		PrefixXDR,
	)

	return &ConnectionShape{
		CompositeElementBase: elem,
	}
}

// Macro returns the macro attribute value.
func (cs *ConnectionShape) Macro() string {
	attr, found := cs.GetAttribute(
		"macro",
		"",
	) //nolint:revive // add-constant: attribute name
	if !found {
		return ""
	}

	return attr.Value()
}

// SetMacro sets the macro attribute.
func (cs *ConnectionShape) SetMacro(
	macro string,
) {
	if macro == "" {
		cs.RemoveAttribute(
			"macro",
			"",
		) //nolint:revive // add-constant: attribute name

		return
	}
	cs.SetAttribute(
		openxml.NewAttribute(
			"",
			"macro", //nolint:revive // add-constant: attribute name
			"",
			macro,
		),
	)
}

// FPublished returns whether the connection shape is published.
func (cs *ConnectionShape) FPublished() bool {
	attr, found := cs.GetAttribute(
		"fPublished",
		"",
	)
	if !found {
		return false
	}

	return attr.Value() == attrValueTrue ||
		attr.Value() == attrValueOne
}

// SetFPublished sets whether the connection shape is published.
//
//nolint:revive // flag-parameter: bool setter is the standard pattern
func (cs *ConnectionShape) SetFPublished(
	value bool,
) {
	if value {
		cs.SetAttribute(
			openxml.NewAttribute(
				"",
				"fPublished",
				"",
				attrValueTrue,
			),
		)
	} else {
		cs.RemoveAttribute("fPublished", "")
	}
}

// NvCxnSpPr returns the non-visual connection shape properties element,
// or nil if not present.
func (cs *ConnectionShape) NvCxnSpPr() *NvCxnSpPr {
	elem := cs.GetElement(
		"nvCxnSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if nvCxnSpPr, ok := elem.(*NvCxnSpPr); ok {
		return nvCxnSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NvCxnSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNvCxnSpPr returns the non-visual connection shape properties,
// creating if needed.
func (cs *ConnectionShape) GetOrCreateNvCxnSpPr() *NvCxnSpPr {
	nvCxnSpPr := cs.NvCxnSpPr()
	if nvCxnSpPr != nil {
		return nvCxnSpPr
	}
	nvCxnSpPr = NewNvCxnSpPr()
	if first := cs.FirstChild(); first != nil {
		cs.InsertBefore(nvCxnSpPr, first)
	} else {
		cs.AppendChild(nvCxnSpPr)
	}

	return nvCxnSpPr
}

// SpPr returns the shape properties element, or nil if not present.
func (cs *ConnectionShape) SpPr() *ShapeProperties {
	elem := cs.GetElement(
		"spPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if spPr, ok := elem.(*ShapeProperties); ok {
		return spPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateSpPr returns the shape properties, creating if needed.
func (cs *ConnectionShape) GetOrCreateSpPr() *ShapeProperties {
	spPr := cs.SpPr()
	if spPr != nil {
		return spPr
	}
	spPr = NewShapeProperties()
	nvCxnSpPr := cs.NvCxnSpPr()
	if nvCxnSpPr != nil {
		cs.InsertAfter(spPr, nvCxnSpPr)
	} else {
		cs.AppendChild(spPr)
	}

	return spPr
}

// Clone creates a deep copy of this ConnectionShape element.
func (cs *ConnectionShape) Clone() openxml.Element {
	cloned := cs.CompositeElementBase.Clone()

	return &ConnectionShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this ConnectionShape element.
func (cs *ConnectionShape) CloneNode(
	deep bool,
) openxml.Element {
	cloned := cs.CompositeElementBase.CloneNode(
		deep,
	)

	return &ConnectionShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NvCxnSpPr represents non-visual connection shape properties (xdr:nvCxnSpPr).
type NvCxnSpPr struct {
	*openxml.CompositeElementBase
}

// NewNvCxnSpPr creates a new NvCxnSpPr element.
func NewNvCxnSpPr() *NvCxnSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"nvCxnSpPr",
		PrefixXDR,
	)

	return &NvCxnSpPr{CompositeElementBase: elem}
}

// CNvPr returns the common non-visual properties element,
// or nil if not present.
func (n *NvCxnSpPr) CNvPr() *CNvPr {
	elem := n.GetElement(
		"cNvPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPr, ok := elem.(*CNvPr); ok {
		return cNvPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCNvPr returns the common non-visual properties, creating if
// needed.
func (n *NvCxnSpPr) GetOrCreateCNvPr() *CNvPr {
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		return cNvPr
	}
	cNvPr = NewCNvPr()
	if first := n.FirstChild(); first != nil {
		n.InsertBefore(cNvPr, first)
	} else {
		n.AppendChild(cNvPr)
	}

	return cNvPr
}

// CNvCxnSpPr returns the non-visual connection shape drawing properties.
func (n *NvCxnSpPr) CNvCxnSpPr() *CNvCxnSpPr {
	elem := n.GetElement(
		"cNvCxnSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvCxnSpPr, ok := elem.(*CNvCxnSpPr); ok {
		return cNvCxnSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvCxnSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCNvCxnSpPr returns the non-visual connection shape drawing
// properties, creating if needed.
func (n *NvCxnSpPr) GetOrCreateCNvCxnSpPr() *CNvCxnSpPr {
	cNvCxnSpPr := n.CNvCxnSpPr()
	if cNvCxnSpPr != nil {
		return cNvCxnSpPr
	}
	cNvCxnSpPr = NewCNvCxnSpPr()
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		n.InsertAfter(cNvCxnSpPr, cNvPr)
	} else {
		n.AppendChild(cNvCxnSpPr)
	}

	return cNvCxnSpPr
}

// Clone creates a deep copy of this NvCxnSpPr element.
func (n *NvCxnSpPr) Clone() openxml.Element {
	cloned := n.CompositeElementBase.Clone()

	return &NvCxnSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NvCxnSpPr element.
func (n *NvCxnSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.CompositeElementBase.CloneNode(
		deep,
	)

	return &NvCxnSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvCxnSpPr represents connection non-visual shape properties (cNvCxnSpPr).
type CNvCxnSpPr struct {
	*openxml.CompositeElementBase
}

// NewCNvCxnSpPr creates a new CNvCxnSpPr element.
func NewCNvCxnSpPr() *CNvCxnSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvCxnSpPr",
		PrefixXDR,
	)

	return &CNvCxnSpPr{CompositeElementBase: elem}
}

// Clone creates a deep copy of this CNvCxnSpPr element.
func (c *CNvCxnSpPr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvCxnSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvCxnSpPr element.
func (c *CNvCxnSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvCxnSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// GroupShape represents the group shape element (xdr:grpSp).
// This is used to group multiple drawing objects together.
type GroupShape struct {
	*openxml.CompositeElementBase
}

// NewGroupShape creates a new GroupShape element.
func NewGroupShape() *GroupShape {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"grpSp",
		PrefixXDR,
	)

	return &GroupShape{
		CompositeElementBase: elem,
	}
}

// NvGrpSpPr returns the non-visual group shape properties element,
// or nil if not present.
func (gs *GroupShape) NvGrpSpPr() *NvGrpSpPr {
	elem := gs.GetElement(
		"nvGrpSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if nvGrpSpPr, ok := elem.(*NvGrpSpPr); ok {
		return nvGrpSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NvGrpSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateNvGrpSpPr returns the non-visual group shape properties,
// creating if needed.
func (gs *GroupShape) GetOrCreateNvGrpSpPr() *NvGrpSpPr {
	nvGrpSpPr := gs.NvGrpSpPr()
	if nvGrpSpPr != nil {
		return nvGrpSpPr
	}
	nvGrpSpPr = NewNvGrpSpPr()
	if first := gs.FirstChild(); first != nil {
		gs.InsertBefore(nvGrpSpPr, first)
	} else {
		gs.AppendChild(nvGrpSpPr)
	}

	return nvGrpSpPr
}

// GrpSpPr returns the group shape properties element, or nil if not present.
func (gs *GroupShape) GrpSpPr() *GrpSpPr {
	elem := gs.GetElement(
		"grpSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if grpSpPr, ok := elem.(*GrpSpPr); ok {
		return grpSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GrpSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateGrpSpPr returns the group shape properties, creating if needed.
func (gs *GroupShape) GetOrCreateGrpSpPr() *GrpSpPr {
	grpSpPr := gs.GrpSpPr()
	if grpSpPr != nil {
		return grpSpPr
	}
	grpSpPr = NewGrpSpPr()
	nvGrpSpPr := gs.NvGrpSpPr()
	if nvGrpSpPr != nil {
		gs.InsertAfter(grpSpPr, nvGrpSpPr)
	} else {
		gs.AppendChild(grpSpPr)
	}

	return grpSpPr
}

// AddShape adds a new Shape element to the group.
func (gs *GroupShape) AddShape() *Shape {
	sp := NewShape()
	gs.AppendChild(sp)

	return sp
}

// AddPicture adds a new Picture element to the group.
func (gs *GroupShape) AddPicture() *DrawingPicture {
	pic := NewDrawingPicture()
	gs.AppendChild(pic)

	return pic
}

// AddConnectionShape adds a new ConnectionShape element to the group.
func (gs *GroupShape) AddConnectionShape() *ConnectionShape {
	cxnSp := NewConnectionShape()
	gs.AppendChild(cxnSp)

	return cxnSp
}

// AddGraphicFrame adds a new GraphicFrame element to the group.
func (gs *GroupShape) AddGraphicFrame() *GraphicFrame {
	gf := NewGraphicFrame()
	gs.AppendChild(gf)

	return gf
}

// AddGroupShape adds a nested GroupShape element.
func (gs *GroupShape) AddGroupShape() *GroupShape {
	grpSp := NewGroupShape()
	gs.AppendChild(grpSp)

	return grpSp
}

// Clone creates a deep copy of this GroupShape element.
func (gs *GroupShape) Clone() openxml.Element {
	cloned := gs.CompositeElementBase.Clone()

	return &GroupShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GroupShape element.
func (gs *GroupShape) CloneNode(
	deep bool,
) openxml.Element {
	cloned := gs.CompositeElementBase.CloneNode(
		deep,
	)

	return &GroupShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// NvGrpSpPr represents non-visual group shape properties (xdr:nvGrpSpPr).
type NvGrpSpPr struct {
	*openxml.CompositeElementBase
}

// NewNvGrpSpPr creates a new NvGrpSpPr element.
func NewNvGrpSpPr() *NvGrpSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"nvGrpSpPr",
		PrefixXDR,
	)

	return &NvGrpSpPr{CompositeElementBase: elem}
}

// CNvPr returns the common non-visual properties element,
// or nil if not present.
//
//nolint:revive // add-constant: element name standard pattern
func (n *NvGrpSpPr) CNvPr() *CNvPr {
	elem := n.GetElement(
		"cNvPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvPr, ok := elem.(*CNvPr); ok {
		return cNvPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvPr{CompositeElementBase: comp}
	}

	return nil
}

// GetOrCreateCNvPr returns the common non-visual properties, creating if
// needed.
func (n *NvGrpSpPr) GetOrCreateCNvPr() *CNvPr {
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		return cNvPr
	}
	cNvPr = NewCNvPr()
	if first := n.FirstChild(); first != nil {
		n.InsertBefore(cNvPr, first)
	} else {
		n.AppendChild(cNvPr)
	}

	return cNvPr
}

// CNvGrpSpPr returns the non-visual group shape drawing properties element.
func (n *NvGrpSpPr) CNvGrpSpPr() *CNvGrpSpPr {
	elem := n.GetElement(
		"cNvGrpSpPr",
		NamespaceSpreadsheetDrawing,
	)
	if elem == nil {
		return nil
	}
	if cNvGrpSpPr, ok := elem.(*CNvGrpSpPr); ok {
		return cNvGrpSpPr
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &CNvGrpSpPr{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GetOrCreateCNvGrpSpPr returns the non-visual group shape drawing
// properties, creating if needed.
func (n *NvGrpSpPr) GetOrCreateCNvGrpSpPr() *CNvGrpSpPr {
	cNvGrpSpPr := n.CNvGrpSpPr()
	if cNvGrpSpPr != nil {
		return cNvGrpSpPr
	}
	cNvGrpSpPr = NewCNvGrpSpPr()
	cNvPr := n.CNvPr()
	if cNvPr != nil {
		n.InsertAfter(cNvGrpSpPr, cNvPr)
	} else {
		n.AppendChild(cNvGrpSpPr)
	}

	return cNvGrpSpPr
}

// Clone creates a deep copy of this NvGrpSpPr element.
func (n *NvGrpSpPr) Clone() openxml.Element {
	cloned := n.CompositeElementBase.Clone()

	return &NvGrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this NvGrpSpPr element.
func (n *NvGrpSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := n.CompositeElementBase.CloneNode(
		deep,
	)

	return &NvGrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CNvGrpSpPr represents connection non-visual group shape properties
// (xdr:cNvGrpSpPr).
type CNvGrpSpPr struct {
	*openxml.CompositeElementBase
}

// NewCNvGrpSpPr creates a new CNvGrpSpPr element.
func NewCNvGrpSpPr() *CNvGrpSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"cNvGrpSpPr",
		PrefixXDR,
	)

	return &CNvGrpSpPr{CompositeElementBase: elem}
}

// Clone creates a deep copy of this CNvGrpSpPr element.
func (c *CNvGrpSpPr) Clone() openxml.Element {
	cloned := c.CompositeElementBase.Clone()

	return &CNvGrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this CNvGrpSpPr element.
func (c *CNvGrpSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := c.CompositeElementBase.CloneNode(
		deep,
	)

	return &CNvGrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// GrpSpPr represents group shape properties (xdr:grpSpPr).
type GrpSpPr struct {
	*openxml.CompositeElementBase
}

// NewGrpSpPr creates a new GrpSpPr element.
func NewGrpSpPr() *GrpSpPr {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"grpSpPr",
		PrefixXDR,
	)

	return &GrpSpPr{CompositeElementBase: elem}
}

// BWMode returns the black and white mode.
func (g *GrpSpPr) BWMode() string {
	attr, found := g.GetAttribute("bwMode", "")
	if !found {
		return ""
	}

	return attr.Value()
}

// SetBWMode sets the black and white mode.
func (g *GrpSpPr) SetBWMode(
	mode string,
) {
	if mode == "" {
		g.RemoveAttribute("bwMode", "")

		return
	}
	g.SetAttribute(
		openxml.NewAttribute(
			"",
			"bwMode",
			"",
			mode,
		),
	)
}

// Clone creates a deep copy of this GrpSpPr element.
func (g *GrpSpPr) Clone() openxml.Element {
	cloned := g.CompositeElementBase.Clone()

	return &GrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this GrpSpPr element.
func (g *GrpSpPr) CloneNode(
	deep bool,
) openxml.Element {
	cloned := g.CompositeElementBase.CloneNode(
		deep,
	)

	return &GrpSpPr{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
