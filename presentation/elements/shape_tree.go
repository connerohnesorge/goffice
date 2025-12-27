//nolint:revive // This file contains shape tree element implementation.
package elements

import (
	"strconv"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// Global shape ID counter for generating unique IDs.
var shapeIdCounter uint32

// nextShapeId returns the next unique shape ID.
func nextShapeId() uint32 {
	return atomic.AddUint32(&shapeIdCounter, 1)
}

// ShapeTree represents the shape tree element (p:spTree).
// This is the container for all visual elements on a slide.
type ShapeTree struct {
	*openxml.CompositeElementBase
}

// NewShapeTree creates a new ShapeTree element.
func NewShapeTree() *ShapeTree {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"spTree",
		PrefixP,
	)
	st := &ShapeTree{CompositeElementBase: elem}

	// Add required non-visual group shape properties
	st.AppendChild(
		NewNonVisualGroupShapeProperties(),
	)
	// Add required group shape properties
	st.AppendChild(NewGroupShapeProperties())

	return st
}

// NonVisualGroupShapeProperties returns the non-visual group shape properties.
func (st *ShapeTree) NonVisualGroupShapeProperties() *NonVisualGroupShapeProperties {
	elem := st.GetElement(
		"nvGrpSpPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if nvgsp, ok := elem.(*NonVisualGroupShapeProperties); ok {
		return nvgsp
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &NonVisualGroupShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// GroupShapeProperties returns the group shape properties.
func (st *ShapeTree) GroupShapeProperties() *GroupShapeProperties {
	elem := st.GetElement(
		"grpSpPr",
		NamespacePresentationML,
	)
	if elem == nil {
		return nil
	}
	if gsp, ok := elem.(*GroupShapeProperties); ok {
		return gsp
	}
	if comp := wrapCompositeElement(elem); comp != nil {
		return &GroupShapeProperties{
			CompositeElementBase: comp,
		}
	}

	return nil
}

// Shapes returns all shape elements in the shape tree.
func (st *ShapeTree) Shapes() []*Shape {
	var shapes []*Shape
	for child := range st.Children() {
		if child.LocalName() == "sp" &&
			child.NamespaceURI() == NamespacePresentationML {
			if sp, ok := child.(*Shape); ok {
				shapes = append(shapes, sp)
			} else if comp := wrapCompositeElement(child); comp != nil {
				shapes = append(shapes, &Shape{CompositeElementBase: comp})
			}
		}
	}

	return shapes
}

// Pictures returns all picture elements in the shape tree.
func (st *ShapeTree) Pictures() []*Picture {
	var pics []*Picture
	for child := range st.Children() {
		if child.LocalName() == "pic" &&
			child.NamespaceURI() == NamespacePresentationML {
			if pic, ok := child.(*Picture); ok {
				pics = append(pics, pic)
			} else if comp := wrapCompositeElement(child); comp != nil {
				pics = append(pics, &Picture{CompositeElementBase: comp})
			}
		}
	}

	return pics
}

// GroupShapes returns all group shape elements in the shape tree.
func (st *ShapeTree) GroupShapes() []*GroupShape {
	var groups []*GroupShape
	for child := range st.Children() {
		if child.LocalName() == "grpSp" &&
			child.NamespaceURI() == NamespacePresentationML {
			if gs, ok := child.(*GroupShape); ok {
				groups = append(groups, gs)
			} else if comp := wrapCompositeElement(child); comp != nil {
				groups = append(groups, &GroupShape{CompositeElementBase: comp})
			}
		}
	}

	return groups
}

// GraphicFrames returns all graphic frame elements in the shape tree.
func (st *ShapeTree) GraphicFrames() []*GraphicFrame {
	var frames []*GraphicFrame
	for child := range st.Children() {
		if child.LocalName() == "graphicFrame" &&
			child.NamespaceURI() == NamespacePresentationML {
			if gf, ok := child.(*GraphicFrame); ok {
				frames = append(frames, gf)
			} else if comp := wrapCompositeElement(child); comp != nil {
				frames = append(frames, &GraphicFrame{CompositeElementBase: comp})
			}
		}
	}

	return frames
}

// ConnectionShapes returns all connection shape elements in the shape tree.
func (st *ShapeTree) ConnectionShapes() []*ConnectionShape {
	var connectors []*ConnectionShape
	for child := range st.Children() {
		if child.LocalName() == "cxnSp" &&
			child.NamespaceURI() == NamespacePresentationML {
			if cs, ok := child.(*ConnectionShape); ok {
				connectors = append(
					connectors,
					cs,
				)
			} else if comp := wrapCompositeElement(child); comp != nil {
				connectors = append(connectors, &ConnectionShape{CompositeElementBase: comp})
			}
		}
	}

	return connectors
}

// AddShape adds a new shape element.
func (st *ShapeTree) AddShape() *Shape {
	sp := NewShape()
	st.AppendChild(sp)

	return sp
}

// AddPicture adds a new picture element.
func (st *ShapeTree) AddPicture(
	relId string,
) *Picture {
	pic := NewPicture(relId)
	st.AppendChild(pic)

	return pic
}

// AddGroupShape adds a new group shape element.
func (st *ShapeTree) AddGroupShape() *GroupShape {
	gs := NewGroupShape()
	st.AppendChild(gs)

	return gs
}

// AddGraphicFrame adds a new graphic frame element.
func (st *ShapeTree) AddGraphicFrame() *GraphicFrame {
	gf := NewGraphicFrame()
	st.AppendChild(gf)

	return gf
}

// AddConnectionShape adds a new connection shape element.
func (st *ShapeTree) AddConnectionShape() *ConnectionShape {
	cs := NewConnectionShape()
	st.AppendChild(cs)

	return cs
}

// Clone creates a deep copy of this ShapeTree element.
func (st *ShapeTree) Clone() openxml.Element {
	cloned := st.CompositeElementBase.Clone()

	return &ShapeTree{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// NonVisualGroupShapeProperties (p:nvGrpSpPr)
// ===========================================================================

// NonVisualGroupShapeProperties represents non-visual group shape properties (p:nvGrpSpPr).
type NonVisualGroupShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewNonVisualGroupShapeProperties creates a new NonVisualGroupShapeProperties element.
func NewNonVisualGroupShapeProperties() *NonVisualGroupShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvGrpSpPr",
		PrefixP,
	)
	nvgsp := &NonVisualGroupShapeProperties{
		CompositeElementBase: elem,
	}

	// Add required child elements
	cNvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvPr",
		PrefixP,
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute("", "id", "", "1"),
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute("", "name", "", ""),
	)
	nvgsp.AppendChild(cNvPr)

	cNvGrpSpPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvGrpSpPr",
		PrefixP,
	)
	nvgsp.AppendChild(cNvGrpSpPr)

	nvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPr",
		PrefixP,
	)
	nvgsp.AppendChild(nvPr)

	return nvgsp
}

// Clone creates a deep copy of this NonVisualGroupShapeProperties element.
func (nvgsp *NonVisualGroupShapeProperties) Clone() openxml.Element {
	cloned := nvgsp.CompositeElementBase.Clone()

	return &NonVisualGroupShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// GroupShapeProperties (p:grpSpPr)
// ===========================================================================

// GroupShapeProperties represents group shape properties (p:grpSpPr).
type GroupShapeProperties struct {
	*openxml.CompositeElementBase
}

// NewGroupShapeProperties creates a new GroupShapeProperties element.
func NewGroupShapeProperties() *GroupShapeProperties {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"grpSpPr",
		PrefixP,
	)

	return &GroupShapeProperties{
		CompositeElementBase: elem,
	}
}

// Clone creates a deep copy of this GroupShapeProperties element.
func (gsp *GroupShapeProperties) Clone() openxml.Element {
	cloned := gsp.CompositeElementBase.Clone()

	return &GroupShapeProperties{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// GroupShape (p:grpSp)
// ===========================================================================

// GroupShape represents a group shape element (p:grpSp).
type GroupShape struct {
	*openxml.CompositeElementBase
}

// NewGroupShape creates a new GroupShape element.
func NewGroupShape() *GroupShape {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"grpSp",
		PrefixP,
	)
	gs := &GroupShape{CompositeElementBase: elem}

	// Add required non-visual properties
	gs.AppendChild(
		NewNonVisualGroupShapeProperties(),
	)
	// Add required group shape properties
	gs.AppendChild(NewGroupShapeProperties())

	return gs
}

// Clone creates a deep copy of this GroupShape element.
func (gs *GroupShape) Clone() openxml.Element {
	cloned := gs.CompositeElementBase.Clone()

	return &GroupShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// GraphicFrame (p:graphicFrame)
// ===========================================================================

// GraphicFrame represents a graphic frame element (p:graphicFrame).
// Used for charts, tables, diagrams, and other OLE objects.
type GraphicFrame struct {
	*openxml.CompositeElementBase
}

// NewGraphicFrame creates a new GraphicFrame element.
func NewGraphicFrame() *GraphicFrame {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"graphicFrame",
		PrefixP,
	)
	gf := &GraphicFrame{
		CompositeElementBase: elem,
	}

	// Add required non-visual properties
	nvGraphicFramePr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvGraphicFramePr",
		PrefixP,
	)

	cNvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvPr",
		PrefixP,
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute("", "id", "", "1"),
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			"Graphic Frame",
		),
	)
	nvGraphicFramePr.AppendChild(cNvPr)

	cNvGraphicFramePr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvGraphicFramePr",
		PrefixP,
	)
	nvGraphicFramePr.AppendChild(
		cNvGraphicFramePr,
	)

	nvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPr",
		PrefixP,
	)
	nvGraphicFramePr.AppendChild(nvPr)

	gf.AppendChild(nvGraphicFramePr)

	// Add transform
	xfrm := openxml.NewCompositeElement(
		NamespacePresentationML,
		"xfrm",
		PrefixP,
	)
	gf.AppendChild(xfrm)

	return gf
}

// SetPosition sets the position of the graphic frame in EMUs.
func (gf *GraphicFrame) SetPosition(x, y int) {
	xfrmElem := gf.GetElement(
		"xfrm",
		NamespacePresentationML,
	)
	if xfrmElem == nil {
		return
	}
	xfrm, ok := xfrmElem.(*openxml.CompositeElementBase)
	if !ok {
		return
	}

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

// SetSize sets the size of the graphic frame in EMUs.
func (gf *GraphicFrame) SetSize(cx, cy int) {
	xfrmElem := gf.GetElement(
		"xfrm",
		NamespacePresentationML,
	)
	if xfrmElem == nil {
		return
	}
	xfrm, ok := xfrmElem.(*openxml.CompositeElementBase)
	if !ok {
		return
	}

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

// Clone creates a deep copy of this GraphicFrame element.
func (gf *GraphicFrame) Clone() openxml.Element {
	cloned := gf.CompositeElementBase.Clone()

	return &GraphicFrame{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// ConnectionShape (p:cxnSp)
// ===========================================================================

// ConnectionShape represents a connection shape element (p:cxnSp).
// Used for connectors between shapes.
type ConnectionShape struct {
	*openxml.CompositeElementBase
}

// NewConnectionShape creates a new ConnectionShape element.
func NewConnectionShape() *ConnectionShape {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cxnSp",
		PrefixP,
	)
	cs := &ConnectionShape{
		CompositeElementBase: elem,
	}

	// Add required non-visual properties
	nvCxnSpPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvCxnSpPr",
		PrefixP,
	)

	cNvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvPr",
		PrefixP,
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute("", "id", "", "1"),
	)
	cNvPr.SetAttribute(
		openxml.NewAttribute(
			"",
			"name",
			"",
			"Connector",
		),
	)
	nvCxnSpPr.AppendChild(cNvPr)

	cNvCxnSpPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"cNvCxnSpPr",
		PrefixP,
	)
	nvCxnSpPr.AppendChild(cNvCxnSpPr)

	nvPr := openxml.NewCompositeElement(
		NamespacePresentationML,
		"nvPr",
		PrefixP,
	)
	nvCxnSpPr.AppendChild(nvPr)

	cs.AppendChild(nvCxnSpPr)

	// Add shape properties
	spPr := openxml.NewCompositeElement(
		NamespaceDrawingML,
		"spPr",
		PrefixA,
	)
	cs.AppendChild(spPr)

	return cs
}

// Clone creates a deep copy of this ConnectionShape element.
func (cs *ConnectionShape) Clone() openxml.Element {
	cloned := cs.CompositeElementBase.Clone()

	return &ConnectionShape{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// ===========================================================================
// ContentPart (p:contentPart) - for embedding external content
// ===========================================================================

// ContentPart represents a content part element (p:contentPart).
// Used for embedding external content like 3D models.
type ContentPart struct {
	*openxml.CompositeElementBase
}

// NewContentPart creates a new ContentPart element.
func NewContentPart(relId string) *ContentPart {
	elem := openxml.NewCompositeElement(
		NamespacePresentationML,
		"contentPart",
		PrefixP,
	)
	cp := &ContentPart{CompositeElementBase: elem}
	cp.SetRelId(relId)

	return cp
}

// RelId returns the relationship ID.
func (cp *ContentPart) RelId() string {
	attr, found := cp.GetAttribute(
		"id",
		NamespaceRelationships,
	)
	if !found {
		return ""
	}

	return attr.Value()
}

// SetRelId sets the relationship ID.
func (cp *ContentPart) SetRelId(relId string) {
	cp.SetAttribute(openxml.NewAttribute(
		NamespaceRelationships,
		"id",
		PrefixR,
		relId,
	))
}

// Clone creates a deep copy of this ContentPart element.
func (cp *ContentPart) Clone() openxml.Element {
	cloned := cp.CompositeElementBase.Clone()

	return &ContentPart{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}
