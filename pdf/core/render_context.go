// Package core provides PDF document creation and low-level PDF primitives.
package core

import "math"

// RenderingContext manages the rendering state for converting OOXML content to PDF.
// It handles origin transformations, nested coordinate contexts, content bounds,
// and the transformation stack necessary for correctly positioning elements.
//
// OOXML uses a top-left origin coordinate system while PDF uses bottom-left.
// This context manages that transformation along with nested contexts like:
// - Page content areas (accounting for margins)
// - Table cells
// - Shapes within groups
// - Header/footer regions
type RenderingContext struct {
	// pageWidth is the page width in PDF points.
	pageWidth float64
	// pageHeight is the page height in PDF points.
	pageHeight float64

	// originStack maintains a stack of origin offsets for nested contexts.
	// Each entry represents an (x, y) offset from the page origin.
	originStack []Point

	// transformStack maintains a stack of transformation matrices.
	transformStack []PDFMatrix

	// contentBoundsStack maintains a stack of content bounds for clipping.
	contentBoundsStack []Rectangle

	// currentMargins stores the current page margins.
	currentMargins Margins

	// coordTransformer handles the base coordinate transformation.
	coordTransformer *CoordTransformer
}

// Point represents a 2D point with X and Y coordinates.
type Point struct {
	X, Y float64
}

// NewPoint creates a new Point with the given coordinates.
func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// Add returns a new point that is the sum of this point and another.
func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Sub returns a new point that is the difference of this point and another.
func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// Scale returns a new point scaled by the given factor.
func (p Point) Scale(factor float64) Point {
	return Point{X: p.X * factor, Y: p.Y * factor}
}

// Distance returns the distance between this point and another.
func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y

	return math.Sqrt(dx*dx + dy*dy)
}

// NewRenderingContext creates a new rendering context for a page.
// pageWidth and pageHeight should be in PDF points.
func NewRenderingContext(
	pageWidth, pageHeight float64,
) *RenderingContext {
	rc := &RenderingContext{
		pageWidth:   pageWidth,
		pageHeight:  pageHeight,
		originStack: make([]Point, 0, 8),
		transformStack: make(
			[]PDFMatrix,
			0,
			8,
		),
		contentBoundsStack: make(
			[]Rectangle,
			0,
			8,
		),
		coordTransformer: NewCoordTransformer(
			pageHeight,
		),
	}

	// Initialize with identity at the root level
	rc.originStack = append(
		rc.originStack,
		NewPoint(0, 0),
	)
	rc.transformStack = append(
		rc.transformStack,
		IdentityMatrix(),
	)
	rc.contentBoundsStack = append(
		rc.contentBoundsStack,
		NewRectangleFromSize(
			0,
			0,
			pageWidth,
			pageHeight,
		),
	)

	return rc
}

// NewRenderingContextFromPageSize creates a rendering context from a PageSize.
func NewRenderingContextFromPageSize(
	size PageSize,
) *RenderingContext {
	return NewRenderingContext(
		size.Width,
		size.Height,
	)
}

// NewRenderingContextFromPageOptions creates a rendering context from page options.
// This automatically accounts for margins in the initial content bounds.
func NewRenderingContextFromPageOptions(
	opts *PageOptions,
) *RenderingContext {
	rc := NewRenderingContext(
		opts.Size.Width,
		opts.Size.Height,
	)
	rc.currentMargins = opts.Margins

	// Set initial content bounds to the content area (page minus margins)
	contentArea := opts.ContentArea()
	rc.contentBoundsStack[0] = contentArea

	return rc
}

// GetPageWidth returns the page width in PDF points.
func (rc *RenderingContext) GetPageWidth() float64 {
	return rc.pageWidth
}

// GetPageHeight returns the page height in PDF points.
func (rc *RenderingContext) GetPageHeight() float64 {
	return rc.pageHeight
}

// GetPageSize returns the page size.
func (rc *RenderingContext) GetPageSize() PageSize {
	return PageSize{
		Width:  rc.pageWidth,
		Height: rc.pageHeight,
	}
}

// GetMargins returns the current page margins.
func (rc *RenderingContext) GetMargins() Margins {
	return rc.currentMargins
}

// SetMargins sets the current page margins and updates the root content bounds.
func (rc *RenderingContext) SetMargins(
	margins Margins,
) {
	rc.currentMargins = margins
	if len(rc.contentBoundsStack) > 0 {
		rc.contentBoundsStack[0] = Rectangle{
			LLX: margins.Left,
			LLY: margins.Bottom,
			URX: rc.pageWidth - margins.Right,
			URY: rc.pageHeight - margins.Top,
		}
	}
}

// PushOrigin pushes a new origin offset onto the stack.
// All subsequent coordinate transformations will be relative to this origin.
// The offset (x, y) is in OOXML coordinates (top-left origin).
// This is useful for rendering content within:
// - Table cells
// - Shape groups
// - Header/footer regions
func (rc *RenderingContext) PushOrigin(
	x, y float64,
) {
	currentOrigin := rc.CurrentOrigin()
	newOrigin := Point{
		X: currentOrigin.X + x,
		Y: currentOrigin.Y + y,
	}
	rc.originStack = append(
		rc.originStack,
		newOrigin,
	)
}

// PushOriginPDF pushes a new origin offset using PDF coordinates (bottom-left).
// This is useful when you have coordinates already in PDF space.
func (rc *RenderingContext) PushOriginPDF(
	x, y float64,
) {
	currentOrigin := rc.CurrentOrigin()
	// Convert from PDF Y to OOXML Y for storage
	ooxmlY := rc.pageHeight - y
	newOrigin := Point{
		X: currentOrigin.X + x,
		Y: currentOrigin.Y + (rc.pageHeight - ooxmlY),
	}
	rc.originStack = append(
		rc.originStack,
		newOrigin,
	)
}

// PopOrigin removes the most recent origin offset from the stack.
// It returns false if trying to pop the root origin.
func (rc *RenderingContext) PopOrigin() bool {
	if len(rc.originStack) <= 1 {
		return false
	}
	rc.originStack = rc.originStack[:len(rc.originStack)-1]

	return true
}

// CurrentOrigin returns the current cumulative origin offset.
func (rc *RenderingContext) CurrentOrigin() Point {
	if len(rc.originStack) == 0 {
		return Point{X: 0, Y: 0}
	}

	return rc.originStack[len(rc.originStack)-1]
}

// OriginDepth returns the current nesting depth of origins (0 = root).
func (rc *RenderingContext) OriginDepth() int {
	return len(rc.originStack) - 1
}

// PushTransform pushes a transformation matrix onto the stack.
// This combines with the current transformation.
func (rc *RenderingContext) PushTransform(
	matrix PDFMatrix,
) {
	current := rc.CurrentTransform()
	combined := current.Multiply(matrix)
	rc.transformStack = append(
		rc.transformStack,
		combined,
	)
}

// PushOOXMLTransform pushes an OOXML transform onto the stack.
// It handles the conversion from OOXML to PDF coordinate space.
func (rc *RenderingContext) PushOOXMLTransform(
	xfrm OOXMLTransform,
) {
	matrix := OOXMLTransformToPDFMatrix(
		xfrm,
		rc.pageHeight,
	)
	rc.PushTransform(matrix)
}

// PopTransform removes the most recent transformation from the stack.
// Returns false if trying to pop the root transformation.
func (rc *RenderingContext) PopTransform() bool {
	if len(rc.transformStack) <= 1 {
		return false
	}
	rc.transformStack = rc.transformStack[:len(rc.transformStack)-1]

	return true
}

// CurrentTransform returns the current combined transformation matrix.
func (rc *RenderingContext) CurrentTransform() PDFMatrix {
	if len(rc.transformStack) == 0 {
		return IdentityMatrix()
	}

	return rc.transformStack[len(rc.transformStack)-1]
}

// TransformDepth returns the current nesting depth of transforms (0 = root).
func (rc *RenderingContext) TransformDepth() int {
	return len(rc.transformStack) - 1
}

// SetContentBounds sets the current content bounds for clipping.
// The bounds are in PDF coordinates (bottom-left origin).
func (rc *RenderingContext) SetContentBounds(
	bounds Rectangle,
) {
	if len(rc.contentBoundsStack) > 0 {
		rc.contentBoundsStack[len(rc.contentBoundsStack)-1] = bounds
	}
}

// PushContentBounds pushes a new content bounds onto the stack.
// This is useful for nested clipping regions like table cells or shapes.
// The bounds are in PDF coordinates.
func (rc *RenderingContext) PushContentBounds(
	bounds Rectangle,
) {
	// Intersect with current bounds if needed
	current := rc.GetContentBounds()
	intersected := intersectRectangles(
		current,
		bounds,
	)
	rc.contentBoundsStack = append(
		rc.contentBoundsStack,
		intersected,
	)
}

// PopContentBounds removes the most recent content bounds from the stack.
// Returns false if trying to pop the root bounds.
func (rc *RenderingContext) PopContentBounds() bool {
	if len(rc.contentBoundsStack) <= 1 {
		return false
	}
	rc.contentBoundsStack = rc.contentBoundsStack[:len(rc.contentBoundsStack)-1]

	return true
}

// GetContentBounds returns the current content bounds.
func (rc *RenderingContext) GetContentBounds() Rectangle {
	if len(rc.contentBoundsStack) == 0 {
		return NewRectangleFromSize(
			0,
			0,
			rc.pageWidth,
			rc.pageHeight,
		)
	}

	return rc.contentBoundsStack[len(rc.contentBoundsStack)-1]
}

// ContentBoundsDepth returns the current nesting depth of content bounds (0 = root).
func (rc *RenderingContext) ContentBoundsDepth() int {
	return len(rc.contentBoundsStack) - 1
}

// TransformPoint transforms a point from OOXML coordinates (top-left origin)
// to PDF coordinates (bottom-left origin), applying the current origin offset
// and transformation matrix.
func (rc *RenderingContext) TransformPoint(
	x, y float64,
) (pdfX, pdfY float64) {
	// Apply origin offset
	origin := rc.CurrentOrigin()
	ooxmlX := x + origin.X
	ooxmlY := y + origin.Y

	// Convert to PDF coordinates (flip Y axis)
	outX := ooxmlX
	outY := rc.pageHeight - ooxmlY

	// Apply current transformation matrix
	transform := rc.CurrentTransform()
	if !transform.IsIdentity() {
		outX, outY = transform.TransformPoint(
			outX,
			outY,
		)
	}

	return outX, outY
}

// TransformPointPDF transforms a point that is already in PDF coordinates
// by applying the current transformation matrix only.
func (rc *RenderingContext) TransformPointPDF(
	x, y float64,
) (pdfX, pdfY float64) {
	transform := rc.CurrentTransform()

	return transform.TransformPoint(x, y)
}

// TransformPointEMU transforms a point from OOXML EMU coordinates to PDF coordinates.
func (rc *RenderingContext) TransformPointEMU(
	emuX, emuY int64,
) (pdfX, pdfY float64) {
	// Convert EMU to points first
	ptX := float64(emuX) * PointsPerEMU
	ptY := float64(emuY) * PointsPerEMU

	// Then apply standard transformation
	return rc.TransformPoint(ptX, ptY)
}

// TransformRect transforms a rectangle from OOXML coordinates to PDF coordinates.
// The input rectangle should be in OOXML space (top-left origin, x/y/width/height).
// Returns a Rectangle in PDF space (bottom-left origin).
func (rc *RenderingContext) TransformRect(
	x, y, width, height float64,
) Rectangle {
	// Transform the top-left corner of the rectangle
	pdfX1, pdfY1 := rc.TransformPoint(x, y)
	// Transform the bottom-right corner
	pdfX2, pdfY2 := rc.TransformPoint(
		x+width,
		y+height,
	)

	// Create rectangle with proper PDF orientation (lower-left to upper-right)
	// Since we flipped Y, the top-left in OOXML becomes somewhere higher in PDF
	llx := math.Min(pdfX1, pdfX2)
	lly := math.Min(pdfY1, pdfY2)
	urx := math.Max(pdfX1, pdfX2)
	ury := math.Max(pdfY1, pdfY2)

	return NewRectangle(llx, lly, urx, ury)
}

// TransformRectEMU transforms a rectangle from OOXML EMU coordinates to PDF coordinates.
func (rc *RenderingContext) TransformRectEMU(
	emuX, emuY, emuWidth, emuHeight int64,
) Rectangle {
	x := float64(emuX) * PointsPerEMU
	y := float64(emuY) * PointsPerEMU
	w := float64(emuWidth) * PointsPerEMU
	h := float64(emuHeight) * PointsPerEMU

	return rc.TransformRect(x, y, w, h)
}

// TransformSize transforms a size (width, height) applying only scaling.
// Unlike TransformPoint, this does not apply origin offset or Y-flip.
func (rc *RenderingContext) TransformSize(
	width, height float64,
) (float64, float64) {
	transform := rc.CurrentTransform()

	return transform.TransformVector(
		width,
		height,
	)
}

// TransformSizeEMU transforms a size from EMU to PDF points.
func (rc *RenderingContext) TransformSizeEMU(
	emuWidth, emuHeight int64,
) (float64, float64) {
	width := float64(emuWidth) * PointsPerEMU
	height := float64(emuHeight) * PointsPerEMU

	return rc.TransformSize(width, height)
}

// IsPointInContentBounds checks if a point (in PDF coordinates) is within the current content bounds.
func (rc *RenderingContext) IsPointInContentBounds(
	x, y float64,
) bool {
	bounds := rc.GetContentBounds()

	return bounds.Contains(x, y)
}

// IsRectInContentBounds checks if a rectangle (in PDF coordinates) is fully within the content bounds.
func (rc *RenderingContext) IsRectInContentBounds(
	rect Rectangle,
) bool {
	bounds := rc.GetContentBounds()

	return rect.LLX >= bounds.LLX &&
		rect.LLY >= bounds.LLY &&
		rect.URX <= bounds.URX &&
		rect.URY <= bounds.URY
}

// RectIntersectsContentBounds checks if a rectangle intersects with the content bounds.
func (rc *RenderingContext) RectIntersectsContentBounds(
	rect Rectangle,
) bool {
	bounds := rc.GetContentBounds()

	return !(rect.URX < bounds.LLX || rect.LLX > bounds.URX ||
		rect.URY < bounds.LLY || rect.LLY > bounds.URY)
}

// ClipToContentBounds clips a rectangle to the current content bounds.
// Returns the clipped rectangle and whether any portion remains visible.
func (rc *RenderingContext) ClipToContentBounds(
	rect Rectangle,
) (Rectangle, bool) {
	bounds := rc.GetContentBounds()
	clipped := intersectRectangles(rect, bounds)

	// Check if the clipped rectangle has valid dimensions
	valid := clipped.Width() > 0 &&
		clipped.Height() > 0

	return clipped, valid
}

// Reset clears all stacks and resets to the initial state.
func (rc *RenderingContext) Reset() {
	rc.originStack = rc.originStack[:1]
	rc.originStack[0] = NewPoint(0, 0)

	rc.transformStack = rc.transformStack[:1]
	rc.transformStack[0] = IdentityMatrix()

	rc.contentBoundsStack = rc.contentBoundsStack[:1]
	rc.contentBoundsStack[0] = NewRectangleFromSize(
		0,
		0,
		rc.pageWidth,
		rc.pageHeight,
	)
	if !rc.currentMargins.IsZero() {
		rc.SetMargins(rc.currentMargins)
	}
}

// SaveState saves the current state of all stacks.
// Returns a state token that can be used to restore.
func (rc *RenderingContext) SaveState() RenderState {
	return RenderState{
		originDepth:    len(rc.originStack),
		transformDepth: len(rc.transformStack),
		boundsDepth: len(
			rc.contentBoundsStack,
		),
	}
}

// RestoreState restores the stacks to a previously saved state.
func (rc *RenderingContext) RestoreState(
	state RenderState,
) {
	if state.originDepth > 0 &&
		state.originDepth <= len(rc.originStack) {
		rc.originStack = rc.originStack[:state.originDepth]
	}
	if state.transformDepth > 0 &&
		state.transformDepth <= len(
			rc.transformStack,
		) {
		rc.transformStack = rc.transformStack[:state.transformDepth]
	}
	if state.boundsDepth > 0 &&
		state.boundsDepth <= len(
			rc.contentBoundsStack,
		) {
		rc.contentBoundsStack = rc.contentBoundsStack[:state.boundsDepth]
	}
}

// RenderState represents a saved rendering state.
type RenderState struct {
	originDepth    int
	transformDepth int
	boundsDepth    int
}

// GetCoordTransformer returns the underlying coordinate transformer.
// This provides access to lower-level transformation functions.
func (rc *RenderingContext) GetCoordTransformer() *CoordTransformer {
	return rc.coordTransformer
}

// ContentAreaContext represents a context for rendering within a specific content area.
// This is a convenience wrapper for common nested context patterns.
type ContentAreaContext struct {
	rc            *RenderingContext
	savedState    RenderState
	contentBounds Rectangle
}

// EnterContentArea creates a new content area context and pushes the necessary state.
// The bounds should be in OOXML coordinates (top-left origin).
// Use ExitContentArea when done rendering in this area.
func (rc *RenderingContext) EnterContentArea(
	x, y, width, height float64,
) *ContentAreaContext {
	savedState := rc.SaveState()

	// Push origin for this content area
	rc.PushOrigin(x, y)

	// Transform bounds to PDF coordinates and push
	pdfBounds := rc.TransformRect(
		0,
		0,
		width,
		height,
	)
	rc.PushContentBounds(pdfBounds)

	return &ContentAreaContext{
		rc:            rc,
		savedState:    savedState,
		contentBounds: pdfBounds,
	}
}

// ExitContentArea restores the rendering context to the state before entering.
func (c *ContentAreaContext) ExitContentArea() {
	c.rc.RestoreState(c.savedState)
}

// GetBounds returns the bounds of this content area in PDF coordinates.
func (c *ContentAreaContext) GetBounds() Rectangle {
	return c.contentBounds
}

// TableCellContext represents a context for rendering within a table cell.
type TableCellContext struct {
	*ContentAreaContext
	row, col   int
	cellWidth  float64
	cellHeight float64
}

// EnterTableCell creates a new table cell context.
// The position and size should be in points (OOXML space).
func (rc *RenderingContext) EnterTableCell(
	x, y, width, height float64,
	row, col int,
) *TableCellContext {
	cac := rc.EnterContentArea(
		x,
		y,
		width,
		height,
	)

	return &TableCellContext{
		ContentAreaContext: cac,
		row:                row,
		col:                col,
		cellWidth:          width,
		cellHeight:         height,
	}
}

// GetCellPosition returns the row and column of this cell.
func (t *TableCellContext) GetCellPosition() (row, col int) {
	return t.row, t.col
}

// GetCellSize returns the width and height of this cell in points.
func (t *TableCellContext) GetCellSize() (width, height float64) {
	return t.cellWidth, t.cellHeight
}

// ShapeGroupContext represents a context for rendering within a shape group.
type ShapeGroupContext struct {
	*ContentAreaContext
	groupTransform OOXMLTransform
}

// EnterShapeGroup creates a new shape group context.
// The groupTransform should be the group's own transformation from the OOXML.
func (rc *RenderingContext) EnterShapeGroup(
	groupTransform OOXMLTransform,
) *ShapeGroupContext {
	// Get the group's position and size in points
	offX, offY := groupTransform.OffsetPoints()
	extCx, extCy := groupTransform.ExtentPoints()

	cac := rc.EnterContentArea(
		offX,
		offY,
		extCx,
		extCy,
	)

	// Push the group's transformation (rotation, flips)
	if groupTransform.HasTransform() {
		rc.PushOOXMLTransform(groupTransform)
	}

	return &ShapeGroupContext{
		ContentAreaContext: cac,
		groupTransform:     groupTransform,
	}
}

// ExitShapeGroup restores the rendering context.
func (s *ShapeGroupContext) ExitShapeGroup() {
	if s.groupTransform.HasTransform() {
		s.rc.PopTransform()
	}
	s.ExitContentArea()
}

// HeaderFooterContext represents a context for rendering headers or footers.
type HeaderFooterContext struct {
	*ContentAreaContext
	isHeader bool
}

// EnterHeader creates a context for rendering in the page header area.
// The header area is defined by the top margin.
func (rc *RenderingContext) EnterHeader() *HeaderFooterContext {
	margins := rc.GetMargins()
	width := rc.pageWidth - margins.Left - margins.Right
	height := margins.Top

	cac := rc.EnterContentArea(
		margins.Left,
		0,
		width,
		height,
	)

	return &HeaderFooterContext{
		ContentAreaContext: cac,
		isHeader:           true,
	}
}

// EnterFooter creates a context for rendering in the page footer area.
// The footer area is defined by the bottom margin.
func (rc *RenderingContext) EnterFooter() *HeaderFooterContext {
	margins := rc.GetMargins()
	width := rc.pageWidth - margins.Left - margins.Right
	height := margins.Bottom

	// Footer is at the bottom of the page in OOXML coordinates
	footerY := rc.pageHeight - margins.Bottom

	cac := rc.EnterContentArea(
		margins.Left,
		footerY,
		width,
		height,
	)

	return &HeaderFooterContext{
		ContentAreaContext: cac,
		isHeader:           false,
	}
}

// IsHeader returns true if this is a header context, false for footer.
func (h *HeaderFooterContext) IsHeader() bool {
	return h.isHeader
}

// intersectRectangles computes the intersection of two rectangles.
// If they don't intersect, returns a zero-sized rectangle.
func intersectRectangles(
	a, b Rectangle,
) Rectangle {
	llx := math.Max(a.LLX, b.LLX)
	lly := math.Max(a.LLY, b.LLY)
	urx := math.Min(a.URX, b.URX)
	ury := math.Min(a.URY, b.URY)

	// If no intersection, return a degenerate rectangle
	if llx >= urx || lly >= ury {
		return Rectangle{
			LLX: llx,
			LLY: lly,
			URX: llx,
			URY: lly,
		}
	}

	return Rectangle{
		LLX: llx,
		LLY: lly,
		URX: urx,
		URY: ury,
	}
}
