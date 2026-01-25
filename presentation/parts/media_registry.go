//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"fmt"
	"strings"
	"sync"
)

// MediaType represents the type of media (video or audio).
type MediaType int

const (
	// MediaTypeVideo represents video media.
	MediaTypeVideo MediaType = iota
	// MediaTypeAudio represents audio media.
	MediaTypeAudio
)

// MediaFormat represents a registered media format.
type MediaFormat struct {
	// Type is whether this is video or audio.
	Type MediaType

	// Extensions that this format supports.
	Extensions []string

	// ContentTypes that this format supports.
	ContentTypes []string

	// Magic bytes for format detection.
	MagicBytes [][]byte

	// MaxSize is the maximum file size in bytes (0 for no limit).
	MaxSize int64

	// RequiresCodec indicates if this format may require additional codecs.
	RequiresCodec bool

	// SupportedInPowerPoint indicates native PowerPoint support.
	SupportedInPowerPoint bool
}

// MediaRegistry manages registration and detection of media formats.
type MediaRegistry struct {
	mu      sync.RWMutex
	formats map[string]*MediaFormat
}

// Global media registry instance.
var globalMediaRegistry = &MediaRegistry{
	formats: make(map[string]*MediaFormat),
}

// RegisterFormat registers a media format.
func (mr *MediaRegistry) RegisterFormat(name string, format *MediaFormat) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	mr.formats[name] = format
}

// DetectFormat detects the media format from content and filename.
func (mr *MediaRegistry) DetectFormat(
	data []byte,
	filename string,
) (*MediaFormat, error) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()

	// First try magic bytes
	for _, format := range mr.formats {
		if len(data) > 0 && len(format.MagicBytes) > 0 {
			for _, magic := range format.MagicBytes {
				if len(data) >= len(magic) &&
					equalBytes(data[:len(magic)], magic) {
					return format, nil
				}
			}
		}
	}

	// Fall back to extension
	ext := strings.ToLower(getFileExtension(filename))
	if ext != "" {
		for _, format := range mr.formats {
			for _, formatExt := range format.Extensions {
				if formatExt == ext {
					return format, nil
				}
			}
		}
	}

	// Fall back to content type if available
	contentType := getContentTypeFromFilename(filename)
	if contentType != "" {
		for _, format := range mr.formats {
			for _, formatType := range format.ContentTypes {
				if formatType == contentType {
					return format, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("unknown media format for file: %s", filename)
}

// ValidateMedia validates media content against format constraints.
func (mr *MediaRegistry) ValidateMedia(
	data []byte,
	filename string,
) error {
	format, err := mr.DetectFormat(data, filename)
	if err != nil {
		return err
	}

	// Check size constraint
	if format.MaxSize > 0 && int64(len(data)) > format.MaxSize {
		return fmt.Errorf("media file exceeds maximum size of %d bytes", format.MaxSize)
	}

	// Check PowerPoint support
	if !format.SupportedInPowerPoint {
		return fmt.Errorf("media format %s is not natively supported in PowerPoint", filename)
	}

	return nil
}

// GetSupportedFormats returns all supported formats for a media type.
func (mr *MediaRegistry) GetSupportedFormats(
	mediaType MediaType,
) []*MediaFormat {
	mr.mu.RLock()
	defer mr.mu.RUnlock()

	var result []*MediaFormat
	for _, format := range mr.formats {
		if format.Type == mediaType && format.SupportedInPowerPoint {
			result = append(result, format)
		}
	}
	return result
}

// equalBytes compares two byte slices for equality.
func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// getContentTypeFromFilename attempts to determine content type from filename.
func getContentTypeFromFilename(filename string) string {
	ext := strings.ToLower(getFileExtension(filename))
	switch ext {
	case ".mp4":
		return "video/mp4"
	case ".avi":
		return "video/x-msvideo"
	case ".mov":
		return "video/quicktime"
	case ".wmv":
		return "video/x-ms-wmv"
	case ".webm":
		return "video/webm"
	case ".mkv":
		return "video/x-matroska"
	case ".ogv":
		return "video/ogg"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".wma":
		return "audio/x-ms-wma"
	case ".m4a":
		return "audio/mp4"
	case ".ogg":
		return "audio/ogg"
	case ".flac":
		return "audio/flac"
	default:
		return ""
	}
}

// RegisterVideoFormats registers all supported video formats.
func RegisterVideoFormats() {
	// MP4 format
	globalMediaRegistry.RegisterFormat("mp4", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".mp4", ".m4v"},
		ContentTypes:          []string{"video/mp4", "video/x-m4v"},
		MagicBytes:            mp4Magic,
		MaxSize:               0, // No limit
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// WMV format
	globalMediaRegistry.RegisterFormat("wmv", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".wmv", ".asf"},
		ContentTypes:          []string{"video/x-ms-wmv", "video/x-ms-asf"},
		MagicBytes:            [][]byte{wmvMagic},
		MaxSize:               0,
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// AVI format
	globalMediaRegistry.RegisterFormat("avi", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".avi"},
		ContentTypes:          []string{"video/x-msvideo", "video/avi"},
		MagicBytes:            [][]byte{aviMagic},
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: true,
	})

	// MOV format
	globalMediaRegistry.RegisterFormat("mov", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".mov", ".qt"},
		ContentTypes:          []string{"video/quicktime"},
		MagicBytes:            mp4Magic, // MOV uses similar structure
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: true,
	})

	// WebM format (not natively supported)
	globalMediaRegistry.RegisterFormat("webm", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".webm"},
		ContentTypes:          []string{"video/webm"},
		MagicBytes:            [][]byte{webmMagic},
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: false,
	})

	// MKV format (not natively supported)
	globalMediaRegistry.RegisterFormat("mkv", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".mkv"},
		ContentTypes:          []string{"video/x-matroska"},
		MagicBytes:            [][]byte{mkvMagic},
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: false,
	})

	// OGV format (not natively supported)
	globalMediaRegistry.RegisterFormat("ogv", &MediaFormat{
		Type:                  MediaTypeVideo,
		Extensions:            []string{".ogv", ".ogg"},
		ContentTypes:          []string{"video/ogg", "video/theora"},
		MagicBytes:            [][]byte{ogvMagic},
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: false,
	})
}

// RegisterAudioFormats registers all supported audio formats.
func RegisterAudioFormats() {
	// MP3 format
	globalMediaRegistry.RegisterFormat("mp3", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".mp3"},
		ContentTypes:          []string{"audio/mpeg", "audio/mp3"},
		MagicBytes:            [][]byte{{0xFF, 0xFB}, {0xFF, 0xF3}, {0xFF, 0xF2}, {0x49, 0x44, 0x33}}, // MPEG sync or ID3
		MaxSize:               0,
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// WAV format
	globalMediaRegistry.RegisterFormat("wav", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".wav", ".wave"},
		ContentTypes:          []string{"audio/wav", "audio/x-wav", "audio/wave"},
		MagicBytes:            [][]byte{{0x52, 0x49, 0x46, 0x46}}, // "RIFF"
		MaxSize:               0,
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// WMA format
	globalMediaRegistry.RegisterFormat("wma", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".wma"},
		ContentTypes:          []string{"audio/x-ms-wma"},
		MagicBytes:            [][]byte{wmvMagic}, // Same as WMV
		MaxSize:               0,
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// M4A format
	globalMediaRegistry.RegisterFormat("m4a", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".m4a", ".m4p"},
		ContentTypes:          []string{"audio/mp4", "audio/x-m4a"},
		MagicBytes:            mp4Magic, // Same as MP4
		MaxSize:               0,
		RequiresCodec:         false,
		SupportedInPowerPoint: true,
	})

	// OGG format
	globalMediaRegistry.RegisterFormat("ogg", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".ogg", ".oga"},
		ContentTypes:          []string{"audio/ogg", "audio/vorbis"},
		MagicBytes:            [][]byte{ogvMagic}, // Same as OGV
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: false,
	})

	// FLAC format
	globalMediaRegistry.RegisterFormat("flac", &MediaFormat{
		Type:                  MediaTypeAudio,
		Extensions:            []string{".flac"},
		ContentTypes:          []string{"audio/flac", "audio/x-flac"},
		MagicBytes:            [][]byte{{0x66, 0x4C, 0x61, 0x43}}, // "fLaC"
		MaxSize:               0,
		RequiresCodec:         true,
		SupportedInPowerPoint: false,
	})
}

// Initialize media formats.
func init() {
	RegisterVideoFormats()
	RegisterAudioFormats()
}

// Helper functions for accessing the global registry.

// DetectMediaFormat detects the media format from content and filename.
func DetectMediaFormat(data []byte, filename string) (*MediaFormat, error) {
	return globalMediaRegistry.DetectFormat(data, filename)
}

// ValidateMedia validates media content against format constraints.
func ValidateMedia(data []byte, filename string) error {
	return globalMediaRegistry.ValidateMedia(data, filename)
}

// GetSupportedVideoFormats returns all supported video formats.
func GetSupportedVideoFormats() []*MediaFormat {
	return globalMediaRegistry.GetSupportedFormats(MediaTypeVideo)
}

// GetSupportedAudioFormats returns all supported audio formats.
func GetSupportedAudioFormats() []*MediaFormat {
	return globalMediaRegistry.GetSupportedFormats(MediaTypeAudio)
}
