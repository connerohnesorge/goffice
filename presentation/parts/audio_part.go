//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package parts

import (
	"bytes"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/connerohnesorge/goffice/openxml"
)

// AudioType represents the type of an audio file.
type AudioType int

const (
	// AudioTypeMp3 represents an MP3 audio file.
	AudioTypeMp3 AudioType = iota

	// AudioTypeWav represents a WAV audio file.
	AudioTypeWav

	// AudioTypeWma represents a WMA (Windows Media Audio) file.
	AudioTypeWma

	// AudioTypeM4a represents an M4A (MPEG-4 Audio) file.
	AudioTypeM4a

	// AudioTypeOgg represents an Ogg Vorbis audio file.
	AudioTypeOgg

	// AudioTypeFlac represents a FLAC audio file.
	AudioTypeFlac
)

// Audio extension constants.
const (
	extMP3   = ".mp3"
	extWAV   = ".wav"
	extWMA   = ".wma"
	extM4A   = ".m4a"
	extOGG   = ".ogg"
	extFLAC  = ".flac"
)

// Audio content type constants.
const (
	contentTypeAudioMPEG  = "audio/mpeg"
	contentTypeAudioWAV   = "audio/wav"
	contentTypeAudioWMA   = "audio/x-ms-wma"
	contentTypeAudioMP4   = "audio/mp4"
	contentTypeAudioOGG   = "audio/ogg"
	contentTypeAudioFLAC  = "audio/flac"
)

// String returns the string representation of the audio type.
func (at AudioType) String() string {
	switch at {
	case AudioTypeMp3:
		return "mp3"
	case AudioTypeWav:
		return "wav"
	case AudioTypeWma:
		return "wma"
	case AudioTypeM4a:
		return "m4a"
	case AudioTypeOgg:
		return "ogg"
	case AudioTypeFlac:
		return "flac"
	default:
		return "mp3"
	}
}

// Extension returns the file extension for this audio type.
func (at AudioType) Extension() string {
	switch at {
	case AudioTypeMp3:
		return extMP3
	case AudioTypeWav:
		return extWAV
	case AudioTypeWma:
		return extWMA
	case AudioTypeM4a:
		return extM4A
	case AudioTypeOgg:
		return extOGG
	case AudioTypeFlac:
		return extFLAC
	default:
		return extMP3
	}
}

// ContentType returns the MIME content type for this audio type.
func (at AudioType) ContentType() string {
	switch at {
	case AudioTypeMp3:
		return contentTypeAudioMPEG
	case AudioTypeWav:
		return contentTypeAudioWAV
	case AudioTypeWma:
		return contentTypeAudioWMA
	case AudioTypeM4a:
		return contentTypeAudioMP4
	case AudioTypeOgg:
		return contentTypeAudioOGG
	case AudioTypeFlac:
		return contentTypeAudioFLAC
	default:
		return contentTypeAudioMPEG
	}
}

// IsSupported returns true if this audio type is supported by PowerPoint.
func (at AudioType) IsSupported() bool {
	switch at {
	case AudioTypeMp3, AudioTypeWav, AudioTypeWma, AudioTypeM4a:
		return true // Native support
	case AudioTypeOgg, AudioTypeFlac:
		return false // Not natively supported
	default:
		return false
	}
}

// AudioPart represents an audio part (ppt/media/audio1.mp3, etc.).
type AudioPart struct {
	*baseMediaPart

	// audioType is the type of this audio.
	audioType AudioType

	// streamReader holds the streaming reader when using streaming mode.
	streamReader io.Reader

	// streamSize holds the size when streaming.
	streamSize int64
}

// Counter for generating unique audio filenames.
var audioUniqueCounter uint64

// NewAudioPartForSlide creates a new audio part for a slide.
func NewAudioPartForSlide(
	slidePart *SlidePart,
	audioType AudioType,
	streaming bool,
) (*AudioPart, error) {
	num := atomic.AddUint64(
		&audioUniqueCounter,
		1,
	)
	uri := fmt.Sprintf(
		"/ppt/media/audio%d%s",
		num,
		audioType.Extension(),
	)

	packPart, relID, err := slidePart.addChildPart(
		uri,
		audioType.ContentType(),
		RelationshipTypeAudio,
	)
	if err != nil {
		return nil, err
	}

	basePart := newBaseMediaPart(
		uri,
		audioType.ContentType(),
		packPart,
		slidePart,
		streaming,
	)

	ap := &AudioPart{
		baseMediaPart: basePart,
		audioType:     audioType,
		streamSize:    0,
	}

	ap.SetRelationshipID(relID)

	// Add to slide part's child parts
	if err := slidePart.AddPart(ap, relID); err != nil {
		return nil, err
	}

	return ap, nil
}

// AudioType returns the type of this audio.
func (ap *AudioPart) AudioType() AudioType {
	return ap.audioType
}

// SetStream sets the audio content from a reader.
func (ap *AudioPart) SetStream(r io.Reader) error {
	if ap.streaming {
		var size int64
		teeReader := io.TeeReader(r, &sizeCounter{size: &size})

		ap.streamReader = teeReader
		ap.streamSize = size
		ap.setSize(size)
		ap.MarkDirty()

		return nil
	}

	return ap.baseMediaPart.SetStream(r)
}

// GetStream returns a reader for the audio content.
func (ap *AudioPart) GetStream() io.Reader {
	if ap.streaming && ap.streamReader != nil {
		return ap.streamReader
	}

	return ap.baseMediaPart.GetStream()
}

// GetData returns the raw audio data.
func (ap *AudioPart) GetData() ([]byte, error) {
	if ap.streaming {
		if ap.streamReader == nil {
			return nil, nil
		}

		return io.ReadAll(ap.streamReader)
	}

	stream := ap.GetStream()
	if stream == nil {
		return nil, nil
	}

	return io.ReadAll(stream)
}

// FeedDataBytes sets the audio data from a byte slice.
func (ap *AudioPart) FeedDataBytes(data []byte) error {
	if ap.streaming {
		return ap.SetStream(bytes.NewReader(data))
	}

	ap.SetData(data)
	ap.setSize(int64(len(data)))

	return nil
}

// Ensure AudioPart implements MediaPart.
var _ MediaPart = (*AudioPart)(nil)

// AudioPartFactory creates an AudioPart from a URI and container.
func AudioPartFactory(
	uri string,
	container openxml.OpenXmlPartContainer,
) openxml.OpenXmlPart {
	packPart := container.GetPackagingPart(uri)
	if packPart == nil {
		return nil
	}

	contentType := packPart.ContentType()
	audioType := audioTypeFromContentType(contentType)

	streaming := ShouldStream(packPart.Size())

	basePart := newBaseMediaPart(
		uri,
		contentType,
		packPart,
		container,
		streaming,
	)

	return &AudioPart{
		baseMediaPart: basePart,
		audioType:     audioType,
	}
}

// audioTypeFromContentType determines the AudioType from a content type string.
func audioTypeFromContentType(
	contentType string,
) AudioType {
	switch contentType {
	case contentTypeAudioMPEG, "audio/mp3":
		return AudioTypeMp3
	case contentTypeAudioWAV, "audio/x-wav", "audio/wave":
		return AudioTypeWav
	case contentTypeAudioWMA:
		return AudioTypeWma
	case contentTypeAudioMP4, "audio/x-m4a":
		return AudioTypeM4a
	case contentTypeAudioOGG:
		return AudioTypeOgg
	case contentTypeAudioFLAC:
		return AudioTypeFlac
	default:
		return AudioTypeMp3
	}
}

// AudioTypeFromExtension determines the AudioType from a file extension.
func AudioTypeFromExtension(
	ext string,
) AudioType {
	switch ext {
	case extMP3:
		return AudioTypeMp3
	case extWAV:
		return AudioTypeWav
	case extWMA:
		return AudioTypeWma
	case extM4A:
		return AudioTypeM4a
	case extOGG:
		return AudioTypeOgg
	case extFLAC:
		return AudioTypeFlac
	default:
		return AudioTypeMp3
	}
}

// AudioTypeFromFilename determines the AudioType from a filename.
func AudioTypeFromFilename(
	filename string,
) AudioType {
	ext := getFileExtension(filename)

	return AudioTypeFromExtension(ext)
}

// AudioTypeFromMagicBytes determines the AudioType from magic bytes.
func AudioTypeFromMagicBytes(
	data []byte,
) (AudioType, bool) {
	if len(data) < 4 {
		return AudioTypeMp3, false
	}

	// Check RIFF/WAV
	if bytes.HasPrefix(data, []byte("RIFF")) {
		if len(data) >= 12 && string(data[8:12]) == "WAVE" {
			return AudioTypeWav, true
		}
	}

	// Check MP3 (ID3 or sync)
	if bytes.HasPrefix(data, []byte("ID3")) {
		return AudioTypeMp3, true
	}
	if data[0] == 0xFF && (data[1]&0xE0) == 0xE0 {
		return AudioTypeMp3, true
	}

	// Check FLAC
	if bytes.HasPrefix(data, []byte("fLaC")) {
		return AudioTypeFlac, true
	}

	// Check Ogg
	if bytes.HasPrefix(data, []byte("OggS")) {
		return AudioTypeOgg, true
	}

	return AudioTypeMp3, false
}

// DetectAudioType attempts to detect the audio type from data, falling back to extension.
func DetectAudioType(
	data []byte,
	filename string,
) AudioType {
	// First try magic bytes
	if audioType, detected := AudioTypeFromMagicBytes(data); detected {
		return audioType
	}

	// Fall back to extension
	if filename != "" {
		return AudioTypeFromFilename(filename)
	}

	// Default to MP3
	return AudioTypeMp3
}

// Register audio content types.
func init() {
	audioTypes := []AudioType{
		AudioTypeMp3,
		AudioTypeWav,
		AudioTypeWma,
		AudioTypeM4a,
		AudioTypeOgg,
		AudioTypeFlac,
	}

	for _, at := range audioTypes {
		openxml.RegisterPartType(
			&openxml.PartTypeInfo{
				ContentType:      at.ContentType(),
				RelationshipType: RelationshipTypeAudio,
				Factory:          AudioPartFactory,
				DefaultURI: fmt.Sprintf(
					"/ppt/media/audio1%s",
					at.Extension(),
				),
				IsFixedContentType: true,
			},
		)
	}
}
