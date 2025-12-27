package drawing

import (
	"math"
	"strings"
	"testing"
)

func TestRect(t *testing.T) {
	r := NewRect(10, 20, 100, 50)
	if r.X != 10 || r.Y != 20 || r.Width != 100 ||
		r.Height != 50 {
		t.Errorf("NewRect: unexpected values")
	}
}

func TestRect_Center(t *testing.T) {
	r := NewRect(0, 0, 100, 50)
	cx, cy := r.Center()
	if cx != 50 || cy != 25 {
		t.Errorf(
			"Center: expected (50, 25), got (%v, %v)",
			cx,
			cy,
		)
	}
}

func TestRect_Right(t *testing.T) {
	r := NewRect(10, 20, 100, 50)
	if r.Right() != 110 {
		t.Errorf(
			"Right: expected 110, got %v",
			r.Right(),
		)
	}
}

func TestRect_Top(t *testing.T) {
	r := NewRect(10, 20, 100, 50)
	if r.Top() != 70 {
		t.Errorf(
			"Top: expected 70, got %v",
			r.Top(),
		)
	}
}

func TestNoFill(t *testing.T) {
	f := NoFill{}
	if f.Type() != FillTypeNone {
		t.Error(
			"NoFill should have type FillTypeNone",
		)
	}
	if !f.IsNone() {
		t.Error(
			"NoFill.IsNone should return true",
		)
	}
	if f.ContentStream(
		NewRect(0, 0, 100, 100),
	) != "" {
		t.Error(
			"NoFill.ContentStream should return empty string",
		)
	}
}

func TestSolidFill(t *testing.T) {
	f := NewSolidFill(Red)
	if f.Type() != FillTypeSolid {
		t.Error(
			"SolidFill should have type FillTypeSolid",
		)
	}
	if f.IsNone() {
		t.Error(
			"SolidFill with opaque color should not be none",
		)
	}

	content := f.ContentStream(
		NewRect(0, 0, 100, 100),
	)
	if !strings.Contains(content, "rg") {
		t.Errorf(
			"SolidFill.ContentStream should contain 'rg', got %q",
			content,
		)
	}
}

func TestSolidFill_Transparent(t *testing.T) {
	f := NewSolidFill(Transparent)
	if !f.IsNone() {
		t.Error(
			"SolidFill with transparent color should be none",
		)
	}
	if f.ContentStream(
		NewRect(0, 0, 100, 100),
	) != "" {
		t.Error(
			"SolidFill with transparent color should return empty content",
		)
	}
}

func TestGradientStop(t *testing.T) {
	stop := NewGradientStop(0.5, Red)
	if stop.Position != 0.5 {
		t.Errorf(
			"GradientStop.Position: expected 0.5, got %v",
			stop.Position,
		)
	}
	if stop.Color != Red {
		t.Error(
			"GradientStop.Color should be Red",
		)
	}
}

func TestGradientStop_Clamping(t *testing.T) {
	stop := NewGradientStop(1.5, Red)
	if stop.Position != 1 {
		t.Errorf(
			"GradientStop should clamp position to 1, got %v",
			stop.Position,
		)
	}

	stop = NewGradientStop(-0.5, Red)
	if stop.Position != 0 {
		t.Errorf(
			"GradientStop should clamp position to 0, got %v",
			stop.Position,
		)
	}
}

func TestLinearGradient(t *testing.T) {
	g := NewLinearGradient(45, Red, Blue)
	if g.Type() != FillTypeLinearGradient {
		t.Error(
			"LinearGradient should have type FillTypeLinearGradient",
		)
	}
	if g.IsNone() {
		t.Error(
			"LinearGradient with stops should not be none",
		)
	}
	if g.Angle != 45 {
		t.Errorf(
			"LinearGradient.Angle: expected 45, got %v",
			g.Angle,
		)
	}
	if len(g.Stops) != 2 {
		t.Errorf(
			"LinearGradient should have 2 stops, got %d",
			len(g.Stops),
		)
	}
}

func TestLinearGradientWithStops(t *testing.T) {
	stops := []GradientStop{
		{Position: 1, Color: Blue},
		{Position: 0, Color: Red},
		{Position: 0.5, Color: Green},
	}
	g := NewLinearGradientWithStops(90, stops)

	// Stops should be sorted
	if g.Stops[0].Position != 0 {
		t.Error(
			"Stops should be sorted by position",
		)
	}
	if g.Stops[1].Position != 0.5 {
		t.Error(
			"Stops should be sorted by position",
		)
	}
	if g.Stops[2].Position != 1 {
		t.Error(
			"Stops should be sorted by position",
		)
	}
}

func TestLinearGradient_AddStop(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)
	g.AddStop(0.5, Green)

	if len(g.Stops) != 3 {
		t.Errorf(
			"AddStop: expected 3 stops, got %d",
			len(g.Stops),
		)
	}
	// Stops should still be sorted
	if g.Stops[1].Position != 0.5 {
		t.Error(
			"AddStop: stops should remain sorted",
		)
	}
}

func TestLinearGradient_SetExtend(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)
	g.SetExtend(GradientExtendRepeat)

	if g.Extend != GradientExtendRepeat {
		t.Error(
			"SetExtend should change extend mode",
		)
	}
}

func TestLinearGradient_SetOpacity(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)
	g.SetOpacity(0.5)

	if g.Opacity != 0.5 {
		t.Errorf(
			"SetOpacity: expected 0.5, got %v",
			g.Opacity,
		)
	}
}

func TestLinearGradient_StartPoint(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	x, y := g.StartPoint(bounds)

	// For angle 0, start should be to the left of center
	cx, cy := bounds.Center()
	if x >= cx {
		t.Error(
			"StartPoint for angle 0 should be left of center",
		)
	}
	if math.Abs(y-cy) > 0.001 {
		t.Error(
			"StartPoint for angle 0 should be at vertical center",
		)
	}
}

func TestLinearGradient_EndPoint(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	x, y := g.EndPoint(bounds)

	// For angle 0, end should be to the right of center
	cx, cy := bounds.Center()
	if x <= cx {
		t.Error(
			"EndPoint for angle 0 should be right of center",
		)
	}
	if math.Abs(y-cy) > 0.001 {
		t.Error(
			"EndPoint for angle 0 should be at vertical center",
		)
	}
}

func TestLinearGradient_ColorAt(t *testing.T) {
	g := NewLinearGradient(0, Red, Blue)

	// At position 0, should be Red
	c := g.ColorAt(0)
	if c.R != 1 || c.B != 0 {
		t.Error("ColorAt(0) should be Red")
	}

	// At position 1, should be Blue
	c = g.ColorAt(1)
	if c.R != 0 || c.B != 1 {
		t.Error("ColorAt(1) should be Blue")
	}

	// At position 0.5, should be blend
	c = g.ColorAt(0.5)
	if c.R < 0.4 || c.R > 0.6 || c.B < 0.4 ||
		c.B > 0.6 {
		t.Errorf(
			"ColorAt(0.5) should be blend, got (%v, %v, %v)",
			c.R,
			c.G,
			c.B,
		)
	}
}

func TestLinearGradient_ColorAt_MultipleStops(
	t *testing.T,
) {
	g := NewLinearGradient(0, Red, Blue)
	g.AddStop(0.5, Green)

	// At 0.25, should be between Red and Green
	c := g.ColorAt(0.25)
	if c.G <= 0 {
		t.Error(
			"ColorAt(0.25) should have some green",
		)
	}

	// At 0.75, should be between Green and Blue
	c = g.ColorAt(0.75)
	if c.B <= 0 {
		t.Error(
			"ColorAt(0.75) should have some blue",
		)
	}
}

func TestLinearGradient_ContentStream(
	t *testing.T,
) {
	g := NewLinearGradient(0, Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	content := g.ContentStream(bounds)

	// Should produce some color output
	if content == "" {
		t.Error(
			"ContentStream should not be empty",
		)
	}
	if !strings.Contains(content, "rg") {
		t.Error(
			"ContentStream should contain 'rg' operator",
		)
	}
}

func TestLinearGradient_ShadingDict(
	t *testing.T,
) {
	g := NewLinearGradient(0, Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	dict := g.ShadingDict(bounds)

	if dict["ShadingType"] != 2 {
		t.Error(
			"ShadingDict should have ShadingType=2 (axial)",
		)
	}
	if dict["ColorSpace"] != "/DeviceRGB" {
		t.Error(
			"ShadingDict should use DeviceRGB",
		)
	}
	coords, ok := dict["Coords"].([]float64)
	if !ok || len(coords) != 4 {
		t.Error(
			"ShadingDict should have 4 coordinates",
		)
	}
}

func TestRadialGradient(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	if g.Type() != FillTypeRadialGradient {
		t.Error(
			"RadialGradient should have type FillTypeRadialGradient",
		)
	}
	if g.IsNone() {
		t.Error(
			"RadialGradient with stops should not be none",
		)
	}
	if g.CenterX != 0.5 || g.CenterY != 0.5 {
		t.Error(
			"RadialGradient should default to center",
		)
	}
	if g.Radius != 0.5 {
		t.Errorf(
			"RadialGradient should default radius=0.5, got %v",
			g.Radius,
		)
	}
}

func TestRadialGradientWithStops(t *testing.T) {
	stops := []GradientStop{
		{Position: 1, Color: Blue},
		{Position: 0, Color: Red},
	}
	g := NewRadialGradientWithStops(stops)

	// Stops should be sorted
	if g.Stops[0].Position != 0 ||
		g.Stops[1].Position != 1 {
		t.Error(
			"Stops should be sorted by position",
		)
	}
}

func TestRadialGradient_SetCenter(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	g.SetCenter(0.25, 0.75)

	if g.CenterX != 0.25 || g.CenterY != 0.75 {
		t.Error(
			"SetCenter should update center coordinates",
		)
	}
}

func TestRadialGradient_SetFocus(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	g.SetFocus(0.3, 0.3)

	if g.FocusX != 0.3 || g.FocusY != 0.3 {
		t.Error(
			"SetFocus should update focus coordinates",
		)
	}
}

func TestRadialGradient_SetRadius(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	g.SetRadius(0.75)

	if g.Radius != 0.75 {
		t.Errorf(
			"SetRadius: expected 0.75, got %v",
			g.Radius,
		)
	}
}

func TestRadialGradient_AddStop(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	g.AddStop(0.5, Green)

	if len(g.Stops) != 3 {
		t.Errorf(
			"AddStop: expected 3 stops, got %d",
			len(g.Stops),
		)
	}
}

func TestRadialGradient_CenterPoint(
	t *testing.T,
) {
	g := NewRadialGradient(Red, Blue)
	bounds := NewRect(0, 0, 100, 50)
	x, y := g.CenterPoint(bounds)

	if x != 50 || y != 25 {
		t.Errorf(
			"CenterPoint: expected (50, 25), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestRadialGradient_FocusPoint(t *testing.T) {
	g := NewRadialGradient(Red, Blue)
	g.SetFocus(0.25, 0.25)
	bounds := NewRect(0, 0, 100, 100)
	x, y := g.FocusPoint(bounds)

	if x != 25 || y != 25 {
		t.Errorf(
			"FocusPoint: expected (25, 25), got (%v, %v)",
			x,
			y,
		)
	}
}

func TestRadialGradient_RadiusValue(
	t *testing.T,
) {
	g := NewRadialGradient(Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	r := g.RadiusValue(bounds)

	diagonal := math.Sqrt(100*100 + 100*100)
	expected := diagonal * 0.5
	if math.Abs(r-expected) > 0.001 {
		t.Errorf(
			"RadiusValue: expected %v, got %v",
			expected,
			r,
		)
	}
}

func TestRadialGradient_ColorAt(t *testing.T) {
	g := NewRadialGradient(Red, Blue)

	c := g.ColorAt(0)
	if c.R != 1 {
		t.Error("ColorAt(0) should be Red")
	}

	c = g.ColorAt(1)
	if c.B != 1 {
		t.Error("ColorAt(1) should be Blue")
	}
}

func TestRadialGradient_ContentStream(
	t *testing.T,
) {
	g := NewRadialGradient(Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	content := g.ContentStream(bounds)

	if content == "" {
		t.Error(
			"ContentStream should not be empty",
		)
	}
}

func TestRadialGradient_ShadingDict(
	t *testing.T,
) {
	g := NewRadialGradient(Red, Blue)
	bounds := NewRect(0, 0, 100, 100)
	dict := g.ShadingDict(bounds)

	if dict["ShadingType"] != 3 {
		t.Error(
			"ShadingDict should have ShadingType=3 (radial)",
		)
	}
	coords, ok := dict["Coords"].([]float64)
	if !ok || len(coords) != 6 {
		t.Error(
			"ShadingDict should have 6 coordinates (x0, y0, r0, x1, y1, r1)",
		)
	}
}

func TestPatternFill(t *testing.T) {
	p := NewPatternFill(
		PatternTypeHorizontalLines,
		Blue,
	)
	if p.Type() != FillTypePattern {
		t.Error(
			"PatternFill should have type FillTypePattern",
		)
	}
	if p.PatternType != PatternTypeHorizontalLines {
		t.Error(
			"PatternFill should have correct pattern type",
		)
	}
	if p.Color != Blue {
		t.Error(
			"PatternFill should have correct color",
		)
	}
}

func TestPatternFill_SetBackground(t *testing.T) {
	p := NewPatternFill(PatternTypeDots, Red)
	p.SetBackground(White)

	if p.BackgroundColor != White {
		t.Error(
			"SetBackground should update background color",
		)
	}
}

func TestPatternFill_SetSpacing(t *testing.T) {
	p := NewPatternFill(
		PatternTypeVerticalLines,
		Red,
	)
	p.SetSpacing(8.0)

	if p.Spacing != 8.0 {
		t.Errorf(
			"SetSpacing: expected 8.0, got %v",
			p.Spacing,
		)
	}
}

func TestPatternFill_SetAngle(t *testing.T) {
	p := NewPatternFill(
		PatternTypeDiagonalLines,
		Red,
	)
	p.SetAngle(45)

	if p.Angle != 45 {
		t.Errorf(
			"SetAngle: expected 45, got %v",
			p.Angle,
		)
	}
}

func TestPatternFill_IsNone(t *testing.T) {
	p := NewPatternFill(
		PatternTypeSolid,
		Transparent,
	)
	p.SetBackground(Transparent)

	if !p.IsNone() {
		t.Error(
			"PatternFill with both colors transparent should be none",
		)
	}

	p.Color = Red
	if p.IsNone() {
		t.Error(
			"PatternFill with opaque color should not be none",
		)
	}
}

func TestPatternFill_ContentStream(t *testing.T) {
	p := NewPatternFill(PatternTypeSolid, Red)
	content := p.ContentStream(
		NewRect(0, 0, 100, 100),
	)

	if !strings.Contains(content, "rg") {
		t.Errorf(
			"ContentStream should contain 'rg', got %q",
			content,
		)
	}
}

func TestFillRect(t *testing.T) {
	bounds := NewRect(10, 20, 100, 50)
	fill := NewSolidFill(Red)
	content := FillRect(bounds, fill)

	if !strings.Contains(content, "rg") {
		t.Error(
			"FillRect should contain fill color",
		)
	}
	if !strings.Contains(content, "re f") {
		t.Error(
			"FillRect should contain rectangle and fill operators",
		)
	}
	if !strings.Contains(content, "10") ||
		!strings.Contains(content, "20") {
		t.Error(
			"FillRect should contain rectangle coordinates",
		)
	}
}

func TestFillRect_NoFill(t *testing.T) {
	bounds := NewRect(0, 0, 100, 100)
	content := FillRect(bounds, nil)

	if content != "" {
		t.Error(
			"FillRect with nil fill should return empty string",
		)
	}

	content = FillRect(bounds, NoFill{})
	if content != "" {
		t.Error(
			"FillRect with NoFill should return empty string",
		)
	}
}

func TestFillPath(t *testing.T) {
	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)
	fill := NewSolidFill(Blue)
	bounds := NewRect(0, 0, 100, 100)

	content := FillPath(path, fill, bounds)

	if !strings.Contains(content, "rg") {
		t.Error(
			"FillPath should contain fill color",
		)
	}
	if !strings.Contains(content, "f\n") {
		t.Error(
			"FillPath should contain fill operator",
		)
	}
}

func TestFillPath_NoFill(t *testing.T) {
	path := NewPathBuilder()
	path.Rectangle(0, 0, 100, 100)
	bounds := NewRect(0, 0, 100, 100)

	content := FillPath(path, nil, bounds)
	if content != "" {
		t.Error(
			"FillPath with nil fill should return empty string",
		)
	}

	content = FillPath(
		nil,
		NewSolidFill(Red),
		bounds,
	)
	if content != "" {
		t.Error(
			"FillPath with nil path should return empty string",
		)
	}
}

func TestSetFill(t *testing.T) {
	fill := NewSolidFill(Green)
	bounds := NewRect(0, 0, 100, 100)
	content := SetFill(fill, bounds)

	if !strings.Contains(content, "rg") {
		t.Error(
			"SetFill should contain fill color",
		)
	}
	if strings.Contains(content, "f") &&
		!strings.Contains(content, "rg") {
		t.Error(
			"SetFill should not contain fill operator",
		)
	}
}

func TestSetFill_NoFill(t *testing.T) {
	bounds := NewRect(0, 0, 100, 100)
	content := SetFill(nil, bounds)

	if content != "" {
		t.Error(
			"SetFill with nil should return empty string",
		)
	}
}

func TestGradient_EmptyStops(t *testing.T) {
	lg := &LinearGradient{Stops: nil}
	if !lg.IsNone() {
		t.Error(
			"LinearGradient with no stops should be none",
		)
	}

	rg := &RadialGradient{Stops: nil}
	if !rg.IsNone() {
		t.Error(
			"RadialGradient with no stops should be none",
		)
	}
}

func TestGradient_SingleStop(t *testing.T) {
	lg := &LinearGradient{
		Stops: []GradientStop{
			{Position: 0.5, Color: Red},
		},
	}
	c := lg.ColorAt(0.25)
	if c != Red {
		t.Error(
			"Gradient with single stop should always return that color",
		)
	}

	rg := &RadialGradient{
		Stops: []GradientStop{
			{Position: 0.5, Color: Blue},
		},
	}
	c = rg.ColorAt(0.75)
	if c != Blue {
		t.Error(
			"Gradient with single stop should always return that color",
		)
	}
}

func TestLinearGradient_ColorFunction_MultipleStops(
	t *testing.T,
) {
	g := NewLinearGradient(0, Red, Blue)
	g.AddStop(0.5, Green)

	fn := g.colorFunction()
	if fn["FunctionType"] != 3 {
		t.Error(
			"Multi-stop gradient should use stitching function (type 3)",
		)
	}
	functions, ok := fn["Functions"].([]map[string]interface{})
	if !ok || len(functions) != 2 {
		t.Error(
			"Should have 2 sub-functions for 3 stops",
		)
	}
}

func TestRadialGradient_ColorFunction_TwoStops(
	t *testing.T,
) {
	g := NewRadialGradient(Red, Blue)

	fn := g.colorFunction()
	if fn["FunctionType"] != 2 {
		t.Error(
			"Two-stop gradient should use exponential function (type 2)",
		)
	}
}

func TestFillType_Constants(t *testing.T) {
	// Verify fill type constants are distinct
	types := []FillType{
		FillTypeNone,
		FillTypeSolid,
		FillTypeLinearGradient,
		FillTypeRadialGradient,
		FillTypePattern,
	}

	for i := range types {
		for j := i + 1; j < len(types); j++ {
			if types[i] == types[j] {
				t.Errorf(
					"FillType constants should be distinct: %d == %d",
					types[i],
					types[j],
				)
			}
		}
	}
}

func TestPatternType_Constants(t *testing.T) {
	// Verify pattern type constants are distinct
	types := []PatternType{
		PatternTypeSolid,
		PatternTypeHorizontalLines,
		PatternTypeVerticalLines,
		PatternTypeDiagonalLines,
		PatternTypeCrossHatch,
		PatternTypeDots,
	}

	for i := range types {
		for j := i + 1; j < len(types); j++ {
			if types[i] == types[j] {
				t.Errorf(
					"PatternType constants should be distinct: %d == %d",
					types[i],
					types[j],
				)
			}
		}
	}
}

func TestGradientExtend_Constants(t *testing.T) {
	// Verify gradient extend constants are distinct
	extends := []GradientExtend{
		GradientExtendPad,
		GradientExtendRepeat,
		GradientExtendReflect,
	}

	for i := range extends {
		for j := i + 1; j < len(extends); j++ {
			if extends[i] == extends[j] {
				t.Errorf(
					"GradientExtend constants should be distinct: %d == %d",
					extends[i],
					extends[j],
				)
			}
		}
	}
}
