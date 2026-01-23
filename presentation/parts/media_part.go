//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"io"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// MediaPart interface defines common methods for media parts (video, audio).
type MediaPart interface {
	openxml.OpenXmlPart

	// RelationshipID returns the relationship ID of this part.
	RelationshipID() string

	// GetStream returns a reader for the media content.
	GetStream() io.Reader

	// SetStream sets the media content from a reader.
	SetStream(r io.Reader) error

	// GetSize returns the size of the media content in bytes.
	GetSize() int64

	// GetContentType returns the MIME content type of the media.
	GetContentType() string

	// IsStreaming returns true if this part uses streaming (for large files).
	IsStreaming() bool

	// GetData returns the raw media data.
	GetData() ([]byte, error)

	// FeedDataBytes sets the media data from a byte slice.
	FeedDataBytes(data []byte) error
}

// baseMediaPart provides common implementation for media parts.
type baseMediaPart struct {
	*openxml.OpenXmlPartData

	// streaming indicates if this part uses streaming for large files.
	streaming bool

	// size of the media content in bytes.
	size int64
}

// ensure baseMediaPart implements MediaPart.
var _ MediaPart = (*baseMediaPart)(nil)

// newBaseMediaPart creates a new base media part.
func newBaseMediaPart(
	uri, contentType string,
	packPart *packaging.Part,
	container openxml.OpenXmlPartContainer,
	streaming bool,
) *baseMediaPart {
	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)

	return &baseMediaPart{
		OpenXmlPartData: partData,
		streaming:       streaming,
		size:            0,
	}
}

// GetContentType returns the MIME content type of the media.
func (bmp *baseMediaPart) GetContentType() string {
	return bmp.ContentType()
}

// IsStreaming returns true if this part uses streaming.
func (bmp *baseMediaPart) IsStreaming() bool {
	return bmp.streaming
}

// GetSize returns the size of the media content.
func (bmp *baseMediaPart) GetSize() int64 {
	return bmp.size
}

// GetData returns the raw media data.
func (bmp *baseMediaPart) GetData() ([]byte, error) {
	return bmp.OpenXmlPartData.GetData(), nil
}

// FeedDataBytes sets the media data from a byte slice.
func (bmp *baseMediaPart) FeedDataBytes(data []byte) error {
	bmp.SetData(data)
	bmp.setSize(int64(len(data)))
	return nil
}

// setSize sets the size of the media content.
func (bmp *baseMediaPart) setSize(size int64) {
	bmp.size = size
}

// SetStream sets the media content from a reader.
// For streaming parts, this stores a reference to the reader.
// For non-streaming parts, it reads all data into memory.
func (bmp *baseMediaPart) SetStream(r io.Reader) error {
	if bmp.streaming {
		// For streaming, we need to handle this differently
		// The actual streaming implementation will be in the concrete types
		return nil
	}

	// For non-streaming, use the default implementation
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	bmp.SetData(data)
	bmp.setSize(int64(len(data)))

	return nil
}