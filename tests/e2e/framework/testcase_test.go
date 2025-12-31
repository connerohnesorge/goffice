package framework

import (
	"encoding/json"
	"testing"
	"time"
)

//
// ChartSpec Tests
//

func TestChartSpec_Creation(t *testing.T) {
	tests := []struct {
		name      string
		chartSpec ChartSpec
		wantErr   bool
	}{
		{
			name: "valid bar clustered chart",
			chartSpec: ChartSpec{
				Type:  ChartTypeBarClustered,
				Title: "Sales by Region",
				Legend: &LegendSpec{
					Position:   LegendPositionBottom,
					ShowLegend: true,
				},
				Data: ChartData{
					Categories: []string{
						"Q1",
						"Q2",
						"Q3",
						"Q4",
					},
					Series: []SeriesData{
						{
							Name: "North",
							Values: []float64{
								100,
								120,
								140,
								160,
							},
							Color: "4472C4",
						},
						{
							Name: "South",
							Values: []float64{
								80,
								90,
								100,
								110,
							},
							Color: "ED7D31",
						},
					},
				},
				Style: ChartStyle{
					ColorPalette: []string{
						"4472C4",
						"ED7D31",
						"A5A5A5",
					},
					FontSize:       12,
					FontFamily:     "Calibri",
					ShowDataLabels: false,
					GridLines: GridLineSpec{
						MajorHorizontal: true,
						MinorHorizontal: false,
						MajorVertical:   false,
						MinorVertical:   false,
					},
				},
				Axes: &AxesSpec{
					Category: AxisSpec{
						Title: "Quarter",
					},
					Value: AxisSpec{
						Title:        "Sales ($)",
						NumberFormat: "#,##0",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid pie chart",
			chartSpec: ChartSpec{
				Type:  ChartTypePie,
				Title: "Market Share",
				Data: ChartData{
					Categories: []string{
						"Product A",
						"Product B",
						"Product C",
					},
					Series: []SeriesData{
						{
							Name: "Share",
							Values: []float64{
								45.5,
								30.2,
								24.3,
							},
							Color: "4472C4",
						},
					},
				},
				Style: ChartStyle{
					ColorPalette: []string{
						"4472C4",
						"ED7D31",
						"A5A5A5",
						"FFC000",
					},
					FontSize:       11,
					FontFamily:     "Arial",
					ShowDataLabels: true,
				},
			},
			wantErr: false,
		},
		{
			name: "valid line chart with markers",
			chartSpec: ChartSpec{
				Type:  ChartTypeLineMarkers,
				Title: "Temperature Trend",
				Data: ChartData{
					Categories: []string{
						"Jan",
						"Feb",
						"Mar",
						"Apr",
						"May",
					},
					Series: []SeriesData{
						{
							Name: "2023",
							Values: []float64{
								32.1,
								35.4,
								45.2,
								55.8,
								65.3,
							},
							Color:       "5B9BD5",
							MarkerStyle: "circle",
							Smooth:      false,
						},
						{
							Name: "2024",
							Values: []float64{
								30.5,
								33.9,
								44.1,
								57.2,
								67.8,
							},
							Color:       "ED7D31",
							MarkerStyle: "square",
							Smooth:      false,
						},
					},
				},
				Style: ChartStyle{
					FontSize:   10,
					FontFamily: "Calibri",
					GridLines: GridLineSpec{
						MajorHorizontal: true,
						MajorVertical:   true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid scatter chart",
			chartSpec: ChartSpec{
				Type:  ChartTypeScatter,
				Title: "Height vs Weight",
				Data: ChartData{
					Series: []SeriesData{
						{
							Name: "Sample Data",
							XValues: []float64{
								160,
								165,
								170,
								175,
								180,
							},
							Values: []float64{
								60,
								65,
								70,
								75,
								80,
							},
							Color: "70AD47",
						},
					},
				},
				Style: ChartStyle{
					FontSize:   12,
					FontFamily: "Arial",
				},
				Axes: &AxesSpec{
					Category: AxisSpec{
						Title: "Height (cm)",
					},
					Value: AxisSpec{
						Title: "Weight (kg)",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify basic field assignment
			if tt.chartSpec.Type == "" {
				t.Error(
					"ChartSpec.Type should not be empty",
				)
			}
			if tt.chartSpec.Title == "" {
				t.Error(
					"ChartSpec.Title should not be empty",
				)
			}
			if len(
				tt.chartSpec.Data.Series,
			) == 0 {
				t.Error(
					"ChartSpec.Data.Series should not be empty",
				)
			}

			// Verify series data
			for i, series := range tt.chartSpec.Data.Series {
				if series.Name == "" {
					t.Errorf(
						"Series %d: Name should not be empty",
						i,
					)
				}
				if len(series.Values) == 0 {
					t.Errorf(
						"Series %d: Values should not be empty",
						i,
					)
				}
			}

			// Verify style
			if tt.chartSpec.Style.FontSize <= 0 {
				t.Error(
					"ChartSpec.Style.FontSize should be positive",
				)
			}
			if tt.chartSpec.Style.FontFamily == "" {
				t.Error(
					"ChartSpec.Style.FontFamily should not be empty",
				)
			}
		})
	}
}

func TestChartType_Values(t *testing.T) {
	tests := []struct {
		name      string
		chartType ChartType
		valid     bool
	}{
		{
			"bar clustered",
			ChartTypeBarClustered,
			true,
		},
		{
			"bar stacked",
			ChartTypeBarStacked,
			true,
		},
		{
			"bar percent stacked",
			ChartTypeBarPercentStacked,
			true,
		},
		{"column", ChartTypeColumn, true},
		{"line", ChartTypeLine, true},
		{
			"line smooth",
			ChartTypeLineSmooth,
			true,
		},
		{
			"line markers",
			ChartTypeLineMarkers,
			true,
		},
		{"pie", ChartTypePie, true},
		{"doughnut", ChartTypeDoughnut, true},
		{"area", ChartTypeArea, true},
		{
			"area stacked",
			ChartTypeAreaStacked,
			true,
		},
		{"scatter", ChartTypeScatter, true},
		{
			"scatter smooth",
			ChartTypeScatterSmooth,
			true,
		},
		{
			"scatter line",
			ChartTypeScatterLine,
			true,
		},
		{"bubble", ChartTypeBubble, true},
		{"radar", ChartTypeRadar, true},
		{
			"radar filled",
			ChartTypeRadarFilled,
			true,
		},
		{"stock", ChartTypeStock, true},
		{"surface", ChartTypeSurface, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.chartType == "" && tt.valid {
				t.Error(
					"ChartType should not be empty for valid type",
				)
			}
			if string(tt.chartType) == "" &&
				tt.valid {
				t.Error(
					"ChartType string value should not be empty",
				)
			}
		})
	}
}

func TestLegendPosition_Values(t *testing.T) {
	tests := []struct {
		name     string
		position LegendPosition
	}{
		{"top", LegendPositionTop},
		{"bottom", LegendPositionBottom},
		{"left", LegendPositionLeft},
		{"right", LegendPositionRight},
		{"top right", LegendPositionTopRight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.position == "" {
				t.Error(
					"LegendPosition should not be empty",
				)
			}
		})
	}
}

//
// ShapeSpec Tests
//

func TestShapeSpec_Creation(t *testing.T) {
	tests := []struct {
		name      string
		shapeSpec ShapeSpec
	}{
		{
			name: "solid fill rectangle",
			shapeSpec: ShapeSpec{
				Type: ShapeTypeRectangle,
				Fill: FillSpec{
					Type:  FillTypeSolid,
					Color: "4472C4",
				},
				Stroke: StrokeSpec{
					Width:     12700, // 1pt
					Color:     "000000",
					DashStyle: DashStyleSolid,
					Cap:       CapStyleFlat,
					Join:      JoinStyleMiter,
				},
				Effects: []EffectSpec{},
			},
		},
		{
			name: "gradient fill ellipse",
			shapeSpec: ShapeSpec{
				Type: ShapeTypeEllipse,
				Fill: FillSpec{
					Type: FillTypeGradient,
					Gradient: &GradientSpec{
						Type:  GradientTypeLinear,
						Angle: 45,
						Stops: []GradientStop{
							{
								Position: 0.0,
								Color:    "4472C4",
							},
							{
								Position: 0.5,
								Color:    "ED7D31",
							},
							{
								Position: 1.0,
								Color:    "A5A5A5",
							},
						},
					},
				},
				Stroke: StrokeSpec{
					Width:     25400, // 2pt
					Color:     "ED7D31",
					DashStyle: DashStyleDash,
					Cap:       CapStyleRound,
					Join:      JoinStyleRound,
				},
			},
		},
		{
			name: "pattern fill with effects",
			shapeSpec: ShapeSpec{
				Type: ShapeTypeRoundRect,
				Fill: FillSpec{
					Type: FillTypePattern,
					Pattern: &PatternSpec{
						Type:       PatternTypeDots,
						Foreground: "4472C4",
						Background: "FFFFFF",
					},
				},
				Stroke: StrokeSpec{
					Width:     12700,
					Color:     "000000",
					DashStyle: DashStyleSolid,
					Cap:       CapStyleSquare,
					Join:      JoinStyleBevel,
				},
				Effects: []EffectSpec{
					{
						Type: EffectTypeShadow,
						Shadow: &ShadowSpec{
							Type:         ShadowTypeOuter,
							Angle:        45,
							Distance:     38100,
							BlurRadius:   50800,
							Color:        "000000",
							Transparency: 0.5,
						},
					},
					{
						Type: EffectTypeGlow,
						Glow: &GlowSpec{
							Radius: 25400,
							Color:  "FFD700",
						},
					},
				},
			},
		},
		{
			name: "shape with text",
			shapeSpec: ShapeSpec{
				Type: ShapeTypeRectangle,
				Fill: FillSpec{
					Type:  FillTypeSolid,
					Color: "4472C4",
				},
				Stroke: StrokeSpec{
					Width:     12700,
					Color:     "FFFFFF",
					DashStyle: DashStyleSolid,
					Cap:       CapStyleFlat,
					Join:      JoinStyleMiter,
				},
				Text: &TextSpec{
					Content: "Hello World",
					DefaultRun: RunSpec{
						FontFamily: "Arial",
						FontSize:   18,
						Color:      "FFFFFF",
						Bold:       true,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify type
			if tt.shapeSpec.Type == "" {
				t.Error(
					"ShapeSpec.Type should not be empty",
				)
			}

			// Verify fill
			if tt.shapeSpec.Fill.Type == "" {
				t.Error(
					"ShapeSpec.Fill.Type should not be empty",
				)
			}

			// Verify stroke
			if tt.shapeSpec.Stroke.Width < 0 {
				t.Error(
					"ShapeSpec.Stroke.Width should not be negative",
				)
			}
			if tt.shapeSpec.Stroke.DashStyle == "" {
				t.Error(
					"ShapeSpec.Stroke.DashStyle should not be empty",
				)
			}
		})
	}
}

func TestShapeType_Values(t *testing.T) {
	types := []ShapeType{
		ShapeTypeRectangle,
		ShapeTypeRoundRect,
		ShapeTypeEllipse,
		ShapeTypeTriangle,
		ShapeTypeDiamond,
		ShapeTypePentagon,
		ShapeTypeHexagon,
		ShapeTypeOctagon,
		ShapeTypeStar5,
		ShapeTypeArrowRight,
		ShapeTypeCallout,
	}

	for _, st := range types {
		t.Run(string(st), func(t *testing.T) {
			if st == "" {
				t.Error(
					"ShapeType should not be empty",
				)
			}
		})
	}
}

func TestFillType_Values(t *testing.T) {
	types := []FillType{
		FillTypeSolid,
		FillTypeGradient,
		FillTypePattern,
		FillTypeNone,
	}

	for _, ft := range types {
		t.Run(string(ft), func(t *testing.T) {
			if ft == "" {
				t.Error(
					"FillType should not be empty",
				)
			}
		})
	}
}

func TestGradientSpec_Creation(t *testing.T) {
	tests := []struct {
		name     string
		gradient GradientSpec
	}{
		{
			name: "linear gradient",
			gradient: GradientSpec{
				Type:  GradientTypeLinear,
				Angle: 90,
				Stops: []GradientStop{
					{
						Position: 0.0,
						Color:    "FF0000",
					},
					{
						Position: 1.0,
						Color:    "00FF00",
					},
				},
			},
		},
		{
			name: "radial gradient",
			gradient: GradientSpec{
				Type: GradientTypeRadial,
				Stops: []GradientStop{
					{
						Position: 0.0,
						Color:    "FFFFFF",
					},
					{
						Position: 0.5,
						Color:    "CCCCCC",
					},
					{
						Position: 1.0,
						Color:    "000000",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.gradient.Type == "" {
				t.Error(
					"GradientSpec.Type should not be empty",
				)
			}
			if len(tt.gradient.Stops) < 2 {
				t.Error(
					"GradientSpec should have at least 2 stops",
				)
			}
			for i, stop := range tt.gradient.Stops {
				if stop.Position < 0 ||
					stop.Position > 1 {
					t.Errorf(
						"Stop %d: Position should be between 0 and 1",
						i,
					)
				}
				if stop.Color == "" {
					t.Errorf(
						"Stop %d: Color should not be empty",
						i,
					)
				}
			}
		})
	}
}

func TestEffectSpec_Creation(t *testing.T) {
	tests := []struct {
		name   string
		effect EffectSpec
	}{
		{
			name: "shadow effect",
			effect: EffectSpec{
				Type: EffectTypeShadow,
				Shadow: &ShadowSpec{
					Type:         ShadowTypeOuter,
					Angle:        45,
					Distance:     38100,
					BlurRadius:   50800,
					Color:        "000000",
					Transparency: 0.5,
				},
			},
		},
		{
			name: "glow effect",
			effect: EffectSpec{
				Type: EffectTypeGlow,
				Glow: &GlowSpec{
					Radius: 25400,
					Color:  "FFD700",
				},
			},
		},
		{
			name: "reflection effect",
			effect: EffectSpec{
				Type: EffectTypeReflection,
				Reflection: &ReflectionSpec{
					BlurRadius:    50800,
					StartOpacity:  1.0,
					EndOpacity:    0.0,
					Distance:      38100,
					Direction:     90,
					FadeDirection: 90,
					StartPosition: 0.0,
					EndPosition:   1.0,
				},
			},
		},
		{
			name: "soft edge effect",
			effect: EffectSpec{
				Type: EffectTypeSoftEdge,
				SoftEdge: &SoftEdgeSpec{
					Radius: 12700,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.effect.Type == "" {
				t.Error(
					"EffectSpec.Type should not be empty",
				)
			}
		})
	}
}

//
// TextSpec Tests
//

func TestTextSpec_Creation(t *testing.T) {
	tests := []struct {
		name     string
		textSpec TextSpec
	}{
		{
			name: "simple text",
			textSpec: TextSpec{
				Content: "Hello World",
				DefaultRun: RunSpec{
					FontFamily: "Arial",
					FontSize:   12,
					Color:      "000000",
					Bold:       false,
					Italic:     false,
					Underline:  UnderlineNone,
				},
			},
		},
		{
			name: "rich text with multiple paragraphs",
			textSpec: TextSpec{
				Paragraphs: []ParagraphSpec{
					{
						Runs: []RunSpec{
							{
								Text:       "Bold Title",
								FontFamily: "Arial",
								FontSize:   18,
								Color:      "000000",
								Bold:       true,
							},
						},
						Alignment: TextAlignCenter,
					},
					{
						Runs: []RunSpec{
							{
								Text:       "Normal text with ",
								FontFamily: "Calibri",
								FontSize:   12,
								Color:      "000000",
							},
							{
								Text:       "italic",
								FontFamily: "Calibri",
								FontSize:   12,
								Color:      "000000",
								Italic:     true,
							},
							{
								Text:       " and ",
								FontFamily: "Calibri",
								FontSize:   12,
								Color:      "000000",
							},
							{
								Text:       "underline",
								FontFamily: "Calibri",
								FontSize:   12,
								Color:      "000000",
								Underline:  UnderlineSingle,
							},
						},
						Alignment:   TextAlignLeft,
						LineSpacing: 1.5,
					},
				},
				DefaultRun: RunSpec{
					FontFamily: "Calibri",
					FontSize:   12,
					Color:      "000000",
				},
			},
		},
		{
			name: "bulleted list",
			textSpec: TextSpec{
				Paragraphs: []ParagraphSpec{
					{
						Runs: []RunSpec{
							{
								Text:       "First bullet point",
								FontFamily: "Arial",
								FontSize:   12,
								Color:      "000000",
							},
						},
						BulletChar: "•",
						Alignment:  TextAlignLeft,
					},
					{
						Runs: []RunSpec{
							{
								Text:       "Second bullet point",
								FontFamily: "Arial",
								FontSize:   12,
								Color:      "000000",
							},
						},
						BulletChar: "•",
						Alignment:  TextAlignLeft,
					},
				},
				DefaultRun: RunSpec{
					FontFamily: "Arial",
					FontSize:   12,
					Color:      "000000",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify default run
			if tt.textSpec.DefaultRun.FontFamily == "" {
				t.Error(
					"TextSpec.DefaultRun.FontFamily should not be empty",
				)
			}
			if tt.textSpec.DefaultRun.FontSize <= 0 {
				t.Error(
					"TextSpec.DefaultRun.FontSize should be positive",
				)
			}

			// Verify paragraphs if present
			for i, para := range tt.textSpec.Paragraphs {
				if len(para.Runs) == 0 {
					t.Errorf(
						"Paragraph %d should have at least one run",
						i,
					)
				}
				for j, run := range para.Runs {
					if run.Text == "" {
						t.Errorf(
							"Paragraph %d, Run %d: Text should not be empty",
							i,
							j,
						)
					}
				}
			}
		})
	}
}

func TestTextAlignment_Values(t *testing.T) {
	alignments := []TextAlignment{
		TextAlignLeft,
		TextAlignCenter,
		TextAlignRight,
		TextAlignJustify,
	}

	for _, align := range alignments {
		t.Run(string(align), func(t *testing.T) {
			if align == "" {
				t.Error(
					"TextAlignment should not be empty",
				)
			}
		})
	}
}

func TestUnderlineType_Values(t *testing.T) {
	types := []UnderlineType{
		UnderlineNone,
		UnderlineSingle,
		UnderlineDouble,
		UnderlineDotted,
		UnderlineDash,
	}

	for _, ut := range types {
		t.Run(string(ut), func(t *testing.T) {
			if ut == "" {
				t.Error(
					"UnderlineType should not be empty",
				)
			}
		})
	}
}

//
// SlideSpec Tests
//

func TestSlideSpec_Creation(t *testing.T) {
	tests := []struct {
		name      string
		slideSpec SlideSpec
	}{
		{
			name: "blank slide",
			slideSpec: SlideSpec{
				Index:  0,
				Layout: "Blank",
				Elements: []ElementSpec{
					{
						Type: ElementTypeShape,
						Position: Position{
							X: 914400, // 1"
							Y: 914400, // 1"
						},
						Size: Size{
							Width:  4572000, // 5"
							Height: 2286000, // 2.5"
						},
						ZIndex: 1,
						Shape: &ShapeSpec{
							Type: ShapeTypeRectangle,
							Fill: FillSpec{
								Type:  FillTypeSolid,
								Color: "4472C4",
							},
							Stroke: StrokeSpec{
								Width:     12700,
								Color:     "000000",
								DashStyle: DashStyleSolid,
								Cap:       CapStyleFlat,
								Join:      JoinStyleMiter,
							},
						},
					},
				},
			},
		},
		{
			name: "slide with chart",
			slideSpec: SlideSpec{
				Index:  1,
				Layout: "Title Only",
				Background: &BackgroundSpec{
					Fill: FillSpec{
						Type:  FillTypeSolid,
						Color: "FFFFFF",
					},
				},
				Elements: []ElementSpec{
					{
						Type: ElementTypeChart,
						Position: Position{
							X: 914400,
							Y: 1828800,
						},
						Size: Size{
							Width:  7315200,
							Height: 4572000,
						},
						ZIndex: 1,
						Chart: &ChartSpec{
							Type:  ChartTypeBarClustered,
							Title: "Sales Data",
							Data: ChartData{
								Categories: []string{
									"Q1",
									"Q2",
									"Q3",
								},
								Series: []SeriesData{
									{
										Name: "Revenue",
										Values: []float64{
											100,
											150,
											200,
										},
										Color: "4472C4",
									},
								},
							},
							Style: ChartStyle{
								FontSize:   12,
								FontFamily: "Calibri",
							},
						},
					},
				},
			},
		},
		{
			name: "slide with text",
			slideSpec: SlideSpec{
				Index:  2,
				Layout: "Blank",
				Elements: []ElementSpec{
					{
						Type: ElementTypeText,
						Position: Position{
							X: 914400,
							Y: 914400,
						},
						Size: Size{
							Width:  7315200,
							Height: 914400,
						},
						ZIndex: 1,
						Text: &TextSpec{
							Content: "Title Text",
							DefaultRun: RunSpec{
								FontFamily: "Arial",
								FontSize:   24,
								Color:      "000000",
								Bold:       true,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.slideSpec.Layout == "" {
				t.Error(
					"SlideSpec.Layout should not be empty",
				)
			}
			if len(tt.slideSpec.Elements) == 0 {
				t.Error(
					"SlideSpec should have at least one element",
				)
			}

			for i, elem := range tt.slideSpec.Elements {
				if elem.Type == "" {
					t.Errorf(
						"Element %d: Type should not be empty",
						i,
					)
				}
				if elem.Size.Width <= 0 {
					t.Errorf(
						"Element %d: Width should be positive",
						i,
					)
				}
				if elem.Size.Height <= 0 {
					t.Errorf(
						"Element %d: Height should be positive",
						i,
					)
				}
			}
		})
	}
}

func TestElementType_Values(t *testing.T) {
	types := []ElementType{
		ElementTypeShape,
		ElementTypeChart,
		ElementTypeText,
		ElementTypeImage,
		ElementTypeTable,
	}

	for _, et := range types {
		t.Run(string(et), func(t *testing.T) {
			if et == "" {
				t.Error(
					"ElementType should not be empty",
				)
			}
		})
	}
}

func TestSlideSize_Presets(t *testing.T) {
	tests := []struct {
		name string
		size SlideSize
	}{
		{"4x3", SlideSize4x3},
		{"16x9", SlideSize16x9},
		{"letter", SlideSizeLetter},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.size.Width <= 0 {
				t.Error(
					"SlideSize.Width should be positive",
				)
			}
			if tt.size.Height <= 0 {
				t.Error(
					"SlideSize.Height should be positive",
				)
			}
		})
	}
}

//
// TestCase Tests
//

func TestNewTestCase(t *testing.T) {
	tc := NewTestCase(
		"test_001",
		"My Test",
		CategoryChart,
	)

	if tc.ID != "test_001" {
		t.Errorf(
			"Expected ID 'test_001', got '%s'",
			tc.ID,
		)
	}
	if tc.Name != "My Test" {
		t.Errorf(
			"Expected Name 'My Test', got '%s'",
			tc.Name,
		)
	}
	if tc.Category != CategoryChart {
		t.Errorf(
			"Expected Category 'chart', got '%s'",
			tc.Category,
		)
	}
	if tc.Created.IsZero() {
		t.Error(
			"Created timestamp should not be zero",
		)
	}
	if tc.Metadata == nil {
		t.Error("Metadata should be initialized")
	}
	if tc.Tags == nil {
		t.Error("Tags should be initialized")
	}
}

func TestTestCase_Serialization(t *testing.T) {
	// Create a test case
	tc := NewTestCase(
		"test_serialize_001",
		"Serialization Test",
		CategoryChart,
	)
	tc.Description = "Test JSON serialization"
	tc.Author = "Test Author"
	tc.Tags = []string{"test", "serialization"}
	tc.Spec = TestSpec{
		SlideCount: 1,
		SlideSize:  SlideSize16x9,
		Slides: []SlideSpec{
			{
				Index:  0,
				Layout: "Blank",
				Elements: []ElementSpec{
					{
						Type: ElementTypeChart,
						Position: Position{
							X: 914400,
							Y: 914400,
						},
						Size: Size{
							Width:  7315200,
							Height: 4572000,
						},
						ZIndex: 1,
						Chart: &ChartSpec{
							Type:  ChartTypeBarClustered,
							Title: "Test Chart",
							Data: ChartData{
								Categories: []string{
									"A",
									"B",
									"C",
								},
								Series: []SeriesData{
									{
										Name: "Series 1",
										Values: []float64{
											10,
											20,
											30,
										},
										Color: "4472C4",
									},
								},
							},
							Style: ChartStyle{
								FontSize:   12,
								FontFamily: "Arial",
							},
						},
					},
				},
			},
		},
	}

	// Serialize to JSON
	data, err := tc.ToJSON()
	if err != nil {
		t.Fatalf(
			"Failed to serialize to JSON: %v",
			err,
		)
	}

	// Deserialize from JSON
	tc2, err := FromJSON(data)
	if err != nil {
		t.Fatalf(
			"Failed to deserialize from JSON: %v",
			err,
		)
	}

	// Verify fields
	if tc2.ID != tc.ID {
		t.Errorf(
			"Expected ID '%s', got '%s'",
			tc.ID,
			tc2.ID,
		)
	}
	if tc2.Name != tc.Name {
		t.Errorf(
			"Expected Name '%s', got '%s'",
			tc.Name,
			tc2.Name,
		)
	}
	if tc2.Description != tc.Description {
		t.Errorf(
			"Expected Description '%s', got '%s'",
			tc.Description,
			tc2.Description,
		)
	}
	if tc2.Category != tc.Category {
		t.Errorf(
			"Expected Category '%s', got '%s'",
			tc.Category,
			tc2.Category,
		)
	}
	if tc2.Author != tc.Author {
		t.Errorf(
			"Expected Author '%s', got '%s'",
			tc.Author,
			tc2.Author,
		)
	}
	if len(tc2.Tags) != len(tc.Tags) {
		t.Errorf(
			"Expected %d tags, got %d",
			len(tc.Tags),
			len(tc2.Tags),
		)
	}
	if tc2.Spec.SlideCount != tc.Spec.SlideCount {
		t.Errorf(
			"Expected SlideCount %d, got %d",
			tc.Spec.SlideCount,
			tc2.Spec.SlideCount,
		)
	}
	if len(
		tc2.Spec.Slides,
	) != len(
		tc.Spec.Slides,
	) {
		t.Errorf(
			"Expected %d slides, got %d",
			len(tc.Spec.Slides),
			len(tc2.Spec.Slides),
		)
	}
}

func TestTestCase_ComplexSerialization(
	t *testing.T,
) {
	// Create a complex test case with all element types
	tc := NewTestCase(
		"test_complex_001",
		"Complex Test",
		CategoryIntegration,
	)
	tc.Description = "Test all element types"
	tc.Tags = []string{
		"integration",
		"all-elements",
	}
	tc.Spec = TestSpec{
		SlideCount: 2,
		SlideSize:  SlideSize16x9,
		Slides: []SlideSpec{
			{
				Index:  0,
				Layout: "Blank",
				Background: &BackgroundSpec{
					Fill: FillSpec{
						Type: FillTypeGradient,
						Gradient: &GradientSpec{
							Type:  GradientTypeLinear,
							Angle: 45,
							Stops: []GradientStop{
								{
									Position: 0.0,
									Color:    "FFFFFF",
								},
								{
									Position: 1.0,
									Color:    "F0F0F0",
								},
							},
						},
					},
				},
				Elements: []ElementSpec{
					{
						Type: ElementTypeShape,
						Position: Position{
							X: 914400,
							Y: 914400,
						},
						Size: Size{
							Width:  2286000,
							Height: 1143000,
						},
						ZIndex: 1,
						Shape: &ShapeSpec{
							Type: ShapeTypeRectangle,
							Fill: FillSpec{
								Type:  FillTypeSolid,
								Color: "4472C4",
							},
							Stroke: StrokeSpec{
								Width:     12700,
								Color:     "000000",
								DashStyle: DashStyleSolid,
								Cap:       CapStyleFlat,
								Join:      JoinStyleMiter,
							},
							Effects: []EffectSpec{
								{
									Type: EffectTypeShadow,
									Shadow: &ShadowSpec{
										Type:         ShadowTypeOuter,
										Angle:        45,
										Distance:     38100,
										BlurRadius:   50800,
										Color:        "000000",
										Transparency: 0.5,
									},
								},
							},
						},
					},
					{
						Type: ElementTypeText,
						Position: Position{
							X: 3657600,
							Y: 914400,
						},
						Size: Size{
							Width:  4572000,
							Height: 914400,
						},
						ZIndex: 2,
						Text: &TextSpec{
							Paragraphs: []ParagraphSpec{
								{
									Runs: []RunSpec{
										{
											Text:       "Title Text",
											FontFamily: "Arial",
											FontSize:   24,
											Color:      "000000",
											Bold:       true,
										},
									},
									Alignment: TextAlignCenter,
								},
							},
							DefaultRun: RunSpec{
								FontFamily: "Arial",
								FontSize:   12,
								Color:      "000000",
							},
						},
					},
				},
			},
			{
				Index:  1,
				Layout: "Title Only",
				Elements: []ElementSpec{
					{
						Type: ElementTypeChart,
						Position: Position{
							X: 914400,
							Y: 1828800,
						},
						Size: Size{
							Width:  7315200,
							Height: 4572000,
						},
						ZIndex: 1,
						Chart: &ChartSpec{
							Type:  ChartTypeLineMarkers,
							Title: "Trend Analysis",
							Legend: &LegendSpec{
								Position:   LegendPositionBottom,
								ShowLegend: true,
							},
							Data: ChartData{
								Categories: []string{
									"Jan",
									"Feb",
									"Mar",
									"Apr",
								},
								Series: []SeriesData{
									{
										Name: "2023",
										Values: []float64{
											100,
											120,
											140,
											160,
										},
										Color:       "4472C4",
										MarkerStyle: "circle",
									},
									{
										Name: "2024",
										Values: []float64{
											110,
											130,
											150,
											170,
										},
										Color:       "ED7D31",
										MarkerStyle: "square",
									},
								},
							},
							Style: ChartStyle{
								ColorPalette: []string{
									"4472C4",
									"ED7D31",
								},
								FontSize:       12,
								FontFamily:     "Calibri",
								ShowDataLabels: false,
								GridLines: GridLineSpec{
									MajorHorizontal: true,
									MajorVertical:   false,
								},
							},
							Axes: &AxesSpec{
								Category: AxisSpec{
									Title: "Month",
								},
								Value: AxisSpec{
									Title:        "Value",
									NumberFormat: "#,##0",
								},
							},
						},
					},
				},
			},
		},
		Theme: &ThemeSpec{
			Name: "Custom Theme",
			ColorScheme: ColorScheme{
				Accent1:           "4472C4",
				Accent2:           "ED7D31",
				Accent3:           "A5A5A5",
				Accent4:           "FFC000",
				Accent5:           "5B9BD5",
				Accent6:           "70AD47",
				Dark1:             "000000",
				Dark2:             "1F4E78",
				Light1:            "FFFFFF",
				Light2:            "E7E6E6",
				Hyperlink:         "0563C1",
				FollowedHyperlink: "954F72",
			},
			FontScheme: FontScheme{
				MajorFont: "Arial",
				MinorFont: "Calibri",
			},
		},
	}

	// Serialize
	data, err := tc.ToJSON()
	if err != nil {
		t.Fatalf(
			"Failed to serialize complex test case: %v",
			err,
		)
	}

	// Deserialize
	tc2, err := FromJSON(data)
	if err != nil {
		t.Fatalf(
			"Failed to deserialize complex test case: %v",
			err,
		)
	}

	// Verify structure
	if len(tc2.Spec.Slides) != 2 {
		t.Errorf(
			"Expected 2 slides, got %d",
			len(tc2.Spec.Slides),
		)
	}
	if tc2.Spec.Theme == nil {
		t.Error("Theme should not be nil")
	}
	if tc2.Spec.Theme.Name != "Custom Theme" {
		t.Errorf(
			"Expected theme name 'Custom Theme', got '%s'",
			tc2.Spec.Theme.Name,
		)
	}

	// Verify first slide elements
	if len(tc2.Spec.Slides[0].Elements) != 2 {
		t.Errorf(
			"Expected 2 elements on first slide, got %d",
			len(tc2.Spec.Slides[0].Elements),
		)
	}

	// Verify chart on second slide
	if len(tc2.Spec.Slides[1].Elements) != 1 {
		t.Errorf(
			"Expected 1 element on second slide, got %d",
			len(tc2.Spec.Slides[1].Elements),
		)
	}
	chartElem := tc2.Spec.Slides[1].Elements[0]
	if chartElem.Type != ElementTypeChart {
		t.Errorf(
			"Expected chart element, got %s",
			chartElem.Type,
		)
	}
	if chartElem.Chart == nil {
		t.Fatal("Chart should not be nil")
	}
	if len(chartElem.Chart.Data.Series) != 2 {
		t.Errorf(
			"Expected 2 series, got %d",
			len(chartElem.Chart.Data.Series),
		)
	}
}

func TestDefaultTestConfig(t *testing.T) {
	config := DefaultTestConfig()

	if config.RenderDPI <= 0 {
		t.Error("RenderDPI should be positive")
	}
	if config.RenderBackend == "" {
		t.Error(
			"RenderBackend should not be empty",
		)
	}
	if config.DiffThreshold < 0 ||
		config.DiffThreshold > 1 {
		t.Error(
			"DiffThreshold should be between 0 and 1",
		)
	}
	if config.DiffAlgorithm == "" {
		t.Error(
			"DiffAlgorithm should not be empty",
		)
	}
	if config.GeneratorTimeout <= 0 {
		t.Error(
			"GeneratorTimeout should be positive",
		)
	}
	if config.RenderTimeout <= 0 {
		t.Error(
			"RenderTimeout should be positive",
		)
	}
	if config.RetryCount < 0 {
		t.Error(
			"RetryCount should not be negative",
		)
	}
	if config.RetryDelay < 0 {
		t.Error(
			"RetryDelay should not be negative",
		)
	}
}

func TestTestCategory_Values(t *testing.T) {
	categories := []TestCategory{
		CategoryChart,
		CategoryShape,
		CategoryText,
		CategoryImage,
		CategoryTable,
		CategoryIntegration,
	}

	for _, cat := range categories {
		t.Run(string(cat), func(t *testing.T) {
			if cat == "" {
				t.Error(
					"TestCategory should not be empty",
				)
			}
		})
	}
}

func TestRenderBackend_Values(t *testing.T) {
	backends := []RenderBackend{
		RenderBackendLibreOffice,
		RenderBackendPowerPoint,
		RenderBackendPDF,
		RenderBackendCustom,
	}

	for _, backend := range backends {
		t.Run(
			string(backend),
			func(t *testing.T) {
				if backend == "" {
					t.Error(
						"RenderBackend should not be empty",
					)
				}
			},
		)
	}
}

func TestDiffAlgorithm_Values(t *testing.T) {
	algorithms := []DiffAlgorithm{
		DiffAlgorithmPixelPerfect,
		DiffAlgorithmPerceptual,
		DiffAlgorithmSSIM,
		DiffAlgorithmMSE,
		DiffAlgorithmPSNR,
	}

	for _, algo := range algorithms {
		t.Run(string(algo), func(t *testing.T) {
			if algo == "" {
				t.Error(
					"DiffAlgorithm should not be empty",
				)
			}
		})
	}
}

//
// Color Handling Tests
//

func TestColorHandling(t *testing.T) {
	tests := []struct {
		name  string
		color string
		valid bool
	}{
		{"6-digit hex", "4472C4", true},
		{"uppercase hex", "ED7D31", true},
		{"lowercase hex", "a5a5a5", true},
		{
			"8-digit hex with alpha",
			"80000000",
			true,
		},
		{"black", "000000", true},
		{"white", "FFFFFF", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a simple shape with this color
			spec := ShapeSpec{
				Type: ShapeTypeRectangle,
				Fill: FillSpec{
					Type:  FillTypeSolid,
					Color: tt.color,
				},
				Stroke: StrokeSpec{
					Width:     12700,
					Color:     tt.color,
					DashStyle: DashStyleSolid,
					Cap:       CapStyleFlat,
					Join:      JoinStyleMiter,
				},
			}

			// Serialize and deserialize
			data, err := json.Marshal(spec)
			if err != nil {
				t.Fatalf(
					"Failed to marshal shape spec: %v",
					err,
				)
			}

			var spec2 ShapeSpec
			err = json.Unmarshal(data, &spec2)
			if err != nil {
				t.Fatalf(
					"Failed to unmarshal shape spec: %v",
					err,
				)
			}

			if spec2.Fill.Color != tt.color {
				t.Errorf(
					"Expected fill color '%s', got '%s'",
					tt.color,
					spec2.Fill.Color,
				)
			}
			if spec2.Stroke.Color != tt.color {
				t.Errorf(
					"Expected stroke color '%s', got '%s'",
					tt.color,
					spec2.Stroke.Color,
				)
			}
		})
	}
}

//
// Configuration Options Tests
//

func TestTestConfig_CustomValues(t *testing.T) {
	config := TestConfig{
		RenderDPI:          150,
		RenderBackend:      RenderBackendPDF,
		DiffThreshold:      0.05,
		DiffAlgorithm:      DiffAlgorithmSSIM,
		IgnoreAntialiasing: false,
		GeneratorTimeout:   30,
		RenderTimeout:      60,
		GenerateDiffImage:  false,
		SaveIntermediates:  true,
		RetryCount:         5,
		RetryDelay:         10,
	}

	// Serialize
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf(
			"Failed to marshal config: %v",
			err,
		)
	}

	// Deserialize
	var config2 TestConfig
	err = json.Unmarshal(data, &config2)
	if err != nil {
		t.Fatalf(
			"Failed to unmarshal config: %v",
			err,
		)
	}

	// Verify all fields
	if config2.RenderDPI != config.RenderDPI {
		t.Errorf(
			"Expected RenderDPI %d, got %d",
			config.RenderDPI,
			config2.RenderDPI,
		)
	}
	if config2.RenderBackend != config.RenderBackend {
		t.Errorf(
			"Expected RenderBackend %s, got %s",
			config.RenderBackend,
			config2.RenderBackend,
		)
	}
	if config2.DiffThreshold != config.DiffThreshold {
		t.Errorf(
			"Expected DiffThreshold %f, got %f",
			config.DiffThreshold,
			config2.DiffThreshold,
		)
	}
	if config2.DiffAlgorithm != config.DiffAlgorithm {
		t.Errorf(
			"Expected DiffAlgorithm %s, got %s",
			config.DiffAlgorithm,
			config2.DiffAlgorithm,
		)
	}
	if config2.IgnoreAntialiasing != config.IgnoreAntialiasing {
		t.Errorf(
			"Expected IgnoreAntialiasing %v, got %v",
			config.IgnoreAntialiasing,
			config2.IgnoreAntialiasing,
		)
	}
	if config2.GeneratorTimeout != config.GeneratorTimeout {
		t.Errorf(
			"Expected GeneratorTimeout %d, got %d",
			config.GeneratorTimeout,
			config2.GeneratorTimeout,
		)
	}
	if config2.RenderTimeout != config.RenderTimeout {
		t.Errorf(
			"Expected RenderTimeout %d, got %d",
			config.RenderTimeout,
			config2.RenderTimeout,
		)
	}
	if config2.GenerateDiffImage != config.GenerateDiffImage {
		t.Errorf(
			"Expected GenerateDiffImage %v, got %v",
			config.GenerateDiffImage,
			config2.GenerateDiffImage,
		)
	}
	if config2.SaveIntermediates != config.SaveIntermediates {
		t.Errorf(
			"Expected SaveIntermediates %v, got %v",
			config.SaveIntermediates,
			config2.SaveIntermediates,
		)
	}
	if config2.RetryCount != config.RetryCount {
		t.Errorf(
			"Expected RetryCount %d, got %d",
			config.RetryCount,
			config2.RetryCount,
		)
	}
	if config2.RetryDelay != config.RetryDelay {
		t.Errorf(
			"Expected RetryDelay %d, got %d",
			config.RetryDelay,
			config2.RetryDelay,
		)
	}
}

//
// Edge Cases and Validation Tests
//

func TestTestCase_EmptySlides(t *testing.T) {
	tc := NewTestCase(
		"empty_test",
		"Empty Test",
		CategoryChart,
	)
	tc.Spec = TestSpec{
		SlideCount: 0,
		SlideSize:  SlideSize16x9,
		Slides:     []SlideSpec{},
	}

	// Should serialize without error
	data, err := tc.ToJSON()
	if err != nil {
		t.Fatalf(
			"Failed to serialize empty test: %v",
			err,
		)
	}

	// Should deserialize without error
	tc2, err := FromJSON(data)
	if err != nil {
		t.Fatalf(
			"Failed to deserialize empty test: %v",
			err,
		)
	}

	if len(tc2.Spec.Slides) != 0 {
		t.Errorf(
			"Expected 0 slides, got %d",
			len(tc2.Spec.Slides),
		)
	}
}

func TestTestCase_Metadata(t *testing.T) {
	tc := NewTestCase(
		"meta_test",
		"Metadata Test",
		CategoryChart,
	)
	tc.Metadata["key1"] = "value1"
	tc.Metadata["key2"] = "value2"
	tc.Metadata["version"] = "1.0"

	// Serialize
	data, err := tc.ToJSON()
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Deserialize
	tc2, err := FromJSON(data)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	if len(tc2.Metadata) != 3 {
		t.Errorf(
			"Expected 3 metadata entries, got %d",
			len(tc2.Metadata),
		)
	}
	if tc2.Metadata["key1"] != "value1" {
		t.Errorf(
			"Expected metadata key1='value1', got '%s'",
			tc2.Metadata["key1"],
		)
	}
	if tc2.Metadata["version"] != "1.0" {
		t.Errorf(
			"Expected metadata version='1.0', got '%s'",
			tc2.Metadata["version"],
		)
	}
}

func TestTestCase_TimestampPersistence(
	t *testing.T,
) {
	// Create test case with specific timestamp
	now := time.Now()
	tc := &TestCase{
		ID:       "time_test",
		Name:     "Timestamp Test",
		Category: CategoryChart,
		Created:  now,
		Config:   DefaultTestConfig(),
		Metadata: make(map[string]string),
		Tags:     []string{},
	}

	// Serialize
	data, err := tc.ToJSON()
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Deserialize
	tc2, err := FromJSON(data)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	// Compare timestamps (allowing for minor serialization differences)
	diff := tc2.Created.Sub(now)
	if diff > time.Second || diff < -time.Second {
		t.Errorf(
			"Timestamp difference too large: %v",
			diff,
		)
	}
}
