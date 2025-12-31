package drawing

import (
	"strings"
	"testing"

	"github.com/connerohnesorge/goffice-pdf/core"
)

func TestDiagramRenderer_RenderPlaceholder(
	t *testing.T,
) {
	t.Run("BasicPlaceholder", func(t *testing.T) {
		// Create a mock page that records operations
		mockPage := core.NewMockPage()
		ctx := core.NewRenderingContext(612, 792).
			WithPage(mockPage)
		renderer := NewDiagramRenderer(ctx)

		// Render placeholder
		err := renderer.RenderPlaceholder(
			100,
			200,
			300,
			150,
		)
		if err != nil {
			t.Fatalf(
				"RenderPlaceholder failed: %v",
				err,
			)
		}

		// Verify that graphics state was saved and restored
		if !mockPage.HasCall(
			"SaveGraphicsState()",
		) {
			t.Error(
				"Expected SaveGraphicsState call",
			)
		}
		if !mockPage.HasCall(
			"RestoreGraphicsState()",
		) {
			t.Error(
				"Expected RestoreGraphicsState call",
			)
		}

		// Verify that a filled rectangle was drawn (background)
		foundFillRect := false
		foundStrokeRect := false
		for _, call := range mockPage.Calls {
			if strings.Contains(
				call,
				"DrawRectangle",
			) {
				if strings.Contains(
					call,
					"true, false",
				) {
					foundFillRect = true
				}
				if strings.Contains(
					call,
					"false, true",
				) {
					foundStrokeRect = true
				}
			}
		}
		if !foundFillRect {
			t.Error(
				"Expected filled rectangle (background)",
			)
		}
		if !foundStrokeRect {
			t.Error(
				"Expected stroked rectangle (border)",
			)
		}

		// Verify that dash pattern was set
		foundDash := false
		for _, call := range mockPage.Calls {
			if strings.Contains(
				call,
				"SetLineDashPattern",
			) {
				foundDash = true

				break
			}
		}
		if !foundDash {
			t.Error(
				"Expected SetLineDashPattern call for dashed border",
			)
		}

		// Verify that text was drawn
		foundText := false
		for _, call := range mockPage.Calls {
			if strings.Contains(
				call,
				"SmartArt Diagram",
			) {
				foundText = true

				break
			}
		}
		if !foundText {
			t.Error(
				"Expected 'SmartArt Diagram' label in DrawText call",
			)
		}

		// Verify font was set to Helvetica
		if mockPage.FontName != "Helvetica" {
			t.Errorf(
				"Expected Helvetica font, got %s",
				mockPage.FontName,
			)
		}
		if mockPage.FontSize != 12.0 {
			t.Errorf(
				"Expected font size 12.0, got %g",
				mockPage.FontSize,
			)
		}
	})

	t.Run("DifferentSizes", func(t *testing.T) {
		mockPage := core.NewMockPage()
		ctx := core.NewRenderingContext(612, 792).
			WithPage(mockPage)
		renderer := NewDiagramRenderer(ctx)

		tests := []struct {
			name string
			x, y float64
			w, h float64
		}{
			{"Small", 50, 50, 100, 80},
			{"Medium", 100, 100, 300, 200},
			{"Large", 50, 50, 500, 400},
			{"Wide", 100, 100, 500, 100},
			{"Tall", 100, 100, 100, 500},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockPage.Reset()
				err := renderer.RenderPlaceholder(
					tt.x,
					tt.y,
					tt.w,
					tt.h,
				)
				if err != nil {
					t.Fatalf(
						"RenderPlaceholder failed for %s: %v",
						tt.name,
						err,
					)
				}

				if mockPage.CallCount() == 0 {
					t.Errorf(
						"No calls recorded for %s",
						tt.name,
					)
				}

				// Verify dimensions were used
				foundRect := false
				for _, call := range mockPage.Calls {
					if strings.Contains(
						call,
						"DrawRectangle",
					) {
						foundRect = true

						break
					}
				}
				if !foundRect {
					t.Errorf(
						"No rectangle drawn for %s",
						tt.name,
					)
				}
			})
		}
	})

	t.Run("NoPageSet", func(t *testing.T) {
		// Create context without setting Page
		ctx := core.NewRenderingContext(612, 792)
		renderer := NewDiagramRenderer(ctx)

		// Should not error even when Page is nil
		err := renderer.RenderPlaceholder(
			100,
			200,
			300,
			150,
		)
		if err != nil {
			t.Fatalf(
				"RenderPlaceholder should not error with nil Page: %v",
				err,
			)
		}
	})
}

func TestDiagramRenderer_RenderDiagramBounds(
	t *testing.T,
) {
	t.Run(
		"RenderBoundsHelper",
		func(t *testing.T) {
			mockPage := core.NewMockPage()
			ctx := core.NewRenderingContext(612, 792).
				WithPage(mockPage)
			renderer := NewDiagramRenderer(ctx)

			bounds := RenderBounds{
				X:      100,
				Y:      200,
				Width:  300,
				Height: 150,
			}

			err := renderer.RenderDiagramBounds(
				bounds,
			)
			if err != nil {
				t.Fatalf(
					"RenderDiagramBounds failed: %v",
					err,
				)
			}

			if mockPage.CallCount() == 0 {
				t.Error("No calls recorded")
			}

			// Should produce the same result as RenderPlaceholder
			foundText := false
			for _, call := range mockPage.Calls {
				if strings.Contains(
					call,
					"SmartArt Diagram",
				) {
					foundText = true

					break
				}
			}
			if !foundText {
				t.Error(
					"Expected diagram label in DrawText call",
				)
			}
		},
	)
}

func TestNewDiagramRenderer(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		ctx := core.NewRenderingContext(612, 792)
		renderer := NewDiagramRenderer(ctx)

		if renderer == nil {
			t.Fatal(
				"NewDiagramRenderer returned nil",
			)
		}
		if renderer.ctx != ctx {
			t.Error(
				"Renderer context not set correctly",
			)
		}
	})
}

// TestDiagramPlaceholderVisualProperties verifies the visual appearance properties
// of the placeholder rendering.
func TestDiagramPlaceholderVisualProperties(
	t *testing.T,
) {
	mockPage := core.NewMockPage()
	ctx := core.NewRenderingContext(612, 792).
		WithPage(mockPage)
	renderer := NewDiagramRenderer(ctx)

	err := renderer.RenderPlaceholder(
		100,
		200,
		300,
		150,
	)
	if err != nil {
		t.Fatalf(
			"RenderPlaceholder failed: %v",
			err,
		)
	}

	// Verify the mock page received the expected method calls
	if !mockPage.HasCall("SaveGraphicsState()") {
		t.Error("Expected SaveGraphicsState call")
	}
	if !mockPage.HasCall(
		"RestoreGraphicsState()",
	) {
		t.Error(
			"Expected RestoreGraphicsState call",
		)
	}

	// Check that fill and stroke colors were set
	foundFillColor := false
	foundStrokeColor := false
	foundDashPattern := false
	foundFont := false
	foundText := false

	for _, call := range mockPage.Calls {
		if strings.HasPrefix(
			call,
			"SetFillColor(",
		) {
			foundFillColor = true
		}
		if strings.HasPrefix(
			call,
			"SetStrokeColor(",
		) {
			foundStrokeColor = true
		}
		if strings.HasPrefix(
			call,
			"SetLineDashPattern(",
		) {
			foundDashPattern = true
		}
		if strings.HasPrefix(call, "SetFont(") {
			foundFont = true
		}
		if strings.HasPrefix(call, "DrawText(") {
			foundText = true
		}
	}

	if !foundFillColor {
		t.Error("Expected SetFillColor call")
	}
	if !foundStrokeColor {
		t.Error("Expected SetStrokeColor call")
	}
	if !foundDashPattern {
		t.Error(
			"Expected SetLineDashPattern call",
		)
	}
	if !foundFont {
		t.Error("Expected SetFont call")
	}
	if !foundText {
		t.Error("Expected DrawText call")
	}

	// Verify graphics state was properly balanced
	if mockPage.GraphicsStateDepth != 0 {
		t.Errorf(
			"Graphics state depth should be 0, got %d",
			mockPage.GraphicsStateDepth,
		)
	}
}
