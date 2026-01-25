package presentation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFluentMediaAPI(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "fluent.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer pres.Close()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	// Create a dummy video file
	videoFile := filepath.Join(tempDir, "test.mp4")
	videoData := make([]byte, 16)
	copy(videoData[4:], []byte("ftyp"))
	if err := os.WriteFile(videoFile, videoData, 0644); err != nil {
		t.Fatalf("failed to create test video file: %v", err)
	}

	// Use fluent API for video
	video, err := pres.AddVideo(0, videoFile)
	if err != nil {
		t.Fatalf("AddVideo failed: %v", err)
	}

	video.SetPosition(100, 100).
		SetSize(200, 200).
		SetAutoStart(true).
		SetLoop(true).
		SetMuted(true).
		SetVolume(50000)

	// Create a dummy audio file
	audioFile := filepath.Join(tempDir, "test.mp3")
	audioData := make([]byte, 16)
	copy(audioData, []byte("ID3"))
	if err := os.WriteFile(audioFile, audioData, 0644); err != nil {
		t.Fatalf("failed to create test audio file: %v", err)
	}

	// Use fluent API for audio
	audio, err := pres.AddAudio(0, audioFile)
	if err != nil {
		t.Fatalf("AddAudio failed: %v", err)
	}

	audio.SetPosition(400, 400).
		SetSize(50, 50).
		SetLoop(true).
		SetMuted(false).
		SetVolume(100000)

	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}
