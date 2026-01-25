// This file implements the high-level Drawing/Image API.
package spreadsheet

import (
	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
)

// EMU conversion constants.
const (
	// emusPerPixel is the number of EMUs (English Metric Units) per pixel
	// at standard 96 DPI. 914400 EMUs = 1 inch, 96 pixels = 1 inch.
	emusPerPixel = 9525

	// emusPerInch is the number of EMUs per inch.
	emusPerInch = 914400

	// emusPerCm is the number of EMUs per centimeter.
	// 914400 / 2.54 = 360000 EMUs per cm.
	emusPerCm = 360000
)

// Drawing represents a collection of drawing objects in a worksheet.
type Drawing struct {
	sheet        *Sheet
	drawingsPart *parts.DrawingsPart
}

// findImagePartByRelID finds an image part by its relationship ID.
func (d *Drawing) findImagePartByRelID(
	relID string,
) *parts.ImagePart {
	for _, ip := range d.drawingsPart.ImageParts() {
		if ip.RelationshipID() == relID {
			return ip
		}
	}

	return nil
}

// parseOneCellAnchorImages extracts images from one-cell anchors.
func (d *Drawing) parseOneCellAnchorImages(
	oneCellAnchor *elements.OneCellAnchor,
) *Image {
	pic := oneCellAnchor.Picture()
	if pic == nil {
		return nil
	}

	// Find the image part by relationship ID
	imagePart := d.findImagePartForPicture(pic)
	if imagePart == nil {
		return nil
	}

	// Extract position information from the from marker
	from := oneCellAnchor.From()
	var fromCell CellRef
	var offsetX, offsetY int64
	if from != nil {
		// Convert from 0-based marker to 1-based CellRef
		fromCell = NewCellRef(
			from.Row()+1,
			from.Col()+1,
		)
		offsetX = from.ColOff()
		offsetY = from.RowOff()
	}

	// Extract size from extent
	var width, height int64
	if ext := oneCellAnchor.Ext(); ext != nil {
		width = ext.Cx()
		height = ext.Cy()
	}

	return &Image{
		sheet:        d.sheet,
		drawingsPart: d.drawingsPart,
		imagePart:    imagePart,
		fromCell:     fromCell,
		offsetX:      offsetX,
		offsetY:      offsetY,
		width:        width,
		height:       height,
		anchor:       oneCellAnchor,
	}
}

// parseTwoCellAnchorImages extracts images from two-cell anchors.
func (d *Drawing) parseTwoCellAnchorImages(
	twoCellAnchor *elements.TwoCellAnchor,
) *Image {
	pic := twoCellAnchor.Picture()
	if pic == nil {
		return nil
	}

	// Find the image part by relationship ID
	imagePart := d.findImagePartForPicture(pic)
	if imagePart == nil {
		return nil
	}

	// Extract position information
	from := twoCellAnchor.From()
	to := twoCellAnchor.To()
	var fromCell, toCell CellRef
	var fromOffsetX, fromOffsetY, toOffsetX, toOffsetY int64

	if from != nil {
		// Convert from 0-based marker to 1-based CellRef
		fromCell = NewCellRef(
			from.Row()+1,
			from.Col()+1,
		)
		fromOffsetX = from.ColOff()
		fromOffsetY = from.RowOff()
	}
	if to != nil {
		// Convert from 0-based marker to 1-based CellRef
		toCell = NewCellRef(
			to.Row()+1,
			to.Col()+1,
		)
		toOffsetX = to.ColOff()
		toOffsetY = to.RowOff()
	}

	return &Image{
		sheet:        d.sheet,
		drawingsPart: d.drawingsPart,
		imagePart:    imagePart,
		fromCell:     fromCell,
		toCell:       &toCell,
		offsetX:      fromOffsetX,
		offsetY:      fromOffsetY,
		toOffsetX:    toOffsetX,
		toOffsetY:    toOffsetY,
		anchor:       twoCellAnchor,
	}
}

// findImagePartForPicture finds the image part for a picture element.
func (d *Drawing) findImagePartForPicture(
	pic *elements.DrawingPicture,
) *parts.ImagePart {
	blipFill := pic.BlipFill()
	if blipFill == nil {
		return nil
	}

	blip := blipFill.Blip()
	if blip == nil {
		return nil
	}

	relID := blip.Embed()

	return d.findImagePartByRelID(relID)
}

// Images returns all images in the drawing.
func (d *Drawing) Images() []*Image {
	if d.drawingsPart == nil {
		return nil
	}

	drawing := d.drawingsPart.Drawing()
	if drawing == nil {
		return nil
	}

	var images []*Image

	// Parse one-cell anchors for pictures
	for _, oneCellAnchor := range drawing.OneCellAnchors() {
		if img := d.parseOneCellAnchorImages(oneCellAnchor); img != nil {
			images = append(images, img)
		}
	}

	// Parse two-cell anchors for pictures
	for _, twoCellAnchor := range drawing.TwoCellAnchors() {
		if img := d.parseTwoCellAnchorImages(twoCellAnchor); img != nil {
			images = append(images, img)
		}
	}

	return images
}

// parseOneCellAnchorShapes extracts shapes from one-cell anchors.
func (*Drawing) parseOneCellAnchorShapes(
	oneCellAnchor *elements.OneCellAnchor,
) *Shape {
	shapeElem := oneCellAnchor.Shape()
	if shapeElem == nil {
		return nil
	}

	// Extract position information
	from := oneCellAnchor.From()
	var fromCell CellRef
	if from != nil {
		// Convert from 0-based marker to 1-based CellRef
		fromCell = NewCellRef(
			from.Row()+1,
			from.Col()+1,
		)
	}

	shape := &Shape{
		fromCell: fromCell,
		anchor:   oneCellAnchor,
	}

	// Try to extract text from text body
	if txBody := shapeElem.TxBody(); txBody != nil {
		// Text extraction would require parsing paragraph/run structure
		// For now, we just note that text body exists
		shape.text = "" // Would need deeper parsing
	}

	return shape
}

// parseTwoCellAnchorShapes extracts shapes from two-cell anchors.
func (*Drawing) parseTwoCellAnchorShapes(
	twoCellAnchor *elements.TwoCellAnchor,
) *Shape {
	shapeElem := twoCellAnchor.Shape()
	if shapeElem == nil {
		return nil
	}

	// Extract position information
	from := twoCellAnchor.From()
	to := twoCellAnchor.To()
	var fromCell, toCell CellRef

	if from != nil {
		// Convert from 0-based marker to 1-based CellRef
		fromCell = NewCellRef(
			from.Row()+1,
			from.Col()+1,
		)
	}
	if to != nil {
		// Convert from 0-based marker to 1-based CellRef
		toCell = NewCellRef(
			to.Row()+1,
			to.Col()+1,
		)
	}

	shape := &Shape{
		fromCell: fromCell,
		toCell:   &toCell,
		anchor:   twoCellAnchor,
	}

	// Try to extract text from text body
	if txBody := shapeElem.TxBody(); txBody != nil {
		// Text extraction would require parsing paragraph/run structure
		shape.text = "" // Would need deeper parsing
	}

	return shape
}

// Shapes returns all shapes in the drawing.
func (d *Drawing) Shapes() []*Shape {
	if d.drawingsPart == nil {
		return nil
	}

	drawing := d.drawingsPart.Drawing()
	if drawing == nil {
		return nil
	}

	oneCellAnchors := drawing.OneCellAnchors()
	twoCellAnchors := drawing.TwoCellAnchors()
	shapes := make(
		[]*Shape,
		0,
		len(oneCellAnchors)+len(twoCellAnchors),
	)

	// Parse one-cell anchors for shapes
	for _, oneCellAnchor := range oneCellAnchors {
		shape := d.parseOneCellAnchorShapes(
			oneCellAnchor,
		)
		if shape == nil {
			continue
		}
		shapes = append(shapes, shape)
	}

	// Parse two-cell anchors for shapes
	for _, twoCellAnchor := range twoCellAnchors {
		shape := d.parseTwoCellAnchorShapes(
			twoCellAnchor,
		)
		if shape == nil {
			continue
		}
		shapes = append(shapes, shape)
	}

	return shapes
}
