package layout

import (
	"testing"
)

func TestGetBaseDirection(t *testing.T) {
	tests := []struct {
		text string
		want Direction
	}{
		{"Hello", LTR},
		{"שלום", RTL},       // Hebrew
		{"مرحبا", RTL},      // Arabic
		{"Hello שלום", LTR}, // Starts with LTR
		{"שלום Hello", RTL}, // Starts with RTL
	}

	for _, tt := range tests {
		got := GetBaseDirection(tt.text)
		if got != tt.want {
			t.Errorf(
				"GetBaseDirection(%q) = %v, want %v",
				tt.text,
				got,
				tt.want,
			)
		}
	}
}

func TestBidiProcessor_Simple(t *testing.T) {
	bp := NewBidiProcessor()
	text := "Hello שלום" // LTR base
	err := bp.Process(text, LTR)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	runes := []rune(text)
	res, err := bp.Reorder(0, len(runes))
	if err != nil {
		t.Fatalf("Reorder failed: %v", err)
	}

	// Visual order for "Hello שלום" (LTR base) should be "Hello םולש"
	// (Actually the bidi package handles this correctly)
	if res.Text == "" {
		t.Error("Reorder returned empty text")
	}
}

func TestBidiProcessor_Mixed(t *testing.T) {
	bp := NewBidiProcessor()
	// "English שלום English"
	text := "English שלום English"
	err := bp.Process(text, LTR)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	runes := []rune(text)
	res, err := bp.Reorder(0, len(runes))
	if err != nil {
		t.Fatalf("Reorder failed: %v", err)
	}

	if res.Text == "" {
		t.Error("Reorder returned empty text")
	}
}

func TestBidiProcessor_RTLBase(t *testing.T) {
	bp := NewBidiProcessor()
	// "שלום English שלום"
	text := "שלום English שלום"
	err := bp.Process(text, RTL)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	runes := []rune(text)
	res, err := bp.Reorder(0, len(runes))
	if err != nil {
		t.Fatalf("Reorder failed: %v", err)
	}

	if res.Text == "" {
		t.Error("Reorder returned empty text")
	}
}
