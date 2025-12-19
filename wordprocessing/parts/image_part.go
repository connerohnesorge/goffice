package parts

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// ImageType represents the type of an image.
type ImageType int

const (
	// ImageTypePng represents a PNG image.
	ImageTypePng ImageType = iota

	// ImageTypeJpeg represents a JPEG image.
	ImageTypeJpeg

	// ImageTypeGif represents a GIF image.
	ImageTypeGif

	// ImageTypeBmp represents a BMP image.
	ImageTypeBmp

	// ImageTypeTiff represents a TIFF image.
	ImageTypeTiff

	// ImageTypeEmf represents an EMF (Enhanced Metafile) image.
	ImageTypeEmf

	// ImageTypeWmf represents a WMF (Windows Metafile) image.
	ImageTypeWmf

	// ImageTypeIcon represents an ICO icon.
	ImageTypeIcon
)

// String returns the string representation of the image type.
func (it ImageType) String() string {
	switch it {
	case ImageTypePng:
		return "png"
	case ImageTypeJpeg:
		return "jpeg"
	case ImageTypeGif:
		return "gif"
	case ImageTypeBmp:
		return "bmp"
	case ImageTypeTiff:
		return "tiff"
	case ImageTypeEmf:
		return "emf"
	case ImageTypeWmf:
		return "wmf"
	case ImageTypeIcon:
		return "ico"
	default:
		return "png"
	}
}

// Extension returns the file extension for this image type.
func (it ImageType) Extension() string {
	switch it {
	case ImageTypePng:
		return ".png"
	case ImageTypeJpeg:
		return ".jpeg"
	case ImageTypeGif:
		return ".gif"
	case ImageTypeBmp:
		return ".bmp"
	case ImageTypeTiff:
		return ".tiff"
	case ImageTypeEmf:
		return ".emf"
	case ImageTypeWmf:
		return ".wmf"
	case ImageTypeIcon:
		return ".ico"
	default:
		return ".png"
	}
}

// ContentType returns the MIME content type for this image type.
func (it ImageType) ContentType() string {
	switch it {
	case ImageTypePng:
		return "image/png"
	case ImageTypeJpeg:
		return "image/jpeg"
	case ImageTypeGif:
		return "image/gif"
	case ImageTypeBmp:
		return "image/bmp"
	case ImageTypeTiff:
		return "image/tiff"
	case ImageTypeEmf:
		return "image/x-emf"
	case ImageTypeWmf:
		return "image/x-wmf"
	case ImageTypeIcon:
		return "image/x-icon"
	default:
		return "image/png"
	}
}

// ImagePart represents an image part (word/media/image1.png, etc.).
type ImagePart struct {
	*openxml.OpenXmlPartData

	// imageType is the type of this image.
	imageType ImageType
}

// Relationship type for images.
const (
	RelationshipTypeImage = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
)

// Counter for generating unique image filenames.
var imageCounter uint64

// newImagePart creates a new image part.
func newImagePart(
	mainPart *MainPart,
	imageType ImageType,
) (*ImagePart, error) {
	num := atomic.AddUint64(&imageCounter, 1)
	uri := fmt.Sprintf(
		"/word/media/image%d%s",
		num,
		imageType.Extension(),
	)

	packPart, relID, err := mainPart.addChildPart(
		uri,
		imageType.ContentType(),
		RelationshipTypeImage,
	)
	if err != nil {
		return nil, err
	}

	partData := openxml.NewOpenXmlPartData(
		uri,
		imageType.ContentType(),
		packPart,
		mainPart,
	)
	partData.SetRelationshipID(relID)

	ip := &ImagePart{
		OpenXmlPartData: partData,
		imageType:       imageType,
	}

	// Add to main part's child parts
	if err := mainPart.AddPart(ip, relID); err != nil {
		return nil, err
	}

	return ip, nil
}

// ImageType returns the type of this image.
func (ip *ImagePart) ImageType() ImageType {
	return ip.imageType
}

// FeedData reads image data from the provided reader and sets it as the part content.
func (ip *ImagePart) FeedData(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	ip.SetData(data)

	return nil
}

// GetStream returns a reader for the part content.
func (ip *ImagePart) GetStream() io.Reader {
	return ip.OpenXmlPartData.GetStream()
}

// GetData returns the raw image data.
func (ip *ImagePart) GetData() ([]byte, error) {
	stream := ip.GetStream()
	if stream == nil {
		return nil, nil
	}

	return io.ReadAll(stream)
}

// FeedDataBytes sets the image data from a byte slice.
func (ip *ImagePart) FeedDataBytes(data []byte) {
	ip.SetData(data)
}

// Ensure ImagePart implements OpenXmlPart.
var _ openxml.OpenXmlPart = (*ImagePart)(nil)

// ImagePartFactory creates an ImagePart from a URI and container.
func ImagePartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	pkg := container.Package()
	if pkg == nil {
		return nil
	}

	packPart, err := pkg.Part(uri)
	if err != nil {
		return nil
	}

	contentType := packPart.ContentType()
	imageType := imageTypeFromContentType(
		contentType,
	)

	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)

	return &ImagePart{
		OpenXmlPartData: partData,
		imageType:       imageType,
	}
}

// imageTypeFromContentType determines the ImageType from a content type string.
func imageTypeFromContentType(
	contentType string,
) ImageType {
	switch contentType {
	case "image/png":
		return ImageTypePng
	case "image/jpeg", "image/jpg":
		return ImageTypeJpeg
	case "image/gif":
		return ImageTypeGif
	case "image/bmp":
		return ImageTypeBmp
	case "image/tiff":
		return ImageTypeTiff
	case "image/x-emf":
		return ImageTypeEmf
	case "image/x-wmf":
		return ImageTypeWmf
	case "image/x-icon":
		return ImageTypeIcon
	default:
		return ImageTypePng
	}
}

// ImageTypeFromExtension determines the ImageType from a file extension.
func ImageTypeFromExtension(
	ext string,
) ImageType {
	ext = strings.ToLower(ext)
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	switch ext {
	case ".png":
		return ImageTypePng
	case ".jpg", ".jpeg":
		return ImageTypeJpeg
	case ".gif":
		return ImageTypeGif
	case ".bmp":
		return ImageTypeBmp
	case ".tif", ".tiff":
		return ImageTypeTiff
	case ".emf":
		return ImageTypeEmf
	case ".wmf":
		return ImageTypeWmf
	case ".ico":
		return ImageTypeIcon
	default:
		return ImageTypePng
	}
}

// ImageTypeFromFilename determines the ImageType from a filename.
func ImageTypeFromFilename(
	filename string,
) ImageType {
	ext := filepath.Ext(filename)

	return ImageTypeFromExtension(ext)
}

// Magic byte signatures for image formats
var (
	pngMagic = []byte{
		0x89,
		0x50,
		0x4E,
		0x47,
		0x0D,
		0x0A,
		0x1A,
		0x0A,
	}
	jpegMagic = []byte{0xFF, 0xD8, 0xFF}
	gifMagic  = []byte{0x47, 0x49, 0x46}
	bmpMagic  = []byte{0x42, 0x4D}
	// TIFF can be either little-endian or big-endian
	tiffLEMagic = []byte{0x49, 0x49, 0x2A, 0x00}
	tiffBEMagic = []byte{0x4D, 0x4D, 0x00, 0x2A}
	// EMF starts with EMF signature
	emfMagic = []byte{0x01, 0x00, 0x00, 0x00}
	// WMF starts with placeable header or standard header
	wmfPlaceableMagic = []byte{
		0xD7,
		0xCD,
		0xC6,
		0x9A,
	}
	// ICO starts with reserved (0) and type (1 for icon)
	icoMagic = []byte{0x00, 0x00, 0x01, 0x00}
)

// ImageTypeFromMagicBytes determines the ImageType from the file's magic bytes.
// Returns the detected image type and true if detected, or ImageTypePng and false if unknown.
func ImageTypeFromMagicBytes(
	data []byte,
) (ImageType, bool) {
	if len(data) < 8 {
		return ImageTypePng, false
	}

	if bytes.HasPrefix(data, pngMagic) {
		return ImageTypePng, true
	}
	if bytes.HasPrefix(data, jpegMagic) {
		return ImageTypeJpeg, true
	}
	if bytes.HasPrefix(data, gifMagic) {
		return ImageTypeGif, true
	}
	if bytes.HasPrefix(data, bmpMagic) {
		return ImageTypeBmp, true
	}
	if bytes.HasPrefix(data, tiffLEMagic) ||
		bytes.HasPrefix(data, tiffBEMagic) {
		return ImageTypeTiff, true
	}
	if bytes.HasPrefix(data, wmfPlaceableMagic) {
		return ImageTypeWmf, true
	}
	if bytes.HasPrefix(data, icoMagic) {
		return ImageTypeIcon, true
	}
	// EMF detection needs more context as 0x01000000 is not unique
	// Check for EMF header at offset 40
	if len(data) >= 44 &&
		bytes.HasPrefix(data, emfMagic) {
		// Additional check: EMF has " EMF" at offset 40
		if len(data) >= 44 && data[40] == 0x20 &&
			data[41] == 0x45 &&
			data[42] == 0x4D &&
			data[43] == 0x46 {
			return ImageTypeEmf, true
		}
	}

	return ImageTypePng, false
}

// DetectImageType attempts to detect the image type from data, falling back to extension.
func DetectImageType(
	data []byte,
	filename string,
) ImageType {
	// First try magic bytes
	if imgType, detected := ImageTypeFromMagicBytes(data); detected {
		return imgType
	}

	// Fall back to extension
	if filename != "" {
		return ImageTypeFromFilename(filename)
	}

	// Default to PNG
	return ImageTypePng
}

// Register image content types (we register them by content type, not relationship).
func init() {
	imageTypes := []ImageType{
		ImageTypePng,
		ImageTypeJpeg,
		ImageTypeGif,
		ImageTypeBmp,
		ImageTypeTiff,
		ImageTypeEmf,
		ImageTypeWmf,
		ImageTypeIcon,
	}

	for _, it := range imageTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:      it.ContentType(),
				RelationshipType: RelationshipTypeImage,
				Factory:          ImagePartFactory,
				DefaultURI: fmt.Sprintf(
					"/word/media/image1%s",
					it.Extension(),
				),
				IsFixedContentType: true,
			},
		)
	}
}
