package core

import (
	"math"
	"testing"
)

func TestOrientation_String(t *testing.T) {
	tests := []struct {
		o    Orientation
		want string
	}{
		{OrientationPortrait, "portrait"},
		{OrientationLandscape, "landscape"},
	}

	for _, tt := range tests {
		if got := tt.o.String(); got != tt.want {
			t.Errorf(
				"Orientation(%d).String() = %q, want %q",
				tt.o,
				got,
				tt.want,
			)
		}
	}
}

func TestAdditionalPageSizes(t *testing.T) {
	tests := []struct {
		name       string
		size       PageSize
		wantWidth  float64 // Expected width in mm (approximate)
		wantHeight float64 // Expected height in mm (approximate)
	}{
		{"A0", PageSizeA0, 841, 1189},
		{"A1", PageSizeA1, 594, 841},
		{"A2", PageSizeA2, 420, 594},
		{"A5", PageSizeA5, 148, 210},
		{"A6", PageSizeA6, 105, 148},
		{"B4", PageSizeB4, 250, 353},
		{"B5", PageSizeB5, 176, 250},
		{
			"Executive",
			PageSizeExecutive,
			184.15,
			266.7,
		}, // 7.25in x 10.5in
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widthMM := tt.size.WidthMM()
			heightMM := tt.size.HeightMM()

			if math.Abs(
				widthMM-tt.wantWidth,
			) > 1 {
				t.Errorf(
					"%s width = %vmm, want ~%vmm",
					tt.name,
					widthMM,
					tt.wantWidth,
				)
			}
			if math.Abs(
				heightMM-tt.wantHeight,
			) > 1 {
				t.Errorf(
					"%s height = %vmm, want ~%vmm",
					tt.name,
					heightMM,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestNewPageSize(t *testing.T) {
	size := NewPageSize(500, 700)
	if size.Width != 500 || size.Height != 700 {
		t.Errorf(
			"NewPageSize(500, 700) = %+v",
			size,
		)
	}
}

func TestNewPageSizeFromDimensions(t *testing.T) {
	width := Inches(8.5)
	height := Inches(11)
	size := NewPageSizeFromDimensions(
		width,
		height,
	)

	if size.Width != 612 {
		t.Errorf(
			"Width = %v, want 612",
			size.Width,
		)
	}
	if size.Height != 792 {
		t.Errorf(
			"Height = %v, want 792",
			size.Height,
		)
	}
}

func TestNewPageSizeFromUnit(t *testing.T) {
	size := NewPageSizeFromUnit(210, 297, UnitMM)

	// A4 is 210mm x 297mm = 595.28 x 841.89 points
	if math.Abs(size.Width-595.28) > 0.1 {
		t.Errorf(
			"Width = %v, want ~595.28",
			size.Width,
		)
	}
	if math.Abs(size.Height-841.89) > 0.1 {
		t.Errorf(
			"Height = %v, want ~841.89",
			size.Height,
		)
	}
}

func TestPageSize_WithOrientation(t *testing.T) {
	tests := []struct {
		name        string
		size        PageSize
		orientation Orientation
		wantWidth   float64
		wantHeight  float64
	}{
		{
			"portrait to portrait",
			PageSizeA4,
			OrientationPortrait,
			PageSizeA4.Width,
			PageSizeA4.Height,
		},
		{
			"portrait to landscape",
			PageSizeA4,
			OrientationLandscape,
			PageSizeA4.Height,
			PageSizeA4.Width,
		},
		{
			"landscape to landscape",
			PageSizeA4.Landscape(),
			OrientationLandscape,
			PageSizeA4.Height,
			PageSizeA4.Width,
		},
		{
			"landscape to portrait",
			PageSizeA4.Landscape(),
			OrientationPortrait,
			PageSizeA4.Width,
			PageSizeA4.Height,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.size.WithOrientation(
				tt.orientation,
			)
			if result.Width != tt.wantWidth ||
				result.Height != tt.wantHeight {
				t.Errorf(
					"WithOrientation(%v) = %+v, want {Width: %v, Height: %v}",
					tt.orientation,
					result,
					tt.wantWidth,
					tt.wantHeight,
				)
			}
		})
	}
}

func TestPageSize_Portrait(t *testing.T) {
	// Start with landscape
	landscape := PageSizeA4.Landscape()
	portrait := landscape.Portrait()

	if !portrait.IsPortrait() {
		t.Error(
			"Portrait() should return portrait orientation",
		)
	}
	if portrait.Width != PageSizeA4.Width ||
		portrait.Height != PageSizeA4.Height {
		t.Errorf(
			"Portrait() = %+v, want %+v",
			portrait,
			PageSizeA4,
		)
	}

	// Already portrait should stay the same
	alreadyPortrait := PageSizeA4.Portrait()
	if alreadyPortrait.Width != PageSizeA4.Width ||
		alreadyPortrait.Height != PageSizeA4.Height {
		t.Errorf(
			"Portrait() on portrait = %+v, want %+v",
			alreadyPortrait,
			PageSizeA4,
		)
	}
}

func TestPageSize_IsPortrait(t *testing.T) {
	if !PageSizeA4.IsPortrait() {
		t.Error("A4 should be portrait")
	}
	if PageSizeA4.Landscape().IsPortrait() {
		t.Error(
			"A4 Landscape should not be portrait",
		)
	}

	// Square should be considered portrait (height >= width)
	square := NewPageSize(100, 100)
	if !square.IsPortrait() {
		t.Error(
			"Square should be considered portrait",
		)
	}
}

func TestPageSize_Orientation(t *testing.T) {
	if PageSizeA4.Orientation() != OrientationPortrait {
		t.Error(
			"A4 orientation should be portrait",
		)
	}
	if PageSizeA4.Landscape().
		Orientation() !=
		OrientationLandscape {
		t.Error(
			"A4 Landscape orientation should be landscape",
		)
	}
}

func TestPageSize_UnitConversions(t *testing.T) {
	// Test US Letter
	if PageSizeLetter.WidthInches() != 8.5 {
		t.Errorf(
			"Letter width in inches = %v, want 8.5",
			PageSizeLetter.WidthInches(),
		)
	}
	if PageSizeLetter.HeightInches() != 11 {
		t.Errorf(
			"Letter height in inches = %v, want 11",
			PageSizeLetter.HeightInches(),
		)
	}

	// Test A4 in mm
	if math.Abs(PageSizeA4.WidthMM()-210) > 0.1 {
		t.Errorf(
			"A4 width in mm = %v, want ~210",
			PageSizeA4.WidthMM(),
		)
	}
	if math.Abs(PageSizeA4.HeightMM()-297) > 0.1 {
		t.Errorf(
			"A4 height in mm = %v, want ~297",
			PageSizeA4.HeightMM(),
		)
	}

	// Test A4 in cm
	if math.Abs(PageSizeA4.WidthCM()-21) > 0.1 {
		t.Errorf(
			"A4 width in cm = %v, want ~21",
			PageSizeA4.WidthCM(),
		)
	}
	if math.Abs(
		PageSizeA4.HeightCM()-29.7,
	) > 0.1 {
		t.Errorf(
			"A4 height in cm = %v, want ~29.7",
			PageSizeA4.HeightCM(),
		)
	}
}

func TestMargins(t *testing.T) {
	t.Run("NewMargins", func(t *testing.T) {
		m := NewMargins(10, 20, 30, 40)
		if m.Top != 10 || m.Right != 20 ||
			m.Bottom != 30 ||
			m.Left != 40 {
			t.Errorf("NewMargins = %+v", m)
		}
	})

	t.Run(
		"NewMarginsFromDimensions",
		func(t *testing.T) {
			m := NewMarginsFromDimensions(
				Inches(1),
				Inches(0.5),
				Inches(1),
				Inches(0.5),
			)
			if m.Top != 72 || m.Right != 36 ||
				m.Bottom != 72 ||
				m.Left != 36 {
				t.Errorf(
					"NewMarginsFromDimensions = %+v",
					m,
				)
			}
		},
	)

	t.Run(
		"NewMarginsUniform",
		func(t *testing.T) {
			m := NewMarginsUniform(50)
			if m.Top != 50 || m.Right != 50 ||
				m.Bottom != 50 ||
				m.Left != 50 {
				t.Errorf(
					"NewMarginsUniform(50) = %+v",
					m,
				)
			}
		},
	)

	t.Run(
		"NewMarginsSymmetric",
		func(t *testing.T) {
			m := NewMarginsSymmetric(72, 36)
			if m.Top != 72 || m.Right != 36 ||
				m.Bottom != 72 ||
				m.Left != 36 {
				t.Errorf(
					"NewMarginsSymmetric(72, 36) = %+v",
					m,
				)
			}
		},
	)

	t.Run("NoMargins", func(t *testing.T) {
		m := NoMargins()
		if !m.IsZero() {
			t.Error("NoMargins should be zero")
		}
	})

	t.Run("DefaultMargins", func(t *testing.T) {
		m := DefaultMargins()
		oneInch := InchesToPoints(1)
		if m.Top != oneInch ||
			m.Right != oneInch ||
			m.Bottom != oneInch ||
			m.Left != oneInch {
			t.Errorf(
				"DefaultMargins = %+v, want all %v",
				m,
				oneInch,
			)
		}
	})

	t.Run("NarrowMargins", func(t *testing.T) {
		m := NarrowMargins()
		halfInch := InchesToPoints(0.5)
		if m.Top != halfInch ||
			m.Right != halfInch ||
			m.Bottom != halfInch ||
			m.Left != halfInch {
			t.Errorf(
				"NarrowMargins = %+v, want all %v",
				m,
				halfInch,
			)
		}
	})

	t.Run("WideMargins", func(t *testing.T) {
		m := WideMargins()
		if m.Top != InchesToPoints(1) ||
			m.Right != InchesToPoints(1.5) ||
			m.Bottom != InchesToPoints(1) ||
			m.Left != InchesToPoints(1.5) {
			t.Errorf("WideMargins = %+v", m)
		}
	})
}

func TestMargins_Totals(t *testing.T) {
	m := NewMargins(10, 20, 30, 40)

	if m.HorizontalTotal() != 60 {
		t.Errorf(
			"HorizontalTotal() = %v, want 60",
			m.HorizontalTotal(),
		)
	}
	if m.VerticalTotal() != 40 {
		t.Errorf(
			"VerticalTotal() = %v, want 40",
			m.VerticalTotal(),
		)
	}
}

func TestMargins_IsZero(t *testing.T) {
	if !NoMargins().IsZero() {
		t.Error("NoMargins.IsZero() = false")
	}
	if DefaultMargins().IsZero() {
		t.Error("DefaultMargins.IsZero() = true")
	}
}

func TestRectangle(t *testing.T) {
	t.Run("NewRectangle", func(t *testing.T) {
		r := NewRectangle(10, 20, 110, 220)
		if r.LLX != 10 || r.LLY != 20 ||
			r.URX != 110 ||
			r.URY != 220 {
			t.Errorf("NewRectangle = %+v", r)
		}
	})

	t.Run(
		"NewRectangleFromSize",
		func(t *testing.T) {
			r := NewRectangleFromSize(
				10,
				20,
				100,
				200,
			)
			if r.LLX != 10 || r.LLY != 20 ||
				r.URX != 110 ||
				r.URY != 220 {
				t.Errorf(
					"NewRectangleFromSize = %+v",
					r,
				)
			}
		},
	)

	t.Run(
		"NewRectangleForPageSize",
		func(t *testing.T) {
			r := NewRectangleForPageSize(
				PageSizeA4,
			)
			if r.LLX != 0 || r.LLY != 0 ||
				r.URX != PageSizeA4.Width ||
				r.URY != PageSizeA4.Height {
				t.Errorf(
					"NewRectangleForPageSize = %+v",
					r,
				)
			}
		},
	)
}

func TestRectangle_Dimensions(t *testing.T) {
	r := NewRectangle(10, 20, 110, 220)

	if r.Width() != 100 {
		t.Errorf(
			"Width() = %v, want 100",
			r.Width(),
		)
	}
	if r.Height() != 200 {
		t.Errorf(
			"Height() = %v, want 200",
			r.Height(),
		)
	}

	size := r.ToPageSize()
	if size.Width != 100 || size.Height != 200 {
		t.Errorf(
			"ToPageSize() = %+v, want {100, 200}",
			size,
		)
	}
}

func TestRectangle_Contains(t *testing.T) {
	r := NewRectangle(0, 0, 100, 100)

	tests := []struct {
		x, y float64
		want bool
	}{
		{50, 50, true},   // Center
		{0, 0, true},     // Corner
		{100, 100, true}, // Opposite corner
		{-1, 50, false},  // Outside left
		{101, 50, false}, // Outside right
		{50, -1, false},  // Outside bottom
		{50, 101, false}, // Outside top
	}

	for _, tt := range tests {
		if got := r.Contains(tt.x, tt.y); got != tt.want {
			t.Errorf(
				"Contains(%v, %v) = %v, want %v",
				tt.x,
				tt.y,
				got,
				tt.want,
			)
		}
	}
}

func TestRectangle_Inset(t *testing.T) {
	r := NewRectangle(0, 0, 100, 200)
	margins := NewMargins(10, 20, 30, 40)
	inset := r.Inset(margins)

	if inset.LLX != 40 { // Left margin
		t.Errorf(
			"Inset LLX = %v, want 40",
			inset.LLX,
		)
	}
	if inset.LLY != 30 { // Bottom margin
		t.Errorf(
			"Inset LLY = %v, want 30",
			inset.LLY,
		)
	}
	if inset.URX != 80 { // 100 - right margin
		t.Errorf(
			"Inset URX = %v, want 80",
			inset.URX,
		)
	}
	if inset.URY != 190 { // 200 - top margin
		t.Errorf(
			"Inset URY = %v, want 190",
			inset.URY,
		)
	}
}

func TestRectangle_Array(t *testing.T) {
	r := NewRectangle(0, 0, 612, 792)
	arr := r.Array()

	if len(arr) != 4 {
		t.Fatalf(
			"Array length = %d, want 4",
			len(arr),
		)
	}
}

func TestPageBoxes(t *testing.T) {
	t.Run("NewPageBoxes", func(t *testing.T) {
		media := NewRectangle(0, 0, 612, 792)
		boxes := NewPageBoxes(media)

		if boxes.MediaBox != media {
			t.Errorf(
				"MediaBox = %+v, want %+v",
				boxes.MediaBox,
				media,
			)
		}
		if boxes.CropBox != nil ||
			boxes.BleedBox != nil ||
			boxes.TrimBox != nil ||
			boxes.ArtBox != nil {
			t.Error(
				"Optional boxes should be nil",
			)
		}
	})

	t.Run(
		"NewPageBoxesFromSize",
		func(t *testing.T) {
			boxes := NewPageBoxesFromSize(
				PageSizeLetter,
			)
			if boxes.MediaBox.Width() != 612 ||
				boxes.MediaBox.Height() != 792 {
				t.Errorf(
					"MediaBox size = %vx%v, want 612x792",
					boxes.MediaBox.Width(),
					boxes.MediaBox.Height(),
				)
			}
		},
	)

	t.Run("WithBoxes", func(t *testing.T) {
		media := NewRectangle(0, 0, 612, 792)
		boxes := NewPageBoxes(media)

		crop := NewRectangle(10, 10, 602, 782)
		boxes = boxes.WithCropBox(crop)
		if boxes.CropBox == nil ||
			*boxes.CropBox != crop {
			t.Error("CropBox not set correctly")
		}

		bleed := NewRectangle(5, 5, 607, 787)
		boxes = boxes.WithBleedBox(bleed)
		if boxes.BleedBox == nil ||
			*boxes.BleedBox != bleed {
			t.Error("BleedBox not set correctly")
		}

		trim := NewRectangle(20, 20, 592, 772)
		boxes = boxes.WithTrimBox(trim)
		if boxes.TrimBox == nil ||
			*boxes.TrimBox != trim {
			t.Error("TrimBox not set correctly")
		}

		art := NewRectangle(30, 30, 582, 762)
		boxes = boxes.WithArtBox(art)
		if boxes.ArtBox == nil ||
			*boxes.ArtBox != art {
			t.Error("ArtBox not set correctly")
		}
	})

	t.Run("WithBleed", func(t *testing.T) {
		media := NewRectangle(0, 0, 612, 792)
		trim := NewRectangle(9, 9, 603, 783)
		boxes := NewPageBoxes(
			media,
		).WithTrimBox(trim).
			WithBleed(9)

		if boxes.BleedBox == nil {
			t.Fatal("BleedBox should be set")
		}
		// Bleed extends 9 points beyond trim
		if boxes.BleedBox.LLX != 0 ||
			boxes.BleedBox.LLY != 0 ||
			boxes.BleedBox.URX != 612 ||
			boxes.BleedBox.URY != 792 {
			t.Errorf(
				"BleedBox = %+v, expected to extend 9pt beyond trim",
				boxes.BleedBox,
			)
		}
	})
}

func TestPageBoxes_EffectiveBoxes(t *testing.T) {
	media := NewRectangle(0, 0, 612, 792)
	crop := NewRectangle(10, 10, 602, 782)

	t.Run(
		"defaults to MediaBox",
		func(t *testing.T) {
			boxes := NewPageBoxes(media)

			if boxes.EffectiveCropBox() != media {
				t.Error(
					"EffectiveCropBox should default to MediaBox",
				)
			}
			if boxes.EffectiveBleedBox() != media {
				t.Error(
					"EffectiveBleedBox should default to MediaBox",
				)
			}
			if boxes.EffectiveTrimBox() != media {
				t.Error(
					"EffectiveTrimBox should default to MediaBox",
				)
			}
			if boxes.EffectiveArtBox() != media {
				t.Error(
					"EffectiveArtBox should default to MediaBox",
				)
			}
		},
	)

	t.Run("uses set boxes", func(t *testing.T) {
		boxes := NewPageBoxes(
			media,
		).WithCropBox(crop)

		if boxes.EffectiveCropBox() != crop {
			t.Error(
				"EffectiveCropBox should use set CropBox",
			)
		}
		// Others should fall back to CropBox
		if boxes.EffectiveBleedBox() != crop {
			t.Error(
				"EffectiveBleedBox should fall back to CropBox",
			)
		}
		if boxes.EffectiveTrimBox() != crop {
			t.Error(
				"EffectiveTrimBox should fall back to CropBox",
			)
		}
		if boxes.EffectiveArtBox() != crop {
			t.Error(
				"EffectiveArtBox should fall back to CropBox",
			)
		}
	})
}

func TestPageOptions(t *testing.T) {
	t.Run(
		"DefaultPageOptions",
		func(t *testing.T) {
			opts := DefaultPageOptions()

			if opts.Size != PageSizeA4 {
				t.Errorf(
					"Default size = %+v, want A4",
					opts.Size,
				)
			}
			if opts.Margins.IsZero() {
				t.Error(
					"Default margins should not be zero",
				)
			}
			if opts.UserUnit != 1.0 {
				t.Errorf(
					"Default UserUnit = %v, want 1.0",
					opts.UserUnit,
				)
			}
		},
	)

	t.Run("NewPageOptions", func(t *testing.T) {
		opts := NewPageOptions(PageSizeLetter)

		if opts.Size != PageSizeLetter {
			t.Errorf(
				"Size = %+v, want Letter",
				opts.Size,
			)
		}
	})

	t.Run("WithMargins", func(t *testing.T) {
		margins := NewMarginsUniform(36)
		opts := NewPageOptions(
			PageSizeA4,
		).WithMargins(margins)

		if opts.Margins != margins {
			t.Errorf(
				"Margins = %+v, want %+v",
				opts.Margins,
				margins,
			)
		}
	})

	t.Run("WithNoMargins", func(t *testing.T) {
		opts := NewPageOptions(
			PageSizeA4,
		).WithNoMargins()

		if !opts.Margins.IsZero() {
			t.Error("Margins should be zero")
		}
	})

	t.Run("WithRotation", func(t *testing.T) {
		tests := []struct {
			input int
			want  int
		}{
			{0, 0},
			{90, 90},
			{180, 180},
			{270, 270},
			{360, 0},
			{45, 90},   // Rounds up
			{135, 180}, // Rounds up
			{-90, 270}, // Negative
			{450, 90},  // > 360
		}

		for _, tt := range tests {
			opts := NewPageOptions(
				PageSizeA4,
			).WithRotation(tt.input)
			if opts.Rotation != tt.want {
				t.Errorf(
					"WithRotation(%d) = %d, want %d",
					tt.input,
					opts.Rotation,
					tt.want,
				)
			}
		}
	})

	t.Run("WithUserUnit", func(t *testing.T) {
		opts := NewPageOptions(
			PageSizeA4,
		).WithUserUnit(2.0)
		if opts.UserUnit != 2.0 {
			t.Errorf(
				"UserUnit = %v, want 2.0",
				opts.UserUnit,
			)
		}

		// Test clamping
		opts = NewPageOptions(
			PageSizeA4,
		).WithUserUnit(0.5)
		if opts.UserUnit != 1.0 {
			t.Errorf(
				"UserUnit should be clamped to 1.0, got %v",
				opts.UserUnit,
			)
		}

		opts = NewPageOptions(
			PageSizeA4,
		).WithUserUnit(100000)
		if opts.UserUnit != 75000 {
			t.Errorf(
				"UserUnit should be clamped to 75000, got %v",
				opts.UserUnit,
			)
		}
	})

	t.Run(
		"Orientation methods",
		func(t *testing.T) {
			opts := NewPageOptions(
				PageSizeA4,
			).Landscape()
			if !opts.Size.IsLandscape() {
				t.Error(
					"Landscape() should produce landscape",
				)
			}

			opts = opts.Portrait()
			if !opts.Size.IsPortrait() {
				t.Error(
					"Portrait() should produce portrait",
				)
			}
		},
	)
}

func TestPageOptions_ContentArea(t *testing.T) {
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(NewMargins(72, 72, 72, 72))

	area := opts.ContentArea()
	if area.LLX != 72 {
		t.Errorf(
			"ContentArea.LLX = %v, want 72",
			area.LLX,
		)
	}
	if area.LLY != 72 {
		t.Errorf(
			"ContentArea.LLY = %v, want 72",
			area.LLY,
		)
	}
	if area.URX != 540 { // 612 - 72
		t.Errorf(
			"ContentArea.URX = %v, want 540",
			area.URX,
		)
	}
	if area.URY != 720 { // 792 - 72
		t.Errorf(
			"ContentArea.URY = %v, want 720",
			area.URY,
		)
	}

	if opts.ContentWidth() != 468 { // 612 - 144
		t.Errorf(
			"ContentWidth = %v, want 468",
			opts.ContentWidth(),
		)
	}
	if opts.ContentHeight() != 648 { // 792 - 144
		t.Errorf(
			"ContentHeight = %v, want 648",
			opts.ContentHeight(),
		)
	}
}

func TestCoordinateSystem(t *testing.T) {
	pageHeight := 792.0 // Letter height

	t.Run(
		"default bottom-left origin",
		func(t *testing.T) {
			cs := NewCoordinateSystem(pageHeight)

			if cs.Origin != OriginBottomLeft {
				t.Error(
					"Default origin should be BottomLeft",
				)
			}
			if cs.UserUnit != 1.0 {
				t.Errorf(
					"Default UserUnit = %v, want 1.0",
					cs.UserUnit,
				)
			}

			// Y should not be transformed
			y := cs.TransformY(100)
			if y != 100 {
				t.Errorf(
					"TransformY(100) = %v, want 100",
					y,
				)
			}
		},
	)

	t.Run("top-left origin", func(t *testing.T) {
		cs := NewCoordinateSystem(
			pageHeight,
		).WithTopLeftOrigin()

		if cs.Origin != OriginTopLeft {
			t.Error("Origin should be TopLeft")
		}

		// Y should be flipped
		y := cs.TransformY(100)
		if y != 692 { // 792 - 100
			t.Errorf(
				"TransformY(100) = %v, want 692",
				y,
			)
		}

		// Point at top-left corner (0,0) should map to (0, pageHeight)
		x, y := cs.TransformPoint(0, 0)
		if x != 0 || y != 792 {
			t.Errorf(
				"TransformPoint(0,0) = (%v,%v), want (0,792)",
				x,
				y,
			)
		}
	})

	t.Run("user units", func(t *testing.T) {
		cs := NewCoordinateSystem(
			pageHeight,
		).WithUserUnit(2.0)

		// 72 points in user units = 36
		if cs.ScaleToUserUnits(72) != 36 {
			t.Errorf(
				"ScaleToUserUnits(72) = %v, want 36",
				cs.ScaleToUserUnits(72),
			)
		}

		// 36 user units = 72 points
		if cs.ScaleFromUserUnits(36) != 72 {
			t.Errorf(
				"ScaleFromUserUnits(36) = %v, want 72",
				cs.ScaleFromUserUnits(36),
			)
		}
	})
}

func TestPageSizeLedger(t *testing.T) {
	// Ledger is landscape Tabloid (17x11 inches)
	if PageSizeLedger.Width != PageSizeTabloid.Height {
		t.Errorf(
			"Ledger width = %v, want %v (Tabloid height)",
			PageSizeLedger.Width,
			PageSizeTabloid.Height,
		)
	}
	if PageSizeLedger.Height != PageSizeTabloid.Width {
		t.Errorf(
			"Ledger height = %v, want %v (Tabloid width)",
			PageSizeLedger.Height,
			PageSizeTabloid.Width,
		)
	}
	if !PageSizeLedger.IsLandscape() {
		t.Error("Ledger should be landscape")
	}
}

// TestDocumentWithPageOptions tests that page options work with document creation
func TestDocumentWithPageOptions(t *testing.T) {
	doc, err := NewDocument()
	if err != nil {
		t.Fatalf("NewDocument() error = %v", err)
	}
	defer func() { _ = doc.Close() }()

	// Test creating page with custom size
	customSize := NewPageSizeFromUnit(
		200,
		300,
		UnitMM,
	)
	page, err := doc.AddPage(customSize)
	if err != nil {
		t.Fatalf("AddPage() error = %v", err)
	}

	// Verify dimensions
	if math.Abs(
		page.Width()-MMToPoints(200),
	) > 0.1 {
		t.Errorf(
			"Page width = %v, want %v",
			page.Width(),
			MMToPoints(200),
		)
	}
	if math.Abs(
		page.Height()-MMToPoints(300),
	) > 0.1 {
		t.Errorf(
			"Page height = %v, want %v",
			page.Height(),
			MMToPoints(300),
		)
	}
}

// TestAllStandardPageSizes ensures all standard page sizes are valid
func TestAllStandardPageSizes(t *testing.T) {
	sizes := map[string]PageSize{
		"A0":        PageSizeA0,
		"A1":        PageSizeA1,
		"A2":        PageSizeA2,
		"A3":        PageSizeA3,
		"A4":        PageSizeA4,
		"A5":        PageSizeA5,
		"A6":        PageSizeA6,
		"B4":        PageSizeB4,
		"B5":        PageSizeB5,
		"Letter":    PageSizeLetter,
		"Legal":     PageSizeLegal,
		"Tabloid":   PageSizeTabloid,
		"Ledger":    PageSizeLedger,
		"Executive": PageSizeExecutive,
		"Statement": PageSizeStatement,
		"Folio":     PageSizeFolio,
		"Quarto":    PageSizeQuarto,
	}

	for name, size := range sizes {
		t.Run(name, func(t *testing.T) {
			if size.Width <= 0 {
				t.Errorf(
					"%s has invalid width: %v",
					name,
					size.Width,
				)
			}
			if size.Height <= 0 {
				t.Errorf(
					"%s has invalid height: %v",
					name,
					size.Height,
				)
			}

			// Most standard sizes should be portrait (except Ledger)
			if name != "Ledger" &&
				size.Width > size.Height {
				t.Errorf(
					"%s should be portrait but width > height",
					name,
				)
			}
		})
	}
}
