//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"bytes"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// VideoType represents the type of a video.
type VideoType int

const (
	// VideoTypeMp4 represents an MP4 video.
	VideoTypeMp4 VideoType = iota

	// VideoTypeAvi represents an AVI video.
	VideoTypeAvi

	// VideoTypeMov represents a MOV (QuickTime) video.
	VideoTypeMov

	// VideoTypeWmv represents a WMV (Windows Media Video) video.
	VideoTypeWmv

	// VideoTypeWebm represents a WebM video.
	VideoTypeWebm

	// VideoTypeMkv represents an MKV (Matroska) video.
	VideoTypeMkv

	// VideoTypeOgv represents an OGV (Ogg Video) video.
	VideoTypeOgv
)

// String returns the string representation of the video type.
func (vt VideoType) String() string {
	switch vt {
	case VideoTypeMp4:
		return "mp4"
	case VideoTypeAvi:
		return "avi"
	case VideoTypeMov:
		return "mov"
	case VideoTypeWmv:
		return "wmv"
	case VideoTypeWebm:
		return "webm"
	case VideoTypeMkv:
		return "mkv"
	case VideoTypeOgv:
		return "ogv"
	default:
		return "mp4"
	}
}

// Extension returns the file extension for this video type.
func (vt VideoType) Extension() string {
	switch vt {
	case VideoTypeMp4:
		return ".mp4"
	case VideoTypeAvi:
		return ".avi"
	case VideoTypeMov:
		return ".mov"
	case VideoTypeWmv:
		return ".wmv"
	case VideoTypeWebm:
		return ".webm"
	case VideoTypeMkv:
		return ".mkv"
	case VideoTypeOgv:
		return ".ogv"
	default:
		return ".mp4"
	}
}

// ContentType returns the MIME content type for this video type.
func (vt VideoType) ContentType() string {
	switch vt {
	case VideoTypeMp4:
		return "video/mp4"
	case VideoTypeAvi:
		return "video/x-msvideo"
	case VideoTypeMov:
		return "video/quicktime"
	case VideoTypeWmv:
		return "video/x-ms-wmv"
	case VideoTypeWebm:
		return "video/webm"
	case VideoTypeMkv:
		return "video/x-matroska"
	case VideoTypeOgv:
		return "video/ogg"
	default:
		return "video/mp4"
	}
}

// IsSupported returns true if this video type is supported by PowerPoint.
func (vt VideoType) IsSupported() bool {
	switch vt {
	case VideoTypeMp4, VideoTypeWmv:
		return true // Native support
	case VideoTypeAvi, VideoTypeMov:
		return true // May require codec, but generally supported
	default:
		return false // WebM, MKV, OGV not natively supported
	}
}

// VideoPart represents a video part (ppt/media/video1.mp4, etc.).
type VideoPart struct {
	*baseMediaPart

	// videoType is the type of this video.
	videoType VideoType

	// streamReader holds the streaming reader when using streaming mode.
	streamReader io.Reader

	// streamSize holds the size when streaming.
	streamSize int64
}

// Counter for generating unique video filenames.
var videoUniqueCounter uint64

// Streaming threshold (100MB) - videos larger than this will use streaming.
const streamingThreshold = 100 * 1024 * 1024

// NewVideoPartForSlide creates a new video part for a slide.
func NewVideoPartForSlide(
	slidePart *SlidePart,
	videoType VideoType,
	streaming bool,
) (*VideoPart, error) {
	num := atomic.AddUint64(
		&videoUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/media/video%d%s",
		num,
		videoType.Extension(),
	)

	packPart, relID, err := slidePart.addChildPart(
		uri,
		videoType.ContentType(),
		RelationshipTypeVideo,
	)
	if err != nil {
		return nil, err
	}

	basePart := newBaseMediaPart(
		uri,
		videoType.ContentType(),
		packPart,
		slidePart,
		streaming,
	)

	vp := &VideoPart{
		baseMediaPart: basePart,
		videoType:     videoType,
		streamSize:    0,
	}

	vp.SetRelationshipID(relID)

	// Add to slide part's child parts
	if err := slidePart.AddPart(vp, relID); err != nil {
		return nil, err
	}

	return vp, nil
}

// VideoType returns the type of this video.
func (vp *VideoPart) VideoType() VideoType {
	return vp.videoType
}

// SetStream sets the video content from a reader.
// For streaming parts, it stores the reader reference.
// For non-streaming parts, it reads all data into memory.
func (vp *VideoPart) SetStream(r io.Reader) error {
	if vp.streaming {
		// For streaming, we need to determine the size first
		// We'll use a tee reader to count bytes while reading
		var size int64
		teeReader := io.TeeReader(r, &sizeCounter{size: &size})

		vp.streamReader = teeReader
		vp.streamSize = size
		vp.setSize(size)

		// Mark as dirty since we have content
		vp.MarkDirty()

		return nil
	}

	// For non-streaming, use the parent implementation
	return vp.baseMediaPart.SetStream(r)
}

// GetStream returns a reader for the video content.
func (vp *VideoPart) GetStream() io.Reader {
	if vp.streaming && vp.streamReader != nil {
		return vp.streamReader
	}

	// Fall back to the default implementation
	return vp.baseMediaPart.GetStream()
}

// FeedData reads video data from the provided reader and sets it as the part content.
func (vp *VideoPart) FeedData(r io.Reader) error {
	if vp.streaming {
		return vp.SetStream(r)
	}

	// For non-streaming, read all data
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	vp.SetData(data)
	vp.setSize(int64(len(data)))

	return nil
}

// FeedDataBytes sets the video data from a byte slice.
func (vp *VideoPart) FeedDataBytes(data []byte) error {
	if vp.streaming {
		// For streaming mode, create a reader from the bytes
		return vp.SetStream(bytes.NewReader(data))
	}

	// For non-streaming, use the default implementation
	vp.SetData(data)
	vp.setSize(int64(len(data)))
	return nil
}

// GetData returns the raw video data.
// Note: For streaming parts, this will read all data into memory.
func (vp *VideoPart) GetData() ([]byte, error) {
	if vp.streaming {
		if vp.streamReader == nil {
			return nil, nil
		}

		// Read all data from the stream
		return io.ReadAll(vp.streamReader)
	}

	// For non-streaming, use the default implementation
	stream := vp.GetStream()
	if stream == nil {
		return nil, nil
	}

	return io.ReadAll(stream)
}

// ShouldStream returns true if the given size suggests streaming should be used.
func ShouldStream(size int64) bool {
	return size > streamingThreshold
}

// Ensure VideoPart implements MediaPart.
var _ MediaPart = (*VideoPart)(nil)

// VideoPartFactory creates a VideoPart from a URI and container.
func VideoPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	contentType := packPart.ContentType()
	videoType := videoTypeFromContentType(contentType)

	// Determine if we should use streaming based on size
	streaming := ShouldStream(packPart.Size())

	basePart := newBaseMediaPart(
		uri,
		contentType,
		packPart,
		container,
		streaming,
	)

	return &VideoPart{
		baseMediaPart: basePart,
		videoType:     videoType,
	}
}

// videoTypeFromContentType determines the VideoType from a content type string.
func videoTypeFromContentType(
	contentType string,
) VideoType {
	switch contentType {
	case "video/mp4":
		return VideoTypeMp4
	case "video/x-msvideo", "video/avi":
		return VideoTypeAvi
	case "video/quicktime":
		return VideoTypeMov
	case "video/x-ms-wmv":
		return VideoTypeWmv
	case "video/webm":
		return VideoTypeWebm
	case "video/x-matroska":
		return VideoTypeMkv
	case "video/ogg":
		return VideoTypeOgv
	default:
		return VideoTypeMp4
	}
}

// VideoTypeFromExtension determines the VideoType from a file extension.
func VideoTypeFromExtension(
	ext string,
) VideoType {
	switch ext {
	case ".mp4":
		return VideoTypeMp4
	case ".avi":
		return VideoTypeAvi
	case ".mov":
		return VideoTypeMov
	case ".wmv":
		return VideoTypeWmv
	case ".webm":
		return VideoTypeWebm
	case ".mkv":
		return VideoTypeMkv
	case ".ogv":
		return VideoTypeOgv
	default:
		return VideoTypeMp4
	}
}

// VideoTypeFromFilename determines the VideoType from a filename.
func VideoTypeFromFilename(
	filename string,
) VideoType {
	ext := getFileExtension(filename)
	return VideoTypeFromExtension(ext)
}

// getFileExtension extracts the file extension from a filename.
func getFileExtension(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i:]
		}
		if filename[i] == '/' || filename[i] == '\\' {
			break
		}
	}
	return ""
}

// Magic byte signatures for video formats
var (
	// MP4 signature - can start with ftyp or moov
	mp4Magic = [][]byte{
		{0x66, 0x74, 0x79, 0x70}, // "ftyp"
		{0x6D, 0x6F, 0x6F, 0x76}, // "moov"
		{0x6D, 0x64, 0x61, 0x74}, // "mdat"
	}

	// AVI signature - "RIFF" + file size + "AVI "
	aviMagic = []byte{0x52, 0x49, 0x46, 0x46}

	// MOV signature - same as MP4 (both are QuickTime containers)
	movMagic = mp4Magic

	// WMV/ASF signature - typically starts with 0x30 0x26 0xB2 0x75
	wmvMagic = []byte{0x30, 0x26, 0xB2, 0x75}

	// WebM signature - EBML header
	webmMagic = []byte{0x1A, 0x45, 0xDF, 0xA3}

	// MKV signature - same as WebM (both are Matroska)
	mkvMagic = webmMagic

	// OGV signature - Ogg container
	ogvMagic = []byte{0x4F, 0x67, 0x67, 0x53} // "OggS"
)

// VideoTypeFromMagicBytes determines the VideoType from the file's magic bytes.
// Returns the detected video type and true if detected, or VideoTypeMp4 and false if unknown.
func VideoTypeFromMagicBytes(
	data []byte,
) (VideoType, bool) {
	if len(data) < 16 {
		return VideoTypeMp4, false
	}

	// Check AVI
	if bytes.HasPrefix(data, aviMagic) {
		// Additional check for "AVI " at offset 8
		if len(data) >= 12 && string(data[8:12]) == "AVI " {
			return VideoTypeAvi, true
		}
	}

	// Check MP4/MOV
	for _, magic := range mp4Magic {
		if bytes.Contains(data[:16], magic) {
			// MP4 and MOV are structurally similar
			// We'll need additional context to distinguish
			// For now, default to MP4
			return VideoTypeMp4, true
		}
	}

	// Check WMV
	if bytes.HasPrefix(data, wmvMagic) {
		return VideoTypeWmv, true
	}

	// Check WebM/MKV
	if bytes.HasPrefix(data, webmMagic) {
		// Additional check for "webm" or "matroska" in the header
		if bytes.Contains(data, []byte("webm")) {
			return VideoTypeWebm, true
		}
		if bytes.Contains(data, []byte("matroska")) {
			return VideoTypeMkv, true
		}
	}

	// Check OGV
	if bytes.HasPrefix(data, ogvMagic) {
		return VideoTypeOgv, true
	}

	return VideoTypeMp4, false
}

// DetectVideoType attempts to detect the video type from data, falling back to extension.
func DetectVideoType(
	data []byte,
	filename string,
) VideoType {
	// First try magic bytes
	if videoType, detected := VideoTypeFromMagicBytes(data); detected {
		return videoType
	}

	// Fall back to extension
	if filename != "" {
		return VideoTypeFromFilename(filename)
	}

	// Default to MP4
	return VideoTypeMp4
}

// sizeCounter is a writer that counts bytes written to it.
type sizeCounter struct {
	size *int64
}

func (sc *sizeCounter) Write(p []byte) (n int, err error) {
	*sc.size += int64(len(p))
	return len(p), nil
}

// Register video content types.
func init() {
	videoTypes := []VideoType{
		VideoTypeMp4,
		VideoTypeAvi,
		VideoTypeMov,
		VideoTypeWmv,
		VideoTypeWebm,
		VideoTypeMkv,
		VideoTypeOgv,
	}

	for _, vt := range videoTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:      vt.ContentType(),
				RelationshipType: RelationshipTypeVideo,
				Factory:          VideoPartFactory,
				DefaultURI: fmt.Sprintf(
					"/ppt/media/video1%s",
					vt.Extension(),
				),
				IsFixedContentType: true,
			},
		)
	}
}