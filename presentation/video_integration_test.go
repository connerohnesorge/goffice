//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package presentation

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/packaging"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// TestAddVideoFromFile tests adding a video from a file.
func TestAddVideoFromFile(t *testing.T) {
	// Create a temporary video file
	tempDir := t.TempDir()
	videoFile := filepath.Join(tempDir, "test.mp4")

	// Write test video data (simulated MP4 with ftyp header)
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70} // MP4 header
	if err := os.WriteFile(videoFile, videoData, 0644); err != nil {
		t.Fatalf("Failed to create test video file: %v", err)
	}

	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Add video from file
	videoPart, err := pres.AddVideoFromFile(0, videoFile, nil)
	if err != nil {
		t.Fatalf("AddVideoFromFile() error = %v", err)
	}

	// Verify video was added
	if videoPart == nil {
		t.Fatal("AddVideoFromFile() returned nil video part")
	}

	// Check video properties
	if videoPart.GetContentType() != "video/mp4" {
		t.Errorf("ContentType = %v, want %v", videoPart.GetContentType(), "video/mp4")
	}

	// Verify data
	data, err := videoPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}
	if !bytes.Equal(data, videoData) {
		t.Errorf("Video data mismatch")
	}

	// Verify size
	if videoPart.GetSize() != int64(len(videoData)) {
		t.Errorf("Size = %v, want %v", videoPart.GetSize(), len(videoData))
	}
}

// TestAddVideoFromReader tests adding a video from an io.Reader.
func TestAddVideoFromReader(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Create test video data
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70} // MP4 header
	reader := bytes.NewReader(videoData)

	// Add video from reader
	videoPart, err := pres.AddVideoFromReader(0, reader, "video.mp4", nil)
	if err != nil {
		t.Fatalf("AddVideoFromReader() error = %v", err)
	}

	// Verify video was added
	if videoPart == nil {
		t.Fatal("AddVideoFromReader() returned nil video part")
	}

	// Verify data
	gotData, err := videoPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}
	if !bytes.Equal(gotData, videoData) {
		t.Errorf("Video data mismatch")
	}
}

// TestAddVideoFromBytes tests adding a video from bytes.
func TestAddVideoFromBytes(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Create test video data
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70} // MP4 header

	// Add video from bytes
	videoPart, err := pres.AddVideoFromBytes(0, videoData, "video.mp4", nil)
	if err != nil {
		t.Fatalf("AddVideoFromBytes() error = %v", err)
	}

	// Verify video was added
	if videoPart == nil {
		t.Fatal("AddVideoFromBytes() returned nil video part")
	}

	// Verify data
	gotData, err := videoPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}
	if !bytes.Equal(gotData, videoData) {
		t.Errorf("Video data mismatch")
	}
}

// TestAddVideoValidation tests video format validation.
func TestAddVideoValidation(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	tests := []struct {
		name        string
		filename    string
		data        []byte
		expectError bool
	}{
		{
			name:        "Valid MP4",
			filename:    "video.mp4",
			data:        []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70},
			expectError: false,
		},
		{
			name:        "Unsupported format",
			filename:    "video.webm",
			data:        []byte{0x1A, 0x45, 0xDF, 0xA3}, // WebM header
			expectError: true,
		},
		{
			name:        "Invalid extension",
			filename:    "video.xyz",
			data:        []byte{0x00, 0x00, 0x00, 0x20},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &AddVideoOptions{
				ValidateFormat: true,
				AutoDetectType: true,
			}

			_, err := pres.AddVideoFromBytes(0, tt.data, tt.filename, opts)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestGetVideos tests retrieving all videos from presentation.
func TestGetVideos(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)

	// Add multiple slides with videos
	slide1 := pres.AddSlide()
	slide2 := pres.AddSlide()

	// Add videos
	videoData1 := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70} // MP4
	videoData2 := []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x41, 0x56, 0x49, 0x20} // AVI

	video1, err := pres.AddVideoFromBytes(0, videoData1, "video1.mp4", nil)
	if err != nil {
		t.Fatalf("Failed to add video1: %v", err)
	}

	video2, err := pres.AddVideoFromBytes(1, videoData2, "video2.avi", nil)
	if err != nil {
		t.Fatalf("Failed to add video2: %v", err)
	}

	// Get all videos
	videos := pres.GetVideos()
	if len(videos) != 2 {
		t.Errorf("GetVideos() returned %d videos, want 2", len(videos))
	}

	// Verify we got the right videos
	foundVideo1, foundVideo2 := false, false
	for _, v := range videos {
		if v == video1 {
			foundVideo1 = true
		}
		if v == video2 {
			foundVideo2 = true
		}
	}
	if !foundVideo1 {
		t.Error("video1 not found in GetVideos() result")
	}
	if !foundVideo2 {
		t.Error("video2 not found in GetVideos() result")
	}
}

// TestRemoveVideo tests removing a video from presentation.
func TestRemoveVideo(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Add a video
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	video, err := pres.AddVideoFromBytes(0, videoData, "video.mp4", nil)
	if err != nil {
		t.Fatalf("Failed to add video: %v", err)
	}

	// Verify video exists
	videos := pres.GetVideos()
	if len(videos) != 1 {
		t.Fatalf("Expected 1 video, got %d", len(videos))
	}

	// Remove the video
	if err := pres.RemoveVideo(video); err != nil {
		t.Fatalf("RemoveVideo() error = %v", err)
	}

	// Verify video is removed
	videos = pres.GetVideos()
	if len(videos) != 0 {
		t.Errorf("Expected 0 videos after removal, got %d", len(videos))
	}
}

// TestStreamingVideo tests streaming video functionality.
func TestStreamingVideo(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Create a large video data (simulating a large file)
	largeVideoData := make([]byte, 150*1024*1024) // 150MB
	copy(largeVideoData, []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}) // MP4 header

	// Add video with streaming threshold below file size
	opts := &AddVideoOptions{
		StreamingThreshold: 100 * 1024 * 1024, // 100MB
	}

	// Use a reader to simulate streaming
	reader := bytes.NewReader(largeVideoData)
	videoPart, err := pres.AddVideoFromReader(0, reader, "large-video.mp4", opts)
	if err != nil {
		t.Fatalf("AddVideoFromReader() error = %v", err)
	}

	// Verify it's a streaming part
	if !videoPart.IsStreaming() {
		t.Error("Expected streaming video part")
	}

	// Verify size
	if videoPart.GetSize() != int64(len(largeVideoData)) {
		t.Errorf("Size = %v, want %v", videoPart.GetSize(), len(largeVideoData))
	}
}

// TestGetVideoInfo tests getting video information.
func TestGetVideoInfo(t *testing.T) {
	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Add a video
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	video, err := pres.AddVideoFromBytes(0, videoData, "video.mp4", nil)
	if err != nil {
		t.Fatalf("Failed to add video: %v", err)
	}

	// Get video info
	info := GetVideoInfo(video)

	// Verify info
	if info.URI != video.URI() {
		t.Errorf("URI = %v, want %v", info.URI, video.URI())
	}
	if info.ContentType != "video/mp4" {
		t.Errorf("ContentType = %v, want %v", info.ContentType, "video/mp4")
	}
	if info.Size != int64(len(videoData)) {
		t.Errorf("Size = %v, want %v", info.Size, len(videoData))
	}
	if info.IsStreaming {
		t.Error("IsStreaming should be false")
	}
	if info.Format != "MP4" {
		t.Errorf("Format = %v, want %v", info.Format, "MP4")
	}
	if !info.Supported {
		t.Error("Supported should be true")
	}
}

// TestVideoRoundTrip tests saving and loading presentation with videos.
func TestVideoRoundTrip(t *testing.T) {
	// Create presentation with video
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Add a video
	videoData := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	video, err := pres.AddVideoFromBytes(0, videoData, "video.mp4", nil)
	if err != nil {
		t.Fatalf("Failed to add video: %v", err)
	}

	// Save presentation
	tempFile := filepath.Join(t.TempDir(), "test.pptx")
	if err := pkg.Save(tempFile); err != nil {
		t.Fatalf("Failed to save presentation: %v", err)
	}

	// Load presentation
	loadedPkg, err := packaging.OpenPackage(tempFile)
	if err != nil {
		t.Fatalf("Failed to load presentation: %v", err)
	}

	loadedPres := NewPresentation(loadedPkg)

	// Verify video is preserved
	videos := loadedPres.GetVideos()
	if len(videos) != 1 {
		t.Fatalf("Expected 1 video after loading, got %d", len(videos))
	}

	// Verify video content
	loadedData, err := videos[0].GetData()
	if err != nil {
		t.Fatalf("Failed to get loaded video data: %v", err)
	}
	if !bytes.Equal(loadedData, videoData) {
		t.Error("Video data mismatch after round trip")
	}
}

// TestMediaRegistry tests the media format registry.
func TestMediaRegistry(t *testing.T) {
	// Test video format detection
	mp4Data := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70}
	format, err := DetectMediaFormat(mp4Data, "video.mp4")
	if err != nil {
		t.Fatalf("DetectMediaFormat() error = %v", err)
	}
	if format.Type != MediaTypeVideo {
		t.Errorf("Format type = %v, want %v", format.Type, MediaTypeVideo)
	}

	// Test validation
	if err := ValidateMedia(mp4Data, "video.mp4"); err != nil {
		t.Errorf("ValidateMedia() error = %v", err)
	}

	// Test unsupported format
	webmData := []byte{0x1A, 0x45, 0xDF, 0xA3}
	err = ValidateMedia(webmData, "video.webm")
	if err == nil {
		t.Error("Expected validation error for unsupported format")
	}

	// Test supported formats
	videoFormats := GetSupportedVideoFormats()
	if len(videoFormats) == 0 {
		t.Error("No supported video formats found")
	}

	audioFormats := GetSupportedAudioFormats()
	if len(audioFormats) == 0 {
		t.Error("No supported audio formats found")
	}
}

// TestStreamingPart tests the streaming part implementation.
func TestStreamingPart(t *testing.T) {
	// Create test data
	testData := []byte("streaming test data")

	// Create presentation
	pkg := packaging.NewPackage()
	pres := NewPresentation(pkg)
	slide := pres.AddSlide()

	// Create streaming video part
	videoPart, err := parts.NewStreamingVideoPartForSlide(slide.SlidePart(), parts.VideoTypeMp4)
	if err != nil {
		t.Fatalf("NewStreamingVideoPartForSlide() error = %v", err)
	}

	// Set stream
	reader := bytes.NewReader(testData)
	if err := videoPart.SetStream(reader); err != nil {
		t.Fatalf("SetStream() error = %v", err)
	}

	// Get stream and verify data
	gotStream := videoPart.GetStream()
	if gotStream == nil {
		t.Fatal("GetStream() returned nil")
	}

	gotData, err := io.ReadAll(gotStream)
	if err != nil {
		t.Fatalf("Failed to read from stream: %v", err)
	}
	if !bytes.Equal(gotData, testData) {
		t.Errorf("Stream data mismatch")
	}

	// Verify size
	if videoPart.GetSize() != int64(len(testData)) {
		t.Errorf("Size = %v, want %v", videoPart.GetSize(), len(testData))
	}

	// Close streaming part
	if err := videoPart.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

// Test helper to create mock reader with specific size.
type sizeReader struct {
	size int
	pos  int
}

func (r *sizeReader) Read(p []byte) (n int, err error) {
	if r.pos >= r.size {
		return 0, io.EOF
	}
	n = len(p)
	if r.pos+n > r.size {
		n = r.size - r.pos
	}
	for i := 0; i < n; i++ {
		p[i] = byte(r.pos + i)
	}
	r.pos += n
	return n, nil
}

func (r *sizeReader) Size() int64 {
	return int64(r.size)
}

// TestShouldUseStreamingForSize tests the streaming decision logic.
func TestShouldUseStreamingForSize(t *testing.T) {
	tests := []struct {
		size     int64
		expected bool
	}{
		{50 * 1024 * 1024, false},  // 50MB
		{99 * 1024 * 1024, false},  // 99MB
		{100 * 1024 * 1024, false}, // Exactly 100MB
		{101 * 1024 * 1024, true},  // 101MB
		{200 * 1024 * 1024, true},  // 200MB
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			if got := shouldUseStreamingForSize(tt.size); got != tt.expected {
				t.Errorf("shouldUseStreamingForSize(%d) = %v, want %v", tt.size, got, tt.expected)
			}
		})
	}
}
