//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"bytes"
	"fmt"
	"iter"
	"testing"

	"github.com/connerohnesorge/goffice/openxml"
	"github.com/connerohnesorge/goffice/openxml/features"
	"github.com/connerohnesorge/goffice/packaging"
)

// TestVideoType tests video type detection and properties.
func TestVideoType(t *testing.T) {
	tests := []struct {
		name        string
		videoType   VideoType
		expectedExt string
		expectedCT  string
		supported   bool
	}{
		{
			name:        "MP4 video",
			videoType:   VideoTypeMp4,
			expectedExt: ".mp4",
			expectedCT:  "video/mp4",
			supported:   true,
		},
		{
			name:        "AVI video",
			videoType:   VideoTypeAvi,
			expectedExt: ".avi",
			expectedCT:  "video/x-msvideo",
			supported:   true,
		},
		{
			name:        "MOV video",
			videoType:   VideoTypeMov,
			expectedExt: ".mov",
			expectedCT:  "video/quicktime",
			supported:   true,
		},
		{
			name:        "WMV video",
			videoType:   VideoTypeWmv,
			expectedExt: ".wmv",
			expectedCT:  "video/x-ms-wmv",
			supported:   true,
		},
		{
			name:        "WebM video",
			videoType:   VideoTypeWebm,
			expectedExt: ".webm",
			expectedCT:  "video/webm",
			supported:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.videoType.Extension(); got != tt.expectedExt {
				t.Errorf("Extension() = %v, want %v", got, tt.expectedExt)
			}
			if got := tt.videoType.ContentType(); got != tt.expectedCT {
				t.Errorf("ContentType() = %v, want %v", got, tt.expectedCT)
			}
			if got := tt.videoType.IsSupported(); got != tt.supported {
				t.Errorf("IsSupported() = %v, want %v", got, tt.supported)
			}
		})
	}
}

// TestVideoTypeFromExtension tests video type detection from file extension.
func TestVideoTypeFromExtension(t *testing.T) {
	tests := []struct {
		ext      string
		expected VideoType
	}{
		{".mp4", VideoTypeMp4},
		{".avi", VideoTypeAvi},
		{".mov", VideoTypeMov},
		{".wmv", VideoTypeWmv},
		{".webm", VideoTypeWebm},
		{".mkv", VideoTypeMkv},
		{".ogv", VideoTypeOgv},
		{"mp4", VideoTypeMp4},  // Without dot
		{".MP4", VideoTypeMp4}, // Case insensitive
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := VideoTypeFromExtension(tt.ext); got != tt.expected {
				t.Errorf("VideoTypeFromExtension(%q) = %v, want %v", tt.ext, got, tt.expected)
			}
		})
	}
}

// TestVideoTypeFromMagicBytes tests video type detection from file content.
func TestVideoTypeFromMagicBytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected VideoType
		found    bool
	}{
		{
			name:     "MP4 with ftyp",
			data:     append(make([]byte, 4), append([]byte("ftyp"), make([]byte, 8)...)...), // "ftyp" at offset 4
			expected: VideoTypeMp4,
			found:    true,
		},
		{
			name:     "AVI",
			data:     append([]byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x41, 0x56, 0x49, 0x20}, make([]byte, 4)...), // "RIFF" + "AVI "
			expected: VideoTypeAvi,
			found:    true,
		},
		{
			name:     "WMV",
			data:     append([]byte{0x30, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11}, make([]byte, 8)...), // ASF header
			expected: VideoTypeWmv,
			found:    true,
		},
		{
			name:     "Too short",
			data:     []byte{0x00, 0x00},
			expected: VideoTypeMp4,
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := VideoTypeFromMagicBytes(tt.data)
			if got != tt.expected {
				t.Errorf("VideoTypeFromMagicBytes() type = %v, want %v", got, tt.expected)
			}
			if found != tt.found {
				t.Errorf("VideoTypeFromMagicBytes() found = %v, want %v", found, tt.found)
			}
		})
	}
}

// TestDetectVideoType tests the combined detection method.
func TestDetectVideoType(t *testing.T) {
	// Test with magic bytes
	aviData := append([]byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x41, 0x56, 0x49, 0x20}, make([]byte, 4)...)
	if got := DetectVideoType(aviData, ""); got != VideoTypeAvi {
		t.Errorf("DetectVideoType() with magic = %v, want %v", got, VideoTypeAvi)
	}

	// Test with filename fallback
	if got := DetectVideoType(nil, "video.mp4"); got != VideoTypeMp4 {
		t.Errorf("DetectVideoType() with filename = %v, want %v", got, VideoTypeMp4)
	}

	// Test with both (magic takes precedence)
	if got := DetectVideoType(aviData, "video.mp4"); got != VideoTypeAvi {
		t.Errorf("DetectVideoType() with both = %v, want %v", got, VideoTypeAvi)
	}

	// Test default
	if got := DetectVideoType(nil, ""); got != VideoTypeMp4 {
		t.Errorf("DetectVideoType() default = %v, want %v", got, VideoTypeMp4)
	}
}

// TestVideoPartBasic tests basic VideoPart functionality without needing a slide.
func TestVideoPartBasic(t *testing.T) {
	// Create a packaging part for the video
	pkg, _ := packaging.Create("test.pptx")
	videoPackPart, err := pkg.CreatePart("/ppt/media/video1.mp4", "video/mp4")
	if err != nil {
		t.Fatalf("Failed to create packaging part: %v", err)
	}

	// Create a mock container
	container := &mockContainer{
		packPart: videoPackPart,
	}

	// Create video part using factory
	partIface := VideoPartFactory("/ppt/media/video1.mp4", container)
	if partIface == nil {
		t.Fatal("VideoPartFactory returned nil")
	}

	videoPart, ok := partIface.(*VideoPart)
	if !ok {
		t.Fatalf("VideoPartFactory returned %T, want *VideoPart", partIface)
	}

	// Check properties
	if videoPart.VideoType() != VideoTypeMp4 {
		t.Errorf("VideoType() = %v, want %v", videoPart.VideoType(), VideoTypeMp4)
	}
	if videoPart.GetContentType() != "video/mp4" {
		t.Errorf("GetContentType() = %v, want %v", videoPart.GetContentType(), "video/mp4")
	}
	if videoPart.IsStreaming() {
		t.Error("IsStreaming() = true, want false")
	}

	// Test setting data
	testData := []byte("test video data")
	videoPart.SetData(testData)
	videoPart.setSize(int64(len(testData)))

	// Verify data
	gotData, _ := videoPart.GetData()
	if !bytes.Equal(gotData, testData) {
		t.Errorf("GetData() = %v, want %v", gotData, testData)
	}

	// Verify size
	if videoPart.GetSize() != int64(len(testData)) {
		t.Errorf("GetSize() = %v, want %v", videoPart.GetSize(), len(testData))
	}
}

// TestVideoPartStreaming tests streaming video part.
func TestVideoPartStreaming(t *testing.T) {
	// Create a packaging part for the video
	pkg, _ := packaging.Create("test2.pptx")
	videoPackPart, err := pkg.CreatePart("/ppt/media/video1.mp4", "video/mp4")
	if err != nil {
		t.Fatalf("Failed to create packaging part: %v", err)
	}

	// Create a mock container
	container := &mockContainer{
		packPart: videoPackPart,
	}

	// Create video part using factory with streaming enabled
	videoPart := &VideoPart{
		baseMediaPart: &baseMediaPart{
			OpenXmlPartData: openxml.NewOpenXmlPartData(
				"/ppt/media/video1.mp4",
				"video/mp4",
				videoPackPart,
				container,
			),
			streaming: true,
			size:      0,
		},
		videoType: VideoTypeMp4,
	}

	// Check streaming flag
	if !videoPart.IsStreaming() {
		t.Error("IsStreaming() = false, want true")
	}

	// Test setting stream
	testData := []byte("streaming test data")
	reader := bytes.NewReader(testData)
	if err := videoPart.SetStream(reader); err != nil {
		t.Fatalf("SetStream() error = %v", err)
	}

	// Verify we can get the data back
	gotData, err := videoPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}
	if !bytes.Equal(gotData, testData) {
		t.Errorf("GetData() = %v, want %v", gotData, testData)
	}
}

// TestShouldStream tests the streaming threshold logic.
func TestShouldStream(t *testing.T) {
	tests := []struct {
		size     int64
		expected bool
	}{
		{1024 * 1024, false},       // 1MB - no streaming
		{50 * 1024 * 1024, false},  // 50MB - no streaming
		{99 * 1024 * 1024, false},  // 99MB - no streaming
		{100 * 1024 * 1024, false}, // Exactly 100MB - no streaming
		{100*1024*1024 + 1, true},  // 100MB + 1 byte - streaming
		{200 * 1024 * 1024, true},  // 200MB - streaming
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("size_%d", tt.size), func(t *testing.T) {
			if got := ShouldStream(tt.size); got != tt.expected {
				t.Errorf("ShouldStream(%d) = %v, want %v", tt.size, got, tt.expected)
			}
		})
	}
}

// TestVideoPartFactory tests the factory function.
func TestVideoPartFactory(t *testing.T) {
	// Create a test package with a video part
	pkg, _ := packaging.Create("test3.pptx")

	// Create the video part in the package
	videoURI := "/ppt/media/video1.mp4"
	videoPart, err := pkg.CreatePart(videoURI, "video/mp4")
	if err != nil {
		t.Fatalf("Failed to create packaging part: %v", err)
	}

	// Create a mock container
	container := &mockContainer{
		packPart: videoPart,
	}

	// Test factory
	result := VideoPartFactory(videoURI, container)
	if result == nil {
		t.Fatal("VideoPartFactory() returned nil")
	}

	vp, ok := result.(*VideoPart)
	if !ok {
		t.Fatalf("VideoPartFactory() returned %T, want *VideoPart", result)
	}

	if vp.VideoType() != VideoTypeMp4 {
		t.Errorf("VideoType() = %v, want %v", vp.VideoType(), VideoTypeMp4)
	}
}

// mockContainer implements OpenXmlPartContainer for testing.
type mockContainer struct {
	packPart *packaging.Part
}

func (m *mockContainer) Parts() iter.Seq[openxml.OpenXmlPart] {
	return func(yield func(openxml.OpenXmlPart) bool) {}
}

func (m *mockContainer) GetPartById(id string) (openxml.OpenXmlPart, error) {
	return nil, openxml.ErrPartNotFound
}

func (m *mockContainer) GetPartsOfType(contentType string) iter.Seq[openxml.OpenXmlPart] {
	return func(yield func(openxml.OpenXmlPart) bool) {}
}

func (m *mockContainer) AddPart(part openxml.OpenXmlPart, id string) error {
	return nil
}

func (m *mockContainer) DeletePart(id string) error {
	return nil
}

func (m *mockContainer) Features() *features.FeatureCollection {
	return features.NewFeatureCollection()
}

func (m *mockContainer) URI() string {
	return "/ppt/slides/slide1.xml"
}

func (m *mockContainer) Package() *packaging.Package {
	return nil
}

func (m *mockContainer) GetPackagingPart(uri string) *packaging.Part {
	if uri == "/ppt/media/video1.mp4" {
		return m.packPart
	}
	return nil
}
