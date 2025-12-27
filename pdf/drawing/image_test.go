package drawing

import (
	"bytes"
	"compress/zlib"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"testing"
)

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected ImageFormat
	}{
		{
			name: "JPEG",
			data: []byte{
				0xFF,
				0xD8,
				0xFF,
				0xE0,
				0x00,
				0x10,
				0x4A,
				0x46,
			},
			expected: ImageFormatJPEG,
		},
		{
			name: "PNG",
			data: []byte{
				0x89,
				0x50,
				0x4E,
				0x47,
				0x0D,
				0x0A,
				0x1A,
				0x0A,
			},
			expected: ImageFormatPNG,
		},
		{
			name: "GIF87a",
			data: []byte{
				0x47,
				0x49,
				0x46,
				0x38,
				0x37,
				0x61,
				0x00,
				0x00,
			},
			expected: ImageFormatGIF,
		},
		{
			name: "GIF89a",
			data: []byte{
				0x47,
				0x49,
				0x46,
				0x38,
				0x39,
				0x61,
				0x00,
				0x00,
			},
			expected: ImageFormatGIF,
		},
		{
			name: "BMP",
			data: []byte{
				0x42,
				0x4D,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
			expected: ImageFormatBMP,
		},
		{
			name: "Unknown",
			data: []byte{
				0x00,
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
			},
			expected: ImageFormatUnknown,
		},
		{
			name:     "Too short",
			data:     []byte{0xFF, 0xD8},
			expected: ImageFormatUnknown,
		},
		{
			name:     "Empty",
			data:     []byte{},
			expected: ImageFormatUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DetectFormat(tt.data)
			if result != tt.expected {
				t.Errorf(
					"DetectFormat() = %v, want %v",
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestImageFormatString(t *testing.T) {
	tests := []struct {
		format   ImageFormat
		expected string
	}{
		{ImageFormatJPEG, "JPEG"},
		{ImageFormatPNG, "PNG"},
		{ImageFormatGIF, "GIF"},
		{ImageFormatBMP, "BMP"},
		{ImageFormatUnknown, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.format.String() != tt.expected {
				t.Errorf(
					"ImageFormat.String() = %v, want %v",
					tt.format.String(),
					tt.expected,
				)
			}
		})
	}
}

func TestColorSpaceString(t *testing.T) {
	tests := []struct {
		cs       ColorSpace
		expected string
	}{
		{ColorSpaceDeviceRGB, "DeviceRGB"},
		{ColorSpaceDeviceGray, "DeviceGray"},
		{ColorSpaceDeviceCMYK, "DeviceCMYK"},
		{ColorSpaceIndexed, "Indexed"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.cs.String() != tt.expected {
				t.Errorf(
					"ColorSpace.String() = %v, want %v",
					tt.cs.String(),
					tt.expected,
				)
			}
		})
	}
}

func TestNewImageXObject(t *testing.T) {
	img := NewImageXObject("Im1")

	if img.Name != "Im1" {
		t.Errorf("Name = %v, want Im1", img.Name)
	}
	if img.BitsPerComponent != 8 {
		t.Errorf(
			"BitsPerComponent = %v, want 8",
			img.BitsPerComponent,
		)
	}
	if img.ColorSpace != ColorSpaceDeviceRGB {
		t.Errorf(
			"ColorSpace = %v, want DeviceRGB",
			img.ColorSpace,
		)
	}
	if img.ColorComponents != 3 {
		t.Errorf(
			"ColorComponents = %v, want 3",
			img.ColorComponents,
		)
	}
	if !img.Interpolate {
		t.Error(
			"Interpolate should be true by default",
		)
	}
}

func TestImageXObjectAspectRatio(t *testing.T) {
	img := NewImageXObject("test")
	img.Width = 200
	img.Height = 100

	ratio := img.AspectRatio()
	if ratio != 2.0 {
		t.Errorf(
			"AspectRatio() = %v, want 2.0",
			ratio,
		)
	}

	// Test zero height
	img.Height = 0
	ratio = img.AspectRatio()
	if ratio != 0 {
		t.Errorf(
			"AspectRatio() with zero height = %v, want 0",
			ratio,
		)
	}
}

func TestImageXObjectHasAlpha(t *testing.T) {
	img := NewImageXObject("test")

	if img.HasAlpha() {
		t.Error(
			"HasAlpha() should be false without SMask",
		)
	}

	img.SMask = NewImageXObject("smask")
	if !img.HasAlpha() {
		t.Error(
			"HasAlpha() should be true with SMask",
		)
	}
}

func TestImageXObjectDictionary(t *testing.T) {
	img := NewImageXObject("test")
	img.Width = 100
	img.Height = 100
	img.Filter = "DCTDecode"

	dict := img.Dictionary()

	if dict["Type"] != "/XObject" {
		t.Errorf(
			"Type = %v, want /XObject",
			dict["Type"],
		)
	}
	if dict["Subtype"] != "/Image" {
		t.Errorf(
			"Subtype = %v, want /Image",
			dict["Subtype"],
		)
	}
	if dict["Width"] != 100 {
		t.Errorf(
			"Width = %v, want 100",
			dict["Width"],
		)
	}
	if dict["Height"] != 100 {
		t.Errorf(
			"Height = %v, want 100",
			dict["Height"],
		)
	}
	if dict["Filter"] != "/DCTDecode" {
		t.Errorf(
			"Filter = %v, want /DCTDecode",
			dict["Filter"],
		)
	}
}

func TestImageManager(t *testing.T) {
	mgr := NewImageManager()

	if mgr.Count() != 0 {
		t.Errorf(
			"Count() = %v, want 0",
			mgr.Count(),
		)
	}

	name1 := mgr.NextName()
	if name1 != "Im1" {
		t.Errorf(
			"NextName() = %v, want Im1",
			name1,
		)
	}

	name2 := mgr.NextName()
	if name2 != "Im2" {
		t.Errorf(
			"NextName() = %v, want Im2",
			name2,
		)
	}

	img := NewImageXObject("Im1")
	mgr.Add(img)

	if mgr.Count() != 1 {
		t.Errorf(
			"Count() after Add = %v, want 1",
			mgr.Count(),
		)
	}

	retrieved, ok := mgr.Get("Im1")
	if !ok {
		t.Error(
			"Get() should return true for existing image",
		)
	}
	if retrieved != img {
		t.Error("Get() returned different image")
	}

	_, ok = mgr.Get("Im999")
	if ok {
		t.Error(
			"Get() should return false for non-existing image",
		)
	}
}

func TestImageManagerSetPrefix(t *testing.T) {
	mgr := NewImageManager()
	mgr.SetPrefix("Img")

	name := mgr.NextName()
	if name != "Img1" {
		t.Errorf(
			"NextName() with custom prefix = %v, want Img1",
			name,
		)
	}
}

func TestImageManagerAutoName(t *testing.T) {
	mgr := NewImageManager()

	img := NewImageXObject("")
	mgr.Add(img)

	if img.Name != "Im1" {
		t.Errorf(
			"Auto-generated name = %v, want Im1",
			img.Name,
		)
	}
}

func createTestJPEG(width, height int) []byte {
	img := image.NewRGBA(
		image.Rect(0, 0, width, height),
	)

	// Fill with a simple pattern
	for y := range height {
		for x := range width {
			img.Set(
				x,
				y,
				color.RGBA{
					uint8(x % 256),
					uint8(y % 256),
					128,
					255,
				},
			)
		}
	}

	var buf bytes.Buffer
	_ = jpeg.Encode(
		&buf,
		img,
		&jpeg.Options{Quality: 90},
	)

	return buf.Bytes()
}

func createTestPNG(
	width, height int,
	withAlpha bool,
) []byte {
	var img image.Image
	if withAlpha {
		rgba := image.NewRGBA(
			image.Rect(0, 0, width, height),
		)
		for y := range height {
			for x := range width {
				rgba.Set(
					x,
					y,
					color.RGBA{
						uint8(x % 256),
						uint8(y % 256),
						128,
						uint8((x + y) % 256),
					},
				)
			}
		}
		img = rgba
	} else {
		nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
		for y := range height {
			for x := range width {
				nrgba.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 128, 255})
			}
		}
		img = nrgba
	}

	var buf bytes.Buffer
	_ = png.Encode(&buf, img)

	return buf.Bytes()
}

func TestLoadJPEG(t *testing.T) {
	jpegData := createTestJPEG(100, 80)

	img, err := LoadJPEG(jpegData, "test")
	if err != nil {
		t.Fatalf("LoadJPEG() error = %v", err)
	}

	if img.Width != 100 {
		t.Errorf(
			"Width = %v, want 100",
			img.Width,
		)
	}
	if img.Height != 80 {
		t.Errorf(
			"Height = %v, want 80",
			img.Height,
		)
	}
	if img.Filter != "DCTDecode" {
		t.Errorf(
			"Filter = %v, want DCTDecode",
			img.Filter,
		)
	}
	if img.OriginalFormat != ImageFormatJPEG {
		t.Errorf(
			"OriginalFormat = %v, want JPEG",
			img.OriginalFormat,
		)
	}

	// Verify data is passed through unchanged
	if !bytes.Equal(img.Data, jpegData) {
		t.Error(
			"JPEG data should be passed through unchanged",
		)
	}
}

func TestLoadPNG(t *testing.T) {
	pngData := createTestPNG(50, 60, false)

	img, err := LoadPNG(pngData, "test")
	if err != nil {
		t.Fatalf("LoadPNG() error = %v", err)
	}

	if img.Width != 50 {
		t.Errorf("Width = %v, want 50", img.Width)
	}
	if img.Height != 60 {
		t.Errorf(
			"Height = %v, want 60",
			img.Height,
		)
	}
	if img.Filter != "FlateDecode" {
		t.Errorf(
			"Filter = %v, want FlateDecode",
			img.Filter,
		)
	}
	if img.OriginalFormat != ImageFormatPNG {
		t.Errorf(
			"OriginalFormat = %v, want PNG",
			img.OriginalFormat,
		)
	}
}

func TestLoadPNGWithAlpha(t *testing.T) {
	pngData := createTestPNG(40, 40, true)

	img, err := LoadPNG(pngData, "test")
	if err != nil {
		t.Fatalf("LoadPNG() error = %v", err)
	}

	if !img.HasAlpha() {
		t.Error("Image should have alpha channel")
	}
	if img.SMask == nil {
		t.Error(
			"SMask should not be nil for image with alpha",
		)
	}
	if img.SMask.Width != 40 {
		t.Errorf(
			"SMask Width = %v, want 40",
			img.SMask.Width,
		)
	}
	if img.SMask.ColorSpace != ColorSpaceDeviceGray {
		t.Errorf(
			"SMask ColorSpace = %v, want DeviceGray",
			img.SMask.ColorSpace,
		)
	}
}

func TestLoadImage(t *testing.T) {
	t.Run("JPEG", func(t *testing.T) {
		data := createTestJPEG(100, 100)
		img, err := LoadImage(data, "test")
		if err != nil {
			t.Fatalf(
				"LoadImage(JPEG) error = %v",
				err,
			)
		}
		if img.OriginalFormat != ImageFormatJPEG {
			t.Errorf(
				"Format = %v, want JPEG",
				img.OriginalFormat,
			)
		}
	})

	t.Run("PNG", func(t *testing.T) {
		data := createTestPNG(100, 100, false)
		img, err := LoadImage(data, "test")
		if err != nil {
			t.Fatalf(
				"LoadImage(PNG) error = %v",
				err,
			)
		}
		if img.OriginalFormat != ImageFormatPNG {
			t.Errorf(
				"Format = %v, want PNG",
				img.OriginalFormat,
			)
		}
	})

	t.Run("Empty data", func(t *testing.T) {
		_, err := LoadImage([]byte{}, "test")
		if err != ErrNoImageData {
			t.Errorf(
				"LoadImage(empty) error = %v, want ErrNoImageData",
				err,
			)
		}
	})
}

func TestLoadFromGoImage(t *testing.T) {
	goImg := image.NewRGBA(
		image.Rect(0, 0, 32, 32),
	)
	for y := range 32 {
		for x := range 32 {
			goImg.Set(
				x,
				y,
				color.RGBA{255, 128, 64, 255},
			)
		}
	}

	img, err := LoadFromGoImage(goImg, "test")
	if err != nil {
		t.Fatalf(
			"LoadFromGoImage() error = %v",
			err,
		)
	}

	if img.Width != 32 {
		t.Errorf("Width = %v, want 32", img.Width)
	}
	if img.Height != 32 {
		t.Errorf(
			"Height = %v, want 32",
			img.Height,
		)
	}
}

func TestDrawImageOperator(t *testing.T) {
	result := DrawImageOperator(
		"Im1",
		100,
		200,
		300,
		400,
	)

	if result == "" {
		t.Error(
			"DrawImageOperator should not return empty string",
		)
	}

	// Check for expected operators
	expected := []string{
		"q\n",
		"300 0 0 400 100 200 cm\n",
		"/Im1 Do\n",
		"Q\n",
	}

	for _, exp := range expected {
		if !bytes.Contains(
			[]byte(result),
			[]byte(exp),
		) {
			t.Errorf(
				"Result missing expected operator: %q",
				exp,
			)
		}
	}
}

func TestDrawImage(t *testing.T) {
	result := DrawImage(
		10,
		20,
		100,
		50,
		"TestImg",
	)

	if result == "" {
		t.Error(
			"DrawImage should not return empty string",
		)
	}
	if !bytes.Contains(
		[]byte(result),
		[]byte("/TestImg Do"),
	) {
		t.Error(
			"Result should contain image Do operator",
		)
	}
}

func TestDrawImageAt(t *testing.T) {
	img := NewImageXObject("Im1")
	img.Width = 200
	img.Height = 100

	result := DrawImageAt(50, 60, img)

	if !bytes.Contains(
		[]byte(result),
		[]byte("/Im1 Do"),
	) {
		t.Error(
			"Result should contain image Do operator",
		)
	}
	// Should use original dimensions
	if !bytes.Contains(
		[]byte(result),
		[]byte("200 0 0 100 50 60 cm"),
	) {
		t.Error(
			"Result should use original image dimensions",
		)
	}
}

func TestDrawImageFit(t *testing.T) {
	img := NewImageXObject("Im1")
	img.Width = 200
	img.Height = 100

	rect := Rect{
		X:      0,
		Y:      0,
		Width:  100,
		Height: 100,
	}

	t.Run("FitModeContain", func(t *testing.T) {
		result := DrawImageFit(
			rect,
			img,
			FitModeContain,
		)
		if result == "" {
			t.Error(
				"DrawImageFit should not return empty string",
			)
		}
	})

	t.Run("FitModeCover", func(t *testing.T) {
		result := DrawImageFit(
			rect,
			img,
			FitModeCover,
		)
		if result == "" {
			t.Error(
				"DrawImageFit should not return empty string",
			)
		}
	})

	t.Run("FitModeStretch", func(t *testing.T) {
		result := DrawImageFit(
			rect,
			img,
			FitModeStretch,
		)
		if result == "" {
			t.Error(
				"DrawImageFit should not return empty string",
			)
		}
	})

	t.Run("FitModeNone", func(t *testing.T) {
		result := DrawImageFit(
			rect,
			img,
			FitModeNone,
		)
		if result == "" {
			t.Error(
				"DrawImageFit should not return empty string",
			)
		}
	})

	t.Run("Nil image", func(t *testing.T) {
		result := DrawImageFit(
			rect,
			nil,
			FitModeContain,
		)
		if result != "" {
			t.Error(
				"DrawImageFit with nil image should return empty string",
			)
		}
	})
}

func TestImageBounds(t *testing.T) {
	img := NewImageXObject("Im1")
	img.Width = 200
	img.Height = 100

	rect := Rect{
		X:      0,
		Y:      0,
		Width:  100,
		Height: 100,
	}

	t.Run("FitModeContain", func(t *testing.T) {
		bounds := ImageBounds(
			rect,
			img,
			FitModeContain,
		)
		// Image is 2:1 aspect, rect is 1:1
		// Should fit to width, resulting in 100x50 centered
		if bounds.Width != 100 {
			t.Errorf(
				"Width = %v, want 100",
				bounds.Width,
			)
		}
		if bounds.Height != 50 {
			t.Errorf(
				"Height = %v, want 50",
				bounds.Height,
			)
		}
		if bounds.Y != 25 {
			t.Errorf(
				"Y = %v, want 25 (centered)",
				bounds.Y,
			)
		}
	})

	t.Run("FitModeStretch", func(t *testing.T) {
		bounds := ImageBounds(
			rect,
			img,
			FitModeStretch,
		)
		if bounds != rect {
			t.Errorf(
				"Bounds = %v, want %v",
				bounds,
				rect,
			)
		}
	})

	t.Run("FitModeNone", func(t *testing.T) {
		bounds := ImageBounds(
			rect,
			img,
			FitModeNone,
		)
		if bounds.Width != 200 {
			t.Errorf(
				"Width = %v, want 200",
				bounds.Width,
			)
		}
		if bounds.Height != 100 {
			t.Errorf(
				"Height = %v, want 100",
				bounds.Height,
			)
		}
	})
}

func TestTransparencyGroup(t *testing.T) {
	tg := DefaultTransparencyGroup()

	if tg.ColorSpace != ColorSpaceDeviceRGB {
		t.Errorf(
			"ColorSpace = %v, want DeviceRGB",
			tg.ColorSpace,
		)
	}

	dict := tg.Dictionary()
	if dict["Type"] != "/Group" {
		t.Errorf(
			"Type = %v, want /Group",
			dict["Type"],
		)
	}
	if dict["S"] != "/Transparency" {
		t.Errorf(
			"S = %v, want /Transparency",
			dict["S"],
		)
	}
}

func TestBeginEndTransparencyGroup(t *testing.T) {
	begin := BeginTransparencyGroup()
	end := EndTransparencyGroup()

	if begin != "q\n" {
		t.Errorf(
			"BeginTransparencyGroup() = %q, want %q",
			begin,
			"q\n",
		)
	}
	if end != "Q\n" {
		t.Errorf(
			"EndTransparencyGroup() = %q, want %q",
			end,
			"Q\n",
		)
	}
}

func TestSetImageAlpha(t *testing.T) {
	// Full opacity should return empty
	result := SetImageAlpha(1.0)
	if result != "" {
		t.Errorf(
			"SetImageAlpha(1.0) = %q, want empty",
			result,
		)
	}

	// Partial opacity should set graphics state
	result = SetImageAlpha(0.5)
	if result == "" {
		t.Error(
			"SetImageAlpha(0.5) should not return empty",
		)
	}
	if !bytes.Contains(
		[]byte(result),
		[]byte("/GS_A50 gs"),
	) {
		t.Errorf(
			"SetImageAlpha(0.5) = %q, should contain /GS_A50 gs",
			result,
		)
	}
}

func TestImageAlphaGraphicsState(t *testing.T) {
	gs := ImageAlphaGraphicsState(0.75)

	if gs["Type"] != "/ExtGState" {
		t.Errorf(
			"Type = %v, want /ExtGState",
			gs["Type"],
		)
	}
	if gs["ca"] != 0.75 {
		t.Errorf("ca = %v, want 0.75", gs["ca"])
	}
	if gs["CA"] != 0.75 {
		t.Errorf("CA = %v, want 0.75", gs["CA"])
	}
}

func TestInlineImage(t *testing.T) {
	ii := NewInlineImage(10, 10)

	if ii.Width != 10 {
		t.Errorf("Width = %v, want 10", ii.Width)
	}
	if ii.Height != 10 {
		t.Errorf(
			"Height = %v, want 10",
			ii.Height,
		)
	}
	if ii.BitsPerComponent != 8 {
		t.Errorf(
			"BitsPerComponent = %v, want 8",
			ii.BitsPerComponent,
		)
	}

	// Test content stream generation
	ii.Data = []byte{0, 0, 0} // Minimal data
	ii.Filter = "FlateDecode"
	result := ii.ContentStream()

	if !bytes.Contains(
		[]byte(result),
		[]byte("BI"),
	) {
		t.Error(
			"Result should contain BI operator",
		)
	}
	if !bytes.Contains(
		[]byte(result),
		[]byte("EI"),
	) {
		t.Error(
			"Result should contain EI operator",
		)
	}
	if !bytes.Contains(
		[]byte(result),
		[]byte("/W 10"),
	) {
		t.Error("Result should contain width")
	}
	if !bytes.Contains(
		[]byte(result),
		[]byte("/H 10"),
	) {
		t.Error("Result should contain height")
	}
}

func TestParseJPEGInfo(t *testing.T) {
	jpegData := createTestJPEG(100, 80)

	info, err := ParseJPEGInfo(jpegData)
	if err != nil {
		t.Fatalf(
			"ParseJPEGInfo() error = %v",
			err,
		)
	}

	if info.Width != 100 {
		t.Errorf(
			"Width = %v, want 100",
			info.Width,
		)
	}
	if info.Height != 80 {
		t.Errorf(
			"Height = %v, want 80",
			info.Height,
		)
	}
	if info.Components != 3 {
		t.Errorf(
			"Components = %v, want 3",
			info.Components,
		)
	}
	if info.ColorSpace != ColorSpaceDeviceRGB {
		t.Errorf(
			"ColorSpace = %v, want DeviceRGB",
			info.ColorSpace,
		)
	}

	// Test invalid data
	_, err = ParseJPEGInfo(
		[]byte{0x00, 0x00, 0x00},
	)
	if err != ErrInvalidImageData {
		t.Errorf(
			"ParseJPEGInfo(invalid) error = %v, want ErrInvalidImageData",
			err,
		)
	}
}

func TestParsePNGInfo(t *testing.T) {
	pngData := createTestPNG(50, 60, true)

	info, err := ParsePNGInfo(pngData)
	if err != nil {
		t.Fatalf("ParsePNGInfo() error = %v", err)
	}

	if info.Width != 50 {
		t.Errorf(
			"Width = %v, want 50",
			info.Width,
		)
	}
	if info.Height != 60 {
		t.Errorf(
			"Height = %v, want 60",
			info.Height,
		)
	}
	if !info.HasAlpha {
		t.Error(
			"HasAlpha should be true for RGBA image",
		)
	}

	// Test invalid data
	_, err = ParsePNGInfo(
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
	)
	if err != ErrInvalidImageData {
		t.Errorf(
			"ParsePNGInfo(invalid) error = %v, want ErrInvalidImageData",
			err,
		)
	}
}

func TestImageReader(t *testing.T) {
	jpegData := createTestJPEG(32, 32)
	reader := bytes.NewReader(jpegData)

	ir := NewImageReader(reader)
	img, err := ir.Read("test")
	if err != nil {
		t.Fatalf(
			"ImageReader.Read() error = %v",
			err,
		)
	}

	if img.Width != 32 {
		t.Errorf("Width = %v, want 32", img.Width)
	}
	if img.Height != 32 {
		t.Errorf(
			"Height = %v, want 32",
			img.Height,
		)
	}
}

func TestScaleImageDimensions(t *testing.T) {
	tests := []struct {
		name                 string
		width, height        int
		maxW, maxH           float64
		expectedW, expectedH float64
	}{
		{
			name:  "No constraints",
			width: 200, height: 100,
			maxW: 0, maxH: 0,
			expectedW: 200, expectedH: 100,
		},
		{
			name:  "Width constraint only",
			width: 200, height: 100,
			maxW: 100, maxH: 0,
			expectedW: 100, expectedH: 50,
		},
		{
			name:  "Height constraint only",
			width: 200, height: 100,
			maxW: 0, maxH: 50,
			expectedW: 100, expectedH: 50,
		},
		{
			name:  "Both constraints, width limiting",
			width: 400, height: 100,
			maxW: 200, maxH: 100,
			expectedW: 200, expectedH: 50,
		},
		{
			name:  "Both constraints, height limiting",
			width: 100, height: 400,
			maxW: 200, maxH: 100,
			expectedW: 25, expectedH: 100,
		},
		{
			name:  "Already within constraints",
			width: 100, height: 50,
			maxW: 200, maxH: 100,
			expectedW: 100, expectedH: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, h := ScaleImageDimensions(
				tt.width,
				tt.height,
				tt.maxW,
				tt.maxH,
			)
			if w != tt.expectedW {
				t.Errorf(
					"Width = %v, want %v",
					w,
					tt.expectedW,
				)
			}
			if h != tt.expectedH {
				t.Errorf(
					"Height = %v, want %v",
					h,
					tt.expectedH,
				)
			}
		})
	}
}

func TestImageDataCompression(t *testing.T) {
	// Create a simple RGB image
	goImg := image.NewRGBA(
		image.Rect(0, 0, 16, 16),
	)
	for y := range 16 {
		for x := range 16 {
			goImg.Set(
				x,
				y,
				color.RGBA{255, 0, 0, 255},
			)
		}
	}

	img, err := LoadFromGoImage(goImg, "test")
	if err != nil {
		t.Fatalf(
			"LoadFromGoImage() error = %v",
			err,
		)
	}

	// Verify data is compressed
	if img.Filter != "FlateDecode" {
		t.Errorf(
			"Filter = %v, want FlateDecode",
			img.Filter,
		)
	}

	// Decompress and verify
	r, err := zlib.NewReader(
		bytes.NewReader(img.Data),
	)
	if err != nil {
		t.Fatalf(
			"zlib.NewReader() error = %v",
			err,
		)
	}

	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(r)
	if err != nil {
		t.Fatalf("Decompress error = %v", err)
	}
	_ = r.Close()

	// Should have 16*16*3 = 768 bytes of RGB data
	if decompressed.Len() != 768 {
		t.Errorf(
			"Decompressed size = %v, want 768",
			decompressed.Len(),
		)
	}
}

func TestExtractImageData(t *testing.T) {
	goImg := image.NewRGBA(image.Rect(0, 0, 2, 2))
	goImg.Set(0, 0, color.RGBA{255, 0, 0, 128})
	goImg.Set(1, 0, color.RGBA{0, 255, 0, 64})
	goImg.Set(0, 1, color.RGBA{0, 0, 255, 255})
	goImg.Set(1, 1, color.RGBA{255, 255, 255, 0})

	rgb, alpha := extractImageData(goImg, true)

	// Check RGB data length
	if len(rgb) != 12 { // 2x2x3
		t.Errorf(
			"RGB length = %v, want 12",
			len(rgb),
		)
	}

	// Check alpha data length
	if len(alpha) != 4 { // 2x2
		t.Errorf(
			"Alpha length = %v, want 4",
			len(alpha),
		)
	}

	// Check first pixel RGB
	if rgb[0] != 255 || rgb[1] != 0 ||
		rgb[2] != 0 {
		t.Errorf(
			"First pixel RGB = (%v, %v, %v), want (255, 0, 0)",
			rgb[0],
			rgb[1],
			rgb[2],
		)
	}

	// Check alpha values
	if alpha[0] != 128 {
		t.Errorf(
			"First pixel alpha = %v, want 128",
			alpha[0],
		)
	}
}

func TestImageHasAlpha(t *testing.T) {
	t.Run("RGBA has alpha", func(t *testing.T) {
		img := image.NewRGBA(
			image.Rect(0, 0, 10, 10),
		)
		if !imageHasAlpha(img) {
			t.Error(
				"RGBA image should have alpha",
			)
		}
	})

	t.Run("RGB has no alpha", func(t *testing.T) {
		// Use YCbCr which doesn't have alpha
		img := image.NewYCbCr(
			image.Rect(0, 0, 10, 10),
			image.YCbCrSubsampleRatio444,
		)
		if imageHasAlpha(img) {
			t.Error(
				"YCbCr image should not have alpha",
			)
		}
	})

	t.Run(
		"Paletted with alpha",
		func(t *testing.T) {
			palette := []color.Color{
				color.RGBA{255, 0, 0, 255},
				color.RGBA{
					0,
					255,
					0,
					128,
				}, // Semi-transparent
			}
			img := image.NewPaletted(
				image.Rect(0, 0, 10, 10),
				palette,
			)
			if !imageHasAlpha(img) {
				t.Error(
					"Paletted image with semi-transparent color should have alpha",
				)
			}
		},
	)

	t.Run(
		"Paletted without alpha",
		func(t *testing.T) {
			palette := []color.Color{
				color.RGBA{255, 0, 0, 255},
				color.RGBA{0, 255, 0, 255},
			}
			img := image.NewPaletted(
				image.Rect(0, 0, 10, 10),
				palette,
			)
			if imageHasAlpha(img) {
				t.Error(
					"Paletted image with all opaque colors should not have alpha",
				)
			}
		},
	)
}
