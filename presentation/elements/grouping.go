//nolint:revive // This file contains high-level grouping operations.
package elements

import (
	"math"

	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/openxml"
)

// computeBoundingBox calculates the bounding box that encloses all shapes.
// Returns (x, y, cx, cy) where x,y is the top-left corner and cx,cy is the size.
// All values are in EMUs (English Metric Units).
func computeBoundingBox(shapes []interface{}) (x, y, cx, cy drawingml.EMU) {
	if len(shapes) == 0 {
		return 0, 0, 0, 0
	}

	minX := drawingml.EMU(math.MaxInt64)
	minY := drawingml.EMU(math.MaxInt64)
	maxX := drawingml.EMU(math.MinInt64)
	maxY := drawingml.EMU(math.MinInt64)

	for _, shape := range shapes {
		var shapeX, shapeY, shapeCx, shapeCy drawingml.EMU

		// Extract transform based on shape type
		switch s := shape.(type) {
		case *Shape:
			if spPr := s.ShapeProperties(); spPr != nil {
				shapeX, shapeY, shapeCx, shapeCy = getTransformFromProperties(spPr)
			}
		case *Picture:
			if spPr := s.ShapeProperties(); spPr != nil {
				shapeX, shapeY, shapeCx, shapeCy = getTransformFromProperties(spPr)
			}
		case *GroupShape:
			if gspPr := s.GroupShapeProperties(); gspPr != nil {
				if xfrm := gspPr.Transform(); xfrm != nil {
					shapeX, shapeY, shapeCx, shapeCy = getTransformBounds(xfrm)
				}
			}
		case *ConnectionShape:
			// ConnectionShape has spPr (DrawingML namespace) but no accessor method
			if spPrElem := s.GetElement("spPr", NamespaceDrawingML); spPrElem != nil {
				if comp, ok := spPrElem.(*openxml.CompositeElementBase); ok {
					spPr := &ShapeProperties{CompositeElementBase: comp}
					shapeX, shapeY, shapeCx, shapeCy = getTransformFromProperties(spPr)
				}
			}
		case *GraphicFrame:
			// GraphicFrame has xfrm directly (PresentationML namespace)
			if xfrmElem := s.GetElement("xfrm", NamespacePresentationML); xfrmElem != nil {
				shapeX, shapeY, shapeCx, shapeCy = getTransformBounds(xfrmElem)
			}
		default:
			// Skip unknown shape types
			continue
		}

		// Update bounding box
		if shapeX < minX {
			minX = shapeX
		}
		if shapeY < minY {
			minY = shapeY
		}
		if shapeX+shapeCx > maxX {
			maxX = shapeX + shapeCx
		}
		if shapeY+shapeCy > maxY {
			maxY = shapeY + shapeCy
		}
	}

	// If no valid shapes found, return zero
	if minX == drawingml.EMU(math.MaxInt64) {
		return 0, 0, 0, 0
	}

	return minX, minY, maxX - minX, maxY - minY
}

// getTransformFromProperties extracts transform bounds from ShapeProperties.
// Returns (x, y, cx, cy) in EMUs.
func getTransformFromProperties(spPr *ShapeProperties) (x, y, cx, cy drawingml.EMU) {
	xfrmElem := spPr.GetElement("xfrm", NamespaceDrawingML)
	if xfrmElem == nil {
		return 0, 0, 0, 0
	}

	return getTransformBounds(xfrmElem)
}

// getTransformBounds extracts offset and extent from a transform element.
// Returns (x, y, cx, cy) in EMUs.
func getTransformBounds(xfrmElem interface{}) (x, y, cx, cy drawingml.EMU) {
	// Try to cast to Transform2D first
	if xfrm, ok := xfrmElem.(*drawingml.Transform2D); ok {
		off := xfrm.Offset()
		ext := xfrm.Extent()

		return off.X, off.Y, ext.Cx, ext.Cy
	}

	// Fallback to CompositeElementBase
	if comp, ok := xfrmElem.(*openxml.CompositeElementBase); ok {
		// Wrap it in a Transform2D for convenience
		xfrm := &drawingml.Transform2D{CompositeElementBase: comp}
		off := xfrm.Offset()
		ext := xfrm.Extent()

		return off.X, off.Y, ext.Cx, ext.Cy
	}

	return 0, 0, 0, 0
}

// adjustCoordinatesRelativeToGroup adjusts a shape's position to be relative
// to the group's origin instead of absolute slide coordinates.
func adjustCoordinatesRelativeToGroup(shape interface{}, groupOffsetX, groupOffsetY drawingml.EMU) {
	switch s := shape.(type) {
	case *Shape:
		if spPr := s.ShapeProperties(); spPr != nil {
			adjustShapePropertiesOffset(spPr, groupOffsetX, groupOffsetY)
		}
	case *Picture:
		if spPr := s.ShapeProperties(); spPr != nil {
			adjustShapePropertiesOffset(spPr, groupOffsetX, groupOffsetY)
		}
	case *GroupShape:
		if gspPr := s.GroupShapeProperties(); gspPr != nil {
			if xfrm := gspPr.Transform(); xfrm != nil {
				adjustTransformOffset(xfrm, groupOffsetX, groupOffsetY)
			}
		}
	case *ConnectionShape:
		// ConnectionShape has spPr (DrawingML namespace) but no accessor method
		if spPrElem := s.GetElement("spPr", NamespaceDrawingML); spPrElem != nil {
			if comp, ok := spPrElem.(*openxml.CompositeElementBase); ok {
				spPr := &ShapeProperties{CompositeElementBase: comp}
				adjustShapePropertiesOffset(spPr, groupOffsetX, groupOffsetY)
			}
		}
	case *GraphicFrame:
		// GraphicFrame has xfrm directly (PresentationML namespace)
		if xfrmElem := s.GetElement("xfrm", NamespacePresentationML); xfrmElem != nil {
			adjustTransformOffset(xfrmElem, groupOffsetX, groupOffsetY)
		}
	}
}

// adjustShapePropertiesOffset adjusts the offset in shape properties.
func adjustShapePropertiesOffset(spPr *ShapeProperties, groupOffsetX, groupOffsetY drawingml.EMU) {
	xfrmElem := spPr.GetElement("xfrm", NamespaceDrawingML)
	if xfrmElem == nil {
		return
	}

	adjustTransformOffset(xfrmElem, groupOffsetX, groupOffsetY)
}

// adjustTransformOffset subtracts the group offset from the shape's offset.
func adjustTransformOffset(xfrmElem interface{}, groupOffsetX, groupOffsetY drawingml.EMU) {
	// Try to cast to Transform2D first
	if xfrm, ok := xfrmElem.(*drawingml.Transform2D); ok {
		off := xfrm.Offset()
		xfrm.SetOffset(drawingml.Offset{
			X: off.X - groupOffsetX,
			Y: off.Y - groupOffsetY,
		})

		return
	}

	// Fallback to CompositeElementBase
	if comp, ok := xfrmElem.(*openxml.CompositeElementBase); ok {
		xfrm := &drawingml.Transform2D{CompositeElementBase: comp}
		off := xfrm.Offset()
		xfrm.SetOffset(drawingml.Offset{
			X: off.X - groupOffsetX,
			Y: off.Y - groupOffsetY,
		})
	}
}

// CreateGroup groups the given shapes together on the slide.
// The shapes are removed from the slide's shape tree, added to a new GroupShape,
// and the GroupShape is added to the slide's shape tree.
// Child shape coordinates are adjusted to be relative to the group origin.
func (s *Slide) CreateGroup(shapes ...interface{}) *GroupShape {
	if len(shapes) == 0 {
		return nil
	}

	st := s.GetOrCreateShapeTree()

	// Calculate bounding box of all shapes
	x, y, cx, cy := computeBoundingBox(shapes)

	// Create new GroupShape
	group := NewGroupShape()

	// Set group's transform (offset = bounding box origin, extent = bounding box size)
	gspPr := group.GroupShapeProperties()
	if gspPr != nil {
		xfrm := gspPr.GetOrCreateTransform()
		if xfrm != nil {
			// Wrap in Transform2D for convenience
			xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}
			xfrmT2D.SetOffset(drawingml.Offset{X: x, Y: y})
			xfrmT2D.SetExtent(drawingml.Extent{Cx: cx, Cy: cy})

			// Set childOffset and childExtent to define the child coordinate space
			// childOffset = group origin (same as offset)
			// childExtent = group size (same as extent)
			xfrmT2D.SetChildOffset(x, y)
			xfrmT2D.SetChildExtent(cx, cy)
		}
	}

	// For each shape:
	// 1. Remove from slide's shape tree
	// 2. Adjust coordinates relative to group origin
	// 3. Add to GroupShape
	for _, shape := range shapes {
		// Remove from slide
		if elem, ok := shape.(openxml.Element); ok {
			st.RemoveChild(elem)
		}

		// Adjust coordinates relative to group origin
		adjustCoordinatesRelativeToGroup(shape, x, y)

		// Add to group
		if elem, ok := shape.(openxml.Element); ok {
			group.AppendChild(elem)
		}
	}

	// Add GroupShape to slide's shape tree
	st.AppendChild(group)

	return group
}

// UngroupShape removes the group and adds its children back to the slide.
// Computes absolute transforms for all children and removes the group from the slide.
// Returns the slice of ungrouped shapes.
func (s *Slide) UngroupShape(group *GroupShape) []interface{} {
	st := s.GetOrCreateShapeTree()

	// Get group's transform
	var groupOffsetX, groupOffsetY drawingml.EMU
	if gspPr := group.GroupShapeProperties(); gspPr != nil {
		if xfrm := gspPr.Transform(); xfrm != nil {
			// Wrap in Transform2D for convenience
			xfrmT2D := &drawingml.Transform2D{CompositeElementBase: xfrm}
			off := xfrmT2D.Offset()
			groupOffsetX = off.X
			groupOffsetY = off.Y
		}
	}

	// Collect all children and compute their absolute positions
	ungroupedShapes := make([]interface{}, 0, 10)

	// Iterate over children
	for child := range group.Children() {
		// Skip non-visual properties
		if child.LocalName() == "nvGrpSpPr" || child.LocalName() == "grpSpPr" {
			continue
		}

		// Store the child
		ungroupedShapes = append(ungroupedShapes, child)
	}

	// Remove group from slide
	st.RemoveChild(group)

	// For each child shape:
	// 1. Set absolute position (group offset + child offset)
	// 2. Add to slide's shape tree
	for _, shape := range ungroupedShapes {
		// Adjust to absolute coordinates (add group offset back)
		adjustCoordinatesToAbsolute(shape, groupOffsetX, groupOffsetY)

		// Add to slide
		if elem, ok := shape.(openxml.Element); ok {
			st.AppendChild(elem)
		}
	}

	return ungroupedShapes
}

// adjustCoordinatesToAbsolute adjusts a shape's position from group-relative
// to absolute slide coordinates.
func adjustCoordinatesToAbsolute(shape interface{}, groupOffsetX, groupOffsetY drawingml.EMU) {
	switch s := shape.(type) {
	case *Shape:
		if spPr := s.ShapeProperties(); spPr != nil {
			adjustShapePropertiesOffsetAbsolute(spPr, groupOffsetX, groupOffsetY)
		}
	case *Picture:
		if spPr := s.ShapeProperties(); spPr != nil {
			adjustShapePropertiesOffsetAbsolute(spPr, groupOffsetX, groupOffsetY)
		}
	case *GroupShape:
		if gspPr := s.GroupShapeProperties(); gspPr != nil {
			if xfrm := gspPr.Transform(); xfrm != nil {
				adjustTransformOffsetAbsolute(xfrm, groupOffsetX, groupOffsetY)
			}
		}
	case *ConnectionShape:
		// ConnectionShape has spPr (DrawingML namespace) but no accessor method
		if spPrElem := s.GetElement("spPr", NamespaceDrawingML); spPrElem != nil {
			if comp, ok := spPrElem.(*openxml.CompositeElementBase); ok {
				spPr := &ShapeProperties{CompositeElementBase: comp}
				adjustShapePropertiesOffsetAbsolute(spPr, groupOffsetX, groupOffsetY)
			}
		}
	case *GraphicFrame:
		// GraphicFrame has xfrm directly (PresentationML namespace)
		if xfrmElem := s.GetElement("xfrm", NamespacePresentationML); xfrmElem != nil {
			adjustTransformOffsetAbsolute(xfrmElem, groupOffsetX, groupOffsetY)
		}
	}
}

// adjustShapePropertiesOffsetAbsolute adjusts the offset in shape properties to absolute.
func adjustShapePropertiesOffsetAbsolute(spPr *ShapeProperties, groupOffsetX, groupOffsetY drawingml.EMU) {
	xfrmElem := spPr.GetElement("xfrm", NamespaceDrawingML)
	if xfrmElem == nil {
		return
	}

	adjustTransformOffsetAbsolute(xfrmElem, groupOffsetX, groupOffsetY)
}

// adjustTransformOffsetAbsolute adds the group offset to the shape's offset.
func adjustTransformOffsetAbsolute(xfrmElem interface{}, groupOffsetX, groupOffsetY drawingml.EMU) {
	// Try to cast to Transform2D first
	if xfrm, ok := xfrmElem.(*drawingml.Transform2D); ok {
		off := xfrm.Offset()
		xfrm.SetOffset(drawingml.Offset{
			X: off.X + groupOffsetX,
			Y: off.Y + groupOffsetY,
		})

		return
	}

	// Fallback to CompositeElementBase
	if comp, ok := xfrmElem.(*openxml.CompositeElementBase); ok {
		xfrm := &drawingml.Transform2D{CompositeElementBase: comp}
		off := xfrm.Offset()
		xfrm.SetOffset(drawingml.Offset{
			X: off.X + groupOffsetX,
			Y: off.Y + groupOffsetY,
		})
	}
}
