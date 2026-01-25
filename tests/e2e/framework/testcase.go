package framework

import (
	"encoding/json"
	"time"
)

// TestCase represents a single visual comparison test between goffice and Open-XML-SDK.
// It defines the presentation content to generate and configuration for comparison.
type TestCase struct {
	// Metadata
	ID          string       `json:"id"`          // Unique identifier (e.g., "chart_bar_clustered_01")
	Name        string       `json:"name"`        // Human-readable name
	Description string       `json:"description"` // What this test verifies
	Category    TestCategory `json:"category"`    // Chart, Shape, Text, Integration
	Tags        []string     `json:"tags"`        // For filtering (e.g., "chart", "bar", "basic")
	Created     time.Time    `json:"created"`
	Author      string       `json:"author"`

	// Test specification
	Spec     TestSpec          `json:"spec"`     // Actual test data
	Config   TestConfig        `json:"config"`   // Test-specific configuration
	Metadata map[string]string `json:"metadata"` // Extensible metadata
}

// TestCategory categorizes test cases
type TestCategory string

const (
	CategoryChart       TestCategory = "chart"
	CategoryShape       TestCategory = "shape"
	CategoryText        TestCategory = "text"
	CategoryImage       TestCategory = "image"
	CategoryTable       TestCategory = "table"
	CategoryIntegration TestCategory = "integration"
)

// TestSpec defines the presentation content to generate.
// Both Go and C# generators use this specification to create identical presentations.
type TestSpec struct {
	// Slide configuration
	SlideCount int       `json:"slide_count"`
	SlideSize  SlideSize `json:"slide_size"`

	// Content per slide
	Slides []SlideSpec `json:"slides"`

	// Global settings
	Theme *ThemeSpec `json:"theme,omitempty"`
	Fonts []FontSpec `json:"fonts,omitempty"`
}

// SlideSize defines presentation dimensions (EMUs)
type SlideSize struct {
	Width  int64 `json:"width"`  // EMUs (914400 = 1 inch)
	Height int64 `json:"height"` // EMUs
}

var (
	SlideSize4x3 = SlideSize{
		Width:  9144000,
		Height: 6858000,
	} // 10" × 7.5"
	SlideSize16x9 = SlideSize{
		Width:  9144000,
		Height: 5143500,
	} // 10" × 5.625"
	SlideSizeLetter = SlideSize{
		Width:  7772400,
		Height: 10058400,
	} // 8.5" × 11"
)

// SlideSpec defines content for a single slide
type SlideSpec struct {
	Index      int             `json:"index"`  // 0-based slide index
	Layout     string          `json:"layout"` // Layout name (e.g., "Blank", "Title Only")
	Background *BackgroundSpec `json:"background,omitempty"`
	Elements   []ElementSpec   `json:"elements"` // Shapes, charts, text, etc.
}

// ElementSpec defines a single element on a slide (shape, chart, text, etc.)
type ElementSpec struct {
	Type     ElementType `json:"type"`     // Shape, Chart, Text, Image, Table
	Position Position    `json:"position"` // X, Y in EMUs
	Size     Size        `json:"size"`     // Width, Height in EMUs
	ZIndex   int         `json:"z_index"`  // Layering order

	// Type-specific data (only one should be set)
	Shape *ShapeSpec `json:"shape,omitempty"`
	Chart *ChartSpec `json:"chart,omitempty"`
	Text  *TextSpec  `json:"text,omitempty"`
	Image *ImageSpec `json:"image,omitempty"`
	Table *TableSpec `json:"table,omitempty"`
}

// Position in EMUs
type Position struct {
	X int64 `json:"x"` // EMUs from left edge
	Y int64 `json:"y"` // EMUs from top edge
}

// Size in EMUs
type Size struct {
	Width  int64 `json:"width"`  // EMUs
	Height int64 `json:"height"` // EMUs
}

// ElementType identifies element kind
type ElementType string

const (
	ElementTypeShape ElementType = "shape"
	ElementTypeChart ElementType = "chart"
	ElementTypeText  ElementType = "text"
	ElementTypeImage ElementType = "image"
	ElementTypeTable ElementType = "table"
)

//
// Chart Specifications
//

// ChartSpec defines a chart element
type ChartSpec struct {
	Type   ChartType   `json:"type"` // Bar, Line, Pie, etc.
	Title  string      `json:"title"`
	Legend *LegendSpec `json:"legend,omitempty"`
	Data   ChartData   `json:"data"`
	Style  ChartStyle  `json:"style"`
	Axes   *AxesSpec   `json:"axes,omitempty"` // Nil for pie/doughnut
}

// ChartType enum
type ChartType string

const (
	ChartTypeBarClustered      ChartType = "bar_clustered"
	ChartTypeBarStacked        ChartType = "bar_stacked"
	ChartTypeBarPercentStacked ChartType = "bar_percent_stacked"
	ChartTypeColumn            ChartType = "column"
	ChartTypeLine              ChartType = "line"
	ChartTypeLineSmooth        ChartType = "line_smooth"
	ChartTypeLineMarkers       ChartType = "line_markers"
	ChartTypePie               ChartType = "pie"
	ChartTypeDoughnut          ChartType = "doughnut"
	ChartTypeArea              ChartType = "area"
	ChartTypeAreaStacked       ChartType = "area_stacked"
	ChartTypeScatter           ChartType = "scatter"
	ChartTypeScatterSmooth     ChartType = "scatter_smooth"
	ChartTypeScatterLine       ChartType = "scatter_line"
	ChartTypeBubble            ChartType = "bubble"
	ChartTypeRadar             ChartType = "radar"
	ChartTypeRadarFilled       ChartType = "radar_filled"
	ChartTypeStock             ChartType = "stock"
	ChartTypeSurface           ChartType = "surface"
)

// ChartData holds series and categories
type ChartData struct {
	Categories []string     `json:"categories"` // X-axis labels or pie slice names
	Series     []SeriesData `json:"series"`     // One or more data series
}

// SeriesData defines a single data series
type SeriesData struct {
	Name        string    `json:"name"`                   // Series name
	Values      []float64 `json:"values"`                 // Y-values (or sizes for bubble)
	XValues     []float64 `json:"x_values,omitempty"`     // For scatter/bubble
	Color       string    `json:"color"`                  // Hex color (e.g., "4472C4")
	MarkerStyle string    `json:"marker_style,omitempty"` // For line/scatter
	Smooth      bool      `json:"smooth,omitempty"`       // For line/area
}

// ChartStyle defines visual styling
type ChartStyle struct {
	ColorPalette   []string     `json:"color_palette"` // Hex colors for auto-coloring
	FontSize       int          `json:"font_size"`     // Points (e.g., 12)
	FontFamily     string       `json:"font_family"`   // "Arial", "Calibri", etc.
	ShowDataLabels bool         `json:"show_data_labels"`
	GridLines      GridLineSpec `json:"grid_lines"`
}

// GridLineSpec configures gridlines
type GridLineSpec struct {
	MajorHorizontal bool `json:"major_horizontal"`
	MinorHorizontal bool `json:"minor_horizontal"`
	MajorVertical   bool `json:"major_vertical"`
	MinorVertical   bool `json:"minor_vertical"`
}

// LegendSpec configures legend
type LegendSpec struct {
	Position   LegendPosition `json:"position"` // Top, Bottom, Left, Right, TopRight
	ShowLegend bool           `json:"show_legend"`
}

// LegendPosition enum
type LegendPosition string

const (
	LegendPositionTop      LegendPosition = "top"
	LegendPositionBottom   LegendPosition = "bottom"
	LegendPositionLeft     LegendPosition = "left"
	LegendPositionRight    LegendPosition = "right"
	LegendPositionTopRight LegendPosition = "top_right"
)

// AxesSpec configures axes
type AxesSpec struct {
	Category AxisSpec `json:"category"`
	Value    AxisSpec `json:"value"`
}

// AxisSpec configures a single axis
type AxisSpec struct {
	Title             string  `json:"title"`
	Min               float64 `json:"min,omitempty"`                 // Auto if 0
	Max               float64 `json:"max,omitempty"`                 // Auto if 0
	MajorUnit         float64 `json:"major_unit,omitempty"`          // Auto if 0
	MinorUnit         float64 `json:"minor_unit,omitempty"`          // Auto if 0
	NumberFormat      string  `json:"number_format,omitempty"`       // "0.00", "#,##0", etc.
	TickLabelRotation int     `json:"tick_label_rotation,omitempty"` // Degrees (-90 to 90)
}

//
// Shape Specifications
//

// ShapeSpec defines a shape element
type ShapeSpec struct {
	Type    ShapeType    `json:"type"` // Rectangle, Ellipse, etc.
	Fill    FillSpec     `json:"fill"`
	Stroke  StrokeSpec   `json:"stroke"`
	Effects []EffectSpec `json:"effects,omitempty"`
	Text    *TextSpec    `json:"text,omitempty"` // Optional text inside shape
}

// ShapeType enum (matches DrawingML presets)
type ShapeType string

const (
	ShapeTypeRectangle  ShapeType = "rect"
	ShapeTypeRoundRect  ShapeType = "roundRect"
	ShapeTypeEllipse    ShapeType = "ellipse"
	ShapeTypeTriangle   ShapeType = "triangle"
	ShapeTypeDiamond    ShapeType = "diamond"
	ShapeTypePentagon   ShapeType = "pentagon"
	ShapeTypeHexagon    ShapeType = "hexagon"
	ShapeTypeOctagon    ShapeType = "octagon"
	ShapeTypeStar5      ShapeType = "star5"
	ShapeTypeArrowRight ShapeType = "rightArrow"
	ShapeTypeCallout    ShapeType = "wedgeRectCallout"
)

// FillSpec defines fill styling
type FillSpec struct {
	Type     FillType      `json:"type"`               // Solid, Gradient, Pattern, None
	Color    string        `json:"color,omitempty"`    // Hex color (for solid)
	Gradient *GradientSpec `json:"gradient,omitempty"` // Gradient definition
	Pattern  *PatternSpec  `json:"pattern,omitempty"`  // Pattern definition
}

// FillType enum
type FillType string

const (
	FillTypeSolid    FillType = "solid"
	FillTypeGradient FillType = "gradient"
	FillTypePattern  FillType = "pattern"
	FillTypeNone     FillType = "none"
)

// GradientSpec defines gradient fill
type GradientSpec struct {
	Type  GradientType   `json:"type"`            // Linear, Radial
	Angle int            `json:"angle,omitempty"` // Degrees (for linear)
	Stops []GradientStop `json:"stops"`
}

// GradientType enum
type GradientType string

const (
	GradientTypeLinear GradientType = "linear"
	GradientTypeRadial GradientType = "radial"
)

// GradientStop defines a color stop
type GradientStop struct {
	Position float64 `json:"position"` // 0.0 to 1.0
	Color    string  `json:"color"`    // Hex color
}

// PatternSpec defines pattern fill
type PatternSpec struct {
	Type       PatternType `json:"type"`       // Dots, Grid, DiagonalStripe, etc.
	Foreground string      `json:"foreground"` // Hex color
	Background string      `json:"background"` // Hex color
}

// PatternType enum (subset of DrawingML patterns)
type PatternType string

const (
	PatternTypeDots           PatternType = "dots"
	PatternTypeGrid           PatternType = "grid"
	PatternTypeDiagonalStripe PatternType = "diagonalStripe"
	PatternTypeCheckered      PatternType = "checkered"
)

// StrokeSpec defines outline/stroke styling
type StrokeSpec struct {
	Width     int64     `json:"width"`      // EMUs
	Color     string    `json:"color"`      // Hex color
	DashStyle DashStyle `json:"dash_style"` // Solid, Dot, Dash, etc.
	Cap       CapStyle  `json:"cap"`        // Flat, Round, Square
	Join      JoinStyle `json:"join"`       // Miter, Round, Bevel
}

// DashStyle enum
type DashStyle string

const (
	DashStyleSolid    DashStyle = "solid"
	DashStyleDot      DashStyle = "dot"
	DashStyleDash     DashStyle = "dash"
	DashStyleDashDot  DashStyle = "dashDot"
	DashStyleLongDash DashStyle = "longDash"
)

// CapStyle enum
type CapStyle string

const (
	CapStyleFlat   CapStyle = "flat"
	CapStyleRound  CapStyle = "round"
	CapStyleSquare CapStyle = "square"
)

// JoinStyle enum
type JoinStyle string

const (
	JoinStyleMiter JoinStyle = "miter"
	JoinStyleRound JoinStyle = "round"
	JoinStyleBevel JoinStyle = "bevel"
)

// EffectSpec defines visual effects
type EffectSpec struct {
	Type       EffectType      `json:"type"` // Shadow, Glow, Reflection, etc.
	Shadow     *ShadowSpec     `json:"shadow,omitempty"`
	Glow       *GlowSpec       `json:"glow,omitempty"`
	Reflection *ReflectionSpec `json:"reflection,omitempty"`
	SoftEdge   *SoftEdgeSpec   `json:"soft_edge,omitempty"`
}

// EffectType enum
type EffectType string

const (
	EffectTypeShadow     EffectType = "shadow"
	EffectTypeGlow       EffectType = "glow"
	EffectTypeReflection EffectType = "reflection"
	EffectTypeSoftEdge   EffectType = "soft_edge"
)

// ShadowSpec defines shadow effect
type ShadowSpec struct {
	Type         ShadowType `json:"type"`         // Outer, Inner, Perspective
	Angle        int        `json:"angle"`        // Degrees
	Distance     int64      `json:"distance"`     // EMUs
	BlurRadius   int64      `json:"blur_radius"`  // EMUs
	Color        string     `json:"color"`        // Hex with optional alpha (e.g., "80000000" for semi-transparent black)
	Transparency float64    `json:"transparency"` // 0.0 (opaque) to 1.0 (transparent)
}

// ShadowType enum
type ShadowType string

const (
	ShadowTypeOuter       ShadowType = "outer"
	ShadowTypeInner       ShadowType = "inner"
	ShadowTypePerspective ShadowType = "perspective"
)

// GlowSpec defines glow effect
type GlowSpec struct {
	Radius int64  `json:"radius"` // EMUs
	Color  string `json:"color"`  // Hex with optional alpha
}

// ReflectionSpec defines reflection effect
type ReflectionSpec struct {
	BlurRadius    int64   `json:"blur_radius"`    // EMUs
	StartOpacity  float64 `json:"start_opacity"`  // 0.0 to 1.0
	EndOpacity    float64 `json:"end_opacity"`    // 0.0 to 1.0
	Distance      int64   `json:"distance"`       // EMUs
	Direction     int     `json:"direction"`      // Degrees
	FadeDirection int     `json:"fade_direction"` // Degrees
	StartPosition float64 `json:"start_position"` // 0.0 to 1.0
	EndPosition   float64 `json:"end_position"`   // 0.0 to 1.0
}

// SoftEdgeSpec defines soft edge effect
type SoftEdgeSpec struct {
	Radius int64 `json:"radius"` // EMUs
}

//
// Text Specifications
//

// TextSpec defines text content and formatting
type TextSpec struct {
	Content    string          `json:"content"`              // Raw text or placeholder
	Paragraphs []ParagraphSpec `json:"paragraphs,omitempty"` // For rich text
	DefaultRun RunSpec         `json:"default_run"`          // Default formatting
}

// ParagraphSpec defines a paragraph
type ParagraphSpec struct {
	Runs        []RunSpec     `json:"runs"`                   // Text runs with formatting
	Alignment   TextAlignment `json:"alignment"`              // Left, Center, Right, Justify
	LineSpacing float64       `json:"line_spacing,omitempty"` // 1.0 = single, 1.5 = 1.5x, etc.
	SpaceBefore int64         `json:"space_before,omitempty"` // EMUs
	SpaceAfter  int64         `json:"space_after,omitempty"`  // EMUs
	Indent      int64         `json:"indent,omitempty"`       // EMUs
	BulletChar  string        `json:"bullet_char,omitempty"`  // "•", "-", etc. (empty = no bullet)
}

// RunSpec defines a text run with formatting
type RunSpec struct {
	Text          string        `json:"text"`
	FontFamily    string        `json:"font_family"` // "Arial", "Calibri", etc.
	FontSize      int           `json:"font_size"`   // Points (e.g., 18)
	Color         string        `json:"color"`       // Hex color
	Bold          bool          `json:"bold"`
	Italic        bool          `json:"italic"`
	Underline     UnderlineType `json:"underline"` // None, Single, Double, etc.
	Strikethrough bool          `json:"strikethrough"`
}

// TextAlignment enum
type TextAlignment string

const (
	TextAlignLeft    TextAlignment = "left"
	TextAlignCenter  TextAlignment = "center"
	TextAlignRight   TextAlignment = "right"
	TextAlignJustify TextAlignment = "justify"
)

// UnderlineType enum
type UnderlineType string

const (
	UnderlineNone   UnderlineType = "none"
	UnderlineSingle UnderlineType = "single"
	UnderlineDouble UnderlineType = "double"
	UnderlineDotted UnderlineType = "dotted"
	UnderlineDash   UnderlineType = "dash"
)

//
// Image & Other Element Specs
//

// ImageSpec defines an image element
type ImageSpec struct {
	Path    string       `json:"path"` // Relative to fixtures/images/
	AltText string       `json:"alt_text"`
	Crop    *CropSpec    `json:"crop,omitempty"`
	Effects []EffectSpec `json:"effects,omitempty"`
}

// CropSpec defines image cropping
type CropSpec struct {
	Left   float64 `json:"left"` // 0.0 to 1.0 (percentage)
	Top    float64 `json:"top"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}

// TableSpec defines a table element
type TableSpec struct {
	Rows    int               `json:"rows"`
	Columns int               `json:"columns"`
	Cells   [][]TableCellSpec `json:"cells"`
	Style   TableStyleSpec    `json:"style"`
}

// TableCellSpec defines a single table cell
type TableCellSpec struct {
	Text       string   `json:"text"`
	Span       CellSpan `json:"span,omitempty"`
	Fill       FillSpec `json:"fill"`
	TextFormat RunSpec  `json:"text_format"`
}

// CellSpan defines cell merging
type CellSpan struct {
	RowSpan int `json:"row_span"` // 1 = no merge
	ColSpan int `json:"col_span"` // 1 = no merge
}

// TableStyleSpec defines table styling
type TableStyleSpec struct {
	HeaderRow     bool `json:"header_row"`
	TotalRow      bool `json:"total_row"`
	BandedRows    bool `json:"banded_rows"`
	FirstColumn   bool `json:"first_column"`
	LastColumn    bool `json:"last_column"`
	BandedColumns bool `json:"banded_columns"`
}

//
// Background & Theme Specs
//

// BackgroundSpec defines slide background
type BackgroundSpec struct {
	Fill FillSpec `json:"fill"`
}

// ThemeSpec defines presentation theme
type ThemeSpec struct {
	Name        string      `json:"name"`
	ColorScheme ColorScheme `json:"color_scheme"`
	FontScheme  FontScheme  `json:"font_scheme"`
}

// ColorScheme defines theme colors
type ColorScheme struct {
	Accent1           string `json:"accent1"` // Hex
	Accent2           string `json:"accent2"`
	Accent3           string `json:"accent3"`
	Accent4           string `json:"accent4"`
	Accent5           string `json:"accent5"`
	Accent6           string `json:"accent6"`
	Dark1             string `json:"dark1"` // Text/background
	Dark2             string `json:"dark2"`
	Light1            string `json:"light1"` // Background/text
	Light2            string `json:"light2"`
	Hyperlink         string `json:"hyperlink"`
	FollowedHyperlink string `json:"followed_hyperlink"`
}

// FontScheme defines theme fonts
type FontScheme struct {
	MajorFont string `json:"major_font"` // Headings
	MinorFont string `json:"minor_font"` // Body text
}

// FontSpec defines a font to embed
type FontSpec struct {
	Family string `json:"family"` // "Arial", "Calibri", etc.
	Path   string `json:"path"`   // Relative to fixtures/fonts/
}

//
// Test Configuration
//

// TestConfig configures test execution behavior including rendering, comparison, and retry logic.
// Can be set globally or per-test case for fine-grained control.
type TestConfig struct {
	// Rendering
	RenderDPI     int           `json:"render_dpi"`     // DPI for PNG output (default: 300)
	RenderBackend RenderBackend `json:"render_backend"` // LibreOffice, PowerPoint, Custom

	// Comparison
	DiffThreshold      float64       `json:"diff_threshold"`      // 0.0 to 1.0 (0 = identical, 1 = completely different)
	DiffAlgorithm      DiffAlgorithm `json:"diff_algorithm"`      // PixelPerfect, Perceptual, SSIM
	IgnoreAntialiasing bool          `json:"ignore_antialiasing"` // Tolerate minor AA differences

	// Timeout
	GeneratorTimeout int `json:"generator_timeout"` // Seconds
	RenderTimeout    int `json:"render_timeout"`    // Seconds

	// Reporting
	GenerateDiffImage bool `json:"generate_diff_image"` // Create red overlay PNG
	SaveIntermediates bool `json:"save_intermediates"`  // Keep PDF, PPTX files

	// Retry
	RetryCount int `json:"retry_count"` // Number of retries on failure
	RetryDelay int `json:"retry_delay"` // Seconds between retries
}

// RenderBackend enum
type RenderBackend string

const (
	RenderBackendLibreOffice RenderBackend = "libreoffice"
	RenderBackendPowerPoint  RenderBackend = "powerpoint"
	RenderBackendPDF         RenderBackend = "pdf" // Use goffice PDF rendering
	RenderBackendCustom      RenderBackend = "custom"
)

// DiffAlgorithm enum
type DiffAlgorithm string

const (
	DiffAlgorithmPixelPerfect DiffAlgorithm = "pixel_perfect" // Exact pixel match
	DiffAlgorithmPerceptual   DiffAlgorithm = "perceptual"    // Human-visible differences
	DiffAlgorithmSSIM         DiffAlgorithm = "ssim"          // Structural similarity index
	DiffAlgorithmMSE          DiffAlgorithm = "mse"           // Mean squared error
	DiffAlgorithmPSNR         DiffAlgorithm = "psnr"          // Peak signal-to-noise ratio
)

//
// Helper Methods
//

// NewTestCase creates a test case with sensible defaults.
// Automatically sets creation time and initializes configuration.
func NewTestCase(
	id, name string,
	category TestCategory,
) *TestCase {
	return &TestCase{
		ID:       id,
		Name:     name,
		Category: category,
		Created:  time.Now(),
		Config:   DefaultTestConfig(),
		Metadata: make(map[string]string),
		Tags:     []string{},
	}
}

// DefaultTestConfig returns sensible defaults for test configuration.
// Uses LibreOffice rendering at 300 DPI with 1% perceptual difference threshold.
func DefaultTestConfig() TestConfig {
	return TestConfig{
		RenderDPI:          300,
		RenderBackend:      RenderBackendLibreOffice,
		DiffThreshold:      0.01, // 1% difference threshold
		DiffAlgorithm:      DiffAlgorithmPerceptual,
		IgnoreAntialiasing: true,
		GeneratorTimeout:   60,
		RenderTimeout:      120,
		GenerateDiffImage:  true,
		SaveIntermediates:  false,
		RetryCount:         2,
		RetryDelay:         5,
	}
}

// ToJSON serializes test case to JSON
func (tc *TestCase) ToJSON() ([]byte, error) {
	return json.MarshalIndent(tc, "", "  ")
}

// FromJSON deserializes test case from JSON
func FromJSON(data []byte) (*TestCase, error) {
	var tc TestCase
	err := json.Unmarshal(data, &tc)

	return &tc, err
}
