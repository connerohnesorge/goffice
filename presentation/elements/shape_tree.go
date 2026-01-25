//nolint:revive // This file contains shape tree element implementation.
package elements

import (
	"strconv"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

const (
	localNamePic          = "pic"
	localNameGroupShape   = "grpSp"
	localNameGraphicFrame = "graphicFrame"
	localNameConnShape    = "cxnSp"
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
		if child.LocalName() == localNamePic &&
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
		if child.LocalName() == localNameGroupShape &&
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
		if child.LocalName() == localNameGraphicFrame &&
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
		if child.LocalName() == localNameConnShape &&
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

// Transform returns the transform element (a:xfrm) if it exists.
// Returns nil if no transform element is present.
func (gsp *GroupShapeProperties) Transform() *openxml.CompositeElementBase {
	xfrmElem := gsp.GetElement(
		"xfrm",
		NamespaceDrawingML,
	)
	if xfrmElem == nil {
		return nil
	}
	if xfrm, ok := xfrmElem.(*openxml.CompositeElementBase); ok {
		return xfrm
	}

	return nil
}

// SetTransform sets the transform element (a:xfrm).
// If transform is nil, removes any existing transform element.
func (gsp *GroupShapeProperties) SetTransform(xfrm *openxml.CompositeElementBase) {
	// Remove existing transform
	if existing := gsp.Transform(); existing != nil {
		gsp.RemoveChild(existing)
	}
	// Add new transform if provided
	if xfrm != nil {
		// Insert as first child
		if first := gsp.FirstChild(); first != nil {
			gsp.InsertBefore(xfrm, first)
		} else {
			gsp.AppendChild(xfrm)
		}
	}
}

// GetOrCreateTransform returns the transform element, creating it if it doesn't exist.
func (gsp *GroupShapeProperties) GetOrCreateTransform() *openxml.CompositeElementBase {
	xfrm := gsp.Transform()
	if xfrm != nil {
		return xfrm
	}
	xfrm = openxml.NewCompositeElement(
		NamespaceDrawingML,
		"xfrm",
		PrefixA,
	)
	// Insert as first child
	if first := gsp.FirstChild(); first != nil {
		gsp.InsertBefore(xfrm, first)
	} else {
		gsp.AppendChild(xfrm)
	}

	return xfrm
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

// GroupShapeProperties returns the group shape properties.
func (gs *GroupShape) GroupShapeProperties() *GroupShapeProperties {
	elem := gs.GetElement(
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

// NonVisualGroupShapeProperties returns the non-visual group shape properties.
func (gs *GroupShape) NonVisualGroupShapeProperties() *NonVisualGroupShapeProperties {
	elem := gs.GetElement(
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

// Shapes returns all shape elements within this group.
func (gs *GroupShape) Shapes() []*Shape {
	var shapes []*Shape
	for child := range gs.Children() {
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

// Pictures returns all picture elements within this group.
func (gs *GroupShape) Pictures() []*Picture {
	var pics []*Picture
	for child := range gs.Children() {
		if child.LocalName() == localNamePic &&
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

// GroupShapes returns all nested group shape elements within this group.
func (gs *GroupShape) GroupShapes() []*GroupShape {
	var groups []*GroupShape
	for child := range gs.Children() {
		if child.LocalName() == localNameGroupShape &&
			child.NamespaceURI() == NamespacePresentationML {
			if g, ok := child.(*GroupShape); ok {
				groups = append(groups, g)
			} else if comp := wrapCompositeElement(child); comp != nil {
				groups = append(groups, &GroupShape{CompositeElementBase: comp})
			}
		}
	}

	return groups
}

// ConnectionShapes returns all connection shape elements within this group.
func (gs *GroupShape) ConnectionShapes() []*ConnectionShape {
	var connectors []*ConnectionShape
	for child := range gs.Children() {
		if child.LocalName() == localNameConnShape &&
			child.NamespaceURI() == NamespacePresentationML {
			if cs, ok := child.(*ConnectionShape); ok {
				connectors = append(connectors, cs)
			} else if comp := wrapCompositeElement(child); comp != nil {
				connectors = append(connectors, &ConnectionShape{CompositeElementBase: comp})
			}
		}
	}

	return connectors
}

// GraphicFrames returns all graphic frame elements within this group.
func (gs *GroupShape) GraphicFrames() []*GraphicFrame {
	var frames []*GraphicFrame
	for child := range gs.Children() {
		if child.LocalName() == localNameGraphicFrame &&
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

// AddShape adds a new shape element to this group.
func (gs *GroupShape) AddShape() *Shape {
	sp := NewShape()
	gs.AppendChild(sp)

	return sp
}

// AddPicture adds a new picture element to this group.
func (gs *GroupShape) AddPicture(relId string) *Picture {
	pic := NewPicture(relId)
	gs.AppendChild(pic)

	return pic
}

// AddGroupShape adds a new nested group shape element to this group.
func (gs *GroupShape) AddGroupShape() *GroupShape {
	nested := NewGroupShape()
	gs.AppendChild(nested)

	return nested
}

// AddConnectionShape adds a new connection shape element to this group.
func (gs *GroupShape) AddConnectionShape() *ConnectionShape {
	cs := NewConnectionShape()
	gs.AppendChild(cs)

	return cs
}

// AddGraphicFrame adds a new graphic frame element to this group.
func (gs *GroupShape) AddGraphicFrame() *GraphicFrame {
	gf := NewGraphicFrame()
	gs.AppendChild(gf)

	return gf
}

// RemoveChild removes a child element from this group.
func (gs *GroupShape) RemoveChild(child openxml.Element) {
	gs.CompositeElementBase.RemoveChild(child)
}

// Clear removes all child shapes from this group.
// This preserves the non-visual properties and group shape properties,
// only removing shape, picture, group, connection, and graphic frame children.
func (gs *GroupShape) Clear() {
	// Collect children to remove (can't modify while iterating)
	var toRemove []openxml.Element
	for child := range gs.Children() {
		switch child.LocalName() {
		case "sp", "pic", "grpSp", "cxnSp", "graphicFrame":
			if child.NamespaceURI() == NamespacePresentationML {
				toRemove = append(toRemove, child)
			}
		}
	}

	// Remove collected children
	for _, child := range toRemove {
		gs.RemoveChild(child)
	}
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

// ConnectionInfo represents information about a shape connection.
type ConnectionInfo struct {
	ShapeID             string
	ConnectionSiteIndex int
}

// StartConnection returns the start connection information.
// ok is false if no start connection is set.
func (cs *ConnectionShape) StartConnection() (conn ConnectionInfo, ok bool) {
	// Navigate to nvCxnSpPr/cNvCxnSpPr element
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return ConnectionInfo{}, false
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return ConnectionInfo{}, false
	}

	// Look for a:stCxn element
	stCxn := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML)
	if stCxn == nil {
		return ConnectionInfo{}, false
	}

	// Read id and idx attributes
	idAttr, foundID := stCxn.GetAttribute("id", "")
	if !foundID {
		return ConnectionInfo{}, false
	}

	idxAttr, foundIdx := stCxn.GetAttribute("idx", "")
	if !foundIdx {
		return ConnectionInfo{}, false
	}

	idx, err := strconv.Atoi(idxAttr.Value())
	if err != nil {
		return ConnectionInfo{}, false
	}

	return ConnectionInfo{
		ShapeID:             idAttr.Value(),
		ConnectionSiteIndex: idx,
	}, true
}

// SetStartConnection sets the start connection to a shape's connection site.
func (cs *ConnectionShape) SetStartConnection(shapeID string, connectionSiteIndex int) {
	// Get nvCxnSpPr/cNvCxnSpPr structure
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return
	}

	// Remove existing stCxn if present
	if existing := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML); existing != nil {
		cNvCxnSpPr.RemoveChild(existing)
	}

	// Create new a:stCxn element
	stCxn := openxml.NewCompositeElement(NamespaceDrawingML, "stCxn", PrefixA)
	stCxn.SetAttribute(openxml.NewAttribute("", "id", "", shapeID))
	stCxn.SetAttribute(openxml.NewAttribute("", "idx", "", strconv.Itoa(connectionSiteIndex)))

	// Insert before endCxn if present, otherwise append
	if endCxn := cNvCxnSpPr.GetElement("endCxn", NamespaceDrawingML); endCxn != nil {
		cNvCxnSpPr.InsertBefore(stCxn, endCxn)
	} else {
		cNvCxnSpPr.AppendChild(stCxn)
	}
}

// EndConnection returns the end connection information.
// ok is false if no end connection is set.
func (cs *ConnectionShape) EndConnection() (conn ConnectionInfo, ok bool) {
	// Navigate to nvCxnSpPr/cNvCxnSpPr element
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return ConnectionInfo{}, false
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return ConnectionInfo{}, false
	}

	// Look for a:endCxn element
	endCxn := cNvCxnSpPr.GetElement("endCxn", NamespaceDrawingML)
	if endCxn == nil {
		return ConnectionInfo{}, false
	}

	// Read id and idx attributes
	idAttr, foundID := endCxn.GetAttribute("id", "")
	if !foundID {
		return ConnectionInfo{}, false
	}

	idxAttr, foundIdx := endCxn.GetAttribute("idx", "")
	if !foundIdx {
		return ConnectionInfo{}, false
	}

	idx, err := strconv.Atoi(idxAttr.Value())
	if err != nil {
		return ConnectionInfo{}, false
	}

	return ConnectionInfo{
		ShapeID:             idAttr.Value(),
		ConnectionSiteIndex: idx,
	}, true
}

// SetEndConnection sets the end connection to a shape's connection site.
func (cs *ConnectionShape) SetEndConnection(shapeID string, connectionSiteIndex int) {
	// Get nvCxnSpPr/cNvCxnSpPr structure
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return
	}

	// Remove existing endCxn if present
	if existing := cNvCxnSpPr.GetElement("endCxn", NamespaceDrawingML); existing != nil {
		cNvCxnSpPr.RemoveChild(existing)
	}

	// Create new a:endCxn element
	endCxn := openxml.NewCompositeElement(NamespaceDrawingML, "endCxn", PrefixA)
	endCxn.SetAttribute(openxml.NewAttribute("", "id", "", shapeID))
	endCxn.SetAttribute(openxml.NewAttribute("", "idx", "", strconv.Itoa(connectionSiteIndex)))

	// Append endCxn
	cNvCxnSpPr.AppendChild(endCxn)
}

// ClearStartConnection removes the start connection.
func (cs *ConnectionShape) ClearStartConnection() {
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return
	}

	if stCxn := cNvCxnSpPr.GetElement("stCxn", NamespaceDrawingML); stCxn != nil {
		cNvCxnSpPr.RemoveChild(stCxn)
	}
}

// ClearEndConnection removes the end connection.
func (cs *ConnectionShape) ClearEndConnection() {
	nvCxnSpPr := wrapCompositeElement(cs.GetElement("nvCxnSpPr", NamespacePresentationML))
	if nvCxnSpPr == nil {
		return
	}

	cNvCxnSpPr := wrapCompositeElement(nvCxnSpPr.GetElement("cNvCxnSpPr", NamespacePresentationML))
	if cNvCxnSpPr == nil {
		return
	}

	if endCxn := cNvCxnSpPr.GetElement("endCxn", NamespaceDrawingML); endCxn != nil {
		cNvCxnSpPr.RemoveChild(endCxn)
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
