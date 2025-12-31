// Package core provides PDF document creation and low-level PDF primitives.
package core

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// Orientation represents the page orientation.
type Orientation int

const (
	// OrientationPortrait is the default portrait orientation (height > width).
	OrientationPortrait Orientation = iota
	// OrientationLandscape is landscape orientation (width > height).
	OrientationLandscape
)

// String returns the string representation of the orientation.
func (o Orientation) String() string {
	if o == OrientationLandscape {
		return "landscape"
	}

	return "portrait"
}

// Additional common page sizes in PDF points (72 points = 1 inch).
// ISO A series sizes (mm).
var (
	// PageSizeA0 is the ISO A0 page size (841mm x 1189mm).
	PageSizeA0 = PageSize{
		Width:  MMToPoints(841),
		Height: MMToPoints(1189),
	}
	// PageSizeA1 is the ISO A1 page size (594mm x 841mm).
	PageSizeA1 = PageSize{
		Width:  MMToPoints(594),
		Height: MMToPoints(841),
	}
	// PageSizeA2 is the ISO A2 page size (420mm x 594mm).
	PageSizeA2 = PageSize{
		Width:  MMToPoints(420),
		Height: MMToPoints(594),
	}
	// PageSizeA5 is the ISO A5 page size (148mm x 210mm).
	PageSizeA5 = PageSize{
		Width:  MMToPoints(148),
		Height: MMToPoints(210),
	}
	// PageSizeA6 is the ISO A6 page size (105mm x 148mm).
	PageSizeA6 = PageSize{
		Width:  MMToPoints(105),
		Height: MMToPoints(148),
	}
)

// ISO B series sizes (mm).
var (
	// PageSizeB4 is the ISO B4 page size (250mm x 353mm).
	PageSizeB4 = PageSize{
		Width:  MMToPoints(250),
		Height: MMToPoints(353),
	}
	// PageSizeB5 is the ISO B5 page size (176mm x 250mm).
	PageSizeB5 = PageSize{
		Width:  MMToPoints(176),
		Height: MMToPoints(250),
	}
)

// US paper sizes.
var (
	// PageSizeExecutive is the US Executive page size (7.25in x 10.5in).
	PageSizeExecutive = PageSize{
		Width:  InchesToPoints(7.25),
		Height: InchesToPoints(10.5),
	}
	// PageSizeLedger is the US Ledger page size (17in x 11in) - landscape Tabloid.
	PageSizeLedger = PageSizeTabloid.Landscape()
	// PageSizeStatement is the US Statement page size (5.5in x 8.5in).
	PageSizeStatement = PageSize{
		Width:  InchesToPoints(5.5),
		Height: InchesToPoints(8.5),
	}
	// PageSizeFolio is the Folio page size (8.5in x 13in).
	PageSizeFolio = PageSize{
		Width:  InchesToPoints(8.5),
		Height: InchesToPoints(13),
	}
	// PageSizeQuarto is the Quarto page size (8.5in x 10.83in).
	PageSizeQuarto = PageSize{
		Width:  InchesToPoints(8.5),
		Height: InchesToPoints(10.83),
	}
)

// NewPageSize creates a custom page size from width and height in points.
func NewPageSize(width, height float64) PageSize {
	return PageSize{Width: width, Height: height}
}

// NewPageSizeFromDimensions creates a page size from Dimension values.
func NewPageSizeFromDimensions(
	width, height Dimension,
) PageSize {
	return PageSize{
		Width:  width.ToPoints(),
		Height: height.ToPoints(),
	}
}

// NewPageSizeFromUnit creates a page size with the given unit.
func NewPageSizeFromUnit(
	width, height float64,
	unit Unit,
) PageSize {
	return PageSize{
		Width:  unit.ToPoints(width),
		Height: unit.ToPoints(height),
	}
}

// WithOrientation returns the page size in the specified orientation.
func (ps PageSize) WithOrientation(
	orientation Orientation,
) PageSize {
	isCurrentlyLandscape := ps.Width > ps.Height
	wantLandscape := orientation == OrientationLandscape

	if isCurrentlyLandscape == wantLandscape {
		return ps
	}

	return ps.Landscape()
}

// Portrait returns the page size in portrait orientation (height > width).
func (ps PageSize) Portrait() PageSize {
	if ps.Width > ps.Height {
		return PageSize{
			Width:  ps.Height,
			Height: ps.Width,
		}
	}

	return ps
}

// IsPortrait returns true if the page size is in portrait orientation.
func (ps PageSize) IsPortrait() bool {
	return ps.Height >= ps.Width
}

// Orientation returns the current orientation of the page size.
func (ps PageSize) Orientation() Orientation {
	if ps.IsLandscape() {
		return OrientationLandscape
	}

	return OrientationPortrait
}

// WidthInches returns the width in inches.
func (ps PageSize) WidthInches() float64 {
	return PointsToInches(ps.Width)
}

// HeightInches returns the height in inches.
func (ps PageSize) HeightInches() float64 {
	return PointsToInches(ps.Height)
}

// WidthMM returns the width in millimeters.
func (ps PageSize) WidthMM() float64 {
	return PointsToMM(ps.Width)
}

// HeightMM returns the height in millimeters.
func (ps PageSize) HeightMM() float64 {
	return PointsToMM(ps.Height)
}

// WidthCM returns the width in centimeters.
func (ps PageSize) WidthCM() float64 {
	return PointsToCM(ps.Width)
}

// HeightCM returns the height in centimeters.
func (ps PageSize) HeightCM() float64 {
	return PointsToCM(ps.Height)
}

// Margins represents the page margins.
type Margins struct {
	Top    float64 // Top margin in points.
	Right  float64 // Right margin in points.
	Bottom float64 // Bottom margin in points.
	Left   float64 // Left margin in points.
}

// NewMargins creates margins with the specified values in points.
func NewMargins(
	top, right, bottom, left float64,
) Margins {
	return Margins{
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Left:   left,
	}
}

// NewMarginsFromDimensions creates margins from Dimension values.
func NewMarginsFromDimensions(
	top, right, bottom, left Dimension,
) Margins {
	return Margins{
		Top:    top.ToPoints(),
		Right:  right.ToPoints(),
		Bottom: bottom.ToPoints(),
		Left:   left.ToPoints(),
	}
}

// NewMarginsUniform creates margins with the same value on all sides.
func NewMarginsUniform(margin float64) Margins {
	return Margins{
		Top:    margin,
		Right:  margin,
		Bottom: margin,
		Left:   margin,
	}
}

// NewMarginsSymmetric creates margins with symmetric horizontal and vertical values.
func NewMarginsSymmetric(
	vertical, horizontal float64,
) Margins {
	return Margins{
		Top:    vertical,
		Right:  horizontal,
		Bottom: vertical,
		Left:   horizontal,
	}
}

// NoMargins returns zero margins.
func NoMargins() Margins {
	return Margins{}
}

// DefaultMargins returns standard 1-inch margins.
func DefaultMargins() Margins {
	oneInch := InchesToPoints(1)

	return NewMarginsUniform(oneInch)
}

// NarrowMargins returns narrow 0.5-inch margins.
func NarrowMargins() Margins {
	halfInch := InchesToPoints(0.5)

	return NewMarginsUniform(halfInch)
}

// WideMargins returns wide 1.5-inch left/right and 1-inch top/bottom margins.
func WideMargins() Margins {
	return Margins{
		Top:    InchesToPoints(1),
		Right:  InchesToPoints(1.5),
		Bottom: InchesToPoints(1),
		Left:   InchesToPoints(1.5),
	}
}

// HorizontalTotal returns the sum of left and right margins.
func (m Margins) HorizontalTotal() float64 {
	return m.Left + m.Right
}

// VerticalTotal returns the sum of top and bottom margins.
func (m Margins) VerticalTotal() float64 {
	return m.Top + m.Bottom
}

// IsZero returns true if all margins are zero.
func (m Margins) IsZero() bool {
	return m.Top == 0 && m.Right == 0 &&
		m.Bottom == 0 &&
		m.Left == 0
}

// Rectangle represents a rectangular area on a page.
type Rectangle struct {
	LLX float64 // Lower-left X coordinate in points.
	LLY float64 // Lower-left Y coordinate in points.
	URX float64 // Upper-right X coordinate in points.
	URY float64 // Upper-right Y coordinate in points.
}

// NewRectangle creates a rectangle from corner coordinates.
func NewRectangle(
	llx, lly, urx, ury float64,
) Rectangle {
	return Rectangle{
		LLX: llx,
		LLY: lly,
		URX: urx,
		URY: ury,
	}
}

// NewRectangleFromSize creates a rectangle from origin and size.
func NewRectangleFromSize(
	x, y, width, height float64,
) Rectangle {
	return Rectangle{
		LLX: x,
		LLY: y,
		URX: x + width,
		URY: y + height,
	}
}

// NewRectangleForPageSize creates a rectangle for a page size starting at origin.
func NewRectangleForPageSize(
	size PageSize,
) Rectangle {
	return Rectangle{
		LLX: 0,
		LLY: 0,
		URX: size.Width,
		URY: size.Height,
	}
}

// Width returns the width of the rectangle.
func (r Rectangle) Width() float64 {
	return r.URX - r.LLX
}

// Height returns the height of the rectangle.
func (r Rectangle) Height() float64 {
	return r.URY - r.LLY
}

// ToPageSize converts the rectangle to a PageSize.
func (r Rectangle) ToPageSize() PageSize {
	return PageSize{
		Width:  r.Width(),
		Height: r.Height(),
	}
}

// toTypesRect converts to pdfcpu types.Rectangle.
//
//nolint:unused // exported for future use by rendering implementations
func (r Rectangle) toTypesRect() *types.Rectangle {
	return &types.Rectangle{
		LL: types.Point{X: r.LLX, Y: r.LLY},
		UR: types.Point{X: r.URX, Y: r.URY},
	}
}

// Array returns the rectangle as a PDF array [llx lly urx ury].
func (r Rectangle) Array() types.Array {
	return types.Array{
		types.Float(r.LLX),
		types.Float(r.LLY),
		types.Float(r.URX),
		types.Float(r.URY),
	}
}

// String returns the rectangle as a string "x y w h" for the PDF re operator.
func (r Rectangle) String() string {
	return fmt.Sprintf("%s %s %s %s",
		FormatFloat(r.LLX),
		FormatFloat(r.LLY),
		FormatFloat(r.Width()),
		FormatFloat(r.Height()))
}

// Contains returns true if the point is within the rectangle.
func (r Rectangle) Contains(x, y float64) bool {
	return x >= r.LLX && x <= r.URX &&
		y >= r.LLY &&
		y <= r.URY
}

// Inset returns a new rectangle inset by the given margins.
func (r Rectangle) Inset(
	margins Margins,
) Rectangle {
	return Rectangle{
		LLX: r.LLX + margins.Left,
		LLY: r.LLY + margins.Bottom,
		URX: r.URX - margins.Right,
		URY: r.URY - margins.Top,
	}
}

// PageBoxes represents the various page boundary boxes defined in PDF.
// See PDF Reference section 14.11.2 for details.
type PageBoxes struct {
	// MediaBox defines the boundaries of the physical medium.
	// This is required for all pages.
	MediaBox Rectangle
	// CropBox defines the visible region of the page.
	// Defaults to MediaBox if not specified.
	CropBox *Rectangle
	// BleedBox defines the region for production clipping.
	// Defaults to CropBox if not specified.
	BleedBox *Rectangle
	// TrimBox defines the intended finished page dimensions.
	// Defaults to CropBox if not specified.
	TrimBox *Rectangle
	// ArtBox defines the meaningful content area.
	// Defaults to CropBox if not specified.
	ArtBox *Rectangle
}

// NewPageBoxes creates page boxes with just the MediaBox set.
func NewPageBoxes(mediaBox Rectangle) PageBoxes {
	return PageBoxes{MediaBox: mediaBox}
}

// NewPageBoxesFromSize creates page boxes from a page size.
func NewPageBoxesFromSize(
	size PageSize,
) PageBoxes {
	return PageBoxes{
		MediaBox: NewRectangleForPageSize(size),
	}
}

// WithCropBox sets the crop box.
func (pb PageBoxes) WithCropBox(
	r Rectangle,
) PageBoxes {
	pb.CropBox = &r

	return pb
}

// WithBleedBox sets the bleed box.
func (pb PageBoxes) WithBleedBox(
	r Rectangle,
) PageBoxes {
	pb.BleedBox = &r

	return pb
}

// WithTrimBox sets the trim box.
func (pb PageBoxes) WithTrimBox(
	r Rectangle,
) PageBoxes {
	pb.TrimBox = &r

	return pb
}

// WithArtBox sets the art box.
func (pb PageBoxes) WithArtBox(
	r Rectangle,
) PageBoxes {
	pb.ArtBox = &r

	return pb
}

// WithBleed creates page boxes with a bleed margin.
// The bleed box extends beyond the trim box by the specified amount.
func (pb PageBoxes) WithBleed(
	bleedAmount float64,
) PageBoxes {
	trim := pb.EffectiveTrimBox()
	pb.BleedBox = &Rectangle{
		LLX: trim.LLX - bleedAmount,
		LLY: trim.LLY - bleedAmount,
		URX: trim.URX + bleedAmount,
		URY: trim.URY + bleedAmount,
	}

	return pb
}

// EffectiveCropBox returns the CropBox or MediaBox if not set.
func (pb PageBoxes) EffectiveCropBox() Rectangle {
	if pb.CropBox != nil {
		return *pb.CropBox
	}

	return pb.MediaBox
}

// EffectiveBleedBox returns the BleedBox or CropBox if not set.
func (pb PageBoxes) EffectiveBleedBox() Rectangle {
	if pb.BleedBox != nil {
		return *pb.BleedBox
	}

	return pb.EffectiveCropBox()
}

// EffectiveTrimBox returns the TrimBox or CropBox if not set.
func (pb PageBoxes) EffectiveTrimBox() Rectangle {
	if pb.TrimBox != nil {
		return *pb.TrimBox
	}

	return pb.EffectiveCropBox()
}

// EffectiveArtBox returns the ArtBox or CropBox if not set.
func (pb PageBoxes) EffectiveArtBox() Rectangle {
	if pb.ArtBox != nil {
		return *pb.ArtBox
	}

	return pb.EffectiveCropBox()
}

// PageOptions contains options for creating a page.
type PageOptions struct {
	// Size is the page size in points.
	Size PageSize
	// Margins are the page margins.
	Margins Margins
	// Boxes are the page boundary boxes.
	Boxes *PageBoxes
	// Rotation is the page rotation in degrees (0, 90, 180, 270).
	Rotation int
	// UserUnit is the size of a user space unit in points.
	// Default is 1.0 (1 user unit = 1 point).
	// PDF 1.6+, maximum value is 75000.
	UserUnit float64
}

// DefaultPageOptions returns default page options with A4 size and default margins.
func DefaultPageOptions() *PageOptions {
	return &PageOptions{
		Size:     PageSizeA4,
		Margins:  DefaultMargins(),
		UserUnit: 1.0,
	}
}

// NewPageOptions creates page options with the given size and default margins.
func NewPageOptions(size PageSize) *PageOptions {
	return &PageOptions{
		Size:     size,
		Margins:  DefaultMargins(),
		UserUnit: 1.0,
	}
}

// WithMargins sets the margins for the page options.
func (po *PageOptions) WithMargins(
	margins Margins,
) *PageOptions {
	po.Margins = margins

	return po
}

// WithNoMargins sets zero margins.
func (po *PageOptions) WithNoMargins() *PageOptions {
	po.Margins = NoMargins()

	return po
}

// WithRotation sets the page rotation.
func (po *PageOptions) WithRotation(
	degrees int,
) *PageOptions {
	// Normalize to valid values
	d := degrees % 360
	if d < 0 {
		d += 360
	}
	// Round to nearest 90 degrees
	po.Rotation = ((d + 45) / 90) * 90
	if po.Rotation == 360 {
		po.Rotation = 0
	}

	return po
}

// WithUserUnit sets the user unit (PDF 1.6+).
func (po *PageOptions) WithUserUnit(
	unit float64,
) *PageOptions {
	u := unit
	if u < 1.0 {
		u = 1.0
	}
	if u > 75000 {
		u = 75000
	}
	po.UserUnit = u

	return po
}

// WithBoxes sets the page boxes.
func (po *PageOptions) WithBoxes(
	boxes PageBoxes,
) *PageOptions {
	po.Boxes = &boxes

	return po
}

// Landscape returns page options with landscape orientation.
func (po *PageOptions) Landscape() *PageOptions {
	po.Size = po.Size.Landscape()

	return po
}

// Portrait returns page options with portrait orientation.
func (po *PageOptions) Portrait() *PageOptions {
	po.Size = po.Size.Portrait()

	return po
}

// ContentArea returns the rectangle representing the area inside the margins.
func (po *PageOptions) ContentArea() Rectangle {
	return Rectangle{
		LLX: po.Margins.Left,
		LLY: po.Margins.Bottom,
		URX: po.Size.Width - po.Margins.Right,
		URY: po.Size.Height - po.Margins.Top,
	}
}

// ContentWidth returns the width of the content area (page width minus margins).
func (po *PageOptions) ContentWidth() float64 {
	return po.Size.Width - po.Margins.HorizontalTotal()
}

// ContentHeight returns the height of the content area (page height minus margins).
func (po *PageOptions) ContentHeight() float64 {
	return po.Size.Height - po.Margins.VerticalTotal()
}

// CoordinateSystem represents the coordinate system configuration for a page.
type CoordinateSystem struct {
	// Origin specifies where the origin is located.
	Origin CoordinateOrigin
	// UserUnit is the size of a user space unit in points.
	UserUnit float64
	// PageHeight is the page height in points (needed for origin transformation).
	PageHeight float64
}

// CoordinateOrigin specifies the location of the coordinate system origin.
type CoordinateOrigin int

const (
	// OriginBottomLeft is the default PDF origin (bottom-left corner).
	OriginBottomLeft CoordinateOrigin = iota
	// OriginTopLeft is the OOXML origin (top-left corner).
	OriginTopLeft
)

// NewCoordinateSystem creates a new coordinate system for a page.
func NewCoordinateSystem(
	pageHeight float64,
) *CoordinateSystem {
	return &CoordinateSystem{
		Origin:     OriginBottomLeft,
		UserUnit:   1.0,
		PageHeight: pageHeight,
	}
}

// WithTopLeftOrigin sets the origin to top-left (for OOXML compatibility).
func (cs *CoordinateSystem) WithTopLeftOrigin() *CoordinateSystem {
	cs.Origin = OriginTopLeft

	return cs
}

// WithUserUnit sets the user unit.
func (cs *CoordinateSystem) WithUserUnit(
	unit float64,
) *CoordinateSystem {
	cs.UserUnit = unit

	return cs
}

// TransformY transforms a Y coordinate based on the origin setting.
// If origin is top-left, Y is flipped to PDF's bottom-left origin.
func (cs *CoordinateSystem) TransformY(
	y float64,
) float64 {
	if cs.Origin == OriginTopLeft {
		return cs.PageHeight - y
	}

	return y
}

// TransformPoint transforms a point based on the coordinate system.
func (cs *CoordinateSystem) TransformPoint(
	x, y float64,
) (tx, ty float64) {
	return x, cs.TransformY(y)
}

// ScaleToUserUnits scales a point value to user units.
func (cs *CoordinateSystem) ScaleToUserUnits(
	value float64,
) float64 {
	return value / cs.UserUnit
}

// ScaleFromUserUnits scales from user units to points.
func (cs *CoordinateSystem) ScaleFromUserUnits(
	value float64,
) float64 {
	return value * cs.UserUnit
}

// LineCap specifies the shape of the endpoints for an open path.
// This controls how line ends are rendered when stroking paths.
// The cap style affects only the endpoints of open subpaths and does not
// affect the joins between connected segments or the shape of closed subpaths.
type LineCap int

const (
	// LineCapButt truncates the line exactly at the endpoint.
	// The stroke is squared off at the endpoint of the path with no projection.
	LineCapButt LineCap = 0
	// LineCapRound rounds the endpoint with a semicircular arc.
	// A semicircular arc with diameter equal to the line width is drawn around
	// the endpoint and filled in.
	LineCapRound LineCap = 1
	// LineCapSquare extends the line with a square projection.
	// The stroke continues beyond the endpoint by half the line width and is
	// squared off.
	LineCapSquare LineCap = 2
)

// LineJoin specifies the shape at the corners of stroked paths.
// This controls how the junction of two line segments is rendered when stroking paths.
type LineJoin int

const (
	// LineJoinMiter creates pointed corners.
	// The outer edges of the strokes extend until they meet at an angle.
	// If the angle is too sharp (exceeds the miter limit), a bevel join is used instead.
	LineJoinMiter LineJoin = 0
	// LineJoinRound creates rounded corners.
	// A circular arc with diameter equal to the line width is drawn around
	// the point where the segments meet.
	LineJoinRound LineJoin = 1
	// LineJoinBevel creates beveled corners.
	// The two segments are finished with butt caps and the notch between them
	// is filled with a triangle.
	LineJoinBevel LineJoin = 2
)

// PageImage represents an image resource that can be added to a page.
// Implementations provide access to image dimensions and raw data for embedding
// into PDF documents.
//
// The image data format depends on the specific implementation but typically
// includes formats like JPEG, PNG, or raw pixel data.
type PageImage interface {
	// Width returns the image width in pixels.
	Width() int
	// Height returns the image height in pixels.
	Height() int
	// Data returns the raw image data in the implementation-specific format.
	// For JPEG images this would be the JPEG-encoded data,
	// for PNG this would be the PNG-encoded data, etc.
	Data() []byte
}

// PageDrawer defines the drawing API for PDF pages.
// It provides methods for drawing shapes, setting colors, managing graphics state,
// and rendering text and images. The interface abstracts PDF page content stream
// operations, allowing for both actual PDF generation and testing with mock implementations.
//
// Coordinate System:
// By default, PDF uses a bottom-left origin coordinate system where:
//   - X increases to the right
//   - Y increases upward
//   - Units are in points (1/72 inch)
//
// Use CoordinateSystem and Transform methods to work with different coordinate systems
// (e.g., top-left origin for OOXML compatibility).
//
// Graphics State:
// Drawing operations are affected by the current graphics state, which includes:
//   - Fill and stroke colors
//   - Line width, dash pattern, cap, and join styles
//   - Current transformation matrix
//   - Font and font size
//
// Use SaveGraphicsState and RestoreGraphicsState to preserve and restore state
// when making temporary changes.
//
// Common Usage Pattern:
//
//	page := doc.AddPage(core.PageSizeA4)
//	page.SetFillColor(0, 0, 0)        // Black fill
//	page.SetStrokeColor(1, 0, 0)      // Red stroke
//	page.SetLineWidth(2)
//	page.DrawRectangle(100, 100, 200, 150, true, true)  // Filled and stroked
//
// Implementations:
//   - PageImpl: Actual PDF page generation
//   - MockPage: Testing without PDF generation
type PageDrawer interface {
	// DrawRectangle draws a rectangle at the specified position.
	// Parameters x and y define the bottom-left corner of the rectangle in user space units.
	// The width and height define the rectangle dimensions.
	// If fill is true, the rectangle is filled using the current fill color.
	// If stroke is true, the rectangle outline is drawn using the current stroke color and line width.
	// Both fill and stroke can be true to create a filled rectangle with an outline.
	//
	// Example:
	//   page.SetFillColor(0.9, 0.9, 0.9)  // Light gray
	//   page.SetStrokeColor(0, 0, 0)      // Black
	//   page.SetLineWidth(1)
	//   page.DrawRectangle(50, 50, 100, 75, true, true)  // Gray rectangle with black border
	DrawRectangle(
		x, y, width, height float64,
		fill, stroke bool,
	)

	// DrawCircle draws a circle at the specified center point.
	// Parameters cx and cy define the center of the circle in user space units.
	// The radius defines the circle size.
	// If fill is true, the circle is filled using the current fill color.
	// If stroke is true, the circle outline is drawn using the current stroke color and line width.
	//
	// The circle is approximated using Bézier curves as PDF does not have a native circle primitive.
	//
	// Example:
	//   page.SetFillColor(1, 0, 0)  // Red
	//   page.DrawCircle(200, 200, 50, true, false)  // Filled red circle, no outline
	DrawCircle(
		cx, cy, radius float64,
		fill, stroke bool,
	)

	// DrawEllipse draws an ellipse at the specified center point.
	// Parameters cx and cy define the center of the ellipse in user space units.
	// The rx and ry parameters define the horizontal and vertical radii respectively.
	// If fill is true, the ellipse is filled using the current fill color.
	// If stroke is true, the ellipse outline is drawn using the current stroke color and line width.
	//
	// The ellipse is approximated using Bézier curves as PDF does not have a native ellipse primitive.
	//
	// Example:
	//   page.SetFillColor(0, 0, 1)  // Blue
	//   page.DrawEllipse(300, 300, 80, 40, true, true)  // Blue ellipse with outline
	DrawEllipse(
		cx, cy, rx, ry float64,
		fill, stroke bool,
	)

	// SetFillColor sets the color used for fill operations.
	// Parameters r, g, b are red, green, blue components in the range [0, 1].
	// This affects DrawRectangle, DrawCircle, DrawEllipse when fill is true,
	// as well as filled text rendering.
	//
	// Example:
	//   page.SetFillColor(1, 0, 0)     // Pure red
	//   page.SetFillColor(0.5, 0.5, 0.5)  // Medium gray
	//   page.SetFillColor(0, 0, 0)     // Black
	SetFillColor(r, g, b float64)

	// SetStrokeColor sets the color used for stroke (outline) operations.
	// Parameters r, g, b are red, green, blue components in the range [0, 1].
	// This affects DrawRectangle, DrawCircle, DrawEllipse when stroke is true,
	// as well as line and path stroking operations.
	//
	// Example:
	//   page.SetStrokeColor(0, 0, 1)  // Blue outlines
	SetStrokeColor(r, g, b float64)

	// SetLineWidth sets the width of stroked lines in user space units.
	// This affects all subsequent stroke operations including shape outlines.
	// The default line width is typically 1 point.
	//
	// Example:
	//   page.SetLineWidth(0.5)   // Thin line
	//   page.SetLineWidth(3)     // Thick line
	SetLineWidth(width float64)

	// SetLineDashPattern sets the dash pattern for stroked lines.
	// The pattern parameter is an array of dash and gap lengths.
	// The phase parameter specifies the offset into the pattern at which to start.
	//
	// An empty pattern array creates a solid line.
	// Even-indexed entries are dash lengths, odd-indexed entries are gap lengths.
	//
	// Examples:
	//   page.SetLineDashPattern([]float64{3, 1}, 0)        // 3 on, 1 off
	//   page.SetLineDashPattern([]float64{5, 2, 1, 2}, 0)  // 5 on, 2 off, 1 on, 2 off
	//   page.SetLineDashPattern([]float64{}, 0)            // Solid line
	//   page.SetLineDashPattern([]float64{3, 1}, 1.5)      // Same pattern, offset by 1.5
	SetLineDashPattern(
		pattern []float64,
		phase float64,
	)

	// SetLineCap sets the line cap style for the endpoints of stroked paths.
	// See LineCap constants for available styles (Butt, Round, Square).
	// The default is typically LineCapButt.
	//
	// Example:
	//   page.SetLineCap(core.LineCapRound)  // Rounded line ends
	SetLineCap(lineCap LineCap)

	// SetLineJoin sets the line join style for corners in stroked paths.
	// See LineJoin constants for available styles (Miter, Round, Bevel).
	// The default is typically LineJoinMiter.
	//
	// Example:
	//   page.SetLineJoin(core.LineJoinRound)  // Rounded corners
	SetLineJoin(lineJoin LineJoin)

	// SaveGraphicsState saves the current graphics state on the PDF graphics state stack.
	// This includes all current color, line, font, and transformation settings.
	// Use RestoreGraphicsState to return to this state later.
	//
	// Graphics state saves/restores must be balanced (each save must have a corresponding restore).
	// This is commonly used to make temporary changes without affecting the surrounding context.
	//
	// Example:
	//   page.SaveGraphicsState()
	//   page.SetFillColor(1, 0, 0)      // Red, temporarily
	//   page.DrawRectangle(x, y, w, h, true, false)
	//   page.RestoreGraphicsState()     // Back to previous fill color
	SaveGraphicsState()

	// RestoreGraphicsState restores the graphics state to the most recently saved state.
	// This pops the graphics state stack and restores all settings (colors, line styles,
	// transformation matrix, etc.) to their saved values.
	//
	// Must be paired with a previous SaveGraphicsState call.
	RestoreGraphicsState()

	// Transform applies a transformation matrix to the current transformation matrix.
	// The matrix parameter is a 6-element array [a b c d e f] representing:
	//   x' = a*x + c*y + e
	//   y' = b*x + d*y + f
	//
	// Common transformations:
	//   Translation:  [1 0 0 1 tx ty]
	//   Scaling:      [sx 0 0 sy 0 0]
	//   Rotation:     [cos(θ) sin(θ) -sin(θ) cos(θ) 0 0]
	//
	// Transformations are cumulative and affect all subsequent drawing operations
	// until restored via RestoreGraphicsState.
	//
	// Example:
	//   page.SaveGraphicsState()
	//   page.Transform(PDFMatrix{1, 0, 0, 1, 100, 200})  // Translate by (100, 200)
	//   // Draw operations here are offset by (100, 200)
	//   page.RestoreGraphicsState()
	Transform(matrix PDFMatrix)

	// AddImage adds an image to the page at the specified position and size.
	// The img parameter provides the image data and dimensions.
	// Parameters x and y specify the bottom-left corner of the image in user space units.
	// Parameters width and height specify the rendered size of the image.
	//
	// The image is scaled to fit the specified dimensions.
	// Aspect ratio is not automatically preserved - specify width and height accordingly.
	//
	// Example:
	//   img := loadImage("logo.png")  // Returns a PageImage
	//   page.AddImage(img, 50, 700, 100, 50)  // Place at (50, 700), size 100x50 points
	AddImage(
		img PageImage,
		x, y, width, height float64,
	)

	// DrawText draws text at the specified position.
	// Parameters x and y specify the baseline start position in user space units.
	// The text parameter is the string to render.
	//
	// The current font and font size (set via SetFont) determine the text appearance.
	// The current fill color determines the text color.
	//
	// Note: This is a simple text drawing method. For complex text layout with
	// line breaking, alignment, and formatting, use the text layout engine in pdf/layout.
	//
	// Example:
	//   page.SetFont("Helvetica", 12)
	//   page.SetFillColor(0, 0, 0)  // Black
	//   page.DrawText(100, 500, "Hello, World!")
	DrawText(x, y float64, text string)

	// SetFont sets the current font and size for text drawing.
	// The name parameter is the font name (e.g., "Helvetica", "Times-Roman", "Courier").
	// The size parameter is the font size in points.
	//
	// Standard PDF fonts (always available):
	//   - Helvetica, Helvetica-Bold, Helvetica-Oblique, Helvetica-BoldOblique
	//   - Times-Roman, Times-Bold, Times-Italic, Times-BoldItalic
	//   - Courier, Courier-Bold, Courier-Oblique, Courier-BoldOblique
	//   - Symbol, ZapfDingbats
	//
	// Example:
	//   page.SetFont("Helvetica-Bold", 14)
	//   page.SetFont("Times-Roman", 12)
	SetFont(name string, size float64)

	// WriteContent writes raw PDF content stream operators directly to the page.
	// This is a low-level method for advanced use cases where the high-level
	// drawing methods are insufficient.
	//
	// The content parameter should be valid PDF content stream syntax.
	// Use with caution as invalid content can corrupt the PDF.
	//
	// Example:
	//   page.WriteContent("0.5 g")  // Set gray fill color to 50%
	//   page.WriteContent("100 200 m 200 200 l S")  // Draw a line
	WriteContent(content string)
}
