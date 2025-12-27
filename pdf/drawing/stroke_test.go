package drawing

import (
	"strings"
	"testing"
)

func TestLineCap_String(t *testing.T) {
	tests := []struct {
		cap      LineCap
		expected string
	}{
		{LineCapButt, "butt"},
		{LineCapRound, "round"},
		{LineCapSquare, "square"},
		{LineCap(99), "LineCap(99)"},
	}

	for _, tt := range tests {
		result := tt.cap.String()
		if result != tt.expected {
			t.Errorf(
				"LineCap(%d).String(): expected %q, got %q",
				tt.cap,
				tt.expected,
				result,
			)
		}
	}
}

func TestLineJoin_String(t *testing.T) {
	tests := []struct {
		join     LineJoin
		expected string
	}{
		{LineJoinMiter, "miter"},
		{LineJoinRound, "round"},
		{LineJoinBevel, "bevel"},
		{LineJoin(99), "LineJoin(99)"},
	}

	for _, tt := range tests {
		result := tt.join.String()
		if result != tt.expected {
			t.Errorf(
				"LineJoin(%d).String(): expected %q, got %q",
				tt.join,
				tt.expected,
				result,
			)
		}
	}
}

func TestDashPattern_NewDashPattern(
	t *testing.T,
) {
	dp := NewDashPattern([]float64{3, 2, 1}, 0.5)

	if len(dp.Array) != 3 {
		t.Errorf(
			"Expected array length 3, got %d",
			len(dp.Array),
		)
	}
	if dp.Phase != 0.5 {
		t.Errorf(
			"Expected phase 0.5, got %v",
			dp.Phase,
		)
	}
}

func TestDashPattern_NewDashPatternSimple(
	t *testing.T,
) {
	dp := NewDashPatternSimple(4, 2)

	if len(dp.Array) != 2 {
		t.Errorf(
			"Expected array length 2, got %d",
			len(dp.Array),
		)
	}
	if dp.Array[0] != 4 || dp.Array[1] != 2 {
		t.Errorf(
			"Expected [4, 2], got %v",
			dp.Array,
		)
	}
	if dp.Phase != 0 {
		t.Errorf(
			"Expected phase 0, got %v",
			dp.Phase,
		)
	}
}

func TestDashPattern_IsSolid(t *testing.T) {
	solid := DashPattern{}
	if !solid.IsSolid() {
		t.Error(
			"Empty dash pattern should be solid",
		)
	}

	dashed := DashPattern{Array: []float64{3, 2}}
	if dashed.IsSolid() {
		t.Error(
			"Dash pattern with array should not be solid",
		)
	}
}

func TestDashPattern_ContentStream(t *testing.T) {
	tests := []struct {
		name     string
		dash     DashPattern
		expected string
	}{
		{
			name:     "solid",
			dash:     DashPattern{},
			expected: "[] 0 d",
		},
		{
			name: "simple dash",
			dash: DashPattern{
				Array: []float64{3, 2},
				Phase: 0,
			},
			expected: "[3 2] 0 d",
		},
		{
			name: "dash with phase",
			dash: DashPattern{
				Array: []float64{4, 2},
				Phase: 1,
			},
			expected: "[4 2] 1 d",
		},
		{
			name: "complex dash",
			dash: DashPattern{
				Array: []float64{4, 2, 1, 2},
				Phase: 0.5,
			},
			expected: "[4 2 1 2] 0.5 d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dash.ContentStream()
			if result != tt.expected {
				t.Errorf(
					"ContentStream: expected %q, got %q",
					tt.expected,
					result,
				)
			}
		})
	}
}

func TestDashPattern_ScaledDashPattern(
	t *testing.T,
) {
	original := DashPattern{
		Array: []float64{4, 2},
		Phase: 1,
	}
	scaled := original.ScaledDashPattern(2.0)

	if len(scaled.Array) != 2 {
		t.Errorf(
			"Expected array length 2, got %d",
			len(scaled.Array),
		)
	}
	if scaled.Array[0] != 8 ||
		scaled.Array[1] != 4 {
		t.Errorf(
			"Expected [8, 4], got %v",
			scaled.Array,
		)
	}
	if scaled.Phase != 2 {
		t.Errorf(
			"Expected phase 2, got %v",
			scaled.Phase,
		)
	}

	// Solid pattern should remain unchanged
	solid := DashSolid.ScaledDashPattern(2.0)
	if !solid.IsSolid() {
		t.Error(
			"Scaled solid pattern should still be solid",
		)
	}

	// Scale of 1.0 should return same values
	unchanged := original.ScaledDashPattern(1.0)
	if unchanged.Array[0] != 4 ||
		unchanged.Array[1] != 2 {
		t.Errorf(
			"Scale 1.0 should not change values, got %v",
			unchanged.Array,
		)
	}
}

func TestDashPatternFromPreset(t *testing.T) {
	tests := []struct {
		preset string
		dash   DashPattern
	}{
		{"solid", DashSolid},
		{"dot", DashDot},
		{"dash", DashDash},
		{"lgDash", DashLongDash},
		{"dashDot", DashDashDot},
		{"lgDashDot", DashLongDashDot},
		{"lgDashDotDot", DashLongDashDotDot},
		{"sysDot", DashSystemDot},
		{"sysDash", DashSystemDash},
		{"sysDashDot", DashSystemDashDot},
		{"sysDashDotDot", DashSystemDashDotDot},
		{
			"unknown",
			DashSolid,
		}, // Unknown preset returns solid
	}

	for _, tt := range tests {
		t.Run(tt.preset, func(t *testing.T) {
			result := DashPatternFromPreset(
				tt.preset,
			)
			if len(
				result.Array,
			) != len(
				tt.dash.Array,
			) {
				t.Errorf(
					"DashPatternFromPreset(%q): expected array length %d, got %d",
					tt.preset,
					len(tt.dash.Array),
					len(result.Array),
				)
			}
			for i := range result.Array {
				if result.Array[i] != tt.dash.Array[i] {
					t.Errorf(
						"DashPatternFromPreset(%q): array mismatch at index %d",
						tt.preset,
						i,
					)
				}
			}
		})
	}
}

func TestStrokeStyle_New(t *testing.T) {
	style := NewStrokeStyle()

	if style.Color != Black {
		t.Errorf(
			"Default color should be black, got %v",
			style.Color,
		)
	}
	if style.Width != DefaultLineWidth {
		t.Errorf(
			"Default width should be %v, got %v",
			DefaultLineWidth,
			style.Width,
		)
	}
	if style.Cap != LineCapButt {
		t.Errorf(
			"Default cap should be butt, got %v",
			style.Cap,
		)
	}
	if style.Join != LineJoinMiter {
		t.Errorf(
			"Default join should be miter, got %v",
			style.Join,
		)
	}
	if style.MiterLimit != DefaultMiterLimit {
		t.Errorf(
			"Default miter limit should be %v, got %v",
			DefaultMiterLimit,
			style.MiterLimit,
		)
	}
	if !style.Dash.IsSolid() {
		t.Error("Default dash should be solid")
	}
}

func TestStrokeStyle_NewWithColor(t *testing.T) {
	style := NewStrokeStyleWithColor(Red)

	if style.Color != Red {
		t.Errorf(
			"Color should be red, got %v",
			style.Color,
		)
	}
	if style.Width != DefaultLineWidth {
		t.Errorf(
			"Width should still be default, got %v",
			style.Width,
		)
	}
}

func TestStrokeStyle_NewWithWidth(t *testing.T) {
	style := NewStrokeStyleWithWidth(2.5)

	if style.Width != 2.5 {
		t.Errorf(
			"Width should be 2.5, got %v",
			style.Width,
		)
	}
	if style.Color != Black {
		t.Errorf(
			"Color should still be black, got %v",
			style.Color,
		)
	}
}

func TestStrokeStyle_Setters(t *testing.T) {
	style := NewStrokeStyle().
		SetColor(Blue).
		SetWidth(3.0).
		SetCap(LineCapRound).
		SetJoin(LineJoinBevel).
		SetMiterLimit(5.0).
		SetDash(DashDash)

	if style.Color != Blue {
		t.Errorf(
			"Color should be blue, got %v",
			style.Color,
		)
	}
	if style.Width != 3.0 {
		t.Errorf(
			"Width should be 3.0, got %v",
			style.Width,
		)
	}
	if style.Cap != LineCapRound {
		t.Errorf(
			"Cap should be round, got %v",
			style.Cap,
		)
	}
	if style.Join != LineJoinBevel {
		t.Errorf(
			"Join should be bevel, got %v",
			style.Join,
		)
	}
	if style.MiterLimit != 5.0 {
		t.Errorf(
			"MiterLimit should be 5.0, got %v",
			style.MiterLimit,
		)
	}
	if style.Dash.IsSolid() {
		t.Error("Dash should not be solid")
	}
}

func TestStrokeStyle_SetDashFromPreset(
	t *testing.T,
) {
	style := NewStrokeStyle().SetDashFromPreset("dashDot")

	if len(style.Dash.Array) != 4 {
		t.Errorf(
			"dashDot should have 4 elements, got %d",
			len(style.Dash.Array),
		)
	}
}

func TestStrokeStyle_IsDefault(t *testing.T) {
	defaultStyle := NewStrokeStyle()
	if !defaultStyle.IsDefault() {
		t.Error(
			"NewStrokeStyle should return default values",
		)
	}

	customWidth := NewStrokeStyle().SetWidth(2.0)
	if customWidth.IsDefault() {
		t.Error(
			"Custom width should not be default",
		)
	}

	customCap := NewStrokeStyle().SetCap(LineCapRound)
	if customCap.IsDefault() {
		t.Error(
			"Custom cap should not be default",
		)
	}

	customDash := NewStrokeStyle().SetDash(DashDash)
	if customDash.IsDefault() {
		t.Error(
			"Custom dash should not be default",
		)
	}
}

func TestStrokeStyle_ContentStream(t *testing.T) {
	style := NewStrokeStyle().
		SetColor(Red).
		SetWidth(2.0).
		SetCap(LineCapRound).
		SetJoin(LineJoinBevel).
		SetDash(DashDash)

	result := style.ContentStream()

	// Should contain line width operator
	if !strings.Contains(result, "2 w") {
		t.Error(
			"ContentStream should contain line width operator '2 w'",
		)
	}

	// Should contain line cap operator
	if !strings.Contains(result, "1 J") {
		t.Error(
			"ContentStream should contain line cap operator '1 J'",
		)
	}

	// Should contain line join operator
	if !strings.Contains(result, "2 j") {
		t.Error(
			"ContentStream should contain line join operator '2 j'",
		)
	}

	// Should contain dash pattern
	if !strings.Contains(result, "d") {
		t.Error(
			"ContentStream should contain dash operator 'd'",
		)
	}

	// Should contain stroke color (RG operator)
	if !strings.Contains(result, "RG") {
		t.Error(
			"ContentStream should contain stroke color operator 'RG'",
		)
	}
}

func TestStrokeStyle_ContentStreamMiterLimit(
	t *testing.T,
) {
	// Miter limit should only appear when join is miter and limit is non-default
	style := NewStrokeStyle().
		SetJoin(LineJoinMiter).
		SetMiterLimit(5.0)

	result := style.ContentStream()
	if !strings.Contains(result, "5 M") {
		t.Error(
			"ContentStream should contain miter limit operator '5 M'",
		)
	}

	// With bevel join, miter limit should not appear
	style2 := NewStrokeStyle().
		SetJoin(LineJoinBevel).
		SetMiterLimit(5.0)

	result2 := style2.ContentStream()
	if strings.Contains(result2, " M") {
		t.Error(
			"ContentStream with bevel join should not contain miter limit operator",
		)
	}

	// With default miter limit, should not appear
	style3 := NewStrokeStyle().SetMiterLimit(DefaultMiterLimit)
	result3 := style3.ContentStream()
	if strings.Contains(result3, "10 M") {
		t.Error(
			"ContentStream with default miter limit should not output it",
		)
	}
}

func TestStrokeStyle_ContentStreamMinimal(
	t *testing.T,
) {
	// Default style should only output color
	defaultStyle := NewStrokeStyle()
	result := defaultStyle.ContentStreamMinimal()

	// Should only have the color operator
	if !strings.Contains(result, "RG") {
		t.Error(
			"ContentStreamMinimal should contain stroke color",
		)
	}

	// Should NOT have default operators
	if strings.Contains(result, "1 w") {
		t.Error(
			"ContentStreamMinimal should not output default line width",
		)
	}
	if strings.Contains(result, "0 J") {
		t.Error(
			"ContentStreamMinimal should not output default line cap",
		)
	}
	if strings.Contains(result, "0 j") {
		t.Error(
			"ContentStreamMinimal should not output default line join",
		)
	}

	// Custom style should output non-default values
	customStyle := NewStrokeStyle().
		SetWidth(2.5).
		SetCap(LineCapSquare)

	result2 := customStyle.ContentStreamMinimal()
	if !strings.Contains(result2, "2.5 w") {
		t.Error(
			"ContentStreamMinimal should output custom line width",
		)
	}
	if !strings.Contains(result2, "2 J") {
		t.Error(
			"ContentStreamMinimal should output custom line cap",
		)
	}
}

func TestStrokeStyle_Clone(t *testing.T) {
	original := NewStrokeStyle().
		SetColor(Blue).
		SetWidth(2.0).
		SetDash(DashDash)

	cloned := original.Clone()

	// Modify clone
	cloned.Color = Green
	cloned.Width = 5.0
	cloned.Dash.Array[0] = 99

	// Original should be unchanged
	if original.Color != Blue {
		t.Error(
			"Original color should be unchanged after clone modification",
		)
	}
	if original.Width != 2.0 {
		t.Error(
			"Original width should be unchanged after clone modification",
		)
	}
	if original.Dash.Array[0] != 4 {
		t.Error(
			"Original dash array should be unchanged after clone modification",
		)
	}
}

func TestStrokeStyle_String(t *testing.T) {
	style := NewStrokeStyle().
		SetColor(Red).
		SetWidth(2.0).
		SetCap(LineCapRound).
		SetDash(DashDash)

	result := style.String()

	if !strings.Contains(result, "StrokeStyle") {
		t.Error(
			"String should contain 'StrokeStyle'",
		)
	}
	if !strings.Contains(result, "width:2.00") {
		t.Error("String should contain width")
	}
	if !strings.Contains(result, "cap:round") {
		t.Error("String should contain cap")
	}
	if !strings.Contains(result, "dash") {
		t.Error("String should contain dash")
	}
}

func TestSetLineWidth(t *testing.T) {
	tests := []struct {
		width    float64
		expected string
	}{
		{1, "1 w"},
		{0.5, "0.5 w"},
		{2.5, "2.5 w"},
		{10, "10 w"},
	}

	for _, tt := range tests {
		result := SetLineWidth(tt.width)
		if result != tt.expected {
			t.Errorf(
				"SetLineWidth(%v): expected %q, got %q",
				tt.width,
				tt.expected,
				result,
			)
		}
	}
}

func TestSetLineCap(t *testing.T) {
	tests := []struct {
		cap      LineCap
		expected string
	}{
		{LineCapButt, "0 J"},
		{LineCapRound, "1 J"},
		{LineCapSquare, "2 J"},
	}

	for _, tt := range tests {
		result := SetLineCap(tt.cap)
		if result != tt.expected {
			t.Errorf(
				"SetLineCap(%v): expected %q, got %q",
				tt.cap,
				tt.expected,
				result,
			)
		}
	}
}

func TestSetLineJoin(t *testing.T) {
	tests := []struct {
		join     LineJoin
		expected string
	}{
		{LineJoinMiter, "0 j"},
		{LineJoinRound, "1 j"},
		{LineJoinBevel, "2 j"},
	}

	for _, tt := range tests {
		result := SetLineJoin(tt.join)
		if result != tt.expected {
			t.Errorf(
				"SetLineJoin(%v): expected %q, got %q",
				tt.join,
				tt.expected,
				result,
			)
		}
	}
}

func TestSetMiterLimit(t *testing.T) {
	tests := []struct {
		limit    float64
		expected string
	}{
		{10, "10 M"},
		{5, "5 M"},
		{1.5, "1.5 M"},
	}

	for _, tt := range tests {
		result := SetMiterLimit(tt.limit)
		if result != tt.expected {
			t.Errorf(
				"SetMiterLimit(%v): expected %q, got %q",
				tt.limit,
				tt.expected,
				result,
			)
		}
	}
}

func TestSetDashPattern(t *testing.T) {
	result := SetDashPattern([]float64{4, 2}, 1)
	if result != "[4 2] 1 d" {
		t.Errorf(
			"SetDashPattern: expected '[4 2] 1 d', got %q",
			result,
		)
	}
}

func TestStrokePath(t *testing.T) {
	path := NewPathBuilder().
		MoveTo(0, 0).
		LineTo(100, 100)

	style := NewStrokeStyle().
		SetColor(Red).
		SetWidth(2.0)

	result := StrokePath(path, style)

	// Should contain stroke style operators
	if !strings.Contains(result, "2 w") {
		t.Error(
			"StrokePath should contain line width",
		)
	}
	if !strings.Contains(result, "RG") {
		t.Error(
			"StrokePath should contain stroke color",
		)
	}

	// Should contain path and stroke
	if !strings.Contains(result, "m\n") {
		t.Error(
			"StrokePath should contain moveto",
		)
	}
	if !strings.Contains(result, "l\n") {
		t.Error(
			"StrokePath should contain lineto",
		)
	}
	if !strings.Contains(result, "S\n") {
		t.Error(
			"StrokePath should contain stroke operator",
		)
	}
}

func TestStrokePath_NilInput(t *testing.T) {
	result := StrokePath(nil, NewStrokeStyle())
	if result != "" {
		t.Error(
			"StrokePath with nil path should return empty string",
		)
	}

	path := NewPathBuilder().MoveTo(0, 0)
	result = StrokePath(path, nil)
	if result != "" {
		t.Error(
			"StrokePath with nil style should return empty string",
		)
	}
}

func TestStrokePathMinimal(t *testing.T) {
	path := NewPathBuilder().
		MoveTo(0, 0).
		LineTo(100, 100)

	// Default style - should have minimal operators
	style := NewStrokeStyle()
	result := StrokePathMinimal(path, style)

	// Should NOT have default line width
	if strings.Contains(result, "1 w\n") {
		t.Error(
			"StrokePathMinimal should not output default line width",
		)
	}

	// Should have color and stroke
	if !strings.Contains(result, "RG") {
		t.Error(
			"StrokePathMinimal should contain stroke color",
		)
	}
	if !strings.Contains(result, "S\n") {
		t.Error(
			"StrokePathMinimal should contain stroke operator",
		)
	}
}

func TestNoStroke_ContentStream(t *testing.T) {
	ns := NoStroke{}
	result := ns.ContentStream()
	if result != "" {
		t.Errorf(
			"NoStroke.ContentStream should return empty string, got %q",
			result,
		)
	}
}

func TestPresetStrokeStyles(t *testing.T) {
	// Test that preset stroke styles are valid
	presets := []*StrokeStyle{
		StrokeBlackThin,
		StrokeBlackMedium,
		StrokeBlackThick,
		StrokeBlackDashed,
		StrokeBlackDotted,
		StrokeBlackRounded,
	}

	for _, preset := range presets {
		if preset.Color != Black {
			t.Errorf(
				"Preset stroke color should be black, got %v",
				preset.Color,
			)
		}

		// Should produce valid content stream
		result := preset.ContentStream()
		if result == "" {
			t.Error(
				"Preset stroke should produce non-empty content stream",
			)
		}
		if !strings.Contains(result, "w") {
			t.Error(
				"Preset stroke should contain line width operator",
			)
		}
	}

	// Test specific preset properties
	if StrokeBlackThin.Width != 0.5 {
		t.Errorf(
			"StrokeBlackThin width should be 0.5, got %v",
			StrokeBlackThin.Width,
		)
	}
	if StrokeBlackThick.Width != 2.0 {
		t.Errorf(
			"StrokeBlackThick width should be 2.0, got %v",
			StrokeBlackThick.Width,
		)
	}
	if StrokeBlackDashed.Dash.IsSolid() {
		t.Error(
			"StrokeBlackDashed should have a dash pattern",
		)
	}
	if StrokeBlackDotted.Cap != LineCapRound {
		t.Error(
			"StrokeBlackDotted should have round caps",
		)
	}
	if StrokeBlackRounded.Join != LineJoinRound {
		t.Error(
			"StrokeBlackRounded should have round joins",
		)
	}
}

func TestPresetDashPatterns(t *testing.T) {
	presets := []struct {
		name string
		dash DashPattern
	}{
		{"DashSolid", DashSolid},
		{"DashDot", DashDot},
		{"DashDash", DashDash},
		{"DashLongDash", DashLongDash},
		{"DashDashDot", DashDashDot},
		{"DashLongDashDot", DashLongDashDot},
		{
			"DashLongDashDotDot",
			DashLongDashDotDot,
		},
		{"DashSystemDot", DashSystemDot},
		{"DashSystemDash", DashSystemDash},
		{"DashSystemDashDot", DashSystemDashDot},
		{
			"DashSystemDashDotDot",
			DashSystemDashDotDot,
		},
	}

	for _, tt := range presets {
		t.Run(tt.name, func(t *testing.T) {
			// All presets should produce valid content stream
			result := tt.dash.ContentStream()
			if !strings.Contains(result, "d") {
				t.Errorf(
					"%s should produce dash operator, got %q",
					tt.name,
					result,
				)
			}

			// Only DashSolid should be solid
			if tt.name == "DashSolid" {
				if !tt.dash.IsSolid() {
					t.Error(
						"DashSolid should be solid",
					)
				}
			} else {
				if tt.dash.IsSolid() {
					t.Errorf("%s should not be solid", tt.name)
				}
			}
		})
	}
}

func TestLineCap_Values(t *testing.T) {
	// Verify PDF-standard values
	if LineCapButt != 0 {
		t.Errorf(
			"LineCapButt should be 0, got %d",
			LineCapButt,
		)
	}
	if LineCapRound != 1 {
		t.Errorf(
			"LineCapRound should be 1, got %d",
			LineCapRound,
		)
	}
	if LineCapSquare != 2 {
		t.Errorf(
			"LineCapSquare should be 2, got %d",
			LineCapSquare,
		)
	}
}

func TestLineJoin_Values(t *testing.T) {
	// Verify PDF-standard values
	if LineJoinMiter != 0 {
		t.Errorf(
			"LineJoinMiter should be 0, got %d",
			LineJoinMiter,
		)
	}
	if LineJoinRound != 1 {
		t.Errorf(
			"LineJoinRound should be 1, got %d",
			LineJoinRound,
		)
	}
	if LineJoinBevel != 2 {
		t.Errorf(
			"LineJoinBevel should be 2, got %d",
			LineJoinBevel,
		)
	}
}

func TestStrokeStyle_FluentAPI(t *testing.T) {
	// Test that fluent API returns the same pointer
	style := NewStrokeStyle()
	result := style.
		SetColor(Red).
		SetWidth(2.0).
		SetCap(LineCapRound).
		SetJoin(LineJoinBevel).
		SetMiterLimit(5.0).
		SetDash(DashDash)

	if result != style {
		t.Error(
			"Fluent API should return the same StrokeStyle pointer",
		)
	}
}

func TestStrokeStyle_ComplexContentStream(
	t *testing.T,
) {
	// Create a complex stroke style with all options
	style := NewStrokeStyle().
		SetColor(NewRGB(0.5, 0.25, 0.75)).
		SetWidth(1.5).
		SetCap(LineCapSquare).
		SetJoin(LineJoinMiter).
		SetMiterLimit(4.0).
		SetDash(NewDashPattern([]float64{5, 2, 1, 2}, 1))

	result := style.ContentStream()

	// Verify all components are present
	if !strings.Contains(result, "1.5 w") {
		t.Error(
			"Should contain line width '1.5 w'",
		)
	}
	if !strings.Contains(result, "2 J") {
		t.Error("Should contain line cap '2 J'")
	}
	if !strings.Contains(result, "0 j") {
		t.Error("Should contain line join '0 j'")
	}
	if !strings.Contains(result, "4 M") {
		t.Error(
			"Should contain miter limit '4 M'",
		)
	}
	if !strings.Contains(
		result,
		"[5 2 1 2] 1 d",
	) {
		t.Error(
			"Should contain dash pattern '[5 2 1 2] 1 d'",
		)
	}
	if !strings.Contains(
		result,
		"0.5 0.25 0.75 RG",
	) {
		t.Error(
			"Should contain stroke color '0.5 0.25 0.75 RG'",
		)
	}
}
