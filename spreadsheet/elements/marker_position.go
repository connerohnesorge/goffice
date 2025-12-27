package elements

import (
	"strconv"

	"github.com/connerohnesorge/goffice/openxml"
)

// Position represents the position element (xdr:pos).
// This specifies the absolute position in EMUs for an absolute anchor.
type Position struct {
	*openxml.CompositeElementBase
}

// NewPosition creates a new Position element.
func NewPosition() *Position {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"pos",
		PrefixXDR,
	)

	return &Position{CompositeElementBase: elem}
}

// X returns the X coordinate in EMUs.
func (p *Position) X() int64 {
	attr, found := p.GetAttribute("x", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize64,
	)

	return val
}

// SetX sets the X coordinate in EMUs.
func (p *Position) SetX(x int64) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"x",
			"",
			strconv.FormatInt(x, parseBase10),
		),
	)
}

// Y returns the Y coordinate in EMUs.
func (p *Position) Y() int64 {
	attr, found := p.GetAttribute("y", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize64,
	)

	return val
}

// SetY sets the Y coordinate in EMUs.
func (p *Position) SetY(y int64) {
	p.SetAttribute(
		openxml.NewAttribute(
			"",
			"y",
			"",
			strconv.FormatInt(y, parseBase10),
		),
	)
}

// SetPosition sets both X and Y coordinates.
func (p *Position) SetPosition(x, y int64) {
	p.SetX(x)
	p.SetY(y)
}

// Clone creates a deep copy of this Position element.
func (p *Position) Clone() openxml.Element {
	cloned := p.CompositeElementBase.Clone()

	return &Position{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Position element.
func (p *Position) CloneNode(
	deep bool,
) openxml.Element {
	cloned := p.CompositeElementBase.CloneNode(
		deep,
	)

	return &Position{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// Extent represents the extent element (xdr:ext).
// This specifies the size in EMUs.
type Extent struct {
	*openxml.CompositeElementBase
}

// NewExtent creates a new Extent element.
func NewExtent() *Extent {
	elem := openxml.NewCompositeElement(
		NamespaceSpreadsheetDrawing,
		"ext",
		PrefixXDR,
	)

	return &Extent{CompositeElementBase: elem}
}

// Cx returns the width in EMUs.
func (e *Extent) Cx() int64 {
	attr, found := e.GetAttribute("cx", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize64,
	)

	return val
}

// SetCx sets the width in EMUs.
func (e *Extent) SetCx(cx int64) {
	e.SetAttribute(
		openxml.NewAttribute(
			"",
			"cx",
			"",
			strconv.FormatInt(cx, parseBase10),
		),
	)
}

// Cy returns the height in EMUs.
func (e *Extent) Cy() int64 {
	attr, found := e.GetAttribute("cy", "")
	if !found {
		return 0
	}
	val, _ := strconv.ParseInt(
		attr.Value(),
		parseBase10,
		bitSize64,
	)

	return val
}

// SetCy sets the height in EMUs.
func (e *Extent) SetCy(cy int64) {
	e.SetAttribute(
		openxml.NewAttribute(
			"",
			"cy",
			"",
			strconv.FormatInt(cy, parseBase10),
		),
	)
}

// SetSize sets both width and height.
func (e *Extent) SetSize(cx, cy int64) {
	e.SetCx(cx)
	e.SetCy(cy)
}

// Clone creates a deep copy of this Extent element.
func (e *Extent) Clone() openxml.Element {
	cloned := e.CompositeElementBase.Clone()

	return &Extent{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// CloneNode creates a copy of this Extent element.
func (e *Extent) CloneNode(
	deep bool,
) openxml.Element {
	cloned := e.CompositeElementBase.CloneNode(
		deep,
	)

	return &Extent{
		CompositeElementBase: cloned.(*openxml.CompositeElementBase),
	}
}

// EMU conversion constants.
const (
	// EMUsPerInch is the number of EMUs per inch.
	EMUsPerInch int64 = 914400

	// EMUsPerCm is the number of EMUs per centimeter.
	EMUsPerCm int64 = 360000

	// EMUsPerPoint is the number of EMUs per point.
	EMUsPerPoint int64 = 12700

	// EMUsPerPixel is the number of EMUs per pixel (at 96 DPI).
	EMUsPerPixel int64 = 9525
)

// InchesToEMUs converts inches to EMUs.
func InchesToEMUs(inches float64) int64 {
	return int64(inches * float64(EMUsPerInch))
}

// EMUsToInches converts EMUs to inches.
func EMUsToInches(emus int64) float64 {
	return float64(emus) / float64(EMUsPerInch)
}

// CmToEMUs converts centimeters to EMUs.
func CmToEMUs(cm float64) int64 {
	return int64(cm * float64(EMUsPerCm))
}

// EMUsToCm converts EMUs to centimeters.
func EMUsToCm(emus int64) float64 {
	return float64(emus) / float64(EMUsPerCm)
}

// PointsToEMUs converts points to EMUs.
func PointsToEMUs(points float64) int64 {
	return int64(points * float64(EMUsPerPoint))
}

// EMUsToPoints converts EMUs to points.
func EMUsToPoints(emus int64) float64 {
	return float64(emus) / float64(EMUsPerPoint)
}

// PixelsToEMUs converts pixels to EMUs (at 96 DPI).
func PixelsToEMUs(pixels int) int64 {
	return int64(pixels) * EMUsPerPixel
}

// EMUsToPixels converts EMUs to pixels (at 96 DPI).
func EMUsToPixels(emus int64) int {
	return int(emus / EMUsPerPixel)
}
