package drawing

import "github.com/connerohnesorge/goffice-pdf/core"

// EmbeddedImage wraps an ImageXObject to satisfy core.PageImage and core.ImageResource.
type EmbeddedImage struct {
	xobj *ImageXObject
}

// LoadEmbeddedImage loads raw image data and returns an EmbeddedImage.
func LoadEmbeddedImage(
	data []byte,
	name string,
) (*EmbeddedImage, error) {
	xobj, err := LoadImage(data, name)
	if err != nil {
		return nil, err
	}

	return &EmbeddedImage{xobj: xobj}, nil
}

// Width returns the image width in pixels.
func (e *EmbeddedImage) Width() int {
	if e == nil || e.xobj == nil {
		return 0
	}

	return e.xobj.Width
}

// Height returns the image height in pixels.
func (e *EmbeddedImage) Height() int {
	if e == nil || e.xobj == nil {
		return 0
	}

	return e.xobj.Height
}

// Data returns the image data used for embedding.
func (e *EmbeddedImage) Data() []byte {
	if e == nil || e.xobj == nil {
		return nil
	}

	return e.xobj.Data
}

// ImageColorSpace returns the PDF color space name.
func (e *EmbeddedImage) ImageColorSpace() string {
	if e == nil || e.xobj == nil {
		return ""
	}

	return e.xobj.ColorSpace.String()
}

// ImageBitsPerComponent returns the number of bits per component.
func (e *EmbeddedImage) ImageBitsPerComponent() int {
	if e == nil || e.xobj == nil {
		return 0
	}

	return e.xobj.BitsPerComponent
}

// ImageFilter returns the PDF filter name.
func (e *EmbeddedImage) ImageFilter() string {
	if e == nil || e.xobj == nil {
		return ""
	}

	return e.xobj.Filter
}

// ImageDecode returns the decode array for the image.
func (e *EmbeddedImage) ImageDecode() []float64 {
	if e == nil || e.xobj == nil {
		return nil
	}

	return e.xobj.Decode
}

// ImageDecodeParms returns decode parameters for the filter.
func (e *EmbeddedImage) ImageDecodeParms() map[string]interface{} {
	if e == nil || e.xobj == nil {
		return nil
	}

	return e.xobj.DecodeParms
}

// ImageInterpolate reports whether interpolation is enabled.
func (e *EmbeddedImage) ImageInterpolate() bool {
	if e == nil || e.xobj == nil {
		return false
	}

	return e.xobj.Interpolate
}

// ImageSMask returns the soft mask image, if present.
func (e *EmbeddedImage) ImageSMask() core.ImageResource {
	if e == nil || e.xobj == nil || e.xobj.SMask == nil {
		return nil
	}

	return &EmbeddedImage{xobj: e.xobj.SMask}
}

var _ core.ImageResource = (*EmbeddedImage)(nil)
