
package drawing

import (
	"strings"
)

// ClipRule defines the fill rule used for clipping operations.
type ClipRule int

const (
	// ClipNonZeroRule uses the non-zero winding number rule to determine
	// which points are inside the clipping path.
	ClipNonZeroRule ClipRule = iota

	// ClipEvenOddRule uses the even-odd rule to determine which points
	// are inside the clipping path.
	ClipEvenOddRule
)

// String returns the PDF operator for the clip rule.
func (r ClipRule) String() string {
	switch r {
	case ClipNonZeroRule:
		return "W"
	case ClipEvenOddRule:
		return "W*"
	default:
		return "W"
	}
}

// ClipContext manages clipping paths with proper graphics state nesting.
// It uses the PDF graphics state stack (q/Q operators) to ensure that
// clipping regions can be properly restored.
//
// Clipping in PDF is intersective - each new clipping path intersects
// with the current clipping path. To restore a previous clipping region,
// you must use the graphics state stack.
//
// Example usage:
//
//	clip := drawing.NewClipContext()
//	clip.PushClip(drawing.ClipRect(10, 10, 100, 100), drawing.ClipNonZeroRule)
//	// ... draw content within clipped region ...
//	clip.PopClip()
//	content := clip.String()
type ClipContext struct {
	operators  strings.Builder
	clipDepth  int
	transforms []Transform
}

// NewClipContext creates a new clipping context.
func NewClipContext() *ClipContext {
	return &ClipContext{
		transforms: []Transform{Identity()},
	}
}

// PushClip saves the graphics state and applies a clipping path.
// The clipping path is defined by the given PathBuilder.
// Use PopClip to restore the previous clipping state.
//
// This outputs:
//   - q (save graphics state)
//   - path operators
//   - W or W* (clip operator)
//   - n (end path without painting)
func (c *ClipContext) PushClip(
	path *PathBuilder,
	rule ClipRule,
) *ClipContext {
	// Save graphics state
	c.operators.WriteString("q\n")
	c.clipDepth++

	// Save current transform
	if len(c.transforms) > 0 {
		c.transforms = append(
			c.transforms,
			c.transforms[len(c.transforms)-1],
		)
	} else {
		c.transforms = append(c.transforms, Identity())
	}

	// Write path operators
	c.operators.WriteString(path.String())

	// Apply clipping
	c.operators.WriteString(rule.String())
	c.operators.WriteString(" n\n")

	return c
}

// PushClipWithTransform saves graphics state, applies a transform, then applies clipping.
// This is useful for applying a clipping path in a transformed coordinate system.
func (c *ClipContext) PushClipWithTransform(
	path *PathBuilder,
	rule ClipRule,
	transform Transform,
) *ClipContext {
	// Save graphics state
	c.operators.WriteString("q\n")
	c.clipDepth++

	// Save and update current transform
	currentTransform := Identity()
	if len(c.transforms) > 0 {
		currentTransform = c.transforms[len(c.transforms)-1]
	}
	newTransform := currentTransform.Concat(
		transform,
	)
	c.transforms = append(
		c.transforms,
		newTransform,
	)

	// Apply transformation
	if !transform.IsIdentity() {
		c.operators.WriteString(
			transform.ToCMOperator(),
		)
		c.operators.WriteByte('\n')
	}

	// Write path operators
	c.operators.WriteString(path.String())

	// Apply clipping
	c.operators.WriteString(rule.String())
	c.operators.WriteString(" n\n")

	return c
}

// PopClip restores the graphics state to before the last PushClip.
// This restores the previous clipping region.
func (c *ClipContext) PopClip() *ClipContext {
	if c.clipDepth > 0 {
		c.operators.WriteString("Q\n")
		c.clipDepth--

		// Restore transform
		if len(c.transforms) > 1 {
			c.transforms = c.transforms[:len(c.transforms)-1]
		}
	}

	return c
}

// PopAllClips restores all saved graphics states.
func (c *ClipContext) PopAllClips() *ClipContext {
	for c.clipDepth > 0 {
		c.PopClip()
	}

	return c
}

// CurrentTransform returns the current transformation matrix.
func (c *ClipContext) CurrentTransform() Transform {
	if len(c.transforms) > 0 {
		return c.transforms[len(c.transforms)-1]
	}

	return Identity()
}

// Depth returns the current clipping depth (number of active clips).
func (c *ClipContext) Depth() int {
	return c.clipDepth
}

// WriteOperator writes a raw operator to the content stream.
// This allows adding drawing operations within the clipping context.
func (c *ClipContext) WriteOperator(
	op string,
) *ClipContext {
	c.operators.WriteString(op)
	c.operators.WriteByte('\n')

	return c
}

// WriteContent writes arbitrary content to the stream.
func (c *ClipContext) WriteContent(
	content string,
) *ClipContext {
	c.operators.WriteString(content)

	return c
}

// String returns the accumulated operators.
func (c *ClipContext) String() string {
	return c.operators.String()
}

// Reset clears the clipping context.
func (c *ClipContext) Reset() *ClipContext {
	c.operators.Reset()
	c.clipDepth = 0
	c.transforms = []Transform{Identity()}

	return c
}

// ClipScope provides automatic cleanup for clipping regions using Go's defer.
// It ensures that PopClip is called when the scope exits.
//
// Example usage:
//
//	func drawWithClip(clip *ClipContext) {
//	    scope := NewClipScope(clip, ClipRect(0, 0, 100, 100), ClipNonZeroRule)
//	    defer scope.Close()
//	    // ... draw content within clipped region ...
//	    // PopClip is automatically called when function exits
//	}
type ClipScope struct {
	context *ClipContext
	closed  bool
}

// NewClipScope creates a new clipping scope.
// It immediately pushes a clip and will pop it when Close is called.
func NewClipScope(
	ctx *ClipContext,
	path *PathBuilder,
	rule ClipRule,
) *ClipScope {
	ctx.PushClip(path, rule)

	return &ClipScope{
		context: ctx,
		closed:  false,
	}
}

// NewClipScopeWithTransform creates a clipping scope with a transformation.
func NewClipScopeWithTransform(
	ctx *ClipContext,
	path *PathBuilder,
	rule ClipRule,
	transform Transform,
) *ClipScope {
	ctx.PushClipWithTransform(
		path,
		rule,
		transform,
	)

	return &ClipScope{
		context: ctx,
		closed:  false,
	}
}

// Close ends the clipping scope by calling PopClip.
// It is safe to call Close multiple times; only the first call has effect.
func (s *ClipScope) Close() {
	if !s.closed {
		s.context.PopClip()
		s.closed = true
	}
}

// Clipping path helper functions
// These create PathBuilder instances configured as clipping paths.

// ClipRect creates a rectangular clipping path.
// The rectangle has its lower-left corner at (x, y) with the given width and height.
func ClipRect(
	x, y, width, height float64,
) *PathBuilder {
	return NewPathBuilder().Rectangle(x, y, width, height)
}

// ClipRoundedRect creates a rounded rectangle clipping path.
// The rectangle has its lower-left corner at (x, y) with the given dimensions
// and uniform corner radius.
func ClipRoundedRect(
	x, y, width, height, radius float64,
) *PathBuilder {
	return NewPathBuilder().RoundedRect(x, y, width, height, radius)
}

// ClipRoundedRectVarying creates a rounded rectangle with varying corner radii.
// The radii are specified in order: top-left, top-right, bottom-right, bottom-left.
func ClipRoundedRectVarying(
	x, y, width, height, tl, tr, br, bl float64,
) *PathBuilder {
	return NewPathBuilder().RoundedRectVarying(x, y, width, height, tl, tr, br, bl)
}

// ClipCircle creates a circular clipping path.
// The circle is centered at (cx, cy) with the given radius.
func ClipCircle(
	cx, cy, radius float64,
) *PathBuilder {
	return NewPathBuilder().Circle(cx, cy, radius)
}

// ClipEllipse creates an elliptical clipping path.
// The ellipse is centered at (cx, cy) with radii rx and ry.
func ClipEllipse(
	cx, cy, rx, ry float64,
) *PathBuilder {
	return NewPathBuilder().Ellipse(cx, cy, rx, ry)
}

// ClipPath creates a clipping path from an existing PathBuilder.
// This creates a copy of the path that can be used for clipping.
func ClipPath(path *PathBuilder) *PathBuilder {
	// Create a new PathBuilder with the same content
	newPath := NewPathBuilder()
	newPath.operators.WriteString(path.String())
	newPath.currentX = path.currentX
	newPath.currentY = path.currentY
	newPath.startX = path.startX
	newPath.startY = path.startY
	newPath.hasStart = path.hasStart
	newPath.subpathLen = path.subpathLen

	return newPath
}

// ClipPolygon creates a polygon clipping path from a list of points.
// Points are specified as alternating x, y coordinates.
func ClipPolygon(points ...float64) *PathBuilder {
	return NewPathBuilder().Polygon(points...)
}

// Convenience methods for directly generating clipping operators

// RectClipOperators returns the PDF operators for a rectangular clip.
// This includes the path, clip operator, and end path operator.
func RectClipOperators(
	x, y, width, height float64,
	rule ClipRule,
) string {
	var sb strings.Builder
	sb.WriteString(
		ClipRect(x, y, width, height).String(),
	)
	sb.WriteString(rule.String())
	sb.WriteString(" n\n")

	return sb.String()
}

// CircleClipOperators returns the PDF operators for a circular clip.
func CircleClipOperators(
	cx, cy, radius float64,
	rule ClipRule,
) string {
	var sb strings.Builder
	sb.WriteString(
		ClipCircle(cx, cy, radius).String(),
	)
	sb.WriteString(rule.String())
	sb.WriteString(" n\n")

	return sb.String()
}

// PathClipOperators returns the PDF operators for clipping to an arbitrary path.
func PathClipOperators(
	path *PathBuilder,
	rule ClipRule,
) string {
	var sb strings.Builder
	sb.WriteString(path.String())
	sb.WriteString(rule.String())
	sb.WriteString(" n\n")

	return sb.String()
}

// ScopedClipOperators returns the PDF operators for a complete clipping scope.
// This includes save, clip, content, and restore.
func ScopedClipOperators(
	path *PathBuilder,
	rule ClipRule,
	content string,
) string {
	var sb strings.Builder
	sb.WriteString("q\n")
	sb.WriteString(path.String())
	sb.WriteString(rule.String())
	sb.WriteString(" n\n")
	sb.WriteString(content)
	sb.WriteString("Q\n")

	return sb.String()
}

// ScopedClipWithTransformOperators returns operators for a clipping scope with transform.
func ScopedClipWithTransformOperators(
	path *PathBuilder,
	rule ClipRule,
	transform Transform,
	content string,
) string {
	var sb strings.Builder
	sb.WriteString("q\n")
	if !transform.IsIdentity() {
		sb.WriteString(transform.ToCMOperator())
		sb.WriteByte('\n')
	}
	sb.WriteString(path.String())
	sb.WriteString(rule.String())
	sb.WriteString(" n\n")
	sb.WriteString(content)
	sb.WriteString("Q\n")

	return sb.String()
}

// IntersectClip creates a path that represents the intersection of two clipping paths.
// Note: This is primarily for documentation - in PDF, clipping is always intersective
// when applied sequentially. This function simply combines the paths.
func IntersectClip(
	path1, path2 *PathBuilder,
) *PathBuilder {
	newPath := NewPathBuilder()
	newPath.operators.WriteString(path1.String())
	newPath.operators.WriteString(path2.String())

	return newPath
}

// NestedClipOperators creates operators for nested clipping regions.
// Each level adds a new clipping region that intersects with the previous.
func NestedClipOperators(
	paths []*PathBuilder,
	rule ClipRule,
	content string,
) string {
	var sb strings.Builder

	// Push all clips
	for _, path := range paths {
		sb.WriteString("q\n")
		sb.WriteString(path.String())
		sb.WriteString(rule.String())
		sb.WriteString(" n\n")
	}

	// Write content
	sb.WriteString(content)

	// Pop all clips
	for range paths {
		sb.WriteString("Q\n")
	}

	return sb.String()
}
