# E2E Visual Testing Framework: Detailed Design

## Table of Contents
1. [Context & Background](#context--background)
2. [Architecture Overview](#architecture-overview)
3. [Component Designs](#component-designs)
4. [Code Patterns & Implementation](#code-patterns--implementation)
5. [Nix Integration](#nix-integration)
6. [Visual Comparison Algorithms](#visual-comparison-algorithms)
7. [CI/CD Integration](#cicd-integration)
8. [Migration & Rollout](#migration--rollout)
9. [Trade-offs & Alternatives](#trade-offs--alternatives)
10. [Open Questions](#open-questions)

---

## Context & Background

### Problem Statement

The goffice library implements OOXML (Office Open XML) presentation generation in pure Go, aiming for **feature parity** with Microsoft's authoritative Open-XML-SDK for .NET. However, verifying that our generated `.pptx` files render **visually identically** to the C# SDK's output is currently manual and error-prone.

**Specific challenges:**
1. **Chart Rendering**: 11 chart types × multiple configurations = hundreds of visual combinations
2. **DrawingML Complexity**: 140+ colors, 100+ shapes, effects (shadows, glows), gradients
3. **Text Formatting**: Font families, sizes, colors, bold/italic, alignment, line spacing
4. **Layout Precision**: EMU positioning (914400 EMUs = 1 inch) must be pixel-perfect
5. **Office Version Compatibility**: Office 2007, 2010, 2013, 2016, 2019, 365 may render differently

### Goals

**Primary Goals:**
- **G1**: Automate visual comparison between Go and C# generated presentations
- **G2**: Detect visual regressions in PR reviews before merge
- **G3**: Provide developers with clear, actionable diff reports
- **G4**: Support all presentation features: charts, shapes, images, text, tables
- **G5**: Achieve <5% false positive rate on visual diffs

**Secondary Goals:**
- **G6**: Enable rapid iteration on chart styling (color palettes, marker styles)
- **G7**: Provide baseline for PDF rendering quality assessment
- **G8**: Create reusable framework for Word/Excel visual testing

### Non-Goals

- **NG1**: Test PowerPoint-specific features (animations, slide transitions) unless they affect static rendering
- **NG2**: Replace existing unit/integration tests (this is complementary)
- **NG3**: Guarantee pixel-perfect matching (perceptual similarity is sufficient)
- **NG4**: Support all Office versions simultaneously (focus on 2016+ initially)

---

## Architecture Overview

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Test Case Definitions                     │
│                    (tests/e2e/testcases/)                    │
│  - chart_bar.go, chart_line.go, shapes.go, text.go, etc.   │
└────────────┬─────────────────────────────────────┬──────────┘
             │                                     │
             ▼                                     ▼
┌──────────────────────────┐         ┌──────────────────────────┐
│   Go Generator Harness    │         │   C# Generator Harness    │
│  (tests/e2e/generators/   │         │ (tests/e2e/generators/    │
│         go/)              │         │        csharp/)           │
│  - Uses goffice           │         │ - Uses Open-XML-SDK       │
│  - Reads test definitions │         │ - Reads same definitions  │
│  - Outputs .pptx          │         │ - Outputs .pptx           │
└────────────┬──────────────┘         └────────────┬─────────────┘
             │                                     │
             ▼                                     ▼
┌──────────────────────────┐         ┌──────────────────────────┐
│   go_output.pptx         │         │   csharp_output.pptx      │
└────────────┬──────────────┘         └────────────┬─────────────┘
             │                                     │
             └───────────────┬─────────────────────┘
                             ▼
                 ┌───────────────────────┐
                 │  Rendering Pipeline   │
                 │  (LibreOffice/PPTX)   │
                 │  - Convert to PDF     │
                 │  - Extract as PNG     │
                 └─────────┬─────────────┘
                           │
         ┌─────────────────┴─────────────────┐
         ▼                                   ▼
┌──────────────────┐               ┌──────────────────┐
│ go_output.png    │               │ csharp_output.png│
└────────┬─────────┘               └────────┬─────────┘
         │                                   │
         └────────────┬──────────────────────┘
                      ▼
          ┌───────────────────────┐
          │  Visual Diff Engine   │
          │  - Pixel comparison   │
          │  - Perceptual metrics │
          │  - Diff highlighting  │
          └─────────┬─────────────┘
                    ▼
          ┌───────────────────────┐
          │   HTML Report         │
          │  - Side-by-side view  │
          │  - Diff overlay       │
          │  - Metrics dashboard  │
          └───────────────────────┘
```

### Directory Structure

```
tests/e2e/
├── framework/                  # Core framework code
│   ├── testcase.go            # Test case definition types
│   ├── generator.go           # Generator interface and harness
│   ├── renderer.go            # PPTX → PNG rendering
│   ├── differ.go              # Image comparison engine
│   ├── reporter.go            # HTML report generation
│   └── config.go              # Configuration management
│
├── generators/                # Platform-specific generators
│   ├── go/                    # Go generator (uses goffice)
│   │   ├── main.go           # Entry point
│   │   ├── chart_gen.go      # Chart generation
│   │   ├── shape_gen.go      # Shape generation
│   │   ├── text_gen.go       # Text generation
│   │   └── helpers.go        # Shared utilities
│   │
│   └── csharp/               # C# generator (uses Open-XML-SDK)
│       ├── Program.cs        # Entry point
│       ├── ChartGenerator.cs # Chart generation
│       ├── ShapeGenerator.cs # Shape generation
│       ├── TextGenerator.cs  # Text generation
│       ├── Helpers.cs        # Shared utilities
│       └── TestRunner.csproj # .NET project file
│
├── testcases/                # Test case definitions
│   ├── charts/
│   │   ├── bar_basic.go
│   │   ├── bar_stacked.go
│   │   ├── line_smooth.go
│   │   ├── pie_exploded.go
│   │   └── scatter_xy.go
│   │
│   ├── shapes/
│   │   ├── rectangles.go
│   │   ├── circles.go
│   │   └── complex_paths.go
│   │
│   ├── text/
│   │   ├── formatting.go
│   │   ├── alignment.go
│   │   └── multilevel.go
│   │
│   └── integration/          # Multi-feature tests
│       └── dashboard.go
│
├── fixtures/                 # Test data
│   ├── data/
│   │   ├── sales_data.json  # Sample chart data
│   │   └── categories.json
│   │
│   ├── images/              # Test images
│   │   └── logo.png
│   │
│   └── fonts/               # Embedded fonts for consistency
│       └── Arial.ttf
│
├── golden/                   # Reference images (optional)
│   └── baselines/
│       └── chart_bar_basic_csharp.png
│
├── output/                   # Generated test artifacts
│   ├── pptx/
│   │   ├── go/
│   │   └── csharp/
│   ├── pdf/
│   ├── png/
│   └── diffs/
│
├── reports/                  # Generated HTML reports
│   ├── index.html
│   ├── assets/
│   └── diffs/
│
├── scripts/                  # Helper scripts
│   ├── run_tests.sh
│   ├── update_baselines.sh
│   └── generate_report.sh
│
├── go.mod                    # Go module for framework
├── e2e_test.go              # Main test entry point
└── README.md

```

---

## Component Designs

### 1. Test Case Definition System

**Purpose:** Declarative, platform-agnostic test case specifications

#### Core Types

```go
// framework/testcase.go

package framework

import (
	"encoding/json"
	"time"
)

// TestCase represents a single visual comparison test
type TestCase struct {
	// Metadata
	ID          string            `json:"id"`          // Unique identifier (e.g., "chart_bar_clustered_01")
	Name        string            `json:"name"`        // Human-readable name
	Description string            `json:"description"` // What this test verifies
	Category    TestCategory      `json:"category"`    // Chart, Shape, Text, Integration
	Tags        []string          `json:"tags"`        // For filtering (e.g., "chart", "bar", "basic")
	Created     time.Time         `json:"created"`
	Author      string            `json:"author"`

	// Test specification
	Spec        TestSpec          `json:"spec"`        // Actual test data
	Config      TestConfig        `json:"config"`      // Test-specific configuration
	Metadata    map[string]string `json:"metadata"`    // Extensible metadata
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

// TestSpec defines the presentation content to generate
type TestSpec struct {
	// Slide configuration
	SlideCount    int               `json:"slide_count"`
	SlideSize     SlideSize         `json:"slide_size"`

	// Content per slide
	Slides        []SlideSpec       `json:"slides"`

	// Global settings
	Theme         *ThemeSpec        `json:"theme,omitempty"`
	Fonts         []FontSpec        `json:"fonts,omitempty"`
}

// SlideSize defines presentation dimensions (EMUs)
type SlideSize struct {
	Width  int64 `json:"width"`  // EMUs (914400 = 1 inch)
	Height int64 `json:"height"` // EMUs
}

var (
	SlideSize4x3 = SlideSize{Width: 9144000, Height: 6858000}   // 10" × 7.5"
	SlideSize16x9 = SlideSize{Width: 9144000, Height: 5143500}  // 10" × 5.625"
	SlideSizeLetter = SlideSize{Width: 7772400, Height: 10058400} // 8.5" × 11"
)

// SlideSpec defines content for a single slide
type SlideSpec struct {
	Index       int               `json:"index"`      // 0-based slide index
	Layout      string            `json:"layout"`     // Layout name (e.g., "Blank", "Title Only")
	Background  *BackgroundSpec   `json:"background,omitempty"`
	Elements    []ElementSpec     `json:"elements"`   // Shapes, charts, text, etc.
}

// ElementSpec defines a single element on a slide (shape, chart, text, etc.)
type ElementSpec struct {
	Type        ElementType       `json:"type"`        // Shape, Chart, Text, Image, Table
	Position    Position          `json:"position"`    // X, Y in EMUs
	Size        Size              `json:"size"`        // Width, Height in EMUs
	ZIndex      int               `json:"z_index"`     // Layering order

	// Type-specific data (only one should be set)
	Shape       *ShapeSpec        `json:"shape,omitempty"`
	Chart       *ChartSpec        `json:"chart,omitempty"`
	Text        *TextSpec         `json:"text,omitempty"`
	Image       *ImageSpec        `json:"image,omitempty"`
	Table       *TableSpec        `json:"table,omitempty"`
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
	Type        ChartType         `json:"type"`        // Bar, Line, Pie, etc.
	Title       string            `json:"title"`
	Legend      *LegendSpec       `json:"legend,omitempty"`
	Data        ChartData         `json:"data"`
	Style       ChartStyle        `json:"style"`
	Axes        *AxesSpec         `json:"axes,omitempty"` // Nil for pie/doughnut
}

// ChartType enum
type ChartType string

const (
	ChartTypeBarClustered       ChartType = "bar_clustered"
	ChartTypeBarStacked         ChartType = "bar_stacked"
	ChartTypeBarPercentStacked  ChartType = "bar_percent_stacked"
	ChartTypeColumn             ChartType = "column"
	ChartTypeLine               ChartType = "line"
	ChartTypeLineSmooth         ChartType = "line_smooth"
	ChartTypeLineMarkers        ChartType = "line_markers"
	ChartTypePie                ChartType = "pie"
	ChartTypeDoughnut           ChartType = "doughnut"
	ChartTypeArea               ChartType = "area"
	ChartTypeAreaStacked        ChartType = "area_stacked"
	ChartTypeScatter            ChartType = "scatter"
	ChartTypeScatterSmooth      ChartType = "scatter_smooth"
	ChartTypeScatterLine        ChartType = "scatter_line"
	ChartTypeBubble             ChartType = "bubble"
	ChartTypeRadar              ChartType = "radar"
	ChartTypeRadarFilled        ChartType = "radar_filled"
	ChartTypeStock              ChartType = "stock"
	ChartTypeSurface            ChartType = "surface"
)

// ChartData holds series and categories
type ChartData struct {
	Categories  []string           `json:"categories"` // X-axis labels or pie slice names
	Series      []SeriesData       `json:"series"`     // One or more data series
}

// SeriesData defines a single data series
type SeriesData struct {
	Name        string             `json:"name"`       // Series name
	Values      []float64          `json:"values"`     // Y-values (or sizes for bubble)
	XValues     []float64          `json:"x_values,omitempty"` // For scatter/bubble
	Color       string             `json:"color"`      // Hex color (e.g., "4472C4")
	MarkerStyle string             `json:"marker_style,omitempty"` // For line/scatter
	Smooth      bool               `json:"smooth,omitempty"`       // For line/area
}

// ChartStyle defines visual styling
type ChartStyle struct {
	ColorPalette  []string         `json:"color_palette"` // Hex colors for auto-coloring
	FontSize      int              `json:"font_size"`     // Points (e.g., 12)
	FontFamily    string           `json:"font_family"`   // "Arial", "Calibri", etc.
	ShowDataLabels bool            `json:"show_data_labels"`
	GridLines     GridLineSpec     `json:"grid_lines"`
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
	Position    LegendPosition `json:"position"` // Top, Bottom, Left, Right, TopRight
	ShowLegend  bool           `json:"show_legend"`
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
	Category  AxisSpec `json:"category"`
	Value     AxisSpec `json:"value"`
}

// AxisSpec configures a single axis
type AxisSpec struct {
	Title           string  `json:"title"`
	Min             float64 `json:"min,omitempty"`             // Auto if 0
	Max             float64 `json:"max,omitempty"`             // Auto if 0
	MajorUnit       float64 `json:"major_unit,omitempty"`      // Auto if 0
	MinorUnit       float64 `json:"minor_unit,omitempty"`      // Auto if 0
	NumberFormat    string  `json:"number_format,omitempty"`   // "0.00", "#,##0", etc.
	TickLabelRotation int   `json:"tick_label_rotation,omitempty"` // Degrees (-90 to 90)
}

//
// Shape Specifications
//

// ShapeSpec defines a shape element
type ShapeSpec struct {
	Type        ShapeType       `json:"type"`        // Rectangle, Ellipse, etc.
	Fill        FillSpec        `json:"fill"`
	Stroke      StrokeSpec      `json:"stroke"`
	Effects     []EffectSpec    `json:"effects,omitempty"`
	Text        *TextSpec       `json:"text,omitempty"` // Optional text inside shape
}

// ShapeType enum (matches DrawingML presets)
type ShapeType string

const (
	ShapeTypeRectangle      ShapeType = "rect"
	ShapeTypeRoundRect      ShapeType = "roundRect"
	ShapeTypeEllipse        ShapeType = "ellipse"
	ShapeTypeTriangle       ShapeType = "triangle"
	ShapeTypeDiamond        ShapeType = "diamond"
	ShapeTypePentagon       ShapeType = "pentagon"
	ShapeTypeHexagon        ShapeType = "hexagon"
	ShapeTypeOctagon        ShapeType = "octagon"
	ShapeTypeStar5          ShapeType = "star5"
	ShapeTypeArrowRight     ShapeType = "rightArrow"
	ShapeTypeCallout        ShapeType = "wedgeRectCallout"
)

// FillSpec defines fill styling
type FillSpec struct {
	Type        FillType        `json:"type"`        // Solid, Gradient, Pattern, None
	Color       string          `json:"color,omitempty"`      // Hex color (for solid)
	Gradient    *GradientSpec   `json:"gradient,omitempty"`   // Gradient definition
	Pattern     *PatternSpec    `json:"pattern,omitempty"`    // Pattern definition
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
	Type        GradientType    `json:"type"`        // Linear, Radial
	Angle       int             `json:"angle,omitempty"`       // Degrees (for linear)
	Stops       []GradientStop  `json:"stops"`
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
	Type       PatternType `json:"type"`        // Dots, Grid, DiagonalStripe, etc.
	Foreground string      `json:"foreground"`  // Hex color
	Background string      `json:"background"`  // Hex color
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
	Width       int64        `json:"width"`       // EMUs
	Color       string       `json:"color"`       // Hex color
	DashStyle   DashStyle    `json:"dash_style"`  // Solid, Dot, Dash, etc.
	Cap         CapStyle     `json:"cap"`         // Flat, Round, Square
	Join        JoinStyle    `json:"join"`        // Miter, Round, Bevel
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
	Type        EffectType      `json:"type"`        // Shadow, Glow, Reflection, etc.
	Shadow      *ShadowSpec     `json:"shadow,omitempty"`
	Glow        *GlowSpec       `json:"glow,omitempty"`
	Reflection  *ReflectionSpec `json:"reflection,omitempty"`
	SoftEdge    *SoftEdgeSpec   `json:"soft_edge,omitempty"`
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
	Type        ShadowType `json:"type"`        // Outer, Inner, Perspective
	Angle       int        `json:"angle"`       // Degrees
	Distance    int64      `json:"distance"`    // EMUs
	BlurRadius  int64      `json:"blur_radius"` // EMUs
	Color       string     `json:"color"`       // Hex with optional alpha (e.g., "80000000" for semi-transparent black)
	Transparency float64   `json:"transparency"` // 0.0 (opaque) to 1.0 (transparent)
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
	Radius  int64  `json:"radius"`  // EMUs
	Color   string `json:"color"`   // Hex with optional alpha
}

// ReflectionSpec defines reflection effect
type ReflectionSpec struct {
	BlurRadius   int64   `json:"blur_radius"`    // EMUs
	StartOpacity float64 `json:"start_opacity"`  // 0.0 to 1.0
	EndOpacity   float64 `json:"end_opacity"`    // 0.0 to 1.0
	Distance     int64   `json:"distance"`       // EMUs
	Direction    int     `json:"direction"`      // Degrees
	FadeDirection int    `json:"fade_direction"` // Degrees
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
	Content     string             `json:"content"`     // Raw text or placeholder
	Paragraphs  []ParagraphSpec    `json:"paragraphs,omitempty"` // For rich text
	DefaultRun  RunSpec            `json:"default_run"` // Default formatting
}

// ParagraphSpec defines a paragraph
type ParagraphSpec struct {
	Runs        []RunSpec          `json:"runs"`        // Text runs with formatting
	Alignment   TextAlignment      `json:"alignment"`   // Left, Center, Right, Justify
	LineSpacing float64            `json:"line_spacing,omitempty"` // 1.0 = single, 1.5 = 1.5x, etc.
	SpaceBefore int64              `json:"space_before,omitempty"` // EMUs
	SpaceAfter  int64              `json:"space_after,omitempty"`  // EMUs
	Indent      int64              `json:"indent,omitempty"`       // EMUs
	BulletChar  string             `json:"bullet_char,omitempty"`  // "•", "-", etc. (empty = no bullet)
}

// RunSpec defines a text run with formatting
type RunSpec struct {
	Text          string       `json:"text"`
	FontFamily    string       `json:"font_family"`    // "Arial", "Calibri", etc.
	FontSize      int          `json:"font_size"`      // Points (e.g., 18)
	Color         string       `json:"color"`          // Hex color
	Bold          bool         `json:"bold"`
	Italic        bool         `json:"italic"`
	Underline     UnderlineType `json:"underline"`     // None, Single, Double, etc.
	Strikethrough bool         `json:"strikethrough"`
}

// TextAlignment enum
type TextAlignment string

const (
	TextAlignLeft     TextAlignment = "left"
	TextAlignCenter   TextAlignment = "center"
	TextAlignRight    TextAlignment = "right"
	TextAlignJustify  TextAlignment = "justify"
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
	Path        string      `json:"path"`        // Relative to fixtures/images/
	AltText     string      `json:"alt_text"`
	Crop        *CropSpec   `json:"crop,omitempty"`
	Effects     []EffectSpec `json:"effects,omitempty"`
}

// CropSpec defines image cropping
type CropSpec struct {
	Left   float64 `json:"left"`   // 0.0 to 1.0 (percentage)
	Top    float64 `json:"top"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}

// TableSpec defines a table element
type TableSpec struct {
	Rows     int               `json:"rows"`
	Columns  int               `json:"columns"`
	Cells    [][]TableCellSpec `json:"cells"`
	Style    TableStyleSpec    `json:"style"`
}

// TableCellSpec defines a single table cell
type TableCellSpec struct {
	Text        string       `json:"text"`
	Span        CellSpan     `json:"span,omitempty"`
	Fill        FillSpec     `json:"fill"`
	TextFormat  RunSpec      `json:"text_format"`
}

// CellSpan defines cell merging
type CellSpan struct {
	RowSpan int `json:"row_span"` // 1 = no merge
	ColSpan int `json:"col_span"` // 1 = no merge
}

// TableStyleSpec defines table styling
type TableStyleSpec struct {
	HeaderRow    bool   `json:"header_row"`
	TotalRow     bool   `json:"total_row"`
	BandedRows   bool   `json:"banded_rows"`
	FirstColumn  bool   `json:"first_column"`
	LastColumn   bool   `json:"last_column"`
	BandedColumns bool  `json:"banded_columns"`
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
	Name         string        `json:"name"`
	ColorScheme  ColorScheme   `json:"color_scheme"`
	FontScheme   FontScheme    `json:"font_scheme"`
}

// ColorScheme defines theme colors
type ColorScheme struct {
	Accent1 string `json:"accent1"` // Hex
	Accent2 string `json:"accent2"`
	Accent3 string `json:"accent3"`
	Accent4 string `json:"accent4"`
	Accent5 string `json:"accent5"`
	Accent6 string `json:"accent6"`
	Dark1   string `json:"dark1"`   // Text/background
	Dark2   string `json:"dark2"`
	Light1  string `json:"light1"`  // Background/text
	Light2  string `json:"light2"`
	Hyperlink string `json:"hyperlink"`
	FollowedHyperlink string `json:"followed_hyperlink"`
}

// FontScheme defines theme fonts
type FontScheme struct {
	MajorFont string `json:"major_font"` // Headings
	MinorFont string `json:"minor_font"` // Body text
}

//
// Test Configuration
//

// TestConfig configures test execution
type TestConfig struct {
	// Rendering
	RenderDPI       int              `json:"render_dpi"`       // DPI for PNG output (default: 300)
	RenderBackend   RenderBackend    `json:"render_backend"`   // LibreOffice, PowerPoint, Custom

	// Comparison
	DiffThreshold   float64          `json:"diff_threshold"`   // 0.0 to 1.0 (0 = identical, 1 = completely different)
	DiffAlgorithm   DiffAlgorithm    `json:"diff_algorithm"`   // PixelPerfect, Perceptual, SSIM
	IgnoreAntialiasing bool          `json:"ignore_antialiasing"` // Tolerate minor AA differences

	// Timeout
	GeneratorTimeout int             `json:"generator_timeout"` // Seconds
	RenderTimeout    int             `json:"render_timeout"`    // Seconds

	// Reporting
	GenerateDiffImage bool           `json:"generate_diff_image"` // Create red overlay PNG
	SaveIntermediates bool           `json:"save_intermediates"`  // Keep PDF, PPTX files

	// Retry
	RetryCount      int              `json:"retry_count"`      // Number of retries on failure
	RetryDelay      int              `json:"retry_delay"`      // Seconds between retries
}

// RenderBackend enum
type RenderBackend string

const (
	RenderBackendLibreOffice RenderBackend = "libreoffice"
	RenderBackendPowerPoint  RenderBackend = "powerpoint"
	RenderBackendPDF         RenderBackend = "pdf"           // Use goffice PDF rendering
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

// NewTestCase creates a test case with defaults
func NewTestCase(id, name string, category TestCategory) *TestCase {
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

// DefaultTestConfig returns sensible defaults
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
```

#### Example Test Case Definition

```go
// testcases/charts/bar_basic.go

package charts

import (
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// BarChartBasic creates a simple clustered bar chart test
func BarChartBasic() *framework.TestCase {
	tc := framework.NewTestCase(
		"chart_bar_clustered_basic",
		"Basic Clustered Bar Chart",
		framework.CategoryChart,
	)
	tc.Description = "Tests basic clustered column chart with 3 series and 4 categories"
	tc.Tags = []string{"chart", "bar", "clustered", "basic"}

	// Define slide with chart
	tc.Spec.SlideCount = 1
	tc.Spec.SlideSize = framework.SlideSize16x9
	tc.Spec.Slides = []framework.SlideSpec{
		{
			Index:  0,
			Layout: "Blank",
			Elements: []framework.ElementSpec{
				{
					Type: framework.ElementTypeChart,
					Position: framework.Position{
						X: 914400,  // 1 inch from left
						Y: 914400,  // 1 inch from top
					},
					Size: framework.Size{
						Width:  7315200, // 8 inches
						Height: 4572000, // 5 inches
					},
					Chart: &framework.ChartSpec{
						Type:  framework.ChartTypeColumn,
						Title: "Quarterly Sales by Region",
						Data: framework.ChartData{
							Categories: []string{"Q1", "Q2", "Q3", "Q4"},
							Series: []framework.SeriesData{
								{
									Name:   "North",
									Values: []float64{100, 120, 110, 130},
									Color:  "4472C4", // Blue
								},
								{
									Name:   "South",
									Values: []float64{80, 90, 95, 105},
									Color:  "ED7D31", // Orange
								},
								{
									Name:   "East",
									Values: []float64{70, 85, 90, 100},
									Color:  "A5A5A5", // Gray
								},
							},
						},
						Style: framework.ChartStyle{
							ColorPalette:   []string{"4472C4", "ED7D31", "A5A5A5", "FFC000", "5B9BD5"},
							FontSize:       12,
							FontFamily:     "Calibri",
							ShowDataLabels: false,
							GridLines: framework.GridLineSpec{
								MajorHorizontal: true,
								MinorHorizontal: false,
								MajorVertical:   false,
								MinorVertical:   false,
							},
						},
						Legend: &framework.LegendSpec{
							Position:   framework.LegendPositionRight,
							ShowLegend: true,
						},
						Axes: &framework.AxesSpec{
							Category: framework.AxisSpec{
								Title: "Quarter",
							},
							Value: framework.AxisSpec{
								Title:        "Sales ($K)",
								Min:          0,
								Max:          140,
								MajorUnit:    20,
								NumberFormat: "#,##0",
							},
						},
					},
				},
			},
		},
	}

	return tc
}
```

---

### 2. Generator Harness Design

#### Go Generator

```go
// generators/go/main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

func main() {
	// Parse flags
	testCaseFile := flag.String("test", "", "Path to test case JSON file")
	outputFile := flag.String("output", "", "Path to output PPTX file")
	flag.Parse()

	if *testCaseFile == "" || *outputFile == "" {
		log.Fatal("Usage: go-generator -test <testcase.json> -output <output.pptx>")
	}

	// Load test case
	data, err := os.ReadFile(*testCaseFile)
	if err != nil {
		log.Fatalf("Failed to read test case: %v", err)
	}

	tc, err := framework.FromJSON(data)
	if err != nil {
		log.Fatalf("Failed to parse test case: %v", err)
	}

	// Generate presentation
	if err := generatePresentation(tc, *outputFile); err != nil {
		log.Fatalf("Failed to generate presentation: %v", err)
	}

	fmt.Printf("Successfully generated: %s\n", *outputFile)
}

// generatePresentation creates the PPTX from test case spec
func generatePresentation(tc *framework.TestCase, outputPath string) error {
	// Create presentation document
	doc, err := presentation.New(outputPath, presentation.DocTypePresentation)
	if err != nil {
		return fmt.Errorf("create presentation: %w", err)
	}
	defer doc.Close()

	// Apply slide size
	if tc.Spec.SlideSize.Width > 0 && tc.Spec.SlideSize.Height > 0 {
		// Set slide size via presentation part
		presP art := doc.PresentationPart()
		if presPart != nil {
			pres := presPart.Presentation()
			if pres != nil {
				slideSize := pres.GetOrCreateSlideSize()
				slideSize.SetCx(tc.Spec.SlideSize.Width)
				slideSize.SetCy(tc.Spec.SlideSize.Height)
			}
		}
	}

	// Generate each slide
	for _, slideSpec := range tc.Spec.Slides {
		if err := generateSlide(doc, slideSpec); err != nil {
			return fmt.Errorf("generate slide %d: %w", slideSpec.Index, err)
		}
	}

	// Save
	if err := doc.Save(); err != nil {
		return fmt.Errorf("save presentation: %w", err)
	}

	return nil
}

// generateSlide creates a single slide with elements
func generateSlide(doc *presentation.Document, spec framework.SlideSpec) error {
	// Add slide
	slidePart, err := doc.AddSlide()
	if err != nil {
		return fmt.Errorf("add slide: %w", err)
	}

	slide := slidePart.Slide()
	if slide == nil {
		return fmt.Errorf("slide is nil")
	}

	// Apply background
	if spec.Background != nil {
		if err := applyBackground(slide, spec.Background); err != nil {
			return fmt.Errorf("apply background: %w", err)
		}
	}

	// Generate elements
	for i, elemSpec := range spec.Elements {
		if err := generateElement(slidePart, slide, elemSpec); err != nil {
			return fmt.Errorf("generate element %d: %w", i, err)
		}
	}

	return nil
}

// generateElement creates a single element (chart, shape, text, etc.)
func generateElement(slidePart *presentation.SlidePart, slide *presentation.Slide, spec framework.ElementSpec) error {
	switch spec.Type {
	case framework.ElementTypeChart:
		return generateChart(slidePart, slide, spec)
	case framework.ElementTypeShape:
		return generateShape(slide, spec)
	case framework.ElementTypeText:
		return generateText(slide, spec)
	case framework.ElementTypeImage:
		return generateImage(slidePart, slide, spec)
	case framework.ElementTypeTable:
		return generateTable(slide, spec)
	default:
		return fmt.Errorf("unknown element type: %s", spec.Type)
	}
}
```

```go
// generators/go/chart_gen.go

package main

import (
	"fmt"

	"github.com/connerohnesorge/goffice/presentation"
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/drawingml"
	"github.com/connerohnesorge/goffice/tests/e2e/framework"
)

// generateChart creates a chart element
func generateChart(slidePart *presentation.SlidePart, slide *presentation.Slide, spec framework.ElementSpec) error {
	if spec.Chart == nil {
		return fmt.Errorf("chart spec is nil")
	}

	chartSpec := spec.Chart

	// Create chart space
	chartSpace := drawingml.NewChartSpace()
	chart := chartSpace.Chart()

	// Set title
	if chartSpec.Title != "" {
		title := drawingml.NewTitleWithText(chartSpec.Title)
		chart.SetTitle(title)
	}

	// Get plot area
	plotArea := chart.PlotArea()
	if plotArea == nil {
		plotArea = drawingml.NewPlotArea()
		chart.SetPlotArea(plotArea)
	}

	// Generate chart type
	switch chartSpec.Type {
	case framework.ChartTypeColumn, framework.ChartTypeBarClustered:
		if err := generateBarChart(plotArea, chartSpec, framework.ChartTypeColumn); err != nil {
			return err
		}
	case framework.ChartTypeLine:
		if err := generateLineChart(plotArea, chartSpec, false); err != nil {
			return err
		}
	case framework.ChartTypeLineSmooth:
		if err := generateLineChart(plotArea, chartSpec, true); err != nil {
			return err
		}
	case framework.ChartTypePie:
		if err := generatePieChart(plotArea, chartSpec); err != nil {
			return err
		}
	case framework.ChartTypeDoughnut:
		if err := generateDoughnutChart(plotArea, chartSpec); err != nil {
			return err
		}
	case framework.ChartTypeArea:
		if err := generateAreaChart(plotArea, chartSpec); err != nil {
			return err
		}
	case framework.ChartTypeScatter:
		if err := generateScatterChart(plotArea, chartSpec, drawingml.ScatterStyleMarker); err != nil {
			return err
		}
	// Add other chart types...
	default:
		return fmt.Errorf("unsupported chart type: %s", chartSpec.Type)
	}

	// Add legend
	if chartSpec.Legend != nil && chartSpec.Legend.ShowLegend {
		legend := drawingml.NewLegendWithPosition(mapLegendPosition(chartSpec.Legend.Position))
		chart.SetLegend(legend)
	}

	// Create chart part and link to slide
	chartPart, err := slidePart.AddChartPart()
	if err != nil {
		return fmt.Errorf("add chart part: %w", err)
	}

	chartPart.SetRootElement(chartSpace)

	// Add graphic frame to slide
	shapeTree := slide.GetOrCreateShapeTree()
	gf := shapeTree.AddGraphicFrame()

	// Set position and size
	gf.SetPosition(spec.Position.X, spec.Position.Y)
	gf.SetSize(spec.Size.Width, spec.Size.Height)

	// Link to chart
	elements.LinkGraphicFrameToChart(gf, chartPart.RelationshipID())

	return nil
}

// generateBarChart creates a bar/column chart
func generateBarChart(plotArea *drawingml.PlotArea, spec *framework.ChartSpec, chartType framework.ChartType) error {
	// Determine direction and grouping
	direction := drawingml.BarDirectionCol // Column (vertical)
	if chartType == framework.ChartTypeBarClustered {
		direction = drawingml.BarDirectionBar // Bar (horizontal)
	}

	grouping := drawingml.BarGroupingClustered
	if spec.Type == framework.ChartTypeBarStacked {
		grouping = drawingml.BarGroupingStacked
	} else if spec.Type == framework.ChartTypeBarPercentStacked {
		grouping = drawingml.BarGroupingPercentStacked
	}

	// Create bar chart
	barChart := plotArea.AddBarChart(direction, grouping)

	// Add series
	for idx, series := range spec.Data.Series {
		barSeries := barChart.AddSeries(idx, idx)

		// Set series name
		barSeries.SetSeriesText(drawingml.NewSeriesTextWithValue(series.Name))

		// Set category data (X-axis labels)
		catData := drawingml.NewCategoryAxisData()
		catRef := drawingml.NewStringReferenceWithCache(
			fmt.Sprintf("Sheet1!$A$2:$A$%d", len(spec.Data.Categories)+1),
			spec.Data.Categories,
		)
		catData.AppendChild(catRef)
		barSeries.SetCategoryAxisData(catData)

		// Set values (Y-axis data)
		vals := drawingml.NewValues()
		numRef := drawingml.NewNumberReferenceWithCache(
			fmt.Sprintf("Sheet1!$%c$2:$%c$%d", 'B'+idx, 'B'+idx, len(series.Values)+1),
			series.Values,
		)
		vals.AppendChild(numRef)
		barSeries.SetValues(vals)

		// Apply color
		if series.Color != "" {
			spPr := drawingml.NewChartShapeProperties()
			spPr.SetSolidFill(series.Color)
			barSeries.SetShapeProperties(spPr)
		}
	}

	// Add axes if specified
	if spec.Axes != nil {
		// Category axis
		catAxisID := 1
		valAxisID := 2

		barChart.AddAxisID(catAxisID)
		barChart.AddAxisID(valAxisID)

		catAxis := plotArea.AddCategoryAxis(catAxisID, valAxisID)
		configureAxis(catAxis, spec.Axes.Category)

		valAxis := plotArea.AddValueAxis(valAxisID, catAxisID)
		configureAxis(valAxis, spec.Axes.Value)
	}

	return nil
}

// generateLineChart creates a line chart
func generateLineChart(plotArea *drawingml.PlotArea, spec *framework.ChartSpec, smooth bool) error {
	grouping := drawingml.GroupingStandard
	if spec.Type == framework.ChartTypeLineStacked {
		grouping = drawingml.GroupingStacked
	}

	lineChart := plotArea.AddLineChart(grouping)

	// Add series
	for idx, series := range spec.Data.Series {
		lineSeries := lineChart.AddSeries(idx, idx)

		lineSeries.SetSeriesText(drawingml.NewSeriesTextWithValue(series.Name))

		// Category data
		catData := drawingml.NewCategoryAxisData()
		catRef := drawingml.NewStringReferenceWithCache(
			fmt.Sprintf("Sheet1!$A$2:$A$%d", len(spec.Data.Categories)+1),
			spec.Data.Categories,
		)
		catData.AppendChild(catRef)
		lineSeries.SetCategoryAxisData(catData)

		// Values
		vals := drawingml.NewValues()
		numRef := drawingml.NewNumberReferenceWithCache(
			fmt.Sprintf("Sheet1!$%c$2:$%c$%d", 'B'+idx, 'B'+idx, len(series.Values)+1),
			series.Values,
		)
		vals.AppendChild(numRef)
		lineSeries.SetValues(vals)

		// Apply color
		if series.Color != "" {
			spPr := drawingml.NewChartShapeProperties()
			spPr.SetOutline(12700, series.Color) // 1pt line width
			lineSeries.SetShapeProperties(spPr)
		}

		// Set smooth
		if smooth || series.Smooth {
			lineSeries.SetSmooth(true)
		}

		// Set marker style
		if series.MarkerStyle != "" {
			marker := drawingml.NewMarker()
			marker.SetSymbol(mapMarkerStyle(series.MarkerStyle))
			marker.SetSize(5) // 5pt markers
			lineSeries.SetMarker(marker)
		}
	}

	// Add axes
	if spec.Axes != nil {
		catAxisID := 1
		valAxisID := 2

		lineChart.AddAxisID(catAxisID)
		lineChart.AddAxisID(valAxisID)

		catAxis := plotArea.AddCategoryAxis(catAxisID, valAxisID)
		configureAxis(catAxis, spec.Axes.Category)

		valAxis := plotArea.AddValueAxis(valAxisID, catAxisID)
		configureAxis(valAxis, spec.Axes.Value)
	}

	return nil
}

// generatePieChart creates a pie chart
func generatePieChart(plotArea *drawingml.PlotArea, spec *framework.ChartSpec) error {
	pieChart := plotArea.AddPieChart()

	// Pie charts typically have one series
	if len(spec.Data.Series) == 0 {
		return fmt.Errorf("pie chart requires at least one series")
	}

	series := spec.Data.Series[0]
	pieSeries := pieChart.AddSeries(0, 0)

	pieSeries.SetSeriesText(drawingml.NewSeriesTextWithValue(series.Name))

	// Category data (slice names)
	catData := drawingml.NewCategoryAxisData()
	catRef := drawingml.NewStringReferenceWithCache(
		fmt.Sprintf("Sheet1!$A$2:$A$%d", len(spec.Data.Categories)+1),
		spec.Data.Categories,
	)
	catData.AppendChild(catRef)
	pieSeries.SetCategoryAxisData(catData)

	// Values
	vals := drawingml.NewValues()
	numRef := drawingml.NewNumberReferenceWithCache(
		fmt.Sprintf("Sheet1!$B$2:$B$%d", len(series.Values)+1),
		series.Values,
	)
	vals.AppendChild(numRef)
	pieSeries.SetValues(vals)

	// Color each data point differently
	if len(spec.Style.ColorPalette) > 0 {
		for i := range spec.Data.Categories {
			colorIdx := i % len(spec.Style.ColorPalette)
			color := spec.Style.ColorPalette[colorIdx]

			dPt := drawingml.NewDataPoint()
			dPt.SetIndex(i)
			spPr := drawingml.NewChartShapeProperties()
			spPr.SetSolidFill(color)
			dPt.SetShapeProperties(spPr)
			pieSeries.AppendChild(dPt)
		}
	}

	return nil
}

// configureAxis applies axis spec settings
func configureAxis(axis interface{}, spec framework.AxisSpec) {
	// This is a simplified version - actual implementation would use type assertion
	// to access category/value/date axis specific methods

	// Example for value axis:
	// if valAxis, ok := axis.(*drawingml.ValueAxis); ok {
	//     if spec.Min > 0 {
	//         valAxis.SetMinimum(spec.Min)
	//     }
	//     if spec.Max > 0 {
	//         valAxis.SetMaximum(spec.Max)
	//     }
	//     // ... more configuration
	// }
}

// Helper mappers
func mapLegendPosition(pos framework.LegendPosition) drawingml.LegendPosition {
	switch pos {
	case framework.LegendPositionTop:
		return drawingml.LegendPositionTop
	case framework.LegendPositionBottom:
		return drawingml.LegendPositionBottom
	case framework.LegendPositionLeft:
		return drawingml.LegendPositionLeft
	case framework.LegendPositionRight:
		return drawingml.LegendPositionRight
	case framework.LegendPositionTopRight:
		return drawingml.LegendPositionTopRight
	default:
		return drawingml.LegendPositionRight
	}
}

func mapMarkerStyle(style string) drawingml.MarkerStyle {
	switch style {
	case "circle":
		return drawingml.MarkerStyleCircle
	case "diamond":
		return drawingml.MarkerStyleDiamond
	case "square":
		return drawingml.MarkerStyleSquare
	case "star":
		return drawingml.MarkerStyleStar
	case "triangle":
		return drawingml.MarkerStyleTriangle
	case "x":
		return drawingml.MarkerStyleX
	case "plus":
		return drawingml.MarkerStylePlus
	default:
		return drawingml.MarkerStyleCircle
	}
}
```

#### C# Generator

```csharp
// generators/csharp/Program.cs

using System;
using System.CommandLine;
using System.IO;
using System.Text.Json;
using DocumentFormat.OpenXml;
using DocumentFormat.OpenXml.Packaging;
using DocumentFormat.OpenXml.Presentation;
using D = DocumentFormat.OpenXml.Drawing;
using C = DocumentFormat.OpenXml.Drawing.Charts;
using P = DocumentFormat.OpenXml.Presentation;

namespace E2EGenerator
{
    class Program
    {
        static async Task<int> Main(string[] args)
        {
            var testOption = new Option<FileInfo>(
                name: "--test",
                description: "Path to test case JSON file"
            ) { IsRequired = true };

            var outputOption = new Option<FileInfo>(
                name: "--output",
                description: "Path to output PPTX file"
            ) { IsRequired = true };

            var rootCommand = new RootCommand("C# PPTX Generator for E2E Testing");
            rootCommand.AddOption(testOption);
            rootCommand.AddOption(outputOption);

            rootCommand.SetHandler(async (testFile, outputFile) =>
            {
                await GeneratePresentation(testFile, outputFile);
            }, testOption, outputOption);

            return await rootCommand.InvokeAsync(args);
        }

        static async Task GeneratePresentation(FileInfo testFile, FileInfo outputFile)
        {
            // Load test case
            var json = await File.ReadAllTextAsync(testFile.FullName);
            var testCase = JsonSerializer.Deserialize<TestCase>(json);

            if (testCase == null)
            {
                throw new InvalidOperationException("Failed to deserialize test case");
            }

            // Create presentation
            using (var pres = PresentationDocument.Create(outputFile.FullName, PresentationDocumentType.Presentation))
            {
                var presPart = pres.AddPresentationPart();
                presPart.Presentation = new Presentation();

                // Initialize presentation structure
                var slideIdList = new SlideIdList();
                var slideMasterIdList = new SlideMasterIdList();

                // Set slide size
                var slideSize = new SlideSize
                {
                    Cx = testCase.Spec.SlideSize.Width,
                    Cy = testCase.Spec.SlideSize.Height
                };

                presPart.Presentation.Append(
                    slideMasterIdList,
                    new NotesSize { Cx = 6858000, Cy = 9144000 },
                    slideIdList,
                    slideSize
                );

                // Create slides
                uint slideId = 256;
                foreach (var slideSpec in testCase.Spec.Slides)
                {
                    var slidePart = CreateSlidePart(presPart, slideSpec);

                    var slideId Entry = new SlideId
                    {
                        Id = slideId++,
                        RelationshipId = presPart.GetIdOfPart(slidePart)
                    };
                    slideIdList.Append(slideIdEntry);
                }

                presPart.Presentation.Save();
            }

            Console.WriteLine($"Successfully generated: {outputFile.FullName}");
        }

        static SlidePart CreateSlidePart(PresentationPart presPart, SlideSpec spec)
        {
            var slidePart = presPart.AddNewPart<SlidePart>();

            // Create slide with shape tree
            slidePart.Slide = new Slide(
                new CommonSlideData(
                    new ShapeTree(
                        new P.NonVisualGroupShapeProperties(
                            new P.NonVisualDrawingProperties { Id = 1, Name = "" },
                            new P.NonVisualGroupShapeDrawingProperties(),
                            new ApplicationNonVisualDrawingProperties()),
                        new P.GroupShapeProperties(
                            new D.TransformGroup()))),
                new ColorMapOverride(
                    new D.MasterColorMapping()));

            var shapeTree = slidePart.Slide.CommonSlideData.ShapeTree;

            // Add elements
            uint shapeId = 2;
            foreach (var elemSpec in spec.Elements)
            {
                switch (elemSpec.Type)
                {
                    case "chart":
                        AddChart(slidePart, shapeTree, elemSpec, ref shapeId);
                        break;
                    case "shape":
                        AddShape(shapeTree, elemSpec, ref shapeId);
                        break;
                    case "text":
                        AddText(shapeTree, elemSpec, ref shapeId);
                        break;
                    // ... other types
                }
            }

            slidePart.Slide.Save();
            return slidePart;
        }

        static void AddChart(SlidePart slidePart, ShapeTree shapeTree, ElementSpec spec, ref uint shapeId)
        {
            if (spec.Chart == null) return;

            // Create chart part
            var chartPart = slidePart.AddNewPart<ChartPart>();

            // Generate chart XML
            var chartSpace = CreateChartSpace(spec.Chart);
            chartPart.ChartSpace = chartSpace;
            chartPart.ChartSpace.Save();

            // Create graphic frame to host chart
            var graphicFrame = new GraphicFrame(
                new P.NonVisualGraphicFrameProperties(
                    new P.NonVisualDrawingProperties
                    {
                        Id = shapeId++,
                        Name = $"Chart {shapeId}"
                    },
                    new P.NonVisualGraphicFrameDrawingProperties(),
                    new ApplicationNonVisualDrawingProperties()),
                new P.Transform(
                    new D.Offset { X = spec.Position.X, Y = spec.Position.Y },
                    new D.Extents { Cx = spec.Size.Width, Cy = spec.Size.Height }),
                new D.Graphic(
                    new D.GraphicData(
                        new C.ChartReference { Id = slidePart.GetIdOfPart(chartPart) }
                    )
                    { Uri = "http://schemas.openxmlformats.org/drawingml/2006/chart" }));

            shapeTree.AppendChild(graphicFrame);
        }

        static C.ChartSpace CreateChartSpace(ChartSpec spec)
        {
            var chartSpace = new C.ChartSpace();
            chartSpace.AddNamespaceDeclaration("c", "http://schemas.openxmlformats.org/drawingml/2006/chart");
            chartSpace.AddNamespaceDeclaration("a", "http://schemas.openxmlformats.org/drawingml/2006/main");
            chartSpace.AddNamespaceDeclaration("r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships");

            var chart = new C.Chart();

            // Add title
            if (!string.IsNullOrEmpty(spec.Title))
            {
                chart.Append(new C.Title(
                    new C.ChartText(
                        new C.RichText(
                            new D.BodyProperties(),
                            new D.ListStyle(),
                            new D.Paragraph(
                                new D.Run(
                                    new D.RunProperties { FontSize = 1800 },
                                    new D.Text(spec.Title)))))));
            }

            var plotArea = new C.PlotArea();

            // Create chart type
            switch (spec.Type)
            {
                case "column":
                case "bar_clustered":
                    CreateBarChart(plotArea, spec);
                    break;
                case "line":
                    CreateLineChart(plotArea, spec);
                    break;
                case "pie":
                    CreatePieChart(plotArea, spec);
                    break;
                // ... other types
            }

            // Add legend
            if (spec.Legend?.ShowLegend == true)
            {
                chart.Append(CreateLegend(spec.Legend));
            }

            chart.Append(plotArea);
            chartSpace.Append(chart);

            return chartSpace;
        }

        static void CreateBarChart(C.PlotArea plotArea, ChartSpec spec)
        {
            var barChart = new C.BarChart(
                new C.BarDirection { Val = C.BarDirectionValues.Column },
                new C.BarGrouping { Val = C.BarGroupingValues.Clustered });

            uint seriesIdx = 0;
            foreach (var series in spec.Data.Series)
            {
                var barChartSeries = new C.BarChartSeries(
                    new C.Index { Val = seriesIdx },
                    new C.Order { Val = seriesIdx });

                // Series text
                barChartSeries.Append(new C.SeriesText(
                    new C.StringReference(
                        new C.StringCache(
                            new C.PointCount { Val = 1U },
                            new C.StringPoint(
                                new C.NumericValue(series.Name))
                            { Index = 0U }))));

                // Category axis data
                var catAxisData = new C.CategoryAxisData();
                var strRef = new C.StringReference();
                strRef.Append(new C.Formula($"Sheet1!$A$2:$A${spec.Data.Categories.Count + 1}"));

                var strCache = new C.StringCache();
                strCache.Append(new C.PointCount { Val = (uint)spec.Data.Categories.Count });
                for (int i = 0; i < spec.Data.Categories.Count; i++)
                {
                    strCache.Append(new C.StringPoint(
                        new C.NumericValue(spec.Data.Categories[i]))
                    { Index = (uint)i });
                }
                strRef.Append(strCache);
                catAxisData.Append(strRef);
                barChartSeries.Append(catAxisData);

                // Values
                var vals = new C.Values();
                var numRef = new C.NumberReference();
                numRef.Append(new C.Formula($"Sheet1!${(char)('B' + seriesIdx)}$2:${(char)('B' + seriesIdx)}${series.Values.Count + 1}"));

                var numCache = new C.NumberingCache();
                numCache.Append(new C.FormatCode("General"));
                numCache.Append(new C.PointCount { Val = (uint)series.Values.Count });
                for (int i = 0; i < series.Values.Count; i++)
                {
                    numCache.Append(new C.NumericPoint(
                        new C.NumericValue(series.Values[i].ToString()))
                    { Index = (uint)i });
                }
                numRef.Append(numCache);
                vals.Append(numRef);
                barChartSeries.Append(vals);

                // Apply color
                if (!string.IsNullOrEmpty(series.Color))
                {
                    var spPr = new C.ChartShapeProperties(
                        new D.SolidFill(
                            new D.RgbColorModelHex { Val = series.Color }));
                    barChartSeries.Append(spPr);
                }

                barChart.Append(barChartSeries);
                seriesIdx++;
            }

            // Add axes
            barChart.Append(new C.AxisId { Val = 1U });
            barChart.Append(new C.AxisId { Val = 2U });

            plotArea.Append(barChart);

            // Category axis
            var catAx = new C.CategoryAxis(
                new C.AxisId { Val = 1U },
                new C.Scaling(new C.Orientation { Val = C.OrientationValues.MinMax }),
                new C.AxisPosition { Val = C.AxisPositionValues.Bottom },
                new C.CrossingAxis { Val = 2U });
            plotArea.Append(catAx);

            // Value axis
            var valAx = new C.ValueAxis(
                new C.AxisId { Val = 2U },
                new C.Scaling(new C.Orientation { Val = C.OrientationValues.MinMax }),
                new C.AxisPosition { Val = C.AxisPositionValues.Left },
                new C.MajorGridlines(),
                new C.CrossingAxis { Val = 1U });

            // Apply axis specs
            if (spec.Axes != null)
            {
                if (spec.Axes.Value.Min > 0 || spec.Axes.Value.Max > 0)
                {
                    var scaling = valAx.GetFirstChild<C.Scaling>();
                    if (spec.Axes.Value.Min > 0)
                        scaling.Append(new C.MinAxisValue { Val = spec.Axes.Value.Min });
                    if (spec.Axes.Value.Max > 0)
                        scaling.Append(new C.MaxAxisValue { Val = spec.Axes.Value.Max });
                }
            }

            plotArea.Append(valAx);
        }

        static C.Legend CreateLegend(LegendSpec spec)
        {
            C.LegendPositionValues position = spec.Position switch
            {
                "top" => C.LegendPositionValues.Top,
                "bottom" => C.LegendPositionValues.Bottom,
                "left" => C.LegendPositionValues.Left,
                "right" => C.LegendPositionValues.Right,
                "top_right" => C.LegendPositionValues.TopRight,
                _ => C.LegendPositionValues.Right
            };

            return new C.Legend(
                new C.LegendPosition { Val = position },
                new C.Overlay { Val = false });
        }
    }

    // Data classes matching Go framework types
    public class TestCase
    {
        public string Id { get; set; }
        public string Name { get; set; }
        public TestSpec Spec { get; set; }
    }

    public class TestSpec
    {
        public int SlideCount { get; set; }
        public SlideSize SlideSize { get; set; }
        public List<SlideSpec> Slides { get; set; }
    }

    public class SlideSize
    {
        public long Width { get; set; }
        public long Height { get; set; }
    }

    public class SlideSpec
    {
        public int Index { get; set; }
        public List<ElementSpec> Elements { get; set; }
    }

    public class ElementSpec
    {
        public string Type { get; set; }
        public Position Position { get; set; }
        public Size Size { get; set; }
        public ChartSpec Chart { get; set; }
    }

    public class Position
    {
        public long X { get; set; }
        public long Y { get; set; }
    }

    public class Size
    {
        public long Width { get; set; }
        public long Height { get; set; }
    }

    public class ChartSpec
    {
        public string Type { get; set; }
        public string Title { get; set; }
        public ChartData Data { get; set; }
        public ChartStyle Style { get; set; }
        public LegendSpec Legend { get; set; }
        public AxesSpec Axes { get; set; }
    }

    public class ChartData
    {
        public List<string> Categories { get; set; }
        public List<SeriesData> Series { get; set; }
    }

    public class SeriesData
    {
        public string Name { get; set; }
        public List<double> Values { get; set; }
        public string Color { get; set; }
    }

    public class ChartStyle
    {
        public List<string> ColorPalette { get; set; }
    }

    public class LegendSpec
    {
        public string Position { get; set; }
        public bool ShowLegend { get; set; }
    }

    public class AxesSpec
    {
        public AxisSpec Category { get; set; }
        public AxisSpec Value { get; set; }
    }

    public class AxisSpec
    {
        public string Title { get; set; }
        public double Min { get; set; }
        public double Max { get; set; }
    }
}
```

---

### 3. Rendering Pipeline

```go
// framework/renderer.go

package framework

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Renderer converts PPTX to PNG images
type Renderer interface {
	Render(ctx context.Context, pptxPath, outputDir string) ([]string, error)
}

// LibreOfficeRenderer uses LibreOffice headless for rendering
type LibreOfficeRenderer struct {
	DPI           int
	Format        string // "png", "pdf"
	LibreOfficeBin string // Path to soffice binary
}

// NewLibreOfficeRenderer creates a renderer with defaults
func NewLibreOfficeRenderer() *LibreOfficeRenderer {
	return &LibreOfficeRenderer{
		DPI:           300,
		Format:        "pdf", // First convert to PDF, then to PNG
		LibreOfficeBin: findLibreOffice(),
	}
}

// Render converts PPTX to PNG images (one per slide)
func (r *LibreOfficeRenderer) Render(ctx context.Context, pptxPath, outputDir string) ([]string, error) {
	// Step 1: Convert PPTX → PDF using LibreOffice
	pdfPath, err := r.convertToPDF(ctx, pptxPath, outputDir)
	if err != nil {
		return nil, fmt.Errorf("convert to PDF: %w", err)
	}

	// Step 2: Convert PDF → PNG images using pdftoppm or ImageMagick
	pngPaths, err := r.convertPDFToPNG(ctx, pdfPath, outputDir)
	if err != nil {
		return nil, fmt.Errorf("convert PDF to PNG: %w", err)
	}

	return pngPaths, nil
}

// convertToPDF uses LibreOffice headless to convert PPTX to PDF
func (r *LibreOfficeRenderer) convertToPDF(ctx context.Context, pptxPath, outputDir string) (string, error) {
	// LibreOffice command:
	// soffice --headless --convert-to pdf --outdir <outputDir> <pptxPath>

	cmd := exec.CommandContext(ctx, r.LibreOfficeBin,
		"--headless",
		"--convert-to", "pdf",
		"--outdir", outputDir,
		pptxPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("libreoffice failed: %w\nOutput: %s", err, output)
	}

	// Determine output PDF path
	baseName := filepath.Base(pptxPath)
	pdfName := baseName[:len(baseName)-len(filepath.Ext(baseName))] + ".pdf"
	pdfPath := filepath.Join(outputDir, pdfName)

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return "", fmt.Errorf("PDF not created: %s", pdfPath)
	}

	return pdfPath, nil
}

// convertPDFToPNG extracts each PDF page as a PNG image
func (r *LibreOfficeRenderer) convertPDFToPNG(ctx context.Context, pdfPath, outputDir string) ([]string, error) {
	// Use pdftoppm (from Poppler) for high-quality PNG extraction:
	// pdftoppm -png -r 300 input.pdf output_prefix

	baseName := filepath.Base(pdfPath)
	prefix := baseName[:len(baseName)-len(filepath.Ext(baseName))]
	outputPrefix := filepath.Join(outputDir, prefix)

	cmd := exec.CommandContext(ctx, "pdftoppm",
		"-png",
		"-r", fmt.Sprintf("%d", r.DPI),
		pdfPath,
		outputPrefix,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("pdftoppm failed: %w\nOutput: %s", err, output)
	}

	// pdftoppm outputs: prefix-1.png, prefix-2.png, ...
	// Collect generated PNGs
	pattern := fmt.Sprintf("%s-*.png", outputPrefix)
	pngPaths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("glob PNGs: %w", err)
	}

	if len(pngPaths) == 0 {
		return nil, fmt.Errorf("no PNG files generated from %s", pdfPath)
	}

	return pngPaths, nil
}

// findLibreOffice locates the LibreOffice binary
func findLibreOffice() string {
	// Try common paths
	candidates := []string{
		"/usr/bin/libreoffice",
		"/usr/bin/soffice",
		"/Applications/LibreOffice.app/Contents/MacOS/soffice",
		"C:\\Program Files\\LibreOffice\\program\\soffice.exe",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Fallback to PATH
	if path, err := exec.LookPath("soffice"); err == nil {
		return path
	}

	return "soffice" // Hope it's in PATH
}

// PowerPointRenderer uses PowerPoint automation (Windows/COM) for rendering
type PowerPointRenderer struct {
	DPI int
}

// NewPowerPointRenderer creates a PowerPoint-based renderer
func NewPowerPointRenderer() *PowerPointRenderer {
	return &PowerPointRenderer{
		DPI: 300,
	}
}

// Render uses PowerPoint to export slides as PNG
func (r *PowerPointRenderer) Render(ctx context.Context, pptxPath, outputDir string) ([]string, error) {
	// This requires PowerPoint COM automation (Windows only)
	// Implemented via external PowerShell script or Go COM bindings

	// Example PowerShell script:
	// $ppt = New-Object -ComObject PowerPoint.Application
	// $pres = $ppt.Presentations.Open($pptxPath)
	// foreach ($i in 1..$pres.Slides.Count) {
	//     $pres.Slides[$i].Export("slide_$i.png", "PNG")
	// }
	// $pres.Close()
	// $ppt.Quit()

	script := fmt.Sprintf(`
		$ppt = New-Object -ComObject PowerPoint.Application
		$ppt.Visible = [Microsoft.Office.Core.MsoTriState]::msoFalse
		$pres = $ppt.Presentations.Open("%s", [Microsoft.Office.Core.MsoTriState]::msoFalse)
		$slideCount = $pres.Slides.Count
		for ($i = 1; $i -le $slideCount; $i++) {
			$outputPath = Join-Path "%s" ("slide_" + $i + ".png")
			$pres.Slides.Item($i).Export($outputPath, "PNG", %d, %d)
		}
		$pres.Close()
		$ppt.Quit()
		[System.Runtime.Interopservices.Marshal]::ReleaseComObject($ppt) | Out-Null
	`, pptxPath, outputDir, r.DPI*10, r.DPI*7) // Approximate 10x7 inch slide

	cmd := exec.CommandContext(ctx, "powershell", "-Command", script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("PowerPoint automation failed: %w\nOutput: %s", err, output)
	}

	// Collect generated PNGs
	pattern := filepath.Join(outputDir, "slide_*.png")
	pngPaths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("glob PNGs: %w", err)
	}

	return pngPaths, nil
}

// PDFRenderer uses goffice's own PDF renderer (for Go output only)
type PDFRenderer struct {
	DPI int
}

// NewPDFRenderer creates a PDF-based renderer
func NewPDFRenderer() *PDFRenderer {
	return &PDFRenderer{
		DPI: 300,
	}
}

// Render uses goffice PDF rendering (only works for Go-generated PPTX)
func (r *PDFRenderer) Render(ctx context.Context, pptxPath, outputDir string) ([]string, error) {
	// This would use goffice's pdf package to render directly
	// Not applicable for C# output comparison, but useful for PDF rendering tests

	// Example (pseudocode):
	// doc, err := presentation.Open(pptxPath, false)
	// pdfPath := filepath.Join(outputDir, "output.pdf")
	// pdf.RenderPresentation(doc, pdfPath)
	// return convertPDFToPNG(pdfPath, outputDir)

	return nil, fmt.Errorf("PDFRenderer not implemented yet")
}
```

---

### 4. Visual Comparison Engine

```go
// framework/differ.go

package framework

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/exec"
	"path/filepath"
)

// DiffResult contains comparison metrics
type DiffResult struct {
	Image1Path  string
	Image2Path  string
	DiffPath    string

	// Metrics
	PixelDiff   int     // Number of different pixels
	TotalPixels int     // Total pixel count
	DiffPercent float64 // Percentage difference (0-100)
	SSIM        float64 // Structural Similarity Index (-1 to 1, 1 = identical)
	MSE         float64 // Mean Squared Error
	PSNR        float64 // Peak Signal-to-Noise Ratio (dB)

	// Result
	Passed      bool    // Did comparison pass threshold?
	Threshold   float64 // Configured threshold
}

// Differ performs visual comparison
type Differ interface {
	Compare(ctx context.Context, img1Path, img2Path, diffPath string, config TestConfig) (*DiffResult, error)
}

// ImageMagickDiffer uses ImageMagick's compare tool
type ImageMagickDiffer struct{}

// NewImageMagickDiffer creates an ImageMagick-based differ
func NewImageMagickDiffer() *ImageMagickDiffer {
	return &ImageMagickDiffer{}
}

// Compare uses ImageMagick compare command
func (d *ImageMagickDiffer) Compare(ctx context.Context, img1Path, img2Path, diffPath string, config TestConfig) (*DiffResult, error) {
	result := &DiffResult{
		Image1Path: img1Path,
		Image2Path: img2Path,
		DiffPath:   diffPath,
		Threshold:  config.DiffThreshold,
	}

	// ImageMagick compare command:
	// compare -metric RMSE img1.png img2.png diff.png
	// Outputs to stderr: <value> (<normalized>)

	metricName := mapDiffAlgorithmToImageMagick(config.DiffAlgorithm)

	cmd := exec.CommandContext(ctx, "compare",
		"-metric", metricName,
		img1Path,
		img2Path,
		diffPath,
	)

	output, err := cmd.CombinedOutput()
	// Note: compare returns exit code 1 if images differ, not an error

	// Parse output (format: "1234.5 (0.0123)")
	var rawValue, normalizedValue float64
	_, parseErr := fmt.Sscanf(string(output), "%f (%f)", &rawValue, &normalizedValue)
	if parseErr != nil {
		return nil, fmt.Errorf("parse compare output: %w\nOutput: %s", parseErr, output)
	}

	// Populate result based on algorithm
	switch config.DiffAlgorithm {
	case DiffAlgorithmSSIM:
		result.SSIM = normalizedValue
		result.DiffPercent = (1 - normalizedValue) * 100
		result.Passed = result.DiffPercent <= config.DiffThreshold*100

	case DiffAlgorithmMSE:
		result.MSE = rawValue
		result.DiffPercent = normalizedValue * 100
		result.Passed = result.DiffPercent <= config.DiffThreshold*100

	case DiffAlgorithmPSNR:
		result.PSNR = rawValue
		// PSNR > 30 dB is typically considered good
		result.DiffPercent = math.Max(0, (50-rawValue)/50*100)
		result.Passed = result.DiffPercent <= config.DiffThreshold*100

	default:
		// Pixel-based comparison
		result.DiffPercent = normalizedValue * 100
		result.Passed = result.DiffPercent <= config.DiffThreshold*100
	}

	return result, nil
}

func mapDiffAlgorithmToImageMagick(alg DiffAlgorithm) string {
	switch alg {
	case DiffAlgorithmSSIM:
		return "SSIM"
	case DiffAlgorithmMSE:
		return "MSE"
	case DiffAlgorithmPSNR:
		return "PSNR"
	case DiffAlgorithmPerceptual:
		return "PHASH" // Perceptual hash
	default:
		return "AE" // Absolute error (pixel count)
	}
}

// GoDiffer implements pure Go image comparison
type GoDiffer struct{}

// NewGoDiffer creates a Go-native differ
func NewGoDiffer() *GoDiffer {
	return &GoDiffer{}
}

// Compare performs pixel-by-pixel comparison in pure Go
func (d *GoDiffer) Compare(ctx context.Context, img1Path, img2Path, diffPath string, config TestConfig) (*DiffResult, error) {
	// Load images
	img1, err := loadPNG(img1Path)
	if err != nil {
		return nil, fmt.Errorf("load image1: %w", err)
	}

	img2, err := loadPNG(img2Path)
	if err != nil {
		return nil, fmt.Errorf("load image2: %w", err)
	}

	// Check dimensions
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	if bounds1 != bounds2 {
		return nil, fmt.Errorf("image dimensions differ: %v vs %v", bounds1, bounds2)
	}

	// Create diff image
	diffImg := image.NewRGBA(bounds1)

	// Compare pixels
	var diffPixels int
	totalPixels := bounds1.Dx() * bounds1.Dy()
	var mse float64

	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)

			r1, g1, b1, a1 := c1.RGBA()
			r2, g2, b2, a2 := c2.RGBA()

			// Calculate difference
			dr := int(r1>>8) - int(r2>>8)
			dg := int(g1>>8) - int(g2>>8)
			db := int(b1>>8) - int(b2>>8)
			da := int(a1>>8) - int(a2>>8)

			// MSE contribution
			mse += float64(dr*dr + dg*dg + db*db + da*da)

			// Check if pixel differs
			if dr != 0 || dg != 0 || db != 0 || da != 0 {
				diffPixels++
				// Highlight diff in red
				diffImg.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			} else {
				// Keep original (grayscale)
				gray := uint8((r1 >> 8) / 3)
				diffImg.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: 255})
			}
		}
	}

	// Calculate metrics
	mse = mse / float64(totalPixels*4) // 4 channels
	psnr := 10 * math.Log10(255*255/mse)

	// Save diff image
	if err := savePNG(diffPath, diffImg); err != nil {
		return nil, fmt.Errorf("save diff image: %w", err)
	}

	// Build result
	diffPercent := float64(diffPixels) / float64(totalPixels) * 100

	result := &DiffResult{
		Image1Path:  img1Path,
		Image2Path:  img2Path,
		DiffPath:    diffPath,
		PixelDiff:   diffPixels,
		TotalPixels: totalPixels,
		DiffPercent: diffPercent,
		MSE:         mse,
		PSNR:        psnr,
		Threshold:   config.DiffThreshold,
		Passed:      diffPercent <= config.DiffThreshold*100,
	}

	// Calculate SSIM if needed
	if config.DiffAlgorithm == DiffAlgorithmSSIM {
		result.SSIM = calculateSSIM(img1, img2)
		result.Passed = result.SSIM >= (1 - config.DiffThreshold)
	}

	return result, nil
}

// calculateSSIM computes Structural Similarity Index
func calculateSSIM(img1, img2 image.Image) float64 {
	// Simplified SSIM calculation (windowed approach omitted for brevity)
	bounds := img1.Bounds()

	var meanX, meanY, varX, varY, covXY float64
	n := float64(bounds.Dx() * bounds.Dy())

	// Calculate means
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()

			gray1 := float64(r1+g1+b1) / (3 * 65535)
			gray2 := float64(r2+g2+b2) / (3 * 65535)

			meanX += gray1
			meanY += gray2
		}
	}
	meanX /= n
	meanY /= n

	// Calculate variances and covariance
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()

			gray1 := float64(r1+g1+b1) / (3 * 65535)
			gray2 := float64(r2+g2+b2) / (3 * 65535)

			varX += (gray1 - meanX) * (gray1 - meanX)
			varY += (gray2 - meanY) * (gray2 - meanY)
			covXY += (gray1 - meanX) * (gray2 - meanY)
		}
	}
	varX /= n
	varY /= n
	covXY /= n

	// SSIM formula
	c1 := 0.01 * 0.01
	c2 := 0.03 * 0.03

	numerator := (2*meanX*meanY + c1) * (2*covXY + c2)
	denominator := (meanX*meanX + meanY*meanY + c1) * (varX + varY + c2)

	return numerator / denominator
}

// Helper functions
func loadPNG(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func savePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
```

---

### 5. Report Generation

```go
// framework/reporter.go

package framework

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

// TestReport aggregates all test results
type TestReport struct {
	Generated   time.Time
	TotalTests  int
	PassedTests int
	FailedTests int
	Results     []TestResult

	// Summary stats
	AvgDiffPercent float64
	MaxDiffPercent float64
	MinDiffPercent float64
}

// TestResult stores result of a single test case
type TestResult struct {
	TestCase    *TestCase
	Status      TestStatus
	GoOutput    GeneratorOutput
	CSharpOutput GeneratorOutput
	DiffResults []DiffResult // One per slide
	Duration    time.Duration
	Error       string
}

// GeneratorOutput stores paths to generated artifacts
type GeneratorOutput struct {
	PPTXPath string
	PDFPath  string
	PNGPaths []string
}

// TestStatus enum
type TestStatus string

const (
	TestStatusPass    TestStatus = "pass"
	TestStatusFail    TestStatus = "fail"
	TestStatusError   TestStatus = "error"
	TestStatusSkipped TestStatus = "skipped"
)

// Reporter generates HTML reports
type Reporter struct {
	OutputDir  string
	AssetsDir  string
	Template   *template.Template
}

// NewReporter creates a reporter
func NewReporter(outputDir string) (*Reporter, error) {
	r := &Reporter{
		OutputDir: outputDir,
		AssetsDir: filepath.Join(outputDir, "assets"),
	}

	// Create directories
	if err := os.MkdirAll(r.AssetsDir, 0755); err != nil {
		return nil, fmt.Errorf("create assets dir: %w", err)
	}

	// Load template
	tmpl, err := template.New("report").Funcs(templateFuncs()).Parse(reportTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	r.Template = tmpl

	return r, nil
}

// GenerateReport creates an HTML report
func (r *Reporter) GenerateReport(report *TestReport) error {
	// Calculate summary stats
	r.calculateStats(report)

	// Generate main report HTML
	reportPath := filepath.Join(r.OutputDir, "index.html")
	f, err := os.Create(reportPath)
	if err != nil {
		return fmt.Errorf("create report file: %w", err)
	}
	defer f.Close()

	if err := r.Template.Execute(f, report); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	// Generate JSON summary
	jsonPath := filepath.Join(r.OutputDir, "results.json")
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	// Copy CSS/JS assets
	if err := r.copyAssets(); err != nil {
		return fmt.Errorf("copy assets: %w", err)
	}

	fmt.Printf("Report generated: %s\n", reportPath)
	return nil
}

func (r *Reporter) calculateStats(report *TestReport) {
	if len(report.Results) == 0 {
		return
	}

	var totalDiff float64
	maxDiff := 0.0
	minDiff := 100.0

	for _, result := range report.Results {
		if result.Status != TestStatusPass && result.Status != TestStatusFail {
			continue
		}

		for _, diff := range result.DiffResults {
			totalDiff += diff.DiffPercent
			if diff.DiffPercent > maxDiff {
				maxDiff = diff.DiffPercent
			}
			if diff.DiffPercent < minDiff {
				minDiff = diff.DiffPercent
			}
		}
	}

	count := 0
	for _, r := range report.Results {
		count += len(r.DiffResults)
	}

	if count > 0 {
		report.AvgDiffPercent = totalDiff / float64(count)
	}
	report.MaxDiffPercent = maxDiff
	report.MinDiffPercent = minDiff
}

func (r *Reporter) copyAssets() error {
	// Copy embedded CSS and JS files
	assets := map[string]string{
		"style.css":  styleCSS,
		"script.js":  scriptJS,
	}

	for filename, content := range assets {
		path := filepath.Join(r.AssetsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("write %s: %w", filename, err)
		}
	}

	return nil
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"statusClass": func(status TestStatus) string {
			switch status {
			case TestStatusPass:
				return "pass"
			case TestStatusFail:
				return "fail"
			case TestStatusError:
				return "error"
			default:
				return "skipped"
			}
		},
		"formatDuration": func(d time.Duration) string {
			return fmt.Sprintf("%.2fs", d.Seconds())
		},
		"formatPercent": func(p float64) string {
			return fmt.Sprintf("%.2f%%", p)
		},
		"relPath": func(fullPath, baseDir string) string {
			rel, err := filepath.Rel(baseDir, fullPath)
			if err != nil {
				return fullPath
			}
			return rel
		},
	}
}

// HTML Template (embedded)
const reportTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>E2E Visual Testing Report</title>
    <link rel="stylesheet" href="assets/style.css">
</head>
<body>
    <header>
        <h1>E2E Visual Testing Report</h1>
        <p class="timestamp">Generated: {{ .Generated.Format "2006-01-02 15:04:05" }}</p>
    </header>

    <section class="summary">
        <h2>Summary</h2>
        <div class="stats">
            <div class="stat">
                <div class="stat-value">{{ .TotalTests }}</div>
                <div class="stat-label">Total Tests</div>
            </div>
            <div class="stat pass">
                <div class="stat-value">{{ .PassedTests }}</div>
                <div class="stat-label">Passed</div>
            </div>
            <div class="stat fail">
                <div class="stat-value">{{ .FailedTests }}</div>
                <div class="stat-label">Failed</div>
            </div>
            <div class="stat">
                <div class="stat-value">{{ formatPercent .AvgDiffPercent }}</div>
                <div class="stat-label">Avg Diff</div>
            </div>
            <div class="stat">
                <div class="stat-value">{{ formatPercent .MaxDiffPercent }}</div>
                <div class="stat-label">Max Diff</div>
            </div>
        </div>
    </section>

    <section class="results">
        <h2>Test Results</h2>
        {{ range .Results }}
        <div class="test-result {{ statusClass .Status }}">
            <div class="test-header">
                <h3>{{ .TestCase.Name }} <span class="test-id">({{ .TestCase.ID }})</span></h3>
                <span class="status-badge {{ statusClass .Status }}">{{ .Status }}</span>
            </div>

            <p class="test-description">{{ .TestCase.Description }}</p>

            <div class="test-tags">
                {{ range .TestCase.Tags }}
                <span class="tag">{{ . }}</span>
                {{ end }}
            </div>

            {{ if eq .Status "pass" "fail" }}
            <div class="slides">
                {{ range $idx, $diff := .DiffResults }}
                <div class="slide-comparison">
                    <h4>Slide {{ add $idx 1 }}</h4>

                    <div class="metrics">
                        <span>Diff: {{ formatPercent $diff.DiffPercent }}</span>
                        {{ if ne $diff.SSIM 0.0 }}
                        <span>SSIM: {{ printf "%.4f" $diff.SSIM }}</span>
                        {{ end }}
                        {{ if ne $diff.PSNR 0.0 }}
                        <span>PSNR: {{ printf "%.2f" $diff.PSNR }} dB</span>
                        {{ end }}
                    </div>

                    <div class="image-grid">
                        <div class="image-col">
                            <h5>Go Output</h5>
                            <img src="{{ relPath $diff.Image1Path $.OutputDir }}" alt="Go output">
                        </div>
                        <div class="image-col">
                            <h5>C# Output</h5>
                            <img src="{{ relPath $diff.Image2Path $.OutputDir }}" alt="C# output">
                        </div>
                        <div class="image-col">
                            <h5>Diff</h5>
                            <img src="{{ relPath $diff.DiffPath $.OutputDir }}" alt="Diff overlay">
                        </div>
                    </div>
                </div>
                {{ end }}
            </div>
            {{ end }}

            {{ if eq .Status "error" }}
            <div class="error-message">
                <strong>Error:</strong> {{ .Error }}
            </div>
            {{ end }}

            <div class="test-footer">
                <span>Duration: {{ formatDuration .Duration }}</span>
            </div>
        </div>
        {{ end }}
    </section>

    <script src="assets/script.js"></script>
</body>
</html>
`

// CSS (embedded)
const styleCSS = `
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    background: #f5f5f5;
    padding: 20px;
}

header {
    background: #2c3e50;
    color: white;
    padding: 30px 20px;
    margin-bottom: 30px;
    border-radius: 8px;
}

header h1 {
    font-size: 2em;
    margin-bottom: 5px;
}

.timestamp {
    color: #bdc3c7;
    font-size: 0.9em;
}

.summary {
    background: white;
    padding: 30px;
    margin-bottom: 30px;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.stats {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    gap: 20px;
    margin-top: 20px;
}

.stat {
    text-align: center;
    padding: 20px;
    background: #ecf0f1;
    border-radius: 8px;
}

.stat.pass {
    background: #d5f4e6;
}

.stat.fail {
    background: #fadbd8;
}

.stat-value {
    font-size: 2em;
    font-weight: bold;
    color: #2c3e50;
}

.stat-label {
    color: #7f8c8d;
    font-size: 0.9em;
    margin-top: 5px;
}

.results {
    background: white;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.test-result {
    border-left: 4px solid #3498db;
    padding: 20px;
    margin-bottom: 30px;
    background: #fafafa;
    border-radius: 4px;
}

.test-result.pass {
    border-left-color: #27ae60;
}

.test-result.fail {
    border-left-color: #e74c3c;
}

.test-result.error {
    border-left-color: #f39c12;
}

.test-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
}

.test-header h3 {
    color: #2c3e50;
}

.test-id {
    font-size: 0.8em;
    color: #7f8c8d;
    font-weight: normal;
}

.status-badge {
    padding: 5px 15px;
    border-radius: 20px;
    font-size: 0.85em;
    font-weight: bold;
    text-transform: uppercase;
}

.status-badge.pass {
    background: #27ae60;
    color: white;
}

.status-badge.fail {
    background: #e74c3c;
    color: white;
}

.status-badge.error {
    background: #f39c12;
    color: white;
}

.test-description {
    color: #555;
    margin-bottom: 10px;
}

.test-tags {
    margin-bottom: 15px;
}

.tag {
    display: inline-block;
    background: #3498db;
    color: white;
    padding: 3px 10px;
    border-radius: 3px;
    font-size: 0.8em;
    margin-right: 5px;
    margin-bottom: 5px;
}

.slides {
    margin-top: 20px;
}

.slide-comparison {
    margin-bottom: 30px;
    padding: 20px;
    background: white;
    border-radius: 4px;
}

.slide-comparison h4 {
    color: #2c3e50;
    margin-bottom: 10px;
}

.metrics {
    margin-bottom: 15px;
    color: #7f8c8d;
    font-size: 0.9em;
}

.metrics span {
    margin-right: 15px;
}

.image-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 15px;
}

.image-col h5 {
    color: #555;
    margin-bottom: 10px;
    font-size: 0.9em;
}

.image-col img {
    width: 100%;
    border: 1px solid #ddd;
    border-radius: 4px;
}

.error-message {
    background: #fadbd8;
    color: #c0392b;
    padding: 15px;
    border-radius: 4px;
    margin-top: 15px;
}

.test-footer {
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid #ddd;
    color: #7f8c8d;
    font-size: 0.9em;
}

@media (max-width: 768px) {
    .image-grid {
        grid-template-columns: 1fr;
    }
}
`

// JavaScript (embedded)
const scriptJS = `
// Add image zoom on click
document.addEventListener('DOMContentLoaded', function() {
    const images = document.querySelectorAll('.image-col img');

    images.forEach(img => {
        img.addEventListener('click', function() {
            const overlay = document.createElement('div');
            overlay.style.cssText = 'position:fixed;top:0;left:0;width:100%;height:100%;background:rgba(0,0,0,0.9);z-index:9999;display:flex;align-items:center;justify-content:center;cursor:zoom-out;';

            const zoomedImg = document.createElement('img');
            zoomedImg.src = this.src;
            zoomedImg.style.cssText = 'max-width:90%;max-height:90%;';

            overlay.appendChild(zoomedImg);
            document.body.appendChild(overlay);

            overlay.addEventListener('click', function() {
                document.body.removeChild(overlay);
            });
        });
    });
});
`
```

This design document continues with additional sections covering **Nix Integration**, **Visual Comparison Algorithms**, **CI/CD Integration**, **Migration & Rollout**, **Trade-offs & Alternatives**, and **Open Questions**. Due to character limits, I'll create the rest in the next parts. Should I continue with the remaining sections?

Let me know if you'd like me to complete the design.md with the remaining sections, then move on to creating the spec.md and tasks.md files!
