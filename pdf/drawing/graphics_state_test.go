package drawing

import (
	"strings"
	"testing"
)

func TestNewFullGraphicsState(t *testing.T) {
	gs := NewFullGraphicsState()

	// Check transform is identity
	if !gs.Transform.IsIdentity() {
		t.Error("expected identity transform")
	}

	// Check stroke defaults
	if gs.Stroke.Width != DefaultLineWidth {
		t.Errorf(
			"expected stroke width %v, got %v",
			DefaultLineWidth,
			gs.Stroke.Width,
		)
	}
	if gs.Stroke.Cap != LineCapButt {
		t.Errorf(
			"expected line cap butt, got %v",
			gs.Stroke.Cap,
		)
	}
	if gs.Stroke.Join != LineJoinMiter {
		t.Errorf(
			"expected line join miter, got %v",
			gs.Stroke.Join,
		)
	}
	if gs.Stroke.MiterLimit != DefaultMiterLimit {
		t.Errorf(
			"expected miter limit %v, got %v",
			DefaultMiterLimit,
			gs.Stroke.MiterLimit,
		)
	}

	// Check fill defaults
	if gs.Fill.Color != Black {
		t.Errorf(
			"expected black fill, got %v",
			gs.Fill.Color,
		)
	}

	// Check text defaults
	if gs.Text.FontSize != 12.0 {
		t.Errorf(
			"expected font size 12, got %v",
			gs.Text.FontSize,
		)
	}
	if gs.Text.HorizontalScaling != 100.0 {
		t.Errorf(
			"expected horizontal scaling 100, got %v",
			gs.Text.HorizontalScaling,
		)
	}

	// Check transparency defaults
	if gs.Transparency.StrokeAlpha != 1.0 {
		t.Errorf(
			"expected stroke alpha 1.0, got %v",
			gs.Transparency.StrokeAlpha,
		)
	}
	if gs.Transparency.FillAlpha != 1.0 {
		t.Errorf(
			"expected fill alpha 1.0, got %v",
			gs.Transparency.FillAlpha,
		)
	}
	if gs.Transparency.BlendMode != BlendModeNormal {
		t.Errorf(
			"expected normal blend mode, got %v",
			gs.Transparency.BlendMode,
		)
	}
}

func TestGraphicsStateStackSaveRestore(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	// Verify initial depth is 0
	if gss.Depth() != 0 {
		t.Errorf(
			"expected initial depth 0, got %d",
			gss.Depth(),
		)
	}

	// Save state
	gss.Save()
	if gss.Depth() != 1 {
		t.Errorf(
			"expected depth 1 after save, got %d",
			gss.Depth(),
		)
	}

	// Modify state
	gss.SetLineWidth(5.0)
	gss.SetFillColor(Red)

	// Save again
	gss.Save()
	if gss.Depth() != 2 {
		t.Errorf(
			"expected depth 2 after second save, got %d",
			gss.Depth(),
		)
	}

	// Restore
	gss.Restore()
	if gss.Depth() != 1 {
		t.Errorf(
			"expected depth 1 after restore, got %d",
			gss.Depth(),
		)
	}

	// Check state was restored
	state := gss.CurrentState()
	if state.Stroke.Width != 5.0 {
		t.Errorf(
			"expected stroke width 5.0 after restore, got %v",
			state.Stroke.Width,
		)
	}

	// Restore again
	gss.Restore()
	if gss.Depth() != 0 {
		t.Errorf(
			"expected depth 0 after second restore, got %d",
			gss.Depth(),
		)
	}

	// Check state was restored to defaults
	state = gss.CurrentState()
	if state.Stroke.Width != DefaultLineWidth {
		t.Errorf(
			"expected default stroke width after full restore, got %v",
			state.Stroke.Width,
		)
	}

	// Check operators
	content := gss.String()
	if !strings.Contains(content, "q\n") {
		t.Error("expected 'q' operator in output")
	}
	if !strings.Contains(content, "Q\n") {
		t.Error("expected 'Q' operator in output")
	}
}

func TestGraphicsStateStackTransform(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	// Apply translation
	gss.Translate(100, 50)
	state := gss.CurrentState()
	if state.Transform.E != 100 ||
		state.Transform.F != 50 {
		t.Errorf(
			"expected translation (100, 50), got (%v, %v)",
			state.Transform.E,
			state.Transform.F,
		)
	}

	// Check operator output
	content := gss.String()
	if !strings.Contains(content, "cm\n") {
		t.Error(
			"expected 'cm' operator in output",
		)
	}
}

func TestGraphicsStateStackStrokeSettings(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	// Set various stroke properties
	gss.SetLineWidth(2.5)
	gss.SetLineCap(LineCapRound)
	gss.SetLineJoin(LineJoinBevel)
	gss.SetMiterLimit(5.0)
	gss.SetDashPattern(DashDash)

	state := gss.CurrentState()

	if state.Stroke.Width != 2.5 {
		t.Errorf(
			"expected stroke width 2.5, got %v",
			state.Stroke.Width,
		)
	}
	if state.Stroke.Cap != LineCapRound {
		t.Errorf(
			"expected round cap, got %v",
			state.Stroke.Cap,
		)
	}
	if state.Stroke.Join != LineJoinBevel {
		t.Errorf(
			"expected bevel join, got %v",
			state.Stroke.Join,
		)
	}
	if state.Stroke.MiterLimit != 5.0 {
		t.Errorf(
			"expected miter limit 5.0, got %v",
			state.Stroke.MiterLimit,
		)
	}

	content := gss.String()
	if !strings.Contains(content, "2.5 w") {
		t.Error("expected line width operator")
	}
	if !strings.Contains(content, "1 J") {
		t.Error("expected line cap operator")
	}
	if !strings.Contains(content, "2 j") {
		t.Error("expected line join operator")
	}
	if !strings.Contains(content, "5 M") {
		t.Error("expected miter limit operator")
	}
	if !strings.Contains(content, "d") {
		t.Error("expected dash pattern operator")
	}
}

func TestGraphicsStateStackColors(t *testing.T) {
	gss := NewGraphicsStateStack()

	gss.SetStrokeColor(Red)
	gss.SetFillColor(Blue)

	state := gss.CurrentState()
	if state.Stroke.Color != Red {
		t.Errorf(
			"expected red stroke, got %v",
			state.Stroke.Color,
		)
	}
	if state.Fill.Color != Blue {
		t.Errorf(
			"expected blue fill, got %v",
			state.Fill.Color,
		)
	}

	content := gss.String()
	if !strings.Contains(content, "1 0 0 RG") {
		t.Error("expected stroke color operator")
	}
	if !strings.Contains(content, "0 0 1 rg") {
		t.Error("expected fill color operator")
	}
}

func TestGraphicsStateStackNoDuplicateOperators(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	// Set the same value twice - should only output once
	gss.SetLineWidth(2.0)
	gss.SetLineWidth(
		2.0,
	) // Same value, should be no-op

	content := gss.String()
	count := strings.Count(content, "2 w")
	if count != 1 {
		t.Errorf(
			"expected exactly 1 line width operator, got %d",
			count,
		)
	}
}

func TestGraphicsStateStackReset(t *testing.T) {
	gss := NewGraphicsStateStack()

	gss.Save()
	gss.SetLineWidth(5.0)
	gss.SetFillColor(Green)

	// Reset
	gss.Reset()

	if gss.Depth() != 0 {
		t.Errorf(
			"expected depth 0 after reset, got %d",
			gss.Depth(),
		)
	}

	state := gss.CurrentState()
	if state.Stroke.Width != DefaultLineWidth {
		t.Errorf(
			"expected default stroke width after reset, got %v",
			state.Stroke.Width,
		)
	}

	if gss.String() != "" {
		t.Error(
			"expected empty output after reset",
		)
	}
}

func TestExtGStateEntry(t *testing.T) {
	entry := NewExtGStateEntry("GS1")
	entry.WithStrokeAlpha(0.5).
		WithFillAlpha(0.7).
		WithBlendMode(BlendModeMultiply)

	dict := entry.ToDictEntries()

	if dict["Type"] != "/ExtGState" {
		t.Errorf(
			"expected Type /ExtGState, got %v",
			dict["Type"],
		)
	}
	if dict["CA"] != 0.5 {
		t.Errorf(
			"expected CA 0.5, got %v",
			dict["CA"],
		)
	}
	if dict["ca"] != 0.7 {
		t.Errorf(
			"expected ca 0.7, got %v",
			dict["ca"],
		)
	}
	if dict["BM"] != "/Multiply" {
		t.Errorf(
			"expected BM /Multiply, got %v",
			dict["BM"],
		)
	}

	op := entry.ApplyOperator()
	if op != "/GS1 gs" {
		t.Errorf(
			"expected '/GS1 gs', got '%s'",
			op,
		)
	}
}

func TestExtGStateManager(t *testing.T) {
	manager := NewExtGStateManager()

	// Create alpha state
	alpha := manager.CreateAlpha(0.5, 0.5)
	if alpha.Name != "GS1" {
		t.Errorf(
			"expected name GS1, got %s",
			alpha.Name,
		)
	}

	// Create blend mode state
	blend := manager.CreateBlendMode(
		BlendModeScreen,
	)
	if blend.Name != "GS2" {
		t.Errorf(
			"expected name GS2, got %s",
			blend.Name,
		)
	}

	// Check entries
	entries := manager.Entries()
	if len(entries) != 2 {
		t.Errorf(
			"expected 2 entries, got %d",
			len(entries),
		)
	}

	// Check resource dict
	dict := manager.ResourceDict()
	if dict == nil {
		t.Error("expected non-nil resource dict")
	}
	if _, ok := dict["GS1"]; !ok {
		t.Error("expected GS1 in resource dict")
	}
	if _, ok := dict["GS2"]; !ok {
		t.Error("expected GS2 in resource dict")
	}
}

func TestGraphicsStateStackExtGState(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()
	manager := NewExtGStateManager()

	entry := gss.SetAlpha(manager, 0.5, 0.75)

	// Check that state was updated
	state := gss.CurrentState()
	if state.Transparency.StrokeAlpha != 0.5 {
		t.Errorf(
			"expected stroke alpha 0.5, got %v",
			state.Transparency.StrokeAlpha,
		)
	}
	if state.Transparency.FillAlpha != 0.75 {
		t.Errorf(
			"expected fill alpha 0.75, got %v",
			state.Transparency.FillAlpha,
		)
	}

	// Check that operator was output
	content := gss.String()
	if !strings.Contains(
		content,
		entry.Name+" gs",
	) {
		t.Errorf(
			"expected gs operator with name %s",
			entry.Name,
		)
	}

	// Check used ExtGStates
	used := gss.UsedExtGStates()
	if len(used) != 1 {
		t.Errorf(
			"expected 1 used ExtGState, got %d",
			len(used),
		)
	}
}

func TestGraphicsStateStackWithState(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	// Set initial state
	gss.SetLineWidth(1.0)

	// Use WithState to temporarily change settings
	gss.WithState(
		func(inner *GraphicsStateStack) {
			inner.SetLineWidth(5.0)
			// Inside here, line width is 5.0
			state := inner.CurrentState()
			if state.Stroke.Width != 5.0 {
				t.Errorf(
					"expected line width 5.0 inside WithState, got %v",
					state.Stroke.Width,
				)
			}
		},
	)

	// After WithState, should be restored
	state := gss.CurrentState()
	if state.Stroke.Width != 1.0 {
		t.Errorf(
			"expected line width 1.0 after WithState, got %v",
			state.Stroke.Width,
		)
	}

	// Check that q/Q were emitted
	content := gss.String()
	qCount := strings.Count(content, "q\n")
	qCountEnd := strings.Count(content, "Q\n")
	if qCount != 1 || qCountEnd != 1 {
		t.Errorf(
			"expected 1 q and 1 Q, got %d q and %d Q",
			qCount,
			qCountEnd,
		)
	}
}

func TestGraphicsStateStackWithTransform(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	gss.WithTransform(
		Identity().Translate(100, 50),
		func(inner *GraphicsStateStack) {
			state := inner.CurrentState()
			if state.Transform.E != 100 ||
				state.Transform.F != 50 {
				t.Errorf(
					"expected translation (100, 50) inside WithTransform",
				)
			}
		},
	)

	// After, transform should be identity
	state := gss.CurrentState()
	if !state.Transform.IsIdentity() {
		t.Error(
			"expected identity transform after WithTransform",
		)
	}
}

func TestGraphicsStateStackWithClip(
	t *testing.T,
) {
	gss := NewGraphicsStateStack()

	path := NewPathBuilder().Rectangle(0, 0, 100, 100)

	gss.WithClip(
		path,
		ClipNonZeroRule,
		func(inner *GraphicsStateStack) {
			state := inner.CurrentState()
			if state.Clipping.Depth != 1 {
				t.Errorf(
					"expected clipping depth 1, got %d",
					state.Clipping.Depth,
				)
			}
		},
	)

	content := gss.String()
	if !strings.Contains(content, "W n") {
		t.Error(
			"expected clipping operators in output",
		)
	}
}

func TestStateScope(t *testing.T) {
	gss := NewGraphicsStateStack()

	// Test that scope properly saves and restores
	{
		scope := NewStateScope(gss)
		gss.SetLineWidth(10.0)
		state := gss.CurrentState()
		if state.Stroke.Width != 10.0 {
			t.Errorf(
				"expected line width 10.0, got %v",
				state.Stroke.Width,
			)
		}
		scope.Close()
	}

	state := gss.CurrentState()
	if state.Stroke.Width != DefaultLineWidth {
		t.Errorf(
			"expected default line width after scope close, got %v",
			state.Stroke.Width,
		)
	}

	// Test that multiple Close calls are safe
	scope := NewStateScope(gss)
	scope.Close()
	scope.Close() // Should not panic or cause issues
}

func TestTransformScope(t *testing.T) {
	gss := NewGraphicsStateStack()

	{
		scope := NewTransformScope(
			gss,
			Identity().Rotate(45),
		)
		state := gss.CurrentState()
		if state.Transform.IsIdentity() {
			t.Error(
				"expected non-identity transform inside scope",
			)
		}
		scope.Close()
	}

	state := gss.CurrentState()
	if !state.Transform.IsIdentity() {
		t.Error(
			"expected identity transform after scope close",
		)
	}
}

func TestClippingScope(t *testing.T) {
	gss := NewGraphicsStateStack()

	path := NewPathBuilder().Circle(50, 50, 25)

	{
		scope := NewClippingScope(
			gss,
			path,
			ClipEvenOddRule,
		)
		// Content inside clip
		gss.WriteOperator("0 0 100 100 re f")
		scope.Close()
	}

	content := gss.String()
	if !strings.Contains(content, "W* n") {
		t.Error(
			"expected even-odd clipping operators in output",
		)
	}
}

func TestTextState(t *testing.T) {
	ts := NewTextState()

	if ts.FontSize != 12.0 {
		t.Errorf(
			"expected font size 12, got %v",
			ts.FontSize,
		)
	}
	if ts.HorizontalScaling != 100.0 {
		t.Errorf(
			"expected horizontal scaling 100, got %v",
			ts.HorizontalScaling,
		)
	}
	if ts.RenderingMode != TextRenderFill {
		t.Errorf(
			"expected fill rendering mode, got %v",
			ts.RenderingMode,
		)
	}

	// Test clone
	ts.FontName = "F1"
	ts.FontSize = 24.0
	clone := ts.Clone()
	if clone.FontName != "F1" ||
		clone.FontSize != 24.0 {
		t.Error(
			"clone did not copy values correctly",
		)
	}

	// Modify original, clone should be unchanged
	ts.FontSize = 36.0
	if clone.FontSize != 24.0 {
		t.Error(
			"clone was affected by original modification",
		)
	}
}

func TestTransparencyState(t *testing.T) {
	ts := NewTransparencyState()

	if !ts.IsOpaque() {
		t.Error(
			"expected new transparency state to be opaque",
		)
	}

	ts.StrokeAlpha = 0.5
	if ts.IsOpaque() {
		t.Error("expected non-opaque state")
	}

	ts.StrokeAlpha = 1.0
	ts.FillAlpha = 0.8
	if ts.IsOpaque() {
		t.Error("expected non-opaque state")
	}
}

func TestBlendModes(t *testing.T) {
	modes := []BlendMode{
		BlendModeNormal,
		BlendModeMultiply,
		BlendModeScreen,
		BlendModeOverlay,
		BlendModeDarken,
		BlendModeLighten,
	}

	for _, mode := range modes {
		if mode == "" {
			t.Errorf(
				"blend mode should not be empty string",
			)
		}
	}
}

func TestExtGStatePresets(t *testing.T) {
	manager := NewExtGStateManager()

	opacity50 := ExtGStateOpacity50(manager)
	if opacity50.StrokeAlpha == nil ||
		*opacity50.StrokeAlpha != 0.5 {
		t.Error("expected 50% stroke opacity")
	}
	if opacity50.FillAlpha == nil ||
		*opacity50.FillAlpha != 0.5 {
		t.Error("expected 50% fill opacity")
	}

	multiply := ExtGStateMultiply(manager)
	if multiply.BlendMode != BlendModeMultiply {
		t.Errorf(
			"expected multiply blend mode, got %v",
			multiply.BlendMode,
		)
	}

	screen := ExtGStateScreen(manager)
	if screen.BlendMode != BlendModeScreen {
		t.Errorf(
			"expected screen blend mode, got %v",
			screen.BlendMode,
		)
	}
}

func TestSetStrokeStyle(t *testing.T) {
	gss := NewGraphicsStateStack()

	style := NewStrokeStyle()
	style.Width = 3.0
	style.Cap = LineCapRound
	style.Join = LineJoinRound
	style.Color = Green
	style.Dash = DashDot

	gss.SetStrokeStyle(style)

	state := gss.CurrentState()
	if state.Stroke.Width != 3.0 {
		t.Errorf(
			"expected width 3.0, got %v",
			state.Stroke.Width,
		)
	}
	if state.Stroke.Cap != LineCapRound {
		t.Errorf(
			"expected round cap, got %v",
			state.Stroke.Cap,
		)
	}
	if state.Stroke.Join != LineJoinRound {
		t.Errorf(
			"expected round join, got %v",
			state.Stroke.Join,
		)
	}
	if state.Stroke.Color != Green {
		t.Errorf(
			"expected green color, got %v",
			state.Stroke.Color,
		)
	}

	content := gss.String()
	if !strings.Contains(content, "3 w") {
		t.Error("expected line width in output")
	}
}

func TestFullGraphicsStateClone(t *testing.T) {
	gs := NewFullGraphicsState()
	gs.Transform = Identity().Translate(10, 20)
	gs.Stroke.Width = 5.0
	gs.Stroke.Dash.Array = []float64{2, 3, 4}
	gs.Fill.Color = Red
	gs.Text.FontSize = 24.0
	gs.Transparency.StrokeAlpha = 0.5

	clone := gs.Clone()

	// Verify clone has same values
	if clone.Transform.E != 10 ||
		clone.Transform.F != 20 {
		t.Error("transform not cloned correctly")
	}
	if clone.Stroke.Width != 5.0 {
		t.Error(
			"stroke width not cloned correctly",
		)
	}
	if len(clone.Stroke.Dash.Array) != 3 {
		t.Error("dash array not cloned correctly")
	}
	if clone.Fill.Color != Red {
		t.Error("fill color not cloned correctly")
	}
	if clone.Text.FontSize != 24.0 {
		t.Error(
			"text font size not cloned correctly",
		)
	}
	if clone.Transparency.StrokeAlpha != 0.5 {
		t.Error(
			"transparency not cloned correctly",
		)
	}

	// Modify original, clone should be unchanged
	gs.Stroke.Width = 10.0
	gs.Stroke.Dash.Array[0] = 99
	gs.Transform.E = 100

	if clone.Stroke.Width != 5.0 {
		t.Error(
			"clone was affected by original modification",
		)
	}
	if clone.Stroke.Dash.Array[0] != 2 {
		t.Error(
			"clone dash array was affected by original modification",
		)
	}
}

func TestWithHelperFunctions(t *testing.T) {
	gss := NewGraphicsStateStack()

	// Test WithTranslation
	called := false
	gss.WithTranslation(
		50,
		100,
		func(inner *GraphicsStateStack) {
			called = true
			state := inner.CurrentState()
			if state.Transform.E != 50 ||
				state.Transform.F != 100 {
				t.Error(
					"translation not applied correctly",
				)
			}
		},
	)
	if !called {
		t.Error("callback not called")
	}

	// Test WithScale
	gss.Reset()
	gss.WithScale(
		2,
		3,
		func(inner *GraphicsStateStack) {
			state := inner.CurrentState()
			if state.Transform.A != 2 ||
				state.Transform.D != 3 {
				t.Error(
					"scale not applied correctly",
				)
			}
		},
	)

	// Test WithRotation
	gss.Reset()
	gss.WithRotation(
		90,
		func(inner *GraphicsStateStack) {
			// 90 degree rotation: [0, 1, -1, 0, 0, 0]
			state := inner.CurrentState()
			if state.Transform.IsIdentity() {
				t.Error(
					"rotation should not be identity",
				)
			}
		},
	)
}

func TestSetFlatness(t *testing.T) {
	gss := NewGraphicsStateStack()

	gss.SetFlatness(2.0)

	state := gss.CurrentState()
	if state.Flatness != 2.0 {
		t.Errorf(
			"expected flatness 2.0, got %v",
			state.Flatness,
		)
	}

	content := gss.String()
	if !strings.Contains(content, "2 i") {
		t.Error(
			"expected flatness operator in output",
		)
	}

	// Setting same value should not output again
	gss.SetFlatness(2.0)
	count := strings.Count(gss.String(), "2 i")
	if count != 1 {
		t.Errorf(
			"expected exactly 1 flatness operator, got %d",
			count,
		)
	}
}
