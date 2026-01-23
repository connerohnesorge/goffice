//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package presentation

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/connerohnesorge/goffice/presentation/parts"
)

// AddVideoOptions contains options for adding a video to a slide.
type AddVideoOptions struct {
	// StreamingThreshold is the file size threshold for using streaming (default: 100MB).
	StreamingThreshold int64

	// ValidateFormat enables format validation (default: true).
	ValidateFormat bool

	// AutoDetectType enables automatic detection of video type from content (default: true).
	AutoDetectType bool

	// VideoType specifies the video type when auto-detection is disabled.
	VideoType parts.VideoType
}

// DefaultAddVideoOptions returns default options for adding videos.
func DefaultAddVideoOptions() *AddVideoOptions {
	return &AddVideoOptions{
		StreamingThreshold: 100 * 1024 * 1024, // 100MB
		ValidateFormat:     true,
		AutoDetectType:     true,
		VideoType:          parts.VideoTypeMp4,
	}
}

// AddVideoFromFile adds a video to the presentation from a file path.
func (p *Presentation) AddVideoFromFile(
	slideIndex int,
	filePath string,
	opts *AddVideoOptions,
) (parts.MediaPart, error) {
	if opts == nil {
		opts = DefaultAddVideoOptions()
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open video file: %w", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Determine if we should use streaming
	useStreaming := fileInfo.Size() > opts.StreamingThreshold

	// Detect video type if needed
	var videoType parts.VideoType
	if opts.AutoDetectType {
		// Read first few bytes for detection
		header := make([]byte, 1024)
		n, _ := file.Read(header)
		file.Seek(0, 0) // Reset position

		videoType = parts.DetectVideoType(header[:n], filePath)
	} else {
		videoType = opts.VideoType
	}

	// Validate if requested
	if opts.ValidateFormat {
		header := make([]byte, 1024)
		n, _ := file.Read(header)
		file.Seek(0, 0) // Reset position

		if err := parts.ValidateMedia(header[:n], filePath); err != nil {
			return nil, fmt.Errorf("media validation failed: %w", err)
		}

		// Check if format is supported
		if !videoType.IsSupported() {
			return nil, fmt.Errorf("video format %s is not natively supported in PowerPoint", videoType.String())
		}
	}

	// Get the slide
	slides := p.Slides()
	if slideIndex < 0 || slideIndex >= len(slides) {
		return nil, fmt.Errorf("slide index %d out of range", slideIndex)
	}
	slide := slides[slideIndex]

	// Create the video part
	var videoPart parts.MediaPart
	if useStreaming {
		// Create streaming video part
		streamingPart, err := parts.NewStreamingVideoPartForSlide(
			slide.SlidePart(),
			videoType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create streaming video part: %w", err)
		}
		videoPart = streamingPart
	} else {
		// Create regular video part
		regularPart, err := parts.NewVideoPartForSlide(
			slide.SlidePart(),
			videoType,
			false, // no streaming
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create video part: %w", err)
		}
		videoPart = regularPart
	}

	// Set the content
	if err := videoPart.SetStream(file); err != nil {
		return nil, fmt.Errorf("failed to set video content: %w", err)
	}

	return videoPart, nil
}

// AddVideoFromReader adds a video to the presentation from an io.Reader.
func (p *Presentation) AddVideoFromReader(
	slideIndex int,
	reader io.Reader,
	filename string,
	opts *AddVideoOptions,
) (parts.MediaPart, error) {
	if opts == nil {
		opts = DefaultAddVideoOptions()
	}

	// If auto-detection is enabled and we have a filename, detect from filename
	var videoType parts.VideoType
	if opts.AutoDetectType && filename != "" {
		videoType = parts.VideoTypeFromFilename(filename)
	} else {
		videoType = opts.VideoType
	}

	// Validate if requested
	if opts.ValidateFormat && filename != "" {
		// We need to read some data for validation
		header := make([]byte, 1024)
		n, err := reader.Read(header)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read header for validation: %w", err)
		}

		// Create a new reader that includes the header we just read
		reader = io.MultiReader(bytes.NewReader(header[:n]), reader)

		if err := parts.ValidateMedia(header[:n], filename); err != nil {
			return nil, fmt.Errorf("media validation failed: %w", err)
		}

		// Check if format is supported
		if !videoType.IsSupported() {
			return nil, fmt.Errorf("video format %s is not natively supported in PowerPoint", videoType.String())
		}
	}

	// Get the slide
	slides := p.Slides()
	if slideIndex < 0 || slideIndex >= len(slides) {
		return nil, fmt.Errorf("slide index %d out of range", slideIndex)
	}
	slide := slides[slideIndex]

	// Create the video part (we'll use regular part for reader-based content)
	videoPart, err := parts.NewVideoPartForSlide(
		slide.SlidePart(),
		videoType,
		false, // no streaming for reader-based content
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create video part: %w", err)
	}

	// Set the content
	if err := videoPart.SetStream(reader); err != nil {
		return nil, fmt.Errorf("failed to set video content: %w", err)
	}

	return videoPart, nil
}

// AddVideoFromBytes adds a video to the presentation from a byte slice.
func (p *Presentation) AddVideoFromBytes(
	slideIndex int,
	data []byte,
	filename string,
	opts *AddVideoOptions,
) (parts.MediaPart, error) {
	if opts == nil {
		opts = DefaultAddVideoOptions()
	}

	// Detect video type if needed
	var videoType parts.VideoType
	if opts.AutoDetectType && filename != "" {
		videoType = parts.DetectVideoType(data, filename)
	} else if opts.AutoDetectType {
		videoType = parts.VideoTypeFromMagicBytes(data)
	} else {
		videoType = opts.VideoType
	}

	// Validate if requested
	if opts.ValidateFormat {
		if err := parts.ValidateMedia(data, filename); err != nil {
			return nil, fmt.Errorf("media validation failed: %w", err)
		}

		// Check if format is supported
		if !videoType.IsSupported() {
			return nil, fmt.Errorf("video format %s is not natively supported in PowerPoint", videoType.String())
		}
	}

	// Get the slide
	slides := p.Slides()
	if slideIndex < 0 || slideIndex >= len(slides) {
		return nil, fmt.Errorf("slide index %d out of range", slideIndex)
	}
	slide := slides[slideIndex]

	// Create the video part
	videoPart, err := parts.NewVideoPartForSlide(
		slide.SlidePart(),
		videoType,
		false, // no streaming for byte-based content
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create video part: %w", err)
	}

	// Set the content
	if err := videoPart.FeedDataBytes(data); err != nil {
		return nil, fmt.Errorf("failed to set video content: %w", err)
	}

	return videoPart, nil
}

// GetVideos returns all video parts in the presentation.
func (p *Presentation) GetVideos() []parts.MediaPart {
	var videos []parts.MediaPart

	// Iterate through all slides
	for _, slide := range p.Slides() {
		// Get all parts from the slide
		for part := range slide.SlidePart().Parts() {
			if videoPart, ok := part.(parts.MediaPart); ok {
				// Check if this is a video (not audio)
				contentType := videoPart.GetContentType()
				if strings.HasPrefix(contentType, "video/") {
					videos = append(videos, videoPart)
				}
			}
		}
	}

	return videos
}

// RemoveVideo removes a video from the presentation.
func (p *Presentation) RemoveVideo(video parts.MediaPart) error {
	// Find which slide contains this video
	for _, slide := range p.Slides() {
		slidePart := slide.SlidePart()

		// Get the relationships for this slide
		rels := slidePart.Relationships()
		if rels == nil {
			continue
		}

		// Find the relationship that points to this video
		for _, rel := range rels.Relationships() {
			// Get the part for this relationship
			part, err := slidePart.GetPartById(rel.ID())
			if err != nil {
				continue
			}

			// Check if this is our video
			if part == video {
				// Remove the relationship
				return slidePart.DeletePart(rel.ID())
			}
		}
	}

	return fmt.Errorf("video not found in any slide")
}

// VideoInfo contains information about a video.
type VideoInfo struct {
	URI         string
	ContentType string
	Size        int64
	IsStreaming bool
	Format      string
	Supported   bool
}

// GetVideoInfo returns information about a video.
func GetVideoInfo(video parts.MediaPart) *VideoInfo {
	return &VideoInfo{
		URI:         video.URI(),
		ContentType: video.GetContentType(),
		Size:        video.GetSize(),
		IsStreaming: video.IsStreaming(),
		Format:      getFormatFromContentType(video.GetContentType()),
		Supported:   isFormatSupported(video.GetContentType()),
	}
}

// Helper function to get format name from content type.
func getFormatFromContentType(contentType string) string {
	switch contentType {
	case "video/mp4":
		return "MP4"
	case "video/x-msvideo", "video/avi":
		return "AVI"
	case "video/quicktime":
		return "MOV"
	case "video/x-ms-wmv":
		return "WMV"
	case "video/webm":
		return "WebM"
	case "video/x-matroska":
		return "MKV"
	case "video/ogg":
		return "OGV"
	default:
		return "Unknown"
	}
}

// Helper function to check if format is supported.
func isFormatSupported(contentType string) bool {
	// Based on our MediaRegistry registration
	supportedTypes := []string{
		"video/mp4",
		"video/x-msvideo",
		"video/avi",
		"video/quicktime",
		"video/x-ms-wmv",
	}

	for _, supported := range supportedTypes {
		if contentType == supported {
			return true
		}
	}
	return false
}
