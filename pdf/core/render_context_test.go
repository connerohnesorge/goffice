package core

import (
	"math"
	"testing"
)

func TestPoint_Basic(t *testing.T) {
	p := NewPoint(10, 20)
	if p.X != 10 || p.Y != 20 {
		t.Errorf(
			"NewPoint(10, 20) = %v, want (10, 20)",
			p,
		)
	}
}

func TestPoint_Add(t *testing.T) {
	p1 := NewPoint(10, 20)
	p2 := NewPoint(5, 15)
	result := p1.Add(p2)

	if result.X != 15 || result.Y != 35 {
		t.Errorf(
			"(%v).Add(%v) = %v, want (15, 35)",
			p1,
			p2,
			result,
		)
	}
}

func TestPoint_Sub(t *testing.T) {
	p1 := NewPoint(10, 20)
	p2 := NewPoint(5, 15)
	result := p1.Sub(p2)

	if result.X != 5 || result.Y != 5 {
		t.Errorf(
			"(%v).Sub(%v) = %v, want (5, 5)",
			p1,
			p2,
			result,
		)
	}
}

func TestPoint_Scale(t *testing.T) {
	p := NewPoint(10, 20)
	result := p.Scale(2)

	if result.X != 20 || result.Y != 40 {
		t.Errorf(
			"(%v).Scale(2) = %v, want (20, 40)",
			p,
			result,
		)
	}
}

func TestPoint_Distance(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	result := p1.Distance(p2)

	if !floatEqualsWithTolerance(
		result,
		5,
		0.0001,
	) {
		t.Errorf(
			"Distance from (0,0) to (3,4) = %v, want 5",
			result,
		)
	}
}

func TestNewRenderingContext(t *testing.T) {
	rc := NewRenderingContext(
		612,
		792,
	) // Letter size

	if rc.GetPageWidth() != 612 {
		t.Errorf(
			"GetPageWidth() = %v, want 612",
			rc.GetPageWidth(),
		)
	}
	if rc.GetPageHeight() != 792 {
		t.Errorf(
			"GetPageHeight() = %v, want 792",
			rc.GetPageHeight(),
		)
	}
	if rc.OriginDepth() != 0 {
		t.Errorf(
			"Initial OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
	if rc.TransformDepth() != 0 {
		t.Errorf(
			"Initial TransformDepth() = %v, want 0",
			rc.TransformDepth(),
		)
	}
}

func TestNewRenderingContextFromPageSize(
	t *testing.T,
) {
	rc := NewRenderingContextFromPageSize(
		PageSizeA4,
	)

	if !floatEqualsWithTolerance(
		rc.GetPageWidth(),
		PageSizeA4.Width,
		0.1,
	) {
		t.Errorf(
			"Width = %v, want %v",
			rc.GetPageWidth(),
			PageSizeA4.Width,
		)
	}
	if !floatEqualsWithTolerance(
		rc.GetPageHeight(),
		PageSizeA4.Height,
		0.1,
	) {
		t.Errorf(
			"Height = %v, want %v",
			rc.GetPageHeight(),
			PageSizeA4.Height,
		)
	}
}

func TestNewRenderingContextFromPageOptions(
	t *testing.T,
) {
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(NewMarginsUniform(72))
	// 1 inch margins
	rc := NewRenderingContextFromPageOptions(opts)

	// Check margins
	margins := rc.GetMargins()
	if margins.Top != 72 || margins.Right != 72 ||
		margins.Bottom != 72 ||
		margins.Left != 72 {
		t.Errorf(
			"Margins = %v, want uniform 72",
			margins,
		)
	}

	// Check content bounds reflect margins
	bounds := rc.GetContentBounds()
	expectedWidth := 612.0 - 72.0 - 72.0 // Letter width minus margins
	expectedHeight := 792.0 - 72.0 - 72.0

	if !floatEqualsWithTolerance(
		bounds.Width(),
		expectedWidth,
		0.1,
	) {
		t.Errorf(
			"Content width = %v, want %v",
			bounds.Width(),
			expectedWidth,
		)
	}
	if !floatEqualsWithTolerance(
		bounds.Height(),
		expectedHeight,
		0.1,
	) {
		t.Errorf(
			"Content height = %v, want %v",
			bounds.Height(),
			expectedHeight,
		)
	}
}

func TestRenderingContext_PushPopOrigin(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push an origin
	rc.PushOrigin(100, 50)

	if rc.OriginDepth() != 1 {
		t.Errorf(
			"After push, OriginDepth() = %v, want 1",
			rc.OriginDepth(),
		)
	}

	origin := rc.CurrentOrigin()
	if origin.X != 100 || origin.Y != 50 {
		t.Errorf(
			"CurrentOrigin() = %v, want (100, 50)",
			origin,
		)
	}

	// Pop the origin
	result := rc.PopOrigin()
	if !result {
		t.Error("PopOrigin() should return true")
	}
	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After pop, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}

	// Try to pop the root origin
	result = rc.PopOrigin()
	if result {
		t.Error(
			"PopOrigin() should return false when trying to pop root",
		)
	}
}

func TestRenderingContext_NestedOrigins(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push multiple nested origins
	rc.PushOrigin(100, 100)
	rc.PushOrigin(50, 50)

	if rc.OriginDepth() != 2 {
		t.Errorf(
			"After two pushes, OriginDepth() = %v, want 2",
			rc.OriginDepth(),
		)
	}

	// Current origin should be cumulative
	origin := rc.CurrentOrigin()
	if origin.X != 150 || origin.Y != 150 {
		t.Errorf(
			"Cumulative origin = %v, want (150, 150)",
			origin,
		)
	}

	// Pop one level
	rc.PopOrigin()
	origin = rc.CurrentOrigin()
	if origin.X != 100 || origin.Y != 100 {
		t.Errorf(
			"After pop, origin = %v, want (100, 100)",
			origin,
		)
	}
}

func TestRenderingContext_TransformPoint(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Test basic Y-flip (OOXML origin at 0,0 should be at page height in PDF)
	x, y := rc.TransformPoint(0, 0)
	if x != 0 ||
		!floatEqualsWithTolerance(y, 792, 0.001) {
		t.Errorf(
			"TransformPoint(0, 0) = (%v, %v), want (0, 792)",
			x,
			y,
		)
	}

	// Point at (72, 72) in OOXML should be at (72, 720) in PDF
	x, y = rc.TransformPoint(72, 72)
	if !floatEqualsWithTolerance(x, 72, 0.001) ||
		!floatEqualsWithTolerance(y, 720, 0.001) {
		t.Errorf(
			"TransformPoint(72, 72) = (%v, %v), want (72, 720)",
			x,
			y,
		)
	}
}

func TestRenderingContext_TransformPointWithOrigin(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push an origin offset
	rc.PushOrigin(100, 100)

	// A point at (0, 0) relative to origin should be at (100, 692) in PDF
	// OOXML: (100, 100) -> PDF Y = 792 - 100 = 692
	x, y := rc.TransformPoint(0, 0)
	if !floatEqualsWithTolerance(x, 100, 0.001) ||
		!floatEqualsWithTolerance(y, 692, 0.001) {
		t.Errorf(
			"TransformPoint(0, 0) with origin = (%v, %v), want (100, 692)",
			x,
			y,
		)
	}

	// A point at (50, 50) relative to origin
	// OOXML: (100+50, 100+50) = (150, 150) -> PDF Y = 792 - 150 = 642
	x, y = rc.TransformPoint(50, 50)
	if !floatEqualsWithTolerance(x, 150, 0.001) ||
		!floatEqualsWithTolerance(y, 642, 0.001) {
		t.Errorf(
			"TransformPoint(50, 50) with origin = (%v, %v), want (150, 642)",
			x,
			y,
		)
	}
}

func TestRenderingContext_TransformPointEMU(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// 914400 EMU = 1 inch = 72 points
	x, y := rc.TransformPointEMU(914400, 914400)

	expectedX := 72.0
	expectedY := 792.0 - 72.0 // 720

	if !floatEqualsWithTolerance(
		x,
		expectedX,
		0.001,
	) ||
		!floatEqualsWithTolerance(
			y,
			expectedY,
			0.001,
		) {
		t.Errorf(
			"TransformPointEMU(914400, 914400) = (%v, %v), want (%v, %v)",
			x,
			y,
			expectedX,
			expectedY,
		)
	}
}

func TestRenderingContext_TransformRect(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Transform a 72x72 rectangle at (72, 72) in OOXML
	rect := rc.TransformRect(72, 72, 72, 72)

	// In PDF coordinates:
	// OOXML top-left (72, 72) -> PDF (72, 720)
	// OOXML bottom-right (144, 144) -> PDF (144, 648)
	// Rectangle LLX, LLY should be (72, 648), URX, URY should be (144, 720)

	if !floatEqualsWithTolerance(
		rect.LLX,
		72,
		0.001,
	) {
		t.Errorf(
			"rect.LLX = %v, want 72",
			rect.LLX,
		)
	}
	if !floatEqualsWithTolerance(
		rect.LLY,
		648,
		0.001,
	) {
		t.Errorf(
			"rect.LLY = %v, want 648",
			rect.LLY,
		)
	}
	if !floatEqualsWithTolerance(
		rect.URX,
		144,
		0.001,
	) {
		t.Errorf(
			"rect.URX = %v, want 144",
			rect.URX,
		)
	}
	if !floatEqualsWithTolerance(
		rect.URY,
		720,
		0.001,
	) {
		t.Errorf(
			"rect.URY = %v, want 720",
			rect.URY,
		)
	}
}

func TestRenderingContext_TransformRectEMU(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// 1 inch square at 1 inch from origin
	rect := rc.TransformRectEMU(
		914400,
		914400,
		914400,
		914400,
	)

	// Same as TransformRect test but with EMU input
	if !floatEqualsWithTolerance(
		rect.Width(),
		72,
		0.001,
	) {
		t.Errorf(
			"rect.Width() = %v, want 72",
			rect.Width(),
		)
	}
	if !floatEqualsWithTolerance(
		rect.Height(),
		72,
		0.001,
	) {
		t.Errorf(
			"rect.Height() = %v, want 72",
			rect.Height(),
		)
	}
}

func TestRenderingContext_TransformSize(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Size should not be affected by origin or Y-flip
	w, h := rc.TransformSize(100, 50)
	if w != 100 || h != 50 {
		t.Errorf(
			"TransformSize(100, 50) = (%v, %v), want (100, 50)",
			w,
			h,
		)
	}
}

func TestRenderingContext_PushPopTransform(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push a translation matrix
	rc.PushTransform(TranslationMatrix(100, 100))

	if rc.TransformDepth() != 1 {
		t.Errorf(
			"After push, TransformDepth() = %v, want 1",
			rc.TransformDepth(),
		)
	}

	if rc.CurrentTransform().IsIdentity() {
		t.Error(
			"After push, transform should not be identity",
		)
	}

	// Pop
	result := rc.PopTransform()
	if !result {
		t.Error(
			"PopTransform() should return true",
		)
	}
	if !rc.CurrentTransform().IsIdentity() {
		t.Error(
			"After pop, transform should be identity",
		)
	}
}

func TestRenderingContext_PushOOXMLTransform(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push a rotated OOXML transform
	xfrm := NewOOXMLTransform(
		0,
		0,
		914400,
		914400,
	).WithRotationDegrees(90)
	rc.PushOOXMLTransform(xfrm)

	if rc.TransformDepth() != 1 {
		t.Errorf(
			"After push, TransformDepth() = %v, want 1",
			rc.TransformDepth(),
		)
	}

	// The transform should not be identity
	if rc.CurrentTransform().IsIdentity() {
		t.Error(
			"Rotated transform should not be identity",
		)
	}
}

func TestRenderingContext_ContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Initial bounds should be the full page
	bounds := rc.GetContentBounds()
	if bounds.Width() != 612 ||
		bounds.Height() != 792 {
		t.Errorf(
			"Initial bounds = %v, want (612, 792)",
			bounds,
		)
	}

	// Set new bounds
	newBounds := NewRectangleFromSize(
		72,
		72,
		468,
		648,
	)
	rc.SetContentBounds(newBounds)

	bounds = rc.GetContentBounds()
	if bounds.LLX != 72 || bounds.LLY != 72 {
		t.Errorf(
			"After set, bounds origin = (%v, %v), want (72, 72)",
			bounds.LLX,
			bounds.LLY,
		)
	}
}

func TestRenderingContext_PushPopContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push nested content bounds
	rc.PushContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	if rc.ContentBoundsDepth() != 1 {
		t.Errorf(
			"After push, ContentBoundsDepth() = %v, want 1",
			rc.ContentBoundsDepth(),
		)
	}

	bounds := rc.GetContentBounds()
	if bounds.LLX != 100 || bounds.LLY != 100 {
		t.Errorf(
			"Bounds origin = (%v, %v), want (100, 100)",
			bounds.LLX,
			bounds.LLY,
		)
	}

	// Pop
	result := rc.PopContentBounds()
	if !result {
		t.Error(
			"PopContentBounds() should return true",
		)
	}
	if rc.ContentBoundsDepth() != 0 {
		t.Errorf(
			"After pop, ContentBoundsDepth() = %v, want 0",
			rc.ContentBoundsDepth(),
		)
	}
}

func TestRenderingContext_ContentBoundsIntersection(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Set initial bounds
	rc.SetContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	// Push overlapping bounds - should intersect
	rc.PushContentBounds(
		NewRectangleFromSize(200, 200, 400, 400),
	)

	bounds := rc.GetContentBounds()

	// Intersection should be (200, 200) to (500, 600)
	if bounds.LLX != 200 || bounds.LLY != 200 {
		t.Errorf(
			"Intersected bounds origin = (%v, %v), want (200, 200)",
			bounds.LLX,
			bounds.LLY,
		)
	}
	if bounds.URX != 500 || bounds.URY != 600 {
		t.Errorf(
			"Intersected bounds corner = (%v, %v), want (500, 600)",
			bounds.URX,
			bounds.URY,
		)
	}
}

func TestRenderingContext_IsPointInContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)
	rc.SetContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	// Point inside
	if !rc.IsPointInContentBounds(200, 300) {
		t.Error(
			"Point (200, 300) should be in bounds",
		)
	}

	// Point outside
	if rc.IsPointInContentBounds(50, 50) {
		t.Error(
			"Point (50, 50) should not be in bounds",
		)
	}
}

func TestRenderingContext_IsRectInContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)
	rc.SetContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	// Rect fully inside
	innerRect := NewRectangleFromSize(
		150,
		150,
		100,
		100,
	)
	if !rc.IsRectInContentBounds(innerRect) {
		t.Error("Inner rect should be in bounds")
	}

	// Rect partially outside
	partialRect := NewRectangleFromSize(
		50,
		50,
		200,
		200,
	)
	if rc.IsRectInContentBounds(partialRect) {
		t.Error(
			"Partial rect should not be fully in bounds",
		)
	}
}

func TestRenderingContext_RectIntersectsContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)
	rc.SetContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	// Rect that intersects
	intersectRect := NewRectangleFromSize(
		50,
		50,
		200,
		200,
	)
	if !rc.RectIntersectsContentBounds(
		intersectRect,
	) {
		t.Error(
			"Intersecting rect should intersect bounds",
		)
	}

	// Rect completely outside
	outsideRect := NewRectangleFromSize(
		600,
		600,
		100,
		100,
	)
	if rc.RectIntersectsContentBounds(
		outsideRect,
	) {
		t.Error(
			"Outside rect should not intersect bounds",
		)
	}
}

func TestRenderingContext_ClipToContentBounds(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)
	rc.SetContentBounds(
		NewRectangleFromSize(100, 100, 400, 600),
	)

	// Clip a rect that extends beyond bounds
	rect := NewRectangleFromSize(
		50,
		150,
		600,
		200,
	)
	clipped, valid := rc.ClipToContentBounds(rect)

	if !valid {
		t.Error("Clipped rect should be valid")
	}
	if clipped.LLX != 100 {
		t.Errorf(
			"Clipped LLX = %v, want 100",
			clipped.LLX,
		)
	}
	if clipped.URX != 500 {
		t.Errorf(
			"Clipped URX = %v, want 500",
			clipped.URX,
		)
	}
}

func TestRenderingContext_Reset(t *testing.T) {
	rc := NewRenderingContext(612, 792)

	// Push some state
	rc.PushOrigin(100, 100)
	rc.PushTransform(TranslationMatrix(50, 50))
	rc.PushContentBounds(
		NewRectangleFromSize(100, 100, 200, 200),
	)

	// Reset
	rc.Reset()

	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After reset, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
	if rc.TransformDepth() != 0 {
		t.Errorf(
			"After reset, TransformDepth() = %v, want 0",
			rc.TransformDepth(),
		)
	}
	if rc.ContentBoundsDepth() != 0 {
		t.Errorf(
			"After reset, ContentBoundsDepth() = %v, want 0",
			rc.ContentBoundsDepth(),
		)
	}
}

func TestRenderingContext_SaveRestoreState(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push some initial state
	rc.PushOrigin(100, 100)
	rc.PushTransform(TranslationMatrix(50, 50))

	// Save state
	state := rc.SaveState()

	// Push more state
	rc.PushOrigin(50, 50)
	rc.PushTransform(RotationMatrix(math.Pi / 4))
	rc.PushContentBounds(
		NewRectangleFromSize(100, 100, 200, 200),
	)

	if rc.OriginDepth() != 2 {
		t.Errorf(
			"Before restore, OriginDepth() = %v, want 2",
			rc.OriginDepth(),
		)
	}

	// Restore state
	rc.RestoreState(state)

	if rc.OriginDepth() != 1 {
		t.Errorf(
			"After restore, OriginDepth() = %v, want 1",
			rc.OriginDepth(),
		)
	}
	if rc.TransformDepth() != 1 {
		t.Errorf(
			"After restore, TransformDepth() = %v, want 1",
			rc.TransformDepth(),
		)
	}
}

func TestContentAreaContext_Basic(t *testing.T) {
	rc := NewRenderingContext(612, 792)

	// Enter a content area
	cac := rc.EnterContentArea(100, 100, 200, 200)

	// Origin should be offset
	origin := rc.CurrentOrigin()
	if origin.X != 100 || origin.Y != 100 {
		t.Errorf(
			"In content area, origin = %v, want (100, 100)",
			origin,
		)
	}

	// Get bounds
	bounds := cac.GetBounds()
	if bounds.Width() == 0 ||
		bounds.Height() == 0 {
		t.Error(
			"Content area bounds should not be zero",
		)
	}

	// Exit
	cac.ExitContentArea()

	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After exit, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
}

func TestTableCellContext_Basic(t *testing.T) {
	rc := NewRenderingContext(612, 792)

	// Enter a table cell
	tcc := rc.EnterTableCell(
		72,
		100,
		150,
		50,
		0,
		0,
	)

	// Check cell position
	row, col := tcc.GetCellPosition()
	if row != 0 || col != 0 {
		t.Errorf(
			"Cell position = (%v, %v), want (0, 0)",
			row,
			col,
		)
	}

	// Check cell size
	width, height := tcc.GetCellSize()
	if width != 150 || height != 50 {
		t.Errorf(
			"Cell size = (%v, %v), want (150, 50)",
			width,
			height,
		)
	}

	// Exit
	tcc.ExitContentArea()
}

func TestShapeGroupContext_Basic(t *testing.T) {
	rc := NewRenderingContext(612, 792)

	// Create a group transform
	groupXfrm := NewOOXMLTransform(
		914400,
		914400,
		914400*2,
		914400*2,
	) // 1 inch offset, 2 inch size

	// Enter the shape group
	sgc := rc.EnterShapeGroup(groupXfrm)

	// Origin should be offset by the group position
	origin := rc.CurrentOrigin()
	if !floatEqualsWithTolerance(
		origin.X,
		72,
		0.1,
	) ||
		!floatEqualsWithTolerance(
			origin.Y,
			72,
			0.1,
		) {
		t.Errorf(
			"In shape group, origin = %v, want approximately (72, 72)",
			origin,
		)
	}

	// Exit
	sgc.ExitShapeGroup()

	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After exit, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
}

func TestShapeGroupContext_WithRotation(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Create a rotated group transform
	groupXfrm := NewOOXMLTransform(
		914400,
		914400,
		914400,
		914400,
	).WithRotationDegrees(45)

	// Enter the shape group
	sgc := rc.EnterShapeGroup(groupXfrm)

	// Transform depth should be 1 (for the rotation)
	if rc.TransformDepth() != 1 {
		t.Errorf(
			"In rotated group, TransformDepth() = %v, want 1",
			rc.TransformDepth(),
		)
	}

	// Exit
	sgc.ExitShapeGroup()

	if rc.TransformDepth() != 0 {
		t.Errorf(
			"After exit, TransformDepth() = %v, want 0",
			rc.TransformDepth(),
		)
	}
}

func TestHeaderContext_Basic(t *testing.T) {
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(NewMargins(72, 72, 72, 72))
	rc := NewRenderingContextFromPageOptions(opts)

	// Enter header
	hfc := rc.EnterHeader()

	if !hfc.IsHeader() {
		t.Error(
			"EnterHeader should return a header context",
		)
	}

	bounds := hfc.GetBounds()
	// Header should be in the top margin area (which in PDF is high Y)
	if bounds.Height() == 0 {
		t.Error(
			"Header bounds should have non-zero height",
		)
	}

	// Exit
	hfc.ExitContentArea()
}

func TestFooterContext_Basic(t *testing.T) {
	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(NewMargins(72, 72, 72, 72))
	rc := NewRenderingContextFromPageOptions(opts)

	// Enter footer
	hfc := rc.EnterFooter()

	if hfc.IsHeader() {
		t.Error(
			"EnterFooter should return a footer context",
		)
	}

	bounds := hfc.GetBounds()
	// Footer should have non-zero dimensions
	if bounds.Height() == 0 {
		t.Error(
			"Footer bounds should have non-zero height",
		)
	}

	// Exit
	hfc.ExitContentArea()
}

func TestRenderingContext_GetCoordTransformer(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	ct := rc.GetCoordTransformer()
	if ct == nil {
		t.Error(
			"GetCoordTransformer() should not return nil",
		)
	}
	if ct.PageHeight() != 792 {
		t.Errorf(
			"CoordTransformer page height = %v, want 792",
			ct.PageHeight(),
		)
	}
}

func TestRenderingContext_NestedContentAreas(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Enter first content area (simulating a table)
	tableCtx := rc.EnterContentArea(
		50,
		100,
		500,
		600,
	)

	// Enter a cell within the table
	cellCtx := rc.EnterContentArea(
		10,
		20,
		100,
		30,
	)

	// Origin should be cumulative
	origin := rc.CurrentOrigin()
	if origin.X != 60 || origin.Y != 120 {
		t.Errorf(
			"Nested origin = %v, want (60, 120)",
			origin,
		)
	}

	// Content bounds should be nested
	if rc.ContentBoundsDepth() != 2 {
		t.Errorf(
			"ContentBoundsDepth() = %v, want 2",
			rc.ContentBoundsDepth(),
		)
	}

	// Exit cell
	cellCtx.ExitContentArea()

	// Exit table
	tableCtx.ExitContentArea()

	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After all exits, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
}

func TestRenderingContext_TransformPointPDF(
	t *testing.T,
) {
	rc := NewRenderingContext(612, 792)

	// Push a scale transform
	rc.PushTransform(ScaleMatrix(2, 2))

	// TransformPointPDF should apply only the matrix, not origin or Y-flip
	x, y := rc.TransformPointPDF(100, 100)
	if x != 200 || y != 200 {
		t.Errorf(
			"TransformPointPDF(100, 100) with 2x scale = (%v, %v), want (200, 200)",
			x,
			y,
		)
	}
}

func TestIntersectRectangles(t *testing.T) {
	// Two overlapping rectangles
	a := NewRectangle(0, 0, 100, 100)
	b := NewRectangle(50, 50, 150, 150)

	result := intersectRectangles(a, b)

	if result.LLX != 50 || result.LLY != 50 {
		t.Errorf(
			"Intersection LL = (%v, %v), want (50, 50)",
			result.LLX,
			result.LLY,
		)
	}
	if result.URX != 100 || result.URY != 100 {
		t.Errorf(
			"Intersection UR = (%v, %v), want (100, 100)",
			result.URX,
			result.URY,
		)
	}
}

func TestIntersectRectangles_NoOverlap(
	t *testing.T,
) {
	// Non-overlapping rectangles
	a := NewRectangle(0, 0, 50, 50)
	b := NewRectangle(100, 100, 150, 150)

	result := intersectRectangles(a, b)

	// Result should be degenerate (zero size)
	if result.Width() > 0 || result.Height() > 0 {
		t.Errorf(
			"Non-overlapping intersection should be zero-sized, got %v",
			result,
		)
	}
}

func TestRenderingContext_ComplexScenario(
	t *testing.T,
) {
	// Simulate rendering a document with:
	// - Page with margins
	// - A table
	// - A cell within the table
	// - A shape within the cell

	opts := NewPageOptions(
		PageSizeLetter,
	).WithMargins(NewMarginsUniform(72))
	rc := NewRenderingContextFromPageOptions(opts)

	// Verify initial state
	if rc.OriginDepth() != 0 {
		t.Errorf(
			"Initial OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}

	// Enter table at (0, 0) relative to page origin, 468x300 points
	// Note: Margins affect content bounds but not origin stack
	// The origin stack tracks OOXML space offsets
	tableState := rc.SaveState()
	rc.PushOrigin(0, 0)
	tableBounds := rc.TransformRect(
		0,
		0,
		468,
		300,
	)
	rc.PushContentBounds(tableBounds)

	// Enter cell at (10, 10) relative to table, 100x50 points
	rc.PushOrigin(10, 10)
	cellBounds := rc.TransformRect(0, 0, 100, 50)
	rc.PushContentBounds(cellBounds)

	// Transform a point within the cell
	// Cell is at (10, 10) relative to table, table is at (0, 0) in OOXML space
	// So point (5, 5) in cell is at (0 + 10 + 5, 0 + 10 + 5) = (15, 15) in OOXML
	// PDF Y = 792 - 15 = 777
	x, y := rc.TransformPoint(5, 5)
	if !floatEqualsWithTolerance(x, 15, 0.001) {
		t.Errorf(
			"Complex scenario X = %v, want 15",
			x,
		)
	}
	if !floatEqualsWithTolerance(y, 777, 0.001) {
		t.Errorf(
			"Complex scenario Y = %v, want 777",
			y,
		)
	}

	// Restore to after table creation
	rc.RestoreState(tableState)

	// Should be back to initial state
	if rc.OriginDepth() != 0 {
		t.Errorf(
			"After restore, OriginDepth() = %v, want 0",
			rc.OriginDepth(),
		)
	}
}

func BenchmarkRenderingContext_TransformPoint(
	b *testing.B,
) {
	rc := NewRenderingContext(612, 792)
	rc.PushOrigin(100, 100)

	for range b.N {
		_, _ = rc.TransformPoint(50, 50)
	}
}

func BenchmarkRenderingContext_TransformRect(
	b *testing.B,
) {
	rc := NewRenderingContext(612, 792)
	rc.PushOrigin(100, 100)

	for range b.N {
		_ = rc.TransformRect(50, 50, 200, 100)
	}
}

func BenchmarkRenderingContext_PushPopOrigin(
	b *testing.B,
) {
	rc := NewRenderingContext(612, 792)

	for range b.N {
		rc.PushOrigin(100, 100)
		rc.PopOrigin()
	}
}

func BenchmarkRenderingContext_SaveRestoreState(
	b *testing.B,
) {
	rc := NewRenderingContext(612, 792)
	rc.PushOrigin(100, 100)
	rc.PushTransform(TranslationMatrix(50, 50))

	for range b.N {
		state := rc.SaveState()
		rc.PushOrigin(50, 50)
		rc.RestoreState(state)
	}
}
