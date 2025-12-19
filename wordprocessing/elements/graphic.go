package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Graphic represents the a:graphic element.
type Graphic struct {
	*openxml.CompositeElementBase
}

// NewGraphic creates a new empty Graphic element.
func NewGraphic() *Graphic {
	elem := openxml.NewCompositeElement(NamespaceDrawingML, "graphic", PrefixA)
	return &Graphic{CompositeElementBase: elem}
}

// NewGraphicWithPicture creates a new Graphic element containing a picture.
func NewGraphicWithPicture(relId string, width, height int64) *Graphic {
	g := NewGraphic()

	// Add graphic data
	graphicData := NewGraphicData(NamespaceDrawingMLPicture)
	g.AppendChild(graphicData)

	// Add picture
	pic := NewPicture(relId, width, height)
	graphicData.AppendChild(pic)

	return g
}

// GraphicData returns the graphic data element.
func (g *Graphic) GraphicData() *GraphicData {
	elem := g.GetElement("graphicData", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if gd, ok := elem.(*GraphicData); ok {
		return gd
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &GraphicData{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this Graphic element.
func (g *Graphic) Clone() openxml.Element {
	return &Graphic{
		CompositeElementBase: g.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// GraphicData represents the a:graphicData element.
type GraphicData struct {
	*openxml.CompositeElementBase
}

// NewGraphicData creates a new GraphicData element.
func NewGraphicData(uri string) *GraphicData {
	elem := openxml.NewCompositeElement(NamespaceDrawingML, "graphicData", PrefixA)
	gd := &GraphicData{CompositeElementBase: elem}
	gd.SetAttribute(openxml.NewAttribute("", "uri", "", uri))
	return gd
}

// URI returns the uri attribute.
func (gd *GraphicData) URI() string {
	attr, found := gd.GetAttribute("uri", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// Picture returns the picture element, or nil if not present.
func (gd *GraphicData) Picture() *Picture {
	elem := gd.GetElement("pic", NamespaceDrawingMLPicture)
	if elem == nil {
		return nil
	}
	if p, ok := elem.(*Picture); ok {
		return p
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Picture{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this GraphicData element.
func (gd *GraphicData) Clone() openxml.Element {
	return &GraphicData{
		CompositeElementBase: gd.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Picture represents the pic:pic element for images.
type Picture struct {
	*openxml.CompositeElementBase
}

// NewPicture creates a new Picture element.
func NewPicture(relId string, width, height int64) *Picture {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "pic", PrefixPic)
	pic := &Picture{CompositeElementBase: elem}

	// Add non-visual picture properties
	nvPicPr := NewNonVisualPictureProperties(0, "Picture")
	pic.AppendChild(nvPicPr)

	// Add blip fill
	blipFill := NewBlipFill(relId)
	pic.AppendChild(blipFill)

	// Add shape properties
	spPr := NewShapeProperties(width, height)
	pic.AppendChild(spPr)

	return pic
}

// NonVisualPictureProperties returns the non-visual picture properties.
func (p *Picture) NonVisualPictureProperties() *NonVisualPictureProperties {
	elem := p.GetElement("nvPicPr", NamespaceDrawingMLPicture)
	if elem == nil {
		return nil
	}
	if nvp, ok := elem.(*NonVisualPictureProperties); ok {
		return nvp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &NonVisualPictureProperties{CompositeElementBase: comp}
	}
	return nil
}

// BlipFill returns the blip fill element.
func (p *Picture) BlipFill() *BlipFill {
	elem := p.GetElement("blipFill", NamespaceDrawingMLPicture)
	if elem == nil {
		return nil
	}
	if bf, ok := elem.(*BlipFill); ok {
		return bf
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &BlipFill{CompositeElementBase: comp}
	}
	return nil
}

// ShapeProperties returns the shape properties element.
func (p *Picture) ShapeProperties() *ShapeProperties {
	elem := p.GetElement("spPr", NamespaceDrawingMLPicture)
	if elem == nil {
		return nil
	}
	if sp, ok := elem.(*ShapeProperties); ok {
		return sp
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &ShapeProperties{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this Picture element.
func (p *Picture) Clone() openxml.Element {
	return &Picture{
		CompositeElementBase: p.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// NonVisualPictureProperties represents the pic:nvPicPr element.
type NonVisualPictureProperties struct {
	*openxml.CompositeElementBase
}

// NewNonVisualPictureProperties creates a new NonVisualPictureProperties element.
func NewNonVisualPictureProperties(id int, name string) *NonVisualPictureProperties {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "nvPicPr", PrefixPic)
	nvp := &NonVisualPictureProperties{CompositeElementBase: elem}

	// Add cNvPr
	cNvPr := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "cNvPr", PrefixPic)
	cNvPr.SetAttribute(openxml.NewAttribute("", "id", "", strconv.Itoa(id)))
	cNvPr.SetAttribute(openxml.NewAttribute("", "name", "", name))
	nvp.AppendChild(cNvPr)

	// Add cNvPicPr
	cNvPicPr := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "cNvPicPr", PrefixPic)
	nvp.AppendChild(cNvPicPr)

	return nvp
}

// Clone creates a deep copy of this NonVisualPictureProperties element.
func (nvp *NonVisualPictureProperties) Clone() openxml.Element {
	return &NonVisualPictureProperties{
		CompositeElementBase: nvp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// BlipFill represents the pic:blipFill element.
type BlipFill struct {
	*openxml.CompositeElementBase
}

// NewBlipFill creates a new BlipFill element.
func NewBlipFill(relId string) *BlipFill {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "blipFill", PrefixPic)
	bf := &BlipFill{CompositeElementBase: elem}

	// Add blip
	blip := NewBlip(relId)
	bf.AppendChild(blip)

	// Add stretch with fillRect
	stretch := openxml.NewCompositeElement(NamespaceDrawingML, "stretch", PrefixA)
	fillRect := openxml.NewCompositeElement(NamespaceDrawingML, "fillRect", PrefixA)
	stretch.AppendChild(fillRect)
	bf.AppendChild(stretch)

	return bf
}

// Blip returns the blip element.
func (bf *BlipFill) Blip() *Blip {
	elem := bf.GetElement("blip", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if b, ok := elem.(*Blip); ok {
		return b
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Blip{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this BlipFill element.
func (bf *BlipFill) Clone() openxml.Element {
	return &BlipFill{
		CompositeElementBase: bf.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Blip represents the a:blip element for image references.
type Blip struct {
	*openxml.CompositeElementBase
}

// NewBlip creates a new Blip element.
func NewBlip(relId string) *Blip {
	elem := openxml.NewCompositeElement(NamespaceDrawingML, "blip", PrefixA)
	blip := &Blip{CompositeElementBase: elem}
	blip.SetEmbed(relId)
	return blip
}

// Embed returns the relationship ID to the image part.
func (b *Blip) Embed() string {
	attr, found := b.GetAttribute("embed", NamespaceRelationships)
	if !found {
		return ""
	}
	return attr.Value()
}

// SetEmbed sets the relationship ID to the image part.
func (b *Blip) SetEmbed(relId string) {
	b.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "embed", PrefixR, relId))
}

// CompressionState returns the compression state attribute.
func (b *Blip) CompressionState() string {
	attr, found := b.GetAttribute("cstate", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// SetCompressionState sets the compression state.
// Valid values: "print", "screen", "email", "hqprint", "none"
func (b *Blip) SetCompressionState(cstate string) {
	b.SetAttribute(openxml.NewAttribute("", "cstate", "", cstate))
}

// Clone creates a deep copy of this Blip element.
func (b *Blip) Clone() openxml.Element {
	return &Blip{
		CompositeElementBase: b.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// ShapeProperties represents the pic:spPr element.
type ShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewShapeProperties creates a new ShapeProperties element.
func NewShapeProperties(width, height int64) *ShapeProperties {
	elem := openxml.NewCompositeElement(NamespaceDrawingMLPicture, "spPr", PrefixPic)
	sp := &ShapeProperties{CompositeElementBase: elem}

	// Add transform
	xfrm := NewTransform2D(0, 0, width, height)
	sp.AppendChild(xfrm)

	// Add preset geometry (rectangle)
	prstGeom := NewPresetGeometry("rect")
	sp.AppendChild(prstGeom)

	return sp
}

// Transform2D returns the transform element.
func (sp *ShapeProperties) Transform2D() *Transform2D {
	elem := sp.GetElement("xfrm", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if xfrm, ok := elem.(*Transform2D); ok {
		return xfrm
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &Transform2D{CompositeElementBase: comp}
	}
	return nil
}

// PresetGeometry returns the preset geometry element.
func (sp *ShapeProperties) PresetGeometry() *PresetGeometry {
	elem := sp.GetElement("prstGeom", NamespaceDrawingML)
	if elem == nil {
		return nil
	}
	if pg, ok := elem.(*PresetGeometry); ok {
		return pg
	}
	if comp, ok := elem.(*openxml.CompositeElementBase); ok {
		return &PresetGeometry{CompositeElementBase: comp}
	}
	return nil
}

// Clone creates a deep copy of this ShapeProperties element.
func (sp *ShapeProperties) Clone() openxml.Element {
	return &ShapeProperties{
		CompositeElementBase: sp.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// Transform2D represents the a:xfrm element.
type Transform2D struct {
	*openxml.CompositeElementBase
}

// NewTransform2D creates a new Transform2D element.
func NewTransform2D(offX, offY, extCx, extCy int64) *Transform2D {
	elem := openxml.NewCompositeElement(NamespaceDrawingML, "xfrm", PrefixA)
	xfrm := &Transform2D{CompositeElementBase: elem}

	// Add offset
	off := openxml.NewCompositeElement(NamespaceDrawingML, "off", PrefixA)
	off.SetAttribute(openxml.NewAttribute("", "x", "", strconv.FormatInt(offX, 10)))
	off.SetAttribute(openxml.NewAttribute("", "y", "", strconv.FormatInt(offY, 10)))
	xfrm.AppendChild(off)

	// Add extent
	ext := openxml.NewCompositeElement(NamespaceDrawingML, "ext", PrefixA)
	ext.SetAttribute(openxml.NewAttribute("", "cx", "", strconv.FormatInt(extCx, 10)))
	ext.SetAttribute(openxml.NewAttribute("", "cy", "", strconv.FormatInt(extCy, 10)))
	xfrm.AppendChild(ext)

	return xfrm
}

// Rotation returns the rotation in 60,000ths of a degree.
func (xfrm *Transform2D) Rotation() int {
	attr, found := xfrm.GetAttribute("rot", "")
	if !found {
		return 0
	}
	val, _ := strconv.Atoi(attr.Value())
	return val
}

// SetRotation sets the rotation in 60,000ths of a degree.
func (xfrm *Transform2D) SetRotation(rot int) {
	xfrm.SetAttribute(openxml.NewAttribute("", "rot", "", strconv.Itoa(rot)))
}

// FlipH returns whether the shape is flipped horizontally.
func (xfrm *Transform2D) FlipH() bool {
	attr, found := xfrm.GetAttribute("flipH", "")
	if !found {
		return false
	}
	return attr.Value() == "1"
}

// SetFlipH sets whether the shape is flipped horizontally.
func (xfrm *Transform2D) SetFlipH(flip bool) {
	val := "0"
	if flip {
		val = "1"
	}
	xfrm.SetAttribute(openxml.NewAttribute("", "flipH", "", val))
}

// FlipV returns whether the shape is flipped vertically.
func (xfrm *Transform2D) FlipV() bool {
	attr, found := xfrm.GetAttribute("flipV", "")
	if !found {
		return false
	}
	return attr.Value() == "1"
}

// SetFlipV sets whether the shape is flipped vertically.
func (xfrm *Transform2D) SetFlipV(flip bool) {
	val := "0"
	if flip {
		val = "1"
	}
	xfrm.SetAttribute(openxml.NewAttribute("", "flipV", "", val))
}

// Clone creates a deep copy of this Transform2D element.
func (xfrm *Transform2D) Clone() openxml.Element {
	return &Transform2D{
		CompositeElementBase: xfrm.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}

// PresetGeometry represents the a:prstGeom element.
type PresetGeometry struct {
	*openxml.CompositeElementBase
}

// NewPresetGeometry creates a new PresetGeometry element.
func NewPresetGeometry(prst string) *PresetGeometry {
	elem := openxml.NewCompositeElement(NamespaceDrawingML, "prstGeom", PrefixA)
	pg := &PresetGeometry{CompositeElementBase: elem}
	pg.SetAttribute(openxml.NewAttribute("", "prst", "", prst))

	// Add empty avLst
	avLst := openxml.NewCompositeElement(NamespaceDrawingML, "avLst", PrefixA)
	pg.AppendChild(avLst)

	return pg
}

// Preset returns the preset shape type.
func (pg *PresetGeometry) Preset() string {
	attr, found := pg.GetAttribute("prst", "")
	if !found {
		return ""
	}
	return attr.Value()
}

// SetPreset sets the preset shape type.
func (pg *PresetGeometry) SetPreset(prst string) {
	pg.SetAttribute(openxml.NewAttribute("", "prst", "", prst))
}

// Clone creates a deep copy of this PresetGeometry element.
func (pg *PresetGeometry) Clone() openxml.Element {
	return &PresetGeometry{
		CompositeElementBase: pg.CompositeElementBase.Clone().(*openxml.CompositeElementBase),
	}
}
