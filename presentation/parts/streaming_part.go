//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/packaging"
)

// StreamingPart provides streaming support for large media files.
// This implementation stores large files in temporary files rather than memory.
type StreamingPart struct {
	*openxml.OpenXmlPartData

	// tempFile holds the temporary file for streaming content.
	tempFile *os.File

	// tempFilePath is the path to the temporary file.
	tempFilePath string

	// size of the content in bytes.
	size int64

	// mu protects concurrent access.
	mu sync.RWMutex
}

// newStreamingPart creates a new streaming part.
func newStreamingPart(
	uri, contentType string,
	packPart *packaging.Part,
	container openxml.OpenXmlPartContainer,
) (*StreamingPart, error) {
	partData := openxml.NewOpenXmlPartData(
		uri,
		contentType,
		packPart,
		container,
	)

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "goffice-streaming-*.tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	return &StreamingPart{
		OpenXmlPartData: partData,
		tempFile:        tempFile,
		tempFilePath:    tempFile.Name(),
		size:            0,
	}, nil
}

// SetStream sets the part content from a reader using streaming.
func (sp *StreamingPart) SetStream(r io.Reader) error {
	sp.mu.Lock()
	defer sp.mu.Unlock()

	// Reset the temp file
	if err := sp.tempFile.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate temp file: %w", err)
	}
	if _, err := sp.tempFile.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek temp file: %w", err)
	}

	// Copy data to temp file
	n, err := io.Copy(sp.tempFile, r)
	if err != nil {
		return fmt.Errorf("failed to copy to temp file: %w", err)
	}

	sp.size = n
	sp.MarkDirty()

	return nil
}

// GetStream returns a reader for the part content.
func (sp *StreamingPart) GetStream() io.Reader {
	sp.mu.RLock()
	defer sp.mu.RUnlock()

	// Open the temp file for reading
	file, err := os.Open(sp.tempFilePath)
	if err != nil {
		return nil
	}

	// Return a reader that will close the file when done
	return &autoCloseReader{file: file}
}

// GetSize returns the size of the content.
func (sp *StreamingPart) GetSize() int64 {
	sp.mu.RLock()
	defer sp.mu.RUnlock()
	return sp.size
}

// Close closes the streaming part and cleans up temporary files.
func (sp *StreamingPart) Close() error {
	sp.mu.Lock()
	defer sp.mu.Unlock()

	var err error
	if sp.tempFile != nil {
		err = sp.tempFile.Close()
		sp.tempFile = nil
	}

	// Remove temp file
	if sp.tempFilePath != "" {
		removeErr := os.Remove(sp.tempFilePath)
		if removeErr != nil && err == nil {
			err = removeErr
		}
		sp.tempFilePath = ""
	}

	return err
}

// Save saves the streaming content to the package.
func (sp *StreamingPart) Save() error {
	sp.mu.RLock()
	defer sp.mu.RUnlock()

	if sp.tempFile == nil {
		return fmt.Errorf("no content to save")
	}

	// Sync the temp file to ensure all data is written
	if err := sp.tempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	// Seek to beginning
	if _, err := sp.tempFile.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek temp file: %w", err)
	}

	// Copy to the packaging part
	if packPart := sp.PackagingPart(); packPart != nil {
		return packPart.SetStream(sp.tempFile)
	}

	return nil
}

// autoCloseReader automatically closes the file when EOF is reached.
type autoCloseReader struct {
	file *os.File
	once sync.Once
}

func (acr *autoCloseReader) Read(p []byte) (n int, err error) {
	n, err = acr.file.Read(p)
	if err == io.EOF {
		// Close the file when EOF is reached
		acr.once.Do(func() {
			acr.file.Close()
		})
	}
	return n, err
}

// StreamingVideoPart combines video part with streaming support.
type StreamingVideoPart struct {
	*VideoPart
	*StreamingPart
}

// NewStreamingVideoPartForSlide creates a new streaming video part.
func NewStreamingVideoPartForSlide(
	slidePart *SlidePart,
	videoType VideoType,
) (*StreamingVideoPart, error) {
	// Create the base video part
	videoPart, err := NewVideoPartForSlide(slidePart, videoType, true)
	if err != nil {
		return nil, err
	}

	// Create the streaming part
	streamingPart, err := newStreamingPart(
		videoPart.URI(),
		videoPart.ContentType(),
		videoPart.PackagingPart(),
		slidePart,
	)
	if err != nil {
		return nil, err
	}

	return &StreamingVideoPart{
		VideoPart:     videoPart,
		StreamingPart: streamingPart,
	}, nil
}

// SetStream sets the video content using streaming.
func (svp *StreamingVideoPart) SetStream(r io.Reader) error {
	// Use the streaming part's implementation
	err := svp.StreamingPart.SetStream(r)
	if err != nil {
		return err
	}

	// Update the video part's size
	svp.VideoPart.setSize(svp.StreamingPart.GetSize())

	return nil
}

// GetStream returns a reader for the video content.
func (svp *StreamingVideoPart) GetStream() io.Reader {
	// Use the streaming part's implementation
	return svp.StreamingPart.GetStream()
}

// GetSize returns the size of the video.
func (svp *StreamingVideoPart) GetSize() int64 {
	// Use the streaming part's size
	return svp.StreamingPart.GetSize()
}

// Save saves both the video metadata and streaming content.
func (svp *StreamingVideoPart) Save() error {
	// Save the streaming content first
	if err := svp.StreamingPart.Save(); err != nil {
		return err
	}

	// Then save the video part (which will save relationships, etc.)
	if saveable, ok := any(svp.VideoPart).(openxml.ISaveablePart); ok {
		return saveable.Save()
	}

	return nil
}

// Close closes the streaming video part.
func (svp *StreamingVideoPart) Close() error {
	// Close the streaming part
	if err := svp.StreamingPart.Close(); err != nil {
		return err
	}

	return nil
}

// Ensure StreamingVideoPart implements MediaPart.
var _ MediaPart = (*StreamingVideoPart)(nil)

// StreamingVideoPartFactory creates a StreamingVideoPart from a URI and container.
func StreamingVideoPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// First create a regular video part
	videoPart := VideoPartFactory(uri, container)
	if videoPart == nil {
		return nil
	}

	vp := videoPart.(*VideoPart)

	// Create the streaming part
	streamingPart, err := newStreamingPart(
		vp.URI(),
		vp.ContentType(),
		vp.PackagingPart(),
		container,
	)
	if err != nil {
		return nil
	}

	return &StreamingVideoPart{
		VideoPart:     vp,
		StreamingPart: streamingPart,
	}
}

// StreamingAudioPart combines audio part with streaming support.
type StreamingAudioPart struct {
	*AudioPart
	*StreamingPart
}

// NewStreamingAudioPartForSlide creates a new streaming audio part.
func NewStreamingAudioPartForSlide(
	slidePart *SlidePart,
	audioType AudioType,
) (*StreamingAudioPart, error) {
	// Create the base audio part
	audioPart, err := NewAudioPartForSlide(slidePart, audioType, true)
	if err != nil {
		return nil, err
	}

	// Create the streaming part
	streamingPart, err := newStreamingPart(
		audioPart.URI(),
		audioPart.ContentType(),
		audioPart.PackagingPart(),
		slidePart,
	)
	if err != nil {
		return nil, err
	}

	return &StreamingAudioPart{
		AudioPart:     audioPart,
		StreamingPart: streamingPart,
	}, nil
}

// SetStream sets the audio content using streaming.
func (sap *StreamingAudioPart) SetStream(r io.Reader) error {
	// Use the streaming part's implementation
	err := sap.StreamingPart.SetStream(r)
	if err != nil {
		return err
	}

	// Update the audio part's size
	sap.AudioPart.setSize(sap.StreamingPart.GetSize())

	return nil
}

// GetStream returns a reader for the audio content.
func (sap *StreamingAudioPart) GetStream() io.Reader {
	// Use the streaming part's implementation
	return sap.StreamingPart.GetStream()
}

// GetSize returns the size of the audio.
func (sap *StreamingAudioPart) GetSize() int64 {
	// Use the streaming part's size
	return sap.StreamingPart.GetSize()
}

// Save saves both the audio metadata and streaming content.
func (sap *StreamingAudioPart) Save() error {
	// Save the streaming content first
	if err := sap.StreamingPart.Save(); err != nil {
		return err
	}

	// Then save the audio part
	if saveable, ok := any(sap.AudioPart).(openxml.ISaveablePart); ok {
		return saveable.Save()
	}

	return nil
}

// Close closes the streaming audio part.
func (sap *StreamingAudioPart) Close() error {
	// Close the streaming part
	if err := sap.StreamingPart.Close(); err != nil {
		return err
	}

	return nil
}

// Ensure StreamingAudioPart implements MediaPart.
var _ MediaPart = (*StreamingAudioPart)(nil)

// StreamingAudioPartFactory creates a StreamingAudioPart from a URI and container.
func StreamingAudioPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	// First create a regular audio part
	audioPart := AudioPartFactory(uri, container)
	if audioPart == nil {
		return nil
	}

	ap := audioPart.(*AudioPart)

	// Create the streaming part
	streamingPart, err := newStreamingPart(
		ap.URI(),
		ap.ContentType(),
		ap.PackagingPart(),
		container,
	)
	if err != nil {
		return nil
	}

	return &StreamingAudioPart{
		AudioPart:     ap,
		StreamingPart: streamingPart,
	}
}

// Helper function to determine if streaming should be used.
func shouldUseStreamingForSize(size int64) bool {
	return size > streamingThreshold
}

// Helper function to create appropriate video part based on size.
func createVideoPartForSlideWithSize(
	slidePart *SlidePart,
	videoType VideoType,
	size int64,
) (MediaPart, error) {
	if shouldUseStreamingForSize(size) {
		return NewStreamingVideoPartForSlide(slidePart, videoType)
	}

	// Use regular video part for smaller files
	return NewVideoPartForSlide(slidePart, videoType, false)
}