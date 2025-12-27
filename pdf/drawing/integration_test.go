// Package drawing provides integration tests for PDF drawing operations.
package drawing

import (
	"image"
	"image/color"
	"math"
	"strings"
	"testing"
)

// TestIntegration_PathDrawingGeneratesValidOperators tests path drawing operations.
func TestIntegration_PathDrawingGeneratesValidOperators(
	t *testing.T,
) {
	t.Run(
		"complete path drawing pipeline",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.MoveTo(100, 100)
			path.LineTo(200, 100)
			path.LineTo(200, 200)
			path.LineTo(100, 200)
			path.ClosePath()

			content := path.Stroke()
			verifyPDFOperators(
				t,
				content,
				[]string{
					"m",
					"l",
					"l",
					"l",
					"h",
					"S",
				},
			)
		},
	)

	t.Run(
		"bezier curves generate correct operators",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.MoveTo(0, 0)
			path.CurveTo(25, 50, 75, 50, 100, 0)
			path.QuadraticCurveTo(50, 100, 0, 0)

			content := path.String()
			verifyPDFOperators(
				t,
				content,
				[]string{"m", "c"},
			)
			// QuadraticCurveTo should be converted to cubic
		},
	)

	t.Run(
		"rectangle shorthand",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.Rectangle(50, 50, 200, 100)

			content := path.String()
			if !strings.Contains(content, "re") {
				t.Error(
					"Rectangle should use 're' operator",
				)
			}
		},
	)

	t.Run(
		"circle approximation with bezier curves",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.Circle(200, 200, 50)

			content := path.String()
			// Circle uses 4 Bezier curve segments
			if strings.Count(
				content,
				" c\n",
			) < 4 {
				t.Error(
					"Circle should use at least 4 cubic Bezier curves",
				)
			}
		},
	)

	t.Run(
		"ellipse approximation",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.Ellipse(100, 100, 75, 50)

			content := path.String()
			if !strings.Contains(
				content,
				" c\n",
			) {
				t.Error(
					"Ellipse should use cubic Bezier curves",
				)
			}
		},
	)

	t.Run("arc operations", func(t *testing.T) {
		path := NewPathBuilder()
		path.MoveTo(100, 100)
		path.ArcTo(150, 150, 50, 50, 0, 90, false)

		content := path.String()
		if len(content) == 0 {
			t.Error(
				"Arc should generate PDF operators",
			)
		}
	})

	t.Run("polygon drawing", func(t *testing.T) {
		// Triangle
		path := NewPathBuilder()
		path.Polygon(100, 100, 200, 100, 150, 200)

		content := path.String()
		verifyPDFOperators(
			t,
			content,
			[]string{"m", "l", "l", "h"},
		)
	})

	t.Run(
		"rounded rectangle",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.RoundedRect(50, 50, 200, 100, 10)

			content := path.String()
			// Rounded rect uses lines and curves
			if !strings.Contains(
				content,
				" l\n",
			) ||
				!strings.Contains(
					content,
					" c\n",
				) {
				t.Error(
					"Rounded rectangle should use lines and curves",
				)
			}
		},
	)
}

// TestIntegration_PathPaintingOperations tests fill and stroke operations.
func TestIntegration_PathPaintingOperations(
	t *testing.T,
) {
	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)

	t.Run("stroke operation", func(t *testing.T) {
		content := path.Stroke()
		if !strings.HasSuffix(
			strings.TrimSpace(content),
			"S",
		) {
			t.Error(
				"Stroke should end with 'S' operator",
			)
		}
	})

	t.Run("fill operation", func(t *testing.T) {
		content := path.Fill()
		if !strings.HasSuffix(
			strings.TrimSpace(content),
			"f",
		) {
			t.Error(
				"Fill should end with 'f' operator",
			)
		}
	})

	t.Run(
		"fill even-odd operation",
		func(t *testing.T) {
			content := path.FillEvenOdd()
			if !strings.HasSuffix(
				strings.TrimSpace(content),
				"f*",
			) {
				t.Error(
					"Fill even-odd should end with 'f*' operator",
				)
			}
		},
	)

	t.Run(
		"fill and stroke operation",
		func(t *testing.T) {
			content := path.FillAndStroke()
			if !strings.HasSuffix(
				strings.TrimSpace(content),
				"B",
			) {
				t.Error(
					"Fill and stroke should end with 'B' operator",
				)
			}
		},
	)

	t.Run(
		"close path and stroke operation",
		func(t *testing.T) {
			closedPath := NewPathBuilder()
			closedPath.MoveTo(0, 0)
			closedPath.LineTo(100, 0)
			closedPath.LineTo(100, 100)
			closedPath.ClosePath()
			content := closedPath.Stroke()
			// Verify the path has close operator before stroke
			if !strings.Contains(content, "h") {
				t.Error(
					"Closed path should contain 'h' operator",
				)
			}
			if !strings.HasSuffix(
				strings.TrimSpace(content),
				"S",
			) {
				t.Error(
					"Stroke should end with 'S' operator",
				)
			}
		},
	)

	t.Run(
		"end path operation (for clipping)",
		func(t *testing.T) {
			content := path.EndPath()
			if !strings.HasSuffix(
				strings.TrimSpace(content),
				"n",
			) {
				t.Error(
					"End path should end with 'n' operator",
				)
			}
		},
	)
}

// TestIntegration_TextRendering tests text rendering operations.
func TestIntegration_TextRendering(t *testing.T) {
	t.Run(
		"basic text drawing",
		func(t *testing.T) {
			content := DrawTextAt(
				100,
				700,
				"Hello, World!",
				"/Helvetica",
				12,
			)

			// Should have BT/ET block
			if !strings.Contains(content, "BT") {
				t.Error(
					"Text should start with BT operator",
				)
			}
			if !strings.Contains(content, "ET") {
				t.Error(
					"Text should end with ET operator",
				)
			}
			// Should have Tf operator
			if !strings.Contains(content, "Tf") {
				t.Error(
					"Text should include Tf (font) operator",
				)
			}
			// Should have Tj operator
			if !strings.Contains(content, "Tj") {
				t.Error(
					"Text should include Tj (show text) operator",
				)
			}
		},
	)

	t.Run("text with style", func(t *testing.T) {
		style := NewTextStyle().
			WithFontSize(14).
			WithColor(Red).
			WithCharacterSpacing(0.5)

		content := DrawText(
			100,
			600,
			"Styled text",
			"/Times-Roman",
			style,
		)

		if !strings.Contains(content, "rg") {
			t.Error(
				"Text with color should include rg operator",
			)
		}
		if !strings.Contains(content, "Tc") {
			t.Error(
				"Text with character spacing should include Tc operator",
			)
		}
	})

	t.Run(
		"text builder fluent API",
		func(t *testing.T) {
			tb := NewTextBuilder()
			tb.BeginText()
			tb.SetFont("/Courier", 10)
			tb.SetTextPosition(50, 500)
			tb.SetLeading(14)
			tb.ShowText("First line")
			tb.NextLine()
			tb.ShowText("Second line")
			tb.EndText()

			content := tb.String()

			if strings.Count(content, "BT") != 1 {
				t.Error(
					"Should have exactly one BT operator",
				)
			}
			if strings.Count(content, "ET") != 1 {
				t.Error(
					"Should have exactly one ET operator",
				)
			}
			if !strings.Contains(content, "TL") {
				t.Error(
					"Should include TL (leading) operator",
				)
			}
			if !strings.Contains(content, "T*") {
				t.Error(
					"Should include T* (next line) operator",
				)
			}
		},
	)

	t.Run(
		"text with rendering modes",
		func(t *testing.T) {
			style := NewTextStyle().WithRenderingMode(TextRenderStroke)
			content := DrawText(
				100,
				400,
				"Outline text",
				"/Helvetica",
				style,
			)

			if !strings.Contains(content, "Tr") {
				t.Error(
					"Text with non-default rendering mode should include Tr operator",
				)
			}
		},
	)

	t.Run(
		"escaped special characters",
		func(t *testing.T) {
			content := DrawTextAt(
				100,
				300,
				"Test (with) \\special\\ chars",
				"/Helvetica",
				12,
			)

			// Check that parentheses and backslashes are escaped
			if !strings.Contains(content, "\\(") {
				t.Error(
					"Left parenthesis should be escaped",
				)
			}
			if !strings.Contains(content, "\\)") {
				t.Error(
					"Right parenthesis should be escaped",
				)
			}
			if !strings.Contains(
				content,
				"\\\\",
			) {
				t.Error(
					"Backslash should be escaped",
				)
			}
		},
	)

	t.Run(
		"hex text for non-ASCII",
		func(t *testing.T) {
			tb := NewTextBuilder()
			tb.BeginText()
			tb.SetFont("/F1", 12)
			tb.ShowHexText(
				[]byte{
					0x00,
					0x48,
					0x00,
					0x65,
					0x00,
					0x6C,
					0x00,
					0x6C,
					0x00,
					0x6F,
				},
			)
			tb.EndText()

			content := tb.String()

			if !strings.Contains(content, "<") ||
				!strings.Contains(content, ">") {
				t.Error(
					"Hex text should use angle brackets",
				)
			}
		},
	)
}

// TestIntegration_GraphicsStateOperations tests graphics state management.
func TestIntegration_GraphicsStateOperations(
	t *testing.T,
) {
	t.Run(
		"graphics state save/restore",
		func(t *testing.T) {
			gs := NewGraphicsState()
			gs.Save()
			gs.Translate(100, 100)
			gs.Rotate(45)
			gs.Scale(2, 2)
			gs.Restore()

			content := gs.String()

			if strings.Count(content, "q") != 1 {
				t.Error(
					"Should have one 'q' (save) operator",
				)
			}
			if strings.Count(content, "Q") != 1 {
				t.Error(
					"Should have one 'Q' (restore) operator",
				)
			}
			if strings.Count(content, "cm") < 1 {
				t.Error(
					"Should have 'cm' (transform) operators",
				)
			}
		},
	)

	t.Run(
		"nested graphics states",
		func(t *testing.T) {
			gs := NewGraphicsState()
			gs.Save()
			gs.Translate(50, 50)
			gs.Save()
			gs.Scale(0.5, 0.5)
			gs.Restore()
			gs.Restore()

			content := gs.String()

			if strings.Count(content, "q") != 2 {
				t.Error(
					"Should have two 'q' operators",
				)
			}
			if strings.Count(content, "Q") != 2 {
				t.Error(
					"Should have two 'Q' operators",
				)
			}
		},
	)

	t.Run(
		"transform matrices",
		func(t *testing.T) {
			// Test translation
			t1 := Identity().Translate(100, 200)
			if t1.E != 100 || t1.F != 200 {
				t.Errorf(
					"Translate(100, 200) = e=%v, f=%v, want e=100, f=200",
					t1.E,
					t1.F,
				)
			}

			// Test scaling
			t2 := Identity().Scale(2, 3)
			if t2.A != 2 || t2.D != 3 {
				t.Errorf(
					"Scale(2, 3) = a=%v, d=%v, want a=2, d=3",
					t2.A,
					t2.D,
				)
			}

			// Test rotation
			t3 := Identity().Rotate(90)
			if math.Abs(t3.A) > 0.001 ||
				math.Abs(t3.B-1) > 0.001 {
				t.Errorf(
					"Rotate(90) = a=%v, b=%v, want a~0, b~1",
					t3.A,
					t3.B,
				)
			}

			// Test concatenation
			t4 := Identity().Translate(100, 0).
				Scale(2, 1)
			x, y := t4.TransformPoint(0, 0)
			if x != 100 || y != 0 {
				t.Errorf(
					"TransformPoint(0,0) after translate+scale = (%v, %v), want (100, 0)",
					x,
					y,
				)
			}
		},
	)

	t.Run(
		"inverse transform",
		func(t *testing.T) {
			t1 := Identity().Translate(100, 200).
				Scale(2, 2)
			inv := t1.Inverse()

			// Apply transform then inverse should give identity
			result := t1.Concat(inv)
			if !result.IsIdentity() {
				t.Error(
					"T * T^-1 should be identity",
				)
			}
		},
	)
}

// TestIntegration_ColorOperations tests color operations.
func TestIntegration_ColorOperations(
	t *testing.T,
) {
	t.Run("RGB fill color", func(t *testing.T) {
		c := NewRGB(1, 0, 0) // Red
		content := c.SetFillRGB()

		if !strings.Contains(content, "rg") {
			t.Error(
				"Fill RGB should use 'rg' operator",
			)
		}
		if !strings.Contains(content, "1") &&
			!strings.Contains(content, "0") {
			t.Error("Should contain color values")
		}
	})

	t.Run("RGB stroke color", func(t *testing.T) {
		c := NewRGB(0, 0, 1) // Blue
		content := c.SetStrokeRGB()

		if !strings.Contains(content, "RG") {
			t.Error(
				"Stroke RGB should use 'RG' operator",
			)
		}
	})

	t.Run("grayscale colors", func(t *testing.T) {
		fillGray := SetFillGray(0.5)
		if !strings.Contains(fillGray, "g") {
			t.Error(
				"Fill gray should use 'g' operator",
			)
		}

		strokeGray := SetStrokeGray(0.5)
		if !strings.Contains(strokeGray, "G") {
			t.Error(
				"Stroke gray should use 'G' operator",
			)
		}
	})

	t.Run("CMYK colors", func(t *testing.T) {
		c := NewCMYK(0, 1, 1, 0) // Red in CMYK
		fill := c.SetFillCMYK()
		stroke := c.SetStrokeCMYK()

		if !strings.Contains(fill, "k") {
			t.Error(
				"Fill CMYK should use 'k' operator",
			)
		}
		if !strings.Contains(stroke, "K") {
			t.Error(
				"Stroke CMYK should use 'K' operator",
			)
		}
	})

	t.Run("color parsing", func(t *testing.T) {
		tests := []struct {
			input string
			want  Color
		}{
			{"red", Red},
			{"#FF0000", NewRGB8(255, 0, 0)},
			{"#F00", NewRGB8(255, 0, 0)},
			{
				"rgb(255, 0, 0)",
				NewRGB8(255, 0, 0),
			},
		}

		for _, tt := range tests {
			got := ParseColor(tt.input)
			// Compare with tolerance due to floating point
			if math.Abs(got.R-tt.want.R) > 0.01 ||
				math.Abs(
					got.G-tt.want.G,
				) > 0.01 ||
				math.Abs(got.B-tt.want.B) > 0.01 {
				t.Errorf(
					"ParseColor(%q) = %v, want %v",
					tt.input,
					got,
					tt.want,
				)
			}
		}
	})

	t.Run(
		"color conversions",
		func(t *testing.T) {
			// RGB to CMYK and back
			rgb := NewRGB(1, 0, 0) // Pure red
			cmyk := RGBToCMYK(rgb)
			rgbBack := cmyk.ToRGB()

			if math.Abs(rgb.R-rgbBack.R) > 0.01 ||
				math.Abs(
					rgb.G-rgbBack.G,
				) > 0.01 ||
				math.Abs(rgb.B-rgbBack.B) > 0.01 {
				t.Errorf(
					"RGB->CMYK->RGB conversion failed: %v -> %v -> %v",
					rgb,
					cmyk,
					rgbBack,
				)
			}

			// RGB to HSL and back
			hsl := RGBToHSL(rgb)
			rgbFromHSL := hsl.ToRGB()

			if math.Abs(
				rgb.R-rgbFromHSL.R,
			) > 0.01 ||
				math.Abs(
					rgb.G-rgbFromHSL.G,
				) > 0.01 ||
				math.Abs(
					rgb.B-rgbFromHSL.B,
				) > 0.01 {
				t.Errorf(
					"RGB->HSL->RGB conversion failed: %v -> %v -> %v",
					rgb,
					hsl,
					rgbFromHSL,
				)
			}
		},
	)
}

// TestIntegration_StrokeStyles tests stroke styling operations.
func TestIntegration_StrokeStyles(t *testing.T) {
	t.Run(
		"stroke style content stream",
		func(t *testing.T) {
			style := NewStrokeStyle().
				SetWidth(2.5).
				SetCap(LineCapRound).
				SetJoin(LineJoinRound).
				SetColor(Blue)

			content := style.ContentStream()

			if !strings.Contains(content, "w") {
				t.Error(
					"Stroke style should include line width (w)",
				)
			}
			if !strings.Contains(content, "J") {
				t.Error(
					"Stroke style should include line cap (J)",
				)
			}
			if !strings.Contains(content, "j") {
				t.Error(
					"Stroke style should include line join (j)",
				)
			}
			if !strings.Contains(content, "RG") {
				t.Error(
					"Stroke style should include stroke color (RG)",
				)
			}
		},
	)

	t.Run("dash patterns", func(t *testing.T) {
		tests := []struct {
			pattern DashPattern
			want    string
		}{
			{DashSolid, "[] 0 d"},
			{DashDot, "[1 2] 0 d"},
			{DashDash, "[4 3] 0 d"},
			{DashDashDot, "[4 3 1 3] 0 d"},
		}

		for _, tt := range tests {
			got := tt.pattern.ContentStream()
			if got != tt.want {
				t.Errorf(
					"DashPattern.ContentStream() = %q, want %q",
					got,
					tt.want,
				)
			}
		}
	})

	t.Run(
		"stroke path helper",
		func(t *testing.T) {
			path := NewPathBuilder()
			path.Rectangle(0, 0, 100, 100)

			style := NewStrokeStyle().
				SetWidth(1.5).
				SetColor(Red)

			content := StrokePath(path, style)

			if !strings.Contains(content, "re") {
				t.Error(
					"Should contain rectangle operator",
				)
			}
			if !strings.Contains(content, "S") {
				t.Error(
					"Should contain stroke operator",
				)
			}
		},
	)
}

// TestIntegration_ClippingOperations tests clipping path operations.
func TestIntegration_ClippingOperations(
	t *testing.T,
) {
	t.Run("rectangular clip", func(t *testing.T) {
		content := RectClipOperators(
			50,
			50,
			200,
			100,
			ClipNonZeroRule,
		)

		if !strings.Contains(content, "re") {
			t.Error(
				"Rect clip should contain rectangle operator",
			)
		}
		if !strings.Contains(content, "W") {
			t.Error(
				"Clip should contain W operator",
			)
		}
		if !strings.Contains(content, "n") {
			t.Error(
				"Clip should end with n operator",
			)
		}
	})

	t.Run("circular clip", func(t *testing.T) {
		content := CircleClipOperators(
			100,
			100,
			50,
			ClipNonZeroRule,
		)

		if !strings.Contains(content, "c") {
			t.Error(
				"Circle clip should contain curve operators",
			)
		}
		if !strings.Contains(content, "W") {
			t.Error(
				"Clip should contain W operator",
			)
		}
	})

	t.Run(
		"even-odd clip rule",
		func(t *testing.T) {
			path := ClipRect(0, 0, 100, 100)
			content := PathClipOperators(
				path,
				ClipEvenOddRule,
			)

			if !strings.Contains(content, "W*") {
				t.Error(
					"Even-odd clip should use W* operator",
				)
			}
		},
	)

	t.Run(
		"scoped clip operations",
		func(t *testing.T) {
			path := ClipRect(0, 0, 200, 200)
			innerContent := "0 0 0 rg\n0 0 100 100 re f\n"

			content := ScopedClipOperators(
				path,
				ClipNonZeroRule,
				innerContent,
			)

			if !strings.Contains(content, "q\n") {
				t.Error(
					"Scoped clip should save graphics state",
				)
			}
			if !strings.Contains(
				content,
				"\nQ\n",
			) {
				t.Error(
					"Scoped clip should restore graphics state",
				)
			}
		},
	)

	t.Run(
		"clip context nested operations",
		func(t *testing.T) {
			clip := NewClipContext()
			clip.PushClip(
				ClipRect(10, 10, 100, 100),
				ClipNonZeroRule,
			)
			clip.PushClip(
				ClipCircle(60, 60, 30),
				ClipNonZeroRule,
			)
			clip.WriteContent("% inner content\n")
			clip.PopAllClips()

			content := clip.String()

			if strings.Count(
				content,
				"q\n",
			) != 2 {
				t.Errorf(
					"Should have 2 save operations, got %d",
					strings.Count(content, "q\n"),
				)
			}
			if strings.Count(
				content,
				"Q\n",
			) != 2 {
				t.Errorf(
					"Should have 2 restore operations, got %d",
					strings.Count(content, "Q\n"),
				)
			}
		},
	)
}

// TestIntegration_ImageOperations tests image embedding operations.
func TestIntegration_ImageOperations(
	t *testing.T,
) {
	t.Run(
		"draw image operator",
		func(t *testing.T) {
			content := DrawImage(
				100,
				200,
				300,
				400,
				"Im1",
			)

			// Should save/restore graphics state
			if !strings.Contains(content, "q") {
				t.Error(
					"Draw image should save graphics state",
				)
			}
			if !strings.Contains(content, "Q") {
				t.Error(
					"Draw image should restore graphics state",
				)
			}
			// Should have cm transformation
			if !strings.Contains(content, "cm") {
				t.Error(
					"Draw image should include transformation matrix",
				)
			}
			// Should invoke XObject
			if !strings.Contains(content, "Do") {
				t.Error(
					"Draw image should use Do operator",
				)
			}
			if !strings.Contains(
				content,
				"/Im1",
			) {
				t.Error(
					"Draw image should reference image name",
				)
			}
		},
	)

	t.Run(
		"image format detection",
		func(t *testing.T) {
			tests := []struct {
				data   []byte
				format ImageFormat
				name   string
			}{
				// JPEG: starts with FF D8 FF, need at least 8 bytes
				{
					[]byte{
						0xFF,
						0xD8,
						0xFF,
						0xE0,
						0x00,
						0x10,
						0x4A,
						0x46,
					},
					ImageFormatJPEG,
					"JPEG",
				},
				// PNG: starts with 89 50 4E 47 0D 0A 1A 0A (exactly 8 byte signature)
				{
					[]byte{
						0x89,
						0x50,
						0x4E,
						0x47,
						0x0D,
						0x0A,
						0x1A,
						0x0A,
					},
					ImageFormatPNG,
					"PNG",
				},
				// GIF: starts with GIF89a or GIF87a (6 bytes), need 8 for DetectFormat
				{
					[]byte{
						0x47,
						0x49,
						0x46,
						0x38,
						0x39,
						0x61,
						0x00,
						0x00,
					},
					ImageFormatGIF,
					"GIF",
				},
				// BMP: starts with BM, need 8 bytes total
				{
					[]byte{
						0x42,
						0x4D,
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
					},
					ImageFormatBMP,
					"BMP",
				},
				// Unknown format with 8 bytes
				{
					[]byte{
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
						0x00,
					},
					ImageFormatUnknown,
					"Unknown",
				},
			}

			for _, tt := range tests {
				got := DetectFormat(tt.data)
				if got != tt.format {
					t.Errorf(
						"DetectFormat(%s) = %v, want %v",
						tt.name,
						got,
						tt.format,
					)
				}
			}
		},
	)

	t.Run(
		"create image from Go image",
		func(t *testing.T) {
			// Create a simple test image
			img := image.NewRGBA(
				image.Rect(0, 0, 10, 10),
			)
			for y := range 10 {
				for x := range 10 {
					img.Set(
						x,
						y,
						color.RGBA{
							R: 255,
							G: 0,
							B: 0,
							A: 255,
						},
					)
				}
			}

			xobj, err := LoadFromGoImage(
				img,
				"TestImg",
			)
			if err != nil {
				t.Fatalf(
					"LoadFromGoImage() error = %v",
					err,
				)
			}

			if xobj.Width != 10 {
				t.Errorf(
					"Width = %d, want 10",
					xobj.Width,
				)
			}
			if xobj.Height != 10 {
				t.Errorf(
					"Height = %d, want 10",
					xobj.Height,
				)
			}
			if xobj.ColorSpace != ColorSpaceDeviceRGB {
				t.Errorf(
					"ColorSpace = %v, want DeviceRGB",
					xobj.ColorSpace,
				)
			}
			if len(xobj.Data) == 0 {
				t.Error(
					"Image data should not be empty",
				)
			}
			if xobj.Filter != "FlateDecode" {
				t.Errorf(
					"Filter = %q, want FlateDecode",
					xobj.Filter,
				)
			}
		},
	)

	t.Run(
		"image with alpha channel",
		func(t *testing.T) {
			// Create image with transparency
			img := image.NewRGBA(
				image.Rect(0, 0, 5, 5),
			)
			for y := range 5 {
				for x := range 5 {
					img.Set(
						x,
						y,
						color.RGBA{
							R: 128,
							G: 128,
							B: 128,
							A: 128,
						},
					)
				}
			}

			xobj, err := LoadFromGoImage(
				img,
				"AlphaImg",
			)
			if err != nil {
				t.Fatalf(
					"LoadFromGoImage() error = %v",
					err,
				)
			}

			if !xobj.HasAlpha() {
				t.Error(
					"Image should have alpha channel",
				)
			}
			if xobj.SMask == nil {
				t.Error(
					"SMask should not be nil for image with alpha",
				)
			}
			if xobj.SMask.ColorSpace != ColorSpaceDeviceGray {
				t.Error(
					"SMask should be grayscale",
				)
			}
		},
	)

	t.Run("image fit modes", func(t *testing.T) {
		xobj := &ImageXObject{
			Name:   "FitTest",
			Width:  200,
			Height: 100,
		}

		rect := NewRect(0, 0, 100, 100)

		// Contain mode
		bounds := ImageBounds(
			rect,
			xobj,
			FitModeContain,
		)
		if bounds.Width > rect.Width ||
			bounds.Height > rect.Height {
			t.Error(
				"Contain mode should not exceed bounds",
			)
		}

		// Cover mode
		bounds = ImageBounds(
			rect,
			xobj,
			FitModeCover,
		)
		if bounds.Width < rect.Width &&
			bounds.Height < rect.Height {
			t.Error(
				"Cover mode should cover bounds",
			)
		}

		// Stretch mode
		bounds = ImageBounds(
			rect,
			xobj,
			FitModeStretch,
		)
		if bounds.Width != rect.Width ||
			bounds.Height != rect.Height {
			t.Error(
				"Stretch mode should match bounds exactly",
			)
		}
	})

	t.Run("image manager", func(t *testing.T) {
		mgr := NewImageManager()

		img1 := NewImageXObject("")
		img1.Width = 100
		img1.Height = 100
		mgr.Add(img1)

		img2 := NewImageXObject("")
		img2.Width = 200
		img2.Height = 200
		mgr.Add(img2)

		if mgr.Count() != 2 {
			t.Errorf(
				"Count() = %d, want 2",
				mgr.Count(),
			)
		}

		// Verify names are assigned
		if img1.Name != "Im1" {
			t.Errorf(
				"First image name = %q, want Im1",
				img1.Name,
			)
		}
		if img2.Name != "Im2" {
			t.Errorf(
				"Second image name = %q, want Im2",
				img2.Name,
			)
		}

		// Verify retrieval
		retrieved, ok := mgr.Get("Im1")
		if !ok {
			t.Error("Should find Im1")
		}
		if retrieved.Width != 100 {
			t.Errorf(
				"Retrieved image width = %d, want 100",
				retrieved.Width,
			)
		}
	})

	t.Run("inline image", func(t *testing.T) {
		ii := NewInlineImage(10, 10)
		ii.Data = make([]byte, 300) // 10x10x3 RGB

		content := ii.ContentStream()

		if !strings.Contains(content, "BI") {
			t.Error(
				"Inline image should start with BI",
			)
		}
		if !strings.Contains(content, "ID") {
			t.Error(
				"Inline image should have ID separator",
			)
		}
		if !strings.Contains(content, "EI") {
			t.Error(
				"Inline image should end with EI",
			)
		}
	})
}

// TestIntegration_CoordinateTransformations tests coordinate transformations.
func TestIntegration_CoordinateTransformations(
	t *testing.T,
) {
	t.Run("translation", func(t *testing.T) {
		transform := Identity().Translate(100, 200)
		x, y := transform.TransformPoint(0, 0)

		if x != 100 || y != 200 {
			t.Errorf(
				"TransformPoint(0,0) after Translate(100,200) = (%v,%v), want (100,200)",
				x,
				y,
			)
		}
	})

	t.Run("scaling", func(t *testing.T) {
		transform := Identity().Scale(2, 3)
		x, y := transform.TransformPoint(10, 10)

		if x != 20 || y != 30 {
			t.Errorf(
				"TransformPoint(10,10) after Scale(2,3) = (%v,%v), want (20,30)",
				x,
				y,
			)
		}
	})

	t.Run(
		"rotation 90 degrees",
		func(t *testing.T) {
			transform := Identity().Rotate(90)
			x, y := transform.TransformPoint(
				10,
				0,
			)

			// Point (10,0) rotated 90 degrees should be approximately (0,10)
			if math.Abs(x) > 0.001 ||
				math.Abs(y-10) > 0.001 {
				t.Errorf(
					"TransformPoint(10,0) after Rotate(90) = (%v,%v), want ~(0,10)",
					x,
					y,
				)
			}
		},
	)

	t.Run(
		"combined transformations",
		func(t *testing.T) {
			// Translate, then scale, then rotate
			transform := Identity().
				Translate(50, 50).
				Scale(2, 2).
				Rotate(45)

			// The transforms are applied in order
			x, y := transform.TransformPoint(0, 0)
			// After translate(50,50): (50,50)
			// After scale(2,2) on (50,50): (100,100)
			// After rotate(45): point rotates around origin

			// Verify something reasonable happened
			if math.IsNaN(x) || math.IsNaN(y) {
				t.Error(
					"Transform resulted in NaN",
				)
			}
		},
	)

	t.Run(
		"skew transformation",
		func(t *testing.T) {
			transform := Identity().Skew(45, 0)
			x, y := transform.TransformPoint(
				10,
				10,
			)

			// With 45 degree x-skew, x should increase by y (tan(45) = 1)
			if math.Abs(x-20) > 0.001 || y != 10 {
				t.Errorf(
					"TransformPoint(10,10) after Skew(45,0) = (%v,%v), want (20,10)",
					x,
					y,
				)
			}
		},
	)

	t.Run(
		"CM operator output",
		func(t *testing.T) {
			transform := NewTransformFromMatrix(
				1,
				0,
				0,
				1,
				100,
				200,
			)
			op := transform.ToCMOperator()

			if !strings.Contains(op, "cm") {
				t.Error(
					"ToCMOperator should output 'cm' operator",
				)
			}
			if !strings.Contains(op, "100") ||
				!strings.Contains(op, "200") {
				t.Error(
					"ToCMOperator should contain translation values",
				)
			}
		},
	)
}

// TestIntegration_FillOperations tests fill operations.
func TestIntegration_FillOperations(
	t *testing.T,
) {
	t.Run("solid fill", func(t *testing.T) {
		fill := NewSolidFill(Red)
		content := fill.ContentStream(
			NewRect(0, 0, 100, 100),
		)

		if !strings.Contains(content, "rg") {
			t.Error(
				"Solid fill should output rg operator",
			)
		}
	})

	t.Run("no fill", func(t *testing.T) {
		fill := NoFill{}
		content := fill.ContentStream(
			NewRect(0, 0, 100, 100),
		)

		if content != "" {
			t.Errorf(
				"No fill should output empty string, got %q",
				content,
			)
		}
		if !fill.IsNone() {
			t.Error(
				"NoFill.IsNone() should return true",
			)
		}
	})

	t.Run(
		"transparent fill is none",
		func(t *testing.T) {
			fill := NewSolidFill(Transparent)

			if !fill.IsNone() {
				t.Error(
					"Transparent solid fill should be considered 'none'",
				)
			}
		},
	)
}

// TestIntegration_CompleteDrawingPipeline tests a complete drawing workflow.
func TestIntegration_CompleteDrawingPipeline(
	t *testing.T,
) {
	t.Run(
		"create complex drawing",
		func(t *testing.T) {
			var content strings.Builder

			// Save graphics state
			content.WriteString("q\n")

			// Set stroke style
			stroke := NewStrokeStyle().
				SetWidth(2).
				SetColor(DarkBlue).
				SetCap(LineCapRound).
				SetJoin(LineJoinRound)
			content.WriteString(
				stroke.ContentStream(),
			)
			content.WriteString("\n")

			// Draw a rectangle
			path := NewPathBuilder()
			path.Rectangle(50, 50, 200, 100)
			content.WriteString(path.Stroke())
			content.WriteString("\n")

			// Draw a circle with fill
			content.WriteString(Red.SetFillRGB())
			content.WriteString("\n")
			circlePath := NewPathBuilder()
			circlePath.Circle(300, 100, 40)
			content.WriteString(
				circlePath.FillAndStroke(),
			)
			content.WriteString("\n")

			// Add text
			content.WriteString(
				DrawTextAt(
					50,
					200,
					"Hello, PDF!",
					"/Helvetica",
					24,
				),
			)
			content.WriteString("\n")

			// Restore graphics state
			content.WriteString("Q\n")

			result := content.String()

			// Verify all expected operators are present
			requiredOps := []string{
				"q",
				"Q",
				"w",
				"J",
				"j",
				"RG",
				"re",
				"S",
				"rg",
				"B",
				"BT",
				"ET",
				"Tf",
				"Tj",
			}
			for _, op := range requiredOps {
				if !strings.Contains(result, op) {
					t.Errorf(
						"Complete drawing should contain %q operator",
						op,
					)
				}
			}
		},
	)
}

// verifyPDFOperators checks that the content contains the expected PDF operators.
func verifyPDFOperators(
	t *testing.T,
	content string,
	expectedOps []string,
) {
	t.Helper()
	for _, op := range expectedOps {
		// Check for operator (followed by space or newline)
		hasOp := strings.Contains(
			content,
			op+"\n",
		) ||
			strings.Contains(content, op+" ") ||
			strings.HasSuffix(
				strings.TrimSpace(content),
				op,
			)
		if !hasOp {
			t.Errorf(
				"Content missing expected operator %q\nContent: %s",
				op,
				content,
			)
		}
	}
}

// BenchmarkPathBuilding benchmarks path building operations.
func BenchmarkPathBuilding(b *testing.B) {
	b.Run("simple rectangle", func(b *testing.B) {
		for range b.N {
			path := NewPathBuilder()
			path.Rectangle(0, 0, 100, 100)
			_ = path.String()
		}
	})

	b.Run("complex path", func(b *testing.B) {
		for range b.N {
			path := NewPathBuilder()
			path.MoveTo(0, 0)
			for j := range 100 {
				path.LineTo(
					float64(j*10),
					float64(j*5),
				)
				path.CurveTo(
					float64(j*10+5),
					float64(j*5+10),
					float64(
						j*10+10,
					),
					float64(j*5+10),
					float64(
						j*10+15,
					),
					float64(j*5),
				)
			}
			path.ClosePath()
			_ = path.Stroke()
		}
	})
}

// BenchmarkTransformOperations benchmarks transform operations.
func BenchmarkTransformOperations(b *testing.B) {
	b.Run("single transform", func(b *testing.B) {
		for range b.N {
			t := Identity().Translate(100, 100).
				Rotate(45).
				Scale(2, 2)
			t.TransformPoint(50, 50)
		}
	})

	b.Run(
		"transform concatenation",
		func(b *testing.B) {
			t1 := Identity().Translate(100, 100)
			t2 := Identity().Rotate(45)
			t3 := Identity().Scale(2, 2)

			b.ResetTimer()
			for range b.N {
				_ = t1.Concat(t2).Concat(t3)
			}
		},
	)
}

// BenchmarkTextRendering benchmarks text rendering operations.
func BenchmarkTextRendering(b *testing.B) {
	b.Run("simple text", func(b *testing.B) {
		for range b.N {
			_ = DrawTextAt(
				100,
				100,
				"Hello, World!",
				"/Helvetica",
				12,
			)
		}
	})

	b.Run("styled text", func(b *testing.B) {
		style := NewTextStyle().
			WithFontSize(14).
			WithColor(Red).
			WithCharacterSpacing(0.5)

		b.ResetTimer()
		for range b.N {
			_ = DrawText(
				100,
				100,
				"Styled text content",
				"/Times-Roman",
				style,
			)
		}
	})
}

// BenchmarkColorOperations benchmarks color operations.
func BenchmarkColorOperations(b *testing.B) {
	b.Run("hex parsing", func(b *testing.B) {
		for range b.N {
			_ = ParseHex("#FF5733")
		}
	})

	b.Run("color conversion", func(b *testing.B) {
		c := NewRGB(0.8, 0.5, 0.2)
		b.ResetTimer()
		for range b.N {
			cmyk := RGBToCMYK(c)
			_ = cmyk.ToRGB()
		}
	})
}
