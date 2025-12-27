package drawing

import (
	"strings"
	"testing"
)

func TestClipRule_String(t *testing.T) {
	tests := []struct {
		rule     ClipRule
		expected string
	}{
		{ClipNonZeroRule, "W"},
		{ClipEvenOddRule, "W*"},
	}

	for _, tt := range tests {
		result := tt.rule.String()
		if result != tt.expected {
			t.Errorf(
				"ClipRule.String(): expected %q, got %q",
				tt.expected,
				result,
			)
		}
	}
}

func TestClipContext_PushPopClip(t *testing.T) {
	ctx := NewClipContext()

	// Initially no clips
	if ctx.Depth() != 0 {
		t.Errorf(
			"Initial depth should be 0, got %d",
			ctx.Depth(),
		)
	}

	// Push a rectangular clip
	path := ClipRect(10, 10, 100, 100)
	ctx.PushClip(path, ClipNonZeroRule)

	if ctx.Depth() != 1 {
		t.Errorf(
			"Depth after PushClip should be 1, got %d",
			ctx.Depth(),
		)
	}

	result := ctx.String()
	if !strings.Contains(result, "q\n") {
		t.Error(
			"PushClip should contain 'q' (save graphics state)",
		)
	}
	if !strings.Contains(result, "re\n") {
		t.Error(
			"PushClip should contain rectangle path",
		)
	}
	if !strings.Contains(result, "W n\n") {
		t.Error(
			"PushClip should contain 'W n' (clip non-zero)",
		)
	}

	// Pop the clip
	ctx.PopClip()

	if ctx.Depth() != 0 {
		t.Errorf(
			"Depth after PopClip should be 0, got %d",
			ctx.Depth(),
		)
	}

	result = ctx.String()
	if !strings.Contains(result, "Q\n") {
		t.Error(
			"PopClip should contain 'Q' (restore graphics state)",
		)
	}
}

func TestClipContext_EvenOddRule(t *testing.T) {
	ctx := NewClipContext()
	path := ClipCircle(50, 50, 25)
	ctx.PushClip(path, ClipEvenOddRule)

	result := ctx.String()
	if !strings.Contains(result, "W* n\n") {
		t.Errorf(
			"EvenOdd clip should contain 'W* n', got %q",
			result,
		)
	}
}

func TestClipContext_NestedClips(t *testing.T) {
	ctx := NewClipContext()

	// Push first clip
	ctx.PushClip(
		ClipRect(0, 0, 200, 200),
		ClipNonZeroRule,
	)
	if ctx.Depth() != 1 {
		t.Errorf(
			"Expected depth 1, got %d",
			ctx.Depth(),
		)
	}

	// Push second clip (nested)
	ctx.PushClip(
		ClipCircle(100, 100, 50),
		ClipNonZeroRule,
	)
	if ctx.Depth() != 2 {
		t.Errorf(
			"Expected depth 2, got %d",
			ctx.Depth(),
		)
	}

	// Pop one clip
	ctx.PopClip()
	if ctx.Depth() != 1 {
		t.Errorf(
			"Expected depth 1 after one pop, got %d",
			ctx.Depth(),
		)
	}

	// Pop all remaining
	ctx.PopAllClips()
	if ctx.Depth() != 0 {
		t.Errorf(
			"Expected depth 0 after PopAllClips, got %d",
			ctx.Depth(),
		)
	}

	result := ctx.String()
	if strings.Count(result, "q\n") != 2 {
		t.Errorf(
			"Expected 2 'q' operators, got %d",
			strings.Count(result, "q\n"),
		)
	}
	if strings.Count(result, "Q\n") != 2 {
		t.Errorf(
			"Expected 2 'Q' operators, got %d",
			strings.Count(result, "Q\n"),
		)
	}
}

func TestClipContext_PushClipWithTransform(
	t *testing.T,
) {
	ctx := NewClipContext()
	path := ClipRect(0, 0, 100, 100)
	transform := Identity().Translate(50, 50).
		Scale(2, 2)

	ctx.PushClipWithTransform(
		path,
		ClipNonZeroRule,
		transform,
	)

	result := ctx.String()
	if !strings.Contains(result, "q\n") {
		t.Error(
			"Should contain save graphics state",
		)
	}
	if !strings.Contains(result, "cm\n") {
		t.Error("Should contain transform matrix")
	}
	if !strings.Contains(result, "W n\n") {
		t.Error("Should contain clip operator")
	}

	// Current transform should reflect the applied transform
	currentTransform := ctx.CurrentTransform()
	if currentTransform.IsIdentity() {
		t.Error(
			"Current transform should not be identity after applying transform",
		)
	}
}

func TestClipContext_WriteOperator(t *testing.T) {
	ctx := NewClipContext()
	ctx.PushClip(
		ClipRect(0, 0, 100, 100),
		ClipNonZeroRule,
	)
	ctx.WriteOperator(
		"1 0 0 rg",
	) // Set fill color
	ctx.WriteOperator(
		"0 0 50 50 re f",
	) // Draw and fill rectangle

	result := ctx.String()
	if !strings.Contains(result, "1 0 0 rg\n") {
		t.Error(
			"WriteOperator should add operator to content",
		)
	}
	if !strings.Contains(
		result,
		"0 0 50 50 re f\n",
	) {
		t.Error(
			"WriteOperator should add second operator",
		)
	}
}

func TestClipContext_WriteContent(t *testing.T) {
	ctx := NewClipContext()
	ctx.WriteContent(
		"BT /F1 12 Tf (Hello) Tj ET\n",
	)

	result := ctx.String()
	if result != "BT /F1 12 Tf (Hello) Tj ET\n" {
		t.Errorf(
			"WriteContent: expected raw content, got %q",
			result,
		)
	}
}

func TestClipContext_Reset(t *testing.T) {
	ctx := NewClipContext()
	ctx.PushClip(
		ClipRect(0, 0, 100, 100),
		ClipNonZeroRule,
	)
	ctx.PushClip(
		ClipCircle(50, 50, 25),
		ClipNonZeroRule,
	)

	ctx.Reset()

	if ctx.Depth() != 0 {
		t.Errorf(
			"Depth after reset should be 0, got %d",
			ctx.Depth(),
		)
	}
	if ctx.String() != "" {
		t.Error(
			"Content after reset should be empty",
		)
	}
	if !ctx.CurrentTransform().IsIdentity() {
		t.Error(
			"Transform after reset should be identity",
		)
	}
}

func TestClipScope(t *testing.T) {
	ctx := NewClipContext()

	// Use scope with defer pattern
	func() {
		scope := NewClipScope(
			ctx,
			ClipRect(10, 10, 80, 80),
			ClipNonZeroRule,
		)
		defer scope.Close()

		ctx.WriteOperator("0.5 g")
		ctx.WriteOperator("20 20 60 60 re f")
	}()

	result := ctx.String()
	if !strings.HasPrefix(result, "q\n") {
		t.Error("ClipScope should start with 'q'")
	}
	if !strings.HasSuffix(result, "Q\n") {
		t.Error("ClipScope should end with 'Q'")
	}
}

func TestClipScope_MultipleClose(t *testing.T) {
	ctx := NewClipContext()
	scope := NewClipScope(
		ctx,
		ClipRect(0, 0, 100, 100),
		ClipNonZeroRule,
	)

	// Close multiple times should be safe
	scope.Close()
	scope.Close()
	scope.Close()

	result := ctx.String()
	if strings.Count(result, "Q\n") != 1 {
		t.Errorf(
			"Multiple Close() calls should only result in one 'Q', got %d",
			strings.Count(result, "Q\n"),
		)
	}
}

func TestClipScope_WithTransform(t *testing.T) {
	ctx := NewClipContext()
	transform := Identity().Rotate(45)

	scope := NewClipScopeWithTransform(
		ctx,
		ClipRect(0, 0, 100, 100),
		ClipNonZeroRule,
		transform,
	)
	ctx.WriteOperator("0 0 50 50 re f")
	scope.Close()

	result := ctx.String()
	if !strings.Contains(result, "cm\n") {
		t.Error(
			"ClipScope with transform should include 'cm' operator",
		)
	}
}

func TestClipRect(t *testing.T) {
	path := ClipRect(10, 20, 100, 50)
	result := path.String()
	expected := "10 20 100 50 re\n"
	if result != expected {
		t.Errorf(
			"ClipRect: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestClipRoundedRect(t *testing.T) {
	path := ClipRoundedRect(0, 0, 100, 50, 10)
	result := path.String()

	if !strings.Contains(result, " m\n") {
		t.Error(
			"ClipRoundedRect should contain moveto",
		)
	}
	if !strings.Contains(result, " c\n") {
		t.Error(
			"ClipRoundedRect should contain curveto for corners",
		)
	}
	if !strings.Contains(result, "h\n") {
		t.Error(
			"ClipRoundedRect should close path",
		)
	}
}

func TestClipRoundedRectVarying(t *testing.T) {
	path := ClipRoundedRectVarying(
		0,
		0,
		100,
		50,
		5,
		10,
		15,
		20,
	)
	result := path.String()

	if !strings.Contains(result, " c\n") {
		t.Error(
			"ClipRoundedRectVarying should contain curveto",
		)
	}
}

func TestClipCircle(t *testing.T) {
	path := ClipCircle(50, 50, 25)
	result := path.String()

	// Circle is approximated with 4 Bezier curves
	if strings.Count(result, " c\n") != 4 {
		t.Errorf(
			"ClipCircle should have 4 curves, got %d",
			strings.Count(result, " c\n"),
		)
	}
}

func TestClipEllipse(t *testing.T) {
	path := ClipEllipse(50, 50, 30, 20)
	result := path.String()

	// Ellipse is also 4 Bezier curves
	if strings.Count(result, " c\n") != 4 {
		t.Errorf(
			"ClipEllipse should have 4 curves, got %d",
			strings.Count(result, " c\n"),
		)
	}
}

func TestClipPath(t *testing.T) {
	original := NewPathBuilder()
	original.MoveTo(0, 0).
		LineTo(100, 0).
		LineTo(50, 100).
		ClosePath()

	copy := ClipPath(original)

	// Should have same content
	if copy.String() != original.String() {
		t.Error(
			"ClipPath should create exact copy of path content",
		)
	}

	// Should be independent (modifying one shouldn't affect other)
	copy.LineTo(200, 200)
	if copy.String() == original.String() {
		t.Error(
			"ClipPath copy should be independent of original",
		)
	}
}

func TestClipPolygon(t *testing.T) {
	path := ClipPolygon(0, 0, 100, 0, 50, 100)
	result := path.String()
	expected := "0 0 m\n100 0 l\n50 100 l\nh\n"
	if result != expected {
		t.Errorf(
			"ClipPolygon: expected %q, got %q",
			expected,
			result,
		)
	}
}

func TestRectClipOperators(t *testing.T) {
	result := RectClipOperators(
		10,
		20,
		100,
		50,
		ClipNonZeroRule,
	)

	if !strings.Contains(
		result,
		"10 20 100 50 re\n",
	) {
		t.Error(
			"RectClipOperators should contain rectangle",
		)
	}
	if !strings.Contains(result, "W n\n") {
		t.Error(
			"RectClipOperators should contain clip operator",
		)
	}
}

func TestCircleClipOperators(t *testing.T) {
	result := CircleClipOperators(
		50,
		50,
		25,
		ClipEvenOddRule,
	)

	if !strings.Contains(result, " c\n") {
		t.Error(
			"CircleClipOperators should contain curves",
		)
	}
	if !strings.Contains(result, "W* n\n") {
		t.Error(
			"CircleClipOperators should contain even-odd clip",
		)
	}
}

func TestPathClipOperators(t *testing.T) {
	path := NewPathBuilder().MoveTo(0, 0).
		LineTo(100, 100).
		LineTo(0, 100).
		ClosePath()
	result := PathClipOperators(
		path,
		ClipNonZeroRule,
	)

	if !strings.Contains(result, "h\n") {
		t.Error(
			"PathClipOperators should contain closed path",
		)
	}
	if !strings.Contains(result, "W n\n") {
		t.Error(
			"PathClipOperators should contain clip operator",
		)
	}
}

func TestScopedClipOperators(t *testing.T) {
	path := ClipRect(0, 0, 100, 100)
	content := "0.5 g\n50 50 m\n100 100 l\nS\n"

	result := ScopedClipOperators(
		path,
		ClipNonZeroRule,
		content,
	)

	if !strings.HasPrefix(result, "q\n") {
		t.Error(
			"ScopedClipOperators should start with 'q'",
		)
	}
	if !strings.HasSuffix(result, "Q\n") {
		t.Error(
			"ScopedClipOperators should end with 'Q'",
		)
	}
	if !strings.Contains(result, "W n\n") {
		t.Error(
			"ScopedClipOperators should contain clip operator",
		)
	}
	if !strings.Contains(result, content) {
		t.Error(
			"ScopedClipOperators should contain the content",
		)
	}
}

func TestScopedClipWithTransformOperators(
	t *testing.T,
) {
	path := ClipRect(0, 0, 100, 100)
	transform := Identity().Translate(50, 50)
	content := "0 0 25 25 re f\n"

	result := ScopedClipWithTransformOperators(
		path,
		ClipNonZeroRule,
		transform,
		content,
	)

	if !strings.Contains(result, "cm\n") {
		t.Error(
			"Should contain transform operator",
		)
	}
	if !strings.Contains(result, "q\n") {
		t.Error("Should contain save state")
	}
	if !strings.Contains(result, "Q\n") {
		t.Error("Should contain restore state")
	}
}

func TestScopedClipWithTransformOperators_Identity(
	t *testing.T,
) {
	path := ClipRect(0, 0, 100, 100)
	transform := Identity() // Identity transform
	content := "test content\n"

	result := ScopedClipWithTransformOperators(
		path,
		ClipNonZeroRule,
		transform,
		content,
	)

	if strings.Contains(result, "cm\n") {
		t.Error(
			"Identity transform should not produce 'cm' operator",
		)
	}
}

func TestIntersectClip(t *testing.T) {
	path1 := ClipRect(0, 0, 100, 100)
	path2 := ClipCircle(50, 50, 30)

	combined := IntersectClip(path1, path2)
	result := combined.String()

	// Should contain both paths
	if !strings.Contains(result, "re\n") {
		t.Error(
			"IntersectClip should contain rectangle",
		)
	}
	if !strings.Contains(result, " c\n") {
		t.Error(
			"IntersectClip should contain circle curves",
		)
	}
}

func TestNestedClipOperators(t *testing.T) {
	paths := []*PathBuilder{
		ClipRect(0, 0, 200, 200),
		ClipRect(25, 25, 150, 150),
		ClipCircle(100, 100, 50),
	}
	content := "drawing content\n"

	result := NestedClipOperators(
		paths,
		ClipNonZeroRule,
		content,
	)

	// Should have 3 saves and 3 restores
	if strings.Count(result, "q\n") != 3 {
		t.Errorf(
			"Expected 3 'q' operators, got %d",
			strings.Count(result, "q\n"),
		)
	}
	if strings.Count(result, "Q\n") != 3 {
		t.Errorf(
			"Expected 3 'Q' operators, got %d",
			strings.Count(result, "Q\n"),
		)
	}

	// Should contain content
	if !strings.Contains(result, content) {
		t.Error(
			"Should contain the drawing content",
		)
	}
}

func TestClipContext_PopClipEmpty(t *testing.T) {
	ctx := NewClipContext()

	// PopClip on empty context should be safe
	ctx.PopClip()

	if ctx.Depth() != 0 {
		t.Error(
			"PopClip on empty context should not affect depth",
		)
	}

	// Should not add Q operator when there's nothing to pop
	result := ctx.String()
	if strings.Contains(result, "Q\n") {
		t.Error(
			"PopClip on empty context should not add 'Q' operator",
		)
	}
}

func TestClipContext_PopAllClipsEmpty(
	t *testing.T,
) {
	ctx := NewClipContext()

	// PopAllClips on empty context should be safe
	ctx.PopAllClips()

	if ctx.Depth() != 0 {
		t.Error(
			"PopAllClips on empty context should leave depth at 0",
		)
	}
}

func TestClipContext_CurrentTransform_Initial(
	t *testing.T,
) {
	ctx := NewClipContext()
	transform := ctx.CurrentTransform()

	if !transform.IsIdentity() {
		t.Error(
			"Initial transform should be identity",
		)
	}
}

func TestClipContext_ChainedOperations(
	t *testing.T,
) {
	ctx := NewClipContext()

	// Test method chaining
	ctx.PushClip(ClipRect(0, 0, 100, 100), ClipNonZeroRule).
		WriteOperator("0.5 g").
		WriteContent("custom content\n").
		PopClip()

	result := ctx.String()

	if !strings.Contains(result, "q\n") {
		t.Error("Should contain save")
	}
	if !strings.Contains(result, "0.5 g\n") {
		t.Error("Should contain operator")
	}
	if !strings.Contains(
		result,
		"custom content\n",
	) {
		t.Error("Should contain content")
	}
	if !strings.Contains(result, "Q\n") {
		t.Error("Should contain restore")
	}
}

// Benchmark tests

func BenchmarkClipRect(b *testing.B) {
	for range b.N {
		_ = ClipRect(10, 20, 100, 50)
	}
}

func BenchmarkClipCircle(b *testing.B) {
	for range b.N {
		_ = ClipCircle(50, 50, 25)
	}
}

func BenchmarkClipContext_PushPop(b *testing.B) {
	ctx := NewClipContext()
	path := ClipRect(0, 0, 100, 100)

	b.ResetTimer()
	for range b.N {
		ctx.PushClip(path, ClipNonZeroRule)
		ctx.PopClip()
		ctx.Reset()
	}
}

func BenchmarkScopedClipOperators(b *testing.B) {
	path := ClipRect(0, 0, 100, 100)
	content := "0.5 g\n0 0 50 50 re f\n"

	b.ResetTimer()
	for range b.N {
		_ = ScopedClipOperators(
			path,
			ClipNonZeroRule,
			content,
		)
	}
}
