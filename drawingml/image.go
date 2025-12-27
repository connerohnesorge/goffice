//nolint:revive // file-length-limit: image utilities need to be together
package drawingml

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"  // Register GIF decoder
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"io"
)

// Image content type constants for common image formats.
const (
	// ContentTypeJPEG is the MIME type for JPEG images.
	ContentTypeJPEG = "image/jpeg"
	// ContentTypePNG is the MIME type for PNG images.
	ContentTypePNG = "image/png"
	// ContentTypeGIF is the MIME type for GIF images.
	ContentTypeGIF = "image/gif"
	// ContentTypeBMP is the MIME type for BMP images.
	ContentTypeBMP = "image/bmp"
	// ContentTypeWEBP is the MIME type for WebP images.
	ContentTypeWEBP = "image/webp"
	// ContentTypeTIFF is the MIME type for TIFF images.
	ContentTypeTIFF = "image/tiff"
	// ContentTypeSVG is the MIME type for SVG images.
	ContentTypeSVG = "image/svg+xml"
	// ContentTypeEMF is the MIME type for EMF (Enhanced Metafile) images.
	ContentTypeEMF = "image/x-emf"
	// ContentTypeWMF is the MIME type for WMF (Windows Metafile) images.
	ContentTypeWMF = "image/x-wmf"
)

// DefaultImageDPI is the default DPI (dots per inch) used for image
// calculations. This is the standard screen resolution on Windows.
const DefaultImageDPI = 96.0

// minImageHeaderSize is the minimum number of bytes needed to detect
// image format.
const minImageHeaderSize = 12

// Image format magic byte constants.
const (
	// jpegMagicByte1 is the first byte of JPEG files (FF).
	jpegMagicByte1 = 0xFF
	// jpegMagicByte2 is the second byte of JPEG files (D8).
	jpegMagicByte2 = 0xD8
	// bmpMagicByte1 is the first byte of BMP files (B).
	bmpMagicByte1 = 0x42
	// bmpMagicByte2 is the second byte of BMP files (M).
	bmpMagicByte2 = 0x4D
	// tiffLittleEndian1 is the first byte of little-endian TIFF.
	tiffLittleEndian1 = 0x49
	// tiffBigEndian1 is the first byte of big-endian TIFF.
	tiffBigEndian1 = 0x4D
	// tiffVersion is the TIFF version marker byte (0x2A).
	tiffVersion = 0x2A
	// emfHeaderMinSize is the minimum size needed to check EMF signature.
	emfHeaderMinSize = 44
	// emfHeaderType is the EMR_HEADER type value.
	emfHeaderType = 0x01
	// emfZeroByte is zero byte constant.
	emfZeroByte = 0x00
	// wmfPlaceableMagic1 is first byte of placeable WMF.
	wmfPlaceableMagic1 = 0xD7
	// wmfPlaceableMagic2 is second byte of placeable WMF.
	wmfPlaceableMagic2 = 0xCD
	// wmfPlaceableMagic3 is third byte of placeable WMF.
	wmfPlaceableMagic3 = 0xC6
	// wmfPlaceableMagic4 is fourth byte of placeable WMF.
	wmfPlaceableMagic4 = 0x9A
	// wmfStandardMagic2 is second byte of standard WMF.
	wmfStandardMagic2 = 0x09
	// imageHeaderReadSize is bytes to read for format detection.
	imageHeaderReadSize = 64
	// magicByte3Index is the index for the third magic byte.
	magicByte3Index = 3
)

// Common errors for image operations.
var (
	// ErrEmptyImageData is returned when image data is empty.
	ErrEmptyImageData = errors.New(
		"image data is empty",
	)
	// ErrUnknownImageFormat is returned when the image format cannot
	// be detected.
	ErrUnknownImageFormat = errors.New(
		"unknown image format",
	)
	// ErrInvalidImageDimensions is returned when image dimensions are invalid.
	ErrInvalidImageDimensions = errors.New(
		"invalid image dimensions",
	)
)

// GetImageDimensions returns the width and height of an image in pixels.
// It supports JPEG, PNG, and GIF formats (registered with the image package).
// For other formats, the data must be in a format supported by
// image.DecodeConfig.
func GetImageDimensions(
	data []byte,
) (width, height int, err error) {
	if len(data) == 0 {
		return 0, 0, ErrEmptyImageData
	}

	config, _, err := image.DecodeConfig(
		bytes.NewReader(data),
	)
	if err != nil {
		return 0, 0, err
	}

	return config.Width, config.Height, nil
}

// GetImageDimensionsFromReader returns the width and height of an image
// in pixels from an io.Reader. The reader is consumed during this operation.
func GetImageDimensionsFromReader(
	r io.Reader,
) (width, height int, err error) {
	if r == nil {
		return 0, 0, ErrEmptyImageData
	}

	config, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return config.Width, config.Height, nil
}

// DetectImageType detects the image format from the image data and returns
// the corresponding content type. It examines magic bytes at the start of
// the data to identify the format.
//
// Returns an empty string if the format cannot be detected.
//
//nolint:revive // cyclomatic: many image formats require separate checks
func DetectImageType(data []byte) string {
	if len(data) < minImageHeaderSize {
		return ""
	}

	// JPEG: FF D8 FF
	if data[0] == jpegMagicByte1 &&
		data[1] == jpegMagicByte2 &&
		data[2] == jpegMagicByte1 {
		return ContentTypeJPEG
	}

	// PNG: 89 50 4E 47 0D 0A 1A 0A
	if bytes.HasPrefix(
		data,
		[]byte{
			0x89,
			0x50,
			0x4E,
			0x47,
			0x0D,
			0x0A,
			0x1A,
			0x0A,
		},
	) {
		return ContentTypePNG
	}

	// GIF: 47 49 46 38 (GIF8)
	if bytes.HasPrefix(data, []byte("GIF8")) {
		return ContentTypeGIF
	}

	// BMP: 42 4D (BM)
	if data[0] == bmpMagicByte1 &&
		data[1] == bmpMagicByte2 {
		return ContentTypeBMP
	}

	// WEBP: RIFF....WEBP
	if bytes.HasPrefix(data, []byte("RIFF")) &&
		bytes.Equal(data[8:12], []byte("WEBP")) {
		return ContentTypeWEBP
	}

	// TIFF: 49 49 2A 00 (little endian) or 4D 4D 00 2A (big endian)
	if (data[0] == tiffLittleEndian1 && data[1] == tiffLittleEndian1 &&
		data[2] == tiffVersion && data[magicByte3Index] == emfZeroByte) ||
		(data[0] == tiffBigEndian1 && data[1] == tiffBigEndian1 &&
			data[2] == emfZeroByte && data[magicByte3Index] == tiffVersion) {
		return ContentTypeTIFF
	}

	// SVG: Check for XML declaration or SVG element
	// Note: This is a simple check and may not catch all SVG files
	if bytes.HasPrefix(data, []byte("<?xml")) ||
		bytes.HasPrefix(data, []byte("<svg")) {
		return ContentTypeSVG
	}

	// EMF: 01 00 00 00 (little endian header type)
	// EMF files start with a header record type of 1 (EMR_HEADER)
	if len(data) >= emfHeaderMinSize &&
		data[0] == emfHeaderType &&
		data[1] == emfZeroByte &&
		data[2] == emfZeroByte &&
		data[magicByte3Index] == emfZeroByte {
		// Additional check for EMF signature at offset 40
		if bytes.Equal(
			data[40:emfHeaderMinSize],
			[]byte{0x20, 0x45, 0x4D, 0x46},
		) {
			return ContentTypeEMF
		}
	}

	// WMF: D7 CD C6 9A (placeable WMF) or 01 00 09 00 (standard WMF)
	isPlaceableWMF := data[0] == wmfPlaceableMagic1 &&
		data[1] == wmfPlaceableMagic2 &&
		data[2] == wmfPlaceableMagic3 &&
		data[magicByte3Index] == wmfPlaceableMagic4
	isStandardWMF := data[0] == emfHeaderType &&
		data[1] == emfZeroByte &&
		data[2] == wmfStandardMagic2 &&
		data[magicByte3Index] == emfZeroByte
	if isPlaceableWMF || isStandardWMF {
		return ContentTypeWMF
	}

	return ""
}

// DetectImageTypeFromReader detects the image format from an io.Reader.
// It returns the content type, and a new reader that includes the bytes
// that were read for detection (since the original reader is consumed).
//
// Returns an empty string for content type if the format cannot be detected.
func DetectImageTypeFromReader(
	r io.Reader,
) (contentType string, newReader io.Reader, err error) {
	if r == nil {
		return "", nil, ErrEmptyImageData
	}

	// Read enough bytes for format detection
	header := make([]byte, imageHeaderReadSize)
	n, err := io.ReadFull(r, header)
	if err != nil && err != io.EOF &&
		err != io.ErrUnexpectedEOF {
		return "", nil, err
	}
	header = header[:n]

	if n < minImageHeaderSize {
		// Not enough data, but still return what we have
		return "", bytes.NewReader(header), nil
	}

	contentType = DetectImageType(header)

	// Create a new reader that combines the header with the rest of the data
	newReader = io.MultiReader(
		bytes.NewReader(header),
		r,
	)

	return contentType, newReader, nil
}

// ImageExtension returns the file extension (including the dot) for a given
// content type. Returns an empty string if the content type is not recognized.
func ImageExtension(contentType string) string {
	switch contentType {
	case ContentTypeJPEG:
		return ".jpg"
	case ContentTypePNG:
		return ".png"
	case ContentTypeGIF:
		return ".gif"
	case ContentTypeBMP:
		return ".bmp"
	case ContentTypeWEBP:
		return ".webp"
	case ContentTypeTIFF:
		return ".tiff"
	case ContentTypeSVG:
		return ".svg"
	case ContentTypeEMF:
		return ".emf"
	case ContentTypeWMF:
		return ".wmf"
	default:
		return ""
	}
}

// ImageExtentFromPixels converts pixel dimensions to EMU values at the
// specified DPI. It returns the width and height in EMUs.
func ImageExtentFromPixels(
	widthPx, heightPx int,
	dpi float64,
) (widthEMU, heightEMU EMU) {
	effectiveDPI := dpi
	if effectiveDPI <= 0 {
		effectiveDPI = DefaultImageDPI
	}

	// EMUs per pixel = EMUs per inch / DPI
	emuPerPixel := float64(
		EMUsPerInch,
	) / effectiveDPI
	widthEMU = EMU(float64(widthPx) * emuPerPixel)
	heightEMU = EMU(
		float64(heightPx) * emuPerPixel,
	)

	return widthEMU, heightEMU
}

// ImageExtentFromPixelsDefault converts pixel dimensions to EMU values using
// the default DPI (96).
func ImageExtentFromPixelsDefault(
	widthPx, heightPx int,
) (widthEMU, heightEMU EMU) {
	return ImageExtentFromPixels(
		widthPx,
		heightPx,
		DefaultImageDPI,
	)
}

// ImageExtentFromData extracts image dimensions from the image data and returns
// the dimensions in EMUs at the specified DPI.
func ImageExtentFromData(
	data []byte,
	dpi float64,
) (widthEMU, heightEMU EMU, err error) {
	width, height, err := GetImageDimensions(data)
	if err != nil {
		return 0, 0, err
	}

	widthEMU, heightEMU = ImageExtentFromPixels(
		width,
		height,
		dpi,
	)

	return widthEMU, heightEMU, nil
}

// ImageExtentFromDataDefault extracts image dimensions from the image data and
// returns the dimensions in EMUs using the default DPI (96).
func ImageExtentFromDataDefault(
	data []byte,
) (widthEMU, heightEMU EMU, err error) {
	return ImageExtentFromData(
		data,
		DefaultImageDPI,
	)
}

// FitImageToMaxSize scales image dimensions to fit within the specified maximum
// bounds while maintaining the aspect ratio. If the image already fits within
// the bounds, the original dimensions are returned.
func FitImageToMaxSize(
	widthEMU, heightEMU, maxWidthEMU, maxHeightEMU EMU,
) (newWidth, newHeight EMU) {
	if widthEMU <= 0 || heightEMU <= 0 {
		return widthEMU, heightEMU
	}

	// If image already fits, return original dimensions
	if widthEMU <= maxWidthEMU &&
		heightEMU <= maxHeightEMU {
		return widthEMU, heightEMU
	}

	// Calculate scale factors for both dimensions
	scaleX := float64(
		maxWidthEMU,
	) / float64(
		widthEMU,
	)
	scaleY := float64(
		maxHeightEMU,
	) / float64(
		heightEMU,
	)

	// Use the smaller scale factor to ensure the image fits within both bounds
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	newWidth = EMU(float64(widthEMU) * scale)
	newHeight = EMU(float64(heightEMU) * scale)

	return newWidth, newHeight
}

// FitImageToWidth scales image dimensions to the specified target width
// while maintaining the aspect ratio.
func FitImageToWidth(
	widthEMU, heightEMU, targetWidthEMU EMU,
) (newWidth, newHeight EMU) {
	if widthEMU <= 0 {
		return targetWidthEMU, heightEMU
	}

	scale := float64(
		targetWidthEMU,
	) / float64(
		widthEMU,
	)
	newWidth = targetWidthEMU
	newHeight = EMU(float64(heightEMU) * scale)

	return newWidth, newHeight
}

// FitImageToHeight scales image dimensions to the specified target height
// while maintaining the aspect ratio.
func FitImageToHeight(
	widthEMU, heightEMU, targetHeightEMU EMU,
) (newWidth, newHeight EMU) {
	if heightEMU <= 0 {
		return widthEMU, targetHeightEMU
	}

	scale := float64(
		targetHeightEMU,
	) / float64(
		heightEMU,
	)
	newWidth = EMU(float64(widthEMU) * scale)
	newHeight = targetHeightEMU

	return newWidth, newHeight
}

// ExtentFromImageData creates an Extent from image data using the default
// DPI. This is a convenience function that combines dimension detection
// and EMU conversion.
func ExtentFromImageData(
	data []byte,
) (Extent, error) {
	widthEMU, heightEMU, err := ImageExtentFromDataDefault(
		data,
	)
	if err != nil {
		return Extent{}, err
	}

	return NewExtent(widthEMU, heightEMU), nil
}

// ExtentFromImageDataAtDPI creates an Extent from image data at the
// specified DPI.
func ExtentFromImageDataAtDPI(
	data []byte,
	dpi float64,
) (Extent, error) {
	widthEMU, heightEMU, err := ImageExtentFromData(
		data,
		dpi,
	)
	if err != nil {
		return Extent{}, err
	}

	return NewExtent(widthEMU, heightEMU), nil
}
