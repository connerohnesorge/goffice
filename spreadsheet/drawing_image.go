// Package spreadsheet provides SpreadsheetML support for Excel documents.
// This file implements the Image type and image-related methods.
package spreadsheet

import (
	"fmt"

	"github.com/connerohnesorge/goffice/spreadsheet/elements"
	"github.com/connerohnesorge/goffice/spreadsheet/parts"
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
	toOffsetX    int64    // in EMUs (for two-cell anchors)
	toOffsetY    int64    // in EMUs (for two-cell anchors)
	width        int64    // in EMUs
	height       int64    // in EMUs
	// anchor holds either *elements.OneCellAnchor or
	// *elements.TwoCellAnchor
	anchor any
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

	// Add the anchor to the drawings part
	// Create a OneCellAnchor element with the picture
	drawing := drawingsPart.Drawing()
	if drawing == nil {
		return nil, fmt.Errorf(
			"drawing part has no root element",
		)
	}

	oneCellAnchor := drawing.AddOneCellAnchor()
	img.anchor = oneCellAnchor

	// Set up the from marker with cell position
	// Note: Markers use 0-based indexing, CellRef uses 1-based
	from := oneCellAnchor.GetOrCreateFrom()
	from.SetPosition(cell.Col-1, 0, cell.Row-1, 0)

	// Set up extent (size) - default to 0 for now, user must call SetSize
	ext := oneCellAnchor.GetOrCreateExt()
	ext.SetSize(0, 0)

	// Create the picture element
	pic := oneCellAnchor.GetOrCreatePicture()

	// Set picture properties with a unique ID
	// Use a simple counter-based ID for now
	picID := uint32(
		len(
			drawing.OneCellAnchors(),
		) + len(
			drawing.TwoCellAnchors(),
		),
	)
	pic.SetPictureId(picID, "")

	// Link the picture to the image part via relationship
	relID := imagePart.RelationshipID()
	pic.SetImageRelationship(relID)

	// Add client data
	oneCellAnchor.GetOrCreateClientData()

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

	// Add the anchor to the drawings part
	// Create a TwoCellAnchor element with the picture
	drawing := drawingsPart.Drawing()
	if drawing == nil {
		return nil, fmt.Errorf(
			"drawing part has no root element",
		)
	}

	twoCellAnchor := drawing.AddTwoCellAnchor()
	img.anchor = twoCellAnchor

	// Set up the from marker with cell position
	// Note: Markers use 0-based indexing, CellRef uses 1-based
	fromMarker := twoCellAnchor.GetOrCreateFrom()
	fromMarker.SetPosition(
		from.Col-1,
		0,
		from.Row-1,
		0,
	)

	// Set up the to marker with cell position
	toMarker := twoCellAnchor.GetOrCreateTo()
	toMarker.SetPosition(to.Col-1, 0, to.Row-1, 0)

	// Create the picture element
	pic := twoCellAnchor.GetOrCreatePicture()

	// Set picture properties with a unique ID
	picID := uint32(
		len(
			drawing.OneCellAnchors(),
		) + len(
			drawing.TwoCellAnchors(),
		),
	)
	pic.SetPictureId(picID, "")

	// Link the picture to the image part via relationship
	relID := imagePart.RelationshipID()
	pic.SetImageRelationship(relID)

	// Add client data
	twoCellAnchor.GetOrCreateClientData()

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

// convertTwoCellToOneCellAnchor converts a two-cell anchor to a one-cell
// anchor, preserving image relationship and size.
func (i *Image) convertTwoCellToOneCellAnchor() {
	twoCellAnchor, ok := i.anchor.(*elements.TwoCellAnchor)
	if !ok {
		return
	}

	drawing := i.drawingsPart.Drawing()
	if drawing == nil {
		return
	}

	// Create new one-cell anchor
	oneCellAnchor := drawing.AddOneCellAnchor()
	i.anchor = oneCellAnchor

	// Copy picture data from two-cell to one-cell
	i.copyPictureDataToOneCellAnchor(
		twoCellAnchor,
		oneCellAnchor,
	)

	// Set up extent with current width/height
	ext := oneCellAnchor.GetOrCreateExt()
	ext.SetSize(i.width, i.height)

	oneCellAnchor.GetOrCreateClientData()
}

// copyPictureDataToOneCellAnchor copies picture relationship from
// a two-cell anchor to a one-cell anchor.
func (*Image) copyPictureDataToOneCellAnchor(
	twoCellAnchor *elements.TwoCellAnchor,
	oneCellAnchor *elements.OneCellAnchor,
) {
	oldPic := twoCellAnchor.Picture()
	if oldPic == nil {
		return
	}

	newPic := oneCellAnchor.GetOrCreatePicture()

	blipFill := oldPic.BlipFill()
	if blipFill == nil {
		return
	}

	blip := blipFill.Blip()
	if blip == nil {
		return
	}

	newPic.SetImageRelationship(blip.Embed())
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

	// Update the underlying anchor element
	// If currently a two-cell anchor, we need to convert to one-cell
	i.convertTwoCellToOneCellAnchor()

	// Update the from marker
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		from := oneCellAnchor.GetOrCreateFrom()
		// Note: Markers use 0-based indexing, CellRef uses 1-based
		from.SetPosition(
			cell.Col-1,
			offsetX,
			cell.Row-1,
			offsetY,
		)
	}
}

// SetPositionBetween sets the position using a two-cell anchor.
//
//nolint:revive // argument-limit: all offsets are needed for two-cell anchor
func (i *Image) SetPositionBetween(
	from, to CellRef,
	fromOffsetX, fromOffsetY, toOffsetX, toOffsetY int64,
) {
	i.fromCell = from
	i.toCell = &to
	i.offsetX = fromOffsetX
	i.offsetY = fromOffsetY
	i.toOffsetX = toOffsetX
	i.toOffsetY = toOffsetY

	// Update the underlying anchor element
	// If currently a one-cell anchor, we need to convert to two-cell
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		// Create a new two-cell anchor
		drawing := i.drawingsPart.Drawing()
		if drawing != nil {
			twoCellAnchor := drawing.AddTwoCellAnchor()
			i.anchor = twoCellAnchor

			// Copy picture data from one-cell to two-cell
			if oldPic := oneCellAnchor.Picture(); oldPic != nil {
				newPic := twoCellAnchor.GetOrCreatePicture()
				// Copy image relationship
				if blipFill := oldPic.BlipFill(); blipFill != nil {
					if blip := blipFill.Blip(); blip != nil {
						newPic.SetImageRelationship(
							blip.Embed(),
						)
					}
				}
			}

			twoCellAnchor.GetOrCreateClientData()
		}
	}

	// Update the from and to markers
	if twoCellAnchor, ok := i.anchor.(*elements.TwoCellAnchor); ok {
		fromMarker := twoCellAnchor.GetOrCreateFrom()
		// Note: Markers use 0-based indexing, CellRef uses 1-based
		fromMarker.SetPosition(
			from.Col-1,
			fromOffsetX,
			from.Row-1,
			fromOffsetY,
		)

		toMarker := twoCellAnchor.GetOrCreateTo()
		toMarker.SetPosition(
			to.Col-1,
			toOffsetX,
			to.Row-1,
			toOffsetY,
		)
	}
}

// SetSize sets the image size in EMUs (English Metric Units).
// Note: 914400 EMUs = 1 inch
func (i *Image) SetSize(width, height int64) {
	i.width = width
	i.height = height

	// Update the underlying extent element (only for one-cell anchors)
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		ext := oneCellAnchor.GetOrCreateExt()
		ext.SetSize(width, height)
	}
}

// SetSizeInPixels sets the image size in pixels.
// Uses standard 96 DPI conversion.
func (i *Image) SetSizeInPixels(
	width, height int,
) {
	i.width = int64(width) * emusPerPixel
	i.height = int64(height) * emusPerPixel

	// Update the underlying extent element (only for one-cell anchors)
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		ext := oneCellAnchor.GetOrCreateExt()
		ext.SetSize(i.width, i.height)
	}
}

// SetSizeInInches sets the image size in inches.
func (i *Image) SetSizeInInches(
	width, height float64,
) {
	i.width = int64(width * emusPerInch)
	i.height = int64(height * emusPerInch)

	// Update the underlying extent element (only for one-cell anchors)
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		ext := oneCellAnchor.GetOrCreateExt()
		ext.SetSize(i.width, i.height)
	}
}

// SetSizeInCm sets the image size in centimeters.
func (i *Image) SetSizeInCm(
	width, height float64,
) {
	i.width = int64(width * emusPerCm)
	i.height = int64(height * emusPerCm)

	// Update the underlying extent element (only for one-cell anchors)
	if oneCellAnchor, ok := i.anchor.(*elements.OneCellAnchor); ok {
		ext := oneCellAnchor.GetOrCreateExt()
		ext.SetSize(i.width, i.height)
	}
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
