// This file implements image handling with support for JPEG passthrough,
// PNG handling with alpha channel support, and other image formats.

package drawing

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

// ImageFormat represents the format of an image.
type ImageFormat int

const (
	// ImageFormatUnknown represents an unknown image format.
	ImageFormatUnknown ImageFormat = iota
	// ImageFormatJPEG represents JPEG format (DCT encoding).
	ImageFormatJPEG
	// ImageFormatPNG represents PNG format.
	ImageFormatPNG
	// ImageFormatGIF represents GIF format.
	ImageFormatGIF
	// ImageFormatBMP represents BMP format.
	ImageFormatBMP
)

// String returns the string representation of the image format.
func (f ImageFormat) String() string {
	switch f {
	case ImageFormatJPEG:
		return "JPEG"
	case ImageFormatPNG:
		return "PNG"
	case ImageFormatGIF:
		return "GIF"
	case ImageFormatBMP:
		return "BMP"
	case ImageFormatUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

// ColorSpace represents the PDF color space for images.
type ColorSpace int

const (
	// ColorSpaceDeviceRGB is the RGB color space.
	ColorSpaceDeviceRGB ColorSpace = iota
	// ColorSpaceDeviceGray is the grayscale color space.
	ColorSpaceDeviceGray
	// ColorSpaceDeviceCMYK is the CMYK color space.
	ColorSpaceDeviceCMYK
	// ColorSpaceIndexed is an indexed color space (palette-based).
	ColorSpaceIndexed
)

// String returns the PDF name for the color space.
func (cs ColorSpace) String() string {
	switch cs {
	case ColorSpaceDeviceRGB:
		return "DeviceRGB"
	case ColorSpaceDeviceGray:
		return "DeviceGray"
	case ColorSpaceDeviceCMYK:
		return "DeviceCMYK"
	case ColorSpaceIndexed:
		return "Indexed"
	default:
		return "DeviceRGB"
	}
}

// FitMode specifies how an image should be fitted within bounds.
type FitMode int

const (
	// FitModeContain scales the image to fit entirely within the bounds,
	// preserving aspect ratio. There may be empty space.
	FitModeContain FitMode = iota
	// FitModeCover scales the image to cover the entire bounds,
	// preserving aspect ratio. Parts may be clipped.
	FitModeCover
	// FitModeStretch stretches the image to fill the bounds exactly,
	// ignoring aspect ratio.
	FitModeStretch
	// FitModeNone uses the original image dimensions.
	FitModeNone
)

// Errors for image operations.
var (
	ErrUnsupportedImageFormat = errors.New(
		"unsupported image format",
	)
	ErrInvalidImageData = errors.New(
		"invalid image data",
	)
	ErrImageTooLarge = errors.New(
		"image too large",
	)
	ErrNoImageData = errors.New(
		"no image data",
	)
)

// ImageXObject represents a PDF Image XObject.
// This contains all the data needed to embed an image in a PDF.
type ImageXObject struct {
	// Name is the resource name (e.g., "Im1", "Im2").
	Name string
	// Width is the image width in pixels.
	Width int
	// Height is the image height in pixels.
	Height int
	// BitsPerComponent is the number of bits per color component (typically 8).
	BitsPerComponent int
	// ColorSpace is the image color space.
	ColorSpace ColorSpace
	// ColorComponents is the number of color components.
	ColorComponents int
	// Data is the image stream data (may be compressed).
	Data []byte
	// Filter is the compression filter applied (e.g., "DCTDecode", "FlateDecode").
	Filter string
	// DecodeParms contains decode parameters for the filter.
	DecodeParms map[string]interface{}
	// SMask is the soft mask (alpha channel) for transparency.
	SMask *ImageXObject
	// Decode is the decode array for color mapping.
	Decode []float64
	// Intent is the rendering intent.
	Intent string
	// Interpolate specifies whether to smooth the image when scaled.
	Interpolate bool
	// OriginalFormat is the original format of the source image.
	OriginalFormat ImageFormat
}

// NewImageXObject creates a new ImageXObject with default values.
func NewImageXObject(name string) *ImageXObject {
	return &ImageXObject{
		Name:             name,
		BitsPerComponent: 8,
		ColorSpace:       ColorSpaceDeviceRGB,
		ColorComponents:  3,
		Interpolate:      true,
	}
}

// AspectRatio returns the aspect ratio (width/height) of the image.
func (img *ImageXObject) AspectRatio() float64 {
	if img.Height == 0 {
		return 0
	}

	return float64(
		img.Width,
	) / float64(
		img.Height,
	)
}

// HasAlpha returns true if the image has transparency.
func (img *ImageXObject) HasAlpha() bool {
	return img.SMask != nil
}

// Dictionary returns the PDF dictionary entries for this image XObject.
func (img *ImageXObject) Dictionary() map[string]interface{} {
	dict := map[string]interface{}{
		"Type":             "/XObject",
		"Subtype":          "/Image",
		"Width":            img.Width,
		"Height":           img.Height,
		"BitsPerComponent": img.BitsPerComponent,
		"Interpolate":      img.Interpolate,
	}

	// Color space
	if img.ColorSpace == ColorSpaceIndexed {
		// Indexed color space requires special handling
		dict["ColorSpace"] = "/" + img.ColorSpace.String()
	} else {
		dict["ColorSpace"] = "/" + img.ColorSpace.String()
	}

	// Filter
	if img.Filter != "" {
		dict["Filter"] = "/" + img.Filter
	}

	// Decode parameters
	if len(img.DecodeParms) > 0 {
		dict["DecodeParms"] = img.DecodeParms
	}

	// Decode array
	if len(img.Decode) > 0 {
		dict["Decode"] = img.Decode
	}

	// Rendering intent
	if img.Intent != "" {
		dict["Intent"] = "/" + img.Intent
	}

	return dict
}

// ImageManager tracks embedded images and generates unique names.
type ImageManager struct {
	images    map[string]*ImageXObject
	dataCache map[string]string // Maps data hash to image name for deduplication
	counter   int
	prefix    string
}

// NewImageManager creates a new image manager.
func NewImageManager() *ImageManager {
	return &ImageManager{
		images:    make(map[string]*ImageXObject),
		dataCache: make(map[string]string),
		counter:   0,
		prefix:    "Im",
	}
}

// SetPrefix sets the prefix for image resource names.
func (im *ImageManager) SetPrefix(prefix string) {
	im.prefix = prefix
}

// NextName generates the next unique image name.
func (im *ImageManager) NextName() string {
	im.counter++

	return fmt.Sprintf(
		"%s%d",
		im.prefix,
		im.counter,
	)
}

// Add adds an image to the manager.
func (im *ImageManager) Add(img *ImageXObject) {
	if img.Name == "" {
		img.Name = im.NextName()
	}
	im.images[img.Name] = img
}

// Get retrieves an image by name.
func (im *ImageManager) Get(
	name string,
) (*ImageXObject, bool) {
	img, ok := im.images[name]

	return img, ok
}

// All returns all managed images.
func (im *ImageManager) All() map[string]*ImageXObject {
	return im.images
}

// Count returns the number of managed images.
func (im *ImageManager) Count() int {
	return len(im.images)
}

// DetectFormat detects the image format from raw data.
func DetectFormat(data []byte) ImageFormat {
	if len(data) < 8 {
		return ImageFormatUnknown
	}

	// JPEG: starts with FF D8 FF
	if data[0] == 0xFF && data[1] == 0xD8 &&
		data[2] == 0xFF {
		return ImageFormatJPEG
	}

	// PNG: starts with 89 50 4E 47 0D 0A 1A 0A
	if data[0] == 0x89 && data[1] == 0x50 &&
		data[2] == 0x4E &&
		data[3] == 0x47 &&
		data[4] == 0x0D &&
		data[5] == 0x0A &&
		data[6] == 0x1A &&
		data[7] == 0x0A {
		return ImageFormatPNG
	}

	// GIF: starts with "GIF87a" or "GIF89a"
	if data[0] == 0x47 && data[1] == 0x49 &&
		data[2] == 0x46 {
		if (data[3] == 0x38 && data[4] == 0x37 && data[5] == 0x61) ||
			(data[3] == 0x38 && data[4] == 0x39 && data[5] == 0x61) {
			return ImageFormatGIF
		}
	}

	// BMP: starts with "BM"
	if data[0] == 0x42 && data[1] == 0x4D {
		return ImageFormatBMP
	}

	return ImageFormatUnknown
}

// LoadImage loads an image from raw data and creates an ImageXObject.
func LoadImage(
	data []byte,
	name string,
) (*ImageXObject, error) {
	if len(data) == 0 {
		return nil, ErrNoImageData
	}

	format := DetectFormat(data)

	switch format {
	case ImageFormatJPEG:
		return LoadJPEG(data, name)
	case ImageFormatPNG:
		return LoadPNG(data, name)
	case ImageFormatGIF:
		return LoadGIF(data, name)
	case ImageFormatBMP:
		return LoadBMP(data, name)
	case ImageFormatUnknown:
		// Try to decode as a Go image
		return LoadGoImage(data, name)
	default:
		// Try to decode as a Go image
		return LoadGoImage(data, name)
	}
}

// LoadJPEG loads a JPEG image with passthrough (no re-encoding).
// JPEG data is passed directly to PDF using DCTDecode filter.
func LoadJPEG(
	data []byte,
	name string,
) (*ImageXObject, error) {
	// Decode just to get dimensions and color space
	reader := bytes.NewReader(data)
	config, err := jpeg.DecodeConfig(reader)
	if err != nil {
		return nil, fmt.Errorf(
			"decode JPEG config: %w",
			err,
		)
	}

	img := NewImageXObject(name)
	img.Width = config.Width
	img.Height = config.Height
	img.Data = data
	img.Filter = "DCTDecode"
	img.OriginalFormat = ImageFormatJPEG

	// Determine color space from the image
	switch config.ColorModel {
	case color.GrayModel, color.Gray16Model:
		img.ColorSpace = ColorSpaceDeviceGray
		img.ColorComponents = 1
	case color.YCbCrModel:
		img.ColorSpace = ColorSpaceDeviceRGB
		img.ColorComponents = 3
	case color.CMYKModel:
		img.ColorSpace = ColorSpaceDeviceCMYK
		img.ColorComponents = 4
		// CMYK JPEGs need inverted decode array
		img.Decode = []float64{
			1,
			0,
			1,
			0,
			1,
			0,
			1,
			0,
		}
	default:
		img.ColorSpace = ColorSpaceDeviceRGB
		img.ColorComponents = 3
	}

	return img, nil
}

// LoadPNG loads a PNG image, handling alpha channel transparency.
func LoadPNG(
	data []byte,
	name string,
) (*ImageXObject, error) {
	reader := bytes.NewReader(data)
	goImg, err := png.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf(
			"decode PNG: %w",
			err,
		)
	}

	return encodeGoImage(
		goImg,
		name,
		ImageFormatPNG,
	)
}

// LoadGIF loads a GIF image, converting to RGB.
func LoadGIF(
	data []byte,
	name string,
) (*ImageXObject, error) {
	reader := bytes.NewReader(data)
	goImg, err := gif.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf(
			"decode GIF: %w",
			err,
		)
	}

	return encodeGoImage(
		goImg,
		name,
		ImageFormatGIF,
	)
}

// LoadBMP loads a BMP image.
func LoadBMP(
	data []byte,
	name string,
) (*ImageXObject, error) {
	// Go's standard library doesn't include BMP decoder by default
	// Try to decode using image.Decode which may work if a BMP decoder is registered
	reader := bytes.NewReader(data)
	goImg, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf(
			"decode BMP: %w",
			err,
		)
	}

	return encodeGoImage(
		goImg,
		name,
		ImageFormatBMP,
	)
}

// LoadGoImage attempts to load any image format supported by Go's image package.
func LoadGoImage(
	data []byte,
	name string,
) (*ImageXObject, error) {
	reader := bytes.NewReader(data)
	goImg, format, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf(
			"decode image: %w",
			err,
		)
	}

	var imgFormat ImageFormat
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		imgFormat = ImageFormatJPEG
	case "png":
		imgFormat = ImageFormatPNG
	case "gif":
		imgFormat = ImageFormatGIF
	case "bmp":
		imgFormat = ImageFormatBMP
	default:
		imgFormat = ImageFormatUnknown
	}

	return encodeGoImage(goImg, name, imgFormat)
}

// LoadFromGoImage creates an ImageXObject from a Go image.Image.
func LoadFromGoImage(
	goImg image.Image,
	name string,
) (*ImageXObject, error) {
	return encodeGoImage(
		goImg,
		name,
		ImageFormatUnknown,
	)
}

// encodeGoImage converts a Go image to an ImageXObject.
func encodeGoImage(
	goImg image.Image,
	name string,
	format ImageFormat,
) (*ImageXObject, error) {
	bounds := goImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width == 0 || height == 0 {
		return nil, ErrInvalidImageData
	}

	img := NewImageXObject(name)
	img.Width = width
	img.Height = height
	img.OriginalFormat = format

	// Check if image has alpha channel
	hasAlpha := imageHasAlpha(goImg)

	// Extract RGB and alpha data
	rgbData, alphaData := extractImageData(
		goImg,
		hasAlpha,
	)

	// Compress RGB data with zlib
	var compressedRGB bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedRGB)
	_, err := zlibWriter.Write(rgbData)
	if err != nil {
		return nil, fmt.Errorf(
			"compress RGB data: %w",
			err,
		)
	}
	if err := zlibWriter.Close(); err != nil {
		return nil, fmt.Errorf(
			"close zlib writer: %w",
			err,
		)
	}

	img.Data = compressedRGB.Bytes()
	img.Filter = "FlateDecode"
	img.ColorSpace = ColorSpaceDeviceRGB
	img.ColorComponents = 3

	// Create soft mask for alpha channel
	if hasAlpha && len(alphaData) > 0 {
		smask := NewImageXObject(name + "_smask")
		smask.Width = width
		smask.Height = height
		smask.ColorSpace = ColorSpaceDeviceGray
		smask.ColorComponents = 1

		// Compress alpha data
		var compressedAlpha bytes.Buffer
		alphaWriter := zlib.NewWriter(
			&compressedAlpha,
		)
		_, err := alphaWriter.Write(alphaData)
		if err != nil {
			return nil, fmt.Errorf(
				"compress alpha data: %w",
				err,
			)
		}
		if err := alphaWriter.Close(); err != nil {
			return nil, fmt.Errorf(
				"close alpha zlib writer: %w",
				err,
			)
		}

		smask.Data = compressedAlpha.Bytes()
		smask.Filter = "FlateDecode"
		img.SMask = smask
	}

	return img, nil
}

// imageHasAlpha checks if an image has an alpha channel.
func imageHasAlpha(img image.Image) bool {
	switch img := img.(type) {
	case *image.RGBA, *image.RGBA64, *image.NRGBA, *image.NRGBA64:
		return true
	case *image.Paletted:
		// Check if palette has any transparent colors
		p := img
		for _, c := range p.Palette {
			_, _, _, a := c.RGBA()
			if a < 0xFFFF {
				return true
			}
		}

		return false
	default:
		// Check a sample of pixels
		bounds := img.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y && y < bounds.Min.Y+10; y++ {
			for x := bounds.Min.X; x < bounds.Max.X && x < bounds.Min.X+10; x++ {
				_, _, _, a := img.At(x, y).RGBA()
				if a < 0xFFFF {
					return true
				}
			}
		}

		return false
	}
}

// extractImageData extracts RGB and alpha data from an image.
func extractImageData(
	img image.Image,
	extractAlpha bool,
) (rgb, alpha []byte) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	rgb = make([]byte, width*height*3)
	if extractAlpha {
		alpha = make([]byte, width*height)
	}

	i := 0
	j := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()

			// Convert from 16-bit to 8-bit
			rgb[i] = uint8(r >> 8)
			rgb[i+1] = uint8(g >> 8)
			rgb[i+2] = uint8(b >> 8)
			i += 3

			if extractAlpha {
				alpha[j] = uint8(a >> 8)
				j++
			}
		}
	}

	return rgb, alpha
}

// DrawImageOperator returns the PDF operator to draw an image.
// The image is drawn at (x, y) with the specified width and height.
// The transform matrix places the image correctly in user space.
func DrawImageOperator(
	name string,
	x, y, width, height float64,
) string {
	var sb strings.Builder
	sb.WriteString("q\n")
	sb.WriteString(
		fmt.Sprintf(
			"%s 0 0 %s %s %s cm\n",
			formatFloat(
				width,
			),
			formatFloat(height),
			formatFloat(x),
			formatFloat(y),
		),
	)
	sb.WriteString(fmt.Sprintf("/%s Do\n", name))
	sb.WriteString("Q\n")

	return sb.String()
}

// DrawImage returns the PDF operators to draw an image at (x, y) with specified dimensions.
func DrawImage(
	x, y, width, height float64,
	imageName string,
) string {
	return DrawImageOperator(
		imageName,
		x,
		y,
		width,
		height,
	)
}

// DrawImageAt returns the PDF operators to draw an image at (x, y) at its original size.
func DrawImageAt(
	x, y float64,
	img *ImageXObject,
) string {
	return DrawImageOperator(
		img.Name,
		x,
		y,
		float64(img.Width),
		float64(img.Height),
	)
}

// DrawImageFit returns the PDF operators to draw an image fitted within a rectangle.
func DrawImageFit(
	rect Rect,
	img *ImageXObject,
	mode FitMode,
) string {
	if img == nil {
		return ""
	}

	imgWidth := float64(img.Width)
	imgHeight := float64(img.Height)

	if imgWidth == 0 || imgHeight == 0 {
		return ""
	}

	var drawWidth, drawHeight, drawX, drawY float64

	switch mode {
	case FitModeContain:
		// Scale to fit entirely within bounds
		imgAspect := imgWidth / imgHeight
		rectAspect := rect.Width / rect.Height

		if imgAspect > rectAspect {
			// Image is wider than rect, fit to width
			drawWidth = rect.Width
			drawHeight = rect.Width / imgAspect
		} else {
			// Image is taller than rect, fit to height
			drawHeight = rect.Height
			drawWidth = rect.Height * imgAspect
		}
		// Center within bounds
		drawX = rect.X + (rect.Width-drawWidth)/2
		drawY = rect.Y + (rect.Height-drawHeight)/2

	case FitModeCover:
		// Scale to cover entire bounds
		imgAspect := imgWidth / imgHeight
		rectAspect := rect.Width / rect.Height

		if imgAspect > rectAspect {
			// Image is wider, fit to height
			drawHeight = rect.Height
			drawWidth = rect.Height * imgAspect
		} else {
			// Image is taller, fit to width
			drawWidth = rect.Width
			drawHeight = rect.Width / imgAspect
		}
		// Center (may be clipped)
		drawX = rect.X + (rect.Width-drawWidth)/2
		drawY = rect.Y + (rect.Height-drawHeight)/2

	case FitModeStretch:
		// Stretch to fill exactly
		drawWidth = rect.Width
		drawHeight = rect.Height
		drawX = rect.X
		drawY = rect.Y

	case FitModeNone:
		// Use original size
		drawWidth = imgWidth
		drawHeight = imgHeight
		drawX = rect.X
		drawY = rect.Y

	default:
		drawWidth = rect.Width
		drawHeight = rect.Height
		drawX = rect.X
		drawY = rect.Y
	}

	return DrawImageOperator(
		img.Name,
		drawX,
		drawY,
		drawWidth,
		drawHeight,
	)
}

// ImageBounds calculates the bounds for drawing an image with a specific fit mode.
func ImageBounds(
	rect Rect,
	img *ImageXObject,
	mode FitMode,
) Rect {
	if img == nil || img.Width == 0 ||
		img.Height == 0 {
		return rect
	}

	imgWidth := float64(img.Width)
	imgHeight := float64(img.Height)

	switch mode {
	case FitModeContain:
		imgAspect := imgWidth / imgHeight
		rectAspect := rect.Width / rect.Height

		var drawWidth, drawHeight float64
		if imgAspect > rectAspect {
			drawWidth = rect.Width
			drawHeight = rect.Width / imgAspect
		} else {
			drawHeight = rect.Height
			drawWidth = rect.Height * imgAspect
		}

		return Rect{
			X:      rect.X + (rect.Width-drawWidth)/2,
			Y:      rect.Y + (rect.Height-drawHeight)/2,
			Width:  drawWidth,
			Height: drawHeight,
		}

	case FitModeCover:
		imgAspect := imgWidth / imgHeight
		rectAspect := rect.Width / rect.Height

		var drawWidth, drawHeight float64
		if imgAspect > rectAspect {
			drawHeight = rect.Height
			drawWidth = rect.Height * imgAspect
		} else {
			drawWidth = rect.Width
			drawHeight = rect.Width / imgAspect
		}

		return Rect{
			X:      rect.X + (rect.Width-drawWidth)/2,
			Y:      rect.Y + (rect.Height-drawHeight)/2,
			Width:  drawWidth,
			Height: drawHeight,
		}

	case FitModeStretch:
		return rect

	case FitModeNone:
		return Rect{
			X:      rect.X,
			Y:      rect.Y,
			Width:  imgWidth,
			Height: imgHeight,
		}

	default:
		return rect
	}
}

// TransparencyGroup represents a PDF transparency group for alpha compositing.
type TransparencyGroup struct {
	// ColorSpace is the blending color space.
	ColorSpace ColorSpace
	// Isolated specifies whether the group is isolated from the backdrop.
	Isolated bool
	// Knockout specifies whether elements knock out elements beneath them.
	Knockout bool
}

// DefaultTransparencyGroup returns a default transparency group for images with alpha.
func DefaultTransparencyGroup() *TransparencyGroup {
	return &TransparencyGroup{
		ColorSpace: ColorSpaceDeviceRGB,
		Isolated:   false,
		Knockout:   false,
	}
}

// Dictionary returns the PDF dictionary entries for this transparency group.
func (tg *TransparencyGroup) Dictionary() map[string]interface{} {
	return map[string]interface{}{
		"Type": "/Group",
		"S":    "/Transparency",
		"CS":   "/" + tg.ColorSpace.String(),
		"I":    tg.Isolated,
		"K":    tg.Knockout,
	}
}

// BeginTransparencyGroup returns PDF operators to begin a transparency group.
func BeginTransparencyGroup() string {
	return "q\n"
}

// EndTransparencyGroup returns PDF operators to end a transparency group.
func EndTransparencyGroup() string {
	return "Q\n"
}

// SetImageAlpha returns PDF operators to set the alpha (opacity) for subsequent images.
// Alpha is a value from 0 (transparent) to 1 (opaque).
func SetImageAlpha(alpha float64) string {
	if alpha >= 1.0 {
		return ""
	}
	alpha = clamp01(alpha)
	// This requires a graphics state with the ca (fill alpha) set
	return fmt.Sprintf(
		"/GS_A%d gs\n",
		int(alpha*100),
	)
}

// ImageAlphaGraphicsState returns the graphics state dictionary for the given alpha.
func ImageAlphaGraphicsState(
	alpha float64,
) map[string]interface{} {
	return map[string]interface{}{
		"Type": "/ExtGState",
		"ca":   clamp01(alpha),
		"CA":   clamp01(alpha),
	}
}

// Inline image support

// InlineImage represents a small image embedded directly in the content stream.
// This is more efficient for very small images.
type InlineImage struct {
	Width            int
	Height           int
	BitsPerComponent int
	ColorSpace       ColorSpace
	Data             []byte
	Filter           string
}

// NewInlineImage creates a new inline image.
func NewInlineImage(
	width, height int,
) *InlineImage {
	return &InlineImage{
		Width:            width,
		Height:           height,
		BitsPerComponent: 8,
		ColorSpace:       ColorSpaceDeviceRGB,
	}
}

// ContentStream returns the PDF operators for this inline image.
func (ii *InlineImage) ContentStream() string {
	var sb strings.Builder

	sb.WriteString("BI\n")

	// Image dictionary entries (abbreviated names)
	sb.WriteString(
		fmt.Sprintf("/W %d\n", ii.Width),
	)
	sb.WriteString(
		fmt.Sprintf("/H %d\n", ii.Height),
	)
	sb.WriteString(
		fmt.Sprintf(
			"/BPC %d\n",
			ii.BitsPerComponent,
		),
	)

	// Color space
	switch ii.ColorSpace {
	case ColorSpaceDeviceRGB:
		sb.WriteString("/CS /RGB\n")
	case ColorSpaceDeviceGray:
		sb.WriteString("/CS /G\n")
	case ColorSpaceDeviceCMYK:
		sb.WriteString("/CS /CMYK\n")
	case ColorSpaceIndexed:
		sb.WriteString("/CS /I\n")
	}

	// Filter
	if ii.Filter != "" {
		switch ii.Filter {
		case "FlateDecode":
			sb.WriteString("/F /Fl\n")
		case "DCTDecode":
			sb.WriteString("/F /DCT\n")
		case "ASCIIHexDecode":
			sb.WriteString("/F /AHx\n")
		default:
			sb.WriteString(
				fmt.Sprintf(
					"/F /%s\n",
					ii.Filter,
				),
			)
		}
	}

	sb.WriteString("ID\n")

	// Binary data
	sb.Write(ii.Data)

	sb.WriteString("\nEI\n")

	return sb.String()
}

// JPEG metadata extraction

// JPEGInfo contains metadata extracted from a JPEG file.
type JPEGInfo struct {
	Width         int
	Height        int
	Components    int
	BitsPerSample int
	ColorSpace    ColorSpace
	IsProgressive bool
}

// ParseJPEGInfo extracts metadata from JPEG data without fully decoding.
func ParseJPEGInfo(
	data []byte,
) (*JPEGInfo, error) {
	if len(data) < 2 || data[0] != 0xFF ||
		data[1] != 0xD8 {
		return nil, ErrInvalidImageData
	}

	info := &JPEGInfo{
		BitsPerSample: 8,
	}

	i := 2
	for i < len(data)-1 {
		if data[i] != 0xFF {
			i++

			continue
		}

		marker := data[i+1]
		i += 2

		// Skip padding FFs
		for i < len(data) && data[i] == 0xFF {
			i++
		}

		if i >= len(data) {
			break
		}

		// SOF markers (Start of Frame)
		if (marker >= 0xC0 && marker <= 0xC3) ||
			(marker >= 0xC5 && marker <= 0xC7) ||
			(marker >= 0xC9 && marker <= 0xCB) ||
			(marker >= 0xCD && marker <= 0xCF) {

			if i+7 >= len(data) {
				break
			}

			// Skip length
			i += 2

			info.BitsPerSample = int(data[i])
			info.Height = int(
				binary.BigEndian.Uint16(
					data[i+1 : i+3],
				),
			)
			info.Width = int(
				binary.BigEndian.Uint16(
					data[i+3 : i+5],
				),
			)
			info.Components = int(data[i+5])

			// Determine color space
			switch info.Components {
			case 1:
				info.ColorSpace = ColorSpaceDeviceGray
			case 3:
				info.ColorSpace = ColorSpaceDeviceRGB
			case 4:
				info.ColorSpace = ColorSpaceDeviceCMYK
			}

			// Check for progressive encoding
			if marker == 0xC2 || marker == 0xC6 ||
				marker == 0xCA ||
				marker == 0xCE {
				info.IsProgressive = true
			}

			return info, nil
		}

		// Skip other markers
		if marker != 0x00 && marker != 0xD0 &&
			marker != 0xD9 {
			if i+1 >= len(data) {
				break
			}
			length := int(
				binary.BigEndian.Uint16(
					data[i : i+2],
				),
			)
			i += length
		}
	}

	return nil, ErrInvalidImageData
}

// PNG chunk handling for metadata extraction

// PNGInfo contains metadata extracted from a PNG file.
type PNGInfo struct {
	Width       int
	Height      int
	BitDepth    int
	ColorType   int
	Compression int
	Filter      int
	Interlace   int
	HasAlpha    bool
}

// ParsePNGInfo extracts metadata from PNG data.
func ParsePNGInfo(data []byte) (*PNGInfo, error) {
	// Check PNG signature
	if len(data) < 8 {
		return nil, ErrInvalidImageData
	}
	signature := []byte{
		0x89,
		0x50,
		0x4E,
		0x47,
		0x0D,
		0x0A,
		0x1A,
		0x0A,
	}
	for i := range 8 {
		if data[i] != signature[i] {
			return nil, ErrInvalidImageData
		}
	}

	// Find IHDR chunk
	i := 8
	for i < len(data)-12 {
		length := int(
			binary.BigEndian.Uint32(
				data[i : i+4],
			),
		)
		chunkType := string(data[i+4 : i+8])

		if chunkType == "IHDR" {
			if length < 13 || i+8+13 > len(data) {
				return nil, ErrInvalidImageData
			}

			chunkData := data[i+8 : i+8+length]

			info := &PNGInfo{
				Width: int(
					binary.BigEndian.Uint32(
						chunkData[0:4],
					),
				),
				Height: int(
					binary.BigEndian.Uint32(
						chunkData[4:8],
					),
				),
				BitDepth:    int(chunkData[8]),
				ColorType:   int(chunkData[9]),
				Compression: int(chunkData[10]),
				Filter:      int(chunkData[11]),
				Interlace:   int(chunkData[12]),
			}

			// Determine if image has alpha
			// Color types: 4 = grayscale+alpha, 6 = RGB+alpha
			info.HasAlpha = info.ColorType == 4 ||
				info.ColorType == 6

			return info, nil
		}

		i += 12 + length // 4 (length) + 4 (type) + length + 4 (CRC)
	}

	return nil, ErrInvalidImageData
}

// ImageReader provides streaming image reading capabilities.
type ImageReader struct {
	reader io.Reader
	//nolint:unused // Reserved for future format detection
	format ImageFormat
}

// NewImageReader creates a new image reader from an io.Reader.
func NewImageReader(r io.Reader) *ImageReader {
	return &ImageReader{
		reader: r,
	}
}

// Read reads and loads an image, returning an ImageXObject.
func (ir *ImageReader) Read(
	name string,
) (*ImageXObject, error) {
	data, err := io.ReadAll(ir.reader)
	if err != nil {
		return nil, fmt.Errorf(
			"read image data: %w",
			err,
		)
	}

	return LoadImage(data, name)
}

// ScaleImageDimensions calculates scaled dimensions maintaining aspect ratio.
func ScaleImageDimensions(
	width, height int,
	maxWidth, maxHeight float64,
) (w, h float64) {
	fw := float64(width)
	fh := float64(height)

	if maxWidth <= 0 && maxHeight <= 0 {
		return fw, fh
	}

	scaleW := float64(1)
	scaleH := float64(1)

	if maxWidth > 0 && fw > maxWidth {
		scaleW = maxWidth / fw
	}
	if maxHeight > 0 && fh > maxHeight {
		scaleH = maxHeight / fh
	}

	// Use the smaller scale to ensure both constraints are satisfied
	scale := scaleW
	if scaleH < scaleW {
		scale = scaleH
	}

	return fw * scale, fh * scale
}
