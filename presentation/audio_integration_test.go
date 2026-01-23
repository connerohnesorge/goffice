//nolint:revive // line-length-limit: OOXML content types and relationship URIs are long strings
package presentation

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/connerohnesorge/goffice/presentation/parts"
)

// TestAddAudioFromFile tests adding an audio file from a file path.
func TestAddAudioFromFile(t *testing.T) {
	tempDir := t.TempDir()
	audioFile := filepath.Join(tempDir, "test.mp3")

	// Write minimal MP3 header (ID3 tag)
	audioData := make([]byte, 16)
	copy(audioData, []byte("ID3"))
	if err := os.WriteFile(audioFile, audioData, 0644); err != nil {
		t.Fatalf("Failed to create test audio file: %v", err)
	}

	// Create presentation
	pres, err := New(filepath.Join(tempDir, "test_audio.pptx"), DocTypePresentation)
	if err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}
	defer pres.Close()
	_, _ = pres.AddSlide()

	// Add audio from file
	audioPart, err := pres.AddAudioFromFile(0, audioFile, nil)
	if err != nil {
		t.Fatalf("AddAudioFromFile() error = %v", err)
	}

	if audioPart == nil {
		t.Fatal("AddAudioFromFile() returned nil")
	}

	if audioPart.GetContentType() != "audio/mpeg" {
		t.Errorf("ContentType = %v, want %v", audioPart.GetContentType(), "audio/mpeg")
	}

	data, err := audioPart.GetData()
	if err != nil {
		t.Fatalf("GetData() error = %v", err)
	}
	if !bytes.Equal(data, audioData) {
		t.Error("Audio data mismatch")
	}
}

// TestGetAudios tests retrieving all audio parts.
func TestGetAudios(t *testing.T) {
	pres, err := New(filepath.Join(t.TempDir(), "get_audios.pptx"), DocTypePresentation)
	if err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}
	defer pres.Close()
	_, _ = pres.AddSlide()

	// Add two audio files
	audioData := make([]byte, 16)
	copy(audioData, []byte("ID3"))
	
	_, err = pres.AddAudioFromFile(0, "", &AddAudioOptions{
		AutoDetectType: false,
		AudioType:      parts.AudioTypeMp3,
	})
	// Wait, AddAudioFromFile requires a real file.
	// I'll use a hack or just create real temp files.
	temp1 := filepath.Join(t.TempDir(), "1.mp3")
	temp2 := filepath.Join(t.TempDir(), "2.wav")
	os.WriteFile(temp1, audioData, 0644)
	copy(audioData, []byte("RIFFxxxxWAVE"))
	os.WriteFile(temp2, audioData, 0644)

	pres.AddAudioFromFile(0, temp1, nil)
	pres.AddAudioFromFile(0, temp2, nil)

	audios := pres.GetAudios()
	if len(audios) != 2 {
		t.Errorf("GetAudios() returned %d parts, want 2", len(audios))
	}
}

// TestAddAudioValidation tests audio format validation.
func TestAddAudioValidation(t *testing.T) {
	pres, _ := New(filepath.Join(t.TempDir(), "validation.pptx"), DocTypePresentation)
	_, _ = pres.AddSlide()

	tests := []struct {
		name        string
		filename    string
		data        []byte
		expectError bool
	}{
		{
			name:        "Valid MP3",
			filename:    "audio.mp3",
			data:        []byte("ID3xxxxxxxxxxxxx"),
			expectError: false,
		},
		{
			name:        "Unsupported format",
			filename:    "audio.ogg",
			data:        []byte("OggSxxxxxxxxxxxx"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &AddAudioOptions{
				ValidateFormat: true,
				AutoDetectType: true,
			}
			_, err := pres.AddAudioFromBytes(0, tt.data, tt.filename, opts)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
