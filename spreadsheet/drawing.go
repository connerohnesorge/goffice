// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the high-level Drawing/Image API.
package spreadsheet

import (
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

// Image represents a high-level wrapper around an image in a worksheet.
type Image struct {
	sheet        *Sheet
	drawingsPart *parts.DrawingsPart
	imagePart    *parts.ImagePart
	fromCell     CellRef
	toCell       *CellRef // nil for one-cell anchor
	offsetX      int64    // in EMUs
	offsetY      int64    // in EMUs
	width        int64    // in EMUs
	height       int64    // in EMUs
}

// newImage creates a new Image at the specified cell (one-cell anchor).
//
//nolint:revive // argument-limit: all parameters are needed for image creation
func newImage(
	sheet *Sheet,
	drawingsPart *parts.DrawingsPart,
	cell CellRef,
	data []byte,
	_ string, // contentType - detected from data instead
) (*Image, error) {
	// Detect image type from data
	imageType := parts.DetectImageType(data, "")

	// Add an image part
	imagePart, err := drawingsPart.AddImagePart(
		imageType,
	)
	if err != nil {
		return nil, err
	}

	// Set the image data
	imagePart.FeedDataBytes(data)

	img := &Image{
		sheet:        sheet,
		drawingsPart: drawingsPart,
		imagePart:    imagePart,
		fromCell:     cell,
	}

	// TODO: Add the anchor to the drawings part
	// This would create a OneCellAnchor element with the picture

	return img, nil
}

// newImageBetween creates a new Image anchored between two cells
// (two-cell anchor).
//
//nolint:revive // argument-limit: all parameters are needed for image creation
func newImageBetween(
	sheet *Sheet,
	drawingsPart *parts.DrawingsPart,
	from, to CellRef,
	data []byte,
	_ string, // contentType - detected from data instead
) (*Image, error) {
	// Detect image type from data
	imageType := parts.DetectImageType(data, "")

	// Add an image part
	imagePart, err := drawingsPart.AddImagePart(
		imageType,
	)
	if err != nil {
		return nil, err
	}

	// Set the image data
	imagePart.FeedDataBytes(data)

	img := &Image{
		sheet:        sheet,
		drawingsPart: drawingsPart,
		imagePart:    imagePart,
		fromCell:     from,
		toCell:       &to,
	}

	// TODO: Add the anchor to the drawings part
	// This would create a TwoCellAnchor element with the picture

	return img, nil
}

// FromCell returns the anchor cell (top-left corner).
func (i *Image) FromCell() CellRef {
	return i.fromCell
}

// ToCell returns the end anchor cell (bottom-right corner) for two-cell
// anchors. Returns nil for one-cell anchors.
func (i *Image) ToCell() *CellRef {
	return i.toCell
}

// SetPosition sets the position using a one-cell anchor.
func (i *Image) SetPosition(
	cell CellRef,
	offsetX, offsetY int64,
) {
	i.fromCell = cell
	i.offsetX = offsetX
	i.offsetY = offsetY
	i.toCell = nil
	// TODO: Update the underlying anchor element
}

// SetPositionBetween sets the position using a two-cell anchor.
//
//nolint:revive // argument-limit: all offsets are needed for two-cell anchor
func (i *Image) SetPositionBetween(
	from, to CellRef,
	fromOffsetX, fromOffsetY, _ /* toOffsetX */, _ /* toOffsetY */ int64,
) {
	i.fromCell = from
	i.toCell = &to
	i.offsetX = fromOffsetX
	i.offsetY = fromOffsetY
	// TODO: Store to offsets and update the underlying anchor element
}

// SetSize sets the image size in EMUs (English Metric Units).
// Note: 914400 EMUs = 1 inch
func (i *Image) SetSize(width, height int64) {
	i.width = width
	i.height = height
	// TODO: Update the underlying extent element
}

// SetSizeInPixels sets the image size in pixels.
// Uses standard 96 DPI conversion.
func (i *Image) SetSizeInPixels(
	width, height int,
) {
	i.width = int64(width) * emusPerPixel
	i.height = int64(height) * emusPerPixel
	// TODO: Update the underlying extent element
}

// SetSizeInInches sets the image size in inches.
func (i *Image) SetSizeInInches(
	width, height float64,
) {
	i.width = int64(width * emusPerInch)
	i.height = int64(height * emusPerInch)
	// TODO: Update the underlying extent element
}

// SetSizeInCm sets the image size in centimeters.
func (i *Image) SetSizeInCm(
	width, height float64,
) {
	i.width = int64(width * emusPerCm)
	i.height = int64(height * emusPerCm)
	// TODO: Update the underlying extent element
}

// Width returns the image width in EMUs.
func (i *Image) Width() int64 {
	return i.width
}

// Height returns the image height in EMUs.
func (i *Image) Height() int64 {
	return i.height
}

// ContentType returns the image content type.
func (i *Image) ContentType() string {
	return i.imagePart.ContentType()
}

// Data returns the image data.
func (i *Image) Data() []byte {
	data, _ := i.imagePart.GetImageData()

	return data
}

// Drawing represents a collection of drawing objects in a worksheet.
type Drawing struct {
	// NOTE: Future implementation will store drawing elements
	// from the drawings part
}

// Images returns all images in the drawing.
func (*Drawing) Images() []*Image {
	// TODO: Parse the drawings part and return Image wrappers
	return nil
}

// Shapes returns all shapes in the drawing.
func (*Drawing) Shapes() []*Shape {
	// TODO: Parse the drawings part and return Shape wrappers
	return nil
}

// Shape represents a shape object in a worksheet.
type Shape struct {
	shapeType string
	text      string
	fromCell  CellRef
	toCell    *CellRef
}

// Type returns the shape type.
func (s *Shape) Type() string {
	return s.shapeType
}

// Text returns the shape text.
func (s *Shape) Text() string {
	return s.text
}

// SetText sets the shape text.
func (s *Shape) SetText(text string) {
	s.text = text
	// TODO: Update the underlying shape element
}

// FromCell returns the anchor cell (top-left corner).
func (s *Shape) FromCell() CellRef {
	return s.fromCell
}

// ToCell returns the end anchor cell (bottom-right corner).
func (s *Shape) ToCell() *CellRef {
	return s.toCell
}
