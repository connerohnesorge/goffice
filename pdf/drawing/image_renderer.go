package drawing

import (
	"errors"
	"fmt"

	"github.com/connerohnesorge/goffice-pdf/core"
	"github.com/connerohnesorge/goffice/drawingml"
)

const (
	// placeholderImageDimension is the default dimension for placeholder images.
	placeholderImageDimension = 100.0
)

// ImageRenderer handles rendering of images/pictures.
type ImageRenderer struct {
	ctx *core.RenderingContext
}

// NewImageRenderer creates a new image renderer.
func NewImageRenderer(
	ctx *core.RenderingContext,
) *ImageRenderer {
	return &ImageRenderer{ctx: ctx}
}

// RenderBounds encapsulates the position and size for rendering.
type RenderBounds struct {
	X, Y, Width, Height float64
}

// RenderPicture renders a picture/image.
func (r *ImageRenderer) RenderPicture(
	blipFill *drawingml.BlipFill,
	bounds RenderBounds,
) error {
	if blipFill == nil {
		return errors.New("blip fill is nil")
	}
	if r.ctx == nil || r.ctx.Page == nil {
		return errors.New("rendering context or page is nil")
	}

	// Get the embedded relationship ID
	embedID := blipFill.Embed()
	if embedID == "" {
		return errors.New(
			"no embedded image reference",
		)
	}

	if r.ctx.ImageResolver == nil {
		return errors.New(
			"image resolver is not set",
		)
	}

	data, err := r.ctx.ImageResolver(embedID)
	if err != nil {
		return fmt.Errorf(
			"resolve image %s: %w",
			embedID,
			err,
		)
	}

	embedded, err := LoadEmbeddedImage(
		data,
		embedID,
	)
	if err != nil {
		return fmt.Errorf(
			"load image %s: %w",
			embedID,
			err,
		)
	}

	r.ctx.Page.AddImage(
		embedded,
		bounds.X,
		bounds.Y,
		bounds.Width,
		bounds.Height,
	)

	return nil
}

// CropValues encapsulates the crop percentages for an image.
type CropValues struct {
	Left, Top, Right, Bottom float64
}

// RenderCroppedPicture renders an image with cropping applied.
func (r *ImageRenderer) RenderCroppedPicture(
	blipFill *drawingml.BlipFill,
	bounds RenderBounds,
	crop CropValues,
) error {
	// Cropping is specified as percentages (0-1) of the image dimensions
	// Calculate the visible portion

	visibleWidth := bounds.Width * (1 - crop.Left - crop.Right)
	visibleHeight := bounds.Height * (1 - crop.Top - crop.Bottom)
	visibleX := bounds.X + bounds.Width*crop.Left
	visibleY := bounds.Y + bounds.Height*crop.Top

	// Clip to the visible area
	r.ctx.Page.SaveGraphicsState()
	defer r.ctx.Page.RestoreGraphicsState()

	// Set clipping rectangle
	r.ctx.Page.DrawRectangle(
		visibleX,
		visibleY,
		visibleWidth,
		visibleHeight,
		false,
		false,
	)
	// TODO: Implement clip operation when available in Page interface

	// Render the full image, but only the clipped portion will be visible
	return r.RenderPicture(blipFill, bounds)
}

// RenderStretchedPicture renders an image with stretch fill.
func (r *ImageRenderer) RenderStretchedPicture(
	blipFill *drawingml.BlipFill,
	bounds RenderBounds,
) error {
	// Stretch mode fills the entire shape with the image
	// This is the default behavior
	return r.RenderPicture(blipFill, bounds)
}

// TileOptions encapsulates tiling options for an image.
type TileOptions struct {
	Alignment drawingml.RectAlignValue
	Flip      drawingml.TileFlipValue
}

// RenderTiledPicture renders an image with tile fill.
func (r *ImageRenderer) RenderTiledPicture(
	blipFill *drawingml.BlipFill,
	bounds RenderBounds,
	_ TileOptions, // TODO: Implement tiling pattern support
) error {
	// Tile mode repeats the image to fill the shape
	// This would require creating a tiling pattern in PDF

	// For basic implementation, just render once
	return r.RenderPicture(blipFill, bounds)
}

// GetImageDimensions returns the dimensions of an embedded image.
func (*ImageRenderer) GetImageDimensions(
	_ string, // TODO: Load image from relationship when implemented
) (width, height float64, err error) {
	// This would load the image and get its dimensions
	// For now, return placeholder dimensions
	return placeholderImageDimension, placeholderImageDimension, nil
}

// ImageDimensions encapsulates image width and height.
type ImageDimensions struct {
	Width, Height float64
}

// ScaleImageToFit calculates the scale factor to fit an image in bounds.
//
//nolint:revive // preserveAspectRatio is a valid configuration parameter
func (*ImageRenderer) ScaleImageToFit(
	imageDims, boundsDims ImageDimensions,
	preserveAspectRatio bool,
) (scaleX, scaleY float64) {
	if !preserveAspectRatio {
		// Stretch to fit exactly
		return boundsDims.Width / imageDims.Width, boundsDims.Height / imageDims.Height
	}

	// Preserve aspect ratio - fit within bounds
	scaleX = boundsDims.Width / imageDims.Width
	scaleY = boundsDims.Height / imageDims.Height

	// Use the smaller scale to ensure it fits
	if scaleX < scaleY {
		return scaleX, scaleX
	}

	return scaleY, scaleY
}
