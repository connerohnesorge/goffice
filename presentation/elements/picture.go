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

// ApplicationNonVisualProperties returns the application non-visual properties (p:nvPr).
func (nvpp *NonVisualPictureProperties) ApplicationNonVisualProperties() *ApplicationNonVisualProperties {
	elem := nvpp.GetElement(
		"nvPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if anvp, ok := elem.(*ApplicationNonVisualProperties); ok {
		return anvp
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &ApplicationNonVisualProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// SetVideoFile sets the video file link.
func (nvpp *NonVisualPictureProperties) SetVideoFile(relId string) {
	anvp := nvpp.ApplicationNonVisualProperties()
	if anvp == nil {
		return
	}
	anvp.SetVideoFile(relId)
}

// SetAudioFile sets the audio file link.
func (nvpp *NonVisualPictureProperties) SetAudioFile(relId string) {
	anvp := nvpp.ApplicationNonVisualProperties()
	if anvp == nil {
		return
	}
	anvp.SetAudioFile(relId)
}

// ApplicationNonVisualProperties represents p:nvPr.
type ApplicationNonVisualProperties struct {
	*openxml.CompositeElementBase
}

// SetVideoFile sets the video file link (a:videoFile).
func (anvp *ApplicationNonVisualProperties) SetVideoFile(relId string) {
	// Remove existing if any
	if existing := anvp.GetElement("videoFile", NamespaceDrawingML); existing != nil {
		anvp.RemoveChild(existing)
	}

	vf := openxml.NewLeafElement(NamespaceDrawingML, "videoFile", PrefixA)
	vf.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "link", PrefixR, relId))

	// Insert before extLst if present
	if extLst := anvp.GetElement("extLst", NamespacePresentationML); extLst != nil {
		anvp.InsertBefore(vf, extLst)
	} else {
		anvp.AppendChild(vf)
	}
}

// SetAudioFile sets the audio file link (a:audioFile).
func (anvp *ApplicationNonVisualProperties) SetAudioFile(relId string) {
	// Remove existing if any
	if existing := anvp.GetElement("audioFile", NamespaceDrawingML); existing != nil {
		anvp.RemoveChild(existing)
	}

	af := openxml.NewLeafElement(NamespaceDrawingML, "audioFile", PrefixA)
	af.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "link", PrefixR, relId))

	// Insert before extLst if present
	if extLst := anvp.GetElement("extLst", NamespacePresentationML); extLst != nil {
		anvp.InsertBefore(af, extLst)
	} else {
		anvp.AppendChild(af)
	}
}

// MediaProperties contains playback options for media.
type MediaProperties struct {
	EmbedRelId string
	AutoStart  bool
	Loop       bool
	Muted      bool
	Volume     int // 0 to 100000
}

// SetMediaProperties sets the media properties in extLst.
func (anvp *ApplicationNonVisualProperties) SetMediaProperties(props MediaProperties) {
	extLst := anvp.GetOrCreateMediaExtensionList()

	// Media extension URI for Office 2010
	mediaUri := "{DAA4B4D4-6D71-4841-9C94-3DE7FCFB9230}"
	ext := extLst.GetOrCreateMediaExtension(mediaUri)

	// Create p14:media element
	nsP14 := "http://schemas.microsoft.com/office/powerpoint/2010/main"
	media := openxml.NewLeafElement(nsP14, "media", "p14")

	if props.EmbedRelId != "" {
		media.SetAttribute(openxml.NewAttribute(NamespaceRelationships, "embed", PrefixR, props.EmbedRelId))
	}

	if props.AutoStart {
		media.SetAttribute(openxml.NewAttribute("", "autoStart", "", "1"))
	}

	if props.Loop {
		media.SetAttribute(openxml.NewAttribute("", "loop", "", "1"))
	}

	if props.Muted {
		media.SetAttribute(openxml.NewAttribute("", "mute", "", "1"))
	}

	if props.Volume > 0 {
		media.SetAttribute(openxml.NewAttribute("", "vol", "", strconv.Itoa(props.Volume)))
	}

	// Clear existing children of extension and add new media element
	ext.ClearChildren()
	ext.AppendChild(media)
}

// GetOrCreateMediaExtensionList returns or creates the extension list (p:extLst).
func (anvp *ApplicationNonVisualProperties) GetOrCreateMediaExtensionList() *MediaExtensionList {
	elem := anvp.GetElement("extLst", NamespacePresentationML)
	if elem != nil {
		if extLst, ok := elem.(*MediaExtensionList); ok {
			return extLst
		}
		if comp := wrapCompositeElement(elem); comp != nil {
			return &MediaExtensionList{CompositeElementBase: comp}
		}
	}

	extLst := NewMediaExtensionList()
	anvp.AppendChild(extLst)
	return extLst
}

// MediaExtensionList represents p:extLst.
type MediaExtensionList struct {
	*openxml.CompositeElementBase
}

// NewMediaExtensionList creates a new MediaExtensionList.
func NewMediaExtensionList() *MediaExtensionList {
	return &MediaExtensionList{
		CompositeElementBase: openxml.NewCompositeElement(NamespacePresentationML, "extLst", PrefixP),
	}
}

// GetOrCreateMediaExtension returns or creates an extension (p:ext) with the given URI.
func (el *MediaExtensionList) GetOrCreateMediaExtension(uri string) *MediaExtension {
	for child := range el.Children() {
		if child.LocalName() == "ext" && child.NamespaceURI() == NamespacePresentationML {
			if attr, ok := child.GetAttribute("uri", ""); ok && attr.Value() == uri {
				if ext, ok := child.(*MediaExtension); ok {
					return ext
				}
				if comp := wrapCompositeElement(child); comp != nil {
					return &MediaExtension{CompositeElementBase: comp}
				}
			}
		}
	}

	ext := NewMediaExtension(uri)
	el.AppendChild(ext)
	return ext
}

// MediaExtension represents p:ext.
type MediaExtension struct {
	*openxml.CompositeElementBase
}

// NewMediaExtension creates a new MediaExtension with the given URI.
func NewMediaExtension(uri string) *MediaExtension {
	ext := &MediaExtension{
		CompositeElementBase: openxml.NewCompositeElement(NamespacePresentationML, "ext", PrefixP),
	}
	ext.SetAttribute(openxml.NewAttribute("", "uri", "", uri))
	return ext
}

// ClearChildren removes all child elements.
func (e *MediaExtension) ClearChildren() {
	var toRemove []openxml.Element
	for child := range e.Children() {
		toRemove = append(toRemove, child)
	}
	for _, child := range toRemove {
		e.RemoveChild(child)
	}
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
