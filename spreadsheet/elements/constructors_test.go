package elements

import (
	"testing"
)

// TestConstructorsNilCheck tests that all constructor functions return non-nil values
func TestConstructorsNilCheck(t *testing.T) {
	tests := []struct {
		name string
		fn   func() any
	}{
		{"Workbook", func() any { return NewWorkbook() }},
		{"Worksheet", func() any { return NewWorksheet() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn()
			if result == nil {
				t.Fatalf("New%s() returned nil", tt.name)
			}
		})
	}
}
