//nolint:revive // This file contains picture element implementation.
package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Picture represents a picture element (p:pic).
type Picture struct {
	*openxml.CompositeElementBase
}

// NewPicture creates a new Picture element.
func NewPicture(relId string) *Picture {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"pic",
		PrefixP,
	)
	pic := &Picture{CompositeElementBase: elem}

	// Add required non-visual picture properties
	pic.AppendChild(
		NewNonVisualPictureProperties(),
	)
	// Add required blip fill
	pic.AppendChild(NewBlipFill(relId))
	// Add required shape properties
	pic.AppendChild(NewShapeProperties())

	return pic
}

// NonVisualPictureProperties returns the non-visual picture properties.
func (pic *Picture) NonVisualPictureProperties() *NonVisualPictureProperties {
	elem := pic.GetElement(
		"nvPicPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if nvpp, ok := elem.(*NonVisualPictureProperties); ok {
		return nvpp
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &NonVisualPictureProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// BlipFill returns the blip fill element.
func (pic *Picture) BlipFill() *BlipFill {
	elem := pic.GetElement(
		"blipFill",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if bf, ok := elem.(*BlipFill); ok {
		return bf
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &BlipFill{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// ShapeProperties returns the shape properties.
func (pic *Picture) ShapeProperties() *ShapeProperties {
	elem := pic.GetElement(
		"spPr",
		NamespacePresentationML,
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
func (pic *Picture) GetOrCreateShapeProperties() *ShapeProperties {
	spPr := pic.ShapeProperties()
	if spPr != nil {
		return spPr
	}
	spPr = NewShapeProperties()
	pic.AppendChild(spPr)

	return spPr
}

// SetPosition sets the position of the picture in EMUs.
func (pic *Picture) SetPosition(x, y int) {
	spPr := pic.GetOrCreateShapeProperties()
	spPr.SetOffset(x, y)
}

// SetSize sets the size of the picture in EMUs.
func (pic *Picture) SetSize(cx, cy int) {
	spPr := pic.GetOrCreateShapeProperties()
	spPr.SetExtents(cx, cy)
}

// SetRelId sets the relationship ID for the image.
func (pic *Picture) SetRelId(relId string) {
	bf := pic.BlipFill()
	if bf == nil {
		return
	}
	bf.SetRelId(relId)
}

// Clone creates a deep copy of this Picture element.
func (pic *Picture) Clone() openxml.Element {
	cloned := pic.CompositeElementBase.Clone()

	return &Picture{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// NonVisualPictureProperties (p:nvPicPr)
// ===========================================================================

// NonVisualPictureProperties represents non-visual picture properties (p:nvPicPr).
type NonVisualPictureProperties struct {
	*openxml.CompositeElementBase
}

// NewNonVisualPictureProperties creates a new NonVisualPictureProperties element.
func NewNonVisualPictureProperties() *NonVisualPictureProperties {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPicPr",
		PrefixP,
	)
	nvpp := &NonVisualPictureProperties{
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
			"Picture "+strconv.FormatUint(
				uint64(id),
				10,
			),
		),
	)
	nvpp.AppendChild(cNvPr)

	// Add required cNvPicPr (common non-visual picture properties)
	cNvPicPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvPicPr",
		PrefixP,
	)
	nvpp.AppendChild(cNvPicPr)

	// Add required nvPr (non-visual properties)
	nvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPr",
		PrefixP,
	)
	nvpp.AppendChild(nvPr)

	return nvpp
}

// Id returns the picture ID.
func (nvpp *NonVisualPictureProperties) Id() uint32 {
	cNvPr := nvpp.GetElement(
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

// SetId sets the picture ID.
func (nvpp *NonVisualPictureProperties) SetId(
	id uint32,
) {
	cNvPr := nvpp.GetElement(
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

// Name returns the picture name.
func (nvpp *NonVisualPictureProperties) Name() string {
	cNvPr := nvpp.GetElement(
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

// SetName sets the picture name.
func (nvpp *NonVisualPictureProperties) SetName(
	name string,
) {
	cNvPr := nvpp.GetElement(
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

// Clone creates a deep copy of this NonVisualPictureProperties element.
func (nvpp *NonVisualPictureProperties) Clone() openxml.Element {
	cloned := nvpp.CompositeElementBase.Clone()

	return &NonVisualPictureProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// BlipFill (p:blipFill)
// ===========================================================================

// BlipFill represents a blip fill element (p:blipFill).
type BlipFill struct {
	*openxml.CompositeElementBase
}

// NewBlipFill creates a new BlipFill element.
func NewBlipFill(relId string) *BlipFill {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"blipFill",
		PrefixP,
	)
	bf := &BlipFill{CompositeElementBase: elem}

	// Add blip element with relationship reference
	blip := openxml.NewLeafElement(
		NamespaceDrawingML,
		"blip",
		PrefixA,
	)
	blip.SetAttribute(openxml.NewAttribute(
		NamespaceRelationships,
		"embed",
		PrefixR,
		relId,
	))
	bf.AppendChild(blip)

	// Add stretch element with fillRect
	stretch := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"stretch",
		PrefixA,
	)
	fillRect := openxml.NewLeafElement(
		NamespaceDrawingML,
		"fillRect",
		PrefixA,
	)
	stretch.AppendChild(fillRect)
	bf.AppendChild(stretch)

	return bf
}

// RelId returns the relationship ID for the image.
func (bf *BlipFill) RelId() string {
	blip := bf.GetElement(
		"blip",
		NamespaceDrawingML,
	)
	if blip == nil {
		return ""
	}
	attr, found := blip.GetAttribute(
		"embed",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelId sets the relationship ID for the image.
func (bf *BlipFill) SetRelId(relId string) {
	blip := bf.GetElement(
		"blip",
		NamespaceDrawingML,
	)
	if blip == nil {
		return
	}
	blip.SetAttribute(openxml.NewAttribute(
		NamespaceRelationships,
		"embed",
		PrefixR,
		relId,
	))
}

// Clone creates a deep copy of this BlipFill element.
func (bf *BlipFill) Clone() openxml.Element {
	cloned := bf.CompositeElementBase.Clone()

	return &BlipFill{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
